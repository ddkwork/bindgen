
#include "pch.h"
#ifdef SCRIPT_ENGINE_USER_MODE
extern UINT64 g_CurrentExprEvalResult;
extern BOOLEAN g_CurrentExprEvalResultHasError;
#endif
UINT64 GetValue(PGUEST_REGS GuestRegs, PACTION_BUFFER ActionBuffer,
                PSCRIPT_ENGINE_GENERAL_REGISTERS ScriptGeneralRegisters,
                PSYMBOL Symbol, BOOLEAN ReturnReference);

BOOLEAN ScriptEngineFunctionEq(UINT64 Address, QWORD Value, BOOL *HasError) {
  UNREFERENCED_PARAMETER(HasError);
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety(Address, sizeof(QWORD))) {
    return FALSE;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  *(UINT64 *)Address = Value;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(QWORD));
#endif
  return TRUE;
}

BOOLEAN ScriptEngineFunctionEd(UINT64 Address, DWORD Value, BOOL *HasError) {
  UNREFERENCED_PARAMETER(HasError);
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety(Address, sizeof(DWORD))) {
    return FALSE;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  *(DWORD *)Address = Value;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(DWORD));
#endif
  return TRUE;
}

BOOLEAN ScriptEngineFunctionEb(UINT64 Address, BYTE Value, BOOL *HasError) {
  UNREFERENCED_PARAMETER(HasError);
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety(Address, sizeof(BYTE))) {
    return FALSE;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  *(BYTE *)Address = Value;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperWriteMemorySafeOnTargetProcess(Address, &Value, sizeof(BYTE));
#endif
  return TRUE;
}

BOOLEAN ScriptEngineFunctionEqPa(UINT64 Address, QWORD Value, BOOL *HasError) {
  UNREFERENCED_PARAMETER(HasError);
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical(Address)) {
    return FALSE;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address functions (eq_pa) is not possible in "
      "user-mode\n");
  return FALSE;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperWriteMemorySafeByPhysicalAddress(Address, (UINT64)&Value,
                                               sizeof(QWORD));
#endif
  return TRUE;
}

BOOLEAN ScriptEngineFunctionEdPa(UINT64 Address, DWORD Value, BOOL *HasError) {
  UNREFERENCED_PARAMETER(HasError);
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical(Address)) {
    return FALSE;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address functions (ed_pa) is not possible in "
      "user-mode\n");
  return FALSE;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperWriteMemorySafeByPhysicalAddress(Address, (UINT64)&Value,
                                               sizeof(DWORD));
#endif
  return TRUE;
}

BOOLEAN ScriptEngineFunctionEbPa(UINT64 Address, BYTE Value, BOOL *HasError) {
  UNREFERENCED_PARAMETER(HasError);
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical(Address)) {
    return FALSE;
  }
#endif
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address functions (eb_pa) is not possible in "
      "user-mode\n");
  return FALSE;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  MemoryMapperWriteMemorySafeByPhysicalAddress(Address, (UINT64)&Value,
                                               sizeof(BYTE));
#endif
  return TRUE;
}

BOOLEAN ScriptEngineFunctionCheckAddress(UINT64 Address, UINT32 Length) {
#ifdef SCRIPT_ENGINE_USER_MODE
  if (CheckAccessValidityAndSafety(Address, Length)) {
    return TRUE;
  }
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (CheckAccessValidityAndSafety(Address, Length)) {
    return TRUE;
  }
#endif
  return FALSE;
}

VOID ScriptEngineFunctionMemcpy(UINT64 Destination, UINT64 Source, UINT32 Num,
                                BOOL *HasError) {
  UINT64 PrevReadLen = 0;
  BYTE MovingBuffer[DebuggerScriptEngineMemcpyMovingBufferSize] = {0};
#ifdef SCRIPT_ENGINE_USER_MODE
  if (!CheckAccessValidityAndSafety(Destination, Num)) {
    *HasError = TRUE;
    return;
  }
  if (!CheckAccessValidityAndSafety(Source, Num)) {
    *HasError = TRUE;
    return;
  }
  memcpy((void *)Destination, (void *)Source, Num);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety(Destination, Num)) {
    *HasError = TRUE;
    return;
  }
  if (!CheckAccessValidityAndSafety(Source, Num)) {
    *HasError = TRUE;
    return;
  }
  while (Num > 0) {
    if (Num > DebuggerScriptEngineMemcpyMovingBufferSize) {
      MemoryMapperReadMemorySafeOnTargetProcess(
          Source + PrevReadLen, MovingBuffer,
          DebuggerScriptEngineMemcpyMovingBufferSize);
      MemoryMapperWriteMemorySafeOnTargetProcess(
          Destination + PrevReadLen, MovingBuffer,
          DebuggerScriptEngineMemcpyMovingBufferSize);
      PrevReadLen += DebuggerScriptEngineMemcpyMovingBufferSize;
      Num -= DebuggerScriptEngineMemcpyMovingBufferSize;
    } else {
      MemoryMapperReadMemorySafeOnTargetProcess(Source + PrevReadLen,
                                                MovingBuffer, Num);
      MemoryMapperWriteMemorySafeOnTargetProcess(Destination + PrevReadLen,
                                                 MovingBuffer, Num);
      PrevReadLen += Num;
      Num = 0;
    }
  }
