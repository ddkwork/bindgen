
#include "pch.h"
PVOID g_SharedMessageBuffer = NULL;

int hyperdbg_show_messages(const char *Text) {
  printf("%s", Text);
  return 0;
}

int hyperdbg_show_messages_shared_buffer() {
  printf("%s", (char *)g_SharedMessageBuffer);
  return 0;
}

int hyperdbg_load() {
  char CpuId[13] = {0};
  hyperdbg_u_read_vendor_string(CpuId);
  printf("current processor vendor is : %s\n", CpuId);
  if (strcmp(CpuId, "GenuineIntel") == 0) {
    printf("virtualization technology is vt-x\n");
  } else {
    printf(
        "this program is not designed to run in a non-VT-x "
        "environment !\n");
    return 1;
  }
  if (hyperdbg_u_detect_vmx_support()) {
    printf("vmx operation is supported by your processor\n");
  } else {
    printf("vmx operation is not supported by your processor\n");
    return 1;
  }
  hyperdbg_u_set_text_message_callback(hyperdbg_show_messages);
  hyperdbg_u_connect_remote_debugger_using_named_pipe(
      "\\\\.\\pipe\\HyperDbgPipe", TRUE);
  Sleep(10000);
  hyperdbg_u_run_command((CHAR *)"r");
  hyperdbg_u_run_command((CHAR *)".start path c:\\Windows\\system32\\calc.exe");
  hyperdbg_u_continue_debuggee();
  hyperdbg_u_continue_debuggee();
  hyperdbg_u_continue_debuggee();
  hyperdbg_u_continue_debuggee();
  hyperdbg_u_continue_debuggee();
  hyperdbg_u_continue_debuggee();
  hyperdbg_u_continue_debuggee();
  return 0;
}

int main() {
  if (hyperdbg_load() == 0) {
  } else {
    printf("err, in loading HyperDbg\n");
  }
}
