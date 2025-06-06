
#include "common.h"
#include "kdcom.h"
BOOLEAN Uart16550SetBaud(_Inout_ PCPPORT Port, ULONG Rate);
UART_STATUS Uart16550PutByte(_Inout_ PCPPORT Port, UCHAR Byte,
                             BOOLEAN BusyWait);
UART_STATUS Uart16550GetByte(_Inout_ PCPPORT Port, _Out_ PUCHAR Byte);
CPPORT g_PortDetails = {0};

static VOID WritePortWithIndex8(_In_ PCPPORT Port, const UCHAR Index,
                                const UCHAR Value) {
  PUCHAR Pointer;
  Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
  WRITE_PORT_UCHAR(Pointer, Value);
  return;
}

static UCHAR ReadPortWithIndex8(_In_ PCPPORT Port, const UCHAR Index) {
  PUCHAR Pointer;
  Pointer = (PUCHAR)(Port->Address + Index * Port->ByteWidth);
  return (UCHAR)READ_PORT_UCHAR(Pointer);
}

VOID KdHyperDbgTest(UINT16 Byte) {
  UNREFERENCED_PARAMETER(Byte);
  CPPORT TempPort = {0};
  TempPort.Address = (PUCHAR)0x2f8;
  TempPort.BaudRate = 0x01c200;
  TempPort.Flags = 0;
  TempPort.ByteWidth = 1;
  TempPort.Write = WritePortWithIndex8;
  TempPort.Read = ReadPortWithIndex8;
  Uart16550PutByte(&TempPort, 0x42, TRUE);
  for (size_t i = 0; i < 100; i++) {
  }
}

VOID KdHyperDbgPrepareDebuggeeConnectionPort(UINT32 PortAddress,
                                             UINT32 Baudrate) {
  UINT64 PortAddressUChar = 0ULL;
  PortAddressUChar = PortAddressUChar | PortAddress;
  g_PortDetails.Address = (PUCHAR)PortAddressUChar;
  g_PortDetails.BaudRate = Baudrate;
  g_PortDetails.Flags = 0;
  g_PortDetails.ByteWidth = 1;
  g_PortDetails.Write = WritePortWithIndex8;
  g_PortDetails.Read = ReadPortWithIndex8;
}

VOID KdHyperDbgSendByte(UCHAR Byte, BOOLEAN BusyWait) {
  Uart16550PutByte(&g_PortDetails, Byte, BusyWait);
}

BOOLEAN KdHyperDbgRecvByte(PUCHAR RecvByte) {
  if (Uart16550GetByte(&g_PortDetails, RecvByte) == UartSuccess) {
    return TRUE;
  }
  return FALSE;
}

BOOLEAN Uart16550InitializePortCommon(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions, _Inout_ PCPPORT Port,
    BOOLEAN MemoryMapped, UCHAR AccessSize, UCHAR BitWidth)

{
  UCHAR RegisterValue;
  UNREFERENCED_PARAMETER(LoadOptions);
  UartpSetAccess(Port, MemoryMapped, AccessSize, BitWidth);
  RegisterValue = Port->Read(Port, COM_LCR);
  RegisterValue &= ~LC_DLAB;
  Port->Write(Port, COM_LCR, RegisterValue);
  Port->Write(Port, COM_IEN, 0);
  Port->Write(Port, COM_FCR, FC_CLEAR_TRANSMIT | FC_CLEAR_RECEIVE);
  Uart16550SetBaud(Port, Port->BaudRate);
  Port->Write(Port, COM_FCR, FC_ENABLE);
  Port->Write(Port, COM_MCR, MC_DTRRTS);
  RegisterValue = Port->Read(Port, COM_MSR);
  if (CHECK_FLAG(RegisterValue, SERIAL_MSR_RI)) {
    Port->Flags |= PORT_RING_INDICATOR;
  }
  return TRUE;
}

