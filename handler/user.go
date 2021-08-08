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
	fmt.Println("handleUsersPost called")
	is_organizer := mux.Vars(r)["is_organizer"]
	// if  err != nil {
	// 	http.Error(w, err.Error(), 400)
	// }
	if is_organizer == "true" {

	}
}
