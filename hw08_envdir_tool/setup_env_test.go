package main

import (
	"os"
	"testing"
)

func TestSetupEnv(t *testing.T) {
	t.Run("Should set env", func(t *testing.T) {
		env := Environment{
			"TEST_VAR": EnvValue{
				Value:      "test_value",
				NeedRemove: false,
			},
		}
		_ = os.Unsetenv("TEST_VAR")

		SetupEnv(env)

		value, exists := os.LookupEnv("TEST_VAR")
		if !exists || value != "test_value" {
			t.Errorf("Expected TEST_VAR to be set to 'test_value', got '%v'", value)
		}
	})

	t.Run("Should drop variable if it marks as NeedRemove", func(t *testing.T) {
		env := Environment{
			"TEST_VAR": EnvValue{
				Value:      "",
				NeedRemove: true,
			},
		}
		_ = os.Setenv("TEST_VAR", "I_am_existing")

		SetupEnv(env)

		val, exists := os.LookupEnv("TEST_VAR")
		if exists {
			t.Errorf("Expected TEST_VAR to be removed, but it still exists '%v'", val)
		}
	})
}
