package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/std/mylog"
)

type Window int //:52
const (
	DisassemblyWindow Window = iota
	DumpWindow1
	StackWindow2
	GraphWindow3
	MemMapWindow4
	SymModWindow5
)

type gui struct{}

// SelectionGet    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::SelectionGet
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ window               │ Window                                                       │ Window  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ start                │ duint *                                                      │ *uint   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │ end                  │ duint *                                                      │ *uint   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ bool                                                         │ bool    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ SelectionGet                                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?SelectionGet@Gui@Script@@YA_NW4Window@12@PEA_K1@Z           │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) SelectionGet(window Window, start *uint, end *uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "window",
				Type:  "Window",
				Value: fmt.Sprintf("%v", window),
			},
			{
				Name:  "start",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	return true
}

// SelectionSet    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::SelectionSet
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ window               │ Window                                                       │ Window  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ start                │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 2  │ end                  │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ bool                                                         │ bool    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ SelectionSet                                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?SelectionSet@Gui@Script@@YA_NW4Window@12@_K1@Z              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) SelectionSet(window Window, start uint, end uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionSet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "window",
				Type:  "Window",
				Value: fmt.Sprintf("%v", window),
			},
			{
				Name:  "start",
				Type:  "uint",
				Value: fmt.Sprintf("%v", start),
			},
			{
				Name:  "end",
				Type:  "uint",
				Value: fmt.Sprintf("%v", end),
			},
		},
	))).Request()
	return true
}

// SelectionGetStart    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::SelectionGetStart
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ window               │ Window                                                       │ Window  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ SelectionGetStart                                            │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?SelectionGetStart@Gui@Script@@YA_KW4Window@12@@Z            │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) SelectionGetStart(window Window) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionGetStart").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "window",
				Type:  "Window",
				Value: fmt.Sprintf("%v", window),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// SelectionGetEnd    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::SelectionGetEnd
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ window               │ Window                                                       │ Window  │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ duint                                                        │ uint    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ SelectionGetEnd                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?SelectionGetEnd@Gui@Script@@YA_KW4Window@12@@Z              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) SelectionGetEnd(window Window) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionGetEnd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "window",
				Type:  "Window",
				Value: fmt.Sprintf("%v", window),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// Message    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::Message
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ message              │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ Message                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?Message@Gui@Script@@YAXPEBD@Z                               │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) Message(message *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/Message").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
}

// MessageYesNo    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::MessageYesNo
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ message              │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ bool                                                         │ bool    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ MessageYesNo                                                 │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?MessageYesNo@Gui@Script@@YA_NPEBD@Z                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) MessageYesNo(message *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/MessageYesNo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	return true
}

// InputLine    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::InputLine
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ title                │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ text                 │ char *                                                       │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ bool                                                         │ bool    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ InputLine                                                    │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?InputLine@Gui@Script@@YA_NPEBDPEAD@Z                        │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) InputLine(title *int8, text *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/InputLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	return true
}

// InputValue    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::InputValue
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ title                │ const char *                                                 │ *int8   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 1  │ value                │ duint *                                                      │ *uint   │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ bool                                                         │ bool    │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ InputValue                                                   │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?InputValue@Gui@Script@@YA_NPEBDPEA_K@Z                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) InputValue(title *int8, value *uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/InputValue").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "value",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// Refresh    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::Refresh
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ Refresh                                                      │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?Refresh@Gui@Script@@YAXXZ                                   │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) Refresh() {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/Refresh").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// AddQWidgetTab    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::AddQWidgetTab
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *                                                       │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ AddQWidgetTab                                                │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?AddQWidgetTab@Gui@Script@@YAXPEAX@Z                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) AddQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/AddQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}

// ShowQWidgetTab    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::ShowQWidgetTab
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *                                                       │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ ShowQWidgetTab                                               │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?ShowQWidgetTab@Gui@Script@@YAXPEAX@Z                        │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) ShowQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/ShowQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}

// CloseQWidgetTab    c api name: Script::Gui::Disassembly::Dump::Stack::Graph::MemMap::SymMod::Gui::CloseQWidgetTab
// ┌────┬──────────────────────┬──────────────────────────────────────────────────────────────┬─────────┐
// │ id │         name         │                            c type                            │ go type │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │ 0  │ qWidget              │ void *                                                       │ uintptr │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ return               │ void                                                         │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.currentFile  │ include\_scriptapi_gui.h                                     │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.name         │ CloseQWidgetTab                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.mangledName  │ ?CloseQWidgetTab@Gui@Script@@YAXPEAX@Z                       │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.file         │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.includedFrom │                                                              │         │
// ├────┼──────────────────────┼──────────────────────────────────────────────────────────────┼─────────┤
// │    │ comment.expansionLoc │ D:\workspace\workspace\bindgen\project\x64dbg\include\_scri- │         │
// │    │                      │ ptapi_gui.h                                                  │         │
// └────┴──────────────────────┴──────────────────────────────────────────────────────────────┴─────────┘
func (g *gui) CloseQWidgetTab(qWidget uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/CloseQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
}
