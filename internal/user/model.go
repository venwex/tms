package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	FirstName    string
	LastName     string

	IsActive bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
