

#define HYPERV_CPUID_VENDOR_AND_MAX_FUNCTIONS 0x40000000
#define HYPERV_CPUID_INTERFACE 0x40000001
#define HYPERV_CPUID_VERSION 0x40000002
#define HYPERV_CPUID_FEATURES 0x40000003
#define HYPERV_CPUID_ENLIGHTMENT_INFO 0x40000004
#define HYPERV_CPUID_IMPLEMENT_LIMITS 0x40000005
#define HYPERV_CPUID_CPU_MANAGEMENT_FEATURES 0x40000007
#define HYPERV_CPUID_NESTED_FEATURES 0x4000000A
#define HYPERV_CPUID_ISOLATION_CONFIG 0x4000000C
#define HYPERV_CPUID_VIRT_STACK_INTERFACE 0x40000081
#define HYPERV_VS_INTERFACE_EAX_SIGNATURE 0x31235356
#define HYPERV_CPUID_VIRT_STACK_PROPERTIES 0x40000082

#define HYPERV_VS_PROPERTIES_EAX_EXTENDED_IOAPIC_RTE BIT(2)
#define HYPERV_HYPERVISOR_PRESENT_BIT 0x80000000
#define HYPERV_CPUID_MIN 0x40000005
#define HYPERV_CPUID_MAX 0x4000ffff

#define HV_X64_MWAIT_AVAILABLE BIT(0)

#define HV_X64_GUEST_DEBUGGING_AVAILABLE BIT(1)

#define HV_X64_PERF_MONITOR_AVAILABLE BIT(2)

#define HV_X64_CPU_DYNAMIC_PARTITIONING_AVAILABLE BIT(3)

#define HV_X64_HYPERCALL_XMM_INPUT_AVAILABLE BIT(4)

#define HV_X64_GUEST_IDLE_STATE_AVAILABLE BIT(5)

#define HV_FEATURE_FREQUENCY_MSRS_AVAILABLE BIT(8)

#define HV_FEATURE_GUEST_CRASH_MSR_AVAILABLE BIT(10)

#define HV_FEATURE_DEBUG_MSRS_AVAILABLE BIT(11)

#define HV_X64_HYPERCALL_XMM_OUTPUT_AVAILABLE BIT(15)

#define HV_STIMER_DIRECT_MODE_AVAILABLE BIT(19)

#define HV_X64_AS_SWITCH_RECOMMENDED BIT(0)

#define HV_X64_LOCAL_TLB_FLUSH_RECOMMENDED BIT(1)

#define HV_X64_REMOTE_TLB_FLUSH_RECOMMENDED BIT(2)

#define HV_X64_APIC_ACCESS_RECOMMENDED BIT(3)

#define HV_X64_SYSTEM_RESET_RECOMMENDED BIT(4)

#define HV_X64_RELAXED_TIMING_RECOMMENDED BIT(5)

#define HV_DEPRECATING_AEOI_RECOMMENDED BIT(9)

#define HV_X64_CLUSTER_IPI_RECOMMENDED BIT(10)

#define HV_X64_EX_PROCESSOR_MASKS_RECOMMENDED BIT(11)

#define HV_X64_ENLIGHTENED_VMCS_RECOMMENDED BIT(14)

#define HV_X64_START_LOGICAL_PROCESSOR BIT(0)
#define HV_X64_CREATE_ROOT_VIRTUAL_PROCESSOR BIT(1)
#define HV_X64_PERFORMANCE_COUNTER_SYNC BIT(2)
#define HV_X64_RESERVED_IDENTITY_BIT BIT(31)

#define HV_X64_NO_NONARCH_CORESHARING BIT(18)

#define HV_X64_NESTED_DIRECT_FLUSH BIT(17)
#define HV_X64_NESTED_GUEST_MAPPING_FLUSH BIT(18)
#define HV_X64_NESTED_MSR_BITMAP BIT(19)

