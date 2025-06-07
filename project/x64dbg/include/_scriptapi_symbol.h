#ifndef _API_SYMBOL_H
#define _API_SYMBOL_H

#include "_scriptapi.h"

namespace api
{
    namespace Symbol
    {
        enum SymbolType
        {
            Function, //user-defined function
            Import, //IAT entry
            Export //export
        };

        struct SymbolInfo
        {
            char mod[MAX_MODULE_SIZE];
            duint rva;
            char name[MAX_LABEL_SIZE];
            bool manual;
            SymbolType type;
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(SymbolInfo, mod, rva, name, manual, type)
        };

        ApiResponse GetList(){
                ListOf(SymbolInfo) list;
                return ApiResponse{
                        .success = true,
                        // .type="uintptr",
                        //.result=handle
                };
        }; //caller has the responsibility to free the list
    }; //Symbol
}; //Script

#endif //_API_SYMBOL_H