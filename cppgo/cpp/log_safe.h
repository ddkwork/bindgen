#include <string>
#include <vector>
#include <map>
#include <functional>
#include <fstream>
#include <sstream>
#include <iomanip>
#include <ctime>
#include <filesystem>
#include <cctype>
#include <algorithm>
#include <mutex>
#include <atomic>

using namespace std;
using namespace filesystem;

enum Kind { INFO, TRACE, WARNING, SUCCESS };

enum IntegerType { UNSIGNED, SIGNED };

struct KeyValue {
    string key;
    string value;

    string Value() const {
        if (!value.empty() && value.back() == '\n') { return value.substr(0, value.size() - 1); }
        return value;
    }

    string ToString() const { return key + Value(); }
};

struct Log {
    atomic<Kind> kind = INFO;
    atomic_bool  debug = true;

    recursive_mutex logMutex;

    KeyValue                       row;
    function<void(const string &)> callback;

    string TextIndent(const string &src, bool isLeftAlign = true) const {
        const int    hexDumpIndentLen = 26;
        const string separate = " │ ";

        lock_guard<recursive_mutex> lock(logMutex);

        int spaceLen = hexDumpIndentLen;
        for (char c: src) { spaceLen -= (isprint(c) ? 1 : 0); }

        string spaceStr = spaceLen > 0 ? string(spaceLen, ' ') : "";

        if (isLeftAlign) { return src + spaceStr + separate; }
        return spaceStr + src + separate;
    }

    void SetCallBack(function<void(const string &)> cb) {
        lock_guard<recursive_mutex> lock(logMutex);
        callback = cb;
    }

    string LogPath() { return DataDir() + "/log.log"; }

    string DataDir() { return "."; }

    static Log New() { return Log(); }

    void Init() {
        lock_guard<recursive_mutex> lock(logMutex);

        if (1) {// 简化逻辑，实际应检查环境
            string path = LogPath();
            if (FileLineCountIsMoreThan(path, 2)) { truncateLogFile(path); }
            Trace("--------- key ---------", "------------------ value ------------------");

            if (IsWindows()) { system("go env -w GOFLAGS=-buildmode=exe"); }
            system("go env -w GOPROXY=https://goproxy.cn");
        }
    }

    void ChdirToGithubWorkspace() {
        lock_guard<recursive_mutex> lock(logMutex);

        if (IsGithubAction()) { current_path(GithubWorkspace()); }
        string dir = current_path().string();
        Info("GITHUB_WORKSPACE", dir);
    }

    template<typename T>
    void HexDump(const string &title, T data) {
        HexDump(title, DumpHex(data));
    }

    void Todo(const string &body) { Warning("TODO", body); }

    template<typename T>
    string Hex(const string &title, T v) {
        return Hex(title, FormatInteger(v, UNSIGNED));
    }

    template<typename... Args>
    void Info(const string &title, Args... args) {
        kind.store(INFO, memory_order_relaxed);
        logTitle(title, args...);
    }

    template<typename... Args>
    void Trace(const string &title, Args... args) {
        kind.store(TRACE, memory_order_relaxed);
        logTitle(title, args...);
    }

    template<typename... Args>
    void Warning(const string &title, Args... args) {
        kind.store(WARNING, memory_order_relaxed);
        logTitle(title, args...);
    }

    template<typename... Args>
    void Json(const string &title, Args... args) {
        kind.store(INFO, memory_order_relaxed);
        logTitle(title, formatJson(args)...);
    }

    template<typename... Args>
    void Success(const string &title, Args... args) {
        kind.store(SUCCESS, memory_order_relaxed);
        logTitle(title, args...);
    }

    template<typename T>
    void Struct(const T &object) {
        Info("Struct", typeid(T).name());
    }

    void SetDebug(bool dbg) { debug.store(dbg, memory_order_relaxed); }

    void Request(const string &request, bool body) {
        lock_guard<recursive_mutex> lock(logMutex);
        Info("HTTP Request", request);
    }

    void Response(const string &response, bool body) {
        lock_guard<recursive_mutex> lock(logMutex);
        Info("HTTP Response", response);
    }

    string DumpRequest(const string &request, bool body) {
        lock_guard<recursive_mutex> lock(logMutex);
        return "Dumped request: " + request;
    }

    string DumpResponse(const string &response, bool body) {
        lock_guard<recursive_mutex> lock(logMutex);
        return "Dumped response: " + response;
    }

private:
    bool FileLineCountIsMoreThan(const string &path, int count) {
        lock_guard<recursive_mutex> lock(logMutex);

        if (!exists(path)) return false;

        ifstream file(path);
        int      lines = 0;
        string   line;
        while (getline(file, line)) {
            if (++lines > count) return true;
        }
        return false;
    }

    void truncateLogFile(const string &path) {
        lock_guard<recursive_mutex> lock(logMutex);

        ofstream file(path, ios::trunc);
        file.close();
    }

