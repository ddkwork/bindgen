
#include "pch.h"

VOID DispatchEventEferSysret(VIRTUAL_MACHINE_STATE *VCpu, PVOID Context) {
  BOOLEAN PostEventTriggerReq = FALSE;
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  EventTriggerResult = VmmCallbackTriggerEvents(
      SYSCALL_HOOK_EFER_SYSRET, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    SyscallHookEmulateSYSRET(VCpu);
    HvSuppressRipIncrement(VCpu);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(SYSCALL_HOOK_EFER_SYSRET,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             Context, NULL, VCpu->Regs);
  }
}

VOID DispatchEventEferSyscall(VIRTUAL_MACHINE_STATE *VCpu) {
  BOOLEAN PostEventTriggerReq = FALSE;
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  EventTriggerResult = VmmCallbackTriggerEvents(
      SYSCALL_HOOK_EFER_SYSCALL, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      (PVOID)VCpu->Regs->rax, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    SyscallHookEmulateSYSCALL(VCpu);
    HvSuppressRipIncrement(VCpu);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(SYSCALL_HOOK_EFER_SYSCALL,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             (PVOID)VCpu->Regs->rax, NULL, VCpu->Regs);
  }
}

VOID DispatchEventCpuid(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT64 Context;
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  if (g_Callbacks.UdCheckForCommand != NULL &&
      g_Callbacks.UdCheckForCommand()) {
    return;
  }
  if (g_TriggerEventForCpuids) {
    Context = VCpu->Regs->rax & 0xffffffff;
    EventTriggerResult = VmmCallbackTriggerEvents(
        CPUID_INSTRUCTION_EXECUTION,
        VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION, (PVOID)Context,
        &PostEventTriggerReq, VCpu->Regs);
    if (EventTriggerResult !=
        VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
      HvHandleCpuid(VCpu);
    }
    if (PostEventTriggerReq) {
      VmmCallbackTriggerEvents(CPUID_INSTRUCTION_EXECUTION,
                               VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                               (PVOID)Context, NULL, VCpu->Regs);
    }
  } else {
    HvHandleCpuid(VCpu);
  }
}

VOID DispatchEventTsc(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN IsRdtscp) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  EventTriggerResult = VmmCallbackTriggerEvents(
      TSC_INSTRUCTION_EXECUTION, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      (PVOID)IsRdtscp, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    if (IsRdtscp) {
      CounterEmulateRdtscp(VCpu);
    } else {
      CounterEmulateRdtsc(VCpu);
    }
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(TSC_INSTRUCTION_EXECUTION,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             (PVOID)IsRdtscp, NULL, VCpu->Regs);
  }
}

VOID DispatchEventVmcall(VIRTUAL_MACHINE_STATE *VCpu) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  if (g_TriggerEventForVmcalls) {
    EventTriggerResult =
        VmmCallbackTriggerEvents(VMCALL_INSTRUCTION_EXECUTION,
                                 VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                                 NULL, &PostEventTriggerReq, VCpu->Regs);
    if (EventTriggerResult !=
        VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
      VmxHandleVmcallVmExit(VCpu);
    }
    if (PostEventTriggerReq) {
      VmmCallbackTriggerEvents(CPUID_INSTRUCTION_EXECUTION,
                               VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                               NULL, NULL, VCpu->Regs);
    }
  } else {
    VmxHandleVmcallVmExit(VCpu);
  }
}

VOID DispatchEventMode(VIRTUAL_MACHINE_STATE *VCpu,
                       DEBUGGER_EVENT_MODE_TYPE TargetMode,
                       BOOLEAN HandleState) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  if (g_ExecTrapInitialized) {
    EventTriggerResult = VmmCallbackTriggerEvents(
        TRAP_EXECUTION_MODE_CHANGED,
        VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION, (PVOID)TargetMode,
        &PostEventTriggerReq, VCpu->Regs);
    if (EventTriggerResult !=
            VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT &&
        HandleState) {
      ExecTrapHandleMoveToAdjustedTrapState(VCpu, (UINT64)TargetMode);
    }
  } else {
    if (HandleState) {
      ExecTrapHandleMoveToAdjustedTrapState(VCpu, (UINT64)TargetMode);
    }
  }
}

