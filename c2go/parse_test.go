package c2go

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/ddkwork/bindgen/cc"
)

func newTestConfig(t testing.TB) *cc.Config {
	abi, err := cc.NewABI(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		t.Fatalf("cc.NewABI: %v", err)
	}
	return &cc.Config{
		ABI:        abi,
		CC:         findClExe(),
		Predefined: loadMSVCExtraTypes(),
	}
}

func testParse(t *testing.T, name string, code string) {
	cfg := newTestConfig(t)

	sources := []cc.Source{
		{Name: "<predefined>", Value: cfg.Predefined},
		{Name: "<builtin>", Value: cc.Builtin},
		{Name: "<msvc_types>", Value: loadMSVCTypes()},
		{Name: "<test>", Value: code},
	}
	_, err := cc.Translate(cfg, sources)
	if err != nil {
		t.Errorf("%s: FAIL: %v", name, err)
	} else {
		fmt.Printf("%s: PASS\n", name)
	}
}

func TestMSVCTypes(t *testing.T) { testParse(t, "MSVC types (unsigned long)", `UINT64 x;`) }
func TestInt64(t *testing.T)     { testParse(t, "__int64 typedef", `typedef unsigned __int64 UINT64;`) }
func TestInt64Multi(t *testing.T) {
	testParse(t, "__int64 multi declarator", `typedef unsigned __int64 UINT64, *PUINT64;`)
}
func TestBoolean(t *testing.T)  { testParse(t, "BOOLEAN typedef", `BOOLEAN flag;`) }
func TestHexArray(t *testing.T) { testParse(t, "Hex array", `UINT8 arr[0x0C];`) }
func TestDecimalArray(t *testing.T) {
	testParse(t, "Decimal array", `UINT8 arr[12];`)
}

func TestSizeofExpression(t *testing.T) {
	testParse(t, "sizeof expression macro", `#define SIZEOF_SYMBOL_WITHOUT_LEN sizeof(long long unsigned) * 2`)
}

func TestStructWithTypedef(t *testing.T) {
	testParse(t, "struct with typedef", `typedef struct _MODULE_SYMBOL_DETAIL {
    UINT8 IsSymbolDetailsFound;
    UINT64 BaseAddress;
} MODULE_SYMBOL_DETAIL, *PMODULE_SYMBOL_DETAIL;`)
}

func TestPointerToPointer(t *testing.T) {
	testParse(t, "pointer to pointer typedef", `typedef struct _TEST *PTEST;
typedef struct _TEST {
    int x;
} TEST, *PTEST;`)
}

func TestEnumComma(t *testing.T) {
	testParse(t, "enum with trailing comma", `typedef enum _TEST {
    A = 0,
    B,
    C,
} TEST;`)
}
func TestDeclspec(t *testing.T) { testParse(t, "__declspec", `__declspec(dllimport) void test();`) }
func TestFarCall(t *testing.T)  { testParse(t, "__cdecl", `void __cdecl test(void);`) }
func TestPackedStruct(t *testing.T) {
	testParse(t, "#pragma pack", `#pragma pack(push, 1)
typedef struct _PACKED {
    char a;
    long long b;
} PACKED;
#pragma pack(pop)`)
}

func TestPragmaWarning(t *testing.T) {
	testParse(t, "#pragma warning", `#pragma warning(disable: 4201)`)
}

func TestNakedFunction(t *testing.T) {
	testParse(t, "naked function", `__declspec(naked) void test(void) {}`)
}

func TestInline(t *testing.T) {
	testParse(t, "inline function", `static inline int test(void) { return 0; }`)
}

func TestStaticAssert(t *testing.T) {
	testParse(t, "static_assert", `typedef char static_assert_check[(sizeof(long long) == 8) ? 1 : -1];`)
}

func TestAlignas(t *testing.T) {
	testParse(t, "_Alignas", `typedef struct _ALIGNED {
    _Alignas(16) char x;
} ALIGNED;`)
}

func TestComplexMacro(t *testing.T) {
	testParse(t, "complex macro", `#define MAX_PATH 260
#define SIZEOF_DEBUGGER_READ_MEMORY sizeof(long long unsigned) * 3 + MAX_PATH`)
}

