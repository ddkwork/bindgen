
#pragma once
__declspec(dllexport) NTSTATUS DllInitialize(_In_ PUNICODE_STRING RegistryPath);
__declspec(dllexport) NTSTATUS DllUnload(void);
