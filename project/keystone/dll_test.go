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

func TestKs_version(t *testing.T) {
	k := newKeystone(t)
	var major, minor uint32
	result := k.Ks_version(&major, &minor)
	expected := KsMakeVersion(KsApiMajor, KsApiMinor)
	if result != expected {
		t.Fatalf("ks_version returned %d (major=%d minor=%d), expected %d", result, major, minor, expected)
	}
	t.Logf("keystone version: %d.%d", major, minor)
}

func TestKs_arch_supported(t *testing.T) {
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
			got := k.Ks_arch_supported(tt.arch)
			if got != tt.want {
				t.Errorf("Ks_arch_supported(%v) = %v, want %v (%s)", tt.arch, got, tt.want, tt.reason)
			}
		})
	}
}

func TestKs_openClose_X86_32(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	err := k.Ks_open(KsArchX86, int32(KsMode32), &ks)
	if err != KsErrOk {
		t.Fatalf("Ks_open(X86, Mode32) failed: %d (%s)", err, goString(k.Ks_strerror(err)))
	}
	defer func() {
		err := k.Ks_close(ks)
		if err != KsErrOk {
			t.Errorf("Ks_close failed: %d (%s)", err, goString(k.Ks_strerror(err)))
		}
	}()
	if ks == nil {
		t.Fatal("Ks_open returned nil engine")
	}
}

func TestKs_openClose_X86_64(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	err := k.Ks_open(KsArchX86, int32(KsMode64), &ks)
	if err != KsErrOk {
		t.Fatalf("Ks_open(X86, Mode64) failed: %d (%s)", err, goString(k.Ks_strerror(err)))
	}
	defer func() {
		if closeErr := k.Ks_close(ks); closeErr != KsErrOk {
			t.Errorf("Ks_close failed: %d (%s)", closeErr, goString(k.Ks_strerror(closeErr)))
		}
	}()
	if ks == nil {
		t.Fatal("Ks_open returned nil engine")
	}
}

func TestKs_openInvalidMode(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	err := k.Ks_open(KsArchX86, int32(KsModeBigEndian), &ks)
	if err == KsErrOk {
		k.Ks_close(ks)
		t.Fatal("expected error for X86 + BigEndian mode, got success")
	}
}

func TestKs_asm_X86_32(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	if err := k.Ks_open(KsArchX86, int32(KsMode32), &ks); err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.Ks_close(ks)

	var encoding *uint8
	var encSize uintptr
	var statCount uintptr
	str, pin := cString("inc eax\ncall 0x1000\nmov ecx, edx")
	result := k.Ks_asm(ks, str, 0x0, &encoding, &encSize, &statCount)
	_ = pin
	t.Logf("Ks_asm result=%d errno=%d", result, k.Ks_errno(ks))
	if result != 0 {
		t.Fatalf("Ks_asm failed: %d (%s)", result, goString(k.Ks_strerror(Ks_err(result))))
	}
	if encoding == nil {
		t.Fatal("encoding is nil")
	}
	if encSize == 0 {
		t.Fatal("encSize is 0")
	}
	t.Logf("assembled %d bytes, %d statements", encSize, statCount)
	k.Ks_free(encoding)
}

func TestKs_asm_X86_64(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	if err := k.Ks_open(KsArchX86, int32(KsMode64), &ks); err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.Ks_close(ks)

	var encoding *uint8
	var encSize uintptr
	var statCount uintptr
	str, pin := cString("mov rax, 0x123456789ABCDEF0\nret")
	result := k.Ks_asm(ks, str, 0x400000, &encoding, &encSize, &statCount)
	_ = pin
	if result != 0 {
		t.Fatalf("Ks_asm failed: %d (%s)", result, goString(k.Ks_strerror(Ks_err(result))))
	}
	if encSize == 0 {
		t.Fatal("encSize is 0")
	}
	t.Logf("assembled %d bytes for 64-bit mode", encSize)
	k.Ks_free(encoding)
}

func TestKs_asm_Arm64(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	if err := k.Ks_open(KsArchArm64, int32(KsModeLittleEndian), &ks); err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.Ks_close(ks)

	var encoding *uint8
	var encSize uintptr
	var statCount uintptr
	str, pin := cString("mov x0, #42\nret")
	result := k.Ks_asm(ks, str, 0, &encoding, &encSize, &statCount)
	_ = pin
	if result != 0 {
		t.Fatalf("Ks_asm failed: %d (%s)", result, goString(k.Ks_strerror(Ks_err(result))))
	}
	t.Logf("arm64 assembled %d bytes", encSize)
	k.Ks_free(encoding)
}

func TestKs_strerror(t *testing.T) {
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
			msg := k.Ks_strerror(tt.code)
			if msg == nil {
				t.Fatal("Ks_strerror returned nil")
			}
			s := goString(msg)
			if len(s) > 0 && tt.code != KsErrOk {
				t.Logf("strerror(%d) = %q", tt.code, s)
			}
		})
	}
}

func TestKs_errno(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	err := k.Ks_open(KsArchX86, int32(KsMode32), &ks)
	if err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.Ks_close(ks)

	errno := k.Ks_errno(ks)
	if errno != KsErrOk {
		t.Logf("errno after open: %d (%s)", errno, goString(k.Ks_strerror(errno)))
	}
}

func TestKs_option_Syntax(t *testing.T) {
	k := newKeystone(t)
	var ks *Ks_engine
	if err := k.Ks_open(KsArchX86, int32(KsMode32), &ks); err != KsErrOk {
		t.Fatalf("open: %d", err)
	}
	defer k.Ks_close(ks)

	syntaxValue := uintptr(1)
	err := k.Ks_option(ks, KsOptSyntax, syntaxValue)
	if err != KsErrOk {
		t.Logf("option may not be supported: %d (%s)", err, goString(k.Ks_strerror(err)))
	}
}

func TestKs_free_NilPointer(t *testing.T) {
	k := newKeystone(t)
	k.Ks_free(nil)
	t.Log("Ks_free(nil) completed without error")
}

func TestFullLifecycle_X86(t *testing.T) {
	k := newKeystone(t)

	var ks *Ks_engine
	if err := k.Ks_open(KsArchX86, int32(KsMode32), &ks); err != KsErrOk {
		t.Fatalf("open: %d (%s)", err, goString(k.Ks_strerror(err)))
	}

	if errno := k.Ks_errno(ks); errno != KsErrOk {
		t.Errorf("unexpected errno after open: %d", errno)
	}

	var encoding *uint8
	var encSize, statCount uintptr
	str, pin := cString("nop\npush ebp\npop ebp")
	_ = pin
	if asmResult := k.Ks_asm(ks, str, 0, &encoding, &encSize, &statCount); asmResult != 0 {
		t.Fatalf("asm: %d (%s)", asmResult, goString(k.Ks_strerror(Ks_err(asmResult))))
	}
	if encSize == 0 {
		t.Fatal("expected non-zero encoding size")
	}
	k.Ks_free(encoding)

	if closeErr := k.Ks_close(ks); closeErr != KsErrOk {
		t.Errorf("close: %d (%s)", closeErr, goString(k.Ks_strerror(closeErr)))
	}
}
