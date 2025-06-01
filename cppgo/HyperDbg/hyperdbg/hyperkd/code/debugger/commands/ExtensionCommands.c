
#include "pch.h"

UINT32 ExtensionCommandPerformActionsForApicRequests(
    PDEBUGGER_APIC_REQUEST ApicRequest) {
  BOOLEAN IsUsingX2APIC = FALSE;
  PLAPIC_PAGE BufferToStoreLApic =
      (LAPIC_PAGE *)(((CHAR *)ApicRequest) + sizeof(DEBUGGER_APIC_REQUEST));
  PIO_APIC_ENTRY_PACKETS BufferToStoreIoApic =
      (IO_APIC_ENTRY_PACKETS *)(((CHAR *)ApicRequest) +
                                sizeof(DEBUGGER_APIC_REQUEST));
  if (ApicRequest->ApicType == DEBUGGER_APIC_REQUEST_TYPE_READ_LOCAL_APIC) {
    if (VmFuncApicStoreLocalApicFields(BufferToStoreLApic, &IsUsingX2APIC)) {
      ApicRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      ApicRequest->IsUsingX2APIC = IsUsingX2APIC;
      return sizeof(DEBUGGER_APIC_REQUEST) + sizeof(LAPIC_PAGE);
    } else {
      ApicRequest->KernelStatus = DEBUGGER_ERROR_APIC_ACTIONS_ERROR;
      return sizeof(DEBUGGER_APIC_REQUEST);
    }
  } else if (ApicRequest->ApicType == DEBUGGER_APIC_REQUEST_TYPE_READ_IO_APIC) {
    if (VmFuncApicStoreIoApicFields(BufferToStoreIoApic)) {
      ApicRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      return sizeof(DEBUGGER_APIC_REQUEST) + sizeof(IO_APIC_ENTRY_PACKETS);
    } else {
      ApicRequest->KernelStatus = DEBUGGER_ERROR_APIC_ACTIONS_ERROR;
      return sizeof(DEBUGGER_APIC_REQUEST);
    }
  } else {
    ApicRequest->KernelStatus = DEBUGGER_ERROR_APIC_ACTIONS_ERROR;
    return sizeof(DEBUGGER_APIC_REQUEST);
  }
}

VOID ExtensionCommandPerformQueryIdtEntriesRequest(
    PINTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS IdtQueryRequest,
    BOOLEAN ReadFromVmxRoot) {
  VmFuncIdtQueryEntries(IdtQueryRequest, ReadFromVmxRoot);
  IdtQueryRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
}

VOID ExtensionCommandVa2paAndPa2va(
    PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS AddressDetails,
    BOOLEAN OperateOnVmxRoot) {
  if (OperateOnVmxRoot) {
    if (AddressDetails->IsVirtual2Physical) {
      AddressDetails->PhysicalAddress =
          VirtualAddressToPhysicalAddressOnTargetProcess(
              (PVOID)AddressDetails->VirtualAddress);
      if (AddressDetails->PhysicalAddress == (UINT64)NULL) {
        AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
      } else {
        AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      }
    } else {
      AddressDetails->VirtualAddress =
          PhysicalAddressToVirtualAddressOnTargetProcess(
              (PVOID)AddressDetails->PhysicalAddress);
      AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
    }
  } else {
    if (AddressDetails->ProcessId ==
        HANDLE_TO_UINT32(PsGetCurrentProcessId())) {
      if (AddressDetails->IsVirtual2Physical) {
        AddressDetails->PhysicalAddress = VirtualAddressToPhysicalAddress(
            (PVOID)AddressDetails->VirtualAddress);
        if (AddressDetails->PhysicalAddress == (UINT64)NULL) {
          AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
        } else {
          AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
        }
      } else {
        AddressDetails->VirtualAddress =
            PhysicalAddressToVirtualAddress(AddressDetails->PhysicalAddress);
        AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      }
    } else {
      if (!CommonIsProcessExist(AddressDetails->ProcessId)) {
        AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_PROCESS_ID;
        return;
      }
      if (AddressDetails->IsVirtual2Physical) {
        AddressDetails->PhysicalAddress =
            VirtualAddressToPhysicalAddressByProcessId(
                (PVOID)AddressDetails->VirtualAddress,
                AddressDetails->ProcessId);
        if (AddressDetails->PhysicalAddress == (UINT64)NULL) {
          AddressDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
        } else {
          AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
        }
      } else {
        AddressDetails->VirtualAddress =
            PhysicalAddressToVirtualAddressByProcessId(
                (PVOID)AddressDetails->PhysicalAddress,
                AddressDetails->ProcessId);
        AddressDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      }
    }
  }
}

