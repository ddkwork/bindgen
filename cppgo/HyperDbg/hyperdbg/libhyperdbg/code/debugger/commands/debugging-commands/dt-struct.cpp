
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandDtHelp() {
  ShowMessages(
      "dt !dt : displays information about a local variable, global "
      "variable or data type.\n\n");
  ShowMessages(
      "If you want to read physical memory then add '!' at the "
      "start of the command\n\n");
  ShowMessages(
      "syntax : \tdt [Module!SymbolName (string)] [AddressExpression (string)] "
      "[pid ProcessId (hex)] [padding Padding (yesno)] [offset Offset (yesno)] "
      "[bitfield Bitfield (yesno)] [native Native (yesno)] [decl Declaration "
      "(yesno)] "
      "[def Definitions (yesno)] [func Functions (yesno)] [pragma Pragma "
      "(yesno)] "
      "[prefix Prefix (string)] [suffix Suffix (string)] [inline Expantion "
      "(string)] "
      "[output FileName (string)]\n\n");
  ShowMessages(
      "syntax : \t!dt [Module!SymbolName (string)] [AddressExpression "
      "(string)] "
      "[padding Padding (yesno)] [offset Offset (yesno)] [bitfield Bitfield "
      "(yesno)] "
      "[native Native (yesno)] [decl Declaration (yesno)] [def Definitions "
      "(yesno)] "
      "[func Functions (yesno)] [pragma Pragma (yesno)] [prefix Prefix "
      "(string)] "
      "[suffix Suffix (string)] [inline Expantion (string)] [output FileName "
      "(string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : dt nt!_EPROCESS\n");
  ShowMessages("\t\te.g : dt nt!_EPROCESS fffff8077356f010\n");
  ShowMessages("\t\te.g : dt nt!_EPROCESS $proc\n");
  ShowMessages("\t\te.g : dt nt!_KPROCESS @rax+@rbx+c0\n");
  ShowMessages("\t\te.g : !dt nt!_EPROCESS 1f0300\n");
  ShowMessages("\t\te.g : dt MyModule!_MY_STRUCT 7ff00040 pid 1420\n");
  ShowMessages("\t\te.g : dt nt!_EPROCESS $proc inline all\n");
  ShowMessages("\t\te.g : dt nt!_EPROCESS fffff8077356f010 inline no\n");
}

VOID CommandStructHelp() {
  ShowMessages(
      "struct : displays a data type, enum, or structure derived from PDB "
      "symbols.\n\n");
  ShowMessages(
      "syntax : struct [Module!SymbolName (string)] [offset Offset (yesno)] "
      "[bitfield Bitfield (yesno)] "
      "[native Native (yesno)] [decl Declaration (yesno)] [def Definitions "
      "(yesno)] "
      "[func Functions (yesno)] [pragma Pragma (yesno)] [prefix Prefix "
      "(string)] "
      "[suffix Suffix (string)] [inline Expantion (string)] [output FileName "
      "(string)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : struct nt!_EPROCESS\n");
  ShowMessages("\t\te.g : struct nt!*\n");
  ShowMessages("\t\te.g : struct nt!* output ntheader.h\n");
  ShowMessages("\t\te.g : struct nt!* func yes output ntheader.h\n");
  ShowMessages("\t\te.g : struct nt!* func yes output ntheader.h\n");
  ShowMessages(
      "\t\te.g : struct nt!* func yes output \"c:\\users\\sina\\desktop\\nt "
      "header.h\"\n");
}

