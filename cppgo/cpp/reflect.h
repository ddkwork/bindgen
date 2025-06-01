#pragma once
#include <any>
#include <typeinfo>
#include <type_traits>
#include <functional>
#include <vector>
#include <map>
#include <string>
#include <memory>
#include <stdexcept>
#include <cassert>
#include <iostream>
#include <unordered_map>

enum class Kind {
    Invalid,     // 无效类型
    Bool,        // 布尔
    Int,         // 整型
    Int8,        // 8位整型
    Int16,       // 16位整型
    Int32,       // 32位整型
    Int64,       // 64位整型
    Uint,        // 无符号整型
    Uint8,       // 8位无符号整型
    Uint16,      // 16位无符号整型
    Uint32,      // 32位无符号整型
    Uint64,      // 64位无符号整型
    Uintptr,     // 指针大小无符号整型
    Float32,     // 32位浮点
    Float64,     // 64位浮点
    Complex64,   // 64位复数
    Complex128,  // 128位复数
    Array,       // 数组
    Chan,        // 通道
    Func,        // 函数
    Interface,   // 接口
    Map,         // 映射
    Ptr,         // 指针
    Slice,       // 切片
    String,      // 字符串
    Struct,      // 结构体
    UnsafePointer// 不安全指针
};

enum Flag {
    FlagNone = 0,
    FlagRO = 1 << 0,          // 只读
    FlagIndir = 1 << 1,       // 间接存储
    FlagAddr = 1 << 2,        // 可寻址
    FlagMethod = 1 << 3,      // 方法
    FlagExported = 1 << 4,    // 导出字段
    FlagAnonymous = 1 << 5,   // 匿名字段
    FlagEmbedded = 1 << 6,    // 嵌入字段
    FlagInterface = 1 << 7,   // 接口类型
    FlagNamed = 1 << 8,       // 命名类型
    FlagComparable = 1 << 9,  // 可比较
    FlagAssignable = 1 << 10, // 可赋值
    FlagConvertible = 1 << 11,// 可转换
    FlagVariadic = 1 << 12    // 可变参数函数
};

struct Type {
    virtual ~Type() = default;

    virtual Kind        kind() const = 0;
    virtual std::string name() const = 0;
    virtual std::string string() const = 0;
    virtual int         size() const = 0;
    virtual int         align() const = 0;
    virtual int         flags() const = 0;
    virtual bool        comparable() const = 0;
    virtual bool        assignableTo(const Type &other) const = 0;
    virtual bool        convertibleTo(const Type &other) const = 0;

    virtual int numField() const { return 0; }

    virtual const Type *field(int i) const { return nullptr; }

    virtual std::string fieldName(int i) const { return ""; }

    virtual int fieldOffset(int i) const { return 0; }

    virtual int fieldIndex(const std::string &name) const { return -1; }

    virtual int numMethod() const { return 0; }

    virtual const Type *method(int i) const { return nullptr; }

    virtual std::string methodName(int i) const { return ""; }

    virtual const Type *elem() const { return nullptr; }// 元素类型 (指针/切片/数组)

    virtual const Type *key() const { return nullptr; }// 键类型 (映射)

    virtual const Type *value() const { return nullptr; }// 值类型 (映射)

    virtual int len() const { return 0; }// 长度 (数组)

    virtual bool variadic() const { return false; }// 是否可变参数 (函数)

    virtual int numIn() const { return 0; }// 输入参数数量 (函数)

    virtual int numOut() const { return 0; }// 输出参数数量 (函数)

    virtual const Type *in(int i) const { return nullptr; }// 输入参数类型 (函数)

    virtual const Type *out(int i) const { return nullptr; }// 输出参数类型 (函数)

    bool operator==(const Type &other) const { return kind() == other.kind() && name() == other.name(); }

    bool operator!=(const Type &other) const { return !(*this == other); }
};

struct Value {
    virtual ~Value() = default;

    virtual const Type &type() const = 0;
    virtual bool        isValid() const = 0;
    virtual bool        canAddr() const = 0;
    virtual bool        canSet() const = 0;
    virtual bool        canInterface() const = 0;
    virtual bool        isZero() const = 0;

    virtual Value addr() const = 0;      // 获取地址
    virtual Value elem() const = 0;      // 解引用指针
    virtual Value field(int i) const = 0;// 获取结构体字段
    virtual Value fieldByName(const std::string &name) const = 0;
    virtual Value index(int i) const = 0;                        // 索引访问
    virtual Value mapIndex(const Value &key) const = 0;          // 映射访问
    virtual Value call(const std::vector<Value> &args) const = 0;// 函数调用

    virtual bool        boolValue() const = 0;
    virtual int64_t     intValue() const = 0;
    virtual uint64_t    uintValue() const = 0;
    virtual double      floatValue() const = 0;
    virtual std::string stringValue() const = 0;
    virtual void       *pointerValue() const = 0;

