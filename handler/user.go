package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"meal_api/data"
	"net/http"

	"gopkg.in/go-playground/validator.v9"
)

var (
	IsOrganizerErr = errors.New("usersへのPOSTリクエストのパラメータ is_organizer を指定してください")
)

type UserPostRequestBody struct {
	LineToken                string `json:"line_token" validate:"required"`
	IsCook                   bool   `json:"is_cook" validate:"required"`
	GetResponseNotifications bool   `json:"get_response_notifications" validate:"required"`
	Password                 string `json:"password" validate:"required, min=8, max=30"`
}

type LineProfile struct {
	LineID     string
	LineName   string
	PictureURL string
}

func ReadUserPostBody(r *http.Request) (UserPostRequestBody, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var rbody UserPostRequestBody
	err := json.Unmarshal(body, &rbody)
	if err != nil {
		fmt.Printf(err.Error())
	}
	validate := validator.New()
	err = validate.Struct(rbody)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return rbody, err
}

func HandleInvitedUserPost(w http.ResponseWriter, r *http.Request) {
	// rbody, err := ReadUserPostBody(r)
	// fmt.Println(rbody)
}

func FetchLineProfile(LineToken string) LineProfile {
	// todo: LINE platformから取得するように変更
	return LineProfile{LineID: "id" + LineToken, LineName: "name" + LineToken, PictureURL: "url" + LineToken}
}

type ReturnToOrganizerPost struct {
	UUID string `json:"uuid"`
}

func HandleOrganizersPost(w http.ResponseWriter, r *http.Request) {
	rbody, err := ReadUserPostBody(r)
	if err != nil {
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		return
	}
	var team data.Team
	err = team.Create(rbody.Password)
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
