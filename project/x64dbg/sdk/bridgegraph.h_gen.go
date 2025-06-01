package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

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
