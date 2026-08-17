package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tms/internal/httpx"
)

/*
auth
POST /api/v1/auth/register
POST /api/v1/auth/login

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	FirstName string
	LastName  string

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

*/

type Handler struct {
	service Service
}

func NewHandler(svc Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// slog
		httpx.WriteError(w, http.StatusBadRequest, fmt.Errorf("error decoding body: %s", err))
		return
	}

	err := h.service.SignUp(ctx, req)
	switch err {

	}
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

}
