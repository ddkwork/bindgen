
#include "pch.h"

PTOKEN NewUnknownToken() {
  PTOKEN Token;
  Token = (PTOKEN)malloc(sizeof(TOKEN));
  if (Token == NULL) {
    return NULL;
  }
  Token->Value = (char *)calloc(TOKEN_VALUE_MAX_LEN + 1, sizeof(char));
  if (Token->Value == NULL) {
    free(Token);
    return NULL;
  }
  strcpy(Token->Value, "");
  Token->Type = UNKNOWN;
  Token->Len = 0;
  Token->MaxLen = TOKEN_VALUE_MAX_LEN;
  Token->VariableType = 0;
  return Token;
}

PTOKEN NewToken(TOKEN_TYPE Type, char *Value) {
  PTOKEN Token = (PTOKEN)malloc(sizeof(TOKEN));
  if (Token == NULL) {
    return NULL;
  }
  unsigned int Len = (unsigned int)strlen(Value);
  Token->Type = Type;
  Token->Len = Len;
  Token->MaxLen = Len;
  Token->Value = (char *)calloc(Token->MaxLen + 1, sizeof(char));
  Token->VariableType = 0;
  if (Token->Value == NULL) {
    free(Token);
    return NULL;
  }
  strcpy(Token->Value, Value);
  return Token;
}

void RemoveToken(PTOKEN *Token) {
  free((*Token)->Value);
  free(*Token);
  *Token = NULL;
  return;
}

void PrintToken(PTOKEN Token) {
  if (Token->Type == WHITE_SPACE) {
    printf("< :");
  } else {
    printf("<'%s' : ", Token->Value);
  }
  switch (Token->Type) {
    case GLOBAL_ID:
      printf(" GLOBAL_ID>\n");
      break;
    case GLOBAL_UNRESOLVED_ID:
      printf(" GLOBAL_UNRESOLVED_ID>\n");
      break;
    case LOCAL_ID:
      printf(" LOCAL_ID>\n");
      break;
    case LOCAL_UNRESOLVED_ID:
      printf(" LOCAL_UNRESOLVED_ID>\n");
      break;
    case STATE_ID:
      printf(" STATE_ID>\n");
      break;
    case DECIMAL:
      printf(" DECIMAL>\n");
      break;
    case HEX:
      printf(" HEX>\n");
      break;
    case OCTAL:
      printf(" OCTAL>\n");
      break;
    case BINARY:
      printf(" BINARY>\n");
      break;
    case SPECIAL_TOKEN:
      printf(" SPECIAL_TOKEN>\n");
      break;
    case KEYWORD:
      printf(" KEYWORD>\n");
      break;
    case WHITE_SPACE:
      printf(" WHITE_SPACE>\n");
      break;
    case COMMENT:
      printf(" COMMENT>\n");
      break;
    case REGISTER:
      printf(" REGISTER>\n");
      break;
    case PSEUDO_REGISTER:
      printf(" PSEUDO_REGISTER>\n");
      break;
    case SEMANTIC_RULE:
      printf(" SEMANTIC_RULE>\n");
      break;
    case NON_TERMINAL:
      printf(" NON_TERMINAL>\n");
      break;
    case END_OF_STACK:
      printf(" END_OF_STACK>\n");
      break;
    case STRING:
      printf(" STRING>\n");
      break;
    case WSTRING:
      printf(" WSTRING>\n");
      break;
    case TEMP:
      printf(" TEMP>\n");
      break;
    case UNKNOWN:
      printf(" UNKNOWN>\n");
      break;
    case SCRIPT_VARIABLE_TYPE:
      printf(" SCRIPT_VARIABLE_TYPE>\n");
      break;
    case FUNCTION_ID:
      printf(" FUNCTION_ID>\n");
      break;
    case FUNCTION_PARAMETER_ID:
      printf(" FUNCTION_PARAMETER_ID>\n");
      break;
    default:
      printf(" ERROR>\n");
      break;
  }
}

