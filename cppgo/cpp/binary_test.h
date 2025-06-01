#include "binary.h"
#include <iostream>

int main() {
    Buffer b;

    Binary::WriteUint16(b, Binary::LittleEndian, 0x1234);
    Binary::WriteUint32(b, Binary::LittleEndian, 0x12345678);
    Binary::WriteFloat(b, Binary::LittleEndian, 3.14159f);

    Binary::WriteUvarint(b, 128); // 需要2字节
    Binary::WriteVarint(b, -1024);// 需要2字节

    std::cout << "Buffer size: " << b.Len() << " bytes" << std::endl;

    size_t   offset = 0;
    uint16_t u16 = Binary::ReadUint16(b, Binary::LittleEndian, offset);
    offset += Binary::Size<uint16_t>();
    std::cout << "u16: 0x" << std::hex << u16 << std::dec << std::endl;

    uint32_t u32 = Binary::ReadUint32(b, Binary::LittleEndian, offset);
    offset += Binary::Size<uint32_t>();
    std::cout << "u32: 0x" << std::hex << u32 << std::dec << std::endl;

    float f = Binary::ReadFloat(b, Binary::LittleEndian, offset);
    offset += Binary::Size<float>();
    std::cout << "float: " << f << std::endl;

    uint64_t uvarint = Binary::ReadUvarint(b, &offset);
    std::cout << "uvarint: " << uvarint << std::endl;

    int64_t varint = Binary::ReadVarint(b, &offset);
    std::cout << "varint: " << varint << std::endl;

    std::cout << "Size of 127: " << Binary::Varint::Size(127) << std::endl;
    std::cout << "Size of 128: " << Binary::Varint::Size(128) << std::endl;
    std::cout << "Size of -1: " << Binary::Varint::Size(-1) << std::endl;

    std::cout << "Native byte order: " << (Binary::NativeOrder() == Binary::LittleEndian ? "LittleEndian" : "BigEndian") << std::endl;

    return 0;
}
