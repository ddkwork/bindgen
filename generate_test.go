package main

import (
	"fmt"
	"maps"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"slices"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
	"modernc.org/cc/v4"
)

func TestFixError(t *testing.T) {
	stream.Fmt(".")
	stream.Fix(".")
}

func newMSVCConfig(t testing.TB) *cc.Config {
	abi, err := cc.NewABI(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		t.Fatalf("cc.NewABI: %v", err)
	}
	ewdkPaths := findEWDKIncludePaths()
	cfg := &cc.Config{
		ABI:             abi,
		CC:              findClExe(),
		Predefined:      loadMSVCExtraTypes(),
		IncludePaths:    append([]string{""}, ewdkPaths...),
		SysIncludePaths: ewdkPaths,
	}
	return cfg
}

type enumInfo struct {
	goName      string
	cName       string
	members     *safemap.M[string, enumMember]
	memberOrder []string
	source      string
	lineNo      int
}

type enumMember struct {
	goName string
	value  string
	intVal int64
	hasVal bool
}

type bitfieldInfo struct {
	goName    string
	goType    string
	bitOffset int
	bitWidth  int64
	backing   string
	backType  string
}

type structInfo struct {
	goName         string
	cName          string
	fields         string
	methods        string
	source         string
	originalSource string
	lineNo         int
}

type typedefInfo struct {
	goName string
	goType string
	source string
	isFunc bool
	lineNo int
}

type macroConstInfo struct {
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
	Structs  *safemap.M[string, structInfo]
	Enums    *safemap.M[string, enumInfo]
	Typedefs *safemap.M[string, typedefInfo]
	Macros   *safemap.M[string, macroConstInfo]
	FnMacros *safemap.M[string, macroConstInfo]
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
	SingleFile       bool
	Predefined       string
	SkipMSVCTypes    bool
	ExtraConstants   map[string]extraConst
}

type extraConst struct {
	GoType  string
	GoValue string
}

func Generate(t *testing.T, configs []BindgenConfig) {
	fmt.Println("=== HyperDbg Bindgen ===")

	cfg := newMSVCConfig(t)
	fmt.Printf("CC: %s\n", cfg.CC)
	fmt.Printf("Predefined: %d bytes\n", len(cfg.Predefined))
	fmt.Printf("IncludePaths: %v\n", cfg.IncludePaths)

	for _, bc := range configs {
		processBindgenConfig(t, cfg, bc)
	}

	fmt.Println("\n✅ All binding generations completed!")
}

func TestGenerate(t *testing.T) {
	t.Run("hyperdbg", func(t *testing.T) {
		Generate(t, []BindgenConfig{{
			HeadersDir:       "project/hyperdbg/clone/SDK",
			OutputDir:        "project/hyperdbg/sdk",
			PackageName:      "sdk",
			HeaderOrder:      []string{"HyperDbgSdk.h"},
			ExtraIncludeDirs: []string{"project/hyperdbg/clone", "project/hyperdbg/clone/SDK/headers"},
			SingleFile:       true,
		}})
	})

	t.Run("zydis", func(t *testing.T) {
		Generate(t, []BindgenConfig{{
			HeadersDir:     "project/zydis/clone/zydis/include",
			OutputDir:      "project/zydis",
			PackageName:    "zydis",
			RecurseHeaders: true,
			SingleFile:     true,
			HeaderOrder:    []string{"Zydis/Zydis.h"},
			BindDll:        true,
			DllName:        "zydis.dll",
			DllFuncFilter: func(name string) bool {
				return strings.HasPrefix(name, "Zydis") || strings.HasPrefix(name, "Zyan")
			},
			ExtraIncludeDirs: []string{
				"project/zydis/clone/zydis/dependencies/zycore/include",
			},
		}})
	})
	t.Run("ipmrec", func(t *testing.T) {
		Generate(t, []BindgenConfig{{
			HeadersDir:  "project/ARImpRec/clone",
			OutputDir:   "project/ARImpRec",
			PackageName: "ipmrec",
			HeaderOrder: []string{
				"ARImpRec.h",
			},
			BindDll: true,
			DllName: "ARImpRec.dll",
			DllFuncFilter: func(name string) bool {
				return !strings.HasPrefix(name, "__builtin") && !strings.HasPrefix(name, "_")
			},
		}})
	})

	t.Run("keystone", func(t *testing.T) {
		Generate(t, []BindgenConfig{{
			HeadersDir:  "project/keystone/clone/keystone/include/keystone",
			OutputDir:   "project/keystone",
			PackageName: "keystone",
			HeaderOrder: []string{
				"keystone.h",
				"arm.h",
				"arm64.h",
				"evm.h",
				"hexagon.h",
				"mips.h",
				"ppc.h",
				"riscv.h",
				"sparc.h",
				"systemz.h",
				"x86.h",
			},
			BindDll: true,
			DllName: "keystone.dll",
			DllFuncFilter: func(name string) bool {
				return strings.HasPrefix(name, "ks_")
			},
		}})
	})

	t.Run("windivert", func(t *testing.T) {
		Generate(t, []BindgenConfig{{
			HeadersDir:  "project/WinDivert/clone/include",
			OutputDir:   "project/WinDivert",
			PackageName: "windivert",
			HeaderOrder: []string{
				"windivert_bindgen.h",
			},
			BindDll: true,
			DllName: "WinDivert64.dll",
			DllFuncFilter: func(name string) bool {
				return strings.HasPrefix(name, "WinDivert")
			},
		}})
	})

	t.Run("everything", func(t *testing.T) {
		Generate(t, []BindgenConfig{{
			HeadersDir:  "project/Everything-SDK/clone/include",
			OutputDir:   "project/Everything-SDK",
			PackageName: "everything",
			HeaderOrder: []string{
				"Everything_stub.h",
			},
			BindDll: true,
			DllName: "Everything64.dll",
			DllFuncFilter: func(name string) bool {
				return strings.HasPrefix(name, "Everything") || strings.HasPrefix(name, "EVERYTHING")
			},
		}})
	})

	t.Run("xed", func(t *testing.T) {
		Generate(t, []BindgenConfig{{
			HeadersDir:  "project/xed/clone/include",
			OutputDir:   "project/xed",
			PackageName: "xed",
			HeaderOrder: []string{
				"xed-interface.h",
			},
			SingleFile: true,
			BindDll:    true,
			DllName:    "xed.dll",
			Predefined: `
#define XED_ENCODER
#define XED_ENCODE_ORDER_MAX_OPERANDS 5
#define XED_ENCODE_ORDER_MAX_ENTRIES 35
#define stderr ((void*)0)
int fprintf(void*, const char*, ...);
int fflush(void*);
void abort(void);
`,
			ExtraIncludeDirs: []string{
				"project/xed/clone/include",
				"project/xed/clone/xed/include/public/xed",
			},
			DllFuncFilter: func(name string) bool {
				return strings.HasPrefix(name, "xed_")
			},
		}})
	})

	// {
	// 	Name:        "header",
	// 	HeadersDir:  filepath.Join("..", "..", "HyperDbgUnified", "HyperDbg", "hyperdbg", "libhyperdbg", "header"),
	// 	OutputDir:   filepath.Join("..", "header"),
	// 	PackageName: "header",
	// 	HeaderOrder: []string{
	// 		"assembler.h",
	// 		"commands.h",
	// 		"common.h",
	// 		"communication.h",
	// 		"debugger.h",
	// 		"export.h",
	// 		"forwarding.h",
	// 		"globals.h",
	// 		"help.h",
	// 		"hwdbg-interpreter.h",
	// 		"hwdbg-scripts.h",
	// 		"inipp.h",
	// 		"install.h",
	// 		"kd.h",
	// 		"libhyperdbg.h",
	// 		"list.h",
	// 		"namedpipe.h",
	// 		"objects.h",
	// 		"pci-id.h",
	// 		"pe-parser.h",
	// 		"rev-ctrl.h",
	// 		"script-engine.h",
	// 		"steppings.h",
	// 		"symbol.h",
	// 		"tests.h",
	// 		"transparency.h",
	// 		"ud.h",
	// 	},
	// },
}