#define HV_X64_NESTED_ENLIGHTENED_TLB BIT(22)

#define HV_PARAVISOR_PRESENT BIT(0)

#define HV_ISOLATION_TYPE GENMASK(3, 0)
#define HV_SHARED_GPA_BOUNDARY_ACTIVE BIT(5)
#define HV_SHARED_GPA_BOUNDARY_BITS GENMASK(11, 6)

enum hv_isolation_type {
  HV_ISOLATION_TYPE_NONE = 0,
  HV_ISOLATION_TYPE_VBS = 1,
  HV_ISOLATION_TYPE_SNP = 2
};

#define HV_X64_MSR_GUEST_OS_ID 0x40000000

#define HV_X64_MSR_HYPERCALL 0x40000001

#define HV_REGISTER_VP_INDEX 0x40000002

#define HV_X64_MSR_RESET 0x40000003

#define HV_X64_MSR_VP_RUNTIME 0x40000010

#define HV_REGISTER_TIME_REF_COUNT 0x40000020

#define HV_REGISTER_REFERENCE_TSC 0x40000021

#define HV_X64_MSR_TSC_FREQUENCY 0x40000022

#define HV_X64_MSR_APIC_FREQUENCY 0x40000023

#define HV_X64_MSR_EOI 0x40000070
#define HV_X64_MSR_ICR 0x40000071
#define HV_X64_MSR_TPR 0x40000072
#define HV_X64_MSR_VP_ASSIST_PAGE 0x40000073

#define HV_REGISTER_SCONTROL 0x40000080
#define HV_REGISTER_SVERSION 0x40000081
#define HV_REGISTER_SIEFP 0x40000082
#define HV_REGISTER_SIMP 0x40000083
#define HV_REGISTER_EOM 0x40000084
#define HV_REGISTER_SINT0 0x40000090
#define HV_REGISTER_SINT1 0x40000091
#define HV_REGISTER_SINT2 0x40000092
#define HV_REGISTER_SINT3 0x40000093
#define HV_REGISTER_SINT4 0x40000094
#define HV_REGISTER_SINT5 0x40000095
#define HV_REGISTER_SINT6 0x40000096
#define HV_REGISTER_SINT7 0x40000097
#define HV_REGISTER_SINT8 0x40000098
#define HV_REGISTER_SINT9 0x40000099
#define HV_REGISTER_SINT10 0x4000009A
#define HV_REGISTER_SINT11 0x4000009B
#define HV_REGISTER_SINT12 0x4000009C
#define HV_REGISTER_SINT13 0x4000009D
#define HV_REGISTER_SINT14 0x4000009E
#define HV_REGISTER_SINT15 0x4000009F

#define HV_REGISTER_STIMER0_CONFIG 0x400000B0
#define HV_REGISTER_STIMER0_COUNT 0x400000B1
#define HV_REGISTER_STIMER1_CONFIG 0x400000B2
#define HV_REGISTER_STIMER1_COUNT 0x400000B3
#define HV_REGISTER_STIMER2_CONFIG 0x400000B4
#define HV_REGISTER_STIMER2_COUNT 0x400000B5
#define HV_REGISTER_STIMER3_CONFIG 0x400000B6
#define HV_REGISTER_STIMER3_COUNT 0x400000B7

#define HV_X64_MSR_GUEST_IDLE 0x400000F0

#define HV_REGISTER_CRASH_P0 0x40000100
#define HV_REGISTER_CRASH_P1 0x40000101
#define HV_REGISTER_CRASH_P2 0x40000102
#define HV_REGISTER_CRASH_P3 0x40000103
#define HV_REGISTER_CRASH_P4 0x40000104
#define HV_REGISTER_CRASH_CTL 0x40000105

#define HV_X64_MSR_REENLIGHTENMENT_CONTROL 0x40000106
#define HV_X64_MSR_TSC_EMULATION_CONTROL 0x40000107
#define HV_X64_MSR_TSC_EMULATION_STATUS 0x40000108

