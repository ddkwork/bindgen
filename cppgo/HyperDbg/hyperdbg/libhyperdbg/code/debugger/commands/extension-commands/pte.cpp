
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandPteHelp() {
  ShowMessages(
      "!pte : finds virtual addresses of different paging-levels.\n\n");
  ShowMessages(
      "syntax : \t!pte [VirtualAddress (hex)] [pid ProcessId (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !pte nt!ExAllocatePoolWithTag\n");
  ShowMessages("\t\te.g : !pte nt!ExAllocatePoolWithTag+5\n");
  ShowMessages("\t\te.g : !pte fffff801deadbeef\n");
  ShowMessages("\t\te.g : !pte 0x400021000 pid 1c0\n");
}

VOID CommandPteShowResults(UINT64 TargetVa,
                           PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PteRead) {
  ShowMessages("VA %llx\n", TargetVa);
  ShowMessages(
      "PML4E (PXE) at %016llx\tcontains %016llx\nPDPTE (PPE) at "
      "%016llx\tcontains "
      "%016llx\nPDE at %016llx\tcontains %016llx\n",
      PteRead->Pml4eVirtualAddress, PteRead->Pml4eValue,
      PteRead->PdpteVirtualAddress, PteRead->PdpteValue,
      PteRead->PdeVirtualAddress, PteRead->PdeValue);
  if (PteRead->PdeVirtualAddress == PteRead->PteVirtualAddress) {
    ShowMessages("PDE is a large page, so it doesn't have a PTE\n");
  } else {
    ShowMessages("PTE at %016llx\tcontains %016llx\n",
                 PteRead->PteVirtualAddress, PteRead->PteValue);
  }
}

VOID CommandPte(vector<CommandToken> CommandTokens, string Command) {
  BOOL Status;
  ULONG ReturnedLength;
  UINT64 TargetVa;
  UINT32 Pid = 0;
  DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS AddressDetails = {0};
  if (CommandTokens.size() == 1 || CommandTokens.size() >= 5 ||
      CommandTokens.size() == 3) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandPteHelp();
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
        ShowMessages("err, couldn't resolve error at '%s'\n\n",
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
      CommandPteHelp();
      return;
    }
  }
  AddressDetails.VirtualAddress = TargetVa;
  AddressDetails.ProcessId = Pid;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    if (Pid != 0) {
      ShowMessages(ASSERT_MESSAGE_CANNOT_SPECIFY_PID);
      return;
    }
    KdSendPtePacketToDebuggee(&AddressDetails);
  } else {
    AssertShowMessageReturnStmt(g_DeviceHandle,
                                ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    if (Pid == 0) {
      Pid = GetCurrentProcessId();
      AddressDetails.ProcessId = Pid;
    }
    Status = DeviceIoControl(
        g_DeviceHandle, IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS,
        &AddressDetails, SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS,
        &AddressDetails, SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS,
        &ReturnedLength, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      return;
    }
    if (AddressDetails.KernelStatus != DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
      ShowErrorMessage(AddressDetails.KernelStatus);
      return;
    }
    CommandPteShowResults(TargetVa, &AddressDetails);
  }
}
