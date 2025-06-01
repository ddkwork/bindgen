
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
  Dpc = PlatformMemAllocateNonPagedPool(sizeof(KDPC));
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

BOOLEAN DpcRoutinePerformVirtualization(KDPC *Dpc, PVOID DeferredContext,
                                        PVOID SystemArgument1,
                                        PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  VmxPerformVirtualizationOnSpecificCore();
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
  return TRUE;
}

VOID DpcRoutinePerformChangeMsrBitmapReadOnSingleCore(KDPC *Dpc,
                                                      PVOID DeferredContext,
                                                      PVOID SystemArgument1,
                                                      PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_READ, (UINT64)DeferredContext, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore(KDPC *Dpc,
                                                       PVOID DeferredContext,
                                                       PVOID SystemArgument1,
                                                       PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_WRITE, (UINT64)DeferredContext, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformEnableRdtscExitingOnSingleCore(KDPC *Dpc,
                                                     PVOID DeferredContext,
                                                     PVOID SystemArgument1,
                                                     PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_SET_RDTSC_EXITING, 0, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformEnableRdpmcExitingOnSingleCore(KDPC *Dpc,
                                                     PVOID DeferredContext,
                                                     PVOID SystemArgument1,
                                                     PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_SET_RDPMC_EXITING, 0, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformSetExceptionBitmapOnSingleCore(KDPC *Dpc,
                                                     PVOID DeferredContext,
                                                     PVOID SystemArgument1,
                                                     PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, (UINT64)DeferredContext, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformEnableMovToDebugRegistersExiting(KDPC *Dpc,
                                                       PVOID DeferredContext,
                                                       PVOID SystemArgument1,
                                                       PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformEnableMovToControlRegisterExiting(
    KDPC *Dpc, DEBUGGER_EVENT_OPTIONS *EventOptions, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING,
               EventOptions->OptionalParam1, EventOptions->OptionalParam2, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformSetExternalInterruptExitingOnSingleCore(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING, NULL64_ZERO, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformEnableEferSyscallHookOnSingleCore(KDPC *Dpc,
                                                        PVOID DeferredContext,
                                                        PVOID SystemArgument1,
                                                        PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_ENABLE_SYSCALL_HOOK_EFER, NULL64_ZERO, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutinePerformChangeIoBitmapOnSingleCore(KDPC *Dpc,
                                                 PVOID DeferredContext,
                                                 PVOID SystemArgument1,
                                                 PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  AsmVmxVmcall(VMCALL_CHANGE_IO_BITMAP, (UINT64)DeferredContext, 0, 0);
  SpinlockUnlock(&OneCoreLock);
}

VOID DpcRoutineEnableMovToCr3Exiting(KDPC *Dpc, PVOID DeferredContext,
                                     PVOID SystemArgument1,
                                     PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CR3_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineChangeToMbecSupportedEptp(KDPC *Dpc, PVOID DeferredContext,
                                         PVOID SystemArgument1,
                                         PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_CHANGE_TO_MBEC_SUPPORTED_EPTP, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineRestoreToNormalEptp(KDPC *Dpc, PVOID DeferredContext,
                                   PVOID SystemArgument1,
                                   PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_RESTORE_TO_NORMAL_EPTP, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableOrDisableMbec(KDPC *Dpc, PVOID DeferredContext,
                                   PVOID SystemArgument1,
                                   PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_DISABLE_OR_ENABLE_MBEC, (UINT64)DeferredContext, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableMovToCr3Exiting(KDPC *Dpc, PVOID DeferredContext,
                                      PVOID SystemArgument1,
                                      PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CR3_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableEferSyscallEvents(KDPC *Dpc, PVOID DeferredContext,
                                       PVOID SystemArgument1,
                                       PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_ENABLE_SYSCALL_HOOK_EFER, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableEferSyscallEvents(KDPC *Dpc, PVOID DeferredContext,
                                        PVOID SystemArgument1,
                                        PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_DISABLE_SYSCALL_HOOK_EFER, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnablePml(KDPC *Dpc, PVOID DeferredContext,
                         PVOID SystemArgument1, PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_ENABLE_DIRTY_LOGGING_MECHANISM, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisablePml(KDPC *Dpc, PVOID DeferredContext,
                          PVOID SystemArgument1, PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_DISABLE_DIRTY_LOGGING_MECHANISM, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineChangeMsrBitmapReadOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                             PVOID SystemArgument1,
                                             PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_READ, (UINT64)DeferredContext, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineResetMsrBitmapReadOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                            PVOID SystemArgument1,
                                            PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_RESET_MSR_BITMAP_READ, NULL64_ZERO, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineChangeMsrBitmapWriteOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                              PVOID SystemArgument1,
                                              PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_WRITE, (UINT64)DeferredContext, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineResetMsrBitmapWriteOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                             PVOID SystemArgument1,
                                             PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_RESET_MSR_BITMAP_WRITE, NULL64_ZERO, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableRdtscExitingAllCores(KDPC *Dpc, PVOID DeferredContext,
                                          PVOID SystemArgument1,
                                          PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_SET_RDTSC_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableRdtscExitingAllCores(KDPC *Dpc, PVOID DeferredContext,
                                           PVOID SystemArgument1,
                                           PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_UNSET_RDTSC_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores(
    KDPC *Dpc, DEBUGGER_EVENT_OPTIONS *EventOptions, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS,
               EventOptions->OptionalParam1, EventOptions->OptionalParam2, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableRdpmcExitingAllCores(KDPC *Dpc, PVOID DeferredContext,
                                          PVOID SystemArgument1,
                                          PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_SET_RDPMC_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableRdpmcExitingAllCores(KDPC *Dpc, PVOID DeferredContext,
                                           PVOID SystemArgument1,
                                           PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_UNSET_RDPMC_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineSetExceptionBitmapOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                            PVOID SystemArgument1,
                                            PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, (UINT64)DeferredContext, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineUnsetExceptionBitmapOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                              PVOID SystemArgument1,
                                              PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, (UINT64)DeferredContext, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS,
               NULL64_ZERO, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableMovDebigRegisterExitingAllCores(KDPC *Dpc,
                                                     PVOID DeferredContext,
                                                     PVOID SystemArgument1,
                                                     PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableMovControlRegisterExitingAllCores(
    KDPC *Dpc, DEBUGGER_EVENT_OPTIONS *EventOptions, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING,
               EventOptions->OptionalParam1, EventOptions->OptionalParam2, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableMovControlRegisterExitingAllCores(
    KDPC *Dpc, DEBUGGER_EVENT_OPTIONS *EventOptions, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CONTROL_REGS_EXITING,
               EventOptions->OptionalParam1, EventOptions->OptionalParam2, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableMovDebigRegisterExitingAllCores(KDPC *Dpc,
                                                      PVOID DeferredContext,
                                                      PVOID SystemArgument1,
                                                      PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineSetEnableExternalInterruptExitingOnAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING, 0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(
      VMCALL_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS,
      0, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineChangeIoBitmapOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                        PVOID SystemArgument1,
                                        PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  AsmVmxVmcall(VMCALL_CHANGE_IO_BITMAP, (UINT64)DeferredContext, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineResetIoBitmapOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                       PVOID SystemArgument1,
                                       PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_RESET_IO_BITMAP, NULL64_ZERO, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0,
               0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableNmiVmexitOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                         PVOID SystemArgument1,
                                         PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_SET_VM_EXIT_ON_NMIS, NULL64_ZERO, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableNmiVmexitOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                          PVOID SystemArgument1,
                                          PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_UNSET_VM_EXIT_ON_NMIS, NULL64_ZERO, 0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineEnableDbAndBpExitingOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                              PVOID SystemArgument1,
                                              PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
  AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_DEBUG_BREAKPOINT,
               0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineDisableDbAndBpExitingOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                               PVOID SystemArgument1,
                                               PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0,
               0);
  AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_DEBUG_BREAKPOINT,
               0, 0);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxVmcall(VMCALL_UNHOOK_ALL_PAGES, NULL64_ZERO, NULL64_ZERO, NULL64_ZERO);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores(
    KDPC *Dpc, PVOID DeferredContext, PVOID SystemArgument1,
    PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  EPT_SINGLE_HOOK_UNHOOKING_DETAILS *UnhookingDetail =
      (EPT_SINGLE_HOOK_UNHOOKING_DETAILS *)DeferredContext;
  AsmVmxVmcall(VMCALL_UNHOOK_SINGLE_PAGE, UnhookingDetail->PhysicalAddress,
               UnhookingDetail->OriginalEntry, NULL64_ZERO);
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineInvalidateEptOnAllCores(KDPC *Dpc, PVOID DeferredContext,
                                       PVOID SystemArgument1,
                                       PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  if (DeferredContext == NULL) {
    AsmVmxVmcall(VMCALL_INVEPT_ALL_CONTEXTS, NULL64_ZERO, NULL64_ZERO,
                 NULL64_ZERO);
  } else {
    AsmVmxVmcall(
        VMCALL_INVEPT_SINGLE_CONTEXT,
        g_GuestState[KeGetCurrentProcessorNumberEx(NULL)].EptPointer.AsUInt,
        NULL64_ZERO, NULL64_ZERO);
  }
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineInitializeGuest(KDPC *Dpc, PVOID DeferredContext,
                               PVOID SystemArgument1, PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  AsmVmxSaveState();
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}

VOID DpcRoutineTerminateGuest(KDPC *Dpc, PVOID DeferredContext,
                              PVOID SystemArgument1, PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(DeferredContext);
  if (!VmxTerminate()) {
    LogError("Err, there were an error terminating vmx");
  }
  KeSignalCallDpcSynchronize(SystemArgument2);
  KeSignalCallDpcDone(SystemArgument1);
}
