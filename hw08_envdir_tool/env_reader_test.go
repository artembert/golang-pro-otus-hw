package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func dropTempFolder(t *testing.T, path string) {
	t.Helper()
	err := os.RemoveAll(path)
	if err != nil {
		t.Errorf("Error removing test folder '%v': %v", path, err)
	}
}

func prepareTestDir(t *testing.T) (string, string) {
	t.Helper()
	dir := filepath.Join("tmp", strings.ReplaceAll(t.Name(), "/", "_"))
	err := os.MkdirAll(dir, os.FileMode(0o700))
	if err != nil {
		t.Errorf("Error creating test directory: %v", err)
	}
	outFile := filepath.Join(dir, "out.txt")

	return dir, outFile
}

func TestReadDir(t *testing.T) {
	envsPath := "testdata/env"

	t.Run("Empty dir", func(t *testing.T) {
		dir, _ := prepareTestDir(t)
		defer dropTempFolder(t, dir)

		envs, _ := ReadDir(dir)

		if len(envs) != 0 {
			t.Errorf("Expected empty map, got '%v'", envs)
		}
	})

	t.Run("Handle first line only", func(t *testing.T) {
		envs, _ := ReadDir(envsPath)

		config := envs["BAR"]
		if config.Value != "bar" {
			t.Errorf("Expected 'bar', got '%v'", config.Value)
		}
		if config.NeedRemove != false {
			t.Errorf("Expected NeedRemove == 'false', got '%v'", config.NeedRemove)
		}
	})

	t.Run("Handle empty first line", func(t *testing.T) {
		envs, _ := ReadDir(envsPath)

		config := envs["EMPTY"]
		if config.Value != "" {
			t.Errorf("Expected '', got '%v'", config.Value)
		}
		if config.NeedRemove != false {
			t.Errorf("Expected NeedRemove == 'false', got '%v'", config.NeedRemove)
		}
	})

	t.Run("Replace null-terminated strings with \\n", func(t *testing.T) {
		envs, _ := ReadDir(envsPath)

		config := envs["FOO"]
		if config.Value != "   foo\nwith new line" {
			t.Errorf("Expected '   foo\nwith new line', got '%v'", config.Value)
		}
		if config.NeedRemove != false {
			t.Errorf("Expected NeedRemove == 'false', got '%v'", config.NeedRemove)
		}
	})

	t.Run("Preserve quotes", func(t *testing.T) {
		envs, _ := ReadDir(envsPath)

		config := envs["HELLO"]
		if config.Value != "\"hello\"" {
			t.Errorf("Expected '\"hello\"', got '%v'", config.Value)
		}
		if config.NeedRemove != false {
			t.Errorf("Expected NeedRemove == 'false', got '%v'", config.NeedRemove)
		}
	})

	t.Run("Mark variable as NeedRemove if a file of a zero length", func(t *testing.T) {
		envs, _ := ReadDir(envsPath)

		config := envs["UNSET"]
		if config.Value != "" {
			t.Errorf("Expected '', got '%v'", config.Value)
		}
		if config.NeedRemove != true {
			t.Errorf("Expected NeedRemove == 'true', got '%v'", config.NeedRemove)
		}
	})
}
