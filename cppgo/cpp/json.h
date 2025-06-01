#pragma once

#include "buffer.h"
#include <map>
#include <vector>
#include <string>
#include <sstream>
#include <iomanip>
#include <stdexcept>
#include <type_traits>

enum JsonType { Null, Boolean, Number, String, Array, Object };

struct JsonValue {
    JsonType type;

    union {
        bool                              boolValue;
        double                            numberValue;
        std::string                      *stringValue;
        std::vector<JsonValue>           *arrayValue;
        std::map<std::string, JsonValue> *objectValue;
    };

    JsonValue() : type(Null) {}

    explicit JsonValue(bool v) : type(Boolean), boolValue(v) {}

    explicit JsonValue(double v) : type(Number), numberValue(v) {}

    explicit JsonValue(const std::string &v) : type(String), stringValue(new std::string(v)) {}

    explicit JsonValue(const char *v) : type(String), stringValue(new std::string(v)) {}

    ~JsonValue() { free(); }

    JsonValue(const JsonValue &other) { copyFrom(other); }

    JsonValue &operator=(const JsonValue &other) {
        if (this != &other) {
            free();
            copyFrom(other);
        }
        return *this;
    }

    JsonValue(JsonValue &&other) noexcept { moveFrom(std::move(other)); }

    JsonValue &operator=(JsonValue &&other) noexcept {
        if (this != &other) {
            free();
            moveFrom(std::move(other));
        }
        return *this;
    }

    void addElement(JsonValue value) {
        if (type != Array) {
            free();
            type = Array;
            arrayValue = new std::vector<JsonValue>();
        }
        arrayValue->push_back(std::move(value));
    }

    void setProperty(const std::string &key, JsonValue value) {
        if (type != Object) {
            free();
            type = Object;
            objectValue = new std::map<std::string, JsonValue>();
        }
        (*objectValue)[key] = std::move(value);
    }

    std::string toString() const;

private:
    void free() {
        switch (type) {
            case String:
                delete stringValue;
                break;
            case Array:
                delete arrayValue;
                break;
            case Object:
                delete objectValue;
                break;
            default:
                break;
        }
        type = Null;
    }

    void copyFrom(const JsonValue &other) {
        type = other.type;
        switch (type) {
            case Boolean:
                boolValue = other.boolValue;
                break;
            case Number:
                numberValue = other.numberValue;
                break;
            case String:
                stringValue = new std::string(*other.stringValue);
                break;
            case Array:
                arrayValue = new std::vector<JsonValue>(*other.arrayValue);
                break;
            case Object:
                objectValue = new std::map<std::string, JsonValue>(*other.objectValue);
                break;
            default:
                break;
        }
    }

    void moveFrom(JsonValue &&other) {
        type = other.type;
        switch (type) {
            case Boolean:
                boolValue = other.boolValue;
                break;
            case Number:
                numberValue = other.numberValue;
                break;
            case String:
                stringValue = other.stringValue;
                other.stringValue = nullptr;
                break;
            case Array:
                arrayValue = other.arrayValue;
                other.arrayValue = nullptr;
                break;
            case Object:
                objectValue = other.objectValue;
                other.objectValue = nullptr;
                break;
            default:
                break;
        }
        other.type = Null;
    }
};

struct Json {
    static Buffer      Marshal(const JsonValue &value);
    static std::string MarshalString(const JsonValue &value);

    static JsonValue Unmarshal(const Buffer &json);
    static JsonValue UnmarshalString(const std::string &json);

private:
    static void serialize(std::ostringstream &oss, const JsonValue &value);
    static void serializeString(std::ostringstream &oss, const std::string &str);

    static JsonValue parseValue(const std::string &json, size_t &index);
    static JsonValue parseArray(const std::string &json, size_t &index);
    static JsonValue parseObject(const std::string &json, size_t &index);
    static JsonValue parseString(const std::string &json, size_t &index);
    static JsonValue parseNumber(const std::string &json, size_t &index);

    static void skipWhitespace(const std::string &json, size_t &index);
};

std::string JsonValue::toString() const {
    switch (type) {
        case Null:
            return "null";
        case Boolean:
            return boolValue ? "true" : "false";
        case Number:
            return std::to_string(numberValue);
        case String:
            return "\"" + *stringValue + "\"";
        case Array: {
            std::string s = "[";
            for (size_t i = 0; i < arrayValue->size(); i++) {
                if (i > 0) s += ", ";
                s += (*arrayValue)[i].toString();
            }
            return s + "]";
        }
        case Object: {
            std::string s = "{";
            size_t      i = 0;
            for (const auto &[key, value]: *objectValue) {
                if (i++ > 0) s += ", ";
                s += "\"" + key + "\":" + value.toString();
            }
            return s + "}";
        }
    }
    return "null";
}