func processBindgenConfig(t *testing.T, cfg *cc.Config, bc BindgenConfig) {
	fmt.Printf("\n=== Processing: %s ===\n", bc.PackageName)

	if bc.Predefined != "" {
		cfg.Predefined += bc.Predefined
	}

	cfg.IncludePaths = append(cfg.IncludePaths, bc.HeadersDir)
	cfg.SysIncludePaths = append(cfg.SysIncludePaths, bc.HeadersDir)
	for _, dir := range bc.ExtraIncludeDirs {
		cfg.IncludePaths = append(cfg.IncludePaths, dir)
		cfg.SysIncludePaths = append(cfg.SysIncludePaths, dir)
	}

	headerMap := make(map[string]string)
	var scanDir func(dir string, prefix string)
	scanDir = func(dir string, prefix string) {
		files, err := os.ReadDir(dir)
		if err != nil {
			return
		}
		for _, file := range files {
			path := filepath.Join(dir, file.Name())
			if file.IsDir() {
				if bc.RecurseHeaders {
					scanDir(path, prefix+file.Name()+"/")
				}
				continue
			}
			if !strings.HasSuffix(file.Name(), ".h") {
				continue
			}
			content, err := os.ReadFile(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", path, err)
				continue
			}
			key := prefix + file.Name()
			headerMap[key] = string(content)
		}
	}
	scanDir(bc.HeadersDir, "")
	for _, dir := range bc.ExtraIncludeDirs {
		scanDir(dir, "")
	}

	files, err := os.ReadDir(bc.HeadersDir)
	if err != nil {
		t.Logf("Skip %s: headers directory not found: %s", bc.PackageName, bc.HeadersDir)
		return
	}
	_ = files

	var sources []cc.Source
	sources = append(sources, cc.Source{Name: "<predefined>", Value: cfg.Predefined})
	sources = append(sources, cc.Source{Name: "<builtin>", Value: cc.Builtin})
	if !bc.SkipMSVCTypes {
		sources = append(sources, cc.Source{Name: "<msvc_types>", Value: loadMSVCTypes()})
	}

	for _, name := range bc.HeaderOrder {
		content, ok := headerMap[name]
		if !ok {
			continue
		}
		sources = append(sources, cc.Source{Name: name, Value: content})
		fmt.Printf("Processing: %s\n", name)
	}

	packedFiles := make(map[string]bool)
	for name, content := range headerMap {
		if strings.Contains(content, "#pragma pack(1)") || strings.Contains(content, "#pragma pack(push, 1)") {
			baseName := strings.TrimSuffix(name, ".h")
			baseName = filepath.Base(baseName)
			packedFiles[baseName] = true
		}
	}

	cfg.PragmaHandler = func(toks []cc.Token) error {
		return nil
	}

	ast, err := cc.Translate(cfg, sources)
	if err != nil {
		t.Logf("Translation error (partial parse): %v", err)
		t.Log("Continuing with partial AST...")
	}

	os.MkdirAll(bc.OutputDir, 0o755)

	modPath := bc.ModuleName
	if modPath == "" {
		modPath = "github.com/ddkwork/" + bc.PackageName
	}
	modContent := fmt.Sprintf("module %s\n\ngo 1.26.1\n", modPath)
	if err := os.WriteFile(filepath.Join(bc.OutputDir, "go.mod"), []byte(modContent), 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing go.mod: %v\n", err)
	}

	if ast == nil {
		t.Log("No AST generated, skipping output")
		return
	}

	result := Result{
		Structs:  safemap.New[string, structInfo](true),
		Enums:    safemap.New[string, enumInfo](true),
		Typedefs: safemap.New[string, typedefInfo](true),
		Macros:   safemap.New[string, macroConstInfo](true),
		FnMacros: safemap.New[string, macroConstInfo](true),
		Imports:  make(map[string]bool),
	}

	result.Macros.Set("PAGE_SIZE", macroConstInfo{
		goName:  "PageSize",
		goType:  "uint32",
		value:   "0x1000",
		goValue: "0x1000",
		source:  "Constants.h",
		isVar:   false,
		cName:   "PAGE_SIZE",
		lineNo:  0,
	})

	type prebuiltStruct struct {
		goName, cName, source, originalSource string
		fields                                string
		lineNo                                int
	}
	for _, ps := range []prebuiltStruct{
		{"BUFFER_HEADER", "_BUFFER_HEADER", "Constants.h", "HyperDbgUnified/HyperDbg/hyperdbg/hyperlog/header/Logging.h", "\tOperationNumber uint32\n\tBufferLength    uint32\n\tValid           bool\n", 57},
		{"DEBUGGER_EVENT_ACTION", "_DEBUGGER_EVENT_ACTION", "Constants.h", "HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/header/debugger/core/Debugger.h", "\tTag                       uint64\n\tActionOrderCode           uint32\n\tActionsList               LIST_ENTRY\n\tActionType                DEBUGGER_EVENT_ACTION_TYPE_ENUM\n\tImmediatelySendTheResults bool\n\tScriptConfiguration       DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION\n\tRequestedBuffer           DEBUGGER_EVENT_REQUEST_BUFFER\n\tCustomCodeBufferSize      uint32\n\tCustomCodeBufferAddress   uintptr\n", 79},
		{"DEBUGGER_EVENT", "_DEBUGGER_EVENT", "Constants.h", "HyperDbgUnified/HyperDbg/hyperdbg/hyperkd/header/debugger/core/Debugger.h", "\tTag                    uint64\n\tEventsOfSameTypeList    LIST_ENTRY\n\tEventType              VMM_EVENT_TYPE_ENUM\n\tEnabled                bool\n\tCoreId                 uint32\n\tProcessId              uint32\n\tActionsListHead        LIST_ENTRY\n\tCountOfActions         uint32\n\tEnableShortCircuiting  bool\n\tEventMode              VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE\n\tInitOptions            DEBUGGER_EVENT_OPTIONS\n\tOptions                DEBUGGER_EVENT_OPTIONS\n\tConditionsBufferSize   uint32\n\tConditionBufferAddress uintptr\n", 107},
	} {
		result.Structs.Set(ps.goName, structInfo{
			goName:         ps.goName,
			cName:          ps.cName,
			fields:         ps.fields,
			source:         ps.source,
			originalSource: ps.originalSource,
			lineNo:         ps.lineNo,
		})
	}

	for l := ast.TranslationUnit; l != nil; l = l.TranslationUnit {
		ed := l.ExternalDeclaration
		if ed == nil {
			continue
		}
		sourceFile := findSourceFileName(l, sources, headerMap)
		baseName := strings.TrimSuffix(filepath.Base(sourceFile), ".h")
		isMsvcTypes := sourceFile == "<msvc_types>"
		if sourceFile == "<builtin>" || sourceFile == "<predefined>" {
			continue
		}
		switch ed.Case {
		case cc.ExternalDeclarationDecl:
			decl := ed.Declaration
			if decl == nil || decl.DeclarationSpecifiers == nil {
				continue
			}
			ds := decl.DeclarationSpecifiers
			if decl.InitDeclaratorList != nil {
				for id := decl.InitDeclaratorList; id != nil; id = id.InitDeclaratorList {
					if id.InitDeclarator == nil || id.InitDeclarator.Declarator == nil {
						continue
					}
					tdName := id.InitDeclarator.Declarator.Name()
					if tdName == "" {
						continue
					}
				}
			}
			switch t := ds.Type().(type) {
			case *cc.StructType:
				tag := t.Tag()
				structName := string(tag.Src())
				if structName != "" {
					goStructName := cTagToGoName(structName)
					isPacked := packedFiles[baseName]
					fieldDefs, methodDefs, innerTypes := generateStructFields(t, goStructName, isPacked)

					result.Structs.Set(goStructName, structInfo{
						goName:  goStructName,
						cName:   structName,
						fields:  fieldDefs,
						methods: methodDefs,
						source:  sourceFile,
						lineNo:  int(tag.Position().Line),
					})
					for _, it := range innerTypes {
						it.source = sourceFile
						result.Structs.Set(it.goName, it)
					}
				}

			case *cc.EnumType:
				enumTag := t.Tag()
				enumName := string(enumTag.Src())
				isTypedefEnum := false
				if enumName == "" {
					if decl.InitDeclaratorList != nil {
						for id := decl.InitDeclaratorList; id != nil; id = id.InitDeclaratorList {
							if id.InitDeclarator == nil || id.InitDeclarator.Declarator == nil {
								continue
							}
							tdName := id.InitDeclarator.Declarator.Name()
							if tdName == "" {
								continue
							}
							enumName = string(tdName)
							isTypedefEnum = true
							break
						}
					}
					if enumName == "" {
						continue
					}
				}
				goEnumName := cTagToGoName(enumName)
				members := safemap.New[string, enumMember]()
				var memberOrder []string
				for _, e := range t.Enumerators() {
					if e == nil {
						continue
					}
					cMemberName := e.Token.SrcStr()
					goMemberName := cEnumMemberToGoName(cMemberName)
					valStr, intVal, hasVal := extractEnumeratorValue(e)
					members.Set(goMemberName, enumMember{
						goName: goMemberName,
						value:  valStr,
						intVal: intVal,
						hasVal: hasVal,
					})
					memberOrder = append(memberOrder, goMemberName)
				}
				result.Enums.Set(goEnumName, enumInfo{
					goName:      goEnumName,
					cName:       enumName,
					members:     members,
					memberOrder: memberOrder,
					source:      sourceFile,
					lineNo:      int(enumTag.Position().Line),
				})
				if isTypedefEnum {
					result.Typedefs.Delete(goEnumName)
				}

			case *cc.UnionType:
				unionTag := t.Tag()
				unionName := string(unionTag.Src())
				if unionName != "" {
					goUnionName := cTagToGoName(unionName)
					unionSize := t.Size()
					if unionFields := generateUnionFields(t, unionSize); unionFields != "" {
						var nestedInner []structInfo
						for j := 0; j < t.NumFields(); j++ {
							if uf := t.FieldByIndex(j); uf != nil {
								if usv, ok := uf.Type().(*cc.StructType); ok {
									nestedTag := usv.Tag()
									nestedTagStr := string(nestedTag.Src())
									if nestedTagStr != "" && !hasBitfields(usv) {
										nestedF, nestedM, _ := generateStructFields(usv, mapCTypeToGo(uf.Type()), false)
										nestedInner = append(nestedInner, structInfo{
											goName:  mapCTypeToGo(uf.Type()),
											cName:   nestedTagStr,
											fields:  nestedF,
											methods: nestedM,
										})
									}
								}
							}
						}
						result.Structs.Set(goUnionName+"_", structInfo{
							goName: goUnionName + "_",
							cName:  unionName,
							fields: unionFields,
							source: sourceFile,
						})
						for _, it := range nestedInner {
							it.source = sourceFile
							result.Structs.Set(it.goName, it)
						}
					}
					fieldDefs := fmt.Sprintf("\t%s\n", unionAlignedFFIType(t, unionSize))
					result.Structs.Set(goUnionName, structInfo{
						goName: goUnionName,
						cName:  unionName,
						fields: fieldDefs,
						source: sourceFile,
					})
				}

				if !isMsvcTypes && decl.InitDeclaratorList != nil {
					for id := decl.InitDeclaratorList; id != nil; id = id.InitDeclaratorList {
						if id.InitDeclarator == nil || id.InitDeclarator.Declarator == nil {
							continue
						}
						tdName := id.InitDeclarator.Declarator.Name()
						if tdName == "" {
							continue
						}
						tdType := id.InitDeclarator.Declarator.Type()
						if tdType == nil {
							continue
						}
						goTdName := cTagToGoName(tdName)
						switch vt := tdType.(type) {
						case *cc.EnumType:
							enumTag := vt.Tag()
							tagStr := string(enumTag.Src())
							if tagStr != "" {
								result.Typedefs.Set(goTdName, typedefInfo{
									goName: goTdName,
									goType: cTagToGoName(tagStr),
									source: sourceFile,
								})
							}
						case *cc.StructType:
							structTag := vt.Tag()
							tagStr := string(structTag.Src())
							if tagStr != "" {
								result.Typedefs.Set(goTdName, typedefInfo{
									goName: goTdName,
									goType: cTagToGoName(tagStr),
									source: sourceFile,
								})
							} else {
								isPacked := packedFiles[baseName]
								fieldDefs, methodDefs, innerTypes := generateStructFields(vt, goTdName, isPacked)
								result.Structs.Set(goTdName, structInfo{
									goName:  goTdName,
									cName:   tdName,
									fields:  fieldDefs,
									methods: methodDefs,
									source:  sourceFile,
								})
								for _, it := range innerTypes {
									it.source = sourceFile
									result.Structs.Set(it.goName, it)
								}
							}
						case *cc.UnionType:
							backingName, backingType, bfs := extractUnionBitfields(vt)
							if backingName != "" && len(bfs) > 0 {
								fieldDefs := fmt.Sprintf("\t%s %s\n", backingName, backingType)
								var methodDefs []string
								for _, bf := range bfs {
									methodDefs = append(methodDefs, generateBitfieldGetterSetter(bf, goTdName, backingName, backingType))
								}
								result.Structs.Set(goTdName, structInfo{
									goName:  goTdName,
									cName:   tdName,
									fields:  fieldDefs,
									methods: strings.Join(methodDefs, "\n"),
									source:  sourceFile,
								})
							} else {
								unionSize := vt.Size()
								fieldDefs := fmt.Sprintf("\t%s\n", unionAlignedFFIType(vt, unionSize))
								result.Structs.Set(goTdName, structInfo{
									goName: goTdName,
									cName:  tdName,
									fields: fieldDefs,
									source: sourceFile,
								})
							}
						case *cc.PointerType:
							elem := vt.Elem()
							if ut, ok := elem.(*cc.UnionType); ok {
								elemTag := ut.Tag()
								tagStr := string(elemTag.Src())
								if tagStr != "" {
									result.Typedefs.Set(goTdName, typedefInfo{
										goName: goTdName,
										goType: "*" + cTagToGoName(tagStr),
										source: sourceFile,
									})
								}
							}
						}
					}
				}

			}

			if !isMsvcTypes && decl.InitDeclaratorList != nil {
				for id := decl.InitDeclaratorList; id != nil; id = id.InitDeclaratorList {
					if id.InitDeclarator == nil || id.InitDeclarator.Declarator == nil {
						continue
					}
					declName := id.InitDeclarator.Declarator.Name()
					if declName == "" {
						continue
					}
					declType := id.InitDeclarator.Declarator.Type()
					if declType == nil {
						continue
					}
					goTypeName := mapCTypedefToGo(declName, declType)
					goDeclName := cTagToGoName(declName)
					if goTypeName == "" {
						switch vt := declType.(type) {
						case *cc.StructType:
							isPacked := packedFiles[baseName]
							fieldDefs, methodDefs, innerTypes := generateStructFields(vt, goDeclName, isPacked)
							result.Structs.Set(goDeclName, structInfo{
								goName:  goDeclName,
								cName:   declName,
								fields:  fieldDefs,
								methods: methodDefs,
								source:  sourceFile,
							})
							for _, it := range innerTypes {
								it.source = sourceFile
								result.Structs.Set(it.goName, it)
							}
						case *cc.UnionType:
							backingName, backingType, bfs := extractUnionBitfields(vt)
							if backingName != "" && len(bfs) > 0 {
								fieldDefs := fmt.Sprintf("\t%s %s\n", backingName, backingType)
								var methodDefs []string
								for _, bf := range bfs {
									methodDefs = append(methodDefs, generateBitfieldGetterSetter(bf, goDeclName, backingName, backingType))
								}
								result.Structs.Set(goDeclName, structInfo{
									goName:  goDeclName,
									cName:   declName,
									fields:  fieldDefs,
									methods: strings.Join(methodDefs, "\n"),
									source:  sourceFile,
								})
							} else {
								unionSize := vt.Size()
								fieldDefs := fmt.Sprintf("\t%s\n", unionAlignedFFIType(vt, unionSize))
								result.Structs.Set(goDeclName, structInfo{
									goName: goDeclName,
									cName:  declName,
									fields: fieldDefs,
									source: sourceFile,
								})
							}
						default:
						}
						continue
					}
					if goDeclName == goTypeName {
						continue
					}
					result.Typedefs.Set(goDeclName, typedefInfo{
						goName: goDeclName,
						goType: goTypeName,
						source: sourceFile,
						isFunc: strings.HasPrefix(goTypeName, "func("),
					})
				}
			}
		}
	}

	fnMacroNames := make(map[string]string)
	constMacroNames := make(map[string]string)
	for macroName, macro := range ast.Macros {
		if macro.IsFnLike {
			fnMacroNames[macroName] = cMacroNameToGoName(macroName)
		} else if macro.IsConst {
			goConstName := cMacroNameToGoName(macroName)
			if goConstName != macroName {
				constMacroNames[macroName] = goConstName
			}
		}
	}

	for name, m := range ast.Macros {
		pos := m.Name.Position()
		sourceFile := filepath.Base(pos.Filename)
		if !strings.HasSuffix(sourceFile, ".h") {
			continue
		}
		found := slices.Contains(bc.HeaderOrder, sourceFile)
		if !found && bc.SingleFile {
			_, found = headerMap[sourceFile]
			if !found {
				for k := range headerMap {
					if filepath.Base(k) == sourceFile {
						found = true
						break
					}
				}
			}
		}
		if !found {
			continue
		}

		if m.IsFnLike {
			replList := m.ReplacementList()
			if len(replList) == 0 {
				continue
			}
			var bodyStr string
			for _, tok := range replList {
				bodyStr += string(tok.Src())
			}
			bodyStr = strings.TrimSpace(bodyStr)

			var goParamList []string
			var goParamNames []string
			var cParamNames []string
			for _, p := range m.Params {
				cPName := string(p.Src())
				cParamNames = append(cParamNames, cPName)
				goPName := strings.ToUpper(cPName[:1]) + cPName[1:]
				goParamList = append(goParamList, goPName+" uint32")
				goParamNames = append(goParamNames, goPName)
			}

			goName := cMacroNameToGoName(name)
			goBody := convertCMacroBodyToGo(bodyStr, cParamNames, goParamNames)
			result.FnMacros.Set(goName, macroConstInfo{
				goName:   goName,
				goType:   "func",
				value:    bodyStr,
				source:   sourceFile,
				cName:    name,
				goParams: strings.Join(goParamList, ", "),
				goBody:   goBody,
			})
			continue
		}

		if !m.IsConst {
			continue
		}

		replList := m.ReplacementList()
		if len(replList) == 0 {
			continue
		}

		var valStr string
		for _, tok := range replList {
			valStr += string(tok.Src())
		}
		valStr = strings.TrimSpace(valStr)

		if !isValidGoMacroValue(valStr) {
			continue
		}

		goVal := cleanCMacroValue(valStr)

		isVar := false
		if strings.Contains(goVal, "unsafe.Sizeof(") {
			isVar = true
		}

		identRe := regexp.MustCompile(`[A-Za-z_][A-Za-z0-9_]*`)
		identsInVal := identRe.FindAllString(goVal, -1)
		identsSet := make(map[string]bool, len(identsInVal))
		for _, id := range identsInVal {
			identsSet[id] = true
		}

		for id := range identsSet {
			if goFnName, ok := fnMacroNames[id]; ok {
				if strings.Contains(goVal, id+"(") || strings.Contains(goVal, id+" (") {
					isVar = true
					goVal = strings.ReplaceAll(goVal, id+"(", goFnName+"(")
					goVal = strings.ReplaceAll(goVal, id+" (", goFnName+"(")
				}
			}
			if goCName, ok := constMacroNames[id]; ok {
				re := regexp.MustCompile(`\b` + regexp.QuoteMeta(id) + `\b`)
				goVal = re.ReplaceAllString(goVal, goCName)
			}
		}

		goType := determineMacroGoType(valStr)
		if isVar && strings.Contains(goVal, "unsafe.Sizeof(") {
			goType = "int"
		}

		goName := cMacroNameToGoName(name)
		result.Macros.Set(goName, macroConstInfo{
			goName:  goName,
			goType:  goType,
			value:   valStr,
			goValue: goVal,
			source:  sourceFile,
			isVar:   isVar,
			lineNo:  int(pos.Line),
		})
	}

	headerOrder := bc.HeaderOrder
	if bc.SingleFile && len(headerOrder) > 0 {
		headerOrder = headerOrder[:1]
	}
	for _, fileName := range headerOrder {
		if _, ok := headerMap[fileName]; !ok {
			continue
		}
		baseName := strings.TrimSuffix(fileName, ".h")
		baseName = filepath.Base(baseName)
		var content strings.Builder
		content.WriteString(fmt.Sprintf("package %s\n\n", bc.PackageName))

		hasContent := false
		structCount, enumCount, tdCount := 0, 0, 0
		definedConsts := make(map[string]bool)

		sourceMatch := func(src string) bool {
			if bc.SingleFile {
				if src == bc.HeaderOrder[0] {
					return true
				}
				_, inHeaders := headerMap[src]
				if inHeaders {
					return true
				}
				for k := range headerMap {
					if filepath.Base(k) == src {
						return true
					}
				}
				if src == "unknown.h" || src == "" || src == "<msvc_types>" {
					return true
				}
				return false
			}
			return src == fileName
		}
		var aliasTypes []typedefInfo
		var funcTypes []typedefInfo
		var skipStructNames map[string]bool
		for _, td := range result.Typedefs.Range() {
			if !sourceMatch(td.source) {
				continue
			}
			if result.Enums.Has(td.goName) {
				continue
			}
			if result.Structs.Has(td.goName) {
				if isGoPrimitiveType(td.goType) {
					if skipStructNames == nil {
						skipStructNames = make(map[string]bool)
					}
					skipStructNames[td.goName] = true
				} else {
					continue
				}
			}
			tdCount++
			hasContent = true
			if td.isFunc {
				funcTypes = append(funcTypes, td)
			} else {
				aliasTypes = append(aliasTypes, td)
			}
		}
		if len(aliasTypes) > 0 {
			content.WriteString(fmt.Sprintf("// Source: %s\n", baseName+".h"))
			content.WriteString("type (\n")
			for _, td := range aliasTypes {
				content.WriteString(fmt.Sprintf("\t%s = %s\n", td.goName, td.goType))
			}
			content.WriteString(")\n\n")
		}
		for _, td := range funcTypes {
			content.WriteString(fmt.Sprintf("type %s %s\n\n", td.goName, td.goType))
		}

		for _, ei := range result.Enums.Range() {
			if !sourceMatch(ei.source) {
				continue
			}
			skip := false
			for _, name := range ei.memberOrder {
				if definedConsts[name] {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
			for _, name := range ei.memberOrder {
				definedConsts[name] = true
			}
			enumCount++
			hasContent = true
			content.WriteString(fmt.Sprintf("// Source: %s:%d -> %s\n", filepath.Base(ei.source), ei.lineNo, ei.cName))
			hasNegative := false
			for _, name := range ei.memberOrder {
				if m, ok := ei.members.Get(name); ok && m.hasVal && strings.HasPrefix(m.value, "-") {
					hasNegative = true
					break
				}
			}
			if hasNegative {
				content.WriteString(fmt.Sprintf("type %s int32\n\n", ei.goName))
			} else {
				content.WriteString(fmt.Sprintf("type %s uint32\n\n", ei.goName))
			}

			// Special handling for EXCEPTION_VECTORS to use iota
			if ei.goName == "EXCEPTION_VECTORS" {
				content.WriteString("const (\n")
				content.WriteString("\tExceptionVectorDivideError EXCEPTION_VECTORS = iota\n")
				// Implicit iota for next 31 entries
				for _, name := range ei.memberOrder[1:32] {
					content.WriteString(fmt.Sprintf("\t%s\n", name))
				}
				// Explicit values for Windows-specific entries
				for _, name := range ei.memberOrder[32:] {
					m, _ := ei.members.Get(name)
					content.WriteString(fmt.Sprintf("\t%s EXCEPTION_VECTORS = %s\n", name, m.value))
				}
				content.WriteString(")\n\n")
			} else if len(ei.memberOrder) > 0 {
				content.WriteString("const (\n")
				isSequential := true
				seqStart := int64(0)
				for i, name := range ei.memberOrder {
					if m, ok := ei.members.Get(name); ok && m.hasVal {
						if i == 0 {
							seqStart = m.intVal
						} else if m.intVal != seqStart+int64(i) {
							isSequential = false
							break
						}
					} else {
						isSequential = false
						break
					}
				}
				if isSequential {
					if seqStart == 0 {
						content.WriteString(fmt.Sprintf("\t%s %s = iota\n", ei.memberOrder[0], ei.goName))
					} else {
						content.WriteString(fmt.Sprintf("\t%s %s = %d + iota\n", ei.memberOrder[0], ei.goName, seqStart))
					}
					for _, name := range ei.memberOrder[1:] {
						content.WriteString(fmt.Sprintf("\t%s\n", name))
					}
				} else {
					hasExplicitValues := false
					for _, name := range ei.memberOrder[1:] {
						if m, ok := ei.members.Get(name); ok && m.hasVal {
							hasExplicitValues = true
							break
						}
					}
					if hasExplicitValues {
						for _, name := range ei.memberOrder {
							if m, ok := ei.members.Get(name); ok && m.hasVal {
								content.WriteString(fmt.Sprintf("\t%s %s = %s\n", name, ei.goName, m.value))
							} else {
								content.WriteString(fmt.Sprintf("\t%s\n", name))
							}
						}
					} else if firstMember, ok := ei.members.Get(ei.memberOrder[0]); ok && firstMember.hasVal && firstMember.intVal != 0 {
						content.WriteString(fmt.Sprintf("\t%s %s = %d + iota\n", ei.memberOrder[0], ei.goName, firstMember.intVal))
						for _, name := range ei.memberOrder[1:] {
							content.WriteString(fmt.Sprintf("\t%s\n", name))
						}
					} else {
						content.WriteString(fmt.Sprintf("\t%s %s = iota\n", ei.memberOrder[0], ei.goName))
						for _, name := range ei.memberOrder[1:] {
							content.WriteString(fmt.Sprintf("\t%s\n", name))
						}
					}
				}
				content.WriteString(")\n\n")
			}
			generateEnumStringer(&content, ei)
		}

		usedNames := make(map[string]bool)
		for _, si := range result.Structs.Range() {
			if !sourceMatch(si.source) {
				continue
			}
			if skipStructNames != nil && skipStructNames[si.goName] {
				continue
			}
			structCount++
			hasContent = true
			src := si.source
			if si.originalSource != "" {
				src = si.originalSource
			}
			goName := si.goName
			reserved := map[string]bool{
				"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
				"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
				"uintptr": true, "string": true, "bool": true, "byte": true, "rune": true,
				"float32": true, "float64": true,
			}
			for reserved[goName] || usedNames[goName] {
				goName = goName + "_"
			}
			usedNames[goName] = true
			content.WriteString(fmt.Sprintf("// Source: %s:%d -> %s\n", src, si.lineNo, si.cName))
			content.WriteString(fmt.Sprintf("type %s struct {\n%s}\n", goName, si.fields))
			if si.methods != "" {
				content.WriteString(si.methods)
			}
			if isFFIUnion(si.fields) && !strings.HasSuffix(si.goName, "_") {
				generateUnionAccessorMethods(&content, si.goName, result.Structs)
			}
			content.WriteString("\n")
		}

		var constItems []macroConstInfo
		var varItems []macroConstInfo
		for _, mc := range result.Macros.Range() {
			if !sourceMatch(mc.source) {
				continue
			}
			if isTypeAliasMacro(mc.goValue) {
				continue
			}
			if mc.isVar {
				varItems = append(varItems, mc)
			} else {
				constItems = append(constItems, mc)
			}
		}
		slices.SortFunc(constItems, func(a, b macroConstInfo) int { return a.lineNo - b.lineNo })

		for _, mc := range result.FnMacros.Range() {
			if !sourceMatch(mc.source) {
				continue
			}
			if !isValidGoFnMacroBody(mc.value) {
				continue
			}
			if isCStyleCastMacro(mc.goBody) {
				continue
			}
			cleanedBody := cleanCMacroValue(mc.goBody)
			if !isValidGoFnMacroBody(cleanedBody) {
				continue
			}
			hasContent = true
			simplifiedParams, simplifiedBody, updatedVarItems := simplifyFnMacro(mc, varItems)
			isBool := isBoolExpr(cleanedBody)
			if simplifiedParams != "" {
				body := cleanCMacroValue(resolveMacroValueNames(simplifiedBody))
				if isBool {
					content.WriteString(fmt.Sprintf("func %s(%s) bool {\n\treturn %s\n}\n\n", mc.goName, simplifiedParams, convertCBoolExpr(body)))
				} else {
					content.WriteString(fmt.Sprintf("func %s(%s) uint32 {\n\treturn uint32(%s)\n}\n\n", mc.goName, simplifiedParams, body))
				}
				varItems = updatedVarItems
			} else {
				body := cleanCMacroValue(resolveMacroValueNames(cleanedBody))
				if isBool {
					content.WriteString(fmt.Sprintf("func %s(%s) bool {\n\treturn %s\n}\n\n", mc.goName, mc.goParams, convertCBoolExpr(body)))
				} else {
					content.WriteString(fmt.Sprintf("func %s(%s) uint32 {\n\treturn uint32(%s)\n}\n\n", mc.goName, mc.goParams, body))
				}
			}
		}

		if len(constItems) > 0 {
			hasContent = true
			isErrorCodeGroup := strings.HasPrefix(baseName, "ErrorCodes") || func() bool {
				for _, mc := range constItems {
					if (strings.HasPrefix(mc.goName, "DebuggerError") || strings.HasPrefix(mc.goName, "DebuggerOperation")) && strings.Contains(mc.source, "ErrorCodes") {
						return true
					}
				}
				return false
			}()
			if isErrorCodeGroup {
				var errorCodeItems []macroConstInfo
				var nonErrorCodeItems []macroConstInfo
				for _, mc := range constItems {
					if strings.HasPrefix(mc.goName, "DebuggerError") || strings.HasPrefix(mc.goName, "DebuggerOperation") {
						errorCodeItems = append(errorCodeItems, mc)
					} else {
						nonErrorCodeItems = append(nonErrorCodeItems, mc)
					}
				}
				content.WriteString("// Source: ErrorCodes.h -> Error codes\n")
				generateErrorCodeEnum(&content, errorCodeItems, baseName)
				propagateConstTypes(nonErrorCodeItems)
				content.WriteString("\n// Source: " + baseName + ".h -> Macro constants\nconst (\n")
				for _, mc := range nonErrorCodeItems {
					content.WriteString(fmt.Sprintf("\t%s %s = %s\n", mc.goName, mc.goType, resolveMacroValueNames(mc.goValue)))
				}
				content.WriteString(")\n\n")
			} else {
				propagateConstTypes(constItems)
				content.WriteString(fmt.Sprintf("// Source: %s -> Macro constants\n", baseName+".h"))
				content.WriteString("const (\n")
				for _, mc := range constItems {
					content.WriteString(fmt.Sprintf("\t%s %s = %s\n", mc.goName, mc.goType, resolveMacroValueNames(mc.goValue)))
				}
				content.WriteString(")\n\n")
			}
		}

		if len(varItems) > 0 {
			hasContent = true
			propagateConstTypes(varItems)
			content.WriteString(fmt.Sprintf("// Source: %s -> Macro variables\n", baseName+".h"))
			content.WriteString("var (\n")
			for _, mc := range varItems {
				val := resolveMacroValueNames(mc.goValue)
				if strings.Contains(val, "unsafe.Sizeof(") {
					val = wrapConstRefsForInt(val)
				}
				content.WriteString(fmt.Sprintf("\t%s %s = %s\n", mc.goName, mc.goType, val))
			}
			content.WriteString(")\n\n")
		}

		if !hasContent && len(bc.ExtraConstants) == 0 {
			continue
		}

		if len(bc.ExtraConstants) > 0 {
			content.WriteString("// Extra constants\n")
			content.WriteString("const (\n")
			for name, ec := range bc.ExtraConstants {
				content.WriteString(fmt.Sprintf("\t%s %s = %s\n", name, ec.GoType, ec.GoValue))
			}
			content.WriteString(")\n\n")
		}

		finalContent := addImports(content.String(), bc.PackageName, result.Imports)

		outputFile := filepath.Join(bc.OutputDir, baseName+".go")
		err := os.WriteFile(outputFile, []byte(finalContent), 0o644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", outputFile, err)
			continue
		}
		fmt.Printf("Generated: %s (%d structs, %d enums, %d typedefs)\n", outputFile, structCount, enumCount, tdCount)
	}

	fmt.Printf("\n✅ %s binding generation completed!\n", bc.PackageName)

	if bc.BindDll {
		generateDllBinding(t, ast, bc, result.Typedefs, result.Enums)
	}
}

func addImports(content string, packageName string, extraImports map[string]bool) string {
	imports := make(map[string]bool)
	maps.Copy(imports, extraImports)
	if strings.Contains(content, "fmt.Sprintf") {
		imports[`"fmt"`] = true
	}
	if strings.Contains(content, "unsafe.") {
		imports[`"unsafe"`] = true
	}
	if strings.Contains(content, "time.Duration") {
		imports[`"time"`] = true
	}
	if strings.Contains(content, "bits.ReverseBytes") || strings.Contains(content, "bits.Reverse(") {
		imports[`"math/bits"`] = true
	}
	if len(imports) == 0 {
		return content
	}
	var sortedImports []string
	for imp := range imports {
		sortedImports = append(sortedImports, imp)
	}
	sort.Strings(sortedImports)
	var importBlock strings.Builder
	importBlock.WriteString("import (\n")
	for _, imp := range sortedImports {
		importBlock.WriteString("\t" + imp + "\n")
	}
	importBlock.WriteString(")\n\n")
	return strings.Replace(content, fmt.Sprintf("package %s\n\n", packageName), fmt.Sprintf("package %s\n\n", packageName)+importBlock.String(), 1)
}

func generateEnumStringer(content *strings.Builder, ei enumInfo) {
	if len(ei.memberOrder) == 0 {
		return
	}
	receiver := strings.ToLower(ei.goName[:1])
	content.WriteString(fmt.Sprintf("func (%s %s) String() string {\n", receiver, ei.goName))
	content.WriteString(fmt.Sprintf("\tswitch %s {\n", receiver))
	seenValues := make(map[int64]bool)
	for _, name := range ei.memberOrder {
		m, _ := ei.members.Get(name)
		if seenValues[m.intVal] {
			continue
		}
		seenValues[m.intVal] = true
		humanName := cEnumMemberToHumanString(m.goName, ei.goName)
		content.WriteString(fmt.Sprintf("\tcase %s:\n\t\treturn %q\n", m.goName, humanName))
	}
	content.WriteString("\tdefault:\n")
	content.WriteString(fmt.Sprintf("\t\treturn fmt.Sprintf(\"%s(0x%%X)\", uint32(%s))\n", ei.goName, receiver))
	content.WriteString("\t}\n")
	content.WriteString("}\n\n")
}

func cEnumMemberToHumanString(goName, typeName string) string {
	prefix := typeName + "_"
	s := goName
	if after, ok := strings.CutPrefix(goName, prefix); ok {
		s = after
	} else {
		for i := 0; i < len(typeName) && i < len(goName); i++ {
			if typeName[i] == '_' && goName[i] == '_' {
				s = goName[i+1:]
				break
			}
			if typeName[i] != goName[i] {
				break
			}
		}
	}
	var result strings.Builder
	for i, ch := range s {
		if i > 0 {
			prev := rune(s[i-1])
			if ch >= 'A' && ch <= 'Z' {
				if prev >= 'a' && prev <= 'z' {
					result.WriteByte(' ')
				} else if prev >= '0' && prev <= '9' {
					result.WriteByte(' ')
				} else if i+1 < len(s) && s[i+1] >= 'a' && s[i+1] <= 'z' {
					result.WriteByte(' ')
				}
			} else if ch >= '0' && ch <= '9' {
				if prev >= 'a' && prev <= 'z' {
					result.WriteByte(' ')
				}
			}
		}
		result.WriteRune(ch)
	}
	return result.String()
}

func isBoolExpr(body string) bool {
	boolOps := []string{"==", "!=", "<=", ">=", "&&", "||"}
	for _, op := range boolOps {
		if strings.Contains(body, op) {
			return true
		}
	}
	notRe := regexp.MustCompile(`!\s*\(`)
	return notRe.MatchString(body)
}

func convertCBoolExpr(body string) string {
	doubleNotRe := regexp.MustCompile(`!!\s*\(`)
	if doubleNotRe.MatchString(body) {
		body = doubleNotRe.ReplaceAllStringFunc(body, func(match string) string {
			return "!!("
		})
		for {
			idx := strings.Index(body, "!!(")
			if idx == -1 {
				break
			}
			parenStart := idx + 2
			depth := 0
			end := -1
			for i := parenStart; i < len(body); i++ {
				if body[i] == '(' {
					depth++
				} else if body[i] == ')' {
					depth--
					if depth == 0 {
						end = i
						break
					}
				}
			}
			if end == -1 {
				break
			}
			inner := body[parenStart+1 : end]
			replacement := "(" + inner + ") != 0"
			body = body[:idx] + replacement + body[end+1:]
		}
		return body
	}
	for {
		idx := strings.Index(body, "!(")
		if idx == -1 {
			idx = strings.Index(body, "! (")
		}
		if idx == -1 {
			break
		}
		start := idx
		parenStart := strings.Index(body[start:], "(") + start
		depth := 0
		end := -1
		for i := parenStart; i < len(body); i++ {
			if body[i] == '(' {
				depth++
			} else if body[i] == ')' {
				depth--
				if depth == 0 {
					end = i
					break
				}
			}
		}
		if end == -1 {
			break
		}
		inner := body[parenStart+1 : end]
		replacement := "(" + inner + ") == 0"
		body = body[:start] + replacement + body[end+1:]
	}
	return body
}

func isTypeAliasMacro(val string) bool {
	trimmed := strings.TrimSpace(val)
	if regexp.MustCompile(`\b[0-9]+\b`).MatchString(trimmed) {
		return false
	}
	if strings.Contains(trimmed, "(") && strings.Contains(trimmed, ")") {
		return false
	}
	lower := strings.ToLower(trimmed)
	typeAliasPatterns := []string{
		"unsigned", "long", "int", "short", "void",
		"char", "wchar_t", "intptr_t", "uintptr_t",
		"struct", "union", "enum",
	}
	for _, p := range typeAliasPatterns {
		if strings.Contains(lower, p) {
			return true
		}
	}
	if strings.ContainsAny(trimmed, "[]*") {
		return true
	}
	return false
}

func isFFIUnion(fields string) bool {
	fields = strings.TrimSpace(fields)
	return strings.HasPrefix(fields, "_ [") && strings.HasSuffix(fields, "byte")
}

func generateUnionAccessorMethods(content *strings.Builder, unionGoName string, structs interface {
	Get(string) (structInfo, bool)
},
) {
	accessorName := unionGoName + "_"
	acc, ok := structs.Get(accessorName)
	if !ok {
		return
	}
	lines := strings.SplitSeq(strings.TrimSpace(acc.fields), "\n")
	for line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		fieldName := parts[0]
		fieldType := parts[1]
		content.WriteString(fmt.Sprintf("func (u *%s) %s() *%s { return (*%s)(unsafe.Pointer(u)) }\n",
			unionGoName, fieldName, fieldType, fieldType))
		if nested, ok := structs.Get(fieldType + "_"); ok && !isFFIUnion(nested.fields) {
			content.WriteString(fmt.Sprintf("func (u *%s) %s() *%s { return (*%s)(unsafe.Pointer(u)) }\n",
				unionGoName, fieldName+"Fields", fieldType+"_", fieldType+"_"))
		}
	}
}

func isValidGoFnMacroBody(body string) bool {
	if body == "" {
		return false
	}
	cMarkers := []string{
		";", "{", "}", "#", "?", "->", "/=",
		"*(const", "__attribute", "assert(", "static", "sizeof",
		"ZYAN_", "ZYDIS_", "ZyanU16", "ZyanU32", "ZyanU64",
		"ZyanI16", "ZyanI32", "ZyanI64",
		"uNsigNed", "constZyan", "inline", "(void)",
		"_byteswap_ulong", "_byteswap_uint64", "_byteswap_ushort",
		"UINT64_C", "INT64_C", "__has_include", "__has_builtin",
		"XED_STATIC_CAST", "XED_REINTERPRET_CAST", "XED_CAST",
	}
	for _, m := range cMarkers {
		if strings.Contains(body, m) {
			return false
		}
	}
	uLiteralRe := regexp.MustCompile(`\bu["']`)
	if uLiteralRe.MatchString(body) {
		return false
	}
	if len(body) > 0 && body[0] == ':' {
		return false
	}
	if isSimpleNumber(body) {
		return false
	}
	largeHexRe := regexp.MustCompile(`0x[0-9A-Fa-f]{9,}`)
	if largeHexRe.MatchString(body) {
		return false
	}
	return true
}

func isCStyleCastMacro(body string) bool {
	castPattern := regexp.MustCompile(`\([A-Z][A-Za-z0-9_]*\)\s*\(`)
	return castPattern.MatchString(body)
}

func isSimpleNumber(s string) bool {
	s = strings.TrimSpace(s)
	if s == "" {
		return false
	}
	_, err := strconv.ParseInt(s, 0, 64)
	if err == nil {
		return true
	}
	_, err = strconv.ParseUint(s, 0, 64)
	return err == nil
}

func isValidGoMacroValue(val string) bool {
	if val == "" {
		return false
	}
	if val == "void" || val == "0ull" {
		return false
	}
	if strings.HasPrefix(val, `"`) || strings.HasPrefix(val, `'`) {
		return false
	}
	if strings.Contains(val, "__DATE__") || strings.Contains(val, "__TIME__") {
		return false
	}
	if strings.Contains(val, "&&") || strings.Contains(val, "||") || strings.Contains(val, "?") {
		return false
	}
	if strings.Contains(val, "=='") || strings.Contains(val, "!='") {
		return false
	}
	if strings.Contains(val, "BUILD_MONTH") {
		return false
	}
	if strings.Contains(val, "__attribute") || strings.Contains(val, "static") || strings.Contains(val, "void*") || strings.Contains(val, "(void") {
		return false
	}
	if strings.Contains(val, "inline") || strings.Contains(val, "noinline") || strings.Contains(val, "deprecated") {
		return false
	}
	if strings.Contains(val, "fallthrough") || strings.Contains(val, "Dllimport") || strings.Contains(val, "dllexport") || strings.Contains(val, "DLLIMPORT") || strings.Contains(val, "DLLEXPORT") || strings.Contains(val, "__declspec") || strings.Contains(val, "__assume") {
		return false
	}
	if strings.Contains(val, "{") || strings.Contains(val, "}") || strings.Contains(val, "(void)") {
		return false
	}
	if strings.Contains(val, "ZyanU64") || strings.Contains(val, "ZyanU32") || strings.Contains(val, "ZyanU16") {
		return false
	}
	cStdMacros := []string{
		"INT8_MIN", "INT16_MIN", "INT32_MIN", "INT64_MIN",
		"INT8_MAX", "INT16_MAX", "INT32_MAX", "INT64_MAX",
		"UINT8_MAX", "UINT16_MAX", "UINT32_MAX", "UINT64_MAX",
		"Int8Min", "Int16Min", "Int32Min", "Int64Min",
		"Int8Max", "Int16Max", "Int32Max", "Int64Max",
		"Uint8Max", "Uint16Max", "Uint32Max", "Uint64Max",
	}
	for _, m := range cStdMacros {
		if strings.Contains(val, m) {
			return false
		}
	}
	charLiteralRe := regexp.MustCompile(`'[^']'`)
	if charLiteralRe.MatchString(val) {
		return false
	}
	if strings.Contains(val, "[") && strings.Contains(val, "]") {
		if strings.Contains(val, "__DATE__") || strings.Contains(val, "__TIME__") || strings.Contains(val, "BUILD_MONTH") {
			return false
		}
	}
	return true
}

func cleanCMacroValue(val string) string {
	result := val
	result = strings.ReplaceAll(result, "~", "^")
	re := regexp.MustCompile(`([0-9A-Fa-f])([UuLl]+|[Uu]?[Ii][0-9]+)\b`)
	result = re.ReplaceAllString(result, "$1")
	sizeofRe := regexp.MustCompile(`sizeof\(([A-Za-z_]\w*)\)`)
	result = sizeofRe.ReplaceAllStringFunc(result, func(match string) string {
		sub := sizeofRe.FindStringSubmatch(match)
		if len(sub) >= 2 {
			cType := strings.TrimSpace(sub[1])
			goType := mapCTypeNameToGoForSizeof(cType)
			if isGoPrimitiveType(goType) {
				return fmt.Sprintf("int(unsafe.Sizeof(%s(0)))", goType)
			}
			return fmt.Sprintf("int(unsafe.Sizeof(%s{}))", goType)
		}
		return match
	})
	return result
}

func mapCTypeNameToGoForSizeof(cType string) string {
	switch cType {
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
	case "LONGLONG", "INT64", "Time_t", "time_t":
		return "int64"
	case "long long", "long long int", "Longlong", "Longlongint", "LonglongInt":
		return "int64"
	case "longlong", "longlongint":
		return "int64"
	case "long long unsigned", "unsigned long long":
		return "uint64"
	case "Longlongunsigned", "LonglongUnsigned", "Unsignedlonglong", "UnsignedLongLong":
		return "uint64"
	case "longlongunsigned":
		return "uint64"
	case "long unsigned int", "unsigned long int":
		return "uint32"
	case "Longunsignedint", "LongUnsignedInt", "Unsignedlongint", "UnsignedLongInt":
		return "uint32"
	case "longunsignedint":
		return "uint32"
	case "DEBUGGER_TEST_QUERY_BUFFER":
		return "DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER"
	case "WCHAR", "wchar_t":
		return "rune"
	case "SIZE_T", "HANDLE", "PVOID", "LPVOID":
		return "uintptr"
	default:
		return cTagToGoName(cType)
	}
}

func resolveMacroValueNames(val string) string {
	result := val
	replacements := map[string]string{
		"PAGE_SIZE":                      "PageSize",
		"MaxSerialPacketSize":            "MaxSerialPacketSize",
		"PacketChunkSize":                "PacketChunkSize",
		"MaximumPacketsCapacity":         "MaximumPacketsCapacity",
		"MaximumPacketsCapacityPriority": "MaximumPacketsCapacityPriority",
	}
	for cName, goName := range replacements {
		result = strings.ReplaceAll(result, cName, goName)
	}
	typeReplacements := map[string]string{
		"BYTE":             "uint8",
		"WORD":             "uint16",
		"DWORD":            "uint32",
		"QWORD":            "uint64",
		"BOOL":             "int32",
		"UINT":             "uint32",
		"ULONG":            "uint32",
		"LONG":             "int32",
		"INT":              "int32",
		"UCHAR":            "uint8",
		"USHORT":           "uint16",
		"CHAR":             "int8",
		"SHORT":            "int16",
		"LONGLONG":         "int64",
		"ULONGLONG":        "uint64",
		"DWORD64":          "uint64",
		"ULONG64":          "uint64",
		"Longlongunsigned": "uint64",
		"ZyanStatus":       "uint32",
	}
	for cName, goName := range typeReplacements {
		re := regexp.MustCompile(`\b` + regexp.QuoteMeta(cName) + `\b`)
		result = re.ReplaceAllString(result, goName)
	}
	return result
}

func wrapConstRefsForInt(val string) string {
	result := val
	constRefs := []string{
		"MaximumPacketsCapacityPriority",
		"MaxSerialPacketSize",
		"PacketChunkSize",
		"MaximumPacketsCapacity",
		"PageSize64",
		"PageSize",
	}
	for _, ref := range constRefs {
		re := regexp.MustCompile(`\b` + regexp.QuoteMeta(ref) + `\b`)
		result = re.ReplaceAllString(result, "int("+ref+")")
	}
	return result
}

func determineMacroGoType(val string) string {
	cleaned := cleanCMacroValue(val)
	if strings.HasPrefix(cleaned, "0x") || strings.HasPrefix(cleaned, "0X") {
		if len(cleaned) > 10 {
			return "uint64"
		}
		return "uint32"
	}
	if strings.HasPrefix(cleaned, "-") || strings.Contains(cleaned, "(-") {
		numRe := regexp.MustCompile(`\b(\d{10,})\b`)
		if numRe.MatchString(cleaned) {
			return "int64"
		}
		if strings.Contains(cleaned, "<<") || strings.Contains(cleaned, ">>") {
			return "int64"
		}
		return "int32"
	}
	numRe := regexp.MustCompile(`\b(\d{10,})\b`)
	if numRe.MatchString(cleaned) {
		return "uint64"
	}
	shiftRe := regexp.MustCompile(`<<\s*(\d+)`)
	if matches := shiftRe.FindAllStringSubmatch(cleaned, -1); len(matches) > 0 {
		for _, m := range matches {
			if len(m) >= 2 {
				shiftVal := 0
				fmt.Sscanf(m[1], "%d", &shiftVal)
				if shiftVal >= 32 {
					return "uint64"
				}
			}
		}
	}
	return "uint32"
}

func propagateConstTypes(items []macroConstInfo) {
	typeMap := make(map[string]string)
	for _, mc := range items {
		typeMap[mc.goName] = mc.goType
	}
	changed := true
	for changed {
		changed = false
		for i, mc := range items {
			if mc.goType == "uint64" || mc.goType == "int64" {
				for refName, refType := range typeMap {
					if refType == "uint64" || refType == "int64" {
						continue
					}
					re := regexp.MustCompile(`\b` + regexp.QuoteMeta(refName) + `\b`)
					if re.MatchString(mc.goValue) {
						newType := "uint64"
						if mc.goType == "int64" {
							newType = "int64"
						}
						if typeMap[refName] != newType {
							for j := range items {
								if items[j].goName == refName {
									items[j].goType = newType
									break
								}
							}
							typeMap[refName] = newType
							changed = true
						}
					}
				}
			}
			if mc.goType != "uint64" && mc.goType != "int64" {
				for refName, refType := range typeMap {
					if refType != "uint64" && refType != "int64" {
						continue
					}
					re := regexp.MustCompile(`\b` + regexp.QuoteMeta(refName) + `\b`)
					if re.MatchString(mc.goValue) {
						newType := "uint64"
						if refType == "int64" {
							newType = "int64"
						}
						if items[i].goType != newType {
							items[i].goType = newType
							typeMap[mc.goName] = newType
							changed = true
						}
					}
				}
			}
		}
	}
}

func simplifyFnMacro(fn macroConstInfo, varItems []macroConstInfo) (string, string, []macroConstInfo) {
	type callInfo struct {
		index int
		args  []string
	}
	var calls []callInfo
	for i, v := range varItems {
		if !strings.HasPrefix(v.goValue, fn.goName+"(") || !strings.HasSuffix(v.goValue, ")") {
			continue
		}
		argsStr := v.goValue[len(fn.goName)+1 : len(v.goValue)-1]
		args := splitCallArgs(argsStr)
		if len(args) > 0 {
			calls = append(calls, callInfo{index: i, args: args})
		}
	}
	if len(calls) == 0 {
		return "", "", varItems
	}
	params := parseGoParams(fn.goParams)
	if len(params) != len(calls[0].args) {
		return "", "", varItems
	}
	paramCount := len(params)
	isConst := make([]bool, paramCount)
	constVal := make([]string, paramCount)
	for j := range paramCount {
		val := strings.TrimSpace(calls[0].args[j])
		allSame := true
		for _, c := range calls[1:] {
			if j >= len(c.args) || strings.TrimSpace(c.args[j]) != val {
				allSame = false
				break
			}
		}
		if allSame {
			isConst[j] = true
			constVal[j] = val
		}
	}
	hasConst := false
	hasVarying := false
	for j := range paramCount {
		if isConst[j] {
			hasConst = true
		} else {
			hasVarying = true
		}
	}
	if !hasConst || !hasVarying {
		return "", "", varItems
	}
	var newParams []string
	for j := range paramCount {
		if !isConst[j] {
			newParams = append(newParams, params[j].name+" "+params[j].typ)
		}
	}
	newBody := fn.goBody
	for j := range paramCount {
		if isConst[j] {
			re := regexp.MustCompile(`\b` + regexp.QuoteMeta(params[j].name) + `\b`)
			newBody = re.ReplaceAllString(newBody, constVal[j])
		}
	}
	newVarItems := make([]macroConstInfo, len(varItems))
	copy(newVarItems, varItems)
	for _, c := range calls {
		var newArgs []string
		for j := range paramCount {
			if !isConst[j] {
				newArgs = append(newArgs, strings.TrimSpace(c.args[j]))
			}
		}
		newVarItems[c.index].goValue = fn.goName + "(" + strings.Join(newArgs, ", ") + ")"
	}
	return strings.Join(newParams, ", "), newBody, newVarItems
}

func parseGoParams(params string) []struct{ name, typ string } {
	if params == "" {
		return nil
	}
	parts := strings.Split(params, ",")
	var result []struct{ name, typ string }
	for _, p := range parts {
		p = strings.TrimSpace(p)
		spaceIdx := strings.LastIndex(p, " ")
		if spaceIdx < 0 {
			result = append(result, struct{ name, typ string }{name: p, typ: ""})
		} else {
			result = append(result, struct{ name, typ string }{name: p[:spaceIdx], typ: p[spaceIdx+1:]})
		}
	}
	return result
}

func splitCallArgs(argsStr string) []string {
	var result []string
	depth := 0
	start := 0
	for i, ch := range argsStr {
		if ch == '(' {
			depth++
		} else if ch == ')' {
			depth--
		} else if ch == ',' && depth == 0 {
			result = append(result, strings.TrimSpace(argsStr[start:i]))
			start = i + 1
		}
	}
	if start < len(argsStr) {
		result = append(result, strings.TrimSpace(argsStr[start:]))
	}
	return result
}

func generateErrorCodeEnum(content *strings.Builder, items []macroConstInfo, baseName string) {
	content.WriteString(fmt.Sprintf("// Source: %s -> Error codes\n", baseName+".h"))
	content.WriteString("type DebuggerErrorCode uint32\n\n")
	content.WriteString("const (\n")
	for _, mc := range items {
		content.WriteString(fmt.Sprintf("\t%s DebuggerErrorCode = %s\n", mc.goName, mc.goValue))
	}
	content.WriteString(")\n\n")

	content.WriteString("func (c DebuggerErrorCode) String() string {\n")
	content.WriteString("\tswitch c {\n")
	for _, mc := range items {
		humanName := cMacroNameToHumanString(mc.goName)
		content.WriteString(fmt.Sprintf("\tcase %s:\n\t\treturn %q\n", mc.goName, humanName))
	}
	content.WriteString("\tdefault:\n")
	content.WriteString("\t\treturn fmt.Sprintf(\"DebuggerErrorCode(0x%X)\", uint32(c))\n")
	content.WriteString("\t}\n")
	content.WriteString("}\n\n")
}

func cMacroNameToHumanString(goName string) string {
	s := goName
	if after, ok := strings.CutPrefix(s, "DebuggerError"); ok {
		s = after
	} else if after, ok := strings.CutPrefix(s, "DebuggerOperation"); ok {
		s = after
	}
	var result strings.Builder
	for i, ch := range s {
		if i > 0 {
			prev := rune(s[i-1])
			if ch >= 'A' && ch <= 'Z' {
				if prev >= 'a' && prev <= 'z' {
					result.WriteByte(' ')
				} else if prev >= '0' && prev <= '9' {
					result.WriteByte(' ')
				} else if i+1 < len(s) && s[i+1] >= 'a' && s[i+1] <= 'z' {
					result.WriteByte(' ')
				}
			} else if ch >= '0' && ch <= '9' {
				if prev >= 'A' && prev <= 'Z' && (i < 2 || s[i-2] >= '0' && s[i-2] <= '9') {
				} else if prev >= 'a' && prev <= 'z' {
					result.WriteByte(' ')
				}
			}
		}
		result.WriteRune(ch)
	}
	return result.String()
}

func convertCMacroBodyToGo(body string, cParamNames, goParamNames []string) string {
	result := body
	for i, cName := range cParamNames {
		goName := goParamNames[i]
		result = strings.ReplaceAll(result, cName, goName)
	}
	result = strings.ReplaceAll(result, "~", "^")
	result = strings.ReplaceAll(result, "_byteswap_ulong(", "bits.ReverseBytes32(")
	result = strings.ReplaceAll(result, "_byteswap_uint64(", "bits.ReverseBytes64(")
	result = strings.ReplaceAll(result, "_byteswap_ushort(", "bits.ReverseBytes16(")
	result = regexp.MustCompile(`UINT64_C\s*\(`).ReplaceAllString(result, "uint64(")
	result = regexp.MustCompile(`INT64_C\s*\(`).ReplaceAllString(result, "int64(")
	result = regexp.MustCompile(`__has_include\s*\([^)]*\)`).ReplaceAllString(result, "0")
	notRe := regexp.MustCompile(`!\s*\(`)
	result = notRe.ReplaceAllString(result, "!(")
	result = strings.TrimSpace(result)
	depth := 0
	for _, ch := range result {
		if ch == '(' {
			depth++
		}
	}
	if depth > 2 {
		if strings.HasPrefix(result, "(") && strings.HasSuffix(result, ")") {
			inner := result[1 : len(result)-1]
			balance := 0
			valid := true
			for _, ch := range inner {
				if ch == '(' {
					balance++
				} else if ch == ')' {
					balance--
					if balance < 0 {
						valid = false
						break
					}
				}
			}
			if valid && balance == 0 {
				result = inner
			}
		}
	}
	return result
}

func cTagToGoName(name string) string {
	if strings.HasPrefix(name, "_") {
		name = name[1:]
	}
	result := strings.ToUpper(name[:1]) + name[1:]
	goReserved := map[string]bool{
		"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true,
		"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
		"uintptr": true, "string": true, "bool": true, "byte": true, "rune": true,
		"float32": true, "float64": true, "complex64": true, "complex128": true,
		"error": true, "any": true, "comparable": true,
	}
	if goReserved[result] {
		result = result + "_"
	}
	return result
}

func cEnumMemberToGoName(name string) string {
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

func cFieldNameToGoName(name string) string {
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

func cMacroNameToGoName(name string) string {
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

func extractEnumeratorValue(e *cc.Enumerator) (valueStr string, intVal int64, hasVal bool) {
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

func findSourceFileName(unit *cc.TranslationUnit, sources []cc.Source, headerMap map[string]string) string {
	posStr := unit.Position().String()
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

func generateUnionFields(u *cc.UnionType, unionSize int64) string {
	if u.NumFields() == 0 {
		return ""
	}
	var fields []string
	for i := 0; i < u.NumFields(); i++ {
		f := u.FieldByIndex(i)
		ft := f.Type()
		fieldType := mapCTypeToGo(ft)
		fieldName := cFieldNameToGoName(f.Name())
		if fieldName == "_" || fieldType == "" {
			return ""
		}
		if sv, ok := ft.(*cc.StructType); ok {
			if hasBitfields(sv) {
				_, backingType, _ := extractUnionBitfields(u)
				if backingType != "" {
					fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, backingType))
				} else {
					fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
				}
			} else {
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
			}
			continue
		}
		if _, ok := ft.(*cc.UnionType); ok {
			return ""
		}
		if _, ok := ft.(*cc.ArrayType); ok {
			return ""
		}
		fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
	}
	if len(fields) == 0 {
		return ""
	}
	return strings.Join(fields, "\n") + "\n"
}

func unionAlignedFFIType(u *cc.UnionType, unionSize int64) string {
	maxAlign := 1
	maxSize := int64(0)
	var maxType string
	for i := 0; i < u.NumFields(); i++ {
		f := u.FieldByIndex(i)
		ft := f.Type()
		align := ft.Align()
		goType := mapCTypeToGo(ft)
		if goType == "" {
			continue
		}
		if ft.Size() > maxSize || (ft.Size() == maxSize && align > maxAlign) {
			maxSize = ft.Size()
			maxAlign = align
			maxType = goType
		}
	}
	if maxType != "" && maxSize >= unionSize {
		return "Data " + maxType
	}
	return fmt.Sprintf("Data [%d]byte", unionSize)
}

func unionAlignedFFITypeWithSize(types []cc.Type, unionSize int64) string {
	maxAlign := 1
	maxSize := int64(0)
	var maxType string
	for _, t := range types {
		align := t.Align()
		goType := mapCTypeToGo(t)
		if goType == "" {
			continue
		}
		if t.Size() > maxSize || (t.Size() == maxSize && align > maxAlign) {
			maxSize = t.Size()
			maxAlign = align
			maxType = goType
		}
	}
	if maxType != "" && maxSize >= unionSize {
		return "Data " + maxType
	}
	return fmt.Sprintf("Data [%d]byte", unionSize)
}

func generateStructFields(t *cc.StructType, structGoName string, forcePacked bool) (fieldStr string, methodStr string, innerTypes []structInfo) {
	if isPureBitfieldStruct(t) {
		f, m := generatePureBitfieldStruct(t, structGoName)
		return f, m, nil
	}

	isPacked := forcePacked || t.Align() == 1

	var fields []string
	var methods []string
	anonCount := 0
	currentOffset := int64(0)
	var padding int64

	type bitfieldGroup struct {
		offset      int64
		size        int64
		backingName string
		backingType string
		bitfields   []bitfieldInfo
	}

	type anonUnionField struct {
		name   string
		goType string
		ccType cc.Type
		offset int64
		size   int64
		align  int
	}

	var currentBfGroup *bitfieldGroup
	var prevFieldOffset int64 = -1
	var anonUnionCandidates []anonUnionField

	isAnonUnionCandidate := func(ft cc.Type) bool {
		switch v := ft.(type) {
		case *cc.StructType:
			if v.NumFields() == 0 || ft.Size() == 0 {
				return false
			}
			tag := v.Tag()
			tagStr := string(tag.Src())
			if tagStr == "" {
				return false
			}
			goName := cTagToGoName(tagStr)
			if goName == "" || goName == "struct" || strings.Contains(goName, "{") {
				return false
			}
			return true
		case *cc.UnionType:
			if v.NumFields() == 0 || ft.Size() == 0 {
				return false
			}
			tag := v.Tag()
			tagStr := string(tag.Src())
			if tagStr == "" {
				return false
			}
			goName := cTagToGoName(tagStr)
			if goName == "" || goName == "struct" || strings.Contains(goName, "{") {
				return false
			}
			return true
		}
		return false
	}

	flushAnonUnion := func() bool {
		if len(anonUnionCandidates) < 2 {
			return false
		}
		firstOffset := anonUnionCandidates[0].offset
		allSameOffset := true
		for _, af := range anonUnionCandidates {
			if af.offset != firstOffset {
				allSameOffset = false
				break
			}
		}
		if !allSameOffset {
			return false
		}
		maxSize := int64(0)
		for _, af := range anonUnionCandidates {
			if af.size > maxSize {
				maxSize = af.size
			}
		}
		anonCount++
		unionName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
		var unionFieldStrs []string
		for _, af := range anonUnionCandidates {
			unionFieldStrs = append(unionFieldStrs, fmt.Sprintf("\t%s %s", af.name, af.goType))
			switch v := af.ccType.(type) {
			case *cc.StructType:
				tag := v.Tag()
				tagStr := string(tag.Src())
				nestedFields, nestedMethods, _ := generateStructFields(v, af.goType, isPacked)
				innerTypes = append(innerTypes, structInfo{
					goName:  af.goType,
					cName:   tagStr,
					fields:  nestedFields,
					methods: nestedMethods,
				})
			case *cc.UnionType:
				tag := v.Tag()
				tagStr := string(tag.Src())
				if tagStr == "" {
					innerTypes = append(innerTypes, structInfo{
						goName: af.goType,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
					})
				} else if unionFields := generateUnionFields(v, v.Size()); unionFields != "" {
					for j := 0; j < v.NumFields(); j++ {
						if uf := v.FieldByIndex(j); uf != nil {
							if usv, ok := uf.Type().(*cc.StructType); ok {
								nestedTag := usv.Tag()
								nestedTagStr := string(nestedTag.Src())
								if nestedTagStr != "" && !hasBitfields(usv) {
									nestedF, nestedM, _ := generateStructFields(usv, mapCTypeToGo(uf.Type()), isPacked)
									innerTypes = append(innerTypes, structInfo{
										goName:  mapCTypeToGo(uf.Type()),
										cName:   nestedTagStr,
										fields:  nestedF,
										methods: nestedM,
									})
								}
							}
						}
					}
					innerTypes = append(innerTypes, structInfo{
						goName: af.goType + "_",
						cName:  tagStr,
						fields: unionFields,
					})
					innerTypes = append(innerTypes, structInfo{
						goName: af.goType,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
					})
				} else {
					innerTypes = append(innerTypes, structInfo{
						goName: af.goType,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
					})
				}
			}
		}
		if !isPacked {
			padding := anonUnionCandidates[0].offset - currentOffset
			if padding > 0 {
				fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
			}
		}
		accessorName := unionName + "_"
		innerTypes = append(innerTypes, structInfo{
			goName: accessorName,
			cName:  unionName,
			fields: strings.Join(unionFieldStrs, "\n") + "\n",
		})
		var candidateTypes []cc.Type
		for _, af := range anonUnionCandidates {
			candidateTypes = append(candidateTypes, af.ccType)
		}
		innerTypes = append(innerTypes, structInfo{
			goName: unionName,
			cName:  unionName,
			fields: fmt.Sprintf("\t%s\n", unionAlignedFFITypeWithSize(candidateTypes, maxSize)),
		})
		fields = append(fields, fmt.Sprintf("\t%s %s", "U", unionName))
		currentOffset = anonUnionCandidates[0].offset + maxSize
		anonUnionCandidates = nil
		prevFieldOffset = -1
		return true
	}

	for i := 0; i < t.NumFields(); i++ {
		f := t.FieldByIndex(i)
		ft := f.Type()
		fieldOffset := f.Offset()

		if !f.IsBitfield() && prevFieldOffset >= 0 && isAnonUnionCandidate(ft) {
			fieldType := mapCTypeToGo(ft)
			fieldName := cFieldNameToGoName(f.Name())
			if fieldName != "_" && fieldType != "" {
				if len(anonUnionCandidates) > 0 && fieldOffset != anonUnionCandidates[0].offset {
					if flushAnonUnion() {
						padding = fieldOffset - currentOffset
					} else {
						for _, af := range anonUnionCandidates {
							fieldPadding := af.offset - currentOffset
							currentOffset = af.offset + af.size
							if !isPacked && fieldPadding > 0 {
								fields = append(fields, fmt.Sprintf("\t_  [%d]byte", fieldPadding))
							}
							switch v := af.ccType.(type) {
							case *cc.StructType:
								tag := v.Tag()
								tagStr := string(tag.Src())
								nestedFields, nestedMethods, nestedInner := generateStructFields(v, af.goType, isPacked)
								innerTypes = append(innerTypes, structInfo{
									goName:  af.goType,
									cName:   tagStr,
									fields:  nestedFields,
									methods: nestedMethods,
								})
								innerTypes = append(innerTypes, nestedInner...)
								fields = append(fields, fmt.Sprintf("\t%s %s", af.name, af.goType))
							case *cc.UnionType:
								tag := v.Tag()
								tagStr := string(tag.Src())
								if tagStr == "" {
									innerTypes = append(innerTypes, structInfo{
										goName: af.goType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
									})
								} else if unionFields := generateUnionFields(v, v.Size()); unionFields != "" {
									innerTypes = append(innerTypes, structInfo{
										goName: af.goType + "_",
										cName:  tagStr,
										fields: unionFields,
									})
									innerTypes = append(innerTypes, structInfo{
										goName: af.goType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
									})
								} else {
									innerTypes = append(innerTypes, structInfo{
										goName: af.goType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
									})
								}
								fields = append(fields, fmt.Sprintf("\t%s %s", af.name, af.goType))
							}
						}
					}
					anonUnionCandidates = nil
				}
				anonUnionCandidates = append(anonUnionCandidates, anonUnionField{
					name:   fieldName,
					goType: fieldType,
					ccType: ft,
					offset: fieldOffset,
					size:   ft.Size(),
					align:  ft.Align(),
				})
				prevFieldOffset = fieldOffset
				continue
			}
		}

		if len(anonUnionCandidates) > 0 {
			if flushAnonUnion() {
				padding = fieldOffset - currentOffset
			} else {
				for _, af := range anonUnionCandidates {
					fieldPadding := af.offset - currentOffset
					currentOffset = af.offset + af.size
					if !isPacked && fieldPadding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", fieldPadding))
					}
					switch v := af.ccType.(type) {
					case *cc.StructType:
						tag := v.Tag()
						tagStr := string(tag.Src())
						nestedFields, nestedMethods, nestedInner := generateStructFields(v, af.goType, isPacked)
						innerTypes = append(innerTypes, structInfo{
							goName:  af.goType,
							cName:   tagStr,
							fields:  nestedFields,
							methods: nestedMethods,
						})
						innerTypes = append(innerTypes, nestedInner...)
					case *cc.UnionType:
						tag := v.Tag()
						tagStr := string(tag.Src())
						if tagStr != "" {
							if uf := generateUnionFields(v, v.Size()); uf != "" {
								innerTypes = append(innerTypes, structInfo{
									goName: af.goType + "_",
									cName:  tagStr,
									fields: uf,
								})
								innerTypes = append(innerTypes, structInfo{
									goName: af.goType,
									cName:  tagStr,
									fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
								})
							} else {
								innerTypes = append(innerTypes, structInfo{
									goName: af.goType,
									cName:  tagStr,
									fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
								})
							}
						}
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", af.name, af.goType))
				}
			}
			anonUnionCandidates = nil
		}
		prevFieldOffset = fieldOffset

		if f.IsBitfield() {
			bfName := cFieldNameToGoName(f.Name())
			bfBits := f.ValueBits()
			bfOffset := f.OffsetBits()
			bfType := mapCTypeToGo(ft)

			if currentBfGroup == nil || fieldOffset != currentBfGroup.offset {
				if currentBfGroup != nil {
					padding := currentBfGroup.offset - currentOffset
					if !isPacked && padding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", currentBfGroup.backingName, currentBfGroup.backingType))
					for _, bf := range currentBfGroup.bitfields {
						methods = append(methods, generateBitfieldGetterSetter(bf, structGoName, currentBfGroup.backingName, currentBfGroup.backingType))
					}
					currentOffset = currentBfGroup.offset + currentBfGroup.size
				}
				totalBits := ft.Size() * 8
				backingType := "uint8"
				switch totalBits {
				case 8:
					backingType = "uint8"
				case 16:
					backingType = "uint16"
				case 32:
					backingType = "uint32"
				case 64:
					backingType = "uint64"
				}
				currentBfGroup = &bitfieldGroup{
					offset:      fieldOffset,
					size:        ft.Size(),
					backingName: bfName + "Bits",
					backingType: backingType,
				}
			}
			currentBfGroup.bitfields = append(currentBfGroup.bitfields, bitfieldInfo{
				goName:    bfName,
				goType:    bfType,
				bitOffset: bfOffset,
				bitWidth:  bfBits,
			})
			continue
		}

		if currentBfGroup != nil {
			padding := currentBfGroup.offset - currentOffset
			if !isPacked && padding > 0 {
				fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
			}
			fields = append(fields, fmt.Sprintf("\t%s %s", currentBfGroup.backingName, currentBfGroup.backingType))
			for _, bf := range currentBfGroup.bitfields {
				methods = append(methods, generateBitfieldGetterSetter(bf, structGoName, currentBfGroup.backingName, currentBfGroup.backingType))
			}
			currentOffset = currentBfGroup.offset + currentBfGroup.size
			currentBfGroup = nil
		}

		padding = fieldOffset - currentOffset

		fieldType := mapCTypeToGo(ft)
		fieldName := cFieldNameToGoName(f.Name())

		if fieldName == "_" && f.Name() == "" {
			switch v := ft.(type) {
			case *cc.UnionType:
				backingName, backingType, bfs := extractUnionBitfields(v)
				if backingName != "" && len(bfs) > 0 {
					if !isPacked && padding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", backingName, backingType))
					for _, bf := range bfs {
						methods = append(methods, generateBitfieldGetterSetter(bf, structGoName, backingName, backingType))
					}
					currentOffset = fieldOffset + ft.Size()
					continue
				}
				if unionFields := generateUnionFields(v, v.Size()); unionFields != "" {
					anonCount++
					anonFieldName := fmt.Sprintf("Anon%d", anonCount)
					anonUnionName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
					innerTypes = append(innerTypes, structInfo{
						goName: anonUnionName + "_",
						cName:  anonUnionName,
						fields: unionFields,
					})
					innerTypes = append(innerTypes, structInfo{
						goName: anonUnionName,
						cName:  anonUnionName,
						fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
					})
					if !isPacked && padding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", anonFieldName, anonUnionName))
					currentOffset = fieldOffset + ft.Size()
					continue
				}
				if v.NumFields() > 0 {
					anonCount++
					anonFieldName := fmt.Sprintf("Anon%d", anonCount)
					anonUnionName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
					var unionFieldStrs []string
					for i := 0; i < v.NumFields(); i++ {
						f := v.FieldByIndex(i)
						ft := f.Type()
						fieldType := mapCTypeToGo(ft)
						fieldName := cFieldNameToGoName(f.Name())
						if fieldName == "_" || fieldType == "" {
							continue
						}
						unionFieldStrs = append(unionFieldStrs, fmt.Sprintf("\t%s %s", fieldName, fieldType))
						switch vt := ft.(type) {
						case *cc.StructType:
							tag := vt.Tag()
							tagStr := string(tag.Src())
							nestedFields, nestedMethods, nestedInner := generateStructFields(vt, fieldType, isPacked)
							innerTypes = append(innerTypes, structInfo{
								goName:  fieldType,
								cName:   tagStr,
								fields:  nestedFields,
								methods: nestedMethods,
							})
							innerTypes = append(innerTypes, nestedInner...)
						case *cc.UnionType:
							tag := vt.Tag()
							tagStr := string(tag.Src())
							if tagStr != "" {
								if uf := generateUnionFields(vt, vt.Size()); uf != "" {
									innerTypes = append(innerTypes, structInfo{
										goName: fieldType + "_",
										cName:  tagStr,
										fields: uf,
									})
									innerTypes = append(innerTypes, structInfo{
										goName: fieldType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(vt, vt.Size())),
									})
								} else {
									innerTypes = append(innerTypes, structInfo{
										goName: fieldType,
										cName:  tagStr,
										fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(vt, vt.Size())),
									})
								}
							}
						}
					}
					if len(unionFieldStrs) > 0 {
						accessorName := anonUnionName + "_"
						innerTypes = append(innerTypes, structInfo{
							goName: accessorName,
							cName:  anonUnionName,
							fields: strings.Join(unionFieldStrs, "\n") + "\n",
						})
						innerTypes = append(innerTypes, structInfo{
							goName: anonUnionName,
							cName:  anonUnionName,
							fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
						})
						if !isPacked && padding > 0 {
							fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
						}
						fields = append(fields, fmt.Sprintf("\t%s %s", anonFieldName, anonUnionName))
						currentOffset = fieldOffset + v.Size()
						continue
					}
				}
			case *cc.StructType:
				if hasBitfields(v) {
					anonCount++
					anonFieldName := fmt.Sprintf("Anon%d", anonCount)
					bitfieldGoName := fieldType
					bitfieldFields, bitfieldMethods := generatePureBitfieldStruct(v, bitfieldGoName)
					innerTypes = append(innerTypes, structInfo{
						goName:  bitfieldGoName,
						cName:   bitfieldGoName,
						fields:  bitfieldFields,
						methods: bitfieldMethods,
					})
					if !isPacked && padding > 0 {
						fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
					}
					fields = append(fields, fmt.Sprintf("\t%s %s", anonFieldName, bitfieldGoName))
					currentOffset = fieldOffset + ft.Size()
					continue
				}
			}
		}

		if !isPacked && padding > 0 {
			fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
		}

		if v, ok := ft.(*cc.StructType); ok {
			tag := v.Tag()
			tagStr := string(tag.Src())
			if tagStr != "" && hasBitfields(v) {
				bitfieldGoName := fieldType
				bitfieldFields, bitfieldMethods := generatePureBitfieldStruct(v, bitfieldGoName)
				innerTypes = append(innerTypes, structInfo{
					goName:  bitfieldGoName,
					cName:   bitfieldGoName,
					fields:  bitfieldFields,
					methods: bitfieldMethods,
				})
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, bitfieldGoName))
			} else if tagStr != "" {
				nestedFields, nestedMethods, nestedInner := generateStructFields(v, fieldType, isPacked)
				innerTypes = append(innerTypes, structInfo{
					goName:  fieldType,
					cName:   fieldType,
					fields:  nestedFields,
					methods: nestedMethods,
				})
				innerTypes = append(innerTypes, nestedInner...)
				for _, it := range generateStructFieldsInnerTypes(v, fieldType, isPacked) {
					innerTypes = append(innerTypes, it)
				}
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
			} else if hasBitfields(v) {
				bitfieldGoName := fieldType
				bitfieldFields, bitfieldMethods := generatePureBitfieldStruct(v, bitfieldGoName)
				innerTypes = append(innerTypes, structInfo{
					goName:  bitfieldGoName,
					cName:   bitfieldGoName,
					fields:  bitfieldFields,
					methods: bitfieldMethods,
				})
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, bitfieldGoName))
			} else {
				anonCount++
				anonGoName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount)
				nestedFields, nestedMethods, nestedInner := generateStructFields(v, anonGoName, isPacked)
				innerTypes = append(innerTypes, structInfo{
					goName:  anonGoName,
					cName:   anonGoName,
					fields:  nestedFields,
					methods: nestedMethods,
				})
				innerTypes = append(innerTypes, nestedInner...)
				for _, it := range generateStructFieldsInnerTypes(v, anonGoName, isPacked) {
					innerTypes = append(innerTypes, it)
				}
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, anonGoName))
			}
		} else if uv, ok := ft.(*cc.UnionType); ok {
			tag := uv.Tag()
			tagStr := string(tag.Src())
			unionSize := uv.Size()
			if fieldType != "uint64" {
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
			} else if tagStr == "" {
				anonCount++
				unionGoName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
				if unionFields := generateUnionFields(uv, unionSize); unionFields != "" {
					for j := 0; j < uv.NumFields(); j++ {
						if uf := uv.FieldByIndex(j); uf != nil {
							if usv, ok := uf.Type().(*cc.StructType); ok {
								nestedTag := usv.Tag()
								nestedTagStr := string(nestedTag.Src())
								if nestedTagStr != "" && !hasBitfields(usv) {
									nestedF, nestedM, _ := generateStructFields(usv, mapCTypeToGo(uf.Type()), isPacked)
									innerTypes = append(innerTypes, structInfo{
										goName:  mapCTypeToGo(uf.Type()),
										cName:   nestedTagStr,
										fields:  nestedF,
										methods: nestedM,
									})
								}
							}
						}
					}
					innerTypes = append(innerTypes, structInfo{
						goName: unionGoName + "_",
						cName:  tagStr,
						fields: unionFields,
					})
					innerTypes = append(innerTypes, structInfo{
						goName: unionGoName,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(uv, unionSize)),
					})
				} else {
					innerTypes = append(innerTypes, structInfo{
						goName: unionGoName,
						cName:  unionGoName,
						fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(uv, unionSize)),
					})
				}
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, unionGoName))
			} else if unionFields := generateUnionFields(uv, unionSize); unionFields != "" {
				unionGoName := fieldType
				for j := 0; j < uv.NumFields(); j++ {
					if uf := uv.FieldByIndex(j); uf != nil {
						if usv, ok := uf.Type().(*cc.StructType); ok {
							nestedTag := usv.Tag()
							nestedTagStr := string(nestedTag.Src())
							if nestedTagStr != "" && !hasBitfields(usv) {
								nestedF, nestedM, _ := generateStructFields(usv, mapCTypeToGo(uf.Type()), isPacked)
								innerTypes = append(innerTypes, structInfo{
									goName:  mapCTypeToGo(uf.Type()),
									cName:   nestedTagStr,
									fields:  nestedF,
									methods: nestedM,
								})
							}
						}
					}
				}
				innerTypes = append(innerTypes, structInfo{
					goName: unionGoName + "_",
					cName:  tagStr,
					fields: unionFields,
				})
				innerTypes = append(innerTypes, structInfo{
					goName: unionGoName,
					cName:  tagStr,
					fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(uv, unionSize)),
				})
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, unionGoName))
			} else {
				unionGoName := fieldType
				innerTypes = append(innerTypes, structInfo{
					goName: unionGoName,
					cName:  tagStr,
					fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(uv, unionSize)),
				})
				fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, unionGoName))
			}
		} else if arr, ok := ft.(*cc.ArrayType); ok {
			elem := arr.Elem()
			for innerArr, ok := elem.(*cc.ArrayType); ok; innerArr, ok = elem.(*cc.ArrayType) {
				elem = innerArr.Elem()
			}
			if sv, ok := elem.(*cc.StructType); ok {
				tag := sv.Tag()
				tagStr := string(tag.Src())
				if tagStr != "" {
					elemGoName := cTagToGoName(tagStr)
					nestedFields, nestedMethods, nestedInner := generateStructFields(sv, elemGoName, isPacked)
					innerTypes = append(innerTypes, structInfo{
						goName:  elemGoName,
						cName:   tagStr,
						fields:  nestedFields,
						methods: nestedMethods,
					})
					innerTypes = append(innerTypes, nestedInner...)
					fieldType = elemGoName
				} else {
					anonCount++
					anonElemName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Elem"
					nestedFields, nestedMethods, nestedInner := generateStructFields(sv, anonElemName, isPacked)
					innerTypes = append(innerTypes, structInfo{
						goName:  anonElemName,
						cName:   anonElemName,
						fields:  nestedFields,
						methods: nestedMethods,
					})
					innerTypes = append(innerTypes, nestedInner...)
					fieldType = anonElemName
				}
			} else if uv, ok := elem.(*cc.UnionType); ok {
				tag := uv.Tag()
				tagStr := string(tag.Src())
				unionSize := uv.Size()
				if unionFields := generateUnionFields(uv, unionSize); unionFields != "" {
					elemGoName := cTagToGoName(tagStr)
					if elemGoName == "" {
						elemGoName = fieldType
					}
					innerTypes = append(innerTypes, structInfo{
						goName: elemGoName,
						cName:  tagStr,
						fields: unionFields,
					})
					fieldType = elemGoName
				} else if tagStr != "" {
					elemGoName := cTagToGoName(tagStr)
					innerTypes = append(innerTypes, structInfo{
						goName: elemGoName,
						cName:  tagStr,
						fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(uv, unionSize)),
					})
					fieldType = elemGoName
				} else {
					anonCount++
					anonUnionName := structGoName + "_Anon" + fmt.Sprintf("%d", anonCount) + "Union"
					innerTypes = append(innerTypes, structInfo{
						goName: anonUnionName,
						cName:  anonUnionName,
						fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(uv, unionSize)),
					})
					fieldType = anonUnionName
				}
			} else {
				fieldType = mapCTypeToGo(arr.Elem())
			}
			fields = append(fields, fmt.Sprintf("\t%s [%d]%s", fieldName, arr.Len(), fieldType))
		} else {
			fields = append(fields, fmt.Sprintf("\t%s %s", fieldName, fieldType))
		}
		currentOffset = fieldOffset + ft.Size()
	}

	if currentBfGroup != nil {
		padding := currentBfGroup.offset - currentOffset
		if !isPacked && padding > 0 {
			fields = append(fields, fmt.Sprintf("\t_  [%d]byte", padding))
		}
		fields = append(fields, fmt.Sprintf("\t%s %s", currentBfGroup.backingName, currentBfGroup.backingType))
		for _, bf := range currentBfGroup.bitfields {
			methods = append(methods, generateBitfieldGetterSetter(bf, structGoName, currentBfGroup.backingName, currentBfGroup.backingType))
		}
		currentOffset = currentBfGroup.offset + currentBfGroup.size
	}

	if len(anonUnionCandidates) > 0 {
		if !flushAnonUnion() {
			for _, af := range anonUnionCandidates {
				fieldPadding := af.offset - currentOffset
				currentOffset = af.offset + af.size
				if !isPacked && fieldPadding > 0 {
					fields = append(fields, fmt.Sprintf("\t_  [%d]byte", fieldPadding))
				}
				switch v := af.ccType.(type) {
				case *cc.StructType:
					tag := v.Tag()
					tagStr := string(tag.Src())
					nestedFields, nestedMethods, nestedInner := generateStructFields(v, af.goType, isPacked)
					innerTypes = append(innerTypes, structInfo{
						goName:  af.goType,
						cName:   tagStr,
						fields:  nestedFields,
						methods: nestedMethods,
					})
					innerTypes = append(innerTypes, nestedInner...)
				case *cc.UnionType:
					tag := v.Tag()
					tagStr := string(tag.Src())
					if tagStr != "" {
						if uf := generateUnionFields(v, v.Size()); uf != "" {
							innerTypes = append(innerTypes, structInfo{
								goName: af.goType + "_",
								cName:  tagStr,
								fields: uf,
							})
							innerTypes = append(innerTypes, structInfo{
								goName: af.goType,
								cName:  tagStr,
								fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
							})
						} else {
							innerTypes = append(innerTypes, structInfo{
								goName: af.goType,
								cName:  tagStr,
								fields: fmt.Sprintf("\t%s\n", unionAlignedFFIType(v, v.Size())),
							})
						}
					}
				}
				fields = append(fields, fmt.Sprintf("\t%s %s", af.name, af.goType))
			}
		}
	}

	totalSize := t.Size()
	trailingPadding := totalSize - currentOffset
	if !isPacked && trailingPadding > 0 {
		fields = append(fields, fmt.Sprintf("\t_  [%d]byte", trailingPadding))
	}

	fieldNameCount := make(map[string]int)
	for i, f := range fields {
		parts := strings.SplitN(strings.TrimSpace(f), " ", 2)
		if len(parts) < 2 {
			continue
		}
		name := parts[0]
		if name == "_" {
			continue
		}
		fieldNameCount[name]++
		if fieldNameCount[name] > 1 {
			newName := fmt.Sprintf("%s%d", name, fieldNameCount[name])
			fields[i] = strings.Replace(f, name, newName, 1)
		}
	}

	return strings.Join(fields, "\n") + "\n", strings.Join(methods, "\n"), innerTypes
}

