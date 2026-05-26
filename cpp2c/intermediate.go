package cpp2c

import (
	"regexp"
	"strings"
)

type CppParameter struct {
	ParameterName string
	ParameterType string
	Const         bool
	Pointer       bool
	PointerCount  int
	ByRef         bool
	Optional      bool
	NoExcept      bool

	QtCppOriginalType     *CppParameter // If we rewrote QStringList->QList<String>, this field contains the original QStringList. Otherwise, it's blank
	BecomesConstInVersion *string       // "6,9"
}

func (p CppParameter) String() string {
	return "Param(" + ifv(p.Const, "const ", "") + p.ParameterType +
		ifv(p.Pointer, strings.Repeat("*", p.PointerCount), "") + ifv(p.ByRef, "&", "") +
		ifv(p.Optional, "?", "") + " " + p.ParameterName + ")"
}

func (p *CppParameter) ApplyTypedef(matchedUnderlyingType CppParameter) {
	if p.QtCppOriginalType == nil {
		tmp := *p                  // Copy
		p.QtCppOriginalType = &tmp // Overwrite once only, at the earliest base type
	}
	p.ParameterType = matchedUnderlyingType.ParameterType

	// If this was a pointer to a typedef'd type, or a typedef of a pointer type, we need to preserve that
	p.Const = p.Const || matchedUnderlyingType.Const
	p.Pointer = p.Pointer || matchedUnderlyingType.Pointer
	p.PointerCount += matchedUnderlyingType.PointerCount
	p.ByRef = p.ByRef || matchedUnderlyingType.ByRef
	p.Optional = p.Optional || matchedUnderlyingType.Optional
}

func (p *CppParameter) PointerTo() CppParameter {
	ret := *p // Copy
	ret.Pointer = true
	ret.PointerCount++
	return ret
}

func (p *CppParameter) ConstCast(isConst bool) CppParameter {
	ret := *p // Copy
	ret.Const = isConst
	return ret
}

func (p *CppParameter) GetQtCppType() *CppParameter {
	if p.QtCppOriginalType != nil {
		return p.QtCppOriginalType
	}

	return p
}

func (p CppParameter) QFlagsOf() (CppParameter, bool) {

	if strings.HasPrefix(p.ParameterType, `QFlags<`) {
		ret := ParseSingleTypeString(p.ParameterType[7 : len(p.ParameterType)-1])
		ret.ParameterName = p.ParameterName + "_qf"
		return ret, true
	}

	if under := p.QtCppOriginalType; under != nil {
		if strings.HasPrefix(under.ParameterType, `QFlags<`) {
			ret := ParseSingleTypeString(under.ParameterType[7 : len(under.ParameterType)-1])
			ret.ParameterName = under.ParameterName + "_qf"
			return ret, true
		}
	}

	return CppParameter{}, false
}

func (p CppParameter) IsFlagType() bool {
	_, ok := p.QFlagsOf()
	return ok
}

func (p CppParameter) QtClassType() bool {

	if KnownClassnames.Has(p.ParameterType) {
		return true
	}

	if p.ParameterType == "Scintilla::Internal::Point" {
		return true
	}

	if p.ParameterType == "QString" || p.ParameterType == "QByteArray" {
		return true
	}

	if len(p.ParameterType) >= 2 && p.ParameterType[0] == 'Q' && p.ParameterType[1] >= 'A' && p.ParameterType[1] <= 'Z' {
		if !p.IsKnownEnum() {
			if !KnownTypedefs.Has(p.ParameterType) {
				return true
			}
		}
	}

	return false
}

func (p CppParameter) IsKnownEnum() bool {
	return KnownEnums.Has(p.ParameterType)
}

func (p CppParameter) ResolvedEnumName() string {
	if en, ok := KnownEnums.Get(p.ParameterType); ok {
		return en.Enum.EnumName
	}
	return p.ParameterType
}

func (p CppParameter) QListOf() (CppParameter, string, bool) {
	if strings.HasPrefix(p.ParameterType, "QList<") && strings.HasSuffix(p.ParameterType, `>`) {
		ret := ParseSingleTypeString(p.ParameterType[6 : len(p.ParameterType)-1])
		ret.ParameterName = p.ParameterName + "_lv"
		return ret, "QList", true
	}

	if strings.HasPrefix(p.ParameterType, "QVector<") && strings.HasSuffix(p.ParameterType, `>`) {
		ret := ParseSingleTypeString(p.ParameterType[8 : len(p.ParameterType)-1])
		ret.ParameterName = p.ParameterName + "_vv"
		return ret, "QVector", true
	}

	return CppParameter{}, "", false
}

