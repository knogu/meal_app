package data

import (
)

type User struct {
	LineID string `gorm:"primaryKey"`
	LineName string
	PictureURL string
	IsCook bool
	TeamUUID string `gorm:"size:36"`
	Team Team `gorm:"foreignKey:TeamUUID"`
}
