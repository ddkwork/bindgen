
#pragma once
VOID BroadcastVmxVirtualizationAllCores();
VOID BroadcastEnableNmiExitingAllCores();
VOID BroadcastDisableNmiExitingAllCores();
VOID BroadcastNotifyAllToInvalidateEptAllCores();
VOID BroadcastEnablePmlOnAllProcessors();
VOID BroadcastDisablePmlOnAllProcessors();
VOID BroadcastChangeToMbecSupportedEptpOnAllProcessors();
VOID BroadcastRestoreToNormalEptpOnAllProcessors();
VOID BroadcasDisableMbecOnAllProcessors();
VOID BroadcasEnableMbecOnAllProcessors();
