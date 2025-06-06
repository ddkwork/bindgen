
#pragma once
VOID ProcessEnableOrDisableThreadChangeMonitor(
    PROCESSOR_DEBUGGING_STATE *DbgState, BOOLEAN Enable,
    BOOLEAN IsSwitchByClockIntrrupt);
VOID ProcessTriggerCr3ProcessChange(UINT32 CoreId);
BOOLEAN ProcessHandleProcessChange(PROCESSOR_DEBUGGING_STATE *DbgState);
BOOLEAN ProcessInterpretProcess(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET PidRequest);
BOOLEAN ProcessCheckIfEprocessIsValid(UINT64 Eprocess, UINT64 ActiveProcessHead,
                                      ULONG ActiveProcessLinksOffset);
BOOLEAN ProcessQueryCount(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                              DebuggerUsermodeProcessOrThreadQueryRequest);
BOOLEAN ProcessQueryList(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                             DebuggerUsermodeProcessOrThreadQueryRequest,
                         PVOID AddressToSaveDetail, UINT32 BufferSize);
BOOLEAN ProcessQueryDetails(
    PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET GetInformationProcessRequest);
