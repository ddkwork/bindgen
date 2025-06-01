
#include "pch.h"

BOOLEAN ApplyEventMonitorEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UINT32 TempProcessId;
  BOOLEAN ResultOfApplyingEvent = FALSE;
  UINT64 RemainingSize;
  UINT64 PagesBytes;
  UINT64 ConstEndAddress;
  UINT64 TempStartAddress;
  UINT64 TempEndAddress;
  UINT64 TempNextPageAddr;
  EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR HookingAddresses = {0};
  if (InputFromVmxRoot) {
    TempProcessId = NULL_ZERO;
  } else {
    if (Event->ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES ||
        Event->ProcessId == 0) {
      TempProcessId = HANDLE_TO_UINT32(PsGetCurrentProcessId());
    } else {
      TempProcessId = Event->ProcessId;
    }
  }
  switch (Event->EventType) {
    case HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE:
    case HIDDEN_HOOK_READ_AND_EXECUTE:
      HookingAddresses.SetHookForRead = TRUE;
      HookingAddresses.SetHookForWrite = TRUE;
      HookingAddresses.SetHookForExec = TRUE;
      break;
    case HIDDEN_HOOK_WRITE_AND_EXECUTE:
      HookingAddresses.SetHookForRead = FALSE;
      HookingAddresses.SetHookForWrite = TRUE;
      HookingAddresses.SetHookForExec = FALSE;
      break;
    case HIDDEN_HOOK_READ_AND_WRITE:
    case HIDDEN_HOOK_READ:
      HookingAddresses.SetHookForRead = TRUE;
      HookingAddresses.SetHookForWrite = TRUE;
      HookingAddresses.SetHookForExec = FALSE;
      break;
    case HIDDEN_HOOK_WRITE:
      HookingAddresses.SetHookForRead = FALSE;
      HookingAddresses.SetHookForWrite = TRUE;
      HookingAddresses.SetHookForExec = FALSE;
      break;
    case HIDDEN_HOOK_EXECUTE:
      HookingAddresses.SetHookForRead = FALSE;
      HookingAddresses.SetHookForWrite = FALSE;
      HookingAddresses.SetHookForExec = TRUE;
      break;
    default:
      LogError("Err, Invalid monitor hook type");
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID;
      goto EventNotApplied;
      break;
  }
  HookingAddresses.Tag = Event->Tag;
  TempStartAddress = Event->InitOptions.OptionalParam1;
  TempEndAddress = Event->InitOptions.OptionalParam2;
  ConstEndAddress = TempEndAddress;
  PagesBytes = (UINT64)PAGE_ALIGN(TempStartAddress);
  PagesBytes = TempEndAddress - PagesBytes;
  PagesBytes = PagesBytes / PAGE_SIZE;
  RemainingSize = TempEndAddress - TempStartAddress;
  for (size_t i = 0; i <= PagesBytes; i++) {
    if (RemainingSize >= PAGE_SIZE) {
      TempEndAddress = (TempStartAddress +
                        ((UINT64)PAGE_ALIGN(TempStartAddress + PAGE_SIZE) -
                         TempStartAddress)) -
                       1;
      RemainingSize = ConstEndAddress - TempEndAddress - 1;
    } else {
      TempNextPageAddr = (UINT64)PAGE_ALIGN(TempStartAddress + RemainingSize);
      if (TempNextPageAddr > ((UINT64)PAGE_ALIGN(TempStartAddress))) {
        TempEndAddress = TempNextPageAddr - 1;
        RemainingSize = RemainingSize - (TempEndAddress - TempStartAddress) - 1;
      } else {
        TempEndAddress = TempStartAddress + RemainingSize;
        RemainingSize = 0;
      }
    }
    HookingAddresses.StartAddress = TempStartAddress;
    HookingAddresses.EndAddress = TempEndAddress;
    if ((DEBUGGER_HOOK_MEMORY_TYPE)Event->InitOptions.OptionalParam3 ==
        DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS) {
      HookingAddresses.MemoryType = DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS;
    } else {
      HookingAddresses.MemoryType = DEBUGGER_MEMORY_HOOK_VIRTUAL_ADDRESS;
    }
    ResultOfApplyingEvent = DebuggerEventEnableMonitorReadWriteExec(
        &HookingAddresses, TempProcessId, InputFromVmxRoot);
    if (!ResultOfApplyingEvent) {
      if (InputFromVmxRoot) {
        TerminateEptHookUnHookAllHooksByHookingTagFromVmxRootAndApplyInvalidation(
            Event->Tag);
      } else {
        ConfigureEptHookUnHookAllByHookingTag(Event->Tag);
      }
      break;
    } else {
      if (!InputFromVmxRoot) {
        PoolManagerCheckAndPerformAllocationAndDeallocation();
      }
    }
    TempStartAddress = TempEndAddress + 1;
  }
  if (InputFromVmxRoot) {
    HaltedBroadcastInvalidateSingleContextAllCores();
  }
  if ((DEBUGGER_HOOK_MEMORY_TYPE)Event->Options.OptionalParam3 ==
      DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS) {
    Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
    Event->Options.OptionalParam2 = Event->InitOptions.OptionalParam2;
  } else {
    if (InputFromVmxRoot) {
      Event->Options.OptionalParam1 =
          VirtualAddressToPhysicalAddressOnTargetProcess(
              (PVOID)Event->InitOptions.OptionalParam1);
      Event->Options.OptionalParam2 =
          VirtualAddressToPhysicalAddressOnTargetProcess(
              (PVOID)Event->InitOptions.OptionalParam2);
    } else {
      Event->Options.OptionalParam1 =
          VirtualAddressToPhysicalAddressByProcessId(
              (PVOID)Event->InitOptions.OptionalParam1, TempProcessId);
      Event->Options.OptionalParam2 =
          VirtualAddressToPhysicalAddressByProcessId(
              (PVOID)Event->InitOptions.OptionalParam2, TempProcessId);
    }
  }
  Event->Options.OptionalParam3 = Event->InitOptions.OptionalParam1;
  Event->Options.OptionalParam4 = Event->InitOptions.OptionalParam2;
  Event->Options.OptionalParam5 = Event->InitOptions.OptionalParam3;
  if (!ResultOfApplyingEvent) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error = DebuggerGetLastError();
    goto EventNotApplied;
  }
  return TRUE;
