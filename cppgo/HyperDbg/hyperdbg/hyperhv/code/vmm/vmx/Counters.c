
#include "pch.h"

VOID CounterEmulateRdtsc(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT64 Tsc = __rdtsc();
  PGUEST_REGS GuestRegs = VCpu->Regs;
  GuestRegs->rax = 0x00000000ffffffff & Tsc;
  GuestRegs->rdx = 0x00000000ffffffff & (Tsc >> 32);
}

VOID CounterEmulateRdtscp(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 Aux = 0;
  UINT64 Tsc = __rdtscp(&Aux);
  PGUEST_REGS GuestRegs = VCpu->Regs;
  GuestRegs->rax = 0x00000000ffffffff & Tsc;
  GuestRegs->rdx = 0x00000000ffffffff & (Tsc >> 32);
  GuestRegs->rcx = 0x00000000ffffffff & Aux;
}

VOID CounterEmulateRdpmc(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 EcxReg = 0;
  PGUEST_REGS GuestRegs = VCpu->Regs;
  EcxReg = GuestRegs->rcx & 0xffffffff;
  UINT64 Pmc = __readpmc(EcxReg);
  GuestRegs->rax = 0x00000000ffffffff & Pmc;
  GuestRegs->rdx = 0x00000000ffffffff & (Pmc >> 32);
}

VOID CounterSetPreemptionTimer(UINT32 TimerValue) {
  VmxVmwrite64(VMCS_GUEST_VMX_PREEMPTION_TIMER_VALUE, TimerValue);
}

VOID CounterClearPreemptionTimer() {
  VmxVmwrite64(VMCS_GUEST_VMX_PREEMPTION_TIMER_VALUE, NULL64_ZERO);
}
