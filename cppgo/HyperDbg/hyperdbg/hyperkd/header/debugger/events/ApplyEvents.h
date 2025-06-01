
#pragma once
BOOLEAN ApplyEventMonitorEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
BOOLEAN ApplyEventEptHookExecCcEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
BOOLEAN ApplyEventEpthookInlineEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
BOOLEAN ApplyEventTrapModeChangeEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventRdmsrExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventWrmsrExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventInOutExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventTscExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventRdpmcExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventMov2DebugRegExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventControlRegisterAccessedEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventExceptionEvent(PDEBUGGER_EVENT Event,
                              PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                              BOOLEAN InputFromVmxRoot);
VOID ApplyEventInterruptEvent(PDEBUGGER_EVENT Event,
                              PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                              BOOLEAN InputFromVmxRoot);
VOID ApplyEventEferSyscallHookEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventEferSysretHookEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventVmcallExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventCpuidExecutionEvent(
    PDEBUGGER_EVENT Event, PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
VOID ApplyEventTracingEvent(PDEBUGGER_EVENT Event,
                            PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                            BOOLEAN InputFromVmxRoot);
