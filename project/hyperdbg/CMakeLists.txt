cmake_minimum_required(VERSION 3.28)

list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_SOURCE_DIR}")
include(${CMAKE_CURRENT_SOURCE_DIR}/ewdk.cmake)

project(HyperDbgUnified C CXX ASM_MASM)
set_property(GLOBAL PROPERTY USE_FOLDERS ON)




set(WDK_NTDDI_VERSION "0x0A000008")

set(CMAKE_BUILD_TYPE Debug CACHE STRING "Build type" FORCE)
set(CMAKE_C_FLAGS_DEBUG "/Od /Zi /RTC1")
set(CMAKE_C_FLAGS "/Od /Zi /RTC1")
set(CMAKE_CXX_FLAGS_DEBUG "/Od /Zi /RTC1")
set(CMAKE_CXX_FLAGS "/Od /Zi /RTC1")

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

include_directories(
    "."
    "clone/HyperDbg/hyperdbg/include"
    "clone/HyperDbg/hyperdbg/dependencies/zydis/include"
    "clone/HyperDbg/hyperdbg/dependencies/zydis/src"
    "clone/HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/include"
    "clone/HyperDbg/hyperdbg/dependencies"
)

add_definitions(
    -DZYCORE_STATIC_BUILD
    -DZYDIS_STATIC_BUILD
)

file(GLOB ZydisSourceFiles "clone/HyperDbg/hyperdbg/dependencies/zydis/src/*.c")
file(GLOB ZycoreSourceFiles "clone/HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/src/*.c")
set(ZydisSourceFiles ${ZydisSourceFiles} ${ZycoreSourceFiles})

file(GLOB_RECURSE HyperkdSourceFiles "clone/HyperDbg/hyperdbg/hyperkd/code/*.c")
file(GLOB HyperkdAsmFiles "clone/HyperDbg/hyperdbg/hyperkd/code/assembly/*.asm")

file(GLOB_RECURSE HyperhvSourceFiles "clone/HyperDbg/hyperdbg/hyperhv/code/*.c")
list(FILTER HyperhvSourceFiles EXCLUDE REGEX ".*code/common/UnloadDll\\.c$")
file(GLOB HyperhvAsmFiles "clone/HyperDbg/hyperdbg/hyperhv/code/assembly/*.asm")

file(GLOB HyperlogSourceFiles "clone/HyperDbg/hyperdbg/hyperlog/code/*.c")

file(GLOB KdserialSourceFiles "clone/HyperDbg/hyperdbg/kdserial/*.c")

file(GLOB ScriptEvalSourceFiles "clone/HyperDbg/hyperdbg/script-eval/code/*.c")

file(GLOB HyperevadeSourceFiles "clone/HyperDbg/hyperdbg/hyperevade/code/*.c")

file(GLOB HypertraceSourceFiles "clone/HyperDbg/hyperdbg/hypertrace/code/*.c")

file(GLOB IncludeOptimizationsSourceFiles "clone/HyperDbg/hyperdbg/include/components/optimizations/code/*.c")
file(GLOB IncludeSpinlockSourceFiles "clone/HyperDbg/hyperdbg/include/components/spinlock/code/*.c")
file(GLOB IncludePlatformSourceFiles "clone/HyperDbg/hyperdbg/include/platform/kernel/code/*.c")

wdk_add_library(zydis_kernel STATIC ${ZydisSourceFiles})
target_compile_definitions(zydis_kernel PRIVATE ZYAN_NO_LIBC ZYDIS_NO_LIBC)
target_compile_options(zydis_kernel PRIVATE /Od /Zi /RTC1)

