package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// LabelInfo (:10 )
type LabelInfo struct {
	mod    [256]int8 // C type: char[256]
	rva    uint      // C type: duint
	text   [256]int8 // C type: char[256]
	manual bool      // C type: bool
}
type label struct{}

// Set    c api name: Script::Label::Set
// ┌────┬────────┬──────────────────────────────────┬────────────┐
// │ id │  name  │              c type              │  go type   │
// ├────┼────────┼──────────────────────────────────┼────────────┤
// │ 0  │ info   │ const Script::Label::LabelInfo * │ *LabelInfo │
// ├────┼────────┼──────────────────────────────────┼────────────┤
// │    │ return │ bool                             │ bool       │
// └────┴────────┴──────────────────────────────────┴────────────┘
func (l *label) Set(info *LabelInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*LabelInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// FromString    c api name: Script::Label::FromString
// ┌────┬────────┬─────────────────────────────┬─────────┐
// │ id │  name  │           c type            │ go type │
// ├────┼────────┼─────────────────────────────┼─────────┤
// │ 0  │ label  │ const Script::Label::char * │ *int8   │
// ├────┼────────┼─────────────────────────────┼─────────┤
// │ 1  │ addr   │ Script::Label::duint *      │ *uint   │
// ├────┼────────┼─────────────────────────────┼─────────┤
// │    │ return │ bool                        │ bool    │
// └────┴────────┴─────────────────────────────┴─────────┘
func (l *label) FromString(label *int8, addr *uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/FromString").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "label",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", label),
			},
			{
				Name:  "addr",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// Get    c api name: Script::Label::Get
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Label::duint  │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 1  │ text   │ Script::Label::char * │ *int8   │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ bool                  │ bool    │
// └────┴────────┴───────────────────────┴─────────┘
func (l *label) Get(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	return true
}

// IsTemporary    c api name: Script::Label::IsTemporary
// ┌────┬────────┬──────────────────────┬─────────┐
// │ id │  name  │        c type        │ go type │
// ├────┼────────┼──────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Label::duint │ uint    │
// ├────┼────────┼──────────────────────┼─────────┤
// │    │ return │ bool                 │ bool    │
// └────┴────────┴──────────────────────┴─────────┘
func (l *label) IsTemporary(addr uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/IsTemporary").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// GetInfo    c api name: Script::Label::GetInfo
// ┌────┬────────┬────────────────────────────┬────────────┐
// │ id │  name  │           c type           │  go type   │
// ├────┼────────┼────────────────────────────┼────────────┤
// │ 0  │ addr   │ Script::Label::duint       │ uint       │
// ├────┼────────┼────────────────────────────┼────────────┤
// │ 1  │ info   │ Script::Label::LabelInfo * │ *LabelInfo │
// ├────┼────────┼────────────────────────────┼────────────┤
// │    │ return │ bool                       │ bool       │
// └────┴────────┴────────────────────────────┴────────────┘
func (l *label) GetInfo(addr uint, info *LabelInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*LabelInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// Delete    c api name: Script::Label::Delete
// ┌────┬────────┬──────────────────────┬─────────┐
// │ id │  name  │        c type        │ go type │
// ├────┼────────┼──────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Label::duint │ uint    │
// ├────┼────────┼──────────────────────┼─────────┤
// │    │ return │ bool                 │ bool    │
// └────┴────────┴──────────────────────┴─────────┘
func (l *label) Delete(addr uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// DeleteRange    c api name: Script::Label::DeleteRange
// ┌────┬────────┬──────────────────────┬─────────┐
// │ id │  name  │        c type        │ go type │
// ├────┼────────┼──────────────────────┼─────────┤
// │ 0  │ start  │ Script::Label::duint │ uint    │
// ├────┼────────┼──────────────────────┼─────────┤
// │ 1  │ end    │ Script::Label::duint │ uint    │
// ├────┼────────┼──────────────────────┼─────────┤
// │    │ return │ void                 │         │
// └────┴────────┴──────────────────────┴─────────┘
func (l *label) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// Clear    c api name: Script::Label::Clear
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (l *label) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GetList    c api name: Script::Label::GetList
// ┌────┬────────┬───────────────────────────┬───────────┐
// │ id │  name  │          c type           │  go type  │
// ├────┼────────┼───────────────────────────┼───────────┤
// │ 0  │ list   │ Script::Label::ListInfo * │ *ListInfo │
// ├────┼────────┼───────────────────────────┼───────────┤
// │    │ return │ bool                      │ bool      │
// └────┴────────┴───────────────────────────┴───────────┘
func (l *label) GetList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
