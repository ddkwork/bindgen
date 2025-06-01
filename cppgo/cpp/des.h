#pragma once
#include <cstdint>
#include <vector>
#include <array>
#include <stdexcept>
#include <cstring>
#include <algorithm>

struct des_Error {
    std::string msg;

    des_Error(const std::string &m) : msg(m) {}

    std::string Error() const { return msg; }
};

des_Error des_ErrKeySize = des_Error("invalid key size");
des_Error des_ErrInputSize = des_Error("input not full block");
des_Error des_ErrInvalidParam = des_Error("invalid parameter");

struct des_Cipher {
    std::array<uint64_t, 16> subkeys;

    uint64_t encryptBlock(uint64_t block) const { return permute(block, false); }

    uint64_t decryptBlock(uint64_t block) const { return permute(block, true); }

private:
    uint64_t permute(uint64_t block, bool decrypt) const {
        block = initialPermutation(block);

        uint32_t left = static_cast<uint32_t>(block >> 32);
        uint32_t right = static_cast<uint32_t>(block);

        for (int i = 0; i < 16; i++) {
            int      round = decrypt ? 15 - i : i;
            uint32_t nextLeft = right;
            uint32_t feistel = feistelFunction(right, subkeys[round]);
            right = left ^ feistel;
            left = nextLeft;
        }

        block = (static_cast<uint64_t>(right) << 32) | left;

        return finalPermutation(block);
    }

    uint32_t feistelFunction(uint32_t r, uint64_t key) const {
        uint64_t expanded = expansionPermutation(r);

        expanded ^= key;

        uint32_t substituted = sBoxSubstitution(expanded);

        return pBoxPermutation(substituted);
    }

    static constexpr std::array<int, 64> initialPermutationTable = {58, 50, 42, 34, 26, 18, 10, 2,  60, 52, 44, 36, 28, 20, 12, 4,  62, 54, 46, 38, 30, 22,
                                                                    14, 6,  64, 56, 48, 40, 32, 24, 16, 8,  57, 49, 41, 33, 25, 17, 9,  1,  59, 51, 43, 35,
                                                                    27, 19, 11, 3,  61, 53, 45, 37, 29, 21, 13, 5,  63, 55, 47, 39, 31, 23, 15, 7};

    uint64_t initialPermutation(uint64_t block) const { return permuteBlock(block, initialPermutationTable); }

    static constexpr std::array<int, 64> finalPermutationTable = {40, 8,  48, 16, 56, 24, 64, 32, 39, 7,  47, 15, 55, 23, 63, 31, 38, 6,  46, 14, 54, 22,
                                                                  62, 30, 37, 5,  45, 13, 53, 21, 61, 29, 36, 4,  44, 12, 52, 20, 60, 28, 35, 3,  43, 11,
                                                                  51, 19, 59, 27, 34, 2,  42, 10, 50, 18, 58, 26, 33, 1,  41, 9,  49, 17, 57, 25};

    uint64_t finalPermutation(uint64_t block) const { return permuteBlock(block, finalPermutationTable); }

    static constexpr std::array<int, 48> expansionPermutationTable = {32, 1,  2,  3,  4,  5,  4,  5,  6,  7,  8,  9,  8,  9,  10, 11, 12, 13, 12, 13, 14, 15, 16, 17,
                                                                      16, 17, 18, 19, 20, 21, 20, 21, 22, 23, 24, 25, 24, 25, 26, 27, 28, 29, 28, 29, 30, 31, 32, 1};

    uint64_t expansionPermutation(uint32_t r) const {
        uint64_t result = 0;
        for (int i = 0; i < 48; i++) {
            int      pos = expansionPermutationTable[i] - 1;
            uint64_t bit = (r >> (31 - pos)) & 1;
            result |= (bit << (47 - i));
        }
        return result;
    }

