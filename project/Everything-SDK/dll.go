package everything

import (
	_ "embed"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Everything struct{}

//go:embed Everything64.dll
var dllBytes []byte

var (
	dll            *windows.LazyDLL
	proc_Everything_SetSearchW      *windows.LazyProc
	proc_Everything_SetSearchA      *windows.LazyProc
	proc_Everything_SetMatchPath      *windows.LazyProc
	proc_Everything_SetMatchCase      *windows.LazyProc
	proc_Everything_SetMatchWholeWord      *windows.LazyProc
	proc_Everything_SetRegex      *windows.LazyProc
	proc_Everything_SetMax      *windows.LazyProc
	proc_Everything_SetOffset      *windows.LazyProc
	proc_Everything_SetReplyWindow      *windows.LazyProc
	proc_Everything_SetReplyID      *windows.LazyProc
	proc_Everything_SetSort      *windows.LazyProc
	proc_Everything_SetRequestFlags      *windows.LazyProc
	proc_Everything_GetMatchPath      *windows.LazyProc
	proc_Everything_GetMatchCase      *windows.LazyProc
	proc_Everything_GetMatchWholeWord      *windows.LazyProc
	proc_Everything_GetRegex      *windows.LazyProc
	proc_Everything_GetMax      *windows.LazyProc
	proc_Everything_GetOffset      *windows.LazyProc
	proc_Everything_GetSearchA      *windows.LazyProc
	proc_Everything_GetSearchW      *windows.LazyProc
	proc_Everything_GetLastError      *windows.LazyProc
	proc_Everything_GetReplyWindow      *windows.LazyProc
	proc_Everything_GetReplyID      *windows.LazyProc
	proc_Everything_GetSort      *windows.LazyProc
	proc_Everything_GetRequestFlags      *windows.LazyProc
	proc_Everything_QueryA      *windows.LazyProc
	proc_Everything_QueryW      *windows.LazyProc
	proc_Everything_IsQueryReply      *windows.LazyProc
	proc_Everything_SortResultsByPath      *windows.LazyProc
	proc_Everything_GetNumFileResults      *windows.LazyProc
	proc_Everything_GetNumFolderResults      *windows.LazyProc
	proc_Everything_GetNumResults      *windows.LazyProc
	proc_Everything_GetTotFileResults      *windows.LazyProc
	proc_Everything_GetTotFolderResults      *windows.LazyProc
	proc_Everything_GetTotResults      *windows.LazyProc
	proc_Everything_IsVolumeResult      *windows.LazyProc
	proc_Everything_IsFolderResult      *windows.LazyProc
	proc_Everything_IsFileResult      *windows.LazyProc
	proc_Everything_GetResultFileNameW      *windows.LazyProc
	proc_Everything_GetResultFileNameA      *windows.LazyProc
	proc_Everything_GetResultPathW      *windows.LazyProc
	proc_Everything_GetResultPathA      *windows.LazyProc
	proc_Everything_GetResultFullPathNameA      *windows.LazyProc
	proc_Everything_GetResultFullPathNameW      *windows.LazyProc
	proc_Everything_GetResultListSort      *windows.LazyProc
	proc_Everything_GetResultListRequestFlags      *windows.LazyProc
	proc_Everything_GetResultExtensionW      *windows.LazyProc
	proc_Everything_GetResultExtensionA      *windows.LazyProc
	proc_Everything_GetResultSize      *windows.LazyProc
	proc_Everything_GetResultDateCreated      *windows.LazyProc
	proc_Everything_GetResultDateModified      *windows.LazyProc
	proc_Everything_GetResultDateAccessed      *windows.LazyProc
	proc_Everything_GetResultAttributes      *windows.LazyProc
	proc_Everything_GetResultFileListFileNameW      *windows.LazyProc
	proc_Everything_GetResultFileListFileNameA      *windows.LazyProc
	proc_Everything_GetResultRunCount      *windows.LazyProc
	proc_Everything_GetResultDateRun      *windows.LazyProc
	proc_Everything_GetResultDateRecentlyChanged      *windows.LazyProc
	proc_Everything_GetResultHighlightedFileNameW      *windows.LazyProc
	proc_Everything_GetResultHighlightedFileNameA      *windows.LazyProc
	proc_Everything_GetResultHighlightedPathW      *windows.LazyProc
	proc_Everything_GetResultHighlightedPathA      *windows.LazyProc
	proc_Everything_GetResultHighlightedFullPathAndFileNameW      *windows.LazyProc
	proc_Everything_GetResultHighlightedFullPathAndFileNameA      *windows.LazyProc
	proc_Everything_Reset      *windows.LazyProc
	proc_Everything_CleanUp      *windows.LazyProc
	proc_Everything_GetMajorVersion      *windows.LazyProc
	proc_Everything_GetMinorVersion      *windows.LazyProc
	proc_Everything_GetRevision      *windows.LazyProc
	proc_Everything_GetBuildNumber      *windows.LazyProc
	proc_Everything_Exit      *windows.LazyProc
	proc_Everything_MSIExitAndStopService      *windows.LazyProc
	proc_Everything_MSIStartService      *windows.LazyProc
	proc_Everything_IsDBLoaded      *windows.LazyProc
	proc_Everything_IsAdmin      *windows.LazyProc
	proc_Everything_IsAppData      *windows.LazyProc
	proc_Everything_RebuildDB      *windows.LazyProc
	proc_Everything_UpdateAllFolderIndexes      *windows.LazyProc
	proc_Everything_SaveDB      *windows.LazyProc
	proc_Everything_SaveRunHistory      *windows.LazyProc
	proc_Everything_DeleteRunHistory      *windows.LazyProc
	proc_Everything_GetTargetMachine      *windows.LazyProc
	proc_Everything_IsFastSort      *windows.LazyProc
	proc_Everything_IsFileInfoIndexed      *windows.LazyProc
	proc_Everything_GetRunCountFromFileNameW      *windows.LazyProc
	proc_Everything_GetRunCountFromFileNameA      *windows.LazyProc
	proc_Everything_SetRunCountFromFileNameW      *windows.LazyProc
	proc_Everything_SetRunCountFromFileNameA      *windows.LazyProc
	proc_Everything_IncRunCountFromFileNameW      *windows.LazyProc
	proc_Everything_IncRunCountFromFileNameA      *windows.LazyProc
)

func init() {
	dll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, "Everything64.dll"))
	proc_Everything_SetSearchW = dll.NewProc("Everything_SetSearchW")
	proc_Everything_SetSearchA = dll.NewProc("Everything_SetSearchA")
	proc_Everything_SetMatchPath = dll.NewProc("Everything_SetMatchPath")
	proc_Everything_SetMatchCase = dll.NewProc("Everything_SetMatchCase")
	proc_Everything_SetMatchWholeWord = dll.NewProc("Everything_SetMatchWholeWord")
	proc_Everything_SetRegex = dll.NewProc("Everything_SetRegex")
	proc_Everything_SetMax = dll.NewProc("Everything_SetMax")
	proc_Everything_SetOffset = dll.NewProc("Everything_SetOffset")
	proc_Everything_SetReplyWindow = dll.NewProc("Everything_SetReplyWindow")
	proc_Everything_SetReplyID = dll.NewProc("Everything_SetReplyID")
	proc_Everything_SetSort = dll.NewProc("Everything_SetSort")
	proc_Everything_SetRequestFlags = dll.NewProc("Everything_SetRequestFlags")
	proc_Everything_GetMatchPath = dll.NewProc("Everything_GetMatchPath")
	proc_Everything_GetMatchCase = dll.NewProc("Everything_GetMatchCase")
	proc_Everything_GetMatchWholeWord = dll.NewProc("Everything_GetMatchWholeWord")
	proc_Everything_GetRegex = dll.NewProc("Everything_GetRegex")
	proc_Everything_GetMax = dll.NewProc("Everything_GetMax")
	proc_Everything_GetOffset = dll.NewProc("Everything_GetOffset")
	proc_Everything_GetSearchA = dll.NewProc("Everything_GetSearchA")
	proc_Everything_GetSearchW = dll.NewProc("Everything_GetSearchW")
	proc_Everything_GetLastError = dll.NewProc("Everything_GetLastError")
	proc_Everything_GetReplyWindow = dll.NewProc("Everything_GetReplyWindow")
	proc_Everything_GetReplyID = dll.NewProc("Everything_GetReplyID")
	proc_Everything_GetSort = dll.NewProc("Everything_GetSort")
	proc_Everything_GetRequestFlags = dll.NewProc("Everything_GetRequestFlags")
	proc_Everything_QueryA = dll.NewProc("Everything_QueryA")
	proc_Everything_QueryW = dll.NewProc("Everything_QueryW")
	proc_Everything_IsQueryReply = dll.NewProc("Everything_IsQueryReply")
	proc_Everything_SortResultsByPath = dll.NewProc("Everything_SortResultsByPath")
	proc_Everything_GetNumFileResults = dll.NewProc("Everything_GetNumFileResults")
	proc_Everything_GetNumFolderResults = dll.NewProc("Everything_GetNumFolderResults")
	proc_Everything_GetNumResults = dll.NewProc("Everything_GetNumResults")
	proc_Everything_GetTotFileResults = dll.NewProc("Everything_GetTotFileResults")
	proc_Everything_GetTotFolderResults = dll.NewProc("Everything_GetTotFolderResults")
	proc_Everything_GetTotResults = dll.NewProc("Everything_GetTotResults")
	proc_Everything_IsVolumeResult = dll.NewProc("Everything_IsVolumeResult")
	proc_Everything_IsFolderResult = dll.NewProc("Everything_IsFolderResult")
	proc_Everything_IsFileResult = dll.NewProc("Everything_IsFileResult")
	proc_Everything_GetResultFileNameW = dll.NewProc("Everything_GetResultFileNameW")
	proc_Everything_GetResultFileNameA = dll.NewProc("Everything_GetResultFileNameA")
	proc_Everything_GetResultPathW = dll.NewProc("Everything_GetResultPathW")
	proc_Everything_GetResultPathA = dll.NewProc("Everything_GetResultPathA")
	proc_Everything_GetResultFullPathNameA = dll.NewProc("Everything_GetResultFullPathNameA")
	proc_Everything_GetResultFullPathNameW = dll.NewProc("Everything_GetResultFullPathNameW")
	proc_Everything_GetResultListSort = dll.NewProc("Everything_GetResultListSort")
	proc_Everything_GetResultListRequestFlags = dll.NewProc("Everything_GetResultListRequestFlags")
	proc_Everything_GetResultExtensionW = dll.NewProc("Everything_GetResultExtensionW")
	proc_Everything_GetResultExtensionA = dll.NewProc("Everything_GetResultExtensionA")
	proc_Everything_GetResultSize = dll.NewProc("Everything_GetResultSize")
	proc_Everything_GetResultDateCreated = dll.NewProc("Everything_GetResultDateCreated")
	proc_Everything_GetResultDateModified = dll.NewProc("Everything_GetResultDateModified")
	proc_Everything_GetResultDateAccessed = dll.NewProc("Everything_GetResultDateAccessed")
	proc_Everything_GetResultAttributes = dll.NewProc("Everything_GetResultAttributes")
	proc_Everything_GetResultFileListFileNameW = dll.NewProc("Everything_GetResultFileListFileNameW")
	proc_Everything_GetResultFileListFileNameA = dll.NewProc("Everything_GetResultFileListFileNameA")
	proc_Everything_GetResultRunCount = dll.NewProc("Everything_GetResultRunCount")
	proc_Everything_GetResultDateRun = dll.NewProc("Everything_GetResultDateRun")
	proc_Everything_GetResultDateRecentlyChanged = dll.NewProc("Everything_GetResultDateRecentlyChanged")
	proc_Everything_GetResultHighlightedFileNameW = dll.NewProc("Everything_GetResultHighlightedFileNameW")
	proc_Everything_GetResultHighlightedFileNameA = dll.NewProc("Everything_GetResultHighlightedFileNameA")
	proc_Everything_GetResultHighlightedPathW = dll.NewProc("Everything_GetResultHighlightedPathW")
	proc_Everything_GetResultHighlightedPathA = dll.NewProc("Everything_GetResultHighlightedPathA")
	proc_Everything_GetResultHighlightedFullPathAndFileNameW = dll.NewProc("Everything_GetResultHighlightedFullPathAndFileNameW")
	proc_Everything_GetResultHighlightedFullPathAndFileNameA = dll.NewProc("Everything_GetResultHighlightedFullPathAndFileNameA")
	proc_Everything_Reset = dll.NewProc("Everything_Reset")
	proc_Everything_CleanUp = dll.NewProc("Everything_CleanUp")
	proc_Everything_GetMajorVersion = dll.NewProc("Everything_GetMajorVersion")
	proc_Everything_GetMinorVersion = dll.NewProc("Everything_GetMinorVersion")
	proc_Everything_GetRevision = dll.NewProc("Everything_GetRevision")
	proc_Everything_GetBuildNumber = dll.NewProc("Everything_GetBuildNumber")
	proc_Everything_Exit = dll.NewProc("Everything_Exit")
	proc_Everything_MSIExitAndStopService = dll.NewProc("Everything_MSIExitAndStopService")
	proc_Everything_MSIStartService = dll.NewProc("Everything_MSIStartService")
	proc_Everything_IsDBLoaded = dll.NewProc("Everything_IsDBLoaded")
	proc_Everything_IsAdmin = dll.NewProc("Everything_IsAdmin")
	proc_Everything_IsAppData = dll.NewProc("Everything_IsAppData")
	proc_Everything_RebuildDB = dll.NewProc("Everything_RebuildDB")
	proc_Everything_UpdateAllFolderIndexes = dll.NewProc("Everything_UpdateAllFolderIndexes")
	proc_Everything_SaveDB = dll.NewProc("Everything_SaveDB")
	proc_Everything_SaveRunHistory = dll.NewProc("Everything_SaveRunHistory")
	proc_Everything_DeleteRunHistory = dll.NewProc("Everything_DeleteRunHistory")
	proc_Everything_GetTargetMachine = dll.NewProc("Everything_GetTargetMachine")
	proc_Everything_IsFastSort = dll.NewProc("Everything_IsFastSort")
	proc_Everything_IsFileInfoIndexed = dll.NewProc("Everything_IsFileInfoIndexed")
	proc_Everything_GetRunCountFromFileNameW = dll.NewProc("Everything_GetRunCountFromFileNameW")
	proc_Everything_GetRunCountFromFileNameA = dll.NewProc("Everything_GetRunCountFromFileNameA")
	proc_Everything_SetRunCountFromFileNameW = dll.NewProc("Everything_SetRunCountFromFileNameW")
	proc_Everything_SetRunCountFromFileNameA = dll.NewProc("Everything_SetRunCountFromFileNameA")
	proc_Everything_IncRunCountFromFileNameW = dll.NewProc("Everything_IncRunCountFromFileNameW")
	proc_Everything_IncRunCountFromFileNameA = dll.NewProc("Everything_IncRunCountFromFileNameA")
}

