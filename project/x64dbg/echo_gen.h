

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

namespace nlohmann {
    template<>
    struct adl_serializer<Script::Argument::ArgumentInfo> {
        static void to_json(json &j, const Script::Argument::ArgumentInfo &self) {
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
    struct adl_serializer<Script::Bookmark::BookmarkInfo> {
        static void to_json(json &j, const Script::Bookmark::BookmarkInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rva", self.rva},
                    {"manual", self.manual},
            };
        }
    };

    template<>
    struct adl_serializer<Script::Comment::CommentInfo> {
        static void to_json(json &j, const Script::Comment::CommentInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rva", self.rva},
                    {"text", self.text},
                    {"manual", self.manual},
            };
        }
    };

    template<>
    struct adl_serializer<Script::Function::FunctionInfo> {
        static void to_json(json &j, const Script::Function::FunctionInfo &self) {
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
    struct adl_serializer<Script::Label::LabelInfo> {
        static void to_json(json &j, const Script::Label::LabelInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rva", self.rva},
                    {"text", self.text},
                    {"manual", self.manual},
            };
        }
    };

    template<>
    struct adl_serializer<Script::Module::ModuleInfo> {
        static void to_json(json &j, const Script::Module::ModuleInfo &self) {
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
    struct adl_serializer<Script::Module::ModuleSectionInfo> {
        static void to_json(json &j, const Script::Module::ModuleSectionInfo &self) {
            j = {

                    {"addr", self.addr},
                    {"size", self.size},
                    {"name", self.name},
            };
        }
    };

    template<>
    struct adl_serializer<Script::Module::ModuleExport> {
        static void to_json(json &j, const Script::Module::ModuleExport &self) {
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
    struct adl_serializer<Script::Module::ModuleImport> {
        static void to_json(json &j, const Script::Module::ModuleImport &self) {
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
    struct adl_serializer<Script::Symbol::SymbolInfo> {
        static void to_json(json &j, const Script::Symbol::SymbolInfo &self) {
            j = {

                    {"mod", self.mod},
                    {"rva", self.rva},
                    {"name", self.name},
                    {"manual", self.manual},
                    {"type", self.type},
            };
        }
    };

    template<>
    struct adl_serializer<BridgeCFNode> {
        static void to_json(json &j, const BridgeCFNode &self) {
            j = {

                    {"parentGraph", self.parentGraph},
                    {"start", self.start},
                    {"end", self.end},
                    {"brtrue", self.brtrue},
                    {"brfalse", self.brfalse},
                    {"icount", self.icount},
                    {"terminal", self.terminal},
                    {"split", self.split},
                    {"indirectcall", self.indirectcall},
                    {"userdata", self.userdata},
                    {"exits", self.exits},
                    {"instrs", self.instrs},
            };
        }
    };

    template<>
    struct adl_serializer<BridgeCFGraph> {
        static void to_json(json &j, const BridgeCFGraph &self) {
            j = {

                    {"entryPoint", self.entryPoint},
                    {"userdata", self.userdata},
                    {"nodes", self.nodes},
                    {"parents", self.parents},
            };
        }
    };

}// namespace nlohmann

void dispatch() {
    server.Post("/_scriptapi_argument.h/Add", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Add(params["info"].get<const ArgumentInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_argument.h/Get", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Get(params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>(), params["instructionCount"].get<duint *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_argument.h/GetInfo", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Argument::GetInfo(params["addr"].get<duint>(), params["info"].get<ArgumentInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Overlaps(params["start"].get<duint>(), params["end"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Delete(params["address"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Argument::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>(), params["deleteManual"].get<bool>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Argument::Clear()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_argument.h/GetList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Argument::GetList(params["list"].get<ListInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Assembler::Assemble(params["addr"].get<duint>(), params["dest"].get<unsigned char *>(), params["size"].get<int *>(), params["instruction"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Assembler::AssembleEx(params["addr"].get<duint>(), params["dest"].get<unsigned char *>(), params["size"].get<int *>(), params["instruction"].get<const char *>(), params["error"].get<char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Assembler::AssembleMem(params["addr"].get<duint>(), params["instruction"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Assembler::AssembleMemEx(params["addr"].get<duint>(), params["instruction"].get<const char *>(), params["size"].get<int *>(), params["error"].get<char *>(), params["fillnop"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_bookmark.h/Set", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::Set(params["info"].get<const BookmarkInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::Get(params["addr"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_bookmark.h/GetInfo", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::GetInfo(params["addr"].get<duint>(), params["info"].get<BookmarkInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::Delete(params["addr"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::Clear()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_bookmark.h/GetList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Bookmark::GetList(params["list"].get<ListInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_comment.h/Set", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Comment::Set(params["info"].get<const CommentInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Comment::Get(params["addr"].get<duint>(), params["text"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_comment.h/GetInfo", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Comment::GetInfo(params["addr"].get<duint>(), params["info"].get<CommentInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Comment::Delete(params["addr"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Comment::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Comment::Clear()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_comment.h/GetList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Comment::GetList(params["list"].get<ListInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::Wait()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::Run()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::Pause()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::Stop()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::StepIn()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::StepOver()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::StepOut()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::SetBreakpoint(params["address"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::DeleteBreakpoint(params["address"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::DisableBreakpoint(params["address"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_debug.h/SetHardwareBreakpoint", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::SetHardwareBreakpoint(params["address"].get<duint>(), params["Type"].get<HardwareType>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Debug::DeleteHardwareBreakpoint(params["address"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/Get", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::Get(params["flag"].get<FlagEnum>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_flag.h/Set", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::Set(params["flag"].get<FlagEnum>(), params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetZF()};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetZF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetOF()};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetOF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetCF()};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetCF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetPF()};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetPF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetSF()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetSF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetTF()};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetTF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetAF()};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetAF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetDF()};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetDF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::GetIF()};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Flag::SetIF(params["value"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/Add", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Function::Add(params["info"].get<const FunctionInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Function::Get(params["addr"].get<duint>(), params["start"].get<duint *>(), params["end"].get<duint *>(), params["instructionCount"].get<duint *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/GetInfo", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Function::GetInfo(params["addr"].get<duint>(), params["info"].get<FunctionInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Function::Overlaps(params["start"].get<duint>(), params["end"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Function::Delete(params["address"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Function::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Function::Clear()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_function.h/GetList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Function::GetList(params["list"].get<ListInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/SelectionGet", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::SelectionGet(params["window"].get<Window>(), params["start"].get<duint *>(), params["end"].get<duint *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/SelectionSet", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::SelectionSet(params["window"].get<Window>(), params["start"].get<duint>(), params["end"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/SelectionGetStart", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::SelectionGetStart(params["window"].get<Window>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/SelectionGetEnd", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::SelectionGetEnd(params["window"].get<Window>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/Message", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::Message(params["message"].get<const char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/MessageYesNo", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::MessageYesNo(params["message"].get<const char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/InputLine", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::InputLine(params["title"].get<const char *>(), params["text"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/InputValue", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::InputValue(params["title"].get<const char *>(), params["value"].get<duint *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::Refresh()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/AddQWidgetTab", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::AddQWidgetTab(params["qWidget"].get<void *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/ShowQWidgetTab", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::ShowQWidgetTab(params["qWidget"].get<void *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_gui.h/CloseQWidgetTab", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Gui::CloseQWidgetTab(params["qWidget"].get<void *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/Set", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::Set(params["info"].get<const LabelInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::FromString(params["label"].get<const char *>(), params["addr"].get<duint *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::Get(params["addr"].get<duint>(), params["text"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::IsTemporary(params["addr"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/GetInfo", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::GetInfo(params["addr"].get<duint>(), params["info"].get<LabelInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::Delete(params["addr"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::DeleteRange(params["start"].get<duint>(), params["end"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::Clear()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_label.h/GetList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Label::GetList(params["list"].get<ListInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::Read(params["addr"].get<duint>(), params["data"].get<void *>(), params["size"].get<duint>(), params["sizeRead"].get<duint *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::Write(params["addr"].get<duint>(), params["data"].get<const void *>(), params["size"].get<duint>(), params["sizeWritten"].get<duint *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::IsValidPtr(params["addr"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::RemoteAlloc(params["addr"].get<duint>(), params["size"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::RemoteFree(params["addr"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::GetProtect(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::SetProtect(params["addr"].get<duint>(), params["protect"].get<unsigned int>(), params["size"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::GetBase(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::GetSize(params["addr"].get<duint>(), params["reserved"].get<bool>(), params["cache"].get<bool>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadByte(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WriteByte(params["addr"].get<duint>(), params["data"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadWord(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WriteWord(params["addr"].get<duint>(), params["data"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadDword(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WriteDword(params["addr"].get<duint>(), params["data"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadQword(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WriteQword(params["addr"].get<duint>(), params["data"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::ReadPtr(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Memory::WritePtr(params["addr"].get<duint>(), params["data"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_misc.h/ParseExpression", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Misc::ParseExpression(params["expression"].get<const char *>(), params["value"].get<duint *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_misc.h/RemoteGetProcAddress", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Misc::RemoteGetProcAddress(params["module"].get<const char *>(), params["api"].get<const char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_misc.h/ResolveLabel", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Misc::ResolveLabel(params["label"].get<const char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_misc.h/Alloc", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Misc::Alloc(params["size"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_misc.h/Free", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Misc::Free(params["ptr"].get<void *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/InfoFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::InfoFromAddr(params["addr"].get<duint>(), params["info"].get<ModuleInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/InfoFromName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::InfoFromName(params["name"].get<const char *>(), params["info"].get<ModuleInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::BaseFromAddr(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::BaseFromName(params["name"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::SizeFromAddr(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::SizeFromName(params["name"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::NameFromAddr(params["addr"].get<duint>(), params["name"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::PathFromAddr(params["addr"].get<duint>(), params["path"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::PathFromName(params["name"].get<const char *>(), params["path"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::EntryFromAddr(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::EntryFromName(params["name"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionCountFromAddr(params["addr"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionCountFromName(params["name"].get<const char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/SectionFromAddr", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionFromAddr(params["addr"].get<duint>(), params["number"].get<int>(), params["section"].get<ModuleSectionInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/SectionFromName", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionFromName(params["name"].get<const char *>(), params["number"].get<int>(), params["section"].get<ModuleSectionInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionListFromAddr(params["addr"].get<duint>(), params["list"].get<ListInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::SectionListFromName(params["name"].get<const char *>(), params["list"].get<ListInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModuleInfo", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleInfo(params["info"].get<ModuleInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleBase()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleSize()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleEntry()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleSectionCount()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleName(params["name"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModulePath(params["path"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetMainModuleSectionList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetMainModuleSectionList(params["list"].get<ListInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetList", [](const Request &req, Response &res) {
        try {
            BridgeList<Script::Module::ModuleInfo>  bridgeList;
            bool                                    ok = Script::Module::GetList(&bridgeList);
            std::vector<Script::Module::ModuleInfo> moduleVector;
            if (ok) {
                BridgeList<Script::Module::ModuleInfo>::ToVector(&bridgeList, moduleVector, true);
            }
            ApiResponse resp{.success = ok, .type = "array", .result = moduleVector};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetExports", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetExports(params["mod"].get<const ModuleInfo *>(), params["list"].get<ListInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_module.h/GetImports", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Module::GetImports(params["mod"].get<const ModuleInfo *>(), params["list"].get<ListInfo *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::Find(params["data"].get<unsigned char *>(), params["datasize"].get<duint>(), params["pattern"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::FindMem(params["start"].get<duint>(), params["size"].get<duint>(), params["pattern"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::Write(params["data"].get<unsigned char *>(), params["datasize"].get<duint>(), params["pattern"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::WriteMem(params["start"].get<duint>(), params["size"].get<duint>(), params["pattern"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::SearchAndReplace(params["data"].get<unsigned char *>(), params["datasize"].get<duint>(), params["searchpattern"].get<const char *>(), params["replacepattern"].get<const char *>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Pattern::SearchAndReplaceMem(params["start"].get<duint>(), params["size"].get<duint>(), params["searchpattern"].get<const char *>(), params["replacepattern"].get<const char *>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::Get(params["reg"].get<RegisterEnum>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::Set(params["reg"].get<RegisterEnum>(), params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::Size()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR0()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR0(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR1()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR1(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR2()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR2(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR3()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR3(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR6()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR6(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDR7()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDR7(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEAX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEAX(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetAX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetAX(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetAH()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetAH(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetAL()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetAL(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEBX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEBX(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBX(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBH()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBH(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBL()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBL(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetECX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetECX(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCX(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCH()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCH(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCL()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCL(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEDX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEDX(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDX(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDH()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDH(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDL()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDL(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEDI()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEDI(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDI()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDI(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetESI()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetESI(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetSI()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetSI(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEBP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEBP(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBP(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetESP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetESP(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetSP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetSP(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetEIP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetEIP(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRAX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRAX(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRBX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRBX(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRCX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRCX(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRDX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRDX(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRSI()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRSI(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetSIL()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetSIL(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRDI()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRDI(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetDIL()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetDIL(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRBP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRBP(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetBPL()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetBPL(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRSP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRSP(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetSPL()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetSPL(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetRIP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetRIP(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR8()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR8(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR8D()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR8D(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR8W()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR8W(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR8B()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR8B(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR9()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR9(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR9D()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR9D(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR9W()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR9W(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR9B()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR9B(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR10()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR10(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR10D()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR10D(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR10W()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR10W(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR10B()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR10B(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR11()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR11(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR11D()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR11D(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR11W()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR11W(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR11B()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR11B(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR12()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR12(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR12D()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR12D(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR12W()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR12W(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR12B()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR12B(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR13()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR13(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR13D()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR13D(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR13W()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR13W(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR13B()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR13B(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR14()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR14(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR14D()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR14D(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR14W()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR14W(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR14B()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR14B(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR15()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR15(params["value"].get<unsigned long long>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR15D()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR15D(params["value"].get<unsigned int>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR15W()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR15W(params["value"].get<unsigned short>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetR15B()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetR15B(params["value"].get<unsigned char>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCAX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCAX(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCBX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCBX(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCCX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCCX(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCDX()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCDX(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCDI()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCDI(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCSI()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCSI(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCBP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCBP(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCSP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCSP(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCIP()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCIP(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::GetCFLAGS()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Register::SetCFLAGS(params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Stack::Pop()};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Stack::Push(params["value"].get<duint>())};
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
            ApiResponse resp{.success = true, .type = "", .result = Script::Stack::Peek(params["offset"].get<int>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/_scriptapi_symbol.h/GetList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = Script::Symbol::GetList(params["list"].get<ListInfo *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgegraph.h/Free", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::Free(params["graphList"].get<const BridgeCFGraphList *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgegraph.h/ToNodeList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFNode::ToNodeList()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgegraph.h/__debugbreak", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = __debugbreak()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgegraph.h/AddNode", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::AddNode(params["node"].get<const BridgeCFNode &>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgegraph.h/AddParent", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::AddParent(params["child"].get<int>(), params["parent"].get<int>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgegraph.h/ToGraphList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::ToGraphList()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeInit", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeInit()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeLoadLibraryCheckedW", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeLoadLibraryCheckedW(params["szDll"].get<const wchar_t *>(), params["allowFailure"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeLoadLibraryCheckedA", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeLoadLibraryCheckedA(params["szDll"].get<const char *>(), params["allowFailure"].get<bool>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeStart", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeStart()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeAlloc", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeAlloc(params["size"].get<size_t>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeFree", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeFree(params["ptr"].get<void *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeSettingGet", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeSettingGet(params["section"].get<const char *>(), params["key"].get<const char *>(), params["value"].get<char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeSettingGetUint", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeSettingGetUint(params["section"].get<const char *>(), params["key"].get<const char *>(), params["value"].get<duint *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeSettingSet", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeSettingSet(params["section"].get<const char *>(), params["key"].get<const char *>(), params["value"].get<const char *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeSettingSetUint", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeSettingSetUint(params["section"].get<const char *>(), params["key"].get<const char *>(), params["value"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeSettingFlush", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeSettingFlush()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeSettingRead", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeSettingRead(params["errorLine"].get<int *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeGetDbgVersion", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeGetDbgVersion()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeIsProcessElevated", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeIsProcessElevated()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeGetNtBuildNumber", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeGetNtBuildNumber()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeUserDirectory", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeUserDirectory()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/BridgeIsARM64Emulated", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = BridgeIsARM64Emulated()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/Data", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = UBridgeCFNodeList::? $BridgeList::Data() };
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/Count", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = UBridgeCFNodeList::? $BridgeList::Count() };
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/Cleanup", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = UBridgeCFNodeList::? $BridgeList::Cleanup() };
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/CopyData", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = UBridgeCFNodeList::? $BridgeList::CopyData(params["listInfo"].get<ListInfo *>(), params["listData"].get<const std::vector<BridgeCFNodeList> &>()) };
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/Free", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::Free(params["graphList"].get<const BridgeCFGraphList *>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/ToVector", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = UBridgeCFNodeList::? $BridgeList::ToVector(params["listInfo"].get<const ListInfo *>(), params["listData"].get<std::vector<BridgeCFNodeList> &>(), params["freedata"].get<bool>()) };
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/ToNodeList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFNode::ToNodeList()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/__debugbreak", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = __debugbreak()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/AddNode", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::AddNode(params["node"].get<const BridgeCFNode &>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/AddParent", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::AddParent(params["child"].get<duint>(), params["parent"].get<duint>())};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/ToGraphList", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = ::BridgeCFGraph::ToGraphList()};
            res.set_content(json(resp).dump(), "application/json");
        } catch (const std::exception &e) {
            res.set_content(json{{"success", false}, {"error", e.what()}}, "application/json");
        }
    });

    server.Post("/bridgemain.h/__builtin_memcpy", [](const Request &req, Response &res) {
        try {
            auto arg = nlohmann::json::parse(req.body).get<std::vector<Param>>();
            json params;
            for (const auto &param: arg) {
                params[param.name] = param.value;
            }
            ApiResponse resp{.success = true, .type = "", .result = __builtin_memcpy(params[""].get<void *>(), params[""].get<const void *>(), params[""].get<unsigned long long>())};
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

    if (g_httpServerRunning) {
        stopHttpServer();
    }

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
        if (g_httpServerThread != NULL) {
            WaitForSingleObject(g_httpServerThread, 1000);
            CloseHandle(g_httpServerThread);
            g_httpServerThread = NULL;
        }
    }
}

#endif
