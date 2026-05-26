package cpp2c

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

var clangTimeout = 10 * time.Minute

func clangExec(inputHeader string, cflags []string) (*AstNode, error) {
	ctx, cancel := context.WithTimeout(context.Background(), clangTimeout)
	defer cancel()

	clangArgs := []string{`-x`, `c++`}
	clangArgs = append(clangArgs, cflags...)
	clangArgs = append(clangArgs, `-Xclang`, `-ast-dump=json`, `-fsyntax-only`, inputHeader)

	cmd := exec.CommandContext(ctx, "clang", clangArgs...)

	var outBuf bytes.Buffer
	var stderrBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &stderrBuf

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("could not start clang: %w", err)
	}

	if err := cmd.Wait(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return nil, fmt.Errorf("clang timed out after %v for %s", clangTimeout, inputHeader)
		}
		if stderrBuf.Len() > 0 {
			return nil, fmt.Errorf("could not wait for clang: %w\nstderr:\n%s", err, stderrBuf.String())
		}
		return nil, fmt.Errorf("could not wait for clang: %w", err)
	}

	raw := outBuf.Bytes()
	if len(raw) == 0 {
		return nil, fmt.Errorf("clang produced no output for %s", inputHeader)
	}

	writeAstCache(raw, inputHeader)
	return filterAst(bytes.NewReader(raw), isQtProjectFromCflags(cflags, inputHeader))
}

func astPath(inputHeader string) string {
	return CacheFileRoot(inputHeader) + ".json"
}

func writeAstCache(data []byte, inputHeader string) {
	path := astPath(inputHeader)
	os.MkdirAll(filepath.Dir(path), 0755)
	err := os.WriteFile(path, data, 0644)
	if err != nil {
		panic("could not write AST cache for " + inputHeader + ": " + err.Error())
	}
}

func readAstCache(inputHeader string) ([]byte, error) {
	return os.ReadFile(astPath(inputHeader))
}
