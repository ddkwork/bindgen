
#include "pch.h"

VOID CommandUnhideHelp() {
  ShowMessages(
      "!unhide : reverts the transparency measures of the '!hide' command and "
      "exits the transparent mode.\n\n");
  ShowMessages("syntax : \t!unhide\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !unhide\n");
}

BOOLEAN HyperDbgDisableTransparentMode() {
  BOOLEAN Status;
  ULONG ReturnedLength;
  DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE UnhideRequest = {0};
  AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED,
                              AssertReturnFalse);
  UnhideRequest.IsHide = FALSE;
  Status = DeviceIoControl(
      g_DeviceHandle,
      IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER,
      &UnhideRequest, SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE,
      &UnhideRequest, SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE,
      &ReturnedLength, NULL);
  if (!Status) {
    ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
    return FALSE;
  }
  if (UnhideRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    ShowMessages("transparent debugging successfully disabled :)\n");
    return TRUE;
  } else {
    ShowErrorMessage(UnhideRequest.KernelStatus);
    return FALSE;
  }
}

VOID CommandUnhide(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() >= 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandUnhideHelp();
    return;
  }
  HyperDbgDisableTransparentMode();
}
