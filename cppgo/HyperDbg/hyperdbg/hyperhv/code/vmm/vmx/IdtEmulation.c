
#include "pch.h"

VOID IdtEmulationQueryIdtEntriesRequest(
    PINTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS IdtQueryRequest,
    BOOLEAN ReadFromVmxRoot) {
  SIDT_ENTRY IdtrReg;
  KIDT_ENTRY *IdtEntries;
  if (!ReadFromVmxRoot) {
    __sidt(&IdtrReg);
    IdtEntries = (KIDT_ENTRY *)IdtrReg.IdtBase;
  } else {
    IdtEntries = (KIDT_ENTRY *)GetGuestIdtr();
  }
  for (UINT32 i = 0; i < MAX_NUMBER_OF_IDT_ENTRIES; i++) {
    IdtQueryRequest->IdtEntry[i] =
        (UINT64)((unsigned long long)IdtEntries[i].HighestPart << 32) |
        ((unsigned long long)IdtEntries[i].HighPart << 16) |
        (unsigned long long)IdtEntries[i].LowPart;
  }
}

VOID IdtEmulationCreateInterruptGate(
    PVOID Handler, SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64 *Entry) {
  UINT64 Offset = (UINT64)Handler;
  Entry->OffsetLow = (Offset >> 0) & 0xFFFF;
  Entry->OffsetMiddle = (Offset >> 16) & 0xFFFF;
  Entry->OffsetHigh = (Offset >> 32) & 0xFFFFFFFF;
#if USE_INTERRUPT_STACK_TABLE == TRUE
  Entry->InterruptStackTable = 1;
#else
  Entry->InterruptStackTable = 0;
#endif
}

VOID IdtEmulationPrepareHostIdt(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64 *VmxHostIdt =
      (SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64 *)VCpu->HostIdt;
  SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64 *WindowsIdt =
      (SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64 *)AsmGetIdtBase();
  RtlZeroMemory(VmxHostIdt, HOST_IDT_DESCRIPTOR_COUNT *
                                sizeof(SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64));
  RtlCopyBytes(
      VmxHostIdt, WindowsIdt,
      HOST_IDT_DESCRIPTOR_COUNT * sizeof(SEGMENT_DESCRIPTOR_INTERRUPT_GATE_64));

  IdtEmulationCreateInterruptGate((PVOID)InterruptHandler0,
                                  &VmxHostIdt[EXCEPTION_VECTOR_DIVIDE_ERROR]);
  IdtEmulationCreateInterruptGate((PVOID)InterruptHandler2,
                                  &VmxHostIdt[EXCEPTION_VECTOR_NMI]);
  IdtEmulationCreateInterruptGate((PVOID)InterruptHandler4,
                                  &VmxHostIdt[EXCEPTION_VECTOR_OVERFLOW]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler5,
      &VmxHostIdt[EXCEPTION_VECTOR_BOUND_RANGE_EXCEEDED]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler6, &VmxHostIdt[EXCEPTION_VECTOR_UNDEFINED_OPCODE]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler7,
      &VmxHostIdt[EXCEPTION_VECTOR_NO_MATH_COPROCESSOR]);
  IdtEmulationCreateInterruptGate((PVOID)InterruptHandler8,
                                  &VmxHostIdt[EXCEPTION_VECTOR_DOUBLE_FAULT]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler10,
      &VmxHostIdt[EXCEPTION_VECTOR_INVALID_TASK_SEGMENT_SELECTOR]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler11,
      &VmxHostIdt[EXCEPTION_VECTOR_SEGMENT_NOT_PRESENT]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler12,
      &VmxHostIdt[EXCEPTION_VECTOR_STACK_SEGMENT_FAULT]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler13,
      &VmxHostIdt[EXCEPTION_VECTOR_GENERAL_PROTECTION_FAULT]);
  IdtEmulationCreateInterruptGate((PVOID)InterruptHandler16,
                                  &VmxHostIdt[EXCEPTION_VECTOR_MATH_FAULT]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler17, &VmxHostIdt[EXCEPTION_VECTOR_ALIGNMENT_CHECK]);
  IdtEmulationCreateInterruptGate((PVOID)InterruptHandler18,
                                  &VmxHostIdt[EXCEPTION_VECTOR_MACHINE_CHECK]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler19,
      &VmxHostIdt[EXCEPTION_VECTOR_SIMD_FLOATING_POINT_NUMERIC_ERROR]);
  IdtEmulationCreateInterruptGate(
      (PVOID)InterruptHandler20,
      &VmxHostIdt[EXCEPTION_VECTOR_VIRTUAL_EXCEPTION]);
  IdtEmulationCreateInterruptGate((PVOID)InterruptHandler30,
                                  &VmxHostIdt[EXCEPTION_VECTOR_RESERVED11]);
}

