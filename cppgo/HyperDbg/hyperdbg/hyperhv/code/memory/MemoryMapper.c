
#include "pch.h"

_Use_decl_annotations_

    UINT64
    MemoryMapperGetIndex(PAGING_LEVEL Level, UINT64 Va) {
  UINT64 Result = Va;
  Result >>= 12 + Level * 9;
  return Result;
}

_Use_decl_annotations_

    UINT32
    MemoryMapperGetOffset(PAGING_LEVEL Level, UINT64 Va) {
  UINT32 Result = (UINT32)MemoryMapperGetIndex(Level, Va);
  Result &= (1 << 9) - 1;
  return Result;
}

_Use_decl_annotations_

    PVOID
    MemoryMapperGetPteVa(PVOID Va, PAGING_LEVEL Level) {
  CR3_TYPE Cr3;
  Cr3.Flags = __readcr3();
  return MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, Cr3);
}

_Use_decl_annotations_

    PVOID
    MemoryMapperGetPteVaByCr3(PVOID Va, PAGING_LEVEL Level,
                              CR3_TYPE TargetCr3) {
  PPAGE_ENTRY PageEntry = NULL;
  CR3_TYPE CurrentProcessCr3 = {0};
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(TargetCr3);
  PageEntry = MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, TargetCr3);
  SwitchToPreviousProcess(CurrentProcessCr3);
  return PageEntry;
}

_Use_decl_annotations_

    PVOID
    MemoryMapperGetPteVaOnTargetProcess(PVOID Va, PAGING_LEVEL Level) {
  PPAGE_ENTRY PageEntry = NULL;
  CR3_TYPE GuestCr3;
  CR3_TYPE CurrentProcessCr3 = {0};
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(GuestCr3);
  PageEntry = MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, GuestCr3);
  SwitchToPreviousProcess(CurrentProcessCr3);
  return PageEntry;
}

BOOLEAN MemoryMapperCheckPteIsPresentOnTargetProcess(PVOID Va,
                                                     PAGING_LEVEL Level) {
  PPAGE_ENTRY PageEntry = NULL;
  CR3_TYPE GuestCr3;
  CR3_TYPE CurrentProcessCr3 = {0};
  BOOLEAN Result = FALSE;
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(GuestCr3);
  PageEntry = MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, GuestCr3);
  if (PageEntry == NULL) {
    Result = FALSE;
  } else {
    if (PageEntry->Fields.Present == TRUE) {
      Result = TRUE;
    } else {
      Result = FALSE;
    }
  }
  SwitchToPreviousProcess(CurrentProcessCr3);
  return Result;
}

_Use_decl_annotations_

    PVOID
    MemoryMapperSetExecuteDisableToPteOnTargetProcess(PVOID Va, BOOLEAN Set) {
  PPAGE_ENTRY PageEntry = NULL;
  CR3_TYPE GuestCr3;
  CR3_TYPE CurrentProcessCr3 = {0};
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(GuestCr3);
  PageEntry = MemoryMapperGetPteVaWithoutSwitchingByCr3(
      Va, PagingLevelPageTable, GuestCr3);
  PageEntry->Fields.ExecuteDisable = Set;
  __invlpg(Va);
  SwitchToPreviousProcess(CurrentProcessCr3);
  return PageEntry;
}

