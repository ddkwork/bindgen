
#include "pch.h"

BOOLEAN DebuggerCommandReadMemory(PDEBUGGER_READ_MEMORY ReadMemRequest,
                                  PVOID UserBuffer, PSIZE_T ReturnSize) {
  UINT32 Pid;
  UINT32 Size;
  UINT64 Address;
  DEBUGGER_READ_MEMORY_TYPE MemType;
  BOOLEAN Is32BitProcess = FALSE;
  Pid = ReadMemRequest->Pid;
  Size = ReadMemRequest->Size;
  Address = ReadMemRequest->Address;
  MemType = ReadMemRequest->MemoryType;
  if (Size && Address != (UINT64)NULL) {
    if (MemoryManagerReadProcessMemoryNormal((HANDLE)Pid, (PVOID)Address,
                                             MemType, (PVOID)UserBuffer, Size,
                                             ReturnSize)) {
      if (ReadMemRequest->MemoryType == DEBUGGER_READ_VIRTUAL_ADDRESS &&
          ReadMemRequest->GetAddressMode) {
        if (ReadMemRequest->Address >= 0xFFFF800000000000 &&
            ReadMemRequest->Address <= 0xFFFFFFFFFFFFFFFF) {
          ReadMemRequest->AddressMode = DEBUGGER_READ_ADDRESS_MODE_64_BIT;
        } else {
          if (UserAccessIsWow64Process((HANDLE)ReadMemRequest->Pid,
                                       &Is32BitProcess)) {
            if (Is32BitProcess) {
              ReadMemRequest->AddressMode = DEBUGGER_READ_ADDRESS_MODE_32_BIT;
            } else {
              ReadMemRequest->AddressMode = DEBUGGER_READ_ADDRESS_MODE_64_BIT;
            }
          } else {
            ReadMemRequest->AddressMode = DEBUGGER_READ_ADDRESS_MODE_64_BIT;
          }
        }
      }
      ReadMemRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
      return TRUE;
    } else {
      ReadMemRequest->KernelStatus =
          DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER;
      return FALSE;
    }
  } else {
    ReadMemRequest->KernelStatus =
        DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER;
    return FALSE;
  }
}

BOOLEAN DebuggerCommandReadMemoryVmxRoot(PDEBUGGER_READ_MEMORY ReadMemRequest,
                                         UCHAR *UserBuffer,
                                         UINT32 *ReturnSize) {
  UINT32 Pid;
  UINT32 Size;
  UINT64 Address;
  UINT64 OffsetInUserBuffer;
  DEBUGGER_READ_MEMORY_TYPE MemType;
  BOOLEAN Is32BitProcess = FALSE;
  PLIST_ENTRY TempList = 0;
  Pid = ReadMemRequest->Pid;
  Size = ReadMemRequest->Size;
  Address = ReadMemRequest->Address;
  MemType = ReadMemRequest->MemoryType;
  if (MemType == DEBUGGER_READ_PHYSICAL_ADDRESS) {
    if (!CheckAddressPhysical(Address)) {
      ReadMemRequest->KernelStatus = DEBUGGER_ERROR_INVALID_PHYSICAL_ADDRESS;
      return FALSE;
    }
    MemoryMapperReadMemorySafeByPhysicalAddress(Address, (UINT64)UserBuffer,
                                                Size);
  } else if (MemType == DEBUGGER_READ_VIRTUAL_ADDRESS) {
    if (!CheckAccessValidityAndSafety(Address, Size)) {
      ReadMemRequest->KernelStatus = DEBUGGER_ERROR_INVALID_ADDRESS;
      return FALSE;
    }
    MemoryMapperReadMemorySafeOnTargetProcess(Address, UserBuffer, Size);
    TempList = &g_BreakpointsListHead;
    while (&g_BreakpointsListHead != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc =
          CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
      if (CurrentBreakpointDesc->Address >= Address &&
          CurrentBreakpointDesc->Address <= Address + Size) {
        OffsetInUserBuffer = CurrentBreakpointDesc->Address - Address;
        if (UserBuffer[OffsetInUserBuffer] == 0xcc) {
          UserBuffer[OffsetInUserBuffer] = CurrentBreakpointDesc->PreviousByte;
        }
      }
    }
  } else {
    ReadMemRequest->KernelStatus = DEBUGGER_ERROR_MEMORY_TYPE_INVALID;
    return FALSE;
  }
  if (ReadMemRequest->MemoryType == DEBUGGER_READ_VIRTUAL_ADDRESS &&
      ReadMemRequest->GetAddressMode) {
    if (ReadMemRequest->Address >= 0xFFFF800000000000 &&
        ReadMemRequest->Address <= 0xFFFFFFFFFFFFFFFF) {
      ReadMemRequest->AddressMode = DEBUGGER_READ_ADDRESS_MODE_64_BIT;
    } else {
      if (UserAccessIsWow64ProcessByEprocess(PsGetCurrentProcess(),
                                             &Is32BitProcess)) {
        if (Is32BitProcess) {
          ReadMemRequest->AddressMode = DEBUGGER_READ_ADDRESS_MODE_32_BIT;
        } else {
          ReadMemRequest->AddressMode = DEBUGGER_READ_ADDRESS_MODE_64_BIT;
        }
      } else {
        ReadMemRequest->AddressMode = DEBUGGER_READ_ADDRESS_MODE_64_BIT;
      }
    }
  }
  ReadMemRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  *ReturnSize = Size;
  return TRUE;
}

