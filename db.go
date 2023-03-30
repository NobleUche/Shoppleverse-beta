package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB
var err error

func ConnectDB() {
	DSN := os.Getenv("DBURL")
	Db, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println(" Db Successfully Connected")
	Db.AutoMigrate(Product{}, Vendor{})
}
