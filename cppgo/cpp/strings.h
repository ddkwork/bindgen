#pragma once

#include "buffer.h"// 包含之前实现的Buffer
#include <cctype>
#include <cstring>
#include <vector>
#include <algorithm>
#include <functional>
#include <locale>

struct Strings {
    Strings() = delete;

    static int Compare(const Buffer &a, const Buffer &b) { return a.Compare(b); }

    static bool Contains(const Buffer &s, const Buffer &substr) { return s.Contains(substr); }

    static bool ContainsAny(const Buffer &s, const Buffer &chars) {
        for (size_t i = 0; i < s.Len(); i++) {
            for (size_t j = 0; j < chars.Len(); j++) {
                if (s.Bytes()[i] == chars.Bytes()[j]) { return true; }
            }
        }
        return false;
    }

    static size_t Count(const Buffer &s, const Buffer &substr) {
        if (substr.Len() == 0) return s.Len() + 1;

        size_t count = 0;
        size_t pos = 0;

        while ((pos = static_cast<size_t>(Index(s, substr, pos))) != (size_t) -1) {
            count++;
            pos += substr.Len();
        }

        return count;
    }

    static bool EqualFold(const Buffer &a, const Buffer &b) {
        if (a.Len() != b.Len()) return false;

        for (size_t i = 0; i < a.Len(); i++) {
            if (std::tolower(a.Bytes()[i]) != std::tolower(b.Bytes()[i])) { return false; }
        }
        return true;
    }

    static std::vector<Buffer> Fields(const Buffer &s) {
        std::vector<Buffer> result;
        size_t              start = 0;
        bool                inField = false;

        for (size_t i = 0; i < s.Len(); i++) {
            uint8_t c = s.Bytes()[i];
            bool    isSpace = std::isspace(c);

            if (inField && isSpace) {
                result.push_back(s.Slice(start, i));
                inField = false;
            } else if (!inField && !isSpace) {
                start = i;
                inField = true;
            }
        }

        if (inField) { result.push_back(s.Slice(start, s.Len())); }

        return result;
    }

    static bool HasPrefix(const Buffer &s, const Buffer &prefix) { return s.HasPrefix(prefix); }

    static bool HasSuffix(const Buffer &s, const Buffer &suffix) { return s.HasSuffix(suffix); }

    static ssize_t Index(const Buffer &s, const Buffer &substr) { return s.Index(substr); }

    static ssize_t Index(const Buffer &s, const Buffer &substr, size_t offset) {
        if (substr.Len() == 0) return 0;
        if (s.Len() < substr.Len()) return -1;

        const uint8_t *start = s.Bytes() + offset;
        const uint8_t *end = s.Bytes() + s.Len() - substr.Len() + 1;

        if (start > end) return -1;

        for (const uint8_t *p = start; p < end; ++p) {
            if (std::memcmp(p, substr.Bytes(), substr.Len()) == 0) { return p - s.Bytes(); }
        }
        return -1;
    }

    static ssize_t IndexAny(const Buffer &s, const Buffer &chars) {
        for (size_t i = 0; i < s.Len(); i++) {
            for (size_t j = 0; j < chars.Len(); j++) {
                if (s.Bytes()[i] == chars.Bytes()[j]) { return i; }
            }
        }
        return -1;
    }

    static Buffer Join(const std::vector<Buffer> &elems, const Buffer &sep) { return Buffer::Join(elems, sep); }

    static ssize_t LastIndex(const Buffer &s, const Buffer &substr) {
        if (substr.Len() == 0) return s.Len();
        if (s.Len() < substr.Len()) return -1;

        for (ssize_t i = s.Len() - substr.Len(); i >= 0; i--) {
            if (std::memcmp(s.Bytes() + i, substr.Bytes(), substr.Len()) == 0) { return i; }
        }
        return -1;
    }

    static ssize_t LastIndexAny(const Buffer &s, const Buffer &chars) {
        for (ssize_t i = s.Len() - 1; i >= 0; i--) {
            for (size_t j = 0; j < chars.Len(); j++) {
                if (s.Bytes()[i] == chars.Bytes()[j]) { return i; }
            }
        }
        return -1;
    }

    static Buffer Repeat(const Buffer &s, size_t count) { return Buffer::Repeat(s, count); }

    static Buffer Replace(const Buffer &s, const Buffer &old, const Buffer &newStr, int n = -1) {
        if (n == 0 || old.Len() == 0) { return s; }

        if (n < 0) {
            n = s.Len() / old.Len() + 1;// 足够大的值
        }

        Buffer  result;
        size_t  last = 0;
        int     count = 0;
        ssize_t pos = 0;

        while ((pos = Index(s, old, last)) != -1 && count < n) {
            result.WriteBuffer(s.Slice(last, pos));
            result.WriteBuffer(newStr);

            last = pos + old.Len();
            count++;
        }

        if (last < s.Len()) { result.WriteBuffer(s.Slice(last, s.Len())); }

        return result;
    }

    static Buffer ReplaceAll(const Buffer &s, const Buffer &old, const Buffer &newStr) { return Replace(s, old, newStr, -1); }