func (p CppParameter) QMapOf() (CppParameter, CppParameter, string, bool) {
	// n.b. Need to block QMap<k,v>::const_terator

	if strings.HasPrefix(p.ParameterType, `QMap<`) && strings.HasSuffix(p.ParameterType, `>`) {
		interior := tokenizeMultipleParameters(p.ParameterType[5 : len(p.ParameterType)-1])
		if len(interior) != 2 {
			panic("QMap<> has unexpected number of template arguments")
		}

		first := ParseSingleTypeString(interior[0])
		first.ParameterName = p.ParameterName + "_mapkey"
		second := ParseSingleTypeString(interior[1])
		second.ParameterName = p.ParameterName + "_mapval"
		return first, second, "QMap", true
	}

	if strings.HasPrefix(p.ParameterType, `QHash<`) && strings.HasSuffix(p.ParameterType, `>`) {
		interior := tokenizeMultipleParameters(p.ParameterType[6 : len(p.ParameterType)-1])
		if len(interior) != 2 {
			panic("QHash<> has unexpected number of template arguments")
		}

		first := ParseSingleTypeString(interior[0])
		first.ParameterName = p.ParameterName + "_hashkey"
		second := ParseSingleTypeString(interior[1])
		second.ParameterName = p.ParameterName + "_hashval"
		return first, second, "QHash", true
	}

	return CppParameter{}, CppParameter{}, "", false
}

func (p CppParameter) QPairOf() (CppParameter, CppParameter, bool) {
	if strings.HasPrefix(p.ParameterType, `QPair<`) && strings.HasSuffix(p.ParameterType, `>`) {
		interior := tokenizeMultipleParameters(p.ParameterType[6 : len(p.ParameterType)-1])
		if len(interior) != 2 {
			panic("QPair<> has unexpected number of template arguments")
		}

		first := ParseSingleTypeString(interior[0])
		first.ParameterName = p.ParameterName + "_first"
		second := ParseSingleTypeString(interior[1])
		second.ParameterName = p.ParameterName + "_second"
		return first, second, true
	}

	return CppParameter{}, CppParameter{}, false
}

func (p CppParameter) QSetOf() (CppParameter, bool) {
	if strings.HasPrefix(p.ParameterType, `QSet<`) {
		ret := ParseSingleTypeString(p.ParameterType[5 : len(p.ParameterType)-1])
		ret.ParameterName = p.ParameterName + "_sv"
		return ret, true
	}

	return CppParameter{}, false
}

func (p CppParameter) QMultiMapOf() bool {
	if strings.HasPrefix(p.ParameterType, `QMultiMap<`) ||
		strings.HasPrefix(p.ParameterType, `QMultiHash<`) {
		return true
	}

	return false
}

func (p CppParameter) IntType() bool {

	if p.IsKnownEnum() {
		return true
	}

	switch p.ParameterType {
	case "int", "unsigned int", "uint",
		"short", "unsigned short", "ushort", "qint16", "quint16", "uint16_t", "int16_t",
		"qint8", "quint8",
		"unsigned char", "signed char", "uchar",
		"long", "unsigned long", "ulong", "qint32", "quint32", "int32_t", "uint32_t",
		"longlong", "ulonglong", "qlonglong", "qulonglong", "qint64", "quint64", "int64_t", "uint64_t", "long long", "unsigned long long",
		"qintptr", "quintptr", "uintptr_t", "intptr_t",
		"qsizetype", "size_t",
		"QIntegerForSizeof<void *>::Unsigned",
		"QIntegerForSizeof<void *>::Signed",
		"QIntegerForSizeof<std::size_t>::Signed",
		"qptrdiff", "ptrdiff_t",
		"double", "float", "qreal":
		return true

	case "char":
		// Only count char as an integer type with cast assertions if it's
		// not possibly a char* string in disguise
		// (However, unsigned chars are always like ints)
		return !p.Pointer

	default:
		return false
	}
}

func (p CppParameter) IsFunctionPointerTypedef() bool {
	if strings.Contains(p.ParameterType, "(*)") {
		return true
	}
	if td, ok := KnownTypedefs.Get(p.ParameterType); ok {
		ut := td.Typedef.UnderlyingType
		if strings.Contains(ut.ParameterType, "(*)") {
			return true
		}
		if ut.Pointer && !ut.QtClassType() && !ut.IsKnownEnum() && !IsDirectCType(ut.ParameterType) && ut.ParameterType != "void" {
			return true
		}
	}
	return false
}

func (p CppParameter) Void() bool {
	return p.ParameterType == "void" && !p.Pointer
}

type CppProperty struct {
	PropertyName string
	PropertyType string
	Visibility   string
}

