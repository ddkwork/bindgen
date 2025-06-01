
#include "pch.h"

VOID CommandTraceHelp() {
  ShowMessages(
      "!trace : traces the execution of user-mode/kernel-mode "
      "instructions.\n\n");
  ShowMessages(
      "syntax : \t!trace [TraceType (string)] [pid ProcessId (hex)] [core "
      "CoreId (hex)] [imm IsImmediate (yesno)] "
      "[sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] "
      "[script { Script (string) }] "
      "[asm condition { Condition (assembly/hex) }] [asm code { Code "
      "(assembly/hex) }] [output {OutputName (string)}]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !trace step-out\n");
  ShowMessages("\t\te.g : !trace step-in pid 1c0\n");
  ShowMessages("\t\te.g : !trace instrument-step core 2 pid 400\n");
  ShowMessages("\t\te.g : !trace instrument-step asm code { nop; nop; nop }\n");
  ShowMessages("\n");
  ShowMessages("valid trace types: \n");
  ShowMessages("\tstep-in : single step-in (regular)\n");
  ShowMessages("\tstep-out : single step-out (regular)\n");
  ShowMessages("\tinstrument-step : single step-in (instrumentation)\n");
}

VOID CommandTrace(vector<CommandToken> CommandTokens, string Command) {
  PDEBUGGER_GENERAL_EVENT_DETAIL Event = NULL;
  PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger = NULL;
  PDEBUGGER_GENERAL_ACTION ActionCustomCode = NULL;
  PDEBUGGER_GENERAL_ACTION ActionScript = NULL;
  UINT32 EventLength;
  UINT32 ActionBreakToDebuggerLength = 0;
  UINT32 ActionCustomCodeLength = 0;
  UINT32 ActionScriptLength = 0;
  DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
  BOOLEAN SetTraceType = FALSE;
  DEBUGGER_EVENT_TRACE_TYPE TargetTrace = DEBUGGER_EVENT_TRACE_TYPE_INVALID;
  if (!InterpretGeneralEventAndActionsFields(
          &CommandTokens, TRAP_EXECUTION_INSTRUCTION_TRACE, &Event,
          &EventLength, &ActionBreakToDebugger, &ActionBreakToDebuggerLength,
          &ActionCustomCode, &ActionCustomCodeLength, &ActionScript,
          &ActionScriptLength, &EventParsingErrorCause)) {
    return;
  }
  if (CommandTokens.size() > 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandTraceHelp();
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  for (auto Section : CommandTokens) {
    if (CompareLowerCaseStrings(Section, "!trace")) {
      continue;
    } else if ((CompareLowerCaseStrings(Section, "step-in") ||
                CompareLowerCaseStrings(Section, "stepin") ||
                CompareLowerCaseStrings(Section, "step")) &&
               !SetTraceType) {
      TargetTrace = DEBUGGER_EVENT_TRACE_TYPE_STEP_IN;
      SetTraceType = TRUE;
    } else if ((CompareLowerCaseStrings(Section, "step-out") ||
                CompareLowerCaseStrings(Section, "stepout")) &&
               !SetTraceType) {
      TargetTrace = DEBUGGER_EVENT_TRACE_TYPE_STEP_OUT;
      SetTraceType = TRUE;
    } else if ((CompareLowerCaseStrings(Section, "step-instrument") ||
                CompareLowerCaseStrings(Section, "instrument-step") ||
                CompareLowerCaseStrings(Section, "instrumentstep") ||
                CompareLowerCaseStrings(Section, "instrument-step-in")) &&
               !SetTraceType) {
      TargetTrace = DEBUGGER_EVENT_TRACE_TYPE_INSTRUMENTATION_STEP_IN;
      SetTraceType = TRUE;
    } else {
      ShowMessages("err, couldn't resolve error at '%s'\n\n",
                   GetCaseSensitiveStringFromCommandToken(Section).c_str());
      CommandTraceHelp();
      FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                                 ActionScript);
    }
  }
  if (!SetTraceType) {
    ShowMessages("please specify the trace type\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  Event->Options.OptionalParam1 = (UINT64)SetTraceType;
  if (!SendEventToKernel(Event, EventLength)) {
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  if (!RegisterActionToEvent(Event, ActionBreakToDebugger,
                             ActionBreakToDebuggerLength, ActionCustomCode,
                             ActionCustomCodeLength, ActionScript,
                             ActionScriptLength)) {
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
}
