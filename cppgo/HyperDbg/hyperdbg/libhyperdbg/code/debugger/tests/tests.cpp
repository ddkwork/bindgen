
#include "pch.h"
extern TCHAR g_TestLocation[MAX_PATH];

BOOLEAN
SetupTestName(_Inout_updates_bytes_all_(BufferLength) PCHAR TestLocation,
              ULONG BufferLength) {
  HANDLE fileHandle;
  DWORD driverLocLen = 0;
  HMODULE ProcHandle = GetModuleHandle(NULL);
  char *Pos;

  GetModuleFileName(ProcHandle, TestLocation, BufferLength);
  Pos = strrchr(TestLocation, '\\');
  if (Pos != NULL) {
    *Pos = '\0';
  }
  if (FAILED(StringCbCat(TestLocation, BufferLength, "\\" TEST_PROCESS_NAME))) {
    return FALSE;
  }
  if ((fileHandle = CreateFile(TestLocation, GENERIC_READ, 0, NULL,
                               OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, NULL)) ==
      INVALID_HANDLE_VALUE) {
    ShowMessages("%s.exe is not loaded\n", TEST_PROCESS_NAME);
    return FALSE;
  }
  if (fileHandle) {
    CloseHandle(fileHandle);
  }
  return TRUE;
}

BOOLEAN CreateProcessAndOpenPipeConnection(PHANDLE ConnectionPipeHandle,
                                           PHANDLE ThreadHandle,
                                           PHANDLE ProcessHandle) {
  HANDLE PipeHandle;
  BOOLEAN SentMessageResult;
  UINT32 ReadBytes;
  char *BufferToRead;
  char *BufferToSend;
  char HandshakeBuffer[] =
      "Hello, Dear Test Process... Yes, I'm HyperDbg Debugger :)";
  PROCESS_INFORMATION ProcessInfo;
  STARTUPINFO StartupInfo;
  char CmdArgs[] = TEST_PROCESS_NAME " im-hyperdbg";
  PipeHandle = NamedPipeServerCreatePipe(
      "\\\\.\\Pipe\\HyperDbgTests", TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE,
      TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  if (!PipeHandle) {
    return FALSE;
  }
  BufferToRead = (char *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  if (!BufferToRead) {
    return FALSE;
  }
  BufferToSend = (char *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  if (!BufferToSend) {
    free(BufferToSend);
    return FALSE;
  }
  RtlZeroMemory(BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  RtlZeroMemory(BufferToSend, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  strcpy(BufferToSend, HandshakeBuffer);
  ZeroMemory(&StartupInfo, sizeof(StartupInfo));
  StartupInfo.cb = sizeof StartupInfo;
  if (!SetupTestName(g_TestLocation, sizeof(g_TestLocation))) {
    free(BufferToRead);
    free(BufferToSend);
    return FALSE;
  }
  if (CreateProcess(g_TestLocation, CmdArgs, NULL, NULL, FALSE,
                    CREATE_NEW_CONSOLE, NULL, NULL, &StartupInfo,
                    &ProcessInfo)) {
    if (!NamedPipeServerWaitForClientConntection(PipeHandle)) {
      free(BufferToRead);
      free(BufferToSend);
      return FALSE;
    }
    ReadBytes = NamedPipeServerReadClientMessage(
        PipeHandle, BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
    if (!ReadBytes) {
      free(BufferToRead);
      free(BufferToSend);
      return FALSE;
    }
    if (strcmp(BufferToRead, "Hey there, Are you HyperDbg?") == 0) {
      SentMessageResult = NamedPipeServerSendMessageToClient(
          PipeHandle, BufferToSend, (UINT32)strlen(BufferToSend) + 1);
      if (!SentMessageResult) {
        free(BufferToRead);
        free(BufferToSend);
        return FALSE;
      }
      RtlZeroMemory(BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
      ReadBytes = NamedPipeServerReadClientMessage(
          PipeHandle, BufferToRead, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
      if (!ReadBytes) {
        free(BufferToRead);
        free(BufferToSend);
        return FALSE;
      }
      if (strcmp(BufferToRead,
                 "Wow! I miss you... Would you plz send test cases?") == 0) {
        *ConnectionPipeHandle = PipeHandle;
        *ThreadHandle = ProcessInfo.hThread;
        *ProcessHandle = ProcessInfo.hProcess;
        free(BufferToRead);
        free(BufferToSend);
        return TRUE;
      }
      ShowMessages("err, could not handshake with the test process\n");
      free(BufferToRead);
      free(BufferToSend);
      return FALSE;
    } else {
      ShowMessages("err, the process could not be started\n");
      free(BufferToRead);
      free(BufferToSend);
      return FALSE;
    }
  }
  free(BufferToRead);
  free(BufferToSend);
  return FALSE;
}

BOOLEAN OpenHyperDbgTestProcess(PHANDLE ThreadHandle, PHANDLE ProcessHandle,
                                CHAR *Args) {
  PROCESS_INFORMATION ProcessInfo = {0};
  STARTUPINFO StartupInfo = {0};
  CHAR CmdArgs[MAX_PATH] = {0};
  StartupInfo.cb = sizeof(StartupInfo);
  if (!SetupTestName(g_TestLocation, sizeof(g_TestLocation))) {
    return FALSE;
  }
  sprintf_s(CmdArgs, sizeof(CmdArgs), "\"%s\" %s", g_TestLocation, Args);
  if (CreateProcessA(NULL, CmdArgs, NULL, NULL, FALSE, CREATE_NEW_CONSOLE, NULL,
                     NULL, &StartupInfo, &ProcessInfo)) {
    *ThreadHandle = ProcessInfo.hThread;
    *ProcessHandle = ProcessInfo.hProcess;
    return TRUE;
  } else {
    ShowMessages("err, CreateProcess failed with error %x\n", GetLastError());
  }
  return FALSE;
}

VOID CloseProcessAndClosePipeConnection(HANDLE ConnectionPipeHandle,
                                        HANDLE ThreadHandle,
                                        HANDLE ProcessHandle) {
  NamedPipeServerCloseHandle(ConnectionPipeHandle);
  CloseHandle(ThreadHandle);
  CloseHandle(ProcessHandle);
}
