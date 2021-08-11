package handler

import (
	"meal_api/own_error"
	"net/http"
)

func HandleResponses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		handleResponsesPost(w, r)
	default:
		err := own_error.BadRequestError{Detail: own_error.MethodNotAllowed{Detail_: "Resource responses response to GET, POST, DELETE, PUT"}}
		handleError(w, err)
	}
}

func handleResponsesPost(w http.ResponseWriter, r *http.Request) {
	// user_id := mux.Vars(r)["user_id"]
}
