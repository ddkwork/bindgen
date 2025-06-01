
#include "pch.h"

BOOLEAN TransparentHideDebugger(
    PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE TransparentModeRequest) {
  if (!g_TransparentMode) {
    g_TransparentModeTrapFlagState = (TRANSPARENT_MODE_TRAP_FLAG_STATE *)
        PlatformMemAllocateZeroedNonPagedPool(
            sizeof(TRANSPARENT_MODE_TRAP_FLAG_STATE));
    BroadcastEnableDbAndBpExitingAllCores();
    g_TransparentMode = TRUE;
    TransparentModeRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
    return TRUE;
  } else {
    TransparentModeRequest->KernelStatus = DEBUGGER_ERROR_DEBUGGER_ALREADY_HIDE;
    return FALSE;
  }
}

BOOLEAN TransparentUnhideDebugger(
    PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE TransparentModeRequest) {
  if (g_TransparentMode) {
    g_TransparentMode = FALSE;
    BroadcastDisableDbAndBpExitingAllCores();
    PlatformMemFreePool(g_TransparentModeTrapFlagState);
    TransparentModeRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
    return TRUE;
  } else {
    TransparentModeRequest->KernelStatus =
        DEBUGGER_ERROR_DEBUGGER_ALREADY_UNHIDE;
    return FALSE;
  }
}

VOID TransparentCPUID(INT32 CpuInfo[], PGUEST_REGS Regs) {
  if (Regs->rax == CPUID_PROCESSOR_AND_PROCESSOR_FEATURE_IDENTIFIERS) {
    CpuInfo[2] &= ~HYPERV_HYPERVISOR_PRESENT_BIT;
  } else if (Regs->rax == CPUID_HV_VENDOR_AND_MAX_FUNCTIONS ||
             Regs->rax == HYPERV_CPUID_INTERFACE) {
    CpuInfo[0] = CpuInfo[1] = CpuInfo[2] = CpuInfo[3] = 0;
  }
}

BOOLEAN TransparentStoreProcessInformation(
    UINT32 ProcessId, UINT32 ThreadId, UINT64 Context,
    TRANSPARENT_MODE_CONTEXT_PARAMS *Params) {
  UINT32 Index;
  BOOLEAN Result;
  BOOLEAN SuccessfullyStored;
  TRANSPARENT_MODE_PROCESS_THREAD_INFORMATION ProcThrdInfo = {0};
  ProcThrdInfo.Fields.ProcessId = ProcessId;
  ProcThrdInfo.Fields.ThreadId = ThreadId;
  SpinlockLock(&TransparentModeTrapListLock);
  Result = BinarySearchPerformSearchItem(
      (UINT64 *)&g_TransparentModeTrapFlagState->ThreadInformation[0],
      g_TransparentModeTrapFlagState->NumberOfItems, &Index,
      ProcThrdInfo.asUInt);
  if (Result) {
    SuccessfullyStored = TRUE;
    goto Return;
  } else {
    SuccessfullyStored = InsertionSortInsertItem(
        (UINT64 *)&g_TransparentModeTrapFlagState->ThreadInformation[0],
        &g_TransparentModeTrapFlagState->NumberOfItems,
        MAXIMUM_NUMBER_OF_THREAD_INFORMATION_FOR_TRANSPARENT_MODE_TRAPS, &Index,
        ProcThrdInfo.asUInt);
    if (SuccessfullyStored) {
      g_TransparentModeTrapFlagState->Context[Index] = Context;
      memcpy(&g_TransparentModeTrapFlagState->Params[Index], Params,
             sizeof(TRANSPARENT_MODE_CONTEXT_PARAMS));
    }
    goto Return;
  }
Return:
  SpinlockUnlock(&TransparentModeTrapListLock);
  return SuccessfullyStored;
}

BOOLEAN TransparentSetTrapFlagAfterSyscall(
    VIRTUAL_MACHINE_STATE *VCpu, UINT32 ProcessId, UINT32 ThreadId,
    UINT64 Context, TRANSPARENT_MODE_CONTEXT_PARAMS *Params) {
  if (!g_TransparentMode) {
    return FALSE;
  }
  if (!TransparentStoreProcessInformation(ProcessId, ThreadId, Context,
                                          Params)) {
    return FALSE;
  }
  VCpu->Regs->r11 |= X86_FLAGS_TF;
  return TRUE;
}

BOOLEAN TransparentCheckAndHandleAfterSyscallTrapFlags(
    VIRTUAL_MACHINE_STATE *VCpu, UINT32 ProcessId, UINT32 ThreadId) {
  RFLAGS Rflags = {0};
  UINT32 Index;
  UINT64 Context = NULL64_ZERO;
  TRANSPARENT_MODE_CONTEXT_PARAMS Params;
  TRANSPARENT_MODE_PROCESS_THREAD_INFORMATION ProcThrdInfo = {0};
  BOOLEAN Result;
  BOOLEAN ResultToReturn;
  Rflags.AsUInt = HvGetRflags();
  if (!Rflags.TrapFlag) {
    return FALSE;
  }
  ProcThrdInfo.Fields.ProcessId = ProcessId;
  ProcThrdInfo.Fields.ThreadId = ThreadId;
  SpinlockLock(&TransparentModeTrapListLock);
  Result = BinarySearchPerformSearchItem(
      (UINT64 *)&g_TransparentModeTrapFlagState->ThreadInformation[0],
      g_TransparentModeTrapFlagState->NumberOfItems, &Index,
      ProcThrdInfo.asUInt);
  if (Result) {
    Context = g_TransparentModeTrapFlagState->Context[Index];
    memcpy(&Params, &g_TransparentModeTrapFlagState->Params[Index],
           sizeof(TRANSPARENT_MODE_CONTEXT_PARAMS));
    HvSetRflagTrapFlag(FALSE);
    InsertionSortDeleteItem(
        (UINT64 *)&g_TransparentModeTrapFlagState->ThreadInformation[0],
        &g_TransparentModeTrapFlagState->NumberOfItems, Index);
    ResultToReturn = TRUE;
    goto ReturnResult;
  } else {
    ResultToReturn = FALSE;
    goto ReturnResult;
  }
ReturnResult:
  SpinlockUnlock(&TransparentModeTrapListLock);
  if (ResultToReturn) {
    TransparentCallbackHandleAfterSyscall(VCpu, ProcessId, ThreadId, Context,
                                          &Params);
  }
  return ResultToReturn;
}

VOID TransparentCallbackHandleAfterSyscall(
    VIRTUAL_MACHINE_STATE *VCpu, UINT32 ProcessId, UINT32 ThreadId,
    UINT64 Context, TRANSPARENT_MODE_CONTEXT_PARAMS *Params) {
  LogInfo(
      "Transparent callback handle the trap flag for process: %x, thread: %x, "
      "rip: %llx, context: %llx (p1: %llx, p2: %llx, p3: %llx, p4: %llx)\n",
      ProcessId, ThreadId, VCpu->LastVmexitRip, Context, Params->OptionalParam1,
      Params->OptionalParam2, Params->OptionalParam3, Params->OptionalParam4);
}
