
#include "pch.h"

VOID CommandMonitorHelp() {
  ShowMessages("!monitor : monitors address range for read and writes.\n\n");
  ShowMessages(
      "syntax : \t!monitor [MemoryType (vapa)] [Attribute (string)] "
      "[FromAddress (hex)] "
      "[ToAddress (hex)] [pid ProcessId (hex)] [core CoreId (hex)] "
      "[imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage "
      "CallingStage (prepostall)] "
      "[buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [asm "
      "condition { Condition (assembly/hex) }] "
      "[asm code { Code (assembly/hex) }] [output {OutputName (string)}]\n");
  ShowMessages(
      "syntax : \t!monitor [MemoryType (vapa)] [Attribute (string)] "
      "[FromAddress (hex)] "
      "[l Length (hex)] [pid ProcessId (hex)] [core CoreId (hex)] "
      "[imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage "
      "CallingStage (prepostall)] "
      "[buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [asm "
      "condition { Condition (assembly/hex) }] "
      "[asm code { Code (assembly/hex) }] [output {OutputName (string)}]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !monitor rw fffff801deadb000 fffff801deadbfff\n");
  ShowMessages("\t\te.g : !monitor rw fffff801deadb000 l 1000\n");
  ShowMessages("\t\te.g : !monitor pa rw c01000 l 1000\n");
  ShowMessages("\t\te.g : !monitor rwx fffff801deadb000 fffff801deadbfff\n");
  ShowMessages("\t\te.g : !monitor rwx fffff801deadb000 l 230d0\n");
  ShowMessages("\t\te.g : !monitor rw nt!Kd_DEFAULT_Mask Kd_DEFAULT_Mask+5\n");
  ShowMessages(
      "\t\te.g : !monitor r fffff801deadb000 fffff801deadbfff pid 400\n");
  ShowMessages(
      "\t\te.g : !monitor w fffff801deadb000 fffff801deadbfff core 2 pid "
      "400\n");
  ShowMessages("\t\te.g : !monitor w c01000 c01000+2500 core 2 pid 400\n");
  ShowMessages(
      "\t\te.g : !monitor x fffff801deadb000 fffff801deadbfff core 2 pid "
      "400\n");
  ShowMessages("\t\te.g : !monitor x fffff801deadb000 l 500 core 2 pid 400\n");
  ShowMessages(
      "\t\te.g : !monitor wx fffff801deadb000 fffff801deadbfff core 2 pid "
      "400\n");
  ShowMessages(
      "\t\te.g : !monitor rw fffff801deadb000 l 1000 script { "
      "printf(\"read/write occurred at the virtual address: %%llx\\n\", "
      "$context); }\n");
  ShowMessages(
      "\t\te.g : !monitor rw fffff801deadb000 l 1000 asm code { nop; nop; nop "
      "}\n");
}

