#include "hex.h"
#include <iostream>

int main() {
    Buffer data = {0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x2C, 0x20, 0x57, 0x6F, 0x72, 0x6C, 0x64, 0x21};

    Buffer hexEncoded = Hex::Encode(data);
    std::cout << "Encoded: " << hexEncoded.String() << std::endl;

    Buffer decoded = Hex::Decode(hexEncoded);
    std::cout << "Decoded: " << decoded.String() << std::endl;

    std::string hexStr = Hex::EncodeToString(data);
    std::cout << "String Encoded: " << hexStr << std::endl;

    Buffer fromStr = Hex::DecodeString(hexStr);
    std::cout << "String Decoded: " << fromStr.String() << std::endl;

    std::cout << "\nHex Dump:\n";
    std::cout << Hex::Dump(data);
    /*
    输出:
    00000000  48 65 6c 6c 6f 2c 20 57  6f 72 6c 64 21              |Hello, World!|
    0000000d
    */

    uint8_t binaryData[] = {0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
                            0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F};

    std::cout << "\nBinary Data Dump:\n";
    std::cout << Hex::Dump(binaryData, sizeof(binaryData));

    try {
        Hex::Decode("12 34 56");// 包含空格
    } catch (const Hex::InvalidByteError &e) { std::cerr << "Error: " << e.what() << std::endl; }

    try {
        Hex::Decode("123");// 奇数长度
    } catch (const Hex::LengthError &e) { std::cerr << "Error: " << e.what() << std::endl; }

    return 0;
}
