package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Book struct {
	Id int `gorm:"column:id"`
	Title string `gorm:"column:title"`
	Author string `gorm:"column:author"`
	Year string `gorm:"column:year"`
}

func addANewBook(db *gorm.DB) {
	newBook := Book{Title: "Povestea celor doua orase", Author: "Charles Dickenson", Year: "1890"}

	result := db.Create(&newBook)
	if result.Error != nil {
		fmt.Println(result.Error)
	}

	fmt.Println(newBook.Id)
}

func findAllBooks(db *gorm.DB) {
	var allBooks []Book
	//db.Table("books").Find(&books)
	db.Find(&allBooks)
	fmt.Print(allBooks)
}

func main() {
	dsn := "user=postgres password=postgres host=localhost dbname=golang-library port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	//addANewBook()

	findAllBooks(db)


	//var oneBook []Book
	//db.Where("title like ?", "%Games of thrones%").Find(&oneBook)
	//fmt.Println(oneBook)
}