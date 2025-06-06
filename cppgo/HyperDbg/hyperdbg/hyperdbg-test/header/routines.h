
#pragma once
BOOLEAN TestCase(std::vector<std::string> &TestCase);
extern "C" {
extern void inline AsmTest();
}
VOID TestCreateLookupTable(HANDLE PipeHandle, PVOID KernelInformation,
                           UINT32 KernelInformationSize);
std::string Uint64ToString(UINT64 value);
BOOLEAN StringReplace(std::string &str, const std::string &from,
                      const std::string &to);
std::string ConvertToString(char *Str);
