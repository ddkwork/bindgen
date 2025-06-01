#pragma once

#include <mutex>
#include <shared_mutex>
#include <condition_variable>
#include <atomic>
#include <functional>
#include <unordered_map>
#include <vector>
#include <memory>
#include <cassert>

struct Mutex {
    std::mutex mtx;

    void Lock() { mtx.lock(); }

    void Unlock() { mtx.unlock(); }
};

struct RWMutex {
    std::shared_mutex mtx;

    void Lock() { mtx.lock(); }

    void Unlock() { mtx.unlock(); }

    void RLock() { mtx.lock_shared(); }

    void RUnlock() { mtx.unlock_shared(); }
};

struct WaitGroup {
    std::atomic<int>        counter{0};
    std::mutex              mtx;
    std::condition_variable cv;

    void Add(int delta) {
        counter += delta;
        if (counter < 0) { throw std::runtime_error("negative WaitGroup counter"); }
    }

    void Done() {
        if (--counter == 0) { cv.notify_all(); }
    }

    void Wait() {
        std::unique_lock<std::mutex> lock(mtx);
        cv.wait(lock, [this] { return counter == 0; });
    }
};

struct Cond {
    std::condition_variable cv;

    void Wait(Mutex *mtx) {
        std::unique_lock<std::mutex> lock(mtx->mtx, std::adopt_lock);
        cv.wait(lock);
        lock.release();// 保持锁状态
    }

    void Signal() { cv.notify_one(); }

    void Broadcast() { cv.notify_all(); }
};

struct Once {
    std::once_flag flag;

    void Do(std::function<void()> f) { std::call_once(flag, f); }
};

template<typename K, typename V>
struct Map {
    RWMutex                  mtx;
    std::unordered_map<K, V> data;

    void Store(const K &key, const V &value) {
        std::unique_lock lock(mtx.mtx);
        data[key] = value;
    }

    V Load(const K &key) {
        std::shared_lock lock(mtx.mtx);
        auto             it = data.find(key);
        if (it != data.end()) { return it->second; }
        throw std::out_of_range("key not found");
    }

    bool LoadOrStore(const K &key, const V &value, V &loaded) {
        std::unique_lock lock(mtx.mtx);
        auto             it = data.find(key);
        if (it != data.end()) {
            loaded = it->second;
            return true;
        }
        data[key] = value;
        return false;
    }

    void Delete(const K &key) {
        std::unique_lock lock(mtx.mtx);
        data.erase(key);
    }

    void Range(std::function<void(const K &, const V &)> f) {
        std::shared_lock lock(mtx.mtx);
        for (const auto &[key, value]: data) { f(key, value); }
    }
};

template<typename T>
struct Pool {
    Mutex                           mtx;
    std::vector<std::unique_ptr<T>> items;

    T *Get() {
        std::lock_guard lock(mtx.mtx);
        if (items.empty()) { return new T(); }
        auto ptr = items.back().release();
        items.pop_back();
        return ptr;
    }

    void Put(T *item) {
        std::lock_guard lock(mtx.mtx);
        items.emplace_back(item);
    }
};

template<typename T>
struct Value {
    Mutex mtx;
    T     value;

    T Load() {
        std::lock_guard lock(mtx.mtx);
        return value;
    }

    void Store(const T &newValue) {
        std::lock_guard lock(mtx.mtx);
        value = newValue;
    }

    T Swap(const T &newValue) {
        std::lock_guard lock(mtx.mtx);
        T               old = value;
        value = newValue;
        return old;
    }

    bool CompareAndSwap(const T &oldValue, const T &newValue) {
        std::lock_guard lock(mtx.mtx);
        if (value == oldValue) {
            value = newValue;
            return true;
        }
        return false;
    }
};

struct MutexGuard {
    Mutex *mtx;

    explicit MutexGuard(Mutex *m) : mtx(m) { mtx->Lock(); }

    ~MutexGuard() { mtx->Unlock(); }

    MutexGuard(const MutexGuard &) = delete;
    MutexGuard &operator=(const MutexGuard &) = delete;
};

struct RLockGuard {
    RWMutex *rw;

    explicit RLockGuard(RWMutex *m) : rw(m) { rw->RLock(); }

    ~RLockGuard() { rw->RUnlock(); }

    RLockGuard(const RLockGuard &) = delete;
    RLockGuard &operator=(const RLockGuard &) = delete;
};

struct LockGuard {
    RWMutex *rw;

    explicit LockGuard(RWMutex *m) : rw(m) { rw->Lock(); }

    ~LockGuard() { rw->Unlock(); }

    LockGuard(const LockGuard &) = delete;
    LockGuard &operator=(const LockGuard &) = delete;
};

void sync_example() {
    Mutex mtx;
    { MutexGuard guard(&mtx); }

    WaitGroup wg;
    wg.Add(1);
    std::thread t([&] { wg.Done(); });
    wg.Wait();
    t.join();

    Map<std::string, int> safeMap;
    safeMap.Store("key", 42);
    int value = safeMap.Load("key");

    Once once;
    once.Do([] {});

    Mutex condMtx;
    Cond  cond;
    bool  ready = false;

    std::thread worker([&] {
        MutexGuard guard(&condMtx);
        while (!ready) { cond.Wait(&condMtx); }
    });

    {
        MutexGuard guard(&condMtx);
        ready = true;
        cond.Signal();
    }

    worker.join();

    Pool<std::vector<int>> vecPool;
    auto                  *vec = vecPool.Get();
    vec->push_back(1);
    vecPool.Put(vec);

    Value<int> atomicInt;
    atomicInt.Store(10);
    int  old = atomicInt.Swap(20);
    bool swapped = atomicInt.CompareAndSwap(20, 30);
}
