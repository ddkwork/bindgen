
#pragma once

typedef enum _SEGMENT_REGISTERS {
  ES = 0,
  CS,
  SS,
  DS,
  FS,
  GS,
  LDTR,
  TR
} SEGMENT_REGISTERS;

#define PASSIVE_LEVEL 0
#define LOW_LEVEL 0
#define APC_LEVEL 1
#define DISPATCH_LEVEL 2
#define CMCI_LEVEL 5
#define CLOCK_LEVEL 13
#define IPI_LEVEL 14
#define DRS_LEVEL 14
#define POWER_LEVEL 14
#define PROFILE_LEVEL 15
#define HIGH_LEVEL 15

#define X86_CR0_PE 0x00000001
#define X86_CR0_MP 0x00000002
#define X86_CR0_EM 0x00000004
#define X86_CR0_TS 0x00000008
#define X86_CR0_ET 0x00000010
#define X86_CR0_NE 0x00000020
#define X86_CR0_WP 0x00010000
#define X86_CR0_AM 0x00040000
#define X86_CR0_NW 0x20000000
#define X86_CR0_CD 0x40000000
#define X86_CR0_PG 0x80000000

#define X86_CR4_VME 0x0001
#define X86_CR4_PVI 0x0002
#define X86_CR4_TSD 0x0004
#define X86_CR4_DE 0x0008
#define X86_CR4_PSE 0x0010
#define X86_CR4_PAE 0x0020
#define X86_CR4_MCE 0x0040
#define X86_CR4_PGE 0x0080
#define X86_CR4_PCE 0x0100
#define X86_CR4_OSFXSR 0x0200
#define X86_CR4_OSXMMEXCPT 0x0400
#define X86_CR4_VMXE 0x2000

#define X86_FLAGS_CF (1 << 0)
#define X86_FLAGS_PF (1 << 2)
#define X86_FLAGS_AF (1 << 4)
#define X86_FLAGS_ZF (1 << 6)
#define X86_FLAGS_SF (1 << 7)
#define X86_FLAGS_TF (1 << 8)
#define X86_FLAGS_IF (1 << 9)
#define X86_FLAGS_DF (1 << 10)
#define X86_FLAGS_OF (1 << 11)
#define X86_FLAGS_STATUS_MASK (0xfff)
#define X86_FLAGS_IOPL_MASK (3 << 12)
#define X86_FLAGS_IOPL_SHIFT (12)
#define X86_FLAGS_IOPL_SHIFT_2ND_BIT (13)
#define X86_FLAGS_NT (1 << 14)
#define X86_FLAGS_RF (1 << 16)
#define X86_FLAGS_VM (1 << 17)
#define X86_FLAGS_AC (1 << 18)
#define X86_FLAGS_VIF (1 << 19)
#define X86_FLAGS_VIP (1 << 20)
#define X86_FLAGS_ID (1 << 21)
#define X86_FLAGS_RESERVED_ONES 0x2
#define X86_FLAGS_RESERVED 0xffc0802a
#define X86_FLAGS_RESERVED_BITS 0xffc38028
#define X86_FLAGS_FIXED 0x00000002

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

#define PCID_NONE 0x000
#define PCID_MASK 0x003

#define CPUID_HV_VENDOR_AND_MAX_FUNCTIONS 0x40000000
#define CPUID_HV_INTERFACE 0x40000001

#define CPUID_ADDR_WIDTH 0x80000008

#define CPUID_PROCESSOR_AND_PROCESSOR_FEATURE_IDENTIFIERS 0x00000001

#define RESERVED_MSR_RANGE_LOW 0x40000000
#define RESERVED_MSR_RANGE_HI 0x400000F0

#define __CPU_INDEX__ KeGetCurrentProcessorNumberEx(NULL)

#define ALIGNMENT_PAGE_SIZE 4096

#define MAXIMUM_ADDRESS 0xffffffffffffffff

#define DPL_USER 3
#define DPL_SYSTEM 0

#define RPL_MASK 3
#define BITS_PER_LONG (sizeof(unsigned long) * 8)
#define ORDER_LONG (sizeof(unsigned long) == 4 ? 5 : 6)
#define BITMAP_ENTRY(_nr, _bmap) ((_bmap))[(_nr) / BITS_PER_LONG]
#define BITMAP_SHIFT(_nr) ((_nr) % BITS_PER_LONG)

#define PAGE_OFFSET(Va) ((PVOID)((ULONG_PTR)(Va) & (PAGE_SIZE - 1)))

#define _XBEGIN_STARTED (~0u)
#define _XABORT_EXPLICIT (1 << 0)
#define _XABORT_RETRY (1 << 1)
#define _XABORT_CONFLICT (1 << 2)
#define _XABORT_CAPACITY (1 << 3)
#define _XABORT_DEBUG (1 << 4)
#define _XABORT_NESTED (1 << 5)
#ifndef _XABORT_CODE
#define _XABORT_CODE(x) (((x) >> 24) & 0xFF)
#endif
typedef SEGMENT_DESCRIPTOR_32 *PSEGMENT_DESCRIPTOR;

typedef struct _CPUID {
  int eax;
  int ebx;
  int ecx;
  int edx;
} CPUID, *PCPUID;

typedef union _CR_FIXED {
  UINT64 Flags;

  struct {
    unsigned long Low;
    long High;
  } Fields;
} CR_FIXED, *PCR_FIXED;

typedef struct _NT_KPROCESS {
  DISPATCHER_HEADER Header;
  LIST_ENTRY ProfileListHead;
  ULONG_PTR DirectoryTableBase;
  UCHAR Data[1];
} NT_KPROCESS, *PNT_KPROCESS;

typedef void (*RunOnLogicalCoreFunc)(ULONG ProcessorId);
UCHAR *PsGetProcessImageFileName(IN PEPROCESS Process);
static NTSTATUS CommonGetHandleFromProcess(_In_ UINT32 ProcessId,
                                           _Out_ PHANDLE Handle);
BOOLEAN CommonAffinityBroadcastToProcessors(_In_ ULONG ProcessorNumber,
                                            _In_ RunOnLogicalCoreFunc Routine);
BOOLEAN CommonIsStringStartsWith(const char *pre, const char *str);
BOOLEAN CommonIsGuestOnUsermode32Bit();
PCHAR CommonGetProcessNameFromProcessControlBlock(PEPROCESS eprocess);
VOID CommonCpuidInstruction(UINT32 Func, UINT32 SubFunc, int *CpuInfo);
VOID CommonWriteDebugInformation(VIRTUAL_MACHINE_STATE *VCpu);
BOOLEAN CommonIsXCr0Valid(XCR0 XCr0);
