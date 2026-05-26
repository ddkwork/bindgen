package cpp2c

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/ddkwork/golibrary/std/stream"
)

type EmitConfig struct {
	HeaderVersionMacro string
	AllowedEnumPrefix  []string
	ValueTypes         map[string]string
	FreeFuncCabiPrefix string
	SkipEnumGeneration bool
	SkipLibmiqtInclude bool
}

type ProjectConfig struct {
	Name               string
	SourceDir          string
	OutputCppDir       string
	OutputCabiDir      string
	Cflags             string
	Headers            []HeaderConfig
	CacheDir           string
	GenerateCMake      func(outDir, packageName string)
	GenerateAuxiliary  func(outDir string)
	AllowedClassPrefix []string
	Emit               EmitConfig
}

type HeaderConfig struct {
	Path      string
	AllowAll  bool
	AllowFunc func(string) bool
}

type QtModuleConfig struct {
	ModuleName   string
	UmbrellaPath string
}

func ProcessProjectHeaders(project *ProjectConfig) {
	ClearFilteredCache(CacheDir)
	ActiveAllowedClassPrefix = project.AllowedClassPrefix

	Generate(
		project.Name,
		func() []string {
			var headers []string
			for _, h := range project.Headers {
				headers = append(headers, h.Path)
			}
			return headers
		}(),
		func(string) bool { return true },
		project.Cflags,
		project.OutputCabiDir,
		func(string, string) bool { return true },
		project.Emit,
	)

	if project.GenerateCMake != nil {
		project.GenerateCMake(project.OutputCabiDir, project.Name)
	}
}

const (
	MaxClangSubprocessCount = 16
	BaseModule              = "github.com/ddkwork/miqt"
)

var GenerateCleaned = false

var CacheDir string

func SplitCflags(s string) []string {
	var ret []string
	var current strings.Builder
	inQuote := false
	for _, r := range s {
		if r == '"' {
			inQuote = !inQuote
		} else if r == ' ' && !inQuote {
			if current.Len() > 0 {
				ret = append(ret, current.String())
				current.Reset()
			}
		} else {
			current.WriteRune(r)
		}
	}
	if current.Len() > 0 {
		ret = append(ret, current.String())
	}
	return ret
}

func CacheFileRoot(inputHeader string) string {
	moduleName := strings.TrimSuffix(filepath.Base(inputHeader), filepath.Ext(inputHeader))
	return filepath.Join(CacheDir, moduleName)
}

func ImportPathForQtPackage(packageName string) string {
	return BaseModule + "/" + packageName
}

func FindHeadersInDir(srcDir string, allowHeader func(string) bool) []string {
	var ret []string
	filepath.Walk(srcDir, func(fullPath string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), `.h`) {
			return nil
		}
		if !allowHeader(fullPath) {
			return nil
		}
		ret = append(ret, fullPath)
		return nil
	})
	return ret
}

func CleanGeneratedFilesInDir(dirpath string) {
	_ = os.MkdirAll(dirpath, 0755)

	existing, err := os.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}

	cleaned := 0
	for _, e := range existing {
		if e.IsDir() {
			continue
		}
		if !strings.HasPrefix(e.Name(), `gen_`) {
			continue
		}
		err := os.Remove(filepath.Join(dirpath, e.Name()))
		if err != nil {
			continue
		}

		cleaned++
	}
}

func CollectGenCabiHeaders(dir string) []string {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	var headers []string
	for _, e := range entries {
		name := e.Name()
		if strings.HasPrefix(name, "gen_cabi_") && strings.HasSuffix(name, ".h") {
			headers = append(headers, name)
		}
	}
	return headers
}

func PkgConfigCflags(packageName string) string {
	stdout, err := exec.Command(`pkg-config`, `--cflags`, packageName).Output()
	if err != nil {
		panic(err)
	}

	return string(stdout)
}

