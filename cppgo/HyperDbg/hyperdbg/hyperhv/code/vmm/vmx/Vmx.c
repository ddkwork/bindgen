
#include "pch.h"

inline UCHAR VmxVmread64(size_t Field, UINT64 FieldValue) {
  return __vmx_vmread((size_t)Field, (size_t *)FieldValue);
}

inline UCHAR VmxVmread32(size_t Field, UINT32 FieldValue) {
  UINT64 TargetField = 0ull;
  TargetField = FieldValue;
  return __vmx_vmread((size_t)Field, (size_t *)TargetField);
}

inline UCHAR VmxVmread16(size_t Field, UINT16 FieldValue) {
  UINT64 TargetField = 0ull;
  TargetField = FieldValue;
  return __vmx_vmread((size_t)Field, (size_t *)TargetField);
}

inline UCHAR VmxVmread64P(size_t Field, UINT64 *FieldValue) {
  return __vmx_vmread((size_t)Field, (size_t *)FieldValue);
}

inline UCHAR VmxVmread32P(size_t Field, UINT32 *FieldValue) {
  UINT64 TargetField = 0ull;
  TargetField = (UINT64)FieldValue;
  return __vmx_vmread((size_t)Field, (size_t *)TargetField);
}

inline UCHAR VmxVmread16P(size_t Field, UINT16 *FieldValue) {
  UINT64 TargetField = 0ull;
  TargetField = (UINT64)FieldValue;
  return __vmx_vmread((size_t)Field, (size_t *)TargetField);
}

inline UCHAR VmxVmwrite64(size_t Field, UINT64 FieldValue) {
  return __vmx_vmwrite((size_t)Field, (size_t)FieldValue);
}

inline UCHAR VmxVmwrite32(size_t Field, UINT32 FieldValue) {
  UINT64 TargetValue = NULL64_ZERO;
  TargetValue = (UINT64)FieldValue;
  return __vmx_vmwrite((size_t)Field, (size_t)TargetValue);
}

inline UCHAR VmxVmwrite16(size_t Field, UINT16 FieldValue) {
  UINT64 TargetValue = NULL64_ZERO;
  TargetValue = (UINT64)FieldValue;
  return __vmx_vmwrite((size_t)Field, (size_t)TargetValue);
}

BOOLEAN VmxCheckVmxSupport() {
  CPUID Data = {0};
  IA32_FEATURE_CONTROL_REGISTER FeatureControlMsr = {0};
  __cpuid((int *)&Data, 1);
  if (!_bittest((const LONG *)&Data.ecx, 5)) {
    return FALSE;
  }
  FeatureControlMsr.AsUInt = __readmsr(IA32_FEATURE_CONTROL);
  if (FeatureControlMsr.EnableVmxOutsideSmx == FALSE) {
    LogError("Err, you should enable vt-x from BIOS");
    return FALSE;
  }
  return TRUE;
}

BOOLEAN VmxGetCurrentExecutionMode() {
  if (g_GuestState) {
    ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
    VIRTUAL_MACHINE_STATE *CurrentVmState = &g_GuestState[CurrentCore];
    return CurrentVmState->IsOnVmxRootMode ? VmxExecutionModeRoot
                                           : VmxExecutionModeNonRoot;
  } else {
    return VmxExecutionModeNonRoot;
  }
}

BOOLEAN VmxGetCurrentLaunchState() {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  VIRTUAL_MACHINE_STATE *CurrentVmState = &g_GuestState[CurrentCore];
  return CurrentVmState->HasLaunched;
}

BOOLEAN VmxInitialize() {
  ULONG ProcessorsCount;
  if (!VmxPerformVirtualizationOnAllCores()) {
    return FALSE;
  }
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  for (size_t ProcessorID = 0; ProcessorID < ProcessorsCount; ProcessorID++) {
    VIRTUAL_MACHINE_STATE *GuestState = &g_GuestState[ProcessorID];
    if (!VmxAllocateVmmStack(GuestState)) {
      return FALSE;
    }
    if (!VmxAllocateMsrBitmap(GuestState)) {
      return FALSE;
    }
    if (!VmxAllocateIoBitmaps(GuestState)) {
      return FALSE;
    }
#if USE_DEFAULT_OS_IDT_AS_HOST_IDT == FALSE
    if (!VmxAllocateHostIdt(GuestState)) {
      return FALSE;
    }
#endif
#if USE_DEFAULT_OS_GDT_AS_HOST_GDT == FALSE
    if (!VmxAllocateHostGdt(GuestState)) {
      return FALSE;
    }
    if (!VmxAllocateHostTss(GuestState)) {
      return FALSE;
    }
#endif
#if USE_INTERRUPT_STACK_TABLE == TRUE
    if (!VmxAllocateHostInterruptStack(GuestState)) {
      return FALSE;
    }
#endif
  }
  g_MsrBitmapInvalidMsrs = VmxAllocateInvalidMsrBimap();
  if (g_MsrBitmapInvalidMsrs == NULL) {
    return FALSE;
  }
  KeGenericCallDpc(DpcRoutineInitializeGuest, 0x0);
  if (AsmVmxVmcall(VMCALL_TEST, 0x22, 0x333, 0x4444) == STATUS_SUCCESS) {
    return TRUE;
  } else {
    return FALSE;
  }
}

