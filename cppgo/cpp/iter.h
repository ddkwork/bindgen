#pragma once
#include <functional>
#include <vector>
#include <map>
#include <stdexcept>
#include <concepts>

template<typename V>
struct Seq {
    using YieldFunc = std::function<bool(V)>;
    std::function<void(YieldFunc)> iterFunc;

    void operator()(YieldFunc yield) const {
        if (iterFunc) iterFunc(yield);
    }

    template<typename Out>
    Seq<Out> Map(std::function<Out(V)> mapper) const {
        return Seq<Out>{[=](std::function<bool(Out)> yield) { iterFunc([&](V value) { return yield(mapper(value)); }); }};
    }

    Seq<V> Filter(std::function<bool(V)> predicate) const {
        return Seq<V>{[=](YieldFunc yield) {
            iterFunc([&](V value) {
                if (predicate(value)) { return yield(value); }
                return true;
            });
        }};
    }

    std::vector<V> Collect() const {
        std::vector<V> result;
        iterFunc([&](V value) {
            result.push_back(value);
            return true;
        });
        return result;
    }

    void ForEach(std::function<void(V)> action) const {
        iterFunc([&](V value) {
            action(value);
            return true;
        });
    }

    V Find(std::function<bool(V)> predicate, V defaultValue = V{}) const {
        V    result = defaultValue;
        bool found = false;

        iterFunc([&](V value) {
            if (predicate(value)) {
                result = value;
                found = true;
                return false;
            }
            return true;
        });

        return found ? result : defaultValue;
    }
};

template<typename K, typename V>
struct Seq2 {
    using YieldFunc = std::function<bool(K, V)>;
    std::function<void(YieldFunc)> iterFunc;

    void operator()(YieldFunc yield) const {
        if (iterFunc) iterFunc(yield);
    }

    Seq<K> Keys() const {
        return Seq<K>{[=](std::function<bool(K)> yield) { iterFunc([&](K key, V) { return yield(key); }); }};
    }

    Seq<V> Values() const {
        return Seq<V>{[=](std::function<bool(V)> yield) { iterFunc([&](K, V value) { return yield(value); }); }};
    }

    template<typename K2, typename V2>
    Seq2<K2, V2> MapPairs(std::function<std::pair<K2, V2>(K, V)> mapper) const {
        return Seq2<K2, V2>{[=](std::function<bool(K2, V2)> yield) {
            iterFunc([&](K key, V value) {
                auto [k2, v2] = mapper(key, value);
                return yield(k2, v2);
            });
        }};
    }

    std::map<K, V> CollectMap() const {
        std::map<K, V> result;
        iterFunc([&](K key, V value) {
            result[key] = value;
            return true;
        });
        return result;
    }
};

template<std::integral T>
Seq<T> Range(T start, T end, T step = 1) {
    return Seq<T>{[=](typename Seq<T>::YieldFunc yield) {
        if (step == 0) throw std::invalid_argument("step cannot be zero");

        if (step > 0) {
            for (T i = start; i < end; i += step) {
                if (!yield(i)) break;
            }
        } else {
            for (T i = start; i > end; i += step) {
                if (!yield(i)) break;
            }
        }
    }};
}

template<typename T>
Seq<T> Slice(const std::vector<T> &vec, int start = 0, int end = -1) {
    return Seq<T>{[=, &vec](typename Seq<T>::YieldFunc yield) {
        int n = static_cast<int>(vec.size());
        int s = (start < 0) ? n + start : start;
        int e = (end < 0) ? n + end : end;
        if (e > n) e = n;

        for (int i = s; i < e; i++) {
            if (!yield(vec[i])) break;
        }
    }};
}

template<typename K, typename V, typename MapType>
Seq2<K, V> IterMap(const MapType &map) {
    return Seq2<K, V>{[&](typename Seq2<K, V>::YieldFunc yield) {
        for (const auto &pair: map) {
            if (!yield(pair.first, pair.second)) break;
        }
    }};
}

template<typename T>
Seq<T> IterVec(const std::vector<T> &vec) {
    return Seq<T>{[&](typename Seq<T>::YieldFunc yield) {
        for (const auto &item: vec) {
            if (!yield(item)) break;
        }
    }};
}

template<typename T>
Seq<T> Chain(std::vector<Seq<T>> seqs) {
    return Seq<T>{[=](typename Seq<T>::YieldFunc yield) {
        for (auto &seq: seqs) {
            bool shouldContinue = true;
            seq([&](T value) {
                shouldContinue = yield(value);
                return shouldContinue;
            });
            if (!shouldContinue) break;
        }
    }};
}

template<typename T>
Seq<T> Take(Seq<T> seq, int count) {
    return Seq<T>{[=](typename Seq<T>::YieldFunc yield) {
        int taken = 0;
        seq([&](T value) {
            if (taken++ < count) { return yield(value); }
            return false;
        });
    }};
}

template<typename T>
Seq<T> Skip(Seq<T> seq, int count) {
    return Seq<T>{[=](typename Seq<T>::YieldFunc yield) {
        int skipped = 0;
        seq([&](T value) {
            if (skipped++ < count) { return true; }
            return yield(value);
        });
    }};
}

void testIterFunctions() {
    auto range = Range(1, 6);
    auto squares = range.Map<int>([](int x) { return x * x; });
    auto result = squares.Collect();

    auto evens = Range(1, 11).Filter([](int x) { return x % 2 == 0; }).Collect();

    auto chained = Chain<int>({Range(1, 4), Range(10, 13)}).Collect();

    std::map<std::string, int> scores = {{"Alice", 90}, {"Bob", 85}};
    auto                       scoreIter = IterMap(scores);

    auto highScorers = scoreIter.Keys().Filter([](const std::string &name) { return name != "Bob"; }).Collect();

    auto partial = Range(1, 11).Skip(3).Take(4).Collect();

    Range(1, 4).ForEach([](int x) { std::cout << x << " "; });
}