EventNotApplied:
  return FALSE;
}

BOOLEAN ApplyEventEptHookExecCcEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UINT32 TempProcessId;
  if (InputFromVmxRoot) {
    HaltedBroadcastSetExceptionBitmapAllCores(EXCEPTION_VECTOR_BREAKPOINT);
    if (!ConfigureEptHookFromVmxRoot(
            (PVOID)Event->InitOptions.OptionalParam1)) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DebuggerGetLastError();
      goto EventNotApplied;
    } else {
      HaltedBroadcastInvalidateSingleContextAllCores();
    }
  } else {
    if (Event->ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES ||
        Event->ProcessId == 0) {
      TempProcessId = HANDLE_TO_UINT32(PsGetCurrentProcessId());
    } else {
      TempProcessId = Event->ProcessId;
    }
    if (!ConfigureEptHook((PVOID)Event->InitOptions.OptionalParam1,
                          TempProcessId)) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DebuggerGetLastError();
      goto EventNotApplied;
    }
  }
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
  return TRUE;
EventNotApplied:
  return FALSE;
}

BOOLEAN ApplyEventEpthookInlineEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UINT32 TempProcessId;
  if (InputFromVmxRoot) {
    Event->Options.OptionalParam1 =
        VirtualAddressToPhysicalAddressOnTargetProcess(
            (PVOID)Event->InitOptions.OptionalParam1);
    if (!ConfigureEptHook2FromVmxRoot(KeGetCurrentProcessorNumberEx(NULL),
                                      (PVOID)Event->InitOptions.OptionalParam1,
                                      NULL)) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DebuggerGetLastError();
      goto EventNotApplied;
    } else {
      HaltedBroadcastInvalidateSingleContextAllCores();
    }
  } else {
    if (Event->ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES ||
        Event->ProcessId == 0) {
      TempProcessId = HANDLE_TO_UINT32(PsGetCurrentProcessId());
    } else {
      TempProcessId = Event->ProcessId;
    }
    Event->Options.OptionalParam1 = VirtualAddressToPhysicalAddressByProcessId(
        (PVOID)Event->InitOptions.OptionalParam1, TempProcessId);
    if (!ConfigureEptHook2(KeGetCurrentProcessorNumberEx(NULL),
                           (PVOID)Event->InitOptions.OptionalParam1, NULL,
                           TempProcessId)) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DebuggerGetLastError();
      goto EventNotApplied;
    }
  }
  return TRUE;
