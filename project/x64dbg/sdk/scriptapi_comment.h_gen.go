package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// CommentInfo (:10 )
type CommentInfo struct {
	mod    [256]int8 // C type: char[256]
	rva    uint      // C type: duint
	text   [256]int8 // C type: char[256]
	manual bool      // C type: bool
}
type comment struct{}

// Set    c api name: Script::Comment::Set
// ┌────┬────────┬──────────────────────────────────────┬──────────────┐
// │ id │  name  │                c type                │   go type    │
// ├────┼────────┼──────────────────────────────────────┼──────────────┤
// │ 0  │ info   │ const Script::Comment::CommentInfo * │ *CommentInfo │
// ├────┼────────┼──────────────────────────────────────┼──────────────┤
// │    │ return │ bool                                 │ bool         │
// └────┴────────┴──────────────────────────────────────┴──────────────┘
func (c *comment) Set(info *CommentInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*CommentInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// Get    c api name: Script::Comment::Get
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │ 1  │ text   │ char * │ *int8   │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (c *comment) Get(addr uint, text *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// GetInfo    c api name: Script::Comment::GetInfo
// ┌────┬────────┬───────────────┬──────────────┐
// │ id │  name  │    c type     │   go type    │
// ├────┼────────┼───────────────┼──────────────┤
// │ 0  │ addr   │ duint         │ uint         │
// ├────┼────────┼───────────────┼──────────────┤
// │ 1  │ info   │ CommentInfo * │ *CommentInfo │
// ├────┼────────┼───────────────┼──────────────┤
// │    │ return │ bool          │ bool         │
// └────┴────────┴───────────────┴──────────────┘
func (c *comment) GetInfo(addr uint, info *CommentInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*CommentInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// Delete    c api name: Script::Comment::Delete
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (c *comment) Delete(addr uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// DeleteRange    c api name: Script::Comment::DeleteRange
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ start  │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │ 1  │ end    │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (c *comment) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// Clear    c api name: Script::Comment::Clear
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (c *comment) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GetList    c api name: Script::Comment::GetList
// ┌────┬────────┬────────────┬───────────┐
// │ id │  name  │   c type   │  go type  │
// ├────┼────────┼────────────┼───────────┤
// │ 0  │ list   │ ListInfo * │ *ListInfo │
// ├────┼────────┼────────────┼───────────┤
// │    │ return │ bool       │ bool      │
// └────┴────────┴────────────┴───────────┘
func (c *comment) GetList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
