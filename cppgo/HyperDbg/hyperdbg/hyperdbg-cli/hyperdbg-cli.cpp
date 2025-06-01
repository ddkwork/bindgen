
#include <Windows.h>
#include <conio.h>

#include <iostream>
#include <string>
#include <vector>

#include "SDK/HyperDbgSdk.h"
#include "SDK/imports/user/HyperDbgLibImports.h"
#include "platform/user/header/Environment.h"
using namespace std;

int main(int argc, char *argv[]) {
  BOOLEAN exit_from_debugger = FALSE;
  string previous_command;
  BOOLEAN reset = FALSE;
  SetConsoleOutputCP(CP_UTF8);
  printf("HyperDbg Debugger [version: %s, build: %s]\n", CompleteVersion,
         BuildVersion);
  printf("Please visit https://docs.hyperdbg.org for more information...\n");
  printf("HyperDbg is released under the GNU Public License v3 (GPLv3).\n\n");
  if (argc != 1) {
    if (!strcmp(argv[1], "--script")) {
      hyperdbg_u_script_read_file_and_execute_commandline(argc, argv);
    } else {
      printf("err, invalid command line options passed to the HyperDbg!\n");
      return 1;
    }
  }
  while (!exit_from_debugger) {
    hyperdbg_u_show_signature();
    string current_command = "";
    reset = TRUE;
  GetMultiLinecCommand:
    string temp_command = "";
    getline(cin, temp_command);
    if (cin.fail() || cin.eof()) {
      cin.clear();
      printf("\n\n");
      continue;
    }
    if (hyperdbg_u_check_multiline_command((CHAR *)temp_command.c_str(),
                                           reset) == TRUE) {
      reset = FALSE;
      current_command += temp_command + "\n";
      printf("> ");
      goto GetMultiLinecCommand;
    } else {
      reset = TRUE;
      current_command += temp_command;
    }
    if (!current_command.compare("") &&
        hyperdbg_u_continue_previous_command()) {
      current_command = previous_command;
    } else {
      previous_command = current_command;
    }
    INT CommandExecutionResult =
        hyperdbg_u_run_command((CHAR *)current_command.c_str());
    if (CommandExecutionResult == 1) {
      exit_from_debugger = true;
    }
    if (CommandExecutionResult != 2) {
      printf("\n");
    }
  }
  return 0;
}
