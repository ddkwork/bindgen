#pragma once
#include <cstdint>
#include <vector>
#include <array>
#include <string>
#include <cstring>

struct md5_Digest {
    std::array<uint8_t, 16> value;

    std::string toString() const {
        const char  hex[] = "0123456789abcdef";
        std::string s(32, ' ');
        for (int i = 0; i < 16; i++) {
            s[2 * i] = hex[value[i] >> 4];
            s[2 * i + 1] = hex[value[i] & 0x0f];
        }
        return s;
    }

    bool operator==(const md5_Digest &other) const { return value == other.value; }

    bool operator!=(const md5_Digest &other) const { return !(*this == other); }
};

struct md5_Hash {
    uint32_t                s0, s1, s2, s3;// 状态寄存器 (ABCD)
    uint64_t                len;           // 已处理数据长度（位）
    std::array<uint8_t, 64> buf;           // 输入缓冲区
    size_t                  nx;            // 缓冲区中有效字节数

    md5_Hash() { reset(); }

    void reset() {
        s0 = 0x67452301;
        s1 = 0xEFCDAB89;
        s2 = 0x98BADCFE;
        s3 = 0x10325476;
        len = 0;
        nx = 0;
        buf.fill(0);
    }

    void write(const uint8_t *data, size_t size) {
        len += size * 8;// 更新总位数

        if (nx > 0) {
            size_t n = std::min(size, 64 - nx);
            std::memcpy(buf.data() + nx, data, n);
            nx += n;
            if (nx == 64) {
                block(buf.data());
                nx = 0;
            }
            data += n;
            size -= n;
        }

        while (size >= 64) {
            block(data);
            data += 64;
            size -= 64;
        }

        if (size > 0) {
            std::memcpy(buf.data(), data, size);
            nx = size;
        }
    }

