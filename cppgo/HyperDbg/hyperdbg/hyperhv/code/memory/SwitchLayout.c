
#include "pch.h"

_Use_decl_annotations_

    CR3_TYPE
    SwitchToProcessMemoryLayout(UINT32 ProcessId) {
  UINT64 GuestCr3;
  PEPROCESS TargetEprocess;
  CR3_TYPE CurrentProcessCr3 = {0};
  if (PsLookupProcessByProcessId((HANDLE)ProcessId, &TargetEprocess) !=
      STATUS_SUCCESS) {
    return CurrentProcessCr3;
  }
  NT_KPROCESS *CurrentProcess = (NT_KPROCESS *)(TargetEprocess);
  GuestCr3 = CurrentProcess->DirectoryTableBase;
  CurrentProcessCr3.Flags = __readcr3();
  __writecr3(GuestCr3);
  ObDereferenceObject(TargetEprocess);
  return CurrentProcessCr3;
}

CR3_TYPE SwitchToCurrentProcessMemoryLayout() {
  CR3_TYPE GuestCr3;
  CR3_TYPE CurrentProcessCr3 = {0};
  GuestCr3.Flags = LayoutGetCurrentProcessCr3().Flags;
  CurrentProcessCr3.Flags = __readcr3();
  __writecr3(GuestCr3.Flags);
  return CurrentProcessCr3;
}

_Use_decl_annotations_

    CR3_TYPE
    SwitchToProcessMemoryLayoutByCr3(CR3_TYPE TargetCr3) {
  CR3_TYPE CurrentProcessCr3 = {0};
  CurrentProcessCr3.Flags = __readcr3();
  __writecr3(TargetCr3.Flags);
  return CurrentProcessCr3;
}

_Use_decl_annotations_

    VOID
    SwitchToPreviousProcess(CR3_TYPE PreviousProcess) {
  __writecr3(PreviousProcess.Flags);
}
