
#include "pch.h"

VOID CommandGgHelp() {
  ShowMessages("gg : shows and changes the operating processor.\n\n");
  ShowMessages("syntax : \tgg\n");
  ShowMessages("syntax : \tgg [wp]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : gg \n");
  ShowMessages("\t\te.g : gg wp \n");
}

VOID CommandGg(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() == 1) {
    ShowMessages("Good Game! :)\n");
  } else if (CommandTokens.size() == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "wp")) {
    ShowMessages("Good Game, Well Played! :)\n");
  } else {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandGgHelp();
    return;
  }
}
