#include <ntddk.h>
#include <intrin.h>

#pragma pack(push, 1)// 严格内存布局

constexpr uint64_t EPT_READ = 1ull << 0;
constexpr uint64_t EPT_WRITE = 1ull << 1;
constexpr uint64_t EPT_EXECUTE = 1ull << 2;
constexpr uint64_t EPT_MEMORY_TYPE_WB = 0x6ull << 3;// 回写内存类型
constexpr uint64_t EPT_IGNORE_PAT = 1ull << 6;
constexpr uint64_t EPT_ACCESSED = 1ull << 8;
constexpr uint64_t EPT_DIRTY = 1ull << 9;

constexpr uint64_t EPT_PDE_PS = 1ull << 7;  // 2MB页
constexpr uint64_t EPT_PDPTE_PS = 1ull << 7;// 1GB页

constexpr uint64_t PHYS_ADDR_MASK = 0x000FFFFFFFFFF000ull;

struct EptController {
    uint64_t        *pml4 = nullptr;
    PHYSICAL_ADDRESS pml4Phys;

    uint64_t *pdpt = nullptr;
    uint64_t *pd = nullptr;

    EptController() { Initialize(); }

    NTSTATUS Initialize() {
        pml4 = static_cast<uint64_t *>(ExAllocatePool2(POOL_FLAG_NON_PAGED | POOL_FLAG_CACHE_ALIGNED,
                                                       PAGE_SIZE,
                                                       'L4MP'// 'PML4'的tag
                                                       ));

        if (!pml4) return STATUS_INSUFFICIENT_RESOURCES;

        pml4Phys = MmGetPhysicalAddress(pml4);
        RtlZeroMemory(pml4, PAGE_SIZE);

        pdpt = static_cast<uint64_t *>(ExAllocatePool2(POOL_FLAG_NON_PAGED | POOL_FLAG_CACHE_ALIGNED,
                                                       PAGE_SIZE,
                                                       'TPDP'// 'PDPT'的tag
                                                       ));

        if (!pdpt) {
            ExFreePool(pml4);
            pml4 = nullptr;
            return STATUS_INSUFFICIENT_RESOURCES;
        }
        RtlZeroMemory(pdpt, PAGE_SIZE);

        pd = static_cast<uint64_t *>(ExAllocatePool2(POOL_FLAG_NON_PAGED | POOL_FLAG_CACHE_ALIGNED,
                                                     PAGE_SIZE,
                                                     'DP__'// 'PD'的tag
                                                     ));

        if (!pd) {
            ExFreePool(pdpt);
            ExFreePool(pml4);
            pdpt = nullptr;
            pml4 = nullptr;
            return STATUS_INSUFFICIENT_RESOURCES;
        }
        RtlZeroMemory(pd, PAGE_SIZE);

        PHYSICAL_ADDRESS pdptPhys = MmGetPhysicalAddress(pdpt);
        pml4[0] = pdptPhys.QuadPart | EPT_READ | EPT_WRITE | EPT_EXECUTE;

        PHYSICAL_ADDRESS pdPhys = MmGetPhysicalAddress(pd);
        pdpt[0] = pdPhys.QuadPart | EPT_READ | EPT_WRITE | EPT_EXECUTE;

        return STATUS_SUCCESS;
    }

    ~EptController() {
        if (pd) ExFreePool(pd);
        if (pdpt) ExFreePool(pdpt);
        if (pml4) ExFreePool(pml4);
        pd = nullptr;
        pdpt = nullptr;
        pml4 = nullptr;
    }

    uint64_t GetEptPointer() const {
        constexpr uint64_t EPT_POINTER_WB = 0x6ull;// 回写内存类型
        return (pml4Phys.QuadPart & PHYS_ADDR_MASK) | EPT_POINTER_WB;
    }

