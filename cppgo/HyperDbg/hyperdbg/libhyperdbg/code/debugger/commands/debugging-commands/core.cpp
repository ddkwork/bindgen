
#include "pch.h"
extern ULONG g_CurrentRemoteCore;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandCoreHelp() {
  ShowMessages("~ : shows and changes the operating processor.\n\n");
  ShowMessages("syntax : \t~\n");
  ShowMessages("syntax : \t~ [CoreNumber (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : ~ \n");
  ShowMessages("\t\te.g : ~ 2 \n");
}

VOID CommandCore(vector<CommandToken> CommandTokens, string Command) {
  UINT32 TargetCore = 0;
  if (CommandTokens.size() != 1 && CommandTokens.size() != 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandCoreHelp();
    return;
  }
  if (!g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages("err, you're not connected to any debuggee\n");
    return;
  }
  if (CommandTokens.size() == 1) {
    ShowMessages("current processor : 0x%x\n", g_CurrentRemoteCore);
  } else if (CommandTokens.size() == 2) {
    if (!ConvertTokenToUInt32(CommandTokens.at(1), &TargetCore)) {
      ShowMessages(
          "please specify a correct hex value for the core that you "
          "want to operate on it\n\n");
      CommandCoreHelp();
      return;
    }
    KdSendSwitchCorePacketToDebuggee(TargetCore);
  }
}
