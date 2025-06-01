
#include "pch.h"

_Use_decl_annotations_

    BOOLEAN
    VmxAllocateVmxonRegion(VIRTUAL_MACHINE_STATE *VCpu) {
  IA32_VMX_BASIC_REGISTER VmxBasicMsr = {0};
  SIZE_T VmxonSize;
  UINT8 VmxonStatus;
  UINT8 *VmxonRegion;
  UINT64 VmxonRegionPhysicalAddr;
  UINT64 AlignedVmxonRegion;
  UINT64 AlignedVmxonRegionPhysicalAddr;
#ifdef ENV_WINDOWS
  if (KeGetCurrentIrql() > DISPATCH_LEVEL) KeRaiseIrqlToDpcLevel();
#endif
  VmxonSize = 2 * VMXON_SIZE;
  VmxonRegion = PlatformMemAllocateContiguousZeroedMemory(VmxonSize +
                                                          ALIGNMENT_PAGE_SIZE);
  if (VmxonRegion == NULL) {
    LogError("Err, couldn't allocate buffer for VMXON region");
    return FALSE;
  }
  VmxonRegionPhysicalAddr = VirtualAddressToPhysicalAddress(VmxonRegion);
  AlignedVmxonRegion =
      (UINT64)((ULONG_PTR)(VmxonRegion + ALIGNMENT_PAGE_SIZE - 1) &
               ~(ALIGNMENT_PAGE_SIZE - 1));
  LogDebugInfo("VMXON Region Address : %llx", AlignedVmxonRegion);
  AlignedVmxonRegionPhysicalAddr =
      (UINT64)((ULONG_PTR)(VmxonRegionPhysicalAddr + ALIGNMENT_PAGE_SIZE - 1) &
               ~(ALIGNMENT_PAGE_SIZE - 1));
  LogDebugInfo("VMXON Region Physical Address : %llx",
               AlignedVmxonRegionPhysicalAddr);
  VmxBasicMsr.AsUInt = __readmsr(IA32_VMX_BASIC);
  LogDebugInfo("Revision Identifier (IA32_VMX_BASIC - MSR 0x480) : 0x%x",
               VmxBasicMsr.VmcsRevisionId);
  *(UINT64 *)AlignedVmxonRegion = VmxBasicMsr.VmcsRevisionId;
  VmxonStatus = __vmx_on(&AlignedVmxonRegionPhysicalAddr);
  if (VmxonStatus) {
    LogError("Err, executing vmxon instruction failed with status : %d",
             VmxonStatus);
    return FALSE;
  }
  VCpu->VmxonRegionPhysicalAddress = AlignedVmxonRegionPhysicalAddr;
  VCpu->VmxonRegionVirtualAddress = (UINT64)VmxonRegion;
  return TRUE;
}

_Use_decl_annotations_

    BOOLEAN
    VmxAllocateVmcsRegion(VIRTUAL_MACHINE_STATE *VCpu) {
  IA32_VMX_BASIC_REGISTER VmxBasicMsr = {0};
  SIZE_T VmcsSize;
  UINT8 *VmcsRegion;
  UINT64 VmcsPhysicalAddr;
  UINT64 AlignedVmcsRegion;
  UINT64 AlignedVmcsRegionPhysicalAddr;
#ifdef ENV_WINDOWS
  if (KeGetCurrentIrql() > DISPATCH_LEVEL) KeRaiseIrqlToDpcLevel();
#endif
  VmcsSize = 2 * VMCS_SIZE;
  VmcsRegion =
      PlatformMemAllocateContiguousZeroedMemory(VmcsSize + ALIGNMENT_PAGE_SIZE);
  if (VmcsRegion == NULL) {
    LogError("Err, couldn't allocate Buffer for VMCS region");
    return FALSE;
  }
  VmcsPhysicalAddr = VirtualAddressToPhysicalAddress(VmcsRegion);
  AlignedVmcsRegion =
      (UINT64)((ULONG_PTR)(VmcsRegion + ALIGNMENT_PAGE_SIZE - 1) &
               ~(ALIGNMENT_PAGE_SIZE - 1));
  LogDebugInfo("VMCS region address : %llx", AlignedVmcsRegion);
  AlignedVmcsRegionPhysicalAddr =
      (UINT64)((ULONG_PTR)(VmcsPhysicalAddr + ALIGNMENT_PAGE_SIZE - 1) &
               ~(ALIGNMENT_PAGE_SIZE - 1));
  LogDebugInfo("VMCS region physical address : %llx",
               AlignedVmcsRegionPhysicalAddr);
  VmxBasicMsr.AsUInt = __readmsr(IA32_VMX_BASIC);
  LogDebugInfo("Revision Identifier (IA32_VMX_BASIC - MSR 0x480) : 0x%x",
               VmxBasicMsr.VmcsRevisionId);
  *(UINT64 *)AlignedVmcsRegion = VmxBasicMsr.VmcsRevisionId;
  VCpu->VmcsRegionPhysicalAddress = AlignedVmcsRegionPhysicalAddr;
  VCpu->VmcsRegionVirtualAddress = (UINT64)VmcsRegion;
  return TRUE;
}

BOOLEAN VmxAllocateVmmStack(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  VCpu->VmmStack =
      (UINT64)PlatformMemAllocateZeroedNonPagedPool(VMM_STACK_SIZE);
  if (VCpu->VmmStack == NULL64_ZERO) {
    LogError("Err, insufficient memory in allocating vmm stack");
    return FALSE;
  }
  LogDebugInfo("VMM Stack for logical processor : 0x%llx", VCpu->VmmStack);
  return TRUE;
}

