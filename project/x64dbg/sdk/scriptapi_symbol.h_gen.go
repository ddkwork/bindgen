package sdk

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
	response := safeGet("_scriptapi_symbol.h/GetList", map[string]string{})
	if len(response) == 0 {
		return
	}
}
