
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandIoapicHelp() {
  ShowMessages("!ioapic : shows the details of I/O APIC entries.\n\n");
  ShowMessages("syntax : \t!ioapic\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !ioapic\n");
}

BOOLEAN HyperDbgGetIoApic(IO_APIC_ENTRY_PACKETS *IoApic) {
  BOOLEAN IsUsingX2APIC;
  return CommandApicSendRequest(DEBUGGER_APIC_REQUEST_TYPE_READ_IO_APIC, IoApic,
                                sizeof(IO_APIC_ENTRY_PACKETS), &IsUsingX2APIC);
}

ULONG CommandIoapicShowRedir(PUCHAR Desc, BOOLEAN CommandReg, BOOLEAN DestSelf,
                             ULONGLONG Lh, ULONGLONG Ll) {
  static PUCHAR DelMode[] = {(PUCHAR) "FixedDel", (PUCHAR) "LowestDl",
                             (PUCHAR) "res010  ", (PUCHAR) "remoterd",
                             (PUCHAR) "NMI     ", (PUCHAR) "RESET   ",
                             (PUCHAR) "res110  ", (PUCHAR) "ExtINTA "};
  static PUCHAR DesShDesc[] = {(PUCHAR) "", (PUCHAR) "  Dest=Self",
                               (PUCHAR) "   Dest=ALL", (PUCHAR) " Dest=Othrs"};
  ULONG Del, Dest, Delstat, Rirr, Trig, Masked, Destsh;
  Del = (Ll >> 8) & 0x7;
  Dest = (Ll >> 11) & 0x1;
  Delstat = (Ll >> 12) & 0x1;
  Rirr = (Ll >> 14) & 0x1;
  Trig = (Ll >> 15) & 0x1;
  Masked = (Ll >> 16) & 0x1;
  Destsh = (Ll >> 18) & 0x3;
  if (CommandReg) {
    Masked = 0;
  }
  ShowMessages("%s: %s  Vec:%02X  %s  ", Desc, SeparateTo64BitValue(Ll).c_str(),
               Ll & 0xff, DelMode[Del]);
  if (DestSelf) {
    ShowMessages("%s", DesShDesc[1]);
  } else if (CommandReg && Destsh) {
    ShowMessages("%s", DesShDesc[Destsh]);
  } else {
    if (Dest) {
      ShowMessages("Lg:%08x", Lh);
    } else {
      ShowMessages("PhysDest:%02X", (Lh >> 56) & 0xFF);
    }
  }
  ShowMessages("%s  %s  %s  %s\n", Delstat ? "-Pend" : "     ",
               Trig ? "level" : "edge ", Rirr ? "rirr" : "    ",
               Masked ? "masked" : "      ");
  return 0;
}

VOID CommandIoapicShowIoApicEntries(IO_APIC_ENTRY_PACKETS *IoApicPackets) {
  UINT32 Index = 0;
  UINT32 Max;
  UCHAR Desc[40];
  UINT64 ll, lh;
  UINT64 ApicBasePa = IoApicPackets->ApicBasePa;
  ll = IoApicPackets->IoLl;
  Max = (ll >> 16) & 0xff;
  ShowMessages("IoApic @ %08x  ID:%x (%x)  Arb:%x\t I/O APIC VA: %s\n",
               ApicBasePa, IoApicPackets->IoIdReg >> 24, ll & 0xFF,
               IoApicPackets->IoArbIdReg,
               SeparateTo64BitValue(IoApicPackets->ApicBaseVa).c_str());
  Max *= 2;
  for (Index = 0; Index <= Max; Index += 2) {
    if (Index >= MAX_NUMBER_OF_IO_APIC_ENTRIES) {
      ShowMessages(
          "err, there are additional entries in the I/O APIC that are not "
          "fully displayed in the results");
      return;
    }
    ll = IoApicPackets->LlLhData[Index];
    lh = IoApicPackets->LlLhData[Index + 1];
    sprintf((CHAR *)Desc, "Inti%02X.", Index / 2);
    CommandIoapicShowRedir(Desc, FALSE, FALSE, lh, ll);
  }
}

VOID CommandIoapic(vector<CommandToken> CommandTokens, string Command) {
  IO_APIC_ENTRY_PACKETS *IoApicPackets = NULL;
  if (CommandTokens.size() != 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandIoapicHelp();
    return;
  }
  IoApicPackets =
      (IO_APIC_ENTRY_PACKETS *)malloc(sizeof(IO_APIC_ENTRY_PACKETS));
  if (IoApicPackets == NULL) {
    ShowMessages("err, allocating buffer for receiving I/O APIC");
  }
  RtlZeroMemory(IoApicPackets, sizeof(IO_APIC_ENTRY_PACKETS));
  if (HyperDbgGetIoApic(IoApicPackets) == TRUE) {
    CommandIoapicShowIoApicEntries(IoApicPackets);
  }
  free(IoApicPackets);
}
