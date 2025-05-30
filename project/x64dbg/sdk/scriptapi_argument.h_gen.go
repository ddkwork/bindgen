package sdk

// ArgumentInfo (:10)
type ArgumentInfo struct {
	Mod              [256]int8 // C type: char[256]
	RvaStart         uint      // C type: duint
	RvaEnd           uint      // C type: duint
	Manual           bool      // C type: bool
	Instructioncount uint      // C type: duint
}
type argument struct{}

func (a *argument) Add(info *ArgumentInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Add").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "info",
				Type:  "*ArgumentInfo ",
				Value: info,
			},
		},
	))).Request()
	// todo handle response into result
}
func (a *argument) Get(addr uint, start *uint, end *uint, instructionCount *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
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
			Param{
				Name:  "instructionCount",
				Type:  "*uint ",
				Value: instructionCount,
			},
		},
	))).Request()
	// todo handle response into result
}
func (a *argument) GetInfo(addr uint, info *ArgumentInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/GetInfo").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "info",
				Type:  "*ArgumentInfo ",
				Value: info,
			},
		},
	))).Request()
	// todo handle response into result
}
func (a *argument) Overlaps(start uint, end uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Overlaps").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (a *argument) Delete(address uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Delete").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "address",
				Type:  "uint",
				Value: address,
			},
		},
	))).Request()
	// todo handle response into result
}
func (a *argument) DeleteRange(start uint, end uint, deleteManual bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/DeleteRange").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
			Param{
				Name:  "deleteManual",
				Type:  "bool",
				Value: deleteManual,
			},
		},
	))).Request()
	// todo handle response into result
}
func (a *argument) Clear() {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/Clear").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (a *argument) GetList(list *ListInfo) {
	Client.Post().Url("http://localhost:8888/_scriptapi_argument.h/GetList").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
