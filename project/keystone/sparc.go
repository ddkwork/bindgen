package keystone

import (
	"fmt"
)

// Source: sparc.h:13 -> ks_err_asm_sparc
type Ks_err_asm_sparc uint32

const (
	KsErrAsmSparcInvalidoperand Ks_err_asm_sparc = 512
	KsErrAsmSparcMissingfeature Ks_err_asm_sparc = 513
	KsErrAsmSparcMnemonicfail Ks_err_asm_sparc = 514
)

func (k Ks_err_asm_sparc) String() string {
	switch k {
	case KsErrAsmSparcInvalidoperand:
		return "Ks Err Asm Sparc Invalidoperand"
	case KsErrAsmSparcMissingfeature:
		return "Ks Err Asm Sparc Missingfeature"
	case KsErrAsmSparcMnemonicfail:
		return "Ks Err Asm Sparc Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_sparc(0x%X)", uint32(k))
	}
}

