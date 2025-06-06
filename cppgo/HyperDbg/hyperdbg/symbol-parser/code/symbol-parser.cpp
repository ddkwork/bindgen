
#include "pch.h"
std::vector<PSYMBOL_LOADED_MODULE_DETAILS> g_LoadedModules;
BOOLEAN g_IsLoadedModulesInitialized = FALSE;
BOOLEAN g_AbortLoadingExecution = FALSE;
CHAR *g_CurrentModuleName = NULL;
PVOID g_MessageHandler = NULL;
SymbolMapCallback g_SymbolMapForDisassembler = NULL;

void SymSetTextMessageCallback(PVOID Handler) {
  g_MessageHandler = Handler;
  pdbex_set_logging_method_export(Handler);
}

VOID ShowMessages(const char *Fmt, ...) {
  va_list ArgList;
  va_list Args;
  if (g_MessageHandler == NULL) {
    va_start(Args, Fmt);
    vprintf(Fmt, Args);
    va_end(Args);
  } else {
    char TempMessage[COMMUNICATION_BUFFER_SIZE +
                     TCP_END_OF_BUFFER_CHARS_COUNT] = {0};
    va_start(ArgList, Fmt);
    int sprintfresult = vsprintf_s(TempMessage, Fmt, ArgList);
    va_end(ArgList);
    if (sprintfresult != -1) {
      ((SendMessageWithParamCallback)g_MessageHandler)(TempMessage);
    }
  }
}

PSYMBOL_LOADED_MODULE_DETAILS SymGetModuleBaseFromSearchMask(
    const char *SearchMask, BOOLEAN SetModuleNameGlobally) {
  string Token;
  char ModuleName[_MAX_FNAME] = {0};
  int Index = 0;
  char Ch = NULL;
  if (!g_IsLoadedModulesInitialized || SearchMask == NULL) {
    return NULL;
  }
  string SearchMaskString(SearchMask);
  char Delimiter = '!';
  if (SearchMaskString.find(Delimiter) != std::string::npos) {
    Token = SearchMaskString.substr(0, SearchMaskString.find(Delimiter));
    strcpy(ModuleName, Token.c_str());
    if (ModuleName[0] == '\0') {
      return NULL;
    }
    while (ModuleName[Index]) {
      Ch = ModuleName[Index];
      ModuleName[Index] = tolower(Ch);
      Index++;
    }
  } else {
    RtlZeroMemory(ModuleName, _MAX_FNAME);
    ModuleName[0] = 'n';
    ModuleName[1] = 't';
  }
  for (auto item : g_LoadedModules) {
    if (strcmp((const char *)item->ModuleName, ModuleName) == 0) {
      if (SetModuleNameGlobally) {
        g_CurrentModuleName = (char *)item->ModuleName;
      }
      return item;
    }
    if (strcmp((const char *)item->ModuleAlternativeName, ModuleName) == 0) {
      if (SetModuleNameGlobally) {
        g_CurrentModuleName = (char *)item->ModuleAlternativeName;
      }
      return item;
    }
  }
  return NULL;
}

BOOLEAN SymGetFieldOffsetFromModule(UINT64 Base, WCHAR *TypeName,
                                    WCHAR *FieldName, UINT32 *FieldOffset) {
  BOOLEAN Found = FALSE;
  const DWORD SizeOfStruct =
      sizeof(SYMBOL_INFOW) + ((MAX_SYM_NAME - 1) * sizeof(wchar_t));
  uint8_t SymbolInfoBuffer[SizeOfStruct];
  auto SymbolInfo = PSYMBOL_INFOW(SymbolInfoBuffer);
  SymbolInfo->SizeOfStruct = sizeof(SYMBOL_INFOW);
  SymbolInfo->MaxNameLen = MAX_SYM_NAME;
  if (!SymGetTypeFromNameW(GetCurrentProcess(), Base, TypeName, SymbolInfo)) {
    return FALSE;
  }
  const ULONG TypeIndex = SymbolInfo->TypeIndex;
  DWORD ChildrenCount = 0;
  if (!SymGetTypeInfo(GetCurrentProcess(), Base, TypeIndex,
                      TI_GET_CHILDRENCOUNT, &ChildrenCount)) {
    return FALSE;
  }
  auto FindChildrenParamsBacking = std::make_unique<uint8_t[]>(
      sizeof(_TI_FINDCHILDREN_PARAMS) + ((ChildrenCount - 1) * sizeof(ULONG)));
  auto FindChildrenParams =
      (_TI_FINDCHILDREN_PARAMS *)FindChildrenParamsBacking.get();
  FindChildrenParams->Count = ChildrenCount;
  if (!SymGetTypeInfo(GetCurrentProcess(), Base, TypeIndex, TI_FINDCHILDREN,
                      FindChildrenParams)) {
    return FALSE;
  }
  for (DWORD ChildIdx = 0; ChildIdx < ChildrenCount; ChildIdx++) {
    const ULONG ChildId = FindChildrenParams->ChildId[ChildIdx];
    WCHAR *ChildName = nullptr;
    SymGetTypeInfo(GetCurrentProcess(), Base, ChildId, TI_GET_SYMNAME,
                   &ChildName);
    UINT64 ChildSize = 0;
    SymGetTypeInfo(GetCurrentProcess(), Base, ChildId, TI_GET_LENGTH,
                   &ChildSize);
    Found = FALSE;
    if (wcscmp(ChildName, FieldName) == 0) {
      const IMAGEHLP_SYMBOL_TYPE_INFO Info =
          (ChildSize == 1) ? TI_GET_BITPOSITION : TI_GET_OFFSET;
      SymGetTypeInfo(GetCurrentProcess(), Base, ChildId, Info, FieldOffset);
      Found = TRUE;
    }
    LocalFree(ChildName);
    ChildName = nullptr;
    if (Found) {
      break;
    }
  }
  return Found;
}

