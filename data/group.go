package data

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

type UserGroup struct {
	UUID string `gorm:"primaryKey;size:36"`
	Password string `gorm:"size:60"`
}

func (group *UserGroup) CreateGroup(password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password),12)
	group.Password = string(hashed)
	u, _ := uuid.NewRandom()
	uu := u.String()
    group.UUID = uu
	result := db.Create(&group)
	fmt.Println(result.Error)
	return result.Error
}