BOOLEAN CommandDtAndStructConvertHyperDbgArgsToPdbex(
    vector<CommandToken> ExtraArgs, std::string &PdbexArgs, UINT32 *ProcessId) {
  UINT32 TargetProcessId = NULL;
  BOOLEAN NextItemIsYesNo = FALSE;
  BOOLEAN NextItemIsString = FALSE;
  BOOLEAN NextItemIsInline = FALSE;
  BOOLEAN NextItemIsFileName = FALSE;
  BOOLEAN NextItemIsProcessId = FALSE;
  PdbexArgs = "";
  for (auto Item : ExtraArgs) {
    if (NextItemIsFileName) {
      PdbexArgs += "\"" + GetCaseSensitiveStringFromCommandToken(Item) + "\" ";
      NextItemIsFileName = FALSE;
      continue;
    }
    if (NextItemIsProcessId) {
      if (!ConvertTokenToUInt32(Item, &TargetProcessId)) {
        ShowMessages("err, you should enter a valid process id\n\n");
        return FALSE;
      }
      NextItemIsProcessId = FALSE;
      continue;
    }
    if (NextItemIsYesNo) {
      if (CompareLowerCaseStrings(Item, "yes")) {
        PdbexArgs += " ";
      } else if (CompareLowerCaseStrings(Item, "no")) {
        PdbexArgs += "- ";
      } else {
        ShowMessages("err, please insert 'yes' or 'no' as the argument\n\n");
        return FALSE;
      }
      NextItemIsYesNo = FALSE;
      continue;
    }
    if (NextItemIsInline) {
      if (CompareLowerCaseStrings(Item, "none")) {
        PdbexArgs += "n ";
      } else if (CompareLowerCaseStrings(Item, "all")) {
        PdbexArgs += "a ";
      } else if (CompareLowerCaseStrings(Item, "unnamed") ||
                 CompareLowerCaseStrings(Item, "unamed")) {
        PdbexArgs += "i ";
      } else {
        ShowMessages(
            "err, please insert 'none', 'inline', or 'all' as the "
            "argument\n\n");
        return FALSE;
      }
      NextItemIsInline = FALSE;
      continue;
    }
    if (NextItemIsString) {
      PdbexArgs += GetCaseSensitiveStringFromCommandToken(Item) + " ";
      NextItemIsString = FALSE;
      continue;
    }
    if (CompareLowerCaseStrings(Item, "pid")) {
      NextItemIsProcessId = TRUE;
    } else if (CompareLowerCaseStrings(Item, "output")) {
      NextItemIsFileName = TRUE;
      PdbexArgs += "-o ";
    } else if (CompareLowerCaseStrings(Item, "inline")) {
      NextItemIsInline = TRUE;
      PdbexArgs += "-e ";
    } else if (CompareLowerCaseStrings(Item, "prefix")) {
      NextItemIsString = TRUE;
      PdbexArgs += "-r ";
    } else if (CompareLowerCaseStrings(Item, "suffix")) {
      NextItemIsString = TRUE;
      PdbexArgs += "-g ";
    } else if (CompareLowerCaseStrings(Item, "padding")) {
      NextItemIsYesNo = TRUE;
      PdbexArgs += "-p";
    } else if (CompareLowerCaseStrings(Item, "offset") ||
               CompareLowerCaseStrings(Item, "offsets")) {
      NextItemIsYesNo = TRUE;
      PdbexArgs += "-x";
    } else if (CompareLowerCaseStrings(Item, "bitfield") ||
               CompareLowerCaseStrings(Item, "bitfields")) {
      NextItemIsYesNo = TRUE;
      PdbexArgs += "-b";
    } else if (CompareLowerCaseStrings(Item, "native")) {
      NextItemIsYesNo = TRUE;
      PdbexArgs += "-i";
    } else if (CompareLowerCaseStrings(Item, "decl")) {
      NextItemIsYesNo = TRUE;
      PdbexArgs += "-n";
    } else if (CompareLowerCaseStrings(Item, "def")) {
      NextItemIsYesNo = TRUE;
      PdbexArgs += "-l";
    } else if (CompareLowerCaseStrings(Item, "func")) {
      NextItemIsYesNo = TRUE;
      PdbexArgs += "-f";
    } else if (CompareLowerCaseStrings(Item, "pragma")) {
      NextItemIsYesNo = TRUE;
      PdbexArgs += "-z";
    } else {
      ShowMessages("err, unknown argument at '%s'\n\n",
                   GetCaseSensitiveStringFromCommandToken(Item).c_str());
      return FALSE;
    }
  }
  if (NextItemIsYesNo || NextItemIsString || NextItemIsInline ||
      NextItemIsFileName || NextItemIsProcessId) {
    ShowMessages("err, incomplete argument\n\n");
    return FALSE;
  }
  *ProcessId = TargetProcessId;
  return TRUE;
}

