package service

import (
	"EfimBot/internal/models"
	"EfimBot/internal/repository"
	"context"
)

type ProjectService struct {
	projectRepo repository.Project
}

func NewProjectService(projectRepo repository.Project) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

func (p *ProjectService) Create(ctx context.Context, project models.Project) error {
	return p.projectRepo.Create(ctx, project)
}
func (p *ProjectService) MakeCompleted(ctx context.Context, ID int) error {
	return p.projectRepo.MakeCompleted(ctx, ID)
}
func (p *ProjectService) GetID(ctx context.Context, codeID string) (ID int, err error) {
	ID, err = p.projectRepo.GetID(ctx, codeID)
	if err != nil {
		return 0, err
	}
	return ID, nil
}
