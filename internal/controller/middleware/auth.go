package middleware

import (
	"net/http"
	"strings"

	"github.com/Nuriddin-Olimjon/url-shortener/pkg/apperrors"
	"github.com/Nuriddin-Olimjon/url-shortener/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.PasetoMaker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(AuthorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := apperrors.NewAuthorization("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := apperrors.NewAuthorization("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != AuthorizationTypeBearer {
			err := apperrors.NewAuthorization("unsupported authorization type " + authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
			return
		}

		ctx.Set(AuthorizationPayloadKey, payload)
		ctx.Next()
	}
}
