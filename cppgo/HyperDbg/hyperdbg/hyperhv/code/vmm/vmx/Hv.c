
#include "pch.h"

UINT32 HvAdjustControls(UINT32 Ctl, UINT32 Msr) {
  MSR MsrValue = {0};
  MsrValue.Flags = __readmsr(Msr);
  Ctl &= MsrValue.Fields.High;
  Ctl |= MsrValue.Fields.Low;
  return Ctl;
}

BOOLEAN HvSetGuestSelector(PVOID GdtBase, UINT32 SegmentRegister,
                           UINT16 Selector) {
  VMX_SEGMENT_SELECTOR SegmentSelector = {0};
  SegmentGetDescriptor(GdtBase, Selector, &SegmentSelector);
  if (Selector == 0x0) {
    SegmentSelector.Attributes.Unusable = TRUE;
  }
  VmxVmwrite64(VMCS_GUEST_ES_SELECTOR + SegmentRegister * 2, Selector);
  VmxVmwrite64(VMCS_GUEST_ES_LIMIT + SegmentRegister * 2,
               SegmentSelector.Limit);
  VmxVmwrite64(VMCS_GUEST_ES_ACCESS_RIGHTS + SegmentRegister * 2,
               SegmentSelector.Attributes.AsUInt);
  VmxVmwrite64(VMCS_GUEST_ES_BASE + SegmentRegister * 2, SegmentSelector.Base);
  return TRUE;
}

VOID HvHandleCpuid(VIRTUAL_MACHINE_STATE *VCpu) {
  INT32 CpuInfo[4];
  PGUEST_REGS Regs = VCpu->Regs;
  __cpuidex(CpuInfo, (INT32)Regs->rax, (INT32)Regs->rcx);
  if (!g_TransparentMode) {
    if (Regs->rax == CPUID_PROCESSOR_AND_PROCESSOR_FEATURE_IDENTIFIERS) {
      CpuInfo[2] |= HYPERV_HYPERVISOR_PRESENT_BIT;
    } else if (Regs->rax == CPUID_HV_VENDOR_AND_MAX_FUNCTIONS) {
      CpuInfo[0] = HYPERV_CPUID_INTERFACE;
      CpuInfo[1] = 'epyH';
      CpuInfo[2] = 'gbDr';
      CpuInfo[3] = 0;
    } else if (Regs->rax == HYPERV_CPUID_INTERFACE) {
      CpuInfo[0] = '0#vH';
      CpuInfo[1] = CpuInfo[2] = CpuInfo[3] = 0;
    } else {
      TransparentCPUID(CpuInfo, Regs);
    }
  }
  Regs->rax = CpuInfo[0];
  Regs->rbx = CpuInfo[1];
  Regs->rcx = CpuInfo[2];
  Regs->rdx = CpuInfo[3];
}

VOID HvHandleControlRegisterAccess(
    VIRTUAL_MACHINE_STATE *VCpu,
    VMX_EXIT_QUALIFICATION_MOV_CR *CrExitQualification) {
  UINT64 *RegPtr;
  UINT64 NewCr3;
  CR3_TYPE NewCr3Reg;
  RegPtr =
      (UINT64 *)&VCpu->Regs->rax + CrExitQualification->GeneralPurposeRegister;

  switch (CrExitQualification->AccessType) {
    case VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR: {
      switch (CrExitQualification->ControlRegister) {
        case VMX_EXIT_QUALIFICATION_REGISTER_CR0:
          VmxVmwrite64(VMCS_GUEST_CR0, *RegPtr);
          VmxVmwrite64(VMCS_CTRL_CR0_READ_SHADOW, *RegPtr);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_CR3:
          NewCr3 = (*RegPtr & ~(1ULL << 63));
          NewCr3Reg.Flags = NewCr3;
          VmxVmwrite64(VMCS_GUEST_CR3, NewCr3Reg.Flags);
          VpidInvvpidSingleContext(VPID_TAG);
          InterceptionCallbackTriggerCr3ProcessChange(VCpu->CoreId);
          if (g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger) {
            InterceptionCallbackCr3VmexitsForThreadInterception(VCpu->CoreId,
                                                                NewCr3Reg);
          }
          if (g_ExecTrapInitialized) {
            ExecTrapHandleCr3Vmexit(VCpu);
          }
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_CR4:
          VmxVmwrite64(VMCS_GUEST_CR4, *RegPtr);
          VmxVmwrite64(VMCS_CTRL_CR4_READ_SHADOW, *RegPtr);
          break;
        default:
          LogWarning(
              "Unsupported register 0x%x in handling control registers access",
              CrExitQualification->ControlRegister);
          break;
      }
    } break;
    case VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR: {
      switch (CrExitQualification->ControlRegister) {
        case VMX_EXIT_QUALIFICATION_REGISTER_CR0:
          __vmx_vmread(VMCS_GUEST_CR0, RegPtr);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_CR3:
          __vmx_vmread(VMCS_GUEST_CR3, RegPtr);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_CR4:
          __vmx_vmread(VMCS_GUEST_CR4, RegPtr);
          break;
        default:
          LogWarning(
              "Unsupported register 0x%x in handling control registers access",
              CrExitQualification->ControlRegister);
          break;
      }
    } break;
    default:
      LogWarning(
          "Unsupported operation 0x%x in handling control registers access",
          CrExitQualification->AccessType);
      break;
  }
}

