
#include "pch.h"

_Use_decl_annotations_

    VOID
    VpidInvvpid(INVVPID_TYPE Type, INVVPID_DESCRIPTOR *Descriptor) {
  INVVPID_DESCRIPTOR *TargetDescriptor = NULL;
  INVVPID_DESCRIPTOR ZeroDescriptor = {0};
  if (!Descriptor) {
    TargetDescriptor = &ZeroDescriptor;
  } else {
    TargetDescriptor = Descriptor;
  }
  AsmInvvpid(Type, TargetDescriptor);
}

_Use_decl_annotations_

    VOID
    VpidInvvpidIndividualAddress(UINT16 Vpid, UINT64 LinearAddress) {
  INVVPID_DESCRIPTOR Descriptor = {Vpid, 0, 0, LinearAddress};
  VpidInvvpid(InvvpidIndividualAddress, &Descriptor);
}

VOID VpidInvvpidSingleContext(UINT16 Vpid) {
  INVVPID_DESCRIPTOR Descriptor = {Vpid, 0, 0, 0};
  VpidInvvpid(InvvpidSingleContext, &Descriptor);
}

VOID VpidInvvpidAllContext() { VpidInvvpid(InvvpidAllContext, NULL); }

VOID VpidInvvpidSingleContextRetainingGlobals(UINT16 Vpid) {
  INVVPID_DESCRIPTOR Descriptor = {Vpid, 0, 0, 0};
  VpidInvvpid(InvvpidSingleContextRetainingGlobals, &Descriptor);
}
