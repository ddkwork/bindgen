
#include "pch.h"

BOOLEAN ExecTrapTraverseThroughOsPageTables(PVMM_EPT_PAGE_TABLE EptTable,
                                            CR3_TYPE TargetCr3,
                                            CR3_TYPE KernelCr3) {
  CR3_TYPE Cr3;
  UINT64 TempCr3;
  PUINT64 Cr3Va;
  PUINT64 PdptVa;
  PUINT64 PdVa;
  PUINT64 PtVa;
  BOOLEAN IsLargePage = FALSE;
  CR3_TYPE CurrentProcessCr3 = {0};
  CurrentProcessCr3 = SwitchToProcessMemoryLayoutByCr3(KernelCr3);
  Cr3.Flags = TargetCr3.Flags;
  TempCr3 = Cr3.Fields.PageFrameNumber << 12;
  PVOID EptPmlEntry4 = EptGetPml1OrPml2Entry(
      EptTable, Cr3.Fields.PageFrameNumber << 12, &IsLargePage);
  if (EptPmlEntry4 != NULL) {
    if (IsLargePage) {
      ((PEPT_PML2_ENTRY)EptPmlEntry4)->ReadAccess = TRUE;
      ((PEPT_PML2_ENTRY)EptPmlEntry4)->WriteAccess = TRUE;
    } else {
      ((PEPT_PML1_ENTRY)EptPmlEntry4)->ReadAccess = TRUE;
      ((PEPT_PML1_ENTRY)EptPmlEntry4)->WriteAccess = TRUE;
    }
  } else {
    LogInfo("null address");
  }
  Cr3Va = (UINT64 *)PhysicalAddressToVirtualAddress(TempCr3);
  if (Cr3Va == NULL) {
    SwitchToPreviousProcess(CurrentProcessCr3);
    return FALSE;
  }
  for (size_t i = 0; i < 512; i++) {
    PPAGE_ENTRY Pml4e = (PAGE_ENTRY *)&Cr3Va[i];
    if (Pml4e->Fields.Present) {
      IsLargePage = FALSE;
      EptPmlEntry4 = EptGetPml1OrPml2Entry(
          EptTable, Pml4e->Fields.PageFrameNumber << 12, &IsLargePage);
      if (EptPmlEntry4 != NULL) {
        if (IsLargePage) {
          ((PEPT_PML2_ENTRY)EptPmlEntry4)->ReadAccess = TRUE;
          ((PEPT_PML2_ENTRY)EptPmlEntry4)->WriteAccess = TRUE;
        } else {
          ((PEPT_PML1_ENTRY)EptPmlEntry4)->ReadAccess = TRUE;
          ((PEPT_PML1_ENTRY)EptPmlEntry4)->WriteAccess = TRUE;
        }
      } else {
        LogInfo("null address");
      }
      PdptVa = (UINT64 *)PhysicalAddressToVirtualAddress(
          Pml4e->Fields.PageFrameNumber << 12);
      if (PdptVa != NULL) {
        for (size_t j = 0; j < 512; j++) {
          PPAGE_ENTRY Pdpte = (PAGE_ENTRY *)&PdptVa[j];
          if (Pdpte->Fields.Present) {
            IsLargePage = FALSE;
            PVOID EptPmlEntry3 = EptGetPml1OrPml2Entry(
                EptTable, Pdpte->Fields.PageFrameNumber << 12, &IsLargePage);
            if (EptPmlEntry3 != NULL) {
              if (IsLargePage) {
                ((PEPT_PML2_ENTRY)EptPmlEntry3)->ReadAccess = TRUE;
                ((PEPT_PML2_ENTRY)EptPmlEntry3)->WriteAccess = TRUE;
              } else {
                ((PEPT_PML1_ENTRY)EptPmlEntry3)->ReadAccess = TRUE;
                ((PEPT_PML1_ENTRY)EptPmlEntry3)->WriteAccess = TRUE;
              }
            } else {
              LogInfo("null address");
            }
            if (Pdpte->Fields.LargePage) {
              continue;
            }
            PdVa = (UINT64 *)PhysicalAddressToVirtualAddress(
                Pdpte->Fields.PageFrameNumber << 12);
            if (PdVa != NULL) {
              for (size_t k = 0; k < 512; k++) {
                if (PdVa == (PUINT64)0xfffffffffffffe00) {
                  continue;
                }
                PPAGE_ENTRY Pde = (PAGE_ENTRY *)&PdVa[k];
                if (Pde->Fields.Present) {
                  IsLargePage = FALSE;
                  PVOID EptPmlEntry2 = EptGetPml1OrPml2Entry(
                      EptTable, Pde->Fields.PageFrameNumber << 12,
                      &IsLargePage);
                  if (EptPmlEntry2 != NULL) {
                    if (IsLargePage) {
                      ((PEPT_PML2_ENTRY)EptPmlEntry2)->ReadAccess = TRUE;
                      ((PEPT_PML2_ENTRY)EptPmlEntry2)->WriteAccess = TRUE;
                    } else {
                      ((PEPT_PML1_ENTRY)EptPmlEntry2)->ReadAccess = TRUE;
                      ((PEPT_PML1_ENTRY)EptPmlEntry2)->WriteAccess = TRUE;
                    }
                  } else {
                    LogInfo("null address");
                  }
                  if (Pde->Fields.LargePage) {
                    continue;
                  }
                  PtVa = (UINT64 *)PhysicalAddressToVirtualAddress(
                      Pde->Fields.PageFrameNumber << 12);
                  if (PtVa != NULL) {
                    for (size_t l = 0; l < 512; l++) {
                    }
                  }
                }
              }
            }
          }
        }
      }
    }
  }
  SwitchToPreviousProcess(CurrentProcessCr3);
  return TRUE;
}

