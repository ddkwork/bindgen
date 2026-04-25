package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"modernc.org/cc/v4"
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

func TestMultilineComment(t *testing.T) {
	testParse(t, "multiline comment", `
/*
 * Multi
 * line
 * comment
 */
int x;`)
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