func saveEmbeddedDLL(data []byte, name string) string {
	tmpDir := os.TempDir()
	p := filepath.Join(tmpDir, name)
	os.WriteFile(p, data, 0644)
	return p
}

func (e *Everything) SetSearchW(LpString *uint16) {
	proc_Everything_SetSearchW.Call(uintptr(unsafe.Pointer(LpString)))
}

func (e *Everything) SetSearchA(LpString *int8) {
	proc_Everything_SetSearchA.Call(uintptr(unsafe.Pointer(LpString)))
}

func (e *Everything) SetMatchPath(BEnable int32) {
	proc_Everything_SetMatchPath.Call(uintptr(BEnable))
}

func (e *Everything) SetMatchCase(BEnable int32) {
	proc_Everything_SetMatchCase.Call(uintptr(BEnable))
}

func (e *Everything) SetMatchWholeWord(BEnable int32) {
	proc_Everything_SetMatchWholeWord.Call(uintptr(BEnable))
}

func (e *Everything) SetRegex(BEnable int32) {
	proc_Everything_SetRegex.Call(uintptr(BEnable))
}

func (e *Everything) SetMax(DwMax uint32) {
	proc_Everything_SetMax.Call(uintptr(DwMax))
}

func (e *Everything) SetOffset(DwOffset uint32) {
	proc_Everything_SetOffset.Call(uintptr(DwOffset))
}

