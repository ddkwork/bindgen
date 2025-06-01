
#include "pch.h"

BOOLEAN InsertionSortInsertItem(UINT64 ArrayPtr[], UINT32 *NumberOfItems,
                                UINT32 MaxNumOfItems, UINT32 *Index,
                                UINT64 Key) {
  UINT32 Idx;
  if (*NumberOfItems >= MaxNumOfItems) {
    return FALSE;
  }
  Idx = *NumberOfItems;
  while (Idx > 0 && ArrayPtr[Idx - 1] > Key) {
    ArrayPtr[Idx] = ArrayPtr[Idx - 1];
    Idx = Idx - 1;
  }
  ArrayPtr[Idx] = Key;
  *Index = Idx;
  (*NumberOfItems)++;
  return TRUE;
}

BOOLEAN InsertionSortDeleteItem(UINT64 ArrayPtr[], UINT32 *NumberOfItems,
                                UINT32 Index) {
  if (Index >= *NumberOfItems) {
    return FALSE;
  }
  for (size_t i = Index + 1; i < *NumberOfItems; i++) {
    ArrayPtr[i - 1] = ArrayPtr[i];
  }
  (*NumberOfItems)--;
  return TRUE;
}
