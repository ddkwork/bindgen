
#include "pch.h"

BOOLEAN inline LogCheckVmxOperation() {
  CHECK_VMX_OPERATION VmxOperationCheck =
      g_MsgTracingCallbacks.VmxOperationCheck;
  if (VmxOperationCheck == NULL) {
    return FALSE;
  }
  return VmxOperationCheck();
}

BOOLEAN inline LogCheckImmediateSend(UINT32 OperationCode) {
  CHECK_IMMEDIATE_MESSAGE_SENDING ImmediateMessageCheck =
      g_MsgTracingCallbacks.CheckImmediateMessageSending;
  if (ImmediateMessageCheck == NULL) {
    return FALSE;
  }
  return ImmediateMessageCheck(OperationCode);
}

BOOLEAN inline LogSendImmediateMessage(CHAR *OptionalBuffer,
                                       UINT32 OptionalBufferLength,
                                       UINT32 OperationCode) {
  SEND_IMMEDIATE_MESSAGE SendImmediateMessage =
      g_MsgTracingCallbacks.SendImmediateMessage;
  if (SendImmediateMessage == NULL) {
    return FALSE;
  }
  return SendImmediateMessage(OptionalBuffer, OptionalBufferLength,
                              OperationCode);
}

BOOLEAN LogInitialize(MESSAGE_TRACING_CALLBACKS *MsgTracingCallbacks) {
  ULONG ProcessorsCount;
  ProcessorsCount = KeQueryActiveProcessorCount(0);
  MessageBufferInformation =
      PlatformMemAllocateZeroedNonPagedPool(sizeof(LOG_BUFFER_INFORMATION) * 2);
  if (!MessageBufferInformation) {
    return FALSE;
  }
  VmxTempMessage = NULL;
  VmxTempMessage =
      PlatformMemAllocateZeroedNonPagedPool(PacketChunkSize * ProcessorsCount);
  if (!VmxTempMessage) {
    PlatformMemFreePool(MessageBufferInformation);
    MessageBufferInformation = NULL;
    return FALSE;
  }
  VmxLogMessage = NULL;
  VmxLogMessage =
      PlatformMemAllocateZeroedNonPagedPool(PacketChunkSize * ProcessorsCount);
  if (!VmxLogMessage) {
    PlatformMemFreePool(MessageBufferInformation);
    MessageBufferInformation = NULL;
    PlatformMemFreePool(VmxTempMessage);
    VmxTempMessage = NULL;
    return FALSE;
  }
  VmxRootLoggingLock = 0;
  for (int i = 0; i < 2; i++) {
    KeInitializeSpinLock(&MessageBufferInformation[i].BufferLock);
    KeInitializeSpinLock(
        &MessageBufferInformation[i].BufferLockForNonImmMessage);
    MessageBufferInformation[i].BufferStartAddress =
        (UINT64)PlatformMemAllocateNonPagedPool(LogBufferSize);
    MessageBufferInformation[i].BufferForMultipleNonImmediateMessage =
        (UINT64)PlatformMemAllocateNonPagedPool(PacketChunkSize);
    if (!MessageBufferInformation[i].BufferStartAddress ||
        !MessageBufferInformation[i].BufferForMultipleNonImmediateMessage) {
      return FALSE;
    }
    MessageBufferInformation[i].BufferStartAddressPriority =
        (UINT64)PlatformMemAllocateNonPagedPool(LogBufferSizePriority);
    if (!MessageBufferInformation[i].BufferStartAddressPriority) {
      return FALSE;
    }
    RtlZeroMemory((void *)MessageBufferInformation[i].BufferStartAddress,
                  LogBufferSize);
    RtlZeroMemory((void *)MessageBufferInformation[i]
                      .BufferForMultipleNonImmediateMessage,
                  PacketChunkSize);
    RtlZeroMemory(
        (void *)MessageBufferInformation[i].BufferStartAddressPriority,
        LogBufferSizePriority);
    MessageBufferInformation[i].BufferEndAddress =
        (UINT64)MessageBufferInformation[i].BufferStartAddress + LogBufferSize;
    MessageBufferInformation[i].BufferEndAddressPriority =
        (UINT64)MessageBufferInformation[i].BufferStartAddressPriority +
        LogBufferSizePriority;
  }
  RtlCopyBytes(&g_MsgTracingCallbacks, MsgTracingCallbacks,
               sizeof(MESSAGE_TRACING_CALLBACKS));
  return TRUE;
}

