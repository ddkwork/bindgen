
#include "pch.h"

VOID TerminateExternalInterruptEvent(PDEBUGGER_EVENT Event,
                                     BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->ExternalInterruptOccurredEventsHead) >
      1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();
    } else {
      ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();
    }
    TempList = &g_Events->ExternalInterruptOccurredEventsHead;
    while (&g_Events->ExternalInterruptOccurredEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();
    } else {
      ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();
    }
  }
}

VOID TerminateHiddenHookReadAndWriteAndExecuteEvent(PDEBUGGER_EVENT Event,
                                                    BOOLEAN InputFromVmxRoot) {
  if (InputFromVmxRoot) {
    TerminateEptHookUnHookAllHooksByHookingTagFromVmxRootAndApplyInvalidation(
        Event->Tag);
  } else {
    ConfigureEptHookUnHookAllByHookingTag(Event->Tag);
  }
}

VOID TerminateHiddenHookExecCcEvent(PDEBUGGER_EVENT Event,
                                    BOOLEAN InputFromVmxRoot) {
  if (InputFromVmxRoot) {
    TerminateEptHookUnHookSingleAddressFromVmxRootAndApplyInvalidation(
        Event->Options.OptionalParam1, (UINT64)NULL);
  } else {
    ConfigureEptHookUnHookSingleAddress(Event->Options.OptionalParam1,
                                        (UINT64)NULL, Event->ProcessId);
  }
}

VOID TerminateHiddenHookExecDetoursEvent(PDEBUGGER_EVENT Event,
                                         BOOLEAN InputFromVmxRoot) {
  if (InputFromVmxRoot) {
    TerminateEptHookUnHookSingleAddressFromVmxRootAndApplyInvalidation(
        (UINT64)NULL, Event->Options.OptionalParam1);
  } else {
    ConfigureEptHookUnHookSingleAddress(
        (UINT64)NULL, Event->Options.OptionalParam1, Event->ProcessId);
  }
}

VOID TerminateRdmsrExecutionEvent(PDEBUGGER_EVENT Event,
                                  BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->RdmsrInstructionExecutionEventsHead) >
      1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetMsrBitmapReadAllCores();
    } else {
      ExtensionCommandResetChangeAllMsrBitmapReadAllCores();
    }
    TempList = &g_Events->RdmsrInstructionExecutionEventsHead;
    while (&g_Events->RdmsrInstructionExecutionEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetMsrBitmapReadAllCores();
    } else {
      ExtensionCommandResetChangeAllMsrBitmapReadAllCores();
    }
  }
}

VOID TerminateWrmsrExecutionEvent(PDEBUGGER_EVENT Event,
                                  BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->WrmsrInstructionExecutionEventsHead) >
      1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetMsrBitmapWriteAllCores();
    } else {
      ExtensionCommandResetAllMsrBitmapWriteAllCores();
    }
    TempList = &g_Events->WrmsrInstructionExecutionEventsHead;
    while (&g_Events->WrmsrInstructionExecutionEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetMsrBitmapWriteAllCores();
    } else {
      ExtensionCommandResetAllMsrBitmapWriteAllCores();
    }
  }
}

VOID TerminateExceptionEvent(PDEBUGGER_EVENT Event, BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->ExceptionOccurredEventsHead) > 1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetExceptionBitmapOnlyOnClearingExceptionEventsAllCores();
    } else {
      ExtensionCommandResetExceptionBitmapAllCores();
    }
    TempList = &g_Events->ExceptionOccurredEventsHead;
    while (&g_Events->ExceptionOccurredEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetExceptionBitmapOnlyOnClearingExceptionEventsAllCores();
    } else {
      ExtensionCommandResetExceptionBitmapAllCores();
    }
  }
}

VOID TerminateInInstructionExecutionEvent(PDEBUGGER_EVENT Event,
                                          BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->InInstructionExecutionEventsHead) > 1 ||
      DebuggerEventListCount(&g_Events->OutInstructionExecutionEventsHead) >=
          1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetIoBitmapAllCores();
    } else {
      ExtensionCommandIoBitmapResetAllCores();
    }
    TempList = &g_Events->InInstructionExecutionEventsHead;
    while (&g_Events->InInstructionExecutionEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetIoBitmapAllCores();
    } else {
      ExtensionCommandIoBitmapResetAllCores();
    }
  }
}

