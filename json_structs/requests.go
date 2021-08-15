package json_structs

import (
	"meal_api/xer"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type UserPostRequestBody struct {
	IsCook                   bool   `json:"is_cook"`
	GetResponseNotifications bool   `json:"get_response_notifications"`
	Password                 string `json:"password" validate:"required,min=8,max=30"`
}

func NewUserPostRequestBody(c *gin.Context) (params UserPostRequestBody, err error) {
	c.ShouldBindJSON(&params)
	if err != nil {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.JsonFormatInvalid, Detail: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.ParamInvalid, Detail: err.Error()})
	}
	return
}

type UserPutParams struct {
	IsCook                   bool `json:"is_cook"`
	GetResponseNotifications bool `json:"get_response_notifications"`
}

func NewUserPutParams(c *gin.Context) (params UserPutParams, err error) {
	c.ShouldBindJSON(&params)
	if err != nil {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.JsonFormatInvalid, Detail: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.ParamInvalid, Detail: err.Error()})
	}
	return
}

type ResponsesParams struct {
	EventID  int       `json:"event_id"`
	Date     time.Time `json:"date"`
	IsNeeded bool      `json:"is_needed"`
}

func NewResponseParams(c *gin.Context) (params ResponsesParams, err error) {
	c.ShouldBindJSON(&params)
	if err != nil {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.JsonFormatInvalid, Detail: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.ParamInvalid, Detail: err.Error()})
	}
	return
}

type UpdateResponseParams struct {
	IsNeeded bool `json:"is_needed"`
}

func NewUpdateResponseParams(c *gin.Context) (params UpdateResponseParams, err error) {
	c.ShouldBindJSON(&params)
	if err != nil {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.JsonFormatInvalid, Detail: err.Error()})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.ParamInvalid, Detail: err.Error()})
	}
	return
}