type CppMethod struct {
	MethodName         string       // C++ method name, unless OverrideMethodName is set, in which case a nice alternative name
	OverrideMethodName string       // C++ method name, present only if we changed the target
	ReturnType         CppParameter // Name not used
	Parameters         []CppParameter
	IsStatic           bool
	IsSignal           bool
	IsConst            bool
	Noexcept           string
	IsVariable         bool
	IsVirtual          bool
	IsPureVirtual      bool           // Virtual method was declared with = 0 i.e. there is no base method here to call
	IsProtected        bool           // If true, we can't call this method but we may still be able to overload it
	HiddenParams       []CppParameter // Populated if there is an overload with more parameters
	VariableFieldName  string

	// Special quirks
	RequireGOOS              *string // constructs a `if runtime.GOOS = {foo}` block in the Go side, no effect on CABI / C++ sides
	RequireCpp               *string // constructs a `#if {foo}` preprocessor block in the C++ side, no effect on Go / CABI sides
	BecomesNonConstInVersion *string // "6,7"
}

func (m CppMethod) CppCallTarget() string {
	if m.OverrideMethodName != "" {
		return m.OverrideMethodName
	}
	return m.MethodName
}

func (m *CppMethod) Rename(newName string) {
	if m.OverrideMethodName == "" {
		m.OverrideMethodName = m.MethodName
	} else {
		// If it was already set, we're already a level of overload resolution deep - preserve it
	}

	m.MethodName = newName
}

func IsArgcArgv(params []CppParameter, pos int) bool {
	// Check if the arguments starting at position=pos are the argc/argv pattern.
	// QApplication/QGuiApplication constructors are the only expected example of this.
	return (len(params) > pos+1 &&
		params[pos].ParameterName == "argc" &&
		params[pos].ParameterType == "int" &&
		params[pos].ByRef &&
		params[pos+1].ParameterName == "argv" &&
		params[pos+1].ParameterType == "char") &&
		params[pos+1].Pointer &&
		params[pos+1].PointerCount == 2
}

func IsReceiverMethod(params []CppParameter, pos int) bool {
	// Check if the arguments starting at position=pos are the receiver/member pattern.
	// QMenu->addAction is the main example of this
	return (len(params) > pos+1 &&
		params[pos].ParameterName == "receiver" &&
		params[pos].ParameterType == "QObject" &&
		params[pos].Pointer &&
		params[pos+1].ParameterName == "member" &&
		params[pos+1].ParameterType == "char" &&
		params[pos+1].Pointer)
}

func (m CppMethod) IsReceiverMethod() bool {
	// Returns true if any of the parameters use the receiever-method pattern
	for i := 0; i < len(m.Parameters); i++ {
		if IsReceiverMethod(m.Parameters, i) {
			return true
		}
	}
	return false
}

func (m CppMethod) SafeMethodName() string {
	// Strip redundant Qt prefix, we know these are all Qt functions
	tmp := strings.TrimPrefix(m.MethodName, "qt_")

	// Operator-overload methods have names not representable in binding
	// languages. Replace more specific cases first
	replacer := strings.NewReplacer(

		// `operator ` with a trailing space only occurs in conversion operators
		// Add a fake _ here, but it will be replaced with camelcase in the regex below
		`operator `, `To `,
		`::`, `__`, // e.g. `operator QCborError::Code`

		`==`, `Equal`,
		`!=`, `NotEqual`,
		`>=`, `GreaterOrEqual`,
		`<=`, `LesserOrEqual`,
		`=`, `Assign`,

		`<<`, `ShiftLeft`, // Qt classes use it more for stream functions e.g. in QDataStream
		`>>`, `ShiftRight`,
		`>`, `Greater`,
		`<`, `Lesser`,

		`+`, `Plus`,
		`-`, `Minus`,
		`*`, `Multiply`,
		`/`, `Divide`,
		`%`, `Modulo`,

		`&&`, `LogicalAnd`,
		`||`, `LogicalOr`,
		`!`, `Not`,
		`&`, `BitwiseAnd`,
		`|`, `BitwiseOr`,
		`~`, `BitwiseXor`,
		`^`, `BitwiseNot`,

		`->`, `PointerDereference`,
		`[]`, `Subscript`,
		`()`, `Call`,
	)
	tmp = replacer.Replace(tmp)

	// Replace spaces (e.g. `operator long long` with CamelCase
	tmp = regexp.MustCompile(` ([a-zA-Z])`).ReplaceAllStringFunc(tmp, func(match string) string { return strings.ToUpper(match[1:]) })

	// Replace spaces before underscores => To_GUID)
	tmp = strings.ReplaceAll(tmp, ` _`, `_`)
	// Replace spaces before digits => Type2)
	tmp = regexp.MustCompile(` (\d)`).ReplaceAllString(tmp, `${1}`)

	// Also replace any underscore_case with CamelCase
	// Only catch lowercase letters in this one, not uppercase, as it causes a
	// lot of churn for Scintilla
	tmp = regexp.MustCompile(`_([a-z])`).ReplaceAllStringFunc(tmp, func(match string) string { return strings.ToUpper(match[1:]) })

	return tmp
}

