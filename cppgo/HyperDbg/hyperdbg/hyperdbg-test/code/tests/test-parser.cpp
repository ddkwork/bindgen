
#include "pch.h"
typedef char **CHAR_PTR_PTR;

CHAR_PTR_PTR createTestCaseArray(const std::vector<std::string> &testCases) {
  CHAR_PTR_PTR testCaseArray =
      (CHAR_PTR_PTR)malloc(testCases.size() * sizeof(UINT64));
  for (size_t i = 0; i < testCases.size(); ++i) {
    testCaseArray[i] = (char *)malloc(testCases[i].length() + 1);
    if (testCaseArray[i] == NULL) {
      return NULL;
    }
    std::strcpy(testCaseArray[i], testCases[i].c_str());
  }
  return testCaseArray;
}

VOID freeTestCaseArray(CHAR_PTR_PTR testCaseArray, size_t size) {
  for (size_t i = 0; i < size; ++i) {
    free(testCaseArray[i]);
  }
  free(testCaseArray);
}

std::vector<std::pair<std::string, std::vector<std::string>>> parseTestCases(
    const std::string &filename) {
  std::ifstream file(filename);
  std::string line;
  std::string command;
  std::string currentToken;
  std::vector<std::string> tokens;
  std::vector<std::pair<std::string, std::vector<std::string>>> testCases;
  bool isCommand = false;
  bool addNewline = false;
  const std::string tokenDelimiter = "----------------------------------";
  const std::string commandDelimiter =
      "_____________________________________________________________";
  while (std::getline(file, line)) {
    if (line == commandDelimiter) {
      addNewline = false;
      if (!command.empty()) {
        if (!currentToken.empty()) {
          tokens.push_back(currentToken);
          currentToken.clear();
        }
        testCases.push_back({command, tokens});
        command.clear();
        tokens.clear();
      }
      isCommand = true;
    } else if (line == tokenDelimiter) {
      addNewline = false;
      isCommand = false;
      if (!currentToken.empty()) {
        tokens.push_back(currentToken);
        currentToken.clear();
      }
    } else {
      if (isCommand) {
        if (addNewline) command += "\n";
        command += line;
      } else {
        if (addNewline) currentToken += "\n";
        currentToken += line;
      }
      addNewline = true;
    }
  }
  if (!command.empty()) {
    if (!currentToken.empty()) {
      tokens.push_back(currentToken);
    }
    testCases.push_back({command, tokens});
  }
  return testCases;
}

int countBackslashNUpToPosition(const std::string &str, std::size_t limit) {
  int count = 0;
  std::string::size_type pos = 0;
  std::string target = "\\n";
  while ((pos = str.find(target, pos)) != std::string::npos && pos < limit) {
    ++count;
    pos += target.length();
  }
  return count;
}

VOID ShowParsedCommandAndTokens(
    const std::pair<std::string, std::vector<std::string>> &testCase,
    UINT32 failedTokenNum, UINT32 failedTokenPosition) {
  UINT32 tokenNum = 0;
  string showingCommand = testCase.first;
  std::string::size_type pos = 0;
  while ((pos = showingCommand.find("\n", pos)) != std::string::npos) {
    showingCommand.replace(pos, 1, "\\n");
    pos += 2;
  }
  std::cout << "Command: \"" << showingCommand << "\"" << std::endl;
  std::cout << "____________________________________\n";
  std::cout << "Expected Tokens: " << std::endl;
  for (const auto &token : testCase.second) {
    string showingToken = token;
    pos = 0;
    while ((pos = showingToken.find("\n", pos)) != std::string::npos) {
      showingToken.replace(pos, 1, "\\n");
      pos += 2;
    }
    if (tokenNum == failedTokenNum) {
      std::cout << "  x ";
    } else {
      std::cout << "  - ";
    }
    std::cout << "\"" << showingToken << "\"" << std::endl;
    if (tokenNum == failedTokenNum) {
      std::cout << "     ";
      int countOfSpaces =
          countBackslashNUpToPosition(showingToken, failedTokenPosition);
      countOfSpaces += failedTokenPosition;
      for (int i = 0; i < countOfSpaces; i++) {
        std::cout << " ";
      }
      std::cout << "^" << std::endl;
    }
    tokenNum++;
  }
}

BOOLEAN TestCommandParser() {
  BOOLEAN overallResult = TRUE;
  int testNum = 0;
  CHAR filePath[MAX_PATH] = {0};
  UINT32 failedTokenNum = 0;
  UINT32 failedTokenPosition = 0;
  if (!hyperdbg_u_setup_path_for_filename(COMMAND_PARSER_TEST_CASES_FILE,
                                          filePath, MAX_PATH, TRUE)) {
    cout << "[-] Could not find the test case files" << endl;
    return FALSE;
  }
  auto testCases = parseTestCases(filePath);
  cout << "Perfom testing test cases with parsed file:" << endl;
  for (const auto &testCase : testCases) {
    testNum++;
    CHAR_PTR_PTR testCaseArray = createTestCaseArray(testCase.second);
    if (hyperdbg_u_test_command_parser(
            (CHAR *)testCase.first.c_str(), (UINT32)testCase.second.size(),
            testCaseArray, &failedTokenNum, &failedTokenPosition)) {
      cout << "[+] Test number " << testNum << " Passed " << endl;
    } else {
      overallResult = FALSE;
      cout << "\n============================================================"
           << endl;
      cout << "\n*********************          " << endl;
      cout << "*** Error details ***          " << endl;
      cout << "*********************          " << endl;
      cout << "\nParsed tokens from HyperDbg main command parser:\n" << endl;
      hyperdbg_u_test_command_parser_show_tokens(
          (CHAR *)testCase.first.c_str());
      cout << "\n============================================================"
           << endl;
      cout << "\nThe parsed command and tokens (From file):" << endl;
      ShowParsedCommandAndTokens(testCase, failedTokenNum, failedTokenPosition);
      cout << "\n[-] Test number " << testNum << " Failed " << endl;
      cout << "============================================================\n"
           << endl;
      break;
    }
    freeTestCaseArray(testCaseArray, testCase.second.size());
  }
  return overallResult;
}
