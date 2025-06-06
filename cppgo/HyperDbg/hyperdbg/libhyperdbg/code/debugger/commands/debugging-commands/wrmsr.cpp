
#include "pch.h"

VOID CommandWrmsrHelp() {
  ShowMessages("wrmsr : writes on a model-specific register (MSR).\n\n");
  ShowMessages(
      "syntax : \twrmsr [Msr (hex)] [Value (hex)] [core CoreNumber (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : wrmsr c0000082 fffff8077356f010\n");
  ShowMessages("\t\te.g : wrmsr c0000082 @rcx\n");
  ShowMessages("\t\te.g : wrmsr c0000082 @rcx+@rdx+12\n");
  ShowMessages("\t\te.g : wrmsr c0000082 fffff8077356f010 core 2\n");
}

VOID CommandWrmsr(vector<CommandToken> CommandTokens, string Command) {
  BOOL Status;
  UINT64 Msr;
  DEBUGGER_READ_AND_WRITE_ON_MSR MsrWriteRequest = {0};
  BOOL IsNextCoreId = FALSE;
  BOOL SetMsr = FALSE;
  BOOL SetValue = FALSE;
  UINT64 Value = 0;
  UINT32 CoreNumer = DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES;
  BOOLEAN IsFirstCommand = TRUE;
  if (CommandTokens.size() >= 6) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandWrmsrHelp();
    return;
  }
  for (auto Section : CommandTokens) {
    if (IsFirstCommand == TRUE) {
      IsFirstCommand = FALSE;
      continue;
    }
    if (IsNextCoreId) {
      if (!ConvertTokenToUInt32(Section, &CoreNumer)) {
        ShowMessages("please specify a correct hex value for core id\n\n");
        CommandWrmsrHelp();
        return;
      }
      IsNextCoreId = FALSE;
      continue;
    }
    if (CompareLowerCaseStrings(Section, "core")) {
      IsNextCoreId = TRUE;
      continue;
    }
    if (!SetMsr) {
      if (!ConvertTokenToUInt64(Section, &Msr)) {
        ShowMessages("please specify a correct hex value to be read\n\n");
        CommandWrmsrHelp();
        return;
      } else {
        SetMsr = TRUE;
        continue;
      }
    }
    if (SetMsr) {
      if (!SymbolConvertNameOrExprToAddress(
              GetCaseSensitiveStringFromCommandToken(Section), &Value)) {
        ShowMessages(
            "please specify a correct hex value or an expression to put on the "
            "msr\n\n");
        CommandWrmsrHelp();
        return;
      } else {
        SetValue = TRUE;
        continue;
      }
    }
  }
  if (!SetMsr) {
    ShowMessages("please specify a correct hex value to write\n\n");
    CommandWrmsrHelp();
    return;
  }
  if (!SetValue) {
    ShowMessages("please specify a correct hex value to put on msr\n\n");
    CommandWrmsrHelp();
    return;
  }
  if (IsNextCoreId) {
    ShowMessages("please specify a correct hex value for core\n\n");
    CommandWrmsrHelp();
    return;
  }
  AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED,
                              AssertReturn);
  MsrWriteRequest.ActionType = DEBUGGER_MSR_WRITE;
  MsrWriteRequest.Msr = Msr;
  MsrWriteRequest.CoreNumber = CoreNumer;
  MsrWriteRequest.Value = Value;
  Status = DeviceIoControl(
      g_DeviceHandle, IOCTL_DEBUGGER_READ_OR_WRITE_MSR, &MsrWriteRequest,
      SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR, NULL, NULL, NULL, NULL);
  if (!Status) {
    ShowMessages(
        "ioctl failed with code (%x), either msr index or core id is invalid\n",
        GetLastError());
    return;
  }
  ShowMessages("\n");
}
