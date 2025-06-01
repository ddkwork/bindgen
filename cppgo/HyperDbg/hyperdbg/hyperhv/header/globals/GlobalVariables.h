
#pragma once

COMPATIBILITY_CHECKS_STATUS g_CompatibilityCheck;

VMM_CALLBACKS g_Callbacks;

VIRTUAL_MACHINE_STATE *g_GuestState;

MEMORY_MAPPER_ADDRESSES *g_MemoryMapper;

EPT_STATE *g_EptState;

LIST_ENTRY g_EptHook2sDetourListHead;

BOOLEAN g_IsEptHook2sDetourListInitialized;

BOOLEAN g_TransparentMode;

TRANSPARENT_MODE_TRAP_FLAG_STATE *g_TransparentModeTrapFlagState;

VOID *g_ApicBase;

VOID *g_IoApicBase;

BOOLEAN g_NmiBroadcastingInitialized;

PVOID g_NmiHandlerForKeDeregisterNmiCallback;

BOOLEAN g_IsUnsafeSyscallOrSysretHandling;

UINT64 *g_MsrBitmapInvalidMsrs;

BOOLEAN g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger;

BOOLEAN g_ModeBasedExecutionControlState;

BOOLEAN g_TriggerEventForVmcalls;

BOOLEAN g_TriggerEventForCpuids;

BOOLEAN g_ExecTrapInitialized;

BOOLEAN g_ExecTrapUnInitializationStarted;

USER_KERNEL_EXECUTION_TRAP_STATE g_ExecTrapState;

BOOLEAN g_IsInterceptingInstructions;

BOOLEAN g_WaitingForInterruptWindowToInjectPageFault;

UINT64 g_PageFaultInjectionAddressFrom;

UINT64 g_PageFaultInjectionAddressTo;

UINT32 g_PageFaultInjectionErrorCode;
