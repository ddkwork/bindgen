
#include "pch.h"

VOID CommandMsrreadHelp() {
  ShowMessages("!msrread : detects the execution of rdmsr instructions.\n\n");
  ShowMessages(
      "syntax : \t!msrread [Msr (hex)] [pid ProcessId (hex)] "
      "[core CoreId (hex)] [imm IsImmediate (yesno)] [sc EnableShortCircuiting "
      "(onoff)] "
      "[stage CallingStage (prepostall)] [buffer PreAllocatedBuffer (hex)] "
      "[script { Script (string) }] "
      "[asm condition { Condition (assembly/hex) }] [asm code { Code "
      "(assembly/hex) }] [output {OutputName (string)}]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !msrread\n");
  ShowMessages("\t\te.g : !msrread 0xc0000082\n");
  ShowMessages("\t\te.g : !msread pid 400\n");
  ShowMessages("\t\te.g : !msrread core 2 pid 400\n");
  ShowMessages(
      "\t\te.g : !msrread script { printf(\"msr read with the 'ecx' register "
      "equal to: %%llx\\n\", $context); }\n");
  ShowMessages("\t\te.g : !msrread asm code { nop; nop; nop }\n");
}

VOID CommandMsrread(vector<CommandToken> CommandTokens, string Command) {
  PDEBUGGER_GENERAL_EVENT_DETAIL Event = NULL;
  PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger = NULL;
  PDEBUGGER_GENERAL_ACTION ActionCustomCode = NULL;
  PDEBUGGER_GENERAL_ACTION ActionScript = NULL;
  UINT32 EventLength;
  UINT32 ActionBreakToDebuggerLength = 0;
  UINT32 ActionCustomCodeLength = 0;
  UINT32 ActionScriptLength = 0;
  UINT64 SpecialTarget = DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS;
  BOOLEAN GetAddress = FALSE;
  DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
  if (!InterpretGeneralEventAndActionsFields(
          &CommandTokens, RDMSR_INSTRUCTION_EXECUTION, &Event, &EventLength,
          &ActionBreakToDebugger, &ActionBreakToDebuggerLength,
          &ActionCustomCode, &ActionCustomCodeLength, &ActionScript,
          &ActionScriptLength, &EventParsingErrorCause)) {
    return;
  }
  for (auto Section : CommandTokens) {
    if (CompareLowerCaseStrings(Section, "!msrread") ||
        CompareLowerCaseStrings(Section, "!msread")) {
      continue;
    } else if (!GetAddress) {
      if (!ConvertTokenToUInt64(Section, &SpecialTarget)) {
        ShowMessages("unknown parameter '%s'\n\n",
                     GetCaseSensitiveStringFromCommandToken(Section).c_str());
        CommandMsrreadHelp();
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger,
                                   ActionCustomCode, ActionScript);
        return;
      } else {
        GetAddress = TRUE;
      }
    } else {
      ShowMessages("unknown parameter '%s'\n\n",
                   GetCaseSensitiveStringFromCommandToken(Section).c_str());
      CommandMsrreadHelp();
      FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                                 ActionScript);
      return;
    }
  }
  Event->Options.OptionalParam1 = SpecialTarget;
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
