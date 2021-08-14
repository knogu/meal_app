package data

import (
	"meal_api/json_structs"
	"meal_api/xer"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type User struct {
	LineID                   string `gorm:"primaryKey"`
	LineName                 string
	PictureURL               string
	IsCook                   bool
	GetResponseNotifications bool
	TeamUUID                 string `gorm:"size:36"`
	Team                     Team
}

type LineProfile struct {
	LineID     string
	LineName   string
	PictureURL string
}

func FetchLineProfile(LineToken string) LineProfile {
	// todo: LINE platformから取得するように変更
	return LineProfile{LineID: "id_" + LineToken, LineName: "name_" + LineToken, PictureURL: "url_" + LineToken}
}

func CreateUserByRequestBody(rbody json_structs.UserPostRequestBody, team_uuid string) (User, error) {
	line_profile := FetchLineProfile(rbody.LineToken)
	user := User{
		LineID:                   line_profile.LineID,
		LineName:                 line_profile.LineName,
		PictureURL:               line_profile.PictureURL,
		IsCook:                   rbody.IsCook,
		GetResponseNotifications: rbody.GetResponseNotifications,
		TeamUUID:                 team_uuid,
	}
	Result := Db.Create(&user)

	return user, errors.WithStack(Result.Error)
}

func IsAuthorized(userIdByPath string, userIdByToken string) (err error) {
	if userIdByPath != userIdByToken {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.NotAuthorized})
	}
	return err
}

func FetchUserById(user_id string) (user User, err error) {
	Result := Db.First(&user, "line_id=?", user_id)
	if errors.Is(Result.Error, gorm.ErrRecordNotFound) {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.UserNotFound})
	}
	return user, errors.WithStack(err)
}

func UpdateUserSetting(user_id string, userSettings json_structs.UserSettings) error {
	user, err := FetchUserById(user_id)
	if err != nil {
		return err
	}
	user.IsCook = userSettings.IsCook
	user.GetResponseNotifications = userSettings.GetResponseNotifications
	result := Db.Save(&user)
	return errors.WithStack(result.Error)
}

func UserIsAuthorizedEvents(eventID int, userIDByToken string) (err error) {
	user, err := FetchUserById(userIDByToken)
	if err != nil {
		return err
	}
	event, err := FetchEventById(eventID)
	if err != nil {
		return err
	}

	if user.TeamUUID != event.TeamUUID {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.NotAuthorized})
	}
	return err
}
