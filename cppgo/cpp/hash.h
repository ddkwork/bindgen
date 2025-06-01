#pragma once
#include <cstdint>
#include <string>
#include <vector>
#include <array>
#include <functional>
#include <memory>
#include <cstring>
#include <stdexcept>

struct hash_Hash {
    virtual ~hash_Hash() = default;

    virtual void Reset() = 0;

    virtual void Write(const uint8_t *data, size_t size) = 0;

    virtual size_t Size() const = 0;

    virtual size_t BlockSize() const = 0;

    virtual std::vector<uint8_t> Sum(const std::vector<uint8_t> &prefix = {}) = 0;

    void WriteString(const std::string &s) { Write(reinterpret_cast<const uint8_t *>(s.data()), s.size()); }

    template<typename T>
    void WriteInt(T value) {
        static_assert(std::is_integral<T>::value, "Integer required");
        Write(reinterpret_cast<const uint8_t *>(&value), sizeof(T));
    }
};

struct hash_Adler32 : public hash_Hash {
    uint32_t                a = 1;
    uint32_t                b = 0;
    static constexpr size_t mod = 65521;

    void Reset() override {
        a = 1;
        b = 0;
    }

    void Write(const uint8_t *data, size_t size) override {
        for (size_t i = 0; i < size; i++) {
            a = (a + data[i]) % mod;
            b = (b + a) % mod;
        }
    }

    size_t Size() const override { return 4; }

    size_t BlockSize() const override { return 4; }

    std::vector<uint8_t> Sum(const std::vector<uint8_t> &prefix = {}) override {
        uint32_t             sum = (b << 16) | a;
        std::vector<uint8_t> result(prefix);
        for (int i = 0; i < 4; i++) { result.push_back(static_cast<uint8_t>(sum >> (24 - i * 8))); }
        return result;
    }

    static uint32_t Checksum(const uint8_t *data, size_t size) {
        hash_Adler32 hasher;
        hasher.Write(data, size);
        auto sum = hasher.Sum();
        return (static_cast<uint32_t>(sum[0]) << 24) | (static_cast<uint32_t>(sum[1]) << 16) | (static_cast<uint32_t>(sum[2]) << 8) | static_cast<uint32_t>(sum[3]);
    }
};

struct hash_Fnv32 : public hash_Hash {
    uint32_t       sum = 0;
    const uint32_t prime;
    const uint32_t offset;

    hash_Fnv32(uint32_t p, uint32_t off) : prime(p), offset(off) { Reset(); }

    void Reset() override { sum = offset; }

    void Write(const uint8_t *data, size_t size) override {
        for (size_t i = 0; i < size; i++) {
            sum *= prime;
            sum ^= data[i];
        }
    }

    size_t Size() const override { return 4; }

    size_t BlockSize() const override { return 1; }

    std::vector<uint8_t> Sum(const std::vector<uint8_t> &prefix = {}) override {
        std::vector<uint8_t> result(prefix);
        for (int i = 0; i < 4; i++) { result.push_back(static_cast<uint8_t>(sum >> (24 - i * 8))); }
        return result;
    }
};

struct hash_Fnv64 : public hash_Hash {
    uint64_t       sum = 0;
    const uint64_t prime;
    const uint64_t offset;

    hash_Fnv64(uint64_t p, uint64_t off) : prime(p), offset(off) { Reset(); }

    void Reset() override { sum = offset; }

    void Write(const uint8_t *data, size_t size) override {
        for (size_t i = 0; i < size; i++) {
            sum *= prime;
            sum ^= data[i];
        }
    }

    size_t Size() const override { return 8; }

    size_t BlockSize() const override { return 1; }

    std::vector<uint8_t> Sum(const std::vector<uint8_t> &prefix = {}) override {
        std::vector<uint8_t> result(prefix);
        for (int i = 0; i < 8; i++) { result.push_back(static_cast<uint8_t>(sum >> (56 - i * 8))); }
        return result;
    }
};

struct hash_Crc32 : public hash_Hash {
    uint32_t                  crc = 0;
    static constexpr uint32_t polynomial = 0xEDB88320;
    static uint32_t           table[256];
    static bool               tableInitialized;

    hash_Crc32() {
        if (!tableInitialized) {
            InitializeTable();
            tableInitialized = true;
        }
        Reset();
    }

    static void InitializeTable() {
        for (uint32_t i = 0; i < 256; i++) {
            uint32_t c = i;
            for (int j = 0; j < 8; j++) {
                if (c & 1) {
                    c = polynomial ^ (c >> 1);
                } else {
                    c >>= 1;
                }
            }
            table[i] = c;
        }
    }

    void Reset() override { crc = 0xFFFFFFFF; }

