package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type bridgegraph struct{}

// Free    c api name: Free
// ┌────┬───────────┬───────────────────────────┬────────────────────┐
// │ id │   name    │          c type           │      go type       │
// ├────┼───────────┼───────────────────────────┼────────────────────┤
// │ 0  │ graphList │ const BridgeCFGraphList * │ *BridgeCFGraphList │
// ├────┼───────────┼───────────────────────────┼────────────────────┤
// │    │ return    │ void                      │                    │
// └────┴───────────┴───────────────────────────┴────────────────────┘
func (b *bridgegraph) Free(graphList *BridgeCFGraphList) {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/Free").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "graphList",
				Type:  "*BridgeCFGraphList ",
				Value: fmt.Sprintf("%v", graphList),
			},
		},
	))).Request()
}

// ToNodeList    c api name: ToNodeList
// ┌────┬────────┬──────────────────┬──────────────────┐
// │ id │  name  │      c type      │     go type      │
// ├────┼────────┼──────────────────┼──────────────────┤
// │    │ return │ BridgeCFNodeList │ BridgeCFNodeList │
// └────┴────────┴──────────────────┴──────────────────┘
func (b *bridgegraph) ToNodeList() BridgeCFNodeList {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/ToNodeList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: BridgeCFNodeList")
}

// __debugbreak    c api name: __debugbreak
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (b *bridgegraph) __debugbreak() {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/__debugbreak").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// AddNode    c api name: AddNode
// ┌────┬────────┬──────────────────────┬───────────────┐
// │ id │  name  │        c type        │    go type    │
// ├────┼────────┼──────────────────────┼───────────────┤
// │ 0  │ node   │ const BridgeCFNode & │ *BridgeCFNode │
// ├────┼────────┼──────────────────────┼───────────────┤
// │    │ return │ void                 │               │
// └────┴────────┴──────────────────────┴───────────────┘
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
}

// AddParent    c api name: AddParent
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ child  │ int    │ int     │
// ├────┼────────┼────────┼─────────┤
// │ 1  │ parent │ int    │ int     │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
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
}

// ToGraphList    c api name: ToGraphList
// ┌────┬────────┬───────────────────┬───────────────────┐
// │ id │  name  │      c type       │      go type      │
// ├────┼────────┼───────────────────┼───────────────────┤
// │    │ return │ BridgeCFGraphList │ BridgeCFGraphList │
// └────┴────────┴───────────────────┴───────────────────┘
func (b *bridgegraph) ToGraphList() BridgeCFGraphList {
	Client.Post().Url("http://localhost:8888/bridgegraph.h/ToGraphList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: BridgeCFGraphList")
}
