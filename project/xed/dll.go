package xed

import (
	_ "embed"
	"os"
	"path/filepath"
	"unsafe"

	"golang.org/x/sys/windows"
)

type Xed struct{}

//go:embed xed.dll
var dllBytes []byte

var (
	dll                                                             *windows.LazyDLL
	proc_xed_strlen                                                 *windows.LazyProc
	proc_xed_strcat                                                 *windows.LazyProc
	proc_xed_strcpy                                                 *windows.LazyProc
	proc_xed_strncpy                                                *windows.LazyProc
	proc_xed_strncat                                                *windows.LazyProc
	proc_xed_operand_enum_t2str                                     *windows.LazyProc
	proc_xed_operand_enum_t_last                                    *windows.LazyProc
	proc_xed_tables_init                                            *windows.LazyProc
	proc_xed_internal_assert                                        *windows.LazyProc
	proc_xed_register_abort_function                                *windows.LazyProc
	proc_xed_itoa                                                   *windows.LazyProc
	proc_xed_itoa_bin                                               *windows.LazyProc
	proc_xed_itoa_hex_zeros                                         *windows.LazyProc
	proc_xed_itoa_hex                                               *windows.LazyProc
	proc_xed_itoa_hex_ul                                            *windows.LazyProc
	proc_xed_set_log_file                                           *windows.LazyProc
	proc_xed_set_verbosity                                          *windows.LazyProc
	proc_xed_sign_extend32_64                                       *windows.LazyProc
	proc_xed_sign_extend16_64                                       *windows.LazyProc
	proc_xed_sign_extend8_64                                        *windows.LazyProc
	proc_xed_sign_extend16_32                                       *windows.LazyProc
	proc_xed_sign_extend8_32                                        *windows.LazyProc
	proc_xed_sign_extend8_16                                        *windows.LazyProc
	proc_xed_sign_extend_arbitrary_to_32                            *windows.LazyProc
	proc_xed_sign_extend_arbitrary_to_64                            *windows.LazyProc
	proc_xed_zero_extend32_64                                       *windows.LazyProc
	proc_xed_zero_extend16_64                                       *windows.LazyProc
	proc_xed_zero_extend8_64                                        *windows.LazyProc
	proc_xed_zero_extend16_32                                       *windows.LazyProc
	proc_xed_zero_extend8_32                                        *windows.LazyProc
	proc_xed_zero_extend8_16                                        *windows.LazyProc
	proc_xed_get_byte                                               *windows.LazyProc
	proc_xed_shortest_width_unsigned                                *windows.LazyProc
	proc_xed_shortest_width_signed                                  *windows.LazyProc
	proc_xed_category_enum_t2str                                    *windows.LazyProc
	proc_xed_category_enum_t_last                                   *windows.LazyProc
	proc_xed_extension_enum_t2str                                   *windows.LazyProc
	proc_xed_extension_enum_t_last                                  *windows.LazyProc
	proc_xed_iclass_enum_t2str                                      *windows.LazyProc
	proc_xed_iclass_enum_t_last                                     *windows.LazyProc
	proc_xed_operand_visibility_enum_t2str                          *windows.LazyProc
	proc_xed_operand_visibility_enum_t_last                         *windows.LazyProc
	proc_xed_operand_action_enum_t2str                              *windows.LazyProc
	proc_xed_operand_action_enum_t_last                             *windows.LazyProc
	proc_xed_operand_convert_enum_t2str                             *windows.LazyProc
	proc_xed_operand_convert_enum_t_last                            *windows.LazyProc
	proc_xed_operand_type_enum_t2str                                *windows.LazyProc
	proc_xed_operand_type_enum_t_last                               *windows.LazyProc
	proc_xed_nonterminal_enum_t2str                                 *windows.LazyProc
	proc_xed_nonterminal_enum_t_last                                *windows.LazyProc
	proc_xed_operand_width_enum_t2str                               *windows.LazyProc
	proc_xed_operand_width_enum_t_last                              *windows.LazyProc
	proc_xed_operand_element_xtype_enum_t2str                       *windows.LazyProc
	proc_xed_operand_element_xtype_enum_t_last                      *windows.LazyProc
	proc_xed_reg_enum_t2str                                         *windows.LazyProc
	proc_xed_reg_enum_t_last                                        *windows.LazyProc
	proc_xed_attribute_enum_t2str                                   *windows.LazyProc
	proc_xed_attribute_enum_t_last                                  *windows.LazyProc
	proc_xed_exception_enum_t2str                                   *windows.LazyProc
	proc_xed_exception_enum_t_last                                  *windows.LazyProc
	proc_xed_iform_enum_t2str                                       *windows.LazyProc
	proc_xed_iform_enum_t_last                                      *windows.LazyProc
	proc_xed_isa_set_enum_t2str                                     *windows.LazyProc
	proc_xed_isa_set_enum_t_last                                    *windows.LazyProc
	proc_xed_iform_map                                              *windows.LazyProc
	proc_xed_iform_max_per_iclass                                   *windows.LazyProc
	proc_xed_iform_first_per_iclass                                 *windows.LazyProc
	proc_xed_iform_to_category                                      *windows.LazyProc
	proc_xed_iform_to_extension                                     *windows.LazyProc
	proc_xed_iform_to_isa_set                                       *windows.LazyProc
	proc_xed_iform_to_iclass_string_att                             *windows.LazyProc
	proc_xed_iform_to_iclass_string_intel                           *windows.LazyProc
	proc_xed_operand_width_bits                                     *windows.LazyProc
	proc_xed_operand_print                                          *windows.LazyProc
	proc_xed_operand_read                                           *windows.LazyProc
	proc_xed_operand_read_only                                      *windows.LazyProc
	proc_xed_operand_written                                        *windows.LazyProc
	proc_xed_operand_written_only                                   *windows.LazyProc
	proc_xed_operand_read_and_written                               *windows.LazyProc
	proc_xed_operand_conditional_read                               *windows.LazyProc
	proc_xed_operand_conditional_write                              *windows.LazyProc
	proc_xed_inst_cpl                                               *windows.LazyProc
	proc_xed_inst_operand                                           *windows.LazyProc
	proc_xed_inst_flag_info_index                                   *windows.LazyProc
	proc_xed_inst_get_attribute                                     *windows.LazyProc
	proc_xed_inst_get_attributes                                    *windows.LazyProc
	proc_xed_attribute_max                                          *windows.LazyProc
	proc_xed_attribute                                              *windows.LazyProc
	proc_xed_inst_table_base                                        *windows.LazyProc
	proc_xed_flag_enum_t2str                                        *windows.LazyProc
	proc_xed_flag_enum_t_last                                       *windows.LazyProc
	proc_xed_flag_action_enum_t2str                                 *windows.LazyProc
	proc_xed_flag_action_enum_t_last                                *windows.LazyProc
	proc_xed_flag_set_print                                         *windows.LazyProc
	proc_xed_flag_set_is_subset_of                                  *windows.LazyProc
	proc_xed_flag_action_get_flag_name                              *windows.LazyProc
	proc_xed_flag_action_get_action                                 *windows.LazyProc
	proc_xed_flag_action_action_invalid                             *windows.LazyProc
	proc_xed_flag_action_print                                      *windows.LazyProc
	proc_xed_flag_action_read_flag                                  *windows.LazyProc
	proc_xed_flag_action_writes_flag                                *windows.LazyProc
	proc_xed_flag_action_read_action                                *windows.LazyProc
	proc_xed_flag_action_write_action                               *windows.LazyProc
	proc_xed_simple_flag_get_nflags                                 *windows.LazyProc
	proc_xed_simple_flag_get_read_flag_set                          *windows.LazyProc
	proc_xed_simple_flag_get_written_flag_set                       *windows.LazyProc
	proc_xed_simple_flag_get_undefined_flag_set                     *windows.LazyProc
	proc_xed_simple_flag_get_may_write                              *windows.LazyProc
	proc_xed_simple_flag_get_must_write                             *windows.LazyProc
	proc_xed_simple_flag_get_flag_action                            *windows.LazyProc
	proc_xed_simple_flag_reads_flags                                *windows.LazyProc
	proc_xed_simple_flag_writes_flags                               *windows.LazyProc
	proc_xed_simple_flag_print                                      *windows.LazyProc
	proc_xed_chip_enum_t2str                                        *windows.LazyProc
	proc_xed_chip_enum_t_last                                       *windows.LazyProc
	proc_xed_operand_element_type_enum_t2str                        *windows.LazyProc
	proc_xed_operand_element_type_enum_t_last                       *windows.LazyProc
	proc_xed_error_enum_t2str                                       *windows.LazyProc
	proc_xed_error_enum_t_last                                      *windows.LazyProc
	proc_xed_get_chip_features                                      *windows.LazyProc
	proc_xed_modify_chip_features                                   *windows.LazyProc
	proc_xed_set_decoder_modes                                      *windows.LazyProc
	proc_xed_decode                                                 *windows.LazyProc
	proc_xed_decode_with_features                                   *windows.LazyProc
	proc_xed_ild_decode                                             *windows.LazyProc
	proc_xed_address_width_enum_t2str                               *windows.LazyProc
	proc_xed_address_width_enum_t_last                              *windows.LazyProc
	proc_xed_machine_mode_enum_t2str                                *windows.LazyProc
	proc_xed_machine_mode_enum_t_last                               *windows.LazyProc
	proc_xed_state_print                                            *windows.LazyProc
	proc_xed_syntax_enum_t2str                                      *windows.LazyProc
	proc_xed_syntax_enum_t_last                                     *windows.LazyProc
	proc_xed_reg_class_enum_t2str                                   *windows.LazyProc
	proc_xed_reg_class_enum_t_last                                  *windows.LazyProc
	proc_xed_reg_class                                              *windows.LazyProc
	proc_xed_gpr_reg_class                                          *windows.LazyProc
	proc_xed_get_largest_enclosing_register                         *windows.LazyProc
	proc_xed_get_largest_enclosing_register32                       *windows.LazyProc
	proc_xed_get_register_width_bits                                *windows.LazyProc
	proc_xed_get_register_width_bits64                              *windows.LazyProc
	proc_xed_operand_values_init                                    *windows.LazyProc
	proc_xed_operand_values_init_set_mode                           *windows.LazyProc
	proc_xed_operand_values_set_mode                                *windows.LazyProc
	proc_xed_operand_values_init_keep_mode                          *windows.LazyProc
	proc_xed_operand_values_dump                                    *windows.LazyProc
	proc_xed_operand_values_print_short                             *windows.LazyProc
	proc_xed_operand_values_has_real_rep                            *windows.LazyProc
	proc_xed_operand_values_has_rep_prefix                          *windows.LazyProc
	proc_xed_operand_values_has_repne_prefix                        *windows.LazyProc
	proc_xed_operand_values_clear_rep                               *windows.LazyProc
	proc_xed_operand_values_get_atomic                              *windows.LazyProc
	proc_xed_operand_values_has_lock_prefix                         *windows.LazyProc
	proc_xed_operand_values_lockable                                *windows.LazyProc
	proc_xed_operand_values_using_default_segment                   *windows.LazyProc
	proc_xed_operand_values_get_effective_operand_width             *windows.LazyProc
	proc_xed_operand_values_get_effective_address_width             *windows.LazyProc
	proc_xed_operand_values_get_stack_address_width                 *windows.LazyProc
	proc_xed_operand_values_has_memory_displacement                 *windows.LazyProc
	proc_xed_operand_values_has_branch_displacement                 *windows.LazyProc
	proc_xed_operand_values_has_displacement                        *windows.LazyProc
	proc_xed_operand_values_get_displacement_for_memop              *windows.LazyProc
	proc_xed_operand_values_has_immediate                           *windows.LazyProc
	proc_xed_operand_values_has_address_size_prefix                 *windows.LazyProc
	proc_xed_operand_values_has_operand_size_prefix                 *windows.LazyProc
	proc_xed_operand_values_has_66_prefix                           *windows.LazyProc
	proc_xed_operand_values_mandatory_66_prefix                     *windows.LazyProc
	proc_xed_operand_values_has_rexw_prefix                         *windows.LazyProc
	proc_xed_operand_values_has_segment_prefix                      *windows.LazyProc
	proc_xed_operand_values_segment_prefix                          *windows.LazyProc
	proc_xed_operand_values_get_long_mode                           *windows.LazyProc
	proc_xed_operand_values_get_real_mode                           *windows.LazyProc
	proc_xed_operand_values_get_pp_vex_prefix                       *windows.LazyProc
	proc_xed_operand_values_accesses_memory                         *windows.LazyProc
	proc_xed_operand_values_number_of_memory_operands               *windows.LazyProc
	proc_xed_operand_values_get_memory_operand_length               *windows.LazyProc
	proc_xed_operand_values_get_base_reg                            *windows.LazyProc
	proc_xed_operand_values_get_index_reg                           *windows.LazyProc
	proc_xed_operand_values_get_seg_reg                             *windows.LazyProc
	proc_xed_operand_values_get_scale                               *windows.LazyProc
	proc_xed_operand_values_memop_without_modrm                     *windows.LazyProc
	proc_xed_operand_values_has_modrm_byte                          *windows.LazyProc
	proc_xed_operand_values_has_sib_byte                            *windows.LazyProc
	proc_xed_operand_values_branch_not_taken_hint                   *windows.LazyProc
	proc_xed_operand_values_branch_taken_hint                       *windows.LazyProc
	proc_xed_operand_values_ignored_branch_not_taken_hint           *windows.LazyProc
	proc_xed_operand_values_ignored_branch_taken_hint               *windows.LazyProc
	proc_xed_operand_values_cet_no_track                            *windows.LazyProc
	proc_xed_operand_values_is_nop                                  *windows.LazyProc
	proc_xed_operand_values_get_immediate_int64                     *windows.LazyProc
	proc_xed_operand_values_get_immediate_uint64                    *windows.LazyProc
	proc_xed_operand_values_get_immediate_is_signed                 *windows.LazyProc
	proc_xed_operand_values_get_immediate_byte                      *windows.LazyProc
	proc_xed_operand_values_get_second_immediate                    *windows.LazyProc
	proc_xed_operand_values_get_memory_displacement_length          *windows.LazyProc
	proc_xed_operand_values_get_memory_displacement_length_bits     *windows.LazyProc
	proc_xed_operand_values_get_memory_displacement_length_bits_raw *windows.LazyProc
	proc_xed_operand_values_get_memory_displacement_int64           *windows.LazyProc
	proc_xed_operand_values_get_memory_displacement_int64_raw       *windows.LazyProc
	proc_xed_operand_values_get_memory_displacement_byte            *windows.LazyProc
	proc_xed_operand_values_get_branch_displacement_length          *windows.LazyProc
	proc_xed_operand_values_get_branch_displacement_length_bits     *windows.LazyProc
	proc_xed_operand_values_get_branch_displacement_int64           *windows.LazyProc
	proc_xed_operand_values_get_branch_displacement_byte            *windows.LazyProc
	proc_xed_operand_values_get_iclass                              *windows.LazyProc
	proc_xed_operand_values_zero_immediate                          *windows.LazyProc
	proc_xed_operand_values_zero_branch_displacement                *windows.LazyProc
	proc_xed_operand_values_zero_memory_displacement                *windows.LazyProc
	proc_xed_operand_values_set_lock                                *windows.LazyProc
	proc_xed_operand_values_zero_segment_override                   *windows.LazyProc
	proc_xed_operand_values_set_iclass                              *windows.LazyProc
	proc_xed_operand_values_set_effective_operand_width             *windows.LazyProc
	proc_xed_operand_values_set_effective_address_width             *windows.LazyProc
	proc_xed_operand_values_set_memory_operand_length               *windows.LazyProc
	proc_xed_operand_values_set_memory_displacement                 *windows.LazyProc
	proc_xed_operand_values_set_memory_displacement_bits            *windows.LazyProc
	proc_xed_operand_values_set_relbr                               *windows.LazyProc
	proc_xed_operand_values_set_absbr                               *windows.LazyProc
	proc_xed_operand_values_set_branch_displacement                 *windows.LazyProc
	proc_xed_operand_values_set_branch_displacement_bits            *windows.LazyProc
	proc_xed_operand_values_set_immediate_signed                    *windows.LazyProc
	proc_xed_operand_values_set_immediate_signed_bits               *windows.LazyProc
	proc_xed_operand_values_set_immediate_unsigned                  *windows.LazyProc
	proc_xed_operand_values_set_immediate_unsigned_bits             *windows.LazyProc
	proc_xed_operand_values_set_base_reg                            *windows.LazyProc
	proc_xed_operand_values_set_seg_reg                             *windows.LazyProc
	proc_xed_operand_values_set_index_reg                           *windows.LazyProc
	proc_xed_operand_values_set_scale                               *windows.LazyProc
	proc_xed_operand_values_set_operand_reg                         *windows.LazyProc
	proc_xed_encoder_request_get_iclass                             *windows.LazyProc
	proc_xed_encoder_request_set_iclass                             *windows.LazyProc
	proc_xed_encoder_request_set_effective_operand_width            *windows.LazyProc
	proc_xed_encoder_request_set_effective_address_size             *windows.LazyProc
	proc_xed_encoder_request_set_reg                                *windows.LazyProc
	proc_xed_encoder_request_set_operand_order                      *windows.LazyProc
	proc_xed_encoder_request_get_operand_order                      *windows.LazyProc
	proc_xed_encoder_request_set_relbr                              *windows.LazyProc
	proc_xed_encoder_request_set_absbr                              *windows.LazyProc
	proc_xed_encoder_request_set_branch_displacement                *windows.LazyProc
	proc_xed_encoder_request_set_ptr                                *windows.LazyProc
	proc_xed_encoder_request_set_uimm0                              *windows.LazyProc
	proc_xed_encoder_request_set_uimm0_bits                         *windows.LazyProc
	proc_xed_encoder_request_set_uimm1                              *windows.LazyProc
	proc_xed_encoder_request_set_simm                               *windows.LazyProc
	proc_xed_encoder_request_set_memory_displacement                *windows.LazyProc
	proc_xed_encoder_request_set_agen                               *windows.LazyProc
	proc_xed_encoder_request_set_mem0                               *windows.LazyProc
	proc_xed_encoder_request_set_mem1                               *windows.LazyProc
	proc_xed_encoder_request_set_memory_operand_length              *windows.LazyProc
	proc_xed_encoder_request_set_seg0                               *windows.LazyProc
	proc_xed_encoder_request_set_seg1                               *windows.LazyProc
	proc_xed_encoder_request_set_base0                              *windows.LazyProc
	proc_xed_encoder_request_set_base1                              *windows.LazyProc
	proc_xed_encoder_request_set_index                              *windows.LazyProc
	proc_xed_encoder_request_set_scale                              *windows.LazyProc
	proc_xed_encoder_request_zero_operand_order                     *windows.LazyProc
	proc_xed_encoder_request_zero_set_mode                          *windows.LazyProc
	proc_xed_encoder_request_zero                                   *windows.LazyProc
	proc_xed_encoder_request_init_from_decode                       *windows.LazyProc
	proc_xed_encode_request_print                                   *windows.LazyProc
	proc_xed_encode                                                 *windows.LazyProc
	proc_xed_encode_nop                                             *windows.LazyProc
	proc_xed_convert_to_encoder_request                             *windows.LazyProc
	proc_xed_patch_disp                                             *windows.LazyProc
	proc_xed_patch_brdisp                                           *windows.LazyProc
	proc_xed_patch_imm0                                             *windows.LazyProc
	proc_xed_operand_action_read                                    *windows.LazyProc
	proc_xed_operand_action_read_only                               *windows.LazyProc
	proc_xed_operand_action_written                                 *windows.LazyProc
	proc_xed_operand_action_written_only                            *windows.LazyProc
	proc_xed_operand_action_read_and_written                        *windows.LazyProc
	proc_xed_operand_action_conditional_read                        *windows.LazyProc
	proc_xed_operand_action_conditional_write                       *windows.LazyProc
	proc_xed_get_version                                            *windows.LazyProc
	proc_xed_get_copyright                                          *windows.LazyProc
	proc_xed_format_set_options                                     *windows.LazyProc
	proc_xed_init_print_info                                        *windows.LazyProc
	proc_xed_decoded_inst_get_attribute                             *windows.LazyProc
	proc_xed_decoded_inst_get_attributes                            *windows.LazyProc
	proc_xed_decoded_inst_is_xacquire                               *windows.LazyProc
	proc_xed_decoded_inst_is_xrelease                               *windows.LazyProc
	proc_xed_decoded_inst_has_mpx_prefix                            *windows.LazyProc
	proc_xed_decoded_inst_is_apx_zu                                 *windows.LazyProc
	proc_xed_decoded_inst_get_modrm                                 *windows.LazyProc
	proc_xed_decoded_inst_masked_vector_operation                   *windows.LazyProc
	proc_xed_decoded_inst_vector_length_bits                        *windows.LazyProc
	proc_xed_decoded_inst_get_nprefixes                             *windows.LazyProc
	proc_xed_decoded_inst_operand_length_bits                       *windows.LazyProc
	proc_xed_decoded_inst_operand_length                            *windows.LazyProc
	proc_xed_decoded_inst_operand_elements                          *windows.LazyProc
	proc_xed_decoded_inst_operand_element_size_bits                 *windows.LazyProc
	proc_xed_decoded_inst_operand_element_type                      *windows.LazyProc
	proc_xed_decoded_inst_operand_action                            *windows.LazyProc
	proc_xed_decoded_inst_masking                                   *windows.LazyProc
	proc_xed_decoded_inst_merging                                   *windows.LazyProc
	proc_xed_decoded_inst_zeroing                                   *windows.LazyProc
	proc_xed_decoded_inst_avx512_dest_elements                      *windows.LazyProc
	proc_xed_decoded_inst_zero_set_mode                             *windows.LazyProc
	proc_xed_decoded_inst_zero_keep_mode                            *windows.LazyProc
	proc_xed_decoded_inst_zero                                      *windows.LazyProc
	proc_xed_decoded_inst_zero_keep_mode_from_operands              *windows.LazyProc
	proc_xed_decoded_inst_get_operand_width                         *windows.LazyProc
	proc_xed_decoded_inst_valid_for_chip                            *windows.LazyProc
	proc_xed_decoded_inst_valid_for_features                        *windows.LazyProc
	proc_xed_decoded_inst_dump                                      *windows.LazyProc
	proc_xed_decoded_inst_dump_xed_format                           *windows.LazyProc
	proc_xed_format_context                                         *windows.LazyProc
	proc_xed_format_generic                                         *windows.LazyProc
	proc_xed_decoded_inst_get_seg_reg                               *windows.LazyProc
	proc_xed_decoded_inst_get_base_reg                              *windows.LazyProc
	proc_xed_decoded_inst_get_index_reg                             *windows.LazyProc
	proc_xed_decoded_inst_get_scale                                 *windows.LazyProc
	proc_xed_decoded_inst_get_memory_displacement                   *windows.LazyProc
	proc_xed_decoded_inst_get_memory_displacement_width             *windows.LazyProc
	proc_xed_decoded_inst_get_memory_displacement_width_bits        *windows.LazyProc
	proc_xed_decoded_inst_get_branch_displacement                   *windows.LazyProc
	proc_xed_decoded_inst_get_branch_displacement_width             *windows.LazyProc
	proc_xed_decoded_inst_get_branch_displacement_width_bits        *windows.LazyProc
	proc_xed_decoded_inst_get_unsigned_immediate                    *windows.LazyProc
	proc_xed_decoded_inst_get_immediate_is_signed                   *windows.LazyProc
	proc_xed_decoded_inst_get_immediate_width                       *windows.LazyProc
	proc_xed_decoded_inst_get_immediate_width_bits                  *windows.LazyProc
	proc_xed_decoded_inst_get_signed_immediate                      *windows.LazyProc
	proc_xed_decoded_inst_get_reg                                   *windows.LazyProc
	proc_xed_decoded_inst_has_default_flags_values                  *windows.LazyProc
	proc_xed_decoded_inst_get_default_flags_values                  *windows.LazyProc
	proc_xed_decoded_inst_get_rflags_info                           *windows.LazyProc
	proc_xed_decoded_inst_uses_rflags                               *windows.LazyProc
	proc_xed_decoded_inst_number_of_memory_operands                 *windows.LazyProc
	proc_xed_decoded_inst_mem_read                                  *windows.LazyProc
	proc_xed_decoded_inst_mem_written                               *windows.LazyProc
	proc_xed_decoded_inst_mem_written_only                          *windows.LazyProc
	proc_xed_decoded_inst_conditionally_writes_registers            *windows.LazyProc
	proc_xed_decoded_inst_get_memory_operand_length                 *windows.LazyProc
	proc_xed_decoded_inst_get_memop_address_width                   *windows.LazyProc
	proc_xed_decoded_inst_is_prefetch                               *windows.LazyProc
	proc_xed_decoded_inst_is_broadcast                              *windows.LazyProc
	proc_xed_decoded_inst_is_broadcast_instruction                  *windows.LazyProc
	proc_xed_decoded_inst_uses_embedded_broadcast                   *windows.LazyProc
	proc_xed_decoded_inst_set_scale                                 *windows.LazyProc
	proc_xed_decoded_inst_set_memory_displacement                   *windows.LazyProc
	proc_xed_decoded_inst_set_branch_displacement                   *windows.LazyProc
	proc_xed_decoded_inst_set_immediate_signed                      *windows.LazyProc
	proc_xed_decoded_inst_set_immediate_unsigned                    *windows.LazyProc
	proc_xed_decoded_inst_set_memory_displacement_bits              *windows.LazyProc
	proc_xed_decoded_inst_set_branch_displacement_bits              *windows.LazyProc
	proc_xed_decoded_inst_set_immediate_signed_bits                 *windows.LazyProc
	proc_xed_decoded_inst_set_immediate_unsigned_bits               *windows.LazyProc
	proc_xed_classify_amx                                           *windows.LazyProc
	proc_xed_classify_avx512                                        *windows.LazyProc
	proc_xed_classify_avx512_maskop                                 *windows.LazyProc
	proc_xed_classify_avx                                           *windows.LazyProc
	proc_xed_classify_sse                                           *windows.LazyProc
	proc_xed_classify_apx_foundation                                *windows.LazyProc
	proc_xed_classify_apx                                           *windows.LazyProc
	proc_xed_iformfl_enum_t_last                                    *windows.LazyProc
	proc_xed_rep_remove                                             *windows.LazyProc
	proc_xed_repe_map                                               *windows.LazyProc
	proc_xed_repne_map                                              *windows.LazyProc
	proc_xed_rep_map                                                *windows.LazyProc
	proc_xed_norep_map                                              *windows.LazyProc
	proc_xed_agen_register_callback                                 *windows.LazyProc
	proc_xed_agen                                                   *windows.LazyProc
	proc_xed_cpuid_rec_enum_t2str                                   *windows.LazyProc
	proc_xed_cpuid_rec_enum_t_last                                  *windows.LazyProc
	proc_xed_cpuid_group_enum_t2str                                 *windows.LazyProc
	proc_xed_cpuid_group_enum_t_last                                *windows.LazyProc
	proc_xed_get_cpuid_group_enum_for_isa_set                       *windows.LazyProc
	proc_xed_get_cpuid_rec_enum_for_group                           *windows.LazyProc
	proc_xed_get_cpuid_rec                                          *windows.LazyProc
	proc_xed_isa_set_is_valid_for_chip                              *windows.LazyProc
)

