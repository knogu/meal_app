package data

import (
)

type UserGroup struct {
	UUID string `gorm:"primaryKey"`
	Password string
}
