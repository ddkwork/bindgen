package zydis

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
)

func TestZydisVersion(t *testing.T) {
	z := &Zydis{}
	ver := z.GetVersion()
	major := uint16((ver & 0xFFFF000000000000) >> 48)
	minor := uint16((ver & 0x0000FFFF00000000) >> 32)
	patch := uint16((ver & 0x00000000FFFF0000) >> 16)
	t.Logf("Zydis version: %d.%d.%d (0x%016X)", major, minor, patch, ver)
	if ver == 0 {
		t.Fatal("GetVersion returned 0")
	}
}

func TestIsFeatureEnabled(t *testing.T) {
	z := &Zydis{}
	for _, feat := range []struct {
		name     string
		value    ZydisFeature
		expected bool
	}{
		{"Decoder", ZydisFeatureDecoder, true},
		{"Encoder", ZydisFeatureEncoder, true},
		{"Formatter", ZydisFeatureFormatter, true},
		{"AVX512", ZydisFeatureAvx512, true},
		{"Knc", ZydisFeatureKnc, true},
		{"Segment", ZydisFeatureSegment, true},
	} {
		status := z.IsFeatureEnabled(feat.value)
		enabled := status == ZyanStatusTrue
		t.Logf("Feature %s: status=0x%X enabled=%v", feat.name, status, enabled)
		if enabled != feat.expected {
			t.Errorf("Feature %s: expected enabled=%v, got %v (status=0x%X)", feat.name, feat.expected, enabled, status)
		}
	}
}

func TestMnemonicGetStringWrapped(t *testing.T) {
	z := &Zydis{}
	str := z.MnemonicGetStringWrapped(ZydisMnemonicNop)
	if str == nil || str.Data == nil {
		t.Fatal("MnemonicGetStringWrapped returned nil")
	}
	result := unsafe.String((*byte)(unsafe.Pointer(str.Data)), int(str.Size))
	if result != "nop" {
		t.Fatalf("expected 'nop', got '%s'", result)
	}
	t.Logf("MnemonicGetStringWrapped(NOP): %s", result)
}

func TestMnemonicGetStringWrapped_MOV(t *testing.T) {
	z := &Zydis{}
	str := z.MnemonicGetStringWrapped(ZydisMnemonicMov)
	if str == nil || str.Data == nil {
		t.Fatal("MnemonicGetStringWrapped returned nil")
	}
	result := unsafe.String((*byte)(unsafe.Pointer(str.Data)), int(str.Size))
	if !strings.HasPrefix(result, "mov") {
		t.Fatalf("expected 'mov*', got '%s'", result)
	}
	t.Logf("MnemonicGetStringWrapped(MOV): %s", result)
}

func TestRegisterGetStringWrapped_RAX(t *testing.T) {
	z := &Zydis{}
	str := z.RegisterGetStringWrapped(ZydisRegisterRax)
	if str == nil || str.Data == nil {
		t.Fatal("RegisterGetStringWrapped returned nil")
	}
	result := unsafe.String((*byte)(unsafe.Pointer(str.Data)), int(str.Size))
	if result != "rax" {
		t.Fatalf("expected 'rax', got '%s'", result)
	}
	t.Logf("RegisterGetStringWrapped(RAX): %s", result)
}

func TestRegisterGetStringWrapped_RSP(t *testing.T) {
	z := &Zydis{}
	str := z.RegisterGetStringWrapped(ZydisRegisterRsp)
	if str == nil || str.Data == nil {
		t.Fatal("RegisterGetStringWrapped returned nil")
	}
	result := unsafe.String((*byte)(unsafe.Pointer(str.Data)), int(str.Size))
	if result != "rsp" {
		t.Fatalf("expected 'rsp', got '%s'", result)
	}
	t.Logf("RegisterGetStringWrapped(RSP): %s", result)
}

func TestDecoderInit_Long64(t *testing.T) {
	z := &Zydis{}
	var decoder ZydisDecoder
	status := z.DecoderInit(&decoder, ZydisMachineModeLong64, ZydisStackWidth64)
	if status != ZyanStatusSuccess {
		t.Fatalf("DecoderInit(Long64,64) failed: 0x%X", status)
	}
	t.Log("DecoderInit(Long64,64) succeeded")
}