VOID IdtEmulationhandleHostInterrupt(
    _Inout_ INTERRUPT_TRAP_FRAME *IntrTrapFrame) {
  UINT64 PageFaultCr2;
  ULONG CurrentCore;
  BOOLEAN Interruptible;
  VMX_INTERRUPTIBILITY_STATE InterruptibilityState = {0};
  CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  VIRTUAL_MACHINE_STATE *VCpu = &g_GuestState[CurrentCore];
  VCpu->LastExceptionOccuredInHost = IntrTrapFrame->vector;
  switch (IntrTrapFrame->vector) {
    case EXCEPTION_VECTOR_NMI:
      if (!VmxBroadcastHandleNmiCallback((PVOID)IntrTrapFrame, FALSE)) {
        VmxVmread32P(VMCS_GUEST_INTERRUPTIBILITY_STATE,
                     &InterruptibilityState.AsUInt);
        Interruptible = !InterruptibilityState.BlockingByNmi;
        if (Interruptible) {
          EventInjectNmi(VCpu);
        } else {
          HvSetNmiWindowExiting(TRUE);
        }
      }
      break;
    case EXCEPTION_VECTOR_PAGE_FAULT:
      PageFaultCr2 = __readcr2();
      LogInfo(
          "Page-fault received, rip: %llx, rsp: %llx, error: %llx, CR2: %llx",
          IntrTrapFrame->rip, IntrTrapFrame->rsp, IntrTrapFrame->error,
          PageFaultCr2);
      break;
    default:
      LogInfo("Host exception, rip: %llx, rsp: %llx, error: %llx, vector: %x",
              IntrTrapFrame->rip, IntrTrapFrame->rsp, IntrTrapFrame->error,
              IntrTrapFrame->vector);
      break;
  }
}

VOID IdtEmulationHandlePageFaults(_Inout_ VIRTUAL_MACHINE_STATE *VCpu,
                                  _In_ VMEXIT_INTERRUPT_INFORMATION
                                      InterruptExit) {
  UINT32 ErrorCode = 0;
  PAGE_FAULT_EXCEPTION PageFaultErrorCode = {0};
  UINT64 PageFaultAddress = 0;
  VmxVmread32P(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &ErrorCode);
  PageFaultErrorCode.AsUInt = ErrorCode;
  __vmx_vmread(VMCS_EXIT_QUALIFICATION, &PageFaultAddress);
  if (!DebuggingCallbackConditionalPageFaultException(
          VCpu->CoreId, PageFaultAddress, PageFaultErrorCode.AsUInt)) {
    EventInjectPageFaults(VCpu, InterruptExit, PageFaultAddress,
                          PageFaultErrorCode);
  }
}

VOID IdtEmulationHandleExceptionAndNmi(_Inout_ VIRTUAL_MACHINE_STATE *VCpu,
                                       _In_ VMEXIT_INTERRUPT_INFORMATION
                                           InterruptExit) {
  switch (InterruptExit.Vector) {
    case EXCEPTION_VECTOR_BREAKPOINT: {
      UINT64 GuestRip = NULL64_ZERO;
      BYTE TargetMem = NULL_ZERO;
      __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
      MemoryMapperReadMemorySafe(GuestRip, &TargetMem, sizeof(BYTE));
      if (!EptCheckAndHandleBreakpoint(VCpu) || TargetMem == 0xcc) {
        if (!DebuggingCallbackHandleBreakpointException(VCpu->CoreId)) {
          HvSuppressRipIncrement(VCpu);
          EventInjectBreakpoint();
        }
      }
    } break;
    case EXCEPTION_VECTOR_UNDEFINED_OPCODE:
      if (!SyscallHookHandleUD(VCpu)) {
        EventInjectUndefinedOpcode(VCpu);
      }
      break;
    case EXCEPTION_VECTOR_PAGE_FAULT:
      IdtEmulationHandlePageFaults(VCpu, InterruptExit);
      break;
    case EXCEPTION_VECTOR_DEBUG_BREAKPOINT:
      if (g_TransparentMode &&
          TransparentCheckAndHandleAfterSyscallTrapFlags(
              VCpu, HANDLE_TO_UINT32(PsGetCurrentProcessId()),
              HANDLE_TO_UINT32(PsGetCurrentThreadId()))) {
      } else if (!DebuggingCallbackHandleDebugBreakpointException(
                     VCpu->CoreId)) {
        EventInjectInterruptOrException(InterruptExit);
      }
      break;
    case EXCEPTION_VECTOR_NMI:
      if (VCpu->EnableExternalInterruptsOnContinue ||
          VCpu->EnableExternalInterruptsOnContinueMtf ||
          VCpu->RegisterBreakOnMtf) {
      } else {
        EventInjectInterruptOrException(InterruptExit);
      }
      break;
    default:
      EventInjectInterruptOrException(InterruptExit);
      break;
  }
}

