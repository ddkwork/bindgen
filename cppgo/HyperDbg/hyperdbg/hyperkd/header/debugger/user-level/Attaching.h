
#pragma once

#define MAX_USER_ACTIONS_FOR_THREADS 3

#define MAX_THREADS_IN_A_PROCESS_HOLDER 100

#define MAX_CR3_IN_A_PROCESS 4

typedef struct _USERMODE_DEBUGGING_PROCESS_DETAILS {
  UINT64 Token;
  BOOLEAN Enabled;
  PVOID PebAddressToMonitor;
  UINT32 ActiveThreadId;
  GUEST_REGS Registers;
  UINT64 Context;
  LIST_ENTRY AttachedProcessList;
  UINT64 UsermodeReservedBuffer;
  UINT64 EntrypointOfMainModule;
  UINT64 BaseAddressOfMainModule;
  PEPROCESS Eprocess;
  UINT32 ProcessId;
  BOOLEAN Is32Bit;
  BOOLEAN IsOnTheStartingPhase;
  BOOLEAN IsOnThreadInterceptingPhase;
  BOOLEAN CheckCallBackForInterceptingFirstInstruction;
  CR3_TYPE InterceptedCr3[MAX_CR3_IN_A_PROCESS];
  LIST_ENTRY ThreadsListHead;
} USERMODE_DEBUGGING_PROCESS_DETAILS, *PUSERMODE_DEBUGGING_PROCESS_DETAILS;

BOOLEAN AttachingInitialize();
BOOLEAN AttachingCheckPageFaultsWithUserDebugger(UINT32 CoreId, UINT64 Address,
                                                 UINT32 PageFaultErrorCode);
BOOLEAN AttachingConfigureInterceptingThreads(UINT64 ProcessDebuggingToken,
                                              BOOLEAN Enable);
BOOLEAN AttachingHandleCr3VmexitsForThreadInterception(UINT32 CoreId,
                                                       CR3_TYPE NewCr3);
VOID AttachingTargetProcess(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS Request);
VOID AttachingHandleEntrypointInterception(PROCESSOR_DEBUGGING_STATE *DbgState);
VOID AttachingRemoveAndFreeAllProcessDebuggingDetails();
PUSERMODE_DEBUGGING_PROCESS_DETAILS AttachingFindProcessDebuggingDetailsByToken(
    UINT64 Token);
PUSERMODE_DEBUGGING_PROCESS_DETAILS
AttachingFindProcessDebuggingDetailsByProcessId(UINT32 ProcessId);
BOOLEAN AttachingQueryDetailsOfActiveDebuggingThreadsAndProcesses(
    PVOID BufferToStoreDetails, UINT32 BufferSize);
BOOLEAN AttachingCheckUnhandledEptViolation(UINT32 CoreId,
                                            UINT64 ViolationQualification,
                                            UINT64 GuestPhysicalAddr);
BOOLEAN AttachingReachedToValidLoadedModule(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetail);
