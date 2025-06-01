#pragma once
#include <string>
#include <sstream>
#include <vector>
#include <iomanip>
#include <cctype>
#include <cmath>
#include <cstdio>
#include <stdexcept>
#include <type_traits>
#include <any>
#include <memory>

struct fmt_Error {
    std::string msg;

    fmt_Error() {}

    fmt_Error(const std::string &m) : msg(m) {}

    std::string Error() const { return msg; }

    operator bool() const { return !msg.empty(); }
};

fmt_Error fmt_ErrInvalidVerb = fmt_Error("invalid verb");
fmt_Error fmt_ErrMissingArg = fmt_Error("missing argument");
fmt_Error fmt_ErrNoVerb = fmt_Error("no verb for value");
fmt_Error fmt_ErrBadWidth = fmt_Error("bad width value");
fmt_Error fmt_ErrBadPrec = fmt_Error("bad precision value");

struct fmt_State {
    std::ostringstream    buf;
    int                   argNum = 0;
    bool                  reordered = false;
    std::vector<std::any> args;
    std::map<int, int>    argMap;

    void writeByte(char c) { buf << c; }

    void writeString(const std::string &s) { buf << s; }

    std::string string() const { return buf.str(); }

    int argIndex(int index) {
        if (reordered) {
            auto it = argMap.find(index);
            if (it != argMap.end()) { return it->second; }
            return -1;
        }
        return index;
    }
};

struct fmt {
    std::function<std::string(const std::string &, const std::vector<std::any> &)> Sprintf = [](const std::string &format, const std::vector<std::any> &args) {
        fmt_State st;
        st.args = args;
        doPrintf(st, format);
        return st.string();
    };

    std::function<void(const std::string &, const std::vector<std::any> &)> Printf = [](const std::string &format, const std::vector<std::any> &args) {
        fmt_State st;
        st.args = args;
        doPrintf(st, format);
        std::cout << st.string();
    };

    std::function<void(FILE *, const std::string &, const std::vector<std::any> &)> Fprintf = [](FILE *f, const std::string &format, const std::vector<std::any> &args) {
        fmt_State st;
        st.args = args;
        doPrintf(st, format);
        fprintf(f, "%s", st.string().c_str());
    };

    std::function<std::string(const std::vector<std::any> &)> Sprint = [](const std::vector<std::any> &args) {
        fmt_State st;
        printArgs(st, args);
        return st.string();
    };

    std::function<void(const std::vector<std::any> &)> Println = [](const std::vector<std::any> &args) {
        fmt_State st;
        printArgs(st, args);
        std::cout << st.string() << "\n";
    };

    std::function<fmt_Error(const std::string &, const std::vector<std::any> &)> Errorf = [](const std::string &format, const std::vector<std::any> &args) {
        return fmt_Error(Sprintf(format, args));
    };

private:
    static void doPrintf(fmt_State &st, const std::string &format) {
        st.argNum = 0;
        for (size_t i = 0; i < format.size();) {
            char c = format[i];
            if (c != '%') {
                st.writeByte(c);
                i++;
                continue;
            }

            i++;
            if (i >= format.size()) { throw std::runtime_error("incomplete format specifier"); }

            if (format[i] == '%') {
                st.writeByte('%');
                i++;
                continue;
            }

            int argIndex = -1;
            if (std::isdigit(format[i])) {
                argIndex = 0;
                while (i < format.size() && std::isdigit(format[i])) {
                    argIndex = argIndex * 10 + (format[i] - '0');
                    i++;
                }
                if (i < format.size() && format[i] == '[') {
                    st.reordered = true;
                    i++;
                    int mappedIndex = 0;
                    while (i < format.size() && std::isdigit(format[i])) {
                        mappedIndex = mappedIndex * 10 + (format[i] - '0');
                        i++;
                    }
                    if (i >= format.size() || format[i] != ']') { throw std::runtime_error("missing closing bracket"); }
                    i++;
                    st.argMap[argIndex] = mappedIndex;
                    argIndex = mappedIndex;
                } else {
                    argIndex = st.argIndex(argIndex);
                }
            }

            std::string flags;
            while (i < format.size()) {
                switch (format[i]) {
                    case '+':
                    case '-':
                    case ' ':
                    case '0':
                    case '#':
                        flags += format[i];
                        i++;
                        break;
                    default:
                        goto end_flags;
                }
            }
        end_flags:

            int width = -1;
            if (std::isdigit(format[i])) {
                width = 0;
                while (i < format.size() && std::isdigit(format[i])) {
                    width = width * 10 + (format[i] - '0');
                    i++;
                }
            } else if (format[i] == '*') {
                i++;
                if (st.argNum >= st.args.size()) { throw std::runtime_error(fmt_ErrMissingArg.msg); }
                auto w = st.args[st.argNum++];
                try {
                    width = std::any_cast<int>(w);
                } catch (const std::bad_any_cast &) { throw std::runtime_error("width must be integer"); }
            }

            int prec = -1;
            if (format[i] == '.') {
                i++;
                if (std::isdigit(format[i])) {
                    prec = 0;
                    while (i < format.size() && std::isdigit(format[i])) {
                        prec = prec * 10 + (format[i] - '0');
                        i++;
                    }
                } else if (format[i] == '*') {
                    i++;
                    if (st.argNum >= st.args.size()) { throw std::runtime_error(fmt_ErrMissingArg.msg); }
                    auto p = st.args[st.argNum++];
                    try {
                        prec = std::any_cast<int>(p);
                    } catch (const std::bad_any_cast &) { throw std::runtime_error("precision must be integer"); }
                }
            }

            if (argIndex == -1) { argIndex = st.argNum++; }

            if (argIndex >= static_cast<int>(st.args.size())) { throw std::runtime_error(fmt_ErrMissingArg.msg); }

            char verb = 0;
            if (i < format.size()) {
                verb = format[i++];
            } else {
                throw std::runtime_error(fmt_ErrInvalidVerb.msg);
            }

            const auto &arg = st.args[argIndex];
            formatValue(st, arg, verb, flags, width, prec);
        }
    }

