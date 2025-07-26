package handler

import (
	"context"
	"net/http"
	"time"
)

type Handler struct {
	httpServer *http.Server
}

func (h *Handler) Run(port string, handler http.Handler) error {
	h.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   30 * time.Second,
	}

	return h.httpServer.ListenAndServe()
}

func (h *Handler) Shutdown(ctx context.Context) error {
	return h.httpServer.Shutdown(ctx)
}
