
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

BOOLEAN HyperDbgReadMemory(UINT64 TargetAddress,
                           DEBUGGER_READ_MEMORY_TYPE MemoryType,
                           DEBUGGER_READ_READING_TYPE ReadingType, UINT32 Pid,
                           UINT32 Size, BOOLEAN GetAddressMode,
                           DEBUGGER_READ_MEMORY_ADDRESS_MODE *AddressMode,
                           BYTE *TargetBufferToStore, UINT32 *ReturnLength) {
  BOOL Status;
  ULONG ReturnedLength;
  DEBUGGER_READ_MEMORY ReadMem = {0};
  UINT32 SizeOfTargetBuffer;
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    AssertShowMessageReturnStmt(
        g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
  }
  ReadMem.Address = TargetAddress;
  ReadMem.Pid = Pid;
  ReadMem.Size = Size;
  ReadMem.MemoryType = MemoryType;
  ReadMem.ReadingType = ReadingType;
  ReadMem.GetAddressMode = GetAddressMode;
  SizeOfTargetBuffer = sizeof(DEBUGGER_READ_MEMORY) + (Size * sizeof(CHAR));
  DEBUGGER_READ_MEMORY *MemReadRequest =
      (DEBUGGER_READ_MEMORY *)malloc(SizeOfTargetBuffer);
  if (MemReadRequest == NULL) {
    return FALSE;
  }
  ZeroMemory(MemReadRequest, SizeOfTargetBuffer);
  memcpy(MemReadRequest, &ReadMem, sizeof(DEBUGGER_READ_MEMORY));
  if (g_IsSerialConnectedToRemoteDebuggee) {
    if (!KdSendReadMemoryPacketToDebuggee(MemReadRequest, SizeOfTargetBuffer)) {
      std::free(MemReadRequest);
      return FALSE;
    }
  } else {
    Status = DeviceIoControl(g_DeviceHandle, IOCTL_DEBUGGER_READ_MEMORY,
                             MemReadRequest, SIZEOF_DEBUGGER_READ_MEMORY,
                             MemReadRequest, SizeOfTargetBuffer,
                             &ReturnedLength, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      std::free(MemReadRequest);
      return FALSE;
    }
  }
  if (MemReadRequest->KernelStatus != DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    ShowErrorMessage(MemReadRequest->KernelStatus);
    std::free(MemReadRequest);
    return FALSE;
  } else {
    if (g_IsSerialConnectedToRemoteDebuggee) {
      *ReturnLength = MemReadRequest->ReturnLength;
    } else {
      ReturnedLength -= SIZEOF_DEBUGGER_READ_MEMORY;
      *ReturnLength = ReturnedLength;
    }
    if (GetAddressMode) {
      *AddressMode = MemReadRequest->AddressMode;
    }
    memcpy(TargetBufferToStore,
           ((unsigned char *)MemReadRequest) + sizeof(DEBUGGER_READ_MEMORY),
           *ReturnLength);
    std::free(MemReadRequest);
    return TRUE;
  }
}

