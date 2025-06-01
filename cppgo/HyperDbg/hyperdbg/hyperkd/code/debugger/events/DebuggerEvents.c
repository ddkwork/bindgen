
#include "pch.h"

VOID DebuggerEventEnableEferOnAllProcessors() {
  ConfigureEnableEferSyscallEventsOnAllProcessors();
}

VOID DebuggerEventDisableEferOnAllProcessors() {
  ConfigureDisableEferSyscallEventsOnAllProcessors();
}

VOID DebuggerEventEnableMovToCr3ExitingOnAllProcessors() {
  ConfigureEnableMovToCr3ExitingOnAllProcessors();
}

VOID DebuggerEventDisableMovToCr3ExitingOnAllProcessors() {
  ConfigureDisableMovToCr3ExitingOnAllProcessors();
}

BOOLEAN DebuggerEventEnableMonitorReadWriteExec(
    EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *HookingDetails,
    UINT32 ProcessId, BOOLEAN ApplyDirectlyFromVmxRoot) {
  if (!HookingDetails->SetHookForRead && !HookingDetails->SetHookForWrite &&
      !HookingDetails->SetHookForExec) {
    return FALSE;
  }
  if (HookingDetails->SetHookForWrite) {
    HookingDetails->SetHookForRead = TRUE;
  }
  if (ApplyDirectlyFromVmxRoot) {
    return ConfigureEptHookMonitorFromVmxRoot(
        KeGetCurrentProcessorNumberEx(NULL), HookingDetails);
  } else {
    return ConfigureEptHookMonitor(KeGetCurrentProcessorNumberEx(NULL),
                                   HookingDetails, ProcessId);
  }
}

BOOLEAN DebuggerCheckProcessOrThreadChange(_In_ UINT32 CoreId) {
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  if (DbgState->ThreadOrProcessTracingDetails
          .InterceptClockInterruptsForThreadChange ||
      DbgState->ThreadOrProcessTracingDetails
          .InterceptClockInterruptsForProcessChange) {
    if (DbgState->ThreadOrProcessTracingDetails
            .InterceptClockInterruptsForThreadChange) {
      return ThreadHandleThreadChange(DbgState);
    } else {
      return ProcessHandleProcessChange(DbgState);
    }
  }
  return FALSE;
}
