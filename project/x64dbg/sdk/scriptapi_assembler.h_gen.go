package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/mylog"
)

type assembler struct{}

func (a *assembler) Assemble(addr uint, dest *byte, size *int, instruction string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/Assemble").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "dest",
				Type:  "*byte ",
				Value: dest,
			},
			Param{
				Name:  "size",
				Type:  "*int ",
				Value: size,
			},
			Param{
				Name:  "instruction",
				Type:  "string",
				Value: instruction,
			},
		},
	))).Request()
	// todo handle response into result
}
func (a *assembler) AssembleEx(addr uint, dest *byte, size *int, instruction string, error *int8) {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/AssembleEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "dest",
				Type:  "*byte ",
				Value: dest,
			},
			Param{
				Name:  "size",
				Type:  "*int ",
				Value: size,
			},
			Param{
				Name:  "instruction",
				Type:  "string",
				Value: instruction,
			},
			Param{
				Name:  "error",
				Type:  "*int8 ",
				Value: error,
			},
		},
	))).Request()
	// todo handle response into result
}
func (a *assembler) AssembleMem(addr uint, instruction string) {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/AssembleMem").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "instruction",
				Type:  "string",
				Value: instruction,
			},
		},
	))).Request()
	// todo handle response into result
}
func (a *assembler) AssembleMemEx(addr uint, instruction string, size *int, error *int8, fillnop bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_assembler.h/AssembleMemEx").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "instruction",
				Type:  "string",
				Value: instruction,
			},
			Param{
				Name:  "size",
				Type:  "*int ",
				Value: size,
			},
			Param{
				Name:  "error",
				Type:  "*int8 ",
				Value: error,
			},
			Param{
				Name:  "fillnop",
				Type:  "bool",
				Value: fillnop,
			},
		},
	))).Request()
	// todo handle response into result
}
