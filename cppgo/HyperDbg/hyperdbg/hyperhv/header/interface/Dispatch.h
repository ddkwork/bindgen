
#pragma once
VOID DispatchEventEferSysret(VIRTUAL_MACHINE_STATE *VCpu, PVOID Context);
VOID DispatchEventEferSyscall(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventCpuid(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventTsc(VIRTUAL_MACHINE_STATE *VCpu, BOOLEAN IsRdtscp);
VOID DispatchEventVmcall(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventMode(VIRTUAL_MACHINE_STATE *VCpu,
                       DEBUGGER_EVENT_MODE_TYPE TargetMode,
                       BOOLEAN HandleState);
VOID DispatchEventMovToCr3(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventIO(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventRdmsr(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventWrmsr(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventRdpmc(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventMovToFromControlRegisters(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventMov2DebugRegs(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventException(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventExternalInterrupts(VIRTUAL_MACHINE_STATE *VCpu);
VOID DispatchEventHiddenHookExecCc(VIRTUAL_MACHINE_STATE *VCpu, PVOID Context);
VOID DispatchEventHiddenHookExecDetours(VIRTUAL_MACHINE_STATE *VCpu,
                                        PVOID Context);
VOID DispatchEventHiddenHookPageReadWriteExecReadPostEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context);
VOID DispatchEventHiddenHookPageReadWriteExecWritePostEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context);
VOID DispatchEventHiddenHookPageReadWriteExecExecutePostEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context);
VOID DispatchEventInstrumentationTrace(VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN DispatchEventHiddenHookPageReadWriteExecuteReadPreEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context,
    BOOLEAN *IsTriggeringPostEventAllowed);
BOOLEAN DispatchEventHiddenHookPageReadWriteExecuteWritePreEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context,
    BOOLEAN *IsTriggeringPostEventAllowed);
BOOLEAN DispatchEventHiddenHookPageReadWriteExecuteExecutePreEvent(
    VIRTUAL_MACHINE_STATE *VCpu, PVOID Context,
    BOOLEAN *IsTriggeringPostEventAllowed);
