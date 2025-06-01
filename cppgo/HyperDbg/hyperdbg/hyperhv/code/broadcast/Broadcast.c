
#include "pch.h"

VOID BroadcastVmxVirtualizationAllCores() {
  KeGenericCallDpc(DpcRoutinePerformVirtualization, NULL);
}

VOID BroadcastEnableDbAndBpExitingAllCores() {
  KeGenericCallDpc(DpcRoutineEnableDbAndBpExitingOnAllCores, NULL);
}

VOID BroadcastDisableDbAndBpExitingAllCores() {
  KeGenericCallDpc(DpcRoutineDisableDbAndBpExitingOnAllCores, NULL);
}

VOID BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores() {
  KeGenericCallDpc(DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores, NULL);
}

VOID BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores() {
  KeGenericCallDpc(DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores,
                   NULL);
}

VOID BroadcastEnableNmiExitingAllCores() {
  KeGenericCallDpc(DpcRoutineEnableNmiVmexitOnAllCores, NULL);
}

VOID BroadcastDisableNmiExitingAllCores() {
  KeGenericCallDpc(DpcRoutineDisableNmiVmexitOnAllCores, NULL);
}

VOID BroadcastNotifyAllToInvalidateEptAllCores() {
  KeGenericCallDpc(DpcRoutineInvalidateEptOnAllCores, (PVOID)TRUE);
}

VOID BroadcastEnableRdtscExitingAllCores() {
  KeGenericCallDpc(DpcRoutineEnableRdtscExitingAllCores, NULL);
}

VOID BroadcastDisableRdtscExitingAllCores() {
  KeGenericCallDpc(DpcRoutineDisableRdtscExitingAllCores, NULL);
}

VOID BroadcastChangeAllMsrBitmapReadAllCores(UINT64 BitmapMask) {
  KeGenericCallDpc(DpcRoutineChangeMsrBitmapReadOnAllCores, (PVOID)BitmapMask);
}

VOID BroadcastResetChangeAllMsrBitmapReadAllCores() {
  KeGenericCallDpc(DpcRoutineResetMsrBitmapReadOnAllCores, NULL);
}

VOID BroadcastChangeAllMsrBitmapWriteAllCores(UINT64 BitmapMask) {
  KeGenericCallDpc(DpcRoutineChangeMsrBitmapWriteOnAllCores, (PVOID)BitmapMask);
}

VOID BroadcastResetAllMsrBitmapWriteAllCores() {
  KeGenericCallDpc(DpcRoutineResetMsrBitmapWriteOnAllCores, NULL);
}

VOID BroadcastDisableRdtscExitingForClearingEventsAllCores() {
  KeGenericCallDpc(DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores,
                   NULL);
}

VOID BroadcastDisableMov2ControlRegsExitingForClearingEventsAllCores(
    PDEBUGGER_EVENT_OPTIONS BroadcastingOption) {
  KeGenericCallDpc(DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores,
                   BroadcastingOption);
}

VOID BroadcastDisableMov2DebugRegsExitingForClearingEventsAllCores() {
  KeGenericCallDpc(DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores,
                   NULL);
}

VOID BroadcastEnableRdpmcExitingAllCores() {
  KeGenericCallDpc(DpcRoutineEnableRdpmcExitingAllCores, NULL);
}

VOID BroadcastDisableRdpmcExitingAllCores() {
  KeGenericCallDpc(DpcRoutineDisableRdpmcExitingAllCores, NULL);
}

VOID BroadcastSetExceptionBitmapAllCores(UINT64 ExceptionIndex) {
  KeGenericCallDpc(DpcRoutineSetExceptionBitmapOnAllCores,
                   (PVOID)ExceptionIndex);
}

VOID BroadcastUnsetExceptionBitmapAllCores(UINT64 ExceptionIndex) {
  KeGenericCallDpc(DpcRoutineUnsetExceptionBitmapOnAllCores,
                   (PVOID)ExceptionIndex);
}

VOID BroadcastResetExceptionBitmapAllCores() {
  KeGenericCallDpc(
      DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores,
      NULL);
}

VOID BroadcastEnableMovControlRegisterExitingAllCores(
    PDEBUGGER_EVENT_OPTIONS BroadcastingOption) {
  KeGenericCallDpc(DpcRoutineEnableMovControlRegisterExitingAllCores,
                   BroadcastingOption);
}

VOID BroadcastDisableMovToControlRegistersExitingAllCores(
    PDEBUGGER_EVENT_OPTIONS BroadcastingOption) {
  KeGenericCallDpc(DpcRoutineDisableMovControlRegisterExitingAllCores,
                   BroadcastingOption);
}

VOID BroadcastEnableMovDebugRegistersExitingAllCores() {
  KeGenericCallDpc(DpcRoutineEnableMovDebigRegisterExitingAllCores, NULL);
}

VOID BroadcastDisableMovDebugRegistersExitingAllCores() {
  KeGenericCallDpc(DpcRoutineDisableMovDebigRegisterExitingAllCores, NULL);
}

VOID BroadcastSetExternalInterruptExitingAllCores() {
  KeGenericCallDpc(DpcRoutineSetEnableExternalInterruptExitingOnAllCores, NULL);
}

VOID BroadcastUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() {
  KeGenericCallDpc(
      DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores,
      NULL);
}

VOID BroadcastIoBitmapChangeAllCores(UINT64 Port) {
  KeGenericCallDpc(DpcRoutineChangeIoBitmapOnAllCores, (PVOID)Port);
}

VOID BroadcastIoBitmapResetAllCores() {
  KeGenericCallDpc(DpcRoutineResetIoBitmapOnAllCores, NULL);
}

VOID BroadcastEnableMovToCr3ExitingOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineEnableMovToCr3Exiting, 0x0);
}

VOID BroadcastChangeToMbecSupportedEptpOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineChangeToMbecSupportedEptp, 0x0);
}

VOID BroadcastRestoreToNormalEptpOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineRestoreToNormalEptp, 0x0);
}

VOID BroadcasDisableMbecOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineEnableOrDisableMbec, 0x0);
}

VOID BroadcasEnableMbecOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineEnableOrDisableMbec, (PVOID)0x1);
}

VOID BroadcastDisableMovToCr3ExitingOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineDisableMovToCr3Exiting, 0x0);
}

VOID BroadcastEnableEferSyscallEventsOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineEnableEferSyscallEvents, 0x0);
}

VOID BroadcastDisableEferSyscallEventsOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineDisableEferSyscallEvents, 0x0);
}

VOID BroadcastEnablePmlOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineEnablePml, 0x0);
}

VOID BroadcastDisablePmlOnAllProcessors() {
  KeGenericCallDpc(DpcRoutineDisablePml, 0x0);
}
