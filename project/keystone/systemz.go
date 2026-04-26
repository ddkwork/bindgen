package keystone

import (
	"fmt"
)

// Source: systemz.h:13 -> ks_err_asm_systemz
type Ks_err_asm_systemz uint32

const (
	KsErrAsmSystemzInvalidoperand Ks_err_asm_systemz = 512 + iota
	KsErrAsmSystemzMissingfeature
	KsErrAsmSystemzMnemonicfail
)

func (k Ks_err_asm_systemz) String() string {
	switch k {
	case KsErrAsmSystemzInvalidoperand:
		return "Ks Err Asm Systemz Invalidoperand"
	case KsErrAsmSystemzMissingfeature:
		return "Ks Err Asm Systemz Missingfeature"
	case KsErrAsmSystemzMnemonicfail:
		return "Ks Err Asm Systemz Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_systemz(0x%X)", uint32(k))
	}
}
