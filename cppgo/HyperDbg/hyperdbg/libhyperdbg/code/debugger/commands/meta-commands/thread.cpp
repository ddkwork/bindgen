
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandThreadHelp() {
  ShowMessages(
      ".thread, .thread2 : shows and changes the threads. "
      "This command needs public symbols for 'ntoskrnl.exe' if "
      "you want to see the threads list. Please visit the "
      "documentation to know about the difference between '.thread' "
      "and '.thread2'.\n\n");
  ShowMessages("syntax : \t.thread\n");
  ShowMessages("syntax : \t.thread [list] [process Eprocess (hex)]\n");
  ShowMessages("syntax : \t.thread [tid ThreadId (hex)]\n");
  ShowMessages("syntax : \t.thread [thread Ethread (hex)]\n");
  ShowMessages("syntax : \t.thread2 [tid ThreadId (hex)]\n");
  ShowMessages("syntax : \t.thread2 [thread Ethread (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : .thread\n");
  ShowMessages("\t\te.g : .thread tid 48a4\n");
  ShowMessages("\t\te.g : .thread2 tid 48a4\n");
  ShowMessages("\t\te.g : .thread thread ffff948c`c8970200\n");
  ShowMessages("\t\te.g : .thread list\n");
  ShowMessages("\t\te.g : .thread list process ffff948c`a1279880\n");
}

BOOLEAN CommandThreadListThreads(UINT64 Eprocess) {
  UINT32 ThreadListHeadOffset = 0;
  UINT32 ThreadListEntryOffset = 0;
  UINT32 CidOffset = 0;
  UINT32 OffsetOfActiveProcessLinks = 0;
  UINT64 AddressOfActiveProcessHead = 0;
  DEBUGGEE_THREAD_LIST_NEEDED_DETAILS ThreadListNeededItems = {0};
  if (ScriptEngineGetFieldOffsetWrapper((CHAR *)"nt!_EPROCESS",
                                        (CHAR *)"ThreadListHead",
                                        &ThreadListHeadOffset) &&
      ScriptEngineGetFieldOffsetWrapper((CHAR *)"nt!_ETHREAD",
                                        (CHAR *)"ThreadListEntry",
                                        &ThreadListEntryOffset) &&
      ScriptEngineGetFieldOffsetWrapper((CHAR *)"nt!_ETHREAD", (CHAR *)"Cid",
                                        &CidOffset) &&
      ScriptEngineGetFieldOffsetWrapper((CHAR *)"nt!_EPROCESS",
                                        (CHAR *)"ActiveProcessLinks",
                                        &OffsetOfActiveProcessLinks) &&
      SymbolConvertNameOrExprToAddress("nt!PsActiveProcessHead",
                                       &AddressOfActiveProcessHead)) {
    ThreadListNeededItems.ThreadListHeadOffset = ThreadListHeadOffset;
    ThreadListNeededItems.ThreadListEntryOffset = ThreadListEntryOffset;
    ThreadListNeededItems.CidOffset = CidOffset;
    ThreadListNeededItems.ActiveProcessLinksOffset = OffsetOfActiveProcessLinks;
    ThreadListNeededItems.PsActiveProcessHead = AddressOfActiveProcessHead;
    ThreadListNeededItems.Process = Eprocess;
    if (!g_IsSerialConnectedToRemoteDebuggee) {
      ObjectShowProcessesOrThreadList(FALSE, NULL, NULL,
                                      &ThreadListNeededItems);
    } else {
      KdSendSwitchThreadPacketToDebuggee(
          DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_LIST, NULL, NULL, FALSE,
          &ThreadListNeededItems);
    }
    return TRUE;
  } else {
    return FALSE;
  }
}

VOID CommandThread(vector<CommandToken> CommandTokens, string Command) {
  UINT32 TargetThreadId = 0;
  UINT64 TargetThread = 0;
  UINT64 TargetProcess = 0;
  BOOLEAN CheckByClkIntr = FALSE;
  if (CommandTokens.size() >= 5) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandThreadHelp();
    return;
  }
  if (CommandTokens.size() == 1) {
    if (!g_IsSerialConnectedToRemoteDebuggee) {
      ObjectShowProcessesOrThreadDetails(FALSE);
    } else {
      KdSendSwitchThreadPacketToDebuggee(
          DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS, NULL, NULL,
          FALSE, NULL);
    }
  } else if (CommandTokens.size() == 2) {
    if (CompareLowerCaseStrings(CommandTokens.at(1), "list")) {
      if (!CommandThreadListThreads(NULL)) {
        ShowMessages(
            "err, the need offset to iterate over threads not found, "
            "make sure to load ntoskrnl.exe's PDB file. use '.help .sym' for "
            "more information\n");
        return;
      }
    } else {
      ShowMessages(
          "err, unknown parameter at '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
      CommandThreadHelp();
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
    if (CompareLowerCaseStrings(CommandTokens.at(1), "tid")) {
      if (!ConvertTokenToUInt32(CommandTokens.at(2), &TargetThreadId)) {
        ShowMessages(
            "please specify a correct hex value for the thread id that you "
            "want to operate on it\n\n");
        CommandThreadHelp();
        return;
      }
    } else if (CompareLowerCaseStrings(CommandTokens.at(1), "thread")) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)),
              &TargetThread)) {
        ShowMessages(
            "please specify a correct hex value for the thread (nt!_ETHREAD) "
            "that you "
            "want to operate on it\n\n");
        CommandThreadHelp();
        return;
      }
    } else if (CompareLowerCaseStrings(CommandTokens.at(1), "list") &&
               CompareLowerCaseStrings(CommandTokens.at(2), "process")) {
      ShowMessages("please specify a hex value for the process\n\n");
      CommandThreadHelp();
      return;
    } else {
      ShowMessages(
          "err, unknown parameter at '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2)).c_str());
      CommandThreadHelp();
      return;
    }
    if (CompareLowerCaseStrings(CommandTokens.at(0), ".thread2")) {
      CheckByClkIntr = FALSE;
    } else {
      CheckByClkIntr = TRUE;
    }
    KdSendSwitchThreadPacketToDebuggee(
        DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH, TargetThreadId,
        TargetThread, CheckByClkIntr, NULL);
  } else if (CommandTokens.size() == 4) {
    if (CompareLowerCaseStrings(CommandTokens.at(1), "list")) {
      if (CompareLowerCaseStrings(CommandTokens.at(2), "process")) {
        if (!SymbolConvertNameOrExprToAddress(
                GetCaseSensitiveStringFromCommandToken(CommandTokens.at(3)),
                &TargetProcess)) {
          ShowMessages(
              "please specify a correct hex value for the process "
              "(nt!_EPROCESS) that you "
              "want to see its threads\n\n");
          CommandThreadHelp();
          return;
        }
        if (!CommandThreadListThreads(TargetProcess)) {
          ShowMessages(
              "err, the need offset to iterate over threads not found, "
              "make sure to load ntoskrnl.exe's PDB file. use '.help .sym' for "
              "more information\n");
          return;
        }
      } else {
        ShowMessages("err, unknown parameter at '%s'\n\n",
                     GetCaseSensitiveStringFromCommandToken(CommandTokens.at(2))
                         .c_str());
        CommandThreadHelp();
        return;
      }
    } else {
      ShowMessages(
          "err, unknown parameter at '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(1)).c_str());
      CommandThreadHelp();
      return;
    }
  } else {
    ShowMessages("invalid parameter\n\n");
    CommandThreadHelp();
    return;
  }
}
