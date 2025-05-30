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
    server.Post("/_scriptapi_flag.h/SetIF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetIF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowCpu", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiShowCpu(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetTF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetTF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::Get(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptSetIp", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptSetIp(params["line"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/SetHardwareBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::SetHardwareBreakpoint(params["address"].get<duint>(), params["Type"].get<HardwareType>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateDisable", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateDisable(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCIP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearAutoLabelRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgClearAutoLabelRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/SetBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::SetBreakpoint(params["address"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetRowCount", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceSetRowCount(params["count"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::GetList(params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR14D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR14D(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptCmdExec", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptCmdExec(params["command"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Comment::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WriteQword", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WriteQword(params["addr"].get<duint>(), params["data"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR12B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR12B(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDumpAtN", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiDumpAtN(params["va"].get<duint>(), params["index"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetBpList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetBpList(params["Type"].get<BPXTYPE>(), params["list"].get<BPMAP *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddStatusBarMessage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAddStatusBarMessage(params["msg"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::Set(params["flag"].get<FlagEnum>(), params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/FromString", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::FromString(params["label"].get<const char *>(), params["addr"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearBookmarkRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgClearBookmarkRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR7", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR7(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgArgumentAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgArgumentAdd(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR12B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR12B(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/ToGraphList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::ToGraphList(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetCurrentTaskProgress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceSetCurrentTaskProgress(params["progress"].get<int>(), params["taskTitle"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR15D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR15D(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiFlushLog", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiFlushLog(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateDumpView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateDumpView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Comment::GetInfo(params["addr"].get<duint>(), params["info"].get<CommentInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetDebugEngine", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetDebugEngine(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetPF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetPF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRBX(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateRegisterView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateRegisterView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRDX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSelectInSymbolsTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSelectInSymbolsTab(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleName(params["name"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateSideBar", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateSideBar(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptSetIp", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiScriptSetIp(params["line"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptBpGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptBpGet(params["line"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetSF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetSF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgAssembleAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgAssembleAt(params["addr"].get<duint>(), params["instruction"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR2", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR2(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR12D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR12D(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBL(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgStackCommentGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgStackCommentGet(params["addr"].get<duint>(), params["comment"].get<STACK_COMMENT *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/Free", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Misc::Free(params["ptr"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR12D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR12D(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR14", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR14(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMenuPrepare", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgMenuPrepare(params["hMenu"].get<GUIMENUTYPE>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetBpxTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetBpxTypeAt(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/SelectionGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::SelectionGet(params["window"].get<Window>(), params["start"].get<duint *>(), params["end"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::Set(params["reg"].get<RegisterEnum>(), params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDI(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetZF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetZF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiFoldDisassembly", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiFoldDisassembly(params["startAddress"].get<duint>(), params["length"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR8D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR8D(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRBP(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR11D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR11D(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgDelEncodeTypeRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgDelEncodeTypeRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetAX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/AddQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::AddQWidgetTab(params["qWidget"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/GetSize", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::GetSize(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceGetCellContent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceGetCellContent(params["row"].get<int>(), params["col"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Comment::GetList(params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEAX(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_assembler.h/Assemble", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Assembler::Assemble(params["addr"].get<duint>(), params["dest"].get<unsigned char *>(), params["size"].get<int *>(), params["instruction"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetLabelAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSetLabelAt(params["addr"].get<duint>(), params["text"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLogRedirectStop", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiLogRedirectStop(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Argument::GetList(params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/ShowQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::ShowQWidgetTab(params["qWidget"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR9D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR9D(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEDI(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR3", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR3(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddRecentFile", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAddRecentFile(params["file"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_assembler.h/AssembleEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Assembler::AssembleEx(params["addr"].get<duint>(), params["dest"].get<unsigned char *>(), params["size"].get<int *>(), params["instruction"].get<const char *>(), params["error"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddLogMessage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAddLogMessage(params["msg"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/Alloc", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Misc::Alloc(params["size"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR8B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR8B(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsValidExpression", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgIsValidExpression(params["expression"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Clear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptBpToggle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptBpToggle(params["line"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctionDel", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgFunctionDel(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLoadSourceFileEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiLoadSourceFileEx(params["path"].get<const char *>(), params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetTimeWastedCounter", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetTimeWastedCounter(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Comment::Get(params["addr"].get<duint>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemMap", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgMemMap(params["memmap"].get<MEMMAP *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::Set(params["info"].get<const BookmarkInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Overlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Function::Overlaps(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/BaseFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::BaseFromName(params["name"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetDebuggeeNotes", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSetDebuggeeNotes(params["text"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetECX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetECX(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Overlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Overlaps(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateDisassemblyView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateDisassemblyView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/Write", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::Write(params["addr"].get<duint>(), params["data"].get<const void *>(), params["size"].get<duint>(), params["sizeWritten"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetRegDumpEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetRegDumpEx(params["regdump"].get<REGDUMP *>(), params["size"].get<size_t>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingRead", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeSettingRead(params["errorLine"].get<int *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCH(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgDisasmFastAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgDisasmFastAt(params["addr"].get<duint>(), params["basicinfo"].get<BASIC_INSTRUCTION_INFO *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceDeleteAllColumns", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceDeleteAllColumns(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCAX(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/InfoFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::InfoFromAddr(params["addr"].get<duint>(), params["info"].get<ModuleInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SizeFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::SizeFromName(params["name"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetLabelAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetLabelAt(params["addr"].get<duint>(), params["segment"].get<SEGMENTREG>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/Pause", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::Pause(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptMsgyn", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiScriptMsgyn(params["message"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSelectionGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSelectionGet(params["hWindow"].get<GUISELECTIONTYPE>(), params["selection"].get<SELECTIONDATA *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Delete(params["address"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeGetNtBuildNumber", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeGetNtBuildNumber(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/GetBase", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::GetBase(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetProgress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceSetProgress(params["progress"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiExecuteOnGuiThread", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiExecuteOnGuiThread(params["cbGuiThread"].get<GUICALLBACK>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRSP(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEBP(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR15D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR15D(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR6", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR6(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAddQWidgetTab(params["qWidget"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCSI(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetIF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetIF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuAddSeparator", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuAddSeparator(params["hMenu"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::Get(params["flag"].get<FlagEnum>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetECX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetECX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeUserDirectory", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeUserDirectory(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetImports", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetImports(params["mod"].get<const ModuleInfo *>(), params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/DeleteHardwareBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::DeleteHardwareBreakpoint(params["address"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptSetTitle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiScriptSetTitle(params["title"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR13", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR13(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMenuEntryClicked", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgMenuEntryClicked(params["hEntry"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptError", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiScriptError(params["line"].get<int>(), params["message"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/AddNode", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::AddNode(params["node"].get<const BridgeCFNode &>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemWrite", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgMemWrite(params["va"].get<duint>(), params["src"].get<const void *>(), params["size"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetXrefCountAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetXrefCountAt(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptAbort", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptAbort(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR8D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR8D(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgLoopDel", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgLoopDel(params["depth"].get<int>(), params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceInitialize", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceInitialize(params["name"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSymbolEnumRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSymbolEnumRange(params["start"].get<duint>(), params["end"].get<duint>(), params["symbolMask"].get<unsigned int>(), params["cbSymbolEnum"].get<CBSYMBOLENUM>(), params["user"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEDX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR12W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR12W(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Function::Delete(params["address"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgExit", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgExit(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBL(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRDI(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/Size", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::Size(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetOF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetOF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgLoopGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgLoopGet(params["depth"].get<int>(), params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetSP(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiScriptClear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR11", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR11(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiCloseQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiCloseQWidgetTab(params["qWidget"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptUnload", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptUnload(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSelectInMemoryMap", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSelectInMemoryMap(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/__debugbreak", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = __debugbreak(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR13B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR13B(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR11B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR11B(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR10D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR10D(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/StepIn", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::StepIn(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR3", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR3(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/IsValidPtr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::IsValidPtr(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetEncodeTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetEncodeTypeAt(params["addr"].get<duint>(), params["size"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetDisassembly", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGetDisassembly(params["addr"].get<duint>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddFavouriteCommand", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAddFavouriteCommand(params["name"].get<const char *>(), params["shortcut"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRDI(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR9B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR9B(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBH(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/SearchAndReplace", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::SearchAndReplace(params["data"].get<unsigned char *>(), params["datasize"].get<duint>(), params["searchpattern"].get<const char *>(), params["replacepattern"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAutoCompleteClearAll", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAutoCompleteClearAll(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateArgumentWidget", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateArgumentWidget(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Function::GetList(params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR2", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR2(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGotoTrace", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGotoTrace(params["index"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetDebugStateFast", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSetDebugStateFast(params["state"].get<DBGSTATE>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCIP(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR13D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR13D(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEBX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetIcon", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuSetIcon(params["hMenu"].get<int>(), params["icon"].get<const ICONDATA *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/WriteMem", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::WriteMem(params["start"].get<duint>(), params["size"].get<duint>(), params["pattern"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptRun", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptRun(params["destline"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgModBaseFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgModBaseFromName(params["name"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Add", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Add(params["info"].get<const ArgumentInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/Read", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::Read(params["addr"].get<duint>(), params["data"].get<void *>(), params["size"].get<duint>(), params["sizeRead"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCDI(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCBP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAutoCompleteAddCmd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAutoCompleteAddCmd(params["cmd"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCX(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetDF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetDF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadDword", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadDword(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRBP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateTraceBrowser", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateTraceBrowser(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetEncodeSizeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetEncodeSizeAt(params["addr"].get<duint>(), params["codesize"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetCellContent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceSetCellContent(params["row"].get<int>(), params["col"].get<int>(), params["str"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowThreads", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiShowThreads(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetAutoBookmarkAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSetAutoBookmarkAt(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuClear(params["hMenu"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddLogMessageHtml", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAddLogMessageHtml(params["msg"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetESP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetESP(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiRepaintTableView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiRepaintTableView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDL(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/CloseQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::CloseQWidgetTab(params["qWidget"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiScriptAdd(params["count"].get<int>(), params["lines"].get<const char **>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/ResolveLabel", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Misc::ResolveLabel(params["label"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateThreadView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateThreadView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearCommentRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgClearCommentRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeIsProcessElevated", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeIsProcessElevated(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR10B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR10B(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/Stop", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::Stop(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR9D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR9D(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCCX(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLoadGraph", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiLoadGraph(params["graph"].get<BridgeCFGraphList *>(), params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBPL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBPL(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/Write", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::Write(params["data"].get<unsigned char *>(), params["datasize"].get<duint>(), params["pattern"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGraphAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGraphAt(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::GetInfo(params["addr"].get<duint>(), params["info"].get<BookmarkInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Add", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Function::Add(params["info"].get<const FunctionInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionFromAddr(params["addr"].get<duint>(), params["number"].get<int>(), params["section"].get<ModuleSectionInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDIL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDIL(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::GetList(params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR15W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR15W(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetProcessId", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetProcessId(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDL(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRDX(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDumpAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiDumpAt(params["va"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCFLAGS", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCFLAGS(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiTypeClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiTypeClear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetSIL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetSIL(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctionGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgFunctionGet(params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Function::Clear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_symbol.h/GetList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Symbol::GetList(params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolSetProgress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSymbolSetProgress(params["percent"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR8B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR8B(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetTebAddress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetTebAddress(params["ThreadId"].get<DWORD>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddFavouriteTool", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAddFavouriteTool(params["name"].get<const char *>(), params["description"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiStackDumpAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiStackDumpAt(params["addr"].get<duint>(), params["csp"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionFromName(params["name"].get<const char *>(), params["number"].get<int>(), params["section"].get<ModuleSectionInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRSI(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionCountFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionCountFromName(params["name"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptStep", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptStep(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetCommentAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetCommentAt(params["addr"].get<duint>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDIL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDIL(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgXrefAddMulti", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgXrefAddMulti(params["edges"].get<const XREF_EDGE *>(), params["count"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEIP(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::Delete(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetSymbolInfoAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetSymbolInfoAt(params["addr"].get<duint>(), params["info"].get<SYMBOLINFO *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowTrace", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiShowTrace(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCL(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiTranslateText", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiTranslateText(params["Source"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCBX(params["value"].get<duint>()};
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

    server.Post("/_scriptapi_memory.h/RemoteFree", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::RemoteFree(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDisplayWarning", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiDisplayWarning(params["title"].get<const char *>(), params["text"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/DisableBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::DisableBreakpoint(params["address"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUnregisterScriptLanguage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUnregisterScriptLanguage(params["id"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetSI(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR6", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR6(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRCX(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::Set(params["info"].get<const LabelInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadPtr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadPtr(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetEntryIcon", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuSetEntryIcon(params["hEntry"].get<int>(), params["icon"].get<const ICONDATA *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/FindMem", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::FindMem(params["start"].get<duint>(), params["size"].get<duint>(), params["pattern"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/InputLine", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::InputLine(params["title"].get<const char *>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetCommentAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSetCommentAt(params["addr"].get<duint>(), params["text"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearAutoFunctionRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgClearAutoFunctionRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRSP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::Clear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR9W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR9W(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgXrefAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgXrefAdd(params["addr"].get<duint>(), params["from"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR11", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR11(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/PathFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::PathFromAddr(params["addr"].get<duint>(), params["path"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSearchGetCellContent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceSearchGetCellContent(params["row"].get<int>(), params["col"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/EntryFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::EntryFromName(params["name"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetThreadList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetThreadList(params["list"].get<THREADLIST *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetThreadId", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetThreadId(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadQword", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadQword(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/DeleteBreakpoint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::DeleteBreakpoint(params["address"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolUpdateModuleList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSymbolUpdateModuleList(params["count"].get<int>(), params["modules"].get<SYMBOLMODULEINFO *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR7", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR7(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Get(params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>(), params["instructionCount"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLogRedirect", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiLogRedirect(params["filename"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_stack.h/Push", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Stack::Push(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR15W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR15W(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR0", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR0(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiEnableLog", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiEnableLog(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeInit", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeInit(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetPebAddress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetPebAddress(params["ProcessId"].get<DWORD>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCSP(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemIsValidReadPtr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgMemIsValidReadPtr(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR13", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR13(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateWindowTitle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateWindowTitle(params["filename"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/IsTemporary", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::IsTemporary(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR11W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR11W(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiRegisterScriptLanguage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiRegisterScriptLanguage(params["info"].get<SCRIPTTYPEINFO *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/Run", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::Run(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptGetBranchInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptGetBranchInfo(params["line"].get<int>(), params["info"].get<SCRIPTBRANCH *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/SelectionGetEnd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::SelectionGetEnd(params["window"].get<Window>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowQWidgetTab", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiShowQWidgetTab(params["qWidget"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEIP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemFindBaseAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgMemFindBaseAddr(params["addr"].get<duint>(), params["size"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLogSave", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiLogSave(params["filename"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/SetProtect", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::SetProtect(params["addr"].get<duint>(), params["protect"].get<unsigned int>(), params["size"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgArgumentOverlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgArgumentOverlaps(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDR1", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR1(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR9", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR9(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCSP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiInvalidateSymbolSource", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiInvalidateSymbolSource(params["base"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetEntryName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuSetEntryName(params["hEntry"].get<int>(), params["name"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/PathFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::PathFromName(params["name"].get<const char *>(), params["path"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Function::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRIP(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsBpDisabled", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgIsBpDisabled(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR14", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR14(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgLoopOverlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgLoopOverlaps(params["depth"].get<int>(), params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetAutoCommentAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSetAutoCommentAt(params["addr"].get<duint>(), params["text"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetAH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetAH(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleSectionList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleSectionList(params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR8W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR8W(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDH(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/Wait", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::Wait(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR9W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR9W(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceAddCommand", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceAddCommand(params["title"].get<const char *>(), params["command"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateAllViews", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateAllViews(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCDI(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Function::Get(params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>(), params["instructionCount"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetStringAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetStringAt(params["addr"].get<duint>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetAF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetAF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetMainThreadId", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGetMainThreadId(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgCmdExecDirect", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgCmdExecDirect(params["cmd"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiOpenTraceFile", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiOpenTraceFile(params["fileName"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgInit", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgInit(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeIsARM64Emulated", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeIsARM64Emulated(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModulePath", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModulePath(params["path"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetTF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetTF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCAX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgValToString", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgValToString(params["s"].get<const char *>(), params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemRead", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgMemRead(params["va"].get<duint>(), params["dest"].get<void *>(), params["size"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingSetUint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeSettingSetUint(params["section"].get<const char *>(), params["key"].get<const char *>(), params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateTimeWastedCounter", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateTimeWastedCounter(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearLabelRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgClearLabelRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetAutoLabelAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSetAutoLabelAt(params["addr"].get<duint>(), params["text"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetEntryHotkey", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuSetEntryHotkey(params["hEntry"].get<int>(), params["hack"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/Set", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Comment::Set(params["info"].get<const CommentInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSelChanged", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSelChanged(params["hWindow"].get<int>(), params["VA"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolLogClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSymbolLogClear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR10", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR10(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/MessageYesNo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::MessageYesNo(params["message"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Argument::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>(), params["deleteManual"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSearchGetRowCount", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceSearchGetRowCount(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgXrefGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgXrefGet(params["addr"].get<duint>(), params["info"].get<XREF_INFO *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateMemoryView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateMemoryView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/InputValue", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::InputValue(params["title"].get<const char *>(), params["value"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::GetInfo(params["addr"].get<duint>(), params["info"].get<LabelInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_stack.h/Pop", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Stack::Pop(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_function.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Function::GetInfo(params["addr"].get<duint>(), params["info"].get<FunctionInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR10W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR10W(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEAX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetCF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetCF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCSI(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetThreadHandle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetThreadHandle(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolLogAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSymbolLogAdd(params["message"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetDF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetDF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBPL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBPL(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WriteByte", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WriteByte(params["addr"].get<duint>(), params["data"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR13D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR13D(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDisableLog", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiDisableLog(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetDebugState", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSetDebugState(params["state"].get<DBGSTATE>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgAnalyzeFunction", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgAnalyzeFunction(params["entry"].get<duint>(), params["graph"].get<BridgeCFGraphList *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR14B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR14B(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WritePtr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WritePtr(params["addr"].get<duint>(), params["data"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/Refresh", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::Refresh(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptMessage", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiScriptMessage(params["message"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiProcessEvents", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiProcessEvents(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiLogClear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiLogClear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSymbolRefreshCurrent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSymbolRefreshCurrent(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/EntryFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::EntryFromAddr(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/Free", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::Free(params["graphList"].get<const BridgeCFGraphList *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetEncodeType", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSetEncodeType(params["addr"].get<duint>(), params["size"].get<duint>(), params["Type"].get<ENCODETYPE>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetEBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEBP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR10", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR10(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceAddColumn", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceAddColumn(params["width"].get<int>(), params["title"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetFavouriteToolShortcut", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSetFavouriteToolShortcut(params["name"].get<const char *>(), params["shortcut"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WriteWord", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WriteWord(params["addr"].get<duint>(), params["data"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR14W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR14W(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/ToNodeList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFNode::ToNodeList(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetCF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetCF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetGlobalNotes", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGetGlobalNotes(params["text"].get<char **>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetGlobalNotes", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSetGlobalNotes(params["text"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSetLastException", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSetLastException(params["exception"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateTypeWidget", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateTypeWidget(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetXrefTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetXrefTypeAt(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/WriteDword", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WriteDword(params["addr"].get<duint>(), params["data"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleBase", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleBase(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadWord", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadWord(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceReloadData", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceReloadData(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR11W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR11W(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleSectionCount", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleSectionCount(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionCountFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionCountFromAddr(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR15", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR15(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetSIL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetSIL(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDI(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetAX(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEDI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEDI(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDX(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetPF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetPF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetEncodeTypeBuffer", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetEncodeTypeBuffer(params["addr"].get<duint>(), params["size"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingGetUint", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeSettingGetUint(params["section"].get<const char *>(), params["key"].get<const char *>(), params["value"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsJumpGoingToExecute", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgIsJumpGoingToExecute(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/ReadByte", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadByte(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_assembler.h/AssembleMem", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Assembler::AssembleMem(params["addr"].get<duint>(), params["instruction"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetAF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetAF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR9B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR9B(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetESP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetESP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR10B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR10B(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::Clear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiShowReferences", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiShowReferences(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR14W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR14W(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetESI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetESI(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetAH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetAH(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR8W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR8W(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRAX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetArgTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetArgTypeAt(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiCloseApplication", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiCloseApplication(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuAdd(params["hMenu"].get<int>(), params["title"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetSF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetSF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SizeFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::SizeFromAddr(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRIP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRIP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetBranchDestination", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetBranchDestination(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsRunning", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgIsRunning(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetActiveView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGetActiveView(params["activeView"].get<ACTIVEVIEW *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/GetProtect", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::GetProtect(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/Clear", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Comment::Clear(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR13W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR13W(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_stack.h/Peek", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Stack::Peek(params["offset"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptGetLineType", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptGetLineType(params["line"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDH(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSettingsUpdated", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSettingsUpdated(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR11D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR11D(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetWatchList", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetWatchList(params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctionOverlaps", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgFunctionOverlaps(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearAutoBookmarkRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgClearAutoBookmarkRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetAL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetAL(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/BaseFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::BaseFromAddr(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingSet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeSettingSet(params["section"].get<const char *>(), params["key"].get<const char *>(), params["value"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/Message", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::Message(params["message"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/SearchAndReplaceMem", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::SearchAndReplaceMem(params["start"].get<duint>(), params["size"].get<duint>(), params["searchpattern"].get<const char *>(), params["replacepattern"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/SelectionSet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::SelectionSet(params["window"].get<Window>(), params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeSettingGet(params["section"].get<const char *>(), params["key"].get<const char *>(), params["value"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR10D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR10D(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgReleaseEncodeTypeBuffer", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgReleaseEncodeTypeBuffer(params["buffer"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR8", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR8(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_assembler.h/AssembleMemEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Assembler::AssembleMemEx(params["addr"].get<duint>(), params["instruction"].get<const char *>(), params["size"].get<int *>(), params["error"].get<char *>(), params["fillnop"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgMemGetPageSize", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgMemGetPageSize(params["base"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEBX(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRSI(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuRemove", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuRemove(params["hEntryMenu"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRCX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/ParseExpression", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Misc::ParseExpression(params["expression"].get<const char *>(), params["value"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR13B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR13B(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleSize", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleSize(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCDX(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetModuleAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetModuleAt(params["addr"].get<duint>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR12", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR12(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR9", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR9(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR14B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR14B(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleInfo(params["info"].get<ModuleInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgDelEncodeTypeSegment", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgDelEncodeTypeSegment(params["start"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetLineWindow", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGetLineWindow(params["title"].get<const char *>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_misc.h/RemoteGetProcAddress", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Misc::RemoteGetProcAddress(params["module"].get<const char *>(), params["api"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetSPL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetSPL(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR15B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR15B(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetDebuggeeNotes", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGetDebuggeeNotes(params["text"].get<char **>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/NameFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::NameFromAddr(params["addr"].get<duint>(), params["name"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceGetRowCount", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceGetRowCount(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeFree", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeFree(params["ptr"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetSPL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetSPL(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCBP(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgScriptLoad", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgScriptLoad(params["filename"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR8", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR8(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/GetOF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetOF(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSymbolEnum", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSymbolEnum(params["base"].get<duint>(), params["cbSymbolEnum"].get<CBSYMBOLENUM>(), params["user"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetFunctionTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetFunctionTypeAt(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_argument.h/GetInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Argument::GetInfo(params["addr"].get<duint>(), params["info"].get<ArgumentInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_pattern.h/Find", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::Find(params["data"].get<unsigned char *>(), params["datasize"].get<duint>(), params["pattern"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetAL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetAL(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetProcessHandle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetProcessHandle(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateGraphView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateGraphView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSymbolEnumFromCache", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSymbolEnumFromCache(params["base"].get<duint>(), params["cbSymbolEnum"].get<CBSYMBOLENUM>(), params["user"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetWindowHandle", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGetWindowHandle(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAutoCompleteDelCmd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAutoCompleteDelCmd(params["cmd"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_bookmark.h/DeleteRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiSelectionSet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiSelectionSet(params["hWindow"].get<GUISELECTIONTYPE>(), params["selection"].get<const SELECTIONDATA *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR11B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR11B(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetSI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetSI(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuAddEntry", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuAddEntry(params["hMenu"].get<int>(), params["title"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgArgumentGet", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgArgumentGet(params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/StepOut", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::StepOut(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBX(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeSettingFlush", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeSettingFlush(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR12W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR12W(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR13W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR13W(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiExecuteOnGuiThreadEx", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiExecuteOnGuiThreadEx(params["cbGuiThread"].get<GUICALLBACKEX>(), params["userdata"].get<void *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionListFromAddr", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionListFromAddr(params["addr"].get<duint>(), params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgClearAutoCommentRange", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgClearAutoCommentRange(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetMainModuleEntry", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleEntry(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR15", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR15(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetBP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBP(params["value"].get<unsigned short>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdatePatches", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdatePatches(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiAddInfoLine", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiAddInfoLine(params["infoLine"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetRBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRBX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiReferenceSetSearchStartCol", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiReferenceSetSearchStartCol(params["col"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetEDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEDX(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeStart", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeStart(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetESI", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetESI(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR1", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR1(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_comment.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Comment::Delete(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_gui.h/SelectionGetStart", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Gui::SelectionGetStart(params["window"].get<Window>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgValFromString", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgValFromString(params["s"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgegraph.h/AddParent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::AddParent(params["child"].get<int>(), params["parent"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeGetDbgVersion", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeGetDbgVersion(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::Get(params["reg"].get<RegisterEnum>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateWatchView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateWatchView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR15B", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR15B(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgXrefDelAll", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgXrefDelAll(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCH(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetSymbolInfo", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetSymbolInfo(params["symbolptr"].get<const SYMBOLPTR *>(), params["info"].get<SYMBOLINFO *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgDisasmAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgDisasmAt(params["addr"].get<duint>(), params["instr"].get<DISASM_INSTR *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiGetCurrentGraph", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiGetCurrentGraph(params["graphList"].get<BridgeCFGraphList *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsRunLocked", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgIsRunLocked(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_debug.h/StepOver", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Debug::StepOver(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCBX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCBX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetR14D", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR14D(params["value"].get<unsigned int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgCmdExec", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgCmdExec(params["cmd"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiTypeAddNode", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiTypeAddNode(params["parent"].get<void *>(), params["Type"].get<const TYPEDESCRIPTOR *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgWinEvent", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgWinEvent(params["message"].get<MSG *>(), params["result"].get<long *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiDisasmAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiDisasmAt(params["addr"].get<duint>(), params["cip"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCL", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCL(params["value"].get<unsigned char>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetCFLAGS", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCFLAGS(params["value"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiIsLogEnabled", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiIsLogEnabled(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetBH", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBH(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateBreakpointsView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateBreakpointsView(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/InfoFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::InfoFromName(params["name"].get<const char *>(), params["info"].get<ModuleInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgArgumentDel", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgArgumentDel(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateSEHChain", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateSEHChain(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCCX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCCX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetDR0", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR0(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgSetAutoFunctionAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgSetAutoFunctionAt(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiMenuSetName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiMenuSetName(params["hMenu"].get<int>(), params["name"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetLoopTypeAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetLoopTypeAt(params["addr"].get<duint>(), params["depth"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiFocusView", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiFocusView(params["hWindow"].get<int>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_memory.h/RemoteAlloc", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Memory::RemoteAlloc(params["addr"].get<duint>(), params["size"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_flag.h/SetZF", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetZF(params["value"].get<bool>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/Delete", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::Delete(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/SetRAX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRAX(params["value"].get<unsigned long long>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiIsUpdateDisabled", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiIsUpdateDisabled(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctions", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgFunctions(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetSP", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetSP(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/BridgeAlloc", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = BridgeAlloc(params["size"].get<size_t>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_label.h/Get", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Label::Get(params["addr"].get<duint>(), params["text"].get<char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgGetBookmarkAt", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgGetBookmarkAt(params["addr"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiUpdateCallStack", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiUpdateCallStack(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/SectionListFromName", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionListFromName(params["name"].get<const char *>(), params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR10W", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR10W(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgFunctionAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgFunctionAdd(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/GuiScriptSetInfoLine", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = GuiScriptSetInfoLine(params["line"].get<int>(), params["info"].get<const char *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetCDX", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCDX(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgIsDebugging", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgIsDebugging(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgLoopAdd", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgLoopAdd(params["start"].get<duint>(), params["end"].get<duint>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_register.h/GetR12", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR12(};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/_scriptapi_module.h/GetExports", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetExports(params["mod"].get<const ModuleInfo *>(), params["list"].get<ListInfo *>()};
res.set_content(json(resp).dump(), "application/json");
       } catch (const std::exception &e) { res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json"); }
   });

    server.Post("/bridgemain.h/DbgWinEventGlobal", [](const Request &req, Response &res) {
       try {
auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
           json params;
           for (const auto &param: arg) { params[param.name] = param.value; }
ApiResponse resp{.success = true, .type = "", .result = DbgWinEventGlobal(params["message"].get<MSG *>()};
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