VOID LogUnInitialize() {
  for (int i = 0; i < 2; i++) {
    if (MessageBufferInformation[i].BufferStartAddress != NULL64_ZERO) {
      PlatformMemFreePool(
          (PVOID)MessageBufferInformation[i].BufferStartAddress);
    }
    if (MessageBufferInformation[i].BufferStartAddressPriority != NULL64_ZERO) {
      PlatformMemFreePool(
          (PVOID)MessageBufferInformation[i].BufferStartAddressPriority);
    }
    if (MessageBufferInformation[i].BufferForMultipleNonImmediateMessage !=
        NULL64_ZERO) {
      PlatformMemFreePool((PVOID)MessageBufferInformation[i]
                              .BufferForMultipleNonImmediateMessage);
    }
  }
  PlatformMemFreePool((PVOID)MessageBufferInformation);
  MessageBufferInformation = NULL;
}

BOOLEAN LogCallbackCheckIfBufferIsFull(BOOLEAN Priority) {
  UINT32 Index;
  BOOLEAN IsVmxRoot;
  UINT32 CurrentIndexToWrite = NULL_ZERO;
  UINT32 CurrentIndexToWritePriority = NULL_ZERO;
  IsVmxRoot = LogCheckVmxOperation();
  if (IsVmxRoot) {
    Index = 1;
  } else {
    Index = 0;
  }
  if (Priority) {
    CurrentIndexToWritePriority =
        MessageBufferInformation[Index].CurrentIndexToWritePriority;
    if (MessageBufferInformation[Index].CurrentIndexToWritePriority >
        MaximumPacketsCapacityPriority - 1) {
      CurrentIndexToWritePriority = 0;
    }
  } else {
    CurrentIndexToWrite = MessageBufferInformation[Index].CurrentIndexToWrite;
    if (MessageBufferInformation[Index].CurrentIndexToWrite >
        MaximumPacketsCapacity - 1) {
      CurrentIndexToWrite = 0;
    }
  }
  BUFFER_HEADER *Header;
  if (Priority) {
    Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                                   .BufferStartAddressPriority +
                               (CurrentIndexToWritePriority *
                                (PacketChunkSize + sizeof(BUFFER_HEADER))));
  } else {
    Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                                   .BufferStartAddress +
                               (CurrentIndexToWrite *
                                (PacketChunkSize + sizeof(BUFFER_HEADER))));
  }
  return Header->Valid;
}

_Use_decl_annotations_

    BOOLEAN
    LogCallbackSendBuffer(UINT32 OperationCode, PVOID Buffer,
                          UINT32 BufferLength, BOOLEAN Priority) {
  UINT32 Index;
  BOOLEAN IsVmxRoot;
  KIRQL OldIRQL = NULL_ZERO;
  if (BufferLength > PacketChunkSize - 1 || BufferLength == 0) {
    return FALSE;
  }
  IsVmxRoot = LogCheckVmxOperation();
  if (LogCheckImmediateSend(OperationCode)) {
    if (!IsVmxRoot) {
      OldIRQL = KeRaiseIrqlToDpcLevel();
    }
    LogSendImmediateMessage(Buffer, BufferLength, OperationCode);
    if (!IsVmxRoot) {
      KeLowerIrql(OldIRQL);
    }
    return TRUE;
  }
  if (IsVmxRoot) {
    Index = 1;
    SpinlockLock(&VmxRootLoggingLock);
  } else {
    Index = 0;
    KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
  }
  if (Priority) {
    if (MessageBufferInformation[Index].CurrentIndexToWritePriority >
        MaximumPacketsCapacityPriority - 1) {
      MessageBufferInformation[Index].CurrentIndexToWritePriority = 0;
    }
  } else {
    if (MessageBufferInformation[Index].CurrentIndexToWrite >
        MaximumPacketsCapacity - 1) {
      MessageBufferInformation[Index].CurrentIndexToWrite = 0;
    }
  }
  BUFFER_HEADER *Header;
  if (Priority) {
    Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                                   .BufferStartAddressPriority +
                               (MessageBufferInformation[Index]
                                    .CurrentIndexToWritePriority *
                                (PacketChunkSize + sizeof(BUFFER_HEADER))));
  } else {
    Header =
        (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                              .BufferStartAddress +
                          (MessageBufferInformation[Index].CurrentIndexToWrite *
                           (PacketChunkSize + sizeof(BUFFER_HEADER))));
  }
  Header->OperationNumber = OperationCode;
  Header->BufferLength = BufferLength;
  Header->Valid = TRUE;
  PVOID SavingBuffer;
  if (Priority) {
    SavingBuffer =
        (PVOID)((UINT64)MessageBufferInformation[Index]
                    .BufferStartAddressPriority +
                (MessageBufferInformation[Index].CurrentIndexToWritePriority *
                 (PacketChunkSize + sizeof(BUFFER_HEADER))) +
                sizeof(BUFFER_HEADER));
  } else {
    SavingBuffer =
        (PVOID)((UINT64)MessageBufferInformation[Index].BufferStartAddress +
                (MessageBufferInformation[Index].CurrentIndexToWrite *
                 (PacketChunkSize + sizeof(BUFFER_HEADER))) +
                sizeof(BUFFER_HEADER));
  }
  RtlCopyBytes(SavingBuffer, Buffer, BufferLength);
  if (Priority) {
    MessageBufferInformation[Index].CurrentIndexToWritePriority =
        MessageBufferInformation[Index].CurrentIndexToWritePriority + 1;
  } else {
    MessageBufferInformation[Index].CurrentIndexToWrite =
        MessageBufferInformation[Index].CurrentIndexToWrite + 1;
  }
  if (g_GlobalNotifyRecord != NULL) {
    g_GlobalNotifyRecord->CheckVmxRootMessagePool = IsVmxRoot;
    KeInsertQueueDpc(&g_GlobalNotifyRecord->Dpc, g_GlobalNotifyRecord, NULL);
    g_GlobalNotifyRecord = NULL;
  }
  if (IsVmxRoot) {
    SpinlockUnlock(&VmxRootLoggingLock);
  } else {
    KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
  }
  return TRUE;
}