    NTSTATUS MapPage(uint64_t guestPhys, uint64_t hostPhys, uint64_t flags) {
        if (KeGetCurrentIrql() > DISPATCH_LEVEL) {
            DbgPrint("MapPage called at invalid IRQL: %d\n", KeGetCurrentIrql());
            return STATUS_UNSUCCESSFUL;
        }

        uint64_t pml4Index = (guestPhys >> 39) & 0x1FF;
        uint64_t pdptIndex = (guestPhys >> 30) & 0x1FF;
        uint64_t pdIndex = (guestPhys >> 21) & 0x1FF;
        uint64_t ptIndex = (guestPhys >> 12) & 0x1FF;

        if (pml4Index != 0) {
            DbgPrint("Only first PML4 supported\n");
            return STATUS_UNSUCCESSFUL;
        }
        uint64_t *pdptBase = pdpt;// 使用预分配的PDPT

        if (pdptIndex != 0) {
            DbgPrint("Only first PDPT supported\n");
            return STATUS_UNSUCCESSFUL;
        }
        uint64_t *pdBase = pd;// 使用预分配的PD

        uint64_t *pdEntry = &pdBase[pdIndex];
        uint64_t *ptTable = nullptr;

        if ((*pdEntry & 1) == 0) {
            ptTable = static_cast<uint64_t *>(ExAllocatePool2(POOL_FLAG_NON_PAGED,
                                                              PAGE_SIZE,
                                                              'TP__'// 'PT'的tag
                                                              ));

            if (!ptTable) return STATUS_INSUFFICIENT_RESOURCES;

            RtlZeroMemory(ptTable, PAGE_SIZE);
            PHYSICAL_ADDRESS ptPhys = MmGetPhysicalAddress(ptTable);

            *pdEntry = ptPhys.QuadPart | flags | EPT_READ | EPT_WRITE | EPT_EXECUTE;
        } else {
            PHYSICAL_ADDRESS ptPhys;
            ptPhys.QuadPart = *pdEntry & PHYS_ADDR_MASK;
            ptTable = static_cast<uint64_t *>(MmGetVirtualForPhysical(ptPhys));
        }

        uint64_t *ptEntry = &ptTable[ptIndex];
        if ((*ptEntry & 1) != 0) {
            DbgPrint("Page already mapped: %llx\n", guestPhys);
            return STATUS_ALREADY_REGISTERED;
        }

        *ptEntry = (hostPhys & PHYS_ADDR_MASK) | flags | EPT_READ | EPT_WRITE | EPT_EXECUTE;
        return STATUS_SUCCESS;
    }

    NTSTATUS MapLargePage(uint64_t guestPhys, uint64_t hostPhys, uint64_t flags) {
        if (KeGetCurrentIrql() > DISPATCH_LEVEL) {
            DbgPrint("MapLargePage called at invalid IRQL: %d\n", KeGetCurrentIrql());
            return STATUS_UNSUCCESSFUL;
        }

        uint64_t pml4Index = (guestPhys >> 39) & 0x1FF;
        uint64_t pdptIndex = (guestPhys >> 30) & 0x1FF;
        uint64_t pdIndex = (guestPhys >> 21) & 0x1FF;

        if ((guestPhys & 0x1FFFFF) || (hostPhys & 0x1FFFFF)) {
            DbgPrint("Address not 2MB aligned\n");
            return STATUS_INVALID_PARAMETER;
        }

        if (pml4Index != 0 || pdptIndex != 0) {
            DbgPrint("Only first PML4/PDPT supported\n");
            return STATUS_UNSUCCESSFUL;
        }

        uint64_t *pdEntry = &pd[pdIndex];
        if ((*pdEntry & 1) != 0) {
            DbgPrint("Large page already mapped: %llx\n", guestPhys);
            return STATUS_ALREADY_REGISTERED;
        }

        *pdEntry = (hostPhys & PHYS_ADDR_MASK) | flags | EPT_PDE_PS | EPT_READ | EPT_WRITE | EPT_EXECUTE;

        return STATUS_SUCCESS;
    }

    NTSTATUS SetupInVmcs() const {
        if (!pml4) return STATUS_UNSUCCESSFUL;

        uint64_t eptPtr = GetEptPointer();
        u8       status = __vmx_vmwrite(0x0000201A, eptPtr);// VMCS_CTRL_EPT_PTR

        return (status == 0) ? STATUS_SUCCESS : STATUS_UNSUCCESSFUL;
    }

    NTSTATUS EnableEptControls() {
        uint64_t execControls = 0;
        uint8_t  status = __vmx_vmread(0x00004002, &execControls);
        if (status != 0) return STATUS_UNSUCCESSFUL;

        execControls |= (1ull << 1);

        status = __vmx_vmwrite(0x00004002, execControls);
        return (status == 0) ? STATUS_SUCCESS : STATUS_UNSUCCESSFUL;
    }
};

#pragma pack(pop)

NTSTATUS SetupEptForVirtualization() {
    EptController ept;
    NTSTATUS      status = ept.Initialize();
    if (!NT_SUCCESS(status)) {
        DbgPrint("EPT initialization failed: 0x%X\n", status);
        return status;
    }

    uint64_t base = 0;
    for (; base < 0x400000; base += PAGE_SIZE) {
        status = ept.MapPage(base, base, EPT_MEMORY_TYPE_WB);
        if (!NT_SUCCESS(status)) {
            DbgPrint("Failed to map page at 0x%llX: 0x%X\n", base, status);
            break;
        }
    }

    status = ept.MapLargePage(0x400000, 0x400000, EPT_MEMORY_TYPE_WB);
    if (!NT_SUCCESS(status)) { DbgPrint("Failed to map large page: 0x%X\n", status); }

    status = ept.SetupInVmcs();
    if (!NT_SUCCESS(status)) {
        DbgPrint("Failed to set EPT in VMCS: 0x%X\n", status);
        return status;
    }

    status = ept.EnableEptControls();
    if (!NT_SUCCESS(status)) {
        DbgPrint("Failed to enable EPT controls: 0x%X\n", status);
        return status;
    }

    DbgPrint("EPT configured successfully!\n");
    return STATUS_SUCCESS;
}
