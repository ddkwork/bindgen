
#include "pch.h"

VOID KdInitializeKernelDebugger() {
  PoolManagerRequestAllocation(sizeof(DEBUGGEE_BP_DESCRIPTOR),
                               MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE,
                               BREAKPOINT_DEFINITION_STRUCTURE);
  BroadcastEnableDbAndBpExitingAllCores();
  RtlZeroMemory(&g_IgnoreBreaksToDebugger,
                sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));
  g_MaximumBreakpointId = 0;
  InitializeListHead(&g_BreakpointsListHead);
  KdInitializeInstantEventPools();
  g_KernelDebuggerState = TRUE;
}

VOID KdUninitializeKernelDebugger() {
  ULONG ProcessorsCount;
  if (g_KernelDebuggerState) {
    ProcessorsCount = KeQueryActiveProcessorCount(0);
    g_KernelDebuggerState = FALSE;
    RtlZeroMemory(&g_IgnoreBreaksToDebugger,
                  sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));
    BreakpointRemoveAllBreakpoints();
    BroadcastDisableDbAndBpExitingAllCores();
  }
}

BOOLEAN KdCheckImmediateMessagingMechanism(UINT32 OperationCode) {
  return (g_KernelDebuggerState &&
          !(OperationCode & OPERATION_MANDATORY_DEBUGGEE_BIT));
}

VOID KdInitializeInstantEventPools() {
  PoolManagerRequestAllocation(REGULAR_INSTANT_EVENT_CONDITIONAL_BUFFER,
                               MAXIMUM_REGULAR_INSTANT_EVENTS,
                               INSTANT_REGULAR_EVENT_BUFFER);
  PoolManagerRequestAllocation(REGULAR_INSTANT_EVENT_ACTION_BUFFER,
                               MAXIMUM_REGULAR_INSTANT_EVENTS,
                               INSTANT_REGULAR_EVENT_ACTION_BUFFER);
#if MAXIMUM_BIG_INSTANT_EVENTS >= \
    1 PoolManagerRequestAllocation(BIG_INSTANT_EVENT_CONDITIONAL_BUFFER,
                               MAXIMUM_BIG_INSTANT_EVENTS,
                               INSTANT_BIG_EVENT_BUFFER);
                               PoolManagerRequestAllocation(
                                   BIG_INSTANT_EVENT_ACTION_BUFFER,
                                   MAXIMUM_BIG_INSTANT_EVENTS,
                                   INSTANT_BIG_EVENT_ACTION_BUFFER);
#endif ConfigureEptHookReservePreallocatedPoolsForEptHooks(
      MAXIMUM_REGULAR_INSTANT_EVENTS -
      MAXIMUM_NUMBER_OF_INITIAL_PREALLOCATED_EPT_HOOKS);
      if (!PoolManagerCheckAndPerformAllocationAndDeallocation()) {
        LogWarning(
            "Warning, cannot allocate the pre-allocated pools for EPT hooks");
      }
}

VOID KdDummyDPC(PKDPC Dpc, PVOID DeferredContext, PVOID SystemArgument1,
                PVOID SystemArgument2) {
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  LogInfo("I'm here %llx\n", DeferredContext);
}

VOID KdFireDpc(PVOID Routine, PVOID Parameter) {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  KeInitializeDpc(g_DbgState[CurrentCore].KdDpcObject,
                  (PKDEFERRED_ROUTINE)Routine, Parameter);
  KeInsertQueueDpc(g_DbgState[CurrentCore].KdDpcObject, NULL, NULL);
}

BOOLEAN KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId(
    UINT32 CoreId, DEBUGGER_THREAD_PROCESS_TRACING TracingType) {
  BOOLEAN Result = FALSE;
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  switch (TracingType) {
    case DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_THREAD_CHANGE:
      Result = DbgState->ThreadOrProcessTracingDetails
                   .InterceptClockInterruptsForThreadChange;
      break;
    case DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_PROCESS_CHANGE:
      Result = DbgState->ThreadOrProcessTracingDetails
                   .InterceptClockInterruptsForProcessChange;
      break;
    case DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_DEBUG_REGISTER_INTERCEPTION:
      Result = DbgState->ThreadOrProcessTracingDetails
                   .DebugRegisterInterceptionState;
      break;
    case DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_WAITING_FOR_MOV_CR3_VM_EXITS:
      Result = DbgState->ThreadOrProcessTracingDetails.IsWatingForMovCr3VmExits;
      break;
    default:
      LogError(
          "Err, debugger encountered an unknown query type for querying "
          "process or thread interception details");
      break;
  }
  return Result;
}

_Use_decl_annotations_

    BYTE
    KdComputeDataChecksum(PVOID Buffer, UINT32 Length) {
  BYTE CalculatedCheckSum = 0;
  BYTE Temp = 0;
  while (Length--) {
    Temp = *(BYTE *)Buffer;
    CalculatedCheckSum = CalculatedCheckSum + Temp;
    Buffer = (PVOID)((UINT64)Buffer + 1);
  }
  return CalculatedCheckSum;
}

_Use_decl_annotations_

    BOOLEAN
    KdResponsePacketToDebugger(DEBUGGER_REMOTE_PACKET_TYPE PacketType,
                               DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION Response,
                               CHAR *OptionalBuffer,
                               UINT32 OptionalBufferLength) {
  DEBUGGER_REMOTE_PACKET Packet = {0};
  BOOLEAN Result = FALSE;
  Packet.Indicator = INDICATOR_OF_HYPERDBG_PACKET;
  Packet.TypeOfThePacket = PacketType;
  Packet.RequestedActionOfThePacket = Response;
  if (OptionalBuffer == NULL || OptionalBufferLength == 0) {
    Packet.Checksum =
        KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
                              sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
    ScopedSpinlock(DebuggerResponseLock,
                   Result = SerialConnectionSend(
                       (CHAR *)&Packet, sizeof(DEBUGGER_REMOTE_PACKET)));
  } else {
    Packet.Checksum =
        KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
                              sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
    Packet.Checksum +=
        KdComputeDataChecksum((PVOID)OptionalBuffer, OptionalBufferLength);
    ScopedSpinlock(DebuggerResponseLock,
                   Result = SerialConnectionSendTwoBuffers(
                       (CHAR *)&Packet, sizeof(DEBUGGER_REMOTE_PACKET),
                       OptionalBuffer, OptionalBufferLength));
  }
  if (g_IgnoreBreaksToDebugger.PauseBreaksUntilSpecialMessageSent &&
      g_IgnoreBreaksToDebugger.SpeialEventResponse == Response) {
    RtlZeroMemory(&g_IgnoreBreaksToDebugger,
                  sizeof(DEBUGGEE_REQUEST_TO_IGNORE_BREAKS_UNTIL_AN_EVENT));
  }
  return Result;
}

_Use_decl_annotations_

    BOOLEAN
    KdLoggingResponsePacketToDebugger(CHAR *OptionalBuffer,
                                      UINT32 OptionalBufferLength,
                                      UINT32 OperationCode) {
  DEBUGGER_REMOTE_PACKET Packet = {0};
  BOOLEAN Result = FALSE;
  Packet.Indicator = INDICATOR_OF_HYPERDBG_PACKET;
  Packet.TypeOfThePacket = DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER;
  Packet.RequestedActionOfThePacket =
      DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM;
  Packet.Checksum =
      KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
                            sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
  Packet.Checksum +=
      KdComputeDataChecksum((PVOID)&OperationCode, sizeof(UINT32));
  Packet.Checksum +=
      KdComputeDataChecksum((PVOID)OptionalBuffer, OptionalBufferLength);
  ScopedSpinlock(DebuggerResponseLock,
                 Result = SerialConnectionSendThreeBuffers(
                     (CHAR *)&Packet, sizeof(DEBUGGER_REMOTE_PACKET),
                     (CHAR *)&OperationCode, sizeof(UINT32), OptionalBuffer,
                     OptionalBufferLength));
  return Result;
}

