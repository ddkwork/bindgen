#pragma once
#include <string>
#include <functional>
#include <iostream>
#include <sstream>
#include <vector>
#include <cmath>
#include <memory>
#include <cstring>
#include <algorithm>
#include <exception>
#include <map>
#include <regex>
#include <fstream>
#include <typeinfo>
#include <type_traits>

struct AssertionError {
    std::string Message;
    std::string File;
    int         Line;

    AssertionError(const std::string &msg, const std::string &f, int l) : Message(msg), File(f), Line(l) {}
};

struct AssertionContext {
    const char *File;
    int         Line;
    std::string CustomMsg;

    AssertionContext(const char *f, int l, const std::string &msg = "") : File(f), Line(l), CustomMsg(msg) {}

    void Fail(const std::string &message) {
        std::string fullMsg = CustomMsg.empty() ? message : CustomMsg + ": " + message;
        throw AssertionError(fullMsg, File, Line);
    }
};

struct Assert {
    static void Equal(AssertionContext &ctx, const auto &expected, const auto &actual) {
        if (expected != actual) {
            std::ostringstream ss;
            ss << "Not equal: \n"
               << "expected: " << expected << "\n"
               << "actual  : " << actual;
            ctx.Fail(ss.str());
        }
    }

    static void Equal(AssertionContext &ctx, float expected, float actual, float epsilon = 1e-6) {
        if (std::fabs(expected - actual) > epsilon) {
            std::ostringstream ss;
            ss << "Not equal: \n"
               << "expected: " << expected << "\n"
               << "actual  : " << actual << "\n"
               << "epsilon : " << epsilon;
            ctx.Fail(ss.str());
        }
    }

    static void Equal(AssertionContext &ctx, double expected, double actual, double epsilon = 1e-6) {
        if (std::fabs(expected - actual) > epsilon) {
            std::ostringstream ss;
            ss << "Not equal: \n"
               << "expected: " << expected << "\n"
               << "actual  : " << actual << "\n"
               << "epsilon : " << epsilon;
            ctx.Fail(ss.str());
        }
    }

    static void Equal(AssertionContext &ctx, const char *expected, const char *actual) {
        if (strcmp(expected, actual) != 0) {
            std::ostringstream ss;
            ss << "Not equal: \n"
               << "expected: \"" << expected << "\"\n"
               << "actual  : \"" << actual << "\"";
            ctx.Fail(ss.str());
        }
    }

    static void Equal(AssertionContext &ctx, const std::string &expected, const std::string &actual) {
        if (expected != actual) {
            std::ostringstream ss;
            ss << "Not equal: \n"
               << "expected: \"" << expected << "\"\n"
               << "actual  : \"" << actual << "\"";
            ctx.Fail(ss.str());
        }
    }

    template<typename T>
    static void Equal(AssertionContext &ctx, const std::vector<T> &expected, const std::vector<T> &actual) {
        if (expected.size() != actual.size()) {
            std::ostringstream ss;
            ss << "Size mismatch: \n"
               << "expected: " << expected.size() << " elements\n"
               << "actual  : " << actual.size() << " elements";
            ctx.Fail(ss.str());
        }

        for (size_t i = 0; i < expected.size(); i++) {
            if (expected[i] != actual[i]) {
                std::ostringstream ss;
                ss << "Element mismatch at index " << i << ": \n"
                   << "expected: " << expected[i] << "\n"
                   << "actual  : " << actual[i];
                ctx.Fail(ss.str());
            }
        }
    }

    static void NotEqual(AssertionContext &ctx, const auto &expected, const auto &actual) {
        if (expected == actual) {
            std::ostringstream ss;
            ss << "Values should not be equal: " << actual;
            ctx.Fail(ss.str());
        }
    }

    static void True(AssertionContext &ctx, bool condition) {
        if (!condition) { ctx.Fail("Expected true but got false"); }
    }

    static void False(AssertionContext &ctx, bool condition) {
        if (condition) { ctx.Fail("Expected false but got true"); }
    }

