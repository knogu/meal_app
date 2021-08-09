package data

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Team struct {
	UUID     string `gorm:"primaryKey;size:36"`
	Password string `gorm:"size:60"`
}

func (team *Team) Create(password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	team.Password = string(hashed)
	u, _ := uuid.NewRandom()
	uu := u.String()
	team.UUID = uu
	result := Db.Create(&team)
	return result.Error
}
