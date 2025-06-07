#ifndef _API_BOOKMARK_H
#define _API_BOOKMARK_H

#include "_scriptapi.h"

namespace api {
    namespace Bookmark {
        struct BookmarkInfo {
            char  mod[MAX_MODULE_SIZE];
            duint rva;
            bool  manual;
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(BookmarkInfo, mod, rva, manual)
        };

        ApiResponse Set(duint addr, bool manual = false) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Set(const BookmarkInfo *info) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Get(duint addr) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse GetInfo(duint addr, BookmarkInfo *info) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Delete(duint addr) {
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
            ListOf(BookmarkInfo) list;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//caller has the responsibility to free the list
    };// namespace Bookmark
};// namespace api

#endif//_API_BOOKMARK_H