func TestDecoderInit_Legacy32(t *testing.T) {
	z := &Zydis{}
	var decoder ZydisDecoder
	status := z.DecoderInit(&decoder, ZydisMachineModeLegacy32, ZydisStackWidth32)
	if status != ZyanStatusSuccess {
		t.Fatalf("DecoderInit(Legacy32,32) failed: 0x%X", status)
	}
	t.Log("DecoderInit(Legacy32,32) succeeded")
}

func disasmText(disasm *ZydisDisassembledInstruction) string {
	n := 0
	for n < len(disasm.Text) && disasm.Text[n] != 0 {
		n++
	}
	return unsafe.String((*byte)(unsafe.Pointer(&disasm.Text[0])), n)
}

func formatHexBytes(b []byte) string {
	var sb strings.Builder
	for i, v := range b {
		if i > 0 {
			sb.WriteByte(' ')
		}
		fmt.Fprintf(&sb, "%02X", v)
	}
	return sb.String()
}

func TestDisassembleIntel_NOP(t *testing.T) {
	z := &Zydis{}
	code := []byte{0x90}
	var disasm ZydisDisassembledInstruction
	status := z.DisassembleIntel(ZydisMachineModeLong64, 0, unsafe.Pointer(&code[0]), uintptr(len(code)), &disasm)
	if status != ZyanStatusSuccess {
		t.Fatalf("DisassembleIntel(NOP) failed: 0x%X", status)
	}
	text := disasmText(&disasm)
	t.Logf("DisassembleIntel(NOP): status=OK text='%s' len=%d", text, len(text))
}

func TestDisassembleIntel_RET(t *testing.T) {
	z := &Zydis{}
	code := []byte{0xC3}
	var disasm ZydisDisassembledInstruction
	status := z.DisassembleIntel(ZydisMachineModeLong64, 0, unsafe.Pointer(&code[0]), uintptr(len(code)), &disasm)
	if status != ZyanStatusSuccess {
		t.Fatalf("DisassembleIntel(RET) failed: 0x%X", status)
	}
	text := disasmText(&disasm)
	t.Logf("DisassembleIntel(RET): status=OK text='%s' len=%d", text, len(text))
}

func TestDisassembleIntel_MOV(t *testing.T) {
	z := &Zydis{}
	code := []byte{0x48, 0x89, 0xD8}
	var disasm ZydisDisassembledInstruction
	status := z.DisassembleIntel(ZydisMachineModeLong64, 0x401000, unsafe.Pointer(&code[0]), uintptr(len(code)), &disasm)
	if status != ZyanStatusSuccess {
		t.Fatalf("DisassembleIntel(MOV RAX,RBX) failed: 0x%X", status)
	}
	text := disasmText(&disasm)
	t.Logf("DisassembleIntel(MOV RAX,RBX @0x401000): %s", text)
}

func TestDisassembleATT_NOP(t *testing.T) {
	z := &Zydis{}
	code := []byte{0x90}
	var disasm ZydisDisassembledInstruction
	status := z.DisassembleATT(ZydisMachineModeLong64, 0, unsafe.Pointer(&code[0]), uintptr(len(code)), &disasm)
	if status != ZyanStatusSuccess {
		t.Fatalf("DisassembleATT(NOP) failed: 0x%X", status)
	}
	text := disasmText(&disasm)
	t.Logf("DisassembleATT(NOP): %s", text)
}

func TestDecoderDecodeFull_NOP(t *testing.T) {
	z := &Zydis{}
	var decoder ZydisDecoder
	z.DecoderInit(&decoder, ZydisMachineModeLong64, ZydisStackWidth64)

	code := []byte{0x90}
	var instruction ZydisDecodedInstruction
	var operands [ZydisMaxOperandCount]ZydisDecodedOperand
	status := z.DecoderDecodeFull(&decoder, unsafe.Pointer(&code[0]), uintptr(len(code)), &instruction, &operands[0])
	if status != ZyanStatusSuccess {
		t.Fatalf("DecoderDecodeFull(NOP) failed: 0x%X", status)
	}
	if instruction.Mnemonic != ZydisMnemonicNop {
		t.Fatalf("expected NOP, got Mnemonic=%v", instruction.Mnemonic)
	}
	if instruction.Length != 1 {
		t.Fatalf("expected length 1, got %d", instruction.Length)
	}
	t.Logf("decoded NOP: mnemonic=%v len=%d operand_count=%d", instruction.Mnemonic, instruction.Length, instruction.OperandCount)
}

