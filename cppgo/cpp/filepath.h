#pragma once

#include <string>
#include <vector>
#include <stdexcept>
#include <algorithm>
#include <cctype>
#include <sys/stat.h>
#include <unistd.h>
#include <dirent.h>
#include <climits>
#include <cstring>
#include <iostream>

struct Filepath {
    static const char Separator = '/';
    static const char ListSeparator = ':';

    static std::string Clean(const std::string &path) {
        if (path.empty()) return ".";

        std::vector<std::string> parts;
        std::string              cleanPath;

        size_t start = 0;
        bool   absolute = !path.empty() && path[0] == Separator;

        if (absolute) {
            cleanPath += Separator;
            start = 1;
        }

        size_t pos = start;
        while (pos < path.size()) {
            size_t end = path.find(Separator, pos);
            if (end == std::string::npos) end = path.size();

            std::string part = path.substr(pos, end - pos);
            pos = end + 1;

            if (part.empty() || part == ".") {
            } else if (part == "..") {
                if (!parts.empty() && parts.back() != "..") {
                    parts.pop_back();
                } else if (!absolute) {
                    parts.push_back("..");
                }
            } else {
                parts.push_back(part);
            }
        }

        for (size_t i = 0; i < parts.size(); i++) {
            if (i > 0 || absolute) cleanPath += Separator;
            cleanPath += parts[i];
        }

        if (cleanPath.empty()) return ".";
        return cleanPath;
    }

    static bool IsAbs(const std::string &path) { return !path.empty() && path[0] == Separator; }

    static std::string Join(const std::vector<std::string> &paths) {
        if (paths.empty()) return "";

        std::string result;
        bool        needsSeparator = false;

        for (const auto &part: paths) {
            if (part.empty()) continue;

            if (needsSeparator) result += Separator;
            result += part;

            needsSeparator = (result.back() != Separator && !part.empty());
        }

        return Clean(result);
    }

    static std::string Join(const std::string &a, const std::string &b) { return Join({a, b}); }

    static std::string Dir(const std::string &path) {
        if (path.empty()) return ".";

        size_t lastSep = path.find_last_of(Separator);
        if (lastSep == std::string::npos) {
            return ".";// 没有目录部分
        }

        if (lastSep == 0) {
            return "/";// 根目录
        }

        return Clean(path.substr(0, lastSep));
    }

    static std::string Base(const std::string &path) {
        if (path.empty()) return ".";

        size_t lastSep = path.find_last_of(Separator);
        if (lastSep == std::string::npos) {
            return path;// 没有目录部分，直接返回文件名
        }

        return path.substr(lastSep + 1);
    }

    static std::string Ext(const std::string &path) {
        std::string base = Base(path);
        size_t      dotPos = base.find_last_of('.');

        if (dotPos == std::string::npos || dotPos == 0) { return ""; }

        if (dotPos != std::string::npos && dotPos + 1 < base.size() && base[dotPos + 1] == Separator) { return ""; }

        return base.substr(dotPos);
    }

    static std::string Stem(const std::string &path) {
        std::string base = Base(path);
        size_t      dotPos = base.find_last_of('.');

        if (dotPos == std::string::npos || dotPos == 0) { return base; }

        return base.substr(0, dotPos);
    }

