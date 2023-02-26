package controller

import (
	"net/http"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

func (c *Controller) RegisterUser(ctx *gin.Context) {
	var (
		payload entity.CreateUserParams
	)

	if ok := bindJson(ctx, &payload); !ok {
		return
	}

	rsp, err := c.userService.CreateUser(ctx, payload)
	if ok := handleError(ctx, err); !ok {
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}
