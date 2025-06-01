
#include "pch.h"
extern UINT32 g_ProcessIdOfLatestStartingProcess;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandKillHelp() {
  ShowMessages(".kill : terminates the current running process.\n\n");
  ShowMessages("syntax : \t.kill \n");
}

VOID CommandKill(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() != 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandKillHelp();
    return;
  }
  if (g_ActiveProcessDebuggingState.IsActive) {
    if (!UdKillProcess(g_ActiveProcessDebuggingState.ProcessId)) {
      ShowMessages("process does not exists, is it already terminated?\n");
    }
  } else if (g_ProcessIdOfLatestStartingProcess != NULL) {
    if (!UdKillProcess(g_ProcessIdOfLatestStartingProcess)) {
      ShowMessages("process does not exists, is it already terminated?\n");
    }
    g_ProcessIdOfLatestStartingProcess = NULL;
  } else {
    ShowMessages("nothing to terminate!\n");
    return;
  }
}
