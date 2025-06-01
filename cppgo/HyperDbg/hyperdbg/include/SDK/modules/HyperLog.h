
#pragma once

typedef BOOLEAN (*CHECK_VMX_OPERATION)();

typedef BOOLEAN (*CHECK_IMMEDIATE_MESSAGE_SENDING)(UINT32 OperationCode);

typedef BOOLEAN (*SEND_IMMEDIATE_MESSAGE)(CHAR *OptionalBuffer,
                                          UINT32 OptionalBufferLength,
                                          UINT32 OperationCode);

typedef struct _MESSAGE_TRACING_CALLBACKS {
  CHECK_VMX_OPERATION VmxOperationCheck;
  CHECK_IMMEDIATE_MESSAGE_SENDING CheckImmediateMessageSending;
  SEND_IMMEDIATE_MESSAGE SendImmediateMessage;
} MESSAGE_TRACING_CALLBACKS, *PMESSAGE_TRACING_CALLBACKS;
