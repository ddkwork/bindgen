package sdk

import (
	"encoding/json"
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
)

type RegisterEnum int // :10
const (
	DR0    RegisterEnum = iota // 0
	DR1                        // 1
	DR2                        // 2
	DR3                        // 3
	DR6                        // 4
	DR7                        // 5
	EAX                        // 6
	AX                         // 7
	AH                         // 8
	AL                         // 9
	EBX                        // 10
	BX                         // 11
	BH                         // 12
	BL                         // 13
	ECX                        // 14
	CX                         // 15
	CH                         // 16
	CL                         // 17
	EDX                        // 18
	DX                         // 19
	DH                         // 20
	DL                         // 21
	EDI                        // 22
	DI                         // 23
	ESI                        // 24
	SI                         // 25
	EBP                        // 26
	BP                         // 27
	ESP                        // 28
	SP                         // 29
	EIP                        // 30
	RAX                        // 31
	RBX                        // 32
	RCX                        // 33
	RDX                        // 34
	RSI                        // 35
	SIL                        // 36
	RDI                        // 37
	DIL                        // 38
	RBP                        // 39
	BPL                        // 40
	RSP                        // 41
	SPL                        // 42
	RIP                        // 43
	R8                         // 44
	R8D                        // 45
	R8W                        // 46
	R8B                        // 47
	R9                         // 48
	R9D                        // 49
	R9W                        // 50
	R9B                        // 51
	R10                        // 52
	R10D                       // 53
	R10W                       // 54
	R10B                       // 55
	R11                        // 56
	R11D                       // 57
	R11W                       // 58
	R11B                       // 59
	R12                        // 60
	R12D                       // 61
	R12W                       // 62
	R12B                       // 63
	R13                        // 64
	R13D                       // 65
	R13W                       // 66
	R13B                       // 67
	R14                        // 68
	R14D                       // 69
	R14W                       // 70
	R14B                       // 71
	R15                        // 72
	R15D                       // 73
	R15W                       // 74
	R15B                       // 75
	CIP                        // 76
	CSP                        // 77
	CAX                        // 78
	CBX                        // 79
	CCX                        // 80
	CDX                        // 81
	CDI                        // 82
	CSI                        // 83
	CBP                        // 84
	CFLAGS                     // 85
)

type register struct{}

func (r *register) Get(reg RegisterEnum) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/Get").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "reg",
				Type:  "RegisterEnum",
				Value: fmt.Sprintf("%v", reg),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) Set(reg RegisterEnum, value uint) {
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
	// todo handle response into result
}

func (r *register) Size() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/Size").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDR0() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR0").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDR0(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR0").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDR1() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR1").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDR1(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR1").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDR2() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR2").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDR2(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR2").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDR3() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR3").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDR3(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR3").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDR6() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR6").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDR6(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR6").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDR7() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDR7").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDR7(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDR7").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetEAX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetEAX(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetAX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetAX(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetAH() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetAH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetAH(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetAH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetAL() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetAL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetAL(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetAL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetEBX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetEBX(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetBX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetBX(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetBH() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetBH(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetBL() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetBL(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetECX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetECX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetECX(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetECX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCX(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCH() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCH(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCL() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCL(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetEDX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetEDX(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDX(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDH() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDH(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDH").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDL() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDL(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetEDI() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetEDI(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDI() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDI(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetESI() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetESI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetESI(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetESI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetSI() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetSI(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetEBP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetEBP(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetBP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetBP(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetESP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetESP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetESP(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetESP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetSP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetSP(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetEIP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetEIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetEIP(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetEIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRAX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRAX(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRBX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRBX(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRCX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRCX(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRDX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRDX(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRSI() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRSI(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetSIL() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetSIL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetSIL(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetSIL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRDI() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRDI(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetDIL() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetDIL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetDIL(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetDIL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRBP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRBP(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetBPL() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetBPL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetBPL(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetBPL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRSP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRSP(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetSPL() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetSPL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetSPL(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetSPL").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetRIP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetRIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetRIP(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetRIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR8() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR8").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR8(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR8").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR8D() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR8D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR8D(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR8D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR8W() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR8W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR8W(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR8W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR8B() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR8B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR8B(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR8B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR9() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR9").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR9(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR9").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR9D() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR9D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR9D(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR9D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR9W() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR9W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR9W(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR9W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR9B() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR9B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR9B(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR9B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR10() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR10").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR10(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR10").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR10D() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR10D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR10D(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR10D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR10W() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR10W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR10W(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR10W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR10B() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR10B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR10B(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR10B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR11() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR11").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR11(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR11").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR11D() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR11D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR11D(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR11D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR11W() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR11W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR11W(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR11W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR11B() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR11B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR11B(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR11B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR12() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR12").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR12(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR12").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR12D() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR12D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR12D(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR12D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR12W() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR12W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR12W(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR12W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR12B() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR12B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR12B(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR12B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR13() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR13").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR13(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR13").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR13D() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR13D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR13D(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR13D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR13W() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR13W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR13W(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR13W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR13B() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR13B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR13B(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR13B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR14() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR14").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR14(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR14").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR14D() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR14D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR14D(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR14D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR14W() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR14W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR14W(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR14W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR14B() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR14B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR14B(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR14B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR15() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR15").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR15(value uint64) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR15").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint64",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR15D() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR15D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR15D(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR15D").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR15W() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR15W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR15W(value uint16) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR15W").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint16",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetR15B() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetR15B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetR15B(value byte) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetR15B").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "byte",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCAX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCAX(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCAX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCBX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCBX(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCBX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCCX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCCX(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCCX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCDX() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCDX(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCDX").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCDI() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCDI(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCDI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCSI() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCSI(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCSI").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCBP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCBP(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCBP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCSP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCSP(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCSP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCIP() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCIP(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCIP").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}

func (r *register) GetCFLAGS() {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/GetCFLAGS").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{},
	))).Request()
	// todo handle response into result
}

func (r *register) SetCFLAGS(value uint) {
	Client.Post().Url("http://localhost:8888/_scriptapi_register.h/SetCFLAGS").SetJsonHead().Body(mylog.Check2(json.Marshal(
		[]Param{
			{
				Name:  "value",
				Type:  "uint",
				Value: fmt.Sprintf("%v", value),
			},
		},
	))).Request()
	// todo handle response into result
}
