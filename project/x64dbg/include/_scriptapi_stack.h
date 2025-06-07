#ifndef _API_STACK_H
#define _API_STACK_H

#include "_scriptapi.h"

namespace api {
    namespace Stack {
        ApiResponse Pop() {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Push(duint value) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }//returns the previous top, equal to Peek(1)

        ApiResponse Peek(int offset = 0) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }//offset is in multiples of Register::Size(), for easy x32/x64 portability
    };// namespace Stack
};// namespace api

#endif//_API_STACK_H