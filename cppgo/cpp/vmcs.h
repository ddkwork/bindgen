#include <ntddk.h>
#include <intrin.h>

#pragma pack(push, 1)// 确保紧密内存布局

struct alignas(4096) VMCS {// 4096字节对齐
    u32 revision_id;
    u32 abort_indicator;

    u8 data[4096 - sizeof(u32) * 2];

    VMCS() : revision_id(0), abort_indicator(0) { RtlZeroMemory(&data, sizeof(data)); }

    static NTSTATUS Create(_Out_ VMCS **ppVmcs, _In_ u32 revisionId) {
        *ppVmcs = static_cast<VMCS *>(ExAllocatePool2(POOL_FLAG_NON_PAGED | POOL_FLAG_CACHE_ALIGNED, sizeof(VMCS), 'SCMV'));

        if (!*ppVmcs) return STATUS_INSUFFICIENT_RESOURCES;

        new (*ppVmcs) VMCS();

        (*ppVmcs)->revision_id = revisionId;
        return STATUS_SUCCESS;
    }

    NTSTATUS Write(u32 field, u64 value) const {
        if (KeGetCurrentIrql() > DISPATCH_LEVEL) {
            KdPrint(("VMCS::Write called at invalid IRQL: %d\n", KeGetCurrentIrql()));
            return STATUS_UNSUCCESSFUL;
        }

        u8 status = __vmx_vmwrite(field, value);
        return (status == 0) ? STATUS_SUCCESS : STATUS_UNSUCCESSFUL;
    }

    NTSTATUS Read(u32 field, _Out_ u64 *pValue) const {
        if (KeGetCurrentIrql() > DISPATCH_LEVEL) {
            KdPrint(("VMCS::Read called at invalid IRQL: %d\n", KeGetCurrentIrql()));
            return STATUS_UNSUCCESSFUL;
        }

        u8 status = __vmx_vmread(field, pValue);
        return (status == 0) ? STATUS_SUCCESS : STATUS_UNSUCCESSFUL;
    }

    NTSTATUS SetControlBit(u32 controlField, u32 flag) {
        u64      currentValue = 0;
        NTSTATUS status = Read(controlField, &currentValue);
        if (!NT_SUCCESS(status)) return status;

        return Write(controlField, currentValue | flag);
    }

    NTSTATUS ClearControlBit(u32 controlField, u32 flag) {
        u64      currentValue = 0;
        NTSTATUS status = Read(controlField, &currentValue);
        if (!NT_SUCCESS(status)) return status;

        return Write(controlField, currentValue & ~flag);
    }

    NTSTATUS EnterVmxOperation() const {
        PHYSICAL_ADDRESS phys = MmGetPhysicalAddress(const_cast<VMCS *>(this));

        u8 status = __vmx_on(&phys.QuadPart);
        return (status == 0) ? STATUS_SUCCESS : STATUS_UNSUCCESSFUL;
    }

    static NTSTATUS ExitVmxOperation() {
        u8 status = __vmx_off();
        return (status == 0) ? STATUS_SUCCESS : STATUS_UNSUCCESSFUL;
    }
};

#pragma pack(pop)

NTSTATUS InitializeVmcs(VMCS *vmcs) {
    NTSTATUS status = vmcs->Write(0x6C14, __readcr0());// HOST_CR0
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->Write(0x6C16, __readcr4());// HOST_CR4
    if (!NT_SUCCESS(status)) return status;

    u64 cr3;
    __movq(&cr3, cr3);
    status = vmcs->Write(0x6C02, cr3);// HOST_CR3
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->Write(0x6800, __readcr0());// GUEST_CR0
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->Write(0x6802, __readcr4());// GUEST_CR4
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->Write(0x6806, cr3);// GUEST_CR3
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->SetControlBit(0x4002, 0x80000000);// CPU_BASED_VM_EXEC_CONTROL
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->SetControlBit(0x4000, 0x3F);// PIN_BASED_VM_EXEC_CONTROL
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->SetControlBit(0x401E, 0xFFFF);// VM_EXIT_CONTROLS
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->SetControlBit(0x401C, 0xFFFF);// VM_ENTRY_CONTROLS
    if (!NT_SUCCESS(status)) return status;

    status = vmcs->Write(0x4402, static_cast<u64>(0x1));// EXCEPTION_BITMAP
    if (!NT_SUCCESS(status)) return status;

    return STATUS_SUCCESS;
}

NTSTATUS DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath) {
    UNREFERENCED_PARAMETER(RegistryPath);
    DriverObject->DriverUnload = [](PDRIVER_OBJECT DrvObj) { VMCS::ExitVmxOperation(); };

    const u32 vmxBasicMsr = __readmsr(0x480);// IA32_VMX_BASIC_MSR
    const u32 revisionId = vmxBasicMsr & 0x7FFFFFFF;

    VMCS    *vmcs = nullptr;
    NTSTATUS status = VMCS::Create(&vmcs, revisionId);
    if (!NT_SUCCESS(status)) {
        KdPrint(("Failed to create VMCS area: 0x%X\n", status));
        return status;
    }

    status = vmcs->EnterVmxOperation();
    if (!NT_SUCCESS(status)) {
        KdPrint(("VMXON failed: 0x%X\n", status));
        ExFreePool(vmcs);
        return status;
    }

    status = InitializeVmcs(vmcs);
    if (!NT_SUCCESS(status)) {
        KdPrint(("VMCS initialization failed: 0x%X\n", status));
        VMCS::ExitVmxOperation();
        ExFreePool(vmcs);
        return status;
    }

    u8 statusCode = __vmx_vmptrld(MmGetPhysicalAddress(vmcs).QuadPart);
    if (statusCode != 0) {
        KdPrint(("VMPTRLD failed: %d\n", statusCode));
        VMCS::ExitVmxOperation();
        ExFreePool(vmcs);
        return STATUS_UNSUCCESSFUL;
    }

    KdPrint(("VMCS initialized successfully!\n"));

    return STATUS_SUCCESS;
}