func init() {
	dll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, "xed.dll"))
	proc_xed_strlen = dll.NewProc("xed_strlen")
	proc_xed_strcat = dll.NewProc("xed_strcat")
	proc_xed_strcpy = dll.NewProc("xed_strcpy")
	proc_xed_strncpy = dll.NewProc("xed_strncpy")
	proc_xed_strncat = dll.NewProc("xed_strncat")
	proc_xed_operand_enum_t2str = dll.NewProc("xed_operand_enum_t2str")
	proc_xed_operand_enum_t_last = dll.NewProc("xed_operand_enum_t_last")
	proc_xed_tables_init = dll.NewProc("xed_tables_init")
	proc_xed_internal_assert = dll.NewProc("xed_internal_assert")
	proc_xed_register_abort_function = dll.NewProc("xed_register_abort_function")
	proc_xed_itoa = dll.NewProc("xed_itoa")
	proc_xed_itoa_bin = dll.NewProc("xed_itoa_bin")
	proc_xed_itoa_hex_zeros = dll.NewProc("xed_itoa_hex_zeros")
	proc_xed_itoa_hex = dll.NewProc("xed_itoa_hex")
	proc_xed_itoa_hex_ul = dll.NewProc("xed_itoa_hex_ul")
	proc_xed_set_log_file = dll.NewProc("xed_set_log_file")
	proc_xed_set_verbosity = dll.NewProc("xed_set_verbosity")
	proc_xed_sign_extend32_64 = dll.NewProc("xed_sign_extend32_64")
	proc_xed_sign_extend16_64 = dll.NewProc("xed_sign_extend16_64")
	proc_xed_sign_extend8_64 = dll.NewProc("xed_sign_extend8_64")
	proc_xed_sign_extend16_32 = dll.NewProc("xed_sign_extend16_32")
	proc_xed_sign_extend8_32 = dll.NewProc("xed_sign_extend8_32")
	proc_xed_sign_extend8_16 = dll.NewProc("xed_sign_extend8_16")
	proc_xed_sign_extend_arbitrary_to_32 = dll.NewProc("xed_sign_extend_arbitrary_to_32")
	proc_xed_sign_extend_arbitrary_to_64 = dll.NewProc("xed_sign_extend_arbitrary_to_64")
	proc_xed_zero_extend32_64 = dll.NewProc("xed_zero_extend32_64")
	proc_xed_zero_extend16_64 = dll.NewProc("xed_zero_extend16_64")
	proc_xed_zero_extend8_64 = dll.NewProc("xed_zero_extend8_64")
	proc_xed_zero_extend16_32 = dll.NewProc("xed_zero_extend16_32")
	proc_xed_zero_extend8_32 = dll.NewProc("xed_zero_extend8_32")
	proc_xed_zero_extend8_16 = dll.NewProc("xed_zero_extend8_16")
	proc_xed_get_byte = dll.NewProc("xed_get_byte")
	proc_xed_shortest_width_unsigned = dll.NewProc("xed_shortest_width_unsigned")
	proc_xed_shortest_width_signed = dll.NewProc("xed_shortest_width_signed")
	proc_xed_category_enum_t2str = dll.NewProc("xed_category_enum_t2str")
	proc_xed_category_enum_t_last = dll.NewProc("xed_category_enum_t_last")
	proc_xed_extension_enum_t2str = dll.NewProc("xed_extension_enum_t2str")
	proc_xed_extension_enum_t_last = dll.NewProc("xed_extension_enum_t_last")
	proc_xed_iclass_enum_t2str = dll.NewProc("xed_iclass_enum_t2str")
	proc_xed_iclass_enum_t_last = dll.NewProc("xed_iclass_enum_t_last")
	proc_xed_operand_visibility_enum_t2str = dll.NewProc("xed_operand_visibility_enum_t2str")
	proc_xed_operand_visibility_enum_t_last = dll.NewProc("xed_operand_visibility_enum_t_last")
	proc_xed_operand_action_enum_t2str = dll.NewProc("xed_operand_action_enum_t2str")
	proc_xed_operand_action_enum_t_last = dll.NewProc("xed_operand_action_enum_t_last")
	proc_xed_operand_convert_enum_t2str = dll.NewProc("xed_operand_convert_enum_t2str")
	proc_xed_operand_convert_enum_t_last = dll.NewProc("xed_operand_convert_enum_t_last")
	proc_xed_operand_type_enum_t2str = dll.NewProc("xed_operand_type_enum_t2str")
	proc_xed_operand_type_enum_t_last = dll.NewProc("xed_operand_type_enum_t_last")
	proc_xed_nonterminal_enum_t2str = dll.NewProc("xed_nonterminal_enum_t2str")
	proc_xed_nonterminal_enum_t_last = dll.NewProc("xed_nonterminal_enum_t_last")
	proc_xed_operand_width_enum_t2str = dll.NewProc("xed_operand_width_enum_t2str")
	proc_xed_operand_width_enum_t_last = dll.NewProc("xed_operand_width_enum_t_last")
	proc_xed_operand_element_xtype_enum_t2str = dll.NewProc("xed_operand_element_xtype_enum_t2str")
	proc_xed_operand_element_xtype_enum_t_last = dll.NewProc("xed_operand_element_xtype_enum_t_last")
	proc_xed_reg_enum_t2str = dll.NewProc("xed_reg_enum_t2str")
	proc_xed_reg_enum_t_last = dll.NewProc("xed_reg_enum_t_last")
	proc_xed_attribute_enum_t2str = dll.NewProc("xed_attribute_enum_t2str")
	proc_xed_attribute_enum_t_last = dll.NewProc("xed_attribute_enum_t_last")
	proc_xed_exception_enum_t2str = dll.NewProc("xed_exception_enum_t2str")
	proc_xed_exception_enum_t_last = dll.NewProc("xed_exception_enum_t_last")
	proc_xed_iform_enum_t2str = dll.NewProc("xed_iform_enum_t2str")
	proc_xed_iform_enum_t_last = dll.NewProc("xed_iform_enum_t_last")
	proc_xed_isa_set_enum_t2str = dll.NewProc("xed_isa_set_enum_t2str")
	proc_xed_isa_set_enum_t_last = dll.NewProc("xed_isa_set_enum_t_last")
	proc_xed_iform_map = dll.NewProc("xed_iform_map")
	proc_xed_iform_max_per_iclass = dll.NewProc("xed_iform_max_per_iclass")
	proc_xed_iform_first_per_iclass = dll.NewProc("xed_iform_first_per_iclass")
	proc_xed_iform_to_category = dll.NewProc("xed_iform_to_category")
	proc_xed_iform_to_extension = dll.NewProc("xed_iform_to_extension")
	proc_xed_iform_to_isa_set = dll.NewProc("xed_iform_to_isa_set")
	proc_xed_iform_to_iclass_string_att = dll.NewProc("xed_iform_to_iclass_string_att")
	proc_xed_iform_to_iclass_string_intel = dll.NewProc("xed_iform_to_iclass_string_intel")
	proc_xed_operand_width_bits = dll.NewProc("xed_operand_width_bits")
	proc_xed_operand_print = dll.NewProc("xed_operand_print")
	proc_xed_operand_read = dll.NewProc("xed_operand_read")
	proc_xed_operand_read_only = dll.NewProc("xed_operand_read_only")
	proc_xed_operand_written = dll.NewProc("xed_operand_written")
	proc_xed_operand_written_only = dll.NewProc("xed_operand_written_only")
	proc_xed_operand_read_and_written = dll.NewProc("xed_operand_read_and_written")
	proc_xed_operand_conditional_read = dll.NewProc("xed_operand_conditional_read")
	proc_xed_operand_conditional_write = dll.NewProc("xed_operand_conditional_write")
	proc_xed_inst_cpl = dll.NewProc("xed_inst_cpl")
	proc_xed_inst_operand = dll.NewProc("xed_inst_operand")
	proc_xed_inst_flag_info_index = dll.NewProc("xed_inst_flag_info_index")
	proc_xed_inst_get_attribute = dll.NewProc("xed_inst_get_attribute")
	proc_xed_inst_get_attributes = dll.NewProc("xed_inst_get_attributes")
	proc_xed_attribute_max = dll.NewProc("xed_attribute_max")
	proc_xed_attribute = dll.NewProc("xed_attribute")
	proc_xed_inst_table_base = dll.NewProc("xed_inst_table_base")
	proc_xed_flag_enum_t2str = dll.NewProc("xed_flag_enum_t2str")
	proc_xed_flag_enum_t_last = dll.NewProc("xed_flag_enum_t_last")
	proc_xed_flag_action_enum_t2str = dll.NewProc("xed_flag_action_enum_t2str")
	proc_xed_flag_action_enum_t_last = dll.NewProc("xed_flag_action_enum_t_last")
	proc_xed_flag_set_print = dll.NewProc("xed_flag_set_print")
	proc_xed_flag_set_is_subset_of = dll.NewProc("xed_flag_set_is_subset_of")
	proc_xed_flag_action_get_flag_name = dll.NewProc("xed_flag_action_get_flag_name")
	proc_xed_flag_action_get_action = dll.NewProc("xed_flag_action_get_action")
	proc_xed_flag_action_action_invalid = dll.NewProc("xed_flag_action_action_invalid")
	proc_xed_flag_action_print = dll.NewProc("xed_flag_action_print")
	proc_xed_flag_action_read_flag = dll.NewProc("xed_flag_action_read_flag")
	proc_xed_flag_action_writes_flag = dll.NewProc("xed_flag_action_writes_flag")
	proc_xed_flag_action_read_action = dll.NewProc("xed_flag_action_read_action")
	proc_xed_flag_action_write_action = dll.NewProc("xed_flag_action_write_action")
	proc_xed_simple_flag_get_nflags = dll.NewProc("xed_simple_flag_get_nflags")
	proc_xed_simple_flag_get_read_flag_set = dll.NewProc("xed_simple_flag_get_read_flag_set")
	proc_xed_simple_flag_get_written_flag_set = dll.NewProc("xed_simple_flag_get_written_flag_set")
	proc_xed_simple_flag_get_undefined_flag_set = dll.NewProc("xed_simple_flag_get_undefined_flag_set")
	proc_xed_simple_flag_get_may_write = dll.NewProc("xed_simple_flag_get_may_write")
	proc_xed_simple_flag_get_must_write = dll.NewProc("xed_simple_flag_get_must_write")
	proc_xed_simple_flag_get_flag_action = dll.NewProc("xed_simple_flag_get_flag_action")
	proc_xed_simple_flag_reads_flags = dll.NewProc("xed_simple_flag_reads_flags")
	proc_xed_simple_flag_writes_flags = dll.NewProc("xed_simple_flag_writes_flags")
	proc_xed_simple_flag_print = dll.NewProc("xed_simple_flag_print")
	proc_xed_chip_enum_t2str = dll.NewProc("xed_chip_enum_t2str")
	proc_xed_chip_enum_t_last = dll.NewProc("xed_chip_enum_t_last")
	proc_xed_operand_element_type_enum_t2str = dll.NewProc("xed_operand_element_type_enum_t2str")
	proc_xed_operand_element_type_enum_t_last = dll.NewProc("xed_operand_element_type_enum_t_last")
	proc_xed_error_enum_t2str = dll.NewProc("xed_error_enum_t2str")
	proc_xed_error_enum_t_last = dll.NewProc("xed_error_enum_t_last")
	proc_xed_get_chip_features = dll.NewProc("xed_get_chip_features")
	proc_xed_modify_chip_features = dll.NewProc("xed_modify_chip_features")
	proc_xed_set_decoder_modes = dll.NewProc("xed_set_decoder_modes")
	proc_xed_decode = dll.NewProc("xed_decode")
	proc_xed_decode_with_features = dll.NewProc("xed_decode_with_features")
	proc_xed_ild_decode = dll.NewProc("xed_ild_decode")
	proc_xed_address_width_enum_t2str = dll.NewProc("xed_address_width_enum_t2str")
	proc_xed_address_width_enum_t_last = dll.NewProc("xed_address_width_enum_t_last")
	proc_xed_machine_mode_enum_t2str = dll.NewProc("xed_machine_mode_enum_t2str")
	proc_xed_machine_mode_enum_t_last = dll.NewProc("xed_machine_mode_enum_t_last")
	proc_xed_state_print = dll.NewProc("xed_state_print")
	proc_xed_syntax_enum_t2str = dll.NewProc("xed_syntax_enum_t2str")
	proc_xed_syntax_enum_t_last = dll.NewProc("xed_syntax_enum_t_last")
	proc_xed_reg_class_enum_t2str = dll.NewProc("xed_reg_class_enum_t2str")
	proc_xed_reg_class_enum_t_last = dll.NewProc("xed_reg_class_enum_t_last")
	proc_xed_reg_class = dll.NewProc("xed_reg_class")
	proc_xed_gpr_reg_class = dll.NewProc("xed_gpr_reg_class")
	proc_xed_get_largest_enclosing_register = dll.NewProc("xed_get_largest_enclosing_register")
	proc_xed_get_largest_enclosing_register32 = dll.NewProc("xed_get_largest_enclosing_register32")
	proc_xed_get_register_width_bits = dll.NewProc("xed_get_register_width_bits")
	proc_xed_get_register_width_bits64 = dll.NewProc("xed_get_register_width_bits64")
	proc_xed_operand_values_init = dll.NewProc("xed_operand_values_init")
	proc_xed_operand_values_init_set_mode = dll.NewProc("xed_operand_values_init_set_mode")
	proc_xed_operand_values_set_mode = dll.NewProc("xed_operand_values_set_mode")
	proc_xed_operand_values_init_keep_mode = dll.NewProc("xed_operand_values_init_keep_mode")
	proc_xed_operand_values_dump = dll.NewProc("xed_operand_values_dump")
	proc_xed_operand_values_print_short = dll.NewProc("xed_operand_values_print_short")
	proc_xed_operand_values_has_real_rep = dll.NewProc("xed_operand_values_has_real_rep")
	proc_xed_operand_values_has_rep_prefix = dll.NewProc("xed_operand_values_has_rep_prefix")
	proc_xed_operand_values_has_repne_prefix = dll.NewProc("xed_operand_values_has_repne_prefix")
	proc_xed_operand_values_clear_rep = dll.NewProc("xed_operand_values_clear_rep")
	proc_xed_operand_values_get_atomic = dll.NewProc("xed_operand_values_get_atomic")
	proc_xed_operand_values_has_lock_prefix = dll.NewProc("xed_operand_values_has_lock_prefix")
	proc_xed_operand_values_lockable = dll.NewProc("xed_operand_values_lockable")
	proc_xed_operand_values_using_default_segment = dll.NewProc("xed_operand_values_using_default_segment")
	proc_xed_operand_values_get_effective_operand_width = dll.NewProc("xed_operand_values_get_effective_operand_width")
	proc_xed_operand_values_get_effective_address_width = dll.NewProc("xed_operand_values_get_effective_address_width")
	proc_xed_operand_values_get_stack_address_width = dll.NewProc("xed_operand_values_get_stack_address_width")
	proc_xed_operand_values_has_memory_displacement = dll.NewProc("xed_operand_values_has_memory_displacement")
	proc_xed_operand_values_has_branch_displacement = dll.NewProc("xed_operand_values_has_branch_displacement")
	proc_xed_operand_values_has_displacement = dll.NewProc("xed_operand_values_has_displacement")
	proc_xed_operand_values_get_displacement_for_memop = dll.NewProc("xed_operand_values_get_displacement_for_memop")
	proc_xed_operand_values_has_immediate = dll.NewProc("xed_operand_values_has_immediate")
	proc_xed_operand_values_has_address_size_prefix = dll.NewProc("xed_operand_values_has_address_size_prefix")
	proc_xed_operand_values_has_operand_size_prefix = dll.NewProc("xed_operand_values_has_operand_size_prefix")
	proc_xed_operand_values_has_66_prefix = dll.NewProc("xed_operand_values_has_66_prefix")
	proc_xed_operand_values_mandatory_66_prefix = dll.NewProc("xed_operand_values_mandatory_66_prefix")
	proc_xed_operand_values_has_rexw_prefix = dll.NewProc("xed_operand_values_has_rexw_prefix")
	proc_xed_operand_values_has_segment_prefix = dll.NewProc("xed_operand_values_has_segment_prefix")
	proc_xed_operand_values_segment_prefix = dll.NewProc("xed_operand_values_segment_prefix")
	proc_xed_operand_values_get_long_mode = dll.NewProc("xed_operand_values_get_long_mode")
	proc_xed_operand_values_get_real_mode = dll.NewProc("xed_operand_values_get_real_mode")
	proc_xed_operand_values_get_pp_vex_prefix = dll.NewProc("xed_operand_values_get_pp_vex_prefix")
	proc_xed_operand_values_accesses_memory = dll.NewProc("xed_operand_values_accesses_memory")
	proc_xed_operand_values_number_of_memory_operands = dll.NewProc("xed_operand_values_number_of_memory_operands")
	proc_xed_operand_values_get_memory_operand_length = dll.NewProc("xed_operand_values_get_memory_operand_length")
	proc_xed_operand_values_get_base_reg = dll.NewProc("xed_operand_values_get_base_reg")
	proc_xed_operand_values_get_index_reg = dll.NewProc("xed_operand_values_get_index_reg")
	proc_xed_operand_values_get_seg_reg = dll.NewProc("xed_operand_values_get_seg_reg")
	proc_xed_operand_values_get_scale = dll.NewProc("xed_operand_values_get_scale")
	proc_xed_operand_values_memop_without_modrm = dll.NewProc("xed_operand_values_memop_without_modrm")
	proc_xed_operand_values_has_modrm_byte = dll.NewProc("xed_operand_values_has_modrm_byte")
	proc_xed_operand_values_has_sib_byte = dll.NewProc("xed_operand_values_has_sib_byte")
	proc_xed_operand_values_branch_not_taken_hint = dll.NewProc("xed_operand_values_branch_not_taken_hint")
	proc_xed_operand_values_branch_taken_hint = dll.NewProc("xed_operand_values_branch_taken_hint")
	proc_xed_operand_values_ignored_branch_not_taken_hint = dll.NewProc("xed_operand_values_ignored_branch_not_taken_hint")
	proc_xed_operand_values_ignored_branch_taken_hint = dll.NewProc("xed_operand_values_ignored_branch_taken_hint")
	proc_xed_operand_values_cet_no_track = dll.NewProc("xed_operand_values_cet_no_track")
	proc_xed_operand_values_is_nop = dll.NewProc("xed_operand_values_is_nop")
	proc_xed_operand_values_get_immediate_int64 = dll.NewProc("xed_operand_values_get_immediate_int64")
	proc_xed_operand_values_get_immediate_uint64 = dll.NewProc("xed_operand_values_get_immediate_uint64")
	proc_xed_operand_values_get_immediate_is_signed = dll.NewProc("xed_operand_values_get_immediate_is_signed")
	proc_xed_operand_values_get_immediate_byte = dll.NewProc("xed_operand_values_get_immediate_byte")
	proc_xed_operand_values_get_second_immediate = dll.NewProc("xed_operand_values_get_second_immediate")
	proc_xed_operand_values_get_memory_displacement_length = dll.NewProc("xed_operand_values_get_memory_displacement_length")
	proc_xed_operand_values_get_memory_displacement_length_bits = dll.NewProc("xed_operand_values_get_memory_displacement_length_bits")
	proc_xed_operand_values_get_memory_displacement_length_bits_raw = dll.NewProc("xed_operand_values_get_memory_displacement_length_bits_raw")
	proc_xed_operand_values_get_memory_displacement_int64 = dll.NewProc("xed_operand_values_get_memory_displacement_int64")
	proc_xed_operand_values_get_memory_displacement_int64_raw = dll.NewProc("xed_operand_values_get_memory_displacement_int64_raw")
	proc_xed_operand_values_get_memory_displacement_byte = dll.NewProc("xed_operand_values_get_memory_displacement_byte")
	proc_xed_operand_values_get_branch_displacement_length = dll.NewProc("xed_operand_values_get_branch_displacement_length")
	proc_xed_operand_values_get_branch_displacement_length_bits = dll.NewProc("xed_operand_values_get_branch_displacement_length_bits")
	proc_xed_operand_values_get_branch_displacement_int64 = dll.NewProc("xed_operand_values_get_branch_displacement_int64")
	proc_xed_operand_values_get_branch_displacement_byte = dll.NewProc("xed_operand_values_get_branch_displacement_byte")
	proc_xed_operand_values_get_iclass = dll.NewProc("xed_operand_values_get_iclass")
	proc_xed_operand_values_zero_immediate = dll.NewProc("xed_operand_values_zero_immediate")
	proc_xed_operand_values_zero_branch_displacement = dll.NewProc("xed_operand_values_zero_branch_displacement")
	proc_xed_operand_values_zero_memory_displacement = dll.NewProc("xed_operand_values_zero_memory_displacement")
	proc_xed_operand_values_set_lock = dll.NewProc("xed_operand_values_set_lock")
	proc_xed_operand_values_zero_segment_override = dll.NewProc("xed_operand_values_zero_segment_override")
	proc_xed_operand_values_set_iclass = dll.NewProc("xed_operand_values_set_iclass")
	proc_xed_operand_values_set_effective_operand_width = dll.NewProc("xed_operand_values_set_effective_operand_width")
	proc_xed_operand_values_set_effective_address_width = dll.NewProc("xed_operand_values_set_effective_address_width")
	proc_xed_operand_values_set_memory_operand_length = dll.NewProc("xed_operand_values_set_memory_operand_length")
	proc_xed_operand_values_set_memory_displacement = dll.NewProc("xed_operand_values_set_memory_displacement")
	proc_xed_operand_values_set_memory_displacement_bits = dll.NewProc("xed_operand_values_set_memory_displacement_bits")
	proc_xed_operand_values_set_relbr = dll.NewProc("xed_operand_values_set_relbr")
	proc_xed_operand_values_set_absbr = dll.NewProc("xed_operand_values_set_absbr")
	proc_xed_operand_values_set_branch_displacement = dll.NewProc("xed_operand_values_set_branch_displacement")
	proc_xed_operand_values_set_branch_displacement_bits = dll.NewProc("xed_operand_values_set_branch_displacement_bits")
	proc_xed_operand_values_set_immediate_signed = dll.NewProc("xed_operand_values_set_immediate_signed")
	proc_xed_operand_values_set_immediate_signed_bits = dll.NewProc("xed_operand_values_set_immediate_signed_bits")
	proc_xed_operand_values_set_immediate_unsigned = dll.NewProc("xed_operand_values_set_immediate_unsigned")
	proc_xed_operand_values_set_immediate_unsigned_bits = dll.NewProc("xed_operand_values_set_immediate_unsigned_bits")
	proc_xed_operand_values_set_base_reg = dll.NewProc("xed_operand_values_set_base_reg")
	proc_xed_operand_values_set_seg_reg = dll.NewProc("xed_operand_values_set_seg_reg")
	proc_xed_operand_values_set_index_reg = dll.NewProc("xed_operand_values_set_index_reg")
	proc_xed_operand_values_set_scale = dll.NewProc("xed_operand_values_set_scale")
	proc_xed_operand_values_set_operand_reg = dll.NewProc("xed_operand_values_set_operand_reg")
	proc_xed_encoder_request_get_iclass = dll.NewProc("xed_encoder_request_get_iclass")
	proc_xed_encoder_request_set_iclass = dll.NewProc("xed_encoder_request_set_iclass")
	proc_xed_encoder_request_set_effective_operand_width = dll.NewProc("xed_encoder_request_set_effective_operand_width")
	proc_xed_encoder_request_set_effective_address_size = dll.NewProc("xed_encoder_request_set_effective_address_size")
	proc_xed_encoder_request_set_reg = dll.NewProc("xed_encoder_request_set_reg")
	proc_xed_encoder_request_set_operand_order = dll.NewProc("xed_encoder_request_set_operand_order")
	proc_xed_encoder_request_get_operand_order = dll.NewProc("xed_encoder_request_get_operand_order")
	proc_xed_encoder_request_set_relbr = dll.NewProc("xed_encoder_request_set_relbr")
	proc_xed_encoder_request_set_absbr = dll.NewProc("xed_encoder_request_set_absbr")
	proc_xed_encoder_request_set_branch_displacement = dll.NewProc("xed_encoder_request_set_branch_displacement")
	proc_xed_encoder_request_set_ptr = dll.NewProc("xed_encoder_request_set_ptr")
	proc_xed_encoder_request_set_uimm0 = dll.NewProc("xed_encoder_request_set_uimm0")
	proc_xed_encoder_request_set_uimm0_bits = dll.NewProc("xed_encoder_request_set_uimm0_bits")
	proc_xed_encoder_request_set_uimm1 = dll.NewProc("xed_encoder_request_set_uimm1")
	proc_xed_encoder_request_set_simm = dll.NewProc("xed_encoder_request_set_simm")
	proc_xed_encoder_request_set_memory_displacement = dll.NewProc("xed_encoder_request_set_memory_displacement")
	proc_xed_encoder_request_set_agen = dll.NewProc("xed_encoder_request_set_agen")
	proc_xed_encoder_request_set_mem0 = dll.NewProc("xed_encoder_request_set_mem0")
	proc_xed_encoder_request_set_mem1 = dll.NewProc("xed_encoder_request_set_mem1")
	proc_xed_encoder_request_set_memory_operand_length = dll.NewProc("xed_encoder_request_set_memory_operand_length")
	proc_xed_encoder_request_set_seg0 = dll.NewProc("xed_encoder_request_set_seg0")
	proc_xed_encoder_request_set_seg1 = dll.NewProc("xed_encoder_request_set_seg1")
	proc_xed_encoder_request_set_base0 = dll.NewProc("xed_encoder_request_set_base0")
	proc_xed_encoder_request_set_base1 = dll.NewProc("xed_encoder_request_set_base1")
	proc_xed_encoder_request_set_index = dll.NewProc("xed_encoder_request_set_index")
	proc_xed_encoder_request_set_scale = dll.NewProc("xed_encoder_request_set_scale")
	proc_xed_encoder_request_zero_operand_order = dll.NewProc("xed_encoder_request_zero_operand_order")
	proc_xed_encoder_request_zero_set_mode = dll.NewProc("xed_encoder_request_zero_set_mode")
	proc_xed_encoder_request_zero = dll.NewProc("xed_encoder_request_zero")
	proc_xed_encoder_request_init_from_decode = dll.NewProc("xed_encoder_request_init_from_decode")
	proc_xed_encode_request_print = dll.NewProc("xed_encode_request_print")
	proc_xed_encode = dll.NewProc("xed_encode")
	proc_xed_encode_nop = dll.NewProc("xed_encode_nop")
	proc_xed_convert_to_encoder_request = dll.NewProc("xed_convert_to_encoder_request")
	proc_xed_patch_disp = dll.NewProc("xed_patch_disp")
	proc_xed_patch_brdisp = dll.NewProc("xed_patch_brdisp")
	proc_xed_patch_imm0 = dll.NewProc("xed_patch_imm0")
	proc_xed_operand_action_read = dll.NewProc("xed_operand_action_read")
	proc_xed_operand_action_read_only = dll.NewProc("xed_operand_action_read_only")
	proc_xed_operand_action_written = dll.NewProc("xed_operand_action_written")
	proc_xed_operand_action_written_only = dll.NewProc("xed_operand_action_written_only")
	proc_xed_operand_action_read_and_written = dll.NewProc("xed_operand_action_read_and_written")
	proc_xed_operand_action_conditional_read = dll.NewProc("xed_operand_action_conditional_read")
	proc_xed_operand_action_conditional_write = dll.NewProc("xed_operand_action_conditional_write")
	proc_xed_get_version = dll.NewProc("xed_get_version")
	proc_xed_get_copyright = dll.NewProc("xed_get_copyright")
	proc_xed_format_set_options = dll.NewProc("xed_format_set_options")
	proc_xed_init_print_info = dll.NewProc("xed_init_print_info")
	proc_xed_decoded_inst_get_attribute = dll.NewProc("xed_decoded_inst_get_attribute")
	proc_xed_decoded_inst_get_attributes = dll.NewProc("xed_decoded_inst_get_attributes")
	proc_xed_decoded_inst_is_xacquire = dll.NewProc("xed_decoded_inst_is_xacquire")
	proc_xed_decoded_inst_is_xrelease = dll.NewProc("xed_decoded_inst_is_xrelease")
	proc_xed_decoded_inst_has_mpx_prefix = dll.NewProc("xed_decoded_inst_has_mpx_prefix")
	proc_xed_decoded_inst_is_apx_zu = dll.NewProc("xed_decoded_inst_is_apx_zu")
	proc_xed_decoded_inst_get_modrm = dll.NewProc("xed_decoded_inst_get_modrm")
	proc_xed_decoded_inst_masked_vector_operation = dll.NewProc("xed_decoded_inst_masked_vector_operation")
	proc_xed_decoded_inst_vector_length_bits = dll.NewProc("xed_decoded_inst_vector_length_bits")
	proc_xed_decoded_inst_get_nprefixes = dll.NewProc("xed_decoded_inst_get_nprefixes")
	proc_xed_decoded_inst_operand_length_bits = dll.NewProc("xed_decoded_inst_operand_length_bits")
	proc_xed_decoded_inst_operand_length = dll.NewProc("xed_decoded_inst_operand_length")
	proc_xed_decoded_inst_operand_elements = dll.NewProc("xed_decoded_inst_operand_elements")
	proc_xed_decoded_inst_operand_element_size_bits = dll.NewProc("xed_decoded_inst_operand_element_size_bits")
	proc_xed_decoded_inst_operand_element_type = dll.NewProc("xed_decoded_inst_operand_element_type")
	proc_xed_decoded_inst_operand_action = dll.NewProc("xed_decoded_inst_operand_action")
	proc_xed_decoded_inst_masking = dll.NewProc("xed_decoded_inst_masking")
	proc_xed_decoded_inst_merging = dll.NewProc("xed_decoded_inst_merging")
	proc_xed_decoded_inst_zeroing = dll.NewProc("xed_decoded_inst_zeroing")
	proc_xed_decoded_inst_avx512_dest_elements = dll.NewProc("xed_decoded_inst_avx512_dest_elements")
	proc_xed_decoded_inst_zero_set_mode = dll.NewProc("xed_decoded_inst_zero_set_mode")
	proc_xed_decoded_inst_zero_keep_mode = dll.NewProc("xed_decoded_inst_zero_keep_mode")
	proc_xed_decoded_inst_zero = dll.NewProc("xed_decoded_inst_zero")
	proc_xed_decoded_inst_zero_keep_mode_from_operands = dll.NewProc("xed_decoded_inst_zero_keep_mode_from_operands")
	proc_xed_decoded_inst_get_operand_width = dll.NewProc("xed_decoded_inst_get_operand_width")
	proc_xed_decoded_inst_valid_for_chip = dll.NewProc("xed_decoded_inst_valid_for_chip")
	proc_xed_decoded_inst_valid_for_features = dll.NewProc("xed_decoded_inst_valid_for_features")
	proc_xed_decoded_inst_dump = dll.NewProc("xed_decoded_inst_dump")
	proc_xed_decoded_inst_dump_xed_format = dll.NewProc("xed_decoded_inst_dump_xed_format")
	proc_xed_format_context = dll.NewProc("xed_format_context")
	proc_xed_format_generic = dll.NewProc("xed_format_generic")
	proc_xed_decoded_inst_get_seg_reg = dll.NewProc("xed_decoded_inst_get_seg_reg")
	proc_xed_decoded_inst_get_base_reg = dll.NewProc("xed_decoded_inst_get_base_reg")
	proc_xed_decoded_inst_get_index_reg = dll.NewProc("xed_decoded_inst_get_index_reg")
	proc_xed_decoded_inst_get_scale = dll.NewProc("xed_decoded_inst_get_scale")
	proc_xed_decoded_inst_get_memory_displacement = dll.NewProc("xed_decoded_inst_get_memory_displacement")
	proc_xed_decoded_inst_get_memory_displacement_width = dll.NewProc("xed_decoded_inst_get_memory_displacement_width")
	proc_xed_decoded_inst_get_memory_displacement_width_bits = dll.NewProc("xed_decoded_inst_get_memory_displacement_width_bits")
	proc_xed_decoded_inst_get_branch_displacement = dll.NewProc("xed_decoded_inst_get_branch_displacement")
	proc_xed_decoded_inst_get_branch_displacement_width = dll.NewProc("xed_decoded_inst_get_branch_displacement_width")
	proc_xed_decoded_inst_get_branch_displacement_width_bits = dll.NewProc("xed_decoded_inst_get_branch_displacement_width_bits")
	proc_xed_decoded_inst_get_unsigned_immediate = dll.NewProc("xed_decoded_inst_get_unsigned_immediate")
	proc_xed_decoded_inst_get_immediate_is_signed = dll.NewProc("xed_decoded_inst_get_immediate_is_signed")
	proc_xed_decoded_inst_get_immediate_width = dll.NewProc("xed_decoded_inst_get_immediate_width")
	proc_xed_decoded_inst_get_immediate_width_bits = dll.NewProc("xed_decoded_inst_get_immediate_width_bits")
	proc_xed_decoded_inst_get_signed_immediate = dll.NewProc("xed_decoded_inst_get_signed_immediate")
	proc_xed_decoded_inst_get_reg = dll.NewProc("xed_decoded_inst_get_reg")
	proc_xed_decoded_inst_has_default_flags_values = dll.NewProc("xed_decoded_inst_has_default_flags_values")
	proc_xed_decoded_inst_get_default_flags_values = dll.NewProc("xed_decoded_inst_get_default_flags_values")
	proc_xed_decoded_inst_get_rflags_info = dll.NewProc("xed_decoded_inst_get_rflags_info")
	proc_xed_decoded_inst_uses_rflags = dll.NewProc("xed_decoded_inst_uses_rflags")
	proc_xed_decoded_inst_number_of_memory_operands = dll.NewProc("xed_decoded_inst_number_of_memory_operands")
	proc_xed_decoded_inst_mem_read = dll.NewProc("xed_decoded_inst_mem_read")
	proc_xed_decoded_inst_mem_written = dll.NewProc("xed_decoded_inst_mem_written")
	proc_xed_decoded_inst_mem_written_only = dll.NewProc("xed_decoded_inst_mem_written_only")
	proc_xed_decoded_inst_conditionally_writes_registers = dll.NewProc("xed_decoded_inst_conditionally_writes_registers")
	proc_xed_decoded_inst_get_memory_operand_length = dll.NewProc("xed_decoded_inst_get_memory_operand_length")
	proc_xed_decoded_inst_get_memop_address_width = dll.NewProc("xed_decoded_inst_get_memop_address_width")
	proc_xed_decoded_inst_is_prefetch = dll.NewProc("xed_decoded_inst_is_prefetch")
	proc_xed_decoded_inst_is_broadcast = dll.NewProc("xed_decoded_inst_is_broadcast")
	proc_xed_decoded_inst_is_broadcast_instruction = dll.NewProc("xed_decoded_inst_is_broadcast_instruction")
	proc_xed_decoded_inst_uses_embedded_broadcast = dll.NewProc("xed_decoded_inst_uses_embedded_broadcast")
	proc_xed_decoded_inst_set_scale = dll.NewProc("xed_decoded_inst_set_scale")
	proc_xed_decoded_inst_set_memory_displacement = dll.NewProc("xed_decoded_inst_set_memory_displacement")
	proc_xed_decoded_inst_set_branch_displacement = dll.NewProc("xed_decoded_inst_set_branch_displacement")
	proc_xed_decoded_inst_set_immediate_signed = dll.NewProc("xed_decoded_inst_set_immediate_signed")
	proc_xed_decoded_inst_set_immediate_unsigned = dll.NewProc("xed_decoded_inst_set_immediate_unsigned")
	proc_xed_decoded_inst_set_memory_displacement_bits = dll.NewProc("xed_decoded_inst_set_memory_displacement_bits")
	proc_xed_decoded_inst_set_branch_displacement_bits = dll.NewProc("xed_decoded_inst_set_branch_displacement_bits")
	proc_xed_decoded_inst_set_immediate_signed_bits = dll.NewProc("xed_decoded_inst_set_immediate_signed_bits")
	proc_xed_decoded_inst_set_immediate_unsigned_bits = dll.NewProc("xed_decoded_inst_set_immediate_unsigned_bits")
	proc_xed_classify_amx = dll.NewProc("xed_classify_amx")
	proc_xed_classify_avx512 = dll.NewProc("xed_classify_avx512")
	proc_xed_classify_avx512_maskop = dll.NewProc("xed_classify_avx512_maskop")
	proc_xed_classify_avx = dll.NewProc("xed_classify_avx")
	proc_xed_classify_sse = dll.NewProc("xed_classify_sse")
	proc_xed_classify_apx_foundation = dll.NewProc("xed_classify_apx_foundation")
	proc_xed_classify_apx = dll.NewProc("xed_classify_apx")
	proc_xed_iformfl_enum_t_last = dll.NewProc("xed_iformfl_enum_t_last")
	proc_xed_rep_remove = dll.NewProc("xed_rep_remove")
	proc_xed_repe_map = dll.NewProc("xed_repe_map")
	proc_xed_repne_map = dll.NewProc("xed_repne_map")
	proc_xed_rep_map = dll.NewProc("xed_rep_map")
	proc_xed_norep_map = dll.NewProc("xed_norep_map")
	proc_xed_agen_register_callback = dll.NewProc("xed_agen_register_callback")
	proc_xed_agen = dll.NewProc("xed_agen")
	proc_xed_cpuid_rec_enum_t2str = dll.NewProc("xed_cpuid_rec_enum_t2str")
	proc_xed_cpuid_rec_enum_t_last = dll.NewProc("xed_cpuid_rec_enum_t_last")
	proc_xed_cpuid_group_enum_t2str = dll.NewProc("xed_cpuid_group_enum_t2str")
	proc_xed_cpuid_group_enum_t_last = dll.NewProc("xed_cpuid_group_enum_t_last")
	proc_xed_get_cpuid_group_enum_for_isa_set = dll.NewProc("xed_get_cpuid_group_enum_for_isa_set")
	proc_xed_get_cpuid_rec_enum_for_group = dll.NewProc("xed_get_cpuid_rec_enum_for_group")
	proc_xed_get_cpuid_rec = dll.NewProc("xed_get_cpuid_rec")
	proc_xed_isa_set_is_valid_for_chip = dll.NewProc("xed_isa_set_is_valid_for_chip")
}

