package data

type User struct {
	LineID                   string `gorm:"primaryKey"`
	LineName                 string
	PictureURL               string
	IsCook                   bool
	GetResponseNotifications bool
	TeamUUID                 string `gorm:"size:36"`
	Team                     Team   `gorm:"foreignKey:TeamUUID"`
}
