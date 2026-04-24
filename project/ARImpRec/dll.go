package ipmrec

import (
	_ "embed"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Ipmrec struct{}

//go:embed ARImpRec.dll
var dllBytes []byte

var (
	dll            *windows.LazyDLL
	proc_UnpackPdataSection      *windows.LazyProc
	proc_GetNameFileOptimized      *windows.LazyProc
	proc_RebuildSectionsFromArmadillo      *windows.LazyProc
	proc_TryGetImportedFunction      *windows.LazyProc
	proc_SearchAndRebuildImportsNoNewSection      *windows.LazyProc
	proc_SearchAndRebuildImportsIATOptimized      *windows.LazyProc
	proc_SearchAndRebuildImports      *windows.LazyProc
	proc_GetProcNameAndOrdinal      *windows.LazyProc
	proc_GetProcOrdinal      *windows.LazyProc
	proc_GetProcName      *windows.LazyProc
	proc_GetAllVAddressesOfImports      *windows.LazyProc
)

func init() {
	dll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, "ARImpRec.dll"))
	proc_UnpackPdataSection = dll.NewProc("UnpackPdataSection@12")
	proc_GetNameFileOptimized = dll.NewProc("GetNameFileOptimized@8")
	proc_RebuildSectionsFromArmadillo = dll.NewProc("RebuildSectionsFromArmadillo@12")
	proc_TryGetImportedFunction = dll.NewProc("TryGetImportedFunction@24")
	proc_SearchAndRebuildImportsNoNewSection = dll.NewProc("SearchAndRebuildImportsNoNewSection@28")
	proc_SearchAndRebuildImportsIATOptimized = dll.NewProc("SearchAndRebuildImportsIATOptimized@28")
	proc_SearchAndRebuildImports = dll.NewProc("SearchAndRebuildImports@28")
	proc_GetProcNameAndOrdinal = dll.NewProc("GetProcNameAndOrdinal@20")
	proc_GetProcOrdinal = dll.NewProc("GetProcOrdinal@8")
	proc_GetProcName = dll.NewProc("GetProcName@16")
	proc_GetAllVAddressesOfImports = dll.NewProc("GetAllVAddressesOfImports@16")
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

func (i *Ipmrec) UnpackPdataSection(MSNameOfProtected *PChar, MSNameOfDumped *PChar, MSWarning *PChar) Integer {
	r1, _, _ := proc_UnpackPdataSection.Call(uintptr(unsafe.Pointer(MSNameOfProtected)), uintptr(unsafe.Pointer(MSNameOfDumped)), uintptr(unsafe.Pointer(MSWarning)))
	return Integer(r1)
}

func (i *Ipmrec) GetNameFileOptimized(MSFileNameOrig *PChar, MSFileNameOptimized *PChar) Integer {
	r1, _, _ := proc_GetNameFileOptimized.Call(uintptr(unsafe.Pointer(MSFileNameOrig)), uintptr(unsafe.Pointer(MSFileNameOptimized)))
	return Integer(r1)
}

func (i *Ipmrec) RebuildSectionsFromArmadillo(MSNameOfProtected *PChar, MSNameOfDumped *PChar, MSWarning *PChar) Integer {
	r1, _, _ := proc_RebuildSectionsFromArmadillo.Call(uintptr(unsafe.Pointer(MSNameOfProtected)), uintptr(unsafe.Pointer(MSNameOfDumped)), uintptr(unsafe.Pointer(MSWarning)))
	return Integer(r1)
}

func (i *Ipmrec) TryGetImportedFunction(IRProcessId Cardinal, IRVAddress Cardinal, IROrdinal *PCardinal, IRHint *PCardinal, IRFunctionName *PChar, IRModule *PChar) Integer {
	r1, _, _ := proc_TryGetImportedFunction.Call(uintptr(unsafe.Pointer(&IRProcessId)), uintptr(unsafe.Pointer(&IRVAddress)), uintptr(unsafe.Pointer(IROrdinal)), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRFunctionName)), uintptr(unsafe.Pointer(IRModule)))
	return Integer(r1)
}

func (i *Ipmrec) SearchAndRebuildImportsNoNewSection(IRProcessId Cardinal, IRNameOfDumped *PChar, IROEP Cardinal, IRSaveOEPToFile IRSaveOEPToFile, IRIATRVA *PCardinal, IRIATSize *PCardinal, IRWarning *PChar) Integer {
	r1, _, _ := proc_SearchAndRebuildImportsNoNewSection.Call(uintptr(unsafe.Pointer(&IRProcessId)), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(unsafe.Pointer(&IROEP)), uintptr(unsafe.Pointer(&IRSaveOEPToFile)), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return Integer(r1)
}

func (i *Ipmrec) SearchAndRebuildImportsIATOptimized(IRProcessId Cardinal, IRNameOfDumped *PChar, IROEP Cardinal, IRSaveOEPToFile IRSaveOEPToFile, IRIATRVA *PCardinal, IRIATSize *PCardinal, IRWarning *PChar) Integer {
	r1, _, _ := proc_SearchAndRebuildImportsIATOptimized.Call(uintptr(unsafe.Pointer(&IRProcessId)), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(unsafe.Pointer(&IROEP)), uintptr(unsafe.Pointer(&IRSaveOEPToFile)), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return Integer(r1)
}

func (i *Ipmrec) SearchAndRebuildImports(IRProcessId Cardinal, IRNameOfDumped *PChar, IROEP Cardinal, IRSaveOEPToFile IRSaveOEPToFile, IRIATRVA *PCardinal, IRIATSize *PCardinal, IRWarning *PChar) Integer {
	r1, _, _ := proc_SearchAndRebuildImports.Call(uintptr(unsafe.Pointer(&IRProcessId)), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(unsafe.Pointer(&IROEP)), uintptr(unsafe.Pointer(&IRSaveOEPToFile)), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return Integer(r1)
}

func (i *Ipmrec) GetProcNameAndOrdinal(IRHModule HMODULE, IRAddress Cardinal, IROrdinal *PCardinal, IRHint *PCardinal, IRProcName *PChar) Integer {
	r1, _, _ := proc_GetProcNameAndOrdinal.Call(uintptr(unsafe.Pointer(IRHModule)), uintptr(unsafe.Pointer(&IRAddress)), uintptr(unsafe.Pointer(IROrdinal)), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRProcName)))
	return Integer(r1)
}

func (i *Ipmrec) GetProcOrdinal(IRHModule HMODULE, IRAddress Cardinal) Cardinal {
	r1, _, _ := proc_GetProcOrdinal.Call(uintptr(unsafe.Pointer(IRHModule)), uintptr(unsafe.Pointer(&IRAddress)))
	return Cardinal(r1)
}

func (i *Ipmrec) GetProcName(IRHModule HMODULE, IRAddress Cardinal, IRHint *PCardinal, IRProcName *PChar) Integer {
	r1, _, _ := proc_GetProcName.Call(uintptr(unsafe.Pointer(IRHModule)), uintptr(unsafe.Pointer(&IRAddress)), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRProcName)))
	return Integer(r1)
}

func (i *Ipmrec) GetAllVAddressesOfImports(IRProcessId Cardinal, IROEP Cardinal, IRVAddressImports *PCardinal, IRNumberOfImports Integer) Integer {
	r1, _, _ := proc_GetAllVAddressesOfImports.Call(uintptr(unsafe.Pointer(&IRProcessId)), uintptr(unsafe.Pointer(&IROEP)), uintptr(unsafe.Pointer(IRVAddressImports)), uintptr(unsafe.Pointer(&IRNumberOfImports)))
	return Integer(r1)
}