BOOLEAN IdtEmulationInjectInterruptWhenInterruptWindowIsOpen(
    _Inout_ VIRTUAL_MACHINE_STATE *VCpu,
    _In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit) {
  BOOLEAN FoundAPlaceForFutureInjection = FALSE;
  for (size_t i = 0; i < PENDING_INTERRUPTS_BUFFER_CAPACITY; i++) {
    if (VCpu->PendingExternalInterrupts[i] == NULL_ZERO) {
      VCpu->PendingExternalInterrupts[i] = InterruptExit.AsUInt;
      FoundAPlaceForFutureInjection = TRUE;
      break;
    }
  }
  return FoundAPlaceForFutureInjection;
}

VOID IdtEmulationHandleExternalInterrupt(_Inout_ VIRTUAL_MACHINE_STATE *VCpu,
                                         _In_ VMEXIT_INTERRUPT_INFORMATION
                                             InterruptExit) {
  BOOLEAN Interruptible = TRUE;
  VMX_INTERRUPTIBILITY_STATE InterruptibilityState = {0};
  RFLAGS GuestRflags = {0};
  if (VCpu->EnableExternalInterruptsOnContinue ||
      VCpu->EnableExternalInterruptsOnContinueMtf) {
    IdtEmulationInjectInterruptWhenInterruptWindowIsOpen(VCpu, InterruptExit);
    HvSuppressRipIncrement(VCpu);
  } else if (InterruptExit.Valid && InterruptExit.InterruptionType ==
                                        INTERRUPT_TYPE_EXTERNAL_INTERRUPT) {
    VmxVmread64P(VMCS_GUEST_RFLAGS, &GuestRflags.AsUInt);
    VmxVmread32P(VMCS_GUEST_INTERRUPTIBILITY_STATE,
                 &InterruptibilityState.AsUInt);
    Interruptible = GuestRflags.InterruptEnableFlag &&
                    !InterruptibilityState.BlockingByMovSs;
    if (Interruptible) {
      EventInjectInterruptOrException(InterruptExit);
    } else {
      IdtEmulationInjectInterruptWhenInterruptWindowIsOpen(VCpu, InterruptExit);
      HvSetInterruptWindowExiting(TRUE);
    }
    HvSuppressRipIncrement(VCpu);
  } else {
    Interruptible = FALSE;
    LogError(
        "Err, why we are here? it's a vm-exit due to the external"
        "interrupt and its type is not external interrupt? weird!");
  }
}

VOID IdtEmulationHandleNmiWindowExiting(_Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  EventInjectNmi(VCpu);
  HvSetNmiWindowExiting(FALSE);
}

BOOLEAN IdtEmulationInjectPageFaultWhenInterruptWindowsIsOpen(
    _Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  if (g_PageFaultInjectionAddressFrom > g_PageFaultInjectionAddressTo) {
    g_WaitingForInterruptWindowToInjectPageFault = FALSE;
    return FALSE;
  }
  EventInjectPageFaultWithCr2(VCpu, g_PageFaultInjectionAddressFrom,
                              g_PageFaultInjectionErrorCode);
  if (MemoryMapperCheckIfPdeIsLargePageOnTargetProcess(
          (PVOID)g_PageFaultInjectionAddressFrom)) {
    g_PageFaultInjectionAddressFrom =
        g_PageFaultInjectionAddressFrom + SIZE_2_MB;
  } else {
    g_PageFaultInjectionAddressFrom =
        g_PageFaultInjectionAddressFrom + PAGE_SIZE;
  }
  return TRUE;
}

VOID IdtEmulationHandleInterruptWindowExiting(
    _Inout_ VIRTUAL_MACHINE_STATE *VCpu) {
  VMEXIT_INTERRUPT_INFORMATION InterruptExit = {0};
  BOOLEAN InjectPageFault = FALSE;
  if (g_WaitingForInterruptWindowToInjectPageFault) {
    InjectPageFault =
        IdtEmulationInjectPageFaultWhenInterruptWindowsIsOpen(VCpu);
  }
  if (!InjectPageFault) {
    for (size_t i = 0; i < PENDING_INTERRUPTS_BUFFER_CAPACITY; i++) {
      if (VCpu->PendingExternalInterrupts[i] != NULL_ZERO) {
        InterruptExit.AsUInt = VCpu->PendingExternalInterrupts[i];
        VCpu->PendingExternalInterrupts[i] = NULL_ZERO;
        break;
      }
    }
    if (InterruptExit.AsUInt == 0) {
      HvSetInterruptWindowExiting(FALSE);
    } else {
      EventInjectInterruptOrException(InterruptExit);
    }
  }
  HvSuppressRipIncrement(VCpu);
}
