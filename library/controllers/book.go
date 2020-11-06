package controllers

import (
	"../models"
	bookRepository "../repository/book"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {}

var books []models.Book
var bookRepo bookRepository.BookRepository

func (c *Controller) Init() Controller {
	bookRepo = bookRepository.BookRepository{}
	return Controller{}
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}

		books = bookRepo.GetBooks(db, book, books)

		json.NewEncoder(w).Encode(books)
	}
}

func (c Controller) GetOneBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r * http.Request) {
		params := mux.Vars(r)
		var book models.Book
		i, _ := strconv.Atoi(params["id"])

		book = bookRepo.GetBook(db, book, i)

		json.NewEncoder(w).Encode(book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook models.Book
		var bookId int

		err := json.NewDecoder(r.Body).Decode(&newBook)
		logFatal(err)

		bookId = bookRepo.AddBook(db, newBook)

		json.NewEncoder(w).Encode(bookId)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newBook models.Book
		json.NewDecoder(r.Body).Decode(&newBook)

		rowsUpdated := bookRepo.UpdateBook(db, newBook)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		id, _ := strconv.Atoi(params["id"])

		rowsDeleted := bookRepo.RemoveBook(db, id)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}

