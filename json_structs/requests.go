package json_structs

import (
	"meal_api/own_error"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type UserPostRequestBody struct {
	LineToken                string `json:"line_token" validate:"required"`
	IsCook                   bool   `json:"is_cook"`
	GetResponseNotifications bool   `json:"get_response_notifications"`
	Password                 string `json:"password" validate:"required,min=8,max=30"`
}

func NewUserPostRequestBody(c *gin.Context) (params UserPostRequestBody, err error) {
	c.ShouldBindJSON(&params)
	if err != nil {
		errtype := own_error.JsonFormatNotValid{Detail_: err.Error()}
		err = errors.WithStack(own_error.BadRequestError{Detail: errtype})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errtype := own_error.ParamNotValid{Detail_: err.Error()}
		err = errors.WithStack(own_error.BadRequestError{Detail: errtype})
	}
	return
}

type UserSettings struct {
	IsCook                   bool `json:"is_cook"`
	GetResponseNotifications bool `json:"get_response_notifications"`
}

type UserPutParams struct {
	LineToken    string       `json:"line_token" validate:"required"`
	UserSettings UserSettings `json:"user_settings"`
}

func NewUserPutParams(c *gin.Context) (params UserPutParams, err error) {
	c.ShouldBindJSON(&params)
	if err != nil {
		errtype := own_error.JsonFormatNotValid{Detail_: err.Error()}
		err = errors.WithStack(own_error.BadRequestError{Detail: errtype})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errtype := own_error.ParamNotValid{Detail_: err.Error()}
		err = errors.WithStack(own_error.BadRequestError{Detail: errtype})
	}
	return
}

type SpecifiedResponseParams struct {
	LineToken string    `json:"line_token" validate:"required"`
	EventID   int       `json:"event_id"`
	Date      time.Time `json:"date"`
	IsNeeded  bool      `json:"is_needed"`
}
