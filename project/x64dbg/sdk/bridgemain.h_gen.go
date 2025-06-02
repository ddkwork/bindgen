package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
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

// BridgeCFNodeList (:10 )
type BridgeCFNodeList struct {
	parentGraph  uint     // C type: duint
	start        uint     // C type: duint
	end          uint     // C type: duint
	brtrue       uint     // C type: duint
	brfalse      uint     // C type: duint
	icount       uint     // C type: duint
	terminal     bool     // C type: bool
	split        bool     // C type: bool
	indirectcall bool     // C type: bool
	userdata     uintptr  // C type: void *
	exits        ListInfo // C type: ListInfo
	instrs       ListInfo // C type: ListInfo
}

// BridgeCFGraphList (:26 )
type BridgeCFGraphList struct {
	entryPoint uint     // C type: duint
	userdata   uintptr  // C type: void *
	nodes      ListInfo // C type: ListInfo
}

// BridgeCFNode (D:\workspace\workspace\bindgen\project\x64dbg\include\bridgegraph.h:41 )
type BridgeCFNode struct {
	parentGraph  uint    // C type: duint
	start        uint    // C type: duint
	end          uint    // C type: duint
	brtrue       uint    // C type: duint
	brfalse      uint    // C type: duint
	icount       uint    // C type: duint
	terminal     bool    // C type: bool
	split        bool    // C type: bool
	indirectcall bool    // C type: bool
	userdata     uintptr // C type: void *
	exits        any     // C type: std::vector<duint>
	instrs       any     // C type: std::vector<BridgeCFInstruction>
}

// BridgeCFGraph (:112 )
type BridgeCFGraph struct {
	entryPoint uint    // C type: duint
	userdata   uintptr // C type: void *
	nodes      any     // C type: std::unordered_map<duint, BridgeCFNode>
	parents    any     // C type: std::unordered_map<duint, std::unordered_set<duint>>
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

// BridgeLoadLibraryCheckedW    c api name: BridgeLoadLibraryCheckedW
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ szDll                │ const wchar_t *           │ * rune  │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 1  │ allowFailure         │ bool                      │ bool    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ HMODULE                   │ uintptr │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ BridgeLoadLibraryCheckedW │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeLoadLibraryCheckedW │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
func (b *bridgemain) BridgeLoadLibraryCheckedW(szDll *rune, allowFailure bool) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeLoadLibraryCheckedW").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "szDll",
				Type:  "* rune",
				Value: fmt.Sprintf("%v", szDll),
			},
			{
				Name:  "allowFailure",
				Type:  "bool",
				Value: fmt.Sprintf("%v", allowFailure),
			},
		},
	))).Request()
	panic("not support return type: HMODULE")
}

// BridgeLoadLibraryCheckedA    c api name: BridgeLoadLibraryCheckedA
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ szDll                │ const char *              │ *int8   │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 1  │ allowFailure         │ bool                      │ bool    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ HMODULE                   │ uintptr │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ BridgeLoadLibraryCheckedA │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeLoadLibraryCheckedA │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
func (b *bridgemain) BridgeLoadLibraryCheckedA(szDll *int8, allowFailure bool) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeLoadLibraryCheckedA").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "szDll",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", szDll),
			},
			{
				Name:  "allowFailure",
				Type:  "bool",
				Value: fmt.Sprintf("%v", allowFailure),
			},
		},
	))).Request()
	panic("not support return type: HMODULE")
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

// Data    c api name: Data
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────────────┐
// │ id │         name         │                            c type                            │      go type      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ return               │ BridgeCFNodeList *                                           │ *BridgeCFNodeList │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.name         │ Data                                                         │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.mangledName  │ ?Data@?$BridgeList@UBridgeCFNodeList@@@@QEBAPEAUBridgeCFNod- │                   │
// │    │                      │ eList@@XZ                                                    │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.file         │                                                              │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │                   │
// │    │                      │ emain.h                                                      │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.expansionLoc │                                                              │                   │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────────────┘
func (b *bridgemain) Data() *BridgeCFNodeList {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Data").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: BridgeCFNodeList *")
}

// Count    c api name: Count
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ int                                                          │ int     │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ Count                                                        │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?Count@?$BridgeList@UBridgeCFNodeList@@@@QEBAHXZ             │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │         │
// │    │                      │ emain.h                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (b *bridgemain) Count() int {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Count").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return 0
}

