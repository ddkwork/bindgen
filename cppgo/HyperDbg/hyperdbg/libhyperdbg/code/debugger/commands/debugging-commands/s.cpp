
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandSearchMemoryHelp() {
  ShowMessages(
      "sb !sb sd !sd sq !sq : searches a contiguous memory for a "
      "special byte pattern\n");
  ShowMessages("sb  Byte and ASCII characters\n");
  ShowMessages("sd  Double-word values (4 bytes)\n");
  ShowMessages("sq  Quad-word values (8 bytes). \n");
  ShowMessages(
      "\n If you want to search in physical (address) memory then add '!' "
      "at the start of the command\n");
  ShowMessages(
      "syntax : \tsb [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] "
      "[pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \tsd [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] "
      "[pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \tsq [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] "
      "[pid ProcessId (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : sb nt!ExAllocatePoolWithTag 90 85 95 l ffff \n");
  ShowMessages("\t\te.g : sb nt!ExAllocatePoolWithTag+5 90 85 95 l ffff \n");
  ShowMessages("\t\te.g : sb @rcx+5 90 85 95 l ffff \n");
  ShowMessages("\t\te.g : sb fffff8077356f010 90 85 95 l ffff \n");
  ShowMessages("\t\te.g : sd fffff8077356f010 90423580 l ffff pid 1c0 \n");
  ShowMessages("\t\te.g : !sq 100000 9090909090909090 l ffff\n");
  ShowMessages("\t\te.g : !sq @rdx+r12 9090909090909090 l ffff\n");
  ShowMessages(
      "\t\te.g : !sq 100000 9090909090909090 9090909090909090 "
      "9090909090909090 l ffffff\n");
}

VOID CommandSearchSendRequest(UINT64 *BufferToSendAsIoctl,
                              UINT32 BufferToSendAsIoctlSize) {
  BOOL Status;
  UINT64 CurrentValue;
  PUINT64 ResultsBuffer = NULL;
  ResultsBuffer = (PUINT64)malloc(MaximumSearchResults * sizeof(UINT64));
  ZeroMemory(ResultsBuffer, MaximumSearchResults * sizeof(UINT64));
  Status = DeviceIoControl(g_DeviceHandle, IOCTL_DEBUGGER_SEARCH_MEMORY,
                           BufferToSendAsIoctl, BufferToSendAsIoctlSize,
                           ResultsBuffer, MaximumSearchResults * sizeof(UINT64),
                           NULL, NULL);
  if (!Status) {
    ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
    free(ResultsBuffer);
    return;
  }
  for (size_t i = 0; i < MaximumSearchResults; i++) {
    CurrentValue = ResultsBuffer[i];
    if (CurrentValue == NULL) {
      if (i == 0) {
        ShowMessages("not found\n");
      }
      break;
    }
    ShowMessages("%llx\n", CurrentValue);
  }
  free(ResultsBuffer);
}

