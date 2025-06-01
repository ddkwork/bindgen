
#include "pch.h"

VOID ProtectedHvChangeExceptionBitmapWithIntegrityCheck(
    VIRTUAL_MACHINE_STATE *VCpu, UINT32 CurrentMask,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  if (VmmCallbackQueryTerminateProtectedResource(
          VCpu->CoreId, PROTECTED_HV_RESOURCES_EXCEPTION_BITMAP, &CurrentMask,
          PassOver)) {
    return;
  }
  if (g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger) {
    CurrentMask |= 1 << EXCEPTION_VECTOR_PAGE_FAULT;
  }
  if (EptHookGetCountOfEpthooks(FALSE) != 0) {
    CurrentMask |= 1 << EXCEPTION_VECTOR_BREAKPOINT;
  }
  if (g_TransparentMode) {
    CurrentMask |= 1 << EXCEPTION_VECTOR_DEBUG_BREAKPOINT;
    CurrentMask |= 1 << EXCEPTION_VECTOR_BREAKPOINT;
  }
  HvWriteExceptionBitmap(CurrentMask);
}

VOID ProtectedHvSetExceptionBitmap(VIRTUAL_MACHINE_STATE *VCpu,
                                   UINT32 IdtIndex) {
  UINT32 ExceptionBitmap = 0;
  ExceptionBitmap = HvReadExceptionBitmap();
  if (IdtIndex == DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES) {
    ExceptionBitmap = 0xffffffff;
  } else {
    ExceptionBitmap |= 1 << IdtIndex;
  }
  ProtectedHvChangeExceptionBitmapWithIntegrityCheck(VCpu, ExceptionBitmap,
                                                     PASSING_OVER_NONE);
}

VOID ProtectedHvUnsetExceptionBitmap(VIRTUAL_MACHINE_STATE *VCpu,
                                     UINT32 IdtIndex) {
  UINT32 ExceptionBitmap = 0;
  ExceptionBitmap = HvReadExceptionBitmap();
  if (IdtIndex == DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES) {
    ExceptionBitmap = 0x0;
  } else {
    ExceptionBitmap &= ~(1 << IdtIndex);
  }
  ProtectedHvChangeExceptionBitmapWithIntegrityCheck(VCpu, ExceptionBitmap,
                                                     PASSING_OVER_NONE);
}

VOID ProtectedHvResetExceptionBitmapToClearEvents(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 ExceptionBitmap = 0;
  ProtectedHvChangeExceptionBitmapWithIntegrityCheck(
      VCpu, ExceptionBitmap, PASSING_OVER_EXCEPTION_EVENTS);
}

VOID ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands(
    VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 ExceptionBitmap = 0;
  ExceptionBitmap = HvReadExceptionBitmap();
  ExceptionBitmap &= ~(1 << EXCEPTION_VECTOR_UNDEFINED_OPCODE);
  ProtectedHvChangeExceptionBitmapWithIntegrityCheck(
      VCpu, ExceptionBitmap,
      PASSING_OVER_UD_EXCEPTIONS_FOR_SYSCALL_SYSRET_HOOK);
}

VOID ProtectedHvApplySetExternalInterruptExiting(
    VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  UINT32 PinBasedControls = 0;
  UINT32 VmExitControls = 0;
  if (Set == FALSE) {
    if (VmmCallbackQueryTerminateProtectedResource(
            VCpu->CoreId, PROTECTED_HV_RESOURCES_EXTERNAL_INTERRUPT_EXITING,
            NULL, PassOver)) {
      return;
    }
  }
  VmxVmread32P(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, &PinBasedControls);
  VmxVmread32P(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, &VmExitControls);
  if (Set) {
    PinBasedControls |= PIN_BASED_VM_EXECUTION_CONTROLS_EXTERNAL_INTERRUPT;
    VmExitControls |= VM_EXIT_ACK_INTR_ON_EXIT;
  } else {
    PinBasedControls &= ~PIN_BASED_VM_EXECUTION_CONTROLS_EXTERNAL_INTERRUPT;
    VmExitControls &= ~VM_EXIT_ACK_INTR_ON_EXIT;
  }
  VmxVmwrite64(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, PinBasedControls);
  VmxVmwrite64(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, VmExitControls);
}

VOID ProtectedHvSetExternalInterruptExiting(VIRTUAL_MACHINE_STATE *VCpu,
                                            BOOLEAN Set) {
  ProtectedHvApplySetExternalInterruptExiting(VCpu, Set, PASSING_OVER_NONE);
}

VOID ProtectedHvExternalInterruptExitingForDisablingInterruptCommands(
    VIRTUAL_MACHINE_STATE *VCpu) {
  ProtectedHvApplySetExternalInterruptExiting(VCpu, FALSE,
                                              PASSING_OVER_INTERRUPT_EVENTS);
}

VOID ProtectedHvSetTscVmexit(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set,
                             PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  UINT32 CpuBasedVmExecControls = 0;
  if (Set == FALSE) {
    if (VmmCallbackQueryTerminateProtectedResource(
            VCpu->CoreId, PROTECTED_HV_RESOURCES_RDTSC_RDTSCP_EXITING, NULL,
            PassOver)) {
      return;
    }
  }
  VmxVmread32P(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_RDTSC_EXITING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_RDTSC_EXITING;
  }
  VmxVmwrite64(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               CpuBasedVmExecControls);
}

