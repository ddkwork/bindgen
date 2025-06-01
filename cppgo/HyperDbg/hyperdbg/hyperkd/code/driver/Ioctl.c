
#include "pch.h"

NTSTATUS DrvDispatchIoControl(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
  UNREFERENCED_PARAMETER(DeviceObject);
  PIO_STACK_LOCATION IrpStack;
  PREGISTER_NOTIFY_BUFFER RegisterEventRequest;
  PDEBUGGER_READ_MEMORY DebuggerReadMemRequest;
  PDEBUGGER_READ_AND_WRITE_ON_MSR DebuggerReadOrWriteMsrRequest;
  PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE DebuggerHideAndUnhideRequest;
  PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS DebuggerPteRequest;
  PDEBUGGER_PAGE_IN_REQUEST DebuggerPageinRequest;
  PDEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET PcitreeRequest;
  PDEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET PcidevinfoRequest;
  PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS DebuggerVa2paAndPa2vaRequest;
  PDEBUGGER_EDIT_MEMORY DebuggerEditMemoryRequest;
  PDEBUGGER_SEARCH_MEMORY DebuggerSearchMemoryRequest;
  PDEBUGGER_GENERAL_EVENT_DETAIL DebuggerNewEventRequest;
  PDEBUGGER_MODIFY_EVENTS DebuggerModifyEventRequest;
  PDEBUGGER_FLUSH_LOGGING_BUFFERS DebuggerFlushBuffersRequest;
  PDEBUGGER_PREALLOC_COMMAND DebuggerReservePreallocPoolRequest;
  PDEBUGGER_PREACTIVATE_COMMAND DebuggerPreactivationRequest;
  PDEBUGGER_APIC_REQUEST DebuggerApicRequest;
  PINTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS DebuggerQueryIdtRequest;
  PDEBUGGER_UD_COMMAND_PACKET DebuggerUdCommandRequest;
  PUSERMODE_LOADED_MODULE_DETAILS DebuggerUsermodeModulesRequest;
  PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
  DebuggerUsermodeProcessOrThreadQueryRequest;
  PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET GetInformationProcessRequest;
  PREVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST RevServiceRequest;
  PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET GetInformationThreadRequest;
  PDEBUGGER_PERFORM_KERNEL_TESTS DebuggerKernelTestRequest;
  PDEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL
  DebuggerCommandExecutionFinishedRequest;
  PDEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER
  DebuggerSendUsermodeMessageRequest;
  PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
  DebuggerSendBufferFromDebuggeeToDebuggerRequest;
  PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS
  DebuggerAttachOrDetachToThreadRequest;
  PDEBUGGER_PREPARE_DEBUGGEE DebuggeeRequest;
  PDEBUGGER_PAUSE_PACKET_RECEIVED DebuggerPauseKernelRequest;
  PDEBUGGER_GENERAL_ACTION DebuggerNewActionRequest;
  PVOID BufferToStoreThreadsAndProcessesDetails;
  NTSTATUS Status;
  ULONG InBuffLength;
  ULONG OutBuffLength;
  SIZE_T ReturnSize;
  BOOLEAN DoNotChangeInformation = FALSE;
  PoolManagerCheckAndPerformAllocationAndDeallocation();
  if (g_AllowIOCTLFromUsermode) {
    IrpStack = IoGetCurrentIrpStackLocation(Irp);
    switch (IrpStack->Parameters.DeviceIoControl.IoControlCode) {
      case IOCTL_REGISTER_EVENT:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_REGISTER_EVENT ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        RegisterEventRequest =
            (PREGISTER_NOTIFY_BUFFER)Irp->AssociatedIrp.SystemBuffer;
        switch (RegisterEventRequest->Type) {
          case IRP_BASED:
            LogRegisterIrpBasedNotification((PVOID)Irp, &Status);
            break;
          case EVENT_BASED:
            if (LogRegisterEventBasedNotification((PVOID)Irp)) {
              Status = STATUS_SUCCESS;
            } else {
              Status = STATUS_UNSUCCESSFUL;
            }
            break;
          default:
            LogError("Err, unknown notification type from user-mode");
            Status = STATUS_INVALID_PARAMETER;
            break;
        }
        break;
      case IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL:
        g_AllowIOCTLFromUsermode = FALSE;
        LogCallbackSendBuffer(OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS, "$",
                              sizeof(CHAR), TRUE);
        Status = STATUS_SUCCESS;
        break;
      case IOCTL_TERMINATE_VMX:
        DebuggerUninitialize();
        VmFuncUninitVmm();
        Status = STATUS_SUCCESS;
        break;
      case IOCTL_DEBUGGER_READ_MEMORY:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_READ_MEMORY ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerReadMemRequest =
            (PDEBUGGER_READ_MEMORY)Irp->AssociatedIrp.SystemBuffer;
        if (DebuggerCommandReadMemory(
                DebuggerReadMemRequest,
                ((CHAR *)DebuggerReadMemRequest) + SIZEOF_DEBUGGER_READ_MEMORY,
                &ReturnSize) == TRUE) {
          Irp->IoStatus.Information = ReturnSize + SIZEOF_DEBUGGER_READ_MEMORY;
        } else {
          Irp->IoStatus.Information = SIZEOF_DEBUGGER_READ_MEMORY;
        }
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_READ_OR_WRITE_MSR:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerReadOrWriteMsrRequest =
            (PDEBUGGER_READ_AND_WRITE_ON_MSR)Irp->AssociatedIrp.SystemBuffer;
        if (DebuggerReadOrWriteMsrRequest->ActionType != DEBUGGER_MSR_WRITE) {
          if (!OutBuffLength) {
            Status = STATUS_INVALID_PARAMETER;
            break;
          }
        }
        Status = DebuggerReadOrWriteMsr(DebuggerReadOrWriteMsrRequest,
                                        (UINT64 *)DebuggerReadOrWriteMsrRequest,
                                        &ReturnSize);
        if (Status == STATUS_SUCCESS) {
          Irp->IoStatus.Information = ReturnSize;
          DoNotChangeInformation = TRUE;
        }
        break;
      case IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerPteRequest = (PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS)
                                 Irp->AssociatedIrp.SystemBuffer;
        ExtensionCommandPte(DebuggerPteRequest, FALSE);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_REGISTER_EVENT:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                sizeof(DEBUGGER_GENERAL_EVENT_DETAIL) ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerNewEventRequest =
            (PDEBUGGER_GENERAL_EVENT_DETAIL)Irp->AssociatedIrp.SystemBuffer;
        DebuggerParseEvent(
            DebuggerNewEventRequest,
            (PDEBUGGER_EVENT_AND_ACTION_RESULT)Irp->AssociatedIrp.SystemBuffer,
            FALSE);
        Irp->IoStatus.Information = sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT);
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                sizeof(DEBUGGER_GENERAL_ACTION) ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerNewActionRequest =
            (PDEBUGGER_GENERAL_ACTION)Irp->AssociatedIrp.SystemBuffer;
        DebuggerParseAction(
            DebuggerNewActionRequest,
            (PDEBUGGER_EVENT_AND_ACTION_RESULT)Irp->AssociatedIrp.SystemBuffer,
            FALSE);
        Irp->IoStatus.Information = sizeof(DEBUGGER_EVENT_AND_ACTION_RESULT);
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerHideAndUnhideRequest =
            (PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE)
                Irp->AssociatedIrp.SystemBuffer;
        if (DebuggerHideAndUnhideRequest->IsHide == TRUE) {
          TransparentHideDebugger(DebuggerHideAndUnhideRequest);
        } else {
          TransparentUnhideDebugger(DebuggerHideAndUnhideRequest);
        }
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerVa2paAndPa2vaRequest =
            (PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS)Irp->AssociatedIrp.SystemBuffer;
        ExtensionCommandVa2paAndPa2va(DebuggerVa2paAndPa2vaRequest, FALSE);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_EDIT_MEMORY:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_EDIT_MEMORY ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerEditMemoryRequest =
            (PDEBUGGER_EDIT_MEMORY)Irp->AssociatedIrp.SystemBuffer;
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength !=
            SIZEOF_DEBUGGER_EDIT_MEMORY +
                DebuggerEditMemoryRequest->CountOf64Chunks * sizeof(UINT64)) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerCommandEditMemory(DebuggerEditMemoryRequest);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_EDIT_MEMORY;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_SEARCH_MEMORY:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_SEARCH_MEMORY ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength ||
            OutBuffLength < MaximumSearchResults * sizeof(UINT64)) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerSearchMemoryRequest =
            (PDEBUGGER_SEARCH_MEMORY)Irp->AssociatedIrp.SystemBuffer;
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength !=
            SIZEOF_DEBUGGER_SEARCH_MEMORY +
                DebuggerSearchMemoryRequest->CountOf64Chunks * sizeof(UINT64)) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        if (DebuggerCommandSearchMemory(DebuggerSearchMemoryRequest) !=
            STATUS_SUCCESS) {
          RtlZeroMemory(DebuggerSearchMemoryRequest,
                        MaximumSearchResults * sizeof(UINT64));
        }
        Irp->IoStatus.Information = MaximumSearchResults * sizeof(UINT64);
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_MODIFY_EVENTS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                sizeof(DEBUGGER_MODIFY_EVENTS) ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerModifyEventRequest =
            (PDEBUGGER_MODIFY_EVENTS)Irp->AssociatedIrp.SystemBuffer;
        DebuggerParseEventsModification(
            DebuggerModifyEventRequest, FALSE,
            EnableInstantEventMechanism ? g_KernelDebuggerState : FALSE);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_MODIFY_EVENTS;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerFlushBuffersRequest =
            (PDEBUGGER_FLUSH_LOGGING_BUFFERS)Irp->AssociatedIrp.SystemBuffer;
        DebuggerCommandFlush(DebuggerFlushBuffersRequest);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerAttachOrDetachToThreadRequest =
            (PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS)
                Irp->AssociatedIrp.SystemBuffer;
        AttachingTargetProcess(DebuggerAttachOrDetachToThreadRequest);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_PREPARE_DEBUGGEE:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_PREPARE_DEBUGGEE ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggeeRequest =
            (PDEBUGGER_PREPARE_DEBUGGEE)Irp->AssociatedIrp.SystemBuffer;
        SerialConnectionPrepare(DebuggeeRequest);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_PREPARE_DEBUGGEE;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_PAUSE_PACKET_RECEIVED:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerPauseKernelRequest =
            (PDEBUGGER_PAUSE_PACKET_RECEIVED)Irp->AssociatedIrp.SystemBuffer;
        KdHaltSystem(DebuggerPauseKernelRequest);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerCommandExecutionFinishedRequest =
            (PDEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL)
                Irp->AssociatedIrp.SystemBuffer;
        DebuggerCommandSignalExecutionState(
            DebuggerCommandExecutionFinishedRequest);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerSendUsermodeMessageRequest =
            (PDEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER)
                Irp->AssociatedIrp.SystemBuffer;
        if (DebuggerSendUsermodeMessageRequest->Length == NULL_ZERO ||
            IrpStack->Parameters.DeviceIoControl.InputBufferLength !=
                SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER +
                    DebuggerSendUsermodeMessageRequest->Length) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerCommandSendMessage(DebuggerSendUsermodeMessageRequest);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerSendBufferFromDebuggeeToDebuggerRequest =
            (PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER)
                Irp->AssociatedIrp.SystemBuffer;
        if (DebuggerSendBufferFromDebuggeeToDebuggerRequest->LengthOfBuffer ==
                NULL_ZERO ||
            IrpStack->Parameters.DeviceIoControl.InputBufferLength !=
                SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER +
                    DebuggerSendBufferFromDebuggeeToDebuggerRequest
                        ->LengthOfBuffer) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerCommandSendGeneralBufferToDebugger(
            DebuggerSendBufferFromDebuggeeToDebuggerRequest);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_PERFROM_KERNEL_SIDE_TESTS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerKernelTestRequest =
            (PDEBUGGER_PERFORM_KERNEL_TESTS)Irp->AssociatedIrp.SystemBuffer;
        TestKernelPerformTests(DebuggerKernelTestRequest);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_RESERVE_PRE_ALLOCATED_POOLS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_PREALLOC_COMMAND ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerReservePreallocPoolRequest =
            (PDEBUGGER_PREALLOC_COMMAND)Irp->AssociatedIrp.SystemBuffer;
        DebuggerCommandReservePreallocatedPools(
            DebuggerReservePreallocPoolRequest);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_PREALLOC_COMMAND;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_PREACTIVATE_FUNCTIONALITY:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_PREACTIVATE_COMMAND ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerPreactivationRequest =
            (PDEBUGGER_PREACTIVATE_COMMAND)Irp->AssociatedIrp.SystemBuffer;
        DebuggerCommandPreactivateFunctionality(DebuggerPreactivationRequest);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_PREACTIVATE_COMMAND;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_PERFROM_ACTIONS_ON_APIC:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_APIC_REQUEST ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerApicRequest =
            (PDEBUGGER_APIC_REQUEST)Irp->AssociatedIrp.SystemBuffer;
        Irp->IoStatus.Information =
            ExtensionCommandPerformActionsForApicRequests(DebuggerApicRequest);
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_QUERY_IDT_ENTRY:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerQueryIdtRequest = (PINTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS)
                                      Irp->AssociatedIrp.SystemBuffer;
        ExtensionCommandPerformQueryIdtEntriesRequest(DebuggerQueryIdtRequest,
                                                      FALSE);
        Irp->IoStatus.Information =
            SIZEOF_INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_SEND_USER_DEBUGGER_COMMANDS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                sizeof(DEBUGGER_UD_COMMAND_PACKET) ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerUdCommandRequest =
            (PDEBUGGER_UD_COMMAND_PACKET)Irp->AssociatedIrp.SystemBuffer;
        UdDispatchUsermodeCommands(DebuggerUdCommandRequest);
        Irp->IoStatus.Information = sizeof(DEBUGGER_UD_COMMAND_PACKET);
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES:
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        BufferToStoreThreadsAndProcessesDetails =
            (PVOID)Irp->AssociatedIrp.SystemBuffer;
        AttachingQueryDetailsOfActiveDebuggingThreadsAndProcesses(
            BufferToStoreThreadsAndProcessesDetails, OutBuffLength);
        Irp->IoStatus.Information = OutBuffLength;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_GET_USER_MODE_MODULE_DETAILS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                sizeof(USERMODE_LOADED_MODULE_DETAILS) ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerUsermodeModulesRequest =
            (PUSERMODE_LOADED_MODULE_DETAILS)Irp->AssociatedIrp.SystemBuffer;
        UserAccessGetLoadedModules(DebuggerUsermodeModulesRequest,
                                   OutBuffLength);
        Irp->IoStatus.Information = OutBuffLength;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                sizeof(DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS) ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerUsermodeProcessOrThreadQueryRequest =
            (PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS)
                Irp->AssociatedIrp.SystemBuffer;
        if (DebuggerUsermodeProcessOrThreadQueryRequest->QueryType ==
            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_COUNT) {
          ProcessQueryCount(DebuggerUsermodeProcessOrThreadQueryRequest);
        } else if (
            DebuggerUsermodeProcessOrThreadQueryRequest->QueryType ==
            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_COUNT) {
          ThreadQueryCount(DebuggerUsermodeProcessOrThreadQueryRequest);
        }
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                sizeof(DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS) ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerUsermodeProcessOrThreadQueryRequest =
            (PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS)
                Irp->AssociatedIrp.SystemBuffer;
        if (DebuggerUsermodeProcessOrThreadQueryRequest->QueryType ==
            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_LIST) {
          ProcessQueryList(DebuggerUsermodeProcessOrThreadQueryRequest,
                           DebuggerUsermodeProcessOrThreadQueryRequest,
                           OutBuffLength);
        } else if (
            DebuggerUsermodeProcessOrThreadQueryRequest->QueryType ==
            DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_LIST) {
          ThreadQueryList(DebuggerUsermodeProcessOrThreadQueryRequest,
                          DebuggerUsermodeProcessOrThreadQueryRequest,
                          OutBuffLength);
        }
        Irp->IoStatus.Information = OutBuffLength;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_QUERY_CURRENT_THREAD:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        GetInformationThreadRequest =
            (PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET)
                Irp->AssociatedIrp.SystemBuffer;
        ThreadQueryDetails(GetInformationThreadRequest);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_QUERY_CURRENT_PROCESS:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        GetInformationProcessRequest =
            (PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET)
                Irp->AssociatedIrp.SystemBuffer;
        ProcessQueryDetails(GetInformationProcessRequest);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_REQUEST_REV_MACHINE_SERVICE:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength || !OutBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        RevServiceRequest = (PREVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST)
                                Irp->AssociatedIrp.SystemBuffer;
        ConfigureInitializeExecTrapOnAllProcessors();
        Irp->IoStatus.Information =
            SIZEOF_REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_DEBUGGER_BRING_PAGES_IN:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGER_PAGE_IN_REQUEST ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        DebuggerPageinRequest =
            (PDEBUGGER_PAGE_IN_REQUEST)Irp->AssociatedIrp.SystemBuffer;
        DebuggerCommandBringPagein(DebuggerPageinRequest);
        Irp->IoStatus.Information = SIZEOF_DEBUGGER_PAGE_IN_REQUEST;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_PCIE_ENDPOINT_ENUM:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        PcitreeRequest = (PDEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET)
                             Irp->AssociatedIrp.SystemBuffer;
        ExtensionCommandPcitree(PcitreeRequest, FALSE);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      case IOCTL_PCIDEVINFO_ENUM:
        if (IrpStack->Parameters.DeviceIoControl.InputBufferLength <
                SIZEOF_DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET ||
            Irp->AssociatedIrp.SystemBuffer == NULL) {
          Status = STATUS_INVALID_PARAMETER;
          LogError("Err, invalid parameter to IOCTL dispatcher");
          break;
        }
        InBuffLength = IrpStack->Parameters.DeviceIoControl.InputBufferLength;
        OutBuffLength = IrpStack->Parameters.DeviceIoControl.OutputBufferLength;
        if (!InBuffLength) {
          Status = STATUS_INVALID_PARAMETER;
          break;
        }
        PcidevinfoRequest = (PDEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET)
                                Irp->AssociatedIrp.SystemBuffer;
        ExtensionCommandPcidevinfo(PcidevinfoRequest, FALSE);
        Irp->IoStatus.Information =
            SIZEOF_DEBUGGEE_PCIDEVINFO_REQUEST_RESPONSE_PACKET;
        Status = STATUS_SUCCESS;
        DoNotChangeInformation = TRUE;
        break;
      default:
        LogError("Err, unknown IOCTL");
        Status = STATUS_NOT_IMPLEMENTED;
        break;
    }
  } else {
    Status = STATUS_SUCCESS;
  }
  if (Status != STATUS_PENDING) {
    Irp->IoStatus.Status = Status;
    if (!DoNotChangeInformation) {
      Irp->IoStatus.Information = 0;
    }
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
  }
  return Status;
}
