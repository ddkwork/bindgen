package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/std/mylog"
)

type DBGSTATE int //D:\workspace\workspace\bindgen\project\x64dbg\include\bridgemain.h:113
const (
	initialized DBGSTATE = iota
	paused               = 1
	running              = 2
	stopped              = 3
)

type SEGMENTREG int //:121
const (
	SEG_DEFAULT SEGMENTREG = iota
	SEG_ES                 = 1
	SEG_DS                 = 2
	SEG_FS                 = 3
	SEG_GS                 = 4
	SEG_CS                 = 5
	SEG_SS                 = 6
)

type ADDRINFOFLAGS int //:132
const (
	flagmodule       ADDRINFOFLAGS = iota + 1
	flaglabel                      = 2
	flagcomment                    = 4
	flagbookmark                   = 8
	flagfunction                   = 16
	flagloop                       = 32
	flagargs                       = 64
	flagNoFuncOffset               = 128
)

type BPXTYPE int //:144
const (
	bp_none      BPXTYPE = iota + 0
	bp_normal            = 1
	bp_hardware          = 2
	bp_memory            = 4
	bp_dll               = 8
	bp_exception         = 16
)

type FUNCTYPE int //:154
const (
	FUNC_NONE   FUNCTYPE = iota
	FUNC_BEGIN           = 1
	FUNC_MIDDLE          = 2
	FUNC_END             = 3
	FUNC_SINGLE          = 4
)

type LOOPTYPE int //:163
const (
	LOOP_NONE   LOOPTYPE = iota
	LOOP_BEGIN           = 1
	LOOP_MIDDLE          = 2
	LOOP_ENTRY           = 3
	LOOP_END             = 4
	LOOP_SINGLE          = 5
)

type XREFTYPE int //:173
const (
	XREF_NONE XREFTYPE = iota
	XREF_DATA          = 1
	XREF_JMP           = 2
	XREF_CALL          = 3
)

type ARGTYPE int //:181
const (
	ARG_NONE   ARGTYPE = iota
	ARG_BEGIN          = 1
	ARG_MIDDLE         = 2
	ARG_END            = 3
	ARG_SINGLE         = 4
)

type DBGMSG int //:190
const (
	DBG_SCRIPT_LOAD                DBGMSG = iota
	DBG_SCRIPT_UNLOAD                     = 1
	DBG_SCRIPT_RUN                        = 2
	DBG_SCRIPT_STEP                       = 3
	DBG_SCRIPT_BPTOGGLE                   = 4
	DBG_SCRIPT_BPGET                      = 5
	DBG_SCRIPT_CMDEXEC                    = 6
	DBG_SCRIPT_ABORT                      = 7
	DBG_SCRIPT_GETLINETYPE                = 8
	DBG_SCRIPT_SETIP                      = 9
	DBG_SCRIPT_GETBRANCHINFO              = 10
	DBG_SYMBOL_ENUM                       = 11
	DBG_ASSEMBLE_AT                       = 12
	DBG_MODBASE_FROM_NAME                 = 13
	DBG_DISASM_AT                         = 14
	DBG_STACK_COMMENT_GET                 = 15
	DBG_GET_THREAD_LIST                   = 16
	DBG_SETTINGS_UPDATED                  = 17
	DBG_DISASM_FAST_AT                    = 18
	DBG_MENU_ENTRY_CLICKED                = 19
	DBG_FUNCTION_GET                      = 20
	DBG_FUNCTION_OVERLAPS                 = 21
	DBG_FUNCTION_ADD                      = 22
	DBG_FUNCTION_DEL                      = 23
	DBG_LOOP_GET                          = 24
	DBG_LOOP_OVERLAPS                     = 25
	DBG_LOOP_ADD                          = 26
	DBG_LOOP_DEL                          = 27
	DBG_IS_RUN_LOCKED                     = 28
	DBG_IS_BP_DISABLED                    = 29
	DBG_SET_AUTO_COMMENT_AT               = 30
	DBG_DELETE_AUTO_COMMENT_RANGE         = 31
	DBG_SET_AUTO_LABEL_AT                 = 32
	DBG_DELETE_AUTO_LABEL_RANGE           = 33
	DBG_SET_AUTO_BOOKMARK_AT              = 34
	DBG_DELETE_AUTO_BOOKMARK_RANGE        = 35
	DBG_SET_AUTO_FUNCTION_AT              = 36
	DBG_DELETE_AUTO_FUNCTION_RANGE        = 37
	DBG_GET_STRING_AT                     = 38
	DBG_GET_FUNCTIONS                     = 39
	DBG_WIN_EVENT                         = 40
	DBG_WIN_EVENT_GLOBAL                  = 41
	DBG_INITIALIZE_LOCKS                  = 42
	DBG_DEINITIALIZE_LOCKS                = 43
	DBG_GET_TIME_WASTED_COUNTER           = 44
	DBG_SYMBOL_ENUM_FROMCACHE             = 45
	DBG_DELETE_COMMENT_RANGE              = 46
	DBG_DELETE_LABEL_RANGE                = 47
	DBG_DELETE_BOOKMARK_RANGE             = 48
	DBG_GET_XREF_COUNT_AT                 = 49
	DBG_GET_XREF_TYPE_AT                  = 50
	DBG_XREF_ADD                          = 51
	DBG_XREF_DEL_ALL                      = 52
	DBG_XREF_GET                          = 53
	DBG_GET_ENCODE_TYPE_BUFFER            = 54
	DBG_ENCODE_TYPE_GET                   = 55
	DBG_DELETE_ENCODE_TYPE_RANGE          = 56
	DBG_ENCODE_SIZE_GET                   = 57
	DBG_DELETE_ENCODE_TYPE_SEG            = 58
	DBG_RELEASE_ENCODE_TYPE_BUFFER        = 59
	DBG_ARGUMENT_GET                      = 60
	DBG_ARGUMENT_OVERLAPS                 = 61
	DBG_ARGUMENT_ADD                      = 62
	DBG_ARGUMENT_DEL                      = 63
	DBG_GET_WATCH_LIST                    = 64
	DBG_SELCHANGED                        = 65
	DBG_GET_PROCESS_HANDLE                = 66
	DBG_GET_THREAD_HANDLE                 = 67
	DBG_GET_PROCESS_ID                    = 68
	DBG_GET_THREAD_ID                     = 69
	DBG_GET_PEB_ADDRESS                   = 70
	DBG_GET_TEB_ADDRESS                   = 71
	DBG_ANALYZE_FUNCTION                  = 72
	DBG_MENU_PREPARE                      = 73
	DBG_GET_SYMBOL_INFO                   = 74
	DBG_GET_DEBUG_ENGINE                  = 75
	DBG_GET_SYMBOL_INFO_AT                = 76
	DBG_XREF_ADD_MULTI                    = 77
)

type SCRIPTLINETYPE int //:272
const (
	linecommand SCRIPTLINETYPE = iota
	linebranch                 = 1
	linelabel                  = 2
	linecomment                = 3
	lineempty                  = 4
)

type SCRIPTBRANCHTYPE int //:281
const (
	scriptnobranch SCRIPTBRANCHTYPE = iota
	scriptjmp                       = 1
	scriptjnejnz                    = 2
	scriptjejz                      = 3
	scriptjbjl                      = 4
	scriptjajg                      = 5
	scriptjbejle                    = 6
	scriptjaejge                    = 7
	scriptcall                      = 8
)

type DISASM_INSTRTYPE int //:294
const (
	instr_normal DISASM_INSTRTYPE = iota
	instr_branch                  = 1
	instr_stack                   = 2
)

type DISASM_ARGTYPE int //:301
const (
	arg_normal DISASM_ARGTYPE = iota
	arg_memory                = 1
)

type STRING_TYPE int //:307
const (
	str_none    STRING_TYPE = iota
	str_ascii               = 1
	str_unicode             = 2
)

type THREADPRIORITY int //:314
const (
	_PriorityIdle        THREADPRIORITY = iota + -15
	_PriorityAboveNormal                = 1
	_PriorityBelowNormal
	_PriorityHighest = 2
	_PriorityLowest
	_PriorityNormal
	_PriorityTimeCritical = 15
	_PriorityUnknown      = 2147483647
)

type THREADWAITREASON int //:326
const (
	_Executive        THREADWAITREASON = iota + 0
	_FreePage                          = 1
	_PageIn                            = 2
	_PoolAllocation                    = 3
	_DelayExecution                    = 4
	_Suspended                         = 5
	_UserRequest                       = 6
	_WrExecutive                       = 7
	_WrFreePage                        = 8
	_WrPageIn                          = 9
	_WrPoolAllocation                  = 10
	_WrDelayExecution                  = 11
	_WrSuspended                       = 12
	_WrUserRequest                     = 13
	_WrEventPair                       = 14
	_WrQueue                           = 15
	_WrLpcReceive                      = 16
	_WrLpcReply                        = 17
	_WrVirtualMemory                   = 18
	_WrPageOut                         = 19
	_WrRendezvous                      = 20
	_Spare2                            = 21
	_Spare3                            = 22
	_Spare4                            = 23
	_Spare5                            = 24
	_WrCalloutStack                    = 25
	_WrKernel                          = 26
	_WrResource                        = 27
	_WrPushLock                        = 28
	_WrMutex                           = 29
	_WrQuantumEnd                      = 30
	_WrDispatchInt                     = 31
	_WrPreempted                       = 32
	_WrYieldExecution                  = 33
	_WrFastMutex                       = 34
	_WrGuardedMutex                    = 35
	_WrRundown                         = 36
)

type MEMORY_SIZE int //:367
const (
	size_byte    MEMORY_SIZE = iota + 1
	size_word                = 2
	size_dword               = 4
	size_qword               = 8
	size_xmmword             = 16
	size_ymmword             = 32
)

type ENCODETYPE int //:377
const (
	enc_unknown ENCODETYPE = iota
	enc_byte               = 1
	enc_word               = 2
	enc_dword              = 3
	enc_fword              = 4
	enc_qword              = 5
	enc_tbyte              = 6
	enc_oword              = 7
	enc_mmword             = 8
	enc_xmmword            = 9
	enc_ymmword            = 10
	enc_zmmword            = 11
	enc_real4              = 12
	enc_real8              = 13
	enc_real10             = 14
	enc_ascii              = 15
	enc_unicode            = 16
	enc_code               = 17
	enc_junk               = 18
	enc_middle             = 19
)

type WATCHVARTYPE int //:401
const (
	TYPE_UINT    WATCHVARTYPE = iota
	TYPE_INT                  = 1
	TYPE_FLOAT                = 2
	TYPE_ASCII                = 3
	TYPE_UNICODE              = 4
	TYPE_INVALID              = 5
)

type WATCHDOGMODE int //:411
const (
	MODE_DISABLED  WATCHDOGMODE = iota
	MODE_ISTRUE                 = 1
	MODE_ISFALSE                = 2
	MODE_CHANGED                = 3
	MODE_UNCHANGED              = 4
)

type BPHWTYPE int //:420
const (
	hw_access  BPHWTYPE = iota
	hw_write            = 1
	hw_execute          = 2
)

type BPMEMTYPE int //:427
const (
	mem_access  BPMEMTYPE = iota
	mem_read              = 1
	mem_write             = 2
	mem_execute           = 3
)

type BPDLLTYPE int //:435
const (
	dll_load   BPDLLTYPE = iota + 1
	dll_unload           = 2
	dll_all              = 3
)

type BPEXTYPE int //:442
const (
	ex_firstchance  BPEXTYPE = iota + 1
	ex_secondchance          = 2
	ex_all                   = 3
)

type BPHWSIZE int //:449
const (
	hw_byte  BPHWSIZE = iota
	hw_word           = 1
	hw_dword          = 2
	hw_qword          = 3
)

type SYMBOLTYPE int //:457
const (
	sym_import SYMBOLTYPE = iota
	sym_export            = 1
	sym_symbol            = 2
)

type MODULEPARTY int //:469
const (
	mod_user   MODULEPARTY = iota
	mod_system             = 1
)

type DEBUG_ENGINE int //:475
const (
	DebugEngineTitanEngine  DEBUG_ENGINE = iota
	DebugEngineGleeBug                   = 1
	DebugEngineStaticEngine              = 2
)

type GUIMENUTYPE int //:1032
const (
	GUI_PLUGIN_MENU GUIMENUTYPE = iota
	GUI_DISASM_MENU             = 1
	GUI_DUMP_MENU               = 2
	GUI_STACK_MENU              = 3
	GUI_GRAPH_MENU              = 4
	GUI_MEMMAP_MENU             = 5
	GUI_SYMMOD_MENU             = 6
)

type GUISELECTIONTYPE int //:1045
const (
	GUI_DISASSEMBLY GUISELECTIONTYPE = iota
	GUI_DUMP                         = 1
	GUI_STACK                        = 2
	GUI_GRAPH                        = 3
	GUI_MEMMAP                       = 4
	GUI_SYMMOD                       = 5
	GUI_THREADS                      = 6
)

type GUIMSG int //:1059
const (
	GUI_DISASSEMBLE_AT              GUIMSG = iota
	GUI_SET_DEBUG_STATE                    = 1
	GUI_ADD_MSG_TO_LOG                     = 2
	GUI_CLEAR_LOG                          = 3
	GUI_UPDATE_REGISTER_VIEW               = 4
	GUI_UPDATE_DISASSEMBLY_VIEW            = 5
	GUI_UPDATE_BREAKPOINTS_VIEW            = 6
	GUI_UPDATE_WINDOW_TITLE                = 7
	GUI_GET_WINDOW_HANDLE                  = 8
	GUI_DUMP_AT                            = 9
	GUI_SCRIPT_ADD                         = 10
	GUI_SCRIPT_CLEAR                       = 11
	GUI_SCRIPT_SETIP                       = 12
	GUI_SCRIPT_ERROR                       = 13
	GUI_SCRIPT_SETTITLE                    = 14
	GUI_SCRIPT_SETINFOLINE                 = 15
	GUI_SCRIPT_MESSAGE                     = 16
	GUI_SCRIPT_MSGYN                       = 17
	GUI_SYMBOL_LOG_ADD                     = 18
	GUI_SYMBOL_LOG_CLEAR                   = 19
	GUI_SYMBOL_SET_PROGRESS                = 20
	GUI_SYMBOL_UPDATE_MODULE_LIST          = 21
	GUI_REF_ADDCOLUMN                      = 22
	GUI_REF_SETROWCOUNT                    = 23
	GUI_REF_GETROWCOUNT                    = 24
	GUI_REF_DELETEALLCOLUMNS               = 25
	GUI_REF_SETCELLCONTENT                 = 26
	GUI_REF_GETCELLCONTENT                 = 27
	GUI_REF_RELOADDATA                     = 28
	GUI_REF_SETSINGLESELECTION             = 29
	GUI_REF_SETPROGRESS                    = 30
	GUI_REF_SETCURRENTTASKPROGRESS         = 31
	GUI_REF_SETSEARCHSTARTCOL              = 32
	GUI_STACK_DUMP_AT                      = 33
	GUI_UPDATE_DUMP_VIEW                   = 34
	GUI_UPDATE_THREAD_VIEW                 = 35
	GUI_ADD_RECENT_FILE                    = 36
	GUI_SET_LAST_EXCEPTION                 = 37
	GUI_GET_DISASSEMBLY                    = 38
	GUI_MENU_ADD                           = 39
	GUI_MENU_ADD_ENTRY                     = 40
	GUI_MENU_ADD_SEPARATOR                 = 41
	GUI_MENU_CLEAR                         = 42
	GUI_SELECTION_GET                      = 43
	GUI_SELECTION_SET                      = 44
	GUI_GETLINE_WINDOW                     = 45
	GUI_AUTOCOMPLETE_ADDCMD                = 46
	GUI_AUTOCOMPLETE_DELCMD                = 47
	GUI_AUTOCOMPLETE_CLEARALL              = 48
	GUI_SCRIPT_ENABLEHIGHLIGHTING          = 49
	GUI_ADD_MSG_TO_STATUSBAR               = 50
	GUI_UPDATE_SIDEBAR                     = 51
	GUI_REPAINT_TABLE_VIEW                 = 52
	GUI_UPDATE_PATCHES                     = 53
	GUI_UPDATE_CALLSTACK                   = 54
	GUI_UPDATE_SEHCHAIN                    = 55
	GUI_SYMBOL_REFRESH_CURRENT             = 56
	GUI_UPDATE_MEMORY_VIEW                 = 57
	GUI_REF_INITIALIZE                     = 58
	GUI_LOAD_SOURCE_FILE                   = 59
	GUI_MENU_SET_ICON                      = 60
	GUI_MENU_SET_ENTRY_ICON                = 61
	GUI_SHOW_CPU                           = 62
	GUI_ADD_QWIDGET_TAB                    = 63
	GUI_SHOW_QWIDGET_TAB                   = 64
	GUI_CLOSE_QWIDGET_TAB                  = 65
	GUI_EXECUTE_ON_GUI_THREAD              = 66
	GUI_UPDATE_TIME_WASTED_COUNTER         = 67
	GUI_SET_GLOBAL_NOTES                   = 68
	GUI_GET_GLOBAL_NOTES                   = 69
	GUI_SET_DEBUGGEE_NOTES                 = 70
	GUI_GET_DEBUGGEE_NOTES                 = 71
	GUI_DUMP_AT_N                          = 72
	GUI_DISPLAY_WARNING                    = 73
	GUI_REGISTER_SCRIPT_LANG               = 74
	GUI_UNREGISTER_SCRIPT_LANG             = 75
	GUI_UPDATE_ARGUMENT_VIEW               = 76
	GUI_FOCUS_VIEW                         = 77
	GUI_UPDATE_WATCH_VIEW                  = 78
	GUI_LOAD_GRAPH                         = 79
	GUI_GRAPH_AT                           = 80
	GUI_UPDATE_GRAPH_VIEW                  = 81
	GUI_SET_LOG_ENABLED                    = 82
	GUI_ADD_FAVOURITE_TOOL                 = 83
	GUI_ADD_FAVOURITE_COMMAND              = 84
	GUI_SET_FAVOURITE_TOOL_SHORTCUT        = 85
	GUI_FOLD_DISASSEMBLY                   = 86
	GUI_SELECT_IN_MEMORY_MAP               = 87
	GUI_GET_ACTIVE_VIEW                    = 88
	GUI_MENU_SET_ENTRY_CHECKED             = 89
	GUI_ADD_INFO_LINE                      = 90
	GUI_PROCESS_EVENTS                     = 91
	GUI_TYPE_ADDNODE                       = 92
	GUI_TYPE_CLEAR                         = 93
	GUI_UPDATE_TYPE_WIDGET                 = 94
	GUI_CLOSE_APPLICATION                  = 95
	GUI_MENU_SET_VISIBLE                   = 96
	GUI_MENU_SET_ENTRY_VISIBLE             = 97
	GUI_MENU_SET_NAME                      = 98
	GUI_MENU_SET_ENTRY_NAME                = 99
	GUI_FLUSH_LOG                          = 100
	GUI_MENU_SET_ENTRY_HOTKEY              = 101
	GUI_REF_SEARCH_GETROWCOUNT             = 102
	GUI_REF_SEARCH_GETCELLCONTENT          = 103
	GUI_MENU_REMOVE                        = 104
	GUI_REF_ADDCOMMAND                     = 105
	GUI_OPEN_TRACE_FILE                    = 106
	GUI_UPDATE_TRACE_BROWSER               = 107
	GUI_INVALIDATE_SYMBOL_SOURCE           = 108
	GUI_GET_CURRENT_GRAPH                  = 109
	GUI_SHOW_REF                           = 110
	GUI_SELECT_IN_SYMBOLS_TAB              = 111
	GUI_GOTO_TRACE                         = 112
	GUI_SHOW_TRACE                         = 113
	GUI_GET_MAIN_THREAD_ID                 = 114
	GUI_ADD_MSG_TO_LOG_HTML                = 115
	GUI_IS_LOG_ENABLED                     = 116
	GUI_IS_DEBUGGER_FOCUSED_UNUSED         = 117
	GUI_SAVE_LOG                           = 118
	GUI_REDIRECT_LOG                       = 119
	GUI_STOP_REDIRECT_LOG                  = 120
	GUI_SHOW_THREADS                       = 121
)