VOID KdHandleDebugEventsWhenKernelDebuggerIsAttached(
    PROCESSOR_DEBUGGING_STATE *DbgState, BOOLEAN TrapSetByDebugger) {
  DEBUGGER_TRIGGERED_EVENT_DETAILS TargetContext = {0};
  BOOLEAN IgnoreDebugEvent = FALSE;
  UINT64 LastVmexitRip = VmFuncGetLastVmexitRip(DbgState->CoreId);
  TargetContext.Context = (PVOID)LastVmexitRip;
  if (TrapSetByDebugger) {
    if (!BreakpointCheckAndHandleDebuggerDefinedBreakpoints(
            DbgState, LastVmexitRip, DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
            FALSE)) {
      if (g_HardwareDebugRegisterDetailsForStepOver.Address != (UINT64)NULL) {
        if (LastVmexitRip ==
            g_HardwareDebugRegisterDetailsForStepOver.Address) {
          if (g_HardwareDebugRegisterDetailsForStepOver.ProcessId ==
                  HANDLE_TO_UINT32(PsGetCurrentProcessId()) &&
              g_HardwareDebugRegisterDetailsForStepOver.ThreadId ==
                  HANDLE_TO_UINT32(PsGetCurrentThreadId())) {
            RtlZeroMemory(&g_HardwareDebugRegisterDetailsForStepOver,
                          sizeof(HARDWARE_DEBUG_REGISTER_DETAILS));
          } else {
            IgnoreDebugEvent = TRUE;
            SetDebugRegisters(
                DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER,
                BREAK_ON_INSTRUCTION_FETCH, FALSE,
                g_HardwareDebugRegisterDetailsForStepOver.Address);
          }
        }
      }
      if (!IgnoreDebugEvent) {
        KdHandleBreakpointAndDebugBreakpoints(
            DbgState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED, &TargetContext);
      }
    }
  } else {
    KdHandleBreakpointAndDebugBreakpoints(
        DbgState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT,
        &TargetContext);
  }
}

VOID KdApplyTasksPreHaltCore(PROCESSOR_DEBUGGING_STATE *DbgState) {
  if (DbgState->ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent ==
      TRUE) {
    ProcessEnableOrDisableThreadChangeMonitor(
        DbgState, FALSE,
        DbgState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);
  }
  if (DbgState->ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent ==
      TRUE) {
    ThreadEnableOrDisableThreadChangeMonitor(
        DbgState, FALSE,
        DbgState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt);
  }
}

VOID KdApplyTasksPostContinueCore(PROCESSOR_DEBUGGING_STATE *DbgState) {
  if (DbgState->HardwareDebugRegisterForStepping != (UINT64)NULL) {
    SetDebugRegisters(DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER,
                      BREAK_ON_INSTRUCTION_FETCH, FALSE,
                      DbgState->HardwareDebugRegisterForStepping);
    DbgState->HardwareDebugRegisterForStepping = (UINT64)NULL;
  }
}

_Use_decl_annotations_

    VOID
    KdContinueDebuggee(
        PROCESSOR_DEBUGGING_STATE *DbgState,
        BOOLEAN PauseBreaksUntilSpecialMessageSent,
        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION SpeialEventResponse) {
  if (PauseBreaksUntilSpecialMessageSent) {
    g_IgnoreBreaksToDebugger.PauseBreaksUntilSpecialMessageSent = TRUE;
    g_IgnoreBreaksToDebugger.SpeialEventResponse = SpeialEventResponse;
  }
  VmFuncCheckAndEnableExternalInterrupts(DbgState->CoreId);
  ULONG ProcessorsCount = KeQueryActiveProcessorCount(0);
  for (size_t i = 0; i < ProcessorsCount; i++) {
    SpinlockUnlock(&g_DbgState[i].Lock);
  }
}

VOID KdContinueDebuggeeJustCurrentCore(PROCESSOR_DEBUGGING_STATE *DbgState) {
  DbgState->DoNotNmiNotifyOtherCoresByThisCore = TRUE;
  SpinlockUnlock(&DbgState->Lock);
}

_Use_decl_annotations_

    BOOLEAN
    KdReadRegisters(PROCESSOR_DEBUGGING_STATE *DbgState,
                    PDEBUGGEE_REGISTER_READ_DESCRIPTION ReadRegisterRequest) {
  GUEST_EXTRA_REGISTERS ERegs = {0};
  if (ReadRegisterRequest->RegisterId == DEBUGGEE_SHOW_ALL_REGISTERS) {
    memcpy((void *)((CHAR *)ReadRegisterRequest +
                    sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION)),
           DbgState->Regs, sizeof(GUEST_REGS));
    ERegs.CS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_CS);
    ERegs.SS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_SS);
    ERegs.DS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_DS);
    ERegs.ES = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_ES);
    ERegs.FS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_FS);
    ERegs.GS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_GS);
    ERegs.RFLAGS = DebuggerGetRegValueWrapper(NULL, REGISTER_RFLAGS);
    ERegs.RIP = DebuggerGetRegValueWrapper(NULL, REGISTER_RIP);
    memcpy((void *)((CHAR *)ReadRegisterRequest +
                    sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION) +
                    sizeof(GUEST_REGS)),
           &ERegs, sizeof(GUEST_EXTRA_REGISTERS));
  } else {
    ReadRegisterRequest->Value = DebuggerGetRegValueWrapper(
        DbgState->Regs, ReadRegisterRequest->RegisterId);
  }
  return TRUE;
}

_Use_decl_annotations_

    BOOLEAN
    KdReadMemory(PGUEST_REGS Regs,
                 PDEBUGGEE_REGISTER_READ_DESCRIPTION ReadRegisterRequest) {
  GUEST_EXTRA_REGISTERS ERegs = {0};
  if (ReadRegisterRequest->RegisterId == DEBUGGEE_SHOW_ALL_REGISTERS) {
    memcpy((void *)((CHAR *)ReadRegisterRequest +
                    sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION)),
           Regs, sizeof(GUEST_REGS));
    ERegs.CS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_CS);
    ERegs.SS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_SS);
    ERegs.DS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_DS);
    ERegs.ES = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_ES);
    ERegs.FS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_FS);
    ERegs.GS = (UINT16)DebuggerGetRegValueWrapper(NULL, REGISTER_GS);
    ERegs.RFLAGS = DebuggerGetRegValueWrapper(NULL, REGISTER_RFLAGS);
    ERegs.RIP = DebuggerGetRegValueWrapper(NULL, REGISTER_RIP);
    memcpy((void *)((CHAR *)ReadRegisterRequest +
                    sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION) +
                    sizeof(GUEST_REGS)),
           &ERegs, sizeof(GUEST_EXTRA_REGISTERS));
  } else {
    ReadRegisterRequest->Value =
        DebuggerGetRegValueWrapper(Regs, ReadRegisterRequest->RegisterId);
  }
  return TRUE;
}

BOOLEAN KdSwitchCore(PROCESSOR_DEBUGGING_STATE *DbgState,
                     DEBUGGEE_CHANGE_CORE_PACKET *ChangeCorePacket) {
  ULONG ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (DbgState->CoreId == ChangeCorePacket->NewCore) {
    ChangeCorePacket->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
    return FALSE;
  }
  if (ChangeCorePacket->NewCore >= ProcessorsCount) {
    ChangeCorePacket->Result =
        DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE;
    return FALSE;
  }
  if (!KdCheckTargetCoreIsLocked(ChangeCorePacket->NewCore)) {
    ChangeCorePacket->Result =
        DEBUGGER_ERROR_TARGET_SWITCHING_CORE_IS_NOT_LOCKED;
    return FALSE;
  }
  VmFuncCheckAndEnableExternalInterrupts(DbgState->CoreId);
  DbgState->MainDebuggingCore = FALSE;
  g_DbgState[ChangeCorePacket->NewCore].MainDebuggingCore = TRUE;
  ChangeCorePacket->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return TRUE;
}