EventNotApplied:
  return FALSE;
}

VOID ApplyEventRdmsrExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastChangeAllMsrBitmapReadAllCores(
          Event->InitOptions.OptionalParam1);
    } else {
      ExtensionCommandChangeAllMsrBitmapReadAllCores(
          Event->InitOptions.OptionalParam1);
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineChangeAllMsrBitmapReadOnSingleCore(
          Event->CoreId, Event->InitOptions.OptionalParam1);
    } else {
      ConfigureChangeMsrBitmapReadOnSingleCore(
          Event->CoreId, Event->InitOptions.OptionalParam1);
    }
  }
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
}

VOID ApplyEventWrmsrExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastChangeAllMsrBitmapWriteAllCores(
          Event->InitOptions.OptionalParam1);
    } else {
      ExtensionCommandChangeAllMsrBitmapWriteAllCores(
          Event->InitOptions.OptionalParam1);
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineChangeAllMsrBitmapWriteOnSingleCore(
          Event->CoreId, Event->InitOptions.OptionalParam1);
    } else {
      ConfigureChangeMsrBitmapWriteOnSingleCore(
          Event->CoreId, Event->InitOptions.OptionalParam1);
    }
  }
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
}

VOID ApplyEventInOutExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastChangeAllIoBitmapAllCores(
          Event->InitOptions.OptionalParam1);
    } else {
      ExtensionCommandIoBitmapChangeAllCores(Event->InitOptions.OptionalParam1);
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineChangeIoBitmapOnSingleCore(
          Event->CoreId, Event->InitOptions.OptionalParam1);
    } else {
      ConfigureChangeIoBitmapOnSingleCore(Event->CoreId,
                                          Event->InitOptions.OptionalParam1);
    }
  }
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
}

VOID ApplyEventTscExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastEnableRdtscExitingAllCores();
    } else {
      ExtensionCommandEnableRdtscExitingAllCores();
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineEnableRdtscExitingOnSingleCore(Event->CoreId);
    } else {
      ConfigureEnableRdtscExitingOnSingleCore(Event->CoreId);
    }
  }
}

VOID ApplyEventRdpmcExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastEnableRdpmcExitingAllCores();
    } else {
      ExtensionCommandEnableRdpmcExitingAllCores();
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineEnableRdpmcExitingOnSingleCore(Event->CoreId);
    } else {
      ConfigureEnableRdpmcExitingOnSingleCore(Event->CoreId);
    }
  }
}

VOID ApplyEventMov2DebugRegExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (ResultsToReturn) {
      HaltedBroadcastEnableMov2DebugRegsExitingAllCores();
    } else {
      ExtensionCommandEnableMovDebugRegistersExitingAllCores();
    }
  } else {
    if (ResultsToReturn) {
      HaltedRoutineEnableMov2DebugRegsExitingOnSingleCore(Event->CoreId);
    } else {
      ConfigureEnableMovToDebugRegistersExitingOnSingleCore(Event->CoreId);
    }
  }
}

VOID ApplyEventControlRegisterAccessedEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
  Event->Options.OptionalParam2 = Event->InitOptions.OptionalParam2;
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastEnableMovToCrExitingAllCores(&Event->Options);
    } else {
      ExtensionCommandEnableMovControlRegisterExitingAllCores(Event);
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineEnableMovToCrExitingOnSingleCore(Event->CoreId,
                                                    &Event->Options);
    } else {
      ConfigureEnableMovToControlRegisterExitingOnSingleCore(Event->CoreId,
                                                             &Event->Options);
    }
  }
}

