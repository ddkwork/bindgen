
#include "pch.h"

volatile LONG OneCoreLock;

NTSTATUS DpcRoutineRunTaskOnSingleCore(UINT32 CoreNumber, PVOID Routine,
                                       PVOID DeferredContext) {
  PRKDPC Dpc;
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (CoreNumber >= ProcessorsCount) {
    return STATUS_INVALID_PARAMETER;
  }
  Dpc = PlatformMemAllocateZeroedNonPagedPool(sizeof(KDPC));
  if (!Dpc) {
    return STATUS_INSUFFICIENT_RESOURCES;
  }
  KeInitializeDpc(Dpc, (PKDEFERRED_ROUTINE)Routine, DeferredContext);
  KeSetTargetProcessorDpc(Dpc, (CCHAR)CoreNumber);
  if (!SpinlockTryLock(&OneCoreLock)) {
    PlatformMemFreePool(Dpc);
    return STATUS_UNSUCCESSFUL;
  }
  KeInsertQueueDpc(Dpc, NULL, NULL);
  SpinlockLock(&OneCoreLock);
  SpinlockUnlock(&OneCoreLock);
  PlatformMemFreePool(Dpc);
  return STATUS_SUCCESS;
}

VOID DpcRoutinePerformWriteMsr(KDPC *Dpc, PVOID DeferredContext,
                               PVOID SystemArgument1, PVOID SystemArgument2) {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  PROCESSOR_DEBUGGING_STATE *CurrentDebuggingState = &g_DbgState[CurrentCore];
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  __writemsr(CurrentDebuggingState->MsrState.Msr,
             CurrentDebuggingState->MsrState.Value);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformReadMsr(KDPC *Dpc, PVOID DeferredContext,
                              PVOID SystemArgument1, PVOID SystemArgument2) {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  PROCESSOR_DEBUGGING_STATE *CurrentDebuggingState = &g_DbgState[CurrentCore];
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  CurrentDebuggingState->MsrState.Value =
      __readmsr(CurrentDebuggingState->MsrState.Msr);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutineWriteMsrToAllCores(KDPC *Dpc, PVOID DeferredContext,
                                  PVOID SystemArgument1,
                                  PVOID SystemArgument2) {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  PROCESSOR_DEBUGGING_STATE *CurrentDebuggingState = &g_DbgState[CurrentCore];
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  __writemsr(CurrentDebuggingState->MsrState.Msr,
             CurrentDebuggingState->MsrState.Value);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineReadMsrToAllCores(KDPC *Dpc, PVOID DeferredContext,
                                 PVOID SystemArgument1, PVOID SystemArgument2) {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  PROCESSOR_DEBUGGING_STATE *CurrentDebuggingState = &g_DbgState[CurrentCore];
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  CurrentDebuggingState->MsrState.Value =
      __readmsr(CurrentDebuggingState->MsrState.Msr);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineVmExitAndHaltSystemAllCores(KDPC *Dpc, PVOID DeferredContext,
                                           PVOID SystemArgument1,
                                           PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  VmFuncVmxVmcall(DEBUGGER_VMCALL_VM_EXIT_HALT_SYSTEM, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}
