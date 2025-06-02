package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// BookmarkInfo (:10 )
type BookmarkInfo struct {
	mod    [256]int8 // C type: char[256]
	rva    uint      // C type: duint
	manual bool      // C type: bool
}
type bookmark struct{}

// Set    c api name: Script::Bookmark::Set
// ┌────┬────────┬────────────────────────────────────────┬───────────────┐
// │ id │  name  │                 c type                 │    go type    │
// ├────┼────────┼────────────────────────────────────────┼───────────────┤
// │ 0  │ info   │ const Script::Bookmark::BookmarkInfo * │ *BookmarkInfo │
// ├────┼────────┼────────────────────────────────────────┼───────────────┤
// │    │ return │ bool                                   │ bool          │
// └────┴────────┴────────────────────────────────────────┴───────────────┘
func (b *bookmark) Set(info *BookmarkInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*BookmarkInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// Get    c api name: Script::Bookmark::Get
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (b *bookmark) Get(addr uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// GetInfo    c api name: Script::Bookmark::GetInfo
// ┌────┬────────┬────────────────┬───────────────┐
// │ id │  name  │     c type     │    go type    │
// ├────┼────────┼────────────────┼───────────────┤
// │ 0  │ addr   │ duint          │ uint          │
// ├────┼────────┼────────────────┼───────────────┤
// │ 1  │ info   │ BookmarkInfo * │ *BookmarkInfo │
// ├────┼────────┼────────────────┼───────────────┤
// │    │ return │ bool           │ bool          │
// └────┴────────┴────────────────┴───────────────┘
func (b *bookmark) GetInfo(addr uint, info *BookmarkInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*BookmarkInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// Delete    c api name: Script::Bookmark::Delete
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (b *bookmark) Delete(addr uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// DeleteRange    c api name: Script::Bookmark::DeleteRange
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ start  │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │ 1  │ end    │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (b *bookmark) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

// Clear    c api name: Script::Bookmark::Clear
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (b *bookmark) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// GetList    c api name: Script::Bookmark::GetList
// ┌────┬────────┬────────────┬───────────┐
// │ id │  name  │   c type   │  go type  │
// ├────┼────────┼────────────┼───────────┤
// │ 0  │ list   │ ListInfo * │ *ListInfo │
// ├────┼────────┼────────────┼───────────┤
// │    │ return │ bool       │ bool      │
// └────┴────────┴────────────┴───────────┘
func (b *bookmark) GetList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
