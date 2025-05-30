package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type SymbolType int // :10
const (
	Function SymbolType = iota // 0
	Import                     // 1
	Export                     // 2
)

// SymbolInfo (:17)
type SymbolInfo struct {
	Mod    [256]int8  // C type: char[256]
	Rva    uint       // C type: duint
	Name   [256]int8  // C type: char[256]
	Manual bool       // C type: bool
	Type   SymbolType // C type: SymbolType
}
type symbol struct{}

func (s *symbol) GetList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_symbol.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
