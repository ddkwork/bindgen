package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type DBGSTATE int // D:\workspace\workspace\mcp\x64dbg\include\bridgemain.h:113
const (
	initialized DBGSTATE = iota // 0
	paused                      // 1
	running                     // 2
	stopped                     // 3
)

type SEGMENTREG int // :121
const (
	SEG_DEFAULT SEGMENTREG = iota // 0
	SEG_ES                        // 1
	SEG_DS                        // 2
	SEG_FS                        // 3
	SEG_GS                        // 4
	SEG_CS                        // 5
	SEG_SS                        // 6
)

type ADDRINFOFLAGS int // :132
const (
	flagmodule       ADDRINFOFLAGS = iota // 1
	flaglabel                      = 2    // 2
	flagcomment                    = 4    // 4
	flagbookmark                   = 8    // 8
	flagfunction                   = 16   // 16
	flagloop                       = 32   // 32
	flagargs                       = 64   // 64
	flagNoFuncOffset               = 128  // 128
)

type BPXTYPE int // :144
const (
	bp_none      BPXTYPE = iota // 0
	bp_normal            = 1    // 1
	bp_hardware          = 2    // 2
	bp_memory            = 4    // 4
	bp_dll               = 8    // 8
	bp_exception         = 16   // 16
)

type FUNCTYPE int // :154
const (
	FUNC_NONE   FUNCTYPE = iota // 0
	FUNC_BEGIN                  // 1
	FUNC_MIDDLE                 // 2
	FUNC_END                    // 3
	FUNC_SINGLE                 // 4
)

type LOOPTYPE int // :163
const (
	LOOP_NONE   LOOPTYPE = iota // 0
	LOOP_BEGIN                  // 1
	LOOP_MIDDLE                 // 2
	LOOP_ENTRY                  // 3
	LOOP_END                    // 4
	LOOP_SINGLE                 // 5
)

type XREFTYPE int // :173
const (
	XREF_NONE XREFTYPE = iota // 0
	XREF_DATA                 // 1
	XREF_JMP                  // 2
	XREF_CALL                 // 3
)

type ARGTYPE int // :181
const (
	ARG_NONE   ARGTYPE = iota // 0
	ARG_BEGIN                 // 1
	ARG_MIDDLE                // 2
	ARG_END                   // 3
	ARG_SINGLE                // 4
)

type DBGMSG int // :190
const (
	DBG_SCRIPT_LOAD                DBGMSG = iota // 0
	DBG_SCRIPT_UNLOAD                            // 1
	DBG_SCRIPT_RUN                               // 2
	DBG_SCRIPT_STEP                              // 3
	DBG_SCRIPT_BPTOGGLE                          // 4
	DBG_SCRIPT_BPGET                             // 5
	DBG_SCRIPT_CMDEXEC                           // 6
	DBG_SCRIPT_ABORT                             // 7
	DBG_SCRIPT_GETLINETYPE                       // 8
	DBG_SCRIPT_SETIP                             // 9
	DBG_SCRIPT_GETBRANCHINFO                     // 10
	DBG_SYMBOL_ENUM                              // 11
	DBG_ASSEMBLE_AT                              // 12
	DBG_MODBASE_FROM_NAME                        // 13
	DBG_DISASM_AT                                // 14
	DBG_STACK_COMMENT_GET                        // 15
	DBG_GET_THREAD_LIST                          // 16
	DBG_SETTINGS_UPDATED                         // 17
	DBG_DISASM_FAST_AT                           // 18
	DBG_MENU_ENTRY_CLICKED                       // 19
	DBG_FUNCTION_GET                             // 20
	DBG_FUNCTION_OVERLAPS                        // 21
	DBG_FUNCTION_ADD                             // 22
	DBG_FUNCTION_DEL                             // 23
	DBG_LOOP_GET                                 // 24
	DBG_LOOP_OVERLAPS                            // 25
	DBG_LOOP_ADD                                 // 26
	DBG_LOOP_DEL                                 // 27
	DBG_IS_RUN_LOCKED                            // 28
	DBG_IS_BP_DISABLED                           // 29
	DBG_SET_AUTO_COMMENT_AT                      // 30
	DBG_DELETE_AUTO_COMMENT_RANGE                // 31
	DBG_SET_AUTO_LABEL_AT                        // 32
	DBG_DELETE_AUTO_LABEL_RANGE                  // 33
	DBG_SET_AUTO_BOOKMARK_AT                     // 34
	DBG_DELETE_AUTO_BOOKMARK_RANGE               // 35
	DBG_SET_AUTO_FUNCTION_AT                     // 36
	DBG_DELETE_AUTO_FUNCTION_RANGE               // 37
	DBG_GET_STRING_AT                            // 38
	DBG_GET_FUNCTIONS                            // 39
	DBG_WIN_EVENT                                // 40
	DBG_WIN_EVENT_GLOBAL                         // 41
	DBG_INITIALIZE_LOCKS                         // 42
	DBG_DEINITIALIZE_LOCKS                       // 43
	DBG_GET_TIME_WASTED_COUNTER                  // 44
	DBG_SYMBOL_ENUM_FROMCACHE                    // 45
	DBG_DELETE_COMMENT_RANGE                     // 46
	DBG_DELETE_LABEL_RANGE                       // 47
	DBG_DELETE_BOOKMARK_RANGE                    // 48
	DBG_GET_XREF_COUNT_AT                        // 49
	DBG_GET_XREF_TYPE_AT                         // 50
	DBG_XREF_ADD                                 // 51
	DBG_XREF_DEL_ALL                             // 52
	DBG_XREF_GET                                 // 53
	DBG_GET_ENCODE_TYPE_BUFFER                   // 54
	DBG_ENCODE_TYPE_GET                          // 55
	DBG_DELETE_ENCODE_TYPE_RANGE                 // 56
	DBG_ENCODE_SIZE_GET                          // 57
	DBG_DELETE_ENCODE_TYPE_SEG                   // 58
	DBG_RELEASE_ENCODE_TYPE_BUFFER               // 59
	DBG_ARGUMENT_GET                             // 60
	DBG_ARGUMENT_OVERLAPS                        // 61
	DBG_ARGUMENT_ADD                             // 62
	DBG_ARGUMENT_DEL                             // 63
	DBG_GET_WATCH_LIST                           // 64
	DBG_SELCHANGED                               // 65
	DBG_GET_PROCESS_HANDLE                       // 66
	DBG_GET_THREAD_HANDLE                        // 67
	DBG_GET_PROCESS_ID                           // 68
	DBG_GET_THREAD_ID                            // 69
	DBG_GET_PEB_ADDRESS                          // 70
	DBG_GET_TEB_ADDRESS                          // 71
	DBG_ANALYZE_FUNCTION                         // 72
	DBG_MENU_PREPARE                             // 73
	DBG_GET_SYMBOL_INFO                          // 74
	DBG_GET_DEBUG_ENGINE                         // 75
	DBG_GET_SYMBOL_INFO_AT                       // 76
	DBG_XREF_ADD_MULTI                           // 77
)