func TestLineContinuation(t *testing.T) {
	testParse(t, "line continuation", `#define TEST_MACRO \
    123`)
}

func TestAttribute(t *testing.T) {
	testParse(t, "__attribute__", `void test(void) __attribute__((noreturn));`)
}

func TestWchar_t(t *testing.T) {
	testParse(t, "wchar_t string", `typedef unsigned short wchar_t;
wchar_t *s = L"test";`)
}

func TestBitfield(t *testing.T) {
	testParse(t, "bitfield", `typedef struct _BITFIELD {
    unsigned int x:1;
    unsigned int y:2;
} BITFIELD;`)
}

func TestUnion(t *testing.T) {
	testParse(t, "union", `typedef union _UNION {
    int x;
    char y;
} UNION;`)
}

func TestCallbackType(t *testing.T) {
	testParse(t, "callback typedef", `typedef void (*Callback)(int x, char *y);`)
}

func TestGLFWAllocatefun(t *testing.T) {
	testParse(t, "GLFW allocate function pointer", `typedef void* (* GLFWallocatefun)(size_t size, void* user);`)
}

func TestParseStdintH(t *testing.T) {
	cfg := newTestConfig(t)
	cfg.PragmaHandler = func(toks []cc.Token) error { return nil }

	// Test 1: just try parsing stdint.h content directly
	t.Run("inline", func(t *testing.T) {
		content := `
#pragma once
#define _STDINT
#if 1
#pragma warning(push)
#pragma warning(disable: 4201)
typedef signed char        my_int8_t;
typedef short              my_int16_t;
typedef int                my_int32_t;
typedef long long          my_int64_t;
#endif
`
		sources := []cc.Source{
			{Name: "<predefined>", Value: cfg.Predefined},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<test>", Value: content},
		}
		_, err := cc.Translate(cfg, sources)
		if err != nil {
			t.Fatalf("inline FAIL: %v", err)
		}
	})

	// Test 2: parse actual stdint.h from MSVC include path (via system include)
	// This should work once __pragma is defined in loadMSVCExtraTypes()
	t.Run("system_stdint", func(t *testing.T) {
		// Use real MSVC sysinclude paths like Generate does
		realCfg := NewMSVCConfig(t)
		realCfg.PragmaHandler = func(toks []cc.Token) error { return nil }
		sources := []cc.Source{
			{Name: "<predefined>", Value: realCfg.Predefined},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<msvc_types>", Value: loadMSVCTypes()},
			{Name: "<test>", Value: "#include <stdint.h>\nuint64_t x;\n"},
		}
		_, err := cc.Translate(realCfg, sources)
		if err != nil {
			t.Fatalf("system_stdint FAIL: %v", err)
		}
		t.Log("system_stdint PASS")
	})

	// Test 3: full include chain simulating GLFW including stdint.h
	t.Run("include_vcruntime", func(t *testing.T) {
		realCfg := NewMSVCConfig(t)
		realCfg.PragmaHandler = func(toks []cc.Token) error { return nil }
		sources := []cc.Source{
			{Name: "<predefined>", Value: realCfg.Predefined},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<msvc_types>", Value: loadMSVCTypes()},
			{Name: "<test>", Value: "#include <vcruntime.h>\nint x;\n"},
		}
		_, err := cc.Translate(realCfg, sources)
		if err != nil {
			t.Fatalf("include_vcruntime FAIL: %v", err)
		}
		t.Log("include_vcruntime PASS")
	})
}

func TestMultilineComment(t *testing.T) {
	testParse(t, "multiline comment", `
/*
 * Multi
 * line
 * comment
 */
int x;`)
}

