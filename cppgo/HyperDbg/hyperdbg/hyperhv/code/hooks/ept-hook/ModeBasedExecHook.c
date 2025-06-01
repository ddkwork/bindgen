
#include "pch.h"

BOOLEAN ModeBasedExecHookDisableUserModeExecution(
    PVMM_EPT_PAGE_TABLE EptTable) {
  for (size_t i = 0; i < VMM_EPT_PML4E_COUNT; i++) {
    EptTable->PML4[i].UserModeExecute = FALSE;
  }
  for (size_t i = 0; i < VMM_EPT_PML3E_COUNT; i++) {
    EptTable->PML3[i].UserModeExecute = TRUE;
  }
  for (size_t i = 0; i < VMM_EPT_PML3E_COUNT; i++) {
    for (size_t j = 0; j < VMM_EPT_PML2E_COUNT; j++) {
      EptTable->PML2[i][j].UserModeExecute = TRUE;
    }
  }
  return TRUE;
}

BOOLEAN ModeBasedExecHookDisableKernelModeExecution(
    PVMM_EPT_PAGE_TABLE EptTable) {
  for (size_t i = 0; i < VMM_EPT_PML4E_COUNT; i++) {
    EptTable->PML4[i].UserModeExecute = TRUE;
    EptTable->PML4[i].ExecuteAccess = FALSE;
  }
  for (size_t i = 0; i < VMM_EPT_PML3E_COUNT; i++) {
    EptTable->PML3[i].UserModeExecute = TRUE;
  }
  for (size_t i = 0; i < VMM_EPT_PML3E_COUNT; i++) {
    for (size_t j = 0; j < VMM_EPT_PML2E_COUNT; j++) {
      EptTable->PML2[i][j].UserModeExecute = TRUE;
    }
  }
  return TRUE;
}

BOOLEAN ModeBasedExecHookEnableUsermodeExecution(PVMM_EPT_PAGE_TABLE EptTable) {
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
  return TRUE;
}

VOID ModeBasedExecHookEnableOrDisable(VIRTUAL_MACHINE_STATE *VCpu,
                                      UINT32 State) {
  if (State == 0x0) {
    HvSetModeBasedExecutionEnableFlag(FALSE);
    VCpu->MbecEnabled = FALSE;
  } else {
    HvSetModeBasedExecutionEnableFlag(TRUE);
  }
}

BOOLEAN ModeBasedExecHookInitialize() {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (!g_CompatibilityCheck.ModeBasedExecutionSupport) {
    return FALSE;
  }
  if (!g_CompatibilityCheck.ExecuteOnlySupport) {
    return FALSE;
  }
  if (g_ModeBasedExecutionControlState) {
    return FALSE;
  }
  for (size_t i = 0; i < ProcessorsCount; i++) {
    ModeBasedExecHookEnableUsermodeExecution(g_GuestState[i].EptPageTable);
  }
  BroadcastNotifyAllToInvalidateEptAllCores();
  g_ModeBasedExecutionControlState = TRUE;
  return TRUE;
}

VOID ModeBasedExecHookUninitialize() {
  BroadcasDisableMbecOnAllProcessors();
  g_ModeBasedExecutionControlState = FALSE;
}
