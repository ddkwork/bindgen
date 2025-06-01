
#include "pch.h"

BOOLEAN BreakpointCheckAndPerformActionsOnTrapFlags(
    UINT32 ProcessId, UINT32 ThreadId, BOOLEAN *TrapSetByDebugger) {
  UINT32 Index;
  DEBUGGER_PROCESS_THREAD_INFORMATION ProcThrdInfo = {0};
  BOOLEAN Result;
  BOOLEAN ResultToReturn;
  RFLAGS Rflags = {0};
  Rflags.AsUInt = VmFuncGetRflags();
  ProcThrdInfo.Fields.ProcessId = ProcessId;
  ProcThrdInfo.Fields.ThreadId = ThreadId;
  SpinlockLock(&BreakpointCommandTrapListLock);
  Result = BinarySearchPerformSearchItem(
      (UINT64 *)&g_TrapFlagState.ThreadInformation[0],
      g_TrapFlagState.NumberOfItems, &Index, ProcThrdInfo.asUInt);
  *TrapSetByDebugger = Result;
  if (!Result && !Rflags.TrapFlag) {
    ResultToReturn = TRUE;
    goto Return;
  } else if (!Result && Rflags.TrapFlag) {
    ResultToReturn = FALSE;
    goto Return;
  } else {
    VmFuncSetRflagTrapFlag(FALSE);
    InsertionSortDeleteItem((UINT64 *)&g_TrapFlagState.ThreadInformation[0],
                            &g_TrapFlagState.NumberOfItems, Index);
    ResultToReturn = TRUE;
    goto Return;
  }
Return:
  SpinlockUnlock(&BreakpointCommandTrapListLock);
  return ResultToReturn;
}

BOOLEAN BreakpointTriggerCallbacks(PROCESSOR_DEBUGGING_STATE *DbgState,
                                   UINT32 ProcessId, UINT32 ThreadId) {
  UNREFERENCED_PARAMETER(DbgState);
  UNREFERENCED_PARAMETER(ProcessId);
  UNREFERENCED_PARAMETER(ThreadId);
  return FALSE;
}

BOOLEAN BreakpointRestoreTheTrapFlagOnceTriggered(UINT32 ProcessId,
                                                  UINT32 ThreadId) {
  UINT32 Index;
  BOOLEAN Result;
  BOOLEAN SuccessfullyStored;
  DEBUGGER_PROCESS_THREAD_INFORMATION ProcThrdInfo = {0};
  ProcThrdInfo.Fields.ProcessId = ProcessId;
  ProcThrdInfo.Fields.ThreadId = ThreadId;
  SpinlockLock(&BreakpointCommandTrapListLock);
  Result = BinarySearchPerformSearchItem(
      (UINT64 *)&g_TrapFlagState.ThreadInformation[0],
      g_TrapFlagState.NumberOfItems, &Index, ProcThrdInfo.asUInt);
  if (Result) {
    SuccessfullyStored = TRUE;
    goto Return;
  } else {
    SuccessfullyStored =
        InsertionSortInsertItem((UINT64 *)&g_TrapFlagState.ThreadInformation[0],
                                &g_TrapFlagState.NumberOfItems,
                                MAXIMUM_NUMBER_OF_THREAD_INFORMATION_FOR_TRAPS,
                                &Index, ProcThrdInfo.asUInt);
    goto Return;
  }
Return:
  SpinlockUnlock(&BreakpointCommandTrapListLock);
  return SuccessfullyStored;
}

