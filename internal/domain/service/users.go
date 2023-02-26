package service

import (
	"context"
	"log"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository/sqlc"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/password"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type UserService interface {
	CreateUser(ctx context.Context, params entity.CreateUserParams) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}

func NewUserService(repo repository.Store) UserService {
	return &userService{
		repo: repo,
	}
}

// Compilation checks
var _ UserService = (*userService)(nil)

type userService struct {
	repo repository.Store
}

func (s *userService) CreateUser(ctx context.Context, params entity.CreateUserParams) (entity.User, error) {
	user := entity.User{}

	hashedPassword, err := password.HashPassword(params.Password)
	if err != nil {
		return user, err
	}

	dbUser, err := s.repo.CreateUser(ctx, sqlc.CreateUserParams{
		FullName: params.FullName,
		Username: params.Username,
		Password: hashedPassword,
	})

	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			constraint := pqErr.ConstraintName

			log.Println("Constraint", constraint)
		}
		return user, err
	}

	user = entity.User{
		ID:       dbUser.ID,
		FullName: dbUser.FullName,
		Username: dbUser.Username,
	}
	return user, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	user := entity.User{}

	dbUser, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println()
			// err = model.NewNotFound("user", "username "+username)
		}
		return user, err
	}

	user = entity.User{
		ID:       dbUser.ID,
		FullName: dbUser.FullName,
		Username: dbUser.Username,
	}

	return user, nil
}
