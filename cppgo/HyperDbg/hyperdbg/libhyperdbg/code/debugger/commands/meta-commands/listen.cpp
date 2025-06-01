
#include "pch.h"
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern BOOLEAN g_IsConnectedToRemoteDebugger;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandListenHelp() {
  ShowMessages(
      ".listen : listens for a client to connect to HyperDbg (works as "
      "a guest server).\n\n");
  ShowMessages(
      "note : \tif you don't specify port then HyperDbg uses the "
      "default port (%s)\n",
      DEFAULT_PORT);
  ShowMessages("syntax : \t.listen [Port (decimal)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .listen\n");
  ShowMessages("\t\te.g : .listen 50000\n");
}

VOID CommandListen(vector<CommandToken> CommandTokens, string Command) {
  string Port;
  if (CommandTokens.size() >= 3) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandListenHelp();
    return;
  }
  if (g_IsConnectedToHyperDbgLocally || g_IsConnectedToRemoteDebuggee ||
      g_IsConnectedToRemoteDebugger) {
    ShowMessages(
        "you're connected to a debugger, please use '.disconnect' "
        "command\n");
    return;
  }
  if (g_IsSerialConnectedToRemoteDebuggee ||
      g_IsSerialConnectedToRemoteDebugger) {
    ShowMessages(
        "you're connected to a an instance of HyperDbg, please use "
        "'.debug close' command\n");
    return;
  }
  if (CommandTokens.size() == 1) {
    ShowMessages("listening on %s ...\n", DEFAULT_PORT);
    RemoteConnectionListen(DEFAULT_PORT);
    return;
  } else if (CommandTokens.size() == 2) {
    Port = GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1));
    if (!IsNumber(Port) || stoi(Port) > 65535 || stoi(Port) < 0) {
      ShowMessages("incorrect port\n");
      return;
    }
    ShowMessages("listening on %s ...\n", Port.c_str());
    RemoteConnectionListen(Port.c_str());
  } else {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandListenHelp();
    return;
  }
}