BOOLEAN SymGetDataTypeSizeFromModule(UINT64 Base, WCHAR *TypeName,
                                     UINT64 *TypeSize) {
  const DWORD SizeOfStruct =
      sizeof(SYMBOL_INFOW) + ((MAX_SYM_NAME - 1) * sizeof(wchar_t));
  uint8_t SymbolInfoBuffer[SizeOfStruct];
  auto SymbolInfo = PSYMBOL_INFOW(SymbolInfoBuffer);
  SymbolInfo->SizeOfStruct = sizeof(SYMBOL_INFOW);
  SymbolInfo->MaxNameLen = MAX_SYM_NAME;
  if (!SymGetTypeFromNameW(GetCurrentProcess(), Base, TypeName, SymbolInfo)) {
    return FALSE;
  }
  if (!SymGetTypeInfo(GetCurrentProcess(), Base, SymbolInfo->TypeIndex,
                      TI_GET_LENGTH, TypeSize)) {
    return FALSE;
  }
  return TRUE;
}

BOOLEAN SymInit() {
  BOOL Ret = FALSE;
  DWORD Options = 0;
  Options = SymGetOptions();
  Options |= SYMOPT_DEBUG;
  Options |= SYMOPT_CASE_INSENSITIVE;
  SymSetOptions(Options);
  Ret = SymInitialize(GetCurrentProcess(), NULL, FALSE);
  if (!Ret) {
    ShowMessages("err, symbol init failed (%x)\n", GetLastError());
    return FALSE;
  }
  g_IsLoadedModulesInitialized = TRUE;
  return TRUE;
}

BOOLEAN SymCheckAndRemoveWow64Prefix(const char *ModuleAddress,
                                     const char *PdbFileName,
                                     std::string &CustomModuleName) {
  char ModuleName[_MAX_FNAME] = {0};
  char PdbModuleName[_MAX_FNAME] = {0};
  _splitpath(ModuleAddress, NULL, NULL, ModuleName, NULL);
  _splitpath(PdbFileName, NULL, NULL, PdbModuleName, NULL);
  std::transform(ModuleName, ModuleName + strlen(ModuleName), ModuleName,
                 [](unsigned char c) { return std::tolower(c); });
  std::transform(PdbModuleName, PdbModuleName + strlen(PdbModuleName),
                 PdbModuleName,
                 [](unsigned char c) { return std::tolower(c); });
  if (std::strcmp(ModuleName, PdbModuleName) == 0) {
    return FALSE;
  } else {
    std::string LowerModuleAddress(ModuleAddress);
    std::transform(LowerModuleAddress.begin(), LowerModuleAddress.end(),
                   LowerModuleAddress.begin(),
                   [](unsigned char c) { return std::tolower(c); });
    if (PdbModuleName[0] == 'w' &&
        LowerModuleAddress.find(":\\windows\\system32") != std::string::npos) {
      CustomModuleName = std::string(ModuleName);
      return TRUE;
    } else {
      return FALSE;
    }
  }
}

BOOLEAN SymCheckNtoskrnlPrefix(const char *PdbFileName,
                               std::string &CustomModuleName) {
  char PdbModuleName[_MAX_FNAME] = {0};
  _splitpath(PdbFileName, NULL, NULL, PdbModuleName, NULL);
  std::transform(PdbModuleName, PdbModuleName + strlen(PdbModuleName),
                 PdbModuleName,
                 [](unsigned char c) { return std::tolower(c); });
  if (strcmp(PdbModuleName, ("ntkrnlmp")) == 0 ||
      strcmp(PdbModuleName, ("ntoskrnl")) == 0 ||
      strcmp(PdbModuleName, ("ntkrpamp")) == 0 ||
      strcmp(PdbModuleName, ("ntkrnlpa")) == 0) {
    CustomModuleName = "nt";
    return TRUE;
  }
  return FALSE;
}

