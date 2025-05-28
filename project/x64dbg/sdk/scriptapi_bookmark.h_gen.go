package sdk

// BookmarkInfo (:10)
type BookmarkInfo struct {
	Mod    [256]int8 // C type: char[256]
	Rva    uint      // C type: duint
	Manual bool      // C type: bool
}
type bookmark struct{}

func (b *bookmark) Set(info *BookmarkInfo) {
	response := safeGet("_scriptapi_bookmark.h/Set", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bookmark) Get(addr uint) {
	response := safeGet("_scriptapi_bookmark.h/Get", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bookmark) GetInfo(addr uint, info *BookmarkInfo) {
	response := safeGet("_scriptapi_bookmark.h/GetInfo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bookmark) Delete(addr uint) {
	response := safeGet("_scriptapi_bookmark.h/Delete", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bookmark) DeleteRange(start uint, end uint) {
	response := safeGet("_scriptapi_bookmark.h/DeleteRange", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bookmark) Clear() {
	response := safeGet("_scriptapi_bookmark.h/Clear", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (b *bookmark) GetList(list *ListInfo) {
	response := safeGet("_scriptapi_bookmark.h/GetList", map[string]string{})
	if len(response) == 0 {
		return
	}
}
