
#include "pch.h"
extern HANDLE g_IsDriverLoadedSuccessfully;
extern HANDLE g_DeviceHandle;
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsDebuggerModulesLoaded;

VOID CommandLoadHelp() {
  ShowMessages("load : installs the drivers and load the modules.\n\n");
  ShowMessages("syntax : \tload [ModuleName (string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : load vmm\n");
}

VOID CommandLoad(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() != 2) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandLoadHelp();
    return;
  }
  if (!g_IsConnectedToHyperDbgLocally) {
    ShowMessages(
        "you're not connected to any instance of HyperDbg, did you "
        "use '.connect'? \n");
    return;
  }
  if (CompareLowerCaseStrings(CommandTokens.at(1), "vmm")) {
    if (g_DeviceHandle) {
      ShowMessages(
          "handle of the driver found, if you use 'load' before, please "
          "first unload it then call 'unload'\n");
      return;
    }
    ShowMessages("loading the vmm driver\n");
    if (HyperDbgInstallVmmDriver() == 1 || HyperDbgLoadVmmModule() == 1) {
      ShowMessages("failed to install or load the driver\n");
      return;
    }
    SymbolLocalReload(GetCurrentProcessId());
  } else {
    ShowMessages("err, module not found\n");
  }
}
