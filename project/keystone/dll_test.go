package keystone

import (
	"testing"
	"unsafe"
)

func newKeystone(t *testing.T) *Keystone {
	t.Helper()
	return &Keystone{}
}

func goString(p *int8) string {
	if p == nil {
		return "<nil>"
	}
	var b []byte
	for i := 0; ; i++ {
		c := unsafe.Slice((*byte)(unsafe.Pointer(p)), i+1)[i]
		if c == 0 {
			break
		}
		b = append(b, c)
	}
	return string(b)
}

func cString(s string) (*int8, []byte) {
	b := append([]byte(s), 0)
	return (*int8)(unsafe.Pointer(&b[0])), b
}

func TestKsVersion(t *testing.T) {
	k := newKeystone(t)
	var major, minor uint32
	result := k.KsVersion(&major, &minor)
	expected := KsMakeVersion(KsApiMajor, KsApiMinor)
	if result != expected {
		t.Fatalf("ks_version returned %d (major=%d minor=%d), expected %d", result, major, minor, expected)
	}
	t.Logf("keystone version: %d.%d", major, minor)
}

func TestKsArchSupported(t *testing.T) {
	k := newKeystone(t)

	tests := []struct {
		arch   Ks_arch
		want   bool
		reason string
	}{
		{KsArchX86, true, "x86 should be supported"},
		{KsArchArm64, true, "arm64 should be supported"},
		{KsArchArm, true, "arm should be supported"},
		{KsArchMips, true, "mips should be supported"},
		{KsArchMax, false, "KsArchMax is sentinel, not a real arch"},
	}

	for _, tt := range tests {
		t.Run(tt.arch.String(), func(t *testing.T) {
			got := k.KsArchSupported(tt.arch)
			if got != tt.want {
				t.Errorf("KsArchSupported(%v) = %v, want %v (%s)", tt.arch, got, tt.want, tt.reason)
			}
		})
	}
}

func TestKsOpenClose_X86_32(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	err := k.KsOpen(KsArchX86, int32(KsMode32), &ks)
	if err != KsErrOk {
		t.Fatalf("KsOpen(X86, Mode32) failed: %d (%s)", err, goString(k.KsStrerror(err)))
	}
	defer func() {
		err := k.KsClose(ks)
		if err != KsErrOk {
			t.Errorf("KsClose failed: %d (%s)", err, goString(k.KsStrerror(err)))
		}
	}()
	if ks == nil {
		t.Fatal("KsOpen returned nil engine")
	}
}

func TestKsOpenClose_X86_64(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	err := k.KsOpen(KsArchX86, int32(KsMode64), &ks)
	if err != KsErrOk {
		t.Fatalf("KsOpen(X86, Mode64) failed: %d (%s)", err, goString(k.KsStrerror(err)))
	}
	defer func() {
		if closeErr := k.KsClose(ks); closeErr != KsErrOk {
			t.Errorf("KsClose failed: %d (%s)", closeErr, goString(k.KsStrerror(closeErr)))
		}
	}()
	if ks == nil {
		t.Fatal("KsOpen returned nil engine")
	}
}

func TestKsOpenInvalidMode(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	err := k.KsOpen(KsArchX86, int32(KsModeBigEndian), &ks)
	if err == KsErrOk {
		k.KsClose(ks)
		t.Fatal("expected error for X86 + BigEndian mode, got success")
	}
}

func TestKsAsm_X86_32(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	if err := k.KsOpen(KsArchX86, int32(KsMode32), &ks); err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.KsClose(ks)

	var encoding *uint8
	var encSize uintptr
	var statCount uintptr
	str, pin := cString("inc eax\ncall 0x1000\nmov ecx, edx")
	result := k.KsAsm(ks, str, 0x0, &encoding, &encSize, &statCount)
	_ = pin
	t.Logf("KsAsm result=%d errno=%d", result, k.KsErrno(ks))
	if result != 0 {
		t.Fatalf("KsAsm failed: %d (%s)", result, goString(k.KsStrerror(Ks_err(result))))
	}
	if encoding == nil {
		t.Fatal("encoding is nil")
	}
	if encSize == 0 {
		t.Fatal("encSize is 0")
	}
	t.Logf("assembled %d bytes, %d statements", encSize, statCount)
	k.KsFree(encoding)
}

