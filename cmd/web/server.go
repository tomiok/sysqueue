package web

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

type Server struct {
	*http.Server
}

func NewServer(addr string, handler chi.Router) *Server {
	return &Server{
		Server: &http.Server{
			Addr:              addr,
			Handler:           handler,
			TLSConfig:         nil,
			ReadTimeout:       10 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
			WriteTimeout:      10 * time.Second,
			IdleTimeout:       10 * time.Second,
			MaxHeaderBytes:    1 << 20,
		},
	}
}

func NewRouter() chi.Router {
	return chi.NewRouter()
}

// Start runs ListenAndServe on the http.Server with graceful shutdown.
func (s *Server) Start() {
	logger.Info("server is running", "port", s.Addr)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("closed Server error", err.Error())
		}
	}()
	s.gracefulShutdown()
}

func (s *Server) gracefulShutdown() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT)
	sig := <-quit
	logger.Info("server is shutting down", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.SetKeepAlivesEnabled(false)
	if err := s.Shutdown(ctx); err != nil {
		logger.Error("could not gracefully shutdown the server", err.Error())
	}
	logger.Info("server stopped")
}