BOOLEAN VmxPerformVirtualizationOnAllCores() {
  PAGED_CODE();
  if (!VmxCheckVmxSupport()) {
    LogError("Err, VMX is not supported in this machine");
    return FALSE;
  }
  g_EptState = PlatformMemAllocateZeroedNonPagedPool(sizeof(EPT_STATE));
  if (!g_EptState) {
    LogError("Err, insufficient memory");
    return FALSE;
  }
  InitializeListHead(&g_EptState->HookedPagesList);
  if (!EptCheckFeatures()) {
    LogError("Err, your processor doesn't support all EPT features");
    return FALSE;
  } else {
    LogDebugInfo("Your processor supports all EPT features");
    if (!EptBuildMtrrMap()) {
      LogError("Err, could not build MTRR memory map");
      return FALSE;
    }
    LogDebugInfo("MTRR memory map built successfully");
  }
  if (!PoolManagerInitialize()) {
    LogError("Err, could not initialize pool manager");
    return FALSE;
  }
  if (!EptLogicalProcessorInitialize()) {
    return FALSE;
  }
  BroadcastVmxVirtualizationAllCores();
  return TRUE;
}

BOOLEAN VmxPerformVirtualizationOnSpecificCore() {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  VIRTUAL_MACHINE_STATE *VCpu = &g_GuestState[CurrentCore];
  LogDebugInfo("Allocating vmx regions for logical core %d", CurrentCore);
  AsmEnableVmxOperation();
  VmxFixCr4AndCr0Bits();
  LogDebugInfo("VMX-Operation enabled successfully");
  if (!VmxAllocateVmxonRegion(VCpu)) {
    LogError("Err, allocating memory for vmxon region was not successful");
    return FALSE;
  }
  if (!VmxAllocateVmcsRegion(VCpu)) {
    LogError("Err, allocating memory for vmcs region was not successful");
    return FALSE;
  }
  return TRUE;
}

VOID VmxFixCr4AndCr0Bits() {
  CR_FIXED CrFixed = {0};
  CR4 Cr4 = {0};
  CR0 Cr0 = {0};
  CrFixed.Flags = __readmsr(IA32_VMX_CR0_FIXED0);
  Cr0.AsUInt = __readcr0();
  Cr0.AsUInt |= CrFixed.Fields.Low;
  CrFixed.Flags = __readmsr(IA32_VMX_CR0_FIXED1);
  Cr0.AsUInt &= CrFixed.Fields.Low;
  __writecr0(Cr0.AsUInt);
  CrFixed.Flags = __readmsr(IA32_VMX_CR4_FIXED0);
  Cr4.AsUInt = __readcr4();
  Cr4.AsUInt |= CrFixed.Fields.Low;
  CrFixed.Flags = __readmsr(IA32_VMX_CR4_FIXED1);
  Cr4.AsUInt &= CrFixed.Fields.Low;
  __writecr4(Cr4.AsUInt);
}

BOOLEAN VmxCheckIsOnVmxRoot() {
  UINT64 VmcsLink = 0;
  __try {
    if (!__vmx_vmread(VMCS_GUEST_VMCS_LINK_POINTER, &VmcsLink)) {
      if (VmcsLink != 0) {
        return TRUE;
      }
    }
  } __except (1) {
  }
  return FALSE;
}