UINT32 SymLoadFileSymbol(UINT64 BaseAddress, const char *PdbFileName,
                         const char *CustomModuleName) {
  DWORD FileSize = 0;
  int Index = 0;
  char Ch = NULL;
  char ModuleName[_MAX_FNAME] = {0};
  char AlternateModuleName[_MAX_FNAME] = {0};
  PSYMBOL_LOADED_MODULE_DETAILS ModuleDetails = NULL;
  if (!g_IsLoadedModulesInitialized) {
    SymInit();
  }
  if (!SymGetFileParams(PdbFileName, FileSize)) {
    ShowMessages("err, cannot obtain file parameters (internal error)\n");
    return -1;
  }
  _splitpath(PdbFileName, NULL, NULL, ModuleName, NULL);
  strcpy(AlternateModuleName, ModuleName);
  while (ModuleName[Index]) {
    Ch = ModuleName[Index];
    ModuleName[Index] = tolower(Ch);
    Index++;
  }
  ModuleDetails = (SYMBOL_LOADED_MODULE_DETAILS *)malloc(
      sizeof(SYMBOL_LOADED_MODULE_DETAILS));
  if (ModuleDetails == NULL) {
    ShowMessages("err, allocating buffer for storing symbol details (%x)\n",
                 GetLastError());
    return -1;
  }
  RtlZeroMemory(ModuleDetails, sizeof(SYMBOL_LOADED_MODULE_DETAILS));
  ModuleDetails->ModuleBase = SymLoadModule64(
      GetCurrentProcess(), NULL, PdbFileName, NULL, BaseAddress, FileSize);
  if (ModuleDetails->ModuleBase == NULL) {
    ShowMessages("err, loading symbols failed (%x)\n", GetLastError());
    free(ModuleDetails);
    return -1;
  }
#ifndef DoNotShowDetailedResult ShowMessages("loading symbols for: %s\n", \
                                             PdbFilePath);
  ShowMessages("load address: %I64x\n", ModuleDetails.ModuleBase);
  SymShowSymbolInfo(ModuleDetails.ModuleBase);
#endif
  ModuleDetails->BaseAddress = BaseAddress;
  strcpy((char *)ModuleDetails->ModuleName, ModuleName);
  strcpy((char *)ModuleDetails->PdbFilePath, PdbFileName);
  if (CustomModuleName != NULL) {
    strcpy((char *)ModuleDetails->ModuleAlternativeName, CustomModuleName);
  }
  g_LoadedModules.push_back(ModuleDetails);
  return 0;
}

UINT32 SymUnloadModuleSymbol(char *ModuleName) {
  BOOLEAN OneModuleFound = FALSE;
  BOOL Ret = FALSE;
  UINT32 Index = 0;
  for (auto item : g_LoadedModules) {
    Index++;
    if (strcmp(item->ModuleName, ModuleName) == 0 ||
        strcmp(item->ModuleAlternativeName, ModuleName) == 0) {
      Ret = SymUnloadModule64(GetCurrentProcess(), item->ModuleBase);
      if (!Ret) {
        ShowMessages("err, unload symbol failed (%x)\n", GetLastError());
        return -1;
      }
      OneModuleFound = TRUE;
      free(item);
      break;
    }
  }
  if (!OneModuleFound) {
    return -1;
  }
  std::vector<PSYMBOL_LOADED_MODULE_DETAILS>::iterator it =
      g_LoadedModules.begin();
  std::advance(it, --Index);
  g_LoadedModules.erase(it);
  return 0;
}

UINT32 SymUnloadAllSymbols() {
  BOOL Ret = FALSE;
  BOOLEAN IsAnythingLoaded = FALSE;
  if (!g_IsLoadedModulesInitialized) {
    return 0;
  }
  for (auto item : g_LoadedModules) {
    IsAnythingLoaded = TRUE;
    Ret = SymUnloadModule64(GetCurrentProcess(), item->ModuleBase);
    if (!Ret) {
    }
    free(item);
  }
  if (!IsAnythingLoaded) {
    return 0;
  }
  g_LoadedModules.clear();
  Ret = SymCleanup(GetCurrentProcess());
  if (!Ret) {
    ShowMessages("err, symbol cleanup failed (%x)\n", GetLastError());
    return 0;
  }
  g_IsLoadedModulesInitialized = FALSE;
  return 0;
}

UINT64 SymConvertNameToAddress(const char *FunctionOrVariableName,
                               PBOOLEAN WasFound) {
  BOOLEAN Found = FALSE;
  UINT64 Address = NULL;
  UINT64 Buffer[(sizeof(SYMBOL_INFO) + MAX_SYM_NAME * sizeof(CHAR) +
                 sizeof(UINT64) - 1) /
                sizeof(UINT64)];
  PSYMBOL_INFO Symbol = (PSYMBOL_INFO)Buffer;
  string FinalModuleName;
  string TempName(FunctionOrVariableName);
  string ExtractedModuleName;
  string FunctionName;
  *WasFound = FALSE;
  Symbol->SizeOfStruct = sizeof(SYMBOL_INFO);
  Symbol->MaxNameLen = MAX_SYM_NAME;
  size_t FoundIndex = TempName.find('!');
  if (FoundIndex != std::string::npos) {
    ExtractedModuleName = TempName.substr(0, FoundIndex);
    FunctionName = TempName.substr(FoundIndex + 1);
    std::transform(ExtractedModuleName.begin(), ExtractedModuleName.end(),
                   ExtractedModuleName.begin(),
                   [](unsigned char c) { return std::tolower(c); });
    for (auto item : g_LoadedModules) {
      if (strcmp((const char *)item->ModuleName, ExtractedModuleName.c_str()) ==
          0) {
        string ModuleName(item->ModuleName);
        FinalModuleName = ModuleName + "!" + FunctionName;
        break;
      }
      if (strcmp((const char *)item->ModuleAlternativeName,
                 ExtractedModuleName.c_str()) == 0) {
        string ModuleName(item->ModuleName);
        FinalModuleName = ModuleName + "!" + FunctionName;
        break;
      }
    }
  } else {
    for (auto item : g_LoadedModules) {
      if (strcmp((const char *)item->ModuleAlternativeName, "nt") == 0) {
        string ModuleName(item->ModuleName);
        FinalModuleName = ModuleName + "!" + TempName;
        break;
      }
    }
  }
  if (FinalModuleName.empty()) {
    *WasFound = FALSE;
    return NULL;
  }
  if (SymFromName(GetCurrentProcess(), FinalModuleName.c_str(), Symbol)) {
    Found = TRUE;
    Address = Symbol->Address;
  } else {
    Found = FALSE;
  }
  *WasFound = Found;
  return Address;
}