UINT32 LogMarkAllAsRead(BOOLEAN IsVmxRoot) {
  UINT32 Index;
  UINT32 ResultsOfBuffersSetToRead = 0;
  KIRQL OldIRQL = NULL_ZERO;
  if (IsVmxRoot) {
    Index = 1;
    SpinlockLock(&VmxRootLoggingLock);
  } else {
    Index = 0;
    KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
  }
  for (size_t i = 0; i < MaximumPacketsCapacity; i++) {
    BUFFER_HEADER *Header =
        (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                              .BufferStartAddress +
                          (MessageBufferInformation[Index].CurrentIndexToSend *
                           (PacketChunkSize + sizeof(BUFFER_HEADER))));
    if (!Header->Valid) {
      if (IsVmxRoot) {
        SpinlockUnlock(&VmxRootLoggingLock);
      } else {
        KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
      }
      return ResultsOfBuffersSetToRead;
    }
    ResultsOfBuffersSetToRead++;
    PVOID SendingBuffer =
        (PVOID)((UINT64)MessageBufferInformation[Index].BufferStartAddress +
                (MessageBufferInformation[Index].CurrentIndexToSend *
                 (PacketChunkSize + sizeof(BUFFER_HEADER))) +
                sizeof(BUFFER_HEADER));
    Header->Valid = FALSE;
    RtlZeroMemory(SendingBuffer, Header->BufferLength);
    if (MessageBufferInformation[Index].CurrentIndexToSend >
        MaximumPacketsCapacity - 2) {
      MessageBufferInformation[Index].CurrentIndexToSend = 0;
    } else {
      MessageBufferInformation[Index].CurrentIndexToSend =
          MessageBufferInformation[Index].CurrentIndexToSend + 1;
    }
  }
  if (IsVmxRoot) {
    SpinlockUnlock(&VmxRootLoggingLock);
  } else {
    KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
  }
  return ResultsOfBuffersSetToRead;
}

