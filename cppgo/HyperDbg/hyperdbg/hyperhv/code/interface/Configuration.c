
#include "pch.h"

VOID ConfigureEnableMovToCr3ExitingOnAllProcessors() {
  g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger = TRUE;
  BroadcastEnableMovToCr3ExitingOnAllProcessors();
}

BOOLEAN ConfigureInitializeExecTrapOnAllProcessors() {
  return ExecTrapInitialize();
}

VOID ConfigureUninitializeExecTrapOnAllProcessors() { ExecTrapUninitialize(); }

BOOLEAN ConfigureExecTrapAddProcessToWatchingList(UINT32 ProcessId) {
  return ExecTrapAddProcessToWatchingList(ProcessId);
}

BOOLEAN ConfigureExecTrapRemoveProcessFromWatchingList(UINT32 ProcessId) {
  return ExecTrapRemoveProcessFromWatchingList(ProcessId);
}

VOID ConfigureModeBasedExecHookUninitializeOnAllProcessors() {
  ModeBasedExecHookUninitialize();
}

VOID ConfigureDirtyLoggingInitializeOnAllProcessors() {
  DirtyLoggingInitialize();
}

VOID ConfigureDirtyLoggingUninitializeOnAllProcessors() {
  DirtyLoggingUninitialize();
}

VOID ConfigureDisableMovToCr3ExitingOnAllProcessors() {
  g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger = FALSE;
  BroadcastDisableMovToCr3ExitingOnAllProcessors();
}

VOID ConfigureEnableEferSyscallEventsOnAllProcessors() {
  BroadcastEnableEferSyscallEventsOnAllProcessors();
}

VOID ConfigureDisableEferSyscallEventsOnAllProcessors() {
  BroadcastDisableEferSyscallEventsOnAllProcessors();
}

BOOLEAN ConfigureEptHookUnHookAllByHookingTag(UINT64 HookingTag) {
  return EptHookUnHookAllByHookingTag(HookingTag);
}

BOOLEAN ConfigureEptHookUnHookSingleHookByHookingTagFromVmxRoot(
    UINT64 HookingTag,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails) {
  return EptHookUnHookSingleHookByHookingTagFromVmxRoot(HookingTag,
                                                        TargetUnhookingDetails);
}

BOOLEAN ConfigureEptHookUnHookSingleAddress(UINT64 VirtualAddress,
                                            UINT64 PhysAddress,
                                            UINT32 ProcessId) {
  return EptHookUnHookSingleAddress(VirtualAddress, PhysAddress, ProcessId);
}

BOOLEAN ConfigureEptHookUnHookSingleAddressFromVmxRoot(
    UINT64 VirtualAddress, UINT64 PhysAddress,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails) {
  return EptHookUnHookSingleAddressFromVmxRoot(VirtualAddress, PhysAddress,
                                               TargetUnhookingDetails);
}

VOID ConfigureEptHookAllocateExtraHookingPagesForMemoryMonitorsAndExecEptHooks(
    UINT32 Count) {
  EptHookAllocateExtraHookingPagesForMemoryMonitorsAndExecEptHooks(Count);
}

VOID ConfigureEptHookReservePreallocatedPoolsForEptHooks(UINT32 Count) {
  EptHookReservePreallocatedPoolsForEptHooks(Count);
}

BOOLEAN ConfigureEptHook(PVOID TargetAddress, UINT32 ProcessId) {
  return EptHook(TargetAddress, ProcessId);
}

BOOLEAN ConfigureEptHookFromVmxRoot(PVOID TargetAddress) {
  return EptHookFromVmxRoot(TargetAddress);
}

BOOLEAN ConfigureEptHook2(UINT32 CoreId, PVOID TargetAddress,
                          PVOID HookFunction, UINT32 ProcessId) {
  return EptHookInlineHook(&g_GuestState[CoreId], TargetAddress, HookFunction,
                           ProcessId);
}

BOOLEAN ConfigureEptHookMonitor(
    UINT32 CoreId, EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *HookingDetails,
    UINT32 ProcessId) {
  return EptHookMonitorHook(&g_GuestState[CoreId], HookingDetails, ProcessId);
}

