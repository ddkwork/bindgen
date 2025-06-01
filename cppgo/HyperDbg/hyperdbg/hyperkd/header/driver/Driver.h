
#pragma once

NTSTATUS DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath);
VOID DrvUnload(PDRIVER_OBJECT DriverObject);

NTSTATUS DrvCreate(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS DrvRead(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS DrvWrite(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS DrvClose(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS DrvUnsupported(PDEVICE_OBJECT DeviceObject, PIRP Irp);
NTSTATUS DrvDispatchIoControl(PDEVICE_OBJECT DeviceObject, PIRP Irp);
