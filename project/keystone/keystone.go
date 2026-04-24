package keystone

import (
	"fmt"
)

// Source: keystone.h
type (
	Ks_engine = Ks_struct
)

type Ks_sym_resolver func(*int8, *uint64) bool

// Source: keystone.h:57 -> ks_arch
type Ks_arch uint32

const (
	KsArchArm Ks_arch = 1
	KsArchArm64 Ks_arch = 2
	KsArchMips Ks_arch = 3
	KsArchX86 Ks_arch = 4
	KsArchPpc Ks_arch = 5
	KsArchSparc Ks_arch = 6
	KsArchSystemz Ks_arch = 7
	KsArchHexagon Ks_arch = 8
	KsArchEvm Ks_arch = 9
	KsArchRiscv Ks_arch = 10
	KsArchMax Ks_arch = 11
)

func (k Ks_arch) String() string {
	switch k {
	case KsArchArm:
		return "Ks Arch Arm"
	case KsArchArm64:
		return "Ks Arch Arm 64"
	case KsArchMips:
		return "Ks Arch Mips"
	case KsArchX86:
		return "Ks Arch X86"
	case KsArchPpc:
		return "Ks Arch Ppc"
	case KsArchSparc:
		return "Ks Arch Sparc"
	case KsArchSystemz:
		return "Ks Arch Systemz"
	case KsArchHexagon:
		return "Ks Arch Hexagon"
	case KsArchEvm:
		return "Ks Arch Evm"
	case KsArchRiscv:
		return "Ks Arch Riscv"
	case KsArchMax:
		return "Ks Arch Max"
	default:
		return fmt.Sprintf("Ks_arch(0x%X)", uint32(k))
	}
}

// Source: keystone.h:72 -> ks_mode
type Ks_mode uint32

const (
	KsModeLittleEndian Ks_mode = 0
	KsModeBigEndian Ks_mode = 1073741824
	KsModeArm Ks_mode = 1
	KsModeThumb Ks_mode = 16
	KsModeV8 Ks_mode = 64
	KsModeMicro Ks_mode = 16
	KsModeMips3 Ks_mode = 32
	KsModeMips32r6 Ks_mode = 64
	KsModeMips32 Ks_mode = 4
	KsModeMips64 Ks_mode = 8
	KsMode16 Ks_mode = 2
	KsMode32 Ks_mode = 4
	KsMode64 Ks_mode = 8
	KsModePpc32 Ks_mode = 4
	KsModePpc64 Ks_mode = 8
	KsModeQpx Ks_mode = 16
	KsModeRiscv32 Ks_mode = 4
	KsModeRiscv64 Ks_mode = 8
	KsModeSparc32 Ks_mode = 4
	KsModeSparc64 Ks_mode = 8
	KsModeV9 Ks_mode = 16
)

func (k Ks_mode) String() string {
	switch k {
	case KsModeLittleEndian:
		return "Ks Mode Little Endian"
	case KsModeBigEndian:
		return "Ks Mode Big Endian"
	case KsModeArm:
		return "Ks Mode Arm"
	case KsModeThumb:
		return "Ks Mode Thumb"
	case KsModeV8:
		return "Ks Mode V8"
	case KsModeMips3:
		return "Ks Mode Mips 3"
	case KsModeMips32:
		return "Ks Mode Mips 32"
	case KsModeMips64:
		return "Ks Mode Mips 64"
	case KsMode16:
		return "Ks Mode 16"
	default:
		return fmt.Sprintf("Ks_mode(0x%X)", uint32(k))
	}
}

// Source: keystone.h:109 -> ks_err
type Ks_err uint32

