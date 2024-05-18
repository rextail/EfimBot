package service

import (
	"EfimBot/internal/dto"
	"EfimBot/internal/models"
	"EfimBot/internal/repository"
	"context"
)

type DepartmentService struct {
	departmentRepo repository.Department
}

func NewDepartmentService(departmentRepo repository.Department) *DepartmentService {
	return &DepartmentService{departmentRepo: departmentRepo}
}

func (d *DepartmentService) Create(ctx context.Context, department dto.Department) error {
	return d.departmentRepo.Create(ctx, models.Department{Name: department.Name, Manager: department.Manager})
}

func (d *DepartmentService) Get(ctx context.Context, name string) (models.Department, error) {
	return d.departmentRepo.Get(ctx, name)
}
