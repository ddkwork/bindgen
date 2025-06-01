
#include "pch.h"

_Must_inspect_result_ _Success_(return == TRUE)

    static EPT_HOOKED_PAGE_DETAIL *EptHookFindByPhysAddress(
        _In_ UINT64 PhysicalBaseAddress) {
  LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL,
                     PageHookList, CurrEntity) {
    if (CurrEntity->PhysicalBaseAddress == PhysicalBaseAddress) {
      return CurrEntity;
    }
  }
  return NULL;
}

static UINT64 EptHookCalcBreakpointOffset(
    _In_ PVOID TargetAddress, _In_ EPT_HOOKED_PAGE_DETAIL *HookedEntry) {
  UINT64 TargetAddressInFakePageContent;
  UINT64 PageOffset;
  TargetAddressInFakePageContent = (UINT64)&HookedEntry->FakePageContents;
  TargetAddressInFakePageContent =
      (UINT64)PAGE_ALIGN(TargetAddressInFakePageContent);
  PageOffset = (UINT64)PAGE_OFFSET(TargetAddress);
  TargetAddressInFakePageContent = TargetAddressInFakePageContent + PageOffset;
  return TargetAddressInFakePageContent;
}

VOID EptHookReservePreallocatedPoolsForEptHooks(UINT32 Count) {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  PoolManagerRequestAllocation(sizeof(VMM_EPT_DYNAMIC_SPLIT),
                               Count * ProcessorsCount,
                               SPLIT_2MB_PAGING_TO_4KB_PAGE);
  PoolManagerRequestAllocation(sizeof(EPT_HOOKED_PAGE_DETAIL), Count,
                               TRACKING_HOOKED_PAGES);
  PoolManagerRequestAllocation(MAX_EXEC_TRAMPOLINE_SIZE, Count,
                               EXEC_TRAMPOLINE);
  PoolManagerRequestAllocation(sizeof(HIDDEN_HOOKS_DETOUR_DETAILS), Count,
                               DETOUR_HOOK_DETAILS);
}

VOID EptHookAllocateExtraHookingPagesForMemoryMonitorsAndExecEptHooks(
    UINT32 Count) {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  PoolManagerRequestAllocation(sizeof(VMM_EPT_DYNAMIC_SPLIT),
                               Count * ProcessorsCount,
                               SPLIT_2MB_PAGING_TO_4KB_PAGE);
  PoolManagerRequestAllocation(sizeof(EPT_HOOKED_PAGE_DETAIL), Count,
                               TRACKING_HOOKED_PAGES);
}

static BOOLEAN EptHookCreateHookPage(_Inout_ VIRTUAL_MACHINE_STATE *VCpu,
                                     _In_ PVOID TargetAddress,
                                     _In_ CR3_TYPE ProcessCr3) {
  ULONG ProcessorsCount;
  EPT_PML1_ENTRY ChangedEntry;
  SIZE_T PhysicalBaseAddress;
  PVOID VirtualTarget;
  PVOID TargetBuffer;
  UINT64 TargetAddressInFakePageContent;
  PEPT_PML1_ENTRY TargetPage;
  PEPT_HOOKED_PAGE_DETAIL HookedPage;
  CR3_TYPE Cr3OfCurrentProcess;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  VirtualTarget = PAGE_ALIGN(TargetAddress);
  PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(
      VirtualTarget, ProcessCr3);
  if (!PhysicalBaseAddress) {
    VmmCallbackSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
    return FALSE;
  }
  HookedPage = (EPT_HOOKED_PAGE_DETAIL *)PoolManagerRequestPool(
      TRACKING_HOOKED_PAGES, TRUE, sizeof(EPT_HOOKED_PAGE_DETAIL));
  if (!HookedPage) {
    VmmCallbackSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
    return FALSE;
  }
  HookedPage->IsHiddenBreakpoint = TRUE;
  HookedPage->VirtualAddress = (UINT64)TargetAddress;
  HookedPage->PhysicalBaseAddress = PhysicalBaseAddress;
  HookedPage->PhysicalBaseAddressOfFakePageContents =
      (SIZE_T)VirtualAddressToPhysicalAddress(
          &HookedPage->FakePageContents[0]) /
      PAGE_SIZE;
  HookedPage->IsExecutionHook = TRUE;
  HookedPage->BreakpointAddresses[0] = (UINT64)TargetAddress;
  HookedPage->CountOfBreakpoints = 1;
  TargetAddressInFakePageContent =
      EptHookCalcBreakpointOffset(TargetAddress, HookedPage);
  Cr3OfCurrentProcess = SwitchToProcessMemoryLayoutByCr3(ProcessCr3);
  MemoryMapperReadMemorySafe((UINT64)VirtualTarget,
                             &HookedPage->FakePageContents, PAGE_SIZE);
  *(BYTE *)TargetAddressInFakePageContent = 0xcc;
  SwitchToPreviousProcess(Cr3OfCurrentProcess);
  for (size_t i = 0; i < ProcessorsCount; i++) {
    TargetBuffer = (PVOID)PoolManagerRequestPool(
        SPLIT_2MB_PAGING_TO_4KB_PAGE, TRUE, sizeof(VMM_EPT_DYNAMIC_SPLIT));
    if (!TargetBuffer) {
      PoolManagerFreePool((UINT64)HookedPage);
      VmmCallbackSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
      return FALSE;
    }
    if (!EptSplitLargePage(g_GuestState[i].EptPageTable, TargetBuffer,
                           PhysicalBaseAddress)) {
      PoolManagerFreePool((UINT64)HookedPage);
      PoolManagerFreePool((UINT64)TargetBuffer);
      LogDebugInfo("Err, could not split page for the address : 0x%llx",
                   PhysicalBaseAddress);
      VmmCallbackSetLastError(
          DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES);
      return FALSE;
    }
    TargetPage =
        EptGetPml1Entry(g_GuestState[i].EptPageTable, PhysicalBaseAddress);
    if (!TargetPage) {
      PoolManagerFreePool((UINT64)HookedPage);
      PoolManagerFreePool((UINT64)TargetBuffer);
      VmmCallbackSetLastError(
          DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS);
      return FALSE;
    }
    HookedPage->OriginalEntry = *TargetPage;
    ChangedEntry = *TargetPage;
    ChangedEntry.ReadAccess = 0;
    ChangedEntry.WriteAccess = 0;
    ChangedEntry.ExecuteAccess = 1;
    ChangedEntry.PageFrameNumber =
        HookedPage->PhysicalBaseAddressOfFakePageContents;
    if (i == 0) {
      HookedPage->ChangedEntry = ChangedEntry;
      InsertHeadList(&g_EptState->HookedPagesList, &(HookedPage->PageHookList));
    }
    TargetPage->AsUInt = ChangedEntry.AsUInt;
    if (VCpu->CoreId == i && g_GuestState[i].HasLaunched) {
      EptInveptSingleContext(VCpu->EptPointer.AsUInt);
    }
  }
  return TRUE;
}

