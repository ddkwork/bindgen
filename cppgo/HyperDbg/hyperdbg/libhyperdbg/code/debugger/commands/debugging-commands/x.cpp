
#include "pch.h"

VOID CommandXHelp() {
  ShowMessages(
      "x : searches and shows symbols (wildcard) and corresponding "
      "addresses.\n\n");
  ShowMessages("syntax : \tx [Module!Name (wildcard string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : x nt!ExAllocatePoolWithTag \n");
  ShowMessages("\t\te.g : x nt!ExAllocatePool* \n");
  ShowMessages("\t\te.g : x nt!* \n");
}

VOID CommandX(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() != 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandXHelp();
    return;
  }
  ScriptEngineSearchSymbolForMaskWrapper(
      GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
}
