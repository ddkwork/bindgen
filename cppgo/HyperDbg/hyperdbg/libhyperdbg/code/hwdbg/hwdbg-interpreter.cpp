
#include "pch.h"
extern HWDBG_INSTANCE_INFORMATION g_HwdbgInstanceInfo;
extern BOOLEAN g_HwdbgInstanceInfoIsValid;
extern std::vector<UINT32> g_HwdbgPortConfiguration;

BOOLEAN HwdbgInterpretPacket(PVOID BufferReceived, UINT32 LengthReceived) {
  PHWDBG_INSTANCE_INFORMATION InstanceInfoPacket;
  PUINT32 InstanceInfoPorts;
  DEBUGGER_REMOTE_PACKET *TheActualPacket = NULL;
  BOOLEAN Result = FALSE;
  if (g_HwdbgInstanceInfoIsValid) {
    TheActualPacket =
        (DEBUGGER_REMOTE_PACKET *)(((CHAR *)BufferReceived) +
                                   g_HwdbgInstanceInfo.debuggeeAreaOffset);
  } else {
    TheActualPacket =
        (DEBUGGER_REMOTE_PACKET *)(((CHAR *)BufferReceived) +
                                   DEFAULT_INITIAL_DEBUGGEE_TO_DEBUGGER_OFFSET);
  }
  if (TheActualPacket->Indicator == INDICATOR_OF_HYPERDBG_PACKET) {
    if (TheActualPacket->TypeOfThePacket !=
        DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER_HARDWARE_LEVEL) {
      ShowMessages("err, unknown packet received from the debuggee\n");
      return FALSE;
    }
    switch (TheActualPacket->RequestedActionOfThePacket) {
      case hwdbgResponseSuccessOrErrorMessage:
        Result = TRUE;
        break;
      case hwdbgResponseInstanceInfo:
        Result = TRUE;
        InstanceInfoPacket =
            (HWDBG_INSTANCE_INFORMATION *)(((CHAR *)TheActualPacket) +
                                           sizeof(DEBUGGER_REMOTE_PACKET));
        InstanceInfoPorts = (UINT32 *)(((CHAR *)InstanceInfoPacket) +
                                       sizeof(HWDBG_INSTANCE_INFORMATION));
        RtlCopyMemory(&g_HwdbgInstanceInfo, InstanceInfoPacket,
                      sizeof(HWDBG_INSTANCE_INFORMATION));
        g_HwdbgPortConfiguration.clear();
        g_HwdbgInstanceInfoIsValid = TRUE;
        for (size_t i = 0; i < g_HwdbgInstanceInfo.numberOfPorts; i++) {
          g_HwdbgPortConfiguration.push_back(InstanceInfoPorts[i]);
        }
        ScriptEngineSetHwdbgInstanceInfo(&g_HwdbgInstanceInfo);
        break;
      default:
        Result = FALSE;
        ShowMessages(
            "err, unknown packet request received from the debuggee\n");
        break;
    }
  }
  return Result;
}

std::vector<UINT32> HwdbgParseStringMemoryLine(const std::string &Line) {
  std::vector<UINT32> Values;
  std::stringstream Ss(Line);
  std::string Token;
  std::getline(Ss, Token, ':');
  while (std::getline(Ss, Token, ' ')) {
    if (Token.length() == 8 &&
        std::all_of(Token.begin(), Token.end(), ::isxdigit)) {
      Values.push_back(static_cast<UINT32>(std::stoul(Token, nullptr, 16)));
    }
  }
  return Values;
}

BOOLEAN HwdbgInterpreterFillMemoryFromFile(const TCHAR *FileName,
                                           UINT32 *MemoryBuffer,
                                           size_t BufferSize) {
  std::ifstream File(FileName);
  std::string Line;
  BOOLEAN Result = TRUE;
  size_t Index = 0;
  if (!File.is_open()) {
    ShowMessages("err, unable to open file %s\n", FileName);
    return FALSE;
  }
  while (getline(File, Line)) {
    if (Index >= BufferSize) {
      Result = FALSE;
      ShowMessages(
          "err, buffer overflow, file contains more data than buffer can "
          "hold\n");
      break;
    }
    vector<UINT32> Values = HwdbgParseStringMemoryLine(Line);
    for (UINT32 Value : Values) {
      if (Index < BufferSize) {
        MemoryBuffer[Index++] = Value;
      } else {
        ShowMessages(
            "err, buffer overflow, file contains more data than buffer can "
            "hold\n");
        File.close();
        return FALSE;
      }
    }
  }
  File.close();
  return Result;
}

