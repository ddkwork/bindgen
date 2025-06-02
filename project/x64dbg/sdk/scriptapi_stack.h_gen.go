package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type stack struct{}

// Pop    c api name: Script::Stack::Pop
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (s *stack) Pop() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_stack.h/Pop").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// Push    c api name: Script::Stack::Push
// ┌────┬────────┬──────────────────────┬─────────┐
// │ id │  name  │        c type        │ go type │
// ├────┼────────┼──────────────────────┼─────────┤
// │ 0  │ value  │ Script::Stack::duint │ uint    │
// ├────┼────────┼──────────────────────┼─────────┤
// │    │ return │ duint                │ uint    │
// └────┴────────┴──────────────────────┴─────────┘
func (s *stack) Push(value uint) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_stack.h/Push").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// Peek    c api name: Script::Stack::Peek
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │ 0  │ offset │ Script::Stack::int │ int     │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ duint              │ uint    │
// └────┴────────┴────────────────────┴─────────┘
func (s *stack) Peek(offset int) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_stack.h/Peek").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "offset",
				Type:  "int",
				Value: fmt.Sprintf("%v", offset),
			},
		},
	))).Request()
	panic("not support return type: duint")
}