// ListInfo (D:\workspace\workspace\bindgen\project\x64dbg\include\bridgelist.h:4 )
type ListInfo struct {
	count int     // C type: int
	size  uint    // C type: size_t
	data  uintptr // C type: void *
}

// BridgeCFInstruction (D:\workspace\workspace\bindgen\project\x64dbg\include\bridgegraph.h:4 )
type BridgeCFInstruction struct {
	addr uint     // C type: duint
	data [15]byte // C type: unsigned char[15]
}

// BridgeCFGraphList (:26 )
type BridgeCFGraphList struct {
	entryPoint uint     // C type: duint
	userdata   uintptr  // C type: void *
	nodes      ListInfo // C type: ListInfo
}

// MEMPAGE (:488 )
type MEMPAGE struct {
	mbi  uintptr   // C type: MEMORY_BASIC_INFORMATION
	info [256]int8 // C type: char[256]
}

// MEMMAP (:494 )
type MEMMAP struct {
	count int      // C type: int
	page  *MEMPAGE // C type: MEMPAGE *
}

// BPMAP (:523 )
type BPMAP struct {
	count int      // C type: int
	bp    *uintptr // C type: BRIDGEBP *
}

// FUNCTION (:541 )
type FUNCTION struct {
	start      uint // C type: duint
	end        uint // C type: duint
	instrcount uint // C type: duint
}

// LOOP (:548 )
type LOOP struct {
	depth      int  // C type: int
	start      uint // C type: duint
	end        uint // C type: duint
	instrcount uint // C type: duint
}

// SYMBOLMODULEINFO (:609 )
type SYMBOLMODULEINFO struct {
	base uint      // C type: duint
	name [256]int8 // C type: char[256]
}

// SYMBOLCBINFO (:615 )
type SYMBOLCBINFO struct {
	base         uint    // C type: duint
	cbSymbolEnum uintptr // C type: CBSYMBOLENUM
	user         uintptr // C type: void *
	start        uint    // C type: duint
	end          uint    // C type: duint
	symbolMask   uint    // C type: unsigned int
}

// YMMREGISTER (:700 )
type YMMREGISTER struct {
	Low  uintptr // C type: XMMREGISTER
	High uintptr // C type: XMMREGISTER
}

// X87FPUREGISTER (:706 )
type X87FPUREGISTER struct {
	data     [10]byte // C type: BYTE[10]
	st_value int      // C type: int
	tag      int      // C type: int
}

// X87FPU (:713 )
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

// REGISTERCONTEXT (:725 )
type REGISTERCONTEXT struct {
	cax          uintptr         // C type: ULONG_PTR
	ccx          uintptr         // C type: ULONG_PTR
	cdx          uintptr         // C type: ULONG_PTR
	cbx          uintptr         // C type: ULONG_PTR
	csp          uintptr         // C type: ULONG_PTR
	cbp          uintptr         // C type: ULONG_PTR
	csi          uintptr         // C type: ULONG_PTR
	cdi          uintptr         // C type: ULONG_PTR
	r8           uintptr         // C type: ULONG_PTR
	r9           uintptr         // C type: ULONG_PTR
	r10          uintptr         // C type: ULONG_PTR
	r11          uintptr         // C type: ULONG_PTR
	r12          uintptr         // C type: ULONG_PTR
	r13          uintptr         // C type: ULONG_PTR
	r14          uintptr         // C type: ULONG_PTR
	r15          uintptr         // C type: ULONG_PTR
	cip          uintptr         // C type: ULONG_PTR
	eflags       uintptr         // C type: ULONG_PTR
	gs           uint16          // C type: unsigned short
	fs           uint16          // C type: unsigned short
	es           uint16          // C type: unsigned short
	ds           uint16          // C type: unsigned short
	cs           uint16          // C type: unsigned short
	ss           uint16          // C type: unsigned short
	dr0          uintptr         // C type: ULONG_PTR
	dr1          uintptr         // C type: ULONG_PTR
	dr2          uintptr         // C type: ULONG_PTR
	dr3          uintptr         // C type: ULONG_PTR
	dr6          uintptr         // C type: ULONG_PTR
	dr7          uintptr         // C type: ULONG_PTR
	RegisterArea [80]byte        // C type: BYTE[80]
	x87fpu       X87FPU          // C type: X87FPU
	MxCsr        uint32          // C type: DWORD
	XmmRegisters [16]uintptr     // C type: XMMREGISTER[16]
	YmmRegisters [16]YMMREGISTER // C type: YMMREGISTER[16]
}

// LASTERROR (:771 )
type LASTERROR struct {
	code uint32    // C type: DWORD
	name [128]int8 // C type: char[128]
}

// LASTSTATUS (:777 )
type LASTSTATUS struct {
	code uint32    // C type: DWORD
	name [128]int8 // C type: char[128]
}

// REGDUMP (:783 )
type REGDUMP struct {
	regcontext           REGISTERCONTEXT   // C type: REGISTERCONTEXT
	flags                uintptr           // C type: FLAGS
	x87FPURegisters      [8]X87FPUREGISTER // C type: X87FPUREGISTER[8]
	mmx                  [8]uint64         // C type: unsigned long long[8]
	MxCsrFields          uintptr           // C type: MXCSRFIELDS
	x87StatusWordFields  uintptr           // C type: X87STATUSWORDFIELDS
	x87ControlWordFields uintptr           // C type: X87CONTROLWORDFIELDS
	lastError            LASTERROR         // C type: LASTERROR
	lastStatus           LASTSTATUS        // C type: LASTSTATUS
}

// DISASM_ARG (:796 )
type DISASM_ARG struct {
	Type     DISASM_ARGTYPE // C type: DISASM_ARGTYPE
	segment  SEGMENTREG     // C type: SEGMENTREG
	mnemonic [64]int8       // C type: char[64]
	constant uint           // C type: duint
	value    uint           // C type: duint
	memvalue uint           // C type: duint
}

// DISASM_INSTR (:806 )
type DISASM_INSTR struct {
	instruction [64]int8         // C type: char[64]
	Type        DISASM_INSTRTYPE // C type: DISASM_INSTRTYPE
	argcount    int              // C type: int
	instr_size  int              // C type: int
	arg         [3]DISASM_ARG    // C type: DISASM_ARG[3]
}

// STACK_COMMENT (:815 )
type STACK_COMMENT struct {
	color   [8]int8   // C type: char[8]
	comment [512]int8 // C type: char[512]
}

// THREADINFO (:821 )
type THREADINFO struct {
	ThreadNumber       int       // C type: int
	Handle             uintptr   // C type: HANDLE
	ThreadId           uint32    // C type: DWORD
	ThreadStartAddress uint      // C type: duint
	ThreadLocalBase    uint      // C type: duint
	threadName         [256]int8 // C type: char[256]
}

// THREADALLINFO (:831 )
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

// THREADLIST (:845 )
type THREADLIST struct {
	count         int            // C type: int
	list          *THREADALLINFO // C type: THREADALLINFO *
	CurrentThread int            // C type: int
}

// MEMORY_INFO (:852 )
type MEMORY_INFO struct {
	value    uint        // C type: duint
	size     MEMORY_SIZE // C type: MEMORY_SIZE
	mnemonic [64]int8    // C type: char[64]
}

// VALUE_INFO (:859 )
type VALUE_INFO struct {
	value uint    // C type: duint
	size  uintptr // C type: VALUE_SIZE
}

// SCRIPTBRANCH (:881 )
type SCRIPTBRANCH struct {
	Type        SCRIPTBRANCHTYPE // C type: SCRIPTBRANCHTYPE
	dest        int              // C type: int
	branchlabel [256]int8        // C type: char[256]
}

// XREF_RECORD (:897 )
type XREF_RECORD struct {
	addr uint     // C type: duint
	Type XREFTYPE // C type: XREFTYPE
}

// XREF_INFO (:903 )
type XREF_INFO struct {
	refcount   uint         // C type: duint
	references *XREF_RECORD // C type: XREF_RECORD *
}

// XREF_EDGE (:909 )
type XREF_EDGE struct {
	address uint // C type: duint
	from    uint // C type: duint
}

// SYMBOLPTR_ (:915 )
type SYMBOLPTR_ struct {
	modbase uint    // C type: duint
	symbol  uintptr // C type: const void *
}

// CELLINFO (:1193 )
type CELLINFO struct {
	row int   // C type: int
	col int   // C type: int
	str *int8 // C type: const char *
}

// SELECTIONDATA (:1200 )
type SELECTIONDATA struct {
	start uint // C type: duint
	end   uint // C type: duint
}

// ICONDATA (:1206 )
type ICONDATA struct {
	data uintptr // C type: const void *
	size uint    // C type: duint
}

// SCRIPTTYPEINFO (:1212 )
type SCRIPTTYPEINFO struct {
	name            [64]int8 // C type: char[64]
	id              int      // C type: int
	execute         uintptr  // C type: GUISCRIPTEXECUTE
	completeCommand uintptr  // C type: GUISCRIPTCOMPLETER
}

// ACTIVEVIEW (:1220 )
type ACTIVEVIEW struct {
	titleHwnd uintptr   // C type: void *
	classHwnd uintptr   // C type: void *
	title     [512]int8 // C type: char[512]
	className [512]int8 // C type: char[512]
}
type bridgemain struct{}

// BridgeInit    c api name: BridgeInit
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ const wchar_t *      │ * rune  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeInit           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeInit           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeInit() *rune {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeInit").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: const wchar_t *")
}

// BridgeStart    c api name: BridgeStart
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ const wchar_t *      │ * rune  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeStart          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeStart          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeStart() *rune {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeStart").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: const wchar_t *")
}

// BridgeAlloc    c api name: BridgeAlloc
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ size                 │ size_t               │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeAlloc          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeAlloc          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeAlloc(size uint) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeAlloc").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	panic("not support return type: void *")
}

// BridgeFree    c api name: BridgeFree
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ ptr                  │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeFree           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeFree           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeFree(ptr uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeFree").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ptr",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", ptr),
			},
		},
	))).Request()
}

// BridgeSettingGet    c api name: BridgeSettingGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ section              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ key                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ value                │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingGet     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingGet     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingGet(section *int8, key *int8, value *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// BridgeSettingGetUint    c api name: BridgeSettingGetUint
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ section              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ key                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ value                │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingGetUint │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingGetUint │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingGetUint(section *int8, key *int8, value *uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingGetUint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// BridgeSettingSet    c api name: BridgeSettingSet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ section              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ key                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ value                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingSet     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingSet     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingSet(section *int8, key *int8, value *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingSet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// BridgeSettingSetUint    c api name: BridgeSettingSetUint
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ section              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ key                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ value                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingSetUint │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingSetUint │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingSetUint(section *int8, key *int8, value uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingSetUint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// BridgeSettingFlush    c api name: BridgeSettingFlush
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingFlush   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingFlush   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingFlush() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingFlush").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// BridgeSettingRead    c api name: BridgeSettingRead
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ errorLine            │ int *                │ *int    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingRead    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingRead    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingRead(errorLine *int) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingRead").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "errorLine",
				Type:  "*int ",
				Value: fmt.Sprintf("%v", errorLine),
			},
		},
	))).Request()
	return true
}

// BridgeGetDbgVersion    c api name: BridgeGetDbgVersion
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeGetDbgVersion  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeGetDbgVersion  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeGetDbgVersion() int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeGetDbgVersion").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return 0
}

// BridgeIsProcessElevated    c api name: BridgeIsProcessElevated
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ bool                    │ bool    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ BridgeIsProcessElevated │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeIsProcessElevated │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) BridgeIsProcessElevated() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeIsProcessElevated").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// BridgeGetNtBuildNumber    c api name: BridgeGetNtBuildNumber
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ unsigned int           │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ BridgeGetNtBuildNumber │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeGetNtBuildNumber │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) BridgeGetNtBuildNumber() uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeGetNtBuildNumber").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// BridgeUserDirectory    c api name: BridgeUserDirectory
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ const wchar_t *      │ * rune  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeUserDirectory  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeUserDirectory  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeUserDirectory() *rune {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeUserDirectory").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: const wchar_t *")
}

// BridgeIsARM64Emulated    c api name: BridgeIsARM64Emulated
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ bool                  │ bool    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ BridgeIsARM64Emulated │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeIsARM64Emulated │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) BridgeIsARM64Emulated() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeIsARM64Emulated").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// DbgInit    c api name: DbgInit
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgInit              │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgInit              │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgInit() *int8 {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgInit").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: const char *")
}

// DbgExit    c api name: DbgExit
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgExit              │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgExit              │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgExit() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgExit").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// DbgMemRead    c api name: DbgMemRead
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ va                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ dest                 │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ size                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemRead           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemRead           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemRead(va uint, dest uintptr, size uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemRead").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "va",
				Type:  "uint",
				Value: fmt.Sprintf("%v", va),
			},
			{
				Name:  "dest",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", dest),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	return true
}

// DbgMemWrite    c api name: DbgMemWrite
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ va                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ src                  │ const void *         │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ size                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemWrite          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemWrite          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemWrite(va uint, src uintptr, size uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemWrite").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "va",
				Type:  "uint",
				Value: fmt.Sprintf("%v", va),
			},
			{
				Name:  "src",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", src),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	return true
}

// DbgMemGetPageSize    c api name: DbgMemGetPageSize
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ base                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemGetPageSize    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemGetPageSize    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemGetPageSize(base uint) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemGetPageSize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "base",
				Type:  "uint",
				Value: fmt.Sprintf("%v", base),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgMemFindBaseAddr    c api name: DbgMemFindBaseAddr
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ size                 │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemFindBaseAddr   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemFindBaseAddr   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemFindBaseAddr(addr uint, size *uint) uint {
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
	panic("not support return type: duint")
}

// DbgCmdExec    c api name: DbgCmdExec
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ cmd                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgCmdExec           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgCmdExec           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgCmdExec(cmd *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgCmdExec").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
	return true
}

// DbgCmdExecDirect    c api name: DbgCmdExecDirect
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ cmd                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgCmdExecDirect     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgCmdExecDirect     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgCmdExecDirect(cmd *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgCmdExecDirect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
	return true
}

// DbgMemMap    c api name: DbgMemMap
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ memmap               │ MEMMAP *             │ *MEMMAP │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemMap            │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemMap            │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemMap(memmap *MEMMAP) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemMap").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "memmap",
				Type:  "*MEMMAP ",
				Value: fmt.Sprintf("%v", memmap),
			},
		},
	))).Request()
	return true
}

// DbgIsValidExpression    c api name: DbgIsValidExpression
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ expression           │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsValidExpression │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsValidExpression │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgIsValidExpression(expression *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsValidExpression").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "expression",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", expression),
			},
		},
	))).Request()
	return true
}

// DbgIsDebugging    c api name: DbgIsDebugging
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsDebugging       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsDebugging       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgIsDebugging() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsDebugging").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// DbgIsJumpGoingToExecute    c api name: DbgIsJumpGoingToExecute
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ bool                    │ bool    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsJumpGoingToExecute │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsJumpGoingToExecute │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) DbgIsJumpGoingToExecute(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsJumpGoingToExecute").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgGetLabelAt    c api name: DbgGetLabelAt
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ addr                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ segment              │ SEGMENTREG           │ SEGMENTREG │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 2  │ text                 │ char *               │ *int8      │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ bool                 │ bool       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgGetLabelAt        │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgGetLabelAt        │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgGetLabelAt(addr uint, segment SEGMENTREG, text *int8) bool {
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
	return true
}

// DbgSetLabelAt    c api name: DbgSetLabelAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetLabelAt        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetLabelAt        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetLabelAt(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetLabelAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
	return true
}

// DbgClearLabelRange    c api name: DbgClearLabelRange
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearLabelRange   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearLabelRange   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgGetCommentAt    c api name: DbgGetCommentAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetCommentAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetCommentAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetCommentAt(addr uint, text *int8) bool {
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
	return true
}

// DbgSetCommentAt    c api name: DbgSetCommentAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetCommentAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetCommentAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetCommentAt(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetCommentAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
	return true
}

// DbgClearCommentRange    c api name: DbgClearCommentRange
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearCommentRange │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearCommentRange │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgGetBookmarkAt    c api name: DbgGetBookmarkAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetBookmarkAt     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetBookmarkAt     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetBookmarkAt(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBookmarkAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgClearBookmarkRange    c api name: DbgClearBookmarkRange
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ start                │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 1  │ end                  │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearBookmarkRange │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearBookmarkRange │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
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
}

// DbgGetModuleAt    c api name: DbgGetModuleAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetModuleAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetModuleAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetModuleAt(addr uint, text *int8) bool {
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
	return true
}

// DbgGetBpxTypeAt    c api name: DbgGetBpxTypeAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ BPXTYPE              │ BPXTYPE │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetBpxTypeAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetBpxTypeAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetBpxTypeAt(addr uint) BPXTYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBpxTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: BPXTYPE")
}

// DbgValFromString    c api name: DbgValFromString
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ s                    │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgValFromString     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgValFromString     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgValFromString(s *int8) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgValFromString").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "s",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", s),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgGetRegDumpEx    c api name: DbgGetRegDumpEx
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ regdump              │ REGDUMP *            │ *REGDUMP │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ size                 │ size_t               │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ bool                 │ bool     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetRegDumpEx      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetRegDumpEx      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetRegDumpEx(regdump *REGDUMP, size uint) bool {
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
	return true
}