BOOLEAN Uart16550LegacyInitializePort(
    _In_opt_ _Null_terminated_ PCHAR LoadOptions, _Inout_ PCPPORT Port,
    BOOLEAN MemoryMapped, UCHAR AccessSize, UCHAR BitWidth)

{
  UNREFERENCED_PARAMETER(AccessSize);
  UNREFERENCED_PARAMETER(BitWidth);
  UNREFERENCED_PARAMETER(MemoryMapped);
  switch ((ULONG_PTR)Port->Address) {
    case 1:
      Port->Address = (PUCHAR)COM1_PORT;
      break;
    case 2:
      Port->Address = (PUCHAR)COM2_PORT;
      break;
    case 3:
      Port->Address = (PUCHAR)COM3_PORT;
      break;
    case 4:
      Port->Address = (PUCHAR)COM4_PORT;
      break;
    default:
      return FALSE;
  }
  Port->Flags = 0;
  return Uart16550InitializePortCommon(LoadOptions, Port, FALSE,
                                       AcpiGenericAccessSizeByte, 8);
}

BOOLEAN Uart16550InitializePort(_In_opt_ _Null_terminated_ PCHAR LoadOptions,
                                _Inout_ PCPPORT Port, BOOLEAN MemoryMapped,
                                UCHAR AccessSize, UCHAR BitWidth)

{
  UCHAR RegisterBitWidth;
  UNREFERENCED_PARAMETER(AccessSize);
  UNREFERENCED_PARAMETER(BitWidth);
#if defined(_ARM_) || defined(_ARM64_)
  RegisterBitWidth = 32;
  if (MemoryMapped == FALSE) {
    return FALSE;
  }
  Port->Flags = PORT_DEFAULT_RATE;
#else
  RegisterBitWidth = 8;
  Port->Flags = 0;
#endif
  return Uart16550InitializePortCommon(LoadOptions, Port, MemoryMapped,
                                       AcpiGenericAccessSizeByte,
                                       RegisterBitWidth);
}

BOOLEAN Uart16550MmInitializePort(_In_opt_ _Null_terminated_ PCHAR LoadOptions,
                                  _Inout_ PCPPORT Port, BOOLEAN MemoryMapped,
                                  UCHAR AccessSize, UCHAR BitWidth)

{
  Port->Flags = PORT_DEFAULT_RATE;
  return Uart16550InitializePortCommon(LoadOptions, Port, MemoryMapped,
                                       AccessSize, BitWidth);
}

BOOLEAN Uart16550SetBaudCommon(_Inout_ PCPPORT Port, ULONG Rate, ULONG Clock)

{
  UCHAR Lcr;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return FALSE;
  }
  if ((Rate == 0) || (Clock == 0)) {
    return FALSE;
  }
  const ULONG DivisorLatch = Clock / Rate;
  Lcr = Port->Read(Port, COM_LCR);
  Lcr |= LC_DLAB;
  Port->Write(Port, COM_LCR, Lcr);
  Port->Write(Port, COM_DLM, (UCHAR)((DivisorLatch >> 8) & 0xFF));
  Port->Write(Port, COM_DLL, (UCHAR)(DivisorLatch & 0xFF));
  Port->Write(Port, COM_LCR, 3);
  Port->BaudRate = Rate;
  return TRUE;
}

BOOLEAN Uart16550SetBaud(_Inout_ PCPPORT Port, ULONG Rate)

{
  if (CHECK_FLAG(Port->Flags, PORT_DEFAULT_RATE)) {
    return FALSE;
  }
  return Uart16550SetBaudCommon(Port, Rate, CLOCK_RATE);
}

UART_STATUS Uart16550GetByte(_Inout_ PCPPORT Port, _Out_ PUCHAR Byte)

