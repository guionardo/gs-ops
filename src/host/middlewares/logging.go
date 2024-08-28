package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
)

func logging(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.InfoContext(r.Context(), fmt.Sprintf("%s %s %s %s", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent()), "requestID", requestID)

			}()
			next.ServeHTTP(w, r)
		})
	}
}