BOOLEAN CommandDtShowDataBasedOnSymbolTypes(const char *TypeName,
                                            UINT64 Address, BOOLEAN IsStruct,
                                            PVOID BufferAddress,
                                            UINT32 TargetPid,
                                            BOOLEAN IsPhysicalAddress,
                                            const char *AdditionalParameters) {
  UINT64 StructureSize = 0;
  BOOLEAN ResultOfFindingSize = FALSE;
  DEBUGGER_DT_COMMAND_OPTIONS DtOptions = {0};
  if (g_IsSerialConnectedToRemoteDebuggee && TargetPid != 0) {
    ShowMessages(ASSERT_MESSAGE_CANNOT_SPECIFY_PID);
    return FALSE;
  } else if (TargetPid == NULL) {
    if (g_ActiveProcessDebuggingState.IsActive) {
      TargetPid = g_ActiveProcessDebuggingState.ProcessId;
    } else {
      TargetPid = GetCurrentProcessId();
    }
  }
  DtOptions.TypeName = TypeName;
  DtOptions.Address = Address;
  DtOptions.IsStruct = IsStruct;
  DtOptions.BufferAddress = NULL;
  DtOptions.TargetPid = TargetPid;
  DtOptions.AdditionalParameters = AdditionalParameters;
  if (Address != NULL) {
    ResultOfFindingSize =
        ScriptEngineGetDataTypeSizeWrapper((char *)TypeName, &StructureSize);
    if (!ResultOfFindingSize || StructureSize == 0) {
      ShowMessages("err, couldn't resolve error at '%s'\n", TypeName);
      return FALSE;
    }
    DtOptions.SizeOfTypeName = StructureSize;
    HyperDbgShowMemoryOrDisassemble(
        DEBUGGER_SHOW_COMMAND_DT, Address,
        IsPhysicalAddress ? DEBUGGER_READ_PHYSICAL_ADDRESS
                          : DEBUGGER_READ_VIRTUAL_ADDRESS,
        READ_FROM_KERNEL, TargetPid, (UINT32)StructureSize, &DtOptions);
    return TRUE;
  } else {
    return ScriptEngineShowDataBasedOnSymbolTypesWrapper(
        TypeName, Address, IsStruct, BufferAddress, AdditionalParameters);
  }
}