static BOOLEAN EptHookUpdateHookPage(
    _In_ PVOID TargetAddress, _Inout_ EPT_HOOKED_PAGE_DETAIL *HookedEntry) {
  UINT64 TargetAddressInFakePageContent;
  BYTE OriginalByte;
  if (HookedEntry == NULL) return FALSE;
  if (HookedEntry->CountOfBreakpoints >= MaximumHiddenBreakpointsOnPage) {
    VmmCallbackSetLastError(
        DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT);
    return FALSE;
  }
  TargetAddressInFakePageContent =
      EptHookCalcBreakpointOffset(TargetAddress, HookedEntry);
  OriginalByte = *(BYTE *)TargetAddressInFakePageContent;
  HookedEntry->BreakpointAddresses[HookedEntry->CountOfBreakpoints] =
      (UINT64)TargetAddress;
  HookedEntry
      ->PreviousBytesOnBreakpointAddresses[HookedEntry->CountOfBreakpoints] =
      OriginalByte;
  HookedEntry->CountOfBreakpoints = HookedEntry->CountOfBreakpoints + 1;
  *(BYTE *)TargetAddressInFakePageContent = 0xcc;
  return TRUE;
}

PVOID ExAllocatePoolWithTagHook(POOL_TYPE PoolType, SIZE_T NumberOfBytes,
                                ULONG Tag) {
  LogInfo(
      "ExAllocatePoolWithTag Called with : Tag = 0x%x, Number Of Bytes = 0x%x, "
      "Pool Type = 0x%x ",
      Tag, NumberOfBytes, PoolType);
  return ExAllocatePoolWithTagOrig(PoolType, NumberOfBytes, Tag);
}

BOOLEAN EptHookPerformPageHook(VIRTUAL_MACHINE_STATE *VCpu, PVOID TargetAddress,
                               CR3_TYPE ProcessCr3) {
  SIZE_T PhysicalBaseAddress;
  PVOID VirtualTarget;
  EPT_HOOKED_PAGE_DETAIL *HookedEntry = NULL;
  VirtualTarget = PAGE_ALIGN(TargetAddress);
  PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(
      VirtualTarget, ProcessCr3);
  if (!PhysicalBaseAddress) {
    VmmCallbackSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
    return FALSE;
  }
  HookedEntry = EptHookFindByPhysAddress(PhysicalBaseAddress);
  if (HookedEntry != NULL) {
    return EptHookUpdateHookPage(TargetAddress, HookedEntry);
  } else {
    return EptHookCreateHookPage(VCpu, TargetAddress, ProcessCr3);
  }
}

BOOLEAN EptHookPerformHook(PVOID TargetAddress, UINT32 ProcessId,
                           BOOLEAN ApplyDirectlyFromVmxRoot) {
  if (ApplyDirectlyFromVmxRoot) {
    DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
    DirectVmcallOptions.OptionalParam1 = (UINT64)TargetAddress;
    DirectVmcallOptions.OptionalParam2 = LayoutGetCurrentProcessCr3().Flags;
    if (DirectVmcallSetHiddenBreakpointHook(KeGetCurrentProcessorNumberEx(NULL),
                                            &DirectVmcallOptions) ==
        STATUS_SUCCESS) {
      LogDebugInfo("Hidden breakpoint hook applied from VMX Root Mode");
      return TRUE;
    }
  } else {
    BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores();
    if (AsmVmxVmcall(VMCALL_SET_HIDDEN_CC_BREAKPOINT, (UINT64)TargetAddress,
                     LayoutGetCr3ByProcessId(ProcessId).Flags,
                     (UINT64)NULL64_ZERO) == STATUS_SUCCESS) {
      LogDebugInfo("Hidden breakpoint hook applied from VMX Root Mode");
      BroadcastNotifyAllToInvalidateEptAllCores();
      return TRUE;
    }
  }
  return FALSE;
}

BOOLEAN EptHook(PVOID TargetAddress, UINT32 ProcessId) {
  if (VmxGetCurrentExecutionMode() == TRUE) {
    return FALSE;
  }
  return EptHookPerformHook(TargetAddress, ProcessId, FALSE);
}

BOOLEAN EptHookFromVmxRoot(PVOID TargetAddress) {
  if (VmxGetCurrentExecutionMode() == FALSE) {
    return FALSE;
  }
  return EptHookPerformHook(TargetAddress, NULL_ZERO, TRUE);
}

BOOLEAN EptHookRestoreSingleHookToOriginalEntry(VIRTUAL_MACHINE_STATE *VCpu,
                                                SIZE_T PhysicalAddress,
                                                UINT64 OriginalEntry) {
  PEPT_PML1_ENTRY TargetPage;
  if (VmxGetCurrentExecutionMode() == FALSE) {
    return FALSE;
  }
  TargetPage = EptGetPml1Entry(VCpu->EptPageTable, PhysicalAddress);
  if (TargetPage != NULL) {
    TargetPage->AsUInt = OriginalEntry;
    EptInveptSingleContext(VCpu->EptPointer.AsUInt);
    return TRUE;
  }
  return FALSE;
}

VOID EptHookRestoreAllHooksToOriginalEntry(VIRTUAL_MACHINE_STATE *VCpu) {
  PEPT_PML1_ENTRY TargetPage;
  if (VmxGetCurrentExecutionMode() == FALSE) {
    return;
  }
  LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL,
                     PageHookList, HookedEntry) {
    TargetPage =
        EptGetPml1Entry(VCpu->EptPageTable, HookedEntry->PhysicalBaseAddress);
    TargetPage->AsUInt = HookedEntry->OriginalEntry.AsUInt;
  }
  EptInveptSingleContext(VCpu->EptPointer.AsUInt);
}