// DbgValToString    c api name: DbgValToString
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ s                    │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ value                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgValToString       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgValToString       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgValToString(s *int8, value uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgValToString").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "s",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", s),
			},
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// DbgMemIsValidReadPtr    c api name: DbgMemIsValidReadPtr
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemIsValidReadPtr │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemIsValidReadPtr │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemIsValidReadPtr(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemIsValidReadPtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgGetBpList    c api name: DbgGetBpList
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ Type                 │ BPXTYPE              │ BPXTYPE │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ list                 │ BPMAP *              │ *BPMAP  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetBpList         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetBpList         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetBpList(Type BPXTYPE, list *BPMAP) int {
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
	return 0
}

// DbgGetFunctionTypeAt    c api name: DbgGetFunctionTypeAt
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ FUNCTYPE             │ FUNCTYPE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetFunctionTypeAt │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetFunctionTypeAt │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetFunctionTypeAt(addr uint) FUNCTYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetFunctionTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: FUNCTYPE")
}

// DbgGetLoopTypeAt    c api name: DbgGetLoopTypeAt
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ depth                │ int                  │ int      │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ LOOPTYPE             │ LOOPTYPE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetLoopTypeAt     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetLoopTypeAt     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetLoopTypeAt(addr uint, depth int) LOOPTYPE {
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
	panic("not support return type: LOOPTYPE")
}

// DbgGetBranchDestination    c api name: DbgGetBranchDestination
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetBranchDestination │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetBranchDestination │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) DbgGetBranchDestination(addr uint) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBranchDestination").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgScriptLoad    c api name: DbgScriptLoad
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ filename             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptLoad        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptLoad        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptLoad(filename *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptLoad").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
}

// DbgScriptUnload    c api name: DbgScriptUnload
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptUnload      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptUnload      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptUnload() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptUnload").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// DbgScriptRun    c api name: DbgScriptRun
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ destline             │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptRun         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptRun         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgScriptStep    c api name: DbgScriptStep
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptStep        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptStep        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptStep() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptStep").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// DbgScriptBpToggle    c api name: DbgScriptBpToggle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptBpToggle    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptBpToggle    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptBpToggle(line int) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptBpToggle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	return true
}

// DbgScriptBpGet    c api name: DbgScriptBpGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptBpGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptBpGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptBpGet(line int) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptBpGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	return true
}

// DbgScriptCmdExec    c api name: DbgScriptCmdExec
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ command              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptCmdExec     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptCmdExec     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptCmdExec(command *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptCmdExec").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "command",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", command),
			},
		},
	))).Request()
	return true
}

// DbgScriptAbort    c api name: DbgScriptAbort
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptAbort       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptAbort       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptAbort() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptAbort").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// DbgScriptGetLineType    c api name: DbgScriptGetLineType
// ┌────┬──────────────────────┬──────────────────────┬────────────────┐
// │ id │         name         │        c type        │    go type     │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │ 0  │ line                 │ int                  │ int            │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ return               │ SCRIPTLINETYPE       │ SCRIPTLINETYPE │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.name         │ DbgScriptGetLineType │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.mangledName  │ DbgScriptGetLineType │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.file         │                      │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.includedFrom │                      │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.expansionLoc │                      │                │
// └────┴──────────────────────┴──────────────────────┴────────────────┘
func (b *bridgemain) DbgScriptGetLineType(line int) SCRIPTLINETYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptGetLineType").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	panic("not support return type: SCRIPTLINETYPE")
}

// DbgScriptSetIp    c api name: DbgScriptSetIp
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptSetIp       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptSetIp       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgScriptGetBranchInfo    c api name: DbgScriptGetBranchInfo
// ┌────┬──────────────────────┬────────────────────────┬───────────────┐
// │ id │         name         │         c type         │    go type    │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │ 0  │ line                 │ int                    │ int           │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │ 1  │ info                 │ SCRIPTBRANCH *         │ *SCRIPTBRANCH │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ return               │ bool                   │ bool          │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.name         │ DbgScriptGetBranchInfo │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.mangledName  │ DbgScriptGetBranchInfo │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.file         │                        │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.includedFrom │                        │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.expansionLoc │                        │               │
// └────┴──────────────────────┴────────────────────────┴───────────────┘
func (b *bridgemain) DbgScriptGetBranchInfo(line int, info *SCRIPTBRANCH) bool {
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
	return true
}

// DbgSymbolEnum    c api name: DbgSymbolEnum
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ base                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ cbSymbolEnum         │ CBSYMBOLENUM         │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ user                 │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSymbolEnum        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSymbolEnum        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSymbolEnum(base uint, cbSymbolEnum uintptr, user uintptr) bool {
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
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	return true
}

// DbgSymbolEnumFromCache    c api name: DbgSymbolEnumFromCache
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ base                 │ duint                  │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ cbSymbolEnum         │ CBSYMBOLENUM           │ uintptr │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 2  │ user                 │ void *                 │ uintptr │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ bool                   │ bool    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ DbgSymbolEnumFromCache │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSymbolEnumFromCache │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) DbgSymbolEnumFromCache(base uint, cbSymbolEnum uintptr, user uintptr) bool {
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
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	return true
}

// DbgSymbolEnumRange    c api name: DbgSymbolEnumRange
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ symbolMask           │ unsigned int         │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 3  │ cbSymbolEnum         │ CBSYMBOLENUM         │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 4  │ user                 │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSymbolEnumRange   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSymbolEnumRange   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSymbolEnumRange(start uint, end uint, symbolMask uint, cbSymbolEnum uintptr, user uintptr) bool {
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
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	return true
}

// DbgAssembleAt    c api name: DbgAssembleAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ instruction          │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgAssembleAt        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgAssembleAt        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgAssembleAt(addr uint, instruction *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgAssembleAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "instruction",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", instruction),
			},
		},
	))).Request()
	return true
}

// DbgModBaseFromName    c api name: DbgModBaseFromName
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ name                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgModBaseFromName   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgModBaseFromName   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgModBaseFromName(name *int8) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgModBaseFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgDisasmAt    c api name: DbgDisasmAt
// ┌────┬──────────────────────┬──────────────────────┬───────────────┐
// │ id │         name         │        c type        │    go type    │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │ 0  │ addr                 │ duint                │ uint          │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │ 1  │ instr                │ DISASM_INSTR *       │ *DISASM_INSTR │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ return               │ void                 │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.name         │ DbgDisasmAt          │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.mangledName  │ DbgDisasmAt          │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.file         │                      │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.includedFrom │                      │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.expansionLoc │                      │               │
// └────┴──────────────────────┴──────────────────────┴───────────────┘
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
}

// DbgStackCommentGet    c api name: DbgStackCommentGet
// ┌────┬──────────────────────┬──────────────────────┬────────────────┐
// │ id │         name         │        c type        │    go type     │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │ 0  │ addr                 │ duint                │ uint           │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │ 1  │ comment              │ STACK_COMMENT *      │ *STACK_COMMENT │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ return               │ bool                 │ bool           │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.name         │ DbgStackCommentGet   │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.mangledName  │ DbgStackCommentGet   │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.file         │                      │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.includedFrom │                      │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.expansionLoc │                      │                │
// └────┴──────────────────────┴──────────────────────┴────────────────┘
func (b *bridgemain) DbgStackCommentGet(addr uint, comment *STACK_COMMENT) bool {
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
	return true
}

// DbgGetThreadList    c api name: DbgGetThreadList
// ┌────┬──────────────────────┬──────────────────────┬─────────────┐
// │ id │         name         │        c type        │   go type   │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │ 0  │ list                 │ THREADLIST *         │ *THREADLIST │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ return               │ void                 │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.name         │ DbgGetThreadList     │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.mangledName  │ DbgGetThreadList     │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.file         │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.includedFrom │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.expansionLoc │                      │             │
// └────┴──────────────────────┴──────────────────────┴─────────────┘
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
}

// DbgSettingsUpdated    c api name: DbgSettingsUpdated
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSettingsUpdated   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSettingsUpdated   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSettingsUpdated() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSettingsUpdated").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// DbgDisasmFastAt    c api name: DbgDisasmFastAt
// ┌────┬──────────────────────┬──────────────────────────┬──────────┐
// │ id │         name         │          c type          │ go type  │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                    │ uint     │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │ 1  │ basicinfo            │ BASIC_INSTRUCTION_INFO * │ *uintptr │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ return               │ void                     │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h     │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.name         │ DbgDisasmFastAt          │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgDisasmFastAt          │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.file         │                          │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.includedFrom │                          │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.expansionLoc │                          │          │
// └────┴──────────────────────┴──────────────────────────┴──────────┘
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
}

// DbgMenuEntryClicked    c api name: DbgMenuEntryClicked
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hEntry               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMenuEntryClicked  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMenuEntryClicked  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgFunctionGet    c api name: DbgFunctionGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ start                │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ end                  │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgFunctionGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgFunctionGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgFunctionGet(addr uint, start *uint, end *uint) bool {
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
	return true
}

// DbgFunctionOverlaps    c api name: DbgFunctionOverlaps
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgFunctionOverlaps  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgFunctionOverlaps  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgFunctionOverlaps(start uint, end uint) bool {
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
	return true
}

// DbgFunctionAdd    c api name: DbgFunctionAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgFunctionAdd       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgFunctionAdd       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgFunctionAdd(start uint, end uint) bool {
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
	return true
}

// DbgFunctionDel    c api name: DbgFunctionDel
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgFunctionDel       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgFunctionDel       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgFunctionDel(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgFunctionDel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgArgumentGet    c api name: DbgArgumentGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ start                │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ end                  │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgArgumentGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgArgumentGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgArgumentGet(addr uint, start *uint, end *uint) bool {
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
	return true
}

// DbgArgumentOverlaps    c api name: DbgArgumentOverlaps
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgArgumentOverlaps  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgArgumentOverlaps  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgArgumentOverlaps(start uint, end uint) bool {
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
	return true
}

// DbgArgumentAdd    c api name: DbgArgumentAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgArgumentAdd       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgArgumentAdd       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgArgumentAdd(start uint, end uint) bool {
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
	return true
}

// DbgArgumentDel    c api name: DbgArgumentDel
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgArgumentDel       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgArgumentDel       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgArgumentDel(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgArgumentDel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgLoopGet    c api name: DbgLoopGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ depth                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ start                │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 3  │ end                  │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgLoopGet           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgLoopGet           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgLoopGet(depth int, addr uint, start *uint, end *uint) bool {
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
	return true
}

// DbgLoopOverlaps    c api name: DbgLoopOverlaps
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ depth                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgLoopOverlaps      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgLoopOverlaps      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgLoopOverlaps(depth int, start uint, end uint) bool {
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
	return true
}

// DbgLoopAdd    c api name: DbgLoopAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgLoopAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgLoopAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgLoopAdd(start uint, end uint) bool {
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
	return true
}

// DbgLoopDel    c api name: DbgLoopDel
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ depth                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgLoopDel           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgLoopDel           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgLoopDel(depth int, addr uint) bool {
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
	return true
}

// DbgXrefAdd    c api name: DbgXrefAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ from                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgXrefAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgXrefAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgXrefAdd(addr uint, from uint) bool {
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
	return true
}

// DbgXrefDelAll    c api name: DbgXrefDelAll
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgXrefDelAll        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgXrefDelAll        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgXrefDelAll(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgXrefDelAll").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgXrefGet    c api name: DbgXrefGet
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ addr                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ info                 │ XREF_INFO *          │ *XREF_INFO │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ bool                 │ bool       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgXrefGet           │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgXrefGet           │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgXrefGet(addr uint, info *XREF_INFO) bool {
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
	return true
}

// DbgGetXrefCountAt    c api name: DbgGetXrefCountAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ size_t               │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetXrefCountAt    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetXrefCountAt    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetXrefCountAt(addr uint) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetXrefCountAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: size_t")
}

// DbgGetXrefTypeAt    c api name: DbgGetXrefTypeAt
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ XREFTYPE             │ XREFTYPE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetXrefTypeAt     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetXrefTypeAt     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetXrefTypeAt(addr uint) XREFTYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetXrefTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: XREFTYPE")
}

// DbgIsRunLocked    c api name: DbgIsRunLocked
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsRunLocked       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsRunLocked       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgIsRunLocked() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsRunLocked").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// DbgIsBpDisabled    c api name: DbgIsBpDisabled
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsBpDisabled      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsBpDisabled      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgIsBpDisabled(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsBpDisabled").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgSetAutoCommentAt    c api name: DbgSetAutoCommentAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetAutoCommentAt  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetAutoCommentAt  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetAutoCommentAt(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoCommentAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
	return true
}

// DbgClearAutoCommentRange    c api name: DbgClearAutoCommentRange
// ┌────┬──────────────────────┬──────────────────────────┬─────────┐
// │ id │         name         │          c type          │ go type │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │ 0  │ start                │ duint                    │ uint    │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │ 1  │ end                  │ duint                    │ uint    │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ return               │ void                     │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h     │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearAutoCommentRange │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearAutoCommentRange │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.file         │                          │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.includedFrom │                          │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                          │         │
// └────┴──────────────────────┴──────────────────────────┴─────────┘
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
}

// DbgSetAutoLabelAt    c api name: DbgSetAutoLabelAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetAutoLabelAt    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetAutoLabelAt    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetAutoLabelAt(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoLabelAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
	return true
}

// DbgClearAutoLabelRange    c api name: DbgClearAutoLabelRange
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ start                │ duint                  │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ end                  │ duint                  │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearAutoLabelRange │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearAutoLabelRange │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
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
}

// DbgSetAutoBookmarkAt    c api name: DbgSetAutoBookmarkAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetAutoBookmarkAt │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetAutoBookmarkAt │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetAutoBookmarkAt(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoBookmarkAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgClearAutoBookmarkRange    c api name: DbgClearAutoBookmarkRange
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ start                │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 1  │ end                  │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ void                      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearAutoBookmarkRange │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearAutoBookmarkRange │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
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
}

// DbgSetAutoFunctionAt    c api name: DbgSetAutoFunctionAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetAutoFunctionAt │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetAutoFunctionAt │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetAutoFunctionAt(start uint, end uint) bool {
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
	return true
}

// DbgClearAutoFunctionRange    c api name: DbgClearAutoFunctionRange
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ start                │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 1  │ end                  │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ void                      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearAutoFunctionRange │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearAutoFunctionRange │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
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
}

// DbgGetStringAt    c api name: DbgGetStringAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetStringAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetStringAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetStringAt(addr uint, text *int8) bool {
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
	return true
}

// DbgFunctions    c api name: DbgFunctions
// ┌────┬──────────────────────┬──────────────────────┬───────────────┐
// │ id │         name         │        c type        │    go type    │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ return               │ const DBGFUNCTIONS * │ *DBGFUNCTIONS │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.name         │ DbgFunctions         │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.mangledName  │ DbgFunctions         │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.file         │                      │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.includedFrom │                      │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.expansionLoc │                      │               │
// └────┴──────────────────────┴──────────────────────┴───────────────┘
func (b *bridgemain) DbgFunctions() *DBGFUNCTIONS {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgFunctions").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: const DBGFUNCTIONS *")
}

// DbgWinEvent    c api name: DbgWinEvent
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ message              │ MSG *                │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ result               │ long *               │ *int32   │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ bool                 │ bool     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgWinEvent          │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgWinEvent          │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgWinEvent(message *uintptr, result *int32) bool {
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
	return true
}

// DbgWinEventGlobal    c api name: DbgWinEventGlobal
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ message              │ MSG *                │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ bool                 │ bool     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgWinEventGlobal    │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgWinEventGlobal    │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgWinEventGlobal(message *uintptr) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgWinEventGlobal").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	return true
}

// DbgIsRunning    c api name: DbgIsRunning
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsRunning         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsRunning         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgIsRunning() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsRunning").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// DbgGetTimeWastedCounter    c api name: DbgGetTimeWastedCounter
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetTimeWastedCounter │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetTimeWastedCounter │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) DbgGetTimeWastedCounter() uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetTimeWastedCounter").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// DbgGetArgTypeAt    c api name: DbgGetArgTypeAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ ARGTYPE              │ ARGTYPE │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetArgTypeAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetArgTypeAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetArgTypeAt(addr uint) ARGTYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetArgTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: ARGTYPE")
}

// DbgGetEncodeTypeBuffer    c api name: DbgGetEncodeTypeBuffer
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                  │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ size                 │ duint *                │ *uint   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void *                 │ uintptr │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetEncodeTypeBuffer │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetEncodeTypeBuffer │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) DbgGetEncodeTypeBuffer(addr uint, size *uint) uintptr {
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
	panic("not support return type: void *")
}

// DbgReleaseEncodeTypeBuffer    c api name: DbgReleaseEncodeTypeBuffer
// ┌────┬──────────────────────┬────────────────────────────┬─────────┐
// │ id │         name         │           c type           │ go type │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 0  │ buffer               │ void *                     │ uintptr │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ return               │ void                       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.name         │ DbgReleaseEncodeTypeBuffer │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgReleaseEncodeTypeBuffer │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.file         │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                            │         │
// └────┴──────────────────────┴────────────────────────────┴─────────┘
func (b *bridgemain) DbgReleaseEncodeTypeBuffer(buffer uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgReleaseEncodeTypeBuffer").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "buffer",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", buffer),
			},
		},
	))).Request()
}

// DbgGetEncodeTypeAt    c api name: DbgGetEncodeTypeAt
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ addr                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ size                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ ENCODETYPE           │ ENCODETYPE │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgGetEncodeTypeAt   │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgGetEncodeTypeAt   │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgGetEncodeTypeAt(addr uint, size uint) ENCODETYPE {
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
	panic("not support return type: ENCODETYPE")
}

// DbgGetEncodeSizeAt    c api name: DbgGetEncodeSizeAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ codesize             │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetEncodeSizeAt   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetEncodeSizeAt   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetEncodeSizeAt(addr uint, codesize uint) uint {
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
	panic("not support return type: duint")
}

// DbgSetEncodeType    c api name: DbgSetEncodeType
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ addr                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ size                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 2  │ Type                 │ ENCODETYPE           │ ENCODETYPE │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ bool                 │ bool       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgSetEncodeType     │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgSetEncodeType     │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgSetEncodeType(addr uint, size uint, Type ENCODETYPE) bool {
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
	return true
}

// DbgDelEncodeTypeRange    c api name: DbgDelEncodeTypeRange
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ start                │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 1  │ end                  │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ DbgDelEncodeTypeRange │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgDelEncodeTypeRange │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
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
}

// DbgDelEncodeTypeSegment    c api name: DbgDelEncodeTypeSegment
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ start                │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ DbgDelEncodeTypeSegment │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgDelEncodeTypeSegment │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
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
}

