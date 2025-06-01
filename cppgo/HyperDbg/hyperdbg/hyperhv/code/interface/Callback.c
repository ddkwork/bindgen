
#include "pch.h"

BOOLEAN LogCallbackPrepareAndSendMessageToQueue(UINT32 OperationCode,
                                                BOOLEAN IsImmediateMessage,
                                                BOOLEAN ShowCurrentSystemTime,
                                                BOOLEAN Priority,
                                                const char *Fmt, ...) {
  BOOLEAN Result;
  va_list ArgList;
  if (g_Callbacks.LogCallbackPrepareAndSendMessageToQueueWrapper == NULL) {
    return FALSE;
  }
  va_start(ArgList, Fmt);
  Result = g_Callbacks.LogCallbackPrepareAndSendMessageToQueueWrapper(
      OperationCode, IsImmediateMessage, ShowCurrentSystemTime, Priority, Fmt,
      ArgList);
  va_end(ArgList);
  return Result;
}

BOOLEAN LogCallbackSendMessageToQueue(UINT32 OperationCode,
                                      BOOLEAN IsImmediateMessage,
                                      CHAR *LogMessage, UINT32 BufferLen,
                                      BOOLEAN Priority) {
  if (g_Callbacks.LogCallbackSendMessageToQueue == NULL) {
    return FALSE;
  }
  return g_Callbacks.LogCallbackSendMessageToQueue(
      OperationCode, IsImmediateMessage, LogMessage, BufferLen, Priority);
}

BOOLEAN LogCallbackCheckIfBufferIsFull(BOOLEAN Priority) {
  if (g_Callbacks.LogCallbackCheckIfBufferIsFull == NULL) {
    return FALSE;
  }
  return g_Callbacks.LogCallbackCheckIfBufferIsFull(Priority);
}

BOOLEAN
LogCallbackSendBuffer(_In_ UINT32 OperationCode,
                      _In_reads_bytes_(BufferLength) PVOID Buffer,
                      _In_ UINT32 BufferLength, _In_ BOOLEAN Priority) {
  if (g_Callbacks.LogCallbackSendBuffer == NULL) {
    return FALSE;
  }
  return g_Callbacks.LogCallbackSendBuffer(OperationCode, Buffer, BufferLength,
                                           Priority);
}

VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE VmmCallbackTriggerEvents(
    VMM_EVENT_TYPE_ENUM EventType,
    VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE CallingStage, PVOID Context,
    BOOLEAN *PostEventRequired, GUEST_REGS *Regs) {
  if (g_Callbacks.VmmCallbackTriggerEvents == NULL) {
    return VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_NO_INITIALIZED;
  }
  return g_Callbacks.VmmCallbackTriggerEvents(EventType, CallingStage, Context,
                                              PostEventRequired, Regs);
}

VOID VmmCallbackSetLastError(UINT32 LastError) {
  if (g_Callbacks.VmmCallbackSetLastError == NULL) {
    return;
  }
  g_Callbacks.VmmCallbackSetLastError(LastError);
}

BOOLEAN VmmCallbackVmcallHandler(UINT32 CoreId, UINT64 VmcallNumber,
                                 UINT64 OptionalParam1, UINT64 OptionalParam2,
                                 UINT64 OptionalParam3) {
  if (g_Callbacks.VmmCallbackVmcallHandler == NULL) {
    return FALSE;
  }
  return g_Callbacks.VmmCallbackVmcallHandler(
      CoreId, VmcallNumber, OptionalParam1, OptionalParam2, OptionalParam3);
}

VOID VmmCallbackRegisteredMtfHandler(UINT32 CoreId) {
  if (g_Callbacks.VmmCallbackRegisteredMtfHandler == NULL) {
    return;
  }
  g_Callbacks.VmmCallbackRegisteredMtfHandler(CoreId);
}

VOID VmmCallbackNmiBroadcastRequestHandler(UINT32 CoreId,
                                           BOOLEAN IsOnVmxNmiHandler) {
  if (g_Callbacks.VmmCallbackNmiBroadcastRequestHandler == NULL) {
    return;
  }
  g_Callbacks.VmmCallbackNmiBroadcastRequestHandler(CoreId, IsOnVmxNmiHandler);
}

BOOLEAN VmmCallbackQueryTerminateProtectedResource(
    UINT32 CoreId, PROTECTED_HV_RESOURCES_TYPE ResourceType, PVOID Context,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver) {
  if (g_Callbacks.VmmCallbackQueryTerminateProtectedResource == NULL) {
    return FALSE;
  }
  return g_Callbacks.VmmCallbackQueryTerminateProtectedResource(
      CoreId, ResourceType, Context, PassOver);
}

BOOLEAN VmmCallbackRestoreEptState(UINT32 CoreId) {
  if (g_Callbacks.VmmCallbackRestoreEptState == NULL) {
    return FALSE;
  }
  return g_Callbacks.VmmCallbackRestoreEptState(CoreId);
}

BOOLEAN VmmCallbackUnhandledEptViolation(UINT32 CoreId,
                                         UINT64 ViolationQualification,
                                         UINT64 GuestPhysicalAddr) {
  if (g_Callbacks.VmmCallbackCheckUnhandledEptViolations == NULL) {
    return FALSE;
  }
  return g_Callbacks.VmmCallbackCheckUnhandledEptViolations(
      CoreId, ViolationQualification, GuestPhysicalAddr);
}

BOOLEAN DebuggingCallbackHandleBreakpointException(UINT32 CoreId) {
  if (g_Callbacks.DebuggingCallbackHandleBreakpointException == NULL) {
    return FALSE;
  }
  return g_Callbacks.DebuggingCallbackHandleBreakpointException(CoreId);
}

BOOLEAN DebuggingCallbackHandleDebugBreakpointException(UINT32 CoreId) {
  if (g_Callbacks.DebuggingCallbackHandleDebugBreakpointException == NULL) {
    return FALSE;
  }
  return g_Callbacks.DebuggingCallbackHandleDebugBreakpointException(CoreId);
}

BOOLEAN DebuggingCallbackConditionalPageFaultException(
    UINT32 CoreId, UINT64 Address, UINT32 PageFaultErrorCode) {
  if (g_Callbacks.DebuggingCallbackConditionalPageFaultException == NULL) {
    return FALSE;
  }
  return g_Callbacks.DebuggingCallbackConditionalPageFaultException(
      CoreId, Address, PageFaultErrorCode);
}

VOID InterceptionCallbackTriggerCr3ProcessChange(UINT32 CoreId) {
  if (g_Callbacks.InterceptionCallbackTriggerCr3ProcessChange == NULL) {
    return;
  }
  g_Callbacks.InterceptionCallbackTriggerCr3ProcessChange(CoreId);
}

VOID InterceptionCallbackCr3VmexitsForThreadInterception(UINT32 CoreId,
                                                         CR3_TYPE NewCr3) {
  if (g_Callbacks.AttachingHandleCr3VmexitsForThreadInterception == NULL) {
    return;
  }
  g_Callbacks.AttachingHandleCr3VmexitsForThreadInterception(CoreId, NewCr3);
}