func generateStructFieldsInnerTypes(t *cc.StructType, structGoName string, forcePacked bool) []structInfo {
	var result []structInfo
	for i := 0; i < t.NumFields(); i++ {
		f := t.FieldByIndex(i)
		ft := f.Type()
		if v, ok := ft.(*cc.StructType); ok {
			tag := v.Tag()
			tagStr := string(tag.Src())
			if tagStr != "" {
				nestedFields, nestedMethods, nestedInner := generateStructFields(v, tagStr, forcePacked)
				result = append(result, structInfo{
					goName:  tagStr,
					cName:   tagStr,
					fields:  nestedFields,
					methods: nestedMethods,
				})
				result = append(result, nestedInner...)
				result = append(result, generateStructFieldsInnerTypes(v, tagStr, forcePacked)...)
			}
		}
	}
	return result
}

func hasBitfields(t *cc.StructType) bool {
	for i := 0; i < t.NumFields(); i++ {
		if t.FieldByIndex(i).IsBitfield() {
			return true
		}
	}
	return false
}

func isPureBitfieldStruct(t *cc.StructType) bool {
	if t.NumFields() == 0 {
		return false
	}
	for i := 0; i < t.NumFields(); i++ {
		if !t.FieldByIndex(i).IsBitfield() {
			return false
		}
	}
	return true
}

