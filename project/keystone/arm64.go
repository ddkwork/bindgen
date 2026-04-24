package keystone

import (
	"fmt"
)

// Source: arm64.h:13 -> ks_err_asm_arm64
type Ks_err_asm_arm64 uint32

const (
	KsErrAsmArm64Invalidoperand Ks_err_asm_arm64 = 512
	KsErrAsmArm64Missingfeature Ks_err_asm_arm64 = 513
	KsErrAsmArm64Mnemonicfail Ks_err_asm_arm64 = 514
)

func (k Ks_err_asm_arm64) String() string {
	switch k {
	case KsErrAsmArm64Invalidoperand:
		return "Ks Err Asm Arm 64 Invalidoperand"
	case KsErrAsmArm64Missingfeature:
		return "Ks Err Asm Arm 64 Missingfeature"
	case KsErrAsmArm64Mnemonicfail:
		return "Ks Err Asm Arm 64 Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_arm64(0x%X)", uint32(k))
	}
}