type SCRIPTLINETYPE int // :272
const (
	linecommand SCRIPTLINETYPE = iota // 0
	linebranch                        // 1
	linelabel                         // 2
	linecomment                       // 3
	lineempty                         // 4
)

type SCRIPTBRANCHTYPE int // :281
const (
	scriptnobranch SCRIPTBRANCHTYPE = iota // 0
	scriptjmp                              // 1
	scriptjnejnz                           // 2
	scriptjejz                             // 3
	scriptjbjl                             // 4
	scriptjajg                             // 5
	scriptjbejle                           // 6
	scriptjaejge                           // 7
	scriptcall                             // 8
)

type DISASM_INSTRTYPE int // :294
const (
	instr_normal DISASM_INSTRTYPE = iota // 0
	instr_branch                         // 1
	instr_stack                          // 2
)

type DISASM_ARGTYPE int // :301
const (
	arg_normal DISASM_ARGTYPE = iota // 0
	arg_memory                       // 1
)

type STRING_TYPE int // :307
const (
	str_none    STRING_TYPE = iota // 0
	str_ascii                      // 1
	str_unicode                    // 2
)

type THREADPRIORITY int // :314
const (
	_PriorityIdle         THREADPRIORITY = iota       // -15
	_PriorityAboveNormal                 = 1          // 1
	_PriorityBelowNormal                 = -1         // -1
	_PriorityHighest                     = 2          // 2
	_PriorityLowest                      = -2         // -2
	_PriorityNormal                      = 0          // 0
	_PriorityTimeCritical                = 15         // 15
	_PriorityUnknown                     = 2147483647 // 2147483647
)

type THREADWAITREASON int // :326
const (
	_Executive        THREADWAITREASON = iota // 0
	_FreePage                          = 1    // 1
	_PageIn                            = 2    // 2
	_PoolAllocation                    = 3    // 3
	_DelayExecution                    = 4    // 4
	_Suspended                         = 5    // 5
	_UserRequest                       = 6    // 6
	_WrExecutive                       = 7    // 7
	_WrFreePage                        = 8    // 8
	_WrPageIn                          = 9    // 9
	_WrPoolAllocation                  = 10   // 10
	_WrDelayExecution                  = 11   // 11
	_WrSuspended                       = 12   // 12
	_WrUserRequest                     = 13   // 13
	_WrEventPair                       = 14   // 14
	_WrQueue                           = 15   // 15
	_WrLpcReceive                      = 16   // 16
	_WrLpcReply                        = 17   // 17
	_WrVirtualMemory                   = 18   // 18
	_WrPageOut                         = 19   // 19
	_WrRendezvous                      = 20   // 20
	_Spare2                            = 21   // 21
	_Spare3                            = 22   // 22
	_Spare4                            = 23   // 23
	_Spare5                            = 24   // 24
	_WrCalloutStack                    = 25   // 25
	_WrKernel                          = 26   // 26
	_WrResource                        = 27   // 27
	_WrPushLock                        = 28   // 28
	_WrMutex                           = 29   // 29
	_WrQuantumEnd                      = 30   // 30
	_WrDispatchInt                     = 31   // 31
	_WrPreempted                       = 32   // 32
	_WrYieldExecution                  = 33   // 33
	_WrFastMutex                       = 34   // 34
	_WrGuardedMutex                    = 35   // 35
	_WrRundown                         = 36   // 36
)

type MEMORY_SIZE int // :367
const (
	size_byte    MEMORY_SIZE = iota // 1
	size_word                = 2    // 2
	size_dword               = 4    // 4
	size_qword               = 8    // 8
	size_xmmword             = 16   // 16
	size_ymmword             = 32   // 32
)

type ENCODETYPE int // :377
const (
	enc_unknown ENCODETYPE = iota // 0
	enc_byte                      // 1
	enc_word                      // 2
	enc_dword                     // 3
	enc_fword                     // 4
	enc_qword                     // 5
	enc_tbyte                     // 6
	enc_oword                     // 7
	enc_mmword                    // 8
	enc_xmmword                   // 9
	enc_ymmword                   // 10
	enc_zmmword                   // 11
	enc_real4                     // 12
	enc_real8                     // 13
	enc_real10                    // 14
	enc_ascii                     // 15
	enc_unicode                   // 16
	enc_code                      // 17
	enc_junk                      // 18
	enc_middle                    // 19
)

type WATCHVARTYPE int // :401
const (
	TYPE_UINT    WATCHVARTYPE = iota // 0
	TYPE_INT                         // 1
	TYPE_FLOAT                       // 2
	TYPE_ASCII                       // 3
	TYPE_UNICODE                     // 4
	TYPE_INVALID                     // 5
)

type WATCHDOGMODE int // :411
const (
	MODE_DISABLED  WATCHDOGMODE = iota // 0
	MODE_ISTRUE                        // 1
	MODE_ISFALSE                       // 2
	MODE_CHANGED                       // 3
	MODE_UNCHANGED                     // 4
)

type BPHWTYPE int // :420
const (
	hw_access  BPHWTYPE = iota // 0
	hw_write                   // 1
	hw_execute                 // 2
)

type BPMEMTYPE int // :427
const (
	mem_access  BPMEMTYPE = iota // 0
	mem_read                     // 1
	mem_write                    // 2
	mem_execute                  // 3
)

type BPDLLTYPE int // :435
const (
	dll_load   BPDLLTYPE = iota // 1
	dll_unload                  // 2
	dll_all                     // 3
)

type BPEXTYPE int // :442
const (
	ex_firstchance  BPEXTYPE = iota // 1
	ex_secondchance                 // 2
	ex_all                          // 3
)

type BPHWSIZE int // :449
const (
	hw_byte  BPHWSIZE = iota // 0
	hw_word                  // 1
	hw_dword                 // 2
	hw_qword                 // 3
)

type SYMBOLTYPE int // :457
const (
	sym_import SYMBOLTYPE = iota // 0
	sym_export                   // 1
	sym_symbol                   // 2
)

type MODULEPARTY int // :469
const (
	mod_user   MODULEPARTY = iota // 0
	mod_system                    // 1
)

type DEBUG_ENGINE int // :475
const (
	DebugEngineTitanEngine  DEBUG_ENGINE = iota // 0
	DebugEngineGleeBug                          // 1
	DebugEngineStaticEngine                     // 2
)

type GUIMENUTYPE int // :1032
const (
	GUI_PLUGIN_MENU GUIMENUTYPE = iota // 0
	GUI_DISASM_MENU                    // 1
	GUI_DUMP_MENU                      // 2
	GUI_STACK_MENU                     // 3
	GUI_GRAPH_MENU                     // 4
	GUI_MEMMAP_MENU                    // 5
	GUI_SYMMOD_MENU                    // 6
)

type GUISELECTIONTYPE int // :1045
const (
	GUI_DISASSEMBLY GUISELECTIONTYPE = iota // 0
	GUI_DUMP                                // 1
	GUI_STACK                               // 2
	GUI_GRAPH                               // 3
	GUI_MEMMAP                              // 4
	GUI_SYMMOD                              // 5
	GUI_THREADS                             // 6
)

