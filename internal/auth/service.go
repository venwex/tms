package auth

import (
	"context"
	"fmt"
	"strings"
	"tms/internal/user"

	"github.com/google/uuid"
)

type Service interface {
	SignUp(ctx context.Context, req RegisterRequest) error
	SignIn(ctx context.Context, req LoginRequest) (*Tokens, error)
}

type service struct {
	repository users.Repository
}

func NewService(repository users.Repository) Service {
	return &service{
		repository: repository,
	}
}

/*
type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
*/

func (s *service) SignUp(ctx context.Context, req RegisterRequest) error {
	if strings.TrimSpace(req.Email) == "" {
		return ErrRequiredEmail
	}

	if strings.TrimSpace(req.Password) == "" {
		return ErrRequiredPassword
	}

	if strings.TrimSpace(req.FirstName) == "" {
		return ErrRequiredFirstName
	}

	if strings.TrimSpace(req.LastName) == "" {
		return ErrRequiredLastName
	}

	if len(req.Password) < 8 {
		return ErrInvalidPassword
	}

	email := strings.ToLower(strings.TrimSpace(req.Email))
	firstName := strings.TrimSpace(req.FirstName)
	lastName := strings.TrimSpace(req.LastName)

	passwordHash, err := hashPassword(req.Password)
	if err != nil {
		return fmt.Errorf("could not hash password: %w", err)
	}

	user := &users.User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: passwordHash,
		FirstName:    firstName,
		LastName:     lastName,
	}

	err = s.repository.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	return nil
}

func (s *service) SignIn(ctx context.Context, req LoginRequest) (*Tokens, error) {
	if strings.TrimSpace(req.Email) == "" {
		return nil, ErrRequiredEmail
	}

	if strings.TrimSpace(req.Password) == "" {
		return nil, ErrRequiredPassword
	}

	if len(req.Password) < 8 {
		return nil, ErrInvalidPassword
	}

	return nil, nil
}
