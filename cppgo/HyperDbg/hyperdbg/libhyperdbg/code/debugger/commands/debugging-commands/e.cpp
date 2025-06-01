
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandEditMemoryHelp() {
  ShowMessages(
      "eb !eb ed !ed eq !eq : edits the memory at specific address \n");
  ShowMessages("eb  Byte and ASCII characters\n");
  ShowMessages("ed  Double-word values (4 bytes)\n");
  ShowMessages("eq  Quad-word values (8 bytes). \n");
  ShowMessages(
      "\n If you want to edit physical (address) memory then add '!' "
      "at the start of the command\n");
  ShowMessages(
      "syntax : \teb [Address (hex)] [Contents (hex)] [pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \ted [Address (hex)] [Contents (hex)] [pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \teq [Address (hex)] [Contents (hex)] [pid ProcessId (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : eb fffff8077356f010 90 \n");
  ShowMessages("\t\te.g : eb nt!Kd_DEFAULT_Mask ff ff ff ff \n");
  ShowMessages("\t\te.g : eb nt!Kd_DEFAULT_Mask+10+@rcx ff ff ff ff \n");
  ShowMessages("\t\te.g : eb fffff8077356f010 90 90 90 90 \n");
  ShowMessages("\t\te.g : !eq 100000 9090909090909090\n");
  ShowMessages("\t\te.g : !eq nt!ExAllocatePoolWithTag+55 9090909090909090\n");
  ShowMessages(
      "\t\te.g : !eq 100000 9090909090909090 9090909090909090 "
      "9090909090909090 9090909090909090 9090909090909090\n");
}

BOOLEAN WriteMemoryContent(UINT64 AddressToEdit,
                           DEBUGGER_EDIT_MEMORY_TYPE MemoryType,
                           DEBUGGER_EDIT_MEMORY_BYTE_SIZE ByteSize, UINT32 Pid,
                           UINT32 CountOf64Chunks, UINT64 *BufferToEdit) {
  BOOL Status;
  BOOLEAN StatusReturn = FALSE;
  DEBUGGER_EDIT_MEMORY *FinalBuffer;
  DEBUGGER_EDIT_MEMORY EditMemoryRequest = {0};
  UINT32 FinalSize = 0;
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    AssertShowMessageReturnStmt(
        g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
  }
  EditMemoryRequest.ProcessId = Pid;
  EditMemoryRequest.Address = AddressToEdit;
  EditMemoryRequest.CountOf64Chunks = CountOf64Chunks;
  EditMemoryRequest.MemoryType = MemoryType;
  EditMemoryRequest.ByteSize = ByteSize;
  FinalSize = (CountOf64Chunks * sizeof(UINT64)) + SIZEOF_DEBUGGER_EDIT_MEMORY;
  EditMemoryRequest.FinalStructureSize = FinalSize;
  FinalBuffer = (DEBUGGER_EDIT_MEMORY *)malloc(FinalSize);
  if (!FinalBuffer) {
    ShowMessages("unable to allocate memory\n\n");
    return FALSE;
  }
  ZeroMemory(FinalBuffer, FinalSize);
  memcpy((PVOID)FinalBuffer, &EditMemoryRequest, SIZEOF_DEBUGGER_EDIT_MEMORY);
  memcpy((UINT64 *)((UINT64)FinalBuffer + SIZEOF_DEBUGGER_EDIT_MEMORY),
         BufferToEdit, (CountOf64Chunks * sizeof(UINT64)));
  if (g_IsSerialConnectedToRemoteDebuggee) {
    if (!KdSendEditMemoryPacketToDebuggee(FinalBuffer, FinalSize)) {
      free(FinalBuffer);
      return FALSE;
    }
  } else {
    Status = DeviceIoControl(g_DeviceHandle, IOCTL_DEBUGGER_EDIT_MEMORY,
                             FinalBuffer, FinalSize, FinalBuffer,
                             SIZEOF_DEBUGGER_EDIT_MEMORY, NULL, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      free(FinalBuffer);
      return FALSE;
    }
  }
  if (FinalBuffer->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    free(FinalBuffer);
    return TRUE;
  } else {
    ShowErrorMessage(FinalBuffer->Result);
    free(FinalBuffer);
    return FALSE;
  }
}

BOOLEAN HyperDbgWriteMemory(PVOID DestinationAddress,
                            DEBUGGER_EDIT_MEMORY_TYPE MemoryType,
                            UINT32 ProcessId, PVOID SourceAddress,
                            UINT32 NumberOfBytes) {
  UINT32 RequiredBytes = 0;
  DEBUGGER_EDIT_MEMORY_BYTE_SIZE ByteSize;
  UINT64 *TargetBuffer;
  UINT32 FinalSize = 0;
  BOOLEAN Result = FALSE;
  BYTE *BufferToEdit = (BYTE *)SourceAddress;
  ByteSize = EDIT_BYTE;
  RequiredBytes = NumberOfBytes * sizeof(UINT64);
  TargetBuffer = (UINT64 *)malloc(RequiredBytes);
  if (!TargetBuffer) {
    return FALSE;
  }
  ZeroMemory(TargetBuffer, FinalSize);
  for (size_t i = 0; i < NumberOfBytes; i++) {
    TargetBuffer[i] = BufferToEdit[i];
  }
  Result = WriteMemoryContent((UINT64)DestinationAddress, MemoryType, ByteSize,
                              ProcessId, NumberOfBytes, TargetBuffer);
  free(TargetBuffer);
  return Result;
}

