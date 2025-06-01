#pragma once
#include "libc.h"
#include "fmt.h"

struct Buffer {
    unsigned char *data;
    size_t         len;
    size_t         cap;

    static Buffer New(size_t capacity = 64) {
        Buffer buf{};
        buf.data = nullptr;
        buf.len = 0;
        buf.cap = 0;
        if (capacity > 0) {
            buf.data = static_cast<unsigned char *>(Alloc(capacity));
            if (buf.data) { buf.cap = capacity; }
        }
        return buf;
    }

    static Buffer FromBytes(const unsigned char *b, size_t n) {
        Buffer buf = New(n);
        if (buf.data && b) {
            Copy(buf.data, b, n);
            buf.len = n;
        }
        return buf;
    }

    static Buffer FromString(const char *s) {
        if (!s) return New();
        size_t length = 0;
        while (s[length] != '\0') length++;
        return FromBytes(reinterpret_cast<const unsigned char *>(s), length);
    }

    void Reset() {
        if (data) {
            Free(data);
            data = nullptr;
        }
        len = 0;
        cap = 0;
    }

    size_t Write(const unsigned char *b, size_t n) {
        if (!b || n == 0) return 0;
        size_t needed = len + n;
        if (needed > cap) {
            size_t newCap = cap * 2;
            if (newCap < needed) newCap = needed;
            if (newCap < 16) newCap = 16;// 最小容量
            auto *newData = static_cast<unsigned char *>(Alloc(newCap));
            if (!newData) return 0;
            if (data) {
                Copy(newData, data, len);
                Free(data);
            }
            data = newData;
            cap = newCap;
        }
        Copy(data + len, b, n);
        len += n;
        return n;
    }

    size_t WriteString(const char *s) {
        if (!s) return 0;
        size_t length = 0;
        while (s[length] != '\0') length++;
        return Write(reinterpret_cast<const unsigned char *>(s), length);
    }

    size_t Read(unsigned char *b, size_t max) {
        if (!b || max == 0 || len == 0) return 0;
        size_t n = (len < max) ? len : max;
        Copy(b, data, n);
        size_t remaining = len - n;
        if (remaining > 0) { Move(data, data + n, remaining); }
        len = remaining;
        return n;
    }

    [[nodiscard]] Buffer Slice(size_t start, size_t end) const {
        if (start >= len) { return New(); }
        if (end > len || end < start) end = len;
        size_t n = end - start;
        if (n == 0) return New();
        return FromBytes(data + start, n);
    }

    [[nodiscard]] bool Equals(const Buffer &other) const {
        if (len != other.len) return false;
        return Compare(data, other.data, len);
    }

    [[nodiscard]] char *ToString() const {
        if (!data || len == 0) return nullptr;
        char *s = static_cast<char *>(Alloc(len + 1));
        if (!s) return nullptr;
        Copy(s, data, len);
        s[len] = '\0';
        return s;
    }

    [[nodiscard]] const unsigned char *Bytes() const { return data; }

    [[nodiscard]] size_t Len() const { return len; }

    [[nodiscard]] size_t Cap() const { return cap; }
};