    bool IsWindows() {
#ifdef _WIN32
        return true;
#else
        return false;
#endif
    }

    string GithubWorkspace() {
        char *env = getenv("GITHUB_WORKSPACE");
        return env ? env : ".";
    }

    bool IsGithubAction() { return getenv("GITHUB_ACTION") != nullptr; }

    string DumpHex(const string &data) {
        lock_guard<recursive_mutex> lock(logMutex);

        ostringstream oss;
        for (char c: data) { oss << hex << setw(2) << setfill('0') << (int) (unsigned char) c << " "; }
        return oss.str();
    }

    template<typename T>
    string FormatInteger(T v, IntegerType type) {
        lock_guard<recursive_mutex> lock(logMutex);

        ostringstream oss;
        if (type == UNSIGNED) {
            oss << "0x" << hex << uppercase << v;
        } else {
            oss << dec << v;
        }
        return oss.str();
    }

    template<typename... Args>
    void logTitle(const string &title, Args... args) {
        lock_guard<recursive_mutex> lock(logMutex);

        ostringstream oss;
        appendArgs(oss, args...);

        string value = oss.str();
        string timestamp = getTimestamp();

        string logLine = "[" + timestamp + "] ";
        switch (kind.load(memory_order_relaxed)) {
            case INFO:
                logLine += "INFO: ";
                break;
            case TRACE:
                logLine += "TRACE: ";
                break;
            case WARNING:
                logLine += "WARNING: ";
                break;
            case SUCCESS:
                logLine += "SUCCESS: ";
                break;
        }
        logLine += TextIndent(title) + value;

        if (debug.load(memory_order_relaxed)) { cout << logLine << endl; }

        ofstream logFile(LogPath(), ios::app);
        logFile << logLine << endl;
        logFile.close();

        if (callback) { callback(logLine); }
    }

    void appendArgs(ostringstream &oss) {}

    template<typename T, typename... Args>
    void appendArgs(ostringstream &oss, T first, Args... args) {
        oss << first << " ";
        appendArgs(oss, args...);
    }

    string getTimestamp() {
        lock_guard<recursive_mutex> lock(logMutex);

        time_t now = time(nullptr);
        tm     timeinfo;
        localtime_r(&now, &timeinfo);
        char buffer[80];
        strftime(buffer, sizeof(buffer), "%Y-%m-%d %H:%M:%S", &timeinfo);
        return string(buffer);
    }

    template<typename... Args>
    vector<string> formatJson(Args... args) {
        lock_guard<recursive_mutex> lock(logMutex);

        vector<string> result = {"{"};
        addJsonItems(result, args...);
        result.push_back("}");
        return result;
    }

    void addJsonItems(vector<string> &result) {}

    template<typename T, typename... Args>
    void addJsonItems(vector<string> &result, T item, Args... args) {
        result.push_back("\"" + string(item) + "\", ");
        addJsonItems(result, args...);
    }
};

Log l = Log::New();

template<typename T>
void HexDump(const string &title, T data) {
    l.HexDump(title, data);
}

void Todo(const string &body) { l.Todo(body); }

template<typename T>
string Hex(const string &title, T v) {
    return l.Hex(title, v);
}

template<typename... Args>
void Info(const string &title, Args... args) {
    l.Info(title, args...);
}

template<typename... Args>
void Trace(const string &title, Args... args) {
    l.Trace(title, args...);
}

template<typename... Args>
void Warning(const string &title, Args... args) {
    l.Warning(title, args...);
}

template<typename T>
void Json(const string &title, T data) {
    l.Json(title, data);
}

template<typename... Args>
void Success(const string &title, Args... args) {
    l.Success(title, args...);
}

template<typename T>
void Struct(T object) {
    l.Struct(object);
}

void SetDebug(bool debug) { l.SetDebug(debug); }

void Request(const string &request, bool body) { l.Request(request, body); }

void Response(const string &response, bool body) { l.Response(response, body); }

string DumpRequest(const string &request, bool body) { return l.DumpRequest(request, body); }

string DumpResponse(const string &response, bool body) { return l.DumpResponse(response, body); }

void concurrentLogTest() {
    vector<thread> threads;

    for (int i = 0; i < 10; ++i) {
        threads.emplace_back([i]() {
            for (int j = 0; j < 100; ++j) { Info("Thread " + to_string(i), "Message " + to_string(j)); }
        });
    }

    for (auto &t: threads) { t.join(); }
}

int main() {
    l.Init();

    concurrentLogTest();

    Info("System Started", "Version", "1.0.0");
    Warning("Low Disk Space", "Free: 100MB");
    Success("Operation Complete", "Time", "2.5s");
    Todo("Implement feature XYZ");

    HexDump("Memory", "HelloWorld");

    struct Point {
        int x, y;
    };

    Point p{10, 20};
    Struct(p);

    Request("GET /api/user", true);

    SetDebug(true);

    return 0;
}
