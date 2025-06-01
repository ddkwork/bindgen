
#include "pch.h"

VOID MtfHandleVmexit(VIRTUAL_MACHINE_STATE *VCpu) {
  BOOLEAN IsMtfHandled = FALSE;
  HvSuppressRipIncrement(VCpu);
  VCpu->IgnoreMtfUnset = FALSE;
  if (g_Callbacks.BreakpointCheckAndHandleReApplyingBreakpoint != NULL &&
      g_Callbacks.BreakpointCheckAndHandleReApplyingBreakpoint(VCpu->CoreId)) {
    IsMtfHandled = TRUE;
  }
  if (VCpu->MtfEptHookRestorePoint) {
    IsMtfHandled = TRUE;
    EptHookHandleMonitorTrapFlag(VCpu);
    VCpu->MtfEptHookRestorePoint = NULL;
    HvEnableAndCheckForPreviousExternalInterrupts(VCpu);
  }
  if (VCpu->RegisterBreakOnMtf) {
    IsMtfHandled = TRUE;
    VCpu->RegisterBreakOnMtf = FALSE;
    VmmCallbackRegisteredMtfHandler(VCpu->CoreId);
  } else if (g_Callbacks.KdCheckAndHandleNmiCallback != NULL &&
             g_Callbacks.KdCheckAndHandleNmiCallback(VCpu->CoreId)) {
    IsMtfHandled = TRUE;
  }
  if (VCpu->IgnoreOneMtf) {
    IsMtfHandled = TRUE;
    VCpu->IgnoreOneMtf = FALSE;
  }
  if (!IsMtfHandled) {
    LogError("Err, why MTF occurred?!");
  }
  if (!VCpu->IgnoreMtfUnset) {
    HvSetMonitorTrapFlag(FALSE);
  } else {
    VCpu->IgnoreMtfUnset = FALSE;
  }
}
