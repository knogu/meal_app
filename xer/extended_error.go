package xer

// extended error

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ErrType struct {
	StatusCode int
	Summary    string
}

var (
	JsonFormatInvalid = ErrType{StatusCode: 400, Summary: "request json format is invalid. Failed to parse json"}
	ParamInvalid      = ErrType{StatusCode: 400, Summary: "request json validation failed"}
	WrongPassword     = ErrType{StatusCode: 400, Summary: "password is wrong"}

	TeamNotFound  = ErrType{StatusCode: 404, Summary: "team not found"}
	EventNotFound = ErrType{StatusCode: 404, Summary: "event not found"}
	UserNotFound  = ErrType{StatusCode: 404, Summary: "user not found"}

	NotAuthorized    = ErrType{StatusCode: 403, Summary: "you are not authorized"}
	MethodNotAllowed = ErrType{StatusCode: 405, Summary: "Method not allowed"}
)

type Err4xx struct {
	ErrType
	Detail string
}

type ErrJsonOutput struct {
	Summary string `json:"error_summary"`
	Detail  string `json:"error_detail"`
}

func (err Err4xx) Error() string {
	return "Summary: " + err.Summary + " Detail: " + err.Detail
}

func (err4xx Err4xx) Return(c *gin.Context) {
	fmt.Println("4xx Error:", err4xx.Error())
	c.JSON(err4xx.StatusCode, gin.H{"summary": err4xx.Summary, "detail": err4xx.Detail})
	return
}

func Process500(c *gin.Context, err error) {
	fmt.Printf("%+v\n", err)
	fmt.Printf("process500")
	c.JSON(500, gin.H{"error": fmt.Sprintf("%+v\n", err)})
	return
}
