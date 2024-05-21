package service

import (
	"EfimBot/internal/dto"
	"EfimBot/internal/models"
	"EfimBot/internal/repository"
	dberrors "EfimBot/internal/repository/db_errors"
	"context"
	"database/sql"
	"errors"
	"log"
)

type ProjectService struct {
	projectRepo repository.Project
}

func NewProjectService(projectRepo repository.Project) *ProjectService {
	return &ProjectService{projectRepo: projectRepo}
}

func (p *ProjectService) Create(ctx context.Context, project dto.CreateProjectInput) error {
	model := models.Project{Code: project.Code, Name: project.Name, ProjectManager: project.ProjectManager}
	err := p.projectRepo.Create(ctx, model)
	if err != nil {
		log.Printf("Failed to create project %v", err)
		return err
	}
	return nil
}

func (p *ProjectService) Complete(ctx context.Context, ID int) error {
	return p.projectRepo.Complete(ctx, ID)
}

func (p *ProjectService) GetID(ctx context.Context, codeID string) (ID int, err error) {
	return p.projectRepo.GetID(ctx, codeID)
}

func (p *ProjectService) GetByID(ctx context.Context, projectID int) (project models.Project, err error) {
	project, err = p.projectRepo.GetByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Can't find project with requested id: %d, err: %v", projectID, err)
			return models.Project{}, dberrors.ErrNoRows
		}
		return models.Project{}, err
	}
	return project, nil
}
