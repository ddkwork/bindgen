#pragma once
#include <cstdint>
#include <vector>
#include <array>
#include <stdexcept>
#include <cstring>
#include <algorithm>

struct aes_Error {
    std::string msg;

    aes_Error(const std::string &m) : msg(m) {}

    std::string Error() const { return msg; }
};

aes_Error aes_ErrKeySize = aes_Error("invalid key size");
aes_Error aes_ErrInputSize = aes_Error("input not full block");
aes_Error aes_ErrInvalidParam = aes_Error("invalid parameter");

struct aes_Cipher {
    int                   rounds;// 加密轮数
    std::vector<uint32_t> enc;   // 加密轮密钥
    std::vector<uint32_t> dec;   // 解密轮密钥

    void encryptBlock(std::array<uint8_t, 16> &dst, const std::array<uint8_t, 16> &src) const { encrypt(dst, src, enc); }

    void decryptBlock(std::array<uint8_t, 16> &dst, const std::array<uint8_t, 16> &src) const { decrypt(dst, src, dec); }

private:
    void encrypt(std::array<uint8_t, 16> &dst, const std::array<uint8_t, 16> &src, const std::vector<uint32_t> &rk) const {
        uint32_t s0 = load32(src.data()) ^ rk[0];
        uint32_t s1 = load32(src.data() + 4) ^ rk[1];
        uint32_t s2 = load32(src.data() + 8) ^ rk[2];
        uint32_t s3 = load32(src.data() + 12) ^ rk[3];

        for (int r = 1; r < rounds; r++) {
            uint32_t t0 = te0[getByte(s0, 3)] ^ te1[getByte(s1, 2)] ^ te2[getByte(s2, 1)] ^ te3[getByte(s3, 0)] ^ rk[4 * r];
            uint32_t t1 = te0[getByte(s1, 3)] ^ te1[getByte(s2, 2)] ^ te2[getByte(s3, 1)] ^ te3[getByte(s0, 0)] ^ rk[4 * r + 1];
            uint32_t t2 = te0[getByte(s2, 3)] ^ te1[getByte(s3, 2)] ^ te2[getByte(s0, 1)] ^ te3[getByte(s1, 0)] ^ rk[4 * r + 2];
            uint32_t t3 = te0[getByte(s3, 3)] ^ te1[getByte(s0, 2)] ^ te2[getByte(s1, 1)] ^ te3[getByte(s2, 0)] ^ rk[4 * r + 3];

            s0 = t0;
            s1 = t1;
            s2 = t2;
            s3 = t3;
        }

        uint32_t t0 =
                (te2[getByte(s0, 3)] & 0xff000000) ^ (te3[getByte(s1, 2)] & 0x00ff0000) ^ (te0[getByte(s2, 1)] & 0x0000ff00) ^ (te1[getByte(s3, 0)] & 0x000000ff) ^ rk[4 * rounds];

        uint32_t t1 = (te2[getByte(s1, 3)] & 0xff000000) ^ (te3[getByte(s2, 2)] & 0x00ff0000) ^ (te0[getByte(s3, 1)] & 0x0000ff00) ^ (te1[getByte(s0, 0)] & 0x000000ff) ^
                      rk[4 * rounds + 1];

        uint32_t t2 = (te2[getByte(s2, 3)] & 0xff000000) ^ (te3[getByte(s3, 2)] & 0x00ff0000) ^ (te0[getByte(s0, 1)] & 0x0000ff00) ^ (te1[getByte(s1, 0)] & 0x000000ff) ^
                      rk[4 * rounds + 2];

        uint32_t t3 = (te2[getByte(s3, 3)] & 0xff000000) ^ (te3[getByte(s0, 2)] & 0x00ff0000) ^ (te0[getByte(s1, 1)] & 0x0000ff00) ^ (te1[getByte(s2, 0)] & 0x000000ff) ^
                      rk[4 * rounds + 3];

        store32(dst.data(), t0);
        store32(dst.data() + 4, t1);
        store32(dst.data() + 8, t2);
        store32(dst.data() + 12, t3);
    }

