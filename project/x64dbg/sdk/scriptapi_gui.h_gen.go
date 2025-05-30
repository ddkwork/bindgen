package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/mylog"
)

type Window int // :52
const (
	DisassemblyWindow Window = iota // 0
	DumpWindow                      // 1
	StackWindow                     // 2
	GraphWindow                     // 3
	MemMapWindow                    // 4
	SymModWindow                    // 5
)

type gui struct{}

func (g *gui) SelectionGet(window Window, start *uint, end *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionGet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "window",
				Type:  "Window",
				Value: window,
			},
			Param{
				Name:  "start",
				Type:  "*uint ",
				Value: start,
			},
			Param{
				Name:  "end",
				Type:  "*uint ",
				Value: end,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) SelectionSet(window Window, start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionSet").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "window",
				Type:  "Window",
				Value: window,
			},
			Param{
				Name:  "start",
				Type:  "uint",
				Value: start,
			},
			Param{
				Name:  "end",
				Type:  "uint",
				Value: end,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) SelectionGetStart(window Window) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionGetStart").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "window",
				Type:  "Window",
				Value: window,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) SelectionGetEnd(window Window) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionGetEnd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "window",
				Type:  "Window",
				Value: window,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) Message(message string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/Message").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "message",
				Type:  "string",
				Value: message,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) MessageYesNo(message string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/MessageYesNo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "message",
				Type:  "string",
				Value: message,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) InputLine(title string, text *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/InputLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "title",
				Type:  "string",
				Value: title,
			},
			Param{
				Name:  "text",
				Type:  "*int8 ",
				Value: text,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) InputValue(title string, value *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/InputValue").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "title",
				Type:  "string",
				Value: title,
			},
			Param{
				Name:  "value",
				Type:  "*uint ",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) Refresh() {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/Refresh").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (g *gui) AddQWidgetTab(qWidget *uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/AddQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: qWidget,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) ShowQWidgetTab(qWidget *uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/ShowQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: qWidget,
			},
		},
	))).Request()
	// todo handle response into result
}
func (g *gui) CloseQWidgetTab(qWidget *uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/CloseQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: qWidget,
			},
		},
	))).Request()
	// todo handle response into result
}
