package cpp2c

import "strings"

type AllowedPlatformInfo interface {
	GoBuildTag() string
	CxxIf() string
}

type PureVirtualDefault struct {
	ReturnType string
	MethodSig  string
	Body       string
}

var ActiveAllowedClassPrefix []string

func DefaultAllowClass(className string) bool {
	switch className {
	case "bool", "char", "signed char", "unsigned char",
		"short", "unsigned short", "int", "unsigned int",
		"long", "unsigned long", "long long", "unsigned long long",
		"float", "double", "long double",
		"void", "size_t", "ssize_t", "ptrdiff_t",
		"int8_t", "int16_t", "int32_t", "int64_t",
		"uint8_t", "uint16_t", "uint32_t", "uint64_t",
		"ImU8", "ImU16", "ImU32", "ImU64",
		"ImS8", "ImS16", "ImS32", "ImS64",
		"ImWchar", "ImWchar16", "ImWchar32":
		return true
	}

	if len(ActiveAllowedClassPrefix) > 0 {
		allowed := false
		for _, prefix := range ActiveAllowedClassPrefix {
			if strings.HasPrefix(className, prefix) {
				allowed = true
				break
			}
		}
		if !allowed {
			return false
		}
	}

	if strings.HasPrefix(className, `std::`) {
		return false
	}

	return true
}

func DefaultAllowMethod(className string, mm CppMethod) error {
	for _, p := range mm.Parameters {
		if p.ParameterType == "void" && p.Pointer {
			return ErrTooComplex
		}
		if p.ParameterName == "" && (p.ParameterType == "int" || p.ParameterType == "...") {
			return ErrTooComplex
		}
		if p.ParameterType == "va_list" {
			return ErrTooComplex
		}
	}
	if mm.ReturnType.ParameterType == "va_list" {
		return ErrTooComplex
	}
	return nil
}

func DefaultAllowFreeFunction(fn CppFreeFunction) error {
	for _, p := range fn.Parameters {
		if p.ParameterType == "void" && p.Pointer {
			return ErrTooComplex
		}
		if p.ParameterName == "" && (p.ParameterType == "int" || p.ParameterType == "...") {
			return ErrTooComplex
		}
		if p.ParameterType == "va_list" {
			return ErrTooComplex
		}
	}
	if fn.ReturnType.ParameterType == "va_list" {
		return ErrTooComplex
	}
	return nil
}

func DefaultAllowSignal(mm CppMethod) bool {
	if !mm.ReturnType.Void() {
		return false
	}
	return true
}
