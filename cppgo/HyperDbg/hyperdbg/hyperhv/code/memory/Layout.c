
#include "pch.h"

_Use_decl_annotations_

    CR3_TYPE
    LayoutGetCr3ByProcessId(UINT32 ProcessId) {
  PEPROCESS TargetEprocess;
  CR3_TYPE ProcessCr3 = {0};
  if (PsLookupProcessByProcessId((HANDLE)ProcessId, &TargetEprocess) !=
      STATUS_SUCCESS) {
    return ProcessCr3;
  }
  NT_KPROCESS *CurrentProcess = (NT_KPROCESS *)(TargetEprocess);
  ProcessCr3.Flags = CurrentProcess->DirectoryTableBase;
  ObDereferenceObject(TargetEprocess);
  return ProcessCr3;
}

CR3_TYPE LayoutGetCurrentProcessCr3() {
  CR3_TYPE GuestCr3;
  NT_KPROCESS *CurrentProcess = (NT_KPROCESS *)(PsGetCurrentProcess());
  GuestCr3.Flags = CurrentProcess->DirectoryTableBase;
  return GuestCr3;
}

CR3_TYPE LayoutGetExactGuestProcessCr3() {
  CR3_TYPE GuestCr3 = {0};
  __vmx_vmread(VMCS_GUEST_CR3, &GuestCr3.Flags);
  return GuestCr3;
}

UINT64 LayoutGetSystemDirectoryTableBase() {
  NT_KPROCESS *SystemProcess = (NT_KPROCESS *)(PsInitialSystemProcess);
  return SystemProcess->DirectoryTableBase;
}
