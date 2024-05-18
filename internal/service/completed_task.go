package service

import (
	"EfimBot/internal/repository"
)

type CompletedTaskServices struct {
	completedTaskRepo repository.CompletedTask
}

func NewCompletedTaskServices(repo repository.CompletedTask) *CompletedTaskServices {
	return &CompletedTaskServices{completedTaskRepo: repo}
}