BOOLEAN ExecTrapAllocateUserDisabledMbecEptPageTable() {
  PVMM_EPT_PAGE_TABLE ModeBasedEptTable;
  EPT_POINTER EPTP = {0};
  ModeBasedEptTable = EptAllocateAndCreateIdentityPageTable();
  if (ModeBasedEptTable == NULL) {
    return FALSE;
  }
  ModeBasedExecHookDisableUserModeExecution(ModeBasedEptTable);
  g_EptState->ModeBasedUserDisabledEptPageTable = ModeBasedEptTable;
  EPTP.MemoryType = MEMORY_TYPE_WRITE_BACK;
  EPTP.EnableAccessAndDirtyFlags = TRUE;
  EPTP.PageWalkLength = 3;
  EPTP.PageFrameNumber =
      (SIZE_T)VirtualAddressToPhysicalAddress(&ModeBasedEptTable->PML4) /
      PAGE_SIZE;
  g_EptState->ModeBasedUserDisabledEptPointer.AsUInt = EPTP.AsUInt;
  return TRUE;
}

BOOLEAN ExecTrapAllocateKernelDisabledMbecEptPageTable() {
  PVMM_EPT_PAGE_TABLE ModeBasedEptTable;
  EPT_POINTER EPTP = {0};
  ModeBasedEptTable = EptAllocateAndCreateIdentityPageTable();
  if (ModeBasedEptTable == NULL) {
    return FALSE;
  }
  ModeBasedExecHookDisableKernelModeExecution(ModeBasedEptTable);
  g_EptState->ModeBasedKernelDisabledEptPageTable = ModeBasedEptTable;
  EPTP.MemoryType = MEMORY_TYPE_WRITE_BACK;
  EPTP.EnableAccessAndDirtyFlags = TRUE;
  EPTP.PageWalkLength = 3;
  EPTP.PageFrameNumber =
      (SIZE_T)VirtualAddressToPhysicalAddress(&ModeBasedEptTable->PML4) /
      PAGE_SIZE;
  g_EptState->ModeBasedKernelDisabledEptPointer.AsUInt = EPTP.AsUInt;
  return TRUE;
}

BOOLEAN ExecTrapEnableExecuteOnlyPages(PVMM_EPT_PAGE_TABLE EptTable) {
  INT64 RemainingSize = 0;
  UINT64 CurrentAddress = 0;
  for (size_t i = 0; i < VMM_EPT_PML4E_COUNT; i++) {
    EptTable->PML4[i].UserModeExecute = TRUE;
  }
  for (size_t i = 0; i < VMM_EPT_PML3E_COUNT; i++) {
    EptTable->PML3[i].UserModeExecute = TRUE;
  }
  for (size_t i = 0; i < VMM_EPT_PML3E_COUNT; i++) {
    for (size_t j = 0; j < VMM_EPT_PML2E_COUNT; j++) {
      EptTable->PML2[i][j].UserModeExecute = TRUE;
    }
  }
  for (size_t i = 0; i < MAX_PHYSICAL_RAM_RANGE_COUNT; i++) {
    if (PhysicalRamRegions[i].RamPhysicalAddress != NULL64_ZERO) {
      RemainingSize = (INT64)PhysicalRamRegions[i].RamSize;
      CurrentAddress = PhysicalRamRegions[i].RamPhysicalAddress;
      while (RemainingSize > 0) {
        PEPT_PML2_ENTRY EptEntry = EptGetPml2Entry(EptTable, CurrentAddress);
        EptEntry->WriteAccess = FALSE;
        CurrentAddress += SIZE_2_MB;
        RemainingSize -= SIZE_2_MB;
      }
    }
  }
  return TRUE;
}

