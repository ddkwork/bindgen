package ipmrec

import (
	"testing"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
	"golang.org/x/sys/windows"
)

var a = &Ipmrec{}

func TestLoadDLL(t *testing.T) {
	t.Log("ARImpRec.dll embedded and ready")
}

func TestUnpackPdataSection(t *testing.T) {
	protected := byteslice.PtrFromString[int8]("nonexistent_protected.exe")
	dumped := byteslice.PtrFromString[int8]("nonexistent_dumped.exe")
	warning := make([]int8, 512)
	ret := a.UnpackPdataSection(protected, dumped, byteslice.PtrFromAnySlice[int8](warning))
	t.Logf("return=%d warning=%s", ret, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](warning)))
}

func TestGetNameFileOptimized(t *testing.T) {
	fileOrig := byteslice.PtrFromString[int8]("test_orig.exe")
	fileOpt := make([]int8, 512)
	ret := a.GetNameFileOptimized(fileOrig, byteslice.PtrFromAnySlice[int8](fileOpt))
	t.Logf("return=%d optimized=%s", ret, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](fileOpt)))
}

func TestRebuildSectionsFromArmadillo(t *testing.T) {
	protected := byteslice.PtrFromString[int8]("nonexistent_protected.exe")
	dumped := byteslice.PtrFromString[int8]("nonexistent_dumped.exe")
	warning := make([]int8, 512)
	ret := a.RebuildSectionsFromArmadillo(protected, dumped, byteslice.PtrFromAnySlice[int8](warning))
	t.Logf("return=%d warning=%s", ret, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](warning)))
}

func TestTryGetImportedFunction(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	var ordVal uint32
	var hintVal uint32
	ppOrd := &ordVal
	ppHint := &hintVal
	funcName := make([]int8, 256)
	module := make([]int8, 256)
	ret := a.TryGetImportedFunction(pid, 0, &ppOrd, &ppHint, byteslice.PtrFromAnySlice[int8](funcName), byteslice.PtrFromAnySlice[int8](module))
	t.Logf("return=%d ordinal=0x%08X hint=0x%08X funcName=%s module=%s",
		ret, ordVal, hintVal, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](funcName)), byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](module)))
}

func TestSearchAndRebuildImports(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	dumpedFile := byteslice.PtrFromString[int8]("nonexistent_dumped.exe")
	var iatRVAVal uint32
	var iatSizeVal uint32
	ppIATRVA := &iatRVAVal
	ppIATSize := &iatSizeVal
	warning := make([]int8, 512)
	ret := a.SearchAndRebuildImports(pid, dumpedFile, 0, 0, &ppIATRVA, &ppIATSize, byteslice.PtrFromAnySlice[int8](warning))
	t.Logf("return=%d IATRVA=0x%08X IATSize=0x%08X warning=%s",
		ret, iatRVAVal, iatSizeVal, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](warning)))
}

func TestSearchAndRebuildImportsIATOptimized(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	dumpedFile := byteslice.PtrFromString[int8]("nonexistent_dumped.exe")
	var iatRVAVal uint32
	var iatSizeVal uint32
	ppIATRVA := &iatRVAVal
	ppIATSize := &iatSizeVal
	warning := make([]int8, 512)
	ret := a.SearchAndRebuildImportsIATOptimized(pid, dumpedFile, 0, 0, &ppIATRVA, &ppIATSize, byteslice.PtrFromAnySlice[int8](warning))
	t.Logf("return=%d IATRVA=0x%08X IATSize=0x%08X warning=%s",
		ret, iatRVAVal, iatSizeVal, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](warning)))
}