VOID KdCheckUserModePriorityBuffers() {
  if (LogCallbackCheckIfBufferIsFull(TRUE)) {
    LogWarning(
        "Warning, the user-mode priority buffers are full, thus the new action "
        "replaces "
        "previously unserviced actions. As the result, some functionalities "
        "might not work correctly!\n"
        "For more information please visit: "
        "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n");
  }
}

VOID KdCloseConnectionAndUnloadDebuggee() {
  KdCheckUserModePriorityBuffers();
  LogCallbackSendBuffer(OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM,
                        "$", sizeof(CHAR), TRUE);
}

_Use_decl_annotations_

    VOID
    KdReloadSymbolDetailsInDebuggee(PDEBUGGEE_SYMBOL_REQUEST_PACKET SymPacket) {
  KdCheckUserModePriorityBuffers();
  LogCallbackSendBuffer(OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL,
                        SymPacket, sizeof(DEBUGGEE_SYMBOL_REQUEST_PACKET),
                        TRUE);
}

VOID KdNotifyDebuggeeForUserInput(DEBUGGEE_USER_INPUT_PACKET *Descriptor,
                                  UINT32 Len) {
  KdCheckUserModePriorityBuffers();
  LogCallbackSendBuffer(OPERATION_DEBUGGEE_USER_INPUT, Descriptor, Len, TRUE);
}

VOID KdSendFormatsFunctionResult(UINT64 Value) {
  DEBUGGEE_FORMATS_PACKET FormatsPacket = {0};
  FormatsPacket.Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  FormatsPacket.Value = Value;
  KdResponsePacketToDebugger(
      DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
      DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS,
      (CHAR *)&FormatsPacket, sizeof(DEBUGGEE_FORMATS_PACKET));
}

VOID KdSendCommandFinishedSignal(UINT32 CoreId) {
  KdHandleBreakpointAndDebugBreakpointsCallback(
      CoreId, DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED,
      NULL);
}

_Use_decl_annotations_

    VOID
    KdHandleHaltsWhenNmiReceivedFromVmxRoot(
        PROCESSOR_DEBUGGING_STATE *DbgState) {
  KdHandleNmi(DbgState);
  DbgState->NmiState.NmiCalledInVmxRootRelatedToHaltDebuggee = FALSE;
}

VOID KdCustomDebuggerBreakSpinlockLock(PROCESSOR_DEBUGGING_STATE *DbgState,
                                       volatile LONG *Lock) {
  unsigned wait = 1;
  while (!SpinlockTryLock(Lock)) {
    for (unsigned i = 0; i < wait; ++i) {
      _mm_pause();
    }
    if (DbgState->NmiState.WaitingToBeLocked) {
      VmFuncChangeIgnoreOneMtfState(DbgState->CoreId, TRUE);
      if (DbgState->NmiState.NmiCalledInVmxRootRelatedToHaltDebuggee) {
        KdHandleHaltsWhenNmiReceivedFromVmxRoot(DbgState);
      } else {
        KdHandleNmi(DbgState);
      }
    }
    if (wait * 2 > 65536) {
      wait = 65536;
    } else {
      wait = wait * 2;
    }
  }
}

VOID KdHandleNmiBroadcastDebugBreaks(UINT32 CoreId, BOOLEAN IsOnVmxNmiHandler) {
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  DbgState->NmiState.WaitingToBeLocked = TRUE;
  if (IsOnVmxNmiHandler) {
    DbgState->NmiState.NmiCalledInVmxRootRelatedToHaltDebuggee = TRUE;
    VmFuncSetMonitorTrapFlag(TRUE);
  } else {
    KdHandleNmi(DbgState);
  }
}

_Use_decl_annotations_

    VOID
    KdHandleBreakpointAndDebugBreakpointsCallback(
        UINT32 CoreId, DEBUGGEE_PAUSING_REASON Reason,
        PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails) {
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  KdHandleBreakpointAndDebugBreakpoints(DbgState, Reason, EventDetails);
}

_Use_decl_annotations_

    VOID
    KdHandleRegisteredMtfCallback(UINT32 CoreId) {
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  if (DbgState->TracingMode) {
    TracingHandleMtf(DbgState);
  } else {
    UINT64 CsSel = NULL64_ZERO;
    DEBUGGER_TRIGGERED_EVENT_DETAILS TargetContext = {0};
    UINT64 LastVmexitRip = VmFuncGetLastVmexitRip(CoreId);
    CsSel = VmFuncGetCsSelector();
    KdCheckGuestOperatingModeChanges(DbgState->InstrumentationStepInTrace.CsSel,
                                     (UINT16)CsSel);
    DbgState->InstrumentationStepInTrace.CsSel = 0;
    if (!BreakpointCheckAndHandleDebuggerDefinedBreakpoints(
            DbgState, LastVmexitRip, DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
            TRUE)) {
      TargetContext.Context = (PVOID)LastVmexitRip;
      KdHandleBreakpointAndDebugBreakpoints(
          DbgState,
          DbgState->IgnoreDisasmInNextPacket
              ? DEBUGGEE_PAUSING_REASON_DEBUGGEE_TRACKING_STEPPED
              : DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
          &TargetContext);
    }
  }
}

_Use_decl_annotations_

    VOID
    KdHandleBreakpointAndDebugBreakpoints(
        PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGEE_PAUSING_REASON Reason,
        PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails) {
  KdCustomDebuggerBreakSpinlockLock(DbgState, &DebuggerHandleBreakpointLock);
  if (g_IgnoreBreaksToDebugger.PauseBreaksUntilSpecialMessageSent) {
    SpinlockUnlock(&DebuggerHandleBreakpointLock);
    return;
  }
  DbgState->MainDebuggingCore = TRUE;
  DbgState->NmiState.WaitingToBeLocked = FALSE;
  SpinlockLock(&DbgState->Lock);
  g_DebuggeeHaltReason = Reason;
  if (EventDetails != NULL) {
    RtlCopyMemory(&g_EventTriggerDetail, EventDetails,
                  sizeof(DEBUGGER_TRIGGERED_EVENT_DETAILS));
  }
  if (DbgState->DoNotNmiNotifyOtherCoresByThisCore == TRUE) {
    DbgState->DoNotNmiNotifyOtherCoresByThisCore = FALSE;
  } else {
    SpinlockLock(&DebuggerResponseLock);
    VmFuncNmiBroadcastRequest(DbgState->CoreId);
    SpinlockUnlock(&DebuggerResponseLock);
  }
  KdManageSystemHaltOnVmxRoot(DbgState, EventDetails);
  g_DebuggeeHaltReason = DEBUGGEE_PAUSING_REASON_NOT_PAUSED;
  RtlZeroMemory(&g_EventTriggerDetail,
                sizeof(DEBUGGER_TRIGGERED_EVENT_DETAILS));
  if (DbgState->MainDebuggingCore) {
    DbgState->MainDebuggingCore = FALSE;
    SpinlockUnlock(&DebuggerHandleBreakpointLock);
  }
}

_Use_decl_annotations_

    BOOLEAN
    KdCheckAndHandleNmiCallback(UINT32 CoreId) {
  BOOLEAN Result = FALSE;
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  if (DbgState->NmiState.WaitingToBeLocked) {
    Result = TRUE;
    if (DbgState->NmiState.NmiCalledInVmxRootRelatedToHaltDebuggee) {
      KdHandleHaltsWhenNmiReceivedFromVmxRoot(DbgState);
    } else {
      KdHandleNmi(DbgState);
    }
  }
  return Result;
}

_Use_decl_annotations_

    VOID
    KdHandleNmi(PROCESSOR_DEBUGGING_STATE *DbgState) {
  DbgState->MainDebuggingCore = FALSE;
  DbgState->NmiState.WaitingToBeLocked = FALSE;
  SpinlockLock(&DbgState->Lock);
  KdManageSystemHaltOnVmxRoot(DbgState, NULL);
  if (DbgState->MainDebuggingCore) {
    DbgState->MainDebuggingCore = FALSE;
    SpinlockUnlock(&DebuggerHandleBreakpointLock);
  }
}

