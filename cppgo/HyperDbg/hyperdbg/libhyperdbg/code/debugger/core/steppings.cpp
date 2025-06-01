
#include "pch.h"
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

BOOLEAN SteppingInstrumentationStepIn() {
  DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;
  if (g_ActiveProcessDebuggingState.IsActive) {
    ShowMessages(
        "the instrumentation step-in is only supported in Debugger Mode\n");
    return FALSE;
  }
  RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN;
  return KdSendStepPacketToDebuggee(RequestFormat);
}

BOOLEAN SteppingInstrumentationStepInForTracking() {
  DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;
  if (g_ActiveProcessDebuggingState.IsActive) {
    ShowMessages(
        "the instrumentation step-in is only supported in Debugger Mode\n");
    return FALSE;
  }
  RequestFormat =
      DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN_FOR_TRACKING;
  return KdSendStepPacketToDebuggee(RequestFormat);
}

BOOLEAN SteppingRegularStepIn() {
  DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;
  RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    return KdSendStepPacketToDebuggee(RequestFormat);
  } else if (g_ActiveProcessDebuggingState.IsActive &&
             g_ActiveProcessDebuggingState.IsPaused) {
    UdSendStepPacketToDebuggee(
        g_ActiveProcessDebuggingState.ProcessDebuggingToken,
        g_ActiveProcessDebuggingState.ThreadId, RequestFormat);
    return TRUE;
  } else {
    return FALSE;
  }
}

BOOLEAN SteppingStepOver() {
  DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;
  RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    return KdSendStepPacketToDebuggee(RequestFormat);
  } else if (g_ActiveProcessDebuggingState.IsActive &&
             g_ActiveProcessDebuggingState.IsPaused) {
    UdSendStepPacketToDebuggee(
        g_ActiveProcessDebuggingState.ProcessDebuggingToken,
        g_ActiveProcessDebuggingState.ThreadId, RequestFormat);
    return TRUE;
  } else {
    return FALSE;
  }
}

BOOLEAN SteppingStepOverForGu(BOOLEAN LastInstruction) {
  DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;
  if (!LastInstruction) {
    RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU;
  } else {
    RequestFormat =
        DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU_LAST_INSTRUCTION;
  }
  if (g_IsSerialConnectedToRemoteDebuggee) {
    return KdSendStepPacketToDebuggee(RequestFormat);
  } else if (g_ActiveProcessDebuggingState.IsActive &&
             g_ActiveProcessDebuggingState.IsPaused) {
    UdSendStepPacketToDebuggee(
        g_ActiveProcessDebuggingState.ProcessDebuggingToken,
        g_ActiveProcessDebuggingState.ThreadId, RequestFormat);
    return TRUE;
  } else {
    return FALSE;
  }
}
