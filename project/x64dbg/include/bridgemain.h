#ifndef BRIDGEMAINApi_H
#define BRIDGEMAINApi_H
#include "_scriptapi.h"

namespace bridgemain {
    ApiResponse BridgeLoadLibraryCheckedW_fake(const wchar_t *szDll, bool allowFailure) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeLoadLibraryCheckedA_fake(std::string szDll, bool allowFailure) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeSettingGet_fake(std::string section, std::string key, std::string value) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeSettingGetUint_fake(std::string section, std::string key, duint *value) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeSettingSet_fake(std::string section, std::string key, std::string value) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeSettingSetUint_fake(std::string section, std::string key, duint value) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeSettingFlush_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeSettingRead_fake(int *errorLine) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeGetDbgVersion_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeIsProcessElevated_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeGetNtBuildNumber_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeUserDirectory_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse BridgeIsARM64Emulated_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgInit_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgExit_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgMemRead_fake(duint va, void *dest, duint size) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgMemWrite_fake(duint va, const void *src, duint size) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgMemGetPageSize_fake(duint base) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgMemFindBaseAddr_fake(duint addr, duint *size) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgCmdExec_fake(std::string cmd) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgCmdExecDirect_fake(std::string cmd) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgMemMap_fake(MEMMAP *memmap) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgIsValidExpression_fake(std::string expression) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgIsDebugging_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgIsJumpGoingToExecute_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetLabelAt_fake(duint addr, SEGMENTREG segment, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSetLabelAt_fake(duint addr, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgClearLabelRange_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetCommentAt_fake(duint addr, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSetCommentAt_fake(duint addr, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgClearCommentRange_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetBookmarkAt_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSetBookmarkAt_fake(duint addr, bool isbookmark) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgClearBookmarkRange_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetModuleAt_fake(duint addr, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetBpxTypeAt_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgValFromString_fake(std::string string) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetRegDumpEx_fake(REGDUMP *regdump, size_t size) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgValToString_fake(std::string string, duint value) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgMemIsValidReadPtr_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetBpList_fake(BPXTYPE type, BPMAP *list) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetFunctionTypeAt_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetLoopTypeAt_fake(duint addr, int depth) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetBranchDestination_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptLoad_fake(std::string filename) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptUnload_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptRun_fake(int destline) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptStep_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptBpToggle_fake(int line) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptBpGet_fake(int line) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptCmdExec_fake(std::string command) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptAbort_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptGetLineType_fake(int line) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptSetIp_fake(int line) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgScriptGetBranchInfo_fake(int line, SCRIPTBRANCH *info) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSymbolEnum_fake(duint base, CBSYMBOLENUM cbSymbolEnum, void *user) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSymbolEnumFromCache_fake(duint base, CBSYMBOLENUM cbSymbolEnum, void *user) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSymbolEnumRange_fake(duint start, duint end, unsigned int symbolMask, CBSYMBOLENUM cbSymbolEnum, void *user) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgAssembleAt_fake(duint addr, std::string instruction) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgModBaseFromName_fake(std::string name) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgDisasmAt_fake(duint addr, DISASM_INSTR *instr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgStackCommentGet_fake(duint addr, STACK_COMMENT *comment) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetThreadList_fake(THREADLIST *list) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSettingsUpdated_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgDisasmFastAt_fake(duint addr, BASIC_INSTRUCTION_INFO *basicinfo) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgMenuEntryClicked_fake(int hEntry) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgFunctionGet_fake(duint addr, duint *start, duint *end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgFunctionOverlaps_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgFunctionAdd_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgFunctionDel_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgArgumentGet_fake(duint addr, duint *start, duint *end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgArgumentOverlaps_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgArgumentAdd_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgArgumentDel_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgLoopGet_fake(int depth, duint addr, duint *start, duint *end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgLoopOverlaps_fake(int depth, duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgLoopAdd_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgLoopDel_fake(int depth, duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgXrefAdd_fake(duint addr, duint from) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgXrefDelAll_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgXrefGet_fake(duint addr, XREF_INFO *info) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetXrefCountAt_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetXrefTypeAt_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgIsRunLocked_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgIsBpDisabled_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSetAutoCommentAt_fake(duint addr, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgClearAutoCommentRange_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSetAutoLabelAt_fake(duint addr, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgClearAutoLabelRange_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSetAutoBookmarkAt_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgClearAutoBookmarkRange_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSetAutoFunctionAt_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgClearAutoFunctionRange_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetStringAt_fake(duint addr, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgFunctions_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgIsRunning_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetTimeWastedCounter_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetArgTypeAt_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetEncodeTypeAt_fake(duint addr, duint size) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetEncodeSizeAt_fake(duint addr, duint codesize) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSetEncodeType_fake(duint addr, duint size, ENCODETYPE type) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgDelEncodeTypeRange_fake(duint start, duint end) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgDelEncodeTypeSegment_fake(duint start) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetWatchList_fake(ListOf(WATCHINFO) list) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgSelChanged_fake(int hWindow, duint VA) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetProcessHandle_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetThreadHandle_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetProcessId_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetThreadId_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetPebAddress_fake(DWORD ProcessId) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetTebAddress_fake(DWORD ThreadId) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgAnalyzeFunction_fake(duint entry, BridgeCFGraphList *graph) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgEval_fake(std::string expression, bool *DEFAULT_PARAM(success, nullptr)) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetSymbolInfo_fake(const SYMBOLPTR *symbolptr, SYMBOLINFO *info) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetDebugEngine_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgGetSymbolInfoAt_fake(duint addr, SYMBOLINFO *info) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgXrefAddMulti_fake(const XREF_EDGE *edges, duint count) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse DbgMenuPrepare_fake(GUIMENUTYPE hMenu) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiTranslateText_fake(std::string Source) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiDisasmAt_fake(duint addr, duint cip) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSetDebugState_fake(DBGSTATE state) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSetDebugStateFast_fake(DBGSTATE state) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAddLogMessage_fake(std::string msg) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAddLogMessageHtml_fake(std::string msg) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiLogClear_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiLogSave_fake(std::string filename) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiLogRedirect_fake(std::string filename) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiLogRedirectStop_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateAllViews_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateRegisterView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateDisassemblyView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateBreakpointsView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateWindowTitle_fake(std::string filename) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiGetWindowHandle_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiDumpAt_fake(duint va) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptAdd_fake(int count, std::string *lines) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptClear_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptSetIp_fake(int line) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptError_fake(int line, std::string message) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptSetTitle_fake(std::string title) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptSetInfoLine_fake(int line, std::string info) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptMessage_fake(std::string message) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptMsgyn_fake(std::string message) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiScriptEnableHighlighting_fake(bool enable) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSymbolLogAdd_fake(std::string message) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSymbolLogClear_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSymbolSetProgress_fake(int percent) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSymbolUpdateModuleList_fake(int count, SYMBOLMODULEINFO *modules) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSymbolRefreshCurrent_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceAddColumn_fake(int width, std::string title) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceSetRowCount_fake(int count) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceGetRowCount_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceSearchGetRowCount_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceDeleteAllColumns_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceInitialize_fake(std::string name) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceSetCellContent_fake(int row, int col, std::string str) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceGetCellContent_fake(int row, int col) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceSearchGetCellContent_fake(int row, int col) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceReloadData_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceSetSingleSelection_fake(int index, bool scroll) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceSetProgress_fake(int progress) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceSetCurrentTaskProgress_fake(int progress, std::string taskTitle) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceSetSearchStartCol_fake(int col) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiStackDumpAt_fake(duint addr, duint csp) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateDumpView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateWatchView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateThreadView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateMemoryView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAddRecentFile_fake(std::string file) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSetLastException_fake(unsigned int exception) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiGetDisassembly_fake(duint addr, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSelectionGet_fake(GUISELECTIONTYPE hWindow, SELECTIONDATA *selection) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSelectionSet_fake(GUISELECTIONTYPE hWindow, const SELECTIONDATA *selection) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiGetLineWindow_fake(std::string title, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAutoCompleteAddCmd_fake(std::string cmd) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAutoCompleteDelCmd_fake(std::string cmd) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAutoCompleteClearAll_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAddStatusBarMessage_fake(std::string msg) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateSideBar_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiRepaintTableView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdatePatches_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateCallStack_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateSEHChain_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiLoadSourceFileEx_fake(std::string path, duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiMenuSetEntryChecked_fake(int hEntry, bool checked) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiMenuSetVisible_fake(int hMenu, bool visible) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiMenuSetEntryVisible_fake(int hEntry, bool visible) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiMenuSetName_fake(int hMenu, std::string name) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiMenuSetEntryName_fake(int hEntry, std::string name) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiMenuSetEntryHotkey_fake(int hEntry, std::string hack) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiShowCpu_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiShowThreads_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateTimeWastedCounter_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiDumpAtN_fake(duint va, int index) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiDisplayWarning_fake(std::string title, std::string text) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateArgumentWidget_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiFocusView_fake(int hWindow) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiIsUpdateDisabled_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateEnable_fake(bool updateNow) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateDisable_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiGraphAt_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateGraphView_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiDisableLog_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiEnableLog_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiIsLogEnabled_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAddFavouriteTool_fake(std::string name, std::string description) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAddFavouriteCommand_fake(std::string name, std::string shortcut) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSetFavouriteToolShortcut_fake(std::string name, std::string shortcut) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiFoldDisassembly_fake(duint startAddress, duint length) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSelectInMemoryMap_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiGetActiveView_fake(ACTIVEVIEW *activeView) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiAddInfoLine_fake(std::string infoLine) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiProcessEvents_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiTypeClear_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateTypeWidget_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiCloseApplication_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiFlushLog_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiReferenceAddCommand_fake(std::string title, std::string command) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiUpdateTraceBrowser_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiOpenTraceFile_fake(std::string fileName) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiInvalidateSymbolSource_fake(duint base) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiShowReferences_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiSelectInSymbolsTab_fake(duint addr) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiGotoTrace_fake(duint index) {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiShowTrace_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

    ApiResponse GuiGetMainThreadId_fake() {
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success = true,
                // .type="uintptr",
                //.result=handle
        };
    }

}// namespace bridgemain

#endif// BRIDGEMAINApi_H