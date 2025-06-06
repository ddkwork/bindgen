
#include "pch.h"
extern BOOLEAN g_RtmSupport;
extern UINT32 g_VirtualAddressWidth;

string SeparateTo64BitValue(UINT64 Value) {
  ostringstream OstringStream;
  string Temp;
  OstringStream << setw(16) << setfill('0') << hex << Value;
  Temp = OstringStream.str();
  Temp.insert(8, 1, '`');
  return Temp;
}

VOID PrintBits(const UINT32 Size, const void *Ptr) {
  unsigned char *Buf = (unsigned char *)Ptr;
  unsigned char Byte;
  int i, j;
  for (i = Size - 1; i >= 0; i--) {
    for (j = 7; j >= 0; j--) {
      Byte = (Buf[i] >> j) & 1;
      ShowMessages("%u", Byte);
    }
    ShowMessages(" ", Byte);
  }
}

BOOL Replace(std::string &str, const std::string &from, const std::string &to) {
  size_t start_pos = str.find(from);
  if (start_pos == std::string::npos) return FALSE;
  str.replace(start_pos, from.size(), to);
  return TRUE;
}

VOID ReplaceAll(string &str, const string &from, const string &to) {
  size_t SartPos = 0;
  if (from.empty()) return;
  while ((SartPos = str.find(from, SartPos)) != std::string::npos) {
    str.replace(SartPos, from.length(), to);
    SartPos += to.length();
  }
}

const vector<string> Split(const string &s, const char &c) {
  string buff{""};
  vector<string> v;
  for (auto n : s) {
    if (n != c)
      buff += n;
    else if (n == c && !buff.empty()) {
      v.push_back(buff);
      buff.clear();
    }
  }
  if (!buff.empty()) v.push_back(buff);
  return v;
}

BOOLEAN IsNumber(const string &str) {
  return !str.empty() &&
         (str.find_first_not_of("[0123456789]") == std::string::npos);
}

BOOLEAN IsHexNotation(const string &s) {
  BOOLEAN IsAnyThing = FALSE;
  for (auto &CptrChar : s) {
    IsAnyThing = TRUE;
    if (!isxdigit(CptrChar)) {
      return FALSE;
    }
  }
  if (IsAnyThing) {
    return TRUE;
  }
  return FALSE;
}

BOOLEAN IsDecimalNotation(const string &s) {
  BOOLEAN IsAnyThing = FALSE;
  for (auto &CptrChar : s) {
    IsAnyThing = TRUE;
    if (!isdigit(CptrChar)) {
      return FALSE;
    }
  }
  if (IsAnyThing) {
    return TRUE;
  }
  return FALSE;
}

vector<char> HexToBytes(const string &hex) {
  vector<char> Bytes;
  for (unsigned int i = 0; i < hex.length(); i += 2) {
    std::string byteString = hex.substr(i, 2);
    char byte = (char)strtol(byteString.c_str(), NULL, 16);
    Bytes.push_back(byte);
  }
  return Bytes;
}

BOOLEAN ConvertStringToUInt64(string TextToConvert, PUINT64 Result) {
  BOOLEAN IsDecimal = FALSE;
  if (TextToConvert.rfind("0x", 0) == 0 || TextToConvert.rfind("0X", 0) == 0 ||
      TextToConvert.rfind("\\x", 0) == 0 ||
      TextToConvert.rfind("\\X", 0) == 0) {
    TextToConvert = TextToConvert.erase(0, 2);
  } else if (TextToConvert.rfind('x', 0) == 0 ||
             TextToConvert.rfind('X', 0) == 0) {
    TextToConvert = TextToConvert.erase(0, 1);
  } else if (TextToConvert.rfind("0n", 0) == 0 ||
             TextToConvert.rfind("0N", 0) == 0 ||
             TextToConvert.rfind("\\n", 0) == 0 ||
             TextToConvert.rfind("\\N", 0) == 0) {
    TextToConvert = TextToConvert.erase(0, 2);
    IsDecimal = TRUE;
  } else if (TextToConvert.rfind('n', 0) == 0 ||
             TextToConvert.rfind('N', 0) == 0) {
    TextToConvert = TextToConvert.erase(0, 1);
    IsDecimal = TRUE;
  }
  TextToConvert.erase(remove(TextToConvert.begin(), TextToConvert.end(), '`'),
                      TextToConvert.end());
  if (IsDecimal) {
    if (!IsDecimalNotation(TextToConvert)) {
      return FALSE;
    } else {
      errno = 0;
      char *unparsed = NULL;
      const char *s = TextToConvert.c_str();
      const unsigned long long int n = strtoull(s, &unparsed, 10);
      if (errno || (!n && s == unparsed)) {
        return FALSE;
      }
      *Result = n;
      return TRUE;
    }
  } else {
    if (!IsHexNotation(TextToConvert)) {
      return FALSE;
    } else {
      const char *Text = TextToConvert.c_str();
      errno = 0;
      unsigned long long result = strtoull(Text, NULL, 16);
      *Result = result;
      if (errno == EINVAL) {
        return FALSE;
      } else if (errno == ERANGE) {
        return TRUE;
      }
      return TRUE;
    }
  }
}