BOOLEAN VmxVirtualizeCurrentSystem(PVOID GuestStack) {
  UINT64 ErrorCode = 0;
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  VIRTUAL_MACHINE_STATE *VCpu = &g_GuestState[CurrentCore];
  LogDebugInfo("Virtualizing current system (logical core : 0x%x)",
               CurrentCore);
#if USE_DEFAULT_OS_IDT_AS_HOST_IDT == FALSE IdtEmulationPrepareHostIdt(VCpu);
#endif
#if USE_DEFAULT_OS_GDT_AS_HOST_GDT == FALSE
  SegmentPrepareHostGdt((SEGMENT_DESCRIPTOR_32 *)AsmGetGdtBase(),
                        AsmGetGdtLimit(), AsmGetTr(), VCpu->HostInterruptStack,
                        (SEGMENT_DESCRIPTOR_32 *)VCpu->HostGdt,
                        (TASK_STATE_SEGMENT_64 *)VCpu->HostTss);
#endif
  if (!VmxClearVmcsState(VCpu)) {
    LogError("Err, failed to clear vmcs");
    return FALSE;
  }
  if (!VmxLoadVmcs(VCpu)) {
    LogError("Err, failed to load vmcs");
    return FALSE;
  }
  LogDebugInfo("Setting up VMCS for current logical core");
  VmxSetupVmcs(VCpu, GuestStack);
  LogDebugInfo("Executing VMLAUNCH on logical core %d", CurrentCore);
  VCpu->HasLaunched = TRUE;
  __vmx_vmlaunch();
  VCpu->HasLaunched = FALSE;
  __vmx_vmread(VMCS_VM_INSTRUCTION_ERROR, &ErrorCode);
  LogError("Err, unable to execute VMLAUNCH, status : 0x%llx", ErrorCode);
  __vmx_off();
  LogError("Err, VMXOFF Executed Successfully but it was because of an error");
  return FALSE;
}

BOOLEAN VmxTerminate() {
  NTSTATUS Status = STATUS_SUCCESS;
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  VIRTUAL_MACHINE_STATE *VCpu = &g_GuestState[CurrentCore];
  Status = AsmVmxVmcall(VMCALL_VMXOFF, NULL64_ZERO, NULL64_ZERO, NULL64_ZERO);
  if (Status == STATUS_SUCCESS) {
    LogDebugInfo("VMX terminated on logical core %d\n", CurrentCore);
    MmFreeContiguousMemory((PVOID)VCpu->VmxonRegionVirtualAddress);
    MmFreeContiguousMemory((PVOID)VCpu->VmcsRegionVirtualAddress);
    PlatformMemFreePool((PVOID)VCpu->VmmStack);
    PlatformMemFreePool((PVOID)VCpu->MsrBitmapVirtualAddress);
    PlatformMemFreePool((PVOID)VCpu->IoBitmapVirtualAddressA);
    PlatformMemFreePool((PVOID)VCpu->IoBitmapVirtualAddressB);
#if USE_DEFAULT_OS_IDT_AS_HOST_IDT == FALSE
    PlatformMemFreePool((PVOID)VCpu->HostIdt);
#endif
#if USE_DEFAULT_OS_GDT_AS_HOST_GDT == FALSE
    PlatformMemFreePool((PVOID)VCpu->HostGdt);
    PlatformMemFreePool((PVOID)VCpu->HostTss);
#endif
#if USE_INTERRUPT_STACK_TABLE == TRUE
    PlatformMemFreePool((PVOID)VCpu->HostInterruptStack);
#endif
    return TRUE;
  }
  return FALSE;
}

VOID VmxVmptrst() {
  PHYSICAL_ADDRESS VmcsPhysicalAddr;
  VmcsPhysicalAddr.QuadPart = 0;
  __vmx_vmptrst((unsigned __int64 *)&VmcsPhysicalAddr);
  LogDebugInfo("VMPTRST result : %llx", VmcsPhysicalAddr);
}

_Use_decl_annotations_

    BOOLEAN
    VmxClearVmcsState(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT8 VmclearStatus;
  VmclearStatus = __vmx_vmclear(&VCpu->VmcsRegionPhysicalAddress);
  LogDebugInfo("VMCS VMCLEAR status : 0x%x", VmclearStatus);
  if (VmclearStatus) {
    LogDebugInfo("VMCS failed to clear, status : 0x%x", VmclearStatus);
    __vmx_off();
    return FALSE;
  }
  return TRUE;
}

_Use_decl_annotations_

    BOOLEAN
    VmxLoadVmcs(VIRTUAL_MACHINE_STATE *VCpu) {
  int VmptrldStatus;
  VmptrldStatus = __vmx_vmptrld(&VCpu->VmcsRegionPhysicalAddress);
  if (VmptrldStatus) {
    LogDebugInfo("VMCS failed to load, status : 0x%x", VmptrldStatus);
    return FALSE;
  }
  return TRUE;
}

