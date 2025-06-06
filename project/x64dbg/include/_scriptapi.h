#ifndef _SCRIPTAPI_H_
#define _SCRIPTAPI_H_

//dispatch build need
#include "../json.hpp"
#include "../pluginsdk/_scriptapi.h"
#include "../pluginsdk/bridgemain.h"
#include "../pluginsdk/_scriptapi_argument.h"
#include "../pluginsdk/_scriptapi_assembler.h"
#include "../pluginsdk/_scriptapi_bookmark.h"
#include "../pluginsdk/_scriptapi_comment.h"
#include "../pluginsdk/_scriptapi_debug.h"
#include "../pluginsdk/_scriptapi_flag.h"
#include "../pluginsdk/_scriptapi_function.h"
#include "../pluginsdk/_scriptapi_gui.h"
#include "../pluginsdk/_scriptapi_label.h"
#include "../pluginsdk/_scriptapi_memory.h"
#include "../pluginsdk/_scriptapi_misc.h"
#include "../pluginsdk/_scriptapi_module.h"
#include "../pluginsdk/_scriptapi_pattern.h"
#include "../pluginsdk/_scriptapi_register.h"
#include "../pluginsdk/_scriptapi_stack.h"
#include "../pluginsdk/_scriptapi_symbol.h"

using json = nlohmann::json;