BOOLEAN HwdbgInterpreterFillFileFromMemory(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, const TCHAR *FileName,
    UINT32 *MemoryBuffer, size_t BufferSize,
    HWDBG_ACTION_ENUMS RequestedAction) {
  std::ofstream File(FileName);
  if (!File.is_open()) {
    printf("err, unable to open file %s\n", FileName);
    return FALSE;
  }
  size_t Address = 0;
  for (size_t I = 0; I < BufferSize / sizeof(UINT32); ++I) {
    File << std::hex << std::setw(8) << std::setfill('0') << MemoryBuffer[I];
    File << " ; +0x" << std::hex << std::setw(1) << std::setfill('0')
         << Address;
    if (I == 0) {
      File << "   | Checksum";
    } else if (I == 1) {
      File << "   | Checksum";
    } else if (I == 2) {
      File << "   | Indicator";
    } else if (I == 3) {
      File << "   | Indicator";
    } else if (I == 4) {
      File << "  | TypeOfThePacket - DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL (0x4)";
    } else if (I == 5) {
      File << "  | RequestedActionOfThePacket - Value" << " (0x" << std::hex
           << std::setw(1) << std::setfill('0') << RequestedAction << ")";
    } else if (I == 6) {
      File << "  | Start of Optional Data";
    }
    File << "\n";
    Address += 4;
  }
  if (g_HwdbgInstanceInfoIsValid) {
    while (Address < InstanceInfo->sharedMemorySize) {
      File << "00000000 ; +0x" << std::hex << std::setw(1) << std::setfill('0')
           << Address;
      Address += 4;
      if (Address < InstanceInfo->sharedMemorySize) {
        File << "\n";
      }
    }
  }
  File.close();
  return TRUE;
}

SIZE_T HwdbgComputeNumberOfFlipFlopsNeeded(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, UINT32 NumberOfStages) {
  SIZE_T NumberOfNeededFlipFlopsInTargetDevice = 0;
  NumberOfNeededFlipFlopsInTargetDevice +=
      (NumberOfStages *
       (g_HwdbgInstanceInfo.maximumNumberOfSupportedGetScriptOperators +
        g_HwdbgInstanceInfo.maximumNumberOfSupportedSetScriptOperators) *
       g_HwdbgInstanceInfo.scriptVariableLength * sizeof(HWDBG_SHORT_SYMBOL) /
       sizeof(UINT64));
  NumberOfNeededFlipFlopsInTargetDevice +=
      (NumberOfStages * g_HwdbgInstanceInfo.scriptVariableLength *
       (sizeof(HWDBG_SHORT_SYMBOL) / sizeof(UINT64)) / 2);
  NumberOfNeededFlipFlopsInTargetDevice +=
      (NumberOfStages *
       g_HwdbgInstanceInfo.numberOfSupportedLocalAndGlobalVariables *
       g_HwdbgInstanceInfo.scriptVariableLength);
  NumberOfNeededFlipFlopsInTargetDevice +=
      (NumberOfStages *
       g_HwdbgInstanceInfo.numberOfSupportedTemporaryVariables *
       g_HwdbgInstanceInfo.scriptVariableLength);
  NumberOfNeededFlipFlopsInTargetDevice +=
      (NumberOfStages *
       Log2Ceil(
           g_HwdbgInstanceInfo.maximumNumberOfStages *
           (g_HwdbgInstanceInfo.maximumNumberOfSupportedGetScriptOperators +
            g_HwdbgInstanceInfo.maximumNumberOfSupportedSetScriptOperators +
            1)) *
       2);
  NumberOfNeededFlipFlopsInTargetDevice += (NumberOfStages);
  NumberOfNeededFlipFlopsInTargetDevice +=
      (NumberOfStages * g_HwdbgInstanceInfo.numberOfPins);
  return NumberOfNeededFlipFlopsInTargetDevice;
}