BOOLEAN BreakpointCheckAndHandleDebugBreakpoint(UINT32 CoreId) {
  BOOLEAN TrapSetByDebugger;
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  BOOLEAN HandledByDebuggerRoutines = TRUE;
  if (BreakpointCheckAndPerformActionsOnTrapFlags(
          HANDLE_TO_UINT32(PsGetCurrentProcessId()),
          HANDLE_TO_UINT32(PsGetCurrentThreadId()), &TrapSetByDebugger)) {
    if (DbgState->ThreadOrProcessTracingDetails
            .DebugRegisterInterceptionState) {
      ThreadHandleThreadChange(DbgState);
    } else if (g_UserDebuggerState == TRUE &&
               (g_IsWaitingForUserModeProcessEntryToBeCalled ||
                g_IsWaitingForReturnAndRunFromPageFault)) {
      AttachingHandleEntrypointInterception(DbgState);
    } else if (g_KernelDebuggerState == TRUE) {
      KdHandleDebugEventsWhenKernelDebuggerIsAttached(DbgState,
                                                      TrapSetByDebugger);
    } else if (UdCheckAndHandleBreakpointsAndDebugBreaks(
                   DbgState,
                   DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK,
                   NULL)) {
    } else {
      HandledByDebuggerRoutines = FALSE;
      LogError(
          "Err, trap is supposed to be handled by the debugger, but none of "
          "routines handled it");
    }
  } else {
    if (g_InterceptDebugBreaks) {
      HandledByDebuggerRoutines = FALSE;
    } else if (g_KernelDebuggerState == TRUE) {
      KdHandleDebugEventsWhenKernelDebuggerIsAttached(DbgState,
                                                      TrapSetByDebugger);
    } else if (UdCheckAndHandleBreakpointsAndDebugBreaks(
                   DbgState,
                   DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK,
                   NULL)) {
    } else {
      HandledByDebuggerRoutines = FALSE;
    }
  }
  return HandledByDebuggerRoutines;
}

BOOLEAN BreakpointClear(PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor) {
  BYTE TargetMem = NULL_ZERO;
  if (!CheckAccessValidityAndSafety(BreakpointDescriptor->Address,
                                    sizeof(BYTE))) {
    MemoryMapperReadMemorySafeByPhysicalAddress(
        BreakpointDescriptor->PhysAddress, (UINT64)&TargetMem, sizeof(BYTE));
    if (TargetMem != 0xcc) {
      return FALSE;
    }
  }
  MemoryMapperWriteMemorySafeByPhysicalAddress(
      BreakpointDescriptor->PhysAddress,
      (UINT64)&BreakpointDescriptor->PreviousByte, sizeof(BYTE));
  BreakpointDescriptor->Enabled = FALSE;
  BreakpointDescriptor->AvoidReApplyBreakpoint = TRUE;
  return TRUE;
}

VOID BreakpointClearAndDeallocateMemory(
    PDEBUGGEE_BP_DESCRIPTOR BreakpointDesc) {
  BreakpointClear(BreakpointDesc);
  RemoveEntryList(&BreakpointDesc->BreakpointsList);
  PoolManagerFreePool((UINT64)BreakpointDesc);
}

BOOLEAN BreakpointCheckAndHandleReApplyingBreakpoint(UINT32 CoreId) {
  BOOLEAN Result = FALSE;
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  if (DbgState->SoftwareBreakpointState != NULL) {
    BYTE BreakpointByte = 0xcc;
    Result = TRUE;
    MemoryMapperWriteMemorySafeByPhysicalAddress(
        DbgState->SoftwareBreakpointState->PhysAddress, (UINT64)&BreakpointByte,
        sizeof(BYTE));
    if (DbgState->SoftwareBreakpointState->SetRflagsIFBitOnMtf) {
      RFLAGS Rflags = {0};
      Rflags.AsUInt = VmFuncGetRflags();
      Rflags.InterruptEnableFlag = TRUE;
      VmFuncSetRflags(Rflags.AsUInt);
      DbgState->SoftwareBreakpointState->SetRflagsIFBitOnMtf = FALSE;
    }
    DbgState->SoftwareBreakpointState = NULL;
  }
  return Result;
}

