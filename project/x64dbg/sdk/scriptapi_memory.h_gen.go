package sdk

type memory struct{}

func (m *memory) Read(addr uint, data *uintptr, size uint, sizeRead *uint) {
	response := safeGet("_scriptapi_memory.h/Read", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) Write(addr uint, data *uintptr, size uint, sizeWritten *uint) {
	response := safeGet("_scriptapi_memory.h/Write", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) IsValidPtr(addr uint) {
	response := safeGet("_scriptapi_memory.h/IsValidPtr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) RemoteAlloc(addr uint, size uint) {
	response := safeGet("_scriptapi_memory.h/RemoteAlloc", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) RemoteFree(addr uint) {
	response := safeGet("_scriptapi_memory.h/RemoteFree", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) GetProtect(addr uint, reserved bool, cache bool) {
	response := safeGet("_scriptapi_memory.h/GetProtect", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) SetProtect(addr uint, protect uint, size uint) {
	response := safeGet("_scriptapi_memory.h/SetProtect", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) GetBase(addr uint, reserved bool, cache bool) {
	response := safeGet("_scriptapi_memory.h/GetBase", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) GetSize(addr uint, reserved bool, cache bool) {
	response := safeGet("_scriptapi_memory.h/GetSize", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) ReadByte(addr uint) {
	response := safeGet("_scriptapi_memory.h/ReadByte", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) WriteByte(addr uint, data byte) {
	response := safeGet("_scriptapi_memory.h/WriteByte", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) ReadWord(addr uint) {
	response := safeGet("_scriptapi_memory.h/ReadWord", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) WriteWord(addr uint, data uint16) {
	response := safeGet("_scriptapi_memory.h/WriteWord", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) ReadDword(addr uint) {
	response := safeGet("_scriptapi_memory.h/ReadDword", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) WriteDword(addr uint, data uint) {
	response := safeGet("_scriptapi_memory.h/WriteDword", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) ReadQword(addr uint) {
	response := safeGet("_scriptapi_memory.h/ReadQword", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) WriteQword(addr uint, data uint64) {
	response := safeGet("_scriptapi_memory.h/WriteQword", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) ReadPtr(addr uint) {
	response := safeGet("_scriptapi_memory.h/ReadPtr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *memory) WritePtr(addr uint, data uint) {
	response := safeGet("_scriptapi_memory.h/WritePtr", map[string]string{})
	if len(response) == 0 {
		return
	}
}
