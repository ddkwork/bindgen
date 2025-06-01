
#pragma once
#define HYPERDBG_SCRIPT_ENGINE
#include "platform/user/header/Environment.h"
#define WIN32_LEAN_AND_MEAN
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <windows.h>

#include "../include/SDK/headers/ScriptEngineCommonDefinitions.h"
#include "SDK/HyperDbgSdk.h"
#include "SDK/headers/HardwareDebugger.h"
#include "SDK/imports/user/HyperDbgScriptImports.h"
#include "SDK/imports/user/HyperDbgSymImports.h"
#include "common.h"
#include "globals.h"
#include "hardware.h"
#include "parse-table.h"
#include "scanner.h"
#include "script-engine.h"
#include "type.h"
