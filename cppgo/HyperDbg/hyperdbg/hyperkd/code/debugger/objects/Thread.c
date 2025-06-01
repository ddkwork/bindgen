
#include "pch.h"

BOOLEAN ThreadHandleThreadChange(PROCESSOR_DEBUGGING_STATE *DbgState) {
  if ((g_ThreadSwitch.ThreadId != NULL_ZERO &&
       g_ThreadSwitch.ThreadId == HANDLE_TO_UINT32(PsGetCurrentThreadId())) ||
      (g_ThreadSwitch.Thread != NULL64_ZERO &&
       g_ThreadSwitch.Thread == PsGetCurrentThread())) {
    KdHandleBreakpointAndDebugBreakpoints(
        DbgState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED, NULL);
    return TRUE;
  }
  return FALSE;
}

BOOLEAN ThreadSwitch(PROCESSOR_DEBUGGING_STATE *DbgState, UINT32 ThreadId,
                     PETHREAD EThread, BOOLEAN CheckByClockInterrupt) {
  g_ThreadSwitch.Thread = NULL64_ZERO;
  g_ThreadSwitch.ThreadId = NULL_ZERO;
  if (ThreadId == NULL_ZERO && EThread == NULL64_ZERO) {
    return FALSE;
  }
  if (EThread != NULL) {
    if (CheckAccessValidityAndSafety((UINT64)EThread, sizeof(BYTE))) {
      g_ThreadSwitch.Thread = EThread;
    } else {
      return FALSE;
    }
  } else if (ThreadId != NULL_ZERO) {
    g_ThreadSwitch.ThreadId = ThreadId;
  }
  HaltedCoreBroadcastTaskAllCores(
      DbgState, DEBUGGER_HALTED_CORE_TASK_SET_THREAD_INTERCEPTION, TRUE, TRUE,
      (PVOID)CheckByClockInterrupt);
  return TRUE;
}

BOOLEAN ThreadShowList(
    PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS ThreadListSymbolInfo,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS QueryAction,
    UINT32 *CountOfThreads, PVOID ListSaveBuffer, UINT64 ListSaveBuffSize) {
  UINT64 ThreadListHead;
  UINT32 EnumerationCount = 0;
  UINT64 Thread = (UINT64)NULL;
  LIST_ENTRY ThreadLinks = {0};
  CLIENT_ID ThreadCid = {0};
  UINT32 MaximumBufferCount = 0;
  PDEBUGGEE_THREAD_LIST_DETAILS_ENTRY SavingEntries = ListSaveBuffer;
  if (QueryAction ==
          DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT &&
      CountOfThreads == NULL) {
    return FALSE;
  }
  if (QueryAction ==
          DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS &&
      (ListSaveBuffer == NULL || ListSaveBuffSize == 0)) {
    return FALSE;
  }
  if (QueryAction ==
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS) {
    MaximumBufferCount =
        (UINT32)(ListSaveBuffSize / sizeof(DEBUGGEE_THREAD_LIST_DETAILS_ENTRY));
  }
  UINT32 ThreadListHeadOffset = ThreadListSymbolInfo->ThreadListHeadOffset;
  UINT32 ThreadListEntryOffset = ThreadListSymbolInfo->ThreadListEntryOffset;
  UINT32 CidOffset = ThreadListSymbolInfo->CidOffset;
  UINT32 ActiveProcessLinksOffset =
      ThreadListSymbolInfo->ActiveProcessLinksOffset;
  UINT64 PsActiveProcessHeadAddress = ThreadListSymbolInfo->PsActiveProcessHead;
  if (ThreadListHeadOffset == NULL_ZERO || ThreadListEntryOffset == NULL_ZERO ||
      CidOffset == NULL_ZERO || ActiveProcessLinksOffset == NULL_ZERO ||
      PsActiveProcessHeadAddress == NULL64_ZERO) {
    return FALSE;
  }
  if (ThreadListSymbolInfo->Process == NULL64_ZERO) {
    ThreadListSymbolInfo->Process = (UINT64)PsGetCurrentProcess();
    ThreadListHead = (UINT64)PsGetCurrentProcess() + ThreadListHeadOffset;
  } else {
    ThreadListHead =
        (UINT64)ThreadListSymbolInfo->Process + ThreadListHeadOffset;
  }
  if (!CheckAccessValidityAndSafety(ThreadListHead, sizeof(BYTE))) {
    return FALSE;
  }
  if (!ProcessCheckIfEprocessIsValid(ThreadListSymbolInfo->Process,
                                     PsActiveProcessHeadAddress,
                                     ActiveProcessLinksOffset)) {
    return FALSE;
  }
  if (QueryAction ==
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY) {
    Log("PROCESS\t%llx\tIMAGE\t%s\n", ThreadListSymbolInfo->Process,
        CommonGetProcessNameFromProcessControlBlock(
            (PEPROCESS)ThreadListSymbolInfo->Process));
  }
  MemoryMapperReadMemorySafe(ThreadListHead, &ThreadLinks, sizeof(ThreadLinks));
  Thread = (UINT64)ThreadLinks.Flink - ThreadListEntryOffset;
  do {
    MemoryMapperReadMemorySafe(Thread + CidOffset, &ThreadCid,
                               sizeof(ThreadCid));
    switch (QueryAction) {
      case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY:
        Log("\tTHREAD\t%llx (%llx.%llx)\n", Thread, ThreadCid.UniqueProcess,
            ThreadCid.UniqueThread);
        break;
      case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT:
        EnumerationCount++;
        break;
      case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS:
        EnumerationCount++;
        if (EnumerationCount == MaximumBufferCount - 1) {
          goto ReturnEnd;
        }
        SavingEntries[EnumerationCount - 1].Eprocess =
            ThreadListSymbolInfo->Process;
        SavingEntries[EnumerationCount - 1].ProcessId =
            HANDLE_TO_UINT32(ThreadCid.UniqueProcess);
        SavingEntries[EnumerationCount - 1].ThreadId =
            HANDLE_TO_UINT32(ThreadCid.UniqueThread);
        SavingEntries[EnumerationCount - 1].Ethread = Thread;
        MemoryMapperReadMemorySafe(
            (UINT64)CommonGetProcessNameFromProcessControlBlock(
                (PEPROCESS)ThreadListSymbolInfo->Process),
            &SavingEntries[EnumerationCount - 1].ImageFileName, 15);
        break;
      default:
        break;
    }
    MemoryMapperReadMemorySafe(Thread + ThreadListEntryOffset, &ThreadLinks,
                               sizeof(ThreadLinks));
    Thread = (UINT64)ThreadLinks.Flink - ThreadListEntryOffset;
  } while ((UINT64)ThreadLinks.Flink != ThreadListHead);
ReturnEnd:
  if (QueryAction ==
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT) {
    *CountOfThreads = EnumerationCount;
  }
  return TRUE;
}

