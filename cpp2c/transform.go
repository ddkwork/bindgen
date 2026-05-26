package cpp2c

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

var (
	generatedClassesMutex sync.Mutex
	generatedClasses      = make(map[string]bool)
)

func ResetGeneratedClasses() {
	generatedClassesMutex.Lock()
	defer generatedClassesMutex.Unlock()
	generatedClasses = make(map[string]bool)
}

func markClassGenerated(className string) bool {
	generatedClassesMutex.Lock()
	defer generatedClassesMutex.Unlock()
	if generatedClasses[className] {
		return false
	}
	generatedClasses[className] = true
	return true
}

// AstTransformChildClasses expands all classes with child class definitions
// into new top-level entries within the header.
func AstTransformChildClasses(parsed *CppParsedHeader) {
	var taken []CppClass

	for i, c := range parsed.Classes {
		nestedClasses := takeChildren(&parsed.Classes[i])
		taken = append(taken, nestedClasses...)

		parsed.Typedefs = append(parsed.Typedefs, c.ChildTypedefs...)
		parsed.Enums = append(parsed.Enums, c.ChildEnums...)
		for _, nc := range nestedClasses {
			parsed.Typedefs = append(parsed.Typedefs, nc.ChildTypedefs...)
			parsed.Enums = append(parsed.Enums, nc.ChildEnums...)
		}
	}

	parsed.Classes = append(parsed.Classes, taken...)
}

// takeChildren recursively takes the children of the class.
func takeChildren(c *CppClass) []CppClass {
	if len(c.ChildClassdefs) == 0 {
		return []CppClass{}
	}

	var ret []CppClass
	for _, child := range c.ChildClassdefs {
		ret = append(ret, takeChildren(&child)...)
		child.ChildClassdefs = nil
		ret = append(ret, child)
	}
	c.ChildClassdefs = nil
	return ret
}

// astTransformOptional expands all methods with optional parameters into
// explicit additional versions.
func astTransformOptional(parsed *CppParsedHeader) {
	for i, c := range parsed.Classes {
		anyChange := false

		for j, m := range c.Methods {
			if m.IsVirtual {
				continue
			}

			optionalStart := -1
			for k, p := range m.Parameters {
				if p.Optional {
					optionalStart = k
					break
				}
			}
			if optionalStart == -1 {
				continue
			}

			anyChange = true

			for x := optionalStart; x < len(m.Parameters); x++ {
				dupMethod := m
				dupMethod.Parameters = m.Parameters[0 : x+1]
				dupMethod.HiddenParams = m.Parameters[x+1:]
				c.Methods = append(c.Methods, dupMethod)
			}

			m.HiddenParams = m.Parameters[optionalStart:]
			m.Parameters = m.Parameters[0:optionalStart]
			c.Methods[j] = m
		}

		for j, m := range c.Ctors {
			optionalStart := -1
			for k, p := range m.Parameters {
				if p.Optional {
					optionalStart = k
					break
				}
			}
			if optionalStart == -1 {
				continue
			}

			anyChange = true

			for x := optionalStart; x < len(m.Parameters); x++ {
				dupCtor := m
				dupCtor.Parameters = m.Parameters[0 : x+1]
				dupCtor.HiddenParams = m.Parameters[x+1:]
				c.Ctors = append(c.Ctors, dupCtor)
			}

			m.HiddenParams = m.Parameters[optionalStart:]
			m.Parameters = m.Parameters[0:optionalStart]
			c.Ctors[j] = m
		}

		if anyChange {
			parsed.Classes[i] = c
		}
	}
}

func (p CppParameter) renderTypeForMethod() string {
	return strings.NewReplacer(" ", "", "::", "").Replace(p.ParameterType)
}

