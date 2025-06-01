
#include "pch.h"
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
extern CommandType g_CommandsList;
extern BOOLEAN g_ShouldPreviousCommandBeContinued;
extern BOOLEAN g_IsCommandListInitialized;
extern BOOLEAN g_LogOpened;
extern BOOLEAN g_ExecutingScript;
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsDebuggeeRunning;
extern BOOLEAN g_BreakPrintingOutput;
extern BOOLEAN g_IsInterpreterOnString;
extern BOOLEAN g_IsInterpreterPreviousCharacterABackSlash;
extern BOOLEAN g_RtmSupport;
extern UINT32 g_VirtualAddressWidth;
extern UINT32 g_InterpreterCountOfOpenCurlyBrackets;
extern ULONG g_CurrentRemoteCore;
extern string g_ServerPort;
extern string g_ServerIp;

class CommandParser {
 public:
  std::vector<CommandToken> Parse(const std::string &ConstInput) {
    std::vector<CommandToken> tokens;
    std::string current;
    bool InQuotes = FALSE;
    int IdxBracket = 0;
    std::string input = ConstInput;
    for (size_t i = 0; i < input.length(); ++i) {
      char c = input[i];
      if (c == '/') {
        if (!InQuotes) {
          size_t j = i;
          char c2 = input[++j];
          if (c2 == '/') {
            size_t StrLitEnd = 0;
            size_t StrLitBeg = input.find("\"", i);
            if (StrLitBeg != std::string::npos) {
              if (i) {
                if (input[i - 1] != '\\')  // if not escaped
                {
                  StrLitEnd = input.find("\"", StrLitBeg + 1);
                  if (StrLitEnd != std::string::npos) {
                    std::string StrLit(input.substr(i, StrLitEnd - i + 1));
                  }
                }
              } else {
                StrLitEnd = input.find("\"", StrLitBeg + 1);
                if (StrLitEnd != std::string::npos) {
                  std::string StrLit(input.substr(i, StrLitEnd - i + 1));
                }
              }
            }
            bool CmntEndBrkt = false;
            size_t CloseBrktPos = 0;
            if (IdxBracket) {
              auto pos = (StrLitEnd > i) ? StrLitEnd : i;
              for (CloseBrktPos = input.find("}", pos);
                   CloseBrktPos != std::string::npos;) {
                CloseBrktPos = input.find("}", CloseBrktPos);
                if (input[CloseBrktPos - 1] == '\\') {
                  input.erase(CloseBrktPos - 1, 1);
                  CloseBrktPos += 1;
                } else {
                  break;
                }
              }
            } else {
              CloseBrktPos = std::string::npos;
            }
            size_t NewLineSrtPos =
                input.find("\\n", i);  // "\\n" entered by user
            if (StrLitBeg && StrLitBeg <= NewLineSrtPos &&
                NewLineSrtPos <= StrLitEnd)  // is it within the string literal?
            {
              NewLineSrtPos = std::string::npos;
            }
            size_t NewLineChrPos = input.find('\n', i);
            std::vector<size_t> PosVec = {CloseBrktPos, NewLineSrtPos,
                                          NewLineChrPos};
            auto min = *(min_element(
                PosVec.begin(), PosVec.end()));  // see which one occures first
                                                 // if (min != std::string::npos
                                                 // && input[min - 1] != '\\') {
            std::string comment(input.substr(i, min - i));
            if (IdxBracket) {
              current += comment;
            }
            auto diff = (min > i) ? (min - i) : (i - min);
            i = i + diff - 1;
            continue;
          } else {
            bool IsNewLineEsc = false;
            if (NewLineSrtPos != std::string::npos) {
              IsNewLineEsc = input[NewLineSrtPos - 1] == '\\';
            }
            std::string comment(input.substr(i, input.size()));
            if (IsNewLineEsc) {
              size_t start_pos = 0;
              while ((start_pos = comment.find("\\\\n", start_pos)) !=
                     std::string::npos) {
                comment.replace(start_pos, 3, "\\n");
                start_pos +=
                    2;  // Handles case where 'to' is a substring of 'from'
              }
              IsNewLineEsc = false;
            }
            if (IdxBracket) {
              current += comment;
            }
            i = i + (input.size() - i);
            continue;
          }
        } else if (c2 == '*') {
          size_t EndPose = input.find("*/", i + 2);  // +2 for cases like /*/
          if (EndPose != std::string::npos) {
            std::string comment(input.substr(
                i, EndPose - i + 2));  // */ is two bytes long if (IdxBracket) {
            current += comment;
          }
          i = (i + (EndPose - i)) + 1;  // +1 for /
          continue;
        } else {
        }
      }
    }
  }
  if (InQuotes) {
    if (c == '"') {
      if (input[i - 1] != '\\') {
        InQuotes = FALSE;
        if (!IdxBracket) {
          AddStringToken(tokens, current,
                         TRUE);  // TRUE for StringLiteral type
          current.clear();
          continue;  // dont add " char
        } else {
          current += c;
          continue;  // dont add " char
        }
      } else {
        input.erase(i - 1, 1);
        i--;  // compensate for the removed char if (!IdxBracket) {
        current.pop_back();
      }
      current += c;
      continue;
    }
  }
} if (c == '}') {
  if (input[i - 1] != '\\') {
    if (IdxBracket) {
      if (!InQuotes)  // not closing "
      {
        IdxBracket--;
      }
      if (!IdxBracket)  // is closing }
      {
        AddBracketStringToken(tokens, current);
        current.clear();
        continue;
      }
    }
  } else if (!InQuotes) {
    input.erase(i - 1, 1);
    i--;                 // compensate for the removed char
    current.pop_back();  // remove last read \\

  }
}
if (((c == ' ' && !InQuotes) || c == '    ') && !InQuotes &&
    !IdxBracket)  // finding seperator space char // Tab seperator added
                  // too
{
  if (!current.empty() && current != " ") {
    AddToken(tokens, current);
    current.clear();
    continue;
  }
  continue;           // avoid adding extra space char
} else if (c == '"')  // string literal is adjacent to previous command
{
  if (i)  // check if this " is the first char to avoid out of range check
  {
    if (input[i - 1] != ' ' && !IdxBracket && !current.empty() && !InQuotes) {
      AddStringToken(tokens, current);
      current.clear();
    }
    if (input[i - 1] != '\\') {
      InQuotes = TRUE;
      if (!IdxBracket) {
        continue;
      }
    }
  } else {
    InQuotes = TRUE;
    if (!IdxBracket) {
      continue;
    }
  }
} else if (c == '{' && !InQuotes) {
  if (i) {
    if (input[i - 1] != '\\') {
      if (input[i - 1] != ' ' && !IdxBracket) {
        AddToken(tokens, current);
        current.clear();
      }
      IdxBracket++;
      if (IdxBracket == 1) continue;
    } else {
      input.erase(i - 1, 1);
      i--;
      current.pop_back();
    }
  } else {
    IdxBracket++;
    if (IdxBracket == 1) continue;
  }
}
if (c == '\\' && !InQuotes) {
  if (current.empty() && input[i + 1] == 'n') {
    i++;
    continue;
  }
}
current += c;
}
if (!current.empty() && current != " ") {
  AddToken(tokens, current);
}
if (IdxBracket) {
}
if (InQuotes) {
}
return tokens;
}

