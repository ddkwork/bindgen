package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// ModuleInfo (:10 )
type ModuleInfo struct {
	base         uint      // C type: duint
	size         uint      // C type: duint
	entry        uint      // C type: duint
	sectionCount int       // C type: int
	name         [256]int8 // C type: char[256]
	path         [256]int8 // C type: char[256]
}

// ModuleSectionInfo (:20 )
type ModuleSectionInfo struct {
	addr uint     // C type: duint
	size uint     // C type: duint
	name [50]int8 // C type: char[50]
}

// ModuleExport (:27 )
type ModuleExport struct {
	ordinal         uint      // C type: duint
	rva             uint      // C type: duint
	va              uint      // C type: duint
	forwarded       bool      // C type: bool
	forwardName     [512]int8 // C type: char[512]
	name            [512]int8 // C type: char[512]
	undecoratedName [512]int8 // C type: char[512]
}

// ModuleImport (:38 )
type ModuleImport struct {
	iatRva          uint      // C type: duint
	iatVa           uint      // C type: duint
	ordinal         uint      // C type: duint
	name            [512]int8 // C type: char[512]
	undecoratedName [512]int8 // C type: char[512]
}
type module struct{}

// InfoFromAddr    c api name: Script::Module::InfoFromAddr
// ┌────┬────────┬──────────────┬─────────────┐
// │ id │  name  │    c type    │   go type   │
// ├────┼────────┼──────────────┼─────────────┤
// │ 0  │ addr   │ duint        │ uint        │
// ├────┼────────┼──────────────┼─────────────┤
// │ 1  │ info   │ ModuleInfo * │ *ModuleInfo │
// ├────┼────────┼──────────────┼─────────────┤
// │    │ return │ bool         │ bool        │
// └────┴────────┴──────────────┴─────────────┘
func (m *module) InfoFromAddr(addr uint, info *ModuleInfo) bool {
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
	return true
}

// InfoFromName    c api name: Script::Module::InfoFromName
// ┌────┬────────┬──────────────────────────────┬─────────────┐
// │ id │  name  │            c type            │   go type   │
// ├────┼────────┼──────────────────────────────┼─────────────┤
// │ 0  │ name   │ const Script::Module::char * │ *int8       │
// ├────┼────────┼──────────────────────────────┼─────────────┤
// │ 1  │ info   │ ModuleInfo *                 │ *ModuleInfo │
// ├────┼────────┼──────────────────────────────┼─────────────┤
// │    │ return │ bool                         │ bool        │
// └────┴────────┴──────────────────────────────┴─────────────┘
func (m *module) InfoFromName(name *int8, info *ModuleInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/InfoFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "info",
				Type:  "*ModuleInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// BaseFromAddr    c api name: Script::Module::BaseFromAddr
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (m *module) BaseFromAddr(addr uint) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/BaseFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// BaseFromName    c api name: Script::Module::BaseFromName
// ┌────┬────────┬──────────────────────────────┬─────────┐
// │ id │  name  │            c type            │ go type │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │ 0  │ name   │ const Script::Module::char * │ *int8   │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │    │ return │ duint                        │ uint    │
// └────┴────────┴──────────────────────────────┴─────────┘
func (m *module) BaseFromName(name *int8) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/BaseFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// SizeFromAddr    c api name: Script::Module::SizeFromAddr
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (m *module) SizeFromAddr(addr uint) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SizeFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// SizeFromName    c api name: Script::Module::SizeFromName
// ┌────┬────────┬──────────────────────────────┬─────────┐
// │ id │  name  │            c type            │ go type │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │ 0  │ name   │ const Script::Module::char * │ *int8   │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │    │ return │ duint                        │ uint    │
// └────┴────────┴──────────────────────────────┴─────────┘
func (m *module) SizeFromName(name *int8) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SizeFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// NameFromAddr    c api name: Script::Module::NameFromAddr
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │ 1  │ name   │ char * │ *int8   │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (m *module) NameFromAddr(addr uint, name *int8) bool {
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
	return true
}

// PathFromAddr    c api name: Script::Module::PathFromAddr
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │ 1  │ path   │ char * │ *int8   │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (m *module) PathFromAddr(addr uint, path *int8) bool {
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
	return true
}

// PathFromName    c api name: Script::Module::PathFromName
// ┌────┬────────┬──────────────────────────────┬─────────┐
// │ id │  name  │            c type            │ go type │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │ 0  │ name   │ const Script::Module::char * │ *int8   │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │ 1  │ path   │ char *                       │ *int8   │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │    │ return │ bool                         │ bool    │
// └────┴────────┴──────────────────────────────┴─────────┘
func (m *module) PathFromName(name *int8, path *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/PathFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "path",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", path),
			},
		},
	))).Request()
	return true
}

