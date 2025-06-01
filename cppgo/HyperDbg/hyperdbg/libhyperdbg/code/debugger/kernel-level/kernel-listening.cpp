
#include "pch.h"
extern DEBUGGER_SYNCRONIZATION_EVENTS_STATE
    g_KernelSyncronizationObjectsHandleTable
        [DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS];
extern BYTE g_CurrentRunningInstruction[MAXIMUM_INSTR_SIZE];
extern OVERLAPPED g_OverlappedIoStructureForReadDebugger;
extern OVERLAPPED g_OverlappedIoStructureForWriteDebugger;
extern HANDLE g_SerialRemoteComPortHandle;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsDebuggeeRunning;
extern BOOLEAN g_IgnoreNewLoggingMessages;
extern BOOLEAN g_SharedEventStatus;
extern BOOLEAN g_IsRunningInstruction32Bit;
extern BOOLEAN g_OutputSourcesInitialized;
extern ULONG g_CurrentRemoteCore;
extern DEBUGGER_EVENT_AND_ACTION_RESULT g_DebuggeeResultOfRegisteringEvent;
extern DEBUGGER_EVENT_AND_ACTION_RESULT g_DebuggeeResultOfAddingActionsToEvent;
extern UINT64 g_ResultOfEvaluatedExpression;
extern UINT32 g_ErrorStateOfResultOfEvaluatedExpression;
extern UINT64 g_KernelBaseAddress;

