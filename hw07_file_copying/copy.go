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
		return err
	}

	if err != nil {
		fmt.Println("error seeking file:", err)
		return ErrOffsetExceedsFileSize
	}

	var buffer []byte = nil
	if limit != 0 {
		buffer = make([]byte, limit)
	}

	_, err = io.ReadFull(sourceFile, buffer)
	if err != nil {
		fmt.Println("error reading file:", err)
		return ErrFileRead
	}

	distFile, err := createFile(toPath)
	if err != nil {
		fmt.Println("error creating file:", err)
		return err
	}

	_, err = io.CopyBuffer(distFile, sourceFile, buffer)
	if err != nil {
		fmt.Println("error writing to file:", err)
		return ErrFileWrite
	}

	err = distFile.Close()
	if err != nil {
		fmt.Println("error closing file:", err)
		return err
	}
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

func createFile(path string) (*os.File, error) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("error creating file:", err)
	}

	return file, err
}