Buffer Json::Marshal(const JsonValue &value) { return Buffer(Json::MarshalString(value)); }

std::string Json::MarshalString(const JsonValue &value) {
    std::ostringstream oss;
    Json::serialize(oss, value);
    return oss.str();
}

JsonValue Json::Unmarshal(const Buffer &json) { return Json::UnmarshalString(json.String()); }

JsonValue Json::UnmarshalString(const std::string &json) {
    size_t index = 0;
    skipWhitespace(json, index);
    return parseValue(json, index);
}

void Json::serialize(std::ostringstream &oss, const JsonValue &value) {
    switch (value.type) {
        case Null:
            oss << "null";
            break;
        case Boolean:
            oss << (value.boolValue ? "true" : "false");
            break;
        case Number:
            oss << value.numberValue;
            break;
        case String:
            serializeString(oss, *value.stringValue);
            break;
        case Array: {
            oss << "[";
            for (size_t i = 0; i < value.arrayValue->size(); i++) {
                if (i > 0) oss << ",";
                serialize(oss, (*value.arrayValue)[i]);
            }
            oss << "]";
            break;
        }
        case Object: {
            oss << "{";
            size_t i = 0;
            for (const auto &[key, val]: *value.objectValue) {
                if (i++ > 0) oss << ",";
                serializeString(oss, key);
                oss << ":";
                serialize(oss, val);
            }
            oss << "}";
            break;
        }
    }
}

void Json::serializeString(std::ostringstream &oss, const std::string &str) {
    oss << "\"";
    for (char c: str) {
        switch (c) {
            case '"':
                oss << "\\\"";
                break;
            case '\\':
                oss << "\\\\";
                break;
            case '\b':
                oss << "\\b";
                break;
            case '\f':
                oss << "\\f";
                break;
            case '\n':
                oss << "\\n";
                break;
            case '\r':
                oss << "\\r";
                break;
            case '\t':
                oss << "\\t";
                break;
            default:
                if (static_cast<unsigned char>(c) < 0x20) {
                    oss << "\\u" << std::hex << std::setw(4) << std::setfill('0') << static_cast<int>(c);
                } else {
                    oss << c;
                }
                break;
        }
    }
    oss << "\"";
}

void Json::skipWhitespace(const std::string &json, size_t &index) {
    while (index < json.size() && isspace(json[index])) { index++; }
}

JsonValue Json::parseValue(const std::string &json, size_t &index) {
    skipWhitespace(json, index);
    if (index >= json.size()) { throw std::runtime_error("Unexpected end of JSON"); }

    char c = json[index];
    if (c == '{') {
        return parseObject(json, index);
    } else if (c == '[') {
        return parseArray(json, index);
    } else if (c == '"') {
        return parseString(json, index);
    } else if (c == '-' || (c >= '0' && c <= '9')) {
        return parseNumber(json, index);
    } else if (c == 't' && json.substr(index, 4) == "true") {
        index += 4;
        return JsonValue(true);
    } else if (c == 'f' && json.substr(index, 5) == "false") {
        index += 5;
        return JsonValue(false);
    } else if (c == 'n' && json.substr(index, 4) == "null") {
        index += 4;
        return JsonValue();
    }
    throw std::runtime_error("Unexpected token: " + std::string(1, c));
}

JsonValue Json::parseObject(const std::string &json, size_t &index) {
    JsonValue result;
    result.type = Object;
    result.objectValue = new std::map<std::string, JsonValue>();

    index++;// 跳过 '{'
    skipWhitespace(json, index);

    if (index < json.size() && json[index] == '}') {
        index++;// 空对象
        return result;
    }

    while (index < json.size()) {
        skipWhitespace(json, index);
        if (json[index] != '"') { throw std::runtime_error("Expected string for object key"); }

        JsonValue keyValue = parseString(json, index);
        if (keyValue.type != String) { throw std::runtime_error("Expected string for object key"); }
        std::string key = *keyValue.stringValue;

        skipWhitespace(json, index);
        if (index >= json.size() || json[index] != ':') { throw std::runtime_error("Expected colon after key in object"); }
        index++;

        JsonValue value = parseValue(json, index);
        result.objectValue->insert({key, std::move(value)});

        skipWhitespace(json, index);
        if (index < json.size() && json[index] == ',') {
            index++;
            continue;
        } else if (index < json.size() && json[index] == '}') {
            index++;
            break;
        } else {
            throw std::runtime_error("Expected comma or closing brace in object");
        }
    }

    return result;
}

