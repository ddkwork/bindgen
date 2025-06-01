
#include "pch.h"

VOID VmFuncPerformRipIncrement(UINT32 CoreId) {
  HvPerformRipIncrement(&g_GuestState[CoreId]);
}

VOID VmFuncSuppressRipIncrement(UINT32 CoreId) {
  HvSuppressRipIncrement(&g_GuestState[CoreId]);
}

VOID VmFuncChangeMtfUnsettingState(UINT32 CoreId, BOOLEAN Set) {
  g_GuestState[CoreId].IgnoreMtfUnset = Set;
}

VOID VmFuncChangeIgnoreOneMtfState(UINT32 CoreId, BOOLEAN Set) {
  g_GuestState[CoreId].IgnoreOneMtf = Set;
}

VOID VmFuncRegisterMtfBreak(UINT32 CoreId) {
  g_GuestState[CoreId].RegisterBreakOnMtf = TRUE;
}

VOID VmFuncUnRegisterMtfBreak(UINT32 CoreId) {
  g_GuestState[CoreId].RegisterBreakOnMtf = FALSE;
}

VOID VmFuncSetMonitorTrapFlag(BOOLEAN Set) { HvSetMonitorTrapFlag(Set); }

VOID VmFuncSetRflagTrapFlag(BOOLEAN Set) { HvSetRflagTrapFlag(Set); }

VOID VmFuncSetLoadDebugControls(BOOLEAN Set) { HvSetLoadDebugControls(Set); }

VOID VmFuncSetSaveDebugControls(BOOLEAN Set) { HvSetSaveDebugControls(Set); }

VOID VmFuncSetPmcVmexit(BOOLEAN Set) { HvSetPmcVmexit(Set); }

VOID VmFuncSetMovControlRegsExiting(BOOLEAN Set, UINT64 ControlRegister,
                                    UINT64 MaskRegister) {
  HvSetMovControlRegsExiting(Set, ControlRegister, MaskRegister);
}

VOID VmFuncSetMovToCr3Vmexit(UINT32 CoreId, BOOLEAN Set) {
  HvSetMovToCr3Vmexit(&g_GuestState[CoreId], Set);
}

VOID VmFuncWriteExceptionBitmap(UINT32 BitmapMask) {
  HvWriteExceptionBitmap(BitmapMask);
}

UINT32 VmFuncReadExceptionBitmap() { return HvReadExceptionBitmap(); }

VOID VmFuncSetInterruptWindowExiting(BOOLEAN Set) {
  HvSetInterruptWindowExiting(Set);
}

VOID VmFuncSetNmiWindowExiting(BOOLEAN Set) { HvSetNmiWindowExiting(Set); }

VOID VmFuncSetNmiExiting(BOOLEAN Set) { HvSetNmiExiting(Set); }

VOID VmFuncSetExceptionBitmap(UINT32 CoreId, UINT32 IdtIndex) {
  HvSetExceptionBitmap(&g_GuestState[CoreId], IdtIndex);
}

VOID VmFuncUnsetExceptionBitmap(UINT32 CoreId, UINT32 IdtIndex) {
  HvUnsetExceptionBitmap(&g_GuestState[CoreId], IdtIndex);
}

VOID VmFuncSetExternalInterruptExiting(UINT32 CoreId, BOOLEAN Set) {
  HvSetExternalInterruptExiting(&g_GuestState[CoreId], Set);
}

VOID VmFuncSetRdtscExiting(UINT32 CoreId, BOOLEAN Set) {
  HvSetRdtscExiting(&g_GuestState[CoreId], Set);
}

VOID VmFuncSetMovDebugRegsExiting(UINT32 CoreId, BOOLEAN Set) {
  HvSetMovDebugRegsExiting(&g_GuestState[CoreId], Set);
}

UINT64 VmFuncGetLastVmexitRip(UINT32 CoreId) {
  return g_GuestState[CoreId].LastVmexitRip;
}

VOID VmFuncInjectPendingExternalInterrupts(UINT32 CoreId) {
  HvInjectPendingExternalInterrupts(&g_GuestState[CoreId]);
}

UINT16 VmFuncGetCsSelector() { return HvGetCsSelector(); }

UINT64 VmFuncGetRflags() { return HvGetRflags(); }

VOID VmFuncSetRflags(UINT64 Rflags) { HvSetRflags(Rflags); }

UINT64 VmFuncGetRip() { return HvGetRip(); }

VOID VmFuncSetRip(UINT64 Rip) { HvSetRip(Rip); }

UINT64 VmFuncGetInterruptibilityState() { return HvGetInterruptibilityState(); }

UINT64 VmFuncClearSteppingBits(UINT64 Interruptibility) {
  return HvClearSteppingBits(Interruptibility);
}

VOID VmFuncSetInterruptibilityState(UINT64 InterruptibilityState) {
  HvSetInterruptibilityState(InterruptibilityState);
}

BOOLEAN VmFuncNmiBroadcastRequest(UINT32 CoreId) {
  return VmxBroadcastNmi(&g_GuestState[CoreId], NMI_BROADCAST_ACTION_REQUEST);
}

BOOLEAN VmFuncNmiBroadcastInvalidateEptSingleContext(UINT32 CoreId) {
  return VmxBroadcastNmi(
      &g_GuestState[CoreId],
      NMI_BROADCAST_ACTION_INVALIDATE_EPT_CACHE_SINGLE_CONTEXT);
}

BOOLEAN VmFuncNmiBroadcastInvalidateEptAllContexts(UINT32 CoreId) {
  return VmxBroadcastNmi(
      &g_GuestState[CoreId],
      NMI_BROADCAST_ACTION_INVALIDATE_EPT_CACHE_ALL_CONTEXTS);
}