func TestDecoderDecodeFull_XOR_EAX_EAX(t *testing.T) {
	z := &Zydis{}
	var decoder ZydisDecoder
	z.DecoderInit(&decoder, ZydisMachineModeLong64, ZydisStackWidth64)

	code := []byte{0x31, 0xC0}
	var instruction ZydisDecodedInstruction
	var operands [ZydisMaxOperandCount]ZydisDecodedOperand
	status := z.DecoderDecodeFull(&decoder, unsafe.Pointer(&code[0]), uintptr(len(code)), &instruction, &operands[0])
	if status != ZyanStatusSuccess {
		t.Fatalf("DecoderDecodeFull(XOR EAX,EAX) failed: 0x%X", status)
	}
	if instruction.Mnemonic != ZydisMnemonicXor {
		t.Fatalf("expected XOR, got Mnemonic=%v", instruction.Mnemonic)
	}
	t.Logf("decoded XOR EAX,EAX: mnemonic=%v len=%d operand_count=%d", instruction.Mnemonic, instruction.Length, instruction.OperandCount)
}

func TestDecoderDecodeFull_Invalid(t *testing.T) {
	z := &Zydis{}
	var decoder ZydisDecoder
	z.DecoderInit(&decoder, ZydisMachineModeLong64, ZydisStackWidth64)

	code := []byte{0xFF}
	var instruction ZydisDecodedInstruction
	var operands [ZydisMaxOperandCount]ZydisDecodedOperand
	status := z.DecoderDecodeFull(&decoder, unsafe.Pointer(&code[0]), uintptr(len(code)), &instruction, &operands[0])
	if status == ZyanStatusSuccess {
		t.Fatal("expected failure for invalid instruction 0xFF, got success")
	}
	t.Logf("correctly rejected invalid instruction: status=0x%X", status)
}

func TestEncoderNopFill(t *testing.T) {
	z := &Zydis{}
	buf := make([]uintptr, 16)
	status := z.EncoderNopFill(unsafe.Pointer(&buf[0]), 15)
	if status != ZyanStatusSuccess {
		t.Fatalf("EncoderNopFill failed: 0x%X", status)
	}
	t.Log("EncoderNopFill(15 bytes) succeeded")
}

func TestXxx(t *testing.T) {
	z := &Zydis{}
	data := []byte{
		0x51, 0x8D, 0x45, 0xFF, 0x50, 0xFF, 0x75, 0x0C, 0xFF, 0x75,
		0x08, 0xFF, 0x15, 0xA0, 0xA5, 0x48, 0x76, 0x85, 0xC0, 0x0F,
		0x88, 0xFC, 0xDA, 0x02, 0x00,
	}
	runtimeAddress := uint64(0x007FFFFFFF400000)
	offset := 0
	for offset < len(data) {
		var disasm ZydisDisassembledInstruction
		status := z.DisassembleIntel(ZydisMachineModeLong64, runtimeAddress,
			unsafe.Pointer(&data[offset]),
			uintptr(len(data)-offset), &disasm)
		if status != ZyanStatusSuccess {
			t.Logf("Decode failed at offset %d: 0x%X", offset, status)
			break
		}
		text := disasmText(&disasm)
		t.Logf("%016X  %s", runtimeAddress, text)
		offset += int(disasm.Info.Length)
		runtimeAddress += uint64(disasm.Info.Length)
	}
}

