package cpp2c

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// CppComment renders a string safely in a C++ block comment.
// It strips interior nested comments.
func CppComment(s string) string {
	// Remove nested comments
	uncomment := strings.NewReplacer("/*", "", "*/", "")
	return "/* " + uncomment.Replace(s) + " */ "
}

func CReservedWord(s string) bool {
	switch s {
	case "default", "self":
		return true
	default:
		return false
	}
}

// isArrayFieldType checks if the parameter's type string contains array brackets,
// indicating it's a C/C++ array like ImVec4[61].
func isArrayFieldType(p CppParameter) bool {
	return strings.Contains(p.ParameterType, "[")
}

// arrayFieldElementType extracts the element type from an array type string.
// e.g. "ImVec4[61]" returns "ImVec4", "ImGuiIO::bool[5]" returns "bool".
func arrayFieldElementType(p CppParameter) string {
	t := p.ParameterType
	if idx := strings.Index(t, "["); idx >= 0 {
		elem := t[:idx]
		if strings.Contains(elem, "::") {
			parts := strings.Split(elem, "::")
			return parts[len(parts)-1]
		}
		return elem
	}
	return t
}

// isLikelyStructType checks if a type name looks like a C++ struct/class/typedef
// name (starts with letter or underscore, contains only valid identifier chars).
func isLikelyStructType(name string) bool {
	if name == "" {
		return false
	}
	if !((name[0] >= 'A' && name[0] <= 'Z') || (name[0] >= 'a' && name[0] <= 'z') || name[0] == '_') {
		return false
	}
	for _, c := range name {
		if c == ' ' || c == '<' || c == '>' || c == '(' || c == ')' || c == ',' {
			return false
		}
	}
	return true
}

func isCabiIntegerType(t string) bool {
	switch t {
	case "int", "unsigned int", "intptr_t", "uintptr_t",
		"int64_t", "uint64_t", "int32_t", "uint32_t",
		"int16_t", "uint16_t", "int8_t", "uint8_t":
		return true
	}
	return false
}

// IsDirectCType returns true if the type name is a valid C type that doesn't need mapping.
func IsDirectCType(t string) bool {
	switch t {
	case "void", "bool", "char", "short", "int", "long", "float", "double",
		"unsigned", "signed", "size_t", "__size_t", "ptrdiff_t", "int8_t", "uint8_t",
		"int16_t", "uint16_t", "int32_t", "uint32_t", "int64_t", "uint64_t",
		"intptr_t", "uintptr_t", "wchar_t", "char16_t", "char32_t",
		"unsigned char", "signed char", "unsigned short", "unsigned int", "unsigned long":
		return true
	}
	return false
}

func (p CppParameter) cParameterName() string {
	// Also make the first letter uppercase so it becomes public in Go
	parmName := p.ParameterName
	if CReservedWord(parmName) {
		parmName += "Val"
	}

	return parmName
}

func cabiCallbackName(c CppClass, m CppMethod) string {
	return "miqt_exec_callback_" + CabiClassName(c.ClassName) + "_" + m.SafeMethodName()
}

func cabiNewName(c CppClass, i int) string {
	return CabiClassName(c.ClassName) + `_new` + maybeSuffix(i)
}

func cabiDeleteName(c CppClass) string {
	return CabiClassName(c.ClassName) + `_delete`
}

func cabiVirtBaseName(c CppClass) string {
	return CabiClassName(c.ClassName) + `_virtbase`
}

func cabiMethodName(c CppClass, m CppMethod) string {
	return CabiClassName(c.ClassName) + `_` + m.SafeMethodName()
}

func cabiConnectName(c CppClass, m CppMethod) string {
	return CabiClassName(c.ClassName) + `_connect_` + m.SafeMethodName()
}

func cabiVirtualBaseName(c CppClass, m CppMethod) string {
	return CabiClassName(c.ClassName) + `_virtualbase_` + m.SafeMethodName()
}

func cabiProtectedBaseName(c CppClass, m CppMethod) string {
	return CabiClassName(c.ClassName) + `_protectedbase_` + m.SafeMethodName()
}

func cabiOverrideVirtualName(c CppClass, m CppMethod) string {
	return CabiClassName(c.ClassName) + `_override_virtual_` + m.SafeMethodName()
}

func CppSubclassName(c CppClass) string {
	return "MiqtVirtual" + strings.Replace(c.ClassName, `::`, ``, -1)
}

func (p CppParameter) RenderTypeCabi() string {

	switch p.ParameterType {
	case "quint8":
		ret := "uint8_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "quint16":
		ret := "uint16_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "quint32":
		ret := "uint32_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "quint64":
		ret := "uint64_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qint8":
		ret := "int8_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qint16":
		ret := "int16_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qint32":
		ret := "int32_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qint64":
		ret := "int64_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qlonglong", "long long":
		ret := "long long"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qulonglong", "unsigned long long":
		ret := "unsigned long long"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qintptr", "QIntegerForSizeof<void *>::Signed":
		ret := "intptr_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "quintptr", "uintptr", "QIntegerForSizeof<void *>::Unsigned":
		ret := "uintptr_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qsizetype", "qptrdiff", "QIntegerForSizeof<std::size_t>::Signed":
		ret := "intptr_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qreal":
		ret := "double"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	case "qfloat16":
		ret := "_Float16"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	}

	// QtPrivate:: types are internal, not real API classes - map to void*
	if strings.Contains(p.ParameterType, "QtPrivate::") || strings.HasPrefix(p.ParameterType, "QtPrivate__") {
		ret := "void"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		} else {
			ret += "*"
		}
		return ret
	}

	if p.ParameterType == "QString" || p.ParameterType == "QByteArray" {
		return "struct miqt_string"
	}

	switch {
	case strings.HasPrefix(p.ParameterType, "QFlags<"):
		return "int"

	case strings.Contains(p.ParameterType, "QIntegerForSizeof<"):
		ret := "intptr_t"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret

	case strings.HasPrefix(p.ParameterType, "QPartialOrdering"):
		return "int"

	case strings.HasPrefix(p.ParameterType, "QSpan<"):
		return "void*"

	case strings.HasPrefix(p.ParameterType, "QBindable<"):
		return "void*"

	case p.ParameterType == "QMetaObject::Connection" || p.ParameterType == "QByteArrayView::const_pointer" || p.ParameterType == "QByteArrayView::const_iterator":
		return "void*"
	case p.ParameterType == "QByteArray::FromBase64Result" || p.ParameterType == "QByteArray::DataPointer" ||
		p.ParameterType == "QArrayDataPointer<char>" || p.ParameterType == "QArrayDataPointer<void>":
		return "void*"
	case p.ParameterType == "QChar::fromUcs4::R" || p.ParameterType == "QChar::fromUcs2::R":
		return "int"
	case p.ParameterType == "QUtf8StringView" || p.ParameterType == "QAnyStringView" ||
		p.ParameterType == "QStringView" || p.ParameterType == "QByteArrayView" ||
		p.ParameterType == "QLatin1StringView" || p.ParameterType == "QLatin1String" ||
		p.ParameterType == "QStringRef":
		return "void*"
	case p.ParameterType == "Disambiguated_t":
		ret := "int"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		return ret
	case p.ParameterType == "QBasicAtomicInteger<int>":
		ret := "int"
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		return ret
	case strings.HasPrefix(p.ParameterType, "QByteArrayView::") || strings.HasPrefix(p.ParameterType, "QStringView::"):
		return "void*"
	case strings.HasPrefix(p.ParameterType, "QByteArray::") && (strings.HasSuffix(p.ParameterType, "iterator") || strings.HasSuffix(p.ParameterType, "reverse_iterator")):
		return "void*"
	case strings.HasPrefix(p.ParameterType, "QString::") && (strings.HasSuffix(p.ParameterType, "iterator") || strings.HasSuffix(p.ParameterType, "reverse_iterator")):
		return "void*"
	}

	if inner, _, ok := p.QListOf(); ok {
		return "struct miqt_array " + CppComment("of "+inner.RenderTypeCabi())
	}

	if inner, ok := p.QSetOf(); ok {
		return "struct miqt_array " + CppComment("set of "+inner.RenderTypeCabi())
	}

	if inner1, inner2, _, ok := p.QMapOf(); ok {
		return "struct miqt_map " + CppComment("of "+inner1.RenderTypeCabi()+" to "+inner2.RenderTypeCabi())
	}

	if inner1, inner2, ok := p.QPairOf(); ok {
		return "struct miqt_map " + CppComment("tuple of "+inner1.RenderTypeCabi()+" and "+inner2.RenderTypeCabi())
	}

	switch p.ParameterType {
	case "iterator", "const_iterator", "reverse_iterator", "const_reverse_iterator", "key_value_iterator", "const_key_value_iterator":
		return "void*"
	case "strong_ordering", "weak_ordering", "partial_ordering":
		return "int"
	case "GLenum", "QOpenGLTimeMonitor::GLenum", "QOpenGLFramebufferObject::GLenum":
		return "unsigned int"
	case "GLuint", "QOpenGLTimeMonitor::GLuint", "QOpenGLFramebufferObject::GLuint":
		return "unsigned int"
	case "GLuint64", "QOpenGLTimeMonitor::GLuint64", "QOpenGLFramebufferObject::GLuint64":
		return "uint64_t"
	case "GLint":
		return "int"
	case "GLfloat":
		return "float"
	case "GLdouble":
		return "double"
	case "GLsizei":
		return "int"
	case "GLboolean":
		return "unsigned char"
	case "GLbitfield":
		return "unsigned int"
	case "GLsync":
		return "void*"
	case "QQmlV4FunctionPtr":
		return "void*"
	case "CreateTextureOption":
		return "int"
	case "va_list":
		return "void*"
	case "HWND", "HINSTANCE", "HDC", "HGLRC", "HMENU", "HICON", "HCURSOR", "HBRUSH", "HRGN", "HMONITOR":
		return "void*"
	case "WPARAM":
		return "uintptr_t"
	case "LPARAM", "LRESULT":
		return "intptr_t"
	case "HRESULT":
		return "int32_t"
	}

	if strings.HasPrefix(p.ParameterType, "QVariantRef<") || strings.HasPrefix(p.ParameterType, "QVariantPointer<") {
		return "void*"
	}

	if (p.Pointer || p.ByRef) && p.QtClassType() {
		if !KnownClassnames.Has(p.ParameterType) && strings.Contains(p.ParameterType, "::") && !p.IsKnownEnum() {
			if !ActiveFilter.AllowClass(p.ParameterType) {
				if p.Pointer {
					return "int*"
				}
				return "int"
			}
		}
		if p.PointerCount > 1 {
			return CabiClassName(p.ParameterType) + strings.Repeat("*", p.PointerCount)
		}
		return CabiClassName(p.ParameterType) + "*"
	}

	if p.QtClassType() && !p.Pointer {
		if !KnownClassnames.Has(p.ParameterType) && strings.Contains(p.ParameterType, "::") && !p.IsKnownEnum() {
			if !ActiveFilter.AllowClass(p.ParameterType) {
				return "int"
			}
		}
		return CabiClassName(p.ParameterType) + "*"
	}

	// https://github.com/qt/qtbase/blob/v5.15.16-lts-lgpl/src/corelib/global/qglobal.h#L233
	// https://github.com/qt/qtbase/blob/v6.9.0/src/corelib/global/qtypes.h#L50
	ret := p.ParameterType
	switch p.ParameterType {
	case "uchar":
		ret = "unsigned char"
	case "ushort":
		ret = "unsigned short"
	case "uint":
		ret = "unsigned int"
	case "ulong":
		ret = "unsigned long"
	case "std::byte":
		ret = "uint8_t"
	case "std::chrono::milliseconds":
		ret = "int64_t"
	case "std::chrono::days":
		ret = "int64_t"
	case "std::chrono::nanoseconds":
		ret = "int64_t"
	case "duration<long long, milli>":
		ret = "int64_t"
	case "duration<long long, nano>":
		ret = "int64_t"
	case "duration<long long>":
		ret = "int64_t"
	case "__size_t":
		ret = "size_t"
	}

	if p.Const {
		// This is needed for const-correctness for calling some overloads
		// e.g. QShortcut ctor taking (QWidget* parent, const char* member) signal -
		// the signal/slot requires that member is const, not just plain char*
		ret = "const " + ret
	}

	if ft, ok := p.QFlagsOf(); ok {
		if e, ok := KnownEnums.Get(ft.ParameterType); ok {
			ret = e.Enum.UnderlyingType.RenderTypeCabi()
		} else {
			ret = "int"
		}

	} else if e, ok := KnownEnums.Get(p.ParameterType); ok {
		ret = e.Enum.UnderlyingType.RenderTypeCabi()

	}

	// Catch-all for unknown template/STL types that can't be expressed in C
	if strings.Contains(ret, "<") && !strings.HasPrefix(ret, "QList<") &&
		!strings.HasPrefix(ret, "QVector<") && !strings.HasPrefix(ret, "QHash<") &&
		!strings.HasPrefix(ret, "QMap<") && !strings.HasPrefix(ret, "QSet<") &&
		!strings.HasPrefix(ret, "QPair<") && !strings.HasPrefix(ret, "QFlags<") {
		ret = "void*"
		if p.Const {
			ret = "const " + ret
		}
		return ret
	}

	// Catch-all for function pointer types - map to void*
	if strings.Contains(p.ParameterType, "(*)") {
		ret := "void*"
		if p.Const {
			ret = "const " + ret
		}
		if p.Pointer {
			ret += strings.Repeat("*", p.PointerCount)
		}
		if p.ByRef {
			ret += "*"
		}
		return ret
	}

	// Check if this is a typedef to a function pointer
	if td, ok := KnownTypedefs.Get(p.ParameterType); ok {
		ut := td.Typedef.UnderlyingType
		if strings.Contains(ut.ParameterType, "(*)") || (ut.Pointer && !ut.QtClassType() && !ut.IsKnownEnum() && !IsDirectCType(ut.ParameterType) && ut.ParameterType != "void") {
			ret := "void*"
			if p.Const {
				ret = "const " + ret
			}
			if p.Pointer {
				ret += strings.Repeat("*", p.PointerCount)
			}
			if p.ByRef {
				ret += "*"
			}
			return ret
		}
	}

	// Catch-all for unknown non-class types (enum-like or inner class refs without prefix)
	// If the type is not a known Qt class and not a basic C type, treat it as int.
	// If the type HAS been remapped by the switch above (e.g. qfloat16 → _Float16,
	// duration<...> → int64_t), keep the remapped type.
	// For pointer/reference types, use void* instead of int* since the actual
	// pointed-to type is unknown and int* is not compatible with struct pointers.
	if !p.QtClassType() && !p.IsKnownEnum() && !p.IsFlagType() &&
		!IsDirectCType(p.ParameterType) && !strings.Contains(p.ParameterType, "<") {
		if ret == p.ParameterType {
			if p.Pointer || p.ByRef {
				ret = "void"
			} else {
				ret = "int"
			}
		}
	}

	if p.Pointer {
		ret += strings.Repeat("*", p.PointerCount)
	} else if p.ByRef {
		ret += "*"
	}

	return ret
}