// astTransformOverloads renames methods if another method exists with the same name.
func astTransformOverloads(parsed *CppParsedHeader) {
	for i, c := range parsed.Classes {
		anyChange := false
		existing := map[string]struct{}{}
		for j, m := range c.Methods {
			originalProposal := m.SafeMethodName()
			proposedName := originalProposal

			if _, ok := existing[proposedName]; !ok {
				existing[proposedName] = struct{}{}
				continue
			}

			anyChange = true

			proposedName = func() (proposedName string) {
				if len(m.Parameters) == 1 {
					if len(m.Parameters[0].ParameterName) > 1 && !strings.HasPrefix(m.Parameters[0].ParameterName, "param") {
						proposedName = originalProposal + "With" + titleCase(m.Parameters[0].ParameterName)
					} else {
						proposedName = originalProposal + "With" + titleCase(m.Parameters[0].renderTypeForMethod())
					}
					if _, ok := existing[proposedName]; !ok {
						return proposedName
					}
				}

				ctr := 2
				for {
					proposedName = fmt.Sprintf("%s%d", originalProposal, ctr)
					if _, ok := existing[proposedName]; !ok {
						return proposedName
					}
					ctr++
				}
			}()

			existing[proposedName] = struct{}{}
			m.Rename(proposedName)
			c.Methods[j] = m
		}

		if anyChange {
			parsed.Classes[i] = c
		}
	}
}

// astTransformConstructorOrder creates a canonical ordering for constructors.
func astTransformConstructorOrder(packageName string, parsed *CppParsedHeader) {
	checkIsDefaultCtor := func(candidate *CppMethod) bool {
		return len(candidate.Parameters) == 1 &&
			candidate.Parameters[0].ParameterType == "QWidget" &&
			candidate.Parameters[0].Pointer
	}

	checkIsQVariantListCtor := func(candidate *CppMethod) bool {
		if len(candidate.Parameters) != 1 {
			return false
		}
		if t, _, ok := candidate.Parameters[0].QListOf(); ok && t.ParameterType == "QVariant" {
			return true
		}
		return false
	}

	for i, c := range parsed.Classes {
		sort.SliceStable(c.Ctors, func(i, j int) bool {
			ic := checkIsDefaultCtor(&c.Ctors[i])
			jc := checkIsDefaultCtor(&c.Ctors[j])
			if ic && !jc {
				return true
			}

			ic = checkIsQVariantListCtor(&c.Ctors[i])
			jc = checkIsQVariantListCtor(&c.Ctors[j])
			if !ic && jc {
				return true
			}

			if packageName == "qt" && c.ClassName == "QDateTime" {
				ic = (len(c.Ctors[i].Parameters) == 0)
				jc = (len(c.Ctors[j].Parameters) == 0)
				if !ic && jc {
					return true
				}
			}

			return false
		})

		parsed.Classes[i] = c
	}
}

// astTransformApplyQuirks applies quirk flags to the class/method.
func astTransformApplyQuirks(packageName string, parsed *CppParsedHeader) {
	for i, c := range parsed.Classes {
		for j := range c.Ctors {
			ActiveFilter.ApplyQuirks(packageName, c.ClassName, &parsed.Classes[i].Ctors[j])
		}
		for j := range c.Methods {
			ActiveFilter.ApplyQuirks(packageName, c.ClassName, &parsed.Classes[i].Methods[j])
		}
	}
}

// astTransformRedundant merges duplicate enum definitions.
type astTransformRedundant struct {
	preserve map[string]*CppEnum
}

func (a *astTransformRedundant) Process(parsed *CppParsedHeader) {
	for i, e := range parsed.Enums {
		prev, ok := a.preserve[e.EnumName]
		if !ok {
			a.preserve[e.EnumName] = &parsed.Enums[i]
			continue
		}

		if prev.UnderlyingType != e.UnderlyingType {
			panic("Enum " + e.EnumName + " is defined multiple times with different underlying types")
		}

		prev.Entries = append(prev.Entries, e.Entries...)
		parsed.Enums[i].EnumName = ""
	}
}

