
#include "pch.h"

BOOLEAN DebuggerVmcallHandler(UINT32 CoreId, UINT64 VmcallNumber,
                              UINT64 OptionalParam1, UINT64 OptionalParam2,
                              UINT64 OptionalParam3) {
  UNREFERENCED_PARAMETER(OptionalParam3);
  BOOLEAN Result = FALSE;
  PROCESSOR_DEBUGGING_STATE *DbgState = &g_DbgState[CoreId];
  switch (VmcallNumber) {
    case DEBUGGER_VMCALL_VM_EXIT_HALT_SYSTEM: {
      KdHandleBreakpointAndDebugBreakpoints(
          DbgState, DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER, NULL);
      Result = TRUE;
      break;
    }
    case DEBUGGER_VMCALL_VM_EXIT_HALT_SYSTEM_AS_A_RESULT_OF_TRIGGERING_EVENT: {
      DEBUGGER_TRIGGERED_EVENT_DETAILS *TriggeredEventDetail =
          (DEBUGGER_TRIGGERED_EVENT_DETAILS *)OptionalParam1;
      PGUEST_REGS TempReg = NULL;
      TempReg = DbgState->Regs;
      DbgState->Regs = (GUEST_REGS *)OptionalParam2;
      KdHandleBreakpointAndDebugBreakpoints(
          DbgState, DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED,
          TriggeredEventDetail);
      DbgState->Regs = TempReg;
      Result = TRUE;
      break;
    }
    case DEBUGGER_VMCALL_SIGNAL_DEBUGGER_EXECUTION_FINISHED: {
      KdSendCommandFinishedSignal(CoreId);
      Result = TRUE;
      break;
    }
    case DEBUGGER_VMCALL_SEND_MESSAGES_TO_DEBUGGER: {
      if (OptionalParam1 != NULL64_ZERO && OptionalParam2 != NULL64_ZERO) {
        KdLoggingResponsePacketToDebugger((CHAR *)OptionalParam1,
                                          (UINT32)OptionalParam2,
                                          OPERATION_LOG_INFO_MESSAGE);
      }
      Result = TRUE;
      break;
    }
    case DEBUGGER_VMCALL_SEND_GENERAL_BUFFER_TO_DEBUGGER: {
      PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
      DebuggeeBufferRequest =
          (DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER *)
              OptionalParam1;
      KdResponsePacketToDebugger(
          DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER,
          DebuggeeBufferRequest->RequestedAction,
          (CHAR
               *)((UINT64)DebuggeeBufferRequest +
                  (SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER)),
          DebuggeeBufferRequest->LengthOfBuffer);
      if (DebuggeeBufferRequest->PauseDebuggeeWhenSent) {
        DbgState->IgnoreDisasmInNextPacket = TRUE;
        KdHandleBreakpointAndDebugBreakpoints(
            DbgState, DEBUGGEE_PAUSING_REASON_PAUSE, NULL);
      }
      Result = TRUE;
      break;
    }
    default:
      Result = FALSE;
      LogError("Err, invalid VMCALL in top-level debugger");
      break;
  }
  return Result;
}