std::string TokenTypeToString(CommandParsingTokenType Type) {
  switch (Type) {
    case CommandParsingTokenType::Num:
      return "Num";
    case CommandParsingTokenType::String:
      return "String";
    case CommandParsingTokenType::StringLiteral:
      return "StringLiteral";
    case CommandParsingTokenType::BracketString:
      return "BracketString";
    default:
      return "Unknown";
  }
}

VOID PrintTokens(const std::vector<CommandToken> &Tokens) {
  const int sz = 200;
  int g_s1Len = 0, g_s2Len = 0, g_s3Len = 0;
  int s1 = 0, s2 = 0, s3 = 0;
  char LineToPrint1[sz], LineToPrint2[sz], LineToPrint3[sz];
  for (const auto &Token : Tokens) {
    s1 = snprintf(LineToPrint1, sz, "CommandParsingTokenType: %s ",
                  TokenTypeToString(std::get<0>(Token)).c_str());
    s2 = snprintf(LineToPrint2, sz, ", Value 1: '%s'",
                  std::get<1>(Token).c_str());
    s3 = snprintf(LineToPrint3, sz, ", Value 2 (lower): '%s'",
                  std::get<2>(Token).c_str());
    if (s1 > g_s1Len) g_s1Len = s1;
    if (s2 > g_s2Len) g_s2Len = s2;
    if (s3 > g_s3Len) g_s3Len = s3;
  }
  for (const auto &Token : Tokens) {
    auto CaseSensitiveText = std::get<1>(Token);
    auto LowerCaseText = std::get<2>(Token);
    if (std::get<0>(Token) == CommandParsingTokenType::BracketString ||
        std::get<0>(Token) == CommandParsingTokenType::String ||
        std::get<0>(Token) == CommandParsingTokenType::StringLiteral) {
      std::string::size_type pos = 0;
      while ((pos = CaseSensitiveText.find("\n", pos)) != std::string::npos) {
        CaseSensitiveText.replace(pos, 1, "\\n");
        pos += 2;
      }
      pos = 0;
      while ((pos = LowerCaseText.find("\n", pos)) != std::string::npos) {
        LowerCaseText.replace(pos, 1, "\\n");
        pos += 2;
      }
    }
    snprintf(LineToPrint1, sz, "CommandParsingTokenType: %s ",
             TokenTypeToString(std::get<0>(Token)).c_str());
    snprintf(LineToPrint2, sz, ", Value 1: '%s'", CaseSensitiveText.c_str());
    snprintf(LineToPrint3, sz, ", Value 2 (lower): '%s'",
             LowerCaseText.c_str());
    ShowMessages("%-*s %-*s %-*s\n", g_s1Len, LineToPrint1, g_s2Len,
                 LineToPrint2, g_s3Len, LineToPrint3);
  }
}

private:
std::string ToLower(const std::string &str) const {
  std::string result = str;
  std::transform(result.begin(), result.end(), result.begin(), ::tolower);
  return result;
}

VOID AddToken(std::vector<CommandToken> &tokens, std::string &str) {
  auto tmp = str;
  UINT64 tmpNum = 0;
  Trim(tmp);
  if (ConvertStringToUInt64(tmp, &tmpNum)) {
    tokens.emplace_back(CommandParsingTokenType::Num, tmp, ToLower(tmp));
  } else {
    AddStringToken(tokens, str);
  }
}

VOID AddStringToken(std::vector<CommandToken> &tokens, std::string &str,
                    BOOL isLiteral = FALSE) {
  auto tmp = str;
  if (!isLiteral) Trim(tmp);
  if (tmp.empty()) return;
  if (isLiteral) {
    tokens.emplace_back(CommandParsingTokenType::StringLiteral, tmp,
                        ToLower(tmp));
  } else {
    tokens.emplace_back(CommandParsingTokenType::String, tmp, ToLower(tmp));
  }
}

VOID AddBracketStringToken(std::vector<CommandToken> &tokens,
                           const std::string &str) {
  tokens.emplace_back(CommandParsingTokenType::BracketString, str,
                      ToLower(str));
}
}
;

int FindDifferencePosition(const char *prsTok, const char *fileTok) {
  int i = 0;
  while (prsTok[i] != '\0' && fileTok[i] != '\0') {
    if (prsTok[i] != fileTok[i]) {
      return i;
    }
    i++;
  }
  if (prsTok[i] != fileTok[i]) {
    return i;
  }
  return -1;
}

BOOLEAN HyperDbgTestCommandParser(CHAR *Command, UINT32 NumberOfTokens,
                                  CHAR **TokensList, UINT32 *FailedTokenNum,
                                  UINT32 *FailedTokenPosition) {
  CommandParser Parser;
  string CommandString(Command);
  std::vector<CommandToken> Tokens = Parser.Parse(CommandString);
  if (Tokens.size() != NumberOfTokens) {
    ShowMessages("err, the number of tokens is not correct\n");
    return FALSE;
  }
  for (UINT32 i = 0; i < Tokens.size(); i++) {
    auto Token = Tokens.at(i);
    auto CaseSensitiveText = std::get<1>(Token);
    if (strcmp(CaseSensitiveText.c_str(), TokensList[i]) != 0) {
      *FailedTokenNum = i;
      *FailedTokenPosition =
          FindDifferencePosition(CaseSensitiveText.c_str(), TokensList[i]);
      ShowMessages("err, the token is not correct\n");
      return FALSE;
    }
  }
  return TRUE;
}

VOID HyperDbgTestCommandParserShowTokens(CHAR *Command) {
  CommandParser Parser;
  string CommandString(Command);
  std::vector<CommandToken> Tokens = Parser.Parse(CommandString);
  Parser.PrintTokens(Tokens);
}

