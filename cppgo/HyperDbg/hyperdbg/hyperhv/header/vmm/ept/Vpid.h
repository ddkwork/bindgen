
#pragma once

#define VPID_TAG 0x1
VOID VpidInvvpid(_In_ INVVPID_TYPE Type,
                 _In_opt_ INVVPID_DESCRIPTOR *Descriptor);
VOID VpidInvvpidIndividualAddress(_In_ UINT16 Vpid, _In_ UINT64 LinearAddress);
VOID VpidInvvpidSingleContext(UINT16 Vpid);
VOID VpidInvvpidAllContext();
VOID VpidInvvpidSingleContextRetainingGlobals(UINT16 Vpid);
