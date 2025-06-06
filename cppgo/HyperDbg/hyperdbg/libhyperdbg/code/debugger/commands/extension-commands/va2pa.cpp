
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandVa2paHelp() {
  ShowMessages("!va2pa : converts virtual address to physical address.\n\n");
  ShowMessages(
      "syntax : \t!va2pa [VirtualAddress (hex)] [pid ProcessId (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !va2pa nt!ExAllocatePoolWithTag\n");
  ShowMessages("\t\te.g : !va2pa nt!ExAllocatePoolWithTag+5\n");
  ShowMessages("\t\te.g : !va2pa @rcx\n");
  ShowMessages("\t\te.g : !va2pa @rcx+5\n");
  ShowMessages("\t\te.g : !va2pa fffff801deadbeef\n");
  ShowMessages("\t\te.g : !va2pa fffff801deadbeef pid 0xc8\n");
}

VOID CommandVa2pa(vector<CommandToken> CommandTokens, string Command) {
  BOOL Status;
  ULONG ReturnedLength;
  UINT64 TargetVa;
  UINT32 Pid = 0;
  DEBUGGER_VA2PA_AND_PA2VA_COMMANDS AddressDetails = {0};
  if (CommandTokens.size() == 1 || CommandTokens.size() >= 5 ||
      CommandTokens.size() == 3) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandVa2paHelp();
    return;
  }
  if (g_ActiveProcessDebuggingState.IsActive) {
    Pid = g_ActiveProcessDebuggingState.ProcessId;
  }
  if (CommandTokens.size() == 2) {
    if (!SymbolConvertNameOrExprToAddress(
            GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)),
            &TargetVa)) {
      ShowMessages(
          "err, couldn't resolve error at '%s'\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
      return;
    }
  } else {
    if (CompareLowerCaseStrings(CommandTokens.at(1), "pid")) {
      if (!ConvertTokenToUInt32(CommandTokens.at(2), &Pid)) {
        ShowMessages("incorrect address, please enter a valid process id\n");
        return;
      }
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(CommandTokens.at(3)),
              &TargetVa)) {
        ShowMessages("err, couldn't resolve error at '%s'\n",
                     GetCaseSensitiveStringFromCommandToken(CommandTokens.at(3))
                         .c_str());
        return;
      }
    } else if (CompareLowerCaseStrings(CommandTokens.at(2), "pid")) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)),
              &TargetVa)) {
        ShowMessages("err, couldn't resolve error at '%s'\n",
                     GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1))
                         .c_str());
        return;
      }
      if (!ConvertTokenToUInt32(CommandTokens.at(3), &Pid)) {
        ShowMessages("incorrect address, please enter a valid process id\n");
        return;
      }
    } else {
      ShowMessages(
          "incorrect use of the '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      CommandVa2paHelp();
      return;
    }
  }
  AddressDetails.VirtualAddress = TargetVa;
  AddressDetails.ProcessId = Pid;
  AddressDetails.IsVirtual2Physical = TRUE;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    if (Pid != 0) {
      ShowMessages(ASSERT_MESSAGE_CANNOT_SPECIFY_PID);
      return;
    }
    KdSendVa2paAndPa2vaPacketToDebuggee(&AddressDetails);
  } else {
    AssertShowMessageReturnStmt(g_DeviceHandle,
                                ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    if (Pid == 0) {
      Pid = GetCurrentProcessId();
      AddressDetails.ProcessId = Pid;
    }
    Status = DeviceIoControl(
        g_DeviceHandle, IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS,
        &AddressDetails, SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS,
        &AddressDetails, SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS,
        &ReturnedLength, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      return;
    }
    if (AddressDetails.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
      ShowMessages("%llx\n", AddressDetails.PhysicalAddress);
    } else {
      ShowErrorMessage(AddressDetails.KernelStatus);
    }
  }
}
