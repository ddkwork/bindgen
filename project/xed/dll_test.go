package xed

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

func newXed(t *testing.T) *Xed {
	x := &Xed{}
	t.Cleanup(func() { _ = x })
	return x
}

func TestTablesInit(t *testing.T) {
	x := newXed(t)
	x.TablesInit()
}

func TestDecodeNop(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	var xedd Xed_decoded_inst_t
	x.DecodedInstZeroSetMode(&xedd, &Xed_state_t{
		Mmode:          XedMachineModeLong64,
		StackAddrWidth: XedAddressWidth64b,
	})

	itext := []byte{0x90}
	err := x.Decode(&xedd, &itext[0], uint32(len(itext)))
	if err != XedErrorNone {
		t.Fatalf("decode NOP failed: %d", err)
	}

	iclass := x.OperandValuesGetIclass(&xedd)
	if iclass != XedIclassNop {
		t.Errorf("expected NOP iclass, got %s", cstring(x.IclassEnumT2str(iclass)))
	}
	length := xedd.DecodedLength
	if length != 1 {
		t.Errorf("expected length 1, got %d", length)
	}
}

func TestDecodeMovRegImm64(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	var xedd Xed_decoded_inst_t
	x.DecodedInstZeroSetMode(&xedd, &Xed_state_t{
		Mmode:          XedMachineModeLong64,
		StackAddrWidth: XedAddressWidth64b,
	})

	itext := []byte{0x48, 0xB8, 0x78, 0x56, 0x34, 0x12, 0xEF, 0xCD, 0xAB, 0x89}
	err := x.Decode(&xedd, &itext[0], uint32(len(itext)))
	if err != XedErrorNone {
		t.Fatalf("decode MOV RAX, imm64 failed: %d", err)
	}

	iclass := x.OperandValuesGetIclass(&xedd)
	if iclass != XedIclassMov {
		t.Errorf("expected MOV iclass, got %s", cstring(x.IclassEnumT2str(iclass)))
	}
	length := xedd.DecodedLength
	if length != 10 {
		t.Errorf("expected length 10, got %d", length)
	}
	imm := x.OperandValuesGetImmediateUint64(&xedd)
	wantImm := uint64(0x89ABCDEF12345678)
	if imm != wantImm {
		t.Errorf("expected imm 0x%016X, got 0x%016X", wantImm, imm)
	}
}

func TestDecodeAddRegMem32(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	var xedd Xed_decoded_inst_t
	x.DecodedInstZeroSetMode(&xedd, &Xed_state_t{
		Mmode:          XedMachineModeLegacy32,
		StackAddrWidth: XedAddressWidth32b,
	})

	itext := []byte{0x03, 0x05, 0x78, 0x56, 0x34, 0x12}
	err := x.Decode(&xedd, &itext[0], uint32(len(itext)))
	if err != XedErrorNone {
		t.Fatalf("decode ADD EAX, [0x12345678] failed: %d", err)
	}

	iclass := x.OperandValuesGetIclass(&xedd)
	if iclass != XedIclassAdd {
		t.Errorf("expected ADD iclass, got %s", cstring(x.IclassEnumT2str(iclass)))
	}
}

func TestFormatIntel(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	var xedd Xed_decoded_inst_t
	x.DecodedInstZeroSetMode(&xedd, &Xed_state_t{
		Mmode:          XedMachineModeLong64,
		StackAddrWidth: XedAddressWidth64b,
	})

	testCases := []struct {
		bytes     []byte
		expectSub string
	}{
		{[]byte{0x90}, "NOP"},
		{[]byte{0x48, 0x89, 0xC8}, "MOV"},
		{[]byte{0xFF, 0xC0}, "INC"},
		{[]byte{0x01, 0xD8}, "ADD"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%X", tc.bytes), func(t *testing.T) {
			x.DecodedInstZeroKeepMode(&xedd)
			err := x.Decode(&xedd, &tc.bytes[0], uint32(len(tc.bytes)))
			if err != XedErrorNone {
				t.Fatalf("decode failed: %d", err)
			}

			buf := make([]int8, 256)
			ok := x.FormatContext(XedSyntaxIntel, &xedd, &buf[0], int32(len(buf)), 0, unsafe.Pointer(uintptr(0)), Xed_disassembly_callback_fn_t(nil))
			if ok == 0 {
				t.Fatal("format_context returned false")
			}

			result := goStringInt8(buf)
			if !strings.Contains(strings.ToUpper(result), tc.expectSub) {
				t.Errorf("expected output containing %q, got %q", tc.expectSub, result)
			}
		})
	}
}