VOID EptHookWriteAbsoluteJump(PCHAR TargetBuffer, SIZE_T TargetAddress) {
  TargetBuffer[0] = 0xe8;
  TargetBuffer[1] = 0x00;
  TargetBuffer[2] = 0x00;
  TargetBuffer[3] = 0x00;
  TargetBuffer[4] = 0x00;
  TargetBuffer[5] = 0x68;
  *((PUINT32)&TargetBuffer[6]) = (UINT32)TargetAddress;
  TargetBuffer[10] = 0xC7;
  TargetBuffer[11] = 0x44;
  TargetBuffer[12] = 0x24;
  TargetBuffer[13] = 0x04;
  *((PUINT32)&TargetBuffer[14]) = (UINT32)(TargetAddress >> 32);
  TargetBuffer[18] = 0xC3;
}

VOID EptHookWriteAbsoluteJump2(PCHAR TargetBuffer, SIZE_T TargetAddress) {
  TargetBuffer[0] = 0x68;
  *((PUINT32)&TargetBuffer[1]) = (UINT32)TargetAddress;
  TargetBuffer[5] = 0xC7;
  TargetBuffer[6] = 0x44;
  TargetBuffer[7] = 0x24;
  TargetBuffer[8] = 0x04;
  *((PUINT32)&TargetBuffer[9]) = (UINT32)(TargetAddress >> 32);
  TargetBuffer[13] = 0xC3;
}

BOOLEAN EptHookInstructionMemory(PEPT_HOOKED_PAGE_DETAIL Hook,
                                 CR3_TYPE ProcessCr3, PVOID TargetFunction,
                                 PVOID TargetFunctionInSafeMemory,
                                 PVOID HookFunction) {
  PHIDDEN_HOOKS_DETOUR_DETAILS DetourHookDetails;
  SIZE_T SizeOfHookedInstructions;
  SIZE_T OffsetIntoPage;
  CR3_TYPE Cr3OfCurrentProcess;
  OffsetIntoPage = ADDRMASK_EPT_PML1_OFFSET((SIZE_T)TargetFunction);
  if ((OffsetIntoPage + 19) > PAGE_SIZE - 1) {
    LogError("Err, function extends past a page boundary");
    return FALSE;
  }
  for (SizeOfHookedInstructions = 0; SizeOfHookedInstructions < 19;
       SizeOfHookedInstructions +=
       DisassemblerLengthDisassembleEngineInVmxRootOnTargetProcess(
           (PVOID)((UINT64)TargetFunctionInSafeMemory +
                   SizeOfHookedInstructions),
           FALSE)) {
  }
  Hook->Trampoline = (CHAR *)PoolManagerRequestPool(EXEC_TRAMPOLINE, TRUE,
                                                    MAX_EXEC_TRAMPOLINE_SIZE);
  if (!Hook->Trampoline) {
    LogError("Err, could not allocate trampoline function buffer");
    return FALSE;
  }
  Cr3OfCurrentProcess = SwitchToProcessMemoryLayoutByCr3(ProcessCr3);
  MemoryMapperReadMemorySafe((UINT64)TargetFunction, Hook->Trampoline,
                             SizeOfHookedInstructions);
  SwitchToPreviousProcess(Cr3OfCurrentProcess);
  EptHookWriteAbsoluteJump2(&Hook->Trampoline[SizeOfHookedInstructions],
                            (SIZE_T)TargetFunction + SizeOfHookedInstructions);
  DetourHookDetails = (HIDDEN_HOOKS_DETOUR_DETAILS *)PoolManagerRequestPool(
      DETOUR_HOOK_DETAILS, TRUE, sizeof(HIDDEN_HOOKS_DETOUR_DETAILS));
  DetourHookDetails->HookedFunctionAddress = TargetFunction;
  DetourHookDetails->ReturnAddress = Hook->Trampoline;
  Hook->AddressOfEptHook2sDetourListEntry = (UINT64)DetourHookDetails;
  InsertHeadList(&g_EptHook2sDetourListHead,
                 &(DetourHookDetails->OtherHooksList));
  EptHookWriteAbsoluteJump(&Hook->FakePageContents[OffsetIntoPage],
                           (SIZE_T)HookFunction);
  return TRUE;
}

