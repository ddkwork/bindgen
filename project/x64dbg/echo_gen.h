﻿//
// Created by Admin on 28/05/2025.
//
#ifndef MCPX64DBG_ECHO_GEN_H
#define MCPX64DBG_ECHO_GEN_H
#include "pluginsdk/bridgemain.h"
#include "pluginsdk/_plugins.h"
#include "pluginsdk/_scriptapi_argument.h"
#include "pluginsdk/_scriptapi_assembler.h"
#include "pluginsdk/_scriptapi_bookmark.h"
#include "pluginsdk/_scriptapi_comment.h"
#include "pluginsdk/_scriptapi_debug.h"
#include "pluginsdk/_scriptapi_flag.h"
#include "pluginsdk/_scriptapi_function.h"
#include "pluginsdk/_scriptapi_gui.h"
#include "pluginsdk/_scriptapi_label.h"
#include "pluginsdk/_scriptapi_memory.h"
#include "pluginsdk/_scriptapi_misc.h"
#include "pluginsdk/_scriptapi_module.h"
#include "pluginsdk/_scriptapi_pattern.h"
#include "pluginsdk/_scriptapi_register.h"
#include "pluginsdk/_scriptapi_stack.h"
#include "pluginsdk/_scriptapi_symbol.h"
#include "pluginsdk/DeviceNameResolver/DeviceNameResolver.h"
#include "pluginsdk/jansson/jansson.h"
#include "pluginsdk/lz4/lz4file.h"
#include "pluginsdk/TitanEngine/TitanEngine.h"
#include "pluginsdk/XEDParse/XEDParse.h"
#include <winsock2.h>
#include <ws2tcpip.h>
#include <map>
#include <string>
#include <functional>
#include <mutex>
#include <ios>
#include <sstream>
#include <algorithm>
#include "httplib.h"
#include "json.hpp"
#pragma comment(lib, "ws2_32.lib")

using namespace httplib;
using json = nlohmann::json;

Server server;
#define DEFAULT_PORT     8888
#define MAX_REQUEST_SIZE 8192

int        g_pluginHandle;
HANDLE     g_httpServerThread = NULL;
bool       g_httpServerRunning = false;
int        httpPort = DEFAULT_PORT;
std::mutex g_httpMutex;

bool         startHttpServer();
void         stopHttpServer();
void         dispatch(SOCKET clientSocket);
DWORD WINAPI HttpServerThread(LPVOID lpParam);
bool         cbEnableHttpServer(int argc, char *argv[]);
bool         cbSetHttpPort(int argc, char *argv[]);

struct Param {
    std::string name;
    std::string type;
    json        value;
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(Param, name, type, value)
};

struct ApiResponse {
    bool        success;
    std::string type;
    json        result;
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(ApiResponse, success, type, result)
};

/*
 {
   json.Marshaler --> adl_serializer
   json.Marshal   --> to_json
   json.Unmarshal --> from_json
 }
 */

namespace nlohmann {
    template<>
    struct adl_serializer<Script::Module::ModuleInfo> {
        static void to_json(json &j, const Script::Module::ModuleInfo &info) {
            j = {

                    {"base", info.base},
                    {"size", info.size},
                    {"entry", info.entry},
                    {"sectionCount", info.sectionCount},
                    {"name", info.name},
                    {"path", info.path},
            };
        }
    };
}// namespace nlohmann

void dispatch() {
    server.Post("/bridgemain.h/DbgMemFindBaseAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) { params[param.name] = param.value; }
            duint       size = 0;
            ApiResponse resp{.success = true, .type = "bool", .result = DbgMemFindBaseAddr(params["addr"].get<duint>(), &size)};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
    });
    server.Post("/_scriptapi_module.h/GetList", [](const Request &req, Response &res) {
        try {
            BridgeList<Script::Module::ModuleInfo>  bridgeList;
            bool                                    ok = Script::Module::GetList(&bridgeList);
            std::vector<Script::Module::ModuleInfo> moduleVector;
            if (ok) { BridgeList<Script::Module::ModuleInfo>::ToVector(&bridgeList, moduleVector, true); }
            ApiResponse resp{.success = ok, .type = "array", .result = moduleVector};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}.dump(), "application/json"); }
    });
}

DWORD WINAPI HttpServerThread(LPVOID lpParam) {
    dispatch();
    server.listen("localhost", httpPort);
    return 0;
}

// Command callback for toggling HTTP server
bool cbEnableHttpServer(int argc, char *argv[]) {
    if (g_httpServerRunning) {
        _plugin_logputs("Stopping HTTP server...");
        stopHttpServer();
        _plugin_logputs("HTTP server stopped");
    } else {
        _plugin_logputs("Starting HTTP server...");
        if (startHttpServer()) {
            _plugin_logprintf("HTTP server started on port %d\n", httpPort);
        } else {
            _plugin_logputs("Failed to start HTTP server");
        }
    }
    return true;
}

// Command callback for changing HTTP server port
bool cbSetHttpPort(int argc, char *argv[]) {
    if (argc < 2) {
        _plugin_logputs("Usage: httpport [port_number]");
        return false;
    }
    int port;
    try {
        port = std::stoi(argv[1]);
    } catch (const std::exception &) {
        _plugin_logputs("Invalid port number");
        return false;
    }
    if (port <= 0 || port > 65535) {
        _plugin_logputs("Port number must be between 1 and 65535");
        return false;
    }
    httpPort = port;
    if (g_httpServerRunning) {
        _plugin_logputs("Restarting HTTP server with new port...");
        stopHttpServer();
        if (startHttpServer()) {
            _plugin_logprintf("HTTP server restarted on port %d\n", httpPort);
        } else {
            _plugin_logputs("Failed to restart HTTP server");
        }
    } else {
        _plugin_logprintf("HTTP port set to %d\n", httpPort);
    }
    return true;
}

// Start the HTTP server
bool startHttpServer() {
    std::lock_guard<std::mutex> lock(g_httpMutex);
    // Stop existing server if running
    if (g_httpServerRunning) { stopHttpServer(); }
    // Create and start the server thread
    g_httpServerThread = CreateThread(NULL, 0, HttpServerThread, NULL, 0, NULL);
    if (g_httpServerThread == NULL) {
        _plugin_logputs("Failed to create HTTP server thread");
        return false;
    }
    g_httpServerRunning = true;
    return true;
}

// Stop the HTTP server
void stopHttpServer() {
    std::lock_guard<std::mutex> lock(g_httpMutex);
    if (g_httpServerRunning) {
        g_httpServerRunning = false;
        server.stop();
        // Wait for the thread to exit
        if (g_httpServerThread != NULL) {
            WaitForSingleObject(g_httpServerThread, 1000);
            CloseHandle(g_httpServerThread);
            g_httpServerThread = NULL;
        }
    }
}

#endif//MCPX64DBG_ECHO_GEN_H
