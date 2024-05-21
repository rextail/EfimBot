package service

import (
	"EfimBot/internal/models"
	"EfimBot/internal/repository"
	"context"
)

type WorkerService struct {
	workerRepo repository.Worker
}

func NewworkerService(workerRepo repository.Worker) *WorkerService {
	return &WorkerService{workerRepo: workerRepo}
}

func (u *WorkerService) Create(ctx context.Context, worker models.Worker) error {
	return u.Create(ctx, worker)
}

func (u *WorkerService) Delete(ctx context.Context, worker models.Worker) error {
	return u.Delete(ctx, worker)
}

func (u *WorkerService) GetID(ctx context.Context, workername string, subdepartment string) (ID string, err error) {
	return u.GetID(ctx, workername, subdepartment)
}
