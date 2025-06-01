
#include "pch.h"

VOID MsrHandleRdmsrVmexit(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 TargetMsr;
  MSR Msr = {0};
  PGUEST_REGS GuestRegs = VCpu->Regs;
  TargetMsr = GuestRegs->rcx & 0xffffffff;
  if ((TargetMsr <= 0x00001FFF) ||
      ((0xC0000000 <= TargetMsr) && (TargetMsr <= 0xC0001FFF)) ||
      (TargetMsr >= RESERVED_MSR_RANGE_LOW &&
       (TargetMsr <= RESERVED_MSR_RANGE_HI))) {
    switch (TargetMsr) {
      case IA32_SYSENTER_CS:
        VmxVmread64P(VMCS_GUEST_SYSENTER_CS, &Msr.Flags);
        break;
      case IA32_SYSENTER_ESP:
        VmxVmread64P(VMCS_GUEST_SYSENTER_ESP, &Msr.Flags);
        break;
      case IA32_SYSENTER_EIP:
        VmxVmread64P(VMCS_GUEST_SYSENTER_EIP, &Msr.Flags);
        break;
      case IA32_GS_BASE:
        VmxVmread64P(VMCS_GUEST_GS_BASE, &Msr.Flags);
        break;
      case IA32_FS_BASE:
        VmxVmread64P(VMCS_GUEST_FS_BASE, &Msr.Flags);
        break;
      case HV_X64_MSR_GUEST_IDLE:
        break;
      default:
        if (TargetMsr <= 0xfff &&
            TestBit(TargetMsr, (unsigned long *)g_MsrBitmapInvalidMsrs) !=
                NULL64_ZERO) {
          EventInjectGeneralProtection();
          return;
        }
        Msr.Flags = __readmsr(TargetMsr);
        if (GuestRegs->rcx == IA32_EFER) {
          IA32_EFER_REGISTER MsrEFER;
          MsrEFER.AsUInt = Msr.Flags;
          MsrEFER.SyscallEnable = TRUE;
          Msr.Flags = MsrEFER.AsUInt;
        }
        break;
    }
    GuestRegs->rax = NULL64_ZERO;
    GuestRegs->rdx = NULL64_ZERO;
    GuestRegs->rax = Msr.Fields.Low;
    GuestRegs->rdx = Msr.Fields.High;
  } else {
    EventInjectGeneralProtection();
    return;
  }
}

VOID MsrHandleWrmsrVmexit(VIRTUAL_MACHINE_STATE *VCpu) {
  UINT32 TargetMsr;
  BOOLEAN UnusedIsKernel;
  MSR Msr = {0};
  PGUEST_REGS GuestRegs = VCpu->Regs;
  TargetMsr = GuestRegs->rcx & 0xffffffff;
  Msr.Fields.Low = (ULONG)GuestRegs->rax;
  Msr.Fields.High = (ULONG)GuestRegs->rdx;
  if ((TargetMsr <= 0x00001FFF) ||
      ((0xC0000000 <= TargetMsr) && (TargetMsr <= 0xC0001FFF)) ||
      (TargetMsr >= RESERVED_MSR_RANGE_LOW &&
       (TargetMsr <= RESERVED_MSR_RANGE_HI))) {
    switch (TargetMsr) {
      case IA32_DS_AREA:
      case IA32_FS_BASE:
      case IA32_GS_BASE:
      case IA32_KERNEL_GS_BASE:
      case IA32_LSTAR:
      case IA32_SYSENTER_EIP:
      case IA32_SYSENTER_ESP:
        if (!CheckAddressCanonicality(Msr.Flags, &UnusedIsKernel)) {
          EventInjectGeneralProtection();
          return;
        }
        break;
    }
    switch (TargetMsr) {
      case IA32_SYSENTER_CS:
        VmxVmwrite64(VMCS_GUEST_SYSENTER_CS, Msr.Flags);
        break;
      case IA32_SYSENTER_ESP:
        VmxVmwrite64(VMCS_GUEST_SYSENTER_ESP, Msr.Flags);
        break;
      case IA32_SYSENTER_EIP:
        VmxVmwrite64(VMCS_GUEST_SYSENTER_EIP, Msr.Flags);
        break;
      case IA32_GS_BASE:
        VmxVmwrite64(VMCS_GUEST_GS_BASE, Msr.Flags);
        break;
      case IA32_FS_BASE:
        VmxVmwrite64(VMCS_GUEST_FS_BASE, Msr.Flags);
        break;
      default:
        if (g_TransparentMode && (TargetMsr >= RESERVED_MSR_RANGE_LOW &&
                                  (TargetMsr <= RESERVED_MSR_RANGE_HI))) {
          LogInfo(
              "WRMSR attempts to write to a reserved MSR range. MSR: %x, rax: "
              "%llx, rdx: %llx, from: %llx",
              TargetMsr, GuestRegs->rax, GuestRegs->rdx, VCpu->LastVmexitRip);
          EventInjectGeneralProtection();
          return;
        }
        __writemsr((unsigned long)GuestRegs->rcx, Msr.Flags);
        break;
    }
  } else {
    EventInjectGeneralProtection();
    return;
  }
}

