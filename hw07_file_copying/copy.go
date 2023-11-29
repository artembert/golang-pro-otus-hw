package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrFileRead              = errors.New("failed to read file")
	ErrFileWrite             = errors.New("failed to write to file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileDoesNotExist      = errors.New("file does not exist")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	sourceFile, err := openFile(fromPath, offset)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
		}
	}(sourceFile)
	if sourceFile == nil {
		return ErrFileRead
	}

	if err != nil {
		fmt.Println("error seeking file:", err)
		return ErrOffsetExceedsFileSize
	}

	distFile, err := os.Create(toPath)
	if err != nil {
		fmt.Println("error creating file:", err)
		return err
	}

	if limit == 0 {
		_, err = io.Copy(distFile, sourceFile)
	} else {
		_, err = io.CopyN(distFile, sourceFile, limit)
	}
	if err != nil {
		if err != io.EOF {
			fmt.Println("error writing to file:", err)
			return ErrFileWrite
		}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
		}
	}(distFile)
	return nil
}

func openFile(path string, offset int64) (*os.File, error) {
	file, err := os.Open(path)
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileDoesNotExist
		} else {
			return nil, ErrUnsupportedFile
		}
	}

	return file, err
}
