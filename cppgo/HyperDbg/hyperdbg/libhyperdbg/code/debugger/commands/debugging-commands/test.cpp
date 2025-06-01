
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandTestHelp() {
  ShowMessages(
      "test : tests essential features of HyperDbg in current machine.\n");
  ShowMessages("syntax : \ttest [Task (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : test query\n");
  ShowMessages("\t\te.g : test trap-status\n");
  ShowMessages("\t\te.g : test pool\n");
  ShowMessages("\t\te.g : test query\n");
  ShowMessages("\t\te.g : test breakpoint on\n");
  ShowMessages("\t\te.g : test breakpoint off\n");
  ShowMessages("\t\te.g : test trap on\n");
  ShowMessages("\t\te.g : test trap off\n");
}

BOOLEAN CommandTestPerformKernelTestsIoctl() {
  BOOL Status;
  ULONG ReturnedLength;
  DEBUGGER_PERFORM_KERNEL_TESTS KernelTestRequest = {0};
  AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED,
                              AssertReturnFalse);
  Status = DeviceIoControl(
      g_DeviceHandle, IOCTL_PERFROM_KERNEL_SIDE_TESTS, &KernelTestRequest,
      SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS, &KernelTestRequest,
      SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS, &ReturnedLength, NULL);
  if (!Status) {
    ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
    return FALSE;
  }
  if (KernelTestRequest.KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    return TRUE;
  } else {
    ShowErrorMessage(KernelTestRequest.KernelStatus);
    return FALSE;
  }
}

VOID CommandTestAllFunctionalities() {
  HANDLE ThreadHandle;
  HANDLE ProcessHandle;
  if (!OpenHyperDbgTestProcess(
          &ThreadHandle, &ProcessHandle,
          (CHAR *)TEST_CASE_PARAMETER_FOR_MAIN_COMMAND_PARSER)) {
    ShowMessages(
        "err, start HyperDbg test process for testing the main command "
        "parser\n");
    return;
  }
  if (!OpenHyperDbgTestProcess(
          &ThreadHandle, &ProcessHandle,
          (CHAR *)TEST_CASE_PARAMETER_FOR_SCRIPT_SEMANTIC_TEST_CASES)) {
    ShowMessages(
        "err, start HyperDbg test process for testing semantic tests\n");
    return;
  }
}

VOID CommandTestAllHwdbg() {
  HANDLE ThreadHandle;
  HANDLE ProcessHandle;
  if (!OpenHyperDbgTestProcess(&ThreadHandle, &ProcessHandle,
                               (CHAR *)TEST_HWDBG_FUNCTIONALITIES)) {
    ShowMessages(
        "err, start HyperDbg test process for testing hwdbg functionalities\n");
    return;
  }
}

BOOLEAN CommandTestPerformTest() {
  BOOLEAN ResultOfTest = FALSE;
  HANDLE PipeHandle;
  HANDLE ThreadHandle;
  HANDLE ProcessHandle;
  UINT32 ReadBytes;
  CHAR *Buffer = NULL;
  Buffer = (CHAR *)malloc(TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  if (!Buffer) {
    ShowMessages("err, enable allocate communication buffer\n");
    return FALSE;
  }
  RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  if (!CreateProcessAndOpenPipeConnection(&PipeHandle, &ThreadHandle,
                                          &ProcessHandle)) {
    ShowMessages("err, enable to connect to the test process\n");
    free(Buffer);
    return FALSE;
  }
SendCommandAndWaitForResponse:
  CHAR TestCommand[] = "this is a test command";
  BOOLEAN SentMessageResult = NamedPipeServerSendMessageToClient(
      PipeHandle, TestCommand, (UINT32)strlen(TestCommand) + 1);
  if (!SentMessageResult) {
    return FALSE;
  }
  RtlZeroMemory(Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  ReadBytes = NamedPipeServerReadClientMessage(
      PipeHandle, (char *)Buffer, TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE);
  if (!ReadBytes) {
    free(Buffer);
    return FALSE;
  }
  goto SendCommandAndWaitForResponse;
  CloseProcessAndClosePipeConnection(PipeHandle, ThreadHandle, ProcessHandle);
  free(Buffer);
  return ResultOfTest;
}

VOID CommandTestQueryState() {
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, query state of the debuggee is only possible when you connected "
        "in debugger mode\n");
    return;
  }
  KdSendTestQueryPacketToDebuggee(TEST_QUERY_HALTING_CORE_STATUS);
}