// Cleanup    c api name: Cleanup
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ Cleanup                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?Cleanup@?$BridgeList@UBridgeCFNodeList@@@@QEAAXXZ           │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │         │
// │    │                      │ emain.h                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (b *bridgemain) Cleanup() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Cleanup").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// CopyData    c api name: CopyData
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────┐
// │ id │         name         │                            c type                            │  go type  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 0  │ listInfo             │ ListInfo *                                                   │ *ListInfo │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 1  │ listData             │ const std::vector<BridgeCFNodeList> &                        │ any       │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ return               │ bool                                                         │ bool      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.name         │ CopyData                                                     │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.mangledName  │ ?CopyData@?$BridgeList@UBridgeCFNodeList@@@@SA_NPEAUListInf- │           │
// │    │                      │ o@@AEBV?$vector@UBridgeCFNodeList@@V?$allocator@UBridgeCFNo- │           │
// │    │                      │ deList@@@std@@@std@@@Z                                       │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.file         │                                                              │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │           │
// │    │                      │ emain.h                                                      │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.expansionLoc │                                                              │           │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────┘
func (b *bridgemain) CopyData(listInfo *ListInfo, listData any) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/CopyData").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "listInfo",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", listInfo),
			},
			{
				Name:  "listData",
				Type:  "any",
				Value: fmt.Sprintf("%v", listData),
			},
		},
	))).Request()
	return true
}

// Free    c api name: Free
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬────────────────────┐
// │ id │         name         │                            c type                            │      go type       │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │ 0  │ graphList            │ const BridgeCFGraphList *                                    │ *BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ return               │ void                                                         │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.name         │ Free                                                         │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.mangledName  │ ?Free@BridgeCFGraph@@SAXPEBUBridgeCFGraphList@@@Z            │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.file         │                                                              │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │                    │
// │    │                      │ emain.h                                                      │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.expansionLoc │                                                              │                    │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴────────────────────┘
func (b *bridgemain) Free(graphList *BridgeCFGraphList) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Free").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graphList",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graphList),
			},
		},
	))).Request()
}

// ToVector    c api name: ToVector
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────┐
// │ id │         name         │                            c type                            │  go type  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 0  │ listInfo             │ const ListInfo *                                             │ *ListInfo │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 1  │ listData             │ std::vector<BridgeCFNodeList> &                              │ any       │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 2  │ freedata             │ bool                                                         │ bool      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ return               │ bool                                                         │ bool      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.name         │ ToVector                                                     │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.mangledName  │ ?ToVector@?$BridgeList@UBridgeCFNodeList@@@@SA_NPEBUListInf- │           │
// │    │                      │ o@@AEAV?$vector@UBridgeCFNodeList@@V?$allocator@UBridgeCFNo- │           │
// │    │                      │ deList@@@std@@@std@@_N@Z                                     │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.file         │                                                              │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │           │
// │    │                      │ emain.h                                                      │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.expansionLoc │                                                              │           │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────┘
func (b *bridgemain) ToVector(listInfo *ListInfo, listData any, freedata bool) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/ToVector").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "listInfo",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", listInfo),
			},
			{
				Name:  "listData",
				Type:  "any",
				Value: fmt.Sprintf("%v", listData),
			},
			{
				Name:  "freedata",
				Type:  "bool",
				Value: fmt.Sprintf("%v", freedata),
			},
		},
	))).Request()
	return true
}

// ToNodeList    c api name: ToNodeList
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬──────────────────┐
// │ id │         name         │                            c type                            │     go type      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼──────────────────┤
// │    │ return               │ BridgeCFNodeList                                             │ BridgeCFNodeList │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼──────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │                  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼──────────────────┤
// │    │ comment.name         │ ToNodeList                                                   │                  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼──────────────────┤
// │    │ comment.mangledName  │ ?ToNodeList@BridgeCFNode@@QEBA?AUBridgeCFNodeList@@XZ        │                  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼──────────────────┤
// │    │ comment.file         │                                                              │                  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼──────────────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │                  │
// │    │                      │ emain.h                                                      │                  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼──────────────────┤
// │    │ comment.expansionLoc │                                                              │                  │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴──────────────────┘
func (b *bridgemain) ToNodeList() BridgeCFNodeList {
	Client.Post().Url("http://localhost:8888/bridgemain.h/ToNodeList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: BridgeCFNodeList")
}

// __debugbreak    c api name: __debugbreak
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ __debugbreak                                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ __debugbreak                                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │         │
// │    │                      │ emain.h                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (b *bridgemain) __debugbreak() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/__debugbreak").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// AddNode    c api name: AddNode
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────────┐
// │ id │         name         │                            c type                            │    go type    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │ 0  │ node                 │ const BridgeCFNode &                                         │ *BridgeCFNode │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ return               │ void                                                         │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.name         │ AddNode                                                      │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.mangledName  │ ?AddNode@BridgeCFGraph@@QEAAXAEBUBridgeCFNode@@@Z            │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.file         │                                                              │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │               │
// │    │                      │ emain.h                                                      │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.expansionLoc │                                                              │               │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────────┘
func (b *bridgemain) AddNode(node *BridgeCFNode) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/AddNode").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "node",
				Type:  "*BridgeCFNode ",
				Value: fmt.Sprintf("%v", node),
			},
		},
	))).Request()
}

