package data

import (
	"meal_api/json_structs"
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
