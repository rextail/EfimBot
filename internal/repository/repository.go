package repository

import (
	"EfimBot/internal/models"
	"context"
)

type Project interface {
	Create(ctx context.Context, project models.Project) error
	Complete(ctx context.Context, projectID int) error
	GetID(ctx context.Context, name string) (id int, err error)
	GetByID(ctx context.Context, projectID int) (project models.Project, err error)
}

type Worker interface {
	Create(ctx context.Context, worker models.Worker) error
	Delete(ctx context.Context, name string, subdepartment string) error
	GetID(ctx context.Context, name string, subdepartment string) (id int, err error)
	GetByID(ctx context.Context, wokerID int) (worker models.Worker, err error)
	AllDepartmentWorkers(ctx context.Context, department string) (workers []models.Worker, err error)
	AllSubDepartmentWorkers(ctx context.Context, subdeparment string) (workers []models.Worker, err error)
}

type Task interface {
	Create(ctx context.Context, code string, description string) error
	AdressInternalTo(ctx context.Context, taskID int, workerID int) error
	AdressExternalTo(ctx context.Context, taskID int, department string) error
	Exist(ctx context.Context, taskID int) (bool, error)
}

type CompletedTask interface {
	Create(ctx context.Context, code string, description string) error
	Delete(ctx context.Context, taskID int) error
}

type Department interface {
	Get(ctx context.Context, name string) (deparment models.Department, err error)
	Create(ctx context.Context, department models.Department) error
	Delete(ctx context.Context, deparment models.Department) error
}

type SubDepartment interface {
	Get(ctx context.Context, name string) (sub models.SubDepartment, err error)
	Create(ctx context.Context, sub models.SubDepartment) error
	Delete(ctx context.Context, sub models.SubDepartment) error
}
