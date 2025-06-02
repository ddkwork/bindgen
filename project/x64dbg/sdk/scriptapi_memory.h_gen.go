package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type memory struct{}

// Read    c api name: Script::Memory::Read
// ┌────┬──────────┬─────────────────────────┬─────────┐
// │ id │   name   │         c type          │ go type │
// ├────┼──────────┼─────────────────────────┼─────────┤
// │ 0  │ addr     │ Script::Memory::duint   │ uint    │
// ├────┼──────────┼─────────────────────────┼─────────┤
// │ 1  │ data     │ Script::Memory::void *  │ uintptr │
// ├────┼──────────┼─────────────────────────┼─────────┤
// │ 2  │ size     │ Script::Memory::duint   │ uint    │
// ├────┼──────────┼─────────────────────────┼─────────┤
// │ 3  │ sizeRead │ Script::Memory::duint * │ *uint   │
// ├────┼──────────┼─────────────────────────┼─────────┤
// │    │ return   │ bool                    │ bool    │
// └────┴──────────┴─────────────────────────┴─────────┘
func (m *memory) Read(addr uint, data uintptr, size uint, sizeRead *uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/Read").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "data",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", data),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "sizeRead",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", sizeRead),
			},
		},
	))).Request()
	return true
}

// Write    c api name: Script::Memory::Write
// ┌────┬─────────────┬──────────────────────────────┬─────────┐
// │ id │    name     │            c type            │ go type │
// ├────┼─────────────┼──────────────────────────────┼─────────┤
// │ 0  │ addr        │ Script::Memory::duint        │ uint    │
// ├────┼─────────────┼──────────────────────────────┼─────────┤
// │ 1  │ data        │ const Script::Memory::void * │ uintptr │
// ├────┼─────────────┼──────────────────────────────┼─────────┤
// │ 2  │ size        │ Script::Memory::duint        │ uint    │
// ├────┼─────────────┼──────────────────────────────┼─────────┤
// │ 3  │ sizeWritten │ Script::Memory::duint *      │ *uint   │
// ├────┼─────────────┼──────────────────────────────┼─────────┤
// │    │ return      │ bool                         │ bool    │
// └────┴─────────────┴──────────────────────────────┴─────────┘
func (m *memory) Write(addr uint, data uintptr, size uint, sizeWritten *uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/Write").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "data",
				Type:  "uintptr",
				Value: fmt.Sprintf("%v", data),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
			{
				Name:  "sizeWritten",
				Type:  "*uint ",
				Value: fmt.Sprintf("%v", sizeWritten),
			},
		},
	))).Request()
	return true
}

// IsValidPtr    c api name: Script::Memory::IsValidPtr
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ bool                  │ bool    │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) IsValidPtr(addr uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/IsValidPtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// RemoteAlloc    c api name: Script::Memory::RemoteAlloc
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 1  │ size   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ duint                 │ uint    │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) RemoteAlloc(addr uint, size uint) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/RemoteAlloc").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// RemoteFree    c api name: Script::Memory::RemoteFree
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ bool                  │ bool    │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) RemoteFree(addr uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/RemoteFree").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	return true
}

// GetProtect    c api name: Script::Memory::GetProtect
// ┌────┬──────────┬───────────────────────┬─────────┐
// │ id │   name   │        c type         │ go type │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 0  │ addr     │ Script::Memory::duint │ uint    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 1  │ reserved │ Script::Memory::bool  │ bool    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 2  │ cache    │ Script::Memory::bool  │ bool    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │    │ return   │ unsigned int          │ uint    │
// └────┴──────────┴───────────────────────┴─────────┘
func (m *memory) GetProtect(addr uint, reserved bool, cache bool) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/GetProtect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "reserved",
				Type:  "bool",
				Value: fmt.Sprintf("%v", reserved),
			},
			{
				Name:  "cache",
				Type:  "bool",
				Value: fmt.Sprintf("%v", cache),
			},
		},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetProtect    c api name: Script::Memory::SetProtect
// ┌────┬─────────┬──────────────────────────────┬─────────┐
// │ id │  name   │            c type            │ go type │
// ├────┼─────────┼──────────────────────────────┼─────────┤
// │ 0  │ addr    │ Script::Memory::duint        │ uint    │
// ├────┼─────────┼──────────────────────────────┼─────────┤
// │ 1  │ protect │ Script::Memory::unsigned int │ uint    │
// ├────┼─────────┼──────────────────────────────┼─────────┤
// │ 2  │ size    │ Script::Memory::duint        │ uint    │
// ├────┼─────────┼──────────────────────────────┼─────────┤
// │    │ return  │ bool                         │ bool    │
// └────┴─────────┴──────────────────────────────┴─────────┘
func (m *memory) SetProtect(addr uint, protect uint, size uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/SetProtect").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "protect",
				Type:  "uint",
				Value: fmt.Sprintf("%v", protect),
			},
			{
				Name:  "size",
				Type:  "uint",
				Value: fmt.Sprintf("%v", size),
			},
		},
	))).Request()
	return true
}