    virtual void setBool(bool v) = 0;
    virtual void setInt(int64_t v) = 0;
    virtual void setUint(uint64_t v) = 0;
    virtual void setFloat(double v) = 0;
    virtual void setString(const std::string &v) = 0;
    virtual void setPointer(void *v) = 0;

    virtual std::any interfaceValue() const = 0;

    virtual std::string string() const = 0;
};

template<typename T>
struct BasicType : public Type {
    Kind        kind() const override;
    std::string name() const override;

    std::string string() const override { return name(); }

    int size() const override { return sizeof(T); }

    int align() const override { return alignof(T); }

    int  flags() const override;
    bool comparable() const override;
    bool assignableTo(const Type &other) const override;
    bool convertibleTo(const Type &other) const override;
};

template<typename T>
struct TypeTraits;

template<>
struct TypeTraits<bool> {
    static constexpr Kind        kind = Kind::Bool;
    static constexpr const char *name = "bool";
    static constexpr int         flags = FlagComparable | FlagAssignable;
};

template<>
struct TypeTraits<int> {
    static constexpr Kind        kind = Kind::Int;
    static constexpr const char *name = "int";
    static constexpr int         flags = FlagComparable | FlagAssignable;
};

template<>
struct TypeTraits<int8_t> {
    static constexpr Kind        kind = Kind::Int8;
    static constexpr const char *name = "int8";
    static constexpr int         flags = FlagComparable | FlagAssignable;
};

template<>
struct TypeTraits<float> {
    static constexpr Kind        kind = Kind::Float32;
    static constexpr const char *name = "float32";
    static constexpr int         flags = FlagComparable | FlagAssignable;
};

template<>
struct TypeTraits<double> {
    static constexpr Kind        kind = Kind::Float64;
    static constexpr const char *name = "float64";
    static constexpr int         flags = FlagComparable | FlagAssignable;
};

template<>
struct TypeTraits<std::string> {
    static constexpr Kind        kind = Kind::String;
    static constexpr const char *name = "string";
    static constexpr int         flags = FlagComparable | FlagAssignable;
};

template<typename T>
Kind BasicType<T>::kind() const {
    return TypeTraits<T>::kind;
}

template<typename T>
std::string BasicType<T>::name() const {
    return TypeTraits<T>::name;
}

template<typename T>
int BasicType<T>::flags() const {
    return TypeTraits<T>::flags;
}

template<typename T>
bool BasicType<T>::comparable() const {
    return flags() & FlagComparable;
}

template<typename T>
bool BasicType<T>::assignableTo(const Type &other) const {
    return *this == other;
}

template<typename T>
bool BasicType<T>::convertibleTo(const Type &other) const {
    return kind() == other.kind();
}

struct StructField {
    std::string name;
    const Type *type;
    int         offset;
    int         flags;
    std::string tag;
};

struct StructType : public Type {
    std::string                typeName;
    std::vector<StructField>   fields;
    std::map<std::string, int> fieldIndex;

    StructType(std::string name) : typeName(std::move(name)) {}

    Kind kind() const override { return Kind::Struct; }

    std::string name() const override { return typeName; }

    std::string string() const override { return "struct " + typeName; }

    int size() const override;
    int align() const override;

    int flags() const override { return FlagComparable | FlagAssignable; }

    bool comparable() const override { return true; }

    bool assignableTo(const Type &other) const override;
    bool convertibleTo(const Type &other) const override;

    int numField() const override { return static_cast<int>(fields.size()); }

    const Type *field(int i) const override { return i < numField() ? fields[i].type : nullptr; }

    std::string fieldName(int i) const override { return i < numField() ? fields[i].name : ""; }

    int fieldOffset(int i) const override { return i < numField() ? fields[i].offset : 0; }

    int fieldIndex(const std::string &name) const override {
        auto it = fieldIndex.find(name);
        return it != fieldIndex.end() ? it->second : -1;
    }

    void addField(std::string name, const Type *type, int flags = FlagExported) {
        int offset = fields.empty() ? 0 : fields.back().offset + fields.back().type->size();
        fields.push_back({std::move(name), type, offset, flags, ""});
        fieldIndex[fields.back().name] = static_cast<int>(fields.size()) - 1;
    }
};

struct PtrType : public Type {
    const Type *elemType;

    explicit PtrType(const Type *elem) : elemType(elem) {}

    Kind kind() const override { return Kind::Ptr; }

    std::string name() const override { return "*" + elemType->name(); }

    std::string string() const override { return "*" + elemType->string(); }

    int size() const override { return sizeof(void *); }

    int align() const override { return alignof(void *); }

    int flags() const override { return FlagComparable | FlagAssignable; }

