#pragma once
#include <cmath>
#include <limits>
#include <stdexcept>
#include <vector>
#include <string>
#include <algorithm>
#include <iostream>
#include <cctype>
#include <climits>
#include <cstring>
#include <iomanip>
#include <type_traits>
#include <concepts>
#include <bit>

constexpr double E = 2.71828182845904523536028747135266249775724709369995957496696763;
constexpr double PI = 3.14159265358979323846264338327950288419716939937510582097494459;
constexpr double PHI = 1.61803398874989484820458683436563811772030917980576286213544862;
constexpr double SQRT2 = 1.41421356237309504880168872420969807856967187537694807317667974;
constexpr double SQRT_E = 1.64872127070012814684865078781416357165377610071014801157507931;
constexpr double SQRT_PI = 1.77245385090551602729816748334114518279754945612238712821380779;
constexpr double SQRT_PHI = 1.27201964951406896425242246173749149171560804184009624861664038;
constexpr double LN2 = 0.693147180559945309417232121458176568075500134360255254120680009;
constexpr double LOG2E = 1.442695040888963407359924681001892137426645954152985934135449406;
constexpr double LN10 = 2.30258509299404568401799145468436420760110148862877297603332790;
constexpr double LOG10E = 0.43429448190325182765112891891660508229439700580366656611445378;

constexpr float  MAX_FLOAT32 = 3.40282346638528859811704183484516925440e+38F;
constexpr float  MIN_FLOAT32 = 1.401298464324817070923729583289916131280e-45F;
constexpr double MAX_FLOAT64 = 1.797693134862315708145274237317043567981e+308;
constexpr double MIN_FLOAT64 = 4.940656458412465441765687928682213723651e-324;

constexpr int          MAX_INT = 2147483647;
constexpr int          MIN_INT = -2147483648;
constexpr int          MAX_INT8 = 127;
constexpr int          MIN_INT8 = -128;
constexpr int          MAX_INT16 = 32767;
constexpr int          MIN_INT16 = -32768;
constexpr int          MAX_INT32 = 2147483647;
constexpr int          MIN_INT32 = -2147483648;
constexpr int64_t      MAX_INT64 = 9223372036854775807LL;
constexpr int64_t      MIN_INT64 = -9223372036854775807LL - 1;
constexpr unsigned int MAX_UINT = 4294967295U;
constexpr uint8_t      MAX_UINT8 = 255;
constexpr uint16_t     MAX_UINT16 = 65535;
constexpr uint32_t     MAX_UINT32 = 4294967295U;
constexpr uint64_t     MAX_UINT64 = 18446744073709551615ULL;

struct BigInt {
    std::vector<uint32_t> digits;
    bool                  negative;

    BigInt() : digits({0}), negative(false) {}

    BigInt(int64_t value) {
        negative = value < 0;
        uint64_t absValue = value < 0 ? static_cast<uint64_t>(-value) : static_cast<uint64_t>(value);

        digits.clear();
        if (absValue == 0) {
            digits.push_back(0);
            return;
        }

        while (absValue > 0) {
            digits.push_back(static_cast<uint32_t>(absValue & 0xFFFFFFFF));
            absValue >>= 32;
        }
        trim();
    }

    BigInt(const std::string &str) : digits({0}), negative(false) {
        if (str.empty()) return;

        size_t start = 0;
        if (str[0] == '-') {
            negative = true;
            start = 1;
        } else if (str[0] == '+') {
            start = 1;
        }

        while (start < str.size() && str[start] == '0') start++;

        if (start == str.size()) {
            digits = {0};
            negative = false;
            return;
        }

        std::string numStr = str.substr(start);
        digits.clear();
        negative = (str[0] == '-');

        while (!numStr.empty()) {
            uint32_t    remainder = 0;
            std::string newStr;

            for (char c: numStr) {
                uint32_t current = remainder * 10 + (c - '0');
                if (current < 10 && newStr.empty()) continue;
                newStr.push_back('0' + current / 10);
                remainder = current % 10;
            }

            digits.push_back(remainder);
            numStr = newStr;
        }

        std::reverse(digits.begin(), digits.end());
        trim();
    }

    void trim() {
        while (!digits.empty() && digits.back() == 0) { digits.pop_back(); }
        if (digits.empty()) {
            digits.push_back(0);
            negative = false;
        }
    }

    bool isZero() const { return digits.size() == 1 && digits[0] == 0; }

    BigInt negate() const {
        BigInt result = *this;
        if (!isZero()) { result.negative = !negative; }
        return result;
    }

