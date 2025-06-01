
#pragma once

#define DEFAULT_INITIAL_DEBUGGEE_TO_DEBUGGER_OFFSET 0x200

#define DEFAULT_INITIAL_DEBUGGER_TO_DEBUGGEE_OFFSET 0x0

#define DEFAULT_INITIAL_BRAM_BUFFER_SIZE 256

#define HWDBG_TEST_READ_INSTANCE_INFO_PATH                              \
  "..\\..\\..\\..\\hwdbg\\sim\\hwdbg\\DebuggerModuleTestingBRAM\\bram_" \
  "instance_info.txt"

#define HWDBG_TEST_WRITE_SCRIPT_BUFFER_PATH \
  "..\\..\\..\\..\\hwdbg\\src\\test\\bram\\script_buffer.hex.txt"

#define HWDBG_TEST_WRITE_INSTANCE_INFO_PATH \
  "..\\..\\..\\..\\hwdbg\\src\\test\\bram\\instance_info.hex.txt"

typedef enum _HWDBG_ACTION_ENUMS {
  hwdbgActionSendInstanceInfo = 1,
  hwdbgActionConfigureScriptBuffer = 2,
} HWDBG_ACTION_ENUMS;

typedef enum _HWDBG_RESPONSE_ENUMS {
  hwdbgResponseSuccessOrErrorMessage = 1,
  hwdbgResponseInstanceInfo = 2,
} HWDBG_RESPONSE_ENUMS;

typedef enum _HWDBG_SUCCESS_OR_ERROR_ENUMS {
  hwdbgOperationWasSuccessful = 0x7FFFFFFF,
  hwdbgErrorInvalidPacket = 1,
} HWDBG_SUCCESS_OR_ERROR_ENUMS;

typedef struct _HWDBG_PORT_INFORMATION_ITEMS {
  UINT32 PortSize;
} HWDBG_PORT_INFORMATION_ITEMS, *PHWDBG_PORT_INFORMATION_ITEMS;

#pragma pack(push, 4)

typedef struct _HWDBG_INSTANCE_INFORMATION {
  UINT32 version;
  UINT32 maximumNumberOfStages;
  UINT32 scriptVariableLength;
  UINT32 numberOfSupportedLocalAndGlobalVariables;
  UINT32 numberOfSupportedTemporaryVariables;
  UINT32 maximumNumberOfSupportedGetScriptOperators;
  UINT32 maximumNumberOfSupportedSetScriptOperators;
  UINT32 sharedMemorySize;
  UINT32 debuggerAreaOffset;
  UINT32 debuggeeAreaOffset;
  UINT32 numberOfPins;
  UINT32 numberOfPorts;

  struct _HWDBG_SCRIPT_CAPABILITIES {
    UINT64 assign_local_global_var : 1;
    UINT64 assign_registers : 1;
    UINT64 assign_pseudo_registers : 1;
    UINT64 conditional_statements_and_comparison_operators : 1;
    UINT64 stack_assignments : 1;
    UINT64 func_or : 1;
    UINT64 func_xor : 1;
    UINT64 func_and : 1;
    UINT64 func_asr : 1;
    UINT64 func_asl : 1;
    UINT64 func_add : 1;
    UINT64 func_sub : 1;
    UINT64 func_mul : 1;
    UINT64 func_div : 1;
    UINT64 func_mod : 1;
    UINT64 func_gt : 1;
    UINT64 func_lt : 1;
    UINT64 func_egt : 1;
    UINT64 func_elt : 1;
    UINT64 func_equal : 1;
    UINT64 func_neq : 1;
    UINT64 func_jmp : 1;
    UINT64 func_jz : 1;
    UINT64 func_jnz : 1;
    UINT64 func_mov : 1;
    UINT64 func_printf : 1;
  } scriptCapabilities;

  UINT32 bramAddrWidth;
  UINT32 bramDataWidth;
} HWDBG_INSTANCE_INFORMATION, *PHWDBG_INSTANCE_INFORMATION;

#pragma pack(pop)

typedef struct _HWDBG_SCRIPT_BUFFER {
  UINT32 scriptNumberOfSymbols;
} HWDBG_SCRIPT_BUFFER, *PHWDBG_SCRIPT_BUFFER;
