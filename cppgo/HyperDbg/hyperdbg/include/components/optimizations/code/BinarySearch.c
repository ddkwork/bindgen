
#include "pch.h"

VOID BinarySearchPrintArray(UINT64 ArrayPtr[], UINT32 NumberOfItems) {
  UINT32 i;
  for (i = 0; i < NumberOfItems; i++) {
    Log("%llx ", ArrayPtr[i]);
  }
  Log("\n");
}

BOOLEAN BinarySearchPerformSearchItem(UINT64 ArrayPtr[], UINT32 NumberOfItems,
                                      UINT32 *ResultIndex, UINT64 Key) {
  UINT32 Position = 0;
  UINT32 Limit = NumberOfItems;
  while (Position < Limit) {
    UINT32 TestPos = Position + ((Limit - Position) >> 1);
    if (ArrayPtr[TestPos] < Key)
      Position = TestPos + 1;
    else
      Limit = TestPos;
  }
  if (Position < NumberOfItems && ArrayPtr[Position] == Key) {
    *ResultIndex = Position;
    return TRUE;
  } else {
    return FALSE;
  }
}
