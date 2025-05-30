package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/mylog"
)

// FunctionInfo (:10)
type FunctionInfo struct {
	Mod              [256]int8 // C type: char[256]
	RvaStart         uint      // C type: duint
	RvaEnd           uint      // C type: duint
	Manual           bool      // C type: bool
	Instructioncount uint      // C type: duint
}
type function struct{}

func (f *function) Add(info *FunctionInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Add").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "info",
				Type:  "*FunctionInfo ",
				Value: info,
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) Get(addr uint, start *uint, end *uint, instructionCount *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "start",
				Type:  "*uint ",
				Value: start,
			},
			Param{
				Name:  "end",
				Type:  "*uint ",
				Value: end,
			},
			Param{
				Name:  "instructionCount",
				Type:  "*uint ",
				Value: instructionCount,
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) GetInfo(addr uint, info *FunctionInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "info",
				Type:  "*FunctionInfo ",
				Value: info,
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) Overlaps(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Overlaps").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "start",
				Type:  "uint",
				Value: start,
			},
			Param{
				Name:  "end",
				Type:  "uint",
				Value: end,
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) Delete(address uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "address",
				Type:  "uint",
				Value: address,
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "start",
				Type:  "uint",
				Value: start,
			},
			Param{
				Name:  "end",
				Type:  "uint",
				Value: end,
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (f *function) GetList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: list,
			},
		},
	))).Request()
	// todo handle response into result
}
