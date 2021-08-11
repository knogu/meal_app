package handler

import (
	"fmt"
	"meal_api/own_error"
	"net/http"

	"github.com/pkg/errors"
)

func handleError(w http.ResponseWriter, err error) {
	switch cause := errors.Cause(err).(type) {
	case own_error.BadRequestError:
		cause.Return(w)
	default:
		process500(w, err)
	}
}

func process500(w http.ResponseWriter, err error) {
	fmt.Printf("%+v\n", err)
	fmt.Printf("process500")
	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	return
}
