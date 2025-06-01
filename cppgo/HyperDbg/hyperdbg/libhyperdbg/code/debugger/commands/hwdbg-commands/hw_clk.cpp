
#include "pch.h"
extern HWDBG_INSTANCE_INFORMATION g_HwdbgInstanceInfo;

VOID CommandHwClkHelp() {
  ShowMessages(
      "!hw_clk : performs actions related to hwdbg hardware debugging events "
      "for each clock cycle.\n\n");
  ShowMessages("syntax : \t!hw_clk  [script { Script (string) }]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !hw_clk script { @hw_pin1 = 0; }\n");
}

BOOLEAN CommandHwClkPerfomTest(vector<CommandToken> CommandTokens,
                               const TCHAR *InstanceFilePathToRead,
                               const TCHAR *InstanceFilePathToSave,
                               const TCHAR *HardwareScriptFilePathToSave,
                               UINT32 InitialBramBufferSize) {
  UINT32 EventLength;
  DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
  PDEBUGGER_GENERAL_EVENT_DETAIL Event = NULL;
  PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger = NULL;
  PDEBUGGER_GENERAL_ACTION ActionCustomCode = NULL;
  PDEBUGGER_GENERAL_ACTION ActionScript = NULL;
  UINT32 ActionBreakToDebuggerLength = 0;
  UINT32 ActionCustomCodeLength = 0;
  UINT32 ActionScriptLength = 0;
  CHAR *ScriptBuffer = NULL;
  BOOLEAN Result = FALSE;
  if (!HwdbgLoadInstanceInfo(InstanceFilePathToRead, InitialBramBufferSize)) {
    return FALSE;
  }
  if (!InterpretGeneralEventAndActionsFields(
          &CommandTokens, (VMM_EVENT_TYPE_ENUM)NULL, &Event, &EventLength,
          &ActionBreakToDebugger, &ActionBreakToDebuggerLength,
          &ActionCustomCode, &ActionCustomCodeLength, &ActionScript,
          &ActionScriptLength, &EventParsingErrorCause)) {
    return FALSE;
  }
  ScriptBuffer =
      (CHAR *)((UINT64)ActionScript + sizeof(DEBUGGER_GENERAL_ACTION));
  HwdbgScriptPrintScriptBuffer(ScriptBuffer, ActionScript->ScriptBufferSize);
  if (!HwdbgScriptCreateHwdbgScript(ScriptBuffer,
                                    ActionScript->ScriptBufferSize,
                                    HardwareScriptFilePathToSave)) {
    ShowMessages("err, unable to create hwdbg script\n");
    Result = FALSE;
    goto FreeAndReturnResult;
  }
  if (!HwdbgWriteTestInstanceInfoRequestIntoFile(&g_HwdbgInstanceInfo,
                                                 InstanceFilePathToSave)) {
    ShowMessages("err, unable to write instance info request\n");
    Result = FALSE;
    goto FreeAndReturnResult;
  }
  Result = TRUE;
FreeAndReturnResult:
  FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode,
                             ActionScript);
  return Result;
}

VOID CommandHwClk(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() >= 2) {
    CommandHwClkPerfomTest(CommandTokens, HWDBG_TEST_READ_INSTANCE_INFO_PATH,
                           HWDBG_TEST_WRITE_INSTANCE_INFO_PATH,
                           HWDBG_TEST_WRITE_SCRIPT_BUFFER_PATH,
                           DEFAULT_INITIAL_BRAM_BUFFER_SIZE);
  } else {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandHwClkHelp();
    return;
  }
}