    md5_Digest sum() {
        uint32_t a = s0, b = s1, c = s2, d = s3;

        std::array<uint8_t, 64> padding;
        size_t                  paddingSize = 64 - nx;
        if (paddingSize < 9) { paddingSize += 64; }

        std::memset(padding.data(), 0, paddingSize);
        padding[0] = 0x80;// 第一位为1，其余为0

        uint64_t bits = len;
        for (int i = 0; i < 8; i++) {
            padding[paddingSize - 8 + i] = static_cast<uint8_t>(bits);
            bits >>= 8;
        }

        write(padding.data(), paddingSize);

        md5_Digest digest;
        store32(digest.value.data(), s0);
        store32(digest.value.data() + 4, s1);
        store32(digest.value.data() + 8, s2);
        store32(digest.value.data() + 12, s3);

        reset();

        return digest;
    }

private:
    void block(const uint8_t *chunk) {
        uint32_t X[16];
        for (int i = 0; i < 16; i++) { X[i] = load32(chunk + 4 * i); }

        uint32_t a = s0;
        uint32_t b = s1;
        uint32_t c = s2;
        uint32_t d = s3;

        a = FF(a, b, c, d, X[0], 7, 0xd76aa478);
        d = FF(d, a, b, c, X[1], 12, 0xe8c7b756);
        c = FF(c, d, a, b, X[2], 17, 0x242070db);
        b = FF(b, c, d, a, X[3], 22, 0xc1bdceee);
        a = FF(a, b, c, d, X[4], 7, 0xf57c0faf);
        d = FF(d, a, b, c, X[5], 12, 0x4787c62a);
        c = FF(c, d, a, b, X[6], 17, 0xa8304613);
        b = FF(b, c, d, a, X[7], 22, 0xfd469501);
        a = FF(a, b, c, d, X[8], 7, 0x698098d8);
        d = FF(d, a, b, c, X[9], 12, 0x8b44f7af);
        c = FF(c, d, a, b, X[10], 17, 0xffff5bb1);
        b = FF(b, c, d, a, X[11], 22, 0x895cd7be);
        a = FF(a, b, c, d, X[12], 7, 0x6b901122);
        d = FF(d, a, b, c, X[13], 12, 0xfd987193);
        c = FF(c, d, a, b, X[14], 17, 0xa679438e);
        b = FF(b, c, d, a, X[15], 22, 0x49b40821);

        a = GG(a, b, c, d, X[1], 5, 0xf61e2562);
        d = GG(d, a, b, c, X[6], 9, 0xc040b340);
        c = GG(c, d, a, b, X[11], 14, 0x265e5a51);
        b = GG(b, c, d, a, X[0], 20, 0xe9b6c7aa);
        a = GG(a, b, c, d, X[5], 5, 0xd62f105d);
        d = GG(d, a, b, c, X[10], 9, 0x02441453);
        c = GG(c, d, a, b, X[15], 14, 0xd8a1e681);
        b = GG(b, c, d, a, X[4], 20, 0xe7d3fbc8);
        a = GG(a, b, c, d, X[9], 5, 0x21e1cde6);
        d = GG(d, a, b, c, X[14], 9, 0xc33707d6);
        c = GG(c, d, a, b, X[3], 14, 0xf4d50d87);
        b = GG(b, c, d, a, X[8], 20, 0x455a14ed);
        a = GG(a, b, c, d, X[13], 5, 0xa9e3e905);
        d = GG(d, a, b, c, X[2], 9, 0xfcefa3f8);
        c = GG(c, d, a, b, X[7], 14, 0x676f02d9);
        b = GG(b, c, d, a, X[12], 20, 0x8d2a4c8a);

        a = HH(a, b, c, d, X[5], 4, 0xfffa3942);
        d = HH(d, a, b, c, X[8], 11, 0x8771f681);
        c = HH(c, d, a, b, X[11], 16, 0x6d9d6122);
        b = HH(b, c, d, a, X[14], 23, 0xfde5380c);
        a = HH(a, b, c, d, X[1], 4, 0xa4beea44);
        d = HH(d, a, b, c, X[4], 11, 0x4bdecfa9);
        c = HH(c, d, a, b, X[7], 16, 0xf6bb4b60);
        b = HH(b, c, d, a, X[10], 23, 0xbebfbc70);
        a = HH(a, b, c, d, X[13], 4, 0x289b7ec6);
        d = HH(d, a, b, c, X[0], 11, 0xeaa127fa);
        c = HH(c, d, a, b, X[3], 16, 0xd4ef3085);
        b = HH(b, c, d, a, X[6], 23, 0x04881d05);
        a = HH(a, b, c, d, X[9], 4, 0xd9d4d039);
        d = HH(d, a, b, c, X[12], 11, 0xe6db99e5);
        c = HH(c, d, a, b, X[15], 16, 0x1fa27cf8);
        b = HH(b, c, d, a, X[2], 23, 0xc4ac5665);

        a = II(a, b, c, d, X[0], 6, 0xf4292244);
        d = II(d, a, b, c, X[7], 10, 0x432aff97);
        c = II(c, d, a, b, X[14], 15, 0xab9423a7);
        b = II(b, c, d, a, X[5], 21, 0xfc93a039);
        a = II(a, b, c, d, X[12], 6, 0x655b59c3);
        d = II(d, a, b, c, X[3], 10, 0x8f0ccc92);
        c = II(c, d, a, b, X[10], 15, 0xffeff47d);
        b = II(b, c, d, a, X[1], 21, 0x85845dd1);
        a = II(a, b, c, d, X[8], 6, 0x6fa87e4f);
        d = II(d, a, b, c, X[15], 10, 0xfe2ce6e0);
        c = II(c, d, a, b, X[6], 15, 0xa3014314);
        b = II(b, c, d, a, X[13], 21, 0x4e0811a1);
        a = II(a, b, c, d, X[4], 6, 0xf7537e82);
        d = II(d, a, b, c, X[11], 10, 0xbd3af235);
        c = II(c, d, a, b, X[2], 15, 0x2ad7d2bb);
        b = II(b, c, d, a, X[9], 21, 0xeb86d391);

        s0 += a;
        s1 += b;
        s2 += c;
        s3 += d;
    }

    static uint32_t F(uint32_t x, uint32_t y, uint32_t z) { return (x & y) | (~x & z); }

    static uint32_t G(uint32_t x, uint32_t y, uint32_t z) { return (x & z) | (y & ~z); }

    static uint32_t H(uint32_t x, uint32_t y, uint32_t z) { return x ^ y ^ z; }

    static uint32_t I(uint32_t x, uint32_t y, uint32_t z) { return y ^ (x | ~z); }

    static uint32_t FF(uint32_t a, uint32_t b, uint32_t c, uint32_t d, uint32_t x, uint32_t s, uint32_t ac) {
        a += F(b, c, d) + x + ac;
        a = rotateLeft(a, s);
        return a + b;
    }