#define HV_X64_MSR_TSC_INVARIANT_CONTROL 0x40000118

#define HV_X64_MSR_STIMER0_COUNT HV_REGISTER_STIMER0_COUNT
#define HV_X64_MSR_STIMER0_CONFIG HV_REGISTER_STIMER0_CONFIG
#define HV_X64_MSR_STIMER1_COUNT HV_REGISTER_STIMER1_COUNT
#define HV_X64_MSR_STIMER1_CONFIG HV_REGISTER_STIMER1_CONFIG
#define HV_X64_MSR_STIMER2_COUNT HV_REGISTER_STIMER2_COUNT
#define HV_X64_MSR_STIMER2_CONFIG HV_REGISTER_STIMER2_CONFIG
#define HV_X64_MSR_STIMER3_COUNT HV_REGISTER_STIMER3_COUNT
#define HV_X64_MSR_STIMER3_CONFIG HV_REGISTER_STIMER3_CONFIG
#define HV_X64_MSR_SCONTROL HV_REGISTER_SCONTROL
#define HV_X64_MSR_SVERSION HV_REGISTER_SVERSION
#define HV_X64_MSR_SIMP HV_REGISTER_SIMP
#define HV_X64_MSR_SIEFP HV_REGISTER_SIEFP
#define HV_X64_MSR_VP_INDEX HV_REGISTER_VP_INDEX
#define HV_X64_MSR_EOM HV_REGISTER_EOM
#define HV_X64_MSR_SINT0 HV_REGISTER_SINT0
#define HV_X64_MSR_SINT15 HV_REGISTER_SINT15
#define HV_X64_MSR_CRASH_P0 HV_REGISTER_CRASH_P0
#define HV_X64_MSR_CRASH_P1 HV_REGISTER_CRASH_P1
#define HV_X64_MSR_CRASH_P2 HV_REGISTER_CRASH_P2
#define HV_X64_MSR_CRASH_P3 HV_REGISTER_CRASH_P3
#define HV_X64_MSR_CRASH_P4 HV_REGISTER_CRASH_P4
#define HV_X64_MSR_CRASH_CTL HV_REGISTER_CRASH_CTL
#define HV_X64_MSR_TIME_REF_COUNT HV_REGISTER_TIME_REF_COUNT
#define HV_X64_MSR_REFERENCE_TSC HV_REGISTER_REFERENCE_TSC

typedef struct _HV_X64_MSR_STIMER_CONFIG_CONTENTS {
  union {
    UINT64 AsUINT64;

    struct {
      UINT64 Enable : 1;
      UINT64 Periodic : 1;
      UINT64 Lazy : 1;
      UINT64 AutoEnable : 1;
      UINT64 ApicVector : 8;
      UINT64 DirectMode : 1;
      UINT64 ReservedZ1 : 1;
      UINT64 SINTx : 4;
      UINT64 ReservedZ2 : 44;
    };
  };
} HV_X64_MSR_STIMER_CONFIG_CONTENTS, *PHV_X64_MSR_STIMER_CONFIG_CONTENTS;

enum hv_mem_host_visibility {
  VMBUS_PAGE_NOT_VISIBLE = 0,
  VMBUS_PAGE_VISIBLE_READ_ONLY = 1,
  VMBUS_PAGE_VISIBLE_READ_WRITE = 3
};

#define HV_MAX_MODIFY_GPA_REP_COUNT ((PAGE_SIZE / sizeof(UINT64)) - 2)

struct hv_gpa_range_for_visibility {
  UINT64 partition_id;
  UINT32 host_visibility : 2;
  UINT32 reserved0 : 30;
  UINT32 reserved1;
  UINT64 gpa_page_list[HV_MAX_MODIFY_GPA_REP_COUNT];
} __packed;

union hv_x64_msr_hypercall_contents {
  UINT64 as_uint64;

