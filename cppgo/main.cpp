#include "driver.h"

void DemoBuffer() {
    Buffer buf = Buffer::New();
    buf.WriteString("简洁Buffer实现!");
    Print("长度: %zu, 容量: %zu\n", buf.Len(), buf.Cap());
    int number = 0x12345678;
    buf.Write(reinterpret_cast<const unsigned char *>(&number), sizeof(number));
    Print("写入后长度: %zu\n", buf.Len());
    unsigned char readData[32];
    size_t        n = buf.Read(readData, sizeof(readData));
    Print("读取 %zu 字节\n", n);
    char *str = buf.ToString();
    if (str) {
        Println("内容: %s", str);
        Free(str);
    }
    Buffer slice = buf.Slice(0, 6);
    char  *sliceStr = slice.ToString();
    if (sliceStr) {
        Println("切片: %s", sliceStr);
        Free(sliceStr);
    }
    Buffer other = Buffer::FromString("简洁Buffer实现!");
    if (buf.Equals(other)) { Println("内容相同"); }
    buf.Reset();
    slice.Reset();
    other.Reset();
}

#ifdef _KERNEL_MODE
extern "C" NTSTATUS DriverEntry(PDRIVER_OBJECT DriverObject, PUNICODE_STRING RegistryPath) {
    Driver driver = {};
    return driver.Run(DriverObject, RegistryPath);
}
#else
#include <windows.h>

int main() {
    auto hDevice = CreateFile(L"\\\\.\\BufferDriver", GENERIC_READ | GENERIC_WRITE, 0, NULL, OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL, NULL);
    if (hDevice == INVALID_HANDLE_VALUE) {
        printf("无法打开设备: %d\n", GetLastError());
        return 1;
    }
    char  writeCmd[] = "WHello";// 'W'命令 + 数据
    DWORD bytesReturned = 0;

    if (!DeviceIoControl(hDevice, CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS), writeCmd, sizeof(writeCmd), NULL, 0, &bytesReturned, NULL)) {
        printf("写入失败: %d\n", GetLastError());
    } else {
        printf("写入成功\n");
    }

    char readBuffer[64] = {0};
    char readCmd = 'R';// 'R'命令

    if (!DeviceIoControl(hDevice,
                         CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS),
                         &readCmd,
                         sizeof(readCmd),
                         readBuffer,
                         sizeof(readBuffer),
                         &bytesReturned,
                         NULL)) {
        printf("读取失败: %d\n", GetLastError());
    } else {
        int bytesRead = static_cast<int>(readBuffer[0]);
        printf("读取 %d 字节: %.*s\n", bytesRead, bytesRead, readBuffer + 1);
    }

    CloseHandle(hDevice);
    return 0;
}

#endif
