#ifndef _API_ARGUMENT_H
#define _API_ARGUMENT_H

#include "_scriptapi.h"

namespace api {
    namespace Argument {
        struct ArgumentInfo {
            char  mod[MAX_MODULE_SIZE];
            duint rvaStart;
            duint rvaEnd;
            bool  manual;
            duint instructioncount;
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(ArgumentInfo, mod, rvaStart, rvaEnd, manual, instructioncount)
        };

        ApiResponse Add(duint start, duint end, bool manual, duint instructionCount = 0) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Add(const ArgumentInfo *info) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Get(duint addr, duint *start = nullptr, duint *end = nullptr, duint *instructionCount = nullptr) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse GetInfo(duint addr, ArgumentInfo *info) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Overlaps(duint start, duint end) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Delete(duint address) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse DeleteRange(duint start, duint end, bool deleteManual = false) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Clear() {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse GetList() {
            ListOf(ArgumentInfo) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//caller has the responsibility to free the list
    };// namespace Argument
};// namespace api

#endif//_API_ARGUMENT_H