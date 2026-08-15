package project

import "net/http"

//POST   /api/v1/projects
//GET    /api/v1/projects
//GET    /api/v1/projects/{id}
//PATCH  /api/v1/projects/{id}
//DELETE /api/v1/projects/{id}

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (h *Handler) CreateProject(w http.ResponseWriter, r *http.Request) {

}