func (e *Everything) SetReplyWindow(HWnd unsafe.Pointer) {
	proc_Everything_SetReplyWindow.Call(uintptr(HWnd))
}

func (e *Everything) SetReplyID(DwId uint32) {
	proc_Everything_SetReplyID.Call(uintptr(DwId))
}

func (e *Everything) SetSort(DwSort uint32) {
	proc_Everything_SetSort.Call(uintptr(DwSort))
}

func (e *Everything) SetRequestFlags(DwRequestFlags uint32) {
	proc_Everything_SetRequestFlags.Call(uintptr(DwRequestFlags))
}

func (e *Everything) GetMatchPath() int32 {
	r1, _, _ := proc_Everything_GetMatchPath.Call()
	return int32(r1)
}

func (e *Everything) GetMatchCase() int32 {
	r1, _, _ := proc_Everything_GetMatchCase.Call()
	return int32(r1)
}

func (e *Everything) GetMatchWholeWord() int32 {
	r1, _, _ := proc_Everything_GetMatchWholeWord.Call()
	return int32(r1)
}

func (e *Everything) GetRegex() int32 {
	r1, _, _ := proc_Everything_GetRegex.Call()
	return int32(r1)
}

func (e *Everything) GetMax() uint32 {
	r1, _, _ := proc_Everything_GetMax.Call()
	return uint32(r1)
}

