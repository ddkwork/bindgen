#pragma once
#include <string>
#include <vector>
#include <map>
#include <functional>
#include <cstdlib>
#include <cstdio>
#include <cstring>
#include <cerrno>
#include <stdexcept>
#include <sstream>
#include <iostream>
#include <memory>
#include <algorithm>
#include <thread>
#include <mutex>
#include <condition_variable>
#include <chrono>
#include <fcntl.h>

#ifdef _WIN32
#include <windows.h>
#include <io.h>
#include <process.h>
#define pipe(fds) _pipe(fds, 4096, _O_BINARY)
#define close     _close
#define read      _read
#define write     _write
#define dup       _dup
#define dup2      _dup2
#define fileno    _fileno
#else
#include <unistd.h>
#include <sys/wait.h>
#include <signal.h>
#endif

struct Error {
    std::string msg;
    int         code;

    Error() : code(0) {}

    Error(const std::string &m) : msg(m), code(-1) {}

    Error(const std::string &m, int c) : msg(m), code(c) {}

    std::string Error() const { return code ? (msg + " (code: " + std::to_string(code) + ")") : msg; }

    bool operator==(const Error &other) const { return code == other.code && msg == other.msg; }

    operator bool() const { return code != 0; }
};

Error ErrNotFound = Error("executable file not found");
Error ErrWait = Error("wait error");
Error ErrKilled = Error("process killed");
Error ErrTimeout = Error("process timeout");

struct ProcessState {
    int         pid;
    int         status;
    bool        exited;
    bool        success;
    int         exitCode;
    std::string error;

    bool Exited() const { return exited; }

    bool Success() const { return success; }

    int ExitCode() const { return exitCode; }

    std::string String() const {
        if (!exited) return "process not exited";
        return success ? "exit status 0" : ("exit status " + std::to_string(exitCode));
    }
};

struct Cmd {
    std::string              path;
    std::vector<std::string> args;
    std::vector<std::string> env;
    std::string              dir;

    FILE *stdin = nullptr;
    FILE *stdout = nullptr;
    FILE *stderr = nullptr;

    ProcessState state;
    bool         started = false;
    bool         finished = false;

#ifdef _WIN32
    HANDLE process = INVALID_HANDLE_VALUE;
    HANDLE thread = INVALID_HANDLE_VALUE;
#else
    int pid = -1;
#endif

    Error Run() {
        auto err = Start();
        if (err) return err;

        return Wait();
    }

    Error Start() {
        if (started) return Error("process already started");

        std::vector<char *> argv;
        argv.push_back(const_cast<char *>(path.c_str()));
        for (auto &arg: args) { argv.push_back(const_cast<char *>(arg.c_str())); }
        argv.push_back(nullptr);

        std::vector<char *> envp;
        for (auto &e: env) { envp.push_back(const_cast<char *>(e.c_str())); }
        envp.push_back(nullptr);

        int stdinPipe[2] = {-1, -1};
        int stdoutPipe[2] = {-1, -1};
        int stderrPipe[2] = {-1, -1};

        if (!stdin && pipe(stdinPipe) != 0) return Error("stdin pipe failed", errno);
        if (!stdout && pipe(stdoutPipe) != 0) return Error("stdout pipe failed", errno);
        if (!stderr && pipe(stderrPipe) != 0) return Error("stderr pipe failed", errno);

#ifdef _WIN32
        STARTUPINFO         si;
        PROCESS_INFORMATION pi;
        ZeroMemory(&si, sizeof(si));
        ZeroMemory(&pi, sizeof(pi));
        si.cb = sizeof(si);

        if (!stdin) {
            si.hStdInput = (HANDLE) _get_osfhandle(stdinPipe[0]);
            close(stdinPipe[1]);// 关闭写端
        }
        if (!stdout) {
            si.hStdOutput = (HANDLE) _get_osfhandle(stdoutPipe[1]);
            close(stdoutPipe[0]);// 关闭读端
        }
        if (!stderr) {
            si.hStdError = (HANDLE) _get_osfhandle(stderrPipe[1]);
            close(stderrPipe[0]);// 关闭读端
        }

        si.dwFlags |= STARTF_USESTDHANDLES;

        std::string cmdLine = path;
        for (auto &arg: args) { cmdLine += " " + arg; }

        BOOL success = CreateProcess(NULL, const_cast<char *>(cmdLine.c_str()), NULL, NULL, TRUE, 0, envp.empty() ? NULL : envp.data(), dir.empty() ? NULL : dir.c_str(), &si, &pi);

        if (!success) {
            DWORD err = GetLastError();
            return Error("CreateProcess failed", err);
        }

        process = pi.hProcess;
        thread = pi.hThread;
        pid = pi.dwProcessId;

        if (!stdin) {
            stdin = _fdopen(stdinPipe[1], "wb");
            if (!stdin) return Error("fdopen stdin failed", errno);
        }
        if (!stdout) {
            stdout = _fdopen(stdoutPipe[0], "rb");
            if (!stdout) return Error("fdopen stdout failed", errno);
        }
        if (!stderr) {
            stderr = _fdopen(stderrPipe[0], "rb");
            if (!stderr) return Error("fdopen stderr failed", errno);
        }
#else
        pid = fork();
        if (pid < 0) { return Error("fork failed", errno); }

        if (pid == 0) {// 子进程
            if (!stdin) {
                dup2(stdinPipe[0], STDIN_FILENO);
                close(stdinPipe[1]);
            }
            if (!stdout) {
                dup2(stdoutPipe[1], STDOUT_FILENO);
                close(stdoutPipe[0]);
            }
            if (!stderr) {
                dup2(stderrPipe[1], STDERR_FILENO);
                close(stderrPipe[0]);
            }

            if (!dir.empty() && chdir(dir.c_str()) != 0) { _exit(127); }

            if (envp.empty()) {
                execvp(path.c_str(), argv.data());
            } else {
                execvpe(path.c_str(), argv.data(), envp.data());
            }

            _exit(127);
        }

        if (!stdin) {
            close(stdinPipe[0]);
            stdin = fdopen(stdinPipe[1], "w");
        }
        if (!stdout) {
            close(stdoutPipe[1]);
            stdout = fdopen(stdoutPipe[0], "r");
        }
        if (!stderr) {
            close(stderrPipe[1]);
            stderr = fdopen(stderrPipe[0], "r");
        }
#endif

        started = true;
        return Error();
    }

