package internalhttp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
)

func loggingMiddleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)

			logger.HTTPRequest(
				r,
				fmt.Sprintf(
					"%s %s %s %d %v %s",
					r.Method,
					r.URL.Path,
					r.Proto,
					http.StatusOK,
					time.Since(startTime),
					r.UserAgent(),
				),
			)
		})
	}
}
