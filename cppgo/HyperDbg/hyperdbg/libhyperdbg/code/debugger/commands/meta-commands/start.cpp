
#include "pch.h"
extern std::wstring g_StartCommandPath;
extern std::wstring g_StartCommandPathAndArguments;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;

VOID CommandStartHelp() {
  ShowMessages(".start : runs a user-mode process.\n\n");
  ShowMessages(
      "syntax : \t.start [path Path (string)] [Parameters (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .start path c:\\reverse_eng\\my_file.exe\n");
  ShowMessages("\t\te.g : .start path \"c:\\reverse eng\\my_file.exe\"\n");
  ShowMessages(
      "\t\te.g : .start path \"c:\\reverse eng\\my_file.exe\" \"arg1\" 2 \"arg "
      "3\"\n");
}

VOID CommandStart(vector<CommandToken> CommandTokens, string Command) {
  vector<string> PathAndArgs;
  string Arguments = "";
  BOOLEAN IsFirstCommand = FALSE;
  BOOLEAN PathIgnored = FALSE;
  BOOLEAN IsNextPath = FALSE;
#if ActivateUserModeDebugger == \
    FALSE if (!g_IsSerialConnectedToRemoteDebugger) {
  ShowMessages(
      "the user-mode debugger in VMI Mode is still in the beta version and "
      "not stable. "
      "we decided to exclude it from this release and release it in future "
      "versions. "
      "if you want to test the user-mode debugger in VMI Mode, you should "
      "build "
      "HyperDbg with special instructions. But starting processes is fully "
      "supported "
      "in the Debugger Mode.\n"
      "(it's not recommended to use it in VMI Mode yet!)\n");
  return;
}
#endif
if (CommandTokens.size() <= 2) {
  ShowMessages(
      "incorrect use of the '%s'\n\n",
      GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
  CommandStartHelp();
  return;
}
if (!CompareLowerCaseStrings(CommandTokens.at(1), "path")) {
  ShowMessages(
      "err, couldn't resolve error at '%s'\n\n",
      GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
  CommandStartHelp();
  return;
}
for (auto Section : CommandTokens) {
  if (!IsFirstCommand) {
    IsFirstCommand = TRUE;
    continue;
  }
  if (!PathIgnored) {
    PathIgnored = TRUE;
    IsNextPath = TRUE;
    continue;
  }
  if (IsNextPath) {
    IsNextPath = FALSE;
    StringToWString(g_StartCommandPath,
                    GetCaseSensitiveStringFromCommandToken(Section));
    continue;
  }
  if (GetCaseSensitiveStringFromCommandToken(Section).find(' ') !=
      std::string::npos) {
    Arguments += "\"" + GetCaseSensitiveStringFromCommandToken(Section) + "\" ";
  } else {
    Arguments += GetCaseSensitiveStringFromCommandToken(Section) + " ";
  }
}
if (Arguments.empty()) {
  UdAttachToProcess(NULL, g_StartCommandPath.c_str(), NULL, FALSE);
} else {
  Arguments.pop_back();
  StringToWString(g_StartCommandPathAndArguments, Arguments);
  UdAttachToProcess(NULL, g_StartCommandPath.c_str(),
                    (WCHAR *)g_StartCommandPathAndArguments.c_str(), FALSE);
}
}