func TestExample_DisassembleSimple(t *testing.T) {
	z := &Zydis{}
	t.Logf("=== Layout Diagnostics ===")
	var di ZydisDisassembledInstruction
	t.Logf("DisassembledInstruction: sizeof=%d Text@%d", unsafe.Sizeof(di), unsafe.Offsetof(di.Text))
	t.Logf("  Operand[0] sizeof=%d", unsafe.Sizeof(di.Operands[0]))
	var op ZydisDecodedOperand_
	t.Logf("DecodedOperand_: sizeof=%d", unsafe.Sizeof(op))
	t.Logf("  Id@%d Vis@%d Act@%d Enc@%d Sz@%d ET@%d ES@%d EC@%d Attr@%d Typ@%d",
		unsafe.Offsetof(op.Id), unsafe.Offsetof(op.Visibility), unsafe.Offsetof(op.Actions),
		unsafe.Offsetof(op.Encoding), unsafe.Offsetof(op.Size), unsafe.Offsetof(op.ElementType),
		unsafe.Offsetof(op.ElementSize), unsafe.Offsetof(op.ElementCount), unsafe.Offsetof(op.Attributes),
		unsafe.Offsetof(op.Type))
	t.Logf("  Reg sizeof=%d Mem sizeof=%d Imm sizeof=%d Ptr sizeof=%d",
		unsafe.Sizeof(ZydisDecodedOperandReg{}),
		unsafe.Sizeof(ZydisDecodedOperandMem{}),
		unsafe.Sizeof(ZydisDecodedOperandImm{}),
		unsafe.Sizeof(ZydisDecodedOperandPtr{}))
	var mem ZydisDecodedOperandMem_
	t.Logf("  Mem: Type@%d Seg@%d Base@%d Idx@%d Scale@%d Disp@%d sizeof=%d",
		unsafe.Offsetof(mem.Type), unsafe.Offsetof(mem.Segment), unsafe.Offsetof(mem.Base),
		unsafe.Offsetof(mem.Index), unsafe.Offsetof(mem.Scale), unsafe.Offsetof(mem.Disp),
		unsafe.Sizeof(mem))
	var imm ZydisDecodedOperandImm_
	t.Logf("  Imm: Signed@%d Rel@%d Val@%d sizeof=%d",
		unsafe.Offsetof(imm.IsSigned), unsafe.Offsetof(imm.IsRelative), unsafe.Offsetof(imm.Value),
		unsafe.Sizeof(imm))
	var insnDiag ZydisDecodedInstruction_
	t.Logf("DecodedInstruction_: sizeof=%d", unsafe.Sizeof(insnDiag))
	t.Logf("  MM@%d Mn@%d Len@%d Enc@%d OM@%d Op@%d SW@%d OW@%d AW@%d OC@%d OCV@%d Attr@%d CPU@%d FPU@%d AVX@%d Meta@%d Raw@%d",
		unsafe.Offsetof(insnDiag.MachineMode), unsafe.Offsetof(insnDiag.Mnemonic),
		unsafe.Offsetof(insnDiag.Length), unsafe.Offsetof(insnDiag.Encoding),
		unsafe.Offsetof(insnDiag.OpcodeMap), unsafe.Offsetof(insnDiag.Opcode),
		unsafe.Offsetof(insnDiag.StackWidth), unsafe.Offsetof(insnDiag.OperandWidth),
		unsafe.Offsetof(insnDiag.AddressWidth), unsafe.Offsetof(insnDiag.OperandCount),
		unsafe.Offsetof(insnDiag.OperandCountVisible), unsafe.Offsetof(insnDiag.Attributes),
		unsafe.Offsetof(insnDiag.CpuFlags), unsafe.Offsetof(insnDiag.FpuFlags),
		unsafe.Offsetof(insnDiag.Avx), unsafe.Offsetof(insnDiag.Meta), unsafe.Offsetof(insnDiag.Raw))
	t.Logf("  AVX sizeof=%d Meta sizeof=%d Raw sizeof=%d",
		unsafe.Sizeof(ZydisDecodedInstructionAvx{}),
		unsafe.Sizeof(ZydisDecodedInstructionMeta{}),
		unsafe.Sizeof(ZydisDecodedInstructionRaw{}))
	var raw ZydisDecodedInstructionRaw_
	t.Logf("  Raw: PC@%d Pfx@%d Enc2@%d Modrm@%d Sib@%d Disp@%d Imm@%d sizeof=%d",
		unsafe.Offsetof(raw.PrefixCount), unsafe.Offsetof(raw.Prefixes),
		unsafe.Offsetof(raw.Encoding2),
		unsafe.Offsetof(raw.Modrm), unsafe.Offsetof(raw.Sib),
		unsafe.Offsetof(raw.Disp), unsafe.Offsetof(raw.Imm),
		unsafe.Sizeof(raw))
	t.Logf("  Raw.Pfx[0] sizeof=%d Disp sizeof=%d Imm[0] sizeof=%d",
		unsafe.Sizeof(raw.Prefixes[0]),
		unsafe.Sizeof(raw.Disp), unsafe.Sizeof(raw.Imm[0]))
	var disp ZydisDecodedInstructionRawDisp_
	t.Logf("    Disp: Val@%d Sz@%d Off@%d sizeof=%d",
		unsafe.Offsetof(disp.Value), unsafe.Offsetof(disp.Size), unsafe.Offsetof(disp.Offset),
		unsafe.Sizeof(disp))
	var imm0 ZydisDecodedInstructionRawImm_
	t.Logf("    Imm: Sig@%d Rel@%d Val@%d Sz@%d Off@%d sizeof=%d",
		unsafe.Offsetof(imm0.IsSigned), unsafe.Offsetof(imm0.IsRelative),
		unsafe.Offsetof(imm0.Value), unsafe.Offsetof(imm0.Size), unsafe.Offsetof(imm0.Offset),
		unsafe.Sizeof(imm0))
	data := []byte{
		0x51, 0x8D, 0x45, 0xFF, 0x50, 0xFF, 0x75, 0x0C, 0xFF, 0x75,
		0x08, 0xFF, 0x15, 0xA0, 0xA5, 0x48, 0x76, 0x85, 0xC0, 0x0F,
		0x88, 0xFC, 0xDA, 0x02, 0x00,
	}
	runtimeAddress := uint64(0x007FFFFFFF400000)
	offset := 0
	var insn ZydisDisassembledInstruction
	for offset < len(data) {
		status := z.DisassembleIntel(
			ZydisMachineModeLong64, runtimeAddress,
			unsafe.Pointer(&data[offset]),
			uintptr(len(data)-offset), &insn,
		)
		if status != ZyanStatusSuccess {
			break
		}
		t.Logf("%016X  %s", runtimeAddress, disasmText(&insn))
		offset += int(insn.Info.Length)
		runtimeAddress += uint64(insn.Info.Length)
	}
}

