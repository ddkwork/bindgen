

#include "Zycore/Format.h"
#include "Zycore/LibC.h"
#include "Zydis/Zydis.h"
#include "pch.h"
#pragma comment(lib, "Zydis.lib")
#pragma comment(lib, "Zycore.lib")
extern UINT32 g_DisassemblerSyntax;
extern std::map<UINT64, LOCAL_FUNCTION_DESCRIPTION> g_DisassemblerSymbolMap;
extern BOOLEAN g_AddressConversion;

typedef struct ZydisSymbol_ {
  ZyanU64 address;

  const char *name;
} ZydisSymbol;

ZydisFormatterFunc default_print_address_absolute;

static ZyanStatus ZydisFormatterPrintAddressAbsolute(
    const ZydisFormatter *formatter, ZydisFormatterBuffer *buffer,
    ZydisFormatterContext *context) {
  ZyanU64 address;
  std::map<UINT64, LOCAL_FUNCTION_DESCRIPTION>::iterator Iterate;
  ZYAN_CHECK(ZydisCalcAbsoluteAddress(context->instruction, context->operand,
                                      context->runtime_address, &address));
  if (g_AddressConversion) {
    Iterate = g_DisassemblerSymbolMap.find(address);
    if (Iterate != g_DisassemblerSymbolMap.end()) {
      ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, ZYDIS_TOKEN_SYMBOL));
      ZyanString *string;
      ZYAN_CHECK(ZydisFormatterBufferGetString(buffer, &string));
      return ZyanStringAppendFormat(
          string, "<%s (%s)>", Iterate->second.ObjectName.c_str(),
          SeparateTo64BitValue(Iterate->first).c_str());
    }
  }
  return default_print_address_absolute(formatter, buffer, context);
}

VOID DisassembleBuffer(ZydisDecoder *decoder, ZyanU64 runtime_address,
                       ZyanU8 *data, ZyanUSize length, uint32_t maximum_instr,
                       BOOLEAN is_x86_64, BOOLEAN show_of_branch_is_taken,
                       PRFLAGS rflags) {
  ZydisFormatter formatter;
  int instr_decoded = 0;
  UINT64 UsedBaseAddress = NULL;
  if (g_DisassemblerSyntax == 1) {
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
  } else if (g_DisassemblerSyntax == 2) {
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_ATT);
  } else if (g_DisassemblerSyntax == 3) {
    ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL_MASM);
  } else {
    ShowMessages("err, in selecting disassembler syntax\n");
    return;
  }
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT,
                            ZYAN_TRUE);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE,
                            ZYAN_TRUE);
  default_print_address_absolute =
      (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
  ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
                        (const void **)&default_print_address_absolute);
  ZydisDecodedOperand operands[ZYDIS_MAX_OPERAND_COUNT];
  ZydisDecodedInstruction instruction;
  char buffer[256];
  while (ZYAN_SUCCESS(
      ZydisDecoderDecodeFull(decoder, data, length, &instruction, operands))) {
    if (g_AddressConversion) {
      if (SymbolShowFunctionNameBasedOnAddress(runtime_address,
                                               &UsedBaseAddress)) {
        ShowMessages(":\n");
      }
    }
    ShowMessages("%s   ", SeparateTo64BitValue(runtime_address).c_str());
    ZydisFormatterFormatInstruction(
        &formatter, &instruction, operands, instruction.operand_count_visible,
        &buffer[0], sizeof(buffer), runtime_address, ZYAN_NULL);
    for (size_t i = 0; i < instruction.length; i++) {
      ZyanU8 MemoryContent = data[i];
      ShowMessages(" %02X", MemoryContent);
    }
#define PaddingLength 12 if (instruction.length < PaddingLength) {
    for (size_t i = 0; i < PaddingLength - instruction.length; i++) {
      ShowMessages("   ");
    }
  }
  if (show_of_branch_is_taken) {
    RFLAGS TempRflags = {0};
    TempRflags.AsUInt = rflags->AsUInt;
    DEBUGGER_CONDITIONAL_JUMP_STATUS ResultOfCondJmp =
        HyperDbgIsConditionalJumpTaken(data, length, TempRflags, is_x86_64);
    if (ResultOfCondJmp == DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN) {
      ShowMessages(" %s [taken]\n", &buffer[0]);
    } else if (ResultOfCondJmp ==
               DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN) {
      ShowMessages(" %s [not taken]\n", &buffer[0]);
    } else {
      ShowMessages(" %s\n", &buffer[0]);
    }
  } else {
    ShowMessages(" %s\n", &buffer[0]);
  }
  data += instruction.length;
  length -= instruction.length;
  runtime_address += instruction.length;
  instr_decoded++;
  if (instr_decoded == maximum_instr) {
    return;
  }
}
}

