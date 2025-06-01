
#include "pch.h"

VOID TracingPerformInstrumentationStepIn(PROCESSOR_DEBUGGING_STATE *DbgState) {
  DisassemblerShowOneInstructionInVmxRootMode(
      (PVOID)VmFuncGetLastVmexitRip(DbgState->CoreId), FALSE);
  DbgState->TracingMode = TRUE;
  VmFuncRegisterMtfBreak(DbgState->CoreId);
  VmFuncEnableMtfAndChangeExternalInterruptState(DbgState->CoreId);
}

VOID TracingHandleMtf(PROCESSOR_DEBUGGING_STATE *DbgState) {
  TracingRestoreSystemState(DbgState);
  MetaDispatchEventInstrumentationTrace(DbgState);
  TracingCheckForContinuingSteps(DbgState);
}

VOID TracingRestoreSystemState(PROCESSOR_DEBUGGING_STATE *DbgState) {
  DbgState->TracingMode = FALSE;
  VmFuncUnRegisterMtfBreak(DbgState->CoreId);
  VmFuncEnableAndCheckForPreviousExternalInterrupts(DbgState->CoreId);
}

VOID TracingCheckForContinuingSteps(PROCESSOR_DEBUGGING_STATE *DbgState) {
  if (DbgState->TracingMode) {
    VmFuncChangeMtfUnsettingState(DbgState->CoreId, TRUE);
  }
}

VOID TracingPerformRegularStepInInstruction(
    PROCESSOR_DEBUGGING_STATE *DbgState) {
  UNREFERENCED_PARAMETER(DbgState);
  UINT64 Interruptibility;
  UINT64 InterruptibilityOld = NULL64_ZERO;
  VmFuncSetRflagTrapFlag(TRUE);
  InterruptibilityOld = VmFuncGetInterruptibilityState();
  Interruptibility = InterruptibilityOld;
  Interruptibility = VmFuncClearSteppingBits(Interruptibility);
  if ((Interruptibility != InterruptibilityOld)) {
    VmFuncSetInterruptibilityState(Interruptibility);
  }
}