BOOLEAN LogReadBuffer(BOOLEAN IsVmxRoot, PVOID BufferToSaveMessage,
                      UINT32 *ReturnedLength) {
  UINT32 Index;
  BOOLEAN PriorityMessageIsAvailable = FALSE;
  KIRQL OldIRQL = NULL_ZERO;
  if (IsVmxRoot) {
    Index = 1;
    SpinlockLock(&VmxRootLoggingLock);
  } else {
    Index = 0;
    KeAcquireSpinLock(&MessageBufferInformation[Index].BufferLock, &OldIRQL);
  }
  BUFFER_HEADER *Header;
  Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                                 .BufferStartAddressPriority +
                             (MessageBufferInformation[Index]
                                  .CurrentIndexToSendPriority *
                              (PacketChunkSize + sizeof(BUFFER_HEADER))));
  if (!Header->Valid) {
    Header =
        (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                              .BufferStartAddress +
                          (MessageBufferInformation[Index].CurrentIndexToSend *
                           (PacketChunkSize + sizeof(BUFFER_HEADER))));
    if (!Header->Valid) {
      if (IsVmxRoot) {
        SpinlockUnlock(&VmxRootLoggingLock);
      } else {
        KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
      }
      return FALSE;
    }
  } else {
    PriorityMessageIsAvailable = TRUE;
  }
  RtlCopyBytes(BufferToSaveMessage, &Header->OperationNumber, sizeof(UINT32));
  PVOID SendingBuffer;
  if (PriorityMessageIsAvailable) {
    SendingBuffer =
        (PVOID)((UINT64)MessageBufferInformation[Index]
                    .BufferStartAddressPriority +
                (MessageBufferInformation[Index].CurrentIndexToSendPriority *
                 (PacketChunkSize + sizeof(BUFFER_HEADER))) +
                sizeof(BUFFER_HEADER));
  } else {
    SendingBuffer =
        (PVOID)((UINT64)MessageBufferInformation[Index].BufferStartAddress +
                (MessageBufferInformation[Index].CurrentIndexToSend *
                 (PacketChunkSize + sizeof(BUFFER_HEADER))) +
                sizeof(BUFFER_HEADER));
  }
  PVOID SavingAddress = (PVOID)((UINT64)BufferToSaveMessage + sizeof(UINT32));
  RtlCopyBytes(SavingAddress, SendingBuffer, Header->BufferLength);
#if ShowMessagesOnDebugger if (Header->OperationNumber <= \
                               OPERATION_LOG_NON_IMMEDIATE_MESSAGE) {
  if (Header->BufferLength > DbgPrintLimitation) {
    for (size_t i = 0; i <= Header->BufferLength / DbgPrintLimitation; i++) {
      if (i != 0) {
        DbgPrint("%s", (char *)((UINT64)SendingBuffer +
                                (DbgPrintLimitation * i) - 2));
      } else {
        DbgPrint("%s",
                 (char *)((UINT64)SendingBuffer + (DbgPrintLimitation * i)));
      }
    }
  } else {
    DbgPrint("%s", (char *)SendingBuffer);
  }
}
#endif
Header->Valid = FALSE;
*ReturnedLength = Header->BufferLength + sizeof(UINT32);
RtlZeroMemory(SendingBuffer, Header->BufferLength);
if (PriorityMessageIsAvailable) {
  if (MessageBufferInformation[Index].CurrentIndexToSendPriority >
      MaximumPacketsCapacityPriority - 2) {
    MessageBufferInformation[Index].CurrentIndexToSendPriority = 0;
  } else {
    MessageBufferInformation[Index].CurrentIndexToSendPriority =
        MessageBufferInformation[Index].CurrentIndexToSendPriority + 1;
  }
} else {
  if (MessageBufferInformation[Index].CurrentIndexToSend >
      MaximumPacketsCapacity - 2) {
    MessageBufferInformation[Index].CurrentIndexToSend = 0;
  } else {
    MessageBufferInformation[Index].CurrentIndexToSend =
        MessageBufferInformation[Index].CurrentIndexToSend + 1;
  }
}
if (IsVmxRoot) {
  SpinlockUnlock(&VmxRootLoggingLock);
} else {
  KeReleaseSpinLock(&MessageBufferInformation[Index].BufferLock, OldIRQL);
}
return TRUE;
}

BOOLEAN LogCheckForNewMessage(BOOLEAN IsVmxRoot, BOOLEAN Priority) {
  UINT32 Index;
  if (IsVmxRoot) {
    Index = 1;
  } else {
    Index = 0;
  }
  BUFFER_HEADER *Header;
  if (Priority) {
    Header = (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                                   .BufferStartAddressPriority +
                               (MessageBufferInformation[Index]
                                    .CurrentIndexToSendPriority *
                                (PacketChunkSize + sizeof(BUFFER_HEADER))));
  } else {
    Header =
        (BUFFER_HEADER *)((UINT64)MessageBufferInformation[Index]
                              .BufferStartAddress +
                          (MessageBufferInformation[Index].CurrentIndexToSend *
                           (PacketChunkSize + sizeof(BUFFER_HEADER))));
  }
  if (!Header->Valid) {
    return FALSE;
  }
  return TRUE;
}