void AppendByte(PTOKEN Token, char c) {
  if (Token->Len >= Token->MaxLen - 1) {
    Token->MaxLen *= 2;
    char *NewValue = (char *)calloc(Token->MaxLen + 1, sizeof(char));
    if (NewValue == NULL) {
      printf("err, could not allocate buffer");
      return;
    }
    memcpy(NewValue, Token->Value, Token->Len);
    free(Token->Value);
    Token->Value = NewValue;
  }
  Token->Value[Token->Len] = c;
  Token->Len++;
}

void AppendWchar(PTOKEN Token, wchar_t c) {
  if (Token->Len >= Token->MaxLen - 2) {
    Token->MaxLen *= 2;
    char *NewValue = (char *)calloc(Token->MaxLen + 2, sizeof(char));
    if (NewValue == NULL) {
      printf("err, could not allocate buffer");
      return;
    }
    memcpy(NewValue, Token->Value, Token->Len);
    free(Token->Value);
    Token->Value = NewValue;
  }
  *((wchar_t *)(Token->Value) + Token->Len / 2) = c;
  Token->Len += 2;
}

PTOKEN CopyToken(PTOKEN Token) {
  PTOKEN TokenCopy = (PTOKEN)malloc(sizeof(TOKEN));
  if (TokenCopy == NULL) {
    return NULL;
  }
  TokenCopy->Type = Token->Type;
  TokenCopy->MaxLen = Token->MaxLen;
  TokenCopy->Len = Token->Len;
  TokenCopy->Value = (char *)calloc(strlen(Token->Value) + 1, sizeof(char));
  TokenCopy->VariableType = Token->VariableType;
  if (TokenCopy->Value == NULL) {
    free(TokenCopy);
    return NULL;
  }
  strcpy(TokenCopy->Value, Token->Value);
  return TokenCopy;
}

PTOKEN_LIST NewTokenList(void) {
  PTOKEN_LIST TokenList = NULL;
  TokenList = (PTOKEN_LIST)malloc(sizeof(*TokenList));
  if (TokenList == NULL) {
    return NULL;
  }
  TokenList->Pointer = 0;
  TokenList->Size = TOKEN_LIST_INIT_SIZE;
  TokenList->Head = (PTOKEN *)malloc(TokenList->Size * sizeof(PTOKEN));
  return TokenList;
}

void RemoveTokenList(PTOKEN_LIST TokenList) {
  PTOKEN Token;
  for (uintptr_t i = 0; i < TokenList->Pointer; i++) {
    Token = *(TokenList->Head + i);
    RemoveToken(&Token);
  }
  free(TokenList->Head);
  free(TokenList);
  return;
}

void PrintTokenList(PTOKEN_LIST TokenList) {
  PTOKEN Token;
  for (uintptr_t i = 0; i < TokenList->Pointer; i++) {
    Token = *(TokenList->Head + i);
    PrintToken(Token);
  }
}

PTOKEN_LIST Push(PTOKEN_LIST TokenList, PTOKEN Token) {
  uintptr_t Head = (uintptr_t)TokenList->Head;
  uintptr_t Pointer = (uintptr_t)TokenList->Pointer;
  PTOKEN *WriteAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));
  *WriteAddr = Token;
  TokenList->Pointer++;
  if (Pointer == TokenList->Size - 1) {
    PTOKEN *NewHead = (PTOKEN *)malloc(2 * TokenList->Size * sizeof(PTOKEN));
    if (NewHead == NULL) {
      printf("err, could not allocate buffer");
      return NULL;
    }
    memcpy(NewHead, TokenList->Head, TokenList->Size * sizeof(PTOKEN));
    free(TokenList->Head);
    TokenList->Size = TokenList->Size * 2;
    TokenList->Head = NewHead;
  }
  return TokenList;
}

PTOKEN Pop(PTOKEN_LIST TokenList) {
  if (TokenList->Pointer > 0) TokenList->Pointer--;
  uintptr_t Head = (uintptr_t)TokenList->Head;
  uintptr_t Pointer = (uintptr_t)TokenList->Pointer;
  PTOKEN *ReadAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));
  return *ReadAddr;
}

