package c2go

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ddkwork/bindgen/cc"
)

func CTagToGoName(name string) string {
	if strings.HasPrefix(name, "_") {
		name = name[1:]
	}
	result := strings.ToUpper(name[:1]) + name[1:]
	if GoReservedTypes[result] {
		result = result + "_"
	}
	return result
}

func CEnumMemberToGoName(name string) string {
	parts := strings.Split(name, "_")
	if len(parts) == 1 {
		return strings.ToUpper(name[:1]) + name[1:]
	}
	var result strings.Builder
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		result.WriteString(strings.ToUpper(part[:1]) + strings.ToLower(part[1:]))
	}
	return result.String()
}

func CFieldNameToGoName(name string) string {
	if name == "" {
		return "_"
	}
	parts := strings.Split(name, "_")
	if len(parts) == 1 {
		return strings.ToUpper(name[:1]) + name[1:]
	}
	var result strings.Builder
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		result.WriteString(strings.ToUpper(part[:1]) + strings.ToLower(part[1:]))
	}
	return result.String()
}

func CMacroNameToGoName(name string) string {
	parts := strings.Split(name, "_")
	if len(parts) == 1 {
		return strings.ToUpper(name[:1]) + name[1:]
	}
	var result strings.Builder
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}
		result.WriteString(strings.ToUpper(part[:1]) + strings.ToLower(part[1:]))
	}
	return result.String()
}

func ExtractEnumeratorValue(e *cc.Enumerator) (valueStr string, intVal int64, hasVal bool) {
	val := e.Value()
	if val == nil {
		return "iota", 0, false
	}
	switch v := val.(type) {
	case cc.Int64Value:
		iv := int64(v)
		return fmt.Sprintf("%d", iv), iv, true
	case cc.UInt64Value:
		uv := uint64(v)
		if uv > 0x7FFFFFFF {
			return fmt.Sprintf("0x%x", uv), int64(uv), true
		}
		return fmt.Sprintf("%d", uv), int64(uv), true
	default:
		return "iota", 0, false
	}
}

