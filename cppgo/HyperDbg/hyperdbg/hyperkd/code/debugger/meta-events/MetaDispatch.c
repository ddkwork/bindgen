
#include "pch.h"

VOID MetaDispatchEventInstrumentationTrace(
    PROCESSOR_DEBUGGING_STATE *DbgState) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  EventTriggerResult =
      DebuggerTriggerEvents(TRAP_EXECUTION_INSTRUCTION_TRACE,
                            VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                            (PVOID)DEBUGGER_EVENT_TRACE_TYPE_STEP_IN,
                            &PostEventTriggerReq, DbgState->Regs);
}
