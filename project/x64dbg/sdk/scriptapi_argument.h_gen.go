package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// ArgumentInfo (:10 )
type ArgumentInfo struct {
	mod              [256]int8 // C type: char[256]
	rvaStart         uint      // C type: duint
	rvaEnd           uint      // C type: duint
	manual           bool      // C type: bool
	instructioncount uint      // C type: duint
}
type argument struct{}

// Add    c api name: Script::Argument::Add
// ┌────┬────────┬────────────────────────────────────────┬───────────────┐
// │ id │  name  │                 c type                 │    go type    │
// ├────┼────────┼────────────────────────────────────────┼───────────────┤
// │ 0  │ info   │ const Script::Argument::ArgumentInfo * │ *ArgumentInfo │
// ├────┼────────┼────────────────────────────────────────┼───────────────┤
// │    │ return │ bool                                   │ bool          │
// └────┴────────┴────────────────────────────────────────┴───────────────┘
func (a *argument) Add(info *ArgumentInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Add").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*ArgumentInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// Get    c api name: Script::Argument::Get
// ┌────┬──────────────────┬───────────────────────────┬─────────┐
// │ id │       name       │          c type           │ go type │
// ├────┼──────────────────┼───────────────────────────┼─────────┤
// │ 0  │ addr             │ Script::Argument::duint   │ uint    │
// ├────┼──────────────────┼───────────────────────────┼─────────┤
// │ 1  │ start            │ Script::Argument::duint * │ *uint   │
// ├────┼──────────────────┼───────────────────────────┼─────────┤
// │ 2  │ end              │ Script::Argument::duint * │ *uint   │
// ├────┼──────────────────┼───────────────────────────┼─────────┤
// │ 3  │ instructionCount │ Script::Argument::duint * │ *uint   │
// ├────┼──────────────────┼───────────────────────────┼─────────┤
// │    │ return           │ bool                      │ bool    │
// └────┴──────────────────┴───────────────────────────┴─────────┘
func (a *argument) Get(addr uint, start *uint, end *uint, instructionCount *uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// GetInfo    c api name: Script::Argument::GetInfo
// ┌────┬────────┬──────────────────────────────────┬───────────────┐
// │ id │  name  │              c type              │    go type    │
// ├────┼────────┼──────────────────────────────────┼───────────────┤
// │ 0  │ addr   │ Script::Argument::duint          │ uint          │
// ├────┼────────┼──────────────────────────────────┼───────────────┤
// │ 1  │ info   │ Script::Argument::ArgumentInfo * │ *ArgumentInfo │
// ├────┼────────┼──────────────────────────────────┼───────────────┤
// │    │ return │ bool                             │ bool          │
// └────┴────────┴──────────────────────────────────┴───────────────┘
func (a *argument) GetInfo(addr uint, info *ArgumentInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*ArgumentInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// Overlaps    c api name: Script::Argument::Overlaps
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ start  │ Script::Argument::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 1  │ end    │ Script::Argument::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (a *argument) Overlaps(start uint, end uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Overlaps").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// Delete    c api name: Script::Argument::Delete
// ┌────┬─────────┬─────────────────────────┬─────────┐
// │ id │  name   │         c type          │ go type │
// ├────┼─────────┼─────────────────────────┼─────────┤
// │ 0  │ address │ Script::Argument::duint │ uint    │
// ├────┼─────────┼─────────────────────────┼─────────┤
// │    │ return  │ bool                    │ bool    │
// └────┴─────────┴─────────────────────────┴─────────┘
func (a *argument) Delete(address uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// DeleteRange    c api name: Script::Argument::DeleteRange
// ┌────┬──────────────┬─────────────────────────┬─────────┐
// │ id │     name     │         c type          │ go type │
// ├────┼──────────────┼─────────────────────────┼─────────┤
// │ 0  │ start        │ Script::Argument::duint │ uint    │
// ├────┼──────────────┼─────────────────────────┼─────────┤
// │ 1  │ end          │ Script::Argument::duint │ uint    │
// ├────┼──────────────┼─────────────────────────┼─────────┤
// │ 2  │ deleteManual │ Script::Argument::bool  │ bool    │
// ├────┼──────────────┼─────────────────────────┼─────────┤
// │    │ return       │ void                    │         │
// └────┴──────────────┴─────────────────────────┴─────────┘
func (a *argument) DeleteRange(start uint, end uint, deleteManual bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
			{
				Name:  "deleteManual",
				Type:  "bool",
				Value: fmt.Sprintf("%v", deleteManual),
			},
		},
	))).Request()
}

// Clear    c api name: Script::Argument::Clear
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (a *argument) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GetList    c api name: Script::Argument::GetList
// ┌────┬────────┬──────────────────────────────┬───────────┐
// │ id │  name  │            c type            │  go type  │
// ├────┼────────┼──────────────────────────────┼───────────┤
// │ 0  │ list   │ Script::Argument::ListInfo * │ *ListInfo │
// ├────┼────────┼──────────────────────────────┼───────────┤
// │    │ return │ bool                         │ bool      │
// └────┴────────┴──────────────────────────────┴───────────┘
func (a *argument) GetList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