    void Write(const uint8_t *data, size_t size) override {
        for (size_t i = 0; i < size; i++) { crc = table[(crc ^ data[i]) & 0xFF] ^ (crc >> 8); }
    }

    size_t Size() const override { return 4; }

    size_t BlockSize() const override { return 1; }

    std::vector<uint8_t> Sum(const std::vector<uint8_t> &prefix = {}) override {
        uint32_t             finalCrc = crc ^ 0xFFFFFFFF;
        std::vector<uint8_t> result(prefix);
        for (int i = 0; i < 4; i++) { result.push_back(static_cast<uint8_t>(finalCrc >> (24 - i * 8))); }
        return result;
    }

    static uint32_t Checksum(const uint8_t *data, size_t size) {
        hash_Crc32 hasher;
        hasher.Write(data, size);
        auto sum = hasher.Sum();
        return (static_cast<uint32_t>(sum[0]) << 24) | (static_cast<uint32_t>(sum[1]) << 16) | (static_cast<uint32_t>(sum[2]) << 8) | static_cast<uint32_t>(sum[3]);
    }
};

uint32_t hash_Crc32::table[256];
bool     hash_Crc32::tableInitialized = false;

struct hash_Sha256 : public hash_Hash {
    std::array<uint32_t, 8> state;
    std::array<uint8_t, 64> buffer;
    uint64_t                length = 0;
    size_t                  bufferSize = 0;

    static constexpr uint32_t k[64] = {0x428a2f98, 0x71374491, 0xb5c0fbcf, 0xe9b5dba5, 0x3956c25b, 0x59f111f1, 0x923f82a4, 0xab1c5ed5, 0xd807aa98, 0x12835b01, 0x243185be,
                                       0x550c7dc3, 0x72be5d74, 0x80deb1fe, 0x9bdc06a7, 0xc19bf174, 0xe49b69c1, 0xefbe4786, 0x0fc19dc6, 0x240ca1cc, 0x2de92c6f, 0x4a7484aa,
                                       0x5cb0a9dc, 0x76f988da, 0x983e5152, 0xa831c66d, 0xb00327c8, 0xbf597fc7, 0xc6e00bf3, 0xd5a79147, 0x06ca6351, 0x14292967, 0x27b70a85,
                                       0x2e1b2138, 0x4d2c6dfc, 0x53380d13, 0x650a7354, 0x766a0abb, 0x81c2c92e, 0x92722c85, 0xa2bfe8a1, 0xa81a664b, 0xc24b8b70, 0xc76c51a3,
                                       0xd192e819, 0xd6990624, 0xf40e3585, 0x106aa070, 0x19a4c116, 0x1e376c08, 0x2748774c, 0x34b0bcb5, 0x391c0cb3, 0x4ed8aa4a, 0x5b9cca4f,
                                       0x682e6ff3, 0x748f82ee, 0x78a5636f, 0x84c87814, 0x8cc70208, 0x90befffa, 0xa4506ceb, 0xbef9a3f7, 0xc67178f2};

    hash_Sha256() { Reset(); }

    void Reset() override {
        state = {0x6a09e667, 0xbb67ae85, 0x3c6ef372, 0xa54ff53a, 0x510e527f, 0x9b05688c, 0x1f83d9ab, 0x5be0cd19};
        buffer.fill(0);
        length = 0;
        bufferSize = 0;
    }

    void Write(const uint8_t *data, size_t size) override {
        length += size;

        while (size > 0) {
            size_t toCopy = std::min(size, 64 - bufferSize);
            std::memcpy(buffer.data() + bufferSize, data, toCopy);
            bufferSize += toCopy;
            data += toCopy;
            size -= toCopy;

            if (bufferSize == 64) {
                ProcessBlock();
                bufferSize = 0;
            }
        }
    }

    void ProcessBlock() {
        std::array<uint32_t, 64> w;

        for (int i = 0; i < 16; i++) {
            w[i] = (static_cast<uint32_t>(buffer[i * 4]) << 24) | (static_cast<uint32_t>(buffer[i * 4 + 1]) << 16) | (static_cast<uint32_t>(buffer[i * 4 + 2]) << 8) |
                   static_cast<uint32_t>(buffer[i * 4 + 3]);
        }

        for (int i = 16; i < 64; i++) {
            uint32_t s0 = std::rotr(w[i - 15], 7) ^ std::rotr(w[i - 15], 18) ^ (w[i - 15] >> 3);
            uint32_t s1 = std::rotr(w[i - 2], 17) ^ std::rotr(w[i - 2], 19) ^ (w[i - 2] >> 10);
            w[i] = w[i - 16] + s0 + w[i - 7] + s1;
        }

        auto [a, b, c, d, e, f, g, h] = state;

        for (int i = 0; i < 64; i++) {
            uint32_t S1 = std::rotr(e, 6) ^ std::rotr(e, 11) ^ std::rotr(e, 25);
            uint32_t ch = (e & f) ^ ((~e) & g);
            uint32_t temp1 = h + S1 + ch + k[i] + w[i];
            uint32_t S0 = std::rotr(a, 2) ^ std::rotr(a, 13) ^ std::rotr(a, 22);
            uint32_t maj = (a & b) ^ (a & c) ^ (b & c);
            uint32_t temp2 = S0 + maj;

            h = g;
            g = f;
            f = e;
            e = d + temp1;
            d = c;
            c = b;
            b = a;
            a = temp1 + temp2;
        }

        state[0] += a;
        state[1] += b;
        state[2] += c;
        state[3] += d;
        state[4] += e;
        state[5] += f;
        state[6] += g;
        state[7] += h;
    }

