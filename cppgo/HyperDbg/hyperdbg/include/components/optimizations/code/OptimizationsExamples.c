
#include "pch.h"

#define MAX_NUM_OF_ARRAY 10

VOID OptimizationExampleInsertionSortAndBinarySearch() {
  UINT64 Arr[MAX_NUM_OF_ARRAY] = {0};
  UINT32 NumberOfItems = 0;
  UINT32 Index;
  InsertionSortInsertItem(Arr, &NumberOfItems, MAX_NUM_OF_ARRAY, &Index, 12);
  InsertionSortInsertItem(Arr, &NumberOfItems, MAX_NUM_OF_ARRAY, &Index, 11);
  InsertionSortInsertItem(Arr, &NumberOfItems, MAX_NUM_OF_ARRAY, &Index, 13);
  InsertionSortInsertItem(Arr, &NumberOfItems, MAX_NUM_OF_ARRAY, &Index, 5);
  InsertionSortInsertItem(Arr, &NumberOfItems, MAX_NUM_OF_ARRAY, &Index, 6);
  InsertionSortInsertItem(Arr, &NumberOfItems, MAX_NUM_OF_ARRAY, &Index, 8);
  BOOLEAN Result =
      BinarySearchPerformSearchItem(Arr, NumberOfItems, &Index, 15);
  if (Result) {
    LogInfo("Index found: %d", Index);
    InsertionSortDeleteItem(Arr, &NumberOfItems, Index);
  } else {
    LogInfo("Index not found!");
  }
  BinarySearchPrintArray(Arr, NumberOfItems);
}