func saveEmbeddedDLL(data []byte, name string) string {
	tmpDir := os.TempDir()
	p := filepath.Join(tmpDir, name)
	os.WriteFile(p, data, 0o644)
	return p
}

func (x *Xed) Strlen(S *int8) Xed_uint_t {
	r1, _, _ := proc_xed_strlen.Call(uintptr(unsafe.Pointer(S)))
	return uint32(r1)
}

func (x *Xed) Strcat(Dst *int8, Src *int8) {
	proc_xed_strcat.Call(uintptr(unsafe.Pointer(Dst)), uintptr(unsafe.Pointer(Src)))
}

func (x *Xed) Strcpy(Dst *int8, Src *int8) {
	proc_xed_strcpy.Call(uintptr(unsafe.Pointer(Dst)), uintptr(unsafe.Pointer(Src)))
}

func (x *Xed) Strncpy(Dst *int8, Src *int8, Len int32) int32 {
	r1, _, _ := proc_xed_strncpy.Call(uintptr(unsafe.Pointer(Dst)), uintptr(unsafe.Pointer(Src)), uintptr(Len))
	return int32(r1)
}

func (x *Xed) Strncat(Dst *int8, Src *int8, Len int32) int32 {
	r1, _, _ := proc_xed_strncat.Call(uintptr(unsafe.Pointer(Dst)), uintptr(unsafe.Pointer(Src)), uintptr(Len))
	return int32(r1)
}

