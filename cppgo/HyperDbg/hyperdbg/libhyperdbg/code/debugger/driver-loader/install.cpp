
#include "pch.h"

BOOLEAN InstallDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName,
                      LPCTSTR ServiceExe) {
  SC_HANDLE SchService;
  DWORD LastError;
  SchService = CreateService(SchSCManager, DriverName, DriverName,
                             SERVICE_ALL_ACCESS, SERVICE_KERNEL_DRIVER,
                             SERVICE_DEMAND_START, SERVICE_ERROR_NORMAL,
                             ServiceExe, NULL, NULL, NULL, NULL, NULL);
  if (SchService == NULL) {
    LastError = GetLastError();
    if (LastError == ERROR_SERVICE_EXISTS) {
      ShowMessages("the service (driver) already exists\n");
      ShowMessages("trying to remove the old instance of the driver first\n");
      ManageDriver(DriverName, NULL, DRIVER_FUNC_STOP);
      if (ManageDriver(DriverName, NULL, DRIVER_FUNC_REMOVE)) {
        ShowMessages(
            "the old instance of the driver is removed successfully\n");
      } else {
        ShowMessages("err, failed to remove the old instance of the driver\n");
        return FALSE;
      }
      ShowMessages("installing the driver again\n");
      if (InstallDriver(SchSCManager, DriverName, ServiceExe)) {
        return TRUE;
      } else {
        ShowMessages(
            "err, failed to install the driver after removing the old "
            "instance\n");
        return FALSE;
      }
    } else if (LastError == ERROR_SERVICE_MARKED_FOR_DELETE) {
      ShowMessages(
          "err, previous instance of the service is not fully deleted. Try "
          "again...\n");
      return FALSE;
    } else {
      ShowMessages("err, CreateService failed (%x)\n", LastError);
      return FALSE;
    }
  }
  if (SchService) {
    CloseServiceHandle(SchService);
  }
  return TRUE;
}

BOOLEAN ManageDriver(LPCTSTR DriverName, LPCTSTR ServiceName, UINT16 Function) {
  SC_HANDLE SchSCManager;
  BOOLEAN Res = TRUE;
  if (!DriverName || (Function == DRIVER_FUNC_INSTALL && !ServiceName)) {
    ShowMessages("invalid Driver or Service provided to ManageDriver() \n");
    return FALSE;
  }
  SchSCManager = OpenSCManager(NULL, NULL, SC_MANAGER_ALL_ACCESS);
  if (!SchSCManager) {
    ShowMessages("err, OpenSCManager failed (%x)\n", GetLastError());
    return FALSE;
  }
  switch (Function) {
    case DRIVER_FUNC_INSTALL:
      if (InstallDriver(SchSCManager, DriverName, ServiceName)) {
        Res = StartDriver(SchSCManager, DriverName);
      } else {
        Res = FALSE;
      }
      break;
    case DRIVER_FUNC_STOP:
      Res = StopDriver(SchSCManager, DriverName);
      break;
    case DRIVER_FUNC_REMOVE:
      Res = RemoveDriver(SchSCManager, DriverName);
      break;
    default:
      ShowMessages("unknown ManageDriver() function \n");
      Res = FALSE;
      break;
  }
  if (SchSCManager) {
    CloseServiceHandle(SchSCManager);
  }
  return Res;
}

BOOLEAN RemoveDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName) {
  SC_HANDLE SchService;
  BOOLEAN Res;
  SchService = OpenService(SchSCManager, DriverName, SERVICE_ALL_ACCESS);
  if (SchService == NULL) {
    ShowMessages("err, OpenService failed (%x)\n", GetLastError());
    return FALSE;
  }
  if (DeleteService(SchService)) {
    Res = TRUE;
  } else {
    ShowMessages("err, DeleteService failed (%x)\n", GetLastError());
    Res = FALSE;
  }
  if (SchService) {
    CloseServiceHandle(SchService);
  }
  return Res;
}

