//
// Created by Admin on 28/05/2025.
//
#ifndef MCPX64DBG_ECHO_GEN_H
#define MCPX64DBG_ECHO_GEN_H

#include "include/api.h"

#include "pluginsdk/_plugins.h"
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
void         dispatch();
DWORD WINAPI HttpServerThread(LPVOID lpParam);
bool         cbEnableHttpServer(int argc, char *argv[]);
bool         cbSetHttpPort(int argc, char *argv[]);

struct Param {
    std::string name;
    std::string type;
    json        value;
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(Param, name, type, value)
};

/*
 {
   json.Marshaler --> adl_serializer
   json.Marshal   --> to_json
   json.Unmarshal --> from_json
 }
 */
/*
namespace nlohmann {
    template<>
    struct adl_serializer<api::Function::FunctionInfo> {
        static void to_json(json &j, const api::Function::FunctionInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rvaStart", self.rvaStart},
                    {"rvaEnd", self.rvaEnd},
                    {"manual", self.manual},
                    {"instructioncount", self.instructioncount},
            };
        }
    };

    template<>
    struct adl_serializer<api::Bookmark::BookmarkInfo> {
        static void to_json(json &j, const api::Bookmark::BookmarkInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rva", self.rva},
                    {"manual", self.manual},
            };
        }
    };

    template<>
    struct adl_serializer<api::Argument::ArgumentInfo> {
        static void to_json(json &j, const api::Argument::ArgumentInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rvaStart", self.rvaStart},
                    {"rvaEnd", self.rvaEnd},
                    {"manual", self.manual},
                    {"instructioncount", self.instructioncount},
            };
        }
    };

    template<>
    struct adl_serializer<api::Comment::CommentInfo> {
        static void to_json(json &j, const api::Comment::CommentInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rva", self.rva},
                    {"text", self.text},
                    {"manual", self.manual},
            };
        }
    };

    template<>
    struct adl_serializer<api::Module::ModuleInfo> {
        static void to_json(json &j, const api::Module::ModuleInfo &self) {
            j = {

                    {"base", self.base},
                    {"size", self.size},
                    {"entry", self.entry},
                    {"sectionCount", self.sectionCount},
                    {"name", self.name},
                    {"path", self.path},
            };
        }
    };

    template<>
    struct adl_serializer<api::Module::ModuleSectionInfo> {
        static void to_json(json &j, const api::Module::ModuleSectionInfo &self) {
            j = {

                    {"addr", self.addr},
                    {"size", self.size},
                    {"name", self.name},
            };
        }
    };

    template<>
    struct adl_serializer<api::Module::ModuleExport> {
        static void to_json(json &j, const api::Module::ModuleExport &self) {
            j = {

                    {"ordinal", self.ordinal},
                    {"rva", self.rva},
                    {"va", self.va},
                    {"forwarded", self.forwarded},
                    {"forwardName", self.forwardName},
                    {"name", self.name},
                    {"undecoratedName", self.undecoratedName},
            };
        }
    };

    template<>
    struct adl_serializer<api::Module::ModuleImport> {
        static void to_json(json &j, const api::Module::ModuleImport &self) {
            j = {

                    {"iatRva", self.iatRva},
                    {"iatVa", self.iatVa},
                    {"ordinal", self.ordinal},
                    {"name", self.name},
                    {"undecoratedName", self.undecoratedName},
            };
        }
    };

    template<>
    struct adl_serializer<api::Label::LabelInfo> {
        static void to_json(json &j, const api::Label::LabelInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rva", self.rva},
                    {"text", self.text},
                    {"manual", self.manual},
            };
        }
    };

}// namespace nlohmann
*/

