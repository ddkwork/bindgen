
#include "pch.h"

int main(int argc, char *argv[]) {
  if (argc != 2) {
    printf(
        "you should not test functionalities directly, instead use 'test all' "
        "command from HyperDbg...\n");
    return 1;
  }
  if (!strcmp(argv[1], TEST_CASE_PARAMETER_FOR_MAIN_COMMAND_PARSER)) {
    if (TestCommandParser()) {
      printf("\n[*] The main command parser test cases passed successfully\n");
    } else {
      printf("\n[x] The main command parser test cases failed\n");
    }
  } else if (!strcmp(argv[1],
                     TEST_CASE_PARAMETER_FOR_SCRIPT_SEMANTIC_TEST_CASES)) {
    if (TestSemanticScripts()) {
      printf("\n[*] The script semantic test cases passed successfully\n");
    } else {
      printf("\n[x] The script semantic test cases failed\n");
    }
  } else if (!strcmp(argv[1], TEST_HWDBG_FUNCTIONALITIES)) {
    if (HwdbgTestCreateTestCases()) {
      printf("\n[*] The hwdbg test cases passed successfully\n");
    } else {
      printf("\n[x] The hwdbg test cases failed\n");
    }
  } else {
    printf("unknown test case\n");
  }
  printf("\npress any key to exit...");
  _getch();
  return 0;
}
