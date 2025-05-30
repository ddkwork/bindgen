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
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "info",
				Type:  "*CommentInfo ",
				Value: info,
			},
		},
	))).Request()
	// todo handle response into result
}
func (c *comment) Get(addr uint, text *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
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
func (c *comment) GetInfo(addr uint, info *CommentInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "info",
				Type:  "*CommentInfo ",
				Value: info,
			},
		},
	))).Request()
	// todo handle response into result
}
func (c *comment) Delete(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
		},
	))).Request()
	// todo handle response into result
}
func (c *comment) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
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
func (c *comment) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (c *comment) GetList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_comment.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "list",
				Type:  "*ListInfo ",
				Value: list,
			},
		},
	))).Request()
	// todo handle response into result
}
