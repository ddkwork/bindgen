# bindgen

C header to Go binding generator, powered by [modernc.org/cc/v4](https://pkg.go.dev/modernc.org/cc/v4).

## Overview

bindgen parses C headers using a real C compiler frontend and generates idiomatic Go bindings — structs, enums, typedefs, macro constants, and DLL function wrappers — all from a single `go test` invocation.

## Features

- **Full C parsing** — uses modernc.org/cc/v4 (no CGO required), supports MSVC-specific types and EWDK include paths
- **Struct generation** — C structs → Go structs with correct field alignment, including bitfield getter/setter methods
- **Enum generation** — C enums → Go `const` blocks with proper type aliases
- **Typedef resolution** — C typedefs → Go type aliases, including function pointer types
- **Macro constant extraction** — `#define` constants → Go `const`, with arithmetic expression evaluation
- **DLL binding** — generates `syscall.LoadDLL`-based wrappers with configurable function name filtering
- **Single-file or multi-file output** — per-project configuration
- **`byteslice` package** — generic `byte`/`int8` ↔ `string` conversion utilities from shared library (`github.com/ddkwork/golibrary/byteslice`)

## Prerequisites

- Go 1.26.2+
- Windows (MSVC toolchain)
- EWDK or Visual Studio with `cl.exe` on PATH

## Usage

Bindings are generated via `go test`:

```bash
go test -v -run TestGenerate -count=1 -timeout 300s .
```

This processes all project configurations defined in `generate_test.go` and writes Go source files to each project's output directory.

## Supported Projects

| Project | Package | DLL Binding | Description |
|---|---|---|---|
| HyperDbg SDK | `sdk` | — | HyperDbg debugger SDK structures and constants |
| Zydis | `zydis` | `zydis.dll` | Fast x86/x86-64 disassembler |
| Keystone | `keystone` | `keystone.dll` | Multi-architecture assembler engine |
| Intel XED | `xed` | `xed.dll` | x86 encoder/decoder |
| GLFW | `glfw` | `glfw3.dll` | Window and input library |
| WinDivert | `windivert` | `WinDivert64.dll` | Network packet interception |
| Everything SDK | `everything` | `Everything64.dll` | File search engine SDK |
| ARImpRec | `ipmrec` | `ARImpRec.dll` | Import reconstruction library |

## Project Structure

```
bindgen/
├── generate_test.go    # Binding generation logic and project configs
├── parse_test.go       # C parsing test helpers
├── msenv.go            # EWDK/MSVC environment detection
├── cc/v4/              # Vendored modernc.org/cc/v4
└── project/            # Generated bindings output
    ├── hyperdbg/sdk/
    ├── zydis/
    ├── keystone/
    ├── xed/
    ├── glfw/
    ├── WinDivert/
    ├── Everything-SDK/
    └── ARImpRec/
```

## Adding a New Binding

Add a `BindgenConfig` entry in `generate_test.go`'s `TestGenerate`:

```go
t.Run("mylib", func(t *testing.T) {
    Generate(t, []BindgenConfig{{
        HeadersDir:   "project/mylib/clone/include",
        OutputDir:    "project/mylib",
        PackageName:  "mylib",
        HeaderOrder:  []string{"mylib.h"},
        BindDll:      true,
        DllName:      "mylib.dll",
        DllFuncFilter: func(name string) bool {
            return strings.HasPrefix(name, "mylib_")
        },
    }})
})
```

## `byteslice` Package

Generic `byte`/`int8` ↔ `string` conversion utilities for C interop, from shared library [`github.com/ddkwork/golibrary/byteslice`](https://pkg.go.dev/github.com/ddkwork/golibrary/byteslice):

```go
import "github.com/ddkwork/golibrary/byteslice"

// []byte/[]int8 → string (truncates at NUL)
s := byteslice.ToString[byte]([]byte("hello\x00world"))  // "hello"
s := byteslice.ToString[int8]([]int8{104, 101, 108, 108, 111, 0})  // "hello"

// *byte/*int8 → string (reads until NUL)
s := byteslice.PtrToString[byte](p)
s := byteslice.PtrToString[int8](p)

// string → []byte/[]int8 (appends NUL terminator)
slice := byteslice.FromString[byte]("hello")
slice := byteslice.FromString[int8]("hello")

// string → *byte/*int8 (appends NUL terminator)
ptr := byteslice.PtrFromString[byte]("hello")
ptr := byteslice.PtrFromString[int8]("hello")

// []byte ↔ struct (zero-copy)
data := byteslice.FromStruct(&myStruct)
ptr := byteslice.ToStruct[MyStruct](buf)

// []T → []byte (zero-copy)
raw := byteslice.FromSlice(mySlice)
```

## CI

GitHub Actions workflow (`.github/workflows/go.yml`) runs on Windows:

1. Mount EWDK ISO for MSVC headers
2. Run `go test -run TestGenerate` to generate all bindings
3. Verify generated code compiles with `go vet`
4. Run per-project tests (DLL tests require corresponding `.dll` files)

## License

This project uses [modernc.org/cc/v4](https://modernc.org/cc/) which is licensed under the BSD-3-Clause.