#endif
}

VOID ScriptEngineFunctionMemcpyPa(UINT64 Destination, UINT64 Source, UINT32 Num,
                                  BOOL *HasError) {
  UINT64 PrevReadLen = 0;
  BYTE MovingBuffer[DebuggerScriptEngineMemcpyMovingBufferSize] = {0};
#ifdef SCRIPT_ENGINE_USER_MODE
  ShowMessages(
      "err, using physical address functions (memcpy_pa) is not possible in "
      "user-mode\n");
  return;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAddressPhysical(Destination)) {
    *HasError = TRUE;
    return;
  }
  if (!CheckAddressPhysical(Source)) {
    *HasError = TRUE;
    return;
  }
  while (Num > 0) {
    if (Num > DebuggerScriptEngineMemcpyMovingBufferSize) {
      MemoryMapperReadMemorySafeByPhysicalAddress(
          Source + PrevReadLen, (UINT64)MovingBuffer,
          DebuggerScriptEngineMemcpyMovingBufferSize);
      MemoryMapperWriteMemorySafeByPhysicalAddress(
          Destination + PrevReadLen, (UINT64)MovingBuffer,
          DebuggerScriptEngineMemcpyMovingBufferSize);
      PrevReadLen += DebuggerScriptEngineMemcpyMovingBufferSize;
      Num -= DebuggerScriptEngineMemcpyMovingBufferSize;
    } else {
      MemoryMapperReadMemorySafeByPhysicalAddress(Source + PrevReadLen,
                                                  (UINT64)MovingBuffer, Num);
      MemoryMapperWriteMemorySafeByPhysicalAddress(Destination + PrevReadLen,
                                                   (UINT64)MovingBuffer, Num);
      PrevReadLen += Num;
      Num = 0;
    }
  }
#endif
}

UINT64 ScriptEngineFunctionVirtualToPhysical(UINT64 Address) {
#ifdef SCRIPT_ENGINE_USER_MODE
  return NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return VirtualAddressToPhysicalAddressOnTargetProcess((PVOID)Address);
#endif
}

UINT64 ScriptEngineFunctionPhysicalToVirtual(UINT64 Address) {
#ifdef SCRIPT_ENGINE_USER_MODE
  return NULL;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  return PhysicalAddressToVirtualAddressOnTargetProcess((PVOID)Address);
#endif
}

VOID ScriptEngineFunctionPrint(UINT64 Tag, BOOLEAN ImmediateMessagePassing,
                               UINT64 Value) {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages("%llx\n", Value);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  char TempBuffer[20] = {0};
  UINT32 TempBufferLen = sprintf(TempBuffer, "%llx", Value);
  LogSimpleWithTag((UINT32)Tag, ImmediateMessagePassing, TempBuffer,
                   TempBufferLen + 1);
#endif
}

VOID ScriptEngineFunctionTestStatement(UINT64 Tag,
                                       BOOLEAN ImmediateMessagePassing,
                                       UINT64 Value) {
  UNREFERENCED_PARAMETER(Tag);
  UNREFERENCED_PARAMETER(ImmediateMessagePassing);
  UNREFERENCED_PARAMETER(Value);
#ifdef SCRIPT_ENGINE_USER_MODE
  g_CurrentExprEvalResult = Value;
  g_CurrentExprEvalResultHasError = FALSE;
#endif
}

