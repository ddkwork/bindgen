
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandSymHelp() {
  ShowMessages(".sym : performs the symbol actions.\n\n");
  ShowMessages("syntax : \t.sym [table]\n");
  ShowMessages("syntax : \t.sym [reload] [pid ProcessId (hex)]\n");
  ShowMessages("syntax : \t.sym [download]\n");
  ShowMessages("syntax : \t.sym [load]\n");
  ShowMessages("syntax : \t.sym [unload]\n");
  ShowMessages(
      "syntax : \t.sym [add] [base Address (hex)] [path Path (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .sym table\n");
  ShowMessages("\t\te.g : .sym reload\n");
  ShowMessages("\t\te.g : .sym reload pid 3a24\n");
  ShowMessages("\t\te.g : .sym load\n");
  ShowMessages("\t\te.g : .sym download\n");
  ShowMessages(
      "\t\te.g : .sym add base fffff8077356000 path c:\\symbols\\my_dll.pdb\n");
  ShowMessages(
      "\t\te.g : .sym add base fffff8077356000 path \"c:\\symbols "
      "files\\my_dll.pdb\"\n");
  ShowMessages("\t\te.g : .sym unload\n");
}

VOID CommandSym(vector<CommandToken> CommandTokens, string Command) {
  UINT64 BaseAddress = NULL;
  UINT32 UserProcessId = NULL;
  string PathToPdb;
  if (CommandTokens.size() == 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandSymHelp();
    return;
  }
  if (CompareLowerCaseStrings(CommandTokens.at(1), "table")) {
    if (CommandTokens.size() != 2) {
      ShowMessages(
          "incorrect use of the '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      CommandSymHelp();
      return;
    }
    SymbolBuildAndShowSymbolTable();
  } else if (CompareLowerCaseStrings(CommandTokens.at(1), "load") ||
             CompareLowerCaseStrings(CommandTokens.at(1), "download")) {
    if (CommandTokens.size() != 2) {
      ShowMessages(
          "incorrect use of the '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      CommandSymHelp();
      return;
    }
    if (CompareLowerCaseStrings(CommandTokens.at(1), "load")) {
      SymbolLoadOrDownloadSymbols(FALSE, FALSE);
    } else if (CompareLowerCaseStrings(CommandTokens.at(1), "download")) {
      SymbolLoadOrDownloadSymbols(TRUE, FALSE);
    }
  } else if (CompareLowerCaseStrings(CommandTokens.at(1), "reload")) {
    if (CommandTokens.size() != 2 && CommandTokens.size() != 4) {
      ShowMessages(
          "incorrect use of the '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      CommandSymHelp();
      return;
    }
    if (CommandTokens.size() == 4) {
      if (CompareLowerCaseStrings(CommandTokens.at(2), "pid")) {
        if (!ConvertTokenToUInt32(CommandTokens.at(3), &UserProcessId)) {
          ShowMessages(
              "err, couldn't resolve error at '%s'\n\n",
              GetCaseSensitiveStringFromCommandToken(CommandTokens.at(3))
                  .c_str());
          CommandSymHelp();
          return;
        }
      } else {
        ShowMessages("incorrect use of the '%s'\n\n",
                     GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0))
                         .c_str());
        CommandSymHelp();
        return;
      }
    }
    if (g_IsSerialConnectedToRemoteDebuggee) {
      SymbolReloadSymbolTableInDebuggerMode(UserProcessId);
    } else {
      if (UserProcessId == NULL) {
        if (g_ActiveProcessDebuggingState.IsActive) {
          UserProcessId = g_ActiveProcessDebuggingState.ProcessId;
        } else {
          UserProcessId = GetCurrentProcessId();
        }
      }
      if (SymbolLocalReload(UserProcessId)) {
        ShowMessages("symbol table updated successfully\n");
      }
    }
  } else if (CompareLowerCaseStrings(CommandTokens.at(1), "unload")) {
    if (CommandTokens.size() != 2) {
      ShowMessages(
          "incorrect use of the '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      CommandSymHelp();
      return;
    }
    ScriptEngineUnloadAllSymbolsWrapper();
  } else if (CompareLowerCaseStrings(CommandTokens.at(1), "add")) {
    if (CommandTokens.size() < 6) {
      ShowMessages(
          "incorrect use of the '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      CommandSymHelp();
      return;
    }
    if (CompareLowerCaseStrings(CommandTokens.at(2), "base")) {
      if (!ConvertTokenToUInt64(CommandTokens.at(3), &BaseAddress)) {
        ShowMessages(
            "please add a valid hex address to be used as the base "
            "address\n\n");
        CommandSymHelp();
        return;
      }
      if (!CompareLowerCaseStrings(CommandTokens.at(4), "path")) {
        ShowMessages("incorrect use of the '%s'\n\n",
                     GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0))
                         .c_str());
        CommandSymHelp();
        return;
      }
      PathToPdb = GetCaseSensitiveStringFromCommandToken(CommandTokens.at(5));
      if (!IsFileExistA(PathToPdb.c_str())) {
        ShowMessages("pdb file not found\n");
        return;
      }
      ShowMessages("loading module symbol at '%s'\n", PathToPdb.c_str());
      ScriptEngineLoadFileSymbolWrapper(BaseAddress, PathToPdb.c_str(), NULL);
    } else {
      ShowMessages(
          "incorrect use of the '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      CommandSymHelp();
      return;
    }
  } else {
    ShowMessages(
        "unknown parameter at '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
    CommandSymHelp();
    return;
  }
}
