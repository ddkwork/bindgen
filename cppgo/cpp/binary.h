#pragma once

#include "buffer.h"// 使用之前实现的Buffer结构体
#include <cstdint>
#include <cstring>
#include <stdexcept>

enum ByteOrder { LittleEndian = 0, BigEndian };

struct Binary {
    static const ByteOrder LittleEndian = ::LittleEndian;
    static const ByteOrder BigEndian = ::BigEndian;

    struct Varint {
        static size_t Size(uint64_t x) {
            size_t n = 1;
            while (x >= 0x80) {
                x >>= 7;
                n++;
            }
            return n;
        }

        static size_t Size(int64_t x) {
            uint64_t ux = static_cast<uint64_t>(x) << 1;
            if (x < 0) ux = ~ux;
            return Size(ux);
        }

        static size_t Put(Buffer &b, uint64_t x) {
            size_t start = b.Len();
            while (x >= 0x80) {
                b.Append(static_cast<uint8_t>(x | 0x80));
                x >>= 7;
            }
            b.Append(static_cast<uint8_t>(x));
            return b.Len() - start;
        }

        static size_t Put(Buffer &b, int64_t x) {
            uint64_t ux = static_cast<uint64_t>(x) << 1;
            if (x < 0) ux = ~ux;
            return Put(b, ux);
        }

        static size_t Read(const Buffer &b, uint64_t *value, size_t offset = 0) {
            *value = 0;
            for (size_t shift = 0; shift < 64; shift += 7) {
                if (offset >= b.Len()) return 0;

                uint8_t c = b.At(offset++);
                *value |= static_cast<uint64_t>(c & 0x7F) << shift;

                if ((c & 0x80) == 0) { return offset; }
            }
            return 0;// 溢出
        }

        static size_t Read(const Buffer &b, int64_t *value, size_t offset = 0) {
            uint64_t ux;
            size_t   newOffset = Read(b, &ux, offset);
            if (!newOffset) return 0;

            *value = static_cast<int64_t>(ux >> 1);
            if (ux & 1) *value = ~(*value);
            return newOffset;
        }
    };

    template<typename T>
    static size_t Size() {
        return sizeof(T);
    }

    template<>
    static size_t Size<uint64_t>() {
        return 0;
    }// 特化无效，使用Varint::Size

    template<>
    static size_t Size<int64_t>() {
        return 0;
    }// 特化无效，使用Varint::Size

    template<typename T>
    static void Read(const Buffer &b, ByteOrder order, T *data, size_t offset = 0) {
        if (offset + sizeof(T) > b.Len()) { throw std::out_of_range("buffer underflow"); }

        const uint8_t *src = b.Data() + offset;
        uint8_t       *dst = reinterpret_cast<uint8_t *>(data);

        if (order == LittleEndian) {
            for (size_t i = 0; i < sizeof(T); i++) { dst[i] = src[i]; }
        } else {
            for (size_t i = 0; i < sizeof(T); i++) { dst[i] = src[sizeof(T) - 1 - i]; }
        }
    }

    static void Read(const Buffer &b, ByteOrder order, uint64_t *value, size_t offset = 0) { Varint::Read(b, value, offset); }

    static void Read(const Buffer &b, ByteOrder order, int64_t *value, size_t offset = 0) { Varint::Read(b, value, offset); }

    template<typename T>
    static void Write(Buffer &b, ByteOrder order, const T &data) {
        const uint8_t *src = reinterpret_cast<const uint8_t *>(&data);

        if (order == LittleEndian) {
            for (size_t i = 0; i < sizeof(T); i++) { b.Append(src[i]); }
        } else {
            for (size_t i = 0; i < sizeof(T); i++) { b.Append(src[sizeof(T) - 1 - i]); }
        }
    }

    static void Write(Buffer &b, ByteOrder order, uint64_t data) { Varint::Put(b, data); }

    static void Write(Buffer &b, ByteOrder order, int64_t data) { Varint::Put(b, data); }

#define DEFINE_TYPE_READ_WRITE(Type)                                              \
    static Type Read##Type(const Buffer &b, ByteOrder order, size_t offset = 0) { \
        Type value;                                                               \
        Read(b, order, &value, offset);                                           \
        return value;                                                             \
    }                                                                             \
    static void Write##Type(Buffer &b, ByteOrder order, Type value) { Write(b, order, value); }

    DEFINE_TYPE_READ_WRITE(uint16_t)
    DEFINE_TYPE_READ_WRITE(uint32_t)
    DEFINE_TYPE_READ_WRITE(uint64_t)
    DEFINE_TYPE_READ_WRITE(int16_t)
    DEFINE_TYPE_READ_WRITE(int32_t)
    DEFINE_TYPE_READ_WRITE(int64_t)
    DEFINE_TYPE_READ_WRITE(float)
    DEFINE_TYPE_READ_WRITE(double)

#undef DEFINE_TYPE_READ_WRITE

    static uint64_t ReadUvarint(const Buffer &b, size_t *offset = nullptr) {
        uint64_t value;
        size_t   newOffset = Varint::Read(b, &value, offset ? *offset : 0);
        if (offset) *offset = newOffset;
        return value;
    }

    static int64_t ReadVarint(const Buffer &b, size_t *offset = nullptr) {
        int64_t value;
        size_t  newOffset = Varint::Read(b, &value, offset ? *offset : 0);
        if (offset) *offset = newOffset;
        return value;
    }

    static void WriteUvarint(Buffer &b, uint64_t value) { Varint::Put(b, value); }

    static void WriteVarint(Buffer &b, int64_t value) { Varint::Put(b, value); }

    static ByteOrder NativeOrder() {
        union {
            uint16_t value;
            uint8_t  data[2];
        } test = {0x0100};

        return test.data[0] == 0x01 ? BigEndian : LittleEndian;
    }
};

using Binary = struct Binary;