BOOLEAN LogCallbackPrepareAndSendMessageToQueueWrapper(
    UINT32 OperationCode, BOOLEAN IsImmediateMessage,
    BOOLEAN ShowCurrentSystemTime, BOOLEAN Priority, const char *Fmt,
    va_list ArgList) {
  int SprintfResult;
  size_t WrittenSize;
  BOOLEAN IsVmxRootMode;
  BOOLEAN Result = FALSE;
  char *LogMessage = NULL;
  char *TempMessage = NULL;
  char TimeBuffer[20] = {0};
  ULONG CurrentCore = KeGetCurrentProcessorNumberEx(NULL);
  IsVmxRootMode = LogCheckVmxOperation();
  if (IsVmxRootMode) {
    LogMessage = &VmxLogMessage[CurrentCore * PacketChunkSize];
    TempMessage = &VmxTempMessage[CurrentCore * PacketChunkSize];
  } else {
    LogMessage = PlatformMemAllocateNonPagedPool(PacketChunkSize);
    if (LogMessage == NULL) {
      return FALSE;
    }
    TempMessage = PlatformMemAllocateNonPagedPool(PacketChunkSize);
    if (TempMessage == NULL) {
      PlatformMemFreePool(LogMessage);
      return FALSE;
    }
  }
  if (ShowCurrentSystemTime) {
    SprintfResult = vsprintf_s(TempMessage, PacketChunkSize - 1, Fmt, ArgList);
    if (SprintfResult == -1) {
      goto FreeBufferAndReturn;
    }
    TIME_FIELDS TimeFields;
    LARGE_INTEGER SystemTime, LocalTime;
    KeQuerySystemTime(&SystemTime);
    ExSystemTimeToLocalTime(&SystemTime, &LocalTime);
    RtlTimeToTimeFields(&LocalTime, &TimeFields);
    sprintf_s(TimeBuffer, RTL_NUMBER_OF(TimeBuffer), "%02hd:%02hd:%02hd.%03hd",
              TimeFields.Hour, TimeFields.Minute, TimeFields.Second,
              TimeFields.Milliseconds);
    SprintfResult = sprintf_s(
        LogMessage, PacketChunkSize - 1, "(%s - core : %d - vmx-root? %s)\t %s",
        TimeBuffer, CurrentCore, IsVmxRootMode ? "yes" : "no", TempMessage);
    if (SprintfResult == -1) {
      goto FreeBufferAndReturn;
    }
  } else {
    SprintfResult = vsprintf_s(LogMessage, PacketChunkSize - 1, Fmt, ArgList);
    if (SprintfResult == -1) {
      goto FreeBufferAndReturn;
    }
  }
  WrittenSize = strnlen_s(LogMessage, PacketChunkSize - 1);
  if (LogMessage[0] == '\0') {
    goto FreeBufferAndReturn;
  }
  Result =
      LogCallbackSendMessageToQueue(OperationCode, IsImmediateMessage,
                                    LogMessage, (UINT32)WrittenSize, Priority);
FreeBufferAndReturn:
  if (!IsVmxRootMode) {
    PlatformMemFreePool(LogMessage);
    PlatformMemFreePool(TempMessage);
  }
  return Result;
}

BOOLEAN LogCallbackPrepareAndSendMessageToQueue(UINT32 OperationCode,
                                                BOOLEAN IsImmediateMessage,
                                                BOOLEAN ShowCurrentSystemTime,
                                                BOOLEAN Priority,
                                                const char *Fmt, ...) {
  va_list ArgList;
  BOOLEAN Result;
  va_start(ArgList, Fmt);
  Result = LogCallbackPrepareAndSendMessageToQueueWrapper(
      OperationCode, IsImmediateMessage, ShowCurrentSystemTime, Priority, Fmt,
      ArgList);
  va_end(ArgList);
  return Result;
}