    static std::vector<Buffer> Split(const Buffer &s, const Buffer &sep) {
        if (sep.Len() == 0) { return SplitByChar(s); }

        std::vector<Buffer> result;
        size_t              start = 0;
        ssize_t             pos = 0;

        while ((pos = Index(s, sep, start)) != -1) {
            result.push_back(s.Slice(start, pos));
            start = pos + sep.Len();
        }

        result.push_back(s.Slice(start, s.Len()));

        return result;
    }

    static std::vector<Buffer> SplitN(const Buffer &s, const Buffer &sep, int n) {
        if (n <= 0) { return Split(s, sep); }

        if (sep.Len() == 0) { return SplitNByChar(s, n); }

        std::vector<Buffer> result;
        size_t              start = 0;
        int                 count = 0;

        while (count < n - 1) {
            ssize_t pos = Index(s, sep, start);
            if (pos == -1) break;

            result.push_back(s.Slice(start, pos));
            start = pos + sep.Len();
            count++;
        }

        result.push_back(s.Slice(start, s.Len()));
        return result;
    }

    static Buffer ToLower(const Buffer &s) {
        Buffer result(s.Len());
        for (size_t i = 0; i < s.Len(); i++) {
            uint8_t b = s.Bytes()[i];
            result.WriteByte(std::tolower(b));
        }
        return result;
    }

    static Buffer ToTitle(const Buffer &s) { return ToUpper(s); }

    static Buffer ToUpper(const Buffer &s) {
        Buffer result(s.Len());
        for (size_t i = 0; i < s.Len(); i++) {
            uint8_t b = s.Bytes()[i];
            result.WriteByte(std::toupper(b));
        }
        return result;
    }

    static Buffer Trim(const Buffer &s, const Buffer &cutset) {
        if (cutset.Len() == 0) { return s; }

        size_t start = 0;
        size_t end = s.Len();

        for (; start < end; start++) {
            if (!ContainsByte(cutset, s.Bytes()[start])) { break; }
        }

        for (; end > start; end--) {
            if (!ContainsByte(cutset, s.Bytes()[end - 1])) { break; }
        }

        return s.Slice(start, end);
    }

    static Buffer TrimFunc(const Buffer &s, std::function<bool(uint8_t)> f) {
        size_t start = 0;
        size_t end = s.Len();

        for (; start < end; start++) {
            if (!f(s.Bytes()[start])) { break; }
        }

        for (; end > start; end--) {
            if (!f(s.Bytes()[end - 1])) { break; }
        }

        return s.Slice(start, end);
    }

    static Buffer TrimLeft(const Buffer &s, const Buffer &cutset) {
        size_t start = 0;
        for (; start < s.Len(); start++) {
            if (!ContainsByte(cutset, s.Bytes()[start])) { break; }
        }
        return s.Slice(start, s.Len());
    }

    static Buffer TrimLeftFunc(const Buffer &s, std::function<bool(uint8_t)> f) {
        size_t start = 0;
        for (; start < s.Len(); start++) {
            if (!f(s.Bytes()[start])) { break; }
        }
        return s.Slice(start, s.Len());
    }

    static Buffer TrimPrefix(const Buffer &s, const Buffer &prefix) {
        if (s.HasPrefix(prefix)) { return s.Slice(prefix.Len(), s.Len()); }
        return s;
    }

    static Buffer TrimRight(const Buffer &s, const Buffer &cutset) {
        size_t end = s.Len();
        for (; end > 0; end--) {
            if (!ContainsByte(cutset, s.Bytes()[end - 1])) { break; }
        }
        return s.Slice(0, end);
    }

    static Buffer TrimRightFunc(const Buffer &s, std::function<bool(uint8_t)> f) {
        size_t end = s.Len();
        for (; end > 0; end--) {
            if (!f(s.Bytes()[end - 1])) { break; }
        }
        return s.Slice(0, end);
    }

    static Buffer TrimSpace(const Buffer &s) {
        return TrimFunc(s, [](uint8_t c) { return std::isspace(c); });
    }

    static Buffer TrimSuffix(const Buffer &s, const Buffer &suffix) {
        if (s.HasSuffix(suffix)) { return s.Slice(0, s.Len() - suffix.Len()); }
        return s;
    }

private:
    static bool ContainsByte(const Buffer &chars, uint8_t b) {
        for (size_t i = 0; i < chars.Len(); i++) {
            if (chars.Bytes()[i] == b) { return true; }
        }
        return false;
    }

    static std::vector<Buffer> SplitByChar(const Buffer &s) {
        std::vector<Buffer> result;
        for (size_t i = 0; i < s.Len(); i++) { result.push_back(s.Slice(i, i + 1)); }
        return result;
    }

    static std::vector<Buffer> SplitNByChar(const Buffer &s, int n) {
        std::vector<Buffer> result;
        size_t              count = 0;

        for (size_t i = 0; i < s.Len() && count < n - 1; i++, count++) { result.push_back(s.Slice(i, i + 1)); }

        if (count < s.Len()) { result.push_back(s.Slice(count, s.Len())); }

        return result;
    }
};
