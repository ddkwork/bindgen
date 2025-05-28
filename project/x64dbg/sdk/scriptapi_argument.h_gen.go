package sdk

// ArgumentInfo (:10)
type ArgumentInfo struct {
	Mod              [256]int8 // C type: char[256]
	RvaStart         uint      // C type: duint
	RvaEnd           uint      // C type: duint
	Manual           bool      // C type: bool
	Instructioncount uint      // C type: duint
}
type argument struct{}

func (a *argument) Add(info *ArgumentInfo) {
	response := safeGet("_scriptapi_argument.h/Add", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *argument) Get(addr uint, start *uint, end *uint, instructionCount *uint) {
	response := safeGet("_scriptapi_argument.h/Get", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *argument) GetInfo(addr uint, info *ArgumentInfo) {
	response := safeGet("_scriptapi_argument.h/GetInfo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *argument) Overlaps(start uint, end uint) {
	response := safeGet("_scriptapi_argument.h/Overlaps", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *argument) Delete(address uint) {
	response := safeGet("_scriptapi_argument.h/Delete", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *argument) DeleteRange(start uint, end uint, deleteManual bool) {
	response := safeGet("_scriptapi_argument.h/DeleteRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *argument) Clear() {
	response := safeGet("_scriptapi_argument.h/Clear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (a *argument) GetList(list *ListInfo) {
	response := safeGet("_scriptapi_argument.h/GetList", map[string]string{})
	if len(response) == 0 {
		return
	}
}
