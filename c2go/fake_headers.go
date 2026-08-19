package c2go

import (
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

type FakeHeaderFS struct {
	files map[string]string
}

func NewFakeHeaderFS() *FakeHeaderFS {
	return &FakeHeaderFS{files: map[string]string{
		"vadefs.h":     vadefsH(),
		"sal.h":        salH(),
		"vcruntime.h":  vcruntimeH(),
		"stddef.h":     stddefH(),
		"stdint.h":     stdintH(),
		"stdio.h":      stdioH(),
		"stdlib.h":     stdlibH(),
		"string.h":     stringH(),
		"stdbool.h":    stdboolH(),
		"windows.h":    windowsH(),
		"errno.h":      errnoH(),
		"inttypes.h":   inttypesH(),
		"limits.h":     limitsH(),
		"math.h":       mathH(),
		"stdarg.h":     stdargH(),
	}}
}

func (f *FakeHeaderFS) Open(name string) (fs.File, error) {
	name = filepath.ToSlash(name)
	name = strings.TrimPrefix(name, "./")
	content, ok := f.files[name]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return &fakeFile{name: name, content: content, pos: 0}, nil
}

type fakeFile struct {
	name    string
	content string
	pos     int64
}

func (f *fakeFile) Stat() (fs.FileInfo, error) {
	return &fakeFileInfo{name: f.name, size: int64(len(f.content))}, nil
}
func (f *fakeFile) Read(p []byte) (int, error) {
	if f.pos >= int64(len(f.content)) {
		return 0, io.EOF
	}
	n := copy(p, f.content[f.pos:])
	f.pos += int64(n)
	return n, nil
}
func (f *fakeFile) Close() error { return nil }

type fakeFileInfo struct {
	name string
	size int64
}

func (fi *fakeFileInfo) Name() string       { return fi.name }
func (fi *fakeFileInfo) Size() int64        { return fi.size }
func (fi *fakeFileInfo) Mode() fs.FileMode  { return 0444 }
func (fi *fakeFileInfo) ModTime() time.Time { return time.Time{} }
func (fi *fakeFileInfo) IsDir() bool        { return false }
func (fi *fakeFileInfo) Sys() any           { return nil }

func vadefsH() string {
	return `
#ifndef _INC__VADEFS
#define _INC__VADEFS

#pragma once

#define __annotation(x)
#define _In_
#define _In_opt_
#define _Out_
#define _Out_opt_

#endif
`
}

func salH() string {
	return `
#ifndef _SAL_H_DEFINED
#define _SAL_H_DEFINED

#include <vadefs.h>

#define _In_z_
#define _In_opt_z_
#define _In_reads_(s)
#define _In_reads_opt_(s)
#define _In_reads_bytes_(s)
#define _In_reads_bytes_opt_(s)
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
#define _Check_return_
#define _Check_return_opt_
#define _Check_return_wat_
#define _Deref_out_z_
#define _Deref_out_
#define _Deref_out_opt_z_
#define _Deref_out_opt_

#endif
`
}

func vcruntimeH() string {
	return `
#ifndef _VCRUNTIME_H
#define _VCRUNTIME_H

#pragma once

#define _VCRT_COMPILER_PREPROCESSOR 1
#define _UCRT

#ifndef _CRTIMP
    #define _CRTIMP
#endif

#include <sal.h>
#include <vadefs.h>

#pragma warning(push)
#pragma warning(disable: 4201)

#define _CRT_PACKING 8
#define _CRT_BEGIN_C_HEADER
#define _CRT_END_C_HEADER

#ifndef _HAS_EXCEPTIONS
    #define _HAS_EXCEPTIONS 0
#endif

#if !defined _M_CEE && !defined __midl
    void __cdecl __security_init_cookie(void);
    #if defined(_M_IX86)
        void __fastcall __security_check_cookie(unsigned long long _StackCookie);
        __declspec(noreturn) void __cdecl __report_gsfailure(void);
    #else
        void __cdecl __security_check_cookie(unsigned long long _StackCookie);
        __declspec(noreturn) void __cdecl __report_gsfailure(unsigned long long _StackCookie);
    #endif
#endif

extern unsigned long long __security_cookie;

#ifndef _VCRT_BUILD
    #define __vcrt_malloc_normal(_Size) malloc(_Size)
    #define __vcrt_calloc_normal(_Count, _Size) calloc(_Count, _Size)
    #define __vcrt_free_normal(_Memory) free(_Memory)
#endif

#pragma warning(pop)

#endif
`
}

func stddefH() string {
	return `
#ifndef _STDDEF_H_
#define _STDDEF_H_

#include <vcruntime.h>

typedef __SIZE_TYPE__ size_t;
typedef __PTRDIFF_TYPE__ ptrdiff_t;

#define NULL ((void*)0)
#define offsetof(s,m) ((size_t)&(((s*)0)->m))

#endif
`
}

func stdintH() string {
	return `
#ifndef _STDINT_H_
#define _STDINT_H_

#include <vcruntime.h>

typedef signed char        int8_t;
typedef short              int16_t;
typedef int                int32_t;
typedef long long          int64_t;
typedef unsigned char      uint8_t;
typedef unsigned short     uint16_t;
typedef unsigned int       uint32_t;
typedef unsigned long long uint64_t;

typedef long long          intptr_t;
typedef unsigned long long uintptr_t;

#endif
`
}

func stdlibH() string {
	return `
#ifndef _INC_STDLIB
#define _INC_STDLIB

#include <stddef.h>

void* malloc(size_t size);
void* calloc(size_t count, size_t size);
void* realloc(void* ptr, size_t size);
void free(void* ptr);
void exit(int code);
void abort(void);
int atexit(void (*func)(void));
int abs(int n);

#define EXIT_FAILURE 1
#define EXIT_SUCCESS 0
#define RAND_MAX 0x7FFF

#endif
`
}

func stringH() string {
	return `
#ifndef _INC_STRING
#define _INC_STRING

#include <stddef.h>

void* memcpy(void* dest, const void* src, size_t n);
void* memmove(void* dest, const void* src, size_t n);
char* strcpy(char* dest, const char* src);
char* strncpy(char* dest, const char* src, size_t n);
char* strcat(char* dest, const char* src);
size_t strlen(const char* str);
int memcmp(const void* s1, const void* s2, size_t n);
int strcmp(const char* s1, const char* s2);
int strncmp(const char* s1, const char* s2, size_t n);
void* memset(void* s, int c, size_t n);
char* strchr(const char* s, int c);
char* strrchr(const char* s, int c);

#endif
`
}

func stdboolH() string {
	return `
#ifndef _STDBOOL_H
#define _STDBOOL_H

#define bool _Bool
#define true 1
#define false 0

#endif
`
}

func stdioH() string {
	return `
#ifndef _INC_STDIO
#define _INC_STDIO

#include <stddef.h>

#define FILE void

#define stderr ((FILE*)0)
#define stdout ((FILE*)0)
#define stdin ((FILE*)0)

int fprintf(FILE* stream, const char* format, ...);
int printf(const char* format, ...);
int sprintf(char* str, const char* format, ...);
int snprintf(char* str, size_t size, const char* format, ...);
int vfprintf(FILE* stream, const char* format, void* arg);
int vsprintf(char* str, const char* format, void* arg);
int vsnprintf(char* str, size_t size, const char* format, void* arg);
int fscanf(FILE* stream, const char* format, ...);
int scanf(const char* format, ...);
int sscanf(const char* str, const char* format, ...);

FILE* fopen(const char* filename, const char* mode);
FILE* freopen(const char* filename, const char* mode, FILE* stream);
int fclose(FILE* stream);
int fflush(FILE* stream);
void setbuf(FILE* stream, char* buf);
int setvbuf(FILE* stream, char* buf, int mode, size_t size);

size_t fread(void* ptr, size_t size, size_t count, FILE* stream);
size_t fwrite(const void* ptr, size_t size, size_t count, FILE* stream);
int fgetc(FILE* stream);
char* fgets(char* str, int n, FILE* stream);
int fputc(int c, FILE* stream);
int fputs(const char* str, FILE* stream);
int getc(FILE* stream);
int getchar(void);
int putchar(int c);
int puts(const char* str);
int ungetc(int c, FILE* stream);

long ftell(FILE* stream);
int fseek(FILE* stream, long offset, int origin);
void rewind(FILE* stream);
int feof(FILE* stream);
int ferror(FILE* stream);
void clearerr(FILE* stream);
int remove(const char* filename);
int rename(const char* old, const char* newname);
FILE* tmpfile(void);
char* tmpnam(char* str);

#define EOF (-1)
#define SEEK_SET 0
#define SEEK_CUR 1
#define SEEK_END 2

#endif
`
}

func windowsH() string {
	return `
#ifndef _INC_WINDOWS
#define _INC_WINDOWS

#include <stddef.h>
#include <stdint.h>

typedef __WCHAR_TYPE__ wchar_t;

typedef int BOOL;
typedef unsigned long DWORD;
typedef unsigned int UINT;
typedef void* HWND;
typedef void* HANDLE;
typedef void* HMODULE;
typedef void* HINSTANCE;
typedef void* HDC;
typedef void* HMENU;
typedef void* HICON;
typedef void* HCURSOR;
typedef void* HBRUSH;
typedef void* HFONT;
typedef void* HPEN;
typedef void* HBITMAP;
typedef void* HRGN;
typedef void* HPALETTE;
typedef void* HACCEL;
typedef void* HHOOK;

typedef long long INT_PTR;
typedef unsigned long long UINT_PTR;
typedef long long LONG_PTR;
typedef unsigned long long ULONG_PTR;
typedef long long LPARAM;
typedef unsigned long long WPARAM;

typedef unsigned char  UINT8;
typedef unsigned short UINT16;
typedef unsigned int   UINT32;
typedef unsigned long long UINT64;
typedef signed char  INT8;
typedef signed short INT16;
typedef signed int   INT32;
typedef signed long long INT64;

typedef void* PVOID;
typedef void* LPVOID;

typedef struct _OVERLAPPED { ULONG_PTR Internal; ULONG_PTR InternalHigh; DWORD Offset; DWORD OffsetHigh; PVOID Pointer; HANDLE hEvent; } OVERLAPPED, *LPOVERLAPPED;

typedef char CHAR;
typedef unsigned char UCHAR;
typedef short SHORT;
typedef unsigned short USHORT;
typedef int INT;
typedef unsigned int UINT;
typedef long LONG;
typedef unsigned long ULONG;
typedef long long LONGLONG;
typedef unsigned long long ULONGLONG;

typedef __SIZE_TYPE__ SIZE_T;
typedef __SIZE_TYPE__ *PSIZE_T;

typedef struct _LIST_ENTRY { struct _LIST_ENTRY *Flink; struct _LIST_ENTRY *Blink; } LIST_ENTRY, *PLIST_ENTRY;

typedef char* LPSTR;
typedef char* PSTR;
typedef const char* LPCSTR;
typedef const char* PCSTR;
typedef wchar_t* LPWSTR;
typedef wchar_t* PWSTR;
typedef const wchar_t* LPCWSTR;
typedef const wchar_t* PCWSTR;
typedef const void* LPCVOID;

typedef DWORD COLORREF;

#ifndef _FILETIME_DEFINED
#define _FILETIME_DEFINED
typedef struct _FILETIME {
	DWORD dwLowDateTime;
	DWORD dwHighDateTime;
} FILETIME, *PFILETIME, *LPFILETIME;
#endif

#ifndef _LARGE_INTEGER_DEFINED
#define _LARGE_INTEGER_DEFINED
typedef union _LARGE_INTEGER {
	struct {
		DWORD LowPart;
		LONG HighPart;
	};
	LONGLONG QuadPart;
} LARGE_INTEGER, *PLARGE_INTEGER;
typedef union _ULARGE_INTEGER {
	struct {
		DWORD LowPart;
		DWORD HighPart;
	};
	ULONGLONG QuadPart;
} ULARGE_INTEGER, *PULARGE_INTEGER;
#endif

#define TRUE 1
#define FALSE 0

#define MAX_PATH 260
#define INFINITE 0xFFFFFFFF

#define CALLBACK
#define WINAPI
#define APIENTRY
#define FAR

#endif
`
}

func errnoH() string {
	return `
#ifndef _INC_ERRNO
#define _INC_ERRNO

#define EPERM 1
#define ENOENT 2
#define ESRCH 3
#define EINTR 4
#define EIO 5
#define ENXIO 6
#define E2BIG 7
#define ENOEXEC 8
#define EBADF 9
#define ECHILD 10
#define EAGAIN 11
#define ENOMEM 12
#define EACCES 13
#define EFAULT 14
#define EBUSY 16
#define EEXIST 17
#define EXDEV 18
#define ENODEV 19
#define ENOTDIR 20
#define EISDIR 21
#define EINVAL 22
#define ENFILE 23
#define EMFILE 24
#define ENOTTY 25
#define EFBIG 27
#define ENOSPC 28
#define ESPIPE 29
#define EROFS 30
#define EMLINK 31
#define EPIPE 32
#define EDOM 33
#define ERANGE 34
#define EDEADLK 36
#define ENAMETOOLONG 38
#define ENOLCK 39
#define ENOSYS 40
#define ENOTEMPTY 41

extern int errno;

#endif
`
}

func inttypesH() string {
	return `
#ifndef _INC_INTTYPES
#define _INC_INTTYPES

#include <stdint.h>

typedef long long intmax_t;
typedef unsigned long long uintmax_t;

#define PRId8 "d"
#define PRId16 "d"
#define PRId32 "d"
#define PRId64 "lld"
#define PRIi8 "i"
#define PRIi16 "i"
#define PRIi32 "i"
#define PRIi64 "lli"
#define PRIu8 "u"
#define PRIu16 "u"
#define PRIu32 "u"
#define PRIu64 "llu"
#define PRIx8 "x"
#define PRIx16 "x"
#define PRIx32 "x"
#define PRIx64 "llx"
#define PRIX8 "X"
#define PRIX16 "X"
#define PRIX32 "X"
#define PRIX64 "llX"
#define PRIdLEAST8 "d"
#define PRIdLEAST16 "d"
#define PRIdLEAST32 "d"
#define PRIdLEAST64 "lld"
#define PRIiLEAST8 "i"
#define PRIiLEAST16 "i"
#define PRIiLEAST32 "i"
#define PRIiLEAST64 "lli"
#define PRIuLEAST8 "u"
#define PRIuLEAST16 "u"
#define PRIuLEAST32 "u"
#define PRIuLEAST64 "llu"
#define PRIdFAST8 "d"
#define PRIdFAST16 "d"
#define PRIdFAST32 "d"
#define PRIdFAST64 "lld"
#define PRIiFAST8 "i"
#define PRIiFAST16 "i"
#define PRIiFAST32 "i"
#define PRIiFAST64 "lli"
#define PRIuFAST8 "u"
#define PRIuFAST16 "u"
#define PRIuFAST32 "u"
#define PRIuFAST64 "llu"
#define PRIuMAX "llu"
#define PRIdMAX "lld"
#define PRIiMAX "lli"

#define SCNd8 "d"
#define SCNd16 "d"
#define SCNd32 "d"
#define SCNd64 "lld"
#define SCNu8 "u"
#define SCNu16 "u"
#define SCNu32 "u"
#define SCNu64 "llu"

typedef struct {
    intmax_t quot;
    intmax_t rem;
} imaxdiv_t;

intmax_t imaxabs(intmax_t j);
imaxdiv_t imaxdiv(intmax_t numer, intmax_t denom);
intmax_t strtoimax(const char* nptr, char** endptr, int base);
uintmax_t strtoumax(const char* nptr, char** endptr, int base);

#endif
`
}

func limitsH() string {
	return `
#ifndef _INC_LIMITS
#define _INC_LIMITS

#define CHAR_BIT 8
#define SCHAR_MIN (-128)
#define SCHAR_MAX 127
#define UCHAR_MAX 255
#define CHAR_MIN SCHAR_MIN
#define CHAR_MAX SCHAR_MAX
#define SHRT_MIN (-32768)
#define SHRT_MAX 32767
#define USHRT_MAX 65535
#define INT_MIN (-2147483647 - 1)
#define INT_MAX 2147483647
#define UINT_MAX 0xffffffffU
#define LONG_MIN (-2147483647L - 1)
#define LONG_MAX 2147483647L
#define ULONG_MAX 0xffffffffUL
#define LLONG_MAX 9223372036854775807LL
#define LLONG_MIN (-9223372036854775807LL - 1)
#define ULLONG_MAX 0xffffffffffffffffULL

#define MB_LEN_MAX 5

#define I8_MIN (-127i8 - 1)
#define I8_MAX 127i8
#define UI8_MAX 0xffui8

#define I16_MIN (-32767i16 - 1)
#define I16_MAX 32767i16
#define UI16_MAX 0xffffui16

#define I32_MIN (-2147483647i32 - 1)
#define I32_MAX 2147483647i32
#define UI32_MAX 0xffffffffui32

#define I64_MIN (-9223372036854775807i64 - 1)
#define I64_MAX 9223372036854775807i64
#define UI64_MAX 0xffffffffffffffffui64

#endif
`
}

func mathH() string {
	return `
#ifndef _INC_MATH
#define _INC_MATH

#define M_PI 3.14159265358979323846
#define M_PI_2 1.57079632679489661923
#define M_PI_4 0.78539816339744830962
#define M_1_PI 0.31830988618379067154
#define M_2_PI 0.63661977236758134308
#define M_E 2.7182818284590452354
#define M_LOG2E 1.4426950408889634074
#define M_LOG10E 0.43429448190325182765
#define M_LN2 0.69314718055994530942
#define M_LN10 2.30258509299404568402
#define M_SQRT2 1.41421356237309504880
#define M_SQRT1_2 0.70710678118654752440

#define INFINITY ((double)(1.0 / 0.0))
#define NAN ((double)(0.0 / 0.0))
#define HUGE_VAL INFINITY

double sin(double x);
double cos(double x);
double tan(double x);
double asin(double x);
double acos(double x);
double atan(double x);
double atan2(double y, double x);
double sinh(double x);
double cosh(double x);
double tanh(double x);
double exp(double x);
double log(double x);
double log10(double x);
double log2(double x);
double pow(double x, double y);
double sqrt(double x);
double ceil(double x);
double floor(double x);
double fabs(double x);
double fmod(double x, double y);
double ldexp(double x, int exp);
double frexp(double x, int* exp);
double modf(double x, double* iptr);
double round(double x);
double trunc(double x);
double rint(double x);
double nearbyint(double x);
double remainder(double x, double y);
double copysign(double x, double y);
double nan(const char* tagp);
int isnan(double x);
int isinf(double x);
int isfinite(double x);
int fpclassify(double x);
int signbit(double x);
int isnormal(double x);

float sinf(float x);
float cosf(float x);
float tanf(float x);
float sqrtf(float x);
float fabsf(float x);
float floorf(float x);
float ceilf(float x);
float roundf(float x);
float powf(float x, float y);
float expf(float x);
float logf(float x);
float log10f(float x);
float log2f(float x);
float fmodf(float x, float y);
float copysignf(float x, float y);
int isnanf(float x);
int isinff(float x);

#endif
`
}

func stdargH() string {
	return `
#ifndef _INC_STDARG
#define _INC_STDARG

typedef __builtin_va_list va_list;

#define va_start(ap, param) __builtin_va_start(ap, param)
#define va_end(ap)          __builtin_va_end(ap)
#define va_arg(ap, type)    __builtin_va_arg(ap, type)
#define va_copy(dest, src)  __builtin_va_copy(dest, src)

#endif
`
}