JsonValue Json::parseArray(const std::string &json, size_t &index) {
    JsonValue result;
    result.type = Array;
    result.arrayValue = new std::vector<JsonValue>();

    index++;// 跳过 '['
    skipWhitespace(json, index);

    if (index < json.size() && json[index] == ']') {
        index++;// 空数组
        return result;
    }

    while (index < json.size()) {
        JsonValue element = parseValue(json, index);
        result.arrayValue->push_back(std::move(element));

        skipWhitespace(json, index);
        if (index < json.size() && json[index] == ',') {
            index++;
            continue;
        } else if (index < json.size() && json[index] == ']') {
            index++;
            break;
        } else {
            throw std::runtime_error("Expected comma or closing bracket in array");
        }
    }

    return result;
}

JsonValue Json::parseString(const std::string &json, size_t &index) {
    index++;// 跳过开头的引号
    std::string value;
    bool        escape = false;

    while (index < json.size()) {
        char c = json[index++];

        if (escape) {
            escape = false;
            switch (c) {
                case '"':
                    value += '"';
                    break;
                case '\\':
                    value += '\\';
                    break;
                case '/':
                    value += '/';
                    break;
                case 'b':
                    value += '\b';
                    break;
                case 'f':
                    value += '\f';
                    break;
                case 'n':
                    value += '\n';
                    break;
                case 'r':
                    value += '\r';
                    break;
                case 't':
                    value += '\t';
                    break;
                case 'u': {
                    if (index + 4 > json.size()) { throw std::runtime_error("Invalid unicode escape"); }
                    index += 4;  // 跳过4个十六进制数字
                    value += '?';// 用问号替代
                    break;
                }
                default:
                    value += c;
                    break;
            }
        } else {
            if (c == '\\') {
                escape = true;
            } else if (c == '"') {
                break;// 结束字符串
            } else {
                value += c;
            }
        }
    }

    return JsonValue(value);
}

JsonValue Json::parseNumber(const std::string &json, size_t &index) {
    size_t start = index;
    bool   isDouble = false;

    if (json[index] == '-') index++;

    while (index < json.size() && isdigit(json[index])) { index++; }

    if (index < json.size() && json[index] == '.') {
        isDouble = true;
        index++;
        while (index < json.size() && isdigit(json[index])) { index++; }
    }

    if (index < json.size() && (json[index] == 'e' || json[index] == 'E')) {
        isDouble = true;
        index++;
        if (index < json.size() && (json[index] == '+' || json[index] == '-')) { index++; }
        while (index < json.size() && isdigit(json[index])) { index++; }
    }

    std::string numStr = json.substr(start, index - start);
    if (isDouble) {
        return JsonValue(std::stod(numStr));
    } else {
        return JsonValue(static_cast<double>(std::stoll(numStr)));
    }
}

#include "json.h"
#include <iostream>
#include <cassert>

int main() {
    JsonValue person;
    person.setProperty("name", JsonValue("Alice"));
    person.setProperty("age", JsonValue(30));
    person.setProperty("isStudent", JsonValue(false));

    JsonValue hobbies;
    hobbies.addElement(JsonValue("reading"));
    hobbies.addElement(JsonValue("swimming"));
    hobbies.addElement(JsonValue("hiking"));
    person.setProperty("hobbies", std::move(hobbies));

    JsonValue address;
    address.setProperty("street", JsonValue("123 Main St"));
    address.setProperty("city", JsonValue("New York"));
    address.setProperty("zip", JsonValue(10001));
    person.setProperty("address", std::move(address));

    std::string jsonStr = Json::MarshalString(person);
    std::cout << "Serialized JSON:\n" << jsonStr << "\n\n";

    JsonValue parsed = Json::UnmarshalString(jsonStr);

    assert(parsed.type == Object);
    assert(parsed.objectValue->at("name").type == String);
    assert(*parsed.objectValue->at("name").stringValue == "Alice");
    assert(parsed.objectValue->at("age").type == Number);
    assert(parsed.objectValue->at("age").numberValue == 30);
    assert(parsed.objectValue->at("isStudent").type == Boolean);
    assert(parsed.objectValue->at("isStudent").boolValue == false);

    const auto &addr = parsed.objectValue->at("address");
    assert(addr.objectValue->at("street").type == String);
    assert(*addr.objectValue->at("street").stringValue == "123 Main St");

    const auto &hobbiesArray = parsed.objectValue->at("hobbies");
    assert(hobbiesArray.type == Array);
    assert(hobbiesArray.arrayValue->size() == 3);
    assert(*hobbiesArray.arrayValue->at(0).stringValue == "reading");

    parsed.objectValue->at("age") = JsonValue(31);

    std::cout << "Modified age JSON:\n" << Json::MarshalString(parsed) << std::endl;

    return 0;
}
