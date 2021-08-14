package handler

import (
	"meal_api/xer"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func handleError(c *gin.Context, err error) {
	switch cause := errors.Cause(err).(type) {
	case xer.Err4xx:
		cause.Return(c)
	default:
		xer.Process500(c, err)
	}
}