func (x *Xed) OperandEnumT2str(P Xed_operand_enum_t) *int8 {
	r1, _, _ := proc_xed_operand_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandEnumTLast() Xed_operand_enum_t {
	r1, _, _ := proc_xed_operand_enum_t_last.Call()
	return Xed_operand_enum_t(uint32(r1))
}

func (x *Xed) TablesInit() {
	proc_xed_tables_init.Call()
}

func (x *Xed) InternalAssert(S *int8, File *int8, Line int32) {
	proc_xed_internal_assert.Call(uintptr(unsafe.Pointer(S)), uintptr(unsafe.Pointer(File)), uintptr(Line))
}

func (x *Xed) RegisterAbortFunction(Fn Xed_user_abort_function_t, Other unsafe.Pointer) {
	proc_xed_register_abort_function.Call(windows.NewCallback(Fn), uintptr(Other))
}

func (x *Xed) Itoa(Buf *int8, F Xed_uint64_t, Buflen int32) int32 {
	r1, _, _ := proc_xed_itoa.Call(uintptr(unsafe.Pointer(Buf)), *(*uintptr)(unsafe.Pointer(&F)), uintptr(Buflen))
	return int32(r1)
}

func (x *Xed) ItoaBin(Buf *int8, F Xed_uint64_t, Bits_to_print Xed_uint_t, Buflen Xed_uint_t) int32 {
	r1, _, _ := proc_xed_itoa_bin.Call(uintptr(unsafe.Pointer(Buf)), *(*uintptr)(unsafe.Pointer(&F)), uintptr(Bits_to_print), uintptr(Buflen))
	return int32(r1)
}

func (x *Xed) ItoaHexZeros(Buf *int8, F Xed_uint64_t, Bits_to_print Xed_uint_t, Leading_zeros Xed_bool_t, Buflen int32) int32 {
	r1, _, _ := proc_xed_itoa_hex_zeros.Call(uintptr(unsafe.Pointer(Buf)), *(*uintptr)(unsafe.Pointer(&F)), uintptr(Bits_to_print), uintptr(Leading_zeros), uintptr(Buflen))
	return int32(r1)
}

func (x *Xed) ItoaHex(Buf *int8, F Xed_uint64_t, Bits_to_print Xed_uint_t, Buflen int32) int32 {
	r1, _, _ := proc_xed_itoa_hex.Call(uintptr(unsafe.Pointer(Buf)), *(*uintptr)(unsafe.Pointer(&F)), uintptr(Bits_to_print), uintptr(Buflen))
	return int32(r1)
}

func (x *Xed) ItoaHexUl(Buf *int8, F Xed_uint64_t, Bits_to_print Xed_uint_t, Leading_zeros Xed_bool_t, Buflen int32, Lowercase Xed_bool_t) int32 {
	r1, _, _ := proc_xed_itoa_hex_ul.Call(uintptr(unsafe.Pointer(Buf)), *(*uintptr)(unsafe.Pointer(&F)), uintptr(Bits_to_print), uintptr(Leading_zeros), uintptr(Buflen), uintptr(Lowercase))
	return int32(r1)
}

func (x *Xed) SetLogFile(O unsafe.Pointer) {
	proc_xed_set_log_file.Call(uintptr(O))
}

func (x *Xed) SetVerbosity(V int32) {
	proc_xed_set_verbosity.Call(uintptr(V))
}

func (x *Xed) SignExtend3264(X Xed_int32_t) Xed_int64_t {
	r1, _, _ := proc_xed_sign_extend32_64.Call(uintptr(X))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) SignExtend1664(X Xed_int16_t) Xed_int64_t {
	r1, _, _ := proc_xed_sign_extend16_64.Call(uintptr(X))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) SignExtend864(X Xed_int8_t) Xed_int64_t {
	r1, _, _ := proc_xed_sign_extend8_64.Call(uintptr(X))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) SignExtend1632(X Xed_int16_t) Xed_int32_t {
	r1, _, _ := proc_xed_sign_extend16_32.Call(uintptr(X))
	return int32(r1)
}

func (x *Xed) SignExtend832(X Xed_int8_t) Xed_int32_t {
	r1, _, _ := proc_xed_sign_extend8_32.Call(uintptr(X))
	return int32(r1)
}

func (x *Xed) SignExtend816(X Xed_int8_t) Xed_int16_t {
	r1, _, _ := proc_xed_sign_extend8_16.Call(uintptr(X))
	return int16(r1)
}

func (x *Xed) SignExtendArbitraryTo32(X Xed_uint32_t, Bits uint32) Xed_int32_t {
	r1, _, _ := proc_xed_sign_extend_arbitrary_to_32.Call(uintptr(X), uintptr(Bits))
	return int32(r1)
}

func (x *Xed) SignExtendArbitraryTo64(X Xed_uint64_t, Bits uint32) Xed_int64_t {
	r1, _, _ := proc_xed_sign_extend_arbitrary_to_64.Call(*(*uintptr)(unsafe.Pointer(&X)), uintptr(Bits))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) ZeroExtend3264(X Xed_uint32_t) Xed_uint64_t {
	r1, _, _ := proc_xed_zero_extend32_64.Call(uintptr(X))
	return *(*Xed_uint64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) ZeroExtend1664(X Xed_uint16_t) Xed_uint64_t {
	r1, _, _ := proc_xed_zero_extend16_64.Call(uintptr(X))
	return *(*Xed_uint64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) ZeroExtend864(X Xed_uint8_t) Xed_uint64_t {
	r1, _, _ := proc_xed_zero_extend8_64.Call(uintptr(X))
	return *(*Xed_uint64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) ZeroExtend1632(X Xed_uint16_t) Xed_uint32_t {
	r1, _, _ := proc_xed_zero_extend16_32.Call(uintptr(X))
	return uint32(r1)
}

func (x *Xed) ZeroExtend832(X Xed_uint8_t) Xed_uint32_t {
	r1, _, _ := proc_xed_zero_extend8_32.Call(uintptr(X))
	return uint32(r1)
}

func (x *Xed) ZeroExtend816(X Xed_uint8_t) Xed_uint16_t {
	r1, _, _ := proc_xed_zero_extend8_16.Call(uintptr(X))
	return uint16(r1)
}

func (x *Xed) GetByte(X Xed_uint64_t, I uint32, Len uint32) Xed_uint8_t {
	r1, _, _ := proc_xed_get_byte.Call(*(*uintptr)(unsafe.Pointer(&X)), uintptr(I), uintptr(Len))
	return uint8(r1)
}

func (x *Xed) ShortestWidthUnsigned(X Xed_uint64_t, Legal_widths Xed_uint8_t) Xed_uint_t {
	r1, _, _ := proc_xed_shortest_width_unsigned.Call(*(*uintptr)(unsafe.Pointer(&X)), uintptr(Legal_widths))
	return uint32(r1)
}

func (x *Xed) ShortestWidthSigned(X Xed_int64_t, Legal_widths Xed_uint8_t) Xed_uint_t {
	r1, _, _ := proc_xed_shortest_width_signed.Call(*(*uintptr)(unsafe.Pointer(&X)), uintptr(Legal_widths))
	return uint32(r1)
}

func (x *Xed) CategoryEnumT2str(P Xed_category_enum_t) *int8 {
	r1, _, _ := proc_xed_category_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) CategoryEnumTLast() Xed_category_enum_t {
	r1, _, _ := proc_xed_category_enum_t_last.Call()
	return Xed_category_enum_t(uint32(r1))
}

func (x *Xed) ExtensionEnumT2str(P Xed_extension_enum_t) *int8 {
	r1, _, _ := proc_xed_extension_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) ExtensionEnumTLast() Xed_extension_enum_t {
	r1, _, _ := proc_xed_extension_enum_t_last.Call()
	return Xed_extension_enum_t(uint32(r1))
}

func (x *Xed) IclassEnumT2str(P Xed_iclass_enum_t) *int8 {
	r1, _, _ := proc_xed_iclass_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) IclassEnumTLast() Xed_iclass_enum_t {
	r1, _, _ := proc_xed_iclass_enum_t_last.Call()
	return Xed_iclass_enum_t(uint32(r1))
}

func (x *Xed) OperandVisibilityEnumT2str(P Xed_operand_visibility_enum_t) *int8 {
	r1, _, _ := proc_xed_operand_visibility_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandVisibilityEnumTLast() Xed_operand_visibility_enum_t {
	r1, _, _ := proc_xed_operand_visibility_enum_t_last.Call()
	return Xed_operand_visibility_enum_t(uint32(r1))
}

func (x *Xed) OperandActionEnumT2str(P Xed_operand_action_enum_t) *int8 {
	r1, _, _ := proc_xed_operand_action_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandActionEnumTLast() Xed_operand_action_enum_t {
	r1, _, _ := proc_xed_operand_action_enum_t_last.Call()
	return Xed_operand_action_enum_t(uint32(r1))
}

func (x *Xed) OperandConvertEnumT2str(P Xed_operand_convert_enum_t) *int8 {
	r1, _, _ := proc_xed_operand_convert_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandConvertEnumTLast() Xed_operand_convert_enum_t {
	r1, _, _ := proc_xed_operand_convert_enum_t_last.Call()
	return Xed_operand_convert_enum_t(uint32(r1))
}

func (x *Xed) OperandTypeEnumT2str(P Xed_operand_type_enum_t) *int8 {
	r1, _, _ := proc_xed_operand_type_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandTypeEnumTLast() Xed_operand_type_enum_t {
	r1, _, _ := proc_xed_operand_type_enum_t_last.Call()
	return Xed_operand_type_enum_t(uint32(r1))
}

func (x *Xed) NonterminalEnumT2str(P Xed_nonterminal_enum_t) *int8 {
	r1, _, _ := proc_xed_nonterminal_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) NonterminalEnumTLast() Xed_nonterminal_enum_t {
	r1, _, _ := proc_xed_nonterminal_enum_t_last.Call()
	return Xed_nonterminal_enum_t(uint32(r1))
}

func (x *Xed) OperandWidthEnumT2str(P Xed_operand_width_enum_t) *int8 {
	r1, _, _ := proc_xed_operand_width_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandWidthEnumTLast() Xed_operand_width_enum_t {
	r1, _, _ := proc_xed_operand_width_enum_t_last.Call()
	return Xed_operand_width_enum_t(uint32(r1))
}

func (x *Xed) OperandElementXtypeEnumT2str(P Xed_operand_element_xtype_enum_t) *int8 {
	r1, _, _ := proc_xed_operand_element_xtype_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandElementXtypeEnumTLast() Xed_operand_element_xtype_enum_t {
	r1, _, _ := proc_xed_operand_element_xtype_enum_t_last.Call()
	return Xed_operand_element_xtype_enum_t(uint32(r1))
}

func (x *Xed) RegEnumT2str(P Xed_reg_enum_t) *int8 {
	r1, _, _ := proc_xed_reg_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) RegEnumTLast() Xed_reg_enum_t {
	r1, _, _ := proc_xed_reg_enum_t_last.Call()
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) AttributeEnumT2str(P Xed_attribute_enum_t) *int8 {
	r1, _, _ := proc_xed_attribute_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) AttributeEnumTLast() Xed_attribute_enum_t {
	r1, _, _ := proc_xed_attribute_enum_t_last.Call()
	return Xed_attribute_enum_t(uint32(r1))
}

func (x *Xed) ExceptionEnumT2str(P Xed_exception_enum_t) *int8 {
	r1, _, _ := proc_xed_exception_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) ExceptionEnumTLast() Xed_exception_enum_t {
	r1, _, _ := proc_xed_exception_enum_t_last.Call()
	return Xed_exception_enum_t(uint32(r1))
}

func (x *Xed) IformEnumT2str(P Xed_iform_enum_t) *int8 {
	r1, _, _ := proc_xed_iform_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) IformEnumTLast() Xed_iform_enum_t {
	r1, _, _ := proc_xed_iform_enum_t_last.Call()
	return Xed_iform_enum_t(uint32(r1))
}

func (x *Xed) IsaSetEnumT2str(P Xed_isa_set_enum_t) *int8 {
	r1, _, _ := proc_xed_isa_set_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) IsaSetEnumTLast() Xed_isa_set_enum_t {
	r1, _, _ := proc_xed_isa_set_enum_t_last.Call()
	return Xed_isa_set_enum_t(uint32(r1))
}

func (x *Xed) IformMap(Iform Xed_iform_enum_t) *Xed_iform_info_t {
	r1, _, _ := proc_xed_iform_map.Call(uintptr(Iform))
	return (*Xed_iform_info_t)(unsafe.Pointer(r1))
}

func (x *Xed) IformMaxPerIclass(Iclass Xed_iclass_enum_t) Xed_uint32_t {
	r1, _, _ := proc_xed_iform_max_per_iclass.Call(uintptr(Iclass))
	return uint32(r1)
}

func (x *Xed) IformFirstPerIclass(Iclass Xed_iclass_enum_t) Xed_uint32_t {
	r1, _, _ := proc_xed_iform_first_per_iclass.Call(uintptr(Iclass))
	return uint32(r1)
}

func (x *Xed) IformToCategory(Iform Xed_iform_enum_t) Xed_category_enum_t {
	r1, _, _ := proc_xed_iform_to_category.Call(uintptr(Iform))
	return Xed_category_enum_t(uint32(r1))
}

func (x *Xed) IformToExtension(Iform Xed_iform_enum_t) Xed_extension_enum_t {
	r1, _, _ := proc_xed_iform_to_extension.Call(uintptr(Iform))
	return Xed_extension_enum_t(uint32(r1))
}

func (x *Xed) IformToIsaSet(Iform Xed_iform_enum_t) Xed_isa_set_enum_t {
	r1, _, _ := proc_xed_iform_to_isa_set.Call(uintptr(Iform))
	return Xed_isa_set_enum_t(uint32(r1))
}

func (x *Xed) IformToIclassStringAtt(Iform Xed_iform_enum_t) *int8 {
	r1, _, _ := proc_xed_iform_to_iclass_string_att.Call(uintptr(Iform))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) IformToIclassStringIntel(Iform Xed_iform_enum_t) *int8 {
	r1, _, _ := proc_xed_iform_to_iclass_string_intel.Call(uintptr(Iform))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandWidthBits(P *Xed_operand_t, Eosz Xed_uint32_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_width_bits.Call(uintptr(unsafe.Pointer(P)), uintptr(Eosz))
	return uint32(r1)
}

func (x *Xed) OperandPrint(Xedd *Xed_decoded_inst_s, I uint32, Buf *int8, Buflen int32) {
	proc_xed_operand_print.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(I), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
}

func (x *Xed) OperandRead(P *Xed_operand_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_read.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandReadOnly(P *Xed_operand_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_read_only.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandWritten(P *Xed_operand_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_written.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandWrittenOnly(P *Xed_operand_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_written_only.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandReadAndWritten(P *Xed_operand_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_read_and_written.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandConditionalRead(P *Xed_operand_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_conditional_read.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandConditionalWrite(P *Xed_operand_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_conditional_write.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) InstCpl(P *Xed_inst_t) uint32 {
	r1, _, _ := proc_xed_inst_cpl.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) InstOperand(P *Xed_inst_t, I uint32) *Xed_operand_t {
	r1, _, _ := proc_xed_inst_operand.Call(uintptr(unsafe.Pointer(P)), uintptr(I))
	return (*Xed_operand_t)(unsafe.Pointer(r1))
}

func (x *Xed) InstFlagInfoIndex(P *Xed_inst_t) Xed_uint32_t {
	r1, _, _ := proc_xed_inst_flag_info_index.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) InstGetAttribute(P *Xed_inst_t, Attr Xed_attribute_enum_t) Xed_uint32_t {
	r1, _, _ := proc_xed_inst_get_attribute.Call(uintptr(unsafe.Pointer(P)), uintptr(Attr))
	return uint32(r1)
}

func (x *Xed) InstGetAttributes(P *Xed_inst_t) Xed_attributes_t {
	r1, _, _ := proc_xed_inst_get_attributes.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_attributes_t)(unsafe.Pointer(&r1))
}

func (x *Xed) AttributeMax() uint32 {
	r1, _, _ := proc_xed_attribute_max.Call()
	return uint32(r1)
}

func (x *Xed) Attribute(I uint32) Xed_attribute_enum_t {
	r1, _, _ := proc_xed_attribute.Call(uintptr(I))
	return Xed_attribute_enum_t(uint32(r1))
}

func (x *Xed) InstTableBase() *Xed_inst_t {
	r1, _, _ := proc_xed_inst_table_base.Call()
	return (*Xed_inst_t)(unsafe.Pointer(r1))
}

func (x *Xed) FlagEnumT2str(P Xed_flag_enum_t) *int8 {
	r1, _, _ := proc_xed_flag_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) FlagEnumTLast() Xed_flag_enum_t {
	r1, _, _ := proc_xed_flag_enum_t_last.Call()
	return Xed_flag_enum_t(uint32(r1))
}

func (x *Xed) FlagActionEnumT2str(P Xed_flag_action_enum_t) *int8 {
	r1, _, _ := proc_xed_flag_action_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) FlagActionEnumTLast() Xed_flag_action_enum_t {
	r1, _, _ := proc_xed_flag_action_enum_t_last.Call()
	return Xed_flag_action_enum_t(uint32(r1))
}

func (x *Xed) FlagSetPrint(P *Xed_flag_set_t, Buf *int8, Buflen int32) int32 {
	r1, _, _ := proc_xed_flag_set_print.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
	return int32(r1)
}

func (x *Xed) FlagSetIsSubsetOf(P *Xed_flag_set_t, Other *Xed_flag_set_t) Xed_bool_t {
	r1, _, _ := proc_xed_flag_set_is_subset_of.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Other)))
	return uint32(r1)
}

func (x *Xed) FlagActionGetFlagName(P *Xed_flag_action_t) Xed_flag_enum_t {
	r1, _, _ := proc_xed_flag_action_get_flag_name.Call(uintptr(unsafe.Pointer(P)))
	return Xed_flag_enum_t(uint32(r1))
}

func (x *Xed) FlagActionGetAction(P *Xed_flag_action_t, I uint32) Xed_flag_action_enum_t {
	r1, _, _ := proc_xed_flag_action_get_action.Call(uintptr(unsafe.Pointer(P)), uintptr(I))
	return Xed_flag_action_enum_t(uint32(r1))
}

func (x *Xed) FlagActionActionInvalid(A Xed_flag_action_enum_t) Xed_bool_t {
	r1, _, _ := proc_xed_flag_action_action_invalid.Call(uintptr(A))
	return uint32(r1)
}

func (x *Xed) FlagActionPrint(P *Xed_flag_action_t, Buf *int8, Buflen int32) int32 {
	r1, _, _ := proc_xed_flag_action_print.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
	return int32(r1)
}

func (x *Xed) FlagActionReadFlag(P *Xed_flag_action_t) Xed_bool_t {
	r1, _, _ := proc_xed_flag_action_read_flag.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) FlagActionWritesFlag(P *Xed_flag_action_t) Xed_bool_t {
	r1, _, _ := proc_xed_flag_action_writes_flag.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) FlagActionReadAction(A Xed_flag_action_enum_t) Xed_bool_t {
	r1, _, _ := proc_xed_flag_action_read_action.Call(uintptr(A))
	return uint32(r1)
}

func (x *Xed) FlagActionWriteAction(A Xed_flag_action_enum_t) Xed_bool_t {
	r1, _, _ := proc_xed_flag_action_write_action.Call(uintptr(A))
	return uint32(r1)
}

func (x *Xed) SimpleFlagGetNflags(P *Xed_simple_flag_t) uint32 {
	r1, _, _ := proc_xed_simple_flag_get_nflags.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) SimpleFlagGetReadFlagSet(P *Xed_simple_flag_t) *Xed_flag_set_t {
	r1, _, _ := proc_xed_simple_flag_get_read_flag_set.Call(uintptr(unsafe.Pointer(P)))
	return (*Xed_flag_set_t)(unsafe.Pointer(r1))
}

func (x *Xed) SimpleFlagGetWrittenFlagSet(P *Xed_simple_flag_t) *Xed_flag_set_t {
	r1, _, _ := proc_xed_simple_flag_get_written_flag_set.Call(uintptr(unsafe.Pointer(P)))
	return (*Xed_flag_set_t)(unsafe.Pointer(r1))
}

func (x *Xed) SimpleFlagGetUndefinedFlagSet(P *Xed_simple_flag_t) *Xed_flag_set_t {
	r1, _, _ := proc_xed_simple_flag_get_undefined_flag_set.Call(uintptr(unsafe.Pointer(P)))
	return (*Xed_flag_set_t)(unsafe.Pointer(r1))
}

func (x *Xed) SimpleFlagGetMayWrite(P *Xed_simple_flag_t) Xed_bool_t {
	r1, _, _ := proc_xed_simple_flag_get_may_write.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) SimpleFlagGetMustWrite(P *Xed_simple_flag_t) Xed_bool_t {
	r1, _, _ := proc_xed_simple_flag_get_must_write.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) SimpleFlagGetFlagAction(P *Xed_simple_flag_t, I uint32) *Xed_flag_action_t {
	r1, _, _ := proc_xed_simple_flag_get_flag_action.Call(uintptr(unsafe.Pointer(P)), uintptr(I))
	return (*Xed_flag_action_t)(unsafe.Pointer(r1))
}

func (x *Xed) SimpleFlagReadsFlags(P *Xed_simple_flag_t) Xed_bool_t {
	r1, _, _ := proc_xed_simple_flag_reads_flags.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) SimpleFlagWritesFlags(P *Xed_simple_flag_t) Xed_bool_t {
	r1, _, _ := proc_xed_simple_flag_writes_flags.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) SimpleFlagPrint(P *Xed_simple_flag_t, Buf *int8, Buflen int32) int32 {
	r1, _, _ := proc_xed_simple_flag_print.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
	return int32(r1)
}

func (x *Xed) ChipEnumT2str(P Xed_chip_enum_t) *int8 {
	r1, _, _ := proc_xed_chip_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) ChipEnumTLast() Xed_chip_enum_t {
	r1, _, _ := proc_xed_chip_enum_t_last.Call()
	return Xed_chip_enum_t(uint32(r1))
}

func (x *Xed) OperandElementTypeEnumT2str(P Xed_operand_element_type_enum_t) *int8 {
	r1, _, _ := proc_xed_operand_element_type_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) OperandElementTypeEnumTLast() Xed_operand_element_type_enum_t {
	r1, _, _ := proc_xed_operand_element_type_enum_t_last.Call()
	return Xed_operand_element_type_enum_t(uint32(r1))
}

func (x *Xed) ErrorEnumT2str(P Xed_error_enum_t) *int8 {
	r1, _, _ := proc_xed_error_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) ErrorEnumTLast() Xed_error_enum_t {
	r1, _, _ := proc_xed_error_enum_t_last.Call()
	return Xed_error_enum_t(uint32(r1))
}

func (x *Xed) GetChipFeatures(P *Xed_chip_features_t, Chip Xed_chip_enum_t) {
	proc_xed_get_chip_features.Call(uintptr(unsafe.Pointer(P)), uintptr(Chip))
}

func (x *Xed) ModifyChipFeatures(P *Xed_chip_features_t, Isa_set Xed_isa_set_enum_t, Present Xed_bool_t) {
	proc_xed_modify_chip_features.Call(uintptr(unsafe.Pointer(P)), uintptr(Isa_set), uintptr(Present))
}

func (x *Xed) SetDecoderModes(Xedd *Xed_decoded_inst_t, Chip Xed_chip_enum_t, Chip_features *Xed_chip_features_t) {
	proc_xed_set_decoder_modes.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(Chip), uintptr(unsafe.Pointer(Chip_features)))
}

func (x *Xed) Decode(Xedd *Xed_decoded_inst_t, Itext *Xed_uint8_t, Bytes uint32) Xed_error_enum_t {
	r1, _, _ := proc_xed_decode.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(unsafe.Pointer(Itext)), uintptr(Bytes))
	return Xed_error_enum_t(uint32(r1))
}

func (x *Xed) DecodeWithFeatures(Xedd *Xed_decoded_inst_t, Itext *Xed_uint8_t, Bytes uint32, Features *Xed_chip_features_t) Xed_error_enum_t {
	r1, _, _ := proc_xed_decode_with_features.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(unsafe.Pointer(Itext)), uintptr(Bytes), uintptr(unsafe.Pointer(Features)))
	return Xed_error_enum_t(uint32(r1))
}

func (x *Xed) IldDecode(Xedd *Xed_decoded_inst_t, Itext *Xed_uint8_t, Bytes uint32) Xed_error_enum_t {
	r1, _, _ := proc_xed_ild_decode.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(unsafe.Pointer(Itext)), uintptr(Bytes))
	return Xed_error_enum_t(uint32(r1))
}

func (x *Xed) AddressWidthEnumT2str(P Xed_address_width_enum_t) *int8 {
	r1, _, _ := proc_xed_address_width_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) AddressWidthEnumTLast() Xed_address_width_enum_t {
	r1, _, _ := proc_xed_address_width_enum_t_last.Call()
	return Xed_address_width_enum_t(uint32(r1))
}

func (x *Xed) MachineModeEnumT2str(P Xed_machine_mode_enum_t) *int8 {
	r1, _, _ := proc_xed_machine_mode_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) MachineModeEnumTLast() Xed_machine_mode_enum_t {
	r1, _, _ := proc_xed_machine_mode_enum_t_last.Call()
	return Xed_machine_mode_enum_t(uint32(r1))
}

func (x *Xed) StatePrint(P *Xed_state_t, Buf *int8, Buflen int32) int32 {
	r1, _, _ := proc_xed_state_print.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
	return int32(r1)
}

func (x *Xed) SyntaxEnumT2str(P Xed_syntax_enum_t) *int8 {
	r1, _, _ := proc_xed_syntax_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) SyntaxEnumTLast() Xed_syntax_enum_t {
	r1, _, _ := proc_xed_syntax_enum_t_last.Call()
	return Xed_syntax_enum_t(uint32(r1))
}

func (x *Xed) RegClassEnumT2str(P Xed_reg_class_enum_t) *int8 {
	r1, _, _ := proc_xed_reg_class_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) RegClassEnumTLast() Xed_reg_class_enum_t {
	r1, _, _ := proc_xed_reg_class_enum_t_last.Call()
	return Xed_reg_class_enum_t(uint32(r1))
}

func (x *Xed) RegClass(R Xed_reg_enum_t) Xed_reg_class_enum_t {
	r1, _, _ := proc_xed_reg_class.Call(uintptr(R))
	return Xed_reg_class_enum_t(uint32(r1))
}

func (x *Xed) GprRegClass(R Xed_reg_enum_t) Xed_reg_class_enum_t {
	r1, _, _ := proc_xed_gpr_reg_class.Call(uintptr(R))
	return Xed_reg_class_enum_t(uint32(r1))
}

func (x *Xed) GetLargestEnclosingRegister(R Xed_reg_enum_t) Xed_reg_enum_t {
	r1, _, _ := proc_xed_get_largest_enclosing_register.Call(uintptr(R))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) GetLargestEnclosingRegister32(R Xed_reg_enum_t) Xed_reg_enum_t {
	r1, _, _ := proc_xed_get_largest_enclosing_register32.Call(uintptr(R))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) GetRegisterWidthBits(R Xed_reg_enum_t) Xed_uint32_t {
	r1, _, _ := proc_xed_get_register_width_bits.Call(uintptr(R))
	return uint32(r1)
}

func (x *Xed) GetRegisterWidthBits64(R Xed_reg_enum_t) Xed_uint32_t {
	r1, _, _ := proc_xed_get_register_width_bits64.Call(uintptr(R))
	return uint32(r1)
}

func (x *Xed) OperandValuesInit(P *Xed_operand_values_t) {
	proc_xed_operand_values_init.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesInitSetMode(P *Xed_operand_values_t, Dstate *Xed_state_t) {
	proc_xed_operand_values_init_set_mode.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Dstate)))
}

func (x *Xed) OperandValuesSetMode(P *Xed_operand_values_t, Dstate *Xed_state_t) {
	proc_xed_operand_values_set_mode.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Dstate)))
}

func (x *Xed) OperandValuesInitKeepMode(Dst *Xed_operand_values_t, Src *Xed_operand_values_t) {
	proc_xed_operand_values_init_keep_mode.Call(uintptr(unsafe.Pointer(Dst)), uintptr(unsafe.Pointer(Src)))
}

func (x *Xed) OperandValuesDump(Ov *Xed_operand_values_t, Buf *int8, Buflen int32) {
	proc_xed_operand_values_dump.Call(uintptr(unsafe.Pointer(Ov)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
}

func (x *Xed) OperandValuesPrintShort(Ov *Xed_operand_values_t, Buf *int8, Buflen int32) {
	proc_xed_operand_values_print_short.Call(uintptr(unsafe.Pointer(Ov)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
}

func (x *Xed) OperandValuesHasRealRep(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_real_rep.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasRepPrefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_rep_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasRepnePrefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_repne_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesClearRep(P *Xed_operand_values_t) {
	proc_xed_operand_values_clear_rep.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesGetAtomic(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_get_atomic.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasLockPrefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_lock_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesLockable(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_lockable.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesUsingDefaultSegment(P *Xed_operand_values_t, I uint32) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_using_default_segment.Call(uintptr(unsafe.Pointer(P)), uintptr(I))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetEffectiveOperandWidth(P *Xed_operand_values_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_values_get_effective_operand_width.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetEffectiveAddressWidth(P *Xed_operand_values_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_values_get_effective_address_width.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetStackAddressWidth(P *Xed_operand_values_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_values_get_stack_address_width.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasMemoryDisplacement(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_memory_displacement.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasBranchDisplacement(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_branch_displacement.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasDisplacement(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_displacement.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetDisplacementForMemop(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_get_displacement_for_memop.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasImmediate(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_immediate.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasAddressSizePrefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_address_size_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasOperandSizePrefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_operand_size_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHas66Prefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_66_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesMandatory66Prefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_mandatory_66_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasRexwPrefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_rexw_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasSegmentPrefix(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_segment_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesSegmentPrefix(P *Xed_operand_values_t) Xed_reg_enum_t {
	r1, _, _ := proc_xed_operand_values_segment_prefix.Call(uintptr(unsafe.Pointer(P)))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) OperandValuesGetLongMode(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_get_long_mode.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetRealMode(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_get_real_mode.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetPpVexPrefix(P *Xed_operand_values_t) Xed_bits_t {
	r1, _, _ := proc_xed_operand_values_get_pp_vex_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesAccessesMemory(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_accesses_memory.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesNumberOfMemoryOperands(P *Xed_operand_values_t) uint32 {
	r1, _, _ := proc_xed_operand_values_number_of_memory_operands.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetMemoryOperandLength(P *Xed_operand_values_t, Memop_idx uint32) uint32 {
	r1, _, _ := proc_xed_operand_values_get_memory_operand_length.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetBaseReg(P *Xed_operand_values_t, Memop_idx uint32) Xed_reg_enum_t {
	r1, _, _ := proc_xed_operand_values_get_base_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) OperandValuesGetIndexReg(P *Xed_operand_values_t, Memop_idx uint32) Xed_reg_enum_t {
	r1, _, _ := proc_xed_operand_values_get_index_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) OperandValuesGetSegReg(P *Xed_operand_values_t, Memop_idx uint32) Xed_reg_enum_t {
	r1, _, _ := proc_xed_operand_values_get_seg_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) OperandValuesGetScale(P *Xed_operand_values_t) uint32 {
	r1, _, _ := proc_xed_operand_values_get_scale.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesMemopWithoutModrm(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_memop_without_modrm.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasModrmByte(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_modrm_byte.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesHasSibByte(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_has_sib_byte.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesBranchNotTakenHint(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_branch_not_taken_hint.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesBranchTakenHint(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_branch_taken_hint.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesIgnoredBranchNotTakenHint(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_ignored_branch_not_taken_hint.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesIgnoredBranchTakenHint(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_ignored_branch_taken_hint.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesCetNoTrack(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_cet_no_track.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesIsNop(P *Xed_operand_values_t) Xed_bool_t {
	r1, _, _ := proc_xed_operand_values_is_nop.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetImmediateInt64(P *Xed_operand_values_t) Xed_int64_t {
	r1, _, _ := proc_xed_operand_values_get_immediate_int64.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) OperandValuesGetImmediateUint64(P *Xed_operand_values_t) Xed_uint64_t {
	r1, _, _ := proc_xed_operand_values_get_immediate_uint64.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_uint64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) OperandValuesGetImmediateIsSigned(P *Xed_operand_values_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_values_get_immediate_is_signed.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetImmediateByte(P *Xed_operand_values_t, I uint32) Xed_uint8_t {
	r1, _, _ := proc_xed_operand_values_get_immediate_byte.Call(uintptr(unsafe.Pointer(P)), uintptr(I))
	return uint8(r1)
}

func (x *Xed) OperandValuesGetSecondImmediate(P *Xed_operand_values_t) Xed_uint8_t {
	r1, _, _ := proc_xed_operand_values_get_second_immediate.Call(uintptr(unsafe.Pointer(P)))
	return uint8(r1)
}

func (x *Xed) OperandValuesGetMemoryDisplacementLength(P *Xed_operand_values_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_values_get_memory_displacement_length.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetMemoryDisplacementLengthBits(P *Xed_operand_values_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_values_get_memory_displacement_length_bits.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetMemoryDisplacementLengthBitsRaw(P *Xed_operand_values_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_values_get_memory_displacement_length_bits_raw.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetMemoryDisplacementInt64(P *Xed_operand_values_t) Xed_int64_t {
	r1, _, _ := proc_xed_operand_values_get_memory_displacement_int64.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) OperandValuesGetMemoryDisplacementInt64Raw(P *Xed_operand_values_t) Xed_int64_t {
	r1, _, _ := proc_xed_operand_values_get_memory_displacement_int64_raw.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) OperandValuesGetMemoryDisplacementByte(P *Xed_operand_values_t, I uint32) Xed_uint8_t {
	r1, _, _ := proc_xed_operand_values_get_memory_displacement_byte.Call(uintptr(unsafe.Pointer(P)), uintptr(I))
	return uint8(r1)
}

func (x *Xed) OperandValuesGetBranchDisplacementLength(P *Xed_operand_values_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_values_get_branch_displacement_length.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetBranchDisplacementLengthBits(P *Xed_operand_values_t) Xed_uint32_t {
	r1, _, _ := proc_xed_operand_values_get_branch_displacement_length_bits.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) OperandValuesGetBranchDisplacementInt64(P *Xed_operand_values_t) Xed_int64_t {
	r1, _, _ := proc_xed_operand_values_get_branch_displacement_int64.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) OperandValuesGetBranchDisplacementByte(P *Xed_operand_values_t, I uint32) Xed_uint8_t {
	r1, _, _ := proc_xed_operand_values_get_branch_displacement_byte.Call(uintptr(unsafe.Pointer(P)), uintptr(I))
	return uint8(r1)
}

func (x *Xed) OperandValuesGetIclass(P *Xed_operand_values_t) Xed_iclass_enum_t {
	r1, _, _ := proc_xed_operand_values_get_iclass.Call(uintptr(unsafe.Pointer(P)))
	return Xed_iclass_enum_t(uint32(r1))
}

func (x *Xed) OperandValuesZeroImmediate(P *Xed_operand_values_t) {
	proc_xed_operand_values_zero_immediate.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesZeroBranchDisplacement(P *Xed_operand_values_t) {
	proc_xed_operand_values_zero_branch_displacement.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesZeroMemoryDisplacement(P *Xed_operand_values_t) {
	proc_xed_operand_values_zero_memory_displacement.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesSetLock(P *Xed_operand_values_t) {
	proc_xed_operand_values_set_lock.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesZeroSegmentOverride(P *Xed_operand_values_t) {
	proc_xed_operand_values_zero_segment_override.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesSetIclass(P *Xed_operand_values_t, Iclass Xed_iclass_enum_t) {
	proc_xed_operand_values_set_iclass.Call(uintptr(unsafe.Pointer(P)), uintptr(Iclass))
}

func (x *Xed) OperandValuesSetEffectiveOperandWidth(P *Xed_operand_values_t, Width uint32) {
	proc_xed_operand_values_set_effective_operand_width.Call(uintptr(unsafe.Pointer(P)), uintptr(Width))
}

func (x *Xed) OperandValuesSetEffectiveAddressWidth(P *Xed_operand_values_t, Width uint32) {
	proc_xed_operand_values_set_effective_address_width.Call(uintptr(unsafe.Pointer(P)), uintptr(Width))
}

func (x *Xed) OperandValuesSetMemoryOperandLength(P *Xed_operand_values_t, Memop_length uint32) {
	proc_xed_operand_values_set_memory_operand_length.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_length))
}

func (x *Xed) OperandValuesSetMemoryDisplacement(P *Xed_operand_values_t, X Xed_int64_t, Len uint32) {
	proc_xed_operand_values_set_memory_displacement.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&X)), uintptr(Len))
}

func (x *Xed) OperandValuesSetMemoryDisplacementBits(P *Xed_operand_values_t, X Xed_int64_t, Len_bits uint32) {
	proc_xed_operand_values_set_memory_displacement_bits.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&X)), uintptr(Len_bits))
}

func (x *Xed) OperandValuesSetRelbr(P *Xed_operand_values_t) {
	proc_xed_operand_values_set_relbr.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesSetAbsbr(P *Xed_operand_values_t) {
	proc_xed_operand_values_set_absbr.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) OperandValuesSetBranchDisplacement(P *Xed_operand_values_t, X Xed_int64_t, Len uint32) {
	proc_xed_operand_values_set_branch_displacement.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&X)), uintptr(Len))
}

func (x *Xed) OperandValuesSetBranchDisplacementBits(P *Xed_operand_values_t, X Xed_int64_t, Len_bits uint32) {
	proc_xed_operand_values_set_branch_displacement_bits.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&X)), uintptr(Len_bits))
}

func (x *Xed) OperandValuesSetImmediateSigned(P *Xed_operand_values_t, X Xed_int32_t, Bytes uint32) {
	proc_xed_operand_values_set_immediate_signed.Call(uintptr(unsafe.Pointer(P)), uintptr(X), uintptr(Bytes))
}

func (x *Xed) OperandValuesSetImmediateSignedBits(P *Xed_operand_values_t, X Xed_int32_t, Bits uint32) {
	proc_xed_operand_values_set_immediate_signed_bits.Call(uintptr(unsafe.Pointer(P)), uintptr(X), uintptr(Bits))
}

func (x *Xed) OperandValuesSetImmediateUnsigned(P *Xed_operand_values_t, X Xed_uint64_t, Bytes uint32) {
	proc_xed_operand_values_set_immediate_unsigned.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&X)), uintptr(Bytes))
}

func (x *Xed) OperandValuesSetImmediateUnsignedBits(P *Xed_operand_values_t, X Xed_uint64_t, Bits uint32) {
	proc_xed_operand_values_set_immediate_unsigned_bits.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&X)), uintptr(Bits))
}

func (x *Xed) OperandValuesSetBaseReg(P *Xed_operand_values_t, Memop_idx uint32, New_base Xed_reg_enum_t) {
	proc_xed_operand_values_set_base_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx), uintptr(New_base))
}

func (x *Xed) OperandValuesSetSegReg(P *Xed_operand_values_t, Memop_idx uint32, New_seg Xed_reg_enum_t) {
	proc_xed_operand_values_set_seg_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx), uintptr(New_seg))
}

func (x *Xed) OperandValuesSetIndexReg(P *Xed_operand_values_t, Memop_idx uint32, New_index Xed_reg_enum_t) {
	proc_xed_operand_values_set_index_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx), uintptr(New_index))
}

func (x *Xed) OperandValuesSetScale(P *Xed_operand_values_t, Memop_idx Xed_uint_t, New_scale Xed_uint_t) {
	proc_xed_operand_values_set_scale.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx), uintptr(New_scale))
}

func (x *Xed) OperandValuesSetOperandReg(P *Xed_operand_values_t, Operand_name Xed_operand_enum_t, Reg_name Xed_reg_enum_t) {
	proc_xed_operand_values_set_operand_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_name), uintptr(Reg_name))
}

func (x *Xed) EncoderRequestGetIclass(P *Xed_encoder_request_t) Xed_iclass_enum_t {
	r1, _, _ := proc_xed_encoder_request_get_iclass.Call(uintptr(unsafe.Pointer(P)))
	return Xed_iclass_enum_t(uint32(r1))
}

func (x *Xed) EncoderRequestSetIclass(P *Xed_encoder_request_t, Iclass Xed_iclass_enum_t) {
	proc_xed_encoder_request_set_iclass.Call(uintptr(unsafe.Pointer(P)), uintptr(Iclass))
}

func (x *Xed) EncoderRequestSetEffectiveOperandWidth(P *Xed_encoder_request_t, Width_bits Xed_uint_t) {
	proc_xed_encoder_request_set_effective_operand_width.Call(uintptr(unsafe.Pointer(P)), uintptr(Width_bits))
}

func (x *Xed) EncoderRequestSetEffectiveAddressSize(P *Xed_encoder_request_t, Width_bits Xed_uint_t) {
	proc_xed_encoder_request_set_effective_address_size.Call(uintptr(unsafe.Pointer(P)), uintptr(Width_bits))
}

func (x *Xed) EncoderRequestSetReg(P *Xed_encoder_request_t, Operand Xed_operand_enum_t, Reg Xed_reg_enum_t) {
	proc_xed_encoder_request_set_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand), uintptr(Reg))
}

func (x *Xed) EncoderRequestSetOperandOrder(P *Xed_encoder_request_t, Operand_index Xed_uint_t, Name Xed_operand_enum_t) {
	proc_xed_encoder_request_set_operand_order.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_index), uintptr(Name))
}

func (x *Xed) EncoderRequestGetOperandOrder(P *Xed_encoder_request_t, Operand_index Xed_uint_t) Xed_operand_enum_t {
	r1, _, _ := proc_xed_encoder_request_get_operand_order.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_index))
	return Xed_operand_enum_t(uint32(r1))
}

func (x *Xed) EncoderRequestSetRelbr(P *Xed_encoder_request_t) {
	proc_xed_encoder_request_set_relbr.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) EncoderRequestSetAbsbr(P *Xed_encoder_request_t) {
	proc_xed_encoder_request_set_absbr.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) EncoderRequestSetBranchDisplacement(P *Xed_encoder_request_t, Brdisp Xed_int64_t, Nbytes Xed_uint_t) {
	proc_xed_encoder_request_set_branch_displacement.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&Brdisp)), uintptr(Nbytes))
}

func (x *Xed) EncoderRequestSetPtr(P *Xed_encoder_request_t) {
	proc_xed_encoder_request_set_ptr.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) EncoderRequestSetUimm0(P *Xed_encoder_request_t, Uimm Xed_uint64_t, Nbytes Xed_uint_t) {
	proc_xed_encoder_request_set_uimm0.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&Uimm)), uintptr(Nbytes))
}

func (x *Xed) EncoderRequestSetUimm0Bits(P *Xed_encoder_request_t, Uimm Xed_uint64_t, Nbits Xed_uint_t) {
	proc_xed_encoder_request_set_uimm0_bits.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&Uimm)), uintptr(Nbits))
}

func (x *Xed) EncoderRequestSetUimm1(P *Xed_encoder_request_t, Uimm Xed_uint8_t) {
	proc_xed_encoder_request_set_uimm1.Call(uintptr(unsafe.Pointer(P)), uintptr(Uimm))
}

func (x *Xed) EncoderRequestSetSimm(P *Xed_encoder_request_t, Simm Xed_int32_t, Nbytes Xed_uint_t) {
	proc_xed_encoder_request_set_simm.Call(uintptr(unsafe.Pointer(P)), uintptr(Simm), uintptr(Nbytes))
}

func (x *Xed) EncoderRequestSetMemoryDisplacement(P *Xed_encoder_request_t, Memdisp Xed_int64_t, Nbytes Xed_uint_t) {
	proc_xed_encoder_request_set_memory_displacement.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&Memdisp)), uintptr(Nbytes))
}

