package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shasikaud/go-app-01/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name   string `gorm:""json:"name"`
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
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