func TestEncoderRequestLayout(t *testing.T) {
	var req ZydisEncoderRequest_
	t.Logf("EncoderRequest_: sizeof=%d", unsafe.Sizeof(req))
	t.Logf("  MM@%d AE@%d Mn@%d Pfx@%d BT@%d BW@%d ASH@%d OSH@%d OC@%d Op@%d Evex@%d Mvex@%d",
		unsafe.Offsetof(req.MachineMode), unsafe.Offsetof(req.AllowedEncodings),
		unsafe.Offsetof(req.Mnemonic), unsafe.Offsetof(req.Prefixes),
		unsafe.Offsetof(req.BranchType), unsafe.Offsetof(req.BranchWidth),
		unsafe.Offsetof(req.AddressSizeHint), unsafe.Offsetof(req.OperandSizeHint),
		unsafe.Offsetof(req.OperandCount), unsafe.Offsetof(req.Operands),
		unsafe.Offsetof(req.Evex), unsafe.Offsetof(req.Mvex))
	t.Logf("  Operands[0] sizeof=%d Evex sizeof=%d Mvex sizeof=%d",
		unsafe.Sizeof(req.Operands[0]),
		unsafe.Sizeof(req.Evex), unsafe.Sizeof(req.Mvex))
	var op ZydisEncoderOperand
	t.Logf("  EncoderOperand: sizeof=%d", unsafe.Sizeof(op))
}

func TestExample_EncodeMov(t *testing.T) {
	z := &Zydis{}
	var req ZydisEncoderRequest_
	req.MachineMode = ZydisMachineModeLong64
	req.Mnemonic = ZydisMnemonicMov
	req.OperandCount = 2
	req.Operands[0].Type = ZydisOperandTypeRegister
	req.Operands[0].Reg.Value = ZydisRegisterRax
	req.Operands[1].Type = ZydisOperandTypeImmediate
	req.Operands[1].Imm.Data = 0x1337

	var encoded [ZydisMaxInstructionLength]byte
	encodedLen := uintptr(len(encoded))

	status := z.EncoderEncodeInstruction(&req, unsafe.Pointer(&encoded[0]), &encodedLen)
	if status != ZyanStatusSuccess {
		t.Fatalf("ZydisEncoderEncodeInstruction failed: 0x%08X", status)
	}

	hexStr := fmt.Sprintf("%02X", encoded[:encodedLen])
	t.Logf("Encoded 'mov rax, 0x1337': %s (len=%d)", hexStr, encodedLen)

	var disasm ZydisDisassembledInstruction
	status3 := z.DisassembleIntel(ZydisMachineModeLong64, 0, unsafe.Pointer(&encoded[0]), uintptr(encodedLen), &disasm)
	if status3 != ZyanStatusSuccess {
		t.Fatalf("DisassembleIntel failed: 0x%08X", status3)
	}
	text := disasmText(&disasm)
	t.Logf("Re-disassembled: %s", text)
	if !strings.Contains(text, "mov") {
		t.Errorf("expected 'mov' in disassembly, got: %s", text)
	}
}

