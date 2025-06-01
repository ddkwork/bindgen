
#pragma once

typedef struct _MODULE_SYMBOL_DETAIL {
  BOOLEAN IsSymbolDetailsFound;
  BOOLEAN IsLocalSymbolPath;
  BOOLEAN IsSymbolPDBAvaliable;
  BOOLEAN IsUserMode;
  BOOLEAN Is32Bit;
  UINT64 BaseAddress;
  char FilePath[MAX_PATH];
  char ModuleSymbolPath[MAX_PATH];
  char ModuleSymbolGuidAndAge[MAXIMUM_GUID_AND_AGE_SIZE];
} MODULE_SYMBOL_DETAIL, *PMODULE_SYMBOL_DETAIL;

typedef struct _USERMODE_LOADED_MODULE_SYMBOLS {
  UINT64 BaseAddress;
  UINT64 Entrypoint;
  wchar_t FilePath[MAX_PATH];
} USERMODE_LOADED_MODULE_SYMBOLS, *PUSERMODE_LOADED_MODULE_SYMBOLS;

typedef struct _USERMODE_LOADED_MODULE_DETAILS {
  UINT32 ProcessId;
  BOOLEAN OnlyCountModules;
  BOOLEAN Is32Bit;
  UINT32 ModulesCount;
  UINT32 Result;
} USERMODE_LOADED_MODULE_DETAILS, *PUSERMODE_LOADED_MODULE_DETAILS;

typedef VOID (*SymbolMapCallback)(UINT64 Address, char *ModuleName,
                                  char *ObjectName, unsigned int ObjectSize);

typedef struct _DEBUGGER_UPDATE_SYMBOL_TABLE {
  UINT32 TotalSymbols;
  UINT32 CurrentSymbolIndex;
  MODULE_SYMBOL_DETAIL SymbolDetailPacket;
} DEBUGGER_UPDATE_SYMBOL_TABLE, *PDEBUGGER_UPDATE_SYMBOL_TABLE;

typedef struct _DEBUGGEE_SYMBOL_UPDATE_RESULT {
  UINT64 KernelStatus;
} DEBUGGEE_SYMBOL_UPDATE_RESULT, *PDEBUGGEE_SYMBOL_UPDATE_RESULT;