BOOLEAN EptHookPerformPageHookMonitorAndInlineHook(VIRTUAL_MACHINE_STATE *VCpu,
                                                   PVOID HookingDetails,
                                                   CR3_TYPE ProcessCr3,
                                                   UINT32 PageHookMask) {
  ULONG ProcessorsCount;
  EPT_PML1_ENTRY ChangedEntry;
  SIZE_T PhysicalBaseAddress;
  PVOID AlignedTargetVaOrPa;
  PVOID TargetBuffer;
  PVOID TargetAddress;
  PVOID HookFunction;
  UINT64 TargetAddressInSafeMemory;
  PEPT_PML1_ENTRY TargetPage;
  PEPT_HOOKED_PAGE_DETAIL HookedPage;
  CR3_TYPE Cr3OfCurrentProcess;
  PLIST_ENTRY TempList = 0;
  PEPT_HOOKED_PAGE_DETAIL HookedEntry = NULL;
  BOOLEAN UnsetExecute = FALSE;
  BOOLEAN UnsetRead = FALSE;
  BOOLEAN UnsetWrite = FALSE;
  BOOLEAN EptHiddenHook = FALSE;
  UnsetRead = (PageHookMask & PAGE_ATTRIB_READ) ? TRUE : FALSE;
  UnsetWrite = (PageHookMask & PAGE_ATTRIB_WRITE) ? TRUE : FALSE;
  UnsetExecute = (PageHookMask & PAGE_ATTRIB_EXEC) ? TRUE : FALSE;
  EptHiddenHook = (PageHookMask & PAGE_ATTRIB_EXEC_HIDDEN_HOOK) ? TRUE : FALSE;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (EptHiddenHook) {
    TargetAddress = ((EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2 *)HookingDetails)
                        ->TargetAddress;
  } else {
    TargetAddress =
        (PVOID)((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)HookingDetails)
            ->StartAddress;
  }
  AlignedTargetVaOrPa = PAGE_ALIGN(TargetAddress);
  if (!EptHiddenHook &&
      ((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)HookingDetails)
              ->MemoryType == DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS) {
    PhysicalBaseAddress = (SIZE_T)AlignedTargetVaOrPa;
  } else {
    PhysicalBaseAddress = (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(
        AlignedTargetVaOrPa, ProcessCr3);
  }
  if (!PhysicalBaseAddress) {
    VmmCallbackSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
    return FALSE;
  }
  TempList = &g_EptState->HookedPagesList;
  while (&g_EptState->HookedPagesList != TempList->Flink) {
    TempList = TempList->Flink;
    HookedEntry =
        CONTAINING_RECORD(TempList, EPT_HOOKED_PAGE_DETAIL, PageHookList);
    if (HookedEntry->PhysicalBaseAddress == PhysicalBaseAddress) {
      VmmCallbackSetLastError(
          DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE);
      return FALSE;
    }
  }
  HookedPage = (EPT_HOOKED_PAGE_DETAIL *)PoolManagerRequestPool(
      TRACKING_HOOKED_PAGES, TRUE, sizeof(EPT_HOOKED_PAGE_DETAIL));
  if (!HookedPage) {
    VmmCallbackSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
    return FALSE;
  }
  HookedPage->VirtualAddress = (UINT64)TargetAddress;
  HookedPage->PhysicalBaseAddress = PhysicalBaseAddress;
  if (!EptHiddenHook) {
    HookedPage->HookingTag =
        ((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)HookingDetails)->Tag;
    if (((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)HookingDetails)
            ->MemoryType == DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS) {
      HookedPage->StartOfTargetPhysicalAddress =
          (SIZE_T)(((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)
                        HookingDetails)
                       ->StartAddress);
    } else {
      HookedPage->StartOfTargetPhysicalAddress =
          (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(
              (PVOID)(((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)
                           HookingDetails)
                          ->StartAddress),
              ProcessCr3);
    }
    if (!HookedPage->StartOfTargetPhysicalAddress) {
      PoolManagerFreePool((UINT64)HookedPage);
      VmmCallbackSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
      return FALSE;
    }
    if (((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)HookingDetails)
            ->MemoryType == DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS) {
      HookedPage->EndOfTargetPhysicalAddress =
          (SIZE_T)(((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)
                        HookingDetails)
                       ->EndAddress);
    } else {
      HookedPage->EndOfTargetPhysicalAddress =
          (SIZE_T)VirtualAddressToPhysicalAddressByProcessCr3(
              (PVOID)(((EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *)
                           HookingDetails)
                          ->EndAddress),
              ProcessCr3);
    }
    if (!HookedPage->EndOfTargetPhysicalAddress) {
      PoolManagerFreePool((UINT64)HookedPage);
      VmmCallbackSetLastError(DEBUGGER_ERROR_INVALID_ADDRESS);
      return FALSE;
    }
  }
  HookedPage->PhysicalBaseAddressOfFakePageContents =
      (SIZE_T)VirtualAddressToPhysicalAddress(
          &HookedPage->FakePageContents[0]) /
      PAGE_SIZE;
  if (EptHiddenHook) {
    HookedPage->IsExecutionHook = TRUE;
    Cr3OfCurrentProcess = SwitchToProcessMemoryLayoutByCr3(ProcessCr3);
    MemoryMapperReadMemorySafe((UINT64)AlignedTargetVaOrPa,
                               &HookedPage->FakePageContents, PAGE_SIZE);
    SwitchToPreviousProcess(Cr3OfCurrentProcess);
    TargetAddressInSafeMemory =
        EptHookCalcBreakpointOffset(TargetAddress, HookedPage);
    if (((EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2 *)HookingDetails)
            ->HookFunction == NULL) {
      HookFunction = (PVOID)AsmGeneralDetourHook;
    } else {
      HookFunction = ((EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2 *)HookingDetails)
                         ->HookFunction;
    }
    if (!EptHookInstructionMemory(HookedPage, ProcessCr3, TargetAddress,
                                  (PVOID)TargetAddressInSafeMemory,
                                  HookFunction)) {
      PoolManagerFreePool((UINT64)HookedPage);
      VmmCallbackSetLastError(DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK);
      return FALSE;
    }
  }
  for (size_t i = 0; i < ProcessorsCount; i++) {
    TargetBuffer = (PVOID)PoolManagerRequestPool(
        SPLIT_2MB_PAGING_TO_4KB_PAGE, TRUE, sizeof(VMM_EPT_DYNAMIC_SPLIT));
    if (!TargetBuffer) {
      PoolManagerFreePool((UINT64)HookedPage);
      VmmCallbackSetLastError(DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY);
      return FALSE;
    }
    if (!EptSplitLargePage(g_GuestState[i].EptPageTable, TargetBuffer,
                           PhysicalBaseAddress)) {
      PoolManagerFreePool((UINT64)HookedPage);
      PoolManagerFreePool((UINT64)TargetBuffer);
      LogDebugInfo("Err, could not split page for the address : 0x%llx",
                   PhysicalBaseAddress);
      VmmCallbackSetLastError(
          DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES);
      return FALSE;
    }
    TargetPage =
        EptGetPml1Entry(g_GuestState[i].EptPageTable, PhysicalBaseAddress);
    if (!TargetPage) {
      PoolManagerFreePool((UINT64)HookedPage);
      PoolManagerFreePool((UINT64)TargetBuffer);
      VmmCallbackSetLastError(
          DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS);
      return FALSE;
    }
    HookedPage->OriginalEntry = *TargetPage;
    ChangedEntry = *TargetPage;
    if (UnsetRead)
      ChangedEntry.ReadAccess = 0;
    else
      ChangedEntry.ReadAccess = 1;
    if (UnsetWrite)
      ChangedEntry.WriteAccess = 0;
    else
      ChangedEntry.WriteAccess = 1;
    if (UnsetExecute)
      ChangedEntry.ExecuteAccess = 0;
    else
      ChangedEntry.ExecuteAccess = 1;
    if (EptHiddenHook) {
      ChangedEntry.ReadAccess = 0;
      ChangedEntry.WriteAccess = 0;
      ChangedEntry.ExecuteAccess = 1;
      ChangedEntry.PageFrameNumber =
          HookedPage->PhysicalBaseAddressOfFakePageContents;
    }
    if (i == 0) {
      HookedPage->ChangedEntry = ChangedEntry;
      InsertHeadList(&g_EptState->HookedPagesList, &(HookedPage->PageHookList));
    }
    TargetPage->AsUInt = ChangedEntry.AsUInt;
    if (VCpu->CoreId == i && g_GuestState[i].HasLaunched) {
      EptInveptSingleContext(VCpu->EptPointer.AsUInt);
    }
  }
  return TRUE;
}

BOOLEAN EptHookPerformMemoryOrInlineHook(
    VIRTUAL_MACHINE_STATE *VCpu,
    EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2 *EptHook2AddressDetails,
    EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *MemoryAddressDetails,
    UINT32 ProcessId, BOOLEAN EptHiddenHook2,
    BOOLEAN ApplyDirectlyFromVmxRoot) {
  UINT32 PageHookMask = 0;
  PVOID HookDetailsToVmcall = NULL;
  if (MemoryAddressDetails != NULL) {
    HookDetailsToVmcall = MemoryAddressDetails;
    if (MemoryAddressDetails->SetHookForExec &&
        !g_CompatibilityCheck.ExecuteOnlySupport) {
      return FALSE;
    }
    if (!MemoryAddressDetails->SetHookForWrite &&
        MemoryAddressDetails->SetHookForRead) {
      return FALSE;
    }
    if (MemoryAddressDetails->SetHookForRead) {
      PageHookMask |= PAGE_ATTRIB_READ;
    }
    if (MemoryAddressDetails->SetHookForWrite) {
      PageHookMask |= PAGE_ATTRIB_WRITE;
    }
    if (MemoryAddressDetails->SetHookForExec) {
      PageHookMask |= PAGE_ATTRIB_EXEC;
    }
  } else if (EptHook2AddressDetails != NULL) {
    HookDetailsToVmcall = EptHook2AddressDetails;
    if (!g_IsEptHook2sDetourListInitialized) {
      g_IsEptHook2sDetourListInitialized = TRUE;
      InitializeListHead(&g_EptHook2sDetourListHead);
    }
    if (EptHiddenHook2) {
      PageHookMask |= PAGE_ATTRIB_EXEC_HIDDEN_HOOK;
    }
  } else {
    return FALSE;
  }
  if (PageHookMask == 0) {
    return FALSE;
  }
  if (ApplyDirectlyFromVmxRoot) {
    DIRECT_VMCALL_PARAMETERS DirectVmcallOptions = {0};
    DirectVmcallOptions.OptionalParam1 = (UINT64)HookDetailsToVmcall;
    DirectVmcallOptions.OptionalParam2 = PageHookMask;
    DirectVmcallOptions.OptionalParam3 = LayoutGetCurrentProcessCr3().Flags;
    if (DirectVmcallPerformVmcall(VCpu->CoreId, VMCALL_CHANGE_PAGE_ATTRIB,
                                  &DirectVmcallOptions) == STATUS_SUCCESS) {
      return TRUE;
    } else {
      return FALSE;
    }
  } else {
    if (VmxGetCurrentLaunchState()) {
      if (AsmVmxVmcall(VMCALL_CHANGE_PAGE_ATTRIB, (UINT64)HookDetailsToVmcall,
                       PageHookMask,
                       LayoutGetCr3ByProcessId(ProcessId).Flags) ==
          STATUS_SUCCESS) {
        if (VmxGetCurrentExecutionMode() == FALSE) {
          BroadcastNotifyAllToInvalidateEptAllCores();
        } else {
          LogInfo(
              "Err, unable to notify all cores to invalidate their TLB "
              "caches as you called hook on vmx-root mode, however, the "
              "hook is still works");
        }
        return TRUE;
      } else {
        return FALSE;
      }
    } else {
      if (EptHookPerformPageHookMonitorAndInlineHook(
              VCpu, HookDetailsToVmcall, LayoutGetCr3ByProcessId(ProcessId),
              PageHookMask) == TRUE) {
        LogWarning("Hook applied (VM has not launched)");
        return TRUE;
      }
    }
  }
  LogWarning("Err, hook was not applied");
  return FALSE;
}

BOOLEAN EptHookInlineHook(VIRTUAL_MACHINE_STATE *VCpu, PVOID TargetAddress,
                          PVOID HookFunction, UINT32 ProcessId) {
  EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2 HookingDetail = {0};
  if (VmxGetCurrentExecutionMode() == TRUE) {
    return FALSE;
  }
  HookingDetail.TargetAddress = TargetAddress;
  HookingDetail.HookFunction = HookFunction;
  return EptHookPerformMemoryOrInlineHook(VCpu, &HookingDetail, NULL, ProcessId,
                                          TRUE, FALSE);
}

BOOLEAN EptHookMonitorHook(
    VIRTUAL_MACHINE_STATE *VCpu,
    EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *HookingDetails,
    UINT32 ProcessId) {
  if (VmxGetCurrentExecutionMode() == TRUE) {
    return FALSE;
  }
  return EptHookPerformMemoryOrInlineHook(VCpu, NULL, HookingDetails, ProcessId,
                                          FALSE, FALSE);
}

BOOLEAN EptHookInlineHookFromVmxRoot(VIRTUAL_MACHINE_STATE *VCpu,
                                     PVOID TargetAddress, PVOID HookFunction) {
  EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2 HookingDetail = {0};
  if (VmxGetCurrentExecutionMode() == FALSE) {
    return FALSE;
  }
  HookingDetail.TargetAddress = TargetAddress;
  HookingDetail.HookFunction = HookFunction;
  return EptHookPerformMemoryOrInlineHook(VCpu, &HookingDetail, NULL, NULL_ZERO,
                                          TRUE, TRUE);
}

BOOLEAN EptHookMonitorFromVmxRoot(
    VIRTUAL_MACHINE_STATE *VCpu,
    EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *MemoryAddressDetails) {
  if (VmxGetCurrentExecutionMode() == FALSE) {
    return FALSE;
  }
  return EptHookPerformMemoryOrInlineHook(VCpu, NULL, MemoryAddressDetails,
                                          NULL_ZERO, FALSE, TRUE);
}

BOOLEAN EptHookHandleHookedPage(
    VIRTUAL_MACHINE_STATE *VCpu, EPT_HOOKED_PAGE_DETAIL *HookedEntryDetails,
    VMX_EXIT_QUALIFICATION_EPT_VIOLATION ViolationQualification,
    SIZE_T PhysicalAddress, EPT_HOOKS_CONTEXT *LastContext,
    BOOLEAN *IgnoreReadOrWriteOrExec, BOOLEAN *IsExecViolation) {
  UINT64 ExactAccessedVirtualAddress;
  UINT64 AlignedVirtualAddress;
  UINT64 AlignedPhysicalAddress;
  BOOLEAN IsTriggeringPostEventAllowed = FALSE;
  AlignedVirtualAddress =
      (UINT64)PAGE_ALIGN(HookedEntryDetails->VirtualAddress);
  AlignedPhysicalAddress = (UINT64)PAGE_ALIGN(PhysicalAddress);
  ExactAccessedVirtualAddress =
      AlignedVirtualAddress + PhysicalAddress - AlignedPhysicalAddress;
  LastContext->HookingTag = HookedEntryDetails->HookingTag;
  LastContext->PhysicalAddress = PhysicalAddress;
  LastContext->VirtualAddress = ExactAccessedVirtualAddress;
  if (!ViolationQualification.EptReadable &&
      ViolationQualification.ReadAccess) {
    HookedEntryDetails->LastViolation = EPT_HOOKED_LAST_VIOLATION_READ;
    *IgnoreReadOrWriteOrExec =
        DispatchEventHiddenHookPageReadWriteExecuteReadPreEvent(
            VCpu, LastContext, &IsTriggeringPostEventAllowed);
    *IsExecViolation = FALSE;
  } else if (!ViolationQualification.EptWriteable &&
             ViolationQualification.WriteAccess) {
    HookedEntryDetails->LastViolation = EPT_HOOKED_LAST_VIOLATION_WRITE;
    *IgnoreReadOrWriteOrExec =
        DispatchEventHiddenHookPageReadWriteExecuteWritePreEvent(
            VCpu, LastContext, &IsTriggeringPostEventAllowed);
    *IsExecViolation = FALSE;
  } else if (!ViolationQualification.EptExecutable &&
             ViolationQualification.ExecuteAccess) {
    HookedEntryDetails->LastViolation = EPT_HOOKED_LAST_VIOLATION_EXEC;
    *IgnoreReadOrWriteOrExec =
        DispatchEventHiddenHookPageReadWriteExecuteExecutePreEvent(
            VCpu, LastContext, &IsTriggeringPostEventAllowed);
    *IsExecViolation = TRUE;
  } else {
    HookedEntryDetails->IsPostEventTriggerAllowed = FALSE;
    return FALSE;
  }
  if (*IgnoreReadOrWriteOrExec == FALSE) {
    HookedEntryDetails->IsPostEventTriggerAllowed =
        IsTriggeringPostEventAllowed;
  } else {
    HookedEntryDetails->IsPostEventTriggerAllowed = FALSE;
  }
  return TRUE;
}

BOOLEAN EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(UINT64 Address) {
  LIST_FOR_EACH_LINK(g_EptHook2sDetourListHead, HIDDEN_HOOKS_DETOUR_DETAILS,
                     OtherHooksList, CurrentHookedDetails) {
    if (CurrentHookedDetails->HookedFunctionAddress == (PVOID)Address) {
      RemoveEntryList(&CurrentHookedDetails->OtherHooksList);
      if (!PoolManagerFreePool((UINT64)CurrentHookedDetails)) {
        LogError(
            "Err, something goes wrong, the pool not found in the list of "
            "previously allocated pools by pool manager");
      }
      return TRUE;
    }
  }
  return FALSE;
}

UINT32 EptHookGetCountOfEpthooks(BOOLEAN IsEptHook2) {
  UINT32 Count = 0;
  LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL,
                     PageHookList, HookedEntry) {
    if (IsEptHook2) {
      if (HookedEntry->IsHiddenBreakpoint == FALSE) {
        Count++;
      }
    } else {
      if (HookedEntry->IsHiddenBreakpoint == TRUE) {
        Count++;
      }
    }
  }
  return Count;
}

BOOLEAN EptHookUnHookSingleAddressDetoursAndMonitor(
    PEPT_HOOKED_PAGE_DETAIL HookedEntry, BOOLEAN ApplyDirectlyFromVmxRoot,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails) {
  TargetUnhookingDetails->PhysicalAddress = HookedEntry->PhysicalBaseAddress;
  TargetUnhookingDetails->OriginalEntry = HookedEntry->OriginalEntry.AsUInt;
  if (ApplyDirectlyFromVmxRoot) {
    TargetUnhookingDetails->CallerNeedsToRestoreEntryAndInvalidateEpt = TRUE;
  } else {
    TargetUnhookingDetails->CallerNeedsToRestoreEntryAndInvalidateEpt = FALSE;
    KeGenericCallDpc(DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores,
                     TargetUnhookingDetails);
  }
  if (HookedEntry->IsExecutionHook) {
    EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(
        HookedEntry->VirtualAddress);
  }
  RemoveEntryList(&HookedEntry->PageHookList);
  if (!PoolManagerFreePool((UINT64)HookedEntry)) {
    LogError(
        "Err, something goes wrong, the pool not found in the list of "
        "previously allocated pools by pool manager");
    return FALSE;
  }
  return TRUE;
}

VOID EptHookHandleMonitorTrapFlag(VIRTUAL_MACHINE_STATE *VCpu) {
  PVOID TargetPage;
  TargetPage = EptGetPml1Entry(
      VCpu->EptPageTable, VCpu->MtfEptHookRestorePoint->PhysicalBaseAddress);
  EptSetPML1AndInvalidateTLB(VCpu, TargetPage,
                             VCpu->MtfEptHookRestorePoint->ChangedEntry,
                             InveptSingleContext);
  if (VCpu->MtfEptHookRestorePoint->IsPostEventTriggerAllowed) {
    if (VCpu->MtfEptHookRestorePoint->LastViolation ==
        EPT_HOOKED_LAST_VIOLATION_READ) {
      DispatchEventHiddenHookPageReadWriteExecReadPostEvent(
          VCpu, &VCpu->MtfEptHookRestorePoint->LastContextState);
    } else if (VCpu->MtfEptHookRestorePoint->LastViolation ==
               EPT_HOOKED_LAST_VIOLATION_WRITE) {
      DispatchEventHiddenHookPageReadWriteExecWritePostEvent(
          VCpu, &VCpu->MtfEptHookRestorePoint->LastContextState);
    } else if (VCpu->MtfEptHookRestorePoint->LastViolation ==
               EPT_HOOKED_LAST_VIOLATION_EXEC) {
      DispatchEventHiddenHookPageReadWriteExecExecutePostEvent(
          VCpu, &VCpu->MtfEptHookRestorePoint->LastContextState);
    }
  }
  VmmCallbackRestoreEptState(VCpu->CoreId);
}

BOOLEAN EptHookUnHookSingleAddressHiddenBreakpoint(
    PEPT_HOOKED_PAGE_DETAIL HookedEntry, UINT64 VirtualAddress,
    BOOLEAN ApplyDirectlyFromVmxRoot,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails) {
  UINT64 TargetAddressInFakePageContent;
  UINT32 CountOfEntriesWithSameAddr = 0;
  TargetUnhookingDetails->RemoveBreakpointInterception = FALSE;
  for (size_t i = 0; i < HookedEntry->CountOfBreakpoints; i++) {
    if (HookedEntry->BreakpointAddresses[i] == VirtualAddress) {
      if (HookedEntry->CountOfBreakpoints == 1) {
        TargetUnhookingDetails->PhysicalAddress =
            HookedEntry->PhysicalBaseAddress;
        TargetUnhookingDetails->OriginalEntry =
            HookedEntry->OriginalEntry.AsUInt;
        if (ApplyDirectlyFromVmxRoot) {
          TargetUnhookingDetails->CallerNeedsToRestoreEntryAndInvalidateEpt =
              TRUE;
        } else {
          TargetUnhookingDetails->CallerNeedsToRestoreEntryAndInvalidateEpt =
              FALSE;
          KeGenericCallDpc(
              DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores,
              TargetUnhookingDetails);
        }
        RemoveEntryList(&HookedEntry->PageHookList);
        if (!PoolManagerFreePool((UINT64)HookedEntry)) {
          LogError(
              "Err, something goes wrong, the pool not found in the list of "
              "previously allocated pools by pool manager");
        }
        if (EptHookGetCountOfEpthooks(FALSE) == 0) {
          if (ApplyDirectlyFromVmxRoot) {
            TargetUnhookingDetails->RemoveBreakpointInterception = TRUE;
          } else {
            TargetUnhookingDetails->RemoveBreakpointInterception = FALSE;
            BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores();
          }
        }
        return TRUE;
      } else {
        TargetAddressInFakePageContent =
            EptHookCalcBreakpointOffset((PVOID)VirtualAddress, HookedEntry);
        for (size_t j = 0; j < HookedEntry->CountOfBreakpoints; j++) {
          if (HookedEntry->BreakpointAddresses[j] == VirtualAddress) {
            CountOfEntriesWithSameAddr++;
          }
        }
        if (CountOfEntriesWithSameAddr == 1) {
          *(BYTE *)TargetAddressInFakePageContent =
              HookedEntry->PreviousBytesOnBreakpointAddresses[i];
        }
        HookedEntry->BreakpointAddresses[i] = NULL64_ZERO;
        HookedEntry->PreviousBytesOnBreakpointAddresses[i] = 0x0;
        for (size_t j = i; j < HookedEntry->CountOfBreakpoints - 1; j++) {
          HookedEntry->BreakpointAddresses[j] =
              HookedEntry->BreakpointAddresses[j + 1];
          HookedEntry->PreviousBytesOnBreakpointAddresses[j] =
              HookedEntry->PreviousBytesOnBreakpointAddresses[j + 1];
        }
        HookedEntry->CountOfBreakpoints = HookedEntry->CountOfBreakpoints - 1;
        return TRUE;
      }
    }
  }
  return FALSE;
}

BOOLEAN EptHookPerformUnHookSingleAddress(
    UINT64 VirtualAddress, UINT64 PhysAddress, UINT64 HookingTag,
    UINT32 ProcessId, BOOLEAN ApplyDirectlyFromVmxRoot,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails) {
  SIZE_T PhysicalAddress = NULL64_ZERO;
  if (ApplyDirectlyFromVmxRoot ||
      ProcessId == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES || ProcessId == 0) {
    ProcessId = HANDLE_TO_UINT32(PsGetCurrentProcessId());
  }
  if (PhysAddress != NULL64_ZERO) {
    PhysicalAddress = (SIZE_T)PAGE_ALIGN(PhysAddress);
  } else if (HookingTag == NULL64_ZERO) {
    if (ApplyDirectlyFromVmxRoot) {
      PhysicalAddress =
          (SIZE_T)(PAGE_ALIGN(VirtualAddressToPhysicalAddressOnTargetProcess(
              (PVOID)VirtualAddress)));
    } else {
      PhysicalAddress =
          (SIZE_T)(PAGE_ALIGN(VirtualAddressToPhysicalAddressByProcessId(
              (PVOID)VirtualAddress, ProcessId)));
    }
  }
  LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL,
                     PageHookList, CurrEntity) {
    if (CurrEntity->IsHiddenBreakpoint) {
      for (size_t i = 0; i < CurrEntity->CountOfBreakpoints; i++) {
        if (CurrEntity->BreakpointAddresses[i] == VirtualAddress) {
          return EptHookUnHookSingleAddressHiddenBreakpoint(
              CurrEntity, VirtualAddress, ApplyDirectlyFromVmxRoot,
              TargetUnhookingDetails);
        }
      }
    } else {
      if ((HookingTag != NULL64_ZERO && CurrEntity->HookingTag == HookingTag) ||
          CurrEntity->PhysicalBaseAddress == PhysicalAddress) {
        return EptHookUnHookSingleAddressDetoursAndMonitor(
            CurrEntity, ApplyDirectlyFromVmxRoot, TargetUnhookingDetails);
      }
    }
  }
  return FALSE;
}

