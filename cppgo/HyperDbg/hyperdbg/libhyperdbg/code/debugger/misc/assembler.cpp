
#include "pch.h"

VOID AssembleData::ParseAssemblyData() {
  std::string RawAsm = AsmRaw;
  RawAsm.erase(std::remove(RawAsm.begin(), RawAsm.end(), '\n'), RawAsm.end());
  std::regex MultipleSpaces(" +");
  RawAsm = std::regex_replace(RawAsm, MultipleSpaces, " ");
  std::vector<std::string> AssemblyInstructions;
  size_t Pos = 0;
  std::string Delimiter = ";";
  while ((Pos = RawAsm.find(Delimiter)) != std::string::npos) {
    std::string Token = RawAsm.substr(0, Pos);
    if (!Token.empty()) {
      AssemblyInstructions.push_back(Token);
    }
    RawAsm.erase(0, Pos + Delimiter.length());
  }
  if (!RawAsm.empty()) {
    AssemblyInstructions.push_back(RawAsm);
  }
  for (auto &InstructionLine : AssemblyInstructions) {
    std::string Expr{};
    UINT64 ExprAddr{};
    size_t Start{};
    while ((Start = InstructionLine.find('<', Start)) != std::string::npos) {
      size_t End = InstructionLine.find('>', Start);
      if (End != std::string::npos) {
        std::string Expr = InstructionLine.substr(Start + 1, End - Start - 1);
        if (!SymbolConvertNameOrExprToAddress(Expr, &ExprAddr) &&
            ExprAddr == 0) {
          ShowMessages("err, failed to resolve the symbol [ %s ].\n",
                       Expr.c_str());
          Start += Expr.size() + 2;
          continue;
        }
        std::ostringstream Oss;
        Oss << std::hex << std::showbase << ExprAddr;
        InstructionLine.replace(Start, End - Start + 1, Oss.str());
      } else {
        break;
      }
      Start += Expr.size() + 2;
    }
  }
  auto ApndSemCln = [](std::string a, std::string b) {
    return std::move(a) + ';' + std::move(b);
  };
  AsmFixed = std::accumulate(std::next(AssemblyInstructions.begin()),
                             AssemblyInstructions.end(),
                             AssemblyInstructions.at(0), ApndSemCln);
  if (!AsmFixed.empty() && AsmFixed.back() == ';') {
    AsmFixed.pop_back();
  }
  while (!AsmFixed.empty() && AsmFixed.back() == ';') {
    AsmFixed.pop_back();
  }
}

INT AssembleData::Assemble(UINT64 StartAddr, ks_arch Arch, INT Mode,
                           INT Syntax) {
  ks_engine *Ks;
  KsErr = ks_open(Arch, Mode, &Ks);
  if (KsErr != KS_ERR_OK) {
    ShowMessages("err, failed on ks_open()");
    return -1;
  }
  if (Syntax) {
    KsErr = ks_option(Ks, KS_OPT_SYNTAX, Syntax);
    if (KsErr != KS_ERR_OK) {
      ShowMessages("err, failed on ks_option() with error code = %u\n", KsErr);
    }
  }
  KsErr = ks_option(Ks, KS_OPT_SYM_RESOLVER, 0);
  if (KsErr != KS_ERR_OK) {
    ShowMessages("err, failed on ks_option() with error code = %u\n", KsErr);
  }
  if (ks_asm(Ks, AsmFixed.c_str(), StartAddr, &EncodedBytes, &BytesCount,
             &StatementCount)) {
    KsErr = ks_errno(Ks);
    ShowMessages("err, failed on ks_asm() with count = %lu, error code = %u\n",
                 (int)StatementCount, KsErr);
  } else {
    if (BytesCount == 0) {
      ShowMessages(
          "err, the assemble operation returned no bytes, most likely due to "
          "incorrect formatting of assembly snippet\n");
    } else {
      ShowMessages(
          "generated assembly: %lu (decimal) bytes, %lu (decimal) statements "
          "==>> ",
          (int)BytesCount, (int)StatementCount);
      size_t i;
      ShowMessages("%s = ", AsmFixed.c_str());
      for (i = 0; i < BytesCount; i++) {
        ShowMessages("%02x ", EncodedBytes[i]);
        EncBytesIntVec.push_back(static_cast<UINT64>(EncodedBytes[i]));
      }
      ShowMessages("\n");
      ks_close(Ks);
      return 0;
    }
  }
  ks_close(Ks);
  return -1;
}

AssembleData *create_AssembleData() { return new AssembleData; }

BOOLEAN HyperDbgAssembleGetLength(const CHAR *AssemblyCode, UINT64 StartAddress,
                                  UINT32 *Length) {
  AssembleData AssembleData;
  std::string AsmCode = AssemblyCode;
  AssembleData.AsmRaw = AsmCode;
  AssembleData.ParseAssemblyData();
  if (AssembleData.Assemble(StartAddress)) {
    return FALSE;
  } else {
    *Length = (UINT32)AssembleData.BytesCount;
    return TRUE;
  }
}

BOOLEAN HyperDbgAssemble(const CHAR *AssemblyCode, UINT64 StartAddress,
                         PVOID BufferToStoreAssembledData, UINT32 BufferSize) {
  AssembleData AssembleData;
  std::string AsmCode = AssemblyCode;
  AssembleData.AsmRaw = AsmCode;
  AssembleData.ParseAssemblyData();
  if (AssembleData.Assemble(StartAddress)) {
    return FALSE;
  } else {
    if (AssembleData.BytesCount > BufferSize) {
      return FALSE;
    }
    memcpy(BufferToStoreAssembledData, AssembleData.EncodedBytes,
           AssembleData.BytesCount);
    return TRUE;
  }
}
