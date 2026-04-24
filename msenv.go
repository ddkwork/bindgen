package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func findClExe() string { return os.Getenv("CC") }

func findEWDKIncludePaths() []string {
	var rawPaths []string
	rawPaths = append(rawPaths, "SDK/headers")
	for _, p := range strings.Split(os.Getenv("INCLUDE"), ";") {
		if p = strings.TrimSpace(p); p != "" {
			rawPaths = append(rawPaths, p)
		}
	}
	var crtPaths, otherPaths, msvcPaths []string
	for _, p := range rawPaths {
		clean := filepath.Clean(p)
		if strings.Contains(clean, `km\crt`) || strings.Contains(clean, `km/crt`) {
			crtPaths = append(crtPaths, clean)
		} else if strings.Contains(clean, `ucrt`) && !strings.Contains(clean, `VC\Tools\MSVC`) {
			continue
		} else if strings.Contains(clean, `VC\Tools\MSVC`) || strings.Contains(clean, `VC/Tools/MSVC`) {
			msvcPaths = append(msvcPaths, clean)
		} else {
			otherPaths = append(otherPaths, clean)
		}
	}
	var paths []string
	paths = append(paths, crtPaths...)
	paths = append(paths, otherPaths...)
	paths = append(paths, msvcPaths...)
	var clean []string
	seen := make(map[string]bool)
	for _, p := range paths {
		if p == "" || seen[p] {
			continue
		}
		info, err := os.Stat(p)
		if err != nil || !info.IsDir() {
			continue
		}
		seen[p] = true
		clean = append(clean, p)
	}
	fmt.Printf("msenv: loaded %d include paths\n", len(clean))
	return clean
}

