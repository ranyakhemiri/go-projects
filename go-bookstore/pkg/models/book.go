package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ranyakhemiri/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	// gorm.model is a struct that contains the fields ID, CreatedAt, UpdatedAt, DeletedAt
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// initilaize the DB connection
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) createBook() *Book {
	// NewRecord is a gorm function, it avoids for us to write the SQL queries
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func getAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func getBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func deleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
