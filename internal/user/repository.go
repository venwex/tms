package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	CreateUser(ctx context.Context, u *User) error
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, u *User) error {
	query := `
        INSERT INTO users (
            id,
            email,
            password_hash,
            first_name,
            last_name
        )
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err := r.db.Exec(ctx, query,
		u.ID,
		u.Email,
		u.PasswordHash,
		u.FirstName,
		u.LastName,
	)
	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