type GUIMSG int // :1059
const (
	GUI_DISASSEMBLE_AT              GUIMSG = iota // 0
	GUI_SET_DEBUG_STATE                           // 1
	GUI_ADD_MSG_TO_LOG                            // 2
	GUI_CLEAR_LOG                                 // 3
	GUI_UPDATE_REGISTER_VIEW                      // 4
	GUI_UPDATE_DISASSEMBLY_VIEW                   // 5
	GUI_UPDATE_BREAKPOINTS_VIEW                   // 6
	GUI_UPDATE_WINDOW_TITLE                       // 7
	GUI_GET_WINDOW_HANDLE                         // 8
	GUI_DUMP_AT                                   // 9
	GUI_SCRIPT_ADD                                // 10
	GUI_SCRIPT_CLEAR                              // 11
	GUI_SCRIPT_SETIP                              // 12
	GUI_SCRIPT_ERROR                              // 13
	GUI_SCRIPT_SETTITLE                           // 14
	GUI_SCRIPT_SETINFOLINE                        // 15
	GUI_SCRIPT_MESSAGE                            // 16
	GUI_SCRIPT_MSGYN                              // 17
	GUI_SYMBOL_LOG_ADD                            // 18
	GUI_SYMBOL_LOG_CLEAR                          // 19
	GUI_SYMBOL_SET_PROGRESS                       // 20
	GUI_SYMBOL_UPDATE_MODULE_LIST                 // 21
	GUI_REF_ADDCOLUMN                             // 22
	GUI_REF_SETROWCOUNT                           // 23
	GUI_REF_GETROWCOUNT                           // 24
	GUI_REF_DELETEALLCOLUMNS                      // 25
	GUI_REF_SETCELLCONTENT                        // 26
	GUI_REF_GETCELLCONTENT                        // 27
	GUI_REF_RELOADDATA                            // 28
	GUI_REF_SETSINGLESELECTION                    // 29
	GUI_REF_SETPROGRESS                           // 30
	GUI_REF_SETCURRENTTASKPROGRESS                // 31
	GUI_REF_SETSEARCHSTARTCOL                     // 32
	GUI_STACK_DUMP_AT                             // 33
	GUI_UPDATE_DUMP_VIEW                          // 34
	GUI_UPDATE_THREAD_VIEW                        // 35
	GUI_ADD_RECENT_FILE                           // 36
	GUI_SET_LAST_EXCEPTION                        // 37
	GUI_GET_DISASSEMBLY                           // 38
	GUI_MENU_ADD                                  // 39
	GUI_MENU_ADD_ENTRY                            // 40
	GUI_MENU_ADD_SEPARATOR                        // 41
	GUI_MENU_CLEAR                                // 42
	GUI_SELECTION_GET                             // 43
	GUI_SELECTION_SET                             // 44
	GUI_GETLINE_WINDOW                            // 45
	GUI_AUTOCOMPLETE_ADDCMD                       // 46
	GUI_AUTOCOMPLETE_DELCMD                       // 47
	GUI_AUTOCOMPLETE_CLEARALL                     // 48
	GUI_SCRIPT_ENABLEHIGHLIGHTING                 // 49
	GUI_ADD_MSG_TO_STATUSBAR                      // 50
	GUI_UPDATE_SIDEBAR                            // 51
	GUI_REPAINT_TABLE_VIEW                        // 52
	GUI_UPDATE_PATCHES                            // 53
	GUI_UPDATE_CALLSTACK                          // 54
	GUI_UPDATE_SEHCHAIN                           // 55
	GUI_SYMBOL_REFRESH_CURRENT                    // 56
	GUI_UPDATE_MEMORY_VIEW                        // 57
	GUI_REF_INITIALIZE                            // 58
	GUI_LOAD_SOURCE_FILE                          // 59
	GUI_MENU_SET_ICON                             // 60
	GUI_MENU_SET_ENTRY_ICON                       // 61
	GUI_SHOW_CPU                                  // 62
	GUI_ADD_QWIDGET_TAB                           // 63
	GUI_SHOW_QWIDGET_TAB                          // 64
	GUI_CLOSE_QWIDGET_TAB                         // 65
	GUI_EXECUTE_ON_GUI_THREAD                     // 66
	GUI_UPDATE_TIME_WASTED_COUNTER                // 67
	GUI_SET_GLOBAL_NOTES                          // 68
	GUI_GET_GLOBAL_NOTES                          // 69
	GUI_SET_DEBUGGEE_NOTES                        // 70
	GUI_GET_DEBUGGEE_NOTES                        // 71
	GUI_DUMP_AT_N                                 // 72
	GUI_DISPLAY_WARNING                           // 73
	GUI_REGISTER_SCRIPT_LANG                      // 74
	GUI_UNREGISTER_SCRIPT_LANG                    // 75
	GUI_UPDATE_ARGUMENT_VIEW                      // 76
	GUI_FOCUS_VIEW                                // 77
	GUI_UPDATE_WATCH_VIEW                         // 78
	GUI_LOAD_GRAPH                                // 79
	GUI_GRAPH_AT                                  // 80
	GUI_UPDATE_GRAPH_VIEW                         // 81
	GUI_SET_LOG_ENABLED                           // 82
	GUI_ADD_FAVOURITE_TOOL                        // 83
	GUI_ADD_FAVOURITE_COMMAND                     // 84
	GUI_SET_FAVOURITE_TOOL_SHORTCUT               // 85
	GUI_FOLD_DISASSEMBLY                          // 86
	GUI_SELECT_IN_MEMORY_MAP                      // 87
	GUI_GET_ACTIVE_VIEW                           // 88
	GUI_MENU_SET_ENTRY_CHECKED                    // 89
	GUI_ADD_INFO_LINE                             // 90
	GUI_PROCESS_EVENTS                            // 91
	GUI_TYPE_ADDNODE                              // 92
	GUI_TYPE_CLEAR                                // 93
	GUI_UPDATE_TYPE_WIDGET                        // 94
	GUI_CLOSE_APPLICATION                         // 95
	GUI_MENU_SET_VISIBLE                          // 96
	GUI_MENU_SET_ENTRY_VISIBLE                    // 97
	GUI_MENU_SET_NAME                             // 98
	GUI_MENU_SET_ENTRY_NAME                       // 99
	GUI_FLUSH_LOG                                 // 100
	GUI_MENU_SET_ENTRY_HOTKEY                     // 101
	GUI_REF_SEARCH_GETROWCOUNT                    // 102
	GUI_REF_SEARCH_GETCELLCONTENT                 // 103
	GUI_MENU_REMOVE                               // 104
	GUI_REF_ADDCOMMAND                            // 105
	GUI_OPEN_TRACE_FILE                           // 106
	GUI_UPDATE_TRACE_BROWSER                      // 107
	GUI_INVALIDATE_SYMBOL_SOURCE                  // 108
	GUI_GET_CURRENT_GRAPH                         // 109
	GUI_SHOW_REF                                  // 110
	GUI_SELECT_IN_SYMBOLS_TAB                     // 111
	GUI_GOTO_TRACE                                // 112
	GUI_SHOW_TRACE                                // 113
	GUI_GET_MAIN_THREAD_ID                        // 114
	GUI_ADD_MSG_TO_LOG_HTML                       // 115
	GUI_IS_LOG_ENABLED                            // 116
	GUI_IS_DEBUGGER_FOCUSED_UNUSED                // 117
	GUI_SAVE_LOG                                  // 118
	GUI_REDIRECT_LOG                              // 119
	GUI_STOP_REDIRECT_LOG                         // 120
	GUI_SHOW_THREADS                              // 121
)

