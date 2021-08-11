package data

import (
	"time"

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
	UserID   int `gorm:"primaryKey"`
	User     User
	Date     time.Time `gorm:"primaryKey"`
	IsNeeded bool
}