PTOKEN Top(PTOKEN_LIST TokenList) {
  uintptr_t Head = (uintptr_t)TokenList->Head;
  uintptr_t Pointer = (uintptr_t)TokenList->Pointer - 1;
  PTOKEN *ReadAddr = (PTOKEN *)(Head + Pointer * sizeof(PTOKEN));
  return *ReadAddr;
}

char IsHex(char c) {
  if ((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') ||
      (c >= 'A' && c <= 'F'))
    return 1;
  else
    return 0;
}

char IsDecimal(char c) {
  if (c >= '0' && c <= '9')
    return 1;
  else
    return 0;
}

char IsLetter(char c) {
  if ((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'))
    return 1;
  else {
    return 0;
  }
}

char IsUnderscore(char c) {
  if (c >= '_')
    return 1;
  else {
    return 0;
  }
}

char IsBinary(char c) {
  if (c == '0' || c == '1')
    return 1;
  else {
    return 0;
  }
}

char IsOctal(char c) {
  if (c >= '0' && c <= '7')
    return 1;
  else
    return 0;
}

PTOKEN NewTemp(PSCRIPT_ENGINE_ERROR_TYPE Error) {
  static unsigned int TempID = 0;
  int i;
  for (i = 0; i < MAX_TEMP_COUNT; i++) {
    if (CurrentUserDefinedFunction->TempMap[i] == 0) {
      TempID = i;
      CurrentUserDefinedFunction->TempMap[i] = 1;
      break;
    }
  }
  if (i == MAX_TEMP_COUNT) {
    *Error = SCRIPT_ENGINE_ERROR_TEMP_LIST_FULL;
  }
  PTOKEN Temp = NewUnknownToken();
  char TempValue[8];
  sprintf(TempValue, "%d", TempID);
  strcpy(Temp->Value, TempValue);
  Temp->Type = TEMP;
  if (CurrentUserDefinedFunction->MaxTempNumber < (i + 1)) {
    CurrentUserDefinedFunction->MaxTempNumber = i + 1;
  }
  return Temp;
}

void FreeTemp(PTOKEN Temp) {
  int id = (int)DecimalToInt(Temp->Value);
  if (Temp->Type == TEMP) {
    CurrentUserDefinedFunction->TempMap[id] = 0;
  }
}

char IsType1Func(PTOKEN Operator) {
  unsigned int n = ONEOPFUNC1_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, OneOpFunc1[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType2Func(PTOKEN Operator) {
  unsigned int n = ONEOPFUNC2_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, OneOpFunc2[i])) {
      return 1;
    }
  }
  return 0;
}

char IsTwoOperandOperator(PTOKEN Operator) {
  unsigned int n = OPERATORS_TWO_OPERAND_LIST_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, OperatorsTwoOperandList[i])) {
      return 1;
    }
  }
  return 0;
}