wdk_add_library(Hyperhv STATIC ${HyperhvSourceFiles} ${HyperhvAsmFiles} ${IncludeOptimizationsSourceFiles} ${IncludeSpinlockSourceFiles})
target_include_directories(Hyperhv PRIVATE
    "clone/HyperDbg/hyperdbg/hyperhv"
    "clone/HyperDbg/hyperdbg/hyperhv/header"
    "clone/HyperDbg/hyperdbg/hyperhv/code"
    "clone/HyperDbg/hyperdbg/include/components/optimizations/code"
    "clone/HyperDbg/hyperdbg/include/components/spinlock/code"
)
target_compile_definitions(Hyperhv PRIVATE DHYPERDBG_VMM HYPERDBG_HYPEREVADE HYPERDBG_HYPER_LOG HYPERDBG_HYPERTRACE NTDDI_VERSION=0x0A000008)
target_compile_options(Hyperhv PRIVATE /Od /Zi /RTC1)
target_precompile_headers(Hyperhv PRIVATE "clone/HyperDbg/hyperdbg/hyperhv/pch.h")

wdk_add_library(Hyperlog STATIC ${HyperlogSourceFiles})
target_include_directories(Hyperlog PRIVATE
    "clone/HyperDbg/hyperdbg/hyperlog"
    "clone/HyperDbg/hyperdbg/hyperlog/header"
    "clone/HyperDbg/hyperdbg/hyperlog/code"
)
target_compile_definitions(Hyperlog PRIVATE DHYPERDBG_HYPERLOG HYPERDBG_HYPER_LOG HYPERDBG_HYPERTRACE HYPERDBG_HYPEREVADE HYPERDBG_VMM NTDDI_VERSION=0x0A000008)
target_compile_options(Hyperlog PRIVATE /Od /Zi /RTC1)
target_precompile_headers(Hyperlog PRIVATE "clone/HyperDbg/hyperdbg/hyperlog/header/pch.h")

wdk_add_library(Kdserial STATIC ${KdserialSourceFiles})
target_include_directories(Kdserial PRIVATE
    "clone/HyperDbg/hyperdbg/kdserial"
)
target_compile_options(Kdserial PRIVATE /Od /Zi /RTC1)

wdk_add_library(Hyperevade STATIC ${HyperevadeSourceFiles})
target_include_directories(Hyperevade PRIVATE
    "clone/HyperDbg/hyperdbg/hyperevade"
    "clone/HyperDbg/hyperdbg/hyperevade/header"
    "clone/HyperDbg/hyperdbg/hyperevade/code"
)
target_compile_definitions(Hyperevade PRIVATE DHYPERDBG_HYPEREVADE HYPERDBG_HYPEREVADE HYPERDBG_HYPER_LOG HYPERDBG_HYPERTRACE HYPERDBG_VMM NTDDI_VERSION=0x0A000008)
target_compile_options(Hyperevade PRIVATE /Od /Zi /RTC1)
target_precompile_headers(Hyperevade PRIVATE "clone/HyperDbg/hyperdbg/hyperevade/header/pch.h")

wdk_add_library(Hypertrace STATIC ${HypertraceSourceFiles})
target_include_directories(Hypertrace PRIVATE
    "clone/HyperDbg/hyperdbg/hypertrace"
    "clone/HyperDbg/hyperdbg/hypertrace/header"
    "clone/HyperDbg/hyperdbg/hypertrace/code"
)
target_compile_definitions(Hypertrace PRIVATE DHYPERDBG_HYPERTRACE HYPERDBG_HYPERTRACE HYPERDBG_HYPER_LOG HYPERDBG_HYPEREVADE HYPERDBG_VMM NTDDI_VERSION=0x0A000008)
target_compile_options(Hypertrace PRIVATE /Od /Zi /RTC1)
target_precompile_headers(Hypertrace PRIVATE "clone/HyperDbg/hyperdbg/hypertrace/header/pch.h")

