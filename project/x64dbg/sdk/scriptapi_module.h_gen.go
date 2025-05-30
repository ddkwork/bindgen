package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

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
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/InfoFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*ModuleInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) InfoFromName(name string, info *ModuleInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/InfoFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "info",
				Type:  "*ModuleInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) BaseFromAddr(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/BaseFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) BaseFromName(name string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/BaseFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) SizeFromAddr(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SizeFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) SizeFromName(name string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SizeFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) NameFromAddr(addr uint, name *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/NameFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) PathFromAddr(addr uint, path *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/PathFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "path",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", path),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) PathFromName(name string, path *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/PathFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "path",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", path),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) EntryFromAddr(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/EntryFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) EntryFromName(name string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/EntryFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) SectionCountFromAddr(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionCountFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) SectionCountFromName(name string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionCountFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) SectionFromAddr(addr uint, number int, section *ModuleSectionInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "number",
				Type:  "int",
				Value: fmt.Sprintf("%v", number),
			},
			{
				Name:  "section",
				Type:  "*ModuleSectionInfo ",
				Value: fmt.Sprintf("%v", section),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) SectionFromName(name string, number int, section *ModuleSectionInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "number",
				Type:  "int",
				Value: fmt.Sprintf("%v", number),
			},
			{
				Name:  "section",
				Type:  "*ModuleSectionInfo ",
				Value: fmt.Sprintf("%v", section),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) SectionListFromAddr(addr uint, list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionListFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) SectionListFromName(name string, list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionListFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "string",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) GetMainModuleInfo(info *ModuleInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*ModuleInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) GetMainModuleBase() {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleBase").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (m *module) GetMainModuleSize() {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleSize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (m *module) GetMainModuleEntry() {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleEntry").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (m *module) GetMainModuleSectionCount() {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleSectionCount").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (m *module) GetMainModuleName(name *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) GetMainModulePath(path *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModulePath").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "path",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", path),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) GetMainModuleSectionList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleSectionList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) GetList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) GetExports(mod *ModuleInfo, list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetExports").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "mod",
				Type:  "*ModuleInfo ",
				Value: fmt.Sprintf("%v", mod),
			},
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}

func (m *module) GetImports(mod *ModuleInfo, list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetImports").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "mod",
				Type:  "*ModuleInfo ",
				Value: fmt.Sprintf("%v", mod),
			},
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	// todo handle response into result
}