// RenderTypeQtCpp renders the Qt C++ type in the original form from the function
// definition, before any typedefs or transformations were applied.
func (p CppParameter) RenderTypeQtCpp() string {
	return p.GetQtCppType().RenderTypeIntermediateCpp()
}

// RenderTypeIntermediateCpp renders the Qt C++ type WITHOUT resolving the
// interior QtCppOriginalType. This is used for intermediate const_cast<>s.
func (p CppParameter) RenderTypeIntermediateCpp() string {
	cppType := p.ParameterType

	if p.Const {
		cppType = "const " + cppType
	}
	if p.Pointer {
		cppType += strings.Repeat("*", p.PointerCount)
	}
	if p.ByRef {
		cppType += "&"
	}

	return cppType
}

// emitParametersCpp emits the parameter definitions exactly how Qt C++ defines them.
func emitParametersCpp(m CppMethod) string {
	tmp := make([]string, 0, len(m.Parameters))
	for _, p := range m.Parameters {
		tmp = append(tmp, p.RenderTypeQtCpp()+" "+p.cParameterName())
	}

	return strings.Join(tmp, `, `)
}

func emitParametersCppQualified(m CppMethod, className string) string {
	tmp := make([]string, 0, len(m.Parameters))
	for _, p := range m.Parameters {
		rendered := p.RenderTypeQtCpp()
		if needsClassQualification(p, className) {
			rendered = className + "::" + rendered
		}
		tmp = append(tmp, rendered+" "+p.cParameterName())
	}

	return strings.Join(tmp, `, `)
}

func needsClassQualification(p CppParameter, className string) bool {
	if p.QtClassType() || p.Pointer || p.ByRef || p.Const {
		return false
	}
	pt := p.ParameterType
	if p.GetQtCppType() != nil && p.GetQtCppType().ParameterType != pt {
		pt = p.GetQtCppType().ParameterType
	}
	if pt == "" || len(pt) == 0 {
		return false
	}
	if pt[0] == 'Q' || strings.Contains(pt, "::") || strings.Contains(pt, "<") {
		return false
	}
	switch pt {
	case "int", "unsigned int", "uint", "short", "unsigned short", "ushort",
		"long", "unsigned long", "ulong", "long long", "unsigned long long",
		"longlong", "ulonglong", "qlonglong", "qulonglong",
		"qint8", "quint8", "qint16", "quint16", "qint32", "quint32", "qint64", "quint64",
		"int8_t", "uint8_t", "int16_t", "uint16_t", "int32_t", "uint32_t", "int64_t", "uint64_t",
		"double", "float", "qreal", "bool", "char", "unsigned char", "signed char",
		"uchar", "void", "size_t", "qsizetype", "qintptr", "quintptr", "uintptr_t", "intptr_t",
		"qptrdiff", "ptrdiff_t", "QString", "QByteArray", "QVariant",
		"std::string", "std::chrono::milliseconds", "std::chrono::days", "std::chrono::nanoseconds",
		"GLenum", "GLuint", "GLint", "GLfloat", "GLdouble", "GLsizei", "GLboolean",
		"GLbitfield", "GLuint64", "GLsync", "GLubyte", "GLushort", "GLshort", "GLchar":
		return false
	}
	if p.IsKnownEnum() || p.IsFlagType() {
		return true
	}
	if KnownEnums.Has(className + "::" + pt) {
		return true
	}
	return false
}

func emitParameterNames(m CppMethod) string {
	tmp := make([]string, 0, len(m.Parameters))
	for _, p := range m.Parameters {
		tmp = append(tmp, p.cParameterName())
	}

	return strings.Join(tmp, `, `)
}

func emitParameterTypesCpp(m CppMethod, includeHidden bool) string {
	return emitParameterTypesCppQualified(m, "", includeHidden)
}

func emitParameterTypesCppQualified(m CppMethod, className string, includeHidden bool) string {
	tmp := make([]string, 0, len(m.Parameters))
	for _, p := range m.Parameters {
		rendered := p.RenderTypeQtCpp()
		if className != "" && needsClassQualification(p, className) {
			rendered = className + "::" + rendered
		}
		tmp = append(tmp, rendered)
	}
	if includeHidden {
		for _, p := range m.HiddenParams {
			rendered := p.RenderTypeQtCpp()
			if className != "" && needsClassQualification(p, className) {
				rendered = className + "::" + rendered
			}
			tmp = append(tmp, rendered)
		}
	}

	return strings.Join(tmp, `, `)
}

func EmitParametersCabi(m CppMethod, selfType string) string {
	tmp := make([]string, 0, len(m.Parameters)+1)

	if !m.IsStatic && selfType != "" {
		tmp = append(tmp, selfType+" self")
	}

	for _, p := range m.Parameters {
		if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
			continue
		}
		tmp = append(tmp, p.RenderTypeCabi()+" "+p.cParameterName())
	}

	return strings.Join(tmp, ", ")
}

func EmitParametersCABI2CppForwarding(params []CppParameter, indent string, className string) (preamble string, forwarding string) {
	tmp := make([]string, 0, len(params)+1)

	for _, p := range params {
		if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
			continue
		}
		addPre, addFwd := EmitCABI2CppForwarding(p, indent, className)
		preamble += addPre
		tmp = append(tmp, addFwd)
	}

	return preamble, strings.Join(tmp, ", ")
}

func MakeNamePrefix(in string) string {
	replacer := strings.NewReplacer(`[`, `_`, `]`, "", `.`, `_`)
	return replacer.Replace(in)
}

func qualifyInnerTypes(typeStr string, className string) string {
	re := regexp.MustCompile(`(\w+)<([^>]+)>`)
	parentScope := className
	if lastColon := strings.LastIndex(className, "::"); lastColon >= 0 {
		parentScope = className[:lastColon]
	}
	return re.ReplaceAllStringFunc(typeStr, func(match string) string {
		sub := re.FindStringSubmatch(match)
		container := sub[1]
		inner := strings.TrimSpace(sub[2])
		parts := strings.Split(inner, ",")
		for i, part := range parts {
			part = strings.TrimSpace(part)
			if en, ok := KnownEnums.Get(part); ok {
				part = en.Enum.EnumName
			} else if len(part) > 0 && part[0] != 'Q' && !strings.Contains(part, "::") && !isCppPrimitiveType(part) {
				part = parentScope + "::" + part
			}
			parts[i] = part
		}
		return container + "<" + strings.Join(parts, ", ") + ">"
	})
}

func isCppPrimitiveType(t string) bool {
	switch t {
	case "int", "unsigned int", "uint", "short", "unsigned short", "ushort",
		"long", "unsigned long", "ulong", "long long", "unsigned long long",
		"longlong", "ulonglong", "qlonglong", "qulonglong",
		"qint8", "quint8", "qint16", "quint16", "qint32", "quint32", "qint64", "quint64",
		"int8_t", "uint8_t", "int16_t", "uint16_t", "int32_t", "uint32_t", "int64_t", "uint64_t",
		"double", "float", "qreal", "bool", "char", "unsigned char", "signed char",
		"uchar", "void", "size_t", "qsizetype", "qintptr", "quintptr", "uintptr_t", "intptr_t",
		"qptrdiff", "ptrdiff_t", "QString", "QByteArray", "QVariant",
		"GLenum", "GLuint", "GLint", "GLfloat", "GLdouble", "GLsizei", "GLboolean",
		"GLbitfield", "GLuint64", "GLsync", "GLubyte", "GLushort", "GLshort", "GLchar":
		return true
	default:
		return false
	}
}