    static std::string Abs(const std::string &path) {
        if (path.empty()) {
            char cwd[PATH_MAX];
            if (getcwd(cwd, sizeof(cwd)) {
                return cwd;
            }
            return "/";
        }

        if (IsAbs(path)) { return Clean(path); }

        char cwd[PATH_MAX];
        if (!getcwd(cwd, sizeof(cwd))) { return Clean("/" + path); }

        return Clean(std::string(cwd) + "/" + path);
    }

    static time_t ModifiedTime(const std::string &path) {
        struct stat fileStat;
        if (stat(path.c_str(), &fileStat) != 0) { throw std::runtime_error("Cannot get file stats: " + path); }
        return fileStat.st_mtime;
    }

    static off_t Size(const std::string &path) {
        struct stat fileStat;
        if (stat(path.c_str(), &fileStat) != 0) { throw std::runtime_error("Cannot get file stats: " + path); }
        return fileStat.st_size;
    }

    static bool Exists(const std::string &path) {
        struct stat fileStat;
        return stat(path.c_str(), &fileStat) == 0;
    }

    static bool IsDir(const std::string &path) {
        struct stat fileStat;
        if (stat(path.c_str(), &fileStat) != 0) { return false; }
        return S_ISDIR(fileStat.st_mode);
    }

    static bool IsFile(const std::string &path) {
        struct stat fileStat;
        if (stat(path.c_str(), &fileStat) != 0) { return false; }
        return S_ISREG(fileStat.st_mode);
    }

    static std::vector<std::string> Glob(const std::string &pattern) {
        std::string dir = Dir(pattern);
        std::string base = Base(pattern);

        return findFiles(dir, base);
    }

private:
    static std::vector<std::string> findFiles(const std::string &dir, const std::string &pattern) {
        std::vector<std::string> result;

        if (!IsDir(dir)) { return result; }

        DIR *dp = opendir(dir.c_str());
        if (!dp) { return result; }

        dirent *entry;
        while ((entry = readdir(dp)) != nullptr) {
            std::string name = entry->d_name;

            if (name == "." || name == "..") continue;

            std::string fullPath = dir + Separator + name;

            if (IsDir(fullPath)) {
                std::string              subDir = dir + Separator + name;
                std::vector<std::string> subFiles = findFiles(subDir, pattern);
                result.insert(result.end(), subFiles.begin(), subFiles.end());
            }

            if (matchPattern(name, pattern)) { result.push_back(fullPath); }
        }

        closedir(dp);
        return result;
    }

    static bool matchPattern(const std::string &name, const std::string &pattern) {
        if (pattern.empty()) return true;
        if (pattern == "*") return true;

        std::string nameLower = toLower(name);
        std::string patternLower = toLower(pattern);

        size_t i = 0, j = 0;
        size_t starPos = std::string::npos;
        size_t starMatch = 0;

        while (i < nameLower.size()) {
            if (j < patternLower.size() && patternLower[j] == '*') {
                starPos = j;
                starMatch = i;
                j++;
            } else if (j < patternLower.size() && (patternLower[j] == '?' || patternLower[j] == nameLower[i])) {
                i++;
                j++;
            } else if (starPos != std::string::npos) {
                j = starPos + 1;
                starMatch++;
                i = starMatch;
            } else {
                return false;
            }
        }

        while (j < patternLower.size() && patternLower[j] == '*') { j++; }

        return j == patternLower.size();
    }

    static std::string toLower(const std::string &str) {
        std::string result = str;
        std::transform(result.begin(), result.end(), result.begin(), [](unsigned char c) { return std::tolower(c); });
        return result;
    }
};

#include "filepath.h"
#include <iostream>

int main() {
    std::cout << "Clean: " << Filepath::Clean("/usr//local/../bin") << std::endl;

    std::cout << "Is absolute: " << std::boolalpha << Filepath::IsAbs("/usr/bin") << std::endl;

    std::cout << "Join: " << Filepath::Join("/usr", "local", "bin") << std::endl;

    std::cout << "Dir: " << Filepath::Dir("/usr/local/bin/bash") << std::endl;
    std::cout << "Base: " << Filepath::Base("/usr/local/bin/bash") << std::endl;
    std::cout << "Ext: " << Filepath::Ext("/usr/local/bin/bash.tar.gz") << std::endl;
    std::cout << "Stem: " << Filepath::Stem("/usr/local/bin/bash.tar.gz") << std::endl;

    std::cout << "Abs: " << Filepath::Abs("../src") << std::endl;

    if (Filepath::Exists("example.txt")) {
        std::cout << "File size: " << Filepath::Size("example.txt") << " bytes" << std::endl;
        std::cout << "Modified: " << Filepath::ModifiedTime("example.txt") << std::endl;
    }

    std::cout << "Is dir: " << Filepath::IsDir("/usr") << std::endl;
    std::cout << "Is file: " << Filepath::IsFile("/usr/bin/ls") << std::endl;

    std::vector<std::string> files = Filepath::Glob("/tmp/backups/*.tar.gz");
    std::cout << "Found files:\n";
    for (const auto &f: files) { std::cout << " - " << f << std::endl; }

    return 0;
}
