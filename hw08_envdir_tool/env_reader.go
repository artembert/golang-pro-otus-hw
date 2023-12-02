package main

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ToPairsSlice(env Environment) []string {
	var pairs []string
	for key, value := range env {
		pairs = append(pairs, key+"="+value.Value)
	}
	return pairs
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	if dir == "" {
		return Environment{}, nil
	}
	// TODO: prepare envs from dir
	return nil, nil
}
