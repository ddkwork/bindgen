
#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

BOOLEAN ObjectShowProcessesOrThreadDetails(BOOLEAN IsProcess) {
  BOOLEAN Status;
  ULONG ReturnedLength;
  DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET GetInformationProcess = {0};
  DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET GetInformationThread = {0};
  if (IsProcess) {
    Status = DeviceIoControl(g_DeviceHandle, IOCTL_QUERY_CURRENT_PROCESS,
                             &GetInformationProcess,
                             SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET,
                             &GetInformationProcess,
                             SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET,
                             &ReturnedLength, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      return FALSE;
    }
    if (GetInformationProcess.Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
      ShowMessages(
          "process id: %x\nprocess (_EPROCESS): %s\nprocess name (16-Byte): "
          "%s\n",
          GetInformationProcess.ProcessId,
          SeparateTo64BitValue(GetInformationProcess.Process).c_str(),
          GetInformationProcess.ProcessName);
      return TRUE;
    } else {
      ShowErrorMessage(GetInformationProcess.Result);
      return FALSE;
    }
  } else {
    Status = DeviceIoControl(
        g_DeviceHandle, IOCTL_QUERY_CURRENT_THREAD, &GetInformationThread,
        SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET, &GetInformationThread,
        SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET, &ReturnedLength,
        NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      return FALSE;
    }
    if (GetInformationThread.Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
      ShowMessages(
          "thread id: %x (pid: %x)\nthread (_ETHREAD): %s\nprocess "
          "(_EPROCESS): %s\nprocess name (16-Byte): %s\n",
          GetInformationThread.ThreadId, GetInformationThread.ProcessId,
          SeparateTo64BitValue(GetInformationThread.Thread).c_str(),
          SeparateTo64BitValue(GetInformationThread.Process).c_str(),
          GetInformationThread.ProcessName);
      return TRUE;
    } else {
      ShowErrorMessage(GetInformationThread.Result);
      return FALSE;
    }
  }
}

BOOLEAN ObjectShowProcessesOrThreadList(
    BOOLEAN IsProcess,
    PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS SymDetailsForProcessList,
    UINT64 Eprocess,
    PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS SymDetailsForThreadList) {
  BOOLEAN Status;
  ULONG ReturnedLength;
  DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
  QueryCountOfActiveThreadsOrProcessesRequest = {0};
  UINT32 SizeOfBufferForThreadsAndProcessDetails = NULL;
  PVOID Entries = NULL;
  PDEBUGGEE_PROCESS_LIST_DETAILS_ENTRY ProcessEntries = NULL;
  PDEBUGGEE_THREAD_LIST_DETAILS_ENTRY ThreadEntries = NULL;
  AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED,
                              AssertReturnFalse);
  if (IsProcess) {
    QueryCountOfActiveThreadsOrProcessesRequest.QueryType =
        DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_COUNT;
  } else {
    QueryCountOfActiveThreadsOrProcessesRequest.QueryType =
        DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_COUNT;
  }
  QueryCountOfActiveThreadsOrProcessesRequest.QueryAction =
      DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT;
  if (IsProcess) {
    RtlCopyMemory(
        &QueryCountOfActiveThreadsOrProcessesRequest.ProcessListNeededDetails,
        SymDetailsForProcessList, sizeof(DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS));
  } else {
    RtlCopyMemory(
        &QueryCountOfActiveThreadsOrProcessesRequest.ThreadListNeededDetails,
        SymDetailsForThreadList, sizeof(DEBUGGEE_THREAD_LIST_NEEDED_DETAILS));
  }
  Status = DeviceIoControl(
      g_DeviceHandle, IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS,
      &QueryCountOfActiveThreadsOrProcessesRequest,
      SIZEOF_DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS,
      &QueryCountOfActiveThreadsOrProcessesRequest,
      SIZEOF_DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS, &ReturnedLength, NULL);
  if (!Status) {
    ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
    return FALSE;
  }
  if (QueryCountOfActiveThreadsOrProcessesRequest.Result ==
      DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    if (QueryCountOfActiveThreadsOrProcessesRequest.Count == 0) {
      ShowMessages("err, unable to get count of active processes or threads\n");
    } else {
      QueryCountOfActiveThreadsOrProcessesRequest.Count =
          QueryCountOfActiveThreadsOrProcessesRequest.Count + 5;
      if (IsProcess) {
        SizeOfBufferForThreadsAndProcessDetails =
            QueryCountOfActiveThreadsOrProcessesRequest.Count *
            sizeof(DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY);
      } else {
        SizeOfBufferForThreadsAndProcessDetails =
            QueryCountOfActiveThreadsOrProcessesRequest.Count *
            sizeof(DEBUGGEE_THREAD_LIST_DETAILS_ENTRY);
      }
      Entries = (PVOID)malloc(SizeOfBufferForThreadsAndProcessDetails);
      RtlZeroMemory(Entries, SizeOfBufferForThreadsAndProcessDetails);
      if (IsProcess) {
        QueryCountOfActiveThreadsOrProcessesRequest.QueryType =
            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_LIST;
      } else {
        QueryCountOfActiveThreadsOrProcessesRequest.QueryType =
            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_LIST;
      }
      Status = DeviceIoControl(
          g_DeviceHandle, IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES,
          &QueryCountOfActiveThreadsOrProcessesRequest,
          SIZEOF_DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS, Entries,
          SizeOfBufferForThreadsAndProcessDetails, &ReturnedLength, NULL);
      if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
      }
      if (IsProcess) {
        ProcessEntries = (PDEBUGGEE_PROCESS_LIST_DETAILS_ENTRY)Entries;
      } else {
        ThreadEntries = (PDEBUGGEE_THREAD_LIST_DETAILS_ENTRY)Entries;
        ShowMessages("PROCESS\t%llx\tIMAGE\t%s\n", ThreadEntries->Eprocess,
                     ThreadEntries->ImageFileName);
      }
      for (size_t i = 0; i < QueryCountOfActiveThreadsOrProcessesRequest.Count;
           i++) {
        if (IsProcess) {
          if (ProcessEntries[i].Eprocess != NULL) {
            ShowMessages(
                "PROCESS\t%llx\n\tProcess Id: %04x\tDirBase (Kernel Cr3): "
                "%016llx\tImage: %s\n\n",
                ProcessEntries[i].Eprocess, ProcessEntries[i].ProcessId,
                ProcessEntries[i].Cr3, ProcessEntries[i].ImageFileName);
          }
        } else {
          if (ThreadEntries[i].Ethread != NULL) {
            ShowMessages("\tTHREAD\t%llx (%x.%x)\n", ThreadEntries[i].Ethread,
                         ThreadEntries[i].ProcessId, ThreadEntries[i].ThreadId);
          }
        }
      }
    }
    return TRUE;
  } else {
    ShowErrorMessage(
        (UINT32)QueryCountOfActiveThreadsOrProcessesRequest.Result);
    return FALSE;
  }
  return FALSE;
}