BOOLEAN ConvertStringToUInt32(string TextToConvert, PUINT32 Result) {
  BOOLEAN IsDecimal = FALSE;
  if (TextToConvert.rfind("0x", 0) == 0 || TextToConvert.rfind("0X", 0) == 0 ||
      TextToConvert.rfind("\\x", 0) == 0 ||
      TextToConvert.rfind("\\X", 0) == 0) {
    TextToConvert = TextToConvert.erase(0, 2);
  } else if (TextToConvert.rfind('x', 0) == 0 ||
             TextToConvert.rfind('X', 0) == 0) {
    TextToConvert = TextToConvert.erase(0, 1);
  } else if (TextToConvert.rfind("0n", 0) == 0 ||
             TextToConvert.rfind("0N", 0) == 0 ||
             TextToConvert.rfind("\\n", 0) == 0 ||
             TextToConvert.rfind("\\N", 0) == 0) {
    TextToConvert = TextToConvert.erase(0, 2);
    IsDecimal = TRUE;
  } else if (TextToConvert.rfind('n', 0) == 0 ||
             TextToConvert.rfind('N', 0) == 0) {
    TextToConvert = TextToConvert.erase(0, 1);
    IsDecimal = TRUE;
  }
  TextToConvert.erase(remove(TextToConvert.begin(), TextToConvert.end(), '`'),
                      TextToConvert.end());
  if (IsDecimal) {
    if (!IsDecimalNotation(TextToConvert)) {
      return FALSE;
    } else {
      try {
        int i = std::stoi(TextToConvert);
        *Result = i;
        return TRUE;
      } catch (std::invalid_argument const &) {
        return FALSE;
      } catch (std::out_of_range const &) {
        return FALSE;
      }
      return FALSE;
    }
  } else {
    if (!IsHexNotation(TextToConvert)) {
      return FALSE;
    } else {
      UINT32 TempResult;
      TempResult = stoi(TextToConvert, nullptr, 16);
      *Result = TempResult;
      return TRUE;
    }
  }
}

BOOLEAN ConvertTokenToUInt64(CommandToken TargetToken, PUINT64 Result) {
  std::string TargetTokenValue = std::get<1>(TargetToken);
  return ConvertStringToUInt64(TargetTokenValue, Result);
}

std::string GetCaseSensitiveStringFromCommandToken(CommandToken TargetToken) {
  std::string TargetTokenValue = std::get<1>(TargetToken);
  return TargetTokenValue;
}

std::string GetLowerStringFromCommandToken(CommandToken TargetToken) {
  std::string TargetTokenValue = std::get<2>(TargetToken);
  return TargetTokenValue;
}

BOOLEAN CompareLowerCaseStrings(CommandToken TargetToken,
                                const char *StringToCompare) {
  std::string TargetTokenValue = std::get<2>(TargetToken);
  return _stricmp(TargetTokenValue.c_str(), StringToCompare) == 0;
}

BOOLEAN IsTokenBracketString(CommandToken TargetToken) {
  CommandParsingTokenType TargetTokenValue = std::get<0>(TargetToken);
  return TargetTokenValue == CommandParsingTokenType::BracketString;
}

BOOLEAN ConvertTokenToUInt32(CommandToken TargetToken, PUINT32 Result) {
  std::string TargetTokenValue = std::get<1>(TargetToken);
  return ConvertStringToUInt32(TargetTokenValue, Result);
}

BOOLEAN HasEnding(string const &fullString, string const &ending) {
  if (fullString.length() >= ending.length()) {
    return (0 == fullString.compare(fullString.length() - ending.length(),
                                    ending.length(), ending));
  } else {
    return FALSE;
  }
}

BOOLEAN ValidateIP(const string &ip) {
  vector<string> list = Split(ip, '.');
  if (list.size() != 4) return FALSE;
  for (string str : list) {
    if (!IsNumber(str) || stoi(str) > 255 || stoi(str) < 0) return FALSE;
  }
  return TRUE;
}

BOOLEAN VmxSupportDetection() { return AsmVmxSupportDetection(); }