func ParseHeaders(includeFiles []string, cflags []string, matcher HeaderMatcher) []*CppParsedHeader {
	result := make([]*CppParsedHeader, len(includeFiles))

	var wg sync.WaitGroup
	ch := make(chan struct{}, min(runtime.NumCPU(), MaxClangSubprocessCount))

	for i, includeFile := range includeFiles {
		ch <- struct{}{}
		wg.Add(1)

		go func(i int, includeFile string) {
			defer func() {
				wg.Done()
				<-ch
			}()

			result[i] = &CppParsedHeader{Filename: includeFile}
			ast := getFilteredAst(includeFile, cflags)
			if ast == nil {
				return
			}
			parseHeader(ast, "", result[i], matcher)
		}(i, includeFile)
	}
	wg.Wait()
	return result
}

func Generate(packageName string, srcDirs []string, allowHeaderFn func(string) bool, cflagsCombined, outDir string, matcher HeaderMatcher, cfg EmitConfig) {
	var includeFiles []string
	for _, srcDir := range srcDirs {
		if fi, err := os.Stat(srcDir); err == nil && !fi.IsDir() {
			includeFiles = append(includeFiles, srcDir)
		} else {
			includeFiles = append(includeFiles, FindHeadersInDir(srcDir, allowHeaderFn)...)
		}
	}

	cflags := SplitCflags(cflagsCombined)

	if !GenerateCleaned {
		CleanGeneratedFilesInDir(outDir)
		GenerateCleaned = true
	}

	FlushKnownTypes()
	ActiveFilter.InsertTypedefs()

	atr := astTransformRedundant{
		preserve: make(map[string]*CppEnum),
	}

	processHeaders := ParseHeaders(includeFiles, cflags, matcher)

	for _, parsed := range processHeaders {
		AstTransformChildClasses(parsed)
		astTransformApplyQuirks(packageName, parsed)
		astTransformOptional(parsed)
		astTransformOverloads(parsed)
		astTransformConstructorOrder(packageName, parsed)
		atr.Process(parsed)

		AddKnownTypes(packageName, parsed)
	}

	emptyCount := 0
	generatedCount := 0

	for _, parsed := range processHeaders {

		AstTransformTypedefs(parsed)
		astTransformBlocklist(parsed)

		if parsed.Empty() {
			emptyCount++
			continue
		}

		outputName := filepath.Join(outDir, "gen_"+strings.TrimSuffix(filepath.Base(parsed.Filename), `.h`))

		bindingCppSrc, err := emitBindingCpp(parsed, filepath.Base(parsed.Filename), cfg)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(outputName+".cpp", []byte(bindingCppSrc), 0644)
		if err != nil {
			panic(err)
		}

		bindingHSrc, cabiHSrc, err := emitBindingHeader(parsed, filepath.Base(parsed.Filename), packageName, cfg)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(outputName+".h", []byte(bindingHSrc), 0644)
		if err != nil {
			panic(err)
		}

		cabiOutputName := filepath.Join(outDir, "gen_cabi_"+strings.TrimSuffix(filepath.Base(parsed.Filename), `.h`))

		err = os.WriteFile(cabiOutputName+".h", []byte(cabiHSrc), 0644)
		if err != nil {
			panic(err)
		}

		generatedCount++

	}
}

func GenerateCommonFiles(outDir string) {
	os.WriteFile(filepath.Join(outDir, "libmiqt.h"), []byte(`#pragma once
#ifndef MIQT_LIBMIQT_LIBMIQT_H
#define MIQT_LIBMIQT_LIBMIQT_H

#ifdef __cplusplus
#include <cstddef>
#else
#include <stddef.h>
#endif

#ifdef __cplusplus
extern "C" {
#endif

struct miqt_string {
	size_t len;
	char* data;
};

struct miqt_array {
	size_t len;
	void* data;
};

struct miqt_map {
	size_t len;
	void* keys;
	void* values;
};

typedef const char const_char;

#ifdef __cplusplus
}
#endif

#endif
`), 0644)

	os.WriteFile(filepath.Join(outDir, "miqt_export.h"), []byte(`#pragma once
#ifndef MIQT_EXPORT_H
#define MIQT_EXPORT_H

#ifdef _WIN32
	#ifdef MIQT_BUILDING_DLL
		#define MIQT_EXPORT __declspec(dllexport)
	#else
		#define MIQT_EXPORT __declspec(dllimport)
	#endif
	#define MIQT_CALL __cdecl
#else
	#define MIQT_EXPORT __attribute__((visibility("default")))
	#define MIQT_CALL
#endif

#endif // MIQT_EXPORT_H
`), 0644)

	os.WriteFile(filepath.Join(outDir, "gen_callback_stubs.cpp"), []byte(GenerateCallbackStubs(outDir)), 0644)
}

