
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandPcitreeHelp() {
  ShowMessages("!pcitree : enumerates all PCIe endpoints on the debuggee.\n\n");
  ShowMessages("syntax : \t!pcitree\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !pcitree\n");
}

VOID CommandPcitree(vector<CommandToken> CommandTokens, string Command) {
  BOOL Status;
  ULONG ReturnedLength;
  DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET PcitreePacket = {0};
  if (CommandTokens.size() != 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandPcitreeHelp();
    return;
  }
  if (g_IsSerialConnectedToRemoteDebuggee) {
    KdSendPcitreePacketToDebuggee(&PcitreePacket);
  } else {
    AssertShowMessageReturnStmt(g_DeviceHandle,
                                ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    Status = DeviceIoControl(
        g_DeviceHandle, IOCTL_PCIE_ENDPOINT_ENUM, &PcitreePacket,
        SIZEOF_DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET, &PcitreePacket,
        SIZEOF_DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET, &ReturnedLength, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      return;
    }
    if (PcitreePacket.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
      ShowMessages("%-12s | %-9s | %-17s | %s \n%s\n", "DBDF", "VID:DID",
                   "Vendor Name", "Device Name",
                   "-----------------------------------------------------------"
                   "-----------");
      for (UINT8 i = 0; i < (PcitreePacket.DeviceInfoListNum < DEV_MAX_NUM
                                 ? PcitreePacket.DeviceInfoListNum
                                 : DEV_MAX_NUM);
           i++) {
        Vendor *CurrentVendor =
            GetVendorById(PcitreePacket.DeviceInfoList[i].ConfigSpace.VendorId);
        CHAR *CurrentVendorName = (CHAR *)"N/A";
        CHAR *CurrentDeviceName = (CHAR *)"N/A";
        if (CurrentVendor != NULL) {
          CurrentVendorName = CurrentVendor->VendorName;
          Device *CurrentDevice = GetDeviceFromVendor(
              CurrentVendor,
              PcitreePacket.DeviceInfoList[i].ConfigSpace.DeviceId);
          if (CurrentDevice != NULL) {
            CurrentDeviceName = CurrentDevice->DeviceName;
          }
        }
        ShowMessages("%04x:%02x:%02x:%x | %04x:%04x | %-17.*s | %.*s\n", 0,
                     PcitreePacket.DeviceInfoList[i].Bus,
                     PcitreePacket.DeviceInfoList[i].Device,
                     PcitreePacket.DeviceInfoList[i].Function,
                     PcitreePacket.DeviceInfoList[i].ConfigSpace.VendorId,
                     PcitreePacket.DeviceInfoList[i].ConfigSpace.DeviceId,
                     strnlen_s(CurrentVendorName, PCI_NAME_STR_LENGTH),
                     CurrentVendorName,
                     strnlen_s(CurrentDeviceName, PCI_NAME_STR_LENGTH),
                     CurrentDeviceName);
        FreeVendor(CurrentVendor);
      }
      FreePciIdDatabase();
    } else {
      ShowErrorMessage(PcitreePacket.KernelStatus);
    }
  }
}
