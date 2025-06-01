
#include "pch.h"

BOOLEAN EptCheckFeatures(VOID) {
  IA32_VMX_EPT_VPID_CAP_REGISTER VpidRegister;
  IA32_MTRR_DEF_TYPE_REGISTER MTRRDefType;
  VpidRegister.AsUInt = __readmsr(IA32_VMX_EPT_VPID_CAP);
  MTRRDefType.AsUInt = __readmsr(IA32_MTRR_DEF_TYPE);
  if (!VpidRegister.PageWalkLength4 || !VpidRegister.MemoryTypeWriteBack ||
      !VpidRegister.Pde2MbPages) {
    return FALSE;
  }
  if (!VpidRegister.AdvancedVmexitEptViolationsInformation) {
    LogDebugInfo(
        "The processor doesn't report advanced VM-exit information for EPT "
        "violations");
  }
  if (!VpidRegister.ExecuteOnlyPages) {
    g_CompatibilityCheck.ExecuteOnlySupport = FALSE;
    LogDebugInfo(
        "The processor doesn't support execute-only pages, execute hooks won't "
        "work as they're on this feature in our design");
  } else {
    g_CompatibilityCheck.ExecuteOnlySupport = TRUE;
  }
  if (!MTRRDefType.MtrrEnable) {
    LogError("Err, MTRR dynamic ranges are not supported");
    return FALSE;
  }
  LogDebugInfo("All EPT features are present");
  return TRUE;
}

UINT8 EptGetMemoryType(SIZE_T PageFrameNumber, BOOLEAN IsLargePage) {
  UINT8 TargetMemoryType;
  SIZE_T AddressOfPage;
  SIZE_T CurrentMtrrRange;
  MTRR_RANGE_DESCRIPTOR *CurrentMemoryRange;
  AddressOfPage =
      IsLargePage ? PageFrameNumber * SIZE_2_MB : PageFrameNumber * PAGE_SIZE;
  TargetMemoryType = (UINT8)-1;
  for (CurrentMtrrRange = 0;
       CurrentMtrrRange < g_EptState->NumberOfEnabledMemoryRanges;
       CurrentMtrrRange++) {
    CurrentMemoryRange = &g_EptState->MemoryRanges[CurrentMtrrRange];
    if (AddressOfPage >= CurrentMemoryRange->PhysicalBaseAddress &&
        AddressOfPage < CurrentMemoryRange->PhysicalEndAddress) {
      if (CurrentMemoryRange->FixedRange) {
        TargetMemoryType = CurrentMemoryRange->MemoryType;
        break;
      }
      if (TargetMemoryType == MEMORY_TYPE_UNCACHEABLE) {
        TargetMemoryType = CurrentMemoryRange->MemoryType;
        break;
      }
      if (TargetMemoryType == MEMORY_TYPE_WRITE_THROUGH ||
          CurrentMemoryRange->MemoryType == MEMORY_TYPE_WRITE_THROUGH) {
        if (TargetMemoryType == MEMORY_TYPE_WRITE_BACK) {
          TargetMemoryType = MEMORY_TYPE_WRITE_THROUGH;
          continue;
        }
      }
      TargetMemoryType = CurrentMemoryRange->MemoryType;
    }
  }
  if (TargetMemoryType == (UINT8)-1) {
    TargetMemoryType = g_EptState->DefaultMemoryType;
  }
  return TargetMemoryType;
}

