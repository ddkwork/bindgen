
#include "pch.h"

BOOLEAN UdInitializeUserDebugger() {
  if (g_UserDebuggerState) {
    return TRUE;
  }
  if (g_PsGetProcessPeb == NULL || g_PsGetProcessWow64Process == NULL ||
      g_ZwQueryInformationProcess == NULL) {
    LogError("Err, unable to find needed functions for user-debugger");
  }
  g_SeedOfUserDebuggingDetails = DebuggerThreadDebuggingTagStartSeed;
  InitializeListHead(&g_ProcessDebuggingDetailsListHead);
  BroadcastEnableDbAndBpExitingAllCores();
  ThreadHolderAllocateThreadHoldingBuffers();
  g_UserDebuggerState = TRUE;
  return TRUE;
}

VOID UdUninitializeUserDebugger() {
  if (g_UserDebuggerState) {
    g_UserDebuggerState = FALSE;
    AttachingRemoveAndFreeAllProcessDebuggingDetails();
  }
}

VOID UdRestoreToOriginalDirection(
    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails) {
  VmFuncSetRip(ThreadDebuggingDetails->ThreadRip);
}

VOID UdContinueThread(
    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails) {
  UdRestoreToOriginalDirection(ThreadDebuggingDetails);
  VmFuncSuppressRipIncrement(KeGetCurrentProcessorNumberEx(NULL));
  ThreadDebuggingDetails->IsPaused = FALSE;
}

VOID UdStepInstructions(
    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails,
    DEBUGGER_REMOTE_STEPPING_REQUEST SteppingType) {
  UdRestoreToOriginalDirection(ThreadDebuggingDetails);
  switch (SteppingType) {
    case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN:
      VmFuncSetRflagTrapFlag(TRUE);
      if (!BreakpointRestoreTheTrapFlagOnceTriggered(
              HANDLE_TO_UINT32(PsGetCurrentProcessId()),
              HANDLE_TO_UINT32(PsGetCurrentThreadId()))) {
        LogWarning(
            "Warning, it is currently not possible to add the current "
            "process/thread to the list of processes "
            "where the trap flag should be masked. Please ensure that you "
            "manually unset the trap flag");
      }
      break;
    case DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER:
      break;
    default:
      break;
  }
  VmFuncSuppressRipIncrement(KeGetCurrentProcessorNumberEx(NULL));
  ThreadDebuggingDetails->IsPaused = FALSE;
}

BOOLEAN UdPerformCommand(
    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails,
    DEBUGGER_UD_COMMAND_ACTION_TYPE UserAction, UINT64 OptionalParam1,
    UINT64 OptionalParam2, UINT64 OptionalParam3, UINT64 OptionalParam4) {
  UNREFERENCED_PARAMETER(OptionalParam2);
  UNREFERENCED_PARAMETER(OptionalParam3);
  UNREFERENCED_PARAMETER(OptionalParam4);
  switch (UserAction) {
    case DEBUGGER_UD_COMMAND_ACTION_TYPE_CONTINUE:
      UdContinueThread(ThreadDebuggingDetails);
      break;
    case DEBUGGER_UD_COMMAND_ACTION_TYPE_REGULAR_STEP:
      UdStepInstructions(ThreadDebuggingDetails,
                         (DEBUGGER_REMOTE_STEPPING_REQUEST)OptionalParam1);
      break;
    default:
      return FALSE;
      break;
  }
  return TRUE;
}

BOOLEAN UdCheckForCommand() {
  PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails;
  if (!g_UserDebuggerState) {
    return FALSE;
  }
  ThreadDebuggingDetails =
      ThreadHolderGetProcessThreadDetailsByProcessIdAndThreadId(
          HANDLE_TO_UINT32(PsGetCurrentProcessId()),
          HANDLE_TO_UINT32(PsGetCurrentThreadId()));
  if (!ThreadDebuggingDetails) {
    return FALSE;
  }
  if (!ThreadDebuggingDetails->IsPaused) {
    return FALSE;
  }
  for (size_t i = 0; i < MAX_USER_ACTIONS_FOR_THREADS; i++) {
    if (ThreadDebuggingDetails->UdAction[i].ActionType !=
        DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE) {
      UdPerformCommand(ThreadDebuggingDetails,
                       ThreadDebuggingDetails->UdAction[i].ActionType,
                       ThreadDebuggingDetails->UdAction[i].OptionalParam1,
                       ThreadDebuggingDetails->UdAction[i].OptionalParam2,
                       ThreadDebuggingDetails->UdAction[i].OptionalParam3,
                       ThreadDebuggingDetails->UdAction[i].OptionalParam4);
      ThreadDebuggingDetails->UdAction[i].OptionalParam1 = (UINT64)NULL;
      ThreadDebuggingDetails->UdAction[i].OptionalParam2 = (UINT64)NULL;
      ThreadDebuggingDetails->UdAction[i].OptionalParam3 = (UINT64)NULL;
      ThreadDebuggingDetails->UdAction[i].OptionalParam4 = (UINT64)NULL;
      ThreadDebuggingDetails->UdAction[i].ActionType =
          DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE;
      break;
    }
  }
  return TRUE;
}

