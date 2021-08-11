package data

type Event struct {
	ID       uint   `gorm:"primaryKey"`
	TeamUUID string `gorm:"size:36"`
	Team     Team
	Sort     int
	Name     string
}
