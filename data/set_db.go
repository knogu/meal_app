package data

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
func DBinit() {
	dsn := "meal:password@tcp(db:3306)/meal?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db connection succeeded")
	}

	db.AutoMigrate(&UserGroup{}, &User{})
	// db.Create(&User{LineID: "1", LineName: "kotaro", GroupUuId: "11", IsCook: false})
}