_Use_decl_annotations_

    PVOID
    MemoryMapperGetPteVaWithoutSwitchingByCr3(PVOID Va, PAGING_LEVEL Level,
                                              CR3_TYPE TargetCr3) {
  CR3_TYPE Cr3;
  UINT64 TempCr3;
  PUINT64 Cr3Va;
  PUINT64 PdptVa;
  PUINT64 PdVa;
  PUINT64 PtVa;
  UINT32 Offset;
  Cr3.Flags = TargetCr3.Flags;
  TempCr3 = Cr3.Fields.PageFrameNumber << 12;
  Cr3Va = (UINT64 *)PhysicalAddressToVirtualAddress(TempCr3);
  if (Cr3Va == NULL) {
    return NULL;
  }
  Offset = MemoryMapperGetOffset(PagingLevelPageMapLevel4, (UINT64)Va);
  PPAGE_ENTRY Pml4e = (PAGE_ENTRY *)&Cr3Va[Offset];
  if (!Pml4e->Fields.Present || Level == PagingLevelPageMapLevel4) {
    return Pml4e;
  }
  PdptVa = (UINT64 *)PhysicalAddressToVirtualAddress(
      Pml4e->Fields.PageFrameNumber << 12);
  if (PdptVa == NULL) {
    return NULL;
  }
  Offset =
      MemoryMapperGetOffset(PagingLevelPageDirectoryPointerTable, (UINT64)Va);
  PPAGE_ENTRY Pdpte = (PAGE_ENTRY *)&PdptVa[Offset];
  if (!Pdpte->Fields.Present || Pdpte->Fields.LargePage ||
      Level == PagingLevelPageDirectoryPointerTable) {
    return Pdpte;
  }
  PdVa = (UINT64 *)PhysicalAddressToVirtualAddress(Pdpte->Fields.PageFrameNumber
                                                   << 12);
  if (PdVa == NULL) {
    return NULL;
  }
  Offset = MemoryMapperGetOffset(PagingLevelPageDirectory, (UINT64)Va);
  PPAGE_ENTRY Pde = (PAGE_ENTRY *)&PdVa[Offset];
  if (!Pde->Fields.Present || Pde->Fields.LargePage ||
      Level == PagingLevelPageDirectory) {
    return Pde;
  }
  PtVa = (UINT64 *)PhysicalAddressToVirtualAddress(Pde->Fields.PageFrameNumber
                                                   << 12);
  if (PtVa == NULL) {
    return NULL;
  }
  Offset = MemoryMapperGetOffset(PagingLevelPageTable, (UINT64)Va);
  PPAGE_ENTRY Pt = (PAGE_ENTRY *)&PtVa[Offset];
  return Pt;
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperCheckIfPageIsPresentByCr3(PVOID Va, CR3_TYPE TargetCr3) {
  PPAGE_ENTRY PageEntry;
  PageEntry = MemoryMapperGetPteVaByCr3(Va, PagingLevelPageTable, TargetCr3);
  if (PageEntry != NULL && PageEntry->Fields.Present) {
    return TRUE;
  } else {
    return FALSE;
  }
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperCheckIfPageIsNxBitSetByCr3(PVOID Va, CR3_TYPE TargetCr3) {
  PPAGE_ENTRY PageEntry;
  PageEntry = MemoryMapperGetPteVaByCr3(Va, PagingLevelPageTable, TargetCr3);
  if (PageEntry != NULL && !PageEntry->Fields.ExecuteDisable) {
    return TRUE;
  } else {
    return FALSE;
  }
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess(PVOID Va) {
  BOOLEAN Result;
  CR3_TYPE GuestCr3;
  PPAGE_ENTRY PageEntry;
  CR3_TYPE CurrentProcessCr3 = {0};
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(GuestCr3);
  PageEntry = MemoryMapperGetPteVa(Va, PagingLevelPageTable);
  if (PageEntry != NULL && !PageEntry->Fields.ExecuteDisable) {
    Result = TRUE;
  } else {
    Result = FALSE;
  }
  SwitchToPreviousProcess(CurrentProcessCr3);
  return Result;
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperCheckIfPdeIsLargePageOnTargetProcess(PVOID Va) {
  BOOLEAN Result;
  CR3_TYPE GuestCr3;
  PPAGE_ENTRY PageEntry;
  CR3_TYPE CurrentProcessCr3 = {0};
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(GuestCr3);
  PageEntry = MemoryMapperGetPteVa(Va, PagingLevelPageDirectory);
  if (PageEntry != NULL && PageEntry->Fields.LargePage) {
    Result = TRUE;
  } else {
    Result = FALSE;
  }
  SwitchToPreviousProcess(CurrentProcessCr3);
  return Result;
}

_Use_decl_annotations_

    PVOID
    MemoryMapperMapReservedPageRange(SIZE_T Size) {
  return MmAllocateMappingAddress(Size, POOLTAG);
}

_Use_decl_annotations_

    VOID
    MemoryMapperUnmapReservedPageRange(PVOID VirtualAddress) {
  MmFreeMappingAddress(VirtualAddress, POOLTAG);
}

_Use_decl_annotations_

    PVOID
    MemoryMapperGetPte(PVOID VirtualAddress) {
  return MemoryMapperGetPteVa(VirtualAddress, PagingLevelPageTable);
}

_Use_decl_annotations_

    PVOID
    MemoryMapperGetPteByCr3(PVOID VirtualAddress, CR3_TYPE TargetCr3) {
  return MemoryMapperGetPteVaByCr3(VirtualAddress, PagingLevelPageTable,
                                   TargetCr3);
}

_Use_decl_annotations_

    PVOID
    MemoryMapperMapPageAndGetPte(PUINT64 PteAddress) {
  PVOID Va;
  UINT64 Pte;
  Va = MemoryMapperMapReservedPageRange(PAGE_SIZE);
  Pte = (UINT64)MemoryMapperGetPte(Va);
  *PteAddress = Pte;
  return Va;
}

VOID MemoryMapperInitialize() {
  UINT64 TempPte;
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (g_MemoryMapper != NULL) {
    return;
  }
  g_MemoryMapper = PlatformMemAllocateZeroedNonPagedPool(
      sizeof(MEMORY_MAPPER_ADDRESSES) * ProcessorsCount);
  for (size_t i = 0; i < ProcessorsCount; i++) {
    g_MemoryMapper[i].VirualAddressForRead =
        (UINT64)MemoryMapperMapPageAndGetPte(&TempPte);
    g_MemoryMapper[i].PteVirtualAddressForRead = TempPte;
    g_MemoryMapper[i].VirualAddressForWrite =
        (UINT64)MemoryMapperMapPageAndGetPte(&TempPte);
    g_MemoryMapper[i].PteVirtualAddressForWrite = TempPte;
  }
}

VOID MemoryMapperUninitialize() {
  ULONG ProcessorsCount = KeQueryActiveProcessorCount(0);
  for (size_t i = 0; i < ProcessorsCount; i++) {
    if (g_MemoryMapper[i].VirualAddressForRead != NULL64_ZERO) {
      MemoryMapperUnmapReservedPageRange(
          (PVOID)g_MemoryMapper[i].VirualAddressForRead);
    }
    if (g_MemoryMapper[i].VirualAddressForWrite != NULL64_ZERO) {
      MemoryMapperUnmapReservedPageRange(
          (PVOID)g_MemoryMapper[i].VirualAddressForWrite);
    }
    g_MemoryMapper[i].VirualAddressForRead = NULL64_ZERO;
    g_MemoryMapper[i].PteVirtualAddressForRead = NULL64_ZERO;
    g_MemoryMapper[i].VirualAddressForWrite = NULL64_ZERO;
    g_MemoryMapper[i].PteVirtualAddressForWrite = NULL64_ZERO;
  }
  g_MemoryMapper = NULL;
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperReadMemorySafeByPte(PHYSICAL_ADDRESS PaAddressToRead,
                                    PVOID BufferToSaveMemory, SIZE_T SizeToRead,
                                    UINT64 PteVaAddress, UINT64 MappingVa,
                                    BOOLEAN InvalidateVpids) {
  PVOID NewAddress;
  PAGE_ENTRY PageEntry;
  PPAGE_ENTRY Pte = (PAGE_ENTRY *)PteVaAddress;
  PVOID Va = (PVOID)MappingVa;
  PageEntry.Flags = Pte->Flags;
  PageEntry.Fields.Present = 1;
  PageEntry.Fields.Write = 1;
  PageEntry.Fields.Global = 1;
  PageEntry.Fields.PageFrameNumber = PaAddressToRead.QuadPart >> 12;
  Pte->Flags = PageEntry.Flags;
  __invlpg(Va);
  if (InvalidateVpids) {
  }
  NewAddress =
      (PVOID)((UINT64)Va + (PAGE_4KB_OFFSET & (PaAddressToRead.QuadPart)));
  memcpy(BufferToSaveMemory, NewAddress, SizeToRead);
  Pte->Flags = NULL64_ZERO;
  return TRUE;
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperWriteMemorySafeByPte(PVOID SourceVA,
                                     PHYSICAL_ADDRESS PaAddressToWrite,
                                     SIZE_T SizeToWrite, UINT64 PteVaAddress,
                                     UINT64 MappingVa,
                                     BOOLEAN InvalidateVpids) {
  PVOID NewAddress;
  PAGE_ENTRY PageEntry;
  PPAGE_ENTRY Pte = (PAGE_ENTRY *)PteVaAddress;
  PVOID Va = (PVOID)MappingVa;
  PageEntry.Flags = Pte->Flags;
  PageEntry.Fields.Present = 1;
  PageEntry.Fields.Write = 1;
  PageEntry.Fields.Global = 1;
  PageEntry.Fields.PageFrameNumber = PaAddressToWrite.QuadPart >> 12;
  Pte->Flags = PageEntry.Flags;
  __invlpg(Va);
  if (InvalidateVpids) {
  }
  NewAddress =
      (PVOID)((UINT64)Va + (PAGE_4KB_OFFSET & (PaAddressToWrite.QuadPart)));
  memcpy(NewAddress, SourceVA, SizeToWrite);
  Pte->Flags = NULL64_ZERO;
  return TRUE;
}

_Use_decl_annotations_

    UINT64
    MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker(
        MEMORY_MAPPER_WRAPPER_FOR_MEMORY_READ TypeOfRead,
        UINT64 AddressToRead) {
  PHYSICAL_ADDRESS PhysicalAddress = {0};
  switch (TypeOfRead) {
    case MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY:
      PhysicalAddress.QuadPart = AddressToRead;
      break;
    case MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY:
      PhysicalAddress.QuadPart =
          VirtualAddressToPhysicalAddress((PVOID)AddressToRead);
      break;
    default:
      return NULL64_ZERO;
      break;
  }
  return PhysicalAddress.QuadPart;
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperReadMemorySafeByPhysicalAddressWrapper(
        MEMORY_MAPPER_WRAPPER_FOR_MEMORY_READ TypeOfRead, UINT64 AddressToRead,
        UINT64 BufferToSaveMemory, SIZE_T SizeToRead) {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  UINT64 AddressToCheck;
  PHYSICAL_ADDRESS PhysicalAddress;
  if (g_MemoryMapper[CurrentCore].VirualAddressForRead == NULL64_ZERO ||
      g_MemoryMapper[CurrentCore].PteVirtualAddressForRead == NULL64_ZERO) {
    return FALSE;
  }
  AddressToCheck =
      (CHAR *)AddressToRead + SizeToRead - ((CHAR *)PAGE_ALIGN(AddressToRead));
  if (AddressToCheck > PAGE_SIZE) {
    UINT64 ReadSize = AddressToCheck;
    while (SizeToRead != 0) {
      ReadSize = (UINT64)PAGE_ALIGN(AddressToRead + PAGE_SIZE) - AddressToRead;
      if (ReadSize == PAGE_SIZE && SizeToRead < PAGE_SIZE) {
        ReadSize = SizeToRead;
      }

      PhysicalAddress.QuadPart =
          MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker(
              TypeOfRead, AddressToRead);
      if (!MemoryMapperReadMemorySafeByPte(
              PhysicalAddress, (PVOID)BufferToSaveMemory, ReadSize,
              g_MemoryMapper[CurrentCore].PteVirtualAddressForRead,
              g_MemoryMapper[CurrentCore].VirualAddressForRead, FALSE)) {
        return FALSE;
      }
      SizeToRead = SizeToRead - ReadSize;
      AddressToRead = AddressToRead + ReadSize;
      BufferToSaveMemory = BufferToSaveMemory + ReadSize;
    }
    return TRUE;
  } else {
    PhysicalAddress.QuadPart =
        MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker(
            TypeOfRead, AddressToRead);
    return MemoryMapperReadMemorySafeByPte(
        PhysicalAddress, (PVOID)BufferToSaveMemory, SizeToRead,
        g_MemoryMapper[CurrentCore].PteVirtualAddressForRead,
        g_MemoryMapper[CurrentCore].VirualAddressForRead, FALSE);
  }
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperReadMemorySafeByPhysicalAddress(UINT64 PaAddressToRead,
                                                UINT64 BufferToSaveMemory,
                                                SIZE_T SizeToRead) {
  return MemoryMapperReadMemorySafeByPhysicalAddressWrapper(
      MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY, PaAddressToRead,
      BufferToSaveMemory, SizeToRead);
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperReadMemorySafe(UINT64 VaAddressToRead, PVOID BufferToSaveMemory,
                               SIZE_T SizeToRead) {
  return MemoryMapperReadMemorySafeByPhysicalAddressWrapper(
      MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY, VaAddressToRead,
      (UINT64)BufferToSaveMemory, SizeToRead);
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperReadMemorySafeOnTargetProcess(UINT64 VaAddressToRead,
                                              PVOID BufferToSaveMemory,
                                              SIZE_T SizeToRead) {
  CR3_TYPE GuestCr3;
  CR3_TYPE OriginalCr3;
  BOOLEAN Result;
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  OriginalCr3.Flags = __readcr3();
  __writecr3(GuestCr3.Flags);
  Result = MemoryMapperReadMemorySafe(VaAddressToRead, BufferToSaveMemory,
                                      SizeToRead);
  __writecr3(OriginalCr3.Flags);
  return Result;
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperWriteMemorySafeOnTargetProcess(UINT64 Destination, PVOID Source,
                                               SIZE_T Size) {
  CR3_TYPE GuestCr3;
  CR3_TYPE OriginalCr3;
  BOOLEAN Result;
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  OriginalCr3.Flags = __readcr3();
  __writecr3(GuestCr3.Flags);
  Result = MemoryMapperWriteMemorySafe(Destination, Source, Size, GuestCr3);
  __writecr3(OriginalCr3.Flags);
  return Result;
}

_Use_decl_annotations_

    UINT64
    MemoryMapperWriteMemorySafeWrapperAddressMaker(
        MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfWrite,
        UINT64 DestinationAddr, PCR3_TYPE TargetProcessCr3,
        UINT32 TargetProcessId) {
  PHYSICAL_ADDRESS PhysicalAddress = {0};
  switch (TypeOfWrite) {
    case MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY:
      PhysicalAddress.QuadPart = DestinationAddr;
      break;
    case MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE:
      if (TargetProcessId == NULL_ZERO) {
        PhysicalAddress.QuadPart =
            VirtualAddressToPhysicalAddress((PVOID)DestinationAddr);
      } else {
        PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddressByProcessId(
            (PVOID)DestinationAddr, TargetProcessId);
      }
      break;
    case MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE:
      if (TargetProcessCr3 == NULL || TargetProcessCr3->Flags == NULL64_ZERO) {
        PhysicalAddress.QuadPart =
            VirtualAddressToPhysicalAddress((PVOID)DestinationAddr);
      } else {
        PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddressByProcessCr3(
            (PVOID)DestinationAddr, *TargetProcessCr3);
      }
      break;
    default:
      return NULL64_ZERO;
      break;
  }
  return PhysicalAddress.QuadPart;
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperWriteMemorySafeWrapper(
        MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfWrite,
        UINT64 DestinationAddr, UINT64 Source, SIZE_T SizeToWrite,
        PCR3_TYPE TargetProcessCr3, UINT32 TargetProcessId) {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  UINT64 AddressToCheck;
  PHYSICAL_ADDRESS PhysicalAddress;
  if (g_MemoryMapper[CurrentCore].VirualAddressForWrite == NULL64_ZERO ||
      g_MemoryMapper[CurrentCore].PteVirtualAddressForWrite == NULL64_ZERO) {
    return FALSE;
  }
  AddressToCheck = (CHAR *)DestinationAddr + SizeToWrite -
                   ((CHAR *)PAGE_ALIGN(DestinationAddr));
  if (AddressToCheck > PAGE_SIZE) {
    UINT64 WriteSize = AddressToCheck;
    while (SizeToWrite != 0) {
      WriteSize =
          (UINT64)PAGE_ALIGN(DestinationAddr + PAGE_SIZE) - DestinationAddr;
      if (WriteSize == PAGE_SIZE && SizeToWrite < PAGE_SIZE) {
        WriteSize = SizeToWrite;
      }

      PhysicalAddress.QuadPart = MemoryMapperWriteMemorySafeWrapperAddressMaker(
          TypeOfWrite, DestinationAddr, TargetProcessCr3, TargetProcessId);
      if (!MemoryMapperWriteMemorySafeByPte(
              (PVOID)Source, PhysicalAddress, WriteSize,
              g_MemoryMapper[CurrentCore].PteVirtualAddressForWrite,
              g_MemoryMapper[CurrentCore].VirualAddressForWrite, FALSE)) {
        return FALSE;
      }
      SizeToWrite = SizeToWrite - WriteSize;
      DestinationAddr = DestinationAddr + WriteSize;
      Source = Source + WriteSize;
    }
    return TRUE;
  } else {
    PhysicalAddress.QuadPart = MemoryMapperWriteMemorySafeWrapperAddressMaker(
        TypeOfWrite, DestinationAddr, TargetProcessCr3, TargetProcessId);
    return MemoryMapperWriteMemorySafeByPte(
        (PVOID)Source, PhysicalAddress, SizeToWrite,
        g_MemoryMapper[CurrentCore].PteVirtualAddressForWrite,
        g_MemoryMapper[CurrentCore].VirualAddressForWrite, FALSE);
  }
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperWriteMemorySafe(UINT64 Destination, PVOID Source,
                                SIZE_T SizeToWrite, CR3_TYPE TargetProcessCr3) {
  return MemoryMapperWriteMemorySafeWrapper(
      MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE, Destination,
      (UINT64)Source, SizeToWrite, &TargetProcessCr3, NULL_ZERO);
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperWriteMemoryUnsafe(UINT64 Destination, PVOID Source,
                                  SIZE_T SizeToWrite, UINT32 TargetProcessId) {
  return MemoryMapperWriteMemorySafeWrapper(
      MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE, Destination,
      (UINT64)Source, SizeToWrite, NULL, TargetProcessId);
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperWriteMemorySafeByPhysicalAddress(UINT64 DestinationPa,
                                                 UINT64 Source,
                                                 SIZE_T SizeToWrite) {
  return MemoryMapperWriteMemorySafeWrapper(
      MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY, DestinationPa, Source,
      SizeToWrite, NULL, NULL_ZERO);
}

_Use_decl_annotations_

    UINT64
    MemoryMapperReserveUsermodeAddressOnTargetProcess(UINT32 ProcessId,
                                                      BOOLEAN Allocate) {
  NTSTATUS Status;
  PVOID AllocPtr = NULL;
  SIZE_T AllocSize = PAGE_SIZE;
  PEPROCESS SourceProcess;
  KAPC_STATE State = {0};
  if (PsGetCurrentProcessId() != (HANDLE)ProcessId) {
    if (PsLookupProcessByProcessId((HANDLE)ProcessId, &SourceProcess) !=
        STATUS_SUCCESS) {
      return NULL64_ZERO;
    }
    __try {
      KeStackAttachProcess(SourceProcess, &State);
      Status = ZwAllocateVirtualMemory(
          NtCurrentProcess(), &AllocPtr, (ULONG_PTR)NULL, &AllocSize,
          Allocate ? MEM_COMMIT : MEM_RESERVE, PAGE_EXECUTE_READWRITE);
      KeUnstackDetachProcess(&State);
      ObDereferenceObject(SourceProcess);
    } __except (EXCEPTION_EXECUTE_HANDLER) {
      KeUnstackDetachProcess(&State);
      ObDereferenceObject(SourceProcess);
      return NULL64_ZERO;
    }
  } else {
    Status = ZwAllocateVirtualMemory(
        NtCurrentProcess(), &AllocPtr, (ULONG_PTR)NULL, &AllocSize,
        Allocate ? MEM_COMMIT : MEM_RESERVE, PAGE_EXECUTE_READWRITE);
  }
  if (!NT_SUCCESS(Status)) {
    return NULL64_ZERO;
  }
  return (UINT64)AllocPtr;
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperFreeMemoryOnTargetProcess(UINT32 ProcessId, PVOID BaseAddress) {
  NTSTATUS Status;
  SIZE_T AllocSize = PAGE_SIZE;
  PEPROCESS SourceProcess;
  KAPC_STATE State = {0};
  if (PsGetCurrentProcessId() != (HANDLE)ProcessId) {
    if (PsLookupProcessByProcessId((HANDLE)ProcessId, &SourceProcess) !=
        STATUS_SUCCESS) {
      return FALSE;
    }
    __try {
      KeStackAttachProcess(SourceProcess, &State);
      Status = ZwFreeVirtualMemory(NtCurrentProcess(), &BaseAddress, &AllocSize,
                                   MEM_RELEASE);
      KeUnstackDetachProcess(&State);
      ObDereferenceObject(SourceProcess);
    } __except (EXCEPTION_EXECUTE_HANDLER) {
      KeUnstackDetachProcess(&State);
      ObDereferenceObject(SourceProcess);
      return FALSE;
    }
  } else {
    Status = ZwFreeVirtualMemory(NtCurrentProcess(), &BaseAddress, &AllocSize,
                                 MEM_RELEASE);
  }
  if (!NT_SUCCESS(Status)) {
    return FALSE;
  }
  return TRUE;
}

_Use_decl_annotations_

    VOID
    MemoryMapperMapPhysicalAddressToPte(PHYSICAL_ADDRESS PhysicalAddress,
                                        PVOID TargetProcessVirtualAddress,
                                        CR3_TYPE TargetProcessKernelCr3) {
  PPAGE_ENTRY PreviousPteEntry;
  PAGE_ENTRY PageEntry;
  CR3_TYPE CurrentProcessCr3;
  PreviousPteEntry =
      MemoryMapperGetPteVaByCr3(TargetProcessVirtualAddress,
                                PagingLevelPageTable, TargetProcessKernelCr3);
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(TargetProcessKernelCr3);
  PageEntry.Flags = PreviousPteEntry->Flags;
  PageEntry.Fields.Present = 1;
  PageEntry.Fields.Supervisor = 1;
  PageEntry.Fields.Write = 1;
  PageEntry.Fields.Global = 1;
  PageEntry.Fields.PageFrameNumber = PhysicalAddress.QuadPart >> 12;
  PreviousPteEntry->Flags = PageEntry.Flags;
  __invlpg(TargetProcessVirtualAddress);
  SwitchToPreviousProcess(CurrentProcessCr3);
}

_Use_decl_annotations_

    BOOLEAN
    MemoryMapperSetSupervisorBitWithoutSwitchingByCr3(PVOID Va, BOOLEAN Set,
                                                      PAGING_LEVEL Level,
                                                      CR3_TYPE TargetCr3) {
  PPAGE_ENTRY Pml = NULL;
  Pml = MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, TargetCr3);
  if (!Pml) {
    return FALSE;
  }
  if (Set) {
    Pml->Fields.Supervisor = 1;
  } else {
    Pml->Fields.Supervisor = 0;
  }
  return TRUE;
}