// astTransformBlocklist filters out methods using too-complex parameter types,
// and entire classes that may be disallowed.
func astTransformBlocklist(parsed *CppParsedHeader) {
	j := 0
	for _, c := range parsed.Classes {
		if !ActiveFilter.AllowClass(c.ClassName) {
			continue
		}

		if !markClassGenerated(c.ClassName) {
			continue
		}

		if c.Abstract && !ActiveFilter.AllowVirtualForClass(c.ClassName) {
			c.Ctors = nil
		}

		parsed.Classes[j] = c
		j++
	}
	parsed.Classes = parsed.Classes[:j]

	for i, c := range parsed.Classes {
		j := 0
		for _, m := range c.Ctors {
			if err := blocklist_MethodAllowed(&m); err != nil {
				continue
			}
			c.Ctors[j] = m
			j++
		}
		c.Ctors = c.Ctors[:j]

		j = 0
		for _, m := range c.Methods {
			if err := blocklist_MethodAllowed(&m); err != nil {
				continue
			}
			c.Methods[j] = m
			j++
		}
		c.Methods = c.Methods[:j]

		parsed.Classes[i] = c
	}
}

func blocklist_MethodAllowed(m *CppMethod) error {
	if err := ActiveFilter.AllowType(m.ReturnType, true); err != nil {
		return err
	}
	// Block methods with variadic ... parameters - cannot be expressed in C ABI
	for _, p := range m.Parameters {
		if p.ParameterType == "..." || strings.HasPrefix(p.ParameterType, "...") {
			return fmt.Errorf("variadic parameter")
		}
	}
	// Block operator basic_string_view - not a valid C operator
	if strings.Contains(m.MethodName, "basic_string_view") {
		return fmt.Errorf("basic_string_view operator")
	}
	// Block methods using std::string parameters (e.g. fromStdString) - can't express in C ABI
	for _, p := range m.Parameters {
		if strings.HasPrefix(p.ParameterType, "std::string") {
			return fmt.Errorf("std::string parameter")
		}
	}
	// Block methods returning std::string
	if strings.HasPrefix(m.ReturnType.ParameterType, "std::string") {
		return fmt.Errorf("std::string return type")
	}
	// Block QByteArray::erase - uses iterator types that can't be expressed in C ABI
	if m.MethodName == "erase" && len(m.Parameters) > 0 {
		for _, p := range m.Parameters {
			if strings.Contains(p.ParameterType, "iterator") {
				return fmt.Errorf("iterator-based erase")
			}
		}
	}
	// Block methods with QByteArrayView/QStringView/QUtf8StringView/QAnyStringView/QLatin1StringView
	// parameters - these view types can't be constructed from C ABI types
	for _, p := range m.Parameters {
		if p.ParameterType == "QByteArrayView" || p.ParameterType == "QStringView" ||
			p.ParameterType == "QUtf8StringView" || p.ParameterType == "QAnyStringView" ||
			p.ParameterType == "QLatin1StringView" || p.ParameterType == "QLatin1String" ||
			p.ParameterType == "QStringRef" {
			return fmt.Errorf("string view parameter type: %s", p.ParameterType)
		}
	}
	// Block methods returning string view types by value (not pointer)
	if !m.ReturnType.Pointer && !m.ReturnType.ByRef {
		switch m.ReturnType.ParameterType {
		case "QByteArrayView", "QStringView", "QUtf8StringView", "QAnyStringView",
			"QLatin1StringView", "QLatin1String", "QStringRef":
			return fmt.Errorf("string view return type: %s", m.ReturnType.ParameterType)
		}
	}
	// Block methods taking QBasicAtomicInteger by value or reference
	for _, p := range m.Parameters {
		if strings.HasPrefix(p.ParameterType, "QBasicAtomicInteger") {
			return fmt.Errorf("QBasicAtomicInteger parameter")
		}
	}
	if strings.HasPrefix(m.ReturnType.ParameterType, "QBasicAtomicInteger") {
		return fmt.Errorf("QBasicAtomicInteger return type")
	}
	// Block methods returning QByteArray::DataPointer or QByteArray::FromBase64Result by value
	if !m.ReturnType.Pointer && !m.ReturnType.ByRef {
		if strings.HasPrefix(m.ReturnType.ParameterType, "QByteArray::DataPointer") ||
			strings.HasPrefix(m.ReturnType.ParameterType, "QByteArray::FromBase64Result") ||
			strings.HasPrefix(m.ReturnType.ParameterType, "QArrayDataPointer") {
			return fmt.Errorf("internal Qt data type return: %s", m.ReturnType.ParameterType)
		}
	}
	// Block methods with operator basic_string_view in return type
	if strings.Contains(m.ReturnType.ParameterType, "basic_string_view") {
		return fmt.Errorf("basic_string_view return type")
	}
	// Block methods returning QChar::fromUcs4::R or QChar::fromUcs2::R internal types
	if strings.HasPrefix(m.ReturnType.ParameterType, "QChar::fromUcs4::R") ||
		strings.HasPrefix(m.ReturnType.ParameterType, "QChar::fromUcs2::R") {
		return fmt.Errorf("QChar internal return type: %s", m.ReturnType.ParameterType)
	}
	for _, p := range m.Parameters {
		if err := ActiveFilter.AllowType(p, false); err != nil {
			return err
		}
	}
	return nil
}