    int compare(const BigInt &other) const {
        if (negative != other.negative) { return negative ? -1 : 1; }

        if (digits.size() != other.digits.size()) { return (digits.size() < other.digits.size()) ? -1 : 1; }

        for (int i = static_cast<int>(digits.size()) - 1; i >= 0; i--) {
            if (digits[i] != other.digits[i]) { return (digits[i] < other.digits[i]) ? -1 : 1; }
        }

        return 0;
    }

    bool operator==(const BigInt &other) const { return compare(other) == 0; }

    bool operator!=(const BigInt &other) const { return compare(other) != 0; }

    bool operator<(const BigInt &other) const { return compare(other) < 0; }

    bool operator>(const BigInt &other) const { return compare(other) > 0; }

    bool operator<=(const BigInt &other) const { return compare(other) <= 0; }

    bool operator>=(const BigInt &other) const { return compare(other) >= 0; }

    BigInt operator+(const BigInt &other) const {
        if (negative == other.negative) {
            BigInt result;
            result.negative = negative;

            size_t maxSize = std::max(digits.size(), other.digits.size());
            result.digits.resize(maxSize + 1, 0);

            uint64_t carry = 0;
            for (size_t i = 0; i < maxSize; i++) {
                uint64_t sum = carry;
                if (i < digits.size()) sum += digits[i];
                if (i < other.digits.size()) sum += other.digits[i];

                result.digits[i] = static_cast<uint32_t>(sum & 0xFFFFFFFF);
                carry = sum >> 32;
            }

            if (carry) result.digits[maxSize] = static_cast<uint32_t>(carry);
            else
                result.digits.resize(maxSize);

            result.trim();
            return result;
        }

        return negative ? (other - negate()) : (*this - other.negate());
    }

    BigInt operator-(const BigInt &other) const {
        if (negative != other.negative) { return *this + other.negate(); }

        if (negative) { return other.negate() - negate(); }

        int cmp = compare(other);
        if (cmp == 0) return BigInt(0);
        if (cmp < 0) return (other - *this).negate();

        BigInt result;
        result.digits.resize(digits.size());

        uint64_t borrow = 0;
        for (size_t i = 0; i < digits.size(); i++) {
            uint64_t diff = static_cast<uint64_t>(digits[i]) - borrow;
            borrow = 0;

            if (i < other.digits.size()) {
                if (diff < other.digits[i]) {
                    diff += (1ULL << 32);
                    borrow = 1;
                }
                diff -= other.digits[i];
            } else if (borrow) {
                if (diff == 0) {
                    diff = (1ULL << 32) - 1;
                    borrow = 1;
                } else {
                    diff -= 1;
                    borrow = 0;
                }
            }

            result.digits[i] = static_cast<uint32_t>(diff);
        }

        result.trim();
        return result;
    }

    BigInt operator*(const BigInt &other) const {
        if (isZero() || other.isZero()) return BigInt(0);

        BigInt result;
        result.digits.resize(digits.size() + other.digits.size());
        result.negative = negative != other.negative;

        for (size_t i = 0; i < digits.size(); i++) {
            uint64_t carry = 0;
            for (size_t j = 0; j < other.digits.size(); j++) {
                uint64_t product = static_cast<uint64_t>(digits[i]) * other.digits[j] + result.digits[i + j] + carry;
                result.digits[i + j] = static_cast<uint32_t>(product & 0xFFFFFFFF);
                carry = product >> 32;
            }

            if (carry) { result.digits[i + other.digits.size()] = static_cast<uint32_t>(carry); }
        }

        result.trim();
        return result;
    }

    BigInt operator/(const BigInt &divisor) const {
        if (divisor.isZero()) throw std::overflow_error("division by zero");
        if (isZero()) return BigInt(0);
        if (divisor.digits.size() == 1 && divisor.digits[0] == 1) { return divisor.negative ? negate() : *this; }

        int cmp = compare(divisor);
        if (cmp < 0) return BigInt(0);
        if (cmp == 0) return BigInt(negative != divisor.negative ? -1 : 1);

        BigInt dividend = *this;
        dividend.negative = false;
        BigInt div = divisor;
        div.negative = false;

        BigInt quotient;
        BigInt remainder;

        for (int i = static_cast<int>(digits.size()) - 1; i >= 0; i--) {
            remainder = remainder * BigInt(1ULL << 32) + BigInt(static_cast<int64_t>(digits[i]));

            uint32_t qDigit = 0;
            uint64_t low = 0, high = (1ULL << 32) - 1;

            while (low <= high) {
                uint64_t mid = (low + high) / 2;
                BigInt   product = div * BigInt(static_cast<int64_t>(mid));

                if (product <= remainder) {
                    qDigit = static_cast<uint32_t>(mid);
                    low = mid + 1;
                } else {
                    high = mid - 1;
                }
            }

            quotient.digits.insert(quotient.digits.begin(), qDigit);
            BigInt product = div * BigInt(static_cast<int64_t>(qDigit));
            remainder = remainder - product;
        }

        quotient.trim();
        quotient.negative = negative != divisor.negative;
        return quotient;
    }