func GenerateBuildBat(outDir string) {
	batPath := filepath.Join(outDir, "build.bat")
	batContent := `@echo off
cd /d "%~dp0"
if exist build rmdir /s /q build
set "PATH=C:\Program Files\CMake\bin;%PATH%"
cmake -B build -G "Ninja" -DCMAKE_BUILD_TYPE=Release . 2>&1
cmake --build build --config Release 2>&1 | powershell -Command "$input | Tee-Object -FilePath build\error.log"
`
	if err := os.WriteFile(batPath, []byte(batContent), 0644); err != nil {
	}
}

func GenerateCallbackStubs(outDir string) string {
	matches, _ := filepath.Glob(filepath.Join(outDir, "gen_*.cpp"))
	type stubInfo struct {
		fullDecl string
		retType  string
	}
	seen := make(map[string]stubInfo)
	for _, f := range matches {
		data, err := os.ReadFile(f)
		if err != nil {
			continue
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			trimmed := strings.TrimSpace(line)
			trimmed = strings.TrimSuffix(trimmed, "\r")
			if !strings.Contains(trimmed, "miqt_exec_callback_") {
				continue
			}
			if !strings.HasSuffix(trimmed, ";") {
				continue
			}
			decl := strings.TrimSuffix(trimmed, ";")
			parts := strings.SplitN(decl, "(", 2)
			if len(parts) != 2 {
				continue
			}
			retTypeAndName := strings.TrimSpace(parts[0])
			nameIdx := strings.LastIndex(retTypeAndName, " ")
			if nameIdx < 0 {
				continue
			}
			retType := retTypeAndName[:nameIdx]
			funcName := retTypeAndName[nameIdx+1:]
			if !strings.HasPrefix(funcName, "miqt_exec_callback_") {
				continue
			}
			if _, exists := seen[funcName]; !exists {
				seen[funcName] = stubInfo{
					fullDecl: decl,
					retType:  retType,
				}
			}
		}
	}

	var sb strings.Builder
	sb.WriteString("// Stub implementations for Go-side virtual method callbacks.\n")
	sb.WriteString("// These are provided by the Go runtime in normal builds.\n")
	sb.WriteString("// For standalone DLL compilation, these stubs allow linking to succeed.\n")
	sb.WriteString("// Auto-generated from gen_*.cpp callback declarations.\n\n")

	matches, _ = filepath.Glob(filepath.Join(outDir, "gen_*.h"))
	seenIncludes := make(map[string]bool)
	for _, f := range matches {
		base := filepath.Base(f)
		if strings.HasPrefix(base, "gen_cabi_") {
			continue
		}
		if !seenIncludes[base] {
			seenIncludes[base] = true
			sb.WriteString("#include \"" + base + "\"\n")
		}
	}
	sb.WriteString("#include \"libmiqt.h\"\n\n")
	sb.WriteString("extern \"C\" {\n\n")

	for _, info := range seen {
		skip := false
		for _, blocked := range []string{"QNativeInterface__", "QQmlSA__", "QtConcurrent__", "QDtlsClientVerifier__", "QQmlContext__PropertyPair", "QRangeModelImplBase_", "QAccessibleWidget_", "QAccessibleWidgetV2_"} {
			if strings.Contains(info.fullDecl, blocked) {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		retType := info.retType
		decl := info.fullDecl

		parenIdx := strings.Index(decl, "(")
		if parenIdx < 0 {
			writeStubBody(&sb, retType, decl)
			continue
		}
		retTypeAndName := decl[:parenIdx]
		paramsPart := decl[parenIdx+1 : len(decl)-1]
		nameIdx := strings.LastIndex(retTypeAndName, " ")
		if nameIdx < 0 {
			writeStubBody(&sb, retType, decl)
			continue
		}
		funcName := retTypeAndName[nameIdx+1:]

		paramTypes := splitCBParams(paramsPart)

		slotIdx := -1
		for i, pt := range paramTypes {
			trimmed := strings.TrimSpace(pt)
			if trimmed == "intptr_t" {
				slotIdx = i
				break
			}
		}

		if slotIdx < 0 {
			writeStubBody(&sb, retType, decl)
			continue
		}

		var namedParams []string
		var callArgTypes []string
		var callArgs []string

		for i, pt := range paramTypes {
			trimmed := strings.TrimSpace(pt)
			if i == slotIdx {
				namedParams = append(namedParams, trimmed+" _slot")
				continue
			}
			paramName := fmt.Sprintf("_p%d", i)
			namedParams = append(namedParams, trimmed+" "+paramName)
			callArgTypes = append(callArgTypes, trimmed)
			callArgs = append(callArgs, paramName)
		}

		newDecl := retType + " " + funcName + "(" + strings.Join(namedParams, ", ") + ")"

		fnPtrSig := strings.Join(callArgTypes, ", ")
		if fnPtrSig == "" {
			fnPtrSig = "void"
		}
		callStr := fmt.Sprintf("((%s (__stdcall *)(%s))_slot)(%s)", retType, fnPtrSig, strings.Join(callArgs, ", "))

		switch {
		case retType == "void":
			sb.WriteString(newDecl + " { if (_slot) " + callStr + "; }\n")
		case retType == "bool":
			sb.WriteString(newDecl + " { if (_slot) return " + callStr + "; return false; }\n")
		case retType == "int":
			sb.WriteString(newDecl + " { if (_slot) return " + callStr + "; return 0; }\n")
		case retType == "double":
			sb.WriteString(newDecl + " { if (_slot) return " + callStr + "; return 0.0; }\n")
		case strings.HasSuffix(retType, "*"):
			sb.WriteString(newDecl + " { if (_slot) return " + callStr + "; return nullptr; }\n")
		default:
			sb.WriteString(newDecl + " { if (_slot) return " + callStr + "; return {}; }\n")
		}
	}

	sb.WriteString("\n} // extern \"C\"\n")

	return sb.String()
}

func splitCBParams(s string) (result []string) {
	depth := 0
	start := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			depth++
		case ')':
			depth--
		case ',':
			if depth == 0 {
				result = append(result, strings.TrimSpace(s[start:i]))
				start = i + 1
			}
		}
	}
	if start <= len(s) {
		last := strings.TrimSpace(s[start:])
		if last != "" {
			result = append(result, last)
		}
	}
	return
}

