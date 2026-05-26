package c2go

import (
	"github.com/ddkwork/golibrary/std/safemap"
)

var GoReservedTypes = map[string]bool{
	"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
	"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
	"uintptr": true, "string": true, "bool": true, "byte": true, "rune": true,
	"float32": true, "float64": true, "complex64": true, "complex128": true,
	"error": true, "any": true, "comparable": true,
}

type EnumInfo struct {
	goName      string
	cName       string
	members     *safemap.M[string, EnumMember]
	memberOrder []string
	source      string
	lineNo      int
}

type EnumMember struct {
	goName string
	value  string
	intVal int64
	hasVal bool
}

type BitfieldInfo struct {
	goName    string
	goType    string
	bitOffset int
	bitWidth  int64
	backing   string
	backType  string
}

type StructInfo struct {
	goName         string
	cName          string
	fields         string
	methods        string
	source         string
	originalSource string
	lineNo         int
}

type TypedefInfo struct {
	goName string
	goType string
	source string
	isFunc bool
	lineNo int
}

type MacroConstInfo struct {
	goName   string
	goType   string
	value    string
	goValue  string
	source   string
	isVar    bool
	cName    string
	goParams string
	goBody   string
	lineNo   int
}

type Result struct {
	Structs  *safemap.M[string, StructInfo]
	Enums    *safemap.M[string, EnumInfo]
	Typedefs *safemap.M[string, TypedefInfo]
	Macros   *safemap.M[string, MacroConstInfo]
	FnMacros *safemap.M[string, MacroConstInfo]
	Imports  map[string]bool
}

type BindgenConfig struct {
	HeadersDir       string
	OutputDir        string
	PackageName      string
	ModuleName       string
	HeaderOrder      []string
	BindDll          bool
	DllName          string
	DllFuncFilter    func(name string) bool
	ExtraIncludeDirs []string
	RecurseHeaders   bool
	Predefined       string
	ExtraConstants   map[string]ExtraConst
}

type ExtraConst struct {
	GoType  string
	GoValue string
}

type DllFuncInfo struct {
	goName     string
	cName      string
	returnType string
	params     []DllFuncParam
}

type DllFuncParam struct {
	goName    string
	goType    string
	isPointer bool
}
