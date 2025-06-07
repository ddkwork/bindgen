#ifndef _API_MISC_H
#define _API_MISC_H

#include "_scriptapi.h"

namespace api {
    namespace Misc {
        ///     bool success = ParseExpression("[esp+8]", &amp;val)
        ApiResponse ParseExpression(const char *expression, duint *value) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ///     duint addr = RemoteGetProcAddress("kernel32.dll", "GetProcAddress")
        ApiResponse RemoteGetProcAddress(const char *module, const char *api) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ///     duint addr = ResolveLabel("sneaky_crypto")
        ApiResponse ResolveLabel(const char *label) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

    };// namespace Misc
};// namespace api

#endif//_API_MISC_H