VOID KdGuaranteedStepInstruction(PROCESSOR_DEBUGGING_STATE *DbgState) {
  UINT64 CsSel = (UINT64)NULL;
  CsSel = VmFuncGetCsSelector();
  DbgState->InstrumentationStepInTrace.CsSel = (UINT16)CsSel;
  VmFuncRegisterMtfBreak(DbgState->CoreId);
  VmFuncChangeMtfUnsettingState(DbgState->CoreId, TRUE);
  VmFuncDisableExternalInterruptsAndInterruptWindow(DbgState->CoreId);
  VmFuncSetMonitorTrapFlag(TRUE);
}

BOOLEAN KdCheckGuestOperatingModeChanges(UINT16 PreviousCsSelector,
                                         UINT16 CurrentCsSelector) {
  PreviousCsSelector = PreviousCsSelector & ~3;
  CurrentCsSelector = CurrentCsSelector & ~3;
  if (PreviousCsSelector == CurrentCsSelector) {
    return FALSE;
  }
  if ((PreviousCsSelector == KGDT64_R3_CODE ||
       PreviousCsSelector == KGDT64_R3_CMCODE) &&
      CurrentCsSelector == KGDT64_R0_CODE) {
    LogInfo("User-mode -> Kernel-mode\n");
  } else if ((CurrentCsSelector == KGDT64_R3_CODE ||
              CurrentCsSelector == KGDT64_R3_CMCODE) &&
             PreviousCsSelector == KGDT64_R0_CODE) {
    LogInfo("Kernel-mode -> User-mode\n");
  } else if (CurrentCsSelector == KGDT64_R3_CODE &&
             PreviousCsSelector == KGDT64_R3_CMCODE) {
    LogInfo("32-bit User-mode -> 64-bit User-mode (Heaven's gate)\n");
  } else if (PreviousCsSelector == KGDT64_R3_CODE &&
             CurrentCsSelector == KGDT64_R3_CMCODE) {
    LogInfo(
        "64-bit User-mode -> 32-bit User-mode (Return from Heaven's gate)\n");
  } else {
    LogError(
        "Err, unknown changes in cs selector during the instrumentation "
        "step-in\n");
  }
  return TRUE;
}

VOID KdRegularStepInInstruction(PROCESSOR_DEBUGGING_STATE *DbgState) {
  TracingPerformRegularStepInInstruction(DbgState);
  BreakpointRestoreTheTrapFlagOnceTriggered(
      HANDLE_TO_UINT32(PsGetCurrentProcessId()),
      HANDLE_TO_UINT32(PsGetCurrentThreadId()));
}

VOID KdRegularStepOver(PROCESSOR_DEBUGGING_STATE *DbgState,
                       BOOLEAN IsNextInstructionACall, UINT32 CallLength) {
  UINT64 NextAddressForHardwareDebugBp = 0;
  ULONG ProcessorsCount;
  if (IsNextInstructionACall) {
    NextAddressForHardwareDebugBp =
        VmFuncGetLastVmexitRip(DbgState->CoreId) + CallLength;
    ProcessorsCount = KeQueryActiveProcessorCount(0);
    g_HardwareDebugRegisterDetailsForStepOver.Address =
        NextAddressForHardwareDebugBp;
    g_HardwareDebugRegisterDetailsForStepOver.ProcessId =
        HANDLE_TO_UINT32(PsGetCurrentProcessId());
    g_HardwareDebugRegisterDetailsForStepOver.ThreadId =
        HANDLE_TO_UINT32(PsGetCurrentThreadId());
    for (size_t i = 0; i < ProcessorsCount; i++) {
      g_DbgState[i].HardwareDebugRegisterForStepping =
          NextAddressForHardwareDebugBp;
    }
  } else {
    KdRegularStepInInstruction(DbgState);
  }
}

BOOLEAN KdPerformRegisterEvent(
    PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET EventDetailHeader,
    DEBUGGER_EVENT_AND_ACTION_RESULT *DebuggerEventAndActionResult) {
#if EnableInstantEventMechanism
  DEBUGGER_GENERAL_EVENT_DETAIL *GeneralEventDetail = NULL;
  GeneralEventDetail =
      (PDEBUGGER_GENERAL_EVENT_DETAIL)(((CHAR *)EventDetailHeader) +
                                       sizeof(
                                           DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET));
  if (!KdCheckAllCoresAreLocked()) {
    DebuggerEventAndActionResult->IsSuccessful = FALSE;
    DebuggerEventAndActionResult->Error =
        DEBUGGER_ERROR_NOT_ALL_CORES_ARE_LOCKED_FOR_APPLYING_INSTANT_EVENT;
  } else {
    DebuggerParseEvent(GeneralEventDetail, DebuggerEventAndActionResult, TRUE);
  }
  return FALSE;
#else KdCheckUserModePriorityBuffers();
  LogCallbackSendBuffer(
      OPERATION_DEBUGGEE_REGISTER_EVENT,
      ((CHAR *)EventDetailHeader +
       sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)),
      EventDetailHeader->Length, TRUE);
  return TRUE;
#endif
}

BOOLEAN KdPerformAddActionToEvent(
    PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET ActionDetailHeader,
    DEBUGGER_EVENT_AND_ACTION_RESULT *DebuggerEventAndActionResult) {
#if EnableInstantEventMechanism
  DEBUGGER_GENERAL_ACTION *GeneralActionDetail = NULL;
  GeneralActionDetail =
      (PDEBUGGER_GENERAL_ACTION)(((CHAR *)ActionDetailHeader) +
                                 sizeof(
                                     DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET));
  DebuggerParseAction(GeneralActionDetail, DebuggerEventAndActionResult, TRUE);
  return FALSE;
#else KdCheckUserModePriorityBuffers();
  LogCallbackSendBuffer(
      OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT,
      ((CHAR *)ActionDetailHeader +
       sizeof(DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET)),
      ActionDetailHeader->Length, TRUE);
  return TRUE;
#endif
}

VOID KdQueryRflagTrapState() {
  LogInfo(
      "Number of valid entries: 0x%x\n"
      "(Please be aware that only top 0x%x items are considered valid. "
      "There could be other items present in the array, but they are not "
      "valid.)",
      g_TrapFlagState.NumberOfItems, g_TrapFlagState.NumberOfItems);
  for (size_t i = 0; i < MAXIMUM_NUMBER_OF_THREAD_INFORMATION_FOR_TRAPS; i++) {
    LogInfo(
        "g_TrapFlagState.ThreadInformation[%d].ProcessId = %x | ThreadId = %x",
        i, g_TrapFlagState.ThreadInformation[i].Fields.ProcessId,
        g_TrapFlagState.ThreadInformation[i].Fields.ThreadId);
  }
}

BOOLEAN KdCheckAllCoresAreLocked() {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  for (size_t i = 0; i < ProcessorsCount; i++) {
    if (!SpinlockCheckLock(&g_DbgState[i].Lock)) {
      return FALSE;
    }
  }
  return TRUE;
}

BOOLEAN KdCheckTargetCoreIsLocked(UINT32 CoreNumber) {
  if (!SpinlockCheckLock(&g_DbgState[CoreNumber].Lock)) {
    return FALSE;
  } else {
    return TRUE;
  }
}

VOID KdQuerySystemState() {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  Log("================================================ Debugging Lock Info "
      "================================================\n");
  for (size_t i = 0; i < ProcessorsCount; i++) {
    if (SpinlockCheckLock(&g_DbgState[i].Lock)) {
      LogInfo("Core : %d is locked", i);
    } else {
      LogInfo("Core : %d isn't locked", i);
    }
  }
  Log("\n================================================ NMI Receiver State "
      "=======+=========================================\n");
  for (size_t i = 0; i < ProcessorsCount; i++) {
    if (g_DbgState[i].NmiState.NmiCalledInVmxRootRelatedToHaltDebuggee) {
      LogInfo("Core : %d - called from an NMI that is called in VMX-root mode",
              i);
    } else {
      LogInfo(
          "Core : %d - not called from an NMI handler (through the immediate "
          "VM-exit mechanism)",
          i);
    }
  }
}