// AddParent    c api name: AddParent
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ child                │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ parent               │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ AddParent                                                    │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?AddParent@BridgeCFGraph@@QEAAX_K0@Z                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │         │
// │    │                      │ emain.h                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (b *bridgemain) AddParent(child uint, parent uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/AddParent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "child",
				Type:  "uint",
				Value: fmt.Sprintf("%v", child),
			},
			{
				Name:  "parent",
				Type:  "uint",
				Value: fmt.Sprintf("%v", parent),
			},
		},
	))).Request()
}

// ToGraphList    c api name: ToGraphList
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────────────┐
// │ id │         name         │                            c type                            │      go type      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ return               │ BridgeCFGraphList                                            │ BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.name         │ ToGraphList                                                  │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.mangledName  │ ?ToGraphList@BridgeCFGraph@@QEBA?AUBridgeCFGraphList@@XZ     │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.file         │                                                              │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │                   │
// │    │                      │ emain.h                                                      │                   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────────┤
// │    │ comment.expansionLoc │                                                              │                   │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────────────┘
func (b *bridgemain) ToGraphList() BridgeCFGraphList {
	Client.Post().Url("http://localhost:8888/bridgemain.h/ToGraphList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: BridgeCFGraphList")
}

// __builtin_memcpy    c api name: __builtin_memcpy
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │                      │ void *                                                       │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │                      │ const void *                                                 │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │                      │ unsigned long long                                           │ uint64  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void *                                                       │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ __builtin_memcpy                                             │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ __builtin_memcpy                                             │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │         │
// │    │                      │ emain.h                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (b *bridgemain) __builtin_memcpy(uintptr, uintptr, uint64) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/__builtin_memcpy").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v"),
			},
			{
				Name:  "",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v"),
			},
			{
				Name:  "",
				Type:  "uint64",
				Value: fmt.Sprintf("%v"),
			},
		},
	))).Request()
	panic("not support return type: void *")
}

// BridgeLoadLibraryCheckedW    c api name: BridgeLoadLibraryCheckedW
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ szDll                │ const wchar_t *           │ * rune  │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 1  │ allowFailure         │ bool                      │ bool    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ HMODULE                   │ uintptr │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ BridgeLoadLibraryCheckedW │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeLoadLibraryCheckedW │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
func (b *bridgemain) BridgeLoadLibraryCheckedW(szDll *rune, allowFailure bool) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeLoadLibraryCheckedW").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "szDll",
				Type:  "* rune",
				Value: fmt.Sprintf("%v", szDll),
			},
			{
				Name:  "allowFailure",
				Type:  "bool",
				Value: fmt.Sprintf("%v", allowFailure),
			},
		},
	))).Request()
	panic("not support return type: HMODULE")
}

// BridgeLoadLibraryCheckedA    c api name: BridgeLoadLibraryCheckedA
// ┌────┬──────────────────────┬───────────────────────────┬─────────┐
// │ id │         name         │          c type           │ go type │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 0  │ szDll                │ const char *              │ *int8   │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │ 1  │ allowFailure         │ bool                      │ bool    │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ return               │ HMODULE                   │ uintptr │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h      │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.name         │ BridgeLoadLibraryCheckedA │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.mangledName  │ BridgeLoadLibraryCheckedA │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.file         │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.includedFrom │                           │         │
// ├────┼──────────────────────┼───────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                           │         │
// └────┴──────────────────────┴───────────────────────────┴─────────┘
func (b *bridgemain) BridgeLoadLibraryCheckedA(szDll *int8, allowFailure bool) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeLoadLibraryCheckedA").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "szDll",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", szDll),
			},
			{
				Name:  "allowFailure",
				Type:  "bool",
				Value: fmt.Sprintf("%v", allowFailure),
			},
		},
	))).Request()
	panic("not support return type: HMODULE")
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

// CopyData    c api name: CopyData
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────┐
// │ id │         name         │                            c type                            │  go type  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 0  │ listInfo             │ ListInfo *                                                   │ *ListInfo │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 1  │ listData             │ const std::vector<BridgeCFNodeList> &                        │ any       │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ return               │ bool                                                         │ bool      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.name         │ CopyData                                                     │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.mangledName  │ ?CopyData@?$BridgeList@UBridgeCFNodeList@@@@SA_NPEAUListInf- │           │
// │    │                      │ o@@AEBV?$vector@UBridgeCFNodeList@@V?$allocator@UBridgeCFNo- │           │
// │    │                      │ deList@@@std@@@std@@@Z                                       │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.file         │                                                              │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │           │
// │    │                      │ emain.h                                                      │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.expansionLoc │                                                              │           │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────┘
func (b *bridgemain) CopyData(listInfo *ListInfo, listData any) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/CopyData").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "listInfo",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", listInfo),
			},
			{
				Name:  "listData",
				Type:  "any",
				Value: fmt.Sprintf("%v", listData),
			},
		},
	))).Request()
	return true
}