INT HyperDbgInterpreter(CHAR *Command) {
  BOOLEAN HelpCommand = FALSE;
  UINT64 CommandAttributes = NULL;
  CommandType::iterator Iterator;
  CommandParser Parser;
  if (!g_IsCommandListInitialized) {
    InitializeDebugger();
    g_IsCommandListInitialized = TRUE;
  }
  if (g_LogOpened && !g_ExecutingScript) {
    LogopenSaveToFile(Command);
    LogopenSaveToFile("\n");
  }
  string CommandString(Command);
  auto Tokens = Parser.Parse(CommandString);
  if (Tokens.empty()) {
    ShowMessages("\n");
    return 0;
  }
  string FirstCommand = GetLowerStringFromCommandToken(Tokens.front());
  CommandAttributes = GetCommandAttributes(FirstCommand);
  if (CommandAttributes & DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER) {
    g_ShouldPreviousCommandBeContinued = TRUE;
  } else {
    g_ShouldPreviousCommandBeContinued = FALSE;
  }
  if (g_IsConnectedToRemoteDebuggee &&
      !(CommandAttributes &
        DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_REMOTE_CONNECTION)) {
    if (g_BreakPrintingOutput) {
      g_BreakPrintingOutput = FALSE;
    }
    RemoteConnectionSendCommand(Command, (UINT32)strlen(Command) + 1);
    ShowMessages("\n");
    return 2;
  } else if (g_IsSerialConnectedToRemoteDebuggee &&
             !(CommandAttributes &
               DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE)) {
    if (CommandAttributes &
        DEBUGGER_COMMAND_ATTRIBUTE_WONT_STOP_DEBUGGER_AGAIN) {
      KdSendUserInputPacketToDebuggee(Command, (UINT32)strlen(Command) + 1,
                                      TRUE);
      KdSetStatusAndWaitForPause();
    } else {
      KdSendTestQueryPacketToDebuggee(
          TEST_BREAKPOINT_TURN_OFF_BPS_AND_EVENTS_FOR_COMMANDS_IN_REMOTE_COMPUTER);
      KdSendUserInputPacketToDebuggee(Command, (UINT32)strlen(Command) + 1,
                                      FALSE);
      KdSendTestQueryPacketToDebuggee(
          TEST_BREAKPOINT_TURN_ON_BPS_AND_EVENTS_FOR_COMMANDS_IN_REMOTE_COMPUTER);
    }
    return 2;
  }
  if (!FirstCommand.compare(".help") || !FirstCommand.compare("help") ||
      !FirstCommand.compare(".hh")) {
    if (Tokens.size() == 2) {
      HelpCommand = TRUE;
      FirstCommand = GetLowerStringFromCommandToken(Tokens.at(1));
    } else {
      ShowMessages(
          "incorrect use of the '%s'\n\n",
          GetCaseSensitiveStringFromCommandToken(Tokens.at(0)).c_str());
      CommandHelpHelp();
      return 0;
    }
  }
  string CaseSensitiveCommandString(Command);
  Iterator = g_CommandsList.find(FirstCommand);
  if (Iterator == g_CommandsList.end()) {
    if (!HelpCommand) {
      ShowMessages(
          "err, couldn't resolve command at '%s'\n",
          GetCaseSensitiveStringFromCommandToken(Tokens.front()).c_str());
    } else {
      ShowMessages(
          "err, couldn't find the help for the command at '%s'\n",
          GetCaseSensitiveStringFromCommandToken(Tokens.at(1)).c_str());
    }
  } else {
    if (HelpCommand) {
      Iterator->second.CommandHelpFunction();
    } else {
      Iterator->second.CommandFunctionNewParser(Tokens,
                                                CaseSensitiveCommandString);
    }
  }
  if (g_LogOpened && !g_ExecutingScript) {
    LogopenSaveToFile("\n");
  }
  return 0;
}

VOID InterpreterRemoveComments(char *CommandText) {
  BOOLEAN IsComment = FALSE;
  BOOLEAN IsOnBracketString = FALSE;
  BOOLEAN IsOnString = FALSE;
  UINT32 LengthOfCommand = (UINT32)strlen(CommandText);
  for (size_t i = 0; i < LengthOfCommand; i++) {
    if (IsComment) {
      if (CommandText[i] == '\n') {
        IsComment = FALSE;
      } else {
        if (CommandText[i] != '\0') {
          memmove((void *)&CommandText[i], (const void *)&CommandText[i + 1],
                  strlen(CommandText) - i);
          i--;
        }
      }
    } else if (CommandText[i] == '#' && !IsOnString) {
      IsComment = TRUE;
      i--;
    } else if (CommandText[i] == '#' && !IsOnString) {
      IsComment = TRUE;
      i--;
    } else if (CommandText[i] == '"') {
      if (i != 0 && CommandText[i - 1] == '\\') {
      } else if (IsOnString) {
        IsOnString = FALSE;
      } else {
        IsOnString = TRUE;
      }
    }
  }
}

VOID HyperDbgShowSignature() {
  if (g_IsConnectedToRemoteDebuggee) {
    ShowMessages("[%s:%s] HyperDbg> ", g_ServerIp.c_str(),
                 g_ServerPort.c_str());
  } else if (g_ActiveProcessDebuggingState.IsActive) {
    ShowMessages("%x:%x u%sHyperDbg> ", g_ActiveProcessDebuggingState.ProcessId,
                 g_ActiveProcessDebuggingState.ThreadId,
                 g_ActiveProcessDebuggingState.Is32Bit ? "86" : "64");
  } else if (g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages("%x: kHyperDbg> ", g_CurrentRemoteCore);
  } else {
    ShowMessages("HyperDbg> ");
  }
}

BOOLEAN CheckMultilineCommand(CHAR *CurrentCommand, BOOLEAN Reset) {
  UINT32 CurrentCommandLen = 0;
  std::string CurrentCommandStr(CurrentCommand);
  if (Reset) {
    g_IsInterpreterOnString = FALSE;
    g_IsInterpreterPreviousCharacterABackSlash = FALSE;
    g_InterpreterCountOfOpenCurlyBrackets = 0;
  }
  CurrentCommandLen = (UINT32)CurrentCommandStr.length();
  for (size_t i = 0; i < CurrentCommandLen; i++) {
    switch (CurrentCommandStr.at(i)) {
      case '"':
        if (g_IsInterpreterPreviousCharacterABackSlash) {
          g_IsInterpreterPreviousCharacterABackSlash = FALSE;
          break;
        }
        if (g_IsInterpreterOnString)
          g_IsInterpreterOnString = FALSE;
        else
          g_IsInterpreterOnString = TRUE;
        break;
      case '{':
        if (g_IsInterpreterPreviousCharacterABackSlash)
          g_IsInterpreterPreviousCharacterABackSlash = FALSE;
        if (!g_IsInterpreterOnString) g_InterpreterCountOfOpenCurlyBrackets++;
        break;
      case '}':
        if (g_IsInterpreterPreviousCharacterABackSlash)
          g_IsInterpreterPreviousCharacterABackSlash = FALSE;
        if (!g_IsInterpreterOnString &&
            g_InterpreterCountOfOpenCurlyBrackets > 0)
          g_InterpreterCountOfOpenCurlyBrackets--;
        break;
      case '\\':
        if (g_IsInterpreterPreviousCharacterABackSlash)
          g_IsInterpreterPreviousCharacterABackSlash = FALSE;
        else
          g_IsInterpreterPreviousCharacterABackSlash = TRUE;
        break;
      default:
        if (g_IsInterpreterPreviousCharacterABackSlash)
          g_IsInterpreterPreviousCharacterABackSlash = FALSE;
        break;
    }
  }
  if (g_IsInterpreterOnString == FALSE &&
      g_InterpreterCountOfOpenCurlyBrackets == 0) {
    return FALSE;
  } else {
    return TRUE;
  }
}