BOOLEAN ConfigureEptHook2FromVmxRoot(UINT32 CoreId, PVOID TargetAddress,
                                     PVOID HookFunction) {
  return EptHookInlineHookFromVmxRoot(&g_GuestState[CoreId], TargetAddress,
                                      HookFunction);
}

BOOLEAN ConfigureEptHookMonitorFromVmxRoot(
    UINT32 CoreId,
    EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *MemoryAddressDetails) {
  return EptHookMonitorFromVmxRoot(&g_GuestState[CoreId], MemoryAddressDetails);
}

BOOLEAN ConfigureEptHookModifyInstructionFetchState(UINT32 CoreId,
                                                    PVOID PhysicalAddress,
                                                    BOOLEAN IsUnset) {
  return EptHookModifyInstructionFetchState(&g_GuestState[CoreId],
                                            PhysicalAddress, IsUnset);
}

BOOLEAN ConfigureEptHookModifyPageReadState(UINT32 CoreId,
                                            PVOID PhysicalAddress,
                                            BOOLEAN IsUnset) {
  return EptHookModifyPageReadState(&g_GuestState[CoreId], PhysicalAddress,
                                    IsUnset);
}

BOOLEAN ConfigureEptHookModifyPageWriteState(UINT32 CoreId,
                                             PVOID PhysicalAddress,
                                             BOOLEAN IsUnset) {
  return EptHookModifyPageWriteState(&g_GuestState[CoreId], PhysicalAddress,
                                     IsUnset);
}

VOID ConfigureEnableEferSyscallHookOnSingleCore(UINT32 TargetCoreId) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformEnableEferSyscallHookOnSingleCore,
      NULL);
}

VOID ConfigureSetEferSyscallOrSysretHookType(
    DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE SyscallHookType) {
  if (SyscallHookType == DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD) {
    g_IsUnsafeSyscallOrSysretHandling = TRUE;
  } else if (SyscallHookType ==
             DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY) {
    g_IsUnsafeSyscallOrSysretHandling = FALSE;
  }
}

VOID ConfigureSetExternalInterruptExitingOnSingleCore(UINT32 TargetCoreId) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId,
      (PVOID)DpcRoutinePerformSetExternalInterruptExitingOnSingleCore, NULL);
}

VOID ConfigureEnableRdtscExitingOnSingleCore(UINT32 TargetCoreId) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformEnableRdtscExitingOnSingleCore,
      NULL);
}

VOID ConfigureEnableRdpmcExitingOnSingleCore(UINT32 TargetCoreId) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformEnableRdpmcExitingOnSingleCore,
      NULL);
}

VOID ConfigureEnableMovToDebugRegistersExitingOnSingleCore(
    UINT32 TargetCoreId) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformEnableMovToDebugRegistersExiting,
      NULL);
}

VOID ConfigureSetExceptionBitmapOnSingleCore(UINT32 TargetCoreId,
                                             UINT32 BitMask) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformSetExceptionBitmapOnSingleCore,
      (PVOID)BitMask);
}

VOID ConfigureEnableMovToControlRegisterExitingOnSingleCore(
    UINT32 TargetCoreId, DEBUGGER_EVENT_OPTIONS *BroadcastingOption) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformEnableMovToControlRegisterExiting,
      &BroadcastingOption);
}

VOID ConfigureChangeMsrBitmapWriteOnSingleCore(UINT32 TargetCoreId,
                                               UINT64 MsrMask) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore,
      (PVOID)MsrMask);
}

VOID ConfigureChangeMsrBitmapReadOnSingleCore(UINT32 TargetCoreId,
                                              UINT64 MsrMask) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformChangeMsrBitmapReadOnSingleCore,
      (PVOID)MsrMask);
}

VOID ConfigureChangeIoBitmapOnSingleCore(UINT32 TargetCoreId, UINT64 Port) {
  DpcRoutineRunTaskOnSingleCore(
      TargetCoreId, (PVOID)DpcRoutinePerformChangeIoBitmapOnSingleCore,
      (PVOID)Port);
}