func TestSearchAndRebuildImportsNoNewSection(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	dumpedFile := byteslice.PtrFromString[int8]("nonexistent_dumped.exe")
	var iatRVAVal uint32
	var iatSizeVal uint32
	ppIATRVA := &iatRVAVal
	ppIATSize := &iatSizeVal
	warning := make([]int8, 512)
	ret := a.SearchAndRebuildImportsNoNewSection(pid, dumpedFile, 0, 0, &ppIATRVA, &ppIATSize, byteslice.PtrFromAnySlice[int8](warning))
	t.Logf("return=%d IATRVA=0x%08X IATSize=0x%08X warning=%s",
		ret, iatRVAVal, iatSizeVal, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](warning)))
}

func TestGetProcNameAndOrdinal(t *testing.T) {
	kernel32, err := windows.LoadLibrary("kernel32.dll")
	if err != nil {
		t.Fatalf("LoadLibrary kernel32.dll: %v", err)
	}
	defer windows.FreeLibrary(kernel32)
	getProcAddr := windows.NewLazySystemDLL("kernel32.dll").NewProc("GetProcAddress")
	addr, _, _ := getProcAddr.Call(
		uintptr(kernel32),
		uintptr(unsafe.Pointer(byteslice.PtrFromString[int8]("GetCurrentProcessId"))),
	)
	if addr == 0 {
		t.Fatal("GetProcAddress returned NULL")
	}
	hMod := uintptr(kernel32)
	var ordVal uint32
	var hintVal uint32
	ppOrd := &ordVal
	ppHint := &hintVal
	procName := make([]int8, 256)
	ret := a.GetProcNameAndOrdinal(hMod, uint32(addr), &ppOrd, &ppHint, byteslice.PtrFromAnySlice[int8](procName))
	t.Logf("return=%d ordinal=0x%08X hint=0x%08X name=%s", ret, ordVal, hintVal, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](procName)))
}

func TestGetProcOrdinal(t *testing.T) {
	kernel32, err := windows.LoadLibrary("kernel32.dll")
	if err != nil {
		t.Fatalf("LoadLibrary kernel32.dll: %v", err)
	}
	defer windows.FreeLibrary(kernel32)
	getProcAddr := windows.NewLazySystemDLL("kernel32.dll").NewProc("GetProcAddress")
	addr, _, _ := getProcAddr.Call(
		uintptr(kernel32),
		uintptr(unsafe.Pointer(byteslice.PtrFromString[int8]("GetCurrentProcessId"))),
	)
	if addr == 0 {
		t.Fatal("GetProcAddress returned NULL")
	}
	hMod := uintptr(kernel32)
	ret := a.GetProcOrdinal(hMod, uint32(addr))
	t.Logf("return=0x%08X (format: 8000XXXX)", ret)
}

func TestGetProcName(t *testing.T) {
	kernel32, err := windows.LoadLibrary("kernel32.dll")
	if err != nil {
		t.Fatalf("LoadLibrary kernel32.dll: %v", err)
	}
	defer windows.FreeLibrary(kernel32)
	getProcAddr := windows.NewLazySystemDLL("kernel32.dll").NewProc("GetProcAddress")
	addr, _, _ := getProcAddr.Call(
		uintptr(kernel32),
		uintptr(unsafe.Pointer(byteslice.PtrFromString[int8]("GetCurrentProcessId"))),
	)
	if addr == 0 {
		t.Fatal("GetProcAddress returned NULL")
	}
	hMod := uintptr(kernel32)
	var hintVal uint32
	ppHint := &hintVal
	procName := make([]int8, 256)
	ret := a.GetProcName(hMod, uint32(addr), &ppHint, byteslice.PtrFromAnySlice[int8](procName))
	t.Logf("return=%d hint=0x%08X name=%s", ret, hintVal, byteslice.PtrToString(byteslice.PtrFromAnySlice[int8](procName)))
}

func TestGetAllVAddressesOfImports(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	var vAddrVal uint32
	ppVAddr := &vAddrVal
	ret := a.GetAllVAddressesOfImports(pid, 0, &ppVAddr, int32(1024))
	t.Logf("return=%d (num imports found)", ret)
}