VOID HvFillGuestSelectorData(PVOID GdtBase, UINT32 SegmentRegister,
                             UINT16 Selector) {
  VMX_SEGMENT_SELECTOR SegmentSelector = {0};
  SegmentGetDescriptor(GdtBase, Selector, &SegmentSelector);
  if (Selector == 0x0) {
    SegmentSelector.Attributes.Unusable = TRUE;
  }
  SegmentSelector.Attributes.Reserved1 = 0;
  SegmentSelector.Attributes.Reserved2 = 0;
  VmxVmwrite64(VMCS_GUEST_ES_SELECTOR + SegmentRegister * 2, Selector);
  VmxVmwrite64(VMCS_GUEST_ES_LIMIT + SegmentRegister * 2,
               SegmentSelector.Limit);
  VmxVmwrite64(VMCS_GUEST_ES_ACCESS_RIGHTS + SegmentRegister * 2,
               SegmentSelector.Attributes.AsUInt);
  VmxVmwrite64(VMCS_GUEST_ES_BASE + SegmentRegister * 2, SegmentSelector.Base);
}

VOID HvResumeToNextInstruction() {
  UINT64 ResumeRIP = NULL64_ZERO;
  UINT64 CurrentRIP = NULL64_ZERO;
  size_t ExitInstructionLength = 0;
  __vmx_vmread(VMCS_GUEST_RIP, &CurrentRIP);
  __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstructionLength);
  ResumeRIP = CurrentRIP + ExitInstructionLength;
  VmxVmwrite64(VMCS_GUEST_RIP, ResumeRIP);
}

inline VOID HvSuppressRipIncrement(VIRTUAL_MACHINE_STATE *VCpu) {
  VCpu->IncrementRip = FALSE;
}

inline VOID HvPerformRipIncrement(VIRTUAL_MACHINE_STATE *VCpu) {
  VCpu->IncrementRip = TRUE;
}

VOID HvSetMonitorTrapFlag(BOOLEAN Set) {
  UINT32 CpuBasedVmExecControls = 0;
  VmxVmread32P(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_MONITOR_TRAP_FLAG;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_MONITOR_TRAP_FLAG;
  }
  VmxVmwrite64(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               CpuBasedVmExecControls);
}

VOID HvSetRflagTrapFlag(BOOLEAN Set) {
  RFLAGS Rflags = {0};
  Rflags.AsUInt = HvGetRflags();
  Rflags.TrapFlag = Set;
  HvSetRflags(Rflags.AsUInt);
}

VOID HvSetLoadDebugControls(BOOLEAN Set) {
  UINT32 VmentryControls = 0;
  VmxVmread32P(VMCS_CTRL_VMENTRY_CONTROLS, &VmentryControls);
  if (Set) {
    VmentryControls |= VM_ENTRY_LOAD_DEBUG_CONTROLS;
  } else {
    VmentryControls &= ~VM_ENTRY_LOAD_DEBUG_CONTROLS;
  }
  VmxVmwrite64(VMCS_CTRL_VMENTRY_CONTROLS, VmentryControls);
}

VOID HvSetSaveDebugControls(BOOLEAN Set) {
  UINT32 VmexitControls = 0;
  VmxVmread32P(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, &VmexitControls);
  if (Set) {
    VmexitControls |= VM_EXIT_SAVE_DEBUG_CONTROLS;
  } else {
    VmexitControls &= ~VM_EXIT_SAVE_DEBUG_CONTROLS;
  }
  VmxVmwrite64(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, VmexitControls);
}