  struct {
    UINT64 enable : 1;
    UINT64 reserved : 11;
    UINT64 guest_physical_address : 52;
  } __packed;
};

union hv_vp_assist_msr_contents {
  UINT64 as_uint64;

  struct {
    UINT64 enable : 1;
    UINT64 reserved : 11;
    UINT64 pfn : 52;
  } __packed;
};

struct hv_reenlightenment_control {
  UINT64 vector : 8;
  UINT64 reserved1 : 8;
  UINT64 enabled : 1;
  UINT64 reserved2 : 15;
  UINT64 target_vp : 32;
};

struct hv_tsc_emulation_control {
  UINT64 enabled : 1;
  UINT64 reserved : 63;
};

struct hv_tsc_emulation_status {
  UINT64 inprogress : 1;
  UINT64 reserved : 63;
};

#define HV_X64_MSR_HYPERCALL_ENABLE 0x00000001
#define HV_X64_MSR_HYPERCALL_PAGE_ADDRESS_SHIFT 12
#define HV_X64_MSR_HYPERCALL_PAGE_ADDRESS_MASK \
  (~((1ull << HV_X64_MSR_HYPERCALL_PAGE_ADDRESS_SHIFT) - 1))
#define HV_X64_MSR_CRASH_PARAMS \
  (1 + (HV_X64_MSR_CRASH_P4 - HV_X64_MSR_CRASH_P0))
#define HV_IPI_LOW_VECTOR 0x10
#define HV_IPI_HIGH_VECTOR 0xff
#define HV_X64_MSR_VP_ASSIST_PAGE_ENABLE 0x00000001
#define HV_X64_MSR_VP_ASSIST_PAGE_ADDRESS_SHIFT 12
#define HV_X64_MSR_VP_ASSIST_PAGE_ADDRESS_MASK \
  (~((1ull << HV_X64_MSR_VP_ASSIST_PAGE_ADDRESS_SHIFT) - 1))

#define HV_X64_ENLIGHTENED_VMCS_VERSION 0xff
#define HV_X64_MSR_TSC_REFERENCE_ENABLE 0x00000001
#define HV_X64_MSR_TSC_REFERENCE_ADDRESS_SHIFT 12

#define HV_HYPERCALL_MAX_XMM_REGISTERS 6

struct hv_nested_enlightenments_control {
  struct {
    UINT32 directhypercall : 1;
    UINT32 reserved : 31;
  } features;

  struct {
    UINT32 reserved;
  } hypercallControls;
};

struct hv_vp_assist_page {
  UINT32 apic_assist;
  UINT32 reserved1;
  UINT64 vtl_control[3];
  struct hv_nested_enlightenments_control nested_control;
  UINT8 enlighten_vmentry;
  UINT8 reserved2[7];
  UINT64 current_nested_vmcs;
};

