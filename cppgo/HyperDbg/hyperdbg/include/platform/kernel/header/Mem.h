
#pragma once
PVOID PlatformMemAllocateContiguousZeroedMemory(SIZE_T NumberOfBytes);
PVOID PlatformMemAllocateNonPagedPool(SIZE_T NumberOfBytes);
PVOID PlatformMemAllocateNonPagedPoolWithQuota(SIZE_T NumberOfBytes);
PVOID PlatformMemAllocateZeroedNonPagedPool(SIZE_T NumberOfBytes);
VOID PlatformMemFreePool(PVOID BufferAddress);