func generatePureBitfieldStruct(t *cc.StructType, structGoName string) (string, string) {
	totalBytes := t.Size()

	var backingType string
	switch totalBytes {
	case 1:
		backingType = "uint8"
	case 2:
		backingType = "uint16"
	case 4:
		backingType = "uint32"
	case 8:
		backingType = "uint64"
	default:
		if totalBytes <= 16 {
			backingType = "uint64"
		} else {
			backingType = fmt.Sprintf("[%d]byte", totalBytes)
		}
	}

	bfs := extractStructBitfields(t)
	var methods []string
	for _, bf := range bfs {
		methods = append(methods, generateBitfieldGetterSetter(bf, structGoName, "Value", backingType))
	}

	fieldStr := fmt.Sprintf("\tValue %s\n", backingType)
	return fieldStr, strings.Join(methods, "\n")
}

func extractUnionBitfields(u *cc.UnionType) (backingName string, backingType string, bfs []bitfieldInfo) {
	backingIdx := -1
	bitfieldIdx := -1
	for i := 0; i < u.NumFields(); i++ {
		f := u.FieldByIndex(i)
		ft := f.Type()
		switch v := ft.(type) {
		case *cc.StructType:
			if hasBitfields(v) {
				bitfieldIdx = i
			}
		default:
			if !f.IsBitfield() && backingIdx == -1 {
				backingIdx = i
			}
		}
	}
	if backingIdx == -1 || bitfieldIdx == -1 {
		return "", "", nil
	}
	bf := u.FieldByIndex(backingIdx)
	backingName = cFieldNameToGoName(bf.Name())
	backingType = mapCTypeToGo(bf.Type())
	bs := u.FieldByIndex(bitfieldIdx).Type().(*cc.StructType)
	bfs = extractStructBitfields(bs)
	return backingName, backingType, bfs
}