// Free    c api name: Free
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬────────────────────┐
// │ id │         name         │                            c type                            │      go type       │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │ 0  │ graphList            │ const BridgeCFGraphList *                                    │ *BridgeCFGraphList │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ return               │ void                                                         │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.name         │ Free                                                         │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.mangledName  │ ?Free@BridgeCFGraph@@SAXPEBUBridgeCFGraphList@@@Z            │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.file         │                                                              │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │                    │
// │    │                      │ emain.h                                                      │                    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼────────────────────┤
// │    │ comment.expansionLoc │                                                              │                    │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴────────────────────┘
func (b *bridgemain) Free(graphList *BridgeCFGraphList) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Free").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graphList",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graphList),
			},
		},
	))).Request()
}

// ToVector    c api name: ToVector
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────┐
// │ id │         name         │                            c type                            │  go type  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 0  │ listInfo             │ const ListInfo *                                             │ *ListInfo │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 1  │ listData             │ std::vector<BridgeCFNodeList> &                              │ any       │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 2  │ freedata             │ bool                                                         │ bool      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ return               │ bool                                                         │ bool      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.name         │ ToVector                                                     │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.mangledName  │ ?ToVector@?$BridgeList@UBridgeCFNodeList@@@@SA_NPEBUListInf- │           │
// │    │                      │ o@@AEAV?$vector@UBridgeCFNodeList@@V?$allocator@UBridgeCFNo- │           │
// │    │                      │ deList@@@std@@@std@@_N@Z                                     │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.file         │                                                              │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │           │
// │    │                      │ emain.h                                                      │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.expansionLoc │                                                              │           │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────┘
func (b *bridgemain) ToVector(listInfo *ListInfo, listData any, freedata bool) bool {
	Client.Post().Url("http://localhost:8888/bridgemain.h/ToVector").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "listInfo",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", listInfo),
			},
			{
				Name:  "listData",
				Type:  "any",
				Value: fmt.Sprintf("%v", listData),
			},
			{
				Name:  "freedata",
				Type:  "bool",
				Value: fmt.Sprintf("%v", freedata),
			},
		},
	))).Request()
	return true
}

// AddNode    c api name: AddNode
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────────┐
// │ id │         name         │                            c type                            │    go type    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │ 0  │ node                 │ const BridgeCFNode &                                         │ *BridgeCFNode │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ return               │ void                                                         │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.name         │ AddNode                                                      │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.mangledName  │ ?AddNode@BridgeCFGraph@@QEAAXAEBUBridgeCFNode@@@Z            │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.file         │                                                              │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │               │
// │    │                      │ emain.h                                                      │               │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────────┤
// │    │ comment.expansionLoc │                                                              │               │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────────┘
func (b *bridgemain) AddNode(node *BridgeCFNode) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/AddNode").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "node",
				Type:  "*BridgeCFNode ",
				Value: fmt.Sprintf("%v", node),
			},
		},
	))).Request()
}

// AddParent    c api name: AddParent
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ child                │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ parent               │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ AddParent                                                    │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?AddParent@BridgeCFGraph@@QEAAX_K0@Z                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │         │
// │    │                      │ emain.h                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (b *bridgemain) AddParent(child uint, parent uint) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/AddParent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "child",
				Type:  "uint",
				Value: fmt.Sprintf("%v", child),
			},
			{
				Name:  "parent",
				Type:  "uint",
				Value: fmt.Sprintf("%v", parent),
			},
		},
	))).Request()
}

// __builtin_memcpy    c api name: __builtin_memcpy
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │                      │ void *                                                       │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │                      │ const void *                                                 │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │                      │ unsigned long long                                           │ uint64  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void *                                                       │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\bridgemain.h                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ __builtin_memcpy                                             │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ __builtin_memcpy                                             │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │ D:\workspace\workspace\bindgen\project\x64dbg\include\bridg- │         │
// │    │                      │ emain.h                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │                                                              │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (b *bridgemain) __builtin_memcpy(uintptr, uintptr, uint64) uintptr {
	Client.Post().Url("http://localhost:8888/bridgemain.h/__builtin_memcpy").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v"),
			},
			{
				Name:  "",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v"),
			},
			{
				Name:  "",
				Type:  "uint64",
				Value: fmt.Sprintf("%v"),
			},
		},
	))).Request()
	panic("not support return type: void *")
}