struct ApiResponse {
    bool        success;
    std::string type;//return go type
    json        result;
    NLOHMANN_DEFINE_TYPE_INTRUSIVE(ApiResponse, success, type, result)
};
#endif//_SCRIPTAPI_H_
      /*  clang dump ast need
#define MAX_PATH 256
#define MAX_LABEL_SIZE 256
#define MAX_COMMENT_SIZE 512
#define MAX_MODULE_SIZE 256
#define MAX_IMPORT_SIZE 65536
#define MAX_BREAKPOINT_SIZE 256
#define MAX_CONDITIONAL_EXPR_SIZE 256
#define MAX_CONDITIONAL_TEXT_SIZE 256
#define MAX_SCRIPT_LINE_SIZE 2048
#define MAX_THREAD_NAME_SIZE 256
#define MAX_WATCH_NAME_SIZE 256
#define MAX_STRING_SIZE 512
#define MAX_ERROR_SIZE 512
#define RIGHTS_STRING_SIZE (sizeof("ERWCG"))
#define MAX_SECTION_SIZE 10
#define MAX_COMMAND_LINE_SIZE 256
#define MAX_MNEMONIC_SIZE 64

#ifndef PAGE_SIZE
#define PAGE_SIZE 0x1000
#endif // PAGE_SIZE

#ifdef _WIN64
typedef unsigned long long duint;
typedef signed long long dsint;
#else
typedef unsigned long duint;
typedef signed long dsint;
#endif //_WIN64

typedef enum {
       initialized,
       paused,
       running,
       stopped
   } DBGSTATE;

typedef enum {
   SEG_DEFAULT,
   SEG_ES,
   SEG_DS,
   SEG_FS,
   SEG_GS,
   SEG_CS,
   SEG_SS
} SEGMENTREG;

typedef enum {
   flagmodule = 0x1,
   flaglabel = 0x2,
   flagcomment = 0x4,
   flagbookmark = 0x8,
   flagfunction = 0x10,
   flagloop = 0x20,
   flagargs = 0x40,
   flagNoFuncOffset = 0x80
} ADDRINFOFLAGS;

typedef enum {
   bp_none = 0,
   bp_normal = 1,
   bp_hardware = 2,
   bp_memory = 4,
   bp_dll = 8,
   bp_exception = 16
} BPXTYPE;

typedef enum {
   FUNC_NONE,
   FUNC_BEGIN,
   FUNC_MIDDLE,
   FUNC_END,
   FUNC_SINGLE
} FUNCTYPE;

typedef enum {
   LOOP_NONE,
   LOOP_BEGIN,
   LOOP_MIDDLE,
   LOOP_ENTRY,
   LOOP_END,
   LOOP_SINGLE
} LOOPTYPE;

typedef enum {
   XREF_NONE,
   XREF_DATA,
   XREF_JMP,
   XREF_CALL
} XREFTYPE;

typedef enum {
   ARG_NONE,
   ARG_BEGIN,
   ARG_MIDDLE,
   ARG_END,
   ARG_SINGLE
} ARGTYPE;

typedef enum {
   DBG_SCRIPT_LOAD,               // param1=std::string filename,      param2=unused
   DBG_SCRIPT_UNLOAD,             // param1=unused,                    param2=unused
   DBG_SCRIPT_RUN,                // param1=int destline,              param2=unused
   DBG_SCRIPT_STEP,               // param1=unused,                    param2=unused
   DBG_SCRIPT_BPTOGGLE,           // param1=int line,                  param2=unused
   DBG_SCRIPT_BPGET,              // param1=int line,                  param2=unused
   DBG_SCRIPT_CMDEXEC,            // param1=std::string command,       param2=unused
   DBG_SCRIPT_ABORT,              // param1=unused,                    param2=unused
   DBG_SCRIPT_GETLINETYPE,        // param1=int line,                  param2=unused
   DBG_SCRIPT_SETIP,              // param1=int line,                  param2=unused
   DBG_SCRIPT_GETBRANCHINFO,      // param1=int line,                  param2=SCRIPTBRANCH* info
   DBG_SYMBOL_ENUM,               // param1=SYMBOLCBINFO* cbInfo,      param2=unused
   DBG_ASSEMBLE_AT,               // param1=duint addr,                param2=std::string instruction
   DBG_MODBASE_FROM_NAME,         // param1=std::string modname,       param2=unused
   DBG_DISASM_AT,                 // param1=duint addr,                 param2=DISASM_INSTR* instr
   DBG_STACK_COMMENT_GET,         // param1=duint addr,                param2=STACK_COMMENT* comment
   DBG_GET_THREAD_LIST,           // param1=THREADALLINFO* list,       param2=unused
   DBG_SETTINGS_UPDATED,          // param1=unused,                    param2=unused
   DBG_DISASM_FAST_AT,            // param1=duint addr,                param2=BASIC_INSTRUCTION_INFO* basicinfo
   DBG_MENU_ENTRY_CLICKED,        // param1=int hEntry,                param2=unused
   DBG_FUNCTION_GET,              // param1=FUNCTION_LOOP_INFO* info,  param2=unused
   DBG_FUNCTION_OVERLAPS,         // param1=FUNCTION_LOOP_INFO* info,  param2=unused
   DBG_FUNCTION_ADD,              // param1=FUNCTION_LOOP_INFO* info,  param2=unused
   DBG_FUNCTION_DEL,              // param1=FUNCTION_LOOP_INFO* info,  param2=unused
   DBG_LOOP_GET,                  // param1=FUNCTION_LOOP_INFO* info,  param2=unused
   DBG_LOOP_OVERLAPS,             // param1=FUNCTION_LOOP_INFO* info,  param2=unused
   DBG_LOOP_ADD,                  // param1=FUNCTION_LOOP_INFO* info,  param2=unused
   DBG_LOOP_DEL,                  // param1=FUNCTION_LOOP_INFO* info,  param2=unused
   DBG_IS_RUN_LOCKED,             // param1=unused,                    param2=unused
   DBG_IS_BP_DISABLED,            // param1=duint addr,                param2=unused
   DBG_SET_AUTO_COMMENT_AT,       // param1=duint addr,                param2=std::string text
   DBG_DELETE_AUTO_COMMENT_RANGE, // param1=duint start,               param2=duint end
   DBG_SET_AUTO_LABEL_AT,         // param1=duint addr,                param2=std::string text
   DBG_DELETE_AUTO_LABEL_RANGE,   // param1=duint start,               param2=duint end
   DBG_SET_AUTO_BOOKMARK_AT,      // param1=duint addr,                param2=std::string text
   DBG_DELETE_AUTO_BOOKMARK_RANGE,// param1=duint start,               param2=duint end
   DBG_SET_AUTO_FUNCTION_AT,      // param1=duint addr,                param2=std::string text
   DBG_DELETE_AUTO_FUNCTION_RANGE,// param1=duint start,               param2=duint end
   DBG_GET_STRING_AT,             // param1=duint addr,                param2=unused
   DBG_GET_FUNCTIONS,             // param1=unused,                    param2=unused
   DBG_WIN_EVENT,                 // param1=MSG* message,              param2=long* result
   DBG_WIN_EVENT_GLOBAL,          // param1=MSG* message,              param2=unused
   DBG_INITIALIZE_LOCKS,          // param1=unused,                    param2=unused
   DBG_DEINITIALIZE_LOCKS,        // param1=unused,                    param2=unused
   DBG_GET_TIME_WASTED_COUNTER,   // param1=unused,                    param2=unused
   DBG_SYMBOL_ENUM_FROMCACHE,     // param1=SYMBOLCBINFO* cbInfo,      param2=unused
   DBG_DELETE_COMMENT_RANGE,      // param1=duint start,               param2=duint end
   DBG_DELETE_LABEL_RANGE,        // param1=duint start,               param2=duint end
   DBG_DELETE_BOOKMARK_RANGE,     // param1=duint start,               param2=duint end
   DBG_GET_XREF_COUNT_AT,         // param1=duint addr,                param2=unused
   DBG_GET_XREF_TYPE_AT,          // param1=duint addr,                param2=unused
   DBG_XREF_ADD,                  // param1=duint addr,                param2=duint from
   DBG_XREF_DEL_ALL,              // param1=duint addr,                param2=unused
   DBG_XREF_GET,                  // param1=duint addr,                param2=XREF_INFO* info
   DBG_GET_ENCODE_TYPE_BUFFER,    // param1=duint addr,                param2=unused
   DBG_ENCODE_TYPE_GET,           // param1=duint addr,                param2=duint size
   DBG_DELETE_ENCODE_TYPE_RANGE,  // param1=duint start,               param2=duint end
   DBG_ENCODE_SIZE_GET,           // param1=duint addr,                param2=duint codesize
   DBG_DELETE_ENCODE_TYPE_SEG,    // param1=duint addr,                param2=unused
   DBG_RELEASE_ENCODE_TYPE_BUFFER,// param1=void* buffer,              param2=unused
   DBG_ARGUMENT_GET,              // param1=FUNCTION* info,            param2=unused
   DBG_ARGUMENT_OVERLAPS,         // param1=FUNCTION* info,            param2=unused
   DBG_ARGUMENT_ADD,              // param1=FUNCTION* info,            param2=unused
   DBG_ARGUMENT_DEL,              // param1=FUNCTION* info,            param2=unused
   DBG_GET_WATCH_LIST,            // param1=ListOf_fake(WATCHINFO),         param2=unused
   DBG_SELCHANGED,                // param1=hWindow,                   param2=VA
   DBG_GET_PROCESS_HANDLE,        // param1=unused,                    param2=unused
   DBG_GET_THREAD_HANDLE,         // param1=unused,                    param2=unused
   DBG_GET_PROCESS_ID,            // param1=unused,                    param2=unused
   DBG_GET_THREAD_ID,             // param1=unused,                    param2=unused
   DBG_GET_PEB_ADDRESS,           // param1=DWORD ProcessId,           param2=unused
   DBG_GET_TEB_ADDRESS,           // param1=DWORD ThreadId,            param2=unused
   DBG_ANALYZE_FUNCTION,          // param1=BridgeCFGraphList* graph,  param2=duint entry
   DBG_MENU_PREPARE,              // param1=int hMenu,                 param2=unused
   DBG_GET_SYMBOL_INFO,           // param1=void* symbol,              param2=SYMBOLINFO* info
   DBG_GET_DEBUG_ENGINE,          // param1=unused,                    param2-unused
   DBG_GET_SYMBOL_INFO_AT,        // param1=duint addr,                param2=SYMBOLINFO* info
   DBG_XREF_ADD_MULTI,            // param1=const XREF_EDGE* edges,    param2=duint count
} DBGMSG;

typedef enum {
   linecommand,
   linebranch,
   linelabel,
   linecomment,
   lineempty,
} SCRIPTLINETYPE;

typedef enum {
   scriptnobranch,
   scriptjmp,
   scriptjnejnz,
   scriptjejz,
   scriptjbjl,
   scriptjajg,
   scriptjbejle,
   scriptjaejge,
   scriptcall
} SCRIPTBRANCHTYPE;

typedef enum {
   instr_normal,
   instr_branch,
   instr_stack
} DISASM_INSTRTYPE;

typedef enum {
   arg_normal,
   arg_memory
} DISASM_ARGTYPE;

typedef enum {
   str_none,
   str_ascii,
   str_unicode
} STRING_TYPE;

typedef enum {
   _PriorityIdle = -15,
   _PriorityAboveNormal = 1,
   _PriorityBelowNormal = -1,
   _PriorityHighest = 2,
   _PriorityLowest = -2,
   _PriorityNormal = 0,
   _PriorityTimeCritical = 15,
   _PriorityUnknown = 0x7FFFFFFF
} THREADPRIORITY;

typedef enum {
   _Executive = 0,
   _FreePage = 1,
   _PageIn = 2,
   _PoolAllocation = 3,
   _DelayExecution = 4,
   _Suspended = 5,
   _UserRequest = 6,
   _WrExecutive = 7,
   _WrFreePage = 8,
   _WrPageIn = 9,
   _WrPoolAllocation = 10,
   _WrDelayExecution = 11,
   _WrSuspended = 12,
   _WrUserRequest = 13,
   _WrEventPair = 14,
   _WrQueue = 15,
   _WrLpcReceive = 16,
   _WrLpcReply = 17,
   _WrVirtualMemory = 18,
   _WrPageOut = 19,
   _WrRendezvous = 20,
   _Spare2 = 21,
   _Spare3 = 22,
   _Spare4 = 23,
   _Spare5 = 24,
   _WrCalloutStack = 25,
   _WrKernel = 26,
   _WrResource = 27,
   _WrPushLock = 28,
   _WrMutex = 29,
   _WrQuantumEnd = 30,
   _WrDispatchInt = 31,
   _WrPreempted = 32,
   _WrYieldExecution = 33,
   _WrFastMutex = 34,
   _WrGuardedMutex = 35,
   _WrRundown = 36,
} THREADWAITREASON;

typedef enum {
   size_byte = 1,
   size_word = 2,
   size_dword = 4,
   size_qword = 8,
   size_xmmword = 16,
   size_ymmword = 32
} MEMORY_SIZE;

typedef enum {
   enc_unknown,//must be 0
   enc_byte,   //1 byte
   enc_word,   //2 bytes
   enc_dword,  //4 bytes
   enc_fword,  //6 bytes
   enc_qword,  //8 bytes
   enc_tbyte,  //10 bytes
   enc_oword,  //16 bytes
   enc_mmword, //8 bytes
   enc_xmmword,//16 bytes
   enc_ymmword,//32 bytes
   enc_zmmword,//64 bytes avx512 not supported
   enc_real4,  //4 byte float
   enc_real8,  //8 byte double
   enc_real10, //10 byte decimal
   enc_ascii,  //ascii sequence
   enc_unicode,//unicode sequence
   enc_code,   //start of code
   enc_junk,   //junk code
   enc_middle  //middle of data
} ENCODETYPE;

typedef enum {
   TYPE_UINT,   // unsigned integer
   TYPE_INT,    // signed integer
   TYPE_FLOAT,  // single precision floating point value
   TYPE_ASCII,  // ascii string
   TYPE_UNICODE,// unicode string
   TYPE_INVALID // invalid watch expression or data type
} WATCHVARTYPE;

typedef enum {
   MODE_DISABLED,// watchdog is disabled
   MODE_ISTRUE,  // alert if expression is not 0
   MODE_ISFALSE, // alert if expression is 0
   MODE_CHANGED, // alert if expression is changed
   MODE_UNCHANGED// alert if expression is not changed
} WATCHDOGMODE;

typedef enum {
   hw_access,
   hw_write,
   hw_execute
} BPHWTYPE;

typedef enum {
   mem_access,
   mem_read,
   mem_write,
   mem_execute
} BPMEMTYPE;

typedef enum {
   dll_load = 1,
   dll_unload,
   dll_all
} BPDLLTYPE;

typedef enum {
   ex_firstchance = 1,
   ex_secondchance,
   ex_all
} BPEXTYPE;

typedef enum {
   hw_byte,
   hw_word,
   hw_dword,
   hw_qword
} BPHWSIZE;

typedef enum {
   sym_import,
   sym_export,
   sym_symbol
} SYMBOLTYPE;

#define SYMBOL_MASK_IMPORT _fake(1u << sym_import)
#define SYMBOL_MASK_EXPORT _fake(1u << sym_export)
#define SYMBOL_MASK_SYMBOL _fake(1u << sym_symbol)
#define SYMBOL_MASK_ALL    _fake(SYMBOL_MASK_IMPORT | SYMBOL_MASK_EXPORT | SYMBOL_MASK_SYMBOL)

typedef enum {
   mod_user,
   mod_system
} MODULEPARTY;

typedef enum {
   DebugEngineTitanEngine,
   DebugEngineGleeBug,
   DebugEngineStaticEngine,
} DEBUG_ENGINE;

typedef MEMORY_SIZE VALUE_SIZE;

typedef struct DBGFUNCTIONS_ DBGFUNCTIONS;

typedef struct
{
   MEMORY_BASIC_INFORMATION mbi;
   char                     info[MAX_MODULE_SIZE];
} MEMPAGE;

typedef struct
{
   int      count;
   MEMPAGE *page;
} MEMMAP;

typedef struct
{
   BPXTYPE        type;
   duint          addr;
   bool           enabled;
   bool           singleshoot;
   bool           active;
   char           name[MAX_BREAKPOINT_SIZE];
   char           mod[MAX_MODULE_SIZE];
   unsigned short slot;
   // extended part
   unsigned char typeEx;//BPHWTYPE/BPMEMTYPE/BPDLLTYPE/BPEXTYPE
   unsigned char hwSize;//BPHWSIZE
   unsigned int  hitCount;
   bool          fastResume;
   bool          silent;
   char          breakCondition[MAX_CONDITIONAL_EXPR_SIZE];
   char          logText[MAX_CONDITIONAL_TEXT_SIZE];
   char          logCondition[MAX_CONDITIONAL_EXPR_SIZE];
   char          commandText[MAX_CONDITIONAL_TEXT_SIZE];
   char          commandCondition[MAX_CONDITIONAL_EXPR_SIZE];
} BRIDGEBP;

typedef struct
{
   int       count;
   BRIDGEBP *bp;
} BPMAP;

typedef struct
{
   char         WatchName[MAX_WATCH_NAME_SIZE];
   char         Expression[MAX_CONDITIONAL_EXPR_SIZE];
   unsigned int window;
   unsigned int id;
   WATCHVARTYPE varType;
   WATCHDOGMODE watchdogMode;
   duint        value;
   bool         watchdogTriggered;
} WATCHINFO;

typedef struct
{
   duint start;     //OUT
   duint end;       //OUT
   duint instrcount;//OUT
} FUNCTION;

typedef struct
{
   int   depth;     //IN
   duint start;     //OUT
   duint end;       //OUT
   duint instrcount;//OUT
} LOOP;

typedef struct
{
   int      flags;                  //ADDRINFOFLAGS _fake(IN)
   char     module[MAX_MODULE_SIZE];//module the address is in
   char     label[MAX_LABEL_SIZE];
   char     comment[MAX_COMMENT_SIZE];
   bool     isbookmark;
   FUNCTION function;
   LOOP     loop;
   FUNCTION args;
} BRIDGE_ADDRINFO;

typedef struct SYMBOLINFO_ {
   duint      addr;
   char      *decoratedSymbol;
   char      *undecoratedSymbol;
   SYMBOLTYPE type;

   // If true: Use BridgeFree_fake(decoratedSymbol) to deallocate
   // Else: The decoratedSymbol pointer is valid until the module unloads
   bool freeDecorated;

   // If true: Use BridgeFree_fake(undecoratedSymbol) to deallcoate
   // Else: The undecoratedSymbol pointer is valid until the module unloads
   bool freeUndecorated;

   // The entry point pseudo-export has ordinal == 0 _fake(invalid ordinal value)
   DWORD ordinal;
} SYMBOLINFO;

typedef struct
{
   duint base;
   char  name[MAX_MODULE_SIZE];
} SYMBOLMODULEINFO;

typedef struct
{
   duint        base;
   CBSYMBOLENUM cbSymbolEnum;
   void        *user;
   duint        start;
   duint        end;
   unsigned int symbolMask;
} SYMBOLCBINFO;

typedef struct
{
   bool c;
   bool p;
   bool a;
   bool z;
   bool s;
   bool t;
   bool i;
   bool d;
   bool o;
} FLAGS;

typedef struct
{
   bool FZ;
   bool PM;
   bool UM;
   bool OM;
   bool ZM;
   bool IM;
   bool DM;
   bool DAZ;
   bool PE;
   bool UE;
   bool OE;
   bool ZE;
   bool DE;
   bool IE;

   unsigned short RC;
} MXCSRFIELDS;

typedef struct
{
   bool B;
   bool C3;
   bool C2;
   bool C1;
   bool C0;
   bool ES;
   bool SF;
   bool P;
   bool U;
   bool O;
   bool Z;
   bool D;
   bool I;

   unsigned short TOP;

} X87STATUSWORDFIELDS;

typedef struct
{
   bool IC;
   bool IEM;
   bool PM;
   bool UM;
   bool OM;
   bool ZM;
   bool DM;
   bool IM;

   unsigned short RC;
   unsigned short PC;

} X87CONTROLWORDFIELDS;

typedef struct DECLSPEC_ALIGN(16) _XMMREGISTER {
   ULONGLONG Low;
   LONGLONG  High;
} XMMREGISTER;

typedef struct
{
   XMMREGISTER Low; //XMM/SSE part
   XMMREGISTER High;//AVX part
} YMMREGISTER;

typedef struct
{
   BYTE data[10];
   int  st_value;
   int  tag;
} X87FPUREGISTER;

typedef struct
{
   WORD  ControlWord;
   WORD  StatusWord;
   WORD  TagWord;
   DWORD ErrorOffset;
   DWORD ErrorSelector;
   DWORD DataOffset;
   DWORD DataSelector;
   DWORD Cr0NpxState;
} X87FPU;

typedef struct
{
   ULONG_PTR cax;
   ULONG_PTR ccx;
   ULONG_PTR cdx;
   ULONG_PTR cbx;
   ULONG_PTR csp;
   ULONG_PTR cbp;
   ULONG_PTR csi;
   ULONG_PTR cdi;
#ifdef _WIN64
   ULONG_PTR r8;
   ULONG_PTR r9;
   ULONG_PTR r10;
   ULONG_PTR r11;
   ULONG_PTR r12;
   ULONG_PTR r13;
   ULONG_PTR r14;
   ULONG_PTR r15;
#endif//_WIN64
   ULONG_PTR      cip;
   ULONG_PTR      eflags;
   unsigned short gs;
   unsigned short fs;
   unsigned short es;
   unsigned short ds;
   unsigned short cs;
   unsigned short ss;
   ULONG_PTR      dr0;
   ULONG_PTR      dr1;
   ULONG_PTR      dr2;
   ULONG_PTR      dr3;
   ULONG_PTR      dr6;
   ULONG_PTR      dr7;
   BYTE           RegisterArea[80];
   X87FPU         x87fpu;
   DWORD          MxCsr;
#ifdef _WIN64
   XMMREGISTER XmmRegisters[16];
   YMMREGISTER YmmRegisters[16];
#else// x86
   XMMREGISTER XmmRegisters[8];
   YMMREGISTER YmmRegisters[8];
#endif
} REGISTERCONTEXT;

typedef struct
{
   DWORD code;
   char  name[128];
} LASTERROR;

typedef struct
{
   DWORD code;
   char  name[128];
} LASTSTATUS;

typedef struct
{
   REGISTERCONTEXT      regcontext;
   FLAGS                flags;
   X87FPUREGISTER       x87FPURegisters[8];
   unsigned long long   mmx[8];
   MXCSRFIELDS          MxCsrFields;
   X87STATUSWORDFIELDS  x87StatusWordFields;
   X87CONTROLWORDFIELDS x87ControlWordFields;
   LASTERROR            lastError;
   LASTSTATUS           lastStatus;
} REGDUMP;

typedef struct
{
   DISASM_ARGTYPE type;//normal/memory
   SEGMENTREG     segment;
   char           mnemonic[64];
   duint          constant;//constant in the instruction _fake(imm/disp)
   duint          value;   //equal to constant or equal to the register value
   duint          memvalue;//memsize:[value]
} DISASM_ARG;

typedef struct
{
   char             instruction[64];
   DISASM_INSTRTYPE type;
   int              argcount;
   int              instr_size;
   DISASM_ARG       arg[3];
} DISASM_INSTR;

typedef struct
{
   char color[8];//hex color-code
   char comment[MAX_COMMENT_SIZE];
} STACK_COMMENT;

typedef struct
{
   int    ThreadNumber;
   HANDLE Handle;
   DWORD  ThreadId;
   duint  ThreadStartAddress;
   duint  ThreadLocalBase;
   char   threadName[MAX_THREAD_NAME_SIZE];
} THREADINFO;

typedef struct
{
   THREADINFO       BasicInfo;
   duint            ThreadCip;
   DWORD            SuspendCount;
   THREADPRIORITY   Priority;
   THREADWAITREASON WaitReason;
   DWORD            LastError;
   FILETIME         UserTime;
   FILETIME         KernelTime;
   FILETIME         CreationTime;
   ULONG64          Cycles;// Windows Vista or greater
} THREADALLINFO;

typedef struct
{
   int            count;
   THREADALLINFO *list;
   int            CurrentThread;
} THREADLIST;

typedef struct
{
   duint       value;//displacement / addrvalue _fake(rip-relative)
   MEMORY_SIZE size; //byte/word/dword/qword
   char        mnemonic[MAX_MNEMONIC_SIZE];
} MEMORY_INFO;

typedef struct
{
   duint      value;
   VALUE_SIZE size;
} VALUE_INFO;

#define TYPE_VALUE  1
#define TYPE_MEMORY 2
#define TYPE_ADDR   4

typedef struct
{
   DWORD       type; //value|memory|addr
   VALUE_INFO  value;//immediat
   MEMORY_INFO memory;
   duint       addr;  //addrvalue _fake(jumps + calls)
   bool        branch;//jumps/calls
   bool        call;  //instruction is a call
   int         size;
   char        instruction[MAX_MNEMONIC_SIZE * 4];
} BASIC_INSTRUCTION_INFO;

typedef struct
{
   SCRIPTBRANCHTYPE type;
   int              dest;
   char             branchlabel[256];
} SCRIPTBRANCH;

typedef struct
{
   duint addr;
   duint start;
   duint end;
   bool  manual;
   int   depth;
} FUNCTION_LOOP_INFO;

typedef struct
{
   duint    addr;
   XREFTYPE type;
} XREF_RECORD;

typedef struct
{
   duint        refcount;
   XREF_RECORD *references;
} XREF_INFO;

typedef struct
{
   duint address;
   duint from;
} XREF_EDGE;

typedef struct SYMBOLPTR_ {
   duint       modbase;
   const void *symbol;
} SYMBOLPTR;


   typedef enum {
       GUI_PLUGIN_MENU,
       GUI_DISASM_MENU,
       GUI_DUMP_MENU,
       GUI_STACK_MENU,
       GUI_GRAPH_MENU,
       GUI_MEMMAP_MENU,
       GUI_SYMMOD_MENU,
   } GUIMENUTYPE;
   typedef enum {
       GUI_DISASSEMBLY,
       GUI_DUMP,
       GUI_STACK,
       GUI_GRAPH,
       GUI_MEMMAP,
       GUI_SYMMOD,
       GUI_THREADS,
   } GUISELECTIONTYPE;

#define GUI_MAX_LINE_SIZE        65536
#define GUI_MAX_DISASSEMBLY_SIZE 2048

typedef enum {
   GUI_DISASSEMBLE_AT,             // param1=_fake(duint)va,            param2=_fake(duint)cip
   GUI_SET_DEBUG_STATE,            // param1=_fake(DBGSTATE)state,      param2=unused
   GUI_ADD_MSG_TO_LOG,             // param1=_fake(const char*)msg,     param2=unused
   GUI_CLEAR_LOG,                  // param1=unused,               param2=unused
   GUI_UPDATE_REGISTER_VIEW,       // param1=unused,               param2=unused
   GUI_UPDATE_DISASSEMBLY_VIEW,    // param1=unused,               param2=unused
   GUI_UPDATE_BREAKPOINTS_VIEW,    // param1=unused,               param2=unused
   GUI_UPDATE_WINDOW_TITLE,        // param1=_fake(const char*)file,    param2=unused
   GUI_GET_WINDOW_HANDLE,          // param1=unused,               param2=unused
   GUI_DUMP_AT,                    // param1=_fake(duint)va             param2=unused
   GUI_SCRIPT_ADD,                 // param1=int count,            param2=const char** lines
   GUI_SCRIPT_CLEAR,               // param1=unused,               param2=unused
   GUI_SCRIPT_SETIP,               // param1=int line,             param2=unused
   GUI_SCRIPT_ERROR,               // param1=int line,             param2=std::string message
   GUI_SCRIPT_SETTITLE,            // param1=std::string title,    param2=unused
   GUI_SCRIPT_SETINFOLINE,         // param1=int line,             param2=std::string info
   GUI_SCRIPT_MESSAGE,             // param1=std::string message,  param2=unused
   GUI_SCRIPT_MSGYN,               // param1=std::string message,  param2=unused
   GUI_SYMBOL_LOG_ADD,             // param1_fake(const char*)msg,      param2=unused
   GUI_SYMBOL_LOG_CLEAR,           // param1=unused,               param2=unused
   GUI_SYMBOL_SET_PROGRESS,        // param1=int percent           param2=unused
   GUI_SYMBOL_UPDATE_MODULE_LIST,  // param1=int count,            param2=SYMBOLMODULEINFO* modules
   GUI_REF_ADDCOLUMN,              // param1=int width,            param2=_fake(const char*)title
   GUI_REF_SETROWCOUNT,            // param1=int rows,             param2=unused
   GUI_REF_GETROWCOUNT,            // param1=unused,               param2=unused
   GUI_REF_DELETEALLCOLUMNS,       // param1=unused,               param2=unused
   GUI_REF_SETCELLCONTENT,         // param1=_fake(CELLINFO*)info,      param2=unused
   GUI_REF_GETCELLCONTENT,         // param1=int row,              param2=int col
   GUI_REF_RELOADDATA,             // param1=unused,               param2=unused
   GUI_REF_SETSINGLESELECTION,     // param1=int index,            param2=bool scroll
   GUI_REF_SETPROGRESS,            // param1=int progress,         param2=unused
   GUI_REF_SETCURRENTTASKPROGRESS, // param1=int progress,         param2=std::string taskTitle
   GUI_REF_SETSEARCHSTARTCOL,      // param1=int col               param2=unused
   GUI_STACK_DUMP_AT,              // param1=duint addr,           param2=duint csp
   GUI_UPDATE_DUMP_VIEW,           // param1=unused,               param2=unused
   GUI_UPDATE_THREAD_VIEW,         // param1=unused,               param2=unused
   GUI_ADD_RECENT_FILE,            // param1=_fake(const char*)file,    param2=unused
   GUI_SET_LAST_EXCEPTION,         // param1=unsigned int code,    param2=unused
   GUI_GET_DISASSEMBLY,            // param1=duint addr,           param2=char* text
   GUI_MENU_ADD,                   // param1=int hMenu,            param2=std::string title
   GUI_MENU_ADD_ENTRY,             // param1=int hMenu,            param2=std::string title
   GUI_MENU_ADD_SEPARATOR,         // param1=int hMenu,            param2=unused
   GUI_MENU_CLEAR,                 // param1=int hMenu,            param2=unused
   GUI_SELECTION_GET,              // param1=GUISELECTIONTYPE,     param2=SELECTIONDATA* selection
   GUI_SELECTION_SET,              // param1=GUISELECTIONTYPE,     param2=const SELECTIONDATA* selection
   GUI_GETLINE_WINDOW,             // param1=std::string title,    param2=char* text
   GUI_AUTOCOMPLETE_ADDCMD,        // param1=std::string cmd,      param2=ununsed
   GUI_AUTOCOMPLETE_DELCMD,        // param1=std::string cmd,      param2=ununsed
   GUI_AUTOCOMPLETE_CLEARALL,      // param1=unused,               param2=unused
   GUI_SCRIPT_ENABLEHIGHLIGHTING,  // param1=bool enable,          param2=unused
   GUI_ADD_MSG_TO_STATUSBAR,       // param1=std::string msg,      param2=unused
   GUI_UPDATE_SIDEBAR,             // param1=unused,               param2=unused
   GUI_REPAINT_TABLE_VIEW,         // param1=unused,               param2=unused
   GUI_UPDATE_PATCHES,             // param1=unused,               param2=unused
   GUI_UPDATE_CALLSTACK,           // param1=unused,               param2=unused
   GUI_UPDATE_SEHCHAIN,            // param1=unused,               param2=unused
   GUI_SYMBOL_REFRESH_CURRENT,     // param1=unused,               param2=unused
   GUI_UPDATE_MEMORY_VIEW,         // param1=unused,               param2=unused
   GUI_REF_INITIALIZE,             // param1=std::string name,     param2=unused
   GUI_LOAD_SOURCE_FILE,           // param1=std::string path,     param2=duint addr
   GUI_MENU_SET_ICON,              // param1=int hMenu,            param2=ICONINFO*
   GUI_MENU_SET_ENTRY_ICON,        // param1=int hEntry,           param2=ICONINFO*
   GUI_SHOW_CPU,                   // param1=unused,               param2=unused
   GUI_ADD_QWIDGET_TAB,            // param1=QWidget*,             param2=unused
   GUI_SHOW_QWIDGET_TAB,           // param1=QWidget*,             param2=unused
   GUI_CLOSE_QWIDGET_TAB,          // param1=QWidget*,             param2=unused
   GUI_EXECUTE_ON_GUI_THREAD,      // param1=GUICALLBACKEX cb,     param2=void* userdata
   GUI_UPDATE_TIME_WASTED_COUNTER, // param1=unused,               param2=unused
   GUI_SET_GLOBAL_NOTES,           // param1=std::string text,     param2=unused
   GUI_GET_GLOBAL_NOTES,           // param1=char** text,          param2=unused
   GUI_SET_DEBUGGEE_NOTES,         // param1=std::string text,     param2=unused
   GUI_GET_DEBUGGEE_NOTES,         // param1=char** text,          param2=unused
   GUI_DUMP_AT_N,                  // param1=int index,            param2=duint va
   GUI_DISPLAY_WARNING,            // param1=std::string text,     param2=unused
   GUI_REGISTER_SCRIPT_LANG,       // param1=SCRIPTTYPEINFO* info, param2=unused
   GUI_UNREGISTER_SCRIPT_LANG,     // param1=int id,               param2=unused
   GUI_UPDATE_ARGUMENT_VIEW,       // param1=unused,               param2=unused
   GUI_FOCUS_VIEW,                 // param1=int hWindow,          param2=unused
   GUI_UPDATE_WATCH_VIEW,          // param1=unused,               param2=unused
   GUI_LOAD_GRAPH,                 // param1=BridgeCFGraphList*    param2=unused
   GUI_GRAPH_AT,                   // param1=duint addr            param2=unused
   GUI_UPDATE_GRAPH_VIEW,          // param1=unused,               param2=unused
   GUI_SET_LOG_ENABLED,            // param1=bool isEnabled        param2=unused
   GUI_ADD_FAVOURITE_TOOL,         // param1=std::string name      param2=std::string description
   GUI_ADD_FAVOURITE_COMMAND,      // param1=std::string command   param2=std::string shortcut
   GUI_SET_FAVOURITE_TOOL_SHORTCUT,// param1=std::string name      param2=std::string shortcut
   GUI_FOLD_DISASSEMBLY,           // param1=duint startAddress    param2=duint length
   GUI_SELECT_IN_MEMORY_MAP,       // param1=duint addr,           param2=unused
   GUI_GET_ACTIVE_VIEW,            // param1=ACTIVEVIEW*,          param2=unused
   GUI_MENU_SET_ENTRY_CHECKED,     // param1=int hEntry,           param2=bool checked
   GUI_ADD_INFO_LINE,              // param1=std::string infoline, param2=unused
   GUI_PROCESS_EVENTS,             // param1=unused,               param2=unused
   GUI_TYPE_ADDNODE,               // param1=void* parent,         param2=TYPEDESCRIPTOR* type
   GUI_TYPE_CLEAR,                 // param1=unused,               param2=unused
   GUI_UPDATE_TYPE_WIDGET,         // param1=unused,               param2=unused
   GUI_CLOSE_APPLICATION,          // param1=unused,               param2=unused
   GUI_MENU_SET_VISIBLE,           // param1=int hMenu,            param2=bool visible
   GUI_MENU_SET_ENTRY_VISIBLE,     // param1=int hEntry,           param2=bool visible
   GUI_MENU_SET_NAME,              // param1=int hMenu,            param2=std::string name
   GUI_MENU_SET_ENTRY_NAME,        // param1=int hEntry,           param2=std::string name
   GUI_FLUSH_LOG,                  // param1=unused,               param2=unused
   GUI_MENU_SET_ENTRY_HOTKEY,      // param1=int hEntry,           param2=std::string hack
   GUI_REF_SEARCH_GETROWCOUNT,     // param1=unused,               param2=unused
   GUI_REF_SEARCH_GETCELLCONTENT,  // param1=int row,              param2=int col
   GUI_MENU_REMOVE,                // param1=int hEntryMenu,       param2=unused
   GUI_REF_ADDCOMMAND,             // param1=std::string title,    param2=std::string command
   GUI_OPEN_TRACE_FILE,            // param1=std::string file name,param2=unused
   GUI_UPDATE_TRACE_BROWSER,       // param1=unused,               param2=unused
   GUI_INVALIDATE_SYMBOL_SOURCE,   // param1=duint base,           param2=unused
   GUI_GET_CURRENT_GRAPH,          // param1=BridgeCFGraphList*,   param2=unused
   GUI_SHOW_REF,                   // param1=unused,               param2=unused
   GUI_SELECT_IN_SYMBOLS_TAB,      // param1=duint addr,           param2=unused
   GUI_GOTO_TRACE,                 // param1=duint index,          param2=unused
   GUI_SHOW_TRACE,                 // param1=unused,               param2=unused
   GUI_GET_MAIN_THREAD_ID,         // param1=unused,               param2=unused
   GUI_ADD_MSG_TO_LOG_HTML,        // param1=_fake(const char*)msg,     param2=unused
   GUI_IS_LOG_ENABLED,             // param1=unused,               param2=unused
   GUI_IS_DEBUGGER_FOCUSED_UNUSED, // This message is removed, could be used for future purposes
   GUI_SAVE_LOG,                   // param1=std::string file name,param2=unused
   GUI_REDIRECT_LOG,               // param1=std::string file name,param2=unused
   GUI_STOP_REDIRECT_LOG,          // param1=unused,               param2=unused
   GUI_SHOW_THREADS,               // param1=unused,               param2=unused
} GUIMSG;

struct _TYPEDESCRIPTOR;

typedef struct
{
   int         row;
   int         col;
   std::string str;
} CELLINFO;

typedef struct
{
   char               name[64];
   int                id;
   GUISCRIPTEXECUTE   execute;
   GUISCRIPTCOMPLETER completeCommand;
} SCRIPTTYPEINFO;

typedef struct
{
   void *titleHwnd;
   void *classHwnd;
   char  title[MAX_STRING_SIZE];
   char  className[MAX_STRING_SIZE];
} ACTIVEVIEW;

typedef struct _TYPEDESCRIPTOR {
   bool         expanded;//is the type node expanded?
   bool         reverse; //big endian?
   const char  *name;    //type name _fake(int b)
   duint        addr;    //virtual address
   duint        offset;  //offset to addr for the actual location
   int          id;      //type id
   int          size;    //sizeof_fake(type)
   TYPETOSTRING callback;//convert to string
   void        *userdata;//user data
} TYPEDESCRIPTOR;


*/