func (e *Everything) GetOffset() uint32 {
	r1, _, _ := proc_Everything_GetOffset.Call()
	return uint32(r1)
}

func (e *Everything) GetSearchA() *int8 {
	r1, _, _ := proc_Everything_GetSearchA.Call()
	return (*int8)(unsafe.Pointer(r1))
}

func (e *Everything) GetSearchW() *uint16 {
	r1, _, _ := proc_Everything_GetSearchW.Call()
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetLastError() uint32 {
	r1, _, _ := proc_Everything_GetLastError.Call()
	return uint32(r1)
}

func (e *Everything) GetReplyWindow() unsafe.Pointer {
	r1, _, _ := proc_Everything_GetReplyWindow.Call()
	return unsafe.Pointer(r1)
}

func (e *Everything) GetReplyID() uint32 {
	r1, _, _ := proc_Everything_GetReplyID.Call()
	return uint32(r1)
}

func (e *Everything) GetSort() uint32 {
	r1, _, _ := proc_Everything_GetSort.Call()
	return uint32(r1)
}

func (e *Everything) GetRequestFlags() uint32 {
	r1, _, _ := proc_Everything_GetRequestFlags.Call()
	return uint32(r1)
}

func (e *Everything) QueryA(BWait int32) int32 {
	r1, _, _ := proc_Everything_QueryA.Call(uintptr(BWait))
	return int32(r1)
}

func (e *Everything) QueryW(BWait int32) int32 {
	r1, _, _ := proc_Everything_QueryW.Call(uintptr(BWait))
	return int32(r1)
}

func (e *Everything) IsQueryReply(Message uint32, WParam int64, LParam int64, DwId uint32) int32 {
	r1, _, _ := proc_Everything_IsQueryReply.Call(uintptr(Message), *(*uintptr)(unsafe.Pointer(&WParam)), *(*uintptr)(unsafe.Pointer(&LParam)), uintptr(DwId))
	return int32(r1)
}

func (e *Everything) SortResultsByPath() {
	proc_Everything_SortResultsByPath.Call()
}

func (e *Everything) GetNumFileResults() uint32 {
	r1, _, _ := proc_Everything_GetNumFileResults.Call()
	return uint32(r1)
}

func (e *Everything) GetNumFolderResults() uint32 {
	r1, _, _ := proc_Everything_GetNumFolderResults.Call()
	return uint32(r1)
}

func (e *Everything) GetNumResults() uint32 {
	r1, _, _ := proc_Everything_GetNumResults.Call()
	return uint32(r1)
}

func (e *Everything) GetTotFileResults() uint32 {
	r1, _, _ := proc_Everything_GetTotFileResults.Call()
	return uint32(r1)
}

func (e *Everything) GetTotFolderResults() uint32 {
	r1, _, _ := proc_Everything_GetTotFolderResults.Call()
	return uint32(r1)
}

func (e *Everything) GetTotResults() uint32 {
	r1, _, _ := proc_Everything_GetTotResults.Call()
	return uint32(r1)
}

func (e *Everything) IsVolumeResult(DwIndex uint32) int32 {
	r1, _, _ := proc_Everything_IsVolumeResult.Call(uintptr(DwIndex))
	return int32(r1)
}

func (e *Everything) IsFolderResult(DwIndex uint32) int32 {
	r1, _, _ := proc_Everything_IsFolderResult.Call(uintptr(DwIndex))
	return int32(r1)
}

func (e *Everything) IsFileResult(DwIndex uint32) int32 {
	r1, _, _ := proc_Everything_IsFileResult.Call(uintptr(DwIndex))
	return int32(r1)
}

func (e *Everything) GetResultFileNameW(DwIndex uint32) *uint16 {
	r1, _, _ := proc_Everything_GetResultFileNameW.Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultFileNameA(DwIndex uint32) *int8 {
	r1, _, _ := proc_Everything_GetResultFileNameA.Call(uintptr(DwIndex))
	return (*int8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultPathW(DwIndex uint32) *uint16 {
	r1, _, _ := proc_Everything_GetResultPathW.Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultPathA(DwIndex uint32) *int8 {
	r1, _, _ := proc_Everything_GetResultPathA.Call(uintptr(DwIndex))
	return (*int8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultFullPathNameA(DwIndex uint32, Buf *int8, Bufsize uint32) uint32 {
	r1, _, _ := proc_Everything_GetResultFullPathNameA.Call(uintptr(DwIndex), uintptr(unsafe.Pointer(Buf)), uintptr(Bufsize))
	return uint32(r1)
}

func (e *Everything) GetResultFullPathNameW(DwIndex uint32, Wbuf *uint16, Wbuf_size_in_wchars uint32) uint32 {
	r1, _, _ := proc_Everything_GetResultFullPathNameW.Call(uintptr(DwIndex), uintptr(unsafe.Pointer(Wbuf)), uintptr(Wbuf_size_in_wchars))
	return uint32(r1)
}

func (e *Everything) GetResultListSort() uint32 {
	r1, _, _ := proc_Everything_GetResultListSort.Call()
	return uint32(r1)
}

func (e *Everything) GetResultListRequestFlags() uint32 {
	r1, _, _ := proc_Everything_GetResultListRequestFlags.Call()
	return uint32(r1)
}

func (e *Everything) GetResultExtensionW(DwIndex uint32) *uint16 {
	r1, _, _ := proc_Everything_GetResultExtensionW.Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultExtensionA(DwIndex uint32) *int8 {
	r1, _, _ := proc_Everything_GetResultExtensionA.Call(uintptr(DwIndex))
	return (*int8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultSize(DwIndex uint32, LpSize *LARGE_INTEGER) int32 {
	r1, _, _ := proc_Everything_GetResultSize.Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpSize)))
	return int32(r1)
}

func (e *Everything) GetResultDateCreated(DwIndex uint32, LpDateCreated *FILETIME) int32 {
	r1, _, _ := proc_Everything_GetResultDateCreated.Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateCreated)))
	return int32(r1)
}

func (e *Everything) GetResultDateModified(DwIndex uint32, LpDateModified *FILETIME) int32 {
	r1, _, _ := proc_Everything_GetResultDateModified.Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateModified)))
	return int32(r1)
}

func (e *Everything) GetResultDateAccessed(DwIndex uint32, LpDateAccessed *FILETIME) int32 {
	r1, _, _ := proc_Everything_GetResultDateAccessed.Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateAccessed)))
	return int32(r1)
}

