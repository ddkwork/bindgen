
#include "pch.h"
extern UINT64 g_EventTag;
extern UINT64 g_KernelBaseAddress;
extern LIST_ENTRY g_EventTrace;
extern BOOLEAN g_EventTraceInitialized;
extern BOOLEAN g_BreakPrintingOutput;
extern BOOLEAN g_AutoUnpause;
extern BOOLEAN g_OutputSourcesInitialized;
extern LIST_ENTRY g_OutputSources;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern BOOLEAN g_IsConnectedToRemoteDebugger;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

BOOLEAN ShowErrorMessage(UINT32 Error) {
  switch (Error) {
    case DEBUGGER_ERROR_TAG_NOT_EXISTS:
      ShowMessages("err, tag not found (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_INVALID_ACTION_TYPE:
      ShowMessages("err, invalid action type (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO:
      ShowMessages("err, action buffer size is zero (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID:
      ShowMessages("err, the event type is invalid (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT:
      ShowMessages("err, unable to create event (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_INVALID_ADDRESS:
      ShowMessages(
          "err, invalid address (%x)\n"
          "address may be paged-out or unavailable on the page table due to "
          "'demand paging'\n"
          "please refer to "
          "https://docs.hyperdbg.org/tips-and-tricks/considerations/"
          "accessing-invalid-address for further information\n",
          Error);
      break;
    case DEBUGGER_ERROR_INVALID_CORE_ID:
      ShowMessages("err, invalid core id (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES:
      ShowMessages("err, exception index exceed first 32 entries (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID:
      ShowMessages("err, interrupt index is not valid (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER:
      ShowMessages(
          "err, unable to hide or unhide debugger and entr the transparent "
          "mode (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_DEBUGGER_ALREADY_HIDE:
      ShowMessages(
          "err, debugger already hidden in the transparent mode (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER:
      ShowMessages("err, edit memory request has invalid parameters (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS:
      ShowMessages(
          "err, edit memory request has invalid address based on "
          "current process layout, the address might be valid but not "
          "present in the ram (%x)\n"
          "address may be paged-out or unavailable on the page table due to "
          "'demand paging'\n"
          "please refer to "
          "https://docs.hyperdbg.org/tips-and-tricks/considerations/"
          "accessing-invalid-address for further information\n",
          Error);
      break;
    case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS:
      ShowMessages(
          "err, edit memory request has invalid address based on other "
          "process layout (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG:
      ShowMessages("err, modify event with invalid tag (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION:
      ShowMessages("err, modify event with invalid type of action (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_STEPPING_INVALID_PARAMETER:
      ShowMessages("err, invalid parameter passed to stepping core. (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED:
      ShowMessages(
          "err, the target thread not found or the thread is disabled (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE:
      ShowMessages("err, invalid baud rate (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT:
      ShowMessages("err, invalid serial port (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE:
      ShowMessages("err, invalid core selected in switching cores (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS:
      ShowMessages("err, unable to switch to the new process (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT:
      ShowMessages("err, unable to run the script on remote debuggee (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_INVALID_REGISTER_NUMBER:
      ShowMessages("err, invalid register number (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE:
      ShowMessages(
          "err, maximum number of breakpoints are used, you need to "
          "send an ioctl to re-allocate new pre-allocated buffers or "
          "configure HyperDbg to pre-allocated more buffers by "
          "configuring MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS:
      ShowMessages(
          "err, breakpoint already exists on target address, you can "
          "use 'bl' command to view a list of breakpoints (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND:
      ShowMessages("err, breakpoint id is invalid (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED:
      ShowMessages("err, breakpoint already disabled (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED:
      ShowMessages("err, breakpoint already enabled (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_MEMORY_TYPE_INVALID:
      ShowMessages("err, the memory type is invalid (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_INVALID_PROCESS_ID:
      ShowMessages(
          "err, the process id is invalid, make sure to enter the "
          "process id in hex format, or if you want to use it in decimal "
          "format, add '0n' prefix to the number (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED:
      ShowMessages("err, the event is not applied (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER:
      ShowMessages(
          "err, either the process id or the _EPROCESS is invalid or "
          "cannot get the details based on the provided parameters (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER:
      ShowMessages(
          "err, either the thread id, _ETHREAD, or _EPROCESS is invalid or "
          "cannot get the details based on the provided parameters (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT:
      ShowMessages(
          "err, the maximum breakpoint for a single page is hit, "
          "you cannot apply more breakpoints in this page (%x)\n"
          "please refer to "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/customize-build/"
          "number-of-ept-hooks-in-one-page"
          " for further information\n",
          Error);
      break;
    case DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY:
      ShowMessages(
          "err, the pre-allocated buffer is empty, usually this buffer will be "
          "filled at the next IOCTL when the debugger is continued (%x)\n"
          "please visit the documentation for the 'prealloc' command or use "
          "'.help prealloc' to to reserve more pre-allocated pools\n",
          Error);
      break;
    case DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES:
      ShowMessages("err, could not convert 2MB large page to 4KB pages (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS:
      ShowMessages(
          "err, failed to get the PML1 entry of the target address (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE:
      ShowMessages(
          "err, the page modification is not applied, make sure that you don't "
          "put multiple EPT Hooks or Monitors on a single page (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK:
      ShowMessages("err, could not build the EPT hook (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE:
      ShowMessages("err, could not find the allocation type (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_INVALID_TEST_QUERY_INDEX:
      ShowMessages("err, invalid index specified for test query command (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS:
      ShowMessages("err, unable to attach to the target process (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED:
      ShowMessages(
          "err, unable to remove hooks as the entrypoint of user-mode "
          "process is not reached yet (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS:
      ShowMessages("err, unable to remove hooks (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED:
      ShowMessages(
          "err, the routines for getting the PEB is not correctly "
          "initialized (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS:
      ShowMessages(
          "err, unable to detect whether the process was 32-bit "
          "or 64-bit (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS:
      ShowMessages("err, unable to kill the process (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN:
      ShowMessages("err, invalid thread debugging token (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS:
      ShowMessages("err, unable to pause the threads of the process (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS:
      ShowMessages(
          "err, the user debugger is already attached to this "
          "process, please use the '.switch' command to switch "
          "to this process (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS:
      ShowMessages(
          "err, the user debugger is not already attached to "
          "the process (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS:
      ShowMessages(
          "err, the user debugger is not able to detach from "
          "this process as there are paused threads in the "
          "target process, please make sure to remove all "
          "the events and continue the target process, then "
          "perform the detach again (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID:
      ShowMessages(
          "err, unable to switch to the process id or thread id "
          "as the target process id or thread id is not found in "
          "the attached threads list, please view the list of "
          "processes and threads by using the '.switch list' command (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS:
      ShowMessages(
          "err, unable to switch to the process as the process doesn't "
          "contain an active intercepted thread (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS:
      ShowMessages("err, unable to get user-mode modules of the process (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK:
      ShowMessages("err, unable to get the callstack (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS:
      ShowMessages("err, unable to query count of processes or threads (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_USING_SHORT_CIRCUITING_EVENT_WITH_POST_EVENT_MODE_IS_FORBIDDEDN:
      ShowMessages(
          "err, using short-circuiting event with post events is not possible "
          "(%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNKNOWN_TEST_QUERY_RECEIVED:
      ShowMessages("err, unknown test query is received to the debugger (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER:
      ShowMessages("err, invalid process or memory address (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_THE_TRAP_FLAG_LIST_IS_FULL:
      ShowMessages(
          "err, unable to add the current thread/process to the list of trap "
          "flags. "
          "Are you debugging multiple threads or stepping through different "
          "processes "
          "simultaneously? (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS_DOES_NOT_EXISTS:
      ShowMessages("err, process does not exists (already terminated?) (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_MODE_EXECUTION_IS_INVALID:
      ShowMessages("err, the specified execution mode is invalid (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_PROCESS_ID_CANNOT_BE_SPECIFIED_WHILE_APPLYING_EVENT_FROM_VMX_ROOT_MODE:
      ShowMessages(
          "err, you cannot specify process id while the debugger is paused in "
          "the debugger mode. "
          "You can use the '.process' or the '.thread' command to switch to "
          "the target process's "
          "memory layout (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_EVENT_AND_CONDITIONALS:
      ShowMessages(
          "err, the requested buffer for storing event and conditions is "
          "larger than the pre-allocated "
          "buffer size (%x)\nfor more information on how to resolve this "
          "issue, "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND:
      ShowMessages(
          "err, not enough pre-allocated buffer exists for storing the event. "
          "You can use the 'prealloc' "
          "command to fix this issue by pre-allocating more buffers (%x)\nfor "
          "more information "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_BIG_PREALLOCATED_BUFFER_NOT_FOUND:
      ShowMessages(
          "err, the requested event is considered as a \"big instant event\" "
          "and right now, there is no "
          "pre-allocated buffer for storing it. You can use the 'prealloc' "
          "command to fix this issue by "
          "pre-allocating big instant event buffers (%x)\nfor more information "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_CREATE_ACTION_CANNOT_ALLOCATE_BUFFER:
      ShowMessages(
          "err, unable to allocate buffer for the target action (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_ACTION_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND:
      ShowMessages(
          "err, not enough pre-allocated buffer exists for storing the event's "
          "action. You can use the 'prealloc' "
          "command to fix this issue by pre-allocating more buffers (%x)\nfor "
          "more information "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_ACTION_BIG_PREALLOCATED_BUFFER_NOT_FOUND:
      ShowMessages(
          "err, the requested action is considered as a \"big instant event "
          "(action)\" and right now, there is no "
          "pre-allocated buffer for storing it. You can use the 'prealloc' "
          "command to fix this issue by "
          "pre-allocating big instant event's action buffers (%x)\nfor more "
          "information "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_ACTION_BUFFER:
      ShowMessages(
          "err, the requested buffer for storing action is larger than the "
          "pre-allocated "
          "buffer size (%x)\nfor more information on how to resolve this "
          "issue, "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_REQUESTED_OPTIONAL_BUFFER_IS_BIGGER_THAN_DEBUGGERS_SEND_RECEIVE_STACK:
      ShowMessages(
          "err, the requested optional buffer is bigger than the debuggers "
          "send/receive stack, "
          "please select a smaller requested buffer for the target event "
          "(%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_REQUESTED_SAFE_BUFFER_NOT_FOUND:
      ShowMessages(
          "err, by default HyperDbg won't allocate requested safe buffers for "
          "instant events in "
          "the debugger mode. You can use the 'prealloc' command to allocate "
          "(regular) requested safe "
          "buffers before running this command (%x)\nfor more information "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_BIG_REQUESTED_SAFE_BUFFER_NOT_FOUND:
      ShowMessages(
          "err, the requested safe buffer is bigger than regular buffers. You "
          "can use the 'prealloc' "
          "command to allocate (big) requested safe buffers before running "
          "this command (%x)\n"
          "for more information "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_REQUESTED_SAFE_BUFFER:
      ShowMessages(
          "err, the requested buffer for storing safe buffers of the action is "
          "larger than the pre-allocated "
          "buffer size (%x)\nfor more information on how to resolve this "
          "issue, "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_UNABLE_TO_ALLOCATE_REQUESTED_SAFE_BUFFER:
      ShowMessages(
          "err, unable to allocate buffer for the target requested safe buffer "
          "(%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_COULD_NOT_FIND_PREACTIVATION_TYPE:
      ShowMessages("err, invalid type is specified for preactivation (%x)\n",
                   Error);
      break;
    case DEBUGGER_ERROR_THE_MODE_EXEC_TRAP_IS_NOT_INITIALIZED:
      ShowMessages(
          "err, the '!mode' event command cannot be directly initialized in "
          "the Debugger Mode. "
          "To avoid wasting system resources and performance issues we decided "
          "to use another "
          "command to initialize it first then use it. You can use the "
          "'preactivate mode' command "
          "to preactivate this mechanism after that, you can use the '!mode' "
          "event (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_THE_TARGET_EVENT_IS_DISABLED_BUT_CANNOT_BE_CLEARED_PRIRITY_BUFFER_IS_FULL:
      ShowMessages(
          "err, the event(s) that you've requested are disabled, yet HyperDbg "
          "cannot remove (clear) them in "
          "the subsequent run due to the user-mode priority buffers being at "
          "full capacity. "
          "This typically occurs when you attempt to clear numerous events "
          "without resuming the debuggee. "
          "Since these unserviced events remain in the queue, HyperDbg is "
          "unable to clear them. "
          "To address this issue, you can resume the debuggee, allowing all "
          "queued events to be cleared (usually 2 to 10 seconds). "
          "Afterward, you can pause the debuggee again and request the removal "
          "of new events (%x)\n"
          "for more information on how to resolve this issue "
          "please visit: "
          "https://docs.hyperdbg.org/tips-and-tricks/misc/instant-events\n",
          Error);
      break;
    case DEBUGGER_ERROR_NOT_ALL_CORES_ARE_LOCKED_FOR_APPLYING_INSTANT_EVENT:
      ShowMessages(
          "err, the event cannot be applied since not all cores are locked! "
          "If you are using the instant-event mechanism or switching between "
          "cores, this will likely halt the system. "
          "This may be caused by a race condition, but continuing and halting "
          "the debugger might resolve it. "
          "If the problem persists, please open an issue and provide steps to "
          "reproduce it (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_TARGET_SWITCHING_CORE_IS_NOT_LOCKED:
      ShowMessages(
          "err, target core is not locked! "
          "This may be caused by a race condition, continuing and halting the "
          "debugger might resolve it. "
          "If the problem persists, please open an issue and provide steps to "
          "reproduce it (%x)\n",
          Error);
      break;
    case DEBUGGER_ERROR_INVALID_PHYSICAL_ADDRESS:
      ShowMessages("err, invalid physical address (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_APIC_ACTIONS_ERROR:
      ShowMessages("err, could not perform APIC actions (%x)\n", Error);
      break;
    case DEBUGGER_ERROR_DEBUGGER_ALREADY_UNHIDE:
      ShowMessages(
          "err, debugger was not in the hidden transparent-mode (%x)\n", Error);
      break;
    default:
      ShowMessages("err, error not found (%x)\n", Error);
      return FALSE;
      break;
  }
  return TRUE;
}

UINT64 DebuggerGetNtoskrnlBase() {
  UINT64 NtoskrnlBase = NULL;
  PRTL_PROCESS_MODULES Modules = NULL;
  if (SymbolCheckAndAllocateModuleInformation(&Modules) == FALSE) {
    ShowMessages(
        "err, unable to get the module list for getting ntoskrnl base "
        "address\n");
    return NULL64_ZERO;
  }
  for (UINT32 i = 0; i < Modules->NumberOfModules; i++) {
    if (!strcmp((const char *)Modules->Modules[i].FullPathName +
                    Modules->Modules[i].OffsetToFileName,
                "ntoskrnl.exe")) {
      NtoskrnlBase = (UINT64)Modules->Modules[i].ImageBase;
      break;
    }
  }
  free(Modules);
  return NtoskrnlBase;
}

UINT64 DebuggerGetKernelBase() {
  UINT64 KernelBase = NULL;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    KernelBase = g_KernelBaseAddress;
  } else {
    KernelBase = DebuggerGetNtoskrnlBase();
    g_KernelBaseAddress = KernelBase;
  }
  return KernelBase;
}

BOOLEAN DebuggerPauseDebuggee() {
  BOOLEAN StatusIoctl = 0;
  ULONG ReturnedLength = 0;
  DEBUGGER_PAUSE_PACKET_RECEIVED PauseRequest = {0};
  StatusIoctl = DeviceIoControl(
      g_DeviceHandle, IOCTL_PAUSE_PACKET_RECEIVED, &PauseRequest,
      SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED, &PauseRequest,
      SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED, &ReturnedLength, NULL);
  if (!StatusIoctl) {
    ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
    return FALSE;
  }
  if (PauseRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
    return TRUE;
  } else {
    ShowErrorMessage(PauseRequest.Result);
    return FALSE;
  }
  return FALSE;
}

BOOLEAN IsConnectedToAnyInstanceOfDebuggerOrDebuggee() {
  if (g_DeviceHandle) {
    ShowMessages(
        "err, the current system is already connected to the local "
        "debugging, use '.disconnect' to disconnect\n");
    return TRUE;
  } else if (g_IsConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, the current system is already connected to remote "
        "machine (debuggee), use '.disconnect' to disconnect from the "
        "remote machine\n");
    return TRUE;
  } else if (g_IsConnectedToRemoteDebugger) {
    ShowMessages(
        "err, the current system is already connected to remote "
        "machine (debugger), use '.disconnect' to disconnect from the "
        "remote machine from debugger\n");
    return TRUE;
  } else if (g_IsSerialConnectedToRemoteDebuggee) {
    ShowMessages(
        "err, the current system is already connected to remote "
        "machine (debuggee), use '.debug close' to disconnect from the "
        "remote machine\n");
    return TRUE;
  } else if (g_IsSerialConnectedToRemoteDebugger) {
    ShowMessages(
        "err, the current system is already connected to remote "
        "machine (debugger), use '.debug close' to disconnect from the "
        "remote machine from debugger\n");
    return TRUE;
  }
  return FALSE;
}

BOOLEAN IsTagExist(UINT64 Tag) {
  PLIST_ENTRY TempList = 0;
  PDEBUGGER_GENERAL_EVENT_DETAIL CommandDetail = {0};
  if (!g_EventTraceInitialized) {
    return FALSE;
  }
  TempList = &g_EventTrace;
  while (&g_EventTrace != TempList->Blink) {
    TempList = TempList->Blink;
    CommandDetail = CONTAINING_RECORD(TempList, DEBUGGER_GENERAL_EVENT_DETAIL,
                                      CommandsEventList);
    if (CommandDetail->Tag == Tag ||
        Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG) {
      return TRUE;
    }
  }
  return FALSE;
}

BOOLEAN InterpretScript(vector<CommandToken> *CommandTokens,
                        PBOOLEAN ScriptSyntaxErrors, PUINT64 BufferAddress,
                        PUINT32 BufferLength, PUINT32 Pointer,
                        PUINT64 ScriptCodeBuffer) {
  BOOLEAN IsTextVisited = FALSE;
  string TargetBracketString = "";
  vector<int> IndexesToRemove;
  UINT32 Index = 0;
  UINT32 NewIndexToRemove = 0;
  *ScriptSyntaxErrors = TRUE;
  for (auto Section : *CommandTokens) {
    Index++;
    if (IsTextVisited && IsTokenBracketString(Section)) {
      IndexesToRemove.push_back(Index);
      TargetBracketString = GetCaseSensitiveStringFromCommandToken(Section);
      IsTextVisited = FALSE;
      continue;
    }
    if (CompareLowerCaseStrings(Section, "script")) {
      IndexesToRemove.push_back(Index);
      IsTextVisited = TRUE;
      continue;
    }
  }
  if (TargetBracketString.length() == 0) {
    return FALSE;
  }
  if (TargetBracketString.rfind("file:", 0) == 0) {
    std::ifstream t(TargetBracketString.erase(0, 5).c_str());
    std::stringstream buffer;
    buffer << t.rdbuf();
    TargetBracketString = buffer.str();
    if (TargetBracketString.empty()) {
      ShowMessages("err, either script file is not found or it's empty\n");
      *ScriptSyntaxErrors = TRUE;
      return TRUE;
    }
  }
  PVOID CodeBuffer =
      ScriptEngineParseWrapper((char *)TargetBracketString.c_str(), TRUE);
  if (CodeBuffer == NULL) {
    *ScriptSyntaxErrors = TRUE;
    return TRUE;
  } else {
    *ScriptSyntaxErrors = FALSE;
  }
  *BufferAddress = ScriptEngineWrapperGetHead(CodeBuffer);
  *BufferLength = ScriptEngineWrapperGetSize(CodeBuffer);
  *Pointer = ScriptEngineWrapperGetPointer(CodeBuffer);
  *ScriptCodeBuffer = (UINT64)CodeBuffer;
  NewIndexToRemove = 0;
  for (auto IndexToRemove : IndexesToRemove) {
    NewIndexToRemove++;
    CommandTokens->erase(CommandTokens->begin() +
                         (IndexToRemove - NewIndexToRemove));
  }
  return TRUE;
}

BOOLEAN InterpretConditionsAndCodes(vector<CommandToken> *CommandTokens,
                                    BOOLEAN IsConditionBuffer,
                                    PUINT64 BufferAddress,
                                    PUINT32 BufferLength) {
  BOOLEAN IsAsm = FALSE;
  BOOLEAN IsTextVisited = FALSE;
  string TargetBracketString = "";
  string Temp;
  vector<CHAR> ParsedBytes;
  vector<int> IndexesToRemove;
  UCHAR *FinalBuffer;
  UINT32 AssembledByteCount;
  int NewIndexToRemove = 0;
  int Index = 0;
  for (auto Section : *CommandTokens) {
    Index++;
    if (IsTextVisited && IsTokenBracketString(Section)) {
      IndexesToRemove.push_back(Index);
      TargetBracketString = GetCaseSensitiveStringFromCommandToken(Section);
      IsTextVisited = FALSE;
      continue;
    }
    if (!IsConditionBuffer && CompareLowerCaseStrings(Section, "code")) {
      IndexesToRemove.push_back(Index);
      IsTextVisited = TRUE;
      continue;
    }
    if (IsConditionBuffer && CompareLowerCaseStrings(Section, "condition")) {
      IndexesToRemove.push_back(Index);
      IsTextVisited = TRUE;
      continue;
    }
    if (CompareLowerCaseStrings(Section, "asm")) {
      IndexesToRemove.push_back(Index);
      IsAsm = TRUE;
      continue;
    }
  }
  if (TargetBracketString.length() == 0) {
    return FALSE;
  }
  if (IsAsm) {
    AssembleData AssembleData;
    AssembleData.AsmRaw = TargetBracketString;
    AssembleData.ParseAssemblyData();
    AssembleData.AsmFixed += "; ret";
    if (AssembleData.Assemble(0)) {
      ShowMessages("err, assemble error code: '%u'\n\n", AssembleData.KsErr);
      return FALSE;
    } else {
      AssembledByteCount = (UINT32)AssembleData.BytesCount;
      FinalBuffer = (unsigned char *)malloc(AssembledByteCount);
      if (FinalBuffer == NULL) {
        return FALSE;
      }
      memcpy(FinalBuffer, AssembleData.EncodedBytes, AssembledByteCount);
      *BufferAddress = (UINT64)FinalBuffer;
      *BufferLength = AssembledByteCount;
    }
  } else {
    TargetBracketString += "c3";
    if (TargetBracketString.rfind("0x", 0) == 0 ||
        TargetBracketString.rfind("0X", 0) == 0 ||
        TargetBracketString.rfind("\\x", 0) == 0 ||
        TargetBracketString.rfind("\\X", 0) == 0) {
      Temp = TargetBracketString.erase(0, 2);
    } else if (TargetBracketString.rfind('x', 0) == 0 ||
               TargetBracketString.rfind('X', 0) == 0) {
      Temp = TargetBracketString.erase(0, 1);
    } else {
      Temp = std::move(TargetBracketString);
    }
    ReplaceAll(Temp, "\\x", "");
    if (Temp.size() % 2 != 0) {
      Temp.insert(0, 1, '0');
    }
    if (!IsHexNotation(Temp)) {
      ShowMessages("please enter condition code in a hex notation\n");
      return FALSE;
    }
    ParsedBytes = HexToBytes(Temp);
    FinalBuffer = (UCHAR *)malloc(ParsedBytes.size());
    if (FinalBuffer == NULL) {
      return FALSE;
    }
    std::copy(ParsedBytes.begin(), ParsedBytes.end(), FinalBuffer);
    *BufferAddress = (UINT64)FinalBuffer;
    *BufferLength = (UINT32)ParsedBytes.size();
  }
  NewIndexToRemove = 0;
  for (auto IndexToRemove : IndexesToRemove) {
    NewIndexToRemove++;
    CommandTokens->erase(CommandTokens->begin() +
                         (IndexToRemove - NewIndexToRemove));
  }
  return TRUE;
}

BOOLEAN InterpretOutput(vector<CommandToken> *CommandTokens,
                        vector<string> &InputSources) {
  BOOLEAN IsTextVisited = FALSE;
  string TargetBracketString = "";
  vector<int> IndexesToRemove;
  string Token;
  int NewIndexToRemove = 0;
  int Index = 0;
  char Delimiter = ',';
  size_t Pos = 0;
  for (auto Section : *CommandTokens) {
    Index++;
    if (IsTextVisited && IsTokenBracketString(Section)) {
      IndexesToRemove.push_back(Index);
      TargetBracketString = GetCaseSensitiveStringFromCommandToken(Section);
      IsTextVisited = FALSE;
      continue;
    }
    if (CompareLowerCaseStrings(Section, "output")) {
      IndexesToRemove.push_back(Index);
      IsTextVisited = TRUE;
      continue;
    }
  }
  if (TargetBracketString.length() == 0) {
    return FALSE;
  }
  if (TargetBracketString.find(Delimiter) != std::string::npos) {
    while ((Pos = TargetBracketString.find(Delimiter)) != string::npos) {
      Token = TargetBracketString.substr(0, Pos);
      Trim(Token);
      if (!Token.empty()) {
        InputSources.push_back(Token);
      }
      TargetBracketString.erase(0, Pos + sizeof(Delimiter) / sizeof(char));
    }
    if (!TargetBracketString.empty()) {
      InputSources.push_back(TargetBracketString);
    }
  } else {
    InputSources.push_back(TargetBracketString);
  }
  NewIndexToRemove = 0;
  for (auto IndexToRemove : IndexesToRemove) {
    NewIndexToRemove++;
    CommandTokens->erase(CommandTokens->begin() +
                         (IndexToRemove - NewIndexToRemove));
  }
  return TRUE;
}

BOOLEAN SendEventToKernel(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
                          UINT32 EventBufferLength) {
  BOOL Status;
  ULONG ReturnedLength;
  DEBUGGER_EVENT_AND_ACTION_RESULT ReturnedBuffer = {0};
  PDEBUGGER_EVENT_AND_ACTION_RESULT TempRegResult;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    TempRegResult =
        KdSendRegisterEventPacketToDebuggee(Event, EventBufferLength);
    memcpy(&ReturnedBuffer, TempRegResult,
           sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT));
  } else {
    AssertShowMessageReturnStmt(
        g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    Status = DeviceIoControl(g_DeviceHandle, IOCTL_DEBUGGER_REGISTER_EVENT,
                             Event, EventBufferLength, &ReturnedBuffer,
                             sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT),
                             &ReturnedLength, NULL);
    if (!Status) {
      ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
      return FALSE;
    }
  }
  if (ReturnedBuffer.IsSuccessful && ReturnedBuffer.Error == 0) {
    if (!g_IsSerialConnectedToRemoteDebuggee &&
        !g_IsSerialConnectedToRemoteDebugger && g_BreakPrintingOutput &&
        g_AutoUnpause) {
      g_BreakPrintingOutput = FALSE;
      if (g_IsConnectedToRemoteDebuggee) {
        RemoteConnectionSendCommand("g", (UINT32)strlen("g") + 1);
      }
      ShowMessages("\n");
    }
  } else {
    if (ReturnedBuffer.Error != 0) {
      ShowErrorMessage(ReturnedBuffer.Error);
    }
    return FALSE;
  }
  return TRUE;
}

BOOLEAN RegisterActionToEvent(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
                              PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger,
                              UINT32 ActionBreakToDebuggerLength,
                              PDEBUGGER_GENERAL_ACTION ActionCustomCode,
                              UINT32 ActionCustomCodeLength,
                              PDEBUGGER_GENERAL_ACTION ActionScript,
                              UINT32 ActionScriptLength) {
  BOOL Status;
  ULONG ReturnedLength;
  DEBUGGER_EVENT_AND_ACTION_RESULT ReturnedBuffer = {0};
  PDEBUGGER_EVENT_AND_ACTION_RESULT TempAddingResult;
  if (g_IsSerialConnectedToRemoteDebuggee) {
    if (ActionBreakToDebugger != NULL) {
      TempAddingResult = KdSendAddActionToEventPacketToDebuggee(
          ActionBreakToDebugger, ActionBreakToDebuggerLength);
      memcpy(&ReturnedBuffer, TempAddingResult,
             sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT));
    }
    if (ActionCustomCode != NULL) {
      TempAddingResult = KdSendAddActionToEventPacketToDebuggee(
          ActionCustomCode, ActionCustomCodeLength);
      memcpy(&ReturnedBuffer, TempAddingResult,
             sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT));
    }
    if (ActionScript != NULL) {
      TempAddingResult = KdSendAddActionToEventPacketToDebuggee(
          ActionScript, ActionScriptLength);
      memcpy(&ReturnedBuffer, TempAddingResult,
             sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT));
    }
  } else {
    AssertShowMessageReturnStmt(
        g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
    if (ActionBreakToDebugger != NULL) {
      Status = DeviceIoControl(
          g_DeviceHandle, IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT,
          ActionBreakToDebugger, ActionBreakToDebuggerLength, &ReturnedBuffer,
          sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT), &ReturnedLength, NULL);
      if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
      }
    }
    if (ActionCustomCode != NULL) {
      Status = DeviceIoControl(
          g_DeviceHandle, IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT, ActionCustomCode,
          ActionCustomCodeLength, &ReturnedBuffer,
          sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT), &ReturnedLength, NULL);
      if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
      }
    }
    if (ActionScript != NULL) {
      Status = DeviceIoControl(
          g_DeviceHandle, IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT, ActionScript,
          ActionScriptLength, &ReturnedBuffer,
          sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT), &ReturnedLength, NULL);
      if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
      }
    }
  }
  FreeEventsAndActionsMemory(NULL, ActionBreakToDebugger, ActionCustomCode,
                             ActionScript);
  InsertHeadList(&g_EventTrace, &(Event->CommandsEventList));
  return TRUE;
}

UINT64 GetNewDebuggerEventTag() { return g_EventTag++; }

VOID FreeEventsAndActionsMemory(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
                                PDEBUGGER_GENERAL_ACTION ActionBreakToDebugger,
                                PDEBUGGER_GENERAL_ACTION ActionCustomCode,
                                PDEBUGGER_GENERAL_ACTION ActionScript) {
  if (Event != NULL) {
    if (Event->CommandStringBuffer != NULL) {
      free(Event->CommandStringBuffer);
    }
    free(Event);
  }
  if (ActionBreakToDebugger != NULL) {
    free(ActionBreakToDebugger);
  }
  if (ActionCustomCode != NULL) {
    free(ActionCustomCode);
  }
  if (ActionScript != NULL) {
    free(ActionScript);
  }
}

BOOLEAN InterpretGeneralEventAndActionsFields(
    vector<CommandToken> *CommandTokens, VMM_EVENT_TYPE_ENUM EventType,
    PDEBUGGER_GENERAL_EVENT_DETAIL *EventDetailsToFill,
    PUINT32 EventBufferLength,
    PDEBUGGER_GENERAL_ACTION *ActionDetailsToFillBreakToDebugger,
    PUINT32 ActionBufferLengthBreakToDebugger,
    PDEBUGGER_GENERAL_ACTION *ActionDetailsToFillCustomCode,
    PUINT32 ActionBufferLengthCustomCode,
    PDEBUGGER_GENERAL_ACTION *ActionDetailsToFillScript,
    PUINT32 ActionBufferLengthScript,
    PDEBUGGER_EVENT_PARSING_ERROR_CAUSE ReasonForErrorInParsing) {
  BOOLEAN Result = FALSE;
  PDEBUGGER_GENERAL_EVENT_DETAIL TempEvent = NULL;
  PDEBUGGER_GENERAL_ACTION TempActionBreak = NULL;
  PDEBUGGER_GENERAL_ACTION TempActionScript = NULL;
  PDEBUGGER_GENERAL_ACTION TempActionCustomCode = NULL;
  VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE CallingStage =
      VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION;
  UINT32 LengthOfCustomCodeActionBuffer = 0;
  UINT32 LengthOfScriptActionBuffer = 0;
  UINT32 LengthOfBreakActionBuffer = 0;
  UINT64 ConditionBufferAddress;
  UINT32 ConditionBufferLength = 0;
  vector<string> ListOfOutputSources;
  UINT64 CodeBufferAddress;
  UINT32 CodeBufferLength = 0;
  UINT64 ScriptBufferAddress;
  UINT64 ScriptCodeBuffer = 0;
  BOOLEAN HasScriptSyntaxError = 0;
  UINT32 ScriptBufferLength = 0;
  UINT32 ScriptBufferPointer = 0;
  UINT32 LengthOfEventBuffer = 0;
  string CommandString;
  BOOLEAN IsAShortCircuitingEventByDefault = FALSE;
  BOOLEAN HasConditionBuffer = FALSE;
  BOOLEAN HasOutputPath = FALSE;
  BOOLEAN HasCodeBuffer = FALSE;
  BOOLEAN HasScript = FALSE;
  BOOLEAN IsNextCommandPid = FALSE;
  BOOLEAN IsNextCommandCoreId = FALSE;
  BOOLEAN IsNextCommandBufferSize = FALSE;
  BOOLEAN IsNextCommandImmediateMessaging = FALSE;
  BOOLEAN IsNextCommandExecutionStage = FALSE;
  BOOLEAN IsNextCommandSc = FALSE;
  BOOLEAN ImmediateMessagePassing = UseImmediateMessagingByDefaultOnEvents;
  UINT32 CoreId;
  UINT32 ProcessId;
  UINT32 IndexOfValidSourceTags;
  UINT32 RequestBuffer = 0;
  PLIST_ENTRY TempList;
  BOOLEAN OutputSourceFound;
  vector<int> IndexesToRemove;
  vector<UINT64> ListOfValidSourceTags;
  int NewIndexToRemove = 0;
  int Index = 0;
  for (auto Section : *CommandTokens) {
    CommandString.append(GetCaseSensitiveStringFromCommandToken(Section));
    CommandString.append(" ");
  }
  UINT64 BufferOfCommandStringLength = CommandString.size() + 1;
  PVOID BufferOfCommandString = malloc(BufferOfCommandStringLength);
  RtlZeroMemory(BufferOfCommandString, BufferOfCommandStringLength);
  memcpy(BufferOfCommandString, CommandString.c_str(), CommandString.size());
  if (!InterpretConditionsAndCodes(CommandTokens, TRUE, &ConditionBufferAddress,
                                   &ConditionBufferLength)) {
    HasConditionBuffer = FALSE;

  } else {
    HasConditionBuffer = TRUE;
  }
  if (!InterpretConditionsAndCodes(CommandTokens, FALSE, &CodeBufferAddress,
                                   &CodeBufferLength)) {
    HasCodeBuffer = FALSE;
  } else {
    HasCodeBuffer = TRUE;
  }
  if (!InterpretScript(CommandTokens, &HasScriptSyntaxError,
                       &ScriptBufferAddress, &ScriptBufferLength,
                       &ScriptBufferPointer, &ScriptCodeBuffer)) {
    HasScript = FALSE;
  } else {
    if (HasScriptSyntaxError) {
      free(BufferOfCommandString);
      *ReasonForErrorInParsing =
          DEBUGGER_EVENT_PARSING_ERROR_CAUSE_SCRIPT_SYNTAX_ERROR;
      return FALSE;
    }
    HasScript = TRUE;
  }
  if (!InterpretOutput(CommandTokens, ListOfOutputSources)) {
    HasOutputPath = FALSE;
  } else {
    if (ListOfOutputSources.size() == 0) {
      free(BufferOfCommandString);
      ShowMessages("err, no input found\n");
      *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_NO_INPUT;
      return FALSE;
    }
    if (ListOfOutputSources.size() >
        DebuggerOutputSourceMaximumRemoteSourceForSingleEvent) {
      free(BufferOfCommandString);
      ShowMessages(
          "err, based on this build of HyperDbg, the maximum input sources for "
          "a single event is 0x%x sources but you entered 0x%x sources\n",
          DebuggerOutputSourceMaximumRemoteSourceForSingleEvent,
          ListOfOutputSources.size());
      *ReasonForErrorInParsing =
          DEBUGGER_EVENT_PARSING_ERROR_CAUSE_MAXIMUM_INPUT_REACHED;
      return FALSE;
    }
    if (!g_OutputSourcesInitialized) {
      free(BufferOfCommandString);
      ShowMessages(
          "err, the name you entered, not found. Did you use "
          "'output' command to create it?\n");
      *ReasonForErrorInParsing =
          DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_NAME_NOT_FOUND;
      return FALSE;
    }
    for (auto item : ListOfOutputSources) {
      TempList = 0;
      OutputSourceFound = FALSE;
      TempList = &g_OutputSources;
      while (&g_OutputSources != TempList->Flink) {
        TempList = TempList->Flink;
        PDEBUGGER_EVENT_FORWARDING CurrentOutputSourceDetails =
            CONTAINING_RECORD(TempList, DEBUGGER_EVENT_FORWARDING,
                              OutputSourcesList);
        if (strcmp(CurrentOutputSourceDetails->Name,
                   RemoveSpaces(item).c_str()) == 0) {
          if (CurrentOutputSourceDetails->State == EVENT_FORWARDING_CLOSED) {
            free(BufferOfCommandString);
            ShowMessages("err, output source already closed\n");
            *ReasonForErrorInParsing =
                DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_SOURCE_ALREADY_CLOSED;
            return FALSE;
          }
          if (CurrentOutputSourceDetails->State ==
              EVENT_FORWARDING_STATE_NOT_OPENED) {
            ShowMessages(
                "some of the output(s) are not opened, it's not an error, but "
                "please ensure "
                "to open the output using the 'output' command to forward the "
                "results to the "
                "target resource\n");
          }
          OutputSourceFound = TRUE;
          ListOfValidSourceTags.push_back(
              CurrentOutputSourceDetails->OutputUniqueTag);
          break;
        }
      }
      if (!OutputSourceFound) {
        free(BufferOfCommandString);
        ShowMessages(
            "err, the name you entered, not found. Did you use "
            "'output' command to create it?\n");
        *ReasonForErrorInParsing =
            DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_NAME_NOT_FOUND;
        return FALSE;
      }
    }
    HasOutputPath = TRUE;
  }

  LengthOfEventBuffer =
      sizeof(DEBUGGER_GENERAL_EVENT_DETAIL) + ConditionBufferLength;
  TempEvent = (PDEBUGGER_GENERAL_EVENT_DETAIL)malloc(LengthOfEventBuffer);
  RtlZeroMemory(TempEvent, LengthOfEventBuffer);
  if (TempEvent == NULL) {
    ShowMessages("err, allocation error\n");
    *ReasonForErrorInParsing =
        DEBUGGER_EVENT_PARSING_ERROR_CAUSE_ALLOCATION_ERROR;
    goto ReturnWithError;
  }
  TempEvent->IsEnabled = TRUE;
  TempEvent->Tag = GetNewDebuggerEventTag();
  TempEvent->CoreId = DEBUGGER_EVENT_APPLY_TO_ALL_CORES;
  if (g_ActiveProcessDebuggingState.IsActive) {
    ShowMessages(
        "notice: as you're debugging a user-mode application, "
        "this event will only trigger on your current debugging process "
        "(pid:%x). If you want the event from the entire system, "
        "add 'pid all' to the event\n",
        g_ActiveProcessDebuggingState.ProcessId);
    TempEvent->ProcessId = g_ActiveProcessDebuggingState.ProcessId;
  } else {
    TempEvent->ProcessId = DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES;
  }
  TempEvent->EventType = EventType;
  TempEvent->CreationTime = time(0);
  TempEvent->CommandStringBuffer = BufferOfCommandString;
  if (HasConditionBuffer) {
    memcpy((PVOID)((UINT64)TempEvent + sizeof(DEBUGGER_GENERAL_EVENT_DETAIL)),
           (PVOID)ConditionBufferAddress, ConditionBufferLength);
    TempEvent->ConditionBufferSize = ConditionBufferLength;
  }
  if (HasCodeBuffer) {
    LengthOfCustomCodeActionBuffer =
        sizeof(DEBUGGER_GENERAL_ACTION) + CodeBufferLength;
    TempActionCustomCode =
        (PDEBUGGER_GENERAL_ACTION)malloc(LengthOfCustomCodeActionBuffer);
    RtlZeroMemory(TempActionCustomCode, LengthOfCustomCodeActionBuffer);
    memcpy(
        (PVOID)((UINT64)TempActionCustomCode + sizeof(DEBUGGER_GENERAL_ACTION)),
        (PVOID)CodeBufferAddress, CodeBufferLength);
    TempActionCustomCode->EventTag = TempEvent->Tag;
    TempActionCustomCode->ActionType = RUN_CUSTOM_CODE;
    TempActionCustomCode->CustomCodeBufferSize = CodeBufferLength;
    TempEvent->CountOfActions = TempEvent->CountOfActions + 1;
  }
  if (HasScript) {
    LengthOfScriptActionBuffer =
        sizeof(DEBUGGER_GENERAL_ACTION) + ScriptBufferLength;
    TempActionScript =
        (PDEBUGGER_GENERAL_ACTION)malloc(LengthOfScriptActionBuffer);
    RtlZeroMemory(TempActionScript, LengthOfScriptActionBuffer);
    memcpy((PVOID)((UINT64)TempActionScript + sizeof(DEBUGGER_GENERAL_ACTION)),
           (PVOID)ScriptBufferAddress, ScriptBufferLength);
    TempActionScript->EventTag = TempEvent->Tag;
    TempActionScript->ActionType = RUN_SCRIPT;
    TempActionScript->ScriptBufferSize = ScriptBufferLength;
    TempActionScript->ScriptBufferPointer = ScriptBufferPointer;
    TempEvent->CountOfActions = TempEvent->CountOfActions + 1;
    ScriptEngineWrapperRemoveSymbolBuffer((PVOID)ScriptCodeBuffer);
  }
  if (!HasCodeBuffer && !HasScript) {
    LengthOfBreakActionBuffer = sizeof(DEBUGGER_GENERAL_ACTION);
    TempActionBreak =
        (PDEBUGGER_GENERAL_ACTION)malloc(LengthOfBreakActionBuffer);
    RtlZeroMemory(TempActionBreak, LengthOfBreakActionBuffer);
    TempActionBreak->EventTag = TempEvent->Tag;
    TempActionBreak->ActionType = BREAK_TO_DEBUGGER;
    TempEvent->CountOfActions = TempEvent->CountOfActions + 1;
  }
  for (auto Section : *CommandTokens) {
    Index++;
    if (IsNextCommandBufferSize) {
      if (!ConvertTokenToUInt32(Section, &RequestBuffer)) {
        ShowMessages("err, buffer size is invalid\n");
        *ReasonForErrorInParsing =
            DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
      } else {
        if (TempActionBreak != NULL) {
          TempActionBreak->PreAllocatedBuffer = RequestBuffer;
        }
        if (TempActionScript != NULL) {
          TempActionScript->PreAllocatedBuffer = RequestBuffer;
        }
        if (TempActionCustomCode != NULL) {
          TempActionCustomCode->PreAllocatedBuffer = RequestBuffer;
        }
      }
      IsNextCommandBufferSize = FALSE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (IsNextCommandImmediateMessaging) {
      if (CompareLowerCaseStrings(Section, "yes")) {
        ImmediateMessagePassing = TRUE;
      } else if (CompareLowerCaseStrings(Section, "no")) {
        ImmediateMessagePassing = FALSE;
      } else {
        ShowMessages("err, immediate messaging token is invalid\n");
        *ReasonForErrorInParsing =
            DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
      }
      IsNextCommandImmediateMessaging = FALSE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (IsNextCommandExecutionStage) {
      if (CompareLowerCaseStrings(Section, "pre")) {
        CallingStage = VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION;
      } else if (CompareLowerCaseStrings(Section, "post")) {
        CallingStage = VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION;
      } else if (CompareLowerCaseStrings(Section, "all")) {
        CallingStage = VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION;
      } else {
        ShowMessages(
            "err, the specified execution mode is invalid; you can either "
            "choose 'pre' or 'post'\n");
        *ReasonForErrorInParsing =
            DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
      }
      IsNextCommandExecutionStage = FALSE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (IsNextCommandSc) {
      if (CompareLowerCaseStrings(Section, "on")) {
        IsAShortCircuitingEventByDefault = TRUE;
      } else if (CompareLowerCaseStrings(Section, "off")) {
        IsAShortCircuitingEventByDefault = FALSE;
      } else {
        ShowMessages(
            "err, the specified short-circuiting state is invalid; you can "
            "either choose 'on' or 'off'\n");
        *ReasonForErrorInParsing =
            DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
      }
      IsNextCommandSc = FALSE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (IsNextCommandPid) {
      if (CompareLowerCaseStrings(Section, "all")) {
        TempEvent->ProcessId = DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES;
      } else if (!ConvertTokenToUInt32(Section, &ProcessId)) {
        ShowMessages("err, pid is invalid\n");
        *ReasonForErrorInParsing =
            DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
      } else {
        TempEvent->ProcessId = ProcessId;
      }
      IsNextCommandPid = FALSE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (IsNextCommandCoreId) {
      if (!ConvertTokenToUInt32(Section, &CoreId)) {
        ShowMessages("err, core id is invalid\n");
        *ReasonForErrorInParsing =
            DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
      } else {
        TempEvent->CoreId = CoreId;
      }
      IsNextCommandCoreId = FALSE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (CompareLowerCaseStrings(Section, "pid")) {
      IsNextCommandPid = TRUE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (CompareLowerCaseStrings(Section, "core")) {
      IsNextCommandCoreId = TRUE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (CompareLowerCaseStrings(Section, "imm")) {
      IsNextCommandImmediateMessaging = TRUE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (CompareLowerCaseStrings(Section, "stage")) {
      IsNextCommandExecutionStage = TRUE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (CompareLowerCaseStrings(Section, "sc")) {
      IsNextCommandSc = TRUE;
      IndexesToRemove.push_back(Index);
      continue;
    }
    if (CompareLowerCaseStrings(Section, "buffer")) {
      IsNextCommandBufferSize = TRUE;
      IndexesToRemove.push_back(Index);
      continue;
    }
  }
  if (IsNextCommandCoreId) {
    ShowMessages("err, please specify a value for 'core'\n");
    *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
    goto ReturnWithError;
  }
  if (IsNextCommandPid) {
    ShowMessages("err, please specify a value for 'pid'\n");
    *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
    goto ReturnWithError;
  }
  if (IsNextCommandBufferSize) {
    ShowMessages("err, please specify a value for 'buffer'\n");
    *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
    goto ReturnWithError;
  }
  if (IsNextCommandImmediateMessaging) {
    ShowMessages("err, please specify a value for 'imm'\n");
    *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
    goto ReturnWithError;
  }
  if (IsNextCommandExecutionStage) {
    ShowMessages("err, please specify a value for 'stage'\n");
    *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
    goto ReturnWithError;
  }
  if (IsNextCommandSc) {
    ShowMessages("err, please specify a value for 'sc'\n");
    *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
    goto ReturnWithError;
  }
  if ((CallingStage == VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION ||
       CallingStage == VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION) &&
      IsAShortCircuitingEventByDefault) {
    ShowMessages(
        "err, using the short-circuiting mechanism with 'post' or 'all' stage "
        "events "
        "doesn't make sense; it's not supported!\n");
    *ReasonForErrorInParsing =
        DEBUGGER_EVENT_PARSING_ERROR_CAUSE_USING_SHORT_CIRCUITING_IN_POST_EVENTS;
    goto ReturnWithError;
  }
  if (!g_IsSerialConnectedToRemoteDebuggee && TempActionBreak != NULL) {
    ShowMessages(
        "err, the script or assembly code is either not found or invalid. "
        "As a result, the default action is to break. "
        "However, breaking to the debugger is not possible in the VMI Mode. "
        "To achieve full control of the system, you can switch to the Debugger "
        "Mode. "
        "In the VMI Mode, you can still use scripts and run custom code for "
        "local debugging."
        "For more information, please check: "
        "https://docs.hyperdbg.org/using-hyperdbg/prerequisites/"
        "operation-modes\n");
    *ReasonForErrorInParsing =
        DEBUGGER_EVENT_PARSING_ERROR_CAUSE_ATTEMPT_TO_BREAK_ON_VMI_MODE;
    goto ReturnWithError;
  }
  if (!ImmediateMessagePassing && HasOutputPath) {
    ShowMessages(
        "err, non-immediate message passing is not supported in "
        "'output-forwarding mode'\n");
    *ReasonForErrorInParsing =
        DEBUGGER_EVENT_PARSING_ERROR_CAUSE_IMMEDIATE_MESSAGING_IN_EVENT_FORWARDING_MODE;
    goto ReturnWithError;
  }
  if (TempActionBreak != NULL) {
    TempActionBreak->ImmediateMessagePassing = ImmediateMessagePassing;
  }
  if (TempActionScript != NULL) {
    TempActionScript->ImmediateMessagePassing = ImmediateMessagePassing;
  }
  if (TempActionCustomCode != NULL) {
    TempActionCustomCode->ImmediateMessagePassing = ImmediateMessagePassing;
  }
  IndexOfValidSourceTags = 0;
  for (auto item : ListOfValidSourceTags) {
    TempEvent->OutputSourceTags[IndexOfValidSourceTags] = item;
    IndexOfValidSourceTags++;
  }
  if (HasOutputPath) {
    TempEvent->HasCustomOutput = TRUE;
  }
  if (IsAShortCircuitingEventByDefault) {
    TempEvent->EnableShortCircuiting = TRUE;
  }
  TempEvent->EventStage = CallingStage;
  *EventDetailsToFill = TempEvent;
  *EventBufferLength = LengthOfEventBuffer;
  if (TempActionBreak != NULL) {
    *ActionDetailsToFillBreakToDebugger = TempActionBreak;
    *ActionBufferLengthBreakToDebugger = LengthOfBreakActionBuffer;
  }
  if (TempActionScript != NULL) {
    *ActionDetailsToFillScript = TempActionScript;
    *ActionBufferLengthScript = LengthOfScriptActionBuffer;
  }
  if (TempActionCustomCode != NULL) {
    *ActionDetailsToFillCustomCode = TempActionCustomCode;
    *ActionBufferLengthCustomCode = LengthOfCustomCodeActionBuffer;
  }
  for (auto IndexToRemove : IndexesToRemove) {
    NewIndexToRemove++;
    CommandTokens->erase(CommandTokens->begin() +
                         (IndexToRemove - NewIndexToRemove));
  }
  if (!g_EventTraceInitialized) {
    InitializeListHead(&g_EventTrace);
    g_EventTraceInitialized = TRUE;
  }
  *ReasonForErrorInParsing =
      DEBUGGER_EVENT_PARSING_ERROR_CAUSE_SUCCESSFUL_NO_ERROR;
  return TRUE;
ReturnWithError:
  if (BufferOfCommandString) {
    free(BufferOfCommandString);
  }
  if (TempEvent) {
    free(TempEvent);
  }
  if (TempActionBreak != NULL) {
    free(TempActionBreak);
  }
  if (TempActionScript != NULL) {
    free(TempActionScript);
  }
  if (TempActionCustomCode != NULL) {
    free(TempActionCustomCode);
  }
  return FALSE;
}
