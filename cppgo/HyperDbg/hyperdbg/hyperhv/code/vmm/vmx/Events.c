
#include "pch.h"

VOID EventInjectInterruption(INTERRUPT_TYPE InterruptionType,
                             EXCEPTION_VECTORS Vector, BOOLEAN DeliverErrorCode,
                             UINT32 ErrorCode) {
  INTERRUPT_INFO Inject = {0};
  Inject.Fields.Valid = TRUE;
  Inject.Fields.InterruptType = InterruptionType;
  Inject.Fields.Vector = Vector;
  Inject.Fields.DeliverCode = DeliverErrorCode;
  VmxVmwrite64(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, Inject.Flags);
  if (DeliverErrorCode) {
    VmxVmwrite64(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);
  }
}

VOID EventInjectBreakpoint() {
  UINT32 ExitInstrLength;
  EventInjectInterruption(INTERRUPT_TYPE_SOFTWARE_EXCEPTION,
                          EXCEPTION_VECTOR_BREAKPOINT, FALSE, 0);
  VmxVmread32P(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstrLength);
  VmxVmwrite64(VMCS_CTRL_VMENTRY_INSTRUCTION_LENGTH, ExitInstrLength);
}

VOID EventInjectGeneralProtection() {
  UINT32 ExitInstrLength;
  EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION,
                          EXCEPTION_VECTOR_GENERAL_PROTECTION_FAULT, TRUE, 0);
  VmxVmread32P(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstrLength);
  VmxVmwrite64(VMCS_CTRL_VMENTRY_INSTRUCTION_LENGTH, ExitInstrLength);
}

VOID EventInjectUndefinedOpcode(VIRTUAL_MACHINE_STATE *VCpu) {
  EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION,
                          EXCEPTION_VECTOR_UNDEFINED_OPCODE, FALSE, 0);
  HvSuppressRipIncrement(VCpu);
}

VOID EventInjectNmi(VIRTUAL_MACHINE_STATE *VCpu) {
  EventInjectInterruption(INTERRUPT_TYPE_NMI, EXCEPTION_VECTOR_NMI, FALSE, 0);
  HvSuppressRipIncrement(VCpu);
}

VOID EventInjectDebugBreakpoint() {
  EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION,
                          EXCEPTION_VECTOR_DEBUG_BREAKPOINT, FALSE, 0);
}

VOID EventInjectPageFaultWithoutErrorCode(UINT64 PageFaultAddress) {
  PAGE_FAULT_EXCEPTION ErrorCode = {0};
  __writecr2(PageFaultAddress);
  ErrorCode.Execute = 0;
  ErrorCode.Present = 0;
  ErrorCode.UserModeAccess = 0;
  ErrorCode.Write = 0;
  EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION,
                          EXCEPTION_VECTOR_PAGE_FAULT, TRUE, ErrorCode.AsUInt);
}

VOID EventInjectInterruptOrException(
    _In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit) {
  UINT32 ErrorCode = 0;
  VmxVmwrite64(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD,
               InterruptExit.AsUInt);
  if (InterruptExit.ErrorCodeValid) {
    VmxVmread32P(VMCS_VMEXIT_INTERRUPTION_ERROR_CODE, &ErrorCode);
    VmxVmwrite64(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);
  }
}

VOID EventInjectPageFaults(_Inout_ VIRTUAL_MACHINE_STATE *VCpu,
                           _In_ VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                           _In_ UINT64 PageFaultAddress,
                           _In_ PAGE_FAULT_EXCEPTION PageFaultCode) {
  __writecr2(PageFaultAddress);
  HvSuppressRipIncrement(VCpu);
  VmxVmwrite64(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD,
               InterruptExit.AsUInt);
  if (InterruptExit.ErrorCodeValid) {
    VmxVmwrite64(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, PageFaultCode.AsUInt);
  }
}

VOID EventInjectPageFaultRangeAddress(VIRTUAL_MACHINE_STATE *VCpu,
                                      UINT64 AddressFrom, UINT64 AddressTo,
                                      UINT32 PageFaultCode) {
  UNREFERENCED_PARAMETER(VCpu);
  g_WaitingForInterruptWindowToInjectPageFault = TRUE;
  g_PageFaultInjectionAddressFrom = AddressFrom;
  g_PageFaultInjectionAddressTo = AddressTo;
  g_PageFaultInjectionErrorCode = PageFaultCode;
  HvSetInterruptWindowExiting(TRUE);
}

VOID EventInjectPageFaultWithCr2(VIRTUAL_MACHINE_STATE *VCpu, UINT64 Address,
                                 UINT32 PageFaultCode) {
  VMEXIT_INTERRUPT_INFORMATION InterruptInfo = {0};
  PAGE_FAULT_EXCEPTION PageFaultErrorCode = {0};
  InterruptInfo.Vector = EXCEPTION_VECTOR_PAGE_FAULT;
  InterruptInfo.InterruptionType = INTERRUPT_TYPE_HARDWARE_EXCEPTION;
  InterruptInfo.ErrorCodeValid = TRUE;
  InterruptInfo.NmiUnblocking = FALSE;
  InterruptInfo.Valid = TRUE;
  PageFaultErrorCode.AsUInt = PageFaultCode;
  EventInjectPageFaults(VCpu, InterruptInfo, Address, PageFaultErrorCode);
}
