
#include "pch.h"

BOOLEAN LoaderInitVmmAndReversingMachine() {
  MESSAGE_TRACING_CALLBACKS MsgTracingCallbacks = {0};
  VMM_CALLBACKS VmmCallbacks = {0};
  g_AllowIOCTLFromUsermode = TRUE;
  MsgTracingCallbacks.VmxOperationCheck = VmFuncVmxGetCurrentExecutionMode;
  VmmCallbacks.LogCallbackPrepareAndSendMessageToQueueWrapper =
      LogCallbackPrepareAndSendMessageToQueueWrapper;
  VmmCallbacks.LogCallbackSendMessageToQueue = LogCallbackSendMessageToQueue;
  VmmCallbacks.LogCallbackSendBuffer = LogCallbackSendBuffer;
  VmmCallbacks.LogCallbackCheckIfBufferIsFull = LogCallbackCheckIfBufferIsFull;
  if (LogInitialize(&MsgTracingCallbacks)) {
    if (VmFuncInitVmm(&VmmCallbacks)) {
      LogDebugInfo("HyperDbg's hypervisor loaded successfully");
      if (CoreInitReversingMachine()) {
        LogDebugInfo("HyperDbg's reversing machine loaded successfully");
        g_HandleInUse = TRUE;
        return TRUE;
      } else {
        LogError("Err, HyperDbg's debugger was not loaded");
      }
    } else {
      LogError("Err, HyperDbg's hypervisor was not loaded");
    }
  } else {
    LogError("Err, HyperDbg's message tracing module was not loaded");
  }
  g_AllowIOCTLFromUsermode = FALSE;
  return FALSE;
}

VOID LoaderUninitializeLogTracer() {
  LogDebugInfo("Unloading HyperDbg's debugger...\n");
#if !UseDbgPrintInsteadOfUsermodeMessageTracking LogDebugInfo( \
    "Uninitializing logs\n");
  LogUnInitialize();
#endif
}