VOID KdUnlockTheHaltedCore(PROCESSOR_DEBUGGING_STATE *DbgState) {
  SpinlockUnlock(&DbgState->Lock);
}

BOOLEAN KdCheckTheHaltedCore(PROCESSOR_DEBUGGING_STATE *DbgState) {
  return SpinlockCheckLock(&DbgState->Lock);
}

BOOLEAN KdBringPagein(PROCESSOR_DEBUGGING_STATE *DbgState,
                      PDEBUGGER_PAGE_IN_REQUEST PageinRequest) {
  VmFuncEventInjectPageFaultRangeAddress(
      DbgState->CoreId, PageinRequest->VirtualAddressFrom,
      PageinRequest->VirtualAddressTo, PageinRequest->PageFaultErrorCode);
  VmFuncSetRflagTrapFlag(TRUE);
  if (!BreakpointRestoreTheTrapFlagOnceTriggered(
          HANDLE_TO_UINT32(PsGetCurrentProcessId()),
          HANDLE_TO_UINT32(PsGetCurrentThreadId()))) {
    PageinRequest->KernelStatus = DEBUGGER_ERROR_THE_TRAP_FLAG_LIST_IS_FULL;
    return FALSE;
  } else {
    PageinRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
    return TRUE;
  }
}

VOID KdPerformTheTestPacketOperation(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER *TestQueryPacket) {
  switch (TestQueryPacket->RequestType) {
    case TEST_QUERY_HALTING_CORE_STATUS:
      KdQuerySystemState();
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_QUERY_TRAP_STATE:
      KdQueryRflagTrapState();
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_QUERY_PREALLOCATED_POOL_STATE:
      PoolManagerShowPreAllocatedPools();
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_SYNCHRONOUS:
    case TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_ASYNCHRONOUS:
      HaltedCoreBroadcastTaskAllCores(
          DbgState, DEBUGGER_HALTED_CORE_TASK_TEST, TRUE,
          TestQueryPacket->RequestType ==
                  TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_SYNCHRONOUS
              ? TRUE
              : FALSE,
          (PVOID)0x55);
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_SETTING_TARGET_TASKS_ON_TARGET_HALTED_CORES:
      if (!CommonValidateCoreNumber((UINT32)TestQueryPacket->Context)) {
        TestQueryPacket->KernelStatus = DEBUGGER_ERROR_INVALID_CORE_ID;
      } else {
        HaltedCoreRunTaskOnSingleCore((UINT32)TestQueryPacket->Context,
                                      DEBUGGER_HALTED_CORE_TASK_TEST, TRUE,
                                      (PVOID)0x85);
        TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      }
      break;
    case TEST_BREAKPOINT_TURN_OFF_BPS:
      g_InterceptBreakpoints = TRUE;
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_BREAKPOINT_TURN_ON_BPS:
      g_InterceptBreakpoints = FALSE;
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_BREAKPOINT_TURN_OFF_BPS_AND_EVENTS_FOR_COMMANDS_IN_REMOTE_COMPUTER:
      g_InterceptBreakpointsAndEventsForCommandsInRemoteComputer = TRUE;
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_BREAKPOINT_TURN_ON_BPS_AND_EVENTS_FOR_COMMANDS_IN_REMOTE_COMPUTER:
      g_InterceptBreakpointsAndEventsForCommandsInRemoteComputer = FALSE;
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_BREAKPOINT_TURN_OFF_DBS:
      g_InterceptDebugBreaks = TRUE;
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case TEST_BREAKPOINT_TURN_ON_DBS:
      g_InterceptDebugBreaks = FALSE;
      TestQueryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    default:
      TestQueryPacket->KernelStatus =
          DEBUGGER_ERROR_UNKNOWN_TEST_QUERY_RECEIVED;
      break;
  }
}

VOID KdPerformSettingTheStateOfShortCircuiting(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    PDEBUGGER_SHORT_CIRCUITING_EVENT ShortCircuitingEvent) {
  if (ShortCircuitingEvent->IsShortCircuiting) {
    DbgState->ShortCircuitingEvent = TRUE;
  } else {
    DbgState->ShortCircuitingEvent = FALSE;
  }
  ShortCircuitingEvent->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
}

BOOLEAN KdPerformEventQueryAndModification(
    PDEBUGGER_MODIFY_EVENTS ModifyAndQueryEvent) {
  BOOLEAN IsForAllEvents = FALSE;
  BOOLEAN ContinueDebugger = FALSE;
  if (ModifyAndQueryEvent->Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG) {
    IsForAllEvents = TRUE;
  } else if (!DebuggerIsTagValid(ModifyAndQueryEvent->Tag)) {
    ModifyAndQueryEvent->KernelStatus =
        DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG;
    return FALSE;
  }
  if (ModifyAndQueryEvent->TypeOfAction == DEBUGGER_MODIFY_EVENTS_QUERY_STATE) {
    if (!DebuggerIsTagValid(ModifyAndQueryEvent->Tag)) {
      ModifyAndQueryEvent->KernelStatus = DEBUGGER_ERROR_TAG_NOT_EXISTS;
    } else {
      if (DebuggerQueryStateEvent(ModifyAndQueryEvent->Tag)) {
        ModifyAndQueryEvent->IsEnabled = TRUE;
      } else {
        ModifyAndQueryEvent->IsEnabled = FALSE;
      }
      ModifyAndQueryEvent->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
    }
  } else if (ModifyAndQueryEvent->TypeOfAction ==
             DEBUGGER_MODIFY_EVENTS_ENABLE) {
    if (IsForAllEvents) {
      DebuggerEnableOrDisableAllEvents(TRUE);
    } else {
      DebuggerEnableEvent(ModifyAndQueryEvent->Tag);
    }
    ModifyAndQueryEvent->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  } else if (ModifyAndQueryEvent->TypeOfAction ==
             DEBUGGER_MODIFY_EVENTS_DISABLE) {
    if (IsForAllEvents) {
      DebuggerEnableOrDisableAllEvents(FALSE);
    } else {
      DebuggerDisableEvent(ModifyAndQueryEvent->Tag);
    }
    ModifyAndQueryEvent->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  } else if (ModifyAndQueryEvent->TypeOfAction ==
             DEBUGGER_MODIFY_EVENTS_CLEAR) {
#if EnableInstantEventMechanism if (IsForAllEvents) {
    DebuggerEnableOrDisableAllEvents(FALSE);
  } else {
    DebuggerDisableEvent(ModifyAndQueryEvent->Tag);
  }
  if (!LogCallbackCheckIfBufferIsFull(TRUE)) {
    LogCallbackSendBuffer(
        OPERATION_DEBUGGEE_CLEAR_EVENTS_WITHOUT_NOTIFYING_DEBUGGER,
        ModifyAndQueryEvent, sizeof(DEBUGGER_MODIFY_EVENTS), TRUE);
    ModifyAndQueryEvent->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  } else {
    ModifyAndQueryEvent->KernelStatus =
        DEBUGGER_ERROR_THE_TARGET_EVENT_IS_DISABLED_BUT_CANNOT_BE_CLEARED_PRIRITY_BUFFER_IS_FULL;
  }
#else KdCheckUserModePriorityBuffers();
    LogCallbackSendBuffer(OPERATION_DEBUGGEE_CLEAR_EVENTS, ModifyAndQueryEvent,
                          sizeof(DEBUGGER_MODIFY_EVENTS), TRUE);
    ContinueDebugger = TRUE;
#endif
}
else {
  ModifyAndQueryEvent->KernelStatus =
      DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION;
}
return ContinueDebugger;
}

