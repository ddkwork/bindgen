
#pragma once

BOOLEAN HvSetGuestSelector(PVOID GdtBase, UINT32 SegmentRegister,
                           UINT16 Selector);

UINT32 HvAdjustControls(UINT32 Ctl, UINT32 Msr);

VOID HvHandleCpuid(VIRTUAL_MACHINE_STATE *VCpu);

VOID HvFillGuestSelectorData(PVOID GdtBase, UINT32 SegmentRegister,
                             UINT16 Selector);

VOID HvHandleControlRegisterAccess(
    VIRTUAL_MACHINE_STATE *VCpu,
    VMX_EXIT_QUALIFICATION_MOV_CR *CrExitQualification);

VOID HvResumeToNextInstruction();

extern inline VOID HvSuppressRipIncrement(VIRTUAL_MACHINE_STATE *VCpu);

extern inline VOID HvPerformRipIncrement(VIRTUAL_MACHINE_STATE *VCpu);

VOID HvSetMonitorTrapFlag(BOOLEAN Set);

VOID HvSetRflagTrapFlag(BOOLEAN Set);

VOID HvSetLoadDebugControls(BOOLEAN Set);

VOID HvSetSaveDebugControls(BOOLEAN Set);

VOID HvRestoreRegisters();

VOID HvSetPmcVmexit(BOOLEAN Set);

VOID HvSetMovControlRegsExiting(BOOLEAN Set, UINT64 ControlRegister,
                                UINT64 MaskRegister);

VOID HvSetMovToCr3Vmexit(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set);

VOID HvWriteExceptionBitmap(UINT32 BitmapMask);

UINT32 HvReadExceptionBitmap();

VOID HvSetInterruptWindowExiting(BOOLEAN Set);

VOID HvSetPmlEnableFlag(BOOLEAN Set);

VOID HvSetModeBasedExecutionEnableFlag(BOOLEAN Set);

VOID HvSetNmiWindowExiting(BOOLEAN Set);

VOID HvHandleMovDebugRegister(VIRTUAL_MACHINE_STATE *VCpu);

VOID HvSetMovDebugRegsExiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set);

VOID HvSetNmiExiting(BOOLEAN Set);

VOID HvSetVmxPreemptionTimerExiting(BOOLEAN Set);

VOID HvSetExceptionBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 IdtIndex);

VOID HvUnsetExceptionBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 IdtIndex);

VOID HvSetExternalInterruptExiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set);

VOID HvEnableAndCheckForPreviousExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu);

VOID HvSetRdtscExiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set);

UINT16 HvGetCsSelector();

UINT64 HvGetRflags();

VOID HvSetRflags(UINT64 Rflags);

UINT64 HvGetRip();

VOID HvSetRip(UINT64 Rip);

UINT64 HvGetInterruptibilityState();

UINT64 HvClearSteppingBits(UINT64 Interruptibility);

VOID HvSetInterruptibilityState(UINT64 InterruptibilityState);

VOID HvInjectPendingExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu);

VOID HvCheckAndEnableExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu);

VOID HvDisableExternalInterruptsAndInterruptWindow(VIRTUAL_MACHINE_STATE *VCpu);

BOOLEAN HvInitVmm(VMM_CALLBACKS *VmmCallbacks);

VOID HvEnableMtfAndChangeExternalInterruptState(VIRTUAL_MACHINE_STATE *VCpu);

VOID HvPreventExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu);

UINT64 HvGetPendingDebugExceptions();

VOID HvSetPendingDebugExceptions(UINT64 Value);

UINT64 HvGetDebugctl();

VOID HvSetDebugctl(UINT64 Value);

VOID HvHandleTrapFlag();