    bool comparable() const override { return true; }

    bool assignableTo(const Type &other) const override;
    bool convertibleTo(const Type &other) const override;

    const Type *elem() const override { return elemType; }
};

struct SliceType : public Type {
    const Type *elemType;

    explicit SliceType(const Type *elem) : elemType(elem) {}

    Kind kind() const override { return Kind::Slice; }

    std::string name() const override { return "[]" + elemType->name(); }

    std::string string() const override { return "[]" + elemType->string(); }

    int size() const override { return sizeof(void *) * 3; }// ptr, len, cap

    int align() const override { return alignof(void *); }

    int flags() const override { return FlagComparable | FlagAssignable; }

    bool comparable() const override { return false; }// 切片不可比较

    bool assignableTo(const Type &other) const override;
    bool convertibleTo(const Type &other) const override;

    const Type *elem() const override { return elemType; }
};

struct FuncType : public Type {
    const Type               *receiver;// 接收器类型 (方法)
    std::vector<const Type *> inTypes;
    std::vector<const Type *> outTypes;
    bool                      isVariadic;

    FuncType(std::vector<const Type *> in, std::vector<const Type *> out, bool variadic = false)
        : inTypes(std::move(in)),
          outTypes(std::move(out)),
          isVariadic(variadic),
          receiver(nullptr) {}

    Kind kind() const override { return Kind::Func; }

    std::string name() const override { return "func"; }

    std::string string() const override;

    int size() const override { return sizeof(void *); }// 函数指针大小

    int align() const override { return alignof(void *); }

    int flags() const override { return FlagComparable; }

    bool comparable() const override { return false; }// 函数不可比较

    bool assignableTo(const Type &other) const override;
    bool convertibleTo(const Type &other) const override;

    bool variadic() const override { return isVariadic; }

    int numIn() const override { return static_cast<int>(inTypes.size()); }

    int numOut() const override { return static_cast<int>(outTypes.size()); }

    const Type *in(int i) const override { return i < numIn() ? inTypes[i] : nullptr; }

    const Type *out(int i) const override { return i < numOut() ? outTypes[i] : nullptr; }
};

template<typename T>
struct BasicValue : public Value {
    T           value;
    const Type *typePtr;
    int         flags;

    BasicValue(T v, const Type *type, int flags = FlagNone) : value(std::move(v)), typePtr(type), flags(flags) {}

    const Type &type() const override { return *typePtr; }

    bool isValid() const override { return true; }

    bool canAddr() const override { return flags & FlagAddr; }

    bool canSet() const override { return !(flags & FlagRO); }

    bool canInterface() const override { return true; }

    bool isZero() const override { return value == T{}; }

    Value addr() const override;
    Value elem() const override;
    Value field(int i) const override;
    Value fieldByName(const std::string &name) const override;
    Value index(int i) const override;
    Value mapIndex(const Value &key) const override;
    Value call(const std::vector<Value> &args) const override;

    bool        boolValue() const override;
    int64_t     intValue() const override;
    uint64_t    uintValue() const override;
    double      floatValue() const override;
    std::string stringValue() const override;
    void       *pointerValue() const override;

    void setBool(bool v) override;
    void setInt(int64_t v) override;
    void setUint(uint64_t v) override;
    void setFloat(double v) override;
    void setString(const std::string &v) override;
    void setPointer(void *v) override;

    std::any interfaceValue() const override { return value; }

    std::string string() const override;
};

struct StructValue : public Value {
    void             *data;
    const StructType *structType;
    int               flags;

    StructValue(void *ptr, const StructType *type, int flags = FlagAddr) : data(ptr), structType(type), flags(flags) {}

    const Type &type() const override { return *structType; }

    bool isValid() const override { return data != nullptr; }

    bool canAddr() const override { return true; }

    bool canSet() const override { return !(flags & FlagRO); }

    bool canInterface() const override { return true; }

    bool isZero() const override;

    Value addr() const override;
    Value elem() const override;
    Value field(int i) const override;
    Value fieldByName(const std::string &name) const override;
    Value index(int i) const override;
    Value mapIndex(const Value &key) const override;
    Value call(const std::vector<Value> &args) const override;

    bool        boolValue() const override;
    int64_t     intValue() const override;
    uint64_t    uintValue() const override;
    double      floatValue() const override;
    std::string stringValue() const override;

    void *pointerValue() const override { return data; }

    void setBool(bool v) override;
    void setInt(int64_t v) override;
    void setUint(uint64_t v) override;
    void setFloat(double v) override;
    void setString(const std::string &v) override;

    void setPointer(void *v) override { data = v; }

    std::any interfaceValue() const override;

    std::string string() const override;
};

template<typename T>
const Type *TypeOf() {
    static BasicType<T> type;
    return &type;
}

