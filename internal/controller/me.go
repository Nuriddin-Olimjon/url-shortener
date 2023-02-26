package controller

import (
	"net/http"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/controller/middleware"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/token"
	"github.com/gin-gonic/gin"
)

func (c *Controller) GetCurrentUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	rsp, err := c.userService.GetUserByUsername(ctx, authPayload.Username)
	if ok := handleError(ctx, err); !ok {
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}