    static constexpr std::array<std::array<uint8_t, 64>, 8> sBoxes = {
            {{// S1
              14, 4, 13, 1, 2,  15, 11, 8,  3,  10, 6, 12, 5, 9,  0, 7, 0,  15, 7, 4, 14, 2, 13, 1, 10, 6,  12, 11, 9,  5, 3, 8,
              4,  1, 14, 8, 13, 6,  2,  11, 15, 12, 9, 7,  3, 10, 5, 0, 15, 12, 8, 2, 4,  9, 1,  7, 5,  11, 3,  14, 10, 0, 6, 13},
             {// S2
              15, 1,  8, 14, 6,  11, 3,  4, 9, 7, 2,  13, 12, 0, 5, 10, 3,  13, 4,  7, 15, 2,  8, 14, 12, 0, 1, 10, 6, 9, 11, 5,
              0,  14, 7, 11, 10, 4,  13, 1, 5, 8, 12, 6,  9,  3, 2, 15, 13, 8,  10, 1, 3,  15, 4, 2,  11, 6, 7, 12, 0, 5, 14, 9},
             {// S3
              10, 0, 9, 14, 6, 3,  15, 5, 1,  13, 12, 7,  11, 4,  2,  8, 13, 7,  0,  9, 3, 4, 6, 10, 2, 8,  5,  14, 12, 11, 15, 1,
              13, 6, 4, 9,  8, 15, 3,  0, 11, 1,  2,  12, 5,  10, 14, 7, 1,  10, 13, 0, 6, 9, 8, 7,  4, 15, 14, 3,  11, 5,  2,  12},
             {// S4
              7,  13, 14, 3, 0,  6,  9, 10, 1,  2, 8, 5,  11, 12, 4, 15, 13, 8,  11, 5, 6,  15, 0,  3, 4, 7, 2, 12, 1,  10, 14, 9,
              10, 6,  9,  0, 12, 11, 7, 13, 15, 1, 3, 14, 5,  2,  8, 4,  3,  15, 0,  6, 10, 1,  13, 8, 9, 4, 5, 11, 12, 7,  2,  14},
             {// S5
              2, 12, 4, 1,  7,  10, 11, 6, 8,  5, 3,  15, 13, 0, 14, 9,  14, 11, 2,  12, 4, 7,  13, 1,  5, 0,  15, 10, 3,  9, 8, 6,
              4, 2,  1, 11, 10, 13, 7,  8, 15, 9, 12, 5,  6,  3, 0,  14, 11, 8,  12, 7,  1, 14, 2,  13, 6, 15, 0,  9,  10, 4, 5, 3},
             {// S6
              12, 1,  10, 15, 9, 2, 6,  8, 0, 13, 3, 4,  14, 7,  5,  11, 10, 15, 4, 2,  7, 12, 9,  5,  6,  1,  13, 14, 0, 11, 3, 8,
              9,  14, 15, 5,  2, 8, 12, 3, 7, 0,  4, 10, 1,  13, 11, 6,  4,  3,  2, 12, 9, 5,  15, 10, 11, 14, 1,  7,  6, 0,  8, 13},
             {// S7
              4, 11, 2,  14, 15, 0, 8, 13, 3,  12, 9, 7, 5, 10, 6, 1, 13, 0,  11, 7, 4, 9, 1,  10, 14, 3, 5, 12, 2,  15, 8, 6,
              1, 4,  11, 13, 12, 3, 7, 14, 10, 15, 6, 8, 0, 5,  9, 2, 6,  11, 13, 8, 1, 4, 10, 7,  9,  5, 0, 15, 14, 2,  3, 12},
             {// S8
              13, 2,  8, 4, 6, 15, 11, 1, 10, 9, 3,  14, 5,  0, 12, 7, 1, 15, 13, 8, 10, 3,  7, 4,  12, 5,  6, 11, 0, 14, 9, 2,
              7,  11, 4, 1, 9, 12, 14, 2, 0,  6, 10, 13, 15, 3, 5,  8, 2, 1,  14, 7, 4,  10, 8, 13, 15, 12, 9, 0,  3, 5,  6, 11}}};

    uint32_t sBoxSubstitution(uint64_t value) const {
        uint32_t result = 0;
        for (int i = 0; i < 8; i++) {
            uint8_t input = (value >> (42 - i * 6)) & 0x3F;

            uint8_t row = ((input & 0x20) >> 4) | (input & 0x01);
            uint8_t col = (input >> 1) & 0x0F;

            uint8_t sboxValue = sBoxes[i][row * 16 + col];

            result |= (static_cast<uint32_t>(sboxValue) << (28 - i * 4));
        }
        return result;
    }

