
#include "pch.h"

UINT64 DebuggerGetRegValueWrapper(PGUEST_REGS GuestRegs, UINT32 RegId) {
  return GetRegValue(GuestRegs, RegId);
}

UINT32 DebuggerGetLastError() { return g_LastError; }

VOID DebuggerSetLastError(UINT32 LastError) { g_LastError = LastError; }

BOOLEAN DebuggerInitialize() {
  ULONG ProcessorsCount = KeQueryActiveProcessorCount(0);
  PROCESSOR_DEBUGGING_STATE *CurrentDebuggerState = NULL;
  if (!GlobalDebuggingStateAllocateZeroedMemory()) {
    return FALSE;
  }
  if (GlobalEventsAllocateZeroedMemory() == FALSE) {
    return FALSE;
  }
  for (UINT32 i = 0; i < ProcessorsCount; i++) {
    g_DbgState[i].CoreId = i;
  }
  InitializeListHead(&g_Events->EptHookExecCcEventsHead);
  InitializeListHead(&g_Events->HiddenHookReadAndWriteAndExecuteEventsHead);
  InitializeListHead(&g_Events->HiddenHookReadAndWriteEventsHead);
  InitializeListHead(&g_Events->HiddenHookReadAndExecuteEventsHead);
  InitializeListHead(&g_Events->HiddenHookWriteAndExecuteEventsHead);
  InitializeListHead(&g_Events->HiddenHookReadEventsHead);
  InitializeListHead(&g_Events->HiddenHookWriteEventsHead);
  InitializeListHead(&g_Events->HiddenHookExecuteEventsHead);
  InitializeListHead(&g_Events->EptHook2sExecDetourEventsHead);
  InitializeListHead(&g_Events->SyscallHooksEferSyscallEventsHead);
  InitializeListHead(&g_Events->SyscallHooksEferSysretEventsHead);
  InitializeListHead(&g_Events->CpuidInstructionExecutionEventsHead);
  InitializeListHead(&g_Events->RdmsrInstructionExecutionEventsHead);
  InitializeListHead(&g_Events->WrmsrInstructionExecutionEventsHead);
  InitializeListHead(&g_Events->ExceptionOccurredEventsHead);
  InitializeListHead(&g_Events->TscInstructionExecutionEventsHead);
  InitializeListHead(&g_Events->PmcInstructionExecutionEventsHead);
  InitializeListHead(&g_Events->InInstructionExecutionEventsHead);
  InitializeListHead(&g_Events->OutInstructionExecutionEventsHead);
  InitializeListHead(&g_Events->DebugRegistersAccessedEventsHead);
  InitializeListHead(&g_Events->ExternalInterruptOccurredEventsHead);
  InitializeListHead(&g_Events->VmcallInstructionExecutionEventsHead);
  InitializeListHead(&g_Events->TrapExecutionModeChangedEventsHead);
  InitializeListHead(&g_Events->TrapExecutionInstructionTraceEventsHead);
  InitializeListHead(&g_Events->ControlRegister3ModifiedEventsHead);
  InitializeListHead(&g_Events->ControlRegisterModifiedEventsHead);
  g_EnableDebuggerEvents = TRUE;
  VmFuncSetTriggerEventForVmcalls(FALSE);
  VmFuncSetTriggerEventForCpuids(FALSE);
  if (!g_ScriptGlobalVariables) {
    g_ScriptGlobalVariables =
        PlatformMemAllocateNonPagedPool(MAX_VAR_COUNT * sizeof(UINT64));
  }
  if (!g_ScriptGlobalVariables) {
    return FALSE;
  }
  RtlZeroMemory(g_ScriptGlobalVariables, MAX_VAR_COUNT * sizeof(UINT64));
  RtlZeroMemory(&g_TrapFlagState, sizeof(DEBUGGER_TRAP_FLAG_STATE));
  for (size_t i = 0; i < ProcessorsCount; i++) {
    CurrentDebuggerState = &g_DbgState[i];
    if (!CurrentDebuggerState->ScriptEngineCoreSpecificStackBuffer) {
      CurrentDebuggerState->ScriptEngineCoreSpecificStackBuffer =
          PlatformMemAllocateNonPagedPool(MAX_STACK_BUFFER_COUNT *
                                          sizeof(UINT64));
    }
    if (!CurrentDebuggerState->ScriptEngineCoreSpecificStackBuffer) {
      return FALSE;
    }
    RtlZeroMemory(CurrentDebuggerState->ScriptEngineCoreSpecificStackBuffer,
                  MAX_STACK_BUFFER_COUNT * sizeof(UINT64));
  }
  VmFuncVmxBroadcastInitialize();
  if (!AttachingInitialize()) {
    return FALSE;
  }
  ConfigureEptHookReservePreallocatedPoolsForEptHooks(
      MAXIMUM_NUMBER_OF_INITIAL_PREALLOCATED_EPT_HOOKS);
  if (!PoolManagerCheckAndPerformAllocationAndDeallocation()) {
    LogWarning(
        "Warning, cannot allocate the pre-allocated pools for EPT hooks");
  }
  return TRUE;
}

VOID DebuggerUninitialize() {
  ULONG ProcessorsCount;
  PROCESSOR_DEBUGGING_STATE *CurrentDebuggerState = NULL;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  g_EnableDebuggerEvents = FALSE;
  if (g_KernelDebuggerState && EnableInstantEventMechanism) {
    DebuggerClearAllEvents(FALSE, TRUE);
  } else {
    DebuggerClearAllEvents(FALSE, FALSE);
  }
  KdUninitializeKernelDebugger();
  UdUninitializeUserDebugger();
  VmFuncVmxBroadcastUninitialize();
  GlobalEventsFreeMemory();
  if (g_ScriptGlobalVariables != NULL) {
    PlatformMemFreePool(g_ScriptGlobalVariables);
    g_ScriptGlobalVariables = NULL;
  }
  for (SIZE_T i = 0; i < ProcessorsCount; i++) {
    CurrentDebuggerState = &g_DbgState[i];
    if (CurrentDebuggerState->ScriptEngineCoreSpecificStackBuffer != NULL) {
      PlatformMemFreePool(
          CurrentDebuggerState->ScriptEngineCoreSpecificStackBuffer);
      CurrentDebuggerState->ScriptEngineCoreSpecificStackBuffer = NULL;
    }
  }
  GlobalDebuggingStateFreeMemory();
}

PDEBUGGER_EVENT DebuggerCreateEvent(
    BOOLEAN Enabled, UINT32 CoreId, UINT32 ProcessId,
    VMM_EVENT_TYPE_ENUM EventType, UINT64 Tag, DEBUGGER_EVENT_OPTIONS *Options,
    UINT32 ConditionsBufferSize, PVOID ConditionBuffer,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  PDEBUGGER_EVENT Event = NULL;
  UINT32 EventBufferSize = sizeof(DEBUGGER_EVENT) + ConditionsBufferSize;
  if (InputFromVmxRoot) {
    if (REGULAR_INSTANT_EVENT_CONDITIONAL_BUFFER >= EventBufferSize) {
      Event = (DEBUGGER_EVENT *)PoolManagerRequestPool(
          INSTANT_REGULAR_EVENT_BUFFER, TRUE,
          REGULAR_INSTANT_EVENT_CONDITIONAL_BUFFER);
      if (!Event) {
        Event = (DEBUGGER_EVENT *)PoolManagerRequestPool(
            INSTANT_BIG_EVENT_BUFFER, TRUE,
            BIG_INSTANT_EVENT_CONDITIONAL_BUFFER);
        if (!Event) {
          ResultsToReturn->IsSuccessful = FALSE;
          ResultsToReturn->Error =
              DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND;
          return NULL;
        }
      }
    } else if (BIG_INSTANT_EVENT_CONDITIONAL_BUFFER >= EventBufferSize) {
      Event = (DEBUGGER_EVENT *)PoolManagerRequestPool(
          INSTANT_BIG_EVENT_BUFFER, TRUE, BIG_INSTANT_EVENT_CONDITIONAL_BUFFER);
      if (!Event) {
        ResultsToReturn->IsSuccessful = FALSE;
        ResultsToReturn->Error =
            DEBUGGER_ERROR_INSTANT_EVENT_BIG_PREALLOCATED_BUFFER_NOT_FOUND;
        return NULL;
      }
    } else {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error =
          DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_EVENT_AND_CONDITIONALS;
      return NULL;
    }
  } else {
    Event = PlatformMemAllocateZeroedNonPagedPool(EventBufferSize);
    if (!Event) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT;
      return NULL;
    }
  }
  Event->CoreId = CoreId;
  Event->ProcessId = ProcessId;
  Event->Enabled = Enabled;
  Event->EventType = EventType;
  Event->Tag = Tag;
  Event->CountOfActions = 0;
  memcpy(&Event->InitOptions, Options, sizeof(DEBUGGER_EVENT_OPTIONS));
  if (ConditionBuffer != 0) {
    Event->ConditionsBufferSize = ConditionsBufferSize;
    Event->ConditionBufferAddress =
        (PVOID)((UINT64)Event + sizeof(DEBUGGER_EVENT));
    memcpy(Event->ConditionBufferAddress, ConditionBuffer,
           ConditionsBufferSize);
  } else {
    Event->ConditionsBufferSize = 0;
  }
  InitializeListHead(&Event->ActionsListHead);
  return Event;
}

