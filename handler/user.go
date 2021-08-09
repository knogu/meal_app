package handler

import (
	"encoding/json"
	"fmt"
	"meal_api/data"
	"meal_api/json_structs"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/go-playground/validator.v9"
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
	team, err := data.FetchTeamByUUid(team_uuid)
	if err != nil {
		fmt.Println("team not found")
		return400(w)
		return
	}

	err = team.PasswordIsValid(rbody.Password)
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

	team, err := data.CreateTeamByPassword(rbody.Password)
	if err != nil {
		fmt.Println(err)
	}

	_, err = data.CreateUserByRequestBody(rbody, team.UUID)
	if err != nil {
		fmt.Println(err)
	}

	output, err := json.MarshalIndent(ReturnToOrganizerPost{UUID: team.UUID}, "", "\t\t")
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
