package service

import (
	"context"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository/sqlc"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/apperrors"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/password"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type UserService interface {
	CreateUser(ctx context.Context, params entity.CreateUserParams) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
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

func (s *userService) CreateUser(ctx context.Context, params entity.CreateUserParams) (*entity.User, error) {
	hashedPassword, err := password.HashPassword(params.Password)
	if err != nil {
		return nil, err
	}

	dbUser, err := s.repo.CreateUser(ctx, sqlc.CreateUserParams{
		FullName: params.FullName,
		Username: params.Username,
		Password: hashedPassword,
	})
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			if pqErr.ConstraintName == repository.UserUniqueUsernameConstraint {
				err = apperrors.NewConflict("user", "given username")
			}
		}
		return nil, err
	}

	user := &entity.User{
		ID:       dbUser.ID,
		FullName: dbUser.FullName,
		Username: dbUser.Username,
	}
	return user, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	dbUser, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = apperrors.NewNotFound("user", "given username")
		}
		return nil, err
	}

	user := &entity.User{
		ID:       dbUser.ID,
		FullName: dbUser.FullName,
		Username: dbUser.Username,
	}

	return user, nil
}
