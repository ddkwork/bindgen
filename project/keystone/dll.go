package keystone

import (
	_ "embed"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Keystone struct{}

//go:embed keystone.dll
var dllBytes []byte

var (
	dll            *windows.LazyDLL
	proc_ks_version      *windows.LazyProc
	proc_ks_arch_supported      *windows.LazyProc
	proc_ks_open      *windows.LazyProc
	proc_ks_close      *windows.LazyProc
	proc_ks_errno      *windows.LazyProc
	proc_ks_strerror      *windows.LazyProc
	proc_ks_option      *windows.LazyProc
	proc_ks_asm      *windows.LazyProc
	proc_ks_free      *windows.LazyProc
)

func init() {
	dll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, "keystone.dll"))
	proc_ks_version = dll.NewProc("ks_version")
	proc_ks_arch_supported = dll.NewProc("ks_arch_supported")
	proc_ks_open = dll.NewProc("ks_open")
	proc_ks_close = dll.NewProc("ks_close")
	proc_ks_errno = dll.NewProc("ks_errno")
	proc_ks_strerror = dll.NewProc("ks_strerror")
	proc_ks_option = dll.NewProc("ks_option")
	proc_ks_asm = dll.NewProc("ks_asm")
	proc_ks_free = dll.NewProc("ks_free")
}

func saveEmbeddedDLL(data []byte, name string) string {
	tmpDir := os.TempDir()
	p := filepath.Join(tmpDir, name)
	if _, err := os.Stat(p); err == nil {
		return p
	}
	os.WriteFile(p, data, 0644)
	return p
}

func (k *Keystone) Ks_version(Major *uint32, Minor *uint32) uint32 {
	r1, _, _ := proc_ks_version.Call(uintptr(unsafe.Pointer(Major)), uintptr(unsafe.Pointer(Minor)))
	return uint32(r1)
}

func (k *Keystone) Ks_arch_supported(Arch Ks_arch) bool {
	r1, _, _ := proc_ks_arch_supported.Call(uintptr(Arch))
	return r1 != 0
}

func (k *Keystone) Ks_open(Arch Ks_arch, Mode int32, Ks **Ks_engine) Ks_err {
	r1, _, _ := proc_ks_open.Call(uintptr(Arch), uintptr(Mode), uintptr(unsafe.Pointer(Ks)))
	return Ks_err(r1)
}

func (k *Keystone) Ks_close(Ks *Ks_engine) Ks_err {
	r1, _, _ := proc_ks_close.Call(uintptr(unsafe.Pointer(Ks)))
	return Ks_err(r1)
}

func (k *Keystone) Ks_errno(Ks *Ks_engine) Ks_err {
	r1, _, _ := proc_ks_errno.Call(uintptr(unsafe.Pointer(Ks)))
	return Ks_err(r1)
}

func (k *Keystone) Ks_strerror(Code Ks_err) *int8 {
	r1, _, _ := proc_ks_strerror.Call(uintptr(Code))
	return (*int8)(unsafe.Pointer(r1))
}

func (k *Keystone) Ks_option(Ks *Ks_engine, Type Ks_opt_type, Value uintptr) Ks_err {
	r1, _, _ := proc_ks_option.Call(uintptr(unsafe.Pointer(Ks)), uintptr(Type), Value)
	return Ks_err(r1)
}

func (k *Keystone) Ks_asm(Ks *Ks_engine, String *int8, Address uint64, Encoding **uint8, Encoding_size *uintptr, Stat_count *uintptr) int32 {
	r1, _, _ := proc_ks_asm.Call(uintptr(unsafe.Pointer(Ks)), uintptr(unsafe.Pointer(String)), uintptr(Address), uintptr(unsafe.Pointer(Encoding)), uintptr(unsafe.Pointer(Encoding_size)), uintptr(unsafe.Pointer(Stat_count)))
	return int32(r1)
}

func (k *Keystone) Ks_free(P *uint8) {
	proc_ks_free.Call(uintptr(unsafe.Pointer(P)))
}

