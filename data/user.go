package data

import "meal_api/json_structs"

type User struct {
	LineID                   string `gorm:"primaryKey"`
	LineName                 string
	PictureURL               string
	IsCook                   bool
	GetResponseNotifications bool
	TeamUUID                 string `gorm:"size:36"`
	Team                     Team   `gorm:"foreignKey:TeamUUID"`
}

type LineProfile struct {
	LineID     string
	LineName   string
	PictureURL string
}

func FetchLineProfile(LineToken string) LineProfile {
	// todo: LINE platformから取得するように変更
	return LineProfile{LineID: "id" + LineToken, LineName: "name" + LineToken, PictureURL: "url" + LineToken}
}

func (user User) CreateByLineToken(LineToken string) error {
	line_profile := FetchLineProfile(LineToken)
	user.LineID = line_profile.LineID
	user.LineName = line_profile.LineName
	user.PictureURL = line_profile.PictureURL

	Result := Db.Create(&user)

	return Result.Error
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

	return user, Result.Error
}
