package service

import (
	"EfimBot/internal/models"
	"EfimBot/internal/repository"
	"context"
)

type SubDepartmentService struct {
	subRepo repository.SubDepartment
}

func NewSubDepartmentService(subRepo repository.SubDepartment) *SubDepartmentService {
	return &SubDepartmentService{subRepo: subRepo}
}

func (d *SubDepartmentService) Create(ctx context.Context, subdepartment models.SubDepartment) error {
	return d.subRepo.Create(ctx, subdepartment)
}
