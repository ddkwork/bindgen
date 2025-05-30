package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type misc struct{}

func (m *misc) ParseExpression(expression string, value *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/ParseExpression").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "expression",
				Type:  "string",
				Value: fmt.Sprintf("%v", expression),
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

func (m *misc) RemoteGetProcAddress(module string, api string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/RemoteGetProcAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "module",
				Type:  "string",
				Value: fmt.Sprintf("%v", module),
			},
			{
				Name:  "api",
				Type:  "string",
				Value: fmt.Sprintf("%v", api),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *misc) ResolveLabel(label string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/ResolveLabel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "label",
				Type:  "string",
				Value: fmt.Sprintf("%v", label),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *misc) Alloc(size uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/Alloc").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (m *misc) Free(ptr *uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/Free").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