NTSTATUS DebuggerReadOrWriteMsr(
    PDEBUGGER_READ_AND_WRITE_ON_MSR ReadOrWriteMsrRequest, UINT64 *UserBuffer,
    PSIZE_T ReturnSize) {
  NTSTATUS Status;
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (ReadOrWriteMsrRequest->ActionType == DEBUGGER_MSR_WRITE) {
    if (ReadOrWriteMsrRequest->CoreNumber ==
        DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES) {
      for (size_t i = 0; i < ProcessorsCount; i++) {
        g_DbgState[i].MsrState.Msr = ReadOrWriteMsrRequest->Msr;
        g_DbgState[i].MsrState.Value = ReadOrWriteMsrRequest->Value;
      }
      KeGenericCallDpc(DpcRoutineWriteMsrToAllCores, 0x0);
    } else {
      if (ReadOrWriteMsrRequest->CoreNumber >= ProcessorsCount) {
        return STATUS_INVALID_PARAMETER;
      }
      g_DbgState[ReadOrWriteMsrRequest->CoreNumber].MsrState.Msr =
          ReadOrWriteMsrRequest->Msr;
      g_DbgState[ReadOrWriteMsrRequest->CoreNumber].MsrState.Value =
          ReadOrWriteMsrRequest->Value;
      Status =
          DpcRoutineRunTaskOnSingleCore(ReadOrWriteMsrRequest->CoreNumber,
                                        (PVOID)DpcRoutinePerformWriteMsr, NULL);
      *ReturnSize = 0;
      return Status;
    }
    *ReturnSize = 0;
    return STATUS_SUCCESS;
  } else if (ReadOrWriteMsrRequest->ActionType == DEBUGGER_MSR_READ) {
    if (ReadOrWriteMsrRequest->CoreNumber ==
        DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES) {
      for (size_t i = 0; i < ProcessorsCount; i++) {
        g_DbgState[i].MsrState.Msr = ReadOrWriteMsrRequest->Msr;
      }
      KeGenericCallDpc(DpcRoutineReadMsrToAllCores, 0x0);
      for (size_t i = 0; i < ProcessorsCount; i++) {
        UserBuffer[i] = g_DbgState[i].MsrState.Value;
      }
      *ReturnSize = sizeof(UINT64) * ProcessorsCount;
      return STATUS_SUCCESS;
    } else {
      if (ReadOrWriteMsrRequest->CoreNumber >= ProcessorsCount) {
        *ReturnSize = 0;
        return STATUS_INVALID_PARAMETER;
      }
      g_DbgState[ReadOrWriteMsrRequest->CoreNumber].MsrState.Msr =
          ReadOrWriteMsrRequest->Msr;
      Status =
          DpcRoutineRunTaskOnSingleCore(ReadOrWriteMsrRequest->CoreNumber,
                                        (PVOID)DpcRoutinePerformReadMsr, NULL);
      if (Status != STATUS_SUCCESS) {
        *ReturnSize = 0;
        return Status;
      }
      UserBuffer[0] =
          g_DbgState[ReadOrWriteMsrRequest->CoreNumber].MsrState.Value;
      *ReturnSize = sizeof(UINT64);
      return STATUS_SUCCESS;
    }
  }
  *ReturnSize = 0;
  return STATUS_UNSUCCESSFUL;
}

