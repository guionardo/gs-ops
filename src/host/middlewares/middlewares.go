package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func GetMiddlewares(logger *slog.Logger, router *http.ServeMux) http.Handler {
	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	return tracing(nextRequestID)(logging(logger)(router))
}