BOOLEAN BreakpointCheckAndHandleDebuggerDefinedBreakpoints(
    PROCESSOR_DEBUGGING_STATE *DbgState, UINT64 GuestRip,
    DEBUGGEE_PAUSING_REASON Reason, BOOLEAN ChangeMtfState) {
  CR3_TYPE GuestCr3 = {0};
  BOOLEAN IsHandledByBpRoutines = FALSE;
  PLIST_ENTRY TempList = 0;
  UINT64 GuestRipPhysical = (UINT64)NULL;
  DEBUGGER_TRIGGERED_EVENT_DETAILS TargetContext = {0};
  RFLAGS Rflags = {0};
  BOOLEAN AvoidUnsetMtf = FALSE;
  BOOLEAN IgnoreUserHandling = FALSE;
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  GuestRipPhysical =
      VirtualAddressToPhysicalAddressByProcessCr3((PVOID)GuestRip, GuestCr3);
  TempList = &g_BreakpointsListHead;
  while (&g_BreakpointsListHead != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc =
        CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
    if (CurrentBreakpointDesc->PhysAddress == GuestRipPhysical) {
      IsHandledByBpRoutines = TRUE;
      MemoryMapperWriteMemorySafeByPhysicalAddress(
          GuestRipPhysical, (UINT64)&CurrentBreakpointDesc->PreviousByte,
          sizeof(BYTE));
      TargetContext.Context = (PVOID)VmFuncGetLastVmexitRip(DbgState->CoreId);
      if (Reason == DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT) {
        TargetContext.Tag = CurrentBreakpointDesc->BreakpointId;
      }
      DbgState->InstructionLengthHint =
          CurrentBreakpointDesc->InstructionLength;
      if ((CurrentBreakpointDesc->Pid == DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES ||
           CurrentBreakpointDesc->Pid ==
               HANDLE_TO_UINT32(PsGetCurrentProcessId())) &&
          (CurrentBreakpointDesc->Tid == DEBUGGEE_BP_APPLY_TO_ALL_THREADS ||
           CurrentBreakpointDesc->Tid ==
               HANDLE_TO_UINT32(PsGetCurrentThreadId())) &&
          (CurrentBreakpointDesc->Core == DEBUGGEE_BP_APPLY_TO_ALL_CORES ||
           CurrentBreakpointDesc->Core == DbgState->CoreId)) {
        if (CurrentBreakpointDesc->RemoveAfterHit) {
          BreakpointClearAndDeallocateMemory(CurrentBreakpointDesc);
        }
        if (CurrentBreakpointDesc->CheckForCallbacks) {
          IgnoreUserHandling = BreakpointTriggerCallbacks(
              DbgState, HANDLE_TO_UINT32(PsGetCurrentProcessId()),
              HANDLE_TO_UINT32(PsGetCurrentThreadId()));
        }
        if (!IgnoreUserHandling && !g_InterceptBreakpoints &&
            !g_InterceptBreakpointsAndEventsForCommandsInRemoteComputer) {
          KdHandleBreakpointAndDebugBreakpoints(DbgState, Reason,
                                                &TargetContext);
        }
      }
      DbgState->InstructionLengthHint = 0;
      if (!CurrentBreakpointDesc->AvoidReApplyBreakpoint) {
        DbgState->SoftwareBreakpointState = CurrentBreakpointDesc;
        VmFuncSetMonitorTrapFlag(TRUE);
        AvoidUnsetMtf = TRUE;
        Rflags.AsUInt = VmFuncGetRflags();
        if (Rflags.InterruptEnableFlag) {
          Rflags.InterruptEnableFlag = FALSE;
          VmFuncSetRflags(Rflags.AsUInt);
          DbgState->SoftwareBreakpointState->SetRflagsIFBitOnMtf = TRUE;
        }
      }
      VmFuncSuppressRipIncrement(DbgState->CoreId);
      break;
    }
  }
  if (IsHandledByBpRoutines && ChangeMtfState) {
    VmFuncChangeMtfUnsettingState(DbgState->CoreId, AvoidUnsetMtf);
  }
  return IsHandledByBpRoutines;
}

BOOLEAN BreakpointHandleBreakpoints(UINT32 CoreId) {
  DEBUGGER_TRIGGERED_EVENT_DETAILS TargetContext = {0};
  UINT64 GuestRip = 0;
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  if (g_KernelDebuggerState) {
    GuestRip = VmFuncGetRip();
    if (!BreakpointCheckAndHandleDebuggerDefinedBreakpoints(
            DbgState, GuestRip,
            DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT, FALSE)) {
      if (g_InterceptBreakpoints ||
          g_InterceptBreakpointsAndEventsForCommandsInRemoteComputer) {
        return FALSE;
      }
      TargetContext.Context = (PVOID)GuestRip;
      KdHandleBreakpointAndDebugBreakpoints(
          DbgState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT,
          &TargetContext);
      VmFuncPerformRipIncrement(DbgState->CoreId);
    }
  } else {
    return FALSE;
  }
  return TRUE;
}

