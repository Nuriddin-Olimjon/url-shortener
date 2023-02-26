package service

import (
	"context"

	"github.com/Nuriddin-Olimjon/url-shortener/config"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/apperrors"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/password"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/token"
	"github.com/jackc/pgx/v4"
)

type AuthService interface {
	GetAccessToken(ctx context.Context, payload entity.LoginParams) (*entity.LoginResponse, error)
}

func NewAuthService(repo repository.Store, pasetoMaker token.PasetoMaker, config *config.Config) AuthService {
	return &authService{
		repo:        repo,
		pasetoMaker: pasetoMaker,
		config:      config,
	}
}

var _ AuthService = (*authService)(nil)

type authService struct {
	repo        repository.Store
	pasetoMaker token.PasetoMaker
	config      *config.Config
}

func (s *authService) GetAccessToken(ctx context.Context, payload entity.LoginParams) (*entity.LoginResponse, error) {
	admin, err := s.repo.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			err = apperrors.NewAuthorization("user with this username not found")
		}
		return nil, err
	}

	err = password.CheckPassword(payload.Password, admin.Password)
	if err != nil {
		err = apperrors.NewAuthorization("incorrect password")
		return nil, err
	}

	accessToken, accessPayload, err := s.pasetoMaker.CreateToken(
		payload.Username,
		s.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, err
	}

	rsp := &entity.LoginResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	return rsp, nil
}
