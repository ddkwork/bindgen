
#include "pch.h"

_Use_decl_annotations_

    UINT64
    PhysicalAddressToVirtualAddress(UINT64 PhysicalAddress) {
  PHYSICAL_ADDRESS PhysicalAddr;
  PhysicalAddr.QuadPart = PhysicalAddress;
  return (UINT64)MmGetVirtualForPhysical(PhysicalAddr);
}

_Use_decl_annotations_

    UINT64
    PhysicalAddressToVirtualAddressByProcessId(PVOID PhysicalAddress,
                                               UINT32 ProcessId) {
  CR3_TYPE CurrentProcessCr3;
  UINT64 VirtualAddress;
  PHYSICAL_ADDRESS PhysicalAddr;
  CurrentProcessCr3 = SwitchToProcessMemoryLayout(ProcessId);
  if (CurrentProcessCr3.Flags == NULL64_ZERO) {
    return NULL64_ZERO;
  }
  PhysicalAddr.QuadPart = (LONGLONG)PhysicalAddress;
  VirtualAddress = (UINT64)MmGetVirtualForPhysical(PhysicalAddr);
  SwitchToPreviousProcess(CurrentProcessCr3);
  return VirtualAddress;
}

_Use_decl_annotations_

    UINT64
    PhysicalAddressToVirtualAddressByCr3(PVOID PhysicalAddress,
                                         CR3_TYPE TargetCr3) {
  CR3_TYPE CurrentProcessCr3;
  UINT64 VirtualAddress;
  PHYSICAL_ADDRESS PhysicalAddr;
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(TargetCr3);
  if (CurrentProcessCr3.Flags == NULL64_ZERO) {
    return NULL64_ZERO;
  }
  PhysicalAddr.QuadPart = (LONGLONG)PhysicalAddress;
  VirtualAddress = (UINT64)MmGetVirtualForPhysical(PhysicalAddr);
  SwitchToPreviousProcess(CurrentProcessCr3);
  return VirtualAddress;
}

_Use_decl_annotations_

    UINT64
    PhysicalAddressToVirtualAddressOnTargetProcess(PVOID PhysicalAddress) {
  CR3_TYPE GuestCr3;
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  return PhysicalAddressToVirtualAddressByCr3(PhysicalAddress, GuestCr3);
}

_Use_decl_annotations_

    UINT64
    VirtualAddressToPhysicalAddress(_In_ PVOID VirtualAddress) {
  return MmGetPhysicalAddress(VirtualAddress).QuadPart;
}

_Use_decl_annotations_

    UINT64
    VirtualAddressToPhysicalAddressByProcessId(PVOID VirtualAddress,
                                               UINT32 ProcessId) {
  CR3_TYPE CurrentProcessCr3;
  UINT64 PhysicalAddress;
  CurrentProcessCr3 = SwitchToProcessMemoryLayout(ProcessId);
  if (CurrentProcessCr3.Flags == NULL64_ZERO) {
    return NULL64_ZERO;
  }
  PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
  SwitchToPreviousProcess(CurrentProcessCr3);
  return PhysicalAddress;
}

_Use_decl_annotations_

    UINT64
    VirtualAddressToPhysicalAddressByProcessCr3(PVOID VirtualAddress,
                                                CR3_TYPE TargetCr3) {
  CR3_TYPE CurrentProcessCr3;
  UINT64 PhysicalAddress;
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(TargetCr3);
  if (CurrentProcessCr3.Flags == NULL64_ZERO) {
    return NULL64_ZERO;
  }
  PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
  SwitchToPreviousProcess(CurrentProcessCr3);
  return PhysicalAddress;
}

_Use_decl_annotations_

    UINT64
    VirtualAddressToPhysicalAddressOnTargetProcess(PVOID VirtualAddress) {
  CR3_TYPE CurrentCr3;
  CR3_TYPE GuestCr3;
  UINT64 PhysicalAddress;
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  CurrentCr3 = SwitchToProcessMemoryLayoutByCr3(GuestCr3);
  if (CurrentCr3.Flags == NULL64_ZERO) {
    return NULL64_ZERO;
  }
  PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
  SwitchToPreviousProcess(CurrentCr3);
  return PhysicalAddress;
}
