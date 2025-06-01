
#pragma once

#define TEST_PROCESS_NAME "hyperdbg-test.exe"
BOOLEAN CreateProcessAndOpenPipeConnection(PHANDLE ConnectionPipeHandle,
                                           PHANDLE ThreadHandle,
                                           PHANDLE ProcessHandle);
BOOLEAN OpenHyperDbgTestProcess(PHANDLE ThreadHandle, PHANDLE ProcessHandle,
                                CHAR *Args);
VOID CloseProcessAndClosePipeConnection(HANDLE ConnectionPipeHandle,
                                        HANDLE ThreadHandle,
                                        HANDLE ProcessHandle);