BOOLEAN EptHookUnHookAllByHookingTag(UINT64 HookingTag) {
  EPT_SINGLE_HOOK_UNHOOKING_DETAILS TargetUnhookingDetails;
  BOOLEAN AtLeastOneUnhooked = FALSE;
  BOOLEAN UnhookingResult = FALSE;
  if (VmxGetCurrentExecutionMode() == TRUE) {
    return FALSE;
  }
KeepUnhooking:
  UnhookingResult = EptHookPerformUnHookSingleAddress(
      NULL_ZERO, NULL_ZERO, HookingTag, NULL_ZERO, FALSE,
      &TargetUnhookingDetails);
  if (UnhookingResult) {
    AtLeastOneUnhooked = TRUE;
    goto KeepUnhooking;
  }
  return AtLeastOneUnhooked;
}

BOOLEAN EptHookUnHookSingleHookByHookingTagFromVmxRoot(
    UINT64 HookingTag,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails) {
  if (VmxGetCurrentExecutionMode() == FALSE) {
    return FALSE;
  }
  return EptHookPerformUnHookSingleAddress(NULL64_ZERO, NULL64_ZERO, HookingTag,
                                           NULL_ZERO, TRUE,
                                           TargetUnhookingDetails);
}

BOOLEAN EptHookUnHookSingleAddress(UINT64 VirtualAddress, UINT64 PhysAddress,
                                   UINT32 ProcessId) {
  EPT_SINGLE_HOOK_UNHOOKING_DETAILS TargetUnhookingDetails;
  if (VmxGetCurrentExecutionMode() == TRUE) {
    return FALSE;
  }
  return EptHookPerformUnHookSingleAddress(VirtualAddress, PhysAddress,
                                           NULL_ZERO, ProcessId, FALSE,
                                           &TargetUnhookingDetails);
}

