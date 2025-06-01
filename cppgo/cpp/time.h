#pragma once

#include <chrono>
#include <ctime>
#include <string>
#include <sstream>
#include <iomanip>
#include <stdexcept>
#include <cctype>
#include <vector>
#include <regex>
#include <sys/time.h>
#include <unistd.h>
#include <cstring>

constexpr int64_t Nanosecond = 1;
constexpr int64_t Microsecond = 1000 * Nanosecond;
constexpr int64_t Millisecond = 1000 * Microsecond;
constexpr int64_t Second = 1000 * Millisecond;
constexpr int64_t Minute = 60 * Second;
constexpr int64_t Hour = 60 * Minute;
constexpr int64_t Day = 24 * Hour;

constexpr int January = 1;
constexpr int February = 2;
constexpr int March = 3;
constexpr int April = 4;
constexpr int May = 5;
constexpr int June = 6;
constexpr int July = 7;
constexpr int August = 8;
constexpr int September = 9;
constexpr int October = 10;
constexpr int November = 11;
constexpr int December = 12;

constexpr int Sunday = 0;
constexpr int Monday = 1;
constexpr int Tuesday = 2;
constexpr int Wednesday = 3;
constexpr int Thursday = 4;
constexpr int Friday = 5;
constexpr int Saturday = 6;

struct Time {
    int64_t sec; // 从1970-01-01 00:00:00 UTC开始计算的秒数
    int32_t nsec;// 纳秒部分
    int16_t loc; // 时区偏移（分钟）

    static Time Now();
    static Time NowInLocation(int16_t offset);
    static Time Unix(int64_t sec, int32_t nsec);
    static Time Date(int year, int month, int day, int hour, int min, int sec, int nsec, int16_t offset);
    static Time Parse(const std::string &layout, const std::string &value);
};

struct Duration {
    int64_t ns;

    int64_t Nanoseconds() const { return ns; }

    int64_t Microseconds() const { return ns / Microsecond; }

    int64_t Milliseconds() const { return ns / Millisecond; }

    int64_t Seconds() const { return ns / Second; }

    int64_t Minutes() const { return ns / Minute; }

    int64_t Hours() const { return ns / Hour; }

    std::string String() const;
};

struct TimeUtil {
    static Duration Nanoseconds(int64_t n) { return Duration{n}; }

    static Duration Microseconds(int64_t n) { return Duration{n * Microsecond}; }

    static Duration Milliseconds(int64_t n) { return Duration{n * Millisecond}; }

    static Duration Seconds(int64_t n) { return Duration{n * Second}; }

    static Duration Minutes(int64_t n) { return Duration{n * Minute}; }

    static Duration Hours(int64_t n) { return Duration{n * Hour}; }

    static void Sleep(Duration d) {
        struct timespec req;
        req.tv_sec = d.Seconds();
        req.tv_nsec = d.Nanoseconds() % 1000000000;
        nanosleep(&req, NULL);
    }

    static Duration Since(Time t) { return Now().Sub(t); }
};

struct TimeFormat {
    static std::string Format(Time t, const std::string &layout);

    static Time Parse(const std::string &layout, const std::string &value);
};

int Year(Time t) {
    time_t    rawtime = t.sec;
    struct tm tm;
    gmtime_r(&rawtime, &tm);
    return tm.tm_year + 1900;
}

int Month(Time t) {
    time_t    rawtime = t.sec;
    struct tm tm;
    gmtime_r(&rawtime, &tm);
    return tm.tm_mon + 1;
}

int Day(Time t) {
    time_t    rawtime = t.sec;
    struct tm tm;
    gmtime_r(&rawtime, &tm);
    return tm.tm_mday;
}

int Weekday(Time t) {
    time_t    rawtime = t.sec;
    struct tm tm;
    gmtime_r(&rawtime, &tm);
    return tm.tm_wday;
}

int Hour(Time t) {
    time_t    rawtime = t.sec;
    struct tm tm;
    gmtime_r(&rawtime, &tm);
    return tm.tm_hour;
}