_Use_decl_annotations_

    BOOLEAN
    VmxSetupVmcs(VIRTUAL_MACHINE_STATE *VCpu, PVOID GuestStack) {
  UINT32 CpuBasedVmExecControls;
  UINT32 SecondaryProcBasedVmExecControls;
  PVOID HostRsp;
  UINT64 GdtBase = 0;
  IA32_VMX_BASIC_REGISTER VmxBasicMsr = {0};
  VMX_SEGMENT_SELECTOR SegmentSelector = {0};
  VmxBasicMsr.AsUInt = __readmsr(IA32_VMX_BASIC);
  VmxVmwrite64(VMCS_HOST_ES_SELECTOR, AsmGetEs() & 0xF8);
  VmxVmwrite64(VMCS_HOST_CS_SELECTOR, AsmGetCs() & 0xF8);
  VmxVmwrite64(VMCS_HOST_SS_SELECTOR, AsmGetSs() & 0xF8);
  VmxVmwrite64(VMCS_HOST_DS_SELECTOR, AsmGetDs() & 0xF8);
  VmxVmwrite64(VMCS_HOST_FS_SELECTOR, AsmGetFs() & 0xF8);
  VmxVmwrite64(VMCS_HOST_GS_SELECTOR, AsmGetGs() & 0xF8);
  VmxVmwrite64(VMCS_HOST_TR_SELECTOR, AsmGetTr() & 0xF8);
  VmxVmwrite64(VMCS_GUEST_VMCS_LINK_POINTER, ~0ULL);
  VmxVmwrite64(VMCS_GUEST_DEBUGCTL, __readmsr(IA32_DEBUGCTL) & 0xFFFFFFFF);
  VmxVmwrite64(VMCS_GUEST_DEBUGCTL_HIGH, __readmsr(IA32_DEBUGCTL) >> 32);
  VmxVmwrite64(VMCS_CTRL_TSC_OFFSET, 0);
  VmxVmwrite64(VMCS_CTRL_PAGEFAULT_ERROR_CODE_MASK, 0);
  VmxVmwrite64(VMCS_CTRL_PAGEFAULT_ERROR_CODE_MATCH, 0);
  VmxVmwrite64(VMCS_CTRL_VMEXIT_MSR_STORE_COUNT, 0);
  VmxVmwrite64(VMCS_CTRL_VMEXIT_MSR_LOAD_COUNT, 0);
  VmxVmwrite64(VMCS_CTRL_VMENTRY_MSR_LOAD_COUNT, 0);
  VmxVmwrite64(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, 0);
  GdtBase = AsmGetGdtBase();
  HvFillGuestSelectorData((PVOID)GdtBase, ES, AsmGetEs());
  HvFillGuestSelectorData((PVOID)GdtBase, CS, AsmGetCs());
  HvFillGuestSelectorData((PVOID)GdtBase, SS, AsmGetSs());
  HvFillGuestSelectorData((PVOID)GdtBase, DS, AsmGetDs());
  HvFillGuestSelectorData((PVOID)GdtBase, FS, AsmGetFs());
  HvFillGuestSelectorData((PVOID)GdtBase, GS, AsmGetGs());
  HvFillGuestSelectorData((PVOID)GdtBase, LDTR, AsmGetLdtr());
  HvFillGuestSelectorData((PVOID)GdtBase, TR, AsmGetTr());
  VmxVmwrite64(VMCS_GUEST_FS_BASE, __readmsr(IA32_FS_BASE));
  VmxVmwrite64(VMCS_GUEST_GS_BASE, __readmsr(IA32_GS_BASE));
  CpuBasedVmExecControls = HvAdjustControls(
      CPU_BASED_ACTIVATE_IO_BITMAP | CPU_BASED_ACTIVATE_MSR_BITMAP |
          CPU_BASED_ACTIVATE_SECONDARY_CONTROLS,
      VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_PROCBASED_CTLS
                              : IA32_VMX_PROCBASED_CTLS);
  VmxVmwrite64(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               CpuBasedVmExecControls);
  LogDebugInfo("CPU Based VM Exec Controls (Based on %s) : 0x%x",
               VmxBasicMsr.VmxControls ? "IA32_VMX_TRUE_PROCBASED_CTLS"
                                       : "IA32_VMX_PROCBASED_CTLS",
               CpuBasedVmExecControls);
  SecondaryProcBasedVmExecControls = HvAdjustControls(
      CPU_BASED_CTL2_RDTSCP | CPU_BASED_CTL2_ENABLE_EPT |
          CPU_BASED_CTL2_ENABLE_INVPCID | CPU_BASED_CTL2_ENABLE_XSAVE_XRSTORS |
          CPU_BASED_CTL2_ENABLE_VPID,
      IA32_VMX_PROCBASED_CTLS2);
  VmxVmwrite64(VMCS_CTRL_SECONDARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS,
               SecondaryProcBasedVmExecControls);
  LogDebugInfo(
      "Secondary Proc Based VM Exec Controls (IA32_VMX_PROCBASED_CTLS2) : 0x%x",
      SecondaryProcBasedVmExecControls);
  VmxVmwrite64(
      VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS,
      HvAdjustControls(0, VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_PINBASED_CTLS
                                                  : IA32_VMX_PINBASED_CTLS));
  VmxVmwrite64(
      VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS,
      HvAdjustControls(VM_EXIT_HOST_ADDR_SPACE_SIZE,
                       VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_EXIT_CTLS
                                               : IA32_VMX_EXIT_CTLS));
  VmxVmwrite64(
      VMCS_CTRL_VMENTRY_CONTROLS,
      HvAdjustControls(VM_ENTRY_IA32E_MODE, VmxBasicMsr.VmxControls
                                                ? IA32_VMX_TRUE_ENTRY_CTLS
                                                : IA32_VMX_ENTRY_CTLS));
  VmxVmwrite64(VMCS_CTRL_CR0_GUEST_HOST_MASK, 0);
  VmxVmwrite64(VMCS_CTRL_CR4_GUEST_HOST_MASK, 0);
  VmxVmwrite64(VMCS_CTRL_CR0_READ_SHADOW, 0);
  VmxVmwrite64(VMCS_CTRL_CR4_READ_SHADOW, 0);
  VmxVmwrite64(VMCS_GUEST_CR0, __readcr0());
  VmxVmwrite64(VMCS_GUEST_CR3, __readcr3());
  VmxVmwrite64(VMCS_GUEST_CR4, __readcr4());
  VmxVmwrite64(VMCS_GUEST_DR7, 0x400);
  VmxVmwrite64(VMCS_HOST_CR0, __readcr0());
  VmxVmwrite64(VMCS_HOST_CR4, __readcr4());
  VmxVmwrite64(VMCS_HOST_CR3, LayoutGetSystemDirectoryTableBase());
  VmxVmwrite64(VMCS_GUEST_GDTR_BASE, AsmGetGdtBase());
  VmxVmwrite64(VMCS_GUEST_IDTR_BASE, AsmGetIdtBase());
  VmxVmwrite64(VMCS_GUEST_GDTR_LIMIT, AsmGetGdtLimit());
  VmxVmwrite64(VMCS_GUEST_IDTR_LIMIT, AsmGetIdtLimit());
  VmxVmwrite64(VMCS_GUEST_RFLAGS, AsmGetRflags());
  VmxVmwrite64(VMCS_GUEST_SYSENTER_CS, __readmsr(IA32_SYSENTER_CS));
  VmxVmwrite64(VMCS_GUEST_SYSENTER_EIP, __readmsr(IA32_SYSENTER_EIP));
  VmxVmwrite64(VMCS_GUEST_SYSENTER_ESP, __readmsr(IA32_SYSENTER_ESP));
#if USE_DEFAULT_OS_GDT_AS_HOST_GDT == FALSE
  SegmentGetDescriptor((PUCHAR)VCpu->HostGdt, AsmGetTr(), &SegmentSelector);
  VmxVmwrite64(VMCS_HOST_TR_BASE, SegmentSelector.Base);
  VmxVmwrite64(VMCS_HOST_GDTR_BASE, VCpu->HostGdt);
#else
  SegmentGetDescriptor((PUCHAR)AsmGetGdtBase(), AsmGetTr(), &SegmentSelector);
  VmxVmwrite64(VMCS_HOST_TR_BASE, SegmentSelector.Base);
  VmxVmwrite64(VMCS_HOST_GDTR_BASE, AsmGetGdtBase());
#endif
  VmxVmwrite64(VMCS_HOST_FS_BASE, __readmsr(IA32_FS_BASE));
  VmxVmwrite64(VMCS_HOST_GS_BASE, __readmsr(IA32_GS_BASE));
#if USE_DEFAULT_OS_IDT_AS_HOST_IDT == \
    FALSE VmxVmwrite64(VMCS_HOST_IDTR_BASE, VCpu->HostIdt);
#else
  VmxVmwrite64(VMCS_HOST_IDTR_BASE, AsmGetIdtBase());
#endif
  VmxVmwrite64(VMCS_HOST_SYSENTER_CS, __readmsr(IA32_SYSENTER_CS));
  VmxVmwrite64(VMCS_HOST_SYSENTER_EIP, __readmsr(IA32_SYSENTER_EIP));
  VmxVmwrite64(VMCS_HOST_SYSENTER_ESP, __readmsr(IA32_SYSENTER_ESP));
  VmxVmwrite64(VMCS_CTRL_MSR_BITMAP_ADDRESS, VCpu->MsrBitmapPhysicalAddress);
  VmxVmwrite64(VMCS_CTRL_IO_BITMAP_A_ADDRESS, VCpu->IoBitmapPhysicalAddressA);
  VmxVmwrite64(VMCS_CTRL_IO_BITMAP_B_ADDRESS, VCpu->IoBitmapPhysicalAddressB);
  VmxVmwrite64(VMCS_CTRL_EPT_POINTER, VCpu->EptPointer.AsUInt);
  VmxVmwrite64(VIRTUAL_PROCESSOR_ID, VPID_TAG);
  VmxVmwrite64(VMCS_GUEST_RSP, (UINT64)GuestStack);
  VmxVmwrite64(VMCS_GUEST_RIP, (UINT64)AsmVmxRestoreState);
  HostRsp = (PVOID)((UINT64)VCpu->VmmStack + VMM_STACK_SIZE - 1);
  HostRsp = ((PVOID)((ULONG_PTR)(HostRsp) & ~(16 - 1)));
  VmxVmwrite64(VMCS_HOST_RSP, (UINT64)HostRsp);
  VmxVmwrite64(VMCS_HOST_RIP, (UINT64)AsmVmexitHandler);
  return TRUE;
}