func extractStructBitfields(t *cc.StructType) []bitfieldInfo {
	var bfs []bitfieldInfo
	bitPos := 0
	for i := 0; i < t.NumFields(); i++ {
		f := t.FieldByIndex(i)
		if !f.IsBitfield() {
			continue
		}
		bfs = append(bfs, bitfieldInfo{
			goName:    cFieldNameToGoName(f.Name()),
			goType:    mapCTypeToGo(f.Type()),
			bitOffset: bitPos,
			bitWidth:  f.ValueBits(),
		})
		bitPos += int(f.ValueBits())
	}
	return bfs
}

func generateBitfieldGetterSetter(bf bitfieldInfo, structName, backingName, backingType string) string {
	mask := uint64(1)<<bf.bitWidth - 1
	shift := bf.bitOffset
	goName := bf.goName
	receiver := strings.ToLower(structName[:1])

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("\nfunc (%s *%s) Get%s() %s {\n", receiver, structName, goName, bf.goType))
	if shift == 0 {
		sb.WriteString(fmt.Sprintf("\treturn %s(%s.%s & %s(0x%X))\n", bf.goType, receiver, backingName, backingType, mask))
	} else {
		sb.WriteString(fmt.Sprintf("\treturn %s((%s.%s >> %d) & %s(0x%X))\n", bf.goType, receiver, backingName, shift, backingType, mask))
	}
	sb.WriteString("}\n")

	sb.WriteString(fmt.Sprintf("\nfunc (%s *%s) Set%s(val %s) {\n", receiver, structName, goName, bf.goType))
	if shift == 0 {
		sb.WriteString(fmt.Sprintf("\t%s.%s = (%s.%s & ^%s(0x%X)) | %s(%s(val) & 0x%X)\n", receiver, backingName, receiver, backingName, backingType, mask, backingType, bf.goType, mask))
	} else {
		sb.WriteString(fmt.Sprintf("\t%s.%s = (%s.%s & ^%s(0x%X<<%d)) | (%s(%s(val) & 0x%X) << %d)\n", receiver, backingName, receiver, backingName, backingType, mask, shift, backingType, bf.goType, mask, shift))
	}
	sb.WriteString("}\n")

	return sb.String()
}