    static uint32_t GG(uint32_t a, uint32_t b, uint32_t c, uint32_t d, uint32_t x, uint32_t s, uint32_t ac) {
        a += G(b, c, d) + x + ac;
        a = rotateLeft(a, s);
        return a + b;
    }

    static uint32_t HH(uint32_t a, uint32_t b, uint32_t c, uint32_t d, uint32_t x, uint32_t s, uint32_t ac) {
        a += H(b, c, d) + x + ac;
        a = rotateLeft(a, s);
        return a + b;
    }

    static uint32_t II(uint32_t a, uint32_t b, uint32_t c, uint32_t d, uint32_t x, uint32_t s, uint32_t ac) {
        a += I(b, c, d) + x + ac;
        a = rotateLeft(a, s);
        return a + b;
    }

    static uint32_t rotateLeft(uint32_t x, uint32_t n) { return (x << n) | (x >> (32 - n)); }

    static uint32_t load32(const uint8_t *b) {
        return (static_cast<uint32_t>(b[0])) | (static_cast<uint32_t>(b[1]) << 8) | (static_cast<uint32_t>(b[2]) << 16) | (static_cast<uint32_t>(b[3]) << 24);
    }

    static void store32(uint8_t *b, uint32_t w) {
        b[0] = static_cast<uint8_t>(w);
        b[1] = static_cast<uint8_t>(w >> 8);
        b[2] = static_cast<uint8_t>(w >> 16);
        b[3] = static_cast<uint8_t>(w >> 24);
    }
};

struct md5 {
    static md5_Hash New() { return md5_Hash(); }

    static md5_Digest Sum(const std::vector<uint8_t> &data) {
        md5_Hash h = New();
        h.write(data.data(), data.size());
        return h.sum();
    }

    static md5_Digest Sum(const std::string &str) { return Sum(std::vector<uint8_t>(str.begin(), str.end())); }

    static md5_Digest File(const std::string &filename) {
        FILE *file = fopen(filename.c_str(), "rb");
        if (!file) { throw std::runtime_error("cannot open file: " + filename); }

        md5_Hash h = New();
        uint8_t  buffer[4096];
        size_t   bytesRead;

        while ((bytesRead = fread(buffer, 1, sizeof(buffer), file)) > 0) { h.write(buffer, bytesRead); }

        fclose(file);
        return h.sum();
    }
};

void testMD5Functions() {
    std::string input = "Hello, MD5!";
    md5_Digest  digest = md5::Sum(input);
    std::cout << "MD5(\"" << input << "\") = " << digest.toString() << std::endl;

    std::vector<uint8_t> data = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05};
    digest = md5::Sum(data);
    std::cout << "MD5(binary) = " << digest.toString() << std::endl;

    md5_Hash hasher = md5::New();

    std::vector<std::string> chunks = {"This is the first part of a ", "very long message that needs ", "to be hashed in chunks."};

    for (const auto &chunk: chunks) { hasher.write(reinterpret_cast<const uint8_t *>(chunk.data()), chunk.size()); }

    digest = hasher.sum();
    std::cout << "MD5(chunked) = " << digest.toString() << std::endl;

    try {
        digest = md5::File("example.txt");
        std::cout << "MD5(file) = " << digest.toString() << std::endl;
    } catch (const std::exception &e) { std::cerr << "Error: " << e.what() << std::endl; }

    std::string testCases[] = {"",
                               "d41d8cd98f00b204e9800998ecf8427e",
                               "a",
                               "0cc175b9c0f1b6a831c399e269772661",
                               "abc",
                               "900150983cd24fb0d6963f7d28e17f72",
                               "message digest",
                               "f96b697d7cb7938d525a2f31aaf161d0",
                               "abcdefghijklmnopqrstuvwxyz",
                               "c3fcd3d76192e4007dfb496cca67e13b"};

    for (size_t i = 0; i < sizeof(testCases) / sizeof(testCases[0]); i += 2) {
        digest = md5::Sum(testCases[i]);
        std::string expected = testCases[i + 1];
        std::string actual = digest.toString();

        std::cout << "Test \"" << testCases[i] << "\": ";
        if (actual == expected) {
            std::cout << "PASSED";
        } else {
            std::cout << "FAILED (expected " << expected << ", got " << actual << ")";
        }
        std::cout << std::endl;
    }
}