VOID VmxVmresume() {
  UINT64 ErrorCode = 0;
  __vmx_vmresume();
  __vmx_vmread(VMCS_VM_INSTRUCTION_ERROR, &ErrorCode);
  __vmx_off();
  LogError(
      "Err, in executing VMRESUME, status : 0x%llx, last VM-exit reason: 0x%x",
      ErrorCode, g_GuestState[KeGetCurrentProcessorNumberEx(NULL)].ExitReason);
}

UINT64 VmxVmfunc(UINT32 EptpIndex, UINT32 Function) {
  return AsmVmfunc(EptpIndex, Function);
}

VOID VmxVmxoff(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT64 GuestRSP = 0;
  UINT64 GuestRIP = 0;
  UINT64 GuestCr3 = 0;
  UINT64 ExitInstructionLength = 0;
  __vmx_vmread(VMCS_GUEST_CR3, &GuestCr3);
  __writecr3(GuestCr3);
  __vmx_vmread(VMCS_GUEST_RIP, &GuestRIP);
  __vmx_vmread(VMCS_GUEST_RSP, &GuestRSP);
  __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstructionLength);
  GuestRIP += ExitInstructionLength;
  VCpu->VmxoffState.GuestRip = GuestRIP;
  VCpu->VmxoffState.GuestRsp = GuestRSP;
  VCpu->VmxoffState.IsVmxoffExecuted = TRUE;
  HvRestoreRegisters();
  VmxClearVmcsState(VCpu);
  __vmx_off();
  VCpu->HasLaunched = FALSE;
  __writecr4(__readcr4() & (~X86_CR4_VMXE));
}