const (
	KsErrOk Ks_err = 0
	KsErrNomem Ks_err = 1
	KsErrArch Ks_err = 2
	KsErrHandle Ks_err = 3
	KsErrMode Ks_err = 4
	KsErrVersion Ks_err = 5
	KsErrOptInvalid Ks_err = 6
	KsErrAsmExprToken Ks_err = 128
	KsErrAsmDirectiveValueRange Ks_err = 129
	KsErrAsmDirectiveId Ks_err = 130
	KsErrAsmDirectiveToken Ks_err = 131
	KsErrAsmDirectiveStr Ks_err = 132
	KsErrAsmDirectiveComma Ks_err = 133
	KsErrAsmDirectiveRelocName Ks_err = 134
	KsErrAsmDirectiveRelocToken Ks_err = 135
	KsErrAsmDirectiveFpoint Ks_err = 136
	KsErrAsmDirectiveUnknown Ks_err = 137
	KsErrAsmDirectiveEqu Ks_err = 138
	KsErrAsmDirectiveInvalid Ks_err = 139
	KsErrAsmVariantInvalid Ks_err = 140
	KsErrAsmExprBracket Ks_err = 141
	KsErrAsmSymbolModifier Ks_err = 142
	KsErrAsmSymbolRedefined Ks_err = 143
	KsErrAsmSymbolMissing Ks_err = 144
	KsErrAsmRparen Ks_err = 145
	KsErrAsmStatToken Ks_err = 146
	KsErrAsmUnsupported Ks_err = 147
	KsErrAsmMacroToken Ks_err = 148
	KsErrAsmMacroParen Ks_err = 149
	KsErrAsmMacroEqu Ks_err = 150
	KsErrAsmMacroArgs Ks_err = 151
	KsErrAsmMacroLevelsExceed Ks_err = 152
	KsErrAsmMacroStr Ks_err = 153
	KsErrAsmMacroInvalid Ks_err = 154
	KsErrAsmEscBackslash Ks_err = 155
	KsErrAsmEscOctal Ks_err = 156
	KsErrAsmEscSequence Ks_err = 157
	KsErrAsmEscStr Ks_err = 158
	KsErrAsmTokenInvalid Ks_err = 159
	KsErrAsmInsnUnsupported Ks_err = 160
	KsErrAsmFixupInvalid Ks_err = 161
	KsErrAsmLabelInvalid Ks_err = 162
	KsErrAsmFragmentInvalid Ks_err = 163
	KsErrAsmInvalidoperand Ks_err = 512
	KsErrAsmMissingfeature Ks_err = 513
	KsErrAsmMnemonicfail Ks_err = 514
)

func (k Ks_err) String() string {
	switch k {
	case KsErrOk:
		return "Ks Err Ok"
	case KsErrNomem:
		return "Ks Err Nomem"
	case KsErrArch:
		return "Ks Err Arch"
	case KsErrHandle:
		return "Ks Err Handle"
	case KsErrMode:
		return "Ks Err Mode"
	case KsErrVersion:
		return "Ks Err Version"
	case KsErrOptInvalid:
		return "Ks Err Opt Invalid"
	case KsErrAsmExprToken:
		return "Ks Err Asm Expr Token"
	case KsErrAsmDirectiveValueRange:
		return "Ks Err Asm Directive Value Range"
	case KsErrAsmDirectiveId:
		return "Ks Err Asm Directive Id"
	case KsErrAsmDirectiveToken:
		return "Ks Err Asm Directive Token"
	case KsErrAsmDirectiveStr:
		return "Ks Err Asm Directive Str"
	case KsErrAsmDirectiveComma:
		return "Ks Err Asm Directive Comma"
	case KsErrAsmDirectiveRelocName:
		return "Ks Err Asm Directive Reloc Name"
	case KsErrAsmDirectiveRelocToken:
		return "Ks Err Asm Directive Reloc Token"
	case KsErrAsmDirectiveFpoint:
		return "Ks Err Asm Directive Fpoint"
	case KsErrAsmDirectiveUnknown:
		return "Ks Err Asm Directive Unknown"
	case KsErrAsmDirectiveEqu:
		return "Ks Err Asm Directive Equ"
	case KsErrAsmDirectiveInvalid:
		return "Ks Err Asm Directive Invalid"
	case KsErrAsmVariantInvalid:
		return "Ks Err Asm Variant Invalid"
	case KsErrAsmExprBracket:
		return "Ks Err Asm Expr Bracket"
	case KsErrAsmSymbolModifier:
		return "Ks Err Asm Symbol Modifier"
	case KsErrAsmSymbolRedefined:
		return "Ks Err Asm Symbol Redefined"
	case KsErrAsmSymbolMissing:
		return "Ks Err Asm Symbol Missing"
	case KsErrAsmRparen:
		return "Ks Err Asm Rparen"
	case KsErrAsmStatToken:
		return "Ks Err Asm Stat Token"
	case KsErrAsmUnsupported:
		return "Ks Err Asm Unsupported"
	case KsErrAsmMacroToken:
		return "Ks Err Asm Macro Token"
	case KsErrAsmMacroParen:
		return "Ks Err Asm Macro Paren"
	case KsErrAsmMacroEqu:
		return "Ks Err Asm Macro Equ"
	case KsErrAsmMacroArgs:
		return "Ks Err Asm Macro Args"
	case KsErrAsmMacroLevelsExceed:
		return "Ks Err Asm Macro Levels Exceed"
	case KsErrAsmMacroStr:
		return "Ks Err Asm Macro Str"
	case KsErrAsmMacroInvalid:
		return "Ks Err Asm Macro Invalid"
	case KsErrAsmEscBackslash:
		return "Ks Err Asm Esc Backslash"
	case KsErrAsmEscOctal:
		return "Ks Err Asm Esc Octal"
	case KsErrAsmEscSequence:
		return "Ks Err Asm Esc Sequence"
	case KsErrAsmEscStr:
		return "Ks Err Asm Esc Str"
	case KsErrAsmTokenInvalid:
		return "Ks Err Asm Token Invalid"
	case KsErrAsmInsnUnsupported:
		return "Ks Err Asm Insn Unsupported"
	case KsErrAsmFixupInvalid:
		return "Ks Err Asm Fixup Invalid"
	case KsErrAsmLabelInvalid:
		return "Ks Err Asm Label Invalid"
	case KsErrAsmFragmentInvalid:
		return "Ks Err Asm Fragment Invalid"
	case KsErrAsmInvalidoperand:
		return "Ks Err Asm Invalidoperand"
	case KsErrAsmMissingfeature:
		return "Ks Err Asm Missingfeature"
	case KsErrAsmMnemonicfail:
		return "Ks Err Asm Mnemonicfail"
	default:
		return fmt.Sprintf("Ks_err(0x%X)", uint32(k))
	}
}