// ListInfo (D:\workspace\workspace\mcp\x64dbg\include\bridgelist.h:4)
type ListInfo struct {
	Count int      // C type: int
	Size  uint     // C type: size_t
	Data  *uintptr // C type: void *
}

// BridgeCFInstruction (D:\workspace\workspace\mcp\x64dbg\include\bridgegraph.h:4)
type BridgeCFInstruction struct {
	Addr uint     // C type: duint
	Data [15]byte // C type: unsigned char[15]
}

// BridgeCFGraphList (:26)
type BridgeCFGraphList struct {
	EntryPoint uint     // C type: duint
	Userdata   *uintptr // C type: void *
	Nodes      ListInfo // C type: ListInfo
}

// MEMPAGE (:488)
type MEMPAGE struct {
	Mbi  uintptr   // C type: MEMORY_BASIC_INFORMATION
	Info [256]int8 // C type: char[256]
}

// MEMMAP (:494)
type MEMMAP struct {
	Count int      // C type: int
	Page  *MEMPAGE // C type: MEMPAGE *
}

// BPMAP (:523)
type BPMAP struct {
	Count int      // C type: int
	Bp    *uintptr // C type: BRIDGEBP *
}

// FUNCTION (:541)
type FUNCTION struct {
	Start      uint // C type: duint
	End        uint // C type: duint
	Instrcount uint // C type: duint
}

// LOOP (:548)
type LOOP struct {
	Depth      int  // C type: int
	Start      uint // C type: duint
	End        uint // C type: duint
	Instrcount uint // C type: duint
}

// SYMBOLMODULEINFO (:609)
type SYMBOLMODULEINFO struct {
	Base uint      // C type: duint
	Name [256]int8 // C type: char[256]
}

// SYMBOLCBINFO (:615)
type SYMBOLCBINFO struct {
	Base         uint     // C type: duint
	CbSymbolEnum uintptr  // C type: CBSYMBOLENUM
	User         *uintptr // C type: void *
	Start        uint     // C type: duint
	End          uint     // C type: duint
	SymbolMask   uint     // C type: unsigned int
}

// YMMREGISTER (:700)
type YMMREGISTER struct {
	Low  uintptr // C type: XMMREGISTER
	High uintptr // C type: XMMREGISTER
}

// X87FPUREGISTER (:706)
type X87FPUREGISTER struct {
	Data     [10]byte // C type: BYTE[10]
	St_value int      // C type: int
	Tag      int      // C type: int
}

// X87FPU (:713)
type X87FPU struct {
	ControlWord   uint16 // C type: WORD
	StatusWord    uint16 // C type: WORD
	TagWord       uint16 // C type: WORD
	ErrorOffset   uint32 // C type: DWORD
	ErrorSelector uint32 // C type: DWORD
	DataOffset    uint32 // C type: DWORD
	DataSelector  uint32 // C type: DWORD
	Cr0NpxState   uint32 // C type: DWORD
}

// REGISTERCONTEXT (:725)
type REGISTERCONTEXT struct {
	Cax          uintptr         // C type: ULONG_PTR
	Ccx          uintptr         // C type: ULONG_PTR
	Cdx          uintptr         // C type: ULONG_PTR
	Cbx          uintptr         // C type: ULONG_PTR
	Csp          uintptr         // C type: ULONG_PTR
	Cbp          uintptr         // C type: ULONG_PTR
	Csi          uintptr         // C type: ULONG_PTR
	Cdi          uintptr         // C type: ULONG_PTR
	R8           uintptr         // C type: ULONG_PTR
	R9           uintptr         // C type: ULONG_PTR
	R10          uintptr         // C type: ULONG_PTR
	R11          uintptr         // C type: ULONG_PTR
	R12          uintptr         // C type: ULONG_PTR
	R13          uintptr         // C type: ULONG_PTR
	R14          uintptr         // C type: ULONG_PTR
	R15          uintptr         // C type: ULONG_PTR
	Cip          uintptr         // C type: ULONG_PTR
	Eflags       uintptr         // C type: ULONG_PTR
	Gs           uint16          // C type: unsigned short
	Fs           uint16          // C type: unsigned short
	Es           uint16          // C type: unsigned short
	Ds           uint16          // C type: unsigned short
	Cs           uint16          // C type: unsigned short
	Ss           uint16          // C type: unsigned short
	Dr0          uintptr         // C type: ULONG_PTR
	Dr1          uintptr         // C type: ULONG_PTR
	Dr2          uintptr         // C type: ULONG_PTR
	Dr3          uintptr         // C type: ULONG_PTR
	Dr6          uintptr         // C type: ULONG_PTR
	Dr7          uintptr         // C type: ULONG_PTR
	RegisterArea [80]byte        // C type: BYTE[80]
	X87fpu       X87FPU          // C type: X87FPU
	MxCsr        uint32          // C type: DWORD
	XmmRegisters [16]uintptr     // C type: XMMREGISTER[16]
	YmmRegisters [16]YMMREGISTER // C type: YMMREGISTER[16]
}

// LASTERROR (:771)
type LASTERROR struct {
	Code uint32    // C type: DWORD
	Name [128]int8 // C type: char[128]
}

// LASTSTATUS (:777)
type LASTSTATUS struct {
	Code uint32    // C type: DWORD
	Name [128]int8 // C type: char[128]
}

// REGDUMP (:783)
type REGDUMP struct {
	Regcontext           REGISTERCONTEXT   // C type: REGISTERCONTEXT
	Flags                uintptr           // C type: FLAGS
	X87FPURegisters      [8]X87FPUREGISTER // C type: X87FPUREGISTER[8]
	Mmx                  [8]uint64         // C type: unsigned long long[8]
	MxCsrFields          uintptr           // C type: MXCSRFIELDS
	X87StatusWordFields  uintptr           // C type: X87STATUSWORDFIELDS
	X87ControlWordFields uintptr           // C type: X87CONTROLWORDFIELDS
	LastError            LASTERROR         // C type: LASTERROR
	LastStatus           LASTSTATUS        // C type: LASTSTATUS
}

// DISASM_ARG (:796)
type DISASM_ARG struct {
	Type     DISASM_ARGTYPE // C type: DISASM_ARGTYPE
	Segment  SEGMENTREG     // C type: SEGMENTREG
	Mnemonic [64]int8       // C type: char[64]
	Constant uint           // C type: duint
	Value    uint           // C type: duint
	Memvalue uint           // C type: duint
}

