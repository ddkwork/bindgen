#ifndef _API_ASSEMBLER_H
#define _API_ASSEMBLER_H

#include "_scriptapi.h"

namespace api {
    namespace Assembler {

        struct AssemblerResult {
            int         size;
            std::string dest;
            std::string error;
            NLOHMANN_DEFINE_TYPE_INTRUSIVE(AssemblerResult, size, dest, error)
        };

        ApiResponse Assemble(duint addr, std::string instruction) {
            unsigned char dest[16] = {0};
            int           size = 0;
            auto          ok = Script::Assembler::Assemble(addr, (unsigned char *) dest, &size, instruction.c_str());
            auto          result = AssemblerResult{
                             .size = size,
                             .dest = reinterpret_cast<const char *const>(dest),
                             .error = "",
            };
            return ApiResponse{
                    .success = ok,
                    .type = "AssemblerResult",
                    .result = result,
            };
        };//dest[16]

        ApiResponse AssembleEx(duint addr, std::string instruction) {
            unsigned char dest[16] = {0};
            int           size = 0;
            char          err[MAX_ERROR_SIZE] = {0};
            auto          ok = Script::Assembler::AssembleEx(addr, (unsigned char *) dest, &size, instruction.c_str(), err);
            auto          result = AssemblerResult{
                             .size = size,
                             .dest = reinterpret_cast<const char *const>(dest),
                             .error = std::string(err),
            };
            return ApiResponse{
                    .success = ok,
                    .type = "AssemblerResult",
                    .result = result,
            };//dest[16], error[MAX_ERROR_SIZE]
        };//dest[16], error[MAX_ERROR_SIZE]

        ApiResponse AssembleMem(duint addr, std::string instruction) {
            unsigned char dest[16] = {0};
            int           size = 0;
            auto          ok = Script::Assembler::AssembleMem(addr, instruction.c_str());
            auto          result = AssemblerResult{
                             .size = size,
                             .dest = reinterpret_cast<const char *const>(dest),
                             .error = "",
            };
            return ApiResponse{
                    .success = ok,
                    .type = "AssemblerResult",
                    .result = result,
            };
        };

        ApiResponse AssembleMemEx(duint addr, std::string instruction, bool fillnop) {
            char err[MAX_ERROR_SIZE] = {0};
            int  size = 0;
            auto ok = Script::Assembler::AssembleMemEx(addr, instruction.c_str(), &size, const_cast<char *>(err), fillnop);
            auto result = AssemblerResult{
                    .size = size,
                    .dest = "",
                    .error = std::string(err),
            };
            return ApiResponse{
                    .success = ok,
                    .type = "AssemblerResult",
                    .result = result,
            };
        };//error[MAX_ERROR_SIZE]
    };// namespace Assembler
};// namespace api

#endif//_API_ASSEMBLER_H