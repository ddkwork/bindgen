
#include "pch.h"

typedef struct _COMPATIBILITY_CHECKS_STATUS {
  BOOLEAN IsX2Apic;
  BOOLEAN RtmSupport;
  BOOLEAN PmlSupport;
  BOOLEAN ModeBasedExecutionSupport;
  BOOLEAN ExecuteOnlySupport;
  UINT32 VirtualAddressWidth;
  UINT32 PhysicalAddressWidth;
} COMPATIBILITY_CHECKS_STATUS, *PCOMPATIBILITY_CHECKS_STATUS;

VOID CompatibilityCheckPerformChecks();
