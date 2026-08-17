package project

import "context"

type Service interface {
	GetProjects(ctx context.Context) ([]Project, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetProjects(ctx context.Context) ([]Project, error) {
	return s.repository.GetProjects(ctx)
}