BOOLEAN EptBuildMtrrMap(VOID) {
  IA32_MTRR_CAPABILITIES_REGISTER MTRRCap;
  IA32_MTRR_PHYSBASE_REGISTER CurrentPhysBase;
  IA32_MTRR_PHYSMASK_REGISTER CurrentPhysMask;
  IA32_MTRR_DEF_TYPE_REGISTER MTRRDefType;
  PMTRR_RANGE_DESCRIPTOR Descriptor;
  UINT32 CurrentRegister;
  UINT32 NumberOfBitsInMask;
  MTRRCap.AsUInt = __readmsr(IA32_MTRR_CAPABILITIES);
  MTRRDefType.AsUInt = __readmsr(IA32_MTRR_DEF_TYPE);
  if (!MTRRDefType.MtrrEnable) {
    g_EptState->DefaultMemoryType = MEMORY_TYPE_UNCACHEABLE;
    return TRUE;
  }
  g_EptState->DefaultMemoryType = (UINT8)MTRRDefType.DefaultMemoryType;
  if (MTRRCap.FixedRangeSupported && MTRRDefType.FixedRangeMtrrEnable) {
    const UINT32 K64Base = 0x0;
    const UINT32 K64Size = 0x10000;
    IA32_MTRR_FIXED_RANGE_TYPE K64Types = {__readmsr(IA32_MTRR_FIX64K_00000)};
    for (unsigned int i = 0; i < 8; i++) {
      Descriptor =
          &g_EptState->MemoryRanges[g_EptState->NumberOfEnabledMemoryRanges++];
      Descriptor->MemoryType = K64Types.s.Types[i];
      Descriptor->PhysicalBaseAddress = K64Base + (K64Size * i);
      Descriptor->PhysicalEndAddress = K64Base + (K64Size * i) + (K64Size - 1);
      Descriptor->FixedRange = TRUE;
    }
    const UINT32 K16Base = 0x80000;
    const UINT32 K16Size = 0x4000;
    for (unsigned int i = 0; i < 2; i++) {
      IA32_MTRR_FIXED_RANGE_TYPE K16Types = {
          __readmsr(IA32_MTRR_FIX16K_80000 + i)};
      for (unsigned int j = 0; j < 8; j++) {
        Descriptor =
            &g_EptState
                 ->MemoryRanges[g_EptState->NumberOfEnabledMemoryRanges++];
        Descriptor->MemoryType = K16Types.s.Types[j];
        Descriptor->PhysicalBaseAddress =
            (K16Base + (i * K16Size * 8)) + (K16Size * j);
        Descriptor->PhysicalEndAddress =
            (K16Base + (i * K16Size * 8)) + (K16Size * j) + (K16Size - 1);
        Descriptor->FixedRange = TRUE;
      }
    }
    const UINT32 K4Base = 0xC0000;
    const UINT32 K4Size = 0x1000;
    for (unsigned int i = 0; i < 8; i++) {
      IA32_MTRR_FIXED_RANGE_TYPE K4Types = {
          __readmsr(IA32_MTRR_FIX4K_C0000 + i)};
      for (unsigned int j = 0; j < 8; j++) {
        Descriptor =
            &g_EptState
                 ->MemoryRanges[g_EptState->NumberOfEnabledMemoryRanges++];
        Descriptor->MemoryType = K4Types.s.Types[j];
        Descriptor->PhysicalBaseAddress =
            (K4Base + (i * K4Size * 8)) + (K4Size * j);
        Descriptor->PhysicalEndAddress =
            (K4Base + (i * K4Size * 8)) + (K4Size * j) + (K4Size - 1);
        Descriptor->FixedRange = TRUE;
      }
    }
  }
  for (CurrentRegister = 0; CurrentRegister < MTRRCap.VariableRangeCount;
       CurrentRegister++) {
    CurrentPhysBase.AsUInt =
        __readmsr(IA32_MTRR_PHYSBASE0 + (CurrentRegister * 2));
    CurrentPhysMask.AsUInt =
        __readmsr(IA32_MTRR_PHYSMASK0 + (CurrentRegister * 2));
    if (CurrentPhysMask.Valid) {
      Descriptor =
          &g_EptState->MemoryRanges[g_EptState->NumberOfEnabledMemoryRanges++];
      Descriptor->PhysicalBaseAddress =
          CurrentPhysBase.PageFrameNumber * PAGE_SIZE;
      _BitScanForward64((ULONG *)&NumberOfBitsInMask,
                        CurrentPhysMask.PageFrameNumber * PAGE_SIZE);
      Descriptor->PhysicalEndAddress = Descriptor->PhysicalBaseAddress +
                                       ((1ULL << NumberOfBitsInMask) - 1ULL);
      Descriptor->MemoryType = (UCHAR)CurrentPhysBase.Type;
      Descriptor->FixedRange = FALSE;
      LogDebugInfo("MTRR Range: Base=0x%llx End=0x%llx Type=0x%x",
                   Descriptor->PhysicalBaseAddress,
                   Descriptor->PhysicalEndAddress, Descriptor->MemoryType);
    }
  }
  LogDebugInfo("Total MTRR ranges committed: 0x%x",
               g_EptState->NumberOfEnabledMemoryRanges);
  return TRUE;
}