struct hv_enlightened_vmcs {
  UINT32 revision_id;
  UINT32 abort;
  UINT16 host_es_selector;
  UINT16 host_cs_selector;
  UINT16 host_ss_selector;
  UINT16 host_ds_selector;
  UINT16 host_fs_selector;
  UINT16 host_gs_selector;
  UINT16 host_tr_selector;
  UINT16 padding16_1;
  UINT64 host_ia32_pat;
  UINT64 host_ia32_efer;
  UINT64 host_cr0;
  UINT64 host_cr3;
  UINT64 host_cr4;
  UINT64 host_ia32_sysenter_esp;
  UINT64 host_ia32_sysenter_eip;
  UINT64 host_rip;
  UINT32 host_ia32_sysenter_cs;
  UINT32 pin_based_vm_exec_control;
  UINT32 vm_exit_controls;
  UINT32 secondary_vm_exec_control;
  UINT64 io_bitmap_a;
  UINT64 io_bitmap_b;
  UINT64 msr_bitmap;
  UINT16 guest_es_selector;
  UINT16 guest_cs_selector;
  UINT16 guest_ss_selector;
  UINT16 guest_ds_selector;
  UINT16 guest_fs_selector;
  UINT16 guest_gs_selector;
  UINT16 guest_ldtr_selector;
  UINT16 guest_tr_selector;
  UINT32 guest_es_limit;
  UINT32 guest_cs_limit;
  UINT32 guest_ss_limit;
  UINT32 guest_ds_limit;
  UINT32 guest_fs_limit;
  UINT32 guest_gs_limit;
  UINT32 guest_ldtr_limit;
  UINT32 guest_tr_limit;
  UINT32 guest_gdtr_limit;
  UINT32 guest_idtr_limit;
  UINT32 guest_es_ar_bytes;
  UINT32 guest_cs_ar_bytes;
  UINT32 guest_ss_ar_bytes;
  UINT32 guest_ds_ar_bytes;
  UINT32 guest_fs_ar_bytes;
  UINT32 guest_gs_ar_bytes;
  UINT32 guest_ldtr_ar_bytes;
  UINT32 guest_tr_ar_bytes;
  UINT64 guest_es_base;
  UINT64 guest_cs_base;
  UINT64 guest_ss_base;
  UINT64 guest_ds_base;
  UINT64 guest_fs_base;
  UINT64 guest_gs_base;
  UINT64 guest_ldtr_base;
  UINT64 guest_tr_base;
  UINT64 guest_gdtr_base;
  UINT64 guest_idtr_base;
  UINT64 padding64_1[3];
  UINT64 vm_exit_msr_store_addr;
  UINT64 vm_exit_msr_load_addr;
  UINT64 vm_entry_msr_load_addr;
  UINT64 cr3_target_value0;
  UINT64 cr3_target_value1;
  UINT64 cr3_target_value2;
  UINT64 cr3_target_value3;
  UINT32 page_fault_error_code_mask;
  UINT32 page_fault_error_code_match;
  UINT32 cr3_target_count;
  UINT32 vm_exit_msr_store_count;
  UINT32 vm_exit_msr_load_count;
  UINT32 vm_entry_msr_load_count;
  UINT64 tsc_offset;
  UINT64 virtual_apic_page_addr;
  UINT64 vmcs_link_pointer;
  UINT64 guest_ia32_debugctl;
  UINT64 guest_ia32_pat;
  UINT64 guest_ia32_efer;
  UINT64 guest_pdptr0;
  UINT64 guest_pdptr1;
  UINT64 guest_pdptr2;
  UINT64 guest_pdptr3;
  UINT64 guest_pending_dbg_exceptions;
  UINT64 guest_sysenter_esp;
  UINT64 guest_sysenter_eip;
  UINT32 guest_activity_state;
  UINT32 guest_sysenter_cs;
  UINT64 cr0_guest_host_mask;
  UINT64 cr4_guest_host_mask;
  UINT64 cr0_read_shadow;
  UINT64 cr4_read_shadow;
  UINT64 guest_cr0;
  UINT64 guest_cr3;
  UINT64 guest_cr4;
  UINT64 guest_dr7;
  UINT64 host_fs_base;
  UINT64 host_gs_base;
  UINT64 host_tr_base;
  UINT64 host_gdtr_base;
  UINT64 host_idtr_base;
  UINT64 host_rsp;
  UINT64 ept_pointer;
  UINT16 virtual_processor_id;
  UINT16 padding16_2[3];
  UINT64 padding64_2[5];
  UINT64 guest_physical_address;
  UINT32 vm_instruction_error;
  UINT32 vm_exit_reason;
  UINT32 vm_exit_intr_info;
  UINT32 vm_exit_intr_error_code;
  UINT32 idt_vectoring_info_field;
  UINT32 idt_vectoring_error_code;
  UINT32 vm_exit_instruction_len;
  UINT32 vmx_instruction_info;
  UINT64 exit_qualification;
  UINT64 exit_io_instruction_ecx;
  UINT64 exit_io_instruction_esi;
  UINT64 exit_io_instruction_edi;
  UINT64 exit_io_instruction_eip;
  UINT64 guest_linear_address;
  UINT64 guest_rsp;
  UINT64 guest_rflags;
  UINT32 guest_interruptibility_info;
  UINT32 cpu_based_vm_exec_control;
  UINT32 exception_bitmap;
  UINT32 vm_entry_controls;
  UINT32 vm_entry_intr_info_field;
  UINT32 vm_entry_exception_error_code;
  UINT32 vm_entry_instruction_len;
  UINT32 tpr_threshold;
  UINT64 guest_rip;
  UINT32 hv_clean_fields;
  UINT32 hv_padding_32;
  UINT32 hv_synthetic_controls;