BOOLEAN HwdbgInterpreterSendPacketAndBufferToHwdbg(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, const TCHAR *FileName,
    DEBUGGER_REMOTE_PACKET_TYPE PacketType, HWDBG_ACTION_ENUMS RequestedAction,
    CHAR *Buffer, UINT32 BufferLength) {
  DEBUGGER_REMOTE_PACKET Packet = {0};
  SIZE_T CommandMaxSize = 0;
  SIZE_T FinalBufferSize = 0;
  if (g_HwdbgInstanceInfoIsValid) {
    CommandMaxSize =
        InstanceInfo->debuggeeAreaOffset - InstanceInfo->debuggerAreaOffset;
  } else {
    CommandMaxSize = DEFAULT_INITIAL_DEBUGGEE_TO_DEBUGGER_OFFSET -
                     DEFAULT_INITIAL_DEBUGGER_TO_DEBUGGEE_OFFSET;
  }
  if (Buffer == NULL) {
    BufferLength = 0;
  }
  FinalBufferSize = sizeof(DEBUGGER_REMOTE_PACKET) + BufferLength;
  if (FinalBufferSize > CommandMaxSize) {
    ShowMessages(
        "err, buffer is above the maximum buffer size that can be sent to "
        "hwdbg (%d > %d)\n",
        FinalBufferSize, CommandMaxSize);
    return FALSE;
  }
  Packet.Indicator = INDICATOR_OF_HYPERDBG_PACKET;
  Packet.TypeOfThePacket = PacketType;
  Packet.RequestedActionOfThePacket =
      (DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION)RequestedAction;
  Packet.Checksum =
      KdComputeDataChecksum((PVOID)((UINT64)&Packet + 1),
                            sizeof(DEBUGGER_REMOTE_PACKET) - sizeof(BYTE));
  if (Buffer != NULL) {
    Packet.Checksum += KdComputeDataChecksum((PVOID)Buffer, BufferLength);
  }
  if (g_HwdbgInstanceInfoIsValid) {
    FinalBufferSize += InstanceInfo->debuggerAreaOffset;
  } else {
    FinalBufferSize += DEFAULT_INITIAL_DEBUGGER_TO_DEBUGGEE_OFFSET;
  }
  CHAR *FinalBuffer = (CHAR *)malloc(FinalBufferSize);
  if (!FinalBuffer) {
    return FALSE;
  }
  RtlZeroMemory(FinalBuffer, FinalBufferSize);
  SIZE_T Offset = g_HwdbgInstanceInfoIsValid
                      ? InstanceInfo->debuggerAreaOffset
                      : DEFAULT_INITIAL_DEBUGGER_TO_DEBUGGEE_OFFSET;
  memcpy(FinalBuffer + Offset, &Packet, sizeof(DEBUGGER_REMOTE_PACKET));
  if (Buffer != NULL) {
    memcpy(FinalBuffer + Offset + sizeof(DEBUGGER_REMOTE_PACKET), Buffer,
           BufferLength);
  }
  HwdbgInterpreterFillFileFromMemory(InstanceInfo, FileName,
                                     (UINT32 *)FinalBuffer, FinalBufferSize,
                                     RequestedAction);
  free(FinalBuffer);
  return TRUE;
}

