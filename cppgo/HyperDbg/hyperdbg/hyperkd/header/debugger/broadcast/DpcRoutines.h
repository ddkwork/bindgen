
#pragma once
NTSTATUS DpcRoutineRunTaskOnSingleCore(UINT32 CoreNumber, PVOID Routine,
                                       PVOID DeferredContext);
VOID DpcRoutinePerformWriteMsr(KDPC *Dpc, PVOID DeferredContext,
                               PVOID SystemArgument1, PVOID SystemArgument2);
VOID DpcRoutinePerformReadMsr(KDPC *Dpc, PVOID DeferredContext,
                              PVOID SystemArgument1, PVOID SystemArgument2);
VOID DpcRoutineWriteMsrToAllCores(KDPC *Dpc, PVOID DeferredContext,
                                  PVOID SystemArgument1, PVOID SystemArgument2);
VOID DpcRoutineReadMsrToAllCores(KDPC *Dpc, PVOID DeferredContext,
                                 PVOID SystemArgument1, PVOID SystemArgument2);
VOID DpcRoutineVmExitAndHaltSystemAllCores(KDPC *Dpc, PVOID DeferredContext,
                                           PVOID SystemArgument1,
                                           PVOID SystemArgument2);
