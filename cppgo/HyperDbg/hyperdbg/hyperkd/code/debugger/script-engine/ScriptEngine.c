
#include "pch.h"

UINT64 ScriptEngineWrapperGetInstructionPointer() {
  if (VmFuncVmxGetCurrentExecutionMode() == TRUE) {
    return VmFuncGetRip();
  } else {
    return (UINT64)NULL;
  }
}

UINT64 ScriptEngineWrapperGetAddressOfReservedBuffer(
    PDEBUGGER_EVENT_ACTION Action) {
  return Action->RequestedBuffer.RequstBufferAddress;
}

VOID ScriptEngineUpdateTargetCoreDateTime(PROCESSOR_DEBUGGING_STATE *DbgState) {
  LARGE_INTEGER SystemTime, LocalTime;
  KeQuerySystemTime(&SystemTime);
  ExSystemTimeToLocalTime(&SystemTime, &LocalTime);
  RtlTimeToTimeFields(&LocalTime, &DbgState->DateTimeHolder.TimeFields);
  sprintf_s(DbgState->DateTimeHolder.TimeBuffer,
            RTL_NUMBER_OF(DbgState->DateTimeHolder.TimeBuffer),
            "%02hd:%02hd:%02hd.%03hd", DbgState->DateTimeHolder.TimeFields.Hour,
            DbgState->DateTimeHolder.TimeFields.Minute,
            DbgState->DateTimeHolder.TimeFields.Second,
            DbgState->DateTimeHolder.TimeFields.Milliseconds);
  sprintf_s(DbgState->DateTimeHolder.DateBuffer,
            RTL_NUMBER_OF(DbgState->DateTimeHolder.DateBuffer),
            "%04hd-%02hd-%02hd", DbgState->DateTimeHolder.TimeFields.Year,
            DbgState->DateTimeHolder.TimeFields.Month,
            DbgState->DateTimeHolder.TimeFields.Day);
}

UINT64 ScriptEngineGetTargetCoreTime() {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CurrentCore];
  ScriptEngineUpdateTargetCoreDateTime(DbgState);
  return (UINT64)&DbgState->DateTimeHolder.TimeBuffer;
}

UINT64 ScriptEngineGetTargetCoreDate() {
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CurrentCore];
  ScriptEngineUpdateTargetCoreDateTime(DbgState);
  return (UINT64)&DbgState->DateTimeHolder.DateBuffer;
}
