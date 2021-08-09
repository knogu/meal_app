package data

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Team struct {
	UUID     string `gorm:"primaryKey;size:36"`
	Password string `gorm:"size:60"`
}

func (team *Team) PasswordIsValid(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(password))
}

func FetchTeamByUUid(uuid string) (Team, error) {
	var team Team
	result := Db.Where("uuid = ?", uuid).First(&team)
	return team, result.Error
}

func CreateTeamByPassword(password string) (Team, error) {
	var team Team
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	team.Password = string(hashed)
	u, _ := uuid.NewRandom()
	uu := u.String()
	team.UUID = uu
	result := Db.Create(&team)
	return team, result.Error
}