func (e *Everything) GetResultAttributes(DwIndex uint32) uint32 {
	r1, _, _ := proc_Everything_GetResultAttributes.Call(uintptr(DwIndex))
	return uint32(r1)
}

func (e *Everything) GetResultFileListFileNameW(DwIndex uint32) *uint16 {
	r1, _, _ := proc_Everything_GetResultFileListFileNameW.Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultFileListFileNameA(DwIndex uint32) *int8 {
	r1, _, _ := proc_Everything_GetResultFileListFileNameA.Call(uintptr(DwIndex))
	return (*int8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultRunCount(DwIndex uint32) uint32 {
	r1, _, _ := proc_Everything_GetResultRunCount.Call(uintptr(DwIndex))
	return uint32(r1)
}

func (e *Everything) GetResultDateRun(DwIndex uint32, LpDateRun *FILETIME) int32 {
	r1, _, _ := proc_Everything_GetResultDateRun.Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateRun)))
	return int32(r1)
}

func (e *Everything) GetResultDateRecentlyChanged(DwIndex uint32, LpDateRecentlyChanged *FILETIME) int32 {
	r1, _, _ := proc_Everything_GetResultDateRecentlyChanged.Call(uintptr(DwIndex), uintptr(unsafe.Pointer(LpDateRecentlyChanged)))
	return int32(r1)
}