VOID KdDispatchAndPerformCommandsFromDebugger(
    PROCESSOR_DEBUGGING_STATE *DbgState) {
  PDEBUGGEE_CHANGE_CORE_PACKET ChangeCorePacket;
  PDEBUGGEE_STEP_PACKET SteppingPacket;
  PDEBUGGER_FLUSH_LOGGING_BUFFERS FlushPacket;
  PDEBUGGER_CALLSTACK_REQUEST CallstackPacket;
  PDEBUGGER_SINGLE_CALLSTACK_FRAME CallstackFrameBuffer;
  PDEBUGGER_DEBUGGER_TEST_QUERY_BUFFER TestQueryPacket;
  PDEBUGGEE_REGISTER_READ_DESCRIPTION ReadRegisterPacket;
  PDEBUGGEE_REGISTER_WRITE_DESCRIPTION WriteRegisterPacket;
  PDEBUGGER_READ_MEMORY ReadMemoryPacket;
  PDEBUGGER_EDIT_MEMORY EditMemoryPacket;
  PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET ChangeProcessPacket;
  PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET ChangeThreadPacket;
  PDEBUGGEE_SCRIPT_PACKET ScriptPacket;
  PDEBUGGEE_USER_INPUT_PACKET UserInputPacket;
  PDEBUGGER_SEARCH_MEMORY SearchQueryPacket;
  PDEBUGGEE_BP_PACKET BpPacket;
  PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PtePacket;
  PDEBUGGER_APIC_REQUEST ApicPacket;
  PINTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS IdtEntryPacket;
  PDEBUGGER_PAGE_IN_REQUEST PageinPacket;
  PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS Va2paPa2vaPacket;
  PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET BpListOrModifyPacket;
  PDEBUGGEE_SYMBOL_REQUEST_PACKET SymReloadPacket;
  PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET EventRegPacket;
  PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET AddActionPacket;
  PDEBUGGER_MODIFY_EVENTS QueryAndModifyEventPacket;
  PDEBUGGER_SHORT_CIRCUITING_EVENT ShortCircuitingEventPacket;
  UINT32 SizeToSend = 0;
  BOOLEAN UnlockTheNewCore = FALSE;
  UINT32 ReturnSize = 0;
  DEBUGGEE_RESULT_OF_SEARCH_PACKET SearchPacketResult = {0};
  DEBUGGER_EVENT_AND_ACTION_RESULT DebuggerEventAndActionResult = {0};
  PDEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET PcitreePacket = {0};
  PDEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET PcidevinfoPacket = {0};
  while (TRUE) {
    BOOLEAN EscapeFromTheLoop = FALSE;
    CHAR *RecvBuffer = &DbgState->KdRecvBuffer[0];
    UINT32 RecvBufferLength = 0;
    PDEBUGGER_REMOTE_PACKET TheActualPacket =
        (PDEBUGGER_REMOTE_PACKET)RecvBuffer;
    RtlZeroMemory(RecvBuffer, MaxSerialPacketSize);
    if (!SerialConnectionRecvBuffer(RecvBuffer, &RecvBufferLength)) {
      continue;
    }
    if (TheActualPacket->Indicator == INDICATOR_OF_HYPERDBG_PACKET) {
      if (KdComputeDataChecksum((PVOID)&TheActualPacket->Indicator,
                                RecvBufferLength - sizeof(BYTE)) !=
          TheActualPacket->Checksum) {
        LogError("Err, checksum is invalid");
        continue;
      }
      if (TheActualPacket->TypeOfThePacket !=
          DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT) {
        LogError("Err, unknown packet received from the debugger\n");
        continue;
      }
      switch (TheActualPacket->RequestedActionOfThePacket) {
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE:
          KdContinueDebuggee(DbgState, FALSE,
                             DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
          EscapeFromTheLoop = TRUE;
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_STEP:
          SteppingPacket =
              (DEBUGGEE_STEP_PACKET *)(((CHAR *)TheActualPacket) +
                                       sizeof(DEBUGGER_REMOTE_PACKET));
          switch (SteppingPacket->StepType) {
            case DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN:
            case DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN_FOR_TRACKING:
              KdGuaranteedStepInstruction(DbgState);
              KdContinueDebuggeeJustCurrentCore(DbgState);
              if (SteppingPacket->StepType ==
                  DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN_FOR_TRACKING) {
                DbgState->IgnoreDisasmInNextPacket = TRUE;
              }
              EscapeFromTheLoop = TRUE;
              break;
            case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER:
            case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU:
            case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU_LAST_INSTRUCTION:
              KdRegularStepOver(DbgState,
                                SteppingPacket->IsCurrentInstructionACall,
                                SteppingPacket->CallLength);
              KdContinueDebuggee(
                  DbgState, FALSE,
                  DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
              if (SteppingPacket->StepType ==
                  DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU) {
                DbgState->IgnoreDisasmInNextPacket = TRUE;
              }
              EscapeFromTheLoop = TRUE;
              break;
            case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN:
              KdRegularStepInInstruction(DbgState);
              KdContinueDebuggee(
                  DbgState, FALSE,
                  DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
              EscapeFromTheLoop = TRUE;
              break;
            default:
              break;
          }
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CLOSE_AND_UNLOAD_DEBUGGEE:
          KdCloseConnectionAndUnloadDebuggee();
          KdContinueDebuggee(DbgState, FALSE,
                             DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
          EscapeFromTheLoop = TRUE;
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE:
          ChangeCorePacket =
              (DEBUGGEE_CHANGE_CORE_PACKET *)(((CHAR *)TheActualPacket) +
                                              sizeof(DEBUGGER_REMOTE_PACKET));
          if (KdSwitchCore(DbgState, ChangeCorePacket)) {
            EscapeFromTheLoop = TRUE;
            UnlockTheNewCore = TRUE;
          }
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE,
              (CHAR *)ChangeCorePacket, sizeof(DEBUGGEE_CHANGE_CORE_PACKET));
          if (UnlockTheNewCore) {
            UnlockTheNewCore = FALSE;
            SpinlockUnlock(&g_DbgState[ChangeCorePacket->NewCore].Lock);
          }
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_FLUSH_BUFFERS:
          FlushPacket =
              (DEBUGGER_FLUSH_LOGGING_BUFFERS *)(((CHAR *)TheActualPacket) +
                                                 sizeof(
                                                     DEBUGGER_REMOTE_PACKET));
          DebuggerCommandFlush(FlushPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH,
              (CHAR *)FlushPacket, sizeof(DEBUGGER_FLUSH_LOGGING_BUFFERS));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK:
          CallstackPacket =
              (DEBUGGER_CALLSTACK_REQUEST *)(((CHAR *)TheActualPacket) +
                                             sizeof(DEBUGGER_REMOTE_PACKET));
          CallstackFrameBuffer = (DEBUGGER_SINGLE_CALLSTACK_FRAME
                                      *)(((CHAR *)TheActualPacket) +
                                         sizeof(DEBUGGER_REMOTE_PACKET) +
                                         sizeof(DEBUGGER_CALLSTACK_REQUEST));
          if (CallstackPacket->BaseAddress == (UINT64)NULL) {
            CallstackPacket->BaseAddress = DbgState->Regs->rsp;
          }
          if (CallstackWalkthroughStack(
                  CallstackFrameBuffer, &CallstackPacket->FrameCount,
                  CallstackPacket->BaseAddress, CallstackPacket->Size,
                  CallstackPacket->Is32Bit)) {
            CallstackPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
          } else {
            CallstackPacket->KernelStatus =
                DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK;
          }
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK,
              (CHAR *)CallstackPacket, (UINT32)CallstackPacket->BufferSize);
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_TEST_QUERY:
          TestQueryPacket = (DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER
                                 *)(((CHAR *)TheActualPacket) +
                                    sizeof(DEBUGGER_REMOTE_PACKET));
          KdPerformTheTestPacketOperation(DbgState, TestQueryPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY,
              (CHAR *)TestQueryPacket,
              sizeof(DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_REGISTERS:
          ReadRegisterPacket = (DEBUGGEE_REGISTER_READ_DESCRIPTION
                                    *)(((CHAR *)TheActualPacket) +
                                       sizeof(DEBUGGER_REMOTE_PACKET));
          if (KdReadRegisters(DbgState, ReadRegisterPacket)) {
            ReadRegisterPacket->KernelStatus =
                DEBUGGER_OPERATION_WAS_SUCCESSFUL;
          } else {
            ReadRegisterPacket->KernelStatus =
                DEBUGGER_ERROR_INVALID_REGISTER_NUMBER;
          }
          if (ReadRegisterPacket->RegisterId == DEBUGGEE_SHOW_ALL_REGISTERS) {
            SizeToSend = sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION) +
                         sizeof(GUEST_REGS) + sizeof(GUEST_EXTRA_REGISTERS);
          } else {
            SizeToSend = sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION);
          }
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS,
              (CHAR *)ReadRegisterPacket, SizeToSend);
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_WRITE_REGISTER:
          WriteRegisterPacket = (DEBUGGEE_REGISTER_WRITE_DESCRIPTION
                                     *)(((CHAR *)TheActualPacket) +
                                        sizeof(DEBUGGER_REMOTE_PACKET));
          if (SetRegValue(DbgState->Regs, WriteRegisterPacket->RegisterId,
                          WriteRegisterPacket->Value)) {
            WriteRegisterPacket->KernelStatus =
                DEBUGGER_OPERATION_WAS_SUCCESSFUL;
          } else {
            WriteRegisterPacket->KernelStatus =
                DEBUGGER_ERROR_INVALID_REGISTER_NUMBER;
          }
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_WRITE_REGISTER,
              (CHAR *)WriteRegisterPacket,
              sizeof(DEBUGGEE_REGISTER_WRITE_DESCRIPTION));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_MEMORY:
          ReadMemoryPacket =
              (DEBUGGER_READ_MEMORY *)(((CHAR *)TheActualPacket) +
                                       sizeof(DEBUGGER_REMOTE_PACKET));
          if (DebuggerCommandReadMemoryVmxRoot(
                  ReadMemoryPacket,
                  (PVOID)((UINT64)ReadMemoryPacket +
                          sizeof(DEBUGGER_READ_MEMORY)),
                  &ReturnSize)) {
            ReadMemoryPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
          } else {
            ReadMemoryPacket->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
          }
          ReadMemoryPacket->ReturnLength = ReturnSize;
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY,
              (CHAR *)ReadMemoryPacket,
              sizeof(DEBUGGER_READ_MEMORY) + ReturnSize);
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_EDIT_MEMORY:
          EditMemoryPacket =
              (PDEBUGGER_EDIT_MEMORY)(((CHAR *)TheActualPacket) +
                                      sizeof(DEBUGGER_REMOTE_PACKET));
          DebuggerCommandEditMemoryVmxRoot(EditMemoryPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY,
              (CHAR *)EditMemoryPacket, sizeof(DEBUGGER_EDIT_MEMORY));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS:
          ChangeProcessPacket = (DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
                                     *)(((CHAR *)TheActualPacket) +
                                        sizeof(DEBUGGER_REMOTE_PACKET));
          ProcessInterpretProcess(DbgState, ChangeProcessPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS,
              (CHAR *)ChangeProcessPacket,
              sizeof(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD:
          ChangeThreadPacket = (DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
                                    *)(((CHAR *)TheActualPacket) +
                                       sizeof(DEBUGGER_REMOTE_PACKET));
          ThreadInterpretThread(DbgState, ChangeThreadPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD,
              (CHAR *)ChangeThreadPacket,
              sizeof(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_RUN_SCRIPT:
          ScriptPacket =
              (DEBUGGEE_SCRIPT_PACKET *)(((CHAR *)TheActualPacket) +
                                         sizeof(DEBUGGER_REMOTE_PACKET));
          if (DebuggerPerformRunScript(DbgState, NULL, ScriptPacket,
                                       &g_EventTriggerDetail)) {
            ScriptPacket->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
          } else {
            ScriptPacket->Result =
                DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT;
          }
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT,
              (CHAR *)ScriptPacket, sizeof(DEBUGGEE_SCRIPT_PACKET));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_USER_INPUT_BUFFER:
          UserInputPacket =
              (DEBUGGEE_USER_INPUT_PACKET *)(((CHAR *)TheActualPacket) +
                                             sizeof(DEBUGGER_REMOTE_PACKET));
          KdNotifyDebuggeeForUserInput(
              ((DEBUGGEE_USER_INPUT_PACKET *)UserInputPacket),
              sizeof(DEBUGGEE_USER_INPUT_PACKET) + UserInputPacket->CommandLen);
          KdContinueDebuggee(DbgState, FALSE,
                             DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
          EscapeFromTheLoop = TRUE;
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SEARCH_QUERY:
          SearchQueryPacket =
              (DEBUGGER_SEARCH_MEMORY *)(((CHAR *)TheActualPacket) +
                                         sizeof(DEBUGGER_REMOTE_PACKET));
          if (SearchAddressWrapper(
                  NULL, SearchQueryPacket, SearchQueryPacket->Address,
                  SearchQueryPacket->Address + SearchQueryPacket->Length, TRUE,
                  &SearchPacketResult.CountOfResults)) {
            SearchPacketResult.Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
          } else {
            SearchPacketResult.Result = DEBUGGER_ERROR_INVALID_ADDRESS;
          }
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY,
              (CHAR *)&SearchPacketResult,
              sizeof(DEBUGGEE_RESULT_OF_SEARCH_PACKET));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_REGISTER_EVENT:
          EventRegPacket = (DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET
                                *)(((CHAR *)TheActualPacket) +
                                   sizeof(DEBUGGER_REMOTE_PACKET));
          if (KdPerformRegisterEvent(EventRegPacket,
                                     &DebuggerEventAndActionResult)) {
            KdContinueDebuggee(
                DbgState, TRUE,
                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT);
            EscapeFromTheLoop = TRUE;
          } else {
            KdResponsePacketToDebugger(
                DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT,
                (CHAR *)&DebuggerEventAndActionResult,
                sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT));
          }
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_ADD_ACTION_TO_EVENT:
          AddActionPacket = (DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET
                                 *)(((CHAR *)TheActualPacket) +
                                    sizeof(DEBUGGER_REMOTE_PACKET));
          if (KdPerformAddActionToEvent(AddActionPacket,
                                        &DebuggerEventAndActionResult)) {
            KdContinueDebuggee(
                DbgState, TRUE,
                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT);
            EscapeFromTheLoop = TRUE;
          } else {
            KdResponsePacketToDebugger(
                DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT,
                (CHAR *)&DebuggerEventAndActionResult,
                sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT));
          }
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_AND_MODIFY_EVENT:
          QueryAndModifyEventPacket =
              (DEBUGGER_MODIFY_EVENTS *)(((CHAR *)TheActualPacket) +
                                         sizeof(DEBUGGER_REMOTE_PACKET));
          if (KdPerformEventQueryAndModification(QueryAndModifyEventPacket)) {
            KdContinueDebuggee(
                DbgState, TRUE,
                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT);
            EscapeFromTheLoop = TRUE;
          } else {
            KdResponsePacketToDebugger(
                DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
                DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT,
                (CHAR *)QueryAndModifyEventPacket,
                sizeof(DEBUGGER_MODIFY_EVENTS));
          }
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SET_SHORT_CIRCUITING_STATE:
          ShortCircuitingEventPacket =
              (DEBUGGER_SHORT_CIRCUITING_EVENT *)(((CHAR *)TheActualPacket) +
                                                  sizeof(
                                                      DEBUGGER_REMOTE_PACKET));
          KdPerformSettingTheStateOfShortCircuiting(DbgState,
                                                    ShortCircuitingEventPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_STATE,
              (CHAR *)ShortCircuitingEventPacket,
              sizeof(DEBUGGER_SHORT_CIRCUITING_EVENT));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_BP:
          BpPacket = (DEBUGGEE_BP_PACKET *)(((CHAR *)TheActualPacket) +
                                            sizeof(DEBUGGER_REMOTE_PACKET));
          BreakpointAddNew(BpPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP,
              (CHAR *)BpPacket, sizeof(DEBUGGEE_BP_PACKET));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_QUERY_PTE:
          PtePacket = (DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
                           *)(((CHAR *)TheActualPacket) +
                              sizeof(DEBUGGER_REMOTE_PACKET));
          ExtensionCommandPte(PtePacket, TRUE);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE,
              (CHAR *)PtePacket,
              sizeof(DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_PERFORM_ACTIONS_ON_APIC:
          ApicPacket =
              (DEBUGGER_APIC_REQUEST *)(((CHAR *)TheActualPacket) +
                                        sizeof(DEBUGGER_REMOTE_PACKET));
          SizeToSend =
              ExtensionCommandPerformActionsForApicRequests(ApicPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_APIC_REQUESTS,
              (CHAR *)ApicPacket, SizeToSend);
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_IDT_ENTRIES:
          IdtEntryPacket = (INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS
                                *)(((CHAR *)TheActualPacket) +
                                   sizeof(DEBUGGER_REMOTE_PACKET));
          ExtensionCommandPerformQueryIdtEntriesRequest(IdtEntryPacket, TRUE);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_IDT_ENTRIES_REQUESTS,
              (CHAR *)IdtEntryPacket,
              SIZEOF_INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS);
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_INJECT_PAGE_FAULT:
          PageinPacket =
              (DEBUGGER_PAGE_IN_REQUEST *)(((CHAR *)TheActualPacket) +
                                           sizeof(DEBUGGER_REMOTE_PACKET));
          KdBringPagein(DbgState, PageinPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BRINGING_PAGES_IN,
              (CHAR *)PageinPacket, sizeof(DEBUGGER_PAGE_IN_REQUEST));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PA2VA_AND_VA2PA:
          Va2paPa2vaPacket = (DEBUGGER_VA2PA_AND_PA2VA_COMMANDS
                                  *)(((CHAR *)TheActualPacket) +
                                     sizeof(DEBUGGER_REMOTE_PACKET));
          ExtensionCommandVa2paAndPa2va(Va2paPa2vaPacket, TRUE);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA,
              (CHAR *)Va2paPa2vaPacket,
              sizeof(DEBUGGER_VA2PA_AND_PA2VA_COMMANDS));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_LIST_OR_MODIFY_BREAKPOINTS:
          BpListOrModifyPacket = (DEBUGGEE_BP_LIST_OR_MODIFY_PACKET
                                      *)(((CHAR *)TheActualPacket) +
                                         sizeof(DEBUGGER_REMOTE_PACKET));
          BreakpointListOrModify(BpListOrModifyPacket);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS,
              (CHAR *)BpListOrModifyPacket,
              sizeof(DEBUGGEE_BP_LIST_OR_MODIFY_PACKET));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_RELOAD:
          SymReloadPacket =
              (DEBUGGEE_SYMBOL_REQUEST_PACKET *)(((CHAR *)TheActualPacket) +
                                                 sizeof(
                                                     DEBUGGER_REMOTE_PACKET));
          KdReloadSymbolDetailsInDebuggee(SymReloadPacket);
          KdContinueDebuggee(DbgState, FALSE,
                             DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION);
          EscapeFromTheLoop = TRUE;
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PCITREE:
          PcitreePacket = (DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET
                               *)(((CHAR *)TheActualPacket) +
                                  sizeof(DEBUGGER_REMOTE_PACKET));
          ExtensionCommandPcitree(PcitreePacket, TRUE);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PCITREE,
              (CHAR *)PcitreePacket,
              sizeof(DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET));
          break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PCIDEVINFO:
          PcidevinfoPacket = (DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET
                                  *)(((CHAR *)TheActualPacket) +
                                     sizeof(DEBUGGER_REMOTE_PACKET));
          ExtensionCommandPcidevinfo(PcidevinfoPacket, TRUE);
          KdResponsePacketToDebugger(
              DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
              DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PCIDEVINFO,
              (CHAR *)PcidevinfoPacket,
              sizeof(DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET));
          break;
        default:
          LogError("Err, unknown packet action received from the debugger\n");
          break;
      }
    } else {
      LogError(
          "Err, it's not a HyperDbg packet, the packet is probably deformed\n");
      continue;
    }
    if (EscapeFromTheLoop) {
      break;
    }
  }
}

