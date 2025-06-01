
#pragma once
#include "platform/user/header/Environment.h"
#define WIN32_LEAN_AND_MEAN
#define USE_LIB_IA32
#if defined(USE_LIB_IA32)
#pragma warning(push, 0)
#include <ia32-doc/out/ia32.h>
#pragma warning(pop)
typedef RFLAGS *PRFLAGS;
#endif
#define USE_NATIVE_SDK_HEADERS
#define _AMD64_
#if defined(USE__NATIVE_PHNT_HEADERS)
typedef const wchar_t *LPCWCHAR, *PCWCHAR;
#define PHNT_MODE PHNT_MODE_USER
#define PHNT_VERSION PHNT_WIN11
#define PHNT_PATCH_FOR_HYPERDBG TRUE
#include <phnt/phnt.h>
#include <phnt/phnt_windows.h>
#elif defined(USE_NATIVE_SDK_HEADERS)
#include <Windows.h>
#include <platform/user/header/Windows.h>
#include <winioctl.h>
#include <winternl.h>
#endif
#include <VersionHelpers.h>
#include <conio.h>
#include <intrin.h>
#include <inttypes.h>
#include <psapi.h>
#include <shlobj.h>
#include <shlwapi.h>
#include <stdio.h>
#include <stdlib.h>
#include <strsafe.h>
#include <tchar.h>
#include <time.h>
#include <tlhelp32.h>
#include <winsock2.h>
#include <ws2tcpip.h>

#include <algorithm>
#include <array>
#include <bitset>
#include <cctype>
#include <cstring>
#include <fstream>
#include <iomanip>
#include <iostream>
#include <iterator>
#include <list>
#include <locale>
#include <map>
#include <memory>
#include <numeric>
#include <regex>
#include <sstream>
#include <string>
#include <unordered_set>
#include <vector>
#define SCRIPT_ENGINE_USER_MODE
#define HYPERDBG_USER_MODE
#define HYPERDBG_LIBHYPERDBG
#ifndef NDEBUG
#define NDEBUG
#endif
#include "../script-eval/header/ScriptEngineHeader.h"
#include "Configuration.h"
#include "Definition.h"
#include "SDK/HyperDbgSdk.h"
#include "SDK/imports/user/HyperDbgLibImports.h"
#include "SDK/imports/user/HyperDbgScriptImports.h"
#include "header/assembler.h"
#include "header/commands.h"
#include "header/common.h"
#include "header/communication.h"
#include "header/debugger.h"
#include "header/export.h"
#include "header/forwarding.h"
#include "header/help.h"
#include "header/hwdbg-interpreter.h"
#include "header/hwdbg-scripts.h"
#include "header/inipp.h"
#include "header/install.h"
#include "header/kd.h"
#include "header/libhyperdbg.h"
#include "header/list.h"
#include "header/namedpipe.h"
#include "header/objects.h"
#include "header/pe-parser.h"
#include "header/rev-ctrl.h"
#include "header/script-engine.h"
#include "header/steppings.h"
#include "header/symbol.h"
#include "header/tests.h"
#include "header/transparency.h"
#include "header/ud.h"
#include "keystone/keystone.h"
#include "pci-id.h"
#ifdef ENV_WINDOWS
#pragma comment(lib, "ntdll.lib")
#pragma comment(lib, "Shlwapi.lib")
#pragma comment(lib, "Ws2_32.lib")
#pragma comment(lib, "Mswsock.lib")
#pragma comment(lib, "AdvApi32.lib")
#pragma comment(lib, "Psapi.lib")
#pragma comment(lib, "Kernel32.lib")
#endif
