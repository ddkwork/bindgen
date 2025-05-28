package sdk

type pattern struct{}

func (p *pattern) Find(data *byte, datasize uint, pattern string) {
	response := safeGet("_scriptapi_pattern.h/Find", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (p *pattern) FindMem(start uint, size uint, pattern string) {
	response := safeGet("_scriptapi_pattern.h/FindMem", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (p *pattern) Write(data *byte, datasize uint, pattern string) {
	response := safeGet("_scriptapi_pattern.h/Write", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (p *pattern) WriteMem(start uint, size uint, pattern string) {
	response := safeGet("_scriptapi_pattern.h/WriteMem", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (p *pattern) SearchAndReplace(data *byte, datasize uint, searchpattern string, replacepattern string) {
	response := safeGet("_scriptapi_pattern.h/SearchAndReplace", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (p *pattern) SearchAndReplaceMem(start uint, size uint, searchpattern string, replacepattern string) {
	response := safeGet("_scriptapi_pattern.h/SearchAndReplaceMem", map[string]string{})
	if len(response) == 0 {
		return
	}
}
