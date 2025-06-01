
#include "pch.h"

VOID CommandTscHelp() {
  ShowMessages("!tsc : monitors execution of rdtsc/rdtscp instructions.\n\n");
  ShowMessages(
      "syntax : \t!tsc [pid ProcessId (hex)] [core CoreId (hex)] [imm "
      "IsImmediate (yesno)] "
      "[sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] "
      "[buffer PreAllocatedBuffer (hex)] "
      "[script { Script (string) }] [asm condition { Condition (assembly/hex) "
      "}] [asm code { Code (assembly/hex) }] "
      "[output {OutputName (string)}]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !tsc\n");
  ShowMessages("\t\te.g : !tsc pid 400\n");
  ShowMessages("\t\te.g : !tsc core 2 pid 400\n");
  ShowMessages(
      "\t\te.g : !tsc script { printf(\"RDTSC/P instruction called at: "
      "%%llx\\n\", @rip); }\n");
  ShowMessages("\t\te.g : !tsc asm code { nop; nop; nop }\n");
}

VOID CommandTsc(vector<CommandToken> CommandTokens, string Command) {
  PDEBUGGER_GENERAL_EVENT_DETAIL Event = NULL;
  PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger = NULL;
  PDEBUGGER_GENERAL_ACTION ActionCustomCode = NULL;
  PDEBUGGER_GENERAL_ACTION ActionScript = NULL;
  UINT32 EventLength;
  UINT32 ActionBreakToDebuggerLength = 0;
  UINT32 ActionCustomCodeLength = 0;
  UINT32 ActionScriptLength = 0;
  DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
  if (!InterpretGeneralEventAndActionsFields(
          &CommandTokens, TSC_INSTRUCTION_EXECUTION, &Event, &EventLength,
          &ActionBreakToDebugger, &ActionBreakToDebuggerLength,
          &ActionCustomCode, &ActionCustomCodeLength, &ActionScript,
          &ActionScriptLength, &EventParsingErrorCause)) {
    return;
  }
  if (CommandTokens.size() > 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandTscHelp();
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
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
