package service

import (
	"context"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
)

type UserService interface {
	CreateUser(ctx context.Context, params entity.CreateUserParams) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}

type userService struct {
	repo repo.Store
}

func NewUserService(repo repo.Store) UserService {
	return &UserService{
		repo: repo,
	}
}