func FindSourceFileName(unit *cc.TranslationUnit, sources []cc.Source, headerMap map[string]string) string {
	posStr := unit.Position().String()
	posStr = strings.ReplaceAll(posStr, `\`, "/")
	for name := range headerMap {
		if strings.Contains(posStr, name) {
			return name
		}
	}
	for _, src := range sources {
		if strings.HasPrefix(src.Name, "<") {
			continue
		}
		if strings.Contains(posStr, src.Name) {
			return src.Name
		}
	}
	if strings.Contains(posStr, "xed-") {
		base := filepath.Base(posStr)
		for name := range headerMap {
			if strings.Contains(base, name) || strings.Contains(name, base) {
				return name
			}
		}
	}
	return "unknown.h"
}

func MapCTypedefToGo(declName string, t cc.Type) string {
	switch declName {
	case "BOOLEAN", "bool":
		return "bool"
	case "PBOOLEAN":
		return "*uint8"
	case "BOOL", "LONG", "INT", "INT32":
		return "int32"
	case "DWORD", "ULONG", "UINT", "UINT32":
		return "uint32"
	case "QWORD", "DWORD64", "ULONG64", "UINT64", "ULONGLONG":
		return "uint64"
	case "HANDLE":
		return "uintptr"
	case "PVOID", "LPVOID":
		return "uintptr"
	case "SIZE_T":
		return "uintptr"
	case "BYTE", "UCHAR", "UINT8":
		return "uint8"
	case "WORD", "USHORT", "UINT16":
		return "uint16"
	case "CHAR", "INT8":
		return "int8"
	case "SHORT", "INT16":
		return "int16"
	case "LONGLONG", "INT64":
		return "int64"
	case "WCHAR", "wchar_t":
		return "uint16"
	case "PBOOL", "PLONG", "PINT", "PINT32":
		return "*int32"
	case "PDWORD", "PULONG", "PUINT", "PUINT32":
		return "*uint32"
	case "PBYTE", "PUCHAR", "PUINT8":
		return "*uint8"
	case "PWORD", "PUSHORT", "PUINT16":
		return "*uint16"
	case "PCHAR", "PINT8":
		return "*int8"
	case "PSHORT", "PINT16":
		return "*int16"
	case "PQWORD", "PDWORD64", "PULONG64", "PUINT64":
		return "*uint64"
	case "PLONGLONG", "PINT64":
		return "*int64"
	case "PWCHAR", "LPWSTR":
		return "*rune"
	case "LPSTR", "LPCSTR", "LPCWSTR":
		return "*uint8"
	case "Time_t", "time_t", "__time64_t":
		return "time.Duration"
	case "LIST_ENTRY":
		return "LIST_ENTRY"
	case "PLIST_ENTRY":
		return "*LIST_ENTRY"
	}

	switch v := t.(type) {
	case *cc.PredefinedType:
		return MapPredefinedType(v.Kind(), "")
	case *cc.PointerType:
		return MapCPointerTypedefToGo(v)
	case *cc.StructType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return ""
	case *cc.EnumType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return "uint32"
	case *cc.UnionType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return ""
	case *cc.FunctionType:
		return ""
	default:
		goType := MapCTypeToGo(t)
		if goType != "" && goType != "void" && goType != CTagToGoName(declName) {
			return goType
		}
		return ""
	}
}

func MapPredefinedType(kind cc.Kind, fallback string) string {
	switch kind {
	case cc.Void:
		return "uintptr"
	case cc.Char, cc.SChar:
		return "int8"
	case cc.UChar:
		return "uint8"
	case cc.Short:
		return "int16"
	case cc.UShort:
		return "uint16"
	case cc.Int:
		return "int32"
	case cc.UInt:
		return "uint32"
	case cc.Long:
		return "int32"
	case cc.ULong:
		return "uint32"
	case cc.LongLong:
		return "int64"
	case cc.ULongLong:
		return "uint64"
	case cc.Float:
		return "float32"
	case cc.Double:
		return "float64"
	case cc.Bool:
		return "bool"
	default:
		return fallback
	}
}

func MapCPointerTypedefToGo(t *cc.PointerType) string {
	elem := t.Elem()
	if ft, ok := elem.(*cc.FunctionType); ok {
		return MapCFunctionTypeToGo(ft)
	}
	goElem := MapCTypeToGoForPointer(elem)
	return "*" + goElem
}

func MapCFunctionTypeToGo(ft *cc.FunctionType) string {
	goRet := MapCTypeToGo(ft.Result())
	params := ft.Parameters()
	var goParams []string
	for _, p := range params {
		goParamType := MapCTypeToGo(p.Type())
		if goParamType == "" || goParamType == "void" {
			continue
		}
		goParams = append(goParams, goParamType)
	}
	if goRet == "void" {
		return fmt.Sprintf("func(%s)", strings.Join(goParams, ", "))
	}
	return fmt.Sprintf("func(%s) %s", strings.Join(goParams, ", "), goRet)
}

func MapCTypeToGoForPointer(t cc.Type) string {
	if td := t.Typedef(); td != nil {
		tdName := td.Name()
		switch tdName {
		case "BOOLEAN", "bool":
			return "bool"
		case "BOOL", "LONG", "INT", "INT32":
			return "int32"
		case "DWORD", "ULONG", "UINT", "UINT32":
			return "uint32"
		case "QWORD", "DWORD64", "ULONG64", "UINT64", "ULONGLONG":
			return "uint64"
		case "BYTE", "UCHAR", "UINT8":
			return "uint8"
		case "WORD", "USHORT", "UINT16":
			return "uint16"
		case "CHAR", "INT8":
			return "int8"
		case "SHORT", "INT16":
			return "int16"
		case "LONGLONG", "INT64":
			return "int64"
		case "WCHAR", "wchar_t":
			return "uint16"
		case "PVOID", "LPVOID", "HANDLE", "SIZE_T":
			return "uintptr"
		case "Cardinal":
			return "uint32"
		case "Integer":
			return "int32"
		case "HMODULE":
			return "uintptr"
		case "IRSaveOEPToFile":
			return "uint32"
		}
		return CTagToGoName(tdName)
	}

	switch v := t.(type) {
	case *cc.PredefinedType:
		return MapPredefinedType(v.Kind(), "byte")
	case *cc.PointerType:
		return MapCPointerTypedefToGo(v)
	case *cc.ArrayType:
		elemType := MapCTypeToGoForPointer(v.Elem())
		size := v.Len()
		return fmt.Sprintf("[%d]%s", size, elemType)
	case *cc.StructType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return "struct{}"
	case *cc.EnumType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return "uint32"
	case *cc.UnionType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return "struct{}"
	default:
		return "byte"
	}
}

func MapCTypeToGo(t cc.Type) string {
	if td := t.Typedef(); td != nil {
		tdName := td.Name()
		switch tdName {
		case "BOOLEAN", "bool":
			return "bool"
		case "BOOL":
			return "int32"
		case "QWORD", "DWORD64", "ULONG64", "UINT64", "ULONGLONG":
			return "uint64"
		case "PVOID", "LPVOID", "HANDLE":
			return "uintptr"
		case "HWND", "HMODULE", "HINSTANCE", "HDC", "HMENU", "HICON", "HCURSOR", "HBRUSH", "HFONT", "HPEN", "HBITMAP", "HRGN", "HPALETTE", "HACCEL", "HHOOK":
			return "uintptr"
		case "WPARAM":
			return "uintptr"
		case "LPARAM":
			return "uintptr"
		case "FILETIME", "PFILETIME", "LPFILETIME":
			return "uintptr"
		case "LARGE_INTEGER", "PLARGE_INTEGER", "ULARGE_INTEGER", "PULARGE_INTEGER":
			return "uintptr"
		case "COLORREF":
			return "uint32"
		case "wchar_t", "Wchar_t", "WCHAR":
			return "uint16"
		case "SIZE_T":
			return "uintptr"
		case "PSIZE_T":
			return "*uintptr"
		case "DWORD", "ULONG", "UINT", "UINT32":
			return "uint32"
		case "LONG", "INT", "INT32":
			return "int32"
		case "BYTE", "UCHAR", "UINT8":
			return "uint8"
		case "WORD", "USHORT", "UINT16":
			return "uint16"
		case "CHAR", "INT8":
			return "int8"
		case "SHORT", "INT16":
			return "int16"
		case "LONGLONG", "INT64":
			return "int64"
		case "PBOOLEAN":
			return "*uint8"
		case "PBOOL", "PLONG", "PINT", "PINT32":
			return "*int32"
		case "PDWORD", "PULONG", "PUINT", "PUINT32":
			return "*uint32"
		case "PBYTE", "PUCHAR", "PUINT8":
			return "*uint8"
		case "PWORD", "PUSHORT", "PUINT16":
			return "*uint16"
		case "PCHAR", "PINT8":
			return "*int8"
		case "PSHORT", "PINT16":
			return "*int16"
		case "PQWORD", "PDWORD64", "PULONG64", "PUINT64":
			return "*uint64"
		case "PLONGLONG", "PINT64":
			return "*int64"
		case "PWCHAR", "LPWSTR", "LPCWSTR":
			return "*uint16"
		case "LPSTR", "LPCSTR":
			return "*uint8"
		case "LIST_ENTRY":
			return "LIST_ENTRY"
		case "PLIST_ENTRY":
			return "*LIST_ENTRY"
		case "Time_t", "time_t", "__time64_t":
			return "int64"
		case "PAGE_SIZE":
			return "uint32"
		case "ZyanU8":
			return "uint8"
		case "ZyanI8":
			return "int8"
		case "ZyanU16":
			return "uint16"
		case "ZyanI16":
			return "int16"
		case "ZyanU32":
			return "uint32"
		case "ZyanI32":
			return "int32"
		case "ZyanU64":
			return "uint64"
		case "ZyanI64":
			return "int64"
		case "ZyanUSize":
			return "uintptr"
		case "ZyanISize":
			return "uintptr"
		case "ZyanUPointer":
			return "uintptr"
		case "ZyanBool":
			return "uint8"
		case "ZyanStatus":
			return "uint32"
		case "ZyanVTable":
			return "uintptr"
		case "uint32_t", "Uint32_t", "UINT32_T":
			return "uint32"
		case "int32_t", "Int32_t", "INT32_T":
			return "int32"
		case "uint64_t", "Uint64_t", "UINT64_T":
			return "uint64"
		case "int64_t", "Int64_t", "INT64_T":
			return "int64"
		case "uint16_t", "Uint16_t", "UINT16_T":
			return "uint16"
		case "int16_t", "Int16_t", "INT16_T":
			return "int16"
		case "uint8_t", "Uint8_t", "UINT8_T":
			return "uint8"
		case "int8_t", "Int8_t", "INT8_T":
			return "int8"
		case "ssize_t", "Ssize_t", "SSIZE_T":
			return "int"
		case "intptr_t", "Intptr_t", "INTPTR_T":
			return "int"
		case "uintptr_t", "Uintptr_t", "UINTPTR_T":
			return "uintptr"
		case "size_t", "Size_t":
			return "uintptr"
		case "Xed_uint64_t":
			return "uint64"
		case "Xed_uint32_t":
			return "uint32"
		case "Xed_uint16_t":
			return "uint16"
		case "Xed_uint8_t":
			return "uint8"
		case "Xed_int64_t":
			return "int64"
		case "Xed_int32_t":
			return "int32"
		case "Xed_int16_t":
			return "int16"
		case "Xed_int8_t":
			return "int8"
		case "Xed_uint_t":
			return "uint32"
		case "Xed_bool_t":
			return "int32"
		case "Cardinal":
			return "uint32"
		case "Integer":
			return "int32"
		case "PChar":
			return "*int8"
		case "PCardinal":
			return "*uint32"
		case "IRSaveOEPToFile":
			return "uint32"
		case "PHMODULE":
			return "*uintptr"
		}
		return CTagToGoName(tdName)
	}

	switch v := t.(type) {
	case *cc.PredefinedType:
		switch v.Kind() {
		case cc.Void:
			return "void"
		case cc.Char, cc.SChar:
			return "int8"
		case cc.UChar:
			return "uint8"
		case cc.Short:
			return "int16"
		case cc.UShort:
			return "uint16"
		case cc.Int:
			return "int32"
		case cc.UInt:
			return "uint32"
		case cc.Long:
			return "int32"
		case cc.ULong:
			return "uint32"
		case cc.LongLong:
			return "int64"
		case cc.ULongLong:
			return "uint64"
		case cc.Float:
			return "float32"
		case cc.Double:
			return "float64"
		case cc.Bool:
			return "bool"
		default:
			return "int"
		}
	case *cc.PointerType:
		elem := v.Elem()
		goElem := MapCTypeToGo(elem)
		if goElem == "void" {
			return "unsafe.Pointer"
		}
		return "*" + goElem
	case *cc.ArrayType:
		elemType := MapCTypeToGo(v.Elem())
		size := v.Len()
		result := fmt.Sprintf("[%d]%s", size, elemType)
		return result
	case *cc.StructType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return "struct{}"
	case *cc.EnumType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return "uint32"
	case *cc.UnionType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return CTagToGoName(tagStr)
		}
		return "uint64"
	default:
		typeStr := t.String()
		if strings.Contains(typeStr, "BOOLEAN") {
			return "bool"
		}
		return "interface{}"
	}
}

func IsGoPrimitiveType(goType string) bool {
	switch goType {
	case "bool", "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64",
		"float32", "float64", "int", "uint", "uintptr", "*bool", "*int8", "*uint8",
		"*int16", "*uint16", "*int32", "*uint32", "*int64", "*uint64", "*uintptr":
		return true
	default:
		return false
	}
}
