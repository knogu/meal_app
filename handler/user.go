package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"meal_api/data"
	"meal_api/json_structs"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

func ReadUserPostBody(r *http.Request) (json_structs.UserPostRequestBody, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var rbody json_structs.UserPostRequestBody
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

func return400(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Header().Set("Content-Type", "application/json")
	return
}

func HandleInvitedUserPost(w http.ResponseWriter, r *http.Request) {
	rbody, err := ReadUserPostBody(r)
	if err != nil {
		return400(w)
		return
	}
	team_uuid := mux.Vars(r)["team_uuid"]
	var team data.Team
	result := data.Db.Where("uuid = ?", team_uuid).First(&team)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("team not found")
		return400(w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(rbody.Password))
	if err != nil {
		return400(w)
		return
	}

	_, err = data.CreateUserByRequestBody(rbody, team.UUID)
	if err != nil {
		fmt.Println(err)
	}

	return
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
	user := data.User{IsCook: rbody.IsCook, GetResponseNotifications: rbody.GetResponseNotifications, TeamUUID: team.UUID}
	user.CreateByLineToken(rbody.LineToken)

	output, err := json.MarshalIndent(ReturnToOrganizerPost{UUID: team.UUID}, "", "\t\t")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
