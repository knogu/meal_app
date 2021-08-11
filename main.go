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
	// user
	r.HandleFunc("/invited_users/{team_uuid}", handler.HandleInvitedUserPost)
	r.HandleFunc("/organizers", handler.HandleOrganizersPost)
	r.HandleFunc("/users/{user_id}", handler.HandleUsersPut)
	r.HandleFunc("/users/{user_id}/responses", handler.HandleResponses)

	http.ListenAndServe(":80", r)
}
