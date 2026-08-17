package project

import (
	"time"

	"github.com/google/uuid"
)

type GetProjectsResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Key         string    `json:"key"`
	Description string    `json:"description"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