func TestExample_EncodeFromScratch(t *testing.T) {
	z := &Zydis{}
	var buf [256]byte
	writePtr := 0
	remaining := len(buf)

	encode := func(req *ZydisEncoderRequest_) {
		var instrLen uintptr = uintptr(remaining)
		status := z.EncoderEncodeInstruction(req, unsafe.Pointer(&buf[writePtr]), &instrLen)
		if status != ZyanStatusSuccess {
			t.Fatalf("ZydisEncoderEncodeInstruction failed: 0x%08X", status)
		}
		writePtr += int(instrLen)
		remaining -= int(instrLen)
	}

	var req ZydisEncoderRequest_
	req.MachineMode = ZydisMachineModeLong64
	req.Mnemonic = ZydisMnemonicMov
	req.OperandCount = 2
	req.Operands[0].Type = ZydisOperandTypeRegister
	req.Operands[0].Reg.Value = ZydisRegisterRax
	req.Operands[1].Type = ZydisOperandTypeImmediate
	req.Operands[1].Imm.Data = 0x1337
	encode(&req)

	req = ZydisEncoderRequest_{}
	req.MachineMode = ZydisMachineModeLong64
	req.Mnemonic = ZydisMnemonicRet
	encode(&req)

	totalLen := writePtr
	hexParts := make([]string, totalLen)
	for i := 0; i < totalLen; i++ {
		hexParts[i] = fmt.Sprintf("%02X", buf[i])
	}
	t.Logf("Assembled bytes: %s", strings.Join(hexParts, " "))

	runtimeAddr := uint64(0x007FFFFFFF400000)
	offset := 0
	for offset < totalLen {
		var disasm ZydisDisassembledInstruction
		status := z.DisassembleIntel(ZydisMachineModeLong64, runtimeAddr, unsafe.Pointer(&buf[offset]), uintptr(totalLen-offset), &disasm)
		if status != ZyanStatusSuccess {
			break
		}
		text := disasmText(&disasm)
		t.Logf("%016X  %s", runtimeAddr, text)
		offset += int(disasm.Info.Length)
		runtimeAddr += uint64(disasm.Info.Length)
	}
}