BOOLEAN LogCallbackSendMessageToQueue(UINT32 OperationCode,
                                      BOOLEAN IsImmediateMessage,
                                      CHAR *LogMessage, UINT32 BufferLen,
                                      BOOLEAN Priority) {
  BOOLEAN Result;
  UINT32 Index;
  BOOLEAN IsVmxRootMode;
  KIRQL OldIRQL = NULL_ZERO;
  IsVmxRootMode = LogCheckVmxOperation();
#if UseWPPTracing if (OperationCode == OPERATION_LOG_INFO_MESSAGE) {
  HypervisorTraceLevelMessage(TRACE_LEVEL_INFORMATION, HVFS_LOG_INFO, "%s",
                              LogMessage);
}
else if (OperationCode == OPERATION_LOG_WARNING_MESSAGE) {
  HypervisorTraceLevelMessage(TRACE_LEVEL_WARNING, HVFS_LOG_WARNING, "%s",
                              LogMessage);
}
else if (OperationCode == OPERATION_LOG_ERROR_MESSAGE) {
  HypervisorTraceLevelMessage(TRACE_LEVEL_ERROR, HVFS_LOG_ERROR, "%s",
                              LogMessage);
}
else {
  HypervisorTraceLevelMessage(TRACE_LEVEL_NONE, HVFS_LOG, "%s", LogMessage);
}
#else if (IsImmediateMessage) {
  return LogCallbackSendBuffer(OperationCode, LogMessage, BufferLen, Priority);
}
else {
  if (IsVmxRootMode) {
    Index = 1;
    SpinlockLock(&VmxRootLoggingLockForNonImmBuffers);
  } else {
    Index = 0;
    KeAcquireSpinLock(
        &MessageBufferInformation[Index].BufferLockForNonImmMessage, &OldIRQL);
  }
  Result = TRUE;
  if ((MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer +
       BufferLen) > PacketChunkSize - 1 &&
      MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer != 0) {
    Result = LogCallbackSendBuffer(
        OPERATION_LOG_NON_IMMEDIATE_MESSAGE,
        (PVOID)MessageBufferInformation[Index]
            .BufferForMultipleNonImmediateMessage,
        MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer, FALSE);
    MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer = 0;
    RtlZeroMemory((void *)MessageBufferInformation[Index]
                      .BufferForMultipleNonImmediateMessage,
                  PacketChunkSize);
  }
  RtlCopyBytes(
      (void *)(MessageBufferInformation[Index]
                   .BufferForMultipleNonImmediateMessage +
               MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer),
      LogMessage, BufferLen);
  MessageBufferInformation[Index].CurrentLengthOfNonImmBuffer += BufferLen;
  if (IsVmxRootMode) {
    SpinlockUnlock(&VmxRootLoggingLockForNonImmBuffers);
  } else {
    KeReleaseSpinLock(
        &MessageBufferInformation[Index].BufferLockForNonImmMessage, OldIRQL);
  }
  return Result;
}
#endif
}

VOID LogNotifyUsermodeCallback(PKDPC Dpc, PVOID DeferredContext,
                               PVOID SystemArgument1, PVOID SystemArgument2) {
  PNOTIFY_RECORD NotifyRecord;
  PIRP Irp;
  UINT32 Length;
  UNREFERENCED_PARAMETER(Dpc);
  UNREFERENCED_PARAMETER(SystemArgument1);
  UNREFERENCED_PARAMETER(SystemArgument2);
  NotifyRecord = DeferredContext;
  ASSERT(NotifyRecord != NULL);
  _Analysis_assume_(NotifyRecord != NULL);
  switch (NotifyRecord->Type) {
    case IRP_BASED:
      Irp = NotifyRecord->Message.PendingIrp;
      if (Irp != NULL) {
        PCHAR OutBuff;
        ULONG InBuffLength;
        ULONG OutBuffLength;
        PIO_STACK_LOCATION IrpSp;
        if (!(Irp->CurrentLocation <= Irp->StackCount + 1)) {
          DbgPrint(
              "Err, probably two or more functions called DPC for an object");
          return;
        }
        IrpSp = IoGetCurrentIrpStackLocation(Irp);
        InBuffLength = IrpSp->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpSp->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Irp->IoStatus.Status = STATUS_INVALID_PARAMETER;
          IoCompleteRequest(Irp, IO_NO_INCREMENT);
          break;
        }
        if (!Irp->AssociatedIrp.SystemBuffer) {
          return;
        }
        OutBuff = Irp->AssociatedIrp.SystemBuffer;
        Length = 0;
        if (!LogReadBuffer(NotifyRecord->CheckVmxRootMessagePool, OutBuff,
                           &Length)) {
          Irp->IoStatus.Status = STATUS_INVALID_PARAMETER;
          IoCompleteRequest(Irp, IO_NO_INCREMENT);
          break;
        }
        Irp->IoStatus.Information = Length;
        Irp->IoStatus.Status = STATUS_SUCCESS;
        IoCompleteRequest(Irp, IO_NO_INCREMENT);
      }
      break;
    case EVENT_BASED:
      KeSetEvent(NotifyRecord->Message.Event, 0, FALSE);
      ObDereferenceObject(NotifyRecord->Message.Event);
      break;
    default:
      ASSERT(FALSE);
      break;
  }
  if (NotifyRecord != NULL) {
    PlatformMemFreePool(NotifyRecord);
  }
}