BOOLEAN EptHookUnHookSingleAddressFromVmxRoot(
    UINT64 VirtualAddress, UINT64 PhysAddress,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails) {
  if (VmxGetCurrentExecutionMode() == FALSE) {
    return FALSE;
  }
  return EptHookPerformUnHookSingleAddress(VirtualAddress, PhysAddress,
                                           NULL64_ZERO, NULL_ZERO, TRUE,
                                           TargetUnhookingDetails);
}

VOID EptHookUnHookAll() {
  if (VmxGetCurrentExecutionMode() == TRUE) {
    return;
  }
  KeGenericCallDpc(DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores, 0x0);
  LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL,
                     PageHookList, CurrEntity) {
    if (!CurrEntity->IsHiddenBreakpoint) {
      EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(
          CurrEntity->VirtualAddress);
    }
    if (!PoolManagerFreePool((UINT64)CurrEntity)) {
      LogError(
          "Err, something goes wrong, the pool not found in the list of "
          "previously allocated pools by pool manager");
    }
  }
}

PVOID EptHook2GeneralDetourEventHandler(PGUEST_REGS Regs, PVOID CalledFrom) {
  PLIST_ENTRY TempList = 0;
  EPT_HOOKS_CONTEXT TempContext = {0};
  Regs->rsp = (UINT64)Regs - sizeof(GUEST_REGS);
  TempContext.VirtualAddress = (UINT64)CalledFrom;
  TempContext.PhysicalAddress = VirtualAddressToPhysicalAddress(CalledFrom);
  VIRTUAL_MACHINE_STATE *VCpu =
      &g_GuestState[KeGetCurrentProcessorNumberEx(NULL)];
  VCpu->Regs = Regs;
  DispatchEventHiddenHookExecDetours(VCpu, &TempContext);
  TempList = &g_EptHook2sDetourListHead;
  while (&g_EptHook2sDetourListHead != TempList->Flink) {
    TempList = TempList->Flink;
    PHIDDEN_HOOKS_DETOUR_DETAILS CurrentHookedDetails = CONTAINING_RECORD(
        TempList, HIDDEN_HOOKS_DETOUR_DETAILS, OtherHooksList);
    if (CurrentHookedDetails->HookedFunctionAddress == CalledFrom) {
      return CurrentHookedDetails->ReturnAddress;
    }
  }
  return CalledFrom;
}