func (m CppMethod) IsReadonlyOperator() bool {
	targ := m.CppCallTarget()
	switch targ {
	case "operator==", "operator!=",
		"operator<", "operator<=", "operator>", "operator>=":
		return true
	}
	return false
}

type CppEnumEntry struct {
	EntryName  string
	EntryValue string
}

type CppEnum struct {
	EnumName       string
	UnderlyingType CppParameter
	Entries        []CppEnumEntry
}

func (e CppEnum) ShortEnumName() string {

	// Strip back one single :: pair from the generated variable name
	if nameParts := strings.Split(e.EnumName, `::`); len(nameParts) > 1 {
		nameParts = nameParts[0 : len(nameParts)-1]
		return strings.Join(nameParts, `::`)
	}

	// No change
	return e.EnumName
}

type CppClass struct {
	ClassName      string
	Abstract       bool
	Ctors          []CppMethod // only use the parameters
	DirectInherits []string    // other class names. This only includes direct inheritance - use AllInherits() to find recursive inheritance
	Methods        []CppMethod
	Props          []CppProperty
	CanDelete      bool

	ChildTypedefs  []CppTypedef
	ChildClassdefs []CppClass
	ChildEnums     []CppEnum
	PrivateMethods []string
	PrivateSignals []CppMethod
}

// VirtualMethods checks if the class has any virtual methods. This requires global
// state knowledge as virtual methods might have been inherited.
// C++ constructors cannot be virtual.
func (c *CppClass) VirtualMethods() []CppMethod {
	var ret []CppMethod
	var retNames = make(map[string]struct{}, 0) // if name is present, a child class found it first
	var block = slice_to_set(c.PrivateMethods)

	if len(c.Ctors) == 0 {
		// This class can't be constructed
		// Therefore there's no way to get a derived version of it for subclassing
		// (Unless we add custom constructors, but that seems like there would be
		// no in-Qt use for such a thing)
		// Pretend that this class is non-virtual
		return nil
	}

	if !ActiveFilter.AllowVirtualForClass(c.ClassName) {
		return nil
	}

	for _, m := range c.Methods {
		if !m.IsVirtual {
			continue
		}
		if m.IsSignal {
			continue
		}
		if !ActiveFilter.AllowVirtual(m) {
			continue
		}

		ret = append(ret, m)
		retNames[m.CppCallTarget()] = struct{}{}
	}

	for _, privMethod := range c.PrivateMethods {
		block[privMethod] = struct{}{}
	}

	// Go will automatically allow virtual overrides for the base type because
	// the parent struct is nested, but the resulting functions will not work
	// because the C ABI static_cast<> will fail for the base type.
	// Scan all inherited classes
	for _, cinfo := range c.AllInheritsClassInfo() {

		// If a base class is permanently unprojectable, the child classes
		// should be too
		if !ActiveFilter.AllowVirtualForClass(cinfo.Class.ClassName) {
			return nil
		}

		for _, m := range cinfo.Class.Methods {
			if !m.IsVirtual {
				continue
			}
			if m.IsSignal {
				continue
			}
			if !ActiveFilter.AllowVirtual(m) {
				continue
			}
			if _, ok := retNames[m.CppCallTarget()]; ok {
				continue // Already found in a child class
			}

			// It's possible that a child class marked a parent method as private
			// (e.g. Qt 5 QAbstractTableModel marks parent() as private)
			// But then we find the protected version further down
			// Use a blocklist to prevent exposing any deeper methods in the call chain
			if _, ok := block[m.MethodName]; ok {
				continue // Marked as private in a child class
			}

			// The class info we loaded has not had all typedefs applied to it
			// m is copied by value. Mutate it
			applyTypedefs_Method(&m)
			// Same with astTransformBlocklist
			if err := blocklist_MethodAllowed(&m); err != nil {
				continue
			}

			ret = append(ret, m)
			retNames[m.CppCallTarget()] = struct{}{}
		}

		// Append this parent's private-virtuals to blocklist so that we
		// do not consider them for grandparent classes
		for _, privMethod := range cinfo.Class.PrivateMethods {
			block[privMethod] = struct{}{}
		}
	}

	return ret
}

