package keystone

import (
	"fmt"
)

// Source: arm.h:13 -> ks_err_asm_arm
type Ks_err_asm_arm uint32

const (
	KsErrAsmArmInvalidoperand Ks_err_asm_arm = 512
	KsErrAsmArmMissingfeature Ks_err_asm_arm = 513
	KsErrAsmArmMnemonicfail Ks_err_asm_arm = 514
)

func (k Ks_err_asm_arm) String() string {
	switch k {
	case KsErrAsmArmInvalidoperand:
		return "Ks Err Asm Arm Invalidoperand"
	case KsErrAsmArmMissingfeature:
		return "Ks Err Asm Arm Missingfeature"
	case KsErrAsmArmMnemonicfail:
		return "Ks Err Asm Arm Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_arm(0x%X)", uint32(k))
	}
}

