
#include "pch.h"
extern TCHAR g_DriverLocation[MAX_PATH];
extern TCHAR g_DriverName[MAX_PATH];
extern BOOLEAN g_UseCustomDriverLocation;

BOOLEAN hyperdbg_u_detect_vmx_support() { return VmxSupportDetection(); }

VOID hyperdbg_u_read_vendor_string(CHAR *vendor_string) {
  CpuReadVendorString(vendor_string);
}

INT hyperdbg_u_load_vmm() { return HyperDbgLoadVmmModule(); }

INT hyperdbg_u_unload_vmm() { return HyperDbgUnloadVmm(); }

INT hyperdbg_u_install_vmm_driver() { return HyperDbgInstallVmmDriver(); }

INT hyperdbg_u_uninstall_vmm_driver() { return HyperDbgUninstallVmmDriver(); }

INT hyperdbg_u_stop_vmm_driver() { return HyperDbgStopVmmDriver(); }

INT hyperdbg_u_run_command(CHAR *command) {
  return HyperDbgInterpreter(command);
}

BOOLEAN hyperdbg_u_test_command_parser(CHAR *command, UINT32 number_of_tokens,
                                       CHAR **tokens_list,
                                       UINT32 *failed_token_num,
                                       UINT32 *failed_token_position) {
  return HyperDbgTestCommandParser(command, number_of_tokens, tokens_list,
                                   failed_token_num, failed_token_position);
}

VOID hyperdbg_u_test_command_parser_show_tokens(CHAR *command) {
  return HyperDbgTestCommandParserShowTokens(command);
}

VOID hyperdbg_u_show_signature() { HyperDbgShowSignature(); }

VOID hyperdbg_u_set_text_message_callback(PVOID handler) {
  SetTextMessageCallback(handler);
}

PVOID hyperdbg_u_set_text_message_callback_using_shared_buffer(PVOID handler) {
  return SetTextMessageCallbackUsingSharedBuffer(handler);
}

VOID hyperdbg_u_unset_text_message_callback() { UnsetTextMessageCallback(); }

INT hyperdbg_u_script_read_file_and_execute_commandline(INT argc,
                                                        CHAR *argv[]) {
  return ScriptReadFileAndExecuteCommandline(argc, argv);
}

BOOLEAN hyperdbg_u_continue_previous_command() {
  return ContinuePreviousCommand();
}

BOOLEAN hyperdbg_u_check_multiline_command(CHAR *current_command,
                                           BOOLEAN reset) {
  return CheckMultilineCommand(current_command, reset);
}

VOID hyperdbg_u_connect_local_debugger() { ConnectLocalDebugger(); }

BOOLEAN hyperdbg_u_connect_remote_debugger(const CHAR *ip, const CHAR *port) {
  return ConnectRemoteDebugger(ip, port);
}

VOID hyperdbg_u_continue_debuggee() { CommandGRequest(); }

VOID hyperdbg_u_pause_debuggee() { CommandPauseRequest(); }

VOID hyperdbg_u_set_breakpoint(UINT64 address, UINT32 pid, UINT32 tid,
                               UINT32 core_numer) {
  CommandBpRequest(address, pid, tid, core_numer);
}

BOOLEAN hyperdbg_u_set_custom_driver_path(CHAR *driver_file_path,
                                          CHAR *driver_name) {
  if (strlen(driver_file_path) > MAX_PATH) {
    ShowMessages("The driver path is too long, the maximum length is %d\n",
                 MAX_PATH);
    return FALSE;
  }
  if (strlen(driver_name) > MAX_PATH) {
    ShowMessages("The driver name is too long, the maximum length is %d\n",
                 MAX_PATH);
    return FALSE;
  }
  strcpy_s(g_DriverLocation, MAX_PATH, driver_file_path);
  strcpy_s(g_DriverName, MAX_PATH, driver_name);
  g_UseCustomDriverLocation = TRUE;
  return TRUE;
}

VOID hyperdbg_u_use_default_driver_path() { g_UseCustomDriverLocation = FALSE; }

BOOLEAN hyperdbg_u_read_memory(
    UINT64 target_address, DEBUGGER_READ_MEMORY_TYPE memory_type,
    DEBUGGER_READ_READING_TYPE reading_Type, UINT32 pid, UINT32 size,
    BOOLEAN get_address_mode, DEBUGGER_READ_MEMORY_ADDRESS_MODE *address_mode,
    BYTE *target_buffer_to_store, UINT32 *return_length) {
  return HyperDbgReadMemory(target_address, memory_type, reading_Type, pid,
                            size, get_address_mode, address_mode,
                            target_buffer_to_store, return_length);
}

VOID hyperdbg_u_show_memory_or_disassemble(
    DEBUGGER_SHOW_MEMORY_STYLE style, UINT64 address,
    DEBUGGER_READ_MEMORY_TYPE memory_type,
    DEBUGGER_READ_READING_TYPE reading_type, UINT32 pid, UINT32 size,
    PDEBUGGER_DT_COMMAND_OPTIONS dt_details) {
  HyperDbgShowMemoryOrDisassemble(style, address, memory_type, reading_type,
                                  pid, size, dt_details);
}

BOOLEAN hyperdbg_u_read_all_registers(GUEST_REGS *guest_registers,
                                      GUEST_EXTRA_REGISTERS *extra_registers) {
  return HyperDbgReadAllRegisters(guest_registers, extra_registers);
}