func TestZydisEncoderH(t *testing.T) {
	cfg := newTestConfig(t)
	// Test the exact macro chain from Zycore's Defines.h
	// With _MSC_VER >= 1800, ZYAN_STATIC_ASSERT(x) → static_assert(x, #x)
	sources := []cc.Source{
		{Name: "<predefined>", Value: cfg.Predefined},
		{Name: "<builtin>", Value: cc.Builtin},
		{Name: "<msvc_types>", Value: loadMSVCTypes()},
		{Name: "<test>", Value: `
#define ZYDIS_ENCODER_MAX_OPERANDS 5

// Exact Zycore Defines.h macro chain
#if defined(_MSC_VER) && (_MSC_VER >= 1800)
#   define ZYAN_STATIC_ASSERT(x) static_assert(x, #x)
#else
#   define ZYAN_STATIC_ASSERT(x) typedef int zyan_sassert[(x) ? 1 : -1]
#endif

#define ZYAN_NEEDS_BIT(n, b) (((unsigned long)(n) >> (b)) > 0)
#define ZYAN_BITS_TO_REPRESENT(n) \
    ( \
        ZYAN_NEEDS_BIT(n,  0) + ZYAN_NEEDS_BIT(n,  1) + \
        ZYAN_NEEDS_BIT(n,  2) + ZYAN_NEEDS_BIT(n,  3) + \
        ZYAN_NEEDS_BIT(n,  4) + ZYAN_NEEDS_BIT(n,  5) + \
        ZYAN_NEEDS_BIT(n,  6) + ZYAN_NEEDS_BIT(n,  7) + \
        ZYAN_NEEDS_BIT(n,  8) + ZYAN_NEEDS_BIT(n,  9) + \
        ZYAN_NEEDS_BIT(n, 10) + ZYAN_NEEDS_BIT(n, 11) + \
        ZYAN_NEEDS_BIT(n, 12) + ZYAN_NEEDS_BIT(n, 13) + \
        ZYAN_NEEDS_BIT(n, 14) + ZYAN_NEEDS_BIT(n, 15) + \
        ZYAN_NEEDS_BIT(n, 16) + ZYAN_NEEDS_BIT(n, 17) + \
        ZYAN_NEEDS_BIT(n, 18) + ZYAN_NEEDS_BIT(n, 19) + \
        ZYAN_NEEDS_BIT(n, 20) + ZYAN_NEEDS_BIT(n, 21) + \
        ZYAN_NEEDS_BIT(n, 22) + ZYAN_NEEDS_BIT(n, 23) + \
        ZYAN_NEEDS_BIT(n, 24) + ZYAN_NEEDS_BIT(n, 25) + \
        ZYAN_NEEDS_BIT(n, 26) + ZYAN_NEEDS_BIT(n, 27) + \
        ZYAN_NEEDS_BIT(n, 28) + ZYAN_NEEDS_BIT(n, 29) + \
        ZYAN_NEEDS_BIT(n, 30) + ZYAN_NEEDS_BIT(n, 31)   \
    )

ZYAN_STATIC_ASSERT(ZYAN_BITS_TO_REPRESENT(ZYDIS_ENCODER_MAX_OPERANDS) == 3);
`},
	}
	_, err := cc.Translate(cfg, sources)
	if err != nil {
		t.Fatalf("ZydisEncoderH: FAIL: %v", err)
	}
	t.Log("ZydisEncoderH: PASS")
}

func TestDocComment(t *testing.T) {
	testParse(t, "doc comment", `
/**
 * @brief Test function
 */
void test(void);`)
}

func TestMicrosoftMacro(t *testing.T) {
	testParse(t, "MS macro", `#define UNREFERENCED_PARAMETER(x) (x)`)
}

func TestMaxMacro(t *testing.T) {
	testParse(t, "max macro", `#define max(a,b) ((a) > (b) ? (a) : (b))`)
}

func TestForLoop(t *testing.T) {
	testParse(t, "for loop in function", `void test(void) { for (int i = 0; i < 10; i++) {} }`)
}
func TestStringLiteral(t *testing.T) { testParse(t, "string literal", `const char * s = "hello";`) }
func TestHexStringLiteral(t *testing.T) {
	testParse(t, "hex string literal", `const char * s = "\x1b[31m";`)
}

func TestPointerMath(t *testing.T) {
	testParse(t, "pointer math", `void test(void) { int arr[10]; int *p = arr; p = p + 1; }`)
}

