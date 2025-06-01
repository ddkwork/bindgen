
#include "pch.h"
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsDebuggerModulesLoaded;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;

VOID CommandUnloadHelp() {
  ShowMessages(
      "unload : unloads the kernel modules and uninstalls the drivers.\n\n");
  ShowMessages("syntax : \tunload [remove] [ModuleName (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : unload vmm\n");
  ShowMessages("\t\te.g : unload remove vmm\n");
}

VOID CommandUnload(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() != 2 && CommandTokens.size() != 3) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandUnloadHelp();
    return;
  }
  if ((CommandTokens.size() == 2 &&
       CompareLowerCaseStrings(CommandTokens.at(1), "vmm")) ||
      (CommandTokens.size() == 3 &&
       CompareLowerCaseStrings(CommandTokens.at(2), "vmm") &&
       CompareLowerCaseStrings(CommandTokens.at(1), "remove"))) {
    if (!g_IsConnectedToHyperDbgLocally) {
      ShowMessages(
          "you're not connected to any instance of HyperDbg, did you "
          "use '.connect'? \n");
      return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee ||
        g_IsSerialConnectedToRemoteDebugger) {
      ShowMessages(
          "you're connected to a an instance of HyperDbg, please use "
          "'.debug close' command\n");
      return;
    }
    if (g_IsDebuggerModulesLoaded) {
      HyperDbgUnloadVmm();
    } else {
      ShowMessages("there is nothing to unload\n");
    }
    if (CompareLowerCaseStrings(CommandTokens.at(1), "remove")) {
      if (HyperDbgStopVmmDriver()) {
        ShowMessages("err, failed to stop driver\n");
        return;
      }
      if (HyperDbgUninstallVmmDriver()) {
        ShowMessages("err, failed to uninstall the driver\n");
        return;
      }
      ShowMessages("the driver is removed\n");
    }
  } else {
    ShowMessages("err, module not found\n");
  }
}