UINT64 VmxReturnStackPointerForVmxoff() {
  return g_GuestState[KeGetCurrentProcessorNumberEx(NULL)].VmxoffState.GuestRsp;
}

UINT64 VmxReturnInstructionPointerForVmxoff() {
  return g_GuestState[KeGetCurrentProcessorNumberEx(NULL)].VmxoffState.GuestRip;
}

VOID VmxPerformTermination() {
  ULONG ProcessorsCount;
  LogDebugInfo("Terminating VMX...\n");
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  g_TransparentMode = FALSE;
  EptHookUnHookAll();
  ExecTrapUninitialize();
  KeGenericCallDpc(DpcRoutineTerminateGuest, 0x0);
  PlatformMemFreePool(g_MsrBitmapInvalidMsrs);
  g_MsrBitmapInvalidMsrs = NULL;
  for (size_t i = 0; i < ProcessorsCount; i++) {
    if (g_GuestState[i].EptPageTable != NULL) {
      MmFreeContiguousMemory(g_GuestState[i].EptPageTable);
    }
    g_GuestState[i].EptPageTable = NULL;
  }
  PlatformMemFreePool(g_EptState);
  g_EptState = NULL;
  PoolManagerUninitialize();
  MemoryMapperUninitialize();
  GlobalGuestStateFreeMemory();
  LogDebugInfo("VMX operation turned off successfully");
}

