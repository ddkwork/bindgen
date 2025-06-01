﻿
#include "pch.h"
extern HWDBG_INSTANCE_INFORMATION g_HwdbgInstanceInfo;
extern BOOLEAN g_HwdbgInstanceInfoIsValid;
extern PVOID g_MessageHandler;

VOID ShowMessages(const char *Fmt, ...) {
  va_list ArgList;
  va_list Args;
  if (g_MessageHandler == NULL) {
    va_start(Args, Fmt);
    vprintf(Fmt, Args);
    va_end(Args);
  } else {
    char TempMessage[COMMUNICATION_BUFFER_SIZE +
                     TCP_END_OF_BUFFER_CHARS_COUNT] = {0};
    va_start(ArgList, Fmt);
    int sprintfresult = vsprintf_s(
        TempMessage, COMMUNICATION_BUFFER_SIZE + TCP_END_OF_BUFFER_CHARS_COUNT,
        Fmt, ArgList);
    va_end(ArgList);
    if (sprintfresult != -1) {
      ((SendMessageWithParamCallback)g_MessageHandler)(TempMessage);
    }
  }
}

UINT64 ScriptEngineConvertNameToAddress(const char *FunctionOrVariableName,
                                        PBOOLEAN WasFound) {
  return SymConvertNameToAddress(FunctionOrVariableName, WasFound);
}

UINT32 ScriptEngineLoadFileSymbol(UINT64 BaseAddress, const char *PdbFileName,
                                  const char *CustomModuleName) {
  return SymLoadFileSymbol(BaseAddress, PdbFileName, CustomModuleName);
}

VOID ScriptEngineSetTextMessageCallback(PVOID Handler) {
  g_MessageHandler = Handler;
  SymSetTextMessageCallback(Handler);
}

UINT32 ScriptEngineUnloadAllSymbols() { return SymUnloadAllSymbols(); }

UINT32 ScriptEngineUnloadModuleSymbol(char *ModuleName) {
  return SymUnloadModuleSymbol(ModuleName);
}

UINT32 ScriptEngineSearchSymbolForMask(const char *SearchMask) {
  return SymSearchSymbolForMask(SearchMask);
}

BOOLEAN ScriptEngineGetFieldOffset(CHAR *TypeName, CHAR *FieldName,
                                   UINT32 *FieldOffset) {
  return SymGetFieldOffset(TypeName, FieldName, FieldOffset);
}

BOOLEAN ScriptEngineGetDataTypeSize(CHAR *TypeName, UINT64 *TypeSize) {
  return SymGetDataTypeSize(TypeName, TypeSize);
}

BOOLEAN ScriptEngineCreateSymbolTableForDisassembler(void *CallbackFunction) {
  return SymCreateSymbolTableForDisassembler(CallbackFunction);
}

BOOLEAN ScriptEngineConvertFileToPdbPath(const char *LocalFilePath,
                                         char *ResultPath,
                                         size_t ResultPathSize) {
  return SymConvertFileToPdbPath(LocalFilePath, ResultPath, ResultPathSize);
}

BOOLEAN ScriptEngineSymbolInitLoad(PVOID BufferToStoreDetails,
                                   UINT32 StoredLength,
                                   BOOLEAN DownloadIfAvailable,
                                   const char *SymbolPath,
                                   BOOLEAN IsSilentLoad) {
  return SymbolInitLoad(BufferToStoreDetails, StoredLength, DownloadIfAvailable,
                        SymbolPath, IsSilentLoad);
}

BOOLEAN ScriptEngineShowDataBasedOnSymbolTypes(
    const char *TypeName, UINT64 Address, BOOLEAN IsStruct, PVOID BufferAddress,
    const char *AdditionalParameters) {
  return SymShowDataBasedOnSymbolTypes(TypeName, Address, IsStruct,
                                       BufferAddress, AdditionalParameters);
}

VOID ScriptEngineSymbolAbortLoading() { SymbolAbortLoading(); }

BOOLEAN ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(
    const char *LocalFilePath, char *PdbFilePath, char *GuidAndAgeDetails,
    BOOLEAN Is32BitModule) {
  return SymConvertFileToPdbFileAndGuidAndAgeDetails(
      LocalFilePath, PdbFilePath, GuidAndAgeDetails, Is32BitModule);
}

PVOID ScriptEngineParse(char *str) {
  PTOKEN_LIST Stack = NewTokenList();
  PTOKEN_LIST MatchedStack = NewTokenList();
  PSYMBOL_BUFFER CodeBuffer = NewSymbolBuffer();
  UserDefinedFunctionHead = malloc(sizeof(USER_DEFINED_FUNCTION_NODE));
  RtlZeroMemory(UserDefinedFunctionHead, sizeof(USER_DEFINED_FUNCTION_NODE));
  UserDefinedFunctionHead->Name = _strdup("main");
  UserDefinedFunctionHead->IdTable = (unsigned long long)NewTokenList();
  UserDefinedFunctionHead->FunctionParameterIdTable =
      (unsigned long long)NewTokenList();
  UserDefinedFunctionHead->TempMap = calloc(MAX_TEMP_COUNT, 1);
  UserDefinedFunctionHead->VariableType =
      (unsigned long long)VARIABLE_TYPE_VOID;
  CurrentUserDefinedFunction = UserDefinedFunctionHead;
  SCRIPT_ENGINE_ERROR_TYPE Error = SCRIPT_ENGINE_ERROR_FREE;
  char *ErrorMessage = NULL;
  static FirstCall = 1;
  if (FirstCall) {
    GlobalIdTable = NewTokenList();
    FirstCall = 0;
  }
  PTOKEN TopToken = NewUnknownToken();
  int NonTerminalId;
  int TerminalId;
  int RuleId;
  char c;
  BOOL WaitForWaitStatementBooleanExpression = FALSE;
  InputIdx = 0;
  CurrentLine = 0;
  CurrentLineIdx = 0;
  PTOKEN EndToken = NewToken(END_OF_STACK, "$");
  PTOKEN StartToken = NewToken(NON_TERMINAL, START_VARIABLE);
  Push(Stack, EndToken);
  Push(Stack, StartToken);
  c = sgetc(str);
  PTOKEN CurrentIn = Scan(str, &c);
  if (CurrentIn->Type == UNKNOWN) {
    Error = SCRIPT_ENGINE_ERROR_SYNTAX;
    ErrorMessage = HandleError(&Error, str);
    CodeBuffer->Message = ErrorMessage;
    RemoveTokenList(Stack);
    RemoveTokenList(MatchedStack);
    RemoveToken(&CurrentIn);
    return (PVOID)CodeBuffer;
  }
  PSYMBOL TempSymbol = NewSymbol();
  TempSymbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
  TempSymbol->Value = FUNC_ADD;
  PushSymbol(CodeBuffer, TempSymbol);
  RemoveSymbol(&TempSymbol);
  TempSymbol = NewSymbol();
  TempSymbol->Type = SYMBOL_NUM_TYPE;
  TempSymbol->Value = 0xffffffffffffffff;
  PushSymbol(CodeBuffer, TempSymbol);
  RemoveSymbol(&TempSymbol);
  TempSymbol = NewSymbol();
  TempSymbol->Type = SYMBOL_STACK_INDEX_TYPE;
  TempSymbol->Value = 0;
  PushSymbol(CodeBuffer, TempSymbol);
  RemoveSymbol(&TempSymbol);
  TempSymbol = NewSymbol();
  TempSymbol->Type = SYMBOL_STACK_INDEX_TYPE;
  TempSymbol->Value = 0;
  PushSymbol(CodeBuffer, TempSymbol);
  RemoveSymbol(&TempSymbol);
  do {
    RemoveToken(&TopToken);
    TopToken = Pop(Stack);
#ifdef _SCRIPT_ENGINE_LL1_DBG_EN printf("\nTop Token :\n");
    PrintToken(TopToken);
    printf("\nCurrent Input :\n");
    PrintToken(CurrentIn);
    printf("\n");
#endif if (TopToken->Type == NON_TERMINAL) {
    if (!strcmp(TopToken->Value, "BOOLEAN_EXPRESSION")) {
      UINT64 BooleanExpressionSize = BooleanExpressionExtractEnd(
          str, &WaitForWaitStatementBooleanExpression, CurrentIn);
      ScriptEngineBooleanExpresssionParse(BooleanExpressionSize, CurrentIn,
                                          MatchedStack, CodeBuffer, str, &c,
                                          &Error);
      if (Error != SCRIPT_ENGINE_ERROR_FREE) {
        break;
      }
      RemoveToken(&CurrentIn);
      CurrentIn = Scan(str, &c);
      if (CurrentIn->Type == UNKNOWN) {
        Error = SCRIPT_ENGINE_ERROR_UNKNOWN_TOKEN;
        break;
      }
      RemoveToken(&CurrentIn);
      CurrentIn = Scan(str, &c);
      if (CurrentIn->Type == UNKNOWN) {
        Error = SCRIPT_ENGINE_ERROR_UNKNOWN_TOKEN;
        break;
      }
      RemoveToken(&TopToken);
      TopToken = Pop(Stack);
    } else {
      NonTerminalId = GetNonTerminalId(TopToken);
      if (NonTerminalId == INVALID) {
        Error = SCRIPT_ENGINE_ERROR_SYNTAX;
        break;
      }
      TerminalId = GetTerminalId(CurrentIn);
      if (TerminalId == INVALID) {
        Error = SCRIPT_ENGINE_ERROR_SYNTAX;
        break;
      }
      RuleId = ParseTable[NonTerminalId][TerminalId];
      if (RuleId == INVALID) {
        Error = SCRIPT_ENGINE_ERROR_SYNTAX;
        break;
      }
      for (int i = RhsSize[RuleId] - 1; i >= 0; i--) {
        PTOKEN Token = (PTOKEN)&Rhs[RuleId][i];
        if (Token->Type == EPSILON) break;
        PTOKEN DuplicatedToken = CopyToken(Token);
        Push(Stack, DuplicatedToken);
      }
    }
  }
  else if (TopToken->Type == SEMANTIC_RULE) {
    if (!strcmp(TopToken->Value, "@PUSH")) {
      RemoveToken(&TopToken);
      TopToken = Pop(Stack);
      Push(MatchedStack, CurrentIn);
      CurrentIn = Scan(str, &c);
      if (CurrentIn->Type == UNKNOWN) {
        Error = SCRIPT_ENGINE_ERROR_SYNTAX;
        break;
      }
    } else {
      if (!strcmp(TopToken->Value, "@START_OF_FOR")) {
        WaitForWaitStatementBooleanExpression = TRUE;
      }
      CodeGen(MatchedStack, CodeBuffer, TopToken, &Error);
      if (Error != SCRIPT_ENGINE_ERROR_FREE) {
        break;
      }
    }
  }
  else {
    if (!IsEqual(TopToken, CurrentIn)) {
      Error = SCRIPT_ENGINE_ERROR_SYNTAX;
      break;
    } else {
      RemoveToken(&CurrentIn);
      CurrentIn = Scan(str, &c);
      if (CurrentIn->Type == UNKNOWN) {
        Error = SCRIPT_ENGINE_ERROR_SYNTAX;
        break;
      }
    }
  }
#ifdef _SCRIPT_ENGINE_LL1_DBG_EN printf("Stack: \n");
  PrintTokenList(Stack);
  printf("\n");
#endif
}
while (TopToken->Type != END_OF_STACK);
if (Error != SCRIPT_ENGINE_ERROR_FREE) {
  ErrorMessage = HandleError(&Error, str);
} else {
  ErrorMessage = NULL;
  PSYMBOL Symbol;
  for (UINT64 i = 0; i < CodeBuffer->Pointer; i++) {
    Symbol = CodeBuffer->Head + i;
    if (Symbol->Type == SYMBOL_LOCAL_ID_TYPE) {
      Symbol->Type = SYMBOL_TEMP_TYPE;
      Symbol->Value += UserDefinedFunctionHead->MaxTempNumber;
    } else if (Symbol->Type == SYMBOL_VARIABLE_COUNT_TYPE) {
      UINT64 VariableCount = Symbol->Value;
      for (UINT64 j = 0; j < VariableCount; j++) {
        Symbol = CodeBuffer->Head + i + j + 1;
        if ((Symbol->Type & 0x7fffffff) == SYMBOL_LOCAL_ID_TYPE) {
          Symbol->Type = SYMBOL_TEMP_TYPE | (Symbol->Type & 0xffffffff00000000);
          Symbol->Value += UserDefinedFunctionHead->MaxTempNumber;
        }
      }
      i += VariableCount;
    }
  }
  Symbol = CodeBuffer->Head + 1;
  Symbol->Value = CurrentUserDefinedFunction->MaxTempNumber +
                  CurrentUserDefinedFunction->LocalVariableNumber;
}
CodeBuffer->Message = ErrorMessage;
if (Stack) RemoveTokenList(Stack);
if (MatchedStack) RemoveTokenList(MatchedStack);
if (UserDefinedFunctionHead) {
  PUSER_DEFINED_FUNCTION_NODE Node = UserDefinedFunctionHead;
  while (Node) {
    if (Node->Name) free(Node->Name);
    if (Node->IdTable) RemoveTokenList((PTOKEN_LIST)Node->IdTable);
    if (Node->FunctionParameterIdTable)
      RemoveTokenList((PTOKEN_LIST)Node->FunctionParameterIdTable);
    if (Node->TempMap) free(Node->TempMap);
    PUSER_DEFINED_FUNCTION_NODE Temp = Node;
    Node = Node->NextNode;
    free(Temp);
  }
  UserDefinedFunctionHead = 0;
}
if (CurrentIn) RemoveToken(&CurrentIn);
if (TopToken) RemoveToken(&TopToken);
return (PVOID)CodeBuffer;
}

