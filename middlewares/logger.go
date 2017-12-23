package middlewares

import (
	"net/http"
	"time"

	"github.com/tehsis/rabbitscore/services/logger"
)

// Logger is a decorator function to log requests
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		logger.Log().Info(
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
