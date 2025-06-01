
#include "pch.h"

VOID ProcessTriggerCr3ProcessChange(UINT32 CoreId) {
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  if (DbgState->ThreadOrProcessTracingDetails.IsWatingForMovCr3VmExits) {
    ProcessHandleProcessChange(DbgState);
  }
}

BOOLEAN ProcessHandleProcessChange(PROCESSOR_DEBUGGING_STATE *DbgState) {
  if ((g_ProcessSwitch.ProcessId != NULL_ZERO &&
       g_ProcessSwitch.ProcessId ==
           HANDLE_TO_UINT32(PsGetCurrentProcessId())) ||
      (g_ProcessSwitch.Process != (UINT64)NULL &&
       g_ProcessSwitch.Process == PsGetCurrentProcess())) {
    KdHandleBreakpointAndDebugBreakpoints(
        DbgState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED, NULL);
    return TRUE;
  }
  return FALSE;
}

BOOLEAN ProcessSwitch(PROCESSOR_DEBUGGING_STATE *DbgState, UINT32 ProcessId,
                      PEPROCESS EProcess, BOOLEAN IsSwitchByClockIntrrupt) {
  g_ProcessSwitch.Process = NULL;
  g_ProcessSwitch.ProcessId = NULL_ZERO;
  if (ProcessId == NULL_ZERO && EProcess == (PEPROCESS)NULL) {
    return FALSE;
  }
  if (EProcess != NULL) {
    if (CheckAccessValidityAndSafety((UINT64)EProcess, sizeof(BYTE))) {
      g_ProcessSwitch.Process = (PVOID)EProcess;
    } else {
      return FALSE;
    }
  } else if (ProcessId != NULL_ZERO) {
    g_ProcessSwitch.ProcessId = ProcessId;
  }
  HaltedCoreBroadcastTaskAllCores(
      DbgState, DEBUGGER_HALTED_CORE_TASK_SET_PROCESS_INTERCEPTION, TRUE, TRUE,
      (PVOID)IsSwitchByClockIntrrupt);
  return TRUE;
}

VOID ProcessDetectChangeByInterceptingClockInterrupts(
    PROCESSOR_DEBUGGING_STATE *DbgState, BOOLEAN Enable) {
  if (Enable) {
    DbgState->ThreadOrProcessTracingDetails
        .InterceptClockInterruptsForProcessChange = TRUE;
    VmFuncSetExternalInterruptExiting(DbgState->CoreId, TRUE);
  } else {
    DbgState->ThreadOrProcessTracingDetails
        .InterceptClockInterruptsForProcessChange = FALSE;
    VmFuncSetExternalInterruptExiting(DbgState->CoreId, FALSE);
  }
}

VOID ProcessDetectChangeByMov2Cr3Vmexits(PROCESSOR_DEBUGGING_STATE *DbgState,
                                         BOOLEAN Enable) {
  if (Enable) {
    DbgState->ThreadOrProcessTracingDetails.IsWatingForMovCr3VmExits = TRUE;
    VmFuncSetMovToCr3Vmexit(DbgState->CoreId, TRUE);
  } else {
    DbgState->ThreadOrProcessTracingDetails.IsWatingForMovCr3VmExits = FALSE;
    VmFuncSetMovToCr3Vmexit(DbgState->CoreId, FALSE);
  }
}

VOID ProcessEnableOrDisableThreadChangeMonitor(
    PROCESSOR_DEBUGGING_STATE *DbgState, BOOLEAN Enable,
    BOOLEAN IsSwitchByClockIntrrupt) {
  if (Enable) {
    DbgState->ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent = TRUE;
    DbgState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt =
        IsSwitchByClockIntrrupt;
  } else {
    DbgState->ThreadOrProcessTracingDetails.InitialSetProcessChangeEvent =
        FALSE;
    DbgState->ThreadOrProcessTracingDetails.InitialSetByClockInterrupt = FALSE;
  }
  if (!IsSwitchByClockIntrrupt) {
    ProcessDetectChangeByMov2Cr3Vmexits(DbgState, Enable);
  } else {
    ProcessDetectChangeByInterceptingClockInterrupts(DbgState, Enable);
  }
}

