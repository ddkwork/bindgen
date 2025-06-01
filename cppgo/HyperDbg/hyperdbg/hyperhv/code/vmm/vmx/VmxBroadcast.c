
#include "pch.h"

VOID VmxBroadcastInitialize() {
  ApicInitialize();
#if USE_DEFAULT_OS_IDT_AS_HOST_IDT == TRUE
  g_NmiHandlerForKeDeregisterNmiCallback =
      KeRegisterNmiCallback(&VmxBroadcastHandleNmiCallback, NULL);
#endif BroadcastEnableNmiExitingAllCores();
  g_NmiBroadcastingInitialized = TRUE;
}

VOID VmxBroadcastUninitialize() {
  g_NmiBroadcastingInitialized = FALSE;
#if USE_DEFAULT_OS_IDT_AS_HOST_IDT == \
    TRUE KeDeregisterNmiCallback(g_NmiHandlerForKeDeregisterNmiCallback);
#endif BroadcastDisableNmiExitingAllCores();
  ApicUninitialize();
}

BOOLEAN VmxBroadcastHandleNmiCallback(PVOID Context, BOOLEAN Handled) {
  UNREFERENCED_PARAMETER(Context);
  ULONG CurrentCore;
  CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  VIRTUAL_MACHINE_STATE *VCpu = &g_GuestState[CurrentCore];
  if (!VmxBroadcastNmiHandler(VCpu, TRUE)) {
    return Handled;
  } else {
    return TRUE;
  }
}

BOOLEAN VmxBroadcastNmi(VIRTUAL_MACHINE_STATE *VCpu,
                        NMI_BROADCAST_ACTION_TYPE VmxBroadcastAction) {
  ULONG ProcessorsCount;
  if (!g_NmiBroadcastingInitialized) {
    return FALSE;
  }
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  for (size_t i = 0; i < ProcessorsCount; i++) {
    if (i != VCpu->CoreId) {
      SpinlockInterlockedCompareExchange(
          (volatile LONG *)&g_GuestState[i]
              .NmiBroadcastingState.NmiBroadcastAction,
          VmxBroadcastAction, NMI_BROADCAST_ACTION_NONE);
    }
  }
  ApicTriggerGenericNmi();
  return TRUE;
}

BOOLEAN VmxBroadcastNmiHandler(VIRTUAL_MACHINE_STATE *VCpu,
                               BOOLEAN IsOnVmxNmiHandler) {
  NMI_BROADCAST_ACTION_TYPE BroadcastAction;
  BOOLEAN IsHandled = FALSE;
  BroadcastAction = InterlockedExchange(
      (volatile LONG *)&VCpu->NmiBroadcastingState.NmiBroadcastAction,
      NMI_BROADCAST_ACTION_NONE);
  if (BroadcastAction == NMI_BROADCAST_ACTION_NONE) {
    IsHandled = FALSE;
    goto ReturnIsHandled;
  }
  switch (BroadcastAction) {
    case NMI_BROADCAST_ACTION_TEST:
      IsHandled = TRUE;
      break;
    case NMI_BROADCAST_ACTION_REQUEST:
      VmmCallbackNmiBroadcastRequestHandler(VCpu->CoreId, IsOnVmxNmiHandler);
      IsHandled = TRUE;
      break;
    case NMI_BROADCAST_ACTION_INVALIDATE_EPT_CACHE_SINGLE_CONTEXT:
      EptInveptSingleContext(VCpu->EptPointer.AsUInt);
      IsHandled = TRUE;
      break;
    case NMI_BROADCAST_ACTION_INVALIDATE_EPT_CACHE_ALL_CONTEXTS:
      EptInveptAllContexts();
      IsHandled = TRUE;
      break;
    default:
      IsHandled = FALSE;
      LogError("Err, invalid NMI reason received");
      break;
  }
ReturnIsHandled:
  return IsHandled;
}
