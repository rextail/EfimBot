package service

import (
	"EfimBot/internal/models"
	"EfimBot/internal/repository"
	"context"
)

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Create(ctx context.Context, user models.User) error {
	return u.Create(ctx, user)
}

func (u *UserService) Delete(ctx context.Context, user models.User) error {
	return u.Delete(ctx, user)
}

func (u *UserService) GetID(ctx context.Context, username string, subdepartment string) (ID string, err error) {
	return u.GetID(ctx, username, subdepartment)
}
