#ifndef _API_MEMORY_H
#define _API_MEMORY_H

#include "_scriptapi.h"

namespace api {
    namespace Memory {
        ApiResponse Read(duint addr, duint size) {
            duint *sizeRead;;
            void *data;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };
        ApiResponse Write(duint addr,  std::vector<unsigned char> data, duint size){
            duint *sizeWritten;
            return ApiResponse{
                    .success = true,
                    // .type="uintptr",
                    //.result=handle
            };
        };
        ApiResponse IsValidPtr(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse RemoteAlloc(duint addr, duint size){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse RemoteFree(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse GetProtect(duint addr, bool reserved = false, bool cache = true){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse SetProtect(duint addr, unsigned int protect, duint size){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse GetBase(duint addr, bool reserved = false, bool cache = true){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse GetSize(duint addr, bool reserved = false, bool cache = true){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }

        ApiResponse ReadByte(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse WriteByte(duint addr, unsigned char data){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse ReadWord(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse WriteWord(duint addr, unsigned short data){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse ReadDword(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse WriteDword(duint addr, unsigned int data){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse ReadQword(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse WriteQword(duint addr, unsigned long long data){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse ReadPtr(duint addr){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
        ApiResponse WritePtr(duint addr, duint data){
        //auto handle=BridgeLoadLibraryCheckedW(szDll, allowFailure);
        return ApiResponse{
                .success=true,
               // .type="uintptr",
                //.result=handle
        };
    }
    }// namespace Memory
}// namespace api

#endif//_API_MEMORY_H