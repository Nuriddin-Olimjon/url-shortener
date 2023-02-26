package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Nuriddin-Olimjon/url-shortener/pkg/apperrors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// handleError used to handle errors which are received from service layer
func handleError(c *gin.Context, err error) (ok bool) {
	if err != nil {
		status := apperrors.Status(err)
		if status == http.StatusInternalServerError {
			log.Println(err.Error())
			err = apperrors.NewInternal(err.Error())
		}
		c.JSON(status, err)
		ok = false
		return
	}
	ok = true
	return
}

// used to help extract validation errors
type InvalidArgument struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
}

// BindJson is helper function, returns false if json body is not bound
func bindJson(c *gin.Context, req interface{}) bool {
	if c.ContentType() != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		err := apperrors.NewUnsupportedMediaType(msg)

		c.JSON(err.Status(), gin.H{
			"error": err,
		})
		return false
	}

	if err := c.ShouldBind(req); err != nil {

		if errs, ok := err.(validator.ValidationErrors); ok {
			var invalidArgs []InvalidArgument

			for _, err := range errs {
				invalidArgs = append(invalidArgs, InvalidArgument{
					err.Field(),
					err.Tag(),
				})
			}

			err := apperrors.NewBadRequest("Invalid request parameters. See invalidArgs")

			c.JSON(err.Status(), gin.H{
				"error":       err,
				"invalidArgs": invalidArgs,
			})
			return false
		}

		badRequest := apperrors.NewBadRequest(err.Error())
		c.JSON(badRequest.Status(), gin.H{"error": badRequest})
		return false
	}
	return true
}
