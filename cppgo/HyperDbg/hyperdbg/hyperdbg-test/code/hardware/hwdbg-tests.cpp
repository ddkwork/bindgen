
#include "pch.h"
namespace fs = std::filesystem;

std::string commentContent(const std::string &content) {
  std::istringstream iss(content);
  std::string line;
  std::string commentedContent;
  while (std::getline(iss, line)) {
    commentedContent += line + "\n; ";
  }
  return commentedContent;
}

BOOLEAN ReadDirectoryAndCreateHwdbgTestCases(
    const char *HwdbgScriptTestCasesPath) {
  CHAR tempFilePath[MAX_PATH] = {0};
  try {
    for (const auto &entry : fs::directory_iterator(HwdbgScriptTestCasesPath)) {
      if (entry.is_regular_file()) {
        std::string filePath = entry.path().string();
        std::ifstream file(filePath);
        if (file.is_open()) {
          std::string content((std::istreambuf_iterator<char>(file)),
                              std::istreambuf_iterator<char>());
          std::cout << "Reading script test case file "
                    << entry.path().filename().string() << std::endl;
          std::string compiled_version_file_path =
              HWDBG_SCRIPT_TEST_CASE_COMPILED_SCRIPTS_DIRECTORY "\\" +
              entry.path().filename().string() + ".hex.txt";
          printf("File content: %s\n", content.c_str());
          if (!hwdbg_script_run_script(content.c_str(),
                                       HWDBG_TEST_READ_INSTANCE_INFO_PATH,
                                       compiled_version_file_path.c_str(),
                                       DEFAULT_INITIAL_BRAM_BUFFER_SIZE)) {
            std::cout << "[-] Could not run the script: " << filePath
                      << std::endl;
            return FALSE;
          }
          if (!hyperdbg_u_setup_path_for_filename(
                  compiled_version_file_path.c_str(), tempFilePath, MAX_PATH,
                  FALSE)) {
            std::cout << "[-] Could not find the compiled version of the hwdbg "
                         "test case file"
                      << endl;
            return FALSE;
          }
          std::ifstream compiledFile(tempFilePath);
          if (compiledFile.is_open()) {
            std::string compiledContent(
                (std::istreambuf_iterator<char>(compiledFile)),
                std::istreambuf_iterator<char>());
            compiledFile.close();
            std::string commentedContent = commentContent(content);
            std::string newContent =
                "; The raw script file is available "
                "at: " HWDBG_SCRIPT_TEST_CASE_COMPILED_SCRIPTS_DIRECTORY "\\" +
                entry.path().filename().string() + "\n;\n; !hw script " +
                commentedContent + "\n" + compiledContent;
            std::ofstream compiledFileOut(tempFilePath);
            if (compiledFileOut.is_open()) {
              compiledFileOut << newContent;
              compiledFileOut.close();
            } else {
              std::cerr << "Could not open file for writing: "
                        << compiled_version_file_path << std::endl;
            }
          } else {
            std::cerr << "Could not open compiled file: "
                      << compiled_version_file_path << std::endl;
          }
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
    return FALSE;
  }
  return TRUE;
}

BOOLEAN HwdbgTestCreateTestCases() {
  int testNum = 0;
  CHAR dirPath[MAX_PATH] = {0};
  if (!hyperdbg_u_setup_path_for_filename(HWDBG_SCRIPT_TEST_CASE_CODE_DIRECTORY,
                                          dirPath, MAX_PATH, FALSE)) {
    cout << "[-] Could not find the hwdbg test case files" << endl;
    return FALSE;
  }
  return ReadDirectoryAndCreateHwdbgTestCases(dirPath);
}
