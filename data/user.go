package data

import (
)

type User struct {
	LineID string `gorm:"primaryKey"`
	LineName string
	PictureURL string
	IsCook bool
	GroupUUID string `gorm:"size:36"`
	UserGroup UserGroup `gorm:"foreignKey:GroupUUID"`
}
