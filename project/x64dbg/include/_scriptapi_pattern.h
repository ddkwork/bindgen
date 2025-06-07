#ifndef _API_PATTERN_H
#define _API_PATTERN_H

#include "_scriptapi.h"

namespace api {
    namespace Pattern {
        ApiResponse Find(duint datasize, std::string pattern) {
            std::string data;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };

        ApiResponse FindMem(duint start, duint size, std::string pattern) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Write(duint datasize, std::string pattern) {
            std::string data;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };

        ApiResponse WriteMem(duint start, duint size, std::string pattern) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse SearchAndReplace(duint datasize, std::string searchpattern, std::string replacepattern) {
            std::string data;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };

        ApiResponse SearchAndReplaceMem(duint start, duint size, std::string searchpattern, std::string replacepattern) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }
    };// namespace Pattern
};// namespace api

#endif//_API_FIND_H