int ZydisTest() {
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    fputs("Invalid Zydis version\n", ZYAN_STDERR);
    return EXIT_FAILURE;
  }
  ZyanU8 data[] = {0x48, 0x8B, 0x05, 0x39, 0x00, 0x13, 0x00, 0x50, 0xFF,
                   0x15, 0xF2, 0x10, 0x00, 0x00, 0x85, 0xC0, 0x0F, 0x84,
                   0x00, 0x00, 0x00, 0x00, 0xE9, 0xE5, 0x0F, 0x00, 0x00};
  ZydisDecoder decoder;
  ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_STACK_WIDTH_64);
  DisassembleBuffer(&decoder, 0x007FFFFFFF400000, &data[0], sizeof(data),
                    0xffffffff, TRUE, FALSE, NULL);
  return 0;
}

int HyperDbgDisassembler64(unsigned char *BufferToDisassemble,
                           UINT64 BaseAddress, UINT64 Size,
                           UINT32 MaximumInstrDecoded,
                           BOOLEAN ShowBranchIsTakenOrNot, PRFLAGS Rflags) {
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    fputs("Invalid Zydis version\n", ZYAN_STDERR);
    return EXIT_FAILURE;
  }
  ZydisDecoder decoder;
  ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64, ZYDIS_STACK_WIDTH_64);
  DisassembleBuffer(&decoder, BaseAddress, &BufferToDisassemble[0], Size,
                    MaximumInstrDecoded, TRUE, ShowBranchIsTakenOrNot, Rflags);
  return 0;
}

int HyperDbgDisassembler32(unsigned char *BufferToDisassemble,
                           UINT64 BaseAddress, UINT64 Size,
                           UINT32 MaximumInstrDecoded,
                           BOOLEAN ShowBranchIsTakenOrNot, PRFLAGS Rflags) {
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    fputs("Invalid Zydis version\n", ZYAN_STDERR);
    return EXIT_FAILURE;
  }
  ZydisDecoder decoder;
  ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                   ZYDIS_STACK_WIDTH_32);
  DisassembleBuffer(&decoder, (UINT32)BaseAddress, &BufferToDisassemble[0],
                    Size, MaximumInstrDecoded, FALSE, ShowBranchIsTakenOrNot,
                    Rflags);
  return 0;
}

