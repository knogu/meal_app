package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"errors"
)

var (
	IsOrganizerErr = errors.New("usersへのPOSTリクエストのパラメータ is_organizer を指定してください")
)

func HandleUsersPost(w http.ResponseWriter, r *http.Request) {
	var is_organizer bool
	var err error
	is_organizer, err = mux.Vars(r)["is_organizer"]
	if  err != nil {
		http.Error(w, err.Error(), 400)
	}
	if is_organizer {

	}
}
