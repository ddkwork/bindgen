
#include "common.h"
#define USIF_FIFO_STAT 0x00000044
#define USIF_FIFO_STAT_TXFFS 0x00FF0000
#define USIF_FIFO_STAT_RXFFS 0x000000FF
#define USIF_TXD 0x00040000
#define USIF_RXD 0x00080000

BOOLEAN UsifInitializePort(_In_opt_ _Null_terminated_ PCHAR LoadOptions,
                           _Inout_ PCPPORT Port, BOOLEAN MemoryMapped,
                           UCHAR AccessSize, UCHAR BitWidth)

{
  UNREFERENCED_PARAMETER(LoadOptions);
  UNREFERENCED_PARAMETER(AccessSize);
  UNREFERENCED_PARAMETER(BitWidth);
  if (MemoryMapped == FALSE) {
    return FALSE;
  }
  Port->Flags = 0;
  return TRUE;
}

BOOLEAN UsifSetBaud(_Inout_ PCPPORT Port, ULONG Rate)

{
  if ((Port == NULL) || (Port->Address == NULL)) {
    return FALSE;
  }
  Port->BaudRate = Rate;
  return TRUE;
}

UART_STATUS UsifGetByte(_Inout_ PCPPORT Port, _Out_ PUCHAR Byte)

{
  ULONG Stat;
  ULONG Value;
  *Byte = 0;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return UartNotReady;
  }
  Stat = READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_FIFO_STAT));
  if ((Stat & USIF_FIFO_STAT_RXFFS) != 0) {
    Value = READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_RXD));
    *Byte = Value & (UCHAR)0xFF;
    return UartSuccess;
  }
  return UartNoData;
}

UART_STATUS UsifPutByte(_Inout_ PCPPORT Port, UCHAR Byte, BOOLEAN BusyWait)

{
  if ((Port == NULL) || (Port->Address == NULL)) {
    return UartNotReady;
  }
  if (BusyWait != FALSE) {
    while (READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_FIFO_STAT)) &
           (USIF_FIFO_STAT_TXFFS));
  } else if (READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_FIFO_STAT)) &
             (USIF_FIFO_STAT_TXFFS)) {
    return UartNotReady;
  }
  WRITE_REGISTER_ULONG((PULONG)(Port->Address + USIF_TXD), (ULONG)Byte);
  return UartSuccess;
}

BOOLEAN UsifRxReady(_Inout_ PCPPORT Port)

{
  ULONG Stat;
  if ((Port == NULL) || (Port->Address == NULL)) {
    return FALSE;
  }
  Stat = READ_REGISTER_ULONG((PULONG)(Port->Address + USIF_FIFO_STAT));
  if ((Stat & USIF_FIFO_STAT_RXFFS) != 0) {
    return TRUE;
  }
  return FALSE;
}