VOID HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_MEMORY_STYLE Style,
                                     UINT64 Address,
                                     DEBUGGER_READ_MEMORY_TYPE MemoryType,
                                     DEBUGGER_READ_READING_TYPE ReadingType,
                                     UINT32 Pid, UINT32 Size,
                                     PDEBUGGER_DT_COMMAND_OPTIONS DtDetails) {
  UINT32 ReturnedLength;
  UCHAR *Buffer;
  DEBUGGER_READ_MEMORY_ADDRESS_MODE AddressMode;
  BOOLEAN CheckForAddressMode = FALSE;
  BOOLEAN Status = FALSE;
  if (Style == DEBUGGER_SHOW_COMMAND_DISASSEMBLE64 ||
      Style == DEBUGGER_SHOW_COMMAND_DISASSEMBLE32) {
    CheckForAddressMode = TRUE;
  } else {
    CheckForAddressMode = FALSE;
  }
  Buffer = (UCHAR *)malloc(Size);
  Status = HyperDbgReadMemory(Address, MemoryType, ReadingType, Pid, Size,
                              CheckForAddressMode, &AddressMode, (BYTE *)Buffer,
                              &ReturnedLength);
  if (!Status) {
    if (Style == DEBUGGER_SHOW_COMMAND_DUMP) {
      ShowMessages(
          "HyperDbg attempted to access an invalid target address: 0x%llx\n"
          "if you are confident that the address is valid, it may be paged out "
          "or not yet available in the current CR3 page table\n"
          "you can use the '.pagein' command to load this page table into "
          "memory and "
          "trigger a page fault (#PF), please refer to the documentation for "
          "further details\n\n",
          Address);
    }
    std::free(Buffer);
    return;
  }
  switch (Style) {
    case DEBUGGER_SHOW_COMMAND_DT:
      if (Size == ReturnedLength) {
        ScriptEngineShowDataBasedOnSymbolTypesWrapper(
            DtDetails->TypeName, Address, FALSE, Buffer,
            DtDetails->AdditionalParameters);
      } else if (ReturnedLength == 0) {
        ShowMessages("err, invalid address");
      } else {
        ShowMessages(
            "err, invalid address or memory is smaller than the structure "
            "size");
      }
      break;
    case DEBUGGER_SHOW_COMMAND_DB:
      ShowMemoryCommandDB(Buffer, Size, Address, MemoryType, ReturnedLength);
      break;
    case DEBUGGER_SHOW_COMMAND_DC:
      ShowMemoryCommandDC(Buffer, Size, Address, MemoryType, ReturnedLength);
      break;
    case DEBUGGER_SHOW_COMMAND_DD:
      ShowMemoryCommandDD(Buffer, Size, Address, MemoryType, ReturnedLength);
      break;
    case DEBUGGER_SHOW_COMMAND_DQ:
      ShowMemoryCommandDQ(Buffer, Size, Address, MemoryType, ReturnedLength);
      break;
    case DEBUGGER_SHOW_COMMAND_DUMP:
      CommandDumpSaveIntoFile(Buffer, Size);
      break;
    case DEBUGGER_SHOW_COMMAND_DISASSEMBLE64:
      if (AddressMode == DEBUGGER_READ_ADDRESS_MODE_32_BIT &&
          MemoryType == DEBUGGER_READ_VIRTUAL_ADDRESS) {
        ShowMessages(
            "the target address seems to be located in a 32-bit program, if "
            "so, "
            "please consider using the 'u32' instead to utilize the 32-bit "
            "disassembler\n");
      }
      if (ReturnedLength != 0) {
        HyperDbgDisassembler64(Buffer, Address, ReturnedLength, 0, FALSE, NULL);
      } else {
        ShowMessages("err, invalid address\n");
      }
      break;
    case DEBUGGER_SHOW_COMMAND_DISASSEMBLE32:
      if (AddressMode == DEBUGGER_READ_ADDRESS_MODE_64_BIT &&
          MemoryType == DEBUGGER_READ_VIRTUAL_ADDRESS) {
        ShowMessages(
            "the target address seems to be located in a 64-bit program, if "
            "so, "
            "please consider using the 'u' instead to utilize the 64-bit "
            "disassembler\n");
      }
      if (ReturnedLength != 0) {
        HyperDbgDisassembler32(Buffer, Address, ReturnedLength, 0, FALSE, NULL);
      } else {
        ShowMessages("err, invalid address\n");
      }
      break;
  }
  std::free(Buffer);
}

void ShowMemoryCommandDB(unsigned char *OutputBuffer, UINT32 Size,
                         UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType,
                         UINT64 Length) {
  unsigned int Character;
  for (UINT32 i = 0; i < Size; i += 16) {
    if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS) {
      ShowMessages("#\t");
    }
    ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
    for (size_t j = 0; j < 16; j++) {
      if (i + j >= Length) {
        ShowMessages("?? ");
      } else {
        ShowMessages("%02X ", OutputBuffer[i + j]);
      }
    }
    ShowMessages(" ");
    for (size_t j = 0; j < 16; j++) {
      Character = (OutputBuffer[i + j]);
      if (isprint(Character)) {
        ShowMessages("%c", Character);
      } else {
        ShowMessages(".");
      }
    }
    ShowMessages("\n");
  }
}

void ShowMemoryCommandDC(unsigned char *OutputBuffer, UINT32 Size,
                         UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType,
                         UINT64 Length) {
  unsigned int Character;
  for (UINT32 i = 0; i < Size; i += 16) {
    if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS) {
      ShowMessages("#\t");
    }
    ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
    for (size_t j = 0; j < 16; j += 4) {
      if (i + j >= Length) {
        ShowMessages("???????? ");
      } else {
        UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
        ShowMessages("%08X ", OutputBufferVar);
      }
    }
    ShowMessages(" ");
    for (size_t j = 0; j < 16; j++) {
      Character = (OutputBuffer[i + j]);
      if (isprint(Character)) {
        ShowMessages("%c", Character);
      } else {
        ShowMessages(".");
      }
    }
    ShowMessages("\n");
  }
}

void ShowMemoryCommandDD(unsigned char *OutputBuffer, UINT32 Size,
                         UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType,
                         UINT64 Length) {
  for (UINT32 i = 0; i < Size; i += 16) {
    if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS) {
      ShowMessages("#\t");
    }
    ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
    for (size_t j = 0; j < 16; j += 4) {
      if (i + j >= Length) {
        ShowMessages("???????? ");
      } else {
        UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
        ShowMessages("%08X ", OutputBufferVar);
      }
    }
    ShowMessages("\n");
  }
}

void ShowMemoryCommandDQ(unsigned char *OutputBuffer, UINT32 Size,
                         UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType,
                         UINT64 Length) {
  for (UINT32 i = 0; i < Size; i += 16) {
    if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS) {
      ShowMessages("#\t");
    }
    ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
    for (size_t j = 0; j < 16; j += 8) {
      if (i + j >= Length) {
        ShowMessages("???????? ");
      } else {
        UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j + 4]);
        ShowMessages("%08X`", OutputBufferVar);
        OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
        ShowMessages("%08X ", OutputBufferVar);
      }
    }
    ShowMessages("\n");
  }
}