PEPT_PML1_ENTRY EptGetPml1Entry(PVMM_EPT_PAGE_TABLE EptPageTable,
                                SIZE_T PhysicalAddress) {
  SIZE_T Directory, DirectoryPointer, PML4Entry;
  PEPT_PML2_ENTRY PML2;
  PEPT_PML1_ENTRY PML1;
  PEPT_PML2_POINTER PML2Pointer;
  Directory = ADDRMASK_EPT_PML2_INDEX(PhysicalAddress);
  DirectoryPointer = ADDRMASK_EPT_PML3_INDEX(PhysicalAddress);
  PML4Entry = ADDRMASK_EPT_PML4_INDEX(PhysicalAddress);
  if (PML4Entry > 0) {
    return NULL;
  }
  PML2 = &EptPageTable->PML2[DirectoryPointer][Directory];
  if (PML2->LargePage) {
    return NULL;
  }
  PML2Pointer = (PEPT_PML2_POINTER)PML2;
  PML1 = (PEPT_PML1_ENTRY)PhysicalAddressToVirtualAddress(
      PML2Pointer->PageFrameNumber * PAGE_SIZE);
  if (!PML1) {
    return NULL;
  }
  PML1 = &PML1[ADDRMASK_EPT_PML1_INDEX(PhysicalAddress)];
  return PML1;
}

PVOID EptGetPml1OrPml2Entry(PVMM_EPT_PAGE_TABLE EptPageTable,
                            SIZE_T PhysicalAddress, BOOLEAN *IsLargePage) {
  SIZE_T Directory, DirectoryPointer, PML4Entry;
  PEPT_PML2_ENTRY PML2;
  PEPT_PML1_ENTRY PML1;
  PEPT_PML2_POINTER PML2Pointer;
  Directory = ADDRMASK_EPT_PML2_INDEX(PhysicalAddress);
  DirectoryPointer = ADDRMASK_EPT_PML3_INDEX(PhysicalAddress);
  PML4Entry = ADDRMASK_EPT_PML4_INDEX(PhysicalAddress);
  if (PML4Entry > 0) {
    return NULL;
  }
  PML2 = &EptPageTable->PML2[DirectoryPointer][Directory];
  if (PML2->LargePage) {
    *IsLargePage = TRUE;
    return PML2;
  }
  PML2Pointer = (PEPT_PML2_POINTER)PML2;
  PML1 = (PEPT_PML1_ENTRY)PhysicalAddressToVirtualAddress(
      PML2Pointer->PageFrameNumber * PAGE_SIZE);
  if (!PML1) {
    return NULL;
  }
  PML1 = &PML1[ADDRMASK_EPT_PML1_INDEX(PhysicalAddress)];
  *IsLargePage = FALSE;
  return PML1;
}

PEPT_PML2_ENTRY EptGetPml2Entry(PVMM_EPT_PAGE_TABLE EptPageTable,
                                SIZE_T PhysicalAddress) {
  SIZE_T Directory, DirectoryPointer, PML4Entry;
  PEPT_PML2_ENTRY PML2;
  Directory = ADDRMASK_EPT_PML2_INDEX(PhysicalAddress);
  DirectoryPointer = ADDRMASK_EPT_PML3_INDEX(PhysicalAddress);
  PML4Entry = ADDRMASK_EPT_PML4_INDEX(PhysicalAddress);
  if (PML4Entry > 0) {
    return NULL;
  }
  PML2 = &EptPageTable->PML2[DirectoryPointer][Directory];
  return PML2;
}

