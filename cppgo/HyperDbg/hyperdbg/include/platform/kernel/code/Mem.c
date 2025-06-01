
#include "pch.h"

PVOID PlatformMemAllocateContiguousZeroedMemory(SIZE_T NumberOfBytes) {
  PVOID Result = NULL;
  PHYSICAL_ADDRESS MaxPhysicalAddr = {.QuadPart = MAXULONG64};
  Result = MmAllocateContiguousMemory(NumberOfBytes, MaxPhysicalAddr);
  if (Result != NULL) RtlSecureZeroMemory(Result, NumberOfBytes);
  return Result;
}

PVOID PlatformMemAllocateNonPagedPool(SIZE_T NumberOfBytes) {
  PVOID Result = ExAllocatePoolWithTag(NonPagedPool, NumberOfBytes, POOLTAG);
  return Result;
}

PVOID PlatformMemAllocateNonPagedPoolWithQuota(SIZE_T NumberOfBytes) {
  PVOID Result = ExAllocatePool2(POOL_FLAG_NON_PAGED | POOL_FLAG_USE_QUOTA,
                                 NumberOfBytes, POOLTAG);
  return Result;
}

PVOID PlatformMemAllocateZeroedNonPagedPool(SIZE_T NumberOfBytes) {
  PVOID Result = ExAllocatePoolWithTag(NonPagedPool, NumberOfBytes, POOLTAG);
  if (Result != NULL) RtlSecureZeroMemory(Result, NumberOfBytes);
  return Result;
}

VOID PlatformMemFreePool(PVOID BufferAddress) {
  ExFreePoolWithTag(BufferAddress, POOLTAG);
}
