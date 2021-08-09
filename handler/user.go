package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"meal_api/data"
	"net/http"
)

var (
	IsOrganizerErr = errors.New("usersへのPOSTリクエストのパラメータ is_organizer を指定してください")
)

type UserPostRequestBody struct {
	LineToken                string `json:"line_token"`
	IsCook                   bool   `json:"is_cook"`
	GetResponseNotifications bool   `json:"get_response_notifications"`
	Password                 string `json:"password"`
}

type LineProfile struct {
	LineID     string
	LineName   string
	PictureURL string
}

func ReadUserPostBody(r *http.Request) UserPostRequestBody {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var rbody UserPostRequestBody
	err := json.Unmarshal(body, &rbody)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return rbody
}

func HandleInvitedUserPost(w http.ResponseWriter, r *http.Request) {
	rbody := ReadUserPostBody(r)
	fmt.Println(rbody)
}

func FetchLineProfile(LineToken string) LineProfile {
	// todo: LINE platformから取得するように変更
	return LineProfile{LineID: "id" + LineToken, LineName: "name" + LineToken, PictureURL: "url" + LineToken}
}

type ReturnToOrganizerPost struct {
	UUID string `json:"uuid"`
}

func HandleOrganizersPost(w http.ResponseWriter, r *http.Request) {
	rbody := ReadUserPostBody(r)
	var team data.Team
	err := team.Create(rbody.Password)
	if err != nil {
		fmt.Println(err)
	}

	line_profile := FetchLineProfile(rbody.LineToken)
	user := data.User{
		LineID:                   line_profile.LineID,
		LineName:                 line_profile.LineName,
		PictureURL:               line_profile.PictureURL,
		IsCook:                   rbody.IsCook,
		GetResponseNotifications: rbody.GetResponseNotifications,
		TeamUUID:                 team.UUID,
	}
	Result := data.Db.Create(&user)
	if Result.Error != nil {
		fmt.Println(Result.Error)
	}

	output, err := json.MarshalIndent(ReturnToOrganizerPost{UUID: team.UUID}, "", "\t\t")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
