package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/mylog"
)

type misc struct{}

func (m *misc) ParseExpression(expression string, value *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/ParseExpression").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "expression",
				Type:  "string",
				Value: expression,
			},
			Param{
				Name:  "value",
				Type:  "*uint ",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *misc) RemoteGetProcAddress(module string, api string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/RemoteGetProcAddress").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "module",
				Type:  "string",
				Value: module,
			},
			Param{
				Name:  "api",
				Type:  "string",
				Value: api,
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *misc) ResolveLabel(label string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/ResolveLabel").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "label",
				Type:  "string",
				Value: label,
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *misc) Alloc(size uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/Alloc").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "size",
				Type:  "uint",
				Value: size,
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *misc) Free(ptr *uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_misc.h/Free").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "ptr",
				Type:  "*uintptr ",
				Value: ptr,
			},
		},
	))).Request()
	// todo handle response into result
}
