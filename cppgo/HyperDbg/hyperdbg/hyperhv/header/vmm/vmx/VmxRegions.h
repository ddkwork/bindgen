
#pragma once
_Success_(return != FALSE) BOOLEAN
    VmxAllocateVmxonRegion(_Out_ VIRTUAL_MACHINE_STATE *VCpu);
_Success_(return != FALSE) BOOLEAN
    VmxAllocateVmcsRegion(_Out_ VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN VmxAllocateVmmStack(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN VmxAllocateMsrBitmap(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN VmxAllocateIoBitmaps(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
UINT64 *VmxAllocateInvalidMsrBimap();
BOOLEAN VmxAllocateHostIdt(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN VmxAllocateHostGdt(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN VmxAllocateHostTss(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN VmxAllocateHostInterruptStack(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
