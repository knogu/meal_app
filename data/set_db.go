package data

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func DBinit() {
	dsn := "meal:password@tcp(Db:3306)/meal?charset=utf8&parseTime=True&loc=Local"
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Db connection succeeded")
	}

	Db.AutoMigrate(&Team{}, &User{})
}
