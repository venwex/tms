package project

import (
	"log/slog"
	"net/http"
)

//POST   /api/v1/projects
//GET    /api/v1/projects
//GET    /api/v1/projects/{id}
//PATCH  /api/v1/projects/{id}
//DELETE /api/v1/projects/{id}

type Handler struct {
	service Service
	logger  *slog.Logger
}

func NewHandler(svc Service, l *slog.Logger) Handler {
	return Handler{
		service: svc,
		logger:  l,
	}
}

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {

}