BOOL SetPrivilege(HANDLE Token, LPCTSTR Privilege, BOOL EnablePrivilege) {
  TOKEN_PRIVILEGES Tp;
  LUID Luid;
  if (!LookupPrivilegeValue(NULL, Privilege, &Luid)) {
    ShowMessages("err, in LookupPrivilegeValue (%x)\n", GetLastError());
    return FALSE;
  }
  Tp.PrivilegeCount = 1;
  Tp.Privileges[0].Luid = Luid;
  if (EnablePrivilege)
    Tp.Privileges[0].Attributes = SE_PRIVILEGE_ENABLED;
  else
    Tp.Privileges[0].Attributes = 0;
  if (!AdjustTokenPrivileges(Token, FALSE, &Tp, sizeof(TOKEN_PRIVILEGES),
                             (PTOKEN_PRIVILEGES)NULL, (PDWORD)NULL)) {
    ShowMessages("err, in AdjustTokenPrivileges (%x)\n", GetLastError());
    return FALSE;
  }
  if (GetLastError() == ERROR_NOT_ALL_ASSIGNED) {
    ShowMessages(
        "err, the token does not have the specified (debug) privilege (ACCESS "
        "DENIED!)\n");
    ShowMessages("make sure to run it with administrator privileges\n");
    return FALSE;
  }
  return TRUE;
}

static inline void ltrim(std::string &s) {
  s.erase(s.begin(), std::find_if(s.begin(), s.end(),
                                  [](int ch) { return !std::isspace(ch); }));
}

static inline void rtrim(std::string &s) {
  s.erase(std::find_if(s.rbegin(), s.rend(),
                       [](int ch) { return !std::isspace(ch); })
              .base(),
          s.end());
}

void Trim(std::string &s) {
  ltrim(s);
  rtrim(s);
}

std::string RemoveSpaces(std::string str) {
  str.erase(remove(str.begin(), str.end(), ' '), str.end());
  return str;
}

BOOLEAN IsFileExistA(const char *FileName) {
  struct stat buffer;
  return (stat(FileName, &buffer) == 0);
}

BOOLEAN IsFileExistW(const wchar_t *FileName) {
  struct _stat64i32 buffer;
  return (_wstat(FileName, &buffer) == 0);
}

BOOLEAN IsEmptyString(char *Text) {
  size_t Len;
  if (Text == NULL || Text[0] == '\0') {
    return TRUE;
  }
  Len = strlen(Text);
  for (size_t i = 0; i < Len; i++) {
    if (Text[i] != ' ' && Text[i] != '\t' && Text[i] != '\n') {
      return FALSE;
    }
  }
  return TRUE;
}

VOID GetConfigFilePath(PWCHAR ConfigPath) {
  WCHAR CurrentPath[MAX_PATH] = {0};
  GetModuleFileNameW(NULL, CurrentPath, MAX_PATH);
  PathRemoveFileSpecW(CurrentPath);
  PathCombineW(ConfigPath, CurrentPath, CONFIG_FILE_NAME);
}

std::vector<std::string> ListDirectory(const std::string &Directory,
                                       const std::string &Extension) {
  WIN32_FIND_DATAA FindData;
  HANDLE Find = INVALID_HANDLE_VALUE;
  std::string FullPath = Directory + "\\" + Extension;
  std::vector<std::string> DirList;
  Find = FindFirstFileA(FullPath.c_str(), &FindData);
  if (Find == INVALID_HANDLE_VALUE)
    throw std::runtime_error("invalid handle value! please check your path...");
  while (FindNextFileA(Find, &FindData) != 0) {
    DirList.push_back(Directory + "\\" + std::string(FindData.cFileName));
  }
  FindClose(Find);
  return DirList;
}

VOID StringToWString(std::wstring &ws, const std::string &s) {
  std::wstring WsTmp(s.begin(), s.end());
  ws = WsTmp;
}

size_t FindCaseInsensitive(std::string Input, std::string ToSearch,
                           size_t Pos) {
  std::transform(Input.begin(), Input.end(), Input.begin(), ::tolower);
  std::transform(ToSearch.begin(), ToSearch.end(), ToSearch.begin(), ::tolower);
  return Input.find(ToSearch, Pos);
}

size_t FindCaseInsensitiveW(std::wstring Input, std::wstring ToSearch,
                            size_t Pos) {
  std::transform(Input.begin(), Input.end(), Input.begin(), ::tolower);
  std::transform(ToSearch.begin(), ToSearch.end(), ToSearch.begin(), ::tolower);
  return Input.find(ToSearch, Pos);
}

char *ConvertStringVectorToCharPointerArray(const std::string &s) {
  char *pc = new char[s.size() + 1];
  std::strcpy(pc, s.c_str());
  return pc;
}

