package controller

import (
	swagDocs "github.com/Nuriddin-Olimjon/url-shortener/docs/api/swagger"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/controller/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (c *Controller) setupRouter() {

	// Handle all shorts
	c.engine.GET("/:uri", c.HandleShortURI)

	apiPrefix := "/api"
	api := c.engine.Group(apiPrefix)

	// Swagger conf
	swagDocs.SwaggerInfo.BasePath = apiPrefix
	swagDocs.SwaggerInfo.Title = "URL shortener docs"
	swagDocs.SwaggerInfo.Version = "1.1"
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Open endpoints
	api.POST("/register", c.RegisterUser)
	api.POST("/token", c.GetAccessToken)

	// Token auth endpoints
	authRoutes := api.Use(middleware.AuthMiddleware(*c.tokenMaker))
	authRoutes.GET("/me", c.GetCurrentUser)
	authRoutes.GET("/urls", c.GetCurrentUserUrls)
	authRoutes.POST("/short-uri", c.CreateShortURI)
	authRoutes.PUT("/short-uri", c.UpdateShortURI)
}
