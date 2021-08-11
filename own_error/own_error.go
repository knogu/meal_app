package own_error

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// インターフェースを満たすものの例: parameter invalid
type SpecifiedBadRequest interface {
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

type BadRequestError struct {
	Detail SpecifiedBadRequest
}

type ErrJsonOutput struct {
	Summary string `json:"error_summary"`
	Detail  string `json:"error_detail"`
}

func (err BadRequestError) Error() string {
	return "Summary: " + err.Detail.Summary() + " Detail: " + err.Detail.Detail()
}

func (badRequest BadRequestError) Return(w http.ResponseWriter) {
	fmt.Println("4xx Error:", badRequest.Error())
	w.WriteHeader(badRequest.Detail.StatusCode())
	output, err := json.MarshalIndent(ErrJsonOutput{Summary: badRequest.Detail.Summary(), Detail: badRequest.Detail.Detail()}, "", "\t\t")
	if err != nil {
		process500(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func handleError(w http.ResponseWriter, err error) {
	switch cause := errors.Cause(err).(type) {
	case BadRequestError:
		cause.Return(w)
	default:
		process500(w, err)
	}
}

func process500(w http.ResponseWriter, err error) {
	fmt.Printf("%+v\n", err)
	fmt.Printf("process500")
	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	return
}
