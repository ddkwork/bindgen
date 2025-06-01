
#include "pch.h"

VOID HaltedRoutineChangeAllMsrBitmapReadOnSingleCore(UINT32 TargetCoreId,
                                                     UINT64 BitmapMask) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_CHANGE_MSR_BITMAP_READ;
  DirectVmcallOptions.OptionalParam1 = BitmapMask;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineChangeAllMsrBitmapWriteOnSingleCore(UINT32 TargetCoreId,
                                                      UINT64 BitmapMask) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_CHANGE_MSR_BITMAP_WRITE;
  DirectVmcallOptions.OptionalParam1 = BitmapMask;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineChangeIoBitmapOnSingleCore(UINT32 TargetCoreId, UINT64 Port) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_CHANGE_IO_BITMAP;
  DirectVmcallOptions.OptionalParam1 = Port;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineEnableRdpmcExitingOnSingleCore(UINT32 TargetCoreId) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_SET_RDPMC_EXITING;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineEnableRdtscExitingOnSingleCore(UINT32 TargetCoreId) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_SET_RDTSC_EXITING;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineEnableMov2DebugRegsExitingOnSingleCore(UINT32 TargetCoreId) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_ENABLE_MOV_TO_DEBUG_REGS_EXITING;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineEnableExternalInterruptExiting(UINT32 TargetCoreId) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_ENABLE_EXTERNAL_INTERRUPT_EXITING;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineSetExceptionBitmapOnSingleCore(UINT32 TargetCoreId,
                                                 UINT64 ExceptionIndex) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_SET_EXCEPTION_BITMAP;
  DirectVmcallOptions.OptionalParam1 = ExceptionIndex;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineUnSetExceptionBitmapOnSingleCore(UINT32 TargetCoreId,
                                                   UINT64 ExceptionIndex) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_UNSET_EXCEPTION_BITMAP;
  DirectVmcallOptions.OptionalParam1 = ExceptionIndex;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineEnableMovToCrExitingOnSingleCore(
    UINT32 TargetCoreId, DEBUGGER_EVENT_OPTIONS *BroadcastingOption) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_ENABLE_MOV_TO_CONTROL_REGS_EXITING;
  DirectVmcallOptions.OptionalParam1 = BroadcastingOption->OptionalParam1;
  DirectVmcallOptions.OptionalParam2 = BroadcastingOption->OptionalParam2;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineEnableEferSyscallHookOnSingleCore(UINT32 TargetCoreId) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_ENABLE_SYSCALL_HOOK_EFER;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineInvalidateEptAllContextsOnSingleCore(UINT32 TargetCoreId) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_INVEPT_ALL_CONTEXTS;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}

VOID HaltedRoutineInvalidateSingleContextOnSingleCore(UINT32 TargetCoreId) {
  DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
  UINT64 HaltedCoreTask = (UINT64)NULL;
  HaltedCoreTask = DEBUGGER_HALTED_CORE_TASK_INVEPT_SINGLE_CONTEXT;
  HaltedCoreRunTaskOnSingleCore(TargetCoreId, HaltedCoreTask, TRUE,
                                &DirectVmcallOptions);
}
