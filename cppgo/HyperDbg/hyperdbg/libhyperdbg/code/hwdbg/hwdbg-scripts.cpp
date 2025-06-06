
#include "pch.h"
extern HWDBG_INSTANCE_INFORMATION g_HwdbgInstanceInfo;
extern BOOLEAN g_HwdbgInstanceInfoIsValid;

VOID HwdbgScriptPrintScriptBuffer(CHAR *ScriptBuffer, UINT32 ScriptBufferSize) {
  ShowMessages(
      "\nHyperDbg (general) script buffer (size=%d, flip-flops (just "
      "script)=%d):\n\n",
      ScriptBufferSize, ScriptBufferSize * 8);
  for (size_t i = 0; i < ScriptBufferSize; i++) {
    ShowMessages("%02X ", (UINT8)ScriptBuffer[i]);
  }
  ShowMessages("\n");
}

BOOLEAN HwdbgScriptCompressScriptBuffer(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, SYMBOL *ScriptBuffer,
    size_t ScriptBufferSize, UINT32 NumberOfStagesForScript,
    HWDBG_SHORT_SYMBOL **NewScriptBuffer, size_t *NewCompressedBufferSize,
    size_t *NumberOfBytesPerChunk) {
  if (!g_HwdbgInstanceInfoIsValid) {
    ShowMessages("err, the instance info is not valid\n");
    return FALSE;
  }
  if (!(InstanceInfo->scriptVariableLength >= sizeof(BYTE) * 8)) {
    ShowMessages(
        "err, the script variable length should be at least 8 bits (1 byte)\n");
    return FALSE;
  }
  if (HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer(
          InstanceInfo, ScriptBuffer, ScriptBufferSize, NumberOfStagesForScript,
          NewScriptBuffer, NewCompressedBufferSize) == FALSE) {
    ShowMessages(
        "err, unable to convert the script buffer to short symbol buffer\n");
    return FALSE;
  }
  if (!HardwareScriptInterpreterCompressBuffer(
          (UINT64 *)*NewScriptBuffer, *NewCompressedBufferSize,
          InstanceInfo->scriptVariableLength, InstanceInfo->bramDataWidth,
          NewCompressedBufferSize, NumberOfBytesPerChunk)) {
    return FALSE;
  }
  return TRUE;
}

VOID HwdbgScriptPrintFinalScriptBufferAndHardwareDetails(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, size_t NewCompressedBufferSize,
    UINT32 NumberOfStagesForScript, UINT32 NumberOfOperandsForScript,
    HWDBG_SHORT_SYMBOL *NewScriptBuffer,
    size_t NumberOfNeededFlipFlopsInTargetDevice, size_t NumberOfBytesPerChunk,
    UINT32 NumberOfOperandsImplemented) {
  ShowMessages("\n---------------------------------------------------------\n");
  NumberOfNeededFlipFlopsInTargetDevice = HwdbgComputeNumberOfFlipFlopsNeeded(
      InstanceInfo, NumberOfStagesForScript);
  ShowMessages(
      "hwdbg script buffer (buffer size=%d, stages=%d, operands needed: %d - "
      "operands used: %d (%.2f%%), total used flip-flops=%d, number of bytes "
      "per chunk: %d):\n\n",
      NewCompressedBufferSize, NumberOfStagesForScript,
      NumberOfOperandsImplemented, NumberOfOperandsForScript,
      ((float)NumberOfOperandsForScript / (float)NumberOfOperandsImplemented) *
          100,
      NumberOfNeededFlipFlopsInTargetDevice, NumberOfBytesPerChunk);
  for (size_t i = 0; i < NewCompressedBufferSize; i++) {
    ShowMessages("%02X ", (UINT8)((CHAR *)NewScriptBuffer)[i]);
  }
}

BOOLEAN HwdbgScriptWriteScriptConfigurationPacketIntoFile(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, const CHAR *FileName,
    UINT32 NumberOfStagesForScript, UINT32 NumberOfOperandsImplemented,
    HWDBG_SHORT_SYMBOL *NewScriptBuffer, size_t NewCompressedBufferSize) {
  TCHAR TestFilePath[MAX_PATH] = {0};
  ShowMessages("\n\nwriting script configuration packet into the file\n");
  if (SetupPathForFileName(FileName, TestFilePath, sizeof(TestFilePath),
                           FALSE) &&
      HwdbgScriptSendScriptPacket(
          InstanceInfo, TestFilePath,
          NumberOfStagesForScript + NumberOfOperandsImplemented - 1,
          NewScriptBuffer, (UINT32)NewCompressedBufferSize)) {
    ShowMessages("\n[*] script buffer successfully written into file: %s\n",
                 TestFilePath);
    return TRUE;
  } else {
    ShowMessages("err, unable to write script buffer\n");
    return FALSE;
  }
}

