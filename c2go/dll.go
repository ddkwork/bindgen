package c2go

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/cc"
	"github.com/ddkwork/golibrary/std/safemap"
)

var typeDeclRe = regexp.MustCompile(`^type\s+(\w+)\s+`)

func ensureImports(src string, required []string) string {
	importRe := regexp.MustCompile(`(?s)import\s*\((.*?)\)`)
	m := importRe.FindStringSubmatchIndex(src)
	if m == nil {
		var sb strings.Builder
		sb.WriteString("import (\n")
		for _, imp := range required {
			sb.WriteString(fmt.Sprintf("\t%q\n", imp))
		}
		sb.WriteString(")\n\n")
		packageRe := regexp.MustCompile(`(?s)(package\s+\w+\n\n?)`)
		pm := packageRe.FindStringSubmatchIndex(src)
		if pm != nil {
			return src[:pm[1]] + sb.String() + src[pm[1]:]
		}
		return sb.String() + src
	}
	importBlock := src[m[2]:m[3]]
	existing := map[string]bool{}
	for _, line := range strings.Split(importBlock, "\n") {
		line = strings.TrimSpace(line)
		line = strings.Trim(line, `"`)
		if line != "" && !strings.HasPrefix(line, "//") && line != "(" && line != ")" {
			existing[line] = true
		}
	}
	var toAdd []string
	for _, imp := range required {
		if !existing[imp] {
			toAdd = append(toAdd, imp)
		}
	}
	if len(toAdd) == 0 {
		return src
	}
	var addLines strings.Builder
	for _, imp := range toAdd {
		addLines.WriteString(fmt.Sprintf("\t%q\n", imp))
	}
	newImportBlock := importBlock
	if !strings.HasSuffix(strings.TrimSpace(newImportBlock), "\n") {
		newImportBlock += "\n"
	}
	newImportBlock += addLines.String()
	return src[:m[2]] + newImportBlock + src[m[3]:]
}

func collectExistingTypes(dir string) map[string]bool {
	types := map[string]bool{}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return types
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") {
			continue
		}
		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			line = strings.TrimSpace(line)
			if m := typeDeclRe.FindStringSubmatch(line); len(m) > 1 {
				types[m[1]] = true
			}
		}
	}
	return types
}

