
#pragma once
#pragma warning(disable : 4201)
typedef unsigned long long QWORD;
typedef unsigned __int64 UINT64, *PUINT64;
typedef unsigned long DWORD;
typedef int BOOL;
typedef unsigned char BYTE;
typedef unsigned short WORD;
typedef int INT;
typedef unsigned int UINT;
typedef unsigned int *PUINT;
typedef unsigned __int64 ULONG64, *PULONG64;
typedef unsigned __int64 DWORD64, *PDWORD64;
typedef char CHAR;
typedef wchar_t WCHAR;
#define VOID void
typedef unsigned char UCHAR;
typedef unsigned short USHORT;
typedef unsigned long ULONG;
typedef UCHAR BOOLEAN;
typedef BOOLEAN *PBOOLEAN;
typedef signed char INT8, *PINT8;
typedef signed short INT16, *PINT16;
typedef signed int INT32, *PINT32;
typedef signed __int64 INT64, *PINT64;
typedef unsigned char UINT8, *PUINT8;
typedef unsigned short UINT16, *PUINT16;
typedef unsigned int UINT32, *PUINT32;
typedef unsigned __int64 UINT64, *PUINT64;
#define NULL_ZERO 0
#define NULL64_ZERO 0ull
#define FALSE 0
#define TRUE 1
#define UPPER_56_BITS 0xffffffffffffff00
#define UPPER_48_BITS 0xffffffffffff0000
#define UPPER_32_BITS 0xffffffff00000000
#define LOWER_32_BITS 0x00000000ffffffff
#define LOWER_16_BITS 0x000000000000ffff
#define LOWER_8_BITS 0x00000000000000ff
#define SECOND_LOWER_8_BITS 0x000000000000ff00
#define UPPER_48_BITS_AND_LOWER_8_BITS 0xffffffffffff00ff

typedef struct GUEST_REGS {
  UINT64 rax;
  UINT64 rcx;
  UINT64 rdx;
  UINT64 rbx;
  UINT64 rsp;
  UINT64 rbp;
  UINT64 rsi;
  UINT64 rdi;
  UINT64 r8;
  UINT64 r9;
  UINT64 r10;
  UINT64 r11;
  UINT64 r12;
  UINT64 r13;
  UINT64 r14;
  UINT64 r15;
} GUEST_REGS, *PGUEST_REGS;

typedef struct GUEST_EXTRA_REGISTERS {
  UINT16 CS;
  UINT16 DS;
  UINT16 FS;
  UINT16 GS;
  UINT16 ES;
  UINT16 SS;
  UINT64 RFLAGS;
  UINT64 RIP;
} GUEST_EXTRA_REGISTERS, *PGUEST_EXTRA_REGISTERS;

typedef struct _SCRIPT_ENGINE_GENERAL_REGISTERS {
  UINT64 *StackBuffer;
  UINT64 *GlobalVariablesList;
  UINT64 StackIndx;
  UINT64 StackBaseIndx;
  UINT64 ReturnValue;
} SCRIPT_ENGINE_GENERAL_REGISTERS, *PSCRIPT_ENGINE_GENERAL_REGISTERS;

typedef struct _CR3_TYPE {
  union {
    UINT64 Flags;

    struct {
      UINT64 Pcid : 12;
      UINT64 PageFrameNumber : 36;
      UINT64 Reserved1 : 12;
      UINT64 Reserved_2 : 3;
      UINT64 PcidInvalidate : 1;
    } Fields;
  };
} CR3_TYPE, *PCR3_TYPE;