func TestRealExample(t *testing.T) {
	testParse(t, "real example", `#define SIZEOF_DEBUGGER_READ_MEMORY sizeof(long long unsigned) * 3

typedef struct _DEBUGGER_READ_MEMORY {
    UINT64 Address;
    UINT32 Size;
    UINT32 Result;
} DEBUGGER_READ_MEMORY, *PDEBUGGER_READ_MEMORY;

typedef struct _MODULE_SYMBOL_DETAIL {
    BOOL IsSymbolDetailsFound;
    UINT64 BaseAddress;
    char FilePath[260];
} MODULE_SYMBOL_DETAIL, *PMODULE_SYMBOL_DETAIL;`)
}

func TestRunAllParse(t *testing.T) {
	fmt.Println("=== Running all parse tests ===")
	tests := []struct {
		name string
		code string
	}{
		{"MSVC types", `typedef unsigned __int64 UINT64; UINT64 x;`},
		{"BOOLEAN", `BOOLEAN flag;`},
		{"sizeof macro", `#define SIZEOF_X sizeof(int)`},
		{"hex array", `UINT8 arr[0x0C];`},
		{"__declspec", `__declspec(dllimport) void test();`},
		{"__cdecl", `void __cdecl test(void);`},
		{"#pragma pack", `#pragma pack(push,1) struct _X { int a; }; #pragma pack(pop)`},
		{"wchar_t", `typedef unsigned short wchar_t; wchar_t *s = L"x";`},
		{"bitfield", `struct _X { unsigned int x:1; };`},
		{"callback", `typedef void (*CB)(int);`},
		{"inline", `static inline int test(void) { return 0; }`},
		{"trailing comma enum", `enum _X { A, B, };`},
	}

	passed := 0
	failed := 0
	for _, tc := range tests {
		cfg := newTestConfig(t)

		_, err := cc.Translate(
			cfg,
			[]cc.Source{
				{Name: "<predefined>", Value: cfg.Predefined},
				{Name: "<builtin>", Value: cc.Builtin},
				{Name: "<msvc_types>", Value: loadMSVCTypes()},
				{Name: "<test>", Value: tc.code},
			},
		)
		if err != nil {
			fmt.Printf("FAIL: %-30s %v\n", tc.name, err)
			failed++
		} else {
			fmt.Printf("PASS: %s\n", tc.name)
			passed++
		}
	}
	fmt.Printf("\n=== Results: %d passed, %d failed ===\n", passed, failed)
	if failed > 0 {
		t.Fail()
	}
}

func TestGenerateDiag(t *testing.T) {
	cfg := newTestConfig(t)

	sdkHeadersDir := filepath.Join("SDK", "headers")
	files, _ := os.ReadDir(sdkHeadersDir)

	var sources []cc.Source
	sources = append(sources, cc.Source{Name: "<predefined>", Value: cfg.Predefined})
	sources = append(sources, cc.Source{Name: "<builtin>", Value: cc.Builtin})
	sources = append(sources, cc.Source{Name: "<msvc_types>", Value: loadMSVCTypes()})

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".h") {
			continue
		}
		fp := filepath.Join(sdkHeadersDir, file.Name())
		content, _ := os.ReadFile(fp)
		sources = append(sources, cc.Source{Name: file.Name(), Value: string(content)})
	}

	ast, err := cc.Translate(cfg, sources)
	t.Logf("err=%v ast=nil?%v len(sources)=%d", err, ast == nil, len(sources))
	if ast != nil {
		count := 0
		for l := ast.TranslationUnit; l != nil; l = l.TranslationUnit {
			count++
		}
		t.Logf("AST nodes: %d", count)
	}
}

