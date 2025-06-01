
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_AddressConversion;

VOID CommandIdtHelp() {
  ShowMessages(
      "!ioapic : shows entries of Interrupt Descriptor Table (IDT).\n\n");
  ShowMessages("syntax : \t!idt [IdtEntry (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !idt\n");
  ShowMessages("\t\te.g : !idt 1d\n");
}

BOOLEAN HyperDbgGetIdtEntry(
    INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS *IdtPacket) {
  BOOL Status;
  ULONG ReturnedLength;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    if (!KdSendQueryIdtPacketsToDebuggee(IdtPacket)) {
      return FALSE;
    } else {
      return TRUE;
    }
  } else {
    AssertShowMessageReturnStmt(
        g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status = DeviceIoControl(g_DeviceHandle, IOCTL_QUERY_IDT_ENTRY, IdtPacket,
                             SIZEOF_INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS,
                             IdtPacket,
                             SIZEOF_INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS,
                             &ReturnedLength, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      return FALSE;
    }
    if (IdtPacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
      return TRUE;
    } else {
      ShowErrorMessage(IdtPacket->KernelStatus);
      return FALSE;
    }
  }
}

VOID CommandIdt(vector<CommandToken> CommandTokens, string Command) {
  UINT32 IdtEntry;
  INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS *IdtPacket = NULL;
  BOOLEAN ShowAllEntries = TRUE;
  UINT64 UsedBaseAddress = NULL;
  if (CommandTokens.size() == 1) {
    ShowAllEntries = TRUE;
  } else if (CommandTokens.size() == 2) {
    ShowAllEntries = FALSE;
    if (ConvertTokenToUInt32(CommandTokens.at(1), &IdtEntry) == FALSE) {
      ShowMessages("err, invalid IDT entry number\n");
      return;
    }
    if (IdtEntry > MAX_NUMBER_OF_IDT_ENTRIES) {
      ShowMessages("err, invalid IDT entry number\n");
      return;
    }
  } else {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandIdtHelp();
    return;
  }
  IdtPacket = (INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS *)malloc(
      sizeof(INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS));
  if (IdtPacket == NULL) {
    ShowMessages("err, allocating buffer for receiving IDT entries");
  }
  RtlZeroMemory(IdtPacket, sizeof(INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS));
  if (HyperDbgGetIdtEntry(IdtPacket) == TRUE) {
    if (ShowAllEntries) {
      ShowMessages("IDT Entries:\n\n");
      for (UINT32 i = 0; i < MAX_NUMBER_OF_IDT_ENTRIES; i++) {
        ShowMessages("IDT[0x%x:%d]\t: %s\t", i, i,
                     SeparateTo64BitValue(IdtPacket->IdtEntry[i]).c_str());
        if (g_AddressConversion) {
          if (SymbolShowFunctionNameBasedOnAddress(IdtPacket->IdtEntry[i],
                                                   &UsedBaseAddress)) {
          }
        }
        ShowMessages("\n");
      }
    } else {
      ShowMessages("IDT[0x%x:%d] : %s\t", IdtEntry, IdtEntry,
                   SeparateTo64BitValue(IdtPacket->IdtEntry[IdtEntry]).c_str());
      if (g_AddressConversion) {
        if (SymbolShowFunctionNameBasedOnAddress(IdtPacket->IdtEntry[IdtEntry],
                                                 &UsedBaseAddress)) {
        }
        ShowMessages("\n");
      }
    }
  }
  free(IdtPacket);
}
