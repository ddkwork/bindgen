#pragma once
#include <ntifs.h>
#include "Buffer.h"

struct Driver {
    inline static Buffer         buf = Buffer::New(1024);
    inline static PDRIVER_OBJECT DriverObject = nullptr;

    static NTSTATUS Run(PDRIVER_OBJECT driverObject, PUNICODE_STRING registry_path) {
        DriverObject = driverObject;
        static UNICODE_STRING device = RTL_CONSTANT_STRING(L"\\Device\\Driver");
        static UNICODE_STRING symlink = RTL_CONSTANT_STRING(L"\\DosDevices\\Driver");

        PDEVICE_OBJECT deviceObject;
        auto           status = IoCreateDevice(driverObject, 0, &device, FILE_DEVICE_UNKNOWN, 0, FALSE, &deviceObject);

        if (!NT_SUCCESS(status)) return status;
        status = IoCreateSymbolicLink(&symlink, &device);
        if (!NT_SUCCESS(status)) {
            IoDeleteDevice(deviceObject);
            return status;
        }
        for (UCHAR i = 0; i <= IRP_MJ_MAXIMUM_FUNCTION; i++) { driverObject->MajorFunction[i] = Dispatch; }
        driverObject->DriverUnload = Unload;
        return STATUS_SUCCESS;
    }

    static NTSTATUS Complete(PIRP irp, NTSTATUS status, ULONG_PTR info = 0) {
        irp->IoStatus.Status = status;
        irp->IoStatus.Information = info;
        IoCompleteRequest(irp, IO_NO_INCREMENT);
        return status;
    }

    static void Unload(PDRIVER_OBJECT) {
        static UNICODE_STRING symlink = RTL_CONSTANT_STRING(L"\\DosDevices\\Driver");
        IoDeleteSymbolicLink(&symlink);
        if (auto device = DriverObject->DeviceObject) { IoDeleteDevice(device); }
    }

    static NTSTATUS IoControl(PIRP irp) {
        auto stack = IoGetCurrentIrpStackLocation(irp);
        if (stack->Parameters.DeviceIoControl.IoControlCode == 0x801) {
            auto len = buf.Write(static_cast<const unsigned char *>(irp->AssociatedIrp.SystemBuffer), stack->Parameters.DeviceIoControl.InputBufferLength);
            return Complete(irp, STATUS_SUCCESS, len);
        }
        return Complete(irp, STATUS_NOT_IMPLEMENTED);
    }

    static NTSTATUS Dispatch(PDEVICE_OBJECT, PIRP irp) {
        auto stack = IoGetCurrentIrpStackLocation(irp);
        switch (stack->MajorFunction) {
            case IRP_MJ_DEVICE_CONTROL:
                return IoControl(irp);
            default:
                return Complete(irp, STATUS_SUCCESS);
        }
    }
};
