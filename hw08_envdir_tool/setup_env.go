package main

import (
	"fmt"
	"os"
)

func SetupEnv(env Environment) {
	for key, value := range env {
		if value.NeedRemove {
			err := os.Unsetenv(key)
			if err != nil {
				fmt.Println("error while unsetting env: ", err)
			}
		} else {
			err := os.Setenv(key, value.Value)
			if err != nil {
				fmt.Printf("error while setting env by name [%s]: %s\n", key, err)
			}
		}
	}
}