NTSTATUS DebuggerCommandEditMemory(PDEBUGGER_EDIT_MEMORY EditMemRequest) {
  UINT32 LengthOfEachChunk = 0;
  PVOID DestinationAddress = 0;
  PVOID SourceAddress = 0;
  if (EditMemRequest->ByteSize == EDIT_BYTE) {
    LengthOfEachChunk = 1;
  } else if (EditMemRequest->ByteSize == EDIT_DWORD) {
    LengthOfEachChunk = 4;
  } else if (EditMemRequest->ByteSize == EDIT_QWORD) {
    LengthOfEachChunk = 8;
  } else {
    EditMemRequest->Result =
        DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
    return STATUS_UNSUCCESSFUL;
  }
  if (EditMemRequest->MemoryType == EDIT_VIRTUAL_MEMORY) {
    if (EditMemRequest->ProcessId ==
            HANDLE_TO_UINT32(PsGetCurrentProcessId()) &&
        VirtualAddressToPhysicalAddress((PVOID)EditMemRequest->Address) == 0) {
      EditMemRequest->Result =
          DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS;
      return STATUS_UNSUCCESSFUL;
    } else if (VirtualAddressToPhysicalAddressByProcessId(
                   (PVOID)EditMemRequest->Address, EditMemRequest->ProcessId) ==
               0) {
      EditMemRequest->Result =
          DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS;
      return STATUS_UNSUCCESSFUL;
    }
    for (size_t i = 0; i < EditMemRequest->CountOf64Chunks; i++) {
      DestinationAddress =
          (PVOID)((UINT64)EditMemRequest->Address + (i * LengthOfEachChunk));
      SourceAddress =
          (PVOID)((UINT64)EditMemRequest + SIZEOF_DEBUGGER_EDIT_MEMORY +
                  (i * sizeof(UINT64)));
      MemoryMapperWriteMemoryUnsafe((UINT64)DestinationAddress, SourceAddress,
                                    LengthOfEachChunk,
                                    EditMemRequest->ProcessId);
    }
  } else if (EditMemRequest->MemoryType == EDIT_PHYSICAL_MEMORY) {
    if (!CheckAddressPhysical(EditMemRequest->Address)) {
      EditMemRequest->Result = DEBUGGER_ERROR_INVALID_ADDRESS;
      return STATUS_UNSUCCESSFUL;
    }
    for (size_t i = 0; i < EditMemRequest->CountOf64Chunks; i++) {
      DestinationAddress =
          (PVOID)((UINT64)EditMemRequest->Address + (i * LengthOfEachChunk));
      SourceAddress =
          (PVOID)((UINT64)EditMemRequest + SIZEOF_DEBUGGER_EDIT_MEMORY +
                  (i * sizeof(UINT64)));
      if (MemoryManagerWritePhysicalMemoryNormal(
              (PVOID)DestinationAddress, (PVOID)SourceAddress,
              (SIZE_T)LengthOfEachChunk) == FALSE) {
        EditMemRequest->Result = DEBUGGER_ERROR_INVALID_ADDRESS;
        return STATUS_UNSUCCESSFUL;
      }
    }
  } else {
    EditMemRequest->Result =
        DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
    return STATUS_UNSUCCESSFUL;
  }
  EditMemRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return STATUS_SUCCESS;
}