DEBUGGER_CONDITIONAL_JUMP_STATUS HyperDbgIsConditionalJumpTaken(
    unsigned char *BufferToDisassemble, UINT64 BuffLength, RFLAGS Rflags,
    BOOLEAN Isx86_64) {
  ZydisDecoder decoder;
  ZydisFormatter formatter;
  ZydisDecodedOperand operands[ZYDIS_MAX_OPERAND_COUNT];
  UINT64 CurrentRip = 0;
  int instr_decoded = 0;
  ZydisDecodedInstruction instruction;
  UINT32 MaximumInstrDecoded = 1;
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    ShowMessages("invalid Zydis version\n");
    return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
  }
  if (Isx86_64) {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64,
                     ZYDIS_STACK_WIDTH_64);
  } else {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                     ZYDIS_STACK_WIDTH_32);
  }
  ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT,
                            ZYAN_TRUE);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE,
                            ZYAN_TRUE);
  default_print_address_absolute =
      (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
  ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
                        (const void **)&default_print_address_absolute);
  while (ZYAN_SUCCESS(ZydisDecoderDecodeFull(
      &decoder, BufferToDisassemble, BuffLength, &instruction, operands))) {
    switch (instruction.mnemonic) {
      case ZydisMnemonic::ZYDIS_MNEMONIC_JO:
        if (Rflags.OverflowFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JNO:
        if (!Rflags.OverflowFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JS:
        if (Rflags.SignFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JNS:
        if (!Rflags.SignFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JZ:
        if (Rflags.ZeroFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JNZ:
        if (!Rflags.ZeroFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JB:
        if (Rflags.CarryFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JNB:
        if (!Rflags.CarryFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JBE:
        if (Rflags.CarryFlag || Rflags.ZeroFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JNBE:
        if (!Rflags.CarryFlag && !Rflags.ZeroFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JL:
        if (Rflags.SignFlag != Rflags.OverflowFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JNL:
        if (Rflags.SignFlag == Rflags.OverflowFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JLE:
        if (Rflags.ZeroFlag || Rflags.SignFlag != Rflags.OverflowFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JNLE:
        if (!Rflags.ZeroFlag && Rflags.SignFlag == Rflags.OverflowFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JP:
        if (Rflags.ParityFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JNP:
        if (!Rflags.ParityFlag)
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN;
        else
          return DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN;
        break;
      case ZydisMnemonic::ZYDIS_MNEMONIC_JCXZ:
      case ZydisMnemonic::ZYDIS_MNEMONIC_JECXZ:
        return DEBUGGER_CONDITIONAL_JUMP_STATUS_NOT_CONDITIONAL_JUMP;
      default:
        return DEBUGGER_CONDITIONAL_JUMP_STATUS_NOT_CONDITIONAL_JUMP;
        break;
    }
    return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
  }
  return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
}

BOOLEAN HyperDbgCheckWhetherTheCurrentInstructionIsCall(
    unsigned char *BufferToDisassemble, UINT64 BuffLength, BOOLEAN Isx86_64,
    PUINT32 CallLength) {
  ZydisDecoder decoder;
  ZydisFormatter formatter;
  ZydisDecodedOperand operands[ZYDIS_MAX_OPERAND_COUNT];
  UINT64 CurrentRip = 0;
  int instr_decoded = 0;
  ZydisDecodedInstruction instruction;
  char buffer[256];
  UINT32 MaximumInstrDecoded = 1;
  *CallLength = 0;
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    ShowMessages("invalid zydis version\n");
    return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
  }
  if (Isx86_64) {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64,
                     ZYDIS_STACK_WIDTH_64);
  } else {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                     ZYDIS_STACK_WIDTH_32);
  }
  ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT,
                            ZYAN_TRUE);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE,
                            ZYAN_TRUE);
  default_print_address_absolute =
      (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
  ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
                        (const void **)&default_print_address_absolute);
  while (ZYAN_SUCCESS(ZydisDecoderDecodeFull(
      &decoder, BufferToDisassemble, BuffLength, &instruction, operands))) {
    ZydisFormatterFormatInstruction(
        &formatter, &instruction, operands, instruction.operand_count_visible,
        &buffer[0], sizeof(buffer), (ZyanU64)CurrentRip, ZYAN_NULL);
    if (instruction.mnemonic == ZydisMnemonic::ZYDIS_MNEMONIC_CALL) {
      *CallLength = instruction.length;
      return TRUE;
    } else {
      return FALSE;
    }
  }
  return FALSE;
}

UINT32 HyperDbgLengthDisassemblerEngine(unsigned char *BufferToDisassemble,
                                        UINT64 BuffLength, BOOLEAN Isx86_64) {
  ZydisDecoder decoder;
  ZydisFormatter formatter;
  ZydisDecodedOperand operands[ZYDIS_MAX_OPERAND_COUNT];
  UINT64 CurrentRip = 0;
  int instr_decoded = 0;
  ZydisDecodedInstruction instruction;
  UINT32 MaximumInstrDecoded = 1;
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    ShowMessages("invalid Zydis version\n");
    return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
  }
  if (Isx86_64) {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64,
                     ZYDIS_STACK_WIDTH_64);
  } else {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                     ZYDIS_STACK_WIDTH_32);
  }
  ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT,
                            ZYAN_TRUE);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE,
                            ZYAN_TRUE);
  default_print_address_absolute =
      (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
  ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
                        (const void **)&default_print_address_absolute);
  while (ZYAN_SUCCESS(ZydisDecoderDecodeFull(
      &decoder, BufferToDisassemble, BuffLength, &instruction, operands))) {
    return instruction.length;
  }
  return 0;
}

static ZyanStatus ZydisFormatterPrintAddressAbsoluteForTrackingInstructions(
    const ZydisFormatter *formatter, ZydisFormatterBuffer *buffer,
    ZydisFormatterContext *context) {
  ZyanU64 address;
  std::map<UINT64, LOCAL_FUNCTION_DESCRIPTION>::iterator Iterate;
  ZYAN_CHECK(ZydisCalcAbsoluteAddress(context->instruction, context->operand,
                                      context->runtime_address, &address));
  if (g_AddressConversion) {
    Iterate = g_DisassemblerSymbolMap.find(address);
    if (Iterate != g_DisassemblerSymbolMap.end()) {
      ZYAN_CHECK(ZydisFormatterBufferAppend(buffer, ZYDIS_TOKEN_SYMBOL));
      ZyanString *string;
      ZYAN_CHECK(ZydisFormatterBufferGetString(buffer, &string));
      CommandTrackHandleReceivedCallInstructions(
          Iterate->second.ObjectName.c_str(), Iterate->first);
      return ZyanStringAppendFormat(
          string, "<%s (%s)>", Iterate->second.ObjectName.c_str(),
          SeparateTo64BitValue(Iterate->first).c_str());
    }
  }
  CommandTrackHandleReceivedCallInstructions(NULL, address);
  return default_print_address_absolute(formatter, buffer, context);
}

BOOLEAN HyperDbgCheckWhetherTheCurrentInstructionIsCallOrRet(
    unsigned char *BufferToDisassemble, UINT64 CurrentRip, UINT32 BuffLength,
    BOOLEAN Isx86_64, PBOOLEAN IsRet) {
  ZydisDecoder decoder;
  ZydisFormatter formatter;
  ZydisDecodedOperand operands[ZYDIS_MAX_OPERAND_COUNT];
  int instr_decoded = 0;
  ZydisDecodedInstruction instruction;
  char buffer[256];
  UINT32 MaximumInstrDecoded = 1;
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    ShowMessages("invalid zydis version\n");
    return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
  }
  if (!Isx86_64) {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64,
                     ZYDIS_STACK_WIDTH_64);
  } else {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                     ZYDIS_STACK_WIDTH_32);
  }
  ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT,
                            ZYAN_TRUE);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE,
                            ZYAN_TRUE);
  default_print_address_absolute =
      (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsoluteForTrackingInstructions;
  ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
                        (const void **)&default_print_address_absolute);
  while (ZYAN_SUCCESS(ZydisDecoderDecodeFull(
      &decoder, BufferToDisassemble, BuffLength, &instruction, operands))) {
    if (instruction.mnemonic == ZydisMnemonic::ZYDIS_MNEMONIC_CALL) {
      ZydisFormatterFormatInstruction(
          &formatter, &instruction, operands, instruction.operand_count_visible,
          &buffer[0], sizeof(buffer), (ZyanU64)CurrentRip, ZYAN_NULL);
      *IsRet = FALSE;
      return TRUE;
    } else if (instruction.mnemonic == ZydisMnemonic::ZYDIS_MNEMONIC_RET) {
      CommandTrackHandleReceivedRetInstructions(CurrentRip);
      *IsRet = TRUE;
      return TRUE;
    } else {
      return FALSE;
    }
  }
  return FALSE;
}

