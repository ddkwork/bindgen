
#pragma once
#define MAX_PHYSICAL_RAM_RANGE_COUNT 32

typedef struct MODE_BASED_RAM_REGIONS {
  UINT64 RamPhysicalAddress;
  UINT64 RamSize;
} MODE_BASED_RAM_REGIONS, *PMODE_BASED_RAM_REGIONS;

MODE_BASED_RAM_REGIONS PhysicalRamRegions[MAX_PHYSICAL_RAM_RANGE_COUNT];
VOID ModeBasedExecHookEnableOrDisable(VIRTUAL_MACHINE_STATE *VCpu,
                                      UINT32 State);
VOID ModeBasedExecHookUninitialize();
BOOLEAN ModeBasedExecHookInitialize();
BOOLEAN ModeBasedExecHookDisableUserModeExecution(PVMM_EPT_PAGE_TABLE EptTable);
BOOLEAN ModeBasedExecHookDisableKernelModeExecution(
    PVMM_EPT_PAGE_TABLE EptTable);
