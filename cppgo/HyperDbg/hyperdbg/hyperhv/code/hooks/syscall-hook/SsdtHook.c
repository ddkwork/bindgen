
#include "pch.h"

PVOID SyscallHookGetKernelBase(PULONG ImageSize) {
  NTSTATUS Status;
  ZWQUERYSYSTEMINFORMATION ZwQSI = 0;
  UNICODE_STRING RoutineName;
  PVOID ModuleBase = NULL;
  PSYSTEM_MODULE_INFORMATION SystemInfoBuffer = NULL;
  ULONG SystemInfoBufferSize = 0;
  RtlInitUnicodeString(&RoutineName, L"ZwQuerySystemInformation");
  ZwQSI = (ZWQUERYSYSTEMINFORMATION)MmGetSystemRoutineAddress(&RoutineName);
  if (!ZwQSI) return NULL;
  Status = ZwQSI(SystemModuleInformation, &SystemInfoBufferSize, 0,
                 &SystemInfoBufferSize);
  if (!SystemInfoBufferSize) {
    LogError("Err, ZwQuerySystemInformation (1) failed");
    return NULL;
  }
  SystemInfoBuffer =
      (PSYSTEM_MODULE_INFORMATION)PlatformMemAllocateNonPagedPool(
          SystemInfoBufferSize * 2);
  if (!SystemInfoBuffer) {
    LogError("Err, insufficient memory");
    return NULL;
  }
  memset(SystemInfoBuffer, 0, SystemInfoBufferSize * 2);
  Status = ZwQSI(SystemModuleInformation, SystemInfoBuffer,
                 SystemInfoBufferSize * 2, &SystemInfoBufferSize);
  if (NT_SUCCESS(Status)) {
    ModuleBase = SystemInfoBuffer->Module.ImageBase;
    if (ImageSize) *ImageSize = SystemInfoBuffer->Module.ImageSize;
  } else {
    LogError("Err, ZwQuerySystemInformation (2) failed");
    return NULL;
  }
  ExFreePool(SystemInfoBuffer);
  return ModuleBase;
}

BOOLEAN SyscallHookFindSsdt(PUINT64 NtTable, PUINT64 Win32kTable) {
  ULONG KernelSize = 0;
  ULONG_PTR KernelBase;
  const unsigned char KiSystemServiceStartPattern[] = {
      0x8B, 0xF8, 0xC1, 0xEF, 0x07, 0x83, 0xE7,
      0x20, 0x25, 0xFF, 0x0F, 0x00, 0x00};
  const ULONG SignatureSize = sizeof(KiSystemServiceStartPattern);
  BOOLEAN Found = FALSE;
  LONG RelativeOffset = 0;
  ULONG_PTR AddressAfterPattern;
  ULONG_PTR Address;
  SSDTStruct *Shadow;
  PVOID TableNT;
  PVOID TableWin32k;
  KernelBase = (ULONG_PTR)SyscallHookGetKernelBase(&KernelSize);
  if (KernelBase == 0 || KernelSize == 0) return FALSE;
  ULONG KiSSSOffset;
  for (KiSSSOffset = 0; KiSSSOffset < KernelSize - SignatureSize;
       KiSSSOffset++) {
    if (RtlCompareMemory(((unsigned char *)KernelBase + KiSSSOffset),
                         KiSystemServiceStartPattern,
                         SignatureSize) == SignatureSize) {
      Found = TRUE;
      break;
    }
  }
  if (!Found) return FALSE;
  AddressAfterPattern = KernelBase + KiSSSOffset + SignatureSize;
  Address = AddressAfterPattern + 7;
  if ((*(unsigned char *)Address == 0x4c) &&
      (*(unsigned char *)(Address + 1) == 0x8d) &&
      (*(unsigned char *)(Address + 2) == 0x1d)) {
    RelativeOffset = *(LONG *)(Address + 3);
  }
  if (RelativeOffset == 0) return FALSE;
  Shadow = (SSDTStruct *)(Address + RelativeOffset + 7);
  TableNT = (PVOID)Shadow;
  TableWin32k = (PVOID)((ULONG_PTR)Shadow + 0x20);
  *NtTable = (UINT64)TableNT;
  *Win32kTable = (UINT64)TableWin32k;
  return TRUE;
}