// DISASM_INSTR (:806)
type DISASM_INSTR struct {
	Instruction [64]int8         // C type: char[64]
	Type        DISASM_INSTRTYPE // C type: DISASM_INSTRTYPE
	Argcount    int              // C type: int
	Instr_size  int              // C type: int
	Arg         [3]DISASM_ARG    // C type: DISASM_ARG[3]
}

// STACK_COMMENT (:815)
type STACK_COMMENT struct {
	Color   [8]int8   // C type: char[8]
	Comment [512]int8 // C type: char[512]
}

// THREADINFO (:821)
type THREADINFO struct {
	ThreadNumber       int       // C type: int
	Handle             uintptr   // C type: HANDLE
	ThreadId           uint32    // C type: DWORD
	ThreadStartAddress uint      // C type: duint
	ThreadLocalBase    uint      // C type: duint
	ThreadName         [256]int8 // C type: char[256]
}

// THREADALLINFO (:831)
type THREADALLINFO struct {
	BasicInfo    THREADINFO       // C type: THREADINFO
	ThreadCip    uint             // C type: duint
	SuspendCount uint32           // C type: DWORD
	Priority     THREADPRIORITY   // C type: THREADPRIORITY
	WaitReason   THREADWAITREASON // C type: THREADWAITREASON
	LastError    uint32           // C type: DWORD
	UserTime     uint64           // C type: FILETIME
	KernelTime   uint64           // C type: FILETIME
	CreationTime uint64           // C type: FILETIME
	Cycles       uint64           // C type: ULONG64
}

// THREADLIST (:845)
type THREADLIST struct {
	Count         int            // C type: int
	List          *THREADALLINFO // C type: THREADALLINFO *
	CurrentThread int            // C type: int
}

// MEMORY_INFO (:852)
type MEMORY_INFO struct {
	Value    uint        // C type: duint
	Size     MEMORY_SIZE // C type: MEMORY_SIZE
	Mnemonic [64]int8    // C type: char[64]
}

// VALUE_INFO (:859)
type VALUE_INFO struct {
	Value uint    // C type: duint
	Size  uintptr // C type: VALUE_SIZE
}

// SCRIPTBRANCH (:881)
type SCRIPTBRANCH struct {
	Type        SCRIPTBRANCHTYPE // C type: SCRIPTBRANCHTYPE
	Dest        int              // C type: int
	Branchlabel [256]int8        // C type: char[256]
}

// XREF_RECORD (:897)
type XREF_RECORD struct {
	Addr uint     // C type: duint
	Type XREFTYPE // C type: XREFTYPE
}

// XREF_INFO (:903)
type XREF_INFO struct {
	Refcount   uint         // C type: duint
	References *XREF_RECORD // C type: XREF_RECORD *
}

// XREF_EDGE (:909)
type XREF_EDGE struct {
	Address uint // C type: duint
	From    uint // C type: duint
}

// SYMBOLPTR_ (:915)
type SYMBOLPTR_ struct {
	Modbase uint     // C type: duint
	Symbol  *uintptr // C type: const void *
}

// CELLINFO (:1193)
type CELLINFO struct {
	Row int    // C type: int
	Col int    // C type: int
	Str string // C type: const char *
}

// SELECTIONDATA (:1200)
type SELECTIONDATA struct {
	Start uint // C type: duint
	End   uint // C type: duint
}

// ICONDATA (:1206)
type ICONDATA struct {
	Data *uintptr // C type: const void *
	Size uint     // C type: duint
}

// SCRIPTTYPEINFO (:1212)
type SCRIPTTYPEINFO struct {
	Name            [64]int8 // C type: char[64]
	Id              int      // C type: int
	Execute         uintptr  // C type: GUISCRIPTEXECUTE
	CompleteCommand uintptr  // C type: GUISCRIPTCOMPLETER
}

// ACTIVEVIEW (:1220)
type ACTIVEVIEW struct {
	TitleHwnd *uintptr  // C type: void *
	ClassHwnd *uintptr  // C type: void *
	Title     [512]int8 // C type: char[512]
	ClassName [512]int8 // C type: char[512]
}
type bridgemain struct{}

