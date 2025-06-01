
#pragma once
VOID TracingHandleMtf(PROCESSOR_DEBUGGING_STATE *DbgState);
VOID TracingRestoreSystemState(PROCESSOR_DEBUGGING_STATE *DbgState);
VOID TracingCheckForContinuingSteps(PROCESSOR_DEBUGGING_STATE *DbgState);
VOID TracingPerformInstrumentationStepIn(PROCESSOR_DEBUGGING_STATE *DbgState);
VOID TracingPerformRegularStepInInstruction(
    PROCESSOR_DEBUGGING_STATE *DbgState);