PVOID DebuggerAllocateSafeRequestedBuffer(
    SIZE_T SizeOfRequestedSafeBuffer,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  PVOID RequestedBuffer = NULL;
  if (InputFromVmxRoot) {
    if (REGULAR_INSTANT_EVENT_REQUESTED_SAFE_BUFFER >=
        SizeOfRequestedSafeBuffer) {
      RequestedBuffer = (PVOID)PoolManagerRequestPool(
          INSTANT_REGULAR_SAFE_BUFFER_FOR_EVENTS, TRUE,
          REGULAR_INSTANT_EVENT_REQUESTED_SAFE_BUFFER);
      if (!RequestedBuffer) {
        RequestedBuffer = (PVOID)PoolManagerRequestPool(
            INSTANT_BIG_SAFE_BUFFER_FOR_EVENTS, TRUE,
            BIG_INSTANT_EVENT_REQUESTED_SAFE_BUFFER);
        if (!RequestedBuffer) {
          ResultsToReturn->IsSuccessful = FALSE;
          ResultsToReturn->Error =
              DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_REQUESTED_SAFE_BUFFER_NOT_FOUND;
          return NULL;
        }
      }
    } else if (BIG_INSTANT_EVENT_REQUESTED_SAFE_BUFFER >=
               SizeOfRequestedSafeBuffer) {
      RequestedBuffer = (PVOID)PoolManagerRequestPool(
          INSTANT_BIG_SAFE_BUFFER_FOR_EVENTS, TRUE,
          BIG_INSTANT_EVENT_REQUESTED_SAFE_BUFFER);
      if (!RequestedBuffer) {
        ResultsToReturn->IsSuccessful = FALSE;
        ResultsToReturn->Error =
            DEBUGGER_ERROR_INSTANT_EVENT_BIG_REQUESTED_SAFE_BUFFER_NOT_FOUND;
        return NULL;
      }
    } else {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error =
          DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_REQUESTED_SAFE_BUFFER;
      return NULL;
    }
  } else {
    RequestedBuffer =
        PlatformMemAllocateZeroedNonPagedPool(SizeOfRequestedSafeBuffer);
    if (!RequestedBuffer) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error =
          DEBUGGER_ERROR_UNABLE_TO_ALLOCATE_REQUESTED_SAFE_BUFFER;
      return NULL;
    }
  }
  return RequestedBuffer;
}

