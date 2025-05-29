//+build ignore

#include "pluginsdk/_plugins.h"
#include "echo_gen.h"
#include <algorithm>

// Plugin information
#define PLUGIN_NAME    "x64dbg HTTP Server"
#define PLUGIN_VERSION 1

void registerCommands();

//=============================================================================
// Plugin Interface Implementation
//=============================================================================

// Initialize the plugin
bool pluginInit(PLUG_INITSTRUCT *initStruct) {
    initStruct->pluginVersion = PLUGIN_VERSION;
    initStruct->sdkVersion = PLUG_SDKVERSION;
    strncpy_s(initStruct->pluginName, PLUGIN_NAME, _TRUNCATE);
    g_pluginHandle = initStruct->pluginHandle;

    _plugin_logputs("x64dbg HTTP Server plugin loading...");

    // Register commands
    registerCommands();

    // Start the HTTP server
    if (startHttpServer()) {
        _plugin_logprintf("x64dbg HTTP Server started on port %d\n", g_httpPort);
    } else {
        _plugin_logputs("Failed to start HTTP server!");
    }

    _plugin_logputs("x64dbg HTTP Server plugin loaded!");
    return true;
}

// Stop the plugin
void pluginStop() {
    _plugin_logputs("Stopping x64dbg HTTP Server...");
    stopHttpServer();
    _plugin_logputs("x64dbg HTTP Server stopped.");
}

// Plugin setup
bool pluginSetup() { return true; }

// Plugin exports
extern "C" __declspec(dllexport) bool pluginit(PLUG_INITSTRUCT *initStruct) { return pluginInit(initStruct); }

extern "C" __declspec(dllexport) void plugstop() { pluginStop(); }

extern "C" __declspec(dllexport) void plugsetup(PLUG_SETUPSTRUCT *setupStruct) { pluginSetup(); }

// Register plugin commands
void registerCommands() {
    _plugin_registercommand(g_pluginHandle, "httpserver", cbEnableHttpServer, "Toggle HTTP server on/off");
    _plugin_registercommand(g_pluginHandle, "httpport", cbSetHttpPort, "Set HTTP server port");
}