template<typename T>
const Type *TypeOf(const T &value) {
    return TypeOf<T>();
}

template<typename T>
Value ValueOf(T value) {
    return BasicValue<T>(std::move(value), TypeOf<T>(), FlagAddr);
}

template<typename T>
Value ValueOf(T *ptr) {
    return BasicValue<T *>(ptr, TypeOf<T *>(), FlagAddr);
}

Value ZeroOf(const Type &type);

Value New(const Type &type);

Value Convert(const Value &v, const Type &t);

Value TypeAssert(const Value &v, const Type &t);

bool DeepEqual(const Value &a, const Value &b);

struct Person {
    std::string name;
    int         age;
    double      salary;

    void promote(double raise) {
        salary += raise;
        std::cout << name << " promoted! New salary: " << salary << std::endl;
    }

    std::string toString() const { return name + " (" + std::to_string(age) + ", $" + std::to_string(salary) + ")"; }
};

const Type *RegisterPersonType() {
    static StructType personType("Person");
    static bool       registered = false;

    if (!registered) {
        personType.addField("name", TypeOf<std::string>());
        personType.addField("age", TypeOf<int>());
        personType.addField("salary", TypeOf<double>());
        registered = true;
    }

    return &personType;
}

void testReflection() {
    auto intType = TypeOf<int>();
    std::cout << "int type: " << intType->name() << ", size: " << intType->size() << ", kind: " << static_cast<int>(intType->kind()) << std::endl;

    Value intVal = ValueOf(42);
    std::cout << "int value: " << intVal.intValue() << ", string: " << intVal.string() << std::endl;

    if (intVal.canSet()) {
        intVal.setInt(100);
        std::cout << "modified int value: " << intVal.intValue() << std::endl;
    }

    const Type *personType = RegisterPersonType();
    std::cout << "\nPerson type: " << personType->string() << std::endl;
    std::cout << "Fields:\n";
    for (int i = 0; i < personType->numField(); i++) {
        std::cout << "  " << personType->fieldName(i) << ": " << personType->field(i)->name() << ", offset: " << personType->fieldOffset(i) << std::endl;
    }

    Person alice{"Alice", 30, 85000.0};
    Value  personVal = ValueOf(&alice);// 获取指针

    Value nameField = personVal.fieldByName("name");
    Value ageField = personVal.fieldByName("age");
    std::cout << "\nAlice name: " << nameField.stringValue() << ", age: " << ageField.intValue() << std::endl;

    if (ageField.canSet()) {
        ageField.setInt(31);
        std::cout << "Alice new age: " << alice.age << std::endl;
    }

    std::cout << "\nCalling promote method:" << std::endl;
    Value promoteMethod = personVal.fieldByName("promote");// 假设有方法字段
    if (promoteMethod.isValid()) {
        std::vector<Value> args = {ValueOf(5000.0)};
        promoteMethod.call(args);
        std::cout << "New salary: " << alice.salary << std::endl;
    }

    std::vector<int> numbers = {1, 2, 3, 4, 5};
    Value            sliceVal = ValueOf(&numbers);
    std::cout << "\nSlice type: " << sliceVal.type().string() << std::endl;

    Value firstElement = sliceVal.index(0);
    std::cout << "First element: " << firstElement.intValue() << std::endl;

    Value floatVal = ValueOf(3.14f);
    Value intConverted = Convert(floatVal, *TypeOf<int>());
    std::cout << "\nConverted float to int: " << intConverted.intValue() << std::endl;

    Person bob1{"Bob", 25, 75000.0};
    Person bob2{"Bob", 25, 75000.0};
    Value  bobVal1 = ValueOf(&bob1);
    Value  bobVal2 = ValueOf(&bob2);

    bool equal = DeepEqual(bobVal1, bobVal2);
    std::cout << "\nBob1 and Bob2 are " << (equal ? "equal" : "not equal") << std::endl;

    Value newPerson = New(*personType);
    if (newPerson.isValid()) {
        newPerson.fieldByName("name").setString("Charlie");
        newPerson.fieldByName("age").setInt(35);
        newPerson.fieldByName("salary").setFloat(95000.0);

        Person *charlie = static_cast<Person *>(newPerson.pointerValue());
        std::cout << "\nNew person created: " << charlie->toString() << std::endl;
        delete charlie;
    }
}

/*
预期输出:

int type: int, size: 4, kind: 1
int value: 42, string: 42
modified int value: 100

Person type: struct Person
Fields:
  name: string, offset: 0
  age: int, offset: 32
  salary: float64, offset: 36

Alice name: Alice, age: 30
Alice new age: 31

Calling promote method:
Alice promoted! New salary: 90000
New salary: 90000

Slice type: *[]int
First element: 1

Converted float to int: 3

Bob1 and
