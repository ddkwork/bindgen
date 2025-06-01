
#include "pch.h"

VOID CommandEptHookHelp() {
  ShowMessages("!epthook : puts a hidden-hook EPT (hidden breakpoints).\n\n");
  ShowMessages(
      "syntax : \t!epthook [Address (hex)] [pid ProcessId (hex)] [core CoreId "
      "(hex)] "
      "[imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] [script { "
      "Script (string) }] "
      "[asm condition { Condition (assembly/hex) }] [asm code { Code "
      "(assembly/hex) }] [output {OutputName (string)}]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !epthook nt!ExAllocatePoolWithTag\n");
  ShowMessages("\t\te.g : !epthook nt!ExAllocatePoolWithTag+5\n");
  ShowMessages("\t\te.g : !epthook fffff801deadb000\n");
  ShowMessages("\t\te.g : !epthook fffff801deadb000 pid 400\n");
  ShowMessages("\t\te.g : !epthook fffff801deadb000 core 2 pid 400\n");
  ShowMessages(
      "\t\te.g : !epthook fffff801deadb000 script { printf(\"hook triggered "
      "at: %%llx\\n\", $context); }\n");
  ShowMessages(
      "\t\te.g : !epthook fffff801deadb000 asm code { nop; nop; nop }\n");
}

VOID CommandEptHook(vector<CommandToken> CommandTokens, string Command) {
  PDEBUGGER_GENERAL_EVENT_DETAIL Event = NULL;
  PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger = NULL;
  PDEBUGGER_GENERAL_ACTION ActionCustomCode = NULL;
  PDEBUGGER_GENERAL_ACTION ActionScript = NULL;
  UINT32 EventLength;
  UINT32 ActionBreakToDebuggerLength = 0;
  UINT32 ActionCustomCodeLength = 0;
  UINT32 ActionScriptLength = 0;
  BOOLEAN GetAddress = FALSE;
  UINT64 OptionalParam1 = 0;
  DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
  if (CommandTokens.size() < 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandEptHookHelp();
    return;
  }
  if (!InterpretGeneralEventAndActionsFields(
          &CommandTokens, HIDDEN_HOOK_EXEC_CC, &Event, &EventLength,
          &ActionBreakToDebugger, &ActionBreakToDebuggerLength,
          &ActionCustomCode, &ActionCustomCodeLength, &ActionScript,
          &ActionScriptLength, &EventParsingErrorCause)) {
    return;
  }
  if (Event->EventStage != VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION) {
    ShowMessages(
        "the utilization of 'post' or 'all' event calling stages is not "
        "meaningful "
        "for the hidden hook; therefore, this command does not support them\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  for (auto Section : CommandTokens) {
    if (CompareLowerCaseStrings(Section, "!epthook")) {
      continue;
    } else if (!GetAddress) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(Section),
              &OptionalParam1)) {
        ShowMessages("err, couldn't resolve error at '%s'\n\n",
                     GetCaseSensitiveStringFromCommandToken(Section).c_str());
        CommandEptHookHelp();
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger,
                                   ActionCustomCode, ActionScript);
        return;
      } else {
        GetAddress = TRUE;
      }
    } else {
      ShowMessages("unknown parameter '%s'\n\n",
                   GetCaseSensitiveStringFromCommandToken(Section).c_str());
      CommandEptHookHelp();
      FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                                 ActionScript);
      return;
    }
  }
  if (OptionalParam1 == 0) {
    ShowMessages(
        "please choose an address to put the hidden breakpoint on it\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  Event->Options.OptionalParam1 = OptionalParam1;
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
