package main

import "testing"

func TestRunCmd(t *testing.T) {
	t.Run("Should pass error code from child process", func(t *testing.T) {
		returnCode := RunCmd([]string{"./testdata/throw_error_code.sh"}, nil)
		if returnCode != 42 {
			t.Errorf("Expected return code 42, got %v", returnCode)
		}
	})

	t.Run("Should return Zero code if child process finished successfully", func(t *testing.T) {
		returnCode := RunCmd([]string{"echo", "Do nothing, just printing"}, nil)
		if returnCode != 0 {
			t.Errorf("Expected return code 0, got %v", returnCode)
		}
	})

	t.Run("Should return Error code from invalid command", func(t *testing.T) {
		returnCode := RunCmd([]string{"invalid_command", "&(()&"}, nil)
		if returnCode == 0 {
			t.Errorf("Expected non-zero return code, got %v", returnCode)
		}
	})
}
