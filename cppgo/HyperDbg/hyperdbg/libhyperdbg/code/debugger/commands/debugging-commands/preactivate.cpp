
#include "pch.h"

VOID CommandPreactivateHelp() {
  ShowMessages("preactivate : preactivates a special functionality.\n\n");
  ShowMessages("syntax : \tpreactivate  [Type (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : preactivate mode\n");
  ShowMessages("\n");
  ShowMessages("type of activations:\n");
  ShowMessages("\tmode: used for preactivation of the '!mode' event\n");
}

VOID CommandPreactivate(vector<CommandToken> CommandTokens, string Command) {
  BOOL Status;
  ULONG ReturnedLength;
  DEBUGGER_PREACTIVATE_COMMAND PreactivateRequest = {0};
  if (CommandTokens.size() != 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandPreactivateHelp();
    return;
  }
  if (CompareLowerCaseStrings(CommandTokens.at(1), "mode") ||
      CompareLowerCaseStrings(CommandTokens.at(1), "!mode")) {
    PreactivateRequest.Type = DEBUGGER_PREACTIVATE_COMMAND_TYPE_MODE;
  } else {
    ShowMessages(
        "err, couldn't resolve error at '%s'\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
    return;
  }
  AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED,
                              AssertReturn);
  Status = DeviceIoControl(
      g_DeviceHandle, IOCTL_PREACTIVATE_FUNCTIONALITY, &PreactivateRequest,
      SIZEOF_DEBUGGER_PREACTIVATE_COMMAND, &PreactivateRequest,
      SIZEOF_DEBUGGER_PREACTIVATE_COMMAND, &ReturnedLength, NULL);
  if (!Status) {
    ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
    return;
  }
  if (PreactivateRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    ShowMessages("the requested service is activated successfully!\n");
  } else {
    ShowErrorMessage(PreactivateRequest.KernelStatus);
  }
}