VOID ProtectedHvSetMovDebugRegsVmexit(
    VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  UINT32 CpuBasedVmExecControls = 0;
  if (Set == FALSE) {
    if (VmmCallbackQueryTerminateProtectedResource(
            VCpu->CoreId, PROTECTED_HV_RESOURCES_MOV_TO_DEBUG_REGISTER_EXITING,
            NULL, PassOver)) {
      return;
    }
  }
  VmxVmread32P(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_MOV_DR_EXITING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_MOV_DR_EXITING;
  }
  VmxVmwrite64(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               CpuBasedVmExecControls);
}

VOID ProtectedHvSetMovToCrVmexit(BOOLEAN Set, UINT64 ControlRegister,
                                 UINT64 MaskRegister) {
  if (ControlRegister == VMX_EXIT_QUALIFICATION_REGISTER_CR0) {
    if (Set) {
      VmxVmwrite64(VMCS_CTRL_CR0_GUEST_HOST_MASK, MaskRegister);
      VmxVmwrite64(VMCS_CTRL_CR0_READ_SHADOW, __readcr0());
    } else {
      VmxVmwrite64(VMCS_CTRL_CR0_GUEST_HOST_MASK, 0);
      VmxVmwrite64(VMCS_CTRL_CR0_READ_SHADOW, 0);
    }
  } else if (ControlRegister == VMX_EXIT_QUALIFICATION_REGISTER_CR4) {
    if (Set) {
      VmxVmwrite64(VMCS_CTRL_CR4_GUEST_HOST_MASK, MaskRegister);
      VmxVmwrite64(VMCS_CTRL_CR4_READ_SHADOW, __readcr0());
    } else {
      VmxVmwrite64(VMCS_CTRL_CR4_GUEST_HOST_MASK, 0);
      VmxVmwrite64(VMCS_CTRL_CR4_READ_SHADOW, 0);
    }
  }
}

VOID ProtectedHvSetMovControlRegsVmexit(
    VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver, UINT64 ControlRegister,
    UINT64 MaskRegister) {
  if (Set == FALSE) {
    if (VmmCallbackQueryTerminateProtectedResource(
            VCpu->CoreId, PROTECTED_HV_RESOURCES_MOV_CONTROL_REGISTER_EXITING,
            NULL, PassOver)) {
      return;
    }
  }
  ProtectedHvSetMovToCrVmexit(Set, ControlRegister, MaskRegister);
}

VOID ProtectedHvSetMovToCr3Vmexit(
    VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  UINT32 CpuBasedVmExecControls = 0;
  if (Set == FALSE) {
    if (VmmCallbackQueryTerminateProtectedResource(
            VCpu->CoreId, PROTECTED_HV_RESOURCES_MOV_TO_CR3_EXITING, NULL,
            PassOver)) {
      return;
    }
    if (g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger) {
      return;
    }
    if (g_ExecTrapInitialized && !g_ExecTrapUnInitializationStarted) {
      return;
    }
  }
  VmxVmread32P(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_CR3_LOAD_EXITING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_CR3_LOAD_EXITING;
  }
  VmxVmwrite64(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               CpuBasedVmExecControls);
}

VOID ProtectedHvSetRdtscExiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set) {
  ProtectedHvSetTscVmexit(VCpu, Set, PASSING_OVER_NONE);
}

VOID ProtectedHvDisableRdtscExitingForDisablingTscCommands(
    VIRTUAL_MACHINE_STATE *VCpu) {
  ProtectedHvSetTscVmexit(VCpu, FALSE, PASSING_OVER_TSC_EVENTS);
}

VOID ProtectedHvSetMovDebugRegsExiting(VIRTUAL_MACHINE_STATE *VCpu,
                                       BOOLEAN Set) {
  ProtectedHvSetMovDebugRegsVmexit(VCpu, Set, PASSING_OVER_NONE);
}

VOID ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands(
    VIRTUAL_MACHINE_STATE *VCpu) {
  ProtectedHvSetMovDebugRegsVmexit(VCpu, FALSE,
                                   PASSING_OVER_MOV_TO_HW_DEBUG_REGS_EVENTS);
}

VOID ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands(
    VIRTUAL_MACHINE_STATE *VCpu, UINT64 ControlRegister, UINT64 MaskRegister) {
  ProtectedHvSetMovControlRegsVmexit(VCpu, FALSE,
                                     PASSING_OVER_MOV_TO_CONTROL_REGS_EVENTS,
                                     ControlRegister, MaskRegister);
}

VOID ProtectedHvSetMov2Cr3Exiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set) {
  ProtectedHvSetMovToCr3Vmexit(VCpu, Set, PASSING_OVER_NONE);
}

VOID ProtectedHvSetMov2CrExiting(BOOLEAN Set, UINT64 ControlRegister,
                                 UINT64 MaskRegister) {
  ProtectedHvSetMovToCrVmexit(Set, ControlRegister, MaskRegister);
}