// DbgGetWatchList    c api name: DbgGetWatchList
// ┌────┬──────────────────────┬──────────────────────┬───────────┐
// │ id │         name         │        c type        │  go type  │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 0  │ list                 │ ListInfo *           │ *ListInfo │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ return               │ bool                 │ bool      │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.name         │ DbgGetWatchList      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.mangledName  │ DbgGetWatchList      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.file         │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.includedFrom │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.expansionLoc │                      │           │
// └────┴──────────────────────┴──────────────────────┴───────────┘
func (b *bridgemain) DbgGetWatchList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetWatchList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	return true
}

// DbgSelChanged    c api name: DbgSelChanged
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hWindow              │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ VA                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSelChanged        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSelChanged        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgGetProcessHandle    c api name: DbgGetProcessHandle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ HANDLE               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetProcessHandle  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetProcessHandle  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetProcessHandle() uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetProcessHandle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: HANDLE")
}

// DbgGetThreadHandle    c api name: DbgGetThreadHandle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ HANDLE               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetThreadHandle   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetThreadHandle   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetThreadHandle() uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetThreadHandle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: HANDLE")
}

// DbgGetProcessId    c api name: DbgGetProcessId
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ DWORD                │ uint32  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetProcessId      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetProcessId      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetProcessId() uint32 {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetProcessId").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: DWORD")
}

// DbgGetThreadId    c api name: DbgGetThreadId
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ DWORD                │ uint32  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetThreadId       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetThreadId       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetThreadId() uint32 {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetThreadId").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: DWORD")
}

// DbgGetPebAddress    c api name: DbgGetPebAddress
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ ProcessId            │ DWORD                │ uint32  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetPebAddress     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetPebAddress     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetPebAddress(ProcessId uint32) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetPebAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ProcessId",
				Type:  "uint32",
				Value: fmt.Sprintf("%v", ProcessId),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgGetTebAddress    c api name: DbgGetTebAddress
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ ThreadId             │ DWORD                │ uint32  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetTebAddress     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetTebAddress     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetTebAddress(ThreadId uint32) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetTebAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ThreadId",
				Type:  "uint32",
				Value: fmt.Sprintf("%v", ThreadId),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgAnalyzeFunction    c api name: DbgAnalyzeFunction
// ┌────┬──────────────────────┬──────────────────────┬────────────────────┐
// │ id │         name         │        c type        │      go type       │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 0  │ entry                │ duint                │ uint               │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 1  │ graph                │ BridgeCFGraphList *  │ *BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ return               │ bool                 │ bool               │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.name         │ DbgAnalyzeFunction   │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.mangledName  │ DbgAnalyzeFunction   │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.file         │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.includedFrom │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.expansionLoc │                      │                    │
// └────┴──────────────────────┴──────────────────────┴────────────────────┘
func (b *bridgemain) DbgAnalyzeFunction(entry uint, graph *BridgeCFGraphList) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgAnalyzeFunction").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "entry",
				Type:  "uint",
				Value: fmt.Sprintf("%v", entry),
			},
			{
				Name:  "graph",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graph),
			},
		},
	))).Request()
	return true
}

// DbgGetSymbolInfo    c api name: DbgGetSymbolInfo
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ symbolptr            │ const SYMBOLPTR *    │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ info                 │ SYMBOLINFO *         │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ void                 │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetSymbolInfo     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetSymbolInfo     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
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
}

// DbgGetDebugEngine    c api name: DbgGetDebugEngine
// ┌────┬──────────────────────┬──────────────────────┬──────────────┐
// │ id │         name         │        c type        │   go type    │
// ├────┼──────────────────────┼──────────────────────┼──────────────┤
// │    │ return               │ DEBUG_ENGINE         │ DEBUG_ENGINE │
// ├────┼──────────────────────┼──────────────────────┼──────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │              │
// ├────┼──────────────────────┼──────────────────────┼──────────────┤
// │    │ comment.name         │ DbgGetDebugEngine    │              │
// ├────┼──────────────────────┼──────────────────────┼──────────────┤
// │    │ comment.mangledName  │ DbgGetDebugEngine    │              │
// ├────┼──────────────────────┼──────────────────────┼──────────────┤
// │    │ comment.file         │                      │              │
// ├────┼──────────────────────┼──────────────────────┼──────────────┤
// │    │ comment.includedFrom │                      │              │
// ├────┼──────────────────────┼──────────────────────┼──────────────┤
// │    │ comment.expansionLoc │                      │              │
// └────┴──────────────────────┴──────────────────────┴──────────────┘
func (b *bridgemain) DbgGetDebugEngine() DEBUG_ENGINE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetDebugEngine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: DEBUG_ENGINE")
}

// DbgGetSymbolInfoAt    c api name: DbgGetSymbolInfoAt
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ info                 │ SYMBOLINFO *         │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ bool                 │ bool     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetSymbolInfoAt   │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetSymbolInfoAt   │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetSymbolInfoAt(addr uint, info *uintptr) bool {
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
	return true
}

// DbgXrefAddMulti    c api name: DbgXrefAddMulti
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ edges                │ const XREF_EDGE *    │ *XREF_EDGE │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ count                │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgXrefAddMulti      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgXrefAddMulti      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgXrefAddMulti(edges *XREF_EDGE, count uint) uint {
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
	panic("not support return type: duint")
}

// DbgMenuPrepare    c api name: DbgMenuPrepare
// ┌────┬──────────────────────┬──────────────────────┬─────────────┐
// │ id │         name         │        c type        │   go type   │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │ 0  │ hMenu                │ GUIMENUTYPE          │ GUIMENUTYPE │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ return               │ void                 │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.name         │ DbgMenuPrepare       │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.mangledName  │ DbgMenuPrepare       │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.file         │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.includedFrom │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.expansionLoc │                      │             │
// └────┴──────────────────────┴──────────────────────┴─────────────┘
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
}

// GuiTranslateText    c api name: GuiTranslateText
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ Source               │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiTranslateText     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiTranslateText     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiTranslateText(Source *int8) *int8 {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiTranslateText").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "Source",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", Source),
			},
		},
	))).Request()
	panic("not support return type: const char *")
}

// GuiDisasmAt    c api name: GuiDisasmAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ cip                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDisasmAt          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDisasmAt          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSetDebugState    c api name: GuiSetDebugState
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ state                │ DBGSTATE             │ DBGSTATE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ void                 │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ GuiSetDebugState     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ GuiSetDebugState     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
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
}

// GuiSetDebugStateFast    c api name: GuiSetDebugStateFast
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ state                │ DBGSTATE             │ DBGSTATE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ void                 │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ GuiSetDebugStateFast │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ GuiSetDebugStateFast │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
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
}

// GuiAddLogMessage    c api name: GuiAddLogMessage
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ msg                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddLogMessage     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddLogMessage     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddLogMessage(msg *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddLogMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
}

// GuiAddLogMessageHtml    c api name: GuiAddLogMessageHtml
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ msg                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddLogMessageHtml │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddLogMessageHtml │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddLogMessageHtml(msg *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddLogMessageHtml").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
}

// GuiLogClear    c api name: GuiLogClear
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiLogClear          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiLogClear          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiLogClear() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiLogSave    c api name: GuiLogSave
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ filename             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiLogSave           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiLogSave           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiLogSave(filename *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogSave").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
}

// GuiLogRedirect    c api name: GuiLogRedirect
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ filename             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiLogRedirect       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiLogRedirect       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiLogRedirect(filename *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogRedirect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
}

// GuiLogRedirectStop    c api name: GuiLogRedirectStop
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiLogRedirectStop   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiLogRedirectStop   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiLogRedirectStop() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogRedirectStop").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateAllViews    c api name: GuiUpdateAllViews
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateAllViews    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateAllViews    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateAllViews() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateAllViews").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateRegisterView    c api name: GuiUpdateRegisterView
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateRegisterView │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateRegisterView │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateRegisterView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateRegisterView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateDisassemblyView    c api name: GuiUpdateDisassemblyView
// ┌────┬──────────────────────┬──────────────────────────┬─────────┐
// │ id │         name         │          c type          │ go type │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ return               │ void                     │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h     │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateDisassemblyView │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateDisassemblyView │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.file         │                          │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.includedFrom │                          │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                          │         │
// └────┴──────────────────────┴──────────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateDisassemblyView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateDisassemblyView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateBreakpointsView    c api name: GuiUpdateBreakpointsView
// ┌────┬──────────────────────┬──────────────────────────┬─────────┐
// │ id │         name         │          c type          │ go type │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ return               │ void                     │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h     │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateBreakpointsView │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateBreakpointsView │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.file         │                          │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.includedFrom │                          │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                          │         │
// └────┴──────────────────────┴──────────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateBreakpointsView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateBreakpointsView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateWindowTitle    c api name: GuiUpdateWindowTitle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ filename             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateWindowTitle │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateWindowTitle │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateWindowTitle(filename *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateWindowTitle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
}

// GuiGetWindowHandle    c api name: GuiGetWindowHandle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ HWND                 │ HWND    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetWindowHandle   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetWindowHandle   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiGetWindowHandle() HWND {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetWindowHandle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: HWND")
}

// GuiDumpAt    c api name: GuiDumpAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ va                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDumpAt            │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDumpAt            │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiScriptAdd    c api name: GuiScriptAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ count                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ lines                │ const char **        │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptAdd         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptAdd         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptAdd(count int, lines *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "count",
				Type:  "int",
				Value: fmt.Sprintf("%v", count),
			},
			{
				Name:  "lines",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", lines),
			},
		},
	))).Request()
}

// GuiScriptClear    c api name: GuiScriptClear
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptClear       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptClear       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptClear() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiScriptSetIp    c api name: GuiScriptSetIp
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptSetIp       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptSetIp       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiScriptError    c api name: GuiScriptError
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ message              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptError       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptError       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptError(line int, message *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptError").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
}

// GuiScriptSetTitle    c api name: GuiScriptSetTitle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptSetTitle    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptSetTitle    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptSetTitle(title *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptSetTitle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
}

// GuiScriptSetInfoLine    c api name: GuiScriptSetInfoLine
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ info                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptSetInfoLine │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptSetInfoLine │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptSetInfoLine(line int, info *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptSetInfoLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
			{
				Name:  "info",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
}

// GuiScriptMessage    c api name: GuiScriptMessage
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ message              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptMessage     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptMessage     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptMessage(message *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
}

// GuiScriptMsgyn    c api name: GuiScriptMsgyn
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ message              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptMsgyn       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptMsgyn       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptMsgyn(message *int8) int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptMsgyn").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	return 0
}

// GuiSymbolLogAdd    c api name: GuiSymbolLogAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ message              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSymbolLogAdd      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSymbolLogAdd      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiSymbolLogAdd(message *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolLogAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
}

// GuiSymbolLogClear    c api name: GuiSymbolLogClear
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSymbolLogClear    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSymbolLogClear    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiSymbolLogClear() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolLogClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiSymbolSetProgress    c api name: GuiSymbolSetProgress
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ percent              │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSymbolSetProgress │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSymbolSetProgress │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSymbolUpdateModuleList    c api name: GuiSymbolUpdateModuleList
// ┌────┬──────────────────────┬───────────────────────────┬───────────────────┐
// │ id │         name         │          c type           │      go type      │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │ 0  │ count                │ int                       │ int               │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │ 1  │ modules              │ SYMBOLMODULEINFO *        │ *SYMBOLMODULEINFO │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ return               │ void                      │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.name         │ GuiSymbolUpdateModuleList │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.mangledName  │ GuiSymbolUpdateModuleList │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.file         │                           │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.includedFrom │                           │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.expansionLoc │                           │                   │
// └────┴──────────────────────┴───────────────────────────┴───────────────────┘
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
}

// GuiSymbolRefreshCurrent    c api name: GuiSymbolRefreshCurrent
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiSymbolRefreshCurrent │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSymbolRefreshCurrent │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) GuiSymbolRefreshCurrent() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolRefreshCurrent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiReferenceAddColumn    c api name: GuiReferenceAddColumn
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ width                │ int                   │ int     │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 1  │ title                │ const char *          │ *int8   │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceAddColumn │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceAddColumn │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceAddColumn(width int, title *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceAddColumn").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "width",
				Type:  "int",
				Value: fmt.Sprintf("%v", width),
			},
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
}

// GuiReferenceSetRowCount    c api name: GuiReferenceSetRowCount
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ count                │ int                     │ int     │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetRowCount │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetRowCount │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
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
}

// GuiReferenceGetRowCount    c api name: GuiReferenceGetRowCount
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ int                     │ int     │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceGetRowCount │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceGetRowCount │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceGetRowCount() int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceGetRowCount").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return 0
}

// GuiReferenceSearchGetRowCount    c api name: GuiReferenceSearchGetRowCount
// ┌────┬──────────────────────┬───────────────────────────────┬─────────┐
// │ id │         name         │            c type             │ go type │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ return               │ int                           │ int     │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h          │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSearchGetRowCount │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSearchGetRowCount │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.file         │                               │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                               │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                               │         │
// └────┴──────────────────────┴───────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceSearchGetRowCount() int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSearchGetRowCount").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return 0
}

// GuiReferenceDeleteAllColumns    c api name: GuiReferenceDeleteAllColumns
// ┌────┬──────────────────────┬──────────────────────────────┬─────────┐
// │ id │         name         │            c type            │ go type │
// ├────┼──────────────────────┼──────────────────────────────┼─────────┤
// │    │ return               │ void                         │         │
// ├────┼──────────────────────┼──────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h         │         │
// ├────┼──────────────────────┼──────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceDeleteAllColumns │         │
// ├────┼──────────────────────┼──────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceDeleteAllColumns │         │
// ├────┼──────────────────────┼──────────────────────────────┼─────────┤
// │    │ comment.file         │                              │         │
// ├────┼──────────────────────┼──────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                              │         │
// ├────┼──────────────────────┼──────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                              │         │
// └────┴──────────────────────┴──────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceDeleteAllColumns() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceDeleteAllColumns").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiReferenceInitialize    c api name: GuiReferenceInitialize
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ name                 │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceInitialize │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceInitialize │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceInitialize(name *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceInitialize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
}

// GuiReferenceSetCellContent    c api name: GuiReferenceSetCellContent
// ┌────┬──────────────────────┬────────────────────────────┬─────────┐
// │ id │         name         │           c type           │ go type │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 0  │ row                  │ int                        │ int     │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 1  │ col                  │ int                        │ int     │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 2  │ str                  │ const char *               │ *int8   │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ return               │ void                       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetCellContent │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetCellContent │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.file         │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                            │         │
// └────┴──────────────────────┴────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceSetCellContent(row int, col int, str *int8) {
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
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", str),
			},
		},
	))).Request()
}

// GuiReferenceGetCellContent    c api name: GuiReferenceGetCellContent
// ┌────┬──────────────────────┬────────────────────────────┬─────────┐
// │ id │         name         │           c type           │ go type │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 0  │ row                  │ int                        │ int     │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 1  │ col                  │ int                        │ int     │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ return               │ char *                     │ *int8   │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceGetCellContent │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceGetCellContent │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.file         │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                            │         │
// └────┴──────────────────────┴────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceGetCellContent(row int, col int) *int8 {
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
	panic("not support return type: char *")
}

// GuiReferenceSearchGetCellContent    c api name: GuiReferenceSearchGetCellContent
// ┌────┬──────────────────────┬──────────────────────────────────┬─────────┐
// │ id │         name         │              c type              │ go type │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │ 0  │ row                  │ int                              │ int     │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │ 1  │ col                  │ int                              │ int     │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ return               │ char *                           │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h             │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSearchGetCellContent │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSearchGetCellContent │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.file         │                                  │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                  │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                  │         │
// └────┴──────────────────────┴──────────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceSearchGetCellContent(row int, col int) *int8 {
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
	panic("not support return type: char *")
}

// GuiReferenceReloadData    c api name: GuiReferenceReloadData
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceReloadData │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceReloadData │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceReloadData() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceReloadData").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiReferenceSetProgress    c api name: GuiReferenceSetProgress
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ progress             │ int                     │ int     │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetProgress │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetProgress │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
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
}

// GuiReferenceSetCurrentTaskProgress    c api name: GuiReferenceSetCurrentTaskProgress
// ┌────┬──────────────────────┬────────────────────────────────────┬─────────┐
// │ id │         name         │               c type               │ go type │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │ 0  │ progress             │ int                                │ int     │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │ 1  │ taskTitle            │ const char *                       │ *int8   │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ return               │ void                               │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h               │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetCurrentTaskProgress │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetCurrentTaskProgress │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                    │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                    │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                    │         │
// └────┴──────────────────────┴────────────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceSetCurrentTaskProgress(progress int, taskTitle *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSetCurrentTaskProgress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "progress",
				Type:  "int",
				Value: fmt.Sprintf("%v", progress),
			},
			{
				Name:  "taskTitle",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", taskTitle),
			},
		},
	))).Request()
}

// GuiReferenceSetSearchStartCol    c api name: GuiReferenceSetSearchStartCol
// ┌────┬──────────────────────┬───────────────────────────────┬─────────┐
// │ id │         name         │            c type             │ go type │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │ 0  │ col                  │ int                           │ int     │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ return               │ void                          │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h          │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetSearchStartCol │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetSearchStartCol │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.file         │                               │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                               │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                               │         │
// └────┴──────────────────────┴───────────────────────────────┴─────────┘
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
}

// GuiStackDumpAt    c api name: GuiStackDumpAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ csp                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiStackDumpAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiStackDumpAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiUpdateDumpView    c api name: GuiUpdateDumpView
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateDumpView    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateDumpView    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateDumpView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateDumpView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateWatchView    c api name: GuiUpdateWatchView
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateWatchView   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateWatchView   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateWatchView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateWatchView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateThreadView    c api name: GuiUpdateThreadView
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateThreadView  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateThreadView  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateThreadView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateThreadView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateMemoryView    c api name: GuiUpdateMemoryView
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateMemoryView  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateMemoryView  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateMemoryView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateMemoryView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiAddRecentFile    c api name: GuiAddRecentFile
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ file                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddRecentFile     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddRecentFile     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddRecentFile(file *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddRecentFile").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "file",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", file),
			},
		},
	))).Request()
}

// GuiSetLastException    c api name: GuiSetLastException
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ exception            │ unsigned int         │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSetLastException  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSetLastException  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiGetDisassembly    c api name: GuiGetDisassembly
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetDisassembly    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetDisassembly    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiGetDisassembly(addr uint, text *int8) bool {
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
	return true
}

// GuiMenuAdd    c api name: GuiMenuAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiMenuAdd(hMenu int, title *int8) int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
	return 0
}

