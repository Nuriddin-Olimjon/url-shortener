package controller

import (
	"net/http"

	"github.com/Nuriddin-Olimjon/url-shortener/internal/controller/middleware"
	"github.com/Nuriddin-Olimjon/url-shortener/internal/domain/entity"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/token"
	"github.com/gin-gonic/gin"
)

// @Router /short-uri [post]
// @Tags uri
// @Summary Create new short uri
// @Accept json
// @Produce json
// @Param payload body entity.CreateURIParams true "Body"
// @Success 200 {object} entity.URL
func (c *Controller) CreateShortURI(ctx *gin.Context) {
	var (
		payload entity.CreateURIParams
	)

	if ok := bindJson(ctx, &payload); !ok {
		return
	}

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	payload.Username = authPayload.Username

	rsp, err := c.urlService.CreateShortURI(ctx, payload)
	if ok := handleError(ctx, err); !ok {
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

// @Router /short-uri [put]
// @Tags uri
// @Summary Update short uri
// @Accept json
// @Produce json
// @Param payload body entity.UpdateURIParams true "Body"
// @Success 200 {object} entity.URL
func (c *Controller) UpdateShortURI(ctx *gin.Context) {
	var (
		payload entity.UpdateURIParams
	)

	if ok := bindJson(ctx, &payload); !ok {
		return
	}

	authPayload := ctx.MustGet(middleware.AuthorizationPayloadKey).(*token.Payload)
	payload.Username = authPayload.Username

	rsp, err := c.urlService.UpdateShortURI(ctx, payload)
	if ok := handleError(ctx, err); !ok {
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (c *Controller) HandleShortURI(ctx *gin.Context) {
	uri := ctx.Param("uri")
	origin, err := c.urlService.GetOriginalUrlFromShort(ctx, uri)
	if ok := handleError(ctx, err); !ok {
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, origin)
}
