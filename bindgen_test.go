package bindgen

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type Project struct {
	Name        string
	Path        string
	TestFile    string
	RemoteURL   string
	SkipPush    bool
}

var projects = []Project{
	{Name: "glfw", Path: "project/glfw/clone", TestFile: "generate_test.go", RemoteURL: "https://github.com/ddkwork/glfw.git"},
	{Name: "everything", Path: "project/everything", TestFile: "generate_test.go", RemoteURL: "https://github.com/ddkwork/everything.git"},
	{Name: "ARImpRec", Path: "project/ARImpRec", TestFile: "generate_test.go", RemoteURL: "https://github.com/ddkwork/ARImpRec.git"},
	{Name: "keystone", Path: "project/keystone", TestFile: "generate_test.go", RemoteURL: "https://github.com/ddkwork/keystone.git"},
	{Name: "WinDivert", Path: "project/WinDivert", TestFile: "generate_test.go", RemoteURL: "https://github.com/ddkwork/WinDivert.git"},
	{Name: "xed", Path: "project/xed", TestFile: "generate_test.go", RemoteURL: "https://github.com/ddkwork/xed.git"},
	{Name: "zydis", Path: "project/zydis", TestFile: "generate_test.go", RemoteURL: "https://github.com/ddkwork/zydis.git"},
	{Name: "hyperdbgsdk", Path: "project/hyperdbgsdk", TestFile: "generate_test.go", RemoteURL: "", SkipPush: true},
	{Name: "imgui", Path: "project/imgui/gen", TestFile: "project_test.go", RemoteURL: "https://github.com/ddkwork/imgui.git"},
	{Name: "qt6", Path: "project/qt6/gen", TestFile: "project_test.go", RemoteURL: "https://github.com/ddkwork/qt6.git"},
}

func TestAllBindings(t *testing.T) {
	rootDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}

	results := make(map[string]bool)
	var failed []string

	for _, p := range projects {
		t.Run(p.Name, func(t *testing.T) {
			projectDir := filepath.Join(rootDir, p.Path)
			
			if _, err := os.Stat(projectDir); os.IsNotExist(err) {
				t.Skipf("Project directory not found: %s", projectDir)
				return
			}

			start := time.Now()
			cmd := exec.Command("go", "test", "-v", "-run", "TestGenerate", "-timeout", "40m", ".")
			cmd.Dir = projectDir
			output, err := cmd.CombinedOutput()
			duration := time.Since(start)

			if err != nil {
				results[p.Name] = false
				failed = append(failed, p.Name)
				t.Errorf("Test failed (%v):\n%s", duration, string(output))
				return
			}

			results[p.Name] = true
			t.Logf("PASS (%v)", duration)

			if !p.SkipPush && p.RemoteURL != "" {
				if err := pushProject(projectDir); err != nil {
					t.Logf("Push failed: %v", err)
				} else {
					t.Logf("Pushed to %s", p.RemoteURL)
				}
			}
		})
	}

	t.Log("\n=== Summary ===")
	for _, p := range projects {
		status := "FAIL"
		if results[p.Name] {
			status = "PASS"
		}
		t.Logf("%s: %s", p.Name, status)
	}

	if len(failed) > 0 {
		t.Errorf("Failed projects: %v", failed)
	}
}

func pushProject(dir string) error {
	cmd := exec.Command("git", "status", "--porcelain")
	cmd.Dir = dir
	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("git status: %v", err)
	}

	if len(strings.TrimSpace(string(output))) == 0 {
		return nil
	}

	commands := [][]string{
		{"git", "add", "-A"},
		{"git", "commit", "-m", "Update bindings\n\n🤖 Generated with CodeMate"},
		{"git", "push"},
	}

	for _, args := range commands {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			return fmt.Errorf("%s: %v\n%s", strings.Join(args, " "), err, string(out))
		}
	}

	return nil
}
