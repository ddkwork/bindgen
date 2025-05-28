package sdk

// FunctionInfo (:10)
type FunctionInfo struct {
	Mod              [256]int8 // C type: char[256]
	RvaStart         uint      // C type: duint
	RvaEnd           uint      // C type: duint
	Manual           bool      // C type: bool
	Instructioncount uint      // C type: duint
}
type function struct{}

func (f *function) Add(info *FunctionInfo) {
	response := safeGet("_scriptapi_function.h/Add", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *function) Get(addr uint, start *uint, end *uint, instructionCount *uint) {
	response := safeGet("_scriptapi_function.h/Get", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *function) GetInfo(addr uint, info *FunctionInfo) {
	response := safeGet("_scriptapi_function.h/GetInfo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *function) Overlaps(start uint, end uint) {
	response := safeGet("_scriptapi_function.h/Overlaps", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *function) Delete(address uint) {
	response := safeGet("_scriptapi_function.h/Delete", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *function) DeleteRange(start uint, end uint) {
	response := safeGet("_scriptapi_function.h/DeleteRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *function) Clear() {
	response := safeGet("_scriptapi_function.h/Clear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *function) GetList(list *ListInfo) {
	response := safeGet("_scriptapi_function.h/GetList", map[string]string{})
	if len(response) == 0 {
		return
	}
}