// GetBase    c api name: Script::Memory::GetBase
// ┌────┬──────────┬───────────────────────┬─────────┐
// │ id │   name   │        c type         │ go type │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 0  │ addr     │ Script::Memory::duint │ uint    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 1  │ reserved │ Script::Memory::bool  │ bool    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 2  │ cache    │ Script::Memory::bool  │ bool    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │    │ return   │ duint                 │ uint    │
// └────┴──────────┴───────────────────────┴─────────┘
func (m *memory) GetBase(addr uint, reserved bool, cache bool) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/GetBase").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "reserved",
				Type:  "bool",
				Value: fmt.Sprintf("%v", reserved),
			},
			{
				Name:  "cache",
				Type:  "bool",
				Value: fmt.Sprintf("%v", cache),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// GetSize    c api name: Script::Memory::GetSize
// ┌────┬──────────┬───────────────────────┬─────────┐
// │ id │   name   │        c type         │ go type │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 0  │ addr     │ Script::Memory::duint │ uint    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 1  │ reserved │ Script::Memory::bool  │ bool    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │ 2  │ cache    │ Script::Memory::bool  │ bool    │
// ├────┼──────────┼───────────────────────┼─────────┤
// │    │ return   │ duint                 │ uint    │
// └────┴──────────┴───────────────────────┴─────────┘
func (m *memory) GetSize(addr uint, reserved bool, cache bool) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/GetSize").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "reserved",
				Type:  "bool",
				Value: fmt.Sprintf("%v", reserved),
			},
			{
				Name:  "cache",
				Type:  "bool",
				Value: fmt.Sprintf("%v", cache),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// ReadByte    c api name: Script::Memory::ReadByte
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ unsigned char         │ byte    │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) ReadByte(addr uint) byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadByte").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: unsigned char")
}

// WriteByte    c api name: Script::Memory::WriteByte
// ┌────┬────────┬───────────────────────────────┬─────────┐
// │ id │  name  │            c type             │ go type │
// ├────┼────────┼───────────────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint         │ uint    │
// ├────┼────────┼───────────────────────────────┼─────────┤
// │ 1  │ data   │ Script::Memory::unsigned char │ byte    │
// ├────┼────────┼───────────────────────────────┼─────────┤
// │    │ return │ bool                          │ bool    │
// └────┴────────┴───────────────────────────────┴─────────┘
func (m *memory) WriteByte(addr uint, data byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WriteByte").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "data",
				Type:  "byte",
				Value: fmt.Sprintf("%v", data),
			},
		},
	))).Request()
	return true
}

// ReadWord    c api name: Script::Memory::ReadWord
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ unsigned short        │ uint16  │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) ReadWord(addr uint) uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadWord").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: unsigned short")
}

// WriteWord    c api name: Script::Memory::WriteWord
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint          │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 1  │ data   │ Script::Memory::unsigned short │ uint16  │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (m *memory) WriteWord(addr uint, data uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WriteWord").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "data",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", data),
			},
		},
	))).Request()
	return true
}

// ReadDword    c api name: Script::Memory::ReadDword
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ unsigned int          │ uint    │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) ReadDword(addr uint) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadDword").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: unsigned int")
}

// WriteDword    c api name: Script::Memory::WriteDword
// ┌────┬────────┬──────────────────────────────┬─────────┐
// │ id │  name  │            c type            │ go type │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint        │ uint    │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │ 1  │ data   │ Script::Memory::unsigned int │ uint    │
// ├────┼────────┼──────────────────────────────┼─────────┤
// │    │ return │ bool                         │ bool    │
// └────┴────────┴──────────────────────────────┴─────────┘
func (m *memory) WriteDword(addr uint, data uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WriteDword").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "data",
				Type:  "uint",
				Value: fmt.Sprintf("%v", data),
			},
		},
	))).Request()
	return true
}

// ReadQword    c api name: Script::Memory::ReadQword
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ unsigned long long    │ uint64  │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) ReadQword(addr uint) uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadQword").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: unsigned long long")
}

// WriteQword    c api name: Script::Memory::WriteQword
// ┌────┬────────┬────────────────────────────────────┬─────────┐
// │ id │  name  │               c type               │ go type │
// ├────┼────────┼────────────────────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint              │ uint    │
// ├────┼────────┼────────────────────────────────────┼─────────┤
// │ 1  │ data   │ Script::Memory::unsigned long long │ uint64  │
// ├────┼────────┼────────────────────────────────────┼─────────┤
// │    │ return │ bool                               │ bool    │
// └────┴────────┴────────────────────────────────────┴─────────┘
func (m *memory) WriteQword(addr uint, data uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WriteQword").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "data",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", data),
			},
		},
	))).Request()
	return true
}

// ReadPtr    c api name: Script::Memory::ReadPtr
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ duint                 │ uint    │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) ReadPtr(addr uint) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/ReadPtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// WritePtr    c api name: Script::Memory::WritePtr
// ┌────┬────────┬───────────────────────┬─────────┐
// │ id │  name  │        c type         │ go type │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 0  │ addr   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │ 1  │ data   │ Script::Memory::duint │ uint    │
// ├────┼────────┼───────────────────────┼─────────┤
// │    │ return │ bool                  │ bool    │
// └────┴────────┴───────────────────────┴─────────┘
func (m *memory) WritePtr(addr uint, data uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_memory.h/WritePtr").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "addr",
				Type:  "uint",
				Value: fmt.Sprintf("%v", addr),
			},
			{
				Name:  "data",
				Type:  "uint",
				Value: fmt.Sprintf("%v", data),
			},
		},
	))).Request()
	return true
}