VOID CommandSearchMemory(vector<CommandToken> CommandTokens, string Command) {
  UINT64 Address;
  vector<UINT64> ValuesToEdit;
  BOOL SetAddress = FALSE;
  BOOL SetValue = FALSE;
  BOOL SetProcId = FALSE;
  BOOL NextIsProcId = FALSE;
  BOOL SetLength = FALSE;
  BOOL NextIsLength = FALSE;
  DEBUGGER_SEARCH_MEMORY SearchMemoryRequest = {0};
  UINT64 Value = 0;
  UINT64 Length = 0;
  UINT32 ProcId = 0;
  UINT32 CountOfValues = 0;
  UINT32 FinalSize = 0;
  UINT64 *FinalBuffer = NULL;
  BOOLEAN IsFirstCommand = TRUE;
  if (g_ActiveProcessDebuggingState.IsActive) {
    ProcId = g_ActiveProcessDebuggingState.ProcessId;
  }
  if (CommandTokens.size() <= 4) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandSearchMemoryHelp();
    return;
  }
  for (auto Section : CommandTokens) {
    if (IsFirstCommand == TRUE) {
      std::string FirstCommand = GetLowerStringFromCommandToken(Section);
      if (!FirstCommand.compare("!sb")) {
        SearchMemoryRequest.MemoryType = SEARCH_PHYSICAL_MEMORY;
        SearchMemoryRequest.ByteSize = SEARCH_BYTE;
      } else if (!FirstCommand.compare("!sd")) {
        SearchMemoryRequest.MemoryType = SEARCH_PHYSICAL_MEMORY;
        SearchMemoryRequest.ByteSize = SEARCH_DWORD;
      } else if (!FirstCommand.compare("!sq")) {
        SearchMemoryRequest.MemoryType = SEARCH_PHYSICAL_MEMORY;
        SearchMemoryRequest.ByteSize = SEARCH_QWORD;
      } else if (!FirstCommand.compare("sb")) {
        SearchMemoryRequest.MemoryType = SEARCH_VIRTUAL_MEMORY;
        SearchMemoryRequest.ByteSize = SEARCH_BYTE;
      } else if (!FirstCommand.compare("sd")) {
        SearchMemoryRequest.MemoryType = SEARCH_VIRTUAL_MEMORY;
        SearchMemoryRequest.ByteSize = SEARCH_DWORD;
      } else if (!FirstCommand.compare("sq")) {
        SearchMemoryRequest.MemoryType = SEARCH_VIRTUAL_MEMORY;
        SearchMemoryRequest.ByteSize = SEARCH_QWORD;
      } else {
        ShowMessages("unknown error happened!\n\n");
        CommandSearchMemoryHelp();
        return;
      }
      IsFirstCommand = FALSE;
      continue;
    }
    if (NextIsProcId) {
      NextIsProcId = FALSE;
      if (!ConvertTokenToUInt32(Section, &ProcId)) {
        ShowMessages("please specify a correct hex process id\n\n");
        CommandSearchMemoryHelp();
        return;
      } else {
        continue;
      }
    }
    if (NextIsLength) {
      NextIsLength = FALSE;
      if (!ConvertTokenToUInt64(Section, &Length)) {
        ShowMessages("please specify a correct hex length\n\n");
        CommandSearchMemoryHelp();
        return;
      } else {
        SetLength = TRUE;
        continue;
      }
    }
    if (!SetProcId && CompareLowerCaseStrings(Section, "pid")) {
      NextIsProcId = TRUE;
      continue;
    }
    if (!SetLength && CompareLowerCaseStrings(Section, "l")) {
      NextIsLength = TRUE;
      continue;
    }
    if (!SetAddress) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(Section), &Address)) {
        ShowMessages("err, couldn't resolve error at '%s'\n\n",
                     GetCaseSensitiveStringFromCommandToken(Section).c_str());
        CommandSearchMemoryHelp();
        return;
      } else {
        SetAddress = TRUE;
        continue;
      }
    }
    if (SetAddress) {
      std::string TargetVal = GetCaseSensitiveStringFromCommandToken(Section);
      if (TargetVal.rfind("0x", 0) == 0 || TargetVal.rfind("0X", 0) == 0 ||
          TargetVal.rfind("\\x", 0) == 0 || TargetVal.rfind("\\X", 0) == 0) {
        TargetVal = TargetVal.erase(0, 2);
      } else if (TargetVal.rfind('x', 0) == 0 || TargetVal.rfind('X', 0) == 0) {
        TargetVal = TargetVal.erase(0, 1);
      }
      TargetVal.erase(remove(TargetVal.begin(), TargetVal.end(), '`'),
                      TargetVal.end());
      if (SearchMemoryRequest.ByteSize == SEARCH_BYTE &&
          TargetVal.size() >= 3) {
        ShowMessages("please specify a byte (hex) value for 'sb' or '!sb'\n\n");
        return;
      }
      if (SearchMemoryRequest.ByteSize == SEARCH_DWORD &&
          TargetVal.size() >= 9) {
        ShowMessages(
            "please specify a dword (hex) value for 'sd' or '!sd'\n\n");
        return;
      }
      if (SearchMemoryRequest.ByteSize == SEARCH_QWORD &&
          TargetVal.size() >= 17) {
        ShowMessages(
            "please specify a qword (hex) value for 'sq' or '!sq'\n\n");
        return;
      }
      if (!ConvertStringToUInt64(TargetVal, &Value)) {
        ShowMessages(
            "please specify a correct hex value to search in the "
            "memory content\n\n");
        CommandSearchMemoryHelp();
        return;
      } else {
        ValuesToEdit.push_back(Value);
        CountOfValues++;
        if (!SetValue) {
          SetValue = TRUE;
        }
        continue;
      }
    }
  }
  if (g_IsSerialConnectedToRemoteDebuggee && ProcId != 0) {
    ShowMessages(ASSERT_MESSAGE_CANNOT_SPECIFY_PID);
    return;
  }
  if (ProcId == 0) {
    ProcId = GetCurrentProcessId();
  }
  SearchMemoryRequest.ProcessId = ProcId;
  SearchMemoryRequest.Address = Address;
  SearchMemoryRequest.CountOf64Chunks = CountOfValues;
  if (!SetAddress) {
    ShowMessages("please specify a correct hex address\n\n");
    CommandSearchMemoryHelp();
    return;
  }
  if (!SetValue) {
    ShowMessages(
        "please specify a correct hex value as the content to search\n\n");
    CommandSearchMemoryHelp();
    return;
  }
  if (!SetLength) {
    ShowMessages("please specify a correct hex value as the length\n\n");
    CommandSearchMemoryHelp();
    return;
  }
  if (NextIsProcId) {
    ShowMessages("please specify a correct hex value as the process id\n\n");
    CommandSearchMemoryHelp();
    return;
  }
  if (NextIsLength) {
    ShowMessages("please specify a correct hex length\n\n");
    CommandSearchMemoryHelp();
    return;
  }
  FinalSize = (CountOfValues * sizeof(UINT64)) + SIZEOF_DEBUGGER_SEARCH_MEMORY;
  SearchMemoryRequest.FinalStructureSize = FinalSize;
  SearchMemoryRequest.Length = Length;
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    AssertShowMessageReturnStmt(g_DeviceHandle,
                                ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
  }
  FinalBuffer = (UINT64 *)malloc(FinalSize);
  if (!FinalBuffer) {
    ShowMessages("unable to allocate memory\n\n");
    return;
  }
  ZeroMemory(FinalBuffer, FinalSize);
  memcpy(FinalBuffer, &SearchMemoryRequest, SIZEOF_DEBUGGER_SEARCH_MEMORY);
  std::copy(ValuesToEdit.begin(), ValuesToEdit.end(),
            (UINT64 *)((UINT64)FinalBuffer + SIZEOF_DEBUGGER_SEARCH_MEMORY));
  if (g_IsSerialConnectedToRemoteDebuggee) {
    KdSendSearchRequestPacketToDebuggee(FinalBuffer, FinalSize);
  } else {
    CommandSearchSendRequest(FinalBuffer, FinalSize);
  }
  free(FinalBuffer);
}
