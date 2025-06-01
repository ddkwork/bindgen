#pragma once

#include <string>
#include <sstream>
#include <cctype>
#include <cmath>
#include <cerrno>
#include <iomanip>
#include <stdexcept>
#include <cctype>
#include <climits>
#include <cstdlib>
#include <vector>
#include <utility>
#include <unicode/utf8.h>// 需要 ICU 库支持 Unicode

struct NumError : std::runtime_error {
    std::string func;
    std::string num;

    explicit NumError(const std::string &func, const std::string &num, const std::string &msg) : std::runtime_error(msg), func(func), num(num) {}
};

struct StrConv {
    static int Atoi(const std::string &s) {
        try {
            return std::stoi(s);
        } catch (const std::invalid_argument &) { throw NumError("Atoi", s, "invalid syntax"); } catch (const std::out_of_range &) {
            throw NumError("Atoi", s, "value out of range");
        }
    }

    static long ParseInt(const std::string &s, int base, int bitSize) {
        try {
            size_t pos = 0;
            long   value = std::stol(s, &pos, base);

            if (pos != s.size()) { throw std::invalid_argument("trailing characters"); }

            if (bitSize == 8) {
                if (value < SCHAR_MIN || value > SCHAR_MAX) { throw std::out_of_range("value out of range"); }
            } else if (bitSize == 16) {
                if (value < SHRT_MIN || value > SHRT_MAX) { throw std::out_of_range("value out of range"); }
            } else if (bitSize == 32) {
                if (value < INT_MIN || value > INT_MAX) { throw std::out_of_range("value out of range"); }
            }

            return value;
        } catch (const std::invalid_argument &) { throw NumError("ParseInt", s, "invalid syntax"); } catch (const std::out_of_range &) {
            throw NumError("ParseInt", s, "value out of range");
        }
    }

    static unsigned long ParseUint(const std::string &s, int base, int bitSize) {
        try {
            size_t        pos = 0;
            unsigned long value = std::stoul(s, &pos, base);

            if (pos != s.size()) { throw std::invalid_argument("trailing characters"); }

            if (bitSize == 8) {
                if (value > UCHAR_MAX) { throw std::out_of_range("value out of range"); }
            } else if (bitSize == 16) {
                if (value > USHRT_MAX) { throw std::out_of_range("value out of range"); }
            } else if (bitSize == 32) {
                if (value > UINT_MAX) { throw std::out_of_range("value out of range"); }
            }

            return value;
        } catch (const std::invalid_argument &) { throw NumError("ParseUint", s, "invalid syntax"); } catch (const std::out_of_range &) {
            throw NumError("ParseUint", s, "value out of range");
        }
    }

    static std::string Itoa(int i) { return std::to_string(i); }

    static std::string FormatInt(long i, int base) {
        std::ostringstream oss;
        switch (base) {
            case 2:
                oss << std::bitset<64>(i);
                break;
            case 8:
                oss << std::oct << i;
                break;
            case 10:
                oss << i;
                break;
            case 16:
                oss << std::hex << i;
                break;
            default:
                throw std::invalid_argument("invalid base");
        }
        return oss.str();
    }

    static std::string FormatUint(unsigned long i, int base) {
        std::ostringstream oss;
        switch (base) {
            case 2:
                oss << std::bitset<64>(i);
                break;
            case 8:
                oss << std::oct << i;
                break;
            case 10:
                oss << i;
                break;
            case 16:
                oss << std::hex << i;
                break;
            default:
                throw std::invalid_argument("invalid base");
        }
        return oss.str();
    }

    static double ParseFloat(const std::string &s) {
        try {
            return std::stod(s);
        } catch (const std::invalid_argument &) { throw NumError("ParseFloat", s, "invalid syntax"); } catch (const std::out_of_range &) {
            throw NumError("ParseFloat", s, "value out of range");
        }
    }

    static std::string FormatFloat(double f, char fmt, int prec) {
        std::ostringstream oss;
        oss << std::setprecision(prec);

        switch (fmt) {
            case 'f':
            case 'F':
                oss << std::fixed;
                break;
            case 'e':
            case 'E':
                oss << std::scientific;
                break;
            case 'g':
            case 'G':
                if (prec == 0) prec = 1;
                oss << std::defaultfloat << std::setprecision(prec);
                break;
            default:
                throw std::invalid_argument("invalid format");
        }

        oss << f;
        return oss.str();
    }

    static bool ParseBool(const std::string &str) {
        if (str == "true" || str == "1") return true;
        if (str == "false" || str == "0") return false;
        throw NumError("ParseBool", str, "invalid syntax");
    }

    static std::string FormatBool(bool b) { return b ? "true" : "false"; }

    static std::string Quote(const std::string &s) {
        std::ostringstream oss;
        oss << '"';
        for (char c: s) {
            switch (c) {
                case '"':
                    oss << "\\\"";
                    break;
                case '\\':
                    oss << "\\\\";
                    break;
                case '\a':
                    oss << "\\a";
                    break;
                case '\b':
                    oss << "\\b";
                    break;
                case '\f':
                    oss << "\\f";
                    break;
                case '\n':
                    oss << "\\n";
                    break;
                case '\r':
                    oss << "\\r";
                    break;
                case '\t':
                    oss << "\\t";
                    break;
                case '\v':
                    oss << "\\v";
                    break;
                default:
                    if (c < 0x20 || c >= 0x7F) {
                        oss << "\\x" << std::setw(2) << std::setfill('0') << std::hex << static_cast<int>(c);
                    } else {
                        oss << c;
                    }
            }
        }
        oss << '"';
        return oss.str();
    }