    static void Nil(AssertionContext &ctx, const std::shared_ptr<void> &ptr) {
        if (ptr != nullptr) { ctx.Fail("Expected nil but got non-nil"); }
    }

    static void NotNil(AssertionContext &ctx, const std::shared_ptr<void> &ptr) {
        if (ptr == nullptr) { ctx.Fail("Expected non-nil but got nil"); }
    }

    template<typename T>
    static void Nil(AssertionContext &ctx, T *ptr) {
        if (ptr != nullptr) { ctx.Fail("Expected nil but got non-nil"); }
    }

    template<typename T>
    static void NotNil(AssertionContext &ctx, T *ptr) {
        if (ptr == nullptr) { ctx.Fail("Expected non-nil but got nil"); }
    }

    static void Error(AssertionContext &ctx, const std::function<void()> &f) {
        bool hasError = false;
        try {
            f();
        } catch (...) { hasError = true; }

        if (!hasError) { ctx.Fail("Expected error but none was thrown"); }
    }

    static void NoError(AssertionContext &ctx, const std::function<void()> &f) {
        try {
            f();
        } catch (const std::exception &e) { ctx.Fail("Unexpected error: " + std::string(e.what())); } catch (...) {
            ctx.Fail("Unexpected unknown error");
        }
    }

    template<typename Container, typename Item>
    static void Contains(AssertionContext &ctx, const Container &container, const Item &item) {
        if (std::find(container.begin(), container.end(), item) == container.end()) {
            std::ostringstream ss;
            ss << "Container does not contain item: " << item;
            ctx.Fail(ss.str());
        }
    }

    static void Contains(AssertionContext &ctx, const std::string &str, const std::string &substr) {
        if (str.find(substr) == std::string::npos) {
            std::ostringstream ss;
            ss << "String does not contain substring: \"" << substr << "\"\n"
               << "in string: \"" << str << "\"";
            ctx.Fail(ss.str());
        }
    }

    static void Regex(AssertionContext &ctx, const std::string &str, const std::string &pattern) {
        std::regex re(pattern);
        if (!std::regex_search(str, re)) {
            std::ostringstream ss;
            ss << "String does not match pattern: \"" << pattern << "\"\n"
               << "string: \"" << str << "\"";
            ctx.Fail(ss.str());
        }
    }

    static void FileExists(AssertionContext &ctx, const std::string &path) {
        std::ifstream file(path);
        if (!file.good()) { ctx.Fail("File does not exist: " + path); }
    }

    static void DirExists(AssertionContext &ctx, const std::string &path) {
        if (FILE *f = fopen(path.c_str(), "r")) {
            fclose(f);
        } else {
            ctx.Fail("Directory does not exist: " + path);
        }
    }

    static void JSONEq(AssertionContext &ctx, const std::string &expected, const std::string &actual) {
        if (expected != actual) {
            std::ostringstream ss;
            ss << "JSON not equal: \n"
               << "expected: " << expected << "\n"
               << "actual  : " << actual;
            ctx.Fail(ss.str());
        }
    }

    template<typename Expected, typename Actual>
    static void IsType(AssertionContext &ctx, const Actual &value) {
        if (!std::is_same<Expected, Actual>::value) {
            std::ostringstream ss;
            ss << "Type mismatch: \n"
               << "expected: " << typeid(Expected).name() << "\n"
               << "actual  : " << typeid(Actual).name();
            ctx.Fail(ss.str());
        }
    }

    template<typename T>
    static void Len(AssertionContext &ctx, const std::vector<T> &container, size_t expected) {
        if (container.size() != expected) {
            std::ostringstream ss;
            ss << "Length mismatch: \n"
               << "expected: " << expected << "\n"
               << "actual  : " << container.size();
            ctx.Fail(ss.str());
        }
    }

    static void Len(AssertionContext &ctx, const std::string &str, size_t expected) {
        if (str.length() != expected) {
            std::ostringstream ss;
            ss << "Length mismatch: \n"
               << "expected: " << expected << "\n"
               << "actual  : " << str.length();
            ctx.Fail(ss.str());
        }
    }

