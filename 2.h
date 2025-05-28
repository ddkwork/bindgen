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