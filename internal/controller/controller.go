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
	authService service.AuthService
	userService service.UserService
	urlService  service.URLService
	config      *config.Config
	tokenMaker  *token.PasetoMaker
}

func NewController(cfg *config.Config, repo repository.Store) Controller {
	engine := gin.Default()

	pasetoMaker, err := token.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create paseto token maker: %s", err)
	}

	authservice := service.NewAuthService(repo, *pasetoMaker, cfg)
	userService := service.NewUserService(repo)
	urlService := service.NewURLService(repo, cfg.ShortURIDuration)

	controller := Controller{
		engine:      engine,
		authService: authservice,
		userService: userService,
		urlService:  urlService,
		config:      cfg,
		tokenMaker:  pasetoMaker,
	}

	registerCustomValidators()

	controller.setupRouter()

	return controller
}

// Start runs the HTTP server on a specific address.
func (c *Controller) Start(address string) error {
	return c.engine.Run(address)
}
