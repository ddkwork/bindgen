#pragma once
#include <cstdint>
#include <vector>
#include <array>
#include <string>
#include <random>
#include <stdexcept>
#include <cmath>
#include <iostream>
#include <limits>
#include <tuple>
#include <memory>

struct rsa_Error {
    std::string msg;

    rsa_Error(const std::string &m) : msg(m) {}

    std::string Error() const { return msg; }
};

rsa_Error rsa_ErrDecryption = rsa_Error("decryption error");
rsa_Error rsa_ErrVerification = rsa_Error("verification error");
rsa_Error rsa_ErrMessageTooLong = rsa_Error("message too long for RSA key size");
rsa_Error rsa_ErrInvalidKey = rsa_Error("invalid RSA key");

using BigInt = std::vector<uint64_t>;
using Word = uint64_t;
constexpr Word BASE = std::numeric_limits<Word>::max();
constexpr Word BASE_BITS = std::numeric_limits<Word>::digits;

namespace bigint {
    int compare(const BigInt &a, const BigInt &b) {
        if (a.size() != b.size()) { return a.size() < b.size() ? -1 : 1; }
        for (int i = a.size() - 1; i >= 0; i--) {
            if (a[i] != b[i]) { return a[i] < b[i] ? -1 : 1; }
        }
        return 0;
    }

    BigInt add(const BigInt &a, const BigInt &b) {
        size_t n = std::max(a.size(), b.size());
        BigInt result;
        Word   carry = 0;

        for (size_t i = 0; i < n; i++) {
            Word a_val = (i < a.size()) ? a[i] : 0;
            Word b_val = (i < b.size()) ? b[i] : 0;

            Word sum = a_val + b_val + carry;
            carry = (sum < a_val || sum < b_val) ? 1 : 0;

            result.push_back(sum);
        }

        if (carry) { result.push_back(carry); }

        return result;
    }

    BigInt sub(const BigInt &a, const BigInt &b) {
        BigInt result;
        Word   borrow = 0;

        for (size_t i = 0; i < a.size(); i++) {
            Word a_val = a[i];
            Word b_val = (i < b.size()) ? b[i] : 0;

            Word diff;
            if (a_val >= b_val + borrow) {
                diff = a_val - b_val - borrow;
                borrow = 0;
            } else {
                diff = a_val + (BASE - b_val - borrow);
                borrow = 1;
            }

            result.push_back(diff);
        }

        while (result.size() > 1 && result.back() == 0) { result.pop_back(); }

        return result;
    }

    BigInt mul(const BigInt &a, const BigInt &b) {
        size_t len = a.size() + b.size();
        BigInt result(len, 0);

        for (size_t i = 0; i < a.size(); i++) {
            Word carry = 0;
            for (size_t j = 0; j < b.size() || carry > 0; j++) {
                Word b_val = (j < b.size()) ? b[j] : 0;
                Word prod = (static_cast<uint128_t>(a[i]) * b_val) + result[i + j] + carry;

                result[i + j] = static_cast<Word>(prod);

                carry = static_cast<Word>(prod >> WORD_BITS);
            }
        }

        while (result.size() > 1 && result.back() == 0) { result.pop_back(); }

        return result;
    }

    std::pair<BigInt, BigInt> div(const BigInt &a, const BigInt &b) {
        if (b.size() == 0) throw std::runtime_error("division by zero");

        if (compare(a, b) < 0) { return {BigInt{0}, a}; }

        BigInt quotient, remainder;
        BigInt current = a;

        while (compare(current, b) >= 0) {
            current = sub(current, b);
            quotient = add(quotient, BigInt{1});
        }

        remainder = current;
        return {quotient, remainder};
    }

    BigInt modExp(const BigInt &base, const BigInt &exp, const BigInt &mod) {
        BigInt result = {1};
        BigInt base_mod = base;
        BigInt exp_copy = exp;

        if (compare(base, mod) >= 0) { base_mod = div(base, mod).second; }

        while (exp_copy.size() > 0 && (exp_copy.size() > 1 || exp_copy[0] > 0)) {
            if (exp_copy[0] & 1) { result = div(mul(result, base_mod), mod).second; }

            base_mod = div(mul(base_mod, base_mod), mod).second;

            Word carry = 0;
            for (int i = exp_copy.size() - 1; i >= 0; i--) {
                Word word = exp_copy[i];
                exp_copy[i] = (word >> 1) | (carry << (WORD_BITS - 1));
                carry = word & 1;
            }

            while (exp_copy.size() > 1 && exp_copy.back() == 0) { exp_copy.pop_back(); }
        }

        return result;
    }

    std::tuple<BigInt, BigInt, BigInt> extGCD(const BigInt &a, const BigInt &b) {
        if (b.size() == 0) { return {a, {1}, {0}}; }

        auto [quot, rem] = div(a, b);
        auto [d, x, y] = extGCD(b, rem);
        return {d, y, sub(x, mul(quot, y))};
    }