BOOLEAN ExtensionCommandPte(
    PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PteDetails,
    BOOLEAN IsOperatingInVmxRoot) {
  BOOLEAN Result = FALSE;
  CR3_TYPE RestoreCr3 = {0};
  if (IsOperatingInVmxRoot) {
    if (!VirtualAddressToPhysicalAddressOnTargetProcess(
            (PVOID)PteDetails->VirtualAddress)) {
      PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
      return FALSE;
    }
    RestoreCr3.Flags = SwitchToCurrentProcessMemoryLayout().Flags;
  } else {
    if (PteDetails->ProcessId != HANDLE_TO_UINT32(PsGetCurrentProcessId())) {
      if (!CommonIsProcessExist(PteDetails->ProcessId)) {
        PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_PROCESS_ID;
        return FALSE;
      }
      RestoreCr3.Flags =
          SwitchToProcessMemoryLayout(PteDetails->ProcessId).Flags;
    }
    if (!VirtualAddressToPhysicalAddress((PVOID)PteDetails->VirtualAddress)) {
      PteDetails->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
      Result = FALSE;
      goto RestoreTheState;
    }
  }
  PPAGE_ENTRY Pml4e = MemoryMapperGetPteVa((PVOID)PteDetails->VirtualAddress,
                                           PagingLevelPageMapLevel4);
  if (Pml4e) {
    PteDetails->Pml4eVirtualAddress = (UINT64)Pml4e;
    PteDetails->Pml4eValue = Pml4e->Flags;
  }
  PPAGE_ENTRY Pdpte = MemoryMapperGetPteVa(
      (PVOID)PteDetails->VirtualAddress, PagingLevelPageDirectoryPointerTable);
  if (Pdpte) {
    PteDetails->PdpteVirtualAddress = (UINT64)Pdpte;
    PteDetails->PdpteValue = Pdpte->Flags;
  }
  PPAGE_ENTRY Pde = MemoryMapperGetPteVa((PVOID)PteDetails->VirtualAddress,
                                         PagingLevelPageDirectory);
  if (Pde) {
    PteDetails->PdeVirtualAddress = (UINT64)Pde;
    PteDetails->PdeValue = Pde->Flags;
  }
  PPAGE_ENTRY Pte = MemoryMapperGetPteVa((PVOID)PteDetails->VirtualAddress,
                                         PagingLevelPageTable);
  if (Pte) {
    PteDetails->PteVirtualAddress = (UINT64)Pte;
    PteDetails->PteValue = Pte->Flags;
  }
  PteDetails->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  Result = TRUE;
RestoreTheState:
  if (RestoreCr3.Flags != (UINT64)NULL) {
    SwitchToPreviousProcess(RestoreCr3);
  }
  return Result;
}

VOID ExtensionCommandChangeAllMsrBitmapReadAllCores(UINT64 BitmapMask) {
  BroadcastChangeAllMsrBitmapReadAllCores(BitmapMask);
}

VOID ExtensionCommandResetChangeAllMsrBitmapReadAllCores() {
  BroadcastResetChangeAllMsrBitmapReadAllCores();
}

VOID ExtensionCommandChangeAllMsrBitmapWriteAllCores(UINT64 BitmapMask) {
  BroadcastChangeAllMsrBitmapWriteAllCores(BitmapMask);
}

VOID ExtensionCommandResetAllMsrBitmapWriteAllCores() {
  BroadcastResetAllMsrBitmapWriteAllCores();
}

VOID ExtensionCommandEnableRdtscExitingAllCores() {
  BroadcastEnableRdtscExitingAllCores();
}

VOID ExtensionCommandDisableRdtscExitingAllCores() {
  BroadcastDisableRdtscExitingAllCores();
}

