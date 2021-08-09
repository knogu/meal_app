package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"meal_api/data"
	"meal_api/handler"
)

func main() {
	data.DBinit()

	r := mux.NewRouter()
	r.HandleFunc("/invited_users/{team_uuid}", handler.HandleInvitedUserPost)
	r.HandleFunc("/organizers", handler.HandleOrganizersPost)

	http.ListenAndServe(":80", r)
}
