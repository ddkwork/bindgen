package sdk

type FlagEnum int // :10
const (
	ZF FlagEnum = iota // 0
	OF                 // 1
	CF                 // 2
	PF                 // 3
	SF                 // 4
	TF                 // 5
	AF                 // 6
	DF                 // 7
	IF                 // 8
)

type flag struct{}

func (f *flag) Get(flag FlagEnum) {
	response := safeGet("_scriptapi_flag.h/Get", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) Set(flag FlagEnum, value bool) {
	response := safeGet("_scriptapi_flag.h/Set", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetZF() {
	response := safeGet("_scriptapi_flag.h/GetZF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetZF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetZF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetOF() {
	response := safeGet("_scriptapi_flag.h/GetOF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetOF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetOF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetCF() {
	response := safeGet("_scriptapi_flag.h/GetCF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetCF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetCF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetPF() {
	response := safeGet("_scriptapi_flag.h/GetPF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetPF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetPF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetSF() {
	response := safeGet("_scriptapi_flag.h/GetSF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetSF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetSF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetTF() {
	response := safeGet("_scriptapi_flag.h/GetTF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetTF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetTF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetAF() {
	response := safeGet("_scriptapi_flag.h/GetAF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetAF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetAF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetDF() {
	response := safeGet("_scriptapi_flag.h/GetDF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetDF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetDF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) GetIF() {
	response := safeGet("_scriptapi_flag.h/GetIF", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (f *flag) SetIF(value bool) {
	response := safeGet("_scriptapi_flag.h/SetIF", map[string]string{})
	if len(response) == 0 {
		return
	}
}
