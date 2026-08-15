package project

import "context"

type Service interface {
	GetProjects(ctx context.Context) ([]Project, error)
}

type projectService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &projectService{
		repository: repository,
	}
}

func (s *projectService) GetProjects(ctx context.Context) ([]Project, error) {
	return s.repository.GetProjects(ctx)
}
