
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsInstrumentingInstructions;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandIHelp() {
  ShowMessages(
      "i : executes a single instruction (step-in) and guarantees that no "
      "other instruction is executed other than the displayed instruction "
      "including user to the kernel (syscalls) and kernel to the user "
      "(sysrets) and exceptions and page-faults and optionally displays all "
      "registers and flags' resulting values.\n\n");
  ShowMessages("syntax : \ti\n");
  ShowMessages("syntax : \ti [Count (hex)]\n");
  ShowMessages("syntax : \tir\n");
  ShowMessages("syntax : \tir [Count (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : i\n");
  ShowMessages("\t\te.g : ir\n");
  ShowMessages("\t\te.g : ir 1f\n");
}

VOID CommandI(vector<CommandToken> CommandTokens, string Command) {
  UINT32 StepCount;
  if (CommandTokens.size() != 1 && CommandTokens.size() != 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandIHelp();
    return;
  }
  if (g_ActiveProcessDebuggingState.IsActive) {
    ShowMessages(
        "the instrumentation step-in is only supported in Debugger Mode\n");
    return;
  }
  if (CommandTokens.size() == 2) {
    if (!ConvertTokenToUInt32(CommandTokens.at(1), &StepCount)) {
      ShowMessages("please specify a correct hex value for [count]\n\n");
      CommandIHelp();
      return;
    }
  } else {
    StepCount = 1;
  }
  if (g_IsSerialConnectedToRemoteDebuggee) {
    g_IsInstrumentingInstructions = TRUE;
    for (size_t i = 0; i < StepCount; i++) {
      SteppingInstrumentationStepIn();
      if (CompareLowerCaseStrings(CommandTokens.at(0), "ir")) {
        HyperDbgRegisterShowAll();
        if (i != StepCount - 1) {
          ShowMessages("\n");
        }
      }
      if (!g_IsInstrumentingInstructions) {
        break;
      }
    }
    g_IsInstrumentingInstructions = FALSE;
  } else {
    ShowMessages(
        "err, stepping (i) is not valid in the current context, you "
        "should connect to a debuggee\n");
  }
}