BOOLEAN BreakpointWrite(PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor) {
  BYTE PreviousByte = NULL_ZERO;
  BYTE BreakpointByte = 0xcc;
  if (!CheckAccessValidityAndSafety(BreakpointDescriptor->Address,
                                    sizeof(BYTE))) {
    return FALSE;
  }
  MemoryMapperReadMemorySafeOnTargetProcess(BreakpointDescriptor->Address,
                                            &PreviousByte, sizeof(BYTE));
  BreakpointDescriptor->PreviousByte = PreviousByte;
  BreakpointDescriptor->Enabled = TRUE;
  BreakpointDescriptor->AvoidReApplyBreakpoint = FALSE;
  MemoryMapperWriteMemorySafeByPhysicalAddress(
      BreakpointDescriptor->PhysAddress, (UINT64)&BreakpointByte, sizeof(BYTE));
  return TRUE;
}

VOID BreakpointRemoveAllBreakpoints() {
  PLIST_ENTRY TempList = 0;
  TempList = &g_BreakpointsListHead;
  while (&g_BreakpointsListHead != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc =
        CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
    BreakpointClearAndDeallocateMemory(CurrentBreakpointDesc);
  }
}

PDEBUGGEE_BP_DESCRIPTOR BreakpointGetEntryByBreakpointId(UINT64 BreakpointId) {
  PLIST_ENTRY TempList = 0;
  TempList = &g_BreakpointsListHead;
  while (&g_BreakpointsListHead != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc =
        CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
    if (CurrentBreakpointDesc->BreakpointId == BreakpointId) {
      return CurrentBreakpointDesc;
    }
  }
  return NULL;
}

PDEBUGGEE_BP_DESCRIPTOR BreakpointGetEntryByAddress(UINT64 Address) {
  PLIST_ENTRY TempList = 0;
  TempList = &g_BreakpointsListHead;
  while (&g_BreakpointsListHead != TempList->Flink) {
    TempList = TempList->Flink;
    PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc =
        CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
    if (CurrentBreakpointDesc->Address == Address) {
      return CurrentBreakpointDesc;
    }
  }
  return NULL;
}

BOOLEAN BreakpointAddNew(PDEBUGGEE_BP_PACKET BpDescriptorArg) {
  PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor = NULL;
  CR3_TYPE GuestCr3 = {0};
  BOOLEAN IsAddress32Bit = FALSE;
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  if (!CheckAccessValidityAndSafety(BpDescriptorArg->Address, sizeof(BYTE))) {
    BpDescriptorArg->Result =
        DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS;
    return FALSE;
  }
  if (BpDescriptorArg->Core != DEBUGGEE_BP_APPLY_TO_ALL_CORES &&
      !CommonValidateCoreNumber(BpDescriptorArg->Core)) {
    BpDescriptorArg->Result = DEBUGGER_ERROR_INVALID_CORE_ID;
    return FALSE;
  }
  if (BreakpointGetEntryByAddress(BpDescriptorArg->Address) != NULL) {
    BpDescriptorArg->Result =
        DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS;
    return FALSE;
  }
  BreakpointDescriptor = (DEBUGGEE_BP_DESCRIPTOR *)PoolManagerRequestPool(
      BREAKPOINT_DEFINITION_STRUCTURE, TRUE, sizeof(DEBUGGEE_BP_DESCRIPTOR));
  if (BreakpointDescriptor == NULL) {
    BpDescriptorArg->Result =
        DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE;
    return FALSE;
  }
  g_MaximumBreakpointId++;
  BreakpointDescriptor->BreakpointId = g_MaximumBreakpointId;
  BreakpointDescriptor->Address = BpDescriptorArg->Address;
  BreakpointDescriptor->PhysAddress =
      VirtualAddressToPhysicalAddressByProcessCr3(
          (PVOID)BpDescriptorArg->Address, GuestCr3);
  BreakpointDescriptor->Core = BpDescriptorArg->Core;
  BreakpointDescriptor->Pid = BpDescriptorArg->Pid;
  BreakpointDescriptor->Tid = BpDescriptorArg->Tid;
  BreakpointDescriptor->RemoveAfterHit = BpDescriptorArg->RemoveAfterHit;
  BreakpointDescriptor->CheckForCallbacks = BpDescriptorArg->CheckForCallbacks;
  if (BpDescriptorArg->Address & 0xff00000000000000) {
    IsAddress32Bit = FALSE;
  } else {
    IsAddress32Bit = KdIsGuestOnUsermode32Bit();
  }
  BreakpointDescriptor->InstructionLength =
      (UINT16)DisassemblerLengthDisassembleEngineInVmxRootOnTargetProcess(
          (PVOID)BpDescriptorArg->Address, IsAddress32Bit);
  BreakpointDescriptor->Enabled = TRUE;
  InsertHeadList(&g_BreakpointsListHead,
                 &(BreakpointDescriptor->BreakpointsList));
  BreakpointWrite(BreakpointDescriptor);
  BpDescriptorArg->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return TRUE;
}

