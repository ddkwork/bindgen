package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type assembler struct{}

func (a *assembler) Assemble(addr uint, dest *byte, size *int, instruction string) {
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
				Type:  "string",
				Value: fmt.Sprintf("%v", instruction),
			},
		},
	))).Request()
	// todo handle response into result
}

func (a *assembler) AssembleEx(addr uint, dest *byte, size *int, instruction string, error *int8) {
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
				Type:  "string",
				Value: fmt.Sprintf("%v", instruction),
			},
			{
				Name:  "error",
				Type:  "*int8 ",
				Value: fmt.Sprintf("%v", error),
			},
		},
	))).Request()
	// todo handle response into result
}

func (a *assembler) AssembleMem(addr uint, instruction string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/AssembleMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "instruction",
				Type:  "string",
				Value: fmt.Sprintf("%v", instruction),
			},
		},
	))).Request()
	// todo handle response into result
}

func (a *assembler) AssembleMemEx(addr uint, instruction string, size *int, error *int8, fillnop bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/AssembleMemEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "instruction",
				Type:  "string",
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
	// todo handle response into result
}
