package own_error

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// インターフェースを満たすものの例: parameter invalid
type ErrDescription interface {
	Summary() string
	Detail() string
	StatusCode() int
}

type JsonFormatNotValid struct {
	Detail_ string
}

func (jsonFormatNotValid JsonFormatNotValid) StatusCode() int {
	return 400
}

func (jsonFormatNotValid JsonFormatNotValid) Summary() string {
	return "request json format is invalid. Failed to parse json"
}

func (jsonFormatNotValid JsonFormatNotValid) Detail() string {
	return jsonFormatNotValid.Detail_
}

type ParamNotValid struct {
	Detail_ string
}

func (paramNotValid ParamNotValid) StatusCode() int {
	return 400
}

func (paramNotValid ParamNotValid) Summary() string {
	return "request json validation failed"
}

func (paramNotValid ParamNotValid) Detail() string {
	return paramNotValid.Detail_
}

type WrongPassword struct {
	Detail_ string
}

func (wrongPassword WrongPassword) StatusCode() int {
	return 400
}

func (wrongPassword WrongPassword) Summary() string {
	return "password is wrong"
}

func (wrongPassword WrongPassword) Detail() string {
	return wrongPassword.Detail_
}

type TeamNotFound struct {
	Detail_ string
}

func (teamNotFound TeamNotFound) StatusCode() int {
	return 404
}

func (teamNotFound TeamNotFound) Summary() string {
	return "team not found"
}

func (teamNotFound TeamNotFound) Detail() string {
	return teamNotFound.Detail_
}

type UserNotFound struct {
	Detail_ string
}

func (userNotFound UserNotFound) StatusCode() int {
	return 404
}
func (userNotFound UserNotFound) Summary() string {
	return "user not found"
}
func (userNotFound UserNotFound) Detail() string {
	return userNotFound.Detail_
}

type NotAuthorized struct {
	Detail_ string
}

func (notAuthorized NotAuthorized) StatusCode() int {
	return 403
}
func (notAuthorized NotAuthorized) Summary() string {
	return "you are not authorized"
}
func (notAuthorized NotAuthorized) Detail() string {
	return notAuthorized.Detail_
}

type EventNotFound struct {
	Detail_ string
}

func (eventNotFound EventNotFound) StatusCode() int {
	return 404
}
func (eventNotFound EventNotFound) Summary() string {
	return "event not found"
}
func (eventNotFound EventNotFound) Detail() string {
	return eventNotFound.Detail_
}

type MethodNotAllowed struct {
	Detail_ string
}

func (methodNotAllowed MethodNotAllowed) StatusCode() int {
	return 405
}
func (methodNotAllowed MethodNotAllowed) Summary() string {
	return "Method not allowed"
}
func (methodNotAllowed MethodNotAllowed) Detail() string {
	return methodNotAllowed.Detail_
}

type BadRequestError struct {
	ErrDescription
}

type ErrJsonOutput struct {
	Summary string `json:"error_summary"`
	Detail  string `json:"error_detail"`
}

func (err BadRequestError) Error() string {
	return "Summary: " + err.Summary() + " Detail: " + err.Detail()
}

func (badRequest BadRequestError) Return(c *gin.Context) {
	fmt.Println("4xx Error:", badRequest.Error())
	// w.WriteHeader(badRequest.Detail.StatusCode())
	output, err := json.MarshalIndent(ErrJsonOutput{Summary: badRequest.Summary(), Detail: badRequest.Detail()}, "", "\t\t")
	if err != nil {
		Process500(c, err)
	}
	c.JSON(badRequest.StatusCode(), output)
}

func Process500(c *gin.Context, err error) {
	fmt.Printf("%+v\n", err)
	fmt.Printf("process500")
	c.JSON(500, gin.H{"error": fmt.Sprintf("%+v\n", err)})
	return
}
