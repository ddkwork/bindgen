cmake_minimum_required(VERSION 3.15)

include("C:/Program Files/CMake/bin/ewdk.cmake")

project(glfw C)

set(GLFW_SRC "clone/glfw/src")

set(CommonSources
    ${GLFW_SRC}/internal.h
    ${GLFW_SRC}/platform.h
    ${GLFW_SRC}/mappings.h
    ${GLFW_SRC}/context.c
    ${GLFW_SRC}/init.c
    ${GLFW_SRC}/input.c
    ${GLFW_SRC}/monitor.c
    ${GLFW_SRC}/platform.c
    ${GLFW_SRC}/vulkan.c
    ${GLFW_SRC}/window.c
    ${GLFW_SRC}/egl_context.c
    ${GLFW_SRC}/osmesa_context.c
    ${GLFW_SRC}/null_platform.h
    ${GLFW_SRC}/null_joystick.h
    ${GLFW_SRC}/null_init.c
    ${GLFW_SRC}/null_monitor.c
    ${GLFW_SRC}/null_window.c
    ${GLFW_SRC}/null_joystick.c
)

set(Win32Sources
    ${GLFW_SRC}/win32_platform.h
    ${GLFW_SRC}/win32_joystick.h
    ${GLFW_SRC}/win32_init.c
    ${GLFW_SRC}/win32_joystick.c
    ${GLFW_SRC}/win32_monitor.c
    ${GLFW_SRC}/win32_window.c
    ${GLFW_SRC}/wgl_context.c
    ${GLFW_SRC}/win32_time.h
    ${GLFW_SRC}/win32_thread.h
    ${GLFW_SRC}/win32_module.c
    ${GLFW_SRC}/win32_time.c
    ${GLFW_SRC}/win32_thread.c
)

um_dll(glfw3 ${CommonSources} ${Win32Sources})

target_include_directories(glfw3
    PUBLIC
        "${CMAKE_CURRENT_SOURCE_DIR}/clone/glfw/include"
    PRIVATE
        "${CMAKE_CURRENT_SOURCE_DIR}/clone/glfw/src"
)

target_compile_definitions(glfw3
    PRIVATE
        _GLFW_WIN32
        _GLFW_BUILD_DLL
        _CRT_SECURE_NO_WARNINGS
)

target_link_libraries(glfw3 gdi32.lib user32.lib shell32.lib)

target_compile_options(glfw3 PRIVATE /O2 /Zi)
