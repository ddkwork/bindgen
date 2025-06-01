
#pragma once
#define HYPERDBG_SYMBOL_PARSER
using namespace std;
#include <Windows.h>
#include <strsafe.h>

#include <algorithm>
#include <iomanip>
#include <sstream>
#include <string>
#include <vector>

#include "platform/user/header/Environment.h"
#define _NO_CVCONST_H
#include <DbgHelp.h>
#define USE_LIB_IA32
#if defined(USE_LIB_IA32)
#pragma warning(push, 0)
#include <ia32-doc/out/ia32.h>
#pragma warning(pop)
typedef RFLAGS *PRFLAGS;
#endif
#include "../symbol-parser/header/common-utils.h"
#include "../symbol-parser/header/symbol-parser.h"
#include "Definition.h"
#include "SDK/HyperDbgSdk.h"
#include "SDK/imports/user/HyperDbgLibImports.h"
#include "SDK/imports/user/HyperDbgSymImports.h"
#pragma comment(lib, "dbghelp.lib")
#pragma comment(lib, "Urlmon.lib")
