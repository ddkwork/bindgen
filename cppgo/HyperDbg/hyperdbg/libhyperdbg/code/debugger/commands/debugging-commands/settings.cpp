
#include "pch.h"
extern BOOLEAN g_AutoUnpause;
extern BOOLEAN g_AutoFlush;
extern BOOLEAN g_AddressConversion;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern UINT32 g_DisassemblerSyntax;

VOID CommandSettingsHelp() {
  ShowMessages(
      "settings : queries, sets, or changes a value for a special settings "
      "option.\n\n");
  ShowMessages("syntax : \tsettings [OptionName (string)]\n");
  ShowMessages("syntax : \tsettings [OptionName (string)] [Value (hex)]\n");
  ShowMessages("syntax : \tsettings [OptionName (string)] [Value (string)]\n");
  ShowMessages("syntax : \tsettings [OptionName (string)] [on|off]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : settings autounpause\n");
  ShowMessages("\t\te.g : settings autounpause on\n");
  ShowMessages("\t\te.g : settings autounpause off\n");
  ShowMessages("\t\te.g : settings addressconversion on\n");
  ShowMessages("\t\te.g : settings addressconversion off\n");
  ShowMessages("\t\te.g : settings autoflush on\n");
  ShowMessages("\t\te.g : settings autoflush off\n");
  ShowMessages("\t\te.g : settings syntax intel\n");
  ShowMessages("\t\te.g : settings syntax att\n");
  ShowMessages("\t\te.g : settings syntax masm\n");
}

BOOLEAN CommandSettingsGetValueFromConfigFile(std::string OptionName,
                                              std::string &OptionValue) {
  inipp::Ini<char> Ini;
  WCHAR ConfigPath[MAX_PATH] = {0};
  std::string OptionValueFromFile;
  GetConfigFilePath(ConfigPath);
  if (!IsFileExistW(ConfigPath)) {
    return FALSE;
  }
  ifstream Is(ConfigPath);
  Ini.parse(Is);
  inipp::get_value(Ini.sections["DEFAULT"], OptionName, OptionValueFromFile);
  Is.close();
  if (!OptionValueFromFile.empty()) {
    OptionValue = OptionValueFromFile;
    return TRUE;
  } else {
    return FALSE;
  }
}

VOID CommandSettingsSetValueFromConfigFile(std::string OptionName,
                                           std::string OptionValue) {
  inipp::Ini<char> Ini;
  WCHAR ConfigPath[MAX_PATH] = {0};
  GetConfigFilePath(ConfigPath);
  ifstream Is(ConfigPath);
  Ini.parse(Is);
  Is.close();
  Ini.sections["DEFAULT"][OptionName] = OptionValue.c_str();
  Ini.interpolate();
  ofstream Os(ConfigPath);
  Ini.generate(Os);
  Os.close();
}

VOID CommandSettingsLoadDefaultValuesFromConfigFile() {
  string OptionValue;
  if (CommandSettingsGetValueFromConfigFile("AsmSyntax", OptionValue)) {
    if (!OptionValue.compare("intel")) {
      g_DisassemblerSyntax = 1;
    } else if (!OptionValue.compare("att") || !OptionValue.compare("at&t")) {
      g_DisassemblerSyntax = 2;
    } else if (!OptionValue.compare("masm")) {
      g_DisassemblerSyntax = 3;
    } else {
      ShowMessages("err, incorrect assembly syntax settings\n");
    }
  }
  if (CommandSettingsGetValueFromConfigFile("AutoUnpause", OptionValue)) {
    if (!OptionValue.compare("on")) {
      g_AutoUnpause = TRUE;
    } else if (!OptionValue.compare("off")) {
      g_AutoUnpause = FALSE;
    } else {
      ShowMessages("err, incorrect auto unpause settings\n");
    }
  }
  if (CommandSettingsGetValueFromConfigFile("AutoFlush", OptionValue)) {
    if (!OptionValue.compare("on")) {
      g_AutoFlush = TRUE;
    } else if (!OptionValue.compare("off")) {
      g_AutoFlush = FALSE;
    } else {
      ShowMessages("err, incorrect auto flush settings\n");
    }
  }
  if (CommandSettingsGetValueFromConfigFile("AddrConv", OptionValue)) {
    if (!OptionValue.compare("on")) {
      g_AddressConversion = TRUE;
    } else if (!OptionValue.compare("off")) {
      g_AddressConversion = FALSE;
    } else {
      ShowMessages("err, incorrect address conversion settings\n");
    }
  }
}

