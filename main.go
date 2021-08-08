package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"meal_api/handler"
	"meal_api/data"
)

func main() {
	data.DBinit()

	r := mux.NewRouter()
	r.HandleFunc("/invited_users", handler.HandleInvitedUserPost)
	r.HandleFunc("/organizers", handler.HandleOrganizersPost)

	http.ListenAndServe(":80", r)
}
