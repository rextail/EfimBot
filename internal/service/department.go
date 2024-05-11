package service

import (
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

func (d *DepartmentService) Create(ctx context.Context, department models.Department) error {
	return d.departmentRepo.Create(ctx, department)
}