VOID ExecTrapReadRamPhysicalRegions() {
  PHYSICAL_ADDRESS Address;
  LONGLONG Size;
  UINT32 Count = 0;
  PPHYSICAL_MEMORY_RANGE PhysicalMemoryRanges = NULL;
  PhysicalMemoryRanges = MmGetPhysicalMemoryRanges();
  do {
    Address.QuadPart = PhysicalMemoryRanges[Count].BaseAddress.QuadPart;
    Size = PhysicalMemoryRanges[Count].NumberOfBytes.QuadPart;
    if (!Address.QuadPart && !Size) {
      break;
    }
    PhysicalRamRegions[Count].RamPhysicalAddress = Address.QuadPart;
    PhysicalRamRegions[Count].RamSize = Size;
  } while (++Count < MAX_PHYSICAL_RAM_RANGE_COUNT);
  ExFreePool(PhysicalMemoryRanges);
}

BOOLEAN ExecTrapInitialize() {
  if (g_ExecTrapInitialized) {
    return FALSE;
  }
  if (!g_CompatibilityCheck.ModeBasedExecutionSupport) {
    LogInfo(
        "Your processor doesn't support Mode-Based Execution Controls (MBEC), "
        "which is a needed feature for this functionality :(\n"
        "MBEC is available on processors starting from the 7th generation "
        "(Kaby Lake) and onwards");
    return FALSE;
  }
  ExecTrapReadRamPhysicalRegions();
  if (!ExecTrapAllocateUserDisabledMbecEptPageTable()) {
    return FALSE;
  }
  if (!ExecTrapAllocateKernelDisabledMbecEptPageTable()) {
    MmFreeContiguousMemory(g_EptState->ModeBasedUserDisabledEptPageTable);
    g_EptState->ModeBasedUserDisabledEptPageTable = NULL;
    return FALSE;
  }
  if (ModeBasedExecHookInitialize() == FALSE) {
    MmFreeContiguousMemory(g_EptState->ModeBasedUserDisabledEptPageTable);
    g_EptState->ModeBasedUserDisabledEptPageTable = NULL;
    MmFreeContiguousMemory(g_EptState->ModeBasedKernelDisabledEptPageTable);
    g_EptState->ModeBasedKernelDisabledEptPageTable = NULL;
    return FALSE;
  }
  BroadcastChangeToMbecSupportedEptpOnAllProcessors();
  g_ExecTrapInitialized = TRUE;
  BroadcastEnableMovToCr3ExitingOnAllProcessors();
  return TRUE;
}

VOID ExecTrapUninitialize() {
  if (!g_ExecTrapInitialized) {
    return;
  }
  g_ExecTrapUnInitializationStarted = TRUE;
  BroadcastDisableMovToCr3ExitingOnAllProcessors();
  BroadcastRestoreToNormalEptpOnAllProcessors();
  ModeBasedExecHookUninitialize();
  g_ExecTrapInitialized = FALSE;
  g_ExecTrapUnInitializationStarted = FALSE;
  if (g_EptState->ModeBasedUserDisabledEptPageTable != NULL) {
    MmFreeContiguousMemory(g_EptState->ModeBasedUserDisabledEptPageTable);
    g_EptState->ModeBasedUserDisabledEptPageTable = NULL;
  }
  if (g_EptState->ModeBasedKernelDisabledEptPageTable != NULL) {
    MmFreeContiguousMemory(g_EptState->ModeBasedKernelDisabledEptPageTable);
    g_EptState->ModeBasedKernelDisabledEptPageTable = NULL;
  }
}

