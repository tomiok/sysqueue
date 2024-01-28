package handler

import (
	"net/http"
	"sysqueue/internal/queue"
)

type Handler struct {
	*queue.EvtService
	*queue.Service
}

func (h *Handler) Queue(w http.ResponseWriter, r *http.Request) {
	
}