    static void formatValue(fmt_State &st, const std::any &arg, char verb, const std::string &flags, int width, int prec) {
        std::ostringstream oss;
        bool               handled = false;

        if (arg.type() == typeid(int)) {
            int value = std::any_cast<int>(arg);
            handled = formatInt(oss, value, verb, flags, width);
        } else if (arg.type() == typeid(double)) {
            double value = std::any_cast<double>(arg);
            handled = formatFloat(oss, value, verb, flags, width, prec);
        } else if (arg.type() == typeid(const char *)) {
            const char *value = std::any_cast<const char *>(arg);
            handled = formatString(oss, std::string(value), verb, flags, width, prec);
        } else if (arg.type() == typeid(std::string)) {
            std::string value = std::any_cast<std::string>(arg);
            handled = formatString(oss, value, verb, flags, width, prec);
        } else if (arg.type() == typeid(bool)) {
            bool value = std::any_cast<bool>(arg);
            handled = formatBool(oss, value, verb, flags, width);
        } else if (arg.type() == typeid(char)) {
            char value = std::any_cast<char>(arg);
            handled = formatChar(oss, value, verb, flags, width);
        }

        if (!handled) { throw std::runtime_error(fmt_ErrInvalidVerb.msg); }

        st.writeString(oss.str());
    }

    static bool formatInt(std::ostream &os, int value, char verb, const std::string &flags, int width) {
        int  base = 10;
        bool showBase = false;
        bool forceSign = false;

        switch (verb) {
            case 'b':
                base = 2;
                showBase = !flags.empty();
                break;
            case 'o':
                base = 8;
                showBase = flags.find('#') != std::string::npos;
                break;
            case 'x':
            case 'X':
                base = 16;
                showBase = flags.find('#') != std::string::npos;
                break;
            case 'd':
            case 'v':
                break;
            default:
                return false;
        }

        if (flags.find('+') != std::string::npos) { forceSign = true; }

        char         sign = 0;
        unsigned int uvalue = value;
        if (value < 0) {
            sign = '-';
            uvalue = -value;
        } else if (forceSign) {
            sign = '+';
        } else if (flags.find(' ') != std::string::npos) {
            sign = ' ';
        }

        std::string digits;
        const char *digitChars = (verb == 'X') ? "0123456789ABCDEF" : "0123456789abcdef";

        do {
            digits.push_back(digitChars[uvalue % base]);
            uvalue /= base;
        } while (uvalue > 0);

        std::reverse(digits.begin(), digits.end());

        std::string prefix;
        if (sign) { prefix.push_back(sign); }
        if (showBase) {
            switch (verb) {
                case 'b':
                    prefix += "0b";
                    break;
                case 'o':
                    prefix += "0";
                    break;
                case 'x':
                    prefix += "0x";
                    break;
                case 'X':
                    prefix += "0X";
                    break;
            }
        }

        int totalLen = prefix.size() + digits.size();
        int padLen = (width > totalLen) ? width - totalLen : 0;

        if (flags.find('-') != std::string::npos) {
            os << prefix << digits;
            if (padLen > 0) { os << std::string(padLen, ' '); }
        } else if (flags.find('0') != std::string::npos && !prefix.empty()) {
            os << prefix;
            if (padLen > 0) { os << std::string(padLen, '0'); }
            os << digits;
        } else {
            if (padLen > 0) { os << std::string(padLen, ' '); }
            os << prefix << digits;
        }

        return true;
    }

