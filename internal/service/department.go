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

type DepartmentService struct {
	departmentRepo repository.Department
}

func NewDepartmentService(departmentRepo repository.Department) *DepartmentService {
	return &DepartmentService{departmentRepo: departmentRepo}
}

func (d *DepartmentService) Create(ctx context.Context, department dto.Department) error {
	model := models.Department{Name: department.Name, Manager: department.Manager}

	if err := d.departmentRepo.Create(ctx, model); err != nil {
		log.Printf("Failed to create department, err %v", err)
		return err
	}

	return nil
}

func (d *DepartmentService) Get(ctx context.Context, name string) (models.Department, error) {
	department, err := d.departmentRepo.Get(ctx, name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Department{}, dberrors.ErrNoRows
		}
		return models.Department{}, err
	}

	return department, nil
}
