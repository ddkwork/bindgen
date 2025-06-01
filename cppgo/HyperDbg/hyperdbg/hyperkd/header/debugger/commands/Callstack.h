
#pragma once
BOOLEAN CallstackWalkthroughStack(
    PDEBUGGER_SINGLE_CALLSTACK_FRAME AddressToSaveFrames, UINT32 *FrameCount,
    UINT64 StackBaseAddress, UINT32 Size, BOOLEAN Is32Bit);
