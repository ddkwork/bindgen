
#pragma once
typedef EPT_PML4E EPT_PML4_POINTER, *PEPT_PML4_POINTER;
typedef EPT_PDPTE EPT_PML3_POINTER, *PEPT_PML3_POINTER;
typedef EPT_PDE_2MB EPT_PML2_ENTRY, *PEPT_PML2_ENTRY;
typedef EPT_PDE EPT_PML2_POINTER, *PEPT_PML2_POINTER;
typedef EPT_PTE EPT_PML1_ENTRY, *PEPT_PML1_ENTRY;

#define PENDING_INTERRUPTS_BUFFER_CAPACITY 64

#define MaximumHiddenBreakpointsOnPage 40

typedef enum _NMI_BROADCAST_ACTION_TYPE {
  NMI_BROADCAST_ACTION_NONE = 0,
  NMI_BROADCAST_ACTION_TEST,
  NMI_BROADCAST_ACTION_REQUEST,
  NMI_BROADCAST_ACTION_INVALIDATE_EPT_CACHE_SINGLE_CONTEXT,
  NMI_BROADCAST_ACTION_INVALIDATE_EPT_CACHE_ALL_CONTEXTS,
} NMI_BROADCAST_ACTION_TYPE;

typedef enum _EPT_HOOKED_LAST_VIOLATION {
  EPT_HOOKED_LAST_VIOLATION_READ = 1,
  EPT_HOOKED_LAST_VIOLATION_WRITE = 2,
  EPT_HOOKED_LAST_VIOLATION_EXEC = 3
} EPT_HOOKED_LAST_VIOLATION;

#define VMM_EPT_PML4E_COUNT 512

#define VMM_EPT_PML3E_COUNT 512

#define VMM_EPT_PML2E_COUNT 512

#define VMM_EPT_PML1E_COUNT 512

typedef struct _VMM_EPT_PAGE_TABLE {
  DECLSPEC_ALIGN(PAGE_SIZE)
  EPT_PML4_POINTER PML4[VMM_EPT_PML4E_COUNT];

  DECLSPEC_ALIGN(PAGE_SIZE)
  EPT_PML3_POINTER PML3[VMM_EPT_PML3E_COUNT];

  DECLSPEC_ALIGN(PAGE_SIZE)
  EPT_PML2_ENTRY PML2[VMM_EPT_PML3E_COUNT][VMM_EPT_PML2E_COUNT];
} VMM_EPT_PAGE_TABLE, *PVMM_EPT_PAGE_TABLE;

typedef struct _VM_EXIT_TRANSPARENCY {
  UINT64 PreviousTimeStampCounter;
  HANDLE ThreadId;
  UINT64 RevealedTimeStampCounterByRdtsc;
  BOOLEAN CpuidAfterRdtscDetected;
} VM_EXIT_TRANSPARENCY, *PVM_EXIT_TRANSPARENCY;

typedef struct _VMX_VMXOFF_STATE {
  BOOLEAN IsVmxoffExecuted;
  UINT64 GuestRip;
  UINT64 GuestRsp;
} VMX_VMXOFF_STATE, *PVMX_VMXOFF_STATE;

typedef struct _EPT_HOOKED_PAGE_DETAIL {
  DECLSPEC_ALIGN(PAGE_SIZE)
  CHAR FakePageContents[PAGE_SIZE];

  LIST_ENTRY PageHookList;

  UINT64 VirtualAddress;

  UINT64 AddressOfEptHook2sDetourListEntry;

  SIZE_T PhysicalBaseAddress;

  SIZE_T StartOfTargetPhysicalAddress;

  SIZE_T EndOfTargetPhysicalAddress;

  UINT64 HookingTag;

  SIZE_T PhysicalBaseAddressOfFakePageContents;

  EPT_PML1_ENTRY OriginalEntry;

  EPT_PML1_ENTRY ChangedEntry;

  PCHAR Trampoline;

  BOOLEAN IsExecutionHook;

  BOOLEAN IsHiddenBreakpoint;

  BOOLEAN IsMmioShadowing;

  EPT_HOOKS_CONTEXT LastContextState;

  BOOLEAN IsPostEventTriggerAllowed;

  EPT_HOOKED_LAST_VIOLATION LastViolation;

  UINT64 BreakpointAddresses[MaximumHiddenBreakpointsOnPage];

  CHAR PreviousBytesOnBreakpointAddresses[MaximumHiddenBreakpointsOnPage];

  UINT64 CountOfBreakpoints;
} EPT_HOOKED_PAGE_DETAIL, *PEPT_HOOKED_PAGE_DETAIL;

typedef struct _NMI_BROADCASTING_STATE {
  volatile NMI_BROADCAST_ACTION_TYPE NmiBroadcastAction;
} NMI_BROADCASTING_STATE, *PNMI_BROADCASTING_STATE;

typedef struct _VIRTUAL_MACHINE_STATE {
  BOOLEAN IsOnVmxRootMode;
  BOOLEAN IncrementRip;
  BOOLEAN HasLaunched;
  BOOLEAN IgnoreMtfUnset;
  BOOLEAN WaitForImmediateVmexit;
  BOOLEAN EnableExternalInterruptsOnContinue;
  BOOLEAN EnableExternalInterruptsOnContinueMtf;
  BOOLEAN RegisterBreakOnMtf;
  BOOLEAN IgnoreOneMtf;
  BOOLEAN NotNormalEptp;
  BOOLEAN MbecEnabled;
  PUINT64 PmlBufferAddress;
  BOOLEAN Test;
  UINT64 TestNumber;
  GUEST_REGS *Regs;
  UINT32 CoreId;
  UINT32 ExitReason;
  UINT32 ExitQualification;
  UINT64 LastVmexitRip;
  UINT64 VmxonRegionPhysicalAddress;
  UINT64 VmxonRegionVirtualAddress;
  UINT64 VmcsRegionPhysicalAddress;
  UINT64 VmcsRegionVirtualAddress;
  UINT64 VmmStack;
  UINT64 MsrBitmapVirtualAddress;
  UINT64 MsrBitmapPhysicalAddress;
  UINT64 IoBitmapVirtualAddressA;
  UINT64 IoBitmapPhysicalAddressA;
  UINT64 IoBitmapVirtualAddressB;
  UINT64 IoBitmapPhysicalAddressB;
  UINT32 QueuedNmi;
  UINT32 PendingExternalInterrupts[PENDING_INTERRUPTS_BUFFER_CAPACITY];
  VMX_VMXOFF_STATE VmxoffState;
  NMI_BROADCASTING_STATE NmiBroadcastingState;
  VM_EXIT_TRANSPARENCY TransparencyState;
  PEPT_HOOKED_PAGE_DETAIL MtfEptHookRestorePoint;
  UINT8 LastExceptionOccuredInHost;
  UINT64 HostIdt;
  UINT64 HostGdt;
  UINT64 HostTss;
  UINT64 HostInterruptStack;
  EPT_POINTER EptPointer;
  PVMM_EPT_PAGE_TABLE EptPageTable;
} VIRTUAL_MACHINE_STATE, *PVIRTUAL_MACHINE_STATE;
