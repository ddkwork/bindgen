
#pragma once
VOID HaltedRoutineChangeAllMsrBitmapReadOnSingleCore(UINT32 TargetCoreId,
                                                     UINT64 BitmapMask);
VOID HaltedRoutineChangeAllMsrBitmapWriteOnSingleCore(UINT32 TargetCoreId,
                                                      UINT64 BitmapMask);
VOID HaltedRoutineChangeIoBitmapOnSingleCore(UINT32 TargetCoreId, UINT64 Port);
VOID HaltedRoutineEnableRdpmcExitingOnSingleCore(UINT32 TargetCoreId);
VOID HaltedRoutineEnableRdtscExitingOnSingleCore(UINT32 TargetCoreId);
VOID HaltedRoutineEnableMov2DebugRegsExitingOnSingleCore(UINT32 TargetCoreId);
VOID HaltedRoutineEnableExternalInterruptExiting(UINT32 TargetCoreId);
VOID HaltedRoutineSetExceptionBitmapOnSingleCore(UINT32 TargetCoreId,
                                                 UINT64 ExceptionIndex);
VOID HaltedRoutineUnSetExceptionBitmapOnSingleCore(UINT32 TargetCoreId,
                                                   UINT64 ExceptionIndex);
VOID HaltedRoutineEnableMovToCrExitingOnSingleCore(
    UINT32 TargetCoreId, DEBUGGER_EVENT_OPTIONS *BroadcastingOption);
VOID HaltedRoutineEnableEferSyscallHookOnSingleCore(UINT32 TargetCoreId);
VOID HaltedRoutineInvalidateEptAllContextsOnSingleCore(UINT32 TargetCoreId);
VOID HaltedRoutineInvalidateSingleContextOnSingleCore(UINT32 TargetCoreId);