func TestGenerateDiag2(t *testing.T) {
	cfg := newTestConfig(t)
	predef := cfg.Predefined

	tests := []string{
		"BasicTypes.h",
	}
	for i := range tests {
		sources := []cc.Source{
			{Name: "<predefined>", Value: predef},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<msvc_types>", Value: loadMSVCTypes()},
		}
		for j := 0; j <= i; j++ {
			content, _ := os.ReadFile(filepath.Join("SDK", "headers", tests[j]))
			sources = append(sources, cc.Source{Name: tests[j], Value: string(content)})
		}

		t.Logf("=== Testing with MSVC config (predefined=%d bytes) ===", len(predef))
		ast, err := cc.Translate(cfg, sources)
		t.Logf("MSVC files=%v err=%v ast=nil?%v", tests[:i+1], err, ast == nil)

		cfg2 := newTestConfig(t)
		ast2, err2 := cc.Translate(cfg2, []cc.Source{
			{Name: "<predefined>", Value: cfg2.Predefined},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<msvc_types>", Value: loadMSVCTypes()},
			{Name: "<test>", Value: `typedef unsigned __int64   ULONG64, *PULONG64;`},
		})
		t.Logf("WITH_EXTRA err=%v ast=nil?%v", err2, ast2 == nil)

		cfg3 := newTestConfig(t)
		ast3, err3 := cc.Translate(cfg3, []cc.Source{
			{Name: "<predefined>", Value: cfg3.Predefined},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<msvc_types>", Value: loadMSVCTypes()},
			{Name: "<test>", Value: `typedef unsigned __int64   ULONG64, *PULONG64;`},
		})
		t.Logf("MSVC_FULL err=%v ast=nil?%v", err3, ast3 == nil)

		basicContent, errBasic := os.ReadFile(filepath.Join("SDK", "headers", "BasicTypes.h"))
		if errBasic != nil {
			t.Logf("BasicTypes.h not found, skipping partial parse test: %v", errBasic)
			return
		}
		lines := strings.Split(string(basicContent), "\n")
		for _, endLine := range []int{26, 27, 28, 29, 30} {
			partial := strings.Join(lines[:endLine], "\n")
			ast4, err4 := cc.Translate(cfg, []cc.Source{
				{Name: "<predefined>", Value: predef},
				{Name: "<builtin>", Value: cc.Builtin},
				{Name: "<msvc_types>", Value: loadMSVCTypes()},
				{Name: "<basic>", Value: partial},
			})
			t.Logf("BASIC_LINES_%d err=%v nil?%v", endLine, err4, ast4 == nil)
		}

		t.Logf("--- Test skipping line 31 ---")
		skip31 := append(lines[:30], lines[31:]...)
		ast5, err5 := cc.Translate(cfg, []cc.Source{
			{Name: "<predefined>", Value: predef},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<msvc_types>", Value: loadMSVCTypes()},
			{Name: "<basic>", Value: strings.Join(skip31, "\n")},
		})
		t.Logf("SKIP_31 err=%v nil?%v", err5, ast5 == nil)

		t.Logf("--- Minimal repro ---")
		minimal := `typedef unsigned __int64   UINT64, *PUINT64;
typedef unsigned int       UINT;
typedef unsigned __int64   ULONG64, *PULONG64;`
		ast6, err6 := cc.Translate(cfg, []cc.Source{
			{Name: "<predefined>", Value: predef},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<msvc_types>", Value: loadMSVCTypes()},
			{Name: "<test>", Value: minimal},
		})
		t.Logf("MINIMAL err=%v nil?%v", err6, ast6 == nil)
	}
}

func TestDiagPredef(t *testing.T) {
	cfg2 := newTestConfig(t)
	predef2 := cfg2.Predefined

	tests := []struct {
		name string
		code string
	}{
		{"just WCHAR", "typedef wchar_t WCHAR;"},
		{"WCHAR with predef", "#define __predefined_wchar_t int\ntypedef wchar_t WCHAR;"},
		{"WCHAR with typedef first", "typedef int __predefined_wchar_t;\ntypedef wchar_t WCHAR;"},
	}

	for _, tt := range tests {
		ast, err := cc.Translate(cfg2, []cc.Source{
			{Name: "<predefined>", Value: predef2},
			{Name: "<builtin>", Value: cc.Builtin},
			{Name: "<msvc_types>", Value: loadMSVCTypes()},
			{Name: "<test>", Value: tt.code},
		})
		status := "OK"
		if err != nil {
			status = err.Error()
		}
		t.Logf("%s: %s (nil=%v)", tt.name, status, ast == nil)
		_ = ast
	}
}

