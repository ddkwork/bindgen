
#include "pch.h"

UINT64 ScriptEnginePseudoRegGetTid() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return (UINT64)GetCurrentThreadId();
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return (UINT64)PsGetCurrentThreadId();
#endif
}

UINT64 ScriptEnginePseudoRegGetCore() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return (UINT64)GetCurrentProcessorNumber();
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return (UINT64)KeGetCurrentProcessorNumberEx(NULL);
#endif
}

UINT64 ScriptEnginePseudoRegGetPid() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return (UINT64)GetCurrentProcessId();
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return (UINT64)PsGetCurrentProcessId();
#endif
}

CHAR *ScriptEnginePseudoRegGetPname() {
#ifdef SCRIPT_ENGINE_USER_MODE
  HANDLE Handle = OpenProcess(PROCESS_QUERY_INFORMATION | PROCESS_VM_READ,
                              FALSE, GetCurrentProcessId());
  if (Handle) {
    CHAR CurrentModulePath[MAX_PATH] = {0};
    if (GetModuleFileNameEx(Handle, 0, CurrentModulePath, MAX_PATH)) {
      CloseHandle(Handle);
      return PathFindFileNameA(CurrentModulePath);
    } else {
      CloseHandle(Handle);
      return NULL;
    }
  }
  return NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return CommonGetProcessNameFromProcessControlBlock(PsGetCurrentProcess());
#endif
}

UINT64 ScriptEnginePseudoRegGetProc() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return (UINT64)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return (UINT64)PsGetCurrentProcess();
#endif
}

UINT64 ScriptEnginePseudoRegGetThread() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return (UINT64)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return (UINT64)PsGetCurrentThread();
#endif
}

UINT64 ScriptEnginePseudoRegGetPeb() {
#ifdef SCRIPT_ENGINE_USER_MODE
  struct PROCESS_BASIC_INFORMATION {
    PVOID Reserved1;
    PVOID PebBaseAddress;
    PVOID Reserved2[2];
    ULONG_PTR UniqueProcessId;
    PVOID Reserved3;
  };

  struct PEB_LDR_DATA {
    BYTE Reserved1[8];
    PVOID Reserved2[3];
    LIST_ENTRY InMemoryOrderModuleList;
  };

  struct PEB {
    BYTE Reserved1[2];
    BYTE BeingDebugged;
    BYTE Reserved2[1];
    PVOID Reserved3[2];
    struct PEB_LDR_DATA *Ldr;
    PVOID ProcessParameters;
    BYTE Reserved4[104];
    PVOID Reserved5[52];
    PVOID PostProcessInitRoutine;
    BYTE Reserved6[128];
    PVOID Reserved7[1];
    ULONG SessionId;
  };

  struct UNICODE_STRING {
    USHORT Length;
    USHORT MaximumLength;
    PWSTR Buffer;
  };

  struct LDR_MODULE {
    LIST_ENTRY InLoadOrderModuleList;
    LIST_ENTRY InMemoryOrderModuleList;
    LIST_ENTRY InInitializationOrderModuleList;
    PVOID BaseAddress;
    PVOID EntryPoint;
    ULONG SizeOfImage;
    struct UNICODE_STRING FullDllName;
    struct UNICODE_STRING BaseDllName;
    ULONG Flags;
    SHORT LoadCount;
    SHORT TlsIndex;
    LIST_ENTRY HashTableEntry;
    ULONG TimeDateStamp;
  };

  enum PROCESSINFOCLASS {
    ProcessBasicInformation = 0,
    ProcessDebugPort = 7,
    ProcessWow64Information = 26,
    ProcessImageFileName = 27
  };

  LPCWSTR NTDLL_NAME = L"ntdll.dll";
  LPCSTR NTQUERYINFO_NAME = "NtQueryInformationProcess";
  HMODULE NtdllMod;
  HANDLE ThisProcess;
  NTSTATUS NtCallRet;
  ULONG BytesReturned;
  NTSTATUS(WINAPI * QueryInfoProcPtr)
  (HANDLE, enum PROCESSINFOCLASS, PVOID, ULONG, PULONG);
  struct PROCESS_BASIC_INFORMATION BasicInfo;
  struct PEB *PebPtr;

  ThisProcess = GetCurrentProcess();
  NtdllMod = LoadLibraryW(NTDLL_NAME);
  QueryInfoProcPtr =
      (NTSTATUS(WINAPI *)(HANDLE, enum PROCESSINFOCLASS, PVOID, ULONG,
                          PULONG))GetProcAddress(NtdllMod, NTQUERYINFO_NAME);
  NtCallRet = QueryInfoProcPtr(ThisProcess, ProcessBasicInformation, &BasicInfo,
                               sizeof(BasicInfo), &BytesReturned);
  PebPtr = (struct PEB *)BasicInfo.PebBaseAddress;
  return (UINT64)PebPtr;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return NULL64_ZERO;
#endif
}

UINT64 ScriptEnginePseudoRegGetTeb() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return (UINT64)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return (UINT64)PsGetCurrentThreadTeb();
#endif
}

UINT64 ScriptEnginePseudoRegGetIp() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return ScriptEngineWrapperGetInstructionPointer();
#endif
}

UINT64 ScriptEnginePseudoRegGetBuffer(UINT64 *CorrespondingAction) {
#ifdef SCRIPT_ENGINE_USER_MODE
  return (UINT64)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return ScriptEngineWrapperGetAddressOfReservedBuffer(
      (DEBUGGER_EVENT_ACTION *)CorrespondingAction);
#endif
}

UINT64 ScriptEnginePseudoRegGetEventTag(PACTION_BUFFER ActionBuffer) {
#ifdef SCRIPT_ENGINE_USER_MODE
  return (UINT64)NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return ActionBuffer->Tag;
#endif
}

UINT64 ScriptEnginePseudoRegGetEventId(PACTION_BUFFER ActionBuffer) {
#ifdef SCRIPT_ENGINE_USER_MODE
  return NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE return (ActionBuffer->Tag - \
                                         DebuggerEventTagStartSeed);
#endif
}

UINT64 ScriptEnginePseudoRegGetEventStage(PACTION_BUFFER ActionBuffer) {
#ifdef SCRIPT_ENGINE_USER_MODE
  return NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return ActionBuffer->CallingStage;
#endif
}

UINT64 ScriptEnginePseudoRegGetTime() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return ScriptEngineGetTargetCoreTime();
#endif
}

UINT64 ScriptEnginePseudoRegGetDate() {
#ifdef SCRIPT_ENGINE_USER_MODE
  return NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return ScriptEngineGetTargetCoreDate();
#endif
}
