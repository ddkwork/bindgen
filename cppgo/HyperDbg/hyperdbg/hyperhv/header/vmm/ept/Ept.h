
#pragma once

#define PAGE_ATTRIB_READ 0x2
#define PAGE_ATTRIB_WRITE 0x4
#define PAGE_ATTRIB_EXEC 0x8
#define PAGE_ATTRIB_EXEC_HIDDEN_HOOK 0x10

#define SIZE_2_MB ((SIZE_T)(512 * PAGE_SIZE))

#define ADDRMASK_EPT_PML1_OFFSET(_VAR_) ((_VAR_) & 0xFFFULL)

#define ADDRMASK_EPT_PML1_INDEX(_VAR_) (((_VAR_) & 0x1FF000ULL) >> 12)

#define ADDRMASK_EPT_PML2_INDEX(_VAR_) (((_VAR_) & 0x3FE00000ULL) >> 21)

#define ADDRMASK_EPT_PML3_INDEX(_VAR_) (((_VAR_) & 0x7FC0000000ULL) >> 30)

#define ADDRMASK_EPT_PML4_INDEX(_VAR_) (((_VAR_) & 0xFF8000000000ULL) >> 39)

typedef struct _MTRR_RANGE_DESCRIPTOR {
  SIZE_T PhysicalBaseAddress;
  SIZE_T PhysicalEndAddress;
  UCHAR MemoryType;
  BOOLEAN FixedRange;
} MTRR_RANGE_DESCRIPTOR, *PMTRR_RANGE_DESCRIPTOR;

typedef union _IA32_MTRR_FIXED_RANGE_TYPE {
  UINT64 AsUInt;

  struct {
    UINT8 Types[8];
  } s;
} IA32_MTRR_FIXED_RANGE_TYPE;

#define MAX_VARIABLE_RANGE_MTRRS 255

#define NUM_FIXED_RANGE_MTRRS \
  ((1 + 2 + 8) * RTL_NUMBER_OF_FIELD(IA32_MTRR_FIXED_RANGE_TYPE, s.Types))

#define NUM_MTRR_ENTRIES (MAX_VARIABLE_RANGE_MTRRS + NUM_FIXED_RANGE_MTRRS)

typedef struct _EPT_STATE {
  LIST_ENTRY HookedPagesList;
  MTRR_RANGE_DESCRIPTOR MemoryRanges[NUM_MTRR_ENTRIES];
  UINT32 NumberOfEnabledMemoryRanges;
  PVMM_EPT_PAGE_TABLE EptPageTable;
  PVMM_EPT_PAGE_TABLE ModeBasedUserDisabledEptPageTable;
  PVMM_EPT_PAGE_TABLE ModeBasedKernelDisabledEptPageTable;
  EPT_POINTER ModeBasedUserDisabledEptPointer;
  EPT_POINTER ModeBasedKernelDisabledEptPointer;
  EPT_POINTER ExecuteOnlyEptPointer;
  UINT8 DefaultMemoryType;
} EPT_STATE, *PEPT_STATE;

typedef struct _VMM_EPT_DYNAMIC_SPLIT {
  DECLSPEC_ALIGN(PAGE_SIZE)
  EPT_PML1_ENTRY PML1[VMM_EPT_PML1E_COUNT];

  union {
    PEPT_PML2_ENTRY Entry;
    PEPT_PML2_POINTER Pointer;
  } u;

  LIST_ENTRY DynamicSplitList;
} VMM_EPT_DYNAMIC_SPLIT, *PVMM_EPT_DYNAMIC_SPLIT;

BOOLEAN EptSetupPML2Entry(PVMM_EPT_PAGE_TABLE EptPageTable,
                          PEPT_PML2_ENTRY NewEntry, SIZE_T PageFrameNumber);
BOOLEAN EptHandlePageHookExit(_Inout_ VIRTUAL_MACHINE_STATE *VCpu,
                              _In_ VMX_EXIT_QUALIFICATION_EPT_VIOLATION
                                  ViolationQualification,
                              _In_ UINT64 GuestPhysicalAddr);

BOOLEAN EptCheckFeatures(VOID);

BOOLEAN EptBuildMtrrMap(VOID);

PVMM_EPT_PAGE_TABLE EptAllocateAndCreateIdentityPageTable(VOID);

BOOLEAN EptSplitLargePage(PVMM_EPT_PAGE_TABLE EptPageTable,
                          PVOID PreAllocatedBuffer, SIZE_T PhysicalAddress);

PEPT_PML2_ENTRY EptGetPml2Entry(PVMM_EPT_PAGE_TABLE EptPageTable,
                                SIZE_T PhysicalAddress);

BOOLEAN EptLogicalProcessorInitialize(VOID);

BOOLEAN EptHandleEptViolation(VIRTUAL_MACHINE_STATE *VCpu);

PEPT_PML1_ENTRY EptGetPml1Entry(PVMM_EPT_PAGE_TABLE EptPageTable,
                                SIZE_T PhysicalAddress);

PVOID EptGetPml1OrPml2Entry(PVMM_EPT_PAGE_TABLE EptPageTable,
                            SIZE_T PhysicalAddress, BOOLEAN *IsLargePage);

VOID EptHandleMisconfiguration(VOID);

VOID EptSetPML1AndInvalidateTLB(
    _Inout_ VIRTUAL_MACHINE_STATE *VCpu, _Out_ PEPT_PML1_ENTRY EntryAddress,
    _In_ EPT_PML1_ENTRY EntryValue,
    _In_ _Strict_type_match_ INVEPT_TYPE InvalidationType);

BOOLEAN EptCheckAndHandleBreakpoint(VIRTUAL_MACHINE_STATE *VCpu);