BOOLEAN EptSplitLargePage(PVMM_EPT_PAGE_TABLE EptPageTable,
                          PVOID PreAllocatedBuffer, SIZE_T PhysicalAddress) {
  PVMM_EPT_DYNAMIC_SPLIT NewSplit;
  EPT_PML1_ENTRY EntryTemplate;
  SIZE_T EntryIndex;
  PEPT_PML2_ENTRY TargetEntry;
  EPT_PML2_POINTER NewPointer;
  TargetEntry = EptGetPml2Entry(EptPageTable, PhysicalAddress);
  if (!TargetEntry) {
    LogError("Err, an invalid physical address passed");
    return FALSE;
  }
  if (!TargetEntry->LargePage) {
    PoolManagerFreePool((UINT64)PreAllocatedBuffer);
    return TRUE;
  }
  NewSplit = (PVMM_EPT_DYNAMIC_SPLIT)PreAllocatedBuffer;
  if (!NewSplit) {
    LogError("Err, failed to allocate dynamic split memory");
    return FALSE;
  }
  RtlZeroMemory(NewSplit, sizeof(VMM_EPT_DYNAMIC_SPLIT));
  NewSplit->u.Entry = TargetEntry;
  EntryTemplate.AsUInt = 0;
  EntryTemplate.ReadAccess = 1;
  EntryTemplate.WriteAccess = 1;
  EntryTemplate.ExecuteAccess = 1;
  EntryTemplate.MemoryType = TargetEntry->MemoryType;
  EntryTemplate.IgnorePat = TargetEntry->IgnorePat;
  EntryTemplate.SuppressVe = TargetEntry->SuppressVe;
  __stosq((SIZE_T *)&NewSplit->PML1[0], EntryTemplate.AsUInt,
          VMM_EPT_PML1E_COUNT);
  for (EntryIndex = 0; EntryIndex < VMM_EPT_PML1E_COUNT; EntryIndex++) {
    NewSplit->PML1[EntryIndex].PageFrameNumber =
        ((TargetEntry->PageFrameNumber * SIZE_2_MB) / PAGE_SIZE) + EntryIndex;
    NewSplit->PML1[EntryIndex].MemoryType =
        EptGetMemoryType(NewSplit->PML1[EntryIndex].PageFrameNumber, FALSE);
  }
  NewPointer.AsUInt = 0;
  NewPointer.WriteAccess = 1;
  NewPointer.ReadAccess = 1;
  NewPointer.ExecuteAccess = 1;
  NewPointer.PageFrameNumber =
      (SIZE_T)VirtualAddressToPhysicalAddress(&NewSplit->PML1[0]) / PAGE_SIZE;
  RtlCopyMemory(TargetEntry, &NewPointer, sizeof(NewPointer));
  return TRUE;
}

BOOLEAN EptIsValidForLargePage(SIZE_T PageFrameNumber) {
  SIZE_T StartAddressOfPage = PageFrameNumber * SIZE_2_MB;
  SIZE_T EndAddressOfPage = StartAddressOfPage + (SIZE_2_MB - 1);
  MTRR_RANGE_DESCRIPTOR *CurrentMemoryRange;
  SIZE_T CurrentMtrrRange;
  for (CurrentMtrrRange = 0;
       CurrentMtrrRange < g_EptState->NumberOfEnabledMemoryRanges;
       CurrentMtrrRange++) {
    CurrentMemoryRange = &g_EptState->MemoryRanges[CurrentMtrrRange];
    if ((StartAddressOfPage <= CurrentMemoryRange->PhysicalEndAddress &&
         EndAddressOfPage > CurrentMemoryRange->PhysicalEndAddress) ||
        (StartAddressOfPage < CurrentMemoryRange->PhysicalBaseAddress &&
         EndAddressOfPage >= CurrentMemoryRange->PhysicalBaseAddress)) {
      return FALSE;
    }
  }
  return TRUE;
}

