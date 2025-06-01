#pragma once
#include <cstdint>
#include <vector>
#include <string>
#include <system_error>
#include <cerrno>
#include <cstring>
#include <algorithm>
#include <type_traits>
#include <concepts>
#include <memory>
#include <fstream>

struct Error {
    std::string msg;
    int         code;

    Error() : code(0) {}

    Error(const std::string &m) : msg(m), code(-1) {}

    Error(const std::string &m, int c) : msg(m), code(c) {}

    std::string Error() const { return code ? (msg + " (code: " + std::to_string(code) + ")") : msg; }

    bool IsTimeout() const { return code == ETIMEDOUT; }

    bool IsTemporary() const { return code == EAGAIN || code == EWOULDBLOCK || code == EINTR; }
};

Error ErrUnexpectedEOF = Error("unexpected EOF");
Error ErrShortWrite = Error("short write");
Error ErrShortBuffer = Error("short buffer");
Error EOF = Error("EOF");

struct Reader {
    using ReadFunc = std::function<Error(uint8_t *, size_t, size_t *)>;
    ReadFunc readFunc;

    Error Read(uint8_t *p, size_t len, size_t *n) const {
        if (!readFunc) return Error("reader not implemented");
        return readFunc(p, len, n);
    }

    Error ReadAll(std::vector<uint8_t> &data) const {
        const size_t bufSize = 512;
        uint8_t      buf[bufSize];

        while (true) {
            size_t n = 0;
            auto   err = Read(buf, bufSize, &n);
            if (n > 0) { data.insert(data.end(), buf, buf + n); }

            if (err.code) {
                if (err == EOF && data.size() > 0) return Error();
                return err;
            }

            if (n == 0) return Error();
        }
    }

    Error ReadFull(uint8_t *p, size_t len, size_t *n) const {
        *n = 0;
        while (*n < len) {
            size_t nn = 0;
            auto   err = Read(p + *n, len - *n, &nn);
            *n += nn;

            if (err.code) {
                if (err == EOF) err = ErrUnexpectedEOF;
                return err;
            }

            if (nn == 0) return ErrUnexpectedEOF;
        }
        return Error();
    }

    Error Copy(Writer &dst, int64_t &written) const {
        const size_t bufSize = 32 * 1024;
        uint8_t      buf[bufSize];
        written = 0;

        while (true) {
            size_t nr = 0;
            auto   err = Read(buf, bufSize, &nr);
            if (nr > 0) {
                size_t nw = 0;
                err = dst.Write(buf, nr, &nw);
                written += nw;

                if (err.code) return err;
                if (nw != nr) return ErrShortWrite;
            }

            if (err.code) {
                if (err == EOF) return Error();
                return err;
            }

            if (nr == 0) return Error();
        }
    }
};

struct Writer {
    using WriteFunc = std::function<Error(const uint8_t *, size_t, size_t *)>;
    WriteFunc writeFunc;

    Error Write(const uint8_t *p, size_t len, size_t *n) const {
        if (!writeFunc) return Error("writer not implemented");
        return writeFunc(p, len, n);
    }

    Error WriteString(const std::string &s, size_t *n) const { return Write(reinterpret_cast<const uint8_t *>(s.data()), s.size(), n); }
};

struct Closer {
    using CloseFunc = std::function<Error()>;
    CloseFunc closeFunc;

    Error Close() const {
        if (!closeFunc) return Error("closer not implemented");
        return closeFunc();
    }
};

struct Seeker {
    using SeekFunc = std::function<Error(int64_t, int, int64_t *)>;
    SeekFunc seekFunc;

    Error Seek(int64_t offset, int whence, int64_t *ret) const {
        if (!seekFunc) return Error("seeker not implemented");
        return seekFunc(offset, whence, ret);
    }
};

struct ReadWriter {
    Reader reader;
    Writer writer;

    Error Read(uint8_t *p, size_t len, size_t *n) const { return reader.Read(p, len, n); }

    Error Write(const uint8_t *p, size_t len, size_t *n) const { return writer.Write(p, len, n); }
};

struct ReadCloser {
    Reader reader;
    Closer closer;

    Error Read(uint8_t *p, size_t len, size_t *n) const { return reader.Read(p, len, n); }

    Error Close() const { return closer.Close(); }
};

struct WriteCloser {
    Writer writer;
    Closer closer;

    Error Write(const uint8_t *p, size_t len, size_t *n) const { return writer.Write(p, len, n); }

    Error Close() const { return closer.Close(); }
};

struct ReadWriteCloser {
    Reader reader;
    Writer writer;
    Closer closer;

    Error Read(uint8_t *p, size_t len, size_t *n) const { return reader.Read(p, len, n); }

    Error Write(const uint8_t *p, size_t len, size_t *n) const { return writer.Write(p, len, n); }

    Error Close() const { return closer.Close(); }
};

struct ByteReader {
    const uint8_t *data;
    size_t         len;
    size_t         pos;

    Error Read(uint8_t *p, size_t n, size_t *read) const {
        if (pos >= len) return EOF;

        size_t toRead = std::min(n, len - pos);
        std::memcpy(p, data + pos, toRead);
        pos += toRead;
        *read = toRead;

        return toRead < n ? EOF : Error();
    }
};