VOID ApplyEventExceptionEvent(PDEBUGGER_EVENT Event,
                              PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                              BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastSetExceptionBitmapAllCores(
          Event->InitOptions.OptionalParam1);
    } else {
      ExtensionCommandSetExceptionBitmapAllCores(
          Event->InitOptions.OptionalParam1);
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineSetExceptionBitmapOnSingleCore(
          Event->CoreId, Event->InitOptions.OptionalParam1);
    } else {
      ConfigureSetExceptionBitmapOnSingleCore(
          Event->CoreId, (UINT32)Event->InitOptions.OptionalParam1);
    }
  }
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
}

VOID ApplyEventInterruptEvent(PDEBUGGER_EVENT Event,
                              PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                              BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastEnableExternalInterruptExitingAllCores();
    } else {
      ExtensionCommandSetExternalInterruptExitingAllCores();
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineEnableExternalInterruptExiting(Event->CoreId);
    } else {
      ConfigureSetExternalInterruptExitingOnSingleCore(Event->CoreId);
    }
  }
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
}

VOID ApplyEventEferSyscallHookEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE SyscallHookType =
      DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY;
  if (Event->InitOptions.OptionalParam2 ==
      DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD) {
    SyscallHookType = DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD;
  } else if (Event->InitOptions.OptionalParam2 ==
             DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY) {
    SyscallHookType = DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY;
  }
  ConfigureSetEferSyscallOrSysretHookType(SyscallHookType);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastEnableEferSyscallHookAllCores();
    } else {
      DebuggerEventEnableEferOnAllProcessors();
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineEnableEferSyscallHookOnSingleCore(Event->CoreId);
    } else {
      ConfigureEnableEferSyscallHookOnSingleCore(Event->CoreId);
    }
  }
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
  Event->Options.OptionalParam2 = SyscallHookType;
}

VOID ApplyEventEferSysretHookEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE SyscallHookType =
      DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY;
  if (Event->InitOptions.OptionalParam2 ==
      DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD) {
    SyscallHookType = DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD;
  } else if (Event->InitOptions.OptionalParam2 ==
             DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY) {
    SyscallHookType = DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY;
  }
  ConfigureSetEferSyscallOrSysretHookType(SyscallHookType);
  if (Event->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (InputFromVmxRoot) {
      HaltedBroadcastEnableEferSyscallHookAllCores();
    } else {
      DebuggerEventEnableEferOnAllProcessors();
    }
  } else {
    if (InputFromVmxRoot) {
      HaltedRoutineEnableEferSyscallHookOnSingleCore(Event->CoreId);
    } else {
      ConfigureEnableEferSyscallHookOnSingleCore(Event->CoreId);
    }
  }
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
  Event->Options.OptionalParam2 = SyscallHookType;
}

VOID ApplyEventVmcallExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(Event);
  UNREFERENCED_PARAMETER(ResultsToReturn);
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  VmFuncSetTriggerEventForVmcalls(TRUE);
}

BOOLEAN ApplyEventTrapModeChangeEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
  if (InputFromVmxRoot) {
    if (VmFuncQueryModeExecTrap()) {
      ConfigureExecTrapAddProcessToWatchingList(Event->ProcessId);
    } else {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error =
          DEBUGGER_ERROR_THE_MODE_EXEC_TRAP_IS_NOT_INITIALIZED;
      return FALSE;
    }
  } else {
    ConfigureExecTrapAddProcessToWatchingList(Event->ProcessId);
    ConfigureInitializeExecTrapOnAllProcessors();
  }
  return TRUE;
}

VOID ApplyEventCpuidExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(Event);
  UNREFERENCED_PARAMETER(ResultsToReturn);
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  VmFuncSetTriggerEventForCpuids(TRUE);
}

VOID ApplyEventTracingEvent(PDEBUGGER_EVENT Event,
                            PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                            BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(ResultsToReturn);
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  Event->Options.OptionalParam1 = Event->InitOptions.OptionalParam1;
}
