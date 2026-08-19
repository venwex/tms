package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
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
	logger  *slog.Logger
}

func NewHandler(svc Service, l *slog.Logger) *Handler {
	return &Handler{
		service: svc,
		logger:  l.With("module", "handler"),
	}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WarnContext(ctx, "sign-up: invalid request body",
			slog.Any("err", err),
		)
		httpx.WriteError(w, http.StatusBadRequest, fmt.Errorf("error decoding body: %s", err))
		return
	}

	err := h.service.SignUp(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrRequiredEmail),
			errors.Is(err, ErrRequiredPassword),
			errors.Is(err, ErrRequiredFirstName),
			errors.Is(err, ErrRequiredLastName),
			errors.Is(err, ErrInvalidPassword):
			h.logger.DebugContext(ctx, "sign-up: validation failed",
				slog.String("email", req.Email),
				slog.Any("err", err),
			)
			httpx.WriteError(w, http.StatusBadRequest, err)

		default:
			h.logger.ErrorContext(ctx, "sign-up: error registering user",
				slog.String("email", req.Email),
				slog.Any("err", err),
			)
			httpx.WriteError(w, http.StatusInternalServerError, err)
		}

		return
	}

	h.logger.InfoContext(ctx, "sign-up: user successfully registered",
		slog.String("email", req.Email),
	)

	httpx.WriteJSON(w, http.StatusCreated, map[string]any{
		"message": "signed up successfully",
	})
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WarnContext(ctx, "sign-in: invalid request body",
			slog.Any("err", err),
		)
		httpx.WriteError(w, http.StatusBadRequest, fmt.Errorf("error decoding body: %s", err))
		return
	}

	tokens, err := h.service.SignIn(ctx, req)
	if err != nil {
		switch {
		case errors.Is(err, ErrRequiredEmail):
		}
	}
}
