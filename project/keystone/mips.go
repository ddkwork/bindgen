package keystone

import (
	"fmt"
)

// Source: mips.h:13 -> ks_err_asm_mips
type Ks_err_asm_mips uint32

const (
	KsErrAsmMipsInvalidoperand Ks_err_asm_mips = 512 + iota
	KsErrAsmMipsMissingfeature
	KsErrAsmMipsMnemonicfail
)

func (k Ks_err_asm_mips) String() string {
	switch k {
	case KsErrAsmMipsInvalidoperand:
		return "Ks Err Asm Mips Invalidoperand"
	case KsErrAsmMipsMissingfeature:
		return "Ks Err Asm Mips Missingfeature"
	case KsErrAsmMipsMnemonicfail:
		return "Ks Err Asm Mips Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_mips(0x%X)", uint32(k))
	}
}