    BigInt modInverse(const BigInt &a, const BigInt &n) {
        auto [d, x, y] = extGCD(a, n);
        if (compare(d, BigInt{1}) != 0) { throw rsa_ErrInvalidKey; }

        if (compare(x, BigInt{0}) < 0) { x = add(x, n); }

        return div(x, n).second;
    }

    BigInt fromBytes(const std::vector<uint8_t> &bytes) {
        BigInt result;
        for (int i = bytes.size(); i > 0; i -= sizeof(Word)) {
            Word   word = 0;
            size_t chunkSize = std::min(sizeof(Word), static_cast<size_t>(i));
            size_t start = i - chunkSize;

            for (size_t j = 0; j < chunkSize; j++) { word = (word << 8) | bytes[start + j]; }

            result.push_back(word);
        }

        return result;
    }

    std::vector<uint8_t> toBytes(const BigInt &n) {
        std::vector<uint8_t> bytes;
        for (Word word: n) {
            for (int i = sizeof(Word) - 1; i >= 0; i--) { bytes.push_back(static_cast<uint8_t>((word >> (i * 8)) & 0xFF)); }
        }

        while (bytes.size() > 1 && bytes[0] == 0) { bytes.erase(bytes.begin()); }

        return bytes;
    }

    BigInt rand(size_t bits) {
        size_t               bytes = (bits + 7) / 8;
        std::vector<uint8_t> buffer(bytes);

        std::random_device                     rd;
        std::mt19937                           gen(rd());
        std::uniform_int_distribution<uint8_t> dist(0, 255);

        for (size_t i = 0; i < bytes; i++) { buffer[i] = dist(gen); }

        if (bits % 8 != 0) { buffer[0] &= 0xFF >> (8 - (bits % 8)); }
        buffer[0] |= 0x80 >> (8 - (bits % 8));

        return fromBytes(buffer);
    }

    bool isPrime(const BigInt &n) {
        BigInt two = {2};
        BigInt three = {3};
        BigInt i = {5};

        if (n.size() == 1) {
            Word val = n[0];
            if (val < 2) return false;
            if (val == 2 || val == 3) return true;
            if (val % 2 == 0) return false;
        }

        if (compare(div(n, two).second, BigInt{0}) == 0) return false;
        if (compare(div(n, three).second, BigInt{0}) == 0) return false;

        while (compare(i, n) < 0) {
            if (compare(div(n, i).second, BigInt{0}) == 0) return false;
            i = add(i, two);
            i = add(i, two);
        }

        return true;
    }

    BigInt genPrime(size_t bits) {
        while (true) {
            BigInt candidate = rand(bits);
            if ((candidate[0] & 1) == 0) { candidate[0] |= 1; }

            if (isPrime(candidate)) { return candidate; }
        }
    }
};// namespace bigint

struct rsa_PublicKey {
    BigInt N;// modulus
    BigInt E;// public exponent

    static constexpr BigInt DefaultExponent = {0x10001};// 65537

    size_t Size() const { return bigint::toBytes(N).size(); }
};

struct rsa_PrivateKey {
    rsa_PublicKey PublicKey;
    BigInt        D;        // private exponent
    BigInt        Primes[2];// p, q (用于中国剩余定理优化)

    size_t Size() const { return PublicKey.Size(); }
};

struct rsa {
    static rsa_PrivateKey GenerateKey(size_t bits) {
        if (bits < 512) { throw rsa_ErrInvalidKey; }

        rsa_PrivateKey priv;
        rsa_PublicKey &pub = priv.PublicKey;

        pub.E = rsa_PublicKey::DefaultExponent;

        priv.Primes[0] = bigint::genPrime(bits / 2);
        priv.Primes[1] = bigint::genPrime(bits - bits / 2);

        pub.N = bigint::mul(priv.Primes[0], priv.Primes[1]);

        BigInt pMinus1 = bigint::sub(priv.Primes[0], BigInt{1});
        BigInt qMinus1 = bigint::sub(priv.Primes[1], BigInt{1});
        BigInt phiN = bigint::mul(pMinus1, qMinus1);

        priv.D = bigint::modInverse(pub.E, phiN);

        return priv;
    }

    static std::vector<uint8_t> Encrypt(const rsa_PublicKey &pub, const std::vector<uint8_t> &msg, bool pkcsPadding = true) {
        size_t k = pub.Size();
        if (msg.size() > k - 11) { throw rsa_ErrMessageTooLong; }

        std::vector<uint8_t> em;
        if (pkcsPadding) {
            em.push_back(0x00);
            em.push_back(0x02);

            size_t paddingLen = k - msg.size() - 3;
            for (size_t i = 0; i < paddingLen; i++) { em.push_back(rand() % 255 + 1); }

            em.push_back(0x00);
        }

        em.insert(em.end(), msg.begin(), msg.end());

        BigInt m = bigint::fromBytes(em);
        if (bigint::compare(m, pub.N) >= 0) { throw rsa_ErrMessageTooLong; }

        BigInt c = bigint::modExp(m, pub.E, pub.N);

        return bigint::toBytes(c);
    }

