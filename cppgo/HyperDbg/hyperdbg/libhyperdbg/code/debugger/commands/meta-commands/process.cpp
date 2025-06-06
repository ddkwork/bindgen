
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandProcessHelp() {
  ShowMessages(
      ".process, .process2 : shows and changes the processes. "
      "This command needs public symbols for ntoskrnl.exe if "
      "you want to see the processes list. Please visit the "
      "documentation to know about the difference between '.process' "
      "and '.process2'.\n\n");
  ShowMessages("syntax : \t.process\n");
  ShowMessages("syntax : \t.process [list]\n");
  ShowMessages("syntax : \t.process [pid ProcessId (hex)]\n");
  ShowMessages("syntax : \t.process [process Eprocess (hex)]\n");
  ShowMessages("syntax : \t.process2 [pid ProcessId (hex)]\n");
  ShowMessages("syntax : \t.process2 [process Eprocess (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .process\n");
  ShowMessages("\t\te.g : .process list\n");
  ShowMessages("\t\te.g : .process pid 4\n");
  ShowMessages("\t\te.g : .process2 pid 4\n");
  ShowMessages("\t\te.g : .process process ffff948c`c2349280\n");
}

VOID CommandProcess(vector<CommandToken> CommandTokens, string Command) {
  UINT32 TargetProcessId = 0;
  UINT64 TargetProcess = 0;
  UINT64 AddressOfActiveProcessHead = 0;
  UINT32 OffsetOfImageFileName = 0;
  UINT32 OffsetOfUniqueProcessId = 0;
  UINT32 OffsetOfActiveProcessLinks = 0;
  BOOLEAN ResultOfGettingOffsets = FALSE;
  BOOLEAN IsSetByClkIntr = FALSE;
  DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS ProcessListNeededItems = {0};
  if (CommandTokens.size() >= 4) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandProcessHelp();
    return;
  }
  if (CommandTokens.size() == 1) {
    if (!g_IsSerialConnectedToRemoteDebuggee) {
      ObjectShowProcessesOrThreadDetails(TRUE);
    } else {
      KdSendSwitchProcessPacketToDebuggee(
          DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS, NULL, NULL,
          FALSE, NULL);
    }
  } else if (CommandTokens.size() == 2) {
    if (CompareLowerCaseStrings(CommandTokens.at(1), "list")) {
      if (ScriptEngineGetFieldOffsetWrapper((CHAR *)"nt!_EPROCESS",
                                            (CHAR *)"ActiveProcessLinks",
                                            &OffsetOfActiveProcessLinks) &&
          ScriptEngineGetFieldOffsetWrapper((CHAR *)"nt!_EPROCESS",
                                            (CHAR *)"ImageFileName",
                                            &OffsetOfImageFileName) &&
          ScriptEngineGetFieldOffsetWrapper((CHAR *)"nt!_EPROCESS",
                                            (CHAR *)"UniqueProcessId",
                                            &OffsetOfUniqueProcessId) &&
          SymbolConvertNameOrExprToAddress("nt!PsActiveProcessHead",
                                           &AddressOfActiveProcessHead)) {
        ProcessListNeededItems.PsActiveProcessHead = AddressOfActiveProcessHead;
        ProcessListNeededItems.ActiveProcessLinksOffset =
            OffsetOfActiveProcessLinks;
        ProcessListNeededItems.ImageFileNameOffset = OffsetOfImageFileName;
        ProcessListNeededItems.UniquePidOffset = OffsetOfUniqueProcessId;
        if (!g_IsSerialConnectedToRemoteDebuggee) {
          ObjectShowProcessesOrThreadList(TRUE, &ProcessListNeededItems, NULL,
                                          NULL);
        } else {
          KdSendSwitchProcessPacketToDebuggee(
              DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_LIST, NULL, NULL,
              FALSE, &ProcessListNeededItems);
        }
      } else {
        ShowMessages(
            "err, the need offset to iterate over processes not found, "
            "make sure to load ntoskrnl.exe's PDB file. use '.help .sym' for "
            "more information\n");
        return;
      }
    } else {
      ShowMessages(
          "err, unknown parameter at '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
      CommandProcessHelp();
      return;
    }
  } else if (CommandTokens.size() == 3) {
    if (!g_IsSerialConnectedToRemoteDebuggee) {
      ShowMessages(
          "err, you're not connected to any debuggee in Debugger Mode, "
          "you can use the '.attach', or the '.detach' commands if you're "
          "operating in VMI Mode\n");
      return;
    }
    if (CompareLowerCaseStrings(CommandTokens.at(1), "pid")) {
      if (!ConvertTokenToUInt32(CommandTokens.at(2), &TargetProcessId)) {
        ShowMessages(
            "please specify a correct hex value for the process id that you "
            "want to operate on it\n\n");
        CommandProcessHelp();
        return;
      }
    } else if (CompareLowerCaseStrings(CommandTokens.at(1), "process")) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)),
              &TargetProcess)) {
        ShowMessages(
            "please specify a correct hex value for the process (nt!_EPROCESS) "
            "that you "
            "want to operate on it\n\n");
        CommandProcessHelp();
        return;
      }
    } else {
      ShowMessages(
          "err, unknown parameter at '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)).c_str());
      CommandProcessHelp();
      return;
    }
    if (CompareLowerCaseStrings(CommandTokens.at(0), ".process2")) {
      IsSetByClkIntr = FALSE;
    } else {
      IsSetByClkIntr = TRUE;
    }
    KdSendSwitchProcessPacketToDebuggee(
        DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH, TargetProcessId,
        TargetProcess, IsSetByClkIntr, NULL);
  } else {
    ShowMessages("invalid parameter\n\n");
    CommandProcessHelp();
    return;
  }
}