VOID CommandDtAndStruct(vector<CommandToken> CommandTokens, string Command) {
  CommandToken TempTypeNameHolder;
  std::string PdbexArgs = "";
  BOOLEAN IsStruct = FALSE;
  UINT64 TargetAddress = NULL;
  PVOID BufferAddressRetrievedFromDebuggee = NULL;
  UINT32 TargetPid = NULL;
  BOOLEAN IsPhysicalAddress = FALSE;
  if (CompareLowerCaseStrings(CommandTokens.at(0), "struct") ||
      CompareLowerCaseStrings(CommandTokens.at(0), "structure")) {
    IsStruct = TRUE;
  } else {
    IsStruct = FALSE;
  }
  if (!IsStruct && CompareLowerCaseStrings(CommandTokens.at(0), "!dt")) {
    IsPhysicalAddress = TRUE;
  } else {
    IsPhysicalAddress = FALSE;
  }
  if (CommandTokens.size() == 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    if (IsStruct) {
      CommandStructHelp();
    } else {
      CommandDtHelp();
    }
    return;
  }
  std::vector<CommandToken> TempSplitTokens(CommandTokens.begin() + 1,
                                            CommandTokens.end());
  if (TempSplitTokens.size() == 1) {
    CommandDtShowDataBasedOnSymbolTypes(
        GetCaseSensitiveStringFromCommandToken(TempSplitTokens.at(0)).c_str(),
        NULL, IsStruct, NULL, TargetPid, IsPhysicalAddress,
        PDBEX_DEFAULT_CONFIGURATION);
  } else {
    if (IsStruct ||
        !SymbolConvertNameOrExprToAddress(
            GetCaseSensitiveStringFromCommandToken(TempSplitTokens.at(0))
                .c_str(),
            &TargetAddress)) {
      if (IsStruct ||
          !SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(TempSplitTokens.at(1))
                  .c_str(),
              &TargetAddress)) {
        TempTypeNameHolder = TempSplitTokens.at(0);
        TempSplitTokens.erase(TempSplitTokens.begin());
        if (!CommandDtAndStructConvertHyperDbgArgsToPdbex(
                TempSplitTokens, PdbexArgs, &TargetPid)) {
          if (IsStruct) {
            CommandStructHelp();
          } else {
            CommandDtHelp();
          }
          return;
        }
        CommandDtShowDataBasedOnSymbolTypes(
            GetCaseSensitiveStringFromCommandToken(TempTypeNameHolder).c_str(),
            TargetAddress, IsStruct, BufferAddressRetrievedFromDebuggee,
            TargetPid, IsPhysicalAddress, PdbexArgs.c_str());
      } else {
        if (TempSplitTokens.size() == 2) {
          CommandDtShowDataBasedOnSymbolTypes(
              GetCaseSensitiveStringFromCommandToken(TempSplitTokens.at(0))
                  .c_str(),
              TargetAddress, IsStruct, BufferAddressRetrievedFromDebuggee,
              TargetPid, IsPhysicalAddress, PDBEX_DEFAULT_CONFIGURATION);
        } else {
          TempTypeNameHolder = TempSplitTokens.at(0);
          TempSplitTokens.erase(TempSplitTokens.begin());
          TempSplitTokens.erase(TempSplitTokens.begin());
          if (!CommandDtAndStructConvertHyperDbgArgsToPdbex(
                  TempSplitTokens, PdbexArgs, &TargetPid)) {
            if (IsStruct) {
              CommandStructHelp();
            } else {
              CommandDtHelp();
            }
            return;
          }
          CommandDtShowDataBasedOnSymbolTypes(
              GetCaseSensitiveStringFromCommandToken(TempTypeNameHolder)
                  .c_str(),
              TargetAddress, IsStruct, BufferAddressRetrievedFromDebuggee,
              TargetPid, IsPhysicalAddress, PdbexArgs.c_str());
        }
      }
    } else {
      if (TempSplitTokens.size() == 2) {
        CommandDtShowDataBasedOnSymbolTypes(
            GetCaseSensitiveStringFromCommandToken(TempSplitTokens.at(1))
                .c_str(),
            TargetAddress, IsStruct, BufferAddressRetrievedFromDebuggee,
            TargetPid, IsPhysicalAddress, PDBEX_DEFAULT_CONFIGURATION);
      } else {
        TempTypeNameHolder = TempSplitTokens.at(1);
        TempSplitTokens.erase(TempSplitTokens.begin());
        TempSplitTokens.erase(TempSplitTokens.begin());
        if (!CommandDtAndStructConvertHyperDbgArgsToPdbex(
                TempSplitTokens, PdbexArgs, &TargetPid)) {
          if (IsStruct) {
            CommandStructHelp();
          } else {
            CommandDtHelp();
          }
          return;
        }
        CommandDtShowDataBasedOnSymbolTypes(
            GetCaseSensitiveStringFromCommandToken(TempTypeNameHolder).c_str(),
            TargetAddress, IsStruct, BufferAddressRetrievedFromDebuggee,
            TargetPid, IsPhysicalAddress, PdbexArgs.c_str());
      }
    }
  }
}
