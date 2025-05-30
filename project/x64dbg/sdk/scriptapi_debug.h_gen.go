package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/mylog"
)

type HardwareType int // :10
const (
	HardwareAccess  HardwareType = iota // 0
	HardwareWrite                       // 1
	HardwareExecute                     // 2
)

type debug struct{}

func (d *debug) Wait() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/Wait").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (d *debug) Run() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/Run").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (d *debug) Pause() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/Pause").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (d *debug) Stop() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/Stop").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (d *debug) StepIn() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/StepIn").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (d *debug) StepOver() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/StepOver").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (d *debug) StepOut() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/StepOut").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (d *debug) SetBreakpoint(address uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/SetBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (d *debug) DeleteBreakpoint(address uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/DeleteBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (d *debug) DisableBreakpoint(address uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/DisableBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
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

func (d *debug) SetHardwareBreakpoint(address uint, Type HardwareType) {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/SetHardwareBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "address",
				Type:  "uint",
				Value: address,
			},
			Param{
				Name:  "Type",
				Type:  "HardwareType",
				Value: Type,
			},
		},
	))).Request()
	// todo handle response into result
}

func (d *debug) DeleteHardwareBreakpoint(address uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/DeleteHardwareBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
