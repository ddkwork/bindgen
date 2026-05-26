package c2go

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

	"github.com/ddkwork/bindgen/cc"
	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
)

func NewMSVCConfig(t testing.TB) *cc.Config {
	abi, err := cc.NewABI(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		t.Fatalf("cc.NewABI: %v", err)
	}
	cfg := &cc.Config{
		ABI:             abi,
		CC:              findClExe(),
		Predefined:      loadMSVCExtraTypes(),
		IncludePaths:    []string{""},
		SysIncludePaths: []string{""},
		FS:              NewFakeHeaderFS(),
	}
	return cfg
}

func Generate(t *testing.T, configs []BindgenConfig) {
	fmt.Println("===  Bindgen ===")

	cfg := NewMSVCConfig(t)
	fmt.Printf("CC: %s\n", cfg.CC)
	fmt.Printf("Predefined: %d bytes\n", len(cfg.Predefined))
	fmt.Printf("IncludePaths: %v\n", cfg.IncludePaths)

	for _, c := range configs {
		ok := ProcessBindgenConfig(t, cfg, c)
		if !ok {
			t.Logf("Skipping go test for %s: generation failed", c.PackageName)
			continue
		}
		stream.FmtDir(c.OutputDir)
		// stream.RunCommandWithDir(c.OutputDir, "go mod tidy")

		if strings.Contains(c.OutputDir, "ARImpRec") {
			oldArch := os.Getenv("GOARCH")
			os.Setenv("GOARCH", "386")
			stream.RunCommandWithDir(c.OutputDir, "go", "test", "-v", "-skip", "TestGenerate", ".")
			os.Setenv("GOARCH", oldArch)
			continue
		}
		stream.RunCommandWithDir(c.OutputDir, "go test -v -skip TestGenerate .")
	}

	fmt.Println("\nAll binding generations completed!")
}

