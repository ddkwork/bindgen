
#include "pch.h"

BOOLEAN CommonIsProcessExist(UINT32 ProcId) {
  PEPROCESS TargetEprocess;
  if (PsLookupProcessByProcessId((HANDLE)ProcId, &TargetEprocess) !=
      STATUS_SUCCESS) {
    return FALSE;
  } else {
    ObDereferenceObject(TargetEprocess);
    return TRUE;
  }
}

_Use_decl_annotations_

    NTSTATUS
    CommonGetHandleFromProcess(UINT32 ProcessId, PHANDLE Handle) {
  NTSTATUS Status;
  Status = STATUS_SUCCESS;
  OBJECT_ATTRIBUTES ObjAttr = {0};
  CLIENT_ID Cid = {0};
  InitializeObjectAttributes(&ObjAttr, NULL, 0, NULL, NULL);
  Cid.UniqueProcess = (HANDLE)ProcessId;
  Cid.UniqueThread = (HANDLE)0;
  Status = ZwOpenProcess(Handle, PROCESS_ALL_ACCESS, &ObjAttr, &Cid);
  return Status;
}

PCHAR CommonGetProcessNameFromProcessControlBlock(PEPROCESS Eprocess) {
  PCHAR Result = 0;
  Result = (CHAR *)PsGetProcessImageFileName(Eprocess);
  return Result;
}

NTSTATUS CommonUndocumentedNtOpenProcess(PHANDLE ProcessHandle,
                                         ACCESS_MASK DesiredAccess,
                                         HANDLE ProcessId,
                                         KPROCESSOR_MODE AccessMode) {
  NTSTATUS Status = STATUS_SUCCESS;
  ACCESS_STATE AccessState;
  CHAR AuxData[0x200] = {0};
  PEPROCESS ProcessObject = NULL;
  HANDLE ProcHandle = NULL;
  Status = SeCreateAccessState(&AccessState, AuxData, DesiredAccess,
                               (PGENERIC_MAPPING)((PCHAR)*PsProcessType + 52));
  if (!NT_SUCCESS(Status)) {
    return Status;
  }
  AccessState.PreviouslyGrantedAccess |= AccessState.RemainingDesiredAccess;
  AccessState.RemainingDesiredAccess = 0;
  Status = PsLookupProcessByProcessId(ProcessId, &ProcessObject);
  if (!NT_SUCCESS(Status)) {
    SeDeleteAccessState(&AccessState);
    return Status;
  }
  Status = ObOpenObjectByPointer(ProcessObject, 0, &AccessState, 0,
                                 *PsProcessType, AccessMode, &ProcHandle);
  SeDeleteAccessState(&AccessState);
  ObDereferenceObject(ProcessObject);
  if (NT_SUCCESS(Status)) *ProcessHandle = ProcHandle;
  return Status;
}

_Use_decl_annotations_

    BOOLEAN
    CommonKillProcess(UINT32 ProcessId, PROCESS_KILL_METHODS KillingMethod) {
  NTSTATUS Status = STATUS_SUCCESS;
  HANDLE ProcessHandle = NULL;
  PEPROCESS Process = NULL;
  if (ProcessId == NULL_ZERO) {
    return FALSE;
  }
  switch (KillingMethod) {
    case PROCESS_KILL_METHOD_1:
      Status = CommonGetHandleFromProcess(ProcessId, &ProcessHandle);
      if (!NT_SUCCESS(Status) || ProcessHandle == NULL) {
        return FALSE;
      }
      Status = ZwTerminateProcess(ProcessHandle, 0);
      if (!NT_SUCCESS(Status)) {
        return FALSE;
      }
      break;
    case PROCESS_KILL_METHOD_2:
      CommonUndocumentedNtOpenProcess(&ProcessHandle, PROCESS_ALL_ACCESS,
                                      (HANDLE)ProcessId, KernelMode);
      if (ProcessHandle == NULL) {
        return FALSE;
      }
      Status = ZwTerminateProcess(ProcessHandle, 0);
      if (!NT_SUCCESS(Status)) {
        return FALSE;
      }
      break;
    case PROCESS_KILL_METHOD_3:
      Status = MmUnmapViewOfSection(Process,
                                    PsGetProcessSectionBaseAddress(Process));
      ObDereferenceObject(Process);
      break;
    default:
      return FALSE;
      break;
  }
  return TRUE;
}

_Use_decl_annotations_

    BOOLEAN
    CommonValidateCoreNumber(UINT32 CoreNumber) {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  if (CoreNumber >= ProcessorsCount) {
    return FALSE;
  } else {
    return TRUE;
  }
}