// ProtectedMethods checks if the class has any virtual methods. This requires global
// state knowledge as protected methods might have been inherited.
// This only considers protected methods, not protected constructor overloads.
func (c *CppClass) ProtectedMethods() []CppMethod {
	var ret []CppMethod
	var retNames = make(map[string]struct{}, 0) // if name is present, a child class found it first
	var block = slice_to_set(c.PrivateMethods)

	for _, m := range c.Methods {

		// A public method with the same name blocks-out any protected methods
		// with the same name (even if the signature is different).
		// e.g.
		// QTextList::setFormat(const QTextListFormat&) [PUBLIC] inherits
		// QTextBlockGroup inherits
		// QTextObject::setFormat(const QTextFormat&) [PROTECTED]
		// FIXME support selecting the parent overload(?)

		if !m.IsProtected {
			block[m.MethodName] = struct{}{}
			continue
		}
		if m.IsVirtual {
			block[m.MethodName] = struct{}{}
			continue
		}
		if m.IsSignal {
			block[m.MethodName] = struct{}{}
			continue
		}

		ret = append(ret, m)
		retNames[m.CppCallTarget()] = struct{}{}
	}

	for _, privMethod := range c.PrivateMethods {
		block[privMethod] = struct{}{}
	}

	for _, cinfo := range c.AllInheritsClassInfo() {

		for _, m := range cinfo.Class.Methods {
			if !m.IsProtected {
				continue
			}
			if m.IsVirtual {
				continue
			}
			if m.IsSignal {
				continue
			}

			if _, ok := retNames[m.CppCallTarget()]; ok {
				continue // Already found in a child class
			}

			// It's possible that a child class marked a parent method as private
			// (e.g. Qt 5 QAbstractTableModel marks parent() as private)
			// But then we find the protected version further down
			// Use a blocklist to prevent exposing any deeper methods in the call chain
			if _, ok := block[m.MethodName]; ok {
				continue // Marked as private in a child class
			}

			// The class info we loaded has not had all typedefs applied to it
			// m is copied by value. Mutate it
			applyTypedefs_Method(&m)
			// Same with astTransformBlocklist
			if err := blocklist_MethodAllowed(&m); err != nil {
				continue
			}

			ret = append(ret, m)
			retNames[m.CppCallTarget()] = struct{}{}
		}

		// Append this parent's private-virtuals to blocklist so that we
		// do not consider them for grandparent classes
		for _, privMethod := range cinfo.Class.PrivateMethods {
			block[privMethod] = struct{}{}
		}
	}

	return ret
}

// AllInheritsClassInfo recursively finds and lists all the parent classes of this class.
// It returns closest-ancestor-first.
func (c *CppClass) AllInheritsClassInfo() []LookupResultClass {
	var ret []LookupResultClass

	// FIXME prevent duplicates arising from diamond inheritance

	for _, baseClassInfo := range c.DirectInheritClassInfo() {

		ret = append(ret, baseClassInfo)

		recurseInfo := baseClassInfo.Class.AllInheritsClassInfo()
		for _, childClass := range recurseInfo {
			ret = append(ret, childClass)
		}
	}

	return ret
}

// DirectInheritClassInfo looks up the CppClass for each entry in DirectInherits.
func (c *CppClass) DirectInheritClassInfo() []LookupResultClass {
	var ret []LookupResultClass

	for _, inh := range c.DirectInherits {
		cinfo, ok := KnownClassnames.Get(inh)
		if !ok {
			continue
		}

		ret = append(ret, cinfo)
	}

	return ret
}

type CppTypedef struct {
	Alias          string
	UnderlyingType CppParameter
}

type CppFreeFunction struct {
	FunctionName string
	ReturnType   CppParameter
	Parameters   []CppParameter
}

type CppParsedHeader struct {
	Filename      string
	Typedefs      []CppTypedef
	Enums         []CppEnum
	Classes       []CppClass
	FreeFunctions []CppFreeFunction
}

func (c CppParsedHeader) Empty() bool {
	for _, en := range c.Enums {
		if en.EnumName != "" {
			return false
		}
	}
	return len(c.Classes) == 0 && len(c.FreeFunctions) == 0
}

func (c *CppParsedHeader) AddContentFrom(other *CppParsedHeader) {
	c.Classes = append(c.Classes, other.Classes...)
	c.Enums = append(c.Enums, other.Enums...)
	c.Typedefs = append(c.Typedefs, other.Typedefs...)
	c.FreeFunctions = append(c.FreeFunctions, other.FreeFunctions...)
}