    static constexpr std::array<int, 32> pBoxPermutationTable = {16, 7, 20, 21, 29, 12, 28, 17, 1,  15, 23, 26, 5,  18, 31, 10,
                                                                 2,  8, 24, 14, 32, 27, 3,  9,  19, 13, 30, 6,  22, 11, 4,  25};

    uint32_t pBoxPermutation(uint32_t value) const {
        uint32_t result = 0;
        for (int i = 0; i < 32; i++) {
            int      pos = pBoxPermutationTable[i] - 1;
            uint32_t bit = (value >> (31 - pos)) & 1;
            result |= (bit << (31 - i));
        }
        return result;
    }

    uint64_t permuteBlock(uint64_t block, const std::array<int, 64> &table) const {
        uint64_t result = 0;
        for (int i = 0; i < 64; i++) {
            int      pos = table[i] - 1;
            uint64_t bit = (block >> (63 - pos)) & 1;
            result |= (bit << (63 - i));
        }
        return result;
    }
};

struct tripleDES_Cipher {
    des_Cipher cipher1;
    des_Cipher cipher2;
    des_Cipher cipher3;

    uint64_t encryptBlock(uint64_t block) const {
        block = cipher1.encryptBlock(block);
        block = cipher2.decryptBlock(block);
        return cipher3.encryptBlock(block);
    }

    uint64_t decryptBlock(uint64_t block) const {
        block = cipher3.decryptBlock(block);
        block = cipher2.encryptBlock(block);
        return cipher1.decryptBlock(block);
    }
};

struct des {
    static constexpr std::array<int, 56> pc1Table = {57, 49, 41, 33, 25, 17, 9,  1, 58, 50, 42, 34, 26, 18, 10, 2, 59, 51, 43, 35, 27, 19, 11, 3, 60, 52, 44, 36,
                                                     63, 55, 47, 39, 31, 23, 15, 7, 62, 54, 46, 38, 30, 22, 14, 6, 61, 53, 45, 37, 29, 21, 13, 5, 28, 20, 12, 4};

    static constexpr std::array<int, 48> pc2Table = {14, 17, 11, 24, 1,  5,  3,  28, 15, 6,  21, 10, 23, 19, 12, 4,  26, 8,  16, 7,  27, 20, 13, 2,
                                                     41, 52, 31, 37, 47, 55, 30, 40, 51, 45, 33, 48, 44, 49, 39, 56, 34, 53, 46, 42, 50, 36, 29, 32};

    static constexpr std::array<int, 16> shiftTable = {1, 1, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 1};

    static des_Cipher newCipher(const std::vector<uint8_t> &key) {
        if (key.size() != 8) { throw des_ErrKeySize; }

        des_Cipher cipher;
        generateSubkeys(cipher.subkeys, key);
        return cipher;
    }

    static tripleDES_Cipher newTripleDESCipher(const std::vector<uint8_t> &key) {
        if (key.size() != 24) { throw des_ErrKeySize; }

        tripleDES_Cipher cipher;

        std::vector<uint8_t> key1(key.begin(), key.begin() + 8);
        std::vector<uint8_t> key2(key.begin() + 8, key.begin() + 16);
        std::vector<uint8_t> key3(key.begin() + 16, key.end());

        cipher.cipher1 = newCipher(key1);
        cipher.cipher2 = newCipher(key2);
        cipher.cipher3 = newCipher(key3);

        return cipher;
    }

private:
    static void generateSubkeys(std::array<uint64_t, 16> &subkeys, const std::vector<uint8_t> &key) {
        uint64_t key64 = 0;
        for (int i = 0; i < 8; i++) { key64 = (key64 << 8) | key[i]; }

        uint64_t permuted = permuteKey(key64, pc1Table);

        uint32_t left = static_cast<uint32_t>(permuted >> 28);
        uint32_t right = static_cast<uint32_t>(permuted);

        for (int i = 0; i < 16; i++) {
            left = rotateLeft28(left, shiftTable[i]);
            right = rotateLeft28(right, shiftTable[i]);

            uint64_t combined = (static_cast<uint64_t>(left) << 28) | right;

            subkeys[i] = permuteKey(combined, pc2Table);
        }
    }

