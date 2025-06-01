#pragma once

#include "buffer.h"
#include <cctype>
#include <stdexcept>
#include <sstream>
#include <iomanip>
#include <string>

struct HexError : std::runtime_error {
    explicit HexError(const std::string &msg) : std::runtime_error(msg) {}
};

struct Hex {
    static Buffer Encode(const Buffer &src) {
        Buffer            dst(src.Len() * 2);
        static const char hexDigits[] = "0123456789abcdef";

        for (size_t i = 0; i < src.Len(); i++) {
            uint8_t b = src[i];
            dst.Append(hexDigits[b >> 4]);
            dst.Append(hexDigits[b & 0x0F]);
        }
        return dst;
    }

    static std::string EncodeToString(const Buffer &src) { return Encode(src).String(); }

    static Buffer Decode(const Buffer &src) {
        if (src.Len() % 2 != 0) { throw HexError("odd length hex string"); }

        Buffer dst(src.Len() / 2);
        for (size_t i = 0; i < src.Len(); i += 2) {
            uint8_t hi = src[i];
            uint8_t lo = src[i + 1];

            uint8_t h = hexCharToByte(hi);
            uint8_t l = hexCharToByte(lo);

            dst.Append((h << 4) | l);
        }
        return dst;
    }

    static Buffer DecodeString(const std::string &s) {
        Buffer src(s);
        return Decode(src);
    }

    static std::string Dump(const Buffer &src) {
        std::ostringstream oss;
        const size_t       bytesPerLine = 16;

        for (size_t i = 0; i < src.Len(); i += bytesPerLine) {
            oss << std::hex << std::setw(8) << std::setfill('0') << i << "  ";

            for (size_t j = 0; j < bytesPerLine; j++) {
                if (i + j < src.Len()) {
                    oss << std::setw(2) << static_cast<int>(src[i + j]) << " ";
                } else {
                    oss << "   ";
                }

                if (j == 7) oss << " ";// 中间分隔符
            }

            oss << " |";

            for (size_t j = 0; j < bytesPerLine; j++) {
                if (i + j >= src.Len()) break;

                uint8_t b = src[i + j];
                oss << (b >= 0x20 && b < 0x7F ? static_cast<char>(b) : '.');
            }

            oss << "|\n";
        }

        oss << std::setw(8) << std::setfill('0') << src.Len() << "\n";
        return oss.str();
    }

private:
    static uint8_t hexCharToByte(uint8_t c) {
        if (c >= '0' && c <= '9') return c - '0';
        if (c >= 'a' && c <= 'f') return c - 'a' + 10;
        if (c >= 'A' && c <= 'F') return c - 'A' + 10;
        throw HexError("invalid hex character: " + std::string(1, c));
    }
};

#include "hex.h"
#include <iostream>

int main() {
    Buffer data = {0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21};// "Hello, World!"

    Buffer encoded = Hex::Encode(data);
    std::cout << "Encoded: " << encoded.String() << std::endl;

    std::string hexStr = Hex::EncodeToString(data);
    std::cout << "String encoded: " << hexStr << std::endl;

    Buffer decoded = Hex::Decode(encoded);
    std::cout << "Decoded: " << decoded.String() << std::endl;

    std::cout << "Hex dump:\n" << Hex::Dump(data);

    try {
        Buffer invalid = Buffer("123");
        Hex::Decode(invalid);// 奇数长度错误
    } catch (const HexError &e) { std::cerr << "Error: " << e.what() << std::endl; }

    return 0;
}
