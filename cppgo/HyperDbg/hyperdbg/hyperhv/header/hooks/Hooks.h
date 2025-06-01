
#pragma once

#define IS_SYSRET_INSTRUCTION(Code)                                  \
  (*((PUINT8)(Code) + 0) == 0x48 && *((PUINT8)(Code) + 1) == 0x0F && \
   *((PUINT8)(Code) + 2) == 0x07)
#define IS_SYSCALL_INSTRUCTION(Code) \
  (*((PUINT8)(Code) + 0) == 0x0F && *((PUINT8)(Code) + 1) == 0x05)

#define IMAGE_DOS_SIGNATURE 0x5A4D
#define IMAGE_OS2_SIGNATURE 0x454E
#define IMAGE_OS2_SIGNATURE_LE 0x454C
#define IMAGE_VXD_SIGNATURE 0x454C
#define IMAGE_NT_SIGNATURE 0x00004550

typedef struct _SSDTStruct {
  LONG *pServiceTable;
  PVOID pCounterTable;
#ifdef _WIN64
  UINT64 NumberOfServices;
#else
  ULONG NumberOfServices;
#endif
  PCHAR pArgumentTable;
} SSDTStruct, *PSSDTStruct;

typedef struct _HIDDEN_HOOKS_DETOUR_DETAILS {
  LIST_ENTRY OtherHooksList;
  PVOID HookedFunctionAddress;
  PVOID ReturnAddress;
} HIDDEN_HOOKS_DETOUR_DETAILS, *PHIDDEN_HOOKS_DETOUR_DETAILS;

typedef struct _SYSTEM_MODULE_ENTRY {
  HANDLE Section;
  PVOID MappedBase;
  PVOID ImageBase;
  ULONG ImageSize;
  ULONG Flags;
  UINT16 LoadOrderIndex;
  UINT16 InitOrderIndex;
  UINT16 LoadCount;
  UINT16 OffsetToFileName;
  UCHAR FullPathName[256];
} SYSTEM_MODULE_ENTRY, *PSYSTEM_MODULE_ENTRY;

typedef struct _SYSTEM_MODULE_INFORMATION {
  ULONG Count;
  SYSTEM_MODULE_ENTRY Module;
} SYSTEM_MODULE_INFORMATION, *PSYSTEM_MODULE_INFORMATION;

typedef enum _SYSTEM_INFORMATION_CLASS {
  SystemModuleInformation = 11,
  SystemKernelDebuggerInformation = 35
} SYSTEM_INFORMATION_CLASS,
    *PSYSTEM_INFORMATION_CLASS;

typedef NTSTATUS(NTAPI *ZWQUERYSYSTEMINFORMATION)(
    IN SYSTEM_INFORMATION_CLASS SystemInformationClass,
    OUT PVOID SystemInformation, IN ULONG SystemInformationLength,
    OUT PULONG ReturnLength OPTIONAL);
NTSTATUS(*NtCreateFileOrig)
(PHANDLE FileHandle, ACCESS_MASK DesiredAccess,
 POBJECT_ATTRIBUTES ObjectAttributes, PIO_STATUS_BLOCK IoStatusBlock,
 PLARGE_INTEGER AllocationSize, ULONG FileAttributes, ULONG ShareAccess,
 ULONG CreateDisposition, ULONG CreateOptions, PVOID EaBuffer, ULONG EaLength);
VOID SyscallHookTest();
VOID SyscallHookConfigureEFER(VIRTUAL_MACHINE_STATE *VCpu,
                              BOOLEAN EnableEFERSyscallHook);
