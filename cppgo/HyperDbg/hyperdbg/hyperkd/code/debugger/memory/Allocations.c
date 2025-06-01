
#include "pch.h"

BOOLEAN GlobalDebuggingStateAllocateZeroedMemory(VOID) {
  SSIZE_T BufferSizeInByte =
      sizeof(PROCESSOR_DEBUGGING_STATE) * KeQueryActiveProcessorCount(0);
  g_DbgState = PlatformMemAllocateZeroedNonPagedPool(BufferSizeInByte);
  if (!g_DbgState) {
    LogInfo("err, insufficient memory for allocating debugging state\n");
    return FALSE;
  }
  return TRUE;
}

VOID GlobalDebuggingStateFreeMemory(VOID) {
  PlatformMemFreePool(g_DbgState);
  g_DbgState = NULL;
}

BOOLEAN GlobalEventsAllocateZeroedMemory(VOID) {
  if (!g_Events) {
    g_Events = PlatformMemAllocateNonPagedPool(sizeof(DEBUGGER_CORE_EVENTS));
  }
  if (g_Events) {
    RtlZeroBytes(g_Events, sizeof(DEBUGGER_CORE_EVENTS));
  }
  return g_Events != NULL;
}

VOID GlobalEventsFreeMemory(VOID) {
  if (g_Events != NULL) {
    PlatformMemFreePool(g_Events);
    g_Events = NULL;
  }
}