BOOLEAN UdDispatchUsermodeCommands(PDEBUGGER_UD_COMMAND_PACKET ActionRequest) {
  PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails;
  ProcessDebuggingDetails = AttachingFindProcessDebuggingDetailsByToken(
      ActionRequest->ProcessDebuggingDetailToken);
  if (!ProcessDebuggingDetails) {
    return FALSE;
  }
  if (ProcessDebuggingDetails->IsOnThreadInterceptingPhase) {
    AttachingConfigureInterceptingThreads(ProcessDebuggingDetails->Token,
                                          FALSE);
  }
  return ThreadHolderApplyActionToPausedThreads(ProcessDebuggingDetails,
                                                ActionRequest);
}

VOID UdSpinThreadOnNop(
    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails,
    PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails) {
  ThreadDebuggingDetails->ThreadRip = VmFuncGetRip();
  VmFuncSetRip(ProcessDebuggingDetails->UsermodeReservedBuffer);
  ThreadDebuggingDetails->IsPaused = TRUE;
}

VOID UdPrePausingReasons(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails,
    DEBUGGEE_PAUSING_REASON Reason,
    PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails) {
  UNREFERENCED_PARAMETER(DbgState);
  UNREFERENCED_PARAMETER(ThreadDebuggingDetails);
  UNREFERENCED_PARAMETER(EventDetails);
  switch (Reason) {
    case DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK:
      break;
    default:
      break;
  }
}

BOOLEAN UdCheckAndHandleBreakpointsAndDebugBreaks(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGEE_PAUSING_REASON Reason,
    PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails) {
  DEBUGGEE_UD_PAUSED_PACKET PausePacket;
  ULONG ExitInstructionLength = 0;
  UINT32 SizeOfSafeBufferToRead = 0;
  RFLAGS Rflags = {0};
  PUSERMODE_DEBUGGING_PROCESS_DETAILS ProcessDebuggingDetails = NULL;
  PUSERMODE_DEBUGGING_THREAD_DETAILS ThreadDebuggingDetails = NULL;
  UINT64 LastVmexitRip = VmFuncGetLastVmexitRip(DbgState->CoreId);
  if (!g_UserDebuggerState && VmFuncVmxGetCurrentExecutionMode() == FALSE) {
    return FALSE;
  }
  ProcessDebuggingDetails = AttachingFindProcessDebuggingDetailsByProcessId(
      HANDLE_TO_UINT32(PsGetCurrentProcessId()));
  if (!ProcessDebuggingDetails) {
    return FALSE;
  }
  ThreadDebuggingDetails = ThreadHolderFindOrCreateThreadDebuggingDetail(
      HANDLE_TO_UINT32(PsGetCurrentThreadId()), ProcessDebuggingDetails);
  if (!ThreadDebuggingDetails) {
    return FALSE;
  }
  ProcessDebuggingDetails->ActiveThreadId = ThreadDebuggingDetails->ThreadId;
  UdPrePausingReasons(DbgState, ThreadDebuggingDetails, Reason, EventDetails);
  RtlZeroMemory(&PausePacket, sizeof(DEBUGGEE_UD_PAUSED_PACKET));
  PausePacket.PausingReason = Reason;
  PausePacket.ProcessId = HANDLE_TO_UINT32(PsGetCurrentProcessId());
  PausePacket.ThreadId = HANDLE_TO_UINT32(PsGetCurrentThreadId());
  PausePacket.ProcessDebuggingToken = ProcessDebuggingDetails->Token;
  PausePacket.Rip = LastVmexitRip;
  PausePacket.Is32Bit = KdIsGuestOnUsermode32Bit();
  Rflags.AsUInt = VmFuncGetRflags();
  PausePacket.Rflags = Rflags.AsUInt;
  if (EventDetails != NULL) {
    PausePacket.EventTag = EventDetails->Tag;
    PausePacket.EventCallingStage = EventDetails->Stage;
  }
  if (DbgState->InstructionLengthHint != 0) {
    ExitInstructionLength = DbgState->InstructionLengthHint;
  } else {
    SizeOfSafeBufferToRead = (UINT32)(LastVmexitRip & 0xfff);
    SizeOfSafeBufferToRead += MAXIMUM_INSTR_SIZE;
    if (SizeOfSafeBufferToRead >= PAGE_SIZE) {
      SizeOfSafeBufferToRead = SizeOfSafeBufferToRead - PAGE_SIZE;
      SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE - SizeOfSafeBufferToRead;
    } else {
      SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE;
    }
    ExitInstructionLength = SizeOfSafeBufferToRead;
  }
  PausePacket.ReadInstructionLen = (UINT16)ExitInstructionLength;
  MemoryMapperReadMemorySafeOnTargetProcess(
      LastVmexitRip, &PausePacket.InstructionBytesOnRip, ExitInstructionLength);
  RtlCopyMemory(&PausePacket.GuestRegs, DbgState->Regs, sizeof(GUEST_REGS));
  LogCallbackSendBuffer(OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE,
                        &PausePacket, sizeof(DEBUGGEE_UD_PAUSED_PACKET), TRUE);
  UdSpinThreadOnNop(ThreadDebuggingDetails, ProcessDebuggingDetails);
  return TRUE;
}