    size_t Size() const override { return 32; }

    size_t BlockSize() const override { return 64; }

    std::vector<uint8_t> Sum(const std::vector<uint8_t> &prefix = {}) override {
        auto     savedState = state;
        size_t   savedBufferSize = bufferSize;
        auto     savedBuffer = buffer;
        uint64_t savedLength = length;

        buffer[bufferSize++] = 0x80;
        if (bufferSize > 56) {
            std::fill(buffer.begin() + bufferSize, buffer.end(), 0);
            ProcessBlock();
            bufferSize = 0;
        }

        std::fill(buffer.begin() + bufferSize, buffer.begin() + 56, 0);
        uint64_t bitLength = length * 8;
        for (int i = 0; i < 8; i++) { buffer[63 - i] = static_cast<uint8_t>(bitLength >> (i * 8)); }
        ProcessBlock();

        std::vector<uint8_t> result(prefix);
        for (int i = 0; i < 8; i++) {
            uint32_t val = state[i];
            for (int j = 0; j < 4; j++) { result.push_back(static_cast<uint8_t>(val >> (24 - j * 8))); }
        }

        state = savedState;
        bufferSize = savedBufferSize;
        buffer = savedBuffer;
        length = savedLength;

        return result;
    }
};

struct hash {
    static std::unique_ptr<hash_Hash> NewAdler32() { return std::make_unique<hash_Adler32>(); }

    static std::unique_ptr<hash_Hash> NewCrc32() { return std::make_unique<hash_Crc32>(); }

    static std::unique_ptr<hash_Hash> NewFnv1_32() { return std::make_unique<hash_Fnv32>(0x01000193, 0x811C9DC5); }

    static std::unique_ptr<hash_Hash> NewFnv1a_32() { return std::make_unique<hash_Fnv32>(0x01000193, 0x811C9DC5); }

    static std::unique_ptr<hash_Hash> NewFnv1_64() { return std::make_unique<hash_Fnv64>(0x100000001B3, 0xCBF29CE484222325); }

    static std::unique_ptr<hash_Hash> NewFnv1a_64() { return std::make_unique<hash_Fnv64>(0x100000001B3, 0xCBF29CE484222325); }

    static std::unique_ptr<hash_Hash> NewSha256() { return std::make_unique<hash_Sha256>(); }

    static std::vector<uint8_t> Sum32(const std::string &data) {
        auto hasher = NewAdler32();
        hasher->WriteString(data);
        return hasher->Sum();
    }

    static std::vector<uint8_t> Sum64(const std::string &data) {
        auto hasher = NewFnv1_64();
        hasher->WriteString(data);
        return hasher->Sum();
    }

    static std::vector<uint8_t> Sum256(const std::string &data) {
        auto hasher = NewSha256();
        hasher->WriteString(data);
        return hasher->Sum();
    }
};

void testHashFunctions() {
    auto adler = hash::NewAdler32();
    adler->WriteString("Hello, world!");
    auto adlerSum = adler->Sum();
    printf("Adler32: ");
    for (auto b: adlerSum) printf("%02x", b);
    printf("\n");

    auto crc32 = hash::NewCrc32();
    crc32->WriteString("Hello, world!");
    auto crcSum = crc32->Sum();
    printf("CRC32: ");
    for (auto b: crcSum) printf("%02x", b);
    printf("\n");

    auto fnv64 = hash::NewFnv1a_64();
    fnv64->WriteString("Hello, world!");
    auto fnvSum = fnv64->Sum();
    printf("FNV-1a 64: ");
    for (auto b: fnvSum) printf("%02x", b);
    printf("\n");

    auto sha256 = hash::NewSha256();
    sha256->WriteString("Hello, world!");
    auto shaSum = sha256->Sum();
    printf("SHA256: ");
    for (auto b: shaSum) printf("%02x", b);
    printf("\n");

    auto quickSum = hash::Sum256("Hello, world!");
    printf("Quick SHA256: ");
    for (auto b: quickSum) printf("%02x", b);
    printf("\n");
}
