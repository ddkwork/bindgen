
#pragma once

typedef BOOLEAN (*LOG_CALLBACK_PREPARE_AND_SEND_MESSAGE_TO_QUEUE)(
    UINT32 OperationCode, BOOLEAN IsImmediateMessage,
    BOOLEAN ShowCurrentSystemTime, BOOLEAN Priority, const char *Fmt,
    va_list ArgList);

typedef BOOLEAN (*LOG_CALLBACK_SEND_MESSAGE_TO_QUEUE)(
    UINT32 OperationCode, BOOLEAN IsImmediateMessage, CHAR *LogMessage,
    UINT32 BufferLen, BOOLEAN Priority);

typedef BOOLEAN (*LOG_CALLBACK_SEND_BUFFER)(_In_ UINT32 OperationCode,
                                            _In_reads_bytes_(BufferLength)
                                                PVOID Buffer,
                                            _In_ UINT32 BufferLength,
                                            _In_ BOOLEAN Priority);

typedef BOOLEAN (*LOG_CALLBACK_CHECK_IF_BUFFER_IS_FULL)(BOOLEAN Priority);

typedef VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE (
    *VMM_CALLBACK_TRIGGER_EVENTS)(
    VMM_EVENT_TYPE_ENUM EventType,
    VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE CallingStage, PVOID Context,
    BOOLEAN *PostEventRequired, GUEST_REGS *Regs);

typedef BOOLEAN (*DEBUGGING_CALLBACK_HANDLE_BREAKPOINT_EXCEPTION)(
    UINT32 CoreId);

typedef BOOLEAN (*DEBUGGING_CALLBACK_HANDLE_DEBUG_BREAKPOINT_EXCEPTION)(
    UINT32 CoreId);

typedef BOOLEAN (*DEBUGGING_CALLBACK_CONDITIONAL_PAGE_FAULT_EXCEPTION)(
    UINT32 CoreId, UINT64 Address, UINT32 PageFaultErrorCode);

typedef BOOLEAN (*UD_CHECK_FOR_COMMAND)();

typedef VOID (*VMM_CALLBACK_REGISTERED_MTF_HANDLER)(UINT32 CoreId);

typedef BOOLEAN (*VMM_CALLBACK_RESTORE_EPT_STATE)(UINT32 CoreId);

typedef BOOLEAN (*VMM_CALLBACK_CHECK_UNHANDLED_EPT_VIOLATION)(
    UINT32 CoreId, UINT64 ViolationQualification, UINT64 GuestPhysicalAddr);

typedef VOID (*INTERCEPTION_CALLBACK_TRIGGER_CR3_CHANGE)(UINT32 CoreId);

typedef BOOLEAN (*INTERCEPTION_CALLBACK_TRIGGER_CLOCK_AND_IPI)(
    _In_ UINT32 CoreId);

typedef BOOLEAN (*ATTACHING_HANDLE_CR3_EVENTS_FOR_THREAD_INTERCEPTION)(
    UINT32 CoreId, CR3_TYPE NewCr3);

typedef BOOLEAN (*BREAKPOINT_CHECK_AND_HANDLE_REAPPLYING_BREAKPOINT)(
    UINT32 CoreId);

typedef VOID (*VMM_CALLBACK_NMI_BROADCAST_REQUEST_HANDLER)(
    UINT32 CoreId, BOOLEAN IsOnVmxNmiHandler);

typedef BOOLEAN (*KD_CHECK_AND_HANDLE_NMI_CALLBACK)(UINT32 CoreId);

typedef VOID (*VMM_CALLBACK_SET_LAST_ERROR)(UINT32 LastError);

typedef BOOLEAN (*VMM_CALLBACK_QUERY_TERMINATE_PROTECTED_RESOURCE)(
    UINT32 CoreId, PROTECTED_HV_RESOURCES_TYPE ResourceType, PVOID Context,
    PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver);

typedef BOOLEAN (
    *KD_QUERY_DEBUGGER_THREAD_OR_PROCESS_TRACING_DETAILS_BY_CORE_ID)(
    UINT32 CoreId, DEBUGGER_THREAD_PROCESS_TRACING TracingType);

typedef BOOLEAN (*VMM_CALLBACK_VMCALL_HANDLER)(UINT32 CoreId,
                                               UINT64 VmcallNumber,
                                               UINT64 OptionalParam1,
                                               UINT64 OptionalParam2,
                                               UINT64 OptionalParam3);

typedef struct _VMM_CALLBACKS {
  LOG_CALLBACK_PREPARE_AND_SEND_MESSAGE_TO_QUEUE
  LogCallbackPrepareAndSendMessageToQueueWrapper;
  LOG_CALLBACK_SEND_MESSAGE_TO_QUEUE LogCallbackSendMessageToQueue;
  LOG_CALLBACK_SEND_BUFFER LogCallbackSendBuffer;
  LOG_CALLBACK_CHECK_IF_BUFFER_IS_FULL LogCallbackCheckIfBufferIsFull;
  VMM_CALLBACK_TRIGGER_EVENTS VmmCallbackTriggerEvents;
  VMM_CALLBACK_SET_LAST_ERROR VmmCallbackSetLastError;
  VMM_CALLBACK_VMCALL_HANDLER VmmCallbackVmcallHandler;
  VMM_CALLBACK_NMI_BROADCAST_REQUEST_HANDLER
  VmmCallbackNmiBroadcastRequestHandler;
  VMM_CALLBACK_QUERY_TERMINATE_PROTECTED_RESOURCE
  VmmCallbackQueryTerminateProtectedResource;
  VMM_CALLBACK_RESTORE_EPT_STATE VmmCallbackRestoreEptState;
  VMM_CALLBACK_CHECK_UNHANDLED_EPT_VIOLATION
  VmmCallbackCheckUnhandledEptViolations;
  DEBUGGING_CALLBACK_HANDLE_BREAKPOINT_EXCEPTION
  DebuggingCallbackHandleBreakpointException;
  DEBUGGING_CALLBACK_HANDLE_DEBUG_BREAKPOINT_EXCEPTION
  DebuggingCallbackHandleDebugBreakpointException;
  DEBUGGING_CALLBACK_CONDITIONAL_PAGE_FAULT_EXCEPTION
  DebuggingCallbackConditionalPageFaultException;
  INTERCEPTION_CALLBACK_TRIGGER_CR3_CHANGE
  InterceptionCallbackTriggerCr3ProcessChange;
  BREAKPOINT_CHECK_AND_HANDLE_REAPPLYING_BREAKPOINT
  BreakpointCheckAndHandleReApplyingBreakpoint;
  UD_CHECK_FOR_COMMAND UdCheckForCommand;
  KD_CHECK_AND_HANDLE_NMI_CALLBACK KdCheckAndHandleNmiCallback;
  VMM_CALLBACK_REGISTERED_MTF_HANDLER VmmCallbackRegisteredMtfHandler;
  INTERCEPTION_CALLBACK_TRIGGER_CLOCK_AND_IPI
  DebuggerCheckProcessOrThreadChange;
  ATTACHING_HANDLE_CR3_EVENTS_FOR_THREAD_INTERCEPTION
  AttachingHandleCr3VmexitsForThreadInterception;
  KD_QUERY_DEBUGGER_THREAD_OR_PROCESS_TRACING_DETAILS_BY_CORE_ID
  KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId;
} VMM_CALLBACKS, *PVMM_CALLBACKS;