func GenerateDllBinding(t *testing.T, ast *cc.AST, bc BindgenConfig, typedefs *safemap.M[string, TypedefInfo], enums *safemap.M[string, EnumInfo], sources []cc.Source, headerMap map[string]string) {
	type funcWithSource struct {
		fn     DllFuncInfo
		source string
	}
	var funcs []funcWithSource
	funcTypeNames := map[string]bool{}
	seenFuncNames := map[string]bool{}
	for l := ast.TranslationUnit; l != nil; l = l.TranslationUnit {
		ed := l.ExternalDeclaration
		if ed == nil || ed.Case != cc.ExternalDeclarationDecl {
			continue
		}
		decl := ed.Declaration
		if decl == nil || decl.DeclarationSpecifiers == nil || decl.InitDeclaratorList == nil {
			continue
		}
		sourceFile := FindSourceFileName(l, sources, headerMap)
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
			if seenFuncNames[name] {
				continue
			}
			seenFuncNames[name] = true
			dtorType := dtor.Type()
			if dtorType == nil {
				continue
			}
			ft, ok := dtorType.(*cc.FunctionType)
			if !ok {
				if pt, isPtr := dtorType.(*cc.PointerType); isPtr {
					if _, isFn := pt.Elem().(*cc.FunctionType); isFn {
						funcTypeNames[CTagToGoName(name)] = true
					}
				}
				continue
			}
			goRet := MapCTypeToGo(ft.Result())
			if goRet == "" {
				goRet = "uintptr"
			}
			var params []DllFuncParam
			paramIdx := 0
			for _, p := range ft.Parameters() {
				pName := ""
				if p.Declarator != nil {
					pName = p.Declarator.Name()
				}
				pGoType := MapCTypeToGo(p.Type())
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
				params = append(params, DllFuncParam{
					goName:    CTagToGoName(pName),
					goType:    pGoType,
					isPointer: isPtr,
				})
			}
			funcs = append(funcs, funcWithSource{
				fn: DllFuncInfo{
					goName:     CTagToGoName(name),
					cName:      name,
					returnType: goRet,
					params:     params,
				},
				source: sourceFile,
			})
		}
	}
	fmt.Printf("Found %d exported functions for DLL binding\n", len(funcs))

	var allFuncs []DllFuncInfo
	for _, f := range funcs {
		allFuncs = append(allFuncs, f.fn)
	}
	existingTypes := collectExistingTypes(bc.OutputDir)
	dllCommonCode := generateDllCommonCode(bc.DllName, bc.PackageName, allFuncs, typedefs, enums, existingTypes)
	dllFile := filepath.Join(bc.OutputDir, "dll.go")
	os.WriteFile(dllFile, []byte(dllCommonCode), 0o644)
	fmt.Printf("Generated: %s\n", dllFile)

	sourceToOrder := make(map[string]string)
	for _, orderFile := range bc.HeaderOrder {
		for src := range headerMap {
			if src == orderFile || strings.HasSuffix(src, "/" + orderFile) || strings.HasSuffix(src, "\\" + orderFile) {
				sourceToOrder[src] = orderFile
			}
		}
	}
	firstOrder := ""
	if len(bc.HeaderOrder) > 0 {
		firstOrder = bc.HeaderOrder[0]
	}

	type orderGroup struct {
		orderFile string
		funcs     []DllFuncInfo
	}
	groups := make(map[string]*orderGroup)
	seenGroupFuncs := map[string]bool{}
	for _, f := range funcs {
		if seenGroupFuncs[f.fn.cName] {
			continue
		}
		seenGroupFuncs[f.fn.cName] = true
		orderFile, ok := sourceToOrder[f.source]
		if !ok {
			orderFile = firstOrder
		}
		if orderFile == "" {
			continue
		}
		g, ok := groups[orderFile]
		if !ok {
			g = &orderGroup{orderFile: orderFile}
			groups[orderFile] = g
		}
		g.funcs = append(g.funcs, f.fn)
	}

	for _, g := range groups {
		baseName := strings.TrimSuffix(g.orderFile, ".h")
		baseName = filepath.Base(baseName)
		baseName = strings.TrimPrefix(baseName, "gen_cabi_")
		outputFile := filepath.Join(bc.OutputDir, baseName+".go")
		dllCode := generateDllBindingCodeForFile(g.funcs, bc.PackageName, funcTypeNames, typedefs, enums)

		existing, err := os.ReadFile(outputFile)
		if err == nil {
			existingStr := string(existing)
			var newFuncs []DllFuncInfo
			for _, f := range g.funcs {
				goName := strings.Split(f.goName, "@")[0]
				found := strings.Contains(existingStr, goName+"(")
				if !found {
					newFuncs = append(newFuncs, f)
				}
			}
			if len(newFuncs) == 0 {
				continue
			}
			dllCode = generateDllBindingCodeForFile(newFuncs, bc.PackageName, funcTypeNames, typedefs, enums)
			result := existingStr + "\n" + dllCode
			result = AddImports(result, bc.PackageName, nil)
			os.WriteFile(outputFile, []byte(result), 0o644)
		} else {
			var fullCode strings.Builder
			fullCode.WriteString(fmt.Sprintf("package %s\n\n", bc.PackageName))
			fullCode.WriteString("%%_IMPORTLIBS_%%\n")
			fullCode.WriteString(dllCode)
			result := fullCode.String()
			result = AddImports(result, bc.PackageName, nil)
			os.WriteFile(outputFile, []byte(result), 0o644)
		}
		
		if data, err := os.ReadFile(outputFile); err == nil {
			content := string(data)
			if strings.Contains(content, "LPOVERLAPPED") && !strings.Contains(content, "type LPOVERLAPPED") {
				if !strings.Contains(content, `"syscall"`) {
					content = strings.Replace(content, "import (\n", "import (\n\t\"syscall\"\n", 1)
				}
				typeDef := "type LPOVERLAPPED = *syscall.Overlapped\n\n"
				importEnd := ")\n\n"
				content = strings.Replace(content, importEnd, importEnd+typeDef, 1)
				os.WriteFile(outputFile, []byte(content), 0o644)
			}
		}
		fmt.Printf("Generated: %s (%d functions)\n", outputFile, len(g.funcs))
	}

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

func generateDllCommonCode(dllName, pkgName string, allFuncs []DllFuncInfo, typedefs *safemap.M[string, TypedefInfo], enums *safemap.M[string, EnumInfo], existingTypes map[string]bool) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("package %s\n\n", pkgName))
	sb.WriteString(`
	import (
	"crypto/sha256"
	_ "embed"
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

`)

	receiverType := strings.ToUpper(pkgName[:1]) + pkgName[1:]
	sb.WriteString(fmt.Sprintf("type %s struct{}\n\n", receiverType))

	sb.WriteString(fmt.Sprintf("//go:embed %s\nvar dllBytes []byte\n\n", dllName))

	sb.WriteString(`var (
	dll     *windows.LazyDLL
	procCache = make(map[string]*windows.LazyProc)
)

func init() {
	dir, err := os.UserCacheDir()
	if err != nil {
		panic(err)
	}
	dir = filepath.Join(dir, "` + pkgName + `_dll_cache")
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(err)
	}
	sha := sha256.Sum256(dllBytes)
	dllName := fmt.Sprintf("%s-%s.dll", "` + strings.TrimSuffix(dllName, ".dll") + `", base64.RawURLEncoding.EncodeToString(sha[:]))
	filePath := filepath.Join(dir, dllName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, dllBytes, 0644); err != nil {
			panic(err)
		}
	}
	// 用绝对路径加载，避免依赖进程级的 SetDllDirectory 全局状态。
	// SetDllDirectory 是进程级单值，多个包同时调用会互相覆盖，
	// 导致后续 LoadLibrary 找不到本包的 DLL 而 panic。
	// 绝对路径让 LoadLibraryW 直接定位文件，不依赖搜索路径。
	dll = windows.NewLazyDLL(filePath)
}

func getProc(name string) *windows.LazyProc {
	if p, ok := procCache[name]; ok { return p }
	p := dll.NewProc(name)
	procCache[name] = p
	return p
}
`)
	return sb.String()
}

