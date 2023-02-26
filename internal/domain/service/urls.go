package service

import (
	"context"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository"
)

type UrlService interface {
	CreateUrl(ctx context.Context, params entity.Url) (entity.Url, error)
	GetUrlByID(ctx context.Context, id int32) (entity.Url, error)
	GetUrlsByUserID(ctx context.Context, userID int32) ([]entity.Url, error)
}

type urlService struct {
	repo repository.Store
}