// GuiMenuAddEntry    c api name: GuiMenuAddEntry
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuAddEntry      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuAddEntry      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiMenuAddEntry(hMenu int, title *int8) int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuAddEntry").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
	return 0
}

// GuiMenuAddSeparator    c api name: GuiMenuAddSeparator
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuAddSeparator  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuAddSeparator  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiMenuClear    c api name: GuiMenuClear
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuClear         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuClear         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiMenuRemove    c api name: GuiMenuRemove
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hEntryMenu           │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuRemove        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuRemove        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSelectionGet    c api name: GuiSelectionGet
// ┌────┬──────────────────────┬──────────────────────┬──────────────────┐
// │ id │         name         │        c type        │     go type      │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │ 0  │ hWindow              │ GUISELECTIONTYPE     │ GUISELECTIONTYPE │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │ 1  │ selection            │ SELECTIONDATA *      │ *SELECTIONDATA   │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ return               │ bool                 │ bool             │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.name         │ GuiSelectionGet      │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.mangledName  │ GuiSelectionGet      │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.file         │                      │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.includedFrom │                      │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.expansionLoc │                      │                  │
// └────┴──────────────────────┴──────────────────────┴──────────────────┘
func (b *bridgemain) GuiSelectionGet(hWindow GUISELECTIONTYPE, selection *SELECTIONDATA) bool {
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
	return true
}

// GuiSelectionSet    c api name: GuiSelectionSet
// ┌────┬──────────────────────┬───────────────────────┬──────────────────┐
// │ id │         name         │        c type         │     go type      │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │ 0  │ hWindow              │ GUISELECTIONTYPE      │ GUISELECTIONTYPE │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │ 1  │ selection            │ const SELECTIONDATA * │ *SELECTIONDATA   │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ return               │ bool                  │ bool             │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.name         │ GuiSelectionSet       │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.mangledName  │ GuiSelectionSet       │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.file         │                       │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.includedFrom │                       │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.expansionLoc │                       │                  │
// └────┴──────────────────────┴───────────────────────┴──────────────────┘
func (b *bridgemain) GuiSelectionSet(hWindow GUISELECTIONTYPE, selection *SELECTIONDATA) bool {
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
	return true
}

// GuiGetLineWindow    c api name: GuiGetLineWindow
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetLineWindow     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetLineWindow     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiGetLineWindow(title *int8, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetLineWindow").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	return true
}

// GuiAutoCompleteAddCmd    c api name: GuiAutoCompleteAddCmd
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ cmd                  │ const char *          │ *int8   │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiAutoCompleteAddCmd │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAutoCompleteAddCmd │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiAutoCompleteAddCmd(cmd *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAutoCompleteAddCmd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
}

// GuiAutoCompleteDelCmd    c api name: GuiAutoCompleteDelCmd
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ cmd                  │ const char *          │ *int8   │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiAutoCompleteDelCmd │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAutoCompleteDelCmd │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiAutoCompleteDelCmd(cmd *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAutoCompleteDelCmd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
}

// GuiAutoCompleteClearAll    c api name: GuiAutoCompleteClearAll
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiAutoCompleteClearAll │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAutoCompleteClearAll │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) GuiAutoCompleteClearAll() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAutoCompleteClearAll").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiAddStatusBarMessage    c api name: GuiAddStatusBarMessage
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ msg                  │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddStatusBarMessage │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddStatusBarMessage │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiAddStatusBarMessage(msg *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddStatusBarMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
}

// GuiUpdateSideBar    c api name: GuiUpdateSideBar
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateSideBar     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateSideBar     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateSideBar() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateSideBar").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiRepaintTableView    c api name: GuiRepaintTableView
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiRepaintTableView  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiRepaintTableView  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiRepaintTableView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiRepaintTableView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdatePatches    c api name: GuiUpdatePatches
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdatePatches     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdatePatches     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdatePatches() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdatePatches").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateCallStack    c api name: GuiUpdateCallStack
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateCallStack   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateCallStack   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateCallStack() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateCallStack").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiUpdateSEHChain    c api name: GuiUpdateSEHChain
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateSEHChain    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateSEHChain    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateSEHChain() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateSEHChain").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiLoadSourceFileEx    c api name: GuiLoadSourceFileEx
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ path                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiLoadSourceFileEx  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiLoadSourceFileEx  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiLoadSourceFileEx(path *int8, addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLoadSourceFileEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "path",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", path),
			},
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
}

// GuiMenuSetIcon    c api name: GuiMenuSetIcon
// ┌────┬──────────────────────┬──────────────────────┬───────────┐
// │ id │         name         │        c type        │  go type  │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 0  │ hMenu                │ int                  │ int       │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 1  │ icon                 │ const ICONDATA *     │ *ICONDATA │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ return               │ void                 │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.name         │ GuiMenuSetIcon       │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.mangledName  │ GuiMenuSetIcon       │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.file         │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.includedFrom │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.expansionLoc │                      │           │
// └────┴──────────────────────┴──────────────────────┴───────────┘
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
}

// GuiMenuSetEntryIcon    c api name: GuiMenuSetEntryIcon
// ┌────┬──────────────────────┬──────────────────────┬───────────┐
// │ id │         name         │        c type        │  go type  │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 0  │ hEntry               │ int                  │ int       │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 1  │ icon                 │ const ICONDATA *     │ *ICONDATA │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ return               │ void                 │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.name         │ GuiMenuSetEntryIcon  │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.mangledName  │ GuiMenuSetEntryIcon  │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.file         │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.includedFrom │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.expansionLoc │                      │           │
// └────┴──────────────────────┴──────────────────────┴───────────┘
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
}

// GuiMenuSetName    c api name: GuiMenuSetName
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ name                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuSetName       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuSetName       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiMenuSetName(hMenu int, name *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
}

// GuiMenuSetEntryName    c api name: GuiMenuSetEntryName
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hEntry               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ name                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuSetEntryName  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuSetEntryName  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiMenuSetEntryName(hEntry int, name *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetEntryName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntry",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntry),
			},
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
}

// GuiMenuSetEntryHotkey    c api name: GuiMenuSetEntryHotkey
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ hEntry               │ int                   │ int     │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 1  │ hack                 │ const char *          │ *int8   │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuSetEntryHotkey │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuSetEntryHotkey │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiMenuSetEntryHotkey(hEntry int, hack *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetEntryHotkey").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntry",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntry),
			},
			{
				Name:  "hack",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", hack),
			},
		},
	))).Request()
}

// GuiShowCpu    c api name: GuiShowCpu
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiShowCpu           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiShowCpu           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiShowCpu() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowCpu").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiShowThreads    c api name: GuiShowThreads
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiShowThreads       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiShowThreads       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiShowThreads() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowThreads").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiAddQWidgetTab    c api name: GuiAddQWidgetTab
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddQWidgetTab     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddQWidgetTab     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}

// GuiShowQWidgetTab    c api name: GuiShowQWidgetTab
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiShowQWidgetTab    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiShowQWidgetTab    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiShowQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}

// GuiCloseQWidgetTab    c api name: GuiCloseQWidgetTab
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiCloseQWidgetTab   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiCloseQWidgetTab   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiCloseQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiCloseQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}

// GuiExecuteOnGuiThread    c api name: GuiExecuteOnGuiThread
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ cbGuiThread          │ GUICALLBACK           │ uintptr │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiExecuteOnGuiThread │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiExecuteOnGuiThread │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
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
}

// GuiUpdateTimeWastedCounter    c api name: GuiUpdateTimeWastedCounter
// ┌────┬──────────────────────┬────────────────────────────┬─────────┐
// │ id │         name         │           c type           │ go type │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ return               │ void                       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateTimeWastedCounter │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateTimeWastedCounter │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.file         │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                            │         │
// └────┴──────────────────────┴────────────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateTimeWastedCounter() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateTimeWastedCounter").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiSetGlobalNotes    c api name: GuiSetGlobalNotes
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSetGlobalNotes    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSetGlobalNotes    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiSetGlobalNotes(text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetGlobalNotes").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
}

// GuiGetGlobalNotes    c api name: GuiGetGlobalNotes
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ text                 │ char **              │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetGlobalNotes    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetGlobalNotes    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSetDebuggeeNotes    c api name: GuiSetDebuggeeNotes
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSetDebuggeeNotes  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSetDebuggeeNotes  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiSetDebuggeeNotes(text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetDebuggeeNotes").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
}

// GuiGetDebuggeeNotes    c api name: GuiGetDebuggeeNotes
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ text                 │ char **              │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetDebuggeeNotes  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetDebuggeeNotes  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiDumpAtN    c api name: GuiDumpAtN
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ va                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ index                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDumpAtN           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDumpAtN           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiDisplayWarning    c api name: GuiDisplayWarning
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDisplayWarning    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDisplayWarning    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiDisplayWarning(title *int8, text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiDisplayWarning").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
}

// GuiRegisterScriptLanguage    c api name: GuiRegisterScriptLanguage
// ┌────┬──────────────────────┬───────────────────────────┬─────────────────┐
// │ id │         name         │          c type           │     go type     │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │ 0  │ info                 │ SCRIPTTYPEINFO *          │ *SCRIPTTYPEINFO │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ return               │ void                      │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.name         │ GuiRegisterScriptLanguage │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.mangledName  │ GuiRegisterScriptLanguage │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.file         │                           │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.includedFrom │                           │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.expansionLoc │                           │                 │
// └────┴──────────────────────┴───────────────────────────┴─────────────────┘
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
}

// GuiUnregisterScriptLanguage    c api name: GuiUnregisterScriptLanguage
// ┌────┬──────────────────────┬─────────────────────────────┬─────────┐
// │ id │         name         │           c type            │ go type │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │ 0  │ id                   │ int                         │ int     │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ return               │ void                        │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h        │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiUnregisterScriptLanguage │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUnregisterScriptLanguage │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.file         │                             │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                             │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                             │         │
// └────┴──────────────────────┴─────────────────────────────┴─────────┘
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
}

// GuiUpdateArgumentWidget    c api name: GuiUpdateArgumentWidget
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateArgumentWidget │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateArgumentWidget │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateArgumentWidget() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateArgumentWidget").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiFocusView    c api name: GuiFocusView
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hWindow              │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiFocusView         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiFocusView         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiIsUpdateDisabled    c api name: GuiIsUpdateDisabled
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiIsUpdateDisabled  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiIsUpdateDisabled  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiIsUpdateDisabled() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiIsUpdateDisabled").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// GuiUpdateDisable    c api name: GuiUpdateDisable
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateDisable     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateDisable     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateDisable() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateDisable").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiLoadGraph    c api name: GuiLoadGraph
// ┌────┬──────────────────────┬──────────────────────┬────────────────────┐
// │ id │         name         │        c type        │      go type       │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 0  │ graph                │ BridgeCFGraphList *  │ *BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 1  │ addr                 │ duint                │ uint               │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ return               │ bool                 │ bool               │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.name         │ GuiLoadGraph         │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.mangledName  │ GuiLoadGraph         │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.file         │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.includedFrom │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.expansionLoc │                      │                    │
// └────┴──────────────────────┴──────────────────────┴────────────────────┘
func (b *bridgemain) GuiLoadGraph(graph *BridgeCFGraphList, addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLoadGraph").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graph",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graph),
			},
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// GuiGraphAt    c api name: GuiGraphAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGraphAt           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGraphAt           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiGraphAt(addr uint) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGraphAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// GuiUpdateGraphView    c api name: GuiUpdateGraphView
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateGraphView   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateGraphView   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateGraphView() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateGraphView").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiDisableLog    c api name: GuiDisableLog
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDisableLog        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDisableLog        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiDisableLog() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiDisableLog").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiEnableLog    c api name: GuiEnableLog
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiEnableLog         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiEnableLog         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiEnableLog() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiEnableLog").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiIsLogEnabled    c api name: GuiIsLogEnabled
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiIsLogEnabled      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiIsLogEnabled      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiIsLogEnabled() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiIsLogEnabled").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// GuiAddFavouriteTool    c api name: GuiAddFavouriteTool
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ name                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ description          │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddFavouriteTool  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddFavouriteTool  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddFavouriteTool(name *int8, description *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddFavouriteTool").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "description",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", description),
			},
		},
	))).Request()
}

// GuiAddFavouriteCommand    c api name: GuiAddFavouriteCommand
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ name                 │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ shortcut             │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddFavouriteCommand │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddFavouriteCommand │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiAddFavouriteCommand(name *int8, shortcut *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddFavouriteCommand").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "shortcut",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", shortcut),
			},
		},
	))).Request()
}

// GuiSetFavouriteToolShortcut    c api name: GuiSetFavouriteToolShortcut
// ┌────┬──────────────────────┬─────────────────────────────┬─────────┐
// │ id │         name         │           c type            │ go type │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │ 0  │ name                 │ const char *                │ *int8   │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │ 1  │ shortcut             │ const char *                │ *int8   │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ return               │ void                        │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h        │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiSetFavouriteToolShortcut │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSetFavouriteToolShortcut │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.file         │                             │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                             │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                             │         │
// └────┴──────────────────────┴─────────────────────────────┴─────────┘
func (b *bridgemain) GuiSetFavouriteToolShortcut(name *int8, shortcut *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetFavouriteToolShortcut").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "shortcut",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", shortcut),
			},
		},
	))).Request()
}

// GuiFoldDisassembly    c api name: GuiFoldDisassembly
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ startAddress         │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ length               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiFoldDisassembly   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiFoldDisassembly   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSelectInMemoryMap    c api name: GuiSelectInMemoryMap
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSelectInMemoryMap │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSelectInMemoryMap │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiGetActiveView    c api name: GuiGetActiveView
// ┌────┬──────────────────────┬──────────────────────┬─────────────┐
// │ id │         name         │        c type        │   go type   │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │ 0  │ activeView           │ ACTIVEVIEW *         │ *ACTIVEVIEW │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ return               │ void                 │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.name         │ GuiGetActiveView     │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.mangledName  │ GuiGetActiveView     │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.file         │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.includedFrom │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.expansionLoc │                      │             │
// └────┴──────────────────────┴──────────────────────┴─────────────┘
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
}

// GuiAddInfoLine    c api name: GuiAddInfoLine
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ infoLine             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddInfoLine       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddInfoLine       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddInfoLine(infoLine *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddInfoLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "infoLine",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", infoLine),
			},
		},
	))).Request()
}

// GuiProcessEvents    c api name: GuiProcessEvents
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiProcessEvents     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiProcessEvents     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiProcessEvents() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiProcessEvents").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiTypeAddNode    c api name: GuiTypeAddNode
// ┌────┬──────────────────────┬────────────────────────┬──────────┐
// │ id │         name         │         c type         │ go type  │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │ 0  │ parent               │ void *                 │ uintptr  │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │ 1  │ Type                 │ const TYPEDESCRIPTOR * │ *uintptr │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ return               │ void *                 │ uintptr  │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.name         │ GuiTypeAddNode         │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.mangledName  │ GuiTypeAddNode         │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.file         │                        │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.includedFrom │                        │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.expansionLoc │                        │          │
// └────┴──────────────────────┴────────────────────────┴──────────┘
func (b *bridgemain) GuiTypeAddNode(parent uintptr, Type *uintptr) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiTypeAddNode").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "parent",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", parent),
			},
			{
				Name:  "Type",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", Type),
			},
		},
	))).Request()
	panic("not support return type: void *")
}

// GuiTypeClear    c api name: GuiTypeClear
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiTypeClear         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiTypeClear         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiTypeClear() bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiTypeClear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return true
}

// GuiUpdateTypeWidget    c api name: GuiUpdateTypeWidget
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateTypeWidget  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateTypeWidget  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateTypeWidget() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateTypeWidget").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiCloseApplication    c api name: GuiCloseApplication
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiCloseApplication  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiCloseApplication  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiCloseApplication() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiCloseApplication").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiFlushLog    c api name: GuiFlushLog
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiFlushLog          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiFlushLog          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiFlushLog() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiFlushLog").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiReferenceAddCommand    c api name: GuiReferenceAddCommand
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ title                │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ command              │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceAddCommand │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceAddCommand │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceAddCommand(title *int8, command *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceAddCommand").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "command",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", command),
			},
		},
	))).Request()
}

// GuiUpdateTraceBrowser    c api name: GuiUpdateTraceBrowser
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateTraceBrowser │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateTraceBrowser │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateTraceBrowser() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateTraceBrowser").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiOpenTraceFile    c api name: GuiOpenTraceFile
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ fileName             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiOpenTraceFile     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiOpenTraceFile     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiOpenTraceFile(fileName *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiOpenTraceFile").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "fileName",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", fileName),
			},
		},
	))).Request()
}

// GuiInvalidateSymbolSource    c api name: GuiInvalidateSymbolSource
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ base                 │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ void                      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ GuiInvalidateSymbolSource │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiInvalidateSymbolSource │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
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
}

// GuiExecuteOnGuiThreadEx    c api name: GuiExecuteOnGuiThreadEx
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ cbGuiThread          │ GUICALLBACKEX           │ uintptr │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 1  │ userdata             │ void *                  │ uintptr │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiExecuteOnGuiThreadEx │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiExecuteOnGuiThreadEx │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) GuiExecuteOnGuiThreadEx(cbGuiThread uintptr, userdata uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiExecuteOnGuiThreadEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cbGuiThread",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", cbGuiThread),
			},
			{
				Name:  "userdata",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", userdata),
			},
		},
	))).Request()
}

// GuiGetCurrentGraph    c api name: GuiGetCurrentGraph
// ┌────┬──────────────────────┬──────────────────────┬────────────────────┐
// │ id │         name         │        c type        │      go type       │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 0  │ graphList            │ BridgeCFGraphList *  │ *BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ return               │ void                 │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.name         │ GuiGetCurrentGraph   │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.mangledName  │ GuiGetCurrentGraph   │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.file         │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.includedFrom │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.expansionLoc │                      │                    │
// └────┴──────────────────────┴──────────────────────┴────────────────────┘
func (b *bridgemain) GuiGetCurrentGraph(graphList *BridgeCFGraphList) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetCurrentGraph").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graphList",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graphList),
			},
		},
	))).Request()
}

// GuiShowReferences    c api name: GuiShowReferences
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiShowReferences    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiShowReferences    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiShowReferences() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowReferences").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiSelectInSymbolsTab    c api name: GuiSelectInSymbolsTab
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiSelectInSymbolsTab │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSelectInSymbolsTab │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
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
}

// GuiGotoTrace    c api name: GuiGotoTrace
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ index                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGotoTrace         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGotoTrace         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiShowTrace    c api name: GuiShowTrace
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiShowTrace         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiShowTrace         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiShowTrace() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowTrace").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GuiGetMainThreadId    c api name: GuiGetMainThreadId
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ DWORD                │ uint32  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetMainThreadId   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetMainThreadId   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiGetMainThreadId() uint32 {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetMainThreadId").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: DWORD")
}

