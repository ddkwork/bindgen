
#include "pch.h"

UINT32 NamedPipeConnectingAndTransferringBuffers() { return 0; }

HANDLE NamedPipeServerCreatePipe(LPCSTR PipeName, UINT32 OutputBufferSize,
                                 UINT32 InputBufferSize) {
  HANDLE hPipe;
  hPipe =
      CreateNamedPipeA(PipeName, PIPE_ACCESS_DUPLEX,
                       PIPE_TYPE_MESSAGE | PIPE_READMODE_MESSAGE | PIPE_WAIT,
                       PIPE_UNLIMITED_INSTANCES, OutputBufferSize,
                       InputBufferSize, NMPWAIT_USE_DEFAULT_WAIT, NULL);
  if (INVALID_HANDLE_VALUE == hPipe) {
    printf("err, occurred while creating the pipe (%x)\n", GetLastError());
    return NULL;
  }
  return hPipe;
}

BOOLEAN NamedPipeServerWaitForClientConntection(HANDLE PipeHandle) {
  BOOL bClientConnected = ConnectNamedPipe(PipeHandle, NULL);
  if (FALSE == bClientConnected) {
    printf("err, occurred while connecting to the client (%x)\n",
           GetLastError());
    CloseHandle(PipeHandle);
    return FALSE;
  }
  return TRUE;
}

UINT32 NamedPipeServerReadClientMessage(HANDLE PipeHandle, char *BufferToSave,
                                        int MaximumReadBufferLength) {
  DWORD cbBytes;
  BOOL bResult = ReadFile(PipeHandle, BufferToSave, MaximumReadBufferLength,
                          &cbBytes, NULL);
  if ((!bResult) || (0 == cbBytes)) {
    printf("err, occurred while reading from the client (%x)\n",
           GetLastError());
    CloseHandle(PipeHandle);
    return 0;
  }
  return cbBytes;
}

BOOLEAN NamedPipeServerSendMessageToClient(HANDLE PipeHandle,
                                           char *BufferToSend, int BufferSize) {
  DWORD cbBytes;
  BOOLEAN bResult =
      WriteFile(PipeHandle, BufferToSend, BufferSize, &cbBytes, NULL);
  if ((!bResult) || (BufferSize != cbBytes)) {
    printf("Error occurred while writing to the client (%x)\n", GetLastError());
    CloseHandle(PipeHandle);
    return FALSE;
  }
  return TRUE;
}

VOID NamedPipeServerCloseHandle(HANDLE PipeHandle) { CloseHandle(PipeHandle); }

HANDLE NamedPipeClientCreatePipe(LPCSTR PipeName) {
  HANDLE hPipe;
  hPipe = CreateFileA(PipeName, GENERIC_READ | GENERIC_WRITE, 0, NULL,
                      OPEN_EXISTING, 0, NULL);
  if (INVALID_HANDLE_VALUE == hPipe) {
    printf("err, occurred while connecting to the server (%x)\n",
           GetLastError());
    return NULL;
  } else {
    return hPipe;
  }
}

BOOLEAN NamedPipeClientSendMessage(HANDLE PipeHandle, char *BufferToSend,
                                   int BufferSize) {
  DWORD cbBytes;
  BOOL bResult =
      WriteFile(PipeHandle, BufferToSend, BufferSize, &cbBytes, NULL);
  if ((!bResult) || (BufferSize != cbBytes)) {
    printf("err, occurred while writing to the server (%x)\n", GetLastError());
    CloseHandle(PipeHandle);
    CloseHandle(PipeHandle);
    return FALSE;
  } else {
    return TRUE;
  }
}

UINT32 NamedPipeClientReadMessage(HANDLE PipeHandle, char *BufferToRead,
                                  int MaximumSizeOfBuffer) {
  DWORD cbBytes;
  BOOL bResult =
      ReadFile(PipeHandle, BufferToRead, MaximumSizeOfBuffer, &cbBytes, NULL);
  if ((!bResult) || (0 == cbBytes)) {
    printf("err, occurred while reading from the server (%x)\n",
           GetLastError());
    CloseHandle(PipeHandle);
    return NULL;
  }
  return cbBytes;
}

VOID NamedPipeClientClosePipe(HANDLE PipeHandle) { CloseHandle(PipeHandle); }

int NamedPipeServerExample() {
  HANDLE PipeHandle;
  BOOLEAN SentMessageResult;
  UINT32 ReadBytes;
  const int BufferSize = 1024;
  char BufferToRead[BufferSize] = {0};
  char BufferToSend[BufferSize] = "test message to send from server !!!";
  printf("create name pipe\n");
  PipeHandle = NamedPipeServerCreatePipe("\\\\.\\Pipe\\HyperDbgTests",
                                         BufferSize, BufferSize);
  if (!PipeHandle) {
    return 1;
  }
  printf("success!\n");
  printf("wait for the client connection\n");
  if (!NamedPipeServerWaitForClientConntection(PipeHandle)) {
    return 1;
  }
  printf("client connected\n");
  printf("read client message\n");
  ReadBytes =
      NamedPipeServerReadClientMessage(PipeHandle, BufferToRead, BufferSize);
  if (!ReadBytes) {
    return 1;
  }
  printf("Message from client : %s\n", BufferToRead);
  SentMessageResult = NamedPipeServerSendMessageToClient(
      PipeHandle, BufferToSend, (int)strlen(BufferToSend) + 1);
  if (!SentMessageResult) {
    return 1;
  }
  NamedPipeServerCloseHandle(PipeHandle);
  return 0;
}

int NamedPipeClientExample() {
  HANDLE PipeHandle;
  BOOLEAN SentMessageResult;
  UINT32 ReadBytes;
  const int BufferSize = 1024;
  char Buffer[BufferSize] = "test message to send from client !!!";
  PipeHandle = NamedPipeClientCreatePipe("\\\\.\\Pipe\\HyperDbgTests");
  if (!PipeHandle) {
    return 1;
  }
  SentMessageResult =
      NamedPipeClientSendMessage(PipeHandle, Buffer, (int)strlen(Buffer) + 1);
  if (!SentMessageResult) {
    return 1;
  }
  ReadBytes = NamedPipeClientReadMessage(PipeHandle, Buffer, BufferSize);
  if (!ReadBytes) {
    return 1;
  }
  printf("Server sent the following message: %s\n", Buffer);
  NamedPipeClientClosePipe(PipeHandle);
  return 0;
}
