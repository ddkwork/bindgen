#ifndef _API_FUNCTION_H
#define _API_FUNCTION_H

#include "_scriptapi.h"

namespace api {
    namespace Function {
        struct FunctionInfo {
            char  mod[MAX_MODULE_SIZE];
            duint rvaStart;
            duint rvaEnd;
            bool  manual;
            duint instructioncount;
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(FunctionInfo, mod, rvaStart, rvaEnd, manual, instructioncount)
        };

        ApiResponse Add(duint start, duint end, bool manual, duint instructionCount = 0) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Add(const FunctionInfo *info) {
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

        ApiResponse GetInfo(duint addr, FunctionInfo *info) {
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

        ApiResponse DeleteRange(duint start, duint end, bool deleteManual) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse DeleteRange(duint start, duint end) {
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
            ListOf(FunctionInfo) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//caller has the responsibility to free the list
    };// namespace Function
};// namespace api

#endif//_API_FUNCTION_H
