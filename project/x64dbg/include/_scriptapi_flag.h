#ifndef _API_FLAG_H
#define _API_FLAG_H

#include "_scriptapi.h"

namespace api
{
    namespace Flag
    {
        enum FlagEnum
        {
            ZF,
            OF,
            CF,
            PF,
            SF,
            TF,
            AF,
            DF,
            IF
        };

         ApiResponse Get(FlagEnum flag){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse Set(FlagEnum flag, bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }

         ApiResponse GetZF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetZF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetOF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetOF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetCF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetCF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetPF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetPF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetSF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetSF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetTF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetTF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetAF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetAF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetDF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetDF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse GetIF(){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
         ApiResponse SetIF(bool value){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
    };
};

#endif //_API_FLAG_H