
#include "pch.h"
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandDetachHelp() {
  ShowMessages(".detach : detaches from debugging a user-mode process.\n\n");
  ShowMessages("syntax : \t.detach \n");
}

VOID DetachFromProcess() {
  DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS DetachRequest = {0};
  AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED,
                              AssertReturn);
  if (!g_ActiveProcessDebuggingState.IsActive) {
    ShowMessages("you're not attached to any thread\n");
    return;
  }
  UdDetachProcess(g_ActiveProcessDebuggingState.ProcessId,
                  g_ActiveProcessDebuggingState.ProcessDebuggingToken);
}

VOID CommandDetach(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() >= 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandDetachHelp();
    return;
  }
  if (g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, '.attach', and '.detach' commands are only usable "
        "in VMI Mode, you can use the '.process', or the '.thread' "
        "in Debugger Mode\n");
    return;
  }
  DetachFromProcess();
}
