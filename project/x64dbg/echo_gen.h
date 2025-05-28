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

#ifdef _WIN64
#pragma comment(lib, "pluginsdk/x64dbg.lib")
#pragma comment(lib, "pluginsdk/x64bridge.lib")
#pragma comment(lib, "pluginsdk/DeviceNameResolver/DeviceNameResolver_x64.lib")
#pragma comment(lib, "pluginsdk/jansson/jansson_x64.lib")
#pragma comment(lib, "pluginsdk/lz4/lz4_x64.lib")
#pragma comment(lib, "pluginsdk/TitanEngine/TitanEngine_x64.lib")
#pragma comment(lib, "pluginsdk/XEDParse/XEDParse_x64.lib")
#else
#pragma comment(lib, "pluginsdk/x32dbg.lib")
#pragma comment(lib, "pluginsdk/x32bridge.lib")
#pragma comment(lib, "pluginsdk/DeviceNameResolver/DeviceNameResolver_x86.lib")
#pragma comment(lib, "pluginsdk/jansson/jansson_x86.lib")
#pragma comment(lib, "pluginsdk/lz4/lz4_x86.lib")
#pragma comment(lib, "pluginsdk/TitanEngine/TitanEngine_x86.lib")
#pragma comment(lib, "pluginsdk/XEDParse/XEDParse_x86.lib")
#endif //_WIN64问题是为
#pragma comment(lib, "shlwapi.lib")


#include <map>
#include <string>
#include <functional>

void dispatch(const std::string &path) {
    static const std::unordered_map<std::string, std::function<void()>> handlers{
            {"_scriptapi_argument.h/Add",   [] { Script::Argument::Add(nullptr); }},//todo add more logic
            {"_scriptapi_argument.h/Clear", [] { Script::Argument::Clear(); }}
    };
    if (auto it = handlers.find(path); it != handlers.end()) it->second();//todo: 这里的路径解析有问题，发包恢复客户端
}


/*
void echo(const std::string path) {
    // 使用 std::map 来存储路径和对应的操作函数
    std::map<std::string, std::function<void()>> pathActions = {
            {"_scriptapi_argument.h/Add",   []() {
                Script::Argument::Add(nullptr);
            }},
            {"_scriptapi_argument.h/Clear", []() {
                Script::Argument::Clear();
            }},
    };

    // 查找并执行对应的操作
    auto it = pathActions.find(path);
    if (it != pathActions.end()) {
        it->second();
    } else {
        // 默认情况下什么都不做
    }
}


#include <iostream>
#include <string>

// 编译期哈希计算（DJB2算法）
constexpr size_t hash(const char *str) noexcept {
    size_t hash = 5381;
    while (*str != '\0') {
        hash = ((hash << 5) + hash) + *str; // hash = hash * 33 + c
        ++str;
    }
    return hash;
}

void dispatch2(const std::string &path) {
    const size_t h = hash(path.c_str());
    switch (h) {
        case hash("Namespace1::myFunc"):
            if (path != "Namespace1::myFunc") break; // 防哈希碰撞
            std::cout << "Call Namespace1::myFunc\n";
            break;
        case hash("Namespace2::myFunc"):
            if (path != "Namespace2::myFunc") break;
            std::cout << "Call Namespace2::myFunc\n";
            break;
        default:
            std::cerr << "Unknown path: " << path << "\n";
    }
}

int main() {
    dispatch("Namespace1::myFunc");
    dispatch("Namespace2::myFunc");
    dispatch("InvalidPath");
    return 0;
}
 */

#endif //MCPX64DBG_ECHO_GEN_H