VOID TerminateOutInstructionExecutionEvent(PDEBUGGER_EVENT Event,
                                           BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->OutInstructionExecutionEventsHead) >
          1 ||
      DebuggerEventListCount(&g_Events->InInstructionExecutionEventsHead) >=
          1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetIoBitmapAllCores();
    } else {
      ExtensionCommandIoBitmapResetAllCores();
    }
    TempList = &g_Events->OutInstructionExecutionEventsHead;
    while (&g_Events->OutInstructionExecutionEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastResetIoBitmapAllCores();
    } else {
      ExtensionCommandIoBitmapResetAllCores();
    }
  }
}

VOID TerminateVmcallExecutionEvent(PDEBUGGER_EVENT Event,
                                   BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(Event);
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  if (DebuggerEventListCount(&g_Events->VmcallInstructionExecutionEventsHead) >
      1) {
    return;
  } else {
    VmFuncSetTriggerEventForVmcalls(FALSE);
  }
}

VOID TerminateExecTrapModeChangedEvent(PDEBUGGER_EVENT Event,
                                       BOOLEAN InputFromVmxRoot) {
  if (DebuggerEventListCount(&g_Events->TrapExecutionModeChangedEventsHead) >
      1) {
    ConfigureExecTrapRemoveProcessFromWatchingList(Event->ProcessId);
    return;
  } else {
    if (!InputFromVmxRoot) {
      ConfigureUninitializeExecTrapOnAllProcessors();
    }
    ConfigureExecTrapRemoveProcessFromWatchingList(Event->ProcessId);
  }
}

VOID TerminateCpuidExecutionEvent(PDEBUGGER_EVENT Event,
                                  BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(Event);
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  if (DebuggerEventListCount(&g_Events->CpuidInstructionExecutionEventsHead) >
      1) {
    return;
  } else {
    VmFuncSetTriggerEventForCpuids(FALSE);
  }
}

VOID TerminateTscEvent(PDEBUGGER_EVENT Event, BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->TscInstructionExecutionEventsHead) >
      1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableRdtscExitingForClearingTscEventsAllCores();
    } else {
      ExtensionCommandDisableRdtscExitingForClearingEventsAllCores();
    }
    TempList = &g_Events->TscInstructionExecutionEventsHead;
    while (&g_Events->TscInstructionExecutionEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableRdtscExitingForClearingTscEventsAllCores();
    } else {
      ExtensionCommandDisableRdtscExitingForClearingEventsAllCores();
    }
  }
}

VOID TerminatePmcEvent(PDEBUGGER_EVENT Event, BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->PmcInstructionExecutionEventsHead) >
      1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableRdpmcExitingAllCores();
    } else {
      ExtensionCommandDisableRdpmcExitingAllCores();
    }
    TempList = &g_Events->PmcInstructionExecutionEventsHead;
    while (&g_Events->PmcInstructionExecutionEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableRdpmcExitingAllCores();
    } else {
      ExtensionCommandDisableRdpmcExitingAllCores();
    }
  }
}

VOID TerminateControlRegistersEvent(PDEBUGGER_EVENT Event,
                                    BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->ControlRegisterModifiedEventsHead) >
      1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableMov2CrExitingForClearingCrEventsAllCores(
          &Event->Options);
    } else {
      ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores(
          Event);
    }
    TempList = &g_Events->ControlRegisterModifiedEventsHead;
    while (&g_Events->ControlRegisterModifiedEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableMov2CrExitingForClearingCrEventsAllCores(
          &Event->Options);
    } else {
      ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores(
          Event);
    }
  }
}

VOID TerminateDebugRegistersEvent(PDEBUGGER_EVENT Event,
                                  BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->DebugRegistersAccessedEventsHead) > 1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableMov2DrExitingForClearingDrEventsAllCores();
    } else {
      ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores();
    }
    TempList = &g_Events->DebugRegistersAccessedEventsHead;
    while (&g_Events->DebugRegistersAccessedEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableMov2DrExitingForClearingDrEventsAllCores();
    } else {
      ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores();
    }
  }
}

VOID TerminateSyscallHookEferEvent(PDEBUGGER_EVENT Event,
                                   BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->SyscallHooksEferSyscallEventsHead) >
          1 ||
      DebuggerEventListCount(&g_Events->SyscallHooksEferSysretEventsHead) >=
          1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableEferSyscallEventsAllCores();
    } else {
      DebuggerEventDisableEferOnAllProcessors();
    }
    TempList = &g_Events->SyscallHooksEferSyscallEventsHead;
    while (&g_Events->SyscallHooksEferSyscallEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableEferSyscallEventsAllCores();
    } else {
      DebuggerEventDisableEferOnAllProcessors();
    }
  }
}