BOOLEAN ProcessCheckIfEprocessIsValid(UINT64 Eprocess, UINT64 ActiveProcessHead,
                                      ULONG ActiveProcessLinksOffset) {
  UINT64 Process;
  LIST_ENTRY ActiveProcessLinks = {0};
  if (ActiveProcessHead == NULL64_ZERO ||
      ActiveProcessLinksOffset == NULL_ZERO) {
    return FALSE;
  }
  if (CheckAccessValidityAndSafety(ActiveProcessHead, sizeof(BYTE))) {
    MemoryMapperReadMemorySafe(ActiveProcessHead, &ActiveProcessLinks,
                               sizeof(ActiveProcessLinks));
    Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
    do {
      MemoryMapperReadMemorySafe(Process + ActiveProcessLinksOffset,
                                 &ActiveProcessLinks,
                                 sizeof(ActiveProcessLinks));
      if (Process == Eprocess) {
        return TRUE;
      }
      Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
    } while ((UINT64)ActiveProcessLinks.Flink != ActiveProcessHead);
  } else {
    return FALSE;
  }
  return FALSE;
}

BOOLEAN ProcessShowList(
    PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS PorcessListSymbolInfo,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS QueryAction,
    UINT32 *CountOfProcesses, PVOID ListSaveBuffer, UINT64 ListSaveBuffSize) {
  UINT64 Process;
  UINT64 UniquePid = 0;
  LIST_ENTRY ActiveProcessLinks = {0};
  UCHAR ImageFileName[15] = {0};
  CR3_TYPE ProcessCr3 = {0};
  UINT32 EnumerationCount = 0;
  UINT32 MaximumBufferCount = 0;
  PDEBUGGEE_PROCESS_LIST_DETAILS_ENTRY SavingEntries = ListSaveBuffer;
  if (QueryAction ==
          DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT &&
      CountOfProcesses == NULL) {
    return FALSE;
  }
  if (QueryAction ==
          DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS &&
      (ListSaveBuffer == NULL || ListSaveBuffSize == 0)) {
    return FALSE;
  }
  if (QueryAction ==
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS) {
    MaximumBufferCount = (UINT32)(ListSaveBuffSize /
                                  sizeof(DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY));
  }
  UINT64 ActiveProcessHead = PorcessListSymbolInfo->PsActiveProcessHead;
  ULONG ImageFileNameOffset = PorcessListSymbolInfo->ImageFileNameOffset;
  ULONG UniquePidOffset = PorcessListSymbolInfo->UniquePidOffset;
  ULONG ActiveProcessLinksOffset =
      PorcessListSymbolInfo->ActiveProcessLinksOffset;
  if (ActiveProcessHead == NULL64_ZERO || ImageFileNameOffset == NULL_ZERO ||
      UniquePidOffset == NULL_ZERO || ActiveProcessLinksOffset == NULL_ZERO) {
    return FALSE;
  }
  if (CheckAccessValidityAndSafety(ActiveProcessHead, sizeof(BYTE))) {
    MemoryMapperReadMemorySafe(ActiveProcessHead, &ActiveProcessLinks,
                               sizeof(ActiveProcessLinks));
    Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
    do {
      MemoryMapperReadMemorySafe(Process + ImageFileNameOffset, &ImageFileName,
                                 sizeof(ImageFileName));
      MemoryMapperReadMemorySafe(Process + UniquePidOffset, &UniquePid,
                                 sizeof(UniquePid));
      MemoryMapperReadMemorySafe(Process + ActiveProcessLinksOffset,
                                 &ActiveProcessLinks,
                                 sizeof(ActiveProcessLinks));
      NT_KPROCESS *CurrentProcess = (NT_KPROCESS *)(Process);
      ProcessCr3.Flags = CurrentProcess->DirectoryTableBase;
      switch (QueryAction) {
        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY:
          Log("PROCESS\t%llx\n\tProcess Id: %04x\tDirBase (Kernel Cr3): "
              "%016llx\tImage: %s\n\n",
              Process, UniquePid, ProcessCr3.Flags, ImageFileName);
        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT:
          EnumerationCount++;
          break;
        case DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS:
          EnumerationCount++;
          if (EnumerationCount == MaximumBufferCount - 1) {
            goto ReturnEnd;
          }
          SavingEntries[EnumerationCount - 1].Eprocess = Process;
          SavingEntries[EnumerationCount - 1].ProcessId = (UINT32)UniquePid;
          SavingEntries[EnumerationCount - 1].Cr3 = ProcessCr3.Flags;
          RtlCopyMemory(&SavingEntries[EnumerationCount - 1].ImageFileName,
                        ImageFileName, 15);
          break;
        default:
          LogError("Err, invalid action specified for process enumeration");
          break;
      }
      Process = (UINT64)ActiveProcessLinks.Flink - ActiveProcessLinksOffset;
    } while ((UINT64)ActiveProcessLinks.Flink != ActiveProcessHead);
  } else {
    return FALSE;
  }
ReturnEnd:
  if (QueryAction ==
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT) {
    *CountOfProcesses = EnumerationCount;
  }
  return TRUE;
}