VOID HvRestoreRegisters() {
  UINT64 FsBase;
  UINT64 GsBase;
  UINT64 GdtrBase;
  UINT64 GdtrLimit;
  UINT64 IdtrBase;
  UINT64 IdtrLimit;
  UINT16 DsSelector;
  UINT16 EsSelector;
  UINT16 SsSelector;
  UINT16 FsSelector;
  __vmx_vmread(VMCS_GUEST_FS_BASE, &FsBase);
  __writemsr(IA32_FS_BASE, FsBase);
  __vmx_vmread(VMCS_GUEST_GS_BASE, &GsBase);
  __writemsr(IA32_GS_BASE, GsBase);
  __vmx_vmread(VMCS_GUEST_GDTR_BASE, &GdtrBase);
  __vmx_vmread(VMCS_GUEST_GDTR_LIMIT, &GdtrLimit);
  AsmReloadGdtr((void *)GdtrBase, (unsigned long)GdtrLimit);
  VmxVmread16P(VMCS_GUEST_DS_SELECTOR, &DsSelector);
  AsmSetDs(DsSelector);
  VmxVmread16P(VMCS_GUEST_ES_SELECTOR, &EsSelector);
  AsmSetEs(EsSelector);
  VmxVmread16P(VMCS_GUEST_SS_SELECTOR, &SsSelector);
  AsmSetSs(SsSelector);
  VmxVmread16P(VMCS_GUEST_FS_SELECTOR, &FsSelector);
  AsmSetFs(FsSelector);
  __vmx_vmread(VMCS_GUEST_IDTR_BASE, &IdtrBase);
  __vmx_vmread(VMCS_GUEST_IDTR_LIMIT, &IdtrLimit);
  AsmReloadIdtr((void *)IdtrBase, (unsigned long)IdtrLimit);
}

VOID HvSetPmcVmexit(BOOLEAN Set) {
  UINT32 CpuBasedVmExecControls = 0;
  VmxVmread32P(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_RDPMC_EXITING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_RDPMC_EXITING;
  }
  VmxVmwrite64(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               CpuBasedVmExecControls);
}

VOID HvSetMovControlRegsExiting(BOOLEAN Set, UINT64 ControlRegister,
                                UINT64 MaskRegister) {
  ProtectedHvSetMov2CrExiting(Set, ControlRegister, MaskRegister);
}

VOID HvSetMovToCr3Vmexit(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set) {
  ProtectedHvSetMov2Cr3Exiting(VCpu, Set);
}

VOID HvWriteExceptionBitmap(UINT32 BitmapMask) {
  VmxVmwrite64(VMCS_CTRL_EXCEPTION_BITMAP, BitmapMask);
}

UINT32 HvReadExceptionBitmap() {
  UINT32 ExceptionBitmap = 0;
  VmxVmread32P(VMCS_CTRL_EXCEPTION_BITMAP, &ExceptionBitmap);
  return ExceptionBitmap;
}

VOID HvSetInterruptWindowExiting(BOOLEAN Set) {
  UINT32 CpuBasedVmExecControls = 0;
  VmxVmread32P(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_VIRTUAL_INTR_PENDING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_VIRTUAL_INTR_PENDING;
  }
  VmxVmwrite64(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               CpuBasedVmExecControls);
}

VOID HvSetPmlEnableFlag(BOOLEAN Set) {
  UINT32 AdjSecCtrl;
  UINT32 SecondaryProcBasedVmExecControls = 0;
  VmxVmread32P(VMCS_CTRL_SECONDARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &SecondaryProcBasedVmExecControls);
  if (Set) {
    SecondaryProcBasedVmExecControls |=
        IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_FLAG;
  } else {
    SecondaryProcBasedVmExecControls &=
        ~IA32_VMX_PROCBASED_CTLS2_ENABLE_PML_FLAG;
  }
  AdjSecCtrl = HvAdjustControls(SecondaryProcBasedVmExecControls,
                                IA32_VMX_PROCBASED_CTLS2);
  VmxVmwrite64(VMCS_CTRL_SECONDARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               AdjSecCtrl);
}