func (x *Xed) EncoderRequestSetAgen(P *Xed_encoder_request_t) {
	proc_xed_encoder_request_set_agen.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) EncoderRequestSetMem0(P *Xed_encoder_request_t) {
	proc_xed_encoder_request_set_mem0.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) EncoderRequestSetMem1(P *Xed_encoder_request_t) {
	proc_xed_encoder_request_set_mem1.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) EncoderRequestSetMemoryOperandLength(P *Xed_encoder_request_t, Nbytes Xed_uint_t) {
	proc_xed_encoder_request_set_memory_operand_length.Call(uintptr(unsafe.Pointer(P)), uintptr(Nbytes))
}

func (x *Xed) EncoderRequestSetSeg0(P *Xed_encoder_request_t, Seg_reg Xed_reg_enum_t) {
	proc_xed_encoder_request_set_seg0.Call(uintptr(unsafe.Pointer(P)), uintptr(Seg_reg))
}

func (x *Xed) EncoderRequestSetSeg1(P *Xed_encoder_request_t, Seg_reg Xed_reg_enum_t) {
	proc_xed_encoder_request_set_seg1.Call(uintptr(unsafe.Pointer(P)), uintptr(Seg_reg))
}

func (x *Xed) EncoderRequestSetBase0(P *Xed_encoder_request_t, Base_reg Xed_reg_enum_t) {
	proc_xed_encoder_request_set_base0.Call(uintptr(unsafe.Pointer(P)), uintptr(Base_reg))
}