VOID ScriptEngineFunctionSpinlockLock(volatile LONG *Lock, BOOL *HasError) {
#ifdef SCRIPT_ENGINE_USER_MODE SpinlockLock(Lock);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Lock, sizeof(LONG))) {
    *HasError = TRUE;
    return;
  }
  SpinlockLock(Lock);
#endif
}

VOID ScriptEngineFunctionSpinlockUnlock(volatile LONG *Lock, BOOL *HasError) {
#ifdef SCRIPT_ENGINE_USER_MODE SpinlockUnlock(Lock);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Lock, sizeof(LONG))) {
    *HasError = TRUE;
    return;
  }
  SpinlockUnlock(Lock);
#endif
}

VOID ScriptEngineFunctionSpinlockLockCustomWait(volatile long *Lock,
                                                unsigned MaxWait,
                                                BOOL *HasError) {
#ifdef SCRIPT_ENGINE_USER_MODE SpinlockLockWithCustomWait(Lock, MaxWait);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Lock, sizeof(LONG))) {
    *HasError = TRUE;
    return;
  }
  SpinlockLockWithCustomWait(Lock, MaxWait);
#endif
}

UINT64 ScriptEngineFunctionStrlen(const char *Address) {
  UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = strlen(Address);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  Result = VmFuncVmxCompatibleStrlen(Address);
#endif
  return Result;
}

UINT64 ScriptEngineFunctionDisassembleLen(PVOID Address, BOOLEAN Is32Bit) {
  UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = HyperDbgLengthDisassemblerEngine(
      (unsigned char *)Address, MAXIMUM_INSTR_SIZE, Is32Bit ? FALSE : TRUE);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  Result = DisassemblerLengthDisassembleEngineInVmxRootOnTargetProcess(Address,
                                                                       Is32Bit);
#endif
  return Result;
}

UINT64 ScriptEngineFunctionWcslen(const wchar_t *Address) {
  UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = wcslen(Address);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  Result = VmFuncVmxCompatibleWcslen(Address);
#endif
  return Result;
}

long long ScriptEngineFunctionInterlockedExchange(long long volatile *Target,
                                                  long long Value,
                                                  BOOL *HasError) {
  long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Target, sizeof(long long))) {
    *HasError = TRUE;
    return (long long)NULL;
  }
#endif
  Result = InterlockedExchange64(Target, Value);
  return Result;
}

long long ScriptEngineFunctionInterlockedExchangeAdd(long long volatile *Addend,
                                                     long long Value,
                                                     BOOL *HasError) {
  long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Addend, sizeof(long long))) {
    *HasError = TRUE;
    return (long long)NULL;
  }
#endif
  Result = InterlockedExchangeAdd64(Addend, Value);
  return Result;
}

long long ScriptEngineFunctionInterlockedIncrement(long long volatile *Addend,
                                                   BOOL *HasError) {
  long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Addend, sizeof(long long))) {
    *HasError = TRUE;
    return (long long)NULL;
  }
#endif
  Result = InterlockedIncrement64(Addend);
  return Result;
}

long long ScriptEngineFunctionInterlockedDecrement(long long volatile *Addend,
                                                   BOOL *HasError) {
  long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Addend, sizeof(long long))) {
    *HasError = TRUE;
    return (long long)NULL;
  }
#endif
  Result = InterlockedDecrement64(Addend);
  return Result;
}

long long ScriptEngineFunctionInterlockedCompareExchange(
    long long volatile *Destination, long long ExChange, long long Comperand,
    BOOL *HasError) {
  long long Result = 0;
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!CheckAccessValidityAndSafety((UINT64)Destination, sizeof(long long))) {
    *HasError = TRUE;
    return (long long)NULL;
  }
#endif
  Result = InterlockedCompareExchange64(Destination, ExChange, Comperand);
  return Result;
}

VOID ScriptEngineFunctionEventEnable(UINT64 EventId) {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages( \
    "err, enabling events is not possible in user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!DebuggerEnableEvent(EventId + DebuggerEventTagStartSeed)) {
    LogInfo("Invalid tag id (%x)", EventId);
  }
#endif
}

VOID ScriptEngineFunctionEventDisable(UINT64 EventId) {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages( \
    "err, disabling events is not possible in user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (!DebuggerDisableEvent(EventId + DebuggerEventTagStartSeed)) {
    LogInfo("Invalid tag id (%x)", EventId);
  }
#endif
}

