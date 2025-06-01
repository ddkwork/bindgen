
#pragma once
VOID VmxBroadcastInitialize();
VOID VmxBroadcastUninitialize();
BOOLEAN VmxBroadcastHandleNmiCallback(PVOID Context, BOOLEAN Handled);
BOOLEAN VmxBroadcastNmi(VIRTUAL_MACHINE_STATE *VCpu,
                        NMI_BROADCAST_ACTION_TYPE VmxBroadcastAction);
BOOLEAN VmxBroadcastNmiHandler(VIRTUAL_MACHINE_STATE *VCpu,
                               BOOLEAN IsOnVmxNmiHandler);
