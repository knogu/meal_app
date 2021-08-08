package data

import (
	"errors"
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	UUID string
	Password string
}