char IsOneOperandOperator(PTOKEN Operator) {
  unsigned int n = OPERATORS_ONE_OPERAND_LIST_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, OperatorsOneOperandList[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType4Func(PTOKEN Operator) {
  unsigned int n = VARARGFUNC1_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, VarArgFunc1[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType5Func(PTOKEN Operator) {
  unsigned int n = ZEROOPFUNC1_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, ZeroOpFunc1[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType6Func(PTOKEN Operator) {
  unsigned int n = TWOOPFUNC1_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, TwoOpFunc1[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType7Func(PTOKEN Operator) {
  unsigned int n = TWOOPFUNC2_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, TwoOpFunc2[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType8Func(PTOKEN Operator) {
  unsigned int n = THREEOPFUNC1_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, ThreeOpFunc1[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType9Func(PTOKEN Operator) {
  unsigned int n = ONEOPFUNC3_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, OneOpFunc3[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType10Func(PTOKEN Operator) {
  unsigned int n = TWOOPFUNC3_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, TwoOpFunc3[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType11Func(PTOKEN Operator) {
  unsigned int n = THREEOPFUNC3_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, ThreeOpFunc3[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType12Func(PTOKEN Operator) {
  unsigned int n = ONEOPFUNC4_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, OneOpFunc4[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType13Func(PTOKEN Operator) {
  unsigned int n = TWOOPFUNC4_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, TwoOpFunc4[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType14Func(PTOKEN Operator) {
  unsigned int n = THREEOPFUNC2_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, ThreeOpFunc2[i])) {
      return 1;
    }
  }
  return 0;
}

char IsType15Func(PTOKEN Operator) {
  unsigned int n = THREEOPFUNC4_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, ThreeOpFunc4[i])) {
      return 1;
    }
  }
  return 0;
}

char IsAssignmentOperator(PTOKEN Operator) {
  unsigned int n = ASSIGNMENT_OPERATOR_LIST_LENGTH;
  for (unsigned int i = 0; i < n; i++) {
    if (!strcmp(Operator->Value, AssignmentOperatorList[i])) {
      return 1;
    }
  }
  return 0;
}

char IsNoneTerminal(PTOKEN Token) {
  if (Token->Value[0] >= 'A' && Token->Value[0] <= 'Z')
    return 1;
  else
    return 0;
}

char IsSemanticRule(PTOKEN Token) {
  if (Token->Value[0] == '@')
    return 1;
  else
    return 0;
}

int GetNonTerminalId(PTOKEN Token) {
  for (int i = 0; i < NONETERMINAL_COUNT; i++) {
    if (!strcmp(Token->Value, NoneTerminalMap[i])) return i;
  }
  return INVALID;
}

int GetTerminalId(PTOKEN Token) {
  for (int i = 0; i < TERMINAL_COUNT; i++) {
    if (Token->Type == HEX) {
      if (!strcmp("_hex", TerminalMap[i])) return i;
    } else if (Token->Type == GLOBAL_ID ||
               Token->Type == GLOBAL_UNRESOLVED_ID) {
      if (!strcmp("_global_id", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == LOCAL_ID || Token->Type == LOCAL_UNRESOLVED_ID) {
      if (!strcmp("_local_id", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == FUNCTION_ID) {
      if (!strcmp("_function_id", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == FUNCTION_PARAMETER_ID) {
      if (!strcmp("_function_parameter_id", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == REGISTER) {
      if (!strcmp("_register", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == PSEUDO_REGISTER) {
      if (!strcmp("_pseudo_register", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == SCRIPT_VARIABLE_TYPE) {
      if (!strcmp("_script_variable_type", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == DECIMAL) {
      if (!strcmp("_decimal", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == BINARY) {
      if (!strcmp("_binary", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == OCTAL) {
      if (!strcmp("_octal", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == STRING) {
      if (!strcmp("_string", TerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == WSTRING) {
      if (!strcmp("_wstring", TerminalMap[i])) {
        return i;
      }
    } else {
      if (!strcmp(Token->Value, TerminalMap[i])) return i;
    }
  }
  return INVALID;
}

int LalrGetNonTerminalId(PTOKEN Token) {
  for (int i = 0; i < LALR_NONTERMINAL_COUNT; i++) {
    if (!strcmp(Token->Value, LalrNoneTerminalMap[i])) return i;
  }
  return INVALID;
}

int LalrGetTerminalId(PTOKEN Token) {
  for (int i = 0; i < LALR_TERMINAL_COUNT; i++) {
    if (Token->Type == HEX) {
      if (!strcmp("_hex", LalrTerminalMap[i])) return i;
    } else if (Token->Type == GLOBAL_ID ||
               Token->Type == GLOBAL_UNRESOLVED_ID) {
      if (!strcmp("_global_id", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == LOCAL_ID || Token->Type == LOCAL_UNRESOLVED_ID) {
      if (!strcmp("_local_id", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == FUNCTION_ID) {
      if (!strcmp("_function_id", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == FUNCTION_PARAMETER_ID) {
      if (!strcmp("_function_parameter_id", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == REGISTER) {
      if (!strcmp("_register", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == PSEUDO_REGISTER) {
      if (!strcmp("_pseudo_register", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == DECIMAL) {
      if (!strcmp("_decimal", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == BINARY) {
      if (!strcmp("_binary", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == OCTAL) {
      if (!strcmp("_octal", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == STRING) {
      if (!strcmp("_string", LalrTerminalMap[i])) {
        return i;
      }
    } else if (Token->Type == WSTRING) {
      if (!strcmp("_wstring", LalrTerminalMap[i])) {
        return i;
      }
    } else {
      if (!strcmp(Token->Value, LalrTerminalMap[i])) return i;
    }
  }
  return INVALID;
}

char IsEqual(const PTOKEN Token1, const PTOKEN Token2) {
  if (Token1->Type == Token2->Type) {
    if (Token1->Type == SPECIAL_TOKEN) {
      if (!strcmp(Token1->Value, Token2->Value)) {
        return 1;
      }
    } else {
      return 1;
    }
  }
  if (Token1->Type == GLOBAL_ID && Token2->Type == GLOBAL_UNRESOLVED_ID) {
    return 1;
  }
  if (Token1->Type == GLOBAL_UNRESOLVED_ID && Token2->Type == GLOBAL_ID) {
    return 1;
  }
  if (Token1->Type == LOCAL_ID && Token2->Type == LOCAL_UNRESOLVED_ID) {
    return 1;
  }
  if (Token1->Type == LOCAL_UNRESOLVED_ID && Token2->Type == LOCAL_ID) {
    return 1;
  }
  return 0;
}

void SetType(unsigned long long *Val, unsigned char Type) {
  *Val = (unsigned long long int)Type;
}

unsigned long long DecimalToInt(char *str) {
  unsigned long long Acc = 0;
  size_t Len;
  Len = strlen(str);
  for (int i = 0; i < Len; i++) {
    Acc *= 10;
    Acc += (str[i] - '0');
  }
  return Acc;
}

unsigned long long DecimalToSignedInt(char *str) {
  long long Acc = 0;
  size_t Len;
  if (str[0] == '-') {
    Len = strlen(str);
    for (int i = 1; i < Len; i++) {
      Acc *= 10;
      Acc += (str[i] - '0');
    }
    return -Acc;
  } else {
    Len = strlen(str);
    for (int i = 0; i < Len; i++) {
      Acc *= 10;
      Acc += (str[i] - '0');
    }
    return Acc;
  }
}

unsigned long long HexToInt(char *str) {
  char temp;
  size_t len = strlen(str);
  unsigned long long Acc = 0;
  for (int i = 0; i < len; i++) {
    Acc <<= 4;
    if (str[i] >= '0' && str[i] <= '9') {
      temp = str[i] - '0';
    } else if (str[i] >= 'a' && str[i] <= 'f') {
      temp = str[i] - 'a' + 10;
    } else {
      temp = str[i] - 'A' + 10;
    }
    Acc += temp;
  }
  return Acc;
}

unsigned long long OctalToInt(char *str) {
  size_t Len;
  unsigned long long Acc = 0;
  Len = strlen(str);
  for (int i = 0; i < Len; i++) {
    Acc <<= 3;
    Acc += (str[i] - '0');
  }
  return Acc;
}

unsigned long long BinaryToInt(char *str) {
  size_t Len;
  unsigned long long Acc = 0;
  Len = strlen(str);
  for (int i = 0; i < Len; i++) {
    Acc <<= 1;
    Acc += (str[i] - '0');
  }
  return Acc;
}

void RotateLeftStringOnce(char *str) {
  int length = (int)strlen(str);
  char temp = str[0];
  for (int i = 0; i < (length - 1); i++) {
    str[i] = str[i + 1];
  }
  str[length - 1] = temp;
}
