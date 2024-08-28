package host

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/guionardo/gs-ops/src/host/middlewares"
	"github.com/guionardo/gs-ops/src/host/routes"
)

var (
	healthy int32
)

func GetServer(listenAddr string) (*http.Server, *slog.Logger) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	logger.Info("Server is starting...")

	router := http.NewServeMux()
	routes.SetupAPI(router)

	server := &http.Server{
		Addr:    listenAddr,
		Handler: middlewares.GetMiddlewares(logger, router),
		// ErrorLog:     logger.
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	return server, logger
}

func RunServer(server *http.Server, logger slog.Logger) {
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Info("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Error("Could not gracefully shutdown the server", "error", err)
		}
		close(done)
	}()

	logger.Info("Server is ready to handle requests", "address", server.Addr)
	atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error("Could not listen", "address", server.Addr, "error", err)
	}

	<-done
	logger.Info("Server stopped")
}