// BridgeAlloc    c api name: BridgeAlloc
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ size                 │ size_t               │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeAlloc          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeAlloc          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeAlloc(size uint) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeAlloc").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	panic("not support return type: void *")
}

// BridgeFree    c api name: BridgeFree
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ ptr                  │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeFree           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeFree           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeFree(ptr uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeFree").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ptr",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", ptr),
			},
		},
	))).Request()
}

// BridgeSettingGet    c api name: BridgeSettingGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ section              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ key                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ value                │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingGet     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingGet     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingGet(section *int8, key *int8, value *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// BridgeSettingGetUint    c api name: BridgeSettingGetUint
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ section              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ key                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ value                │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingGetUint │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingGetUint │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingGetUint(section *int8, key *int8, value *uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingGetUint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// BridgeSettingSet    c api name: BridgeSettingSet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ section              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ key                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ value                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingSet     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingSet     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingSet(section *int8, key *int8, value *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingSet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// BridgeSettingSetUint    c api name: BridgeSettingSetUint
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ section              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ key                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ value                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingSetUint │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingSetUint │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingSetUint(section *int8, key *int8, value uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingSetUint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "section",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", section),
			},
			{
				Name:  "key",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", key),
			},
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// BridgeSettingRead    c api name: BridgeSettingRead
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ errorLine            │ int *                │ *int    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ BridgeSettingRead    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeSettingRead    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) BridgeSettingRead(errorLine *int) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeSettingRead").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "errorLine",
				Type:  "*int ",
				Value: fmt.Sprintf("%v", errorLine),
			},
		},
	))).Request()
	return true
}

// DbgMemRead    c api name: DbgMemRead
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ va                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ dest                 │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ size                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemRead           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemRead           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemRead(va uint, dest uintptr, size uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemRead").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "va",
				Type:  "uint",
				Value: fmt.Sprintf("%v", va),
			},
			{
				Name:  "dest",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", dest),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	return true
}

// DbgMemWrite    c api name: DbgMemWrite
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ va                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ src                  │ const void *         │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ size                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemWrite          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemWrite          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemWrite(va uint, src uintptr, size uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemWrite").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "va",
				Type:  "uint",
				Value: fmt.Sprintf("%v", va),
			},
			{
				Name:  "src",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", src),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	return true
}

// DbgMemGetPageSize    c api name: DbgMemGetPageSize
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ base                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemGetPageSize    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemGetPageSize    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemGetPageSize(base uint) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemGetPageSize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "base",
				Type:  "uint",
				Value: fmt.Sprintf("%v", base),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgMemFindBaseAddr    c api name: DbgMemFindBaseAddr
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ size                 │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemFindBaseAddr   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemFindBaseAddr   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemFindBaseAddr(addr uint, size *uint) uint {
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
	panic("not support return type: duint")
}

// DbgCmdExec    c api name: DbgCmdExec
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ cmd                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgCmdExec           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgCmdExec           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgCmdExec(cmd *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgCmdExec").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
	return true
}

// DbgCmdExecDirect    c api name: DbgCmdExecDirect
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ cmd                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgCmdExecDirect     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgCmdExecDirect     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgCmdExecDirect(cmd *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgCmdExecDirect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
	return true
}

// DbgMemMap    c api name: DbgMemMap
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ memmap               │ MEMMAP *             │ *MEMMAP │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemMap            │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemMap            │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemMap(memmap *MEMMAP) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemMap").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "memmap",
				Type:  "*MEMMAP ",
				Value: fmt.Sprintf("%v", memmap),
			},
		},
	))).Request()
	return true
}

// DbgIsValidExpression    c api name: DbgIsValidExpression
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ expression           │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsValidExpression │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsValidExpression │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgIsValidExpression(expression *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsValidExpression").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "expression",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", expression),
			},
		},
	))).Request()
	return true
}

// DbgIsJumpGoingToExecute    c api name: DbgIsJumpGoingToExecute
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ bool                    │ bool    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsJumpGoingToExecute │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsJumpGoingToExecute │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) DbgIsJumpGoingToExecute(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsJumpGoingToExecute").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgGetLabelAt    c api name: DbgGetLabelAt
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ addr                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ segment              │ SEGMENTREG           │ SEGMENTREG │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 2  │ text                 │ char *               │ *int8      │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ bool                 │ bool       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgGetLabelAt        │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgGetLabelAt        │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgGetLabelAt(addr uint, segment SEGMENTREG, text *int8) bool {
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
	return true
}

// DbgSetLabelAt    c api name: DbgSetLabelAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetLabelAt        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetLabelAt        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetLabelAt(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetLabelAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
	return true
}

// DbgClearLabelRange    c api name: DbgClearLabelRange
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearLabelRange   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearLabelRange   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgGetCommentAt    c api name: DbgGetCommentAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetCommentAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetCommentAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetCommentAt(addr uint, text *int8) bool {
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
	return true
}

// DbgSetCommentAt    c api name: DbgSetCommentAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetCommentAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetCommentAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetCommentAt(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetCommentAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
	return true
}

// DbgClearCommentRange    c api name: DbgClearCommentRange
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearCommentRange │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearCommentRange │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgGetBookmarkAt    c api name: DbgGetBookmarkAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetBookmarkAt     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetBookmarkAt     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetBookmarkAt(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBookmarkAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgClearBookmarkRange    c api name: DbgClearBookmarkRange
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ start                │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 1  │ end                  │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearBookmarkRange │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearBookmarkRange │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
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
}

// DbgGetModuleAt    c api name: DbgGetModuleAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetModuleAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetModuleAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetModuleAt(addr uint, text *int8) bool {
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
	return true
}

// DbgGetBpxTypeAt    c api name: DbgGetBpxTypeAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ BPXTYPE              │ BPXTYPE │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetBpxTypeAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetBpxTypeAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetBpxTypeAt(addr uint) BPXTYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBpxTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: BPXTYPE")
}

// DbgValFromString    c api name: DbgValFromString
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ s                    │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgValFromString     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgValFromString     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgValFromString(s *int8) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgValFromString").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "s",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", s),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgGetRegDumpEx    c api name: DbgGetRegDumpEx
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ regdump              │ REGDUMP *            │ *REGDUMP │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ size                 │ size_t               │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ bool                 │ bool     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetRegDumpEx      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetRegDumpEx      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetRegDumpEx(regdump *REGDUMP, size uint) bool {
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
	return true
}

// DbgValToString    c api name: DbgValToString
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ s                    │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ value                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgValToString       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgValToString       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgValToString(s *int8, value uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgValToString").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "s",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", s),
			},
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// DbgMemIsValidReadPtr    c api name: DbgMemIsValidReadPtr
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMemIsValidReadPtr │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMemIsValidReadPtr │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgMemIsValidReadPtr(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgMemIsValidReadPtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgGetBpList    c api name: DbgGetBpList
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ Type                 │ BPXTYPE              │ BPXTYPE │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ list                 │ BPMAP *              │ *BPMAP  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetBpList         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetBpList         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetBpList(Type BPXTYPE, list *BPMAP) int {
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
	return 0
}

// DbgGetFunctionTypeAt    c api name: DbgGetFunctionTypeAt
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ FUNCTYPE             │ FUNCTYPE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetFunctionTypeAt │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetFunctionTypeAt │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetFunctionTypeAt(addr uint) FUNCTYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetFunctionTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: FUNCTYPE")
}

// DbgGetLoopTypeAt    c api name: DbgGetLoopTypeAt
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ depth                │ int                  │ int      │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ LOOPTYPE             │ LOOPTYPE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetLoopTypeAt     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetLoopTypeAt     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetLoopTypeAt(addr uint, depth int) LOOPTYPE {
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
	panic("not support return type: LOOPTYPE")
}

// DbgGetBranchDestination    c api name: DbgGetBranchDestination
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetBranchDestination │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetBranchDestination │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) DbgGetBranchDestination(addr uint) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetBranchDestination").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgScriptLoad    c api name: DbgScriptLoad
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ filename             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptLoad        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptLoad        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptLoad(filename *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptLoad").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
}

// DbgScriptRun    c api name: DbgScriptRun
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ destline             │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptRun         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptRun         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgScriptBpToggle    c api name: DbgScriptBpToggle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptBpToggle    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptBpToggle    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptBpToggle(line int) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptBpToggle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	return true
}

// DbgScriptBpGet    c api name: DbgScriptBpGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptBpGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptBpGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptBpGet(line int) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptBpGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	return true
}

// DbgScriptCmdExec    c api name: DbgScriptCmdExec
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ command              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptCmdExec     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptCmdExec     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgScriptCmdExec(command *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptCmdExec").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "command",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", command),
			},
		},
	))).Request()
	return true
}

// DbgScriptGetLineType    c api name: DbgScriptGetLineType
// ┌────┬──────────────────────┬──────────────────────┬────────────────┐
// │ id │         name         │        c type        │    go type     │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │ 0  │ line                 │ int                  │ int            │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ return               │ SCRIPTLINETYPE       │ SCRIPTLINETYPE │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.name         │ DbgScriptGetLineType │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.mangledName  │ DbgScriptGetLineType │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.file         │                      │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.includedFrom │                      │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.expansionLoc │                      │                │
// └────┴──────────────────────┴──────────────────────┴────────────────┘
func (b *bridgemain) DbgScriptGetLineType(line int) SCRIPTLINETYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgScriptGetLineType").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
		},
	))).Request()
	panic("not support return type: SCRIPTLINETYPE")
}

// DbgScriptSetIp    c api name: DbgScriptSetIp
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgScriptSetIp       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgScriptSetIp       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgScriptGetBranchInfo    c api name: DbgScriptGetBranchInfo
// ┌────┬──────────────────────┬────────────────────────┬───────────────┐
// │ id │         name         │         c type         │    go type    │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │ 0  │ line                 │ int                    │ int           │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │ 1  │ info                 │ SCRIPTBRANCH *         │ *SCRIPTBRANCH │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ return               │ bool                   │ bool          │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.name         │ DbgScriptGetBranchInfo │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.mangledName  │ DbgScriptGetBranchInfo │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.file         │                        │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.includedFrom │                        │               │
// ├────┼──────────────────────┼────────────────────────┼───────────────┤
// │    │ comment.expansionLoc │                        │               │
// └────┴──────────────────────┴────────────────────────┴───────────────┘
func (b *bridgemain) DbgScriptGetBranchInfo(line int, info *SCRIPTBRANCH) bool {
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
	return true
}

// DbgSymbolEnum    c api name: DbgSymbolEnum
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ base                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ cbSymbolEnum         │ CBSYMBOLENUM         │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ user                 │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSymbolEnum        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSymbolEnum        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSymbolEnum(base uint, cbSymbolEnum uintptr, user uintptr) bool {
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
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	return true
}

// DbgSymbolEnumFromCache    c api name: DbgSymbolEnumFromCache
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ base                 │ duint                  │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ cbSymbolEnum         │ CBSYMBOLENUM           │ uintptr │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 2  │ user                 │ void *                 │ uintptr │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ bool                   │ bool    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ DbgSymbolEnumFromCache │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSymbolEnumFromCache │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) DbgSymbolEnumFromCache(base uint, cbSymbolEnum uintptr, user uintptr) bool {
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
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	return true
}

// DbgSymbolEnumRange    c api name: DbgSymbolEnumRange
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ symbolMask           │ unsigned int         │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 3  │ cbSymbolEnum         │ CBSYMBOLENUM         │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 4  │ user                 │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSymbolEnumRange   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSymbolEnumRange   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSymbolEnumRange(start uint, end uint, symbolMask uint, cbSymbolEnum uintptr, user uintptr) bool {
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
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", user),
			},
		},
	))).Request()
	return true
}

// DbgAssembleAt    c api name: DbgAssembleAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ instruction          │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgAssembleAt        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgAssembleAt        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgAssembleAt(addr uint, instruction *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgAssembleAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "instruction",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", instruction),
			},
		},
	))).Request()
	return true
}

// DbgModBaseFromName    c api name: DbgModBaseFromName
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ name                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgModBaseFromName   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgModBaseFromName   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgModBaseFromName(name *int8) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgModBaseFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgDisasmAt    c api name: DbgDisasmAt
// ┌────┬──────────────────────┬──────────────────────┬───────────────┐
// │ id │         name         │        c type        │    go type    │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │ 0  │ addr                 │ duint                │ uint          │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │ 1  │ instr                │ DISASM_INSTR *       │ *DISASM_INSTR │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ return               │ void                 │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.name         │ DbgDisasmAt          │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.mangledName  │ DbgDisasmAt          │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.file         │                      │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.includedFrom │                      │               │
// ├────┼──────────────────────┼──────────────────────┼───────────────┤
// │    │ comment.expansionLoc │                      │               │
// └────┴──────────────────────┴──────────────────────┴───────────────┘
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
}

// DbgStackCommentGet    c api name: DbgStackCommentGet
// ┌────┬──────────────────────┬──────────────────────┬────────────────┐
// │ id │         name         │        c type        │    go type     │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │ 0  │ addr                 │ duint                │ uint           │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │ 1  │ comment              │ STACK_COMMENT *      │ *STACK_COMMENT │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ return               │ bool                 │ bool           │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.name         │ DbgStackCommentGet   │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.mangledName  │ DbgStackCommentGet   │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.file         │                      │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.includedFrom │                      │                │
// ├────┼──────────────────────┼──────────────────────┼────────────────┤
// │    │ comment.expansionLoc │                      │                │
// └────┴──────────────────────┴──────────────────────┴────────────────┘
func (b *bridgemain) DbgStackCommentGet(addr uint, comment *STACK_COMMENT) bool {
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
	return true
}

// DbgGetThreadList    c api name: DbgGetThreadList
// ┌────┬──────────────────────┬──────────────────────┬─────────────┐
// │ id │         name         │        c type        │   go type   │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │ 0  │ list                 │ THREADLIST *         │ *THREADLIST │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ return               │ void                 │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.name         │ DbgGetThreadList     │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.mangledName  │ DbgGetThreadList     │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.file         │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.includedFrom │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.expansionLoc │                      │             │
// └────┴──────────────────────┴──────────────────────┴─────────────┘
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
}

// DbgDisasmFastAt    c api name: DbgDisasmFastAt
// ┌────┬──────────────────────┬──────────────────────────┬──────────┐
// │ id │         name         │          c type          │ go type  │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                    │ uint     │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │ 1  │ basicinfo            │ BASIC_INSTRUCTION_INFO * │ *uintptr │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ return               │ void                     │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h     │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.name         │ DbgDisasmFastAt          │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgDisasmFastAt          │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.file         │                          │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.includedFrom │                          │          │
// ├────┼──────────────────────┼──────────────────────────┼──────────┤
// │    │ comment.expansionLoc │                          │          │
// └────┴──────────────────────┴──────────────────────────┴──────────┘
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
}

// DbgMenuEntryClicked    c api name: DbgMenuEntryClicked
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hEntry               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgMenuEntryClicked  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgMenuEntryClicked  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgFunctionGet    c api name: DbgFunctionGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ start                │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ end                  │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgFunctionGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgFunctionGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgFunctionGet(addr uint, start *uint, end *uint) bool {
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
	return true
}

// DbgFunctionOverlaps    c api name: DbgFunctionOverlaps
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgFunctionOverlaps  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgFunctionOverlaps  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgFunctionOverlaps(start uint, end uint) bool {
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
	return true
}

// DbgFunctionAdd    c api name: DbgFunctionAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgFunctionAdd       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgFunctionAdd       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgFunctionAdd(start uint, end uint) bool {
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
	return true
}

// DbgFunctionDel    c api name: DbgFunctionDel
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgFunctionDel       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgFunctionDel       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgFunctionDel(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgFunctionDel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgArgumentGet    c api name: DbgArgumentGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ start                │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ end                  │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgArgumentGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgArgumentGet       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgArgumentGet(addr uint, start *uint, end *uint) bool {
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
	return true
}

// DbgArgumentOverlaps    c api name: DbgArgumentOverlaps
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgArgumentOverlaps  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgArgumentOverlaps  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgArgumentOverlaps(start uint, end uint) bool {
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
	return true
}

// DbgArgumentAdd    c api name: DbgArgumentAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgArgumentAdd       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgArgumentAdd       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgArgumentAdd(start uint, end uint) bool {
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
	return true
}

// DbgArgumentDel    c api name: DbgArgumentDel
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgArgumentDel       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgArgumentDel       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgArgumentDel(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgArgumentDel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgLoopGet    c api name: DbgLoopGet
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ depth                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ start                │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 3  │ end                  │ duint *              │ *uint   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgLoopGet           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgLoopGet           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgLoopGet(depth int, addr uint, start *uint, end *uint) bool {
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
	return true
}

// DbgLoopOverlaps    c api name: DbgLoopOverlaps
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ depth                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 2  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgLoopOverlaps      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgLoopOverlaps      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgLoopOverlaps(depth int, start uint, end uint) bool {
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
	return true
}

// DbgLoopAdd    c api name: DbgLoopAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgLoopAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgLoopAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgLoopAdd(start uint, end uint) bool {
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
	return true
}

// DbgLoopDel    c api name: DbgLoopDel
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ depth                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgLoopDel           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgLoopDel           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgLoopDel(depth int, addr uint) bool {
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
	return true
}

// DbgXrefAdd    c api name: DbgXrefAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ from                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgXrefAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgXrefAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgXrefAdd(addr uint, from uint) bool {
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
	return true
}

// DbgXrefDelAll    c api name: DbgXrefDelAll
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgXrefDelAll        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgXrefDelAll        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgXrefDelAll(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgXrefDelAll").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgXrefGet    c api name: DbgXrefGet
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ addr                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ info                 │ XREF_INFO *          │ *XREF_INFO │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ bool                 │ bool       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgXrefGet           │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgXrefGet           │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgXrefGet(addr uint, info *XREF_INFO) bool {
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
	return true
}

// DbgGetXrefCountAt    c api name: DbgGetXrefCountAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ size_t               │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetXrefCountAt    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetXrefCountAt    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetXrefCountAt(addr uint) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetXrefCountAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: size_t")
}

// DbgGetXrefTypeAt    c api name: DbgGetXrefTypeAt
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ XREFTYPE             │ XREFTYPE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetXrefTypeAt     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetXrefTypeAt     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetXrefTypeAt(addr uint) XREFTYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetXrefTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: XREFTYPE")
}

// DbgIsBpDisabled    c api name: DbgIsBpDisabled
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgIsBpDisabled      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgIsBpDisabled      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgIsBpDisabled(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgIsBpDisabled").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgSetAutoCommentAt    c api name: DbgSetAutoCommentAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetAutoCommentAt  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetAutoCommentAt  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetAutoCommentAt(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoCommentAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
	return true
}

