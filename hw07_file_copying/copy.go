package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrFileWrite             = errors.New("failed to write to file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileDoesNotExist      = errors.New("file does not exist")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	sourceFile, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrFileDoesNotExist
		}
		return ErrUnsupportedFile
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
		}
	}(sourceFile)

	if offset > 0 {
		fileInfo, err := sourceFile.Stat()
		if err != nil {
			return err
		}
		if offset > fileInfo.Size() {
			return ErrOffsetExceedsFileSize
		}
	}

	_, err = sourceFile.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	distFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file:", err)
		}
	}(distFile)

	if limit == 0 {
		_, err = io.Copy(distFile, sourceFile)
	} else {
		_, err = io.CopyN(distFile, sourceFile, limit)
	}
	if err != nil && errors.Is(err, io.EOF) {
		return ErrFileWrite
	}

	return nil
}
