
#include "pch.h"

#define MAXIMUM_NUMBER_OF_PROCESSES_FOR_USER_KERNEL_EXEC_THREAD 100

typedef struct _USER_KERNEL_EXECUTION_TRAP_STATE {
  UINT32 NumberOfItems;
  UINT64 InterceptionProcessIds
      [MAXIMUM_NUMBER_OF_PROCESSES_FOR_USER_KERNEL_EXEC_THREAD];
} USER_KERNEL_EXECUTION_TRAP_STATE, *PUSER_KERNEL_EXECUTION_TRAP_STATE;

VOID ExecTrapHandleCr3Vmexit(VIRTUAL_MACHINE_STATE *VCpu);
VOID ExecTrapChangeToUserDisabledMbecEptp(VIRTUAL_MACHINE_STATE *VCpu);
VOID ExecTrapChangeToKernelDisabledMbecEptp(VIRTUAL_MACHINE_STATE *VCpu);
VOID ExecTrapRestoreToNormalEptp(VIRTUAL_MACHINE_STATE *VCpu);
VOID ExecTrapHandleMoveToAdjustedTrapState(VIRTUAL_MACHINE_STATE *VCpu,
                                           DEBUGGER_EVENT_MODE_TYPE TargetMode);
VOID ExecTrapUninitialize();
BOOLEAN ExecTrapInitialize();
BOOLEAN ExecTrapHandleEptViolationVmexit(
    VIRTUAL_MACHINE_STATE *VCpu,
    VMX_EXIT_QUALIFICATION_EPT_VIOLATION *ViolationQualification);
BOOLEAN ExecTrapAddProcessToWatchingList(UINT32 ProcessId);
BOOLEAN ExecTrapRemoveProcessFromWatchingList(UINT32 ProcessId);
