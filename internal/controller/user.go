package controller

import (
	"net/http"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/controller/middleware"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/token"
	"github.com/gin-gonic/gin"
)

// @Router /me [get]
// @Tags user
// @Summary Get current user info
// @Success 200 {object} entity.User
func (c *Controller) GetCurrentUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	rsp, err := c.userService.GetUserByUsername(ctx, authPayload.Username)
	if ok := handleError(ctx, err); !ok {
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

// @Router /urls [get]
// @Tags user
// @Summary Get current user urls
// @Success 200 {object} []entity.URL
func (c *Controller) GetCurrentUserUrls(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	rsp, err := c.urlService.GetURLsByUsername(ctx, authPayload.Username)
	if ok := handleError(ctx, err); !ok {
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}
