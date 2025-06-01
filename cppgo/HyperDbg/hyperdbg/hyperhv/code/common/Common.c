
#include "pch.h"

_Use_decl_annotations_

    BOOLEAN
    CommonAffinityBroadcastToProcessors(ULONG ProcessorNumber,
                                        RunOnLogicalCoreFunc Routine) {
  KIRQL OldIrql;
  KeSetSystemAffinityThread((KAFFINITY)(1ULL << ProcessorNumber));
  OldIrql = KeRaiseIrqlToDpcLevel();
  Routine(ProcessorNumber);
  KeLowerIrql(OldIrql);
  KeRevertToUserAffinityThread();
  return TRUE;
}

PCHAR CommonGetProcessNameFromProcessControlBlock(PEPROCESS Eprocess) {
  PCHAR Result = 0;
  Result = (CHAR *)PsGetProcessImageFileName(Eprocess);
  return Result;
}

BOOLEAN CommonIsStringStartsWith(const char *pre, const char *str) {
  size_t lenpre = strlen(pre), lenstr = strlen(str);
  return lenstr < lenpre ? FALSE : memcmp(pre, str, lenpre) == 0;
}

VOID CommonCpuidInstruction(UINT32 Func, UINT32 SubFunc, int *CpuInfo) {
  __cpuidex(CpuInfo, Func, SubFunc);
}

BOOLEAN CommonIsGuestOnUsermode32Bit() {
  UINT64 CsSel = NULL64_ZERO;
  CsSel = HvGetCsSelector();
  if (CsSel == KGDT64_R0_CODE) {
    return FALSE;
  } else if ((CsSel & ~3) == KGDT64_R3_CODE) {
    return FALSE;
  } else if ((CsSel & ~3) == KGDT64_R3_CMCODE) {
    return TRUE;
  } else {
    LogError("Err, unknown value for cs, cannot determine wow64 mode");
  }
  return FALSE;
}

VOID CommonWriteDebugInformation(VIRTUAL_MACHINE_STATE *VCpu) {
  LogError(
      "HyperDbg cannot recover from this error, please provide the following "
      "information through the Git issues");
  LogInfo("Target RIP: %llx\n", VCpu->LastVmexitRip);
  CHAR Instruction[MAXIMUM_INSTR_SIZE] = {0};
  MemoryMapperReadMemorySafeOnTargetProcess(VCpu->LastVmexitRip, Instruction,
                                            MAXIMUM_INSTR_SIZE);
  for (size_t i = 0; i < MAXIMUM_INSTR_SIZE; i++) {
    Log("%02X ", Instruction[i] & 0xffU);
  }
  Log("\n");
  DisassemblerShowOneInstructionInVmxRootMode((PVOID)VCpu->LastVmexitRip,
                                              CommonIsGuestOnUsermode32Bit());
  Log("\n");
  Log("RAX=%016llx RBX=%016llx RCX=%016llx\n"
      "RDX=%016llx RSI=% 016llx RDI=%016llx\n"
      "RIP=%016llx RSP=%016llx RBP=%016llx\n"
      "R8 =%016llx R9 =%016llx R10=%016llx\n"
      "R11=%016llx R12=%016llx R13=%016llx\n"
      "R14=%016llx R15=%016llx\n",
      VCpu->Regs->rax, VCpu->Regs->rbx, VCpu->Regs->rcx, VCpu->Regs->rdx,
      VCpu->Regs->rsi, VCpu->Regs->rdi, VCpu->LastVmexitRip, VCpu->Regs->rsp,
      VCpu->Regs->rbp, VCpu->Regs->r8, VCpu->Regs->r9, VCpu->Regs->r10,
      VCpu->Regs->r11, VCpu->Regs->r12, VCpu->Regs->r13, VCpu->Regs->r14,
      VCpu->Regs->r15);
}

BOOLEAN CommonIsXCr0Valid(XCR0 XCr0) {
  CPUID_EAX_0D_ECX_00 XCr0CpuidInfo;
  UINT64 UnsupportedXCr0Bits;
  CommonCpuidInstruction(CPUID_EXTENDED_STATE_INFORMATION, 0,
                         (INT32 *)&XCr0CpuidInfo);
  UnsupportedXCr0Bits =
      ~(XCr0CpuidInfo.Eax.AsUInt | (UINT64)XCr0CpuidInfo.Edx.AsUInt << 32);
  if ((XCr0.AsUInt & UnsupportedXCr0Bits) != 0) {
    return FALSE;
  }
  if (XCr0.X87 != 1) {
    return FALSE;
  }
  if (XCr0.Sse == 0 && XCr0.Avx == 1) {
    return FALSE;
  }
  if (XCr0.Avx == 0) {
    if ((XCr0.Opmask | XCr0.ZmmHi256 | XCr0.ZmmHi16) == 1) {
      return FALSE;
    }
  }
  if (XCr0.Bndcsr != XCr0.Bndreg) {
    return FALSE;
  }
  if ((XCr0.Opmask | XCr0.ZmmHi256 | XCr0.ZmmHi16) == 1) {
    if ((XCr0.Opmask & XCr0.ZmmHi256 & XCr0.ZmmHi16) != 1) {
      return FALSE;
    }
  }
  return TRUE;
}
