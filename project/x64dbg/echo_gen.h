//
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
    server.Post("/_scriptapi_register.h/SetDR3", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/InputLine", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSearchGetCellContent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR13", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiRepaintTableView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_stack.h/Pop", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetFunctionTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetAutoBookmarkAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetAL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemMap", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR14", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/ParseExpression", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR12B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/RemoteAlloc", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR10W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetAutoLabelAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateArgumentWidget", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetDebugStateFast", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetAF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/RemoteFree", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR12D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowReferences", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/Write", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsBpDisabled", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/Refresh", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDIL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR9W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/Wait", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiCloseQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAutoCompleteClearAll", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiInvalidateSymbolSource", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAutoCompleteAddCmd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetDisassembly", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsRunning", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgArgumentDel", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/FindMem", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/Free", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgArgumentGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR8", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetLineWindow", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetExports", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateMemoryView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/SetBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetXrefCountAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/InputValue", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/ResolveLabel", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingSetUint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMenuPrepare", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiEnableLog", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctionDel", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
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
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowTrace", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctionOverlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateAllViews", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetBranchDestination", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR8D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeUserDirectory", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR15D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgCmdExecDirect", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WriteByte", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctionAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgReleaseEncodeTypeBuffer", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetESP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetProgress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/DisableBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgWinEvent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WriteDword", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgStackCommentGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR9D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/EntryFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateDisable", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/AddQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetMainThreadId", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/FromString", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgModBaseFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeIsProcessElevated", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgLoopGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeGetNtBuildNumber", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR11D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/GetSize", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Add", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSearchGetRowCount", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSymbolEnum", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetEncodeTypeBuffer", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuAddEntry", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiTypeClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/NameFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetOF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgDelEncodeTypeRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/WriteMem", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/ToNodeList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR11D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolLogClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR12", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionListFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdatePatches", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetTimeWastedCounter", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR15", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR10", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR15D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiExecuteOnGuiThreadEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/InfoFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiFlushLog", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_stack.h/Push", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptBpToggle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetThreadId", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearBookmarkRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolSetProgress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateRegisterView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetECX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR10W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/Find", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetAL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetDebuggeeNotes", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/MessageYesNo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddRecentFile", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolLogAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateSEHChain", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGraphAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR8", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptSetIp", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleEntry", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetZF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemGetPageSize", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgValFromString", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearLabelRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR3", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetCurrentTaskProgress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetImports", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetCF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR15W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/BaseFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetEntryName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgXrefDelAll", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsValidExpression", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetLabelAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptSetInfoLine", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDisableLog", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSelectionSet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleSectionList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateDumpView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceGetCellContent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/Stop", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetBookmarkAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/GetProtect", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptSetIp", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceAddColumn", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetEncodeSizeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR10D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgInit", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WriteWord", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingRead", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetModuleAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgDelEncodeTypeSegment", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetTF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR10B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR7", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Overlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetCF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetTF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR15B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAutoCompleteDelCmd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR7", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiStackDumpAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemRead", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR6", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionListFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetAH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetAF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SizeFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/BaseFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR12W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetDebuggeeNotes", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLogSave", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearAutoCommentRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSelectionGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiExecuteOnGuiThread", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLogClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetIF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/SelectionGetStart", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetSIL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateDisassemblyView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetLoopTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetESI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetDebugState", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_stack.h/Peek", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/Message", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLoadSourceFileEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMenuEntryClicked", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR2", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSelectInSymbolsTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetTebAddress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearAutoBookmarkRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_assembler.h/AssembleMemEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgWinEventGlobal", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR8B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/GetBase", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/PathFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetEncodeType", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeInit", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetSF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR14W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceReloadData", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR10", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemIsValidReadPtr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateTypeWidget", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetCurrentGraph", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowCpu", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeAlloc", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR6", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetEncodeTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgArgumentOverlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetPF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetECX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetAutoFunctionAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateWatchView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingGetUint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetDF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR1", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgCmdExec", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctions", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR11", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR13W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WritePtr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddLogMessage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetAutoCommentAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR13B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeGetDbgVersion", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleSize", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR14", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetProcessId", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/SearchAndReplace", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetDF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateThreadView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR9", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingFlush", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetGlobalNotes", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR9B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiRegisterScriptLanguage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgXrefGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgDisasmAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDumpAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SizeFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR8W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsDebugging", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR11B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuAddSeparator", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadWord", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSettingsUpdated", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgExit", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLogRedirectStop", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBPL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR1", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddInfoLine", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemFindBaseAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetESI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsRunLocked", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/SetProtect", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/StepIn", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetESP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModulePath", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR13B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetWindowHandle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateSideBar", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetZF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgDisasmFastAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/AddNode", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceAddCommand", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetSF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDIL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateTimeWastedCounter", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDisasmAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiTranslateText", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgArgumentAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiIsUpdateDisabled", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiTypeAddNode", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiIsLogEnabled", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleSectionCount", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemWrite", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceGetRowCount", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetSPL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR10D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolRefreshCurrent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsJumpGoingToExecute", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetSearchStartCol", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR0", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptCmdExec", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/StepOut", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR13D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionCountFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceInitialize", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/__debugbreak", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetArgTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgXrefAddMulti", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDumpAtN", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiOpenTraceFile", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR8W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLoadGraph", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR11W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiFocusView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/SelectionGetEnd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/Run", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR2", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/Alloc", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetXrefTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/SearchAndReplaceMem", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgValToString", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/CloseQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetActiveView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetPF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/Size", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR15W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearAutoLabelRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetCellContent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR14D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeStart", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiProcessEvents", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadQword", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_assembler.h/Assemble", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR10B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgLoopOverlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/SelectionSet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolUpdateModuleList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetSPL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetThreadList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgAnalyzeFunction", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR9B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgAssembleAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSymbolEnumRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetEntryIcon", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiFoldDisassembly", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR11B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR8B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR13", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/ShowQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_symbol.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptLoad", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadPtr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetAH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetRegDumpEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetLabelAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadByte", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptAbort", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR12", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptUnload", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptSetTitle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLogRedirect", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/IsTemporary", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR9", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR14B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptError", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/DeleteHardwareBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetGlobalNotes", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/EntryFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/Read", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeIsARM64Emulated", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateWindowTitle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionCountFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/ToGraphList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleBase", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuRemove", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR13W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Add", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR15", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateCallStack", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCFLAGS", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetSymbolInfoAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetSymbolInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptGetLineType", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateTraceBrowser", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/DeleteBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearAutoFunctionRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetStringAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR11W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetDebugEngine", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetCommentAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddFavouriteCommand", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/PathFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR14W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctionGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptGetBranchInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingSet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetLastException", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiCloseApplication", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateGraphView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateBreakpointsView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetCommentAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Overlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/StepOver", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetPebAddress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR12B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetIF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/Free", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDisplayWarning", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUnregisterScriptLanguage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptMessage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetThreadHandle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetWatchList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetRowCount", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/AddParent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadDword", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR14D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptRun", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WriteQword", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetSIL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptBpGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgXrefAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddLogMessageHtml", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGotoTrace", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddFavouriteTool", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBPL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/Write", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearCommentRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/RemoteGetProcAddress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/SetHardwareBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetBpList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_assembler.h/AssembleMem", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetProcessHandle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSymbolEnumFromCache", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetOF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceDeleteAllColumns", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgLoopDel", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR15B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddStatusBarMessage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSelectInMemoryMap", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR8D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR14B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetEntryHotkey", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR11", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR0", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/Pause", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgLoopAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/InfoFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/IsValidPtr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_assembler.h/AssembleEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowThreads", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR9W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR9D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptStep", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetIcon", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetFavouriteToolShortcut", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCFLAGS", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSelChanged", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/SelectionGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptMsgyn", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeFree", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR13D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR12W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR12D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetBpxTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
              //todo call api with params
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
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

void stopHttpServer() {
    std::lock_guard<std::mutex> lock(g_httpMutex);
    if (g_httpServerRunning) {
        g_httpServerRunning = false;
        server.stop();
        if (g_httpServerThread != NULL) {// Wait for the thread to exit
            WaitForSingleObject(g_httpServerThread, 1000);
            CloseHandle(g_httpServerThread);
            g_httpServerThread = NULL;
        }
    }
}

#endif//MCPX64DBG_ECHO_GEN_H