VOID ScriptEngineFunctionEventClear(UINT64 EventId) {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages( \
    "err, disabling events is not possible in user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  BOOLEAN PoolManagerAllocatedMemory = FALSE;
  if (g_KernelDebuggerState && EnableInstantEventMechanism) {
    PoolManagerAllocatedMemory = TRUE;
  }
  if (!DebuggerClearEvent(EventId + DebuggerEventTagStartSeed,
                          VmFuncVmxGetCurrentExecutionMode(),
                          PoolManagerAllocatedMemory)) {
    LogInfo("Invalid tag id (%x)", EventId);
  }
#endif
}

VOID ScriptEngineFunctionPause(ACTION_BUFFER *ActionDetail,
                               PGUEST_REGS GuestRegs) {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages( \
    "err, breaking is not possible in user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE if (g_KernelDebuggerState &&
      g_DebuggeeHaltReason == DEBUGGEE_PAUSING_REASON_NOT_PAUSED) {
        DEBUGGER_TRIGGERED_EVENT_DETAILS TriggeredEventDetail = {0};
        ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
        TriggeredEventDetail.Tag = ActionDetail->Tag;
        TriggeredEventDetail.Context = (PVOID)ActionDetail->Context;
        if (ActionDetail->CallingStage == 1) {
          TriggeredEventDetail.Stage =
              VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION;
        } else {
          TriggeredEventDetail.Stage =
              VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION;
        }
        if (VmFuncVmxGetCurrentExecutionMode() == TRUE) {
          KdHandleBreakpointAndDebugBreakpointsCallback(
              CurrentCore, DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED,
              &TriggeredEventDetail);
        } else {
          VmFuncVmxVmcall(
              DEBUGGER_VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT,
              (UINT64)&TriggeredEventDetail, (UINT64)GuestRegs, (UINT64)NULL);
        }
      }
      else {
        LogInfo(
            "The 'pause();' function is either called from the vmi-mode or is "
            "evaluated by the '?' command. It's not allowed to use it on "
            "vmi-mode "
            "(local debugging) or by the '?' command");
      }
#endif
}

VOID ScriptEngineFunctionFlush() {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages( \
    "err, it's not possible to flush buffers in user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE LogMarkAllAsRead(TRUE);
  LogMarkAllAsRead(FALSE);
#endif
}

VOID ScriptEngineFunctionShortCircuitingEvent(UINT64 State,
                                              ACTION_BUFFER *ActionDetail) {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages( \
    "err, it's not possible to short-circuit events in user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE if (ActionDetail->CallingStage == 1) {
  LogWarning(
      "Warning, calling the 'event_sc' function in the 'post' calling stage "
      "doesn't make sense as the emulation is already performed!\n"
      "You can use this function in the 'pre' calling stage");
  return;
}
ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
if (State != 0) {
  g_DbgState[CurrentCore].ShortCircuitingEvent = TRUE;
} else {
  g_DbgState[CurrentCore].ShortCircuitingEvent = FALSE;
}
#endif
}

VOID ScriptEngineFunctionFormats(UINT64 Tag, BOOLEAN ImmediateMessagePassing,
                                 UINT64 Value) {
#ifdef SCRIPT_ENGINE_USER_MODE ScriptEngineFunctionTestStatement( \
    Tag, ImmediateMessagePassing, Value);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE if (g_KernelDebuggerState) {
  KdSendFormatsFunctionResult(Value);
}
else {
  char TempBuffer[20] = {0};
  UINT32 TempBufferLen = sprintf(TempBuffer, "%llx\n", Value);
  LogSimpleWithTag((UINT32)Tag, ImmediateMessagePassing, TempBuffer,
                   TempBufferLen + 1);
}
#endif
}

UINT32 CustomStrlen(UINT64 StrAddr, BOOLEAN IsWstring) {
#ifdef SCRIPT_ENGINE_USER_MODE if (IsWstring) {
  return (UINT32)wcslen((const wchar_t *)StrAddr);
}
else {
  return (UINT32)strlen((const char *)StrAddr);
}
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE if (IsWstring) {
return VmFuncVmxCompatibleWcslen((const wchar_t *)StrAddr);
}
else {
  return VmFuncVmxCompatibleStrlen((const CHAR *)StrAddr);
}
#endif
}

BOOLEAN CheckIfStringIsSafe(UINT64 StrAddr, BOOLEAN IsWstring) {
#ifdef SCRIPT_ENGINE_USER_MODE
  return TRUE;
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  if (CheckAccessValidityAndSafety(StrAddr, IsWstring ? 4 : 2)) {
    return TRUE;
  } else {
    return FALSE;
  }
#endif
}

