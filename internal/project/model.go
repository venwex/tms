package project

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID
	Name        string
	Key         string
	Description string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
