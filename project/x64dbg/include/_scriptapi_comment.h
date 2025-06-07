#ifndef _API_COMMENT_H
#define _API_COMMENT_H

#include "_scriptapi.h"

namespace api {
    namespace Comment {
        struct CommentInfo {
            char  mod[MAX_MODULE_SIZE];
            duint rva;
            char  text[MAX_LABEL_SIZE];
            bool  manual;
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(CommentInfo, mod, rva, text, manual)
        };

        ApiResponse Set(duint addr, std::string text, bool manual = false) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Set(const CommentInfo *info) {
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        }

        ApiResponse Get(duint addr) {
            char text[MAX_COMMENT_SIZE];
            auto ok = Script::Comment::Get(addr, text);
            return ApiResponse{
                    .success = ok,
                    .type = "std::string",
                    .result = std::string(text)};
        };//text[MAX_COMMENT_SIZE]

        ApiResponse GetInfo(duint addr, CommentInfo *info) {
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
            //                 ListOf(CommentInfo) list;
            //                 Script::Comment::GetList(&list){
            //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };//caller has the responsibility to free the list
    };// namespace Comment
};// namespace api

#endif//_API_COMMENT_H