BOOLEAN ThreadInterpretThread(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET TidRequest) {
  switch (TidRequest->ActionType) {
    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS:
      TidRequest->ProcessId = HANDLE_TO_UINT32(PsGetCurrentProcessId());
      TidRequest->ThreadId = HANDLE_TO_UINT32(PsGetCurrentThreadId());
      TidRequest->Process = (UINT64)PsGetCurrentProcess();
      TidRequest->Thread = (UINT64)PsGetCurrentThread();
      MemoryMapperReadMemorySafe(
          (UINT64)CommonGetProcessNameFromProcessControlBlock(
              PsGetCurrentProcess()),
          &TidRequest->ProcessName, 16);
      TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH:
      if (!ThreadSwitch(DbgState, TidRequest->ThreadId,
                        (PETHREAD)TidRequest->Thread,
                        TidRequest->CheckByClockInterrupt)) {
        TidRequest->Result =
            DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
        break;
      }
      TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_LIST:
      if (!ThreadShowList(
              &TidRequest->ThreadListSymDetails,
              DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY,
              NULL, NULL, (UINT64)NULL)) {
        TidRequest->Result =
            DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
        break;
      }
      TidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    default:
      TidRequest->Result =
          DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER;
      break;
  }
  if (TidRequest->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    return TRUE;
  } else {
    return FALSE;
  }
}

