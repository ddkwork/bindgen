
#include "pch.h"
extern BOOLEAN g_BreakPrintingOutput;
extern BOOLEAN g_IsDebuggerModulesLoaded;
extern BOOLEAN g_AutoUnpause;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsExecutingSymbolLoadingRoutines;
extern BOOLEAN g_IsInstrumentingInstructions;
extern BOOLEAN g_IgnorePauseRequests;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;

BOOL BreakController(DWORD CtrlType) {
  switch (CtrlType) {
    case CTRL_BREAK_EVENT:
    case CTRL_C_EVENT:
      if (g_IgnorePauseRequests) {
        return TRUE;
      }
      if (g_IsExecutingSymbolLoadingRoutines) {
        g_IsExecutingSymbolLoadingRoutines = FALSE;
        ScriptEngineSymbolAbortLoadingWrapper();
      }
      if (g_IsSerialConnectedToRemoteDebuggee) {
        if (g_IsInstrumentingInstructions) {
          g_IsInstrumentingInstructions = FALSE;
        } else {
          KdBreakControlCheckAndPauseDebugger(TRUE);
        }
      } else if (!g_IsDebuggerModulesLoaded && !g_IsConnectedToRemoteDebuggee) {
      } else {
        if (g_IsInstrumentingInstructions) {
          g_IsInstrumentingInstructions = FALSE;
        } else {
          g_BreakPrintingOutput = TRUE;
          if (g_IsConnectedToRemoteDebuggee) {
            RemoteConnectionSendCommand("pause", (UINT32)strlen("pause") + 1);
          }
          Sleep(300);
          if (!g_IsConnectedToRemoteDebuggee) {
            if (g_AutoUnpause) {
              ShowMessages(
                  "\npausing...\nauto-unpause mode is enabled, "
                  "debugger will automatically continue when you run a new "
                  "event command, if you want to change this behaviour then "
                  "run run 'settings autounpause off'\n\n");
            } else {
              ShowMessages(
                  "\npausing...\nauto-unpause mode is disabled, you "
                  "should run 'g' when you want to continue, otherwise run "
                  "'settings "
                  "autounpause on'\n\n");
            }
            HyperDbgShowSignature();
            if (g_ActiveProcessDebuggingState.IsActive) {
              UdPauseProcess(
                  g_ActiveProcessDebuggingState.ProcessDebuggingToken);
            }
          }
        }
      }
      return TRUE;
    case CTRL_CLOSE_EVENT:
      return TRUE;
    case CTRL_LOGOFF_EVENT:
      return FALSE;
    case CTRL_SHUTDOWN_EVENT:
      return FALSE;
    default:
      return FALSE;
  }
}
