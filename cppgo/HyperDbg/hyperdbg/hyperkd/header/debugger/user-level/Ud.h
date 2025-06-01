
#pragma once
BOOLEAN UdInitializeUserDebugger();
VOID UdUninitializeUserDebugger();
BOOLEAN UdCheckAndHandleBreakpointsAndDebugBreaks(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGEE_PAUSING_REASON Reason,
    PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails);
BOOLEAN UdDispatchUsermodeCommands(PDEBUGGER_UD_COMMAND_PACKET ActionRequest);
BOOLEAN UdCheckForCommand();
