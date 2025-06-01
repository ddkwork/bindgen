#pragma once
#include <vector>
#include <cstdint>
#include <string>
#include <stdexcept>
#include <cstring>
#include <iostream>
#include <memory>
#include <algorithm>

struct Error {
    std::string msg;

    Error() {}

    Error(const std::string &m) : msg(m) {}

    std::string Error() const { return msg; }

    operator bool() const { return !msg.empty(); }
};

Error ErrChecksum = Error("checksum error");
Error ErrHeader = Error("invalid header");
Error ErrShortBuffer = Error("buffer too short");

const uint8_t GZIP_ID1 = 0x1f;
const uint8_t GZIP_ID2 = 0x8b;
const uint8_t DEFLATE = 8;

const uint8_t FTEXT = 1 << 0;
const uint8_t FHCRC = 1 << 1;
const uint8_t FEXTRA = 1 << 2;
const uint8_t FNAME = 1 << 3;
const uint8_t FCOMMENT = 1 << 4;

struct Header {
    std::string Comment;
    std::string Extra;
    std::string Name;
    uint32_t    ModTime;// Unix时间戳
    uint8_t     OS;     // 操作系统类型
    uint8_t     Flags;  // 标志位

    Header() : ModTime(0), OS(255), Flags(0) {}
};

struct Writer {
    std::vector<uint8_t> buf;
    Header               header;
    int                  level = 6;// 默认压缩级别
    bool                 closed = false;
    uint32_t             size = 0;      // 写入的未压缩字节数
    uint32_t             compressed = 0;// 压缩后字节数
    uint32_t             digest = 0;    // CRC32校验和

    Error Init(std::vector<uint8_t> &dst) {
        buf = dst;
        return WriteHeader();
    }

    Error WriteHeader() {
        buf.push_back(GZIP_ID1);
        buf.push_back(GZIP_ID2);    // ID
        buf.push_back(DEFLATE);     // 压缩方法
        buf.push_back(header.Flags);// 标志位

        buf.push_back(header.ModTime & 0xFF);
        buf.push_back((header.ModTime >> 8) & 0xFF);
        buf.push_back((header.ModTime >> 16) & 0xFF);
        buf.push_back((header.ModTime >> 24) & 0xFF);

        buf.push_back(0);        // 额外标志 (通常为0)
        buf.push_back(header.OS);// 操作系统

        if (header.Flags & FEXTRA) {
            if (header.Extra.size() > 0xFFFF) { return Error("extra field too long"); }
            uint16_t len = header.Extra.size();
            buf.push_back(len & 0xFF);
            buf.push_back(len >> 8);
            buf.insert(buf.end(), header.Extra.begin(), header.Extra.end());
        }

        if (header.Flags & FNAME) {
            buf.insert(buf.end(), header.Name.begin(), header.Name.end());
            buf.push_back(0);// 以null结尾
        }

        if (header.Flags & FCOMMENT) {
            buf.insert(buf.end(), header.Comment.begin(), header.Comment.end());
            buf.push_back(0);// 以null结尾
        }

        if (header.Flags & FHCRC) {
            buf.push_back(0);
            buf.push_back(0);
        }

        return Error();
    }

    Error Write(const std::vector<uint8_t> &data) {
        if (closed) return Error("write after close");

        digest = UpdateCRC32(digest, data);
        size += data.size();

        std::vector<uint8_t> compressedData = Deflate(data, level);
        compressed += compressedData.size();

        buf.insert(buf.end(), compressedData.begin(), compressedData.end());
        return Error();
    }

    Error Close() {
        if (closed) return Error();
        closed = true;

        WriteTrailer();
        return Error();
    }

    void WriteTrailer() {
        buf.push_back(digest & 0xFF);
        buf.push_back((digest >> 8) & 0xFF);
        buf.push_back((digest >> 16) & 0xFF);
        buf.push_back((digest >> 24) & 0xFF);

        buf.push_back(size & 0xFF);
        buf.push_back((size >> 8) & 0xFF);
        buf.push_back((size >> 16) & 0xFF);
        buf.push_back((size >> 24) & 0xFF);
    }

    Error Reset(std::vector<uint8_t> &dst) {
        buf = dst;
        size = 0;
        compressed = 0;
        digest = 0;
        closed = false;
        return WriteHeader();
    }