    void decrypt(std::array<uint8_t, 16> &dst, const std::array<uint8_t, 16> &src, const std::vector<uint32_t> &rk) const {
        uint32_t s0 = load32(src.data()) ^ rk[0];
        uint32_t s1 = load32(src.data() + 4) ^ rk[1];
        uint32_t s2 = load32(src.data() + 8) ^ rk[2];
        uint32_t s3 = load32(src.data() + 12) ^ rk[3];

        for (int r = 1; r < rounds; r++) {
            uint32_t t0 = td0[getByte(s0, 3)] ^ td1[getByte(s3, 2)] ^ td2[getByte(s2, 1)] ^ td3[getByte(s1, 0)] ^ rk[4 * r];
            uint32_t t1 = td0[getByte(s1, 3)] ^ td1[getByte(s0, 2)] ^ td2[getByte(s3, 1)] ^ td3[getByte(s2, 0)] ^ rk[4 * r + 1];
            uint32_t t2 = td0[getByte(s2, 3)] ^ td1[getByte(s1, 2)] ^ td2[getByte(s0, 1)] ^ td3[getByte(s3, 0)] ^ rk[4 * r + 2];
            uint32_t t3 = td0[getByte(s3, 3)] ^ td1[getByte(s2, 2)] ^ td2[getByte(s1, 1)] ^ td3[getByte(s0, 0)] ^ rk[4 * r + 3];

            s0 = t0;
            s1 = t1;
            s2 = t2;
            s3 = t3;
        }

        uint32_t t0 =
                (td4[getByte(s0, 3)] & 0xff000000) ^ (td4[getByte(s3, 2)] & 0x00ff0000) ^ (td4[getByte(s2, 1)] & 0x0000ff00) ^ (td4[getByte(s1, 0)] & 0x000000ff) ^ rk[4 * rounds];

        uint32_t t1 = (td4[getByte(s1, 3)] & 0xff000000) ^ (td4[getByte(s0, 2)] & 0x00ff0000) ^ (td4[getByte(s3, 1)] & 0x0000ff00) ^ (td4[getByte(s2, 0)] & 0x000000ff) ^
                      rk[4 * rounds + 1];

        uint32_t t2 = (td4[getByte(s2, 3)] & 0xff000000) ^ (td4[getByte(s1, 2)] & 0x00ff0000) ^ (td4[getByte(s0, 1)] & 0x0000ff00) ^ (td4[getByte(s3, 0)] & 0x000000ff) ^
                      rk[4 * rounds + 2];

        uint32_t t3 = (td4[getByte(s3, 3)] & 0xff000000) ^ (td4[getByte(s2, 2)] & 0x00ff0000) ^ (td4[getByte(s1, 1)] & 0x0000ff00) ^ (td4[getByte(s0, 0)] & 0x000000ff) ^
                      rk[4 * rounds + 3];

        store32(dst.data(), t0);
        store32(dst.data() + 4, t1);
        store32(dst.data() + 8, t2);
        store32(dst.data() + 12, t3);
    }

    static uint32_t load32(const uint8_t *b) {
        return (static_cast<uint32_t>(b[0]) << 24) | (static_cast<uint32_t>(b[1]) << 16) | (static_cast<uint32_t>(b[2]) << 8) | static_cast<uint32_t>(b[3]);
    }

    static void store32(uint8_t *b, uint32_t w) {
        b[0] = static_cast<uint8_t>(w >> 24);
        b[1] = static_cast<uint8_t>(w >> 16);
        b[2] = static_cast<uint8_t>(w >> 8);
        b[3] = static_cast<uint8_t>(w);
    }

    static uint8_t getByte(uint32_t w, int n) { return static_cast<uint8_t>(w >> (24 - 8 * n)); }

    static const std::array<uint32_t, 256> te0;
    static const std::array<uint32_t, 256> te1;
    static const std::array<uint32_t, 256> te2;
    static const std::array<uint32_t, 256> te3;

    static const std::array<uint32_t, 256> td0;
    static const std::array<uint32_t, 256> td1;
    static const std::array<uint32_t, 256> td2;
    static const std::array<uint32_t, 256> td3;
    static const std::array<uint32_t, 256> td4;
};

struct aes_ECB {
    const aes_Cipher &cipher;

    explicit aes_ECB(const aes_Cipher &c) : cipher(c) {}

