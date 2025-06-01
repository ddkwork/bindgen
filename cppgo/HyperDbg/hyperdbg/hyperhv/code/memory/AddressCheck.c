
#include "pch.h"

BOOLEAN CheckAddressValidityUsingTsx(CHAR *Address) {
  UINT32 Status = 0;
  BOOLEAN Result = FALSE;
  CHAR TempContent;
  if ((Status = _xbegin()) == _XBEGIN_STARTED) {
    TempContent = *(CHAR *)Address;
    _xend();
    Result = TRUE;
  } else {
    Result = FALSE;
  }
  return Result;
}

BOOLEAN CheckAddressCanonicality(UINT64 VAddr, PBOOLEAN IsKernelAddress) {
  UINT64 Addr = (UINT64)VAddr;
  UINT64 MaxVirtualAddrLowHalf, MinVirtualAddressHighHalf;
  UINT32 AddrWidth = g_CompatibilityCheck.VirtualAddressWidth;
  MaxVirtualAddrLowHalf = ((UINT64)1ull << (AddrWidth - 1)) - 1;
  MinVirtualAddressHighHalf = ~MaxVirtualAddrLowHalf;
  if ((Addr > MaxVirtualAddrLowHalf) && (Addr < MinVirtualAddressHighHalf)) {
    *IsKernelAddress = FALSE;
    return FALSE;
  }
  if (MinVirtualAddressHighHalf < Addr) {
    *IsKernelAddress = TRUE;
  } else {
    *IsKernelAddress = FALSE;
  }
  return TRUE;
}

BOOLEAN CheckAddressPhysical(UINT64 PAddr) {
  UINT64 Addr = (UINT64)PAddr;
  UINT64 MaxPA;
  UINT32 AddrWidth = g_CompatibilityCheck.PhysicalAddressWidth;
  MaxPA = ((UINT64)1ull << (AddrWidth)) - 1;
  if (Addr > MaxPA) {
    return FALSE;
  }
  return TRUE;
}

BOOLEAN CheckAccessValidityAndSafety(UINT64 TargetAddress, UINT32 Size) {
  CR3_TYPE GuestCr3;
  UINT64 OriginalCr3;
  BOOLEAN IsKernelAddress;
  BOOLEAN Result = FALSE;
  if (!CheckAddressCanonicality(TargetAddress, &IsKernelAddress)) {
    Result = FALSE;
    goto Return;
  }
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  OriginalCr3 = __readcr3();
  __writecr3(GuestCr3.Flags);
  UINT64 AddressToCheck =
      (CHAR *)TargetAddress + Size - ((CHAR *)PAGE_ALIGN(TargetAddress));
  if (AddressToCheck > PAGE_SIZE) {
    UINT64 ReadSize = AddressToCheck;
    while (Size != 0) {
      ReadSize = (UINT64)PAGE_ALIGN(TargetAddress + PAGE_SIZE) - TargetAddress;
      if (ReadSize == PAGE_SIZE && Size < PAGE_SIZE) {
        ReadSize = Size;
      }
      if (!MemoryMapperCheckIfPageIsPresentByCr3((PVOID)TargetAddress,
                                                 GuestCr3)) {
        Result = FALSE;
        goto RestoreCr3;
      }

      Size = (UINT32)(Size - ReadSize);
      TargetAddress = TargetAddress + ReadSize;
    }
  } else {
    if (!MemoryMapperCheckIfPageIsPresentByCr3((PVOID)TargetAddress,
                                               GuestCr3)) {
      Result = FALSE;
      goto RestoreCr3;
    }
  }
  Result = TRUE;
RestoreCr3:
  __writecr3(OriginalCr3);
Return:
  return Result;
}

UINT32 CheckAddressMaximumInstructionLength(PVOID Address) {
  UINT64 SizeOfSafeBufferToRead = 0;
  SizeOfSafeBufferToRead = (UINT64)Address & 0xfff;
  SizeOfSafeBufferToRead += MAXIMUM_INSTR_SIZE;
  if (SizeOfSafeBufferToRead >= PAGE_SIZE) {
    SizeOfSafeBufferToRead = SizeOfSafeBufferToRead - PAGE_SIZE;
    if (CheckAccessValidityAndSafety((UINT64)Address + PAGE_SIZE,
                                     sizeof(CHAR))) {
      SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE;
    } else {
      SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE - SizeOfSafeBufferToRead;
    }
  } else {
    SizeOfSafeBufferToRead = MAXIMUM_INSTR_SIZE;
  }
  return (UINT32)SizeOfSafeBufferToRead;
}
