package handler

import (
	"net/http"
	"meal_api/data"
	"encoding/json"
	"fmt"
	"errors"
)

var (
	IsOrganizerErr = errors.New("usersへのPOSTリクエストのパラメータ is_organizer を指定してください")
)

type UserPostRequestBody struct {
	LineToken string `json:"line_token"`
	IsCook bool `json:"is_cook"`
	GetResponseNotifications bool `json:"get_response_notifications"`
	Password string `json:"password"`
}

func GetUserPostBody(r *http.Request) UserPostRequestBody {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var rbody UserPostRequestBody
	json.Unmarshal(body, &rbody)
	return rbody
}

func HandleInvitedUserPost(w http.ResponseWriter, r *http.Request) {
	rbody := GetUserPostBody(r)
	fmt.Println(rbody)
}

func HandleOrganizersPost(w http.ResponseWriter, r *http.Request) {
	rbody := GetUserPostBody(r)
	var group data.UserGroup
	err := group.CreateGroup(rbody.Password)
	fmt.Println(group.UUID)
	
	if err != nil {
		fmt.Println(err)
	}
}