BOOLEAN DebuggerCommandEditMemoryVmxRoot(PDEBUGGER_EDIT_MEMORY EditMemRequest) {
  UINT32 LengthOfEachChunk = 0;
  PVOID DestinationAddress = 0;
  PVOID SourceAddress = 0;
  if (EditMemRequest->ByteSize == EDIT_BYTE) {
    LengthOfEachChunk = 1;
  } else if (EditMemRequest->ByteSize == EDIT_DWORD) {
    LengthOfEachChunk = 4;
  } else if (EditMemRequest->ByteSize == EDIT_QWORD) {
    LengthOfEachChunk = 8;
  } else {
    EditMemRequest->Result =
        DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
    return FALSE;
  }
  if (EditMemRequest->MemoryType == EDIT_VIRTUAL_MEMORY) {
    if (!CheckAccessValidityAndSafety(
            EditMemRequest->Address,
            EditMemRequest->ByteSize * EditMemRequest->CountOf64Chunks)) {
      EditMemRequest->Result = DEBUGGER_ERROR_INVALID_ADDRESS;
      return FALSE;
    }
    for (size_t i = 0; i < EditMemRequest->CountOf64Chunks; i++) {
      DestinationAddress =
          (PVOID)((UINT64)EditMemRequest->Address + (i * LengthOfEachChunk));
      SourceAddress =
          (PVOID)((UINT64)EditMemRequest + SIZEOF_DEBUGGER_EDIT_MEMORY +
                  (i * sizeof(UINT64)));
      MemoryMapperWriteMemorySafeOnTargetProcess(
          (UINT64)DestinationAddress, SourceAddress, LengthOfEachChunk);
    }
  } else if (EditMemRequest->MemoryType == EDIT_PHYSICAL_MEMORY) {
    if (!CheckAddressPhysical(EditMemRequest->Address)) {
      EditMemRequest->Result = DEBUGGER_ERROR_INVALID_ADDRESS;
      return FALSE;
    }
    for (size_t i = 0; i < EditMemRequest->CountOf64Chunks; i++) {
      DestinationAddress =
          (PVOID)((UINT64)EditMemRequest->Address + (i * LengthOfEachChunk));
      SourceAddress =
          (PVOID)((UINT64)EditMemRequest + SIZEOF_DEBUGGER_EDIT_MEMORY +
                  (i * sizeof(UINT64)));
      MemoryMapperWriteMemorySafeByPhysicalAddress(
          (UINT64)DestinationAddress, (UINT64)SourceAddress, LengthOfEachChunk);
    }
  } else {
    EditMemRequest->Result =
        DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER;
    return FALSE;
  }
  EditMemRequest->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return TRUE;
}

