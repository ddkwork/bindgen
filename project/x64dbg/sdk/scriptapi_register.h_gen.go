package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type RegisterEnum int //:10
const (
	DR0 RegisterEnum = iota
	DR11
	DR22
	DR33
	DR64
	DR75
	EAX6
	AX7
	AH8
	AL9
	EBX10
	BX11
	BH12
	BL13
	ECX14
	CX15
	CH16
	CL17
	EDX18
	DX19
	DH20
	DL21
	EDI22
	DI23
	ESI24
	SI25
	EBP26
	BP27
	ESP28
	SP29
	EIP30
	RAX31
	RBX32
	RCX33
	RDX34
	RSI35
	SIL36
	RDI37
	DIL38
	RBP39
	BPL40
	RSP41
	SPL42
	RIP43
	R844
	R8D45
	R8W46
	R8B47
	R948
	R9D49
	R9W50
	R9B51
	R1052
	R10D53
	R10W54
	R10B55
	R1156
	R11D57
	R11W58
	R11B59
	R1260
	R12D61
	R12W62
	R12B63
	R1364
	R13D65
	R13W66
	R13B67
	R1468
	R14D69
	R14W70
	R14B71
	R1572
	R15D73
	R15W74
	R15B75
	CIP76
	CSP77
	CAX78
	CBX79
	CCX80
	CDX81
	CDI82
	CSI83
	CBP84
	CFLAGS85
)

type register struct{}

// Get    c api name: Script::Register::Get
// ┌────┬────────┬────────────────────────────────┬──────────────┐
// │ id │  name  │             c type             │   go type    │
// ├────┼────────┼────────────────────────────────┼──────────────┤
// │ 0  │ reg    │ Script::Register::RegisterEnum │ RegisterEnum │
// ├────┼────────┼────────────────────────────────┼──────────────┤
// │    │ return │ duint                          │ uint         │
// └────┴────────┴────────────────────────────────┴──────────────┘
func (r *register) Get(reg RegisterEnum) uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "reg",
				Type:  "RegisterEnum",
				Value: fmt.Sprintf("%v", reg),
			},
		},
	))).Request()
	panic("not support return type: duint")
}

// Set    c api name: Script::Register::Set
// ┌────┬────────┬────────────────────────────────┬──────────────┐
// │ id │  name  │             c type             │   go type    │
// ├────┼────────┼────────────────────────────────┼──────────────┤
// │ 0  │ reg    │ Script::Register::RegisterEnum │ RegisterEnum │
// ├────┼────────┼────────────────────────────────┼──────────────┤
// │ 1  │ value  │ Script::Register::duint        │ uint         │
// ├────┼────────┼────────────────────────────────┼──────────────┤
// │    │ return │ bool                           │ bool         │
// └────┴────────┴────────────────────────────────┴──────────────┘
func (r *register) Set(reg RegisterEnum, value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/Set").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "reg",
				Type:  "RegisterEnum",
				Value: fmt.Sprintf("%v", reg),
			},
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// Size    c api name: Script::Register::Size
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ int    │ int     │
// └────┴────────┴────────┴─────────┘
func (r *register) Size() int {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/Size").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	return 0
}