VOID HvSetModeBasedExecutionEnableFlag(BOOLEAN Set) {
  UINT32 AdjSecCtrl;
  UINT32 SecondaryProcBasedVmExecControls = 0;
  VmxVmread32P(VMCS_CTRL_SECONDARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &SecondaryProcBasedVmExecControls);
  if (Set) {
    SecondaryProcBasedVmExecControls |=
        IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_FLAG;
  } else {
    SecondaryProcBasedVmExecControls &=
        ~IA32_VMX_PROCBASED_CTLS2_MODE_BASED_EXECUTE_CONTROL_FOR_EPT_FLAG;
  }
  AdjSecCtrl = HvAdjustControls(SecondaryProcBasedVmExecControls,
                                IA32_VMX_PROCBASED_CTLS2);
  VmxVmwrite64(VMCS_CTRL_SECONDARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               AdjSecCtrl);
}

VOID HvSetNmiWindowExiting(BOOLEAN Set) {
  UINT32 CpuBasedVmExecControls = 0;
  VmxVmread32P(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               &CpuBasedVmExecControls);
  if (Set) {
    CpuBasedVmExecControls |= CPU_BASED_VIRTUAL_NMI_PENDING;
  } else {
    CpuBasedVmExecControls &= ~CPU_BASED_VIRTUAL_NMI_PENDING;
  }
  VmxVmwrite64(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               CpuBasedVmExecControls);
}

VOID HvHandleMovDebugRegister(VIRTUAL_MACHINE_STATE *VCpu) {
  VMX_EXIT_QUALIFICATION_MOV_DR ExitQualification;
  CR4 Cr4;
  DR7 Dr7;
  VMX_SEGMENT_SELECTOR Cs;
  UINT64 *GpRegs = (UINT64 *)VCpu->Regs;
  VmxVmread64P(VMCS_EXIT_QUALIFICATION, &ExitQualification.AsUInt);
  UINT64 GpRegister = GpRegs[ExitQualification.GeneralPurposeRegister];
  Cs = GetGuestCs();
  if (Cs.Attributes.DescriptorPrivilegeLevel != 0) {
    EventInjectGeneralProtection();
    HvSuppressRipIncrement(VCpu);
    return;
  }
  VmxVmread64P(VMCS_GUEST_CR4, &Cr4.AsUInt);
  if (ExitQualification.DebugRegister == 4 ||
      ExitQualification.DebugRegister == 5) {
    if (Cr4.DebuggingExtensions) {
      EventInjectUndefinedOpcode(VCpu);
      return;
    } else {
      ExitQualification.DebugRegister += 2;
    }
  }
  VmxVmread64P(VMCS_GUEST_DR7, &Dr7.AsUInt);
  if (Dr7.GeneralDetect) {
    DR6 Dr6 = {.AsUInt = __readdr(6),
               .BreakpointCondition = 0,
               .DebugRegisterAccessDetected = TRUE};
    __writedr(6, Dr6.AsUInt);
    Dr7.GeneralDetect = FALSE;
    VmxVmwrite64(VMCS_GUEST_DR7, Dr7.AsUInt);
    EventInjectDebugBreakpoint();
    HvSuppressRipIncrement(VCpu);
    return;
  }
  if (ExitQualification.DirectionOfAccess ==
          VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR &&
      (ExitQualification.DebugRegister == VMX_EXIT_QUALIFICATION_REGISTER_DR6 ||
       ExitQualification.DebugRegister ==
           VMX_EXIT_QUALIFICATION_REGISTER_DR7) &&
      (GpRegister >> 32) != 0) {
    EventInjectGeneralProtection();
    HvSuppressRipIncrement(VCpu);
    return;
  }
  switch (ExitQualification.DirectionOfAccess) {
    case VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR:
      switch (ExitQualification.DebugRegister) {
        case VMX_EXIT_QUALIFICATION_REGISTER_DR0:
          __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR0, GpRegister);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR1:
          __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR1, GpRegister);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR2:
          __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR2, GpRegister);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR3:
          __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR3, GpRegister);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR6:
          __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR6, GpRegister);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR7:
          __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR7, GpRegister);
          break;
        default:
          break;
      }
      break;
    case VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR:
      switch (ExitQualification.DebugRegister) {
        case VMX_EXIT_QUALIFICATION_REGISTER_DR0:
          GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR0);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR1:
          GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR1);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR2:
          GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR2);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR3:
          GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR3);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR6:
          GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR6);
          break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR7:
          GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR7);
          break;
        default:
          break;
      }
    default:
      break;
  }
}