    static std::vector<uint8_t> Decrypt(const rsa_PrivateKey &priv, const std::vector<uint8_t> &ciphertext, bool pkcsPadding = true) {
        size_t k = priv.Size();
        if (ciphertext.size() != k) { throw rsa_ErrDecryption; }

        BigInt c = bigint::fromBytes(ciphertext);
        if (bigint::compare(c, priv.PublicKey.N) >= 0) { throw rsa_ErrDecryption; }

        BigInt m = bigint::modExp(c, priv.D, priv.PublicKey.N);

        std::vector<uint8_t> em = bigint::toBytes(m);
        if (em.size() < k) { em.insert(em.begin(), k - em.size(), 0); }

        if (pkcsPadding) {
            if (em[0] != 0x00 || em[1] != 0x02) { throw rsa_ErrDecryption; }

            size_t i = 2;
            while (i < em.size() && em[i] != 0) { i++; }

            if (i < 10 || i == em.size()) { throw rsa_ErrDecryption; }

            return std::vector<uint8_t>(em.begin() + i + 1, em.end());
        } else {
            return em;
        }
    }

    static std::vector<uint8_t> Sign(const rsa_PrivateKey &priv, const std::vector<uint8_t> &digest, bool pkcsPadding = true) { return Encrypt(priv, digest, pkcsPadding); }

    static bool Verify(const rsa_PublicKey &pub, const std::vector<uint8_t> &digest, const std::vector<uint8_t> &signature, bool pkcsPadding = true) {
        try {
            std::vector<uint8_t> recovered = Decrypt(pub, signature, pkcsPadding);
            return recovered == digest;
        } catch (const rsa_Error &) { return false; }
    }
};

struct pkcs1 {
    static std::vector<uint8_t> PadMessage(const rsa_PublicKey &pub, const std::vector<uint8_t> &msg) { return rsa::Encrypt(pub, msg, true); }

    static std::vector<uint8_t> UnpadMessage(const rsa_PrivateKey &priv, const std::vector<uint8_t> &encrypted) { return rsa::Decrypt(priv, encrypted, true); }

    static std::vector<uint8_t> HashPrefix(const std::vector<uint8_t> &digest, const std::string &hashName) {
        static const std::vector<uint8_t> SHA256_ID = {0x30, 0x31, 0x30, 0x0d, 0x06, 0x09, 0x60, 0x86, 0x48, 0x01, 0x65, 0x03, 0x04, 0x02, 0x01, 0x05, 0x00, 0x04, 0x20};

        if (hashName == "SHA256") {
            std::vector<uint8_t> prefixed = SHA256_ID;
            prefixed.insert(prefixed.end(), digest.begin(), digest.end());
            return prefixed;
        }

        return digest;
    }
};

void testRSAFunctions() {
    rsa_PrivateKey privKey = rsa::GenerateKey(1024);
    rsa_PublicKey  pubKey = privKey.PublicKey;

    std::string          message = "Hello, RSA!";
    std::vector<uint8_t> plaintext(message.begin(), message.end());

    std::vector<uint8_t> ciphertext = rsa::Encrypt(pubKey, plaintext);

    std::vector<uint8_t> decrypted = rsa::Decrypt(privKey, ciphertext);
    std::string          recovered(decrypted.begin(), decrypted.end());

    std::cout << "Original: " << message << std::endl;
    std::cout << "Decrypted: " << recovered << std::endl;

    std::vector<uint8_t> digest = {0x31, 0x5f, 0x5b, 0xdb, 0x1d, 0x1d, 0xc9, 0x03, 0x9f, 0x60, 0x97, 0x0a, 0xbb, 0x97, 0xd8, 0xd7,
                                   0xb4, 0xfd, 0xfc, 0xfd, 0x55, 0xb0, 0xb7, 0x65, 0x5f, 0x59, 0xf0, 0x88, 0x71, 0x60, 0x64, 0x8b};

    std::vector<uint8_t> prefix = pkcs1::HashPrefix(digest, "SHA256");
    std::vector<uint8_t> signature = rsa::Sign(privKey, prefix);

    bool valid = rsa::Verify(pubKey, prefix, signature);
    std::cout << "Signature valid: " << std::boolalpha << valid << std::endl;

    try {
        ciphertext.resize(ciphertext.size() - 1);
        decrypted = rsa::Decrypt(privKey, ciphertext);
    } catch (const rsa_Error &e) { std::cerr << "Decryption error: " << e.Error() << std::endl; }

    rsa_PrivateKey smallKey = rsa::GenerateKey(512);
    std::cout << "512-bit key size: " << smallKey.Size() << " bytes" << std::endl;
    std::cout << "2048-bit key size: " << rsa::GenerateKey(2048).Size() << " bytes" << std::endl;
}