UINT32 VmxCompatibleStrlen(const CHAR *S) {
  CHAR Temp = NULL_ZERO;
  UINT32 Count = 0;
  UINT64 AlignedAddress;
  CR3_TYPE GuestCr3;
  CR3_TYPE OriginalCr3;
  AlignedAddress = (UINT64)PAGE_ALIGN((UINT64)S);
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  OriginalCr3.Flags = __readcr3();
  __writecr3(GuestCr3.Flags);
  if (!CheckAccessValidityAndSafety(AlignedAddress, sizeof(CHAR))) {
    __writecr3(OriginalCr3.Flags);
    return 0;
  }
  while (TRUE) {
    MemoryMapperReadMemorySafe((UINT64)S, &Temp, sizeof(CHAR));
    if (Temp != '\0') {
      Count++;
      S++;
    } else {
      __writecr3(OriginalCr3.Flags);
      return Count;
    }
    if (!((UINT64)S & (PAGE_SIZE - 1))) {
      if (!CheckAccessValidityAndSafety((UINT64)S, sizeof(CHAR))) {
        __writecr3(OriginalCr3.Flags);
        return 0;
      }
    }
  }
  __writecr3(OriginalCr3.Flags);
}

UINT32 VmxCompatibleWcslen(const wchar_t *S) {
  wchar_t Temp = NULL_ZERO;
  UINT32 Count = 0;
  UINT64 AlignedAddress;
  CR3_TYPE GuestCr3;
  CR3_TYPE OriginalCr3;
  AlignedAddress = (UINT64)PAGE_ALIGN((UINT64)S);
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  OriginalCr3.Flags = __readcr3();
  __writecr3(GuestCr3.Flags);
  AlignedAddress = (UINT64)PAGE_ALIGN((UINT64)S);
  if (!CheckAccessValidityAndSafety(AlignedAddress, sizeof(wchar_t))) {
    __writecr3(OriginalCr3.Flags);
    return 0;
  }
  while (TRUE) {
    MemoryMapperReadMemorySafe((UINT64)S, &Temp, sizeof(wchar_t));
    if (Temp != '\0\0') {
      Count++;
      S++;
    } else {
      __writecr3(OriginalCr3.Flags);
      return Count;
    }
    if (!((UINT64)S & (PAGE_SIZE - 1))) {
      if (!CheckAccessValidityAndSafety((UINT64)S, sizeof(wchar_t))) {
        __writecr3(OriginalCr3.Flags);
        return 0;
      }
    }
  }
  __writecr3(OriginalCr3.Flags);
}

INT32 VmxCompatibleStrcmp(const CHAR *Address1, const CHAR *Address2,
                          SIZE_T Num, BOOLEAN IsStrncmp) {
  CHAR C1 = NULL_ZERO, C2 = NULL_ZERO;
  INT32 Result = 0;
  UINT32 Count = 0;
  UINT64 AlignedAddress1, AlignedAddress2;
  CR3_TYPE GuestCr3;
  CR3_TYPE OriginalCr3;
  AlignedAddress1 = (UINT64)PAGE_ALIGN((UINT64)Address1);
  AlignedAddress2 = (UINT64)PAGE_ALIGN((UINT64)Address2);
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  OriginalCr3.Flags = __readcr3();
  __writecr3(GuestCr3.Flags);
  if (!CheckAccessValidityAndSafety(AlignedAddress1, sizeof(CHAR)) ||
      !CheckAccessValidityAndSafety(AlignedAddress2, sizeof(CHAR))) {
    __writecr3(OriginalCr3.Flags);
    return 0x2;
  }
  do {
    if (IsStrncmp) {
      if (Count == Num) {
        break;
      } else {
        Count++;
      }
    }

    MemoryMapperReadMemorySafe((UINT64)Address1, &C1, sizeof(CHAR));

    MemoryMapperReadMemorySafe((UINT64)Address2, &C2, sizeof(CHAR));
    Address1++;
    Address2++;
    if (!((UINT64)AlignedAddress1 & (PAGE_SIZE - 1))) {
      if (!CheckAccessValidityAndSafety((UINT64)AlignedAddress1,
                                        sizeof(CHAR))) {
        __writecr3(OriginalCr3.Flags);
        return 0x2;
      }
    }
    if (!((UINT64)AlignedAddress2 & (PAGE_SIZE - 1))) {
      if (!CheckAccessValidityAndSafety((UINT64)AlignedAddress2,
                                        sizeof(CHAR))) {
        __writecr3(OriginalCr3.Flags);
        return 0x2;
      }
    }
    Result = C1 - C2;
  } while (!Result && C2);
  if (Result < 0) {
    Result = -1;
  } else if (Result > 0) {
    Result = 1;
  }
  __writecr3(OriginalCr3.Flags);
  return Result;
}