func generateDllBindingCodeForFile(funcs []DllFuncInfo, pkgName string, funcTypeNames map[string]bool, typedefs *safemap.M[string, TypedefInfo], enums *safemap.M[string, EnumInfo]) string {
	var sb strings.Builder

	receiverType := strings.ToUpper(pkgName[:1]) + pkgName[1:]
	receiver := strings.ToLower(receiverType[:1])
	usedNames := map[string]int{}
	seenFuncs := map[string]bool{}
	for _, fn := range funcs {
		if seenFuncs[fn.cName] {
			continue
		}
		seenFuncs[fn.cName] = true
		goName := strings.TrimPrefix(fn.goName, receiverType)
		goName = strings.TrimLeft(goName, "_")
		goName = GoDllIdent(goName)
		parts := strings.Split(goName, "_")
		for i, p := range parts {
			if len(p) > 0 {
				parts[i] = strings.ToUpper(p[:1]) + p[1:]
			}
		}
		goName = strings.Join(parts, "")
		if count, exists := usedNames[goName]; exists {
			usedNames[goName] = count + 1
			goName = fmt.Sprintf("%s_%d", goName, count)
		} else {
			usedNames[goName] = 1
		}
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
			syscallArgs = append(syscallArgs, DllParamToSyscall(p, funcTypeNames, typedefs, enums))
		}
		argsStr := strings.Join(syscallArgs, ", ")
		if fn.returnType == "void" || fn.returnType == "" {
			sb.WriteString(fmt.Sprintf("\tgetProc(%q).Call(%s)\n", fn.cName, argsStr))
		} else if fn.returnType == "bool" {
			sb.WriteString(fmt.Sprintf("\tr1, _, _ := getProc(%q).Call(%s)\n", fn.cName, argsStr))
			sb.WriteString("\treturn r1 != 0\n")
		} else {
			sb.WriteString(fmt.Sprintf("\tr1, _, _ := getProc(%q).Call(%s)\n", fn.cName, argsStr))
			resolvedRet := ResolveTypedef(fn.returnType, typedefs)
			cast := DllReturnTypeCast(resolvedRet)
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

func GoDllIdent(name string) string {
	if idx := strings.LastIndex(name, "@"); idx >= 0 {
		return name[:idx]
	}
	return name
}

func DllParamToSyscall(p DllFuncParam, funcTypeNames map[string]bool, typedefs *safemap.M[string, TypedefInfo], enums *safemap.M[string, EnumInfo]) string {
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
		return "func() uintptr { if " + p.goName + " { return 1 }; return 0 }()"
	case "unsafe.Pointer":
		return "uintptr(" + p.goName + ")"
	default:
		if strings.HasPrefix(resolvedType, "func(") || funcTypeNames[p.goType] {
			return "func() uintptr { if " + p.goName + " == nil { println(\"Callback is nil\"); return 0 }; return windows.NewCallback(" + p.goName + ") }()"
		}
		if td, ok := typedefs.Get(resolvedType); ok && td.isFunc {
			return "func() uintptr { if " + p.goName + " == nil { println(\"Callback is nil\"); return 0 }; return windows.NewCallback(" + p.goName + ") }()"
		}
		if td, ok := typedefs.Get(p.goType); ok && td.isFunc {
			return "func() uintptr { if " + p.goName + " == nil { println(\"Callback is nil\"); return 0 }; return windows.NewCallback(" + p.goName + ") }()"
		}
		if strings.HasPrefix(p.goType, "*") || p.isPointer {
			return "uintptr(unsafe.Pointer(" + p.goName + "))"
		}
		if IsGoPrimitiveType(resolvedType) {
			return "uintptr(" + p.goName + ")"
		}
		if enums.Has(resolvedType) || enums.Has(p.goType) {
			return "uintptr(" + p.goName + ")"
		}
		return "uintptr(unsafe.Pointer(&" + p.goName + "))"
	}
}

func ResolveTypedef(goType string, typedefs *safemap.M[string, TypedefInfo]) string {
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

func DllReturnTypeCast(goType string) string {
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