func (x *Xed) EncoderRequestSetBase1(P *Xed_encoder_request_t, Base_reg Xed_reg_enum_t) {
	proc_xed_encoder_request_set_base1.Call(uintptr(unsafe.Pointer(P)), uintptr(Base_reg))
}

func (x *Xed) EncoderRequestSetIndex(P *Xed_encoder_request_t, Index_reg Xed_reg_enum_t) {
	proc_xed_encoder_request_set_index.Call(uintptr(unsafe.Pointer(P)), uintptr(Index_reg))
}

func (x *Xed) EncoderRequestSetScale(P *Xed_encoder_request_t, Scale Xed_uint_t) {
	proc_xed_encoder_request_set_scale.Call(uintptr(unsafe.Pointer(P)), uintptr(Scale))
}

func (x *Xed) EncoderRequestZeroOperandOrder(P *Xed_encoder_request_t) {
	proc_xed_encoder_request_zero_operand_order.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) EncoderRequestZeroSetMode(P *Xed_encoder_request_t, Dstate *Xed_state_t) {
	proc_xed_encoder_request_zero_set_mode.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Dstate)))
}

func (x *Xed) EncoderRequestZero(P *Xed_encoder_request_t) {
	proc_xed_encoder_request_zero.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) EncoderRequestInitFromDecode(D *Xed_decoded_inst_s) {
	proc_xed_encoder_request_init_from_decode.Call(uintptr(unsafe.Pointer(D)))
}