int Minute(Time t) {
    time_t    rawtime = t.sec;
    struct tm tm;
    gmtime_r(&rawtime, &tm);
    return tm.tm_min;
}

int Second(Time t) {
    time_t    rawtime = t.sec;
    struct tm tm;
    gmtime_r(&rawtime, &tm);
    return tm.tm_sec;
}

int Nanosecond(Time t) { return t.nsec; }

int YearDay(Time t) {
    time_t    rawtime = t.sec;
    struct tm tm;
    gmtime_r(&rawtime, &tm);
    return tm.tm_yday;
}

Time Add(Time t, Duration d) {
    int64_t totalNsec = static_cast<int64_t>(t.nsec) + d.ns;
    int64_t extraSec = totalNsec / 1000000000;
    return Time{t.sec + static_cast<int64_t>(d.Seconds()) + extraSec, static_cast<int32_t>(totalNsec % 1000000000), t.loc};
}

Duration Sub(Time a, Time b) {
    int64_t secDiff = a.sec - b.sec;
    int64_t nsecDiff = a.nsec - b.nsec;
    return Duration{secDiff * 1000000000 + nsecDiff};
}

bool After(Time a, Time b) {
    if (a.sec != b.sec) return a.sec > b.sec;
    return a.nsec > b.nsec;
}

bool Before(Time a, Time b) {
    if (a.sec != b.sec) return a.sec < b.sec;
    return a.nsec < b.nsec;
}

bool Equal(Time a, Time b) { return a.sec == b.sec && a.nsec == b.nsec; }

Time Time::Now() {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return Time{tv.tv_sec, static_cast<int32_t>(tv.tv_usec * 1000), 0};
}

Time Time::NowInLocation(int16_t offset) {
    struct timeval tv;
    gettimeofday(&tv, NULL);
    return Time{tv.tv_sec, static_cast<int32_t>(tv.tv_usec * 1000), offset};
}

Duration Duration::String() const {
    if (ns < Microsecond) { return std::to_string(ns) + "ns"; }
    if (ns < Millisecond) { return std::to_string(ns / 1000.0) + "µs"; }
    if (ns < Second) { return std::to_string(ns / 1000000.0) + "ms"; }

    int64_t sec = ns / Second;
    int64_t min = sec / 60;
    sec %= 60;
    int64_t hr = min / 60;
    min %= 60;

    std::ostringstream oss;
    if (hr > 0) oss << hr << "h";
    if (min > 0) oss << min << "m";
    if (sec > 0 || (hr == 0 && min == 0)) {
        oss << sec;
        int64_t frac = (ns % Second) / 1000000;// 毫秒部分
        if (frac > 0) oss << "." << std::setw(3) << std::setfill('0') << frac;
        oss << "s";
    }
    return oss.str();
}

std::string TimeFormat::Format(Time t, const std::string &layout) {
    time_t    rawtime = t.sec;
    struct tm tm;
    if (t.loc == 0) {
        gmtime_r(&rawtime, &tm);
    } else {
        time_t adjusted = rawtime + t.loc * 60;
        gmtime_r(&adjusted, &tm);
    }

    std::ostringstream oss;
    const char        *ptr = layout.c_str();

    while (*ptr) {
        if (*ptr == '%') {
            ptr++;
            switch (*ptr) {
                case 'Y':
                    oss << (tm.tm_year + 1900);
                    break;
                case 'm':
                    oss << std::setw(2) << std::setfill('0') << (tm.tm_mon + 1);
                    break;
                case 'd':
                    oss << std::setw(2) << std::setfill('0') << tm.tm_mday;
                    break;
                case 'H':
                    oss << std::setw(2) << std::setfill('0') << tm.tm_hour;
                    break;
                case 'M':
                    oss << std::setw(2) << std::setfill('0') << tm.tm_min;
                    break;
                case 'S':
                    oss << std::setw(2) << std::setfill('0') << tm.tm_sec;
                    break;
                case 's':
                    oss << (t.sec);
                    break;// Unix时间戳
                case 'f':
                    oss << std::setw(9) << std::setfill('0') << t.nsec;
                    break;// 纳秒
                case 'z':
                    if (t.loc == 0) oss << "Z";
                    else
                        oss << std::setw(3) << std::setfill('0') << (t.loc / 60) << ":" << std::setw(2) << (t.loc % 60);
                    break;
                case 'A':
                    oss << (std::string[]{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}[tm.tm_wday]);
                    break;
                case 'a':
                    oss << (std::string[]{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}[tm.tm_wday]);
                    break;
                case 'B':
                    oss << (std::string[]{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}[tm.tm_mon]);
                    break;
                case 'b':
                    oss << (std::string[]{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}[tm.tm_mon]);
                    break;
                default:
                    oss << *ptr;
            }
            ptr++;
        } else {
            oss << *ptr++;
        }
    }

    return oss.str();
}

