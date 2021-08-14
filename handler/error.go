package handler

import (
	"meal_api/own_error"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func handleError(c *gin.Context, err error) {
	switch cause := errors.Cause(err).(type) {
	case own_error.BadRequestError:
		cause.Return(c)
	default:
		own_error.Process500(c, err)
	}
}