  struct {
    UINT32 nested_flush_hypercall : 1;
    UINT32 msr_bitmap : 1;
    UINT32 reserved : 30;
  } hv_enlightenments_control;

  UINT32 hv_vp_id;
  UINT64 hv_vm_id;
  UINT64 partition_assist_page;
  UINT64 padding64_4[4];
  UINT64 guest_bndcfgs;
  UINT64 padding64_5[7];
  UINT64 xss_exit_bitmap;
  UINT64 padding64_6[7];
};

#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_NONE 0
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_IO_BITMAP BIT(0)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_MSR_BITMAP BIT(1)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_CONTROL_GRP2 BIT(2)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_CONTROL_GRP1 BIT(3)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_CONTROL_PROC BIT(4)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_CONTROL_EVENT BIT(5)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_CONTROL_ENTRY BIT(6)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_CONTROL_EXCPN BIT(7)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_CRDR BIT(8)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_CONTROL_XLAT BIT(9)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_GUEST_BASIC BIT(10)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_GUEST_GRP1 BIT(11)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_GUEST_GRP2 BIT(12)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_HOST_POINTER BIT(13)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_HOST_GRP1 BIT(14)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_ENLIGHTENMENTSCONTROL BIT(15)
#define HV_VMX_ENLIGHTENED_CLEAN_FIELD_ALL 0xFFFF

struct hv_partition_assist_pg {
  UINT32 tlb_lock_count;
};

enum hv_interrupt_type {
  HV_X64_INTERRUPT_TYPE_FIXED = 0x0000,
  HV_X64_INTERRUPT_TYPE_LOWESTPRIORITY = 0x0001,
  HV_X64_INTERRUPT_TYPE_SMI = 0x0002,
  HV_X64_INTERRUPT_TYPE_REMOTEREAD = 0x0003,
  HV_X64_INTERRUPT_TYPE_NMI = 0x0004,
  HV_X64_INTERRUPT_TYPE_INIT = 0x0005,
  HV_X64_INTERRUPT_TYPE_SIPI = 0x0006,
  HV_X64_INTERRUPT_TYPE_EXTINT = 0x0007,
  HV_X64_INTERRUPT_TYPE_LOCALINT0 = 0x0008,
  HV_X64_INTERRUPT_TYPE_LOCALINT1 = 0x0009,
  HV_X64_INTERRUPT_TYPE_MAXIMUM = 0x000A,
};

union hv_msi_address_register {
  UINT32 as_uint32;

  struct {
    UINT32 reserved1 : 2;
    UINT32 destination_mode : 1;
    UINT32 redirection_hint : 1;
    UINT32 reserved2 : 8;
    UINT32 destination_id : 8;
    UINT32 msi_base : 12;
  };
};

union hv_msi_data_register {
  UINT32 as_uint32;

  struct {
    UINT32 vector : 8;
    UINT32 delivery_mode : 3;
    UINT32 reserved1 : 3;
    UINT32 level_assert : 1;
    UINT32 trigger_mode : 1;
    UINT32 reserved2 : 16;
  };
};

union hv_msi_entry {
  UINT64 as_uint64;

  struct {
    union hv_msi_address_register address;
    union hv_msi_data_register data;
  } __packed;
};
