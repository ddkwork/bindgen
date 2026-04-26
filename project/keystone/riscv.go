package keystone

import (
	"fmt"
)

// Source: riscv.h:13 -> ks_err_asm_riscv
type Ks_err_asm_riscv uint32

const (
	KsErrAsmRiscvInvalidoperand Ks_err_asm_riscv = 512 + iota
	KsErrAsmRiscvMissingfeature
	KsErrAsmRiscvMnemonicfail
)

func (k Ks_err_asm_riscv) String() string {
	switch k {
	case KsErrAsmRiscvInvalidoperand:
		return "Ks Err Asm Riscv Invalidoperand"
	case KsErrAsmRiscvMissingfeature:
		return "Ks Err Asm Riscv Missingfeature"
	case KsErrAsmRiscvMnemonicfail:
		return "Ks Err Asm Riscv Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_riscv(0x%X)", uint32(k))
	}
}
