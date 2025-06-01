
#include "pch.h"
using namespace std;

VOID CommandPeHelp() {
  ShowMessages(
      ".pe : parses portable executable (PE) files and dump sections.\n\n");
  ShowMessages("syntax : \t.pe [header] [FilePath (string)]\n");
  ShowMessages(
      "syntax : \t.pe [section] [SectionName (string)] [FilePath (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .pe header c:\\reverse\\myfile.exe\n");
  ShowMessages(
      "\t\te.g : .pe section .text \"c:\\reverse files\\myfile.exe\"\n");
  ShowMessages(
      "\t\te.g : .pe section .rdata \"c:\\reverse files\\myfile.exe\"\n");
}

VOID CommandPe(vector<CommandToken> CommandTokens, string Command) {
  BOOLEAN Is32Bit = FALSE;
  wstring Filepath;
  string TempFilePath;
  BOOLEAN ShowDumpOfSection = FALSE;
  if (CommandTokens.size() <= 2) {
    ShowMessages(
        "err, incorrect use of the '%s' command\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandPeHelp();
    return;
  }
  if (CompareLowerCaseStrings(CommandTokens.at(1), "section")) {
    if (CommandTokens.size() == 3) {
      ShowMessages("please specify a valid PE file\n\n");
      CommandPeHelp();
      return;
    }
    ShowDumpOfSection = TRUE;
    TempFilePath = GetCaseSensitiveStringFromCommandToken(CommandTokens.at(3));
  } else if (CompareLowerCaseStrings(CommandTokens.at(1), "header")) {
    ShowDumpOfSection = FALSE;
    TempFilePath = GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2));
  } else {
    ShowMessages(
        "err, couldn't resolve error at '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
    CommandPeHelp();
    return;
  }
  StringToWString(Filepath, TempFilePath);
  if (!PeIsPE32BitOr64Bit(Filepath.c_str(), &Is32Bit)) {
    return;
  }
  if (!ShowDumpOfSection) {
    PeShowSectionInformationAndDump(Filepath.c_str(), NULL, Is32Bit);
  } else {
    PeShowSectionInformationAndDump(
        Filepath.c_str(),
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)).c_str(),
        Is32Bit);
  }
}