// GetDR0    c api name: Script::Register::GetDR0
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetDR0() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR0").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetDR0    c api name: Script::Register::SetDR0
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetDR0(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR0").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDR1    c api name: Script::Register::GetDR1
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetDR1() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR1").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetDR1    c api name: Script::Register::SetDR1
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetDR1(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR1").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDR2    c api name: Script::Register::GetDR2
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetDR2() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR2").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetDR2    c api name: Script::Register::SetDR2
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetDR2(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR2").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDR3    c api name: Script::Register::GetDR3
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetDR3() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR3").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetDR3    c api name: Script::Register::SetDR3
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetDR3(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR3").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDR6    c api name: Script::Register::GetDR6
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetDR6() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR6").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetDR6    c api name: Script::Register::SetDR6
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetDR6(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR6").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDR7    c api name: Script::Register::GetDR7
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetDR7() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR7").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetDR7    c api name: Script::Register::SetDR7
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetDR7(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR7").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetEAX    c api name: Script::Register::GetEAX
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetEAX() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetEAX    c api name: Script::Register::SetEAX
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetEAX(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetAX    c api name: Script::Register::GetAX
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetAX() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetAX    c api name: Script::Register::SetAX
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetAX(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetAH    c api name: Script::Register::GetAH
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetAH() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetAH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetAH    c api name: Script::Register::SetAH
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetAH(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetAH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetAL    c api name: Script::Register::GetAL
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetAL() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetAL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetAL    c api name: Script::Register::SetAL
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetAL(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetAL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetEBX    c api name: Script::Register::GetEBX
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetEBX() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetEBX    c api name: Script::Register::SetEBX
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetEBX(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetBX    c api name: Script::Register::GetBX
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetBX() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetBX    c api name: Script::Register::SetBX
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetBX(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetBH    c api name: Script::Register::GetBH
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetBH() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetBH    c api name: Script::Register::SetBH
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetBH(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetBL    c api name: Script::Register::GetBL
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetBL() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetBL    c api name: Script::Register::SetBL
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetBL(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetECX    c api name: Script::Register::GetECX
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetECX() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetECX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetECX    c api name: Script::Register::SetECX
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetECX(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetECX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCX    c api name: Script::Register::GetCX
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetCX() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetCX    c api name: Script::Register::SetCX
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetCX(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCH    c api name: Script::Register::GetCH
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetCH() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetCH    c api name: Script::Register::SetCH
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetCH(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCL    c api name: Script::Register::GetCL
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetCL() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetCL    c api name: Script::Register::SetCL
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetCL(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetEDX    c api name: Script::Register::GetEDX
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetEDX() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetEDX    c api name: Script::Register::SetEDX
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetEDX(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDX    c api name: Script::Register::GetDX
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetDX() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetDX    c api name: Script::Register::SetDX
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetDX(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDH    c api name: Script::Register::GetDH
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetDH() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetDH    c api name: Script::Register::SetDH
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetDH(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDL    c api name: Script::Register::GetDL
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetDL() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetDL    c api name: Script::Register::SetDL
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetDL(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetEDI    c api name: Script::Register::GetEDI
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetEDI() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetEDI    c api name: Script::Register::SetEDI
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetEDI(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDI    c api name: Script::Register::GetDI
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetDI() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetDI    c api name: Script::Register::SetDI
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetDI(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetESI    c api name: Script::Register::GetESI
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetESI() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetESI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetESI    c api name: Script::Register::SetESI
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetESI(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetESI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetSI    c api name: Script::Register::GetSI
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetSI() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetSI    c api name: Script::Register::SetSI
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetSI(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetEBP    c api name: Script::Register::GetEBP
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetEBP() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetEBP    c api name: Script::Register::SetEBP
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetEBP(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetBP    c api name: Script::Register::GetBP
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetBP() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetBP    c api name: Script::Register::SetBP
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetBP(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetESP    c api name: Script::Register::GetESP
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetESP() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetESP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetESP    c api name: Script::Register::SetESP
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetESP(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetESP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetSP    c api name: Script::Register::GetSP
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetSP() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetSP    c api name: Script::Register::SetSP
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetSP(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetEIP    c api name: Script::Register::GetEIP
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetEIP() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetEIP    c api name: Script::Register::SetEIP
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetEIP(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRAX    c api name: Script::Register::GetRAX
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRAX() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRAX    c api name: Script::Register::SetRAX
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRAX(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRBX    c api name: Script::Register::GetRBX
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRBX() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRBX    c api name: Script::Register::SetRBX
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRBX(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRCX    c api name: Script::Register::GetRCX
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRCX() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRCX    c api name: Script::Register::SetRCX
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRCX(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRDX    c api name: Script::Register::GetRDX
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRDX() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRDX    c api name: Script::Register::SetRDX
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRDX(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRSI    c api name: Script::Register::GetRSI
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRSI() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRSI    c api name: Script::Register::SetRSI
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRSI(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetSIL    c api name: Script::Register::GetSIL
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetSIL() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetSIL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetSIL    c api name: Script::Register::SetSIL
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetSIL(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetSIL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRDI    c api name: Script::Register::GetRDI
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRDI() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRDI    c api name: Script::Register::SetRDI
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRDI(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetDIL    c api name: Script::Register::GetDIL
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetDIL() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDIL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetDIL    c api name: Script::Register::SetDIL
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetDIL(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDIL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRBP    c api name: Script::Register::GetRBP
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRBP() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRBP    c api name: Script::Register::SetRBP
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRBP(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetBPL    c api name: Script::Register::GetBPL
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetBPL() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBPL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetBPL    c api name: Script::Register::SetBPL
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetBPL(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBPL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRSP    c api name: Script::Register::GetRSP
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRSP() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRSP    c api name: Script::Register::SetRSP
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRSP(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetSPL    c api name: Script::Register::GetSPL
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetSPL() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetSPL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetSPL    c api name: Script::Register::SetSPL
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetSPL(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetSPL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetRIP    c api name: Script::Register::GetRIP
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetRIP() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetRIP    c api name: Script::Register::SetRIP
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetRIP(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR8    c api name: Script::Register::GetR8
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetR8() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR8").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetR8    c api name: Script::Register::SetR8
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetR8(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR8").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR8D    c api name: Script::Register::GetR8D
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetR8D() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR8D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetR8D    c api name: Script::Register::SetR8D
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetR8D(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR8D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR8W    c api name: Script::Register::GetR8W
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetR8W() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR8W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetR8W    c api name: Script::Register::SetR8W
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetR8W(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR8W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR8B    c api name: Script::Register::GetR8B
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetR8B() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR8B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetR8B    c api name: Script::Register::SetR8B
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetR8B(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR8B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR9    c api name: Script::Register::GetR9
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetR9() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR9").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetR9    c api name: Script::Register::SetR9
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetR9(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR9").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR9D    c api name: Script::Register::GetR9D
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetR9D() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR9D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetR9D    c api name: Script::Register::SetR9D
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetR9D(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR9D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR9W    c api name: Script::Register::GetR9W
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetR9W() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR9W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetR9W    c api name: Script::Register::SetR9W
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetR9W(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR9W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR9B    c api name: Script::Register::GetR9B
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetR9B() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR9B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetR9B    c api name: Script::Register::SetR9B
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetR9B(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR9B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR10    c api name: Script::Register::GetR10
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetR10() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR10").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetR10    c api name: Script::Register::SetR10
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetR10(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR10").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR10D    c api name: Script::Register::GetR10D
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetR10D() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR10D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetR10D    c api name: Script::Register::SetR10D
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetR10D(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR10D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR10W    c api name: Script::Register::GetR10W
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetR10W() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR10W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetR10W    c api name: Script::Register::SetR10W
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetR10W(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR10W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR10B    c api name: Script::Register::GetR10B
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetR10B() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR10B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetR10B    c api name: Script::Register::SetR10B
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetR10B(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR10B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR11    c api name: Script::Register::GetR11
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetR11() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR11").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetR11    c api name: Script::Register::SetR11
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetR11(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR11").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR11D    c api name: Script::Register::GetR11D
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetR11D() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR11D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetR11D    c api name: Script::Register::SetR11D
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetR11D(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR11D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR11W    c api name: Script::Register::GetR11W
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetR11W() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR11W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetR11W    c api name: Script::Register::SetR11W
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetR11W(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR11W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR11B    c api name: Script::Register::GetR11B
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetR11B() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR11B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetR11B    c api name: Script::Register::SetR11B
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetR11B(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR11B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR12    c api name: Script::Register::GetR12
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetR12() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR12").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetR12    c api name: Script::Register::SetR12
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetR12(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR12").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR12D    c api name: Script::Register::GetR12D
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetR12D() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR12D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetR12D    c api name: Script::Register::SetR12D
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetR12D(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR12D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR12W    c api name: Script::Register::GetR12W
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetR12W() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR12W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetR12W    c api name: Script::Register::SetR12W
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetR12W(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR12W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR12B    c api name: Script::Register::GetR12B
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetR12B() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR12B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetR12B    c api name: Script::Register::SetR12B
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetR12B(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR12B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR13    c api name: Script::Register::GetR13
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetR13() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR13").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetR13    c api name: Script::Register::SetR13
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetR13(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR13").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR13D    c api name: Script::Register::GetR13D
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetR13D() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR13D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetR13D    c api name: Script::Register::SetR13D
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetR13D(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR13D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR13W    c api name: Script::Register::GetR13W
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetR13W() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR13W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetR13W    c api name: Script::Register::SetR13W
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetR13W(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR13W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR13B    c api name: Script::Register::GetR13B
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetR13B() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR13B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetR13B    c api name: Script::Register::SetR13B
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetR13B(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR13B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR14    c api name: Script::Register::GetR14
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetR14() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR14").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetR14    c api name: Script::Register::SetR14
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetR14(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR14").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR14D    c api name: Script::Register::GetR14D
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetR14D() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR14D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetR14D    c api name: Script::Register::SetR14D
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetR14D(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR14D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR14W    c api name: Script::Register::GetR14W
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetR14W() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR14W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetR14W    c api name: Script::Register::SetR14W
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetR14W(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR14W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR14B    c api name: Script::Register::GetR14B
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetR14B() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR14B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetR14B    c api name: Script::Register::SetR14B
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetR14B(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR14B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR15    c api name: Script::Register::GetR15
// ┌────┬────────┬────────────────────┬─────────┐
// │ id │  name  │       c type       │ go type │
// ├────┼────────┼────────────────────┼─────────┤
// │    │ return │ unsigned long long │ uint64  │
// └────┴────────┴────────────────────┴─────────┘
func (r *register) GetR15() uint64 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR15").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned long long")
}

// SetR15    c api name: Script::Register::SetR15
// ┌────┬────────┬──────────────────────────────────────┬─────────┐
// │ id │  name  │                c type                │ go type │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned long long │ uint64  │
// ├────┼────────┼──────────────────────────────────────┼─────────┤
// │    │ return │ bool                                 │ bool    │
// └────┴────────┴──────────────────────────────────────┴─────────┘
func (r *register) SetR15(value uint64) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR15").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR15D    c api name: Script::Register::GetR15D
// ┌────┬────────┬──────────────┬─────────┐
// │ id │  name  │    c type    │ go type │
// ├────┼────────┼──────────────┼─────────┤
// │    │ return │ unsigned int │ uint    │
// └────┴────────┴──────────────┴─────────┘
func (r *register) GetR15D() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR15D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned int")
}

// SetR15D    c api name: Script::Register::SetR15D
// ┌────┬────────┬────────────────────────────────┬─────────┐
// │ id │  name  │             c type             │ go type │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned int │ uint    │
// ├────┼────────┼────────────────────────────────┼─────────┤
// │    │ return │ bool                           │ bool    │
// └────┴────────┴────────────────────────────────┴─────────┘
func (r *register) SetR15D(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR15D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR15W    c api name: Script::Register::GetR15W
// ┌────┬────────┬────────────────┬─────────┐
// │ id │  name  │     c type     │ go type │
// ├────┼────────┼────────────────┼─────────┤
// │    │ return │ unsigned short │ uint16  │
// └────┴────────┴────────────────┴─────────┘
func (r *register) GetR15W() uint16 {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR15W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned short")
}

// SetR15W    c api name: Script::Register::SetR15W
// ┌────┬────────┬──────────────────────────────────┬─────────┐
// │ id │  name  │              c type              │ go type │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned short │ uint16  │
// ├────┼────────┼──────────────────────────────────┼─────────┤
// │    │ return │ bool                             │ bool    │
// └────┴────────┴──────────────────────────────────┴─────────┘
func (r *register) SetR15W(value uint16) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR15W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetR15B    c api name: Script::Register::GetR15B
// ┌────┬────────┬───────────────┬─────────┐
// │ id │  name  │    c type     │ go type │
// ├────┼────────┼───────────────┼─────────┤
// │    │ return │ unsigned char │ byte    │
// └────┴────────┴───────────────┴─────────┘
func (r *register) GetR15B() byte {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR15B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: unsigned char")
}

// SetR15B    c api name: Script::Register::SetR15B
// ┌────┬────────┬─────────────────────────────────┬─────────┐
// │ id │  name  │             c type              │ go type │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::unsigned char │ byte    │
// ├────┼────────┼─────────────────────────────────┼─────────┤
// │    │ return │ bool                            │ bool    │
// └────┴────────┴─────────────────────────────────┴─────────┘
func (r *register) SetR15B(value byte) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR15B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCAX    c api name: Script::Register::GetCAX
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCAX() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCAX    c api name: Script::Register::SetCAX
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCAX(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCBX    c api name: Script::Register::GetCBX
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCBX() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCBX    c api name: Script::Register::SetCBX
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCBX(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCCX    c api name: Script::Register::GetCCX
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCCX() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCCX    c api name: Script::Register::SetCCX
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCCX(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCDX    c api name: Script::Register::GetCDX
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCDX() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCDX    c api name: Script::Register::SetCDX
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCDX(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCDI    c api name: Script::Register::GetCDI
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCDI() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCDI    c api name: Script::Register::SetCDI
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCDI(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCSI    c api name: Script::Register::GetCSI
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCSI() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCSI    c api name: Script::Register::SetCSI
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCSI(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCBP    c api name: Script::Register::GetCBP
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCBP() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCBP    c api name: Script::Register::SetCBP
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCBP(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCSP    c api name: Script::Register::GetCSP
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCSP() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCSP    c api name: Script::Register::SetCSP
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCSP(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCIP    c api name: Script::Register::GetCIP
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCIP() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCIP    c api name: Script::Register::SetCIP
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCIP(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}

// GetCFLAGS    c api name: Script::Register::GetCFLAGS
// ┌────┬────────┬────────┬─────────┐
// │ id │  name  │ c type │ go type │
// ├────┼────────┼────────┼─────────┤
// │    │ return │ duint  │ uint    │
// └────┴────────┴────────┴─────────┘
func (r *register) GetCFLAGS() uint {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCFLAGS").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	panic("not support return type: duint")
}

// SetCFLAGS    c api name: Script::Register::SetCFLAGS
// ┌────┬────────┬─────────────────────────┬─────────┐
// │ id │  name  │         c type          │ go type │
// ├────┼────────┼─────────────────────────┼─────────┤
// │ 0  │ value  │ Script::Register::duint │ uint    │
// ├────┼────────┼─────────────────────────┼─────────┤
// │    │ return │ bool                    │ bool    │
// └────┴────────┴─────────────────────────┴─────────┘
func (r *register) SetCFLAGS(value uint) bool {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCFLAGS").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	return true
}