VOID ExtensionCommandDisableRdtscExitingForClearingEventsAllCores() {
  BroadcastDisableRdtscExitingForClearingEventsAllCores();
}

VOID ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores(
    PDEBUGGER_EVENT Event) {
  BroadcastDisableMov2ControlRegsExitingForClearingEventsAllCores(
      &Event->Options);
}

VOID ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores() {
  BroadcastDisableMov2DebugRegsExitingForClearingEventsAllCores();
}

VOID ExtensionCommandEnableRdpmcExitingAllCores() {
  BroadcastEnableRdpmcExitingAllCores();
}

VOID ExtensionCommandDisableRdpmcExitingAllCores() {
  BroadcastDisableRdpmcExitingAllCores();
}

VOID ExtensionCommandSetExceptionBitmapAllCores(UINT64 ExceptionIndex) {
  BroadcastSetExceptionBitmapAllCores(ExceptionIndex);
}

VOID ExtensionCommandUnsetExceptionBitmapAllCores(UINT64 ExceptionIndex) {
  BroadcastUnsetExceptionBitmapAllCores(ExceptionIndex);
}

VOID ExtensionCommandResetExceptionBitmapAllCores() {
  BroadcastResetExceptionBitmapAllCores();
}

VOID ExtensionCommandEnableMovControlRegisterExitingAllCores(
    PDEBUGGER_EVENT Event) {
  BroadcastEnableMovControlRegisterExitingAllCores(&Event->Options);
}

VOID ExtensionCommandDisableMovToControlRegistersExitingAllCores(
    PDEBUGGER_EVENT Event) {
  BroadcastDisableMovToControlRegistersExitingAllCores(&Event->Options);
}

VOID ExtensionCommandEnableMovDebugRegistersExitingAllCores() {
  BroadcastEnableMovDebugRegistersExitingAllCores();
}

VOID ExtensionCommandDisableMovDebugRegistersExitingAllCores() {
  BroadcastDisableMovDebugRegistersExitingAllCores();
}

VOID ExtensionCommandSetExternalInterruptExitingAllCores() {
  BroadcastSetExternalInterruptExitingAllCores();
}

VOID ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() {
  BroadcastUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();
}

VOID ExtensionCommandIoBitmapChangeAllCores(UINT64 Port) {
  BroadcastIoBitmapChangeAllCores(Port);
}

VOID ExtensionCommandIoBitmapResetAllCores() {
  BroadcastIoBitmapResetAllCores();
}

VOID ExtensionCommandPcitree(
    PDEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET PcitreePacket,
    BOOLEAN OperateOnVmxRoot) {
  DWORD DeviceIdVendorId = 0xFFFFFFFF;
  DWORD ClassCode = 0xFFFFFFFF;
  UINT8 DevNum = 0;
  UNREFERENCED_PARAMETER(OperateOnVmxRoot);
  for (UINT8 b = 0; b < BUS_MAX_NUM; b++) {
    for (UINT8 d = 0; d < DEVICE_MAX_NUM; d++) {
      for (UINT8 f = 0; f < FUNCTION_MAX_NUM; f++) {
        DeviceIdVendorId = (DWORD)PciReadCam(b, d, f, 0, sizeof(DWORD));
        if (DeviceIdVendorId != 0xFFFFFFFF) {
          PcitreePacket->DeviceInfoList[DevNum].Bus = b;
          PcitreePacket->DeviceInfoList[DevNum].Device = d;
          PcitreePacket->DeviceInfoList[DevNum].Function = f;
          PcitreePacket->DeviceInfoList[DevNum].ConfigSpace.VendorId =
              (UINT16)(DeviceIdVendorId & 0xFFFF);
          PcitreePacket->DeviceInfoList[DevNum].ConfigSpace.DeviceId =
              (UINT16)(DeviceIdVendorId >> 16);
          ClassCode = (DWORD)PciReadCam(b, d, f, 0, sizeof(DWORD));
          PcitreePacket->DeviceInfoList[DevNum].ConfigSpace.ClassCode[0] =
              (UINT8)((ClassCode >> 24) & 0xFF);
          PcitreePacket->DeviceInfoList[DevNum].ConfigSpace.ClassCode[1] =
              (UINT8)((ClassCode >> 16) & 0xFF);
          PcitreePacket->DeviceInfoList[DevNum].ConfigSpace.ClassCode[2] =
              (UINT8)((ClassCode >> 8) & 0xFF);
          DevNum++;
          if (DevNum == DEV_MAX_NUM) {
            LogError(
                "Reached maximum number of devices (%u) that can be stored in "
                "debuggee response packet.\n",
                DEV_MAX_NUM);
            break;
          }
        }
      }
    }
  }
  PcitreePacket->DeviceInfoListNum = DevNum;
  if (PcitreePacket->DeviceInfoListNum) {
    PcitreePacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  } else {
    PcitreePacket->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
  }
}