BOOLEAN HwdbgScriptCreateHwdbgScript(
    CHAR *ScriptBuffer, UINT32 ScriptBufferSize,
    const TCHAR *HardwareScriptFilePathToSave) {
  UINT32 NumberOfStagesForScript = 0;
  UINT32 NumberOfOperandsImplemented = 0;
  UINT32 NumberOfOperandsForScript = 0;
  size_t NewCompressedBufferSize = 0;
  size_t NumberOfNeededFlipFlopsInTargetDevice = 0;
  size_t NumberOfBytesPerChunk = 0;
  HWDBG_SHORT_SYMBOL *NewScriptBuffer = NULL;
  if (!HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities(
          &g_HwdbgInstanceInfo, ScriptBuffer, ScriptBufferSize / sizeof(SYMBOL),
          &NumberOfStagesForScript, &NumberOfOperandsForScript,
          &NumberOfOperandsImplemented)) {
    ShowMessages(
        "\n[-] target script is NOT supported by this instance of hwdbg!\n");
    return FALSE;
  } else {
    ShowMessages(
        "\n[+] target script is supported by this instance of hwdbg!\n");
  }
  if (!HwdbgScriptCompressScriptBuffer(
          &g_HwdbgInstanceInfo, (SYMBOL *)ScriptBuffer, ScriptBufferSize,
          NumberOfStagesForScript, &NewScriptBuffer, &NewCompressedBufferSize,
          &NumberOfBytesPerChunk)) {
    ShowMessages("err, unable to compress the script buffer\n");
    return FALSE;
  }
  HwdbgScriptPrintFinalScriptBufferAndHardwareDetails(
      &g_HwdbgInstanceInfo, NewCompressedBufferSize, NumberOfStagesForScript,
      NumberOfOperandsForScript, NewScriptBuffer,
      NumberOfNeededFlipFlopsInTargetDevice, NumberOfBytesPerChunk,
      NumberOfOperandsImplemented);
  if (!HwdbgScriptWriteScriptConfigurationPacketIntoFile(
          &g_HwdbgInstanceInfo, HardwareScriptFilePathToSave,
          NumberOfStagesForScript, NumberOfOperandsImplemented, NewScriptBuffer,
          NewCompressedBufferSize)) {
    ShowMessages("err, unable to write script buffer\n");
    return FALSE;
  }
  if (NewScriptBuffer != NULL) {
    HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer(NewScriptBuffer);
  }
  return TRUE;
}

BOOLEAN HwdbgScriptGetScriptBufferFromRawString(string ScriptString,
                                                PVOID *CodeBuffer,
                                                UINT64 *BufferAddress,
                                                UINT32 *BufferLength,
                                                UINT32 *Pointer) {
  PVOID ResultingCodeBuffer = NULL;
  ResultingCodeBuffer =
      ScriptEngineParseWrapper((char *)ScriptString.c_str(), TRUE);
  if (ResultingCodeBuffer == NULL) {
    return FALSE;
  }
  *BufferAddress = ScriptEngineWrapperGetHead(ResultingCodeBuffer);
  *BufferLength = ScriptEngineWrapperGetSize(ResultingCodeBuffer);
  *Pointer = ScriptEngineWrapperGetPointer(ResultingCodeBuffer);
  *CodeBuffer = ResultingCodeBuffer;
  return TRUE;
}

BOOLEAN HwdbgScriptSendScriptPacket(HWDBG_INSTANCE_INFORMATION *InstanceInfo,
                                    const TCHAR *FileName,
                                    UINT32 NumberOfSymbols,
                                    HWDBG_SHORT_SYMBOL *Buffer,
                                    UINT32 BufferLength) {
  HWDBG_SCRIPT_BUFFER ScriptBuffer = {0};
  BOOLEAN Result = FALSE;
  ScriptBuffer.scriptNumberOfSymbols = NumberOfSymbols;
  CHAR *FinalBuffer =
      (CHAR *)malloc(BufferLength + sizeof(HWDBG_SCRIPT_BUFFER));
  if (!FinalBuffer) {
    return FALSE;
  }
  RtlZeroMemory(FinalBuffer, BufferLength + sizeof(HWDBG_SCRIPT_BUFFER));
  memcpy(FinalBuffer, &ScriptBuffer, sizeof(HWDBG_SCRIPT_BUFFER));
  if (Buffer != NULL) {
    memcpy(FinalBuffer + sizeof(HWDBG_SCRIPT_BUFFER), Buffer, BufferLength);
  }
  Result = HwdbgInterpreterSendPacketAndBufferToHwdbg(
      InstanceInfo, FileName,
      DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL,
      hwdbgActionConfigureScriptBuffer, FinalBuffer,
      BufferLength + sizeof(HWDBG_SCRIPT_BUFFER));
  free(FinalBuffer);
  return Result;
}

BOOLEAN HwdbgScriptRunScript(const CHAR *Script,
                             const TCHAR *InstanceFilePathToRead,
                             const TCHAR *HardwareScriptFilePathToSave,
                             UINT32 InitialBramBufferSize) {
  PVOID CodeBuffer;
  UINT64 BufferAddress;
  UINT32 BufferLength;
  UINT32 Pointer;
  if (!HwdbgLoadInstanceInfo(InstanceFilePathToRead, InitialBramBufferSize)) {
    return FALSE;
  }
  if (!HwdbgScriptGetScriptBufferFromRawString(
          Script, &CodeBuffer, &BufferAddress, &BufferLength, &Pointer)) {
    return FALSE;
  }
  HwdbgScriptPrintScriptBuffer((CHAR *)BufferAddress, BufferLength);
  if (!HwdbgScriptCreateHwdbgScript((CHAR *)BufferAddress, BufferLength,
                                    HardwareScriptFilePathToSave)) {
    ShowMessages("err, unable to create hwdbg script\n");
    return FALSE;
  }
  return TRUE;
}
