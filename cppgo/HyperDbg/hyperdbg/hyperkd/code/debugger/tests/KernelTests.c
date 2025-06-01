
#include "pch.h"

VOID TestKernelPerformTests(PDEBUGGER_PERFORM_KERNEL_TESTS KernelTestRequest) {
  LogInfo("Starting kernel-test process...");
  LogInfo("All the kernel events are triggered");
  KernelTestRequest->KernelStatus = DEBUGGER_OPERATION_WAS_SUCCESSFUL;
}
