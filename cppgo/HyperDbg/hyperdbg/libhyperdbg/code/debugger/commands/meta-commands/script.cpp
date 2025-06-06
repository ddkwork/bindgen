
#include "pch.h"
using namespace std;
extern BOOLEAN g_ExecutingScript;

VOID CommandScriptHelp() {
  ShowMessages(".script : runs a HyperDbg script.\n\n");
  ShowMessages("syntax : .script [FilePath (string)] [Args (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .script C:\\scripts\\script.ds\n");
  ShowMessages("\t\te.g : .script C:\\scripts\\script.ds 95 85 @rsp\n");
  ShowMessages("\t\te.g : .script \"C:\\scripts\\hello world.ds\"\n");
  ShowMessages("\t\te.g : .script \"C:\\scripts\\hello world.ds\" @rax\n");
  ShowMessages(
      "\t\te.g : .script \"C:\\scripts\\hello world.ds\" @rax @rcx+55 $pid\n");
  ShowMessages(
      "\t\te.g : .script \"C:\\scripts\\hello world.ds\" 12 55 @rip\n");
}

VOID CommandScriptRunCommand(std::string Input, vector<string> PathAndArgs) {
  int CommandExecutionResult = 0;
  char *LineContent = NULL;
  int i = 0;
  for (auto item : PathAndArgs) {
    string ToReplace = "$arg" + std::to_string(i);
    i++;
    ReplaceAll(Input, ToReplace, item);
  }
  LineContent = (char *)Input.c_str();
  if (IsEmptyString(LineContent)) {
    return;
  }
  HyperDbgShowSignature();
  ShowMessages("%s\n", LineContent);
  CommandExecutionResult = HyperDbgInterpreter(LineContent);
  ShowMessages("\n");
  if (CommandExecutionResult == 1) {
    exit(0);
  }
}

VOID HyperDbgScriptReadFileAndExecuteCommand(
    std::vector<std::string> &PathAndArgs) {
  std::string Line;
  BOOLEAN IsOpened = FALSE;
  bool Reset = false;
  string CommandToExecute = "";
  string PathOfScriptFile = "";
  PathOfScriptFile = PathAndArgs.at(0);
  ReplaceAll(PathOfScriptFile, "\"", "");
  ifstream File(PathOfScriptFile);
  if (File.is_open()) {
    IsOpened = TRUE;
    g_ExecutingScript = TRUE;
    Reset = true;
    while (std::getline(File, Line)) {
      if (CheckMultilineCommand((char *)Line.c_str(), Reset)) {
        if (Reset) {
          CommandToExecute.clear();
        }
        Reset = false;
        CommandToExecute += Line + "\n";
        continue;
      } else {
        Reset = true;
        CommandToExecute += Line;
      }
      CommandScriptRunCommand(CommandToExecute, PathAndArgs);
      CommandToExecute.clear();
    }
    if (!CommandToExecute.empty()) {
      CommandScriptRunCommand(CommandToExecute, PathAndArgs);
      CommandToExecute.clear();
    }
    g_ExecutingScript = FALSE;
    File.close();
  }
  if (!IsOpened) {
    ShowMessages("err, invalid file specified for the script\n");
  }
}

/**
 * @brief Parsing the command line options for scripts
 * @param argc
 * @param argv
 *
 * @return INT
 */
INT ScriptReadFileAndExecuteCommandline(INT argc, CHAR *argv[]) {
  vector<string> Args;
  for (size_t i = 2; i < argc; i++) {
    std::string TempStr(argv[i]);
    Args.push_back(TempStr);
  }
  if (!Args.empty()) {
    HyperDbgScriptReadFileAndExecuteCommand(Args);
    printf("\n");
  } else {
    printf("err, invalid command line options passed to the HyperDbg!\n");
    return FALSE;
  }
  return TRUE;
}

/**
 * @brief .script command handler
 *
 * @param CommandTokens
 * @param Command
 * @return VOID
 */
VOID CommandScript(vector<CommandToken> CommandTokens, string Command) {
  vector<string> PathAndArgs;
  BOOLEAN IsFirstCommand = FALSE;
  if (CommandTokens.size() == 1) {
    ShowMessages("please specify a file\n\n");
    CommandScriptHelp();
    return;
  }
  for (auto Section : CommandTokens) {
    if (!IsFirstCommand) {
      IsFirstCommand = TRUE;
      continue;
    }
    PathAndArgs.push_back(GetCaseSensitiveStringFromCommandToken(Section));
  }
  HyperDbgScriptReadFileAndExecuteCommand(PathAndArgs);
}