VOID ExtensionCommandPcidevinfo(
    PDEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET PcidevinfoPacket,
    BOOLEAN OperateOnVmxRoot) {
  DWORD DeviceIdVendorId = 0xFFFFFFFF;
  UNREFERENCED_PARAMETER(OperateOnVmxRoot);
  DeviceIdVendorId = (DWORD)PciReadCam(
      PcidevinfoPacket->DeviceInfo.Bus, PcidevinfoPacket->DeviceInfo.Device,
      PcidevinfoPacket->DeviceInfo.Function, 0, 4);
  if (DeviceIdVendorId != 0xFFFFFFFF) {
    DWORD *cs = (DWORD *)&PcidevinfoPacket->DeviceInfo.ConfigSpace;
    for (UINT16 i = 0; i < CAM_CONFIG_SPACE_LENGTH; i += 4) {
      *cs = (DWORD)PciReadCam(
          PcidevinfoPacket->DeviceInfo.Bus, PcidevinfoPacket->DeviceInfo.Device,
          PcidevinfoPacket->DeviceInfo.Function, (BYTE)i, 4);
      cs++;
    }
    if ((PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.HeaderType &
         0x01) << 7 ==
            0 &&
        !PcidevinfoPacket->PrintRaw) {
      for (UINT8 i = 0; i < 5; i++) {
        if ((PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader.ConfigSpaceEp
                 .Bar[i] &
             0x1) == 0) {
          if (((PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                    .ConfigSpaceEp.Bar[i] &
                0x6) >>
               1) == 2) {
            UINT64 BarMsb = PcidevinfoPacket->DeviceInfo.ConfigSpace
                                .DeviceHeader.ConfigSpaceEp.Bar[i + 1];
            UINT64 BarLsb = PcidevinfoPacket->DeviceInfo.ConfigSpace
                                .DeviceHeader.ConfigSpaceEp.Bar[i];
            UINT64 Bar64 =
                ((BarMsb & 0xFFFFFFFF) << 32) + (BarLsb & 0xFFFFFFF0);
            PcidevinfoPacket->DeviceInfo.MmioBarInfo[i].Is64Bit = TRUE;
            if (Bar64 == 0) {
              PcidevinfoPacket->DeviceInfo.MmioBarInfo[i].IsEnabled = FALSE;
              continue;
            }
            PcidevinfoPacket->DeviceInfo.MmioBarInfo[i].Is64Bit = TRUE;
            PcidevinfoPacket->DeviceInfo.MmioBarInfo[i].IsEnabled = TRUE;
            i++;
          } else {
            UINT32 Bar32 = (PcidevinfoPacket->DeviceInfo.ConfigSpace
                                .DeviceHeader.ConfigSpaceEp.Bar[i] &
                            0xFFFFFFF0);
            PcidevinfoPacket->DeviceInfo.MmioBarInfo[i].Is64Bit = FALSE;
            if (Bar32 == 0) {
              PcidevinfoPacket->DeviceInfo.MmioBarInfo[i].IsEnabled = FALSE;
              continue;
            }
            PcidevinfoPacket->DeviceInfo.MmioBarInfo[i].Is64Bit = FALSE;
            PcidevinfoPacket->DeviceInfo.MmioBarInfo[i].IsEnabled = TRUE;
          }
        }
      }
    }
    PcidevinfoPacket->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  } else {
    PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.DeviceId = 0xFFFF;
    PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.VendorId = 0xFFFF;
    PcidevinfoPacket->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
  }
}