wdk_add_library(Platform STATIC ${IncludePlatformSourceFiles})
target_include_directories(Platform PRIVATE
    "clone/HyperDbg/hyperdbg/include/platform/kernel/code"
)
target_compile_definitions(Platform PRIVATE NTDDI_VERSION=0x0A000008)
target_compile_options(Platform PRIVATE /Od /Zi /RTC1)
target_precompile_headers(Platform PRIVATE "clone/HyperDbg/hyperdbg/include/platform/kernel/code/pch.h")

wdk_add_driver(hyperkd
    KMDF 1.15
    ${HyperkdSourceFiles}
    ${HyperkdAsmFiles}
    ${IncludeOptimizationsSourceFiles}
    ${IncludeSpinlockSourceFiles}
    ${ScriptEvalSourceFiles}
)
target_include_directories(hyperkd PRIVATE
    "clone/HyperDbg/hyperdbg/hyperkd"
    "clone/HyperDbg/hyperdbg/hyperkd/header"
    "clone/HyperDbg/hyperdbg/hyperkd/code"
    "clone/HyperDbg/hyperdbg/include/components/optimizations/code"
    "clone/HyperDbg/hyperdbg/include/components/spinlock/code"
    "clone/HyperDbg/hyperdbg/script-eval"
)
target_compile_definitions(hyperkd PRIVATE DHYPERDBG_DEBUGGER HYPERDBG_HYPEREVADE HYPERDBG_HYPER_LOG HYPERDBG_HYPERTRACE HYPERDBG_VMM NTDDI_VERSION=0x0A000008)
target_precompile_headers(hyperkd PRIVATE "clone/HyperDbg/hyperdbg/hyperkd/header/pch.h")
target_compile_options(hyperkd PRIVATE /Od /Zi /RTC1)
target_link_options(hyperkd PRIVATE /DEBUG /FORCE:MULTIPLE)

target_link_libraries(hyperkd
    zydis_kernel
    Hyperhv
    Hyperlog
    Kdserial
    Hyperevade
    Hypertrace
    Platform
)

# ============================================================
# User-mode targets
# ============================================================

set(CMAKE_MSVC_RUNTIME_LIBRARY "MultiThreaded")

um_library(Zydis STATIC ${ZydisSourceFiles})
target_compile_options(Zydis PRIVATE /Od /Zi)

file(GLOB_RECURSE PdbexSourceFiles "clone/HyperDbg/hyperdbg/dependencies/pdbex/Source/*.cpp")
list(FILTER PdbexSourceFiles EXCLUDE REGEX ".*main\\.cpp$")

um_library(Pdbex STATIC ${PdbexSourceFiles})
target_include_directories(Pdbex PRIVATE
    "clone/HyperDbg/hyperdbg/dependencies/pdbex/Source"
    "$ENV{VSINSTALLDIR}DIA SDK\\include"
)
target_compile_definitions(Pdbex PRIVATE _CRT_SECURE_NO_WARNINGS)
target_compile_options(Pdbex PRIVATE /Od /Zi /std:c++latest)

file(GLOB SymbolParserSourceFiles "clone/HyperDbg/hyperdbg/symbol-parser/code/*.cpp")

um_library(SymbolParser STATIC ${SymbolParserSourceFiles})
target_include_directories(SymbolParser PRIVATE
    "clone/HyperDbg/hyperdbg/symbol-parser"
    "clone/HyperDbg/hyperdbg/symbol-parser/header"
    "clone/HyperDbg/hyperdbg/symbol-parser/code"
    "clone/HyperDbg/hyperdbg/include"
    "clone/HyperDbg/hyperdbg/dependencies"
    "clone/HyperDbg/hyperdbg/dependencies/pdbex/Source"
)
target_compile_definitions(SymbolParser PRIVATE HYPERDBG_SYMBOL_PARSER _CRT_SECURE_NO_WARNINGS)
target_compile_options(SymbolParser PRIVATE /Od /Zi)
target_precompile_headers(SymbolParser PRIVATE "clone/HyperDbg/hyperdbg/symbol-parser/pch.h")
target_link_libraries(SymbolParser Pdbex)