PDEBUGGER_EVENT_ACTION DebuggerAddActionToEvent(
    PDEBUGGER_EVENT Event, DEBUGGER_EVENT_ACTION_TYPE_ENUM ActionType,
    BOOLEAN SendTheResultsImmediately,
    PDEBUGGER_EVENT_REQUEST_CUSTOM_CODE InTheCaseOfCustomCode,
    PDEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION InTheCaseOfRunScript,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  PDEBUGGER_EVENT_ACTION Action;
  SIZE_T ActionBufferSize;
  PVOID RequestedBuffer = NULL;
  if (InTheCaseOfCustomCode != NULL) {
    ActionBufferSize = sizeof(DEBUGGER_EVENT_ACTION) +
                       InTheCaseOfCustomCode->CustomCodeBufferSize;
  } else if (InTheCaseOfRunScript != NULL) {
    ActionBufferSize =
        sizeof(DEBUGGER_EVENT_ACTION) + InTheCaseOfRunScript->ScriptLength;
  } else {
    ActionBufferSize = sizeof(DEBUGGER_EVENT_ACTION);
  }
  if (InputFromVmxRoot) {
    if (REGULAR_INSTANT_EVENT_ACTION_BUFFER >= ActionBufferSize) {
      Action = (DEBUGGER_EVENT_ACTION *)PoolManagerRequestPool(
          INSTANT_REGULAR_EVENT_ACTION_BUFFER, TRUE,
          REGULAR_INSTANT_EVENT_ACTION_BUFFER);
      if (!Action) {
        Action = (DEBUGGER_EVENT_ACTION *)PoolManagerRequestPool(
            INSTANT_BIG_EVENT_ACTION_BUFFER, TRUE,
            BIG_INSTANT_EVENT_ACTION_BUFFER);
        if (!Action) {
          ResultsToReturn->IsSuccessful = FALSE;
          ResultsToReturn->Error =
              DEBUGGER_ERROR_INSTANT_EVENT_ACTION_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND;
          return NULL;
        }
      }
    } else if (BIG_INSTANT_EVENT_ACTION_BUFFER >= ActionBufferSize) {
      Action = (DEBUGGER_EVENT_ACTION *)PoolManagerRequestPool(
          INSTANT_BIG_EVENT_ACTION_BUFFER, TRUE,
          BIG_INSTANT_EVENT_ACTION_BUFFER);
      if (!Action) {
        ResultsToReturn->IsSuccessful = FALSE;
        ResultsToReturn->Error =
            DEBUGGER_ERROR_INSTANT_EVENT_ACTION_BIG_PREALLOCATED_BUFFER_NOT_FOUND;
        return NULL;
      }
    } else {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error =
          DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_ACTION_BUFFER;
      return NULL;
    }
  } else {
    Action = PlatformMemAllocateZeroedNonPagedPool(ActionBufferSize);
    if (Action == NULL) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error =
          DEBUGGER_ERROR_UNABLE_TO_CREATE_ACTION_CANNOT_ALLOCATE_BUFFER;
      return NULL;
    }
  }
  if (ActionType == RUN_CUSTOM_CODE && InTheCaseOfCustomCode != NULL &&
      InTheCaseOfCustomCode->OptionalRequestedBufferSize != 0) {
    if (InTheCaseOfCustomCode->OptionalRequestedBufferSize >=
        MaximumPacketsCapacity) {
      if (InputFromVmxRoot) {
        PoolManagerFreePool((UINT64)Action);
      } else {
        PlatformMemFreePool(Action);
      }
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error =
          DEBUGGER_ERROR_INSTANT_EVENT_REQUESTED_OPTIONAL_BUFFER_IS_BIGGER_THAN_DEBUGGERS_SEND_RECEIVE_STACK;
      return NULL;
    }
    RequestedBuffer = DebuggerAllocateSafeRequestedBuffer(
        InTheCaseOfCustomCode->OptionalRequestedBufferSize, ResultsToReturn,
        InputFromVmxRoot);
    if (!RequestedBuffer) {
      if (InputFromVmxRoot) {
        PoolManagerFreePool((UINT64)Action);
      } else {
        PlatformMemFreePool(Action);
      }
      return NULL;
    }
    Action->RequestedBuffer.EnabledRequestBuffer = TRUE;
    Action->RequestedBuffer.RequestBufferSize =
        InTheCaseOfCustomCode->OptionalRequestedBufferSize;
    Action->RequestedBuffer.RequstBufferAddress = (UINT64)RequestedBuffer;
  }
  if (ActionType == RUN_SCRIPT && InTheCaseOfRunScript != NULL &&
      InTheCaseOfRunScript->OptionalRequestedBufferSize != 0) {
    if (InTheCaseOfRunScript->OptionalRequestedBufferSize >=
        MaximumPacketsCapacity) {
      if (InputFromVmxRoot) {
        PoolManagerFreePool((UINT64)Action);
      } else {
        PlatformMemFreePool(Action);
      }
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error =
          DEBUGGER_ERROR_INSTANT_EVENT_REQUESTED_OPTIONAL_BUFFER_IS_BIGGER_THAN_DEBUGGERS_SEND_RECEIVE_STACK;
      return NULL;
    }
    RequestedBuffer = DebuggerAllocateSafeRequestedBuffer(
        InTheCaseOfRunScript->OptionalRequestedBufferSize, ResultsToReturn,
        InputFromVmxRoot);
    if (!RequestedBuffer) {
      if (InputFromVmxRoot) {
        PoolManagerFreePool((UINT64)Action);
      } else {
        PlatformMemFreePool(Action);
      }
      return NULL;
    }
    Action->RequestedBuffer.EnabledRequestBuffer = TRUE;
    Action->RequestedBuffer.RequestBufferSize =
        InTheCaseOfRunScript->OptionalRequestedBufferSize;
    Action->RequestedBuffer.RequstBufferAddress = (UINT64)RequestedBuffer;
  }
  if (ActionType == RUN_CUSTOM_CODE && InTheCaseOfCustomCode != NULL) {
    if (InTheCaseOfCustomCode != NULL &&
        InTheCaseOfCustomCode->CustomCodeBufferSize == 0) {
      if (InputFromVmxRoot) {
        PoolManagerFreePool((UINT64)Action);
        if (RequestedBuffer != NULL) {
          PoolManagerFreePool((UINT64)RequestedBuffer);
        }
      } else {
        PlatformMemFreePool(Action);
        if (RequestedBuffer != NULL) {
          PlatformMemFreePool(RequestedBuffer);
        }
      }
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO;
      return NULL;
    }
    Action->CustomCodeBufferSize = InTheCaseOfCustomCode->CustomCodeBufferSize;
    Action->CustomCodeBufferAddress =
        (PVOID)((UINT64)Action + sizeof(DEBUGGER_EVENT_ACTION));
    memcpy(Action->CustomCodeBufferAddress,
           InTheCaseOfCustomCode->CustomCodeBufferAddress,
           InTheCaseOfCustomCode->CustomCodeBufferSize);
  } else if (ActionType == RUN_SCRIPT && InTheCaseOfRunScript != NULL) {
    if (InTheCaseOfRunScript->ScriptBuffer == NULL64_ZERO ||
        InTheCaseOfRunScript->ScriptLength == NULL_ZERO) {
      if (InputFromVmxRoot) {
        PoolManagerFreePool((UINT64)Action);
        if (RequestedBuffer != 0) {
          PoolManagerFreePool((UINT64)RequestedBuffer);
        }
      } else {
        PlatformMemFreePool(Action);
        if (RequestedBuffer != 0) {
          PlatformMemFreePool(RequestedBuffer);
        }
      }
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO;
      return NULL;
    }
    Action->ScriptConfiguration.ScriptBuffer =
        (UINT64)((BYTE *)Action + sizeof(DEBUGGER_EVENT_ACTION));
    RtlCopyMemory((void *)Action->ScriptConfiguration.ScriptBuffer,
                  (const void *)InTheCaseOfRunScript->ScriptBuffer,
                  InTheCaseOfRunScript->ScriptLength);
    Action->ScriptConfiguration.ScriptLength =
        InTheCaseOfRunScript->ScriptLength;
    Action->ScriptConfiguration.ScriptPointer =
        InTheCaseOfRunScript->ScriptPointer;
    Action->ScriptConfiguration.OptionalRequestedBufferSize =
        InTheCaseOfRunScript->OptionalRequestedBufferSize;
  }
  Event->CountOfActions++;
  Action->ActionOrderCode = Event->CountOfActions;
  Action->ImmediatelySendTheResults = SendTheResultsImmediately;
  Action->ActionType = ActionType;
  Action->Tag = Event->Tag;
  InsertHeadList(&Event->ActionsListHead, &(Action->ActionsList));
  return Action;
}

BOOLEAN DebuggerRegisterEvent(PDEBUGGER_EVENT Event) {
  PLIST_ENTRY TargetEventList = NULL;
  TargetEventList = DebuggerGetEventListByEventType(Event->EventType);
  if (TargetEventList != NULL) {
    InsertHeadList(TargetEventList, &(Event->EventsOfSameTypeList));
    return TRUE;
  } else {
    return FALSE;
  }
}

VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE DebuggerTriggerEvents(
    VMM_EVENT_TYPE_ENUM EventType,
    VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE CallingStage, PVOID Context,
    BOOLEAN *PostEventRequired, GUEST_REGS *Regs) {
  PROCESSOR_DEBUGGING_STATE *DbgState = NULL;
  DebuggerCheckForCondition *ConditionFunc;
  DEBUGGER_TRIGGERED_EVENT_DETAILS EventTriggerDetail = {0};
  PEPT_HOOKS_CONTEXT EptContext;
  PLIST_ENTRY TempList = 0;
  PLIST_ENTRY TempList2 = 0;
  const PVOID OriginalContext = Context;
  if (!g_EnableDebuggerEvents ||
      g_InterceptBreakpointsAndEventsForCommandsInRemoteComputer) {
    return VMM_CALLBACK_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED;
  }
  DbgState = &g_DbgState[KeGetCurrentProcessorNumberEx(NULL)];
  DbgState->Regs = Regs;
  TempList = DebuggerGetEventListByEventType(EventType);
  TempList2 = TempList;
  if (TempList == NULL) {
    return VMM_CALLBACK_TRIGGERING_EVENT_STATUS_INVALID_EVENT_TYPE;
  }
  while (TempList2 != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGER_EVENT CurrentEvent =
        CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
    if (!CurrentEvent->Enabled) {
      continue;
    }
    if (CurrentEvent->CoreId != DEBUGGER_EVENT_APPLY_TO_ALL_CORES &&
        CurrentEvent->CoreId != DbgState->CoreId) {
      continue;
    }
    if (CurrentEvent->ProcessId != DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES &&
        CurrentEvent->ProcessId != HANDLE_TO_UINT32(PsGetCurrentProcessId())) {
      continue;
    }
    switch (CurrentEvent->EventType) {
      case EXTERNAL_INTERRUPT_OCCURRED:
        if ((UINT64)Context != CurrentEvent->Options.OptionalParam1) {
          continue;
        }
        break;
      case HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE:
      case HIDDEN_HOOK_READ_AND_WRITE:
      case HIDDEN_HOOK_READ_AND_EXECUTE:
      case HIDDEN_HOOK_WRITE_AND_EXECUTE:
      case HIDDEN_HOOK_READ:
      case HIDDEN_HOOK_WRITE:
      case HIDDEN_HOOK_EXECUTE:
        EptContext = (PEPT_HOOKS_CONTEXT)OriginalContext;
        if (EptContext->HookingTag != CurrentEvent->Tag) {
          continue;
        } else {
          Context = (PVOID)EptContext->VirtualAddress;
        }
        break;
      case HIDDEN_HOOK_EXEC_CC:
        if ((UINT64)Context != CurrentEvent->Options.OptionalParam1) {
          continue;
        }
        break;
      case HIDDEN_HOOK_EXEC_DETOURS:
        EptContext = (PEPT_HOOKS_CONTEXT)OriginalContext;
        if (EptContext->PhysicalAddress !=
            CurrentEvent->Options.OptionalParam1) {
          continue;
        } else {
          Context = (PVOID)(EptContext->VirtualAddress);
        }
        break;
      case RDMSR_INSTRUCTION_EXECUTION:
      case WRMSR_INSTRUCTION_EXECUTION:
        if (CurrentEvent->Options.OptionalParam1 !=
                DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS &&
            CurrentEvent->Options.OptionalParam1 != (UINT64)Context) {
          continue;
        }
        break;
      case EXCEPTION_OCCURRED:
        if (CurrentEvent->Options.OptionalParam1 !=
                DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES &&
            CurrentEvent->Options.OptionalParam1 != (UINT64)Context) {
          continue;
        }
        break;
      case IN_INSTRUCTION_EXECUTION:
      case OUT_INSTRUCTION_EXECUTION:
        if (CurrentEvent->Options.OptionalParam1 !=
                DEBUGGER_EVENT_ALL_IO_PORTS &&
            CurrentEvent->Options.OptionalParam1 != (UINT64)Context) {
          continue;
        }
        break;
      case SYSCALL_HOOK_EFER_SYSCALL:
        if (CurrentEvent->Options.OptionalParam1 !=
                DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS &&
            CurrentEvent->Options.OptionalParam1 != (UINT64)Context) {
          continue;
        }
        break;
      case CPUID_INSTRUCTION_EXECUTION:
        if (CurrentEvent->Options.OptionalParam1 != (UINT64)NULL &&
            CurrentEvent->Options.OptionalParam2 != (UINT64)Context) {
          continue;
        }
        break;
      case CONTROL_REGISTER_MODIFIED:
        if (CurrentEvent->Options.OptionalParam1 != (UINT64)Context) {
          continue;
        }
        break;
      case TRAP_EXECUTION_MODE_CHANGED:
        if (CurrentEvent->Options.OptionalParam1 !=
            DEBUGGER_EVENT_MODE_TYPE_USER_MODE_AND_KERNEL_MODE) {
          if ((CurrentEvent->Options.OptionalParam1 ==
                   DEBUGGER_EVENT_MODE_TYPE_USER_MODE &&
               Context == (PVOID)DEBUGGER_EVENT_MODE_TYPE_KERNEL_MODE) ||
              (CurrentEvent->Options.OptionalParam1 ==
                   DEBUGGER_EVENT_MODE_TYPE_KERNEL_MODE &&
               Context == (PVOID)DEBUGGER_EVENT_MODE_TYPE_USER_MODE)) {
            continue;
          }
        }
        break;
      default:
        break;
    }
    if (CallingStage == VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION &&
        (CurrentEvent->EventMode ==
             VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION ||
         CurrentEvent->EventMode ==
             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION)) {
      *PostEventRequired = TRUE;
      if (CurrentEvent->EventMode ==
          VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION) {
        continue;
      }
    }
    if (CurrentEvent->ConditionsBufferSize != 0) {
      ConditionFunc =
          (DebuggerCheckForCondition *)CurrentEvent->ConditionBufferAddress;
      if (AsmDebuggerConditionCodeHandler((UINT64)DbgState->Regs,
                                          (UINT64)Context,
                                          (UINT64)ConditionFunc) == 0) {
        continue;
      }
    }
    DbgState->ShortCircuitingEvent = CurrentEvent->EnableShortCircuiting;
    EventTriggerDetail.Context = Context;
    EventTriggerDetail.Tag = CurrentEvent->Tag;
    EventTriggerDetail.Stage = CallingStage;
    DebuggerPerformActions(DbgState, CurrentEvent, &EventTriggerDetail);
  }
  if (DbgState->ShortCircuitingEvent) {
    DbgState->ShortCircuitingEvent = FALSE;
    return VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT;
  } else {
    return VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL;
  }
}

VOID DebuggerPerformActions(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGER_EVENT *Event,
    DEBUGGER_TRIGGERED_EVENT_DETAILS *EventTriggerDetail) {
  PLIST_ENTRY TempList = 0;
  TempList = &Event->ActionsListHead;
  while (&Event->ActionsListHead != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGER_EVENT_ACTION CurrentAction =
        CONTAINING_RECORD(TempList, DEBUGGER_EVENT_ACTION, ActionsList);
    switch (CurrentAction->ActionType) {
      case BREAK_TO_DEBUGGER:
        DebuggerPerformBreakToDebugger(DbgState, CurrentAction,
                                       EventTriggerDetail);
        break;
      case RUN_SCRIPT:
        DebuggerPerformRunScript(DbgState, CurrentAction, NULL,
                                 EventTriggerDetail);
        break;
      case RUN_CUSTOM_CODE:
        DebuggerPerformRunTheCustomCode(DbgState, CurrentAction,
                                        EventTriggerDetail);
        break;
      default:
        break;
    }
  }
}

BOOLEAN DebuggerPerformRunScript(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGER_EVENT_ACTION *Action,
    DEBUGGEE_SCRIPT_PACKET *ScriptDetails,
    DEBUGGER_TRIGGERED_EVENT_DETAILS *EventTriggerDetail) {
  SYMBOL_BUFFER CodeBuffer = {0};
  ACTION_BUFFER ActionBuffer = {0};
  SYMBOL ErrorSymbol = {0};
  SCRIPT_ENGINE_GENERAL_REGISTERS ScriptGeneralRegisters = {0};
  if (Action != NULL) {
    if (EventTriggerDetail->Stage ==
        VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION) {
      ActionBuffer.CallingStage = 1;
    } else {
      ActionBuffer.CallingStage = 0;
    }
    ActionBuffer.Context = (UINT64)EventTriggerDetail->Context;
    ActionBuffer.Tag = EventTriggerDetail->Tag;
    ActionBuffer.ImmediatelySendTheResults = Action->ImmediatelySendTheResults;
    ActionBuffer.CurrentAction = (UINT64)Action;
    CodeBuffer.Head = (PSYMBOL)Action->ScriptConfiguration.ScriptBuffer;
    CodeBuffer.Size = Action->ScriptConfiguration.ScriptLength;
    CodeBuffer.Pointer = Action->ScriptConfiguration.ScriptPointer;
  } else if (ScriptDetails != NULL) {
    if (EventTriggerDetail->Stage ==
        VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION) {
      ActionBuffer.CallingStage = 1;
    } else {
      ActionBuffer.CallingStage = 0;
    }
    ActionBuffer.Context = (UINT64)EventTriggerDetail->Context;
    ActionBuffer.Tag = EventTriggerDetail->Tag;
    ActionBuffer.ImmediatelySendTheResults = TRUE;
    ActionBuffer.CurrentAction = (UINT64)NULL;
    CodeBuffer.Head =
        (SYMBOL *)((CHAR *)ScriptDetails + sizeof(DEBUGGEE_SCRIPT_PACKET));
    CodeBuffer.Size = ScriptDetails->ScriptBufferSize;
    CodeBuffer.Pointer = ScriptDetails->ScriptBufferPointer;
  } else {
    return FALSE;
  }
  ScriptGeneralRegisters.StackBuffer =
      DbgState->ScriptEngineCoreSpecificStackBuffer;
  ScriptGeneralRegisters.GlobalVariablesList = g_ScriptGlobalVariables;
  RtlZeroMemory(ScriptGeneralRegisters.StackBuffer,
                MAX_STACK_BUFFER_COUNT * sizeof(UINT64));
  UINT64 EXECUTENUMBER = 0;
  for (UINT64 i = 0; i < CodeBuffer.Pointer;) {
    if (ScriptEngineExecute(DbgState->Regs, &ActionBuffer,
                            &ScriptGeneralRegisters, &CodeBuffer, &i,
                            &ErrorSymbol) == TRUE) {
      LogInfo("Err, ScriptEngineExecute, function = % s\n ",
              FunctionNames[ErrorSymbol.Value]);
      break;
    } else if (ScriptGeneralRegisters.StackIndx >= MAX_STACK_BUFFER_COUNT) {
      LogInfo(
          "Err, stack buffer overflow (more information: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/customize-build/"
          "change-script-engine-limitations)\n");
      break;
    } else if (EXECUTENUMBER >= MAX_EXECUTION_COUNT) {
      LogInfo(
          "Err, exceeding the max execution count (more information: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/customize-build/"
          "change-script-engine-limitations)\n");
      break;
    }
    EXECUTENUMBER++;
  }
  return TRUE;
}