func isGoPrimitiveType(goType string) bool {
	switch goType {
	case "bool", "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64",
		"float32", "float64", "int", "uint", "uintptr", "*bool", "*int8", "*uint8",
		"*int16", "*uint16", "*int32", "*uint32", "*int64", "*uint64", "*uintptr":
		return true
	default:
		return false
	}
}

func mapCTypedefToGo(declName string, t cc.Type) string {
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
		return mapPredefinedType(v.Kind(), "")
	case *cc.PointerType:
		return mapCPointerTypedefToGo(v)
	case *cc.StructType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
		}
		return ""
	case *cc.EnumType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
		}
		return "uint32"
	case *cc.UnionType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
		}
		return ""
	case *cc.FunctionType:
		return ""
	default:
		goType := mapCTypeToGo(t)
		if goType != "" && goType != "void" && goType != cTagToGoName(declName) {
			return goType
		}
		return ""
	}
}

func mapPredefinedType(kind cc.Kind, fallback string) string {
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

func mapCPointerTypedefToGo(t *cc.PointerType) string {
	elem := t.Elem()
	if ft, ok := elem.(*cc.FunctionType); ok {
		return mapCFunctionTypeToGo(ft)
	}
	goElem := mapCTypeToGoForPointer(elem)
	return "*" + goElem
}

func mapCFunctionTypeToGo(ft *cc.FunctionType) string {
	goRet := mapCTypeToGo(ft.Result())
	params := ft.Parameters()
	var goParams []string
	for _, p := range params {
		goParamType := mapCTypeToGo(p.Type())
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

func mapCTypeToGoForPointer(t cc.Type) string {
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
		}
		return cTagToGoName(tdName)
	}

	switch v := t.(type) {
	case *cc.PredefinedType:
		return mapPredefinedType(v.Kind(), "byte")
	case *cc.PointerType:
		return mapCPointerTypedefToGo(v)
	case *cc.ArrayType:
		elemType := mapCTypeToGoForPointer(v.Elem())
		size := v.Len()
		return fmt.Sprintf("[%d]%s", size, elemType)
	case *cc.StructType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
		}
		return "struct{}"
	case *cc.EnumType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
		}
		return "uint32"
	case *cc.UnionType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
		}
		return "struct{}"
	default:
		return "byte"
	}
}

