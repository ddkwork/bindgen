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
	dll                                      *windows.LazyDLL
	proc_malloc                              *windows.LazyProc
	proc_calloc                              *windows.LazyProc
	proc_realloc                             *windows.LazyProc
	proc_free                                *windows.LazyProc
	proc_exit                                *windows.LazyProc
	proc_abort                               *windows.LazyProc
	proc_atexit                              *windows.LazyProc
	proc_memcpy                              *windows.LazyProc
	proc_memmove                             *windows.LazyProc
	proc_strcpy                              *windows.LazyProc
	proc_strncpy                             *windows.LazyProc
	proc_strcat                              *windows.LazyProc
	proc_strlen                              *windows.LazyProc
	proc_memcmp                              *windows.LazyProc
	proc_strcmp                              *windows.LazyProc
	proc_strncmp                             *windows.LazyProc
	proc_memset                              *windows.LazyProc
	proc_strchr                              *windows.LazyProc
	proc_UnpackPdataSection                  *windows.LazyProc
	proc_GetNameFileOptimized                *windows.LazyProc
	proc_RebuildSectionsFromArmadillo        *windows.LazyProc
	proc_TryGetImportedFunction              *windows.LazyProc
	proc_SearchAndRebuildImportsNoNewSection *windows.LazyProc
	proc_SearchAndRebuildImportsIATOptimized *windows.LazyProc
	proc_SearchAndRebuildImports             *windows.LazyProc
	proc_GetProcNameAndOrdinal               *windows.LazyProc
	proc_GetProcOrdinal                      *windows.LazyProc
	proc_GetProcName                         *windows.LazyProc
	proc_GetAllVAddressesOfImports           *windows.LazyProc
)

func init() {
	dll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, "ARImpRec.dll"))
	proc_malloc = dll.NewProc("malloc")
	proc_calloc = dll.NewProc("calloc")
	proc_realloc = dll.NewProc("realloc")
	proc_free = dll.NewProc("free")
	proc_exit = dll.NewProc("exit")
	proc_abort = dll.NewProc("abort")
	proc_atexit = dll.NewProc("atexit")
	proc_memcpy = dll.NewProc("memcpy")
	proc_memmove = dll.NewProc("memmove")
	proc_strcpy = dll.NewProc("strcpy")
	proc_strncpy = dll.NewProc("strncpy")
	proc_strcat = dll.NewProc("strcat")
	proc_strlen = dll.NewProc("strlen")
	proc_memcmp = dll.NewProc("memcmp")
	proc_strcmp = dll.NewProc("strcmp")
	proc_strncmp = dll.NewProc("strncmp")
	proc_memset = dll.NewProc("memset")
	proc_strchr = dll.NewProc("strchr")
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
	os.WriteFile(p, data, 0o644)
	return p
}

func (i *Ipmrec) Malloc(Size uint32) unsafe.Pointer {
	r1, _, _ := proc_malloc.Call(uintptr(Size))
	return unsafe.Pointer(r1)
}

func (i *Ipmrec) Calloc(Count uint32, Size uint32) unsafe.Pointer {
	r1, _, _ := proc_calloc.Call(uintptr(Count), uintptr(Size))
	return unsafe.Pointer(r1)
}

func (i *Ipmrec) Realloc(Ptr unsafe.Pointer, Size uint32) unsafe.Pointer {
	r1, _, _ := proc_realloc.Call(uintptr(Ptr), uintptr(Size))
	return unsafe.Pointer(r1)
}

func (i *Ipmrec) Free(Ptr unsafe.Pointer) {
	proc_free.Call(uintptr(Ptr))
}

func (i *Ipmrec) Exit(Code int32) {
	proc_exit.Call(uintptr(Code))
}

func (i *Ipmrec) Abort() {
	proc_abort.Call()
}

func (i *Ipmrec) Atexit(Func *interface{}) int32 {
	r1, _, _ := proc_atexit.Call(uintptr(unsafe.Pointer(Func)))
	return int32(r1)
}

