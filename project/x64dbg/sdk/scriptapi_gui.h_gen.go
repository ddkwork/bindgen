package sdk

import (
	"encoding/json"
	"fmt"
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
	// todo handle response into result
}

func (g *gui) SelectionSet(window Window, start uint, end uint) {
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
	// todo handle response into result
}

func (g *gui) SelectionGetStart(window Window) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionGetStart").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "window",
				Type:  "Window",
				Value: fmt.Sprintf("%v", window),
			},
		},
	))).Request()
	// todo handle response into result
}

func (g *gui) SelectionGetEnd(window Window) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/SelectionGetEnd").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "window",
				Type:  "Window",
				Value: fmt.Sprintf("%v", window),
			},
		},
	))).Request()
	// todo handle response into result
}

func (g *gui) Message(message string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/Message").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "string",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	// todo handle response into result
}

func (g *gui) MessageYesNo(message string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/MessageYesNo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "message",
				Type:  "string",
				Value: fmt.Sprintf("%v", message),
			},
		},
	))).Request()
	// todo handle response into result
}

func (g *gui) InputLine(title string, text *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/InputLine").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "text",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", text),
			},
		},
	))).Request()
	// todo handle response into result
}

func (g *gui) InputValue(title string, value *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/InputValue").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "title",
				Type:  "string",
				Value: fmt.Sprintf("%v", title),
			},
			{
				Name:  "value",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", value),
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
			{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
	// todo handle response into result
}

func (g *gui) ShowQWidgetTab(qWidget *uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/ShowQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
	// todo handle response into result
}

func (g *gui) CloseQWidgetTab(qWidget *uintptr) {
	Client.Post().Url("http://localhost:8888/_scriptapi_gui.h/CloseQWidgetTab").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "qWidget",
				Type:  "*uintptr ",
				Value: fmt.Sprintf("%v", qWidget),
			},
		},
	))).Request()
	// todo handle response into result
}
