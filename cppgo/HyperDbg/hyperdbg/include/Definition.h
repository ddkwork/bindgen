
#pragma once

#define CONFIG_FILE_NAME L"config.ini"

#define VMM_DRIVER_NAME "hyperhv"

#define KERNEL_DEBUGGER_DRIVER_NAME "hyperkd"

#define KERNEL_DEBUGGER_DRIVER_NAME_AND_EXTENSION "hyperkd.sys"

#define TEST_DEFAULT_NAMED_PIPE "\\\\.\\pipe\\HyperDbgPipe"

#define TEST_CASE_MAXIMUM_BUFFERS_TO_COMMUNICATE 0x1000

#define MAX_HWDBG_TESTING_PIN_COUNT 500

#define COMMAND_PARSER_TEST_CASES_FILE \
  "..\\..\\..\\tests\\command-parser\\command-parser-testcases.txt"

#define TEST_HWDBG_FUNCTIONALITIES "test-hwdbg-functionalities"

#define TEST_CASE_PARAMETER_FOR_MAIN_COMMAND_PARSER "test-command-parser"

#define TEST_CASE_PARAMETER_FOR_SCRIPT_SEMANTIC_TEST_CASES \
  "test-script-semantic-test-cases"

#define SCRIPT_ENGINE_TEST_CASES_DIRECTORY "script-test-cases"

#define SCRIPT_SEMANTIC_TEST_CASE_DIRECTORY \
  "..\\..\\..\\tests\\script-engine-test\\semantic-test-cases"

#define HWDBG_SCRIPT_TEST_CASE_CODE_DIRECTORY \
  "..\\..\\..\\tests\\hwdbg-tests\\scripts\\codes"

#define HWDBG_SCRIPT_TEST_CASE_COMPILED_SCRIPTS_DIRECTORY \
  "..\\..\\..\\tests\\hwdbg-tests\\scripts\\compiled-scripts"

#define HWDBG_SCRIPT_TEST_CASE_SAMPLE_TESTS_DIRECTORY \
  "..\\..\\..\\tests\\hwdbg-tests\\scripts\\sample-tests"

#define DefaultSpeedOfReadingKernelMessages 30
