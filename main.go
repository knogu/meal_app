package main

import (
	"meal_api/data"
	"meal_api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.DBinit()

	r := gin.Default()
	// user
	r.POST("/invited_users/:team_uuid", handler.HandleInvitedUserPost)
	r.POST("/organizers", handler.HandleOrganizersPost)
	r.PUT("/users/:user_id", handler.HandleUsersPut)

	// responses
	r.POST("/users/:user_id/responses", handler.HandleResponsesPost)
	r.PUT("/users/:user_id/responses/:id", handler.HandleResponsesPut)

	r.Run(":80")
}