void dispatch() {
    server.Post("/_scriptapi_argument.h/Get", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Argument::Get(params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>(), params["instructionCount"].get<duint *>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_argument.h/Overlaps", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Argument::Overlaps(params["start"].get<duint>(), params["end"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_argument.h/Delete", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Argument::Delete(params["address"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_argument.h/DeleteRange", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Argument::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>(), params["deleteManual"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_argument.h/Clear", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Argument::Clear();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_argument.h/GetList", [](const Request &req, Response &res) {
        try {
            auto resp = api::Argument::GetList();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_assembler.h/Assemble", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Assembler::Assemble(params["addr"].get<duint>(), params["instruction"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_assembler.h/AssembleEx", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Assembler::AssembleEx(params["addr"].get<duint>(), params["instruction"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_assembler.h/AssembleMem", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Assembler::AssembleMem(params["addr"].get<duint>(), params["instruction"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_assembler.h/AssembleMemEx", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Assembler::AssembleMemEx(params["addr"].get<duint>(), params["instruction"].get<std::string>(), params["fillnop"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_bookmark.h/Get", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Bookmark::Get(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_bookmark.h/Delete", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Bookmark::Delete(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_bookmark.h/DeleteRange", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Bookmark::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_bookmark.h/Clear", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Bookmark::Clear();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_bookmark.h/GetList", [](const Request &req, Response &res) {
        try {
            auto resp = api::Bookmark::GetList();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_comment.h/Get", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Comment::Get(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_comment.h/Delete", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Comment::Delete(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_comment.h/DeleteRange", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Comment::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_comment.h/Clear", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Comment::Clear();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_comment.h/GetList", [](const Request &req, Response &res) {
        try {
            auto resp = api::Comment::GetList();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/Wait", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Debug::Wait();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/Run", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Debug::Run();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/Pause", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Debug::Pause();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/Stop", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Debug::Stop();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/StepIn", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Debug::StepIn();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/StepOver", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Debug::StepOver();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/StepOut", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Debug::StepOut();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/SetBreakpoint", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Debug::SetBreakpoint(params["address"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/DeleteBreakpoint", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Debug::DeleteBreakpoint(params["address"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/DisableBreakpoint", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Debug::DisableBreakpoint(params["address"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/DeleteHardwareBreakpoint", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Debug::DeleteHardwareBreakpoint(params["address"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetZF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetZF();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetZF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetZF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetOF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetOF();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetOF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetOF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetCF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetCF();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetCF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetCF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetPF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetPF();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetPF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetPF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetSF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetSF();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetSF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetSF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetTF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetTF();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetTF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetTF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetAF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetAF();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetAF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetAF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetDF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetDF();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetDF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetDF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/GetIF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::GetIF();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/SetIF", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Flag::SetIF(params["value"].get<bool>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/Get", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Function::Get(params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>(), params["instructionCount"].get<duint *>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/Overlaps", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Function::Overlaps(params["start"].get<duint>(), params["end"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/Delete", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Function::Delete(params["address"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/DeleteRange", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Function::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/Clear", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Function::Clear();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/GetList", [](const Request &req, Response &res) {
        try {
            auto resp = api::Function::GetList();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/Refresh", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Gui::Refresh();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/FromString", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Label::FromString(params["label"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/Get", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Label::Get(params["addr"].get<duint>(), params["text"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/IsTemporary", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Label::IsTemporary(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/Delete", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Label::Delete(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/DeleteRange", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Label::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>());

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/Clear", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Label::Clear();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/GetList", [](const Request &req, Response &res) {
        try {
            auto resp = api::Label::GetList();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/Read", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::Read(params["addr"].get<duint>(), params["size"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/Write", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::Write(params["addr"].get<duint>(), params["data"].get<std::vector<unsigned char>>(), params["size"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/IsValidPtr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::IsValidPtr(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/RemoteAlloc", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Memory::RemoteAlloc(params["addr"].get<duint>(), params["size"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/RemoteFree", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::RemoteFree(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/GetProtect", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Memory::GetProtect(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/SetProtect", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::SetProtect(params["addr"].get<duint>(), params["protect"].get<unsigned int>(), params["size"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/GetBase", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Memory::GetBase(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/GetSize", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Memory::GetSize(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/ReadByte", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Memory::ReadByte(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/WriteByte", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::WriteByte(params["addr"].get<duint>(), params["data"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/ReadWord", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Memory::ReadWord(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/WriteWord", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::WriteWord(params["addr"].get<duint>(), params["data"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/ReadDword", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Memory::ReadDword(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/WriteDword", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::WriteDword(params["addr"].get<duint>(), params["data"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/ReadQword", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Memory::ReadQword(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/WriteQword", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::WriteQword(params["addr"].get<duint>(), params["data"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/ReadPtr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Memory::ReadPtr(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_memory.h/WritePtr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Memory::WritePtr(params["addr"].get<duint>(), params["data"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/BaseFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::BaseFromAddr(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/BaseFromName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::BaseFromName(params["name"].get<std::string>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/SizeFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::SizeFromAddr(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/SizeFromName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::SizeFromName(params["name"].get<std::string>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/NameFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Module::NameFromAddr(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/PathFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Module::PathFromAddr(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/PathFromName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Module::PathFromName(params["name"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/EntryFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::EntryFromAddr(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/EntryFromName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::EntryFromName(params["name"].get<std::string>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/SectionCountFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "int", .result = api::Module::SectionCountFromAddr(params["addr"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/SectionCountFromName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "int", .result = api::Module::SectionCountFromName(params["name"].get<std::string>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/SectionListFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Module::SectionListFromAddr(params["addr"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/SectionListFromName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Module::SectionListFromName(params["name"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModuleBase", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::GetMainModuleBase()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModuleSize", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::GetMainModuleSize()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModuleEntry", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Module::GetMainModuleEntry()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModuleSectionCount", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "int", .result = api::Module::GetMainModuleSectionCount()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModuleName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Module::GetMainModuleName();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModulePath", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Module::GetMainModulePath();
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModuleSectionList", [](const Request &req, Response &res) {
        try {
            auto resp = api::Module::GetMainModuleSectionList();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetList", [](const Request &req, Response &res) {
        try {
            auto resp = api::Module::GetList();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_pattern.h/Find", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Pattern::Find(params["datasize"].get<duint>(), params["pattern"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_pattern.h/FindMem", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Pattern::FindMem(params["start"].get<duint>(), params["size"].get<duint>(), params["pattern"].get<std::string>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_pattern.h/Write", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Pattern::Write(params["datasize"].get<duint>(), params["pattern"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_pattern.h/WriteMem", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Pattern::WriteMem(params["start"].get<duint>(), params["size"].get<duint>(), params["pattern"].get<std::string>());

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_pattern.h/SearchAndReplace", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp = api::Pattern::SearchAndReplace(params["datasize"].get<duint>(), params["searchpattern"].get<std::string>(), params["replacepattern"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_pattern.h/SearchAndReplaceMem", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Pattern::SearchAndReplaceMem(params["start"].get<duint>(), params["size"].get<duint>(), params["searchpattern"].get<std::string>(), params["replacepattern"].get<std::string>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/Get", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            auto resp= api::Register::Get();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/Set", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::Set(params["reg"].get<api::Register::RegisterEnum>(), params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/Size", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "int", .result = api::Register::Size()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDR0", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetDR0()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDR0", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDR0(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDR1", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetDR1()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDR1", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDR1(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDR2", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetDR2()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDR2", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDR2(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDR3", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetDR3()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDR3", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDR3(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDR6", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetDR6()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDR6", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDR6(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDR7", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetDR7()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDR7", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDR7(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetEAX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetEAX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetEAX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetEAX(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetAX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetAX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetAX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetAX(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetAH", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetAH()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetAH", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetAH(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetAL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetAL()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetAL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetAL(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetEBX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetEBX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetEBX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetEBX(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetBX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetBX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetBX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetBX(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetBH", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetBH()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetBH", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetBH(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetBL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetBL()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetBL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetBL(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetECX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetECX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetECX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetECX(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetCX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCX(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCH", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetCH()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCH", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCH(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetCL()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCL(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetEDX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetEDX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetEDX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetEDX(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetDX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDX(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDH", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetDH()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDH", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDH(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetDL()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDL(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetEDI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetEDI()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetEDI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetEDI(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetDI()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDI(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetESI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetESI()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetESI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetESI(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetSI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetSI()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetSI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetSI(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetEBP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetEBP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetEBP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetEBP(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetBP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetBP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetBP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetBP(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetESP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetESP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetESP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetESP(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetSP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetSP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetSP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetSP(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetEIP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetEIP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetEIP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetEIP(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRAX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRAX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRAX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRAX(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRBX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRBX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRBX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRBX(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRCX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRCX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRCX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRCX(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRDX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRDX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRDX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRDX(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRSI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRSI()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRSI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRSI(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetSIL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetSIL()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetSIL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetSIL(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRDI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRDI()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRDI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRDI(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetDIL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetDIL()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetDIL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetDIL(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRBP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRBP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRBP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRBP(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetBPL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetBPL()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetBPL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetBPL(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRSP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRSP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRSP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRSP(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetSPL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetSPL()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetSPL", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetSPL(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetRIP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetRIP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetRIP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetRIP(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR8", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetR8()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR8", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR8(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR8D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetR8D()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR8D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR8D(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR8W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetR8W()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR8W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR8W(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR8B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetR8B()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR8B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR8B(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR9", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetR9()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR9", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR9(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR9D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetR9D()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR9D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR9D(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR9W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetR9W()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR9W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR9W(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR9B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetR9B()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR9B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR9B(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR10", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetR10()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR10", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR10(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR10D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetR10D()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR10D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR10D(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR10W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetR10W()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR10W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR10W(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR10B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetR10B()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR10B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR10B(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR11", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetR11()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR11", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR11(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR11D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetR11D()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR11D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR11D(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR11W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetR11W()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR11W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR11W(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR11B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetR11B()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR11B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR11B(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR12", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetR12()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR12", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR12(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR12D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetR12D()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR12D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR12D(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR12W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetR12W()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR12W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR12W(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR12B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetR12B()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR12B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR12B(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR13", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetR13()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR13", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR13(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR13D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetR13D()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR13D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR13D(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR13W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetR13W()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR13W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR13W(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR13B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetR13B()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR13B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR13B(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR14", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetR14()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR14", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR14(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR14D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetR14D()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR14D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR14D(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR14W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetR14W()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR14W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR14W(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR14B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetR14B()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR14B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR14B(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR15", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned long long", .result = api::Register::GetR15()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR15", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR15(params["value"].get<unsigned long long>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR15D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned int", .result = api::Register::GetR15D()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR15D", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR15D(params["value"].get<unsigned int>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR15W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned short", .result = api::Register::GetR15W()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR15W", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR15W(params["value"].get<unsigned short>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetR15B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "unsigned char", .result = api::Register::GetR15B()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetR15B", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetR15B(params["value"].get<unsigned char>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCAX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCAX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCAX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCAX(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCBX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCBX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCBX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCBX(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCCX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCCX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCCX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCCX(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCDX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCDX()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCDX", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCDX(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCDI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCDI()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCDI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCDI(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCSI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCSI()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCSI", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCSI(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCBP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCBP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCBP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCBP(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCSP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCSP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCSP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCSP(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCIP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCIP()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCIP", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCIP(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/GetCFLAGS", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Register::GetCFLAGS()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_register.h/SetCFLAGS", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }

            auto resp = api::Register::SetCFLAGS(params["value"].get<duint>());
            res.set_content(json(resp).dump(), "application/json");
            ;
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_stack.h/Pop", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Stack::Pop()};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_stack.h/Push", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Stack::Push(params["value"].get<duint>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_stack.h/Peek", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "duint", .result = api::Stack::Peek(params["offset"].get<int>())};

            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_symbol.h/GetList", [](const Request &req, Response &res) {
        try {
            auto resp = api::Symbol::GetList();
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
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
    if (g_httpServerRunning) {
        stopHttpServer();
    }
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
