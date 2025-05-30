package sdk

import (
	"encoding/json"
	"github.com/ddkwork/golibrary/mylog"
)

type memory struct{}

func (m *memory) Read(addr uint, data *uintptr, size uint, sizeRead *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/Read").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "data",
				Type:  "*uintptr ",
				Value: data,
			},
			Param{
				Name:  "size",
				Type:  "uint",
				Value: size,
			},
			Param{
				Name:  "sizeRead",
				Type:  "*uint ",
				Value: sizeRead,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) Write(addr uint, data *uintptr, size uint, sizeWritten *uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/Write").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "data",
				Type:  "*uintptr ",
				Value: data,
			},
			Param{
				Name:  "size",
				Type:  "uint",
				Value: size,
			},
			Param{
				Name:  "sizeWritten",
				Type:  "*uint ",
				Value: sizeWritten,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) IsValidPtr(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/IsValidPtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (m *memory) RemoteAlloc(addr uint, size uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/RemoteAlloc").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "size",
				Type:  "uint",
				Value: size,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) RemoteFree(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/RemoteFree").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (m *memory) GetProtect(addr uint, reserved bool, cache bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/GetProtect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "reserved",
				Type:  "bool",
				Value: reserved,
			},
			Param{
				Name:  "cache",
				Type:  "bool",
				Value: cache,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) SetProtect(addr uint, protect uint, size uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/SetProtect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "protect",
				Type:  "uint",
				Value: protect,
			},
			Param{
				Name:  "size",
				Type:  "uint",
				Value: size,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) GetBase(addr uint, reserved bool, cache bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/GetBase").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "reserved",
				Type:  "bool",
				Value: reserved,
			},
			Param{
				Name:  "cache",
				Type:  "bool",
				Value: cache,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) GetSize(addr uint, reserved bool, cache bool) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/GetSize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "reserved",
				Type:  "bool",
				Value: reserved,
			},
			Param{
				Name:  "cache",
				Type:  "bool",
				Value: cache,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) ReadByte(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadByte").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (m *memory) WriteByte(addr uint, data byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WriteByte").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "data",
				Type:  "byte",
				Value: data,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) ReadWord(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadWord").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (m *memory) WriteWord(addr uint, data uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WriteWord").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "data",
				Type:  "uint16",
				Value: data,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) ReadDword(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadDword").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (m *memory) WriteDword(addr uint, data uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WriteDword").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "data",
				Type:  "uint",
				Value: data,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) ReadQword(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadQword").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (m *memory) WriteQword(addr uint, data uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WriteQword").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "data",
				Type:  "uint64",
				Value: data,
			},
		},
	))).Request()
	// todo handle response into result
}
func (m *memory) ReadPtr(addr uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadPtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
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
func (m *memory) WritePtr(addr uint, data uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WritePtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			Param{
				Name:  "addr",
				Type:  "uint",
				Value: addr,
			},
			Param{
				Name:  "data",
				Type:  "uint",
				Value: data,
			},
		},
	))).Request()
	// todo handle response into result
}
