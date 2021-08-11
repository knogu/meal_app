package json_structs

import (
	"encoding/json"
	"meal_api/own_error"
	"net/http"

	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

type UserPostRequestBody struct {
	LineToken                string `json:"line_token" validate:"required"`
	IsCook                   bool   `json:"is_cook"`
	GetResponseNotifications bool   `json:"get_response_notifications"`
	Password                 string `json:"password" validate:"required,min=8,max=30"`
}

type UserSettings struct {
	IsCook                   bool `json:"is_cook"`
	GetResponseNotifications bool `json:"get_response_notifications"`
}

type UserPutParams struct {
	LineToken    string       `json:"line_token" validate:"required"`
	UserSettings UserSettings `json:"user_settings"`
}

// ↓ *UserPutParams を UserPutParams にするとバグる、なぜ？
func (params *UserPutParams) ReadRequestBody(r *http.Request) (err error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	err = json.Unmarshal(body, &params)
	if err != nil {
		errtype := own_error.JsonFormatNotValid{Detail_: err.Error()}
		err = errors.WithStack(own_error.BadRequestError{Detail: errtype})
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errtype := own_error.ParamNotValid{Detail_: err.Error()}
		err = errors.WithStack(own_error.BadRequestError{Detail: errtype})
	}
	return err
}