    BigInt operator%(const BigInt &divisor) const {
        BigInt quotient = *this / divisor;
        return *this - (quotient * divisor);
    }

    std::string toString() const {
        if (isZero()) return "0";

        std::string result;
        BigInt      num = *this;
        num.negative = false;

        while (!num.isZero()) {
            BigInt quotient = num / BigInt(10);
            BigInt remainder = num - (quotient * BigInt(10));
            result.push_back('0' + remainder.digits[0]);
            num = quotient;
        }

        if (negative) result.push_back('-');
        std::reverse(result.begin(), result.end());
        return result;
    }
};

template<typename T>
T abs(T x) {
    if constexpr (std::is_floating_point_v<T>) return std::abs(x);
    else if constexpr (std::is_integral_v<T>)
        return x < 0 ? -x : x;
    else
        static_assert(sizeof(T) == 0, "Unsupported type");
}

template<typename T>
T max(T a, T b) {
    return a > b ? a : b;
}

template<typename T>
T min(T a, T b) {
    return a < b ? a : b;
}

double floor(double x) { return std::floor(x); }

double ceil(double x) { return std::ceil(x); }

double trunc(double x) { return std::trunc(x); }

double round(double x) { return std::round(x); }

double exp(double x) { return std::exp(x); }

double log(double x) { return std::log(x); }

double log10(double x) { return std::log10(x); }

double log2(double x) { return std::log2(x); }

double log1p(double x) { return std::log1p(x); }

double expm1(double x) { return std::expm1(x); }

double pow(double x, double y) { return std::pow(x, y); }

double sqrt(double x) { return std::sqrt(x); }

double cbrt(double x) { return std::cbrt(x); }

double hypot(double p, double q) { return std::hypot(p, q); }

double sin(double x) { return std::sin(x); }

double cos(double x) { return std::cos(x); }

double tan(double x) { return std::tan(x); }

double asin(double x) { return std::asin(x); }

double acos(double x) { return std::acos(x); }

double atan(double x) { return std::atan(x); }

double atan2(double y, double x) { return std::atan2(y, x); }

double sinh(double x) { return std::sinh(x); }

double cosh(double x) { return std::cosh(x); }

double tanh(double x) { return std::tanh(x); }

double asinh(double x) { return std::asinh(x); }

double acosh(double x) { return std::acosh(x); }

double atanh(double x) { return std::atanh(x); }

bool isInf(double x, int sign) {
    if (sign > 0) return x == std::numeric_limits<double>::infinity();
    if (sign < 0) return x == -std::numeric_limits<double>::infinity();
    return std::isinf(x);
}

bool isNaN(double x) { return std::isnan(x); }

bool isFinite(double x) { return std::isfinite(x); }

double copysign(double x, double y) { return std::copysign(x, y); }

double nextafter(double x, double y) { return std::nextafter(x, y); }

double erf(double x) {
    if (x == 0) return 0;
    if (std::isinf(x)) return x > 0 ? 1.0 : -1.0;

    const double a1 = 0.254829592;
    const double a2 = -0.284496736;
    const double a3 = 1.421413741;
    const double a4 = -1.453152027;
    const double a5 = 1.061405429;
    const double p = 0.3275911;

    int sign = x < 0 ? -1 : 1;
    x = abs(x);

    double t = 1.0 / (1.0 + p * x);
    double y = 1.0 - (((((a5 * t + a4) * t + a3) * t + a2) * t + a1) * t * std::exp(-x * x));

    return sign * y;
}

double erfc(double x) { return 1.0 - erf(x); }