VOID ExecTrapRestoreToNormalEptp(VIRTUAL_MACHINE_STATE *VCpu) {
  __vmx_vmwrite(VMCS_CTRL_EPT_POINTER, VCpu->EptPointer.AsUInt);
  VCpu->NotNormalEptp = FALSE;
}

VOID ExecTrapChangeToExecuteOnlyEptp(VIRTUAL_MACHINE_STATE *VCpu) {
  __vmx_vmwrite(VMCS_CTRL_EPT_POINTER,
                g_EptState->ExecuteOnlyEptPointer.AsUInt);
  VCpu->NotNormalEptp = TRUE;
}

VOID ExecTrapChangeToUserDisabledMbecEptp(VIRTUAL_MACHINE_STATE *VCpu) {
  __vmx_vmwrite(VMCS_CTRL_EPT_POINTER,
                g_EptState->ModeBasedUserDisabledEptPointer.AsUInt);
  VCpu->NotNormalEptp = TRUE;
}

VOID ExecTrapChangeToKernelDisabledMbecEptp(VIRTUAL_MACHINE_STATE *VCpu) {
  __vmx_vmwrite(VMCS_CTRL_EPT_POINTER,
                g_EptState->ModeBasedKernelDisabledEptPointer.AsUInt);
  VCpu->NotNormalEptp = TRUE;
}

VOID ExecTrapHandleMoveToAdjustedTrapState(
    VIRTUAL_MACHINE_STATE *VCpu, DEBUGGER_EVENT_MODE_TYPE TargetMode) {
  if (TargetMode == DEBUGGER_EVENT_MODE_TYPE_USER_MODE) {
    ExecTrapChangeToKernelDisabledMbecEptp(VCpu);
  } else if (TargetMode == DEBUGGER_EVENT_MODE_TYPE_KERNEL_MODE) {
    ExecTrapChangeToUserDisabledMbecEptp(VCpu);
  }
}

BOOLEAN ExecTrapHandleEptViolationVmexit(
    VIRTUAL_MACHINE_STATE *VCpu,
    VMX_EXIT_QUALIFICATION_EPT_VIOLATION *ViolationQualification) {
  if (!g_ExecTrapInitialized) {
    return FALSE;
  }
  if (!ViolationQualification->EptExecutableForUserMode &&
      ViolationQualification->ExecuteAccess) {
    HvSuppressRipIncrement(VCpu);
    DispatchEventMode(VCpu, DEBUGGER_EVENT_MODE_TYPE_USER_MODE, TRUE);
    return TRUE;
  } else if (!ViolationQualification->EptExecutable &&
             ViolationQualification->ExecuteAccess) {
    HvSuppressRipIncrement(VCpu);
    DispatchEventMode(VCpu, DEBUGGER_EVENT_MODE_TYPE_KERNEL_MODE, TRUE);
  } else {
    return FALSE;
  }
  return TRUE;
}

VOID ExecTrapHandleCr3Vmexit(VIRTUAL_MACHINE_STATE *VCpu) {
  BOOLEAN Result;
  UINT32 Index;
  Result = BinarySearchPerformSearchItem(
      &g_ExecTrapState.InterceptionProcessIds[0], g_ExecTrapState.NumberOfItems,
      &Index, (UINT64)PsGetCurrentProcessId());
  if (Result) {
    HvSetModeBasedExecutionEnableFlag(TRUE);
    VCpu->MbecEnabled = TRUE;
    DispatchEventMode(VCpu, DEBUGGER_EVENT_MODE_TYPE_KERNEL_MODE, TRUE);
  } else if (VCpu->MbecEnabled) {
    HvSetModeBasedExecutionEnableFlag(FALSE);
    VCpu->MbecEnabled = FALSE;
  }
}

BOOLEAN ExecTrapAddProcessToWatchingList(UINT32 ProcessId) {
  UINT32 Index;
  return InsertionSortInsertItem(
      &g_ExecTrapState.InterceptionProcessIds[0],
      &g_ExecTrapState.NumberOfItems,
      MAXIMUM_NUMBER_OF_PROCESSES_FOR_USER_KERNEL_EXEC_THREAD, &Index,
      (UINT64)ProcessId);
}

BOOLEAN ExecTrapRemoveProcessFromWatchingList(UINT32 ProcessId) {
  return InsertionSortDeleteItem(&g_ExecTrapState.InterceptionProcessIds[0],
                                 &g_ExecTrapState.NumberOfItems,
                                 (UINT64)ProcessId);
}
