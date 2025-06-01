
#include "pch.h"

BOOLEAN DisassemblerShowInstructionsInVmxNonRootMode(PVOID Address,
                                                     UINT32 Length,
                                                     BOOLEAN Is32Bit) {
  ZydisDecoder Decoder;
  ZydisFormatter Formatter;
  SIZE_T ReadOffset = 0;
  ZydisDecodedInstruction Instruction;
  ZydisDecodedOperand Operands[ZYDIS_MAX_OPERAND_COUNT];
  ZyanStatus Status;
  CHAR PrintBuffer[128];
  PAGED_CODE();
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    LogError("Err, invalid zydis version");
    return FALSE;
  }
  if (Is32Bit) {
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&Decoder,
                                       ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                                       ZYDIS_STACK_WIDTH_32))) {
      return FALSE;
    }
  } else {
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&Decoder, ZYDIS_MACHINE_MODE_LONG_64,
                                       ZYDIS_STACK_WIDTH_64))) {
      return FALSE;
    }
  }
  if (!ZYAN_SUCCESS(
          ZydisFormatterInit(&Formatter, ZYDIS_FORMATTER_STYLE_INTEL))) {
    return FALSE;
  }
  while ((Status = ZydisDecoderDecodeFull(
              &Decoder, (PVOID)((UINT64)Address + ReadOffset),
              Length - ReadOffset, &Instruction, Operands)) !=
         ZYDIS_STATUS_NO_MORE_DATA) {
    NT_ASSERT(ZYAN_SUCCESS(Status));
    if (!ZYAN_SUCCESS(Status)) {
      ReadOffset++;
      continue;
    }
    const ZyanU64 InstrAddress = (ZyanU64)((UINT64)Address + ReadOffset);
    ZydisFormatterFormatInstruction(
        &Formatter, &Instruction, Operands, Instruction.operand_count_visible,
        PrintBuffer, sizeof(PrintBuffer), InstrAddress, NULL);
    LogInfo("+%-4X 0x%-16llX\t\t%hs\n", (UINT32)ReadOffset, InstrAddress,
            PrintBuffer);
    ReadOffset += Instruction.length;
  }
  return TRUE;
}

BOOLEAN DisassemblerShowOneInstructionInVmxNonRootMode(PVOID Address,
                                                       UINT64 ActualRip,
                                                       BOOLEAN Is32Bit) {
  ZydisDecoder Decoder;
  ZydisFormatter Formatter;
  SIZE_T ReadOffset = 0;
  ZydisDecodedInstruction Instruction;
  ZydisDecodedOperand Operands[ZYDIS_MAX_OPERAND_COUNT];
  ZyanStatus Status;
  CHAR PrintBuffer[128];
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    LogError("Err, invalid zydis version");
    return FALSE;
  }
  if (Is32Bit) {
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&Decoder,
                                       ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                                       ZYDIS_STACK_WIDTH_32))) {
      return FALSE;
    }
  } else {
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&Decoder, ZYDIS_MACHINE_MODE_LONG_64,
                                       ZYDIS_STACK_WIDTH_64))) {
      return FALSE;
    }
  }
  if (!ZYAN_SUCCESS(
          ZydisFormatterInit(&Formatter, ZYDIS_FORMATTER_STYLE_INTEL))) {
    return FALSE;
  }
  while ((Status = ZydisDecoderDecodeFull(
              &Decoder, (PVOID)((UINT64)Address + ReadOffset),
              MAXIMUM_INSTR_SIZE - ReadOffset, &Instruction, Operands)) !=
         ZYDIS_STATUS_NO_MORE_DATA) {
    NT_ASSERT(ZYAN_SUCCESS(Status));
    if (!ZYAN_SUCCESS(Status)) {
      ReadOffset++;
      continue;
    }
    const ZyanU64 InstrAddress = (ZyanU64)((UINT64)ActualRip + ReadOffset);
    ZydisFormatterFormatInstruction(
        &Formatter, &Instruction, Operands, Instruction.operand_count_visible,
        PrintBuffer, sizeof(PrintBuffer), InstrAddress, NULL);
    Log("core: %x | pid: %x - tid: %x,\t %llx \t\t\t\t%hs\n",
        KeGetCurrentProcessorNumberEx(NULL), PsGetCurrentProcessId(),
        PsGetCurrentThreadId(), ActualRip, PrintBuffer);
    break;
  }
  return TRUE;
}

UINT32 DisassemblerLengthDisassembleEngine(PVOID Address, BOOLEAN Is32Bit) {
  ZydisDecoder Decoder;
  ZydisDecodedInstruction Instruction;
  ZydisDecodedOperand Operands[ZYDIS_MAX_OPERAND_COUNT];
  ZyanStatus Status;
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    LogError("Err, invalid zydis version");
    return NULL_ZERO;
  }
  if (Is32Bit) {
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&Decoder,
                                       ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                                       ZYDIS_STACK_WIDTH_32))) {
      return NULL_ZERO;
    }
  } else {
    if (!ZYAN_SUCCESS(ZydisDecoderInit(&Decoder, ZYDIS_MACHINE_MODE_LONG_64,
                                       ZYDIS_STACK_WIDTH_64))) {
      return NULL_ZERO;
    }
  }
  while ((Status = ZydisDecoderDecodeFull(
              &Decoder, (PVOID)((UINT64)Address), MAXIMUM_INSTR_SIZE,
              &Instruction, Operands)) != ZYDIS_STATUS_NO_MORE_DATA) {
    NT_ASSERT(ZYAN_SUCCESS(Status));
    if (!ZYAN_SUCCESS(Status)) {
      return NULL_ZERO;
    }
    return Instruction.length;
  }
  return NULL_ZERO;
}

UINT32 DisassemblerLengthDisassembleEngineInVmxRootOnTargetProcess(
    PVOID Address, BOOLEAN Is32Bit) {
  BYTE SafeMemoryToRead[MAXIMUM_INSTR_SIZE] = {0};
  UINT64 SizeOfSafeBufferToRead = 0;
  SizeOfSafeBufferToRead = CheckAddressMaximumInstructionLength(Address);
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, SafeMemoryToRead,
                                            SizeOfSafeBufferToRead);
  return DisassemblerLengthDisassembleEngine(SafeMemoryToRead, Is32Bit);
}

UINT32 DisassemblerShowOneInstructionInVmxRootMode(PVOID Address,
                                                   BOOLEAN Is32Bit) {
  BYTE SafeMemoryToRead[MAXIMUM_INSTR_SIZE] = {0};
  UINT64 SizeOfSafeBufferToRead = 0;
  SizeOfSafeBufferToRead = CheckAddressMaximumInstructionLength(Address);
  MemoryMapperReadMemorySafeOnTargetProcess((UINT64)Address, SafeMemoryToRead,
                                            SizeOfSafeBufferToRead);
  return DisassemblerShowOneInstructionInVmxNonRootMode(
      SafeMemoryToRead, (UINT64)Address, Is32Bit);
}
