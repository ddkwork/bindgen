#ifndef _API_LABEL_H
#define _API_LABEL_H

#include "_scriptapi.h"

namespace api
{
    namespace Label
    {
        struct LabelInfo
        {
            char mod[MAX_MODULE_SIZE];
            duint rva;
            char text[MAX_LABEL_SIZE];
            bool manual;
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(LabelInfo, mod, rva, text, manual)
        };

         ApiResponse Set(duint addr, std::string text, bool manual = false){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse Set(duint addr, std::string text, bool manual = false, bool temporary = false){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse Set(const LabelInfo* info){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse FromString(std::string label){
             duint* addr;
             return ApiResponse{
                     .success = true,
                     // .type="uintptr",
                     //.result=handle
             };
         };
         ApiResponse Get(duint addr,std::string text){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    } //text[MAX_LABEL_SIZE]
         ApiResponse IsTemporary(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetInfo(duint addr, LabelInfo* info){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse Delete(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse DeleteRange(duint start, duint end){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse Clear(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetList(){
                 ListOf(LabelInfo) list;
                 return ApiResponse{
                         .success = true,
                         // .type="uintptr",
                         //.result=handle
                 };
         }; //caller has the responsibility to free the list
    }; //Label
}; //Script

#endif //_API_LABEL_H