    float Ratio() const {
        if (size == 0) return 1.0f;
        return static_cast<float>(compressed) / size;
    }

private:
    std::vector<uint8_t> Deflate(const std::vector<uint8_t> &data, int level) {
        std::vector<uint8_t> result;

        result.push_back(0x01);// BFINAL=1, BTYPE=00

        size_t pos = 0;
        while (pos < data.size()) {
            size_t blockSize = std::min<size_t>(data.size() - pos, 65535);

            result.push_back(blockSize & 0xFF);
            result.push_back(blockSize >> 8);

            uint16_t nlen = ~blockSize;
            result.push_back(nlen & 0xFF);
            result.push_back(nlen >> 8);

            result.insert(result.end(), data.begin() + pos, data.begin() + pos + blockSize);
            pos += blockSize;
        }

        return result;
    }

    uint32_t UpdateCRC32(uint32_t crc, const std::vector<uint8_t> &data) {
        static const uint32_t poly = 0xEDB88320;
        static uint32_t       table[256];
        static bool           initialized = false;

        if (!initialized) {
            for (uint32_t i = 0; i < 256; i++) {
                uint32_t c = i;
                for (int j = 0; j < 8; j++) {
                    if (c & 1) c = (c >> 1) ^ poly;
                    else
                        c >>= 1;
                }
                table[i] = c;
            }
            initialized = true;
        }

        crc = ~crc;
        for (uint8_t b: data) { crc = table[(crc ^ b) & 0xFF] ^ (crc >> 8); }
        return ~crc;
    }
};

struct Reader {
    std::vector<uint8_t> src;
    size_t               pos = 0;
    Header               header;
    bool                 closed = false;
    uint32_t             digest = 0;    // CRC32校验和
    uint32_t             size = 0;      // 解压后字节数
    uint32_t             compressed = 0;// 压缩数据大小

    Error Init(const std::vector<uint8_t> &data) {
        src = data;
        pos = 0;
        return ReadHeader();
    }

    Error ReadHeader() {
        if (src.size() < 10) return ErrHeader;

        if (src[0] != GZIP_ID1 || src[1] != GZIP_ID2) { return ErrHeader; }

        if (src[2] != DEFLATE) { return Error("unsupported compression method"); }

        header.Flags = src[3];
        header.ModTime =
                (static_cast<uint32_t>(src[4]) << 0) | (static_cast<uint32_t>(src[5]) << 8) | (static_cast<uint32_t>(src[6]) << 16) | (static_cast<uint32_t>(src[7]) << 24);

        header.OS = src[9];
        pos = 10;

        if (header.Flags & FEXTRA) {
            if (pos + 2 > src.size()) return ErrHeader;
            uint16_t xlen = src[pos] | (src[pos + 1] << 8);
            pos += 2;

            if (pos + xlen > src.size()) return ErrHeader;
            header.Extra = std::string(src.begin() + pos, src.begin() + pos + xlen);
            pos += xlen;
        }

        if (header.Flags & FNAME) {
            size_t start = pos;
            while (pos < src.size() && src[pos] != 0) pos++;
            if (pos >= src.size()) return ErrHeader;
            header.Name = std::string(src.begin() + start, src.begin() + pos);
            pos++;// 跳过null
        }

        if (header.Flags & FCOMMENT) {
            size_t start = pos;
            while (pos < src.size() && src[pos] != 0) pos++;
            if (pos >= src.size()) return ErrHeader;
            header.Comment = std::string(src.begin() + start, src.begin() + pos);
            pos++;// 跳过null
        }

        if (header.Flags & FHCRC) {
            if (pos + 2 > src.size()) return ErrHeader;
            pos += 2;
        }

        return Error();
    }

    Error Read(std::vector<uint8_t> &dst) {
        if (closed) return Error("read after close");

        auto [data, err] = Inflate();
        if (err) return err;

        digest = UpdateCRC32(digest, data);
        size += data.size();

        dst.insert(dst.end(), data.begin(), data.end());
        return Error();
    }

