package sdk

type HardwareType int // :10
const (
	HardwareAccess  HardwareType = iota // 0
	HardwareWrite                       // 1
	HardwareExecute                     // 2
)

type debug struct{}

func (d *debug) Wait() {
	response := safeGet("_scriptapi_debug.h/Wait", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) Run() {
	response := safeGet("_scriptapi_debug.h/Run", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) Pause() {
	response := safeGet("_scriptapi_debug.h/Pause", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) Stop() {
	response := safeGet("_scriptapi_debug.h/Stop", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) StepIn() {
	response := safeGet("_scriptapi_debug.h/StepIn", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) StepOver() {
	response := safeGet("_scriptapi_debug.h/StepOver", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) StepOut() {
	response := safeGet("_scriptapi_debug.h/StepOut", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) SetBreakpoint(address uint) {
	response := safeGet("_scriptapi_debug.h/SetBreakpoint", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) DeleteBreakpoint(address uint) {
	response := safeGet("_scriptapi_debug.h/DeleteBreakpoint", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) DisableBreakpoint(address uint) {
	response := safeGet("_scriptapi_debug.h/DisableBreakpoint", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) SetHardwareBreakpoint(address uint, Type HardwareType) {
	response := safeGet("_scriptapi_debug.h/SetHardwareBreakpoint", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (d *debug) DeleteHardwareBreakpoint(address uint) {
	response := safeGet("_scriptapi_debug.h/DeleteHardwareBreakpoint", map[string]string{})
	if len(response) == 0 {
		return
	}
}