VOID DispatchEventMovToCr3(VIRTUAL_MACHINE_STATE *VCpu) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  if (g_ExecTrapInitialized) {
    EventTriggerResult =
        VmmCallbackTriggerEvents(CONTROL_REGISTER_3_MODIFIED,
                                 VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                                 NULL, &PostEventTriggerReq, VCpu->Regs);
    if (EventTriggerResult !=
        VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    }
    if (PostEventTriggerReq) {
      VmmCallbackTriggerEvents(CONTROL_REGISTER_3_MODIFIED,
                               VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                               NULL, NULL, VCpu->Regs);
    }
  } else {
  }
}

VOID DispatchEventIO(VIRTUAL_MACHINE_STATE *VCpu) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult =
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_NO_INITIALIZED;
  VMX_EXIT_QUALIFICATION_IO_INSTRUCTION IoQualification = {
      .AsUInt = VCpu->ExitQualification};
  RFLAGS Flags = {0};
  BOOLEAN PostEventTriggerReq = FALSE;
  VmxVmread64P(VMCS_GUEST_RFLAGS, (UINT64 *)&Flags);
  if (IoQualification.DirectionOfAccess == AccessIn) {
    EventTriggerResult = VmmCallbackTriggerEvents(
        IN_INSTRUCTION_EXECUTION,
        VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
        (PVOID)IoQualification.PortNumber, &PostEventTriggerReq, VCpu->Regs);
  } else if (IoQualification.DirectionOfAccess == AccessOut) {
    EventTriggerResult = VmmCallbackTriggerEvents(
        OUT_INSTRUCTION_EXECUTION,
        VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
        (PVOID)IoQualification.PortNumber, &PostEventTriggerReq, VCpu->Regs);
  }
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    IoHandleIoVmExits(VCpu, IoQualification, Flags);
  }
  if (PostEventTriggerReq) {
    if (IoQualification.DirectionOfAccess == AccessIn) {
      VmmCallbackTriggerEvents(IN_INSTRUCTION_EXECUTION,
                               VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                               (PVOID)IoQualification.PortNumber, NULL,
                               VCpu->Regs);
    } else if (IoQualification.DirectionOfAccess == AccessOut) {
      VmmCallbackTriggerEvents(OUT_INSTRUCTION_EXECUTION,
                               VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                               (PVOID)IoQualification.PortNumber, NULL,
                               VCpu->Regs);
    }
  }
}

VOID DispatchEventRdmsr(VIRTUAL_MACHINE_STATE *VCpu) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  EventTriggerResult = VmmCallbackTriggerEvents(
      RDMSR_INSTRUCTION_EXECUTION,
      VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      (PVOID)(VCpu->Regs->rcx & 0xffffffff), &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    MsrHandleRdmsrVmexit(VCpu);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(RDMSR_INSTRUCTION_EXECUTION,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             (PVOID)(VCpu->Regs->rcx & 0xffffffff), NULL,
                             VCpu->Regs);
  }
}

VOID DispatchEventWrmsr(VIRTUAL_MACHINE_STATE *VCpu) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  EventTriggerResult = VmmCallbackTriggerEvents(
      WRMSR_INSTRUCTION_EXECUTION,
      VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      (PVOID)(VCpu->Regs->rcx & 0xffffffff), &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    MsrHandleWrmsrVmexit(VCpu);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(WRMSR_INSTRUCTION_EXECUTION,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             (PVOID)(VCpu->Regs->rcx & 0xffffffff), NULL,
                             VCpu->Regs);
  }
}

VOID DispatchEventRdpmc(VIRTUAL_MACHINE_STATE *VCpu) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  EventTriggerResult = VmmCallbackTriggerEvents(
      PMC_INSTRUCTION_EXECUTION, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      NULL, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    CounterEmulateRdpmc(VCpu);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(PMC_INSTRUCTION_EXECUTION,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             NULL, NULL, VCpu->Regs);
  }
}

