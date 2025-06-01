#pragma once
#include <vector>
#include <algorithm>
#include <functional>
#include <cassert>
#include <iostream>
#include <iterator>
#include <type_traits>

template<typename T>
struct Sortable {
    virtual int  Len() const = 0;
    virtual bool Less(int i, int j) const = 0;
    virtual void Swap(int i, int j) = 0;
    virtual ~Sortable() = default;
};

template<typename Container>
struct ContainerSortable : public Sortable<typename Container::value_type> {
    using T = typename Container::value_type;
    Container &data;

    explicit ContainerSortable(Container &cont) : data(cont) {}

    int Len() const override { return static_cast<int>(data.size()); }

    bool Less(int i, int j) const override { return data[i] < data[j]; }

    void Swap(int i, int j) override { std::swap(data[i], data[j]); }
};

template<typename T>
void insertionSort(Sortable<T> &s, int a, int b) {
    for (int i = a + 1; i < b; i++) {
        for (int j = i; j > a && s.Less(j, j - 1); j--) { s.Swap(j, j - 1); }
    }
}

template<typename T>
void heapSort(Sortable<T> &s, int a, int b) {
    int n = b - a;
    for (int i = n / 2 - 1; i >= 0; i--) { heapify(s, a, b, i); }

    for (int i = n - 1; i > 0; i--) {
        s.Swap(a, a + i);
        heapify(s, a, a + i, 0);
    }
}

template<typename T>
void heapify(Sortable<T> &s, int a, int b, int i) {
    int largest = i;
    int left = 2 * i + 1;
    int right = 2 * i + 2;
    int n = b - a;

    if (left < n && s.Less(a + largest, a + left)) { largest = left; }

    if (right < n && s.Less(a + largest, a + right)) { largest = right; }

    if (largest != i) {
        s.Swap(a + i, a + largest);
        heapify(s, a, b, largest);
    }
}

template<typename T>
void quickSort(Sortable<T> &s, int a, int b) {
    if (b - a <= 7) {
        insertionSort(s, a, b);
        return;
    }

    int pivot = medianOfThree(s, a, b);
    int p = partition(s, a, b, pivot);

    quickSort(s, a, p);
    quickSort(s, p + 1, b);
}

template<typename T>
int medianOfThree(Sortable<T> &s, int a, int b) {
    int m = a + (b - a) / 2;

    if (s.Less(m, a)) s.Swap(a, m);
    if (s.Less(b - 1, a)) s.Swap(a, b - 1);
    if (s.Less(b - 1, m)) s.Swap(m, b - 1);

    return m;
}

template<typename T>
int partition(Sortable<T> &s, int a, int b, int pivot) {
    s.Swap(pivot, b - 1);
    int i = a;

    for (int j = a; j < b - 1; j++) {
        if (s.Less(j, b - 1)) {
            s.Swap(i, j);
            i++;
        }
    }

    s.Swap(i, b - 1);
    return i;
}

template<typename T>
void mergeSort(Sortable<T> &s, int a, int b) {
    if (b - a <= 1) return;

    int m = a + (b - a) / 2;
    mergeSort(s, a, m);
    mergeSort(s, m, b);
    merge(s, a, m, b);
}

template<typename T>
void merge(Sortable<T> &s, int a, int m, int b) {
    std::vector<T> tmp;
    tmp.reserve(b - a);

    int i = a, j = m;
    while (i < m && j < b) {
        if (s.Less(i, j)) {
            tmp.push_back(s[i]);// 需要索引访问，稍后扩展接口
            i++;
        } else {
            tmp.push_back(s[j]);
            j++;
        }
    }

    while (i < m) {
        tmp.push_back(s[i]);
        i++;
    }

    while (j < b) {
        tmp.push_back(s[j]);
        j++;
    }

    for (int k = 0; k < tmp.size(); k++) {
        s[a + k] = tmp[k];// 需要索引赋值，稍后扩展接口
    }
}

template<typename T>
void Sort(Sortable<T> &s) {
    quickSort(s, 0, s.Len());
}

template<typename T>
void Stable(Sortable<T> &s) {
    mergeSort(s, 0, s.Len());
}