func testEncodeRoundTrip(t *testing.T, x *Xed, inst Xed_encoder_instruction_t, expectBytes int) ([]byte, string) {
	var encReq Xed_encoder_request_t
	x.EncoderRequestZeroSetMode(&encReq, &inst.Mode)

	ok := x.ConvertToEncoderRequest(&encReq, &inst)
	if ok == 0 {
		t.Fatalf("convert_to_encoder_request failed")
	}

	itext := make([]byte, XedMaxInstructionBytes)
	ilen := uint32(len(itext))
	var olen uint32

	err := x.Encode(&encReq, &itext[0], ilen, &olen)
	if err != XedErrorNone {
		t.Fatalf("encode error: %s (%d)", cstring(x.ErrorEnumT2str(err)), err)
	}

	if olen == 0 || olen > XedMaxInstructionBytes {
		t.Fatalf("bad encoded length: %d", olen)
	}

	result := itext[:olen]

	var xedd Xed_decoded_inst_t
	x.DecodedInstZeroSetMode(&xedd, &inst.Mode)
	decErr := x.Decode(&xedd, &result[0], olen)
	if decErr != XedErrorNone {
		t.Fatalf("round-trip decode error: %s (%d)", cstring(x.ErrorEnumT2str(decErr)), decErr)
	}

	buf := make([]int8, 256)
	fmtOk := x.FormatContext(XedSyntaxIntel, &xedd, &buf[0], int32(len(buf)), 0, unsafe.Pointer(uintptr(0)), Xed_disassembly_callback_fn_t(nil))
	disasm := ""
	if fmtOk != 0 {
		disasm = goStringInt8(buf)
	} else {
		disasm = "<format error>"
	}

	return result, disasm
}

func TestEncodeJmpRelbr64(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	dstate := xedState(XedMachineModeLong64)
	inst := xedBuildInst(dstate, XedIclassJmp, 64,
		xedEncOpRelBr(0x11223344, 32),
	)

	result, disasm := testEncodeRoundTrip(t, x, inst, 5)
	t.Logf("encoded: %X -> %s", result, disasm)
}

func TestEncodeAddRegImm32(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	dstate := xedState(XedMachineModeLong64)
	inst := xedBuildInst(dstate, XedIclassAdd, 64,
		xedEncOpReg(XedRegRax),
		xedEncOpImm0(0x77, 8),
	)

	result, disasm := testEncodeRoundTrip(t, x, inst, 3)
	t.Logf("encoded: %X -> %s", result, disasm)
}

func TestEncodeAddRegImm32_32bit(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	dstate := xedState(XedMachineModeLegacy32)
	inst := xedBuildInst(dstate, XedIclassAdd, 0,
		xedEncOpReg(XedRegEax),
		xedEncOpImm0(0x44332211, 32),
	)

	result, disasm := testEncodeRoundTrip(t, x, inst, 5)
	t.Logf("encoded: %X -> %s", result, disasm)
}

func TestEncodeXorRegReg64(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	dstate := xedState(XedMachineModeLong64)
	inst := xedBuildInst(dstate, XedIclassXor, 64,
		xedEncOpReg(XedRegRcx),
		xedEncOpReg(XedRegRdx),
	)

	result, disasm := testEncodeRoundTrip(t, x, inst, 3)
	t.Logf("encoded: %X -> %s", result, disasm)
}

func TestEncodePushReg(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	dstate := xedState(XedMachineModeLegacy32)
	inst := xedBuildInst(dstate, XedIclassPush, 0,
		xedEncOpReg(XedRegEcx),
	)

	result, disasm := testEncodeRoundTrip(t, x, inst, 1)
	t.Logf("encoded: %X -> %s", result, disasm)
}

func TestEncodeMovRegMemDisp64(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	dstate := xedState(XedMachineModeLong64)
	inst := xedBuildInst(dstate, XedIclassMov, 64,
		xedEncOpReg(XedRegRax),
		xedEncOpMemBd(XedRegInvalid, 0x1122334455667788, 64, 64),
	)

	result, disasm := testEncodeRoundTrip(t, x, inst, 10)
	t.Logf("encoded: %X -> %s", result, disasm)
}

func TestEncodeLeaDispOnly(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	dstate := xedState(XedMachineModeLegacy32)
	inst := xedBuildInst(dstate, XedIclassLea, 32,
		xedEncOpReg(XedRegEax),
		xedEncOpMemBd(XedRegInvalid, 0x11223344, 32, 32),
	)

	result, disasm := testEncodeRoundTrip(t, x, inst, 7)
	t.Logf("encoded: %X -> %s", result, disasm)
}

