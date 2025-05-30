package sdk

import (
	"encoding/json"
	"fmt"
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
			{
				Name:  "info",
				Type:  "*FunctionInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) Get(addr uint, start *uint, end *uint, instructionCount *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "start",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", end),
			},
			{
				Name:  "instructionCount",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", instructionCount),
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) GetInfo(addr uint, info *FunctionInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*FunctionInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) Overlaps(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Overlaps").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) Delete(address uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "address",
				Type:  "uint",
				Value: fmt.Sprintf("%v", address),
			},
		},
	))).Request()
	// todo handle response into result
}

func (f *function) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
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
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}
