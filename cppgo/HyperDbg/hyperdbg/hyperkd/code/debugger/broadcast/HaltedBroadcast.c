
#include "pch.h"

VOID HaltedBroadcastChangeAllMsrBitmapReadAllCores(UINT64 BitmapMask) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_CHANGE_MSR_BITMAP_READ;
  DirectVmcallOptions.OptionalParam1 = BitmapMask;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastChangeAllMsrBitmapWriteAllCores(UINT64 BitmapMask) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_CHANGE_MSR_BITMAP_WRITE;
  DirectVmcallOptions.OptionalParam1 = BitmapMask;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastChangeAllIoBitmapAllCores(UINT64 Port) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_CHANGE_IO_BITMAP;
  DirectVmcallOptions.OptionalParam1 = Port;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastEnableRdpmcExitingAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_SET_RDPMC_EXITING;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastEnableRdtscExitingAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_SET_RDTSC_EXITING;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastEnableMov2DebugRegsExitingAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_ENABLE_MOV_TO_DEBUG_REGS_EXITING;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastEnableExternalInterruptExitingAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_ENABLE_EXTERNAL_INTERRUPT_EXITING;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastSetExceptionBitmapAllCores(UINT64 ExceptionIndex) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_SET_EXCEPTION_BITMAP;
  DirectVmcallOptions.OptionalParam1 = ExceptionIndex;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastUnSetExceptionBitmapAllCores(UINT64 ExceptionIndex) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_UNSET_EXCEPTION_BITMAP;
  DirectVmcallOptions.OptionalParam1 = ExceptionIndex;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastEnableMovToCrExitingAllCores(
    DEBUGGER_EVENT_OPTIONS *BroadcastingOption) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_ENABLE_MOV_TO_CONTROL_REGS_EXITING;
  DirectVmcallOptions.OptionalParam1 = BroadcastingOption->OptionalParam1;
  DirectVmcallOptions.OptionalParam2 = BroadcastingOption->OptionalParam2;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastEnableEferSyscallHookAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_ENABLE_SYSCALL_HOOK_EFER;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastInvalidateEptAllContextsAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_INVEPT_ALL_CONTEXTS;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastInvalidateSingleContextAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_INVEPT_SINGLE_CONTEXT;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastUnhookSinglePageAllCores(
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *UnhookingDetail) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_UNHOOK_SINGLE_PAGE;
  DirectVmcallOptions.OptionalParam1 = UnhookingDetail->PhysicalAddress;
  DirectVmcallOptions.OptionalParam2 = UnhookingDetail->OriginalEntry;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask =
      DEBUGGER_HALTED_CORE_TASK_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastResetMsrBitmapReadAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_RESET_MSR_BITMAP_READ;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastResetMsrBitmapWriteAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_RESET_MSR_BITMAP_WRITE;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastResetExceptionBitmapOnlyOnClearingExceptionEventsAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask =
      DEBUGGER_HALTED_CORE_TASK_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastResetIoBitmapAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_RESET_IO_BITMAP;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastDisableRdtscExitingForClearingTscEventsAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask =
      DEBUGGER_HALTED_CORE_TASK_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastDisableRdpmcExitingAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_UNSET_RDPMC_EXITING;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastDisableEferSyscallEventsAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_DISABLE_SYSCALL_HOOK_EFER;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastDisableMov2DrExitingForClearingDrEventsAllCores() {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask =
      DEBUGGER_HALTED_CORE_TASK_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}

VOID HaltedBroadcastDisableMov2CrExitingForClearingCrEventsAllCores(
    DEBUGGER_EVENT_OPTIONS *BroadcastingOption) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask =
      DEBUGGER_HALTED_CORE_TASK_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS;
  DirectVmcallOptions.OptionalParam1 = BroadcastingOption->OptionalParam1;
  DirectVmcallOptions.OptionalParam2 = BroadcastingOption->OptionalParam2;
  HaltedCoreBroadcastTaskAllCores(
      &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)], HaltedCoreTask, TRUE,
      TRUE, &DirectVmcallOptions);
}
