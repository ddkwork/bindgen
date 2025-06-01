
#include "pch.h"

NTSTATUS DriverEntry(PDRIVER_OBJECT DriverObject,
                     PUNICODE_STRING RegistryPath) {
  NTSTATUS Ntstatus = STATUS_SUCCESS;
  UINT64 Index = 0;
  PDEVICE_OBJECT DeviceObject = NULL;
  UNICODE_STRING DriverName =
      RTL_CONSTANT_STRING(L"\\Device\\HyperDbgReversingMachineDevice");
  UNICODE_STRING DosDeviceName =
      RTL_CONSTANT_STRING(L"\\DosDevices\\HyperDbgReversingMachineDevice");
  UNREFERENCED_PARAMETER(RegistryPath);
  UNREFERENCED_PARAMETER(DriverObject);
  ExInitializeDriverRuntime(DrvRtPoolNxOptIn);
  Ntstatus = IoCreateDevice(DriverObject, 0, &DriverName, FILE_DEVICE_UNKNOWN,
                            FILE_DEVICE_SECURE_OPEN, FALSE, &DeviceObject);
  if (Ntstatus == STATUS_SUCCESS) {
    for (Index = 0; Index < IRP_MJ_MAXIMUM_FUNCTION; Index++)
      DriverObject->MajorFunction[Index] = DrvUnsupported;
    DbgPrint("Setting device major functions");
    DriverObject->MajorFunction[IRP_MJ_CLOSE] = DrvClose;
    DriverObject->MajorFunction[IRP_MJ_CREATE] = DrvCreate;
    DriverObject->MajorFunction[IRP_MJ_READ] = DrvRead;
    DriverObject->MajorFunction[IRP_MJ_WRITE] = DrvWrite;
    DriverObject->MajorFunction[IRP_MJ_DEVICE_CONTROL] = DrvDispatchIoControl;
    DriverObject->DriverUnload = DrvUnload;
    IoCreateSymbolicLink(&DosDeviceName, &DriverName);
  }
  if (DeviceObject != NULL) {
    DeviceObject->Flags |= DO_BUFFERED_IO;
  }
  DbgPrint("HyperDbg's device and major functions are loaded");
  ASSERT(NT_SUCCESS(Ntstatus));
  return Ntstatus;
}

VOID DrvUnload(PDRIVER_OBJECT DriverObject) {
  UNICODE_STRING DosDeviceName;
  RtlInitUnicodeString(&DosDeviceName,
                       L"\\DosDevices\\HyperDbgReversingMachineDevice");
  IoDeleteSymbolicLink(&DosDeviceName);
  IoDeleteDevice(DriverObject->DeviceObject);
  LoaderUninitializeLogTracer();
}

NTSTATUS DrvCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
  UNREFERENCED_PARAMETER(DeviceObject);
  LUID DebugPrivilege = {SE_DEBUG_PRIVILEGE, 0};
  if (!SeSinglePrivilegeCheck(DebugPrivilege, Irp->RequestorMode)) {
    Irp->IoStatus.Status = STATUS_ACCESS_DENIED;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_ACCESS_DENIED;
  }
  if (g_HandleInUse) {
    Irp->IoStatus.Status = STATUS_SUCCESS;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_SUCCESS;
  }
  if (LoaderInitVmmAndReversingMachine()) {
    Irp->IoStatus.Status = STATUS_SUCCESS;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_SUCCESS;
  } else {
    Irp->IoStatus.Status = STATUS_UNSUCCESSFUL;
    Irp->IoStatus.Information = 0;
    IoCompleteRequest(Irp, IO_NO_INCREMENT);
    return STATUS_UNSUCCESSFUL;
  }
}

NTSTATUS DrvRead(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
  UNREFERENCED_PARAMETER(DeviceObject);
  DbgPrint("This function is not used");
  Irp->IoStatus.Status = STATUS_SUCCESS;
  Irp->IoStatus.Information = 0;
  IoCompleteRequest(Irp, IO_NO_INCREMENT);
  return STATUS_SUCCESS;
}

NTSTATUS DrvWrite(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
  UNREFERENCED_PARAMETER(DeviceObject);
  DbgPrint("This function is not used");
  Irp->IoStatus.Status = STATUS_SUCCESS;
  Irp->IoStatus.Information = 0;
  IoCompleteRequest(Irp, IO_NO_INCREMENT);
  return STATUS_SUCCESS;
}

NTSTATUS DrvClose(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
  UNREFERENCED_PARAMETER(DeviceObject);
  g_HandleInUse = FALSE;
  Irp->IoStatus.Status = STATUS_SUCCESS;
  Irp->IoStatus.Information = 0;
  IoCompleteRequest(Irp, IO_NO_INCREMENT);
  return STATUS_SUCCESS;
}

NTSTATUS DrvUnsupported(PDEVICE_OBJECT DeviceObject, PIRP Irp) {
  UNREFERENCED_PARAMETER(DeviceObject);
  DbgPrint("This function is not supported");
  Irp->IoStatus.Status = STATUS_SUCCESS;
  Irp->IoStatus.Information = 0;
  IoCompleteRequest(Irp, IO_NO_INCREMENT);
  return STATUS_SUCCESS;
}