func (x *Xed) EncodeRequestPrint(P *Xed_encoder_request_t, Buf *int8, Buflen Xed_uint_t) {
	proc_xed_encode_request_print.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
}

func (x *Xed) Encode(R *Xed_encoder_request_t, Array *Xed_uint8_t, Ilen uint32, Olen *uint32) Xed_error_enum_t {
	r1, _, _ := proc_xed_encode.Call(uintptr(unsafe.Pointer(R)), uintptr(unsafe.Pointer(Array)), uintptr(Ilen), uintptr(unsafe.Pointer(Olen)))
	return Xed_error_enum_t(uint32(r1))
}

func (x *Xed) EncodeNop(Array *Xed_uint8_t, Ilen uint32) Xed_error_enum_t {
	r1, _, _ := proc_xed_encode_nop.Call(uintptr(unsafe.Pointer(Array)), uintptr(Ilen))
	return Xed_error_enum_t(uint32(r1))
}

func (x *Xed) ConvertToEncoderRequest(Out *Xed_encoder_request_t, In *Xed_encoder_instruction_t) Xed_bool_t {
	r1, _, _ := proc_xed_convert_to_encoder_request.Call(uintptr(unsafe.Pointer(Out)), uintptr(unsafe.Pointer(In)))
	return uint32(r1)
}

func (x *Xed) PatchDisp(Xedd *Xed_decoded_inst_t, Itext *Xed_uint8_t, Disp Xed_enc_displacement_t) Xed_bool_t {
	r1, _, _ := proc_xed_patch_disp.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(unsafe.Pointer(Itext)), uintptr(unsafe.Pointer(&Disp)))
	return uint32(r1)
}

func (x *Xed) PatchBrdisp(Xedd *Xed_decoded_inst_t, Itext *Xed_uint8_t, Disp Xed_encoder_operand_t) Xed_bool_t {
	r1, _, _ := proc_xed_patch_brdisp.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(unsafe.Pointer(Itext)), uintptr(unsafe.Pointer(&Disp)))
	return uint32(r1)
}

func (x *Xed) PatchImm0(Xedd *Xed_decoded_inst_t, Itext *Xed_uint8_t, Imm0 Xed_encoder_operand_t) Xed_bool_t {
	r1, _, _ := proc_xed_patch_imm0.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(unsafe.Pointer(Itext)), uintptr(unsafe.Pointer(&Imm0)))
	return uint32(r1)
}

func (x *Xed) OperandActionRead(Rw Xed_operand_action_enum_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_action_read.Call(uintptr(Rw))
	return uint32(r1)
}

func (x *Xed) OperandActionReadOnly(Rw Xed_operand_action_enum_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_action_read_only.Call(uintptr(Rw))
	return uint32(r1)
}

func (x *Xed) OperandActionWritten(Rw Xed_operand_action_enum_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_action_written.Call(uintptr(Rw))
	return uint32(r1)
}

func (x *Xed) OperandActionWrittenOnly(Rw Xed_operand_action_enum_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_action_written_only.Call(uintptr(Rw))
	return uint32(r1)
}

func (x *Xed) OperandActionReadAndWritten(Rw Xed_operand_action_enum_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_action_read_and_written.Call(uintptr(Rw))
	return uint32(r1)
}

func (x *Xed) OperandActionConditionalRead(Rw Xed_operand_action_enum_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_action_conditional_read.Call(uintptr(Rw))
	return uint32(r1)
}

func (x *Xed) OperandActionConditionalWrite(Rw Xed_operand_action_enum_t) Xed_uint_t {
	r1, _, _ := proc_xed_operand_action_conditional_write.Call(uintptr(Rw))
	return uint32(r1)
}

func (x *Xed) GetVersion() *int8 {
	r1, _, _ := proc_xed_get_version.Call()
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) GetCopyright() *int8 {
	r1, _, _ := proc_xed_get_copyright.Call()
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) FormatSetOptions(Format_options Xed_format_options_t) {
	proc_xed_format_set_options.Call(uintptr(unsafe.Pointer(&Format_options)))
}

func (x *Xed) InitPrintInfo(Pi *Xed_print_info_t) {
	proc_xed_init_print_info.Call(uintptr(unsafe.Pointer(Pi)))
}

func (x *Xed) DecodedInstGetAttribute(P *Xed_decoded_inst_t, Attr Xed_attribute_enum_t) Xed_uint32_t {
	r1, _, _ := proc_xed_decoded_inst_get_attribute.Call(uintptr(unsafe.Pointer(P)), uintptr(Attr))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetAttributes(P *Xed_decoded_inst_t) Xed_attributes_t {
	r1, _, _ := proc_xed_decoded_inst_get_attributes.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_attributes_t)(unsafe.Pointer(&r1))
}

func (x *Xed) DecodedInstIsXacquire(P *Xed_decoded_inst_t) Xed_uint32_t {
	r1, _, _ := proc_xed_decoded_inst_is_xacquire.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstIsXrelease(P *Xed_decoded_inst_t) Xed_uint32_t {
	r1, _, _ := proc_xed_decoded_inst_is_xrelease.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstHasMpxPrefix(P *Xed_decoded_inst_t) Xed_uint32_t {
	r1, _, _ := proc_xed_decoded_inst_has_mpx_prefix.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstIsApxZu(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_is_apx_zu.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetModrm(P *Xed_decoded_inst_t) Xed_uint8_t {
	r1, _, _ := proc_xed_decoded_inst_get_modrm.Call(uintptr(unsafe.Pointer(P)))
	return uint8(r1)
}

func (x *Xed) DecodedInstMaskedVectorOperation(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_masked_vector_operation.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstVectorLengthBits(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_vector_length_bits.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetNprefixes(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_nprefixes.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstOperandLengthBits(P *Xed_decoded_inst_t, Operand_index uint32) uint32 {
	r1, _, _ := proc_xed_decoded_inst_operand_length_bits.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_index))
	return uint32(r1)
}

func (x *Xed) DecodedInstOperandLength(P *Xed_decoded_inst_t, Operand_index uint32) uint32 {
	r1, _, _ := proc_xed_decoded_inst_operand_length.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_index))
	return uint32(r1)
}

func (x *Xed) DecodedInstOperandElements(P *Xed_decoded_inst_t, Operand_index uint32) uint32 {
	r1, _, _ := proc_xed_decoded_inst_operand_elements.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_index))
	return uint32(r1)
}

func (x *Xed) DecodedInstOperandElementSizeBits(P *Xed_decoded_inst_t, Operand_index uint32) uint32 {
	r1, _, _ := proc_xed_decoded_inst_operand_element_size_bits.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_index))
	return uint32(r1)
}

