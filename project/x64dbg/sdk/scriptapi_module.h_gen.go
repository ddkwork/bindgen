package sdk

// ModuleInfo (:10)
type ModuleInfo struct {
	Base         uint      // C type: duint
	Size         uint      // C type: duint
	Entry        uint      // C type: duint
	SectionCount int       // C type: int
	Name         [256]int8 // C type: char[256]
	Path         [256]int8 // C type: char[256]
}

// ModuleSectionInfo (:20)
type ModuleSectionInfo struct {
	Addr uint     // C type: duint
	Size uint     // C type: duint
	Name [50]int8 // C type: char[50]
}

// ModuleExport (:27)
type ModuleExport struct {
	Ordinal         uint      // C type: duint
	Rva             uint      // C type: duint
	Va              uint      // C type: duint
	Forwarded       bool      // C type: bool
	ForwardName     [512]int8 // C type: char[512]
	Name            [512]int8 // C type: char[512]
	UndecoratedName [512]int8 // C type: char[512]
}

// ModuleImport (:38)
type ModuleImport struct {
	IatRva          uint      // C type: duint
	IatVa           uint      // C type: duint
	Ordinal         uint      // C type: duint
	Name            [512]int8 // C type: char[512]
	UndecoratedName [512]int8 // C type: char[512]
}
type module struct{}

func (m *module) InfoFromAddr(addr uint, info *ModuleInfo) {
	response := safeGet("_scriptapi_module.h/InfoFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) InfoFromName(name string, info *ModuleInfo) {
	response := safeGet("_scriptapi_module.h/InfoFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) BaseFromAddr(addr uint) {
	response := safeGet("_scriptapi_module.h/BaseFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) BaseFromName(name string) {
	response := safeGet("_scriptapi_module.h/BaseFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) SizeFromAddr(addr uint) {
	response := safeGet("_scriptapi_module.h/SizeFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) SizeFromName(name string) {
	response := safeGet("_scriptapi_module.h/SizeFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) NameFromAddr(addr uint, name *int8) {
	response := safeGet("_scriptapi_module.h/NameFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) PathFromAddr(addr uint, path *int8) {
	response := safeGet("_scriptapi_module.h/PathFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) PathFromName(name string, path *int8) {
	response := safeGet("_scriptapi_module.h/PathFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) EntryFromAddr(addr uint) {
	response := safeGet("_scriptapi_module.h/EntryFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) EntryFromName(name string) {
	response := safeGet("_scriptapi_module.h/EntryFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) SectionCountFromAddr(addr uint) {
	response := safeGet("_scriptapi_module.h/SectionCountFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) SectionCountFromName(name string) {
	response := safeGet("_scriptapi_module.h/SectionCountFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) SectionFromAddr(addr uint, number int, section *ModuleSectionInfo) {
	response := safeGet("_scriptapi_module.h/SectionFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) SectionFromName(name string, number int, section *ModuleSectionInfo) {
	response := safeGet("_scriptapi_module.h/SectionFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) SectionListFromAddr(addr uint, list *ListInfo) {
	response := safeGet("_scriptapi_module.h/SectionListFromAddr", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) SectionListFromName(name string, list *ListInfo) {
	response := safeGet("_scriptapi_module.h/SectionListFromName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetMainModuleInfo(info *ModuleInfo) {
	response := safeGet("_scriptapi_module.h/GetMainModuleInfo", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetMainModuleBase() {
	response := safeGet("_scriptapi_module.h/GetMainModuleBase", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetMainModuleSize() {
	response := safeGet("_scriptapi_module.h/GetMainModuleSize", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetMainModuleEntry() {
	response := safeGet("_scriptapi_module.h/GetMainModuleEntry", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetMainModuleSectionCount() {
	response := safeGet("_scriptapi_module.h/GetMainModuleSectionCount", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetMainModuleName(name *int8) {
	response := safeGet("_scriptapi_module.h/GetMainModuleName", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetMainModulePath(path *int8) {
	response := safeGet("_scriptapi_module.h/GetMainModulePath", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetMainModuleSectionList(list *ListInfo) {
	response := safeGet("_scriptapi_module.h/GetMainModuleSectionList", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetList(list *ListInfo) {
	response := safeGet("_scriptapi_module.h/GetList", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetExports(mod *ModuleInfo, list *ListInfo) {
	response := safeGet("_scriptapi_module.h/GetExports", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (m *module) GetImports(mod *ModuleInfo, list *ListInfo) {
	response := safeGet("_scriptapi_module.h/GetImports", map[string]string{})
	if len(response) == 0 {
		return
	}
}