func (i *Ipmrec) Memcpy(Dest unsafe.Pointer, Src unsafe.Pointer, N uint32) unsafe.Pointer {
	r1, _, _ := proc_memcpy.Call(uintptr(Dest), uintptr(Src), uintptr(N))
	return unsafe.Pointer(r1)
}

func (i *Ipmrec) Memmove(Dest unsafe.Pointer, Src unsafe.Pointer, N uint32) unsafe.Pointer {
	r1, _, _ := proc_memmove.Call(uintptr(Dest), uintptr(Src), uintptr(N))
	return unsafe.Pointer(r1)
}

func (i *Ipmrec) Strcpy(Dest *int8, Src *int8) *int8 {
	r1, _, _ := proc_strcpy.Call(uintptr(unsafe.Pointer(Dest)), uintptr(unsafe.Pointer(Src)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Ipmrec) Strncpy(Dest *int8, Src *int8, N uint32) *int8 {
	r1, _, _ := proc_strncpy.Call(uintptr(unsafe.Pointer(Dest)), uintptr(unsafe.Pointer(Src)), uintptr(N))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Ipmrec) Strcat(Dest *int8, Src *int8) *int8 {
	r1, _, _ := proc_strcat.Call(uintptr(unsafe.Pointer(Dest)), uintptr(unsafe.Pointer(Src)))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Ipmrec) Strlen(Str *int8) uint32 {
	r1, _, _ := proc_strlen.Call(uintptr(unsafe.Pointer(Str)))
	return uint32(r1)
}

func (i *Ipmrec) Memcmp(S1 unsafe.Pointer, S2 unsafe.Pointer, N uint32) int32 {
	r1, _, _ := proc_memcmp.Call(uintptr(S1), uintptr(S2), uintptr(N))
	return int32(r1)
}

func (i *Ipmrec) Strcmp(S1 *int8, S2 *int8) int32 {
	r1, _, _ := proc_strcmp.Call(uintptr(unsafe.Pointer(S1)), uintptr(unsafe.Pointer(S2)))
	return int32(r1)
}

func (i *Ipmrec) Strncmp(S1 *int8, S2 *int8, N uint32) int32 {
	r1, _, _ := proc_strncmp.Call(uintptr(unsafe.Pointer(S1)), uintptr(unsafe.Pointer(S2)), uintptr(N))
	return int32(r1)
}

func (i *Ipmrec) Memset(C unsafe.Pointer, Val int32, N uint32) unsafe.Pointer {
	r1, _, _ := proc_memset.Call(uintptr(C), uintptr(Val), uintptr(N))
	return unsafe.Pointer(r1)
}

func (i *Ipmrec) Strchr(S *int8, C int32) *int8 {
	r1, _, _ := proc_strchr.Call(uintptr(unsafe.Pointer(S)), uintptr(C))
	return (*int8)(unsafe.Pointer(r1))
}

func (i *Ipmrec) UnpackPdataSection(MSNameOfProtected *int8, MSNameOfDumped *int8, MSWarning *int8) int32 {
	r1, _, _ := proc_UnpackPdataSection.Call(uintptr(unsafe.Pointer(MSNameOfProtected)), uintptr(unsafe.Pointer(MSNameOfDumped)), uintptr(unsafe.Pointer(MSWarning)))
	return int32(r1)
}

func (i *Ipmrec) GetNameFileOptimized(MSFileNameOrig *int8, MSFileNameOptimized *int8) int32 {
	r1, _, _ := proc_GetNameFileOptimized.Call(uintptr(unsafe.Pointer(MSFileNameOrig)), uintptr(unsafe.Pointer(MSFileNameOptimized)))
	return int32(r1)
}

func (i *Ipmrec) RebuildSectionsFromArmadillo(MSNameOfProtected *int8, MSNameOfDumped *int8, MSWarning *int8) int32 {
	r1, _, _ := proc_RebuildSectionsFromArmadillo.Call(uintptr(unsafe.Pointer(MSNameOfProtected)), uintptr(unsafe.Pointer(MSNameOfDumped)), uintptr(unsafe.Pointer(MSWarning)))
	return int32(r1)
}

func (i *Ipmrec) TryGetImportedFunction(IRProcessId uint32, IRVAddress uint32, IROrdinal **uint32, IRHint **uint32, IRFunctionName *int8, IRModule *int8) int32 {
	r1, _, _ := proc_TryGetImportedFunction.Call(uintptr(IRProcessId), uintptr(IRVAddress), uintptr(unsafe.Pointer(IROrdinal)), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRFunctionName)), uintptr(unsafe.Pointer(IRModule)))
	return int32(r1)
}