BOOLEAN SyscallHookHandleUD(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN SyscallHookEmulateSYSRET(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN SyscallHookEmulateSYSCALL(_Inout_ VIRTUAL_MACHINE_STATE *VCpu);
PVOID(*ExAllocatePoolWithTagOrig)
(POOL_TYPE PoolType, SIZE_T NumberOfBytes, ULONG Tag);

#define MAX_EXEC_TRAMPOLINE_SIZE 100

BOOLEAN EptHookPerformPageHook(VIRTUAL_MACHINE_STATE *VCpu, PVOID TargetAddress,
                               CR3_TYPE ProcessCr3);

BOOLEAN EptHookPerformPageHookMonitorAndInlineHook(VIRTUAL_MACHINE_STATE *VCpu,
                                                   PVOID HookingDetails,
                                                   CR3_TYPE ProcessCr3,
                                                   UINT32 PageHookMask);

BOOLEAN EptHook(PVOID TargetAddress, UINT32 ProcessId);

BOOLEAN EptHookFromVmxRoot(PVOID TargetAddress);

BOOLEAN EptHookInlineHook(VIRTUAL_MACHINE_STATE *VCpu, PVOID TargetAddress,
                          PVOID HookFunction, UINT32 ProcessId);

BOOLEAN EptHookMonitorHook(
    VIRTUAL_MACHINE_STATE *VCpu,
    EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *HookingDetails,
    UINT32 ProcessId);

BOOLEAN EptHookInlineHookFromVmxRoot(VIRTUAL_MACHINE_STATE *VCpu,
                                     PVOID TargetAddress, PVOID HookFunction);

BOOLEAN EptHookMonitorFromVmxRoot(
    VIRTUAL_MACHINE_STATE *VCpu,
    EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR *MemoryAddressDetails);

BOOLEAN EptHookHandleHookedPage(
    VIRTUAL_MACHINE_STATE *VCpu, EPT_HOOKED_PAGE_DETAIL *HookedEntryDetails,
    VMX_EXIT_QUALIFICATION_EPT_VIOLATION ViolationQualification,
    SIZE_T PhysicalAddress, EPT_HOOKS_CONTEXT *LastContext,
    BOOLEAN *IgnoreReadOrWriteOrExec, BOOLEAN *IsExecViolation);

BOOLEAN EptHookRestoreSingleHookToOriginalEntry(VIRTUAL_MACHINE_STATE *VCpu,
                                                SIZE_T PhysicalAddress,
                                                UINT64 OriginalEntry);

VOID EptHookRestoreAllHooksToOriginalEntry(VIRTUAL_MACHINE_STATE *VCpu);

VOID EptHookUnHookAll();

BOOLEAN EptHookUnHookAllByHookingTag(UINT64 HookingTag);

BOOLEAN EptHookUnHookSingleAddress(UINT64 VirtualAddress, UINT64 PhysAddress,
                                   UINT32 ProcessId);

BOOLEAN EptHookUnHookSingleHookByHookingTagFromVmxRoot(
    UINT64 HookingTag,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails);

BOOLEAN EptHookUnHookSingleAddressFromVmxRoot(
    UINT64 VirtualAddress, UINT64 PhysAddress,
    EPT_SINGLE_HOOK_UNHOOKING_DETAILS *TargetUnhookingDetails);

UINT32 EptHookGetCountOfEpthooks(BOOLEAN IsEptHook2);

BOOLEAN EptHookRemoveEntryAndFreePoolFromEptHook2sDetourList(UINT64 Address);

PVOID EptHook2GeneralDetourEventHandler(PGUEST_REGS Regs, PVOID CalledFrom);

VOID EptHookAllocateExtraHookingPagesForMemoryMonitorsAndExecEptHooks(
    UINT32 Count);

VOID EptHookReservePreallocatedPoolsForEptHooks(UINT32 Count);

BOOLEAN EptHookModifyInstructionFetchState(VIRTUAL_MACHINE_STATE *VCpu,
                                           PVOID PhysicalAddress,
                                           BOOLEAN IsUnset);

BOOLEAN EptHookModifyPageReadState(VIRTUAL_MACHINE_STATE *VCpu,
                                   PVOID PhysicalAddress, BOOLEAN IsUnset);

BOOLEAN EptHookModifyPageWriteState(VIRTUAL_MACHINE_STATE *VCpu,
                                    PVOID PhysicalAddress, BOOLEAN IsUnset);

VOID EptHookHandleMonitorTrapFlag(VIRTUAL_MACHINE_STATE *VCpu);