    static uint32_t rotateLeft28(uint32_t value, int shift) { return ((value << shift) | (value >> (28 - shift))) & 0x0FFFFFFF; }

    static uint64_t permuteKey(uint64_t key, const std::array<int, 56> &table) {
        uint64_t result = 0;
        for (int i = 0; i < 56; i++) {
            int      pos = table[i] - 1;
            uint64_t bit = (key >> (63 - pos)) & 1;
            result |= (bit << (55 - i));
        }
        return result;
    }

    static uint64_t permuteKey(uint64_t key, const std::array<int, 48> &table) {
        uint64_t result = 0;
        for (int i = 0; i < 48; i++) {
            int      pos = table[i] - 1;
            uint64_t bit = (key >> (55 - pos)) & 1;// PC-2使用56位输入
            result |= (bit << (47 - i));
        }
        return result;
    }
};

struct des_ECB {
    const des_Cipher &cipher;

    explicit des_ECB(const des_Cipher &c) : cipher(c) {}

    std::vector<uint8_t> encrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 8 != 0) { throw des_ErrInputSize; }

        std::vector<uint8_t> result;
        for (size_t i = 0; i < data.size(); i += 8) {
            uint64_t block = readBlock(data, i);
            block = cipher.encryptBlock(block);
            writeBlock(result, block);
        }
        return result;
    }

    std::vector<uint8_t> decrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 8 != 0) { throw des_ErrInputSize; }

        std::vector<uint8_t> result;
        for (size_t i = 0; i < data.size(); i += 8) {
            uint64_t block = readBlock(data, i);
            block = cipher.decryptBlock(block);
            writeBlock(result, block);
        }
        return result;
    }

private:
    static uint64_t readBlock(const std::vector<uint8_t> &data, size_t offset) {
        uint64_t block = 0;
        for (int i = 0; i < 8; i++) { block = (block << 8) | data[offset + i]; }
        return block;
    }

    static void writeBlock(std::vector<uint8_t> &out, uint64_t block) {
        for (int i = 0; i < 8; i++) { out.push_back(static_cast<uint8_t>(block >> (56 - i * 8))); }
    }
};

struct des_CBC {
    const des_Cipher &cipher;
    uint64_t          iv;

    des_CBC(const des_Cipher &c, const std::vector<uint8_t> &ivData) : cipher(c) {
        if (ivData.size() != 8) { throw des_ErrInputSize; }
        iv = des_ECB::readBlock(ivData, 0);
    }

    std::vector<uint8_t> encrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 8 != 0) { throw des_ErrInputSize; }

        std::vector<uint8_t> result;
        uint64_t             prev = iv;

        for (size_t i = 0; i < data.size(); i += 8) {
            uint64_t block = des_ECB::readBlock(data, i);
            block ^= prev;
            block = cipher.encryptBlock(block);
            des_ECB::writeBlock(result, block);
            prev = block;
        }
        return result;
    }

    std::vector<uint8_t> decrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 8 != 0) { throw des_ErrInputSize; }

        std::vector<uint8_t> result;
        uint64_t             prev = iv;

        for (size_t i = 0; i < data.size(); i += 8) {
            uint64_t block = des_ECB::readBlock(data, i);
            uint64_t decrypted = cipher.decryptBlock(block);
            decrypted ^= prev;
            des_ECB::writeBlock(result, decrypted);
            prev = block;
        }
        return result;
    }
};

struct des {
    static des_Cipher newCipher(const std::vector<uint8_t> &key) { return des::newCipher(key); }

    static tripleDES_Cipher newTripleDESCipher(const std::vector<uint8_t> &key) { return des::newTripleDESCipher(key); }

    static des_ECB newECBEncrypter(des_Cipher &cipher) { return des_ECB(cipher); }

    static des_CBC newCBCEncrypter(des_Cipher &cipher, const std::vector<uint8_t> &iv) { return des_CBC(cipher, iv); }

    static des_ECB newECBDecrypter(des_Cipher &cipher) { return des_ECB(cipher); }

    static des_CBC newCBCDecrypter(des_Cipher &cipher, const std::vector<uint8_t> &iv) { return des_CBC(cipher, iv); }

    static std::vector<uint8_t> padPKCS7(const std::vector<uint8_t> &data, size_t blockSize) {
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
};

struct tripleDES_ECB {
    const tripleDES_Cipher &cipher;

