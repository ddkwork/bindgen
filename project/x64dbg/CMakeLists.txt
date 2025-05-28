cmake_minimum_required(VERSION 3.15)
project(MCPx64dbg)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Define architecture (x86 or x64)
set(BUILD_X64 OFF)

# Set output file extension based on architecture
if (BUILD_X64)
    set(PLUGIN_EXT ".dp64")
    set(ARCH_BITS 64)
    set(ARCH_SUFFIX "_x64")
else ()
    set(PLUGIN_EXT ".dp32")
    set(ARCH_BITS 32)
    set(ARCH_SUFFIX "_x86")
endif ()

# Set paths to x64dbg SDK
set(sdkPath "${CMAKE_SOURCE_DIR}")

# Include directories
include_directories(
        ${sdkPath}
        ${sdkPath}/pluginsdk
        ${sdkPath}/nlohmann  # Add nlohmann include path
        ${CMAKE_SOURCE_DIR}/include
)

# Add source files
file(GLOB SOURCES "*.cpp")

# Create shared library (plugin)
add_library(${PROJECT_NAME} SHARED ${SOURCES})

# Set output name
set_target_properties(${PROJECT_NAME} PROPERTIES
        PREFIX ""
        OUTPUT_NAME "${PROJECT_NAME}"
        SUFFIX "${PLUGIN_EXT}"
)

# Define macro for architecture
if (BUILD_X64)
    target_compile_definitions(${PROJECT_NAME} PRIVATE -D_WIN64)
else ()
    target_compile_definitions(${PROJECT_NAME} PRIVATE -D_WIN32)
endif ()

# Add path to library files
link_directories(${sdkPath}/pluginsdk)

# Check if libraries exist and adjust architecture suffix if needed
if (NOT EXISTS "${sdkPath}/pluginsdk/TitanEngine/TitanEngine${ARCH_SUFFIX}.lib")
    message(STATUS "TitanEngine${ARCH_SUFFIX}.lib not found, trying TitanEngine.lib")
    set(TITAN_ENGINE_LIB "${sdkPath}/pluginsdk/TitanEngine/TitanEngine.lib")
else ()
    set(TITAN_ENGINE_LIB "${sdkPath}/pluginsdk/TitanEngine/TitanEngine${ARCH_SUFFIX}.lib")
endif ()

if (NOT EXISTS "${sdkPath}/pluginsdk/DeviceNameResolver/DeviceNameResolver${ARCH_SUFFIX}.lib")
    message(STATUS "DeviceNameResolver${ARCH_SUFFIX}.lib not found, trying DeviceNameResolver.lib")
    set(DEVICE_RESOLVER_LIB "${sdkPath}/pluginsdk/DeviceNameResolver/DeviceNameResolver.lib")
else ()
    set(DEVICE_RESOLVER_LIB "${sdkPath}/pluginsdk/DeviceNameResolver/DeviceNameResolver${ARCH_SUFFIX}.lib")
endif ()

if (NOT EXISTS "${sdkPath}/pluginsdk/jansson/jansson${ARCH_SUFFIX}.lib")
    message(STATUS "jansson${ARCH_SUFFIX}.lib not found, trying jansson.lib")
    set(JANSSON_LIB "${sdkPath}/pluginsdk/jansson/jansson.lib")
else ()
    set(JANSSON_LIB "${sdkPath}/pluginsdk/jansson/jansson${ARCH_SUFFIX}.lib")
endif ()

if (NOT EXISTS "${sdkPath}/pluginsdk/lz4/lz4${ARCH_SUFFIX}.lib")
    message(STATUS "lz4${ARCH_SUFFIX}.lib not found, trying lz4.lib")
    set(LZ4_LIB "${sdkPath}/pluginsdk/lz4/lz4.lib")
else ()
    set(LZ4_LIB "${sdkPath}/pluginsdk/lz4/lz4${ARCH_SUFFIX}.lib")
endif ()

if (NOT EXISTS "${sdkPath}/pluginsdk/XEDParse/XEDParse${ARCH_SUFFIX}.lib")
    message(STATUS "XEDParse${ARCH_SUFFIX}.lib not found, trying XEDParse.lib")
    set(XEDPARSE_LIB "${sdkPath}/pluginsdk/XEDParse/XEDParse.lib")
else ()
    set(XEDPARSE_LIB "${sdkPath}/pluginsdk/XEDParse/XEDParse${ARCH_SUFFIX}.lib")
endif ()

# Link x64dbg libraries and dependencies
target_link_libraries(${PROJECT_NAME}
        ws2_32                         # For winsock2
        winhttp                        # WinHTTP library
        ${TITAN_ENGINE_LIB}
        ${DEVICE_RESOLVER_LIB}
        ${JANSSON_LIB}
        ${LZ4_LIB}
        ${XEDPARSE_LIB}
)
if (CMAKE_SIZEOF_VOID_P EQUAL 4)
    target_link_libraries(${PROJECT_NAME}
            ${sdkPath}/pluginsdk/x32dbg.lib
            ${sdkPath}/pluginsdk/x32bridge.lib

    )
elseif (CMAKE_SIZEOF_VOID_P EQUAL 8)
    target_link_libraries(${PROJECT_NAME}
            ${sdkPath}/pluginsdk/x64dbg.lib
            ${sdkPath}/pluginsdk/x64bridge.lib
    )
endif ()


# Just show information about where the plugin is built
add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
        COMMAND ${CMAKE_COMMAND} -E echo ""
        COMMAND ${CMAKE_COMMAND} -E echo "=== Plugin build complete ==="
        COMMAND ${CMAKE_COMMAND} -E echo "Plugin built at: ${CMAKE_BINARY_DIR}/${CMAKE_CFG_INTDIR}/${PROJECT_NAME}${PLUGIN_EXT}"
        COMMAND ${CMAKE_COMMAND} -E echo "Install manually to: $ENV{ProgramFiles}/x64dbg/release/x${ARCH_BITS}/plugins/${PROJECT_NAME}${PLUGIN_EXT}"
        COMMAND ${CMAKE_COMMAND} -E echo "================================="
)