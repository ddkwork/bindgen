
#pragma once

#define DEBUGGER_DEBUG_REGISTER_FOR_STEP_OVER 0

#define DEBUGGER_DEBUG_REGISTER_FOR_THREAD_MANAGEMENT 1

typedef struct _DEBUGGER_CORE_EVENTS {
  LIST_ENTRY HiddenHookReadAndWriteAndExecuteEventsHead;
  LIST_ENTRY HiddenHookReadAndWriteEventsHead;
  LIST_ENTRY HiddenHookReadAndExecuteEventsHead;
  LIST_ENTRY HiddenHookWriteAndExecuteEventsHead;
  LIST_ENTRY HiddenHookReadEventsHead;
  LIST_ENTRY HiddenHookWriteEventsHead;
  LIST_ENTRY HiddenHookExecuteEventsHead;
  LIST_ENTRY EptHook2sExecDetourEventsHead;
  LIST_ENTRY EptHookExecCcEventsHead;
  LIST_ENTRY SyscallHooksEferSyscallEventsHead;
  LIST_ENTRY SyscallHooksEferSysretEventsHead;
  LIST_ENTRY CpuidInstructionExecutionEventsHead;
  LIST_ENTRY RdmsrInstructionExecutionEventsHead;
  LIST_ENTRY WrmsrInstructionExecutionEventsHead;
  LIST_ENTRY ExceptionOccurredEventsHead;
  LIST_ENTRY TscInstructionExecutionEventsHead;
  LIST_ENTRY PmcInstructionExecutionEventsHead;
  LIST_ENTRY InInstructionExecutionEventsHead;
  LIST_ENTRY OutInstructionExecutionEventsHead;
  LIST_ENTRY DebugRegistersAccessedEventsHead;
  LIST_ENTRY ExternalInterruptOccurredEventsHead;
  LIST_ENTRY VmcallInstructionExecutionEventsHead;
  LIST_ENTRY TrapExecutionModeChangedEventsHead;
  LIST_ENTRY TrapExecutionInstructionTraceEventsHead;
  LIST_ENTRY ControlRegister3ModifiedEventsHead;
  LIST_ENTRY ControlRegisterModifiedEventsHead;
} DEBUGGER_CORE_EVENTS, *PDEBUGGER_CORE_EVENTS;

typedef struct _DEBUGGER_EVENT_ACTION {
  UINT64 Tag;
  UINT32 ActionOrderCode;
  LIST_ENTRY ActionsList;
  DEBUGGER_EVENT_ACTION_TYPE_ENUM ActionType;
  BOOLEAN ImmediatelySendTheResults;
  DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
  ScriptConfiguration;
  DEBUGGER_EVENT_REQUEST_BUFFER
  RequestedBuffer;
  UINT32 CustomCodeBufferSize;
  PVOID CustomCodeBufferAddress;
} DEBUGGER_EVENT_ACTION, *PDEBUGGER_EVENT_ACTION;

typedef struct _DEBUGGER_EVENT {
  UINT64 Tag;
  LIST_ENTRY EventsOfSameTypeList;
  VMM_EVENT_TYPE_ENUM EventType;
  BOOLEAN Enabled;
  UINT32 CoreId;
  UINT32
  ProcessId;
  LIST_ENTRY ActionsListHead;
  UINT32 CountOfActions;
  BOOLEAN EnableShortCircuiting;
  VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE EventMode;
  DEBUGGER_EVENT_OPTIONS InitOptions;
  DEBUGGER_EVENT_OPTIONS Options;
  UINT32 ConditionsBufferSize;
  PVOID ConditionBufferAddress;
} DEBUGGER_EVENT, *PDEBUGGER_EVENT;

typedef UINT64 DebuggerCheckForCondition(PGUEST_REGS Regs, PVOID Context);

typedef PVOID DebuggerRunCustomCodeFunc(PVOID PreAllocatedBufferAddress,
                                        PGUEST_REGS Regs, PVOID Context);
