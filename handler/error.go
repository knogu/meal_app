package handler

import (
	"fmt"
	"meal_api/xer"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func handleError(c *gin.Context, err error) {
	switch cause := errors.Cause(err).(type) {
	case xer.Err4xx:
		handle4xx(c, cause)
	default:
		handle500(c, err)
	}
}

func handle4xx(c *gin.Context, err4xx xer.Err4xx) {
	fmt.Println("4xx Error:", err4xx.Error())
	c.JSON(err4xx.StatusCode, gin.H{"summary": err4xx.Summary, "detail": err4xx.Detail})
	return
}

func handle500(c *gin.Context, err error) {
	fmt.Printf("%+v\n", err)
	fmt.Printf("process500")
	c.JSON(500, gin.H{"error": fmt.Sprintf("%+v\n", err)})
	return
}
