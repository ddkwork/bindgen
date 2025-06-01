
#pragma once
VOID DebuggerEventEnableEferOnAllProcessors();
VOID DebuggerEventDisableEferOnAllProcessors();
VOID DebuggerEventEnableMovToCr3ExitingOnAllProcessors();
VOID DebuggerEventDisableMovToCr3ExitingOnAllProcessors();
BOOLEAN DebuggerEventEnableMonitorReadWriteExec(
    EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *HookingDetails,
    UINT32 ProcessId, BOOLEAN ApplyDirectlyFromVmxRoot);
BOOLEAN DebuggerCheckProcessOrThreadChange(_In_ UINT32 CoreId);
