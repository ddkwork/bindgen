
#include "pch.h"
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern BOOLEAN g_IsConnectedToRemoteDebugger;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;
extern string g_ServerPort;
extern string g_ServerIp;

VOID CommandStatusHelp() {
  ShowMessages(
      ".status | status : gets the status of current debugger in local "
      "system (if you connected to a remote system then '.status' "
      "shows the state of current debugger, while 'status' shows the "
      "state of remote debuggee).\n\n");
  ShowMessages("syntax : \t.status\n");
  ShowMessages("syntax : \tstatus\n");
}

VOID CommandStatus(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() != 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandStatusHelp();
  }
  if (g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages("remote debugging - debugger ('debugger mode')\n");
  } else if (g_IsSerialConnectedToRemoteDebugger) {
    ShowMessages("remote debugging - debuggee ('debugger mode')\n");
  } else if (g_IsConnectedToRemoteDebuggee) {
    ShowMessages("remote debugging ('vmi mode'), ip : %s:%s \n",
                 g_ServerIp.c_str(), g_ServerPort.c_str());
  } else if (g_IsConnectedToHyperDbgLocally) {
    ShowMessages("local debugging ('vmi mode')\n");
  } else if (g_IsConnectedToRemoteDebugger) {
    ShowMessages(
        "a remote debugger connected to this system in ('vmi "
        "mode'), ip : %s:%s \n",
        g_ServerIp.c_str(), g_ServerPort.c_str());
  } else {
    ShowMessages("err, you're not connected to any instance of HyperDbg\n");
  }
}