BOOLEAN PerformSearchAddress(UINT64 *AddressToSaveResults,
                             PDEBUGGER_SEARCH_MEMORY SearchMemRequest,
                             UINT64 StartAddress, UINT64 EndAddress,
                             BOOLEAN IsDebuggeePaused,
                             PUINT32 CountOfMatchedCases) {
  UINT32 CountOfOccurance = 0;
  UINT64 Cmp64 = 0;
  UINT32 IndexToArrayOfResults = 0;
  UINT32 LengthOfEachChunk = 0;
  PVOID TempSourceAddress = 0;
  PVOID SourceAddress = 0;
  BOOLEAN StillMatch = FALSE;
  UINT64 TempValue = (UINT64)NULL;
  CR3_TYPE CurrentProcessCr3 = {0};
  if (SearchMemRequest->ByteSize == SEARCH_BYTE) {
    LengthOfEachChunk = 1;
  } else if (SearchMemRequest->ByteSize == SEARCH_DWORD) {
    LengthOfEachChunk = 4;
  } else if (SearchMemRequest->ByteSize == SEARCH_QWORD) {
    LengthOfEachChunk = 8;
  } else {
    return FALSE;
  }
  if (SearchMemRequest->MemoryType == SEARCH_VIRTUAL_MEMORY ||
      SearchMemRequest->MemoryType == SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY) {
    if (IsDebuggeePaused) {
      CurrentProcessCr3 =
          SwitchToProcessMemoryLayoutByCr3(LayoutGetCurrentProcessCr3());
    } else {
      if (SearchMemRequest->ProcessId !=
          HANDLE_TO_UINT32(PsGetCurrentProcessId())) {
        CurrentProcessCr3 =
            SwitchToProcessMemoryLayout(SearchMemRequest->ProcessId);
      }
    }
    SourceAddress =
        (PVOID)((UINT64)SearchMemRequest + SIZEOF_DEBUGGER_SEARCH_MEMORY);
    for (size_t BaseIterator = (size_t)StartAddress;
         BaseIterator < ((UINT64)EndAddress);
         BaseIterator += LengthOfEachChunk) {
      if (IsDebuggeePaused) {
        MemoryMapperReadMemorySafe((UINT64)BaseIterator, &Cmp64,
                                   LengthOfEachChunk);
      } else {
        RtlCopyMemory(&Cmp64, (PVOID)BaseIterator, LengthOfEachChunk);
      }
      TempValue = *(UINT64 *)SourceAddress;
      if (Cmp64 == TempValue) {
        StillMatch = TRUE;
        for (size_t i = LengthOfEachChunk;
             i < SearchMemRequest->CountOf64Chunks; i++) {
          TempSourceAddress =
              (PVOID)((UINT64)SearchMemRequest + SIZEOF_DEBUGGER_SEARCH_MEMORY +
                      (i * sizeof(UINT64)));
          if (IsDebuggeePaused) {
            MemoryMapperReadMemorySafe(
                (UINT64)(BaseIterator + (LengthOfEachChunk * i)), &Cmp64,
                LengthOfEachChunk);
          } else {
            RtlCopyMemory(&Cmp64,
                          (PVOID)(BaseIterator + (LengthOfEachChunk * i)),
                          LengthOfEachChunk);
          }
          if (IsDebuggeePaused) {
            MemoryMapperReadMemorySafe((UINT64)TempSourceAddress, &TempValue,
                                       sizeof(UINT64));
          } else {
            TempValue = *(UINT64 *)TempSourceAddress;
          }
          if (!(Cmp64 == TempValue)) {
            StillMatch = FALSE;
            break;
          }
        }
        if (StillMatch) {
          CountOfOccurance++;
          if (IsDebuggeePaused) {
            if (SearchMemRequest->MemoryType ==
                SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY) {
              Log("%llx\n",
                  VirtualAddressToPhysicalAddress((PVOID)BaseIterator));
            } else {
              Log("%llx\n", BaseIterator);
            }
          } else {
            if (SearchMemRequest->MemoryType ==
                SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY) {
              AddressToSaveResults[IndexToArrayOfResults] =
                  VirtualAddressToPhysicalAddress((PVOID)BaseIterator);
            } else {
              AddressToSaveResults[IndexToArrayOfResults] = BaseIterator;
            }
          }
          if (MaximumSearchResults > IndexToArrayOfResults) {
            IndexToArrayOfResults++;
          } else {
            *CountOfMatchedCases = CountOfOccurance;
            return TRUE;
          }
        }
      } else {
        continue;
      }
    }
    if (IsDebuggeePaused || SearchMemRequest->ProcessId !=
                                HANDLE_TO_UINT32(PsGetCurrentProcessId())) {
      SwitchToPreviousProcess(CurrentProcessCr3);
    }
  } else if (SearchMemRequest->MemoryType == SEARCH_PHYSICAL_MEMORY) {
    LogError(
        "Err, searching physical memory is not allowed without virtual "
        "address");
    return FALSE;
  } else {
    return FALSE;
  }
  *CountOfMatchedCases = CountOfOccurance;
  return TRUE;
}

