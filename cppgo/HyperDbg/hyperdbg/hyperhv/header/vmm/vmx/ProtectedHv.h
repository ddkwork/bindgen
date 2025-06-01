
#pragma once
VOID ProtectedHvSetExceptionBitmap(VIRTUAL_MACHINE_STATE *VCpu,
                                   UINT32 IdtIndex);
VOID ProtectedHvUnsetExceptionBitmap(VIRTUAL_MACHINE_STATE *VCpu,
                                     UINT32 IdtIndex);
VOID ProtectedHvResetExceptionBitmapToClearEvents(VIRTUAL_MACHINE_STATE *VCpu);
VOID ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands(
    VIRTUAL_MACHINE_STATE *VCpu);
VOID ProtectedHvSetExternalInterruptExiting(VIRTUAL_MACHINE_STATE *VCpu,
                                            BOOLEAN Set);
VOID ProtectedHvExternalInterruptExitingForDisablingInterruptCommands(
    VIRTUAL_MACHINE_STATE *VCpu);
VOID ProtectedHvSetRdtscExiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set);
VOID ProtectedHvDisableRdtscExitingForDisablingTscCommands(
    VIRTUAL_MACHINE_STATE *VCpu);
VOID ProtectedHvSetMovDebugRegsExiting(VIRTUAL_MACHINE_STATE *VCpu,
                                       BOOLEAN Set);
VOID ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands(
    VIRTUAL_MACHINE_STATE *VCpu);
VOID ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands(
    VIRTUAL_MACHINE_STATE *VCpu, UINT64 ControlRegister, UINT64 MaskRegister);
VOID ProtectedHvSetMov2CrExiting(BOOLEAN Set, UINT64 ControlRegister,
                                 UINT64 MaskRegister);
VOID ProtectedHvSetMov2Cr3Exiting(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN Set);