BOOLEAN HyperDbgCheckWhetherTheCurrentInstructionIsRet(
    unsigned char *BufferToDisassemble, UINT64 BuffLength, BOOLEAN Isx86_64) {
  ZydisDecoder decoder;
  ZydisFormatter formatter;
  ZydisDecodedOperand operands[ZYDIS_MAX_OPERAND_COUNT];
  UINT64 CurrentRip = 0;
  int instr_decoded = 0;
  ZydisDecodedInstruction instruction;
  char buffer[256];
  UINT32 MaximumInstrDecoded = 1;
  if (ZydisGetVersion() != ZYDIS_VERSION) {
    ShowMessages("invalid zydis version\n");
    return DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR;
  }
  if (Isx86_64) {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_64,
                     ZYDIS_STACK_WIDTH_64);
  } else {
    ZydisDecoderInit(&decoder, ZYDIS_MACHINE_MODE_LONG_COMPAT_32,
                     ZYDIS_STACK_WIDTH_32);
  }
  ZydisFormatterInit(&formatter, ZYDIS_FORMATTER_STYLE_INTEL);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SEGMENT,
                            ZYAN_TRUE);
  ZydisFormatterSetProperty(&formatter, ZYDIS_FORMATTER_PROP_FORCE_SIZE,
                            ZYAN_TRUE);
  default_print_address_absolute =
      (ZydisFormatterFunc)&ZydisFormatterPrintAddressAbsolute;
  ZydisFormatterSetHook(&formatter, ZYDIS_FORMATTER_FUNC_PRINT_ADDRESS_ABS,
                        (const void **)&default_print_address_absolute);
  while (ZYAN_SUCCESS(ZydisDecoderDecodeFull(
      &decoder, BufferToDisassemble, BuffLength, &instruction, operands))) {
    ZydisFormatterFormatInstruction(
        &formatter, &instruction, operands, instruction.operand_count_visible,
        &buffer[0], sizeof(buffer), (ZyanU64)CurrentRip, ZYAN_NULL);
    if (instruction.mnemonic == ZydisMnemonic::ZYDIS_MNEMONIC_RET) {
      return TRUE;
    } else {
      return FALSE;
    }
  }
  return FALSE;
}
