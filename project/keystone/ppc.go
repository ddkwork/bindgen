package keystone

import (
	"fmt"
)

// Source: ppc.h:13 -> ks_err_asm_ppc
type Ks_err_asm_ppc uint32

const (
	KsErrAsmPpcInvalidoperand Ks_err_asm_ppc = 512
	KsErrAsmPpcMissingfeature Ks_err_asm_ppc = 513
	KsErrAsmPpcMnemonicfail Ks_err_asm_ppc = 514
)

func (k Ks_err_asm_ppc) String() string {
	switch k {
	case KsErrAsmPpcInvalidoperand:
		return "Ks Err Asm Ppc Invalidoperand"
	case KsErrAsmPpcMissingfeature:
		return "Ks Err Asm Ppc Missingfeature"
	case KsErrAsmPpcMnemonicfail:
		return "Ks Err Asm Ppc Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_ppc(0x%X)", uint32(k))
	}
}

