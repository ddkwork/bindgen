
#pragma once

#define USE_DEFAULT_OS_GDT_AS_HOST_GDT FALSE

#define HOST_GDT_DESCRIPTOR_COUNT 10

#define HOST_INTERRUPT_STACK_SIZE 4 * PAGE_SIZE

#define USE_INTERRUPT_STACK_TABLE FALSE
_Success_(return) BOOLEAN
    SegmentGetDescriptor(_In_ PUCHAR GdtBase, _In_ UINT16 Selector,
                         _Out_ PVMX_SEGMENT_SELECTOR SegmentSelector);
VOID SegmentPrepareHostGdt(SEGMENT_DESCRIPTOR_32 *OsGdtBase, UINT16 OsGdtLimit,
                           UINT16 TrSelector, UINT64 HostStack,
                           SEGMENT_DESCRIPTOR_32 *AllocatedHostGdt,
                           TASK_STATE_SEGMENT_64 *AllocatedHostTss);
