package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("Offset: 0, Limit: 0", func(t *testing.T) {
		dir := filepath.Join("tmp", strings.Replace(t.Name(), "/", "_", -1))
		err := os.MkdirAll(dir, os.FileMode(0700))
		if err != nil {
			t.Errorf("Error creating test directory: %v", err)
		}
		defer func(path string) {
			err := os.RemoveAll(path)
			if err != nil {
				t.Errorf("Error removing test folder '%v': %v", dir, err)
			}
		}(dir)
		outFile := filepath.Join(dir, "out.txt")
		err = Copy("testdata/input.txt", outFile, 0, 0)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})
}
