#pragma once

#include <string>
#include <vector>
#include <stdexcept>
#include <algorithm>
#include <cctype>
#include <cstdlib>
#include <ctime>
#include <fcntl.h>
#include <unistd.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <errno.h>
#include <dirent.h>
#include <climits>
#include <cstring>
#include <iostream>
#include <fstream>

struct OSError : std::runtime_error {
    explicit OSError(const std::string &msg, int code = 0) : std::runtime_error(msg + " (errno: " + std::to_string(code) + ")") {}
};

constexpr char PathSeparator = '/';
constexpr char PathListSeparator = ':';

constexpr int O_RDONLY = 00;
constexpr int O_WRONLY = 01;
constexpr int O_RDWR = 02;
constexpr int O_CREAT = 0100;
constexpr int O_TRUNC = 01000;
constexpr int O_APPEND = 02000;
constexpr int O_DIRECT = 040000;
constexpr int O_EXCL = 0200;

constexpr int PermRead = 0400;   // r--, owner read
constexpr int PermWrite = 0200;  // -w-, owner write
constexpr int PermExecute = 0100;// --x, owner execute
constexpr int PermDir = 0777;    // drwxrwxrwx

struct FileInfo {
    std::string name;   // 文件名
    size_t      size;   // 文件大小（字节）
    time_t      modTime;// 最后修改时间
    bool        isDir;  // 是否为目录
    mode_t      mode;   // 文件模式
};

struct File {
    int fd = -1;

    explicit File(int fileDescriptor) : fd(fileDescriptor) {}

    size_t Read(void *buffer, size_t length) {
        if (fd == -1) { throw OSError("File not open", EBADF); }

        ssize_t n = read(fd, buffer, length);
        if (n < 0) { throw OSError("Read error", errno); }
        return static_cast<size_t>(n);
    }

    size_t Write(const void *buffer, size_t length) {
        if (fd == -1) { throw OSError("File not open", EBADF); }

        ssize_t n = write(fd, buffer, length);
        if (n < 0) { throw OSError("Write error", errno); }
        return static_cast<size_t>(n);
    }

    void Close() {
        if (fd != -1) {
            close(fd);
            fd = -1;
        }
    }

    off_t Seek(off_t offset, int whence) {
        if (fd == -1) { throw OSError("File not open", EBADF); }

        off_t newPos = lseek(fd, offset, whence);
        if (newPos == -1) { throw OSError("Seek error", errno); }
        return newPos;
    }

    void Sync() {
        if (fd == -1) { throw OSError("File not open", EBADF); }
        fsync(fd);
    }
};

struct OS {
    static File OpenFile(const std::string &name, int flags, int perm = PermRead | PermWrite) {
        int fd = open(name.c_str(), flags, perm);
        if (fd == -1) { throw OSError("Cannot open file: " + name, errno); }
        return File(fd);
    }

    static File Open(const std::string &name) { return OpenFile(name, O_RDONLY); }

    static File Create(const std::string &name) { return OpenFile(name, O_CREAT | O_RDWR | O_TRUNC); }

