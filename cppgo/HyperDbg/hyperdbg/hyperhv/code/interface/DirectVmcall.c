
#include "pch.h"

NTSTATUS DirectVmcallTest(UINT32 CoreId,
                          DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId], VMCALL_TEST,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallPerformVmcall(
    UINT32 CoreId, UINT64 VmcallNumber,
    DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId], VmcallNumber,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallChangeMsrBitmapRead(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_CHANGE_MSR_BITMAP_READ,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallChangeMsrBitmapWrite(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_CHANGE_MSR_BITMAP_WRITE,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallChangeIoBitmap(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_CHANGE_IO_BITMAP, DirectVmcallOptions);
}

NTSTATUS DirectVmcallEnableRdpmcExiting(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_SET_RDPMC_EXITING, DirectVmcallOptions);
}

NTSTATUS DirectVmcallEnableRdtscpExiting(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_SET_RDTSC_EXITING, DirectVmcallOptions);
}

NTSTATUS DirectVmcallEnableMov2DebugRegsExiting(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallSetExceptionBitmap(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_SET_EXCEPTION_BITMAP, DirectVmcallOptions);
}

NTSTATUS DirectVmcallEnableExternalInterruptExiting(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallEnableMovToCrExiting(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallEnableEferSyscall(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_ENABLE_SYSCALL_HOOK_EFER,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallSetHiddenBreakpointHook(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_SET_HIDDEN_CC_BREAKPOINT,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallInvalidateEptAllContexts(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_INVEPT_ALL_CONTEXTS, DirectVmcallOptions);
}

NTSTATUS DirectVmcallInvalidateSingleContext(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_INVEPT_SINGLE_CONTEXT, DirectVmcallOptions);
}

NTSTATUS DirectVmcallUnsetExceptionBitmap(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_UNSET_EXCEPTION_BITMAP,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallUnhookSinglePage(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_UNHOOK_SINGLE_PAGE, DirectVmcallOptions);
}

NTSTATUS
DirectVmcallSetDisableExternalInterruptExitingOnlyOnClearingInterruptEvents(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId],
      VMCALL_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS,
      DirectVmcallOptions);
}

NTSTATUS DirectVmcallResetMsrBitmapRead(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_RESET_MSR_BITMAP_READ, DirectVmcallOptions);
}

NTSTATUS DirectVmcallResetMsrBitmapWrite(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_RESET_MSR_BITMAP_WRITE,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallResetExceptionBitmapOnlyOnClearingExceptionEvents(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId],
      VMCALL_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS,
      DirectVmcallOptions);
}

NTSTATUS DirectVmcallResetIoBitmap(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_RESET_IO_BITMAP, DirectVmcallOptions);
}

NTSTATUS DirectVmcallDisableRdtscExitingForClearingTscEvents(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS,
      DirectVmcallOptions);
}

NTSTATUS DirectVmcallDisableRdpmcExiting(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId], VMCALL_UNSET_RDPMC_EXITING, DirectVmcallOptions);
}

NTSTATUS DirectVmcallDisableEferSyscallEvents(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(&g_GuestState[CoreId],
                                      VMCALL_DISABLE_SYSCALL_HOOK_EFER,
                                      DirectVmcallOptions);
}

NTSTATUS DirectVmcallDisableMov2DrExitingForClearingDrEvents(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId],
      VMCALL_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS,
      DirectVmcallOptions);
}

NTSTATUS DirectVmcallDisableMov2CrExitingForClearingCrEvents(
    UINT32 CoreId, DIRECT_VMCALL_PARAMETERS *DirectVmcallOptions) {
  return VmxVmcallDirectVmcallHandler(
      &g_GuestState[CoreId],
      VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS, DirectVmcallOptions);
}