VOID ApplyFormatSpecifier(const CHAR *CurrentSpecifier, CHAR *FinalBuffer,
                          PUINT32 CurrentProcessedPositionFromStartOfFormat,
                          PUINT32 CurrentPositionInFinalBuffer, UINT64 Val,
                          UINT32 SizeOfFinalBuffer) {
  UINT32 TempBufferLen = 0;
  CHAR TempBuffer[50 + 1] = {0};
  *CurrentProcessedPositionFromStartOfFormat =
      *CurrentProcessedPositionFromStartOfFormat +
      (UINT32)strlen(CurrentSpecifier);
  sprintf_s(TempBuffer, sizeof(TempBuffer), CurrentSpecifier, Val);
  TempBufferLen = (UINT32)strlen(TempBuffer);
  if (*CurrentPositionInFinalBuffer + TempBufferLen > SizeOfFinalBuffer) {
    return;
  }
  memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], TempBuffer,
         TempBufferLen);
  *CurrentPositionInFinalBuffer = *CurrentPositionInFinalBuffer + TempBufferLen;
}

size_t WcharToChar(const wchar_t *src, char *dest, size_t dest_len) {
  wchar_t Code;
  size_t i;
  i = 0;
  while ((src[i] != '\0') && i < (dest_len - 1)) {
    Code = src[i];
    if (Code < 128)
      dest[i] = (char)Code;
    else {
      dest[i] = '?';
      if (Code >= 0xD800 && Code <= 0xD8FF) {
        i++;
      }
    }
    i++;
  }
  return i - 1;
}

BOOLEAN ApplyStringFormatSpecifier(
    const CHAR *CurrentSpecifier, CHAR *FinalBuffer,
    PUINT32 CurrentProcessedPositionFromStartOfFormat,
    PUINT32 CurrentPositionInFinalBuffer, UINT64 Val, BOOLEAN IsWstring,
    UINT32 SizeOfFinalBuffer) {
  UINT32 StringSize;
  wchar_t WstrBuffer[50];
  CHAR AsciiBuffer[sizeof(WstrBuffer) / 2];
  UINT32 StringSizeInByte;
  UINT32 CountOfBlocks;
  UINT32 CopiedBlockLen;
  if (!CheckIfStringIsSafe(Val, IsWstring)) {
    return FALSE;
  }
  *CurrentProcessedPositionFromStartOfFormat +=
      (UINT32)strlen(CurrentSpecifier);
  StringSize = CustomStrlen(Val, IsWstring);
  if (*CurrentPositionInFinalBuffer + StringSize > SizeOfFinalBuffer) {
    return TRUE;
  }
  if (IsWstring) {
    StringSizeInByte = StringSize * 2;
    if (StringSizeInByte % sizeof(WstrBuffer) == 0) {
      CountOfBlocks = StringSizeInByte / sizeof(WstrBuffer);
    } else {
      CountOfBlocks = (StringSizeInByte / sizeof(WstrBuffer)) + 1;
    }
    for (size_t i = 0; i < CountOfBlocks; i++) {
      RtlZeroMemory(WstrBuffer, sizeof(WstrBuffer));
      RtlZeroMemory(AsciiBuffer, sizeof(AsciiBuffer));
      if (i == CountOfBlocks - 1) {
#ifdef SCRIPT_ENGINE_USER_MODE
        memcpy(WstrBuffer, (void *)(Val + (i * sizeof(WstrBuffer))),
               StringSizeInByte % sizeof(WstrBuffer));
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
        MemoryMapperReadMemorySafeOnTargetProcess(
            (UINT64)(Val + (i * sizeof(WstrBuffer))), WstrBuffer,
            StringSizeInByte % sizeof(WstrBuffer));
#endif
      } else {
#ifdef SCRIPT_ENGINE_USER_MODE
        memcpy(WstrBuffer, (void *)(Val + (i * sizeof(WstrBuffer))),
               sizeof(WstrBuffer));
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
        MemoryMapperReadMemorySafeOnTargetProcess(
            (UINT64)(Val + (i * sizeof(WstrBuffer))), WstrBuffer,
            sizeof(WstrBuffer));
#endif
      }
      CopiedBlockLen =
          (UINT32)WcharToChar(WstrBuffer, AsciiBuffer, sizeof(AsciiBuffer) + 1);
      memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], (void *)AsciiBuffer,
             CopiedBlockLen + 1);
      *CurrentPositionInFinalBuffer += CopiedBlockLen + 1;
    }
  } else {
#ifdef SCRIPT_ENGINE_USER_MODE
    memcpy(&FinalBuffer[*CurrentPositionInFinalBuffer], (void *)Val,
           StringSize);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE MemoryMapperReadMemorySafeOnTargetProcess(
        Val, &FinalBuffer[*CurrentPositionInFinalBuffer], StringSize);
#endif
        *CurrentPositionInFinalBuffer += StringSize;
  }
  return TRUE;
}

