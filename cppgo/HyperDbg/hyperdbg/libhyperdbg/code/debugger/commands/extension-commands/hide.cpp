
#include "pch.h"
extern UINT64 g_CpuidAverage;
extern UINT64 g_CpuidStandardDeviation;
extern UINT64 g_CpuidMedian;
extern UINT64 g_RdtscAverage;
extern UINT64 g_RdtscStandardDeviation;
extern UINT64 g_RdtscMedian;
extern BOOLEAN g_TransparentResultsMeasured;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandHideHelp() {
  ShowMessages(
      "!hide : tries to make HyperDbg transparent from anti-debugging "
      "and anti-hypervisor methods (this is a work under progress and new "
      "methods will be added frequently).\n\n");
  ShowMessages("syntax : \t!hide\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !hide\n");
}

BOOLEAN HyperDbgEnableTransparentMode() {
  BOOLEAN Status;
  ULONG ReturnedLength;
  DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE HideRequest = {0};
  AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED,
                              AssertReturnFalse);
  HideRequest.IsHide = TRUE;
  Status = DeviceIoControl(
      g_DeviceHandle,
      IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER, &HideRequest,
      SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE, &HideRequest,
      SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE, &ReturnedLength,
      NULL);
  if (!Status) {
    ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
    return FALSE;
  }
  if (HideRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    ShowMessages("transparent debugging successfully enabled :)\n");
    return TRUE;
  } else {
    ShowErrorMessage(HideRequest.KernelStatus);
    return FALSE;
  }
}

VOID CommandHide(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() != 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandHideHelp();
    return;
  }
  HyperDbgEnableTransparentMode();
}