func loadMSVCExtraTypes() string {
	return `
#define _MSC_VER 1941
#define _MSC_FULL_VER 1941341200
#define _WIN64 1
#define _WIN32 1
#define _M_X64 100
#define _M_AMD64 100
#define _INTEGRAL_MAX_BITS 64

#define __SIZE_TYPE__ unsigned long long
#define __WCHAR_TYPE__ unsigned short
#define __PTRDIFF_TYPE__ long long
#define __INT8_TYPE__ signed char
#define __INT16_TYPE__ short
#define __INT32_TYPE__ int
#define __INT64_TYPE__ long long
#define __UINT8_TYPE__ unsigned char
#define __UINT16_TYPE__ unsigned short
#define __UINT32_TYPE__ unsigned int
#define __UINT64_TYPE__ unsigned long long

#define __declspec(x)
#define __cdecl
#define __stdcall
#define __fastcall
#define __thiscall
#define __inline inline
#define __forceinline
#define __unaligned
#define __w64
#define __ptr64
#define __noop ((void)0)

#define _alignof(x) __alignof(x)
#define __alignof __alignof
#define __uuidof(x) 0

#define NULL ((void*)0)

#define VOID void
#define CALLBACK
#define WINAPI
#define APIENTRY
#define CDECL
#define PASCAL
#define STDMETHODCALLTYPE
#define STDMETHODVCALLTYPE(x)

#define MAX_PATH 260
#define IN
#define OUT
#define OPTIONAL
#define UNREFERENCED_PARAMETER(x) (x)
#define DBGUNREFERENCED_LOCAL_VARIABLE(x) (x)
#define ASSERT(x)
#define DEBUG_BREAK()

#pragma warning(disable: 4201)
#pragma warning(disable: 4214)
#pragma warning(disable: 4820)

#define _VCRUNTIME_H
#define _INC_CRTDEFS
#define _INC__VADEFS
#define _SAL_H_DEFINED
#define _INC_CORECRT
#define _VCRT_COMPILER_PREPROCESSOR 1
#define _UCRT
#define _CRT_PACKING 8
#define _CRTIMP
#define _CRT_BEGIN_C_HEADER
#define _CRT_END_C_HEADER

#define _In_
#define _In_opt_
#define _In_z_
#define _In_opt_z_
#define _In_reads_(s)
#define _In_reads_opt_(s)
#define _In_reads_bytes_(s)
#define _In_reads_bytes_opt_(s)
#define _Out_
#define _Out_opt_
#define _Out_writes_(s)
#define _Out_writes_opt_(s)
#define _Out_writes_bytes_(s)
#define _Out_writes_bytes_opt_(s)
#define _Out_writes_z_(s)
#define _Out_writes_opt_z_(s)
#define _Out_writes_to_(s,c)
#define _Out_writes_to_opt_(s,c)
#define _Out_writes_all_(s)
#define _Out_writes_all_opt_(s)
#define _Outptr_
#define _Outptr_opt_
#define _Inout_
#define _Inout_opt_
#define _Inout_writes_(s)
#define _Inout_writes_opt_(s)
#define _Inout_writes_bytes_(s)
#define _Inout_writes_bytes_opt_(s)
#define _Inout_writes_z_(s)
#define _Inout_writes_opt_z_(s)
#define _Ret_maybenull_
#define _Ret_notnull_
#define _Ret_null_
#define _Success_(x)
#define _Use_decl_annotations_
#define _Post_writable_byte_size_(s)
#define _Pre_notnull_
#define _Post_satisfies_(x)
#define _IRQL_requires_min_(x)
#define _IRQL_requires_max_(x)
#define _IRQL_requires_(x)
#define _At_(x,y)
#define _When_(x,y)
#define _Acquires_lock_(x)
#define _Releases_lock_(x)
#define __analysis_assume(x)
#define __in
#define __in_opt
#define __out
#define __out_opt
#define __inout
#define __inout_opt
#define __in_ecount(s)
#define __in_ecount_opt(s)
#define __out_ecount(s)
#define __out_ecount_opt(s)
#define __in_bcount(s)
#define __in_bcount_opt(s)
#define __out_bcount(s)
#define __out_bcount_opt(s)
#define __inout_ecount(s)
#define __inout_ecount_opt(s)
#define __deref_out_ecount(s)
#define __deref_out_ecount_opt(s)
#define __out_ecount_part(s,c)
#define __field_ecount(s)
#define __field_bcount(s)
#define __struct_bcount(s)
#define _Writable_bytes_(s)
#define _Inexpressible_(s)
#define _Field_range_(a,b)
#define _Post_equal_to_(x)
#define _Post_satisfies_(x)
#define _UCRT_DISABLE_CLANG_WARNINGS
#define _CRT_DEPRECATE_TEXT(x)
#define __int64 long long
#define __int32 int
#define __int16 short
#define __int8 char
#define __cdecl
#define _CRT_ALTERNATIVE_INLINES
#define _CRTIMP2
#define _CRTIMP_ALT
#define _MRTIMP
#define _MRTIMP2
#define _MCRTIMP
#define _CRTIMP_PURE
#define _CRTEXP
#define _CRTEXP_PURE
#define _PGLOBAL
#define _AGLOBAL
#define _CRTDATA(x) x
#define _CRTIMP2_FUNCTION(type) type __cdecl
#define _CRTIMP2_PURE_FUNCTION(type) type __cdecl
#define _CRTIMP2_MEMBER_FUNCTION(type) type __cdecl
#define _CRTIMP2_PURE_MEMBER_FUNCTION(type) type __cdecl
#define _MRTIMP_FUNCTION(type) type __cdecl
#define _MRTIMP2_FUNCTION(type) type __cdecl
#define _MSVCR80_FUNCTION(type) type __cdecl
#define _MSVCR80_FUNCTION_2(sal, type) type __cdecl
#define _CRTIMP_TYPEINFO
#define _CRTIMP_PURE_TYPEINFO
#define __STDC_SECURE_LIB__ 200411L
#define __GOT_SECURE_LIB__ __STDC_SECURE_LIB__
#define __STDC_WANT_SECURE_LIB__ 1
#define _CRT_SECURE_NO_WARNINGS
#define _CRT_NOFORCE_MANIFEST
#define _W64
#define _CRT_PACKING 8
#define _CRT_INSECURE_DEPRECATE(x)
#define _CRT_INSECURE_DEPRECATE_GLOBALS(x)
#define _CRT_OBSOLETE(x)
#define _Deref_out_z_
#define _Deref_out_
#define _Deref_out_opt_z_
#define _Deref_out_opt_
#define _CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES 0
#define _CRT_SECURE_CPP_OVERLOAD_STANDARD_NAMES_COUNT 0
#define _CRT_SECURE_CPP_OVERLOAD_SECURE_NAMES 0
#define __CRTDECL __cdecl
#define _CRT_JIT_INTRINSIC
#define _CRT_MANAGED_HEAP_DEPRECATE
#define _CRT_CONST_CORRECT_OVERLOADS
#define _CRT_DECLARE_NONSTDC_NAMES(x) 1
#define _CRT_NONSTDC_DEPRECATE(x)
#define _CRT_NONSTDC_NO_WARNINGS
#define _CRT_SECURE_NO_DEPRECATE
#define _CRT_NONSTDC_NO_DEPRECATE
#define _DEPRECATE_SUPPORTED 1
#define __STDC_WANT_SECURE_LIB__ 1
#define _SECURECRT_FILL_BUFFER_PATTERN 0xFD
#define _CRT_SECURE_CPP_NOTHROW
#define _HAS_EXCEPTIONS 0
#define _locale_t void*
#define _LDOUBLE
#define _CRT_ATOF_DEFINED
#define _CRT_ERRNO_DEFINED
#define _CRT_SYS_ERR_DATA
#define _M_CEE_PURE
#define _INC_STDLIB
#define _INC_SEARCH
#define _INC_MALLOC
#define _INC_STDIO
#define _STDIO_DEFINED
#define _STDBOOL
#define __bool_true_false_are_defined 1
#define bool _Bool
#define false 0
#define true 1
#define _VCRT_COMPILER_PREPROCESSOR 1
#define _STDINT
#define _Check_return_
#define _Check_return_opt_
#define _Check_return_wat_
#define __ALTDECL __cdecl
#define FILE void
`
}