VOID ScriptEngineFunctionPrintf(
    PGUEST_REGS GuestRegs, ACTION_BUFFER *ActionDetail,
    SCRIPT_ENGINE_GENERAL_REGISTERS *ScriptGeneralRegisters, UINT64 Tag,
    BOOLEAN ImmediateMessagePassing, char *Format, UINT64 ArgCount,
    PSYMBOL FirstArg, BOOLEAN *HasError) {
  char FinalBuffer[PacketChunkSize] = {0};
  UINT32 CurrentPositionInFinalBuffer = 0;
  UINT32 CurrentProcessedPositionFromStartOfFormat = 0;
  BOOLEAN WithoutAnyFormatSpecifier = TRUE;
  UINT64 Val;
  UINT32 Position;
  UINT32 LenOfFormats = (UINT32)strlen(Format) + 1;
  PSYMBOL Symbol;
  *HasError = FALSE;
  for (int i = 0; i < ArgCount; i++) {
    WithoutAnyFormatSpecifier = FALSE;
    Symbol = FirstArg + i;
    Position = (Symbol->Type >> 32) + 1;
    SYMBOL TempSymbol = {0};
    memcpy(&TempSymbol, Symbol, sizeof(SYMBOL));
    TempSymbol.Type &= 0x7fffffff;
    Val = GetValue(GuestRegs, ActionDetail, ScriptGeneralRegisters, &TempSymbol,
                   FALSE);
    CHAR PercentageChar = Format[Position];

    if (CurrentProcessedPositionFromStartOfFormat != Position) {
      UINT32 StringLen = Position - CurrentProcessedPositionFromStartOfFormat;
      if (CurrentPositionInFinalBuffer + StringLen < sizeof(FinalBuffer)) {
        memcpy(&FinalBuffer[CurrentPositionInFinalBuffer],
               &Format[CurrentProcessedPositionFromStartOfFormat], StringLen);
        CurrentProcessedPositionFromStartOfFormat += StringLen;
        CurrentPositionInFinalBuffer += StringLen;
      }
    }
    if (PercentageChar == '%') {
      CHAR FormatSpecifier[5] = {0};
      FormatSpecifier[0] = '%';
      CHAR IndicatorChar2 = Format[Position + 1];
      if (IndicatorChar2 == 'l' || IndicatorChar2 == 'w' ||
          IndicatorChar2 == 'h') {
        FormatSpecifier[1] = IndicatorChar2;
        if (IndicatorChar2 == 'l' && Format[Position + 2] == 'l') {
          FormatSpecifier[2] = 'l';
          FormatSpecifier[3] = Format[Position + 3];
        } else {
          FormatSpecifier[2] = Format[Position + 2];
        }
      } else {
        FormatSpecifier[1] = IndicatorChar2;
      }
      if (!strncmp(FormatSpecifier, "%s", 2)) {
        if (!ApplyStringFormatSpecifier(
                "%s", FinalBuffer, &CurrentProcessedPositionFromStartOfFormat,
                &CurrentPositionInFinalBuffer, Val, FALSE,
                sizeof(FinalBuffer))) {
          *HasError = TRUE;
          return;
        }
      } else if (!strncmp(FormatSpecifier, "%ls", 3) ||
                 !strncmp(FormatSpecifier, "%ws", 3)) {
        if (!ApplyStringFormatSpecifier(
                "%ws", FinalBuffer, &CurrentProcessedPositionFromStartOfFormat,
                &CurrentPositionInFinalBuffer, Val, TRUE,
                sizeof(FinalBuffer))) {
          *HasError = TRUE;
          return;
        }
      } else {
        ApplyFormatSpecifier(FormatSpecifier, FinalBuffer,
                             &CurrentProcessedPositionFromStartOfFormat,
                             &CurrentPositionInFinalBuffer, Val,
                             sizeof(FinalBuffer));
      }
    }
  }
  if (WithoutAnyFormatSpecifier) {
    if (LenOfFormats < sizeof(FinalBuffer)) {
      memcpy(FinalBuffer, Format, LenOfFormats);
    }
  } else {
    if (LenOfFormats > CurrentProcessedPositionFromStartOfFormat) {
      UINT32 RemainedLen =
          LenOfFormats - CurrentProcessedPositionFromStartOfFormat;
      if (CurrentPositionInFinalBuffer + RemainedLen < sizeof(FinalBuffer)) {
        memcpy(&FinalBuffer[CurrentPositionInFinalBuffer],
               &Format[CurrentProcessedPositionFromStartOfFormat], RemainedLen);
      }
    }
  }
#ifdef SCRIPT_ENGINE_USER_MODE printf("%s", FinalBuffer);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  LogSimpleWithTag((UINT32)Tag, ImmediateMessagePassing, FinalBuffer,
                   (UINT32)strlen(FinalBuffer) + 1);
