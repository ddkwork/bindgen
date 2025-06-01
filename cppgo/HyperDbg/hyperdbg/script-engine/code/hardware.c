
#include "pch.h"

VOID HardwareScriptInterpreterShowScriptCapabilities(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo) {
  ShowMessages("\nThis debuggee supports the following operatiors:\n");
  ShowMessages(
      "\tlocal and global variable assignments: %s (maximum number of var: %d) "
      "\n",
      InstanceInfo->scriptCapabilities.assign_local_global_var
          ? "supported"
          : "not supported",
      InstanceInfo->numberOfSupportedLocalAndGlobalVariables);
  ShowMessages("\tregisters (pin/ports) assignment: %s \n",
               InstanceInfo->scriptCapabilities.assign_registers
                   ? "supported"
                   : "not supported");
  ShowMessages("\tpseudo-registers assignment: %s \n",
               InstanceInfo->scriptCapabilities.assign_pseudo_registers
                   ? "supported"
                   : "not supported");
  ShowMessages("\tconditional statements and comparison operators: %s \n",
               InstanceInfo->scriptCapabilities
                       .conditional_statements_and_comparison_operators
                   ? "supported"
                   : "not supported");
  ShowMessages("\tor: %s \n", InstanceInfo->scriptCapabilities.func_or
                                  ? "supported"
                                  : "not supported");
  ShowMessages("\txor: %s \n", InstanceInfo->scriptCapabilities.func_xor
                                   ? "supported"
                                   : "not supported");
  ShowMessages("\tand: %s \n", InstanceInfo->scriptCapabilities.func_and
                                   ? "supported"
                                   : "not supported");
  ShowMessages("\tarithmetic shift right: %s \n",
               InstanceInfo->scriptCapabilities.func_asr ? "supported"
                                                         : "not supported");
  ShowMessages("\tarithmetic shift left: %s \n",
               InstanceInfo->scriptCapabilities.func_asl ? "supported"
                                                         : "not supported");
  ShowMessages("\taddition: %s \n", InstanceInfo->scriptCapabilities.func_add
                                        ? "supported"
                                        : "not supported");
  ShowMessages("\tsubtraction: %s \n", InstanceInfo->scriptCapabilities.func_sub
                                           ? "supported"
                                           : "not supported");
  ShowMessages("\tmultiplication: %s \n",
               InstanceInfo->scriptCapabilities.func_mul ? "supported"
                                                         : "not supported");
  ShowMessages("\tdivision: %s \n", InstanceInfo->scriptCapabilities.func_div
                                        ? "supported"
                                        : "not supported");
  ShowMessages("\tmodulus: %s \n", InstanceInfo->scriptCapabilities.func_mod
                                       ? "supported"
                                       : "not supported");
  ShowMessages("\tgreater than: %s \n",
               (InstanceInfo->scriptCapabilities.func_gt &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tless than: %s \n",
               (InstanceInfo->scriptCapabilities.func_lt &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tgreater than or equal to: %s \n",
               (InstanceInfo->scriptCapabilities.func_egt &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tless than or equal to: %s \n",
               (InstanceInfo->scriptCapabilities.func_elt &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tequal: %s \n",
               (InstanceInfo->scriptCapabilities.func_equal &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tnot equal: %s \n",
               (InstanceInfo->scriptCapabilities.func_neq &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tjump: %s \n",
               (InstanceInfo->scriptCapabilities.func_jmp &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tjump if zero: %s \n",
               (InstanceInfo->scriptCapabilities.func_jz &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tjump if not zero: %s \n",
               (InstanceInfo->scriptCapabilities.func_jnz &&
                InstanceInfo->scriptCapabilities
                    .conditional_statements_and_comparison_operators)
                   ? "supported"
                   : "not supported");
  ShowMessages("\tmove: %s \n", InstanceInfo->scriptCapabilities.func_mov
                                    ? "supported"
                                    : "not supported");
  ShowMessages("\tprintf: %s \n", InstanceInfo->scriptCapabilities.func_printf
                                      ? "supported"
                                      : "not supported");
  ShowMessages("\n");
}

BOOLEAN HardwareScriptInterpreterCheckScriptBufferWithScriptCapabilities(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, PVOID ScriptBuffer,
    UINT32 CountOfScriptSymbolChunks, UINT32 *NumberOfStages,
    UINT32 *NumberOfOperands, UINT32 *NumberOfOperandsImplemented) {
  BOOLEAN NotSupported = FALSE;
  SYMBOL *SymbolArray = (SYMBOL *)ScriptBuffer;
  UINT32 Stages = 0;
  UINT32 Operands = 0;
  UINT32 NumberOfGetOperands = 0;
  UINT32 NumberOfSetOperands = 0;
  for (size_t i = 0; i < CountOfScriptSymbolChunks; i++) {
    if (SymbolArray[i].Type != SYMBOL_SEMANTIC_RULE_TYPE) {
      Operands++;
      ShowMessages(
          "  \t%lld. found a non-semnatic rule (operand) | type: 0x%llx, "
          "value: 0x%llx\n",
          i, SymbolArray[i].Type, SymbolArray[i].Value);
      switch (SymbolArray[i].Type) {
        case SYMBOL_GLOBAL_ID_TYPE:
        case SYMBOL_LOCAL_ID_TYPE:
          if (!InstanceInfo->scriptCapabilities.assign_local_global_var) {
            NotSupported = TRUE;
            ShowMessages(
                "err, global/local variable assignment is not supported\n");
          }
          if (SymbolArray[i].Value >=
              InstanceInfo->numberOfSupportedLocalAndGlobalVariables) {
            NotSupported = TRUE;
            ShowMessages(
                "err, global/local variable index is out of range of supported "
                "by this instance of hwdbg\n");
          }
          break;
        case SYMBOL_UNDEFINED:
        case SYMBOL_NUM_TYPE:
          break;
        case SYMBOL_REGISTER_TYPE:
          if (!InstanceInfo->scriptCapabilities.assign_registers) {
            NotSupported = TRUE;
            ShowMessages("err, register assignment is not supported\n");
          }
          break;
        case SYMBOL_PSEUDO_REG_TYPE:
          if (!InstanceInfo->scriptCapabilities.assign_pseudo_registers) {
            NotSupported = TRUE;
            ShowMessages("err, pseudo register index is not supported\n");
          }
          break;
        case SYMBOL_TEMP_TYPE:
          if (!InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages(
                "err, temporary variables (for conditional statement) is not "
                "supported\n");
          }
          if (SymbolArray[i].Value >=
              InstanceInfo->numberOfSupportedTemporaryVariables) {
            NotSupported = TRUE;
            ShowMessages(
                "err, temp variable index (number of operands for conditional "
                "statements) is out of range of supported by this instance of "
                "hwdbg\n");
          }
          break;
        case SYMBOL_STACK_INDEX_TYPE:
          if (!InstanceInfo->scriptCapabilities.stack_assignments) {
            NotSupported = TRUE;
            ShowMessages(
                "err, temporary variables (for conditional statement) is not "
                "supported\n");
          }
          break;
        default:
          NotSupported = TRUE;
          ShowMessages("err, unknown operand type: %lld (0x%llx)\n",
                       SymbolArray[i].Type, SymbolArray[i].Type);
          break;
      }
    } else {
      Stages++;
      ShowMessages(
          "- %lld. found a semnatic rule (operator) | type: 0x%llx, value: "
          "0x%llx\n",
          i, SymbolArray[i].Type, SymbolArray[i].Value);
      if (FuncGetNumberOfOperands(SymbolArray[i].Type, &NumberOfGetOperands,
                                  &NumberOfSetOperands) == FALSE) {
        NotSupported = TRUE;
        ShowMessages("err, unknown operand type for the operator (0x%llx)\n",
                     SymbolArray[i].Type);
        return FALSE;
      }
      switch (SymbolArray[i].Value) {
        case FUNC_OR:
          if (!InstanceInfo->scriptCapabilities.func_or) {
            NotSupported = TRUE;
            ShowMessages("err, OR is not supported by the debuggee\n");
          }
          break;
        case FUNC_XOR:
          if (!InstanceInfo->scriptCapabilities.func_xor) {
            NotSupported = TRUE;
            ShowMessages("err, XOR is not supported by the debuggee\n");
          }
          break;
        case FUNC_AND:
          if (!InstanceInfo->scriptCapabilities.func_and) {
            NotSupported = TRUE;
            ShowMessages("err, AND is not supported by the debuggee\n");
          }
          break;
        case FUNC_ASR:
          if (!InstanceInfo->scriptCapabilities.func_asr) {
            NotSupported = TRUE;
            ShowMessages(
                "err, arithmetic shift right is not supported by the "
                "debuggee\n");
          }
          break;
        case FUNC_ASL:
          if (!InstanceInfo->scriptCapabilities.func_asl) {
            NotSupported = TRUE;
            ShowMessages(
                "err, arithmetic shift left is not supported by the "
                "debuggee\n");
          }
          break;
        case FUNC_ADD:
          if (!InstanceInfo->scriptCapabilities.func_add) {
            NotSupported = TRUE;
            ShowMessages("err, addition is not supported by the debuggee\n");
          }
          break;
        case FUNC_SUB:
          if (!InstanceInfo->scriptCapabilities.func_sub) {
            NotSupported = TRUE;
            ShowMessages("err, subtraction is not supported by the debuggee\n");
          }
          break;
        case FUNC_MUL:
          if (!InstanceInfo->scriptCapabilities.func_mul) {
            NotSupported = TRUE;
            ShowMessages(
                "err, multiplication is not supported by the debuggee\n");
          }
          break;
        case FUNC_DIV:
          if (!InstanceInfo->scriptCapabilities.func_div) {
            NotSupported = TRUE;
            ShowMessages("err, division is not supported by the debuggee\n");
          }
          break;
        case FUNC_MOD:
          if (!InstanceInfo->scriptCapabilities.func_mod) {
            NotSupported = TRUE;
            ShowMessages("err, modulus is not supported by the debuggee\n");
          }
          break;
        case FUNC_GT:
          if (!InstanceInfo->scriptCapabilities.func_gt ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages(
                "err, greater than is not supported by the debuggee\n");
          }
          break;
        case FUNC_LT:
          if (!InstanceInfo->scriptCapabilities.func_lt ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages("err, less than is not supported by the debuggee\n");
          }
          break;
        case FUNC_EGT:
          if (!InstanceInfo->scriptCapabilities.func_egt ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages(
                "err, greater than or equal to is not supported by the "
                "debuggee\n");
          }
          break;
        case FUNC_ELT:
          if (!InstanceInfo->scriptCapabilities.func_elt ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages(
                "err, less than or equal to is not supported by the "
                "debuggee\n");
          }
          break;
        case FUNC_EQUAL:
          if (!InstanceInfo->scriptCapabilities.func_equal ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages("err, equal is not supported by the debuggee\n");
          }
          break;
        case FUNC_NEQ:
          if (!InstanceInfo->scriptCapabilities.func_neq ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages("err, not equal is not supported by the debuggee\n");
          }
          break;
        case FUNC_JMP:
          if (!InstanceInfo->scriptCapabilities.func_jmp ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages("err, jump is not supported by the debuggee\n");
          }
          break;
        case FUNC_JZ:
          if (!InstanceInfo->scriptCapabilities.func_jz ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages(
                "err, jump if zero is not supported by the debuggee\n");
          }
          break;
        case FUNC_JNZ:
          if (!InstanceInfo->scriptCapabilities.func_jnz ||
              !InstanceInfo->scriptCapabilities
                   .conditional_statements_and_comparison_operators) {
            NotSupported = TRUE;
            ShowMessages(
                "err, jump if not zero is not supported by the debuggee\n");
          }
          break;
        case FUNC_MOV:
          if (!InstanceInfo->scriptCapabilities.func_mov) {
            NotSupported = TRUE;
            ShowMessages("err, move is not supported by the debuggee\n");
          }
          break;
        case FUNC_PRINTF:
          if (!InstanceInfo->scriptCapabilities.func_printf) {
            NotSupported = TRUE;
            ShowMessages("err, printf is not supported by the debuggee\n");
          }
          break;
        default:
          NotSupported = TRUE;
          ShowMessages("err, undefined operator for hwdbg: %lld (0x%llx)\n",
                       SymbolArray[i].Type, SymbolArray[i].Type);
          break;
      }
    }
  }
  *NumberOfStages = Stages;
  *NumberOfOperands = Operands;
  *NumberOfOperandsImplemented =
      *NumberOfStages *
      (InstanceInfo->maximumNumberOfSupportedGetScriptOperators +
       InstanceInfo->maximumNumberOfSupportedSetScriptOperators);
  if (NotSupported) {
    return FALSE;
  } else {
    return TRUE;
  }
}

BOOLEAN HardwareScriptInterpreterCompressBuffer(UINT64 *Buffer,
                                                size_t BufferLength,
                                                UINT32 ScriptVariableLength,
                                                UINT32 BramDataWidth,
                                                size_t *NewBufferSize,
                                                size_t *NumberOfBytesPerChunk) {
  if (ScriptVariableLength <= 7 || ScriptVariableLength > 64) {
    ShowMessages("err, invalid bit size, it should be between 7 and 64\n");
    return FALSE;
  }
  if (ScriptVariableLength > BramDataWidth) {
    ShowMessages(
        "err, script variable length cannot be more than the BRAM data "
        "width\n");
    return FALSE;
  }
  size_t NumberOfChunks = BufferLength / sizeof(UINT64);
  size_t NewBytesPerChunk = (BramDataWidth + 7) / 8;
  *NumberOfBytesPerChunk = NewBytesPerChunk;
  *NewBufferSize = NumberOfChunks * NewBytesPerChunk;
  UINT8 *TempBuffer = (UINT8 *)malloc(*NewBufferSize);
  if (TempBuffer == NULL) {
    ShowMessages("err, memory allocation failed\n");
    return FALSE;
  }
  for (size_t i = 0; i < NumberOfChunks; ++i) {
    uint64_t Chunk = Buffer[i];
    for (size_t j = 0; j < NewBytesPerChunk; ++j) {
      TempBuffer[i * NewBytesPerChunk + j] =
          (uint8_t)((Chunk >> (j * 8)) & 0xFF);
    }
  }
  RtlZeroMemory(Buffer, BufferLength);
  memcpy(Buffer, TempBuffer, *NewBufferSize);
  free(TempBuffer);
  return TRUE;
}

BOOLEAN HardwareScriptInterpreterConvertSymbolToHwdbgShortSymbolBuffer(
    HWDBG_INSTANCE_INFORMATION *InstanceInfo, SYMBOL *SymbolBuffer,
    size_t SymbolBufferLength, UINT32 NumberOfStages,
    HWDBG_SHORT_SYMBOL **NewShortSymbolBuffer, size_t *NewBufferSize) {
  if (!g_HwdbgInstanceInfoIsValid) {
    ShowMessages("err, instance info is not valid\n");
    return FALSE;
  }
  UINT32 NumberOfOperands =
      InstanceInfo->maximumNumberOfSupportedGetScriptOperators +
      InstanceInfo->maximumNumberOfSupportedSetScriptOperators;
  SIZE_T NumberOfSymbols = SymbolBufferLength / sizeof(SymbolBuffer[0]);
  *NewBufferSize =
      NumberOfStages * (NumberOfOperands + 1) * sizeof(HWDBG_SHORT_SYMBOL);
  HWDBG_SHORT_SYMBOL *HwdbgShortSymbolBuffer =
      (HWDBG_SHORT_SYMBOL *)malloc(*NewBufferSize);
  if (!HwdbgShortSymbolBuffer) {
    ShowMessages("err, could not allocate compression buffer\n");
    return FALSE;
  }
  RtlZeroMemory(HwdbgShortSymbolBuffer, *NewBufferSize);
  UINT32 IndexOfShortSymbolBuffer = 0;
  for (UINT32 i = 0; i < NumberOfSymbols; i++) {
    if (SymbolBuffer[i].Type == SYMBOL_SEMANTIC_RULE_TYPE) {
      HwdbgShortSymbolBuffer[IndexOfShortSymbolBuffer].Type =
          SymbolBuffer[i].Type;
      HwdbgShortSymbolBuffer[IndexOfShortSymbolBuffer].Value =
          SymbolBuffer[i].Value;
      UINT32 NumberOfGetOperands = 0;
      UINT32 NumberOfSetOperands = 0;
      if (!FuncGetNumberOfOperands(SymbolBuffer[i].Value, &NumberOfGetOperands,
                                   &NumberOfSetOperands)) {
        ShowMessages("err, unknown operand type for the operator (0x%llx)\n",
                     SymbolBuffer[i].Value);
        free(HwdbgShortSymbolBuffer);
        return FALSE;
      }
      if (NumberOfGetOperands >
          InstanceInfo->maximumNumberOfSupportedGetScriptOperators) {
        ShowMessages(
            "err, the number of get operands is more than the maximum "
            "supported get operands\n");
        free(HwdbgShortSymbolBuffer);
        return FALSE;
      }
      if (NumberOfSetOperands >
          InstanceInfo->maximumNumberOfSupportedSetScriptOperators) {
        ShowMessages(
            "err, the number of set operands is more than the maximum "
            "supported set operands\n");
        free(HwdbgShortSymbolBuffer);
        return FALSE;
      }
      for (size_t j = 0; j < NumberOfGetOperands; j++) {
        i++;
        IndexOfShortSymbolBuffer++;
        if (SymbolBuffer[i].Type == SYMBOL_SEMANTIC_RULE_TYPE) {
          ShowMessages("err, not expecting a semantic rule at operand: %llx\n",
                       SymbolBuffer[i].Value);
          free(HwdbgShortSymbolBuffer);
          return FALSE;
        }
        HwdbgShortSymbolBuffer[IndexOfShortSymbolBuffer].Type =
            SymbolBuffer[i].Type;
        HwdbgShortSymbolBuffer[IndexOfShortSymbolBuffer].Value =
            SymbolBuffer[i].Value;
      }
      IndexOfShortSymbolBuffer =
          IndexOfShortSymbolBuffer +
          InstanceInfo->maximumNumberOfSupportedGetScriptOperators -
          NumberOfGetOperands;
      for (size_t j = 0; j < NumberOfSetOperands; j++) {
        i++;
        IndexOfShortSymbolBuffer++;
        if (SymbolBuffer[i].Type == SYMBOL_SEMANTIC_RULE_TYPE) {
          ShowMessages("err, not expecting a semantic rule at operand: %llx\n",
                       SymbolBuffer[i].Value);
          free(HwdbgShortSymbolBuffer);
          return FALSE;
        }
        HwdbgShortSymbolBuffer[IndexOfShortSymbolBuffer].Type =
            SymbolBuffer[i].Type;
        HwdbgShortSymbolBuffer[IndexOfShortSymbolBuffer].Value =
            SymbolBuffer[i].Value;
      }
      IndexOfShortSymbolBuffer =
          IndexOfShortSymbolBuffer +
          InstanceInfo->maximumNumberOfSupportedSetScriptOperators -
          NumberOfSetOperands;
      IndexOfShortSymbolBuffer++;
    } else {
      ShowMessages("err, not expecting a non-semantic rule at: %llx\n",
                   SymbolBuffer[i].Type);
      free(HwdbgShortSymbolBuffer);
      return FALSE;
    }
  }
  *NewShortSymbolBuffer = (HWDBG_SHORT_SYMBOL *)HwdbgShortSymbolBuffer;
  return TRUE;
}

VOID HardwareScriptInterpreterFreeHwdbgShortSymbolBuffer(
    HWDBG_SHORT_SYMBOL *NewShortSymbolBuffer) {
  if (NewShortSymbolBuffer) {
    free(NewShortSymbolBuffer);
  }
}
