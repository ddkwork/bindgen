
#pragma once

#define MaximumRequestsQueueDepth 300
#define NumberOfPreAllocatedBuffers 10

typedef struct _POOL_TABLE {
  UINT64 Address;
  SIZE_T Size;
  POOL_ALLOCATION_INTENTION Intention;
  LIST_ENTRY PoolsList;
  BOOLEAN IsBusy;
  BOOLEAN ShouldBeFreed;
  BOOLEAN AlreadyFreed;
} POOL_TABLE, *PPOOL_TABLE;

typedef struct _REQUEST_NEW_ALLOCATION {
  SIZE_T Size;
  UINT32 Count;
  POOL_ALLOCATION_INTENTION Intention;
} REQUEST_NEW_ALLOCATION, *PREQUEST_NEW_ALLOCATION;

REQUEST_NEW_ALLOCATION *g_RequestNewAllocation;

volatile LONG LockForRequestAllocation;

volatile LONG LockForReadingPool;

BOOLEAN g_IsNewRequestForAllocationReceived;

BOOLEAN g_IsNewRequestForDeAllocation;

LIST_ENTRY g_ListOfAllocatedPoolsHead;
static BOOLEAN PlmgrAllocateRequestNewAllocation(SIZE_T NumberOfBytes);
static VOID PlmgrFreeRequestNewAllocation(VOID);

BOOLEAN PoolManagerInitialize();

VOID PoolManagerUninitialize();