// EntryFromAddr    c api name: Script::Module::EntryFromAddr
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (m *module) EntryFromAddr(addr uint) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/EntryFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// EntryFromName    c api name: Script::Module::EntryFromName
// ┌────┬────────┬──────────────────────────────┬─────────┐
// │ id │  name  │            c type            │ go type │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │ 0  │ name   │ const Script::Module::char * │ *int8   │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │    │ return │ duint                        │ uint    │
// └────┴────────┴──────────────────────────────┴─────────┘
func (m *module) EntryFromName(name *int8) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/EntryFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// SectionCountFromAddr    c api name: Script::Module::SectionCountFromAddr
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ addr   │ duint  │ uint    │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ int    │ int     │
// └────┴────────┴────────┴─────────┘
func (m *module) SectionCountFromAddr(addr uint) int {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionCountFromAddr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return 0
}

// SectionCountFromName    c api name: Script::Module::SectionCountFromName
// ┌────┬────────┬──────────────────────────────┬─────────┐
// │ id │  name  │            c type            │ go type │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │ 0  │ name   │ const Script::Module::char * │ *int8   │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │    │ return │ int                          │ int     │
// └────┴────────┴──────────────────────────────┴─────────┘
func (m *module) SectionCountFromName(name *int8) int {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionCountFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	return 0
}

// SectionFromAddr    c api name: Script::Module::SectionFromAddr
// ┌────┬─────────┬─────────────────────┬────────────────────┐
// │ id │  name   │       c type        │      go type       │
// ├────┼─────────┼─────────────────────┼────────────────────┤
// │ 0  │ addr    │ duint               │ uint               │
// ├────┼─────────┼─────────────────────┼────────────────────┤
// │ 1  │ number  │ int                 │ int                │
// ├────┼─────────┼─────────────────────┼────────────────────┤
// │ 2  │ section │ ModuleSectionInfo * │ *ModuleSectionInfo │
// ├────┼─────────┼─────────────────────┼────────────────────┤
// │    │ return  │ bool                │ bool               │
// └────┴─────────┴─────────────────────┴────────────────────┘
func (m *module) SectionFromAddr(addr uint, number int, section *ModuleSectionInfo) bool {
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
	return true
}

// SectionFromName    c api name: Script::Module::SectionFromName
// ┌────┬─────────┬──────────────────────────────┬────────────────────┐
// │ id │  name   │            c type            │      go type       │
// ├────┼─────────┼──────────────────────────────┼────────────────────┤
// │ 0  │ name    │ const Script::Module::char * │ *int8              │
// ├────┼─────────┼──────────────────────────────┼────────────────────┤
// │ 1  │ number  │ int                          │ int                │
// ├────┼─────────┼──────────────────────────────┼────────────────────┤
// │ 2  │ section │ ModuleSectionInfo *          │ *ModuleSectionInfo │
// ├────┼─────────┼──────────────────────────────┼────────────────────┤
// │    │ return  │ bool                         │ bool               │
// └────┴─────────┴──────────────────────────────┴────────────────────┘
func (m *module) SectionFromName(name *int8, number int, section *ModuleSectionInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
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
	return true
}

// SectionListFromAddr    c api name: Script::Module::SectionListFromAddr
// ┌────┬────────┬────────────┬───────────┐
// │ id │  name  │   c type   │  go type  │
// ├────┼────────┼────────────┼───────────┤
// │ 0  │ addr   │ duint      │ uint      │
// ├────┼────────┼────────────┼───────────┤
// │ 1  │ list   │ ListInfo * │ *ListInfo │
// ├────┼────────┼────────────┼───────────┤
// │    │ return │ bool       │ bool      │
// └────┴────────┴────────────┴───────────┘
func (m *module) SectionListFromAddr(addr uint, list *ListInfo) bool {
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
	return true
}

// SectionListFromName    c api name: Script::Module::SectionListFromName
// ┌────┬────────┬──────────────────────────────┬───────────┐
// │ id │  name  │            c type            │  go type  │
// ├────┼────────┼──────────────────────────────┼───────────┤
// │ 0  │ name   │ const Script::Module::char * │ *int8     │
// ├────┼────────┼──────────────────────────────┼───────────┤
// │ 1  │ list   │ ListInfo *                   │ *ListInfo │
// ├────┼────────┼──────────────────────────────┼───────────┤
// │    │ return │ bool                         │ bool      │
// └────┴────────┴──────────────────────────────┴───────────┘
func (m *module) SectionListFromName(name *int8, list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/SectionListFromName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	return true
}

