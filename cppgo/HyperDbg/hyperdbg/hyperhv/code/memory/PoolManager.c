
#include "pch.h"

BOOLEAN PlmgrAllocateRequestNewAllocation(SIZE_T NumberOfBytes) {
  g_RequestNewAllocation = PlatformMemAllocateZeroedNonPagedPool(NumberOfBytes);
  if (!g_RequestNewAllocation) {
    return FALSE;
  }
  return TRUE;
}

VOID PlmgrFreeRequestNewAllocation(VOID) {
  PlatformMemFreePool(g_RequestNewAllocation);
  g_RequestNewAllocation = NULL;
}

BOOLEAN PoolManagerInitialize() {
  SIZE_T BufferSize =
      MaximumRequestsQueueDepth * sizeof(REQUEST_NEW_ALLOCATION);
  if (!PlmgrAllocateRequestNewAllocation(BufferSize)) {
    LogError("Err, insufficient memory");
    return FALSE;
  }
  InitializeListHead(&g_ListOfAllocatedPoolsHead);
  g_IsNewRequestForDeAllocation = FALSE;
  return TRUE;
}

VOID PoolManagerUninitialize() {
  PLIST_ENTRY ListTemp = 0;
  ListTemp = &g_ListOfAllocatedPoolsHead;
  SpinlockLock(&LockForReadingPool);
  while (&g_ListOfAllocatedPoolsHead != ListTemp->Flink) {
    ListTemp = ListTemp->Flink;
    PPOOL_TABLE PoolTable =
        (PPOOL_TABLE)CONTAINING_RECORD(ListTemp, POOL_TABLE, PoolsList);
    if (!PoolTable->AlreadyFreed) {
      PlatformMemFreePool((PVOID)PoolTable->Address);
    }
    RemoveEntryList(&PoolTable->PoolsList);
    PlatformMemFreePool(PoolTable);
  }
  SpinlockUnlock(&LockForReadingPool);
  PlmgrFreeRequestNewAllocation();
}

BOOLEAN PoolManagerFreePool(UINT64 AddressToFree) {
  PLIST_ENTRY ListTemp = 0;
  BOOLEAN Result = FALSE;
  ListTemp = &g_ListOfAllocatedPoolsHead;
  SpinlockLock(&LockForReadingPool);
  while (&g_ListOfAllocatedPoolsHead != ListTemp->Flink) {
    ListTemp = ListTemp->Flink;
    PPOOL_TABLE PoolTable =
        (PPOOL_TABLE)CONTAINING_RECORD(ListTemp, POOL_TABLE, PoolsList);
    if (PoolTable->Address == AddressToFree) {
      PoolTable->ShouldBeFreed = TRUE;
      Result = TRUE;
      g_IsNewRequestForDeAllocation = TRUE;
      break;
    }
  }
  SpinlockUnlock(&LockForReadingPool);
  return Result;
}

VOID PoolManagerShowPreAllocatedPools() {
  PLIST_ENTRY ListTemp = 0;
  ListTemp = &g_ListOfAllocatedPoolsHead;
  while (&g_ListOfAllocatedPoolsHead != ListTemp->Flink) {
    ListTemp = ListTemp->Flink;
    PPOOL_TABLE PoolTable =
        (PPOOL_TABLE)CONTAINING_RECORD(ListTemp, POOL_TABLE, PoolsList);
    LogInfo(
        "Pool details, Pool intention: %x | Pool address: %llx | Pool state: "
        "%s | Should be freed: %s | Already freed: %s\n",
        PoolTable->Intention, PoolTable->Address,
        PoolTable->IsBusy ? "used" : "free",
        PoolTable->ShouldBeFreed ? "true" : "false",
        PoolTable->AlreadyFreed ? "true" : "false");
  }
}

UINT64 PoolManagerRequestPool(POOL_ALLOCATION_INTENTION Intention,
                              BOOLEAN RequestNewPool, UINT32 Size) {
  UINT64 Address = 0;
  ScopedSpinlock(
      LockForReadingPool, LIST_FOR_EACH_LINK(g_ListOfAllocatedPoolsHead,
                                             POOL_TABLE, PoolsList, PoolTable) {
        if (PoolTable->Intention == Intention && PoolTable->IsBusy == FALSE) {
          PoolTable->IsBusy = TRUE;
          Address = PoolTable->Address;
          break;
        }
      });
  if (RequestNewPool) {
    PoolManagerRequestAllocation(Size, 1, Intention);
  }
  return Address;
}

