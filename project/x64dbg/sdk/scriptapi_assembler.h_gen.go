package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type assembler struct{}

// Assemble    c api name: Script::Assembler::Assemble
// ┌────┬─────────────┬─────────────────┬─────────┐
// │ id │    name     │     c type      │ go type │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 0  │ addr        │ duint           │ uint    │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 1  │ dest        │ unsigned char * │ *byte   │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 2  │ size        │ int *           │ *int    │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 3  │ instruction │ const char *    │ *int8   │
// ├────┼─────────────┼─────────────────┼─────────┤
// │    │ return      │ bool            │ bool    │
// └────┴─────────────┴─────────────────┴─────────┘
func (a *assembler) Assemble(addr uint, dest *byte, size *int, instruction *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/Assemble").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "dest",
				Type:  "*byte ",
				Value: fmt.Sprintf("%v", dest),
			},
			{
				Name:  "size",
				Type:  "*int ",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "instruction",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", instruction),
			},
		},
	))).Request()
	return true
}

// AssembleEx    c api name: Script::Assembler::AssembleEx
// ┌────┬─────────────┬─────────────────┬─────────┐
// │ id │    name     │     c type      │ go type │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 0  │ addr        │ duint           │ uint    │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 1  │ dest        │ unsigned char * │ *byte   │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 2  │ size        │ int *           │ *int    │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 3  │ instruction │ const char *    │ *int8   │
// ├────┼─────────────┼─────────────────┼─────────┤
// │ 4  │ error       │ char *          │ *int8   │
// ├────┼─────────────┼─────────────────┼─────────┤
// │    │ return      │ bool            │ bool    │
// └────┴─────────────┴─────────────────┴─────────┘
func (a *assembler) AssembleEx(addr uint, dest *byte, size *int, instruction *int8, error *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/AssembleEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "dest",
				Type:  "*byte ",
				Value: fmt.Sprintf("%v", dest),
			},
			{
				Name:  "size",
				Type:  "*int ",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "instruction",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", instruction),
			},
			{
				Name:  "error",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", error),
			},
		},
	))).Request()
	return true
}

// AssembleMem    c api name: Script::Assembler::AssembleMem
// ┌────┬─────────────┬──────────────┬─────────┐
// │ id │    name     │    c type    │ go type │
// ├────┼─────────────┼──────────────┼─────────┤
// │ 0  │ addr        │ duint        │ uint    │
// ├────┼─────────────┼──────────────┼─────────┤
// │ 1  │ instruction │ const char * │ *int8   │
// ├────┼─────────────┼──────────────┼─────────┤
// │    │ return      │ bool         │ bool    │
// └────┴─────────────┴──────────────┴─────────┘
func (a *assembler) AssembleMem(addr uint, instruction *int8) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/AssembleMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "instruction",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", instruction),
			},
		},
	))).Request()
	return true
}

// AssembleMemEx    c api name: Script::Assembler::AssembleMemEx
// ┌────┬─────────────┬──────────────┬─────────┐
// │ id │    name     │    c type    │ go type │
// ├────┼─────────────┼──────────────┼─────────┤
// │ 0  │ addr        │ duint        │ uint    │
// ├────┼─────────────┼──────────────┼─────────┤
// │ 1  │ instruction │ const char * │ *int8   │
// ├────┼─────────────┼──────────────┼─────────┤
// │ 2  │ size        │ int *        │ *int    │
// ├────┼─────────────┼──────────────┼─────────┤
// │ 3  │ error       │ char *       │ *int8   │
// ├────┼─────────────┼──────────────┼─────────┤
// │ 4  │ fillnop     │ bool         │ bool    │
// ├────┼─────────────┼──────────────┼─────────┤
// │    │ return      │ bool         │ bool    │
// └────┴─────────────┴──────────────┴─────────┘
func (a *assembler) AssembleMemEx(addr uint, instruction *int8, size *int, error *int8, fillnop bool) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/AssembleMemEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "instruction",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", instruction),
			},
			{
				Name:  "size",
				Type:  "*int ",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "error",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", error),
			},
			{
				Name:  "fillnop",
				Type:  "bool",
				Value: fmt.Sprintf("%v", fillnop),
			},
		},
	))).Request()
	return true
}