BOOLEAN EptSetupPML2Entry(PVMM_EPT_PAGE_TABLE EptPageTable,
                          PEPT_PML2_ENTRY NewEntry, SIZE_T PageFrameNumber) {
  PVOID TargetBuffer;
  NewEntry->PageFrameNumber = PageFrameNumber;
  if (EptIsValidForLargePage(PageFrameNumber)) {
    NewEntry->MemoryType = EptGetMemoryType(PageFrameNumber, TRUE);
    return TRUE;
  } else {
    TargetBuffer =
        (PVOID)PlatformMemAllocateNonPagedPool(sizeof(VMM_EPT_DYNAMIC_SPLIT));
    if (!TargetBuffer) {
      LogError("Err, cannot allocate page for splitting edge large pages");
      return FALSE;
    }
    return EptSplitLargePage(EptPageTable, TargetBuffer,
                             PageFrameNumber * SIZE_2_MB);
  }
}

PVMM_EPT_PAGE_TABLE EptAllocateAndCreateIdentityPageTable(VOID) {
  PVMM_EPT_PAGE_TABLE PageTable;
  EPT_PML3_POINTER RWXTemplate;
  EPT_PML2_ENTRY PML2EntryTemplate;
  SIZE_T EntryGroupIndex;
  SIZE_T EntryIndex;
  PageTable =
      PlatformMemAllocateContiguousZeroedMemory(sizeof(VMM_EPT_PAGE_TABLE));
  if (PageTable == NULL) {
    LogError("Err, failed to allocate memory for PageTable");
    return NULL;
  }
  PageTable->PML4[0].PageFrameNumber =
      (SIZE_T)VirtualAddressToPhysicalAddress(&PageTable->PML3[0]) / PAGE_SIZE;
  PageTable->PML4[0].ReadAccess = 1;
  PageTable->PML4[0].WriteAccess = 1;
  PageTable->PML4[0].ExecuteAccess = 1;
  RWXTemplate.AsUInt = 0;
  RWXTemplate.ReadAccess = 1;
  RWXTemplate.WriteAccess = 1;
  RWXTemplate.ExecuteAccess = 1;
  __stosq((SIZE_T *)&PageTable->PML3[0], RWXTemplate.AsUInt,
          VMM_EPT_PML3E_COUNT);
  for (EntryIndex = 0; EntryIndex < VMM_EPT_PML3E_COUNT; EntryIndex++) {
    PageTable->PML3[EntryIndex].PageFrameNumber =
        (SIZE_T)VirtualAddressToPhysicalAddress(
            &PageTable->PML2[EntryIndex][0]) /
        PAGE_SIZE;
  }
  PML2EntryTemplate.AsUInt = 0;
  PML2EntryTemplate.WriteAccess = 1;
  PML2EntryTemplate.ReadAccess = 1;
  PML2EntryTemplate.ExecuteAccess = 1;
  PML2EntryTemplate.LargePage = 1;
  __stosq((SIZE_T *)&PageTable->PML2[0], PML2EntryTemplate.AsUInt,
          VMM_EPT_PML3E_COUNT * VMM_EPT_PML2E_COUNT);
  for (EntryGroupIndex = 0; EntryGroupIndex < VMM_EPT_PML3E_COUNT;
       EntryGroupIndex++) {
    for (EntryIndex = 0; EntryIndex < VMM_EPT_PML2E_COUNT; EntryIndex++) {
      EptSetupPML2Entry(PageTable,
                        &PageTable->PML2[EntryGroupIndex][EntryIndex],
                        (EntryGroupIndex * VMM_EPT_PML2E_COUNT) + EntryIndex);
    }
  }
  return PageTable;
}

