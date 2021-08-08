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

func CreateGroup(password string) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password),12)
	u, _ := uuid.NewRandom()
	uu := u.String()
    group := UserGroup{UUID: uu, Password: string(hashed)}
	fmt.Println("pointer")
	fmt.Println(&group)
	result := db.Create(&group)
	fmt.Println("created")
	return result.Error
}