BOOLEAN SearchAddressWrapper(PUINT64 AddressToSaveResults,
                             PDEBUGGER_SEARCH_MEMORY SearchMemRequest,
                             UINT64 StartAddress, UINT64 EndAddress,
                             BOOLEAN IsDebuggeePaused,
                             PUINT32 CountOfMatchedCases) {
  CR3_TYPE CurrentProcessCr3;
  UINT64 BaseAddress = 0;
  UINT64 RealPhysicalAddress = 0;
  UINT64 TempValue = (UINT64)NULL;
  UINT64 TempStartAddress = (UINT64)NULL;
  BOOLEAN DoesBaseAddrSaved = FALSE;
  BOOLEAN SearchResult = FALSE;
  *CountOfMatchedCases = 0;
  if (SearchMemRequest->MemoryType == SEARCH_VIRTUAL_MEMORY) {
    TempStartAddress = StartAddress;
    StartAddress = (UINT64)PAGE_ALIGN(StartAddress);
    if (IsDebuggeePaused) {
      CurrentProcessCr3 =
          SwitchToProcessMemoryLayoutByCr3(LayoutGetCurrentProcessCr3());
    } else {
      CurrentProcessCr3 =
          SwitchToProcessMemoryLayout(SearchMemRequest->ProcessId);
    }
    while (StartAddress < EndAddress) {
      TempValue = VirtualAddressToPhysicalAddress((PVOID)StartAddress);
      if (TempValue != 0) {
        if (!DoesBaseAddrSaved) {
          BaseAddress = TempStartAddress;
          DoesBaseAddrSaved = TRUE;
        }
      } else {
        break;
      }
      StartAddress += PAGE_SIZE;
    }
    SwitchToPreviousProcess(CurrentProcessCr3);
    if (DoesBaseAddrSaved && StartAddress > BaseAddress) {
      SearchResult = PerformSearchAddress(
          AddressToSaveResults, SearchMemRequest, BaseAddress, EndAddress,
          IsDebuggeePaused, CountOfMatchedCases);
    } else {
      return FALSE;
    }
  } else if (SearchMemRequest->MemoryType == SEARCH_PHYSICAL_MEMORY) {
    RealPhysicalAddress = SearchMemRequest->Address;
    if (IsDebuggeePaused) {
      SearchMemRequest->Address =
          PhysicalAddressToVirtualAddressOnTargetProcess((PVOID)StartAddress);
      EndAddress =
          PhysicalAddressToVirtualAddressOnTargetProcess((PVOID)EndAddress);
    } else if (SearchMemRequest->ProcessId ==
               HANDLE_TO_UINT32(PsGetCurrentProcessId())) {
      SearchMemRequest->Address = PhysicalAddressToVirtualAddress(StartAddress);
      EndAddress = PhysicalAddressToVirtualAddress(EndAddress);
    } else {
      SearchMemRequest->Address = PhysicalAddressToVirtualAddressByProcessId(
          (PVOID)StartAddress, SearchMemRequest->ProcessId);
      EndAddress = PhysicalAddressToVirtualAddressByProcessId(
          (PVOID)EndAddress, SearchMemRequest->ProcessId);
    }
    SearchMemRequest->MemoryType = SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY;
    SearchResult = PerformSearchAddress(AddressToSaveResults, SearchMemRequest,
                                        SearchMemRequest->Address, EndAddress,
                                        IsDebuggeePaused, CountOfMatchedCases);
    SearchMemRequest->MemoryType = SEARCH_PHYSICAL_MEMORY;
    SearchMemRequest->Address = RealPhysicalAddress;
  }
  return SearchResult;
}

