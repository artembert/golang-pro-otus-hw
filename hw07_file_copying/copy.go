package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	// Place your code here.
	return nil

}

func openFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println("error opening file:", err)
		}
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
		}
	}(file)

	return file, err
}

func createFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("error creating file:", err)
	}

	return file, err
}