func TestExample_RewriteCode(t *testing.T) {
	z := &Zydis{}

	addBytes := []byte{0x01, 0xC0}
	jzBytes := []byte{0x0F, 0x84, 0x00, 0x00, 0x00, 0x00}

	testCases := []struct {
		name    string
		bytes   []byte
		wantSub bool
		wantJnz bool
	}{
		{"add eax,eax", addBytes, true, false},
		{"jz", jzBytes, false, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var decoder ZydisDecoder_
			status := z.DecoderInit(&decoder, ZydisMachineModeLong64, ZydisStackWidth64)
			if status != ZyanStatusSuccess {
				t.Fatalf("DecoderInit failed: 0x%08X", status)
			}

			var instr ZydisDecodedInstruction
			var operands [10]ZydisDecodedOperand
			status = z.DecoderDecodeFull(&decoder, unsafe.Pointer(&tc.bytes[0]), uintptr(len(tc.bytes)), &instr, &operands[0])
			if status != ZyanStatusSuccess {
				t.Fatalf("DecoderDecodeFull failed: 0x%08X", status)
			}

			var fmt_ ZydisFormatter_
			z.FormatterInit(&fmt_, ZydisFormatterStyleIntel)

			var fmtBuf [256]int8
			z.FormatterFormatInstruction(&fmt_, &instr, &operands[0], instr.OperandCountVisible, &fmtBuf[0], uintptr(len(fmtBuf)), 0, nil)
			t.Logf("Original: %s", byteslice.ToString(fmtBuf[:]))

			var encReq ZydisEncoderRequest_
			status = z.EncoderDecodedInstructionToEncoderRequest(&instr, &operands[0], instr.OperandCountVisible, &encReq)
			if status != ZyanStatusSuccess {
				t.Fatalf("EncoderDecodedInstructionToEncoderRequest failed: 0x%08X", status)
			}

			if tc.wantSub && encReq.Mnemonic == ZydisMnemonicAdd {
				encReq.Mnemonic = ZydisMnemonicSub
			}
			if tc.wantJnz && encReq.Mnemonic == ZydisMnemonicJz {
				encReq.Mnemonic = ZydisMnemonicJnz
			}

			for i := uint8(0); i < encReq.OperandCount; i++ {
				op := &encReq.Operands[i]
				switch op.Type {
				case ZydisOperandTypeImmediate:
					op.Imm.Data = 0x42
				case ZydisOperandTypeMemory:
					op.Mem.Displacement = 0x1337
				}
			}

			var newBytes [ZydisMaxInstructionLength]byte
			var newLen uintptr = uintptr(len(newBytes))
			status = z.EncoderEncodeInstruction(&encReq, unsafe.Pointer(&newBytes[0]), &newLen)
			if status != ZyanStatusSuccess {
				t.Fatalf("EncoderEncodeInstruction failed: 0x%08X", status)
			}

			hexParts := make([]string, newLen)
			for i := uintptr(0); i < newLen; i++ {
				hexParts[i] = fmt.Sprintf("%02X", newBytes[i])
			}
			t.Logf("New bytes: %s", strings.Join(hexParts, " "))

			var newInstr ZydisDecodedInstruction
			var newOps [10]ZydisDecodedOperand
			status = z.DecoderDecodeFull(&decoder, unsafe.Pointer(&newBytes[0]), newLen, &newInstr, &newOps[0])
			if status != ZyanStatusSuccess {
				t.Fatalf("re-decode failed: 0x%08X", status)
			}

			var newFmtBuf [256]int8
			z.FormatterFormatInstruction(&fmt_, &newInstr, &newOps[0], newInstr.OperandCountVisible, &newFmtBuf[0], uintptr(len(newFmtBuf)), 0, nil)
			t.Logf("New:      %s", byteslice.ToString(newFmtBuf[:]))

			if tc.wantSub && newInstr.Mnemonic != ZydisMnemonicSub {
				t.Errorf("expected sub, got %v", newInstr.Mnemonic)
			}
			if tc.wantJnz && newInstr.Mnemonic != ZydisMnemonicJnz {
				t.Errorf("expected jnz, got %v", newInstr.Mnemonic)
			}
		})
	}
}

func TestExample_Formatter01_SymbolResolver(t *testing.T) {
	z := &Zydis{}

	data := []byte{
		0x48, 0x8B, 0x05, 0x39, 0x00, 0x13, 0x00,
		0x50,
		0xFF, 0x15, 0xF2, 0x10, 0x00, 0x00,
		0x85, 0xC0,
		0x0F, 0x84, 0x00, 0x00, 0x00, 0x00,
		0xE9, 0xE5, 0x0F, 0x00, 0x00,
	}

	var decoder ZydisDecoder_
	status := z.DecoderInit(&decoder, ZydisMachineModeLong64, ZydisStackWidth64)
	if status != ZyanStatusSuccess {
		t.Fatalf("DecoderInit failed: 0x%08X", status)
	}

	var fmt_ ZydisFormatter_
	z.FormatterInit(&fmt_, ZydisFormatterStyleIntel)
	z.FormatterSetProperty(&fmt_, ZydisFormatterPropForceSegment, 1)
	z.FormatterSetProperty(&fmt_, ZydisFormatterPropForceSize, 1)

	runtimeAddr := uint64(0x007FFFFFFF400000)
	offset := 0
	for offset < len(data) {
		var instr ZydisDecodedInstruction
		var operands [10]ZydisDecodedOperand
		status = z.DecoderDecodeFull(&decoder, unsafe.Pointer(&data[offset]), uintptr(len(data)-offset), &instr, &operands[0])
		if status != ZyanStatusSuccess {
			break
		}

		var fmtBuf [256]int8
		z.FormatterFormatInstruction(&fmt_, &instr, &operands[0], instr.OperandCountVisible, &fmtBuf[0], uintptr(len(fmtBuf)), runtimeAddr, nil)
		text := byteslice.ToString(fmtBuf[:])
		t.Logf("%016X  %s", runtimeAddr, text)

		offset += int(instr.Length)
		runtimeAddr += uint64(instr.Length)
	}
}

