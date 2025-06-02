package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type misc struct{}

// ParseExpression    c api name: Script::Misc::ParseExpression
// ┌────┬────────────┬────────────────────────────┬─────────┐
// │ id │    name    │           c type           │ go type │
// ├────┼────────────┼────────────────────────────┼─────────┤
// │ 0  │ expression │ const Script::Misc::char * │ *int8   │
// ├────┼────────────┼────────────────────────────┼─────────┤
// │ 1  │ value      │ Script::Misc::duint *      │ *uint   │
// ├────┼────────────┼────────────────────────────┼─────────┤
// │    │ return     │ bool                       │ bool    │
// └────┴────────────┴────────────────────────────┴─────────┘
func (m *misc) ParseExpression(expression *int8, value *uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/ParseExpression").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "expression",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", expression),
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

// RemoteGetProcAddress    c api name: Script::Misc::RemoteGetProcAddress
// ┌────┬────────┬────────────────────────────┬─────────┐
// │ id │  name  │           c type           │ go type │
// ├────┼────────┼────────────────────────────┼─────────┤
// │ 0  │ module │ const Script::Misc::char * │ *int8   │
// ├────┼────────┼────────────────────────────┼─────────┤
// │ 1  │ api    │ const Script::Misc::char * │ *int8   │
// ├────┼────────┼────────────────────────────┼─────────┤
// │    │ return │ duint                      │ uint    │
// └────┴────────┴────────────────────────────┴─────────┘
func (m *misc) RemoteGetProcAddress(module *int8, api *int8) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/RemoteGetProcAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "module",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", module),
			},
			{
				Name:  "api",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", api),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// ResolveLabel    c api name: Script::Misc::ResolveLabel
// ┌────┬────────┬────────────────────────────┬─────────┐
// │ id │  name  │           c type           │ go type │
// ├────┼────────┼────────────────────────────┼─────────┤
// │ 0  │ label  │ const Script::Misc::char * │ *int8   │
// ├────┼────────┼────────────────────────────┼─────────┤
// │    │ return │ duint                      │ uint    │
// └────┴────────┴────────────────────────────┴─────────┘
func (m *misc) ResolveLabel(label *int8) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/ResolveLabel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "label",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", label),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// Alloc    c api name: Script::Misc::Alloc
// ┌────┬────────┬─────────────────────┬─────────┐
// │ id │  name  │       c type        │ go type │
// ├────┼────────┼─────────────────────┼─────────┤
// │ 0  │ size   │ Script::Misc::duint │ uint    │
// ├────┼────────┼─────────────────────┼─────────┤
// │    │ return │ void *              │ uintptr │
// └────┴────────┴─────────────────────┴─────────┘
func (m *misc) Alloc(size uint) uintptr {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/Alloc").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// Free    c api name: Script::Misc::Free
// ┌────┬────────┬──────────────────────┬─────────┐
// │ id │  name  │        c type        │ go type │
// ├────┼────────┼──────────────────────┼─────────┤
// │ 0  │ ptr    │ Script::Misc::void * │ uintptr │
// ├────┼────────┼──────────────────────┼─────────┤
// │    │ return │ void                 │         │
// └────┴────────┴──────────────────────┴─────────┘
func (m *misc) Free(ptr uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/Free").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "ptr",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", ptr),
			},
		},
	))).Request()
}