BOOLEAN SymGetFieldOffset(CHAR *TypeName, CHAR *FieldName,
                          UINT32 *FieldOffset) {
  BOOL Ret = FALSE;
  UINT32 Index = 0;
  PSYMBOL_LOADED_MODULE_DETAILS SymbolInfo = NULL;
  BOOLEAN Result = FALSE;
  SymbolInfo = SymGetModuleBaseFromSearchMask(TypeName, TRUE);
  if (SymbolInfo == NULL) {
    return FALSE;
  }
  while (TypeName[Index] != '\0') {
    if (TypeName[Index] == '!') {
      TypeName = (CHAR *)(TypeName + Index + 1);
      break;
    }
    Index++;
  }
  const size_t TypeNameSize = strlen(TypeName) + 1;
  WCHAR *TypeNameW = (WCHAR *)malloc(sizeof(wchar_t) * TypeNameSize);
  if (TypeNameW == NULL) {
    printf("err, could not allocate buffer");
    return FALSE;
  }
  RtlZeroMemory(TypeNameW, sizeof(wchar_t) * TypeNameSize);
  mbstowcs(TypeNameW, TypeName, TypeNameSize);
  const size_t FieldNameSize = strlen(FieldName) + 1;
  WCHAR *FieldNameW = (WCHAR *)malloc(sizeof(wchar_t) * FieldNameSize);
  if (FieldNameW == NULL) {
    printf("err, could not allocate buffer");
    free(TypeNameW);
    return FALSE;
  }
  RtlZeroMemory(FieldNameW, sizeof(wchar_t) * FieldNameSize);
  mbstowcs(FieldNameW, FieldName, FieldNameSize);
  Result = SymGetFieldOffsetFromModule(SymbolInfo->ModuleBase, TypeNameW,
                                       FieldNameW, FieldOffset);
  free(TypeNameW);
  free(FieldNameW);
  return Result;
}

BOOLEAN SymGetDataTypeSize(CHAR *TypeName, UINT64 *TypeSize) {
  BOOL Ret = FALSE;
  UINT32 Index = 0;
  PSYMBOL_LOADED_MODULE_DETAILS SymbolInfo = NULL;
  BOOLEAN Result = FALSE;
  SymbolInfo = SymGetModuleBaseFromSearchMask(TypeName, TRUE);
  if (SymbolInfo == NULL) {
    return FALSE;
  }
  while (TypeName[Index] != '\0') {
    if (TypeName[Index] == '!') {
      TypeName = (CHAR *)(TypeName + Index + 1);
      break;
    }
    Index++;
  }
  const size_t TypeNameSize = strlen(TypeName) + 1;
  WCHAR *TypeNameW = (WCHAR *)malloc(sizeof(wchar_t) * TypeNameSize);
  if (TypeNameW == NULL) {
    printf("err, could not allocate buffer");
    return FALSE;
  }
  RtlZeroMemory(TypeNameW, sizeof(wchar_t) * TypeNameSize);
  mbstowcs(TypeNameW, TypeName, TypeNameSize);
  Result =
      SymGetDataTypeSizeFromModule(SymbolInfo->ModuleBase, TypeNameW, TypeSize);
  free(TypeNameW);
  return Result;
}

UINT32 SymSearchSymbolForMask(const char *SearchMask) {
  BOOL Ret = FALSE;
  PSYMBOL_LOADED_MODULE_DETAILS SymbolInfo = NULL;
  SymbolInfo = SymGetModuleBaseFromSearchMask(SearchMask, TRUE);
  if (SymbolInfo == NULL) {
    return -1;
  }
  Ret = SymEnumSymbols(GetCurrentProcess(), SymbolInfo->ModuleBase, SearchMask,
                       SymDisplayMaskSymbolsCallback, NULL);
  if (!Ret) {
    ShowMessages("err, symbol enum failed (%x)\n", GetLastError());
  }
  return 0;
}

BOOLEAN SymCreateSymbolTableForDisassembler(void *CallbackFunction) {
  BOOL Ret = FALSE;
  BOOLEAN Result = TRUE;
  g_SymbolMapForDisassembler = (SymbolMapCallback)CallbackFunction;
  for (auto item : g_LoadedModules) {
    g_CurrentModuleName = (char *)item->ModuleName;
    Ret = SymEnumSymbols(GetCurrentProcess(), item->BaseAddress, NULL,
                         SymDeliverDisassemblerSymbolMapCallback, NULL);
    if (!Ret) {
      Result = FALSE;
    }
  }
  return Result;
}

