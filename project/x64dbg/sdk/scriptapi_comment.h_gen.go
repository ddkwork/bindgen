package sdk

// CommentInfo (:10)
type CommentInfo struct {
	Mod    [256]int8 // C type: char[256]
	Rva    uint      // C type: duint
	Text   [256]int8 // C type: char[256]
	Manual bool      // C type: bool
}
type comment struct{}

func (c *comment) Set(info *CommentInfo) {
	response := safeGet("_scriptapi_comment.h/Set", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (c *comment) Get(addr uint, text *int8) {
	response := safeGet("_scriptapi_comment.h/Get", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (c *comment) GetInfo(addr uint, info *CommentInfo) {
	response := safeGet("_scriptapi_comment.h/GetInfo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (c *comment) Delete(addr uint) {
	response := safeGet("_scriptapi_comment.h/Delete", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (c *comment) DeleteRange(start uint, end uint) {
	response := safeGet("_scriptapi_comment.h/DeleteRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (c *comment) Clear() {
	response := safeGet("_scriptapi_comment.h/Clear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (c *comment) GetList(list *ListInfo) {
	response := safeGet("_scriptapi_comment.h/GetList", map[string]string{})
	if len(response) == 0 {
		return
	}
}
