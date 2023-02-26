package service

import (
	"context"
	"time"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository/sqlc"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/apperrors"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/urlgenerator"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"gopkg.in/guregu/null.v4"
)

type URLService interface {
	GetURLsByUsername(ctx context.Context, username string) ([]entity.URL, error)
	CreateShortURI(ctx context.Context, payload entity.CreateURIParams) (*entity.URL, error)
	UpdateShortURI(ctx context.Context, payload entity.UpdateURIParams) (*entity.URL, error)
	GetOriginalUrlFromShort(ctx context.Context, short_uri string) (string, error)
}

func NewURLService(repo repository.Store, urlDuration time.Duration) URLService {
	return &urlService{
		repo:        repo,
		urlDuration: urlDuration,
	}
}

var _ URLService = (*urlService)(nil)

type urlService struct {
	repo        repository.Store
	urlDuration time.Duration
}

func (s *urlService) GetURLsByUsername(ctx context.Context, username string) ([]entity.URL, error) {
	dbUser, err := s.repo.GetUserByUsername(ctx, username)
	if err == pgx.ErrNoRows {
		return nil, apperrors.NewNotFound("user", "given username")
	}
	if err != nil {
		return nil, err
	}

	dbUrls, err := s.repo.GetUserURLS(ctx, dbUser.ID)
	if err != nil {
		return nil, err
	}

	urls := []entity.URL{}
	for _, url := range dbUrls {
		urls = append(urls, entity.URL{
			ID:             url.ID,
			ShortUri:       url.ShortUri.String,
			UserID:         url.UserID,
			RequestedCount: url.RequestedCount,
			OriginalUrl:    url.OriginalUrl,
			ExpiresAt:      url.ExpiresAt,
		})
	}
	return urls, nil
}

func (s *urlService) CreateShortURI(ctx context.Context, payload entity.CreateURIParams) (*entity.URL, error) {
	dbUser, err := s.repo.GetUserByUsername(ctx, payload.Username)
	if err == pgx.ErrNoRows {
		return nil, apperrors.NewNotFound("user", "this token")
	}
	if err != nil {
		return nil, err
	}

	if payload.ShortURI != "" {
		_, err := s.repo.GetURLByShortURI(ctx, null.StringFrom(payload.ShortURI))
		if err == nil {
			return nil, apperrors.NewConflict("short_uri", "given value")
		}
		if err != nil && err != pgx.ErrNoRows {
			return nil, err
		}
	}

	// ***Begin transaction***
	tx, txQuery, err := s.repo.StartTX(ctx)
	if err != nil {
		return nil, err
	}

	urlID, err := txQuery.CreateURL(ctx, sqlc.CreateURLParams{
		UserID:      dbUser.ID,
		OriginalUrl: payload.OriginalUrl,
		ExpiresAt:   time.Now().Add(s.urlDuration),
	})
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	if payload.ShortURI == "" {
		newURI, err := urlgenerator.Encode(urlID)
		if err != nil {
			tx.Rollback(ctx)
			return nil, err
		}
		payload.ShortURI = newURI
	}

	dbURL, err := txQuery.SetShortURLByID(ctx, sqlc.SetShortURLByIDParams{
		ID:        urlID,
		ShortUri:  null.StringFrom(payload.ShortURI),
		ExpiresAt: time.Now().Add(s.urlDuration),
	})
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	// ***Commit transaction***
	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	url := entity.URL{
		ID:             dbURL.ID,
		ShortUri:       dbURL.ShortUri.String,
		UserID:         dbURL.UserID,
		RequestedCount: dbURL.RequestedCount,
		OriginalUrl:    dbURL.OriginalUrl,
		ExpiresAt:      dbURL.ExpiresAt,
	}

	return &url, nil
}

func (s *urlService) UpdateShortURI(ctx context.Context, payload entity.UpdateURIParams) (*entity.URL, error) {
	dbUser, err := s.repo.GetUserByUsername(ctx, payload.Username)
	if err == pgx.ErrNoRows {
		return nil, apperrors.NewNotFound("user", "this token")
	}
	if err != nil {
		return nil, err
	}

	dbURL, err := s.repo.GetURLByShortURI(ctx, null.StringFrom(payload.OldShortURI))
	if err == pgx.ErrNoRows {
		return nil, apperrors.NewNotFound("short_uri", "given value")
	}
	if err != nil {
		return nil, err
	}

	if dbURL.UserID != dbUser.ID {
		return nil, apperrors.NewForbidden("this URL doesn't belong to authenticated user")
	}

	dbURL, err = s.repo.SetShortURLByID(ctx, sqlc.SetShortURLByIDParams{
		ShortUri:  null.StringFrom(payload.NewShortURI),
		ExpiresAt: time.Now().Add(s.urlDuration),
		ID:        dbURL.ID,
	})
	if pqErr, ok := err.(*pgconn.PgError); ok {
		if pqErr.ConstraintName == repository.ShortUriUniqueConstraint {
			err = apperrors.NewConflict("short_uri", "given value")
		}
	}
	if err != nil {
		return nil, err
	}

	url := entity.URL{
		ID:             dbURL.ID,
		ShortUri:       dbURL.ShortUri.String,
		UserID:         dbURL.UserID,
		RequestedCount: dbURL.RequestedCount,
		OriginalUrl:    dbURL.OriginalUrl,
		ExpiresAt:      dbURL.ExpiresAt,
	}

	return &url, nil
}

func (s *urlService) GetOriginalUrlFromShort(ctx context.Context, short_uri string) (string, error) {
	url, err := s.repo.GetURLByShortURI(ctx, null.StringFrom(short_uri))
	if err == pgx.ErrNoRows {
		return "", apperrors.NewNotFound("url", "given short")
	}
	if err != nil {
		return "", nil
	}

	if url.ExpiresAt.Before(time.Now()) {
		return "", apperrors.NewBadRequest("this short url is expired")
	}

	err = s.repo.IncreaseURLRequestedCount(ctx, null.StringFrom(short_uri))
	if err != nil {
		return "", nil
	}

	return url.OriginalUrl, nil
}