NTSTATUS DebuggerCommandSearchMemory(PDEBUGGER_SEARCH_MEMORY SearchMemRequest) {
  PUINT64 SearchResultsStorage = NULL;
  PUINT64 UsermodeBuffer = NULL;
  UINT64 AddressFrom = 0;
  UINT64 AddressTo = 0;
  UINT64 CurrentValue = 0;
  UINT32 ResultsIndex = 0;
  UINT32 CountOfResults = 0;
  if (SearchMemRequest->ProcessId !=
          HANDLE_TO_UINT32(PsGetCurrentProcessId()) &&
      !CommonIsProcessExist(SearchMemRequest->ProcessId)) {
    return STATUS_INVALID_PARAMETER;
  }
  UsermodeBuffer = (UINT64 *)SearchMemRequest;
  AddressFrom = SearchMemRequest->Address;
  AddressTo = SearchMemRequest->Address + SearchMemRequest->Length;
  SearchResultsStorage = PlatformMemAllocateZeroedNonPagedPool(
      MaximumSearchResults * sizeof(UINT64));
  if (SearchResultsStorage == NULL) {
    return STATUS_INSUFFICIENT_RESOURCES;
  }
  SearchAddressWrapper(SearchResultsStorage, SearchMemRequest, AddressFrom,
                       AddressTo, FALSE, &CountOfResults);
  RtlZeroMemory(SearchMemRequest, MaximumSearchResults * sizeof(UINT64));
  for (size_t i = 0; i < MaximumSearchResults; i++) {
    CurrentValue = SearchResultsStorage[i];
    if (CurrentValue == (UINT64)NULL) {
      break;
    }
    if (CurrentValue >= AddressFrom && CurrentValue <= AddressTo) {
      UsermodeBuffer[ResultsIndex] = CurrentValue;
      ResultsIndex++;
    }
  }
  PlatformMemFreePool(SearchResultsStorage);
  return STATUS_SUCCESS;
}

NTSTATUS DebuggerCommandFlush(
    PDEBUGGER_FLUSH_LOGGING_BUFFERS DebuggerFlushBuffersRequest) {
  DebuggerFlushBuffersRequest->CountOfMessagesThatSetAsReadFromVmxRoot =
      LogMarkAllAsRead(TRUE);
  DebuggerFlushBuffersRequest->CountOfMessagesThatSetAsReadFromVmxNonRoot =
      LogMarkAllAsRead(FALSE);
  DebuggerFlushBuffersRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return STATUS_SUCCESS;
}

NTSTATUS DebuggerCommandSignalExecutionState(
    PDEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL
        DebuggerFinishedExecutionRequest) {
  VmFuncVmxVmcall(DEBUGGER_VMCALL_SIGNAL_DEBUGGER_EXECUTION_FINISHED, 0, 0, 0);
  DebuggerFinishedExecutionRequest->KernelStatus =
      DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return STATUS_SUCCESS;
}

NTSTATUS DebuggerCommandSendMessage(PDEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER
                                        DebuggerSendUsermodeMessageRequest) {
  VmFuncVmxVmcall(DEBUGGER_VMCALL_SEND_MESSAGES_TO_DEBUGGER,
                  (UINT64)DebuggerSendUsermodeMessageRequest +
                      (SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER),
                  DebuggerSendUsermodeMessageRequest->Length, 0);
  DebuggerSendUsermodeMessageRequest->KernelStatus =
      DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return STATUS_SUCCESS;
}

NTSTATUS DebuggerCommandSendGeneralBufferToDebugger(
    PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
        DebuggeeBufferRequest) {
  VmFuncVmxVmcall(DEBUGGER_VMCALL_SEND_GENERAL_BUFFER_TO_DEBUGGER,
                  (UINT64)DebuggeeBufferRequest, 0, 0);
  DebuggeeBufferRequest->KernelResult = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return STATUS_SUCCESS;
}