func mapCTypeToGo(t cc.Type) string {
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
		case "PWCHAR", "LPWSTR":
			return "*uint16"
		case "LPSTR", "LPCSTR", "LPCWSTR":
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
			return "intptr"
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
		}
		return cTagToGoName(tdName)
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
		goElem := mapCTypeToGo(elem)
		if goElem == "void" {
			return "unsafe.Pointer"
		}
		return "*" + goElem
	case *cc.ArrayType:
		elemType := mapCTypeToGo(v.Elem())
		size := v.Len()
		result := fmt.Sprintf("[%d]%s", size, elemType)
		return result
	case *cc.StructType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
		}
		return "struct{}"
	case *cc.EnumType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
		}
		return "uint32"
	case *cc.UnionType:
		tag := v.Tag()
		tagStr := string(tag.Src())
		if tagStr != "" {
			return cTagToGoName(tagStr)
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

type dllFuncInfo struct {
	goName     string
	cName      string
	returnType string
	params     []dllFuncParam
}

type dllFuncParam struct {
	goName    string
	goType    string
	isPointer bool
}

func generateDllBinding(t *testing.T, ast *cc.AST, bc BindgenConfig, typedefs *safemap.M[string, typedefInfo], enums *safemap.M[string, enumInfo]) {
	var funcs []dllFuncInfo
	funcTypeNames := map[string]bool{}
	for l := ast.TranslationUnit; l != nil; l = l.TranslationUnit {
		ed := l.ExternalDeclaration
		if ed == nil || ed.Case != cc.ExternalDeclarationDecl {
			continue
		}
		decl := ed.Declaration
		if decl == nil || decl.DeclarationSpecifiers == nil || decl.InitDeclaratorList == nil {
			continue
		}
		for id := decl.InitDeclaratorList; id != nil; id = id.InitDeclaratorList {
			if id.InitDeclarator == nil || id.InitDeclarator.Declarator == nil {
				continue
			}
			dtor := id.InitDeclarator.Declarator
			name := dtor.Name()
			if name == "" {
				continue
			}
			if bc.DllFuncFilter != nil && !bc.DllFuncFilter(name) {
				continue
			}
			dtorType := dtor.Type()
			if dtorType == nil {
				continue
			}
			ft, ok := dtorType.(*cc.FunctionType)
			if !ok {
				if pt, isPtr := dtorType.(*cc.PointerType); isPtr {
					if _, isFn := pt.Elem().(*cc.FunctionType); isFn {
						funcTypeNames[cTagToGoName(name)] = true
					}
				}
				continue
			}
			goRet := mapCTypeToGo(ft.Result())
			if goRet == "" {
				goRet = "uintptr"
			}
			var params []dllFuncParam
			paramIdx := 0
			for _, p := range ft.Parameters() {
				pName := ""
				if p.Declarator != nil {
					pName = p.Declarator.Name()
				}
				pGoType := mapCTypeToGo(p.Type())
				if pGoType == "" {
					pGoType = "uintptr"
				}
				if pGoType == "void" {
					continue
				}
				if pName == "" {
					paramIdx++
					pName = fmt.Sprintf("arg%d", paramIdx)
				} else {
					paramIdx++
				}
				_, isPtr := p.Type().(*cc.PointerType)
				params = append(params, dllFuncParam{
					goName:    cTagToGoName(pName),
					goType:    pGoType,
					isPointer: isPtr,
				})
			}
			funcs = append(funcs, dllFuncInfo{
				goName:     cTagToGoName(name),
				cName:      name,
				returnType: goRet,
				params:     params,
			})
		}
	}
	fmt.Printf("Found %d exported functions for DLL binding\n", len(funcs))

	dllCode := generateDllBindingCode(funcs, bc.DllName, bc.PackageName, funcTypeNames, typedefs, enums)
	dllFile := filepath.Join(bc.OutputDir, "dll.go")
	os.WriteFile(dllFile, []byte(dllCode), 0o644)
	fmt.Printf("Generated: %s\n", dllFile)

	if bc.DllName != "" {
		dstDll := filepath.Join(bc.OutputDir, bc.DllName)
		copied := false
		searchDirs := []string{bc.HeadersDir, filepath.Dir(bc.HeadersDir)}
		for _, baseDir := range searchDirs {
			for _, srcDir := range []string{"build", "dll"} {
				srcDll := filepath.Join(baseDir, srcDir, bc.DllName)
				if data, err := os.ReadFile(srcDll); err == nil {
					os.WriteFile(dstDll, data, 0o644)
					fmt.Printf("Copied: %s -> %s\n", srcDll, dstDll)
					copied = true
					break
				}
			}
			if copied {
				break
			}
			entries, err := os.ReadDir(baseDir)
			if err != nil {
				continue
			}
			for _, entry := range entries {
				if !entry.IsDir() {
					continue
				}
				srcDll := filepath.Join(baseDir, entry.Name(), bc.DllName)
				if data, err := os.ReadFile(srcDll); err == nil {
					os.WriteFile(dstDll, data, 0o644)
					fmt.Printf("Copied: %s -> %s\n", srcDll, dstDll)
					copied = true
					break
				}
			}
			if copied {
				break
			}
		}
		if !copied {
			fmt.Printf("Warning: %s not found in build/ or dll/, skip copy\n", bc.DllName)
		}
	}
}

func generateDllBindingCode(funcs []dllFuncInfo, dllName, pkgName string, funcTypeNames map[string]bool, typedefs *safemap.M[string, typedefInfo], enums *safemap.M[string, enumInfo]) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	sb.WriteString("import (\n")
	sb.WriteString("\t_ \"embed\"\n")
	sb.WriteString("\t\"os\"\n")
	sb.WriteString("\t\"path/filepath\"\n")
	hasCallbacks := false
	for _, fn := range funcs {
		if fn.returnType != "void" && fn.returnType != "" && strings.Contains(fn.returnType, "func(") {
			hasCallbacks = true
		}
		for _, p := range fn.params {
			if strings.Contains(p.goType, "func(") {
				hasCallbacks = true
			}
			if funcTypeNames[p.goType] {
			}
		}
		if hasCallbacks {
			break
		}
	}
	sb.WriteString("\t\"unsafe\"\n\n")
	sb.WriteString("\t\"golang.org/x/sys/windows\"\n")
	sb.WriteString(")\n\n")

	receiverType := strings.ToUpper(pkgName[:1]) + pkgName[1:]
	receiver := strings.ToLower(receiverType[:1])
	sb.WriteString(fmt.Sprintf("type %s struct{}\n\n", receiverType))

	sb.WriteString(fmt.Sprintf("//go:embed %s\nvar dllBytes []byte\n\n", dllName))

	sb.WriteString("var (\n")
	sb.WriteString("\tdll            *windows.LazyDLL\n")
	for _, fn := range funcs {
		sb.WriteString(fmt.Sprintf("\tproc_%s      *windows.LazyProc\n", goDllIdent(fn.cName)))
	}
	sb.WriteString(")\n\n")

	sb.WriteString("func init() {\n")
	sb.WriteString(fmt.Sprintf("\tdll = windows.NewLazyDLL(saveEmbeddedDLL(dllBytes, %q))\n", dllName))
	for _, fn := range funcs {
		sb.WriteString(fmt.Sprintf("\tproc_%s = dll.NewProc(%q)\n", goDllIdent(fn.cName), fn.cName))
	}
	sb.WriteString("}\n\n")

	sb.WriteString("func saveEmbeddedDLL(data []byte, name string) string {\n")
	sb.WriteString("\ttmpDir := os.TempDir()\n")
	sb.WriteString("\tp := filepath.Join(tmpDir, name)\n")
	sb.WriteString("\tos.WriteFile(p, data, 0644)\n")
	sb.WriteString("\treturn p\n")
	sb.WriteString("}\n\n")

	for _, fn := range funcs {
		goName := strings.TrimPrefix(fn.goName, receiverType)
		goName = strings.TrimLeft(goName, "_")
		goName = goDllIdent(goName)
		parts := strings.Split(goName, "_")
		for i, p := range parts {
			if len(p) > 0 {
				parts[i] = strings.ToUpper(p[:1]) + p[1:]
			}
		}
		goName = strings.Join(parts, "")
		var goParams []string
		for _, p := range fn.params {
			goParams = append(goParams, fmt.Sprintf("%s %s", p.goName, p.goType))
		}
		paramsStr := strings.Join(goParams, ", ")

		returnSig := ""
		if fn.returnType != "void" && fn.returnType != "" {
			returnSig = " " + fn.returnType
		}
		sb.WriteString(fmt.Sprintf("func (%s *%s) %s(%s)%s {\n", receiver, receiverType, goName, paramsStr, returnSig))

		var syscallArgs []string
		for _, p := range fn.params {
			syscallArgs = append(syscallArgs, dllParamToSyscall(p, funcTypeNames, typedefs, enums))
		}
		argsStr := strings.Join(syscallArgs, ", ")
		if fn.returnType == "void" || fn.returnType == "" {
			sb.WriteString(fmt.Sprintf("\tproc_%s.Call(%s)\n", goDllIdent(fn.cName), argsStr))
		} else if fn.returnType == "bool" {
			sb.WriteString(fmt.Sprintf("\tr1, _, _ := proc_%s.Call(%s)\n", goDllIdent(fn.cName), argsStr))
			sb.WriteString("\treturn r1 != 0\n")
		} else {
			sb.WriteString(fmt.Sprintf("\tr1, _, _ := proc_%s.Call(%s)\n", goDllIdent(fn.cName), argsStr))
			resolvedRet := resolveTypedef(fn.returnType, typedefs)
			cast := dllReturnTypeCast(resolvedRet)
			smallIntTypes := map[string]bool{"uintptr": true, "uint32": true, "int32": true, "uint16": true, "int16": true, "uint8": true, "int8": true}
			if enums.Has(resolvedRet) || enums.Has(fn.returnType) {
				sb.WriteString(fmt.Sprintf("\treturn %s(uint32(r1))\n", fn.returnType))
			} else if cast != "" && smallIntTypes[cast] {
				sb.WriteString(fmt.Sprintf("\treturn %s(r1)\n", cast))
			} else if resolvedRet == "unsafe.Pointer" {
				sb.WriteString("\treturn unsafe.Pointer(r1)\n")
			} else if after, ok := strings.CutPrefix(resolvedRet, "*"); ok {
				sb.WriteString(fmt.Sprintf("\treturn (*%s)(unsafe.Pointer(r1))\n", after))
			} else if fn.returnType != "" && fn.returnType != "void" {
				sb.WriteString(fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&r1))\n", fn.returnType))
			} else {
				sb.WriteString(fmt.Sprintf("\treturn %s(r1)\n", cast))
			}
		}
		sb.WriteString("}\n\n")
	}

	return sb.String()
}

