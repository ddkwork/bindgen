package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/std/mylog"
)

type SymbolType int //:10
const (
	Function SymbolType = iota
	Import1
	Export2
)

// SymbolInfo (:17 )
type SymbolInfo struct {
	mod    [256]int8  // C type: char[256]
	rva    uint       // C type: duint
	name   [256]int8  // C type: char[256]
	manual bool       // C type: bool
	Type   SymbolType // C type: SymbolType
}
type symbol struct{}

// GetList    c api name: Script::Symbol::GetList
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬───────────┐
// │ id │         name         │                            c type                            │  go type  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │ 0  │ list                 │ ListInfo *                                                   │ *ListInfo │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ return               │ bool                                                         │ bool      │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.currentFile  │ include\_scriptapi_symbol.h                                  │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.name         │ GetList                                                      │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.mangledName  │ ?GetList@Symbol@Script@@YA_NPEAUListInfo@@@Z                 │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.file         │                                                              │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.includedFrom │                                                              │           │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼───────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │           │
// │    │                      │ ptapi_symbol.h                                               │           │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴───────────┘
func (s *symbol) GetList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_symbol.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
