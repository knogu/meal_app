package data

import (
	"meal_api/xer"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Event struct {
	ID       uint   `gorm:"primaryKey"`
	TeamUUID string `gorm:"size:36"`
	Team     Team
	Sort     int
	Name     string
}

func FetchEventById(id int) (event Event, err error) {
	Result := Db.First(&event, id)
	if errors.Is(Result.Error, gorm.ErrRecordNotFound) {
		err = xer.Err4xx{ErrType: xer.EventNotFound}
	}
	return event, errors.WithStack(err)
}