func writeStubBody(sb *strings.Builder, retType, fullDecl string) {
	switch {
	case retType == "void":
		sb.WriteString(fullDecl + " {}\n")
	case retType == "bool":
		sb.WriteString(fullDecl + " { return false; }\n")
	case retType == "int":
		sb.WriteString(fullDecl + " { return 0; }\n")
	case retType == "double":
		sb.WriteString(fullDecl + " { return 0.0; }\n")
	case strings.HasSuffix(retType, "*"):
		sb.WriteString(fullDecl + " { return nullptr; }\n")
	default:
		sb.WriteString(fullDecl + " { return {}; }\n")
	}
}

func RunBuild(outDir string, dllName string) {
	println("Building...")
	println("ai不要动不动就停止控制台，最少等待30分钟，应该等控制台执行完毕")
	stream.RunCommandWithDir(outDir, "cmd","/c", "build.bat")
	//todo 增加构建时差打印
	dllPath := filepath.Join(outDir, "build", dllName)
	if _, err := os.Stat(dllPath); err != nil {
		println(fmt.Sprintf("DLL not found after build: %s", dllPath))
	errlogPath := filepath.Join(outDir, "build", "error.log")
		println("Error log path:", errlogPath)
		println("让ai阅读这个构建失败日志文件的所有内容修复cabi生成器的bug")
		panic("Build failed")
	}
	println("Build successful: " + dllPath)
}
