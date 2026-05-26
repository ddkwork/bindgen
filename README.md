# bindgen - C/C++ to Go Bindings Generator

**纯 Go 实现，无 CGO 依赖** - 将任意 C/C++ 项目自动生成为 Go 语言绑定。

## 愿景

让任何 C/C++ 库都能快速、无痛地获得 Go 语言绑定：

- 🚀 **无 CGO** - 纯 Go 实现，使用 `golang.org/x/sys/windows` 动态加载 DLL
- ⚡ **编译快** - 无 CGO 开销，编译速度极快
- 🔄 **交叉编译** - 轻松支持跨平台编译
- 📦 **DLL 嵌入** - 动态库嵌入到 Go 二进制，无需手动分发
- 🛠️ **自动生成** - 从 C/C++ 头文件自动生成所有绑定代码

## 绑定的项目

| 项目 | 描述 | 仓库 |
|------|------|------|
| **GLFW** | 跨平台窗口和输入库 | [github.com/ddkwork/glfw](https://github.com/ddkwork/glfw) |
| **Everything SDK** | Windows 文件搜索 | [github.com/ddkwork/everything](https://github.com/ddkwork/everything) |
| **Intel XED** | x86 指令编码/解码 | [github.com/ddkwork/xed](https://github.com/ddkwork/xed) |
| **Zydis** | x86/x64 反汇编器 | [github.com/ddkwork/zydis](https://github.com/ddkwork/zydis) |
| **ImGui** | 即时模式 GUI | [github.com/ddkwork/imgui](https://github.com/ddkwork/imgui) |
| **Qt6** | 跨平台 GUI 框架 | [github.com/ddkwork/qt6](https://github.com/ddkwork/qt6) |
| **Keystone** | 多架构汇编引擎 | [github.com/ddkwork/keystone](https://github.com/ddkwork/keystone) |
| **WinDivert** | Windows 网络包拦截 | [github.com/ddkwork/WinDivert](https://github.com/ddkwork/WinDivert) |
| **ARImpRec** | PE 导入表重建 | [github.com/ddkwork/ARImpRec](https://github.com/ddkwork/ARImpRec) |

## 架构

```
bindgen/
├── c2go/              # C 到 Go 绑定生成核心
│   ├── generate.go          # 主生成逻辑
│   ├── fake_headers.go      # 虚拟 C 标准库头文件
│   ├── msenv.go             # MSVC 编译器环境
│   └── types.go             # 类型定义
├── cc/                # C 预处理和解析器
├── cpp2c/             # C++ 到 C 转换
├── bindgen_test.go    # 集成测试（测试所有项目）
└── project/           # 绑定项目示例
    ├── glfw/
    ├── xed/
    ├── zydis/
    └── ...
```

## 使用方法

### 1. 创建配置文件

```go
// generate_test.go
package mylib

import (
    "testing"
    "github.com/ddkwork/bindgen/c2go"
)

func TestGenerate(t *testing.T) {
    c2go.Generate(t, []c2go.BindgenConfig{{
        HeadersDir:     "include",           // C 头文件目录
        OutputDir:      ".",                 // Go 输出目录
        PackageName:    "mylib",             // Go 包名
        HeaderOrder:    []string{"mylib.h"}, // 处理顺序
        BindDll:        true,                // 生成 DLL 绑定
        DllName:        "mylib.dll",         // DLL 名称
        DllFuncFilter: func(name string) bool {
            return true // 过滤导出函数
        },
    }})
}
```

### 2. 运行生成

```bash
go test -run TestGenerate -v
```

### 3. 使用绑定

```go
package main

import "github.com/ddkwork/mylib"

func main() {
    lib := &mylib.Mylib{}
    lib.Init()
    defer lib.Destroy()
    
    // 使用生成的函数...
}
```

## 核心特性

### Fake Headers

内置虚拟 C 标准库头文件，无需依赖系统头文件：

- `stddef.h`, `stdint.h` - 标准类型（size_t, int32_t 等）
- `stdlib.h`, `stdio.h`, `string.h` - 标准函数
- `windows.h` - Windows API 类型（HWND, DWORD 等）

### 类型映射

自动将 C 类型映射为 Go 类型：

| C 类型 | Go 类型 |
|--------|---------|
| `int32_t` | `int32` |
| `uint64_t` | `uint64` |
| `size_t` | `uintptr` |
| `void*` | `unsafe.Pointer` |
| `char*` | `*int8` |
| 函数指针 | `func(...) uintptr` |

### DLL 嵌入

自动嵌入 DLL 到 Go 二进制：

```go
//go:embed mylib.dll
var dllBytes []byte
```

运行时解压到用户缓存目录，无需手动分发 DLL。

## 运行所有测试

```bash
go test -v -run TestAllBindings
```

自动测试所有绑定项目，成功的项目自动推送到 GitHub。

## 依赖

- Go 1.26+
- `golang.org/x/sys/windows` - Windows API 调用
- MSVC cl.exe (可选，用于预处理)

## 许可证

MIT License

## 贡献

欢迎提交 Issue 和 Pull Request！