INT32 VmxCompatibleWcscmp(const wchar_t *Address1, const wchar_t *Address2,
                          SIZE_T Num, BOOLEAN IsWcsncmp) {
  wchar_t C1 = NULL_ZERO, C2 = NULL_ZERO;
  INT32 Result = 0;
  UINT32 Count = 0;
  UINT64 AlignedAddress1, AlignedAddress2;
  CR3_TYPE GuestCr3;
  CR3_TYPE OriginalCr3;
  AlignedAddress1 = (UINT64)PAGE_ALIGN((UINT64)Address1);
  AlignedAddress2 = (UINT64)PAGE_ALIGN((UINT64)Address2);
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  OriginalCr3.Flags = __readcr3();
  __writecr3(GuestCr3.Flags);
  if (!CheckAccessValidityAndSafety(AlignedAddress1, sizeof(wchar_t)) ||
      !CheckAccessValidityAndSafety(AlignedAddress2, sizeof(wchar_t))) {
    __writecr3(OriginalCr3.Flags);
    return 0x2;
  }
  do {
    if (IsWcsncmp) {
      if (Count == Num) {
        break;
      } else {
        Count++;
      }
    }

    MemoryMapperReadMemorySafe((UINT64)Address1, &C1, sizeof(wchar_t));

    MemoryMapperReadMemorySafe((UINT64)Address2, &C2, sizeof(wchar_t));
    Address1++;
    Address2++;
    if (!((UINT64)AlignedAddress1 & (PAGE_SIZE - 1))) {
      if (!CheckAccessValidityAndSafety((UINT64)AlignedAddress1,
                                        sizeof(wchar_t))) {
        __writecr3(OriginalCr3.Flags);
        return 0x2;
      }
    }
    if (!((UINT64)AlignedAddress2 & (PAGE_SIZE - 1))) {
      if (!CheckAccessValidityAndSafety((UINT64)AlignedAddress2,
                                        sizeof(wchar_t))) {
        __writecr3(OriginalCr3.Flags);
        return 0x2;
      }
    }
    Result = C1 - C2;
  } while (!Result && C2);
  if (Result < 0) {
    Result = -1;
  } else if (Result > 0) {
    Result = 1;
  }
  __writecr3(OriginalCr3.Flags);
  return Result;
}

INT32 VmxCompatibleMemcmp(const CHAR *Address1, const CHAR *Address2,
                          size_t Count) {
  CHAR C1 = NULL_ZERO, C2 = NULL_ZERO;
  INT32 Result = 0;
  UINT64 AlignedAddress1, AlignedAddress2;
  CR3_TYPE GuestCr3;
  CR3_TYPE OriginalCr3;
  AlignedAddress1 = (UINT64)PAGE_ALIGN((UINT64)Address1);
  AlignedAddress2 = (UINT64)PAGE_ALIGN((UINT64)Address2);
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  OriginalCr3.Flags = __readcr3();
  __writecr3(GuestCr3.Flags);
  if (!CheckAccessValidityAndSafety(AlignedAddress1, sizeof(wchar_t)) ||
      !CheckAccessValidityAndSafety(AlignedAddress2, sizeof(wchar_t))) {
    __writecr3(OriginalCr3.Flags);
    return 0x2;
  }
  while (Count-- > 0 && !Result) {
    MemoryMapperReadMemorySafe((UINT64)Address1, &C1, sizeof(CHAR));

    MemoryMapperReadMemorySafe((UINT64)Address2, &C2, sizeof(CHAR));
    Address1++;
    Address2++;
    if (!((UINT64)AlignedAddress1 & (PAGE_SIZE - 1))) {
      if (!CheckAccessValidityAndSafety((UINT64)AlignedAddress1,
                                        sizeof(wchar_t))) {
        __writecr3(OriginalCr3.Flags);
        return 0x2;
      }
    }
    if (!((UINT64)AlignedAddress2 & (PAGE_SIZE - 1))) {
      if (!CheckAccessValidityAndSafety((UINT64)AlignedAddress2,
                                        sizeof(wchar_t))) {
        __writecr3(OriginalCr3.Flags);
        return 0x2;
      }
    }
    Result = C1 - C2;
  }
  if (Result < 0) {
    Result = -1;
  } else if (Result > 0) {
    Result = 1;
  }
  __writecr3(OriginalCr3.Flags);
  return Result;
}