BOOLEAN EptLogicalProcessorInitialize(VOID) {
  ULONG ProcessorsCount;
  PVMM_EPT_PAGE_TABLE PageTable;
  EPT_POINTER EPTP = {0};
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  for (size_t i = 0; i < ProcessorsCount; i++) {
    PageTable = EptAllocateAndCreateIdentityPageTable();
    if (!PageTable) {
      for (size_t j = 0; j < ProcessorsCount; j++) {
        if (g_GuestState[j].EptPageTable != NULL) {
          MmFreeContiguousMemory(g_GuestState[j].EptPageTable);
          g_GuestState[j].EptPageTable = NULL;
        }
      }
      LogError("Err, unable to allocate memory for EPT");
      return FALSE;
    }
    g_GuestState[i].EptPageTable = PageTable;
    EPTP.MemoryType = g_EptState->DefaultMemoryType;
    EPTP.EnableAccessAndDirtyFlags = TRUE;
    EPTP.PageWalkLength = 3;
    EPTP.PageFrameNumber =
        (SIZE_T)VirtualAddressToPhysicalAddress(&PageTable->PML4) / PAGE_SIZE;
    g_GuestState[i].EptPointer = EPTP;
  }
  return TRUE;
}

_Use_decl_annotations_

    BOOLEAN
    EptHandlePageHookExit(
        VIRTUAL_MACHINE_STATE *VCpu,
        VMX_EXIT_QUALIFICATION_EPT_VIOLATION ViolationQualification,
        UINT64 GuestPhysicalAddr) {
  PVOID TargetPage;
  UINT64 CurrentRip;
  UINT32 CurrentInstructionLength;
  BOOLEAN IsHandled = FALSE;
  BOOLEAN ResultOfHandlingHook = FALSE;
  BOOLEAN IgnoreReadOrWriteOrExec = FALSE;
  BOOLEAN IsExecViolation = FALSE;
  LIST_FOR_EACH_LINK(g_EptState->HookedPagesList, EPT_HOOKED_PAGE_DETAIL,
                     PageHookList, HookedEntry) {
    if (HookedEntry->PhysicalBaseAddress ==
        (SIZE_T)PAGE_ALIGN(GuestPhysicalAddr)) {
      if (GuestPhysicalAddr >= HookedEntry->StartOfTargetPhysicalAddress &&
          GuestPhysicalAddr <= HookedEntry->EndOfTargetPhysicalAddress) {
        ResultOfHandlingHook = EptHookHandleHookedPage(
            VCpu, HookedEntry, ViolationQualification, GuestPhysicalAddr,
            &HookedEntry->LastContextState, &IgnoreReadOrWriteOrExec,
            &IsExecViolation);
      } else {
        ResultOfHandlingHook = TRUE;
      }
      if (ResultOfHandlingHook) {
        if (!IgnoreReadOrWriteOrExec) {
          TargetPage = EptGetPml1Entry(VCpu->EptPageTable,
                                       HookedEntry->PhysicalBaseAddress);
          EptSetPML1AndInvalidateTLB(VCpu, TargetPage,
                                     HookedEntry->OriginalEntry,
                                     InveptSingleContext);
          VCpu->MtfEptHookRestorePoint = HookedEntry;
          HvEnableMtfAndChangeExternalInterruptState(VCpu);
        }
      }
      IsHandled = TRUE;
      break;
    }
  }
  if (IgnoreReadOrWriteOrExec) {
    HvSuppressRipIncrement(VCpu);
    if (!IsExecViolation) {
      CurrentRip = HvGetRip();
      CurrentInstructionLength =
          DisassemblerLengthDisassembleEngineInVmxRootOnTargetProcess(
              (PVOID)CurrentRip, CommonIsGuestOnUsermode32Bit());
      CurrentRip = CurrentRip + CurrentInstructionLength;
      HvSetRip(CurrentRip);
    }
  } else {
    HvSuppressRipIncrement(VCpu);
  }
  return IsHandled;
}

