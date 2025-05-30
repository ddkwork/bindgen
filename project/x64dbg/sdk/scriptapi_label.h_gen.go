package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// LabelInfo (:10)
type LabelInfo struct {
	Mod    [256]int8 // C type: char[256]
	Rva    uint      // C type: duint
	Text   [256]int8 // C type: char[256]
	Manual bool      // C type: bool
}
type label struct{}

func (l *label) Set(info *LabelInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*LabelInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (l *label) FromString(label string, addr *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/FromString").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "label",
				Type:  "string",
				Value: fmt.Sprintf("%v", label),
			},
			{
				Name:  "addr",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	// todo handle response into result
}

func (l *label) Get(addr uint, text *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
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

func (l *label) IsTemporary(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/IsTemporary").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (l *label) GetInfo(addr uint, info *LabelInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*LabelInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (l *label) Delete(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (l *label) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
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

func (l *label) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (l *label) GetList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_label.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
