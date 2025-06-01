#include <iostream>
#include "buffer.h"

int main() {
    Buffer buf;
    buf.WriteString("Hello, ");
    buf.WriteBytes(reinterpret_cast<const uint8_t *>("Buffer!"), 7);

    std::cout << "Buffer: " << buf.String() << std::endl;
    std::cout << "Length: " << buf.Len() << std::endl;

    Buffer firstWord = buf.Next(5);
    std::cout << "First word: " << firstWord.String() << std::endl;

    try {
        buf.UnreadByte();
        buf.UnreadByte();
        std::cout << "After unread: " << buf.String() << std::endl;
    } catch (const std::exception &e) { std::cerr << "Error: " << e.what() << std::endl; }

    Buffer prefix("Hello");
    std::cout << "Has prefix 'Hello': " << buf.HasPrefix(prefix) << std::endl;

    Buffer needle("Buffer");
    std::cout << "Contains 'Buffer': " << buf.Contains(needle) << std::endl;
    std::cout << "Index of 'Buffer': " << buf.Index(needle) << std::endl;

    std::vector<Buffer> parts = {Buffer("Go"), Buffer("C++"), Buffer("Rust")};

    Buffer sep(", ");
    Buffer joined = Buffer::Join(parts, sep);
    std::cout << "Joined: " << joined.String() << std::endl;

    Buffer repeated = Buffer::Repeat(Buffer("*"), 10);
    std::cout << "Repeated: " << repeated.String() << std::endl;

    {
        std::ofstream out("test.txt");
        buf.WriteTo(out);
    }

    Buffer fileBuf;
    {
        std::ifstream in("test.txt");
        fileBuf.ReadFrom(in);
    }
    std::cout << "File content: " << fileBuf.String() << std::endl;

    return 0;
}