    static bool formatFloat(std::ostream &os, double value, char verb, const std::string &flags, int width, int prec) {
        if (prec < 0) prec = 6;

        switch (verb) {
            case 'f':
            case 'F':
                os << std::fixed;
                break;
            case 'e':
            case 'E':
                os << std::scientific;
                if (verb == 'E') os << std::uppercase;
                break;
            case 'g':
            case 'G':
                os << std::defaultfloat;
                if (verb == 'G') os << std::uppercase;
                break;
            case 'v':// 默认格式
                break;
            default:
                return false;
        }

        if (flags.find('+') != std::string::npos) { os << std::showpos; }

        if (width > 0) { os << std::setw(width); }
        os << std::setprecision(prec);

        os << value;
        return true;
    }

    static bool formatString(std::ostream &os, const std::string &value, char verb, const std::string &flags, int width, int prec) {
        if (verb != 's' && verb != 'q' && verb != 'v') { return false; }

        std::string s = value;
        if (prec >= 0 && s.size() > static_cast<size_t>(prec)) { s = s.substr(0, prec); }

        if (verb == 'q') { s = "\"" + s + "\""; }

        if (width > 0) {
            int padLen = width - s.size();
            if (padLen > 0) {
                if (flags.find('-') != std::string::npos) {
                    os << s << std::string(padLen, ' ');
                } else {
                    os << std::string(padLen, ' ') << s;
                }
                return true;
            }
        }

        os << s;
        return true;
    }

    static bool formatBool(std::ostream &os, bool value, char verb, const std::string &flags, int width) {
        std::string s;
        switch (verb) {
            case 't':
            case 'v':
                s = value ? "true" : "false";
                break;
            case 'd':
                s = value ? "1" : "0";
                break;
            default:
                return false;
        }

        if (width > 0) {
            int padLen = width - s.size();
            if (padLen > 0) {
                if (flags.find('-') != std::string::npos) {
                    os << s << std::string(padLen, ' ');
                } else {
                    os << std::string(padLen, ' ') << s;
                }
                return true;
            }
        }

        os << s;
        return true;
    }

    static bool formatChar(std::ostream &os, char value, char verb, const std::string &flags, int width) {
        if (verb != 'c' && verb != 'v') { return false; }

        if (width > 1) {
            int padLen = width - 1;
            if (padLen > 0) {
                if (flags.find('-') != std::string::npos) {
                    os << value << std::string(padLen, ' ');
                } else {
                    os << std::string(padLen, ' ') << value;
                }
                return true;
            }
        }

        os << value;
        return true;
    }

    static void printArgs(fmt_State &st, const std::vector<std::any> &args) {
        for (size_t i = 0; i < args.size(); i++) {
            if (i > 0) { st.writeString(" "); }
            formatValue(st, args[i], 'v', "", -1, -1);
        }
    }
};

fmt fmt;

void testFmtFunctions() {
    std::string name = "Alice";
    int         age = 30;
    fmt.Printf("Hello %s, you are %d years old\n", {name, age});

    double pi = 3.1415926535;
    fmt.Printf("PI = %.2f\n", {pi});

    fmt.Printf("|%-10s|%10d|\n", {"Hello", 123});

    fmt.Printf("Mixed: %s %d %.3f %t\n", {"text", 42, 3.14159, true});

    std::string s = fmt.Sprintf("%s-%d-%.1f", {"test", 123, 45.6});
    fmt.Println({s});

    try {
        int a = 10, b = 0;
        if (b == 0) { throw fmt.Errorf("division by zero: %d/%d", {a, b}); }
    } catch (const std::exception &e) { fmt.Println({e.what()}); }
}