BOOLEAN EptHookModifyInstructionFetchState(VIRTUAL_MACHINE_STATE *VCpu,
                                           PVOID PhysicalAddress,
                                           BOOLEAN IsUnset) {
  PVOID PmlEntry = NULL;
  BOOLEAN IsLargePage = FALSE;
  PmlEntry = EptGetPml1OrPml2Entry(g_EptState->EptPageTable,
                                   (SIZE_T)PhysicalAddress, &IsLargePage);
  if (PmlEntry) {
    if (IsLargePage) {
      if (IsUnset) {
        ((PEPT_PML2_ENTRY)PmlEntry)->ExecuteAccess = FALSE;
      } else {
        ((PEPT_PML2_ENTRY)PmlEntry)->ExecuteAccess = TRUE;
      }
    } else {
      if (IsUnset) {
        ((PEPT_PML1_ENTRY)PmlEntry)->ExecuteAccess = FALSE;
      } else {
        ((PEPT_PML1_ENTRY)PmlEntry)->ExecuteAccess = TRUE;
      }
    }
  } else {
    return FALSE;
  }
  EptInveptSingleContext(VCpu->EptPointer.AsUInt);
  return TRUE;
}

BOOLEAN EptHookModifyPageReadState(VIRTUAL_MACHINE_STATE *VCpu,
                                   PVOID PhysicalAddress, BOOLEAN IsUnset) {
  PVOID PmlEntry = NULL;
  BOOLEAN IsLargePage = FALSE;
  PmlEntry = EptGetPml1OrPml2Entry(g_EptState->EptPageTable,
                                   (SIZE_T)PhysicalAddress, &IsLargePage);
  if (PmlEntry) {
    if (IsLargePage) {
      if (IsUnset) {
        ((PEPT_PML2_ENTRY)PmlEntry)->ReadAccess = FALSE;
      } else {
        ((PEPT_PML2_ENTRY)PmlEntry)->ReadAccess = TRUE;
      }
    } else {
      if (IsUnset) {
        ((PEPT_PML1_ENTRY)PmlEntry)->ReadAccess = FALSE;
      } else {
        ((PEPT_PML1_ENTRY)PmlEntry)->ReadAccess = TRUE;
      }
    }
  } else {
    return FALSE;
  }
  EptInveptSingleContext(VCpu->EptPointer.AsUInt);
  return TRUE;
}

BOOLEAN EptHookModifyPageWriteState(VIRTUAL_MACHINE_STATE *VCpu,
                                    PVOID PhysicalAddress, BOOLEAN IsUnset) {
  PVOID PmlEntry = NULL;
  BOOLEAN IsLargePage = FALSE;
  PmlEntry = EptGetPml1OrPml2Entry(g_EptState->EptPageTable,
                                   (SIZE_T)PhysicalAddress, &IsLargePage);
  if (PmlEntry) {
    if (IsLargePage) {
      if (IsUnset) {
        ((PEPT_PML2_ENTRY)PmlEntry)->WriteAccess = FALSE;
      } else {
        ((PEPT_PML2_ENTRY)PmlEntry)->WriteAccess = TRUE;
      }
    } else {
      if (IsUnset) {
        ((PEPT_PML1_ENTRY)PmlEntry)->WriteAccess = FALSE;
      } else {
        ((PEPT_PML1_ENTRY)PmlEntry)->WriteAccess = TRUE;
      }
    }
  } else {
    return FALSE;
  }
  EptInveptSingleContext(VCpu->EptPointer.AsUInt);
  return TRUE;
}