func loadStringH() string {
	return `
#ifndef _SIZE_T_DEFINED
typedef unsigned long long size_t;
#define _SIZE_T_DEFINED
#endif
#ifndef NULL
#define NULL ((void*)0)
#endif
void* memcpy(void* dest, const void* src, size_t n);
size_t strlen(const char* s);
char* strdup(const char* s);
`
}

func TestQt6HeaderParsing(t *testing.T) {
	cfg := newTestConfig(t)

	tests := []struct {
		name string
		code string
	}{
		{
			name: "ifdef_cplusplus_class_fwd",
			code: `
#ifdef __cplusplus
class QApplication;
class QEvent;
class QMetaObject;
class QObject;
class QStyle;
class QWidget;
#else
typedef struct QApplication QApplication;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QStyle QStyle;
typedef struct QWidget QWidget;
#endif

MIQT_EXPORT QApplication* QApplication_new(int* argc, char** argv);
`,
		},
		{
			name: "ifdef_cplusplus_with_using",
			code: `
#ifdef __cplusplus
class QAbstractEventDispatcher;
class QCoreApplication;
using Qt_ApplicationAttribute = Qt::ApplicationAttribute;
#else
typedef struct QAbstractEventDispatcher QAbstractEventDispatcher;
typedef struct QCoreApplication QCoreApplication;
typedef int Qt_ApplicationAttribute;
#endif

MIQT_EXPORT QCoreApplication* QCoreApplication_new(int* argc, char** argv);
`,
		},
		{
			name: "ifdef_cplusplus_with_typedef_inner",
			code: `
#ifdef __cplusplus
#if defined(WORKAROUND_INNER_CLASS_DEFINITION_QEventLoop__ProcessEventsFlags)
typedef QEventLoop::ProcessEventsFlags QEventLoop__ProcessEventsFlags;
#else
class QEventLoop__ProcessEventsFlags;
#endif
#else
typedef struct QEventLoop__ProcessEventsFlags QEventLoop__ProcessEventsFlags;
#endif

MIQT_EXPORT void test_func(QEventLoop__ProcessEventsFlags* flags);
`,
		},
		{
			name: "c_only_opaque_structs",
			code: `
typedef struct QApplication QApplication;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;

MIQT_EXPORT QApplication* QApplication_new(int* argc, char** argv);
`,
		},
		{
			name: "full_qapplication_header",
			code: `
#pragma once
#ifndef MIQT_QT6_GEN_QAPPLICATION_H
#define MIQT_QT6_GEN_QAPPLICATION_H

#pragma GCC diagnostic ignored "-Wdeprecated-declarations"

#include "libmiqt.h"
#include "miqt_export.h"

#ifdef __cplusplus
extern "C" {
#endif

#ifdef __cplusplus
class QApplication;
class QEvent;
class QMetaObject;
class QObject;
class QStyle;
class QWidget;
#else
typedef struct QApplication QApplication;
typedef struct QEvent QEvent;
typedef struct QMetaObject QMetaObject;
typedef struct QObject QObject;
typedef struct QStyle QStyle;
typedef struct QWidget QWidget;
#endif

MIQT_EXPORT QApplication*QApplication_new(int*argc, char** argv);
MIQT_EXPORT int QApplication_exec();

#ifdef __cplusplus
}
#endif

#endif
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sources := []cc.Source{
				{Name: "<predefined>", Value: cfg.Predefined},
				{Name: "<builtin>", Value: cc.Builtin},
				{Name: "<msvc_types>", Value: loadMSVCTypes()},
				{Name: "<test>", Value: tc.code},
			}
			ast, err := cc.Translate(cfg, sources)
			if err != nil {
				t.Errorf("PARSE ERROR: %v", err)
			} else if ast == nil {
				t.Error("AST is nil (partial/failed parse)")
			} else {
				count := 0
				for l := ast.TranslationUnit; l != nil; l = l.TranslationUnit {
					count++
				}
				t.Logf("OK: AST nodes=%d", count)
			}
			_ = ast
		})
	}
}

func TestLibmiqtHeaderParsing(t *testing.T) {
	cfg := newTestConfig(t)

	tests := []struct {
		name           string
		includeStringH bool
		code           string
	}{
		{
			name:           "libmiqt_with_string_h",
			includeStringH: true,
			code: `
#pragma once
#ifndef MIQT_LIBMIQT_H
#define MIQT_LIBMIQT_H

#include <string.h>

struct miqt_string {
	size_t len;
	char* data;
};

struct miqt_array {
	size_t len;
	void* data;
};

#endif
`,
		},
		{
			name:           "libmiqt_without_string_h",
			includeStringH: false,
			code: `
#pragma once
#ifndef MIQT_LIBMIQT_H
#define MIQT_LIBMIQT_H

struct miqt_string {
	size_t len;
	char* data;
};

struct miqt_array {
	size_t len;
	void* data;
};

#endif
`,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sources := []cc.Source{
				{Name: "<predefined>", Value: cfg.Predefined},
				{Name: "<builtin>", Value: cc.Builtin},
				{Name: "<msvc_types>", Value: loadMSVCTypes()},
			}
			if tc.includeStringH {
				sources = append(sources, cc.Source{Name: "<string.h>", Value: loadStringH()})
			}
			sources = append(sources, cc.Source{Name: "<test>", Value: tc.code})

			ast, err := cc.Translate(cfg, sources)
			if err != nil {
				t.Errorf("PARSE ERROR: %v", err)
			} else if ast == nil {
				t.Error("AST is nil (partial/failed parse)")
			} else {
				count := 0
				for l := ast.TranslationUnit; l != nil; l = l.TranslationUnit {
					count++
				}
				t.Logf("OK: AST nodes=%d", count)
			}
			_ = ast
		})
	}
}

func TestGLFW3Header(t *testing.T) {
	glfwPath := filepath.Join("..", "project", "glfw", "clone", "glfw", "include", "GLFW", "glfw3.h")
	content, err := os.ReadFile(glfwPath)
	if err != nil {
		t.Skipf("glfw3.h not found at %s: %v", glfwPath, err)
	}

	cfg := NewMSVCConfig(t)
	cfg.PragmaHandler = func(toks []cc.Token) error { return nil }
	cfg.Predefined += `
#define GLFW_INCLUDE_NONE
#define GLFW_DLL
`

	fullSources := []cc.Source{
		{Name: "<predefined>", Value: cfg.Predefined},
		{Name: "<builtin>", Value: cc.Builtin},
		{Name: "<msvc_types>", Value: loadMSVCTypes()},
		{Name: "GLFW/glfw3.h", Value: string(content)},
	}

	_, err = cc.Translate(cfg, fullSources)
	if err != nil {
		t.Fatalf("glfw3.h parse failed: %v", err)
	}
}

func TestChineseCharacterPanic(t *testing.T) {
	tests := []struct {
		name    string
		source  string
		wantErr string
	}{
		{
			name:    "chinese_before_define",
			source:  "你\n#define FOO 1\nint x;\n",
			wantErr: "non-ASCII",
		},
		{
			name:    "chinese_in_function",
			source:  "int foo() { return 好; }\n",
			wantErr: "non-ASCII",
		},
		{
			name:    "chinese_in_identifier_middle",
			source:  "int fo好o;\n",
			wantErr: "non-ASCII",
		},
		{
			name:    "chinese_after_typedef",
			source:  "typedef int 类型;\n",
			wantErr: "non-ASCII",
		},
	}

	abi, err := cc.NewABI(runtime.GOOS, runtime.GOARCH)
	if err != nil {
		t.Fatal(err)
	}
	cfg := &cc.Config{
		ABI: abi,
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sources := []cc.Source{
				{Name: "<test>", Value: tc.source},
			}
			_, err := cc.Translate(cfg, sources)
			if err == nil {
				t.Fatal("expected error for Chinese character, got nil")
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("error %q does not contain %q", err.Error(), tc.wantErr)
			}
			t.Logf("OK: got expected error: %v", err)
		})
	}
}