Time Time::Parse(const std::string &layout, const std::string &value) {
    time_t    now = ::time(nullptr);
    struct tm tm;
    memset(&tm, 0, sizeof(tm));

    const char *fmt = layout.c_str();
    const char *val = value.c_str();

    while (*fmt && *val) {
        if (*fmt == '%') {
            fmt++;
            switch (*fmt) {
                case 'Y':
                    if (sscanf(val, "%4d", &tm.tm_year) != 1) goto parse_error;
                    tm.tm_year -= 1900;
                    val += 4;
                    break;
                case 'm':
                    if (sscanf(val, "%2d", &tm.tm_mon) != 1) goto parse_error;
                    tm.tm_mon -= 1;
                    val += 2;
                    break;
                case 'd':
                    if (sscanf(val, "%2d", &tm.tm_mday) != 1) goto parse_error;
                    val += 2;
                    break;
                case 'H':
                    if (sscanf(val, "%2d", &tm.tm_hour) != 1) goto parse_error;
                    val += 2;
                    break;
                case 'M':
                    if (sscanf(val, "%2d", &tm.tm_min) != 1) goto parse_error;
                    val += 2;
                    break;
                case 'S':
                    if (sscanf(val, "%2d", &tm.tm_sec) != 1) goto parse_error;
                    val += 2;
                    break;
                default:
                    fmt++;
                    continue;
            }
            fmt++;
        } else {
            if (*fmt != *val) goto parse_error;
            fmt++;
            val++;
        }
    }

    tm.tm_isdst = -1;
    time_t sec = mktime(&tm);
    if (sec == -1) goto parse_error;

    return Time{sec, 0, 0};

parse_error:
    throw std::runtime_error("parse error: " + value);
}

int main() {
    Time now = Time::Now();

    std::cout << "Current time: " << TimeFormat::Format(now, "%Y-%m-%d %H:%M:%S") << "\n";

    std::cout << "Year: " << Year(now) << "\n";
    std::cout << "Month: " << Month(now) << "\n";
    std::cout << "Day: " << Day(now) << "\n";
    std::cout << "Hour: " << Hour(now) << "\n";
    std::cout << "Minute: " << Minute(now) << "\n";
    std::cout << "Second: " << Second(now) << "\n";

    Duration oneHour = TimeUtil::Hours(1);
    Time     inOneHour = Add(now, oneHour);

    if (After(inOneHour, now)) { std::cout << "inOneHour is after now\n"; }

    Duration diff = Sub(inOneHour, now);
    std::cout << "Difference: " << diff.Hours() << " hours\n";

    Time parsed = Time::Parse("2006-01-02", "2023-12-25");
    std::cout << "Parsed time: " << TimeFormat::Format(parsed, "%B %d, %Y") << "\n";

    std::cout << "Sleeping for 1 second...\n";
    TimeUtil::Sleep(TimeUtil::Seconds(1));
    std::cout << "Awake!\n";

    return 0;
}
