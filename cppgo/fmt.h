
#ifndef CPPGO_FMT_H
#define CPPGO_FMT_H
#pragma once
#if defined(_KERNEL_MODE)
#include <wdm.h>
#include <ntstrsafe.h>
#include <cstdarg>

static inline void Print(const char *fmt, ...) {
#if DBG
    va_list args;
    va_start(args, fmt);
    DbgPrintEx(DPFLTR_IHVDRIVER_ID, DPFLTR_ERROR_LEVEL, fmt, args);
    va_end(args);
#endif
}

static inline void Println(const char *fmt, ...) {
#if DBG
    va_list args;
    va_start(args, fmt);
    DbgPrintEx(DPFLTR_IHVDRIVER_ID, DPFLTR_ERROR_LEVEL, fmt, args);
    DbgPrintEx(DPFLTR_IHVDRIVER_ID, DPFLTR_ERROR_LEVEL, "\n", NULL);
    va_end(args);
#endif
}
#else
#include <cstdlib>
#include <cstring>
#include <cstdio>
#include <cstdarg>

static inline void Print(const char *fmt, ...) {
    va_list args;
    va_start(args, fmt);
    vprintf(fmt, args);
    va_end(args);
}

static inline void Println(const char *fmt, ...) {
    va_list args;
    va_start(args, fmt);
    vprintf(fmt, args);
    printf("\n");
    va_end(args);
}
#endif
#endif//CPPGO_FMT_H
