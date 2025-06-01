
#define HYPERDBG_KERNEL_MODE
#define HYPERDBG_RM
#include <Windef.h>
#include <ntifs.h>

#include "Configuration.h"
#include "SDK/HyperDbgSdk.h"
#include "SDK/imports/kernel/HyperDbgHyperLogImports.h"
#include "SDK/imports/kernel/HyperDbgHyperLogIntrinsics.h"
#include "SDK/imports/kernel/HyperDbgVmmImports.h"
#include "SDK/modules/HyperLog.h"
#include "SDK/modules/VMM.h"
#include "components/spinlock/header/Spinlock.h"
#include "core/Core.h"
#include "driver/Driver.h"
#include "driver/Loader.h"
#include "macros/MetaMacros.h"
#include "misc/Global.h"