func TestExample_Formatter03_Tokenizer(t *testing.T) {
	z := &Zydis{}

	data := []byte{
		0x62, 0xF1, 0x6C, 0x5F, 0xC2, 0x54, 0x98, 0x40, 0x0F,
	}

	var decoder ZydisDecoder_
	status := z.DecoderInit(&decoder, ZydisMachineModeLong64, ZydisStackWidth64)
	if status != ZyanStatusSuccess {
		t.Fatalf("DecoderInit failed: 0x%08X", status)
	}

	var fmt_ ZydisFormatter_
	z.FormatterInit(&fmt_, ZydisFormatterStyleIntel)
	z.FormatterSetProperty(&fmt_, ZydisFormatterPropForceSegment, 1)
	z.FormatterSetProperty(&fmt_, ZydisFormatterPropForceSize, 1)

	var instr ZydisDecodedInstruction
	var operands [10]ZydisDecodedOperand
	status = z.DecoderDecodeFull(&decoder, unsafe.Pointer(&data[0]), uintptr(len(data)), &instr, &operands[0])
	if status != ZyanStatusSuccess {
		t.Fatalf("DecoderDecodeFull failed: 0x%08X", status)
	}

	var fmtBuf [256]byte
	var token *ZydisFormatterTokenConst
	status = z.FormatterTokenizeInstruction(&fmt_, &instr, &operands[0], instr.OperandCountVisible, unsafe.Pointer(&fmtBuf[0]), uintptr(len(fmtBuf)), 0, &token, nil)
	if status != ZyanStatusSuccess {
		t.Fatalf("FormatterTokenizeInstruction failed: 0x%08X", status)
	}

	tokenTypeNames := map[uint32]string{
		ZydisTokenInvalid: "INVALID", ZydisTokenWhitespace: "WHITESPACE",
		ZydisTokenDelimiter: "DELIMITER", ZydisTokenParenthesisOpen: "PARENTHESIS_OPEN",
		ZydisTokenParenthesisClose: "PARENTHESIS_CLOSE", ZydisTokenPrefix: "PREFIX",
		ZydisTokenMnemonic: "MNEMONIC", ZydisTokenRegister: "REGISTER",
		ZydisTokenAddressAbs: "ADDRESS_ABS", ZydisTokenAddressRel: "ADDRESS_REL",
		ZydisTokenDisplacement: "DISPLACEMENT", ZydisTokenImmediate: "IMMEDIATE",
		ZydisTokenTypecast: "TYPECAST", ZydisTokenDecorator: "DECORATOR",
		ZydisTokenSymbol: "SYMBOL",
	}

	tokenCount := 0
	for token != nil {
		var tType ZydisTokenType
		var tValue ZyanConstCharPointer
		status = z.FormatterTokenGetValue(token, &tType, &tValue)
		if status != ZyanStatusSuccess {
			break
		}

		typeName := "UNKNOWN"
		if name, ok := tokenTypeNames[uint32(tType)]; ok {
			typeName = name
		}

		var valStr string
		if tValue != nil {
			valStr = byteslice.ToString(unsafe.Slice((*int8)(unsafe.Pointer(tValue)), 64))
		}
		t.Logf("ZYDIS_TOKEN_%-17s (%02X): \"%s\"", typeName, tType, valStr)

		var nextToken *ZydisFormatterTokenConst
		nextStatus := z.FormatterTokenNext(&token)
		if nextStatus != ZyanStatusSuccess {
			break
		}
		_ = nextToken
		tokenCount++
		if tokenCount > 50 {
			t.Log("too many tokens, stopping")
			break
		}
	}

	if tokenCount == 0 {
		t.Error("no tokens produced")
	}
}