// Source: keystone.h:173 -> ks_opt_type
type Ks_opt_type uint32

const (
	KsOptSyntax Ks_opt_type = 1
	KsOptSymResolver Ks_opt_type = 2
)

func (k Ks_opt_type) String() string {
	switch k {
	case KsOptSyntax:
		return "Ks Opt Syntax"
	case KsOptSymResolver:
		return "Ks Opt Sym Resolver"
	default:
		return fmt.Sprintf("Ks_opt_type(0x%X)", uint32(k))
	}
}

// Source: keystone.h:180 -> ks_opt_value
type Ks_opt_value uint32

const (
	KsOptSyntaxIntel Ks_opt_value = 1
	KsOptSyntaxAtt Ks_opt_value = 2
	KsOptSyntaxNasm Ks_opt_value = 4
	KsOptSyntaxMasm Ks_opt_value = 8
	KsOptSyntaxGas Ks_opt_value = 16
	KsOptSyntaxRadix16 Ks_opt_value = 32
)

func (k Ks_opt_value) String() string {
	switch k {
	case KsOptSyntaxIntel:
		return "Ks Opt Syntax Intel"
	case KsOptSyntaxAtt:
		return "Ks Opt Syntax Att"
	case KsOptSyntaxNasm:
		return "Ks Opt Syntax Nasm"
	case KsOptSyntaxMasm:
		return "Ks Opt Syntax Masm"
	case KsOptSyntaxGas:
		return "Ks Opt Syntax Gas"
	case KsOptSyntaxRadix16:
		return "Ks Opt Syntax Radix 16"
	default:
		return fmt.Sprintf("Ks_opt_value(0x%X)", uint32(k))
	}
}

// Source: keystone.h:38 -> ks_struct
type Ks_struct struct {

}

func KsMakeVersion(Major uint32, Minor uint32) uint32 {
	return uint32(((Major<<8)+Minor))
}

// Source: keystone.h -> Macro constants
const (
	KsApiMajor uint32 = 0
	KsApiMinor uint32 = 9
	KsVersionMajor uint32 = KsApiMajor
	KsVersionMinor uint32 = KsApiMinor
	KsVersionExtra uint32 = 2
	KsErrAsm uint32 = 128
	KsErrAsmArch uint32 = 512
)