void CodeGen(PTOKEN_LIST MatchedStack, PSYMBOL_BUFFER CodeBuffer,
             PTOKEN Operator, PSCRIPT_ENGINE_ERROR_TYPE Error) {
  PTOKEN Op0 = NULL;
  PTOKEN Op1 = NULL;
  PTOKEN Op2 = NULL;
  PTOKEN Temp = NULL;
  PSYMBOL OperatorSymbol = NULL;
  PSYMBOL Op0Symbol = NULL;
  PSYMBOL Op1Symbol = NULL;
  PSYMBOL Op2Symbol = NULL;
  PSYMBOL TempSymbol = NULL;
  VARIABLE_TYPE *VariableType = NULL;
  OperatorSymbol = ToSymbol(Operator, Error);
#ifdef _SCRIPT_ENGINE_CODEGEN_DBG_EN printf("Operator :\n");
  PrintToken(Operator);
  printf("\n");
  printf("Semantic Stack:\n");
  PrintTokenList(MatchedStack);
  printf("\n");
  printf("Code Buffer:\n");
  PrintSymbolBuffer((PVOID)CodeBuffer);
  printf(".\n.\n.\n\n");
#endif while (TRUE) {
  if (!strcmp(Operator->Value, "@START_OF_USER_DEFINED_FUNCTION")) {
    Op0 = Pop(MatchedStack);
    VariableType = HandleType(MatchedStack);
    if (VariableType->Kind == TY_UNKNOWN) {
      *Error = SCRIPT_ENGINE_ERROR_UNDEFINED_VARIABLE_TYPE;
      break;
    }
    PSYMBOL JumpInstruction = NewSymbol();
    JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    JumpInstruction->Value = FUNC_JMP;
    PushSymbol(CodeBuffer, JumpInstruction);
    RemoveSymbol(&JumpInstruction);
    PSYMBOL JumpAddressSymbol = NewSymbol();
    JumpAddressSymbol->Type = SYMBOL_NUM_TYPE;
    JumpAddressSymbol->Value = 0xffffffffffffffff;
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    RemoveSymbol(&JumpAddressSymbol);
    PUSER_DEFINED_FUNCTION_NODE Node = UserDefinedFunctionHead;
    while (Node->NextNode) {
      Node = Node->NextNode;
    }
    Node->NextNode = malloc(sizeof(USER_DEFINED_FUNCTION_NODE));
    RtlZeroMemory(Node->NextNode, sizeof(USER_DEFINED_FUNCTION_NODE));
    CurrentUserDefinedFunction = Node->NextNode;
    CurrentUserDefinedFunction->Name = _strdup(Op0->Value);
    CurrentUserDefinedFunction->Address = CodeBuffer->Pointer;
    CurrentUserDefinedFunction->VariableType = (long long unsigned)VariableType;
    CurrentUserDefinedFunction->IdTable = (unsigned long long)NewTokenList();
    CurrentUserDefinedFunction->FunctionParameterIdTable =
        (unsigned long long)NewTokenList();
    CurrentUserDefinedFunction->TempMap = calloc(MAX_TEMP_COUNT, 1);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    TempSymbol->Value = FUNC_PUSH;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_STACK_BASE_INDEX_TYPE;
    TempSymbol->Value = 0;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    TempSymbol->Value = FUNC_MOV;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_STACK_INDEX_TYPE;
    TempSymbol->Value = 0;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_STACK_BASE_INDEX_TYPE;
    TempSymbol->Value = 0;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    TempSymbol->Value = FUNC_ADD;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_NUM_TYPE;
    TempSymbol->Value = 0xffffffffffffffff;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_STACK_INDEX_TYPE;
    TempSymbol->Value = 0;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_STACK_INDEX_TYPE;
    TempSymbol->Value = 0;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
  } else if (!strcmp(Operator->Value, "@FUNCTION_PARAMETER")) {
    Op0 = Pop(MatchedStack);
    VariableType = HandleType(MatchedStack);
    if (VariableType->Kind == TY_UNKNOWN) {
      *Error = SCRIPT_ENGINE_ERROR_UNDEFINED_VARIABLE_TYPE;
      break;
    }
    NewFunctionParameterIdentifier(Op0);
    CurrentUserDefinedFunction->ParameterNumber++;
  } else if (!strcmp(Operator->Value, "@END_OF_USER_DEFINED_FUNCTION")) {
    UINT64 CurrentPointer = CodeBuffer->Pointer;
    PSYMBOL Symbol = NULL;
    if (!CurrentUserDefinedFunction) {
      *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
      break;
    }
    for (UINT64 i = CurrentUserDefinedFunction->Address; i < CurrentPointer;
         i++) {
      Symbol = CodeBuffer->Head + i;
      if (Symbol->Type == SYMBOL_LOCAL_ID_TYPE) {
        Symbol->Type = SYMBOL_TEMP_TYPE;
        Symbol->Value += CurrentUserDefinedFunction->MaxTempNumber;
      } else if (Symbol->Type == SYMBOL_VARIABLE_COUNT_TYPE) {
        UINT64 VariableCount = Symbol->Value;
        for (UINT64 j = 0; j < VariableCount; j++) {
          Symbol = CodeBuffer->Head + i + j + 1;
          if ((Symbol->Type & 0x7fffffff) == SYMBOL_LOCAL_ID_TYPE) {
            Symbol->Type =
                SYMBOL_TEMP_TYPE | (Symbol->Type & 0xffffffff00000000);
            Symbol->Value += CurrentUserDefinedFunction->MaxTempNumber;
          }
        }
        i += VariableCount;
      }
    }
    Symbol = CodeBuffer->Head + CurrentUserDefinedFunction->Address + 6;
    Symbol->Value = CurrentUserDefinedFunction->MaxTempNumber +
                    CurrentUserDefinedFunction->LocalVariableNumber;
    for (UINT64 i = CurrentUserDefinedFunction->Address; i < CurrentPointer;
         i++) {
      Symbol = CodeBuffer->Head + i;
      if (Symbol->Type == SYMBOL_SEMANTIC_RULE_TYPE &&
          Symbol->Value == FUNC_JMP &&
          (CodeBuffer->Head + i + 1)->Value == 0xfffffffffffffff0) {
        (CodeBuffer->Head + i + 1)->Value = CurrentPointer;
        i++;
      }
    }
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    TempSymbol->Value = FUNC_MOV;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_STACK_BASE_INDEX_TYPE;
    TempSymbol->Value = 0;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_STACK_INDEX_TYPE;
    TempSymbol->Value = 0;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    TempSymbol->Value = FUNC_POP;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_STACK_BASE_INDEX_TYPE;
    TempSymbol->Value = 0;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    TempSymbol = NewSymbol();
    TempSymbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    TempSymbol->Value = FUNC_RET;
    PushSymbol(CodeBuffer, TempSymbol);
    RemoveSymbol(&TempSymbol);
    Symbol = CodeBuffer->Head + CurrentUserDefinedFunction->Address - 1;
    Symbol->Value = CodeBuffer->Pointer;
    CurrentUserDefinedFunction = UserDefinedFunctionHead;
  } else if (!strcmp(Operator->Value,
                     "@RETURN_OF_USER_DEFINED_FUNCTION_WITHOUT_VALUE")) {
    if (!CurrentUserDefinedFunction) {
      *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
      break;
    }
    if (((VARIABLE_TYPE *)CurrentUserDefinedFunction->VariableType)->Kind !=
        TY_VOID) {
      *Error = SCRIPT_ENGINE_ERROR_NON_VOID_FUNCTION_NOT_RETURNING_VALUE;
      break;
    }
    PSYMBOL Symbol = NewSymbol();
    Symbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    Symbol->Value = FUNC_JMP;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_NUM_TYPE;
    Symbol->Value = 0xfffffffffffffff0;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
  } else if (!strcmp(Operator->Value,
                     "@RETURN_OF_USER_DEFINED_FUNCTION_WITH_VALUE")) {
    if (!CurrentUserDefinedFunction) {
      *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
      break;
    }
    if (((VARIABLE_TYPE *)CurrentUserDefinedFunction->VariableType)->Kind ==
        TY_VOID) {
      *Error = SCRIPT_ENGINE_ERROR_VOID_FUNCTION_RETURNING_VALUE;
      break;
    }
    PSYMBOL Symbol = NewSymbol();
    Symbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    Symbol->Value = FUNC_MOV;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    FreeTemp(Op0);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_RETURN_VALUE_TYPE;
    Symbol->Value = 0;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    Symbol->Value = FUNC_JMP;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_NUM_TYPE;
    Symbol->Value = 0xfffffffffffffff0;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (
      !strcmp(
          Operator->Value,
          "@END_OF_CALLING_USER_DEFINED_FUNCTION_WITHOUT_RETURNING_VALUE") ||
      !strcmp(Operator->Value,
              "@END_OF_CALLING_USER_DEFINED_FUNCTION_WITH_RETURNING_VALUE")) {
    PSYMBOL Symbol = NULL;
    PSYMBOL TempSymbol = NULL;
    int VariableNum = 0;
    PTOKEN FunctionToken = NULL;
    while (MatchedStack->Pointer > 0) {
      FunctionToken = Pop(MatchedStack);
      if (FunctionToken->Type == FUNCTION_ID) {
        break;
      } else {
        VariableNum++;
        Symbol = NewSymbol();
        Symbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
        Symbol->Value = FUNC_PUSH;
        PushSymbol(CodeBuffer, Symbol);
        RemoveSymbol(&Symbol);
        Symbol = ToSymbol(FunctionToken, Error);
        PushSymbol(CodeBuffer, Symbol);
        RemoveSymbol(&Symbol);
        RemoveToken(&FunctionToken);
      }
    }
    PUSER_DEFINED_FUNCTION_NODE Node =
        GetUserDefinedFunctionNode(FunctionToken);
    if (!Node) {
      *Error = SCRIPT_ENGINE_ERROR_UNDEFINED_FUNCTION;
      break;
    }
    if (VariableNum != Node->ParameterNumber) {
      *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
      break;
    }
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    Symbol->Value = FUNC_CALL;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_NUM_TYPE;
    Symbol->Value = Node->Address;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    Symbol->Value = FUNC_SUB;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_NUM_TYPE;
    Symbol->Value = Node->ParameterNumber;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_STACK_INDEX_TYPE;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    Symbol = NewSymbol();
    Symbol->Type = SYMBOL_STACK_INDEX_TYPE;
    PushSymbol(CodeBuffer, Symbol);
    RemoveSymbol(&Symbol);
    if (!strcmp(Operator->Value,
                "@END_OF_CALLING_USER_DEFINED_FUNCTION_WITH_RETURNING_VALUE")) {
      if (((VARIABLE_TYPE *)Node->VariableType)->Kind == TY_VOID) {
        *Error = SCRIPT_ENGINE_ERROR_VOID_FUNCTION_RETURNING_VALUE;
        break;
      }
      Temp = NewTemp(Error);
      Push(MatchedStack, Temp);
      Symbol = NewSymbol();
      Symbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
      Symbol->Value = FUNC_MOV;
      PushSymbol(CodeBuffer, Symbol);
      RemoveSymbol(&Symbol);
      Symbol = NewSymbol();
      Symbol->Type = SYMBOL_RETURN_VALUE_TYPE;
      Symbol->Value = 0;
      PushSymbol(CodeBuffer, Symbol);
      RemoveSymbol(&Symbol);
      TempSymbol = ToSymbol(Temp, Error);
      PushSymbol(CodeBuffer, TempSymbol);
      RemoveSymbol(&TempSymbol);
      if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
        break;
      }
    }
    RemoveToken(&FunctionToken);
  } else if (!strcmp(Operator->Value, "@MULTIPLE_ASSIGNMENT")) {
    PSYMBOL Symbol = NewSymbol();
    Symbol->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    Symbol->Value = FUNC_MOV;
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    for (int i = MatchedStack->Pointer; i > 0; i--) {
      Op1 = Top(MatchedStack);
      if (Op1->Type == TEMP || Op1->Type == HEX || Op1->Type == OCTAL ||
          Op1->Type == BINARY || Op1->Type == PSEUDO_REGISTER) {
        *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
        Pop(MatchedStack);
        break;
      } else if (Op1->Type == GLOBAL_UNRESOLVED_ID) {
        PushSymbol(CodeBuffer, Symbol);
        PushSymbol(CodeBuffer, Op0Symbol);
        Op1Symbol = NewSymbol();
        free((void *)Op1Symbol->Value);
        Op1Symbol->Value = NewGlobalIdentifier(Op1);
        SetType(&Op1Symbol->Type, SYMBOL_GLOBAL_ID_TYPE);
        Pop(MatchedStack);
        PushSymbol(CodeBuffer, Op1Symbol);
      } else if (Op1->Type == LOCAL_UNRESOLVED_ID) {
        PushSymbol(CodeBuffer, Symbol);
        PushSymbol(CodeBuffer, Op0Symbol);
        Op1Symbol = NewSymbol();
        free((void *)Op1Symbol->Value);
        Op1Symbol->Value = NewLocalIdentifier(Op1);
        SetType(&Op1Symbol->Type, SYMBOL_LOCAL_ID_TYPE);
        Pop(MatchedStack);
        PushSymbol(CodeBuffer, Op1Symbol);
        RemoveSymbol(&Op1Symbol);
      } else if (Op1->Type == LOCAL_ID || Op1->Type == GLOBAL_ID ||
                 Op1->Type == FUNCTION_PARAMETER_ID || Op1->Type == REGISTER) {
        PushSymbol(CodeBuffer, Symbol);
        PushSymbol(CodeBuffer, Op0Symbol);
        Op1Symbol = ToSymbol(Op1, Error);
        PushSymbol(CodeBuffer, Op1Symbol);
        Pop(MatchedStack);
        RemoveSymbol(&Op1Symbol);
      } else {
        break;
      }
    }
    RemoveSymbol(&Symbol);
    Op1 = 0;
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      FreeTemp(Op0);
      break;
    }
    if (MatchedStack->Pointer > 0) {
      if (Top(MatchedStack)->Type == SCRIPT_VARIABLE_TYPE) {
        VariableType = HandleType(MatchedStack);
        if (VariableType->Kind == TY_UNKNOWN) {
          *Error = SCRIPT_ENGINE_ERROR_UNDEFINED_VARIABLE_TYPE;
          break;
        }
      }
    }
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (!strcmp(Operator->Value, "@MOV")) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    if (Op1->Type == GLOBAL_UNRESOLVED_ID) {
      Op1Symbol = NewSymbol();
      free((void *)Op1Symbol->Value);
      Op1Symbol->Value = NewGlobalIdentifier(Op1);
      SetType(&Op1Symbol->Type, SYMBOL_GLOBAL_ID_TYPE);
    } else if (Op1->Type == LOCAL_UNRESOLVED_ID) {
      Op1Symbol = NewSymbol();
      free((void *)Op1Symbol->Value);
      Op1Symbol->Value = NewLocalIdentifier(Op1);
      SetType(&Op1Symbol->Type, SYMBOL_LOCAL_ID_TYPE);
    } else {
      Op1Symbol = ToSymbol(Op1, Error);
    }
    if (MatchedStack->Pointer > 0) {
      if (Top(MatchedStack)->Type == SCRIPT_VARIABLE_TYPE) {
        VariableType = HandleType(MatchedStack);
        if (VariableType->Kind == TY_UNKNOWN) {
          *Error = SCRIPT_ENGINE_ERROR_UNDEFINED_VARIABLE_TYPE;
          break;
        }
      }
    }
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    FreeTemp(Op0);
    FreeTemp(Op1);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsType2Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsType1Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsType4Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    PSYMBOL_BUFFER TempStack = NewSymbolBuffer();
    UINT32 OperandCount = 0;
    do {
      if (Op1) {
        RemoveToken(&Op1);
      }
      Op1 = Pop(MatchedStack);
      if (Op1->Type != SEMANTIC_RULE) {
        Op1Symbol = ToSymbol(Op1, Error);
        FreeTemp(Op1);
        PushSymbol(TempStack, Op1Symbol);
        RemoveSymbol(&Op1Symbol);
        OperandCount++;
        if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
          RemoveSymbolBuffer((PVOID)TempStack);
          break;
        }
      }
    } while (
        !(Op1->Type == SEMANTIC_RULE && !strcmp(Op1->Value, "@VARGSTART")));
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    FreeTemp(Op0);
    char *Format = Op0->Value;
    PSYMBOL OperandCountSymbol = NewSymbol();
    OperandCountSymbol->Type = SYMBOL_VARIABLE_COUNT_TYPE;
    OperandCountSymbol->Value = OperandCount;
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, OperandCountSymbol);
    RemoveSymbol(&OperandCountSymbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      RemoveSymbolBuffer((PVOID)TempStack);
      break;
    }
    unsigned int FirstArgPointer = CodeBuffer->Pointer;
    PSYMBOL Symbol;
    unsigned int ArgCount = TempStack->Pointer;
    for (int i = TempStack->Pointer - 1; i >= 0; i--) {
      Symbol = TempStack->Head + i;
      PushSymbol(CodeBuffer, Symbol);
    }
    PSYMBOL FirstArg =
        (PSYMBOL)((unsigned long long)CodeBuffer->Head +
                  (unsigned long long)(FirstArgPointer * sizeof(SYMBOL)));
    RemoveSymbolBuffer((PVOID)TempStack);
    UINT32 i = 0;
    char *Str = Format;
    do {
      if (*Str == '%') {
        CHAR Temp = *(Str + 1);
        if (Temp == 'd' || Temp == 'i' || Temp == 'u' || Temp == 'o' ||
            Temp == 'x' || Temp == 'c' || Temp == 'p' || Temp == 's' ||
            !strncmp(Str, "%ws", 3) || !strncmp(Str, "%ls", 3) ||
            !strncmp(Str, "%ld", 3) || !strncmp(Str, "%li", 3) ||
            !strncmp(Str, "%lu", 3) || !strncmp(Str, "%lo", 3) ||
            !strncmp(Str, "%lx", 3) || !strncmp(Str, "%hd", 3) ||
            !strncmp(Str, "%hi", 3) || !strncmp(Str, "%hu", 3) ||
            !strncmp(Str, "%ho", 3) || !strncmp(Str, "%hx", 3) ||
            !strncmp(Str, "%lld", 4) || !strncmp(Str, "%lli", 4) ||
            !strncmp(Str, "%llu", 4) || !strncmp(Str, "%llo", 4) ||
            !strncmp(Str, "%llx", 4)) {
          if (i < ArgCount) {
            Symbol = FirstArg + i;
          } else {
            *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
            break;
          }
          Symbol->Type &= 0xffffffff;
          Symbol->Type |= (UINT64)(Str - Format - 1) << 32;
          i++;
        }
      }
      Str++;
    } while (*Str);
    if (i != ArgCount) {
      *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
    }
    if (*Error == SCRIPT_ENGINE_ERROR_SYNTAX) {
      break;
    }
  } else if (IsType5Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
  } else if (!strcmp(Operator->Value, "@IGNORE_LVALUE")) {
    Op0 = Pop(MatchedStack);
  } else if (IsType6Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    FreeTemp(Op1);
  } else if (IsType7Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    FreeTemp(Op0);
    FreeTemp(Op1);
  } else if (IsType8Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    Op2 = Pop(MatchedStack);
    Op2Symbol = ToSymbol(Op2, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    PushSymbol(CodeBuffer, Op2Symbol);
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op2);
    FreeTemp(Op0);
    FreeTemp(Op1);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsType14Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    Op2 = Pop(MatchedStack);
    Op2Symbol = ToSymbol(Op2, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    PushSymbol(CodeBuffer, Op2Symbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    FreeTemp(Op0);
    FreeTemp(Op1);
    FreeTemp(Op2);
  } else if (IsAssignmentOperator(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsTwoOperandOperator(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    FreeTemp(Op1);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsOneOperandOperator(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (!strcmp(Operator->Value, "@VARGSTART")) {
    PTOKEN OperatorCopy = CopyToken(Operator);
    Push(MatchedStack, OperatorCopy);
  } else if (!strcmp(Operator->Value, "@START_OF_IF")) {
    PTOKEN OperatorCopy = CopyToken(Operator);
    Push(MatchedStack, OperatorCopy);
  } else if (!strcmp(Operator->Value, "@JZ")) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    PSYMBOL JumpAddressSymbol = NewSymbol();
    JumpAddressSymbol->Type = SYMBOL_NUM_TYPE;
    JumpAddressSymbol->Value = 0xffffffffffffffff;
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    RemoveSymbol(&JumpAddressSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    char str[20] = {0};
    sprintf(str, "%llu", (UINT64)CodeBuffer->Pointer);
    PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
    Push(MatchedStack, CurrentAddressToken);
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (!strcmp(Operator->Value, "@JMP_TO_END_AND_JZCOMPLETED")) {
    UINT64 CurrentPointer = CodeBuffer->Pointer;
    PTOKEN JumpSemanticAddressToken = Pop(MatchedStack);
    UINT64 JumpSemanticAddress = DecimalToInt(JumpSemanticAddressToken->Value);
    PSYMBOL JumpAddressSymbol =
        (PSYMBOL)(CodeBuffer->Head + JumpSemanticAddress - 2);
    JumpAddressSymbol->Value = CurrentPointer + 2;
    RemoveToken(&JumpSemanticAddressToken);
    PSYMBOL JumpInstruction = NewSymbol();
    JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    JumpInstruction->Value = FUNC_JMP;
    PushSymbol(CodeBuffer, JumpInstruction);
    RemoveSymbol(&JumpInstruction);
    JumpAddressSymbol = NewSymbol();
    JumpAddressSymbol->Type = SYMBOL_NUM_TYPE;
    JumpAddressSymbol->Value = 0xffffffffffffffff;
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    RemoveSymbol(&JumpAddressSymbol);
    char str[20] = {0};
    sprintf(str, "%llu", CurrentPointer);
    PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
    Push(MatchedStack, CurrentAddressToken);
  } else if (!strcmp(Operator->Value, "@END_OF_IF")) {
    UINT64 CurrentPointer = CodeBuffer->Pointer;
    PTOKEN JumpSemanticAddressToken = Pop(MatchedStack);
    PSYMBOL JumpAddressSymbol;
    while (strcmp(JumpSemanticAddressToken->Value, "@START_OF_IF")) {
      UINT64 JumpSemanticAddress =
          DecimalToInt(JumpSemanticAddressToken->Value);
      JumpAddressSymbol = (PSYMBOL)(CodeBuffer->Head + JumpSemanticAddress + 1);
      JumpAddressSymbol->Value = CurrentPointer;
      RemoveToken(&JumpSemanticAddressToken);
      JumpSemanticAddressToken = Pop(MatchedStack);
    }
    RemoveToken(&JumpSemanticAddressToken);
  } else if (!strcmp(Operator->Value, "@START_OF_WHILE")) {
    PTOKEN OperatorCopy = CopyToken(Operator);
    Push(MatchedStack, OperatorCopy);
    char str[20] = {0};
    sprintf(str, "%llu", (UINT64)CodeBuffer->Pointer);
    PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
    Push(MatchedStack, CurrentAddressToken);
  } else if (!strcmp(Operator->Value, "@START_OF_WHILE_COMMANDS")) {
    UINT64 CurrentPointer = CodeBuffer->Pointer;
    PTOKEN JzToken = NewToken(SEMANTIC_RULE, "@JZ");
    RemoveSymbol(&OperatorSymbol);
    OperatorSymbol = ToSymbol(JzToken, Error);
    RemoveToken(&JzToken);
    PSYMBOL JumpAddressSymbol = NewSymbol();
    JumpAddressSymbol->Type = SYMBOL_NUM_TYPE;
    JumpAddressSymbol->Value = 0xffffffffffffffff;
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    PTOKEN StartOfWhileToken = Pop(MatchedStack);
    char str[20];
    sprintf(str, "%llu", CurrentPointer + 1);
    PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
    Push(MatchedStack, CurrentAddressToken);
    Push(MatchedStack, StartOfWhileToken);
    PushSymbol(CodeBuffer, OperatorSymbol);
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    PushSymbol(CodeBuffer, Op0Symbol);
    RemoveSymbol(&JumpAddressSymbol);
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (!strcmp(Operator->Value, "@END_OF_WHILE")) {
    PSYMBOL JumpInstruction = NewSymbol();
    JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    JumpInstruction->Value = FUNC_JMP;
    PushSymbol(CodeBuffer, JumpInstruction);
    RemoveSymbol(&JumpInstruction);
    PTOKEN JumpAddressToken = Pop(MatchedStack);
    UINT64 JumpAddress = DecimalToInt(JumpAddressToken->Value);
    PSYMBOL JumpAddressSymbol = ToSymbol(JumpAddressToken, Error);
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    RemoveSymbol(&JumpAddressSymbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    UINT64 CurrentPointer = CodeBuffer->Pointer;
    do {
      RemoveToken(&JumpAddressToken);
      JumpAddressToken = Pop(MatchedStack);
      if (!strcmp(JumpAddressToken->Value, "@START_OF_WHILE")) {
        break;
      }
      JumpAddress = DecimalToInt(JumpAddressToken->Value);
      JumpAddressSymbol = (PSYMBOL)(CodeBuffer->Head + JumpAddress);
      JumpAddressSymbol->Value = CurrentPointer;
    } while (TRUE);
    RemoveToken(&JumpAddressToken);
  } else if (!strcmp(Operator->Value, "@START_OF_DO_WHILE")) {
    PTOKEN OperatorCopy = CopyToken(Operator);
    Push(MatchedStack, OperatorCopy);
    char str[20];
    sprintf(str, "%llu", (UINT64)CodeBuffer->Pointer);
    PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
    Push(MatchedStack, CurrentAddressToken);
  } else if (!strcmp(Operator->Value, "@END_OF_DO_WHILE")) {
    PSYMBOL JumpInstruction = NewSymbol();
    JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    JumpInstruction->Value = FUNC_JNZ;
    PushSymbol(CodeBuffer, JumpInstruction);
    RemoveSymbol(&JumpInstruction);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    PTOKEN JumpAddressToken = Pop(MatchedStack);
    UINT64 JumpAddress = DecimalToInt(JumpAddressToken->Value);
    PSYMBOL JumpAddressSymbol = ToSymbol(JumpAddressToken, Error);
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    PushSymbol(CodeBuffer, Op0Symbol);
    RemoveSymbol(&JumpAddressSymbol);
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    UINT64 CurrentPointer = CodeBuffer->Pointer;
    do {
      RemoveToken(&JumpAddressToken);
      JumpAddressToken = Pop(MatchedStack);
      if (!strcmp(JumpAddressToken->Value, "@START_OF_DO_WHILE")) {
        break;
      }
      JumpAddress = DecimalToInt(JumpAddressToken->Value);
#ifdef _SCRIPT_ENGINE_LL1_DBG_EN printf("Jz Jump Address = %lld\n", \
                                        JumpAddress);
#endif
      JumpAddressSymbol = (PSYMBOL)(CodeBuffer->Head + JumpAddress);
      JumpAddressSymbol->Value = CurrentPointer;
    } while (TRUE);
    RemoveToken(&JumpAddressToken);
  } else if (!strcmp(Operator->Value, "@START_OF_FOR")) {
    PTOKEN OperatorCopy = CopyToken(Operator);
    Push(MatchedStack, OperatorCopy);
    char str[20] = {0};
    sprintf(str, "%llu", (UINT64)CodeBuffer->Pointer);
    PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
    Push(MatchedStack, CurrentAddressToken);
  } else if (!strcmp(Operator->Value, "@FOR_INC_DEC")) {
    PSYMBOL JnzInstruction = NewSymbol();
    JnzInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    JnzInstruction->Value = FUNC_JZ;
    PushSymbol(CodeBuffer, JnzInstruction);
    RemoveSymbol(&JnzInstruction);
    PSYMBOL JnzAddressSymbol = NewSymbol();
    JnzAddressSymbol->Type = SYMBOL_NUM_TYPE;
    JnzAddressSymbol->Value = 0xffffffffffffffff;
    PushSymbol(CodeBuffer, JnzAddressSymbol);
    RemoveSymbol(&JnzAddressSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    PSYMBOL JumpInstruction = NewSymbol();
    JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    JumpInstruction->Value = FUNC_JMP;
    PushSymbol(CodeBuffer, JumpInstruction);
    RemoveSymbol(&JumpInstruction);
    PSYMBOL JumpAddressSymbol = NewSymbol();
    JumpAddressSymbol->Type = SYMBOL_NUM_TYPE;
    JumpAddressSymbol->Value = 0xffffffffffffffff;
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    RemoveSymbol(&JumpAddressSymbol);
    PTOKEN StartOfForAddressToken = Pop(MatchedStack);
    char str[20] = {0};
    sprintf(str, "%llu", (UINT64)CodeBuffer->Pointer);
    PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
    Push(MatchedStack, CurrentAddressToken);
    Push(MatchedStack, StartOfForAddressToken);
  } else if (!strcmp(Operator->Value, "@START_OF_FOR_COMMANDS")) {
    PSYMBOL JumpInstruction = NewSymbol();
    JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    JumpInstruction->Value = FUNC_JMP;
    PushSymbol(CodeBuffer, JumpInstruction);
    RemoveSymbol(&JumpInstruction);
    PTOKEN JumpAddressToken = Pop(MatchedStack);
    UINT64 JumpAddress = DecimalToInt(JumpAddressToken->Value);
    PSYMBOL JumpAddressSymbol = ToSymbol(JumpAddressToken, Error);
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    RemoveToken(&JumpAddressToken);
    RemoveSymbol(&JumpAddressSymbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    UINT64 CurrentPointer = CodeBuffer->Pointer;
    JumpAddressToken = Pop(MatchedStack);
    JumpAddress = DecimalToInt(JumpAddressToken->Value);
    JumpAddressSymbol = (PSYMBOL)(CodeBuffer->Head + JumpAddress - 1);
    JumpAddressSymbol->Value = CurrentPointer;
    char str[20] = {0};
    sprintf(str, "%llu", JumpAddress - 4);
    PTOKEN JzAddressToken = NewToken(DECIMAL, str);
    Push(MatchedStack, JzAddressToken);
    PTOKEN IncDecToken = NewToken(SEMANTIC_RULE, "@INC_DEC");
    Push(MatchedStack, IncDecToken);
    Push(MatchedStack, JumpAddressToken);
  } else if (!strcmp(Operator->Value, "@END_OF_FOR")) {
    PSYMBOL JumpInstruction = NewSymbol();
    JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
    JumpInstruction->Value = FUNC_JMP;
    PushSymbol(CodeBuffer, JumpInstruction);
    RemoveSymbol(&JumpInstruction);
    PTOKEN JumpAddressToken = Pop(MatchedStack);
    UINT64 JumpAddress = DecimalToInt(JumpAddressToken->Value);
    PSYMBOL JumpAddressSymbol = ToSymbol(JumpAddressToken, Error);
    PushSymbol(CodeBuffer, JumpAddressSymbol);
    RemoveSymbol(&JumpAddressSymbol);
    RemoveToken(&JumpAddressToken);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    JumpAddressToken = Pop(MatchedStack);
    UINT64 CurrentPointer = CodeBuffer->Pointer;
    do {
      RemoveToken(&JumpAddressToken);
      JumpAddressToken = Pop(MatchedStack);
      if (!strcmp(JumpAddressToken->Value, "@START_OF_FOR")) {
        break;
      } else {
        JumpAddress = DecimalToInt(JumpAddressToken->Value);
        JumpAddressSymbol = (PSYMBOL)(CodeBuffer->Head + JumpAddress);
        JumpAddressSymbol->Value = CurrentPointer;
      }
    } while (TRUE);
    RemoveToken(&JumpAddressToken);
  } else if (!strcmp(Operator->Value, "@BREAK")) {
    PTOKEN_LIST TempStack = NewTokenList();
    PTOKEN TempToken;
    do {
      TempToken = Pop(MatchedStack);
      if ((!strcmp(TempToken->Value, "@START_OF_FOR")) ||
          (!strcmp(TempToken->Value, "@START_OF_WHILE")) ||
          (!strcmp(TempToken->Value, "@START_OF_DO_WHILE"))) {
        Push(MatchedStack, TempToken);
        UINT64 CurrentPointer = CodeBuffer->Pointer + 1;
        char str[20];
        sprintf(str, "%llu", CurrentPointer);
        PTOKEN CurrentAddressToken = NewToken(DECIMAL, str);
        Push(MatchedStack, CurrentAddressToken);
        PSYMBOL JumpInstruction = NewSymbol();
        JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
        JumpInstruction->Value = FUNC_JMP;
        PushSymbol(CodeBuffer, JumpInstruction);
        RemoveSymbol(&JumpInstruction);
        PSYMBOL JumpAddressSymbol = NewSymbol();
        JumpAddressSymbol->Type = SYMBOL_NUM_TYPE;
        JumpAddressSymbol->Value = 0xffffffffffffffff;
        PushSymbol(CodeBuffer, JumpAddressSymbol);
        RemoveSymbol(&JumpAddressSymbol);
        do {
          TempToken = Pop(TempStack);
          Push(MatchedStack, TempToken);
        } while (TempStack->Pointer != 0);
        break;
      } else if (MatchedStack->Pointer == 0) {
        *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
        RemoveToken(&TempToken);
        break;
      } else {
        Push(TempStack, TempToken);
      }
    } while (TRUE);
    RemoveTokenList(TempStack);
  } else if (!strcmp(Operator->Value, "@CONTINUE")) {
    PTOKEN_LIST TempStack = NewTokenList();
    PTOKEN TempToken;
    do {
      TempToken = Pop(MatchedStack);
      if (!strcmp(TempToken->Value, "@INC_DEC")) {
        Push(MatchedStack, TempToken);
        PSYMBOL JumpInstruction = NewSymbol();
        JumpInstruction->Type = SYMBOL_SEMANTIC_RULE_TYPE;
        JumpInstruction->Value = FUNC_JMP;
        PushSymbol(CodeBuffer, JumpInstruction);
        RemoveSymbol(&JumpInstruction);
        TempToken = Pop(TempStack);
        Push(MatchedStack, TempToken);
        PSYMBOL JumpAddressSymbol = NewSymbol();
        JumpAddressSymbol->Type = SYMBOL_NUM_TYPE;
        JumpAddressSymbol->Value = DecimalToInt(TempToken->Value);
        PushSymbol(CodeBuffer, JumpAddressSymbol);
        RemoveSymbol(&JumpAddressSymbol);
        do {
          TempToken = Pop(TempStack);
          Push(MatchedStack, TempToken);
        } while (TempStack->Pointer != 0);
        break;
      } else if (MatchedStack->Pointer == 0) {
        *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
        break;
      } else {
        Push(TempStack, TempToken);
      }
    } while (TRUE);
    RemoveTokenList(TempStack);
  } else if (IsType9Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsType10Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    FreeTemp(Op1);
  } else if (IsType11Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    Op2 = Pop(MatchedStack);
    Op2Symbol = ToSymbol(Op2, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    PushSymbol(CodeBuffer, Op2Symbol);
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    FreeTemp(Op1);
    FreeTemp(Op2);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsType12Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else if (IsType13Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    FreeTemp(Op1);
  } else if (IsType15Func(Operator)) {
    PushSymbol(CodeBuffer, OperatorSymbol);
    Op0 = Pop(MatchedStack);
    Op0Symbol = ToSymbol(Op0, Error);
    Op1 = Pop(MatchedStack);
    Op1Symbol = ToSymbol(Op1, Error);
    Op2 = Pop(MatchedStack);
    Op2Symbol = ToSymbol(Op2, Error);
    PushSymbol(CodeBuffer, Op0Symbol);
    PushSymbol(CodeBuffer, Op1Symbol);
    PushSymbol(CodeBuffer, Op2Symbol);
    Temp = NewTemp(Error);
    Push(MatchedStack, Temp);
    TempSymbol = ToSymbol(Temp, Error);
    PushSymbol(CodeBuffer, TempSymbol);
    FreeTemp(Op0);
    FreeTemp(Op1);
    FreeTemp(Op2);
    if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
      break;
    }
  } else {
    *Error = SCRIPT_ENGINE_ERROR_UNHANDLED_SEMANTIC_RULE;
  }
  break;
}
#ifdef _SCRIPT_ENGINE_CODEGEN_DBG_EN printf("Semantic Stack:\n");
PrintTokenList(MatchedStack);
printf("\n");
printf("Code Buffer:\n");
PrintSymbolBuffer((PVOID)CodeBuffer);
printf("------------------------------------------\n\n");
#endif
if (Op0) RemoveToken(&Op0);
if (Op1) RemoveToken(&Op1);
if (Op2) RemoveToken(&Op2);
RemoveSymbol(&OperatorSymbol);
if (Op0Symbol) RemoveSymbol(&Op0Symbol);
if (Op1Symbol) RemoveSymbol(&Op1Symbol);
if (Op2Symbol) RemoveSymbol(&Op2Symbol);
if (TempSymbol) RemoveSymbol(&TempSymbol);
return;
}

UINT64 BooleanExpressionExtractEnd(char *str,
                                   BOOL *WaitForWaitStatementBooleanExpression,
                                   PTOKEN CurrentIn) {
  UINT64 BooleanExpressionSize = 0;
  if (*WaitForWaitStatementBooleanExpression) {
    while (str[InputIdx + BooleanExpressionSize - 1] != ';') {
      BooleanExpressionSize += 1;
    }
    *WaitForWaitStatementBooleanExpression = FALSE;
    return InputIdx + BooleanExpressionSize - 1;
  } else {
    int OpenParanthesesCount = 1;
    if (!strcmp(CurrentIn->Value, "(")) {
      OpenParanthesesCount++;
    }
    while (str[InputIdx + BooleanExpressionSize - 1] != '\0') {
      if (str[InputIdx + BooleanExpressionSize - 1] == ')') {
        OpenParanthesesCount--;
        if (OpenParanthesesCount == 0) {
          return InputIdx + BooleanExpressionSize - 1;
        }
      } else if (str[InputIdx + BooleanExpressionSize - 1] == '(') {
        OpenParanthesesCount++;
      }
      BooleanExpressionSize++;
    }
  }
  return -1;
}

void ScriptEngineBooleanExpresssionParse(UINT64 BooleanExpressionSize,
                                         PTOKEN FirstToken,
                                         PTOKEN_LIST MatchedStack,
                                         PSYMBOL_BUFFER CodeBuffer, char *str,
                                         char *c,
                                         PSCRIPT_ENGINE_ERROR_TYPE Error) {
  PTOKEN_LIST Stack = NewTokenList();
  PTOKEN State = NewToken(STATE_ID, "0");
  Push(Stack, State);
#ifdef _SCRIPT_ENGINE_LALR_DBG_EN printf("Boolean Expression: ");
  printf("%s", FirstToken->Value);
  for (int i = InputIdx - 1; i < BooleanExpressionSize; i++) {
    printf("%c", str[i]);
  }
  printf("\n\n");
#endif
  PTOKEN EndToken = NewToken(END_OF_STACK, "$");
  PTOKEN CurrentIn = CopyToken(FirstToken);
  PTOKEN TopToken = NULL;
  PTOKEN Lhs = NULL;
  PTOKEN Temp = NULL;
  PTOKEN Operand = NULL;
  PTOKEN SemanticRule = NULL;
  int Action = INVALID;
  int StateId = 0;
  int Goto = 0;
  int InputPointer = 0;
  int RhsSize = 0;
  unsigned int InputIdxTemp;
  char Ctemp;
  while (1) {
    TopToken = Top(Stack);
    int TerminalId = LalrGetTerminalId(CurrentIn);
    StateId = (int)DecimalToSignedInt(TopToken->Value);
    if (StateId == INVALID || TerminalId < 0) {
      *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
      break;
    }
    Action = LalrActionTable[StateId][TerminalId];
#ifdef _SCRIPT_ENGINE_LALR_DBG_EN printf("Stack :\n");
    PrintTokenList(Stack);
    printf("Action : %d\n\n", Action);
#endif if (Action == LALR_ACCEPT) {
    *Error = SCRIPT_ENGINE_ERROR_FREE;
    break;
  }
  if (Action == INVALID) {
    *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
    break;
  }
  if (Action == 0) {
    *Error = SCRIPT_ENGINE_ERROR_SYNTAX;
    break;
  } else if (Action >= 0) {
    StateId = Action;
    Push(Stack, CurrentIn);
    char buffer[20] = {0};
    sprintf(buffer, "%d", StateId);
    State = NewToken(STATE_ID, buffer);
    Push(Stack, State);
    InputIdxTemp = InputIdx;
    Ctemp = *c;
    CurrentIn = Scan(str, c);
    if (InputIdx - 1 > BooleanExpressionSize) {
      InputIdx = InputIdxTemp;
      *c = Ctemp;
      RemoveToken(&CurrentIn);
      CurrentIn = CopyToken(EndToken);
    }
  } else if (Action < 0) {
    StateId = -Action;
    Lhs = (PTOKEN)&LalrLhs[StateId - 1];
    RhsSize = LalrGetRhsSize(StateId - 1);
    SemanticRule = (PTOKEN)&LalrSemanticRules[StateId - 1];
    for (int i = 0; i < 2 * RhsSize; i++) {
      Temp = Pop(Stack);
      if (SemanticRule->Type == SEMANTIC_RULE &&
          !strcmp(SemanticRule->Value, "@PUSH")) {
        if (LalrIsOperandType(Temp)) {
          Operand = Temp;
          Push(MatchedStack, Operand);
        } else {
          RemoveToken(&Temp);
        }
      } else {
        RemoveToken(&Temp);
      }
    }
    if (SemanticRule->Type == SEMANTIC_RULE) {
      if (!strcmp(SemanticRule->Value, "@PUSH")) {
      } else {
        CodeGen(MatchedStack, CodeBuffer, SemanticRule, Error);
        if (*Error != SCRIPT_ENGINE_ERROR_FREE) {
          break;
        }
      }
    }
    Temp = Top(Stack);
    StateId = (int)DecimalToSignedInt(Temp->Value);
    Goto = LalrGotoTable[StateId][LalrGetNonTerminalId(Lhs)];
    PTOKEN LhsCopy = CopyToken(Lhs);
    char buffer[20] = {0};
    sprintf(buffer, "%d", Goto);
    State = NewToken(STATE_ID, buffer);
    Push(Stack, LhsCopy);
    Push(Stack, State);
  }
}
if (EndToken) RemoveToken(&EndToken);
if (Stack) RemoveTokenList(Stack);
if (CurrentIn) RemoveToken(&CurrentIn);
return;
}

PSYMBOL NewSymbol(void) {
  PSYMBOL Symbol;
  Symbol = (PSYMBOL)malloc(sizeof(SYMBOL));
  if (Symbol == NULL) {
    return NULL;
  }
  Symbol->Value = 0;
  Symbol->Len = 0;
  Symbol->Type = 0;
  return Symbol;
}

PSYMBOL NewStringSymbol(PTOKEN Token) {
  PSYMBOL Symbol;
  int BufferSize = (SIZE_SYMBOL_WITHOUT_LEN + Token->Len) / sizeof(SYMBOL) + 1;
  Symbol = (PSYMBOL)calloc(sizeof(SYMBOL), BufferSize);
  if (Symbol == NULL) {
    return NULL;
  }
  memcpy(&Symbol->Value, Token->Value, Token->Len);
  SetType(&Symbol->Type, SYMBOL_STRING_TYPE);
  Symbol->Len = Token->Len;
  return Symbol;
}

PSYMBOL NewWstringSymbol(PTOKEN Token) {
  PSYMBOL Symbol;
  int BufferSize = (SIZE_SYMBOL_WITHOUT_LEN + Token->Len) / sizeof(SYMBOL) + 1;
  Symbol = (PSYMBOL)malloc(BufferSize * sizeof(SYMBOL));
  if (Symbol == NULL) {
    return NULL;
  }
  memcpy(&Symbol->Value, Token->Value, Token->Len);
  SetType(&Symbol->Type, SYMBOL_WSTRING_TYPE);
  Symbol->Len = Token->Len;
  return Symbol;
}

unsigned int GetSymbolHeapSize(PSYMBOL Symbol) {
  int Temp = (SIZE_SYMBOL_WITHOUT_LEN + (int)Symbol->Len) / sizeof(SYMBOL) + 1;
  return Temp;
}

void RemoveSymbol(PSYMBOL *Symbol) {
  free(*Symbol);
  *Symbol = NULL;
  return;
}

void PrintSymbol(PVOID Symbol) {
  PSYMBOL Sym = (PSYMBOL)Symbol;
  if (Sym->Type & 0xffffffff00000000) {
    printf("Type = @VARGSTART\n");
    return;
  }
  printf("Type = %s, ", SymbolTypeNames[Sym->Type]);
  if (Sym->Type == SYMBOL_SEMANTIC_RULE_TYPE) {
    printf("Value = %s\n", FunctionNames[Sym->Value]);
  } else if (Sym->Type == SYMBOL_STRING_TYPE) {
    printf("Value = %s\n", (char *)&Sym->Value);
  } else if (Sym->Type == SYMBOL_WSTRING_TYPE) {
    printf("Value = %ls\n", (wchar_t *)&Sym->Value);
  } else {
    printf("Value = %lld\n", Sym->Value);
  }
}

PSYMBOL ToSymbol(PTOKEN Token, PSCRIPT_ENGINE_ERROR_TYPE Error) {
  PSYMBOL Symbol = NewSymbol();
  switch (Token->Type) {
    case GLOBAL_ID:
      Symbol->Value = GetGlobalIdentifierVal(Token);
      SetType(&Symbol->Type, SYMBOL_GLOBAL_ID_TYPE);
      return Symbol;
    case LOCAL_ID:
      Symbol->Value = GetLocalIdentifierVal(Token);
      SetType(&Symbol->Type, SYMBOL_LOCAL_ID_TYPE);
      return Symbol;
    case DECIMAL:
      Symbol->Value = DecimalToInt(Token->Value);
      SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
      return Symbol;
    case HEX:
      Symbol->Value = HexToInt(Token->Value);
      SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
      return Symbol;
    case OCTAL:
      Symbol->Value = OctalToInt(Token->Value);
      SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
      return Symbol;
    case BINARY:
      Symbol->Value = BinaryToInt(Token->Value);
      SetType(&Symbol->Type, SYMBOL_NUM_TYPE);
      return Symbol;
    case REGISTER:
      Symbol->Value = RegisterToInt(Token->Value);
      SetType(&Symbol->Type, SYMBOL_REGISTER_TYPE);
      return Symbol;
    case PSEUDO_REGISTER:
      Symbol->Value = PseudoRegToInt(Token->Value);
      SetType(&Symbol->Type, SYMBOL_PSEUDO_REG_TYPE);
      return Symbol;
    case SEMANTIC_RULE:
      Symbol->Value = SemanticRuleToInt(Token->Value);
      SetType(&Symbol->Type, SYMBOL_SEMANTIC_RULE_TYPE);
      return Symbol;
    case TEMP:
      Symbol->Value = DecimalToInt(Token->Value);
      SetType(&Symbol->Type, SYMBOL_TEMP_TYPE);
      return Symbol;
    case STRING:
      RemoveSymbol(&Symbol);
      return NewStringSymbol(Token);
    case WSTRING:
      RemoveSymbol(&Symbol);
      return NewWstringSymbol(Token);
    case FUNCTION_PARAMETER_ID:
      Symbol->Value = GetFunctionParameterIdentifier(Token);
      SetType(&Symbol->Type, SYMBOL_FUNCTION_PARAMETER_ID_TYPE);
      return Symbol;
    default:
      *Error = SCRIPT_ENGINE_ERROR_UNRESOLVED_VARIABLE;
      Symbol->Type = INVALID;
      Symbol->Value = INVALID;
      return Symbol;
  }
}

PSYMBOL_BUFFER NewSymbolBuffer(void) {
  PSYMBOL_BUFFER SymbolBuffer;
  SymbolBuffer = (PSYMBOL_BUFFER)malloc(sizeof(*SymbolBuffer));
  if (SymbolBuffer == NULL) {
    return NULL;
  }
  SymbolBuffer->Pointer = 0;
  SymbolBuffer->Size = SYMBOL_BUFFER_INIT_SIZE;
  SymbolBuffer->Head = (PSYMBOL)malloc(SymbolBuffer->Size * sizeof(SYMBOL));
  SymbolBuffer->Message = NULL;
  return SymbolBuffer;
}

void RemoveSymbolBuffer(PVOID SymbolBuffer) {
  PSYMBOL_BUFFER SymBuf = (PSYMBOL_BUFFER)SymbolBuffer;
  free(SymBuf->Message);
  free(SymBuf->Head);
  free(SymBuf);
}

PSYMBOL_BUFFER PushSymbol(PSYMBOL_BUFFER SymbolBuffer, const PSYMBOL Symbol) {
  uintptr_t Head = (uintptr_t)SymbolBuffer->Head;
  uintptr_t Pointer = (uintptr_t)SymbolBuffer->Pointer;
  PSYMBOL WriteAddr = (PSYMBOL)(Head + Pointer * sizeof(SYMBOL));
  if (Symbol->Type == SYMBOL_STRING_TYPE ||
      Symbol->Type == SYMBOL_WSTRING_TYPE) {
    SymbolBuffer->Pointer += GetSymbolHeapSize(Symbol);
    if (SymbolBuffer->Pointer >= SymbolBuffer->Size - 1) {
      unsigned int NewSize = SymbolBuffer->Size;
      do {
        NewSize *= 2;
      } while (NewSize <= SymbolBuffer->Pointer);
      PSYMBOL NewHead = (PSYMBOL)malloc(NewSize * sizeof(SYMBOL));
      if (NewHead == NULL) {
        printf("err, could not allocate buffer");
        return NULL;
      }
      memcpy(NewHead, SymbolBuffer->Head, SymbolBuffer->Size * sizeof(SYMBOL));
      free(SymbolBuffer->Head);
      SymbolBuffer->Size = NewSize;
      SymbolBuffer->Head = NewHead;
    }
    WriteAddr = (PSYMBOL)((uintptr_t)SymbolBuffer->Head +
                          (uintptr_t)Pointer * (uintptr_t)sizeof(SYMBOL));
    WriteAddr->Type = Symbol->Type;
    WriteAddr->Len = Symbol->Len;
    memcpy((char *)&WriteAddr->Value, (char *)&Symbol->Value, Symbol->Len);
  } else {
    *WriteAddr = *Symbol;
    SymbolBuffer->Pointer++;
    if (Pointer >= SymbolBuffer->Size - 1) {
      PSYMBOL NewHead =
          (PSYMBOL)malloc(2 * SymbolBuffer->Size * sizeof(SYMBOL));
      if (NewHead == NULL) {
        printf("err, could not allocate buffer");
        return NULL;
      }
      memcpy(NewHead, SymbolBuffer->Head, SymbolBuffer->Size * sizeof(SYMBOL));
      free(SymbolBuffer->Head);
      SymbolBuffer->Size *= 2;
      SymbolBuffer->Head = NewHead;
    }
  }
  return SymbolBuffer;
}

void PrintSymbolBuffer(const PVOID SymbolBuffer) {
  PSYMBOL_BUFFER SymBuff = (PSYMBOL_BUFFER)SymbolBuffer;
  PSYMBOL Symbol;
  printf("CodeBuffer:\n");
  for (unsigned int i = 0; i < SymBuff->Pointer;) {
    Symbol = SymBuff->Head + i;
    printf("Address = %d, ", i);
    PrintSymbol((PVOID)Symbol);
    if (Symbol->Type == SYMBOL_STRING_TYPE ||
        Symbol->Type == SYMBOL_WSTRING_TYPE) {
      int temp = GetSymbolHeapSize(Symbol);
      i += temp;
    } else {
      i++;
    }
  }
}

unsigned long long int RegisterToInt(char *str) {
  for (int i = 0; i < REGISTER_MAP_LIST_LENGTH; i++) {
    if (!strcmp(str, RegisterMapList[i].Name)) {
      return RegisterMapList[i].Type;
    }
  }
  if (g_HwdbgInstanceInfoIsValid) {
    const char *Ptr;
    UINT32 Num = 0;
    if (strncmp(str, "hw_pin", 6) == 0) {
      Ptr = str + 6;
      if (*Ptr == '\0') {
        return INVALID;
      }
      while (*Ptr) {
        if (!isdigit((unsigned char)*Ptr)) {
          return INVALID;
        }
        Ptr++;
      }
      Num = atoi(str + 6);
      if (Num >= g_HwdbgInstanceInfo.numberOfPins) {
        return INVALID;
      } else {
        return Num;
      }
    }
    if (strncmp(str, "hw_port", 7) == 0) {
      Ptr = str + 7;
      if (*Ptr == '\0') {
        return INVALID;
      }
      while (*Ptr) {
        if (!isdigit((unsigned char)*Ptr)) {
          return INVALID;
        }
        Ptr++;
      }
      Num = atoi(str + 7);
      if (Num >= g_HwdbgInstanceInfo.numberOfPorts) {
        return INVALID;
      } else {
        return Num + g_HwdbgInstanceInfo.numberOfPins;
      }
    }
  }
  return INVALID;
}

unsigned long long int PseudoRegToInt(char *str) {
  for (int i = 0; i < PSEUDO_REGISTER_MAP_LIST_LENGTH; i++) {
    if (!strcmp(str, PseudoRegisterMapList[i].Name)) {
      return PseudoRegisterMapList[i].Type;
    }
  }
  return INVALID;
}

unsigned long long int SemanticRuleToInt(char *str) {
  for (int i = 0; i < SEMANTIC_RULES_MAP_LIST_LENGTH; i++) {
    if (!strcmp(str, SemanticRulesMapList[i].Name)) {
      return SemanticRulesMapList[i].Type;
    }
  }
  return INVALID;
}

char *HandleError(PSCRIPT_ENGINE_ERROR_TYPE Error, char *str) {
  unsigned int LineEnd;
  for (int i = InputIdx;; i++) {
    if (str[i] == '\n' || str[i] == '\0') {
      LineEnd = i;
      break;
    }
  }
  int MessageSize = 16 + 100 + (CurrentTokenIdx - CurrentLineIdx) +
                    (LineEnd - CurrentLineIdx);
  char *Message = (char *)malloc(MessageSize);
  if (Message == NULL) {
    printf("err, could not allocate buffer");
    return NULL;
  }
  strcpy(Message, "Line ");
  char Line[16] = {0};
  sprintf(Line, "%d:\n", CurrentLine);
  strcat(Message, Line);
  strncat(Message, (str + CurrentLineIdx), LineEnd - CurrentLineIdx);
  strcat(Message, "\n");
  char Space = ' ';
  int n = (CurrentTokenIdx - CurrentLineIdx);
  for (int i = 0; i < n; i++) {
    strncat(Message, &Space, 1);
  }
  strcat(Message, "^\n");
  switch (*Error) {
    case SCRIPT_ENGINE_ERROR_SYNTAX:
      strcat(Message, "Syntax Error: ");
      strcat(Message, "Invalid Syntax");
      return Message;
    case SCRIPT_ENGINE_ERROR_UNKNOWN_TOKEN:
      strcat(Message, "Syntax Error: ");
      strcat(Message, "Unknown Token");
      return Message;
    case SCRIPT_ENGINE_ERROR_UNRESOLVED_VARIABLE:
      strcat(Message, "Syntax Error: ");
      strcat(Message, "Unresolved Variable");
      return Message;
    case SCRIPT_ENGINE_ERROR_UNHANDLED_SEMANTIC_RULE:
      strcat(Message, "Syntax Error: ");
      strcat(Message, "Unhandled Semantic Rule");
      return Message;
    case SCRIPT_ENGINE_ERROR_TEMP_LIST_FULL:
      strcat(Message, "Internal Error: ");
      strcat(Message,
             "Please split the expression to many smaller expressions.");
      return Message;
    case SCRIPT_ENGINE_ERROR_UNDEFINED_FUNCTION:
      strcat(Message, "Undefined Function");
      return Message;
    case SCRIPT_ENGINE_ERROR_UNDEFINED_VARIABLE_TYPE:
      strcat(Message, "Undefined Variable Type");
      return Message;
    case SCRIPT_ENGINE_ERROR_VOID_FUNCTION_RETURNING_VALUE:
      strcat(Message, "Returning a value in void function");
      return Message;
    case SCRIPT_ENGINE_ERROR_NON_VOID_FUNCTION_NOT_RETURNING_VALUE:
      strcat(Message, "Not returning a value in noo-void function");
      return Message;
    default:
      strcat(Message, "Unknown Error: ");
      return Message;
  }
}

int GetGlobalIdentifierVal(PTOKEN Token) {
  PTOKEN CurrentToken;
  for (uintptr_t i = 0; i < GlobalIdTable->Pointer; i++) {
    CurrentToken = *(GlobalIdTable->Head + i);
    if (!strcmp(Token->Value, CurrentToken->Value)) {
      return (int)i;
    }
  }
  return -1;
}

int GetLocalIdentifierVal(PTOKEN Token) {
  PTOKEN CurrentToken;
  for (uintptr_t i = 0;
       i < ((PTOKEN_LIST)CurrentUserDefinedFunction->IdTable)->Pointer; i++) {
    CurrentToken =
        *(((PTOKEN_LIST)CurrentUserDefinedFunction->IdTable)->Head + i);
    if (!strcmp(Token->Value, CurrentToken->Value)) {
      return (int)i;
    }
  }
  return -1;
}

int NewGlobalIdentifier(PTOKEN Token) {
  PTOKEN CopiedToken = CopyToken(Token);
  GlobalIdTable = Push(GlobalIdTable, CopiedToken);
  return GlobalIdTable->Pointer - 1;
}

int NewLocalIdentifier(PTOKEN Token) {
  PTOKEN CopiedToken = CopyToken(Token);
  Push(((PTOKEN_LIST)CurrentUserDefinedFunction->IdTable), CopiedToken);
  CurrentUserDefinedFunction->LocalVariableNumber++;
  return ((PTOKEN_LIST)CurrentUserDefinedFunction->IdTable)->Pointer - 1;
}

int NewFunctionParameterIdentifier(PTOKEN Token) {
  PTOKEN CopiedToken = CopyToken(Token);
  Push(((PTOKEN_LIST)CurrentUserDefinedFunction->FunctionParameterIdTable),
       CopiedToken);
  return ((PTOKEN_LIST)CurrentUserDefinedFunction->FunctionParameterIdTable)
             ->Pointer -
         1;
}

int GetFunctionParameterIdentifier(PTOKEN Token) {
  PTOKEN CurrentToken;
  for (uintptr_t i = 0;
       i < ((PTOKEN_LIST)CurrentUserDefinedFunction->FunctionParameterIdTable)
               ->Pointer;
       i++) {
    CurrentToken =
        *(((PTOKEN_LIST)CurrentUserDefinedFunction->FunctionParameterIdTable)
              ->Head +
          i);
    if (!strcmp(Token->Value, CurrentToken->Value)) {
      return (int)i;
    }
  }
  return -1;
}

PUSER_DEFINED_FUNCTION_NODE GetUserDefinedFunctionNode(PTOKEN Token) {
  PUSER_DEFINED_FUNCTION_NODE Node = UserDefinedFunctionHead;
  while (Node) {
    if (!strcmp((const char *)Token->Value, Node->Name)) {
      return Node;
      break;
    }
    Node = Node->NextNode;
  }
  return 0;
}

int LalrGetRhsSize(int RuleId) {
  int Counter = 0;
  int N = LalrRhsSize[RuleId];
  for (int i = 0; i < N; i++) {
    if (LalrRhs[RuleId][i].Type != EPSILON &&
        LalrRhs[RuleId][i].Type != SEMANTIC_RULE) {
      Counter++;
    }
  }
  return Counter;
}

BOOL LalrIsOperandType(PTOKEN Token) {
  if (Token->Type == GLOBAL_ID) {
    return TRUE;
  } else if (Token->Type == GLOBAL_UNRESOLVED_ID) {
    return TRUE;
  }
  if (Token->Type == LOCAL_ID) {
    return TRUE;
  } else if (Token->Type == LOCAL_UNRESOLVED_ID) {
    return TRUE;
  } else if (Token->Type == FUNCTION_PARAMETER_ID) {
    return TRUE;
  } else if (Token->Type == FUNCTION_ID) {
    return TRUE;
  } else if (Token->Type == DECIMAL) {
    return TRUE;
  } else if (Token->Type == HEX) {
    return TRUE;
  } else if (Token->Type == OCTAL) {
    return TRUE;
  } else if (Token->Type == BINARY) {
    return TRUE;
  } else if (Token->Type == REGISTER) {
    return TRUE;
  } else if (Token->Type == PSEUDO_REGISTER) {
    return TRUE;
  } else if (Token->Type == TEMP) {
    return TRUE;
  } else if (Token->Type == STRING) {
    return TRUE;
  } else if (Token->Type == WSTRING) {
    return TRUE;
  }
  return FALSE;
}

BOOLEAN ScriptEngineSetHwdbgInstanceInfo(
    HWDBG_INSTANCE_INFORMATION *InstancInfo) {
  memcpy(&g_HwdbgInstanceInfo, InstancInfo, sizeof(HWDBG_INSTANCE_INFORMATION));
  g_HwdbgInstanceInfoIsValid = TRUE;
  return TRUE;
}

BOOLEAN FuncGetNumberOfOperands(UINT64 FuncType, UINT32 *NumberOfGetOperands,
                                UINT32 *NumberOfSetOperands) {
  BOOLEAN Result = FALSE;
  switch (FuncType) {
    case FUNC_INC:
      *NumberOfGetOperands = 1;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_DEC:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_OR:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_XOR:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_AND:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_ASL:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_ADD:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_SUB:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_MUL:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_DIV:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_MOD:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_GT:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_LT:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_EGT:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_ELT:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_EQUAL:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_NEQ:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    case FUNC_JMP:
      *NumberOfGetOperands = 1;
      *NumberOfSetOperands = 0;
      Result = TRUE;
      break;
    case FUNC_JZ:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 0;
      Result = TRUE;
      break;
    case FUNC_JNZ:
      *NumberOfGetOperands = 2;
      *NumberOfSetOperands = 0;
      Result = TRUE;
      break;
    case FUNC_MOV:
      *NumberOfGetOperands = 1;
      *NumberOfSetOperands = 1;
      Result = TRUE;
      break;
    default:
      Result = FALSE;
      break;
  }
  return Result;
}
