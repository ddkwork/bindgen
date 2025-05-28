package sdk

type Window int // :52
const (
	DisassemblyWindow Window = iota // 0
	DumpWindow                      // 1
	StackWindow                     // 2
	GraphWindow                     // 3
	MemMapWindow                    // 4
	SymModWindow                    // 5
)

type gui struct{}

func (g *gui) SelectionGet(window Window, start *uint, end *uint) {
	response := safeGet("_scriptapi_gui.h/SelectionGet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) SelectionSet(window Window, start uint, end uint) {
	response := safeGet("_scriptapi_gui.h/SelectionSet", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) SelectionGetStart(window Window) {
	response := safeGet("_scriptapi_gui.h/SelectionGetStart", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) SelectionGetEnd(window Window) {
	response := safeGet("_scriptapi_gui.h/SelectionGetEnd", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) Message(message string) {
	response := safeGet("_scriptapi_gui.h/Message", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) MessageYesNo(message string) {
	response := safeGet("_scriptapi_gui.h/MessageYesNo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) InputLine(title string, text *int8) {
	response := safeGet("_scriptapi_gui.h/InputLine", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) InputValue(title string, value *uint) {
	response := safeGet("_scriptapi_gui.h/InputValue", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) Refresh() {
	response := safeGet("_scriptapi_gui.h/Refresh", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) AddQWidgetTab(qWidget *uintptr) {
	response := safeGet("_scriptapi_gui.h/AddQWidgetTab", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) ShowQWidgetTab(qWidget *uintptr) {
	response := safeGet("_scriptapi_gui.h/ShowQWidgetTab", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (g *gui) CloseQWidgetTab(qWidget *uintptr) {
	response := safeGet("_scriptapi_gui.h/CloseQWidgetTab", map[string]string{})
	if len(response) == 0 {
		return
	}
}
