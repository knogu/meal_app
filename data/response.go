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
	EventID  int
	Event    Event
	UserID   int
	User     User
	IsNeeded bool
	Date     time.Time
}