    std::vector<uint8_t> encrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 16 != 0) { throw aes_ErrInputSize; }

        std::vector<uint8_t> result;
        result.resize(data.size());

        for (size_t i = 0; i < data.size(); i += 16) {
            std::array<uint8_t, 16> src, dst;
            std::copy(data.begin() + i, data.begin() + i + 16, src.begin());
            cipher.encryptBlock(dst, src);
            std::copy(dst.begin(), dst.end(), result.begin() + i);
        }
        return result;
    }

    std::vector<uint8_t> decrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 16 != 0) { throw aes_ErrInputSize; }

        std::vector<uint8_t> result;
        result.resize(data.size());

        for (size_t i = 0; i < data.size(); i += 16) {
            std::array<uint8_t, 16> src, dst;
            std::copy(data.begin() + i, data.begin() + i + 16, src.begin());
            cipher.decryptBlock(dst, src);
            std::copy(dst.begin(), dst.end(), result.begin() + i);
        }
        return result;
    }
};

struct aes_CBC {
    const aes_Cipher       &cipher;
    std::array<uint8_t, 16> iv;

    aes_CBC(const aes_Cipher &c, const std::array<uint8_t, 16> &ivData) : cipher(c), iv(ivData) {}

    std::vector<uint8_t> encrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 16 != 0) { throw aes_ErrInputSize; }

        std::vector<uint8_t> result;
        result.resize(data.size());
        std::array<uint8_t, 16> prev = iv;

        for (size_t i = 0; i < data.size(); i += 16) {
            std::array<uint8_t, 16> src, dst;
            std::copy(data.begin() + i, data.begin() + i + 16, src.begin());

            for (int j = 0; j < 16; j++) { src[j] ^= prev[j]; }

            cipher.encryptBlock(dst, src);
            std::copy(dst.begin(), dst.end(), result.begin() + i);
            prev = dst;
        }
        return result;
    }

    std::vector<uint8_t> decrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 16 != 0) { throw aes_ErrInputSize; }

        std::vector<uint8_t> result;
        result.resize(data.size());
        std::array<uint8_t, 16> prev = iv;

        for (size_t i = 0; i < data.size(); i += 16) {
            std::array<uint8_t, 16> src, dst;
            std::copy(data.begin() + i, data.begin() + i + 16, src.begin());

            cipher.decryptBlock(dst, src);

            for (int j = 0; j < 16; j++) { dst[j] ^= prev[j]; }

            std::copy(dst.begin(), dst.end(), result.begin() + i);
            prev = src;
        }
        return result;
    }
};

struct aes {
    static aes_Cipher newCipher(const std::vector<uint8_t> &key) {
        switch (key.size()) {
            case 16:
                return expandKey(key, 10);
            case 24:
                return expandKey(key, 12);
            case 32:
                return expandKey(key, 14);
            default:
                throw aes_ErrKeySize;
        }
    }

    static aes_ECB newECBEncrypter(aes_Cipher &cipher) { return aes_ECB(cipher); }

    static aes_CBC newCBCEncrypter(aes_Cipher &cipher, const std::array<uint8_t, 16> &iv) { return aes_CBC(cipher, iv); }

    static aes_ECB newECBDecrypter(aes_Cipher &cipher) { return aes_ECB(cipher); }

    static aes_CBC newCBCDecrypter(aes_Cipher &cipher, const std::array<uint8_t, 16> &iv) { return aes_CBC(cipher, iv); }

    static std::vector<uint8_t> padPKCS7(const std::vector<uint8_t> &data, size_t blockSize = 16) {
        size_t               padding = blockSize - (data.size() % blockSize);
        std::vector<uint8_t> padded = data;
        padded.insert(padded.end(), padding, static_cast<uint8_t>(padding));
        return padded;
    }

