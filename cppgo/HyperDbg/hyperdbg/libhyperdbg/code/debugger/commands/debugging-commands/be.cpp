
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandBeHelp() {
  ShowMessages("be : enables a breakpoint using breakpoint id.\n\n");
  ShowMessages("syntax : \tbe [BreakpointId (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : be 0\n");
  ShowMessages("\t\te.g : be 2\n");
}

VOID CommandBe(vector<CommandToken> CommandTokens, string Command) {
  UINT64 BreakpointId;
  DEBUGGEE_BP_LIST_OR_MODIFY_PACKET Request = {0};
  if (CommandTokens.size() != 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandBeHelp();
    return;
  }
  if (!ConvertTokenToUInt64(CommandTokens.at(1), &BreakpointId)) {
    ShowMessages("please specify a correct hex value for breakpoint id\n\n");
    CommandBeHelp();
    return;
  }
  if (g_IsSerialConnectedToRemoteDebuggee) {
    Request.Request = DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE;
    Request.BreakpointId = BreakpointId;
    KdSendListOrModifyPacketToDebuggee(&Request);
  } else {
    ShowMessages(
        "err, enabling breakpoints is only valid if you connected to "
        "a debuggee in debugger-mode\n");
  }
}