    static void Empty(AssertionContext &ctx, const auto &container) {
        if (!container.empty()) { ctx.Fail("Container is not empty"); }
    }

    static void Empty(AssertionContext &ctx, const std::string &str) {
        if (!str.empty()) { ctx.Fail("String is not empty"); }
    }

    static void NotEmpty(AssertionContext &ctx, const auto &container) {
        if (container.empty()) { ctx.Fail("Container is empty"); }
    }

    static void NotEmpty(AssertionContext &ctx, const std::string &str) {
        if (str.empty()) { ctx.Fail("String is empty"); }
    }
};

struct WithMessage {
    AssertionContext &Ctx;
    std::string       Msg;

    WithMessage(AssertionContext &c, const std::string &m) : Ctx(c), Msg(m) {}

    template<typename... Args>
    void Equal(const Args &...args) {
        Ctx.CustomMsg = Msg;
        Assert::Equal(Ctx, args...);
    }

    template<typename... Args>
    void NotEqual(const Args &...args) {
        Ctx.CustomMsg = Msg;
        Assert::NotEqual(Ctx, args...);
    }

    void True(bool condition) {
        Ctx.CustomMsg = Msg;
        Assert::True(Ctx, condition);
    }
};

struct TestSuite {
    std::string                 Name;
    std::function<void()>       Body;
    std::vector<AssertionError> Failures;
};

struct TestRegistry {
    static std::map<std::string, TestSuite> &Tests() {
        static std::map<std::string, TestSuite> registry;
        return registry;
    }

    static void RegisterTest(const std::string &name, const std::function<void()> &body) { Tests()[name] = {name, body, {}}; }

    static int RunAll() {
        int failed = 0;
        for (auto &pair: Tests()) {
            TestSuite &test = pair.second;
            std::cout << "=== RUN   " << test.Name << "\n";

            try {
                test.Body();
                std::cout << "--- PASS: " << test.Name << "\n";
            } catch (const AssertionError &e) {
                test.Failures.push_back(e);
                failed++;
                std::cout << "--- FAIL: " << test.Name << "\n";
                std::cout << "    " << e.File << ":" << e.Line << ": " << e.Message << "\n";
            } catch (const std::exception &e) {
                failed++;
                std::cout << "--- FAIL: " << test.Name << "\n";
                std::cout << "    Unexpected exception: " << e.what() << "\n";
            } catch (...) {
                failed++;
                std::cout << "--- FAIL: " << test.Name << "\n";
                std::cout << "    Unknown exception\n";
            }
        }

        if (failed > 0) {
            std::cout << "\nFAIL: " << failed << " of " << Tests().size() << " tests failed\n";
        } else {
            std::cout << "\nPASS: All " << Tests().size() << " tests passed\n";
        }

        return failed;
    }
};

#define TEST(testName)                                                                   \
    static void testName##_Body();                                                       \
    static struct TestReg_##testName {                                                   \
        TestReg_##testName() { TestRegistry::RegisterTest(#testName, testName##_Body); } \
    } testName##_Reg;                                                                    \
    static void testName##_Body()

#define CTX ::AssertionContext(__FILE__, __LINE__)

TEST(StringEquality) {
    std::string a = "hello";
    std::string b = "hello";
    Assert::Equal(CTX, a, b);
    Assert::Equal(CTX, "world", "world");
}

TEST(VectorOperations) {
    std::vector<int> v = {1, 2, 3};
    Assert::Contains(CTX, v, 2);
    Assert::NotContains(CTX, v, 4);
    Assert::Len(CTX, v, 3);
}

TEST(ErrorHandling) {
    Assert::Error(CTX, [] { throw std::runtime_error("test"); });
    Assert::NoError(CTX, [] { int x = 42; });

    auto ctx = CTX;
    WithMessage(ctx, "Custom error message").Equal(1, 2);
}

TEST(TypeChecking) {
    int x = 42;
    Assert::IsType<int>(CTX, x);
    Assert::IsType<double>(CTX, 3.14);
}

int main() { return TestRegistry::RunAll(); }
