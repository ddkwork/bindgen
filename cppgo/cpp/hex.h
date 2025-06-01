#pragma once

#include "buffer.h"// 使用之前实现的Buffer
#include <cctype>
#include <stdexcept>
#include <sstream>
#include <iomanip>
#include <string>

struct Hex {
    struct InvalidByteError : public std::runtime_error {
        InvalidByteError(char byte) : std::runtime_error("encoding/hex: invalid byte: " + std::string(1, byte)) {}
    };

    struct LengthError : public std::runtime_error {
        LengthError() : std::runtime_error("encoding/hex: odd length hex string") {}
    };

    static Buffer Encode(const Buffer &src) {
        Buffer dst(src.Len() * 2);
        EncodeBuffer(dst, src.Data(), src.Len());
        return dst;
    }

    static Buffer Encode(const uint8_t *data, size_t len) {
        Buffer dst(len * 2);
        EncodeBuffer(dst, data, len);
        return dst;
    }

    static void EncodeBuffer(Buffer &dst, const uint8_t *src, size_t len) {
        static const char hexDigits[] = "0123456789abcdef";

        uint8_t *p = dst.Data();
        if (dst.Cap() < len * 2) { throw std::out_of_range("encoding/hex: buffer too small"); }

        for (size_t i = 0; i < len; i++) {
            *p++ = hexDigits[src[i] >> 4];
            *p++ = hexDigits[src[i] & 0x0F];
        }

        dst.Truncate(len * 2);// 设置实际长度
    }

    static Buffer Decode(const Buffer &src) {
        if (src.Len() % 2 != 0) { throw LengthError(); }

        Buffer dst(src.Len() / 2);
        DecodeBuffer(dst, src.Data(), src.Len());
        return dst;
    }

    static Buffer Decode(const std::string &s) { return Decode(Buffer(s)); }

    static void DecodeBuffer(Buffer &dst, const uint8_t *src, size_t len) {
        if (len % 2 != 0) { throw LengthError(); }

        if (dst.Cap() < len / 2) { throw std::out_of_range("encoding/hex: buffer too small"); }

        for (size_t i = 0; i < len; i += 2) { dst.Append((fromHexChar(src[i]) << 4) | fromHexChar(src[i + 1])); }
    }

    static std::string EncodeToString(const Buffer &src) { return Encode(src).String(); }

    static std::string EncodeToString(const uint8_t *data, size_t len) { return Encode(data, len).String(); }

    static Buffer DecodeString(const std::string &s) { return Decode(Buffer(s)); }

    static std::string Dump(const Buffer &src) {
        std::ostringstream oss;
        Dump(oss, src.Data(), src.Len());
        return oss.str();
    }

    static std::string Dump(const uint8_t *data, size_t len) {
        std::ostringstream oss;
        Dump(oss, data, len);
        return oss.str();
    }

    static std::string ToString(const Buffer &src, size_t offset = 0) {
        std::ostringstream oss;
        Dump(oss, src.Data() + offset, src.Len() - offset);
        return oss.str();
    }

private:
    static uint8_t fromHexChar(uint8_t c) {
        switch (c) {
            case '0' ... '9':
                return c - '0';
            case 'a' ... 'f':
                return c - 'a' + 10;
            case 'A' ... 'F':
                return c - 'A' + 10;
            default:
                throw InvalidByteError(c);
        }
    }

    static void Dump(std::ostream &os, const uint8_t *data, size_t len) {
        const size_t bytesPerLine = 16;
        char         buf[bytesPerLine * 3 + bytesPerLine + 1];

        for (size_t i = 0; i < len; i += bytesPerLine) {
            size_t n = len - i;
            if (n > bytesPerLine) n = bytesPerLine;

            std::memset(buf, ' ', sizeof(buf));
            buf[sizeof(buf) - 1] = '\0';

            os << std::hex << std::setw(8) << std::setfill('0') << i << "  ";

            for (size_t j = 0; j < n; j++) {
                size_t pos = j * 3;
                if (j >= bytesPerLine / 2) pos++;// 中间空格

                uint8_t b = data[i + j];
                buf[pos] = "0123456789abcdef"[b >> 4];
                buf[pos + 1] = "0123456789abcdef"[b & 0x0F];

                char &textChar = buf[bytesPerLine * 3 + (bytesPerLine > 8 ? 1 : 0) + j];
                textChar = (b >= 0x20 && b < 0x7F) ? b : '.';
            }

            if (n < bytesPerLine) {
                for (size_t j = n; j < bytesPerLine; j++) {
                    size_t pos = j * 3;
                    if (j >= bytesPerLine / 2) pos++;// 中间空格

                    buf[pos] = buf[pos + 1] = ' ';
                }
            }

            buf[bytesPerLine * 3] = '|';
            buf[bytesPerLine * 3 + n + 1] = '|';

            os << buf << "\n";
        }

        os << std::hex << std::setw(8) << len << "\n";
        os << std::dec;// 恢复十进制输出
    }

    static void Dump(std::ostream &os, const uint8_t *data, size_t len) {
        const int bytesPerLine = 16;
        char      hex[bytesPerLine * 3 + 1];
        char      ascii[bytesPerLine + 1];

        for (size_t i = 0; i < len; i += bytesPerLine) {
            os << std::setw(8) << std::setfill('0') << std::hex << i << "  ";

            std::memset(hex, ' ', sizeof(hex));
            hex[bytesPerLine * 3] = '\0';

            std::memset(ascii, ' ', sizeof(ascii));
            ascii[bytesPerLine] = '\0';

            size_t lineLength = (len - i) > bytesPerLine ? bytesPerLine : (len - i);

            for (size_t j = 0; j < lineLength; j++) {
                uint8_t b = data[i + j];
                size_t  pos = j * 3;

                hex[pos] = "0123456789abcdef"[b >> 4];
                hex[pos + 1] = "0123456789abcdef"[b & 0x0F];

                ascii[j] = (b >= 0x20 && b < 0x7F) ? b : '.';
            }

            hex[bytesPerLine * 3] = ' ';

            os << hex << "  |" << ascii << "|\n";
        }

        os << std::setw(8) << std::setfill('0') << std::hex << len << "\n";
        os << std::dec;// 恢复十进制输出
    }
};
