package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

// CommentInfo (:10)
type CommentInfo struct {
	Mod    [256]int8 // C type: char[256]
	Rva    uint      // C type: duint
	Text   [256]int8 // C type: char[256]
	Manual bool      // C type: bool
}
type comment struct{}

func (c *comment) Set(info *CommentInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "info",
				Type:  "*CommentInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (c *comment) Get(addr uint, text *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (c *comment) GetInfo(addr uint, info *CommentInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "info",
				Type:  "*CommentInfo ",
				Value: fmt.Sprintf("%v", info),
			},
		},
	))).Request()
	// todo handle response into result
}

func (c *comment) Delete(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (c *comment) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (c *comment) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (c *comment) GetList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
