
#include "pch.h"

VOID CommandFormatsHelp() {
  ShowMessages(
      ".formats : shows a value or register in different formats.\n\n");
  ShowMessages("syntax : \t.formats [Expression (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .formats nt!ExAllocatePoolWithTag\n");
  ShowMessages("\t\te.g : .formats nt!Kd_DEFAULT_Mask\n");
  ShowMessages("\t\te.g : .formats nt!Kd_DEFAULT_Mask+5\n");
  ShowMessages("\t\te.g : .formats 55\n");
  ShowMessages("\t\te.g : .formats @rax\n");
  ShowMessages("\t\te.g : .formats @rbx+@rcx\n");
  ShowMessages("\t\te.g : .formats $pid\n");
}

VOID CommandFormatsShowResults(UINT64 U64Value) {
  time_t t;
  struct tm *tmp;
  char MY_TIME[50];
  unsigned int Character;
  time(&t);
  tmp = localtime(&t);
  strftime(MY_TIME, sizeof(MY_TIME), "%x - %I:%M%p", tmp);
  ShowMessages("evaluate expression:\n");
  ShowMessages("Hex :        %s\n", SeparateTo64BitValue(U64Value).c_str());
  ShowMessages("Decimal :    %d\n", U64Value);
  ShowMessages("Octal :      %o\n", U64Value);
  ShowMessages("Binary :     ");
  PrintBits(sizeof(UINT64), &U64Value);
  ShowMessages("\nChar :       ");
  unsigned char *TempCharacter = (unsigned char *)&U64Value;
  for (size_t j = 0; j < sizeof(UINT64); j++) {
    Character = (unsigned int)TempCharacter[j];
    if (isprint(Character)) {
      ShowMessages("%c", Character);
    } else {
      ShowMessages(".");
    }
  }
  ShowMessages("\nTime :       %s\n", MY_TIME);
  ShowMessages("Float :      %4.2f %+.0e %E\n", U64Value, U64Value, U64Value);
  ShowMessages("Double :     %.*e\n", DECIMAL_DIG, U64Value);
}

VOID CommandFormats(vector<CommandToken> CommandTokens, string Command) {
  UINT64 ConstantValue = 0;
  BOOLEAN HasError = TRUE;
  if (CommandTokens.size() == 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandFormatsHelp();
    return;
  }
  Trim(Command);
  Command.erase(
      0, GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).size());
  Trim(Command);
  ConstantValue = ScriptEngineEvalSingleExpression(Command, &HasError);
  if (HasError) {
    ShowMessages("err, couldn't resolve error at '%s'\n", Command.c_str());
  } else {
    CommandFormatsShowResults(ConstantValue);
  }
}