VOID DispatchEventMov2DebugRegs(VIRTUAL_MACHINE_STATE *VCpu) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  if (g_Callbacks.KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId !=
          NULL &&
      g_Callbacks.KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId(
          VCpu->CoreId,
          DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_DEBUG_REGISTER_INTERCEPTION)) {
    return;
  }
  EventTriggerResult = VmmCallbackTriggerEvents(
      DEBUG_REGISTERS_ACCESSED, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      NULL, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    HvHandleMovDebugRegister(VCpu);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(DEBUG_REGISTERS_ACCESSED,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             NULL, NULL, VCpu->Regs);
  }
}

VOID DispatchEventMovToFromControlRegisters(VIRTUAL_MACHINE_STATE *VCpu) {
  BOOLEAN ModifyReg;
  VMX_EXIT_QUALIFICATION_MOV_CR *CrExitQualification;
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  UINT32 ExitQualification = 0;
  VmxVmread32P(VMCS_EXIT_QUALIFICATION, &ExitQualification);
  CrExitQualification = (VMX_EXIT_QUALIFICATION_MOV_CR *)&ExitQualification;
  if (CrExitQualification->AccessType ==
      VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR) {
    ModifyReg = TRUE;
  } else {
    ModifyReg = FALSE;
  }
  EventTriggerResult = VmmCallbackTriggerEvents(
      ModifyReg ? CONTROL_REGISTER_MODIFIED : CONTROL_REGISTER_READ,
      VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      (PVOID)CrExitQualification->ControlRegister, &PostEventTriggerReq,
      VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    HvHandleControlRegisterAccess(VCpu, CrExitQualification);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(
        ModifyReg ? CONTROL_REGISTER_MODIFIED : CONTROL_REGISTER_READ,
        VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
        (PVOID)CrExitQualification->ControlRegister, NULL, VCpu->Regs);
  }
}

VOID DispatchEventException(VIRTUAL_MACHINE_STATE *VCpu) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  VMEXIT_INTERRUPT_INFORMATION InterruptExit = {0};
  VmxVmread32P(VMCS_VMEXIT_INTERRUPTION_INFORMATION, &InterruptExit.AsUInt);
  if (InterruptExit.InterruptionType == INTERRUPT_TYPE_NMI &&
      InterruptExit.Vector == EXCEPTION_VECTOR_NMI) {
    if (!VCpu->RegisterBreakOnMtf && VmxBroadcastNmiHandler(VCpu, FALSE)) {
      return;
    }
  }
  EventTriggerResult = VmmCallbackTriggerEvents(
      EXCEPTION_OCCURRED, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      (PVOID)InterruptExit.Vector, &PostEventTriggerReq, VCpu->Regs);
  if (VCpu->RegisterBreakOnMtf) {
    return;
  }
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    IdtEmulationHandleExceptionAndNmi(VCpu, InterruptExit);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(EXCEPTION_OCCURRED,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             (PVOID)InterruptExit.Vector, NULL, VCpu->Regs);
  }
}

VOID DispatchEventExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu) {
  VMEXIT_INTERRUPT_INFORMATION InterruptExit = {0};
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  VmxVmread32P(VMCS_VMEXIT_INTERRUPTION_INFORMATION, &InterruptExit.AsUInt);
  if (VCpu->WaitForImmediateVmexit &&
      InterruptExit.Vector == IMMEDIATE_VMEXIT_MECHANISM_VECTOR_FOR_SELF_IPI) {
    HvSetExternalInterruptExiting(VCpu, FALSE);
    HvSuppressRipIncrement(VCpu);
    VmxMechanismHandleImmediateVmexit(VCpu);
    return;
  }
  if ((InterruptExit.Vector == CLOCK_INTERRUPT) ||
      (VCpu->CoreId != 0 && InterruptExit.Vector == IPI_INTERRUPT)) {
    if (g_Callbacks.DebuggerCheckProcessOrThreadChange != NULL) {
      g_Callbacks.DebuggerCheckProcessOrThreadChange(VCpu->CoreId);
    }
  }
  EventTriggerResult = VmmCallbackTriggerEvents(
      EXTERNAL_INTERRUPT_OCCURRED,
      VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      (PVOID)InterruptExit.Vector, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult !=
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    IdtEmulationHandleExternalInterrupt(VCpu, InterruptExit);
  }
  if (PostEventTriggerReq) {
    VmmCallbackTriggerEvents(EXTERNAL_INTERRUPT_OCCURRED,
                             VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                             (PVOID)InterruptExit.Vector, NULL, VCpu->Regs);
  }
}