UINT64 DebuggerGetRegValueWrapper(PGUEST_REGS GuestRegs, UINT32 RegId);
UINT32 DebuggerGetLastError();
VOID DebuggerSetLastError(UINT32 LastError);
BOOLEAN DebuggerInitialize();
VOID DebuggerUninitialize();
PDEBUGGER_EVENT DebuggerCreateEvent(
    BOOLEAN Enabled, UINT32 CoreId, UINT32 ProcessId,
    VMM_EVENT_TYPE_ENUM EventType, UINT64 Tag, DEBUGGER_EVENT_OPTIONS *Options,
    UINT32 ConditionsBufferSize, PVOID ConditionBuffer,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
PDEBUGGER_EVENT_ACTION DebuggerAddActionToEvent(
    PDEBUGGER_EVENT Event, DEBUGGER_EVENT_ACTION_TYPE_ENUM ActionType,
    BOOLEAN SendTheResultsImmediately,
    PDEBUGGER_EVENT_REQUEST_CUSTOM_CODE InTheCaseOfCustomCode,
    PDEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION InTheCaseOfRunScript,
    PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
    BOOLEAN InputFromVmxRoot);
BOOLEAN DebuggerRegisterEvent(PDEBUGGER_EVENT Event);
VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE DebuggerTriggerEvents(
    VMM_EVENT_TYPE_ENUM EventType,
    VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE CallingStage, PVOID Context,
    BOOLEAN *PostEventRequired, GUEST_REGS *Regs);
PDEBUGGER_EVENT DebuggerGetEventByTag(UINT64 Tag);
BOOLEAN DebuggerRemoveEvent(UINT64 Tag, BOOLEAN PoolManagerAllocatedMemory);
BOOLEAN DebuggerQueryDebuggerStatus();
BOOLEAN DebuggerParseEvent(PDEBUGGER_GENERAL_EVENT_DETAIL EventDetails,
                           PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                           BOOLEAN InputFromVmxRoot);
BOOLEAN DebuggerApplyEvent(PDEBUGGER_EVENT Event,
                           PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                           BOOLEAN InputFromVmxRoot);
BOOLEAN DebuggerParseAction(PDEBUGGER_GENERAL_ACTION Action,
                            PDEBUGGER_EVENT_AND_ACTION_RESULT ResultsToReturn,
                            BOOLEAN InputFromVmxRoot);
BOOLEAN DebuggerParseEventsModification(
    PDEBUGGER_MODIFY_EVENTS DebuggerEventModificationRequest,
    BOOLEAN InputFromVmxRoot, BOOLEAN PoolManagerAllocatedMemory);
BOOLEAN DebuggerTerminateEvent(UINT64 Tag, BOOLEAN InputFromVmxRoot);
UINT32 DebuggerEventListCount(PLIST_ENTRY TargetEventList);
VOID DebuggerClearAllEvents(BOOLEAN InputFromVmxRoot,
                            BOOLEAN PoolManagerAllocatedMemory);
BOOLEAN DebuggerEnableOrDisableAllEvents(BOOLEAN IsEnable);
UINT32 DebuggerEventListCountByCore(PLIST_ENTRY TargetEventList,
                                    UINT32 TargetCore);
UINT32 DebuggerEventListCountByEventType(VMM_EVENT_TYPE_ENUM EventType,
                                         UINT32 TargetCore);
UINT32 DebuggerExceptionEventBitmapMask(UINT32 CoreIndex);
BOOLEAN DebuggerIsTagValid(UINT64 Tag);
BOOLEAN DebuggerEnableEvent(UINT64 Tag);
BOOLEAN DebuggerQueryStateEvent(UINT64 Tag);
BOOLEAN DebuggerDisableEvent(UINT64 Tag);
BOOLEAN DebuggerClearEvent(UINT64 Tag, BOOLEAN InputFromVmxRoot,
                           BOOLEAN PoolManagerAllocatedMemory);
VOID DebuggerPerformActions(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGER_EVENT *Event,
    DEBUGGER_TRIGGERED_EVENT_DETAILS *EventTriggerDetail);
VOID DebuggerPerformBreakToDebugger(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGER_EVENT_ACTION *Action,
    DEBUGGER_TRIGGERED_EVENT_DETAILS *EventTriggerDetail);
BOOLEAN DebuggerPerformRunScript(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGER_EVENT_ACTION *Action,
    DEBUGGEE_SCRIPT_PACKET *ScriptDetails,
    DEBUGGER_TRIGGERED_EVENT_DETAILS *EventTriggerDetail);
VOID DebuggerPerformRunTheCustomCode(
    PROCESSOR_DEBUGGING_STATE *DbgState, DEBUGGER_EVENT_ACTION *Action,
    DEBUGGER_TRIGGERED_EVENT_DETAILS *EventTriggerDetail);
PLIST_ENTRY DebuggerGetEventListByEventType(VMM_EVENT_TYPE_ENUM EventType);
