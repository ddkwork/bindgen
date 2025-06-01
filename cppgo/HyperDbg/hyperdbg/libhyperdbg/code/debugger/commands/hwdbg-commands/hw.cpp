
#include "pch.h"
extern HWDBG_INSTANCE_INFORMATION g_HwdbgInstanceInfo;
extern BOOLEAN g_HwdbgInstanceInfoIsValid;

VOID CommandHwHelp() {
  ShowMessages("!hw : runs a hardware script in the target device.\n\n");
  ShowMessages("syntax : \t!hw script [script { Script (string) }]\n");
  ShowMessages("syntax : \t!hw script [unload]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !hw script { @hw_pin1 = 0; }\n");
  ShowMessages("\t\te.g : !hw unload\n");
}

VOID CommandHw(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() >= 2 &&
      CompareLowerCaseStrings(CommandTokens.at(1), "script")) {
    HwdbgScriptRunScript(
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)).c_str(),
        HWDBG_TEST_READ_INSTANCE_INFO_PATH, HWDBG_TEST_WRITE_SCRIPT_BUFFER_PATH,
        DEFAULT_INITIAL_BRAM_BUFFER_SIZE);
  } else if (CommandTokens.size() >= 2 &&
             (CompareLowerCaseStrings(CommandTokens.at(1), "eval") ||
              CompareLowerCaseStrings(CommandTokens.at(1), "evaluation"))) {
    ScriptEngineWrapperTestParserForHwdbg(
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)));
  } else if (CommandTokens.size() == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "unload")) {
    g_HwdbgInstanceInfoIsValid = FALSE;
  } else {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandHwHelp();
    return;
  }
}
