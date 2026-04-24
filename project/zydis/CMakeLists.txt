cmake_minimum_required(VERSION 3.15)

list(APPEND CMAKE_MODULE_PATH "${CMAKE_CURRENT_SOURCE_DIR}")
include(${CMAKE_CURRENT_SOURCE_DIR}/ewdk.cmake)

project(zydis C)



file(GLOB ZydisSourceFiles "zydis/src/*.c")
file(GLOB ZycoreSourceFiles "zydis/dependencies/zycore/src/*.c")
set(AllSourceFiles ${ZydisSourceFiles} ${ZycoreSourceFiles})

um_dll(zydis SHARED ${AllSourceFiles})

target_include_directories(zydis
    PUBLIC
        "${CMAKE_CURRENT_SOURCE_DIR}/zydis/include"
    PRIVATE
        "${CMAKE_CURRENT_SOURCE_DIR}/zydis/src"
        "${CMAKE_CURRENT_SOURCE_DIR}/zydis/dependencies/zycore/include"
)

target_compile_definitions(zydis
    PRIVATE
        ZYDIS_SHOULD_EXPORT
        ZYCORE_SHOULD_EXPORT
        _CRT_SECURE_NO_WARNINGS
)

target_compile_options(zydis PRIVATE /O2 /Zi)