BOOLEAN ListeningSerialPortInDebugger() {
  PDEBUGGER_PREPARE_DEBUGGEE InitPacket;
  PDEBUGGER_REMOTE_PACKET TheActualPacket;
  PDEBUGGEE_KD_PAUSED_PACKET PausePacket;
  PDEBUGGEE_MESSAGE_PACKET MessagePacket;
  PDEBUGGEE_CHANGE_CORE_PACKET ChangeCorePacket;
  PDEBUGGEE_SCRIPT_PACKET ScriptPacket;
  PDEBUGGEE_FORMATS_PACKET FormatsPacket;
  PDEBUGGER_EVENT_AND_ACTION_RESULT EventAndActionPacket;
  PDEBUGGER_UPDATE_SYMBOL_TABLE SymbolUpdatePacket;
  PDEBUGGER_MODIFY_EVENTS EventModifyAndQueryPacket;
  PDEBUGGEE_SYMBOL_UPDATE_RESULT SymbolReloadFinishedPacket;
  PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET ChangeProcessPacket;
  PDEBUGGEE_RESULT_OF_SEARCH_PACKET SearchResultsPacket;
  PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET ChangeThreadPacket;
  PDEBUGGER_FLUSH_LOGGING_BUFFERS FlushPacket;
  PDEBUGGER_CALLSTACK_REQUEST CallstackPacket;
  PDEBUGGER_SINGLE_CALLSTACK_FRAME CallstackFramePacket;
  PDEBUGGER_DEBUGGER_TEST_QUERY_BUFFER TestQueryPacket;
  PDEBUGGEE_REGISTER_READ_DESCRIPTION ReadRegisterPacket;
  PDEBUGGEE_REGISTER_WRITE_DESCRIPTION WriteRegisterPacket;
  PDEBUGGER_APIC_REQUEST ApicRequestPacket;
  PDEBUGGER_READ_MEMORY ReadMemoryPacket;
  PDEBUGGER_EDIT_MEMORY EditMemoryPacket;
  PDEBUGGEE_BP_PACKET BpPacket;
  PDEBUGGER_SHORT_CIRCUITING_EVENT ShortCircuitingPacket;
  PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PtePacket;
  PDEBUGGER_PAGE_IN_REQUEST PageinPacket;
  PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS Va2paPa2vaPacket;
  PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET ListOrModifyBreakpointPacket;
  BOOLEAN ShowSignatureWhenDisconnected = FALSE;
  PVOID CallerAddress = NULL;
  UINT32 CallerSize = NULL_ZERO;
  PDEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET PcitreePacket;
  PINTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS IdtEntryRequestPacket;
  PDEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET PcidevinfoPacket;
StartAgain:
  CHAR BufferToReceive[MaxSerialPacketSize] = {0};
  UINT32 LengthReceived = 0;
  if (!KdReceivePacketFromDebuggee(BufferToReceive, &LengthReceived)) {
    if (LengthReceived == 0 && BufferToReceive[0] == NULL) {
      ShowMessages("\nthe remote connection is closed\n");
      if (g_IsSerialConnectedToRemoteDebuggee) {
        CommandEventsClearAllEventsAndResetTags();
        if (g_IsDebuggeeRunning == FALSE) {
          ShowSignatureWhenDisconnected = TRUE;
        }
      }
      KdCloseConnection();
      if (ShowSignatureWhenDisconnected) {
        ShowSignatureWhenDisconnected = FALSE;
        ShowMessages("\n");
      }
      return FALSE;
    } else {
      ShowMessages("err, invalid buffer received\n");
      goto StartAgain;
    }
  }
  if (LengthReceived == 1 && BufferToReceive[0] == NULL) {
    goto StartAgain;
  }
  TheActualPacket = (PDEBUGGER_REMOTE_PACKET)BufferToReceive;
  if (TheActualPacket->Indicator == INDICATOR_OF_HYPERDBG_PACKET) {
    if (KdComputeDataChecksum((PVOID)&TheActualPacket->Indicator,
                              LengthReceived - sizeof(BYTE)) !=
        TheActualPacket->Checksum) {
      ShowMessages("\nerr, checksum is invalid\n");
      goto StartAgain;
    }
    if (TheActualPacket->TypeOfThePacket !=
        DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER) {
      ShowMessages("\nerr, unknown packet received from the debuggee\n");
      goto StartAgain;
    }
    switch (TheActualPacket->RequestedActionOfThePacket) {
      case DEBUGGER_REMOTE_PACKET_PING_AND_SEND_SUPPORTED_VERSION:
        KdSendResponseOfThePingPacket();
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_STARTED:
        InitPacket =
            (DEBUGGER_PREPARE_DEBUGGEE *)(((CHAR *)TheActualPacket) +
                                          sizeof(DEBUGGER_REMOTE_PACKET));
        g_KernelBaseAddress = InitPacket->KernelBaseAddress;
        ShowMessages("connected to debuggee %s\n", InitPacket->OsName);
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_STARTED_PACKET_RECEIVED);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM:
        MessagePacket =
            (DEBUGGEE_MESSAGE_PACKET *)(((CHAR *)TheActualPacket) +
                                        sizeof(DEBUGGER_REMOTE_PACKET));
        if (!g_OutputSourcesInitialized ||
            !ForwardingCheckAndPerformEventForwarding(
                MessagePacket->OperationCode, MessagePacket->Message,
                (UINT32)strlen(MessagePacket->Message))) {
          if (!g_IgnoreNewLoggingMessages) {
            ShowMessages("%s", MessagePacket->Message);
          }
        }
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION:
        g_IgnoreNewLoggingMessages = TRUE;
        PausePacket =
            (DEBUGGEE_KD_PAUSED_PACKET *)(((CHAR *)TheActualPacket) +
                                          sizeof(DEBUGGER_REMOTE_PACKET));
        g_IsDebuggeeRunning = FALSE;
        g_CurrentRemoteCore = PausePacket->CurrentCore;
        RtlZeroMemory(g_CurrentRunningInstruction, MAXIMUM_INSTR_SIZE);
        memcpy(g_CurrentRunningInstruction, &PausePacket->InstructionBytesOnRip,
               MAXIMUM_INSTR_SIZE);
        g_IsRunningInstruction32Bit = PausePacket->IsProcessorOn32BitMode;
        switch (PausePacket->PausingReason) {
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT:
            if (PausePacket->EventTag != NULL) {
              ShowMessages("breakpoint 0x%x hit\n", PausePacket->EventTag);
            }
            break;
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED:
            if (PausePacket->EventTag != NULL) {
              if (PausePacket->EventCallingStage ==
                  VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION) {
                ShowMessages("event 0x%x triggered (post)\n",
                             PausePacket->EventTag - DebuggerEventTagStartSeed);
              } else {
                ShowMessages("event 0x%x triggered (pre)\n",
                             PausePacket->EventTag - DebuggerEventTagStartSeed);
              }
            }
            break;
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED:
            ShowMessages("switched to the specified process\n");
            break;
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED:
            ShowMessages("switched to the specified thread\n");
            break;
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_STARTING_MODULE_LOADED:
            ShowMessages(
                "the target module is loaded and a breakpoint is set to the "
                "entrypoint\n"
                "press 'g' to reach to the entrypoint of the main module...\n");
            break;
          default:
            break;
        }
        if (!PausePacket->IgnoreDisassembling) {
          if (PausePacket->ReadInstructionLen != MAXIMUM_INSTR_SIZE) {
            if (HyperDbgLengthDisassemblerEngine(
                    PausePacket->InstructionBytesOnRip, MAXIMUM_INSTR_SIZE,
                    PausePacket->IsProcessorOn32BitMode ? FALSE : TRUE) >
                PausePacket->ReadInstructionLen) {
              ShowMessages(
                  "oOh, no! there might be a misinterpretation in "
                  "disassembling the current instruction\n");
            }
          }
          if (!PausePacket->IsProcessorOn32BitMode) {
            HyperDbgDisassembler64(PausePacket->InstructionBytesOnRip,
                                   PausePacket->Rip, MAXIMUM_INSTR_SIZE, 1,
                                   TRUE, (PRFLAGS)&PausePacket->Rflags);
          } else {
            HyperDbgDisassembler32(PausePacket->InstructionBytesOnRip,
                                   PausePacket->Rip, MAXIMUM_INSTR_SIZE, 1,
                                   TRUE, (PRFLAGS)&PausePacket->Rflags);
          }
        }
        switch (PausePacket->PausingReason) {
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT:
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT:
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED:
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED:
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED:
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED:
            DbgReceivedKernelResponse(
                DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);
            break;
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_TRACKING_STEPPED:
            CommandTrackHandleReceivedInstructions(
                &PausePacket->InstructionBytesOnRip[0], MAXIMUM_INSTR_SIZE,
                PausePacket->IsProcessorOn32BitMode ? FALSE : TRUE,
                PausePacket->Rip);
            DbgReceivedKernelResponse(
                DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);
            break;
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_STARTING_MODULE_LOADED:
            ShowMessages("\n");
            DbgReceivedKernelResponse(
                DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);
            break;
          case DEBUGGEE_PAUSING_REASON_PAUSE:
            break;
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED:
            DbgReceivedKernelResponse(
                DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CORE_SWITCHING_RESULT);
            break;
          case DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED:
            ShowMessages("\n");
            DbgReceivedKernelResponse(
                DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_DEBUGGEE_FINISHED_COMMAND_EXECUTION);
            break;
          case DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER:
            DbgReceivedKernelResponse(
                DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PAUSED_DEBUGGEE_DETAILS);
            break;
          default:
            ShowMessages("err, unknown pausing reason is received\n");
            break;
        }
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE:
        ChangeCorePacket =
            (DEBUGGEE_CHANGE_CORE_PACKET *)(((CHAR *)TheActualPacket) +
                                            sizeof(DEBUGGER_REMOTE_PACKET));
        if (ChangeCorePacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          ShowMessages("current operating core changed to 0x%x\n",
                       ChangeCorePacket->NewCore);
        } else {
          ShowErrorMessage(ChangeCorePacket->Result);
          DbgReceivedKernelResponse(
              DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CORE_SWITCHING_RESULT);
        }
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS:
        ChangeProcessPacket =
            (DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET));
        if (ChangeProcessPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          if (ChangeProcessPacket->ActionType ==
              DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS) {
            ShowMessages(
                "process id: %x\nprocess (_EPROCESS): %s\nprocess name "
                "(16-Byte): %s\n",
                ChangeProcessPacket->ProcessId,
                SeparateTo64BitValue(ChangeProcessPacket->Process).c_str(),
                &ChangeProcessPacket->ProcessName);
          } else if (ChangeProcessPacket->ActionType ==
                     DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH) {
            ShowMessages(
                "press 'g' to continue the debuggee, if the pid or the "
                "process object address is valid then the debuggee will "
                "be automatically paused when it attached to the target "
                "process\n");
          }
        } else {
          ShowErrorMessage(ChangeProcessPacket->Result);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PROCESS_SWITCHING_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY:
        SearchResultsPacket =
            (DEBUGGEE_RESULT_OF_SEARCH_PACKET *)(((CHAR *)TheActualPacket) +
                                                 sizeof(
                                                     DEBUGGER_REMOTE_PACKET));
        if (SearchResultsPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          if (SearchResultsPacket->CountOfResults == 0) {
            ShowMessages("not found\n");
          }
        } else {
          ShowErrorMessage(SearchResultsPacket->Result);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SEARCH_QUERY_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD:
        ChangeThreadPacket =
            (DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET));
        if (ChangeThreadPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          if (ChangeThreadPacket->ActionType ==
              DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS) {
            ShowMessages(
                "thread id: %x (pid: %x)\nthread (_ETHREAD): %s\nprocess "
                "(_EPROCESS): %s\nprocess name (16-Byte): %s\n",
                ChangeThreadPacket->ThreadId, ChangeThreadPacket->ProcessId,
                SeparateTo64BitValue(ChangeThreadPacket->Thread).c_str(),
                SeparateTo64BitValue(ChangeThreadPacket->Process).c_str(),
                &ChangeThreadPacket->ProcessName);
          } else if (ChangeThreadPacket->ActionType ==
                     DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH) {
            ShowMessages(
                "press 'g' to continue the debuggee, if the tid or the "
                "thread object address is valid then the debuggee will "
                "be automatically paused when it attached to the target "
                "thread\n");
          }
        } else {
          ShowErrorMessage(ChangeThreadPacket->Result);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_THREAD_SWITCHING_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH:
        FlushPacket =
            (DEBUGGER_FLUSH_LOGGING_BUFFERS *)(((CHAR *)TheActualPacket) +
                                               sizeof(DEBUGGER_REMOTE_PACKET));
        if (FlushPacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          ShowMessages(
              "flushing buffers was successful, total %d messages were "
              "cleared.\n",
              FlushPacket->CountOfMessagesThatSetAsReadFromVmxNonRoot +
                  FlushPacket->CountOfMessagesThatSetAsReadFromVmxRoot);
        } else {
          ShowErrorMessage(FlushPacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_FLUSH_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK:
        CallstackPacket =
            (DEBUGGER_CALLSTACK_REQUEST *)(((CHAR *)TheActualPacket) +
                                           sizeof(DEBUGGER_REMOTE_PACKET));
        CallstackFramePacket =
            (DEBUGGER_SINGLE_CALLSTACK_FRAME
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET) +
                    sizeof(DEBUGGER_CALLSTACK_REQUEST));
        if (CallstackPacket->KernelStatus ==
            DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          CallstackShowFrames(CallstackFramePacket, CallstackPacket->FrameCount,
                              CallstackPacket->DisplayMethod,
                              CallstackPacket->Is32Bit);
        } else {
          ShowErrorMessage(CallstackPacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CALLSTACK_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY:
        TestQueryPacket =
            (DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET));
        if (TestQueryPacket->KernelStatus ==
            DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          switch (TestQueryPacket->RequestType) {
            case TEST_BREAKPOINT_TURN_OFF_BPS:
              ShowMessages(
                  "breakpoint interception (#BP) is deactivated\n"
                  "from now, the breakpoints will be re-injected into the "
                  "guest debuggee\n");
              break;
            case TEST_BREAKPOINT_TURN_ON_BPS:
              ShowMessages("breakpoint interception (#BP) is activated\n");
              break;
            case TEST_BREAKPOINT_TURN_OFF_DBS:
              ShowMessages(
                  "debug break interception (#DB) is deactivated\n"
                  "from now, the debug breaks will be re-injected into the "
                  "guest debuggee\n");
              break;
            case TEST_BREAKPOINT_TURN_ON_DBS:
              ShowMessages("debug break interception (#DB) is activated\n");
              break;
            default:
              break;
          }
        } else {
          ShowErrorMessage(TestQueryPacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_TEST_QUERY);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT:
        ScriptPacket =
            (DEBUGGEE_SCRIPT_PACKET *)(((CHAR *)TheActualPacket) +
                                       sizeof(DEBUGGER_REMOTE_PACKET));
        if (ScriptPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
        } else {
          ShowErrorMessage(ScriptPacket->Result);
        }
        if (ScriptPacket->IsFormat) {
          DbgReceivedKernelResponse(
              DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_FORMATS_RESULT);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_RUNNING_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS:
        FormatsPacket =
            (DEBUGGEE_FORMATS_PACKET *)(((CHAR *)TheActualPacket) +
                                        sizeof(DEBUGGER_REMOTE_PACKET));
        g_ErrorStateOfResultOfEvaluatedExpression = FormatsPacket->Result;
        g_ResultOfEvaluatedExpression = FormatsPacket->Value;
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT:
        EventAndActionPacket =
            (DEBUGGER_EVENT_AND_ACTION_RESULT *)(((CHAR *)TheActualPacket) +
                                                 sizeof(
                                                     DEBUGGER_REMOTE_PACKET));
        memcpy(&g_DebuggeeResultOfRegisteringEvent, EventAndActionPacket,
               sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT));
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_REGISTER_EVENT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT:
        EventAndActionPacket =
            (DEBUGGER_EVENT_AND_ACTION_RESULT *)(((CHAR *)TheActualPacket) +
                                                 sizeof(
                                                     DEBUGGER_REMOTE_PACKET));
        memcpy(&g_DebuggeeResultOfAddingActionsToEvent, EventAndActionPacket,
               sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT));
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_ADD_ACTION_TO_EVENT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT:
        EventModifyAndQueryPacket =
            (DEBUGGER_MODIFY_EVENTS *)(((CHAR *)TheActualPacket) +
                                       sizeof(DEBUGGER_REMOTE_PACKET));
        if (EventModifyAndQueryPacket->KernelStatus !=
            DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          ShowErrorMessage((UINT32)EventModifyAndQueryPacket->KernelStatus);
        } else if (EventModifyAndQueryPacket->TypeOfAction ==
                   DEBUGGER_MODIFY_EVENTS_QUERY_STATE) {
          g_SharedEventStatus = EventModifyAndQueryPacket->IsEnabled;
        } else {
          CommandEventsHandleModifiedEvent(EventModifyAndQueryPacket->Tag,
                                           EventModifyAndQueryPacket);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_MODIFY_AND_QUERY_EVENT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED:
        SymbolReloadFinishedPacket =
            (DEBUGGEE_SYMBOL_UPDATE_RESULT *)(((CHAR *)TheActualPacket) +
                                              sizeof(DEBUGGER_REMOTE_PACKET));
        if (SymbolReloadFinishedPacket->KernelStatus !=
            DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          ShowErrorMessage((UINT32)SymbolReloadFinishedPacket->KernelStatus);
        } else {
          SymbolInitialReload();
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SYMBOL_RELOAD);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS:
        ReadRegisterPacket =
            (DEBUGGEE_REGISTER_READ_DESCRIPTION *)(((CHAR *)TheActualPacket) +
                                                   sizeof(
                                                       DEBUGGER_REMOTE_PACKET));
        DbgWaitGetRequestData(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_REGISTERS,
            &CallerAddress, &CallerSize);
        memcpy(CallerAddress, ReadRegisterPacket, CallerSize);
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_REGISTERS);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_WRITE_REGISTER:
        WriteRegisterPacket =
            (DEBUGGEE_REGISTER_WRITE_DESCRIPTION
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET));
        DbgWaitGetRequestData(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_WRITE_REGISTER,
            &CallerAddress, &CallerSize);
        memcpy(CallerAddress, WriteRegisterPacket, CallerSize);
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_WRITE_REGISTER);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_APIC_REQUESTS:
        ApicRequestPacket =
            (DEBUGGER_APIC_REQUEST *)(((CHAR *)TheActualPacket) +
                                      sizeof(DEBUGGER_REMOTE_PACKET));
        DbgWaitGetRequestData(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_APIC_ACTIONS,
            &CallerAddress, &CallerSize);
        memcpy(CallerAddress, ApicRequestPacket, CallerSize);
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_APIC_ACTIONS);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_IDT_ENTRIES_REQUESTS:
        IdtEntryRequestPacket =
            (INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET));
        DbgWaitGetRequestData(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IDT_ENTRIES,
            &CallerAddress, &CallerSize);
        memcpy(CallerAddress, IdtEntryRequestPacket, CallerSize);
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IDT_ENTRIES);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY:
        ReadMemoryPacket =
            (DEBUGGER_READ_MEMORY *)(((CHAR *)TheActualPacket) +
                                     sizeof(DEBUGGER_REMOTE_PACKET));
        DbgWaitGetRequestData(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_MEMORY,
            &CallerAddress, &CallerSize);
        memcpy(CallerAddress, ReadMemoryPacket, CallerSize);
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_MEMORY);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY:
        EditMemoryPacket =
            (DEBUGGER_EDIT_MEMORY *)(((CHAR *)TheActualPacket) +
                                     sizeof(DEBUGGER_REMOTE_PACKET));
        DbgWaitGetRequestData(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_EDIT_MEMORY,
            &CallerAddress, &CallerSize);
        memcpy(CallerAddress, EditMemoryPacket, CallerSize);
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_EDIT_MEMORY);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP:
        BpPacket = (DEBUGGEE_BP_PACKET *)(((CHAR *)TheActualPacket) +
                                          sizeof(DEBUGGER_REMOTE_PACKET));
        if (BpPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
        } else {
          ShowErrorMessage(BpPacket->Result);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_BP);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_STATE:
        ShortCircuitingPacket =
            (DEBUGGER_SHORT_CIRCUITING_EVENT *)(((CHAR *)TheActualPacket) +
                                                sizeof(DEBUGGER_REMOTE_PACKET));
        if (ShortCircuitingPacket->KernelStatus ==
            DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          ShowMessages(
              "the event's short-circuiting state changed to %s\n",
              ShortCircuitingPacket->IsShortCircuiting ? "'on'" : "'off'");
        } else {
          ShowErrorMessage((UINT32)ShortCircuitingPacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SHORT_CIRCUITING_EVENT_STATE);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE:
        PtePacket =
            (DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET));
        if (PtePacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          CommandPteShowResults(PtePacket->VirtualAddress, PtePacket);
        } else {
          ShowErrorMessage(PtePacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PTE_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BRINGING_PAGES_IN:
        PageinPacket =
            (DEBUGGER_PAGE_IN_REQUEST *)(((CHAR *)TheActualPacket) +
                                         sizeof(DEBUGGER_REMOTE_PACKET));
        if (PageinPacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          ShowMessages(
              "the page-fault is delivered to the target thread\n"
              "press 'g' to continue debuggee (the current thread will execute "
              "ONLY one instruction and will be halted again)...\n");
        } else {
          ShowErrorMessage(PageinPacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PAGE_IN_STATE);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA:
        Va2paPa2vaPacket =
            (DEBUGGER_VA2PA_AND_PA2VA_COMMANDS *)(((CHAR *)TheActualPacket) +
                                                  sizeof(
                                                      DEBUGGER_REMOTE_PACKET));
        if (Va2paPa2vaPacket->KernelStatus ==
            DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          if (Va2paPa2vaPacket->IsVirtual2Physical) {
            ShowMessages("%llx\n", Va2paPa2vaPacket->PhysicalAddress);
          } else {
            ShowMessages("%llx\n", Va2paPa2vaPacket->VirtualAddress);
          }
        } else {
          ShowErrorMessage(Va2paPa2vaPacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_VA2PA_AND_PA2VA_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS:
        ListOrModifyBreakpointPacket =
            (DEBUGGEE_BP_LIST_OR_MODIFY_PACKET *)(((CHAR *)TheActualPacket) +
                                                  sizeof(
                                                      DEBUGGER_REMOTE_PACKET));
        if (ListOrModifyBreakpointPacket->Result ==
            DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
        } else {
          ShowErrorMessage(ListOrModifyBreakpointPacket->Result);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_LIST_OR_MODIFY_BREAKPOINTS);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_UPDATE_SYMBOL_INFO:
        SymbolUpdatePacket =
            (DEBUGGER_UPDATE_SYMBOL_TABLE *)(((CHAR *)TheActualPacket) +
                                             sizeof(DEBUGGER_REMOTE_PACKET));
        SymbolBuildAndUpdateSymbolTable(
            &SymbolUpdatePacket->SymbolDetailPacket);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PCITREE:
        PcitreePacket =
            (DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET));
        if (PcitreePacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          ShowMessages("%-12s | %-9s | %-17s | %s \n%s\n", "DBDF", "VID:DID",
                       "Vendor Name", "Device Name",
                       "-------------------------------------------------------"
                       "---------------");
          for (UINT8 i = 0; i < (PcitreePacket->DeviceInfoListNum < DEV_MAX_NUM
                                     ? PcitreePacket->DeviceInfoListNum
                                     : DEV_MAX_NUM);
               i++) {
            Vendor *CurrentVendor = GetVendorById(
                PcitreePacket->DeviceInfoList[i].ConfigSpace.VendorId);
            char *CurrentVendorName = (char *)"N/A";
            char *CurrentDeviceName = (char *)"N/A";
            if (CurrentVendor != NULL) {
              CurrentVendorName = CurrentVendor->VendorName;
              Device *CurrentDevice = GetDeviceFromVendor(
                  CurrentVendor,
                  PcitreePacket->DeviceInfoList[i].ConfigSpace.DeviceId);
              if (CurrentDevice != NULL) {
                CurrentDeviceName = CurrentDevice->DeviceName;
              }
            }
            ShowMessages("%04x:%02x:%02x:%x | %04x:%04x | %-17.*s | %.*s\n", 0,
                         PcitreePacket->DeviceInfoList[i].Bus,
                         PcitreePacket->DeviceInfoList[i].Device,
                         PcitreePacket->DeviceInfoList[i].Function,
                         PcitreePacket->DeviceInfoList[i].ConfigSpace.VendorId,
                         PcitreePacket->DeviceInfoList[i].ConfigSpace.DeviceId,
                         strnlen_s(CurrentVendorName, PCI_NAME_STR_LENGTH),
                         CurrentVendorName,
                         strnlen_s(CurrentDeviceName, PCI_NAME_STR_LENGTH),
                         CurrentDeviceName);
            FreeVendor(CurrentVendor);
          }
          FreePciIdDatabase();
        } else {
          ShowErrorMessage(PcitreePacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PCITREE_RESULT);
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PCIDEVINFO:
        PcidevinfoPacket =
            (DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET
                 *)(((CHAR *)TheActualPacket) + sizeof(DEBUGGER_REMOTE_PACKET));
        if (PcidevinfoPacket->KernelStatus ==
            DEBUGGER_OPERATION_WAS_SUCCESSFUL) {
          const char *PciHeaderTypeAsString[] = {
              "Endpoint", "PCI-to-PCI Bridge", "PCI-to-CardBus Bridge"};
          const char *PciMmioBarTypeAsString[] = {"32-bit Wide", "Reserved",
                                                  "64-bit Wide", "Reserved"};
          UINT8 BarNumOffset = 0;
          ShowMessages(
              "PCI configuration space (CAM) for device %04x:%02x:%02x:%x\n", 0,
              PcidevinfoPacket->DeviceInfo.Bus,
              PcidevinfoPacket->DeviceInfo.Device,
              PcidevinfoPacket->DeviceInfo.Function);
          if (!PcidevinfoPacket->PrintRaw) {
            Vendor *CurrentVendor = GetVendorById(
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.VendorId);
            CHAR *CurrentVendorName = (CHAR *)"N/A";
            CHAR *CurrentDeviceName = (CHAR *)"N/A";
            if (CurrentVendor != NULL) {
              CurrentVendorName = CurrentVendor->VendorName;
              Device *CurrentDevice = GetDeviceFromVendor(
                  CurrentVendor, PcidevinfoPacket->DeviceInfo.ConfigSpace
                                     .CommonHeader.DeviceId);
              if (CurrentDevice != NULL) {
                CurrentDeviceName = CurrentDevice->DeviceName;
              }
            }
            ShowMessages(
                "\nCommon Header:\nVID:DID: %04x:%04x\nVendor Name: "
                "%-17.*s\nDevice Name: %.*s\nCommand: %04x\n",
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.VendorId,
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.DeviceId,
                strnlen_s(CurrentVendorName, PCI_NAME_STR_LENGTH),
                CurrentVendorName,
                strnlen_s(CurrentDeviceName, PCI_NAME_STR_LENGTH),
                CurrentDeviceName,
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.Command);
            if ((PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                     .HeaderType &
                 0x01)
                    << 7 ==
                0) {
              ShowMessages("  Memory Space: %u\n  I/O Space: %u\n",
                           (PcidevinfoPacket->DeviceInfo.ConfigSpace
                                .CommonHeader.Command &
                            0x2) >>
                               1,
                           (PcidevinfoPacket->DeviceInfo.ConfigSpace
                                .CommonHeader.Command &
                            0x1));
            }
            ShowMessages(
                "Status: %04x\nRevision ID: %02x\nClass Code: "
                "%06x\nCacheLineSize: %02x\nPrimaryLatencyTimer: "
                "%02x\nHeaderType: %s (%02x)\n  Multi-function Device: "
                "%s\nBist: %02x\n",
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.Status,
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                    .RevisionId,
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.ClassCode,
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                    .CacheLineSize,
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                    .PrimaryLatencyTimer,
                (PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                     .HeaderType &
                 0x3f) < 2
                    ? PciHeaderTypeAsString[(
                          PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                              .HeaderType &
                          0x1)]
                    : "Unknown",
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                    .HeaderType,
                (PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                     .HeaderType &
                 0x1)
                    ? "True"
                    : "False",
                PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader.Bist);
            FreeVendor(CurrentVendor);
            FreePciIdDatabase();
            ShowMessages("\nDevice Header:\n");
            if ((PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                     .HeaderType &
                 0x01)
                    << 7 ==
                0) {
              for (UINT8 i = 0; i < 5; i++) {
                if ((PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                         .ConfigSpaceEp.Bar[i] &
                     0x1) == 0) {
                  if (((PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                            .ConfigSpaceEp.Bar[i] &
                        0x6) >>
                       1) == 2) {
                    UINT64 BarMsb = PcidevinfoPacket->DeviceInfo.ConfigSpace
                                        .DeviceHeader.ConfigSpaceEp.Bar[i + 1];
                    UINT64 BarLsb = PcidevinfoPacket->DeviceInfo.ConfigSpace
                                        .DeviceHeader.ConfigSpaceEp.Bar[i];
                    UINT64 ActualBar =
                        ((BarMsb & 0xFFFFFFFF) << 32) + (BarLsb & 0xFFFFFFF0);
                    ShowMessages(
                        "BAR%u %s\n BAR Type: MMIO\n MMIO BAR Type: %s "
                        "(%02x)\n BAR MSB: %08x\n BAR LSB: %08x\n BAR "
                        "(actual): %016llx\n Prefetchable: %s\n",
                        i - BarNumOffset,
                        ((PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                              .Command &
                          0x2) >>
                             1 ==
                         0) || !PcidevinfoPacket->DeviceInfo.MmioBarInfo[i]
                                    .IsEnabled
                            ? "[disabled]"
                            : "",
                        PciMmioBarTypeAsString[(PcidevinfoPacket->DeviceInfo
                                                    .ConfigSpace.DeviceHeader
                                                    .ConfigSpaceEp.Bar[i] &
                                                0x6) >>
                                               1],
                        (PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                             .ConfigSpaceEp.Bar[i] &
                         0x6) >>
                            1,
                        BarMsb, BarLsb, ActualBar,
                        (PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                             .ConfigSpaceEp.Bar[i] &
                         0x8 >> 3)
                            ? "True"
                            : "False");
                    i++;
                    BarNumOffset++;
                  } else {
                    UINT32 ActualBar = (PcidevinfoPacket->DeviceInfo.ConfigSpace
                                            .DeviceHeader.ConfigSpaceEp.Bar[i] &
                                        0xFFFFFFF0);
                    ShowMessages(
                        "BAR%u %s\n BAR Type: MMIO\n BAR: %08x\n BAR (actual): "
                        "%08x\n Prefetchable: %s\n",
                        i - BarNumOffset,
                        ((PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                              .Command &
                          0x2) >>
                             1 ==
                         0) || !PcidevinfoPacket->DeviceInfo.MmioBarInfo[i]
                                    .IsEnabled
                            ? "[disabled]"
                            : "",
                        PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                            .ConfigSpaceEp.Bar[i],
                        ActualBar,
                        (PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                             .ConfigSpaceEp.Bar[i] &
                         0x8 >> 3)
                            ? "True"
                            : "False");
                  }
                } else {
                  UINT32 ActualBar32 = PcidevinfoPacket->DeviceInfo.ConfigSpace
                                           .DeviceHeader.ConfigSpaceEp.Bar[i] &
                                       0xFFFFFFFC;
                  ShowMessages(
                      "BAR%u %s\n BAR Type: Port IO\n BAR: %08x\n BAR "
                      "(actual): %08x\n Reserved: %u\n",
                      i - BarNumOffset,
                      ((PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                            .Command &
                        0x1) == 0)
                          ? "[disabled]"
                          : "",
                      PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                          .ConfigSpaceEp.Bar[i],
                      ActualBar32,
                      (PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                           .ConfigSpaceEp.Bar[i] &
                       0x2) >>
                          1);
                }
              }
              ShowMessages(
                  "Cardbus CIS Pointer: %08x\nSubsystem Vendor ID: "
                  "%04x\nSubsystem ID: %04x\nROM BAR: %08x\nCapabilities "
                  "Pointer: %02x\nReserved (0xD): %06x\nReserved (0xE): "
                  "%08x\nInterrupt Line: %02x\nInterrupt Pin: %02x\nMin Grant: "
                  "%02x\nMax latency: %02x\n",
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.CardBusCISPtr,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.SubVendorId,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.SubSystemId,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.ROMBar,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.CapabilitiesPtr,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.Reserved,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.Reserved1,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.InterruptLine,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.InterruptPin,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.MinGnt,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpaceEp.MaxLat);
            } else if ((PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                            .HeaderType &
                        0x3f) == 1) {
              ShowMessages("BAR0: %08x\nBAR1: %08x\n",
                           PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                               .ConfigSpacePtpBridge.Bar[0],
                           PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                               .ConfigSpacePtpBridge.Bar[1]);
              ShowMessages(
                  "Primary Bus Number: %02x\nSecondary Bus Number: "
                  "%02x\nSubordinate Bus Number: %02x\nSecondary Latency "
                  "Timer: %02x\nI/O Base: %02x\nI/O Limit: %02x\nSecondary "
                  "Status: %04x\nMemory Base: %04x\nMemory Limit: "
                  "%04x\nPrefetchable Memory Base: %04x\nPrefetchable Memory "
                  "Limit: %04x\nPrefetchable Base Upper 32 Bits: "
                  "%08x\nPrefetchable Limit Upper 32 Bits: %08x\nI/O Base "
                  "Upper 16 Bits: %04x\nI/O Limit Upper 16 Bits: "
                  "%04x\nCapability Pointer: %02x\nReserved: %06x\nROM BAR: "
                  "%08x\nInterrupt Line: %02x\nInterrupt Pin: %02x\nBridge "
                  "Control: %04x\n",
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.PrimaryBusNumber,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.SecondaryBusNumber,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.SubordinateBusNumber,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.SecondaryLatencyTimer,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.IoBase,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.IoLimit,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.SecondaryStatus,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.MemoryBase,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.MemoryLimit,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.PrefetchableMemoryBase,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.PrefetchableMemoryLimit,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.PrefetchableBaseUpper32b,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.PrefetchableLimitUpper32b,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.IoLimitUpper16b,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.IoBaseUpper16b,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.CapabilityPtr,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.Reserved,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.ROMBar,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.InterruptLine,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.InterruptPin,
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.DeviceHeader
                      .ConfigSpacePtpBridge.BridgeControl);
            } else if ((PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                            .HeaderType &
                        0x3f) == 2) {
              ShowMessages(
                  "Parsing header type %s (%02x) currently unsupported\n",
                  PciHeaderTypeAsString[PcidevinfoPacket->DeviceInfo.ConfigSpace
                                            .CommonHeader.HeaderType &
                                        0x01],
                  PcidevinfoPacket->DeviceInfo.ConfigSpace.CommonHeader
                          .HeaderType &
                      0x01);
            } else {
              ShowMessages("\nDevice Header:\nUnknown header type %02x\n",
                           (PcidevinfoPacket->DeviceInfo.ConfigSpace
                                .CommonHeader.HeaderType &
                            0x3f));
            }
          } else {
            UINT32 *cs = (UINT32 *)&PcidevinfoPacket->DeviceInfo.ConfigSpace;
            ShowMessages(
                "    00 01 02 03 04 05 06 07 08 09 0a 0b 0c 0d 0e 0f\n");
            for (UINT16 i = 0; i < CAM_CONFIG_SPACE_LENGTH; i += 16) {
              ShowMessages("%02x: ", i);
              for (UINT8 j = 0; j < 16; j++) {
                ShowMessages("%02x ", *(((BYTE *)cs) + j));
              }
              for (UINT8 j = 0; j < 16; j++) {
                char c = (char)*(cs + j);
                if (c >= 32 && c <= 126) {
                  ShowMessages("%c", c);
                } else {
                  ShowMessages(".");
                }
              }
              ShowMessages("\n");
              cs += 4;
            }
          }
        } else {
          ShowErrorMessage(PcidevinfoPacket->KernelStatus);
        }
        DbgReceivedKernelResponse(
            DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PCIDEVINFO_RESULT);
        break;
      default:
        ShowMessages("err, unknown packet action received from the debugger\n");
        break;
    }
  } else {
    ShowMessages("err, invalid packet received\n");
  }
  goto StartAgain;
  return TRUE;
}

BOOLEAN ListeningSerialPortInDebuggee() {
StartAgain:
  BOOL Status;
  char SerialBuffer[MaxSerialPacketSize] = {0};
  DWORD EventMask = 0;
  char ReadData = NULL;
  DWORD NoBytesRead = 0;
  UINT32 Loop = 0;
  PDEBUGGER_REMOTE_PACKET TheActualPacket =
      (PDEBUGGER_REMOTE_PACKET)SerialBuffer;
  Status = SetCommMask(g_SerialRemoteComPortHandle, EV_RXCHAR);
  if (Status == FALSE) {
  }
  Status = WaitCommEvent(g_SerialRemoteComPortHandle, &EventMask, NULL);
  if (Status == FALSE) {
  }
  do {
    Status = ReadFile(g_SerialRemoteComPortHandle, &ReadData, sizeof(ReadData),
                      &NoBytesRead, NULL);
    if (!Status || !(MaxSerialPacketSize > Loop)) {
      ShowMessages(
          "err, a buffer received in debuggee which exceeds the "
          "buffer limitation\n");
      goto StartAgain;
    }
    SerialBuffer[Loop] = ReadData;
    if (KdCheckForTheEndOfTheBuffer(&Loop, (BYTE *)SerialBuffer)) {
      break;
    }
    ++Loop;
  } while (NoBytesRead > 0);
  if (Loop == 1 && SerialBuffer[0] == NULL) {
    goto StartAgain;
  }
  if (TheActualPacket->Indicator == INDICATOR_OF_HYPERDBG_PACKET) {
    if (KdComputeDataChecksum((PVOID)&TheActualPacket->Indicator,
                              Loop - sizeof(BYTE)) !=
        TheActualPacket->Checksum) {
      ShowMessages("err checksum is invalid\n");
      goto StartAgain;
    }
    if (TheActualPacket->TypeOfThePacket !=
        DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE) {
      ShowMessages("err, unknown packet received from the debugger\n");
      goto StartAgain;
    }
    switch (TheActualPacket->RequestedActionOfThePacket) {
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE:
        if (!DebuggerPauseDebuggee()) {
          ShowMessages(
              "err, debugger tries to pause the debuggee but the "
              "attempt was unsuccessful\n");
        }
        break;
      case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DO_NOT_READ_ANY_PACKET:
        return TRUE;
        break;
      default:
        ShowMessages("err, unknown packet action received from the debugger\n");
        break;
    }
  } else {
    DebugBreak();
  }
  goto StartAgain;
  return TRUE;
}

DWORD WINAPI ListeningSerialPauseDebuggerThread(PVOID Param) {
  ListeningSerialPortInDebugger();
  return 0;
}

DWORD WINAPI ListeningSerialPauseDebuggeeThread(PVOID Param) {
  ListeningSerialPortInDebuggee();
  return 0;
}