func loadMSVCTypes() string {
	return `
typedef __WCHAR_TYPE__ wchar_t;
typedef __SIZE_TYPE__  size_t;
typedef __SIZE_TYPE__  SIZE_T;
typedef __PTRDIFF_TYPE__ ptrdiff_t;
typedef __SIZE_TYPE__  uintptr_t;
typedef __PTRDIFF_TYPE__ intptr_t;
typedef __UINT8_TYPE__  uint8_t;
typedef __UINT16_TYPE__ uint16_t;
typedef __UINT32_TYPE__ uint32_t;
typedef __UINT64_TYPE__ uint64_t;
typedef __INT8_TYPE__   int8_t;
typedef __INT16_TYPE__  int16_t;
typedef __INT32_TYPE__  int32_t;
typedef __INT64_TYPE__  int64_t;
typedef long long      time_t;
typedef int errno_t;
typedef long __time32_t;
typedef long long __time64_t;
typedef char * va_list;
typedef struct _LIST_ENTRY {
    struct _LIST_ENTRY *Flink;
    struct _LIST_ENTRY *Blink;
} LIST_ENTRY, *PLIST_ENTRY;
typedef int                   BOOL;
typedef unsigned char         BOOLEAN;
typedef unsigned long         DWORD;
typedef unsigned short        WORD;
typedef void                 *PVOID;
typedef void                 *LPVOID;
typedef void                 *HANDLE;
typedef unsigned long long    QWORD;
typedef unsigned __int64      UINT64, *PUINT64;
typedef unsigned __int32      UINT32, *PUINT32;
typedef unsigned __int16      UINT16, *PUINT16;
typedef unsigned __int8       UINT8, *PUINT8;
typedef __int64               INT64, *PINT64;
typedef signed int            INT32, *PINT32;
typedef signed short          INT16, *PINT16;
typedef signed char           INT8, *PINT8;
typedef unsigned __int64      ULONG64, *PULONG64;
typedef unsigned __int64      DWORD64, *PDWORD64;
typedef char                  CHAR;
typedef wchar_t               WCHAR;
typedef unsigned char         UCHAR;
typedef short                 SHORT;
typedef unsigned short        USHORT;
typedef int                   INT;
typedef unsigned int          UINT;
typedef unsigned int         *PUINT;
typedef long                  LONG;
typedef unsigned long         ULONG;

#define VOID void
`
}