// AstTransformTypedefs replaces the ParameterType with any known typedef value.
func AstTransformTypedefs(parsed *CppParsedHeader) {
	for i, c := range parsed.Classes {
		for j, m := range c.Methods {
			applyTypedefs_Method(&m)
			c.Methods[j] = m
		}
		for j, m := range c.Ctors {
			applyTypedefs_Method(&m)
			c.Ctors[j] = m
		}
		parsed.Classes[i] = c
	}

	for i, e := range parsed.Enums {
		e.UnderlyingType = ApplyTypedefs(e.UnderlyingType)
		parsed.Enums[i] = e
	}
}

func ApplyTypedefs(p CppParameter) CppParameter {
	for {
		td, ok := KnownTypedefs.Get(p.ParameterType)
		if !ok {
			break
		}
		p.ApplyTypedef(td.Typedef.UnderlyingType)
	}

	if t, containerType, ok := p.QListOf(); ok {
		t2 := ApplyTypedefs(t)
		t2.QtCppOriginalType = nil
		if p.QtCppOriginalType == nil {
			tmp := p
			p.QtCppOriginalType = &tmp
		}
		p.ParameterType = containerType + `<` + t2.RenderTypeQtCpp() + `>`
	} else if kType, vType, containerType, ok := p.QMapOf(); ok {
		kType2 := ApplyTypedefs(kType)
		kType2.QtCppOriginalType = nil
		vType2 := ApplyTypedefs(vType)
		vType2.QtCppOriginalType = nil
		if p.QtCppOriginalType == nil {
			tmp := p
			p.QtCppOriginalType = &tmp
		}
		p.ParameterType = containerType + `<` + kType2.RenderTypeQtCpp() + `, ` + vType2.RenderTypeQtCpp() + `>`
	} else if kType, vType, ok := p.QPairOf(); ok {
		kType2 := ApplyTypedefs(kType)
		kType2.QtCppOriginalType = nil
		vType2 := ApplyTypedefs(vType)
		vType2.QtCppOriginalType = nil
		if p.QtCppOriginalType == nil {
			tmp := p
			p.QtCppOriginalType = &tmp
		}
		p.ParameterType = `QPair<` + kType2.RenderTypeQtCpp() + `, ` + vType2.RenderTypeQtCpp() + `>`
	}

	return p
}

func applyTypedefs_Method(m *CppMethod) {
	for k, p := range m.Parameters {
		transformed := ApplyTypedefs(p)
		m.Parameters[k] = transformed
	}
	m.ReturnType = ApplyTypedefs(m.ReturnType)
}