func ProcessBindgenConfig(t *testing.T, cfg *cc.Config, bc BindgenConfig) bool {
	fmt.Printf("\n=== Processing: %s ===\n", bc.PackageName)

	if bc.Predefined != "" {
		cfg.Predefined += bc.Predefined
	}

	cfg.IncludePaths = append(cfg.IncludePaths, bc.HeadersDir)
	cfg.SysIncludePaths = append(cfg.SysIncludePaths, bc.HeadersDir)
	for _, dir := range bc.ExtraIncludeDirs {
		cfg.IncludePaths = append(cfg.IncludePaths, dir)
		cfg.SysIncludePaths = append([]string{dir}, cfg.SysIncludePaths...)
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
		return false
	}
	_ = files

	var sources []cc.Source
	sources = append(sources, cc.Source{Name: "<predefined>", Value: cfg.Predefined})
	sources = append(sources, cc.Source{Name: "<builtin>", Value: cc.Builtin})

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

	var ast *cc.AST
	func() {
		defer func() {
			if r := recover(); r != nil {
				panicMsg := fmt.Sprintf("%v", r)
				buf := make([]byte, 4096)
				n := runtime.Stack(buf, false)
				t.Logf("cc.Translate panicked:\n  %s\nStack:\n%s", panicMsg, string(buf[:n]))
				var sourceInfo strings.Builder
				sourceInfo.WriteString("Sources being translated:\n")
				for _, s := range sources {
					v := fmt.Sprintf("%v", s.Value)
					sourceInfo.WriteString(fmt.Sprintf("  Source: %s (%d bytes)\n", s.Name, len(v)))
				}
				t.Log(sourceInfo.String())
				if strings.Contains(panicMsg, "-->") {
					parts := strings.SplitN(panicMsg, "-->", 2)
					loc := strings.TrimSpace(parts[1])
					lineNo := 0
					if idx := strings.Index(loc, ":"); idx > 0 {
						if n, e := strconv.Atoi(loc[idx+1:]); e == nil {
							lineNo = n
						}
						loc = loc[:idx]
					}
					for _, s := range sources {
						if strings.HasSuffix(loc, s.Name) || strings.HasSuffix(s.Name, loc) {
							content := fmt.Sprintf("%v", s.Value)
							lines := strings.Split(content, "\n")
							t.Logf("  Source file: %s", s.Name)
							if lineNo > 0 && lineNo <= len(lines) {
								start := lineNo - 3
								if start < 0 {
									start = 0
								}
								end := lineNo + 2
								if end > len(lines) {
									end = len(lines)
								}
								for i := start; i < end; i++ {
									marker := "   "
									if i+1 == lineNo {
										marker = ">>>"
									}
									t.Logf("  %s %4d | %s", marker, i+1, lines[i])
								}
							}
							break
						}
					}
				}
				ast = nil
				err = fmt.Errorf("cc.Translate panic: %v", r)
				t.FailNow()
			}
		}()
		ast, err = cc.Translate(cfg, sources)
	}()
	if err != nil {
		t.Logf("Translation error (partial parse): %v", err)
		if len(sources) > 0 {
			lastSrc := sources[len(sources)-1]
			v := fmt.Sprintf("%v", lastSrc.Value)
			t.Logf("Last source: %s (%d bytes)", lastSrc.Name, len(v))
		}
		t.Log("Continuing with partial AST...")
	}

	os.MkdirAll(bc.OutputDir, 0o755)

	if ast == nil {
		t.Log("No AST generated, skipping output")
		return false
	}

	result := Result{
		Structs:  safemap.New[string, StructInfo](true),
		Enums:    safemap.New[string, EnumInfo](true),
		Typedefs: safemap.New[string, TypedefInfo](true),
		Macros:   safemap.New[string, MacroConstInfo](true),
		FnMacros: safemap.New[string, MacroConstInfo](true),
		Imports:  make(map[string]bool),
	}

	result.Macros.Set("PAGE_SIZE", MacroConstInfo{
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
		result.Structs.Set(ps.goName, StructInfo{
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
		sourceFile := FindSourceFileName(l, sources, headerMap)
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
					goStructName := CTagToGoName(structName)
					isPacked := packedFiles[baseName]
					fieldDefs, methodDefs, innerTypes := GenerateStructFields(t, goStructName, isPacked)

					result.Structs.Set(goStructName, StructInfo{
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
				goEnumName := CTagToGoName(enumName)
				members := safemap.New[string, EnumMember]()
				var memberOrder []string
				for _, e := range t.Enumerators() {
					if e == nil {
						continue
					}
					cMemberName := e.Token.SrcStr()
					goMemberName := CEnumMemberToGoName(cMemberName)
					valStr, intVal, hasVal := ExtractEnumeratorValue(e)
					members.Set(goMemberName, EnumMember{
						goName: goMemberName,
						value:  valStr,
						intVal: intVal,
						hasVal: hasVal,
					})
					memberOrder = append(memberOrder, goMemberName)
				}
				result.Enums.Set(goEnumName, EnumInfo{
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
					goUnionName := CTagToGoName(unionName)
					unionSize := t.Size()
					if unionFields := GenerateUnionFields(t, unionSize); unionFields != "" {
						var nestedInner []StructInfo
						for j := 0; j < t.NumFields(); j++ {
							if uf := t.FieldByIndex(j); uf != nil {
								if usv, ok := uf.Type().(*cc.StructType); ok {
									nestedTag := usv.Tag()
									nestedTagStr := string(nestedTag.Src())
									if nestedTagStr != "" && !HasBitfields(usv) {
										nestedF, nestedM, _ := GenerateStructFields(usv, MapCTypeToGo(uf.Type()), false)
										nestedInner = append(nestedInner, StructInfo{
											goName:  MapCTypeToGo(uf.Type()),
											cName:   nestedTagStr,
											fields:  nestedF,
											methods: nestedM,
										})
									}
								}
							}
						}
						result.Structs.Set(goUnionName+"_", StructInfo{
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
					fieldDefs := fmt.Sprintf("\t%s\n", UnionAlignedFFIType(t, unionSize))
					result.Structs.Set(goUnionName, StructInfo{
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
						goTdName := CTagToGoName(tdName)
						switch vt := tdType.(type) {
						case *cc.EnumType:
							enumTag := vt.Tag()
							tagStr := string(enumTag.Src())
							if tagStr != "" {
								result.Typedefs.Set(goTdName, TypedefInfo{
									goName: goTdName,
									goType: CTagToGoName(tagStr),
									source: sourceFile,
								})
							}
						case *cc.StructType:
							structTag := vt.Tag()
							tagStr := string(structTag.Src())
							if tagStr != "" {
								result.Typedefs.Set(goTdName, TypedefInfo{
									goName: goTdName,
									goType: CTagToGoName(tagStr),
									source: sourceFile,
								})
							} else {
								isPacked := packedFiles[baseName]
								fieldDefs, methodDefs, innerTypes := GenerateStructFields(vt, goTdName, isPacked)
								result.Structs.Set(goTdName, StructInfo{
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
							backingName, backingType, bfs := ExtractUnionBitfields(vt)
							if backingName != "" && len(bfs) > 0 {
								fieldDefs := fmt.Sprintf("\t%s %s\n", backingName, backingType)
								var methodDefs []string
								for _, bf := range bfs {
									methodDefs = append(methodDefs, GenerateBitfieldGetterSetter(bf, goTdName, backingName, backingType))
								}
								result.Structs.Set(goTdName, StructInfo{
									goName:  goTdName,
									cName:   tdName,
									fields:  fieldDefs,
									methods: strings.Join(methodDefs, "\n"),
									source:  sourceFile,
								})
							} else {
								unionSize := vt.Size()
								fieldDefs := fmt.Sprintf("\t%s\n", UnionAlignedFFIType(vt, unionSize))
								result.Structs.Set(goTdName, StructInfo{
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
									result.Typedefs.Set(goTdName, TypedefInfo{
										goName: goTdName,
										goType: "*" + CTagToGoName(tagStr),
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
					goTypeName := MapCTypedefToGo(declName, declType)
					goDeclName := CTagToGoName(declName)
					if goTypeName == "" {
						switch vt := declType.(type) {
						case *cc.StructType:
							isPacked := packedFiles[baseName]
							fieldDefs, methodDefs, innerTypes := GenerateStructFields(vt, goDeclName, isPacked)
							result.Structs.Set(goDeclName, StructInfo{
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
							backingName, backingType, bfs := ExtractUnionBitfields(vt)
							if backingName != "" && len(bfs) > 0 {
								fieldDefs := fmt.Sprintf("\t%s %s\n", backingName, backingType)
								var methodDefs []string
								for _, bf := range bfs {
									methodDefs = append(methodDefs, GenerateBitfieldGetterSetter(bf, goDeclName, backingName, backingType))
								}
								result.Structs.Set(goDeclName, StructInfo{
									goName:  goDeclName,
									cName:   declName,
									fields:  fieldDefs,
									methods: strings.Join(methodDefs, "\n"),
									source:  sourceFile,
								})
							} else {
								unionSize := vt.Size()
								fieldDefs := fmt.Sprintf("\t%s\n", UnionAlignedFFIType(vt, unionSize))
								result.Structs.Set(goDeclName, StructInfo{
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
						if st, ok := declType.(*cc.StructType); ok {
							if st.NumFields() == 0 {
								result.Structs.Set(goDeclName, StructInfo{
									goName: goDeclName,
									cName:  declName,
									fields: "",
									source: sourceFile,
								})
							}
						}
						continue
					}
					result.Typedefs.Set(goDeclName, TypedefInfo{
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
			fnMacroNames[macroName] = CMacroNameToGoName(macroName)
		} else if macro.IsConst {
			goConstName := CMacroNameToGoName(macroName)
			if goConstName != macroName {
				constMacroNames[macroName] = goConstName
			}
		}
	}

	headerOrder := bc.HeaderOrder

	orderIncludes := make(map[string]map[string]bool)
	includeRe := regexp.MustCompile(`#\s*include\s+["<]([^">]+)[">]`)
	var collectIncludes func(orderFile string) map[string]bool
	collectIncludes = func(orderFile string) map[string]bool {
		includes := make(map[string]bool)
		queue := []string{orderFile}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]

			content, ok := headerMap[current]
			if !ok {
				var found bool
				for key, val := range headerMap {
					if strings.HasSuffix(key, "/"+current) || strings.HasSuffix(key, "\\"+current) || key == current {
						content = val
						found = true
						break
					}
				}
				if !found {
					continue
				}
			}

			for key := range headerMap {
				if key == current || strings.HasSuffix(key, "/"+current) || strings.HasSuffix(key, "\\"+current) {
					includes[key] = true
				}
			}

			for _, match := range includeRe.FindAllStringSubmatch(content, -1) {
				incName := match[1]
				for key := range headerMap {
					if key == incName || strings.HasSuffix(key, "/"+incName) || strings.HasSuffix(key, "\\"+incName) {
						if !includes[key] {
							includes[key] = true
							queue = append(queue, key)
						}
					}
				}
			}
		}
		return includes
	}
	for _, orderFile := range headerOrder {
		orderIncludes[orderFile] = collectIncludes(orderFile)
	}

	sourceToOrder := make(map[string]string)
	for _, orderFile := range headerOrder {
		for key := range orderIncludes[orderFile] {
			sourceToOrder[key] = orderFile
			sourceToOrder[filepath.Base(key)] = orderFile
		}
	}

	for name, m := range ast.Macros {
		pos := m.Name.Position()
		sourceFile := filepath.Base(pos.Filename)
		if !strings.HasSuffix(sourceFile, ".h") {
			continue
		}
		if _, ok := sourceToOrder[sourceFile]; !ok {
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

			paramType := "uint32"
			if needsUint64(bodyStr) {
				paramType = "uint64"
			}

			var goParamList []string
			var goParamNames []string
			var cParamNames []string
			for _, p := range m.Params {
				cPName := string(p.Src())
				cParamNames = append(cParamNames, cPName)
				goPName := strings.ToUpper(cPName[:1]) + cPName[1:]
				goParamList = append(goParamList, goPName+" "+paramType)
				goParamNames = append(goParamNames, goPName)
			}

			goName := CMacroNameToGoName(name)
			goBody := ConvertCMacroBodyToGo(bodyStr, cParamNames, goParamNames)
			result.FnMacros.Set(goName, MacroConstInfo{
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

		skipMacros := map[string]bool{
			"APIENTRY": true, "GLAPIENTRY": true, "GLAPI": true,
			"__stdcall": true, "__cdecl": true, "__fastcall": true,
			"CALLBACK": true, "WINAPI": true,
			"WINGDIAPI": true, "GLFWAPI": true,
			"GLFW_EXPOSE_NATIVE_": true,
		}
		if skipMacros[name] {
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

		if !IsValidGoMacroValue(valStr) {
			continue
		}

		goVal := CleanCMacroValue(valStr)

		// Skip macros that are just function aliases (e.g., #define Foo FooW)
		if goVal == valStr && !strings.Contains(valStr, "(") && !strings.Contains(valStr, "0x") && !strings.Contains(valStr, "0X") {
			isPureIdent := true
			hasLetter := false
			for _, c := range valStr {
				if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_') {
					isPureIdent = false
					break
				}
				if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_' {
					hasLetter = true
				}
			}
			_, isKnownConst := constMacroNames[valStr]
			if isPureIdent && hasLetter && !isKnownConst {
				continue
			}
		}

		isVar := false
		if strings.Contains(goVal, "unsafe.Sizeof(") {
			isVar = true
		}

		identsInVal := ExtractIdentifiers(goVal)
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
				goVal = ReplaceWord(goVal, id, goCName)
			}
		}

		goType := DetermineMacroGoType(valStr)
		if isVar && strings.Contains(goVal, "unsafe.Sizeof(") {
			goType = "int"
		}

		goName := CMacroNameToGoName(name)
		result.Macros.Set(goName, MacroConstInfo{
			goName:  goName,
			goType:  goType,
			value:   valStr,
			goValue: goVal,
			source:  sourceFile,
			isVar:   isVar,
			lineNo:  int(pos.Line),
		})
	}

	headerOrder = bc.HeaderOrder
	definedTypedefs := make(map[string]bool)
	definedStructs := make(map[string]bool)
	headerOrderSet := make(map[string]bool)
	for _, f := range headerOrder {
		headerOrderSet[f] = true
		headerOrderSet[filepath.Base(f)] = true
	}

	globalDefinedConsts := make(map[string]bool)
	for _, fileName := range headerOrder {
		if _, ok := headerMap[fileName]; !ok {
			continue
		}
		baseName := strings.TrimSuffix(fileName, ".h")
		baseName = filepath.Base(baseName)
		baseName = strings.TrimPrefix(baseName, "gen_cabi_")
		var content strings.Builder
		content.WriteString(fmt.Sprintf("package %s\n\n", bc.PackageName))

		hasContent := false
		structCount, enumCount, tdCount := 0, 0, 0
		definedConsts := globalDefinedConsts

		sourceMatch := func(src string) bool {
			srcBase := filepath.Base(src)
			if src == fileName || srcBase == fileName {
				return true
			}
			if src == "<predefined>" || src == "<builtin>" {
				return true
			}
			// Fake headers (windows.h, stddef.h, etc.) should be included
			fakeHeaders := map[string]bool{
				"windows.h": true, "stddef.h": true, "stdint.h": true,
				"stdlib.h": true, "stdio.h": true, "string.h": true,
				"vcruntime.h": true, "sal.h": true, "vadefs.h": true,
			}
			if fakeHeaders[src] || fakeHeaders[srcBase] {
				return true
			}
			if headerOrderSet[src] || headerOrderSet[srcBase] {
				return true
			}
			orderFile, ok := sourceToOrder[src]
			if !ok {
				orderFile, ok = sourceToOrder[srcBase]
			}
			return ok && (orderFile == fileName || filepath.Base(orderFile) == fileName || strings.TrimSuffix(orderFile, ".h") == fileName || strings.TrimSuffix(filepath.Base(orderFile), ".h") == fileName)
		}
		var aliasTypes []TypedefInfo
		var funcTypes []TypedefInfo
		var skipStructNames map[string]bool
		skipBasicTypedefs := map[string]bool{
			"Cardinal": true, "Integer": true, "PChar": true,
			"PCardinal": true, "HMODULE": true, "IRSaveOEPToFile": true,
			// Windows handle types - map to Go primitives in function signatures
			"HWND": true, "HANDLE": true, "HINSTANCE": true, "HDC": true,
			"HMENU": true, "HICON": true, "HCURSOR": true, "HBRUSH": true,
			"HFONT": true, "HPEN": true, "HBITMAP": true, "HRGN": true,
			"HPALETTE": true, "HACCEL": true, "HHOOK": true,
			"WPARAM": true, "LPARAM": true,
		}
		for _, td := range result.Typedefs.Range() {
			if !sourceMatch(td.source) {
				continue
			}
			if skipBasicTypedefs[td.goName] {
				continue
			}
			if definedTypedefs[td.goName] {
				continue
			}
			if result.Enums.Has(td.goName) {
				continue
			}
			if result.Structs.Has(td.goName) {
				if IsGoPrimitiveType(td.goType) {
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
			definedTypedefs[td.goName] = true
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
			goType := td.goType
			if idx := strings.LastIndex(goType, ") "); idx >= 0 {
				goType = goType[:idx+1] + " uintptr"
			} else {
				goType = strings.TrimSuffix(goType, ")") + ") uintptr"
			}
			content.WriteString(fmt.Sprintf("type %s %s\n\n", td.goName, goType))
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
			GenerateEnumStringer(&content, ei)
		}

		usedNames := make(map[string]bool)
		var structDecls []string
		for _, si := range result.Structs.Range() {
			if !sourceMatch(si.source) {
				continue
			}
			if skipStructNames != nil && skipStructNames[si.goName] {
				continue
			}
			if definedStructs[si.goName] {
				continue
			}
			structCount++
			hasContent = true
			definedStructs[si.goName] = true
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
			structDecl := fmt.Sprintf("\t%s struct {", goName)
			if si.fields != "" {
				structDecl += "\n" + si.fields + "\t}"
			} else {
				structDecl += "}"
			}
			structDecl += fmt.Sprintf(" // %s:%d -> %s\n", filepath.Base(src), si.lineNo, si.cName)
			structDecls = append(structDecls, structDecl)
			if si.methods != "" {
				content.WriteString(si.methods)
			}
			if IsFFIUnion(si.fields) && !strings.HasSuffix(si.goName, "_") {
				GenerateUnionAccessorMethods(&content, si.goName, result.Structs)
			}
		}
		if len(structDecls) > 0 {
			content.WriteString("type (\n")
			for _, decl := range structDecls {
				content.WriteString(decl)
			}
			content.WriteString(")\n\n")
		}

		var constItems []MacroConstInfo
		var varItems []MacroConstInfo
		for _, mc := range result.Macros.Range() {
			if !sourceMatch(mc.source) {
				continue
			}
			if IsTypeAliasMacro(mc.goValue) {
				continue
			}
			if mc.isVar {
				if !definedConsts[mc.goName] {
					varItems = append(varItems, mc)
				}
			} else {
				if !definedConsts[mc.goName] {
					constItems = append(constItems, mc)
				}
			}
		}
		for _, mc := range constItems {
			definedConsts[mc.goName] = true
		}
		for _, mc := range varItems {
			definedConsts[mc.goName] = true
		}
		slices.SortStableFunc(constItems, func(a, b MacroConstInfo) int {
			if a.source != b.source {
				return strings.Compare(a.source, b.source)
			}
			return a.lineNo - b.lineNo
		})
		slices.SortStableFunc(varItems, func(a, b MacroConstInfo) int {
			if a.source != b.source {
				return strings.Compare(a.source, b.source)
			}
			return a.lineNo - b.lineNo
		})

		for _, mc := range result.FnMacros.Range() {
			if !sourceMatch(mc.source) {
				continue
			}
			if definedConsts[mc.goName] {
				continue
			}
			definedConsts[mc.goName] = true
			if !IsValidGoFnMacroBody(mc.value) {
				continue
			}
			mc.goBody = StripCStyleCast(mc.goBody)
			cleanedBody := CleanCMacroValue(mc.goBody)
			if !IsValidGoFnMacroBody(cleanedBody) {
				continue
			}
			hasContent = true
			simplifiedParams, simplifiedBody, updatedVarItems := SimplifyFnMacro(mc, varItems)
			isBool := IsBoolExpr(cleanedBody)
			retType := "uint32"
			retCast := "uint32"
			if strings.Contains(mc.goParams, "uint64") {
				retType = "uint64"
				retCast = "uint64"
			}
			if simplifiedParams != "" {
				body := CleanCMacroValue(ResolveMacroValueNames(simplifiedBody))
				if isBool {
					content.WriteString(fmt.Sprintf("func %s(%s) bool {\n\treturn %s\n}\n\n", mc.goName, simplifiedParams, ConvertCBoolExpr(body)))
				} else {
					content.WriteString(fmt.Sprintf("func %s(%s) %s {\n\treturn %s(%s)\n}\n\n", mc.goName, simplifiedParams, retType, retCast, body))
				}
				varItems = updatedVarItems
			} else {
				body := CleanCMacroValue(ResolveMacroValueNames(cleanedBody))
				if isBool {
					content.WriteString(fmt.Sprintf("func %s(%s) bool {\n\treturn %s\n}\n\n", mc.goName, mc.goParams, ConvertCBoolExpr(body)))
				} else {
					content.WriteString(fmt.Sprintf("func %s(%s) %s {\n\treturn %s(%s)\n}\n\n", mc.goName, mc.goParams, retType, retCast, body))
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
				var errorCodeItems []MacroConstInfo
				var nonErrorCodeItems []MacroConstInfo
				for _, mc := range constItems {
					if strings.HasPrefix(mc.goName, "DebuggerError") || strings.HasPrefix(mc.goName, "DebuggerOperation") {
						errorCodeItems = append(errorCodeItems, mc)
					} else {
						nonErrorCodeItems = append(nonErrorCodeItems, mc)
					}
				}
				content.WriteString("// Source: ErrorCodes.h -> Error codes\n")
				GenerateErrorCodeEnum(&content, errorCodeItems, baseName)
				PropagateConstTypes(nonErrorCodeItems)
				content.WriteString("\n// Source: " + baseName + ".h -> Macro constants\nconst (\n")
				for _, mc := range nonErrorCodeItems {
					content.WriteString(fmt.Sprintf("\t%s %s = %s\n", mc.goName, mc.goType, ResolveMacroValueNames(mc.goValue)))
				}
				content.WriteString(")\n\n")
			} else {
				PropagateConstTypes(constItems)
				content.WriteString(fmt.Sprintf("// Source: %s -> Macro constants\n", baseName+".h"))
				content.WriteString("const (\n")
				for _, mc := range constItems {
					content.WriteString(fmt.Sprintf("\t%s %s = %s\n", mc.goName, mc.goType, ResolveMacroValueNames(mc.goValue)))
				}
				content.WriteString(")\n\n")
			}
		}

		if len(varItems) > 0 {
			hasContent = true
			PropagateConstTypes(varItems)
			content.WriteString(fmt.Sprintf("// Source: %s -> Macro variables\n", baseName+".h"))
			content.WriteString("var (\n")
			for _, mc := range varItems {
				val := ResolveMacroValueNames(mc.goValue)
				if strings.Contains(val, "unsafe.Sizeof(") {
					val = WrapConstRefsForInt(val)
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

		finalContent := AddImports(content.String(), bc.PackageName, result.Imports)

		finalContent = ReplaceRemainingSizeof(finalContent)

		outputFile := filepath.Join(bc.OutputDir, baseName+".go")
		err := os.WriteFile(outputFile, []byte(finalContent), 0o644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing file %s: %v\n", outputFile, err)
			continue
		}
		fmt.Printf("Generated: %s (%d structs, %d enums, %d typedefs)\n", outputFile, structCount, enumCount, tdCount)
	}

	fmt.Printf("\n%s binding generation completed!\n", bc.PackageName)

	if bc.BindDll {
		GenerateDllBinding(t, ast, bc, result.Typedefs, result.Enums, sources, headerMap)
	}
	return true
}

func AddImports(content string, packageName string, extraImports map[string]bool) string {
	imports := make(map[string]bool)
	maps.Copy(imports, extraImports)
	if strings.Contains(content, "fmt.Sprintf") && !strings.Contains(content, `"fmt"`) {
		imports[`"fmt"`] = true
	}
	if strings.Contains(content, "unsafe.") && !strings.Contains(content, `"unsafe"`) {
		imports[`"unsafe"`] = true
	}
	if strings.Contains(content, "windows.") && !strings.Contains(content, `"golang.org/x/sys/windows"`) {
		imports[`"golang.org/x/sys/windows"`] = true
	}
	if strings.Contains(content, "time.Duration") && !strings.Contains(content, `"time"`) {
		imports[`"time"`] = true
	}
	if (strings.Contains(content, "bits.ReverseBytes") || strings.Contains(content, "bits.Reverse(")) && !strings.Contains(content, `"math/bits"`) {
		imports[`"math/bits"`] = true
	}
	if len(imports) == 0 {
		content = strings.Replace(content, "%%_IMPORTLIBS_%%\n", "", 1)
		return content
	}
	var sortedImports []string
	for imp := range imports {
		sortedImports = append(sortedImports, imp)
	}
	sort.Strings(sortedImports)
	if strings.Contains(content, "import (") {
		for _, imp := range sortedImports {
			content = strings.Replace(content, "import (", "import (\n\t"+imp, 1)
		}
		return content
	}
	var importBlock strings.Builder
	importBlock.WriteString("import (\n")
	for _, imp := range sortedImports {
		importBlock.WriteString("\t" + imp + "\n")
	}
	importBlock.WriteString(")\n\n")
	
	needLpOverlapped := strings.Contains(content, "LPOVERLAPPED") && !strings.Contains(content, "type LPOVERLAPPED")
	if needLpOverlapped {
		if !strings.Contains(importBlock.String(), `"syscall"`) {
			importBlockStr := importBlock.String()
			importBlockStr = strings.Replace(importBlockStr, "import (\n", "import (\n\t\"syscall\"\n", 1)
			importBlock = strings.Builder{}
			importBlock.WriteString(importBlockStr)
		}
	}
	
	if strings.Contains(content, "%%_IMPORTLIBS_%%") {
		result := strings.Replace(content, "%%_IMPORTLIBS_%%", importBlock.String(), 1)
		if needLpOverlapped {
			typeDef := "type LPOVERLAPPED = *syscall.Overlapped\n\n"
			importEnd := ")\n\n"
			result = strings.Replace(result, importEnd, importEnd+typeDef, 1)
		}
		return result
	}
	packageLine := fmt.Sprintf("package %s\n", packageName)
	if strings.Contains(content, packageLine+"\n") {
		result := strings.Replace(content, packageLine+"\n", packageLine+"\n"+importBlock.String(), 1)
		if needLpOverlapped {
			typeDef := "type LPOVERLAPPED = *syscall.Overlapped\n\n"
			importEnd := ")\n\n"
			result = strings.Replace(result, importEnd, importEnd+typeDef, 1)
		}
		return result
	}
	result := strings.Replace(content, packageLine, packageLine+"\n"+importBlock.String(), 1)
	if needLpOverlapped {
		typeDef := "type LPOVERLAPPED = *syscall.Overlapped\n\n"
		importEnd := ")\n\n"
		result = strings.Replace(result, importEnd, importEnd+typeDef, 1)
	}
	return result
}

func GenerateEnumStringer(content *strings.Builder, ei EnumInfo) {
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
		humanName := CEnumMemberToHumanString(m.goName, ei.goName)
		content.WriteString(fmt.Sprintf("\tcase %s:\n\t\treturn %q\n", m.goName, humanName))
	}
	content.WriteString("\tdefault:\n")
	content.WriteString(fmt.Sprintf("\t\treturn fmt.Sprintf(\"%s(0x%%X)\", uint32(%s))\n", ei.goName, receiver))
	content.WriteString("\t}\n")
	content.WriteString("}\n\n")
}

func CEnumMemberToHumanString(goName, typeName string) string {
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

func IsBoolExpr(body string) bool {
	boolOps := []string{"==", "!=", "<=", ">=", "&&", "||"}
	for _, op := range boolOps {
		if strings.Contains(body, op) {
			return true
		}
	}
	return HasNotParen(body)
}

func ConvertCBoolExpr(body string) string {
	if HasDoubleNotParen(body) {
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

func IsTypeAliasMacro(val string) bool {
	trimmed := strings.TrimSpace(val)
	if ContainsDigit(trimmed) {
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

func IsFFIUnion(fields string) bool {
	fields = strings.TrimSpace(fields)
	return strings.HasPrefix(fields, "_ [") && strings.HasSuffix(fields, "byte")
}

func GenerateUnionAccessorMethods(content *strings.Builder, unionGoName string, structs interface {
	Get(string) (StructInfo, bool)
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
		if nested, ok := structs.Get(fieldType + "_"); ok && !IsFFIUnion(nested.fields) {
			content.WriteString(fmt.Sprintf("func (u *%s) %s() *%s { return (*%s)(unsafe.Pointer(u)) }\n",
				unionGoName, fieldName+"Fields", fieldType+"_", fieldType+"_"))
		}
	}
}

func IsValidGoFnMacroBody(body string) bool {
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
	if IsULiteral(body) {
		return false
	}
	if len(body) > 0 && body[0] == ':' {
		return false
	}
	if IsSimpleNumber(body) {
		return false
	}
	if IsLargeHex(body) {
		return false
	}
	return true
}

func IsCStyleCastMacro(body string) bool {
	return IsCStyleCast(body)
}

func IsSimpleNumber(s string) bool {
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

func IsValidGoMacroValue(val string) bool {
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
	if strings.Contains(val, "__cdecl") || strings.Contains(val, "__stdcall") || strings.Contains(val, "__fastcall") || strings.Contains(val, "__thiscall") || strings.Contains(val, "__vectorcall") {
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
	if IsCharLiteral(val) {
		return false
	}
	if strings.Contains(val, "[") && strings.Contains(val, "]") {
		if strings.Contains(val, "__DATE__") || strings.Contains(val, "__TIME__") || strings.Contains(val, "BUILD_MONTH") {
			return false
		}
	}
	return true
}

func CleanCMacroValue(val string) string {
	result := val
	result = strings.ReplaceAll(result, "~", "^")
	result = StripAllNumberSuffixes(result)
	for {
		typeName, found := ParseSizeofType(result)
		if !found {
			break
		}
		goType := MapCTypeNameToGoForSizeof(typeName)
		var replacement string
		if IsGoPrimitiveType(goType) {
			replacement = fmt.Sprintf("int(unsafe.Sizeof(%s(0)))", goType)
		} else {
			replacement = fmt.Sprintf("int(unsafe.Sizeof(%s{}))", goType)
		}
		result = strings.ReplaceAll(result, "sizeof("+typeName+")", replacement)
	}
	return result
}

func ReplaceRemainingSizeof(content string) string {
	for {
		idx := strings.Index(content, "sizeof(")
		if idx < 0 {
			break
		}
		parenStart := idx + len("sizeof(")
		if parenStart >= len(content) {
			break
		}
		depth := 1
		end := -1
		for i := parenStart; i < len(content); i++ {
			if content[i] == '(' {
				depth++
			} else if content[i] == ')' {
				depth--
				if depth == 0 {
					end = i
					break
				}
			}
		}
		if end < 0 {
			break
		}
		typeName := strings.TrimSpace(content[parenStart:end])
		goType := MapCTypeNameToGoForSizeof(typeName)
		var replacement string
		if goType == "" || typeName == "" {
			replacement = "0"
		} else if IsGoPrimitiveType(goType) {
			replacement = fmt.Sprintf("int(unsafe.Sizeof(%s(0)))", goType)
		} else {
			replacement = fmt.Sprintf("int(unsafe.Sizeof(%s{}))", goType)
		}
		oldExpr := content[idx : end+1]
		content = strings.ReplaceAll(content, oldExpr, replacement)
	}
	return content
}

func StripAllNumberSuffixes(s string) string {
	runes := []rune(s)
	var buf strings.Builder
	i := 0
	for i < len(runes) {
		if runes[i] >= '0' && runes[i] <= '9' {
			start := i
			for i < len(runes) && ((runes[i] >= '0' && runes[i] <= '9') ||
				(runes[i] >= 'a' && runes[i] <= 'f') || (runes[i] >= 'A' && runes[i] <= 'F') ||
				runes[i] == 'x' || runes[i] == 'X') {
				i++
			}
			buf.WriteString(string(runes[start:i]))
			for i < len(runes) {
				ch := runes[i]
				if ch == 'U' || ch == 'u' || ch == 'L' || ch == 'l' || ch == 'Z' || ch == 'z' {
					i++
				} else if ch == 'I' || ch == 'i' {
					i++
					for i < len(runes) && runes[i] >= '0' && runes[i] <= '9' {
						i++
					}
				} else {
					break
				}
			}
		} else {
			buf.WriteRune(runes[i])
			i++
		}
	}
	return buf.String()
}

func MapCTypeNameToGoForSizeof(cType string) string {
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
	case "Cardinal":
		return "uint32"
	case "Integer":
		return "int32"
	case "HMODULE":
		return "uintptr"
	case "IRSaveOEPToFile":
		return "uint32"
	default:
		return CTagToGoName(cType)
	}
}

func ResolveMacroValueNames(val string) string {
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
		result = ReplaceWord(result, cName, goName)
	}
	return result
}

func WrapConstRefsForInt(val string) string {
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
		result = ReplaceWord(result, ref, "int("+ref+")")
	}
	return result
}

func DetermineMacroGoType(val string) string {
	cleaned := CleanCMacroValue(val)
	if strings.HasPrefix(cleaned, "0x") || strings.HasPrefix(cleaned, "0X") {
		if len(cleaned) > 10 {
			return "uint64"
		}
		return "uint32"
	}
	if strings.HasPrefix(cleaned, "-") || strings.Contains(cleaned, "(-") {
		if _, found := FindLargeNum(cleaned); found {
			return "int64"
		}
		if strings.Contains(cleaned, "<<") || strings.Contains(cleaned, ">>") {
			return "int64"
		}
		return "int32"
	}
	if _, found := FindLargeNum(cleaned); found {
		return "uint64"
	}
	if shiftVal, found := FindShiftAmount(cleaned); found && shiftVal >= 32 {
		return "uint64"
	}
	return "uint32"
}

func PropagateConstTypes(items []MacroConstInfo) {
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
					if HasWord(mc.goValue, refName) {
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
					if HasWord(mc.goValue, refName) {
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
	for i, mc := range items {
		if strings.HasPrefix(mc.goValue, "(-") || strings.HasPrefix(mc.goValue, "-") {
			for refName, refType := range typeMap {
				if HasWord(mc.goValue, refName) && refType == "uint32" {
					items[i].goType = "int32"
					typeMap[mc.goName] = "int32"
					for j := range items {
						if items[j].goName == refName {
							items[j].goType = "int32"
							break
						}
					}
					typeMap[refName] = "int32"
				}
			}
		}
	}
}

func SimplifyFnMacro(fn MacroConstInfo, varItems []MacroConstInfo) (string, string, []MacroConstInfo) {
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
		args := SplitCallArgs(argsStr)
		if len(args) > 0 {
			calls = append(calls, callInfo{index: i, args: args})
		}
	}
	if len(calls) == 0 {
		return "", "", varItems
	}
	params := ParseGoParams(fn.goParams)
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
			newBody = ReplaceWord(newBody, params[j].name, constVal[j])
		}
	}
	newVarItems := make([]MacroConstInfo, len(varItems))
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

func ParseGoParams(params string) []struct{ name, typ string } {
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

func SplitCallArgs(argsStr string) []string {
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

func GenerateErrorCodeEnum(content *strings.Builder, items []MacroConstInfo, baseName string) {
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
		humanName := CMacroNameToHumanString(mc.goName)
		content.WriteString(fmt.Sprintf("\tcase %s:\n\t\treturn %q\n", mc.goName, humanName))
	}
	content.WriteString("\tdefault:\n")
	content.WriteString("\t\treturn fmt.Sprintf(\"DebuggerErrorCode(0x%X)\", uint32(c))\n")
	content.WriteString("\t}\n")
	content.WriteString("}\n\n")
}

func CMacroNameToHumanString(goName string) string {
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

func ConvertCMacroBodyToGo(body string, cParamNames, goParamNames []string) string {
	result := body
	for i, cName := range cParamNames {
		goName := goParamNames[i]
		result = strings.ReplaceAll(result, cName, goName)
	}
	result = strings.ReplaceAll(result, "~", "^")
	result = strings.ReplaceAll(result, "_byteswap_ulong(", "bits.ReverseBytes32(")
	result = strings.ReplaceAll(result, "_byteswap_uint64(", "bits.ReverseBytes64(")
	result = strings.ReplaceAll(result, "_byteswap_ushort(", "bits.ReverseBytes16(")
	result = strings.ReplaceAll(result, "UINT64_C(", "uint64(")
	result = strings.ReplaceAll(result, "INT64_C(", "int64(")
	for {
		start := strings.Index(result, "__has_include(")
		if start < 0 {
			break
		}
		end := strings.Index(result[start:], ")")
		if end < 0 {
			break
		}
		result = result[:start] + "0" + result[start+end+1:]
	}
	result = strings.ReplaceAll(result, "! (", "!(")
	for i := 0; i < len(result)-1; i++ {
		if result[i] == '!' && i+1 < len(result) && result[i+1] == '(' {
			result = result[:i] + "!(" + result[i+2:]
			break
		}
	}
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