// DbgClearAutoCommentRange    c api name: DbgClearAutoCommentRange
// ┌────┬──────────────────────┬──────────────────────────┬─────────┐
// │ id │         name         │          c type          │ go type │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │ 0  │ start                │ duint                    │ uint    │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │ 1  │ end                  │ duint                    │ uint    │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ return               │ void                     │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h     │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearAutoCommentRange │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearAutoCommentRange │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.file         │                          │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.includedFrom │                          │         │
// ├────┼──────────────────────┼──────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                          │         │
// └────┴──────────────────────┴──────────────────────────┴─────────┘
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
}

// DbgSetAutoLabelAt    c api name: DbgSetAutoLabelAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetAutoLabelAt    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetAutoLabelAt    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetAutoLabelAt(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoLabelAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
	return true
}

// DbgClearAutoLabelRange    c api name: DbgClearAutoLabelRange
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ start                │ duint                  │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ end                  │ duint                  │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearAutoLabelRange │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearAutoLabelRange │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
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
}

// DbgSetAutoBookmarkAt    c api name: DbgSetAutoBookmarkAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetAutoBookmarkAt │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetAutoBookmarkAt │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetAutoBookmarkAt(addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgSetAutoBookmarkAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DbgClearAutoBookmarkRange    c api name: DbgClearAutoBookmarkRange
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ start                │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 1  │ end                  │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ void                      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearAutoBookmarkRange │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearAutoBookmarkRange │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
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
}

// DbgSetAutoFunctionAt    c api name: DbgSetAutoFunctionAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ start                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ end                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSetAutoFunctionAt │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSetAutoFunctionAt │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgSetAutoFunctionAt(start uint, end uint) bool {
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
	return true
}

// DbgClearAutoFunctionRange    c api name: DbgClearAutoFunctionRange
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ start                │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 1  │ end                  │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ void                      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ DbgClearAutoFunctionRange │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgClearAutoFunctionRange │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
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
}

// DbgGetStringAt    c api name: DbgGetStringAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetStringAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetStringAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetStringAt(addr uint, text *int8) bool {
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
	return true
}

// DbgWinEvent    c api name: DbgWinEvent
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ message              │ MSG *                │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ result               │ long *               │ *int32   │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ bool                 │ bool     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgWinEvent          │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgWinEvent          │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgWinEvent(message *uintptr, result *int32) bool {
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
	return true
}

// DbgWinEventGlobal    c api name: DbgWinEventGlobal
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ message              │ MSG *                │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ bool                 │ bool     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgWinEventGlobal    │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgWinEventGlobal    │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgWinEventGlobal(message *uintptr) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgWinEventGlobal").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	return true
}

// DbgGetArgTypeAt    c api name: DbgGetArgTypeAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ ARGTYPE              │ ARGTYPE │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetArgTypeAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetArgTypeAt      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetArgTypeAt(addr uint) ARGTYPE {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetArgTypeAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: ARGTYPE")
}

// DbgGetEncodeTypeBuffer    c api name: DbgGetEncodeTypeBuffer
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                  │ uint    │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ size                 │ duint *                │ *uint   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void *                 │ uintptr │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetEncodeTypeBuffer │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetEncodeTypeBuffer │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) DbgGetEncodeTypeBuffer(addr uint, size *uint) uintptr {
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
	panic("not support return type: void *")
}

// DbgReleaseEncodeTypeBuffer    c api name: DbgReleaseEncodeTypeBuffer
// ┌────┬──────────────────────┬────────────────────────────┬─────────┐
// │ id │         name         │           c type           │ go type │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 0  │ buffer               │ void *                     │ uintptr │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ return               │ void                       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.name         │ DbgReleaseEncodeTypeBuffer │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgReleaseEncodeTypeBuffer │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.file         │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                            │         │
// └────┴──────────────────────┴────────────────────────────┴─────────┘
func (b *bridgemain) DbgReleaseEncodeTypeBuffer(buffer uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgReleaseEncodeTypeBuffer").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "buffer",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", buffer),
			},
		},
	))).Request()
}

// DbgGetEncodeTypeAt    c api name: DbgGetEncodeTypeAt
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ addr                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ size                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ ENCODETYPE           │ ENCODETYPE │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgGetEncodeTypeAt   │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgGetEncodeTypeAt   │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgGetEncodeTypeAt(addr uint, size uint) ENCODETYPE {
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
	panic("not support return type: ENCODETYPE")
}

// DbgGetEncodeSizeAt    c api name: DbgGetEncodeSizeAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ codesize             │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetEncodeSizeAt   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetEncodeSizeAt   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetEncodeSizeAt(addr uint, codesize uint) uint {
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
	panic("not support return type: duint")
}

// DbgSetEncodeType    c api name: DbgSetEncodeType
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ addr                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ size                 │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 2  │ Type                 │ ENCODETYPE           │ ENCODETYPE │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ bool                 │ bool       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgSetEncodeType     │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgSetEncodeType     │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgSetEncodeType(addr uint, size uint, Type ENCODETYPE) bool {
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
	return true
}

// DbgDelEncodeTypeRange    c api name: DbgDelEncodeTypeRange
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ start                │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 1  │ end                  │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ DbgDelEncodeTypeRange │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgDelEncodeTypeRange │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
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
}

// DbgDelEncodeTypeSegment    c api name: DbgDelEncodeTypeSegment
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ start                │ duint                   │ uint    │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ DbgDelEncodeTypeSegment │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgDelEncodeTypeSegment │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
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
}

// DbgGetWatchList    c api name: DbgGetWatchList
// ┌────┬──────────────────────┬──────────────────────┬───────────┐
// │ id │         name         │        c type        │  go type  │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 0  │ list                 │ ListInfo *           │ *ListInfo │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ return               │ bool                 │ bool      │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.name         │ DbgGetWatchList      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.mangledName  │ DbgGetWatchList      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.file         │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.includedFrom │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.expansionLoc │                      │           │
// └────┴──────────────────────┴──────────────────────┴───────────┘
func (b *bridgemain) DbgGetWatchList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetWatchList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	return true
}

// DbgSelChanged    c api name: DbgSelChanged
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hWindow              │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ VA                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgSelChanged        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgSelChanged        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// DbgGetPebAddress    c api name: DbgGetPebAddress
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ ProcessId            │ DWORD                │ uint32  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetPebAddress     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetPebAddress     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetPebAddress(ProcessId uint32) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetPebAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ProcessId",
				Type:  "uint32",
				Value: fmt.Sprintf("%v", ProcessId),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgGetTebAddress    c api name: DbgGetTebAddress
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ ThreadId             │ DWORD                │ uint32  │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ DbgGetTebAddress     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ DbgGetTebAddress     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) DbgGetTebAddress(ThreadId uint32) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgGetTebAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ThreadId",
				Type:  "uint32",
				Value: fmt.Sprintf("%v", ThreadId),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// DbgAnalyzeFunction    c api name: DbgAnalyzeFunction
// ┌────┬──────────────────────┬──────────────────────┬────────────────────┐
// │ id │         name         │        c type        │      go type       │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 0  │ entry                │ duint                │ uint               │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 1  │ graph                │ BridgeCFGraphList *  │ *BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ return               │ bool                 │ bool               │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.name         │ DbgAnalyzeFunction   │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.mangledName  │ DbgAnalyzeFunction   │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.file         │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.includedFrom │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.expansionLoc │                      │                    │
// └────┴──────────────────────┴──────────────────────┴────────────────────┘
func (b *bridgemain) DbgAnalyzeFunction(entry uint, graph *BridgeCFGraphList) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/DbgAnalyzeFunction").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "entry",
				Type:  "uint",
				Value: fmt.Sprintf("%v", entry),
			},
			{
				Name:  "graph",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graph),
			},
		},
	))).Request()
	return true
}

// DbgGetSymbolInfo    c api name: DbgGetSymbolInfo
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ symbolptr            │ const SYMBOLPTR *    │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ info                 │ SYMBOLINFO *         │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ void                 │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetSymbolInfo     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetSymbolInfo     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
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
}

// DbgGetSymbolInfoAt    c api name: DbgGetSymbolInfoAt
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ addr                 │ duint                │ uint     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 1  │ info                 │ SYMBOLINFO *         │ *uintptr │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ bool                 │ bool     │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ DbgGetSymbolInfoAt   │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ DbgGetSymbolInfoAt   │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
func (b *bridgemain) DbgGetSymbolInfoAt(addr uint, info *uintptr) bool {
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
	return true
}

// DbgXrefAddMulti    c api name: DbgXrefAddMulti
// ┌────┬──────────────────────┬──────────────────────┬────────────┐
// │ id │         name         │        c type        │  go type   │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 0  │ edges                │ const XREF_EDGE *    │ *XREF_EDGE │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │ 1  │ count                │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ return               │ duint                │ uint       │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.name         │ DbgXrefAddMulti      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.mangledName  │ DbgXrefAddMulti      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.file         │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.includedFrom │                      │            │
// ├────┼──────────────────────┼──────────────────────┼────────────┤
// │    │ comment.expansionLoc │                      │            │
// └────┴──────────────────────┴──────────────────────┴────────────┘
func (b *bridgemain) DbgXrefAddMulti(edges *XREF_EDGE, count uint) uint {
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
	panic("not support return type: duint")
}

// DbgMenuPrepare    c api name: DbgMenuPrepare
// ┌────┬──────────────────────┬──────────────────────┬─────────────┐
// │ id │         name         │        c type        │   go type   │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │ 0  │ hMenu                │ GUIMENUTYPE          │ GUIMENUTYPE │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ return               │ void                 │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.name         │ DbgMenuPrepare       │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.mangledName  │ DbgMenuPrepare       │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.file         │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.includedFrom │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.expansionLoc │                      │             │
// └────┴──────────────────────┴──────────────────────┴─────────────┘
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
}

// GuiTranslateText    c api name: GuiTranslateText
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ Source               │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiTranslateText     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiTranslateText     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiTranslateText(Source *int8) *int8 {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiTranslateText").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "Source",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", Source),
			},
		},
	))).Request()
	panic("not support return type: const char *")
}

// GuiDisasmAt    c api name: GuiDisasmAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ cip                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDisasmAt          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDisasmAt          │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSetDebugState    c api name: GuiSetDebugState
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ state                │ DBGSTATE             │ DBGSTATE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ void                 │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ GuiSetDebugState     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ GuiSetDebugState     │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
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
}

// GuiSetDebugStateFast    c api name: GuiSetDebugStateFast
// ┌────┬──────────────────────┬──────────────────────┬──────────┐
// │ id │         name         │        c type        │ go type  │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │ 0  │ state                │ DBGSTATE             │ DBGSTATE │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ return               │ void                 │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.name         │ GuiSetDebugStateFast │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.mangledName  │ GuiSetDebugStateFast │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.file         │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.includedFrom │                      │          │
// ├────┼──────────────────────┼──────────────────────┼──────────┤
// │    │ comment.expansionLoc │                      │          │
// └────┴──────────────────────┴──────────────────────┴──────────┘
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
}

// GuiAddLogMessage    c api name: GuiAddLogMessage
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ msg                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddLogMessage     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddLogMessage     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddLogMessage(msg *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddLogMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
}

// GuiAddLogMessageHtml    c api name: GuiAddLogMessageHtml
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ msg                  │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddLogMessageHtml │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddLogMessageHtml │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddLogMessageHtml(msg *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddLogMessageHtml").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
}

// GuiLogSave    c api name: GuiLogSave
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ filename             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiLogSave           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiLogSave           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiLogSave(filename *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogSave").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
}

// GuiLogRedirect    c api name: GuiLogRedirect
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ filename             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiLogRedirect       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiLogRedirect       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiLogRedirect(filename *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLogRedirect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
}

// GuiUpdateWindowTitle    c api name: GuiUpdateWindowTitle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ filename             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiUpdateWindowTitle │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUpdateWindowTitle │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiUpdateWindowTitle(filename *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiUpdateWindowTitle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "filename",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", filename),
			},
		},
	))).Request()
}

// GuiDumpAt    c api name: GuiDumpAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ va                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDumpAt            │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDumpAt            │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiScriptAdd    c api name: GuiScriptAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ count                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ lines                │ const char **        │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptAdd         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptAdd         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptAdd(count int, lines *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "count",
				Type:  "int",
				Value: fmt.Sprintf("%v", count),
			},
			{
				Name:  "lines",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", lines),
			},
		},
	))).Request()
}

// GuiScriptSetIp    c api name: GuiScriptSetIp
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptSetIp       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptSetIp       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiScriptError    c api name: GuiScriptError
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ message              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptError       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptError       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptError(line int, message *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptError").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
}

// GuiScriptSetTitle    c api name: GuiScriptSetTitle
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptSetTitle    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptSetTitle    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptSetTitle(title *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptSetTitle").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
}

// GuiScriptSetInfoLine    c api name: GuiScriptSetInfoLine
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ line                 │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ info                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptSetInfoLine │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptSetInfoLine │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptSetInfoLine(line int, info *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptSetInfoLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "line",
				Type:  "int",
				Value: fmt.Sprintf("%v", line),
			},
			{
				Name:  "info",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
}

// GuiScriptMessage    c api name: GuiScriptMessage
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ message              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptMessage     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptMessage     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptMessage(message *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
}

// GuiScriptMsgyn    c api name: GuiScriptMsgyn
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ message              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiScriptMsgyn       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiScriptMsgyn       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiScriptMsgyn(message *int8) int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiScriptMsgyn").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	return 0
}

// GuiSymbolLogAdd    c api name: GuiSymbolLogAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ message              │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSymbolLogAdd      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSymbolLogAdd      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiSymbolLogAdd(message *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSymbolLogAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
}

// GuiSymbolSetProgress    c api name: GuiSymbolSetProgress
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ percent              │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSymbolSetProgress │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSymbolSetProgress │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSymbolUpdateModuleList    c api name: GuiSymbolUpdateModuleList
// ┌────┬──────────────────────┬───────────────────────────┬───────────────────┐
// │ id │         name         │          c type           │      go type      │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │ 0  │ count                │ int                       │ int               │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │ 1  │ modules              │ SYMBOLMODULEINFO *        │ *SYMBOLMODULEINFO │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ return               │ void                      │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.name         │ GuiSymbolUpdateModuleList │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.mangledName  │ GuiSymbolUpdateModuleList │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.file         │                           │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.includedFrom │                           │                   │
// ├────┼──────────────────────┼───────────────────────────┼───────────────────┤
// │    │ comment.expansionLoc │                           │                   │
// └────┴──────────────────────┴───────────────────────────┴───────────────────┘
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
}

// GuiReferenceAddColumn    c api name: GuiReferenceAddColumn
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ width                │ int                   │ int     │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 1  │ title                │ const char *          │ *int8   │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceAddColumn │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceAddColumn │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceAddColumn(width int, title *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceAddColumn").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "width",
				Type:  "int",
				Value: fmt.Sprintf("%v", width),
			},
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
}

// GuiReferenceSetRowCount    c api name: GuiReferenceSetRowCount
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ count                │ int                     │ int     │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetRowCount │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetRowCount │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
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
}

// GuiReferenceInitialize    c api name: GuiReferenceInitialize
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ name                 │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceInitialize │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceInitialize │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceInitialize(name *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceInitialize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
}

// GuiReferenceSetCellContent    c api name: GuiReferenceSetCellContent
// ┌────┬──────────────────────┬────────────────────────────┬─────────┐
// │ id │         name         │           c type           │ go type │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 0  │ row                  │ int                        │ int     │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 1  │ col                  │ int                        │ int     │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 2  │ str                  │ const char *               │ *int8   │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ return               │ void                       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetCellContent │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetCellContent │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.file         │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                            │         │
// └────┴──────────────────────┴────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceSetCellContent(row int, col int, str *int8) {
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
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", str),
			},
		},
	))).Request()
}

// GuiReferenceGetCellContent    c api name: GuiReferenceGetCellContent
// ┌────┬──────────────────────┬────────────────────────────┬─────────┐
// │ id │         name         │           c type           │ go type │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 0  │ row                  │ int                        │ int     │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │ 1  │ col                  │ int                        │ int     │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ return               │ char *                     │ *int8   │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h       │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceGetCellContent │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceGetCellContent │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.file         │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                            │         │
// ├────┼──────────────────────┼────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                            │         │
// └────┴──────────────────────┴────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceGetCellContent(row int, col int) *int8 {
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
	panic("not support return type: char *")
}

// GuiReferenceSearchGetCellContent    c api name: GuiReferenceSearchGetCellContent
// ┌────┬──────────────────────┬──────────────────────────────────┬─────────┐
// │ id │         name         │              c type              │ go type │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │ 0  │ row                  │ int                              │ int     │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │ 1  │ col                  │ int                              │ int     │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ return               │ char *                           │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h             │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSearchGetCellContent │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSearchGetCellContent │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.file         │                                  │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                  │         │
// ├────┼──────────────────────┼──────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                  │         │
// └────┴──────────────────────┴──────────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceSearchGetCellContent(row int, col int) *int8 {
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
	panic("not support return type: char *")
}

// GuiReferenceSetProgress    c api name: GuiReferenceSetProgress
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ progress             │ int                     │ int     │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetProgress │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetProgress │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
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
}

// GuiReferenceSetCurrentTaskProgress    c api name: GuiReferenceSetCurrentTaskProgress
// ┌────┬──────────────────────┬────────────────────────────────────┬─────────┐
// │ id │         name         │               c type               │ go type │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │ 0  │ progress             │ int                                │ int     │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │ 1  │ taskTitle            │ const char *                       │ *int8   │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ return               │ void                               │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h               │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetCurrentTaskProgress │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetCurrentTaskProgress │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                    │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                    │         │
// ├────┼──────────────────────┼────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                    │         │
// └────┴──────────────────────┴────────────────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceSetCurrentTaskProgress(progress int, taskTitle *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceSetCurrentTaskProgress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "progress",
				Type:  "int",
				Value: fmt.Sprintf("%v", progress),
			},
			{
				Name:  "taskTitle",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", taskTitle),
			},
		},
	))).Request()
}

// GuiReferenceSetSearchStartCol    c api name: GuiReferenceSetSearchStartCol
// ┌────┬──────────────────────┬───────────────────────────────┬─────────┐
// │ id │         name         │            c type             │ go type │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │ 0  │ col                  │ int                           │ int     │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ return               │ void                          │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h          │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceSetSearchStartCol │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceSetSearchStartCol │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.file         │                               │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                               │         │
// ├────┼──────────────────────┼───────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                               │         │
// └────┴──────────────────────┴───────────────────────────────┴─────────┘
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
}

