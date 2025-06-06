
#include "pch.h"
extern UINT64 g_OutputSourceTag;
extern LIST_ENTRY g_OutputSources;
extern LIST_ENTRY g_EventTrace;

UINT64 ForwardingGetNewOutputSourceTag() { return g_OutputSourceTag++; }

DEBUGGER_OUTPUT_SOURCE_STATUS ForwardingOpenOutputSource(
    PDEBUGGER_EVENT_FORWARDING SourceDescriptor) {
  if (SourceDescriptor->State == EVENT_FORWARDING_CLOSED) {
    return DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED;
  }
  if (SourceDescriptor->State == EVENT_FORWARDING_STATE_OPENED) {
    return DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_OPENED;
  }
  SourceDescriptor->State = EVENT_FORWARDING_STATE_OPENED;
  if (SourceDescriptor->Type == EVENT_FORWARDING_FILE) {
    return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED;
  } else if (SourceDescriptor->Type == EVENT_FORWARDING_NAMEDPIPE) {
    return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED;
  } else if (SourceDescriptor->Type == EVENT_FORWARDING_TCP) {
    return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED;
  } else if (SourceDescriptor->Type == EVENT_FORWARDING_MODULE) {
    return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED;
  }
  return DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR;
}

DEBUGGER_OUTPUT_SOURCE_STATUS ForwardingCloseOutputSource(
    PDEBUGGER_EVENT_FORWARDING SourceDescriptor) {
  if (SourceDescriptor->State == EVENT_FORWARDING_CLOSED) {
    return DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED;
  }
  if (SourceDescriptor->State == EVENT_FORWARDING_STATE_NOT_OPENED ||
      SourceDescriptor->State != EVENT_FORWARDING_STATE_OPENED) {
    return DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR;
  }
  SourceDescriptor->State = EVENT_FORWARDING_CLOSED;
  if (SourceDescriptor->Type == EVENT_FORWARDING_FILE) {
    CloseHandle(SourceDescriptor->Handle);
    return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED;
  } else if (SourceDescriptor->Type == EVENT_FORWARDING_TCP) {
    CommunicationClientShutdownConnection(SourceDescriptor->Socket);
    CommunicationClientCleanup(SourceDescriptor->Socket);
    return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED;
  } else if (SourceDescriptor->Type == EVENT_FORWARDING_NAMEDPIPE) {
    NamedPipeClientClosePipe(SourceDescriptor->Handle);
    return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED;
  } else if (SourceDescriptor->Type == EVENT_FORWARDING_MODULE) {
    FreeLibrary(SourceDescriptor->Module);
    return DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED;
  }
  return DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR;
}

VOID *ForwardingCreateOutputSource(DEBUGGER_EVENT_FORWARDING_TYPE SourceType,
                                   const string &Description, SOCKET *Socket,
                                   HMODULE *Module) {
  string IpPortDelimiter;
  string Ip;
  string Port;
  if (SourceType == EVENT_FORWARDING_FILE) {
    HANDLE FileHandle = CreateFileA(Description.c_str(), GENERIC_WRITE, 0, NULL,
                                    OPEN_ALWAYS, FILE_ATTRIBUTE_NORMAL, NULL);
    return (void *)FileHandle;
  } else if (SourceType == EVENT_FORWARDING_MODULE) {
    HMODULE ModuleHandle = LoadLibraryA(Description.c_str());
    if (ModuleHandle == NULL) {
      ShowMessages("err, unable to load the module\n");
      return INVALID_HANDLE_VALUE;
    }
    hyperdbg_event_forwarding_t hyperdbg_event_forwarding =
        (hyperdbg_event_forwarding_t)GetProcAddress(
            ModuleHandle, "hyperdbg_event_forwarding");
    if (hyperdbg_event_forwarding == NULL) {
      ShowMessages(
          "err, unable to find the 'hyperdbg_event_forwarding' function\n");
      return INVALID_HANDLE_VALUE;
    }
    *Module = ModuleHandle;
    return (void *)hyperdbg_event_forwarding;
  } else if (SourceType == EVENT_FORWARDING_NAMEDPIPE) {
    HANDLE PipeHandle = NamedPipeClientCreatePipe(Description.c_str());
    if (!PipeHandle) {
      return INVALID_HANDLE_VALUE;
    }
    return (void *)PipeHandle;
  } else if (SourceType == EVENT_FORWARDING_TCP) {
    if (Description.find(':') != std::string::npos) {
      IpPortDelimiter = ':';
      size_t find = Description.find(IpPortDelimiter);
      Ip = Description.substr(0, find);
      Port = Description.substr(find + 1, find + Description.size());
      if (CommunicationClientConnectToServer(Ip.c_str(), Port.c_str(),
                                             Socket) == 0) {
        return (void *)TRUE;
      } else {
        return INVALID_HANDLE_VALUE;
      }
    } else {
      return INVALID_HANDLE_VALUE;
    }
  }
  return INVALID_HANDLE_VALUE;
}

