package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// BridgeCFNodeList (:10)
type BridgeCFNodeList struct {
	ParentGraph  int      // C type: int
	Start        int      // C type: int
	End          int      // C type: int
	Brtrue       int      // C type: int
	Brfalse      int      // C type: int
	Icount       int      // C type: int
	Terminal     bool     // C type: bool
	Split        bool     // C type: bool
	Indirectcall bool     // C type: bool
	Userdata     *uintptr // C type: void *
	Exits        int      // C type: int
	Instrs       int      // C type: int
}

// BridgeCFNode (D:\workspace\workspace\mcp\x64dbg\include\bridgegraph.h:41)
type BridgeCFNode struct {
	ParentGraph  int      // C type: int
	Start        int      // C type: int
	End          int      // C type: int
	Brtrue       int      // C type: int
	Brfalse      int      // C type: int
	Icount       int      // C type: int
	Terminal     bool     // C type: bool
	Split        bool     // C type: bool
	Indirectcall bool     // C type: bool
	Userdata     *uintptr // C type: void *
	Exits        int      // C type: int
	Instrs       uintptr  // C type: std::vector<BridgeCFInstruction>
}

// BridgeCFGraph (:112)
type BridgeCFGraph struct {
	EntryPoint int      // C type: int
	Userdata   *uintptr // C type: void *
	Nodes      int      // C type: int
}
type bridgegraph struct{}

func (b *bridgegraph) Free(graphList *uintptr) {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/Free").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (b *bridgegraph) ToNodeList() {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/ToNodeList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgegraph) __debugbreak() {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/__debugbreak").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (b *bridgegraph) AddNode(node *BridgeCFNode) {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/AddNode").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (b *bridgegraph) AddParent(child int, parent int) {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/AddParent").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "child",
				Type:  "int",
				Value: fmt.Sprintf("%v", child),
			},
			{
				Name:  "parent",
				Type:  "int",
				Value: fmt.Sprintf("%v", parent),
			},
		},
	))).Request()
	// todo handle response into result
}

func (b *bridgegraph) ToGraphList() {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/ToGraphList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
