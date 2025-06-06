
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandEvalHelp() {
  ShowMessages("? : evaluates and execute expressions in debuggee.\n\n");
  ShowMessages("syntax : \t? [Expression (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : ? print(dq(poi(@rcx)));\n");
  ShowMessages("\t\te.g : ? json(dq(poi(@rcx)));\n");
}

BOOLEAN CommandEvalCheckTestcase() {
  string Line;
  BOOLEAN IsOpened = FALSE;
  UINT64 ExpectedValue = 0;
  BOOLEAN ExpectError = FALSE;
  string Expr = "";
  std::vector<std::string> TestFiles;
  try {
    TestFiles = ListDirectory(SCRIPT_ENGINE_TEST_CASES_DIRECTORY, "*.txt");
  } catch (const std::exception &) {
    ShowMessages(
        "err, test-cases not found, make sure to run test-environment.py "
        "before running test cases\n");
    return FALSE;
  }
  for (auto item : TestFiles) {
    ifstream File(item.c_str());
    if (File.is_open()) {
      IsOpened = TRUE;
      ShowMessages("Running from : %s\n\n", item.c_str());
      while (std::getline(File, Line)) {
        ShowMessages("Test-case number : %s\n", Line.c_str());
        if (!std::getline(File, Line)) {
          goto ErrorMessage;
        }
        Expr = Line;
        ShowMessages("Statement : %s\n", Line.c_str());
        if (!std::getline(File, Line)) {
          goto ErrorMessage;
        }
        if (!Line.compare("$error$")) {
          ShowMessages("Expected result : %s\n", Line.c_str());
          ExpectError = TRUE;
          ExpectedValue = NULL;
        } else if (!ConvertStringToUInt64(Line, &ExpectedValue)) {
          ShowMessages("err, the expected results are in incorrect format\n");
          goto ErrorMessage;
        } else {
          ExpectError = FALSE;
          ShowMessages("Expected result : %llx\n", ExpectedValue);
        }
        Expr.append(" ");
        if (ScriptAutomaticStatementsTestWrapper(Expr, ExpectedValue,
                                                 ExpectError)) {
          ShowMessages("Test result : Passed\n");
        } else {
          ShowMessages("Test result : Failed\n");
        }
        if (!std::getline(File, Line)) {
          goto ErrorMessage;
        }
        if (Line.compare("$end$")) {
          goto ErrorMessage;
        }
        ShowMessages(
            "\n------------------------------------------------------------"
            "\n\n");
      }
      File.close();
    }
  }
  if (!IsOpened) {
    ShowMessages("err, could not find files for script engine test-cases\n");
    return FALSE;
  }
  return TRUE;
ErrorMessage:
  ShowMessages(
      "\nerr, testing failed! incorrect file format for the testing "
      "script engine\nmake sure files have a correct ending format\n");
  return FALSE;
}

VOID CommandEval(vector<CommandToken> CommandTokens, string Command) {
  PVOID CodeBuffer;
  UINT64 BufferAddress;
  UINT32 BufferLength;
  UINT32 Pointer;
  if (CommandTokens.size() == 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandEvalHelp();
    return;
  }
  Trim(Command);
  Command.erase(
      0, GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).size());
  Trim(Command);
  if (!Command.compare("test")) {
    if (!CommandEvalCheckTestcase()) {
      ShowMessages(
          "err, testing script engine test-cases was not successful\n");
    } else {
      ShowMessages("testing script engine test-cases was successful\n");
    }
    return;
  }
  if (g_IsSerialConnectedToRemoteDebuggee) {
    CodeBuffer = ScriptEngineParseWrapper((char *)Command.c_str(), TRUE);
    if (CodeBuffer == NULL) {
      return;
    }
    BufferAddress = ScriptEngineWrapperGetHead(CodeBuffer);
    BufferLength = ScriptEngineWrapperGetSize(CodeBuffer);
    Pointer = ScriptEngineWrapperGetPointer(CodeBuffer);
    KdSendScriptPacketToDebuggee(BufferAddress, BufferLength, Pointer, FALSE);
    ScriptEngineWrapperRemoveSymbolBuffer(CodeBuffer);
  } else {
    ShowMessages(
        "this command should not be used while you're in VMI-Mode or not in "
        "debugger-mode, "
        "the results that you see is a simulated result for TESTING "
        "script-engine "
        "and is not based on the status of your system. You can use this "
        "command, "
        "ONLY in debugger-mode\n\n");
    ShowMessages("test expression : %s \n", Command.c_str());
    ScriptEngineWrapperTestParser(Command);
  }
}
