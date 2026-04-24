package keystone

import (
	"fmt"
)

// Source: evm.h:13 -> ks_err_asm_evm
type Ks_err_asm_evm uint32

const (
	KsErrAsmEvmInvalidoperand Ks_err_asm_evm = 512
	KsErrAsmEvmMissingfeature Ks_err_asm_evm = 513
	KsErrAsmEvmMnemonicfail Ks_err_asm_evm = 514
)

func (k Ks_err_asm_evm) String() string {
	switch k {
	case KsErrAsmEvmInvalidoperand:
		return "Ks Err Asm Evm Invalidoperand"
	case KsErrAsmEvmMissingfeature:
		return "Ks Err Asm Evm Missingfeature"
	case KsErrAsmEvmMnemonicfail:
		return "Ks Err Asm Evm Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_evm(0x%X)", uint32(k))
	}
}

