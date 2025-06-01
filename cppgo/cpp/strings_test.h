#include <iostream>
#include "strings.h"

int main() {
    Buffer str = Buffer("GoStringPackage");

    std::cout << "Contains 'String': " << Strings::Contains(str, Buffer("String")) << std::endl;

    std::cout << "Has prefix 'Go': " << Strings::HasPrefix(str, Buffer("Go")) << std::endl;

    std::cout << "Index of 'P': " << Strings::Index(str, Buffer("P")) << std::endl;

    Buffer lower = Strings::ToLower(str);
    std::cout << "Lowercase: " << lower.String() << std::endl;

    Buffer upper = Strings::ToUpper(str);
    std::cout << "Uppercase: " << upper.String() << std::endl;

    Buffer replaced = Strings::ReplaceAll(str, Buffer("Package"), Buffer("PackageC++"));
    std::cout << "Replaced: " << replaced.String() << std::endl;

    Buffer spaceStr = Buffer("   Hello World!   \n");
    Buffer trimmed = Strings::TrimSpace(spaceStr);
    std::cout << "Trimmed: '" << trimmed.String() << "'" << std::endl;

    Buffer              csv = Buffer("apple,banana,orange");
    std::vector<Buffer> fruits = Strings::Split(csv, Buffer(","));
    std::cout << "Fruits split: ";
    for (const auto &fruit: fruits) { std::cout << fruit.String() << " "; }
    std::cout << std::endl;

    std::vector<Buffer> languages = {Buffer("Go"), Buffer("C++"), Buffer("Rust"), Buffer("Python")};
    Buffer              joined = Strings::Join(languages, Buffer(" | "));
    std::cout << "Joined: " << joined.String() << std::endl;

    Buffer              text = Buffer("   This is   a test   of fields  ");
    std::vector<Buffer> words = Strings::Fields(text);
    std::cout << "Words: ";
    for (const auto &word: words) { std::cout << word.String() << " "; }
    std::cout << std::endl;

    std::cout << "Count of 'is': " << Strings::Count(str, Buffer("is")) << std::endl;

    Buffer prefixed = Strings::TrimPrefix(str, Buffer("Go"));
    std::cout << "After trim prefix: " << prefixed.String() << std::endl;

    return 0;
}
