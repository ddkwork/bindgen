#pragma once

#include <vector>
#include <cctype>
#include <cstring>
#include <algorithm>
#include <string>
#include <stdexcept>
#include <functional>

struct Buffer {
private:
    std::vector<uint8_t> buf_;
    size_t               off_ = 0;// 读偏移量

    void checkIndex(size_t i) const {
        if (i >= Len()) throw std::out_of_range("index out of range");
    }

    void ensureCapacity(size_t add) {
        if (Len() + add > Cap()) { grow(Len() + add - Cap()); }
    }

public:
    Buffer() = default;

    explicit Buffer(size_t size) : buf_(size) {}

    explicit Buffer(const std::string &str) : buf_(str.begin(), str.end()) {}

    explicit Buffer(const void *p, size_t len) : buf_(static_cast<const uint8_t *>(p), static_cast<const uint8_t *>(p) + len) {}

    size_t Len() const { return buf_.size() - off_; }

    size_t Cap() const { return buf_.capacity(); }

    void Reset() {
        off_ = 0;
        buf_.clear();
    }

    const uint8_t *Bytes() const { return buf_.data() + off_; }

    std::string String() const { return std::string(buf_.begin() + off_, buf_.end()); }

    uint8_t NextByte() {
        if (off_ >= buf_.size()) throw std::out_of_range("end of buffer");
        return buf_[off_++];
    }

    void ReadBytes(uint8_t *p, size_t n) {
        if (off_ + n > buf_.size()) throw std::out_of_range("end of buffer");
        std::memcpy(p, buf_.data() + off_, n);
        off_ += n;
    }

    Buffer Next(size_t n) {
        if (off_ + n > buf_.size()) throw std::out_of_range("end of buffer");
        Buffer result(buf_.data() + off_, n);
        off_ += n;
        return result;
    }

    void WriteByte(uint8_t b) {
        ensureCapacity(1);
        buf_.push_back(b);
    }

    void WriteBytes(const uint8_t *p, size_t n) {
        ensureCapacity(n);
        buf_.insert(buf_.end(), p, p + n);
    }

    void WriteString(const std::string &s) { WriteBytes(reinterpret_cast<const uint8_t *>(s.data()), s.size()); }

    void WriteBuffer(const Buffer &b) { WriteBytes(b.Bytes(), b.Len()); }

    void Grow(size_t n) {
        if (n <= Cap() - Len()) return;

        std::vector<uint8_t> newBuf;
        newBuf.reserve(Len() + n);

        if (Len() > 0) { newBuf.insert(newBuf.end(), buf_.begin() + off_, buf_.end()); }

        buf_ = std::move(newBuf);
        off_ = 0;
    }

    uint8_t ReadByte() { return NextByte(); }

    int ReadByte() {// 兼容Go的返回int类型
        if (off_ >= buf_.size()) return -1;
        return static_cast<int>(buf_[off_++]);
    }

    size_t Read(void *p, size_t n) {
        const size_t avail = Len();
        if (n > avail) n = avail;
        ReadBytes(static_cast<uint8_t *>(p), n);
        return n;
    }

    void UnreadByte() {
        if (off_ == 0) throw std::runtime_error("nothing to unread");
        off_--;
    }

    void Truncate(size_t n) {
        if (n > Len()) n = Len();
        buf_.resize(off_ + n);
    }

    size_t WriteTo(std::ostream &os) {
        if (Len() == 0) return 0;
        os.write(reinterpret_cast<const char *>(buf_.data() + off_), Len());
        const size_t written = Len();
        off_ = buf_.size();// 所有数据已读
        return written;
    }

    size_t ReadFrom(std::istream &is) {
        size_t       totalRead = 0;
        const size_t initial = buf_.size();

        while (is) {
            const size_t oldCap = Cap();
            const size_t newCap = (oldCap == 0) ? 4096 : oldCap * 2;
            Grow(newCap - Len());

            is.read(reinterpret_cast<char *>(buf_.data() + buf_.size()), Cap() - Len());
            const size_t read = is.gcount();

            buf_.resize(buf_.size() + read);
            totalRead += read;
        }

        return totalRead;
    }

    int Compare(const Buffer &b) const {
        const size_t minLen = std::min(Len(), b.Len());
        int          cmp = std::memcmp(Bytes(), b.Bytes(), minLen);
        return (cmp != 0) ? cmp : static_cast<int>(Len()) - static_cast<int>(b.Len());
    }

    bool Contains(const Buffer &sub) const {
        if (sub.Len() == 0) return true;
        if (Len() < sub.Len()) return false;

        const uint8_t *start = Bytes();
        const uint8_t *end = Bytes() + Len() - sub.Len() + 1;

        for (const uint8_t *p = start; p < end; ++p) {
            if (std::memcmp(p, sub.Bytes(), sub.Len()) == 0) { return true; }
        }
        return false;
    }

    bool HasPrefix(const Buffer &prefix) const { return Len() >= prefix.Len() && std::memcmp(Bytes(), prefix.Bytes(), prefix.Len()) == 0; }

    bool HasSuffix(const Buffer &suffix) const { return Len() >= suffix.Len() && std::memcmp(Bytes() + Len() - suffix.Len(), suffix.Bytes(), suffix.Len()) == 0; }

    ssize_t Index(const Buffer &sub) const {
        if (sub.Len() == 0) return 0;
        if (Len() < sub.Len()) return -1;

        const uint8_t *start = Bytes();
        const uint8_t *end = Bytes() + Len() - sub.Len() + 1;

        for (const uint8_t *p = start; p < end; ++p) {
            if (std::memcmp(p, sub.Bytes(), sub.Len()) == 0) { return p - start; }
        }
        return -1;
    }

    Buffer Slice(size_t start, size_t end) const {
        if (start > Len()) return Buffer();
        if (end > Len()) end = Len();
        if (start > end) return Buffer();

        return Buffer(Bytes() + start, end - start);
    }

    static Buffer Join(const std::vector<Buffer> &buffers, const Buffer &sep) {
        if (buffers.empty()) return Buffer();

        size_t totalLen = sep.Len() * (buffers.size() - 1);
        for (const auto &b: buffers) { totalLen += b.Len(); }

        Buffer result(totalLen);
        result.Grow(totalLen);

        bool first = true;
        for (const auto &b: buffers) {
            if (!first) { result.WriteBuffer(sep); }
            first = false;
            result.WriteBuffer(b);
        }

        return result;
    }

    static Buffer Repeat(const Buffer &b, size_t count) {
        if (count == 0) return Buffer();
        if (count == 1) return b;

        Buffer result;
        result.Grow(b.Len() * count);

        for (size_t i = 0; i < count; i++) { result.WriteBuffer(b); }

        return result;
    }
};