VOID HwdbgShowIntanceInfo(HWDBG_INSTANCE_INFORMATION *InstanceInfo) {
  UINT32 PortNum = 0;
  ShowMessages("Debuggee Version: 0x%x\n", InstanceInfo->version);
  ShowMessages("Debuggee Maximum Number Of Stages: 0x%x\n",
               InstanceInfo->maximumNumberOfStages);
  ShowMessages("Debuggee Script Variable Length: 0x%x\n",
               InstanceInfo->scriptVariableLength);
  ShowMessages(
      "Debuggee Number of Supported Local (and global) Variables: 0x%x\n",
      InstanceInfo->numberOfSupportedLocalAndGlobalVariables);
  ShowMessages("Debuggee Number of Supported Temporary Variables: 0x%x\n",
               InstanceInfo->numberOfSupportedTemporaryVariables);
  ShowMessages(
      "Debuggee Maximum Number Of Supported GET Script Operators: 0x%x\n",
      InstanceInfo->maximumNumberOfSupportedGetScriptOperators);
  ShowMessages(
      "Debuggee Maximum Number Of Supported SET Script Operators: 0x%x\n",
      InstanceInfo->maximumNumberOfSupportedSetScriptOperators);
  ShowMessages("Debuggee Shared Memory Size: 0x%x\n",
               InstanceInfo->sharedMemorySize);
  ShowMessages("Debuggee Debugger Area Offset: 0x%x\n",
               InstanceInfo->debuggerAreaOffset);
  ShowMessages("Debuggee Debuggee Area Offset: 0x%x\n",
               InstanceInfo->debuggeeAreaOffset);
  ShowMessages("Debuggee Script Capabilities Mask: 0x%llx\n",
               InstanceInfo->scriptCapabilities);
  HardwareScriptInterpreterShowScriptCapabilities(&g_HwdbgInstanceInfo);
  ShowMessages("Debuggee Number Of Pins: 0x%x\n", InstanceInfo->numberOfPins);
  ShowMessages("Debuggee Number Of Ports: 0x%x\n", InstanceInfo->numberOfPorts);
  ShowMessages("Debuggee BRAM Address Width: 0x%x\n",
               InstanceInfo->bramAddrWidth);
  ShowMessages("Debuggee BRAM Data Width: 0x%x (%d bit)\n",
               InstanceInfo->bramDataWidth, InstanceInfo->bramDataWidth);
  for (auto item : g_HwdbgPortConfiguration) {
    ShowMessages("Port number %d ($hw_port%d): 0x%x\n", PortNum, PortNum, item);
    PortNum++;
  }
}

BOOLEAN HwdbgReadInstanceInfoFromFile(const TCHAR *FileName,
                                      UINT32 *MemoryBuffer, size_t BufferSize) {
  TCHAR TestFilePath[MAX_PATH] = {0};
  if (SetupPathForFileName(HWDBG_TEST_READ_INSTANCE_INFO_PATH, TestFilePath,
                           sizeof(TestFilePath), TRUE) &&
      HwdbgInterpreterFillMemoryFromFile(TestFilePath, MemoryBuffer,
                                         BufferSize)) {
    for (SIZE_T I = 0; I < BufferSize; ++I) {
      ShowMessages("%08x ", MemoryBuffer[I]);
      ShowMessages("\n");
    }
    return TRUE;
  }
  return FALSE;
}

BOOLEAN HwdbgWriteTestInstanceInfoRequestIntoFile(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, const CHAR *FileName) {
  TCHAR TestFilePath[MAX_PATH] = {0};
  if (SetupPathForFileName(FileName, TestFilePath, sizeof(TestFilePath),
                           FALSE) &&
      HwdbgInterpreterSendPacketAndBufferToHwdbg(
          InstanceInfo, TestFilePath,
          DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL,
          hwdbgActionSendInstanceInfo, NULL, NULL_ZERO)) {
    ShowMessages("[*] instance info successfully written into file: %s\n",
                 TestFilePath);
    return TRUE;
  }
  return FALSE;
}

BOOLEAN HwdbgLoadInstanceInfo(const TCHAR *InstanceFilePathToRead,
                              UINT32 InitialBramBufferSize) {
  UINT32 *MemoryBuffer = NULL;
  MemoryBuffer = (UINT32 *)malloc(InitialBramBufferSize * sizeof(UINT32));
  if (MemoryBuffer == NULL) {
    ShowMessages(
        "err, unable to allocate memory for the instance info packet of the "
        "debuggee");
    return FALSE;
  }
  if (HwdbgReadInstanceInfoFromFile(InstanceFilePathToRead, MemoryBuffer,
                                    InitialBramBufferSize)) {
    ShowMessages("instance info read successfully\n");
  } else {
    ShowMessages("err, unable to read instance info packet of the debuggee");
    free(MemoryBuffer);
    return FALSE;
  }
  if (HwdbgInterpretPacket(MemoryBuffer, InitialBramBufferSize)) {
    ShowMessages("instance info interpreted successfully\n");
    HwdbgShowIntanceInfo(&g_HwdbgInstanceInfo);
  } else {
    ShowMessages(
        "err, unable to interpret instance info packet of the debuggee");
    free(MemoryBuffer);
    return FALSE;
  }
  free(MemoryBuffer);
  return TRUE;
}