func TestEncodeMultipleInstructions(t *testing.T) {
	x := newXed(t)
	x.TablesInit()

	type testCase struct {
		name    string
		mode    Xed_machine_mode_enum_t
		eosz    uint32
		iclass  Xed_iclass_enum_t
		ops     []encOp
		setAddr bool
		addrVal uint32
	}

	tests := []testCase{
		{"JMP rel64", XedMachineModeLong64, 64, XedIclassJmp, []encOp{xedEncOpRelBr(0x11223344, 32)}, false, 0},
		{"XOR reg32", XedMachineModeLegacy32, 0, XedIclassXor, []encOp{xedEncOpReg(XedRegEcx), xedEncOpReg(XedRegEdx)}, false, 0},
		{"XOR reg64", XedMachineModeLong64, 64, XedIclassXor, []encOp{xedEncOpReg(XedRegRcx), xedEncOpReg(XedRegRdx)}, false, 0},
		{"PUSH reg32", XedMachineModeLegacy32, 0, XedIclassPush, []encOp{xedEncOpReg(XedRegEcx)}, false, 0},
		{"PUSH reg64", XedMachineModeLong64, 64, XedIclassPush, []encOp{xedEncOpReg(XedRegRcx)}, false, 0},
		{"ADD EAX, imm8", XedMachineModeLong64, 64, XedIclassAdd, []encOp{xedEncOpReg(XedRegRax), xedEncOpImm0(0x77, 8)}, false, 0},
		{"ADD EAX, imm32", XedMachineModeLong64, 64, XedIclassAdd, []encOp{xedEncOpReg(XedRegRax), xedEncOpImm0(0x44332211, 32)}, false, 0},
		{"MOV RAX, [rip+disp]", XedMachineModeLong64, 64, XedIclassMov, []encOp{xedEncOpReg(XedRegRax), xedEncOpMemBd(XedRegInvalid, 0x11223344, 32, 64)}, false, 0},
		{"MOV CR3,RDI", XedMachineModeLong64, 64, XedIclassMovCr, []encOp{xedEncOpReg(XedRegCr3), xedEncOpReg(XedRegRdi)}, false, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			dstate := xedState(tc.mode)
			inst := xedBuildInst(dstate, tc.iclass, tc.eosz, tc.ops...)
			if tc.setAddr {
				xedSetAddr(&inst, tc.addrVal)
			}
			result, disasm := testEncodeRoundTrip(t, x, inst, 0)
			t.Logf("encoded: %v -> %s", result, disasm)
		})
	}
}

func TestVersion(t *testing.T) {
	x := newXed(t)
	x.TablesInit()
	ver := x.GetVersion()
	if ver == nil {
		t.Fatal("get_version returned nil")
	}
	s := goString((*[256]byte)(unsafe.Pointer(ver))[:])
	t.Logf("XED version: %s", s)
	if len(s) == 0 {
		t.Error("version string is empty")
	}
}

func TestRegisterNames(t *testing.T) {
	x := newXed(t)
	registers := []struct {
		reg  Xed_reg_enum_t
		want string
	}{
		{XedRegRax, "RAX"},
		{XedRegRcx, "RCX"},
		{XedRegRdx, "RDX"},
		{XedRegRbx, "RBX"},
		{XedRegRsp, "RSP"},
		{XedRegRbp, "RBP"},
		{XedRegRsi, "RSI"},
		{XedRegRdi, "RDI"},
		{XedRegEax, "EAX"},
		{XedRegEcx, "ECX"},
		{XedRegEdx, "EDX"},
	}

	for _, tc := range registers {
		t.Run(tc.want, func(t *testing.T) {
			s := x.RegEnumT2str(tc.reg)
			if s == nil {
				t.Fatalf("reg_enum_t2str(%d) returned nil", tc.reg)
			}
			name := cstring(s)
			if name != tc.want {
				t.Errorf("got %q want %q", name, tc.want)
			}
		})
	}
}

func newAssembler(t *testing.T) *Assembler {
	x := newXed(t)
	x.TablesInit()
	return NewAssembler(x)
}

func TestAssembleNop(t *testing.T) {
	a := newAssembler(t)
	bytes, err := a.Assemble("NOP", 64)
	if err != nil {
		t.Fatalf("assemble NOP: %v", err)
	}
	if len(bytes) != 1 || bytes[0] != 0x90 {
		t.Errorf("expected [0x90], got %X", bytes)
	}
}

func TestAssembleMovRegImm64(t *testing.T) {
	a := newAssembler(t)
	bytes, err := a.Assemble("MOV RAX, 0x1122334455667788", 64)
	if err != nil {
		t.Fatalf("assemble MOV RAX, imm64: %v", err)
	}
	t.Logf("MOV RAX, 0x1122334455667788 -> %X", bytes)
	disasm, err := a.Disassemble(bytes, 64)
	if err != nil {
		t.Fatalf("disassemble: %v", err)
	}
	t.Logf("disasm: %s", disasm)
}

