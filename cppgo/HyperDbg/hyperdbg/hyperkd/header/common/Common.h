
#pragma once
UCHAR *PsGetProcessImageFileName(IN PEPROCESS Process);
PVOID PsGetProcessSectionBaseAddress(PEPROCESS Process);
NTKERNELAPI NTSTATUS NTAPI SeCreateAccessState(PACCESS_STATE AccessState,
                                               PVOID AuxData,
                                               ACCESS_MASK DesiredAccess,
                                               PGENERIC_MAPPING Mapping);
NTKERNELAPI VOID NTAPI SeDeleteAccessState(PACCESS_STATE AccessState);
NTSTATUS MmUnmapViewOfSection(PEPROCESS Process, PVOID BaseAddress);

typedef struct _NT_KPROCESS {
  DISPATCHER_HEADER Header;
  LIST_ENTRY ProfileListHead;
  ULONG_PTR DirectoryTableBase;
  UCHAR Data[1];
} NT_KPROCESS, *PNT_KPROCESS;

typedef enum _PROCESS_KILL_METHODS {
  PROCESS_KILL_METHOD_1 = 0,
  PROCESS_KILL_METHOD_2,
  PROCESS_KILL_METHOD_3,
} PROCESS_KILL_METHODS;

#define KGDT64_NULL (0 * 16)
#define KGDT64_R0_CODE (1 * 16)
#define KGDT64_R0_DATA (1 * 16) + 8
#define KGDT64_R3_CMCODE (2 * 16)
#define KGDT64_R3_DATA (2 * 16) + 8
#define KGDT64_R3_CODE (3 * 16)
#define KGDT64_SYS_TSS (4 * 16)
#define KGDT64_R3_CMTEB (5 * 16)
#define KGDT64_R0_CMCODE (6 * 16)
#define KGDT64_LAST (7 * 16)
BOOLEAN CommonIsProcessExist(UINT32 ProcId);
PCHAR CommonGetProcessNameFromProcessControlBlock(PEPROCESS Eprocess);
BOOLEAN CommonKillProcess(UINT32 ProcessId, PROCESS_KILL_METHODS KillingMethod);
BOOLEAN CommonValidateCoreNumber(UINT32 CoreNumber);