PVOID SyscallHookGetFunctionAddress(INT32 ApiNumber, BOOLEAN GetFromWin32k) {
  SSDTStruct *SSDT;
  BOOLEAN Result;
  ULONG_PTR SSDTbase;
  UINT64 NtTable, Win32kTable;
  Result = SyscallHookFindSsdt(&NtTable, &Win32kTable);
  if (!Result) {
    LogError("Err, SSDT not found");
    return 0;
  }
  if (!GetFromWin32k) {
    SSDT = (SSDTStruct *)NtTable;
  } else {
    ApiNumber = ApiNumber - 0x1000;
    SSDT = (SSDTStruct *)Win32kTable;
  }
  SSDTbase = (ULONG_PTR)SSDT->pServiceTable;
  if (!SSDTbase) {
    LogError("Err, ServiceTable not found");
    return 0;
  }
  return (PVOID)((SSDT->pServiceTable[ApiNumber] >> 4) + SSDTbase);
}

NTSTATUS NtCreateFileHook(PHANDLE FileHandle, ACCESS_MASK DesiredAccess,
                          POBJECT_ATTRIBUTES ObjectAttributes,
                          PIO_STATUS_BLOCK IoStatusBlock,
                          PLARGE_INTEGER AllocationSize, ULONG FileAttributes,
                          ULONG ShareAccess, ULONG CreateDisposition,
                          ULONG CreateOptions, PVOID EaBuffer, ULONG EaLength) {
  NTSTATUS ConvertStatus;
  UNICODE_STRING ObjectName;
  ANSI_STRING FileNameA;
  ObjectName.Buffer = NULL;
  __try {
    ProbeForRead(FileHandle, sizeof(HANDLE), 1);
    ProbeForRead(ObjectAttributes, sizeof(OBJECT_ATTRIBUTES), 1);
    ProbeForRead(ObjectAttributes->ObjectName, sizeof(UNICODE_STRING), 1);
    ProbeForRead(ObjectAttributes->ObjectName->Buffer,
                 ObjectAttributes->ObjectName->Length, 1);
    FileHandle = *FileHandle;
    ObjectName.Length = ObjectAttributes->ObjectName->Length;
    ObjectName.MaximumLength = ObjectAttributes->ObjectName->MaximumLength;
    ObjectName.Buffer =
        PlatformMemAllocateZeroedNonPagedPool(ObjectName.MaximumLength);
    RtlCopyUnicodeString(&ObjectName, ObjectAttributes->ObjectName);
    ConvertStatus = RtlUnicodeStringToAnsiString(
        &FileNameA, ObjectAttributes->ObjectName, TRUE);
    LogInfo("NtCreateFile called for : %s", FileNameA.Buffer);
  } __except (EXCEPTION_EXECUTE_HANDLER) {
  }
  if (ObjectName.Buffer) {
    PlatformMemFreePool(ObjectName.Buffer);
  }
  return NtCreateFileOrig(FileHandle, DesiredAccess, ObjectAttributes,
                          IoStatusBlock, AllocationSize, FileAttributes,
                          ShareAccess, CreateDisposition, CreateOptions,
                          EaBuffer, EaLength);
}

VOID SyscallHookTest() {
  INT32 ApiNumberOfNtCreateFile = 0x0055;
  PVOID ApiLocationFromSSDTOfNtCreateFile =
      SyscallHookGetFunctionAddress(ApiNumberOfNtCreateFile, FALSE);
  if (!ApiLocationFromSSDTOfNtCreateFile) {
    LogError("Err, address finding for syscall base address");
    return;
  }
  if (EptHookInlineHook(&g_GuestState[KeGetCurrentProcessorNumberEx(NULL)],
                        ApiLocationFromSSDTOfNtCreateFile,
                        (PVOID)NtCreateFileHook,
                        HANDLE_TO_UINT32(PsGetCurrentProcessId()))) {
    LogInfo("Hook appkied to address of API Number : 0x%x at %llx\n",
            ApiNumberOfNtCreateFile, ApiLocationFromSSDTOfNtCreateFile);
  }
}