BOOLEAN EptHandleEptViolation(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT64 GuestPhysicalAddr;
  VMX_EXIT_QUALIFICATION_EPT_VIOLATION ViolationQualification = {
      .AsUInt = VCpu->ExitQualification};
  __vmx_vmread(VMCS_GUEST_PHYSICAL_ADDRESS, &GuestPhysicalAddr);
  if (ExecTrapHandleEptViolationVmexit(VCpu, &ViolationQualification)) {
    return TRUE;
  } else if (EptHandlePageHookExit(VCpu, ViolationQualification,
                                   GuestPhysicalAddr)) {
    return TRUE;
  } else if (VmmCallbackUnhandledEptViolation(
                 VCpu->CoreId, (UINT64)ViolationQualification.AsUInt,
                 GuestPhysicalAddr)) {
    return TRUE;
  }
  LogError("Err, unexpected EPT violation at RIP: %llx", VCpu->LastVmexitRip);
  DbgBreakPoint();
  return FALSE;
}

VOID EptHandleMisconfiguration(VOID) {
  UINT64 GuestPhysicalAddr = 0;
  __vmx_vmread(VMCS_GUEST_PHYSICAL_ADDRESS, &GuestPhysicalAddr);
  LogInfo("EPT Misconfiguration!");
  LogError(
      "Err, a field in the EPT paging structure was invalid, faulting guest "
      "address : 0x%llx",
      GuestPhysicalAddr);
}

_Use_decl_annotations_

    VOID
    EptSetPML1AndInvalidateTLB(VIRTUAL_MACHINE_STATE *VCpu,
                               PEPT_PML1_ENTRY EntryAddress,
                               EPT_PML1_ENTRY EntryValue,
                               INVEPT_TYPE InvalidationType) {
  EntryAddress->AsUInt = EntryValue.AsUInt;
  if (InvalidationType == InveptSingleContext) {
    EptInveptSingleContext(VCpu->EptPointer.AsUInt);
  } else if (InvalidationType == InveptAllContext) {
    EptInveptAllContexts();
  } else {
    LogError("Err, invalid invalidation parameter");
  }
}

BOOLEAN EptCheckAndHandleEptHookBreakpoints(VIRTUAL_MACHINE_STATE *VCpu,
                                            UINT64 GuestRip) {
  PVOID TargetPage;
  PLIST_ENTRY TempList;
  BOOLEAN IsHandledByEptHook = FALSE;
  TempList = &g_EptState->HookedPagesList;
  while (&g_EptState->HookedPagesList != TempList->Flink) {
    TempList = TempList->Flink;
    PEPT_HOOKED_PAGE_DETAIL HookedEntry =
        CONTAINING_RECORD(TempList, EPT_HOOKED_PAGE_DETAIL, PageHookList);
    if (HookedEntry->IsExecutionHook) {
      for (size_t i = 0; i < HookedEntry->CountOfBreakpoints; i++) {
        if (HookedEntry->BreakpointAddresses[i] == GuestRip) {
          DispatchEventHiddenHookExecCc(VCpu, (PVOID)GuestRip);
          TargetPage = EptGetPml1Entry(VCpu->EptPageTable,
                                       HookedEntry->PhysicalBaseAddress);
          EptSetPML1AndInvalidateTLB(VCpu, TargetPage,
                                     HookedEntry->OriginalEntry,
                                     InveptSingleContext);
          VCpu->MtfEptHookRestorePoint = HookedEntry;
          HvEnableMtfAndChangeExternalInterruptState(VCpu);
          IsHandledByEptHook = TRUE;
          break;
        }
      }
    }
  }
  return IsHandledByEptHook;
}

BOOLEAN EptCheckAndHandleBreakpoint(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT64 GuestRip = 0;
  BOOLEAN IsHandledByEptHook;
  __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
  HvSuppressRipIncrement(VCpu);
  IsHandledByEptHook = EptCheckAndHandleEptHookBreakpoints(VCpu, GuestRip);
  return IsHandledByEptHook;
}