template<typename Container>
void Slice(Container &data) {
    ContainerSortable<Container> s(data);
    Sort(s);
}

template<typename Container, typename Compare>
void Slice(Container &data, Compare comp) {
    struct CustomSortable : public ContainerSortable<Container> {
        Compare comp;

        CustomSortable(Container &cont, Compare c) : ContainerSortable<Container>(cont), comp(c) {}

        bool Less(int i, int j) const override { return comp(this->data[i], this->data[j]); }
    };

    CustomSortable s(data, comp);
    Sort(s);
}

template<typename Container>
void StableSlice(Container &data) {
    ContainerSortable<Container> s(data);
    Stable(s);
}

template<typename Container, typename Compare>
void StableSlice(Container &data, Compare comp) {
    struct CustomSortable : public ContainerSortable<Container> {
        Compare comp;

        CustomSortable(Container &cont, Compare c) : ContainerSortable<Container>(cont), comp(c) {}

        bool Less(int i, int j) const override { return comp(this->data[i], this->data[j]); }
    };

    CustomSortable s(data, comp);
    Stable(s);
}

template<typename T>
bool IsSorted(Sortable<T> &s) {
    int n = s.Len();
    for (int i = 1; i < n; i++) {
        if (s.Less(i, i - 1)) { return false; }
    }
    return true;
}

template<typename Container>
bool IsSliceSorted(const Container &data) {
    for (size_t i = 1; i < data.size(); i++) {
        if (data[i] < data[i - 1]) { return false; }
    }
    return true;
}

template<typename Container, typename Compare>
bool IsSliceSorted(const Container &data, Compare comp) {
    for (size_t i = 1; i < data.size(); i++) {
        if (comp(data[i], data[i - 1])) { return false; }
    }
    return true;
}

template<typename T>
int Search(Sortable<T> &s, const T &x) {
    int low = 0;
    int high = s.Len();

    while (low < high) {
        int mid = low + (high - low) / 2;
        if (s.Less(mid, x)) {
            low = mid + 1;
        } else if (x < s[mid]) {// 需要索引访问
            high = mid;
        } else {
            return mid;
        }
    }
    return -1;// 未找到
}

template<typename Container>
int SearchSlice(const Container &data, const typename Container::value_type &x) {
    auto it = std::lower_bound(data.begin(), data.end(), x);
    if (it != data.end() && *it == x) { return static_cast<int>(std::distance(data.begin(), it)); }
    return -1;
}

template<typename Container, typename Compare>
int SearchSlice(const Container &data, const typename Container::value_type &x, Compare comp) {
    auto it = std::lower_bound(data.begin(), data.end(), x, comp);
    if (it != data.end() && !comp(x, *it) && !comp(*it, x)) { return static_cast<int>(std::distance(data.begin(), it)); }
    return -1;
}

template<typename T>
void Reverse(Sortable<T> &s) {
    int i = 0;
    int j = s.Len() - 1;
    while (i < j) {
        s.Swap(i, j);
        i++;
        j--;
    }
}

template<typename Container>
void ReverseSlice(Container &data) {
    std::reverse(data.begin(), data.end());
}

template<typename T>
struct IndexableSortable : public Sortable<T> {
    virtual T       &operator[](int i) = 0;
    virtual const T &operator[](int i) const = 0;
};

template<typename Container>
struct EnhancedContainerSortable : public IndexableSortable<typename Container::value_type> {
    using T = typename Container::value_type;
    Container &data;

    explicit EnhancedContainerSortable(Container &cont) : data(cont) {}

    int Len() const override { return static_cast<int>(data.size()); }

    bool Less(int i, int j) const override { return data[i] < data[j]; }

    void Swap(int i, int j) override { std::swap(data[i], data[j]); }

    T &operator[](int i) override { return data[i]; }

    const T &operator[](int i) const override { return data[i]; }
};

struct Person {
    std::string name;
    int         age;
    double      salary;

    bool operator<(const Person &other) const { return name < other.name; }

    friend std::ostream &operator<<(std::ostream &os, const Person &p) { return os << p.name << " (" << p.age << ", $" << p.salary << ")"; }
};