func (x *Xed) DecodedInstOperandElementType(P *Xed_decoded_inst_t, Operand_index uint32) Xed_operand_element_type_enum_t {
	r1, _, _ := proc_xed_decoded_inst_operand_element_type.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_index))
	return Xed_operand_element_type_enum_t(uint32(r1))
}

func (x *Xed) DecodedInstOperandAction(P *Xed_decoded_inst_t, Operand_index uint32) Xed_operand_action_enum_t {
	r1, _, _ := proc_xed_decoded_inst_operand_action.Call(uintptr(unsafe.Pointer(P)), uintptr(Operand_index))
	return Xed_operand_action_enum_t(uint32(r1))
}

func (x *Xed) DecodedInstMasking(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_masking.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstMerging(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_merging.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstZeroing(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_zeroing.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstAvx512DestElements(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_avx512_dest_elements.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstZeroSetMode(P *Xed_decoded_inst_t, Dstate *Xed_state_t) {
	proc_xed_decoded_inst_zero_set_mode.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Dstate)))
}

func (x *Xed) DecodedInstZeroKeepMode(P *Xed_decoded_inst_t) {
	proc_xed_decoded_inst_zero_keep_mode.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) DecodedInstZero(P *Xed_decoded_inst_t) {
	proc_xed_decoded_inst_zero.Call(uintptr(unsafe.Pointer(P)))
}

func (x *Xed) DecodedInstZeroKeepModeFromOperands(P *Xed_decoded_inst_t, Operands *Xed_operand_values_t) {
	proc_xed_decoded_inst_zero_keep_mode_from_operands.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Operands)))
}

func (x *Xed) DecodedInstGetOperandWidth(P *Xed_decoded_inst_t) Xed_uint32_t {
	r1, _, _ := proc_xed_decoded_inst_get_operand_width.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstValidForChip(P *Xed_decoded_inst_t, Chip Xed_chip_enum_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_valid_for_chip.Call(uintptr(unsafe.Pointer(P)), uintptr(Chip))
	return uint32(r1)
}

func (x *Xed) DecodedInstValidForFeatures(P *Xed_decoded_inst_t, Chip_features *Xed_chip_features_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_valid_for_features.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Chip_features)))
	return uint32(r1)
}

func (x *Xed) DecodedInstDump(P *Xed_decoded_inst_t, Buf *int8, Buflen int32) {
	proc_xed_decoded_inst_dump.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen))
}

func (x *Xed) DecodedInstDumpXedFormat(P *Xed_decoded_inst_t, Buf *int8, Buflen int32, Runtime_address Xed_uint64_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_dump_xed_format.Call(uintptr(unsafe.Pointer(P)), uintptr(unsafe.Pointer(Buf)), uintptr(Buflen), *(*uintptr)(unsafe.Pointer(&Runtime_address)))
	return uint32(r1)
}

func (x *Xed) FormatContext(Syntax Xed_syntax_enum_t, Xedd *Xed_decoded_inst_t, Out_buffer *int8, Buffer_len int32, Runtime_instruction_address Xed_uint64_t, Context unsafe.Pointer, Symbolic_callback Xed_disassembly_callback_fn_t) Xed_bool_t {
	r1, _, _ := proc_xed_format_context.Call(uintptr(Syntax), uintptr(unsafe.Pointer(Xedd)), uintptr(unsafe.Pointer(Out_buffer)), uintptr(Buffer_len), *(*uintptr)(unsafe.Pointer(&Runtime_instruction_address)), uintptr(Context), windows.NewCallback(Symbolic_callback))
	return uint32(r1)
}

func (x *Xed) FormatGeneric(Pi *Xed_print_info_t) Xed_bool_t {
	r1, _, _ := proc_xed_format_generic.Call(uintptr(unsafe.Pointer(Pi)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetSegReg(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_reg_enum_t {
	r1, _, _ := proc_xed_decoded_inst_get_seg_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) DecodedInstGetBaseReg(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_reg_enum_t {
	r1, _, _ := proc_xed_decoded_inst_get_base_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) DecodedInstGetIndexReg(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_reg_enum_t {
	r1, _, _ := proc_xed_decoded_inst_get_index_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) DecodedInstGetScale(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_scale.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetMemoryDisplacement(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_int64_t {
	r1, _, _ := proc_xed_decoded_inst_get_memory_displacement.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) DecodedInstGetMemoryDisplacementWidth(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_memory_displacement_width.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetMemoryDisplacementWidthBits(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_memory_displacement_width_bits.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetBranchDisplacement(P *Xed_decoded_inst_t) Xed_int64_t {
	r1, _, _ := proc_xed_decoded_inst_get_branch_displacement.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_int64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) DecodedInstGetBranchDisplacementWidth(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_branch_displacement_width.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetBranchDisplacementWidthBits(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_branch_displacement_width_bits.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetUnsignedImmediate(P *Xed_decoded_inst_t) Xed_uint64_t {
	r1, _, _ := proc_xed_decoded_inst_get_unsigned_immediate.Call(uintptr(unsafe.Pointer(P)))
	return *(*Xed_uint64_t)(unsafe.Pointer(&r1))
}

func (x *Xed) DecodedInstGetImmediateIsSigned(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_immediate_is_signed.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetImmediateWidth(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_immediate_width.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetImmediateWidthBits(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_get_immediate_width_bits.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetSignedImmediate(P *Xed_decoded_inst_t) Xed_int32_t {
	r1, _, _ := proc_xed_decoded_inst_get_signed_immediate.Call(uintptr(unsafe.Pointer(P)))
	return int32(r1)
}

func (x *Xed) DecodedInstGetReg(P *Xed_decoded_inst_t, Reg_operand Xed_operand_enum_t) Xed_reg_enum_t {
	r1, _, _ := proc_xed_decoded_inst_get_reg.Call(uintptr(unsafe.Pointer(P)), uintptr(Reg_operand))
	return Xed_reg_enum_t(uint32(r1))
}

func (x *Xed) DecodedInstHasDefaultFlagsValues(Xedd *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_has_default_flags_values.Call(uintptr(unsafe.Pointer(Xedd)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetDefaultFlagsValues(Xedd *Xed_decoded_inst_t, P *Xed_flag_dfv_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_get_default_flags_values.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetRflagsInfo(P *Xed_decoded_inst_t) *Xed_simple_flag_t {
	r1, _, _ := proc_xed_decoded_inst_get_rflags_info.Call(uintptr(unsafe.Pointer(P)))
	return (*Xed_simple_flag_t)(unsafe.Pointer(r1))
}

func (x *Xed) DecodedInstUsesRflags(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_uses_rflags.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstNumberOfMemoryOperands(P *Xed_decoded_inst_t) Xed_uint_t {
	r1, _, _ := proc_xed_decoded_inst_number_of_memory_operands.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstMemRead(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_mem_read.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return uint32(r1)
}

func (x *Xed) DecodedInstMemWritten(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_mem_written.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return uint32(r1)
}

func (x *Xed) DecodedInstMemWrittenOnly(P *Xed_decoded_inst_t, Mem_idx uint32) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_mem_written_only.Call(uintptr(unsafe.Pointer(P)), uintptr(Mem_idx))
	return uint32(r1)
}

func (x *Xed) DecodedInstConditionallyWritesRegisters(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_conditionally_writes_registers.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetMemoryOperandLength(P *Xed_decoded_inst_t, Memop_idx uint32) uint32 {
	r1, _, _ := proc_xed_decoded_inst_get_memory_operand_length.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx))
	return uint32(r1)
}

func (x *Xed) DecodedInstGetMemopAddressWidth(P *Xed_decoded_inst_t, Memop_idx Xed_uint_t) uint32 {
	r1, _, _ := proc_xed_decoded_inst_get_memop_address_width.Call(uintptr(unsafe.Pointer(P)), uintptr(Memop_idx))
	return uint32(r1)
}

func (x *Xed) DecodedInstIsPrefetch(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_is_prefetch.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstIsBroadcast(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_is_broadcast.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstIsBroadcastInstruction(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_is_broadcast_instruction.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstUsesEmbeddedBroadcast(P *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_decoded_inst_uses_embedded_broadcast.Call(uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) DecodedInstSetScale(P *Xed_decoded_inst_t, Scale Xed_uint_t) {
	proc_xed_decoded_inst_set_scale.Call(uintptr(unsafe.Pointer(P)), uintptr(Scale))
}

func (x *Xed) DecodedInstSetMemoryDisplacement(P *Xed_decoded_inst_t, Disp Xed_int64_t, Length_bytes Xed_uint_t) {
	proc_xed_decoded_inst_set_memory_displacement.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&Disp)), uintptr(Length_bytes))
}

func (x *Xed) DecodedInstSetBranchDisplacement(P *Xed_decoded_inst_t, Disp Xed_int64_t, Length_bytes Xed_uint_t) {
	proc_xed_decoded_inst_set_branch_displacement.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&Disp)), uintptr(Length_bytes))
}

func (x *Xed) DecodedInstSetImmediateSigned(P *Xed_decoded_inst_t, X Xed_int32_t, Length_bytes Xed_uint_t) {
	proc_xed_decoded_inst_set_immediate_signed.Call(uintptr(unsafe.Pointer(P)), uintptr(X), uintptr(Length_bytes))
}

func (x *Xed) DecodedInstSetImmediateUnsigned(P *Xed_decoded_inst_t, X Xed_uint64_t, Length_bytes Xed_uint_t) {
	proc_xed_decoded_inst_set_immediate_unsigned.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&X)), uintptr(Length_bytes))
}

func (x *Xed) DecodedInstSetMemoryDisplacementBits(P *Xed_decoded_inst_t, Disp Xed_int64_t, Length_bits Xed_uint_t) {
	proc_xed_decoded_inst_set_memory_displacement_bits.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&Disp)), uintptr(Length_bits))
}

func (x *Xed) DecodedInstSetBranchDisplacementBits(P *Xed_decoded_inst_t, Disp Xed_int64_t, Length_bits Xed_uint_t) {
	proc_xed_decoded_inst_set_branch_displacement_bits.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&Disp)), uintptr(Length_bits))
}

func (x *Xed) DecodedInstSetImmediateSignedBits(P *Xed_decoded_inst_t, X Xed_int32_t, Length_bits Xed_uint_t) {
	proc_xed_decoded_inst_set_immediate_signed_bits.Call(uintptr(unsafe.Pointer(P)), uintptr(X), uintptr(Length_bits))
}

func (x *Xed) DecodedInstSetImmediateUnsignedBits(P *Xed_decoded_inst_t, X Xed_uint64_t, Length_bits Xed_uint_t) {
	proc_xed_decoded_inst_set_immediate_unsigned_bits.Call(uintptr(unsafe.Pointer(P)), *(*uintptr)(unsafe.Pointer(&X)), uintptr(Length_bits))
}

func (x *Xed) ClassifyAmx(D *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_classify_amx.Call(uintptr(unsafe.Pointer(D)))
	return uint32(r1)
}

func (x *Xed) ClassifyAvx512(D *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_classify_avx512.Call(uintptr(unsafe.Pointer(D)))
	return uint32(r1)
}

func (x *Xed) ClassifyAvx512Maskop(D *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_classify_avx512_maskop.Call(uintptr(unsafe.Pointer(D)))
	return uint32(r1)
}

func (x *Xed) ClassifyAvx(D *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_classify_avx.Call(uintptr(unsafe.Pointer(D)))
	return uint32(r1)
}

func (x *Xed) ClassifySse(D *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_classify_sse.Call(uintptr(unsafe.Pointer(D)))
	return uint32(r1)
}

func (x *Xed) ClassifyApxFoundation(D *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_classify_apx_foundation.Call(uintptr(unsafe.Pointer(D)))
	return uint32(r1)
}

func (x *Xed) ClassifyApx(D *Xed_decoded_inst_t) Xed_bool_t {
	r1, _, _ := proc_xed_classify_apx.Call(uintptr(unsafe.Pointer(D)))
	return uint32(r1)
}

func (x *Xed) IformflEnumTLast() Xed_iformfl_enum_t {
	r1, _, _ := proc_xed_iformfl_enum_t_last.Call()
	return Xed_iformfl_enum_t(uint32(r1))
}

func (x *Xed) RepRemove(X Xed_iclass_enum_t) Xed_iclass_enum_t {
	r1, _, _ := proc_xed_rep_remove.Call(uintptr(X))
	return Xed_iclass_enum_t(uint32(r1))
}

func (x *Xed) RepeMap(Iclass Xed_iclass_enum_t) Xed_iclass_enum_t {
	r1, _, _ := proc_xed_repe_map.Call(uintptr(Iclass))
	return Xed_iclass_enum_t(uint32(r1))
}

func (x *Xed) RepneMap(Iclass Xed_iclass_enum_t) Xed_iclass_enum_t {
	r1, _, _ := proc_xed_repne_map.Call(uintptr(Iclass))
	return Xed_iclass_enum_t(uint32(r1))
}

func (x *Xed) RepMap(Iclass Xed_iclass_enum_t) Xed_iclass_enum_t {
	r1, _, _ := proc_xed_rep_map.Call(uintptr(Iclass))
	return Xed_iclass_enum_t(uint32(r1))
}

func (x *Xed) NorepMap(Iclass Xed_iclass_enum_t) Xed_iclass_enum_t {
	r1, _, _ := proc_xed_norep_map.Call(uintptr(Iclass))
	return Xed_iclass_enum_t(uint32(r1))
}

func (x *Xed) AgenRegisterCallback(Register_fn Xed_register_callback_fn_t, Segment_fn Xed_segment_base_callback_fn_t) {
	proc_xed_agen_register_callback.Call(windows.NewCallback(Register_fn), windows.NewCallback(Segment_fn))
}

func (x *Xed) Agen(Xedd *Xed_decoded_inst_t, Memop_index uint32, Context unsafe.Pointer, Out_address *Xed_uint64_t) Xed_error_enum_t {
	r1, _, _ := proc_xed_agen.Call(uintptr(unsafe.Pointer(Xedd)), uintptr(Memop_index), uintptr(Context), uintptr(unsafe.Pointer(Out_address)))
	return Xed_error_enum_t(uint32(r1))
}

func (x *Xed) CpuidRecEnumT2str(P Xed_cpuid_rec_enum_t) *int8 {
	r1, _, _ := proc_xed_cpuid_rec_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) CpuidRecEnumTLast() Xed_cpuid_rec_enum_t {
	r1, _, _ := proc_xed_cpuid_rec_enum_t_last.Call()
	return Xed_cpuid_rec_enum_t(uint32(r1))
}

func (x *Xed) CpuidGroupEnumT2str(P Xed_cpuid_group_enum_t) *int8 {
	r1, _, _ := proc_xed_cpuid_group_enum_t2str.Call(uintptr(P))
	return (*int8)(unsafe.Pointer(r1))
}

func (x *Xed) CpuidGroupEnumTLast() Xed_cpuid_group_enum_t {
	r1, _, _ := proc_xed_cpuid_group_enum_t_last.Call()
	return Xed_cpuid_group_enum_t(uint32(r1))
}

func (x *Xed) GetCpuidGroupEnumForIsaSet(Isaset Xed_isa_set_enum_t, I Xed_uint_t) Xed_cpuid_group_enum_t {
	r1, _, _ := proc_xed_get_cpuid_group_enum_for_isa_set.Call(uintptr(Isaset), uintptr(I))
	return Xed_cpuid_group_enum_t(uint32(r1))
}

func (x *Xed) GetCpuidRecEnumForGroup(Group Xed_cpuid_group_enum_t, I Xed_uint_t) Xed_cpuid_rec_enum_t {
	r1, _, _ := proc_xed_get_cpuid_rec_enum_for_group.Call(uintptr(Group), uintptr(I))
	return Xed_cpuid_rec_enum_t(uint32(r1))
}

func (x *Xed) GetCpuidRec(Cpuid_bit Xed_cpuid_rec_enum_t, P *Xed_cpuid_rec_t) Xed_bool_t {
	r1, _, _ := proc_xed_get_cpuid_rec.Call(uintptr(Cpuid_bit), uintptr(unsafe.Pointer(P)))
	return uint32(r1)
}

func (x *Xed) IsaSetIsValidForChip(Isa_set Xed_isa_set_enum_t, Chip Xed_chip_enum_t) Xed_bool_t {
	r1, _, _ := proc_xed_isa_set_is_valid_for_chip.Call(uintptr(Isa_set), uintptr(Chip))
	return uint32(r1)
}
