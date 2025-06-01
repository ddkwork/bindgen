
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

VOID CommandReadMemoryAndDisassemblerHelp() {
  ShowMessages(
      "db dc dd dq !db !dc !dd !dq & u u64 !u !u64 u2 u32 !u2 !u32 : reads the "
      " "
      "memory different shapes (hex) and disassembler\n");
  ShowMessages("db  Byte and ASCII characters\n");
  ShowMessages("dc  Double-word values (4 bytes) and ASCII characters\n");
  ShowMessages("dd  Double-word values (4 bytes)\n");
  ShowMessages("dq  Quad-word values (8 bytes). \n");
  ShowMessages("u u64 Disassembler at the target address (x64) \n");
  ShowMessages("u2 u32  Disassembler at the target address (x86) \n");
  ShowMessages(
      "\nIf you want to read physical memory then add '!' at the "
      "start of the command\n");
  ShowMessages("you can also disassemble physical memory using '!u'\n\n");
  ShowMessages(
      "syntax : \tdb [Address (hex)] [l Length (hex)] [pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \tdc [Address (hex)] [l Length (hex)] [pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \tdd [Address (hex)] [l Length (hex)] [pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \tdq [Address (hex)] [l Length (hex)] [pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \tu [Address (hex)] [l Length (hex)] [pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \tu64 [Address (hex)] [l Length (hex)] [pid ProcessId "
      "(hex)]\n");
  ShowMessages(
      "syntax : \tu2 [Address (hex)] [l Length (hex)] [pid ProcessId (hex)]\n");
  ShowMessages(
      "syntax : \tu32 [Address (hex)] [l Length (hex)] [pid ProcessId "
      "(hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : db nt!Kd_DEFAULT_Mask\n");
  ShowMessages("\t\te.g : db nt!Kd_DEFAULT_Mask+10\n");
  ShowMessages("\t\te.g : db @rax\n");
  ShowMessages("\t\te.g : db @rax+50\n");
  ShowMessages("\t\te.g : db fffff8077356f010\n");
  ShowMessages("\t\te.g : !dq 100000\n");
  ShowMessages("\t\te.g : !dq @rax+77\n");
  ShowMessages("\t\te.g : u32 @eip\n");
  ShowMessages("\t\te.g : u nt!ExAllocatePoolWithTag\n");
  ShowMessages("\t\te.g : u nt!ExAllocatePoolWithTag+30\n");
  ShowMessages("\t\te.g : u fffff8077356f010\n");
  ShowMessages("\t\te.g : u fffff8077356f010+@rcx\n");
}

