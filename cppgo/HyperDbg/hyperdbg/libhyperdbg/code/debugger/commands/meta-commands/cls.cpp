
#include "pch.h"

VOID CommandClsHelp() {
  ShowMessages(".cls : clears the screen.\n\n");
  ShowMessages("syntax : \t.cls\n");
}

VOID CommandCls(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() != 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandClsHelp();
    return;
  }
  system("cls");
}