{
  UCHAR Data;
  UCHAR Lsr;
  UCHAR Msr;
  *Byte = 0;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return UartNotReady;
  }
  Lsr = Port->Read(Port, COM_LSR);
  if (Lsr == SERIAL_LSR_NOT_PRESENT) {
    return UartNotReady;
  }
  if (CHECK_FLAG(Lsr, COM_DATRDY)) {
    if (CHECK_FLAG(Lsr, COM_PE) || CHECK_FLAG(Lsr, COM_FE) ||
        CHECK_FLAG(Lsr, COM_OE)) {
      return UartError;
    }
    Data = Port->Read(Port, COM_DAT);
    if (CHECK_FLAG(Port->Flags, PORT_MODEM_CONTROL)) {
      Msr = Port->Read(Port, COM_MSR);
      if (CHECK_FLAG(Msr, MS_CD) == FALSE) {
        return UartNoData;
      }
    }
    *Byte = Data;
    return UartSuccess;
  } else {
    Msr = Port->Read(Port, COM_MSR);
    if ((CHECK_FLAG(Port->Flags, PORT_RING_INDICATOR) &&
         !CHECK_FLAG(Msr, SERIAL_MSR_RI)) ||
        (!CHECK_FLAG(Port->Flags, PORT_RING_INDICATOR) &&
         CHECK_FLAG(Msr, SERIAL_MSR_RI))) {
      Port->Flags |= PORT_MODEM_CONTROL;
    }
    return UartNoData;
  }
}

UART_STATUS Uart16550PutByte(_Inout_ PCPPORT Port, UCHAR Byte, BOOLEAN BusyWait)

{
  UCHAR Lsr;
  UCHAR Msr;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return UartNotReady;
  }
  if (CHECK_FLAG(Port->Flags, PORT_MODEM_CONTROL)) {
    Msr = Port->Read(Port, COM_MSR);
    while ((Msr & MS_DSRCTSCD) != MS_DSRCTSCD) {
      if (!CHECK_FLAG(Msr, MS_CD)) {
        Lsr = Port->Read(Port, COM_LSR);
        if (CHECK_FLAG(Port->Flags, COM_DATRDY)) {
          Port->Read(Port, COM_DAT);
        }
      }
      Msr = Port->Read(Port, COM_MSR);
    }
  }
  Lsr = Port->Read(Port, COM_LSR);
  if (Lsr == SERIAL_LSR_NOT_PRESENT) {
    return UartNotReady;
  }
  while (!CHECK_FLAG(Lsr, COM_OUTRDY)) {
    Msr = Port->Read(Port, COM_MSR);
    if ((CHECK_FLAG(Port->Flags, PORT_RING_INDICATOR) &&
         !CHECK_FLAG(Msr, SERIAL_MSR_RI)) ||
        (!CHECK_FLAG(Port->Flags, PORT_RING_INDICATOR) &&
         CHECK_FLAG(Msr, SERIAL_MSR_RI))) {
      Port->Flags |= PORT_MODEM_CONTROL;
    }
    if (BusyWait == FALSE) {
      return UartNotReady;
    }
    Lsr = Port->Read(Port, COM_LSR);
  }
  Port->Write(Port, COM_DAT, Byte);
  return UartSuccess;
}

BOOLEAN Uart16550RxReady(_Inout_ PCPPORT Port)

{
  UCHAR Lsr;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return FALSE;
  }
  Lsr = Port->Read(Port, COM_LSR);
  if (Lsr == SERIAL_LSR_NOT_PRESENT) {
    return FALSE;
  }
  if (CHECK_FLAG(Lsr, COM_DATRDY)) {
    return TRUE;
  }
  return FALSE;
}

UART_HARDWARE_DRIVER Legacy16550HardwareDriver = {
    Uart16550LegacyInitializePort, Uart16550SetBaud, Uart16550GetByte,
    Uart16550PutByte, Uart16550RxReady};
UART_HARDWARE_DRIVER Uart16550HardwareDriver = {
    Uart16550InitializePort, Uart16550SetBaud, Uart16550GetByte,
    Uart16550PutByte, Uart16550RxReady};
UART_HARDWARE_DRIVER MM16550HardwareDriver = {
    Uart16550MmInitializePort, Uart16550SetBaud, Uart16550GetByte,
    Uart16550PutByte, Uart16550RxReady};
