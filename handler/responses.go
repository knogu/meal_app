package handler

import (
	"meal_api/data"
	"meal_api/json_structs"
	"meal_api/xer"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func HandleResponsesPost(c *gin.Context) {
	userIDByPath := c.Param("user_id")
	lineToken, err := validateLineToken(c)
	if err != nil {
		handleError(c, err)
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
		handleError(c, err)
		return
	}

	param, err := json_structs.NewUpdateResponseParams(c)
	if err != nil {
		handleError(c, err)
		return
	}
	response, err := validateSpecifiedResponse(c, lineToken)
	if err != nil {
		handleError(c, err)
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
		handleError(c, err)
		return
	}

	response, err := validateSpecifiedResponse(c, lineToken)
	if err != nil {
		handleError(c, err)
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

var (
	indexDaysDuration = 10
)

func HandleResponsesGet(c *gin.Context) {
	lineToken, err := validateLineToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	userIDByToken := data.FetchLineProfile(lineToken).LineID
	err = data.IsAuthorizedInTeam(userIDByToken, c.Query("user_id"))
	if err != nil {
		handleError(c, err)
		return
	}
	user, err := data.FetchUserById(userIDByToken)
	t, err := time.Parse(time.RFC3339, c.Query("from_date"))
	if err != nil {
		handleError(c, xer.Err4xx{ErrType: xer.TimeParseFailed})
		return
	}
	eventsJSON, err := user.EventsWithResponses(t, indexDaysDuration)
	if err != nil {
		handle500(c, err)
		return
	}

	c.JSON(http.StatusOK, eventsJSON)
	return
}

func HandleTeamResponsesGet(c *gin.Context) {
	lineToken, err := validateLineToken(c)
	if err != nil {
		handleError(c, err)
		return
	}
	userIDByToken := data.FetchLineProfile(lineToken).LineID
	user, err := data.FetchUserById(userIDByToken)
	if err != nil {
		handleError(c, err)
		return
	}
	team, err := data.FetchTeamByUUid(user.TeamUUID)
	if err != nil {
		handleError(c, err)
		return
	}
	t, err := time.Parse(time.RFC3339, c.Query("from_date"))
	if err != nil {
		handleError(c, xer.Err4xx{ErrType: xer.TimeParseFailed})
		return
	}

	eventsJSON, err := team.EventsWithResponses(t, indexDaysDuration)
	if err != nil {
		handle500(c, err)
		return
	}

	c.JSON(http.StatusOK, eventsJSON)
	return
}