VOID HvSetNmiExiting(BOOLEAN Set) {
  UINT32 PinBasedControls = 0;
  UINT32 VmExitControls = 0;
  VmxVmread32P(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, &PinBasedControls);
  VmxVmread32P(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, &VmExitControls);
  if (Set) {
    PinBasedControls |= PIN_BASED_VM_EXECUTION_CONTROLS_NMI_EXITING;
    VmExitControls |= VM_EXIT_ACK_INTR_ON_EXIT;
  } else {
    PinBasedControls &= ~PIN_BASED_VM_EXECUTION_CONTROLS_NMI_EXITING;
    VmExitControls &= ~VM_EXIT_ACK_INTR_ON_EXIT;
  }
  VmxVmwrite64(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, PinBasedControls);
  VmxVmwrite64(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, VmExitControls);
}

VOID HvSetVmxPreemptionTimerExiting(BOOLEAN Set) {
  UINT32 PinBasedControls = 0;
  VmxVmread32P(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, &PinBasedControls);
  if (Set) {
    PinBasedControls |= PIN_BASED_VM_EXECUTION_CONTROLS_ACTIVE_VMX_TIMER;
  } else {
    PinBasedControls &= ~PIN_BASED_VM_EXECUTION_CONTROLS_ACTIVE_VMX_TIMER;
  }
  VmxVmwrite64(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, PinBasedControls);
}

VOID HvSetExceptionBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 IdtIndex) {
  ProtectedHvSetExceptionBitmap(VCpu, IdtIndex);
}

VOID HvUnsetExceptionBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 IdtIndex) {
  ProtectedHvUnsetExceptionBitmap(VCpu, IdtIndex);
}

VOID HvSetExternalInterruptExiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set) {
  ProtectedHvSetExternalInterruptExiting(VCpu, Set);
}

VOID HvEnableAndCheckForPreviousExternalInterrupts(
    VIRTUAL_MACHINE_STATE *VCpu) {
  if (VCpu->EnableExternalInterruptsOnContinueMtf) {
    HvSetExternalInterruptExiting(VCpu, FALSE);
    if (VCpu->PendingExternalInterrupts[0] != NULL_ZERO) {
      HvSetInterruptWindowExiting(TRUE);
    }
    VCpu->EnableExternalInterruptsOnContinueMtf = FALSE;
  }
}

VOID HvSetRdtscExiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set) {
  ProtectedHvSetRdtscExiting(VCpu, Set);
}

VOID HvSetMovDebugRegsExiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set) {
  ProtectedHvSetMovDebugRegsExiting(VCpu, Set);
}

UINT16 HvGetCsSelector() {
  UINT64 CsSel = NULL64_ZERO;
  __vmx_vmread(VMCS_GUEST_CS_SELECTOR, &CsSel);
  return CsSel & 0xffff;
}

UINT64 HvGetRflags() {
  UINT64 Rflags = NULL64_ZERO;
  __vmx_vmread(VMCS_GUEST_RFLAGS, &Rflags);
  return Rflags;
}

VOID HvSetRflags(UINT64 Rflags) { VmxVmwrite64(VMCS_GUEST_RFLAGS, Rflags); }

UINT64 HvGetRip() {
  UINT64 Rip = NULL64_ZERO;
  __vmx_vmread(VMCS_GUEST_RIP, &Rip);
  return Rip;
}

VOID HvSetRip(UINT64 Rip) { VmxVmwrite64(VMCS_GUEST_RIP, Rip); }

UINT64 HvGetInterruptibilityState() {
  UINT64 InterruptibilityState = NULL64_ZERO;
  __vmx_vmread(VMCS_GUEST_INTERRUPTIBILITY_STATE, &InterruptibilityState);
  return InterruptibilityState;
}

UINT64 HvClearSteppingBits(UINT64 Interruptibility) {
  UINT64 InterruptibilityState = Interruptibility;
  InterruptibilityState &= ~(GUEST_INTR_STATE_STI | GUEST_INTR_STATE_MOV_SS);
  return InterruptibilityState;
}

VOID HvSetInterruptibilityState(UINT64 InterruptibilityState) {
  VmxVmwrite64(VMCS_GUEST_INTERRUPTIBILITY_STATE, InterruptibilityState);
}