    Error Close() {
        if (closed) return Error();
        closed = true;

        if (pos + 8 > src.size()) return ErrChecksum;

        uint32_t storedCRC = (static_cast<uint32_t>(src[pos]) << 0) | (static_cast<uint32_t>(src[pos + 1]) << 8) | (static_cast<uint32_t>(src[pos + 2]) << 16) |
                             (static_cast<uint32_t>(src[pos + 3]) << 24);
        pos += 4;

        uint32_t storedSize = (static_cast<uint32_t>(src[pos]) << 0) | (static_cast<uint32_t>(src[pos + 1]) << 8) | (static_cast<uint32_t>(src[pos + 2]) << 16) |
                              (static_cast<uint32_t>(src[pos + 3]) << 24);
        pos += 4;

        if (storedSize != size) { return Error("size mismatch"); }

        if (storedCRC != digest) { return ErrChecksum; }

        return Error();
    }

    Error Reset(const std::vector<uint8_t> &data) {
        src = data;
        pos = 0;
        digest = 0;
        size = 0;
        compressed = 0;
        closed = false;
        return ReadHeader();
    }

    float Ratio() const {
        if (compressed == 0) return 1.0f;
        return static_cast<float>(size) / compressed;
    }

private:
    std::pair<std::vector<uint8_t>, Error> Inflate() {
        std::vector<uint8_t> result;

        while (pos < src.size()) {
            if (pos + 1 > src.size()) { return {result, ErrShortBuffer}; }

            uint8_t bfinal = src[pos] & 0x01;
            uint8_t btype = (src[pos] >> 1) & 0x03;
            pos++;

            if (btype == 0) {// 无压缩
                if (pos + 4 > src.size()) return {result, ErrShortBuffer};

                uint16_t len = src[pos] | (src[pos + 1] << 8);
                uint16_t nlen = src[pos + 2] | (src[pos + 3] << 8);
                pos += 4;

                if (len != (uint16_t) (~nlen)) { return {result, Error("length mismatch")}; }

                if (pos + len > src.size()) { return {result, ErrShortBuffer}; }

                result.insert(result.end(), src.begin() + pos, src.begin() + pos + len);
                pos += len;
                compressed += len + 4;// 包括头部
            } else {
                return {result, Error("compressed blocks not supported in this implementation")};
            }

            if (bfinal) break;
        }

        return {result, Error()};
    }

    uint32_t UpdateCRC32(uint32_t crc, const std::vector<uint8_t> &data) {
        static const uint32_t poly = 0xEDB88320;
        static uint32_t       table[256];
        static bool           initialized = false;

        if (!initialized) {
            for (uint32_t i = 0; i < 256; i++) {
                uint32_t c = i;
                for (int j = 0; j < 8; j++) {
                    if (c & 1) c = (c >> 1) ^ poly;
                    else
                        c >>= 1;
                }
                table[i] = c;
            }
            initialized = true;
        }

        crc = ~crc;
        for (uint8_t b: data) { crc = table[(crc ^ b) & 0xFF] ^ (crc >> 8); }
        return ~crc;
    }
};

struct ZlibWriter {
    std::vector<uint8_t> buf;
    int                  level = 6;
    bool                 closed = false;
    uint32_t             digest = 0;// Adler32校验和
    uint32_t             size = 0;

    Error Init(std::vector<uint8_t> &dst) {
        buf = dst;
        return WriteHeader();
    }

    Error WriteHeader() {
        uint8_t cmf = 0x78;// CM=8 (DEFLATE), CINFO=7 (32K窗口)
        uint8_t flg = 0;

        if (level == 1) flg = 0x01;// 最快压缩
        else if (level == 9)
            flg = 0x03;// 最佳压缩
        else
            flg = 0x02;// 默认

        while ((cmf * 256 + flg) % 31 != 0) { flg++; }

        buf.push_back(cmf);
        buf.push_back(flg);
        return Error();
    }

    Error Write(const std::vector<uint8_t> &data) {
        if (closed) return Error("write after close");

        digest = UpdateAdler32(digest, data);
        size += data.size();

        std::vector<uint8_t> compressedData = Deflate(data, level);

        buf.insert(buf.end(), compressedData.begin(), compressedData.end());
        return Error();
    }

    Error Close() {
        if (closed) return Error();
        closed = true;

        buf.push_back((digest >> 24) & 0xFF);
        buf.push_back((digest >> 16) & 0xFF);
        buf.push_back((digest >> 8) & 0xFF);
        buf.push_back(digest & 0xFF);
        return Error();
    }

private:
    std::vector<uint8_t> Deflate(const std::vector<uint8_t> &data, int level) {
        std::vector<uint8_t> result;

        result.push_back(0x01);// BFINAL=1, BTYPE=00

        size_t pos = 0;
        while (pos < data.size()) {
            size_t blockSize = std::min<size_t>(data.size() - pos, 65535);

            result.push_back(blockSize & 0xFF);
            result.push_back(blockSize >> 8);

            uint16_t nlen = ~blockSize;
            result.push_back(nlen & 0xFF);
            result.push_back(nlen >> 8);

            result.insert(result.end(), data.begin() + pos, data.begin() + pos + blockSize);
            pos += blockSize;
        }

        return result;
    }