BOOLEAN ForwardingPerformEventForwarding(
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetail, CHAR *Message,
    UINT32 MessageLength) {
  BOOLEAN Result = FALSE;
  PLIST_ENTRY TempList = 0;
  for (size_t i = 0; i < DebuggerOutputSourceMaximumRemoteSourceForSingleEvent;
       i++) {
    if (EventDetail->OutputSourceTags[i] == NULL) {
      return Result;
    }
    TempList = &g_OutputSources;
    while (&g_OutputSources != TempList->Flink) {
      TempList = TempList->Flink;
      PDEBUGGER_EVENT_FORWARDING CurrentOutputSourceDetails = CONTAINING_RECORD(
          TempList, DEBUGGER_EVENT_FORWARDING, OutputSourcesList);
      if (EventDetail->OutputSourceTags[i] ==
          CurrentOutputSourceDetails->OutputUniqueTag) {
        if (CurrentOutputSourceDetails->State ==
            EVENT_FORWARDING_STATE_OPENED) {
          switch (CurrentOutputSourceDetails->Type) {
            case EVENT_FORWARDING_NAMEDPIPE:
              Result = ForwardingSendToNamedPipe(
                  CurrentOutputSourceDetails->Handle, Message, MessageLength);
              break;
            case EVENT_FORWARDING_FILE:
              Result = ForwardingWriteToFile(CurrentOutputSourceDetails->Handle,
                                             Message, MessageLength);
              break;
            case EVENT_FORWARDING_TCP:
              Result = ForwardingSendToTcpSocket(
                  CurrentOutputSourceDetails->Socket, Message, MessageLength);
              break;
            case EVENT_FORWARDING_MODULE:
              ((hyperdbg_event_forwarding_t)CurrentOutputSourceDetails->Handle)(
                  Message, MessageLength);
              Result = TRUE;
              break;
            default:
              break;
          }
        }
        break;
      }
    }
  }
  return FALSE;
}

BOOLEAN ForwardingCheckAndPerformEventForwarding(UINT32 OperationCode,
                                                 CHAR *Message,
                                                 UINT32 MessageLength) {
  PLIST_ENTRY TempList;
  BOOLEAN OutputSourceFound = FALSE;
  TempList = &g_EventTrace;
  while (&g_EventTrace != TempList->Blink) {
    TempList = TempList->Blink;
    PDEBUGGER_GENERAL_EVENT_DETAIL EventDetail = CONTAINING_RECORD(
        TempList, DEBUGGER_GENERAL_EVENT_DETAIL, CommandsEventList);
    if (EventDetail->HasCustomOutput &&
        (UINT32)EventDetail->Tag == OperationCode) {
      OutputSourceFound = TRUE;
      if (!ForwardingPerformEventForwarding(EventDetail, Message,
                                            MessageLength)) {
        ShowMessages(
            "err, there was an error transferring the "
            "message to the remote sources\n");
      }
      break;
    }
  }
  return OutputSourceFound;
}

BOOLEAN ForwardingWriteToFile(HANDLE FileHandle, CHAR *Message,
                              UINT32 MessageLength) {
  DWORD BytesWritten = 0;
  BOOL ErrorFlag = FALSE;
  ErrorFlag =
      WriteFile(FileHandle, Message, MessageLength, &BytesWritten, NULL);
  return TRUE;
  if (ErrorFlag == FALSE) {
    return FALSE;
  } else {
    if (BytesWritten != MessageLength) {
      return FALSE;
    } else {
      return TRUE;
    }
  }
  return FALSE;
}

BOOLEAN ForwardingSendToNamedPipe(HANDLE NamedPipeHandle, CHAR *Message,
                                  UINT32 MessageLength) {
  BOOLEAN SentMessageResult;
  SentMessageResult =
      NamedPipeClientSendMessage(NamedPipeHandle, Message, MessageLength);
  if (!SentMessageResult) {
    return FALSE;
  }
  return TRUE;
}

BOOLEAN ForwardingSendToTcpSocket(SOCKET TcpSocket, CHAR *Message,
                                  UINT32 MessageLength) {
  if (CommunicationClientSendMessage(TcpSocket, Message, MessageLength) != 0) {
    return FALSE;
  }
  return TRUE;
}
