
#include "pch.h"

int TestBit(int BitNumber, unsigned long *addr) {
  return (BITMAP_ENTRY(BitNumber, addr) >> BITMAP_SHIFT(BitNumber)) & 1;
}

void ClearBit(int BitNumber, unsigned long *addr) {
  BITMAP_ENTRY(BitNumber, addr) &= ~(1UL << BITMAP_SHIFT(BitNumber));
}

void SetBit(int BitNumber, unsigned long *addr) {
  BITMAP_ENTRY(BitNumber, addr) |= (1UL << BITMAP_SHIFT(BitNumber));
}
