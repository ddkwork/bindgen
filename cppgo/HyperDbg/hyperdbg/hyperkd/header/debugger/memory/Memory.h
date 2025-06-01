
#pragma once

typedef struct _PAGE_ENTRY {
  union {
    UINT64 Flags;
    PML4E_64 Pml4;
    PDPTE_1GB_64 PdptLarge;
    PDPTE_64 Pdpt;
    PDE_2MB_64 PdLarge;
    PDE_64 Pd;
    PTE_64 Pt;

    struct {
      UINT64 Present : 1;
      UINT64 Write : 1;
      UINT64 Supervisor : 1;
      UINT64 PageLevelWriteThrough : 1;
      UINT64 PageLevelCacheDisable : 1;
      UINT64 Accessed : 1;
      UINT64 Dirty : 1;
      UINT64 LargePage : 1;
      UINT64 Global : 1;
      UINT64 Ignored1 : 3;
      UINT64 PageFrameNumber : 36;
      UINT64 Reserved1 : 4;
      UINT64 Ignored2 : 7;
      UINT64 ProtectionKey : 4;
      UINT64 ExecuteDisable : 1;
    } Fields;
  };
} PAGE_ENTRY, *PPAGE_ENTRY;