    static void Remove(const std::string &name) {
        if (unlink(name.c_str()) {
            throw OSError("Cannot remove file: " + name, errno);
        }
    }

    static void Mkdir(const std::string &name, int perm = PermDir) {
        if (mkdir(name.c_str(), perm)) { throw OSError("Cannot create directory: " + name, errno); }
    }

    static void RemoveAll(const std::string &path) {
        struct stat st;
        if (lstat(path.c_str(), &st)) { throw OSError("Cannot stat path: " + path, errno); }

        if (S_ISDIR(st.st_mode)) {
            DIR *dir = opendir(path.c_str());
            if (!dir) { throw OSError("Cannot open directory: " + path, errno); }

            dirent *entry;
            while ((entry = readdir(dir))) {
                std::string name = entry->d_name;
                if (name == "." || name == "..") continue;

                std::string fullPath = path + PathSeparator + name;
                RemoveAll(fullPath);
            }

            closedir(dir);
            rmdir(path.c_str());
        } else {
            remove(path.c_str());
        }
    }

    static std::vector<FileInfo> ReadDir(const std::string &name) {
        std::vector<FileInfo> entries;
        DIR                  *dir = opendir(name.c_str());
        if (!dir) { throw OSError("Cannot open directory: " + name, errno); }

        dirent *entry;
        while ((entry = readdir(dir))) {
            std::string entryName = entry->d_name;
            if (entryName == "." || entryName == "..") continue;

            FileInfo info;
            info.name = entryName;
            std::string fullPath = name + PathSeparator + entryName;

            struct stat st;
            if (lstat(fullPath.c_str(), &st) == 0) {
                info.size = st.st_size;
                info.modTime = st.st_mtime;
                info.mode = st.st_mode;
                info.isDir = S_ISDIR(st.st_mode);
            }

            entries.push_back(info);
        }

        closedir(dir);
        return entries;
    }

    static FileInfo Stat(const std::string &name) {
        struct stat st;
        if (stat(name.c_str(), &st)) { throw OSError("Cannot get file info: " + name, errno); }

        FileInfo info;
        info.size = st.st_size;
        info.modTime = st.st_mtime;
        info.isDir = S_ISDIR(st.st_mode);
        info.mode = st.st_mode;

        size_t pos = name.find_last_of(PathSeparator);
        info.name = (pos != std::string::npos) ? name.substr(pos + 1) : name;

        return info;
    }

    static bool IsExist(const std::string &path) { return access(path.c_str(), F_OK) == 0; }

    static bool IsDir(const std::string &path) {
        struct stat st;
        if (lstat(path.c_str(), &st)) return false;
        return S_ISDIR(st.st_mode);
    }

    static bool IsRegular(const std::string &path) {
        struct stat st;
        if (lstat(path.c_str(), &st)) return false;
        return S_ISREG(st.st_mode);
    }

    static std::string Getenv(const std::string &key) {
        const char *value = std::getenv(key.c_str());
        return value ? value : "";
    }

    static void Setenv(const std::string &key, const std::string &value) {
        if (setenv(key.c_str(), value.c_str(), 1)) { throw OSError("Setenv failed", errno); }
    }

    static void Unsetenv(const std::string &key) {
        if (unsetenv(key.c_str())) { throw OSError("Unsetenv failed", errno); }
    }

    static std::vector<std::pair<std::string, std::string>> Environ() {
        std::vector<std::pair<std::string, std::string>> env;

        extern char **environ;
        for (char **envp = environ; *envp; envp++) {
            std::string entry(*envp);
            size_t      pos = entry.find('=');
            if (pos != std::string::npos) {
                std::string key = entry.substr(0, pos);
                std::string value = entry.substr(pos + 1);
                env.emplace_back(key, value);
            }
        }

        return env;
    }

    static void Exit(int code) { std::exit(code); }

    static int Getpid() { return getpid(); }

    static int Getppid() { return getppid(); }

    static std::string Getwd() {
        char cwd[PATH_MAX];
        if (!getcwd(cwd, sizeof(cwd))) { throw OSError("Getwd failed", errno); }
        return cwd;
    }

    static void Chdir(const std::string &path) {
        if (chdir(path.c_str())) { throw OSError("Chdir failed: " + path, errno); }
    }

    static std::string Getuid() { return std::to_string(getuid()); }

    static std::string Getgid() { return std::to_string(getgid()); }

    static std::string Geteuid() { return std::to_string(geteuid()); }

    static std::string Getegid() { return std::to_string(getegid()); }

    static time_t Now() { return time(nullptr); }

    static void StdoutToFile(const std::string &path) {
        std::ofstream file(path, std::ios::out | std::ios::app);
        if (!file) { throw OSError("Cannot open file for stdout: " + path, errno); }

        std::cout.rdbuf(file.rdbuf());
    }

    static void StderrToFile(const std::string &path) {
        std::ofstream file(path, std::ios::out | std::ios::app);
        if (!file) { throw OSError("Cannot open file for stderr: " + path, errno); }

        std::cerr.rdbuf(file.rdbuf());
    }

    static void WriteFile(const std::string &path, const std::string &data) {
        File file = OS::OpenFile(path, O_WRONLY | O_CREAT | O_TRUNC);
        file.Write(data.data(), data.size());
        file.Close();
    }

    static std::string ReadFile(const std::string &path) {
        File file = OS::OpenFile(path, O_RDONLY);

        struct stat st;
        if (fstat(file.fd, &st)) {
            file.Close();
            throw OSError("Stat failed: " + path, errno);
        }

        std::string content;
        content.resize(st.st_size);

        size_t totalRead = 0;
        while (totalRead < static_cast<size_t>(st.st_size)) {
            size_t n = file.Read(&content[totalRead], st.st_size - totalRead);
            if (n == 0) break;
            totalRead += n;
        }

        file.Close();
        return content;
    }
};

#include "os.h"
#include <iostream>

int main() {
    try {
        OS::WriteFile("test.txt", "Hello, OS Module!\n");
        std::cout << "File content: " << OS::ReadFile("test.txt") << "\n";

        OS::Mkdir("test_dir");

        OS::WriteFile("test_dir/test2.txt", "Another file\n");

        std::cout << "\nDirectory contents:\n";
        auto entries = OS::ReadDir("test_dir");
        for (const auto &e: entries) { std::cout << " - " << e.name << " (" << e.size << " bytes)\n"; }

        FileInfo info = OS::Stat("test.txt");
        std::cout << "\nFile info: " << info.name << ", size: " << info.size << ", isDir: " << std::boolalpha << info.isDir << "\n";

        std::cout << "\nHome directory: " << OS::Getenv("HOME") << "\n";

        std::cout << "PID: " << OS::Getpid() << ", PPID: " << OS::Getppid() << "\n";

        std::string cwd = OS::Getwd();
        std::cout << "Current dir: " << cwd << "\n";

        OS::Chdir("test_dir");
        std::cout << "Now in: " << OS::Getwd() << "\n";
        OS::Chdir(cwd);// 切换回来

        std::cout << "UID: " << OS::Getuid() << ", GID: " << OS::Getgid() << "\n";

        OS::RemoveAll("test_dir");
        OS::Remove("test.txt");

    } catch (const OSError &e) {
        std::cerr << "OS Error: " << e.what() << "\n";
        return 1;
    } catch (const std::exception &e) {
        std::cerr << "Error: " << e.what() << "\n";
        return 1;
    }

    return 0;
}
