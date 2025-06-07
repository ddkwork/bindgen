#ifndef _API_MODULE_H
#define _API_MODULE_H

#include "_scriptapi.h"

namespace api {
    namespace Module {
        struct ModuleInfo {
            duint base;
            duint size;
            duint entry;
            int   sectionCount;
            char  name[MAX_MODULE_SIZE];
            char  path[MAX_PATH];
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(ModuleInfo, base, size, entry, sectionCount, name, path)
        };

        struct ModuleSectionInfo {
            duint addr;
            duint size;
            char  name[MAX_SECTION_SIZE * 5];
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(ModuleSectionInfo, addr, size, name)
        };

        struct ModuleExport {
            duint ordinal;
            duint rva;
            duint va;
            bool  forwarded;
            char  forwardName[MAX_STRING_SIZE];
            char  name[MAX_STRING_SIZE];
            char  undecoratedName[MAX_STRING_SIZE];
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(ModuleExport, ordinal, rva, va, forwarded, forwardName, name, undecoratedName)
        };

        struct ModuleImport {
            duint iatRva;
            duint iatVa;
            duint ordinal;//equal to -1 if imported by name
            char  name[MAX_STRING_SIZE];
            char  undecoratedName[MAX_STRING_SIZE];
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(ModuleImport, iatRva, iatVa, ordinal, name, undecoratedName)
        };

        ApiResponse InfoFromAddr(duint addr, ModuleInfo *info) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse InfoFromName(std::string name, ModuleInfo *info) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse BaseFromAddr(duint addr) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse BaseFromName(std::string name) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse SizeFromAddr(duint addr) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse SizeFromName(std::string name) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse NameFromAddr(duint addr) {
            char name[MAX_MODULE_SIZE] = {0};
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//name[MAX_MODULE_SIZE]

        ApiResponse PathFromAddr(duint addr) {
            char path[MAX_PATH] = {0};
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//path[MAX_PATH]

        ApiResponse PathFromName(std::string name) {
            char path[MAX_PATH] = {0};
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//path[MAX_PATH]

        ApiResponse EntryFromAddr(duint addr) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse EntryFromName(std::string name) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse SectionCountFromAddr(duint addr) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse SectionCountFromName(std::string name) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse SectionFromAddr(duint addr, int number, ModuleSectionInfo *section) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse SectionFromName(std::string name, int number, ModuleSectionInfo *section) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse SectionListFromAddr(duint addr) {
            ListOf(ModuleSectionInfo) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };

        ApiResponse SectionListFromName(std::string name) {
            ListOf(ModuleSectionInfo) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };

        ApiResponse GetMainModuleInfo(ModuleInfo *info) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse GetMainModuleBase() {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse GetMainModuleSize() {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse GetMainModuleEntry() {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse GetMainModuleSectionCount() {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse GetMainModuleName() {
            char name[MAX_MODULE_SIZE] = {0};
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//name[MAX_MODULE_SIZE]

        ApiResponse GetMainModulePath() {
            char path[MAX_PATH] = {0};
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//path[MAX_PATH]

        ApiResponse GetMainModuleSectionList() {
            ListOf(ModuleSectionInfo) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//caller has the responsibility to free the list

        ApiResponse GetList() {
            ListOf(ModuleInfo) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//caller has the responsibility to free the list

        ApiResponse GetExports(const ModuleInfo *mod) {
            ListOf(ModuleExport) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//caller has the responsibility to free the list

        ApiResponse GetImports(const ModuleInfo *mod) {
            ListOf(ModuleImport) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//caller has the responsibility to free the list
    };// namespace Module
};// namespace api

#endif//_API_MODULE_H
