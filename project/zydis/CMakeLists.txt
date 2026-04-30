cmake_minimum_required(VERSION 3.15)

include("C:/Program Files/CMake/bin/ewdk.cmake")

project(zydis C)



file(GLOB ZydisSourceFiles "clone/zydis/src/*.c")
file(GLOB ZycoreSourceFiles "clone/zydis/dependencies/zycore/src/*.c")
set(AllSourceFiles ${ZydisSourceFiles} ${ZycoreSourceFiles})

um_dll(zydis SHARED ${AllSourceFiles})

target_include_directories(zydis
    PUBLIC
        "${CMAKE_CURRENT_SOURCE_DIR}/clone/zydis/include"
    PRIVATE
        "${CMAKE_CURRENT_SOURCE_DIR}/clone/zydis/src"
        "${CMAKE_CURRENT_SOURCE_DIR}/clone/zydis/dependencies/zycore/include"
)

target_compile_definitions(zydis
    PRIVATE
        ZYDIS_SHOULD_EXPORT
        ZYCORE_SHOULD_EXPORT
        _CRT_SECURE_NO_WARNINGS
)

target_compile_options(zydis PRIVATE /O2 /Zi)
