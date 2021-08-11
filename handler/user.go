package handler

import (
	"encoding/json"
	"meal_api/data"
	"meal_api/json_structs"
	"meal_api/own_error"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

func ReadUserPostBody(r *http.Request) (json_structs.UserPostRequestBody, error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var rbody json_structs.UserPostRequestBody
	err := json.Unmarshal(body, &rbody)
	if err != nil {
		errtype := own_error.JsonFormatNotValid{Detail_: err.Error()}
		return rbody, errors.WithStack(own_error.BadRequestError{Detail: errtype})
	}

	validate := validator.New()
	err = validate.Struct(rbody)
	if err != nil {
		errtype := own_error.ParamNotValid{Detail_: err.Error()}
		return rbody, errors.WithStack(own_error.BadRequestError{Detail: errtype})
	}
	return rbody, nil
}

func HandleInvitedUserPost(w http.ResponseWriter, r *http.Request) {
	rbody, err := ReadUserPostBody(r)
	if err != nil {
		handleError(w, err)
	}

	team_uuid := mux.Vars(r)["team_uuid"]
	team, err := data.FetchTeamByUUid(team_uuid)
	if err != nil {
		handleError(w, err)
	}

	err = team.PasswordIsValid(rbody.Password)
	if err != nil {
		handleError(w, err)
	}

	_, err = data.CreateUserByRequestBody(rbody, team.UUID)
	if err != nil {
		handleError(w, err)
	}

	return
}

type ReturnToOrganizerPost struct {
	UUID string `json:"uuid"`
}

func HandleOrganizersPost(w http.ResponseWriter, r *http.Request) {
	rbody, err := ReadUserPostBody(r)
	if err != nil {
		handleError(w, err)
		return
	}

	team, err := data.CreateTeamByPassword(rbody.Password)
	if err != nil {
		handleError(w, err)
	}

	_, err = data.CreateUserByRequestBody(rbody, team.UUID)
	if err != nil {
		handleError(w, err)
	}

	err = team.CreateDefaultEvents()
	if err != nil {
		handleError(w, err)
	}

	output, err := json.MarshalIndent(ReturnToOrganizerPost{UUID: team.UUID}, "", "\t\t")
	if err != nil {
		handleError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func HandleUsersPut(w http.ResponseWriter, r *http.Request) {
	user_id := mux.Vars(r)["user_id"]
	var params json_structs.UserPutParams
	params.ReadRequestBody(r)
	err := data.IsAuthorized(user_id, params.LineToken)
	if err != nil {
		handleError(w, err)
	}

	err = data.UpdateUserSetting(user_id, params.UserSettings)
	if err != nil {
		handleError(w, err)
	}
	return
}