BOOLEAN PoolManagerAllocateAndAddToPoolTable(
    SIZE_T Size, UINT32 Count, POOL_ALLOCATION_INTENTION Intention) {
  for (size_t i = 0; i < Count; i++) {
    POOL_TABLE *SinglePool = NULL;
    SinglePool = PlatformMemAllocateZeroedNonPagedPool(sizeof(POOL_TABLE));
    if (!SinglePool) {
      LogError("Err, insufficient memory");
      return FALSE;
    }
    SinglePool->Address = (UINT64)PlatformMemAllocateZeroedNonPagedPool(Size);
    if (!SinglePool->Address) {
      PlatformMemFreePool(SinglePool);
      LogError("Err, insufficient memory");
      return FALSE;
    }
    SinglePool->Intention = Intention;
    SinglePool->IsBusy = FALSE;
    SinglePool->ShouldBeFreed = FALSE;
    SinglePool->AlreadyFreed = FALSE;
    SinglePool->Size = Size;
    InsertHeadList(&g_ListOfAllocatedPoolsHead, &(SinglePool->PoolsList));
  }
  return TRUE;
}

BOOLEAN PoolManagerCheckAndPerformAllocationAndDeallocation() {
  BOOLEAN Result = TRUE;
  PLIST_ENTRY ListTemp = 0;
  if (VmxGetCurrentExecutionMode() == TRUE) {
    return FALSE;
  }
  PAGED_CODE();
  if (g_IsNewRequestForAllocationReceived) {
    for (SIZE_T i = 0; i < MaximumRequestsQueueDepth; i++) {
      REQUEST_NEW_ALLOCATION *CurrentItem = &g_RequestNewAllocation[i];
      if (CurrentItem->Size != 0) {
        Result = PoolManagerAllocateAndAddToPoolTable(
            CurrentItem->Size, CurrentItem->Count, CurrentItem->Intention);
        CurrentItem->Count = 0;
        CurrentItem->Intention = 0;
        CurrentItem->Size = 0;
      }
    }
  }
  if (g_IsNewRequestForDeAllocation) {
    ListTemp = &g_ListOfAllocatedPoolsHead;
    SpinlockLock(&LockForReadingPool);
    while (&g_ListOfAllocatedPoolsHead != ListTemp->Flink) {
      ListTemp = ListTemp->Flink;
      PPOOL_TABLE PoolTable =
          (PPOOL_TABLE)CONTAINING_RECORD(ListTemp, POOL_TABLE, PoolsList);
      if (PoolTable->ShouldBeFreed && !PoolTable->AlreadyFreed) {
        PoolTable->AlreadyFreed = TRUE;
        PlatformMemFreePool((PVOID)PoolTable->Address);
        RemoveEntryList(&PoolTable->PoolsList);
        PlatformMemFreePool(PoolTable);
      }
    }
    SpinlockUnlock(&LockForReadingPool);
  }
  g_IsNewRequestForDeAllocation = FALSE;
  g_IsNewRequestForAllocationReceived = FALSE;
  return Result;
}

BOOLEAN PoolManagerRequestAllocation(SIZE_T Size, UINT32 Count,
                                     POOL_ALLOCATION_INTENTION Intention) {
  BOOLEAN FoundAPlace = FALSE;
  SpinlockLock(&LockForRequestAllocation);
  for (SIZE_T i = 0; i < MaximumRequestsQueueDepth; i++) {
    REQUEST_NEW_ALLOCATION *CurrentItem = &g_RequestNewAllocation[i];
    if (CurrentItem->Size == 0) {
      CurrentItem->Count = Count;
      CurrentItem->Intention = Intention;
      CurrentItem->Size = Size;
      FoundAPlace = TRUE;
      break;
    }
  }
  if (!FoundAPlace) {
    SpinlockUnlock(&LockForRequestAllocation);
    return FALSE;
  }
  g_IsNewRequestForAllocationReceived = TRUE;
  SpinlockUnlock(&LockForRequestAllocation);
  return TRUE;
}
