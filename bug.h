//  struct ModuleInfo
//        {
//            int sectionCount;
//        };
//typedef enum
//{
//    initialized,
//    paused,
//    running,
//    stopped
//} DBGSTATE;
//
//typedef struct
//{
//    int count; //Number of element in the list.
//    size_t size; //Size of list in bytes (used for type checking).
//    void* data; //Pointer to the list contents. Must be deleted by the caller using BridgeFree (or BridgeList::Free).
//} ListInfo;

// enum Color {
//     RED,
//     GREEN,
//     YELLOW =99
// };

//typedef enum
//{
//    SEG_DEFAULT,
//    SEG_ES,
//    SEG_DS,
//    SEG_FS,
//    SEG_GS,
//    SEG_CS,
//    SEG_SS
//} SEGMENTREG;


// struct Car {
//     char make[50];
//     int year;
// };
//
//int add (int a, int b){
//    return a+b;
//};

//typedef unsigned char (*CHECK_VMX_OPERATION)();
//#define MAX_PATH 260
//typedef unsigned __int64   UINT64, *PUINT64;

//typedef struct _CR3_TYPE
//{
//    union
//    {
//        UINT64 Flags;
//
//        struct
//        {
//            UINT64 Pcid : 12;
//            UINT64 PageFrameNumber : 36;
//            UINT64 Reserved1 : 12;
//            UINT64 Reserved_2 : 3;
//            UINT64 PcidInvalidate : 1;
//        } Fields;
//    };
//} CR3_TYPE, *PCR3_TYPE;

//namespace Script
//{
//    namespace Debug
//    {
//        enum HardwareType
//        {
//            HardwareAccess,
//            HardwareWrite,
//            HardwareExecute
//        };
//    }
//}

namespace Script
{
    namespace Module
    {
        struct ModuleInfo
        {
            int sectionCount;
         }; //Module
    }
}; //Script

namespace Script
{
    namespace Argument
    {
        struct ArgumentInfo
        {
            bool manual;
        };

        SCRIPT_EXPORT bool Add(duint start, duint end, bool manual, duint instructionCount = 0);
        SCRIPT_EXPORT bool Add(const ArgumentInfo* info);
        SCRIPT_EXPORT bool Get(duint addr, duint* start = nullptr, duint* end = nullptr, duint* instructionCount = nullptr);
        SCRIPT_EXPORT bool GetInfo(duint addr, ArgumentInfo* info);
        SCRIPT_EXPORT bool Overlaps(duint start, duint end);
        SCRIPT_EXPORT bool Delete(duint address);
        SCRIPT_EXPORT void DeleteRange(duint start, duint end, bool deleteManual = false);
        SCRIPT_EXPORT void Clear();
        SCRIPT_EXPORT bool GetList(ListOf(ArgumentInfo) list); //caller has the responsibility to free the list
    }; //Argument
}; //Script