NTSTATUS DebuggerCommandReservePreallocatedPools(
    PDEBUGGER_PREALLOC_COMMAND PreallocRequest) {
  switch (PreallocRequest->Type) {
    case DEBUGGER_PREALLOC_COMMAND_TYPE_THREAD_INTERCEPTION:
      PoolManagerRequestAllocation(sizeof(USERMODE_DEBUGGING_THREAD_HOLDER),
                                   PreallocRequest->Count,
                                   PROCESS_THREAD_HOLDER);
      break;
    case DEBUGGER_PREALLOC_COMMAND_TYPE_MONITOR:
      ConfigureEptHookAllocateExtraHookingPagesForMemoryMonitorsAndExecEptHooks(
          PreallocRequest->Count);
      break;
    case DEBUGGER_PREALLOC_COMMAND_TYPE_EPTHOOK:
      ConfigureEptHookAllocateExtraHookingPagesForMemoryMonitorsAndExecEptHooks(
          PreallocRequest->Count);
      break;
    case DEBUGGER_PREALLOC_COMMAND_TYPE_EPTHOOK2:
      ConfigureEptHookReservePreallocatedPoolsForEptHooks(
          PreallocRequest->Count);
      break;
    case DEBUGGER_PREALLOC_COMMAND_TYPE_REGULAR_EVENT:
      PoolManagerRequestAllocation(REGULAR_INSTANT_EVENT_CONDITIONAL_BUFFER,
                                   PreallocRequest->Count,
                                   INSTANT_REGULAR_EVENT_BUFFER);
      PoolManagerRequestAllocation(REGULAR_INSTANT_EVENT_ACTION_BUFFER,
                                   PreallocRequest->Count,
                                   INSTANT_REGULAR_EVENT_ACTION_BUFFER);
      break;
    case DEBUGGER_PREALLOC_COMMAND_TYPE_BIG_EVENT:
      PoolManagerRequestAllocation(BIG_INSTANT_EVENT_CONDITIONAL_BUFFER,
                                   PreallocRequest->Count,
                                   INSTANT_BIG_EVENT_BUFFER);
      PoolManagerRequestAllocation(BIG_INSTANT_EVENT_ACTION_BUFFER,
                                   PreallocRequest->Count,
                                   INSTANT_BIG_EVENT_ACTION_BUFFER);
      break;
    case DEBUGGER_PREALLOC_COMMAND_TYPE_REGULAR_SAFE_BUFFER:
      PoolManagerRequestAllocation(REGULAR_INSTANT_EVENT_REQUESTED_SAFE_BUFFER,
                                   PreallocRequest->Count,
                                   INSTANT_REGULAR_SAFE_BUFFER_FOR_EVENTS);
      break;
    case DEBUGGER_PREALLOC_COMMAND_TYPE_BIG_SAFE_BUFFER:
      PoolManagerRequestAllocation(BIG_INSTANT_EVENT_REQUESTED_SAFE_BUFFER,
                                   PreallocRequest->Count,
                                   INSTANT_BIG_SAFE_BUFFER_FOR_EVENTS);
      break;
    default:
      PreallocRequest->KernelStatus =
          DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE;
      return STATUS_UNSUCCESSFUL;
  }
  PoolManagerCheckAndPerformAllocationAndDeallocation();
  PreallocRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return STATUS_SUCCESS;
}

NTSTATUS DebuggerCommandPreactivateFunctionality(
    PDEBUGGER_PREACTIVATE_COMMAND PreactivateRequest) {
  switch (PreactivateRequest->Type) {
    case DEBUGGER_PREACTIVATE_COMMAND_TYPE_MODE:
      ConfigureInitializeExecTrapOnAllProcessors();
      break;
    default:
      PreactivateRequest->KernelStatus =
          DEBUGGER_ERROR_COULD_NOT_FIND_PREACTIVATION_TYPE;
      return STATUS_UNSUCCESSFUL;
  }
  PreactivateRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return STATUS_SUCCESS;
}

BOOLEAN DebuggerCommandBringPagein(PDEBUGGER_PAGE_IN_REQUEST PageinRequest) {
  LogInfo("Page-request is received!");
  PageinRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return TRUE;
}
