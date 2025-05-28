package mcp

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ddkwork/golibrary/stream"

	"mcp/bindgen"
)

const (
	newSdk    = "include"
	targetDir = "sdk"
)

func TestName(t *testing.T) {
	//fakeError.Walk(".")
	//bind(targetDir, func(s string) bool {
	//	return false
	//},"2.h")

	// os.RemoveAll("cache")
	RepackSdk()
	bindgen.Bind(targetDir, func(s string) bool {
		return filepath.Base(s) == "bridgemain.h"
	}, filepath.Join(newSdk, "bridgemain.h"))
}

func TestWalk(t *testing.T) {
	// os.RemoveAll("cache")
	RepackSdk()
	bindgen.Walk(
		newSdk,
		targetDir,
		func(s string) bool {
			for k := range map[string]bool{
				//"_plugins.h": true,
			} {
				if filepath.Base(s) == k {
					return false
				}
			}
			return stream.DirDepth(s) == 1
		},
		func(s string) bool {
			return filepath.Base(s) == "bridgemain.h"
		},
	)
}

const pluginsdk = "todo/snapshot_2025-03-15_15-57/pluginsdk"

func RepackSdk() { // 官方的sdk对于clang而言有瑕疵，需要patch
	filepath.Walk(pluginsdk, func(path string, info fs.FileInfo, err error) error {
		if strings.HasPrefix(filepath.Base(path), "_scriptapi_") {
			stream.CopyFile(path, filepath.Join(newSdk, filepath.Base(path)))
		}
		return err
	})
	stream.CopyFile(filepath.Join(pluginsdk, "bridgelist.h"), filepath.Join(newSdk, "bridgelist.h"))
	stream.CopyFile(filepath.Join(pluginsdk, "bridgegraph.h"), filepath.Join(newSdk, "bridgegraph.h"))
	stream.WriteTruncate(filepath.Join(newSdk, "_scriptapi.h"), `
#include "bridgelist.h"

//Debugger defines
#define MAX_PATH 256
#define MAX_LABEL_SIZE 256
#define MAX_COMMENT_SIZE 512
#define MAX_MODULE_SIZE 256
#define MAX_IMPORT_SIZE 65536
#define MAX_BREAKPOINT_SIZE 256
#define MAX_CONDITIONAL_EXPR_SIZE 256
#define MAX_CONDITIONAL_TEXT_SIZE 256
#define MAX_SCRIPT_LINE_SIZE 2048
#define MAX_THREAD_NAME_SIZE 256
#define MAX_WATCH_NAME_SIZE 256
#define MAX_STRING_SIZE 512
#define MAX_ERROR_SIZE 512
#define RIGHTS_STRING_SIZE (sizeof("ERWCG"))
#define MAX_SECTION_SIZE 10
#define MAX_COMMAND_LINE_SIZE 256
#define MAX_MNEMONIC_SIZE 64

#ifndef PAGE_SIZE
#define PAGE_SIZE 0x1000
#endif // PAGE_SIZE

#ifndef PLUG_IMPEXP
#ifdef BUILD_DBG
#define PLUG_IMPEXP __declspec(dllexport)
#else
#define PLUG_IMPEXP __declspec(dllimport)
#endif //BUILD_DBG
#endif //PLUG_IMPEXP

#define SCRIPT_EXPORT PLUG_IMPEXP

#ifdef _WIN64
typedef unsigned long long duint;
typedef signed long long dsint;
#else
typedef unsigned long duint;
typedef signed long dsint;
#endif //_WIN64

#ifndef BRIDGE_IMPEXP
#ifdef BUILD_BRIDGE
#define BRIDGE_IMPEXP __declspec(dllexport)
#else
#define BRIDGE_IMPEXP __declspec(dllimport)
#endif //BUILD_BRIDGE
#endif //BRIDGE_IMPEXP

`)

	g := stream.NewGeneratedFile()
	for s := range stream.ReadFileToLines(filepath.Join(pluginsdk, "bridgemain.h")) {
		if strings.HasPrefix(s, "//") {
			continue
		}
		g.P(s)
	}
	stream.WriteTruncate(filepath.Join(newSdk, "bridgemain.h"), g.String())
}