VOID BreakpointListAllBreakpoint() {
  BOOLEAN IsListEmpty = TRUE;
  PLIST_ENTRY TempList = 0;
  TempList = &g_BreakpointsListHead;
  while (&g_BreakpointsListHead != TempList->Blink) {
    TempList = TempList->Blink;
    PDEBUGGEE_BP_DESCRIPTOR CurrentBreakpointDesc =
        CONTAINING_RECORD(TempList, DEBUGGEE_BP_DESCRIPTOR, BreakpointsList);
    if (IsListEmpty) {
      Log("Id   Address           Status\n");
      Log("--   ---------------   --------");
      IsListEmpty = FALSE;
    }
    Log("\n%02x   %016llx  %s", CurrentBreakpointDesc->BreakpointId,
        CurrentBreakpointDesc->Address,
        CurrentBreakpointDesc->Enabled ? "enabled" : "disabled");
    if (CurrentBreakpointDesc->Core != DEBUGGEE_BP_APPLY_TO_ALL_CORES) {
      Log(" core = %x ", CurrentBreakpointDesc->Core);
    }
    if (CurrentBreakpointDesc->Pid != DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES) {
      Log(" pid = %x ", CurrentBreakpointDesc->Pid);
    }
    if (CurrentBreakpointDesc->Tid != DEBUGGEE_BP_APPLY_TO_ALL_THREADS) {
      Log(" tid = %x ", CurrentBreakpointDesc->Tid);
    }
  }
  if (IsListEmpty) {
    Log("Breakpoints list is empty");
  }
}

BOOLEAN BreakpointListOrModify(
    PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET ListOrModifyBreakpoints) {
  PDEBUGGEE_BP_DESCRIPTOR BreakpointDescriptor = NULL;
  if (ListOrModifyBreakpoints->Request ==
      DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS) {
    BreakpointListAllBreakpoint();
  } else if (ListOrModifyBreakpoints->Request ==
             DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE) {
    BreakpointDescriptor =
        BreakpointGetEntryByBreakpointId(ListOrModifyBreakpoints->BreakpointId);
    if (BreakpointDescriptor == NULL) {
      ListOrModifyBreakpoints->Result = DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND;
      return FALSE;
    }
    if (BreakpointDescriptor->Enabled) {
      ListOrModifyBreakpoints->Result =
          DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED;
      return FALSE;
    }
    BreakpointWrite(BreakpointDescriptor);
  } else if (ListOrModifyBreakpoints->Request ==
             DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE) {
    BreakpointDescriptor =
        BreakpointGetEntryByBreakpointId(ListOrModifyBreakpoints->BreakpointId);
    if (BreakpointDescriptor == NULL) {
      ListOrModifyBreakpoints->Result = DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND;
      return FALSE;
    }
    if (!BreakpointDescriptor->Enabled) {
      ListOrModifyBreakpoints->Result =
          DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED;
      return FALSE;
    }
    BreakpointClear(BreakpointDescriptor);
  } else if (ListOrModifyBreakpoints->Request ==
             DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR) {
    BreakpointDescriptor =
        BreakpointGetEntryByBreakpointId(ListOrModifyBreakpoints->BreakpointId);
    if (BreakpointDescriptor == NULL) {
      ListOrModifyBreakpoints->Result = DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND;
      return FALSE;
    }
    BreakpointClearAndDeallocateMemory(BreakpointDescriptor);
  }
  ListOrModifyBreakpoints->Result = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
  return TRUE;
}
