
#pragma once

char *VmxLogMessage;

char *VmxTempMessage;

typedef struct _NOTIFY_RECORD {
  NOTIFY_TYPE Type;

  union {
    PKEVENT Event;
    PIRP PendingIrp;
  } Message;

  KDPC Dpc;
  BOOLEAN CheckVmxRootMessagePool;
} NOTIFY_RECORD, *PNOTIFY_RECORD;

typedef struct _BUFFER_HEADER {
  UINT32 OperationNumber;
  UINT32 BufferLength;
  BOOLEAN Valid;
} BUFFER_HEADER, *PBUFFER_HEADER;

typedef struct _LOG_BUFFER_INFORMATION {
  KSPIN_LOCK BufferLock;
  KSPIN_LOCK BufferLockForNonImmMessage;
  UINT64 BufferForMultipleNonImmediateMessage;
  UINT32 CurrentLengthOfNonImmBuffer;
  UINT64 BufferStartAddress;
  UINT64 BufferEndAddress;
  UINT32 CurrentIndexToSend;
  UINT32 CurrentIndexToWrite;
  UINT64 BufferStartAddressPriority;
  UINT64 BufferEndAddressPriority;
  UINT32 CurrentIndexToSendPriority;
  UINT32 CurrentIndexToWritePriority;
} LOG_BUFFER_INFORMATION, *PLOG_BUFFER_INFORMATION;

LOG_BUFFER_INFORMATION *MessageBufferInformation;

volatile LONG VmxRootLoggingLock;

volatile LONG VmxRootLoggingLockForNonImmBuffers;

NOTIFY_RECORD *g_GlobalNotifyRecord;

MESSAGE_TRACING_CALLBACKS g_MsgTracingCallbacks;
BOOLEAN LogReadBuffer(BOOLEAN IsVmxRoot, PVOID BufferToSaveMessage,
                      UINT32 *ReturnedLength);
VOID LogNotifyUsermodeCallback(PKDPC Dpc, PVOID DeferredContext,
                               PVOID SystemArgument1, PVOID SystemArgument2);
