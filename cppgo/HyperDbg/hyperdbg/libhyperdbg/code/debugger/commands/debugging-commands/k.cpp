
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsRunningInstruction32Bit;

VOID CommandKHelp() {
  ShowMessages("k, kd, kq : shows the callstack of the thread.\n\n");
  ShowMessages("syntax : \tk\n");
  ShowMessages("syntax : \tkd\n");
  ShowMessages("syntax : \tkq\n");
  ShowMessages("syntax : \tk [base StackAddress (hex)] [l Length (hex)]\n");
  ShowMessages("syntax : \tkd [base StackAddress (hex)] [l Length (hex)]\n");
  ShowMessages("syntax : \tkq [base StackAddress (hex)] [l Length (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : k\n");
  ShowMessages("\t\te.g : k l 100\n");
  ShowMessages("\t\te.g : kd base 0x77356f010\n");
  ShowMessages("\t\te.g : kq base fffff8077356f010\n");
  ShowMessages("\t\te.g : kq base @rbx-10\n");
  ShowMessages("\t\te.g : kq base fffff8077356f010 l 100\n");
}

VOID CommandK(vector<CommandToken> CommandTokens, string Command) {
  UINT64 BaseAddress = NULL;
  UINT32 Length = PAGE_SIZE;
  BOOLEAN IsFirstCommand = TRUE;
  BOOLEAN IsNextBase = FALSE;
  BOOLEAN IsNextLength = FALSE;
  if (CommandTokens.size() >= 6) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandKHelp();
    return;
  }
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, tracing callstack is not possible when you're not "
        "connected to a debuggee\n");
    return;
  }
  if (g_IsRunningInstruction32Bit) {
    Length = PAGE_SIZE;
  } else {
    Length = PAGE_SIZE * 2;
  }
  for (auto Section : CommandTokens) {
    if (IsFirstCommand) {
      IsFirstCommand = FALSE;
      continue;
    }
    if (IsNextBase == TRUE) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(Section), &BaseAddress)) {
        ShowMessages("err, couldn't resolve error at '%s'\n",
                     GetCaseSensitiveStringFromCommandToken(Section).c_str());
        return;
      }
      IsNextBase = FALSE;
      continue;
    }
    if (IsNextLength == TRUE) {
      if (!ConvertTokenToUInt32(Section, &Length)) {
        ShowMessages("err, you should enter a valid length\n\n");
        return;
      }
      IsNextLength = FALSE;
      continue;
    }
    if (CompareLowerCaseStrings(Section, "l")) {
      IsNextLength = TRUE;
      continue;
    }
    if (CompareLowerCaseStrings(Section, "base")) {
      IsNextBase = TRUE;
      continue;
    }
    ShowMessages(
        "err, incorrect use of the '%s' command\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandKHelp();
    return;
  }
  if (IsNextLength || IsNextBase) {
    ShowMessages(
        "incorrect use of the '%s' command\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandKHelp();
    return;
  }
  if (CompareLowerCaseStrings(CommandTokens.at(0), "k")) {
    KdSendCallStackPacketToDebuggee(
        BaseAddress, Length, DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITHOUT_PARAMS,
        g_IsRunningInstruction32Bit);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "kq")) {
    KdSendCallStackPacketToDebuggee(
        BaseAddress, Length, DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITH_PARAMS,
        FALSE);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "kd")) {
    KdSendCallStackPacketToDebuggee(
        BaseAddress, Length, DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITH_PARAMS,
        TRUE);
  }
}
