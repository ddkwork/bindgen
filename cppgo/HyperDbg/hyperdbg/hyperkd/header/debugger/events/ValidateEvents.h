
#pragma once
BOOLEAN ValidateEventMonitor(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
                             PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                             BOOLEAN InputFromVmxRoot);
BOOLEAN ValidateEventException(
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
BOOLEAN ValidateEventInterrupt(
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
BOOLEAN ValidateEventTrapExec(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
                              PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                              BOOLEAN InputFromVmxRoot);
BOOLEAN ValidateEventEptHookHiddenBreakpointAndInlineHooks(
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