VOID CommandEditMemory(vector<CommandToken> CommandTokens, string Command) {
  UINT64 Address;
  UINT64 *FinalBuffer;
  vector<UINT64> ValuesToEdit;
  DEBUGGER_EDIT_MEMORY_TYPE MemoryType;
  DEBUGGER_EDIT_MEMORY_BYTE_SIZE ByteSize;
  BOOL SetAddress = FALSE;
  BOOL SetValue = FALSE;
  BOOL SetProcId = FALSE;
  BOOL NextIsProcId = FALSE;
  UINT64 Value = 0;
  UINT32 ProcId = 0;
  UINT32 CountOfValues = 0;
  UINT32 FinalSize = 0;
  BOOLEAN IsFirstCommand = TRUE;
  if (g_ActiveProcessDebuggingState.IsActive) {
    ProcId = g_ActiveProcessDebuggingState.ProcessId;
  }
  if (CommandTokens.size() <= 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandEditMemoryHelp();
    return;
  }
  for (auto Section : CommandTokens) {
    if (IsFirstCommand) {
      if (CompareLowerCaseStrings(Section, "!eb")) {
        MemoryType = EDIT_PHYSICAL_MEMORY;
        ByteSize = EDIT_BYTE;
      } else if (CompareLowerCaseStrings(Section, "!ed")) {
        MemoryType = EDIT_PHYSICAL_MEMORY;
        ByteSize = EDIT_DWORD;
      } else if (CompareLowerCaseStrings(Section, "!eq")) {
        MemoryType = EDIT_PHYSICAL_MEMORY;
        ByteSize = EDIT_QWORD;
      } else if (CompareLowerCaseStrings(Section, "eb")) {
        MemoryType = EDIT_VIRTUAL_MEMORY;
        ByteSize = EDIT_BYTE;
      } else if (CompareLowerCaseStrings(Section, "ed")) {
        MemoryType = EDIT_VIRTUAL_MEMORY;
        ByteSize = EDIT_DWORD;
      } else if (CompareLowerCaseStrings(Section, "eq")) {
        MemoryType = EDIT_VIRTUAL_MEMORY;
        ByteSize = EDIT_QWORD;
      } else {
        ShowMessages("unknown error happened !\n\n");
        CommandEditMemoryHelp();
        return;
      }
      IsFirstCommand = FALSE;
      continue;
    }
    if (NextIsProcId) {
      NextIsProcId = FALSE;
      if (!ConvertTokenToUInt32(Section, &ProcId)) {
        ShowMessages("please specify a correct hex process id\n\n");
        CommandEditMemoryHelp();
        return;
      } else {
        continue;
      }
    }
    if (!SetProcId && CompareLowerCaseStrings(Section, "pid")) {
      NextIsProcId = TRUE;
      continue;
    }
    if (!SetAddress) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(Section), &Address)) {
        ShowMessages("err, couldn't resolve error at '%s'\n\n",
                     GetCaseSensitiveStringFromCommandToken(Section).c_str());
        CommandEditMemoryHelp();
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
      if (ByteSize == EDIT_BYTE && TargetVal.size() >= 3) {
        ShowMessages("please specify a byte (hex) value for 'eb' or '!eb'\n\n");
        return;
      }
      if (ByteSize == EDIT_DWORD && TargetVal.size() >= 9) {
        ShowMessages(
            "please specify a dword (hex) value for 'ed' or '!ed'\n\n");
        return;
      }
      if (ByteSize == EDIT_QWORD && TargetVal.size() >= 17) {
        ShowMessages(
            "please specify a qword (hex) value for 'eq' or '!eq'\n\n");
        return;
      }
      if (!ConvertStringToUInt64(TargetVal, &Value)) {
        ShowMessages(
            "please specify a correct hex value to change the memory "
            "content\n\n");
        CommandEditMemoryHelp();
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
  if (!SetAddress) {
    ShowMessages("please specify a correct hex address\n\n");
    CommandEditMemoryHelp();
    return;
  }
  if (!SetValue) {
    ShowMessages(
        "please specify a correct hex value as the content to edit\n\n");
    CommandEditMemoryHelp();
    return;
  }
  if (NextIsProcId) {
    ShowMessages("please specify a correct hex value as the process id\n\n");
    CommandEditMemoryHelp();
    return;
  }
  FinalSize = (CountOfValues * sizeof(UINT64));
  FinalBuffer = (UINT64 *)malloc(FinalSize);
  if (!FinalBuffer) {
    ShowMessages("unable to allocate memory\n\n");
    return;
  }
  ZeroMemory(FinalBuffer, FinalSize);
  std::copy(ValuesToEdit.begin(), ValuesToEdit.end(), FinalBuffer);
  WriteMemoryContent(Address, MemoryType, ByteSize, ProcId, CountOfValues,
                     FinalBuffer);
  free(FinalBuffer);
}