BOOLEAN ContinuePreviousCommand() {
  BOOLEAN Result = g_ShouldPreviousCommandBeContinued;
  g_ShouldPreviousCommandBeContinued = FALSE;
  if (Result) {
    return TRUE;
  } else {
    return FALSE;
  }
}

UINT64 GetCommandAttributes(const string &FirstCommand) {
  CommandType::iterator Iterator;
  Iterator = g_CommandsList.find(FirstCommand);
  if (Iterator == g_CommandsList.end()) {
    return DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL;
  } else {
    return Iterator->second.CommandAttrib;
  }
  return NULL;
}

VOID InitializeDebugger() {
  InitializeCommandsDictionary();
  ScriptEngineSetTextMessageCallbackWrapper(ShowMessages);
  if (!SetConsoleCtrlHandler(BreakController, TRUE)) {
    ShowMessages(
        "err, when registering CTRL+C and CTRL+BREAK Signals "
        "handler\n");
  }
  g_VirtualAddressWidth = Getx86VirtualAddressWidth();
  g_RtmSupport = CheckCpuSupportRtm();
  CommandSettingsLoadDefaultValuesFromConfigFile();
}

VOID InitializeCommandsDictionary() {
  g_CommandsList[".help"] = {NULL, &CommandHelpHelp,
                             DEBUGGER_COMMAND_HELP_ATTRIBUTES};
  g_CommandsList[".hh"] = {NULL, &CommandHelpHelp,
                           DEBUGGER_COMMAND_HELP_ATTRIBUTES};
  g_CommandsList["help"] = {NULL, &CommandHelpHelp,
                            DEBUGGER_COMMAND_HELP_ATTRIBUTES};
  g_CommandsList["clear"] = {&CommandCls, &CommandClsHelp,
                             DEBUGGER_COMMAND_CLEAR_ATTRIBUTES};
  g_CommandsList[".cls"] = {&CommandCls, &CommandClsHelp,
                            DEBUGGER_COMMAND_CLEAR_ATTRIBUTES};
  g_CommandsList["cls"] = {&CommandCls, &CommandClsHelp,
                           DEBUGGER_COMMAND_CLEAR_ATTRIBUTES};
  g_CommandsList[".connect"] = {&CommandConnect, &CommandConnectHelp,
                                DEBUGGER_COMMAND_CONNECT_ATTRIBUTES};
  g_CommandsList["connect"] = {&CommandConnect, &CommandConnectHelp,
                               DEBUGGER_COMMAND_CONNECT_ATTRIBUTES};
  g_CommandsList[".listen"] = {&CommandListen, &CommandListenHelp,
                               DEBUGGER_COMMAND_LISTEN_ATTRIBUTES};
  g_CommandsList["listen"] = {&CommandListen, &CommandListenHelp,
                              DEBUGGER_COMMAND_LISTEN_ATTRIBUTES};
  g_CommandsList["g"] = {&CommandG, &CommandGHelp,
                         DEBUGGER_COMMAND_G_ATTRIBUTES};
  g_CommandsList["go"] = {&CommandG, &CommandGHelp,
                          DEBUGGER_COMMAND_G_ATTRIBUTES};
  g_CommandsList["gg"] = {&CommandGg, &CommandGgHelp,
                          DEBUGGER_COMMAND_GG_ATTRIBUTES};
  g_CommandsList[".attach"] = {&CommandAttach, &CommandAttachHelp,
                               DEBUGGER_COMMAND_ATTACH_ATTRIBUTES};
  g_CommandsList["attach"] = {&CommandAttach, &CommandAttachHelp,
                              DEBUGGER_COMMAND_ATTACH_ATTRIBUTES};
  g_CommandsList[".detach"] = {&CommandDetach, &CommandDetachHelp,
                               DEBUGGER_COMMAND_DETACH_ATTRIBUTES};
  g_CommandsList["detach"] = {&CommandDetach, &CommandDetachHelp,
                              DEBUGGER_COMMAND_DETACH_ATTRIBUTES};
  g_CommandsList[".start"] = {&CommandStart, &CommandStartHelp,
                              DEBUGGER_COMMAND_START_ATTRIBUTES};
  g_CommandsList["start"] = {&CommandStart, &CommandStartHelp,
                             DEBUGGER_COMMAND_START_ATTRIBUTES};
  g_CommandsList[".restart"] = {&CommandRestart, &CommandRestartHelp,
                                DEBUGGER_COMMAND_RESTART_ATTRIBUTES};
  g_CommandsList["restart"] = {&CommandRestart, &CommandRestartHelp,
                               DEBUGGER_COMMAND_RESTART_ATTRIBUTES};
  g_CommandsList[".switch"] = {&CommandSwitch, &CommandSwitchHelp,
                               DEBUGGER_COMMAND_SWITCH_ATTRIBUTES};
  g_CommandsList["switch"] = {&CommandSwitch, &CommandSwitchHelp,
                              DEBUGGER_COMMAND_SWITCH_ATTRIBUTES};
  g_CommandsList[".kill"] = {&CommandKill, &CommandKillHelp,
                             DEBUGGER_COMMAND_KILL_ATTRIBUTES};
  g_CommandsList["kill"] = {&CommandKill, &CommandKillHelp,
                            DEBUGGER_COMMAND_KILL_ATTRIBUTES};
  g_CommandsList[".process"] = {&CommandProcess, &CommandProcessHelp,
                                DEBUGGER_COMMAND_PROCESS_ATTRIBUTES};
  g_CommandsList[".process2"] = {&CommandProcess, &CommandProcessHelp,
                                 DEBUGGER_COMMAND_PROCESS_ATTRIBUTES};
  g_CommandsList["process"] = {&CommandProcess, &CommandProcessHelp,
                               DEBUGGER_COMMAND_PROCESS_ATTRIBUTES};
  g_CommandsList["process2"] = {&CommandProcess, &CommandProcessHelp,
                                DEBUGGER_COMMAND_PROCESS_ATTRIBUTES};
  g_CommandsList[".thread"] = {&CommandThread, &CommandThreadHelp,
                               DEBUGGER_COMMAND_THREAD_ATTRIBUTES};
  g_CommandsList[".thread2"] = {&CommandThread, &CommandThreadHelp,
                                DEBUGGER_COMMAND_THREAD_ATTRIBUTES};
  g_CommandsList["thread"] = {&CommandThread, &CommandThreadHelp,
                              DEBUGGER_COMMAND_THREAD_ATTRIBUTES};
  g_CommandsList["thread2"] = {&CommandThread, &CommandThreadHelp,
                               DEBUGGER_COMMAND_THREAD_ATTRIBUTES};
  g_CommandsList["sleep"] = {&CommandSleep, &CommandSleepHelp,
                             DEBUGGER_COMMAND_SLEEP_ATTRIBUTES};
  g_CommandsList["event"] = {&CommandEvents, &CommandEventsHelp,
                             DEBUGGER_COMMAND_EVENTS_ATTRIBUTES};
  g_CommandsList["events"] = {&CommandEvents, &CommandEventsHelp,
                              DEBUGGER_COMMAND_EVENTS_ATTRIBUTES};
  g_CommandsList["setting"] = {&CommandSettings, &CommandSettingsHelp,
                               DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES};
  g_CommandsList["settings"] = {&CommandSettings, &CommandSettingsHelp,
                                DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES};
  g_CommandsList[".setting"] = {&CommandSettings, &CommandSettingsHelp,
                                DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES};
  g_CommandsList[".settings"] = {&CommandSettings, &CommandSettingsHelp,
                                 DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES};
  g_CommandsList[".disconnect"] = {&CommandDisconnect, &CommandDisconnectHelp,
                                   DEBUGGER_COMMAND_DISCONNECT_ATTRIBUTES};
  g_CommandsList["disconnect"] = {&CommandDisconnect, &CommandDisconnectHelp,
                                  DEBUGGER_COMMAND_DISCONNECT_ATTRIBUTES};
  g_CommandsList[".debug"] = {&CommandDebug, &CommandDebugHelp,
                              DEBUGGER_COMMAND_DEBUG_ATTRIBUTES};
  g_CommandsList["debug"] = {&CommandDebug, &CommandDebugHelp,
                             DEBUGGER_COMMAND_DEBUG_ATTRIBUTES};
  g_CommandsList[".status"] = {&CommandStatus, &CommandStatusHelp,
                               DEBUGGER_COMMAND_DOT_STATUS_ATTRIBUTES};
  g_CommandsList["status"] = {&CommandStatus, &CommandStatusHelp,
                              DEBUGGER_COMMAND_STATUS_ATTRIBUTES};
  g_CommandsList["load"] = {&CommandLoad, &CommandLoadHelp,
                            DEBUGGER_COMMAND_LOAD_ATTRIBUTES};
  g_CommandsList[".load"] = {&CommandLoad, &CommandLoadHelp,
                             DEBUGGER_COMMAND_LOAD_ATTRIBUTES};
  g_CommandsList["exit"] = {&CommandExit, &CommandExitHelp,
                            DEBUGGER_COMMAND_EXIT_ATTRIBUTES};
  g_CommandsList[".exit"] = {&CommandExit, &CommandExitHelp,
                             DEBUGGER_COMMAND_EXIT_ATTRIBUTES};
  g_CommandsList["flush"] = {&CommandFlush, &CommandFlushHelp,
                             DEBUGGER_COMMAND_FLUSH_ATTRIBUTES};
  g_CommandsList["pause"] = {&CommandPause, &CommandPauseHelp,
                             DEBUGGER_COMMAND_PAUSE_ATTRIBUTES};
  g_CommandsList[".pause"] = {&CommandPause, &CommandPauseHelp,
                              DEBUGGER_COMMAND_PAUSE_ATTRIBUTES};
  g_CommandsList["unload"] = {&CommandUnload, &CommandUnloadHelp,
                              DEBUGGER_COMMAND_UNLOAD_ATTRIBUTES};
  g_CommandsList[".script"] = {&CommandScript, &CommandScriptHelp,
                               DEBUGGER_COMMAND_SCRIPT_ATTRIBUTES};
  g_CommandsList["script"] = {&CommandScript, &CommandScriptHelp,
                              DEBUGGER_COMMAND_SCRIPT_ATTRIBUTES};
  g_CommandsList["output"] = {&CommandOutput, &CommandOutputHelp,
                              DEBUGGER_COMMAND_OUTPUT_ATTRIBUTES};
  g_CommandsList["print"] = {&CommandPrint, &CommandPrintHelp,
                             DEBUGGER_COMMAND_PRINT_ATTRIBUTES};
  g_CommandsList["?"] = {&CommandEval, &CommandEvalHelp,
                         DEBUGGER_COMMAND_EVAL_ATTRIBUTES};
  g_CommandsList["eval"] = {&CommandEval, &CommandEvalHelp,
                            DEBUGGER_COMMAND_EVAL_ATTRIBUTES};
  g_CommandsList["evaluate"] = {&CommandEval, &CommandEvalHelp,
                                DEBUGGER_COMMAND_EVAL_ATTRIBUTES};
  g_CommandsList[".logopen"] = {&CommandLogopen, &CommandLogopenHelp,
                                DEBUGGER_COMMAND_LOGOPEN_ATTRIBUTES};
  g_CommandsList[".logclose"] = {&CommandLogclose, &CommandLogcloseHelp,
                                 DEBUGGER_COMMAND_LOGCLOSE_ATTRIBUTES};
  g_CommandsList[".pagein"] = {&CommandPagein, &CommandPageinHelp,
                               DEBUGGER_COMMAND_PAGEIN_ATTRIBUTES};
  g_CommandsList["pagein"] = {&CommandPagein, &CommandPageinHelp,
                              DEBUGGER_COMMAND_PAGEIN_ATTRIBUTES};
  g_CommandsList["test"] = {&CommandTest, &CommandTestHelp,
                            DEBUGGER_COMMAND_TEST_ATTRIBUTES};
  g_CommandsList[".test"] = {&CommandTest, &CommandTestHelp,
                             DEBUGGER_COMMAND_TEST_ATTRIBUTES};
  g_CommandsList["cpu"] = {&CommandCpu, &CommandCpuHelp,
                           DEBUGGER_COMMAND_CPU_ATTRIBUTES};
  g_CommandsList["wrmsr"] = {&CommandWrmsr, &CommandWrmsrHelp,
                             DEBUGGER_COMMAND_WRMSR_ATTRIBUTES};
  g_CommandsList["rdmsr"] = {&CommandRdmsr, &CommandRdmsrHelp,
                             DEBUGGER_COMMAND_RDMSR_ATTRIBUTES};
  g_CommandsList["!va2pa"] = {&CommandVa2pa, &CommandVa2paHelp,
                              DEBUGGER_COMMAND_VA2PA_ATTRIBUTES};
  g_CommandsList["!pa2va"] = {&CommandPa2va, &CommandPa2vaHelp,
                              DEBUGGER_COMMAND_PA2VA_ATTRIBUTES};
  g_CommandsList[".formats"] = {&CommandFormats, &CommandFormatsHelp,
                                DEBUGGER_COMMAND_FORMATS_ATTRIBUTES};
  g_CommandsList[".format"] = {&CommandFormats, &CommandFormatsHelp,
                               DEBUGGER_COMMAND_FORMATS_ATTRIBUTES};
  g_CommandsList["!pte"] = {&CommandPte, &CommandPteHelp,
                            DEBUGGER_COMMAND_PTE_ATTRIBUTES};
  g_CommandsList["~"] = {&CommandCore, &CommandCoreHelp,
                         DEBUGGER_COMMAND_CORE_ATTRIBUTES};
  g_CommandsList["core"] = {&CommandCore, &CommandCoreHelp,
                            DEBUGGER_COMMAND_CORE_ATTRIBUTES};
  g_CommandsList["!apic"] = {&CommandApic, &CommandApicHelp,
                             DEBUGGER_COMMAND_APIC_ATTRIBUTES};
  g_CommandsList["!lapic"] = {&CommandApic, &CommandApicHelp,
                              DEBUGGER_COMMAND_APIC_ATTRIBUTES};
  g_CommandsList["!localapic"] = {&CommandApic, &CommandApicHelp,
                                  DEBUGGER_COMMAND_APIC_ATTRIBUTES};
  g_CommandsList["!ioapic"] = {&CommandIoapic, &CommandIoapicHelp,
                               DEBUGGER_COMMAND_IOAPIC_ATTRIBUTES};
  g_CommandsList["!monitor"] = {&CommandMonitor, &CommandMonitorHelp,
                                DEBUGGER_COMMAND_MONITOR_ATTRIBUTES};
  g_CommandsList["!vmcall"] = {&CommandVmcall, &CommandVmcallHelp,
                               DEBUGGER_COMMAND_VMCALL_ATTRIBUTES};
  g_CommandsList["!epthook"] = {&CommandEptHook, &CommandEptHookHelp,
                                DEBUGGER_COMMAND_EPTHOOK_ATTRIBUTES};
  g_CommandsList["bp"] = {&CommandBp, &CommandBpHelp,
                          DEBUGGER_COMMAND_BP_ATTRIBUTES};
  g_CommandsList["bl"] = {&CommandBl, &CommandBlHelp,
                          DEBUGGER_COMMAND_BD_ATTRIBUTES};
  g_CommandsList["be"] = {&CommandBe, &CommandBeHelp,
                          DEBUGGER_COMMAND_BD_ATTRIBUTES};
  g_CommandsList["bd"] = {&CommandBd, &CommandBdHelp,
                          DEBUGGER_COMMAND_BD_ATTRIBUTES};
  g_CommandsList["bc"] = {&CommandBc, &CommandBcHelp,
                          DEBUGGER_COMMAND_BD_ATTRIBUTES};
  g_CommandsList["!epthook2"] = {&CommandEptHook2, &CommandEptHook2Help,
                                 DEBUGGER_COMMAND_EPTHOOK2_ATTRIBUTES};
  g_CommandsList["!cpuid"] = {&CommandCpuid, &CommandCpuidHelp,
                              DEBUGGER_COMMAND_CPUID_ATTRIBUTES};
  g_CommandsList["!msrread"] = {&CommandMsrread, &CommandMsrreadHelp,
                                DEBUGGER_COMMAND_MSRREAD_ATTRIBUTES};
  g_CommandsList["!msread"] = {&CommandMsrread, &CommandMsrreadHelp,
                               DEBUGGER_COMMAND_MSRREAD_ATTRIBUTES};
  g_CommandsList["!msrwrite"] = {&CommandMsrwrite, &CommandMsrwriteHelp,
                                 DEBUGGER_COMMAND_MSRWRITE_ATTRIBUTES};
  g_CommandsList["!tsc"] = {&CommandTsc, &CommandTscHelp,
                            DEBUGGER_COMMAND_TSC_ATTRIBUTES};
  g_CommandsList["!pmc"] = {&CommandPmc, &CommandPmcHelp,
                            DEBUGGER_COMMAND_PMC_ATTRIBUTES};
  g_CommandsList["!crwrite"] = {&CommandCrwrite, &CommandCrwriteHelp,
                                DEBUGGER_COMMAND_CRWRITE_ATTRIBUTES};
  g_CommandsList["!dr"] = {&CommandDr, &CommandDrHelp,
                           DEBUGGER_COMMAND_DR_ATTRIBUTES};
  g_CommandsList["!ioin"] = {&CommandIoin, &CommandIoinHelp,
                             DEBUGGER_COMMAND_IOIN_ATTRIBUTES};
  g_CommandsList["!ioout"] = {&CommandIoout, &CommandIooutHelp,
                              DEBUGGER_COMMAND_IOOUT_ATTRIBUTES};
  g_CommandsList["!iout"] = {&CommandIoout, &CommandIooutHelp,
                             DEBUGGER_COMMAND_IOOUT_ATTRIBUTES};
  g_CommandsList["!exception"] = {&CommandException, &CommandExceptionHelp,
                                  DEBUGGER_COMMAND_EXCEPTION_ATTRIBUTES};
  g_CommandsList["!interrupt"] = {&CommandInterrupt, &CommandInterruptHelp,
                                  DEBUGGER_COMMAND_INTERRUPT_ATTRIBUTES};
  g_CommandsList["!syscall"] = {&CommandSyscallAndSysret, &CommandSyscallHelp,
                                DEBUGGER_COMMAND_SYSCALL_ATTRIBUTES};
  g_CommandsList["!syscall2"] = {&CommandSyscallAndSysret, &CommandSyscallHelp,
                                 DEBUGGER_COMMAND_SYSCALL_ATTRIBUTES};
  g_CommandsList["!sysret"] = {&CommandSyscallAndSysret, &CommandSysretHelp,
                               DEBUGGER_COMMAND_SYSRET_ATTRIBUTES};
  g_CommandsList["!sysret2"] = {&CommandSyscallAndSysret, &CommandSysretHelp,
                                DEBUGGER_COMMAND_SYSRET_ATTRIBUTES};
  g_CommandsList["!mode"] = {&CommandMode, &CommandModeHelp,
                             DEBUGGER_COMMAND_MODE_ATTRIBUTES};
  g_CommandsList["!trace"] = {&CommandTrace, &CommandTraceHelp,
                              DEBUGGER_COMMAND_TRACE_ATTRIBUTES};
  g_CommandsList["!hide"] = {&CommandHide, &CommandHideHelp,
                             DEBUGGER_COMMAND_HIDE_ATTRIBUTES};
  g_CommandsList["!unhide"] = {&CommandUnhide, &CommandUnhideHelp,
                               DEBUGGER_COMMAND_UNHIDE_ATTRIBUTES};
  g_CommandsList["!measure"] = {&CommandMeasure, &CommandMeasureHelp,
                                DEBUGGER_COMMAND_MEASURE_ATTRIBUTES};
  g_CommandsList["lm"] = {&CommandLm, &CommandLmHelp,
                          DEBUGGER_COMMAND_LM_ATTRIBUTES};
  g_CommandsList["p"] = {&CommandP, &CommandPHelp,
                         DEBUGGER_COMMAND_P_ATTRIBUTES};
  g_CommandsList["pr"] = {&CommandP, &CommandPHelp,
                          DEBUGGER_COMMAND_P_ATTRIBUTES};
  g_CommandsList["t"] = {&CommandT, &CommandTHelp,
                         DEBUGGER_COMMAND_T_ATTRIBUTES};
  g_CommandsList["tr"] = {&CommandT, &CommandTHelp,
                          DEBUGGER_COMMAND_T_ATTRIBUTES};
  g_CommandsList["i"] = {&CommandI, &CommandIHelp,
                         DEBUGGER_COMMAND_I_ATTRIBUTES};
  g_CommandsList["ir"] = {&CommandI, &CommandIHelp,
                          DEBUGGER_COMMAND_I_ATTRIBUTES};
  g_CommandsList["gu"] = {&CommandGu, &CommandGuHelp,
                          DEBUGGER_COMMAND_GU_ATTRIBUTES};
  g_CommandsList["db"] = {&CommandReadMemoryAndDisassembler,
                          &CommandReadMemoryAndDisassemblerHelp,
                          DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["dc"] = {&CommandReadMemoryAndDisassembler,
                          &CommandReadMemoryAndDisassemblerHelp,
                          DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["dd"] = {&CommandReadMemoryAndDisassembler,
                          &CommandReadMemoryAndDisassemblerHelp,
                          DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["dq"] = {&CommandReadMemoryAndDisassembler,
                          &CommandReadMemoryAndDisassemblerHelp,
                          DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["!db"] = {&CommandReadMemoryAndDisassembler,
                           &CommandReadMemoryAndDisassemblerHelp,
                           DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["!dc"] = {&CommandReadMemoryAndDisassembler,
                           &CommandReadMemoryAndDisassemblerHelp,
                           DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["!dd"] = {&CommandReadMemoryAndDisassembler,
                           &CommandReadMemoryAndDisassemblerHelp,
                           DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["!dq"] = {&CommandReadMemoryAndDisassembler,
                           &CommandReadMemoryAndDisassemblerHelp,
                           DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["!u"] = {&CommandReadMemoryAndDisassembler,
                          &CommandReadMemoryAndDisassemblerHelp,
                          DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["u"] = {&CommandReadMemoryAndDisassembler,
                         &CommandReadMemoryAndDisassemblerHelp,
                         DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["!u64"] = {&CommandReadMemoryAndDisassembler,
                            &CommandReadMemoryAndDisassemblerHelp,
                            DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["u64"] = {&CommandReadMemoryAndDisassembler,
                           &CommandReadMemoryAndDisassemblerHelp,
                           DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["!u2"] = {&CommandReadMemoryAndDisassembler,
                           &CommandReadMemoryAndDisassemblerHelp,
                           DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["u2"] = {&CommandReadMemoryAndDisassembler,
                          &CommandReadMemoryAndDisassemblerHelp,
                          DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["!u32"] = {&CommandReadMemoryAndDisassembler,
                            &CommandReadMemoryAndDisassemblerHelp,
                            DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["u32"] = {&CommandReadMemoryAndDisassembler,
                           &CommandReadMemoryAndDisassemblerHelp,
                           DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES};
  g_CommandsList["eb"] = {&CommandEditMemory, &CommandEditMemoryHelp,
                          DEBUGGER_COMMAND_E_ATTRIBUTES};
  g_CommandsList["ed"] = {&CommandEditMemory, &CommandEditMemoryHelp,
                          DEBUGGER_COMMAND_E_ATTRIBUTES};
  g_CommandsList["eq"] = {&CommandEditMemory, &CommandEditMemoryHelp,
                          DEBUGGER_COMMAND_E_ATTRIBUTES};
  g_CommandsList["!eb"] = {&CommandEditMemory, &CommandEditMemoryHelp,
                           DEBUGGER_COMMAND_E_ATTRIBUTES};
  g_CommandsList["!ed"] = {&CommandEditMemory, &CommandEditMemoryHelp,
                           DEBUGGER_COMMAND_E_ATTRIBUTES};
  g_CommandsList["!eq"] = {&CommandEditMemory, &CommandEditMemoryHelp,
                           DEBUGGER_COMMAND_E_ATTRIBUTES};
  g_CommandsList["sb"] = {&CommandSearchMemory, &CommandSearchMemoryHelp,
                          DEBUGGER_COMMAND_S_ATTRIBUTES};
  g_CommandsList["sd"] = {&CommandSearchMemory, &CommandSearchMemoryHelp,
                          DEBUGGER_COMMAND_S_ATTRIBUTES};
  g_CommandsList["sq"] = {&CommandSearchMemory, &CommandSearchMemoryHelp,
                          DEBUGGER_COMMAND_S_ATTRIBUTES};
  g_CommandsList["!sb"] = {&CommandSearchMemory, &CommandSearchMemoryHelp,
                           DEBUGGER_COMMAND_S_ATTRIBUTES};
  g_CommandsList["!sd"] = {&CommandSearchMemory, &CommandSearchMemoryHelp,
                           DEBUGGER_COMMAND_S_ATTRIBUTES};
  g_CommandsList["!sq"] = {&CommandSearchMemory, &CommandSearchMemoryHelp,
                           DEBUGGER_COMMAND_S_ATTRIBUTES};
  g_CommandsList["r"] = {&CommandR, &CommandRHelp,
                         DEBUGGER_COMMAND_R_ATTRIBUTES};
  g_CommandsList[".sympath"] = {&CommandSympath, &CommandSympathHelp,
                                DEBUGGER_COMMAND_SYMPATH_ATTRIBUTES};
  g_CommandsList["sympath"] = {&CommandSympath, &CommandSympathHelp,
                               DEBUGGER_COMMAND_SYMPATH_ATTRIBUTES};
  g_CommandsList[".sym"] = {&CommandSym, &CommandSymHelp,
                            DEBUGGER_COMMAND_SYM_ATTRIBUTES};
  g_CommandsList["sym"] = {&CommandSym, &CommandSymHelp,
                           DEBUGGER_COMMAND_SYM_ATTRIBUTES};
  g_CommandsList["x"] = {&CommandX, &CommandXHelp,
                         DEBUGGER_COMMAND_X_ATTRIBUTES};
  g_CommandsList["prealloc"] = {&CommandPrealloc, &CommandPreallocHelp,
                                DEBUGGER_COMMAND_PREALLOC_ATTRIBUTES};
  g_CommandsList["preallocate"] = {&CommandPrealloc, &CommandPreallocHelp,
                                   DEBUGGER_COMMAND_PREALLOC_ATTRIBUTES};
  g_CommandsList["preallocation"] = {&CommandPrealloc, &CommandPreallocHelp,
                                     DEBUGGER_COMMAND_PREALLOC_ATTRIBUTES};
  g_CommandsList["preactivate"] = {&CommandPreactivate, &CommandPreactivateHelp,
                                   DEBUGGER_COMMAND_PREACTIVATE_ATTRIBUTES};
  g_CommandsList["preactive"] = {&CommandPreactivate, &CommandPreactivateHelp,
                                 DEBUGGER_COMMAND_PREACTIVATE_ATTRIBUTES};
  g_CommandsList["preactivation"] = {&CommandPreactivate,
                                     &CommandPreactivateHelp,
                                     DEBUGGER_COMMAND_PREACTIVATE_ATTRIBUTES};
  g_CommandsList["k"] = {&CommandK, &CommandKHelp,
                         DEBUGGER_COMMAND_K_ATTRIBUTES};
  g_CommandsList["kd"] = {&CommandK, &CommandKHelp,
                          DEBUGGER_COMMAND_K_ATTRIBUTES};
  g_CommandsList["kq"] = {&CommandK, &CommandKHelp,
                          DEBUGGER_COMMAND_K_ATTRIBUTES};
  g_CommandsList["dt"] = {&CommandDtAndStruct, &CommandDtHelp,
                          DEBUGGER_COMMAND_DT_ATTRIBUTES};
  g_CommandsList["!dt"] = {&CommandDtAndStruct, &CommandDtHelp,
                           DEBUGGER_COMMAND_DT_ATTRIBUTES};
  g_CommandsList["struct"] = {&CommandDtAndStruct, &CommandStructHelp,
                              DEBUGGER_COMMAND_STRUCT_ATTRIBUTES};
  g_CommandsList["structure"] = {&CommandDtAndStruct, &CommandStructHelp,
                                 DEBUGGER_COMMAND_STRUCT_ATTRIBUTES};
  g_CommandsList[".pe"] = {&CommandPe, &CommandPeHelp,
                           DEBUGGER_COMMAND_PE_ATTRIBUTES};
  g_CommandsList["!rev"] = {&CommandRev, &CommandRevHelp,
                            DEBUGGER_COMMAND_REV_ATTRIBUTES};
  g_CommandsList["rev"] = {&CommandRev, &CommandRevHelp,
                           DEBUGGER_COMMAND_REV_ATTRIBUTES};
  g_CommandsList["!track"] = {&CommandTrack, &CommandTrackHelp,
                              DEBUGGER_COMMAND_TRACK_ATTRIBUTES};
  g_CommandsList["track"] = {&CommandTrack, &CommandTrackHelp,
                             DEBUGGER_COMMAND_TRACK_ATTRIBUTES};
  g_CommandsList[".dump"] = {&CommandDump, &CommandDumpHelp,
                             DEBUGGER_COMMAND_DUMP_ATTRIBUTES};
  g_CommandsList["dump"] = {&CommandDump, &CommandDumpHelp,
                            DEBUGGER_COMMAND_DUMP_ATTRIBUTES};
  g_CommandsList["!dump"] = {&CommandDump, &CommandDumpHelp,
                             DEBUGGER_COMMAND_DUMP_ATTRIBUTES};
  g_CommandsList["a"] = {&CommandAssemble, &CommandAssembleHelp,
                         DEBUGGER_COMMAND_A_ATTRIBUTES};
  g_CommandsList["asm"] = {&CommandAssemble, &CommandAssembleHelp,
                           DEBUGGER_COMMAND_A_ATTRIBUTES};
  g_CommandsList["assemble"] = {&CommandAssemble, &CommandAssembleHelp,
                                DEBUGGER_COMMAND_A_ATTRIBUTES};
  g_CommandsList["assembly"] = {&CommandAssemble, &CommandAssembleHelp,
                                DEBUGGER_COMMAND_A_ATTRIBUTES};
  g_CommandsList["!a"] = {&CommandAssemble, &CommandAssembleHelp,
                          DEBUGGER_COMMAND_A_ATTRIBUTES};
  g_CommandsList["!asm"] = {&CommandAssemble, &CommandAssembleHelp,
                            DEBUGGER_COMMAND_A_ATTRIBUTES};
  g_CommandsList["!assemble"] = {&CommandAssemble, &CommandAssembleHelp,
                                 DEBUGGER_COMMAND_A_ATTRIBUTES};
  g_CommandsList["!assembly"] = {&CommandAssemble, &CommandAssembleHelp,
                                 DEBUGGER_COMMAND_A_ATTRIBUTES};
  g_CommandsList["!pcitree"] = {&CommandPcitree, &CommandPcitreeHelp,
                                DEBUGGER_COMMAND_PCITREE_ATTRIBUTES};
  g_CommandsList["!pcietree"] = {&CommandPcitree, &CommandPcitreeHelp,
                                 DEBUGGER_COMMAND_PCITREE_ATTRIBUTES};
  g_CommandsList["!pcicam"] = {&CommandPcicam, &CommandPcicamHelp,
                               DEBUGGER_COMMAND_PCICAM_ATTRIBUTES};
  g_CommandsList["!idt"] = {&CommandIdt, &CommandIdtHelp,
                            DEBUGGER_COMMAND_IDT_ATTRIBUTES};
  g_CommandsList["!hw_clk"] = {&CommandHwClk, &CommandHwClkHelp,
                               DEBUGGER_COMMAND_HWDBG_HW_CLK_ATTRIBUTES};
  g_CommandsList["!hw_clock"] = {&CommandHwClk, &CommandHwClkHelp,
                                 DEBUGGER_COMMAND_HWDBG_HW_CLK_ATTRIBUTES};
  g_CommandsList["!hwdbg_clock"] = {&CommandHwClk, &CommandHwClkHelp,
                                    DEBUGGER_COMMAND_HWDBG_HW_CLK_ATTRIBUTES};
  g_CommandsList["!hwdbg_clock"] = {&CommandHwClk, &CommandHwClkHelp,
                                    DEBUGGER_COMMAND_HWDBG_HW_CLK_ATTRIBUTES};
  g_CommandsList["!hw"] = {&CommandHw, &CommandHwHelp,
                           DEBUGGER_COMMAND_HWDBG_HW_ATTRIBUTES};
}
