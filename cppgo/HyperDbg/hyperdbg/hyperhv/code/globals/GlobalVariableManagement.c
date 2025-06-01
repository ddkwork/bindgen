
#include "pch.h"

BOOLEAN GlobalGuestStateAllocateZeroedMemory(VOID) {
  SSIZE_T BufferSizeInByte =
      sizeof(VIRTUAL_MACHINE_STATE) * KeQueryActiveProcessorCount(0);
  if (!g_GuestState) {
    g_GuestState = PlatformMemAllocateNonPagedPool(BufferSizeInByte);
    if (!g_GuestState) {
      LogError("Err, insufficient memory\n");
      return FALSE;
    }
  }
  RtlZeroMemory(g_GuestState, BufferSizeInByte);
  return TRUE;
}

VOID GlobalGuestStateFreeMemory(VOID) {
  PlatformMemFreePool(g_GuestState);
  g_GuestState = NULL;
}
