
#pragma once
#define _NO_CRT_STDIO_INLINE
#pragma warning(disable : 4201)
#include "platform/kernel/header/Environment.h"
#ifdef ENV_WINDOWS
#include <Windef.h>
#include <ntddk.h>
#include <ntstrsafe.h>
#endif
#define HYPERDBG_KERNEL_MODE
#define HYPERDBG_HYPER_LOG
#include "Logging.h"
#include "SDK/HyperDbgSdk.h"
#include "SDK/imports/kernel/HyperDbgHyperLogImports.h"
#include "SDK/modules/HyperLog.h"
#include "UnloadDll.h"
#include "components/spinlock/header/Spinlock.h"
#include "platform/kernel/header/Mem.h"
