package loggerzap

import (
	"fmt"
	"os"
	"testing"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/loglevel"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	cases := []struct {
		message  string
		logLevel loglevel.LogLevel
	}{
		{
			message:  "Test Info Message",
			logLevel: loglevel.Info,
		},
		{
			message:  "Something goes wrong",
			logLevel: loglevel.Error,
		},
		{
			message:  "Deprecated function called",
			logLevel: loglevel.Warn,
		},
		{
			message:  "Connection established",
			logLevel: loglevel.Debug,
		},
	}
	for _, testCase := range cases {
		t.Run(string(testCase.logLevel), func(t *testing.T) {
			tmpfile, err := os.CreateTemp("", "log")
			if err != nil {
				t.Fatal(err)
			}
			defer func() {
				_ = os.Remove(tmpfile.Name())
			}()

			logger, err := New(loglevel.Debug, tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}
			switch testCase.logLevel {
			case loglevel.Info:
				logger.Info(testCase.message)
			case loglevel.Error:
				logger.Error(testCase.message)
			case loglevel.Debug:
				logger.Debug(testCase.message)
			case loglevel.Warn:
				logger.Warn(testCase.message)
			}

			content, err := os.ReadFile(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}

			require.Contains(t, string(content), testCase.message)
			require.Contains(t, string(content), fmt.Sprintf("\"level\":\"%s\"", testCase.logLevel))
		})
	}
}