VOID VmFuncInvalidateEptSingleContext(UINT32 CoreId) {
  EptInveptSingleContext(g_GuestState[CoreId].EptPointer.AsUInt);
}

VOID VmFuncInvalidateEptAllContexts() { EptInveptAllContexts(); }

VOID VmFuncCheckAndEnableExternalInterrupts(UINT32 CoreId) {
  HvCheckAndEnableExternalInterrupts(&g_GuestState[CoreId]);
}

VOID VmFuncDisableExternalInterruptsAndInterruptWindow(UINT32 CoreId) {
  HvDisableExternalInterruptsAndInterruptWindow(&g_GuestState[CoreId]);
}

BOOLEAN VmFuncInitVmm(VMM_CALLBACKS *VmmCallbacks) {
  return HvInitVmm(VmmCallbacks);
}

VOID VmFuncUninitVmm() { VmxPerformTermination(); }

BOOLEAN VmFuncVmxGetCurrentExecutionMode() {
  return VmxGetCurrentExecutionMode();
}

BOOLEAN VmFuncQueryModeExecTrap() { return g_ExecTrapInitialized; }

VOID VmFuncSetTriggerEventForVmcalls(BOOLEAN Set) {
  g_TriggerEventForVmcalls = Set;
}

VOID VmFuncSetTriggerEventForCpuids(BOOLEAN Set) {
  g_TriggerEventForCpuids = Set;
}

UINT32 VmFuncVmxCompatibleStrlen(const CHAR *s) {
  return VmxCompatibleStrlen(s);
}

UINT32 VmFuncVmxCompatibleWcslen(const wchar_t *s) {
  return VmxCompatibleWcslen(s);
}

VOID VmFuncEventInjectPageFaultWithCr2(UINT32 CoreId, UINT64 Address,
                                       UINT32 PageFaultCode) {
  EventInjectPageFaultWithCr2(&g_GuestState[CoreId], Address, PageFaultCode);
}

VOID VmFuncEventInjectPageFaultRangeAddress(UINT32 CoreId, UINT64 AddressFrom,
                                            UINT64 AddressTo,
                                            UINT32 PageFaultCode) {
  EventInjectPageFaultRangeAddress(&g_GuestState[CoreId], AddressFrom,
                                   AddressTo, PageFaultCode);
}

VOID VmFuncEventInjectInterruption(UINT32 InterruptionType, UINT32 Vector,
                                   BOOLEAN DeliverErrorCode, UINT32 ErrorCode) {
  EventInjectInterruption(InterruptionType, Vector, DeliverErrorCode,
                          ErrorCode);
}

NTSTATUS VmFuncVmxVmcall(unsigned long long VmcallNumber,
                         unsigned long long OptionalParam1,
                         unsigned long long OptionalParam2,
                         unsigned long long OptionalParam3) {
  return AsmVmxVmcall(VmcallNumber, OptionalParam1, OptionalParam2,
                      OptionalParam3);
}

VOID VmFuncVmxBroadcastInitialize() { VmxBroadcastInitialize(); }

VOID VmFuncVmxBroadcastUninitialize() { VmxBroadcastUninitialize(); }

VOID VmFuncEventInjectBreakpoint() { EventInjectBreakpoint(); }

INT32 VmFuncVmxCompatibleStrcmp(const CHAR *Address1, const CHAR *Address2) {
  return VmxCompatibleStrcmp(Address1, Address2, NULL_ZERO, FALSE);
}

INT32 VmFuncVmxCompatibleStrncmp(const CHAR *Address1, const CHAR *Address2,
                                 SIZE_T Num) {
  return VmxCompatibleStrcmp(Address1, Address2, Num, TRUE);
}

INT32 VmFuncVmxCompatibleWcscmp(const wchar_t *Address1,
                                const wchar_t *Address2) {
  return VmxCompatibleWcscmp(Address1, Address2, NULL_ZERO, FALSE);
}

INT32 VmFuncVmxCompatibleWcsncmp(const wchar_t *Address1,
                                 const wchar_t *Address2, SIZE_T Num) {
  return VmxCompatibleWcscmp(Address1, Address2, Num, TRUE);
}

INT32 VmFuncVmxCompatibleMemcmp(const CHAR *Address1, const CHAR *Address2,
                                size_t Count) {
  return VmxCompatibleMemcmp(Address1, Address2, Count);
}

VOID VmFuncEnableMtfAndChangeExternalInterruptState(UINT32 CoreId) {
  HvEnableMtfAndChangeExternalInterruptState(&g_GuestState[CoreId]);
}

VOID VmFuncEnableAndCheckForPreviousExternalInterrupts(UINT32 CoreId) {
  HvEnableAndCheckForPreviousExternalInterrupts(&g_GuestState[CoreId]);
}

BOOLEAN VmFuncApicStoreLocalApicFields(PLAPIC_PAGE LocalApicBuffer,
                                       PBOOLEAN IsUsingX2APIC) {
  return ApicStoreLocalApicFields(LocalApicBuffer, IsUsingX2APIC);
}

BOOLEAN VmFuncApicStoreIoApicFields(IO_APIC_ENTRY_PACKETS *IoApicPackets) {
  return ApicStoreIoApicFields(IoApicPackets);
}

VOID VmFuncIdtQueryEntries(
    PINTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS IdtQueryRequest,
    BOOLEAN ReadFromVmxRoot) {
  IdtEmulationQueryIdtEntriesRequest(IdtQueryRequest, ReadFromVmxRoot);
}
