package main

import (
	"net/http"
	"github.com/gorilla/mux"

	"meal_api/handler"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.HandleUsers)

	http.ListenAndServe(":80",r)
}