file(GLOB ScriptEngineSourceFiles "clone/HyperDbg/hyperdbg/script-engine/code/*.c")
list(APPEND ScriptEngineSourceFiles "clone/HyperDbg/hyperdbg/script-engine/script_include.c")

um_library(ScriptEngine STATIC ${ScriptEngineSourceFiles})
target_include_directories(ScriptEngine PRIVATE
    "clone/HyperDbg/hyperdbg/script-engine"
    "clone/HyperDbg/hyperdbg/script-engine/header"
    "clone/HyperDbg/hyperdbg/script-engine/code"
    "clone/HyperDbg/hyperdbg/include"
    "clone/HyperDbg/hyperdbg/script-eval"
)
target_compile_definitions(ScriptEngine PRIVATE SCRIPTENGINE_EXPORTS HYPERDBG_SCRIPT_ENGINE HYPERDBG_SYMBOL_PARSER _CRT_SECURE_NO_WARNINGS)
target_compile_options(ScriptEngine PRIVATE /Od /Zi)
target_precompile_headers(ScriptEngine PRIVATE "clone/HyperDbg/hyperdbg/script-engine/header/pch.h")
target_link_libraries(ScriptEngine SymbolParser)

file(GLOB_RECURSE LibhyperdbgSourceFiles "clone/HyperDbg/hyperdbg/libhyperdbg/code/*.cpp")
list(FILTER LibhyperdbgSourceFiles EXCLUDE REGEX ".*dllmain\\.cpp$")
list(APPEND LibhyperdbgSourceFiles "clone/HyperDbg/hyperdbg/libhyperdbg/pch.cpp")
file(GLOB LibhyperdbgAsmFiles "clone/HyperDbg/hyperdbg/libhyperdbg/code/assembly/*.asm")
file(GLOB ScriptEvalCppFiles "clone/HyperDbg/hyperdbg/script-eval/code/*.c")

foreach(SRC ${ScriptEvalCppFiles})
    get_filename_component(SRC_NAME ${SRC} NAME)
    configure_file(${SRC} "${CMAKE_CURRENT_BINARY_DIR}/script-eval-cpp/${SRC_NAME}pp" COPYONLY)
endforeach()

file(GLOB ScriptEvalCppRenamedFiles "${CMAKE_CURRENT_BINARY_DIR}/script-eval-cpp/*.cpp")

um_library(ScriptEvalUser STATIC ${ScriptEvalCppRenamedFiles})
target_include_directories(ScriptEvalUser PRIVATE
    "clone/HyperDbg/hyperdbg/script-eval"
    "clone/HyperDbg/hyperdbg/include"
    "clone/HyperDbg/hyperdbg/dependencies"
    "clone/HyperDbg/hyperdbg/dependencies/phnt"
    "clone/HyperDbg/hyperdbg/dependencies/zydis/include"
    "clone/HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/include"
    "clone/HyperDbg/hyperdbg/libhyperdbg"
    "clone/HyperDbg/hyperdbg/script-engine/header"
    "clone/HyperDbg/hyperdbg/symbol-parser/header"
)
target_compile_definitions(ScriptEvalUser PRIVATE
    HYPERDBG_LIBHYPERDBG
    HYPERDBG_SCRIPT_ENGINE
    HYPERDBG_SYMBOL_PARSER
    ZYCORE_STATIC_DEFINE
    ZYDIS_STATIC_DEFINE
    HPRDBGCTRL_EXPORTS
    NTDDI_VERSION=0x06010000
    WINVER=0x0601
    _WIN32_WINNT=0x0601
    _CRT_SECURE_NO_WARNINGS
)
target_compile_options(ScriptEvalUser PRIVATE /Od /Zi)
target_precompile_headers(ScriptEvalUser PRIVATE "clone/HyperDbg/hyperdbg/libhyperdbg/pch.h")

