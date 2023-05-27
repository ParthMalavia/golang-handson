package config

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() *gorm.DB {
	d, err := gorm.Open("mysql", "parth:Crest@123@tcp(127.0.0.1:3306)/testdb?parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("Connected to Database")
	return db
}
