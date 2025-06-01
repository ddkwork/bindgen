
#include "pch.h"

BOOLEAN ReadPhysicalMemoryUsingMapIoSpace(PVOID PhysicalAddress, PVOID Buffer,
                                          SIZE_T BufferSize) {
  PHYSICAL_ADDRESS PhysicalAddressTemp = {0};
  PhysicalAddressTemp.QuadPart = (LONGLONG)PhysicalAddress;
  PVOID VirtualAddress = MmMapIoSpaceEx(PhysicalAddressTemp, BufferSize,
                                        PAGE_READWRITE | PAGE_NOCACHE);
  if (VirtualAddress == NULL) {
    return FALSE;
  }
  RtlCopyMemory(Buffer, VirtualAddress, BufferSize);
  MmUnmapIoSpace(VirtualAddress, BufferSize);
  return TRUE;
}

BOOLEAN WritePhysicalMemoryUsingMapIoSpace(PVOID PhysicalAddress, PVOID Buffer,
                                           SIZE_T BufferSize) {
  PHYSICAL_ADDRESS PhysicalAddressTemp = {0};
  PhysicalAddressTemp.QuadPart = (LONGLONG)PhysicalAddress;
  PVOID VirtualAddress = MmMapIoSpaceEx(PhysicalAddressTemp, BufferSize,
                                        PAGE_READWRITE | PAGE_NOCACHE);
  if (VirtualAddress == NULL) {
    return FALSE;
  }
  RtlCopyMemory(VirtualAddress, Buffer, BufferSize);
  MmUnmapIoSpace(VirtualAddress, BufferSize);
  return TRUE;
}

BOOLEAN MemoryManagerReadProcessMemoryNormal(HANDLE PID, PVOID Address,
                                             DEBUGGER_READ_MEMORY_TYPE MemType,
                                             PVOID UserBuffer, SIZE_T Size,
                                             PSIZE_T ReturnSize) {
  PEPROCESS SourceProcess;
  MM_COPY_ADDRESS CopyAddress = {0};
  KAPC_STATE State = {0};
  PHYSICAL_ADDRESS TempPhysicalAddress = {0};
  if (PsGetCurrentProcessId() != PID &&
      MemType == DEBUGGER_READ_VIRTUAL_ADDRESS) {
    if (PsLookupProcessByProcessId(PID, &SourceProcess) != STATUS_SUCCESS) {
      return FALSE;
    }
    __try {
      KeStackAttachProcess(SourceProcess, &State);
      TempPhysicalAddress = MmGetPhysicalAddress(Address);
      KeUnstackDetachProcess(&State);
      CopyAddress.PhysicalAddress.QuadPart = TempPhysicalAddress.QuadPart;
      MmCopyMemory(UserBuffer, CopyAddress, Size, MM_COPY_MEMORY_PHYSICAL,
                   ReturnSize);
      ObDereferenceObject(SourceProcess);
      return TRUE;
    } __except (EXCEPTION_EXECUTE_HANDLER) {
      KeUnstackDetachProcess(&State);
      ObDereferenceObject(SourceProcess);
      return FALSE;
    }
  } else {
    __try {
      if (MemType == DEBUGGER_READ_VIRTUAL_ADDRESS) {
        CopyAddress.VirtualAddress = Address;
        MmCopyMemory(UserBuffer, CopyAddress, Size, MM_COPY_MEMORY_VIRTUAL,
                     ReturnSize);
      } else if (MemType == DEBUGGER_READ_PHYSICAL_ADDRESS) {
        if (!CheckAddressPhysical((UINT64)Address)) {
          return FALSE;
        }
        CopyAddress.PhysicalAddress.QuadPart = (LONGLONG)Address;
        if (MmCopyMemory(UserBuffer, CopyAddress, Size, MM_COPY_MEMORY_PHYSICAL,
                         ReturnSize) != STATUS_SUCCESS &&
            *ReturnSize == 0) {
          if (AsmVmxVmcall(VMCALL_READ_PHYSICAL_MEMORY, (UINT64)Address,
                           (UINT64)UserBuffer,
                           (UINT64)Size) == STATUS_SUCCESS) {
            *ReturnSize = Size;
            return TRUE;
          } else {
            *ReturnSize = 0;
            return FALSE;
          }
        }
      } else {
        return FALSE;
      }
      return TRUE;
    } __except (EXCEPTION_EXECUTE_HANDLER) {
      return FALSE;
    }
  }
}

BOOLEAN MemoryManagerWritePhysicalMemoryNormal(PVOID TargetAddress,
                                               PVOID UserBuffer, SIZE_T Size) {
  if (AsmVmxVmcall(VMCALL_WRITE_PHYSICAL_MEMORY, (UINT64)TargetAddress,
                   (UINT64)UserBuffer, (UINT64)Size) == STATUS_SUCCESS) {
    return TRUE;
  } else {
    return FALSE;
  }
}
