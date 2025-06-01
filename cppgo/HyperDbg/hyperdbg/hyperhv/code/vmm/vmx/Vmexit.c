
#include "pch.h"

BOOLEAN VmxVmexitHandler(_Inout_ PGUEST_REGS GuestRegs) {
  UINT32 ExitReason = 0;
  BOOLEAN Result = FALSE;
  VIRTUAL_MACHINE_STATE *VCpu = NULL;
  VCpu = &g_GuestState[KeGetCurrentProcessorNumberEx(NULL)];
  VCpu->Regs = GuestRegs;
  VCpu->IsOnVmxRootMode = TRUE;
  VmxVmread32P(VMCS_EXIT_REASON, &ExitReason);
  ExitReason &= 0xffff;
  VCpu->ExitReason = ExitReason;
  VCpu->IncrementRip = TRUE;
  __vmx_vmread(VMCS_GUEST_RIP, &VCpu->LastVmexitRip);
  __vmx_vmread(VMCS_GUEST_RSP, &VCpu->Regs->rsp);
  VmxVmread32P(VMCS_EXIT_QUALIFICATION, &VCpu->ExitQualification);
  switch (ExitReason) {
    case VMX_EXIT_REASON_TRIPLE_FAULT: {
      VmxHandleTripleFaults(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_VMCLEAR:
    case VMX_EXIT_REASON_EXECUTE_VMPTRLD:
    case VMX_EXIT_REASON_EXECUTE_VMPTRST:
    case VMX_EXIT_REASON_EXECUTE_VMREAD:
    case VMX_EXIT_REASON_EXECUTE_VMRESUME:
    case VMX_EXIT_REASON_EXECUTE_VMWRITE:
    case VMX_EXIT_REASON_EXECUTE_VMXOFF:
    case VMX_EXIT_REASON_EXECUTE_VMXON:
    case VMX_EXIT_REASON_EXECUTE_VMLAUNCH: {
      EventInjectUndefinedOpcode(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_INVEPT:
    case VMX_EXIT_REASON_EXECUTE_INVVPID:
    case VMX_EXIT_REASON_EXECUTE_GETSEC:
    case VMX_EXIT_REASON_EXECUTE_INVD: {
      EventInjectUndefinedOpcode(VCpu);
      break;
    }
    case VMX_EXIT_REASON_MOV_CR: {
      DispatchEventMovToFromControlRegisters(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_RDMSR: {
      DispatchEventRdmsr(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_WRMSR: {
      DispatchEventWrmsr(VCpu);
      break;
    }
    case VMX_EXIT_REASON_IO_SMI:
    case VMX_EXIT_REASON_SMI: {
      LogInfo("VM-exit reason SMM %llx | qual: %llx", ExitReason,
              VCpu->ExitQualification);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_CPUID: {
      DispatchEventCpuid(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_IO_INSTRUCTION: {
      DispatchEventIO(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EPT_VIOLATION: {
      if (EptHandleEptViolation(VCpu) == FALSE) {
        LogError("Err, there were errors in handling EPT violation");
      }
      break;
    }
    case VMX_EXIT_REASON_EPT_MISCONFIGURATION: {
      EptHandleMisconfiguration();
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_VMCALL: {
      DispatchEventVmcall(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXCEPTION_OR_NMI: {
      DispatchEventException(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXTERNAL_INTERRUPT: {
      DispatchEventExternalInterrupts(VCpu);
      break;
    }
    case VMX_EXIT_REASON_INTERRUPT_WINDOW: {
      IdtEmulationHandleInterruptWindowExiting(VCpu);
      break;
    }
    case VMX_EXIT_REASON_NMI_WINDOW: {
      IdtEmulationHandleNmiWindowExiting(VCpu);
      break;
    }
    case VMX_EXIT_REASON_MONITOR_TRAP_FLAG: {
      MtfHandleVmexit(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_HLT: {
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_RDTSC:
    case VMX_EXIT_REASON_EXECUTE_RDTSCP: {
      DispatchEventTsc(
          VCpu, ExitReason == VMX_EXIT_REASON_EXECUTE_RDTSCP ? TRUE : FALSE);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_RDPMC: {
      DispatchEventRdpmc(VCpu);
      break;
    }
    case VMX_EXIT_REASON_MOV_DR: {
      DispatchEventMov2DebugRegs(VCpu);
      break;
    }
    case VMX_EXIT_REASON_EXECUTE_XSETBV: {
      VmxHandleXsetbv(VCpu);
      break;
    }
    case VMX_EXIT_REASON_VMX_PREEMPTION_TIMER_EXPIRED: {
      VmxHandleVmxPreemptionTimerVmexit(VCpu);
      break;
    }
    case VMX_EXIT_REASON_PAGE_MODIFICATION_LOG_FULL: {
      DirtyLoggingHandleVmexits(VCpu);
      break;
    }
    default: {
      LogError("Err, unknown vmexit, reason : 0x%llx", ExitReason);
      break;
    }
  }
  if (!VCpu->VmxoffState.IsVmxoffExecuted && VCpu->IncrementRip) {
    if (g_TransparentMode) {
      HvHandleTrapFlag();
    }
    HvResumeToNextInstruction();
  }
  if (VCpu->VmxoffState.IsVmxoffExecuted) {
    Result = TRUE;
  }
  VCpu->IsOnVmxRootMode = FALSE;
  return Result;
}