BOOLEAN MsrHandleSetMsrBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 Msr,
                              BOOLEAN ReadDetection, BOOLEAN WriteDetection) {
  if (!ReadDetection && !WriteDetection) {
    return FALSE;
  }
  if (Msr <= 0x00001FFF) {
    if (ReadDetection) {
      SetBit(Msr, (unsigned long *)VCpu->MsrBitmapVirtualAddress);
    }
    if (WriteDetection) {
      SetBit(Msr, (unsigned long *)VCpu->MsrBitmapVirtualAddress + 2048);
    }
  } else if ((0xC0000000 <= Msr) && (Msr <= 0xC0001FFF)) {
    if (ReadDetection) {
      SetBit(Msr - 0xC0000000,
             (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 1024));
    }
    if (WriteDetection) {
      SetBit(Msr - 0xC0000000,
             (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 3072));
    }
  } else {
    return FALSE;
  }
  return TRUE;
}

BOOLEAN MsrHandleUnSetMsrBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 Msr,
                                BOOLEAN ReadDetection, BOOLEAN WriteDetection) {
  if (!ReadDetection && !WriteDetection) {
    return FALSE;
  }
  if (Msr <= 0x00001FFF) {
    if (ReadDetection) {
      ClearBit(Msr, (unsigned long *)VCpu->MsrBitmapVirtualAddress);
    }
    if (WriteDetection) {
      ClearBit(Msr, (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 2048));
    }
  } else if ((0xC0000000 <= Msr) && (Msr <= 0xC0001FFF)) {
    if (ReadDetection) {
      ClearBit(Msr - 0xC0000000,
               (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 1024));
    }
    if (WriteDetection) {
      ClearBit(Msr - 0xC0000000,
               (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 3072));
    }
  } else {
    return FALSE;
  }
  return TRUE;
}

VOID MsrHandleFilterMsrReadBitmap(VIRTUAL_MACHINE_STATE *VCpu) {
  ClearBit(0x102, (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 1024));
  ClearBit(0xe7, (unsigned long *)VCpu->MsrBitmapVirtualAddress);
  ClearBit(0xe8, (unsigned long *)VCpu->MsrBitmapVirtualAddress);
}

VOID MsrHandleFilterMsrWriteBitmap(VIRTUAL_MACHINE_STATE *VCpu) {
  ClearBit(0x102, (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 3072));
  ClearBit(0xe7, (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 2048));
  ClearBit(0xe8, (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 2048));
  ClearBit(0x48, (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 2048));
  ClearBit(0x49, (unsigned long *)(VCpu->MsrBitmapVirtualAddress + 2048));
}

VOID MsrHandlePerformMsrBitmapReadChange(VIRTUAL_MACHINE_STATE *VCpu,
                                         UINT32 MsrMask) {
  if (MsrMask == DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS) {
    memset((void *)VCpu->MsrBitmapVirtualAddress, 0xff, 2048);
    MsrHandleFilterMsrReadBitmap(VCpu);
  } else {
    MsrHandleSetMsrBitmap(VCpu, MsrMask, TRUE, FALSE);
  }
}

VOID MsrHandlePerformMsrBitmapReadReset(VIRTUAL_MACHINE_STATE *VCpu) {
  memset((void *)VCpu->MsrBitmapVirtualAddress, 0x0, 2048);
}

VOID MsrHandlePerformMsrBitmapWriteChange(VIRTUAL_MACHINE_STATE *VCpu,
                                          UINT32 MsrMask) {
  if (MsrMask == DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS) {
    memset((void *)((UINT64)VCpu->MsrBitmapVirtualAddress + 2048), 0xff, 2048);
    MsrHandleFilterMsrWriteBitmap(VCpu);
  } else {
    MsrHandleSetMsrBitmap(VCpu, MsrMask, FALSE, TRUE);
  }
}

VOID MsrHandlePerformMsrBitmapWriteReset(VIRTUAL_MACHINE_STATE *VCpu) {
  memset((void *)((UINT64)VCpu->MsrBitmapVirtualAddress + 2048), 0x0, 2048);
}