    explicit tripleDES_ECB(const tripleDES_Cipher &c) : cipher(c) {}

    std::vector<uint8_t> encrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 8 != 0) { throw des_ErrInputSize; }

        std::vector<uint8_t> result;
        for (size_t i = 0; i < data.size(); i += 8) {
            uint64_t block = des_ECB::readBlock(data, i);
            block = cipher.encryptBlock(block);
            des_ECB::writeBlock(result, block);
        }
        return result;
    }

    std::vector<uint8_t> decrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 8 != 0) { throw des_ErrInputSize; }

        std::vector<uint8_t> result;
        for (size_t i = 0; i < data.size(); i += 8) {
            uint64_t block = des_ECB::readBlock(data, i);
            block = cipher.decryptBlock(block);
            des_ECB::writeBlock(result, block);
        }
        return result;
    }
};

struct tripleDES_CBC {
    const tripleDES_Cipher &cipher;
    uint64_t                iv;

    tripleDES_CBC(const tripleDES_Cipher &c, const std::vector<uint8_t> &ivData) : cipher(c) {
        if (ivData.size() != 8) { throw des_ErrInputSize; }
        iv = des_ECB::readBlock(ivData, 0);
    }

    std::vector<uint8_t> encrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 8 != 0) { throw des_ErrInputSize; }

        std::vector<uint8_t> result;
        uint64_t             prev = iv;

        for (size_t i = 0; i < data.size(); i += 8) {
            uint64_t block = des_ECB::readBlock(data, i);
            block ^= prev;
            block = cipher.encryptBlock(block);
            des_ECB::writeBlock(result, block);
            prev = block;
        }
        return result;
    }

    std::vector<uint8_t> decrypt(const std::vector<uint8_t> &data) const {
        if (data.size() % 8 != 0) { throw des_ErrInputSize; }

        std::vector<uint8_t> result;
        uint64_t             prev = iv;

        for (size_t i = 0; i < data.size(); i += 8) {
            uint64_t block = des_ECB::readBlock(data, i);
            uint64_t decrypted = cipher.decryptBlock(block);
            decrypted ^= prev;
            des_ECB::writeBlock(result, decrypted);
            prev = block;
        }
        return result;
    }
};

void testDESFunctions() {
    std::vector<uint8_t> key = {0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF};
    std::vector<uint8_t> plaintext = {'H', 'e', 'l', 'l', 'o', '!', '!', '!'};
    std::vector<uint8_t> iv = {0x12, 0x34, 0x56, 0x78, 0x90, 0xAB, 0xCD, 0xEF};

    auto desCipher = des::newCipher(key);

    auto ecbEnc = des::newECBEncrypter(desCipher);
    auto encrypted = ecbEnc.encrypt(plaintext);
    auto ecbDec = des::newECBDecrypter(desCipher);
    auto decrypted = ecbDec.decrypt(encrypted);

    auto cbcEnc = des::newCBCEncrypter(desCipher, iv);
    encrypted = cbcEnc.encrypt(plaintext);
    auto cbcDec = des::newCBCDecrypter(desCipher, iv);
    decrypted = cbcDec.decrypt(encrypted);

    std::vector<uint8_t> key3 = {0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0xFE, 0xDC, 0xBA, 0x98, 0x76, 0x54, 0x32, 0x10, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67};

    auto des3Cipher = des::newTripleDESCipher(key3);

    tripleDES_ECB tdesECB(des3Cipher);
    encrypted = tdesECB.encrypt(plaintext);
    decrypted = tdesECB.decrypt(encrypted);

    tripleDES_CBC tdesCBC(des3Cipher, iv);
    encrypted = tdesCBC.encrypt(plaintext);
    decrypted = tdesCBC.decrypt(encrypted);

    std::string          longText = "This is a longer message that needs padding!";
    std::vector<uint8_t> longData(longText.begin(), longText.end());

    auto padded = des::padPKCS7(longData, 8);

    encrypted = cbcEnc.encrypt(padded);

    decrypted = cbcDec.decrypt(encrypted);

    auto        unpadded = des::unpadPKCS7(decrypted);
    std::string result(unpadded.begin(), unpadded.end());
}