func (i *Ipmrec) SearchAndRebuildImportsNoNewSection(IRProcessId uint32, IRNameOfDumped *int8, IROEP uint32, IRSaveOEPToFile uint32, IRIATRVA **uint32, IRIATSize **uint32, IRWarning *int8) int32 {
	r1, _, _ := proc_SearchAndRebuildImportsNoNewSection.Call(uintptr(IRProcessId), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(IROEP), uintptr(IRSaveOEPToFile), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return int32(r1)
}

func (i *Ipmrec) SearchAndRebuildImportsIATOptimized(IRProcessId uint32, IRNameOfDumped *int8, IROEP uint32, IRSaveOEPToFile uint32, IRIATRVA **uint32, IRIATSize **uint32, IRWarning *int8) int32 {
	r1, _, _ := proc_SearchAndRebuildImportsIATOptimized.Call(uintptr(IRProcessId), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(IROEP), uintptr(IRSaveOEPToFile), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return int32(r1)
}

func (i *Ipmrec) SearchAndRebuildImports(IRProcessId uint32, IRNameOfDumped *int8, IROEP uint32, IRSaveOEPToFile uint32, IRIATRVA **uint32, IRIATSize **uint32, IRWarning *int8) int32 {
	r1, _, _ := proc_SearchAndRebuildImports.Call(uintptr(IRProcessId), uintptr(unsafe.Pointer(IRNameOfDumped)), uintptr(IROEP), uintptr(IRSaveOEPToFile), uintptr(unsafe.Pointer(IRIATRVA)), uintptr(unsafe.Pointer(IRIATSize)), uintptr(unsafe.Pointer(IRWarning)))
	return int32(r1)
}

func (i *Ipmrec) GetProcNameAndOrdinal(IRHModule uintptr, IRAddress uint32, IROrdinal **uint32, IRHint **uint32, IRProcName *int8) int32 {
	r1, _, _ := proc_GetProcNameAndOrdinal.Call(IRHModule, uintptr(IRAddress), uintptr(unsafe.Pointer(IROrdinal)), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRProcName)))
	return int32(r1)
}

func (i *Ipmrec) GetProcOrdinal(IRHModule uintptr, IRAddress uint32) uint32 {
	r1, _, _ := proc_GetProcOrdinal.Call(IRHModule, uintptr(IRAddress))
	return uint32(r1)
}

func (i *Ipmrec) GetProcName(IRHModule uintptr, IRAddress uint32, IRHint **uint32, IRProcName *int8) int32 {
	r1, _, _ := proc_GetProcName.Call(IRHModule, uintptr(IRAddress), uintptr(unsafe.Pointer(IRHint)), uintptr(unsafe.Pointer(IRProcName)))
	return int32(r1)
}

func (i *Ipmrec) GetAllVAddressesOfImports(IRProcessId uint32, IROEP uint32, IRVAddressImports **uint32, IRNumberOfImports int32) int32 {
	r1, _, _ := proc_GetAllVAddressesOfImports.Call(uintptr(IRProcessId), uintptr(IROEP), uintptr(unsafe.Pointer(IRVAddressImports)), uintptr(IRNumberOfImports))
	return int32(r1)
}
