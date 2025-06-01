
#include "pch.h"

BOOLEAN LoaderInitVmmAndDebugger() {
  MESSAGE_TRACING_CALLBACKS MsgTracingCallbacks = {0};
  VMM_CALLBACKS VmmCallbacks = {0};
  g_AllowIOCTLFromUsermode = TRUE;
  MsgTracingCallbacks.VmxOperationCheck = VmFuncVmxGetCurrentExecutionMode;
  MsgTracingCallbacks.CheckImmediateMessageSending =
      KdCheckImmediateMessagingMechanism;
  MsgTracingCallbacks.SendImmediateMessage = KdLoggingResponsePacketToDebugger;
  VmmCallbacks.LogCallbackPrepareAndSendMessageToQueueWrapper =
      LogCallbackPrepareAndSendMessageToQueueWrapper;
  VmmCallbacks.LogCallbackSendMessageToQueue = LogCallbackSendMessageToQueue;
  VmmCallbacks.LogCallbackSendBuffer = LogCallbackSendBuffer;
  VmmCallbacks.LogCallbackCheckIfBufferIsFull = LogCallbackCheckIfBufferIsFull;
  VmmCallbacks.VmmCallbackTriggerEvents = DebuggerTriggerEvents;
  VmmCallbacks.VmmCallbackSetLastError = DebuggerSetLastError;
  VmmCallbacks.VmmCallbackVmcallHandler = DebuggerVmcallHandler;
  VmmCallbacks.VmmCallbackRegisteredMtfHandler = KdHandleRegisteredMtfCallback;
  VmmCallbacks.VmmCallbackNmiBroadcastRequestHandler =
      KdHandleNmiBroadcastDebugBreaks;
  VmmCallbacks.VmmCallbackQueryTerminateProtectedResource =
      TerminateQueryDebuggerResource;
  VmmCallbacks.VmmCallbackRestoreEptState =
      UserAccessCheckForLoadedModuleDetails;
  VmmCallbacks.VmmCallbackCheckUnhandledEptViolations =
      AttachingCheckUnhandledEptViolation;
  VmmCallbacks.DebuggingCallbackHandleBreakpointException =
      BreakpointHandleBreakpoints;
  VmmCallbacks.DebuggingCallbackHandleDebugBreakpointException =
      BreakpointCheckAndHandleDebugBreakpoint;
  VmmCallbacks.BreakpointCheckAndHandleReApplyingBreakpoint =
      BreakpointCheckAndHandleReApplyingBreakpoint;
  VmmCallbacks.UdCheckForCommand = UdCheckForCommand;
  VmmCallbacks.DebuggerCheckProcessOrThreadChange =
      DebuggerCheckProcessOrThreadChange;
  VmmCallbacks.DebuggingCallbackConditionalPageFaultException =
      AttachingCheckPageFaultsWithUserDebugger;
  VmmCallbacks.AttachingHandleCr3VmexitsForThreadInterception =
      AttachingHandleCr3VmexitsForThreadInterception;
  VmmCallbacks.KdCheckAndHandleNmiCallback = KdCheckAndHandleNmiCallback;
  VmmCallbacks.KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId =
      KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId;
  VmmCallbacks.InterceptionCallbackTriggerCr3ProcessChange =
      ProcessTriggerCr3ProcessChange;
  if (LogInitialize(&MsgTracingCallbacks)) {
    if (VmFuncInitVmm(&VmmCallbacks)) {
      LogDebugInfo("HyperDbg's hypervisor loaded successfully");
      if (DebuggerInitialize()) {
        LogDebugInfo("HyperDbg's debugger loaded successfully");
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
