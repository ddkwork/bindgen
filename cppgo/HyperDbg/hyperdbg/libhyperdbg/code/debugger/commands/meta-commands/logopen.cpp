
#include "pch.h"
using namespace std;
extern BOOLEAN g_LogOpened;
extern ofstream g_LogOpenFile;

VOID CommandLogopenHelp() {
  ShowMessages(".logopen : saves commands and results in a file.\n\n");
  ShowMessages("syntax : \t.logopen [FilePath (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .logopen c:\\users\\sina\\desktop\\log.txt\n");
  ShowMessages(
      "\t\te.g : .logopen \"c:\\users\\sina\\desktop\\log with space.txt\"\n");
}

VOID CommandLogopen(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() != 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandLogopenHelp();
    return;
  }
  if (g_LogOpened) {
    ShowMessages(
        "log was opened previously, you have the close it first "
        "(using .logclose)\n");
    return;
  }
  g_LogOpenFile.open(
      GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
  if (g_LogOpenFile.is_open()) {
    g_LogOpened = TRUE;
    time_t t = time(NULL);
    struct tm tm = *localtime(&t);
    ShowMessages(
        "save commands and results into file : %s (%d-%02d-%02d "
        "%02d:%02d:%02d)\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str(),
        tm.tm_year + 1900, tm.tm_mon + 1, tm.tm_mday, tm.tm_hour, tm.tm_min,
        tm.tm_sec);
  } else {
    ShowMessages(
        "unable to open file : %s\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
    return;
  }
}

VOID LogopenSaveToFile(const char *Text) { g_LogOpenFile << Text; }