VOID ThreadDetectChangeByDebugRegisterOnGs(PROCESSOR_DEBUGGING_STATE *DbgState,
                                           BOOLEAN Enable) {
  UINT64 MsrGsBase;
  if (Enable) {
    MsrGsBase = __readmsr(IA32_GS_BASE);
    MsrGsBase += 0x188;
    DbgState->ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs =
        MsrGsBase;
    DbgState->ThreadOrProcessTracingDetails.DebugRegisterInterceptionState =
        TRUE;
    VmFuncSetLoadDebugControls(TRUE);
    VmFuncSetSaveDebugControls(TRUE);
    VmFuncSetExceptionBitmap(DbgState->CoreId,
                             EXCEPTION_VECTOR_DEBUG_BREAKPOINT);
    SetDebugRegisters(
        DEBUGGER_DEBUG_REGISTER_FOR_THREAD_MANAGEMENT, BREAK_ON_WRITE_ONLY,
        TRUE,
        DbgState->ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs);
    VmFuncSetMovDebugRegsExiting(DbgState->CoreId, TRUE);
  } else {
    DbgState->ThreadOrProcessTracingDetails.DebugRegisterInterceptionState =
        FALSE;
    VmFuncSetMovDebugRegsExiting(DbgState->CoreId, FALSE);
    VmFuncSetLoadDebugControls(FALSE);
    VmFuncSetSaveDebugControls(FALSE);
    VmFuncUnsetExceptionBitmap(DbgState->CoreId,
                               EXCEPTION_VECTOR_DEBUG_BREAKPOINT);
    DbgState->ThreadOrProcessTracingDetails.CurrentThreadLocationOnGs =
        (UINT64)NULL;
  }
}

VOID ThreadDetectChangeByInterceptingClockInterrupts(
    PROCESSOR_DEBUGGING_STATE *DbgState, BOOLEAN Enable) {
  if (Enable) {
    DbgState->ThreadOrProcessTracingDetails
        .InterceptClockInterruptsForThreadChange = TRUE;
    VmFuncSetExternalInterruptExiting(DbgState->CoreId, TRUE);
  } else {
    DbgState->ThreadOrProcessTracingDetails
        .InterceptClockInterruptsForThreadChange = FALSE;
    VmFuncSetExternalInterruptExiting(DbgState->CoreId, FALSE);
  }
}

VOID ThreadEnableOrDisableThreadChangeMonitor(
    PROCESSOR_DEBUGGING_STATE *DbgState, BOOLEAN Enable,
    BOOLEAN IsSwitchByClockIntrrupt) {
  if (Enable) {
    DbgState->ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent = TRUE;
    DbgState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt =
        IsSwitchByClockIntrrupt;
  } else {
    DbgState->ThreadOrProcessTracingDetails.InitialSetThreadChangeEvent = FALSE;
    DbgState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt = FALSE;
  }
  if (!IsSwitchByClockIntrrupt) {
    ThreadDetectChangeByDebugRegisterOnGs(DbgState, Enable);
  } else {
    ThreadDetectChangeByInterceptingClockInterrupts(DbgState, Enable);
  }
}

BOOLEAN ThreadQueryCount(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                             DebuggerUsermodeProcessOrThreadQueryRequest) {
  BOOLEAN Result = FALSE;
  Result = ThreadShowList(
      &DebuggerUsermodeProcessOrThreadQueryRequest->ThreadListNeededDetails,
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT,
      &DebuggerUsermodeProcessOrThreadQueryRequest->Count, NULL, (UINT64)NULL);
  if (Result && DebuggerUsermodeProcessOrThreadQueryRequest->Count != 0) {
    DebuggerUsermodeProcessOrThreadQueryRequest->Result =
        DEBUGGER_OPERATION_WAS_SUCCESSFUL;
    return TRUE;
  }
  DebuggerUsermodeProcessOrThreadQueryRequest->Result =
      DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS;
  return FALSE;
}

BOOLEAN ThreadQueryList(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                            DebuggerUsermodeProcessOrThreadQueryRequest,
                        PVOID AddressToSaveDetail, UINT32 BufferSize) {
  BOOLEAN Result = FALSE;
  Result = ThreadShowList(
      &DebuggerUsermodeProcessOrThreadQueryRequest->ThreadListNeededDetails,
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS,
      NULL, AddressToSaveDetail, BufferSize);
  return Result;
}

BOOLEAN ThreadQueryDetails(
    PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET GetInformationThreadRequest) {
  GetInformationThreadRequest->ProcessId =
      HANDLE_TO_UINT32(PsGetCurrentProcessId());
  GetInformationThreadRequest->Process = (UINT64)PsGetCurrentProcess();
  GetInformationThreadRequest->Thread = (UINT64)PsGetCurrentThread();
  GetInformationThreadRequest->ThreadId =
      HANDLE_TO_UINT32(PsGetCurrentThreadId());
  RtlCopyMemory(
      &GetInformationThreadRequest->ProcessName,
      CommonGetProcessNameFromProcessControlBlock(PsGetCurrentProcess()), 15);
  GetInformationThreadRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return TRUE;
}