func EmitCABI2CppForwarding(p CppParameter, indent string, className string) (preamble string, forwarding string) {

	nameprefix := MakeNamePrefix(p.cParameterName())

	// For inner classes (className contains ::), qualify any unqualified
	// template parameters that match the inner class short name
	qualifiedType := p.ParameterType
	if strings.Contains(className, "::") && strings.Contains(p.ParameterType, "<") {
		qualifiedType = qualifyInnerTypes(p.ParameterType, className)
	}

	// Variadic ... cannot be forwarded via static_cast
	if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
		return preamble, p.cParameterName()
	}

	switch p.ParameterType {
	case "QString":
		preamble += indent + "QString " + nameprefix + "_QString = QString::fromUtf8(" + p.cParameterName() + ".data, " + p.cParameterName() + ".len);\n"
		return preamble, nameprefix + "_QString"

	case "QByteArray":
		preamble += indent + "QByteArray " + nameprefix + "_QByteArray(" + p.cParameterName() + ".data, " + p.cParameterName() + ".len);\n"
		return preamble, nameprefix + "_QByteArray"

	case "std::chrono::milliseconds":
		return preamble, "std::chrono::milliseconds(" + p.cParameterName() + ")"
	case "std::chrono::days":
		return preamble, "std::chrono::days(" + p.cParameterName() + ")"
	case "std::chrono::nanoseconds":
		return preamble, "std::chrono::nanoseconds(" + p.cParameterName() + ")"
	case "duration<long long, milli>":
		return preamble, "std::chrono::milliseconds(" + p.cParameterName() + ")"
	case "duration<long long, nano>":
		return preamble, "std::chrono::nanoseconds(" + p.cParameterName() + ")"
	case "duration<long long>":
		return preamble, "std::chrono::seconds(" + p.cParameterName() + ")"
	}

	if listType, _, ok := p.QListOf(); ok {

		listCppType := p.GetQtCppType().ParameterType
		if className != "" && !strings.Contains(listCppType, "::") {
			listCppType = qualifyInnerTypes(listCppType, className)
		}
		preamble += indent + listCppType + " " + nameprefix + "_QList;\n"
		preamble += indent + nameprefix + "_QList.reserve(" + p.cParameterName() + ".len);\n"

		preamble += indent + listType.RenderTypeCabi() + "* " + nameprefix + "_arr = static_cast<" + listType.RenderTypeCabi() + "*>(" + p.cParameterName() + ".data);\n"
		preamble += indent + "for(size_t i = 0; i < " + p.cParameterName() + ".len; ++i) {\n"

		listType.ParameterName = nameprefix + "_arr[i]"
		addPre, addFwd := EmitCABI2CppForwarding(listType, indent+"\t", className)
		preamble += addPre
		preamble += indent + "\t" + nameprefix + "_QList.push_back(" + addFwd + ");\n"

		preamble += indent + "}\n"

		// Support passing QList<>* (very rare, but used in qnetwork)
		if p.Pointer {
			return preamble, "&" + nameprefix + "_QList"
		} else {
			return preamble, nameprefix + "_QList"
		}

	} else if kType, vType, mapContainerType, ok := p.QMapOf(); ok {
		preamble += indent + p.GetQtCppType().ParameterType + " " + nameprefix + "_QMap;\n"

		// This container may be a QMap or a QHash
		// QHash supports .reserve(), but QMap doesn't
		if mapContainerType == "QHash" {
			preamble += indent + nameprefix + "_QMap.reserve(" + p.cParameterName() + ".len);\n"
		}

		preamble += indent + kType.RenderTypeCabi() + "* " + nameprefix + "_karr = static_cast<" + kType.RenderTypeCabi() + "*>(" + p.cParameterName() + ".keys);\n"
		preamble += indent + vType.RenderTypeCabi() + "* " + nameprefix + "_varr = static_cast<" + vType.RenderTypeCabi() + "*>(" + p.cParameterName() + ".values);\n"
		preamble += indent + "for(size_t i = 0; i < " + p.cParameterName() + ".len; ++i) {\n"

		kType.ParameterName = nameprefix + "_karr[i]"
		addPreK, addFwdK := EmitCABI2CppForwarding(kType, indent+"\t", className)
		preamble += addPreK

		vType.ParameterName = nameprefix + "_varr[i]"
		addPreV, addFwdV := EmitCABI2CppForwarding(vType, indent+"\t", className)
		preamble += addPreV

		preamble += indent + "\t" + nameprefix + "_QMap[" + addFwdK + "] = " + addFwdV + ";\n"

		preamble += indent + "}\n"
		return preamble, nameprefix + "_QMap"

	} else if kType, vType, ok := p.QPairOf(); ok {
		preamble += indent + p.GetQtCppType().ParameterType + " " + nameprefix + "_QPair;\n"

		preamble += indent + kType.RenderTypeCabi() + "* " + nameprefix + "_first_arr = static_cast<" + kType.RenderTypeCabi() + "*>(" + p.cParameterName() + ".keys);\n"
		preamble += indent + vType.RenderTypeCabi() + "* " + nameprefix + "_second_arr = static_cast<" + vType.RenderTypeCabi() + "*>(" + p.cParameterName() + ".values);\n"

		kType.ParameterName = nameprefix + "_first_arr[0]"
		addPreK, addFwdK := EmitCABI2CppForwarding(kType, indent+"\t", className)
		preamble += addPreK

		vType.ParameterName = nameprefix + "_second_arr[0]"
		addPreV, addFwdV := EmitCABI2CppForwarding(vType, indent+"\t", className)
		preamble += addPreV

		preamble += indent + nameprefix + "_QPair.first = " + addFwdK + ";\n"
		preamble += indent + nameprefix + "_QPair.second = " + addFwdV + ";\n"

		return preamble, nameprefix + "_QPair"

	} else if p.IsFunctionPointerTypedef() {
		if p.Pointer {
			return preamble, "reinterpret_cast<" + p.ParameterType + "*>(" + p.cParameterName() + ")"
		}
		return preamble, "(" + p.ParameterType + ")(" + p.cParameterName() + ")"

	} else if !p.Pointer && (p.IsFlagType() || p.IntType() || p.IsKnownEnum()) {
		castSrc := p.cParameterName()
		castType := p.ParameterType
		if p.IsKnownEnum() && !p.IsFlagType() {
			resolved := p.ResolvedEnumName()
			if className != "" && !strings.HasPrefix(resolved, className+"::") {
				if KnownEnums.Has(className + "::" + p.ParameterType) {
					castType = className + "::" + p.ParameterType
				} else {
					castType = resolved
				}
			} else {
				castType = resolved
			}
		}
		if p.IsFlagType() {
			castType = p.ParameterType
			if inner, ok := p.QFlagsOf(); ok && !strings.Contains(inner.ParameterType, "::") {
				qualifiedInner := ""
				if className != "" {
					if KnownEnums.Has(className + "::" + inner.ParameterType) {
						qualifiedInner = className + "::" + inner.ParameterType
					}
				}
				if qualifiedInner == "" {
					if en, ok2 := KnownEnums.Get(inner.ParameterType); ok2 {
						qualifiedInner = en.Enum.EnumName
					} else {
						qualifiedInner = "Qt::" + inner.ParameterType
					}
				}
				castType = "QFlags<" + qualifiedInner + ">"
			}
		}

		if p.ByRef { // e.g. QDataStream::operator>>() overloads
			castSrc = "*" + castSrc
		}

		if p.QtCppOriginalType != nil && p.QtCppOriginalType.Const != p.Const {
			return preamble, "static_cast<" + p.RenderTypeQtCpp() + ">(const_cast<" + p.RenderTypeIntermediateCpp() + ">(" + p.cParameterName() + "))"
		}

		if p.GetQtCppType().ParameterType == "qintptr" ||
			p.GetQtCppType().ParameterType == "qsizetype" || // Qt 6 qversionnumber.h: invalid ‘static_cast�?from type ‘ptrdiff_t*�?{aka ‘long int*’} to type ‘qsizetype*�?{aka ‘long long int*’}
			p.ParameterType == "qint8" ||
			(p.IsFlagType() && p.ByRef) ||
			(p.IsKnownEnum() && p.ByRef) ||
			(p.IntType() && p.ByRef) {
			// QDataStream::operator>>() by reference (qint64)
			// QLockFile::getLockInfo() by pointer
			// QTextStream::operator>>() by reference (qlonglong + qulonglong)
			// QDataStream::operator>>() qint8
			// CABI has these as int64_t* (long int) which fails a static_cast to qint64& (long long int&)
			// Hack a hard C-style cast
			if (p.GetQtCppType().ParameterType == "qintptr" || p.GetQtCppType().ParameterType == "qsizetype") && (p.Pointer || p.ByRef) {
				return preamble, "(" + castType + "*)(" + p.cParameterName() + ")"
			}
			if p.IntType() && p.ByRef {
				return preamble, castSrc
			}
			return preamble, "(" + castType + ")(" + castSrc + ")"
		} else {
			// Use static_cast<> safely
			return preamble, "static_cast<" + castType + ">(" + castSrc + ")"
		}

	} else if _, ok := p.QSetOf(); ok {
		preamble += indent + p.GetQtCppType().ParameterType + " " + nameprefix + "_QSet;\n"

		setType, _ := p.QSetOf()
		preamble += indent + setType.RenderTypeCabi() + "* " + nameprefix + "_arr = static_cast<" + setType.RenderTypeCabi() + "*>(" + p.cParameterName() + ".data);\n"
		preamble += indent + "for(size_t i = 0; i < " + p.cParameterName() + ".len; ++i) {\n"

		setType.ParameterName = nameprefix + "_arr[i]"
		addPre, addFwd := EmitCABI2CppForwarding(setType, indent+"\t", className)
		preamble += addPre
		preamble += indent + "\t" + nameprefix + "_QSet.insert(" + addFwd + ");\n"

		preamble += indent + "}\n"

		if p.Pointer {
			return preamble, "&" + nameprefix + "_QSet"
		} else {
			return preamble, nameprefix + "_QSet"
		}

	} else if p.ByRef {
		if p.Pointer {
			// By ref and by pointer
			// This happens for QDataStream &QDataStream::operator>>(char *&s)
			// We are only using one level of indirection
			return preamble, p.cParameterName()
		} else {
			// By ref and not by pointer
			// We changed RenderTypeCabi() to render this as a pointer
			// Need to dereference so we can pass as reference to the actual Qt C++ function
			//tmp = append(tmp, "*"+p.cParameterName())
			return preamble, "*" + p.cParameterName()
		}

	} else if p.QtClassType() && !p.Pointer {
		if strings.Contains(p.ParameterType, "<") {
			return preamble, "reinterpret_cast<" + p.ParameterType + ">(" + p.cParameterName() + ")"
		}
		if strings.Contains(p.cParameterName(), `[`) {
			return preamble, "*(" + p.cParameterName() + ")" // Extra brackets aren't necessary, just nice
		}
		return preamble, "*" + p.cParameterName()

	} else {
		cabiType := p.RenderTypeCabi()

		// When QtPrivate::* type is mapped to void* in C ABI, need reinterpret_cast
		if cabiType == "void*" && strings.Contains(p.ParameterType, "QtPrivate") {
			return preamble, "reinterpret_cast<" + p.ParameterType + ">(" + p.cParameterName() + ")"
		}

		// When a function pointer type is mapped to void* in C ABI, need reinterpret_cast
		if cabiType == "void*" && strings.Contains(p.ParameterType, "(*)") {
			return preamble, "reinterpret_cast<" + p.ParameterType + ">(" + p.cParameterName() + ")"
		}

		// When CABI type is void* and C++ type is a pointer to an unknown class/struct,
		// need static_cast from void* to the actual pointer type
		// For template types, use reinterpret_cast to avoid issues with
		// unqualified template parameters (e.g. ImVector<InnerClass> where InnerClass
		// needs qualification in the outer context)
		if cabiType == "void*" && p.Pointer && !p.ByRef {
			if strings.Contains(qualifiedType, "<") {
				return preamble, "reinterpret_cast<" + qualifiedType + "*>(" + p.cParameterName() + ")"
			}
			return preamble, "static_cast<" + qualifiedType + "*>(" + p.cParameterName() + ")"
		}

		// Windows handle types that are pointer typedefs mapped to void*
		if cabiType == "void*" && !p.Pointer && !p.ByRef {
			switch p.ParameterType {
			case "HWND", "HINSTANCE", "HDC", "HGLRC", "HMENU", "HICON", "HCURSOR", "HBRUSH", "HRGN", "HMONITOR":
				return preamble, "reinterpret_cast<" + p.ParameterType + ">(" + p.cParameterName() + ")"
			}
		}

		// When a template or struct value type is mapped to void* in C ABI,
		// the CABI parameter is a pointer to the value - dereference it
		// For template types, use reinterpret_cast to avoid issues with
		// unqualified template parameters
		if cabiType == "void*" && !p.Pointer && !p.ByRef &&
			(strings.Contains(qualifiedType, "<") || isLikelyStructType(qualifiedType)) {
			if strings.Contains(qualifiedType, "<") {
				return preamble, "*reinterpret_cast<" + qualifiedType + "*>(" + p.cParameterName() + ")"
			}
			return preamble, "*static_cast<" + qualifiedType + "*>(" + p.cParameterName() + ")"
		}

		if isCabiIntegerType(cabiType) && p.ParameterType != cabiType &&
			!strings.HasPrefix(p.ParameterType, "nullptr") &&
			!strings.HasPrefix(p.ParameterType, "std::nullptr") &&
			!strings.HasPrefix(p.ParameterType, "decltype") &&
			!strings.Contains(cabiType, "*") {
			castType := p.ParameterType
			if qtType := p.GetQtCppType(); qtType.ParameterType != p.ParameterType && qtType.ParameterType != "" {
				castType = qtType.ParameterType
			} else if en, ok := KnownEnums.Get(p.ParameterType); ok {
				castType = en.Enum.EnumName
			} else if className != "" && !strings.Contains(p.ParameterType, "::") {
				castType = className + "::" + p.ParameterType
			}
			return preamble, "static_cast<" + castType + ">(" + p.cParameterName() + ")"
		}
		return preamble, p.cParameterName()
	}

}