VOID DebuggerPerformRunTheCustomCode(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGER_EVENT_ACTION *Action,
    DEBUGGER_TRIGGERED_EVENT_DETAILS *EventTriggerDetail) {
  if (Action->CustomCodeBufferSize == 0) {
    return;
  }
  if (Action->RequestedBuffer.RequestBufferSize == 0) {
    AsmDebuggerCustomCodeHandler((UINT64)NULL, (UINT64)DbgState->Regs,
                                 (UINT64)EventTriggerDetail->Context,
                                 (UINT64)Action->CustomCodeBufferAddress);
  } else {
    AsmDebuggerCustomCodeHandler(
        (UINT64)Action->RequestedBuffer.RequstBufferAddress,
        (UINT64)DbgState->Regs, (UINT64)EventTriggerDetail->Context,
        (UINT64)Action->CustomCodeBufferAddress);
  }
}

VOID DebuggerPerformBreakToDebugger(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGER_EVENT_ACTION *Action,
    DEBUGGER_TRIGGERED_EVENT_DETAILS *EventTriggerDetail) {
  UNREFERENCED_PARAMETER(Action);
  if (VmFuncVmxGetCurrentExecutionMode() == TRUE) {
    KdHandleBreakpointAndDebugBreakpoints(
        DbgState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED,
        EventTriggerDetail);
  } else {
    VmFuncVmxVmcall(
        DEBUGGER_VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT,
        (UINT64)EventTriggerDetail, (UINT64)DbgState->Regs, (UINT64)NULL);
  }
}

PDEBUGGER_EVENT DebuggerGetEventByTag(UINT64 Tag) {
  PLIST_ENTRY TempList = 0;
  PLIST_ENTRY TempList2 = 0;
  for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY);
       i++) {
    TempList = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
    TempList2 = TempList;
    while (TempList2 != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag == Tag) {
        return CurrentEvent;
      }
    }
  }
  return NULL;
}

BOOLEAN DebuggerEnableOrDisableAllEvents(BOOLEAN IsEnable) {
  BOOLEAN FindAtLeastOneEvent = FALSE;
  PLIST_ENTRY TempList = 0;
  PLIST_ENTRY TempList2 = 0;
  for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY);
       i++) {
    TempList = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
    TempList2 = TempList;
    while (TempList2 != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (!FindAtLeastOneEvent) {
        FindAtLeastOneEvent = TRUE;
      }
      if (IsEnable) {
        DebuggerEnableEvent(CurrentEvent->Tag);
      } else {
        DebuggerDisableEvent(CurrentEvent->Tag);
      }
    }
  }
  return FindAtLeastOneEvent;
}

BOOLEAN DebuggerTerminateAllEvents(BOOLEAN InputFromVmxRoot) {
  BOOLEAN FindAtLeastOneEvent = FALSE;
  PLIST_ENTRY TempList = 0;
  PLIST_ENTRY TempList2 = 0;
  for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY);
       i++) {
    TempList = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
    TempList2 = TempList;
    while (TempList2 != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (!FindAtLeastOneEvent) {
        FindAtLeastOneEvent = TRUE;
      }
      DebuggerTerminateEvent(CurrentEvent->Tag, InputFromVmxRoot);
    }
  }
  return FindAtLeastOneEvent;
}

BOOLEAN DebuggerRemoveAllEvents(BOOLEAN PoolManagerAllocatedMemory) {
  BOOLEAN FindAtLeastOneEvent = FALSE;
  PLIST_ENTRY TempList = 0;
  PLIST_ENTRY TempList2 = 0;
  for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY);
       i++) {
    TempList = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
    TempList2 = TempList;
    while (TempList2 != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (!FindAtLeastOneEvent) {
        FindAtLeastOneEvent = TRUE;
      }
      DebuggerRemoveEvent(CurrentEvent->Tag, PoolManagerAllocatedMemory);
    }
  }
  return FindAtLeastOneEvent;
}

UINT32 DebuggerEventListCount(PLIST_ENTRY TargetEventList) {
  PLIST_ENTRY TempList = 0;
  UINT32 Counter = 0;
  TempList = TargetEventList;
  while (TargetEventList != TempList->Flink) {
    TempList = TempList->Flink;

    Counter++;
  }
  return Counter;
}

PLIST_ENTRY DebuggerGetEventListByEventType(VMM_EVENT_TYPE_ENUM EventType) {
  PLIST_ENTRY ResultList = NULL;
  switch (EventType) {
    case HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE:
      ResultList = &g_Events->HiddenHookReadAndWriteAndExecuteEventsHead;
      break;
    case HIDDEN_HOOK_READ_AND_WRITE:
      ResultList = &g_Events->HiddenHookReadAndWriteEventsHead;
      break;
    case HIDDEN_HOOK_READ_AND_EXECUTE:
      ResultList = &g_Events->HiddenHookReadAndExecuteEventsHead;
      break;
    case HIDDEN_HOOK_WRITE_AND_EXECUTE:
      ResultList = &g_Events->HiddenHookWriteAndExecuteEventsHead;
      break;
    case HIDDEN_HOOK_READ:
      ResultList = &g_Events->HiddenHookReadEventsHead;
      break;
    case HIDDEN_HOOK_WRITE:
      ResultList = &g_Events->HiddenHookWriteEventsHead;
      break;
    case HIDDEN_HOOK_EXECUTE:
      ResultList = &g_Events->HiddenHookExecuteEventsHead;
      break;
    case HIDDEN_HOOK_EXEC_DETOURS:
      ResultList = &g_Events->EptHook2sExecDetourEventsHead;
      break;
    case HIDDEN_HOOK_EXEC_CC:
      ResultList = &g_Events->EptHookExecCcEventsHead;
      break;
    case SYSCALL_HOOK_EFER_SYSCALL:
      ResultList = &g_Events->SyscallHooksEferSyscallEventsHead;
      break;
    case SYSCALL_HOOK_EFER_SYSRET:
      ResultList = &g_Events->SyscallHooksEferSysretEventsHead;
      break;
    case CPUID_INSTRUCTION_EXECUTION:
      ResultList = &g_Events->CpuidInstructionExecutionEventsHead;
      break;
    case RDMSR_INSTRUCTION_EXECUTION:
      ResultList = &g_Events->RdmsrInstructionExecutionEventsHead;
      break;
    case WRMSR_INSTRUCTION_EXECUTION:
      ResultList = &g_Events->WrmsrInstructionExecutionEventsHead;
      break;
    case EXCEPTION_OCCURRED:
      ResultList = &g_Events->ExceptionOccurredEventsHead;
      break;
    case TSC_INSTRUCTION_EXECUTION:
      ResultList = &g_Events->TscInstructionExecutionEventsHead;
      break;
    case PMC_INSTRUCTION_EXECUTION:
      ResultList = &g_Events->PmcInstructionExecutionEventsHead;
      break;
    case IN_INSTRUCTION_EXECUTION:
      ResultList = &g_Events->InInstructionExecutionEventsHead;
      break;
    case OUT_INSTRUCTION_EXECUTION:
      ResultList = &g_Events->OutInstructionExecutionEventsHead;
      break;
    case DEBUG_REGISTERS_ACCESSED:
      ResultList = &g_Events->DebugRegistersAccessedEventsHead;
      break;
    case EXTERNAL_INTERRUPT_OCCURRED:
      ResultList = &g_Events->ExternalInterruptOccurredEventsHead;
      break;
    case VMCALL_INSTRUCTION_EXECUTION:
      ResultList = &g_Events->VmcallInstructionExecutionEventsHead;
      break;
    case TRAP_EXECUTION_MODE_CHANGED:
      ResultList = &g_Events->TrapExecutionModeChangedEventsHead;
      break;
    case TRAP_EXECUTION_INSTRUCTION_TRACE:
      ResultList = &g_Events->TrapExecutionInstructionTraceEventsHead;
      break;
    case CONTROL_REGISTER_3_MODIFIED:
      ResultList = &g_Events->ControlRegister3ModifiedEventsHead;
      break;
    case CONTROL_REGISTER_MODIFIED:
      ResultList = &g_Events->ControlRegisterModifiedEventsHead;
      break;
    default:
      LogError("Err, wrong event type is specified");
      ResultList = NULL;
      break;
  }
  return ResultList;
}

