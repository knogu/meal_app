package data

import (
	"meal_api/xer"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Team struct {
	UUID     string `gorm:"primaryKey;size:36"`
	Password string `gorm:"size:60"`
}

func (team *Team) PasswordIsValid(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(password))
	if err != nil {
		err = xer.Err4xx{ErrType: xer.WrongPassword}
	}
	return nil
}

func FetchTeamByUUid(uuid string) (Team, error) {
	var team Team
	var err error
	result := Db.Where("uuid = ?", uuid).First(&team)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = xer.Err4xx{ErrType: xer.TeamNotFound}
	}
	return team, err
}

func CreateTeamByPassword(password string) (Team, error) {
	var team Team
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return team, errors.WithStack(err)
	}

	team.Password = string(hashed)
	u, err := uuid.NewRandom()
	if err != nil {
		return team, errors.WithStack(err)
	}

	uu := u.String()
	team.UUID = uu
	result := Db.Create(&team)
	return team, errors.WithStack(result.Error)
}

func (team Team) CreateDefaultEvents() (err error) {
	lunch := Event{Team: team, Sort: 1, Name: "lunch"}
	result := Db.Create(&lunch)
	if result.Error != nil {
		return errors.WithStack(result.Error)
	}
	dinner := Event{Team: team, Sort: 1, Name: "dinner"}
	result = Db.Create(&dinner)
	return result.Error
}
