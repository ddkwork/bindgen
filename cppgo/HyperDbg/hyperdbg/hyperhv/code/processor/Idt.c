
#include "pch.h"

VOID IdtDumpInterruptEntries() {
  KDESCRIPTOR64 descr;
  __sidt(&descr.Limit);
  ULONG n = (descr.Limit + 1) / sizeof(KIDTENTRY64);
  if (n > 0) {
    int i = 0;
    KIDTENTRY64 *pidte = (KIDTENTRY64 *)descr.Base;
    do {
      ULONG_PTR addr = ((ULONG_PTR)pidte->OffsetHigh << 32) +
                       ((ULONG_PTR)pidte->OffsetMiddle << 16) +
                       pidte->OffsetLow;
      LogInfo("Interrupt %u -> %p\n", i++, addr);
    } while (pidte++, --n);
  }
}