func TestAssembleXorRegReg(t *testing.T) {
	a := newAssembler(t)
	bytes, err := a.Assemble("XOR RCX, RDX", 64)
	if err != nil {
		t.Fatalf("assemble XOR RCX, RDX: %v", err)
	}
	t.Logf("XOR RCX, RDX -> %X", bytes)
	disasm, err := a.Disassemble(bytes, 64)
	if err != nil {
		t.Fatalf("disassemble: %v", err)
	}
	t.Logf("disasm: %s", disasm)
	if !strings.Contains(strings.ToUpper(disasm), "XOR") {
		t.Errorf("expected XOR in disassembly, got %q", disasm)
	}
}

func TestAssembleAddRegImm(t *testing.T) {
	a := newAssembler(t)
	bytes, err := a.Assemble("ADD RAX, 0x77", 64)
	if err != nil {
		t.Fatalf("assemble ADD RAX, 0x77: %v", err)
	}
	t.Logf("ADD RAX, 0x77 -> %X", bytes)
	disasm, err := a.Disassemble(bytes, 64)
	if err != nil {
		t.Fatalf("disassemble: %v", err)
	}
	t.Logf("disasm: %s", disasm)
}

func TestAssemblePush(t *testing.T) {
	a := newAssembler(t)
	bytes, err := a.Assemble("PUSH RAX", 64)
	if err != nil {
		t.Fatalf("assemble PUSH RAX: %v", err)
	}
	t.Logf("PUSH RAX -> %X", bytes)
}

func TestAssembleJmp(t *testing.T) {
	a := newAssembler(t)
	bytes, err := a.Assemble("JMP 0x11223344", 64)
	if err != nil {
		t.Fatalf("assemble JMP: %v", err)
	}
	t.Logf("JMP 0x11223344 -> %X", bytes)
	disasm, err := a.Disassemble(bytes, 64)
	if err != nil {
		t.Fatalf("disassemble: %v", err)
	}
	t.Logf("disasm: %s", disasm)
}

func TestAssembleRet(t *testing.T) {
	a := newAssembler(t)
	bytes, err := a.Assemble("RET", 64)
	if err != nil {
		t.Fatalf("assemble RET: %v", err)
	}
	if len(bytes) != 1 || bytes[0] != 0xC3 {
		t.Errorf("expected [0xC3], got %X", bytes)
	}
}

func TestAssembleMovRegMem(t *testing.T) {
	a := newAssembler(t)
	encoded, err := a.Assemble("MOV RAX, QWORD PTR [RIP+0x11223344]", 64)
	if err != nil {
		t.Fatalf("assemble MOV RAX, [RIP+disp]: %v", err)
	}
	expected := []byte{0x48, 0x8B, 0x05, 0x44, 0x33, 0x22, 0x11}
	if !bytes.Equal(encoded, expected) {
		t.Errorf("expected %X, got %X", expected, encoded)
	}
	disasm, err := a.Disassemble(encoded, 64)
	if err != nil {
		t.Fatalf("disassemble: %v", err)
	}
	if !strings.Contains(strings.ToUpper(disasm), "MOV") {
		t.Errorf("disasm should contain MOV: %s", disasm)
	}
}

func TestAssembleMultiple(t *testing.T) {
	a := newAssembler(t)
	tests := []struct {
		line string
		mode int
	}{
		{"NOP", 64},
		{"XOR ECX, EDX", 32},
		{"XOR RCX, RDX", 64},
		{"PUSH ECX", 32},
		{"ADD RAX, 0x77", 64},
		{"RET", 64},
		{"MOV EAX, EBX", 32},
		{"SUB RCX, RDX", 64},
	}
	for _, tc := range tests {
		t.Run(tc.line, func(t *testing.T) {
			bytes, err := a.Assemble(tc.line, tc.mode)
			if err != nil {
				t.Fatalf("assemble %q: %v", tc.line, err)
			}
			disasm, err := a.Disassemble(bytes, tc.mode)
			if err != nil {
				t.Fatalf("disassemble: %v", err)
			}
			t.Logf("%s -> %X -> %s", tc.line, bytes, disasm)
		})
	}
}

func TestAssembleRoundTrip(t *testing.T) {
	a := newAssembler(t)
	tests := []struct {
		line      string
		mode      int
		expectHex string
	}{
		{"NOP", 64, "90"},
		{"RET", 64, "C3"},
		{"PUSH RAX", 64, "50"},
		{"POP RAX", 64, "58"},
	}
	for _, tc := range tests {
		t.Run(tc.line, func(t *testing.T) {
			bytes, err := a.Assemble(tc.line, tc.mode)
			if err != nil {
				t.Fatalf("assemble %q: %v", tc.line, err)
			}
			gotHex := fmt.Sprintf("%X", bytes)
			if gotHex != tc.expectHex {
				t.Errorf("expected %s, got %s", tc.expectHex, gotHex)
			}
		})
	}
}