func goDllIdent(name string) string {
	if idx := strings.LastIndex(name, "@"); idx >= 0 {
		return name[:idx]
	}
	return name
}

func dllParamToSyscall(p dllFuncParam, funcTypeNames map[string]bool, typedefs *safemap.M[string, typedefInfo], enums *safemap.M[string, enumInfo]) string {
	resolvedType := p.goType
	seen := make(map[string]bool)
	for {
		if seen[resolvedType] {
			break
		}
		seen[resolvedType] = true
		if td, ok := typedefs.Get(resolvedType); ok && !td.isFunc {
			resolvedType = td.goType
		} else {
			break
		}
	}
	switch resolvedType {
	case "uintptr":
		return p.goName
	case "uint64", "int64":
		return "*(*uintptr)(unsafe.Pointer(&" + p.goName + "))"
	case "uint32", "int32":
		return "uintptr(" + p.goName + ")"
	case "uint16", "int16":
		return "uintptr(" + p.goName + ")"
	case "uint8", "int8":
		return "uintptr(" + p.goName + ")"
	case "bool":
		return "uintptr(0)"
	case "unsafe.Pointer":
		return "uintptr(" + p.goName + ")"
	default:
		if strings.HasPrefix(resolvedType, "func(") || funcTypeNames[p.goType] {
			return "uintptr(*(*unsafe.Pointer)(unsafe.Pointer(&" + p.goName + ")))"
		}
		if strings.HasPrefix(p.goType, "*") || p.isPointer {
			return "uintptr(unsafe.Pointer(" + p.goName + "))"
		}
		if isGoPrimitiveType(resolvedType) {
			return "uintptr(" + p.goName + ")"
		}
		if enums.Has(resolvedType) || enums.Has(p.goType) {
			return "uintptr(" + p.goName + ")"
		}
		return "uintptr(unsafe.Pointer(&" + p.goName + "))"
	}
}

func resolveTypedef(goType string, typedefs *safemap.M[string, typedefInfo]) string {
	seen := make(map[string]bool)
	resolved := goType
	for {
		if seen[resolved] {
			break
		}
		seen[resolved] = true
		if td, ok := typedefs.Get(resolved); ok && !td.isFunc {
			resolved = td.goType
		} else {
			break
		}
	}
	return resolved
}

func dllReturnTypeCast(goType string) string {
	switch goType {
	case "uintptr":
		return "uintptr"
	case "uint64":
		return "uint64"
	case "uint32":
		return "uint32"
	case "int32":
		return "int32"
	case "int16":
		return "int16"
	case "uint16":
		return "uint16"
	case "int8":
		return "int8"
	case "uint8":
		return "uint8"
	default:
		if strings.HasPrefix(goType, "*") {
			return ""
		}
		if goType == "unsafe.Pointer" {
			return ""
		}
		return goType
	}
}

func TestDllReceiverName(t *testing.T) {
	tests := []struct {
		pkgName    string
		wantRecv   string
		wantStruct string
	}{
		{"keystone", "k", "Keystone"},
		{"zydis", "z", "Zydis"},
		{"arimp", "a", "Arimp"},
		{"sdk", "s", "Sdk"},
	}
	for _, tt := range tests {
		t.Run(tt.pkgName, func(t *testing.T) {
			code := generateDllBindingCode([]dllFuncInfo{
				{goName: tt.wantStruct + "_test_func", cName: "test_func", returnType: "void", params: nil},
			}, "test.dll", tt.pkgName, nil, nil, nil)
			wantPattern := fmt.Sprintf("func (%s *%s)", tt.wantRecv, tt.wantStruct)
			if !strings.Contains(code, wantPattern) {
				t.Errorf("generated code want receiver %q, got:\n%s", wantPattern, code)
			}
		})
	}
}

func TestSizeTTypeMappingInDllOutput(t *testing.T) {
	dllPath := filepath.Join("project", "patchKeystone", "keystone2", "dll.go")
	content, err := os.ReadFile(dllPath)
	if err != nil {
		t.Skipf("dll.go not found, skip: %v", err)
	}
	s := string(content)

	if strings.Contains(s, "Size_t") {
		t.Error("Size_t should be mapped to uintptr, but found raw Size_t in generated dll.go")
	}

	if !strings.Contains(s, "Value uintptr") {
		t.Error("ks_option Value parameter should be uintptr (mapped from Size_t)")
	}
	if !strings.Contains(s, "*uintptr") {
		t.Error("ks_asm Encoding_size/Stat_count should be *uintptr (mapped from *Size_t)")
	}

	receiverPattern := "func (k *Keystone)"
	if !strings.Contains(s, receiverPattern) {
		t.Errorf("receiver should be %q (first letter lowercase of Keystone), not 'z'", receiverPattern)
	}
}
