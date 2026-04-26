package keystone

import (
	"fmt"
)

// Source: x86.h:13 -> ks_err_asm_x86
type Ks_err_asm_x86 uint32

const (
	KsErrAsmX86Invalidoperand Ks_err_asm_x86 = 512 + iota
	KsErrAsmX86Missingfeature
	KsErrAsmX86Mnemonicfail
)

func (k Ks_err_asm_x86) String() string {
	switch k {
	case KsErrAsmX86Invalidoperand:
		return "Ks Err Asm X86 Invalidoperand"
	case KsErrAsmX86Missingfeature:
		return "Ks Err Asm X86 Missingfeature"
	case KsErrAsmX86Mnemonicfail:
		return "Ks Err Asm X86 Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_x86(0x%X)", uint32(k))
	}
}