// EmitAssignCppToCabi transforms and assigns rvalue to the assignExpression.
// Sample assignExpression: `return `, `auto foo = `
// Sample rvalue: `foo`, `foo(xyz)`
// The return is a complete statement including trailing newline.
func EmitAssignCppToCabi(assignExpression string, p CppParameter, rvalue string, className ...string) string {

	cn := ""
	if len(className) > 0 {
		cn = className[0]
	}

	// For inner classes (cn contains ::), qualify any unqualified
	// template parameters that match the inner class short name
	qualifiedType := p.ParameterType
	if strings.Contains(cn, "::") && strings.Contains(p.ParameterType, "<") {
		qualifiedType = qualifyInnerTypes(p.ParameterType, cn)
	}

	shouldReturn := assignExpression // n.b. already has indent
	afterCall := ""
	assignExpression = strings.TrimLeft(assignExpression, " \t")
	indent := shouldReturn[0 : len(shouldReturn)-len(assignExpression)]

	shouldReturn = shouldReturn[len(indent):]

	namePrefix := MakeNamePrefix(p.cParameterName())

	if p.Void() {
		shouldReturn = ""
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.ParameterType == "QString" {

		if p.Pointer {
			// e.g. QTextStream::String()
			// These are rare, and probably expected to be lightweight references
			// But, a copy is the best we can project it as
			// Un-pointer-ify
			shouldReturn = ifv(p.Const, "const ", "") + "QString* " + namePrefix + "_ret = "
			afterCall = indent + "// Convert QString pointer from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory\n"
			afterCall += indent + "QByteArray " + namePrefix + "_b = " + namePrefix + "_ret->toUtf8();\n"

		} else {
			shouldReturn = ifv(p.Const, "const ", "") + "QString " + namePrefix + "_ret = "
			afterCall = indent + "// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory\n"
			afterCall += indent + "QByteArray " + namePrefix + "_b = " + namePrefix + "_ret.toUtf8();\n"
		}

		afterCall += indent + "struct miqt_string " + namePrefix + "_ms;\n"
		afterCall += indent + namePrefix + "_ms.len = " + namePrefix + "_b.length();\n"
		afterCall += indent + namePrefix + "_ms.data = static_cast<char*>(malloc(" + namePrefix + "_ms.len));\n"
		afterCall += indent + "memcpy(" + namePrefix + "_ms.data, " + namePrefix + "_b.data(), " + namePrefix + "_ms.len);\n"
		afterCall += indent + assignExpression + namePrefix + "_ms;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.ParameterType == "QByteArray" {
		// C++ has given us a QByteArray. CABI needs this as a struct miqt_string
		// Do not free the data, the caller will free it

		shouldReturn = ifv(p.Const, "const ", "") + "QByteArray " + namePrefix + "_qb = "

		afterCall += indent + "struct miqt_string " + namePrefix + "_ms;\n"
		afterCall += indent + namePrefix + "_ms.len = " + namePrefix + "_qb.length();\n"
		afterCall += indent + namePrefix + "_ms.data = static_cast<char*>(malloc(" + namePrefix + "_ms.len));\n"
		afterCall += indent + "memcpy(" + namePrefix + "_ms.data, " + namePrefix + "_qb.data(), " + namePrefix + "_ms.len);\n"
		afterCall += indent + assignExpression + namePrefix + "_ms;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if t, _, ok := p.QListOf(); ok {

		// In some cases rvalue is a function call and the temporary
		// is necessary; in some cases it's a literal and the temporary is
		// elided; but in some cases it's a Qt class and the temporary goes
		// through a copy constructor
		// TODO Detect safe cases where this can be optimized

		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "
		if len(className) > 0 && className[0] != "" {
			qType := qualifyInnerTypes(p.RenderTypeQtCpp(), className[0])
			if qType != p.RenderTypeQtCpp() {
				shouldReturn = qType + " " + namePrefix + "_ret = "
			}
		}

		afterCall += indent + "// Convert QList<> from C++ memory to manually-managed C memory\n"
		afterCall += indent + "" + t.RenderTypeCabi() + "* " + namePrefix + "_arr = static_cast<" + t.RenderTypeCabi() + "*>(malloc(sizeof(" + t.RenderTypeCabi() + ") * " + namePrefix + "_ret.length()));\n"
		afterCall += indent + "for (size_t i = 0, e = " + namePrefix + "_ret.length(); i < e; ++i) {\n"
		afterCall += EmitAssignCppToCabi(indent+"\t"+namePrefix+"_arr[i] = ", t, namePrefix+"_ret[i]")
		afterCall += indent + "}\n"

		afterCall += indent + "struct miqt_array " + namePrefix + "_out;\n"
		afterCall += indent + "" + namePrefix + "_out.len = " + namePrefix + "_ret.length();\n"
		afterCall += indent + "" + namePrefix + "_out.data = static_cast<void*>(" + namePrefix + "_arr);\n"

		afterCall += indent + assignExpression + "" + namePrefix + "_out;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if t, ok := p.QSetOf(); ok {

		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "
		if len(className) > 0 && className[0] != "" {
			qType := qualifyInnerTypes(p.RenderTypeQtCpp(), className[0])
			if qType != p.RenderTypeQtCpp() {
				shouldReturn = qType + " " + namePrefix + "_ret = "
			}
		}

		afterCall += indent + "// Convert QSet<> from C++ memory to manually-managed C memory\n"
		afterCall += indent + "" + t.RenderTypeCabi() + "* " + namePrefix + "_arr = static_cast<" + t.RenderTypeCabi() + "*>(malloc(sizeof(" + t.RenderTypeCabi() + ") * " + namePrefix + "_ret.size()));\n"
		afterCall += indent + "int " + namePrefix + "_ctr = 0;\n"
		afterCall += indent + "QSetIterator<" + t.RenderTypeQtCpp() + "> " + namePrefix + "_itr(" + namePrefix + "_ret);\n"
		afterCall += indent + "while (" + namePrefix + "_itr.hasNext()) {\n"
		afterCall += EmitAssignCppToCabi(indent+"\t"+namePrefix+"_arr["+namePrefix+"_ctr++] = ", t, namePrefix+"_itr.next()")
		afterCall += indent + "}\n"

		afterCall += indent + "struct miqt_array " + namePrefix + "_out;\n"
		afterCall += indent + "" + namePrefix + "_out.len = " + namePrefix + "_ret.size();\n"
		afterCall += indent + "" + namePrefix + "_out.data = static_cast<void*>(" + namePrefix + "_arr);\n"

		afterCall += indent + assignExpression + "" + namePrefix + "_out;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if kType, vType, _, ok := p.QMapOf(); ok {
		// QMap<K,V>

		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "
		if len(className) > 0 && className[0] != "" {
			qType := qualifyInnerTypes(p.RenderTypeQtCpp(), className[0])
			if qType != p.RenderTypeQtCpp() {
				shouldReturn = qType + " " + namePrefix + "_ret = "
			}
		}

		afterCall += indent + "// Convert QMap<> from C++ memory to manually-managed C memory\n"
		afterCall += indent + "" + kType.RenderTypeCabi() + "* " + namePrefix + "_karr = static_cast<" + kType.RenderTypeCabi() + "*>(malloc(sizeof(" + kType.RenderTypeCabi() + ") * " + namePrefix + "_ret.size()));\n"
		afterCall += indent + "" + vType.RenderTypeCabi() + "* " + namePrefix + "_varr = static_cast<" + vType.RenderTypeCabi() + "*>(malloc(sizeof(" + vType.RenderTypeCabi() + ") * " + namePrefix + "_ret.size()));\n"

		afterCall += indent + "int " + namePrefix + "_ctr = 0;\n"
		afterCall += indent + "for (auto " + namePrefix + "_itr = " + namePrefix + "_ret.keyValueBegin(); " + namePrefix + "_itr != " + namePrefix + "_ret.keyValueEnd(); ++" + namePrefix + "_itr) {\n"
		afterCall += EmitAssignCppToCabi(indent+"\t"+namePrefix+"_karr["+namePrefix+"_ctr] = ", kType, namePrefix+"_itr->first")
		afterCall += EmitAssignCppToCabi(indent+"\t"+namePrefix+"_varr["+namePrefix+"_ctr] = ", vType, namePrefix+"_itr->second")
		afterCall += indent + "\t" + namePrefix + "_ctr++;\n"

		afterCall += indent + "}\n"

		afterCall += indent + "struct miqt_map " + namePrefix + "_out;\n"
		afterCall += indent + "" + namePrefix + "_out.len = " + namePrefix + "_ret.size();\n"
		afterCall += indent + "" + namePrefix + "_out.keys = static_cast<void*>(" + namePrefix + "_karr);\n"
		afterCall += indent + "" + namePrefix + "_out.values = static_cast<void*>(" + namePrefix + "_varr);\n"

		afterCall += indent + assignExpression + "" + namePrefix + "_out;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if kType, vType, ok := p.QPairOf(); ok {
		// QPair<T1,T2>

		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "

		afterCall += indent + "// Convert QPair<> from C++ memory to manually-managed C memory\n"
		afterCall += indent + "" + kType.RenderTypeCabi() + "* " + namePrefix + "_first_arr = static_cast<" + kType.RenderTypeCabi() + "*>(malloc(sizeof(" + kType.RenderTypeCabi() + ")));\n"
		afterCall += indent + "" + vType.RenderTypeCabi() + "* " + namePrefix + "_second_arr = static_cast<" + vType.RenderTypeCabi() + "*>(malloc(sizeof(" + vType.RenderTypeCabi() + ")));\n"

		afterCall += EmitAssignCppToCabi(indent+namePrefix+"_first_arr[0] = ", kType, namePrefix+"_ret.first")
		afterCall += EmitAssignCppToCabi(indent+namePrefix+"_second_arr[0] = ", vType, namePrefix+"_ret.second")

		afterCall += indent + "struct miqt_map " + namePrefix + "_out;\n"
		afterCall += indent + "" + namePrefix + "_out.len = 1;\n"
		afterCall += indent + "" + namePrefix + "_out.keys = static_cast<void*>(" + namePrefix + "_first_arr);\n"
		afterCall += indent + "" + namePrefix + "_out.values = static_cast<void*>(" + namePrefix + "_second_arr);\n"

		afterCall += indent + assignExpression + "" + namePrefix + "_out;\n"
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if strings.HasPrefix(p.ParameterType, "std::chrono::") || p.ParameterType == "duration<long long, milli>" || p.ParameterType == "duration<long long, nano>" || p.ParameterType == "duration<long long>" {
		return indent + p.RenderTypeCabi() + " " + namePrefix + "_ret = " + rvalue + ".count();\n" +
			indent + assignExpression + namePrefix + "_ret;\n"

	} else if p.IsFlagType() || p.IsKnownEnum() || p.QtCppOriginalType != nil {
		returnType := p.ParameterType
		if p.IsKnownEnum() && !p.IsFlagType() {
			resolved := p.ResolvedEnumName()
			if cn != "" && !strings.HasPrefix(resolved, cn+"::") {
				if KnownEnums.Has(cn + "::" + p.ParameterType) {
					returnType = cn + "::" + p.ParameterType
				} else {
					returnType = resolved
				}
			} else {
				returnType = resolved
			}
		}
		if p.IsFlagType() {
			origQtType := p.GetQtCppType().ParameterType
			if origQtType != "" && (strings.HasPrefix(origQtType, "QFlags<") || strings.Contains(origQtType, "::")) {
				returnType = origQtType
			}
			if inner, ok := p.QFlagsOf(); ok && !strings.Contains(inner.ParameterType, "::") {
				qualifiedInner := ""
				if cn != "" {
					if KnownEnums.Has(cn + "::" + inner.ParameterType) {
						qualifiedInner = cn + "::" + inner.ParameterType
					}
				}
				if qualifiedInner == "" {
					if en, ok2 := KnownEnums.Get(inner.ParameterType); ok2 {
						qualifiedInner = en.Enum.EnumName
					} else {
						qualifiedInner = "Qt::" + inner.ParameterType
					}
				}
				returnType = "QFlags<" + qualifiedInner + ">"
			}
		}
		if p.Const {
			returnType = "const " + returnType
		}
		if p.Pointer {
			returnType += strings.Repeat("*", p.PointerCount)
		}
		// Function pointer types can't be used directly as variable declarations
		// e.g. `void (*)(int) var = ...` is invalid C++; use `auto` instead
		if strings.Contains(returnType, "(*)") {
			shouldReturn = "auto " + namePrefix + "_ret = "
		} else {
			shouldReturn = returnType + " " + namePrefix + "_ret = "
		}
		if p.QtCppOriginalType != nil && p.QtCppOriginalType.Const != p.Const {
			afterCall += indent + "" + assignExpression + "const_cast<" + p.RenderTypeCabi() + ">(static_cast<" + p.RenderTypeIntermediateCpp() + ">(" + namePrefix + "_ret));\n"
		} else if p.QtCppOriginalType != nil && p.QtCppOriginalType.ParameterType == "qintptr" {
			afterCall += indent + "" + assignExpression + "(" + p.RenderTypeCabi() + ")(" + namePrefix + "_ret);\n"
		} else if p.ByRef {
			afterCall += indent + assignExpression + "reinterpret_cast<" + p.RenderTypeCabi() + ">(&" + namePrefix + "_ret);\n"
		} else {
			afterCall += indent + "" + assignExpression + "static_cast<" + p.RenderTypeCabi() + ">(" + namePrefix + "_ret);\n"
		}
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.QtClassType() && p.ByRef {
		// It's a pointer in disguise, just needs one cast
		shouldReturn = p.RenderTypeQtCpp() + " " + namePrefix + "_ret = "
		afterCall += indent + "// Cast returned reference into pointer\n"
		if p.Const {
			nonConst := p // copy
			nonConst.Const = false
			nonConst.ByRef = false
			nonConst.Pointer = true
			nonConst.PointerCount = 1
			afterCall += indent + "" + assignExpression + "const_cast<" + nonConst.RenderTypeQtCpp() + ">(&" + namePrefix + "_ret);\n"
		} else {
			afterCall += indent + "" + assignExpression + "&" + namePrefix + "_ret;\n"
		}
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if p.QtClassType() && !p.Pointer {
		if strings.Contains(p.ParameterType, "<") {
			return indent + assignExpression + "reinterpret_cast<" + p.RenderTypeCabi() + ">(new " + qualifiedType + "(" + rvalue + "));\n"
		}
		return indent + assignExpression + "new " + qualifiedType + "(" + rvalue + ");\n"

	} else if p.Const {
		shouldReturn += "(" + p.RenderTypeCabi() + ") "
		return indent + shouldReturn + rvalue + ";\n" + afterCall

	} else if strings.Contains(p.RenderTypeCabi(), "*") && !p.Pointer && !p.ByRef {
		// CABI type is a pointer (e.g. void* for iterators) but C++ returns by value
		if strings.Contains(p.ParameterType, "<") || isLikelyStructType(p.ParameterType) {
			// Template or struct value type - take address instead of reinterpret_cast
			// Use const_cast + static_cast to safely handle both const and non-const cases
			return indent + assignExpression + "const_cast<" + p.RenderTypeCabi() + ">(static_cast<const void*>(&" + rvalue + "));\n"
		}
		return indent + assignExpression + "reinterpret_cast<" + p.RenderTypeCabi() + ">(" + rvalue + ");\n"

	} else {
		// Basic type
		if p.ByRef {
			// The C++ type is a reference, the CABI type is a pointer type
			shouldReturn += "&"
		}

		return indent + shouldReturn + rvalue + ";\n" + afterCall
	}

}

func GetCppZeroValue(p CppParameter) string {
	if p.Pointer {
		return "nullptr"
	} else if p.IsKnownEnum() {
		return "(" + p.RenderTypeQtCpp() + ")(0)"
	} else if p.IntType() {
		return "0"
	} else if p.ParameterType == "bool" {
		return "false"
	} else if p.ParameterType == "void" {
		return ""
	} else {
		return p.RenderTypeQtCpp() + "()"
	}
}

func GetCabiZeroValue(p CppParameter) string {
	// n.b. Identical to getCppZeroValue in most cases

	if p.Pointer && !(p.ParameterType == "QString") {
		return GetCppZeroValue(p)

	} else if ev, ok := KnownEnums.Get(p.ParameterType); ok {
		// In CABI the zero value may be the underlying type of an enum instead
		return "(" + ev.Enum.UnderlyingType.RenderTypeCabi() + ")(0)"

	} else if p.IntType() {
		return GetCppZeroValue(p) // default

	} else if p.ParameterType == "bool" {
		return GetCppZeroValue(p)
	} else if p.ParameterType == "void" {
		return GetCppZeroValue(p)

	} else if p.ParameterType == "QString" || p.ParameterType == "QByteArray" {
		return "{}"

	} else if _, _, ok := p.QListOf(); ok {
		return "{}"

	} else if _, ok := p.QSetOf(); ok {
		return "{}"

	} else if _, _, _, ok := p.QMapOf(); ok {
		return "{}"

	} else if _, _, ok := p.QPairOf(); ok {
		return "{}"

	} else if p.RenderTypeCabi() == "int" {
		return "0"

	} else {
		// Difference for Qt classes: Qt C++ can expect to return them by value,
		// but CABI always needs to return them by pointer

		return "nullptr"
	}

}

// getReferencedTypes finds all referenced Qt types in this file.
func getReferencedTypes(src *CppParsedHeader) []string {

	foundTypes := map[string]struct{}{}

	var maybeAddType func(p CppParameter)
	maybeAddType = func(p CppParameter) {
		if p.QtClassType() {
			foundTypes[p.ParameterType] = struct{}{}
		}
		if p.ByRef || p.Pointer {
			if !p.IsKnownEnum() {
				if !KnownTypedefs.Has(p.ParameterType) {
					if p.ParameterType != "" && p.ParameterType != "void" {
						foundTypes[p.ParameterType] = struct{}{}
					}
				}
			}
		}
		if t, containerType, ok := p.QListOf(); ok {
			foundTypes[containerType] = struct{}{} // QList / QVector
			maybeAddType(t)
		}
		if kType, vType, containerType, ok := p.QMapOf(); ok {
			foundTypes[containerType] = struct{}{} // QMap / QHash
			maybeAddType(kType)
			maybeAddType(vType)
		}
		if kType, vType, ok := p.QPairOf(); ok {
			foundTypes["QPair"] = struct{}{}
			maybeAddType(kType)
			maybeAddType(vType)
		}
		if t, ok := p.QSetOf(); ok {
			foundTypes["QSet"] = struct{}{}
			maybeAddType(t)
		}
	}

	for _, c := range src.Classes {

		foundTypes[c.ClassName] = struct{}{}

		for _, ctor := range c.Ctors {
			for _, p := range ctor.Parameters {
				maybeAddType(p)
			}
		}
		for _, m := range c.Methods {
			for _, p := range m.Parameters {
				maybeAddType(p)
			}
			maybeAddType(m.ReturnType)
		}
		for _, vm := range c.VirtualMethods() {
			for _, p := range vm.Parameters {
				maybeAddType(p)
			}
			maybeAddType(vm.ReturnType)
		}
		for _, vm := range c.ProtectedMethods() {
			for _, p := range vm.Parameters {
				maybeAddType(p)
			}
			maybeAddType(vm.ReturnType)
		}
		for _, cn := range c.AllInheritsClassInfo() {
			maybeAddType(CppParameter{
				ParameterType: cn.Class.ClassName,
			})
		}
	}

	for _, fn := range src.FreeFunctions {
		maybeAddType(fn.ReturnType)
		for _, p := range fn.Parameters {
			maybeAddType(p)
		}
	}

	// Some types (e.g. QRgb) are found but are typedefs, not classes
	for _, td := range src.Typedefs {
		delete(foundTypes, td.Alias)
	}

	// Convert to sorted list
	foundTypesList := make([]string, 0, len(foundTypes))
	for ft := range foundTypes {
		if !ActiveFilter.AllowClass(ft) {
			continue
		}
		if isCPrimitiveType(ft) {
			continue
		}
		if strings.Contains(ft, `(*)`) {
			continue
		}
		if strings.HasSuffix(ft, `()`) {
			continue
		}
		if strings.Contains(ft, `<`) {
			continue
		}

		foundTypesList = append(foundTypesList, ft)
	}
	sort.Strings(foundTypesList)

	return foundTypesList
}

// CabiClassName returns the Go / CABI class name for a Qt C++ class.
// Normally this is the same, except for class types that are nested inside another class definition.
func isCPrimitiveType(typeName string) bool {
	switch typeName {
	case "bool", "char", "signed char", "unsigned char",
		"short", "unsigned short", "int", "unsigned int",
		"long", "unsigned long", "long long", "unsigned long long",
		"float", "double", "long double",
		"void", "size_t", "ssize_t", "ptrdiff_t",
		"int8_t", "int16_t", "int32_t", "int64_t",
		"uint8_t", "uint16_t", "uint32_t", "uint64_t",
		"char16_t", "char32_t", "wchar_t",
		"qfloat16",
		"qint8", "qint16", "qint32", "qint64",
		"quint8", "quint16", "quint32", "quint64",
		"qlonglong", "qulonglong",
		"qintptr", "quintptr", "qsizetype", "qptrdiff",
		"qreal",
		"ImU8", "ImU16", "ImU32", "ImU64",
		"ImS8", "ImS16", "ImS32", "ImS64",
		"ImWchar", "ImWchar16", "ImWchar32":
		return true
	}
	return false
}

func CabiClassName(className string) string {

	// Many types are defined in qnamespace.h under Qt::
	// The Go implementation is always called qt.Foo, and these names don't
	// collide with anything, so strip the redundant prefix
	className = strings.TrimPrefix(className, `Qt::`)

	// Must use __ to avoid subclass/method name collision e.g. QPagedPaintDevice::Margins
	return strings.Replace(className, `::`, `__`, -1)
}

func (cfg EmitConfig) isValueType(className string) bool {
	cn := CabiClassName(className)
	_, ok := cfg.ValueTypes[cn]
	return ok
}

func (cfg EmitConfig) isAllowedEnum(enumName string) bool {
	if len(cfg.AllowedEnumPrefix) == 0 {
		return true
	}
	for _, prefix := range cfg.AllowedEnumPrefix {
		if strings.HasPrefix(enumName, prefix) {
			return true
		}
	}
	return false
}

func emitBindingHeader(src *CppParsedHeader, filename string, packageName string, cfg EmitConfig) (string, string, error) {
	cppheader := strings.Builder{}
	cheader := strings.Builder{}

	includeGuard := "MIQT_" + strings.ToUpper(strings.Replace(strings.Replace(packageName, `/`, `_`, -1), `-`, `_`, -1)) + "_GEN_" + strings.ToUpper(strings.Replace(strings.Replace(filename, `.`, `_`, -1), `-`, `_`, -1))
	cheaderIncludeGuard := "CABI_" + strings.ToUpper(strings.Replace(strings.Replace(packageName, `/`, `_`, -1), `-`, `_`, -1)) + "_GEN_" + strings.ToUpper(strings.Replace(strings.Replace(filename, `.`, `_`, -1), `-`, `_`, -1))

	bindingInclude := "libmiqt.h"
	exportInclude := "miqt_export.h"
	if cfg.SkipLibmiqtInclude {
		bindingInclude = ""
	}

	preamble := `#pragma once
#ifndef ` + includeGuard + `
#define ` + includeGuard + `

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>

`
	if bindingInclude != "" {
		preamble += `#include "` + bindingInclude + `"
`
	}
	preamble += `#include "` + exportInclude + `"

`
	cppheader.WriteString(preamble)

	cheaderPreamble := `#pragma once
#ifndef ` + cheaderIncludeGuard + `
#define ` + cheaderIncludeGuard + `

`
	if bindingInclude != "" {
		cheaderPreamble += `#include "` + bindingInclude + `"
#include "` + exportInclude + `"

`
	} else {
		cheaderPreamble += `#include "` + exportInclude + `"

`
	}
	cheader.WriteString(cheaderPreamble)

	foundTypesList := getReferencedTypes(src)

	cppheader.WriteString("#ifdef __cplusplus\n")

	for _, ft := range foundTypesList {
		if ActiveFilter.PreventStructDeclaration(ft) && !ActiveFilter.ShouldEmitForwardDeclaration(ft) {
			continue
		}

		if strings.Contains(ft, `<`) {
			continue
		}

		if cfg.isValueType(ft) || ActiveFilter.ShouldEmitStructDefinition(ft) {
			continue
		}

		if strings.Contains(ft, `::`) {
			cppheader.WriteString(`struct ` + CabiClassName(ft) + `;\n`)

		} else {
			cppheader.WriteString(`class ` + ft + ";\n")
		}
	}

	cppheader.WriteString("#else\n")

	for _, ft := range foundTypesList {
		if ActiveFilter.PreventStructDeclaration(ft) && !ActiveFilter.ShouldEmitForwardDeclaration(ft) {
			continue
		}
		if strings.Contains(ft, `<`) {
			continue
		}
		if cfg.isValueType(ft) || ActiveFilter.ShouldEmitStructDefinition(ft) {
			continue
		}
		cppheader.WriteString(`typedef struct ` + CabiClassName(ft) + " " + CabiClassName(ft) + ";\n")
	}

	cppheader.WriteString("#endif\n")

	cppheader.WriteString("#ifdef __cplusplus\nextern \"C\" {\n#endif\n")

	for _, ft := range foundTypesList {
		if ActiveFilter.PreventStructDeclaration(ft) && !ActiveFilter.ShouldEmitForwardDeclaration(ft) {
			continue
		}
		if strings.Contains(ft, `<`) {
			continue
		}
		if cfg.isValueType(ft) || ActiveFilter.ShouldEmitStructDefinition(ft) {
			continue
		}
		cheader.WriteString(`typedef struct ` + CabiClassName(ft) + " " + CabiClassName(ft) + ";\n")
	}

	cppheader.WriteString("\n")
	cheader.WriteString("\n")

	if cfg.HeaderVersionMacro != "" {
		cppheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n\n")
		cheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n\n")
	}

	if !cfg.SkipEnumGeneration {
		for _, e := range src.Enums {
			if e.EnumName == "" {
				continue
			}
			if !cfg.isAllowedEnum(e.EnumName) {
				continue
			}
			enumCabiName := CabiClassName(e.EnumName)
			cppheader.WriteString(fmt.Sprintf("// 来源: 枚举 (%s)\n", e.EnumName))
			enumSrc := "typedef enum {\n"
			for _, ee := range e.Entries {
				if ee.EntryValue != "" {
					enumSrc += "\t" + ee.EntryName + " = " + ee.EntryValue + ",\n"
				} else {
					enumSrc += "\t" + ee.EntryName + ",\n"
				}
			}
			enumSrc += "} " + enumCabiName + ";\n\n"
			cppheader.WriteString(enumSrc)
			cheader.WriteString(enumSrc)
		}
	}

	if cfg.HeaderVersionMacro != "" {
		cppheader.WriteString("#endif\n\n")
		cheader.WriteString("#endif\n\n")
	}

	var autoStructCppDefs []string
	var autoStructCDefs []string
	for _, c := range src.Classes {
		if ActiveFilter.ShouldEmitStructDefinition(c.ClassName) {
			if len(c.Props) == 0 {
				continue
			}
			cName := CabiClassName(c.ClassName)
			var cFields []string
			var cppFields []string
			for _, p := range c.Props {
				cFields = append(cFields, p.PropertyType+" "+p.PropertyName)
				cppFields = append(cppFields, p.PropertyType+" "+p.PropertyName+";")
			}
			autoStructCDefs = append(autoStructCDefs, "typedef struct "+cName+" { "+strings.Join(cFields, "; ")+"; } "+cName+";")
			if strings.Contains(c.ClassName, "::") {
				autoStructCppDefs = append(autoStructCppDefs, "typedef "+c.ClassName+" "+cName+";")
			} else {
				autoStructCppDefs = append(autoStructCppDefs, "struct "+cName+" { "+strings.Join(cppFields, " ")+" };")
			}
		}
	}
	if len(autoStructCDefs) > 0 {
		if cfg.HeaderVersionMacro != "" {
			cppheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n")
			cheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n")
		}
		cppheader.WriteString("#ifdef __cplusplus\n")
		for _, def := range autoStructCppDefs {
			cppheader.WriteString(def + "\n")
		}
		cppheader.WriteString("#else\n")
		for _, def := range autoStructCDefs {
			cppheader.WriteString(def + "\n")
		}
		cppheader.WriteString("#endif\n")
		for _, def := range autoStructCDefs {
			cheader.WriteString(def + "\n")
		}
		if cfg.HeaderVersionMacro != "" {
			cppheader.WriteString("#endif\n")
			cheader.WriteString("#endif\n")
		}
		cppheader.WriteString("\n")
		cheader.WriteString("\n")
	}

	if cfg.HeaderVersionMacro != "" && len(cfg.ValueTypes) > 0 {
		hasEmpty := false
		for _, def := range cfg.ValueTypes {
			if def == "" {
				hasEmpty = true
				break
			}
		}
		if hasEmpty {
			cppheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n")
			cheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n")
			for name, def := range cfg.ValueTypes {
				if def == "" {
					cppheader.WriteString("#ifdef __cplusplus\nclass " + name + ";\n#else\ntypedef struct " + name + " " + name + ";\n#endif\n")
					cheader.WriteString("typedef struct " + name + " " + name + ";\n")
				}
			}
			cppheader.WriteString("#endif\n")
			cheader.WriteString("#endif\n")
			cppheader.WriteString("\n")
			cheader.WriteString("\n")
		}
	}

	var autoTypedefs []string
	for _, td := range src.Typedefs {
		if ActiveFilter.ShouldEmitEnumDefinition(td.Alias) {
			if strings.Contains(td.UnderlyingType.ParameterType, "(*)") {
				continue
			}
			autoTypedefs = append(autoTypedefs, "typedef "+td.UnderlyingType.ParameterType+" "+td.Alias+";")
		}
	}
	if len(autoTypedefs) > 0 {
		if cfg.HeaderVersionMacro != "" {
			cppheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n")
			cheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n")
		}
		for _, def := range autoTypedefs {
			cppheader.WriteString(def + "\n")
			cheader.WriteString(def + "\n")
		}
		if cfg.HeaderVersionMacro != "" {
			cppheader.WriteString("#endif\n")
			cheader.WriteString("#endif\n")
		}
		cppheader.WriteString("\n")
		cheader.WriteString("\n")
	}

	if len(cfg.ValueTypes) > 0 {
		if cfg.HeaderVersionMacro != "" {
			cppheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n")
			cheader.WriteString("#ifndef " + cfg.HeaderVersionMacro + "\n")
		}
		cppheader.WriteString("#ifdef __cplusplus\n")
		var vtNames []string
		for name := range cfg.ValueTypes {
			vtNames = append(vtNames, name)
		}
		sort.Strings(vtNames)
		for _, name := range vtNames {
			def := cfg.ValueTypes[name]
			if def == "" {
				cppheader.WriteString("class " + name + ";\n")
				continue
			}
			cppDef := strings.TrimPrefix(def, "typedef ")
			cppDef = strings.TrimSuffix(cppDef, " "+name+";")
			cppDef += ";"
			cppheader.WriteString(cppDef + "\n")
		}
		cppheader.WriteString("#else\n")
		for _, name := range vtNames {
			def := cfg.ValueTypes[name]
			if def == "" {
				cppheader.WriteString("typedef struct " + name + " " + name + ";\n")
				continue
			}
			cppheader.WriteString(def + "\n")
		}
		cppheader.WriteString("#endif\n")
		for _, name := range vtNames {
			def := cfg.ValueTypes[name]
			if def == "" {
				cheader.WriteString("typedef struct " + name + " " + name + ";\n")
				continue
			}
			cheader.WriteString(def + "\n")
		}
		if cfg.HeaderVersionMacro != "" {
			cppheader.WriteString("#endif\n\n")
			cheader.WriteString("#endif\n\n")
		} else {
			cppheader.WriteString("\n")
			cheader.WriteString("\n")
		}
	}

	cppheader.WriteString("#ifndef MIQT_TYPES_ONLY\n\n")
	cheader.WriteString("#ifndef MIQT_TYPES_ONLY\n\n")

	for _, c := range src.Classes {
		className := CabiClassName(c.ClassName)
		cppheader.WriteString(fmt.Sprintf("// 来源: 类 (%s)\n", c.ClassName))
		virtualMethods := c.VirtualMethods()
		protectedMethods := c.ProtectedMethods()

		for i, ctor := range c.Ctors {
			hasVariadic := false
			for _, p := range ctor.Parameters {
				if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
					hasVariadic = true
					break
				}
			}
			if hasVariadic {
				continue
			}
			ctorLine := fmt.Sprintf("MIQT_EXPORT %s* %s(%s);\n", className, cabiNewName(c, i), emitParametersCabiConstructor(&c, &ctor))
			cppheader.WriteString("// 来源: 构造函数\n")
			cppheader.WriteString(ctorLine)
			cheader.WriteString(ctorLine)
		}

		if len(c.DirectInheritClassInfo()) > 0 {
			virtBaseLine := "MIQT_EXPORT void " + cabiVirtBaseName(c) + "(" + className + "* src"
			for _, baseClass := range c.DirectInheritClassInfo() {
				virtBaseLine += ", " + CabiClassName(baseClass.Class.ClassName) + "** outptr_" + CabiClassName(baseClass.Class.ClassName)
			}
			virtBaseLine += ");\n"
			cppheader.WriteString(virtBaseLine)
			cheader.WriteString(virtBaseLine)
		}

		for _, m := range c.Methods {
			if m.IsProtected && !m.IsVirtual {
				continue
			}
			hasVariadic := false
			for _, p := range m.Parameters {
				if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
					hasVariadic = true
					break
				}
			}
			if hasVariadic {
				continue
			}

			if m.ReturnType.BecomesConstInVersion != nil {
				versionLine := "// This method's return type was changed from non-const to const in Qt " + *m.ReturnType.BecomesConstInVersion + "\n" +
					"#if QT_VERSION >= QT_VERSION_CHECK(" + strings.Replace(*m.ReturnType.BecomesConstInVersion, `.`, `,`, -1) + ",0)\n" +
					fmt.Sprintf("MIQT_EXPORT %s %s(%s);\n", "const "+m.ReturnType.RenderTypeCabi(), cabiMethodName(c, m), EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*")) +
					"#else\n" +
					fmt.Sprintf("MIQT_EXPORT %s %s(%s);\n", m.ReturnType.RenderTypeCabi(), cabiMethodName(c, m), EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*")) +
					"#endif\n"
				cppheader.WriteString(versionLine)
				cheader.WriteString(versionLine)
			} else {
				isArrayField := m.IsVariable && ((strings.HasPrefix(m.MethodName, "set") && len(m.Parameters) > 0 && isArrayFieldType(m.Parameters[0])) || (!strings.HasPrefix(m.MethodName, "set") && isArrayFieldType(m.ReturnType)))
				if isArrayField {
					var elemType string
					if strings.HasPrefix(m.MethodName, "set") {
						elemType = arrayFieldElementType(m.Parameters[0])
					} else {
						elemType = arrayFieldElementType(m.ReturnType)
					}
					var arrayLine string
					if strings.HasPrefix(m.MethodName, "set") {
						arrayLine = fmt.Sprintf("MIQT_EXPORT void %s_%s(%s* self, const %s* %s);\n",
							className, m.SafeMethodName(), className, elemType, m.VariableFieldName)
					} else {
						arrayLine = fmt.Sprintf("MIQT_EXPORT %s* %s_%s(const %s* self);\n",
							elemType, className, m.SafeMethodName(), className)
					}
					cppheader.WriteString(arrayLine)
					cheader.WriteString(arrayLine)
				} else {
					methodLine := fmt.Sprintf("MIQT_EXPORT %s %s(%s);\n", m.ReturnType.RenderTypeCabi(), cabiMethodName(c, m), EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*"))
					cppheader.WriteString("// 来源: 类方法\n")
					cppheader.WriteString(methodLine)
					cheader.WriteString(methodLine)
				}
			}

			if m.IsSignal {
				signalLine := fmt.Sprintf("MIQT_EXPORT %s %s(%s* self, intptr_t slot);\n", m.ReturnType.RenderTypeCabi(), cabiConnectName(c, m), className)
				cppheader.WriteString("// 来源: 信号\n")
				cppheader.WriteString(signalLine)
				cheader.WriteString(signalLine)
			}
		}

		if len(c.Methods) > 0 {
			cppheader.WriteString("\n")
			cheader.WriteString("\n")
		}

		for _, m := range virtualMethods {
			overrideLine := fmt.Sprintf("MIQT_EXPORT bool %s(%s* self, intptr_t slot);\n", cabiOverrideVirtualName(c, m), "void")
			virtBaseLine := fmt.Sprintf("MIQT_EXPORT %s %s(%s);\n", m.ReturnType.RenderTypeCabi(), cabiVirtualBaseName(c, m), EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+"void"+"*"))
			cppheader.WriteString("// 来源: 虚函数\n")
			cppheader.WriteString(overrideLine)
			cheader.WriteString(overrideLine)
			cppheader.WriteString(virtBaseLine)
			cheader.WriteString(virtBaseLine)
		}
		if len(virtualMethods) > 0 {
			cppheader.WriteString("\n")
			cheader.WriteString("\n")

			for _, m := range protectedMethods {
				protectedLine := fmt.Sprintf("MIQT_EXPORT %s %s(bool* _dynamic_cast_ok, %s);\n", m.ReturnType.RenderTypeCabi(), cabiProtectedBaseName(c, m), EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+"void"+"*"))
				cppheader.WriteString("// 来源: 受保护方法\n")
				cppheader.WriteString(protectedLine)
				cheader.WriteString(protectedLine)
			}
			if len(protectedMethods) > 0 {
				cppheader.WriteString("\n")
				cheader.WriteString("\n")
			}
		}

		for _, m := range c.PrivateSignals {
			privSignalLine := fmt.Sprintf("MIQT_EXPORT %s %s(%s* self, intptr_t slot);\n", m.ReturnType.RenderTypeCabi(), cabiConnectName(c, m), className)
			cppheader.WriteString("// 来源: 私有信号\n")
			cppheader.WriteString(privSignalLine)
			cheader.WriteString(privSignalLine)
		}
		if len(c.PrivateSignals) > 0 {
			cppheader.WriteString("\n")
			cheader.WriteString("\n")
		}

		if c.CanDelete {
			deleteLine := fmt.Sprintf("MIQT_EXPORT void %s(%s* self);\n", cabiDeleteName(c), className)
			cppheader.WriteString("// 来源: 析构函数\n")
			cppheader.WriteString(deleteLine)
			cheader.WriteString(deleteLine)
		}

		cppheader.WriteString("\n")
		cheader.WriteString("\n")
	}

	freeFuncOverloadCount := make(map[string]int)
	freeFuncOverloadSeen := make(map[string]int)
	for _, fn := range src.FreeFunctions {
		isVariadic := false
		for _, p := range fn.Parameters {
			if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
				isVariadic = true
				break
			}
		}
		if isVariadic {
			continue
		}
		baseName := CabiClassName(fn.FunctionName)
		if cfg.FreeFuncCabiPrefix != "" {
			baseName = cfg.FreeFuncCabiPrefix + baseName
		}
		freeFuncOverloadCount[baseName]++
	}

	for _, fn := range src.FreeFunctions {
		hasVariadic := false
		for _, p := range fn.Parameters {
			if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
				hasVariadic = true
				break
			}
		}
		if hasVariadic {
			continue
		}
		funcName := CabiClassName(fn.FunctionName)
		if cfg.FreeFuncCabiPrefix != "" {
			funcName = cfg.FreeFuncCabiPrefix + funcName
		}
		if freeFuncOverloadCount[funcName] > 1 {
			freeFuncOverloadSeen[funcName]++
			funcName += fmt.Sprintf("_%d", freeFuncOverloadSeen[funcName])
		}
		params := make([]string, 0, len(fn.Parameters))
		for _, p := range fn.Parameters {
			params = append(params, p.RenderTypeCabi()+" "+p.cParameterName())
		}
		freeFuncLine := fmt.Sprintf("MIQT_EXPORT %s %s(%s);\n", fn.ReturnType.RenderTypeCabi(), funcName, strings.Join(params, ", "))
		cppheader.WriteString("// 来源: 自由函数\n")
		cppheader.WriteString(freeFuncLine)
		cheader.WriteString(freeFuncLine)
	}

	ending := `#endif // MIQT_TYPES_ONLY

#ifdef __cplusplus
} /* extern C */
#endif

#endif
`
	cppheader.WriteString(ending)

	cheaderEnding := `#endif // MIQT_TYPES_ONLY

#endif
`
	cheader.WriteString(cheaderEnding)
	return cppheader.String(), cheader.String(), nil
}

func emitParametersCabiConstructor(c *CppClass, ctor *CppMethod) string {

	slist := make([]string, 0, len(ctor.Parameters))
	for _, p := range ctor.Parameters {
		if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
			continue
		}
		slist = append(slist, p.RenderTypeCabi()+" "+p.cParameterName())
	}

	return strings.Join(slist, `, `)
}