string SymSeparateTo64BitValue(UINT64 Value) {
  ostringstream OstringStream;
  string Temp;
  OstringStream << setw(16) << setfill('0') << hex << Value;
  Temp = OstringStream.str();
  Temp.insert(8, 1, '`');
  return Temp;
}

BOOL SymGetFileParams(const char *FileName, DWORD &FileSize) {
  if (FileName == 0) {
    return FALSE;
  }
  char FileExt[_MAX_EXT] = {0};
  _splitpath(FileName, NULL, NULL, NULL, FileExt);
  if (strcmp(FileExt, (".pdb")) == 0 || strcmp(FileExt, (".PDB")) == 0) {
    if (!SymGetFileSize(FileName, FileSize)) {
      return FALSE;
    }
  } else {
    FileSize = 0;
    return FALSE;
  }
  return TRUE;
}

BOOL SymGetFileSize(const char *FileName, DWORD &FileSize) {
  if (FileName == 0) {
    return FALSE;
  }
  HANDLE hFile = CreateFileA(FileName, GENERIC_READ, FILE_SHARE_READ, NULL,
                             OPEN_EXISTING, 0, NULL);
  if (hFile == INVALID_HANDLE_VALUE) {
    ShowMessages("err, unable to open symbol file (%x)\n", GetLastError());
    return FALSE;
  }
  FileSize = GetFileSize(hFile, NULL);
  if (FileSize == INVALID_FILE_SIZE) {
    ShowMessages("err, unable to get symbol file size (%x)\n", GetLastError());
  }
  if (!CloseHandle(hFile)) {
    ShowMessages("err, unable to close symbol file (%x)\n", GetLastError());
  }
  return (FileSize != INVALID_FILE_SIZE);
}

VOID SymShowSymbolInfo(UINT64 ModuleBase) {
  IMAGEHLP_MODULE64 ModuleInfo;
  memset(&ModuleInfo, 0, sizeof(ModuleInfo));
  ModuleInfo.SizeOfStruct = sizeof(ModuleInfo);
  BOOL Ret = SymGetModuleInfo64(GetCurrentProcess(), ModuleBase, &ModuleInfo);
  if (!Ret) {
    ShowMessages("err, unable to get symbol file information (%x)\n",
                 GetLastError());
    return;
  }
  switch (ModuleInfo.SymType) {
    case SymNone:
      ShowMessages("no symbols available for the module\n");
      break;
    case SymExport:
      ShowMessages("loaded symbols: Exports\n");
      break;
    case SymCoff:
      ShowMessages("loaded symbols: COFF\n");
      break;
    case SymCv:
      ShowMessages("loaded symbols: CodeView\n");
      break;
    case SymSym:
      ShowMessages("loaded symbols: SYM\n");
      break;
    case SymVirtual:
      ShowMessages("loaded symbols: Virtual\n");
      break;
    case SymPdb:
      ShowMessages("loaded symbols: PDB\n");
      break;
    case SymDia:
      ShowMessages("loaded symbols: DIA\n");
      break;
    case SymDeferred:
      ShowMessages("loaded symbols: Deferred\n");
      break;
    default:
      ShowMessages("loaded symbols: Unknown format\n");
      break;
  }
  if (ModuleInfo.ImageName[0] != '\0') {
    ShowMessages("image name: %s\n", ModuleInfo.ImageName);
  }
  if (ModuleInfo.LoadedImageName[0] != '\0') {
    ShowMessages("loaded image name: %s\n", ModuleInfo.LoadedImageName);
  }
  if (ModuleInfo.LoadedPdbName[0] != '\0') {
    ShowMessages("PDB file name: %s\n", ModuleInfo.LoadedPdbName);
  }
  if (ModuleInfo.PdbUnmatched || ModuleInfo.DbgUnmatched) {
    ShowMessages("warning, unmatched symbols\n");
  }
  ShowMessages("line numbers: %s\n",
               ModuleInfo.LineNumbers ? "available" : "not available");
  ShowMessages("global symbols: %s\n",
               ModuleInfo.GlobalSymbols ? "available" : "not available");
  ShowMessages("type information: %s\n",
               ModuleInfo.TypeInfo ? ("Available") : ("Not available"));
  ShowMessages("source indexing: %s\n",
               ModuleInfo.SourceIndexed ? "yes" : "no");
  ShowMessages("public symbols: %s\n",
               ModuleInfo.Publics ? "available" : "not available");
}

BOOL CALLBACK SymDisplayMaskSymbolsCallback(SYMBOL_INFO *SymInfo,
                                            ULONG SymbolSize,
                                            PVOID UserContext) {
  if (SymInfo != 0) {
    SymShowSymbolDetails(*SymInfo);
  }
  return TRUE;
}