    Error Wait() {
        if (!started) return Error("process not started");
        if (finished) return Error("process already finished");

#ifdef _WIN32
        DWORD result = WaitForSingleObject(process, INFINITE);
        if (result == WAIT_FAILED) { return Error("WaitForSingleObject failed", GetLastError()); }

        DWORD exitCode;
        if (!GetExitCodeProcess(process, &exitCode)) { return Error("GetExitCodeProcess failed", GetLastError()); }

        CloseHandle(process);
        CloseHandle(thread);

        state.exited = true;
        state.exitCode = static_cast<int>(exitCode);
        state.success = (exitCode == 0);
#else
        int status;
        if (waitpid(pid, &status, 0) < 0) { return Error("waitpid failed", errno); }

        state.exited = true;
        if (WIFEXITED(status)) {
            state.exitCode = WEXITSTATUS(status);
            state.success = (state.exitCode == 0);
        } else if (WIFSIGNALED(status)) {
            state.exitCode = WTERMSIG(status);
            state.success = false;
            state.error = "terminated by signal";
        }
#endif

        if (stdin && stdin != stdin) fclose(stdin);
        if (stdout && stdout != stdout) fclose(stdout);
        if (stderr && stderr != stderr) fclose(stderr);

        stdin = stdout = stderr = nullptr;
        finished = true;
        return Error();
    }

    Error Kill() {
        if (!started || finished) return Error("process not running");

#ifdef _WIN32
        if (!TerminateProcess(process, 1)) { return Error("TerminateProcess failed", GetLastError()); }
#else
        if (kill(pid, SIGKILL) != 0) { return Error("kill failed", errno); }
#endif

        state.exited = true;
        state.success = false;
        state.exitCode = -1;
        state.error = "killed";
        finished = true;
        return Error();
    }

    Error CombinedOutput(std::string &output) {
        if (stdout || stderr) { return Error("stdout or stderr already set"); }

        int outPipe[2];
        if (pipe(outPipe) != 0) { return Error("pipe failed", errno); }

        stdout = stderr =
#ifdef _WIN32
                _fdopen(outPipe[1], "wb");
#else
                fdopen(outPipe[1], "w");
#endif

        auto err = Run();
        if (err) return err;

        char    buffer[4096];
        ssize_t n;
        while ((n = read(outPipe[0], buffer, sizeof(buffer))) > 0) { output.append(buffer, n); }

        close(outPipe[0]);
        return Error();
    }

