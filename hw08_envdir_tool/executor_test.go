package main

import "testing"

func TestRunCmd(t *testing.T) {
	t.Run("Should pass error code from child process", func(t *testing.T) {
		returnCode := RunCmd([]string{"./testdata/throw_error_code.sh"}, nil)
		if returnCode != 42 {
			t.Errorf("Expected return code 42, got %v", returnCode)
		}
	})
}