VOID TerminateSysretHookEferEvent(PDEBUGGER_EVENT Event,
                                  BOOLEAN InputFromVmxRoot) {
  PLIST_ENTRY TempList = 0;
  DEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn = {0};
  if (DebuggerEventListCount(&g_Events->SyscallHooksEferSysretEventsHead) > 1 ||
      DebuggerEventListCount(&g_Events->SyscallHooksEferSyscallEventsHead) >
          1) {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableEferSyscallEventsAllCores();
    } else {
      DebuggerEventDisableEferOnAllProcessors();
    }
    TempList = &g_Events->SyscallHooksEferSysretEventsHead;
    while (&g_Events->SyscallHooksEferSysretEventsHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag != Event->Tag) {
        DebuggerApplyEvent(CurrentEvent, &ResultsToReturn, InputFromVmxRoot);
        if (!ResultsToReturn.IsSuccessful) {
          LogInfo("Err, unable to re-apply previous events");
        }
      }
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedBroadcastDisableEferSyscallEventsAllCores();
    } else {
      DebuggerEventDisableEferOnAllProcessors();
    }
  }
}

BOOLEAN TerminateQueryDebuggerResourceExceptionBitmap(
    UINT32 CoreId, UINT32 *BitmapMask,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  if (!(PassOver & PASSING_OVER_EXCEPTION_EVENTS)) {
    *BitmapMask = *BitmapMask | DebuggerExceptionEventBitmapMask(CoreId);
  }
  if (!(PassOver & PASSING_OVER_UD_EXCEPTIONS_FOR_SYSCALL_SYSRET_HOOK)) {
    if (DebuggerEventListCountByEventType(SYSCALL_HOOK_EFER_SYSCALL, CoreId) !=
            0 ||
        DebuggerEventListCountByEventType(SYSCALL_HOOK_EFER_SYSRET, CoreId) !=
            0) {
      *BitmapMask = *BitmapMask | (1 << EXCEPTION_VECTOR_UNDEFINED_OPCODE);
    }
  }
  if (DebuggerQueryDebuggerStatus()) {
    *BitmapMask = *BitmapMask | (1 << EXCEPTION_VECTOR_BREAKPOINT);
    *BitmapMask = *BitmapMask | (1 << EXCEPTION_VECTOR_DEBUG_BREAKPOINT);
  }
  if (KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId(
          CoreId,
          DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_DEBUG_REGISTER_INTERCEPTION)) {
    *BitmapMask = *BitmapMask | (1 << EXCEPTION_VECTOR_DEBUG_BREAKPOINT);
  }
  return FALSE;
}

BOOLEAN TerminateQueryDebuggerResourceExternalInterruptExiting(
    UINT32 CoreId, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  if (!(PassOver & PASSING_OVER_INTERRUPT_EVENTS)) {
    if (DebuggerEventListCountByEventType(EXTERNAL_INTERRUPT_OCCURRED,
                                          CoreId) != 0) {
      return TRUE;
    }
  }
  if (KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId(
          CoreId,
          DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_THREAD_CHANGE) ||
      KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId(
          CoreId,
          DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_PROCESS_CHANGE)) {
    return TRUE;
  }
  return FALSE;
}

BOOLEAN TerminateQueryDebuggerResourceTscExiting(
    UINT32 CoreId, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  if (!(PassOver & PASSING_OVER_TSC_EVENTS)) {
    if (DebuggerEventListCountByEventType(TSC_INSTRUCTION_EXECUTION, CoreId) !=
        0) {
      return TRUE;
    }
  }
  return FALSE;
}

BOOLEAN TerminateQueryDebuggerResourceMov2DebugRegExiting(
    UINT32 CoreId, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  if (!(PassOver & PASSING_OVER_MOV_TO_HW_DEBUG_REGS_EVENTS)) {
    if (DebuggerEventListCountByEventType(DEBUG_REGISTERS_ACCESSED, CoreId) !=
        0) {
      return TRUE;
    }
  }
  if (KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId(
          CoreId,
          DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_DEBUG_REGISTER_INTERCEPTION)) {
    return TRUE;
  }
  return FALSE;
}