#endif
}

VOID ScriptEngineFunctionEventInject(UINT32 InterruptionType, UINT32 Vector,
                                     BOOL *HasError) {
  UNREFERENCED_PARAMETER(HasError);
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages( \
    "err, event_inject is not supported in user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE if (Vector >= 256 || InterruptionType >= 8) {
  LogInfo("Err, invalid event vector or interruption type is specified");
  return;
}
else {
  VmFuncEventInjectInterruption(InterruptionType, Vector, FALSE, 0);
}
#endif
}

VOID ScriptEngineFunctionEventInjectErrorCode(UINT32 InterruptionType,
                                              UINT32 Vector, UINT32 ErrorCode,
                                              BOOL *HasError) {
  UNREFERENCED_PARAMETER(HasError);
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages( \
    "err, event_inject is not supported in user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE if (Vector >= 256 || InterruptionType >= 8) {
  LogInfo("Err, invalid event vector or interruption type is specified");
  return;
}
else {
  VmFuncEventInjectInterruption(InterruptionType, Vector, TRUE, ErrorCode);
}
#endif
}

UINT64 ScriptEngineFunctionStrcmp(const char *Address1, const char *Address2) {
  UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = strcmp(Address1, Address2);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  Result = VmFuncVmxCompatibleStrcmp(Address1, Address2);
#endif
  return Result;
}

UINT64 ScriptEngineFunctionStrncmp(const char *Address1, const char *Address2,
                                   size_t Num) {
  UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = strncmp(Address1, Address2, Num);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  Result = VmFuncVmxCompatibleStrncmp(Address1, Address2, Num);
#endif
  return Result;
}

UINT64 ScriptEngineFunctionWcscmp(const wchar_t *Address1,
                                  const wchar_t *Address2) {
  UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = wcscmp(Address1, Address2);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  Result = VmFuncVmxCompatibleWcscmp(Address1, Address2);
#endif
  return Result;
}

UINT64 ScriptEngineFunctionWcsncmp(const wchar_t *Address1,
                                   const wchar_t *Address2, size_t Num) {
  UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = wcsncmp(Address1, Address2, Num);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  Result = VmFuncVmxCompatibleWcsncmp(Address1, Address2, Num);
#endif
  return Result;
}

UINT64 ScriptEngineFunctionMemcmp(const char *Address1, const char *Address2,
                                  size_t Count) {
  UINT64 Result = 0;
#ifdef SCRIPT_ENGINE_USER_MODE
  Result = memcmp(Address1, Address2, Count);
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
  Result = VmFuncVmxCompatibleMemcmp(Address1, Address2, Count);
#endif
  return Result;
}

VOID ScriptEngineFunctionEventTraceInstrumentationStep() {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages(
      "err, it's not possible to call event_trace_instrumentation_step "
      "function in the user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
      ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
      TracingPerformInstrumentationStepIn(&g_DbgState[CurrentCore]);
#endif
}

VOID ScriptEngineFunctionEventTraceStepIn() {
#ifdef SCRIPT_ENGINE_USER_MODE ShowMessages(
      "err, it's not possible to call event_trace_instrumentation_step "
      "function in the user-mode\n");
#endif
#ifdef SCRIPT_ENGINE_KERNEL_MODE
      ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
      TracingPerformRegularStepInInstruction(&g_DbgState[CurrentCore]);
#endif
}