    static std::vector<uint8_t> unpadPKCS7(const std::vector<uint8_t> &data) {
        if (data.empty()) return data;
        uint8_t padding = data.back();
        if (padding > data.size()) return data;
        return std::vector<uint8_t>(data.begin(), data.end() - padding);
    }

private:
    static aes_Cipher expandKey(const std::vector<uint8_t> &key, int rounds) {
        aes_Cipher cipher;
        cipher.rounds = rounds;

        int keyWords = key.size() / 4;
        int expandedSize = 4 * (rounds + 1);

        cipher.enc.resize(expandedSize);
        cipher.dec.resize(expandedSize);

        for (int i = 0; i < keyWords; i++) {
            cipher.enc[i] = (static_cast<uint32_t>(key[4 * i]) << 24) | (static_cast<uint32_t>(key[4 * i + 1]) << 16) | (static_cast<uint32_t>(key[4 * i + 2]) << 8) |
                            static_cast<uint32_t>(key[4 * i + 3]);
        }

        for (int i = keyWords; i < expandedSize; i++) {
            uint32_t temp = cipher.enc[i - 1];

            if (i % keyWords == 0) {
                temp = subWord(rotWord(temp)) ^ rcon[i / keyWords];
            } else if (keyWords > 6 && i % keyWords == 4) {
                temp = subWord(temp);
            }

            cipher.enc[i] = cipher.enc[i - keyWords] ^ temp;
        }

        for (int i = 0; i < expandedSize; i += 4) {
            uint32_t ei0 = cipher.enc[i];
            uint32_t ei1 = cipher.enc[i + 1];
            uint32_t ei2 = cipher.enc[i + 2];
            uint32_t ei3 = cipher.enc[i + 3];

            cipher.dec[i] = invMixColumn(invMixColumn(invMixColumn(ei0)));
            cipher.dec[i + 1] = invMixColumn(invMixColumn(invMixColumn(ei1)));
            cipher.dec[i + 2] = invMixColumn(invMixColumn(invMixColumn(ei2)));
            cipher.dec[i + 3] = invMixColumn(invMixColumn(invMixColumn(ei3)));
        }

        return cipher;
    }

    static uint32_t subWord(uint32_t w) {
        uint32_t res = 0;
        for (int i = 0; i < 4; i++) {
            uint8_t b = static_cast<uint8_t>(w >> (24 - 8 * i));
            res |= static_cast<uint32_t>(sbox[b]) << (24 - 8 * i);
        }
        return res;
    }

    static uint32_t rotWord(uint32_t w) { return (w << 8) | (w >> 24); }

    static uint32_t invMixColumn(uint32_t w) {
        uint32_t res = 0;
        for (int i = 0; i < 4; i++) {
            uint8_t b = static_cast<uint8_t>(w >> (24 - 8 * i));
            uint8_t b0 = mul(0x0e, b);
            uint8_t b1 = mul(0x0b, b);
            uint8_t b2 = mul(0x0d, b);
            uint8_t b3 = mul(0x09, b);
            res |= static_cast<uint32_t>(b0 ^ b1 ^ b2 ^ b3) << (24 - 8 * i);
        }
        return res;
    }

    static uint8_t mul(uint8_t a, uint8_t b) {
        uint8_t p = 0;
        for (int i = 0; i < 8; i++) {
            if (b & 1) p ^= a;
            bool hi = a & 0x80;
            a <<= 1;
            if (hi) a ^= 0x1b;// x^8 + x^4 + x^3 + x + 1
            b >>= 1;
        }
        return p;
    }

    static const std::array<uint8_t, 256> sbox;

    static const std::array<uint32_t, 11> rcon;
};

