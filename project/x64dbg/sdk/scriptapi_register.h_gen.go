package sdk

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
	response := safeGet("_scriptapi_register.h/Get", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) Set(reg RegisterEnum, value uint) {
	response := safeGet("_scriptapi_register.h/Set", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) Size() {
	response := safeGet("_scriptapi_register.h/Size", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDR0() {
	response := safeGet("_scriptapi_register.h/GetDR0", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDR0(value uint) {
	response := safeGet("_scriptapi_register.h/SetDR0", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDR1() {
	response := safeGet("_scriptapi_register.h/GetDR1", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDR1(value uint) {
	response := safeGet("_scriptapi_register.h/SetDR1", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDR2() {
	response := safeGet("_scriptapi_register.h/GetDR2", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDR2(value uint) {
	response := safeGet("_scriptapi_register.h/SetDR2", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDR3() {
	response := safeGet("_scriptapi_register.h/GetDR3", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDR3(value uint) {
	response := safeGet("_scriptapi_register.h/SetDR3", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDR6() {
	response := safeGet("_scriptapi_register.h/GetDR6", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDR6(value uint) {
	response := safeGet("_scriptapi_register.h/SetDR6", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDR7() {
	response := safeGet("_scriptapi_register.h/GetDR7", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDR7(value uint) {
	response := safeGet("_scriptapi_register.h/SetDR7", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetEAX() {
	response := safeGet("_scriptapi_register.h/GetEAX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetEAX(value uint) {
	response := safeGet("_scriptapi_register.h/SetEAX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetAX() {
	response := safeGet("_scriptapi_register.h/GetAX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetAX(value uint16) {
	response := safeGet("_scriptapi_register.h/SetAX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetAH() {
	response := safeGet("_scriptapi_register.h/GetAH", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetAH(value byte) {
	response := safeGet("_scriptapi_register.h/SetAH", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetAL() {
	response := safeGet("_scriptapi_register.h/GetAL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetAL(value byte) {
	response := safeGet("_scriptapi_register.h/SetAL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetEBX() {
	response := safeGet("_scriptapi_register.h/GetEBX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetEBX(value uint) {
	response := safeGet("_scriptapi_register.h/SetEBX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetBX() {
	response := safeGet("_scriptapi_register.h/GetBX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetBX(value uint16) {
	response := safeGet("_scriptapi_register.h/SetBX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetBH() {
	response := safeGet("_scriptapi_register.h/GetBH", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetBH(value byte) {
	response := safeGet("_scriptapi_register.h/SetBH", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetBL() {
	response := safeGet("_scriptapi_register.h/GetBL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetBL(value byte) {
	response := safeGet("_scriptapi_register.h/SetBL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetECX() {
	response := safeGet("_scriptapi_register.h/GetECX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetECX(value uint) {
	response := safeGet("_scriptapi_register.h/SetECX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCX() {
	response := safeGet("_scriptapi_register.h/GetCX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCX(value uint16) {
	response := safeGet("_scriptapi_register.h/SetCX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCH() {
	response := safeGet("_scriptapi_register.h/GetCH", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCH(value byte) {
	response := safeGet("_scriptapi_register.h/SetCH", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCL() {
	response := safeGet("_scriptapi_register.h/GetCL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCL(value byte) {
	response := safeGet("_scriptapi_register.h/SetCL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetEDX() {
	response := safeGet("_scriptapi_register.h/GetEDX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetEDX(value uint) {
	response := safeGet("_scriptapi_register.h/SetEDX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDX() {
	response := safeGet("_scriptapi_register.h/GetDX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDX(value uint16) {
	response := safeGet("_scriptapi_register.h/SetDX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDH() {
	response := safeGet("_scriptapi_register.h/GetDH", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDH(value byte) {
	response := safeGet("_scriptapi_register.h/SetDH", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDL() {
	response := safeGet("_scriptapi_register.h/GetDL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDL(value byte) {
	response := safeGet("_scriptapi_register.h/SetDL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetEDI() {
	response := safeGet("_scriptapi_register.h/GetEDI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetEDI(value uint) {
	response := safeGet("_scriptapi_register.h/SetEDI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDI() {
	response := safeGet("_scriptapi_register.h/GetDI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDI(value uint16) {
	response := safeGet("_scriptapi_register.h/SetDI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetESI() {
	response := safeGet("_scriptapi_register.h/GetESI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetESI(value uint) {
	response := safeGet("_scriptapi_register.h/SetESI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetSI() {
	response := safeGet("_scriptapi_register.h/GetSI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetSI(value uint16) {
	response := safeGet("_scriptapi_register.h/SetSI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetEBP() {
	response := safeGet("_scriptapi_register.h/GetEBP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetEBP(value uint) {
	response := safeGet("_scriptapi_register.h/SetEBP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetBP() {
	response := safeGet("_scriptapi_register.h/GetBP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetBP(value uint16) {
	response := safeGet("_scriptapi_register.h/SetBP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetESP() {
	response := safeGet("_scriptapi_register.h/GetESP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetESP(value uint) {
	response := safeGet("_scriptapi_register.h/SetESP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetSP() {
	response := safeGet("_scriptapi_register.h/GetSP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetSP(value uint16) {
	response := safeGet("_scriptapi_register.h/SetSP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetEIP() {
	response := safeGet("_scriptapi_register.h/GetEIP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetEIP(value uint) {
	response := safeGet("_scriptapi_register.h/SetEIP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRAX() {
	response := safeGet("_scriptapi_register.h/GetRAX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRAX(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRAX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRBX() {
	response := safeGet("_scriptapi_register.h/GetRBX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRBX(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRBX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRCX() {
	response := safeGet("_scriptapi_register.h/GetRCX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRCX(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRCX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRDX() {
	response := safeGet("_scriptapi_register.h/GetRDX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRDX(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRDX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRSI() {
	response := safeGet("_scriptapi_register.h/GetRSI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRSI(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRSI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetSIL() {
	response := safeGet("_scriptapi_register.h/GetSIL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetSIL(value byte) {
	response := safeGet("_scriptapi_register.h/SetSIL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRDI() {
	response := safeGet("_scriptapi_register.h/GetRDI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRDI(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRDI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetDIL() {
	response := safeGet("_scriptapi_register.h/GetDIL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetDIL(value byte) {
	response := safeGet("_scriptapi_register.h/SetDIL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRBP() {
	response := safeGet("_scriptapi_register.h/GetRBP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRBP(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRBP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetBPL() {
	response := safeGet("_scriptapi_register.h/GetBPL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetBPL(value byte) {
	response := safeGet("_scriptapi_register.h/SetBPL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRSP() {
	response := safeGet("_scriptapi_register.h/GetRSP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRSP(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRSP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetSPL() {
	response := safeGet("_scriptapi_register.h/GetSPL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetSPL(value byte) {
	response := safeGet("_scriptapi_register.h/SetSPL", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetRIP() {
	response := safeGet("_scriptapi_register.h/GetRIP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetRIP(value uint64) {
	response := safeGet("_scriptapi_register.h/SetRIP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR8() {
	response := safeGet("_scriptapi_register.h/GetR8", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR8(value uint64) {
	response := safeGet("_scriptapi_register.h/SetR8", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR8D() {
	response := safeGet("_scriptapi_register.h/GetR8D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR8D(value uint) {
	response := safeGet("_scriptapi_register.h/SetR8D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR8W() {
	response := safeGet("_scriptapi_register.h/GetR8W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR8W(value uint16) {
	response := safeGet("_scriptapi_register.h/SetR8W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR8B() {
	response := safeGet("_scriptapi_register.h/GetR8B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR8B(value byte) {
	response := safeGet("_scriptapi_register.h/SetR8B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR9() {
	response := safeGet("_scriptapi_register.h/GetR9", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR9(value uint64) {
	response := safeGet("_scriptapi_register.h/SetR9", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR9D() {
	response := safeGet("_scriptapi_register.h/GetR9D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR9D(value uint) {
	response := safeGet("_scriptapi_register.h/SetR9D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR9W() {
	response := safeGet("_scriptapi_register.h/GetR9W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR9W(value uint16) {
	response := safeGet("_scriptapi_register.h/SetR9W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR9B() {
	response := safeGet("_scriptapi_register.h/GetR9B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR9B(value byte) {
	response := safeGet("_scriptapi_register.h/SetR9B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR10() {
	response := safeGet("_scriptapi_register.h/GetR10", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR10(value uint64) {
	response := safeGet("_scriptapi_register.h/SetR10", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR10D() {
	response := safeGet("_scriptapi_register.h/GetR10D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR10D(value uint) {
	response := safeGet("_scriptapi_register.h/SetR10D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR10W() {
	response := safeGet("_scriptapi_register.h/GetR10W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR10W(value uint16) {
	response := safeGet("_scriptapi_register.h/SetR10W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR10B() {
	response := safeGet("_scriptapi_register.h/GetR10B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR10B(value byte) {
	response := safeGet("_scriptapi_register.h/SetR10B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR11() {
	response := safeGet("_scriptapi_register.h/GetR11", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR11(value uint64) {
	response := safeGet("_scriptapi_register.h/SetR11", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR11D() {
	response := safeGet("_scriptapi_register.h/GetR11D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR11D(value uint) {
	response := safeGet("_scriptapi_register.h/SetR11D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR11W() {
	response := safeGet("_scriptapi_register.h/GetR11W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR11W(value uint16) {
	response := safeGet("_scriptapi_register.h/SetR11W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR11B() {
	response := safeGet("_scriptapi_register.h/GetR11B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR11B(value byte) {
	response := safeGet("_scriptapi_register.h/SetR11B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR12() {
	response := safeGet("_scriptapi_register.h/GetR12", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR12(value uint64) {
	response := safeGet("_scriptapi_register.h/SetR12", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR12D() {
	response := safeGet("_scriptapi_register.h/GetR12D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR12D(value uint) {
	response := safeGet("_scriptapi_register.h/SetR12D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR12W() {
	response := safeGet("_scriptapi_register.h/GetR12W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR12W(value uint16) {
	response := safeGet("_scriptapi_register.h/SetR12W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR12B() {
	response := safeGet("_scriptapi_register.h/GetR12B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR12B(value byte) {
	response := safeGet("_scriptapi_register.h/SetR12B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR13() {
	response := safeGet("_scriptapi_register.h/GetR13", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR13(value uint64) {
	response := safeGet("_scriptapi_register.h/SetR13", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR13D() {
	response := safeGet("_scriptapi_register.h/GetR13D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR13D(value uint) {
	response := safeGet("_scriptapi_register.h/SetR13D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR13W() {
	response := safeGet("_scriptapi_register.h/GetR13W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR13W(value uint16) {
	response := safeGet("_scriptapi_register.h/SetR13W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR13B() {
	response := safeGet("_scriptapi_register.h/GetR13B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR13B(value byte) {
	response := safeGet("_scriptapi_register.h/SetR13B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR14() {
	response := safeGet("_scriptapi_register.h/GetR14", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR14(value uint64) {
	response := safeGet("_scriptapi_register.h/SetR14", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR14D() {
	response := safeGet("_scriptapi_register.h/GetR14D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR14D(value uint) {
	response := safeGet("_scriptapi_register.h/SetR14D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR14W() {
	response := safeGet("_scriptapi_register.h/GetR14W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR14W(value uint16) {
	response := safeGet("_scriptapi_register.h/SetR14W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR14B() {
	response := safeGet("_scriptapi_register.h/GetR14B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR14B(value byte) {
	response := safeGet("_scriptapi_register.h/SetR14B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR15() {
	response := safeGet("_scriptapi_register.h/GetR15", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR15(value uint64) {
	response := safeGet("_scriptapi_register.h/SetR15", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR15D() {
	response := safeGet("_scriptapi_register.h/GetR15D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR15D(value uint) {
	response := safeGet("_scriptapi_register.h/SetR15D", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR15W() {
	response := safeGet("_scriptapi_register.h/GetR15W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR15W(value uint16) {
	response := safeGet("_scriptapi_register.h/SetR15W", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetR15B() {
	response := safeGet("_scriptapi_register.h/GetR15B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetR15B(value byte) {
	response := safeGet("_scriptapi_register.h/SetR15B", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCAX() {
	response := safeGet("_scriptapi_register.h/GetCAX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCAX(value uint) {
	response := safeGet("_scriptapi_register.h/SetCAX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCBX() {
	response := safeGet("_scriptapi_register.h/GetCBX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCBX(value uint) {
	response := safeGet("_scriptapi_register.h/SetCBX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCCX() {
	response := safeGet("_scriptapi_register.h/GetCCX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCCX(value uint) {
	response := safeGet("_scriptapi_register.h/SetCCX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCDX() {
	response := safeGet("_scriptapi_register.h/GetCDX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCDX(value uint) {
	response := safeGet("_scriptapi_register.h/SetCDX", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCDI() {
	response := safeGet("_scriptapi_register.h/GetCDI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCDI(value uint) {
	response := safeGet("_scriptapi_register.h/SetCDI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCSI() {
	response := safeGet("_scriptapi_register.h/GetCSI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCSI(value uint) {
	response := safeGet("_scriptapi_register.h/SetCSI", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCBP() {
	response := safeGet("_scriptapi_register.h/GetCBP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCBP(value uint) {
	response := safeGet("_scriptapi_register.h/SetCBP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCSP() {
	response := safeGet("_scriptapi_register.h/GetCSP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCSP(value uint) {
	response := safeGet("_scriptapi_register.h/SetCSP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCIP() {
	response := safeGet("_scriptapi_register.h/GetCIP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCIP(value uint) {
	response := safeGet("_scriptapi_register.h/SetCIP", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) GetCFLAGS() {
	response := safeGet("_scriptapi_register.h/GetCFLAGS", map[string]string{})
	if len(response) == 0 {
		return
	}
}

func (r *register) SetCFLAGS(value uint) {
	response := safeGet("_scriptapi_register.h/SetCFLAGS", map[string]string{})
	if len(response) == 0 {
		return
	}
}