VOID CommonCpuidInstruction(UINT32 Func, UINT32 SubFunc, int *CpuInfo) {
  __cpuidex(CpuInfo, Func, SubFunc);
}

UINT32 Getx86VirtualAddressWidth() {
  int Regs[4];
  CommonCpuidInstruction(CPUID_ADDR_WIDTH, 0, Regs);
  return ((Regs[0] >> 8) & 0x0ff);
}

BOOLEAN CheckCpuSupportRtm() {
  int Regs1[4];
  int Regs2[4];
  BOOLEAN Result;
  CommonCpuidInstruction(0, 0, Regs1);
  CommonCpuidInstruction(7, 0, Regs2);
  Result = Regs1[0] >= 0x7 && (Regs2[1] & 0x4800) == 0x4800;
  return Result;
}

BOOLEAN CheckAddressCanonicality(UINT64 VAddr, PBOOLEAN IsKernelAddress) {
  UINT64 Addr = (UINT64)VAddr;
  UINT64 MaxVirtualAddrLowHalf, MinVirtualAddressHighHalf;
  UINT32 AddrWidth = g_VirtualAddressWidth;
  MaxVirtualAddrLowHalf = ((UINT64)1ull << (AddrWidth - 1)) - 1;
  MinVirtualAddressHighHalf = ~MaxVirtualAddrLowHalf;
  if ((Addr > MaxVirtualAddrLowHalf) && (Addr < MinVirtualAddressHighHalf)) {
    *IsKernelAddress = FALSE;
    return FALSE;
  }
  if (MinVirtualAddressHighHalf < Addr) {
    *IsKernelAddress = TRUE;
  } else {
    *IsKernelAddress = FALSE;
  }
  return TRUE;
}

BOOLEAN CheckAddressValidityUsingTsx(UINT64 Address) {
  UINT32 Status = 0;
  BOOLEAN Result = FALSE;
  CHAR TempContent;
  if ((Status = _xbegin()) == _XBEGIN_STARTED) {
    TempContent = *(CHAR *)Address;
    _xend();
    Result = TRUE;
  } else {
    Result = FALSE;
  }
  return Result;
}

BOOLEAN CheckAccessValidityAndSafety(UINT64 TargetAddress, UINT32 Size) {
  BOOLEAN IsKernelAddress;
  BOOLEAN Result = FALSE;
  if (!CheckAddressCanonicality(TargetAddress, &IsKernelAddress)) {
    return FALSE;
  }
  if (IsKernelAddress) {
    return FALSE;
  }
  if (g_RtmSupport) {
    UINT64 AddressToCheck =
        (CHAR *)TargetAddress + Size - ((CHAR *)PAGE_ALIGN(TargetAddress));
    if (AddressToCheck > PAGE_SIZE) {
      UINT32 ReadSize = 0;
      while (Size != 0) {
        ReadSize = (UINT32)((UINT64)PAGE_ALIGN(TargetAddress + PAGE_SIZE) -
                            TargetAddress);
        if (ReadSize == PAGE_SIZE && Size < PAGE_SIZE) {
          ReadSize = Size;
        }
        if (!CheckAddressValidityUsingTsx(TargetAddress)) {
          return FALSE;
        }

        Size = Size - ReadSize;
        TargetAddress = TargetAddress + ReadSize;
      }
    } else {
      if (!CheckAddressValidityUsingTsx(TargetAddress)) {
        return FALSE;
      }
    }
  } else {
    return FALSE;
    UINT64 AddressToCheck =
        (CHAR *)TargetAddress + Size - ((CHAR *)PAGE_ALIGN(TargetAddress));
    if (AddressToCheck > PAGE_SIZE) {
      UINT32 ReadSize = 0;
      while (Size != 0) {
        ReadSize = (UINT32)((UINT64)PAGE_ALIGN(TargetAddress + PAGE_SIZE) -
                            TargetAddress);
        if (ReadSize == PAGE_SIZE && Size < PAGE_SIZE) {
          ReadSize = Size;
        }
        try {
          UINT64 ReadingTest = *((UINT64 *)TargetAddress);
        } catch (...) {
          return FALSE;
        }

        Size = Size - ReadSize;
        TargetAddress = TargetAddress + ReadSize;
      }
    } else {
      try {
        UINT64 ReadingTest = *((UINT64 *)TargetAddress);
      } catch (...) {
        return FALSE;
      }
    }
  }
  return TRUE;
}

UINT32 Log2Ceil(UINT32 n) {
  if (n == 0) return 0;
  n--;
  UINT32 log2Floor = 0;
  while (n >>= 1) {
    log2Floor++;
  }
  return log2Floor + 1;
}