VOID HvInjectPendingExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu) {
  if (VCpu->PendingExternalInterrupts[0] != NULL_ZERO) {
    HvSetInterruptWindowExiting(TRUE);
  }
}

VOID HvCheckAndEnableExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu) {
  if (VCpu->EnableExternalInterruptsOnContinue) {
    HvSetExternalInterruptExiting(VCpu, FALSE);
    HvInjectPendingExternalInterrupts(VCpu);
    VCpu->EnableExternalInterruptsOnContinue = FALSE;
  }
}

VOID HvDisableExternalInterruptsAndInterruptWindow(
    VIRTUAL_MACHINE_STATE *VCpu) {
  HvSetExternalInterruptExiting(VCpu, TRUE);
  HvSetInterruptWindowExiting(FALSE);
  VCpu->EnableExternalInterruptsOnContinue = TRUE;
}

BOOLEAN HvInitVmm(VMM_CALLBACKS *VmmCallbacks) {
  ULONG ProcessorsCount;
  BOOLEAN Result = FALSE;
  RtlCopyMemory(&g_Callbacks, VmmCallbacks, sizeof(VMM_CALLBACKS));
  CompatibilityCheckPerformChecks();
  Result = GlobalGuestStateAllocateZeroedMemory();
  if (!Result) {
    return FALSE;
  }
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  for (UINT32 i = 0; i < ProcessorsCount; i++) {
    g_GuestState[i].CoreId = i;
  }
  MemoryMapperInitialize();
  g_TransparentMode = FALSE;
  g_WaitingForInterruptWindowToInjectPageFault = FALSE;
  return VmxInitialize();
}

VOID HvEnableMtfAndChangeExternalInterruptState(VIRTUAL_MACHINE_STATE *VCpu) {
  HvSetMonitorTrapFlag(TRUE);
  HvSetExternalInterruptExiting(VCpu, TRUE);
  HvSetInterruptWindowExiting(FALSE);
  VCpu->EnableExternalInterruptsOnContinueMtf = TRUE;
}

VOID HvPreventExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu) {
  HvSetExternalInterruptExiting(VCpu, TRUE);
  HvSetInterruptWindowExiting(FALSE);
  VCpu->EnableExternalInterruptsOnContinueMtf = TRUE;
}

UINT64 HvGetPendingDebugExceptions() {
  UINT64 Value;
  VmxVmread64P(VMCS_GUEST_PENDING_DEBUG_EXCEPTIONS, &Value);
  return Value;
}

VOID HvSetPendingDebugExceptions(UINT64 Value) {
  VmxVmwrite64(VMCS_GUEST_PENDING_DEBUG_EXCEPTIONS, Value);
}

UINT64 HvGetDebugctl() {
  UINT32 LowPart;
  UINT32 HighPart;
  VmxVmread32P(VMCS_GUEST_DEBUGCTL, &LowPart);
  VmxVmread32P(VMCS_GUEST_DEBUGCTL_HIGH, &HighPart);
  return (UINT64)HighPart << 32 | LowPart;
}

VOID HvSetDebugctl(UINT64 Value) {
  VmxVmwrite64(VMCS_GUEST_DEBUGCTL, Value & 0xFFFFFFFF);
  VmxVmwrite64(VMCS_GUEST_DEBUGCTL_HIGH, Value >> 32);
}

VOID HvHandleTrapFlag() {
  IA32_DEBUGCTL_REGISTER Debugctl = {.AsUInt = HvGetDebugctl()};
  RFLAGS GuestRFlags = {.AsUInt = GetGuestRFlags()};
  VMX_PENDING_DEBUG_EXCEPTIONS PendingDebugExceptions;
  VMX_INTERRUPTIBILITY_STATE InterruptibilityState;
  if (GuestRFlags.TrapFlag && !Debugctl.Btf) {
    PendingDebugExceptions.AsUInt = HvGetPendingDebugExceptions();
    PendingDebugExceptions.Bs = 1;
    HvSetPendingDebugExceptions(PendingDebugExceptions.AsUInt);
    InterruptibilityState.AsUInt = (UINT32)HvGetInterruptibilityState();
    if (InterruptibilityState.BlockingByMovSs) {
      InterruptibilityState.BlockingByMovSs = 0;
      HvSetInterruptibilityState(InterruptibilityState.AsUInt);
    }
  }
}
