
#include "pch.h"

static unsigned MaxWait = 65536;

BOOLEAN SpinlockTryLock(volatile LONG *Lock) {
  return (!(*Lock) && !_interlockedbittestandset(Lock, 0));
}

void SpinlockLock(volatile LONG *Lock) {
  unsigned wait = 1;
  while (!SpinlockTryLock(Lock)) {
    for (unsigned i = 0; i < wait; ++i) {
      _mm_pause();
    }
    if (wait * 2 > MaxWait) {
      wait = MaxWait;
    } else {
      wait = wait * 2;
    }
  }
}

void SpinlockLockWithCustomWait(volatile LONG *Lock, unsigned MaximumWait) {
  unsigned wait = 1;
  while (!SpinlockTryLock(Lock)) {
    for (unsigned i = 0; i < wait; ++i) {
      _mm_pause();
    }
    if (wait * 2 > MaximumWait) {
      wait = MaximumWait;
    } else {
      wait = wait * 2;
    }
  }
}

void SpinlockUnlock(volatile LONG *Lock) { *Lock = 0; }