BOOL CALLBACK SymDeliverDisassemblerSymbolMapCallback(SYMBOL_INFO *SymInfo,
                                                      ULONG SymbolSize,
                                                      PVOID UserContext) {
  if (SymInfo != 0 && g_SymbolMapForDisassembler != NULL) {
    g_SymbolMapForDisassembler(SymInfo->Address, g_CurrentModuleName,
                               SymInfo->Name, SymInfo->Size);
  }
  return TRUE;
}

VOID SymShowSymbolDetails(SYMBOL_INFO &SymInfo) {
  if (g_CurrentModuleName == NULL) {
    ShowMessages("%s ", SymSeparateTo64BitValue(SymInfo.Address).c_str());
  } else {
    ShowMessages("%s  %s!", SymSeparateTo64BitValue(SymInfo.Address).c_str(),
                 g_CurrentModuleName);
  }
  ShowMessages("%s\n", SymInfo.Name);
#ifndef DoNotShowDetailedResult ShowMessages(" size: %u", SymInfo.Size);
  ShowMessages(" symbol: %s  ", SymTagStr(SymInfo.Tag));
#endif
}

const char *SymTagStr(ULONG Tag) {
  switch (Tag) {
    case SymTagNull:
      return ("Null");
    case SymTagExe:
      return ("Exe");
    case SymTagCompiland:
      return ("Compiland");
    case SymTagCompilandDetails:
      return ("CompilandDetails");
    case SymTagCompilandEnv:
      return ("CompilandEnv");
    case SymTagFunction:
      return ("Function");
    case SymTagBlock:
      return ("Block");
    case SymTagData:
      return ("Data");
    case SymTagAnnotation:
      return ("Annotation");
    case SymTagLabel:
      return ("Label");
    case SymTagPublicSymbol:
      return ("PublicSymbol");
    case SymTagUDT:
      return ("UDT");
    case SymTagEnum:
      return ("Enum");
    case SymTagFunctionType:
      return ("FunctionType");
    case SymTagPointerType:
      return ("PointerType");
    case SymTagArrayType:
      return ("ArrayType");
    case SymTagBaseType:
      return ("BaseType");
    case SymTagTypedef:
      return ("Typedef");
    case SymTagBaseClass:
      return ("BaseClass");
    case SymTagFriend:
      return ("Friend");
    case SymTagFunctionArgType:
      return ("FunctionArgType");
    case SymTagFuncDebugStart:
      return ("FuncDebugStart");
    case SymTagFuncDebugEnd:
      return ("FuncDebugEnd");
    case SymTagUsingNamespace:
      return ("UsingNamespace");
    case SymTagVTableShape:
      return ("VTableShape");
    case SymTagVTable:
      return ("VTable");
    case SymTagCustom:
      return ("Custom");
    case SymTagThunk:
      return ("Thunk");
    case SymTagCustomType:
      return ("CustomType");
    case SymTagManagedType:
      return ("ManagedType");
    case SymTagDimension:
      return ("Dimension");
    default:
      return ("Unknown");
  }
  return ("");
}

BOOLEAN SymConvertFileToPdbPath(const char *LocalFilePath, char *ResultPath,
                                size_t ResultPathSize) {
  HRESULT Result;
  BOOL Ret;
  SYMSRV_INDEX_INFO SymInfo = {0};
  const char *FormatStr =
      "%s/%08x%04x%04x%02x%02x%02x%02x%02x%02x%02x%02x%x/%s";
  SymInfo.sizeofstruct = sizeof(SYMSRV_INDEX_INFO);
  if (LocalFilePath == NULL && ResultPath == NULL) {
    return FALSE;
  }
  Ret = SymSrvGetFileIndexInfo(LocalFilePath, &SymInfo, 0);
  if (Ret) {
    Result = StringCchPrintfA(
        ResultPath, ResultPathSize, FormatStr, SymInfo.pdbfile,
        SymInfo.guid.Data1, SymInfo.guid.Data2, SymInfo.guid.Data3,
        SymInfo.guid.Data4[0], SymInfo.guid.Data4[1], SymInfo.guid.Data4[2],
        SymInfo.guid.Data4[3], SymInfo.guid.Data4[4], SymInfo.guid.Data4[5],
        SymInfo.guid.Data4[6], SymInfo.guid.Data4[7], SymInfo.age,
        SymInfo.pdbfile);
    if (FAILED(Result)) {
      return FALSE;
    }
    return TRUE;
  } else {
    return FALSE;
  }
  return FALSE;
}

VOID SymConvertWow64CompatibilityPaths(const char *LocalFilePath,
                                       std::string &Wow64ConvertedPath) {
  std::string FilePath(LocalFilePath);
  std::transform(FilePath.begin(), FilePath.end(), FilePath.begin(), ::tolower);
  size_t pos = FilePath.find(":\\windows\\system32");
  if (pos != std::string::npos) {
    FilePath.replace(pos, 18, ":\\windows\\syswow64");
  }
  pos = FilePath.find(":\\program files");
  if (pos != std::string::npos) {
    FilePath.replace(pos, 15, ":\\program files (x86)");
  }
  Wow64ConvertedPath = FilePath;
}

