
#pragma once
BOOLEAN InsertionSortInsertItem(UINT64 ArrayPtr[], UINT32 *NumberOfItems,
                                UINT32 MaxNumOfItems, UINT32 *Index,
                                UINT64 Key);
BOOLEAN InsertionSortDeleteItem(UINT64 ArrayPtr[], UINT32 *NumberOfItems,
                                UINT32 Index);
