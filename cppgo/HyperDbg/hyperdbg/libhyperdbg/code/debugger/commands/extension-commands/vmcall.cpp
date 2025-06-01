
#include "pch.h"

VOID CommandVmcallHelp() {
  ShowMessages("!vmcall : monitors execution of VMCALL instruction.\n\n");
  ShowMessages(
      "syntax : \t!vmcall [pid ProcessId (hex)] [core CoreId (hex)] [imm "
      "IsImmediate (yesno)] "
      "[sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] "
      "[buffer PreAllocatedBuffer (hex)] "
      "[script { Script (string) }] [asm condition { Condition (assembly/hex) "
      "}] [asm code { Code (assembly/hex) }] "
      "[output {OutputName (string)}]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !vmcall\n");
  ShowMessages("\t\te.g : !vmcall pid 400\n");
  ShowMessages("\t\te.g : !vmcall core 2 pid 400\n");
  ShowMessages(
      "\t\te.g : !vmcall script { printf(\"VMCALL executed with context: "
      "%%llx\\n\", $context); }\n");
  ShowMessages("\t\te.g : !vmcall asm code { nop; nop; nop }\n");
}

VOID CommandVmcall(vector<CommandToken> CommandTokens, string Command) {
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
          &CommandTokens, VMCALL_INSTRUCTION_EXECUTION, &Event, &EventLength,
          &ActionBreakToDebugger, &ActionBreakToDebuggerLength,
          &ActionCustomCode, &ActionCustomCodeLength, &ActionScript,
          &ActionScriptLength, &EventParsingErrorCause)) {
    return;
  }
  if (CommandTokens.size() > 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandVmcallHelp();
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
