package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"sysqueue/internal/queue"
)

type Handler struct {
	*queue.EvtService
	*queue.Service
}

func (h *Handler) Queue(w http.ResponseWriter, r *http.Request) {
	eventID := chi.URLParam(r, "eventID")

}
