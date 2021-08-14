package handler

import (
	"meal_api/data"
	"meal_api/json_structs"

	"github.com/gin-gonic/gin"
)

func HandleResponsesPost(c *gin.Context) {
	userIDByPath := c.Param("user_id")
	params, err := json_structs.NewSpecifiedResponseParams(c)
	if err != nil {
		handleError(c, err)
		return
	}
	userIDByToken := data.FetchLineProfile(params.LineToken).LineID

	err = AuthorizeResponses(userIDByPath, userIDByToken, params.EventID)
	if err != nil {
		handleError(c, err)
		return
	}

	_, err = data.CreateResponseByParams(userIDByToken, params)
	if err != nil {
		handleError(c, err)
		return
	}

	return
}

func HandleResponsesPut(c *gin.Context) {
	userIDByPath := c.Param("user_id")
	params, err := json_structs.NewSpecifiedResponseParams(c)
	if err != nil {
		handleError(c, err)
		return
	}
	userIDByToken := data.FetchLineProfile(params.LineToken).LineID

	err = AuthorizeResponses(userIDByPath, userIDByToken, params.EventID)
	if err != nil {
		handleError(c, err)
		return
	}

	data.Db.Model(&data.Response{}).Where("user_id=? and event_id=? and date=?", userIDByToken, params.EventID, params.Date).Update("is_needed", params.IsNeeded)

	return
}

func AuthorizeResponses(userIDByPath string, userIDByToken string, eventID int) (err error) {
	err = data.IsAuthorized(userIDByPath, userIDByToken)
	if err != nil {
		return err
	}
	err = data.UserIsAuthorizedEvents(eventID, userIDByToken)
	return err
}
