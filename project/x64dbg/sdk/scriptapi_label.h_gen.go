package sdk

// LabelInfo (:10)
type LabelInfo struct {
	Mod    [256]int8 // C type: char[256]
	Rva    uint      // C type: duint
	Text   [256]int8 // C type: char[256]
	Manual bool      // C type: bool
}
type label struct{}

func (l *label) Set(info *LabelInfo) {
	response := safeGet("_scriptapi_label.h/Set", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (l *label) FromString(label string, addr *uint) {
	response := safeGet("_scriptapi_label.h/FromString", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (l *label) Get(addr uint, text *int8) {
	response := safeGet("_scriptapi_label.h/Get", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (l *label) IsTemporary(addr uint) {
	response := safeGet("_scriptapi_label.h/IsTemporary", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (l *label) GetInfo(addr uint, info *LabelInfo) {
	response := safeGet("_scriptapi_label.h/GetInfo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (l *label) Delete(addr uint) {
	response := safeGet("_scriptapi_label.h/Delete", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (l *label) DeleteRange(start uint, end uint) {
	response := safeGet("_scriptapi_label.h/DeleteRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (l *label) Clear() {
	response := safeGet("_scriptapi_label.h/Clear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (l *label) GetList(list *ListInfo) {
	response := safeGet("_scriptapi_label.h/GetList", map[string]string{})
	if len(response) == 0 {
		return
	}
}