BOOLEAN StartDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName) {
  SC_HANDLE SchService;
  DWORD LastError;
  BOOLEAN Status = TRUE;
  SchService = OpenService(SchSCManager, DriverName, SERVICE_ALL_ACCESS);
  if (SchService == NULL) {
    ShowMessages("err, OpenService failed (%x)\n", GetLastError());
    return FALSE;
  }
  if (!StartService(SchService, 0, NULL)) {
    LastError = GetLastError();
    if (LastError == ERROR_SERVICE_ALREADY_RUNNING) {
    } else if (LastError == ERROR_PATH_NOT_FOUND) {
      ShowMessages(
          "err, path to the driver not found, or the access to the driver file "
          "is limited\n");
      ShowMessages(
          "most of the time, it's because anti-virus software is not finished "
          "scanning the drivers, "
          "so, if you try to load the driver again (re-enter the previous "
          "command), the problem will be solved\n");
      Status = FALSE;
    } else if (LastError == ERROR_INVALID_IMAGE_HASH) {
      ShowMessages(
          "err, failed loading driver\n"
          "it's because either the driver signature enforcement is enabled or "
          "HVCI prevents the driver from loading\n"
          "you should disable the driver signature enforcement by attaching "
          "WinDbg or from the boot menu\n"
          "if the driver signature enforcement is disabled, HVCI might prevent "
          "the driver from loading\n"
          "HyperDbg is not compatible with Virtualization Based Security "
          "(VBS)\n"
          "please follow the instructions from: "
          "https://docs.hyperdbg.org/getting-started/build-and-install \n");
      Status = FALSE;
    } else {
      ShowMessages("err, StartService failure (%x)\n", LastError);
      Status = FALSE;
    }
  }
  if (SchService) {
    CloseServiceHandle(SchService);
  }
  return Status;
}

BOOLEAN StopDriver(SC_HANDLE SchSCManager, LPCTSTR DriverName) {
  BOOLEAN Res = TRUE;
  SC_HANDLE SchService;
  SERVICE_STATUS serviceStatus;
  SchService = OpenService(SchSCManager, DriverName, SERVICE_ALL_ACCESS);
  if (SchService == NULL) {
    ShowMessages("err, OpenService failed (%x)\n", GetLastError());
    return FALSE;
  }
  if (ControlService(SchService, SERVICE_CONTROL_STOP, &serviceStatus)) {
    Res = TRUE;
  } else {
    ShowMessages(
        "warning, failed to stop the driver. Possible reasons include the "
        "driver not currently running or an unsuccessful unload from a "
        "previous run. "
        "This is not an error, HyperDbg tries to remove the previous driver "
        "and load it again (%x)\n",
        GetLastError());
    Res = FALSE;
  }
  if (SchService) {
    CloseServiceHandle(SchService);
  }
  return Res;
}

BOOLEAN
SetupPathForFileName(const CHAR *FileName,
                     _Inout_updates_bytes_all_(BufferLength) PCHAR FileLocation,
                     ULONG BufferLength, BOOLEAN CheckFileExists) {
  HANDLE FileHandle;
  DWORD FileLocLen = 0;
  HMODULE ProcHandle = GetModuleHandle(NULL);
  char *Pos;

  GetModuleFileName(ProcHandle, FileLocation, BufferLength);
  Pos = strrchr(FileLocation, '\\');
  if (Pos != NULL) {
    *Pos = '\0';
  }
  if (FAILED(StringCbCat(FileLocation, BufferLength, "\\"))) {
    return FALSE;
  }
  if (FAILED(StringCbCat(FileLocation, BufferLength, FileName))) {
    return FALSE;
  }
  if (CheckFileExists) {
    if ((FileHandle = CreateFile(FileLocation, GENERIC_READ, 0, NULL,
                                 OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, NULL)) ==
        INVALID_HANDLE_VALUE) {
      ShowMessages("err, target file is not loaded\n");
      return FALSE;
    }
    if (FileHandle) {
      CloseHandle(FileHandle);
    }
  }
  return TRUE;
}
