package service

import "github.com/Nuriddin-Olimjon/url-shortener/internal/repository"

type URLService interface{}

func NewURLService(repo repository.Store) URLService {
	return &urlService{
		repo: repo,
	}
}

type urlService struct {
	repo repository.Store
}

var _ URLService = (*urlService)(nil)