    static std::string QuoteToASCII(const std::string &s) {
        std::ostringstream oss;
        oss << '"';
        for (char c: s) {
            if (c < 0x20 || c >= 0x7F) {
                oss << "\\x" << std::setw(2) << std::setfill('0') << std::hex << static_cast<int>(c);
            } else {
                switch (c) {
                    case '"':
                        oss << "\\\"";
                        break;
                    case '\\':
                        oss << "\\\\";
                        break;
                    default:
                        oss << c;
                }
            }
        }
        oss << '"';
        return oss.str();
    }

    static std::string Unquote(const std::string &s) {
        if (s.size() < 2 || s[0] != '"' || s[s.size() - 1] != '"') { throw NumError("Unquote", s, "missing quotes"); }

        std::string result;
        bool        escape = false;

        for (size_t i = 1; i < s.size() - 1; i++) {
            char c = s[i];

            if (escape) {
                escape = false;
                switch (c) {
                    case 'a':
                        result += '\a';
                        break;
                    case 'b':
                        result += '\b';
                        break;
                    case 'f':
                        result += '\f';
                        break;
                    case 'n':
                        result += '\n';
                        break;
                    case 'r':
                        result += '\r';
                        break;
                    case 't':
                        result += '\t';
                        break;
                    case 'v':
                        result += '\v';
                        break;
                    case 'x': {
                        if (i + 2 >= s.size() - 1) { throw NumError("Unquote", s, "truncated escape sequence"); }
                        std::string hex = s.substr(i + 1, 2);
                        i += 2;
                        try {
                            int value = std::stoi(hex, nullptr, 16);
                            result += static_cast<char>(value);
                        } catch (...) { throw NumError("Unquote", s, "invalid hex escape"); }
                        break;
                    }
                    case 'u': {
                        if (i + 4 >= s.size() - 1) { throw NumError("Unquote", s, "truncated escape sequence"); }
                        std::string hex = s.substr(i + 1, 4);
                        i += 4;
                        try {
                            int value = std::stoi(hex, nullptr, 16);
                            result += unicodeToUTF8(value);
                        } catch (...) { throw NumError("Unquote", s, "invalid unicode escape"); }
                        break;
                    }
                    case 'U': {
                        if (i + 8 >= s.size() - 1) { throw NumError("Unquote", s, "truncated escape sequence"); }
                        std::string hex = s.substr(i + 1, 8);
                        i += 8;
                        try {
                            int value = std::stol(hex, nullptr, 16);
                            result += unicodeToUTF8(value);
                        } catch (...) { throw NumError("Unquote", s, "invalid unicode escape"); }
                        break;
                    }
                    default:
                        result += c;// 其他情况直接添加字符
                }
            } else {
                if (c == '\\') {
                    escape = true;
                } else {
                    result += c;
                }
            }
        }

        return result;
    }

    static std::string AppendInt(std::string dst, long i, int base) { return dst + FormatInt(i, base); }

    static std::string AppendUint(std::string dst, unsigned long i, int base) { return dst + FormatUint(i, base); }

    static std::string AppendFloat(std::string dst, double f, char fmt, int prec) { return dst + FormatFloat(f, fmt, prec); }

    static std::string AppendBool(std::string dst, bool b) { return dst + FormatBool(b); }

    static std::string AppendQuote(std::string dst, const std::string &s) { return dst + Quote(s); }

    static bool CanBackquote(const std::string &s) {
        for (char c: s) {
            if (c == '`' || c == '\n' || c == '\r') { return false; }
            if (c < 0x20 || c > 0x7E) { return false; }
        }
        return true;
    }

private:
    static std::string unicodeToUTF8(int code) {
        std::string utf8;
        if (code <= 0x7F) {
            utf8 += static_cast<char>(code);
        } else if (code <= 0x7FF) {
            utf8 += static_cast<char>(0xC0 | (code >> 6));
            utf8 += static_cast<char>(0x80 | (code & 0x3F));
        } else if (code <= 0xFFFF) {
            utf8 += static_cast<char>(0xE0 | (code >> 12));
            utf8 += static_cast<char>(0x80 | ((code >> 6) & 0x3F));
            utf8 += static_cast<char>(0x80 | (code & 0x3F));
        } else {
            utf8 += static_cast<char>(0xF0 | (code >> 18));
            utf8 += static_cast<char>(0x80 | ((code >> 12) & 0x3F));
            utf8 += static_cast<char>(0x80 | ((code >> 6) & 0x3F));
            utf8 += static_cast<char>(0x80 | (code & 0x3F));
        }
        return utf8;
    }
};

#include "strconv.h"
#include <iostream>

int main() {
    try {
        int i = StrConv::Atoi("12345");
        std::cout << "Atoi: " << i << "\n";

        long hexVal = StrConv::ParseInt("1a3", 16, 32);
        std::cout << "ParseInt (hex): " << hexVal << "\n";

        double d = StrConv::ParseFloat("3.1415926");
        std::cout << "ParseFloat: " << d << "\n";

        bool b = StrConv::ParseBool("true");
        std::cout << "ParseBool: " << std::boolalpha << b << "\n";

        std::string quoted = StrConv::Quote("Hello\nWorld");
        std::cout << "Quote: " << quoted << "\n";

        std::string unquoted = StrConv::Unquote("\"Hello\\nWorld\"");
        std::cout << "Unquote: " << unquoted << "\n";

        std::string data = "Value=";
        data = StrConv::AppendInt(data, 42, 10);
        data = StrConv::AppendBool(data, true);
        std::cout << "Append: " << data << "\n";

        std::cout << "CanBackquote: " << StrConv::CanBackquote("simple`string") << "\n";

    } catch (const NumError &e) { std::cerr << "Error in " << e.func << "(" << e.num << "): " << e.what() << "\n"; } catch (const std::exception &e) {
        std::cerr << "Error: " << e.what() << "\n";
    }

    return 0;
}
