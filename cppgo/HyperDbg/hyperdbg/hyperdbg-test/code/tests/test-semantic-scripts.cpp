
#include "pch.h"
namespace fs = std::filesystem;

VOID ReadDirectoryAndTestSemanticTestcases(const char *ScriptSemanticPath) {
  try {
    for (const auto &entry : fs::directory_iterator(ScriptSemanticPath)) {
      if (entry.is_regular_file()) {
        std::string filePath = entry.path().string();
        std::ifstream file(filePath);
        if (file.is_open()) {
          std::string content((std::istreambuf_iterator<char>(file)),
                              std::istreambuf_iterator<char>());
          std::cout << "Executing file " << entry.path().filename().string()
                    << std::endl;
          hyperdbg_u_run_command((CHAR *)content.c_str());
          std::cout << "--------------------------------------------"
                    << std::endl;
          file.close();
        } else {
          std::cerr << "Could not open file: " << filePath << std::endl;
        }
      }
    }
  } catch (const fs::filesystem_error &e) {
    std::cerr << "Filesystem error: " << e.what() << std::endl;
  }
}

BOOLEAN TestSemanticScripts() {
  int testNum = 0;
  CHAR dirPath[MAX_PATH] = {0};
  if (!hyperdbg_u_setup_path_for_filename(SCRIPT_SEMANTIC_TEST_CASE_DIRECTORY,
                                          dirPath, MAX_PATH, FALSE)) {
    cout << "[-] Could not find the test case files" << endl;
    return FALSE;
  }
  if (!hyperdbg_u_connect_remote_debugger_using_named_pipe(
          TEST_DEFAULT_NAMED_PIPE, TRUE)) {
    cout << "[-] Could not connect to the debugger" << endl;
    return FALSE;
  }
  ReadDirectoryAndTestSemanticTestcases(dirPath);
  hyperdbg_u_debug_close_remote_debugger();
  return TRUE;
}
