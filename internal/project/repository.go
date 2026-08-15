package project

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// POST   /api/v1/projects
// GET    /api/v1/projects
// GET    /api/v1/projects/{id}
// PATCH  /api/v1/projects/{id}
// DELETE /api/v1/projects/{id}

type Repository interface {
	GetProjects(ctx context.Context) ([]Project, error)
	GetProject(ctx context.Context, id uuid.UUID) (*Project, error)
	CreateProject(ctx context.Context, project *Project) error
	UpdateProject(ctx context.Context, project *Project) (*Project, error)
	DeleteProject(ctx context.Context, id uuid.UUID) error
}

type projectRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &projectRepository{
		db: db,
	}
}

func (r *projectRepository) GetProjects(ctx context.Context) ([]Project, error) {
	query := `SELECT id, name, key, description, created_at, updated_at, deleted_at 
    		  FROM project
    		  WHERE deleted_at IS NULL	
			  ORDER BY created_at DESC
    		  `

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var project Project

		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Key,
			&project.Description,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.DeletedAt,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) GetProject(ctx context.Context, id uuid.UUID) (*Project, error) {
	query := `SELECT ID, Name, Description, created_at, updated_at, deleted_at
			  FROM project
			  WHERE ID = $1
				AND deleted_at IS NULL
			  `

	var project Project

	err := r.db.QueryRow(ctx, query, id).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
		&project.UpdatedAt,
		&project.DeletedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrProjectNotFound
	}
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (r *projectRepository) CreateProject(ctx context.Context, project *Project) error {
	query := `INSERT INTO project (id, name, key, description) 
			  VALUES ($1, $2, $3, $4)
			  `

	_, err := r.db.Exec(ctx,
		query,
		project.ID,
		project.Name,
		project.Key,
		project.Description,
	)
	if err != nil {
		return err
	}

	return nil
}

/*
type UpdateProjectRequest struct {
	Name string
	Description string
}
*/

func (r *projectRepository) UpdateProject(ctx context.Context, project *Project) (*Project, error) {
	query := `
		UPDATE project
		SET name = $1,
			description = $2,
			updated_at = NOW()
		WHERE ID = $3
			AND deleted_at IS NULL
		RETURNING 
		    id, name, key, description, 
		    created_at, updated_at, deleted_at
        `

	err := r.db.QueryRow(ctx,
		query,
		project.Name,
		project.Description,
		project.ID,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Key,
		&project.Description,
		&project.CreatedAt,
		&project.UpdatedAt,
		&project.DeletedAt,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrProjectNotFound
	}
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (r *projectRepository) DeleteProject(ctx context.Context, id uuid.UUID) error {
	query := `
		UPDATE project
		SET deleted_at = NOW()
		WHERE id = $1
		AND deleted_at IS NULL
		`

	_, err := r.db.Exec(ctx, query)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrProjectNotFound
	}
	if err != nil {
		return err
	}

	return nil
}
