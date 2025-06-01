
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

VOID CommandApicHelp() {
  ShowMessages(
      "!apic : shows the details of Local APIC in both xAPIC and x2APIC "
      "modes.\n\n");
  ShowMessages("syntax : \t!apic\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : !apic\n");
}

BOOLEAN CommandApicSendRequest(DEBUGGER_APIC_REQUEST_TYPE ApicType,
                               PVOID ApicBuffer, UINT32 ExpectedRequestSize,
                               PBOOLEAN IsUsingX2APIC) {
  BOOL Status;
  ULONG ReturnedLength;
  PDEBUGGER_APIC_REQUEST ApicRequest;
  UINT32 RequestSize = 0;
  RequestSize = SIZEOF_DEBUGGER_APIC_REQUEST + ExpectedRequestSize;
  ApicRequest = (PDEBUGGER_APIC_REQUEST)malloc(RequestSize);
  if (ApicRequest == NULL) {
    return FALSE;
  }
  RtlZeroMemory(ApicRequest, RequestSize);
  ApicRequest->ApicType = ApicType;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    if (!KdSendApicActionPacketsToDebuggee(ApicRequest, RequestSize)) {
      free(ApicRequest);
      return FALSE;
    } else {
      *IsUsingX2APIC = ApicRequest->IsUsingX2APIC;
      RtlCopyMemory(
          ApicBuffer,
          (PVOID)(((CHAR *)ApicRequest) + sizeof(DEBUGGER_APIC_REQUEST)),
          ExpectedRequestSize);
      free(ApicRequest);
      return TRUE;
    }
  } else {
    AssertShowMessageReturnStmt(
        g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status = DeviceIoControl(g_DeviceHandle, IOCTL_PERFROM_ACTIONS_ON_APIC,
                             ApicRequest, SIZEOF_DEBUGGER_APIC_REQUEST,
                             ApicRequest, RequestSize, &ReturnedLength, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      free(ApicRequest);
      return FALSE;
    }
    if (ReturnedLength != RequestSize &&
        ReturnedLength != SIZEOF_DEBUGGER_APIC_REQUEST) {
      ShowMessages("err, apic request failed\n");
      free(ApicRequest);
      return FALSE;
    }
    if (ApicRequest->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
      *IsUsingX2APIC = ApicRequest->IsUsingX2APIC;
      RtlCopyMemory(
          ApicBuffer,
          (PVOID)(((CHAR *)ApicRequest) + sizeof(DEBUGGER_APIC_REQUEST)),
          ExpectedRequestSize);
      free(ApicRequest);
      return TRUE;
    } else {
      ShowErrorMessage(ApicRequest->KernelStatus);
      free(ApicRequest);
      return FALSE;
    }
  }
}

BOOLEAN HyperDbgGetLocalApic(PLAPIC_PAGE LocalApic, PBOOLEAN IsUsingX2APIC) {
  return CommandApicSendRequest(DEBUGGER_APIC_REQUEST_TYPE_READ_LOCAL_APIC,
                                LocalApic, sizeof(LAPIC_PAGE), IsUsingX2APIC);
}

VOID CommandApic(vector<CommandToken> CommandTokens, string Command) {
  BOOLEAN IsUsingX2APIC = FALSE;
  UINT8 i = 0, j = 0;
  UINT32 k = 0;
  UINT32 Reg = 0;
  LAPIC_PAGE LocalApic = {0};
  if (CommandTokens.size() != 1) {
    ShowMessages(
        "incorrect use of the '%s'\n\n",
        GetCaseSensitiveStringFromCommandToken(CommandTokens.at(0)).c_str());
    CommandApicHelp();
    return;
  }
  if (HyperDbgGetLocalApic(&LocalApic, &IsUsingX2APIC) == FALSE) {
    return;
  }
  ShowMessages(
      "***  (%s Mode) PHYSICAL LAPIC ID = %u, Ver = 0x%x, MaxLvtEntry = %d, "
      "DirectedEOI = P%d/E%d, (SW: '%s')\n"
      "     -> TPR = 0x%08x,  PPR = 0x%08x\n"
      "     -> LDR = 0x%08x,  SVR = 0x%08x,  Err = 0x%08x\n"
      "     -> LVT_INT0 = 0x%08x,  LVT_INT1 = 0x%08x\n"
      "     -> LVT_CMCI = 0x%08x,  LVT_PMCR = 0x%08x\n"
      "     -> LVT_TMR  = 0x%08x,  LVT_TSR  = 0x%08x\n"
      "     -> LVT_ERR  = 0x%08x\n"
      "     -> InitialCount = 0x%08x, CurrentCount = 0x%08x, DivideConfig = "
      "0x%08x\n",
      IsUsingX2APIC ? "X2APIC" : "XAPIC", LocalApic.Id >> 24, LocalApic.Version,
      (LocalApic.Version & 0xFF0000) >> 16, (LocalApic.Version >> 24) & 1,
      (LocalApic.SpuriousInterruptVector >> 12) & 1,
      (LocalApic.SpuriousInterruptVector & LAPIC_SVR_FLAG_SW_ENABLE) != 0
          ? "Enabled"
          : "Disabled",
      LocalApic.TPR, LocalApic.ProcessorPriority, LocalApic.LogicalDestination,
      LocalApic.SpuriousInterruptVector, LocalApic.ErrorStatus,
      LocalApic.LvtLINT0, LocalApic.LvtLINT1, LocalApic.LvtCmci,
      LocalApic.LvtPerfMonCounters, LocalApic.LvtTimer,
      LocalApic.LvtThermalSensor, LocalApic.LvtError, LocalApic.InitialCount,
      LocalApic.CurrentCount, LocalApic.DivideConfiguration);
  ShowMessages("ISR : ");
  for (i = 0; i < 8; i++) {
    k = 1;
    Reg = (UINT32)LocalApic.ISR[i * 4];
    for (j = 0; j < 32; j++) {
      if (0 != (Reg & k)) {
        ShowMessages("0x%02hhx ", (UINT8)(i * 32 + j));
      }
      k = k << 1;
    }
  }
  ShowMessages("\n");
  ShowMessages("TMR : ");
  for (i = 0; i < 8; i++) {
    k = 1;
    Reg = (UINT32)LocalApic.TMR[i * 4];
    for (j = 0; j < 32; j++) {
      if (Reg & k) {
        ShowMessages("0x%02hhx ", (UINT8)(i * 32 + j));
      }
      k = k << 1;
    }
  }
  ShowMessages("\n");
  ShowMessages("IRR : ");
  for (i = 0; i < 8; i++) {
    k = 1;
    Reg = (UINT32)LocalApic.IRR[i * 4];
    for (j = 0; j < 32; j++) {
      if (Reg & k) {
        ShowMessages("0x%02hhx ", (UINT8)(i * 32 + j));
      }
      k = k << 1;
    }
  }
  ShowMessages("\n");
}