    Error Output(std::string &output) {
        if (stdout) { return Error("stdout already set"); }

        int outPipe[2];
        if (pipe(outPipe) != 0) { return Error("pipe failed", errno); }

        stdout =
#ifdef _WIN32
                _fdopen(outPipe[1], "wb");
#else
                fdopen(outPipe[1], "w");
#endif

        auto err = Run();
        if (err) return err;

        char    buffer[4096];
        ssize_t n;
        while ((n = read(outPipe[0], buffer, sizeof(buffer))) > 0) { output.append(buffer, n); }

        close(outPipe[0]);
        return Error();
    }

    Error WriteStdin(const std::string &data) {
        if (!stdin) return Error("stdin not set");

        size_t written = fwrite(data.data(), 1, data.size(), stdin);
        if (written < data.size()) { return Error("write stdin failed", errno); }
        fflush(stdin);
        return Error();
    }

    Error CloseStdin() {
        if (stdin) {
            fclose(stdin);
            stdin = nullptr;
        }
        return Error();
    }
};

struct Command {
    static Cmd Create(const std::string &name, const std::vector<std::string> &args) { return Cmd{.path = name, .args = args}; }

    static std::string LookPath(const std::string &file) {
        const char *pathEnv = std::getenv("PATH");
        if (!pathEnv) return "";

        std::string              path = pathEnv;
        std::vector<std::string> paths;
        size_t                   start = 0, end = 0;

        while ((end = path.find(':', start)) != std::string::npos) {
            paths.push_back(path.substr(start, end - start));
            start = end + 1;
        }
        paths.push_back(path.substr(start));

        for (const auto &p: paths) {
            std::string fullPath = p + "/" + file;
#ifdef _WIN32
            fullPath += ".exe";
#endif
            if (access(fullPath.c_str(), X_OK) == 0) { return fullPath; }
        }

        return "";
    }

    static Error Run(const std::string &name, const std::vector<std::string> &args, std::string &output) {
        Cmd cmd = Create(name, args);
        return cmd.CombinedOutput(output);
    }

    static Error Run(const std::string &name, const std::vector<std::string> &args) {
        Cmd cmd = Create(name, args);
        return cmd.Run();
    }
};

void testCmdFunctions() {
    std::string output;
    auto        err = Command::Run("echo", {"Hello, world!"}, output);
    std::cout << "Echo output: " << output << "\n";

    Cmd cmd1 = Command::Create("ls", {"-l"});
    Cmd cmd2 = Command::Create("grep", {"txt"});

    int pipefd[2];
    if (pipe(pipefd) == 0) {
        cmd1.stdout = fdopen(pipefd[1], "w");
        cmd2.stdin = fdopen(pipefd[0], "r");
    }

    err = cmd1.Start();
    if (err) std::cerr << "cmd1 start failed: " << err.Error() << "\n";

    err = cmd2.Start();
    if (err) std::cerr << "cmd2 start failed: " << err.Error() << "\n";

    cmd1.Wait();
    cmd2.Wait();

    char        buffer[1024];
    std::string result;
    while (fgets(buffer, sizeof(buffer), cmd2.stdout)) { result += buffer; }
    std::cout << "Files with 'txt':\n" << result;

    Cmd sleepCmd = Command::Create("sleep", {"5"});
    err = sleepCmd.Start();
    if (err) std::cerr << "sleep start failed: " << err.Error() << "\n";

    auto start = std::chrono::steady_clock::now();
    bool timedOut = false;

    while (!sleepCmd.finished) {
        auto now = std::chrono::steady_clock::now();
        auto elapsed = std::chrono::duration_cast<std::chrono::seconds>(now - start).count();

        if (elapsed > 3) {
            sleepCmd.Kill();
            timedOut = true;
            break;
        }

        std::this_thread::sleep_for(std::chrono::milliseconds(100));
    }

    if (timedOut) {
        std::cout << "Sleep command timed out and killed\n";
    } else {
        std::cout << "Sleep command completed\n";
    }

    Cmd envCmd = Command::Create("env", {});
    envCmd.env = {"MY_VAR=hello", "ANOTHER_VAR=world"};

    std::string envOutput;
    envCmd.CombinedOutput(envOutput);
    std::cout << "Environment output:\n" << envOutput;

    std::string goPath = Command::LookPath("go");
    if (!goPath.empty()) {
        std::cout << "Found Go at: " << goPath << "\n";
    } else {
        std::cout << "Go not found in PATH\n";
    }
}
