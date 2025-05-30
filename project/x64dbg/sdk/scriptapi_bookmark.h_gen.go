package sdk

// BookmarkInfo (:10)
type BookmarkInfo struct {
	Mod    [256]int8 // C type: char[256]
	Rva    uint      // C type: duint
	Manual bool      // C type: bool
}
type bookmark struct{}

func (b *bookmark) Set(info *BookmarkInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "info",
				Type:  "*BookmarkInfo ",
				Value: info,
			},
		},
	))).Request()
	// todo handle response into result
}
func (b *bookmark) Get(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (b *bookmark) GetInfo(addr uint, info *BookmarkInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "info",
				Type:  "*BookmarkInfo ",
				Value: info,
			},
		},
	))).Request()
	// todo handle response into result
}
func (b *bookmark) Delete(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (b *bookmark) DeleteRange(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (b *bookmark) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (b *bookmark) GetList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_bookmark.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
