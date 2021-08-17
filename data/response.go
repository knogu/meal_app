package data

import (
	"meal_api/json_structs"
	"meal_api/xer"
	"time"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

var (
	defaultTime   = 12
	defaultMinite = 0
	defaultSecond = 0
)

type Response struct {
	gorm.Model
	EventID  int `gorm:"primaryKey"`
	Event    Event
	UserID   string `gorm:"primaryKey"`
	User     User
	Date     time.Time `gorm:"primaryKey"`
	IsNeeded bool
}

func CreateResponseByParams(userID string, params json_structs.ResponsesParams) (response Response, err error) {
	response = Response{EventID: params.EventID, UserID: userID, Date: params.Date, IsNeeded: params.IsNeeded}
	result := Db.Create(&response)

	return response, errors.WithStack(result.Error)
}

func FetchResponseByID(id int) (response Response, err error) {
	err = Db.First(&response, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = xer.Err4xx{ErrType: xer.ResponseNotFound}
	}
	return response, errors.WithStack(err)
}

func FetchResponseByMultipleKeys(user_id string, event_id uint, date time.Time) (response Response, err error) {
	err = Db.Where("user_id = ? and event_id = ? and date = ?", user_id, event_id, date).First(&response).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = xer.Err4xx{ErrType: xer.ResponseNotFound}
	}
	return response, errors.WithStack(err)
}
