package handler

import (
	"meal_api/data"
	"meal_api/json_structs"
	"meal_api/xer"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleResponsesPost(c *gin.Context) {
	userIDByPath := c.Param("user_id")
	params, err := json_structs.NewResponseParams(c)
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
	response_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	response, err := data.FetchResponseByID(response_id)
	if err != nil {
		err = xer.Err4xx{ErrType: xer.PathParamInvalid, Detail: "response id must be parsed into int"}
		handleError(c, err)
		return
	}

	err = AuthorizeResponses(userIDByPath, userIDByToken, response.EventID)
	if err != nil {
		handleError(c, err)
		return
	}

	data.Db.Model(&data.Response{}).Where("id=?", response_id).Update("is_needed", params.IsNeeded)

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
