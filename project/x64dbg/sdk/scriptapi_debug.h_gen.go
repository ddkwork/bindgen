package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type HardwareType int //:10
const (
	HardwareAccess HardwareType = iota
	HardwareWrite1
	HardwareExecute2
)

type debug struct{}

// Wait    c api name: Script::Debug::Wait
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (d *debug) Wait() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/Wait").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// Run    c api name: Script::Debug::Run
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (d *debug) Run() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/Run").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// Pause    c api name: Script::Debug::Pause
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (d *debug) Pause() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/Pause").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// Stop    c api name: Script::Debug::Stop
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (d *debug) Stop() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/Stop").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// StepIn    c api name: Script::Debug::StepIn
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (d *debug) StepIn() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/StepIn").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// StepOver    c api name: Script::Debug::StepOver
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (d *debug) StepOver() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/StepOver").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// StepOut    c api name: Script::Debug::StepOut
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ void   │         │
// └────┴────────┴────────┴─────────┘
func (d *debug) StepOut() {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/StepOut").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
}

// SetBreakpoint    c api name: Script::Debug::SetBreakpoint
// ┌────┬─────────┬──────────────────────┬─────────┐
// │ id │  name   │        c type        │ go type │
// ├────┼─────────┼──────────────────────┼─────────┤
// │ 0  │ address │ Script::Debug::duint │ uint    │
// ├────┼─────────┼──────────────────────┼─────────┤
// │    │ return  │ bool                 │ bool    │
// └────┴─────────┴──────────────────────┴─────────┘
func (d *debug) SetBreakpoint(address uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/SetBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "address",
				Type:  "uint",
				Value: fmt.Sprintf("%v", address),
			},
		},
	))).Request()
	return true
}

// DeleteBreakpoint    c api name: Script::Debug::DeleteBreakpoint
// ┌────┬─────────┬──────────────────────┬─────────┐
// │ id │  name   │        c type        │ go type │
// ├────┼─────────┼──────────────────────┼─────────┤
// │ 0  │ address │ Script::Debug::duint │ uint    │
// ├────┼─────────┼──────────────────────┼─────────┤
// │    │ return  │ bool                 │ bool    │
// └────┴─────────┴──────────────────────┴─────────┘
func (d *debug) DeleteBreakpoint(address uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/DeleteBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "address",
				Type:  "uint",
				Value: fmt.Sprintf("%v", address),
			},
		},
	))).Request()
	return true
}

// DisableBreakpoint    c api name: Script::Debug::DisableBreakpoint
// ┌────┬─────────┬──────────────────────┬─────────┐
// │ id │  name   │        c type        │ go type │
// ├────┼─────────┼──────────────────────┼─────────┤
// │ 0  │ address │ Script::Debug::duint │ uint    │
// ├────┼─────────┼──────────────────────┼─────────┤
// │    │ return  │ bool                 │ bool    │
// └────┴─────────┴──────────────────────┴─────────┘
func (d *debug) DisableBreakpoint(address uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/DisableBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "address",
				Type:  "uint",
				Value: fmt.Sprintf("%v", address),
			},
		},
	))).Request()
	return true
}

// SetHardwareBreakpoint    c api name: Script::Debug::SetHardwareBreakpoint
// ┌────┬─────────┬─────────────────────────────┬──────────────┐
// │ id │  name   │           c type            │   go type    │
// ├────┼─────────┼─────────────────────────────┼──────────────┤
// │ 0  │ address │ Script::Debug::duint        │ uint         │
// ├────┼─────────┼─────────────────────────────┼──────────────┤
// │ 1  │ Type    │ Script::Debug::HardwareType │ HardwareType │
// ├────┼─────────┼─────────────────────────────┼──────────────┤
// │    │ return  │ bool                        │ bool         │
// └────┴─────────┴─────────────────────────────┴──────────────┘
func (d *debug) SetHardwareBreakpoint(address uint, Type HardwareType) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/SetHardwareBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "address",
				Type:  "uint",
				Value: fmt.Sprintf("%v", address),
			},
			{
				Name:  "Type",
				Type:  "HardwareType",
				Value: fmt.Sprintf("%v", Type),
			},
		},
	))).Request()
	return true
}

// DeleteHardwareBreakpoint    c api name: Script::Debug::DeleteHardwareBreakpoint
// ┌────┬─────────┬──────────────────────┬─────────┐
// │ id │  name   │        c type        │ go type │
// ├────┼─────────┼──────────────────────┼─────────┤
// │ 0  │ address │ Script::Debug::duint │ uint    │
// ├────┼─────────┼──────────────────────┼─────────┤
// │    │ return  │ bool                 │ bool    │
// └────┴─────────┴──────────────────────┴─────────┘
func (d *debug) DeleteHardwareBreakpoint(address uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_debug.h/DeleteHardwareBreakpoint").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "address",
				Type:  "uint",
				Value: fmt.Sprintf("%v", address),
			},
		},
	))).Request()
	return true
}
