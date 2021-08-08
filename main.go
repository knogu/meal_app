package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"log"

	"meal_api/handler"
	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

func DBinit() {
	dsn := "meal:password@tcp(db:3306)/meal?charset=utf8&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db connection succeeded")
	}
}

func main() {
	DBinit()

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.HandleUsersPost)

	http.ListenAndServe(":80",r)
}