BOOLEAN KdIsGuestOnUsermode32Bit() {
  UINT64 CsSel = (UINT64)NULL;
  CsSel = VmFuncGetCsSelector();
  if (CsSel == KGDT64_R0_CODE) {
    return FALSE;
  } else if ((CsSel & ~3) == KGDT64_R3_CODE) {
    return FALSE;
  } else if ((CsSel & ~3) == KGDT64_R3_CMCODE) {
    return TRUE;
  } else {
    LogError("Err, unknown value for cs, cannot determine wow64 mode");
  }
  return FALSE;
}

VOID KdManageSystemHaltOnVmxRoot(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails) {
  DEBUGGEE_KD_PAUSED_PACKET PausePacket;
  ULONG ExitInstructionLength = 0;
  RFLAGS Rflags = {0};
  UINT64 LastVmexitRip = 0;
  KdApplyTasksPreHaltCore(DbgState);
StartAgain:
  if (DbgState->MainDebuggingCore) {
    RtlZeroMemory(&PausePacket, sizeof(DEBUGGEE_KD_PAUSED_PACKET));
    LastVmexitRip = VmFuncGetRip();
    PausePacket.PausingReason = g_DebuggeeHaltReason;
    PausePacket.CurrentCore = DbgState->CoreId;
    PausePacket.Rip = LastVmexitRip;
    PausePacket.IsProcessorOn32BitMode = KdIsGuestOnUsermode32Bit();
    PausePacket.IgnoreDisassembling = DbgState->IgnoreDisasmInNextPacket;
    DbgState->IgnoreDisasmInNextPacket = FALSE;
    Rflags.AsUInt = VmFuncGetRflags();
    PausePacket.Rflags = Rflags.AsUInt;
    if (EventDetails != NULL) {
      PausePacket.EventTag = EventDetails->Tag;
      PausePacket.EventCallingStage = EventDetails->Stage;
    }
    if (DbgState->InstructionLengthHint != 0) {
      ExitInstructionLength = DbgState->InstructionLengthHint;
    } else {
      ExitInstructionLength =
          CheckAddressMaximumInstructionLength((PVOID)LastVmexitRip);
    }
    PausePacket.ReadInstructionLen = (UINT16)ExitInstructionLength;
    MemoryMapperReadMemorySafeOnTargetProcess(
        LastVmexitRip, &PausePacket.InstructionBytesOnRip,
        ExitInstructionLength);
    KdResponsePacketToDebugger(
        DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
        DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION,
        (CHAR *)&PausePacket, sizeof(DEBUGGEE_KD_PAUSED_PACKET));
    KdDispatchAndPerformCommandsFromDebugger(DbgState);
    if (!DbgState->MainDebuggingCore) {
      goto StartAgain;
    }
  } else {
    DbgState->NmiState.WaitingToBeLocked = FALSE;
    ScopedSpinlock(
        DbgState->Lock, if (DbgState->MainDebuggingCore) {
          g_DebuggeeHaltReason = DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED;
          goto StartAgain;
        });
    if (DbgState->HaltedCoreTask.PerformHaltedTask) {
      DbgState->HaltedCoreTask.PerformHaltedTask = FALSE;
      HaltedCorePerformTargetTask(DbgState, DbgState->HaltedCoreTask.TargetTask,
                                  DbgState->HaltedCoreTask.Context);
      if (DbgState->HaltedCoreTask.LockAgainAfterTask) {
        SpinlockLock(&DbgState->Lock);
        goto StartAgain;
      }
    }
  }
  KdApplyTasksPostContinueCore(DbgState);
}

VOID KdBroadcastHaltOnAllCores() {
  KeGenericCallDpc(DpcRoutineVmExitAndHaltSystemAllCores, NULL);
}

VOID KdHaltSystem(PDEBUGGER_PAUSE_PACKET_RECEIVED PausePacket) {
  VmFuncVmxVmcall(DEBUGGER_VMCALL_VM_EXIT_HALT_SYSTEM, 0, 0, 0);
  PausePacket->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
}
