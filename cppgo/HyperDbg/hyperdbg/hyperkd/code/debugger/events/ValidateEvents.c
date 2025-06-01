
#include "pch.h"

BOOLEAN ValidateEventMonitor(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
                             PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                             BOOLEAN InputFromVmxRoot) {
  UINT32 TempPid;
  TempPid = EventDetails->ProcessId;
  if (TempPid == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES) {
    TempPid = HANDLE_TO_UINT32(PsGetCurrentProcessId());
  }
  if (InputFromVmxRoot &&
      TempPid != HANDLE_TO_UINT32(PsGetCurrentProcessId())) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error =
        DEBUGGER_ERROR_PROCESS_ID_CANNOT_BE_SPECIFIED_WHILE_APPLYING_EVENT_FROM_VMX_ROOT_MODE;
    return FALSE;
  }
  if ((DEBUGGER_HOOK_MEMORY_TYPE)EventDetails->Options.OptionalParam3 ==
      DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS) {
    if (CheckAddressPhysical(EventDetails->Options.OptionalParam1) == FALSE ||
        CheckAddressPhysical(EventDetails->Options.OptionalParam2) == FALSE) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_PHYSICAL_ADDRESS;
      return FALSE;
    }
  } else {
    if (InputFromVmxRoot) {
      if (VirtualAddressToPhysicalAddressOnTargetProcess(
              (PVOID)EventDetails->Options.OptionalParam1) == (UINT64)NULL ||
          VirtualAddressToPhysicalAddressOnTargetProcess(
              (PVOID)EventDetails->Options.OptionalParam2) == (UINT64)NULL) {
        ResultsToReturn->IsSuccessful = FALSE;
        ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_ADDRESS;
        return FALSE;
      }
    } else {
      if (VirtualAddressToPhysicalAddressByProcessId(
              (PVOID)EventDetails->Options.OptionalParam1, TempPid) ==
              (UINT64)NULL ||
          VirtualAddressToPhysicalAddressByProcessId(
              (PVOID)EventDetails->Options.OptionalParam2, TempPid) ==
              (UINT64)NULL) {
        ResultsToReturn->IsSuccessful = FALSE;
        ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_ADDRESS;
        return FALSE;
      }
    }
  }
  if (EventDetails->Options.OptionalParam1 >=
      EventDetails->Options.OptionalParam2) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_ADDRESS;
    return FALSE;
  }
  return TRUE;
}

BOOLEAN ValidateEventException(
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  if (EventDetails->Options.OptionalParam1 !=
          DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES &&
      EventDetails->Options.OptionalParam1 >= 31) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error =
        DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES;
    return FALSE;
  }
  return TRUE;
}

BOOLEAN ValidateEventInterrupt(
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  if (!(EventDetails->Options.OptionalParam1 >= 32 &&
        EventDetails->Options.OptionalParam1 <= 0xff)) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error = DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID;
    return FALSE;
  }
  return TRUE;
}

BOOLEAN ValidateEventTrapExec(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
                              PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                              BOOLEAN InputFromVmxRoot) {
  UNREFERENCED_PARAMETER(InputFromVmxRoot);
  if (EventDetails->Options.OptionalParam1 !=
          DEBUGGER_EVENT_MODE_TYPE_USER_MODE &&
      EventDetails->Options.OptionalParam1 !=
          DEBUGGER_EVENT_MODE_TYPE_KERNEL_MODE &&
      EventDetails->Options.OptionalParam1 !=
          DEBUGGER_EVENT_MODE_TYPE_USER_MODE_AND_KERNEL_MODE) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error = DEBUGGER_ERROR_MODE_EXECUTION_IS_INVALID;
    return FALSE;
  }
  return TRUE;
}

BOOLEAN ValidateEventEptHookHiddenBreakpointAndInlineHooks(
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot) {
  UINT32 TempPid;
  TempPid = EventDetails->ProcessId;
  if (TempPid == DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES) {
    TempPid = HANDLE_TO_UINT32(PsGetCurrentProcessId());
  }
  if (InputFromVmxRoot &&
      TempPid != HANDLE_TO_UINT32(PsGetCurrentProcessId())) {
    ResultsToReturn->IsSuccessful = FALSE;
    ResultsToReturn->Error =
        DEBUGGER_ERROR_PROCESS_ID_CANNOT_BE_SPECIFIED_WHILE_APPLYING_EVENT_FROM_VMX_ROOT_MODE;
    return FALSE;
  }
  if (InputFromVmxRoot) {
    if (VirtualAddressToPhysicalAddressOnTargetProcess(
            (PVOID)EventDetails->Options.OptionalParam1) == (UINT64)NULL) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_ADDRESS;
      return FALSE;
    }
  } else {
    if (VirtualAddressToPhysicalAddressByProcessId(
            (PVOID)EventDetails->Options.OptionalParam1, TempPid) ==
        (UINT64)NULL) {
      ResultsToReturn->IsSuccessful = FALSE;
      ResultsToReturn->Error = DEBUGGER_ERROR_INVALID_ADDRESS;
      return FALSE;
    }
  }
  return TRUE;
}