BOOLEAN hyperdbg_u_read_target_register(REGS_ENUM register_id,
                                        UINT64 *target_register) {
  return HyperDbgReadTargetRegister(register_id, target_register);
}

BOOLEAN hyperdbg_u_write_target_register(REGS_ENUM register_id, UINT64 value) {
  return HyperDbgWriteTargetRegister(register_id, value);
}

BOOLEAN hyperdbg_u_show_all_registers() { return HyperDbgRegisterShowAll(); }

BOOLEAN hyperdbg_u_show_target_register(REGS_ENUM register_id) {
  return HyperDbgRegisterShowTargetRegister(register_id);
}

BOOLEAN hyperdbg_u_write_memory(PVOID destination_address,
                                DEBUGGER_EDIT_MEMORY_TYPE memory_type,
                                UINT32 process_id, PVOID source_address,
                                UINT32 number_of_bytes) {
  return HyperDbgWriteMemory(destination_address, memory_type, process_id,
                             source_address, number_of_bytes);
}

UINT64 hyperdbg_u_get_kernel_base() { return DebuggerGetKernelBase(); }

BOOLEAN hyperdbg_u_connect_remote_debugger_using_com_port(
    const CHAR *port_name, DWORD baudrate, BOOLEAN pause_after_connection) {
  return HyperDbgDebugRemoteDeviceUsingComPort(port_name, baudrate,
                                               pause_after_connection);
}

BOOLEAN hyperdbg_u_connect_remote_debugger_using_named_pipe(
    const CHAR *named_pipe, BOOLEAN pause_after_connection) {
  return HyperDbgDebugRemoteDeviceUsingNamedPipe(named_pipe,
                                                 pause_after_connection);
}

BOOLEAN hyperdbg_u_debug_close_remote_debugger() {
  return HyperDbgDebugCloseRemoteDebugger();
}

BOOLEAN hyperdbg_u_connect_current_debugger_using_com_port(
    const CHAR *port_name, DWORD baudrate) {
  return HyperDbgDebugCurrentDeviceUsingComPort(port_name, baudrate);
}

BOOLEAN hyperdbg_u_start_process(const WCHAR *path) {
  return UdAttachToProcess(NULL, path, NULL, FALSE);
}

BOOLEAN hyperdbg_u_start_process_with_args(const WCHAR *path,
                                           const WCHAR *arguments) {
  return UdAttachToProcess(NULL, path, arguments, FALSE);
}

BOOLEAN hyperdbg_u_assemble_get_length(const CHAR *assembly_code,
                                       UINT64 start_address, UINT32 *length) {
  return HyperDbgAssembleGetLength(assembly_code, start_address, length);
}

BOOLEAN hyperdbg_u_assemble(const CHAR *assembly_code, UINT64 start_address,
                            PVOID buffer_to_store_assembled_data,
                            UINT32 buffer_size) {
  return HyperDbgAssemble(assembly_code, start_address,
                          buffer_to_store_assembled_data, buffer_size);
}

BOOLEAN hyperdbg_u_setup_path_for_filename(const CHAR *filename,
                                           CHAR *file_location,
                                           UINT32 buffer_len,
                                           BOOLEAN check_file_existence) {
  return SetupPathForFileName(filename, file_location, buffer_len,
                              check_file_existence);
}

BOOLEAN hyperdbg_u_stepping_instrumentation_step_in() {
  return SteppingInstrumentationStepIn();
}

BOOLEAN hyperdbg_u_stepping_regular_step_in() {
  return SteppingRegularStepIn();
}

BOOLEAN hyperdbg_u_stepping_step_over() { return SteppingStepOver(); }

BOOLEAN hyperdbg_u_stepping_instrumentation_step_in_for_tracking() {
  return SteppingInstrumentationStepInForTracking();
}

BOOLEAN hyperdbg_u_stepping_step_over_for_gu(BOOLEAN last_instruction) {
  return SteppingStepOverForGu(last_instruction);
}

BOOLEAN hyperdbg_u_get_local_apic(PLAPIC_PAGE local_apic,
                                  BOOLEAN *is_using_x2apic) {
  return HyperDbgGetLocalApic(local_apic, is_using_x2apic);
}

BOOLEAN hyperdbg_u_get_io_apic(IO_APIC_ENTRY_PACKETS *io_apic) {
  return HyperDbgGetIoApic(io_apic);
}

BOOLEAN hyperdbg_u_get_idt_entry(
    INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS *idt_packet) {
  return HyperDbgGetIdtEntry(idt_packet);
}

BOOLEAN hwdbg_script_run_script(const CHAR *script,
                                const CHAR *instance_filepath_to_read,
                                const CHAR *hardware_script_file_path_to_save,
                                UINT32 initial_bram_buffer_size) {
  return HwdbgScriptRunScript(script, instance_filepath_to_read,
                              hardware_script_file_path_to_save,
                              initial_bram_buffer_size);
}

VOID hwdbg_script_engine_wrapper_test_parser(const char *Expr) {
  std::string StrExpr = Expr;
  ScriptEngineWrapperTestParserForHwdbg(StrExpr);
}

BOOLEAN hyperdbg_u_enable_transparent_mode() {
  return HyperDbgEnableTransparentMode();
}

BOOLEAN hyperdbg_u_disable_transparent_mode() {
  return HyperDbgDisableTransparentMode();
}