BOOLEAN LogRegisterIrpBasedNotification(PVOID TargetIrp, LONG *Status) {
  PNOTIFY_RECORD NotifyRecord;
  PIO_STACK_LOCATION IrpStack;
  PREGISTER_NOTIFY_BUFFER RegisterEvent;
  PIRP Irp = (PIRP)TargetIrp;
  if (g_GlobalNotifyRecord == NULL) {
    IrpStack = IoGetCurrentIrpStackLocation(Irp);
    RegisterEvent = (PREGISTER_NOTIFY_BUFFER)Irp->AssociatedIrp.SystemBuffer;
    NotifyRecord =
        PlatformMemAllocateNonPagedPoolWithQuota(sizeof(NOTIFY_RECORD));
    if (NULL == NotifyRecord) {
      *Status = (LONG)STATUS_INSUFFICIENT_RESOURCES;
      return FALSE;
    }
    NotifyRecord->Type = IRP_BASED;
    NotifyRecord->Message.PendingIrp = Irp;
    KeInitializeDpc(&NotifyRecord->Dpc, LogNotifyUsermodeCallback,
                    NotifyRecord);
    IoMarkIrpPending(Irp);
    if (LogCheckForNewMessage(FALSE, TRUE)) {
      NotifyRecord->CheckVmxRootMessagePool = FALSE;
      KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
    } else if (LogCheckForNewMessage(TRUE, TRUE)) {
      NotifyRecord->CheckVmxRootMessagePool = TRUE;
      KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
    } else if (LogCheckForNewMessage(FALSE, FALSE)) {
      NotifyRecord->CheckVmxRootMessagePool = FALSE;
      KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
    } else if (LogCheckForNewMessage(TRUE, FALSE)) {
      NotifyRecord->CheckVmxRootMessagePool = TRUE;
      KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
    } else {
      g_GlobalNotifyRecord = NotifyRecord;
    }
    *Status = (LONG)STATUS_PENDING;
    return TRUE;
  } else {
    *Status = (LONG)STATUS_SUCCESS;
    return TRUE;
  }
}

BOOLEAN LogRegisterEventBasedNotification(PVOID TargetIrp) {
  PNOTIFY_RECORD NotifyRecord;
  NTSTATUS Status;
  PIO_STACK_LOCATION IrpStack;
  PREGISTER_NOTIFY_BUFFER RegisterEvent;
  PIRP Irp = (PIRP)TargetIrp;
  IrpStack = IoGetCurrentIrpStackLocation(Irp);
  RegisterEvent = (PREGISTER_NOTIFY_BUFFER)Irp->AssociatedIrp.SystemBuffer;
  NotifyRecord =
      PlatformMemAllocateNonPagedPoolWithQuota(sizeof(NOTIFY_RECORD));
  if (NULL == NotifyRecord) {
    DbgPrint("Err, unable to allocate memory for notify record\n");
    return FALSE;
  }
  NotifyRecord->Type = EVENT_BASED;
  KeInitializeDpc(&NotifyRecord->Dpc, LogNotifyUsermodeCallback, NotifyRecord);
  Status = ObReferenceObjectByHandle(RegisterEvent->hEvent,
                                     SYNCHRONIZE | EVENT_MODIFY_STATE,
                                     *ExEventObjectType, Irp->RequestorMode,
                                     &NotifyRecord->Message.Event, NULL);
  if (!NT_SUCCESS(Status)) {
    DbgPrint("Err, unable to reference user mode event object, status = 0x%x\n",
             Status);
    PlatformMemFreePool(NotifyRecord);
    return FALSE;
  }
  KeInsertQueueDpc(&NotifyRecord->Dpc, NotifyRecord, NULL);
  return TRUE;
}