struct ByteWriter {
    uint8_t *data;
    size_t   len;
    size_t   pos;

    Error Write(const uint8_t *p, size_t n, size_t *written) const {
        if (pos >= len) return ErrShortWrite;

        size_t toWrite = std::min(n, len - pos);
        std::memcpy(data + pos, p, toWrite);
        pos += toWrite;
        *written = toWrite;

        return toWrite < n ? ErrShortWrite : Error();
    }
};

struct File {
    std::shared_ptr<std::fstream> fs;
    std::string                   name;

    Error Read(uint8_t *p, size_t len, size_t *n) const {
        if (!fs->is_open()) return Error("file not open");

        fs->read(reinterpret_cast<char *>(p), len);
        *n = fs->gcount();

        if (fs->eof() && *n > 0) return Error();
        if (fs->fail() && !fs->eof()) return Error("read error");
        return fs->eof() ? EOF : Error();
    }

    Error Write(const uint8_t *p, size_t len, size_t *n) const {
        if (!fs->is_open()) return Error("file not open");

        auto pos = fs->tellp();
        fs->write(reinterpret_cast<const char *>(p), len);
        *n = fs->tellp() - pos;

        return fs->fail() ? Error("write error") : Error();
    }

    Error Seek(int64_t offset, int whence, int64_t *ret) const {
        if (!fs->is_open()) return Error("file not open");

        std::ios_base::seekdir dir;
        switch (whence) {
            case 0:
                dir = std::ios::beg;
                break;// SEEK_SET
            case 1:
                dir = std::ios::cur;
                break;// SEEK_CUR
            case 2:
                dir = std::ios::end;
                break;// SEEK_END
            default:
                return Error("invalid whence");
        }

        fs->seekg(offset, dir);
        if (fs->fail()) return Error("seek error");

        *ret = fs->tellg();
        return Error();
    }

    Error Close() const {
        if (fs->is_open()) { fs->close(); }
        return Error();
    }

    static File Open(const std::string &name, const std::string &mode) {
        std::ios_base::openmode omode = std::ios::binary;
        if (mode.find('r') != std::string::npos) omode |= std::ios::in;
        if (mode.find('w') != std::string::npos) omode |= std::ios::out;
        if (mode.find('a') != std::string::npos) omode |= std::ios::app;
        if (mode.find('t') != std::string::npos) omode &= ~std::ios::binary;

        auto fs = std::make_shared<std::fstream>(name, omode);
        return File{fs, name};
    }
};

struct IO {
    static Error Copy(Writer &dst, Reader &src, int64_t &written) { return src.Copy(dst, written); }

    static Error CopyN(Writer &dst, Reader &src, int64_t n, int64_t &written) {
        written = 0;
        const size_t bufSize = 32 * 1024;
        uint8_t      buf[bufSize];

        while (written < n) {
            size_t toRead = static_cast<size_t>(std::min<int64_t>(n - written, bufSize));
            size_t nr = 0;
            auto   err = src.Read(buf, toRead, &nr);
            if (nr > 0) {
                size_t nw = 0;
                err = dst.Write(buf, nr, &nw);
                written += nw;

                if (err.code) return err;
                if (nw != nr) return ErrShortWrite;
            }

            if (err.code) {
                if (err == EOF && written == n) return Error();
                return err;
            }

            if (nr == 0) return ErrUnexpectedEOF;
        }
        return Error();
    }

    static Error ReadAll(Reader &r, std::vector<uint8_t> &data) { return r.ReadAll(data); }

    static Error ReadFull(Reader &r, uint8_t *buf, size_t len, size_t *n) { return r.ReadFull(buf, len, n); }

    static Error WriteString(Writer &w, const std::string &s, size_t *n) { return w.WriteString(s, n); }
};

void testIOFunctions() {
    std::string data = "Hello, world!";
    ByteReader  reader{.data = reinterpret_cast<const uint8_t *>(data.data()), .len = data.size(), .pos = 0};

    uint8_t buf[20];
    size_t  n = 0;
    auto    err = reader.Read(buf, 5, &n);
    std::cout << "Read " << n << " bytes: " << std::string(buf, buf + n) << "\n";

    auto   file = File::Open("test.txt", "w");
    size_t written = 0;
    err = file.Write(reinterpret_cast<const uint8_t *>(data.data()), data.size(), &written);
    file.Close();

    auto                 inFile = File::Open("test.txt", "r");
    std::vector<uint8_t> content;
    err = inFile.ReadAll(content);
    std::cout << "File content: " << std::string(content.begin(), content.end()) << "\n";
    inFile.Close();

    ByteWriter writer{.data = buf, .len = sizeof(buf), .pos = 0};

    int64_t totalWritten = 0;
    reader.pos = 0;// 重置读取位置
    err = IO::Copy(writer, reader, totalWritten);
    std::cout << "Copied " << totalWritten << " bytes: " << std::string(buf, buf + totalWritten) << "\n";

    ReadWriteCloser rwc{.reader = reader, .writer = writer, .closer = Closer{[]() { return Error(); }}};

    size_t nr = 0, nw = 0;
    rwc.Read(buf, 5, &nr);
    rwc.Write(buf, nr, &nw);
    rwc.Close();
}
