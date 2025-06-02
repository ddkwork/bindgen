package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// FunctionInfo (:10 )
type FunctionInfo struct {
	mod              [256]int8 // C type: char[256]
	rvaStart         uint      // C type: duint
	rvaEnd           uint      // C type: duint
	manual           bool      // C type: bool
	instructioncount uint      // C type: duint
}
type function struct{}

// Add    c api name: Script::Function::Add
// ┌────┬────────┬──────────────────────┬───────────────┐
// │ id │  name  │        c type        │    go type    │
// ├────┼────────┼──────────────────────┼───────────────┤
// │ 0  │ info   │ const FunctionInfo * │ *FunctionInfo │
// ├────┼────────┼──────────────────────┼───────────────┤
// │    │ return │ bool                 │ bool          │
// └────┴────────┴──────────────────────┴───────────────┘
func (f *function) Add(info *FunctionInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Add").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*FunctionInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// Get    c api name: Script::Function::Get
// ┌────┬──────────────────┬─────────┬─────────┐
// │ id │       name       │ c type  │ go type │
// ├────┼──────────────────┼─────────┼─────────┤
// │ 0  │ addr             │ duint   │ uint    │
// ├────┼──────────────────┼─────────┼─────────┤
// │ 1  │ start            │ duint * │ *uint   │
// ├────┼──────────────────┼─────────┼─────────┤
// │ 2  │ end              │ duint * │ *uint   │
// ├────┼──────────────────┼─────────┼─────────┤
// │ 3  │ instructionCount │ duint * │ *uint   │
// ├────┼──────────────────┼─────────┼─────────┤
// │    │ return           │ bool    │ bool    │
// └────┴──────────────────┴─────────┴─────────┘
func (f *function) Get(addr uint, start *uint, end *uint, instructionCount *uint) bool {
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
	return true
}

// GetInfo    c api name: Script::Function::GetInfo
// ┌────┬────────┬────────────────┬───────────────┐
// │ id │  name  │     c type     │    go type    │
// ├────┼────────┼────────────────┼───────────────┤
// │ 0  │ addr   │ duint          │ uint          │
// ├────┼────────┼────────────────┼───────────────┤
// │ 1  │ info   │ FunctionInfo * │ *FunctionInfo │
// ├────┼────────┼────────────────┼───────────────┤
// │    │ return │ bool           │ bool          │
// └────┴────────┴────────────────┴───────────────┘
func (f *function) GetInfo(addr uint, info *FunctionInfo) bool {
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
	return true
}

// Overlaps    c api name: Script::Function::Overlaps
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ start  │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │ 1  │ end    │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (f *function) Overlaps(start uint, end uint) bool {
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
	return true
}

// Delete    c api name: Script::Function::Delete
// ┌────┬─────────┬────────┬─────────┐
// │ id │  name   │ c type │ go type │
// ├────┼─────────┼────────┼─────────┤
// │ 0  │ address │ duint  │ uint    │
// ├────┼─────────┼────────┼─────────┤
// │    │ return  │ bool   │ bool    │
// └────┴─────────┴────────┴─────────┘
func (f *function) Delete(address uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "address",
				Type:  "uint",
				Value: fmt.Sprintf("%v", address),
			},
		},
	))).Request()
	return true
}

// DeleteRange    c api name: Script::Function::DeleteRange
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ start  │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │ 1  │ end    │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
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
}

// Clear    c api name: Script::Function::Clear
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (f *function) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GetList    c api name: Script::Function::GetList
// ┌────┬────────┬────────────┬───────────┐
// │ id │  name  │   c type   │  go type  │
// ├────┼────────┼────────────┼───────────┤
// │ 0  │ list   │ ListInfo * │ *ListInfo │
// ├────┼────────┼────────────┼───────────┤
// │    │ return │ bool       │ bool      │
// └────┴────────┴────────────┴───────────┘
func (f *function) GetList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_function.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	return true
}
