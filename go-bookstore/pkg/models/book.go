package models

import (
	"github.com/jinzhu/gorm"

	"example/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Publication string
}

func init() {
	db = config.Connect()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var b Book
	db := db.Where("ID=?", ID).Find(&b)
	return &b, db
}

func DeleteBook(ID int64) Book {
	var b Book
	// db.Delete(b, db.Where("ID=?", ID)) // try this syntax
	db.Where("ID=?", ID).Delete(b)
	return b
}