func emitBindingCpp(src *CppParsedHeader, filename string, cfg EmitConfig) (string, error) {
	cppbody := strings.Builder{}

	cppbody.WriteString(`#include <` + filename + ">\n")
	cppbody.WriteString(`#include "gen_` + strings.TrimSuffix(filename, ".h") + `.h"` + "\n")
	cppbody.WriteString("#include <new>\n")

	platformRestriction := ActiveFilter.HeaderPlatformRestriction(filename)
	if platformRestriction != nil {
		cppbody.WriteString(`#if ` + platformRestriction.CxxIf() + "\n\n")
	}

	// Write prototypes for functions that the host language bindings should export
	// for virtual function overrides

	cppbody.WriteString(`
#ifdef __cplusplus
extern "C" {
#endif

`)

	for _, c := range src.Classes {
		writtenCallbacks := make(map[string]bool)

		methodsAndPrivateSignals := append(c.Methods, c.PrivateSignals...)

		for _, m := range methodsAndPrivateSignals {
			if m.IsSignal {
				callbackName := cabiCallbackName(c, m)
				if writtenCallbacks[callbackName] {
					continue
				}
				writtenCallbacks[callbackName] = true
				callback := "void " + callbackName + "(intptr_t"

				for _, p := range m.Parameters {
					callback += ", " + p.RenderTypeCabi()
				}
				callback += ");\n"
				cppbody.WriteString(callback)
			}
		}

		for _, m := range c.VirtualMethods() {
			callbackName := cabiCallbackName(c, m)
			if writtenCallbacks[callbackName] {
				continue
			}
			writtenCallbacks[callbackName] = true
			callback := m.ReturnType.RenderTypeCabi() + " " + callbackName + "(" + ifv(m.IsConst, "const ", "") + CabiClassName(c.ClassName) + "*, intptr_t"

			for _, p := range m.Parameters {
				callback += ", " + p.RenderTypeCabi()
			}
			callback += ");\n"
			cppbody.WriteString(callback)
		}
	}

	cppbody.WriteString(
		`#ifdef __cplusplus
} /* extern C */
#endif

`)

	for _, c := range src.Classes {
		if EmittedClasses.Has(c.ClassName) {
			continue
		}
		EmittedClasses.Set(c.ClassName, true)
		className := CabiClassName(c.ClassName)
		virtualMethods := c.VirtualMethods()
		protectedMethods := c.ProtectedMethods()

		if len(virtualMethods) > 0 {
			subclassName := CppSubclassName(c)

			cppbody.WriteString("class " + subclassName + " final : public " + c.ClassName + " {\n" +
				"public:\n" +
				"\n",
			)

			for _, ctor := range c.Ctors {
				cppbody.WriteString("\t" + subclassName + "(" + emitParametersCpp(ctor) + "): " + c.ClassName + "(" + emitParameterNames(ctor) + ") {}\n")
			}
			cppbody.WriteString("\n")

			if !c.CanDelete {
				cppbody.WriteString(
					"private:\n" +
						"\tvirtual ~" + subclassName + "();\n" + //  = delete;\n" +
						"\n" +
						"public:\n" +
						"\n",
				)
			} else {
				cppbody.WriteString(
					"\tvirtual ~" + subclassName + "() override = default;\n" +
						"\n",
				)
			}

			// Add default implementations for pure virtual functions that were skipped
			pureVirtualDefaults := ActiveFilter.GetPureVirtualDefaults(c.ClassName)
			for _, pv := range pureVirtualDefaults {
				cppbody.WriteString("\tvirtual " + pv.ReturnType + " " + pv.MethodSig + " override { " + pv.Body + " }\n\n")
			}

			for _, m := range virtualMethods {

				{
					var maybeReturn, maybeReturn2 string
					var returnTransformP, returnTransformF string
					if !m.ReturnType.Void() {
						maybeReturn = "return "

						maybeReturn2 = m.ReturnType.RenderTypeCabi() + " callback_return_value = "
						returnParam := m.ReturnType // copy
						returnParam.ParameterName = "callback_return_value"
						returnTransformP, returnTransformF = EmitCABI2CppForwarding(returnParam, "\t\t", c.ClassName)
					}

					handleVarname := "handle__" + m.SafeMethodName()

					cppbody.WriteString(
						"\t// cgo.Handle value for overwritten implementation\n" +
							"\tintptr_t " + handleVarname + " = 0;\n" +
							"\n",
					)

					// In the case of method overloads, we always need to use the
					// original method name (CppCallTarget), not the MethodName

					cppbody.WriteString(
						"\t// Subclass to allow providing a Go implementation\n" +
							"\tvirtual " + m.ReturnType.RenderTypeQtCpp() + " " + m.CppCallTarget() + "(" + emitParametersCpp(m) + ") " + ifv(m.IsConst, "const ", "") + ifv(len(m.Noexcept) > 0, m.Noexcept+" ", "") + "override {\n",
					)

					cppbody.WriteString("\t\tif (" + handleVarname + " == 0) {\n")
					if m.IsPureVirtual {
						if m.ReturnType.Void() {
							cppbody.WriteString("\t\t\treturn; // Pure virtual, there is no base we can call\n")
						} else {
							cppbody.WriteString("\t\t\treturn " + GetCppZeroValue(m.ReturnType) + "; // Pure virtual, there is no base we can call\n")
						}
					} else {
						cppbody.WriteString("\t\t\t" + maybeReturn + c.ClassName + "::" + m.CppCallTarget() + "(" + emitParameterNames(m) + ");\n")

						if m.ReturnType.Void() {
							cppbody.WriteString("\t\t\treturn;\n")
						}
					}
					cppbody.WriteString("\t\t}\n")

					paramArgs := []string{}
					paramArgs = append(paramArgs, "this")
					paramArgs = append(paramArgs, handleVarname)

					var signalCode string

					for i, p := range m.Parameters {
						signalCode += EmitAssignCppToCabi(fmt.Sprintf("\t\t%s sigval%d = ", p.RenderTypeCabi(), i+1), p, p.cParameterName())
						paramArgs = append(paramArgs, fmt.Sprintf("sigval%d", i+1))
					}

					cppbody.WriteString(
						"\n" +
							signalCode +
							"\t\t" + maybeReturn2 + cabiCallbackName(c, m) + "(" + strings.Join(paramArgs, `, `) + ");\n" +
							returnTransformP +
							ifv(maybeReturn == "", "", "\t\treturn "+returnTransformF+";") + "\n" +
							"\t}\n" +
							"\n",
					)
				}

				// If there is a base version of this method, add a helper to
				// allow calling it

				if !m.IsPureVirtual {

					// The virtualbase wrapper needs to take CABI parameters, not
					// real Qt parameters, in case there are protected enum types
					// (e.g. QAbstractItemView::CursorAction)

					// Because (in the Go projection) this is only exposed as a
					// super() argument to a real virtual override, we know that
					// the pointer type correctly points to our subclass and
					// therefore no static_cast<> validation is required

					cppbody.WriteString(
						"\tfriend " + m.ReturnType.RenderTypeCabi() + " " + cabiVirtualBaseName(c, m) + "(" + EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+"void*") + ");\n\n",
					)

				}
			}

			if len(protectedMethods) > 0 {
				cppbody.WriteString("\t// Wrappers to allow calling protected methods:\n")
			}

			for _, m := range protectedMethods {

				// The protectedbase wrapper needs to take CABI parameters, not
				// real Qt parameters, in case there are protected enum types
				// (e.g. QAbstractItemView::CursorAction)

				cppbody.WriteString(
					"\tfriend " + m.ReturnType.RenderTypeCabi() + " " + cabiProtectedBaseName(c, m) + "(bool* _dynamic_cast_ok, " + EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+"void*") + ");\n",
				)
			}

			cppbody.WriteString(
				"};\n" +
					"\n")
		}

		for i, ctor := range c.Ctors {
			hasVariadic := false
			for _, p := range ctor.Parameters {
				if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
					hasVariadic = true
					break
				}
			}
			if hasVariadic {
				continue
			}

			preamble, forwarding := EmitParametersCABI2CppForwarding(ctor.Parameters, "\t", c.ClassName)

			cppbody.WriteString(
				className + "* " + cabiNewName(c, i) + "(" + EmitParametersCabi(ctor, "") + ") {\n",
			)

			if ctor.RequireCpp != nil {
				cppbody.WriteString(
					"#if " + *ctor.RequireCpp + "\n",
				)
			}

			cppbody.WriteString(
				preamble +
					"\treturn new (std::nothrow) " +
					ifv(len(virtualMethods) > 0, CppSubclassName(c), c.ClassName) +
					"(" + forwarding + ");\n",
			)

			if ctor.RequireCpp != nil {
				cppbody.WriteString(
					"#else\n" +
						"\treturn nullptr;\n" +
						"#endif\n",
				)
			}

			cppbody.WriteString(
				"}\n" +
					"\n",
			)

		}

		// Add a helper method to retrieve base class pointers
		// That's because C++ virtual inheritance shifts the pointer; we
		// need the base pointers to call base methods from CGO
		if len(c.DirectInheritClassInfo()) > 0 {
			cppbody.WriteString(
				"void " + cabiVirtBaseName(c) + "(" + className + "* src",
			)
			for _, baseClass := range c.DirectInheritClassInfo() {
				cppbody.WriteString(", " + baseClass.Class.ClassName + "** outptr_" + CabiClassName(baseClass.Class.ClassName))
			}
			cppbody.WriteString(") {\n")
			for _, baseClass := range c.DirectInheritClassInfo() {
				cppbody.WriteString("\t*outptr_" + CabiClassName(baseClass.Class.ClassName) + " = static_cast<" + baseClass.Class.ClassName + "*>(src);\n")
			}
			cppbody.WriteString(
				"}\n" +
					"\n",
			)

		}

		for _, m := range c.Methods {

			// Skip methods with variadic ... parameters (cannot bind via C ABI)
			hasVariadic := false
			for _, p := range m.Parameters {
				if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
					hasVariadic = true
					break
				}
			}
			if strings.HasSuffix(m.SafeMethodName(), "With...") || strings.HasSuffix(m.SafeMethodName(), "WithDotDotDot") {
				hasVariadic = true
			}
			if hasVariadic {
				continue
			}

			// Protected virtual methods will be bound separately (the only
			// useful thing is to expose calling the virtual base)
			// Protected non-virtual methods should always be hidden
			if m.IsProtected {
				continue
			}

			// Need to take an extra 'self' parameter

			preamble, forwarding := EmitParametersCABI2CppForwarding(m.Parameters, "\t", c.ClassName)

			// callTarget is an rvalue representing the full C++ function call.
			callTarget := "self->"
			if m.IsStatic {
				callTarget = c.ClassName + "::"
			}

			callTarget += m.CppCallTarget() + "(" + forwarding + ")"

			// Qt 6.8 moved many operator== implementations from class methods
			// into global operators.
			// By using infix syntax, either can be called
			if m.IsReadonlyOperator() && len(m.Parameters) == 1 {
				operator := m.CppCallTarget()[8:]
				callTarget = "(*self " + operator + " " + forwarding + ")"
			}

			if m.RequireCpp != nil {
				var unavailableRetn string
				if retnCabi := m.ReturnType.RenderTypeCabi(); retnCabi == "void" {
					unavailableRetn = "\treturn;\n"
				} else {
					unavailableRetn = "\t" + retnCabi + " _ret_unavailable;\n" +
						"\treturn _ret_unavailable;\n"
				}

				cppbody.WriteString(fmt.Sprintf(
					"%s %s_%s(%s) {\n"+
						"#if "+*m.RequireCpp+"\n"+
						"%s"+
						"%s"+
						"#else\n"+
						"%s"+
						"#endif\n"+
						"}\n"+
						"\n",
					m.ReturnType.RenderTypeCabi(), className, m.SafeMethodName(), EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*"),
					preamble,
					EmitAssignCppToCabi("\treturn ", m.ReturnType, callTarget, className),
					unavailableRetn,
				))

			} else if m.BecomesNonConstInVersion != nil {

				nonConstCallTarget := "const_cast<" + c.ClassName + "*>(self)->" + m.CppCallTarget() + "(" + forwarding + ")"

				cppbody.WriteString("" +
					m.ReturnType.RenderTypeCabi() + " " + className + "_" + m.SafeMethodName() + "(" + EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*") + ") {\n" +
					preamble + "\n" +
					"// This method was changed from const to non-const in Qt " + *m.BecomesNonConstInVersion + "\n" +
					"#if QT_VERSION < QT_VERSION_CHECK(" + strings.Replace(*m.BecomesNonConstInVersion, `.`, `,`, -1) + ",0)\n" +
					EmitAssignCppToCabi("\treturn ", m.ReturnType, callTarget, className) +
					"#else\n" +
					EmitAssignCppToCabi("\treturn ", m.ReturnType, nonConstCallTarget, className) +
					"#endif\n" +
					"}\n" +
					"\n",
				)

			} else if m.ReturnType.BecomesConstInVersion != nil {

				cppbody.WriteString("" +
					"// This method's return type was changed from non-const to const in Qt " + *m.ReturnType.BecomesConstInVersion + "\n" +
					"#if QT_VERSION >= QT_VERSION_CHECK(" + strings.Replace(*m.ReturnType.BecomesConstInVersion, `.`, `,`, -1) + ",0)\n" +
					"const " + m.ReturnType.RenderTypeCabi() + " " + className + "_" + m.SafeMethodName() + "(" + EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*") + ") {\n" +
					"#else\n" +
					m.ReturnType.RenderTypeCabi() + " " + className + "_" + m.SafeMethodName() + "(" + EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*") + ") {\n" +
					"#endif\n" +
					preamble + "\n" +
					EmitAssignCppToCabi("\treturn ", m.ReturnType, callTarget, className) +
					"}\n\n",
				)

			} else {

				if m.IsVariable {
					isArray := false
					var elemType string
					if strings.HasPrefix(m.MethodName, "set") && len(m.Parameters) > 0 && isArrayFieldType(m.Parameters[0]) {
						isArray = true
						elemType = arrayFieldElementType(m.Parameters[0])
					} else if !strings.HasPrefix(m.MethodName, "set") && isArrayFieldType(m.ReturnType) {
						isArray = true
						elemType = arrayFieldElementType(m.ReturnType)
					}
					if isArray {
						if strings.HasPrefix(m.MethodName, "set") {
							cppbody.WriteString(fmt.Sprintf("void %s_%s(%s* self, const %s* %s) {\n",
								className, m.SafeMethodName(), className, elemType, m.VariableFieldName))
							cppbody.WriteString(fmt.Sprintf("\tmemcpy(self->%s, %s, sizeof(self->%s));\n}\n\n",
								m.VariableFieldName, m.VariableFieldName, m.VariableFieldName))
						} else {
							cppbody.WriteString(fmt.Sprintf("%s* %s_%s(const %s* self) {\n",
								elemType, className, m.SafeMethodName(), className))
							cppbody.WriteString(fmt.Sprintf("\treturn (%s*)self->%s;\n}\n\n", elemType, m.VariableFieldName))
						}
					} else {
						cppbody.WriteString(m.ReturnType.RenderTypeCabi() + " " + className + "_" + m.SafeMethodName() + "(" + EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*") + ") {\n")
						if strings.HasPrefix(m.MethodName, "set") {
							cppbody.WriteString(preamble + "\tself->" + m.VariableFieldName + " = " + forwarding + ";\n}\n\n")
						} else {
							cppbody.WriteString(EmitAssignCppToCabi("\treturn ", m.ReturnType, "self->"+m.VariableFieldName, className) + "}\n\n")
						}
					}
					continue
				}

				cppbody.WriteString(fmt.Sprintf(
					"%s %s_%s(%s) {\n"+
						"%s"+
						"%s"+
						"}\n"+
						"\n",
					m.ReturnType.RenderTypeCabi(), className, m.SafeMethodName(), EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+className+"*"),
					preamble,
					EmitAssignCppToCabi("\treturn ", m.ReturnType, callTarget, className),
				))

			}

			if m.IsSignal {
				// If there are hidden parameters, the type of the signal itself
				// needs to include them
				exactSignal := `static_cast<void (` + c.ClassName + `::*)(` + emitParameterTypesCppQualified(m, c.ClassName, true) + `)` + ifv(m.IsConst, ` const`, ``) + `>(&` + c.ClassName + `::` + m.CppCallTarget() + `)`

				paramArgs := []string{"slot"}
				paramArgDefs := []string{"intptr_t cb"}

				var signalCode string

				for i, p := range m.Parameters {
					signalCode += EmitAssignCppToCabi(fmt.Sprintf("\t\t%s sigval%d = ", p.RenderTypeCabi(), i+1), p, p.cParameterName())
					paramArgs = append(paramArgs, fmt.Sprintf("sigval%d", i+1))
					paramArgDefs = append(paramArgDefs, p.RenderTypeCabi()+" "+p.cParameterName())
				}

				signalCode += "\t\t" + cabiCallbackName(c, m) + "(" + strings.Join(paramArgs, `, `) + ");\n"

				cppbody.WriteString(
					`void ` + cabiConnectName(c, m) + `(` + className + `* self, intptr_t slot) {` + "\n" +
						"\t" + className + `::connect(self, ` + exactSignal + `, self, [=](` + emitParametersCppQualified(m, c.ClassName) + `) {` + "\n" +
						signalCode +
						"\t});\n" +
						"}\n" +
						"\n",
				)
			}

		}

		// FIXME(hack): In some platforms (Android Qt 5), instantiating a
		// protected enum fails in friend context:
		//
		//     QAbstractItemView::State _ret = self_cast->state();
		//     error: 'State' is a protected member of 'QAbstractItemView'
		//
		// @ref https://stackoverflow.com/q/52191903
		// However, it works fine on most other platforms. Probably this
		// is a GCC vs Clang difference.
		fixupProtectedReferences := func(assignStmts string) string {

			// Work around it for this specific class (fingers-crossed) by
			// referencing the protected enum via its subclass name
			ret := strings.Replace(assignStmts, c.ClassName+`::`, CppSubclassName(c)+`::`, -1)

			// Also need to scan parent classes (e.g. QColumnView friend
			// functions refer to its parent QAbstractItemView::State)
			for _, classInherit := range c.AllInheritsClassInfo() {
				ret = strings.Replace(ret, classInherit.Class.ClassName+`::`, CppSubclassName(c)+`::`, -1)
			}

			// The first instance of this class name change affected the very
			// method we're going to call
			// Undo it, but only once
			ret = strings.Replace(ret, "->"+CppSubclassName(c), "->"+c.ClassName, 1)

			return ret
		}

		if len(virtualMethods) > 0 {
			// This is a subclassed class. In that case, we allow calling
			// protected methods
			// This is a standalone function, but it can access the protected
			// method via a friend declaration

			subclassName := CppSubclassName(c)

			// Virtual override helpers
			for _, m := range virtualMethods {

				// Virtual methods: Allow overriding
				// (Never use a const self*)
				// The pointer that we are passed is the base type, not the subclassed
				// type. First cast the void* to the base type, and only then,
				// upclass it

				cppbody.WriteString(
					`bool ` + cabiOverrideVirtualName(c, m) + `(void* self, intptr_t slot) {` + "\n" +
						"\t" + subclassName + "* self_cast = static_cast<" + subclassName + "*>( (" + c.ClassName + "*)(self) );\n" +
						"\tif (self_cast == nullptr) {\n" +
						"\t\treturn false;\n" +
						"\t}\n" +
						"\n" +
						"\tself_cast->handle__" + m.SafeMethodName() + " = slot;\n" +
						"\treturn true;\n" +
						"}\n" +
						"\n",
				)

				// 2. Add CABI function to call the base method

				if !m.IsPureVirtual {
					// This is not generally exposed in the Go binding, but when overriding
					// the method, allows Go code to call super().

					// This calls the target Qt C++ method directly using fully
					// qualified syntax (`MiqtSubclass->QFoo::Bar()`). This method
					// takes and returns CABI types.

					vbpreamble, vbforwarding := EmitParametersCABI2CppForwarding(m.Parameters, "\t", c.ClassName)

					callTarget := "static_cast<" + ifv(m.IsConst, "const ", "") + subclassName + "*>(self)->" + c.ClassName + "::" + m.CppCallTarget() + "(" + vbforwarding + ")"

					cppbody.WriteString(
						m.ReturnType.RenderTypeCabi() + " " + cabiVirtualBaseName(c, m) + "(" + EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+"void*") + ") {\n" +
							vbpreamble +
							fixupProtectedReferences(EmitAssignCppToCabi("\treturn ", m.ReturnType, callTarget, c.ClassName)) +
							"}\n" +
							"\n",
					)

				}
			}

			for _, m := range protectedMethods {

				vbpreamble, vbforwarding := EmitParametersCABI2CppForwarding(m.Parameters, "\t\t", c.ClassName)
				vbCallTarget := "self_cast->" + m.CppCallTarget() + "(" + vbforwarding + ")"

				cppbody.WriteString(
					m.ReturnType.RenderTypeCabi() + " " + cabiProtectedBaseName(c, m) + "(bool* _dynamic_cast_ok, " + EmitParametersCabi(m, ifv(m.IsConst, "const ", "")+"void*") + ") {\n" +
						"\t" + subclassName + "* self_cast = static_cast<" + subclassName + "*>( (" + c.ClassName + "*)(self) );\n" +
						"\tif (self_cast == nullptr) {\n" +
						"\t\t*_dynamic_cast_ok = false;\n" +
						"\t\treturn " + GetCabiZeroValue(m.ReturnType) + ";\n" +
						"\t}\n" +
						"\n" +
						"\t*_dynamic_cast_ok = true;\n" +
						vbpreamble +
						fixupProtectedReferences(EmitAssignCppToCabi("\treturn ", m.ReturnType, vbCallTarget, c.ClassName)) +
						"}\n" +
						"\n",
				)

			}
		}

		for _, m := range c.PrivateSignals {
			exactSignal := `&` + c.ClassName + `::` + m.CppCallTarget()

			paramArgs := []string{"slot"}
			paramArgDefs := []string{"intptr_t cb"}

			var signalCode string

			for i, p := range m.Parameters {
				signalCode += EmitAssignCppToCabi(fmt.Sprintf("\t\t%s sigval%d = ", p.RenderTypeCabi(), i+1), p, p.cParameterName())
				paramArgs = append(paramArgs, fmt.Sprintf("sigval%d", i+1))
				paramArgDefs = append(paramArgDefs, p.RenderTypeCabi()+" "+p.cParameterName())
			}

			signalCode += "\t\t" + cabiCallbackName(c, m) + "(" + strings.Join(paramArgs, `, `) + ");\n"

			cppbody.WriteString(
				`void ` + cabiConnectName(c, m) + `(` + className + `* self, intptr_t slot) {` + "\n" +
					"\t" + className + `::connect(self, ` + exactSignal + `, self, [=](` + emitParametersCppQualified(m, c.ClassName) + `) {` + "\n" +
					signalCode +
					"\t});\n" +
					"}\n" +
					"\n",
			)
		}

		// Delete
		// If we subclassed, our class destructor is always virtual. Therefore
		// we can delete from the self ptr without any static_cast<>
		if c.CanDelete {
			cppbody.WriteString(
				"void " + cabiDeleteName(c) + "(" + className + "* self) {\n" +
					"\tdelete self;\n" +
					"}\n" +
					"\n",
			)
		}
	}

	// Detect overloaded free functions and add suffixes
	freeFuncOverloadCountCpp := make(map[string]int)
	freeFuncOverloadSeenCpp := make(map[string]int)
	for _, fn := range src.FreeFunctions {
		isVariadic := false
		for _, p := range fn.Parameters {
			if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
				isVariadic = true
				break
			}
		}
		if isVariadic {
			continue
		}
		baseName := CabiClassName(fn.FunctionName)
		if cfg.FreeFuncCabiPrefix != "" {
			baseName = cfg.FreeFuncCabiPrefix + baseName
		}
		freeFuncOverloadCountCpp[baseName]++
	}

	for _, fn := range src.FreeFunctions {
		funcName := CabiClassName(fn.FunctionName)
		cabiFuncName := funcName
		if cfg.FreeFuncCabiPrefix != "" {
			cabiFuncName = cfg.FreeFuncCabiPrefix + funcName
		}
		hasVariadic := false
		for _, p := range fn.Parameters {
			if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
				hasVariadic = true
				break
			}
		}
		if hasVariadic {
			continue
		}
		// Add suffix for overloaded free functions
		if freeFuncOverloadCountCpp[cabiFuncName] > 1 {
			freeFuncOverloadSeenCpp[cabiFuncName]++
			cabiFuncName += fmt.Sprintf("_%d", freeFuncOverloadSeenCpp[cabiFuncName])
		}
		preamble, forwarding := EmitParametersCABI2CppForwarding(fn.Parameters, "\t", "")

		params := make([]string, 0, len(fn.Parameters))
		for _, p := range fn.Parameters {
			params = append(params, p.RenderTypeCabi()+" "+p.cParameterName())
		}

		callTarget := fn.FunctionName + "(" + forwarding + ")"

		cppbody.WriteString(fn.ReturnType.RenderTypeCabi() + " " + cabiFuncName + "(" + strings.Join(params, ", ") + ") {\n")
		cppbody.WriteString(preamble)
		if fn.ReturnType.Void() {
			cppbody.WriteString("\t" + callTarget + ";\n")
		} else {
			cppbody.WriteString(EmitAssignCppToCabi("\treturn ", fn.ReturnType, callTarget))
		}
		cppbody.WriteString("}\n\n")
	}

	//

	if platformRestriction != nil {
		cppbody.WriteString(`#endif //` + platformRestriction.CxxIf() + "\n\n")
	}

	return cppbody.String(), nil
}

func GetReferencedTypesForTest(src *CppParsedHeader) []string {
	return getReferencedTypes(src)
}

func EmitBindingHeaderForTest(src *CppParsedHeader, filename string, packageName string, cfg EmitConfig) (string, string, error) {
	return emitBindingHeader(src, filename, packageName, cfg)
}
