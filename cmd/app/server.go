package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

func NewServer(addr string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:              addr,
		Handler:           handler,
		TLSConfig:         nil,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
}

func NewRouter() chi.Router {
	return chi.NewRouter()
}