VOID DispatchEventHiddenHookExecCc(VIRTUAL_MACHINE_STATE *VCpu, PVOID Context) {
  BOOLEAN PostEventTriggerReq = FALSE;
  VmmCallbackTriggerEvents(HIDDEN_HOOK_EXEC_CC,
                           VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                           Context, &PostEventTriggerReq, VCpu->Regs);
}

VOID DispatchEventHiddenHookExecDetours(VIRTUAL_MACHINE_STATE *VCpu,
                                        PVOID Context) {
  BOOLEAN PostEventTriggerReq = FALSE;
  VmmCallbackTriggerEvents(HIDDEN_HOOK_EXEC_DETOURS,
                           VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                           Context, &PostEventTriggerReq, VCpu->Regs);
}

BOOLEAN DispatchEventHiddenHookPageReadWriteExecuteReadPreEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context,
    BOOLEAN *IsTriggeringPostEventAllowed) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  BOOLEAN ShortCircuitingEvent = FALSE;
  EventTriggerResult = VmmCallbackTriggerEvents(
      HIDDEN_HOOK_READ, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION, Context,
      &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_EXECUTE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  return ShortCircuitingEvent;
}

BOOLEAN DispatchEventHiddenHookPageReadWriteExecuteWritePreEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context,
    BOOLEAN *IsTriggeringPostEventAllowed) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  BOOLEAN ShortCircuitingEvent = FALSE;
  EventTriggerResult = VmmCallbackTriggerEvents(
      HIDDEN_HOOK_WRITE, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_WRITE_AND_EXECUTE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  return ShortCircuitingEvent;
}

BOOLEAN DispatchEventHiddenHookPageReadWriteExecuteExecutePreEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context,
    BOOLEAN *IsTriggeringPostEventAllowed) {
  VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE EventTriggerResult;
  BOOLEAN PostEventTriggerReq = FALSE;
  BOOLEAN ShortCircuitingEvent = FALSE;
  EventTriggerResult = VmmCallbackTriggerEvents(
      HIDDEN_HOOK_EXECUTE, VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
      Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_EXECUTE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_WRITE_AND_EXECUTE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  EventTriggerResult =
      VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE,
                               VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION,
                               Context, &PostEventTriggerReq, VCpu->Regs);
  if (EventTriggerResult ==
      VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT) {
    ShortCircuitingEvent = TRUE;
  }
  if (PostEventTriggerReq) {
    *IsTriggeringPostEventAllowed = TRUE;
  }
  return ShortCircuitingEvent;
}

VOID DispatchEventHiddenHookPageReadWriteExecReadPostEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context) {
  VmmCallbackTriggerEvents(HIDDEN_HOOK_READ,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_EXECUTE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
}

VOID DispatchEventHiddenHookPageReadWriteExecWritePostEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context) {
  VmmCallbackTriggerEvents(HIDDEN_HOOK_WRITE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_WRITE_AND_EXECUTE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
}

VOID DispatchEventHiddenHookPageReadWriteExecExecutePostEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context) {
  VmmCallbackTriggerEvents(HIDDEN_HOOK_EXECUTE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_EXECUTE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_WRITE_AND_EXECUTE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
  VmmCallbackTriggerEvents(HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE,
                           VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION,
                           Context, NULL, VCpu->Regs);
}
