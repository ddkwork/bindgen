package ipmrec

import (
	"testing"
	"unsafe"

	"golang.org/x/sys/windows"
)

func bytePtr(s string) *PChar {
	p, _ := windows.BytePtrFromString(s)
	return (*PChar)(unsafe.Pointer(p))
}

func bufPtr(b []byte) *PChar {
	return (*PChar)(unsafe.Pointer(&b[0]))
}

func pcardinalBuf(n int) (*PCardinal, []Cardinal) {
	buf := make([]Cardinal, n)
	p := &buf[0]
	return &p, buf
}

var a = &Ipmrec{}

func TestLoadDLL(t *testing.T) {
	t.Log("ARImpRec.dll embedded and ready")
}

func TestUnpackPdataSection(t *testing.T) {
	protected := bytePtr("nonexistent_protected.exe")
	dumped := bytePtr("nonexistent_dumped.exe")
	warning := make([]byte, 512)
	ret := a.UnpackPdataSection(protected, dumped, bufPtr(warning))
	t.Logf("return=%d warning=%s", ret, goString(bufPtr(warning)))
}

func TestGetNameFileOptimized(t *testing.T) {
	fileOrig := bytePtr("test_orig.exe")
	fileOpt := make([]byte, 512)
	ret := a.GetNameFileOptimized(fileOrig, bufPtr(fileOpt))
	t.Logf("return=%d optimized=%s", ret, goString(bufPtr(fileOpt)))
}

func TestRebuildSectionsFromArmadillo(t *testing.T) {
	protected := bytePtr("nonexistent_protected.exe")
	dumped := bytePtr("nonexistent_dumped.exe")
	warning := make([]byte, 512)
	ret := a.RebuildSectionsFromArmadillo(protected, dumped, bufPtr(warning))
	t.Logf("return=%d warning=%s", ret, goString(bufPtr(warning)))
}

func TestTryGetImportedFunction(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	pOrd, ordBuf := pcardinalBuf(1)
	pHint, hintBuf := pcardinalBuf(1)
	funcName := make([]byte, 256)
	module := make([]byte, 256)
	ret := a.TryGetImportedFunction(pid, 0, pOrd, pHint, bufPtr(funcName), bufPtr(module))
	t.Logf("return=%d ordinal=0x%08X hint=0x%08X funcName=%s module=%s",
		ret, Cardinal(ordBuf[0]), Cardinal(hintBuf[0]), goString(bufPtr(funcName)), goString(bufPtr(module)))
}

func TestSearchAndRebuildImports(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	dumpedFile := bytePtr("nonexistent_dumped.exe")
	pIATRVA, iatRVABuf := pcardinalBuf(1)
	pIATSize, iatSizeBuf := pcardinalBuf(1)
	warning := make([]byte, 512)
	ret := a.SearchAndRebuildImports(pid, dumpedFile, 0, 0, pIATRVA, pIATSize, bufPtr(warning))
	t.Logf("return=%d IATRVA=0x%08X IATSize=0x%08X warning=%s",
		ret, Cardinal(iatRVABuf[0]), Cardinal(iatSizeBuf[0]), goString(bufPtr(warning)))
}

func TestSearchAndRebuildImportsIATOptimized(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	dumpedFile := bytePtr("nonexistent_dumped.exe")
	pIATRVA, iatRVABuf := pcardinalBuf(1)
	pIATSize, iatSizeBuf := pcardinalBuf(1)
	warning := make([]byte, 512)
	ret := a.SearchAndRebuildImportsIATOptimized(pid, dumpedFile, 0, 0, pIATRVA, pIATSize, bufPtr(warning))
	t.Logf("return=%d IATRVA=0x%08X IATSize=0x%08X warning=%s",
		ret, Cardinal(iatRVABuf[0]), Cardinal(iatSizeBuf[0]), goString(bufPtr(warning)))
}

func TestSearchAndRebuildImportsNoNewSection(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	dumpedFile := bytePtr("nonexistent_dumped.exe")
	pIATRVA, iatRVABuf := pcardinalBuf(1)
	pIATSize, iatSizeBuf := pcardinalBuf(1)
	warning := make([]byte, 512)
	ret := a.SearchAndRebuildImportsNoNewSection(pid, dumpedFile, 0, 0, pIATRVA, pIATSize, bufPtr(warning))
	t.Logf("return=%d IATRVA=0x%08X IATSize=0x%08X warning=%s",
		ret, Cardinal(iatRVABuf[0]), Cardinal(iatSizeBuf[0]), goString(bufPtr(warning)))
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
		uintptr(unsafe.Pointer(bytePtr("GetCurrentProcessId"))),
	)
	if addr == 0 {
		t.Fatal("GetProcAddress returned NULL")
	}
	hMod := HMODULE(unsafe.Pointer(kernel32))
	pOrd, ordBuf := pcardinalBuf(1)
	pHint, hintBuf := pcardinalBuf(1)
	procName := make([]byte, 256)
	ret := a.GetProcNameAndOrdinal(hMod, Cardinal(addr), pOrd, pHint, bufPtr(procName))
	t.Logf("return=%d ordinal=0x%08X hint=0x%08X name=%s", ret, Cardinal(ordBuf[0]), Cardinal(hintBuf[0]), goString(bufPtr(procName)))
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
		uintptr(unsafe.Pointer(bytePtr("GetCurrentProcessId"))),
	)
	if addr == 0 {
		t.Fatal("GetProcAddress returned NULL")
	}
	hMod := HMODULE(unsafe.Pointer(kernel32))
	ret := a.GetProcOrdinal(hMod, Cardinal(addr))
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
		uintptr(unsafe.Pointer(bytePtr("GetCurrentProcessId"))),
	)
	if addr == 0 {
		t.Fatal("GetProcAddress returned NULL")
	}
	hMod := HMODULE(unsafe.Pointer(kernel32))
	pHint, hintBuf := pcardinalBuf(1)
	procName := make([]byte, 256)
	ret := a.GetProcName(hMod, Cardinal(addr), pHint, bufPtr(procName))
	t.Logf("return=%d hint=0x%08X name=%s", ret, Cardinal(hintBuf[0]), goString(bufPtr(procName)))
}

func TestGetAllVAddressesOfImports(t *testing.T) {
	pid := windows.GetCurrentProcessId()
	pVAddr, vAddrBuf := pcardinalBuf(1024)
	ret := a.GetAllVAddressesOfImports(pid, 0, pVAddr, Integer(len(vAddrBuf)))
	t.Logf("return=%d (num imports found)", ret)
	for i := 0; i < int(ret) && i < 10; i++ {
		t.Logf("  vAddr[%d] = 0x%08X", i, Cardinal(vAddrBuf[i]))
	}
}

func goString(p *PChar) string {
	if p == nil {
		return ""
	}
	var buf []byte
	pb := (*byte)(unsafe.Pointer(p))
	for *pb != 0 {
		buf = append(buf, *pb)
		pb = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(pb)) + 1))
	}
	return string(buf)
}
