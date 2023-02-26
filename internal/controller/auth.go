package controller

import (
	"net/http"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

// @Router /register [post]
// @Tags user
// @Summary Register new user
// @Accept json
// @Produce json
// @Param payload body entity.CreateUserParams true "Body"
// @Success 200 {object} entity.User
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

// @Router /token [post]
// @Tags auth
// @Summary Create new token
// @Accept json
// @Produce json
// @Param payload body entity.LoginParams true "Body"
// @Success 200 {object} entity.LoginResponse
func (c *Controller) GetAccessToken(ctx *gin.Context) {
	var (
		payload entity.LoginParams
	)
	if ok := bindJson(ctx, &payload); !ok {
		return
	}

	rsp, err := c.authService.GetAccessToken(ctx, payload)
	if ok := handleError(ctx, err); !ok {
		return
	}

	ctx.JSON(http.StatusOK, *rsp)
}