VOID CommandTestQueryTrapState() {
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, query state of the debuggee is only possible when you connected "
        "in debugger mode\n");
    return;
  }
  KdSendTestQueryPacketToDebuggee(TEST_QUERY_TRAP_STATE);
}

VOID CommandTestQueryPreAllocPoolsState() {
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, query state of the debuggee is only possible when you connected "
        "in debugger mode\n");
    return;
  }
  KdSendTestQueryPacketToDebuggee(TEST_QUERY_PREALLOCATED_POOL_STATE);
}

VOID CommandTestSetTargetTaskToHaltedCores(BOOLEAN Synchronous) {
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, query state of the debuggee is only possible when you connected "
        "in debugger mode\n");
    return;
  }
  KdSendTestQueryPacketToDebuggee(
      Synchronous ? TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_SYNCHRONOUS
                  : TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_ASYNCHRONOUS);
}

VOID CommandTestSetTargetTaskToTargetCore(UINT32 CoreNumber) {
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, query state of the debuggee is only possible when you connected "
        "in debugger mode\n");
    return;
  }
  KdSendTestQueryPacketWithContextToDebuggee(
      TEST_SETTING_TARGET_TASKS_ON_TARGET_HALTED_CORES, (UINT64)CoreNumber);
}

VOID CommandTestSetBreakpointState(BOOLEAN State) {
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, query state of the debuggee is only possible when you connected "
        "in debugger mode\n");
    return;
  }
  if (State) {
    KdSendTestQueryPacketToDebuggee(TEST_BREAKPOINT_TURN_ON_BPS);
  } else {
    KdSendTestQueryPacketToDebuggee(TEST_BREAKPOINT_TURN_OFF_BPS);
  }
}

VOID CommandTestSetDebugBreakState(BOOLEAN State) {
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, query state of the debuggee is only possible when you connected "
        "in debugger mode\n");
    return;
  }
  if (State) {
    KdSendTestQueryPacketToDebuggee(TEST_BREAKPOINT_TURN_ON_DBS);
  } else {
    KdSendTestQueryPacketToDebuggee(TEST_BREAKPOINT_TURN_OFF_DBS);
  }
}

VOID CommandTest(vector<CommandToken> CommandTokens, string Command) {
  UINT64 Context = NULL;
  UINT32 CommandSize = (UINT32)CommandTokens.size();
  if (CommandSize == 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandTestHelp();
  } else if (CommandSize == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "query")) {
    CommandTestQueryState();
  } else if (CommandSize == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "trap-status")) {
    CommandTestQueryTrapState();
  } else if (CommandSize == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "pool")) {
    CommandTestQueryPreAllocPoolsState();
  } else if (CommandSize == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "sync-task")) {
    CommandTestSetTargetTaskToHaltedCores(TRUE);
  } else if (CommandSize == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "async-task")) {
    CommandTestSetTargetTaskToHaltedCores(FALSE);
  } else if (CommandSize == 3 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "target-core-task")) {
    if (!ConvertTokenToUInt64(CommandTokens.at(2), &Context)) {
      ShowMessages(
          "err, you should enter a valid hex number as the core id\n\n");
      return;
    }
    CommandTestSetTargetTaskToTargetCore((UINT32)Context);
  } else if (CommandSize == 3 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "breakpoint")) {
    if (CompareLowerCaseStrings(CommandTokens.at(2), "on")) {
      CommandTestSetBreakpointState(TRUE);
    } else if (CompareLowerCaseStrings(CommandTokens.at(2), "off")) {
      CommandTestSetBreakpointState(FALSE);
    } else {
      ShowMessages(
          "err, couldn't resolve error at '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)).c_str());
      return;
    }
  } else if (CommandSize == 3 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "trap")) {
    if (CompareLowerCaseStrings(CommandTokens.at(2), "on")) {
      CommandTestSetDebugBreakState(TRUE);
    } else if (CompareLowerCaseStrings(CommandTokens.at(2), "off")) {
      CommandTestSetDebugBreakState(FALSE);
    } else {
      ShowMessages(
          "err, couldn't resolve error at '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)).c_str());
      return;
    }
  } else if (CommandSize == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "all")) {
    CommandTestAllFunctionalities();
  } else if (CommandSize == 2 &&
             CompareLowerCaseStrings(CommandTokens.at(1), "hwdbg")) {
    CommandTestAllHwdbg();
  } else {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandTestHelp();
    return;
  }
}
