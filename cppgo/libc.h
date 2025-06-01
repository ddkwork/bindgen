
#ifndef CPPGO_LIBC_H
#define CPPGO_LIBC_H
#pragma once
#if defined(_KERNEL_MODE)
#include <ntifs.h>
#include <wdm.h>
#include <ntstrsafe.h>
#include <cstdarg>

static inline void *Alloc(size_t size) {
#if (WDK_NTDDI_VERSION >= NTDDI_WIN10_VB)
    return ExAllocatePool2(POOL_FLAG_NON_PAGED, size, 'bufT');
#else
    return ExAllocatePoolWithTag(NonPagedPool, size, 'bufT');
#endif
}

static inline void Free(void *ptr) {
    if (ptr) ExFreePoolWithTag(ptr, 'bufT');
}

static inline void Copy(void *dest, const void *src, size_t size) { RtlCopyMemory(dest, src, size); }

static inline void Move(void *dest, const void *src, size_t size) { RtlMoveMemory(dest, src, size); }

static inline bool Compare(const void *ptr1, const void *ptr2, size_t size) { return RtlCompareMemory(ptr1, ptr2, size) == size; }

#else
#include <cstdlib>
#include <cstring>
#include <cstdio>
#include <cstdarg>

static inline void *Alloc(size_t size) { return malloc(size); }

static inline void Free(void *ptr) { free(ptr); }

static inline void Copy(void *dest, const void *src, size_t size) { memcpy(dest, src, size); }

static inline void Move(void *dest, const void *src, size_t size) { memmove(dest, src, size); }

static inline bool Compare(const void *ptr1, const void *ptr2, size_t size) { return memcmp(ptr1, ptr2, size) == 0; }

#endif
#endif//CPPGO_LIBC_H
