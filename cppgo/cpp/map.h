#pragma once

#include <vector>
#include <algorithm>
#include <functional>
#include <stdexcept>
#include <memory>
#include <mutex>
#include <shared_mutex>
#include <iostream>
#include <iomanip>
#include <cassert>

template<typename K, typename V>
struct Pair {
    K key;
    V value;

    Pair(const K &k, const V &v) : key(k), value(v) {}

    bool operator<(const Pair &other) const { return key < other.key; }
};

template<typename K, typename V>
struct OrderedMap {
    using PairType = Pair<K, V>;
    using Container = std::vector<PairType>;

    Container                 data;
    mutable std::shared_mutex mtx;// 读写锁

    void Store(const K &key, const V &value);
    bool Load(const K &key, V &value) const;
    bool Delete(const K &key);
    int  Size() const;
    bool Empty() const;
    void Clear();

    void Range(std::function<void(const K &, const V &)> func) const;
    void RangeFrom(const K &start, std::function<void(const K &, const V &)> func) const;
    void RangeTo(const K &end, std::function<void(const K &, const V &)> func) const;
    void RangeBetween(const K &start, const K &end, std::function<void(const K &, const V &)> func) const;

    bool     Contains(const K &key) const;
    V       &At(const K &key);
    const V &At(const K &key) const;
    V       &operator[](const K &key);

    typename Container::iterator begin() { return data.begin(); }

    typename Container::const_iterator begin() const { return data.begin(); }

    typename Container::iterator end() { return data.end(); }

    typename Container::const_iterator end() const { return data.end(); }

    typename Container::iterator       LowerBound(const K &key);
    typename Container::const_iterator LowerBound(const K &key) const;
    typename Container::iterator       UpperBound(const K &key);
    typename Container::const_iterator UpperBound(const K &key) const;

private:
    typename Container::iterator       find_pos(const K &key);
    typename Container::const_iterator find_pos(const K &key) const;
};

template<typename K, typename V>
struct SafeOrderedMap {
    OrderedMap<K, V>          map;
    mutable std::shared_mutex mtx;

    void Store(const K &key, const V &value) {
        std::unique_lock lock(mtx);
        map.Store(key, value);
    }

    bool Load(const K &key, V &value) const {
        std::shared_lock lock(mtx);
        return map.Load(key, value);
    }

    bool Delete(const K &key) {
        std::unique_lock lock(mtx);
        return map.Delete(key);
    }

    int Size() const {
        std::shared_lock lock(mtx);
        return map.Size();
    }

    bool Empty() const {
        std::shared_lock lock(mtx);
        return map.Empty();
    }

    void Clear() {
        std::unique_lock lock(mtx);
        map.Clear();
    }

    void Range(std::function<void(const K &, const V &)> func) const {
        std::shared_lock lock(mtx);
        map.Range(func);
    }

    void RangeFrom(const K &start, std::function<void(const K &, const V &)> func) const {
        std::shared_lock lock(mtx);
        map.RangeFrom(start, func);
    }

    void RangeTo(const K &end, std::function<void(const K &, const V &)> func) const {
        std::shared_lock lock(mtx);
        map.RangeTo(end, func);
    }

    void RangeBetween(const K &start, const K &end, std::function<void(const K &, const V &)> func) const {
        std::shared_lock lock(mtx);
        map.RangeBetween(start, end, func);
    }

    bool Contains(const K &key) const {
        std::shared_lock lock(mtx);
        return map.Contains(key);
    }
};

template<typename K, typename V>
void OrderedMap<K, V>::Store(const K &key, const V &value) {
    auto it = find_pos(key);

    if (it != data.end() && it->key == key) {
        it->value = value;
    } else {
        data.insert(it, PairType(key, value));
    }
}

template<typename K, typename V>
bool OrderedMap<K, V>::Load(const K &key, V &value) const {
    auto it = find_pos(key);

    if (it != data.end() && it->key == key) {
        value = it->value;
        return true;
    }
    return false;
}

template<typename K, typename V>
bool OrderedMap<K, V>::Delete(const K &key) {
    auto it = find_pos(key);

    if (it != data.end() && it->key == key) {
        data.erase(it);
        return true;
    }
    return false;
}

template<typename K, typename V>
int OrderedMap<K, V>::Size() const {
    return static_cast<int>(data.size());
}

template<typename K, typename V>
bool OrderedMap<K, V>::Empty() const {
    return data.empty();
}

template<typename K, typename V>
void OrderedMap<K, V>::Clear() {
    data.clear();
}

template<typename K, typename V>
void OrderedMap<K, V>::Range(std::function<void(const K &, const V &)> func) const {
    for (const auto &pair: data) { func(pair.key, pair.value); }
}

template<typename K, typename V>
void OrderedMap<K, V>::RangeFrom(const K &start, std::function<void(const K &, const V &)> func) const {
    auto it = std::lower_bound(data.begin(), data.end(), PairType(start, V{}));
    for (; it != data.end(); ++it) { func(it->key, it->value); }
}

