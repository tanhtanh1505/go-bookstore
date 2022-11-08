package models

import (
	"github.com/jinzhu/gorm"
	"github.com/tanhtanh1505/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
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

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID = ?", Id).Find(&book)
	return &book, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(&book)
	return book
}