BOOLEAN VmxAllocateMsrBitmap(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  VCpu->MsrBitmapVirtualAddress =
      (UINT64)PlatformMemAllocateZeroedNonPagedPool(PAGE_SIZE);
  if (VCpu->MsrBitmapVirtualAddress == NULL64_ZERO) {
    LogError("Err, insufficient memory in allocating MSR Bitmaps");
    return FALSE;
  }
  VCpu->MsrBitmapPhysicalAddress =
      VirtualAddressToPhysicalAddress((PVOID)VCpu->MsrBitmapVirtualAddress);
  LogDebugInfo("MSR Bitmap virtual address  : 0x%llx",
               VCpu->MsrBitmapVirtualAddress);
  LogDebugInfo("MSR Bitmap physical address : 0x%llx",
               VCpu->MsrBitmapPhysicalAddress);
  return TRUE;
}

BOOLEAN VmxAllocateIoBitmaps(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  VCpu->IoBitmapVirtualAddressA =
      (UINT64)PlatformMemAllocateZeroedNonPagedPool(PAGE_SIZE);
  if (VCpu->IoBitmapVirtualAddressA == NULL64_ZERO) {
    LogError("Err, insufficient memory in allocating I/O Bitmaps A");
    return FALSE;
  }
  VCpu->IoBitmapPhysicalAddressA =
      VirtualAddressToPhysicalAddress((PVOID)VCpu->IoBitmapVirtualAddressA);
  LogDebugInfo("I/O Bitmap A Virtual Address  : 0x%llx",
               VCpu->IoBitmapVirtualAddressA);
  LogDebugInfo("I/O Bitmap A Physical Address : 0x%llx",
               VCpu->IoBitmapPhysicalAddressA);
  VCpu->IoBitmapVirtualAddressB =
      (UINT64)PlatformMemAllocateZeroedNonPagedPool(PAGE_SIZE);
  if (VCpu->IoBitmapVirtualAddressB == NULL64_ZERO) {
    LogError("Err, insufficient memory in allocating I/O Bitmaps B");
    return FALSE;
  }
  VCpu->IoBitmapPhysicalAddressB =
      VirtualAddressToPhysicalAddress((PVOID)VCpu->IoBitmapVirtualAddressB);
  LogDebugInfo("I/O Bitmap B virtual address  : 0x%llx",
               VCpu->IoBitmapVirtualAddressB);
  LogDebugInfo("I/O Bitmap B physical address : 0x%llx",
               VCpu->IoBitmapPhysicalAddressB);
  return TRUE;
}

UINT64 *VmxAllocateInvalidMsrBimap() {
  UINT64 *InvalidMsrBitmap;
  InvalidMsrBitmap = PlatformMemAllocateZeroedNonPagedPool(0x1000 / 0x8);
  if (InvalidMsrBitmap == NULL) {
    return NULL;
  }
  for (UINT32 i = 0; i < 0x1000; ++i) {
    __try {
      __readmsr(i);
    } __except (EXCEPTION_EXECUTE_HANDLER) {
      SetBit(i, (unsigned long *)InvalidMsrBitmap);
    }
  }
  return InvalidMsrBitmap;
}

BOOLEAN VmxAllocateHostIdt(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 IdtSize =
      HOST_IDT_DESCRIPTOR_COUNT * sizeof(SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64);
  if (PAGE_SIZE > IdtSize) {
    IdtSize = PAGE_SIZE;
  }
  VCpu->HostIdt = (UINT64)PlatformMemAllocateZeroedNonPagedPool(IdtSize);
  if (VCpu->HostIdt == NULL64_ZERO) {
    LogError("Err, insufficient memory in allocating host IDT");
    return FALSE;
  }
  LogDebugInfo("Host IDT virtual address : 0x%llx", VCpu->HostIdt);
  return TRUE;
}

BOOLEAN VmxAllocateHostGdt(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 GdtSize = HOST_GDT_DESCRIPTOR_COUNT * sizeof(SEGMENT_DESCRIPTOR_64);
  if (PAGE_SIZE > GdtSize) {
    GdtSize = PAGE_SIZE;
  }
  VCpu->HostGdt = (UINT64)PlatformMemAllocateZeroedNonPagedPool(GdtSize);
  if (VCpu->HostGdt == NULL64_ZERO) {
    LogError("Err, insufficient memory in allocating host GDT");
    return FALSE;
  }
  LogDebugInfo("Host GDT virtual address : 0x%llx", VCpu->HostGdt);
  return TRUE;
}

BOOLEAN VmxAllocateHostTss(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 TssSize = PAGE_SIZE;
  VCpu->HostTss = (UINT64)PlatformMemAllocateZeroedNonPagedPool(TssSize);
  if (VCpu->HostTss == NULL64_ZERO) {
    LogError("Err, insufficient memory in allocating host TSS");
    return FALSE;
  }
  LogDebugInfo("Host TSS virtual address : 0x%llx", VCpu->HostTss);
  return TRUE;
}

BOOLEAN VmxAllocateHostInterruptStack(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  VCpu->HostInterruptStack =
      (UINT64)PlatformMemAllocateZeroedNonPagedPool(HOST_INTERRUPT_STACK_SIZE);
  if (VCpu->HostInterruptStack == NULL64_ZERO) {
    LogError("Err, insufficient memory in allocating host interrupt stack");
    return FALSE;
  }
  LogDebugInfo("Host interrupt stack virtual address : 0x%llx",
               VCpu->HostInterruptStack);
  return TRUE;
}
