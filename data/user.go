package data

import (
)

type User struct {
	LineID string `gorm:"primaryKey"`
	LineName string
	PictureURL string
	IsCook bool
	GroupUuId string `gorm:"size:16"`
	UserGroup UserGroup `gorm:"foreignKey:GroupUuId"`
}
