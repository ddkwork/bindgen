package keystone

import (
	"fmt"
)

// Source: hexagon.h:13 -> ks_err_asm_hexagon
type Ks_err_asm_hexagon uint32

const (
	KsErrAsmHexagonInvalidoperand Ks_err_asm_hexagon = 512 + iota
	KsErrAsmHexagonMissingfeature
	KsErrAsmHexagonMnemonicfail
)

func (k Ks_err_asm_hexagon) String() string {
	switch k {
	case KsErrAsmHexagonInvalidoperand:
		return "Ks Err Asm Hexagon Invalidoperand"
	case KsErrAsmHexagonMissingfeature:
		return "Ks Err Asm Hexagon Missingfeature"
	case KsErrAsmHexagonMnemonicfail:
		return "Ks Err Asm Hexagon Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err_asm_hexagon(0x%X)", uint32(k))
	}
}

