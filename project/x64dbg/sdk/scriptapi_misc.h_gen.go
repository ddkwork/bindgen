package sdk

type misc struct{}

func (m *misc) ParseExpression(expression string, value *uint) {
	response := safeGet("_scriptapi_misc.h/ParseExpression", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *misc) RemoteGetProcAddress(module string, api string) {
	response := safeGet("_scriptapi_misc.h/RemoteGetProcAddress", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *misc) ResolveLabel(label string) {
	response := safeGet("_scriptapi_misc.h/ResolveLabel", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *misc) Alloc(size uint) {
	response := safeGet("_scriptapi_misc.h/Alloc", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *misc) Free(ptr *uintptr) {
	response := safeGet("_scriptapi_misc.h/Free", map[string]string{})
	if len(response) == 0 {
		return
	}
}
