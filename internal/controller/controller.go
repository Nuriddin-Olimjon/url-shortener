package controller

import (
	"log"

	"github.com/Nuriddin-Olimjon/url-shortener/config"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/service"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/repository"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/token"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	engine      *gin.Engine
	userService service.UserService
	urlService  service.URLService
	config      *config.Config
	tokenMaker  *token.PasetoMaker
}

func NewController(cfg *config.Config, repo repository.Store) Controller {
	engine := gin.Default()

	pasetoMaker, err := token.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create paseto token maker: %w", err)
	}

	userService := service.NewUserService(repo)
	urlService := service.NewURLService(repo)

	controller := Controller{
		engine: engine,
		userService: userService,
		urlService: urlService,
		config: cfg,
		tokenMaker: pasetoMaker,
	}

}