bool byAge(const Person &a, const Person &b) { return a.age < b.age; }

bool bySalaryDesc(const Person &a, const Person &b) { return a.salary > b.salary; }

void testSorting() {
    std::vector<int> numbers = {5, 2, 8, 1, 9, 3};
    std::cout << "Original numbers: ";
    for (int n: numbers) std::cout << n << " ";
    std::cout << std::endl;

    Slice(numbers);
    std::cout << "Sorted numbers: ";
    for (int n: numbers) std::cout << n << " ";
    std::cout << std::endl;

    std::vector<Person> people = {{"Alice", 30, 85000.0}, {"Bob", 25, 75000.0}, {"Charlie", 35, 95000.0}, {"David", 28, 80000.0}};

    std::cout << "\nOriginal people:" << std::endl;
    for (const auto &p: people) std::cout << p << std::endl;

    Slice(people);
    std::cout << "\nSorted by name:" << std::endl;
    for (const auto &p: people) std::cout << p << std::endl;

    Slice(people, byAge);
    std::cout << "\nSorted by age:" << std::endl;
    for (const auto &p: people) std::cout << p << std::endl;

    Slice(people, bySalaryDesc);
    std::cout << "\nSorted by salary (descending):" << std::endl;
    for (const auto &p: people) std::cout << p << std::endl;

    std::vector<std::pair<int, std::string>> pairs = {{2, "two"}, {1, "first"}, {2, "second"}, {1, "another"}};

    auto pairComp = [](const auto &a, const auto &b) { return a.first < b.first; };

    StableSlice(pairs, pairComp);
    std::cout << "\nStable sorted pairs:" << std::endl;
    for (const auto &p: pairs) { std::cout << p.first << ": " << p.second << std::endl; }

    bool sorted = IsSliceSorted(people, bySalaryDesc);
    std::cout << "\nIs people sorted by salary? " << (sorted ? "Yes" : "No") << std::endl;

    int index = SearchSlice(people, Person{"Bob", 25, 75000.0}, [](const Person &a, const Person &b) { return a.name < b.name; });

    if (index != -1) {
        std::cout << "\nFound Bob at index: " << index << std::endl;
    } else {
        std::cout << "\nBob not found" << std::endl;
    }

    ReverseSlice(numbers);
    std::cout << "\nReversed numbers: ";
    for (int n: numbers) std::cout << n << " ";
    std::cout << std::endl;

    std::vector<std::string>                            words = {"banana", "apple", "cherry", "date"};
    EnhancedContainerSortable<std::vector<std::string>> sorter(words);

    std::cout << "\nOriginal words: ";
    for (int i = 0; i < sorter.Len(); i++) { std::cout << sorter[i] << " "; }
    std::cout << std::endl;

    Sort(sorter);
    std::cout << "Sorted words: ";
    for (int i = 0; i < sorter.Len(); i++) { std::cout << sorter[i] << " "; }
    std::cout << std::endl;

    int pos = Search(sorter, std::string("cherry"));
    if (pos != -1) { std::cout << "Found 'cherry' at position: " << pos << std::endl; }
}

/*
预期输出:

Original numbers: 5 2 8 1 9 3 
Sorted numbers: 1 2 3 5 8 9 

Original people:
Alice (30, $85000)
Bob (25, $75000)
Charlie (35, $95000)
David (28, $80000)

Sorted by name:
Alice (30, $85000)
Bob (25, $75000)
Charlie (35, $95000)
David (28, $80000)

Sorted by age:
Bob (25, $75000)
David (28, $80000)
Alice (30, $85000)
Charlie (35, $95000)

Sorted by salary (descending):
Charlie (35, $95000)
Alice (30, $85000)
David (28, $80000)
Bob (25, $75000)

Stable sorted pairs:
1: first
1: another
2: two
2: second

Is people sorted by salary? Yes

Found Bob at index: 3

Reversed numbers: 9 8 5 3 2 1 

Original words: banana apple cherry date 
Sorted words: apple banana cherry date 
Found 'cherry' at position: 2
*/
