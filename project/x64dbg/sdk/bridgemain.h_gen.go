package sdk

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
	response := safeGet("bridgemain.h/BridgeInit", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeStart() {
	response := safeGet("bridgemain.h/BridgeStart", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeAlloc(size uint) {
	response := safeGet("bridgemain.h/BridgeAlloc", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeFree(ptr *uintptr) {
	response := safeGet("bridgemain.h/BridgeFree", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeSettingGet(section string, key string, value *int8) {
	response := safeGet("bridgemain.h/BridgeSettingGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeSettingGetUint(section string, key string, value *uint) {
	response := safeGet("bridgemain.h/BridgeSettingGetUint", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeSettingSet(section string, key string, value string) {
	response := safeGet("bridgemain.h/BridgeSettingSet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeSettingSetUint(section string, key string, value uint) {
	response := safeGet("bridgemain.h/BridgeSettingSetUint", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeSettingFlush() {
	response := safeGet("bridgemain.h/BridgeSettingFlush", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeSettingRead(errorLine *int) {
	response := safeGet("bridgemain.h/BridgeSettingRead", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeGetDbgVersion() {
	response := safeGet("bridgemain.h/BridgeGetDbgVersion", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeIsProcessElevated() {
	response := safeGet("bridgemain.h/BridgeIsProcessElevated", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeGetNtBuildNumber() {
	response := safeGet("bridgemain.h/BridgeGetNtBuildNumber", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeUserDirectory() {
	response := safeGet("bridgemain.h/BridgeUserDirectory", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) BridgeIsARM64Emulated() {
	response := safeGet("bridgemain.h/BridgeIsARM64Emulated", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgInit() {
	response := safeGet("bridgemain.h/DbgInit", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgExit() {
	response := safeGet("bridgemain.h/DbgExit", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgMemRead(va uint, dest *uintptr, size uint) {
	response := safeGet("bridgemain.h/DbgMemRead", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgMemWrite(va uint, src *uintptr, size uint) {
	response := safeGet("bridgemain.h/DbgMemWrite", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgMemGetPageSize(base uint) {
	response := safeGet("bridgemain.h/DbgMemGetPageSize", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgMemFindBaseAddr(addr uint, size *uint) {
	response := safeGet("bridgemain.h/DbgMemFindBaseAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgCmdExec(cmd string) {
	response := safeGet("bridgemain.h/DbgCmdExec", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgCmdExecDirect(cmd string) {
	response := safeGet("bridgemain.h/DbgCmdExecDirect", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgMemMap(memmap *MEMMAP) {
	response := safeGet("bridgemain.h/DbgMemMap", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgIsValidExpression(expression string) {
	response := safeGet("bridgemain.h/DbgIsValidExpression", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgIsDebugging() {
	response := safeGet("bridgemain.h/DbgIsDebugging", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgIsJumpGoingToExecute(addr uint) {
	response := safeGet("bridgemain.h/DbgIsJumpGoingToExecute", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetLabelAt(addr uint, segment SEGMENTREG, text *int8) {
	response := safeGet("bridgemain.h/DbgGetLabelAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSetLabelAt(addr uint, text string) {
	response := safeGet("bridgemain.h/DbgSetLabelAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgClearLabelRange(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgClearLabelRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetCommentAt(addr uint, text *int8) {
	response := safeGet("bridgemain.h/DbgGetCommentAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSetCommentAt(addr uint, text string) {
	response := safeGet("bridgemain.h/DbgSetCommentAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgClearCommentRange(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgClearCommentRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetBookmarkAt(addr uint) {
	response := safeGet("bridgemain.h/DbgGetBookmarkAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgClearBookmarkRange(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgClearBookmarkRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetModuleAt(addr uint, text *int8) {
	response := safeGet("bridgemain.h/DbgGetModuleAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetBpxTypeAt(addr uint) {
	response := safeGet("bridgemain.h/DbgGetBpxTypeAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgValFromString(s string) {
	response := safeGet("bridgemain.h/DbgValFromString", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetRegDumpEx(regdump *REGDUMP, size uint) {
	response := safeGet("bridgemain.h/DbgGetRegDumpEx", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgValToString(s string, value uint) {
	response := safeGet("bridgemain.h/DbgValToString", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgMemIsValidReadPtr(addr uint) {
	response := safeGet("bridgemain.h/DbgMemIsValidReadPtr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetBpList(Type BPXTYPE, list *BPMAP) {
	response := safeGet("bridgemain.h/DbgGetBpList", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetFunctionTypeAt(addr uint) {
	response := safeGet("bridgemain.h/DbgGetFunctionTypeAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetLoopTypeAt(addr uint, depth int) {
	response := safeGet("bridgemain.h/DbgGetLoopTypeAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetBranchDestination(addr uint) {
	response := safeGet("bridgemain.h/DbgGetBranchDestination", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptLoad(filename string) {
	response := safeGet("bridgemain.h/DbgScriptLoad", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptUnload() {
	response := safeGet("bridgemain.h/DbgScriptUnload", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptRun(destline int) {
	response := safeGet("bridgemain.h/DbgScriptRun", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptStep() {
	response := safeGet("bridgemain.h/DbgScriptStep", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptBpToggle(line int) {
	response := safeGet("bridgemain.h/DbgScriptBpToggle", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptBpGet(line int) {
	response := safeGet("bridgemain.h/DbgScriptBpGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptCmdExec(command string) {
	response := safeGet("bridgemain.h/DbgScriptCmdExec", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptAbort() {
	response := safeGet("bridgemain.h/DbgScriptAbort", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptGetLineType(line int) {
	response := safeGet("bridgemain.h/DbgScriptGetLineType", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptSetIp(line int) {
	response := safeGet("bridgemain.h/DbgScriptSetIp", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgScriptGetBranchInfo(line int, info *SCRIPTBRANCH) {
	response := safeGet("bridgemain.h/DbgScriptGetBranchInfo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSymbolEnum(base uint, cbSymbolEnum uintptr, user *uintptr) {
	response := safeGet("bridgemain.h/DbgSymbolEnum", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSymbolEnumFromCache(base uint, cbSymbolEnum uintptr, user *uintptr) {
	response := safeGet("bridgemain.h/DbgSymbolEnumFromCache", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSymbolEnumRange(start uint, end uint, symbolMask uint, cbSymbolEnum uintptr, user *uintptr) {
	response := safeGet("bridgemain.h/DbgSymbolEnumRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgAssembleAt(addr uint, instruction string) {
	response := safeGet("bridgemain.h/DbgAssembleAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgModBaseFromName(name string) {
	response := safeGet("bridgemain.h/DbgModBaseFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgDisasmAt(addr uint, instr *DISASM_INSTR) {
	response := safeGet("bridgemain.h/DbgDisasmAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgStackCommentGet(addr uint, comment *STACK_COMMENT) {
	response := safeGet("bridgemain.h/DbgStackCommentGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetThreadList(list *THREADLIST) {
	response := safeGet("bridgemain.h/DbgGetThreadList", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSettingsUpdated() {
	response := safeGet("bridgemain.h/DbgSettingsUpdated", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgDisasmFastAt(addr uint, basicinfo *uintptr) {
	response := safeGet("bridgemain.h/DbgDisasmFastAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgMenuEntryClicked(hEntry int) {
	response := safeGet("bridgemain.h/DbgMenuEntryClicked", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgFunctionGet(addr uint, start *uint, end *uint) {
	response := safeGet("bridgemain.h/DbgFunctionGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgFunctionOverlaps(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgFunctionOverlaps", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgFunctionAdd(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgFunctionAdd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgFunctionDel(addr uint) {
	response := safeGet("bridgemain.h/DbgFunctionDel", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgArgumentGet(addr uint, start *uint, end *uint) {
	response := safeGet("bridgemain.h/DbgArgumentGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgArgumentOverlaps(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgArgumentOverlaps", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgArgumentAdd(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgArgumentAdd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgArgumentDel(addr uint) {
	response := safeGet("bridgemain.h/DbgArgumentDel", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgLoopGet(depth int, addr uint, start *uint, end *uint) {
	response := safeGet("bridgemain.h/DbgLoopGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgLoopOverlaps(depth int, start uint, end uint) {
	response := safeGet("bridgemain.h/DbgLoopOverlaps", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgLoopAdd(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgLoopAdd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgLoopDel(depth int, addr uint) {
	response := safeGet("bridgemain.h/DbgLoopDel", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgXrefAdd(addr uint, from uint) {
	response := safeGet("bridgemain.h/DbgXrefAdd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgXrefDelAll(addr uint) {
	response := safeGet("bridgemain.h/DbgXrefDelAll", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgXrefGet(addr uint, info *XREF_INFO) {
	response := safeGet("bridgemain.h/DbgXrefGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetXrefCountAt(addr uint) {
	response := safeGet("bridgemain.h/DbgGetXrefCountAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetXrefTypeAt(addr uint) {
	response := safeGet("bridgemain.h/DbgGetXrefTypeAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgIsRunLocked() {
	response := safeGet("bridgemain.h/DbgIsRunLocked", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgIsBpDisabled(addr uint) {
	response := safeGet("bridgemain.h/DbgIsBpDisabled", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSetAutoCommentAt(addr uint, text string) {
	response := safeGet("bridgemain.h/DbgSetAutoCommentAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgClearAutoCommentRange(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgClearAutoCommentRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSetAutoLabelAt(addr uint, text string) {
	response := safeGet("bridgemain.h/DbgSetAutoLabelAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgClearAutoLabelRange(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgClearAutoLabelRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSetAutoBookmarkAt(addr uint) {
	response := safeGet("bridgemain.h/DbgSetAutoBookmarkAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgClearAutoBookmarkRange(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgClearAutoBookmarkRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSetAutoFunctionAt(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgSetAutoFunctionAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgClearAutoFunctionRange(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgClearAutoFunctionRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetStringAt(addr uint, text *int8) {
	response := safeGet("bridgemain.h/DbgGetStringAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgFunctions() {
	response := safeGet("bridgemain.h/DbgFunctions", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgWinEvent(message *uintptr, result *int32) {
	response := safeGet("bridgemain.h/DbgWinEvent", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgWinEventGlobal(message *uintptr) {
	response := safeGet("bridgemain.h/DbgWinEventGlobal", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgIsRunning() {
	response := safeGet("bridgemain.h/DbgIsRunning", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetTimeWastedCounter() {
	response := safeGet("bridgemain.h/DbgGetTimeWastedCounter", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetArgTypeAt(addr uint) {
	response := safeGet("bridgemain.h/DbgGetArgTypeAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetEncodeTypeBuffer(addr uint, size *uint) {
	response := safeGet("bridgemain.h/DbgGetEncodeTypeBuffer", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgReleaseEncodeTypeBuffer(buffer *uintptr) {
	response := safeGet("bridgemain.h/DbgReleaseEncodeTypeBuffer", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetEncodeTypeAt(addr uint, size uint) {
	response := safeGet("bridgemain.h/DbgGetEncodeTypeAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetEncodeSizeAt(addr uint, codesize uint) {
	response := safeGet("bridgemain.h/DbgGetEncodeSizeAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSetEncodeType(addr uint, size uint, Type ENCODETYPE) {
	response := safeGet("bridgemain.h/DbgSetEncodeType", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgDelEncodeTypeRange(start uint, end uint) {
	response := safeGet("bridgemain.h/DbgDelEncodeTypeRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgDelEncodeTypeSegment(start uint) {
	response := safeGet("bridgemain.h/DbgDelEncodeTypeSegment", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetWatchList(list *ListInfo) {
	response := safeGet("bridgemain.h/DbgGetWatchList", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgSelChanged(hWindow int, VA uint) {
	response := safeGet("bridgemain.h/DbgSelChanged", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetProcessHandle() {
	response := safeGet("bridgemain.h/DbgGetProcessHandle", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetThreadHandle() {
	response := safeGet("bridgemain.h/DbgGetThreadHandle", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetProcessId() {
	response := safeGet("bridgemain.h/DbgGetProcessId", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetThreadId() {
	response := safeGet("bridgemain.h/DbgGetThreadId", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetPebAddress(ProcessId uint32) {
	response := safeGet("bridgemain.h/DbgGetPebAddress", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetTebAddress(ThreadId uint32) {
	response := safeGet("bridgemain.h/DbgGetTebAddress", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgAnalyzeFunction(entry uint, graph *uintptr) {
	response := safeGet("bridgemain.h/DbgAnalyzeFunction", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetSymbolInfo(symbolptr *uintptr, info *uintptr) {
	response := safeGet("bridgemain.h/DbgGetSymbolInfo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetDebugEngine() {
	response := safeGet("bridgemain.h/DbgGetDebugEngine", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgGetSymbolInfoAt(addr uint, info *uintptr) {
	response := safeGet("bridgemain.h/DbgGetSymbolInfoAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgXrefAddMulti(edges *XREF_EDGE, count uint) {
	response := safeGet("bridgemain.h/DbgXrefAddMulti", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) DbgMenuPrepare(hMenu GUIMENUTYPE) {
	response := safeGet("bridgemain.h/DbgMenuPrepare", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiTranslateText(Source string) {
	response := safeGet("bridgemain.h/GuiTranslateText", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiDisasmAt(addr uint, cip uint) {
	response := safeGet("bridgemain.h/GuiDisasmAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSetDebugState(state DBGSTATE) {
	response := safeGet("bridgemain.h/GuiSetDebugState", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSetDebugStateFast(state DBGSTATE) {
	response := safeGet("bridgemain.h/GuiSetDebugStateFast", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAddLogMessage(msg string) {
	response := safeGet("bridgemain.h/GuiAddLogMessage", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAddLogMessageHtml(msg string) {
	response := safeGet("bridgemain.h/GuiAddLogMessageHtml", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiLogClear() {
	response := safeGet("bridgemain.h/GuiLogClear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiLogSave(filename string) {
	response := safeGet("bridgemain.h/GuiLogSave", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiLogRedirect(filename string) {
	response := safeGet("bridgemain.h/GuiLogRedirect", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiLogRedirectStop() {
	response := safeGet("bridgemain.h/GuiLogRedirectStop", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateAllViews() {
	response := safeGet("bridgemain.h/GuiUpdateAllViews", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateRegisterView() {
	response := safeGet("bridgemain.h/GuiUpdateRegisterView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateDisassemblyView() {
	response := safeGet("bridgemain.h/GuiUpdateDisassemblyView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateBreakpointsView() {
	response := safeGet("bridgemain.h/GuiUpdateBreakpointsView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateWindowTitle(filename string) {
	response := safeGet("bridgemain.h/GuiUpdateWindowTitle", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGetWindowHandle() {
	response := safeGet("bridgemain.h/GuiGetWindowHandle", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiDumpAt(va uint) {
	response := safeGet("bridgemain.h/GuiDumpAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiScriptAdd(count int, lines *string) {
	response := safeGet("bridgemain.h/GuiScriptAdd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiScriptClear() {
	response := safeGet("bridgemain.h/GuiScriptClear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiScriptSetIp(line int) {
	response := safeGet("bridgemain.h/GuiScriptSetIp", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiScriptError(line int, message string) {
	response := safeGet("bridgemain.h/GuiScriptError", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiScriptSetTitle(title string) {
	response := safeGet("bridgemain.h/GuiScriptSetTitle", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiScriptSetInfoLine(line int, info string) {
	response := safeGet("bridgemain.h/GuiScriptSetInfoLine", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiScriptMessage(message string) {
	response := safeGet("bridgemain.h/GuiScriptMessage", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiScriptMsgyn(message string) {
	response := safeGet("bridgemain.h/GuiScriptMsgyn", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSymbolLogAdd(message string) {
	response := safeGet("bridgemain.h/GuiSymbolLogAdd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSymbolLogClear() {
	response := safeGet("bridgemain.h/GuiSymbolLogClear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSymbolSetProgress(percent int) {
	response := safeGet("bridgemain.h/GuiSymbolSetProgress", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSymbolUpdateModuleList(count int, modules *SYMBOLMODULEINFO) {
	response := safeGet("bridgemain.h/GuiSymbolUpdateModuleList", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSymbolRefreshCurrent() {
	response := safeGet("bridgemain.h/GuiSymbolRefreshCurrent", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceAddColumn(width int, title string) {
	response := safeGet("bridgemain.h/GuiReferenceAddColumn", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceSetRowCount(count int) {
	response := safeGet("bridgemain.h/GuiReferenceSetRowCount", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceGetRowCount() {
	response := safeGet("bridgemain.h/GuiReferenceGetRowCount", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceSearchGetRowCount() {
	response := safeGet("bridgemain.h/GuiReferenceSearchGetRowCount", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceDeleteAllColumns() {
	response := safeGet("bridgemain.h/GuiReferenceDeleteAllColumns", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceInitialize(name string) {
	response := safeGet("bridgemain.h/GuiReferenceInitialize", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceSetCellContent(row int, col int, str string) {
	response := safeGet("bridgemain.h/GuiReferenceSetCellContent", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceGetCellContent(row int, col int) {
	response := safeGet("bridgemain.h/GuiReferenceGetCellContent", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceSearchGetCellContent(row int, col int) {
	response := safeGet("bridgemain.h/GuiReferenceSearchGetCellContent", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceReloadData() {
	response := safeGet("bridgemain.h/GuiReferenceReloadData", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceSetProgress(progress int) {
	response := safeGet("bridgemain.h/GuiReferenceSetProgress", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceSetCurrentTaskProgress(progress int, taskTitle string) {
	response := safeGet("bridgemain.h/GuiReferenceSetCurrentTaskProgress", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceSetSearchStartCol(col int) {
	response := safeGet("bridgemain.h/GuiReferenceSetSearchStartCol", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiStackDumpAt(addr uint, csp uint) {
	response := safeGet("bridgemain.h/GuiStackDumpAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateDumpView() {
	response := safeGet("bridgemain.h/GuiUpdateDumpView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateWatchView() {
	response := safeGet("bridgemain.h/GuiUpdateWatchView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateThreadView() {
	response := safeGet("bridgemain.h/GuiUpdateThreadView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateMemoryView() {
	response := safeGet("bridgemain.h/GuiUpdateMemoryView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAddRecentFile(file string) {
	response := safeGet("bridgemain.h/GuiAddRecentFile", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSetLastException(exception uint) {
	response := safeGet("bridgemain.h/GuiSetLastException", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGetDisassembly(addr uint, text *int8) {
	response := safeGet("bridgemain.h/GuiGetDisassembly", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuAdd(hMenu int, title string) {
	response := safeGet("bridgemain.h/GuiMenuAdd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuAddEntry(hMenu int, title string) {
	response := safeGet("bridgemain.h/GuiMenuAddEntry", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuAddSeparator(hMenu int) {
	response := safeGet("bridgemain.h/GuiMenuAddSeparator", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuClear(hMenu int) {
	response := safeGet("bridgemain.h/GuiMenuClear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuRemove(hEntryMenu int) {
	response := safeGet("bridgemain.h/GuiMenuRemove", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSelectionGet(hWindow GUISELECTIONTYPE, selection *SELECTIONDATA) {
	response := safeGet("bridgemain.h/GuiSelectionGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSelectionSet(hWindow GUISELECTIONTYPE, selection *SELECTIONDATA) {
	response := safeGet("bridgemain.h/GuiSelectionSet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGetLineWindow(title string, text *int8) {
	response := safeGet("bridgemain.h/GuiGetLineWindow", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAutoCompleteAddCmd(cmd string) {
	response := safeGet("bridgemain.h/GuiAutoCompleteAddCmd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAutoCompleteDelCmd(cmd string) {
	response := safeGet("bridgemain.h/GuiAutoCompleteDelCmd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAutoCompleteClearAll() {
	response := safeGet("bridgemain.h/GuiAutoCompleteClearAll", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAddStatusBarMessage(msg string) {
	response := safeGet("bridgemain.h/GuiAddStatusBarMessage", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateSideBar() {
	response := safeGet("bridgemain.h/GuiUpdateSideBar", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiRepaintTableView() {
	response := safeGet("bridgemain.h/GuiRepaintTableView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdatePatches() {
	response := safeGet("bridgemain.h/GuiUpdatePatches", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateCallStack() {
	response := safeGet("bridgemain.h/GuiUpdateCallStack", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateSEHChain() {
	response := safeGet("bridgemain.h/GuiUpdateSEHChain", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiLoadSourceFileEx(path string, addr uint) {
	response := safeGet("bridgemain.h/GuiLoadSourceFileEx", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuSetIcon(hMenu int, icon *ICONDATA) {
	response := safeGet("bridgemain.h/GuiMenuSetIcon", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuSetEntryIcon(hEntry int, icon *ICONDATA) {
	response := safeGet("bridgemain.h/GuiMenuSetEntryIcon", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuSetName(hMenu int, name string) {
	response := safeGet("bridgemain.h/GuiMenuSetName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuSetEntryName(hEntry int, name string) {
	response := safeGet("bridgemain.h/GuiMenuSetEntryName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiMenuSetEntryHotkey(hEntry int, hack string) {
	response := safeGet("bridgemain.h/GuiMenuSetEntryHotkey", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiShowCpu() {
	response := safeGet("bridgemain.h/GuiShowCpu", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiShowThreads() {
	response := safeGet("bridgemain.h/GuiShowThreads", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAddQWidgetTab(qWidget *uintptr) {
	response := safeGet("bridgemain.h/GuiAddQWidgetTab", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiShowQWidgetTab(qWidget *uintptr) {
	response := safeGet("bridgemain.h/GuiShowQWidgetTab", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiCloseQWidgetTab(qWidget *uintptr) {
	response := safeGet("bridgemain.h/GuiCloseQWidgetTab", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiExecuteOnGuiThread(cbGuiThread uintptr) {
	response := safeGet("bridgemain.h/GuiExecuteOnGuiThread", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateTimeWastedCounter() {
	response := safeGet("bridgemain.h/GuiUpdateTimeWastedCounter", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSetGlobalNotes(text string) {
	response := safeGet("bridgemain.h/GuiSetGlobalNotes", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGetGlobalNotes(text *int8) {
	response := safeGet("bridgemain.h/GuiGetGlobalNotes", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSetDebuggeeNotes(text string) {
	response := safeGet("bridgemain.h/GuiSetDebuggeeNotes", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGetDebuggeeNotes(text *int8) {
	response := safeGet("bridgemain.h/GuiGetDebuggeeNotes", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiDumpAtN(va uint, index int) {
	response := safeGet("bridgemain.h/GuiDumpAtN", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiDisplayWarning(title string, text string) {
	response := safeGet("bridgemain.h/GuiDisplayWarning", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiRegisterScriptLanguage(info *SCRIPTTYPEINFO) {
	response := safeGet("bridgemain.h/GuiRegisterScriptLanguage", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUnregisterScriptLanguage(id int) {
	response := safeGet("bridgemain.h/GuiUnregisterScriptLanguage", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateArgumentWidget() {
	response := safeGet("bridgemain.h/GuiUpdateArgumentWidget", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiFocusView(hWindow int) {
	response := safeGet("bridgemain.h/GuiFocusView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiIsUpdateDisabled() {
	response := safeGet("bridgemain.h/GuiIsUpdateDisabled", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateDisable() {
	response := safeGet("bridgemain.h/GuiUpdateDisable", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiLoadGraph(graph *uintptr, addr uint) {
	response := safeGet("bridgemain.h/GuiLoadGraph", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGraphAt(addr uint) {
	response := safeGet("bridgemain.h/GuiGraphAt", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateGraphView() {
	response := safeGet("bridgemain.h/GuiUpdateGraphView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiDisableLog() {
	response := safeGet("bridgemain.h/GuiDisableLog", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiEnableLog() {
	response := safeGet("bridgemain.h/GuiEnableLog", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiIsLogEnabled() {
	response := safeGet("bridgemain.h/GuiIsLogEnabled", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAddFavouriteTool(name string, description string) {
	response := safeGet("bridgemain.h/GuiAddFavouriteTool", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAddFavouriteCommand(name string, shortcut string) {
	response := safeGet("bridgemain.h/GuiAddFavouriteCommand", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSetFavouriteToolShortcut(name string, shortcut string) {
	response := safeGet("bridgemain.h/GuiSetFavouriteToolShortcut", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiFoldDisassembly(startAddress uint, length uint) {
	response := safeGet("bridgemain.h/GuiFoldDisassembly", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSelectInMemoryMap(addr uint) {
	response := safeGet("bridgemain.h/GuiSelectInMemoryMap", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGetActiveView(activeView *ACTIVEVIEW) {
	response := safeGet("bridgemain.h/GuiGetActiveView", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiAddInfoLine(infoLine string) {
	response := safeGet("bridgemain.h/GuiAddInfoLine", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiProcessEvents() {
	response := safeGet("bridgemain.h/GuiProcessEvents", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiTypeAddNode(parent *uintptr, Type *uintptr) {
	response := safeGet("bridgemain.h/GuiTypeAddNode", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiTypeClear() {
	response := safeGet("bridgemain.h/GuiTypeClear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateTypeWidget() {
	response := safeGet("bridgemain.h/GuiUpdateTypeWidget", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiCloseApplication() {
	response := safeGet("bridgemain.h/GuiCloseApplication", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiFlushLog() {
	response := safeGet("bridgemain.h/GuiFlushLog", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiReferenceAddCommand(title string, command string) {
	response := safeGet("bridgemain.h/GuiReferenceAddCommand", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiUpdateTraceBrowser() {
	response := safeGet("bridgemain.h/GuiUpdateTraceBrowser", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiOpenTraceFile(fileName string) {
	response := safeGet("bridgemain.h/GuiOpenTraceFile", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiInvalidateSymbolSource(base uint) {
	response := safeGet("bridgemain.h/GuiInvalidateSymbolSource", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiExecuteOnGuiThreadEx(cbGuiThread uintptr, userdata *uintptr) {
	response := safeGet("bridgemain.h/GuiExecuteOnGuiThreadEx", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGetCurrentGraph(graphList *uintptr) {
	response := safeGet("bridgemain.h/GuiGetCurrentGraph", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiShowReferences() {
	response := safeGet("bridgemain.h/GuiShowReferences", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiSelectInSymbolsTab(addr uint) {
	response := safeGet("bridgemain.h/GuiSelectInSymbolsTab", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGotoTrace(index uint) {
	response := safeGet("bridgemain.h/GuiGotoTrace", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiShowTrace() {
	response := safeGet("bridgemain.h/GuiShowTrace", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bridgemain) GuiGetMainThreadId() {
	response := safeGet("bridgemain.h/GuiGetMainThreadId", map[string]string{})
	if len(response) == 0 {
		return
	}
}