func (e *Everything) GetResultHighlightedFileNameW(DwIndex uint32) *uint16 {
	r1, _, _ := proc_Everything_GetResultHighlightedFileNameW.Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedFileNameA(DwIndex uint32) *int8 {
	r1, _, _ := proc_Everything_GetResultHighlightedFileNameA.Call(uintptr(DwIndex))
	return (*int8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedPathW(DwIndex uint32) *uint16 {
	r1, _, _ := proc_Everything_GetResultHighlightedPathW.Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedPathA(DwIndex uint32) *int8 {
	r1, _, _ := proc_Everything_GetResultHighlightedPathA.Call(uintptr(DwIndex))
	return (*int8)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedFullPathAndFileNameW(DwIndex uint32) *uint16 {
	r1, _, _ := proc_Everything_GetResultHighlightedFullPathAndFileNameW.Call(uintptr(DwIndex))
	return (*uint16)(unsafe.Pointer(r1))
}

func (e *Everything) GetResultHighlightedFullPathAndFileNameA(DwIndex uint32) *int8 {
	r1, _, _ := proc_Everything_GetResultHighlightedFullPathAndFileNameA.Call(uintptr(DwIndex))
	return (*int8)(unsafe.Pointer(r1))
}

func (e *Everything) Reset() {
	proc_Everything_Reset.Call()
}

func (e *Everything) CleanUp() {
	proc_Everything_CleanUp.Call()
}

func (e *Everything) GetMajorVersion() uint32 {
	r1, _, _ := proc_Everything_GetMajorVersion.Call()
	return uint32(r1)
}

func (e *Everything) GetMinorVersion() uint32 {
	r1, _, _ := proc_Everything_GetMinorVersion.Call()
	return uint32(r1)
}

func (e *Everything) GetRevision() uint32 {
	r1, _, _ := proc_Everything_GetRevision.Call()
	return uint32(r1)
}

func (e *Everything) GetBuildNumber() uint32 {
	r1, _, _ := proc_Everything_GetBuildNumber.Call()
	return uint32(r1)
}

func (e *Everything) Exit() int32 {
	r1, _, _ := proc_Everything_Exit.Call()
	return int32(r1)
}

func (e *Everything) MSIExitAndStopService(Msihandle unsafe.Pointer) uint32 {
	r1, _, _ := proc_Everything_MSIExitAndStopService.Call(uintptr(Msihandle))
	return uint32(r1)
}

func (e *Everything) MSIStartService(Msihandle unsafe.Pointer) uint32 {
	r1, _, _ := proc_Everything_MSIStartService.Call(uintptr(Msihandle))
	return uint32(r1)
}

func (e *Everything) IsDBLoaded() int32 {
	r1, _, _ := proc_Everything_IsDBLoaded.Call()
	return int32(r1)
}

func (e *Everything) IsAdmin() int32 {
	r1, _, _ := proc_Everything_IsAdmin.Call()
	return int32(r1)
}

func (e *Everything) IsAppData() int32 {
	r1, _, _ := proc_Everything_IsAppData.Call()
	return int32(r1)
}

func (e *Everything) RebuildDB() int32 {
	r1, _, _ := proc_Everything_RebuildDB.Call()
	return int32(r1)
}

func (e *Everything) UpdateAllFolderIndexes() int32 {
	r1, _, _ := proc_Everything_UpdateAllFolderIndexes.Call()
	return int32(r1)
}

func (e *Everything) SaveDB() int32 {
	r1, _, _ := proc_Everything_SaveDB.Call()
	return int32(r1)
}

func (e *Everything) SaveRunHistory() int32 {
	r1, _, _ := proc_Everything_SaveRunHistory.Call()
	return int32(r1)
}

func (e *Everything) DeleteRunHistory() int32 {
	r1, _, _ := proc_Everything_DeleteRunHistory.Call()
	return int32(r1)
}

func (e *Everything) GetTargetMachine() uint32 {
	r1, _, _ := proc_Everything_GetTargetMachine.Call()
	return uint32(r1)
}

func (e *Everything) IsFastSort(SortType uint32) int32 {
	r1, _, _ := proc_Everything_IsFastSort.Call(uintptr(SortType))
	return int32(r1)
}

func (e *Everything) IsFileInfoIndexed(FileInfoType uint32) int32 {
	r1, _, _ := proc_Everything_IsFileInfoIndexed.Call(uintptr(FileInfoType))
	return int32(r1)
}

func (e *Everything) GetRunCountFromFileNameW(LpFileName *uint16) uint32 {
	r1, _, _ := proc_Everything_GetRunCountFromFileNameW.Call(uintptr(unsafe.Pointer(LpFileName)))
	return uint32(r1)
}

func (e *Everything) GetRunCountFromFileNameA(LpFileName *int8) uint32 {
	r1, _, _ := proc_Everything_GetRunCountFromFileNameA.Call(uintptr(unsafe.Pointer(LpFileName)))
	return uint32(r1)
}

func (e *Everything) SetRunCountFromFileNameW(LpFileName *uint16, DwRunCount uint32) int32 {
	r1, _, _ := proc_Everything_SetRunCountFromFileNameW.Call(uintptr(unsafe.Pointer(LpFileName)), uintptr(DwRunCount))
	return int32(r1)
}

func (e *Everything) SetRunCountFromFileNameA(LpFileName *int8, DwRunCount uint32) int32 {
	r1, _, _ := proc_Everything_SetRunCountFromFileNameA.Call(uintptr(unsafe.Pointer(LpFileName)), uintptr(DwRunCount))
	return int32(r1)
}

func (e *Everything) IncRunCountFromFileNameW(LpFileName *uint16) uint32 {
	r1, _, _ := proc_Everything_IncRunCountFromFileNameW.Call(uintptr(unsafe.Pointer(LpFileName)))
	return uint32(r1)
}

func (e *Everything) IncRunCountFromFileNameA(LpFileName *int8) uint32 {
	r1, _, _ := proc_Everything_IncRunCountFromFileNameA.Call(uintptr(unsafe.Pointer(LpFileName)))
	return uint32(r1)
}

