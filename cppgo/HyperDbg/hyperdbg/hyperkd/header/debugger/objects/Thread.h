
#pragma once
BOOLEAN ThreadInterpretThread(
    PROCESSOR_DEBUGGING_STATE *DbgState,
    PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET TidRequest);
VOID ThreadEnableOrDisableThreadChangeMonitor(
    PROCESSOR_DEBUGGING_STATE *DbgState, BOOLEAN Enable,
    BOOLEAN IsSwitchByClockIntrrupt);
BOOLEAN ThreadHandleThreadChange(PROCESSOR_DEBUGGING_STATE *DbgState);
BOOLEAN ThreadQueryCount(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                             DebuggerUsermodeProcessOrThreadQueryRequest);
BOOLEAN ThreadQueryList(PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
                            DebuggerUsermodeProcessOrThreadQueryRequest,
                        PVOID AddressToSaveDetail, UINT32 BufferSize);
BOOLEAN ThreadQueryDetails(
    PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET GetInformationThreadRequest);
