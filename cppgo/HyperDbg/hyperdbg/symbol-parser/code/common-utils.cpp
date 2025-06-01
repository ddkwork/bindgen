
#include "pch.h"

BOOLEAN IsFileExists(const std::string &FileName) {
  struct stat buffer;
  return (stat(FileName.c_str(), &buffer) == 0);
}

BOOLEAN IsDirExists(const std::string &DirPath) {
  DWORD Ftyp = GetFileAttributesA(DirPath.c_str());
  if (Ftyp == INVALID_FILE_ATTRIBUTES) return FALSE;
  if (Ftyp & FILE_ATTRIBUTE_DIRECTORY) return TRUE;
  return FALSE;
}

BOOLEAN CreateDirectoryRecursive(const std::string &Path) {
  size_t Pos = 0;
  do {
    Pos = Path.find_first_of("\\/", Pos + 1);
    CreateDirectoryA(Path.substr(0, Pos).c_str(), NULL);
  } while (Pos != std::string::npos && Pos <= Path.size());
  if (IsDirExists(Path)) return TRUE;
  return FALSE;
}

VOID SplitPathAndArgs(std::vector<std::string> &Qargs,
                      const std::string &Command) {
  int Len = (int)Command.length();
  bool Qot = false, Sqot = false;
  int ArgLen;
  for (int i = 0; i < Len; i++) {
    int start = i;
    if (Command[i] == '\"') {
      Qot = true;
    } else if (Command[i] == '\'')
      Sqot = true;
    if (Qot) {
      i++;
      start++;
      while (i < Len && Command[i] != '\"') i++;
      if (i < Len) Qot = false;
      ArgLen = i - start;
      i++;
    } else if (Sqot) {
      i++;
      while (i < Len && Command[i] != '\'') i++;
      if (i < Len) Sqot = false;
      ArgLen = i - start;
      i++;
    } else {
      while (i < Len && Command[i] != ' ') i++;
      ArgLen = i - start;
    }
    string Temp = Command.substr(start, ArgLen);
    if (!Temp.empty() && Temp != " ") {
      Qargs.push_back(Temp);
    }
  }
  /*
  for (int i = 0; i < Qargs.size(); i++)
  {
      std::cout << Qargs[i] << std::endl;
  }

  std::cout << Qargs.size();

  if (Qot || Sqot)
      std::cout << "One of the quotes is open\n";
  */
}

/**
 * @brief general split command
 *
 * @param s target string
 * @param c splitter (delimiter)
 * @return const vector<string>
 */
const std::vector<std::string> Split(const std::string &s, const char &c) {
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