UINT32 DebuggerEventListCountByCore(PLIST_ENTRY TargetEventList,
                                    UINT32 TargetCore) {
  PLIST_ENTRY TempList = 0;
  UINT32 Counter = 0;
  TempList = TargetEventList;
  while (TargetEventList != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGER_EVENT CurrentEvent =
        CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
    if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES ||
        CurrentEvent->CoreId == TargetCore) {
      Counter++;
    }
  }
  return Counter;
}

UINT32 DebuggerEventListCountByEventType(VMM_EVENT_TYPE_ENUM EventType,
                                         UINT32 TargetCore) {
  PLIST_ENTRY TempList = 0;
  UINT32 Counter = 0;
  PLIST_ENTRY TargetEventList = DebuggerGetEventListByEventType(EventType);
  TempList = TargetEventList;
  while (TargetEventList != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGER_EVENT CurrentEvent =
        CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
    if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES ||
        CurrentEvent->CoreId == TargetCore) {
      Counter++;
    }
  }
  return Counter;
}

UINT32 DebuggerExceptionEventBitmapMask(UINT32 CoreIndex) {
  PLIST_ENTRY TempList = 0;
  UINT32 ExceptionMask = 0;
  TempList = &g_Events->ExceptionOccurredEventsHead;
  while (&g_Events->ExceptionOccurredEventsHead != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGER_EVENT CurrentEvent =
        CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
    if (CurrentEvent->CoreId == DEBUGGER_EVENT_APPLY_TO_ALL_CORES ||
        CurrentEvent->CoreId == CoreIndex) {
      ExceptionMask |= CurrentEvent->Options.OptionalParam1;
    }
  }
  return ExceptionMask;
}

BOOLEAN DebuggerEnableEvent(UINT64 Tag) {
  PDEBUGGER_EVENT Event;
  Event = DebuggerGetEventByTag(Tag);
  if (Event == NULL) {
    return FALSE;
  }
  Event->Enabled = TRUE;
  return TRUE;
}

BOOLEAN DebuggerQueryStateEvent(UINT64 Tag) {
  PDEBUGGER_EVENT Event;
  Event = DebuggerGetEventByTag(Tag);
  if (Event == NULL) {
    return FALSE;
  }
  return Event->Enabled;
}

BOOLEAN DebuggerDisableEvent(UINT64 Tag) {
  PDEBUGGER_EVENT Event;
  Event = DebuggerGetEventByTag(Tag);
  if (Event == NULL) {
    return FALSE;
  }
  Event->Enabled = FALSE;
  return TRUE;
}

BOOLEAN DebuggerClearEvent(UINT64 Tag, BOOLEAN InputFromVmxRoot,
                           BOOLEAN PoolManagerAllocatedMemory) {
  DebuggerDisableEvent(Tag);
  DebuggerTerminateEvent(Tag, InputFromVmxRoot);
  return DebuggerRemoveEvent(Tag, PoolManagerAllocatedMemory);
}

VOID DebuggerClearAllEvents(BOOLEAN InputFromVmxRoot,
                            BOOLEAN PoolManagerAllocatedMemory) {
  DebuggerEnableOrDisableAllEvents(FALSE);
  DebuggerTerminateAllEvents(InputFromVmxRoot);
  DebuggerRemoveAllEvents(PoolManagerAllocatedMemory);
}

BOOLEAN DebuggerIsTagValid(UINT64 Tag) {
  PDEBUGGER_EVENT Event;
  Event = DebuggerGetEventByTag(Tag);
  if (Event == NULL) {
    return FALSE;
  }
  return TRUE;
}

BOOLEAN DebuggerQueryDebuggerStatus() {
  if (g_KernelDebuggerState || g_UserDebuggerState) {
    return TRUE;
  } else {
    return FALSE;
  }
}

BOOLEAN DebuggerRemoveEventFromEventList(UINT64 Tag) {
  PLIST_ENTRY TempList = 0;
  PLIST_ENTRY TempList2 = 0;
  for (size_t i = 0; i < sizeof(DEBUGGER_CORE_EVENTS) / sizeof(LIST_ENTRY);
       i++) {
    TempList = (PLIST_ENTRY)((UINT64)(g_Events) + (i * sizeof(LIST_ENTRY)));
    TempList2 = TempList;
    while (TempList2 != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT CurrentEvent =
          CONTAINING_RECORD(TempList, DEBUGGER_EVENT, EventsOfSameTypeList);
      if (CurrentEvent->Tag == Tag) {
        RemoveEntryList(&CurrentEvent->EventsOfSameTypeList);
        return TRUE;
      }
    }
  }
  return FALSE;
}

BOOLEAN DebuggerRemoveAllActionsFromEvent(PDEBUGGER_EVENT Event,
                                          BOOLEAN PoolManagerAllocatedMemory) {
  PLIST_ENTRY TempList = 0;
  PLIST_ENTRY TempList2 = 0;
  TempList = &Event->ActionsListHead;
  TempList2 = TempList;
  while (TempList2 != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGER_EVENT_ACTION CurrentAction =
        CONTAINING_RECORD(TempList, DEBUGGER_EVENT_ACTION, ActionsList);
    if (CurrentAction->RequestedBuffer.RequestBufferSize != 0 &&
        CurrentAction->RequestedBuffer.RequstBufferAddress != (UINT64)NULL) {
      if (PoolManagerAllocatedMemory) {
        PoolManagerFreePool(CurrentAction->RequestedBuffer.RequstBufferAddress);
      } else {
        PlatformMemFreePool(
            (PVOID)CurrentAction->RequestedBuffer.RequstBufferAddress);
      }
    }
    if (PoolManagerAllocatedMemory) {
      PoolManagerFreePool((UINT64)CurrentAction);
    } else {
      PlatformMemFreePool(CurrentAction);
    }
  }
  return TRUE;
}

BOOLEAN DebuggerRemoveEvent(UINT64 Tag, BOOLEAN PoolManagerAllocatedMemory) {
  PDEBUGGER_EVENT Event;
  if (!DebuggerDisableEvent(Tag)) {
    return FALSE;
  }
  Event = DebuggerGetEventByTag(Tag);
  if (!DebuggerRemoveEventFromEventList(Tag)) {
    return FALSE;
  }
  DebuggerRemoveAllActionsFromEvent(Event, PoolManagerAllocatedMemory);
  if (PoolManagerAllocatedMemory) {
    PoolManagerFreePool((UINT64)Event);
  } else {
    PlatformMemFreePool(Event);
  }
  return TRUE;
}