// GuiStackDumpAt    c api name: GuiStackDumpAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ csp                  │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiStackDumpAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiStackDumpAt       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiAddRecentFile    c api name: GuiAddRecentFile
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ file                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddRecentFile     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddRecentFile     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddRecentFile(file *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddRecentFile").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "file",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", file),
			},
		},
	))).Request()
}

// GuiSetLastException    c api name: GuiSetLastException
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ exception            │ unsigned int         │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSetLastException  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSetLastException  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiGetDisassembly    c api name: GuiGetDisassembly
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetDisassembly    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetDisassembly    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiGetDisassembly(addr uint, text *int8) bool {
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
	return true
}

// GuiMenuAdd    c api name: GuiMenuAdd
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuAdd           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiMenuAdd(hMenu int, title *int8) int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuAdd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
	return 0
}

// GuiMenuAddEntry    c api name: GuiMenuAddEntry
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuAddEntry      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuAddEntry      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiMenuAddEntry(hMenu int, title *int8) int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuAddEntry").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
		},
	))).Request()
	return 0
}

// GuiMenuAddSeparator    c api name: GuiMenuAddSeparator
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuAddSeparator  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuAddSeparator  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiMenuClear    c api name: GuiMenuClear
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuClear         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuClear         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiMenuRemove    c api name: GuiMenuRemove
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hEntryMenu           │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuRemove        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuRemove        │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSelectionGet    c api name: GuiSelectionGet
// ┌────┬──────────────────────┬──────────────────────┬──────────────────┐
// │ id │         name         │        c type        │     go type      │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │ 0  │ hWindow              │ GUISELECTIONTYPE     │ GUISELECTIONTYPE │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │ 1  │ selection            │ SELECTIONDATA *      │ *SELECTIONDATA   │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ return               │ bool                 │ bool             │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.name         │ GuiSelectionGet      │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.mangledName  │ GuiSelectionGet      │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.file         │                      │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.includedFrom │                      │                  │
// ├────┼──────────────────────┼──────────────────────┼──────────────────┤
// │    │ comment.expansionLoc │                      │                  │
// └────┴──────────────────────┴──────────────────────┴──────────────────┘
func (b *bridgemain) GuiSelectionGet(hWindow GUISELECTIONTYPE, selection *SELECTIONDATA) bool {
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
	return true
}

// GuiSelectionSet    c api name: GuiSelectionSet
// ┌────┬──────────────────────┬───────────────────────┬──────────────────┐
// │ id │         name         │        c type         │     go type      │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │ 0  │ hWindow              │ GUISELECTIONTYPE      │ GUISELECTIONTYPE │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │ 1  │ selection            │ const SELECTIONDATA * │ *SELECTIONDATA   │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ return               │ bool                  │ bool             │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.name         │ GuiSelectionSet       │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.mangledName  │ GuiSelectionSet       │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.file         │                       │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.includedFrom │                       │                  │
// ├────┼──────────────────────┼───────────────────────┼──────────────────┤
// │    │ comment.expansionLoc │                       │                  │
// └────┴──────────────────────┴───────────────────────┴──────────────────┘
func (b *bridgemain) GuiSelectionSet(hWindow GUISELECTIONTYPE, selection *SELECTIONDATA) bool {
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
	return true
}

// GuiGetLineWindow    c api name: GuiGetLineWindow
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ char *               │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ bool                 │ bool    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetLineWindow     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetLineWindow     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiGetLineWindow(title *int8, text *int8) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetLineWindow").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	return true
}

// GuiAutoCompleteAddCmd    c api name: GuiAutoCompleteAddCmd
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ cmd                  │ const char *          │ *int8   │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiAutoCompleteAddCmd │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAutoCompleteAddCmd │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiAutoCompleteAddCmd(cmd *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAutoCompleteAddCmd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
}

// GuiAutoCompleteDelCmd    c api name: GuiAutoCompleteDelCmd
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ cmd                  │ const char *          │ *int8   │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiAutoCompleteDelCmd │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAutoCompleteDelCmd │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiAutoCompleteDelCmd(cmd *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAutoCompleteDelCmd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cmd",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", cmd),
			},
		},
	))).Request()
}

// GuiAddStatusBarMessage    c api name: GuiAddStatusBarMessage
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ msg                  │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddStatusBarMessage │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddStatusBarMessage │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiAddStatusBarMessage(msg *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddStatusBarMessage").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "msg",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", msg),
			},
		},
	))).Request()
}

// GuiLoadSourceFileEx    c api name: GuiLoadSourceFileEx
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ path                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiLoadSourceFileEx  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiLoadSourceFileEx  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiLoadSourceFileEx(path *int8, addr uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLoadSourceFileEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "path",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", path),
			},
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
}

// GuiMenuSetIcon    c api name: GuiMenuSetIcon
// ┌────┬──────────────────────┬──────────────────────┬───────────┐
// │ id │         name         │        c type        │  go type  │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 0  │ hMenu                │ int                  │ int       │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 1  │ icon                 │ const ICONDATA *     │ *ICONDATA │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ return               │ void                 │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.name         │ GuiMenuSetIcon       │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.mangledName  │ GuiMenuSetIcon       │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.file         │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.includedFrom │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.expansionLoc │                      │           │
// └────┴──────────────────────┴──────────────────────┴───────────┘
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
}

// GuiMenuSetEntryIcon    c api name: GuiMenuSetEntryIcon
// ┌────┬──────────────────────┬──────────────────────┬───────────┐
// │ id │         name         │        c type        │  go type  │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 0  │ hEntry               │ int                  │ int       │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │ 1  │ icon                 │ const ICONDATA *     │ *ICONDATA │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ return               │ void                 │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.name         │ GuiMenuSetEntryIcon  │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.mangledName  │ GuiMenuSetEntryIcon  │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.file         │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.includedFrom │                      │           │
// ├────┼──────────────────────┼──────────────────────┼───────────┤
// │    │ comment.expansionLoc │                      │           │
// └────┴──────────────────────┴──────────────────────┴───────────┘
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
}

// GuiMenuSetName    c api name: GuiMenuSetName
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hMenu                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ name                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuSetName       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuSetName       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiMenuSetName(hMenu int, name *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hMenu",
				Type:  "int",
				Value: fmt.Sprintf("%v", hMenu),
			},
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
}

// GuiMenuSetEntryName    c api name: GuiMenuSetEntryName
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hEntry               │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ name                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuSetEntryName  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuSetEntryName  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiMenuSetEntryName(hEntry int, name *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetEntryName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntry",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntry),
			},
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
}

// GuiMenuSetEntryHotkey    c api name: GuiMenuSetEntryHotkey
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ hEntry               │ int                   │ int     │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 1  │ hack                 │ const char *          │ *int8   │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiMenuSetEntryHotkey │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiMenuSetEntryHotkey │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
func (b *bridgemain) GuiMenuSetEntryHotkey(hEntry int, hack *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiMenuSetEntryHotkey").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "hEntry",
				Type:  "int",
				Value: fmt.Sprintf("%v", hEntry),
			},
			{
				Name:  "hack",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", hack),
			},
		},
	))).Request()
}

// GuiAddQWidgetTab    c api name: GuiAddQWidgetTab
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddQWidgetTab     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddQWidgetTab     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}

// GuiShowQWidgetTab    c api name: GuiShowQWidgetTab
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiShowQWidgetTab    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiShowQWidgetTab    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiShowQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiShowQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}

// GuiCloseQWidgetTab    c api name: GuiCloseQWidgetTab
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *               │ uintptr │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiCloseQWidgetTab   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiCloseQWidgetTab   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiCloseQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiCloseQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}

// GuiExecuteOnGuiThread    c api name: GuiExecuteOnGuiThread
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ cbGuiThread          │ GUICALLBACK           │ uintptr │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiExecuteOnGuiThread │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiExecuteOnGuiThread │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
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
}

// GuiSetGlobalNotes    c api name: GuiSetGlobalNotes
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSetGlobalNotes    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSetGlobalNotes    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiSetGlobalNotes(text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetGlobalNotes").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
}

// GuiGetGlobalNotes    c api name: GuiGetGlobalNotes
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ text                 │ char **              │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetGlobalNotes    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetGlobalNotes    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSetDebuggeeNotes    c api name: GuiSetDebuggeeNotes
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSetDebuggeeNotes  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSetDebuggeeNotes  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiSetDebuggeeNotes(text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetDebuggeeNotes").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
}

// GuiGetDebuggeeNotes    c api name: GuiGetDebuggeeNotes
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ text                 │ char **              │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGetDebuggeeNotes  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGetDebuggeeNotes  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiDumpAtN    c api name: GuiDumpAtN
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ va                   │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ index                │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDumpAtN           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDumpAtN           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiDisplayWarning    c api name: GuiDisplayWarning
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ title                │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ text                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiDisplayWarning    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiDisplayWarning    │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiDisplayWarning(title *int8, text *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiDisplayWarning").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
}

// GuiRegisterScriptLanguage    c api name: GuiRegisterScriptLanguage
// ┌────┬──────────────────────┬───────────────────────────┬─────────────────┐
// │ id │         name         │          c type           │     go type     │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │ 0  │ info                 │ SCRIPTTYPEINFO *          │ *SCRIPTTYPEINFO │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ return               │ void                      │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.name         │ GuiRegisterScriptLanguage │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.mangledName  │ GuiRegisterScriptLanguage │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.file         │                           │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.includedFrom │                           │                 │
// ├────┼──────────────────────┼───────────────────────────┼─────────────────┤
// │    │ comment.expansionLoc │                           │                 │
// └────┴──────────────────────┴───────────────────────────┴─────────────────┘
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
}

// GuiUnregisterScriptLanguage    c api name: GuiUnregisterScriptLanguage
// ┌────┬──────────────────────┬─────────────────────────────┬─────────┐
// │ id │         name         │           c type            │ go type │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │ 0  │ id                   │ int                         │ int     │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ return               │ void                        │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h        │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiUnregisterScriptLanguage │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiUnregisterScriptLanguage │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.file         │                             │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                             │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                             │         │
// └────┴──────────────────────┴─────────────────────────────┴─────────┘
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
}

// GuiFocusView    c api name: GuiFocusView
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ hWindow              │ int                  │ int     │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiFocusView         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiFocusView         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiLoadGraph    c api name: GuiLoadGraph
// ┌────┬──────────────────────┬──────────────────────┬────────────────────┐
// │ id │         name         │        c type        │      go type       │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 0  │ graph                │ BridgeCFGraphList *  │ *BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 1  │ addr                 │ duint                │ uint               │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ return               │ bool                 │ bool               │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.name         │ GuiLoadGraph         │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.mangledName  │ GuiLoadGraph         │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.file         │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.includedFrom │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.expansionLoc │                      │                    │
// └────┴──────────────────────┴──────────────────────┴────────────────────┘
func (b *bridgemain) GuiLoadGraph(graph *BridgeCFGraphList, addr uint) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiLoadGraph").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graph",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graph),
			},
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// GuiGraphAt    c api name: GuiGraphAt
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGraphAt           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGraphAt           │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiGraphAt(addr uint) uint {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGraphAt").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// GuiAddFavouriteTool    c api name: GuiAddFavouriteTool
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ name                 │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ description          │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddFavouriteTool  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddFavouriteTool  │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddFavouriteTool(name *int8, description *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddFavouriteTool").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "description",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", description),
			},
		},
	))).Request()
}

// GuiAddFavouriteCommand    c api name: GuiAddFavouriteCommand
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ name                 │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ shortcut             │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddFavouriteCommand │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddFavouriteCommand │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiAddFavouriteCommand(name *int8, shortcut *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddFavouriteCommand").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "shortcut",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", shortcut),
			},
		},
	))).Request()
}

// GuiSetFavouriteToolShortcut    c api name: GuiSetFavouriteToolShortcut
// ┌────┬──────────────────────┬─────────────────────────────┬─────────┐
// │ id │         name         │           c type            │ go type │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │ 0  │ name                 │ const char *                │ *int8   │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │ 1  │ shortcut             │ const char *                │ *int8   │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ return               │ void                        │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h        │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.name         │ GuiSetFavouriteToolShortcut │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSetFavouriteToolShortcut │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.file         │                             │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                             │         │
// ├────┼──────────────────────┼─────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                             │         │
// └────┴──────────────────────┴─────────────────────────────┴─────────┘
func (b *bridgemain) GuiSetFavouriteToolShortcut(name *int8, shortcut *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiSetFavouriteToolShortcut").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "shortcut",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", shortcut),
			},
		},
	))).Request()
}

// GuiFoldDisassembly    c api name: GuiFoldDisassembly
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ startAddress         │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 1  │ length               │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiFoldDisassembly   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiFoldDisassembly   │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiSelectInMemoryMap    c api name: GuiSelectInMemoryMap
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiSelectInMemoryMap │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSelectInMemoryMap │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}

// GuiGetActiveView    c api name: GuiGetActiveView
// ┌────┬──────────────────────┬──────────────────────┬─────────────┐
// │ id │         name         │        c type        │   go type   │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │ 0  │ activeView           │ ACTIVEVIEW *         │ *ACTIVEVIEW │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ return               │ void                 │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.name         │ GuiGetActiveView     │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.mangledName  │ GuiGetActiveView     │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.file         │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.includedFrom │                      │             │
// ├────┼──────────────────────┼──────────────────────┼─────────────┤
// │    │ comment.expansionLoc │                      │             │
// └────┴──────────────────────┴──────────────────────┴─────────────┘
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
}

// GuiAddInfoLine    c api name: GuiAddInfoLine
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ infoLine             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiAddInfoLine       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiAddInfoLine       │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiAddInfoLine(infoLine *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiAddInfoLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "infoLine",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", infoLine),
			},
		},
	))).Request()
}

// GuiTypeAddNode    c api name: GuiTypeAddNode
// ┌────┬──────────────────────┬────────────────────────┬──────────┐
// │ id │         name         │         c type         │ go type  │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │ 0  │ parent               │ void *                 │ uintptr  │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │ 1  │ Type                 │ const TYPEDESCRIPTOR * │ *uintptr │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ return               │ void *                 │ uintptr  │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.name         │ GuiTypeAddNode         │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.mangledName  │ GuiTypeAddNode         │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.file         │                        │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.includedFrom │                        │          │
// ├────┼──────────────────────┼────────────────────────┼──────────┤
// │    │ comment.expansionLoc │                        │          │
// └────┴──────────────────────┴────────────────────────┴──────────┘
func (b *bridgemain) GuiTypeAddNode(parent uintptr, Type *uintptr) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiTypeAddNode").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "parent",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", parent),
			},
			{
				Name:  "Type",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", Type),
			},
		},
	))).Request()
	panic("not support return type: void *")
}

// GuiReferenceAddCommand    c api name: GuiReferenceAddCommand
// ┌────┬──────────────────────┬────────────────────────┬─────────┐
// │ id │         name         │         c type         │ go type │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 0  │ title                │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │ 1  │ command              │ const char *           │ *int8   │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ return               │ void                   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h   │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.name         │ GuiReferenceAddCommand │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiReferenceAddCommand │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.file         │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.includedFrom │                        │         │
// ├────┼──────────────────────┼────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                        │         │
// └────┴──────────────────────┴────────────────────────┴─────────┘
func (b *bridgemain) GuiReferenceAddCommand(title *int8, command *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiReferenceAddCommand").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "command",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", command),
			},
		},
	))).Request()
}

// GuiOpenTraceFile    c api name: GuiOpenTraceFile
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ fileName             │ const char *         │ *int8   │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiOpenTraceFile     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiOpenTraceFile     │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
func (b *bridgemain) GuiOpenTraceFile(fileName *int8) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiOpenTraceFile").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "fileName",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", fileName),
			},
		},
	))).Request()
}

// GuiInvalidateSymbolSource    c api name: GuiInvalidateSymbolSource
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ base                 │ duint                     │ uint    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ void                      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ GuiInvalidateSymbolSource │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiInvalidateSymbolSource │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
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
}

// GuiExecuteOnGuiThreadEx    c api name: GuiExecuteOnGuiThreadEx
// ┌────┬──────────────────────┬─────────────────────────┬─────────┐
// │ id │         name         │         c type          │ go type │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 0  │ cbGuiThread          │ GUICALLBACKEX           │ uintptr │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │ 1  │ userdata             │ void *                  │ uintptr │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ return               │ void                    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h    │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.name         │ GuiExecuteOnGuiThreadEx │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiExecuteOnGuiThreadEx │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.file         │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.includedFrom │                         │         │
// ├────┼──────────────────────┼─────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                         │         │
// └────┴──────────────────────┴─────────────────────────┴─────────┘
func (b *bridgemain) GuiExecuteOnGuiThreadEx(cbGuiThread uintptr, userdata uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiExecuteOnGuiThreadEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "cbGuiThread",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", cbGuiThread),
			},
			{
				Name:  "userdata",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", userdata),
			},
		},
	))).Request()
}

// GuiGetCurrentGraph    c api name: GuiGetCurrentGraph
// ┌────┬──────────────────────┬──────────────────────┬────────────────────┐
// │ id │         name         │        c type        │      go type       │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │ 0  │ graphList            │ BridgeCFGraphList *  │ *BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ return               │ void                 │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.name         │ GuiGetCurrentGraph   │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.mangledName  │ GuiGetCurrentGraph   │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.file         │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.includedFrom │                      │                    │
// ├────┼──────────────────────┼──────────────────────┼────────────────────┤
// │    │ comment.expansionLoc │                      │                    │
// └────┴──────────────────────┴──────────────────────┴────────────────────┘
func (b *bridgemain) GuiGetCurrentGraph(graphList *BridgeCFGraphList) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/GuiGetCurrentGraph").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graphList",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graphList),
			},
		},
	))).Request()
}

// GuiSelectInSymbolsTab    c api name: GuiSelectInSymbolsTab
// ┌────┬──────────────────────┬───────────────────────┬─────────┐
// │ id │         name         │        c type         │ go type │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │ 0  │ addr                 │ duint                 │ uint    │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ return               │ void                  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h  │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.name         │ GuiSelectInSymbolsTab │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiSelectInSymbolsTab │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.file         │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.includedFrom │                       │         │
// ├────┼──────────────────────┼───────────────────────┼─────────┤
// │    │ comment.expansionLoc │                       │         │
// └────┴──────────────────────┴───────────────────────┴─────────┘
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
}

// GuiGotoTrace    c api name: GuiGotoTrace
// ┌────┬──────────────────────┬──────────────────────┬─────────┐
// │ id │         name         │        c type        │ go type │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │ 0  │ index                │ duint                │ uint    │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ return               │ void                 │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.name         │ GuiGotoTrace         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.mangledName  │ GuiGotoTrace         │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.file         │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.includedFrom │                      │         │
// ├────┼──────────────────────┼──────────────────────┼─────────┤
// │    │ comment.expansionLoc │                      │         │
// └────┴──────────────────────┴──────────────────────┴─────────┘
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
}