VOID CommandSettingsAddressConversion(vector<CommandToken> CommandTokens) {
  if (CommandTokens.size() == 2) {
    if (g_AddressConversion) {
      ShowMessages("address conversion is enabled\n");
    } else {
      ShowMessages("address conversion is disabled\n");
    }
  } else if (CommandTokens.size() == 3) {
    if (CompareLowerCaseStrings(CommandTokens.at(2), "on")) {
      g_AddressConversion = TRUE;
      CommandSettingsSetValueFromConfigFile("AddrConv", "on");
      ShowMessages("set address conversion to enabled\n");
    } else if (CompareLowerCaseStrings(CommandTokens.at(2), "off")) {
      g_AddressConversion = FALSE;
      CommandSettingsSetValueFromConfigFile("AddrConv", "off");
      ShowMessages("set address conversion to disabled\n");
    } else {
      ShowMessages(
          "incorrect use of the '%s', please use 'help %s' for more "
          "information\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      return;
    }
  } else {
    ShowMessages(
        "incorrect use of the '%s', please use 'help %s' for more "
        "information\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    return;
  }
}

VOID CommandSettingsAutoFlush(vector<CommandToken> CommandTokens) {
  if (CommandTokens.size() == 2) {
    if (g_AutoFlush) {
      ShowMessages("auto-flush is enabled\n");
    } else {
      ShowMessages("auto-flush is disabled\n");
    }
  } else if (CommandTokens.size() == 3) {
    if (CompareLowerCaseStrings(CommandTokens.at(2), "on")) {
      g_AutoFlush = TRUE;
      CommandSettingsSetValueFromConfigFile("AutoFlush", "on");
      ShowMessages("set auto-flush to enabled\n");
    } else if (CompareLowerCaseStrings(CommandTokens.at(2), "off")) {
      g_AutoFlush = FALSE;
      CommandSettingsSetValueFromConfigFile("AutoFlush", "off");
      ShowMessages("set auto-flush to disabled\n");
    } else {
      ShowMessages(
          "incorrect use of the '%s', please use 'help %s' for more "
          "information\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      return;
    }
  } else {
    ShowMessages(
        "incorrect use of the '%s', please use 'help %s' for more "
        "information\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    return;
  }
}

VOID CommandSettingsAutoUpause(vector<CommandToken> CommandTokens) {
  if (CommandTokens.size() == 2) {
    if (g_AutoUnpause) {
      ShowMessages("auto-unpause is enabled\n");
    } else {
      ShowMessages("auto-unpause is disabled\n");
    }
  } else if (CommandTokens.size() == 3) {
    if (CompareLowerCaseStrings(CommandTokens.at(2), "on")) {
      g_AutoUnpause = TRUE;
      CommandSettingsSetValueFromConfigFile("AutoUnpause", "on");
      ShowMessages("set auto-unpause to enabled\n");
    } else if (CompareLowerCaseStrings(CommandTokens.at(2), "off")) {
      g_AutoUnpause = FALSE;
      CommandSettingsSetValueFromConfigFile("AutoUnpause", "off");
      ShowMessages("set auto-unpause to disabled\n");
    } else {
      ShowMessages(
          "incorrect use of the '%s', please use 'help %s' for more "
          "information\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      return;
    }
  } else {
    ShowMessages(
        "incorrect use of the '%s', please use 'help %s' for more "
        "information\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    return;
  }
}

VOID CommandSettingsSyntax(vector<CommandToken> CommandTokens) {
  if (CommandTokens.size() == 2) {
    if (g_DisassemblerSyntax == 1) {
      ShowMessages("disassembler syntax is : intel\n");
    } else if (g_DisassemblerSyntax == 2) {
      ShowMessages("disassembler syntax is : at&t\n");
    } else if (g_DisassemblerSyntax == 3) {
      ShowMessages("disassembler syntax is : masm\n");
    } else {
      ShowMessages("unknown syntax\n");
    }
  } else if (CommandTokens.size() == 3) {
    if (CompareLowerCaseStrings(CommandTokens.at(2), "intel")) {
      g_DisassemblerSyntax = 1;
      CommandSettingsSetValueFromConfigFile("AsmSyntax", "intel");
      ShowMessages("set syntax to intel\n");
    } else if (CompareLowerCaseStrings(CommandTokens.at(2), "att") ||
               CompareLowerCaseStrings(CommandTokens.at(2), "at&t")) {
      g_DisassemblerSyntax = 2;
      CommandSettingsSetValueFromConfigFile("AsmSyntax", "att");
      ShowMessages("set syntax to at&t\n");
    } else if (CompareLowerCaseStrings(CommandTokens.at(2), "masm")) {
      g_DisassemblerSyntax = 3;
      CommandSettingsSetValueFromConfigFile("AsmSyntax", "masm");
      ShowMessages("set syntax to masm\n");
    } else {
      ShowMessages(
          "incorrect use of the '%s', please use 'help %s' for more "
          "information\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      return;
    }
  } else {
    ShowMessages(
        "incorrect use of the '%s', please use 'help %s' for more "
        "information\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    return;
  }
}

VOID CommandSettings(vector<CommandToken> CommandTokens, string Command) {
  if (CommandTokens.size() <= 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandSettingsHelp();
    return;
  }
  if (CompareLowerCaseStrings(CommandTokens.at(1), "autounpause")) {
    CommandSettingsAutoUpause(CommandTokens);
  } else if (CompareLowerCaseStrings(CommandTokens.at(1), "syntax")) {
    if (g_IsConnectedToRemoteDebuggee) {
      RemoteConnectionSendCommand(Command.c_str(),
                                  (UINT32)Command.length() + 1);
    } else {
      CommandSettingsSyntax(CommandTokens);
    }
  } else if (CompareLowerCaseStrings(CommandTokens.at(1), "autoflush")) {
    if (g_IsConnectedToRemoteDebuggee) {
      RemoteConnectionSendCommand(Command.c_str(),
                                  (UINT32)Command.length() + 1);
    } else {
      CommandSettingsAutoFlush(CommandTokens);
    }
  } else if (CompareLowerCaseStrings(CommandTokens.at(1),
                                     "addressconversion")) {
    if (g_IsConnectedToRemoteDebuggee) {
      RemoteConnectionSendCommand(Command.c_str(),
                                  (UINT32)Command.length() + 1);
    } else {
      CommandSettingsAddressConversion(CommandTokens);
    }
  } else {
    ShowMessages(
        "incorrect use of the '%s', please use 'help %s' for more "
        "information\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str(),
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    return;
  }
}