BOOLEAN DebuggerValidateEvent(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
                              PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                              BOOLEAN InputFromVmxRoot) {
  if ((EventDetails->EventStage ==
           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION ||
       EventDetails->EventStage ==
           VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION) &&
      EventDetails->EnableShortCircuiting == TRUE) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error =
        DEBUGGER_ERROR_USING_SHORT_CIRCUITING_EVENT_WITH_POST_EVENT_MODE_IS_FORBIDDEDN;
    return FALSE;
  }
  if (EventDetails->CoreId != DEBUGGER_EVENT_APPLY_TO_ALL_CORES) {
    if (!CommonValidateCoreNumber(EventDetails->CoreId)) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_CORE_ID;
      return FALSE;
    }
  }
  if (EventDetails->ProcessId != DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES &&
      EventDetails->ProcessId != 0) {
    if (!InputFromVmxRoot) {
      if (!CommonIsProcessExist(EventDetails->ProcessId)) {
        ResultsToReturn->IsSuccessful = FALSE;
        ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_PROCESS_ID;
        return FALSE;
      }
    }
  }
  switch (EventDetails->EventType) {
    case EXCEPTION_OCCURRED: {
      if (!ValidateEventException(EventDetails, ResultsToReturn,
                                  InputFromVmxRoot)) {
        return FALSE;
      }
      break;
    }
    case EXTERNAL_INTERRUPT_OCCURRED: {
      if (!ValidateEventInterrupt(EventDetails, ResultsToReturn,
                                  InputFromVmxRoot)) {
        return FALSE;
      }
      break;
    }
    case TRAP_EXECUTION_MODE_CHANGED: {
      if (!ValidateEventTrapExec(EventDetails, ResultsToReturn,
                                 InputFromVmxRoot)) {
        return FALSE;
      }
      break;
    }
    case HIDDEN_HOOK_EXEC_DETOURS:
    case HIDDEN_HOOK_EXEC_CC: {
      if (!ValidateEventEptHookHiddenBreakpointAndInlineHooks(
              EventDetails, ResultsToReturn, InputFromVmxRoot)) {
        return FALSE;
      }
      break;
    }
    case HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE:
    case HIDDEN_HOOK_READ_AND_WRITE:
    case HIDDEN_HOOK_READ_AND_EXECUTE:
    case HIDDEN_HOOK_WRITE_AND_EXECUTE:
    case HIDDEN_HOOK_READ:
    case HIDDEN_HOOK_WRITE:
    case HIDDEN_HOOK_EXECUTE: {
      if (!ValidateEventMonitor(EventDetails, ResultsToReturn,
                                InputFromVmxRoot)) {
        return FALSE;
      }
      break;
    }
    default:
      break;
  }
  return TRUE;
}

