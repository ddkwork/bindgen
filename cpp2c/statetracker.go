package cpp2c

import "github.com/ddkwork/golibrary/std/safemap"

type LookupResultClass struct {
	PackageName string
	Class       CppClass
}

type LookupResultTypedef struct {
	PackageName string
	Typedef     CppTypedef
}

type LookupResultEnum struct {
	PackageName string
	Enum        CppEnum
}

var (
	KnownClassnames = safemap.New[string, LookupResultClass]()
	KnownTypedefs   = safemap.New[string, LookupResultTypedef]()
	KnownEnums      = safemap.New[string, LookupResultEnum]()
	EmittedClasses  = safemap.New[string, bool]()
)

func FlushKnownTypes() {
	KnownClassnames.Reset()
	KnownTypedefs.Reset()
	KnownEnums.Reset()
	EmittedClasses.Reset()
}

func AddKnownTypes(packageName string, parsed *CppParsedHeader) {
	for _, c := range parsed.Classes {
		KnownClassnames.Set(c.ClassName, LookupResultClass{packageName, c})
	}
	for _, td := range parsed.Typedefs {
		KnownTypedefs.Set(td.Alias, LookupResultTypedef{packageName, td})
	}
	for _, en := range parsed.Enums {
		KnownEnums.Set(en.EnumName, LookupResultEnum{packageName, en})
	}
}
