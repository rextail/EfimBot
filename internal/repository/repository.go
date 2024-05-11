package repository

import (
	"EfimBot/internal/models"
	"context"
)

type Project interface {
	Create(ctx context.Context, project models.Project) error
	MakeCompleted(ctx context.Context, ID int) error
	GetID(ctx context.Context, username string) (ID int, err error)
	GetByID(ctx context.Context, ID int) (project *models.Project, err error)
}

type User interface {
	Create(ctx context.Context, user models.User) error
	Delete(ctx context.Context, user models.User) error
	GetID(ctx context.Context, username string, subdepartment string) (ID int, err error)
	AllDepartmentWorkers(ctx context.Context, department string) (workers string, err error)
	AllSubDepartmentWorkers(ctx context.Context, subdeparment string) (workers string, err error)
}

type Task interface {
	Done(ctx context.Context, task models.Task)
	CreateExternal(ctx context.Context, task models.Task)
	AdressExternalTo(ctx context.Context, task models.Task)
	CreateInternalAttachedToUser(ctx context.Context)
}

type Department interface {
	GetID(ctx context.Context, department models.Department) (ID int, err error)
	Create(ctx context.Context, department models.Department) error
	Delete(ctx context.Context, department models.Department) error
	GetAllWorkers(ctx context.Context)
}

type SubDepartment interface {
	GetID(ctx context.Context)
	Create(ctx context.Context)
	Delete(ctx context.Context)
}