VOID CommandMonitor(vector<CommandToken> CommandTokens, string Command) {
  PDEBUGGER_GENERAL_EVENT_DETAIL Event = NULL;
  PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger = NULL;
  PDEBUGGER_GENERAL_ACTION ActionCustomCode = NULL;
  PDEBUGGER_GENERAL_ACTION ActionScript = NULL;
  UINT32 EventLength;
  UINT32 ActionBreakToDebuggerLength = 0;
  UINT32 ActionCustomCodeLength = 0;
  UINT32 ActionScriptLength = 0;
  UINT32 HookLength = 0;
  UINT64 OptionalParam1 = 0;
  UINT64 OptionalParam2 = 0;
  BOOLEAN SetFrom = FALSE;
  BOOLEAN SetTo = FALSE;
  BOOLEAN IsNextLength = FALSE;
  BOOLEAN LengthAlreadySet = FALSE;
  BOOLEAN SetAttributes = FALSE;
  BOOLEAN HookMemoryTypeSet = FALSE;
  DEBUGGER_HOOK_MEMORY_TYPE HookMemoryType =
      DEBUGGER_MEMORY_HOOK_VIRTUAL_ADDRESS;
  DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
  if (CommandTokens.size() < 4) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandMonitorHelp();
    return;
  }
  if (!InterpretGeneralEventAndActionsFields(
          &CommandTokens, HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE, &Event,
          &EventLength, &ActionBreakToDebugger, &ActionBreakToDebuggerLength,
          &ActionCustomCode, &ActionCustomCodeLength, &ActionScript,
          &ActionScriptLength, &EventParsingErrorCause)) {
    return;
  }
  for (auto Section : CommandTokens) {
    if (CompareLowerCaseStrings(Section, "!monitor")) {
      continue;
    } else if (IsNextLength) {
      if (!ConvertTokenToUInt32(Section, &HookLength)) {
        ShowMessages("err, you should enter a valid length\n\n");
        return;
      }
      IsNextLength = FALSE;
      LengthAlreadySet = TRUE;
      SetTo = TRUE;
    } else if (CompareLowerCaseStrings(Section, "r") && !SetAttributes) {
      Event->EventType = HIDDEN_HOOK_READ;
      SetAttributes = TRUE;
    } else if (CompareLowerCaseStrings(Section, "w") && !SetAttributes) {
      Event->EventType = HIDDEN_HOOK_WRITE;
      SetAttributes = TRUE;
    } else if (CompareLowerCaseStrings(Section, "x") && !SetAttributes) {
      Event->EventType = HIDDEN_HOOK_EXECUTE;
      SetAttributes = TRUE;
    } else if ((CompareLowerCaseStrings(Section, "rw") ||
                CompareLowerCaseStrings(Section, "wr")) &&
               !SetAttributes) {
      Event->EventType = HIDDEN_HOOK_READ_AND_WRITE;
      SetAttributes = TRUE;
    } else if ((CompareLowerCaseStrings(Section, "rx") ||
                CompareLowerCaseStrings(Section, "xr")) &&
               !SetAttributes) {
      Event->EventType = HIDDEN_HOOK_READ_AND_EXECUTE;
      SetAttributes = TRUE;
    } else if ((CompareLowerCaseStrings(Section, "wx") ||
                CompareLowerCaseStrings(Section, "xw")) &&
               !SetAttributes) {
      Event->EventType = HIDDEN_HOOK_WRITE_AND_EXECUTE;
      SetAttributes = TRUE;
    } else if ((CompareLowerCaseStrings(Section, "rwx") ||
                CompareLowerCaseStrings(Section, "rxw") ||
                CompareLowerCaseStrings(Section, "wrx") ||
                CompareLowerCaseStrings(Section, "wxr") ||
                CompareLowerCaseStrings(Section, "xrw") ||
                CompareLowerCaseStrings(Section, "xwr")) &&
               !SetAttributes) {
      Event->EventType = HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE;
      SetAttributes = TRUE;
    } else if (CompareLowerCaseStrings(Section, "l") && !SetTo &&
               !LengthAlreadySet) {
      IsNextLength = TRUE;
      continue;
    } else if (CompareLowerCaseStrings(Section, "va") && !HookMemoryTypeSet) {
      HookMemoryType = DEBUGGER_MEMORY_HOOK_VIRTUAL_ADDRESS;
      HookMemoryTypeSet = TRUE;
      continue;
    } else if (CompareLowerCaseStrings(Section, "pa") && !HookMemoryTypeSet) {
      HookMemoryType = DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS;
      HookMemoryTypeSet = TRUE;
      continue;
    } else {
      if (!SetFrom) {
        if (!SymbolConvertNameOrExprToAddress(
                GetCaseSensitiveStringFromCommandToken(Section),
                &OptionalParam1)) {
          ShowMessages("err, couldn't resolve error at '%s'\n\n",
                       GetCaseSensitiveStringFromCommandToken(Section).c_str());
          CommandMonitorHelp();
          FreeEventsAndActionsMemory(Event, ActionBreakToDebugger,
                                     ActionCustomCode, ActionScript);
          return;
        }
        SetFrom = TRUE;
      } else if (!SetTo && !LengthAlreadySet) {
        if (!SymbolConvertNameOrExprToAddress(
                GetCaseSensitiveStringFromCommandToken(Section),
                &OptionalParam2)) {
          ShowMessages("err, couldn't resolve error at '%s'\n\n",
                       GetCaseSensitiveStringFromCommandToken(Section).c_str());
          CommandMonitorHelp();
          FreeEventsAndActionsMemory(Event, ActionBreakToDebugger,
                                     ActionCustomCode, ActionScript);
          return;
        }
        SetTo = TRUE;
      } else {
        ShowMessages("unknown parameter '%s'\n\n",
                     GetCaseSensitiveStringFromCommandToken(Section).c_str());
        CommandMonitorHelp();
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger,
                                   ActionCustomCode, ActionScript);
        return;
      }
    }
  }
  if (!SetFrom || !SetTo) {
    ShowMessages(
        "please choose the 'from' or 'to' values or specify the length\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  if (LengthAlreadySet) {
    OptionalParam2 = OptionalParam1 + HookLength - 1;
  }
  if (OptionalParam1 > OptionalParam2) {
    ShowMessages(
        "please choose the 'from' value first, then choose the 'to' "
        "value\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  if (!SetAttributes) {
    ShowMessages(
        "please specify the attribute(s) that you want to monitor (r, w, x, "
        "rw, rx, wx, rwx)\n");
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                               ActionScript);
    return;
  }
  Event->Options.OptionalParam1 = OptionalParam1;
  Event->Options.OptionalParam2 = OptionalParam2;
  Event->Options.OptionalParam3 = HookMemoryType;
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
