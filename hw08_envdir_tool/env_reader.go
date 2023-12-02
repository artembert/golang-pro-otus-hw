package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
)

var ErrNoFileContentFound = errors.New("no file content found")

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ToPairsSlice(env Environment) []string {
	pairs := make([]string, 0, len(env))
	for key, value := range env {
		pairs = append(pairs, key+"="+value.Value)
	}
	return pairs
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if dir == "" {
		return nil, nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("error while reading dir: ", err)
		return nil, err
	}
	envs := Environment{}
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		if info.IsDir() {
			continue
		}
		env, err := getEnvFromFile(dir, info.Name())
		if err == nil {
			envs[entry.Name()] = env
		}
	}
	return envs, nil
}

func getEnvFromFile(dir, fileName string) (EnvValue, error) {
	file, err := os.Open(path.Join(dir, fileName))
	if err != nil {
		return EnvValue{}, err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return EnvValue{}, err
	}
	if stat.Size() == 0 {
		return EnvValue{Value: "", NeedRemove: true}, nil
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		text = strings.TrimRight(text, " ")
		text = strings.TrimRight(text, "\t")
		text = string(bytes.ReplaceAll([]byte(text), []byte("\x00"), []byte("\n")))
		return EnvValue{Value: text, NeedRemove: false}, nil
	}
	return EnvValue{}, ErrNoFileContentFound
}