const std::array<uint8_t, 256> aes::sbox = {
        {0x63,
         0x7c,
         0x77,
         0x7b,
         0xf2,
         0x6b,
         0x6f,
         0xc5,
         0x30,
         0x01,
         0x67,
         0x2b,
         0xfe,
         0xd7,
         0xab,
         0x76,
         0xca,
         0x82,
         0xc9,
         0x7d,
         0xfa,
         0x59,
         0x47,
         0xf0,
         0xad,
         0xd4,
         0xa2,
         0xaf,
         0x9c,
         0xa4,
         0x72,
         0xc0,
         0xb7,
         0xfd,
         0x93,
         0x26,
         0x36,
         0x3f,
         0xf7,
         0xcc,
         0x34,
         0xa5,
         0xe5,
         0xf1,
         0x71,
         0xd8,
         0x31,
         0x15,
         0x04,
         0xc7,
         0x23,
         0xc3,
         0x18,
         0x96,
         0x05,
         0x9a,
         0x07,
         0x12,
         0x80,
         0xe2,
         0xeb,
         0x27,
         0xb2,
         0x75,
         0x09,
         0x83,
         0x2c,
         0x1a,
         0x1b,
         0x6e,
         0x5a,
         0xa0,
         0x52,
         0x3b,
         0xd6,
         0xb3,
         0x29,
         0xe3,
         0x2f,
         0x84,
         0x53,
         0xd1,
         0x00,
         0xed,
         0x20,
         0xfc,
         0xb1,
         0x5b,
         0x6a,
         0xcb,
         0xbe,
         0x39,
         0x4a,
         0x4c,
         0x58,
         0xcf,
         0xd0,
         0xef,
         0xaa,
         0xfb,
         0x43,
         0x4d,
         0x33,
         0x85,
         0x45,
         0xf9,
         0x02,
         0x7f,
         0x50,
         0x3c,
         0x9f,
         0xa8,
         0x51,
         0xa3,
         0x40,
         0x8f,
         0x92,
         0x9d,
         0x38,
         0xf5,
         0xbc,
         0xb6,
         0xda,
         0x21,
         0x10,
         0xff,
         0xf3,
         0xd2,
         0xcd,
         0x0c,
         0x13,
         0xec,
         0x5f,
         0x97,
         0x44,
         0x17,
         0xc4,
         0xa7,
         0x7e,
         0x3d,
         0x64,
         0x5d,
         0x19,
         0x73,
         0x60,
         0x81,
         0x4f,
         0xdc,
         0x22,
         0x2a,

         const std::array<uint32_t, 11> aes::rcon = {
                 {0x00000000, 0x01000000, 0x02000000, 0x04000000, 0x08000000, 0x10000000, 0x20000000, 0x40000000, 0x80000000, 0x1b000000, 0x36000000}};

const std::array<uint32_t, 256> aes_Cipher::te0 = {{
        0xc66363a5,
        0xf87c7c84,
        0xee777799,
        0xf67b7b8d,
        0xfff2f20d,
        0xd66b6bbd,
        0xde6f6fb1,
        0x91c5c554,
        0x60303050,
        0x02010103,
        0xce6767a9,
        0x562b2b7d,
}};

const std::array<uint32_t, 256> aes_Cipher::te1 = {{
        0xa5c66363,
        0x84f87c7c,
        0x99ee7777,
        0x8df67b7b,
        0x0dfff2f2,
        0xbdd66b6b,
        0xb1de6f6f,
        0x5491c5c5,
        0x50603030,
        0x03020101,
        0xa9ce6767,
        0x7d562b2b,
}};

const std::array<uint32_t, 256> aes_Cipher::te2 = {{
        0x63a5c663,
        0x7c84f87c,
        0x7799ee77,
        0x7b8df67b,
        0xf20dfff2,
        0x6bbdd66b,
        0x6fb1de6f,
        0xc55491c5,
        0x30506030,
        0x01030201,
        0x67a9ce67,
        0x2b7d562b,
}};

const std::array<uint32_t, 256> aes_Cipher::te3 = {{
        0x6363a5c6,
        0x7c7c84f8,
        0x777799ee,
        0x7b7b8df6,
        0xf2f20dff,
        0x6b6bbdd6,
        0x6f6fb1de,
        0xc5c55491,
        0x30305060,
        0x01010302,
        0x6767a9ce,
        0x2b2b7d56,
}};

const std::array<uint32_t, 256> aes_Cipher::td0 = {{
        0x51f4a750,
        0x7e416553,
        0x1a17a4c3,
        0x3a275e96,
        0x3bab6bcb,
        0x1f9d45f1,
        0xacfa58ab,
        0x4be30393,
        0x2030fa55,
        0xad766df6,
        0x88cc7691,
        0xf5024c25,
}};

const std::array<uint32_t, 256> aes_Cipher::td1 = {{
        0x50a7f451,
        0x5365417e,
        0xc3a4171a,
        0x965e273a,
        0xcb6bab3b,
        0xf1459d1f,
        0xab58faac,
        0x9303e34b,
        0x55fa3020,
        0xf66d76ad,
        0x9176cc88,
        0x254c02f5,
}};

const std::array<uint32_t, 256> aes_Cipher::td2 = {{
        0x50f4a751,
        0x5365417e,
        0xc3a4171a,
        0x965e273a,
        0xcb6bab3b,
        0xf1459d1f,
        0xab58faac,
        0x9303e34b,
        0x55fa3020,
        0xf66d76ad,
        0x9176cc88,
        0x254c02f5,
}};

const std::array<uint32_t, 256> aes_Cipher::td3 = {{
        0xa75051f4,
        0x65537e41,
        0xa4c31a17,
        0x5e963a27,
        0x6bcb3bab,
        0x45f11f9d,
        0x58abacfa,
        0x03934be3,
        0xfa552030,
        0x6df6ad76,
        0x769188cc,
        0x4c25f502,
}};

const std::array<uint32_t, 256> aes_Cipher::td4 = {{
        0x00000052,
        0x00000009,
        0x0000006a,
        0x000000d5,
        0x00000030,
        0x00000036,
        0x000000a5,
        0x00000038,
        0x000000bf,
        0x00000040,
        0x000000a3,
        0x0000009e,
}};

struct aes_CTR {
    const aes_Cipher       &cipher;
    std::array<uint8_t, 16> nonce;
    uint64_t                counter;

    aes_CTR(const aes_Cipher &c, const std::array<uint8_t, 12> &nonceData) : cipher(c), counter(0) { std::copy(nonceData.begin(), nonceData.end(), nonce.begin()); }

    std::vector<uint8_t> crypt(const std::vector<uint8_t> &data) {
        std::vector<uint8_t> result;
        result.resize(data.size());

        size_t pos = 0;
        while (pos < data.size()) {
            std::array<uint8_t, 16> keystream;
            generateKeystream(keystream);

            size_t blockSize = std::min<size_t>(16, data.size() - pos);
            for (size_t i = 0; i < blockSize; i++) { result[pos + i] = data[pos + i] ^ keystream[i]; }

            pos += blockSize;
        }
        return result;
    }

private:
    void generateKeystream(std::array<uint8_t, 16> &keystream) {
        std::array<uint8_t, 16> counterBlock;
        std::copy(nonce.begin(), nonce.end(), counterBlock.begin());

        for (int i = 0; i < 8; i++) { counterBlock[15 - i] = static_cast<uint8_t>(counter >> (8 * i)); }

        cipher.encryptBlock(keystream, counterBlock);

        counter++;
    }
};

struct aes_GCM {
    const aes_Cipher       &cipher;
    std::array<uint8_t, 12> nonce;

    aes_GCM(const aes_Cipher &c, const std::array<uint8_t, 12> &nonceData) : cipher(c), nonce(nonceData) {}

    struct AuthData {
        std::vector<uint8_t> data;
        std::vector<uint8_t> tag;
    };

    AuthData seal(const std::vector<uint8_t> &plaintext, const std::vector<uint8_t> &additionalData) {
        uint32_t                counter = 1;
        std::array<uint8_t, 16> counterBlock;
        std::copy(nonce.begin(), nonce.end(), counterBlock.begin());
        counterBlock[15] = 1;// 初始计数器

        std::vector<uint8_t> ciphertext;
        ciphertext.resize(plaintext.size());

        for (size_t i = 0; i < plaintext.size(); i += 16) {
            incrementCounter(counterBlock);

            std::array<uint8_t, 16> keystream;
            cipher.encryptBlock(keystream, counterBlock);

            size_t blockSize = std::min<size_t>(16, plaintext.size() - i);
            for (size_t j = 0; j < blockSize; j++) { ciphertext[i + j] = plaintext[i + j] ^ keystream[j]; }
        }

        std::array<uint8_t, 16> tag = computeTag(additionalData, ciphertext);

        return {ciphertext, std::vector<uint8_t>(tag.begin(), tag.end())};
    }

    std::vector<uint8_t> open(const AuthData &authData, const std::vector<uint8_t> &additionalData) {
        std::array<uint8_t, 16> computedTag = computeTag(additionalData, authData.data);
        if (!std::equal(computedTag.begin(), computedTag.end(), authData.tag.begin())) { throw aes_Error("GCM tag verification failed"); }

        std::vector<uint8_t> plaintext;
        plaintext.resize(authData.data.size());

        std::array<uint8_t, 16> counterBlock;
        std::copy(nonce.begin(), nonce.end(), counterBlock.begin());
        counterBlock[15] = 1;// 初始计数器

        for (size_t i = 0; i < authData.data.size(); i += 16) {
            incrementCounter(counterBlock);

            std::array<uint8_t, 16> keystream;
            cipher.encryptBlock(keystream, counterBlock);

            size_t blockSize = std::min<size_t>(16, authData.data.size() - i);
            for (size_t j = 0; j < blockSize; j++) { plaintext[i + j] = authData.data[i + j] ^ keystream[j]; }
        }

        return plaintext;
    }

private:
    void incrementCounter(std::array<uint8_t, 16> &block) {
        for (int i = 15; i >= 12; i--) {
            if (++block[i] != 0) break;
        }
    }

    std::array<uint8_t, 16> computeTag(const std::vector<uint8_t> &aad, const std::vector<uint8_t> &ciphertext) {
        std::array<uint8_t, 16> hash = {0};

        ghash(hash, aad);

        ghash(hash, ciphertext);

        uint64_t                aadLen = aad.size() * 8;
        uint64_t                ctLen = ciphertext.size() * 8;
        std::array<uint8_t, 16> lenBlock;
        store64(lenBlock.data(), aadLen);
        store64(lenBlock.data() + 8, ctLen);
        ghash(hash, std::vector<uint8_t>(lenBlock.begin(), lenBlock.end()));

        std::array<uint8_t, 16> encHash;
        std::array<uint8_t, 16> j0;
        std::copy(nonce.begin(), nonce.end(), j0.begin());
        j0[15] = 1;
        cipher.encryptBlock(encHash, j0);

        for (int i = 0; i < 16; i++) { encHash[i] ^= hash[i]; }

        return encHash;
    }

    void ghash(std::array<uint8_t, 16> &hash, const std::vector<uint8_t> &data) {
        for (size_t i = 0; i < data.size(); i += 16) {
            for (int j = 0; j < 16 && (i + j) < data.size(); j++) { hash[j] ^= data[i + j]; }
        }
    }

    static void store64(uint8_t *dst, uint64_t n) {
        for (int i = 0; i < 8; i++) {
            dst[7 - i] = static_cast<uint8_t>(n);
            n >>= 8;
        }
    }
};

struct aes {
    static aes_Cipher newCipher(const std::vector<uint8_t> &key) {}

    static aes_ECB newECBEncrypter(aes_Cipher &cipher) { return aes_ECB(cipher); }

    static aes_ECB newECBDecrypter(aes_Cipher &cipher) { return aes_ECB(cipher); }

    static aes_CBC newCBCEncrypter(aes_Cipher &cipher, const std::array<uint8_t, 16> &iv) { return aes_CBC(cipher, iv); }

    static aes_CBC newCBCDecrypter(aes_Cipher &cipher, const std::array<uint8_t, 16> &iv) { return aes_CBC(cipher, iv); }

    static aes_CTR newCTR(aes_Cipher &cipher, const std::array<uint8_t, 12> &nonce) { return aes_CTR(cipher, nonce); }

    static aes_GCM newGCM(aes_Cipher &cipher, const std::array<uint8_t, 12> &nonce) { return aes_GCM(cipher, nonce); }

    static std::vector<uint8_t> padPKCS7(const std::vector<uint8_t> &data, size_t blockSize = 16) {}

    static std::vector<uint8_t> unpadPKCS7(const std::vector<uint8_t> &data) {}
};

void testAESFunctions() {
    std::vector<uint8_t>    key = {0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c};
    std::vector<uint8_t>    plaintext = {'H', 'e', 'l', 'l', 'o', ' ', 'A', 'E', 'S', '!', '!', '!', '!', '!', '!', '!'};
    std::array<uint8_t, 16> iv = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f};

    auto aesCipher = aes::newCipher(key);

    auto ecbEnc = aes::newECBEncrypter(aesCipher);
    auto encrypted = ecbEnc.encrypt(plaintext);
    auto ecbDec = aes::newECBDecrypter(aesCipher);
    auto decrypted = ecbDec.decrypt(encrypted);

    auto cbcEnc = aes::newCBCEncrypter(aesCipher, iv);
    encrypted = cbcEnc.encrypt(plaintext);
    auto cbcDec = aes::newCBCDecrypter(aesCipher, iv);
    decrypted = cbcDec.decrypt(encrypted);

    std::vector<uint8_t> key256(32, 0xaa);
    auto                 aesCipher256 = aes::newCipher(key256);

    std::array<uint8_t, 12> nonce = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b};
    auto                    ctr = aes::newCTR(aesCipher, nonce);
    encrypted = ctr.crypt(plaintext);
    decrypted = ctr.crypt(encrypted);// CTR是对称的

    auto                 gcm = aes::newGCM(aesCipher, nonce);
    std::vector<uint8_t> aad = {0x50, 0x51, 0x52, 0x53};// 附加认证数据
    auto                 sealed = gcm.seal(plaintext, aad);
    decrypted = gcm.open(sealed, aad);

    std::string          longText = "This is a longer message that needs padding for CBC mode!";
    std::vector<uint8_t> longData(longText.begin(), longText.end());

    auto padded = aes::padPKCS7(longData);

    encrypted = cbcEnc.encrypt(padded);

    decrypted = cbcDec.decrypt(encrypted);

    auto        unpadded = aes::unpadPKCS7(decrypted);
    std::string result(unpadded.begin(), unpadded.end());
}