    uint32_t UpdateAdler32(uint32_t adler, const std::vector<uint8_t> &data) {
        uint32_t       s1 = adler & 0xFFFF;
        uint32_t       s2 = (adler >> 16) & 0xFFFF;
        const uint32_t MOD_ADLER = 65521;

        for (uint8_t b: data) {
            s1 = (s1 + b) % MOD_ADLER;
            s2 = (s2 + s1) % MOD_ADLER;
        }
        return (s2 << 16) | s1;
    }
};

struct ZlibReader {
    std::vector<uint8_t> src;
    size_t               pos = 0;
    bool                 closed = false;
    uint32_t             digest = 1;// Adler32初始值
    uint32_t             size = 0;

    Error Init(const std::vector<uint8_t> &data) {
        src = data;
        pos = 0;
        return ReadHeader();
    }

    Error ReadHeader() {
        if (src.size() < 2) return ErrHeader;

        uint8_t cmf = src[0];
        uint8_t flg = src[1];
        pos = 2;

        if ((cmf & 0x0F) != 8) { return Error("unsupported compression method"); }

        if ((cmf * 256 + flg) % 31 != 0) { return ErrHeader; }

        if (flg & 0x20) {
            if (pos + 4 > src.size()) return ErrHeader;
            pos += 4;
        }

        return Error();
    }

    Error Read(std::vector<uint8_t> &dst) {
        if (closed) return Error("read after close");

        auto [data, err] = Inflate();
        if (err) return err;

        digest = UpdateAdler32(digest, data);
        size += data.size();

        dst.insert(dst.end(), data.begin(), data.end());
        return Error();
    }

    Error Close() {
        if (closed) return Error();
        closed = true;

        if (pos + 4 > src.size()) return ErrChecksum;

        uint32_t storedAdler = (static_cast<uint32_t>(src[pos]) << 24) | (static_cast<uint32_t>(src[pos + 1]) << 16) | (static_cast<uint32_t>(src[pos + 2]) << 8) |
                               static_cast<uint32_t>(src[pos + 3]);
        pos += 4;

        if (storedAdler != digest) { return ErrChecksum; }

        return Error();
    }

private:
    std::pair<std::vector<uint8_t>, Error> Inflate() {
        std::vector<uint8_t> result;

        while (pos < src.size()) {
            if (pos + 1 > src.size()) { return {result, ErrShortBuffer}; }

            uint8_t bfinal = src[pos] & 0x01;
            uint8_t btype = (src[pos] >> 1) & 0x03;
            pos++;

            if (btype == 0) {// 无压缩
                if (pos + 4 > src.size()) return {result, ErrShortBuffer};

                uint16_t len = src[pos] | (src[pos + 1] << 8);
                uint16_t nlen = src[pos + 2] | (src[pos + 3] << 8);
                pos += 4;

                if (len != (uint16_t) (~nlen)) { return {result, Error("length mismatch")}; }

                if (pos + len > src.size()) { return {result, ErrShortBuffer}; }

                result.insert(result.end(), src.begin() + pos, src.begin() + pos + len);
                pos += len;
            } else {
                return {result, Error("compressed blocks not supported in this implementation")};
            }

            if (bfinal) break;
        }

        return {result, Error()};
    }

    uint32_t UpdateAdler32(uint32_t adler, const std::vector<uint8_t> &data) {
        uint32_t       s1 = adler & 0xFFFF;
        uint32_t       s2 = (adler >> 16) & 0xFFFF;
        const uint32_t MOD_ADLER = 65521;

        for (uint8_t b: data) {
            s1 = (s1 + b) % MOD_ADLER;
            s2 = (s2 + s1) % MOD_ADLER;
        }
        return (s2 << 16) | s1;
    }
};

std::pair<std::vector<uint8_t>, Error> CompressGzip(const std::vector<uint8_t> &data, int level = 6, const Header &header = Header()) {
    std::vector<uint8_t> result
