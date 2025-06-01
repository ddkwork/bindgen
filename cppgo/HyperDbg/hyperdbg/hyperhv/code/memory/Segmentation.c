
#include "pch.h"

_Use_decl_annotations_

    BOOLEAN
    SegmentGetDescriptor(PUCHAR GdtBase, UINT16 Selector,
                         PVMX_SEGMENT_SELECTOR SegmentSelector) {
  SEGMENT_DESCRIPTOR_32 *DescriptorTable32;
  SEGMENT_DESCRIPTOR_32 *Descriptor32;
  SEGMENT_SELECTOR SegSelector = {.AsUInt = Selector};
  if (!SegmentSelector) return FALSE;
#define SELECTOR_TABLE_LDT 0x1
#define SELECTOR_TABLE_GDT 0x0
  if ((Selector == 0x0) || (SegSelector.Table != SELECTOR_TABLE_GDT)) {
    return FALSE;
  }
  DescriptorTable32 = (SEGMENT_DESCRIPTOR_32 *)(GdtBase);
  Descriptor32 = &DescriptorTable32[SegSelector.Index];
  SegmentSelector->Selector = Selector;
  SegmentSelector->Limit = __segmentlimit(Selector);
  SegmentSelector->Base = ((UINT64)Descriptor32->BaseAddressLow |
                           (UINT64)Descriptor32->BaseAddressMiddle << 16 |
                           (UINT64)Descriptor32->BaseAddressHigh << 24);
  SegmentSelector->Attributes.AsUInt = (AsmGetAccessRights(Selector) >> 8);
  if (SegSelector.Table == 0 && SegSelector.Index == 0) {
    SegmentSelector->Attributes.Unusable = TRUE;
  }
  if ((Descriptor32->Type == SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY) ||
      (Descriptor32->Type == SEGMENT_DESCRIPTOR_TYPE_CALL_GATE)) {
    UINT64 SegmentLimitHigh;
    SegmentLimitHigh = (*(UINT64 *)((PUCHAR)Descriptor32 + 8));
    SegmentSelector->Base =
        (SegmentSelector->Base & 0xffffffff) | (SegmentLimitHigh << 32);
  }
  if (SegmentSelector->Attributes.Granularity) {
    SegmentSelector->Limit = (SegmentSelector->Limit << 12) + 0xfff;
  }
  return TRUE;
}

VOID SegmentPrepareHostGdt(SEGMENT_DESCRIPTOR_32 *OsGdtBase, UINT16 OsGdtLimit,
                           UINT16 TrSelector, UINT64 HostStack,
                           SEGMENT_DESCRIPTOR_32 *AllocatedHostGdt,
                           TASK_STATE_SEGMENT_64 *AllocatedHostTss) {
  RtlCopyBytes(AllocatedHostGdt, OsGdtBase, OsGdtLimit + 1);
  RtlZeroBytes(AllocatedHostTss, sizeof(TASK_STATE_SEGMENT_64));
#if USE_INTERRUPT_STACK_TABLE == TRUE
  UINT64 EndOfStack = 0;
  EndOfStack = ((UINT64)HostStack + HOST_INTERRUPT_STACK_SIZE - 1);
  EndOfStack = ((UINT64)((ULONG_PTR)(EndOfStack) & ~(16 - 1)));
  LogDebugInfo("Host Interrupt Stack, from: %llx, to: %llx", HostStack,
               EndOfStack);
  AllocatedHostTss->Rsp0 = EndOfStack;
  AllocatedHostTss->Rsp1 = EndOfStack;
  AllocatedHostTss->Rsp2 = EndOfStack;
  AllocatedHostTss->Ist1 = EndOfStack;
  AllocatedHostTss->Ist2 = EndOfStack;
  AllocatedHostTss->Ist3 = EndOfStack;
  AllocatedHostTss->Ist4 = EndOfStack;
  AllocatedHostTss->Ist5 = EndOfStack;
  AllocatedHostTss->Ist6 = EndOfStack;
  AllocatedHostTss->Ist7 = EndOfStack;
#else UNREFERENCED_PARAMETER(HostStack);
#endif
  SEGMENT_DESCRIPTOR_64 *GdtTssDesc =
      (SEGMENT_DESCRIPTOR_64 *)&AllocatedHostGdt[TrSelector];
  UINT64 Base = (UINT64)AllocatedHostTss;
  GdtTssDesc->BaseAddressLow = (Base >> 00) & 0xFFFF;
  GdtTssDesc->BaseAddressMiddle = (Base >> 16) & 0xFF;
  GdtTssDesc->BaseAddressHigh = (Base >> 24) & 0xFF;
  GdtTssDesc->BaseAddressUpper = (Base >> 32) & 0xFFFFFFFF;
}