BOOLEAN SymConvertFileToPdbFileAndGuidAndAgeDetails(const char *LocalFilePath,
                                                    char *PdbFilePath,
                                                    char *GuidAndAgeDetails,
                                                    BOOLEAN Is32BitModule) {
  SYMSRV_INDEX_INFO SymInfo = {0};
  std::string Wow64ConvertedPath;
  const char *FormatStrPdbFilePath = "%s";
  const char *ActualLocalFilePath = NULL;
  const char *FormatStrPdbFileGuidAndAgeDetails =
      "%08x%04x%04x%02x%02x%02x%02x%02x%02x%02x%02x%x";
  SymInfo.sizeofstruct = sizeof(SYMSRV_INDEX_INFO);
  if (Is32BitModule) {
    SymConvertWow64CompatibilityPaths(LocalFilePath, Wow64ConvertedPath);
    ActualLocalFilePath = Wow64ConvertedPath.c_str();
  } else {
    ActualLocalFilePath = LocalFilePath;
  }
  BOOL Ret = SymSrvGetFileIndexInfo(ActualLocalFilePath, &SymInfo, 0);
  if (Ret) {
    wsprintfA(PdbFilePath, FormatStrPdbFilePath, SymInfo.pdbfile);
    wsprintfA(GuidAndAgeDetails, FormatStrPdbFileGuidAndAgeDetails,
              SymInfo.guid.Data1, SymInfo.guid.Data2, SymInfo.guid.Data3,
              SymInfo.guid.Data4[0], SymInfo.guid.Data4[1],
              SymInfo.guid.Data4[2], SymInfo.guid.Data4[3],
              SymInfo.guid.Data4[4], SymInfo.guid.Data4[5],
              SymInfo.guid.Data4[6], SymInfo.guid.Data4[7], SymInfo.age);
    return TRUE;
  } else {
    return FALSE;
  }
  return FALSE;
}

BOOLEAN SymbolInitLoad(PVOID BufferToStoreDetails, UINT32 StoredLength,
                       BOOLEAN DownloadIfAvailable, const char *SymbolPath,
                       BOOLEAN IsSilentLoad) {
  string Tmp, SymDir, CustomModuleNameStr;
  const char *CustomModuleName = NULL;
  string SymPath(SymbolPath);
  PMODULE_SYMBOL_DETAIL BufferToStoreDetailsConverted =
      (PMODULE_SYMBOL_DETAIL)BufferToStoreDetails;
  vector<string> SplitedSymPath = Split(SymPath, '*');
  if (SplitedSymPath.size() < 2) return FALSE;
  if (SplitedSymPath[1].find(":\\") == string::npos) return FALSE;
  Tmp = SymDir = SplitedSymPath[1];
  for (size_t i = 0; i < StoredLength / sizeof(MODULE_SYMBOL_DETAIL); i++) {
    if (g_AbortLoadingExecution) {
      g_AbortLoadingExecution = FALSE;
      return FALSE;
    }
    if (!BufferToStoreDetailsConverted[i].IsSymbolDetailsFound) {
      continue;
    }
    if (BufferToStoreDetailsConverted[i].IsLocalSymbolPath) {
      if (IsFileExists(BufferToStoreDetailsConverted[i].ModuleSymbolPath)) {
        BufferToStoreDetailsConverted[i].IsSymbolPDBAvaliable = TRUE;
        if (!IsSilentLoad) {
          ShowMessages("loading symbol '%s'...", Tmp.c_str());
        }
        CustomModuleName = NULL;
        if (BufferToStoreDetailsConverted[i].Is32Bit &&
            SymCheckAndRemoveWow64Prefix(
                BufferToStoreDetailsConverted[i].FilePath,
                BufferToStoreDetailsConverted[i].ModuleSymbolPath,
                CustomModuleNameStr)) {
          CustomModuleName = CustomModuleNameStr.c_str();
        } else if (!BufferToStoreDetailsConverted[i].Is32Bit &&
                   SymCheckNtoskrnlPrefix(
                       BufferToStoreDetailsConverted[i].ModuleSymbolPath,
                       CustomModuleNameStr)) {
          CustomModuleName = CustomModuleNameStr.c_str();
        }
        if (SymLoadFileSymbol(BufferToStoreDetailsConverted[i].BaseAddress,
                              BufferToStoreDetailsConverted[i].ModuleSymbolPath,
                              CustomModuleName) == 0) {
          if (!IsSilentLoad) {
            ShowMessages("\tloaded\n");
          }
        } else {
          if (!IsSilentLoad) {
            ShowMessages("\tnot loaded (already loaded?)\n");
          }
        }
      }
    } else {
      Tmp = SymDir + "\\" + BufferToStoreDetailsConverted[i].ModuleSymbolPath +
            "\\" + BufferToStoreDetailsConverted[i].ModuleSymbolGuidAndAge +
            "\\" + BufferToStoreDetailsConverted[i].ModuleSymbolPath;
      if (DownloadIfAvailable && IsFileExists(Tmp) == FALSE) {
        SymbolPdbDownload(
            BufferToStoreDetailsConverted[i].ModuleSymbolPath,
            BufferToStoreDetailsConverted[i].ModuleSymbolGuidAndAge, SymPath,
            IsSilentLoad);
      }
      if (IsFileExists(Tmp)) {
        BufferToStoreDetailsConverted[i].IsSymbolPDBAvaliable = TRUE;
        if (!IsSilentLoad) {
          ShowMessages("loading symbol '%s'...", Tmp.c_str());
        }
        CustomModuleName = NULL;
        if (BufferToStoreDetailsConverted[i].Is32Bit &&
            SymCheckAndRemoveWow64Prefix(
                BufferToStoreDetailsConverted[i].FilePath, Tmp.c_str(),
                CustomModuleNameStr)) {
          CustomModuleName = CustomModuleNameStr.c_str();
        } else if (!BufferToStoreDetailsConverted[i].Is32Bit &&
                   SymCheckNtoskrnlPrefix(Tmp.c_str(), CustomModuleNameStr)) {
          CustomModuleName = CustomModuleNameStr.c_str();
        }
        if (SymLoadFileSymbol(BufferToStoreDetailsConverted[i].BaseAddress,
                              Tmp.c_str(), CustomModuleName) == 0) {
          if (!IsSilentLoad) {
            ShowMessages("\tloaded\n");
          }
        } else {
          if (!IsSilentLoad) {
            ShowMessages("\tnot loaded (already loaded?)\n");
          }
        }
      }
    }
  }
  return TRUE;
}

