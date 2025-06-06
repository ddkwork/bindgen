
#pragma once
UINT64 ScriptEngineConvertNameToAddressWrapper(
    const char *FunctionOrVariableName, PBOOLEAN WasFound);
UINT32 ScriptEngineLoadFileSymbolWrapper(UINT64 BaseAddress,
                                         const char *PdbFileName,
                                         const char *CustomModuleName);
VOID ScriptEngineSetTextMessageCallbackWrapper(PVOID Handler);
UINT32 ScriptEngineUnloadAllSymbolsWrapper();
UINT32 ScriptEngineUnloadModuleSymbolWrapper(char *ModuleName);
UINT32 ScriptEngineSearchSymbolForMaskWrapper(const char *SearchMask);
BOOLEAN ScriptEngineGetFieldOffsetWrapper(CHAR *TypeName, CHAR *FieldName,
                                          UINT32 *FieldOffset);
BOOLEAN ScriptEngineGetDataTypeSizeWrapper(CHAR *TypeName, UINT64 *TypeSize);
BOOLEAN ScriptEngineCreateSymbolTableForDisassemblerWrapper(
    void *CallbackFunction);
BOOLEAN ScriptEngineConvertFileToPdbPathWrapper(const char *LocalFilePath,
                                                char *ResultPath,
                                                size_t ResultPathSize);
BOOLEAN ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper(
    const char *LocalFilePath, char *PdbFilePath, char *GuidAndAgeDetails,
    BOOLEAN Is32BitModule);
BOOLEAN ScriptEngineSymbolInitLoadWrapper(
    PMODULE_SYMBOL_DETAIL BufferToStoreDetails, UINT32 StoredLength,
    BOOLEAN DownloadIfAvailable, const char *SymbolPath, BOOLEAN IsSilentLoad);
BOOLEAN ScriptEngineShowDataBasedOnSymbolTypesWrapper(
    const char *TypeName, UINT64 Address, BOOLEAN IsStruct, PVOID BufferAddress,
    const char *AdditionalParameters);
VOID ScriptEngineSymbolAbortLoadingWrapper();
VOID ScriptEngineWrapperTestParser(const string &Expr);
VOID ScriptEngineWrapperTestParserForHwdbg(const string &Expr);
BOOLEAN ScriptAutomaticStatementsTestWrapper(const string &Expr,
                                             UINT64 ExpectationValue,
                                             BOOLEAN ExceptError);
PVOID ScriptEngineParseWrapper(char *Expr, BOOLEAN ShowErrorMessageIfAny);
VOID PrintSymbolBufferWrapper(PVOID SymbolBuffer);
UINT64 ScriptEngineWrapperGetHead(PVOID SymbolBuffer);
UINT32 ScriptEngineWrapperGetSize(PVOID SymbolBuffer);
UINT32 ScriptEngineWrapperGetPointer(PVOID SymbolBuffer);
VOID ScriptEngineWrapperRemoveSymbolBuffer(PVOID SymbolBuffer);
UINT64 ScriptEngineEvalUInt64StyleExpressionWrapper(const string &Expr,
                                                    PBOOLEAN HasError);
UINT64 ScriptEngineEvalSingleExpression(string Expr, PBOOLEAN HasError);