BOOLEAN ProcessInterpretProcess(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET PidRequest) {
  switch (PidRequest->ActionType) {
    case DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS:
      PidRequest->ProcessId = HANDLE_TO_UINT32(PsGetCurrentProcessId());
      PidRequest->Process = (UINT64)PsGetCurrentProcess();
      MemoryMapperReadMemorySafe(
          (UINT64)CommonGetProcessNameFromProcessControlBlock(
              PsGetCurrentProcess()),
          &PidRequest->ProcessName, 16);
      PidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH:
      if (!ProcessSwitch(DbgState, PidRequest->ProcessId,
                         (PEPROCESS)PidRequest->Process,
                         PidRequest->IsSwitchByClkIntr)) {
        PidRequest->Result =
            DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
        break;
      }
      PidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    case DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_LIST:
      if (!ProcessShowList(
              &PidRequest->ProcessListSymDetails,
              DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY,
              NULL, NULL, (UINT64)NULL)) {
        PidRequest->Result =
            DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
        break;
      }
      PidRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      break;
    default:
      PidRequest->Result =
          DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER;
      break;
  }
  if (PidRequest->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    return TRUE;
  } else {
    return FALSE;
  }
}

BOOLEAN ProcessQueryCount(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                              DebuggerUsermodeProcessOrThreadQueryRequest) {
  BOOLEAN Result = FALSE;
  Result = ProcessShowList(
      &DebuggerUsermodeProcessOrThreadQueryRequest->ProcessListNeededDetails,
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

BOOLEAN ProcessQueryList(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                             DebuggerUsermodeProcessOrThreadQueryRequest,
                         PVOID AddressToSaveDetail, UINT32 BufferSize) {
  BOOLEAN Result = FALSE;
  Result = ProcessShowList(
      &DebuggerUsermodeProcessOrThreadQueryRequest->ProcessListNeededDetails,
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS,
      NULL, AddressToSaveDetail, BufferSize);
  return Result;
}

BOOLEAN ProcessQueryDetails(
    PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET GetInformationProcessRequest) {
  GetInformationProcessRequest->ProcessId =
      HANDLE_TO_UINT32(PsGetCurrentProcessId());
  GetInformationProcessRequest->Process = (UINT64)PsGetCurrentProcess();
  RtlCopyMemory(
      &GetInformationProcessRequest->ProcessName,
      CommonGetProcessNameFromProcessControlBlock(PsGetCurrentProcess()), 15);
  GetInformationProcessRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return TRUE;
}