double gamma(double x) {
    if (x <= 0) {
        if (x == floor(x)) return std::numeric_limits<double>::infinity();
        return PI / (sin(PI * x) * gamma(1 - x));
    }

    const double g = 7;
    const double p[] = {0.99999999999980993,
                        676.5203681218851,
                        -1259.1392167224028,
                        771.32342877765313,
                        -176.61502916214059,
                        12.507343278686905,
                        -0.13857109526572012,
                        9.9843695780195716e-6,
                        1.5056327351493116e-7};

    if (x < 0.5) return PI / (sin(PI * x) * gamma(1 - x));

    x -= 1;
    double a = p[0];
    double t = x + g + 0.5;

    for (int i = 1; i < 9; i++) a += p[i] / (x + i);

    return sqrt(2 * PI) * pow(t, x + 0.5) * exp(-t) * a;
}

double lgamma(double x) { return std::lgamma(x); }

int iabs(int x) { return x < 0 ? -x : x; }

int64_t iabs64(int64_t x) { return x < 0 ? -x : x; }

template<typename T>
T div(T x, T y) {
    if (y == 0) throw std::overflow_error("integer divide by zero");
    return x / y;
}

template<typename T>
T mod(T x, T y) {
    if (y == 0) throw std::overflow_error("integer modulo by zero");
    return x % y;
}

template<typename T>
int addOvf(T x, T y, T *res) {
    if constexpr (std::is_unsigned_v<T>) {
        T r = x + y;
        if (r < x) return -1;
        *res = r;
        return 0;
    } else {
        if ((y > 0 && x > std::numeric_limits<T>::max() - y) || (y < 0 && x < std::numeric_limits<T>::min() - y)) return -1;
        *res = x + y;
        return 0;
    }
}

int leadingZeros(uint32_t x) { return x ? __builtin_clz(x) : 32; }

int leadingZeros(uint64_t x) { return x ? __builtin_clzll(x) : 64; }

int trailingZeros(uint32_t x) { return x ? __builtin_ctz(x) : 32; }

int trailingZeros(uint64_t x) { return x ? __builtin_ctzll(x) : 64; }

int onesCount(uint32_t x) { return __builtin_popcount(x); }

int onesCount(uint64_t x) { return __builtin_popcountll(x); }

uint32_t rotateLeft(uint32_t x, int k) {
    k &= 31;
    return k ? (x << k) | (x >> (32 - k)) : x;
}

uint64_t rotateLeft64(uint64_t x, int k) {
    k &= 63;
    return k ? (x << k) | (x >> (64 - k)) : x;
}

uint64_t float64bits(double x) {
    uint64_t bits;
    memcpy(&bits, &x, sizeof(double));
    return bits;
}

double float64frombits(uint64_t bits) {
    double value;
    memcpy(&value, &bits, sizeof(uint64_t));
    return value;
}

uint32_t float32bits(float x) {
    uint32_t bits;
    memcpy(&bits, &x, sizeof(float));
    return bits;
}

float float32frombits(uint32_t bits) {
    float value;
    memcpy(&value, &bits, sizeof(uint32_t));
    return value;
}

void testMathFunctions() {
    std::cout << "abs(-3.5) = " << abs(-3.5) << "\n";
    std::cout << "abs(-10) = " << abs(-10) << "\n";
    std::cout << "max(2.3, 5.1) = " << max(2.3, 5.1) << "\n";
    std::cout << "min(2, 5) = " << min(2, 5) << "\n\n";

    std::cout << "floor(3.7) = " << floor(3.7) << "\n";
    std::cout << "ceil(3.2) = " << ceil(3.2) << "\n";
    std::cout << "sin(PI/2) = " << sin(PI / 2) << "\n\n";

    std::cout << "erf(1) ≈ " << erf(1) << "\n";
    std::cout << "gamma(5) = " << gamma(5) << "\n\n";

    BigInt big1("12345678901234567890");
    BigInt big2("98765432109876543210");
    BigInt bigSum = big1 + big2;
    BigInt bigProd = big1 * big2;

    std::cout << "BigInt Sum: " << big1.toString() << " + " << big2.toString() << " = " << bigSum.toString() << "\n";
    std::cout << "BigInt Product: " << big1.toString() << " * " << big2.toString() << " = " << bigProd.toString() << "\n";

    uint64_t bits = float64bits(PI);
    std::cout << "\nPI bits: 0x" << std::hex << bits << std::dec;
    std::cout << " → " << float64frombits(bits) << "\n";

    std::cout << "Leading zeros in 16: " << leadingZeros(16u) << "\n";

    int res;
    int overflow = addOvf(MAX_INT, 1, &res);
    std::cout << "\nMAX_INT + 1: " << (overflow ? "overflow" : "no overflow") << "\n";
}