func TestKsAsm_X86_64(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	if err := k.KsOpen(KsArchX86, int32(KsMode64), &ks); err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.KsClose(ks)

	var encoding *uint8
	var encSize uintptr
	var statCount uintptr
	str, pin := cString("mov rax, 0x123456789ABCDEF0\nret")
	result := k.KsAsm(ks, str, 0x400000, &encoding, &encSize, &statCount)
	_ = pin
	if result != 0 {
		t.Fatalf("KsAsm failed: %d (%s)", result, goString(k.KsStrerror(Ks_err(result))))
	}
	if encSize == 0 {
		t.Fatal("encSize is 0")
	}
	t.Logf("assembled %d bytes for 64-bit mode", encSize)
	k.KsFree(encoding)
}

func TestKsAsm_Arm64(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	if err := k.KsOpen(KsArchArm64, int32(KsModeLittleEndian), &ks); err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.KsClose(ks)

	var encoding *uint8
	var encSize uintptr
	var statCount uintptr
	str, pin := cString("mov x0, #42\nret")
	result := k.KsAsm(ks, str, 0, &encoding, &encSize, &statCount)
	_ = pin
	if result != 0 {
		t.Fatalf("KsAsm failed: %d (%s)", result, goString(k.KsStrerror(Ks_err(result))))
	}
	t.Logf("arm64 assembled %d bytes", encSize)
	k.KsFree(encoding)
}

func TestKsStrerror(t *testing.T) {
	k := newKeystone(t)

	tests := []struct {
		code Ks_err
	}{
		{KsErrOk},
		{KsErrNomem},
		{KsErrArch},
		{KsErrHandle},
		{KsErrMode},
	}

	for _, tt := range tests {
		t.Run(tt.code.String(), func(t *testing.T) {
			msg := k.KsStrerror(tt.code)
			if msg == nil {
				t.Fatal("KsStrerror returned nil")
			}
			s := goString(msg)
			if len(s) > 0 && tt.code != KsErrOk {
				t.Logf("strerror(%d) = %q", tt.code, s)
			}
		})
	}
}

func TestKsErrno(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	err := k.KsOpen(KsArchX86, int32(KsMode32), &ks)
	if err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.KsClose(ks)

	errno := k.KsErrno(ks)
	if errno != KsErrOk {
		t.Logf("errno after open: %d (%s)", errno, goString(k.KsStrerror(errno)))
	}
}

func TestKsOption_Syntax(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	if err := k.KsOpen(KsArchX86, int32(KsMode32), &ks); err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.KsClose(ks)

	syntaxValue := uintptr(1)
	err := k.KsOption(ks, KsOptSyntax, syntaxValue)
	if err != KsErrOk {
		t.Logf("option may not be supported: %d (%s)", err, goString(k.KsStrerror(err)))
	}
}

func TestKsFree_NilPointer(t *testing.T) {
	k := newKeystone(t)
	k.KsFree(nil)
	t.Log("KsFree(nil) completed without error")
}

func TestFullLifecycle_X86(t *testing.T) {
	k := newKeystone(t)

	var ks *Ks_engine
	if err := k.KsOpen(KsArchX86, int32(KsMode32), &ks); err != KsErrOk {
		t.Fatalf("open: %d (%s)", err, goString(k.KsStrerror(err)))
	}

	if errno := k.KsErrno(ks); errno != KsErrOk {
		t.Errorf("unexpected errno after open: %d", errno)
	}

	var encoding *uint8
	var encSize, statCount uintptr
	str, pin := cString("nop\npush ebp\npop ebp")
	_ = pin
	if asmResult := k.KsAsm(ks, str, 0, &encoding, &encSize, &statCount); asmResult != 0 {
		t.Fatalf("asm: %d (%s)", asmResult, goString(k.KsStrerror(Ks_err(asmResult))))
	}
	if encSize == 0 {
		t.Fatal("expected non-zero encoding size")
	}
	k.KsFree(encoding)

	if closeErr := k.KsClose(ks); closeErr != KsErrOk {
		t.Errorf("close: %d (%s)", closeErr, goString(k.KsStrerror(closeErr)))
	}
}
