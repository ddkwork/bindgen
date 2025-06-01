
#pragma once
BOOLEAN SpinlockTryLock(volatile LONG *Lock);
BOOLEAN SpinlockCheckLock(volatile LONG *Lock);
void SpinlockLock(volatile LONG *Lock);
void SpinlockLockWithCustomWait(volatile LONG *Lock, unsigned MaxWait);
void SpinlockUnlock(volatile LONG *Lock);
void SpinlockInterlockedCompareExchange(LONG volatile *Destination,
                                        LONG Exchange, LONG Comperand);
#define ScopedSpinlock(LockObject, CodeToRun)                            \
  MetaScopedExpr(SpinlockLock(&LockObject), SpinlockUnlock(&LockObject), \
                 CodeToRun)
