package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func dropTempFolder(t *testing.T, path string) {
	err := os.RemoveAll(path)
	if err != nil {
		t.Errorf("Error removing test folder '%v': %v", path, err)
	}
}

func prepareTestDir(t *testing.T) (string, string) {
	dir := filepath.Join("tmp", strings.Replace(t.Name(), "/", "_", -1))
	err := os.MkdirAll(dir, os.FileMode(0700))
	if err != nil {
		t.Errorf("Error creating test directory: %v", err)
	}
	outFile := filepath.Join(dir, "out.txt")

	return dir, outFile
}

func TestCopy(t *testing.T) {
	t.Run("Offset: 0, Limit: 0", func(t *testing.T) {
		dir, outFile := prepareTestDir(t)
		defer dropTempFolder(t, dir)

		err := Copy("testdata/input.txt", outFile, 0, 0)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("Offset: 10, Limit: 0", func(t *testing.T) {
		dir, outFile := prepareTestDir(t)
		defer dropTempFolder(t, dir)

		err := Copy("testdata/input.txt", outFile, 10, 0)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("Offset: 0, Limit: 10", func(t *testing.T) {
		dir, outFile := prepareTestDir(t)
		defer dropTempFolder(t, dir)

		err := Copy("testdata/input.txt", outFile, 0, 10)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("Offset: 10, Limit: 10", func(t *testing.T) {
		dir, outFile := prepareTestDir(t)
		defer dropTempFolder(t, dir)

		err := Copy("testdata/input.txt", outFile, 10, 10)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("Expect source file does not exist", func(t *testing.T) {
		dir, outFile := prepareTestDir(t)
		defer dropTempFolder(t, dir)

		err := Copy("testdata/nonexistent.txt", outFile, 10, 10)
		if !errors.Is(err, ErrFileDoesNotExist) {
			t.Errorf("Expected ErrFileDoesNotExist, got %v", err)
		}
	})
}
