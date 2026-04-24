cmake_minimum_required(VERSION 3.22)

list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_SOURCE_DIR}")
include(${CMAKE_CURRENT_SOURCE_DIR}/ewdk.cmake)

project(AllDemos)
set(CMAKE_CXX_STANDARD 14)
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS}")
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS}")



set(LLVM_TARGETS_TO_BUILD "all")
if(LLVM_TARGETS_TO_BUILD STREQUAL "all")
    set(LLVM_TARGETS_TO_BUILD X86 ARM AArch64 Mips PowerPC Sparc SystemZ Hexagon RISCV)
endif()

file(GLOB_RECURSE src_MC "keystone/llvm/lib/MC/*.cpp")
file(GLOB src_Support "keystone/llvm/lib/Support/*.c*")
file(GLOB src_Target "keystone/llvm/lib/Target/*.cpp")

set(src_core ${src_MC} ${src_Support} ${src_Target})

file(GLOB_RECURSE src_AArch64 "keystone/llvm/lib/Target/AArch64/*.cpp")
file(GLOB_RECURSE src_ARM "keystone/llvm/lib/Target/ARM/*.cpp")
file(GLOB_RECURSE src_Hexagon "keystone/llvm/lib/Target/Hexagon/*.cpp")
file(GLOB_RECURSE src_Mips "keystone/llvm/lib/Target/Mips/*.cpp")
file(GLOB_RECURSE src_PowerPC "keystone/llvm/lib/Target/PowerPC/*.cpp")
file(GLOB_RECURSE src_Sparc "keystone/llvm/lib/Target/Sparc/*.cpp")
file(GLOB_RECURSE src_SystemZ "keystone/llvm/lib/Target/SystemZ/*.cpp")
file(GLOB_RECURSE src_X86 "keystone/llvm/lib/Target/X86/*.cpp")
file(GLOB_RECURSE src_RISCV "keystone/llvm/lib/Target/RISCV/*.cpp")

set(KeystoneSourceFiles ${src_core})
foreach(t ${LLVM_TARGETS_TO_BUILD})
    set(KeystoneSourceFiles ${KeystoneSourceFiles} ${src_${t}})
endforeach()
list(APPEND KeystoneSourceFiles
    "keystone/llvm/keystone/ks.cpp"
    "keystone/llvm/keystone/EVMMapping.cpp"
)

set(LLVM_ENUM_TARGETS "")
foreach(t ${LLVM_TARGETS_TO_BUILD})
    string(APPEND LLVM_ENUM_TARGETS "LLVM_TARGET(${t})\n")
endforeach()

set(LLVM_ENUM_ASM_PARSERS "")
foreach(t ${LLVM_TARGETS_TO_BUILD})
    string(APPEND LLVM_ENUM_ASM_PARSERS "LLVM_ASM_PARSER(${t})\n")
endforeach()

set(HAVE_INTTYPES_H 1)
set(HAVE_STDINT_H 1)
set(HAVE_UINT64_T 1)
set(HAVE_ISATTY 1)
set(LLVM_VERSION_MAJOR 3)
set(LLVM_VERSION_MINOR 9)
set(LLVM_VERSION_PATCH 0)
set(LLVM_ON_WIN32 1)
set(LLVM_ENABLE_THREADS 1)
set(LLVM_HAS_ATOMICS 1)
set(PACKAGE_VERSION "3.9.0svn")
set(PACKAGE_NAME "LLVM")
set(LLVM_DEFAULT_TARGET_TRIPLE "x86_64-pc-windows-msvc")
set(LLVM_HOST_TRIPLE "x86_64-pc-windows-msvc")
set(LLVM_NATIVE_ARCH X86)
set(LLVM_NATIVE_TARGET LLVMInitializeX86Target)
set(LLVM_NATIVE_TARGETINFO LLVMInitializeX86TargetInfo)
set(LLVM_NATIVE_TARGETMC LLVMInitializeX86TargetMC)
set(LLVM_NATIVE_ASMPARSER LLVMInitializeX86AsmParser)

configure_file(
    "keystone/llvm/include/llvm/Config/Targets.def.in"
    "${CMAKE_CURRENT_BINARY_DIR}/include/llvm/Config/Targets.def"
)
configure_file(
    "keystone/llvm/include/llvm/Config/AsmParsers.def.in"
    "${CMAKE_CURRENT_BINARY_DIR}/include/llvm/Config/AsmParsers.def"
)
configure_file(
    "keystone/llvm/include/llvm/Support/DataTypes.h.cmake"
    "${CMAKE_CURRENT_BINARY_DIR}/include/llvm/Support/DataTypes.h"
)
configure_file(
    "keystone/llvm/include/llvm/Config/llvm-config.h.cmake"
    "${CMAKE_CURRENT_BINARY_DIR}/include/llvm/Config/llvm-config.h"
)
configure_file(
    "keystone/llvm/include/llvm/Config/config.h.cmake"
    "${CMAKE_CURRENT_BINARY_DIR}/include/llvm/Config/config.h"
)

um_dll(keystone ${KeystoneSourceFiles})

target_include_directories(keystone PRIVATE
    "keystone/include"
    "keystone"
    "keystone/llvm/include"
    "${CMAKE_CURRENT_BINARY_DIR}/include"
    "keystone/llvm/keystone"
    "keystone/llvm/lib/Target/AArch64"
    "keystone/llvm/lib/Target/ARM"
    "keystone/llvm/lib/Target/Hexagon"
    "keystone/llvm/lib/Target/Mips"
    "keystone/llvm/lib/Target/PowerPC"
    "keystone/llvm/lib/Target/Sparc"
    "keystone/llvm/lib/Target/SystemZ"
    "keystone/llvm/lib/Target/X86"
    "keystone/llvm/lib/Target/RISCV"
)

target_compile_definitions(keystone PRIVATE
    _CRT_SECURE_NO_WARNINGS
    KEYSTONE_EXPORTS
    LLVM_MAJOR_VERSION=3
    LLVM_MINOR_VERSION=9
    LLVM_ENABLE_ARCH_X86
    LLVM_ENABLE_ARCH_ARM
    LLVM_ENABLE_ARCH_AArch64
    LLVM_ENABLE_ARCH_Mips
    LLVM_ENABLE_ARCH_PowerPC
    LLVM_ENABLE_ARCH_Sparc
    LLVM_ENABLE_ARCH_SystemZ
    LLVM_ENABLE_ARCH_Hexagon
    LLVM_ENABLE_ARCH_EVM
    LLVM_ENABLE_ARCH_RISCV
)

target_compile_options(keystone PRIVATE /Od /Zi /EHsc)

add_custom_command(TARGET keystone POST_BUILD
    COMMAND ${CMAKE_COMMAND} -E make_directory "${CMAKE_CURRENT_BINARY_DIR}/include/keystone"
    COMMAND ${CMAKE_COMMAND} -E copy_directory "${CMAKE_CURRENT_SOURCE_DIR}/keystone/include/keystone" "${CMAKE_CURRENT_BINARY_DIR}/include/keystone"
    COMMENT "Copying Keystone header files to output directory..."
)