BOOLEAN SymbolPdbDownload(std::string SymName, const std::string &GUID,
                          const std::string &SymPath, BOOLEAN IsSilentLoad) {
  vector<string> SplitedSymPath = Split(SymPath, '*');
  if (SplitedSymPath.size() < 2) return FALSE;
  if (SplitedSymPath[1].find(":\\") == string::npos) return FALSE;
  if (SplitedSymPath[2].find("http:") == string::npos &&
      SplitedSymPath[2].find("https:") == string::npos)
    return FALSE;
  string SymDir = SplitedSymPath[1];
  string SymDownloadServer = SplitedSymPath[2];
  string DownloadURL =
      SymDownloadServer + "/" + SymName + "/" + GUID + "/" + SymName;
  string SymFullDir = SymDir + "\\" + SymName + "\\" + GUID + "\\";
  if (!CreateDirectoryRecursive(SymFullDir)) {
    if (!IsSilentLoad) {
      ShowMessages("err, unable to create sympath directory '%s'\n",
                   SymFullDir.c_str());
    }
    return FALSE;
  }
  if (!IsSilentLoad) {
    ShowMessages("downloading symbol '%s'...", SymName.c_str());
  }
  HRESULT Result =
      URLDownloadToFileA(NULL, DownloadURL.c_str(),
                         (SymFullDir + "\\" + SymName).c_str(), 0, NULL);
  if (Result == S_OK) {
    if (!IsSilentLoad) {
      ShowMessages("\tdownloaded\n");
    }
    return TRUE;
  } else {
    if (!IsSilentLoad) {
      ShowMessages("\tcould not be downloaded (%x) \n", Result);
    }
  }
  return FALSE;
}

VOID SymbolAbortLoading() {
  if (!g_AbortLoadingExecution) {
    g_AbortLoadingExecution = TRUE;
    ShowMessages("\naborting, please wait...\n");
  }
}

BOOLEAN SymShowDataBasedOnSymbolTypes(const char *TypeName, UINT64 Address,
                                      BOOLEAN IsStruct, PVOID BufferAddress,
                                      const char *AdditionalParameters) {
  vector<string> SplitedSymPath;
  char **ArgvArray = NULL;
  PSYMBOL_LOADED_MODULE_DETAILS SymbolInfo = NULL;
  UINT32 SizeOfArgv = 0;
  UINT32 TypeNameIndex = 0;
  SymbolInfo = SymGetModuleBaseFromSearchMask(TypeName, FALSE);
  if (!SymbolInfo) {
    ShowMessages("err, couldn't resolve error at '%s'\n", TypeName);
    return FALSE;
  }
  std::string AdditionalParametersString(AdditionalParameters);
  SplitPathAndArgs(SplitedSymPath, AdditionalParametersString);
  SizeOfArgv = (UINT32)SplitedSymPath.size() + 3;
  ArgvArray = (char **)malloc(SizeOfArgv * sizeof(char *));
  if (ArgvArray == NULL) {
    return FALSE;
  }
  RtlZeroMemory(ArgvArray, SizeOfArgv * sizeof(char *));
  ArgvArray[0] = (char *)NULL;
  while (TypeName[TypeNameIndex] != NULL) {
    if (TypeName[TypeNameIndex] == '!') {
      TypeName = &TypeName[++TypeNameIndex];
      break;
    }
    TypeNameIndex++;
  }
  ArgvArray[1] = (char *)TypeName;
  ArgvArray[2] = SymbolInfo->PdbFilePath;
  for (size_t i = 3; i < SizeOfArgv; i++) {
    ArgvArray[i] = (char *)SplitedSymPath.at(i - 3).c_str();
  }
  if (IsStruct) {
    pdbex_export(SizeOfArgv, ArgvArray, true, BufferAddress);
  } else {
    pdbex_export(SizeOfArgv, ArgvArray, false, BufferAddress);
  }
  free(ArgvArray);
  return TRUE;
}
