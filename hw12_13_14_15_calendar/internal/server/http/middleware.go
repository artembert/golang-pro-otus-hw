package internalhttp

import (
	"fmt"
	"github.com/artembert/golang-pro-otus-hw/hw12_13_14_15_calendar/internal/interfaces/logger"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler, logger logger.Logger) http.HandlerFunc { //nolint:unused
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)

		logger.HttpRequest(r, fmt.Sprintf("%s %s %s %d %v %s", r.Method, r.URL.Path, r.Proto, http.StatusOK, time.Since(startTime), r.UserAgent()))
	}
}
