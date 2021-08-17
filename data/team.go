package data

import (
	"meal_api/xer"
	"sort"
	"time"

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
	dinner := Event{Team: team, Sort: 2, Name: "dinner"}
	result = Db.Create(&dinner)
	return result.Error
}

func (team *Team) FetchTeamResponses(eventID uint, Date time.Time, users []User) (responseListJson []ResponseJson, err error) {
	// eventIDとのリレーションがない場合の例外処理は、とりあえずなし
	for _, user := range users {
		Response, errTemp := FetchResponseByMultipleKeys(user.LineID, eventID, Date)
		if errTemp != nil {
			switch errors.Cause(errTemp).(type) {
			case xer.Err4xx:
				continue
			default:
				return nil, errTemp
			}
		}
		responseJson := ResponseJson{ResponseID: Response.ID, IsNeeded: Response.IsNeeded, UserID: Response.UserID}
		responseListJson = append(responseListJson, responseJson)
	}
	return responseListJson, nil
}

func (team *Team) EventsWithResponses(from time.Time, days int) (dateEvents []DateJson, err error) {
	var Events EventList
	Db.Where("team_uuid = ?", team.UUID).Find(&Events)
	sort.Sort(Events)
	var users []User
	Db.Where("team_uuid = ?", team.UUID).Find(&users)
	for i := 0; i < days; i++ {
		var eventsListJson []EventJson
		date := from.AddDate(0, 0, i)
		for _, event := range Events {
			var eventJson EventJson
			event.setIDandName2Json(&eventJson)
			eventsListJson = append(eventsListJson, eventJson)
			responseListJson, err := team.FetchTeamResponses(event.ID, date, users)
			if err != nil {
				return nil, err
			}
			eventsListJson[len(eventsListJson)-1].Response = responseListJson
		}
		var dateEvent DateJson
		dateEvent.Date = date
		dateEvent.Events = eventsListJson
		dateEvents = append(dateEvents, dateEvent)
	}

	return dateEvents, nil
}