func (b *bridgemain) BridgeInit() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeInit").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeStart() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeStart").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeAlloc(size uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeAlloc").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeFree(ptr *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeFree").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ptr",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", ptr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeSettingGet(section string, key string, value *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "string",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "string",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeSettingGetUint(section string, key string, value *uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingGetUint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "string",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "string",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeSettingSet(section string, key string, value string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingSet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "string",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "string",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "string",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeSettingSetUint(section string, key string, value uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingSetUint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "string",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "string",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeSettingFlush() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingFlush").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeSettingRead(errorLine *int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingRead").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "errorLine",
				Type:  "*int ",
				Value: fmt.Sprintf("%v", errorLine),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeGetDbgVersion() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeGetDbgVersion").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeIsProcessElevated() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeIsProcessElevated").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeGetNtBuildNumber() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeGetNtBuildNumber").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeUserDirectory() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeUserDirectory").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeIsARM64Emulated() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeIsARM64Emulated").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgInit() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgInit").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgExit() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgExit").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgMemRead(va uint, dest *uintptr, size uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemRead").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "va",
				Type:  "uint",
				Value: fmt.Sprintf("%v", va),
			},
			{
				Name:  "dest",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", dest),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgMemWrite(va uint, src *uintptr, size uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemWrite").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "va",
				Type:  "uint",
				Value: fmt.Sprintf("%v", va),
			},
			{
				Name:  "src",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", src),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgMemGetPageSize(base uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemGetPageSize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "base",
				Type:  "uint",
				Value: fmt.Sprintf("%v", base),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgMemFindBaseAddr(addr uint, size *uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemFindBaseAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "size",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgCmdExec(cmd string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgCmdExec").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "string",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgCmdExecDirect(cmd string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgCmdExecDirect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "string",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgMemMap(memmap *MEMMAP) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemMap").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "memmap",
				Type:  "*MEMMAP ",
				Value: fmt.Sprintf("%v", memmap),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgIsValidExpression(expression string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsValidExpression").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "expression",
				Type:  "string",
				Value: fmt.Sprintf("%v", expression),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgIsDebugging() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsDebugging").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgIsJumpGoingToExecute(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsJumpGoingToExecute").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetLabelAt(addr uint, segment SEGMENTREG, text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetLabelAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "segment",
				Type:  "SEGMENTREG",
				Value: fmt.Sprintf("%v", segment),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSetLabelAt(addr uint, text string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetLabelAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "string",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgClearLabelRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgClearLabelRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetCommentAt(addr uint, text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetCommentAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSetCommentAt(addr uint, text string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetCommentAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "string",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgClearCommentRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgClearCommentRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetBookmarkAt(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBookmarkAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgClearBookmarkRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgClearBookmarkRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetModuleAt(addr uint, text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetModuleAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetBpxTypeAt(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBpxTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgValFromString(s string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgValFromString").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "s",
				Type:  "string",
				Value: fmt.Sprintf("%v", s),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetRegDumpEx(regdump *REGDUMP, size uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetRegDumpEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "regdump",
				Type:  "*REGDUMP ",
				Value: fmt.Sprintf("%v", regdump),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgValToString(s string, value uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgValToString").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "s",
				Type:  "string",
				Value: fmt.Sprintf("%v", s),
			},
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgMemIsValidReadPtr(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemIsValidReadPtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetBpList(Type BPXTYPE, list *BPMAP) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBpList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "Type",
				Type:  "BPXTYPE",
				Value: fmt.Sprintf("%v", Type),
			},
			{
				Name:  "list",
				Type:  "*BPMAP ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetFunctionTypeAt(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetFunctionTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetLoopTypeAt(addr uint, depth int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetLoopTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "depth",
				Type:  "int",
				Value: fmt.Sprintf("%v", depth),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetBranchDestination(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBranchDestination").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptLoad(filename string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptLoad").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "string",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptUnload() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptUnload").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptRun(destline int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptRun").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "destline",
				Type:  "int",
				Value: fmt.Sprintf("%v", destline),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptStep() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptStep").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptBpToggle(line int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptBpToggle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptBpGet(line int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptBpGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptCmdExec(command string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptCmdExec").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "command",
				Type:  "string",
				Value: fmt.Sprintf("%v", command),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptAbort() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptAbort").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptGetLineType(line int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptGetLineType").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptSetIp(line int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptSetIp").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgScriptGetBranchInfo(line int, info *SCRIPTBRANCH) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptGetBranchInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
			{
				Name:  "info",
				Type:  "*SCRIPTBRANCH ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSymbolEnum(base uint, cbSymbolEnum uintptr, user *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSymbolEnum").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "base",
				Type:  "uint",
				Value: fmt.Sprintf("%v", base),
			},
			{
				Name:  "cbSymbolEnum",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", cbSymbolEnum),
			},
			{
				Name:  "user",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSymbolEnumFromCache(base uint, cbSymbolEnum uintptr, user *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSymbolEnumFromCache").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "base",
				Type:  "uint",
				Value: fmt.Sprintf("%v", base),
			},
			{
				Name:  "cbSymbolEnum",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", cbSymbolEnum),
			},
			{
				Name:  "user",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSymbolEnumRange(start uint, end uint, symbolMask uint, cbSymbolEnum uintptr, user *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSymbolEnumRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
			{
				Name:  "symbolMask",
				Type:  "uint",
				Value: fmt.Sprintf("%v", symbolMask),
			},
			{
				Name:  "cbSymbolEnum",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", cbSymbolEnum),
			},
			{
				Name:  "user",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgAssembleAt(addr uint, instruction string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgAssembleAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "instruction",
				Type:  "string",
				Value: fmt.Sprintf("%v", instruction),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgModBaseFromName(name string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgModBaseFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgDisasmAt(addr uint, instr *DISASM_INSTR) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgDisasmAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "instr",
				Type:  "*DISASM_INSTR ",
				Value: fmt.Sprintf("%v", instr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgStackCommentGet(addr uint, comment *STACK_COMMENT) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgStackCommentGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "comment",
				Type:  "*STACK_COMMENT ",
				Value: fmt.Sprintf("%v", comment),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetThreadList(list *THREADLIST) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetThreadList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*THREADLIST ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSettingsUpdated() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSettingsUpdated").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgDisasmFastAt(addr uint, basicinfo *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgDisasmFastAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "basicinfo",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", basicinfo),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgMenuEntryClicked(hEntry int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMenuEntryClicked").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntry",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntry),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgFunctionGet(addr uint, start *uint, end *uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgFunctionGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "start",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgFunctionOverlaps(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgFunctionOverlaps").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgFunctionAdd(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgFunctionAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgFunctionDel(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgFunctionDel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgArgumentGet(addr uint, start *uint, end *uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgArgumentGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "start",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgArgumentOverlaps(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgArgumentOverlaps").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgArgumentAdd(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgArgumentAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgArgumentDel(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgArgumentDel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgLoopGet(depth int, addr uint, start *uint, end *uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgLoopGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "depth",
				Type:  "int",
				Value: fmt.Sprintf("%v", depth),
			},
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "start",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgLoopOverlaps(depth int, start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgLoopOverlaps").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "depth",
				Type:  "int",
				Value: fmt.Sprintf("%v", depth),
			},
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgLoopAdd(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgLoopAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgLoopDel(depth int, addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgLoopDel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "depth",
				Type:  "int",
				Value: fmt.Sprintf("%v", depth),
			},
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgXrefAdd(addr uint, from uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgXrefAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "from",
				Type:  "uint",
				Value: fmt.Sprintf("%v", from),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgXrefDelAll(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgXrefDelAll").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgXrefGet(addr uint, info *XREF_INFO) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgXrefGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*XREF_INFO ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetXrefCountAt(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetXrefCountAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetXrefTypeAt(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetXrefTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgIsRunLocked() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsRunLocked").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgIsBpDisabled(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsBpDisabled").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSetAutoCommentAt(addr uint, text string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoCommentAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "string",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgClearAutoCommentRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgClearAutoCommentRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSetAutoLabelAt(addr uint, text string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoLabelAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "string",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgClearAutoLabelRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgClearAutoLabelRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSetAutoBookmarkAt(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoBookmarkAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgClearAutoBookmarkRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgClearAutoBookmarkRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSetAutoFunctionAt(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoFunctionAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgClearAutoFunctionRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgClearAutoFunctionRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetStringAt(addr uint, text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetStringAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgFunctions() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgFunctions").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgWinEvent(message *uintptr, result *int32) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgWinEvent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", message),
			},
			{
				Name:  "result",
				Type:  "*int32 ",
				Value: fmt.Sprintf("%v", result),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgWinEventGlobal(message *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgWinEventGlobal").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgIsRunning() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsRunning").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetTimeWastedCounter() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetTimeWastedCounter").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetArgTypeAt(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetArgTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetEncodeTypeBuffer(addr uint, size *uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetEncodeTypeBuffer").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "size",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgReleaseEncodeTypeBuffer(buffer *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgReleaseEncodeTypeBuffer").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "buffer",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", buffer),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetEncodeTypeAt(addr uint, size uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetEncodeTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetEncodeSizeAt(addr uint, codesize uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetEncodeSizeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "codesize",
				Type:  "uint",
				Value: fmt.Sprintf("%v", codesize),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSetEncodeType(addr uint, size uint, Type ENCODETYPE) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetEncodeType").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "Type",
				Type:  "ENCODETYPE",
				Value: fmt.Sprintf("%v", Type),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgDelEncodeTypeRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgDelEncodeTypeRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgDelEncodeTypeSegment(start uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgDelEncodeTypeSegment").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetWatchList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetWatchList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgSelChanged(hWindow int, VA uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSelChanged").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hWindow",
				Type:  "int",
				Value: fmt.Sprintf("%v", hWindow),
			},
			{
				Name:  "VA",
				Type:  "uint",
				Value: fmt.Sprintf("%v", VA),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetProcessHandle() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetProcessHandle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetThreadHandle() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetThreadHandle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetProcessId() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetProcessId").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetThreadId() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetThreadId").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetPebAddress(ProcessId uint32) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetPebAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ProcessId",
				Type:  "uint32",
				Value: fmt.Sprintf("%v", ProcessId),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetTebAddress(ThreadId uint32) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetTebAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ThreadId",
				Type:  "uint32",
				Value: fmt.Sprintf("%v", ThreadId),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgAnalyzeFunction(entry uint, graph *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgAnalyzeFunction").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "entry",
				Type:  "uint",
				Value: fmt.Sprintf("%v", entry),
			},
			{
				Name:  "graph",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", graph),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetSymbolInfo(symbolptr *uintptr, info *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetSymbolInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "symbolptr",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", symbolptr),
			},
			{
				Name:  "info",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetDebugEngine() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetDebugEngine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgGetSymbolInfoAt(addr uint, info *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetSymbolInfoAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgXrefAddMulti(edges *XREF_EDGE, count uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgXrefAddMulti").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "edges",
				Type:  "*XREF_EDGE ",
				Value: fmt.Sprintf("%v", edges),
			},
			{
				Name:  "count",
				Type:  "uint",
				Value: fmt.Sprintf("%v", count),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) DbgMenuPrepare(hMenu GUIMENUTYPE) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMenuPrepare").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "GUIMENUTYPE",
				Value: fmt.Sprintf("%v", hMenu),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiTranslateText(Source string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiTranslateText").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "Source",
				Type:  "string",
				Value: fmt.Sprintf("%v", Source),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiDisasmAt(addr uint, cip uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiDisasmAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "cip",
				Type:  "uint",
				Value: fmt.Sprintf("%v", cip),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSetDebugState(state DBGSTATE) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetDebugState").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "state",
				Type:  "DBGSTATE",
				Value: fmt.Sprintf("%v", state),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSetDebugStateFast(state DBGSTATE) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetDebugStateFast").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "state",
				Type:  "DBGSTATE",
				Value: fmt.Sprintf("%v", state),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAddLogMessage(msg string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddLogMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "string",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAddLogMessageHtml(msg string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddLogMessageHtml").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "string",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiLogClear() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiLogSave(filename string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogSave").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "string",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiLogRedirect(filename string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogRedirect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "string",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiLogRedirectStop() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogRedirectStop").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateAllViews() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateAllViews").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateRegisterView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateRegisterView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateDisassemblyView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateDisassemblyView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateBreakpointsView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateBreakpointsView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateWindowTitle(filename string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateWindowTitle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "string",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGetWindowHandle() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetWindowHandle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiDumpAt(va uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiDumpAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "va",
				Type:  "uint",
				Value: fmt.Sprintf("%v", va),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiScriptAdd(count int, lines *string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "count",
				Type:  "int",
				Value: fmt.Sprintf("%v", count),
			},
			{
				Name:  "lines",
				Type:  "*string",
				Value: fmt.Sprintf("%v", lines),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiScriptClear() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiScriptSetIp(line int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptSetIp").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiScriptError(line int, message string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptError").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
			{
				Name:  "message",
				Type:  "string",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiScriptSetTitle(title string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptSetTitle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiScriptSetInfoLine(line int, info string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptSetInfoLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
			{
				Name:  "info",
				Type:  "string",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiScriptMessage(message string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "string",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiScriptMsgyn(message string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptMsgyn").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "string",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSymbolLogAdd(message string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolLogAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "string",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSymbolLogClear() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolLogClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSymbolSetProgress(percent int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolSetProgress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "percent",
				Type:  "int",
				Value: fmt.Sprintf("%v", percent),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSymbolUpdateModuleList(count int, modules *SYMBOLMODULEINFO) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolUpdateModuleList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "count",
				Type:  "int",
				Value: fmt.Sprintf("%v", count),
			},
			{
				Name:  "modules",
				Type:  "*SYMBOLMODULEINFO ",
				Value: fmt.Sprintf("%v", modules),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSymbolRefreshCurrent() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolRefreshCurrent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceAddColumn(width int, title string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceAddColumn").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "width",
				Type:  "int",
				Value: fmt.Sprintf("%v", width),
			},
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceSetRowCount(count int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSetRowCount").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "count",
				Type:  "int",
				Value: fmt.Sprintf("%v", count),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceGetRowCount() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceGetRowCount").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceSearchGetRowCount() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSearchGetRowCount").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceDeleteAllColumns() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceDeleteAllColumns").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceInitialize(name string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceInitialize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceSetCellContent(row int, col int, str string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSetCellContent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "row",
				Type:  "int",
				Value: fmt.Sprintf("%v", row),
			},
			{
				Name:  "col",
				Type:  "int",
				Value: fmt.Sprintf("%v", col),
			},
			{
				Name:  "str",
				Type:  "string",
				Value: fmt.Sprintf("%v", str),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceGetCellContent(row int, col int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceGetCellContent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "row",
				Type:  "int",
				Value: fmt.Sprintf("%v", row),
			},
			{
				Name:  "col",
				Type:  "int",
				Value: fmt.Sprintf("%v", col),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceSearchGetCellContent(row int, col int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSearchGetCellContent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "row",
				Type:  "int",
				Value: fmt.Sprintf("%v", row),
			},
			{
				Name:  "col",
				Type:  "int",
				Value: fmt.Sprintf("%v", col),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceReloadData() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceReloadData").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceSetProgress(progress int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSetProgress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "progress",
				Type:  "int",
				Value: fmt.Sprintf("%v", progress),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceSetCurrentTaskProgress(progress int, taskTitle string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSetCurrentTaskProgress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "progress",
				Type:  "int",
				Value: fmt.Sprintf("%v", progress),
			},
			{
				Name:  "taskTitle",
				Type:  "string",
				Value: fmt.Sprintf("%v", taskTitle),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceSetSearchStartCol(col int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSetSearchStartCol").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "col",
				Type:  "int",
				Value: fmt.Sprintf("%v", col),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiStackDumpAt(addr uint, csp uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiStackDumpAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "csp",
				Type:  "uint",
				Value: fmt.Sprintf("%v", csp),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateDumpView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateDumpView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateWatchView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateWatchView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateThreadView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateThreadView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateMemoryView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateMemoryView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAddRecentFile(file string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddRecentFile").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "file",
				Type:  "string",
				Value: fmt.Sprintf("%v", file),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSetLastException(exception uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetLastException").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "exception",
				Type:  "uint",
				Value: fmt.Sprintf("%v", exception),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGetDisassembly(addr uint, text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetDisassembly").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuAdd(hMenu int, title string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuAddEntry(hMenu int, title string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuAddEntry").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuAddSeparator(hMenu int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuAddSeparator").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuClear(hMenu int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuRemove(hEntryMenu int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuRemove").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntryMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntryMenu),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSelectionGet(hWindow GUISELECTIONTYPE, selection *SELECTIONDATA) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSelectionGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hWindow",
				Type:  "GUISELECTIONTYPE",
				Value: fmt.Sprintf("%v", hWindow),
			},
			{
				Name:  "selection",
				Type:  "*SELECTIONDATA ",
				Value: fmt.Sprintf("%v", selection),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSelectionSet(hWindow GUISELECTIONTYPE, selection *SELECTIONDATA) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSelectionSet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hWindow",
				Type:  "GUISELECTIONTYPE",
				Value: fmt.Sprintf("%v", hWindow),
			},
			{
				Name:  "selection",
				Type:  "*SELECTIONDATA ",
				Value: fmt.Sprintf("%v", selection),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGetLineWindow(title string, text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetLineWindow").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAutoCompleteAddCmd(cmd string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAutoCompleteAddCmd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "string",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAutoCompleteDelCmd(cmd string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAutoCompleteDelCmd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "string",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAutoCompleteClearAll() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAutoCompleteClearAll").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAddStatusBarMessage(msg string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddStatusBarMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "string",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateSideBar() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateSideBar").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiRepaintTableView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiRepaintTableView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdatePatches() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdatePatches").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateCallStack() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateCallStack").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateSEHChain() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateSEHChain").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiLoadSourceFileEx(path string, addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLoadSourceFileEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "path",
				Type:  "string",
				Value: fmt.Sprintf("%v", path),
			},
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuSetIcon(hMenu int, icon *ICONDATA) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetIcon").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "icon",
				Type:  "*ICONDATA ",
				Value: fmt.Sprintf("%v", icon),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuSetEntryIcon(hEntry int, icon *ICONDATA) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetEntryIcon").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntry",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntry),
			},
			{
				Name:  "icon",
				Type:  "*ICONDATA ",
				Value: fmt.Sprintf("%v", icon),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuSetName(hMenu int, name string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuSetEntryName(hEntry int, name string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetEntryName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntry",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntry),
			},
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiMenuSetEntryHotkey(hEntry int, hack string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetEntryHotkey").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntry",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntry),
			},
			{
				Name:  "hack",
				Type:  "string",
				Value: fmt.Sprintf("%v", hack),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiShowCpu() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowCpu").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiShowThreads() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowThreads").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAddQWidgetTab(qWidget *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiShowQWidgetTab(qWidget *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiCloseQWidgetTab(qWidget *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiCloseQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiExecuteOnGuiThread(cbGuiThread uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiExecuteOnGuiThread").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cbGuiThread",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", cbGuiThread),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateTimeWastedCounter() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateTimeWastedCounter").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSetGlobalNotes(text string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetGlobalNotes").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "text",
				Type:  "string",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGetGlobalNotes(text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetGlobalNotes").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSetDebuggeeNotes(text string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetDebuggeeNotes").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "text",
				Type:  "string",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGetDebuggeeNotes(text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetDebuggeeNotes").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiDumpAtN(va uint, index int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiDumpAtN").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "va",
				Type:  "uint",
				Value: fmt.Sprintf("%v", va),
			},
			{
				Name:  "index",
				Type:  "int",
				Value: fmt.Sprintf("%v", index),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiDisplayWarning(title string, text string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiDisplayWarning").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "text",
				Type:  "string",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiRegisterScriptLanguage(info *SCRIPTTYPEINFO) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiRegisterScriptLanguage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*SCRIPTTYPEINFO ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUnregisterScriptLanguage(id int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUnregisterScriptLanguage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "id",
				Type:  "int",
				Value: fmt.Sprintf("%v", id),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateArgumentWidget() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateArgumentWidget").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiFocusView(hWindow int) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiFocusView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hWindow",
				Type:  "int",
				Value: fmt.Sprintf("%v", hWindow),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiIsUpdateDisabled() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiIsUpdateDisabled").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateDisable() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateDisable").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiLoadGraph(graph *uintptr, addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLoadGraph").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graph",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", graph),
			},
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGraphAt(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGraphAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateGraphView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateGraphView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiDisableLog() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiDisableLog").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiEnableLog() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiEnableLog").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiIsLogEnabled() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiIsLogEnabled").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAddFavouriteTool(name string, description string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddFavouriteTool").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "description",
				Type:  "string",
				Value: fmt.Sprintf("%v", description),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAddFavouriteCommand(name string, shortcut string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddFavouriteCommand").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "shortcut",
				Type:  "string",
				Value: fmt.Sprintf("%v", shortcut),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSetFavouriteToolShortcut(name string, shortcut string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetFavouriteToolShortcut").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "shortcut",
				Type:  "string",
				Value: fmt.Sprintf("%v", shortcut),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiFoldDisassembly(startAddress uint, length uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiFoldDisassembly").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "startAddress",
				Type:  "uint",
				Value: fmt.Sprintf("%v", startAddress),
			},
			{
				Name:  "length",
				Type:  "uint",
				Value: fmt.Sprintf("%v", length),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSelectInMemoryMap(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSelectInMemoryMap").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGetActiveView(activeView *ACTIVEVIEW) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetActiveView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "activeView",
				Type:  "*ACTIVEVIEW ",
				Value: fmt.Sprintf("%v", activeView),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiAddInfoLine(infoLine string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddInfoLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "infoLine",
				Type:  "string",
				Value: fmt.Sprintf("%v", infoLine),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiProcessEvents() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiProcessEvents").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiTypeAddNode(parent *uintptr, Type *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiTypeAddNode").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "parent",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", parent),
			},
			{
				Name:  "Type",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", Type),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiTypeClear() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiTypeClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateTypeWidget() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateTypeWidget").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiCloseApplication() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiCloseApplication").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiFlushLog() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiFlushLog").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiReferenceAddCommand(title string, command string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceAddCommand").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "command",
				Type:  "string",
				Value: fmt.Sprintf("%v", command),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiUpdateTraceBrowser() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateTraceBrowser").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiOpenTraceFile(fileName string) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiOpenTraceFile").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "fileName",
				Type:  "string",
				Value: fmt.Sprintf("%v", fileName),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiInvalidateSymbolSource(base uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiInvalidateSymbolSource").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "base",
				Type:  "uint",
				Value: fmt.Sprintf("%v", base),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiExecuteOnGuiThreadEx(cbGuiThread uintptr, userdata *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiExecuteOnGuiThreadEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cbGuiThread",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", cbGuiThread),
			},
			{
				Name:  "userdata",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", userdata),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGetCurrentGraph(graphList *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetCurrentGraph").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graphList",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", graphList),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiShowReferences() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowReferences").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiSelectInSymbolsTab(addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSelectInSymbolsTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGotoTrace(index uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGotoTrace").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "index",
				Type:  "uint",
				Value: fmt.Sprintf("%v", index),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiShowTrace() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowTrace").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) GuiGetMainThreadId() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetMainThreadId").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
