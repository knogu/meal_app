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
	r.HandleFunc("/users", handler.HandleUsersPost)

	http.ListenAndServe(":80", r)
}
