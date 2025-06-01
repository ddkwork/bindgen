
#include "pch.h"

BOOLEAN DirtyLoggingInitialize() {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (!g_CompatibilityCheck.PmlSupport) {
    LogWarning(
        "err, dirty logging mechanism is not initialized as the processor "
        "doesn't support PML");
    return FALSE;
  }
  for (size_t i = 0; i < ProcessorsCount; i++) {
    if (g_GuestState[i].PmlBufferAddress == NULL) {
      g_GuestState[i].PmlBufferAddress =
          PlatformMemAllocateNonPagedPool(PAGE_SIZE);
    }
    if (g_GuestState[i].PmlBufferAddress == NULL) {
      for (size_t j = 0; j < ProcessorsCount; j++) {
        if (g_GuestState[j].PmlBufferAddress != NULL) {
          PlatformMemFreePool(g_GuestState[j].PmlBufferAddress);
        }
      }
      return FALSE;
    }
    RtlZeroBytes(g_GuestState[i].PmlBufferAddress, PAGE_SIZE);
  }
  BroadcastEnablePmlOnAllProcessors();
  return TRUE;
}

BOOLEAN DirtyLoggingEnable(VIRTUAL_MACHINE_STATE *VCpu) {
  if (VCpu->PmlBufferAddress == NULL) {
    return FALSE;
  }
  UINT64 PmlPhysAddr = VirtualAddressToPhysicalAddress(VCpu->PmlBufferAddress);
  __vmx_vmwrite(VMCS_CTRL_PML_ADDRESS, PmlPhysAddr);
  __vmx_vmwrite(VMCS_GUEST_PML_INDEX, PML_ENTITY_NUM - 1);
  HvSetPmlEnableFlag(TRUE);
  return TRUE;
}

VOID DirtyLoggingDisable(VIRTUAL_MACHINE_STATE *VCpu) {
  UNREFERENCED_PARAMETER(VCpu);
  __vmx_vmwrite(VMCS_CTRL_PML_ADDRESS, NULL64_ZERO);
  __vmx_vmwrite(VMCS_GUEST_PML_INDEX, 0x0);
  HvSetPmlEnableFlag(FALSE);
}

VOID DirtyLoggingUninitialize() {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  BroadcastDisablePmlOnAllProcessors();
  for (size_t i = 0; i < ProcessorsCount; i++) {
    if (g_GuestState[i].PmlBufferAddress != NULL) {
      PlatformMemFreePool(g_GuestState[i].PmlBufferAddress);
    }
  }
}

VOID DirtyLoggingHandlePageModificationLog(VIRTUAL_MACHINE_STATE *VCpu) {
  for (size_t i = 0; i < PML_ENTITY_NUM; i++) {
    LogInfo("Address : %llx", VCpu->PmlBufferAddress[i]);
  }
}

BOOLEAN DirtyLoggingFlushPmlBuffer(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT64 *PmlBuf;
  UINT16 PmlIdx;
  BOOLEAN IsLargePage;
  PVOID PmlEntry;
  VmxVmread16P(VMCS_GUEST_PML_INDEX, &PmlIdx);
  if (PmlIdx == (PML_ENTITY_NUM - 1)) return FALSE;
  if (PmlIdx >= PML_ENTITY_NUM) {
    PmlIdx = 0;
  } else {
    PmlIdx++;
  }
  PmlBuf = VCpu->PmlBufferAddress;
  for (; PmlIdx < PML_ENTITY_NUM; PmlIdx++) {
    UINT64 AccessedPhysAddr;
    AccessedPhysAddr = PmlBuf[PmlIdx];
    PmlEntry = EptGetPml1OrPml2Entry(VCpu->EptPageTable, AccessedPhysAddr,
                                     &IsLargePage);
    if (PmlEntry == NULL) {
      LogInfo("Err, null page");
      continue;
    }
    if (IsLargePage) {
      ((PEPT_PML2_ENTRY)PmlEntry)->Dirty = FALSE;
    } else {
      ((PEPT_PML1_ENTRY)PmlEntry)->Dirty = FALSE;
    }
  }
  __vmx_vmwrite(VMCS_GUEST_PML_INDEX, PML_ENTITY_NUM - 1);
  return TRUE;
}

VOID DirtyLoggingHandleVmexits(VIRTUAL_MACHINE_STATE *VCpu) {
  LogInfo("Dirty Logging VM-exit");
  DirtyLoggingFlushPmlBuffer(VCpu);
  HvSuppressRipIncrement(VCpu);
}
