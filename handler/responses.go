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
	lineToken, err := validateLineToken(c)
	if err != nil {
		return
	}
	params, err := json_structs.NewResponseParams(c)
	if err != nil {
		handleError(c, err)
		return
	}
	userIDByToken := data.FetchLineProfile(lineToken).LineID

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

func validateSpecifiedResponse(c *gin.Context, lineToken string) (response data.Response, err error) {
	userIDByPath := c.Param("user_id")
	userIDByToken := data.FetchLineProfile(lineToken).LineID
	response_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}
	response, err = data.FetchResponseByID(response_id)
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

	return
}

func HandleResponsesPut(c *gin.Context) {
	lineToken, err := validateLineToken(c)
	if err != nil {
		return
	}

	param, err := json_structs.NewUpdateResponseParams(c)
	if err != nil {
		handleError(c, err)
		return
	}
	response, err := validateSpecifiedResponse(c, lineToken)
	if err != nil {
		return
	}
	response.IsNeeded = param.IsNeeded

	data.Db.Save(&response)
	return
}

func validateLineToken(c *gin.Context) (lineToken string, err error) {
	lineToken = c.Query("line_token")
	if len(lineToken) == 0 {
		err = xer.Err4xx{ErrType: xer.MissingLineToken}
	}
	return lineToken, err
}

func HandleResponsesDelete(c *gin.Context) {
	lineToken, err := validateLineToken(c)
	if err != nil {
		return
	}

	response, err := validateSpecifiedResponse(c, lineToken)
	if err != nil {
		return
	}

	data.Db.Delete(&response)
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
