
#pragma once
#ifdef LIBHYPERDBG_EXPORTS
#define LIBHYPERDBG_API __declspec(dllexport)
#else
#define LIBHYPERDBG_API __declspec(dllimport)
#endif
#define DRIVER_FUNC_INSTALL 0x01
#define DRIVER_FUNC_STOP 0x02
#define DRIVER_FUNC_REMOVE 0x03
BOOLEAN InstallDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName,
                      LPCTSTR ServiceExe);
BOOLEAN RemoveDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName);
BOOLEAN StartDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName);
BOOLEAN StopDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName);
BOOLEAN ManageDriver(_In_ LPCTSTR DriverName, _In_ LPCTSTR ServiceName,
                     _In_ UINT16 Function);
BOOLEAN
SetupPathForFileName(const CHAR *FileName,
                     _Inout_updates_bytes_all_(BufferLength) PCHAR FileLocation,
                     ULONG BufferLength, BOOLEAN CheckFileExists);