um_library(Libhyperdbg STATIC
    ${LibhyperdbgSourceFiles}
    ${LibhyperdbgAsmFiles}
)
target_include_directories(Libhyperdbg PRIVATE
    "clone/HyperDbg/hyperdbg/libhyperdbg"
    "clone/HyperDbg/hyperdbg/libhyperdbg/header"
    "clone/HyperDbg/hyperdbg/libhyperdbg/code"
    "clone/HyperDbg/hyperdbg/include"
    "clone/HyperDbg/hyperdbg/dependencies"
    "clone/HyperDbg/hyperdbg/dependencies/phnt"
    "clone/HyperDbg/hyperdbg/dependencies/zydis/include"
    "clone/HyperDbg/hyperdbg/dependencies/zydis/dependencies/zycore/include"
    "clone/HyperDbg/hyperdbg/dependencies/keystone/include"
    "clone/HyperDbg/hyperdbg/dependencies/pdbex/Source"
    "clone/HyperDbg/hyperdbg/script-eval"
    "clone/HyperDbg/hyperdbg/script-engine/header"
    "clone/HyperDbg/hyperdbg/symbol-parser/header"
)
target_compile_definitions(Libhyperdbg PRIVATE
    HYPERDBG_LIBHYPERDBG
    HYPERDBG_SCRIPT_ENGINE
    HYPERDBG_SYMBOL_PARSER
    ZYCORE_STATIC_DEFINE
    ZYDIS_STATIC_DEFINE
    HPRDBGCTRL_EXPORTS
    NTDDI_VERSION=0x06010000
    WINVER=0x0601
    _WIN32_WINNT=0x0601
    _CRT_SECURE_NO_WARNINGS
)
target_compile_options(Libhyperdbg PRIVATE /Od /Zi)
set_source_files_properties(${LibhyperdbgAsmFiles} PROPERTIES LANGUAGE ASM_MASM)
target_precompile_headers(Libhyperdbg PRIVATE "clone/HyperDbg/hyperdbg/libhyperdbg/pch.h")
target_link_libraries(Libhyperdbg
    ScriptEngine
    SymbolParser
    Pdbex
    Zydis
    ScriptEvalUser
)

wdk_add_executable(hyperdbg-cli "clone/HyperDbg/hyperdbg/hyperdbg-cli/hyperdbg-cli.cpp")
target_include_directories(hyperdbg-cli PRIVATE
    "clone/HyperDbg/hyperdbg/include"
    "clone/HyperDbg/hyperdbg/dependencies"
)
target_compile_definitions(hyperdbg-cli PRIVATE
    HYPERDBG_LIBHYPERDBG
    HYPERDBG_SCRIPT_ENGINE
    HYPERDBG_SYMBOL_PARSER
    HYPERDBG_DEBUG
    _CRT_SECURE_NO_WARNINGS
)
target_compile_options(hyperdbg-cli PRIVATE /Od /Zi)
target_link_options(hyperdbg-cli PRIVATE /DEBUG /FORCE:MULTIPLE
    /ALTERNATENAME:__imp_pdbex_export=pdbex_export
    /ALTERNATENAME:__imp_pdbex_set_logging_method_export=pdbex_set_logging_method_export
)
target_link_libraries(hyperdbg-cli
    Libhyperdbg
    dbghelp
    urlmon
    keystone
)
target_link_directories(hyperdbg-cli PRIVATE
    "clone/HyperDbg/hyperdbg/libraries/keystone/release-lib"
    "${CMAKE_CURRENT_BINARY_DIR}/$<CONFIG>"
)

add_custom_command(TARGET hyperdbg-cli PRE_LINK
    COMMAND ${CMAKE_COMMAND} -E copy_if_different
        "${CMAKE_CURRENT_BINARY_DIR}/Zydis.lib"
        "${CMAKE_CURRENT_BINARY_DIR}/Zycore.lib"
)
