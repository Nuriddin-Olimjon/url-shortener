package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *Controller) setupRouter() {
	basePrefix := "/api"

	api := c.engine.Group(basePrefix)

	// Swagger conf
	// swagDocs.SwaggerInfo.BasePath = basePrefix
	// swagDocs.SwaggerInfo.Title = "Tourniquet api docs"
	// swagDocs.SwaggerInfo.Version = "1"
	// api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Open endpoints
	api.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// api.POST("/login", c.LoginAdmin)

	// // Basic auth endpoints
	// api.POST("/organizations", middleware.BasicAuth(server.config), server.CreateOrganization)
	// api.POST("/locations", middleware.BasicAuth(server.config), server.CreateLocation)
	// api.POST("/admins", middleware.BasicAuth(server.config), server.CreateAdmin)
	// api.POST("/admin-orgs", middleware.BasicAuth(server.config), server.CreateAdminOrg)

	// // Token auth endpoints
	// tokenAuthRoutes := api.Use(middleware.AuthMiddleware(server.tokenMaker))
	// tokenAuthRoutes.GET("/me", server.GetCurrentAdmin)
	// tokenAuthRoutes.GET("/organizations/:id/locations", server.GetOrgLocations)
}
