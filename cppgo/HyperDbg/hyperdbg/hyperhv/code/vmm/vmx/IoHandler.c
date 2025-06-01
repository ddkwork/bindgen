
#include "pch.h"

VOID IoHandleIoVmExits(VIRTUAL_MACHINE_STATE *VCpu,
                       VMX_EXIT_QUALIFICATION_IO_INSTRUCTION IoQualification,
                       RFLAGS Flags) {
  UINT16 Port = 0;
  UINT32 Count = 0;
  UINT32 Size = 0;
  PGUEST_REGS GuestRegs = VCpu->Regs;

  union {
    unsigned char *AsBytePtr;
    unsigned short *AsWordPtr;
    unsigned long *AsDwordPtr;
    void *AsPtr;
    UINT64 AsUint64;
  } PortValue;

  if (IoQualification.StringInstruction) {
    PortValue.AsPtr =
        (PVOID)(IoQualification.DirectionOfAccess == AccessIn ? GuestRegs->rdi
                                                              : GuestRegs->rsi);
  } else {
    PortValue.AsPtr = &GuestRegs->rax;
  }
  Port = (UINT16)IoQualification.PortNumber;
  Count = IoQualification.RepPrefixed ? (GuestRegs->rcx & 0xffffffff) : 1;
  Size = (UINT32)(IoQualification.SizeOfAccess + 1);
  switch (IoQualification.DirectionOfAccess) {
    case AccessIn:
      if (IoQualification.StringInstruction) {
        switch (Size) {
          case 1:
            IoInByteString(Port, (UINT8 *)PortValue.AsBytePtr, Count);
            break;
          case 2:
            IoInWordString(Port, (UINT16 *)PortValue.AsWordPtr, Count);
            break;
          case 4:
            IoInDwordString(Port, (UINT32 *)PortValue.AsDwordPtr, Count);
            break;
        }
      } else {
        switch (Size) {
          case 1:
            *PortValue.AsBytePtr = IoInByte(Port);
            break;
          case 2:
            *PortValue.AsWordPtr = IoInWord(Port);
            break;
          case 4:
            *PortValue.AsDwordPtr = IoInDword(Port);
            break;
        }
      }
      break;
    case AccessOut:
      if (IoQualification.StringInstruction) {
        switch (Size) {
          case 1:
            IoOutByteString(Port, (UINT8 *)PortValue.AsBytePtr, Count);
            break;
          case 2:
            IoOutWordString(Port, (UINT16 *)PortValue.AsWordPtr, Count);
            break;
          case 4:
            IoOutDwordString(Port, (UINT32 *)PortValue.AsDwordPtr, Count);
            break;
        }
      } else {
        switch (Size) {
          case 1:
            IoOutByte(Port, *PortValue.AsBytePtr);
            break;
          case 2:
            IoOutWord(Port, *PortValue.AsWordPtr);
            break;
          case 4:
            IoOutDword(Port, *PortValue.AsDwordPtr);
            break;
        }
      }
      break;
  }
  if (IoQualification.StringInstruction) {
    UINT64 GpReg = IoQualification.DirectionOfAccess == AccessIn
                       ? GuestRegs->rdi
                       : GuestRegs->rsi;
    if (Flags.DirectionFlag) {
      GpReg -= Count * Size;
    } else {
      GpReg += Count * Size;
    }
    if (IoQualification.RepPrefixed) {
      GuestRegs->rcx = 0;
    }
  }
}

BOOLEAN IoHandleSetIoBitmap(VIRTUAL_MACHINE_STATE *VCpu, UINT32 Port) {
  if (Port <= 0x7FFF) {
    SetBit(Port, (unsigned long *)VCpu->IoBitmapVirtualAddressA);
  } else if ((0x8000 <= Port) && (Port <= 0xFFFF)) {
    SetBit(Port - 0x8000, (unsigned long *)VCpu->IoBitmapVirtualAddressB);
  } else {
    return FALSE;
  }
  return TRUE;
}

VOID IoHandlePerformIoBitmapChange(VIRTUAL_MACHINE_STATE *VCpu, UINT32 Port) {
  if (Port == DEBUGGER_EVENT_ALL_IO_PORTS) {
    memset((void *)VCpu->IoBitmapVirtualAddressA, 0xFF, PAGE_SIZE);
    memset((void *)VCpu->IoBitmapVirtualAddressB, 0xFF, PAGE_SIZE);
  } else {
    IoHandleSetIoBitmap(VCpu, Port);
  }
}

VOID IoHandlePerformIoBitmapReset(VIRTUAL_MACHINE_STATE *VCpu) {
  memset((void *)VCpu->IoBitmapVirtualAddressA, 0x0, PAGE_SIZE);
  memset((void *)VCpu->IoBitmapVirtualAddressB, 0x0, PAGE_SIZE);
}