BOOLEAN DebuggerApplyEvent(PDEBUGGER_EVENT Event,
                           PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                           BOOLEAN InputFromVmxRoot) {
  switch (Event->EventType) {
    case HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE:
    case HIDDEN_HOOK_READ_AND_WRITE:
    case HIDDEN_HOOK_READ_AND_EXECUTE:
    case HIDDEN_HOOK_WRITE_AND_EXECUTE:
    case HIDDEN_HOOK_READ:
    case HIDDEN_HOOK_WRITE:
    case HIDDEN_HOOK_EXECUTE: {
      if (!ApplyEventMonitorEvent(Event, ResultsToReturn, InputFromVmxRoot)) {
        goto ClearTheEventAfterCreatingEvent;
      }
      break;
    }
    case HIDDEN_HOOK_EXEC_CC: {
      if (!ApplyEventEptHookExecCcEvent(Event, ResultsToReturn,
                                        InputFromVmxRoot)) {
        goto ClearTheEventAfterCreatingEvent;
      }
      break;
    }
    case HIDDEN_HOOK_EXEC_DETOURS: {
      if (!ApplyEventEpthookInlineEvent(Event, ResultsToReturn,
                                        InputFromVmxRoot)) {
        goto ClearTheEventAfterCreatingEvent;
      }
      break;
    }
    case RDMSR_INSTRUCTION_EXECUTION: {
      ApplyEventRdmsrExecutionEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case WRMSR_INSTRUCTION_EXECUTION: {
      ApplyEventWrmsrExecutionEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case IN_INSTRUCTION_EXECUTION:
    case OUT_INSTRUCTION_EXECUTION: {
      ApplyEventInOutExecutionEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case TSC_INSTRUCTION_EXECUTION: {
      ApplyEventTscExecutionEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case PMC_INSTRUCTION_EXECUTION: {
      ApplyEventRdpmcExecutionEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case DEBUG_REGISTERS_ACCESSED: {
      ApplyEventMov2DebugRegExecutionEvent(Event, ResultsToReturn,
                                           InputFromVmxRoot);
      break;
    }
    case CONTROL_REGISTER_MODIFIED: {
      ApplyEventControlRegisterAccessedEvent(Event, ResultsToReturn,
                                             InputFromVmxRoot);
      break;
    }
    case EXCEPTION_OCCURRED: {
      ApplyEventExceptionEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case EXTERNAL_INTERRUPT_OCCURRED: {
      ApplyEventInterruptEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case SYSCALL_HOOK_EFER_SYSCALL: {
      ApplyEventEferSyscallHookEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case SYSCALL_HOOK_EFER_SYSRET: {
      ApplyEventEferSysretHookEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case VMCALL_INSTRUCTION_EXECUTION: {
      ApplyEventVmcallExecutionEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case TRAP_EXECUTION_MODE_CHANGED: {
      if (!ApplyEventTrapModeChangeEvent(Event, ResultsToReturn,
                                         InputFromVmxRoot)) {
        goto ClearTheEventAfterCreatingEvent;
      }
      break;
    }
    case CPUID_INSTRUCTION_EXECUTION: {
      ApplyEventCpuidExecutionEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    case TRAP_EXECUTION_INSTRUCTION_TRACE: {
      ApplyEventTracingEvent(Event, ResultsToReturn, InputFromVmxRoot);
      break;
    }
    default: {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID;
      goto ClearTheEventAfterCreatingEvent;
      break;
    }
  }
  ResultsToReturn->IsSuccessful = TRUE;
  ResultsToReturn->Error = 0;
  return TRUE;
ClearTheEventAfterCreatingEvent:
  return FALSE;
}

BOOLEAN DebuggerParseEvent(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
                           PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                           BOOLEAN InputFromVmxRoot) {
  PDEBUGGER_EVENT Event;
  if (!DebuggerValidateEvent(EventDetails, ResultsToReturn, InputFromVmxRoot)) {
    return FALSE;
  }
  if (EventDetails->ConditionBufferSize != 0) {
    Event = DebuggerCreateEvent(
        FALSE, EventDetails->CoreId, EventDetails->ProcessId,
        EventDetails->EventType, EventDetails->Tag, &EventDetails->Options,
        EventDetails->ConditionBufferSize,
        (PVOID)((UINT64)EventDetails + sizeof(DEBUGGER_GENERAL_EVENT_DETAIL)),
        ResultsToReturn, InputFromVmxRoot);
  } else {
    Event = DebuggerCreateEvent(
        FALSE, EventDetails->CoreId, EventDetails->ProcessId,
        EventDetails->EventType, EventDetails->Tag, &EventDetails->Options, 0,
        NULL, ResultsToReturn, InputFromVmxRoot);
  }
  if (Event == NULL) {
    return FALSE;
  }
  DebuggerRegisterEvent(Event);
  if (DebuggerApplyEvent(Event, ResultsToReturn, InputFromVmxRoot)) {
    Event->EnableShortCircuiting = EventDetails->EnableShortCircuiting;
    if (EventDetails->EventStage ==
        VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION) {
      Event->EventMode = VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION;
    } else if (EventDetails->EventStage ==
               VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION) {
      Event->EventMode = VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION;
    } else {
      Event->EventMode = VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION;
    }
    return TRUE;
  } else {
    if (Event != NULL) {
      DebuggerRemoveEvent(Event->Tag, InputFromVmxRoot);
    }
    return FALSE;
  }
}

BOOLEAN DebuggerParseAction(PDEBUGGER_GENERAL_ACTION ActionDetails,
                            PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                            BOOLEAN InputFromVmxRoot) {
  DEBUGGER_EVENT_ACTION *Action = NULL;
  PDEBUGGER_EVENT Event = DebuggerGetEventByTag(ActionDetails->EventTag);
  if (Event == NULL) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error = DEBUGGER_ERROR_TAG_NOT_EXISTS;
    return FALSE;
  }
  if (ActionDetails->ActionType == RUN_CUSTOM_CODE) {
    if (ActionDetails->CustomCodeBufferSize == 0) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO;
      return FALSE;
    }
    DEBUGGER_EVENT_REQUEST_CUSTOM_CODE CustomCode = {0};
    CustomCode.CustomCodeBufferSize = ActionDetails->CustomCodeBufferSize;
    CustomCode.CustomCodeBufferAddress =
        (PVOID)((UINT64)ActionDetails + sizeof(DEBUGGER_GENERAL_ACTION));
    CustomCode.OptionalRequestedBufferSize = ActionDetails->PreAllocatedBuffer;
    Action = DebuggerAddActionToEvent(
        Event, RUN_CUSTOM_CODE, ActionDetails->ImmediateMessagePassing,
        &CustomCode, NULL, ResultsToReturn, InputFromVmxRoot);
    if (!Action) {
      return FALSE;
    }
  } else if (ActionDetails->ActionType == RUN_SCRIPT) {
    if (ActionDetails->ScriptBufferSize == 0) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO;
      return FALSE;
    }
    DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION UserScriptConfig = {0};
    UserScriptConfig.ScriptBuffer =
        (UINT64)ActionDetails + sizeof(DEBUGGER_GENERAL_ACTION);
    UserScriptConfig.ScriptLength = ActionDetails->ScriptBufferSize;
    UserScriptConfig.ScriptPointer = ActionDetails->ScriptBufferPointer;
    UserScriptConfig.OptionalRequestedBufferSize =
        ActionDetails->PreAllocatedBuffer;
    Action = DebuggerAddActionToEvent(
        Event, RUN_SCRIPT, ActionDetails->ImmediateMessagePassing, NULL,
        &UserScriptConfig, ResultsToReturn, InputFromVmxRoot);
    if (!Action) {
      return FALSE;
    }
  } else if (ActionDetails->ActionType == BREAK_TO_DEBUGGER) {
    Action = DebuggerAddActionToEvent(
        Event, BREAK_TO_DEBUGGER, ActionDetails->ImmediateMessagePassing, NULL,
        NULL, ResultsToReturn, InputFromVmxRoot);
    if (!Action) {
      return FALSE;
    }
  } else {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_ACTION_TYPE;
    return FALSE;
  }
  DebuggerEnableEvent(Event->Tag);
  ResultsToReturn->IsSuccessful = TRUE;
  ResultsToReturn->Error = 0;
  return TRUE;
}

BOOLEAN DebuggerTerminateEvent(UINT64 Tag, BOOLEAN InputFromVmxRoot) {
  PDEBUGGER_EVENT Event;
  BOOLEAN Result = FALSE;
  Event = DebuggerGetEventByTag(Tag);
  if (Event == NULL) {
    return FALSE;
  }
  switch (Event->EventType) {
    case EXTERNAL_INTERRUPT_OCCURRED: {
      TerminateExternalInterruptEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE:
    case HIDDEN_HOOK_READ_AND_WRITE:
    case HIDDEN_HOOK_READ_AND_EXECUTE:
    case HIDDEN_HOOK_WRITE_AND_EXECUTE:
    case HIDDEN_HOOK_READ:
    case HIDDEN_HOOK_WRITE:
    case HIDDEN_HOOK_EXECUTE: {
      TerminateHiddenHookReadAndWriteAndExecuteEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case HIDDEN_HOOK_EXEC_CC: {
      TerminateHiddenHookExecCcEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case HIDDEN_HOOK_EXEC_DETOURS: {
      TerminateHiddenHookExecDetoursEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case RDMSR_INSTRUCTION_EXECUTION: {
      TerminateRdmsrExecutionEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case WRMSR_INSTRUCTION_EXECUTION: {
      TerminateWrmsrExecutionEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case EXCEPTION_OCCURRED: {
      TerminateExceptionEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case IN_INSTRUCTION_EXECUTION: {
      TerminateInInstructionExecutionEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case OUT_INSTRUCTION_EXECUTION: {
      TerminateOutInstructionExecutionEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case SYSCALL_HOOK_EFER_SYSCALL: {
      TerminateSyscallHookEferEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case SYSCALL_HOOK_EFER_SYSRET: {
      TerminateSysretHookEferEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case VMCALL_INSTRUCTION_EXECUTION: {
      TerminateVmcallExecutionEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case TRAP_EXECUTION_MODE_CHANGED: {
      TerminateExecTrapModeChangedEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case TSC_INSTRUCTION_EXECUTION: {
      TerminateTscEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case PMC_INSTRUCTION_EXECUTION: {
      TerminatePmcEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case DEBUG_REGISTERS_ACCESSED: {
      TerminateDebugRegistersEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case CPUID_INSTRUCTION_EXECUTION: {
      TerminateCpuidExecutionEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    case CONTROL_REGISTER_MODIFIED: {
      TerminateControlRegistersEvent(Event, InputFromVmxRoot);
      Result = TRUE;
      break;
    }
    default:
      LogError("Err, unknown event for termination");
      Result = FALSE;
      break;
  }
  return Result;
}

BOOLEAN DebuggerParseEventsModification(
    PDEBUGGER_MODIFY_EVENTS DebuggerEventModificationRequest,
    BOOLEAN InputFromVmxRoot, BOOLEAN PoolManagerAllocatedMemory) {
  BOOLEAN IsForAllEvents = FALSE;
  if (DebuggerEventModificationRequest->Tag ==
      DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG) {
    IsForAllEvents = TRUE;
  } else if (!DebuggerIsTagValid(DebuggerEventModificationRequest->Tag)) {
    DebuggerEventModificationRequest->KernelStatus =
        DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG;
    return FALSE;
  }
  if (DebuggerEventModificationRequest->TypeOfAction ==
      DEBUGGER_MODIFY_EVENTS_ENABLE) {
    if (IsForAllEvents) {
      DebuggerEnableOrDisableAllEvents(TRUE);
    } else {
      DebuggerEnableEvent(DebuggerEventModificationRequest->Tag);
    }
  } else if (DebuggerEventModificationRequest->TypeOfAction ==
             DEBUGGER_MODIFY_EVENTS_DISABLE) {
    if (IsForAllEvents) {
      DebuggerEnableOrDisableAllEvents(FALSE);
    } else {
      DebuggerDisableEvent(DebuggerEventModificationRequest->Tag);
    }
  } else if (DebuggerEventModificationRequest->TypeOfAction ==
             DEBUGGER_MODIFY_EVENTS_CLEAR) {
    if (IsForAllEvents) {
      DebuggerClearAllEvents(InputFromVmxRoot, PoolManagerAllocatedMemory);
    } else {
      DebuggerClearEvent(DebuggerEventModificationRequest->Tag,
                         InputFromVmxRoot, PoolManagerAllocatedMemory);
    }
  } else if (DebuggerEventModificationRequest->TypeOfAction ==
             DEBUGGER_MODIFY_EVENTS_QUERY_STATE) {
    if (!DebuggerIsTagValid(DebuggerEventModificationRequest->Tag)) {
      DebuggerEventModificationRequest->KernelStatus =
          DEBUGGER_ERROR_TAG_NOT_EXISTS;
      return FALSE;
    }
    if (DebuggerQueryStateEvent(DebuggerEventModificationRequest->Tag)) {
      DebuggerEventModificationRequest->IsEnabled = TRUE;
    } else {
      DebuggerEventModificationRequest->IsEnabled = FALSE;
    }
  } else {
    DebuggerEventModificationRequest->KernelStatus =
        DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION;
    return FALSE;
  }
  DebuggerEventModificationRequest->KernelStatus =
      DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return TRUE;
}
