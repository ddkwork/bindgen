
#include "pch.h"

VOID VmxHandleXsetbv(VIRTUAL_MACHINE_STATE *VCpu) {
  CPUID_EAX_01 CpuidInfo;
  VMX_SEGMENT_SELECTOR GuestCs = GetGuestCs();
  UINT32 XCrIndex = VCpu->Regs->rcx & 0xffffffff;
  XCR0 XCr0 = {.AsUInt = VCpu->Regs->rdx << 32 | VCpu->Regs->rax};
  CR4 GuestCr4 = {.AsUInt = GetGuestCr4()};
  if (XCrIndex != 0 || GuestCs.Attributes.DescriptorPrivilegeLevel != 0 ||
      !CommonIsXCr0Valid(XCr0)) {
    EventInjectGeneralProtection();
    return;
  }
  CommonCpuidInstruction(CPUID_VERSION_INFORMATION, 0, (INT32 *)&CpuidInfo);
  if (CpuidInfo.CpuidFeatureInformationEcx.XsaveXrstorInstruction == 0 ||
      GuestCr4.OsXsave == 0) {
    EventInjectUndefinedOpcode(VCpu);
    return;
  }
  _xsetbv(XCrIndex, XCr0.AsUInt);
}

VOID VmxHandleVmxPreemptionTimerVmexit(VIRTUAL_MACHINE_STATE *VCpu) {
  LogError("Why vm-exit for VMX preemption timer happened?");
  HvSuppressRipIncrement(VCpu);
}

VOID VmxHandleTripleFaults(VIRTUAL_MACHINE_STATE *VCpu) {
  LogError("Err, triple fault error occurred!");
  CommonWriteDebugInformation(VCpu);
  DbgBreakPoint();
}
