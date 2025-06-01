
#include "pch.h"

VOID VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer() {
  HvSetVmxPreemptionTimerExiting(TRUE);
  CounterSetPreemptionTimer(0);
}

VOID VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer() {
  CounterClearPreemptionTimer();
  HvSetVmxPreemptionTimerExiting(FALSE);
}

VOID VmxMechanismCreateImmediateVmexitBySelfIpi() {
  ApicSelfIpi(IMMEDIATE_VMEXIT_MECHANISM_VECTOR_FOR_SELF_IPI);
}

VOID VmxMechanismCreateImmediateVmexit(VIRTUAL_MACHINE_STATE *VCpu) {
  VCpu->WaitForImmediateVmexit = TRUE;
  VmxMechanismCreateImmediateVmexitBySelfIpi();
  HvSetExternalInterruptExiting(VCpu, TRUE);
}

VOID VmxMechanismHandleImmediateVmexit(VIRTUAL_MACHINE_STATE *VCpu) {
  VCpu->WaitForImmediateVmexit = FALSE;
  HvSetExternalInterruptExiting(VCpu, FALSE);
}
