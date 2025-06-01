
#pragma once

volatile LONG BreakpointCommandTrapListLock;
VOID BreakpointRemoveAllBreakpoints();
BOOLEAN BreakpointAddNew(PDEBUGGEE_BP_PACKET BpDescriptorArg);
BOOLEAN BreakpointListOrModify(
    PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET ListOrModifyBreakpoints);
BOOLEAN BreakpointHandleBreakpoints(UINT32 CoreId);
BOOLEAN BreakpointCheckAndHandleDebuggerDefinedBreakpoints(
    PROCESSOR_DEBUGGING_STATE *DbgState, UINT64 GuestRip,
    DEBUGGEE_PAUSING_REASON Reason, BOOLEAN ChangeMtfState);
BOOLEAN BreakpointCheckAndHandleReApplyingBreakpoint(UINT32 CoreId);
BOOLEAN BreakpointCheckAndHandleDebugBreakpoint(UINT32 CoreId);
BOOLEAN BreakpointRestoreTheTrapFlagOnceTriggered(UINT32 ProcessId,
                                                  UINT32 ThreadId);
