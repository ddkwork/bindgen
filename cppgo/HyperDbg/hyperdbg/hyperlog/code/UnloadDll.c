
#include "pch.h"

NTSTATUS DllInitialize(_In_ PUNICODE_STRING RegistryPath) {
  UNREFERENCED_PARAMETER(RegistryPath);
  return STATUS_SUCCESS;
}

NTSTATUS DllUnload(void) { return STATUS_SUCCESS; }
