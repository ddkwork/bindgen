
#include "pch.h"

BOOLEAN CompatibilityCheckCpuSupportForRtm() {
  int Regs1[4];
  int Regs2[4];
  BOOLEAN Result;
  CommonCpuidInstruction(0, 0, Regs1);
  CommonCpuidInstruction(7, 0, Regs2);
  Result = Regs1[0] >= 0x7 && (Regs2[1] & 0x4800) == 0x4800;
  return Result;
}

UINT32 CompatibilityCheckGetX86VirtualAddressWidth() {
  int Regs[4];
  CommonCpuidInstruction(CPUID_ADDR_WIDTH, 0, Regs);
  return ((Regs[0] >> 8) & 0x0ff);
}

UINT32 CompatibilityCheckGetX86PhysicalAddressWidth() {
  int Regs[4];
  CommonCpuidInstruction(CPUID_ADDR_WIDTH, 0, Regs);
  return (Regs[0] & 0x0ff);
}

BOOLEAN CompatibilityCheckModeBasedExecution() {
  UINT32 SecondaryProcBasedVmExecControls = HvAdjustControls(
      IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_FLAG,
      IA32_VMX_PROCBASED_CTLS2);
  if (SecondaryProcBasedVmExecControls &
      IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_FLAG) {
    return TRUE;
  } else {
    return FALSE;
  }
}

BOOLEAN CompatibilityCheckPml() {
  UINT32 SecondaryProcBasedVmExecControls = HvAdjustControls(
      IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_FLAG, IA32_VMX_PROCBASED_CTLS2);
  if (SecondaryProcBasedVmExecControls &
      IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_FLAG) {
    return TRUE;
  } else {
    return FALSE;
  }
}

VOID CompatibilityCheckPerformChecks() {
  g_CompatibilityCheck.RtmSupport = CompatibilityCheckCpuSupportForRtm();
  g_CompatibilityCheck.VirtualAddressWidth =
      CompatibilityCheckGetX86VirtualAddressWidth();
  g_CompatibilityCheck.PhysicalAddressWidth =
      CompatibilityCheckGetX86PhysicalAddressWidth();
  g_CompatibilityCheck.ModeBasedExecutionSupport =
      CompatibilityCheckModeBasedExecution();
  g_CompatibilityCheck.PmlSupport = CompatibilityCheckPml();
  LogDebugInfo(
      "Mode based execution: %s | PML: %s",
      g_CompatibilityCheck.ModeBasedExecutionSupport ? "true" : "false",
      g_CompatibilityCheck.PmlSupport ? "true" : "false");
}
