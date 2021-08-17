package data

import (
	"meal_api/json_structs"
	"meal_api/xer"
	"sort"
	"time"

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

func CreateUserByRequestBody(lineToken string, rbody json_structs.UserPostRequestBody, team_uuid string) (User, error) {
	line_profile := FetchLineProfile(lineToken)
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
	// todo: Result.Errorが↓以外の場合のハンドリング
	if errors.Is(Result.Error, gorm.ErrRecordNotFound) {
		err = errors.WithStack(xer.Err4xx{ErrType: xer.UserNotFound})
	}
	return user, errors.WithStack(err)
}

func UpdateUserSetting(user_id string, userPutParams json_structs.UserPutParams) error {
	user, err := FetchUserById(user_id)
	if err != nil {
		return err
	}
	user.IsCook = userPutParams.IsCook
	user.GetResponseNotifications = userPutParams.GetResponseNotifications
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

func IsAuthorizedInTeam(lineIDByToken string, targetLineID string) (err error) {
	isSelf := lineIDByToken == targetLineID
	requester, err := FetchUserById(targetLineID)
	if err != nil {
		return err
	}
	target, err := FetchUserById(targetLineID)
	if err != nil {
		return err
	}
	isCook := requester.IsCook && requester.Team.UUID == target.Team.UUID
	if (!isSelf) && !(isCook) {
		return xer.Err4xx{ErrType: xer.NotAuthorized}
	}
	return
}

type DateJson struct {
	Date   time.Time
	Events []EventJson
}

type EventJson struct {
	ID   uint
	Name string
	// ↓nullを許容する
	Response interface{}
}

type ResponseJson struct {
	ResponseID uint
	UserID     string
	IsNeeded   bool
}

func (user *User) EventsWithResponses(from time.Time, days int) (dateEvents []DateJson, err error) {
	var Events EventList
	Db.Where("team_uuid = ?", user.TeamUUID).Find(&Events)
	sort.Sort(Events)
	for i := 0; i < days; i++ {
		var eventsListJson []EventJson
		Date := from.AddDate(0, 0, i)
		for _, event := range Events {
			var eventJson EventJson
			event.setIDandName2Json(&eventJson)
			eventsListJson = append(eventsListJson, eventJson)
			Response, errTemp := FetchResponseByMultipleKeys(user.LineID, event.ID, Date)
			if errTemp != nil {
				switch errors.Cause(errTemp).(type) {
				case xer.Err4xx:
					continue
				default:
					return dateEvents, errTemp
				}
			}
			responseJson := ResponseJson{ResponseID: Response.ID, IsNeeded: Response.IsNeeded, UserID: Response.UserID}
			eventsListJson[len(eventsListJson)-1].Response = responseJson
		}
		var dateEvent DateJson
		dateEvent.Date = Date
		dateEvent.Events = eventsListJson
		dateEvents = append(dateEvents, dateEvent)
	}

	return dateEvents, nil
}
