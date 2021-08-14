package handler

import (
	"encoding/json"
	"meal_api/data"
	"meal_api/json_structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleInvitedUserPost(c *gin.Context) {
	team_uuid := c.Param("team_uuid")
	team, err := data.FetchTeamByUUid(team_uuid)
	if err != nil {
		handleError(c, err)
		return
	}

	params, err := json_structs.NewUserPostRequestBody(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = team.PasswordIsValid(params.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	_, err = data.CreateUserByRequestBody(params, team.UUID)
	if err != nil {
		handleError(c, err)
		return
	}

	return
}

type ReturnToOrganizerPost struct {
	UUID string `json:"uuid"`
}

func HandleOrganizersPost(c *gin.Context) {
	params, err := json_structs.NewUserPostRequestBody(c)
	if err != nil {
		handleError(c, err)
		return
	}

	team, err := data.CreateTeamByPassword(params.Password)
	if err != nil {
		handleError(c, err)
		return
	}

	_, err = data.CreateUserByRequestBody(params, team.UUID)
	if err != nil {
		handleError(c, err)
		return
	}

	err = team.CreateDefaultEvents()
	if err != nil {
		handleError(c, err)
		return
	}

	output, err := json.MarshalIndent(ReturnToOrganizerPost{UUID: team.UUID}, "", "\t\t")
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, output)
	return
}

func HandleUsersPut(c *gin.Context) {
	user_id := c.Param("user_id")
	params, err := json_structs.NewUserPutParams(c)
	if err != nil {
		handleError(c, err)
		return
	}

	err = data.IsAuthorized(user_id, params.LineToken)
	if err != nil {
		handleError(c, err)
		return
	}

	err = data.UpdateUserSetting(user_id, params.UserSettings)
	if err != nil {
		handleError(c, err)
	}
	return
}
