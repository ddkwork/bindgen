
#include "pch.h"

VOID CommandModeHelp() {
  ShowMessages(
      "!mode : traps (and possibly blocks) the execution of "
      "user-mode/kernel-mode instructions.\n\n");
  ShowMessages(
      "syntax : \t!mode [Mode (string)] [pid ProcessId (hex)] [core CoreId "
      "(hex)] [imm IsImmediate (yesno)] "
      "[sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] "
      "[script { Script (string) }] "
      "[asm condition { Condition (assembly/hex) }] [asm code { Code "
      "(assembly/hex) }] [output {OutputName (string)}]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !mode u pid 1c0\n");
  ShowMessages("\t\te.g : !mode k pid 1c0\n");
  ShowMessages("\t\te.g : !mode ku pid 1c0\n");
  ShowMessages("\t\te.g : !mode ku core 2 pid 400\n");
  ShowMessages(
      "\t\te.g : !mode u pid 1c0 script { printf(\"kernel -> user transition "
      "occurred!\\n\"); }\n");
  ShowMessages("\t\te.g : !mode ku pid 1c0 asm code { nop; nop; nop }\n");
  ShowMessages("\n");
  ShowMessages(
      "note: this event applies to the target process; thus, you need to "
      "specify the process id\n");
}

VOID CommandMode(vector<CommandToken> CommandTokens, string Command) {
  PDEBUGGER_GENERAL_EVENT_DETAIL Event = NULL;
  PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger = NULL;
  PDEBUGGER_GENERAL_ACTION ActionCustomCode = NULL;
  PDEBUGGER_GENERAL_ACTION ActionScript = NULL;
  UINT32 EventLength;
  UINT32 ActionBreakToDebuggerLength = 0;
  UINT32 ActionCustomCodeLength = 0;
  UINT32 ActionScriptLength = 0;
  DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
  BOOLEAN SetMode = FALSE;
  DEBUGGER_EVENT_MODE_TYPE TargetInterceptionMode =
      DEBUGGER_EVENT_MODE_TYPE_INVALID;
  if (!InterpretGeneralEventAndActionsFields(
          &CommandTokens, TRAP_EXECUTION_MODE_CHANGED, &Event, &EventLength,
          &ActionBreakToDebugger, &ActionBreakToDebuggerLength,
          &ActionCustomCode, &ActionCustomCodeLength, &ActionScript,
          &ActionScriptLength, &EventParsingErrorCause)) {
    return;
  }
  if (Event->EventStage != VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION) {
    ShowMessages(
        "the utilization of 'post' or 'all' event calling stages is not "
        "meaningful "
        "for the mode (user-mode/kernel-mode) change traps; therefore, this "
        "command does not support them\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  if (CommandTokens.size() > 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandModeHelp();
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  for (auto Section : CommandTokens) {
    if (CompareLowerCaseStrings(Section, "!mode")) {
      continue;
    } else if (CompareLowerCaseStrings(Section, "u") && !SetMode) {
      TargetInterceptionMode = DEBUGGER_EVENT_MODE_TYPE_USER_MODE;
      SetMode = TRUE;
    } else if (CompareLowerCaseStrings(Section, "k") && !SetMode) {
      TargetInterceptionMode = DEBUGGER_EVENT_MODE_TYPE_KERNEL_MODE;
      SetMode = TRUE;
    } else if ((CompareLowerCaseStrings(Section, "uk") ||
                CompareLowerCaseStrings(Section, "ku")) &&
               !SetMode) {
      TargetInterceptionMode =
          DEBUGGER_EVENT_MODE_TYPE_USER_MODE_AND_KERNEL_MODE;
      SetMode = TRUE;
    } else {
      ShowMessages("err, couldn't resolve error at '%s'\n\n",
                   GetCaseSensitiveStringFromCommandToken(Section).c_str());
      CommandModeHelp();
      FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                                 ActionScript);
    }
  }
  if (!SetMode) {
    ShowMessages(
        "please specify the mode(s) that you want to intercept their execution "
        "(u, k, ku)\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  if (Event->ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES) {
    ShowMessages(
        "this event only applies to the selected process(es). please specify "
        "the 'pid' or the process id of the target process that you want to "
        "trap its execution\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  Event->Options.OptionalParam1 = (UINT64)TargetInterceptionMode;
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