BOOLEAN TerminateQueryDebuggerResourceMovControlRegsExiting(
    UINT32 CoreId, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  if (!(PassOver & PASSING_OVER_MOV_TO_CONTROL_REGS_EVENTS)) {
    if (DebuggerEventListCountByEventType(CONTROL_REGISTER_MODIFIED, CoreId) !=
        0) {
      return TRUE;
    }
  }
  return FALSE;
}

BOOLEAN TerminateQueryDebuggerResourceMovToCr3Exiting(
    UINT32 CoreId, PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  UNREFERENCED_PARAMETER(PassOver);
  if (KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId(
          CoreId,
          DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_WAITING_FOR_MOV_CR3_VM_EXITS)) {
    return TRUE;
  }
  return FALSE;
}

BOOLEAN TerminateEptHookUnHookSingleAddressFromVmxRootAndApplyInvalidation(
    UINT64 VirtualAddress, UINT64 PhysAddress) {
  BOOLEAN Result = FALSE;
  EPT_SINGLE_HOOK_UNHOOKING_DETAILS TargetUnhookingDetails = {0};
  Result = ConfigureEptHookUnHookSingleAddressFromVmxRoot(
      VirtualAddress, PhysAddress, &TargetUnhookingDetails);
  if (Result == TRUE) {
    if (TargetUnhookingDetails.CallerNeedsToRestoreEntryAndInvalidateEpt) {
      HaltedBroadcastUnhookSinglePageAllCores(&TargetUnhookingDetails);
    }
    if (TargetUnhookingDetails.RemoveBreakpointInterception) {
      HaltedBroadcastUnSetExceptionBitmapAllCores(EXCEPTION_VECTOR_BREAKPOINT);
    }
    return TRUE;
  }
  return FALSE;
}

BOOLEAN
TerminateEptHookUnHookAllHooksByHookingTagFromVmxRootAndApplyInvalidation(
    UINT64 HookingTag) {
  BOOLEAN Result = FALSE;
  BOOLEAN IsAtLeastOneHookRemoved = FALSE;
  EPT_SINGLE_HOOK_UNHOOKING_DETAILS TargetUnhookingDetails = {0};
UnhookNextPossibleTag:
  Result = ConfigureEptHookUnHookSingleHookByHookingTagFromVmxRoot(
      HookingTag, &TargetUnhookingDetails);
  if (Result == TRUE) {
    IsAtLeastOneHookRemoved = TRUE;
    if (TargetUnhookingDetails.CallerNeedsToRestoreEntryAndInvalidateEpt) {
      HaltedBroadcastUnhookSinglePageAllCores(&TargetUnhookingDetails);
    }
    if (TargetUnhookingDetails.RemoveBreakpointInterception) {
      HaltedBroadcastUnSetExceptionBitmapAllCores(EXCEPTION_VECTOR_BREAKPOINT);
    }
    goto UnhookNextPossibleTag;
  }
  return IsAtLeastOneHookRemoved;
}

BOOLEAN TerminateQueryDebuggerResource(
    UINT32 CoreId, PROTECTED_HV_RESOURCES_TYPE ResourceType, PVOID Context,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  BOOLEAN Result = FALSE;
  switch (ResourceType) {
    case PROTECTED_HV_RESOURCES_EXCEPTION_BITMAP:
      Result = TerminateQueryDebuggerResourceExceptionBitmap(
          CoreId, (UINT32 *)Context, PassOver);
      break;
    case PROTECTED_HV_RESOURCES_EXTERNAL_INTERRUPT_EXITING:
      Result = TerminateQueryDebuggerResourceExternalInterruptExiting(CoreId,
                                                                      PassOver);
      break;
    case PROTECTED_HV_RESOURCES_RDTSC_RDTSCP_EXITING:
      Result = TerminateQueryDebuggerResourceTscExiting(CoreId, PassOver);
      break;
    case PROTECTED_HV_RESOURCES_MOV_TO_DEBUG_REGISTER_EXITING:
      Result =
          TerminateQueryDebuggerResourceMov2DebugRegExiting(CoreId, PassOver);
      break;
    case PROTECTED_HV_RESOURCES_MOV_CONTROL_REGISTER_EXITING:
      Result =
          TerminateQueryDebuggerResourceMovControlRegsExiting(CoreId, PassOver);
      break;
    case PROTECTED_HV_RESOURCES_MOV_TO_CR3_EXITING:
      Result = TerminateQueryDebuggerResourceMovToCr3Exiting(CoreId, PassOver);
      break;
    default:
      Result = FALSE;
      LogError("Err, invalid protected type");
      break;
  }
  return Result;
}
