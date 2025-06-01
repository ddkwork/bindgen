#pragma once
#include <cstdint>
#include <vector>
#include <array>
#include <string>
#include <cstring>
#include <functional>
#include <stdexcept>

struct HashInterface {
    virtual ~HashInterface() = default;
    virtual void                 reset() = 0;
    virtual void                 write(const uint8_t *data, size_t size) = 0;
    virtual std::vector<uint8_t> sum() = 0;
    virtual size_t               size() const = 0;
    virtual size_t               blockSize() const = 0;
};

struct md5_Hash : public HashInterface {};

struct sha1_Hash : public HashInterface {
    void reset() override { /* 实现重置逻辑 */ }

    void write(const uint8_t *data, size_t size) override { /* 实现写入逻辑 */ }

    std::vector<uint8_t> sum() override { /* 实现摘要计算 */ return {}; }

    size_t size() const override { return 20; }// 160位

    size_t blockSize() const override { return 64; }
};

struct sha256_Hash : public HashInterface {
    void reset() override { /* 实现重置逻辑 */ }

    void write(const uint8_t *data, size_t size) override { /* 实现写入逻辑 */ }

    std::vector<uint8_t> sum() override { /* 实现摘要计算 */ return {}; }

    size_t size() const override { return 32; }// 256位

    size_t blockSize() const override { return 64; }
};

template<typename Hash>
struct hmac_Hash : public HashInterface {
    Hash   inner;
    Hash   outer;
    bool   finalized;
    size_t blockSize;

    hmac_Hash(const std::vector<uint8_t> &key) : finalized(false) {
        blockSize = inner.blockSize();

        std::vector<uint8_t> processedKey = key;
        if (processedKey.size() > blockSize) {
            inner.reset();
            inner.write(processedKey.data(), processedKey.size());
            processedKey = inner.sum();
        }

        processedKey.resize(blockSize, 0x00);

        std::vector<uint8_t> ipad(blockSize, 0x36);
        std::vector<uint8_t> opad(blockSize, 0x5c);

        for (size_t i = 0; i < blockSize; i++) {
            ipad[i] ^= processedKey[i];
            opad[i] ^= processedKey[i];
        }

        inner.reset();
        inner.write(ipad.data(), ipad.size());
    }

    void reset() override { throw std::runtime_error("hmac: reset not supported in this implementation"); }

    void write(const uint8_t *data, size_t size) override {
        if (finalized) { throw std::runtime_error("hmac: write after sum"); }
        inner.write(data, size);
    }

    std::vector<uint8_t> sum() override {
        if (finalized) { throw std::runtime_error("hmac: sum called multiple times"); }
        finalized = true;

        std::vector<uint8_t> innerDigest = inner.sum();

        outer.reset();
        outer.write(outerPad().data(), blockSize);
        outer.write(innerDigest.data(), innerDigest.size());

        return outer.sum();
    }

    size_t size() const override { return outer.size(); }

    size_t blockSize() const override { return inner.blockSize(); }

private:
    std::vector<uint8_t> outerPad() const {
        std::vector<uint8_t> opad(blockSize, 0x5c);
        return opad;
    }
};

struct hmac {
    template<typename Hash>
    static hmac_Hash<Hash> New(const std::vector<uint8_t> &key) {
        return hmac_Hash<Hash>(key);
    }

    template<typename Hash>
    static std::vector<uint8_t> Sum(const std::vector<uint8_t> &key, const std::vector<uint8_t> &data) {
        hmac_Hash<Hash> h(key);
        h.write(data.data(), data.size());
        return h.sum();
    }

    template<typename Hash>
    static std::vector<uint8_t> Sum(const std::vector<uint8_t> &key, const std::string &data) {
        return Sum<Hash>(key, std::vector<uint8_t>(data.begin(), data.end()));
    }

    template<typename Hash>
    static std::vector<uint8_t> File(const std::vector<uint8_t> &key, const std::string &filename) {
        hmac_Hash<Hash> h(key);
        FILE           *file = fopen(filename.c_str(), "rb");
        if (!file) { throw std::runtime_error("cannot open file: " + filename); }

        uint8_t buffer[4096];
        size_t  bytesRead;

        while ((bytesRead = fread(buffer, 1, sizeof(buffer), file)) > 0) { h.write(buffer, bytesRead); }

        fclose(file);
        return h.sum();
    }

    static bool Equal(const std::vector<uint8_t> &a, const std::vector<uint8_t> &b) {
        if (a.size() != b.size()) { return false; }

        uint8_t result = 0;
        for (size_t i = 0; i < a.size(); i++) { result |= a[i] ^ b[i]; }

        return result == 0;
    }
};

using HMAC_MD5 = hmac_Hash<md5_Hash>;
using HMAC_SHA1 = hmac_Hash<sha1_Hash>;
using HMAC_SHA256 = hmac_Hash<sha256_Hash>;

void testHMACFunctions() {
    std::vector<uint8_t> key = {'k', 'e', 'y'};
    std::string          data = "Hello, HMAC!";

    auto md5Mac = hmac::Sum<md5_Hash>(key, data);

    HMAC_MD5 md5Hasher(key);
    md5Hasher.write(reinterpret_cast<const uint8_t *>(data.data()), data.size());
    auto md5Digest = md5Hasher.sum();

    std::vector<uint8_t> longKey(32, 0xaa);// 32字节密钥
    std::vector<uint8_t> binaryData = {0x00, 0x01, 0x02, 0x03};

    HMAC_SHA256 sha256Hasher(longKey);
    sha256Hasher.write(binaryData.data(), binaryData.size());
    auto sha256Digest = sha256Hasher.sum();

    try {
        auto fileMac = hmac::File<sha1_Hash>(key, "document.txt");
        std::cout << "HMAC-SHA1(file) = ";
        for (auto b: fileMac) { printf("%02x", b); }
        std::cout << std::endl;
    } catch (const std::exception &e) { std::cerr << "Error: " << e.what() << std::endl; }

    std::vector<uint8_t> expectedMac = {0x75, 0x0c, 0x78, 0x3e, 0x6a, 0xb0, 0xb5, 0x03, 0xea, 0xa8, 0x6e, 0x31, 0x0a, 0x5d, 0xb7, 0x38};

    if (hmac::Equal(md5Digest, expectedMac)) {
        std::cout << "HMAC-MD5 verification succeeded" << std::endl;
    } else {
        std::cout << "HMAC-MD5 verification failed" << std::endl;
    }

    struct TestVector {
        std::vector<uint8_t> key;
        std::vector<uint8_t> data;
        std::vector<uint8_t> expected;
    };

    TestVector md5Vectors[] = { {{0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b, 0x0b}, {0x48, 0x69, 0