template<typename K, typename V>
void OrderedMap<K, V>::RangeTo(const K &end, std::function<void(const K &, const V &)> func) const {
    auto it = data.begin();
    for (; it != data.end() && it->key <= end; ++it) { func(it->key, it->value); }
}

template<typename K, typename V>
void OrderedMap<K, V>::RangeBetween(const K &start, const K &end, std::function<void(const K &, const V &)> func) const {
    auto begin_it = std::lower_bound(data.begin(), data.end(), PairType(start, V{}));
    auto end_it = std::upper_bound(data.begin(), data.end(), PairType(end, V{}));

    for (; begin_it != end_it; ++begin_it) { func(begin_it->key, begin_it->value); }
}

template<typename K, typename V>
bool OrderedMap<K, V>::Contains(const K &key) const {
    auto it = find_pos(key);
    return (it != data.end() && it->key == key);
}

template<typename K, typename V>
V &OrderedMap<K, V>::At(const K &key) {
    auto it = find_pos(key);
    if (it == data.end() || it->key != key) { throw std::out_of_range("Key not found"); }
    return it->value;
}

template<typename K, typename V>
const V &OrderedMap<K, V>::At(const K &key) const {
    auto it = find_pos(key);
    if (it == data.end() || it->key != key) { throw std::out_of_range("Key not found"); }
    return it->value;
}

template<typename K, typename V>
V &OrderedMap<K, V>::operator[](const K &key) {
    auto it = find_pos(key);

    if (it != data.end() && it->key == key) { return it->value; }

    it = data.insert(it, PairType(key, V{}));
    return it->value;
}

template<typename K, typename V>
typename OrderedMap<K, V>::Container::iterator OrderedMap<K, V>::LowerBound(const K &key) {
    return std::lower_bound(data.begin(), data.end(), PairType(key, V{}));
}

template<typename K, typename V>
typename OrderedMap<K, V>::Container::const_iterator OrderedMap<K, V>::LowerBound(const K &key) const {
    return std::lower_bound(data.begin(), data.end(), PairType(key, V{}));
}

template<typename K, typename V>
typename OrderedMap<K, V>::Container::iterator OrderedMap<K, V>::UpperBound(const K &key) {
    return std::upper_bound(data.begin(), data.end(), PairType(key, V{}));
}

template<typename K, typename V>
typename OrderedMap<K, V>::Container::const_iterator OrderedMap<K, V>::UpperBound(const K &key) const {
    return std::upper_bound(data.begin(), data.end(), PairType(key, V{}));
}

template<typename K, typename V>
typename OrderedMap<K, V>::Container::iterator OrderedMap<K, V>::find_pos(const K &key) {
    return std::lower_bound(data.begin(), data.end(), PairType(key, V{}));
}

template<typename K, typename V>
typename OrderedMap<K, V>::Container::const_iterator OrderedMap<K, V>::find_pos(const K &key) const {
    return std::lower_bound(data.begin(), data.end(), PairType(key, V{}));
}

template<typename K, typename V>
void PrintOrderedMap(const OrderedMap<K, V> &map) {
    std::cout << "OrderedMap (" << map.Size() << " items):\n";
    map.Range([](const K &key, const V &value) { std::cout << "  " << std::setw(10) << key << " => " << value << "\n"; });
}

int main() {
    OrderedMap<int, std::string> map;

    map.Store(3, "Three");
    map.Store(1, "One");
    map.Store(4, "Four");
    map.Store(2, "Two");
    map.Store(5, "Five");

    PrintOrderedMap(map);
    /* 输出:
        OrderedMap (5 items):
                  1 => One
                  2 => Two
                  3 => Three
                  4 => Four
                  5 => Five
    */

    std::string value;
    if (map.Load(3, value)) { std::cout << "\nValue for key 3: " << value << "\n"; }

    std::cout << "Value for key 2: " << map[2] << "\n";

    map.Store(3, "THREE");
    std::cout << "Updated value for key 3: " << map[3] << "\n";

    map.Delete(4);
    std::cout << "\nAfter deleting key 4:\n";
    PrintOrderedMap(map);

    std::cout << "\nKeys between 2 and 4:\n";
    map.RangeBetween(2, 4, [](int key, const std::string &val) { std::cout << "  " << key << ": " << val << "\n"; });

    SafeOrderedMap<std::string, int> safeMap;
    safeMap.Store("apple", 5);
    safeMap.Store("banana", 8);
    safeMap.Store("cherry", 3);

    int count;
    if (safeMap.Load("banana", count)) { std::cout << "\nBanana count: " << count << "\n"; }

    std::cout << "\nSafeOrderedMap contents:\n";
    safeMap.Range([](const std::string &key, int value) { std::cout << "  " << key << ": " << value << "\n"; });

    return 0;
}
