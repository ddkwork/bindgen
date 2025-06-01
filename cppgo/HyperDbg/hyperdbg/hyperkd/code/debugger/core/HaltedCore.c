
#include "pch.h"

VOID HaltedCoreTaskTest(PROCESSOR_DEBUGGING_STATE *DbgState, PVOID Context) {
  UNREFERENCED_PARAMETER(DbgState);
  LogInfo("Target test task executed on halted core, context: %llx", Context);
}

VOID HaltedCorePerformTargetTask(PROCESSOR_DEBUGGING_STATE *DbgState,
                                 UINT64 TargetTask, PVOID Context) {
  switch (TargetTask) {
    case DEBUGGER_HALTED_CORE_TASK_TEST: {
      HaltedCoreTaskTest(DbgState, Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_RUN_VMCALL: {
      DirectVmcallTest(DbgState->CoreId, (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_SET_PROCESS_INTERCEPTION: {
      ProcessEnableOrDisableThreadChangeMonitor(DbgState, TRUE,
                                                PVOID_TO_BOOLEAN(Context));
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_SET_THREAD_INTERCEPTION: {
      ThreadEnableOrDisableThreadChangeMonitor(DbgState, TRUE,
                                               PVOID_TO_BOOLEAN(Context));
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_CHANGE_MSR_BITMAP_READ: {
      DirectVmcallChangeMsrBitmapRead(DbgState->CoreId,
                                      (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_CHANGE_MSR_BITMAP_WRITE: {
      DirectVmcallChangeMsrBitmapWrite(DbgState->CoreId,
                                       (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_CHANGE_IO_BITMAP: {
      DirectVmcallChangeIoBitmap(DbgState->CoreId,
                                 (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_SET_RDPMC_EXITING: {
      DirectVmcallEnableRdpmcExiting(DbgState->CoreId,
                                     (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_SET_RDTSC_EXITING: {
      DirectVmcallEnableRdtscpExiting(DbgState->CoreId,
                                      (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_ENABLE_MOV_TO_DEBUG_REGS_EXITING: {
      DirectVmcallEnableMov2DebugRegsExiting(
          DbgState->CoreId, (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_SET_EXCEPTION_BITMAP: {
      DirectVmcallSetExceptionBitmap(DbgState->CoreId,
                                     (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_ENABLE_EXTERNAL_INTERRUPT_EXITING: {
      DirectVmcallEnableExternalInterruptExiting(
          DbgState->CoreId, (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_ENABLE_MOV_TO_CONTROL_REGS_EXITING: {
      DirectVmcallEnableMovToCrExiting(DbgState->CoreId,
                                       (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_ENABLE_SYSCALL_HOOK_EFER: {
      DirectVmcallEnableEferSyscall(DbgState->CoreId,
                                    (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_INVEPT_ALL_CONTEXTS: {
      DirectVmcallInvalidateEptAllContexts(DbgState->CoreId,
                                           (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_INVEPT_SINGLE_CONTEXT: {
      DirectVmcallInvalidateSingleContext(DbgState->CoreId,
                                          (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_UNSET_EXCEPTION_BITMAP: {
      DirectVmcallUnsetExceptionBitmap(DbgState->CoreId,
                                       (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_UNHOOK_SINGLE_PAGE: {
      DirectVmcallUnhookSinglePage(DbgState->CoreId,
                                   (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS: {
      DirectVmcallSetDisableExternalInterruptExitingOnlyOnClearingInterruptEvents(
          DbgState->CoreId, (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_RESET_MSR_BITMAP_READ: {
      DirectVmcallResetMsrBitmapRead(DbgState->CoreId,
                                     (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_RESET_MSR_BITMAP_WRITE: {
      DirectVmcallResetMsrBitmapWrite(DbgState->CoreId,
                                      (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS: {
      DirectVmcallResetExceptionBitmapOnlyOnClearingExceptionEvents(
          DbgState->CoreId, (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_RESET_IO_BITMAP: {
      DirectVmcallResetIoBitmap(DbgState->CoreId,
                                (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS: {
      DirectVmcallDisableRdtscExitingForClearingTscEvents(
          DbgState->CoreId, (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_UNSET_RDPMC_EXITING: {
      DirectVmcallDisableRdpmcExiting(DbgState->CoreId,
                                      (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_DISABLE_SYSCALL_HOOK_EFER: {
      DirectVmcallDisableEferSyscallEvents(DbgState->CoreId,
                                           (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS: {
      DirectVmcallDisableMov2DrExitingForClearingDrEvents(
          DbgState->CoreId, (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    case DEBUGGER_HALTED_CORE_TASK_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS: {
      DirectVmcallDisableMov2CrExitingForClearingCrEvents(
          DbgState->CoreId, (DIRECT_VMCALL_PARAMETERS *)Context);
      break;
    }
    default:
      LogWarning("Warning, unknown broadcast on halted core received");
      break;
  }
}

VOID HaltedCoreApplyTaskOnTargetCore(UINT32 TargetCoreId, UINT64 TargetTask,
                                     BOOLEAN LockAgainAfterTask,
                                     PVOID Context) {
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[TargetCoreId];
  DbgState->HaltedCoreTask.PerformHaltedTask = TRUE;
  DbgState->HaltedCoreTask.KernelStatus = (UINT64)NULL;
  DbgState->HaltedCoreTask.LockAgainAfterTask = LockAgainAfterTask;
  DbgState->HaltedCoreTask.TargetTask = TargetTask;
  DbgState->HaltedCoreTask.Context = Context;
  KdUnlockTheHaltedCore(DbgState);
}

VOID HaltedCoreRunTaskOnSingleCore(UINT32 TargetCoreId, UINT64 TargetTask,
                                   BOOLEAN LockAgainAfterTask, PVOID Context) {
  if (TargetCoreId == KeGetCurrentProcessorNumberEx(NULL)) {
    HaltedCorePerformTargetTask(&g_DbgState[TargetCoreId], TargetTask, Context);
  } else {
    HaltedCoreApplyTaskOnTargetCore(TargetCoreId, TargetTask,
                                    LockAgainAfterTask, Context);
  }
}

BOOLEAN HaltedCoreBroadcastTaskAllCores(PROCESSOR_DEBUGGING_STATE *DbgState,
                                        UINT64 TargetTask,
                                        BOOLEAN LockAgainAfterTask,
                                        BOOLEAN Synchronize, PVOID Context) {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (Synchronize && !LockAgainAfterTask) {
    LogWarning(
        "Synchronization is not possible when the locking after the task is "
        "not expected");
    return FALSE;
  }
  for (UINT32 i = 0; i < ProcessorsCount; i++) {
    if (DbgState->CoreId != i) {
      HaltedCoreApplyTaskOnTargetCore(i, TargetTask, LockAgainAfterTask,
                                      Context);
    } else {
      HaltedCorePerformTargetTask(DbgState, TargetTask, Context);
    }
  }
  if (Synchronize) {
    for (size_t i = 0; i < ProcessorsCount; i++) {
      if (DbgState->CoreId != i) {
        while (TRUE) {
          if (KdCheckTheHaltedCore(&g_DbgState[i]) == FALSE) {
            continue;
          } else {
            break;
          }
        }
      }
    }
  }
  return TRUE;
}
