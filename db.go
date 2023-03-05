package main

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var Db *gorm.DB
var err error

func ConnectDB(){
	DSN:= os.Getenv("DBURL")
	Db,err=gorm.Open(postgres.Open(DSN),&gorm.Config{})
	if err!=nil{
		panic(err)
	}
	fmt.Println(" Db Successfully Connected")
	Db.AutoMigrate(Product{},Vendor{})
}