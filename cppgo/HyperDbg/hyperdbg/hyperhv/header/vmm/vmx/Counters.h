
#pragma once
VOID CounterEmulateRdtsc(VIRTUAL_MACHINE_STATE *VCpu);
VOID CounterEmulateRdtscp(VIRTUAL_MACHINE_STATE *VCpu);
VOID CounterEmulateRdpmc(VIRTUAL_MACHINE_STATE *VCpu);
VOID CounterSetPreemptionTimer(UINT32 TimerValue);
VOID CounterClearPreemptionTimer();
