
#pragma once

typedef struct _PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE {
  UINT64 Msr;
  UINT64 Value;
} PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE,
    *PPROCESSOR_DEBUGGING_MSR_READ_OR_WRITE;

typedef struct _DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE {
  UINT16 CsSel;
} DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE,
    *PDEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE;

typedef struct _DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS {
  BOOLEAN InitialSetProcessChangeEvent;
  BOOLEAN InitialSetThreadChangeEvent;
  BOOLEAN InitialSetByClockInterrupt;
  UINT64 CurrentThreadLocationOnGs;
  BOOLEAN DebugRegisterInterceptionState;
  BOOLEAN InterceptClockInterruptsForThreadChange;
  BOOLEAN IsWatingForMovCr3VmExits;
  BOOLEAN InterceptClockInterruptsForProcessChange;
} DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS,
    *PDEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS;

typedef struct _DEBUGGEE_BP_DESCRIPTOR {
  UINT64 BreakpointId;
  LIST_ENTRY BreakpointsList;
  BOOLEAN Enabled;
  UINT64 Address;
  UINT64 PhysAddress;
  UINT32 Pid;
  UINT32 Tid;
  UINT32 Core;
  UINT16 InstructionLength;
  BYTE PreviousByte;
  BOOLEAN SetRflagsIFBitOnMtf;
  BOOLEAN AvoidReApplyBreakpoint;
  BOOLEAN RemoveAfterHit;
  BOOLEAN CheckForCallbacks;
} DEBUGGEE_BP_DESCRIPTOR, *PDEBUGGEE_BP_DESCRIPTOR;

typedef struct _KD_NMI_STATE {
  volatile BOOLEAN NmiCalledInVmxRootRelatedToHaltDebuggee;
  volatile BOOLEAN WaitingToBeLocked;
} KD_NMI_STATE, *PKD_NMI_STATE;

typedef struct _DEBUGGER_PROCESS_THREAD_INFORMATION {
  union {
    UINT64 asUInt;

    struct {
      UINT32 ProcessId;
      UINT32 ThreadId;
    } Fields;
  };
} DEBUGGER_PROCESS_THREAD_INFORMATION, *PDEBUGGER_PROCESS_THREAD_INFORMATION;

typedef struct _DEBUGGER_TRAP_FLAG_STATE {
  UINT32 NumberOfItems;
  DEBUGGER_PROCESS_THREAD_INFORMATION
  ThreadInformation[MAXIMUM_NUMBER_OF_THREAD_INFORMATION_FOR_TRAPS];
} DEBUGGER_TRAP_FLAG_STATE, *PDEBUGGER_TRAP_FLAG_STATE;

typedef struct _DEBUGGEE_HALTED_CORE_TASK {
  BOOLEAN PerformHaltedTask;
  BOOLEAN LockAgainAfterTask;
  UINT64 TargetTask;
  PVOID Context;
  UINT64 KernelStatus;
} DEBUGGEE_HALTED_CORE_TASK, *PDEBUGGEE_HALTED_CORE_TASK;

typedef struct _DATE_TIME_HOLDER {
  TIME_FIELDS TimeFields;
  CHAR TimeBuffer[14];
  CHAR DateBuffer[12];
} DATE_TIME_HOLDER, *PDATE_TIME_HOLDER;

typedef struct _PROCESSOR_DEBUGGING_STATE {
  volatile LONG Lock;
  volatile BOOLEAN MainDebuggingCore;
  GUEST_REGS *Regs;
  UINT32 CoreId;
  BOOLEAN ShortCircuitingEvent;
  BOOLEAN IgnoreDisasmInNextPacket;
  PROCESSOR_DEBUGGING_MSR_READ_OR_WRITE MsrState;
  DATE_TIME_HOLDER DateTimeHolder;
  PDEBUGGEE_BP_DESCRIPTOR SoftwareBreakpointState;
  DEBUGGEE_INSTRUMENTATION_STEP_IN_TRACE InstrumentationStepInTrace;
  BOOLEAN DoNotNmiNotifyOtherCoresByThisCore;
  BOOLEAN TracingMode;
  DEBUGGEE_PROCESS_OR_THREAD_TRACING_DETAILS ThreadOrProcessTracingDetails;
  KD_NMI_STATE NmiState;
  DEBUGGEE_HALTED_CORE_TASK HaltedCoreTask;
  BOOLEAN BreakStarterCore;
  UINT16 InstructionLengthHint;
  UINT64 HardwareDebugRegisterForStepping;
  UINT64 *ScriptEngineCoreSpecificStackBuffer;
  PKDPC KdDpcObject;
  CHAR KdRecvBuffer[MaxSerialPacketSize];
} PROCESSOR_DEBUGGING_STATE, PPROCESSOR_DEBUGGING_STATE;