// GetMainModuleInfo    c api name: Script::Module::GetMainModuleInfo
// ┌────┬────────┬──────────────┬─────────────┐
// │ id │  name  │    c type    │   go type   │
// ├────┼────────┼──────────────┼─────────────┤
// │ 0  │ info   │ ModuleInfo * │ *ModuleInfo │
// ├────┼────────┼──────────────┼─────────────┤
// │    │ return │ bool         │ bool        │
// └────┴────────┴──────────────┴─────────────┘
func (m *module) GetMainModuleInfo(info *ModuleInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*ModuleInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	return true
}

// GetMainModuleBase    c api name: Script::Module::GetMainModuleBase
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (m *module) GetMainModuleBase() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleBase").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// GetMainModuleSize    c api name: Script::Module::GetMainModuleSize
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (m *module) GetMainModuleSize() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleSize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// GetMainModuleEntry    c api name: Script::Module::GetMainModuleEntry
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (m *module) GetMainModuleEntry() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleEntry").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// GetMainModuleSectionCount    c api name: Script::Module::GetMainModuleSectionCount
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ int    │ int     │
// └────┴────────┴────────┴─────────┘
func (m *module) GetMainModuleSectionCount() int {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleSectionCount").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return 0
}

// GetMainModuleName    c api name: Script::Module::GetMainModuleName
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ name   │ char * │ *int8   │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (m *module) GetMainModuleName(name *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleName").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "name",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", name),
			},
		},
	))).Request()
	return true
}

// GetMainModulePath    c api name: Script::Module::GetMainModulePath
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │ 0  │ path   │ char * │ *int8   │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ bool   │ bool    │
// └────┴────────┴────────┴─────────┘
func (m *module) GetMainModulePath(path *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModulePath").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "path",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", path),
			},
		},
	))).Request()
	return true
}

// GetMainModuleSectionList    c api name: Script::Module::GetMainModuleSectionList
// ┌────┬────────┬────────────┬───────────┐
// │ id │  name  │   c type   │  go type  │
// ├────┼────────┼────────────┼───────────┤
// │ 0  │ list   │ ListInfo * │ *ListInfo │
// ├────┼────────┼────────────┼───────────┤
// │    │ return │ bool       │ bool      │
// └────┴────────┴────────────┴───────────┘
func (m *module) GetMainModuleSectionList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetMainModuleSectionList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	return true
}

// GetList    c api name: Script::Module::GetList
// ┌────┬────────┬────────────┬───────────┐
// │ id │  name  │   c type   │  go type  │
// ├────┼────────┼────────────┼───────────┤
// │ 0  │ list   │ ListInfo * │ *ListInfo │
// ├────┼────────┼────────────┼───────────┤
// │    │ return │ bool       │ bool      │
// └────┴────────┴────────────┴───────────┘
func (m *module) GetList(list *ListInfo) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_module.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: fmt.Sprintf("%v", list),
			},
		},
	))).Request()
	return true
}

// GetExports    c api name: Script::Module::GetExports
// ┌────┬────────┬────────────────────────────────────┬─────────────┐
// │ id │  name  │               c type               │   go type   │
// ├────┼────────┼────────────────────────────────────┼─────────────┤
// │ 0  │ mod    │ const Script::Module::ModuleInfo * │ *ModuleInfo │
// ├────┼────────┼────────────────────────────────────┼─────────────┤
// │ 1  │ list   │ ListInfo *                         │ *ListInfo   │
// ├────┼────────┼────────────────────────────────────┼─────────────┤
// │    │ return │ bool                               │ bool        │
// └────┴────────┴────────────────────────────────────┴─────────────┘
func (m *module) GetExports(mod *ModuleInfo, list *ListInfo) bool {
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
	return true
}

// GetImports    c api name: Script::Module::GetImports
// ┌────┬────────┬────────────────────────────────────┬─────────────┐
// │ id │  name  │               c type               │   go type   │
// ├────┼────────┼────────────────────────────────────┼─────────────┤
// │ 0  │ mod    │ const Script::Module::ModuleInfo * │ *ModuleInfo │
// ├────┼────────┼────────────────────────────────────┼─────────────┤
// │ 1  │ list   │ ListInfo *                         │ *ListInfo   │
// ├────┼────────┼────────────────────────────────────┼─────────────┤
// │    │ return │ bool                               │ bool        │
// └────┴────────┴────────────────────────────────────┴─────────────┘
func (m *module) GetImports(mod *ModuleInfo, list *ListInfo) bool {
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
	return true
}
