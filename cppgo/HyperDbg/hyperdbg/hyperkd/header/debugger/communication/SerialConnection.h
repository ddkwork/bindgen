
#pragma once
VOID KdHyperDbgTest(UINT16 Byte);
VOID KdHyperDbgPrepareDebuggeeConnectionPort(UINT32 PortAddress,
                                             UINT32 Baudrate);
VOID KdHyperDbgSendByte(UCHAR Byte, BOOLEAN BusyWait);
BOOLEAN KdHyperDbgRecvByte(PUCHAR RecvByte);
VOID SerialConnectionTest();
NTSTATUS SerialConnectionPrepare(PDEBUGGER_PREPARE_DEBUGGEE DebuggeeRequest);
BOOLEAN SerialConnectionCheckPort(UINT32 SerialPort);
BOOLEAN SerialConnectionCheckBaudrate(DWORD Baudrate);
BOOLEAN SerialConnectionSend(CHAR *Buffer, UINT32 Length);
BOOLEAN SerialConnectionRecvBuffer(CHAR *BufferToSave, UINT32 *LengthReceived);
BOOLEAN SerialConnectionSendTwoBuffers(CHAR *Buffer1, UINT32 Length1,
                                       CHAR *Buffer2, UINT32 Length2);
BOOLEAN SerialConnectionSendThreeBuffers(CHAR *Buffer1, UINT32 Length1,
                                         CHAR *Buffer2, UINT32 Length2,
                                         CHAR *Buffer3, UINT32 Length3);
#define CBR_110 110
#define CBR_300 300
#define CBR_600 600
#define CBR_1200 1200
#define CBR_2400 2400
#define CBR_4800 4800
#define CBR_9600 9600
#define CBR_14400 14400
#define CBR_19200 19200
#define CBR_38400 38400
#define CBR_56000 56000
#define CBR_57600 57600
#define CBR_115200 115200
#define CBR_128000 128000
#define CBR_256000 256000
#define COM1_PORT 0x03F8
#define COM2_PORT 0x02F8
#define COM3_PORT 0x03E8
#define COM4_PORT 0x02E8
