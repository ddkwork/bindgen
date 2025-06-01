
#pragma once
#define PML_ENTITY_NUM 512
BOOLEAN DirtyLoggingInitialize();
BOOLEAN DirtyLoggingEnable(VIRTUAL_MACHINE_STATE *VCpu);
VOID DirtyLoggingDisable(VIRTUAL_MACHINE_STATE *VCpu);
VOID DirtyLoggingUninitialize();
VOID DirtyLoggingHandleVmexits(VIRTUAL_MACHINE_STATE *VCpu);
