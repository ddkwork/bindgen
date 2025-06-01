
#include "pch.h"
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern BOOLEAN g_IsConnectedToRemoteDebugger;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;
extern string g_ServerPort;
extern string g_ServerIp;

VOID CommandConnectHelp() {
  ShowMessages(
      ".connect : connects to a remote or local machine to start "
      "debugging.\n\n");
  ShowMessages("syntax : \t.connect [local]\n");
  ShowMessages("syntax : \t.connect [Ip (string)] [Port (decimal)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .connect local\n");
  ShowMessages("\t\te.g : .connect 192.168.1.5 50000\n");
}

VOID ConnectLocalDebugger() { g_IsConnectedToHyperDbgLocally = TRUE; }

BOOLEAN ConnectRemoteDebugger(const CHAR *Ip, const CHAR *Port) {
  if (!ValidateIP(Ip)) {
    return FALSE;
  }
  if (Port != NULL) {
    if (!IsNumber(Port) || stoi(Port) > 65535 || stoi(Port) < 0) {
      return FALSE;
    }
    g_ServerIp = Ip;
    g_ServerPort = Port;
    RemoteConnectionConnect(Ip, Port);
  } else {
    g_ServerIp = Ip;
    g_ServerPort = DEFAULT_PORT;
    RemoteConnectionConnect(Ip, DEFAULT_PORT);
  }
  return TRUE;
}

VOID CommandConnect(vector<CommandToken> CommandTokens, string Command) {
  string Ip;
  string Port;
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
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandConnectHelp();
    return;
  } else if (CompareLowerCaseStrings(CommandTokens.at(1), "local") &&
             CommandTokens.size() == 2) {
    ShowMessages("local debugging (vmi-mode)\n");
    ConnectLocalDebugger();
    return;
  } else if (CommandTokens.size() == 3 || CommandTokens.size() == 2) {
    Ip = GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1));
    if (CommandTokens.size() == 3) {
      Port = GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2));
    }
    if (!ValidateIP(Ip)) {
      ShowMessages("incorrect ip address\n");
      return;
    }
    if (CommandTokens.size() == 3) {
      if (!IsNumber(Port) || stoi(Port) > 65535 || stoi(Port) < 0) {
        ShowMessages("incorrect port\n");
        return;
      }
      ConnectRemoteDebugger(Ip.c_str(), Port.c_str());
    } else {
      ConnectRemoteDebugger(Ip.c_str(), NULL);
    }
  } else {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandConnectHelp();
    return;
  }
}
