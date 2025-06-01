
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandBdHelp() {
  ShowMessages("bd : disables a breakpoint using breakpoint id.\n\n");
  ShowMessages("syntax : \tbd [BreakpointId (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : bd 0\n");
  ShowMessages("\t\te.g : bd 2\n");
}

VOID CommandBd(vector<CommandToken> CommandTokens, string Command) {
  UINT64 BreakpointId;
  DEBUGGEE_BP_LIST_OR_MODIFY_PACKET Request = {0};
  if (CommandTokens.size() != 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandBdHelp();
    return;
  }
  if (!ConvertTokenToUInt64(CommandTokens.at(1), &BreakpointId)) {
    ShowMessages("please specify a correct hex value for breakpoint id\n\n");
    CommandBdHelp();
    return;
  }
  if (g_IsSerialConnectedToRemoteDebuggee) {
    Request.Request = DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE;
    Request.BreakpointId = BreakpointId;
    KdSendListOrModifyPacketToDebuggee(&Request);
  } else {
    ShowMessages(
        "err, disabling breakpoints is only valid if you connected to "
        "a debuggee in debugger-mode\n");
  }
}
