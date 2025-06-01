
#include "pch.h"

NTSTATUS DrvDispatchIoControl(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
  PIO_STACK_LOCATION IrpStack;
  PREGISTER_NOTIFY_BUFFER RegisterEventRequest;
  NTSTATUS Status;
  UNREFERENCED_PARAMETER(DeviceObject);
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
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
  }
  return Status;
}