VOID CommandReadMemoryAndDisassembler(vector<CommandToken> CommandTokens,
                                      string Command) {
  UINT32 Pid = 0;
  UINT32 Length = 0;
  UINT64 TargetAddress = 0;
  BOOLEAN IsNextProcessId = FALSE;
  BOOLEAN IsFirstCommand = TRUE;
  BOOLEAN IsNextLength = FALSE;
  string FirstCommand =
      GetCaseSensitiveStringFromCommandToken(CommandTokens.front());
  if (g_ActiveProcessDebuggingState.IsActive) {
    Pid = g_ActiveProcessDebuggingState.ProcessId;
  }
  if (CommandTokens.size() == 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandReadMemoryAndDisassemblerHelp();
    return;
  }
  for (auto Section : CommandTokens) {
    if (IsFirstCommand) {
      IsFirstCommand = FALSE;
      continue;
    }
    if (IsNextProcessId == TRUE) {
      if (!ConvertTokenToUInt32(Section, &Pid)) {
        ShowMessages("err, you should enter a valid process id\n\n");
        return;
      }
      IsNextProcessId = FALSE;
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
    if (CompareLowerCaseStrings(Section, "pid")) {
      IsNextProcessId = TRUE;
      continue;
    }
    if (TargetAddress == 0) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(Section),
              &TargetAddress)) {
        ShowMessages("err, couldn't resolve error at '%s'\n",
                     GetCaseSensitiveStringFromCommandToken(Section).c_str());
        return;
      }
    } else {
      ShowMessages(
          "err, incorrect use of the '%s' command\n\n",
          GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
      CommandReadMemoryAndDisassemblerHelp();
      return;
    }
  }
  if (!TargetAddress) {
    ShowMessages("err, please enter a valid address\n\n");
    return;
  }
  if (Length == 0) {
    if (CompareLowerCaseStrings(CommandTokens.at(0), "u") ||
        CompareLowerCaseStrings(CommandTokens.at(0), "!u") ||
        CompareLowerCaseStrings(CommandTokens.at(0), "u64") ||
        CompareLowerCaseStrings(CommandTokens.at(0), "!u64")) {
      Length = 0x40;
    } else {
      Length = 0x80;
    }
  }
  if (IsNextLength || IsNextProcessId) {
    ShowMessages(
        "incorrect use of the '%s' command\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandReadMemoryAndDisassemblerHelp();
    return;
  }
  if (g_IsSerialConnectedToRemoteDebuggee && Pid != 0) {
    ShowMessages(ASSERT_MESSAGE_CANNOT_SPECIFY_PID);
    return;
  }
  if (Pid == 0) {
    Pid = GetCurrentProcessId();
  }
  if (CompareLowerCaseStrings(CommandTokens.at(0), "db")) {
    HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_COMMAND_DB, TargetAddress,
                                    DEBUGGER_READ_VIRTUAL_ADDRESS,
                                    READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "dc")) {
    HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_COMMAND_DC, TargetAddress,
                                    DEBUGGER_READ_VIRTUAL_ADDRESS,
                                    READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "dd")) {
    HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_COMMAND_DD, TargetAddress,
                                    DEBUGGER_READ_VIRTUAL_ADDRESS,
                                    READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "dq")) {
    HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_COMMAND_DQ, TargetAddress,
                                    DEBUGGER_READ_VIRTUAL_ADDRESS,
                                    READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "!db")) {
    HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_COMMAND_DB, TargetAddress,
                                    DEBUGGER_READ_PHYSICAL_ADDRESS,
                                    READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "!dc")) {
    HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_COMMAND_DC, TargetAddress,
                                    DEBUGGER_READ_PHYSICAL_ADDRESS,
                                    READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "!dd")) {
    HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_COMMAND_DD, TargetAddress,
                                    DEBUGGER_READ_PHYSICAL_ADDRESS,
                                    READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "!dq")) {
    HyperDbgShowMemoryOrDisassemble(DEBUGGER_SHOW_COMMAND_DQ, TargetAddress,
                                    DEBUGGER_READ_PHYSICAL_ADDRESS,
                                    READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "u") ||
             CompareLowerCaseStrings(CommandTokens.at(0), "u64")) {
    HyperDbgShowMemoryOrDisassemble(
        DEBUGGER_SHOW_COMMAND_DISASSEMBLE64, TargetAddress,
        DEBUGGER_READ_VIRTUAL_ADDRESS, READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "!u") ||
             CompareLowerCaseStrings(CommandTokens.at(0), "!u64")) {
    HyperDbgShowMemoryOrDisassemble(
        DEBUGGER_SHOW_COMMAND_DISASSEMBLE64, TargetAddress,
        DEBUGGER_READ_PHYSICAL_ADDRESS, READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "u2") ||
             CompareLowerCaseStrings(CommandTokens.at(0), "u32")) {
    HyperDbgShowMemoryOrDisassemble(
        DEBUGGER_SHOW_COMMAND_DISASSEMBLE32, TargetAddress,
        DEBUGGER_READ_VIRTUAL_ADDRESS, READ_FROM_KERNEL, Pid, Length, NULL);
  } else if (CompareLowerCaseStrings(CommandTokens.at(0), "!u2") ||
             CompareLowerCaseStrings(CommandTokens.at(0), "!u32")) {
    HyperDbgShowMemoryOrDisassemble(
        DEBUGGER_SHOW_COMMAND_DISASSEMBLE32, TargetAddress,
        DEBUGGER_READ_PHYSICAL_ADDRESS, READ_FROM_KERNEL, Pid, Length, NULL);
  }
}
