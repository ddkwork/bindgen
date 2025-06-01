package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// ListInfo (D:\workspace\workspace\bindgen\project\x64dbg\include\bridgelist.h:4)
type ListInfo struct {
	Count int      // C type: int
	Size  uint     // C type: size_t
	Data  *uintptr // C type: void *
}

// BridgeCFInstruction (D:\workspace\workspace\bindgen\project\x64dbg\include\bridgegraph.h:4)
type BridgeCFInstruction struct {
	Addr uint     // C type: duint
	Data [15]byte // C type: unsigned char[15]
}

// BridgeCFNodeList (:10)
type BridgeCFNodeList struct {
	ParentGraph  uint     // C type: duint
	Start        uint     // C type: duint
	End          uint     // C type: duint
	Brtrue       uint     // C type: duint
	Brfalse      uint     // C type: duint
	Icount       uint     // C type: duint
	Terminal     bool     // C type: bool
	Split        bool     // C type: bool
	Indirectcall bool     // C type: bool
	Userdata     *uintptr // C type: void *
	Exits        ListInfo // C type: ListInfo
	Instrs       ListInfo // C type: ListInfo
}

// BridgeCFGraphList (:26)
type BridgeCFGraphList struct {
	EntryPoint uint     // C type: duint
	Userdata   *uintptr // C type: void *
	Nodes      ListInfo // C type: ListInfo
}

// BridgeCFNode (D:\workspace\workspace\bindgen\project\x64dbg\include\bridgegraph.h:41)
type BridgeCFNode struct {
	ParentGraph  uint     // C type: duint
	Start        uint     // C type: duint
	End          uint     // C type: duint
	Brtrue       uint     // C type: duint
	Brfalse      uint     // C type: duint
	Icount       uint     // C type: duint
	Terminal     bool     // C type: bool
	Split        bool     // C type: bool
	Indirectcall bool     // C type: bool
	Userdata     *uintptr // C type: void *
	Exits        any      // C type: std::vector<duint>
	Instrs       any      // C type: std::vector<BridgeCFInstruction>
}

// BridgeCFGraph (:112)
type BridgeCFGraph struct {
	EntryPoint uint     // C type: duint
	Userdata   *uintptr // C type: void *
	Nodes      any      // C type: std::unordered_map<duint, BridgeCFNode>
	Parents    any      // C type: std::unordered_map<duint, std::unordered_set<duint>>
}
type bridgemain struct{}

func (b *bridgemain) BridgeInit() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeInit").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeLoadLibraryCheckedW(szDll *wint8_t, allowFailure bool) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeLoadLibraryCheckedW").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "szDll",
				Type:  "*wint8_t ",
				Value: fmt.Sprintf("%v", szDll),
			},
			{
				Name:  "allowFailure",
				Type:  "bool",
				Value: fmt.Sprintf("%v", allowFailure),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) BridgeLoadLibraryCheckedA(szDll string, allowFailure bool) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/BridgeLoadLibraryCheckedA").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "szDll",
				Type:  "string",
				Value: fmt.Sprintf("%v", szDll),
			},
			{
				Name:  "allowFailure",
				Type:  "bool",
				Value: fmt.Sprintf("%v", allowFailure),
			},
		},
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

func (b *bridgemain) Data() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Data").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) Count() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Count").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) Cleanup() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Cleanup").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) CopyData(listInfo *ListInfo, listData any) {
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
	// todo handle response into result
}

func (b *bridgemain) Free(graphList *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/Free").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (b *bridgemain) ToVector(listInfo *ListInfo, listData any, freedata bool) {
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
	// todo handle response into result
}

func (b *bridgemain) ToNodeList() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/ToNodeList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) __debugbreak() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/__debugbreak").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

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
	// todo handle response into result
}

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
	// todo handle response into result
}

func (b *bridgemain) ToGraphList() {
	Client.Post().Url("http://localhost:8888/bridgemain.h/ToGraphList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgemain) __builtin_memcpy(*uintptr, *uintptr, uint64) {
	Client.Post().Url("http://localhost:8888/bridgemain.h/__builtin_memcpy").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v"),
			},
			{
				Name:  "",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v"),
			},
			{
				Name:  "",
				Type:  "uint64",
				Value: fmt.Sprintf("%v"),
			},
		},
	))).Request()
	// todo handle response into result
}
