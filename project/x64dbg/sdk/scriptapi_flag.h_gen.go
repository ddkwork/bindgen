package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/mylog"
)

type FlagEnum int // :10
const (
	ZF FlagEnum = iota // 0
	OF                 // 1
	CF                 // 2
	PF                 // 3
	SF                 // 4
	TF                 // 5
	AF                 // 6
	DF                 // 7
	IF                 // 8
)

type flag struct{}

func (f *flag) Get(flag FlagEnum) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "flag",
				Type:  "FlagEnum",
				Value: flag,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) Set(flag FlagEnum, value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "flag",
				Type:  "FlagEnum",
				Value: flag,
			},
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetZF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetZF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetZF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetZF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetOF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetOF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetOF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetOF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetCF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetCF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetCF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetCF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetPF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetPF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetPF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetPF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetSF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetSF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetSF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetSF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetTF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetTF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetTF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetTF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetAF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetAF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetAF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetAF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetDF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetDF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetDF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetDF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
func (f *flag) GetIF() {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/GetIF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}
func (f *flag) SetIF(value bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_flag.h/SetIF").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "value",
				Type:  "bool",
				Value: value,
			},
		},
	))).Request()
	// todo handle response into result
}
